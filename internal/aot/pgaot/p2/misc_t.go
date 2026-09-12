package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TransferPredicateLocksToNewTarget(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v327 int64
	_ = v327
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v27 = F_get_hash_value(m, v26, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v33 = F_get_hash_value(m, v32, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	v41 = v36 + v27&int32(15)<<(uint(int32(7))%32)
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	v45 = F_LWLockAcquire(m, v43, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v61 = int32(_a_F_TransferPredicateLocksToNewTarget_0)
	v62 = v41 + v61
	v67 = v36 + v33&int32(15)<<(uint(int32(7))%32)
	v69 = v67 + v61
	if base.Ui32(v41) < base.Ui32(v67) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[3]))
	v55 = F_hash_search_with_hash_value(m, v48, int32(_a_F_TransferPredicateLocksToNewTarget_1), v51, int32(2), v23+int32(13))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	F_LWLockRelease(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v88 = int32(0)
	v90 = F_hash_search_with_hash_value(m, v87, l0, v27, v88, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	v73 = F_LWLockAcquire(m, v62, l2^int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v79 = F_LWLockAcquire(m, v69, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v76 = F_LWLockAcquire(m, v69, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	if base.Ui32(v41) <= base.Ui32(v67) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v84 = F_LWLockAcquire(m, v62, l2^int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	if v90 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v98 = F_hash_search_with_hash_value(m, v94, l1, v33, int32(3), v23+int32(12))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v414 = int32(1)
	goto L22
L22:
	;
	if base.Ui32(v41) < base.Ui32(v67) {
		goto L84
	} else {
		goto L85
	}
L23:
	;
	v414 = base.B2i32(v386 == int32(0))
	goto L22
L24:
	;
	if v98 == int32(0) {
		v386 = int32(1)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+12)))
	if v102 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v106 = v98 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v106
	goto L28
L27:
	;
	goto L28
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v98
	v116 = F_LWLockAcquire(m, v111+int32(3584), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v118 = int32(16)
	v119 = v98 + v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v122 = v90 + v118
	if v120 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v123 = v120
	goto L32
L31:
	;
	v123 = v122
	goto L32
L32:
	;
	v140 = int32(0)
	v143 = v123
	goto L33
L33:
	;
	if v122 != v143 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	F_LWLockRelease(m, v354+int32(3584))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L75
	}
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v143)+16))
	v149 = v143 - int32(4)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v150
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	goto L34
L38:
	;
	v152 = int32(8)
	v153 = v143 + v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[4]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v175 = F_hash_search_with_hash_value(m, v165, v143-v152, v168<<(uint(int32(4))%32)^v27, int32(2), v23+int32(12))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v178 = v150
	goto L40
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[4]))
	v182 = int32(4)
	v190 = F_hash_search_with_hash_value(m, v181, v23+v182, v178<<(uint(v182)%32)^v33, int32(3), v23+int32(12))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v178 = v177
	goto L40
L42:
	;
	if v190 != 0 {
		v140 = v348
		v143 = v146
		goto L33
	} else {
		goto L74
	}
L43:
	;
	if v190 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	F_LWLockRelease(m, v195+int32(3584))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+12)))
	if v297 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L47:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	v205 = F_LWLockAcquire(m, v201+int32(3584), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v207 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	F_LWLockRelease(m, v281+int32(3584))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L56
	}
L50:
	;
	if v207 == v119 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v211 = v207
	goto L52
L52:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v232 = int32(8)
	v233 = v211 + v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v239)+4)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v242
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[4]))
	v248 = int32(4)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v211-v248)))
	v257 = F_hash_search_with_hash_value(m, v245, v211-v232, v250<<(uint(v248)%32)^v33, int32(2), v23+int32(14))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L49
L54:
	;
	if v119 != v231 {
		v211 = v231
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v286 = int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v287 != v119 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v290 = v287
	goto L59
L58:
	;
	v290 = int32(0)
	goto L59
L59:
	;
	if v290 != 0 {
		v348 = v286
		goto L42
	} else {
		goto L60
	}
L60:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v295 = F_hash_search_with_hash_value(m, v292, v98, v33, int32(2), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v190 != 0 {
		v140 = v286
		v143 = v146
		goto L33
	} else {
		goto L62
	}
L62:
	;
	v386 = v286
	goto L23
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v190)+24)) = v147
	v348 = v140
	goto L42
L64:
	;
	v301 = v190 + int32(8)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v302 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v190)+24))
	if base.Ui64(v147) <= base.Ui64(v327) {
		v348 = v140
		goto L42
	} else {
		goto L73
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v119
	goto L69
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = v119
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v308)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v301
	v313 = v190 + int32(16)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v316 = v314 + int32(48)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v314)+52))
	if v317 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+52)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v314)+48)) = v316
	goto L72
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+20)) = v316
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+16)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v313
	goto L63
L73:
	;
	goto L63
L74:
	;
	v386 = v348
	goto L23
L75:
	;
	if l2 == int32(0) {
		v386 = v140
		goto L23
	} else {
		goto L76
	}
L76:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	if v361 != v122 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v364 = v361
	goto L79
L78:
	;
	v364 = int32(0)
	goto L79
L79:
	;
	if v364 != 0 {
		v386 = v140
		goto L23
	} else {
		goto L80
	}
L80:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v369 = F_hash_search_with_hash_value(m, v366, v90, v27, int32(2), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v386 = v140
	goto L23
L82:
	;
	F_LWLockRelease(m, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L89
	}
L83:
	;
	F_LWLockRelease(m, v417)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L88
	}
L84:
	;
	v417 = v69
	v418 = v62
	goto L83
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v41) <= base.Ui32(v67) {
		v422 = v69
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v417 = v62
	v418 = v69
	goto L83
L88:
	;
	v422 = v418
	goto L82
L89:
	;
	if l2 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	v428 = F_LWLockAcquire(m, v426, int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	m.G0 = v23 + int32(16)
	return v414 & int32(1)
L93:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[3]))
	v438 = F_hash_search_with_hash_value(m, v431, int32(_a_F_TransferPredicateLocksToNewTarget_1), v434, int32(1), v23+int32(15))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	F_LWLockRelease(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L92
}
func F_t_isalpha_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_pg_mblen_cstr(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_pg_mblen_with_len(m, l0, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != int32(1) {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_t_isalpha_cstr[0])))
				if v17 != int32(1) {
					F_char2wchar(m, v6+int32(4), int32(3), l0, v12, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						if base.Ui32(v35) <= base.Ui32(int32(_a_F_t_isalpha_cstr_0)) {
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(8))%32)))+uint32(_c_F_t_isalpha_cstr[1]))))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32))&int32(31)|v46<<(uint(int32(5))%32))+uint32(_c_F_t_isalpha_cstr[1]))))
							v60 = int32(base.Ui32(v52)>>(uint(v35&int32(7))%32)) & int32(1)
						} else {
							v60 = base.B2i32(base.Ui32(v35) < base.Ui32(int32(_a_F_t_isalpha_cstr_1)))
						}
						v61 = v60
						m.G0 = v6 + int32(16)
						return v61
					}
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					v61 = base.B2i32(base.Ui32((v20|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)))
					m.G0 = v6 + int32(16)
					return v61
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v61 = base.B2i32(base.Ui32((v20|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)))
				m.G0 = v6 + int32(16)
				return v61
			}
		}
	}
}
func F_tblspc_desc(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	switch v11 & int32(240) {
	case 0:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v10 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
		F_appendStringInfo(m, l0, int32(_a_F_tblspc_desc_0), v7)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	default:
		m.G0 = v7 + int32(32)
		return
	case 16:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v22
		F_appendStringInfo(m, l0, int32(_a_F_tblspc_desc_1), v7+int32(16))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_test_lockmode_for_conflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	v6 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v6)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(12))+uint32(_c_F_test_lockmode_for_conflict[0])))
	if base.Ui32(l1) < base.Ui32(int32(3)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v239
L2:
	;
	if v134 != 0 {
		v239 = int32(2)
		goto L1
	} else {
		goto L42
	}
L3:
	;
	v134 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[1]))
	if v25 == l1 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v134 = int32(1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[2]))
	if v29 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v134 = v126
	goto L2
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[3]))
	if v33 == int32(0) {
		v126 = v6
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[4]))
	v97 = int32(0)
	v99 = v29 - int32(1)
	goto L32
L13:
	;
	v38 = v33
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v43 == int32(4) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v126 = int32(0)
	goto L9
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	if v90 != 0 {
		v38 = v90
		goto L14
	} else {
		goto L31
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(1)
	if l1 == v46 {
		v126 = v49
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v53 = v51 - int32(1)
	if v53 < int32(0) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v58 = int32(0)
	v60 = v53
	goto L21
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v66 = int32(2)
	v67 = base.I32_div_s(v60-v58, v66)
	v68 = v67 + v58
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64+v68<<(uint(v66)%32))))
	if v72 == l1 {
		v126 = v49
		goto L9
	} else {
		goto L23
	}
L22:
	;
	goto L16
L23:
	;
	v76 = F_TransactionIdPrecedes(m, v72, l1)
	mBase = m.M
	if v76 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v77 = v68 + int32(1)
	goto L26
L25:
	;
	v77 = v58
	goto L26
L26:
	;
	if v76 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v80 = v60
	goto L29
L28:
	;
	v80 = v68 - int32(1)
	goto L29
L29:
	;
	if v77 <= v80 {
		v58 = v77
		v60 = v80
		goto L21
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	goto L15
L32:
	;
	v104 = int32(2)
	v105 = base.I32_div_s(v99-v97, v104)
	v106 = v105 + v97
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v95+v106<<(uint(v104)%32))))
	v111 = base.B2i32(v110 == l1)
	if v110 == l1 {
		v126 = v111
		goto L9
	} else {
		goto L34
	}
L33:
	;
	v126 = v111
	goto L9
L34:
	;
	v114 = base.B2i32(base.Ui32(v110) < base.Ui32(l1))
	if base.Ui32(v110) < base.Ui32(l1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v115 = v106 + int32(1)
	goto L37
L36:
	;
	v115 = v97
	goto L37
L37:
	;
	if base.Ui32(v110) < base.Ui32(l1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = v99
	goto L40
L39:
	;
	v118 = v106 - int32(1)
	goto L40
L40:
	;
	if v115 <= v118 {
		v97 = v115
		v99 = v118
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L33
L42:
	;
	v135 = F_TransactionIdIsInProgress(m, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v135 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v140 = int32(2)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v140)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v145 = int32(12)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144*v145)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(v140)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154*v145)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v149<<(uint(v140)%32))+uint32(_c_F_test_lockmode_for_conflict[7])))
	goto L48
L46:
	;
	goto L47
L47:
	;
	v174 = int32(0)
	v175 = F_TransactionIdDidAbort(m, l1)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L43
	} else {
		goto L50
	}
L48:
	;
	if int32(base.Ui32(v164)>>(uint(v159)%32))&int32(1) == int32(0) {
		v239 = int32(0)
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v170)
	return int32(0)
L50:
	;
	if v175 != 0 {
		v239 = v174
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v177 = F_TransactionIdDidCommit(m, l1)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(l0) < base.Ui32(int32(4)) {
		v239 = v174
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v177 == int32(0) {
		v239 = v174
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v183 = int32(2)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v183)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v188 = int32(12)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187*v188)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(v183)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197*v188)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v192<<(uint(v183)%32))+uint32(_c_F_test_lockmode_for_conflict[7])))
	goto L55
L55:
	;
	if int32(base.Ui32(v207)>>(uint(v202)%32))&int32(1) == int32(0) {
		v239 = v174
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v213 = int32(4)
	v216 = l3 + v213
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v219 = v217 + int32(12)
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	v222 = int32(16)
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+2)))
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	if v220|v221<<(uint(v222)%32) == v225|v226<<(uint(v222)%32) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	if v236 != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	goto L57
L59:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+4)))
	if v232 == v233 {
		v236 = int32(1)
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v236 = int32(0)
	goto L58
L62:
	;
	goto L61
L63:
	;
	v237 = v213
	goto L65
L64:
	;
	v237 = int32(3)
	goto L65
L65:
	;
	v239 = v237
	goto L1
}
func F_texticlike(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = F_Generic_Text_IC_like(m, v3, v8, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v11 == int32(1))
			}
		}
	}
}
func F_texticlike_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_like_regex_support(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_texticnlike(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = F_Generic_Text_IC_like(m, v3, v8, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v11 != int32(1))
			}
		}
	}
}
func F_texticregexeq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v18 = int32(1)
			v19 = v17 & v18
			if v17 == v18 {
				v22 = int32(4)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v24&int32(254) == int32(2) {
					v33 = v22
				} else {
					v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
				}
				if v24 == int32(1) {
					v36 = v22
				} else {
					v36 = v33
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v19 != 0 {
					v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v50 = F_RE_compile_and_cache(m, v15, int32(27), v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v56 = F_palloc(m, v47<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						v60 = v13
					} else {
						v60 = v8 + int32(4)
					}
					v61 = F_pg_mb2wchar_with_len(m, v60, v56, v47)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						v66 = F_RE_wchar_execute(m, v56, v61, v63, v63, v63)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v56)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v66
							}
						}
					}
				}
			}
		}
	}
}
func F_textin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5&int32(3) == int32(0) {
		v29 = v5
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v64 = v62 + int32(4)
	v65 = F_palloc(m, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v62 = v54 - v5
	goto L1
L3:
	;
	v33 = v29
	goto L12
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v13 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v62 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v18 = v5
	goto L8
L8:
	;
	v22 = v18 + int32(1)
	if v22&int32(3) == int32(0) {
		v29 = v22
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v54 = v22
	goto L2
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v27 != 0 {
		v18 = v22
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = int32(-2139062144)
	if (int32(16843008)-v39|v39)&v42 == v42 {
		v33 = v33 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v48 = v33
	goto L15
L14:
	;
	goto L13
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v52 != 0 {
		v48 = v48 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v54 = v48
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v64 << (uint(int32(2)) % 32)
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	return v65
L21:
	;
	v74 = F__emscripten_memcpy_bulkmem(m, v65+int32(4), v5, v62)
	mBase = m.M
	goto L23
L22:
	;
	goto L23
L23:
	;
	goto L20
}
func F_textltname(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1558), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6) >> (uint(int32(31)) % 32))
	}
}
func F_textne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = F_pg_newlocale_from_collation(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L82
	}
L4:
	;
	return int32(0)
L5:
	;
	v14 = int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v16 == v14 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return v213
L7:
	;
	F_pfree(m, v205)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L81
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_toast_raw_datum_size(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v115 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L45
	}
L11:
	;
	v22 = F_toast_raw_datum_size(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v20 != v22 {
		v213 = v14
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v25 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v27 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v29 = int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v31&v29 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v34 = v29
	goto L18
L17:
	;
	v34 = int32(4)
	goto L18
L18:
	;
	v35 = v25 + v34
	v36 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v38&v36 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v41 = v36
	goto L21
L20:
	;
	v41 = int32(4)
	goto L21
L21:
	;
	v42 = v27 + v41
	v43 = int32(4)
	v44 = v20 - v43
	if base.Ui32(v43) <= base.Ui32(v44) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v107 != v25 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v106 = int32(0)
	goto L22
L24:
	;
	v80 = v75
	v81 = v76
	v82 = v77
	goto L34
L25:
	;
	if (v35|v42)&int32(3) != 0 {
		v75 = v35
		v76 = v42
		v77 = v44
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v68 = v35
	v69 = v42
	v70 = v44
	goto L27
L27:
	;
	if v70 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v52 = v35
	v53 = v42
	v54 = v44
	goto L29
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v57 != v58 {
		v75 = v52
		v76 = v53
		v77 = v54
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v68 = v63
	v69 = v61
	v70 = v65
	goto L27
L31:
	;
	v60 = int32(4)
	v61 = v53 + v60
	v63 = v52 + v60
	v65 = v54 - v60
	if base.Ui32(int32(3)) < base.Ui32(v65) {
		v52 = v63
		v53 = v61
		v54 = v65
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v75 = v68
	v76 = v69
	v77 = v70
	goto L24
L34:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 == v86 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v106 = v85 - v86
	goto L22
L36:
	;
	v88 = int32(1)
	v93 = v82 - v88
	if v93 != 0 {
		v80 = v80 + v88
		v81 = v81 + v88
		v82 = v93
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L23
L40:
	;
	F_pfree(m, v25)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v112 = base.B2i32(v106 != int32(0))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27 != v113 {
		v204 = v112
		v205 = v27
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v213 = v112
	goto L6
L45:
	;
	v118 = v115 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v120 = F_pg_detoast_datum_packed(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v127 = v125 & int32(1)
	if v127 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v128 = v118
	goto L49
L48:
	;
	v128 = v115 + int32(4)
	goto L49
L49:
	;
	if v125 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v157 = int32(1)
	v158 = v120 + v157
	if v122&v157 != 0 {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v131 = int32(4)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v133&int32(254) == int32(2) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v146 = int32(1)
	if v127 != 0 {
		v156 = int32(base.Ui32(v125)>>(uint(v146)%32)) - v146
		goto L50
	} else {
		goto L60
	}
L54:
	;
	v142 = v131
	goto L56
L55:
	;
	v142 = base.B2i32(v133 == int32(18)) << (uint(v131) % 32)
	goto L56
L56:
	;
	if v133 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v145 = v131
	goto L59
L58:
	;
	v145 = v142
	goto L59
L59:
	;
	v156 = v145
	goto L50
L60:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v156 = int32(base.Ui32(v150)>>(uint(int32(2))%32)) - int32(4)
	goto L50
L61:
	;
	v163 = v158
	goto L63
L62:
	;
	v163 = v120 + int32(4)
	goto L63
L63:
	;
	if v122 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v194 = F_varstr_cmp(m, v128, v156, v163, v193, v9)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L75
	}
L65:
	;
	v166 = int32(4)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v168&int32(254) == int32(2) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v181 = int32(1)
	if v122&v181 != 0 {
		v193 = int32(base.Ui32(v122)>>(uint(v181)%32)) - v181
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v177 = v166
	goto L70
L69:
	;
	v177 = base.B2i32(v168 == int32(18)) << (uint(v166) % 32)
	goto L70
L70:
	;
	if v168 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v180 = v166
	goto L73
L72:
	;
	v180 = v177
	goto L73
L73:
	;
	v193 = v180
	goto L64
L74:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v193 = int32(base.Ui32(v187)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L75:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v196 != v115 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v115)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v201 = base.B2i32(v194 != int32(0))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v120 == v202 {
		v213 = v201
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v204 = v201
	v205 = v120
	goto L7
L81:
	;
	v213 = v204
	goto L6
L82:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_textne_0), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(_a_F_textne_1), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_textne_2), int32(1648), int32(_a_F_textne_3))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_throttle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = v5 + base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v9) <= base.Ui64(v7) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = base.I64_div_u_s(v7, v9)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v21 = m.G0
	v22 = int32(16)
	v23 = v21 - v22
	m.G0 = v23
	F___gettimeofday(m, v23)
	mBase = m.M
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+8)))
	m.G0 = v23 + v22
	goto L7
L5:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v75 = base.I64_rem_u_s(v73, v74)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v75
	v80 = m.G0
	v81 = int32(16)
	v82 = v80 - v81
	m.G0 = v82
	F___gettimeofday(m, v82)
	mBase = m.M
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
	m.G0 = v82 + v81
	goto L21
L6:
	;
	goto L5
L7:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = v36 - (v27 + v26*int64(1000000) - int64(946684800000000)) + v11*v12
	if v38 <= int64(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[1]))
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[0]))
	v53 = base.I64_div_u_s(v38, int64(1000))
	v56 = F_WaitLatch(m, v50, int32(41), base.I32_wrap_i64(v53), int32(150994944))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L16
	}
