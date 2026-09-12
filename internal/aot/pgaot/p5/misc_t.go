package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TransferExpandedObject(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v8 != l1 {
		if v8 == int32(0) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
			if v13 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v12
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v12
			}
			if v12 == int32(0) {
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v18
			}
		}
		if l1 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = l1
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v25
			if v25 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v4
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(0)
		}
	} else {
	}
	return v3 + int32(12)
}
func F_TransferPredicateLocksToHeapRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v258 int64
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v323 int64
	_ = v323
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v23 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(48)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v26) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
	if v30 == int32(116) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v35 = v34
	goto L7
L6:
	;
	v35 = v26
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v42 = F_LWLockAcquire(m, v38+int32(3840), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v49 = F_LWLockAcquire(m, v45+int32(25344), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v56 = F_LWLockAcquire(m, v52+int32(25472), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v63 = F_LWLockAcquire(m, v59+int32(25600), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v70 = F_LWLockAcquire(m, v66+int32(25728), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v77 = F_LWLockAcquire(m, v73+int32(25856), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v84 = F_LWLockAcquire(m, v80+int32(25984), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v91 = F_LWLockAcquire(m, v87+int32(26112), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v98 = F_LWLockAcquire(m, v94+int32(26240), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v105 = F_LWLockAcquire(m, v101+int32(26368), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v112 = F_LWLockAcquire(m, v108+int32(26496), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v119 = F_LWLockAcquire(m, v115+int32(26624), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v126 = F_LWLockAcquire(m, v122+int32(26752), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v133 = F_LWLockAcquire(m, v129+int32(26880), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v140 = F_LWLockAcquire(m, v136+int32(27008), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v147 = F_LWLockAcquire(m, v143+int32(27136), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v154 = F_LWLockAcquire(m, v150+int32(27264), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v161 = F_LWLockAcquire(m, v157+int32(3584), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	v171 = F_hash_search_with_hash_value(m, v164, int32(1630556), v167, int32(2), v19+int32(28))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	F_hash_seq_init(m, v19+int32(28), v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v181 = F_hash_seq_search(m, v19+int32(28))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	if v181 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v186 = v2
	v187 = v181
	v193 = v2
	goto L33
L31:
	;
	goto L32
L32:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v396 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	v400 = F_hash_search_with_hash_value(m, v393, int32(1630556), v396, int32(1), v19+int32(8))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L8
	} else {
		goto L70
	}
L33:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v199 != v26 {
		v359 = v186
		v366 = v193
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v374 = F_hash_seq_search(m, v19+int32(28))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L8
	} else {
		goto L68
	}
L36:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v201 != v36 {
		v359 = v186
		v366 = v193
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if v33 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v186 != 0 {
		v232 = v186
		v233 = v193
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	if v203 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v204 == int32(-1) {
		v359 = v186
		v366 = v193
		goto L35
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v187)+20))
	if v234 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v36
	v212 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v215 = F_get_hash_value(m, v212, v19+int32(8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v224 = F_hash_search_with_hash_value(m, v218, v19+int32(8), v215, int32(1), v19+int32(27))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+27)))
	if v226 != 0 {
		v232 = v224
		v233 = v215
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v228 = v224 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v224)+16)) = v228
	v232 = v224
	v233 = v215
	goto L42
L47:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v354 = F_hash_search(m, v350, v187, int32(2), v19+int32(27))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L8
	} else {
		goto L67
	}
L48:
	;
	v238 = v187 + int32(16)
	if v234 == v238 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v241 = v232 + int32(16)
	v250 = v234
	goto L50
L50:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v250)+16))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v250-int32(4))))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v263 = int32(8)
	v264 = v250 + v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+4)) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	*(*int32)(unsafe.Add(mBase, uint32(v266))) = v268
	v271 = *(*int32)(unsafe.Add(mBase, _consts[792]))
	v277 = F_hash_search(m, v271, v250-v263, int32(2), v19+int32(27))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v232
	v282 = *(*int32)(unsafe.Add(mBase, _consts[792]))
	v291 = F_hash_search_with_hash_value(m, v282, v19+int32(8), v261<<(uint(int32(4))%32)^v233, int32(1), v19+int32(27))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+27)))
	if v293 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v262 != v238 {
		v250 = v262
		goto L50
	} else {
		goto L66
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v291)+24)) = v258
	goto L54
L56:
	;
	v297 = v291 + int32(8)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v232)+20))
	if v298 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v291)+24))
	if base.Ui64(v258) <= base.Ui64(v323) {
		goto L54
	} else {
		goto L65
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = v241
	goto L61
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+12)) = v241
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v304)+4)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v297
	v309 = v291 + int32(16)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v312 = v310 + int32(48)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v310)+52))
	if v313 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+52)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v310)+48)) = v312
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+20)) = v312
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v309
	goto L55
L65:
	;
	goto L55
L66:
	;
	goto L51
L67:
	;
	v359 = v232
	v366 = v233
	goto L35
L68:
	;
	if v374 != 0 {
		v186 = v359
		v187 = v374
		v193 = v366
		goto L33
	} else {
		goto L69
	}
L69:
	;
	goto L34
