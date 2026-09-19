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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
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
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int64
	_ = v310
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
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
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v18 + int32(48)
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v25) < base.Ui32(int32(_a_F_TransferPredicateLocksToHeapRelation_0)) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+118)))
	if v29 == int32(116) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = v33
	goto L7
L6:
	;
	v34 = v25
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v41 = F_LWLockAcquire(m, v37+int32(3840), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v48 = F_LWLockAcquire(m, v44+int32(_a_F_TransferPredicateLocksToHeapRelation_1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v55 = F_LWLockAcquire(m, v51+int32(_a_F_TransferPredicateLocksToHeapRelation_2), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v62 = F_LWLockAcquire(m, v58+int32(_a_F_TransferPredicateLocksToHeapRelation_3), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v69 = F_LWLockAcquire(m, v65+int32(_a_F_TransferPredicateLocksToHeapRelation_4), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v76 = F_LWLockAcquire(m, v72+int32(_a_F_TransferPredicateLocksToHeapRelation_5), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v83 = F_LWLockAcquire(m, v79+int32(_a_F_TransferPredicateLocksToHeapRelation_6), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v90 = F_LWLockAcquire(m, v86+int32(_a_F_TransferPredicateLocksToHeapRelation_7), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v97 = F_LWLockAcquire(m, v93+int32(_a_F_TransferPredicateLocksToHeapRelation_8), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v104 = F_LWLockAcquire(m, v100+int32(_a_F_TransferPredicateLocksToHeapRelation_9), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v111 = F_LWLockAcquire(m, v107+int32(_a_F_TransferPredicateLocksToHeapRelation_10), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v118 = F_LWLockAcquire(m, v114+int32(_a_F_TransferPredicateLocksToHeapRelation_11), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v125 = F_LWLockAcquire(m, v121+int32(_a_F_TransferPredicateLocksToHeapRelation_12), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v132 = F_LWLockAcquire(m, v128+int32(_a_F_TransferPredicateLocksToHeapRelation_13), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v139 = F_LWLockAcquire(m, v135+int32(_a_F_TransferPredicateLocksToHeapRelation_14), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v146 = F_LWLockAcquire(m, v142+int32(_a_F_TransferPredicateLocksToHeapRelation_15), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v153 = F_LWLockAcquire(m, v149+int32(_a_F_TransferPredicateLocksToHeapRelation_16), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	v160 = F_LWLockAcquire(m, v156+int32(3584), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[2]))
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[3]))
	v169 = v18 + int32(28)
	v170 = F_hash_search_with_hash_value(m, v163, int32(_a_F_TransferPredicateLocksToHeapRelation_17), v166, int32(2), v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[2]))
	F_hash_seq_init(m, v169, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v176 = F_hash_seq_search(m, v169)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	if v176 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v181 = v2
	v183 = v176
	v190 = v2
	goto L33
L31:
	;
	goto L32
L32:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[2]))
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[3]))
	v382 = F_hash_search_with_hash_value(m, v375, int32(_a_F_TransferPredicateLocksToHeapRelation_17), v378, int32(1), v18+int32(8))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L8
	} else {
		goto L70
	}
L33:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v193 != v25 {
		v343 = v181
		v352 = v190
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v357 = F_hash_seq_search(m, v18+int32(28))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L68
	}
L36:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v195 != v35 {
		v343 = v181
		v352 = v190
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if v32 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v181 != 0 {
		v224 = v181
		v225 = v190
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	if v197 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	if v198 == int32(-1) {
		v343 = v181
		v352 = v190
		goto L35
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v183)+20))
	if v226 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v35
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[2]))
	v208 = v18 + int32(8)
	v209 = F_get_hash_value(m, v206, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[2]))
	v216 = F_hash_search_with_hash_value(m, v212, v208, v209, int32(1), v18+int32(27))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+27)))
	if v218 != 0 {
		v224 = v216
		v225 = v209
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v220 = v216 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v220
	v224 = v216
	v225 = v209
	goto L42
L47:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[2]))
	v338 = F_hash_search(m, v334, v183, int32(2), v18+int32(27))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L8
	} else {
		goto L67
	}
L48:
	;
	v230 = v183 + int32(16)
	if v226 == v230 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v233 = v224 + int32(16)
	v241 = v226
	goto L50
L50:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v241-int32(4))))
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v241)+16))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v257
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[4]))
	v265 = v18 + int32(27)
	v266 = F_hash_search(m, v260, v241-int32(8), int32(2), v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L8
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v224
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[4]))
	v278 = F_hash_search_with_hash_value(m, v271, v18+int32(8), v251<<(uint(int32(4))%32)^v225, int32(1), v265)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+27)))
	if v280 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	if v253 != v230 {
		v241 = v253
		goto L50
	} else {
		goto L66
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v278)+24)) = v252
	goto L54
L56:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	if v283 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v278)+24))
	if base.Ui64(v252) <= base.Ui64(v310) {
		goto L54
	} else {
		goto L65
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v224)+16)) = v233
	goto L61
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+12)) = v233
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+8)) = v289
	v292 = v278 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v292
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v297 = v295 + int32(48)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)+52))
	if v298 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+52)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v295)+48)) = v297
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+20)) = v297
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v304
	v307 = v278 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+4)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v307
	goto L55
L65:
	;
	goto L55
L66:
	;
	goto L51
L67:
	;
	v343 = v224
	v352 = v225
	goto L35
L68:
	;
	if v357 != 0 {
		v181 = v343
		v183 = v357
		v190 = v352
		goto L33
	} else {
		goto L69
	}
L69:
	;
	goto L34
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v385+int32(3584))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v391+int32(_a_F_TransferPredicateLocksToHeapRelation_16))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v397+int32(_a_F_TransferPredicateLocksToHeapRelation_15))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v403+int32(_a_F_TransferPredicateLocksToHeapRelation_14))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v409+int32(_a_F_TransferPredicateLocksToHeapRelation_13))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v415+int32(_a_F_TransferPredicateLocksToHeapRelation_12))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v421+int32(_a_F_TransferPredicateLocksToHeapRelation_11))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v427+int32(_a_F_TransferPredicateLocksToHeapRelation_10))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v433+int32(_a_F_TransferPredicateLocksToHeapRelation_9))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v439+int32(_a_F_TransferPredicateLocksToHeapRelation_8))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v445+int32(_a_F_TransferPredicateLocksToHeapRelation_7))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v451+int32(_a_F_TransferPredicateLocksToHeapRelation_6))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v457+int32(_a_F_TransferPredicateLocksToHeapRelation_5))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v463+int32(_a_F_TransferPredicateLocksToHeapRelation_4))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v469+int32(_a_F_TransferPredicateLocksToHeapRelation_3))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v475+int32(_a_F_TransferPredicateLocksToHeapRelation_2))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v481+int32(_a_F_TransferPredicateLocksToHeapRelation_1))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToHeapRelation[1]))
	F_LWLockRelease(m, v487+int32(3840))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v104 int64
	_ = v104
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int32
	_ = v122
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v152 int64
	_ = v152
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = l1 & int64(281474976710655)
	v19 = int64(base.Ui64(l1)>>(uint(int64(48))%64)) & int64(32767)
	v20 = base.I32_wrap_i64(v19)
	if base.Ui32(v20-int32(_a_F___trunctfsf2_0)) <= base.Ui32(int32(253)) {
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
		v52 = base.B2i32(base.Ui32(int32(_a_F___trunctfsf2_1)) < base.Ui32(v49))
		if base.Ui32(int32(_a_F___trunctfsf2_1)) < base.Ui32(v49) {
			v53 = int32(0)
		} else {
			v53 = v49
		}
		if base.Ui32(int32(_a_F___trunctfsf2_1)) < base.Ui32(v49) {
			v56 = int32(-16255)
		} else {
			v56 = int32(-16256)
		}
		v181 = v53
		v182 = v56 + v20
	} else {
		if base.B2i32(l0|v15 == int64(0))|base.B2i32(v19 != int64(32767)) == int32(0) {
			v181 = base.I32_wrap_i64(int64(base.Ui64(v15)>>(uint(int64(25))%64))) | int32(_a_F___trunctfsf2_2)
			v182 = int32(255)
		} else {
			if base.Ui32(int32(_a_F___trunctfsf2_3)) < base.Ui32(v20) {
				v181 = int32(0)
				v182 = int32(255)
			} else {
				v78 = base.B2i32(v19 == int64(0))
				if v19 == int64(0) {
					v79 = int32(_a_F___trunctfsf2_4)
				} else {
					v79 = int32(_a_F___trunctfsf2_0)
				}
				v80 = v79 - v20
				if int32(112) < v80 {
					v83 = int32(0)
					v181 = v83
					v182 = v83
				} else {
					if v19 == int64(0) {
						v87 = v15
					} else {
						v87 = v15 | int64(281474976710656)
					}
					if v20 != v79 {
						v91 = v12 + int32(16)
						v93 = int32(128) - v80
						if v93&int32(64) != 0 {
							v112 = int64(0)
							v113 = l0 << (uint(base.I64_extend_i32_u(v93+int32(-64))) % 64)
						} else {
							if v93 == int32(0) {
								v112 = l0
								v113 = v87
							} else {
								v104 = base.I64_extend_i32_u(v93)
								v112 = l0 << (uint(v104) % 64)
								v113 = v87<<(uint(v104)%64) | int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v93))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v91))) = v112
						*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v113
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
						v122 = base.B2i32(v117|v118 != int64(0))
					} else {
						v122 = int32(0)
					}
					if v80&int32(64) != 0 {
						v141 = int64(base.Ui64(v87) >> (uint(base.I64_extend_i32_u(v80+int32(-64))) % 64))
						v142 = int64(0)
					} else {
						if v80 == int32(0) {
							v141 = l0
							v142 = v87
						} else {
							v137 = base.I64_extend_i32_u(v80)
							v141 = v87<<(uint(base.I64_extend_i32_u(int32(64)-v80))%64) | int64(base.Ui64(l0)>>(uint(v137)%64))
							v142 = int64(base.Ui64(v87) >> (uint(v137) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v141
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v142
					v146 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					v149 = base.I32_wrap_i64(int64(base.Ui64(v146) >> (uint(int64(25)) % 64)))
					v150 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					v152 = v150 | base.I64_extend_i32_u(v122)
					v156 = v146 & int64(33554431)
					v157 = int64(16777216)
					if v156 == v157 {
						v161 = base.B2i32(v152 == int64(0))
					} else {
						v161 = base.B2i32(base.Ui64(v156) < base.Ui64(v157))
					}
					if v161 == int32(0) {
						v174 = v149 + int32(1)
					} else {
						if v152|(v156^int64(16777216)) != int64(0) {
							v174 = v149
						} else {
							v174 = v149&int32(1) + v149
						}
					}
					v178 = base.B2i32(base.Ui32(int32(_a_F___trunctfsf2_1)) < base.Ui32(v174))
					if base.Ui32(int32(_a_F___trunctfsf2_1)) < base.Ui32(v174) {
						v179 = v174 ^ int32(_a_F___trunctfsf2_5)
					} else {
						v179 = v174
					}
					v181 = v179
					v182 = v178
				}
			}
		}
	}
	m.G0 = v12 + int32(32)
	return base.F32_reinterpret_i32(base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(32))%64)))&int32(-2147483648) | v182<<(uint(int32(23))%32) | v181)
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		if v10 == int32(1) {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
			if v16 == int32(18) {
				v19 = int32(16)
			} else {
				v19 = int32(0)
			}
			if base.Ui32((v16-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v26 = int32(4)
			} else {
				v26 = v19
			}
			v39 = v26
		} else {
			v27 = int32(1)
			if v10&v27 != 0 {
				v39 = int32(base.Ui32(v10)>>(uint(v27)%32)) - v27
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v42 = F_palloc(m, v39+int32(1))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			if v39 != 0 {
				v44 = int32(1)
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
				if v46&v44 != 0 {
					v49 = v44
				} else {
					v49 = int32(4)
				}
				base.MemoryCopy(m, v42, v6+v49, v39)
			} else {
			}
			v53 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v39+v42))) = uint8(v53)
			if v6 != v5 {
				F_pfree(m, v6)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					return v42
				}
			} else {
				return v42
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
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
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
		v194 = int32(-1)
		m.G0 = v13 + int32(16)
		return v194
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp2tm[0]))
		v34 = base.I32_wrap_i64(v29)
		v46 = v34 + int32(_a_F_timestamp2tm_0)
		v47 = int32(_a_F_timestamp2tm_1)
		v48 = base.I32_div_u_s(v46, v47)
		v49 = int32(3)
		v55 = int32(2)
		v60 = base.I32_div_u_s((v48*int32(1073595727)+v46)<<(uint(v55)%32)|v49, v47)
		v63 = v34 + int32(_a_F_timestamp2tm_2) + v48*v49 + v60 + int32(_a_F_timestamp2tm_3)
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
		*(*int32)(unsafe.Add(mBase, uint32(l2+int32(20)))) = v83 + v65<<(uint(int32(2))%32) - int32(_a_F_timestamp2tm_4)
		v91 = v81 + int32(123)
		v95 = int32(base.Ui32(v91*int32(2141)) >> (uint(int32(16)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(l2+int32(12)))) = v91 - int32(base.Ui32(v95*int32(_a_F_timestamp2tm_5))>>(uint(int32(8))%32))
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
				v187 = v137
				v188 = v137
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v188
				v194 = v187
			} else {
				v194 = v137
			}
			m.G0 = v13 + int32(16)
			return v194
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
					v194 = v178
				} else {
					v184 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
					v187 = v178
					v188 = v184
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v188
					v194 = v187
				}
				m.G0 = v13 + int32(16)
				return v194
			}
		}
	}
}
func F_timetypmodin(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14003(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_tlist_same_datatypes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9 = v7
	goto L3
L2:
	;
	v9 = int32(0)
	goto L3
L3:
	;
	if l0 == int32(0) {
		v58 = v9
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return base.B2i32(v58 == int32(0))
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 <= int32(0) {
		v58 = v9
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v18 = v9
	v19 = int32(0)
	goto L7
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v19<<(uint(int32(2))%32))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+26)))
	if v26 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v58 = v50
	goto L4