L13:
	;
	return
L14:
	;
	goto L12
L15:
	;
	if v56&int32(8) == int32(0) {
		goto L4
	} else {
		goto L20
	}
L16:
	;
	if v56&int32(1) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[1]))
	if v63 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	goto L6
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v86 + v85*int64(1000000) - int64(946684800000000)
	goto L3
}
func F_tidle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.B2i32(v27 <= int32(0))
}
func F_timeofday(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = m.G0
	v5 = v3 - int32(288)
	m.G0 = v5
	F___gettimeofday(m, v5+int32(272))
	mBase = m.M
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v5)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v10
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_timeofday[0]))
	v20 = F_pg_localtime(m, v5+int32(8), v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v24 = F_pg_strftime(m, v5+int32(144), int32(128), int32(_a_F_timeofday_0), v20)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+280))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26
			v33 = F_pg_snprintf(m, v5+int32(16), int32(128), v5+int32(144), v5)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v37 = F_cstring_to_text(m, v5+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(288)
					return v37
				}
			}
		}
	}
}
func F_timestamptz2timestamp(m *base.Module, l0 int64) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int64
	_ = v79
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v109 int64
	_ = v109
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v147 int64
	_ = v147
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if base.Ui64(int64(2)) <= base.Ui64(l0-int64(9223372036854775807)) {
		v20 = int32(0)
		v22 = F_timestamp2tm(m, l0, v8+int32(28), v8+int32(36), v8+int32(32), v20, v20)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			if v22 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_2), int32(_a_F_timestamptz2timestamp_3))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+32)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
				if v27 <= int32(-4713) {
					if v27 != int32(-4713) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
						if int32(10) < v32 {
							v43 = v32
							v45 = v8 + int32(8)
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							v51 = base.B2i32(int32(2) < v43)
							if int32(2) < v43 {
								v52 = int32(_a_F_timestamptz2timestamp_5)
							} else {
								v52 = int32(_a_F_timestamptz2timestamp_6)
							}
							v53 = v52 + v27
							v58 = base.I32_div_s(v53, int32(4))
							v61 = base.I32_div_s(v53, int32(-100))
							v64 = base.I32_div_s(v53, int32(400))
							if int32(2) < v43 {
								v68 = int32(1)
							} else {
								v68 = int32(13)
							}
							v73 = base.I32_div_s((v68+v43)*int32(_a_F_timestamptz2timestamp_7), int32(256))
							v79 = base.I64_extend_i32_s(v46 + v53*int32(365) + v58 + v61 + v64 + v73 - int32(_a_F_timestamptz2timestamp_8) - int32(_a_F_timestamptz2timestamp_9))
							v88 = int64(32)
							v89 = int64(20)
							v91 = int64(base.Ui64(v79) >> (uint(v88) % 64))
							v94 = int64(4294967295)
							v95 = int64(500654080)
							v97 = v79 & v94
							v98 = v95 * v97
							v102 = int64(base.Ui64(v98)>>(uint(v88)%64)) + v95*v91
							v109 = v97*v89 + v102&v94
							*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v79*int64(0) + v79>>(uint(int64(63))%64)*int64(86400000000) + v89*v91 + int64(base.Ui64(v102)>>(uint(v88)%64)) + int64(base.Ui64(v109)>>(uint(v88)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v45))) = v98&v94 | v109<<(uint(v88)%64)
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
							v121 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
							if v120 != v121>>(uint(int64(63))%64) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
								v128 = int32(60)
								v137 = base.I64_extend_i32_s(v125+(v126+v127*v128)*v128)*int64(1000000) + v26
								v140 = v137 + v121
								if base.B2i32(v137 < int64(0))^base.B2i32(v140 < v121) != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									if base.Ui64(v140-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
													mBase = m.M
													v189 = m.ExcPending
													if v189 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v147 = v140
										m.G0 = v8 + int32(80)
										return v147
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int64(0)
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
				} else {
					if v27 <= int32(_a_F_timestamptz2timestamp_10) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
						v43 = v37
						v45 = v8 + int32(8)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						v51 = base.B2i32(int32(2) < v43)
						if int32(2) < v43 {
							v52 = int32(_a_F_timestamptz2timestamp_5)
						} else {
							v52 = int32(_a_F_timestamptz2timestamp_6)
						}
						v53 = v52 + v27
						v58 = base.I32_div_s(v53, int32(4))
						v61 = base.I32_div_s(v53, int32(-100))
						v64 = base.I32_div_s(v53, int32(400))
						if int32(2) < v43 {
							v68 = int32(1)
						} else {
							v68 = int32(13)
						}
						v73 = base.I32_div_s((v68+v43)*int32(_a_F_timestamptz2timestamp_7), int32(256))
						v79 = base.I64_extend_i32_s(v46 + v53*int32(365) + v58 + v61 + v64 + v73 - int32(_a_F_timestamptz2timestamp_8) - int32(_a_F_timestamptz2timestamp_9))
						v88 = int64(32)
						v89 = int64(20)
						v91 = int64(base.Ui64(v79) >> (uint(v88) % 64))
						v94 = int64(4294967295)
						v95 = int64(500654080)
						v97 = v79 & v94
						v98 = v95 * v97
						v102 = int64(base.Ui64(v98)>>(uint(v88)%64)) + v95*v91
						v109 = v97*v89 + v102&v94
						*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v79*int64(0) + v79>>(uint(int64(63))%64)*int64(86400000000) + v89*v91 + int64(base.Ui64(v102)>>(uint(v88)%64)) + int64(base.Ui64(v109)>>(uint(v88)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v45))) = v98&v94 | v109<<(uint(v88)%64)
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
						if v120 != v121>>(uint(int64(63))%64) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
							v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
							v128 = int32(60)
							v137 = base.I64_extend_i32_s(v125+(v126+v127*v128)*v128)*int64(1000000) + v26
							v140 = v137 + v121
							if base.B2i32(v137 < int64(0))^base.B2i32(v140 < v121) != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								if base.Ui64(v140-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v147 = v140
									m.G0 = v8 + int32(80)
									return v147
								}
							}
						}
					} else {
						if v27 != int32(_a_F_timestamptz2timestamp_11) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
							if int32(5) < v40 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v43 = v40
								v45 = v8 + int32(8)
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
								v51 = base.B2i32(int32(2) < v43)
								if int32(2) < v43 {
									v52 = int32(_a_F_timestamptz2timestamp_5)
								} else {
									v52 = int32(_a_F_timestamptz2timestamp_6)
								}
								v53 = v52 + v27
								v58 = base.I32_div_s(v53, int32(4))
								v61 = base.I32_div_s(v53, int32(-100))
								v64 = base.I32_div_s(v53, int32(400))
								if int32(2) < v43 {
									v68 = int32(1)
								} else {
									v68 = int32(13)
								}
								v73 = base.I32_div_s((v68+v43)*int32(_a_F_timestamptz2timestamp_7), int32(256))
								v79 = base.I64_extend_i32_s(v46 + v53*int32(365) + v58 + v61 + v64 + v73 - int32(_a_F_timestamptz2timestamp_8) - int32(_a_F_timestamptz2timestamp_9))
								v88 = int64(32)
								v89 = int64(20)
								v91 = int64(base.Ui64(v79) >> (uint(v88) % 64))
								v94 = int64(4294967295)
								v95 = int64(500654080)
								v97 = v79 & v94
								v98 = v95 * v97
								v102 = int64(base.Ui64(v98)>>(uint(v88)%64)) + v95*v91
								v109 = v97*v89 + v102&v94
								*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v79*int64(0) + v79>>(uint(int64(63))%64)*int64(86400000000) + v89*v91 + int64(base.Ui64(v102)>>(uint(v88)%64)) + int64(base.Ui64(v109)>>(uint(v88)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v45))) = v98&v94 | v109<<(uint(v88)%64)
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
								v121 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
								if v120 != v121>>(uint(int64(63))%64) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
									v128 = int32(60)
									v137 = base.I64_extend_i32_s(v125+(v126+v127*v128)*v128)*int64(1000000) + v26
									v140 = v137 + v121
									if base.B2i32(v137 < int64(0))^base.B2i32(v140 < v121) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
													mBase = m.M
													v189 = m.ExcPending
													if v189 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										if base.Ui64(v140-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz2timestamp_0), int32(0))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_timestamptz2timestamp_1), int32(_a_F_timestamptz2timestamp_4), int32(_a_F_timestamptz2timestamp_3))
														mBase = m.M
														v189 = m.ExcPending
														if v189 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v147 = v140
											m.G0 = v8 + int32(80)
											return v147
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
		v147 = l0
		m.G0 = v8 + int32(80)
		return v147
	}
}
func F_timetztypmodout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) <= v7 {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(_a_F_timetztypmodout_0)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
		v14 = F_psprintf(m, int32(_a_F_timetztypmodout_1), v5)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v21 = v14
			m.G0 = v5 + int32(16)
			return v21
		}
	} else {
		v19 = F_pstrdup(m, int32(_a_F_timetztypmodout_0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v19
			m.G0 = v5 + int32(16)
			return v21
		}
	}
}
func F_tokenize_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = base.I64_rotl(v10, int64(32))
		F_errcontext_msg(m, int32(_a_F_tokenize_error_callback_0), v5)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_tolower(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
		v8 = l0 | int32(32)
	} else {
		v8 = l0
	}
	return v8
}
func F_top12(m *base.Module, l0 float64) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
}
func F_touched_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v33 int64
	_ = v33
	var v36 float64
	_ = v36
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 float64
	_ = v40
	var v43 int64
	_ = v43
	var v50 float64
	_ = v50
	var v57 int64
	_ = v57
	var v62 float64
	_ = v62
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v84 int32
	_ = v84
	var v90 float64
	_ = v90
	var v95 int64
	_ = v95
	var v96 float64
	_ = v96
	var v99 int64
	_ = v99
	var v114 int32
	_ = v114
	var v115 float64
	_ = v115
	var v118 int32
	_ = v118
	var v119 float64
	_ = v119
	var v120 int32
	_ = v120
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v154 float64
	_ = v154
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v173 float64
	_ = v173
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v203 float64
	_ = v203
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v209 float64
	_ = v209
	var v212 int64
	_ = v212
	var v229 float64
	_ = v229
	var v230 int32
	_ = v230
	var v231 float64
	_ = v231
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v250 float64
	_ = v250
	var v251 int32
	_ = v251
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v254 float64
	_ = v254
	var v255 float64
	_ = v255
	var v256 int32
	_ = v256
	var v262 float64
	_ = v262
	var v263 int32
	_ = v263
	var v264 float64
	_ = v264
	var v265 int32
	_ = v265
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v268 int32
	_ = v268
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v26
	v29 = v18 + int32(16)
	v30 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v33 = base.I64_reinterpret_f64(v30) & int64(9223372036854775807)
	if base.Ui64(v33) <= base.Ui64(int64(9218868437227405312)) {
		goto L18
	} else {
		goto L19
	}
L1:
	;
	m.G0 = v18 + int32(32)
	return v307
L2:
	;
	v295 = F_lseg_inside_poly(m, l1, v286, l3, l4)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L41
	} else {
		goto L85
	}
L3:
	;
	v307 = int32(1)
	goto L1
L4:
	;
	v250 = F_point_dt(m, l2, v18)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L41
	} else {
		goto L73
	}
L5:
	;
	v229 = F_point_dt(m, l2, v18)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L41
	} else {
		goto L68
	}
L6:
	;
	v209 = *(*float64)(unsafe.Add(mBase, uint32(v206)+8))
	v212 = base.I64_reinterpret_f64(v209) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v208) {
		goto L62
	} else {
		goto L63
	}
L7:
	;
	if base.F64_ne(v30, v194) != 0 {
		v246 = v196
		goto L4
	} else {
		goto L60
	}
L8:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v163)&int64(9223372036854775807)) {
		v246 = v165
		goto L4
	} else {
		goto L50
	}
L9:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v151)))
	if base.Ui64(v153) < base.Ui64(int64(9218868437227405313)) {
		v162 = v148
		v163 = v154
		v165 = v151
		v167 = v153
		goto L8
	} else {
		goto L49
	}
L10:
	;
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	if base.Ui64(v139&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v246 = v136
		goto L4
	} else {
		goto L48
	}
L11:
	;
	v130 = l2 + int32(16)
	if base.Ui64(v33) <= base.Ui64(int64(9218868437227405312)) {
		v148 = v90
		v151 = v130
		v153 = v95
		goto L9
	} else {
		goto L47
	}
L12:
	;
	v114 = l2 + int32(16)
	v115 = F_point_dt(m, v114, v18)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L41
	} else {
		goto L42
	}
L13:
	;
	v96 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v99 = base.I64_reinterpret_f64(v96) & int64(9223372036854775807)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v95) {
		goto L35
	} else {
		goto L36
	}
L14:
	;
	v148 = v40
	v151 = l2 + int32(16)
	v153 = v43
	goto L9
L15:
	;
	if base.F64_ne(v30, v36) != 0 {
		goto L14
	} else {
		goto L33
	}
L16:
	;
	v136 = l2 + int32(16)
	goto L10
L17:
	;
	if base.F64_ne(v30, v36) != 0 {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v38 = int64(9223372036854775807)
	v39 = base.I64_reinterpret_f64(v36) & v38
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v43 = base.I64_reinterpret_f64(v40) & v38
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v43) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui64(v57&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L16
	} else {
		goto L26
	}
L21:
	;
	v84 = base.B2i32(base.Ui64(v39) < base.Ui64(int64(9218868437227405313)))
	goto L15
L22:
	;
	goto L23
L23:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v39) {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui64(base.I64_reinterpret_f64(v50)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v84 = int32(1)
	goto L15
L26:
	;
	v62 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v90 = v62
	v95 = base.I64_reinterpret_f64(v62) & int64(9223372036854775807)
	goto L13
L27:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v30, v36)), float64(1e-06)) == int32(0) {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.F64_eq(v40, v50) != 0 {
		goto L12
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v40, v50)), float64(1e-06)) != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	v79 = l2 + int32(16)
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v79)))
	v162 = v40
	v163 = v80
	v165 = v79
	v167 = v43
	goto L8
L33:
	;
	if v84 != 0 {
		v90 = v40
		v95 = v43
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v99) {
		goto L12
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if base.F64_ne(v90, v96) != 0 {
		goto L11
	} else {
		goto L39
	}
L38:
	;
	goto L11
L39:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v99) {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	goto L12
L41:
	;
	return int32(0)
L42:
	;
	v119 = F_point_dt(m, v114, v29)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v121 = base.F64_add(v115, v119)
	v122 = F_point_dt(m, v18, v29)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	if base.F64_eq(v121, v122) != 0 {
		v286 = v114
		goto L2
	} else {
		goto L45
	}
L45:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v121, v122)), float64(1e-06)) != 0 {
		v286 = v114
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L3
L47:
	;
	v136 = v130
	goto L10
L48:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v203 = v144
	v206 = v136
	v208 = base.I64_reinterpret_f64(v144) & int64(9223372036854775807)
	goto L6
L49:
	;
	v192 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v154)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	v193 = v148
	v194 = v154
	v196 = v151
	v198 = v153
	goto L7
L50:
	;
	v173 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v173)&int64(9223372036854775807)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v192 = int32(1)
	v193 = v162
	v194 = v163
	v196 = v165
	v198 = v167
	goto L7
L52:
	;
	goto L53
L53:
	;
	if base.F64_ne(v30, v163) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v30, v163)), float64(1e-06)) == int32(0) {
		v246 = v165
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if base.F64_eq(v162, v173) != 0 {
		goto L5
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v162, v173)), float64(1e-06)) != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	v246 = v165
	goto L4
L60:
	;
	if v192 == int32(0) {
		v246 = v196
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v203 = v193
	v206 = v196
	v208 = v198
	goto L6
L62:
	;
	if base.Ui64(v212) <= base.Ui64(int64(9218868437227405312)) {
		v246 = v206
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if base.F64_ne(v209, v203) != 0 {
		v246 = v206
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L5
L66:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v212) {
		v246 = v206
		goto L4
	} else {
		goto L67
	}
L67:
	;
	goto L5
L68:
	;
	v231 = F_point_dt(m, l2, v29)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L41
	} else {
		goto L69
	}
L69:
	;
	v233 = base.F64_add(v229, v231)
	v234 = F_point_dt(m, v18, v29)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L41
	} else {
		goto L70
	}
L70:
	;
	if base.F64_eq(v233, v234) != 0 {
		v286 = l2
		goto L2
	} else {
		goto L71
	}
L71:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v233, v234)), float64(1e-06)) != 0 {
		v286 = l2
		goto L2
	} else {
		goto L72
	}
L72:
	;
	goto L3
L73:
	;
	v252 = F_point_dt(m, l2, v29)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L41
	} else {
		goto L74
	}
L74:
	;
	v254 = base.F64_add(v250, v252)
	v255 = F_point_dt(m, v18, v29)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L41
	} else {
		goto L75
	}
L75:
	;
	if base.F64_eq(v254, v255) != 0 {
		v286 = l2
		goto L2
	} else {
		goto L76
	}
L76:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v254, v255)), float64(1e-06)) != 0 {
		v286 = l2
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v262 = F_point_dt(m, v246, v18)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L41
	} else {
		goto L78
	}
L78:
	;
	v264 = F_point_dt(m, v246, v29)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L41
	} else {
		goto L79
	}
L79:
	;
	v266 = base.F64_add(v262, v264)
	v267 = F_point_dt(m, v18, v29)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L41
	} else {
		goto L80
	}
L80:
	;
	if base.F64_eq(v266, v267) != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v286 = v246
	goto L2
L82:
	;
	goto L83
L83:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v266, v267)), float64(1e-06)) != 0 {
		v286 = v246
		goto L2
	} else {
		goto L84
	}
L84:
	;
	goto L3