L70:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v403+int32(3584))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v409+int32(27264))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v415+int32(27136))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v421+int32(27008))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v427+int32(26880))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v433+int32(26752))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v439+int32(26624))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v445+int32(26496))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v451+int32(26368))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v457+int32(26240))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v463+int32(26112))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v469+int32(25984))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v475+int32(25856))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v481+int32(25728))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v487+int32(25600))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v493+int32(25472))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v499+int32(25344))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v505+int32(3840))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	goto L1
}
func F___tan(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v77 int32
	_ = v77
	var v81 float64
	_ = v81
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v95 int64
	_ = v95
	var v97 float64
	_ = v97
	var v101 float64
	_ = v101
	var v113 float64
	_ = v113
	v7 = int32(0)
	v9 = base.I64_reinterpret_f64(l0)
	v13 = base.B2i32(base.Ui64(v9&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
	if v13 == v7 {
		v22 = base.B2i32(int64(0) <= v9)
		if int64(0) <= v9 {
			v23 = l1
		} else {
			v23 = base.F64_neg(l1)
		}
		v27 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(l0)), base.F64_sub(float64(3.061616997868383e-17), v23))
		v28 = float64(0)
		v29 = v22
	} else {
		v27 = l0
		v28 = l1
		v29 = v7
	}
	v30 = base.F64_mul(v27, v27)
	v31 = base.F64_mul(v27, v30)
	v34 = base.F64_mul(v30, v30)
	v73 = base.F64_add(base.F64_mul(v31, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v30, base.F64_add(base.F64_mul(v31, base.F64_add(base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v30, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, base.F64_add(base.F64_mul(v34, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v28)), v28))
	v74 = base.F64_add(v27, v73)
	if v13 == int32(0) {
		v77 = int32(1)
		v81 = base.F64_convert_i32_s(v77 - l2<<(uint(v77)%32))
		v86 = base.F64_add(v27, base.F64_sub(v73, base.F64_div(base.F64_mul(v74, v74), base.F64_add(v74, v81))))
		v88 = base.F64_sub(v81, base.F64_add(v86, v86))
		if v29 != 0 {
			v90 = v88
		} else {
			v90 = base.F64_neg(v88)
		}
		return v90
	} else {
		if l2 != 0 {
			v93 = base.F64_div(float64(-1), v74)
			v95 = int64(-4294967296)
			v97 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v93) & v95)
			v101 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v74) & v95)
			v113 = base.F64_add(base.F64_mul(v93, base.F64_add(base.F64_mul(v97, base.F64_sub(v73, base.F64_sub(v101, v27))), base.F64_add(base.F64_mul(v97, v101), float64(1)))), v97)
		} else {
			v113 = v74
		}
		return v113
	}
}
func F___trunctfsf2(m *base.Module, l0 int64, l1 int64) float32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v99 int64
	_ = v99
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = l1 & int64(281474976710655)
	v19 = int64(base.Ui64(l1)>>(uint(int64(48))%64)) & int64(32767)
	v20 = base.I32_wrap_i64(v19)
	if base.Ui32(v20-int32(16257)) <= base.Ui32(int32(253)) {
		v27 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(int64(25)) % 64)))
		v31 = l1 & int64(33554431)
		v32 = int64(16777216)
		if v31 == v32 {
			v36 = base.B2i32(l0 == int64(0))
		} else {
			v36 = base.B2i32(base.Ui64(v31) < base.Ui64(v32))
		}
		if v36 == int32(0) {
			v49 = v27 + int32(1)
		} else {
			if l0|(v31^int64(16777216)) != int64(0) {
				v49 = v27
			} else {
				v49 = v27&int32(1) + v27
			}
		}
		v52 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v49))
		if base.Ui32(int32(8388607)) < base.Ui32(v49) {
			v53 = int32(0)
		} else {
			v53 = v49
		}
		if base.Ui32(int32(8388607)) < base.Ui32(v49) {
			v56 = int32(-16255)
		} else {
			v56 = int32(-16256)
		}
		v177 = v53
		v178 = v56 + v20
	} else {
		if l0|v15 == int64(0) {
			if base.Ui32(int32(16510)) < base.Ui32(v20) {
				v177 = int32(0)
				v178 = int32(255)
			} else {
				v75 = base.B2i32(v19 == int64(0))
				if v19 == int64(0) {
					v76 = int32(16256)
				} else {
					v76 = int32(16257)
				}
				v77 = v76 - v20
				if int32(112) < v77 {
					v80 = int32(0)
					v177 = v80
					v178 = v80
				} else {
					v83 = v12 + int32(16)
					if v19 == int64(0) {
						v86 = v15
					} else {
						v86 = v15 | int64(281474976710656)
					}
					v88 = int32(128) - v77
					if v88&int32(64) != 0 {
						v107 = int64(0)
						v108 = l0 << (uint(base.I64_extend_i32_u(v88+int32(-64))) % 64)
					} else {
						if v88 == int32(0) {
							v107 = l0
							v108 = v86
						} else {
							v99 = base.I64_extend_i32_u(v88)
							v107 = l0 << (uint(v99) % 64)
							v108 = v86<<(uint(v99)%64) | int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v88))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v83))) = v107
					*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = v108
					if v77&int32(64) != 0 {
						v130 = int64(base.Ui64(v86) >> (uint(base.I64_extend_i32_u(v77+int32(-64))) % 64))
						v131 = int64(0)
					} else {
						if v77 == int32(0) {
							v130 = l0
							v131 = v86
						} else {
							v126 = base.I64_extend_i32_u(v77)
							v130 = v86<<(uint(base.I64_extend_i32_u(int32(64)-v77))%64) | int64(base.Ui64(l0)>>(uint(v126)%64))
							v131 = int64(base.Ui64(v86) >> (uint(v126) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v130
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v131
					v135 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					v138 = base.I32_wrap_i64(int64(base.Ui64(v135) >> (uint(int64(25)) % 64)))
					v139 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					v141 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					v142 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
					v144 = int64(0)
					v148 = v139 | base.I64_extend_i32_u(base.B2i32(v20 != v76)&base.B2i32(v141|v142 != v144))
					v152 = v135 & int64(33554431)
					v153 = int64(16777216)
					if v152 == v153 {
						v157 = base.B2i32(v148 == v144)
					} else {
						v157 = base.B2i32(base.Ui64(v152) < base.Ui64(v153))
					}
					if v157 == int32(0) {
						v170 = v138 + int32(1)
					} else {
						if v148|(v152^int64(16777216)) != int64(0) {
							v170 = v138
						} else {
							v170 = v138&int32(1) + v138
						}
					}
					v174 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v170))
					if base.Ui32(int32(8388607)) < base.Ui32(v170) {
						v175 = v170 ^ int32(8388608)
					} else {
						v175 = v170
					}
					v177 = v175
					v178 = v174
				}
			}
		} else {
			if v19 != int64(32767) {
				if base.Ui32(int32(16510)) < base.Ui32(v20) {
					v177 = int32(0)
					v178 = int32(255)
				} else {
					v75 = base.B2i32(v19 == int64(0))
					if v19 == int64(0) {
						v76 = int32(16256)
					} else {
						v76 = int32(16257)
					}
					v77 = v76 - v20
					if int32(112) < v77 {
						v80 = int32(0)
						v177 = v80
						v178 = v80
					} else {
						v83 = v12 + int32(16)
						if v19 == int64(0) {
							v86 = v15
						} else {
							v86 = v15 | int64(281474976710656)
						}
						v88 = int32(128) - v77
						if v88&int32(64) != 0 {
							v107 = int64(0)
							v108 = l0 << (uint(base.I64_extend_i32_u(v88+int32(-64))) % 64)
						} else {
							if v88 == int32(0) {
								v107 = l0
								v108 = v86
							} else {
								v99 = base.I64_extend_i32_u(v88)
								v107 = l0 << (uint(v99) % 64)
								v108 = v86<<(uint(v99)%64) | int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v88))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v83))) = v107
						*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = v108
						if v77&int32(64) != 0 {
							v130 = int64(base.Ui64(v86) >> (uint(base.I64_extend_i32_u(v77+int32(-64))) % 64))
							v131 = int64(0)
						} else {
							if v77 == int32(0) {
								v130 = l0
								v131 = v86
							} else {
								v126 = base.I64_extend_i32_u(v77)
								v130 = v86<<(uint(base.I64_extend_i32_u(int32(64)-v77))%64) | int64(base.Ui64(l0)>>(uint(v126)%64))
								v131 = int64(base.Ui64(v86) >> (uint(v126) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v130
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v131
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
						v138 = base.I32_wrap_i64(int64(base.Ui64(v135) >> (uint(int64(25)) % 64)))
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v142 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
						v144 = int64(0)
						v148 = v139 | base.I64_extend_i32_u(base.B2i32(v20 != v76)&base.B2i32(v141|v142 != v144))
						v152 = v135 & int64(33554431)
						v153 = int64(16777216)
						if v152 == v153 {
							v157 = base.B2i32(v148 == v144)
						} else {
							v157 = base.B2i32(base.Ui64(v152) < base.Ui64(v153))
						}
						if v157 == int32(0) {
							v170 = v138 + int32(1)
						} else {
							if v148|(v152^int64(16777216)) != int64(0) {
								v170 = v138
							} else {
								v170 = v138&int32(1) + v138
							}
						}
						v174 = base.B2i32(base.Ui32(int32(8388607)) < base.Ui32(v170))
						if base.Ui32(int32(8388607)) < base.Ui32(v170) {
							v175 = v170 ^ int32(8388608)
						} else {
							v175 = v170
						}
						v177 = v175
						v178 = v174
					}
				}
			} else {
				v177 = base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(25))%64))) | int32(4194304)
				v178 = int32(255)
			}
		}
	}
	m.G0 = v12 + int32(32)
	return base.F32_reinterpret_i32(base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(32))%64)))&int32(-2147483648) | v178<<(uint(int32(23))%32) | v177)
}
func F_tag_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	v8 = l1 - int32(1636608432)
	if l0&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(l1) {
			v117 = l0
			v118 = l1
			v119 = v8
			v120 = v8
			v121 = v8
			for {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
				v124 = v123 + v120
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
				v128 = v127 + v121
				v130 = int32(4)
				v132 = v125 + v119 - v128 ^ base.I32_rotl(v128, v130)
				v136 = v124 - v132 ^ base.I32_rotl(v132, int32(6))
				v137 = v128 + v124
				v138 = v132 + v137
				v139 = v136 + v138
				v143 = v137 - v136 ^ base.I32_rotl(v136, int32(8))
				v147 = v138 - v143 ^ base.I32_rotl(v143, int32(16))
				v151 = v139 - v147 ^ base.I32_rotl(v147, int32(19))
				v152 = v143 + v139
				v153 = v147 + v152
				v154 = v151 + v153
				v158 = v152 - v151 ^ base.I32_rotl(v151, v130)
				v159 = int32(12)
				v160 = v117 + v159
				v162 = v118 - v159
				if base.Ui32(int32(11)) < base.Ui32(v162) {
					v117 = v160
					v118 = v162
					v119 = v153
					v120 = v154
					v121 = v158
					continue
				} else {
					break
				}
				break
			}
			v165 = v160
			v166 = v162
			v167 = v153
			v168 = v154
			v169 = v158
		} else {
			v165 = l0
			v166 = l1
			v167 = v8
			v168 = v8
			v169 = v8
		}
		switch v166 - int32(1) {
		case 0:
			v228 = v167
			v229 = v168
			v230 = v169
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 1:
			v221 = v167
			v222 = v168
			v223 = v169
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 2:
			v214 = v167
			v215 = v168
			v216 = v169
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 3:
			v208 = v168
			v209 = v169
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 4:
			v204 = v168
			v205 = v169
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 5:
			v198 = v168
			v199 = v169
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 6:
			v192 = v168
			v193 = v169
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 7:
			v187 = v169
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 8:
			v182 = v169
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 9:
			v177 = v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 10:
			v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+10)))
			v177 = v173<<(uint(int32(24))%32) + v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		default:
			v235 = v167
			v236 = v168
			v237 = v169
		}
	} else {
		if base.Ui32(l1) < base.Ui32(int32(12)) {
			v63 = l0
			v64 = l1
			v65 = v8
			v66 = v8
			v67 = v8
		} else {
			v15 = l0
			v16 = l1
			v17 = v8
			v18 = v8
			v19 = v8
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v22 = v21 + v18
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v26 = v25 + v19
				v28 = int32(4)
				v30 = v23 + v17 - v26 ^ base.I32_rotl(v26, v28)
				v34 = v22 - v30 ^ base.I32_rotl(v30, int32(6))
				v35 = v26 + v22
				v36 = v30 + v35
				v37 = v34 + v36
				v41 = v35 - v34 ^ base.I32_rotl(v34, int32(8))
				v45 = v36 - v41 ^ base.I32_rotl(v41, int32(16))
				v49 = v37 - v45 ^ base.I32_rotl(v45, int32(19))
				v50 = v41 + v37
				v51 = v45 + v50
				v52 = v49 + v51
				v56 = v50 - v49 ^ base.I32_rotl(v49, v28)
				v57 = int32(12)
				v58 = v15 + v57
				v60 = v16 - v57
				if base.Ui32(int32(11)) < base.Ui32(v60) {
					v15 = v58
					v16 = v60
					v17 = v51
					v18 = v52
					v19 = v56
					continue
				} else {
					break
				}
				break
			}
			v63 = v58
			v64 = v60
			v65 = v51
			v66 = v52
			v67 = v56
		}
		switch v64 - int32(1) {
		case 0:
			v114 = v65
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 1:
			v109 = v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 2:
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
			v109 = v105<<(uint(int32(16))%32) + v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 3:
			v102 = v66
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 4:
			v99 = v66
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 5:
			v94 = v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 6:
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
			v94 = v90<<(uint(int32(16))%32) + v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 7:
			v85 = v67
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 8:
			v80 = v67
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 9:
			v75 = v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 10:
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)))
			v75 = v71<<(uint(int32(24))%32) + v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		default:
			v235 = v65
			v236 = v66
			v237 = v67
		}
	}
	v240 = int32(14)
	v242 = v236 ^ v237 - base.I32_rotl(v236, v240)
	v246 = v242 ^ v235 - base.I32_rotl(v242, int32(11))
	v250 = v246 ^ v236 - base.I32_rotl(v246, int32(25))
	v254 = v250 ^ v242 - base.I32_rotl(v250, int32(16))
	v258 = v254 ^ v246 - base.I32_rotl(v254, int32(4))
	v262 = v258 ^ v250 - base.I32_rotl(v258, v240)
	return v262 ^ v254 - base.I32_rotl(v262, int32(24))
}
func F_textout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v12 = int32(4)
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v14&int32(254) == int32(2) {
				v23 = v12
			} else {
				v23 = base.B2i32(v14 == int32(18)) << (uint(v12) % 32)
			}
			if v14 == int32(1) {
				v26 = v12
			} else {
				v26 = v23
			}
			v39 = v26
		} else {
			v27 = int32(1)
			if v9&v27 != 0 {
				v39 = int32(base.Ui32(v9)>>(uint(v27)%32)) - v27
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v42 = F_palloc(m, v39+int32(1))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v44 = int32(1)
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
			if v46&v44 != 0 {
				v49 = v44
			} else {
				v49 = int32(4)
			}
			if v39 != 0 {
				v51 = F__emscripten_memcpy_bulkmem(m, v42, v5+v49, v39)
				mBase = m.M
				v52 = v51
			} else {
				v52 = v42
			}
			v54 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v39+v52))) = uint8(v54)
			if v5 != v4 {
				F_pfree(m, v5)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					return v52
				}
			} else {
				return v52
			}
		}
	}
}
func F_texttoxml(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_xmlparse(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_timestamp2tm(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v145 int64
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = base.I64_div_s(l0, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(l0+int64(86399999999)) {
		v25 = v17 * int64(-86400000000)
	} else {
		v25 = int64(0)
	}
	v26 = v25 + l0
	v29 = v26>>(uint(int64(63))%64) + v17
	if v29 < int64(-2451545) {
		v190 = int32(-1)
		m.G0 = v13 + int32(16)
		return v190
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, _consts[330]))
		v34 = base.I32_wrap_i64(v29)
		v46 = v34 + int32(2483589)
		v47 = int32(146097)
		v48 = base.I32_div_u_s(v46, v47)
		v49 = int32(3)
		v55 = int32(2)
		v60 = base.I32_div_u_s((v48*int32(1073595727)+v46)<<(uint(v55)%32)|v49, v47)
		v63 = v34 + int32(2451545) + v48*v49 + v60 + int32(32104)
		v64 = int32(1461)
		v65 = base.I32_div_u_s(v63, v64)
		v68 = v65*int32(-1461) + v63
		v70 = v68 << (uint(v55) % 32)
		if base.Ui32(v64) <= base.Ui32(v70) {
			v76 = base.I32_rem_u_s(v68+int32(305), int32(365))
			v81 = v76
		} else {
			v80 = base.I32_rem_u_s(v68+int32(306), int32(366))
			v81 = v80
		}
		v83 = base.I32_div_u_s(v70, int32(1461))
		*(*int32)(unsafe.Add(mBase, uint32(l2+int32(20)))) = v83 + v65<<(uint(int32(2))%32) - int32(4800)
		v91 = v81 + int32(123)
		v95 = int32(base.Ui32(v91*int32(2141)) >> (uint(int32(16)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(l2+int32(12)))) = v91 - int32(base.Ui32(v95*int32(7834))>>(uint(int32(8))%32))
		v105 = base.I32_rem_u_s(v95+int32(10), int32(12))
		*(*int32)(unsafe.Add(mBase, uint32(l2+int32(16)))) = v105 + int32(1)
		if v26 < int64(0) {
			v113 = v26 + int64(86400000000)
		} else {
			v113 = v26
		}
		v115 = base.I64_div_s(v113, int64(3600000000))
		*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)) = uint32(v115)
		v120 = base.I64_extend32_s(v115)*int64(-3600000000) + v113
		v122 = base.I64_div_s(v120, int64(60000000))
		*(*uint32)(unsafe.Add(mBase, uint32(l2)+4)) = uint32(v122)
		v127 = base.I64_extend32_s(v122)*int64(-60000000) + v120
		v129 = base.I64_div_s(v127, int64(1000000))
		*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v129)
		v133 = v129*int64(4293967296) + v127
		*(*uint32)(unsafe.Add(mBase, uint32(l3))) = uint32(v133)
		if l1 == int32(0) {
			v137 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v137
			*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = int64(4294967295)
			if l4 != 0 {
				v186 = v137
				v187 = v137
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v187
				v190 = v186
			} else {
				v190 = v137
			}
			m.G0 = v13 + int32(16)
			return v190
		} else {
			v145 = base.I64_div_s(l0-base.I64_extend32_s(v133), int64(1000000))
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v145 + int64(946684800)
			if l5 != 0 {
				v151 = l5
			} else {
				v151 = v33
			}
			v152 = F_pg_localtime(m, v13+int32(8), v151)
			mBase = m.M
			v155 = m.ExcPending
			if v155 != 0 {
				return int32(0)
			} else {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v156 + int32(1900)
				v160 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v160 + int32(1)
				v164 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v164
				v166 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v166
				v168 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v168
				v170 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v170
				v172 = *(*int32)(unsafe.Add(mBase, uint32(v152)+32))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v172
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v152)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v174
				v176 = *(*int32)(unsafe.Add(mBase, uint32(v152)+40))
				*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v176
				v178 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v178 - v174
				if l4 == v178 {
					v190 = v178
				} else {
					v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
					v186 = v178
					v187 = v184
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v187
					v190 = v186
				}
				m.G0 = v13 + int32(16)
				return v190
			}
		}
	}
}
func F_timetypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(222910), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499833), int32(65), int32(278037))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v37 = F_anytime_typmod_check(m, int32(0), v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_tlist_same_datatypes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v4 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = v8
	goto L3
L2:
	;
	v10 = int32(0)
	goto L3
L3:
	;
	if l0 == int32(0) {
		v59 = v10
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return base.B2i32(v59 == int32(0))
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= int32(0) {
		v59 = v10
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v19 = v10
	v20 = v4
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v20<<(uint(int32(2))%32))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+26)))
	if v28 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v59 = v51
	goto L4