L9:
	;
	v52 = v19 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v52 < v53 {
		v18 = v50
		v19 = v52
		goto L7
	} else {
		goto L22
	}
L10:
	;
	return int32(0)
L11:
	;
	if v18 == int32(0) {
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
		v50 = v18
		goto L9
	} else {
		goto L21
	}
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v32 = F_exprType(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v32 != v36 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v39 = v18 + int32(4)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v39) < base.Ui32(v41+v42<<(uint(int32(2))%32)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v47 = v39
	goto L20
L19:
	;
	v47 = int32(0)
	goto L20
L20:
	;
	v50 = v47
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
					v38 = int32(_a_F_tm2timestamp_0)
				} else {
					v38 = int32(_a_F_tm2timestamp_1)
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
				v59 = base.I32_div_s((v54+v31)*int32(_a_F_tm2timestamp_2), int32(256))
				v65 = base.I64_extend_i32_s(v32 + v39*int32(365) + v44 + v47 + v50 + v59 - int32(_a_F_tm2timestamp_3) - int32(_a_F_tm2timestamp_4))
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
		if v15 <= int32(_a_F_tm2timestamp_5) {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v31 = v25
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v37 = base.B2i32(int32(2) < v31)
			if int32(2) < v31 {
				v38 = int32(_a_F_tm2timestamp_0)
			} else {
				v38 = int32(_a_F_tm2timestamp_1)
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
			v59 = base.I32_div_s((v54+v31)*int32(_a_F_tm2timestamp_2), int32(256))
			v65 = base.I64_extend_i32_s(v32 + v39*int32(365) + v44 + v47 + v50 + v59 - int32(_a_F_tm2timestamp_3) - int32(_a_F_tm2timestamp_4))
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
			if v15 != int32(_a_F_tm2timestamp_6) {
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
						v38 = int32(_a_F_tm2timestamp_0)
					} else {
						v38 = int32(_a_F_tm2timestamp_1)
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
					v59 = base.I32_div_s((v54+v31)*int32(_a_F_tm2timestamp_2), int32(256))
					v65 = base.I64_extend_i32_s(v32 + v39*int32(365) + v44 + v47 + v50 + v59 - int32(_a_F_tm2timestamp_3) - int32(_a_F_tm2timestamp_4))
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
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_include_file[0]))
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
								F_errmsg(m, int32(_a_F_tokenize_include_file_0), v11)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_tokenize_include_file_1), int32(460), int32(_a_F_tokenize_include_file_2))
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
							v44 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_include_file[1]))
							F_MemoryContextDelete(m, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								v48 = int32(_a_F_tokenize_include_file_3)
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_trackitem_compare_element[0]))
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformAExprOp[0])))
	if v10 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v110 = F_transformExprRecurse(m, l0, v8)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L41
	}
L2:
	;
	if base.B2i32(v7 == int32(0))|base.B2i32(v73 != int32(36)) != 0 {
		goto L1
	} else {
		goto L28
	}
L3:
	;
	if v8 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v13 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 != int32(1) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 != int32(61) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v25 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v8 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v46 == int32(34) {
		v73 = v45
		goto L2
	} else {
		goto L18
	}
L10:
	;
	if v7 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L11:
	;
	v28 = int32(72)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v29 != v28 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v32 != 0 {
		v45 = v28
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v36 != int32(72) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v39 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v42 == int32(34) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v45 = v42
	goto L9
L18:
	;
	v50 = F_palloc0(m, int32(20))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(52)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v60 == int32(72) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v64
	v66 = F_transformExprRecurse(m, l0, v50)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L26
	}
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)))
	if v63 != 0 {
		v64 = v7
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v64 = v8
	goto L21
L25:
	;
	goto L24
L26:
	;
	return v66
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v73 = v72
	goto L2
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v81 = v79 - int32(22)
	if v81 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v97 = F_transformExprRecurse(m, l0, v8)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L38
	}
L30:
	;
	if v81 == int32(14) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v84 != int32(4) {
		goto L1
	} else {
		goto L36
	}
L33:
	;
	goto L29
L34:
	;
	goto L1
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(3)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v92
	v94 = F_transformExprRecurse(m, l0, v7)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	return v94
L38:
	;
	v99 = F_transformExprRecurse(m, l0, v7)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v105 = F_make_row_comparison_op(m, l0, v101, v102, v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	return v105
L41:
	;
	v112 = F_transformExprRecurse(m, l0, v7)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v116 = F_make_op(m, l0, v114, v110, v112, v109, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	return v116
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	v13 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	if l1|base.B2i32(l8 == v13) == v13 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v20 + int32(112)
	return v349
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = l2
	F_errmsg(m, int32(_a_F_transformAssignmentIndirection_0), v20)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L14
	} else {
		goto L86
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L14
	} else {
		goto L81
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L14
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L14
	} else {
		goto L69
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L14
	} else {
		goto L64
	}
L7:
	;
	v197 = F_exprType(m, l9)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L14
	} else {
		goto L51
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v48 < v49 {
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v47 = v39
	v48 = (l8 - v40) >> (uint(int32(2)) % 32)
	goto L8
L11:
	;
	v28 = F_palloc0(m, int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if l8 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(34)
	v39 = v28
	goto L10
L16:
	;
	v39 = l1
	goto L10
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	v47 = l1
	v48 = v46
	goto L8
L18:
	;
	v59 = v48
	v65 = v13
	goto L21
L19:
	;
	v172 = v13
	goto L20
L20:
	;
	if v172 == int32(0) {
		goto L7
	} else {
		goto L49
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v71 = v68 + v59<<(uint(int32(2))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 != int32(78) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v172 = v152
	goto L20
L23:
	;
	if v73 == int32(77) {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v152 = F_lappend(m, v65, v72)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L14
	} else {
		goto L47
	}
L26:
	;
	if v65 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = F_transformAssignmentSubscripts(m, l0, v47, l2, l4, l5, l6, v65, l7, v71, l9, l10, l11)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
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
	v83 = F_getBaseTypeAndTypmod(m, l4, v20+int32(108))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L14
	} else {
		goto L31
	}
L30:
	;
	v349 = v78
	goto L1
L31:
	;
	v85 = F_typeidTypeRelid(m, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	if v85 == int32(0) {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v90 = F_get_attnum(m, v85, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	if v90 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v90 < int32(0) {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	F_get_atttypetypmodcoll(m, v85, v90, v20+int32(104), v20+int32(100), v20+int32(96))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	v104 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+96))
	v111 = v71 + int32(4)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if base.Ui32(v111) < base.Ui32(v113+v114<<(uint(int32(2))%32)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v119 = v111
	goto L40
L39:
	;
	v119 = v104
	goto L40
L40:
	;
	v120 = F_transformAssignmentIndirection(m, l0, v104, v105, v104, v107, v108, v109, l7, v119, l9, l10, l11)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v123 = F_palloc0(m, int32(20))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = int32(26)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v120
	v133 = F_list_make1_impl(m, int32(1), v20+int32(84))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v90
	v141 = F_list_make1_impl(m, int32(471), v20+int32(80))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v141
	if l4 == v83 {
		v349 = v123
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	v147 = int32(0)
	v150 = F_coerce_to_domain(m, v123, v83, v146, l4, v147, int32(2), l11, v147)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	v349 = v150
	goto L1
L47:
	;
	v155 = v59 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v155 < v156 {
		v59 = v155
		v65 = v152
		goto L21
	} else {
		goto L48
	}
L48:
	;
	goto L22
L49:
	;
	v178 = F_transformAssignmentSubscripts(m, l0, v47, l2, l4, l5, l6, v172, l7, int32(0), l9, l10, l11)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v349 = v178
	goto L1
L51:
	;
	v201 = F_coerce_to_target_type(m, l0, l9, v197, l4, l5, l10, int32(2), int32(-1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	if v201 != 0 {
		v349 = v201
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v210 = F_format_type_be(m, l4)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v212 = F_exprType(m, l9)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v214 = F_format_type_be(m, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l2
	F_errmsg(m, int32(_a_F_transformAssignmentIndirection_1), v20+int32(16))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(_a_F_transformAssignmentIndirection_2), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_transformAssignmentIndirection_3), int32(896), int32(_a_F_transformAssignmentIndirection_4))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
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
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_transformAssignmentIndirection_5), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_transformAssignmentIndirection_3), int32(736), int32(_a_F_transformAssignmentIndirection_4))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
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
	v259 = m.ExcPending
	if v259 != 0 {
		goto L14
	} else {
		goto L70
	}
L70:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v261 = F_format_type_be(m, l4)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v260
	F_errmsg(m, int32(_a_F_transformAssignmentIndirection_6), v20+int32(32))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L14
	} else {
		goto L72
	}
L72:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L14
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_transformAssignmentIndirection_3), int32(785), int32(_a_F_transformAssignmentIndirection_4))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v286 = F_format_type_be(m, l4)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v285
	F_errmsg(m, int32(_a_F_transformAssignmentIndirection_7), v20+int32(48))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L14
	} else {
		goto L78
	}
L78:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_transformAssignmentIndirection_3), int32(794), int32(_a_F_transformAssignmentIndirection_4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
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
	v309 = m.ExcPending
	if v309 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v310
	F_errmsg(m, int32(_a_F_transformAssignmentIndirection_8), v20-int32(-64))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L14
	} else {
		goto L83
	}
L83:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_transformAssignmentIndirection_3), int32(800), int32(_a_F_transformAssignmentIndirection_4))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
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
	F_errhint(m, int32(_a_F_transformAssignmentIndirection_2), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L14
	} else {
		goto L87
	}
L87:
	;
	F_parser_errposition(m, l0, l11)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L14
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_transformAssignmentIndirection_3), int32(886), int32(_a_F_transformAssignmentIndirection_4))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
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
	var v72 int32
	_ = v72
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
		v72 = v7
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
		v72 = v7
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
		v72 = v58
		goto L12
	} else {
		goto L20
	}
L19:
	;
	v72 = v58
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
	m.T0[v86].(func(*base.Module, int32, int32, int32, int32, int32))(m, v74, l4, l0, v72, l5)
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
	F_errmsg(m, int32(_a_F_transformContainerSubscripts_0), v14)
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
	F_errfinish(m, int32(_a_F_transformContainerSubscripts_1), int32(274), int32(_a_F_transformContainerSubscripts_2))
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
	F_errmsg(m, int32(_a_F_transformContainerSubscripts_0), v14+int32(16))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_transformContainerSubscripts_1), int32(323), int32(_a_F_transformContainerSubscripts_2))
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
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	v5 = int32(0)
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v5
	v22 = v5
	goto L7