L85:
	;
	v307 = v295
	goto L1
}
func F_tqueueDestroyReceiver(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_shm_mq_detach(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_tqueueReceiveSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_ExecFetchSlotMinimalTuple(m, l0, v6+int32(15))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v16 = int32(0)
		v18 = F_shm_mq_send(m, v14, v15, v10, v16, v16)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
			if v20 == int32(1) {
				F_pfree(m, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					switch v18 {
					case 0, 2:
						m.G0 = v6 + int32(16)
						return base.B2i32(v18 != int32(2))
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_tqueueReceiveSlot_0), int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_tqueueReceiveSlot_1), int32(74), int32(_a_F_tqueueReceiveSlot_2))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
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
				}
			} else {
				switch v18 {
				case 0, 2:
					m.G0 = v6 + int32(16)
					return base.B2i32(v18 != int32(2))
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_tqueueReceiveSlot_0), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_tqueueReceiveSlot_1), int32(74), int32(_a_F_tqueueReceiveSlot_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
			}
		}
	}
}
func F_trackitem_compare_frequencies_desc_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	return v5 - v7
}
func F_transformDistinctOnClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v463 int32
	_ = v463
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	v5 = int32(0)
	if l1 == v5 {
		v237 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 <= int32(0) {
		v237 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = v5
	v38 = v5
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v38<<(uint(int32(2))%32))))
	v48 = F_findTargetlistEntrySQL92(m, l0, v46, l2, int32(21))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v237 = v222
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v52 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v56 == int32(0) {
		v202 = int32(1)
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v208 = v52
	goto L10
L10:
	;
	v222 = F_lappend_int(m, v33, v208)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L41
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v202
	v208 = v202
	goto L10
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v60 <= int32(0) {
		v202 = int32(1)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v64 = v60 & int32(3)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v66 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v60) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v78 = v66
	v80 = v66
	v82 = int32(0)
	goto L17
L15:
	;
	v120 = v66
	v122 = v66
	goto L16
L16:
	;
	if v64 != 0 {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v94 = v65 + v80<<(uint(int32(2))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if base.Ui32(v78) < base.Ui32(v102) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v120 = v110
	v122 = v112
	goto L16
L19:
	;
	v104 = v102
	goto L21
L20:
	;
	v104 = v78
	goto L21
L21:
	;
	if base.Ui32(v104) < base.Ui32(v100) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v106 = v100
	goto L24
L23:
	;
	v106 = v104
	goto L24
L24:
	;
	if base.Ui32(v106) < base.Ui32(v98) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v108 = v98
	goto L27
L26:
	;
	v108 = v106
	goto L27
L27:
	;
	if base.Ui32(v108) < base.Ui32(v96) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v110 = v96
	goto L30
L29:
	;
	v110 = v108
	goto L30
L30:
	;
	v111 = int32(4)
	v112 = v80 + v111
	v114 = v82 + v111
	if v114 != v60&int32(2147483644) {
		v78 = v110
		v80 = v112
		v82 = v114
		goto L17
	} else {
		goto L31
	}
L31:
	;
	goto L18
L32:
	;
	v138 = v120
	v140 = v122
	v145 = v66
	goto L35
L33:
	;
	v168 = v120
	goto L34
L34:
	;
	v202 = v168 + int32(1)
	goto L11
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v65+v140<<(uint(int32(2))%32))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	if base.Ui32(v138) < base.Ui32(v156) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v168 = v158
	goto L34
L37:
	;
	v158 = v156
	goto L39
L38:
	;
	v158 = v138
	goto L39
L39:
	;
	v159 = int32(1)
	v162 = v145 + v159
	if v162 != v64 {
		v138 = v158
		v140 = v140 + v159
		v145 = v162
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v225 = v38 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v225 < v226 {
		v33 = v222
		v38 = v225
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L5
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L6
	} else {
		goto L129
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L99
	}
L45:
	;
	v358 = v338
	v360 = int32(0)
	goto L74
L46:
	;
	v338 = int32(0)
	v347 = v5
	goto L45
L47:
	;
	goto L48
L48:
	;
	v249 = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v250 <= v249 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v338 = int32(0)
	v347 = v5
	goto L45
L50:
	;
	goto L51
L51:
	;
	v259 = v249
	v262 = int32(0)
	v271 = v5
	goto L52
L52:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273+v259<<(uint(int32(2))%32))))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v279 = int32(0)
	if v237 == v279 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v338 = v324
	v347 = v326
	goto L45
L54:
	;
	if v317 != 0 {
		goto L67
	} else {
		goto L68
	}
L55:
	;
	v317 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v285 <= int32(0) {
		v310 = v279
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v317 = v310
	goto L54
L59:
	;
	v288 = int32(0)
	if v288 < v285 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v291 = v285
	goto L62
L61:
	;
	v291 = v288
	goto L62
L62:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v294 = int32(0)
	goto L63
L63:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v292+v294<<(uint(int32(2))%32))))
	v303 = base.B2i32(v302 == v278)
	if v302 == v278 {
		v310 = v303
		goto L58
	} else {
		goto L65
	}
L64:
	;
	v310 = v303
	goto L58
L65:
	;
	v305 = v294 + int32(1)
	if v305 != v291 {
		v294 = v305
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	if v271&int32(1) != 0 {
		goto L44
	} else {
		goto L70
	}
L68:
	;
	v324 = v262
	goto L69
L69:
	;
	v325 = int32(1)
	v326 = v317 ^ v325
	v328 = v259 + v325
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v328 < v329 {
		v259 = v328
		v262 = v324
		v271 = v326
		goto L52
	} else {
		goto L73
	}
L70:
	;
	v320 = F_copyObjectImpl(m, v277)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v322 = F_lappend(m, v262, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v324 = v322
	goto L69
L73:
	;
	goto L53
L74:
	;
	v368 = int32(0)
	if l1 == v368 {
		v378 = v368
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v237 != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v372 <= v360 {
		v378 = int32(0)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v378 = v374 + v360<<(uint(int32(2))%32)
	goto L76
L79:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v394 = F_get_sortgroupref_tle(m, v392, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L6
	} else {
		goto L89
	}
L80:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v379 <= v360 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v389 = v338
	goto L82
L82:
	;
	return v389
L83:
	;
	v389 = v358
	goto L82
L84:
	;
	if v378 == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v386 = v383 + v360<<(uint(int32(2))%32)
	if v386 != 0 {
		goto L79
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	v358 = v463
	v360 = v360 + int32(1)
	goto L74
L88:
	;
	if v347 != 0 {
		goto L43
	} else {
		goto L97
	}
L89:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	if v396 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	if v358 == int32(0) {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v401 <= int32(0) {
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v410 = int32(0)
	goto L93
L93:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v404+v410<<(uint(int32(2))%32))))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v428 == v396 {
		v463 = v358
		goto L87
	} else {
		goto L95
	}
L94:
	;
	goto L88
L95:
	;
	v431 = v410 + int32(1)
	if v431 != v401 {
		v410 = v431
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v452 = F_exprLocation(m, v391)
	mBase = m.M
	v453 = F_addTargetToGroupList(m, l0, v394, v358, v451, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v463 = v453
	goto L87
L99:
	;
	F_errcode(m, int32(_a_F_transformDistinctOnClause_0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_transformDistinctOnClause_1), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v487 = int32(0)
	if v237 == v487 {
		v494 = v487
		goto L103
	} else {
		goto L104
	}
L102:
	;
	F_parser_errposition(m, l0, v567)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L6
	} else {
		goto L127
	}
L103:
	;
	if l1 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v490 <= int32(0) {
		v494 = v487
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v494 = v493
	goto L103
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L6
	} else {
		goto L124
	}
L107:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v497 <= int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	if v494 == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v502 == int32(0) {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v486 == v505 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	v567 = F_exprLocation(m, v566)
	mBase = m.M
	goto L102
L112:
	;
	v550 = v502
	goto L111
L113:
	;
	goto L114
L114:
	;
	v509 = int32(1)
	goto L115
L115:
	;
	v526 = int32(0)
	if v237 == v526 {
		v535 = v526
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v550 = v541
	goto L111
L117:
	;
	if v497 <= v509 {
		goto L106
	} else {
		goto L120
	}
L118:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v529 <= v509 {
		v535 = v526
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v535 = v531 + v509<<(uint(int32(2))%32)
	goto L117
L120:
	;
	if v535 == int32(0) {
		goto L106
	} else {
		goto L121
	}
L121:
	;
	v541 = v502 + v509<<(uint(int32(2))%32)
	if v541 == int32(0) {
		goto L106
	} else {
		goto L122
	}
L122:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	if v546 != v486 {
		v509 = v509 + int32(1)
		goto L115
	} else {
		goto L123
	}
L123:
	;
	goto L116
L124:
	;
	F_errmsg_internal(m, int32(_a_F_transformDistinctOnClause_2), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_transformDistinctOnClause_3), int32(3187), int32(_a_F_transformDistinctOnClause_4))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errfinish(m, int32(_a_F_transformDistinctOnClause_3), int32(3120), int32(_a_F_transformDistinctOnClause_5))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(_a_F_transformDistinctOnClause_0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_transformDistinctOnClause_1), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	v617 = F_exprLocation(m, v391)
	mBase = m.M
	F_parser_errposition(m, l0, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_transformDistinctOnClause_3), int32(3149), int32(_a_F_transformDistinctOnClause_5))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_treekey_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v99 int32
	_ = v99
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
	if v13 == int32(0) {
		v77 = v13
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v99
L2:
	;
	v99 = (v77 + int32(1)) * (v13 - v12) * int32(10)
	goto L1
L3:
	;
	if v12 == int32(0) {
		v77 = v13
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(8)
	v22 = v3 + v18
	v23 = v4 + v18
	v26 = v13
	v30 = v12
	goto L5
L5:
	;
	v31 = int32(2)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if base.Ui32(v35) < base.Ui32(v36) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v77 = v58
	goto L2
L7:
	;
	v58 = v26 - int32(1)
	if v26 < int32(2) {
		v77 = v58
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v38 = v35
	goto L10
L9:
	;
	v38 = v36
	goto L10
L10:
	;
	v39 = F_memcmp(m, v22+v31, v23+v31, v38)
	mBase = m.M
	if v39 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v35 == v36 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v39 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v43 = int32(10)
	v99 = (v26*v43 + v43) * (v35 - v36)
	goto L1
L15:
	;
	v55 = int32(-10)
	goto L17
L16:
	;
	v55 = int32(10)
	goto L17
L17:
	;
	v99 = (v26 + int32(1)) * v55
	goto L1
L18:
	;
	v61 = int32(9)
	v63 = int32(_a_F_treekey_cmp_0)
	v71 = int32(1)
	if v71 < v30 {
		v22 = v22 + (v35+v61)&v63
		v23 = v23 + (v36+v61)&v63
		v26 = v58
		v30 = v30 - v71
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L6
}
func F_trigramsMatchGraph(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v197 int32
	_ = v197
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = F__emscripten_memset_bulkmem(m, v13, base.I32_extend8_s(v3), v15)
	mBase = m.M
	goto L1
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = F__emscripten_memset_bulkmem(m, v18, base.I32_extend8_s(int32(0)), v20)
	mBase = m.M
	goto L2
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = v3
	v32 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v97)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = v97
	v110 = v3
	goto L18
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26+v32<<(uint(int32(2))%32))))
	v43 = v42 + v30
	if v42 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v82 = v32 + int32(1)
	if v82 != v23 {
		v30 = v43
		v32 = v82
		goto L6
	} else {
		goto L16
	}
L9:
	;
	v48 = v30
	goto L10
L10:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v48))))
	if v59 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v65+v32))) = uint8(v67)
	goto L8
L12:
	;
	v63 = v48 + int32(1)
	if v63 < v43 {
		v48 = v63
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	goto L8
L16:
	;
	goto L7
L17:
	;
	return v197
L18:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99+v110<<(uint(int32(2))%32))))
	v122 = v102 + v119<<(uint(int32(3))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if int32(0) < v123 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v197 = int32(0)
	goto L17
L20:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v131 = int32(0)
	v133 = v108
	goto L23
L21:
	;
	v173 = v108
	goto L22
L22:
	;
	v182 = v110 + int32(1)
	if v182 < v173 {
		v108 = v173
		v110 = v182
		goto L18
	} else {
		goto L30
	}
L23:
	;
	v143 = v127 + v131<<(uint(int32(3))%32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v144))))
	if v146 != int32(1) {
		v164 = v133
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v173 = v164
	goto L22
L25:
	;
	v167 = v131 + int32(1)
	if v167 != v123 {
		v131 = v167
		v133 = v164
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v149 = int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v150 == v149 {
		v197 = v149
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v153 = v150 + v96
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v154 != 0 {
		v164 = v133
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v155)
	*(*int32)(unsafe.Add(mBase, uint32(v99+v133<<(uint(int32(2))%32)))) = v150
	v164 = v133 + v155
	goto L25
L29:
	;
	goto L24
L30:
	;
	goto L19
}
func F_truncate_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v15 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_truncate_cb_wrapper_0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(992)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v19
		v23 = int32(_a_F_truncate_cb_wrapper_1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_cb_wrapper[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_truncate_cb_wrapper[0])) = v12 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(16)
		v33 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+147)) = uint8(v33)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v35
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		v38 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+164)) = uint8(v38)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v37
		m.T0[v15].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14, l1, l2, l3, l4)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_truncate_cb_wrapper[0])) = v44
			m.G0 = v12 + int32(32)
			return
		}
	} else {
		m.G0 = v12 + int32(32)
		return
	}
}
func F_try_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v118 float64
	_ = v118
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 float64
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 float64
	_ = v172
	var v174 float64
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v188 float64
	_ = v188
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 float64
	_ = v259
	var v266 float64
	_ = v266
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(96)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v31 = F_calc_non_nestloop_required_outer(m, l2, l3)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L17
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L7
L6:
	;
	v19 = int32(0)
	goto L7
L7:
	;
	v20 = F_bms_is_member(m, v15, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v24 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = v25
	goto L13
L12:
	;
	v27 = int32(0)
	goto L13
L13:
	;
	v28 = F_bms_is_member(m, v23, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L4
L16:
	;
	F_bms_free(m, v31)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L91
	}
L17:
	;
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v34 = int32(0)
	if v31 == v34 {
		v75 = v34
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v93 = m.G0
	v95 = v93 - int32(16)
	m.G0 = v95
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v100 = *(*float64)(unsafe.Add(mBase, _c_F_try_hashjoin_path[0]))
	if l4 != 0 {
		goto L37
	} else {
		goto L38
	}
L21:
	;
	if v75 == int32(0) {
		goto L16
	} else {
		goto L35
	}
L22:
	;
	goto L21
L23:
	;
	if v33 == int32(0) {
		v75 = v34
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v43 < v44 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v46 = v43
	goto L27
L26:
	;
	v46 = v44
	goto L27
L27:
	;
	if v46 <= int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v49 = int32(1)
	goto L30
L29:
	;
	v49 = v46
	goto L30
L30:
	;
	v50 = int32(8)
	v55 = int32(0)
	goto L31
L31:
	;
	v62 = v55 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v33+v50+v62)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+(v31+v50))))
	v67 = v64 & v66
	v69 = base.B2i32(v67 != int32(0))
	if v67 != 0 {
		v75 = v69
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v75 = v69
	goto L22
L33:
	;
	v71 = v55 + int32(1)
	if v71 != v49 {
		v55 = v71
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L20
L36:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	v219 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v220 = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v225 == v220 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v104 = base.F64_convert_i32_s(v101)
	goto L39
L38:
	;
	v104 = float64(0)
	goto L39
L39:
	;
	v105 = base.F64_mul(v100, v104)
	v107 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v108 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v110 = float64(0)
	v113 = *(*float64)(unsafe.Add(mBase, _c_F_try_hashjoin_path[1]))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	goto L41
L41:
	;
	goto L42
L42:
	;
	v140 = base.F64_add(base.F64_mul(v105, v98), base.F64_add(base.F64_sub(v107, v108), v110))
	v141 = base.F64_add(base.F64_mul(base.F64_add(v105, v113), v97), base.F64_add(base.F64_add(v108, v110), v118))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_hashjoin_path[2])))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_ExecChooseHashTableSize(m, v97, v147, int32(1), int32(0), v149, v95, v95+int32(12), v95+int32(8), v95+int32(4))
	mBase = m.M
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	if int32(2) <= v157 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v161 = *(*float64)(unsafe.Add(mBase, _c_F_try_hashjoin_path[3]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+32))
	v164 = int32(7)
	v166 = int32(-8)
	v168 = int32(24)
	v172 = float64(0.0001220703125)
	v174 = base.F64_ceil(base.F64_mul(base.F64_mul(v98, base.F64_convert_i32_u((v163+v164)&v166+v168)), v172))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+32))
	v188 = base.F64_ceil(base.F64_mul(base.F64_mul(v97, base.F64_convert_i32_u((v177+v164)&v166+v168)), v172))
	v194 = base.F64_add(base.F64_mul(v161, base.F64_add(base.F64_add(v174, v174), v188)), v140)
	v195 = base.F64_add(base.F64_mul(v161, v188), v141)
	goto L51
L50:
	;
	v194 = v140
	v195 = v141
	goto L51
L51:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v194
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v195
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = base.F64_add(v194, v195)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v145 + (v143^int32(1))&int32(255) + v144
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v210
	m.G0 = v95 + int32(16)
	goto L36
L52:
	;
	if v303 == int32(0) {
		goto L16
	} else {
		goto L88
	}
L53:
	;
	v303 = int32(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v229 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v303 = int32(1)
	goto L52
L57:
	;
	goto L58
L58:
	;
	if v31 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v234 = int32(0)
	goto L61
L60:
	;
	v234 = v220
	goto L61
L61:
	;
	if v31 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v237 = int32(25)
	goto L64
L63:
	;
	v237 = int32(24)
	goto L64
L64:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v237))))
	v247 = v220
	goto L65
L65:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v247<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+40))
	if v217 != v255 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v303 = v291
	goto L52
L67:
	;
	if v239 != 0 {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	if v255 <= v217 {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v259 = *(*float64)(unsafe.Add(mBase, uint32(v254)+56))
	if base.F64_le(v219, base.F64_mul(v259, float64(1.01))) == int32(0) {
		goto L67
	} else {
		goto L72
	}
L71:
	;
	v303 = int32(1)
	goto L52
L72:
	;
	v303 = int32(1)
	goto L52
L73:
	;
	goto L66
L74:
	;
	v285 = int32(1)
	v287 = v247 + v285
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v287 < v288 {
		v247 = v287
		goto L65
	} else {
		goto L87
	}
L75:
	;
	v266 = *(*float64)(unsafe.Add(mBase, uint32(v254)+48))
	if base.F64_gt(v218, base.F64_mul(v266, float64(1.01))) == int32(0) {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v272 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if v273 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	v275 = v272
	goto L81
L80:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v254)+64))
	v275 = v274
	goto L81
L81:
	;
	v276 = F_compare_pathkeys(m, v234, v275)
	mBase = m.M
	if v276&int32(-3) != 0 {
		goto L74
	} else {
		goto L82
	}
L82:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	if v279 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v281 = v280
	goto L85
L84:
	;
	v281 = v272
	goto L85
L85:
	;
	v282 = F_bms_equal(m, v31, v281)
	mBase = m.M
	if v282 != 0 {
		v291 = v272
		goto L73
	} else {
		goto L86
	}
L86:
	;
	goto L74
L87:
	;
	v291 = v285
	goto L73