L9:
	;
	v53 = v20 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 < v54 {
		v19 = v51
		v20 = v53
		goto L7
	} else {
		goto L22
	}
L10:
	;
	return v4
L11:
	;
	if v19 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l2 != 0 {
		v51 = v19
		goto L9
	} else {
		goto L21
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v34 = F_exprType(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v34 != v38 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v41 = v19 + int32(4)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v41) < base.Ui32(v43+v44<<(uint(int32(2))%32)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v41
	goto L20
L19:
	;
	v49 = int32(0)
	goto L20
L20:
	;
	v51 = v49
	goto L9
L21:
	;
	goto L10
L22:
	;
	goto L8
}
func F_tm2timestamp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v95 int64
	_ = v95
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v132 int32
	_ = v132
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v156 int32
	_ = v156
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 <= int32(-4713) {
		if v15 != int32(-4713) {
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
			v156 = int32(-1)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if int32(10) < v20 {
				v31 = v20
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v37 = base.B2i32(int32(2) < v31)
				if int32(2) < v31 {
					v38 = int32(4800)
				} else {
					v38 = int32(4799)
				}
				v39 = v38 + v15
				v44 = base.I32_div_s(v39, int32(4))
				v47 = base.I32_div_s(v39, int32(-100))
				v50 = base.I32_div_s(v39, int32(400))
				if int32(2) < v31 {
					v54 = int32(1)
				} else {
					v54 = int32(13)
				}
				v59 = base.I32_div_s((v54+v31)*int32(7834), int32(256))
				v65 = base.I64_extend_i32_s(v32 + v39*int32(365) + v44 + v47 + v50 + v59 - int32(32167) - int32(2451545))
				v74 = int64(32)
				v75 = int64(20)
				v77 = int64(base.Ui64(v65) >> (uint(v74) % 64))
				v80 = int64(4294967295)
				v81 = int64(500654080)
				v83 = v65 & v80
				v84 = v81 * v83
				v88 = int64(base.Ui64(v84)>>(uint(v74)%64)) + v81*v77
				v95 = v83*v75 + v88&v80
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v65*int64(0) + v65>>(uint(int64(63))%64)*int64(86400000000) + v75*v77 + int64(base.Ui64(v88)>>(uint(v74)%64)) + int64(base.Ui64(v95)>>(uint(v74)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v13))) = v84&v80 | v95<<(uint(v74)%64)
				v106 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
				if v106 != v107>>(uint(int64(63))%64) {
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
					v156 = int32(-1)
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v115 = int32(60)
					v124 = base.I64_extend_i32_s(l1) + base.I64_extend_i32_s(v112+(v113+v114*v115)*v115)*int64(1000000)
					v125 = v107 + v124
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = v125
					if base.B2i32(v124 < int64(0))^base.B2i32(v125 < v107) != 0 {
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
						v156 = int32(-1)
					} else {
						if l2 != 0 {
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v137 = base.I64_extend_i32_s(int32(0)-v132)*int64(-1000000) + v125
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v137
							v139 = v137
						} else {
							v139 = v125
						}
						if base.Ui64(v139+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
							v156 = int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
							v156 = int32(-1)
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
				v156 = int32(-1)
			}
		}
	} else {
		if v15 <= int32(5874897) {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v31 = v25
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v37 = base.B2i32(int32(2) < v31)
			if int32(2) < v31 {
				v38 = int32(4800)
			} else {
				v38 = int32(4799)
			}
			v39 = v38 + v15
			v44 = base.I32_div_s(v39, int32(4))
			v47 = base.I32_div_s(v39, int32(-100))
			v50 = base.I32_div_s(v39, int32(400))
			if int32(2) < v31 {
				v54 = int32(1)
			} else {
				v54 = int32(13)
			}
			v59 = base.I32_div_s((v54+v31)*int32(7834), int32(256))
			v65 = base.I64_extend_i32_s(v32 + v39*int32(365) + v44 + v47 + v50 + v59 - int32(32167) - int32(2451545))
			v74 = int64(32)
			v75 = int64(20)
			v77 = int64(base.Ui64(v65) >> (uint(v74) % 64))
			v80 = int64(4294967295)
			v81 = int64(500654080)
			v83 = v65 & v80
			v84 = v81 * v83
			v88 = int64(base.Ui64(v84)>>(uint(v74)%64)) + v81*v77
			v95 = v83*v75 + v88&v80
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v65*int64(0) + v65>>(uint(int64(63))%64)*int64(86400000000) + v75*v77 + int64(base.Ui64(v88)>>(uint(v74)%64)) + int64(base.Ui64(v95)>>(uint(v74)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v13))) = v84&v80 | v95<<(uint(v74)%64)
			v106 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
			v107 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
			if v106 != v107>>(uint(int64(63))%64) {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
				v156 = int32(-1)
			} else {
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v115 = int32(60)
				v124 = base.I64_extend_i32_s(l1) + base.I64_extend_i32_s(v112+(v113+v114*v115)*v115)*int64(1000000)
				v125 = v107 + v124
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v125
				if base.B2i32(v124 < int64(0))^base.B2i32(v125 < v107) != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
					v156 = int32(-1)
				} else {
					if l2 != 0 {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v137 = base.I64_extend_i32_s(int32(0)-v132)*int64(-1000000) + v125
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v137
						v139 = v137
					} else {
						v139 = v125
					}
					if base.Ui64(v139+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
						v156 = int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
						v156 = int32(-1)
					}
				}
			}
		} else {
			if v15 != int32(5874898) {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
				v156 = int32(-1)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if int32(5) < v28 {
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
					v156 = int32(-1)
				} else {
					v31 = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v37 = base.B2i32(int32(2) < v31)
					if int32(2) < v31 {
						v38 = int32(4800)
					} else {
						v38 = int32(4799)
					}
					v39 = v38 + v15
					v44 = base.I32_div_s(v39, int32(4))
					v47 = base.I32_div_s(v39, int32(-100))
					v50 = base.I32_div_s(v39, int32(400))
					if int32(2) < v31 {
						v54 = int32(1)
					} else {
						v54 = int32(13)
					}
					v59 = base.I32_div_s((v54+v31)*int32(7834), int32(256))
					v65 = base.I64_extend_i32_s(v32 + v39*int32(365) + v44 + v47 + v50 + v59 - int32(32167) - int32(2451545))
					v74 = int64(32)
					v75 = int64(20)
					v77 = int64(base.Ui64(v65) >> (uint(v74) % 64))
					v80 = int64(4294967295)
					v81 = int64(500654080)
					v83 = v65 & v80
					v84 = v81 * v83
					v88 = int64(base.Ui64(v84)>>(uint(v74)%64)) + v81*v77
					v95 = v83*v75 + v88&v80
					*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v65*int64(0) + v65>>(uint(int64(63))%64)*int64(86400000000) + v75*v77 + int64(base.Ui64(v88)>>(uint(v74)%64)) + int64(base.Ui64(v95)>>(uint(v74)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v13))) = v84&v80 | v95<<(uint(v74)%64)
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
					if v106 != v107>>(uint(int64(63))%64) {
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
						v156 = int32(-1)
					} else {
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v115 = int32(60)
						v124 = base.I64_extend_i32_s(l1) + base.I64_extend_i32_s(v112+(v113+v114*v115)*v115)*int64(1000000)
						v125 = v107 + v124
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v125
						if base.B2i32(v124 < int64(0))^base.B2i32(v125 < v107) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
							v156 = int32(-1)
						} else {
							if l2 != 0 {
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v137 = base.I64_extend_i32_s(int32(0)-v132)*int64(-1000000) + v125
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v137
								v139 = v137
							} else {
								v139 = v125
							}
							if base.Ui64(v139+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
								v156 = int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
								v156 = int32(-1)
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v13 + int32(16)
	return v156
}
func F_tokenize_include_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_AbsoluteConfigLocation(m, l1, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_open_auth_file(m, v13, l3, l4, l6)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				if l5 == int32(0) {
					F_pfree(m, v13)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[86]))
					if v22 != int32(44) {
						F_pfree(m, v13)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							m.G0 = v11 + int32(16)
							return
						}
					} else {
						v26 = F_errstart(m, l3, int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v26 == int32(0) {
								v48 = l6
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(0)
								F_pfree(m, v13)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
								F_errmsg(m, int32(717524), v11)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									F_errfinish(m, int32(501351), int32(460), int32(387681))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										v48 = l6
										*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(0)
										F_pfree(m, v13)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											m.G0 = v11 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_tokenize_auth_file(m, v13, v15, l2, l3, l4)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = F_FreeFile(m, v15)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if l4 != 0 {
							F_pfree(m, v13)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v11 + int32(16)
								return
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, _consts[377]))
							F_MemoryContextDelete(m, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								v48 = int32(4419156)
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(0)
								F_pfree(m, v13)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_toupper(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
		v8 = l0 & int32(95)
	} else {
		v8 = l0
	}
	return v8
}
func F_trackitem_compare_element(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1037]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_FunctionCall2Coll(m, v6, v7, v9, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_trackitem_compare_frequencies_desc_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	return v5 - v7
}
func F_transformAExprOp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[263])))
	if v9 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v106 = F_transformExprRecurse(m, l0, v7)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L19
	} else {
		goto L37
	}
L2:
	;
	if v72 != int32(36) {
		goto L1
	} else {
		goto L28
	}
L3:
	;
	if v7 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 != int32(1) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v21 != int32(61) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v24 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v7 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v45 == int32(34) {
		v72 = v44
		goto L2
	} else {
		goto L18
	}
L10:
	;
	if v6 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L11:
	;
	v27 = int32(72)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v28 != v27 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v31 != 0 {
		v44 = v27
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v35 != int32(72) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)))
	if v38 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v41 == int32(34) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v44 = v41
	goto L9