L5:
	;
	v52 = v5
	goto L6
L6:
	;
	return v52
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if l3 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v52 = v40
	goto L6
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v40 = F_addTargetToSortList(m, l0, v38, v22, v39, v28)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L16
	}
L10:
	;
	v31 = F_findTargetlistEntrySQL99(m, l0, v29, l2, int32(20))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v36 = F_findTargetlistEntrySQL92(m, l0, v29, l2, int32(20))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	v38 = v31
	goto L9
L15:
	;
	v38 = v36
	goto L9
L16:
	;
	v43 = v20 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 < v44 {
		v20 = v43
		v22 = v40
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	v4 = int32(0)
	if l0&int32(3) != 0 {
		v21 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v22 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		goto L5
	default:
		v220 = v4
		goto L4
	}
L2:
	;
	v16 = base.I32_rem_s(l0, int32(100))
	if v16 != 0 {
		v21 = int32(1)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = base.I32_rem_s(l0, int32(400))
	v21 = base.B2i32(v18 == int32(0))
	goto L1
L4:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	return v226 + (l2 + v220)
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v39 = l0 - base.B2i32(v36 < int32(3))
	v41 = base.I32_div_s(v39, int32(400))
	v43 = base.I32_rem_s(v39, int32(100))
	v46 = base.I32_div_s(v39, int32(-100))
	v47 = int32(1)
	v52 = base.I32_div_s(base.I32_extend8_s(v43), int32(4))
	v58 = base.I32_rem_s(v36+int32(9), int32(12))
	v65 = base.I32_div_s(base.I32_extend16_s(v58*int32(26)+int32(24)), int32(10))
	v70 = int32(7)
	v71 = base.I32_rem_s(v41+v43+v46<<(uint(v47)%32)+base.I32_extend8_s(v52)+base.I32_extend16_s(v65+v47), v70)
	if v71 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v220 = v32 * int32(_a_F_transtime_0)
	goto L4
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = int32(_a_F_transtime_0)
	v25 = v23 * v24
	v27 = v25 - v24
	if int32(59) < v23 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v30 = v25
	goto L10
L9:
	;
	v30 = v27
	goto L10
L10:
	;
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v30
	goto L13
L12:
	;
	v31 = v27
	goto L13
L13:
	;
	v220 = v31
	goto L4
L14:
	;
	v76 = v71 + v70
	goto L16
L15:
	;
	v76 = v71
	goto L16
L16:
	;
	v77 = v35 - v76
	if v77 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v82 = v77 + int32(7)
	goto L19
L18:
	;
	v82 = v77
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v83 < int32(2) {
		v116 = v82
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v127 = v116 * int32(_a_F_transtime_0)
	v129 = v36 - int32(1)
	if v129 <= int32(0) {
		v220 = v127
		goto L4
	} else {
		goto L26
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v21*int32(48)+v36<<(uint(int32(2))%32))+uint32(_c_F_transtime[0])))
	v94 = int32(7)
	v100 = v82
	v104 = int32(1)
	goto L22
L22:
	;
	v111 = v100 + int32(7)
	if v93 <= v111 {
		v116 = v100
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v116 = v82 + v83*v94 - v94
	goto L20
L24:
	;
	v114 = v104 + int32(1)
	if v114 != v83 {
		v100 = v111
		v104 = v114
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v132 = int32(3)
	v133 = v129 & v132
	v137 = v21*int32(48) + int32(_a_F_transtime_1)
	if base.Ui32(v36-int32(2)) < base.Ui32(v132) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v194 = v184
	v198 = v188
	v201 = int32(0)
	goto L35
L28:
	;
	v184 = int32(0)
	v188 = v127
	goto L27
L29:
	;
	goto L30
L30:
	;
	v146 = int32(0)
	v148 = v146
	v152 = v127
	v153 = v146
	goto L31
L31:
	;
	v160 = v137 + v148<<(uint(int32(2))%32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v162 = int32(_a_F_transtime_0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	v176 = v161*v162 + (v164*v162 + v152 + v168*v162 + v172*v162)
	v177 = int32(4)
	v178 = v148 + v177
	v180 = v153 + v177
	if v180 != v129&int32(2147483644) {
		v148 = v178
		v152 = v176
		v153 = v180
		goto L31
	} else {
		goto L33
	}
L32:
	;
	if v133 == int32(0) {
		v220 = v176
		goto L4
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v184 = v178
	v188 = v176
	goto L27
L35:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v137+v194<<(uint(int32(2))%32))))
	v210 = v207*int32(_a_F_transtime_0) + v198
	v211 = int32(1)
	v214 = v201 + v211
	if v214 != v133 {
		v194 = v194 + v211
		v198 = v210
		v201 = v214
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v220 = v210
	goto L4
L37:
	;
	goto L36
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
	var v46 int32
	_ = v46
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
	v46 = l0 + int32(5)
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
		v46 = v46 + int32(3)
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
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_trgm_presence_map[0]))
	v74 = m.T0[v73].(func(*base.Module, int32, int32) int32)(m, v46, l1+int32(5)+v68*int32(3))
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
	v5 = F_AllocateFile(m, l1, int32(_a_F_tsearch_readline_begin_0))
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
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(1163)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l0
				v22 = int32(_a_F_tsearch_readline_begin_1)
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_tsearch_readline_begin[0]))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v23
				*(*int32)(unsafe.Add(mBase, _c_F_tsearch_readline_begin[0])) = l0 + int32(32)
				return base.B2i32(v5 != int32(0))
			}
		} else {
			return base.B2i32(v5 != int32(0))
		}
	}
}
func F_tsm_handler_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_tsm_handler_in_0), int32(372), int32(_a_F_tsm_handler_in_1), int32(_a_F_tsm_handler_in_2), int32(_a_F_tsm_handler_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_tsm_system_handler(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14009(m, l0, int32(257), int32(286), int32(285), int32(284), int32(283), int32(282), int32(700))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1514), int32(0), v4, v5)
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v7 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(787)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	return
L2:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v65 = F_convert_tuples_by_position(m, l2, v40, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L22
	}
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v40 != 0 {
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v10 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = int32(0)
	goto L7
L7:
	;
	v24 = l2 + int32(20) + v17<<(uint(int32(4))%32)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
	if v28 == int32(_a_F_tstoreStartupReceiver_0) {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v32 = v17 + int32(1)
	if v32 != v10 {
		v17 = v32
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L8
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	goto L1
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v46 = F_convert_tuples_by_position(m, l2, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v48 = int32(0)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(788)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v48
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v54 = v10 << (uint(int32(2)) % 32)
	v55 = F_MemoryContextAlloc(m, v52, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L20
	}
L18:
	;
	return
L19:
	;
	v48 = v46
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v59 = F_MemoryContextAlloc(m, v58, v54)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	return
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v65
	if v65 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(789)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v76 = F_MakeTupleTableSlot(m, v74, int32(_a_F_tstoreStartupReceiver_1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v76
	return
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
func F_turkish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1067 int32
	_ = v1067
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1220 int32
	_ = v1220
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1284 int32
	_ = v1284
	var v1292 int32
	_ = v1292
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1451 int32
	_ = v1451
	var v1459 int32
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1774 int32
	_ = v1774
	var v1782 int32
	_ = v1782
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1982 int32
	_ = v1982
	var v1998 int32
	_ = v1998
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2122 int32
	_ = v2122
	var v2126 int32
	_ = v2126
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2148 int32
	_ = v2148
	var v2152 int32
	_ = v2152
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2169 int32
	_ = v2169
	var v2177 int32
	_ = v2177
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2243 int32
	_ = v2243
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2303 int32
	_ = v2303
	var v2307 int32
	_ = v2307
	var v2311 int32
	_ = v2311
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2336 int32
	_ = v2336
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2355 int32
	_ = v2355
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2406 int32
	_ = v2406
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2418 int32
	_ = v2418
	var v2427 int32
	_ = v2427
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2453 int32
	_ = v2453
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2495 int32
	_ = v2495
	var v2505 int32
	_ = v2505
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2514 int32
	_ = v2514
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2550 int32
	_ = v2550
	var v2553 int32
	_ = v2553
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2572 int32
	_ = v2572
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2610 int32
	_ = v2610
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L3
L1:
	;
	return v2616
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v87 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = v7
	goto L36
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 != v9 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v79 = F_slice_del(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L31
	} else {
		goto L32
	}
L5:
	;
	goto L4
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v14))))
	if v18 == int32(39) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L12
L9:
	;
	goto L8
L10:
	;
	if v72 < int32(0) {
		goto L2
	} else {
		goto L30
	}
L12:
	;
	goto L13
L13:
	;
	goto L14
L14:
	;
	v27 = v9
	v29 = int32(1)
	goto L17
L16:
	;
	v72 = v57
	goto L10
L17:
	;
	if v15 <= v27 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v72 = int32(-1)
	goto L10
L20:
	;
	goto L21
L21:
	;
	v34 = v27 + int32(1)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v27))))
	if base.Ui32(v36) < base.Ui32(int32(192)) {
		v57 = v34
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v58 = int32(1)
	if v58 < v29 {
		v27 = v57
		v29 = v29 - v58
		goto L17
	} else {
		goto L29
	}
L23:
	;
	if v15 <= v34 {
		v57 = v34
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v43 = v34
	goto L25
L25:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14+v43))))
	if int32(-65) < v46 {
		v57 = v43
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v57 = v15
	goto L22
L27:
	;
	v50 = v43 + int32(1)
	if v50 != v15 {
		v43 = v50
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L18
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
	v9 = v72
	goto L3
L31:
	;
	return int32(0)
L32:
	;
	if v79 < int32(0) {
		v2616 = v79
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	if v204 < int32(0) {
		v2616 = v87
		goto L1
	} else {
		goto L59
	}
L35:
	;
	v204 = v176
	goto L34
L36:
	;
	if v100 <= v109 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v204 = int32(-1)
	goto L34
L39:
	;
	goto L40
L40:
	;
	v116 = int32(1)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v101))))
	if base.Ui32(v118) < base.Ui32(int32(192)) {
		v175 = v118
		v176 = v116
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if int32(305) < v175 {
		goto L54
	} else {
		goto L55
	}
L42:
	;
	v122 = v109 + int32(1)
	if v122 == v100 {
		v175 = v118
		v176 = v116
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v101))))
	v127 = v125 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v118) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v101))))
	v143 = v141 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v118) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v131 = v109 + int32(2)
	if v131 != v100 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v175 = v118<<(uint(int32(6))%32)&int32(1984) | v127
	v176 = int32(2)
	goto L41
L48:
	;
	goto L47
L49:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v147))))
	v175 = v160&int32(63) | (v118<<(uint(int32(18))%32)&int32(_a_F_turkish_UTF_8_stem_0) | v127<<(uint(int32(12))%32) | v143<<(uint(int32(6))%32))
	v176 = int32(4)
	goto L41
L50:
	;
	v147 = v109 + int32(3)
	if v147 != v100 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v175 = v118<<(uint(int32(12))%32)&int32(_a_F_turkish_UTF_8_stem_1) | v127<<(uint(int32(6))%32) | v143
	v176 = int32(3)
	goto L41
L53:
	;
	goto L52
L54:
	;
	v193 = v176 + v109
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v193
	v109 = v193
	goto L36
L55:
	;
	v180 = v175 - int32(97)
	if v180 < int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v180)>>(uint(int32(3))%32)))+uint32(_c_F_turkish_UTF_8_stem[0]))))
	if int32(base.Ui32(v186)>>(uint(v180&int32(7))%32))&int32(1) != 0 {
		goto L35
	} else {
		goto L57
	}
L57:
	;
	goto L54
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v208 = v207 + v204
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v208
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = v208
	goto L62
L60:
	;
	if v326 < int32(0) {
		v2616 = v87
		goto L1
	} else {
		goto L85
	}
L61:
	;
	v326 = v298
	goto L60
L62:
	;
	if v222 <= v231 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v326 = int32(-1)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v238 = int32(1)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+v223))))
	if base.Ui32(v240) < base.Ui32(int32(192)) {
		v297 = v240
		v298 = v238
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if int32(305) < v297 {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v244 = v231 + int32(1)
	if v244 == v222 {
		v297 = v240
		v298 = v238
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v223))))
	v249 = v247 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v240) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v223))))
	v265 = v263 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v240) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v253 = v231 + int32(2)
	if v253 != v222 {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v297 = v240<<(uint(int32(6))%32)&int32(1984) | v249
	v298 = int32(2)
	goto L67
L74:
	;
	goto L73
L75:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223+v269))))
	v297 = v282&int32(63) | (v240<<(uint(int32(18))%32)&int32(_a_F_turkish_UTF_8_stem_0) | v249<<(uint(int32(12))%32) | v265<<(uint(int32(6))%32))
	v298 = int32(4)
	goto L67