L88:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v308 = F_create_hashjoin_path(m, l0, l1, l5, v12, l6, l2, l3, int32(0), v307, v31, l4)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_add_path(m, l1, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	goto L1
L91:
	;
	goto L1
}
func F_tsm_bernoulli_handler(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
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
			v24 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(278)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(279)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(280)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(281)
			v36 = int32(257)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v36)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22
			m.G0 = v6 + int32(16)
			return v9
		}
	}
}
func F_tsmatchjoinsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.005))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_tsmatchsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 float64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 float32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 float32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 float32
	_ = v180
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 float32
	_ = v208
	var v209 float64
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v226 float64
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 float64
	_ = v241
	var v242 int32
	_ = v242
	var v256 float64
	_ = v256
	var v259 float32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 float64
	_ = v271
	var v272 int32
	_ = v272
	var v287 float64
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 float64
	_ = v292
	var v313 float64
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = float64(0.005)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v30 = F_get_restriction_variable(m, v21, v22, v23, v18+int32(12), v18+int32(8), v18+int32(7))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		if v30 == int32(0) {
			v313 = v20
			v315 = F_Float8GetDatum(m, v313)
			mBase = m.M
			v316 = m.ExcPending
			if v316 != 0 {
				return int32(0)
			} else {
				m.G0 = v18 + int32(80)
				return v315
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v37 != int32(7) {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
				if v40 == int32(0) {
					v313 = v20
					v315 = F_Float8GetDatum(m, v313)
					mBase = m.M
					v316 = m.ExcPending
					if v316 != 0 {
						return int32(0)
					} else {
						m.G0 = v18 + int32(80)
						return v315
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
					m.T0[v43].(func(*base.Module, int32))(m, v40)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v313 = v20
						v315 = F_Float8GetDatum(m, v313)
						mBase = m.M
						v316 = m.ExcPending
						if v316 != 0 {
							return int32(0)
						} else {
							m.G0 = v18 + int32(80)
							return v315
						}
					}
				}
			} else {
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
				if v46 == int32(1) {
					v49 = float64(0)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
					if v50 == int32(0) {
						v313 = v49
						v315 = F_Float8GetDatum(m, v313)
						mBase = m.M
						v316 = m.ExcPending
						if v316 != 0 {
							return int32(0)
						} else {
							m.G0 = v18 + int32(80)
							return v315
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
						m.T0[v53].(func(*base.Module, int32))(m, v50)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v313 = v49
							v315 = F_Float8GetDatum(m, v313)
							mBase = m.M
							v316 = m.ExcPending
							if v316 != 0 {
								return int32(0)
							} else {
								m.G0 = v18 + int32(80)
								return v315
							}
						}
					}
				} else {
					v56 = float64(0.005)
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
					if v57 != int32(3615) {
						v287 = v56
						v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
						if v288 != 0 {
							v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
							m.T0[v289].(func(*base.Module, int32))(m, v288)
							mBase = m.M
							v291 = m.ExcPending
							if v291 != 0 {
								return int32(0)
							} else {
								v292 = float64(0)
								if base.F64_lt(v287, v292) != 0 {
									v313 = v292
								} else {
									if base.F64_gt(v287, float64(1)) == int32(0) {
										v313 = v287
									} else {
										v313 = float64(1)
									}
								}
								v315 = F_Float8GetDatum(m, v313)
								mBase = m.M
								v316 = m.ExcPending
								if v316 != 0 {
									return int32(0)
								} else {
									m.G0 = v18 + int32(80)
									return v315
								}
							}
						} else {
							v292 = float64(0)
							if base.F64_lt(v287, v292) != 0 {
								v313 = v292
							} else {
								if base.F64_gt(v287, float64(1)) == int32(0) {
									v313 = v287
								} else {
									v313 = float64(1)
								}
							}
							v315 = F_Float8GetDatum(m, v313)
							mBase = m.M
							v316 = m.ExcPending
							if v316 != 0 {
								return int32(0)
							} else {
								m.G0 = v18 + int32(80)
								return v315
							}
						}
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
						if v60 != int32(3614) {
							v287 = v56
							v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
							if v288 != 0 {
								v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
								m.T0[v289].(func(*base.Module, int32))(m, v288)
								mBase = m.M
								v291 = m.ExcPending
								if v291 != 0 {
									return int32(0)
								} else {
									v292 = float64(0)
									if base.F64_lt(v287, v292) != 0 {
										v313 = v292
									} else {
										if base.F64_gt(v287, float64(1)) == int32(0) {
											v313 = v287
										} else {
											v313 = float64(1)
										}
									}
									v315 = F_Float8GetDatum(m, v313)
									mBase = m.M
									v316 = m.ExcPending
									if v316 != 0 {
										return int32(0)
									} else {
										m.G0 = v18 + int32(80)
										return v315
									}
								}
							} else {
								v292 = float64(0)
								if base.F64_lt(v287, v292) != 0 {
									v313 = v292
								} else {
									if base.F64_gt(v287, float64(1)) == int32(0) {
										v313 = v287
									} else {
										v313 = float64(1)
									}
								}
								v315 = F_Float8GetDatum(m, v313)
								mBase = m.M
								v316 = m.ExcPending
								if v316 != 0 {
									return int32(0)
								} else {
									m.G0 = v18 + int32(80)
									return v315
								}
							}
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
							if v64 == int32(0) {
								v287 = float64(0)
								v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
								if v288 != 0 {
									v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
									m.T0[v289].(func(*base.Module, int32))(m, v288)
									mBase = m.M
									v291 = m.ExcPending
									if v291 != 0 {
										return int32(0)
									} else {
										v292 = float64(0)
										if base.F64_lt(v287, v292) != 0 {
											v313 = v292
										} else {
											if base.F64_gt(v287, float64(1)) == int32(0) {
												v313 = v287
											} else {
												v313 = float64(1)
											}
										}
										v315 = F_Float8GetDatum(m, v313)
										mBase = m.M
										v316 = m.ExcPending
										if v316 != 0 {
											return int32(0)
										} else {
											m.G0 = v18 + int32(80)
											return v315
										}
									}
								} else {
									v292 = float64(0)
									if base.F64_lt(v287, v292) != 0 {
										v313 = v292
									} else {
										if base.F64_gt(v287, float64(1)) == int32(0) {
											v313 = v287
										} else {
											v313 = float64(1)
										}
									}
									v315 = F_Float8GetDatum(m, v313)
									mBase = m.M
									v316 = m.ExcPending
									if v316 != 0 {
										return int32(0)
									} else {
										m.G0 = v18 + int32(80)
										return v315
									}
								}
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
								if v68 != 0 {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
									v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
									v77 = F_get_attstatsslot(m, v18+int32(44), v68, int32(4), int32(0), int32(3))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										if v77 != 0 {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
											if v79 != v80+int32(2) {
												v85 = v63 + int32(8)
												v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
												v90 = int32(0)
												v93 = F_tsquery_opr_selec(m, v85, v85+v86*int32(12), v90, v90, float32(0))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													v226 = v93
													F_free_attstatsslot(m, v18+int32(44))
													mBase = m.M
													v231 = m.ExcPending
													if v231 != 0 {
														return int32(0)
													} else {
														v256 = v226
														v259 = *(*float32)(unsafe.Add(mBase, uint32(v69+v70)+8))
														v287 = base.F64_mul(v256, base.F64_sub(float64(1), base.F64_promote_f32(v259)))
														v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
														if v288 != 0 {
															v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
															m.T0[v289].(func(*base.Module, int32))(m, v288)
															mBase = m.M
															v291 = m.ExcPending
															if v291 != 0 {
																return int32(0)
															} else {
																v292 = float64(0)
																if base.F64_lt(v287, v292) != 0 {
																	v313 = v292
																} else {
																	if base.F64_gt(v287, float64(1)) == int32(0) {
																		v313 = v287
																	} else {
																		v313 = float64(1)
																	}
																}
																v315 = F_Float8GetDatum(m, v313)
																mBase = m.M
																v316 = m.ExcPending
																if v316 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(80)
																	return v315
																}
															}
														} else {
															v292 = float64(0)
															if base.F64_lt(v287, v292) != 0 {
																v313 = v292
															} else {
																if base.F64_gt(v287, float64(1)) == int32(0) {
																	v313 = v287
																} else {
																	v313 = float64(1)
																}
															}
															v315 = F_Float8GetDatum(m, v313)
															mBase = m.M
															v316 = m.ExcPending
															if v316 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(80)
																return v315
															}
														}
													}
												}
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
												v99 = F_palloc(m, v80<<(uint(int32(3))%32))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													if v80 <= int32(0) {
													} else {
														v103 = int32(0)
														if v80 != int32(1) {
															v108 = v103
															v115 = int32(0)
															for {
																v123 = int32(3)
																v125 = v99 + v108<<(uint(v123)%32)
																v126 = int32(2)
																v127 = v108 << (uint(v126) % 32)
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v96+v127)))
																*(*int32)(unsafe.Add(mBase, uint32(v125))) = v129
																v132 = *(*float32)(unsafe.Add(mBase, uint32(v95+v127)))
																*(*float32)(unsafe.Add(mBase, uint32(v125)+4)) = v132
																v135 = v108 | int32(1)
																v138 = v99 + v135<<(uint(v123)%32)
																v140 = v135 << (uint(v126) % 32)
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v96+v140)))
																*(*int32)(unsafe.Add(mBase, uint32(v138))) = v142
																v145 = *(*float32)(unsafe.Add(mBase, uint32(v140+v95)))
																*(*float32)(unsafe.Add(mBase, uint32(v138)+4)) = v145
																v148 = v108 + v126
																v150 = v115 + v126
																if v150 != v80&int32(2147483646) {
																	v108 = v148
																	v115 = v150
																	continue
																} else {
																	break
																}
																break
															}
															v152 = v148
														} else {
															v152 = v103
														}
														if v80&int32(1) == int32(0) {
														} else {
															v173 = v99 + v152<<(uint(int32(3))%32)
															v175 = v152 << (uint(int32(2)) % 32)
															v177 = *(*int32)(unsafe.Add(mBase, uint32(v96+v175)))
															*(*int32)(unsafe.Add(mBase, uint32(v173))) = v177
															v180 = *(*float32)(unsafe.Add(mBase, uint32(v175+v95)))
															*(*float32)(unsafe.Add(mBase, uint32(v173)+4)) = v180
														}
													}
													v197 = int32(8)
													v198 = v63 + v197
													v199 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
													v208 = *(*float32)(unsafe.Add(mBase, uint32(v95+v79<<(uint(int32(2))%32)-v197)))
													v209 = F_tsquery_opr_selec(m, v198, v198+v199*int32(12), v99, v80, v208)
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v99)
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return int32(0)
														} else {
															v226 = v209
															F_free_attstatsslot(m, v18+int32(44))
															mBase = m.M
															v231 = m.ExcPending
															if v231 != 0 {
																return int32(0)
															} else {
																v256 = v226
																v259 = *(*float32)(unsafe.Add(mBase, uint32(v69+v70)+8))
																v287 = base.F64_mul(v256, base.F64_sub(float64(1), base.F64_promote_f32(v259)))
																v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
																if v288 != 0 {
																	v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
																	m.T0[v289].(func(*base.Module, int32))(m, v288)
																	mBase = m.M
																	v291 = m.ExcPending
																	if v291 != 0 {
																		return int32(0)
																	} else {
																		v292 = float64(0)
																		if base.F64_lt(v287, v292) != 0 {
																			v313 = v292
																		} else {
																			if base.F64_gt(v287, float64(1)) == int32(0) {
																				v313 = v287
																			} else {
																				v313 = float64(1)
																			}
																		}
																		v315 = F_Float8GetDatum(m, v313)
																		mBase = m.M
																		v316 = m.ExcPending
																		if v316 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v18 + int32(80)
																			return v315
																		}
																	}
																} else {
																	v292 = float64(0)
																	if base.F64_lt(v287, v292) != 0 {
																		v313 = v292
																	} else {
																		if base.F64_gt(v287, float64(1)) == int32(0) {
																			v313 = v287
																		} else {
																			v313 = float64(1)
																		}
																	}
																	v315 = F_Float8GetDatum(m, v313)
																	mBase = m.M
																	v316 = m.ExcPending
																	if v316 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v18 + int32(80)
																		return v315
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v233 = v63 + int32(8)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
											v238 = int32(0)
											v241 = F_tsquery_opr_selec(m, v233, v233+v234*int32(12), v238, v238, float32(0))
											mBase = m.M
											v242 = m.ExcPending
											if v242 != 0 {
												return int32(0)
											} else {
												v256 = v241
												v259 = *(*float32)(unsafe.Add(mBase, uint32(v69+v70)+8))
												v287 = base.F64_mul(v256, base.F64_sub(float64(1), base.F64_promote_f32(v259)))
												v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
												if v288 != 0 {
													v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
													m.T0[v289].(func(*base.Module, int32))(m, v288)
													mBase = m.M
													v291 = m.ExcPending
													if v291 != 0 {
														return int32(0)
													} else {
														v292 = float64(0)
														if base.F64_lt(v287, v292) != 0 {
															v313 = v292
														} else {
															if base.F64_gt(v287, float64(1)) == int32(0) {
																v313 = v287
															} else {
																v313 = float64(1)
															}
														}
														v315 = F_Float8GetDatum(m, v313)
														mBase = m.M
														v316 = m.ExcPending
														if v316 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(80)
															return v315
														}
													}
												} else {
													v292 = float64(0)
													if base.F64_lt(v287, v292) != 0 {
														v313 = v292
													} else {
														if base.F64_gt(v287, float64(1)) == int32(0) {
															v313 = v287
														} else {
															v313 = float64(1)
														}
													}
													v315 = F_Float8GetDatum(m, v313)
													mBase = m.M
													v316 = m.ExcPending
													if v316 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(80)
														return v315
													}
												}
											}
										}
									}
								} else {
									v264 = v63 + int32(8)
									v268 = int32(0)
									v271 = F_tsquery_opr_selec(m, v264, v264+v64*int32(12), v268, v268, float32(0))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return int32(0)
									} else {
										v287 = v271
										v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
										if v288 != 0 {
											v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
											m.T0[v289].(func(*base.Module, int32))(m, v288)
											mBase = m.M
											v291 = m.ExcPending
											if v291 != 0 {
												return int32(0)
											} else {
												v292 = float64(0)
												if base.F64_lt(v287, v292) != 0 {
													v313 = v292
												} else {
													if base.F64_gt(v287, float64(1)) == int32(0) {
														v313 = v287
													} else {
														v313 = float64(1)
													}
												}
												v315 = F_Float8GetDatum(m, v313)
												mBase = m.M
												v316 = m.ExcPending
												if v316 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(80)
													return v315
												}
											}
										} else {
											v292 = float64(0)
											if base.F64_lt(v287, v292) != 0 {
												v313 = v292
											} else {
												if base.F64_gt(v287, float64(1)) == int32(0) {
													v313 = v287
												} else {
													v313 = float64(1)
												}
											}
											v315 = F_Float8GetDatum(m, v313)
											mBase = m.M
											v316 = m.ExcPending
											if v316 != 0 {
												return int32(0)
											} else {
												m.G0 = v18 + int32(80)
												return v315
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
func F_tsrange_subdiff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v11 = F_Float8GetDatum(m, base.F64_div(base.F64_sub(base.F64_convert_i64_s(v3), base.F64_convert_i64_s(v6)), float64(1e+06)))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_tstoreShutdownReceiver(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v8 != 0 {
				F_pfree(m, v8)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v13 != 0 {
						F_free_conversion_map(m, v13)
						mBase = m.M
						v15 = m.ExcPending
						if v15 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
							v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							if v18 != 0 {
								F_ExecDropSingleTupleTableSlot(m, v18)
								mBase = m.M
								v20 = m.ExcPending
								if v20 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v18 != 0 {
							F_ExecDropSingleTupleTableSlot(m, v18)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v13 != 0 {
					F_free_conversion_map(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v18 != 0 {
							F_ExecDropSingleTupleTableSlot(m, v18)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v18 != 0 {
						F_ExecDropSingleTupleTableSlot(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v8 != 0 {
			F_pfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v13 != 0 {
					F_free_conversion_map(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v18 != 0 {
							F_ExecDropSingleTupleTableSlot(m, v18)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v18 != 0 {
						F_ExecDropSingleTupleTableSlot(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v13 != 0 {
				F_free_conversion_map(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v18 != 0 {
						F_ExecDropSingleTupleTableSlot(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				if v18 != 0 {
					F_ExecDropSingleTupleTableSlot(m, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
					return
				}
			}
		}
	}
}
func F_tuplehash_start_iterate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = int32(-1)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int64(0) {
		v29 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(0)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10+v12*int32(12))+4))
	if v20 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v29 = v6
	goto L1
L5:
	;
	v29 = v12
	goto L1
L6:
	;
	goto L7
L7:
	;
	v24 = v12 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v24)) < base.Ui64(v7) {
		v12 = v24
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
}
func F_tzload(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
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
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
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
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v725 int32
	_ = v725
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v837 int32
	_ = v837
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
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
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1243 int32
	_ = v1243
	var v1283 int32
	_ = v1283
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1354 int32
	_ = v1354
	var v1366 int32
	_ = v1366
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1404 int32
	_ = v1404
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1610 int32
	_ = v1610
	var v1612 int64
	_ = v1612
	var v1613 int64
	_ = v1613
	var v1614 int64
	_ = v1614
	var v1615 int64
	_ = v1615
	var v1616 int64
	_ = v1616
	var v1620 int64
	_ = v1620
	var v1633 int64
	_ = v1633
	var v1640 int64
	_ = v1640
	var v1645 int64
	_ = v1645
	var v1648 int64
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1656 int32
	_ = v1656
	var v1660 int64
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1850 int32
	_ = v1850
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1895 int32
	_ = v1895
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1925 int32
	_ = v1925
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1968 int32
	_ = v1968
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2133 int32
	_ = v2133
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2218 int64
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2230 int32
	_ = v2230
	var v2236 int32
	_ = v2236
	var v2241 int32
	_ = v2241
	var v2243 int64
	_ = v2243
	var v2244 int64
	_ = v2244
	var v2245 int64
	_ = v2245
	var v2246 int64
	_ = v2246
	var v2247 int64
	_ = v2247
	var v2251 int64
	_ = v2251
	var v2264 int64
	_ = v2264
	var v2271 int64
	_ = v2271
	var v2276 int64
	_ = v2276
	var v2279 int64
	_ = v2279
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2321 int32
	_ = v2321
	var v2322 int32
	_ = v2322
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2412 int32
	_ = v2412
	var v2416 int32
	_ = v2416
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2475 int32
	_ = v2475
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2565 int32
	_ = v2565
	var v2567 int32
	_ = v2567
	var v2574 int32
	_ = v2574
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2589 int32
	_ = v2589
	var v2596 int32
	_ = v2596
	var v2600 int32
	_ = v2600
	var v2603 int32
	_ = v2603
	var v2609 int32
	_ = v2609
	var v2616 int32
	_ = v2616
	var v2620 int32
	_ = v2620
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2637 int32
	_ = v2637
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2666 int32
	_ = v2666
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2700 int32
	_ = v2700
	var v2708 int32
	_ = v2708
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2771 int32
	_ = v2771
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2829 int32
	_ = v2829
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2842 int32
	_ = v2842
	var v2851 int32
	_ = v2851
	var v2893 int32
	_ = v2893
	var v2898 int32
	_ = v2898
	var v2902 int32
	_ = v2902
	var v2907 int32
	_ = v2907
	var v2909 int32
	_ = v2909
	var v2913 int32
	_ = v2913
	var v2919 int32
	_ = v2919
	var v2922 int32
	_ = v2922
	var v2928 int32
	_ = v2928
	var v2932 int32
	_ = v2932
	var v2934 int32
	_ = v2934
	var v2942 int32
	_ = v2942
	var v2943 int32
	_ = v2943
	var v2946 int32
	_ = v2946
	var v2952 int32
	_ = v2952
	var v2956 int32
	_ = v2956
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2970 int32
	_ = v2970
	var v2972 int32
	_ = v2972
	var v2974 int32
	_ = v2974
	var v2977 int32
	_ = v2977
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2994 int32
	_ = v2994
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3008 int32
	_ = v3008
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3026 int32
	_ = v3026
	var v3030 int32
	_ = v3030
	var v3071 int32
	_ = v3071
	var v3076 int32
	_ = v3076
	var v3111 int32
	_ = v3111
	var v3116 int32
	_ = v3116
	var v3121 int32
	_ = v3121
	var v3157 int32
	_ = v3157
	var v3163 int32
	_ = v3163
	var v3203 int32
	_ = v3203
	var v3205 int32
	_ = v3205
	var v3209 int32
	_ = v3209
	var v3218 int32
	_ = v3218
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3270 int64
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3279 int32
	_ = v3279
	var v3317 int64
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3362 int32
	_ = v3362
	var v3367 int32
	_ = v3367
	var v3368 int64
	_ = v3368
	var v3370 int64
	_ = v3370
	var v3372 int64
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3385 int32
	_ = v3385
	var v3428 int32
	_ = v3428
	var v3432 int32
	_ = v3432
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3476 int64
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3521 int32
	_ = v3521
	var v3526 int32
	_ = v3526
	var v3527 int64
	_ = v3527
	var v3529 int64
	_ = v3529
	var v3531 int64
	_ = v3531
	var v3534 int32
	_ = v3534
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3539 int32
	_ = v3539
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3592 int32
	_ = v3592
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3648 int32
	_ = v3648
	var v3650 int32
	_ = v3650
	var v3653 int32
	_ = v3653
	var v3654 int64
	_ = v3654
	var v3656 int64
	_ = v3656
	var v3658 int32
	_ = v3658
	var v3664 int32
	_ = v3664
	var v3665 int64
	_ = v3665
	var v3667 int64
	_ = v3667
	var v3669 int32
	_ = v3669
	var v3670 int32
	_ = v3670
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3718 int32
	_ = v3718
	var v3722 int32
	_ = v3722
	var v3724 int32
	_ = v3724
	var v3727 int32
	_ = v3727
	var v3728 int64
	_ = v3728
	var v3730 int64
	_ = v3730
	var v3774 int32
	_ = v3774
	var v3783 int32
	_ = v3783
	var v3819 int32
	_ = v3819
	var v3822 int32
	_ = v3822
	var v3825 int32
	_ = v3825
	var v3827 int32
	_ = v3827
	var v3871 int32
	_ = v3871
	var v3875 int32
	_ = v3875
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3879 int32
	_ = v3879
	var v3880 int32
	_ = v3880
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3899 int32
	_ = v3899
	var v3900 int32
	_ = v3900
	var v3903 int32
	_ = v3903
	var v3904 int32
	_ = v3904
	var v3907 int32
	_ = v3907
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3920 int64
	_ = v3920
	var v3921 int64
	_ = v3921
	var v3925 int32
	_ = v3925
	var v3929 int32
	_ = v3929
	var v3936 int32
	_ = v3936
	var v3941 int32
	_ = v3941
	var v3944 int32
	_ = v3944
	var v3946 int32
	_ = v3946
	var v3989 int32
	_ = v3989
	var v3991 int32
	_ = v3991
	var v3994 int32
	_ = v3994
	var v3995 int32
	_ = v3995
	var v3997 int32
	_ = v3997
	var v3998 int32
	_ = v3998
	var v4000 int32
	_ = v4000
	var v4001 int32
	_ = v4001
	var v4003 int32
	_ = v4003
	var v4004 int32
	_ = v4004
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4008 int32
	_ = v4008
	var v4009 int32
	_ = v4009
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
	var v4021 int32
	_ = v4021
	var v4022 int32
	_ = v4022
	var v4025 int32
	_ = v4025
	var v4032 int32
	_ = v4032
	var v4033 int32
	_ = v4033
	var v4035 int64
	_ = v4035
	var v4039 int64
	_ = v4039
	var v4043 int32
	_ = v4043
	var v4092 int32
	_ = v4092
	var v4096 int32
	_ = v4096
	var v4139 int32
	_ = v4139
	var v4141 int32
	_ = v4141
	var v4143 int32
	_ = v4143
	var v4147 int32
	_ = v4147
	var v4150 int32
	_ = v4150
	var v4195 int32
	_ = v4195
	var v4199 int32
	_ = v4199
	var v4242 int32
	_ = v4242
	var v4245 int32
	_ = v4245
	var v4249 int32
	_ = v4249
	var v4294 int32
	_ = v4294
	var v4298 int32
	_ = v4298
	var v4343 int32
	_ = v4343
	var v4412 int32
	_ = v4412
	var v4431 int32
	_ = v4431
	v44 = F_emscripten_builtin_malloc(m, int32(_a_F_tzload_0))
	mBase = m.M
	if v44 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_tzload[0]))
	return v48
L2:
	;
	goto L3
L3:
	;
	v50 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)) = uint16(v50)
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v53 = l0
	goto L6
L5:
	;
	v53 = int32(_a_F_tzload_1)
	goto L6
L6:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v57 = v53 + base.B2i32(v54 == int32(58))
	v58 = m.G0
	v60 = v58 - int32(1056)
	m.G0 = v60
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tzload[1])))
	if v63 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_get_share_path(m, int32(_a_F_tzload_2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v253 = v60 + int32(32)
	v254 = int32(_a_F_tzload_2)
	goto L64
L10:
	;
	return int32(0)
L11:
	;
	v71 = int32(_a_F_tzload_2)
	goto L14
L12:
	;
	v130 = v123 + int32(_a_F_tzload_2)
	v131 = int32(_a_F_tzload_3)
	v133 = int32(1024) - v123
	if v133 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	v123 = v114 - v71
	goto L12
L14:
	;
	v99 = v71
	goto L23
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v108 = int32(-2139062144)
	if (int32(16843008)-v105|v105)&v108 == v108 {
		v99 = v99 + int32(4)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v114 = v99
	goto L26
L25:
	;
	goto L24
L26:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v118 != 0 {
		v114 = v114 + int32(1)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L13
L28:
	;
	goto L27
L29:
	;
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_tzload[1])) = uint8(v249)
	goto L9
L30:
	;
	v245 = F_strlen(m, v241)
	mBase = m.M
	goto L29
L31:
	;
	v241 = v131
	goto L30
L32:
	;
	goto L33
L33:
	;
	v139 = v133 - int32(1)
	if (v130^v131)&int32(3) != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v238)
	v241 = v234
	goto L30
L35:
	;
	v219 = v214
	v220 = v215
	v221 = v216
	goto L57
L36:
	;
	if v209 == int32(0) {
		v234 = v207
		v235 = v208
		goto L34
	} else {
		goto L56
	}
L37:
	;
	v207 = v131
	v208 = v130
	v209 = v139
	goto L36
L38:
	;
	goto L39
L39:
	;
	goto L42
L40:
	;
	if v176 == int32(0) {
		v234 = v173
		v235 = v174
		goto L34
	} else {
		goto L49
	}
L41:
	;
	v173 = v131
	v174 = v130
	v175 = v139
	v176 = base.B2i32(v139 != int32(0))
	goto L40
L42:
	;
	if v139 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v152 = v131
	v153 = v130
	v154 = v139
	goto L44
L44:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v156)
	if v156 == int32(0) {
		v214 = v152
		v215 = v153
		v216 = v154
		goto L35
	} else {
		goto L46
	}
L45:
	;
	v173 = v167
	v174 = v161
	v175 = v163
	v176 = v165
	goto L40
L46:
	;
	v160 = int32(1)
	v161 = v153 + v160
	v163 = v154 - v160
	v164 = int32(0)
	v165 = base.B2i32(v163 != v164)
	v167 = v152 + v160
	if v167&int32(3) == v164 {
		v173 = v167
		v174 = v161
		v175 = v163
		v176 = v165
		goto L40
	} else {
		goto L47
	}
L47:
	;
	if v163 != 0 {
		v152 = v167
		v153 = v161
		v154 = v163
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v179 == int32(0) {
		v207 = v173
		v208 = v174
		v209 = v175
		goto L36
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v175) < base.Ui32(int32(4)) {
		v207 = v173
		v208 = v174
		v209 = v175
		goto L36
	} else {
		goto L51
	}
L51:
	;
	v185 = v173
	v186 = v174
	v187 = v175
	goto L52
L52:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v193 = int32(-2139062144)
	if (int32(16843008)-v190|v190)&v193 != v193 {
		v214 = v185
		v215 = v186
		v216 = v187
		goto L35
	} else {
		goto L54
	}
L53:
	;
	v207 = v201
	v208 = v199
	v209 = v203
	goto L36
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v190
	v198 = int32(4)
	v199 = v186 + v198
	v201 = v185 + v198
	v203 = v187 - v198
	if base.Ui32(int32(3)) < base.Ui32(v203) {
		v185 = v201
		v186 = v199
		v187 = v203
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v214 = v207
	v215 = v208
	v216 = v209
	goto L35
L57:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v220))) = uint8(v223)
	if v223 == int32(0) {
		v234 = v219
		v235 = v220
		goto L34
	} else {
		goto L59
	}
L58:
	;
	v234 = v230
	v235 = v228
	goto L34
L59:
	;
	v227 = int32(1)
	v228 = v220 + v227
	v230 = v219 + v227
	v232 = v221 - v227
	if v232 != 0 {
		v219 = v230
		v220 = v228
		v221 = v232
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v372 = v60 + int32(32)
	if v372&int32(3) == int32(0) {
		v396 = v372
		goto L96
	} else {
		goto L97
	}
L62:
	;
	v367 = F_strlen(m, v356)
	mBase = m.M
	goto L61
L64:
	;
	goto L65
L65:
	;
	v261 = int32(1023)
	if (v253^v254)&int32(3) != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v360)
	goto L62
L67:
	;
	v341 = v336
	v342 = v337
	v343 = v338
	goto L89
L68:
	;
	if v331 == int32(0) {
		v356 = v329
		v357 = v330
		goto L66
	} else {
		goto L88
	}
L69:
	;
	v329 = v254
	v330 = v253
	v331 = v261
	goto L68
L70:
	;
	goto L71
L71:
	;
	goto L73
L72:
	;
	goto L81
L73:
	;
	goto L72
L81:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tzload[2])))
	if v301 == int32(0) {
		v329 = v254
		v330 = v253
		v331 = v261
		goto L68
	} else {
		goto L82
	}
L82:
	;
	goto L83
L83:
	;
	v307 = v254
	v308 = v253
	v309 = v261
	goto L84
L84:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v315 = int32(-2139062144)
	if (int32(16843008)-v312|v312)&v315 != v315 {
		v336 = v307
		v337 = v308
		v338 = v309
		goto L67
	} else {
		goto L86
	}
L85:
	;
	v329 = v323
	v330 = v321
	v331 = v325
	goto L68
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v312
	v320 = int32(4)
	v321 = v308 + v320
	v323 = v307 + v320
	v325 = v309 - v320
	if base.Ui32(int32(3)) < base.Ui32(v325) {
		v307 = v323
		v308 = v321
		v309 = v325
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v336 = v329
	v337 = v330
	v338 = v331
	goto L67
L89:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	*(*uint8)(unsafe.Add(mBase, uint32(v342))) = uint8(v345)
	if v345 == int32(0) {
		v356 = v341
		v357 = v342
		goto L66
	} else {
		goto L91
	}
L90:
	;
	v356 = v352
	v357 = v350
	goto L66
L91:
	;
	v349 = int32(1)
	v350 = v342 + v349
	v352 = v341 + v349
	v354 = v343 - v349
	if v354 != 0 {
		v341 = v352
		v342 = v350
		v343 = v354
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	m.G0 = v60 + int32(1056)
	if v1243 < int32(0) {
		goto L315
	} else {
		goto L316
	}
L94:
	;
	if v57&int32(3) == int32(0) {
		v453 = v57
		goto L113
	} else {
		goto L114
	}
L95:
	;
	v429 = v421 - v372
	goto L94
L96:
	;
	v400 = v396
	goto L105
L97:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v380 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v429 = int32(0)
	goto L94
L99:
	;
	goto L100
L100:
	;
	v385 = v372
	goto L101
L101:
	;
	v389 = v385 + int32(1)
	if v389&int32(3) == int32(0) {
		v396 = v389
		goto L96
	} else {
		goto L103
	}
L102:
	;
	v421 = v389
	goto L95
L103:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v394 != 0 {
		v385 = v389
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v409 = int32(-2139062144)
	if (int32(16843008)-v406|v406)&v409 == v409 {
		v400 = v400 + int32(4)
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v415 = v400
	goto L108
L107:
	;
	goto L106
L108:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	if v419 != 0 {
		v415 = v415 + int32(1)
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v421 = v415
	goto L95
L110:
	;
	goto L109
L111:
	;
	if base.Ui32(int32(1023)) < base.Ui32(v429+v486+int32(1)) {
		v1243 = int32(-1)
		goto L93
	} else {
		goto L128
	}
L112:
	;
	v486 = v478 - v57
	goto L111
L113:
	;
	v457 = v453
	goto L122
L114:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v437 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v486 = int32(0)
	goto L111
L116:
	;
	goto L117
L117:
	;
	v442 = v57
	goto L118
L118:
	;
	v446 = v442 + int32(1)
	if v446&int32(3) == int32(0) {
		v453 = v446
		goto L113
	} else {
		goto L120
	}
L119:
	;
	v478 = v446
	goto L112
L120:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	if v451 != 0 {
		v442 = v446
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v466 = int32(-2139062144)
	if (int32(16843008)-v463|v463)&v466 == v466 {
		v457 = v457 + int32(4)
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v472 = v457
	goto L125
L124:
	;
	goto L123
L125:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	if v476 != 0 {
		v472 = v472 + int32(1)
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v478 = v472
	goto L112
L127:
	;
	goto L126
L128:
	;
	if l1 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v496 = v60 + int32(32) + v429
	v497 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v496))) = uint8(v497)
	v500 = v496 + int32(1)
	if (v57^v500)&int32(3) != 0 {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	goto L131
L131:
	;
	v600 = v429
	v608 = v57
	goto L154
L132:
	;
	v575 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v575
	v582 = F_open(m, v60+int32(32), v575, v60+int32(16))
	mBase = m.M
	if v575 <= v582 {
		v1243 = v582
		goto L93
	} else {
		goto L153
	}
L133:
	;
	goto L132
L134:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v555))) = uint8(v554)
	if v554&int32(255) == int32(0) {
		goto L133
	} else {
		goto L149
	}
L135:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v553 = v57
	v554 = v506
	v555 = v500
	goto L134
L136:
	;
	goto L137
L137:
	;
	if v57&int32(3) != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v510 = v57
	v512 = v500
	goto L141
L139:
	;
	v524 = v57
	v526 = v500
	goto L140
L140:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v531 = int32(-2139062144)
	if (int32(16843008)-v528|v528)&v531 != v531 {
		v553 = v524
		v554 = v528
		v555 = v526
		goto L134
	} else {
		goto L145
	}
L141:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510))))
	*(*uint8)(unsafe.Add(mBase, uint32(v512))) = uint8(v513)
	if v513 == int32(0) {
		goto L133
	} else {
		goto L143
	}
L142:
	;
	v524 = v520
	v526 = v518
	goto L140
L143:
	;
	v517 = int32(1)
	v518 = v512 + v517
	v520 = v510 + v517
	if v520&int32(3) != 0 {
		v510 = v520
		v512 = v518
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v536 = v524
	v537 = v528
	v538 = v526
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v537
	v540 = int32(4)
	v541 = v538 + v540
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v536)+4))
	v544 = v536 + v540
	v548 = int32(-2139062144)
	if (v542|(int32(16843008)-v542))&v548 == v548 {
		v536 = v544
		v537 = v542
		v538 = v541
		goto L146
	} else {
		goto L148
	}
L147:
	;
	v553 = v544
	v554 = v542
	v555 = v541
	goto L134
L148:
	;
	goto L147
L149:
	;
	v562 = v553
	v564 = v555
	goto L150
L150:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v564)+1)) = uint8(v565)
	v567 = int32(1)
	if v565 != 0 {
		v562 = v562 + v567
		v564 = v564 + v567
		goto L150
	} else {
		goto L152
	}