L18:
	;
	v49 = F_palloc0(m, int32(20))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(52)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v59 == int32(72) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v63
	v65 = F_transformExprRecurse(m, l0, v49)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v62 != 0 {
		v63 = v6
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v63 = v7
	goto L21
L25:
	;
	goto L24
L26:
	;
	return v65
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v72 = v71
	goto L2
L28:
	;
	if v6 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	switch v77 - int32(22) {
	case 0:
		goto L31
	default:
		goto L1
	case 14:
		goto L30
	}
L30:
	;
	v93 = F_transformExprRecurse(m, l0, v7)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L34
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v80 != int32(4) {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(3)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v88
	v90 = F_transformExprRecurse(m, l0, v6)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	return v90
L34:
	;
	v95 = F_transformExprRecurse(m, l0, v6)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v101 = F_make_row_comparison_op(m, l0, v97, v98, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	return v101
L37:
	;
	v108 = F_transformExprRecurse(m, l0, v6)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v112 = F_make_op(m, l0, v110, v106, v108, v105, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	return v112
}
func F_transformAssignmentIndirection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	v13 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	if l1 != 0 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v20 + int32(112)
	return v346
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l2
	F_errmsg(m, int32(192436), v20)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L14
	} else {
		goto L86
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L14
	} else {
		goto L81
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L14
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L14
	} else {
		goto L69
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L64
	}
L7:
	;
	v194 = F_exprType(m, l9)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L14
	} else {
		goto L51
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v45 < v46 {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	if l7 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v44 = v36
	v45 = (l8 - v37) >> (uint(int32(2)) % 32)
	goto L8
L11:
	;
	if l8 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L12:
	;
	if l8 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = F_palloc0(m, int32(16))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(34)
	v36 = v25
	goto L10
L16:
	;
	v36 = l1
	goto L10
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v44 = l1
	v45 = v43
	goto L8
L18:
	;
	v56 = v45
	v61 = v13
	goto L21
L19:
	;
	v168 = v13
	goto L20
L20:
	;
	if v168 == int32(0) {
		goto L7
	} else {
		goto L49
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v68 = v65 + v56<<(uint(int32(2))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != int32(78) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v168 = v149
	goto L20
L23:
	;
	if v70 == int32(77) {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v149 = F_lappend(m, v61, v69)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L14
	} else {
		goto L47
	}
L26:
	;
	if v61 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v75 = F_transformAssignmentSubscripts(m, l0, v44, l2, l4, l5, l6, v61, l7, v68, l9, l10, l11)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = l5
	v80 = F_getBaseTypeAndTypmod(m, l4, v20+int32(108))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L14
	} else {
		goto L31
	}
L30:
	;
	v346 = v75
	goto L1
L31:
	;
	v82 = F_typeidTypeRelid(m, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	if v82 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v87 = F_get_attnum(m, v82, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	if v87 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v87 < int32(0) {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	F_get_atttypetypmodcoll(m, v82, v87, v20+int32(104), v20+int32(100), v20+int32(96))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v101 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	v108 = v68 + int32(4)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if base.Ui32(v108) < base.Ui32(v110+v111<<(uint(int32(2))%32)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v116 = v108
	goto L40
L39:
	;
	v116 = v101
	goto L40
L40:
	;
	v117 = F_transformAssignmentIndirection(m, l0, v101, v102, v101, v104, v105, v106, l7, v116, l9, l10, l11)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v120 = F_palloc0(m, int32(20))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(26)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v117
	v130 = F_list_make1_impl(m, int32(1), v20+int32(84))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v87
	v138 = F_list_make1_impl(m, int32(471), v20+int32(80))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v120)+12)) = v138
	if l4 == v80 {
		v346 = v120
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v144 = int32(0)
	v147 = F_coerce_to_domain(m, v120, v80, v143, l4, v144, int32(2), l11, v144)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	v346 = v147
	goto L1
L47:
	;
	v152 = v56 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v152 < v153 {
		v56 = v152
		v61 = v149
		goto L21
	} else {
		goto L48
	}
L48:
	;
	goto L22
L49:
	;
	v175 = F_transformAssignmentSubscripts(m, l0, v44, l2, l4, l5, l6, v168, l7, int32(0), l9, l10, l11)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v346 = v175
	goto L1
L51:
	;
	v198 = F_coerce_to_target_type(m, l0, l9, v194, l4, l5, l10, int32(2), int32(-1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	if v198 != 0 {
		v346 = v198
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v207 = F_format_type_be(m, l4)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v209 = F_exprType(m, l9)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v211 = F_format_type_be(m, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	if l3 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l2
	F_errmsg(m, int32(192625), v20+int32(16))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(619207), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(494506), int32(896), int32(255356))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(365908), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(494506), int32(736), int32(255356))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L14
	} else {
		goto L70
	}
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v258 = F_format_type_be(m, l4)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v257
	F_errmsg(m, int32(370611), v20+int32(32))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L14
	} else {
		goto L72
	}
L72:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L14
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(494506), int32(785), int32(255356))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v283 = F_format_type_be(m, l4)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v282
	F_errmsg(m, int32(193939), v20+int32(48))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L14
	} else {
		goto L78
	}
L78:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(494506), int32(794), int32(255356))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L14
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v307
	F_errmsg(m, int32(712659), v20-int32(-64))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L14
	} else {
		goto L83
	}
L83:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(494506), int32(800), int32(255356))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L14
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
	F_errhint(m, int32(619207), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L14
	} else {
		goto L87
	}
L87:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L14
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(494506), int32(886), int32(255356))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L14
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformContainerSubscripts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v7 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l3
	if l5 != 0 {
		v27 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = F_getSubscriptingRoutines(m, v27, v14+int32(24))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v19 = F_getBaseTypeAndTypmod(m, l2, v14+int32(28))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v27 = int32(1005)
	goto L1
L4:
	;
	v27 = int32(1028)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	switch v19 - int32(22) {
	case 0:
		goto L3
	default:
		v27 = v19
		goto L1
	case 8:
		goto L4
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L31
	}
L8:
	;
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l4 == int32(0) {
		v69 = v7
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L25
	}
L12:
	;
	v74 = F_palloc0(m, int32(40))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L22
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v34 <= int32(0) {
		v69 = v7
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(0)
	if v37 < v34 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = v34
	goto L17
L16:
	;
	v40 = v37
	goto L17
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v46 = int32(0)
	goto L18
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41+v46<<(uint(int32(2))%32))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
	if v58 != 0 {
		v69 = v58
		goto L12
	} else {
		goto L20
	}
L19:
	;
	v69 = v58
	goto L12
L20:
	;
	v60 = v46 + int32(1)
	if v60 != v40 {
		v46 = v60
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(14)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v81
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	m.T0[v86].(func(*base.Module, int32, int32, int32, int32, int32))(m, v74, l4, l0, v69, l5)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v89 == int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v14 + int32(32)
	return v74
L25:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v103 = F_format_type_be(m, v27)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v103
	F_errmsg(m, int32(329699), v14)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v109 = F_exprLocation(m, l1)
	mBase = m.M
	F_parser_errposition(m, l0, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(500671), int32(274), int32(118743))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v124 = F_format_type_be(m, v27)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v124
	F_errmsg(m, int32(329699), v14+int32(16))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(500671), int32(323), int32(118743))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformSortClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	v5 = int32(0)
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v48
L2:
	;
	v16 = v5
	v18 = v5
	goto L7
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v9 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v48 = v5
	goto L1
L6:
	;
	goto L5
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v16<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if l3 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = v36
	goto L1
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v36 = F_addTargetToSortList(m, l0, v34, v18, v35, v24)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L16
	}