L76:
	;
	v269 = v231 + int32(3)
	if v269 != v222 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v297 = v240<<(uint(int32(12))%32)&int32(_a_F_turkish_UTF_8_stem_1) | v249<<(uint(int32(6))%32) | v265
	v298 = int32(3)
	goto L67
L79:
	;
	goto L78
L80:
	;
	v315 = v298 + v231
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v315
	v231 = v315
	goto L62
L81:
	;
	v302 = v297 - int32(97)
	if v302 < int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v302)>>(uint(int32(3))%32)))+uint32(_c_F_turkish_UTF_8_stem[0]))))
	if int32(base.Ui32(v308)>>(uint(v302&int32(7))%32))&int32(1) != 0 {
		goto L61
	} else {
		goto L83
	}
L83:
	;
	goto L80
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v330
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v333))) = int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v338 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v338 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v883
	v885 = int32(0)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	if v887 == v885 {
		v2616 = v885
		goto L1
	} else {
		goto L248
	}
L87:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v873
	v875 = F_slice_del(m, l0)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L31
	} else {
		goto L246
	}
L88:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v363 = v337 - v336
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v362 - v363
	v366 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v366 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v342-int32(3) <= v341 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346+v342-int32(1)))))
	if v350 != int32(159) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v355 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_2), int32(4))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L31
	} else {
		goto L92
	}
L92:
	;
	if v355 == int32(0) {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	v360 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L94
L94:
	;
	if v360 != 0 {
		goto L87
	} else {
		goto L95
	}
L95:
	;
	goto L88
L96:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v378 = v377 - v363
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v378
	v381 = v378 - int32(1)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v381 <= v382 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v371 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_3), int32(32))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L31
	} else {
		goto L98
	}
L98:
	;
	if v371 == int32(0) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v376 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L100
L100:
	;
	if v376 != 0 {
		goto L87
	} else {
		goto L101
	}
L101:
	;
	goto L96
L102:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v408 = v407 - v363
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v408
	v410 = int32(3)
	v412 = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v408-v415 < v410 {
		v425 = v412
		goto L110
	} else {
		goto L111
	}
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384+v381))))
	if base.B2i32(v386&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v386)%32)&int32(_a_F_turkish_UTF_8_stem_4) == int32(0)) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v400 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_5), int32(8))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L31
	} else {
		goto L105
	}
L105:
	;
	if v400 == int32(0) {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v405 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L107
L107:
	;
	if v405 != 0 {
		goto L87
	} else {
		goto L108
	}
L108:
	;
	goto L102
L109:
	;
	if v425 != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L109
L111:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v421 = F_memcmp(m, v418+v408-v410, int32(_a_F_turkish_UTF_8_stem_6), v410)
	mBase = m.M
	if v421 != 0 {
		v425 = v412
		goto L110
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v408 - v410
	v425 = int32(1)
	goto L110
L113:
	;
	v427 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L116
L114:
	;
	goto L115
L115:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v429 = v428 - v363
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v429-int32(5) <= v431 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	if v427 != 0 {
		goto L87
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v555 - v363
	v558 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v558 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L119:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v429-int32(1)))))
	switch v439 - int32(97) {
	case 0, 4:
		goto L120
	default:
		goto L118
	}
L120:
	;
	v444 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_7), int32(2))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L31
	} else {
		goto L121
	}
L121:
	;
	if v444 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v450-int32(4) <= v449 {
		v466 = v448
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v530 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v530 == int32(0) {
		goto L118
	} else {
		goto L147
	}
L124:
	;
	v467 = v448 - v450
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v466 - v467
	v470 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v470 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454+v450-int32(1)))))
	if v458 != int32(122) {
		v466 = v448
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v463 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_8), int32(4))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L31
	} else {
		goto L127
	}
L127:
	;
	if v463 != 0 {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v466 = v465
	goto L124
L129:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v490 - v467
	v493 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v493 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v474-int32(2) <= v473 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478+v474-int32(1)))))
	if v482 != int32(114) {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v487 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_9), int32(2))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L31
	} else {
		goto L133
	}
L133:
	;
	if v487 != 0 {
		goto L123
	} else {
		goto L134
	}
L134:
	;
	goto L129
L135:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v515 - v467
	v518 = F_r_mark_sUn(m, l0)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L31
	} else {
		goto L143
	}
L136:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v496 - int32(1)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v498 <= v499 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501+v498))))
	if v503 != int32(109) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v508 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_10), int32(4))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L31
	} else {
		goto L139
	}
L139:
	;
	if v508 == int32(0) {
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v513 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L141
L141:
	;
	if v513 != 0 {
		goto L123
	} else {
		goto L142
	}
L142:
	;
	goto L135
L143:
	;
	if v518 != 0 {
		goto L123
	} else {
		goto L144
	}
L144:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v520 - v467
	v523 = F_r_mark_yUz(m, l0)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L31
	} else {
		goto L145
	}
L145:
	;
	if v523 != 0 {
		goto L123
	} else {
		goto L146
	}
L146:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v525 - v467
	goto L123
L147:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v534-int32(3) <= v533 {
		goto L118
	} else {
		goto L148
	}
L148:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538+v534-int32(1)))))
	if v542 != int32(159) {
		goto L118
	} else {
		goto L149
	}
L149:
	;
	v547 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_2), int32(4))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L31
	} else {
		goto L150
	}
L150:
	;
	if v547 == int32(0) {
		goto L118
	} else {
		goto L151
	}
L151:
	;
	v552 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L152
L152:
	;
	if v552 != 0 {
		goto L87
	} else {
		goto L153
	}
L153:
	;
	goto L118
L154:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v642 - v363
	v645 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v645 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L155:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v562-int32(2) <= v561 {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566+v562-int32(1)))))
	if v570 != int32(114) {
		goto L154
	} else {
		goto L157
	}
L157:
	;
	v575 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_9), int32(2))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L31
	} else {
		goto L158
	}
L158:
	;
	if v575 == int32(0) {
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v579
	v581 = F_slice_del(m, l0)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L31
	} else {
		goto L160
	}
L160:
	;
	if v581 < int32(0) {
		v2616 = v581
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v585
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v588 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v588 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v638))) = int32(0)
	goto L87
L163:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v609 = v587 - v585
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v608 - v609
	v612 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v612 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v592-int32(2) <= v591 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596+v592-int32(1)))))
	if v600 != int32(114) {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	v605 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_11), int32(8))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L31
	} else {
		goto L167
	}
L167:
	;
	if v605 != 0 {
		goto L162
	} else {
		goto L168
	}
L168:
	;
	goto L163
L169:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v623 - v609
	v626 = F_r_mark_ysA(m, l0)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L31
	} else {
		goto L175
	}
L170:
	;
	v617 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_3), int32(32))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L31
	} else {
		goto L171
	}
L171:
	;
	if v617 == int32(0) {
		goto L169
	} else {
		goto L172
	}
L172:
	;
	v622 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L173
L173:
	;
	if v622 != 0 {
		goto L162
	} else {
		goto L174
	}
L174:
	;
	goto L169
L175:
	;
	if v626 != 0 {
		goto L162
	} else {
		goto L176
	}
L176:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v628 - v609
	v631 = F_r_mark_ymUs_(m, l0)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L31
	} else {
		goto L177
	}
L177:
	;
	if v631 != 0 {
		goto L162
	} else {
		goto L178
	}
L178:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v633 + (v585 - v587)
	goto L162
L179:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v688 = v687 - v363
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v688
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v688-int32(4) <= v690 {
		v707 = v688
		goto L195
	} else {
		goto L196
	}
L180:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v649-int32(2) <= v648 {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v649-int32(1)))))
	if v657 != int32(122) {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	v662 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_12), int32(4))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L31
	} else {
		goto L183
	}
L183:
	;
	if v662 == int32(0) {
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v668 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v668 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v679 + (v667 - v666)
	v683 = F_r_mark_ysA(m, l0)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L31
	} else {
		goto L191
	}
L186:
	;
	v673 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_3), int32(32))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L31
	} else {
		goto L187
	}
L187:
	;
	if v673 == int32(0) {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v678 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L189
L189:
	;
	if v678 != 0 {
		goto L87
	} else {
		goto L190
	}
L190:
	;
	goto L185
L191:
	;
	if v683 != 0 {
		goto L87
	} else {
		goto L192
	}
L192:
	;
	goto L179
L193:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v781 - v363
	v784 = int32(0)
	v785 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v785 == v784 {
		v806 = v784
		goto L222
	} else {
		goto L223
	}
L194:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v744
	v746 = F_slice_del(m, l0)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L31
	} else {
		goto L212
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v707
	v709 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v709 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694+v688-int32(1)))))
	if v698 != int32(122) {
		v707 = v688
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v703 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_8), int32(4))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L31
	} else {
		goto L198
	}
L198:
	;
	if v703 != 0 {
		goto L194
	} else {
		goto L199
	}
L199:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v707 = v705 - v363
	goto L195
L200:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v731 - v363
	v734 = F_r_mark_sUn(m, l0)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L31
	} else {
		goto L208
	}
L201:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v714 = v712 - int32(1)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v714 <= v715 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717+v714))))
	if v719 != int32(122) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v724 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_13), int32(4))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L31
	} else {
		goto L204
	}
L204:
	;
	if v724 == int32(0) {
		goto L200
	} else {
		goto L205
	}
L205:
	;
	v729 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L206
L206:
	;
	if v729 != 0 {
		goto L194
	} else {
		goto L207
	}
L207:
	;
	goto L200
L208:
	;
	if v734 != 0 {
		goto L194
	} else {
		goto L209
	}
L209:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v736 - v363
	v739 = F_r_mark_yUm(m, l0)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L31
	} else {
		goto L210
	}
L210:
	;
	if v739 == int32(0) {
		goto L193
	} else {
		goto L211
	}
L211:
	;
	goto L194
L212:
	;
	if v746 < int32(0) {
		v2616 = v746
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v750
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v753 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v753 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v777 + (v750 - v752)
	goto L87
L215:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v757-int32(3) <= v756 {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761+v757-int32(1)))))
	if v765 != int32(159) {
		goto L214
	} else {
		goto L217
	}
L217:
	;
	v770 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_2), int32(4))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L31
	} else {
		goto L218
	}
L218:
	;
	if v770 == int32(0) {
		goto L214
	} else {
		goto L219
	}
L219:
	;
	v775 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L220
L220:
	;
	if v775 != 0 {
		goto L87
	} else {
		goto L221
	}
L221:
	;
	goto L214
L222:
	;
	if v806 == int32(0) {
		goto L86
	} else {
		goto L227
	}
L223:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v789-int32(2) <= v788 {
		v806 = v784
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793+v789-int32(1)))))
	if v797 != int32(114) {
		v806 = v784
		goto L222
	} else {
		goto L225
	}
L225:
	;
	v802 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_11), int32(8))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L31
	} else {
		goto L226
	}
L226:
	;
	v806 = base.B2i32(v802 != int32(0))
	goto L222
L227:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v810
	v812 = F_slice_del(m, l0)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L31
	} else {
		goto L228
	}
L228:
	;
	if v812 < int32(0) {
		v2616 = v812
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v816
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v816-int32(4) <= v819 {
		v837 = v2
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if v837 != 0 {
		goto L234
	} else {
		goto L235
	}
L231:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824+v816-int32(1)))))
	if v828 != int32(122) {
		v837 = v2
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v833 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_8), int32(4))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L31
	} else {
		goto L233
	}
L233:
	;
	v837 = base.B2i32(v833 != int32(0))
	goto L230
L234:
	;
	v863 = F_r_mark_ymUs_(m, l0)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L31
	} else {
		goto L244
	}
L235:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v839 = v818 - v816
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v838 - v839
	v842 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L31
	} else {
		goto L236
	}
L236:
	;
	if v842 != 0 {
		goto L234
	} else {
		goto L237
	}
L237:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v844 - v839
	v847 = F_r_mark_yUm(m, l0)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L31
	} else {
		goto L238
	}