L151:
	;
	goto L133
L152:
	;
	goto L151
L153:
	;
	v585 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v496))) = uint8(v585)
	goto L131
L154:
	;
	v635 = int32(47)
	v636 = F___strchrnul(m, v608, v635)
	mBase = m.M
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636))))
	if v638 == v635 {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	if l1 != 0 {
		goto L280
	} else {
		goto L281
	}
L156:
	;
	v704 = F_AllocateDir(m, v60+int32(32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L10
	} else {
		goto L182
	}
L157:
	;
	if v642 != 0 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	v642 = v636
	goto L160
L159:
	;
	v642 = int32(0)
	goto L160
L160:
	;
	goto L157
L161:
	;
	v701 = v642 - v608
	goto L156
L162:
	;
	goto L163
L163:
	;
	if v608&int32(3) == int32(0) {
		v667 = v608
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v701 = v700
	goto L156
L165:
	;
	v700 = v692 - v608
	goto L164
L166:
	;
	v671 = v667
	goto L175
L167:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v651 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v700 = int32(0)
	goto L164
L169:
	;
	goto L170
L170:
	;
	v656 = v608
	goto L171
L171:
	;
	v660 = v656 + int32(1)
	if v660&int32(3) == int32(0) {
		v667 = v660
		goto L166
	} else {
		goto L173
	}
L172:
	;
	v692 = v660
	goto L165
L173:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	if v665 != 0 {
		v656 = v660
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	v680 = int32(-2139062144)
	if (int32(16843008)-v677|v677)&v680 == v680 {
		v671 = v671 + int32(4)
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v686 = v671
	goto L178
L177:
	;
	goto L176
L178:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686))))
	if v690 != 0 {
		v686 = v686 + int32(1)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v692 = v686
	goto L165
L180:
	;
	goto L179
L181:
	;
	if v712 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L182:
	;
	v709 = F_ReadDirExtended(m, v704, v60+int32(32), int32(15))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	if v709 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v712 = int32(1023) - v600
	v713 = v600 + (v60 + int32(32) | int32(1))
	v725 = v709
	goto L187
L185:
	;
	goto L186
L186:
	;
	F_FreeDir(m, v704)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L10
	} else {
		goto L228
	}
L187:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725)+19)))
	if v756 == int32(46) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L186
L189:
	;
	v875 = F_ReadDirExtended(m, v704, v60+int32(32), int32(15))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L226
	}
L190:
	;
	v760 = v725 + int32(19)
	if v760&int32(3) == int32(0) {
		v784 = v760
		goto L193
	} else {
		goto L194
	}
L191:
	;
	if v817 != v701 {
		goto L189
	} else {
		goto L208
	}
L192:
	;
	v817 = v809 - v760
	goto L191
L193:
	;
	v788 = v784
	goto L202
L194:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760))))
	if v768 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v817 = int32(0)
	goto L191
L196:
	;
	goto L197
L197:
	;
	v773 = v760
	goto L198
L198:
	;
	v777 = v773 + int32(1)
	if v777&int32(3) == int32(0) {
		v784 = v777
		goto L193
	} else {
		goto L200
	}
L199:
	;
	v809 = v777
	goto L192