L10:
	;
	v27 = F_findTargetlistEntrySQL99(m, l0, v25, l2, int32(20))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v32 = F_findTargetlistEntrySQL92(m, l0, v25, l2, int32(20))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v34 = v27
	goto L9
L15:
	;
	v34 = v32
	goto L9
L16:
	;
	v39 = v16 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v39 < v40 {
		v16 = v39
		v18 = v36
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
}
func F_transtime(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	v4 = int32(0)
	if l0&int32(3) != 0 {
		v22 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v23 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	default:
		v237 = v4
		goto L4
	}
L2:
	;
	v17 = base.I32_rem_s(l0, int32(100))
	if v17 != 0 {
		v22 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = base.I32_rem_s(l0, int32(400))
	v22 = base.B2i32(v19 == int32(0))
	goto L1
L4:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	return v245 + (l2 + v237)
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = l0 - base.B2i32(v37 < int32(3))
	v42 = base.I32_div_s(v40, int32(400))
	v44 = base.I32_rem_s(v40, int32(100))
	v47 = base.I32_div_s(v40, int32(-100))
	v48 = int32(1)
	v53 = base.I32_div_s(base.I32_extend8_s(v44), int32(4))
	v59 = base.I32_rem_s(v37+int32(9), int32(12))
	v66 = base.I32_div_s(base.I32_extend16_s(v59*int32(26)+int32(24)), int32(10))
	v71 = int32(7)
	v72 = base.I32_rem_s(v42+v44+v47<<(uint(v48)%32)+base.I32_extend8_s(v53)+base.I32_extend16_s(v66+v48), v71)
	if v72 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v237 = v33 * int32(86400)
	goto L4
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = int32(86400)
	v26 = v24 * v25
	v28 = v26 - v25
	if int32(59) < v24 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v26
	goto L10
L9:
	;
	v31 = v28
	goto L10
L10:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = v31
	goto L13
L12:
	;
	v32 = v28
	goto L13
L13:
	;
	v237 = v32
	goto L4
L14:
	;
	v77 = v72 + v71
	goto L16
L15:
	;
	v77 = v72
	goto L16
L16:
	;
	v78 = v36 - v77
	if v78 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v83 = v78 + int32(7)
	goto L19
L18:
	;
	v83 = v78
	goto L19
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v84 <= int32(1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v134 = v122 * int32(86400)
	if v128 <= int32(0) {
		v237 = v134
		goto L4
	} else {
		goto L28
	}
L21:
	;
	v122 = v83
	v128 = v37 - int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v91 = int32(1)
	v92 = v37 - v91
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v22*int32(48)+v92<<(uint(int32(2))%32))+uint32(_consts[1304])))
	v99 = int32(7)
	v105 = v83
	v108 = v91
	goto L24