L238:
	;
	if v847 != 0 {
		goto L234
	} else {
		goto L239
	}
L239:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v849 - v839
	v852 = F_r_mark_sUn(m, l0)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L31
	} else {
		goto L240
	}
L240:
	;
	if v852 != 0 {
		goto L234
	} else {
		goto L241
	}
L241:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v854 - v839
	v857 = F_r_mark_yUz(m, l0)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L31
	} else {
		goto L242
	}
L242:
	;
	if v857 != 0 {
		goto L234
	} else {
		goto L243
	}
L243:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v859 - v839
	goto L234
L244:
	;
	if v863 != 0 {
		goto L87
	} else {
		goto L245
	}
L245:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v865 + (v816 - v818)
	goto L87
L246:
	;
	if v875 < int32(0) {
		v2616 = v875
		goto L1
	} else {
		goto L247
	}
L247:
	;
	goto L86
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v883
	v891 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v891 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2243
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2243
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2247
	v2249 = int32(2)
	v2251 = int32(0)
	if v2247-v2243 < v2249 {
		v2264 = v2251
		goto L669
	} else {
		goto L670
	}
L250:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v931
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v931
	v934 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v934 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L251:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v895-int32(2) <= v894 {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v899+v895-int32(1)))))
	if v903 != int32(114) {
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v908 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_9), int32(2))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L31
	} else {
		goto L254
	}
L254:
	;
	if v908 == int32(0) {
		goto L250
	} else {
		goto L255
	}
L255:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v912
	v914 = F_slice_del(m, l0)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L31
	} else {
		goto L256
	}
L256:
	;
	if v914 < int32(0) {
		v2616 = v914
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v920 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L31
	} else {
		goto L258
	}
L258:
	;
	if v920 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v924 + (v918 - v919)
	goto L249
L260:
	;
	goto L261
L261:
	;
	if int32(0) <= v920 {
		goto L249
	} else {
		goto L262
	}
L262:
	;
	v2616 = v920
	goto L1
L263:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2232
	v2234 = F_slice_del(m, l0)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L31
	} else {
		goto L666
	}
L264:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2226
	v2228 = F_slice_del(m, l0)
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L31
	} else {
		goto L664
	}
L265:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1144
	v1147 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1147 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L266:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v939 = v937 - int32(1)
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v939 <= v940 {
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942+v939))))
	switch v944 - int32(97) {
	case 0, 4:
		goto L268
	default:
		goto L265
	}
L268:
	;
	v949 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_14), int32(2))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L31
	} else {
		goto L269
	}
L269:
	;
	if v949 == int32(0) {
		goto L265
	} else {
		goto L270
	}
L270:
	;
	v954 = Fn13983(m, l0, int32(110))
	mBase = m.M
	goto L271
L271:
	;
	if v954 == int32(0) {
		goto L265
	} else {
		goto L272
	}
L272:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v957
	v959 = F_slice_del(m, l0)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L31
	} else {
		goto L273
	}
L273:
	;
	if v959 < int32(0) {
		v2616 = v959
		goto L1
	} else {
		goto L274
	}
L274:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v963
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v963-int32(3) <= v966 {
		v985 = v965
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v986 = v965 - v963
	v987 = v985 - v986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v987
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v987
	v990 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L31
	} else {
		goto L284
	}
L276:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970+v963-int32(1)))))
	if v974 != int32(177) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	if v974 != int32(105) {
		v985 = v965
		goto L275
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v981 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_15), int32(2))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L31
	} else {
		goto L281
	}
L280:
	;
	goto L279
L281:
	;
	if v981 != 0 {
		goto L264
	} else {
		goto L282
	}
L282:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v985 = v983
	goto L275
L283:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1118 = v1117 - v986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1118
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1118
	v1121 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L31
	} else {
		goto L322
	}
L284:
	;
	if v990 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v994 - v986
	v997 = int32(0)
	v1003 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1003 == v997 {
		goto L289
	} else {
		goto L290
	}
L286:
	;
	goto L287
L287:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1085
	v1087 = F_slice_del(m, l0)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L31
	} else {
		goto L309
	}
L288:
	;
	if v1082 == int32(0) {
		goto L283
	} else {
		goto L308
	}
L289:
	;
	v1082 = int32(0)
	goto L288
L290:
	;
	goto L291
L291:
	;
	v1011 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_16), int32(105), int32(305), int32(0))
	mBase = m.M
	if v1011 != 0 {
		v1075 = v997
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1082 = v1075
	goto L288
L293:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1014 <= v1015 {
		goto L297
	} else {
		goto L298
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1067
	v1075 = int32(1)
	goto L292
L295:
	;
	v1067 = v1024 - v1012 + v1031
	goto L294
L296:
	;
	v1039 = v1014 - v1012
	v1040 = v1036 + v1039
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1040
	if v1038 < v1040 {
		goto L302
	} else {
		goto L303
	}
L297:
	;
	v1036 = v1012
	v1037 = v1013
	v1038 = v1015
	goto L296
L298:
	;
	goto L299
L299:
	;
	v1020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013+v1014-int32(1)))))
	if v1020 != int32(115) {
		v1036 = v1012
		v1037 = v1013
		v1038 = v1015
		goto L296
	} else {
		goto L300
	}
L300:
	;
	v1024 = v1014 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1024
	v1029 = int32(0)
	v1030 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), v1029)
	mBase = m.M
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1030 == v1029 {
		goto L295
	} else {
		goto L301
	}
L301:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1036 = v1031
	v1037 = v1034
	v1038 = v1035
	goto L296
L302:
	;
	v1046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040+v1037-int32(1)))))
	if v1046 == int32(115) {
		v1075 = v997
		goto L292
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1050 = F_skip_b_utf8(m, v1037, v1040, v1038, int32(1))
	mBase = m.M
	if v1050 < int32(0) {
		v1075 = v997
		goto L292
	} else {
		goto L306
	}
L305:
	;
	goto L304
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1050
	v1058 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1058 != 0 {
		v1075 = v997
		goto L292
	} else {
		goto L307
	}
L307:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1067 = v1059 + v1039
	goto L294
L308:
	;
	goto L287
L309:
	;
	if v1087 < int32(0) {
		v2616 = v1087
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1091
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1094 = v1093 - v1091
	v1095 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L31
	} else {
		goto L311
	}
L311:
	;
	if v1095 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1099 - v1094
	goto L249
L313:
	;
	goto L314
L314:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1102
	v1104 = F_slice_del(m, l0)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L31
	} else {
		goto L315
	}
L315:
	;
	if v1104 < int32(0) {
		v2616 = v1104
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1108 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L31
	} else {
		goto L317
	}
L317:
	;
	if v1108 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1112 - v1094
	goto L249
L319:
	;
	goto L320
L320:
	;
	if int32(0) <= v1108 {
		goto L249
	} else {
		goto L321
	}
L321:
	;
	v2616 = v1108
	goto L1
L322:
	;
	if v1121 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1125 - v986
	goto L249
L324:
	;
	goto L325
L325:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1128
	v1130 = F_slice_del(m, l0)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L31
	} else {
		goto L326
	}
L326:
	;
	if v1130 < int32(0) {
		v2616 = v1130
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1134 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L31
	} else {
		goto L328
	}
L328:
	;
	if v1134 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1138 - v986
	goto L249
L330:
	;
	goto L331
L331:
	;
	if int32(0) <= v1134 {
		goto L249
	} else {
		goto L332
	}
L332:
	;
	v2616 = v1134
	goto L1
L333:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1344
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1344
	v1347 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1347 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L334:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1191-int32(3) <= v1190 {
		v1210 = v1189
		goto L346
	} else {
		goto L347
	}
L335:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1167
	v1169 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1169 == int32(0) {
		goto L333
	} else {
		goto L341
	}
L336:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1151-int32(2) <= v1150 {
		goto L335
	} else {
		goto L337
	}
L337:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155+v1151-int32(1)))))
	switch v1159 - int32(97) {
	case 0, 4:
		goto L338
	default:
		goto L335
	}
L338:
	;
	v1164 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_18), int32(2))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L31
	} else {
		goto L339
	}
L339:
	;
	if v1164 != 0 {
		goto L334
	} else {
		goto L340
	}
L340:
	;
	goto L335
L341:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1174 = v1172 - int32(1)
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1174 <= v1175 {
		goto L333
	} else {
		goto L342
	}
L342:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1177+v1174))))
	switch v1179 - int32(97) {
	case 0, 4:
		goto L343
	default:
		goto L333
	}
L343:
	;
	v1184 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_19), int32(2))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L31
	} else {
		goto L344
	}
L344:
	;
	if v1184 == int32(0) {
		goto L333
	} else {
		goto L345
	}
L345:
	;
	goto L334
L346:
	;
	v1211 = v1189 - v1191
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1210 - v1211
	v1214 = int32(0)
	v1220 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1220 == v1214 {
		goto L355
	} else {
		goto L356
	}
L347:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1195+v1191-int32(1)))))
	if v1199 != int32(177) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	if v1199 != int32(105) {
		v1210 = v1189
		goto L346
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	v1206 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_15), int32(2))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L31
	} else {
		goto L352
	}
L351:
	;
	goto L350
L352:
	;
	if v1206 != 0 {
		goto L263
	} else {
		goto L353
	}
L353:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1210 = v1208
	goto L346
L354:
	;
	if v1299 != 0 {
		goto L374
	} else {
		goto L375
	}
L355:
	;
	v1299 = int32(0)
	goto L354
L356:
	;
	goto L357
L357:
	;
	v1228 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_16), int32(105), int32(305), int32(0))
	mBase = m.M
	if v1228 != 0 {
		v1292 = v1214
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1299 = v1292
	goto L354
L359:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1231 <= v1232 {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1284
	v1292 = int32(1)
	goto L358
L361:
	;
	v1284 = v1241 - v1229 + v1248
	goto L360
L362:
	;
	v1256 = v1231 - v1229
	v1257 = v1253 + v1256
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1257
	if v1255 < v1257 {
		goto L368
	} else {
		goto L369
	}
L363:
	;
	v1253 = v1229
	v1254 = v1230
	v1255 = v1232
	goto L362
L364:
	;
	goto L365
L365:
	;
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1230+v1231-int32(1)))))
	if v1237 != int32(115) {
		v1253 = v1229
		v1254 = v1230
		v1255 = v1232
		goto L362
	} else {
		goto L366
	}
L366:
	;
	v1241 = v1231 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1241
	v1246 = int32(0)
	v1247 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), v1246)
	mBase = m.M
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1247 == v1246 {
		goto L361
	} else {
		goto L367
	}
L367:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1253 = v1248
	v1254 = v1251
	v1255 = v1252
	goto L362
L368:
	;
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257+v1254-int32(1)))))
	if v1263 == int32(115) {
		v1292 = v1214
		goto L358
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v1267 = F_skip_b_utf8(m, v1254, v1257, v1255, int32(1))
	mBase = m.M
	if v1267 < int32(0) {
		v1292 = v1214
		goto L358
	} else {
		goto L372
	}
L371:
	;
	goto L370
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1267
	v1275 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1275 != 0 {
		v1292 = v1214
		goto L358
	} else {
		goto L373
	}
L373:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1284 = v1276 + v1256
	goto L360
L374:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1300
	v1302 = F_slice_del(m, l0)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L31
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1332 - v1211
	v1335 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L31
	} else {
		goto L390
	}
L377:
	;
	if v1302 < int32(0) {
		v2616 = v1302
		goto L1
	} else {
		goto L378
	}
L378:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1306
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1309 = v1308 - v1306
	v1310 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L31
	} else {
		goto L379
	}
L379:
	;
	if v1310 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1314 - v1309
	goto L249
L381:
	;
	goto L382
L382:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1317
	v1319 = F_slice_del(m, l0)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L31
	} else {
		goto L383
	}
L383:
	;
	if v1319 < int32(0) {
		v2616 = v1319
		goto L1
	} else {
		goto L384
	}
L384:
	;
	v1323 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L31
	} else {
		goto L385
	}
L385:
	;
	if v1323 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1327 - v1309
	goto L249
L387:
	;
	goto L388
L388:
	;
	if int32(0) <= v1323 {
		goto L249
	} else {
		goto L389
	}
L389:
	;
	v2616 = v1323
	goto L1
L390:
	;
	if v1335 == int32(0) {
		goto L333
	} else {
		goto L391
	}
L391:
	;
	if int32(0) <= v1335 {
		goto L249
	} else {
		goto L392
	}