L200:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
	if v782 != 0 {
		v773 = v777
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v788)))
	v797 = int32(-2139062144)
	if (int32(16843008)-v794|v794)&v797 == v797 {
		v788 = v788 + int32(4)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v803 = v788
	goto L205
L204:
	;
	goto L203
L205:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
	if v807 != 0 {
		v803 = v803 + int32(1)
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v809 = v803
	goto L192
L207:
	;
	goto L206
L208:
	;
	v821 = v760
	v822 = v608
	v823 = v701
	goto L210
L209:
	;
	if v868 == int32(0) {
		goto L181
	} else {
		goto L225
	}
L210:
	;
	if v823 != 0 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v868 = int32(0)
	goto L209
L212:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822))))
	if v826 == v827 {
		v849 = v826
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	goto L211
L215:
	;
	v851 = int32(1)
	if v849 != 0 {
		v821 = v821 + v851
		v822 = v822 + v851
		v823 = v823 - v851
		goto L210
	} else {
		goto L224
	}
L216:
	;
	if base.Ui32((v826-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v837 = v826 | int32(32)
	goto L219
L218:
	;
	v837 = v826
	goto L219
L219:
	;
	if base.Ui32((v827-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v846 = v827 | int32(32)
	goto L222
L221:
	;
	v846 = v827
	goto L222
L222:
	;
	if v837 == v846 {
		v849 = v837
		goto L215
	} else {
		goto L223
	}
L223:
	;
	v868 = v837 - v846
	goto L209
L224:
	;
	goto L214
L225:
	;
	goto L189
L226:
	;
	if v875 != 0 {
		v725 = v875
		goto L187
	} else {
		goto L227
	}
L227:
	;
	goto L188
L228:
	;
	v1243 = int32(-1)
	goto L93
L229:
	;
	F_FreeDir(m, v704)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L10
	} else {
		goto L261
	}
L230:
	;
	v1033 = F_strlen(m, v1029)
	mBase = m.M
	goto L229
L231:
	;
	v1029 = v760
	goto L230
L232:
	;
	goto L233
L233:
	;
	v927 = v712 - int32(1)
	if (v713^v760)&int32(3) != 0 {
		goto L237
	} else {
		goto L238
	}
L234:
	;
	v1026 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1023))) = uint8(v1026)
	v1029 = v1022
	goto L230
L235:
	;
	v1007 = v1002
	v1008 = v1003
	v1009 = v1004
	goto L257
L236:
	;
	if v997 == int32(0) {
		v1022 = v995
		v1023 = v996
		goto L234
	} else {
		goto L256
	}
L237:
	;
	v995 = v760
	v996 = v713
	v997 = v927
	goto L236
L238:
	;
	goto L239
L239:
	;
	v931 = int32(0)
	if v760&int32(3) == v931 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	if v964 == int32(0) {
		v1022 = v961
		v1023 = v962
		goto L234
	} else {
		goto L249
	}
L241:
	;
	v961 = v760
	v962 = v713
	v963 = v927
	v964 = base.B2i32(v927 != v931)
	goto L240
L242:
	;
	if v927 == int32(0) {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v940 = v760
	v941 = v713
	v942 = v927
	goto L244
L244:
	;
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940))))
	*(*uint8)(unsafe.Add(mBase, uint32(v941))) = uint8(v944)
	if v944 == int32(0) {
		v1002 = v940
		v1003 = v941
		v1004 = v942
		goto L235
	} else {
		goto L246
	}
L245:
	;
	v961 = v955
	v962 = v949
	v963 = v951
	v964 = v953
	goto L240
L246:
	;
	v948 = int32(1)
	v949 = v941 + v948
	v951 = v942 - v948
	v952 = int32(0)
	v953 = base.B2i32(v951 != v952)
	v955 = v940 + v948
	if v955&int32(3) == v952 {
		v961 = v955
		v962 = v949
		v963 = v951
		v964 = v953
		goto L240
	} else {
		goto L247
	}
L247:
	;
	if v951 != 0 {
		v940 = v955
		v941 = v949
		v942 = v951
		goto L244
	} else {
		goto L248
	}
L248:
	;
	goto L245
L249:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961))))
	if v967 == int32(0) {
		v995 = v961
		v996 = v962
		v997 = v963
		goto L236
	} else {
		goto L250
	}
L250:
	;
	if base.Ui32(v963) < base.Ui32(int32(4)) {
		v995 = v961
		v996 = v962
		v997 = v963
		goto L236
	} else {
		goto L251
	}
L251:
	;
	v973 = v961
	v974 = v962
	v975 = v963
	goto L252
L252:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v973)))
	v981 = int32(-2139062144)
	if (int32(16843008)-v978|v978)&v981 != v981 {
		v1002 = v973
		v1003 = v974
		v1004 = v975
		goto L235
	} else {
		goto L254
	}
L253:
	;
	v995 = v989
	v996 = v987
	v997 = v991
	goto L236
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = v978
	v986 = int32(4)
	v987 = v974 + v986
	v989 = v973 + v986
	v991 = v975 - v986
	if base.Ui32(int32(3)) < base.Ui32(v991) {
		v973 = v989
		v974 = v987
		v975 = v991
		goto L252
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v1002 = v995
	v1003 = v996
	v1004 = v997
	goto L235
L257:
	;
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1008))) = uint8(v1011)
	if v1011 == int32(0) {
		v1022 = v1007
		v1023 = v1008
		goto L234
	} else {
		goto L259
	}
L258:
	;
	v1022 = v1018
	v1023 = v1016
	goto L234
L259:
	;
	v1015 = int32(1)
	v1016 = v1008 + v1015
	v1018 = v1007 + v1015
	v1020 = v1009 - v1015
	if v1020 != 0 {
		v1007 = v1018
		v1008 = v1016
		v1009 = v1020
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v1039 = v60 + int32(32)
	v1041 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v1039+v600))) = uint8(v1041)
	v1044 = v600 + int32(1)
	v1047 = v1044 + v1039
	if v1047&int32(3) == int32(0) {
		v1071 = v1047
		goto L264
	} else {
		goto L265
	}
L262:
	;
	if v642 != 0 {
		v600 = v1104 + v1044
		v608 = v642 + int32(1)
		goto L154
	} else {
		goto L279
	}
L263:
	;
	v1104 = v1096 - v1047
	goto L262
L264:
	;
	v1075 = v1071
	goto L273
L265:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	if v1055 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1104 = int32(0)
	goto L262
L267:
	;
	goto L268
L268:
	;
	v1060 = v1047
	goto L269
L269:
	;
	v1064 = v1060 + int32(1)
	if v1064&int32(3) == int32(0) {
		v1071 = v1064
		goto L264
	} else {
		goto L271
	}
L270:
	;
	v1096 = v1064
	goto L263
L271:
	;
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1064))))
	if v1069 != 0 {
		v1060 = v1064
		goto L269
	} else {
		goto L272
	}
L272:
	;
	goto L270
L273:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	v1084 = int32(-2139062144)
	if (int32(16843008)-v1081|v1081)&v1084 == v1084 {
		v1075 = v1075 + int32(4)
		goto L273
	} else {
		goto L275
	}
L274:
	;
	v1090 = v1075
	goto L276
L275:
	;
	goto L274
L276:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090))))
	if v1094 != 0 {
		v1090 = v1090 + int32(1)
		goto L276
	} else {
		goto L278
	}
L277:
	;
	v1096 = v1090
	goto L263
L278:
	;
	goto L277
L279:
	;
	goto L155
L280:
	;
	v1110 = v429 + v60 + int32(33)
	goto L286
L281:
	;
	goto L282
L282:
	;
	v1226 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v1226
	v1231 = F_open(m, v60+int32(32), v1226, v60)
	mBase = m.M
	v1243 = v1231
	goto L93
L283:
	;
	goto L282
L284:
	;
	v1223 = F_strlen(m, v1212)
	mBase = m.M
	goto L283
L286:
	;
	goto L287
L287:
	;
	v1117 = int32(255)
	if (l1^v1110)&int32(3) != 0 {
		goto L291
	} else {
		goto L292
	}
L288:
	;
	v1216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1213))) = uint8(v1216)
	goto L284
L289:
	;
	v1197 = v1192
	v1198 = v1193
	v1199 = v1194
	goto L311
L290:
	;
	if v1187 == int32(0) {
		v1212 = v1185
		v1213 = v1186
		goto L288
	} else {
		goto L310
	}
L291:
	;
	v1185 = v1110
	v1186 = l1
	v1187 = v1117
	goto L290
L292:
	;
	goto L293
L293:
	;
	if v1110&int32(3) == int32(0) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	if v1154 == int32(0) {
		v1212 = v1151
		v1213 = v1152
		goto L288
	} else {
		goto L303
	}
L295:
	;
	v1151 = v1110
	v1152 = l1
	v1153 = v1117
	v1154 = int32(1)
	goto L294
L296:
	;
	goto L297
L297:
	;
	v1130 = v1110
	v1131 = l1
	v1132 = v1117
	goto L298
L298:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1130))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1131))) = uint8(v1134)
	if v1134 == int32(0) {
		v1192 = v1130
		v1193 = v1131
		v1194 = v1132
		goto L289
	} else {
		goto L300
	}
L299:
	;
	v1151 = v1145
	v1152 = v1139
	v1153 = v1141
	v1154 = v1143
	goto L294
L300:
	;
	v1138 = int32(1)
	v1139 = v1131 + v1138
	v1141 = v1132 - v1138
	v1142 = int32(0)
	v1143 = base.B2i32(v1141 != v1142)
	v1145 = v1130 + v1138
	if v1145&int32(3) == v1142 {
		v1151 = v1145
		v1152 = v1139
		v1153 = v1141
		v1154 = v1143
		goto L294
	} else {
		goto L301
	}
L301:
	;
	if v1141 != 0 {
		v1130 = v1145
		v1131 = v1139
		v1132 = v1141
		goto L298
	} else {
		goto L302
	}
L302:
	;
	goto L299
L303:
	;
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	if v1157 == int32(0) {
		v1185 = v1151
		v1186 = v1152
		v1187 = v1153
		goto L290
	} else {
		goto L304
	}
L304:
	;
	if base.Ui32(v1153) < base.Ui32(int32(4)) {
		v1185 = v1151
		v1186 = v1152
		v1187 = v1153
		goto L290
	} else {
		goto L305
	}
L305:
	;
	v1163 = v1151
	v1164 = v1152
	v1165 = v1153
	goto L306
L306:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1163)))
	v1171 = int32(-2139062144)
	if (int32(16843008)-v1168|v1168)&v1171 != v1171 {
		v1192 = v1163
		v1193 = v1164
		v1194 = v1165
		goto L289
	} else {
		goto L308
	}
L307:
	;
	v1185 = v1179
	v1186 = v1177
	v1187 = v1181
	goto L290
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1164))) = v1168
	v1176 = int32(4)
	v1177 = v1164 + v1176
	v1179 = v1163 + v1176
	v1181 = v1165 - v1176
	if base.Ui32(int32(3)) < base.Ui32(v1181) {
		v1163 = v1179
		v1164 = v1177
		v1165 = v1181
		goto L306
	} else {
		goto L309
	}
L309:
	;
	goto L307
L310:
	;
	v1192 = v1185
	v1193 = v1186
	v1194 = v1187
	goto L289
L311:
	;
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1198))) = uint8(v1201)
	if v1201 == int32(0) {
		v1212 = v1197
		v1213 = v1198
		goto L288
	} else {
		goto L313
	}
L312:
	;
	v1212 = v1208
	v1213 = v1206
	goto L288
L313:
	;
	v1205 = int32(1)
	v1206 = v1198 + v1205
	v1208 = v1197 + v1205
	v1210 = v1199 - v1205
	if v1210 != 0 {
		v1197 = v1208
		v1198 = v1206
		v1199 = v1210
		goto L311
	} else {
		goto L314
	}
L314:
	;
	goto L312
L315:
	;
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return int32(44)
L316:
	;
	goto L317
L317:
	;
	v1283 = F_read(m, v1243, v44, int32(_a_F_tzload_4))
	mBase = m.M
	if v1283 <= int32(43) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	if v1283 < int32(0) {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	goto L320
L320:
	;
	v1295 = F_close(m, v1243)
	mBase = m.M
	if int32(0) <= v1295 {
		goto L324
	} else {
		goto L325
	}
L321:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, _c_F_tzload[0]))
	v1291 = v1290
	goto L323
L322:
	;
	v1291 = int32(28)
	goto L323
L323:
	;
	v1292 = F_close(m, v1243)
	mBase = m.M
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return v1291
L324:
	;
	v1299 = l2 + int32(_a_F_tzload_5)
	v1301 = l2 + int32(_a_F_tzload_6)
	v1303 = l2 + int32(_a_F_tzload_7)
	v1305 = l2 + int32(24)
	v1307 = l2 + int32(_a_F_tzload_8)
	v1309 = v44 + int32(44)
	v1354 = int32(4)
	v1366 = v1283
	goto L328
L325:
	;
	goto L326
L326:
	;
	v4431 = *(*int32)(unsafe.Add(mBase, _c_F_tzload[0]))
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return v4431
L327:
	;
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return v4412
L328:
	;
	v1389 = int32(28)
	v1390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(31)))))
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(30)))))
	v1392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+28)))
	v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(29)))))
	v1398 = int32(8)
	v1404 = v1390 | (v1391|(v1392&int32(127)<<(uint(int32(16))%32)|v1397<<(uint(v1398)%32)))<<(uint(v1398)%32)
	if v1392 < int32(0) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	if v2684 < int32(3) {
		goto L511
	} else {
		goto L512
	}
L330:
	;
	v1409 = v1404 | int32(-2147483648)
	goto L332
L331:
	;
	v1409 = v1404
	goto L332
L332:
	;
	if base.Ui32(int32(49)) < base.Ui32(v1409) {
		v4412 = v1389
		goto L327
	} else {
		goto L333
	}
L333:
	;
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(39)))))
	v1413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(38)))))
	v1414 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+36)))
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(37)))))
	v1420 = int32(8)
	v1426 = v1412 | (v1413|(v1414&int32(127)<<(uint(int32(16))%32)|v1419<<(uint(v1420)%32)))<<(uint(v1420)%32)
	if v1414 < int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1431 = v1426 | int32(-2147483648)
	goto L336
L335:
	;
	v1431 = v1426
	goto L336
L336:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1431) {
		v4412 = v1389
		goto L327
	} else {
		goto L337
	}
L337:
	;
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(35)))))
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(34)))))
	v1436 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+32)))
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(33)))))
	v1442 = int32(8)
	v1448 = v1434 | (v1435|(v1436&int32(127)<<(uint(int32(16))%32)|v1441<<(uint(v1442)%32)))<<(uint(v1442)%32)
	if v1436 < int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1453 = v1448 | int32(-2147483648)
	goto L340
L339:
	;
	v1453 = v1448
	goto L340
L340:
	;
	if base.Ui32(int32(1999)) < base.Ui32(v1453) {
		v4412 = v1389
		goto L327
	} else {
		goto L341
	}
L341:
	;
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(43)))))
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(42)))))
	v1458 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+40)))
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(41)))))
	v1464 = int32(8)
	v1470 = v1456 | (v1457|(v1458&int32(127)<<(uint(int32(16))%32)|v1463<<(uint(v1464)%32)))<<(uint(v1464)%32)
	if v1458 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1475 = v1470 | int32(-2147483648)
	goto L344
L343:
	;
	v1475 = v1470
	goto L344
L344:
	;
	if base.Ui32(int32(49)) < base.Ui32(v1475) {
		v4412 = v1389
		goto L327
	} else {
		goto L345
	}
L345:
	;
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(21)))))
	v1479 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+20)))
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(22)))))
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(23)))))
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(27)))))
	v1483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(26)))))
	v1484 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+24)))
	v1489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(25)))))
	v1490 = int32(8)
	v1496 = v1482 | (v1483|(v1484&int32(127)<<(uint(int32(16))%32)|v1489<<(uint(v1490)%32)))<<(uint(v1490)%32)
	if v1484 < int32(0) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1501 = v1496 | int32(-2147483648)
	goto L348
L347:
	;
	v1501 = v1496
	goto L348
L348:
	;
	if v1431 != v1501 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1504 = v1501
	goto L351
L350:
	;
	v1504 = int32(0)
	goto L351
L351:
	;
	if v1504 != 0 {
		v4412 = v1389
		goto L327
	} else {
		goto L352
	}
L352:
	;
	v1509 = int32(8)
	v1515 = (v1479&int32(127)<<(uint(int32(16))%32)|v1478<<(uint(v1509)%32)|v1480)<<(uint(v1509)%32) | v1481
	if v1479 < int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v1520 = v1515 | int32(-2147483648)
	goto L355
L354:
	;
	v1520 = v1515
	goto L355
L355:
	;
	if v1431 != v1520 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1523 = v1520
	goto L358
L357:
	;
	v1523 = int32(0)
	goto L358
L358:
	;
	if v1523 != 0 {
		v4412 = v1389
		goto L327
	} else {
		goto L359
	}
L359:
	;
	v1526 = v1354 + int32(4)
	if v1366 < v1520+v1501+v1409*v1526+v1453+v1453*v1354+v1431*int32(6)+v1475+int32(44) {
		v4412 = v1389
		goto L327
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v1475
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1431
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1453
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1409
	v1543 = int32(0)
	if v1453 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1546 = v1309
	v1547 = int32(0)
	v1549 = v1543
	goto L364
L362:
	;
	v1787 = v1309
	v1788 = v1431
	v1790 = v1543
	goto L363
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1790
	if v1788 <= int32(0) {
		goto L395
	} else {
		goto L396
	}
L364:
	;
	v1588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1546))))
	v1590 = v1588 & int32(127)
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+2)))
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+1)))
	v1593 = base.I32_extend8_s(v1588)
	if v1354 == int32(4) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	v1681 = int32(0)
	if v1681 < v1679 {
		goto L383
	} else {
		goto L384
	}
L366:
	;
	v1649 = v1549 + v1307
	v1650 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1649))) = uint8(v1650)
	if v1547 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L367:
	;
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+3)))
	v1599 = int32(8)
	v1605 = v1596 | (v1590<<(uint(int32(16))%32)|v1592<<(uint(v1599)%32)|v1591)<<(uint(v1599)%32)
	if v1593 < int32(0) {
		goto L370
	} else {
		goto L371
	}
L368:
	;
	goto L369
L369:
	;
	v1612 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+7)))
	v1613 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+6)))
	v1614 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+4)))
	v1615 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+3)))
	v1616 = int64(8)
	v1620 = int64(16)
	v1633 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1546)+5)))
	v1640 = v1612 | (v1613|((v1614|(v1615<<(uint(v1616)%64)|(base.I64_extend_i32_u(v1591)|(base.I64_extend_i32_u(v1590)<<(uint(v1620)%64)|base.I64_extend_i32_u(v1592)<<(uint(v1616)%64)))<<(uint(v1620)%64)))<<(uint(v1620)%64)|v1633<<(uint(v1616)%64)))<<(uint(v1616)%64)
	if v1593 < int32(0) {
		goto L373
	} else {
		goto L374
	}
L370:
	;
	v1610 = v1605 | int32(-2147483648)
	goto L372
L371:
	;
	v1610 = v1605
	goto L372
L372:
	;
	v1648 = base.I64_extend_i32_s(v1610)
	goto L366
L373:
	;
	v1645 = v1640 | int64(-9223372036854775807-1)
	goto L375
L374:
	;
	v1645 = v1640
	goto L375
L375:
	;
	v1648 = v1645
	goto L366
L376:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1305+v1667<<(uint(int32(3))%32)))) = v1648
	v1674 = v1546 + v1354
	v1675 = int32(1)
	v1678 = v1549 + v1675
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1678 < v1679 {
		v1546 = v1674
		v1547 = v1667 + v1675
		v1549 = v1678
		goto L364
	} else {
		goto L382
	}
L377:
	;
	v1667 = int32(0)
	goto L376
L378:
	;
	goto L379
L379:
	;
	v1656 = v1547 - int32(1)
	v1660 = *(*int64)(unsafe.Add(mBase, uint32(v1305+v1656<<(uint(int32(3))%32))))
	if v1660 < v1648 {
		v1667 = v1547
		goto L376
	} else {
		goto L380
	}