L24:
	;
	v117 = v105 + int32(7)
	if v98 <= v117 {
		v122 = v105
		v128 = v92
		goto L20
	} else {
		goto L26
	}
L25:
	;
	v122 = v83 + v84*v99 - v99
	v128 = v92
	goto L20
L26:
	;
	v120 = v108 + int32(1)
	if v120 != v84 {
		v105 = v117
		v108 = v120
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v137 = int32(3)
	v138 = v128 & v137
	if base.Ui32(v37-int32(2)) < base.Ui32(v137) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v138 == int32(0) {
		v237 = v197
		goto L4
	} else {
		goto L36
	}
L30:
	;
	v194 = int32(0)
	v197 = v134
	goto L29
L31:
	;
	goto L32
L32:
	;
	v147 = int32(0)
	v151 = v147
	v154 = v134
	v160 = v147
	goto L33
L33:
	;
	v164 = v22*int32(48) + v151<<(uint(int32(2))%32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[1305])))
	v168 = int32(86400)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[1304])))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[1306])))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[1307])))
	v188 = v167*v168 + (v172*v168 + v154 + v178*v168 + v184*v168)
	v189 = int32(4)
	v190 = v151 + v189
	v192 = v160 + v189
	if v192 != v128&int32(2147483644) {
		v151 = v190
		v154 = v188
		v160 = v192
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v194 = v190
	v197 = v188
	goto L29
L35:
	;
	goto L34
L36:
	;
	v209 = v194
	v212 = v197
	v216 = int32(0)
	goto L37
L37:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v22*int32(48)+v209<<(uint(int32(2))%32))+uint32(_consts[1304])))
	v228 = v225*int32(86400) + v212
	v229 = int32(1)
	v232 = v216 + v229
	if v232 != v138 {
		v209 = v209 + v229
		v212 = v228
		v216 = v232
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v237 = v228
	goto L4
L39:
	;
	goto L38
}
func F_trgm_presence_map(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v102 int32
	_ = v102
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(2)
	v15 = int32(5)
	v16 = int32(base.Ui32(v12)>>(uint(v13)%32)) - v15
	v17 = int32(3)
	v18 = base.I32_div_u_s(v16, v17)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(base.Ui32(v19)>>(uint(v13)%32)) - v15
	v25 = base.I32_div_u_s(v23, v17)
	v26 = F_palloc0(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v23) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v34 = int32(1)
	if base.Ui32(v25) <= base.Ui32(v34) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	return v26
L6:
	;
	v37 = v34
	goto L8
L7:
	;
	v37 = v25
	goto L8
L8:
	;
	v45 = int32(0)
	v48 = l0 + int32(5)
	goto L9
L9:
	;
	if base.Ui32(v16) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v102 = v45 + int32(1)
	if v102 != v37 {
		v45 = v102
		v48 = v48 + int32(3)
		goto L9
	} else {
		goto L23
	}
L12:
	;
	v55 = v18
	v56 = int32(0)
	goto L13
L13:
	;
	v68 = int32(base.Ui32(v55+v56) >> (uint(int32(1)) % 32))
	v73 = *(*int32)(unsafe.Add(mBase, _consts[1457]))
	v74 = m.T0[v73].(func(*base.Module, int32, int32) int32)(m, v48, l1+int32(5)+v68*int32(3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v26))) = uint8(v86)
	goto L11