L392:
	;
	v2616 = v1335
	goto L1
L393:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1507
	v1510 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1510 == int32(0) {
		goto L442
	} else {
		goto L443
	}
L394:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1381 = int32(0)
	v1387 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1387 == v1381 {
		goto L405
	} else {
		goto L406
	}
L395:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1367
	v1369 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1369 == int32(0) {
		goto L393
	} else {
		goto L401
	}
L396:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1351-int32(3) <= v1350 {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1355+v1351-int32(1)))))
	if v1359 != int32(110) {
		goto L395
	} else {
		goto L398
	}
L398:
	;
	v1364 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_20), int32(2))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L31
	} else {
		goto L399
	}
L399:
	;
	if v1364 != 0 {
		goto L394
	} else {
		goto L400
	}
L400:
	;
	goto L395
L401:
	;
	v1374 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_21), int32(4))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L31
	} else {
		goto L402
	}
L402:
	;
	if v1374 == int32(0) {
		goto L393
	} else {
		goto L403
	}
L403:
	;
	goto L394
L404:
	;
	if v1466 != 0 {
		goto L424
	} else {
		goto L425
	}
L405:
	;
	v1466 = int32(0)
	goto L404
L406:
	;
	goto L407
L407:
	;
	v1395 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_16), int32(105), int32(305), int32(0))
	mBase = m.M
	if v1395 != 0 {
		v1459 = v1381
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1466 = v1459
	goto L404
L409:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1398 <= v1399 {
		goto L413
	} else {
		goto L414
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1451
	v1459 = int32(1)
	goto L408
L411:
	;
	v1451 = v1408 - v1396 + v1415
	goto L410
L412:
	;
	v1423 = v1398 - v1396
	v1424 = v1420 + v1423
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1424
	if v1422 < v1424 {
		goto L418
	} else {
		goto L419
	}
L413:
	;
	v1420 = v1396
	v1421 = v1397
	v1422 = v1399
	goto L412
L414:
	;
	goto L415
L415:
	;
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1397+v1398-int32(1)))))
	if v1404 != int32(115) {
		v1420 = v1396
		v1421 = v1397
		v1422 = v1399
		goto L412
	} else {
		goto L416
	}
L416:
	;
	v1408 = v1398 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1408
	v1413 = int32(0)
	v1414 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), v1413)
	mBase = m.M
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1414 == v1413 {
		goto L411
	} else {
		goto L417
	}
L417:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1420 = v1415
	v1421 = v1418
	v1422 = v1419
	goto L412
L418:
	;
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1424+v1421-int32(1)))))
	if v1430 == int32(115) {
		v1459 = v1381
		goto L408
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	v1434 = F_skip_b_utf8(m, v1421, v1424, v1422, int32(1))
	mBase = m.M
	if v1434 < int32(0) {
		v1459 = v1381
		goto L408
	} else {
		goto L422
	}
L421:
	;
	goto L420
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1434
	v1442 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1442 != 0 {
		v1459 = v1381
		goto L408
	} else {
		goto L423
	}
L423:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1451 = v1443 + v1423
	goto L410
L424:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1467
	v1469 = F_slice_del(m, l0)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L31
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1499 + (v1379 - v1380)
	v1503 = F_r_mark_lArI(m, l0)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L31
	} else {
		goto L440
	}
L427:
	;
	if v1469 < int32(0) {
		v2616 = v1469
		goto L1
	} else {
		goto L428
	}
L428:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1473
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1476 = v1475 - v1473
	v1477 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L31
	} else {
		goto L429
	}
L429:
	;
	if v1477 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1481 - v1476
	goto L249
L431:
	;
	goto L432
L432:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1484
	v1486 = F_slice_del(m, l0)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L31
	} else {
		goto L433
	}
L433:
	;
	if v1486 < int32(0) {
		v2616 = v1486
		goto L1
	} else {
		goto L434
	}
L434:
	;
	v1490 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L31
	} else {
		goto L435
	}
L435:
	;
	if v1490 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1494 - v1476
	goto L249
L437:
	;
	goto L438
L438:
	;
	if int32(0) <= v1490 {
		goto L249
	} else {
		goto L439
	}
L439:
	;
	v2616 = v1490
	goto L1
L440:
	;
	if v1503 != 0 {
		goto L249
	} else {
		goto L441
	}
L441:
	;
	goto L393
L442:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1598
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1598
	v1601 = int32(0)
	v1602 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1602 == v1601 {
		v1623 = v1601
		goto L477
	} else {
		goto L478
	}
L443:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1514-int32(2) <= v1513 {
		goto L442
	} else {
		goto L444
	}
L444:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1514-int32(1)))))
	if v1522 != int32(110) {
		goto L442
	} else {
		goto L445
	}
L445:
	;
	v1527 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_22), int32(4))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L31
	} else {
		goto L446
	}
L446:
	;
	if v1527 == int32(0) {
		goto L442
	} else {
		goto L447
	}
L447:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1531
	v1533 = F_slice_del(m, l0)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L31
	} else {
		goto L448
	}
L448:
	;
	if v1533 < int32(0) {
		v2616 = v1533
		goto L1
	} else {
		goto L449
	}
L449:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1537
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1540 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L31
	} else {
		goto L450
	}
L450:
	;
	if v1540 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1542
	v1544 = F_slice_del(m, l0)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L31
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1575 = v1539 - v1537
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1574 - v1575
	v1578 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L31
	} else {
		goto L467
	}
L454:
	;
	if v1544 < int32(0) {
		v2616 = v1544
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1548
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1551 = v1550 - v1548
	v1552 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L31
	} else {
		goto L456
	}
L456:
	;
	if v1552 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1556 - v1551
	goto L249
L458:
	;
	goto L459
L459:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1559
	v1561 = F_slice_del(m, l0)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L31
	} else {
		goto L460
	}
L460:
	;
	if v1561 < int32(0) {
		v2616 = v1561
		goto L1
	} else {
		goto L461
	}
L461:
	;
	v1565 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L31
	} else {
		goto L462
	}
L462:
	;
	if v1565 == int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1569 - v1551
	goto L249
L464:
	;
	goto L465
L465:
	;
	if int32(0) <= v1565 {
		goto L249
	} else {
		goto L466
	}
L466:
	;
	v2616 = v1565
	goto L1
L467:
	;
	if v1578 != 0 {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1580
	v1582 = F_slice_del(m, l0)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L31
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1590 - v1575
	v1593 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L31
	} else {
		goto L475
	}
L471:
	;
	if v1582 < int32(0) {
		v2616 = v1582
		goto L1
	} else {
		goto L472
	}
L472:
	;
	v1586 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L31
	} else {
		goto L473
	}
L473:
	;
	if int32(0) <= v1586 {
		goto L249
	} else {
		goto L474
	}
L474:
	;
	v2616 = v1586
	goto L1
L475:
	;
	if int32(0) <= v1593 {
		goto L249
	} else {
		goto L476
	}
L476:
	;
	v2616 = v1593
	goto L1
L477:
	;
	if v1623 == int32(0) {
		goto L486
	} else {
		goto L487
	}
L478:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1607 = v1605 - int32(1)
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1607 <= v1608 {
		v1623 = v1601
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610+v1607))))
	if v1612 != int32(110) {
		v1623 = v1601
		goto L477
	} else {
		goto L480
	}
L480:
	;
	v1617 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_23), int32(4))
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L31
	} else {
		goto L481
	}
L481:
	;
	if v1617 == int32(0) {
		v1623 = v1601
		goto L477
	} else {
		goto L482
	}
L482:
	;
	v1622 = Fn13983(m, l0, int32(110))
	mBase = m.M
	goto L483
L483:
	;
	v1623 = v1622
	goto L477
L484:
	;
	if v2219 < int32(0) {
		v2616 = v2219
		goto L1
	} else {
		goto L663
	}
L485:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1823
	v1826 = F_r_mark_lArI(m, l0)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L31
	} else {
		goto L548
	}
L486:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1627
	v1629 = int32(0)
	v1630 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1630 == v1629 {
		v1651 = v1629
		goto L489
	} else {
		goto L490
	}
L487:
	;
	goto L488
L488:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1657
	v1659 = F_slice_del(m, l0)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L31
	} else {
		goto L497
	}
L489:
	;
	if v1651 == int32(0) {
		goto L485
	} else {
		goto L496
	}
L490:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1635 = v1633 - int32(1)
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1635 <= v1636 {
		v1651 = v1629
		goto L489
	} else {
		goto L491
	}
L491:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1638+v1635))))
	switch v1640 - int32(97) {
	case 0, 4:
		goto L492
	default:
		v1651 = v1629
		goto L489
	}
L492:
	;
	v1645 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_24), int32(2))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L31
	} else {
		goto L493
	}
L493:
	;
	if v1645 == int32(0) {
		v1651 = v1629
		goto L489
	} else {
		goto L494
	}
L494:
	;
	v1650 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L495
L495:
	;
	v1651 = v1650
	goto L489
L496:
	;
	goto L488
L497:
	;
	if v1659 < int32(0) {
		v2616 = v1659
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1663
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1666 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L31
	} else {
		goto L500
	}
L499:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1693 = v1665 - v1663
	v1694 = v1692 - v1693
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1694
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1694
	v1697 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L31
	} else {
		goto L512
	}
L500:
	;
	if v1666 == int32(0) {
		goto L499
	} else {
		goto L501
	}
L501:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1670
	v1672 = F_slice_del(m, l0)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L31
	} else {
		goto L502
	}
L502:
	;
	if v1672 < int32(0) {
		v2616 = v1672
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v1676 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L31
	} else {
		goto L504
	}
L504:
	;
	v1679 = int32(base.Ui32(v1676) >> (uint(int32(31)) % 32))
	if v1676 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v1681 = v1679
	goto L507
L506:
	;
	v1681 = int32(36)
	goto L507
L507:
	;
	if v1681 == int32(0) {
		goto L249
	} else {
		goto L508
	}
L508:
	;
	if v1681 == int32(36) {
		goto L499
	} else {
		goto L509
	}
L509:
	;
	if v1679 != 0 {
		v2219 = v1676 >> (uint(int32(31)) % 32) & v1676
		goto L484
	} else {
		goto L510
	}
L510:
	;
	goto L249
L511:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1814 - v1693
	v1817 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L31
	} else {
		goto L545
	}
L512:
	;
	if v1697 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1701 - v1693
	v1704 = int32(0)
	v1710 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1710 == v1704 {
		goto L517
	} else {
		goto L518
	}
L514:
	;
	goto L515
L515:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1792
	v1794 = F_slice_del(m, l0)
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L31
	} else {
		goto L537
	}
L516:
	;
	if v1789 == int32(0) {
		goto L511
	} else {
		goto L536
	}
L517:
	;
	v1789 = int32(0)
	goto L516
L518:
	;
	goto L519
L519:
	;
	v1718 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_16), int32(105), int32(305), int32(0))
	mBase = m.M
	if v1718 != 0 {
		v1782 = v1704
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v1789 = v1782
	goto L516
L521:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1721 <= v1722 {
		goto L525
	} else {
		goto L526
	}
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1774
	v1782 = int32(1)
	goto L520
L523:
	;
	v1774 = v1731 - v1719 + v1738
	goto L522
L524:
	;
	v1746 = v1721 - v1719
	v1747 = v1743 + v1746
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1747
	if v1745 < v1747 {
		goto L530
	} else {
		goto L531
	}
L525:
	;
	v1743 = v1719
	v1744 = v1720
	v1745 = v1722
	goto L524
L526:
	;
	goto L527
L527:
	;
	v1727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1720+v1721-int32(1)))))
	if v1727 != int32(115) {
		v1743 = v1719
		v1744 = v1720
		v1745 = v1722
		goto L524
	} else {
		goto L528
	}
L528:
	;
	v1731 = v1721 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1731
	v1736 = int32(0)
	v1737 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), v1736)
	mBase = m.M
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1737 == v1736 {
		goto L523
	} else {
		goto L529
	}
L529:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1743 = v1738
	v1744 = v1741
	v1745 = v1742
	goto L524
L530:
	;
	v1753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747+v1744-int32(1)))))
	if v1753 == int32(115) {
		v1782 = v1704
		goto L520
	} else {
		goto L533
	}
L531:
	;
	goto L532
L532:
	;
	v1757 = F_skip_b_utf8(m, v1744, v1747, v1745, int32(1))
	mBase = m.M
	if v1757 < int32(0) {
		v1782 = v1704
		goto L520
	} else {
		goto L534
	}