L380:
	;
	if v1648 < v1660 {
		v4412 = v1389
		goto L327
	} else {
		goto L381
	}
L381:
	;
	v1665 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1649-int32(1)))) = uint8(v1665)
	v1667 = v1656
	goto L376
L382:
	;
	goto L365
L383:
	;
	v1685 = v1674
	v1686 = v1681
	v1688 = v1681
	v1689 = v1679
	goto L386
L384:
	;
	v1744 = v1674
	v1747 = v1681
	goto L385
L385:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1787 = v1744
	v1788 = v1786
	v1790 = v1747
	goto L363
L386:
	;
	v1727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1685))))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1728 <= v1727 {
		v4412 = v1389
		goto L327
	} else {
		goto L388
	}
L387:
	;
	v1744 = v1740
	v1747 = v1737
	goto L385
L388:
	;
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1686+v1307))))
	if v1731 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1688+v1307))) = uint8(v1727)
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1737 = v1688 + int32(1)
	v1738 = v1734
	goto L391
L390:
	;
	v1737 = v1688
	v1738 = v1689
	goto L391
L391:
	;
	v1739 = int32(1)
	v1740 = v1685 + v1739
	v1742 = v1686 + v1739
	if v1742 < v1738 {
		v1685 = v1740
		v1686 = v1742
		v1688 = v1737
		v1689 = v1738
		goto L386
	} else {
		goto L392
	}
L392:
	;
	goto L387
L393:
	;
	v2166 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2133+v1301))) = uint8(v2166)
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2170 <= v2166 {
		goto L419
	} else {
		goto L420
	}
L394:
	;
	v1958 = int32(3)
	v1959 = v1925 & v1958
	if base.Ui32(v1925-int32(1)) < base.Ui32(v1958) {
		goto L408
	} else {
		goto L409
	}
L395:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if int32(0) < v1832 {
		v1916 = v1787
		v1925 = v1832
		goto L394
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v1836 = v1787
	v1850 = v1543
	goto L399
L398:
	;
	v2124 = v1787
	v2133 = int32(0)
	goto L393
L399:
	;
	v1880 = v1303 + v1850<<(uint(int32(4))%32)
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+3)))
	v1882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+2)))
	v1883 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1836))))
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+1)))
	v1889 = int32(8)
	v1895 = v1881 | (v1882|(v1883&int32(127)<<(uint(int32(16))%32)|v1888<<(uint(v1889)%32)))<<(uint(v1889)%32)
	if v1883 < int32(0) {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1916 = v1911
	v1925 = v1906
	goto L394
L401:
	;
	v1900 = v1895 | int32(-2147483648)
	goto L403
L402:
	;
	v1900 = v1895
	goto L403
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1880))) = v1900
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+4)))
	if base.Ui32(int32(1)) < base.Ui32(v1902) {
		v4412 = v1389
		goto L327
	} else {
		goto L404
	}
L404:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1880)+4)) = uint8(v1902)
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836)+5)))
	if v1906 <= v1907 {
		v4412 = v1389
		goto L327
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1880)+8)) = v1907
	v1911 = v1836 + int32(6)
	v1913 = v1850 + int32(1)
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1913 < v1914 {
		v1836 = v1911
		v1850 = v1913
		goto L399
	} else {
		goto L406
	}
L406:
	;
	goto L400
L407:
	;
	if v1959 == int32(0) {
		v2124 = v2028
		v2133 = v1925
		goto L393
	} else {
		goto L414
	}
L408:
	;
	v2028 = v1916
	v2029 = int32(0)
	goto L407
L409:
	;
	goto L410
L410:
	;
	v1968 = int32(0)
	v1970 = v1916
	v1971 = v1968
	v1976 = v1968
	goto L411
L411:
	;
	v2012 = v1971 + v1301
	v2013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1970))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2012))) = uint8(v2013)
	v2015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1970)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2012)+1)) = uint8(v2015)
	v2017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1970)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2012)+2)) = uint8(v2017)
	v2019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1970)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2012)+3)) = uint8(v2019)
	v2021 = int32(4)
	v2022 = v1971 + v2021
	v2024 = v1970 + v2021
	v2026 = v1976 + v2021
	if v2026 != v1925&int32(-4) {
		v1970 = v2024
		v1971 = v2022
		v1976 = v2026
		goto L411
	} else {
		goto L413
	}
L412:
	;
	v2028 = v2024
	v2029 = v2022
	goto L407
L413:
	;
	goto L412
L414:
	;
	v2072 = v2028
	v2073 = v2029
	v2076 = int32(0)
	goto L415
L415:
	;
	v2115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2072))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2073+v1301))) = uint8(v2115)
	v2117 = int32(1)
	v2120 = v2072 + v2117
	v2122 = v2076 + v2117
	if v2122 != v1959 {
		v2072 = v2120
		v2073 = v2073 + v2117
		v2076 = v2122
		goto L415
	} else {
		goto L417
	}
L416:
	;
	v2124 = v2120
	v2133 = v1925
	goto L393
L417:
	;
	goto L416
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2328
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v2367 {
		goto L441
	} else {
		goto L442
	}
L419:
	;
	v2324 = v2124
	v2328 = int32(0)
	goto L418
L420:
	;
	goto L421
L421:
	;
	v2175 = int32(0)
	v2177 = v2124
	v2181 = v2175
	v2183 = v2175
	v2218 = int64(0)
	goto L422
L422:
	;
	v2219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2177))))
	v2221 = v2219 & int32(127)
	v2222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+2)))
	v2223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+1)))
	v2224 = base.I32_extend8_s(v2219)
	if v1354 == int32(4) {
		goto L425
	} else {
		goto L426
	}
L423:
	;
	v2324 = v2313
	v2328 = v2321
	goto L418
L424:
	;
	if v2279 < int64(0) {
		v4412 = v1389
		goto L327
	} else {
		goto L434
	}
L425:
	;
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+3)))
	v2230 = int32(8)
	v2236 = v2227 | (v2221<<(uint(int32(16))%32)|v2223<<(uint(v2230)%32)|v2222)<<(uint(v2230)%32)
	if v2224 < int32(0) {
		goto L428
	} else {
		goto L429
	}
L426:
	;
	goto L427
L427:
	;
	v2243 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+7)))
	v2244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+6)))
	v2245 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+4)))
	v2246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+3)))
	v2247 = int64(8)
	v2251 = int64(16)
	v2264 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2177)+5)))
	v2271 = v2243 | (v2244|((v2245|(v2246<<(uint(v2247)%64)|(base.I64_extend_i32_u(v2222)|(base.I64_extend_i32_u(v2221)<<(uint(v2251)%64)|base.I64_extend_i32_u(v2223)<<(uint(v2247)%64)))<<(uint(v2251)%64)))<<(uint(v2251)%64)|v2264<<(uint(v2247)%64)))<<(uint(v2247)%64)
	if v2224 < int32(0) {
		goto L431
	} else {
		goto L432
	}
L428:
	;
	v2241 = v2236 | int32(-2147483648)
	goto L430
L429:
	;
	v2241 = v2236
	goto L430
L430:
	;
	v2279 = base.I64_extend_i32_s(v2241)
	goto L424
L431:
	;
	v2276 = v2271 | int64(-9223372036854775807-1)
	goto L433
L432:
	;
	v2276 = v2271
	goto L433
L433:
	;
	v2279 = v2276
	goto L424
L434:
	;
	if v2279-v2218 < int64(2419199) {
		v4412 = v1389
		goto L327
	} else {
		goto L435
	}
L435:
	;
	v2285 = v2177 + v1354
	v2286 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2285))))
	v2291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2285)+1)))
	v2292 = int32(8)
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2285)+2)))
	v2299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2285)+3)))
	v2300 = (v2286&int32(127)<<(uint(int32(16))%32)|v2291<<(uint(v2292)%32)|v2295)<<(uint(v2292)%32) | v2299
	if v2286 < int32(0) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v2305 = v2300 | int32(-2147483648)
	goto L438
L437:
	;
	v2305 = v2300
	goto L438
L438:
	;
	v2306 = int32(1)
	if base.B2i32(v2305 != v2183-v2306)&base.B2i32(v2305 != v2183+v2306) != 0 {
		v4412 = v1389
		goto L327
	} else {
		goto L439
	}
L439:
	;
	v2313 = v2177 + v1526
	v2316 = v1299 + v2181<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v2316)+8)) = base.I64_extend_i32_s(v2305)
	*(*int64)(unsafe.Add(mBase, uint32(v2316))) = v2279
	v2321 = v2181 + int32(1)
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2321 < v2322 {
		v2177 = v2313
		v2181 = v2321
		v2183 = v2305
		v2218 = v2279
		goto L422
	} else {
		goto L440
	}
L440:
	;
	goto L423
L441:
	;
	v2370 = v2324
	v2373 = v2166
	goto L444
L442:
	;
	v2489 = v2324
	goto L443
L443:
	;
	v2531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+4)))
	if v2531 != 0 {
		goto L460
	} else {
		goto L461
	}
L444:
	;
	v2412 = int32(0)
	if v1501 == v2412 {
		goto L447
	} else {
		goto L448
	}
L445:
	;
	v2430 = v2421
	v2431 = v2412
	goto L452
L446:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1303+v2373<<(uint(int32(4))%32))+12)) = uint8(v2422)
	v2428 = v2373 + int32(1)
	if v2428 != v2367 {
		v2370 = v2421
		v2373 = v2428
		goto L444
	} else {
		goto L451
	}
L447:
	;
	v2421 = v2370
	v2422 = int32(0)
	goto L446
L448:
	;
	goto L449
L449:
	;
	v2416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2370))))
	if base.Ui32(int32(1)) < base.Ui32(v2416) {
		v4412 = v1389
		goto L327
	} else {
		goto L450
	}
L450:
	;
	v2421 = v2370 + int32(1)
	v2422 = v2416
	goto L446
L451:
	;
	goto L445
L452:
	;
	if v1520 == int32(0) {
		goto L455
	} else {
		goto L456
	}
L453:
	;
	v2489 = v2480
	goto L443
L454:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1303+v2431<<(uint(int32(4))%32))+13)) = uint8(v2481)
	v2487 = v2431 + int32(1)
	if v2487 != v2367 {
		v2430 = v2480
		v2431 = v2487
		goto L452
	} else {
		goto L459
	}
L455:
	;
	v2480 = v2430
	v2481 = int32(0)
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2430))))
	if base.Ui32(int32(1)) < base.Ui32(v2475) {
		v4412 = v1389
		goto L327
	} else {
		goto L458
	}
L458:
	;
	v2480 = v2430 + int32(1)
	v2481 = v2475
	goto L454
L459:
	;
	goto L453
L460:
	;
	v2533 = v44 - v2489 + v1366
	if v44 == v2489 {
		goto L464
	} else {
		goto L465
	}
L461:
	;
	v2684 = v1366
	goto L462
L462:
	;
	goto L329
L463:
	;
	if base.Ui32(v1354) < base.Ui32(int32(5)) {
		v1354 = v1354 << (uint(int32(1)) % 32)
		v1366 = v2533
		goto L328
	} else {
		goto L509
	}
L464:
	;
	goto L463
L465:
	;
	v2537 = v44 + v2533
	if base.Ui32(v2489-v2537) <= base.Ui32(int32(0)-v2533<<(uint(int32(1))%32)) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	v2544 = F___memcpy(m, v44, v2489, v2533)
	mBase = m.M
	goto L463
L467:
	;
	goto L468
L468:
	;
	v2547 = (v44 ^ v2489) & int32(3)
	if base.Ui32(v44) < base.Ui32(v2489) {
		goto L471
	} else {
		goto L472
	}
L469:
	;
	if v2649 == int32(0) {
		goto L464
	} else {
		goto L505
	}
L470:
	;
	if base.Ui32(v2627) <= base.Ui32(int32(3)) {
		v2648 = v2626
		v2649 = v2627
		v2650 = v2628
		goto L469
	} else {
		goto L501
	}
L471:
	;
	if v2547 != 0 {
		goto L474
	} else {
		goto L475
	}
L472:
	;
	goto L473
L473:
	;
	if v2547 != 0 {
		v2609 = v2533
		goto L484
	} else {
		goto L485
	}
L474:
	;
	v2648 = v2489
	v2649 = v2533
	v2650 = v44
	goto L469
L475:
	;
	goto L476
L476:
	;
	if v44&int32(3) == int32(0) {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v2626 = v2489
	v2627 = v2533
	v2628 = v44
	goto L470
L478:
	;
	goto L479
L479:
	;
	v2554 = v2489
	v2555 = v2533
	v2556 = v44
	goto L480
L480:
	;
	if v2555 == int32(0) {
		goto L464
	} else {
		goto L482
	}
L481:
	;
	v2626 = v2563
	v2627 = v2565
	v2628 = v2567
	goto L470
L482:
	;
	v2560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2554))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2556))) = uint8(v2560)
	v2562 = int32(1)
	v2563 = v2554 + v2562
	v2565 = v2555 - v2562
	v2567 = v2556 + v2562
	if v2567&int32(3) != 0 {
		v2554 = v2563
		v2555 = v2565
		v2556 = v2567
		goto L480
	} else {
		goto L483
	}
L483:
	;
	goto L481
L484:
	;
	if v2609 == int32(0) {
		goto L464
	} else {
		goto L497
	}
L485:
	;
	if v2537&int32(3) != 0 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v2574 = v2533
	goto L489
L487:
	;
	v2589 = v2533
	goto L488
L488:
	;
	if base.Ui32(v2589) <= base.Ui32(int32(3)) {
		v2609 = v2589
		goto L484
	} else {
		goto L493
	}
L489:
	;
	if v2574 == int32(0) {
		goto L464
	} else {
		goto L491
	}
L490:
	;
	v2589 = v2580
	goto L488
L491:
	;
	v2580 = v2574 - int32(1)
	v2581 = v44 + v2580
	v2583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2489+v2580))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2581))) = uint8(v2583)
	if v2581&int32(3) != 0 {
		v2574 = v2580
		goto L489
	} else {
		goto L492
	}
L492:
	;
	goto L490
L493:
	;
	v2596 = v2589
	goto L494
L494:
	;
	v2600 = v2596 - int32(4)
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v2489+v2600)))
	*(*int32)(unsafe.Add(mBase, uint32(v44+v2600))) = v2603
	if base.Ui32(int32(3)) < base.Ui32(v2600) {
		v2596 = v2600
		goto L494
	} else {
		goto L496
	}
L495:
	;
	v2609 = v2600
	goto L484
L496:
	;
	goto L495
L497:
	;
	v2616 = v2609
	goto L498
L498:
	;
	v2620 = v2616 - int32(1)
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2489+v2620))))
	*(*uint8)(unsafe.Add(mBase, uint32(v44+v2620))) = uint8(v2623)
	if v2620 != 0 {
		v2616 = v2620
		goto L498
	} else {
		goto L500
	}
L499:
	;
	goto L464
L500:
	;
	goto L499
L501:
	;
	v2633 = v2626
	v2634 = v2627
	v2635 = v2628
	goto L502
L502:
	;
	v2637 = *(*int32)(unsafe.Add(mBase, uint32(v2633)))
	*(*int32)(unsafe.Add(mBase, uint32(v2635))) = v2637
	v2639 = int32(4)
	v2640 = v2633 + v2639
	v2642 = v2635 + v2639
	v2644 = v2634 - v2639
	if base.Ui32(int32(3)) < base.Ui32(v2644) {
		v2633 = v2640
		v2634 = v2644
		v2635 = v2642
		goto L502
	} else {
		goto L504
	}
L503:
	;
	v2648 = v2640
	v2649 = v2644
	v2650 = v2642
	goto L469
L504:
	;
	goto L503
L505:
	;
	v2655 = v2648
	v2656 = v2649
	v2657 = v2650
	goto L506
L506:
	;
	v2659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2655))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2657))) = uint8(v2659)
	v2661 = int32(1)
	v2666 = v2656 - v2661
	if v2666 != 0 {
		v2655 = v2655 + v2661
		v2656 = v2666
		v2657 = v2657 + v2661
		goto L506
	} else {
		goto L508
	}
L507:
	;
	goto L464
L508:
	;
	goto L507
L509:
	;
	v2684 = v2533
	goto L462
L510:
	;
	v3819 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3819 < int32(2) {
		goto L620
	} else {
		goto L621
	}
L511:
	;
	v3774 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v3774 == int32(0) {
		v4412 = v1389
		goto L327
	} else {
		goto L619
	}
L512:
	;
	v2687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v2687 != int32(10) {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v2692 = v44 + v2684 - int32(1)
	v2693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2692))))
	if v2693 != int32(10) {
		goto L511
	} else {
		goto L514
	}
L514:
	;
	if int32(256) < v2367+int32(2) {
		v3783 = v2367
		goto L510
	} else {
		goto L515
	}
L515:
	;
	v2700 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2692))) = uint8(v2700)
	v2708 = F_tzparse(m, v44+int32(1), v44+int32(_a_F_tzload_4), v2700)
	mBase = m.M
	if v2708 == v2700 {
		goto L511
	} else {
		goto L516
	}
L516:
	;
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_tzload[3])))
	if int32(0) < v2712 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2723 = v2711
	v2724 = int32(0)
	v2728 = v2700
	goto L520
L518:
	;
	v3116 = v2711
	v3121 = v2700
	goto L519
L519:
	;
	if v3121 != v2712 {
		goto L511
	} else {
		goto L581
	}
L520:
	;
	v2764 = v44 + int32(_a_F_tzload_9) + v2724<<(uint(int32(4))%32)
	v2765 = *(*int32)(unsafe.Add(mBase, uint32(v2764)))
	v2766 = v44 + int32(_a_F_tzload_10) + v2765
	v2767 = int32(0)
	if v2767 < v2723 {
		goto L524
	} else {
		goto L525
	}
L521:
	;
	v3116 = v3071
	v3121 = v3076
	goto L519
L522:
	;
	v3111 = v2724 + int32(1)
	if v3111 < v2712 {
		v2723 = v3071
		v2724 = v3111
		v2728 = v3076
		goto L520
	} else {
		goto L580
	}
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2764))) = v3030
	v3071 = v3026
	v3076 = v2728 + int32(1)
	goto L522
L524:
	;
	v2771 = v2767
	goto L527
L525:
	;
	v2851 = v2767
	goto L526
L526:
	;
	if v2766&int32(3) == int32(0) {
		v2909 = v2766
		goto L543
	} else {
		goto L544
	}
L527:
	;
	v2813 = v2771 + v1301
	v2816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2766))))
	v2817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2813))))
	if v2817 == int32(0) {
		v2836 = v2816
		v2837 = v2817
		goto L530
	} else {
		goto L531
	}
L528:
	;
	v2851 = v2723
	goto L526
L529:
	;
	if v2837-v2836 == int32(0) {
		goto L537
	} else {
		goto L538
	}
L530:
	;
	goto L529
L531:
	;
	if v2816 != v2817 {
		v2836 = v2816
		v2837 = v2817
		goto L530
	} else {
		goto L532
	}
L532:
	;
	v2821 = v2813
	v2822 = v2766
	goto L533
L533:
	;
	v2825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2822)+1)))
	v2826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2821)+1)))
	if v2826 == int32(0) {
		v2836 = v2825
		v2837 = v2826
		goto L530
	} else {
		goto L535
	}