L15:
	;
	goto L14
L16:
	;
	if v83 < v82 {
		v55 = v82
		v56 = v83
		goto L13
	} else {
		goto L22
	}
L17:
	;
	if v74 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = v68
	v83 = v56
	goto L16
L19:
	;
	goto L20
L20:
	;
	if v74 == int32(0) {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v82 = v55
	v83 = v68 + int32(1)
	goto L16
L22:
	;
	goto L11
L23:
	;
	goto L10
}
func F_tsearch_readline(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v4 + int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 != v8 {
			F_pfree(m, v8)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = l0 + int32(12)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v21 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v21
				v27 = F_pg_get_line_append(m, v17, v19)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v32 = F_pg_any_to_server(m, v29, v30, int32(6))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
							v35 = F_pstrdup(m, v32)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v39 = v35
								return v39
							}
						}
					} else {
						v39 = int32(0)
						return v39
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = l0 + int32(12)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v21 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v21
			v27 = F_pg_get_line_append(m, v17, v19)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v32 = F_pg_any_to_server(m, v29, v30, int32(6))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
						v35 = F_pstrdup(m, v32)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v39 = v35
							return v39
						}
					}
				} else {
					v39 = int32(0)
					return v39
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = l0 + int32(12)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v21
		*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v21
		v27 = F_pg_get_line_append(m, v17, v19)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v27 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v32 = F_pg_any_to_server(m, v29, v30, int32(6))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v32
					v35 = F_pstrdup(m, v32)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v39 = v35
						return v39
					}
				}
			} else {
				v39 = int32(0)
				return v39
			}
		}
	}
}
func F_tsearch_readline_begin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = F_AllocateFile(m, l1, int32(231353))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
		if v5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
			F_initStringInfo(m, l0+int32(12))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1179)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0
				v22 = int32(4513176)
				v23 = *(*int32)(unsafe.Add(mBase, _consts[141]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v23
				*(*int32)(unsafe.Add(mBase, _consts[141])) = l0 + int32(32)
				return base.B2i32(v5 != int32(0))
			}
		} else {
			return base.B2i32(v5 != int32(0))
		}
	}
}
func F_tsm_handler_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(219859)
			F_errmsg(m, int32(192913), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495301), int32(372), int32(279951))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_tsm_system_handler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_palloc0(m, int32(36))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(440)
		v15 = int32(700)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v15
		v22 = F_list_make1_impl(m, int32(472), v6+int32(8))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(282)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(283)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(284)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(285)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(286)
			v36 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v36)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22
			m.G0 = v6 + int32(16)
			return v9
		}
	}
}
func F_tsq_mcontained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(1530), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_tstoreReceiveSlot_notoast(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_tuplestore_puttupleslot(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_tstoreStartupReceiver(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v9 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v38 = v4
	goto L1
L3:
	;
	goto L4
L4:
	;
	if v8 <= int32(0) {
		v38 = v4
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v18 = int32(0)
	goto L6
L6:
	;
	v26 = l2 + int32(20) + v18<<(uint(int32(4))%32)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+9)))
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v38 = v4
	goto L1
L8:
	;
	v33 = v18 + int32(1)
	if v33 != v8 {
		v18 = v33
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v28 != int32(65535) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v38 = int32(1)
	goto L1
L11:
	;
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	return
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(787)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(789)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v67 = v8 << (uint(int32(2)) % 32)
	v68 = F_MemoryContextAlloc(m, v65, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L24
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v44 = F_convert_tuples_by_position(m, l2, v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v58 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v58
	if v38 == v58 {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	return
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v44
	if v38 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if v44 == int32(0) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(788)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v55 = F_MakeSingleTupleTableSlot(m, v53, int32(1619316))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v55
	return
L23:
	;
	goto L14
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v72 = F_MemoryContextAlloc(m, v71, v67)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v72
	goto L12
}
func F_tuplehash_iterate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v9 = v6
	goto L1
L1:
	;
	if v9&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v25
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = v17 & (v18 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
	v25 = v16 + v18*int32(12)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = v26 & (v27 ^ v21)
	if v29 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v32)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v36 != int32(1) {
		v9 = base.B2i32(v29 == int32(0))
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
}
func F_typeidType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(50636), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499999), int32(584), int32(372624))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v6 + int32(16)
			return v9
		}
	}
}
func F_typeidTypeRelid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(50636), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499999), int32(676), int32(436380))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+84))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_typenameTypeIdAndMod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = F_typenameType(m, l0, l1, l3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7+v8)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v10
		F_ReleaseCatCache(m, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	}
}