L533:
	;
	goto L532
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1757
	v1765 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), int32(0))
	mBase = m.M
	if v1765 != 0 {
		v1782 = v1704
		goto L520
	} else {
		goto L535
	}
L535:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1774 = v1766 + v1746
	goto L522
L536:
	;
	goto L515
L537:
	;
	if v1794 < int32(0) {
		v2616 = v1794
		goto L1
	} else {
		goto L538
	}
L538:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1798
	v1800 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L31
	} else {
		goto L539
	}
L539:
	;
	if v1800 == int32(0) {
		goto L249
	} else {
		goto L540
	}
L540:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1804
	v1806 = F_slice_del(m, l0)
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L31
	} else {
		goto L541
	}
L541:
	;
	if v1806 < int32(0) {
		v2616 = v1806
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v1810 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L31
	} else {
		goto L543
	}
L543:
	;
	if int32(0) <= v1810 {
		goto L249
	} else {
		goto L544
	}
L544:
	;
	v2616 = v1810
	goto L1
L545:
	;
	if int32(0) <= v1817 {
		goto L249
	} else {
		goto L546
	}
L546:
	;
	if int32(base.Ui32(v1817)>>(uint(int32(31))%32)) != 0 {
		v2219 = v1817
		goto L484
	} else {
		goto L547
	}
L547:
	;
	goto L249
L548:
	;
	if v1826 != 0 {
		goto L549
	} else {
		goto L550
	}
L549:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1828
	v1830 = F_slice_del(m, l0)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L31
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1834
	v1836 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L31
	} else {
		goto L554
	}
L552:
	;
	if int32(0) <= v1830 {
		goto L249
	} else {
		goto L553
	}
L553:
	;
	v2616 = v1830
	goto L1
L554:
	;
	v1839 = int32(base.Ui32(v1836) >> (uint(int32(31)) % 32))
	if v1836 != 0 {
		goto L555
	} else {
		goto L556
	}
L555:
	;
	v1841 = v1839
	goto L557
L556:
	;
	v1841 = int32(44)
	goto L557
L557:
	;
	if v1841 == int32(0) {
		goto L249
	} else {
		goto L558
	}
L558:
	;
	v1846 = v1836 >> (uint(int32(31)) % 32) & v1836
	if v1841 != int32(44) {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v2215 = v1839
	v2216 = v1846
	goto L561
L560:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1849
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1849
	v1852 = int32(0)
	v1853 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1853 == v1852 {
		v1873 = v1852
		goto L562
	} else {
		goto L563
	}
L561:
	;
	if v2215 == int32(0) {
		goto L249
	} else {
		goto L662
	}
L562:
	;
	if v1873 != 0 {
		goto L568
	} else {
		goto L569
	}
L563:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1858 = v1856 - int32(1)
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1858 <= v1859 {
		v1873 = v1852
		goto L562
	} else {
		goto L564
	}
L564:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1861+v1858))))
	switch v1863 - int32(97) {
	case 0, 4:
		goto L565
	default:
		v1873 = v1852
		goto L562
	}
L565:
	;
	v1868 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_25), int32(4))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L31
	} else {
		goto L566
	}
L566:
	;
	v1873 = base.B2i32(v1868 != int32(0))
	goto L562
L567:
	;
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2090
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2090
	v2093 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L31
	} else {
		goto L626
	}
L568:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2042
	v2044 = F_slice_del(m, l0)
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L31
	} else {
		goto L609
	}
L569:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1874
	v1876 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v1876 != 0 {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L575
L571:
	;
	v2011 = int32(0)
	goto L572
L572:
	;
	if v2011 != 0 {
		goto L568
	} else {
		goto L600
	}
L573:
	;
	if v2005 != 0 {
		goto L596
	} else {
		goto L597
	}
L574:
	;
	v2005 = v1998
	goto L573
L575:
	;
	if v1889 <= v1890 {
		v1998 = int32(-1)
		goto L574
	} else {
		goto L577
	}
L576:
	;
	v1998 = int32(0)
	goto L574
L577:
	;
	v1907 = int32(1)
	v1908 = v1889 - v1907
	v1910 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1891+v1908))))
	v1912 = v1910 & int32(255)
	if base.B2i32(v1908 == v1890)|base.B2i32(int32(0) <= v1910) != 0 {
		v1970 = v1912
		v1974 = v1907
		goto L578
	} else {
		goto L579
	}
L578:
	;
	if int32(305) < v1970 {
		goto L586
	} else {
		goto L587
	}
L579:
	;
	v1919 = v1912 & int32(63)
	v1921 = v1889 - int32(2)
	v1923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1891+v1921))))
	v1925 = v1923 << (uint(int32(6)) % 32)
	if base.B2i32(v1921 != v1890)&base.B2i32(base.Ui32(v1923) < base.Ui32(int32(192))) == int32(0) {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v1970 = v1925&int32(1984) | v1919
	v1974 = int32(2)
	goto L578
L581:
	;
	goto L582
L582:
	;
	v1938 = v1925&int32(4032) | v1919
	v1940 = v1889 - int32(3)
	v1942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1891+v1940))))
	if base.B2i32(v1940 != v1890)&base.B2i32(base.Ui32(v1942) < base.Ui32(int32(224))) == int32(0) {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v1970 = v1942<<(uint(int32(12))%32)&int32(_a_F_turkish_UTF_8_stem_1) | v1938
	v1974 = int32(3)
	goto L578
L584:
	;
	goto L585
L585:
	;
	v1960 = int32(4)
	v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1889+v1891-v1960))))
	v1970 = v1942<<(uint(int32(12))%32)&int32(_a_F_turkish_UTF_8_stem_26) | v1962&int32(7)<<(uint(int32(18))%32) | v1938
	v1974 = v1960
	goto L578
L586:
	;
	v2005 = v1974
	goto L573
L587:
	;
	goto L588
L588:
	;
	v1976 = v1970 - int32(105)
	if v1976 < int32(0) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v2005 = v1974
	goto L573
L590:
	;
	goto L591
L591:
	;
	v1982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1976)>>(uint(int32(3))%32)))+uint32(_c_F_turkish_UTF_8_stem[1]))))
	if int32(base.Ui32(v1982)>>(uint(v1976&int32(7))%32))&int32(1) == int32(0) {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v2005 = v1974
	goto L573
L593:
	;
	goto L594
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1889 - v1974
	goto L595
L595:
	;
	goto L576
L596:
	;
	v2009 = int32(0)
	goto L598
L597:
	;
	v2008 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L599
L598:
	;
	v2011 = v2009
	goto L572
L599:
	;
	v2009 = v2008
	goto L598
L600:
	;
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2012
	v2014 = int32(0)
	v2015 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2015 == v2014 {
		v2037 = v2014
		goto L601
	} else {
		goto L602
	}
L601:
	;
	if v2037 == int32(0) {
		goto L567
	} else {
		goto L608
	}
L602:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2018 <= v2019 {
		v2037 = v2014
		goto L601
	} else {
		goto L603
	}
L603:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2021+v2018-int32(1)))))
	switch v2025 - int32(97) {
	case 0, 4:
		goto L604
	default:
		v2037 = v2014
		goto L601
	}
L604:
	;
	v2030 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_27), int32(2))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L31
	} else {
		goto L605
	}
L605:
	;
	if v2030 == int32(0) {
		v2037 = v2014
		goto L601
	} else {
		goto L606
	}
L606:
	;
	v2035 = Fn13983(m, l0, int32(121))
	mBase = m.M
	goto L607
L607:
	;
	v2037 = v2035
	goto L601
L608:
	;
	goto L568
L609:
	;
	if v2044 < int32(0) {
		v2616 = v2044
		goto L1
	} else {
		goto L610
	}
L610:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2048
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2051 = F_r_mark_possessives(m, l0)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L31
	} else {
		goto L612
	}
L611:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2078
	v2080 = F_slice_del(m, l0)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L31
	} else {
		goto L622
	}
L612:
	;
	if v2051 != 0 {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2053
	v2055 = F_slice_del(m, l0)
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L31
	} else {
		goto L616
	}
L614:
	;
	goto L615
L615:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2068 + (v2048 - v2050)
	v2072 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L31
	} else {
		goto L620
	}
L616:
	;
	if v2055 < int32(0) {
		v2616 = v2055
		goto L1
	} else {
		goto L617
	}
L617:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2059
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2062 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L31
	} else {
		goto L618
	}
L618:
	;
	if v2062 != 0 {
		goto L611
	} else {
		goto L619
	}
L619:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2064 + (v2059 - v2061)
	goto L611
L620:
	;
	if v2072 == int32(0) {
		goto L249
	} else {
		goto L621
	}
L621:
	;
	goto L611
L622:
	;
	if v2080 < int32(0) {
		v2616 = v2080
		goto L1
	} else {
		goto L623
	}
L623:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2084
	v2086 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L31
	} else {
		goto L624
	}
L624:
	;
	if int32(0) <= v2086 {
		goto L249
	} else {
		goto L625
	}
L625:
	;
	v2616 = v2086
	goto L1
L626:
	;
	if v2093 == int32(0) {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2097
	v2099 = int32(0)
	v2105 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v2105 == v2099 {
		goto L631
	} else {
		goto L632
	}
L628:
	;
	goto L629
L629:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2187
	v2189 = F_slice_del(m, l0)
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L31
	} else {
		goto L651
	}
L630:
	;
	if v2184 == int32(0) {
		goto L249
	} else {
		goto L650
	}
L631:
	;
	v2184 = int32(0)
	goto L630
L632:
	;
	goto L633
L633:
	;
	v2113 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_16), int32(105), int32(305), int32(0))
	mBase = m.M
	if v2113 != 0 {
		v2177 = v2099
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v2184 = v2177
	goto L630
L635:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2116 <= v2117 {
		goto L639
	} else {
		goto L640
	}
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2169
	v2177 = int32(1)
	goto L634
L637:
	;
	v2169 = v2126 - v2114 + v2133
	goto L636
L638:
	;
	v2141 = v2116 - v2114
	v2142 = v2138 + v2141
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2142
	if v2140 < v2142 {
		goto L644
	} else {
		goto L645
	}
L639:
	;
	v2138 = v2114
	v2139 = v2115
	v2140 = v2117
	goto L638
L640:
	;
	goto L641
L641:
	;
	v2122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2115+v2116-int32(1)))))
	if v2122 != int32(115) {
		v2138 = v2114
		v2139 = v2115
		v2140 = v2117
		goto L638
	} else {
		goto L642
	}
L642:
	;
	v2126 = v2116 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2126
	v2131 = int32(0)
	v2132 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), v2131)
	mBase = m.M
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2132 == v2131 {
		goto L637
	} else {
		goto L643
	}
L643:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2138 = v2133
	v2139 = v2136
	v2140 = v2137
	goto L638
L644:
	;
	v2148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2142+v2139-int32(1)))))
	if v2148 == int32(115) {
		v2177 = v2099
		goto L634
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	v2152 = F_skip_b_utf8(m, v2139, v2142, v2140, int32(1))
	mBase = m.M
	if v2152 < int32(0) {
		v2177 = v2099
		goto L634
	} else {
		goto L648
	}
L647:
	;
	goto L646
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2152
	v2160 = F_in_grouping_b_U(m, l0, int32(_a_F_turkish_UTF_8_stem_17), int32(97), int32(305), int32(0))
	mBase = m.M
	if v2160 != 0 {
		v2177 = v2099
		goto L634
	} else {
		goto L649
	}
L649:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2169 = v2161 + v2141
	goto L636
L650:
	;
	goto L629
L651:
	;
	if v2189 < int32(0) {
		v2616 = v2189
		goto L1
	} else {
		goto L652
	}
L652:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2193
	v2195 = F_r_mark_lAr(m, l0)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L31
	} else {
		goto L653
	}
L653:
	;
	if v2195 == int32(0) {
		goto L249
	} else {
		goto L654
	}
L654:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2199
	v2201 = F_slice_del(m, l0)
	mBase = m.M
	v2202 = m.ExcPending
	if v2202 != 0 {
		goto L31
	} else {
		goto L655
	}
L655:
	;
	if v2201 < int32(0) {
		v2616 = v2201
		goto L1
	} else {
		goto L656
	}
L656:
	;
	v2205 = F_r_stem_suffix_chain_before_ki(m, l0)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L31
	} else {
		goto L657
	}