L534:
	;
	v2836 = v2825
	v2837 = v2826
	goto L530
L535:
	;
	v2829 = int32(1)
	if v2825 == v2826 {
		v2821 = v2821 + v2829
		v2822 = v2822 + v2829
		goto L533
	} else {
		goto L536
	}
L536:
	;
	goto L534
L537:
	;
	v3026 = v2723
	v3030 = v2771
	goto L523
L538:
	;
	goto L539
L539:
	;
	v2842 = v2771 + int32(1)
	if v2842 != v2723 {
		v2771 = v2842
		goto L527
	} else {
		goto L540
	}
L540:
	;
	goto L528
L541:
	;
	v2943 = v2942 + v2851
	if int32(49) < v2943 {
		v3071 = v2723
		v3076 = v2728
		goto L522
	} else {
		goto L558
	}
L542:
	;
	v2942 = v2934 - v2766
	goto L541
L543:
	;
	v2913 = v2909
	goto L552
L544:
	;
	v2893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2766))))
	if v2893 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v2942 = int32(0)
	goto L541
L546:
	;
	goto L547
L547:
	;
	v2898 = v2766
	goto L548
L548:
	;
	v2902 = v2898 + int32(1)
	if v2902&int32(3) == int32(0) {
		v2909 = v2902
		goto L543
	} else {
		goto L550
	}
L549:
	;
	v2934 = v2902
	goto L542
L550:
	;
	v2907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2902))))
	if v2907 != 0 {
		v2898 = v2902
		goto L548
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v2913)))
	v2922 = int32(-2139062144)
	if (int32(16843008)-v2919|v2919)&v2922 == v2922 {
		v2913 = v2913 + int32(4)
		goto L552
	} else {
		goto L554
	}
L553:
	;
	v2928 = v2913
	goto L555
L554:
	;
	goto L553
L555:
	;
	v2932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2928))))
	if v2932 != 0 {
		v2928 = v2928 + int32(1)
		goto L555
	} else {
		goto L557
	}
L556:
	;
	v2934 = v2928
	goto L542
L557:
	;
	goto L556
L558:
	;
	v2946 = v2851 + v1301
	if (v2766^v2946)&int32(3) != 0 {
		goto L562
	} else {
		goto L563
	}
L559:
	;
	v3026 = v2943 + int32(1)
	v3030 = v2851
	goto L523
L560:
	;
	goto L559
L561:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3001))) = uint8(v3000)
	if v3000&int32(255) == int32(0) {
		goto L560
	} else {
		goto L576
	}
L562:
	;
	v2952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2766))))
	v2999 = v2766
	v3000 = v2952
	v3001 = v2946
	goto L561
L563:
	;
	goto L564
L564:
	;
	if v2766&int32(3) != 0 {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	v2956 = v2766
	v2958 = v2946
	goto L568
L566:
	;
	v2970 = v2766
	v2972 = v2946
	goto L567
L567:
	;
	v2974 = *(*int32)(unsafe.Add(mBase, uint32(v2970)))
	v2977 = int32(-2139062144)
	if (int32(16843008)-v2974|v2974)&v2977 != v2977 {
		v2999 = v2970
		v3000 = v2974
		v3001 = v2972
		goto L561
	} else {
		goto L572
	}
L568:
	;
	v2959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2956))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2958))) = uint8(v2959)
	if v2959 == int32(0) {
		goto L560
	} else {
		goto L570
	}
L569:
	;
	v2970 = v2966
	v2972 = v2964
	goto L567
L570:
	;
	v2963 = int32(1)
	v2964 = v2958 + v2963
	v2966 = v2956 + v2963
	if v2966&int32(3) != 0 {
		v2956 = v2966
		v2958 = v2964
		goto L568
	} else {
		goto L571
	}
L571:
	;
	goto L569
L572:
	;
	v2982 = v2970
	v2983 = v2974
	v2984 = v2972
	goto L573
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2984))) = v2983
	v2986 = int32(4)
	v2987 = v2984 + v2986
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2982)+4))
	v2990 = v2982 + v2986
	v2994 = int32(-2139062144)
	if (v2988|(int32(16843008)-v2988))&v2994 == v2994 {
		v2982 = v2990
		v2983 = v2988
		v2984 = v2987
		goto L573
	} else {
		goto L575
	}
L574:
	;
	v2999 = v2990
	v3000 = v2988
	v3001 = v2987
	goto L561
L575:
	;
	goto L574
L576:
	;
	v3008 = v2999
	v3010 = v3001
	goto L577
L577:
	;
	v3011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3008)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3010)+1)) = uint8(v3011)
	v3013 = int32(1)
	if v3011 != 0 {
		v3008 = v3008 + v3013
		v3010 = v3010 + v3013
		goto L577
	} else {
		goto L579
	}
L578:
	;
	goto L560
L579:
	;
	goto L578
L580:
	;
	goto L521
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v3116
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3157 < int32(2) {
		v3218 = v3157
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_tzload[4])))
	v3258 = int32(0)
	if v3218 == v3258 {
		v3385 = v3258
		goto L589
	} else {
		goto L590
	}
L583:
	;
	v3163 = v3157
	goto L584
L584:
	;
	v3203 = v3163 - int32(1)
	v3205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307+v3203))))
	v3209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3163+v1307-int32(2)))))
	if v3205 != v3209 {
		v3218 = v3163
		goto L582
	} else {
		goto L586
	}
L585:
	;
	v3218 = int32(1)
	goto L582
L586:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v3203
	if base.Ui32(int32(2)) < base.Ui32(v3163) {
		v3163 = v3203
		goto L584
	} else {
		goto L587
	}
L587:
	;
	goto L585
L588:
	;
	if v2712 <= int32(0) {
		goto L511
	} else {
		goto L611
	}
L589:
	;
	if v3257 <= v3385 {
		goto L588
	} else {
		goto L601
	}
L590:
	;
	if v3257 <= int32(0) {
		v3385 = v3258
		goto L589
	} else {
		goto L591
	}
L591:
	;
	v3270 = *(*int64)(unsafe.Add(mBase, uint32(v3218<<(uint(int32(3))%32)+v1305-int32(8))))
	v3271 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v3279 = v3258
	goto L592
L592:
	;
	v3317 = *(*int64)(unsafe.Add(mBase, uint32(v44+int32(_a_F_tzload_11)+v3279<<(uint(int32(3))%32))))
	v3318 = v3271
	goto L595
L593:
	;
	goto L588
L594:
	;
	if v3270 < v3317+v3372 {
		v3385 = v3279
		goto L589
	} else {
		goto L599
	}
L595:
	;
	v3362 = v3318 - int32(1)
	if v3362 < int32(0) {
		v3372 = int64(0)
		goto L594
	} else {
		goto L597
	}
L596:
	;
	v3370 = *(*int64)(unsafe.Add(mBase, uint32(v3367)+8))
	v3372 = v3370
	goto L594
L597:
	;
	v3367 = v1299 + v3362<<(uint(int32(4))%32)
	v3368 = *(*int64)(unsafe.Add(mBase, uint32(v3367)))
	if v3317 < v3368 {
		v3318 = v3362
		goto L595
	} else {
		goto L598
	}
L598:
	;
	goto L596
L599:
	;
	v3376 = v3279 + int32(1)
	if v3376 != v3257 {
		v3279 = v3376
		goto L592
	} else {
		goto L600
	}
L600:
	;
	goto L593
L601:
	;
	v3428 = v3218
	v3432 = v3385
	goto L602
L602:
	;
	if int32(1999) < v3428 {
		goto L588
	} else {
		goto L604
	}
L603:
	;
	goto L588
L604:
	;
	v3469 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v3470 = int32(3)
	v3476 = *(*int64)(unsafe.Add(mBase, uint32(v44+int32(_a_F_tzload_11)+v3432<<(uint(v3470)%32))))
	v3477 = v3469
	goto L606
L605:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1305+v3428<<(uint(v3470)%32)))) = v3476 + v3531
	v3534 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(_a_F_tzload_12)+v3432))))
	v3538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v3539 = v3537 + v3538
	*(*uint8)(unsafe.Add(mBase, uint32(v1307+v3534))) = uint8(v3539)
	v3541 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3542 = int32(1)
	v3543 = v3541 + v3542
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v3543
	v3546 = v3432 + v3542
	if v3546 != v3257 {
		v3428 = v3543
		v3432 = v3546
		goto L602
	} else {
		goto L610
	}
L606:
	;
	v3521 = v3477 - int32(1)
	if v3521 < int32(0) {
		v3531 = int64(0)
		goto L605
	} else {
		goto L608
	}
L607:
	;
	v3529 = *(*int64)(unsafe.Add(mBase, uint32(v3526)+8))
	v3531 = v3529
	goto L605
L608:
	;
	v3526 = v1299 + v3521<<(uint(int32(4))%32)
	v3527 = *(*int64)(unsafe.Add(mBase, uint32(v3526)))
	if v3476 < v3527 {
		v3477 = v3521
		goto L606
	} else {
		goto L609
	}
L609:
	;
	goto L607
L610:
	;
	goto L603
L611:
	;
	v3592 = int32(1)
	v3595 = v44 + int32(_a_F_tzload_13)
	v3596 = int32(0)
	if v2712 != v3592 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v3603 = v3596
	v3605 = int32(0)
	goto L615
L613:
	;
	v3675 = v3596
	goto L614
L614:
	;
	if v2712&v3592 == int32(0) {
		goto L511
	} else {
		goto L618
	}
L615:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3645 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v3644 + v3645
	v3648 = int32(4)
	v3650 = v1303 + v3644<<(uint(v3648)%32)
	v3653 = v3595 + v3603<<(uint(v3648)%32)
	v3654 = *(*int64)(unsafe.Add(mBase, uint32(v3653)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3650)+8)) = v3654
	v3656 = *(*int64)(unsafe.Add(mBase, uint32(v3653)))
	*(*int64)(unsafe.Add(mBase, uint32(v3650))) = v3656
	v3658 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v3658 + v3645
	v3664 = v1303 + v3658<<(uint(v3648)%32)
	v3665 = *(*int64)(unsafe.Add(mBase, uint32(v3653)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3664)+8)) = v3665
	v3667 = *(*int64)(unsafe.Add(mBase, uint32(v3653)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3664))) = v3667
	v3669 = int32(2)
	v3670 = v3603 + v3669
	v3672 = v3605 + v3669
	if v3672 != v2712&int32(2147483646) {
		v3603 = v3670
		v3605 = v3672
		goto L615
	} else {
		goto L617
	}
L616:
	;
	v3675 = v3670
	goto L614
L617:
	;
	goto L616
L618:
	;
	v3718 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v3718 + int32(1)
	v3722 = int32(4)
	v3724 = v1303 + v3718<<(uint(v3722)%32)
	v3727 = v3595 + v3675<<(uint(v3722)%32)
	v3728 = *(*int64)(unsafe.Add(mBase, uint32(v3727)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3724)+8)) = v3728
	v3730 = *(*int64)(unsafe.Add(mBase, uint32(v3727)))
	*(*int64)(unsafe.Add(mBase, uint32(v3724))) = v3730
	goto L511
L619:
	;
	v3783 = v3774
	goto L510
L620:
	;
	v4092 = int32(0)
	if v3819 <= v4092 {
		v4343 = v4092
		goto L663
	} else {
		goto L664
	}
L621:
	;
	v3822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_tzload[5]))))
	v3825 = v1303 + v3822<<(uint(int32(4))%32)
	v3827 = int32(1)
	goto L622
L622:
	;
	if v3783 <= v3822 {
		goto L625
	} else {
		goto L626
	}
L623:
	;
	v3936 = v3819 - int32(1)
	v3941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3936+v1307))))
	v3944 = v1303 + v3941<<(uint(int32(4))%32)
	v3946 = v3819 - int32(2)
	goto L643
L624:
	;
	goto L623
L625:
	;
	v3929 = v3827 + int32(1)
	if v3929 != v3819 {
		v3827 = v3929
		goto L622
	} else {
		goto L642
	}
L626:
	;
	v3871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3827+v1307))))
	if v3783 <= v3871 {
		goto L625
	} else {
		goto L627
	}
L627:
	;
	v3875 = v1303 + v3871<<(uint(int32(4))%32)
	v3876 = *(*int32)(unsafe.Add(mBase, uint32(v3875)))
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3825)))
	if v3876 != v3877 {
		goto L625
	} else {
		goto L628
	}
L628:
	;
	v3879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3875)+4)))
	v3880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3825)+4)))
	if v3879 != v3880 {
		goto L625
	} else {
		goto L629
	}
L629:
	;
	v3882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3875)+12)))
	v3883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3825)+12)))
	if v3882 != v3883 {
		goto L625
	} else {
		goto L630
	}
L630:
	;
	v3885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3875)+13)))
	v3886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3825)+13)))
	if v3885 != v3886 {
		goto L625
	} else {
		goto L631
	}
L631:
	;
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v3875)+8))
	v3889 = v1301 + v3888
	v3890 = *(*int32)(unsafe.Add(mBase, uint32(v3825)+8))
	v3891 = v1301 + v3890
	v3894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3891))))
	v3895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3889))))
	if v3895 == int32(0) {
		v3914 = v3894
		v3915 = v3895
		goto L633
	} else {
		goto L634
	}
L632:
	;
	if v3915-v3914 != 0 {
		goto L625
	} else {
		goto L640
	}
L633:
	;
	goto L632
L634:
	;
	if v3894 != v3895 {
		v3914 = v3894
		v3915 = v3895
		goto L633
	} else {
		goto L635
	}
L635:
	;
	v3899 = v3889
	v3900 = v3891
	goto L636
L636:
	;
	v3903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3900)+1)))
	v3904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3899)+1)))
	if v3904 == int32(0) {
		v3914 = v3903
		v3915 = v3904
		goto L633
	} else {
		goto L638
	}
L637:
	;
	v3914 = v3903
	v3915 = v3904
	goto L633
L638:
	;
	v3907 = int32(1)
	if v3903 == v3904 {
		v3899 = v3899 + v3907
		v3900 = v3900 + v3907
		goto L636
	} else {
		goto L639
	}
L639:
	;
	goto L637
L640:
	;
	v3920 = *(*int64)(unsafe.Add(mBase, uint32(v1305+v3827<<(uint(int32(3))%32))))
	v3921 = *(*int64)(unsafe.Add(mBase, uint32(v1305)))
	if v3920-v3921 != int64(12622780800) {
		goto L625
	} else {
		goto L641
	}
L641:
	;
	v3925 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v3925)
	goto L624
L642:
	;
	goto L624
L643:
	;
	if v3783 <= v3941 {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	goto L620
L645:
	;
	if int32(0) < v3946 {
		v3946 = v3946 - int32(1)
		goto L643
	} else {
		goto L662
	}
L646:
	;
	v3989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3946+v1307))))
	if v3783 <= v3989 {
		goto L645
	} else {
		goto L647
	}
L647:
	;
	v3991 = *(*int32)(unsafe.Add(mBase, uint32(v3944)))
	v3994 = v1303 + v3989<<(uint(int32(4))%32)
	v3995 = *(*int32)(unsafe.Add(mBase, uint32(v3994)))
	if v3991 != v3995 {
		goto L645
	} else {
		goto L648
	}
L648:
	;
	v3997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3944)+4)))
	v3998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3994)+4)))
	if v3997 != v3998 {
		goto L645
	} else {
		goto L649
	}
L649:
	;
	v4000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3944)+12)))
	v4001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3994)+12)))
	if v4000 != v4001 {
		goto L645
	} else {
		goto L650
	}
L650:
	;
	v4003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3944)+13)))
	v4004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3994)+13)))
	if v4003 != v4004 {
		goto L645
	} else {
		goto L651
	}
L651:
	;
	v4006 = *(*int32)(unsafe.Add(mBase, uint32(v3944)+8))
	v4007 = v1301 + v4006
	v4008 = *(*int32)(unsafe.Add(mBase, uint32(v3994)+8))
	v4009 = v1301 + v4008
	v4012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4009))))
	v4013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4007))))
	if v4013 == int32(0) {
		v4032 = v4012
		v4033 = v4013
		goto L653
	} else {
		goto L654
	}
L652:
	;
	if v4033-v4032 != 0 {
		goto L645
	} else {
		goto L660
	}
L653:
	;
	goto L652
L654:
	;
	if v4012 != v4013 {
		v4032 = v4012
		v4033 = v4013
		goto L653
	} else {
		goto L655
	}
L655:
	;
	v4017 = v4007
	v4018 = v4009
	goto L656
L656:
	;
	v4021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4018)+1)))
	v4022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4017)+1)))
	if v4022 == int32(0) {
		v4032 = v4021
		v4033 = v4022
		goto L653
	} else {
		goto L658
	}
L657:
	;
	v4032 = v4021
	v4033 = v4022
	goto L653
L658:
	;
	v4025 = int32(1)
	if v4021 == v4022 {
		v4017 = v4017 + v4025
		v4018 = v4018 + v4025
		goto L656
	} else {
		goto L659
	}
L659:
	;
	goto L657
L660:
	;
	v4035 = *(*int64)(unsafe.Add(mBase, uint32(v1305+v3936<<(uint(int32(3))%32))))
	v4039 = *(*int64)(unsafe.Add(mBase, uint32(v1305+v3946<<(uint(int32(3))%32))))
	if v4035-v4039 != int64(12622780800) {
		goto L645
	} else {
		goto L661
	}
L661:
	;
	v4043 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+17)) = uint8(v4043)
	goto L620
L662:
	;
	goto L644
L663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_tzload[6]))) = v4343
	v4412 = v4092
	goto L327
L664:
	;
	v4096 = v4092
	goto L666
L665:
	;
	v4343 = int32(0)
	goto L663
L666:
	;
	v4139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4096+v1307))))
	if v4139 != 0 {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	v4143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_tzload[5]))))
	v4147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303+v4143<<(uint(int32(4))%32))+4)))
	if v4147 != int32(1) {
		goto L672
	} else {
		goto L673
	}
L668:
	;
	v4141 = v4096 + int32(1)
	if v3819 != v4141 {
		v4096 = v4141
		goto L666
	} else {
		goto L671
	}
L669:
	;
	goto L670
L670:
	;
	goto L667
L671:
	;
	goto L665
L672:
	;
	v4242 = int32(1)
	if v3783 <= v4242 {
		goto L678
	} else {
		goto L679
	}
L673:
	;
	v4150 = v4143
	goto L674
L674:
	;
	if v4150 <= int32(0) {
		goto L672
	} else {
		goto L676
	}
L675:
	;
	v4343 = v4195
	goto L663
L676:
	;
	v4195 = v4150 - int32(1)
	v4199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303+v4195<<(uint(int32(4))%32))+4)))
	if v4199 != 0 {
		v4150 = v4195
		goto L674
	} else {
		goto L677
	}
L677:
	;
	goto L675
L678:
	;
	v4245 = v4242
	goto L680
L679:
	;
	v4245 = v3783
	goto L680
L680:
	;
	v4249 = int32(0)
	goto L681
L681:
	;
	v4294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(_a_F_tzload_14)+v4249<<(uint(int32(4))%32)))))
	if v4294 != int32(1) {
		v4343 = v4249
		goto L663
	} else {
		goto L683
	}
L682:
	;
	goto L665
L683:
	;
	v4298 = v4249 + int32(1)
	if v4298 != v4245 {
		v4249 = v4298
		goto L681
	} else {
		goto L684
	}
L684:
	;
	goto L682
}