L657:
	;
	if v2205 == int32(0) {
		goto L249
	} else {
		goto L658
	}
L658:
	;
	if v2205 < int32(0) {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v2213 = v2205
	goto L661
L660:
	;
	v2213 = v1846
	goto L661
L661:
	;
	v2215 = int32(base.Ui32(v2205) >> (uint(int32(31)) % 32))
	v2216 = v2213
	goto L561
L662:
	;
	v2219 = v2216
	goto L484
L663:
	;
	goto L249
L664:
	;
	if int32(0) <= v2228 {
		goto L249
	} else {
		goto L665
	}
L665:
	;
	v2616 = v2228
	goto L1
L666:
	;
	if v2234 < int32(0) {
		v2616 = v2234
		goto L1
	} else {
		goto L667
	}
L667:
	;
	goto L249
L668:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2264 == int32(0) {
		goto L674
	} else {
		goto L675
	}
L669:
	;
	goto L668
L670:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2260 = F_memcmp(m, v2257+v2247-v2249, int32(_a_F_turkish_UTF_8_stem_28), v2249)
	mBase = m.M
	if v2260 != 0 {
		v2264 = v2251
		goto L669
	} else {
		goto L671
	}
L671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2247 - v2249
	v2264 = int32(1)
	goto L669
L672:
	;
	v2616 = v2614
	goto L1
L673:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2297
	if v2297 <= v2296 {
		goto L686
	} else {
		goto L687
	}
L674:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2296 = v2268
	v2297 = v2265
	goto L673
L675:
	;
	goto L676
L676:
	;
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2270 = int32(3)
	v2272 = int32(0)
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2269-v2275 < v2270 {
		v2285 = v2272
		goto L679
	} else {
		goto L680
	}
L677:
	;
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2291 <= v2293 {
		v2614 = int32(0)
		goto L672
	} else {
		goto L685
	}
L678:
	;
	if v2285 != 0 {
		goto L682
	} else {
		goto L683
	}
L679:
	;
	goto L678
L680:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2281 = F_memcmp(m, v2278+v2269-v2270, int32(_a_F_turkish_UTF_8_stem_29), v2270)
	mBase = m.M
	if v2281 != 0 {
		v2285 = v2272
		goto L679
	} else {
		goto L681
	}
L681:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2269 - v2270
	v2285 = int32(1)
	goto L679
L682:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2291 = v2286
	goto L677
L683:
	;
	goto L684
L684:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2289 = v2287 + (v2269 - v2265)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2289
	v2291 = v2289
	goto L677
L685:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2296 = v2293
	v2297 = v2295
	goto L673
L686:
	;
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2572
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2572
	v2577 = F_find_among_b(m, l0, int32(_a_F_turkish_UTF_8_stem_30), int32(4))
	mBase = m.M
	v2578 = m.ExcPending
	if v2578 != 0 {
		goto L31
	} else {
		goto L743
	}
L687:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303+v2297-int32(1)))))
	switch v2307 - int32(100) {
	case 0, 3:
		goto L688
	default:
		goto L686
	}
L688:
	;
	v2311 = v2297 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2311
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2336 = v2311
	goto L691
L689:
	;
	if v2442 < int32(0) {
		goto L686
	} else {
		goto L707
	}
L690:
	;
	v2442 = int32(-1)
	goto L689
L691:
	;
	if v2336 <= v2326 {
		goto L690
	} else {
		goto L693
	}
L693:
	;
	v2343 = int32(1)
	v2344 = v2336 - v2343
	v2346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2327+v2344))))
	v2348 = v2346 & int32(255)
	if base.B2i32(v2344 == v2326)|base.B2i32(int32(0) <= v2346) != 0 {
		v2406 = v2348
		v2410 = v2343
		goto L694
	} else {
		goto L695
	}
L694:
	;
	if int32(305) < v2406 {
		goto L702
	} else {
		goto L703
	}
L695:
	;
	v2355 = v2348 & int32(63)
	v2357 = v2336 - int32(2)
	v2359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2327+v2357))))
	v2361 = v2359 << (uint(int32(6)) % 32)
	if base.B2i32(v2357 != v2326)&base.B2i32(base.Ui32(v2359) < base.Ui32(int32(192))) == int32(0) {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	v2406 = v2361&int32(1984) | v2355
	v2410 = int32(2)
	goto L694
L697:
	;
	goto L698
L698:
	;
	v2374 = v2361&int32(4032) | v2355
	v2376 = v2336 - int32(3)
	v2378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2327+v2376))))
	if base.B2i32(v2376 != v2326)&base.B2i32(base.Ui32(v2378) < base.Ui32(int32(224))) == int32(0) {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v2406 = v2378<<(uint(int32(12))%32)&int32(_a_F_turkish_UTF_8_stem_1) | v2374
	v2410 = int32(3)
	goto L694
L700:
	;
	goto L701
L701:
	;
	v2396 = int32(4)
	v2398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2336+v2327-v2396))))
	v2406 = v2378<<(uint(int32(12))%32)&int32(_a_F_turkish_UTF_8_stem_26) | v2398&int32(7)<<(uint(int32(18))%32) | v2374
	v2410 = v2396
	goto L694
L702:
	;
	v2427 = v2336 - v2410
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2427
	v2336 = v2427
	goto L691
L703:
	;
	v2412 = v2406 - int32(97)
	if v2412 < int32(0) {
		goto L702
	} else {
		goto L704
	}
L704:
	;
	v2418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2412)>>(uint(int32(3))%32)))+uint32(_c_F_turkish_UTF_8_stem[0]))))
	if int32(base.Ui32(v2418)>>(uint(v2412&int32(7))%32))&int32(1) == int32(0) {
		goto L702
	} else {
		goto L705
	}
L705:
	;
	v2442 = v2410
	goto L689
L707:
	;
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2446 <= v2447 {
		goto L710
	} else {
		goto L711
	}
L708:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2484 = v2445 - v2446
	v2485 = v2483 - v2484
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2485
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2485 <= v2487 {
		goto L720
	} else {
		goto L721
	}
L709:
	;
	v2479 = F_slice_from_s(m, l0, int32(2), int32(_a_F_turkish_UTF_8_stem_31))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L31
	} else {
		goto L718
	}
L710:
	;
	v2459 = int32(2)
	v2461 = int32(0)
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2463-v2464 < v2459 {
		v2474 = v2461
		goto L714
	} else {
		goto L715
	}
L711:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2449+v2446-int32(1)))))
	if v2453 != int32(97) {
		goto L710
	} else {
		goto L712
	}
L712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2446 - int32(1)
	goto L709
L713:
	;
	if v2474 == int32(0) {
		goto L708
	} else {
		goto L717
	}
L714:
	;
	goto L713
L715:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2470 = F_memcmp(m, v2467+v2463-v2459, int32(_a_F_turkish_UTF_8_stem_32), v2459)
	mBase = m.M
	if v2470 != 0 {
		v2474 = v2461
		goto L714
	} else {
		goto L716
	}
L716:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2463 - v2459
	v2474 = int32(1)
	goto L714
L717:
	;
	goto L709
L718:
	;
	if int32(0) <= v2479 {
		goto L686
	} else {
		goto L719
	}
L719:
	;
	v2614 = v2479
	goto L672
L720:
	;
	v2524 = int32(2)
	v2526 = int32(0)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2528-v2529 < v2524 {
		v2539 = v2526
		goto L729
	} else {
		goto L730
	}
L721:
	;
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2491 = int32(1)
	v2493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2489+v2485-v2491))))
	v2495 = v2493 - int32(101)
	switch (v2495<<(uint(int32(7))%32) | int32(base.Ui32(v2495&int32(254))>>(uint(v2491)%32))) & int32(255) {
	case 0, 2:
		goto L723
	default:
		goto L720
	case 5, 8:
		goto L722
	}
L722:
	;
	v2514 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2485 - v2514
	v2519 = F_slice_from_s(m, l0, v2514, int32(_a_F_turkish_UTF_8_stem_33))
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L31
	} else {
		goto L726
	}
L723:
	;
	v2505 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2485 - v2505
	v2510 = F_slice_from_s(m, l0, v2505, int32(_a_F_turkish_UTF_8_stem_34))
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L31
	} else {
		goto L724
	}
L724:
	;
	if int32(0) <= v2510 {
		goto L686
	} else {
		goto L725
	}
L725:
	;
	v2614 = v2510
	goto L672
L726:
	;
	if int32(0) <= v2519 {
		goto L686
	} else {
		goto L727
	}
L727:
	;
	v2614 = v2519
	goto L672
L728:
	;
	if v2539 == int32(0) {
		goto L732
	} else {
		goto L733
	}
L729:
	;
	goto L728
L730:
	;
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2535 = F_memcmp(m, v2532+v2528-v2524, int32(_a_F_turkish_UTF_8_stem_35), v2524)
	mBase = m.M
	if v2535 != 0 {
		v2539 = v2526
		goto L729
	} else {
		goto L731
	}
L731:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2528 - v2524
	v2539 = int32(1)
	goto L729
L732:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2543 = v2542 - v2484
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2543
	v2545 = int32(2)
	v2547 = int32(0)
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2543-v2550 < v2545 {
		v2560 = v2547
		goto L736
	} else {
		goto L737
	}
L733:
	;
	goto L734
L734:
	;
	v2565 = F_slice_from_s(m, l0, int32(2), int32(_a_F_turkish_UTF_8_stem_36))
	mBase = m.M
	v2566 = m.ExcPending
	if v2566 != 0 {
		goto L31
	} else {
		goto L740
	}
L735:
	;
	if v2560 == int32(0) {
		goto L686
	} else {
		goto L739
	}
L736:
	;
	goto L735
L737:
	;
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2556 = F_memcmp(m, v2553+v2543-v2545, int32(_a_F_turkish_UTF_8_stem_37), v2545)
	mBase = m.M
	if v2556 != 0 {
		v2560 = v2547
		goto L736
	} else {
		goto L738
	}
L738:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2543 - v2545
	v2560 = int32(1)
	goto L736
L739:
	;
	goto L734
L740:
	;
	if v2565 < int32(0) {
		v2614 = v2565
		goto L672
	} else {
		goto L741
	}
L741:
	;
	goto L686
L742:
	;
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2610
	v2614 = int32(1)
	goto L672
L743:
	;
	if v2577 == int32(0) {
		goto L742
	} else {
		goto L744
	}
L744:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2581
	switch v2577 - int32(1) {
	case 0:
		goto L748
	case 1:
		goto L747
	case 2:
		goto L746
	case 3:
		goto L745
	default:
		goto L742
	}
L745:
	;
	v2605 = F_slice_from_s(m, l0, int32(1), int32(_a_F_turkish_UTF_8_stem_38))
	mBase = m.M
	v2606 = m.ExcPending
	if v2606 != 0 {
		goto L31
	} else {
		goto L755
	}
L746:
	;
	v2599 = F_slice_from_s(m, l0, int32(1), int32(_a_F_turkish_UTF_8_stem_39))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L31
	} else {
		goto L753
	}
L747:
	;
	v2593 = F_slice_from_s(m, l0, int32(2), int32(_a_F_turkish_UTF_8_stem_40))
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L31
	} else {
		goto L751
	}
L748:
	;
	v2587 = F_slice_from_s(m, l0, int32(1), int32(_a_F_turkish_UTF_8_stem_41))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L31
	} else {
		goto L749
	}
L749:
	;
	if int32(0) <= v2587 {
		goto L742
	} else {
		goto L750
	}
L750:
	;
	v2614 = v2587
	goto L672
L751:
	;
	if int32(0) <= v2593 {
		goto L742
	} else {
		goto L752
	}
L752:
	;
	v2614 = v2593
	goto L672
L753:
	;
	if int32(0) <= v2599 {
		goto L742
	} else {
		goto L754
	}
L754:
	;
	v2614 = v2599
	goto L672
L755:
	;
	if v2605 < int32(0) {
		v2614 = v2605
		goto L672
	} else {
		goto L756
	}
L756:
	;
	goto L742
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
				F_errmsg_internal(m, int32(_a_F_typeidType_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_typeidType_1), int32(584), int32(_a_F_typeidType_2))
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14014(m, l0, int32(_a_F_typeidTypeRelid_0), int32(676), int32(_a_F_typeidTypeRelid_1), int32(_a_F_typeidTypeRelid_2), int32(82))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
