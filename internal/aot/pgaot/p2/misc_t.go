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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[821]))
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[821]))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v41 = v36 + v27&int32(15)<<(uint(int32(7))%32)
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[824]))
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
	v61 = int32(25344)
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
	v48 = *(*int32)(unsafe.Add(mBase, _consts[821]))
	v51 = *(*int32)(unsafe.Add(mBase, _consts[825]))
	v55 = F_hash_search_with_hash_value(m, v48, int32(1629180), v51, int32(2), v23+int32(13))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[824]))
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
	v87 = *(*int32)(unsafe.Add(mBase, _consts[821]))
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
	v94 = *(*int32)(unsafe.Add(mBase, _consts[821]))
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
	v111 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v354 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v165 = *(*int32)(unsafe.Add(mBase, _consts[822]))
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
	v181 = *(*int32)(unsafe.Add(mBase, _consts[822]))
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
	v195 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v201 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v281 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v245 = *(*int32)(unsafe.Add(mBase, _consts[822]))
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
	v292 = *(*int32)(unsafe.Add(mBase, _consts[821]))
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
	v366 = *(*int32)(unsafe.Add(mBase, _consts[821]))
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
	v426 = *(*int32)(unsafe.Add(mBase, _consts[824]))
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
	v431 = *(*int32)(unsafe.Add(mBase, _consts[821]))
	v434 = *(*int32)(unsafe.Add(mBase, _consts[825]))
	v438 = F_hash_search_with_hash_value(m, v431, int32(1629180), v434, int32(1), v23+int32(15))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[824]))
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
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1005])))
				if v17 != int32(1) {
					F_char2wchar(m, v6+int32(4), int32(3), l0, v12, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						if base.Ui32(v35) <= base.Ui32(int32(131071)) {
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(8))%32)))+uint32(_consts[602]))))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v35)>>(uint(int32(3))%32))&int32(31)|v46<<(uint(int32(5))%32))+uint32(_consts[602]))))
							v60 = int32(base.Ui32(v52)>>(uint(v35&int32(7))%32)) & int32(1)
						} else {
							v60 = base.B2i32(base.Ui32(v35) < base.Ui32(int32(196606)))
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
		F_appendStringInfo(m, l0, int32(697661), v7)
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
		F_appendStringInfo(m, l0, int32(59441), v7+int32(16))
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(12))+uint32(_consts[73])))
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[67]))
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
	v29 = *(*int32)(unsafe.Add(mBase, _consts[68]))
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
	v33 = *(*int32)(unsafe.Add(mBase, _consts[65]))
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
	v95 = *(*int32)(unsafe.Add(mBase, _consts[69]))
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
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v140)%32))+uint32(_consts[71])))
	v145 = int32(12)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144*v145)+uint32(_consts[66])))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(v140)%32))+uint32(_consts[71])))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154*v145)+uint32(_consts[66])))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v149<<(uint(v140)%32))+uint32(_consts[74])))
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
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v183)%32))+uint32(_consts[71])))
	v188 = int32(12)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187*v188)+uint32(_consts[66])))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v13<<(uint(v183)%32))+uint32(_consts[71])))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197*v188)+uint32(_consts[66])))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v192<<(uint(v183)%32))+uint32(_consts[74])))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_strlen(m, v5)
	mBase = m.M
	v8 = v6 + int32(4)
	v9 = F_palloc(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v8 << (uint(int32(2)) % 32)
		if v6 != 0 {
			v18 = F__emscripten_memcpy_bulkmem(m, v9+int32(4), v5, v6)
			mBase = m.M
		} else {
		}
		return v9
	}
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1559), v3, v4, v5)
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
	F_errmsg(m, int32(245508), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(574955), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(500260), int32(1648), int32(106279))
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
	v42 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	v50 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
	v63 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1063]))
	v20 = F_pg_localtime(m, v5+int32(8), v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		v24 = F_pg_strftime(m, v5+int32(144), int32(128), int32(508343), v20)
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
						F_errmsg(m, int32(402548), int32(0))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(496005), int32(6550), int32(237178))
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
								F_errmsg(m, int32(402548), int32(0))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
								v52 = int32(4800)
							} else {
								v52 = int32(4799)
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
							v73 = base.I32_div_s((v68+v43)*int32(7834), int32(256))
							v79 = base.I64_extend_i32_s(v46 + v53*int32(365) + v58 + v61 + v64 + v73 - int32(32167) - int32(2451545))
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
										F_errmsg(m, int32(402548), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
											F_errmsg(m, int32(402548), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
												F_errmsg(m, int32(402548), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
									F_errmsg(m, int32(402548), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
					if v27 <= int32(5874897) {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
						v43 = v37
						v45 = v8 + int32(8)
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						v51 = base.B2i32(int32(2) < v43)
						if int32(2) < v43 {
							v52 = int32(4800)
						} else {
							v52 = int32(4799)
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
						v73 = base.I32_div_s((v68+v43)*int32(7834), int32(256))
						v79 = base.I64_extend_i32_s(v46 + v53*int32(365) + v58 + v61 + v64 + v73 - int32(32167) - int32(2451545))
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
									F_errmsg(m, int32(402548), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
										F_errmsg(m, int32(402548), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
											F_errmsg(m, int32(402548), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
						if v27 != int32(5874898) {
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
									F_errmsg(m, int32(402548), int32(0))
									mBase = m.M
									v184 = m.ExcPending
									if v184 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
										F_errmsg(m, int32(402548), int32(0))
										mBase = m.M
										v184 = m.ExcPending
										if v184 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
									v52 = int32(4800)
								} else {
									v52 = int32(4799)
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
								v73 = base.I32_div_s((v68+v43)*int32(7834), int32(256))
								v79 = base.I64_extend_i32_s(v46 + v53*int32(365) + v58 + v61 + v64 + v73 - int32(32167) - int32(2451545))
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
											F_errmsg(m, int32(402548), int32(0))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
												F_errmsg(m, int32(402548), int32(0))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
													F_errmsg(m, int32(402548), int32(0))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(496005), int32(6554), int32(237178))
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
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(372721)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
		v14 = F_psprintf(m, int32(177326), v5)
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
		v19 = F_pstrdup(m, int32(372721))
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
		F_errcontext_msg(m, int32(716109), v5)
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
								F_errmsg(m, int32(347795), int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498600), int32(74), int32(86268))
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
							F_errmsg(m, int32(347795), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498600), int32(74), int32(86268))
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
	F_errcode(m, int32(393348))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(146559), int32(0))
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
	F_errmsg_internal(m, int32(339273), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(498825), int32(3187), int32(265223))
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
	F_errfinish(m, int32(498825), int32(3120), int32(359924))
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
	F_errcode(m, int32(393348))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(146559), int32(0))
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
	F_errfinish(m, int32(498825), int32(3149), int32(359924))
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
	v63 = int32(131064)
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
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(357586)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(993)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v19
		v23 = int32(4508392)
		v24 = *(*int32)(unsafe.Add(mBase, _consts[88]))
		*(*int32)(unsafe.Add(mBase, _consts[88])) = v12 + int32(4)
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
			*(*int32)(unsafe.Add(mBase, _consts[88])) = v44
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
	v100 = *(*float64)(unsafe.Add(mBase, _consts[485]))
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
	v113 = *(*float64)(unsafe.Add(mBase, _consts[489]))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	goto L41
L41:
	;
	goto L42
L42:
	;
	v140 = base.F64_add(base.F64_mul(v105, v98), base.F64_add(base.F64_sub(v107, v108), v110))
	v141 = base.F64_add(base.F64_mul(base.F64_add(v105, v113), v97), base.F64_add(base.F64_add(v108, v110), v118))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[490])))
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
	v161 = *(*float64)(unsafe.Add(mBase, _consts[486]))
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
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
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v557 int32
	_ = v557
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v781 int32
	_ = v781
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v907 int32
	_ = v907
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v1018 int32
	_ = v1018
	var v1030 int32
	_ = v1030
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1276 int64
	_ = v1276
	var v1277 int64
	_ = v1277
	var v1278 int64
	_ = v1278
	var v1279 int64
	_ = v1279
	var v1280 int64
	_ = v1280
	var v1284 int64
	_ = v1284
	var v1297 int64
	_ = v1297
	var v1304 int64
	_ = v1304
	var v1309 int64
	_ = v1309
	var v1312 int64
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1320 int32
	_ = v1320
	var v1324 int64
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1514 int32
	_ = v1514
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1589 int32
	_ = v1589
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1640 int32
	_ = v1640
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1797 int32
	_ = v1797
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1882 int64
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1900 int32
	_ = v1900
	var v1905 int32
	_ = v1905
	var v1907 int64
	_ = v1907
	var v1908 int64
	_ = v1908
	var v1909 int64
	_ = v1909
	var v1910 int64
	_ = v1910
	var v1911 int64
	_ = v1911
	var v1915 int64
	_ = v1915
	var v1928 int64
	_ = v1928
	var v1935 int64
	_ = v1935
	var v1940 int64
	_ = v1940
	var v1943 int64
	_ = v1943
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2037 int32
	_ = v2037
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2092 int32
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2139 int32
	_ = v2139
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2153 int32
	_ = v2153
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2238 int32
	_ = v2238
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2253 int32
	_ = v2253
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2273 int32
	_ = v2273
	var v2280 int32
	_ = v2280
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2348 int32
	_ = v2348
	var v2351 int32
	_ = v2351
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2364 int32
	_ = v2364
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2435 int32
	_ = v2435
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2506 int32
	_ = v2506
	var v2515 int32
	_ = v2515
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2560 int32
	_ = v2560
	var v2564 int32
	_ = v2564
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2585 int32
	_ = v2585
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2679 int32
	_ = v2679
	var v2684 int32
	_ = v2684
	var v2719 int32
	_ = v2719
	var v2724 int32
	_ = v2724
	var v2729 int32
	_ = v2729
	var v2765 int32
	_ = v2765
	var v2771 int32
	_ = v2771
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2817 int32
	_ = v2817
	var v2826 int32
	_ = v2826
	var v2865 int32
	_ = v2865
	var v2866 int32
	_ = v2866
	var v2878 int64
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2887 int32
	_ = v2887
	var v2925 int64
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2970 int32
	_ = v2970
	var v2975 int32
	_ = v2975
	var v2976 int64
	_ = v2976
	var v2978 int64
	_ = v2978
	var v2980 int64
	_ = v2980
	var v2984 int32
	_ = v2984
	var v2993 int32
	_ = v2993
	var v3036 int32
	_ = v3036
	var v3040 int32
	_ = v3040
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3084 int64
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3129 int32
	_ = v3129
	var v3134 int32
	_ = v3134
	var v3135 int64
	_ = v3135
	var v3137 int64
	_ = v3137
	var v3139 int64
	_ = v3139
	var v3142 int32
	_ = v3142
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3154 int32
	_ = v3154
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3262 int64
	_ = v3262
	var v3264 int64
	_ = v3264
	var v3266 int32
	_ = v3266
	var v3272 int32
	_ = v3272
	var v3273 int64
	_ = v3273
	var v3275 int64
	_ = v3275
	var v3277 int32
	_ = v3277
	var v3278 int32
	_ = v3278
	var v3280 int32
	_ = v3280
	var v3283 int32
	_ = v3283
	var v3326 int32
	_ = v3326
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3336 int64
	_ = v3336
	var v3338 int64
	_ = v3338
	var v3382 int32
	_ = v3382
	var v3391 int32
	_ = v3391
	var v3427 int32
	_ = v3427
	var v3430 int32
	_ = v3430
	var v3433 int32
	_ = v3433
	var v3435 int32
	_ = v3435
	var v3479 int32
	_ = v3479
	var v3483 int32
	_ = v3483
	var v3484 int32
	_ = v3484
	var v3485 int32
	_ = v3485
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3507 int32
	_ = v3507
	var v3508 int32
	_ = v3508
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3515 int32
	_ = v3515
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3528 int64
	_ = v3528
	var v3529 int64
	_ = v3529
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3544 int32
	_ = v3544
	var v3549 int32
	_ = v3549
	var v3552 int32
	_ = v3552
	var v3554 int32
	_ = v3554
	var v3597 int32
	_ = v3597
	var v3599 int32
	_ = v3599
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3608 int32
	_ = v3608
	var v3609 int32
	_ = v3609
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3625 int32
	_ = v3625
	var v3626 int32
	_ = v3626
	var v3629 int32
	_ = v3629
	var v3630 int32
	_ = v3630
	var v3633 int32
	_ = v3633
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3643 int64
	_ = v3643
	var v3647 int64
	_ = v3647
	var v3651 int32
	_ = v3651
	var v3700 int32
	_ = v3700
	var v3704 int32
	_ = v3704
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3751 int32
	_ = v3751
	var v3755 int32
	_ = v3755
	var v3758 int32
	_ = v3758
	var v3803 int32
	_ = v3803
	var v3807 int32
	_ = v3807
	var v3850 int32
	_ = v3850
	var v3853 int32
	_ = v3853
	var v3857 int32
	_ = v3857
	var v3902 int32
	_ = v3902
	var v3906 int32
	_ = v3906
	var v3951 int32
	_ = v3951
	var v4020 int32
	_ = v4020
	var v4039 int32
	_ = v4039
	v44 = F_emscripten_builtin_malloc(m, int32(78408))
	mBase = m.M
	if v44 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[159]))
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
	v53 = int32(374660)
	goto L6
L6:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v57 = v53 + base.B2i32(v54 == int32(58))
	v58 = m.G0
	v60 = v58 - int32(1056)
	m.G0 = v60
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1255])))
	if v63 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_get_share_path(m, int32(4515904))
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
	v197 = v60 + int32(32)
	v198 = int32(4515904)
	goto L47
L10:
	;
	return int32(0)
L11:
	;
	v71 = int32(4515904)
	v72 = F_strlen(m, v71)
	mBase = m.M
	v74 = v72 + v71
	v75 = int32(372439)
	v77 = int32(1024) - v72
	if v77 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1255])) = uint8(v193)
	goto L9
L13:
	;
	v189 = F_strlen(m, v185)
	mBase = m.M
	goto L12
L14:
	;
	v185 = v75
	goto L13
L15:
	;
	goto L16
L16:
	;
	v83 = v77 - int32(1)
	if (v74^v75)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v182)
	v185 = v178
	goto L13
L18:
	;
	v163 = v158
	v164 = v159
	v165 = v160
	goto L40
L19:
	;
	if v153 == int32(0) {
		v178 = v151
		v179 = v152
		goto L17
	} else {
		goto L39
	}
L20:
	;
	v151 = v75
	v152 = v74
	v153 = v83
	goto L19
L21:
	;
	goto L22
L22:
	;
	goto L25
L23:
	;
	if v120 == int32(0) {
		v178 = v117
		v179 = v118
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v117 = v75
	v118 = v74
	v119 = v83
	v120 = base.B2i32(v83 != int32(0))
	goto L23
L25:
	;
	if v83 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v96 = v75
	v97 = v74
	v98 = v83
	goto L27
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v100)
	if v100 == int32(0) {
		v158 = v96
		v159 = v97
		v160 = v98
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v117 = v111
	v118 = v105
	v119 = v107
	v120 = v109
	goto L23
L29:
	;
	v104 = int32(1)
	v105 = v97 + v104
	v107 = v98 - v104
	v108 = int32(0)
	v109 = base.B2i32(v107 != v108)
	v111 = v96 + v104
	if v111&int32(3) == v108 {
		v117 = v111
		v118 = v105
		v119 = v107
		v120 = v109
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v107 != 0 {
		v96 = v111
		v97 = v105
		v98 = v107
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v123 == int32(0) {
		v151 = v117
		v152 = v118
		v153 = v119
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v119) < base.Ui32(int32(4)) {
		v151 = v117
		v152 = v118
		v153 = v119
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v129 = v117
	v130 = v118
	v131 = v119
	goto L35
L35:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v137 = int32(-2139062144)
	if (int32(16843008)-v134|v134)&v137 != v137 {
		v158 = v129
		v159 = v130
		v160 = v131
		goto L18
	} else {
		goto L37
	}
L36:
	;
	v151 = v145
	v152 = v143
	v153 = v147
	goto L19
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v134
	v142 = int32(4)
	v143 = v130 + v142
	v145 = v129 + v142
	v147 = v131 - v142
	if base.Ui32(int32(3)) < base.Ui32(v147) {
		v129 = v145
		v130 = v143
		v131 = v147
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v158 = v151
	v159 = v152
	v160 = v153
	goto L18
L40:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v167)
	if v167 == int32(0) {
		v178 = v163
		v179 = v164
		goto L17
	} else {
		goto L42
	}
L41:
	;
	v178 = v174
	v179 = v172
	goto L17
L42:
	;
	v171 = int32(1)
	v172 = v164 + v171
	v174 = v163 + v171
	v176 = v165 - v171
	if v176 != 0 {
		v163 = v174
		v164 = v172
		v165 = v176
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v317 = F_strlen(m, v60+int32(32))
	mBase = m.M
	v318 = F_strlen(m, v57)
	mBase = m.M
	if base.Ui32(int32(1023)) < base.Ui32(v317+v318+int32(1)) {
		v907 = int32(-1)
		goto L76
	} else {
		goto L77
	}
L45:
	;
	v311 = F_strlen(m, v300)
	mBase = m.M
	goto L44
L47:
	;
	goto L48
L48:
	;
	v205 = int32(1023)
	if (v197^v198)&int32(3) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v304)
	goto L45
L50:
	;
	v285 = v280
	v286 = v281
	v287 = v282
	goto L72
L51:
	;
	if v275 == int32(0) {
		v300 = v273
		v301 = v274
		goto L49
	} else {
		goto L71
	}
L52:
	;
	v273 = v198
	v274 = v197
	v275 = v205
	goto L51
L53:
	;
	goto L54
L54:
	;
	goto L56
L55:
	;
	goto L64
L56:
	;
	goto L55
L64:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1256])))
	if v245 == int32(0) {
		v273 = v198
		v274 = v197
		v275 = v205
		goto L51
	} else {
		goto L65
	}
L65:
	;
	goto L66
L66:
	;
	v251 = v198
	v252 = v197
	v253 = v205
	goto L67
L67:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v259 = int32(-2139062144)
	if (int32(16843008)-v256|v256)&v259 != v259 {
		v280 = v251
		v281 = v252
		v282 = v253
		goto L50
	} else {
		goto L69
	}
L68:
	;
	v273 = v267
	v274 = v265
	v275 = v269
	goto L51
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v256
	v264 = int32(4)
	v265 = v252 + v264
	v267 = v251 + v264
	v269 = v253 - v264
	if base.Ui32(int32(3)) < base.Ui32(v269) {
		v251 = v267
		v252 = v265
		v253 = v269
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v280 = v273
	v281 = v274
	v282 = v275
	goto L50
L72:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v289)
	if v289 == int32(0) {
		v300 = v285
		v301 = v286
		goto L49
	} else {
		goto L74
	}
L73:
	;
	v300 = v296
	v301 = v294
	goto L49
L74:
	;
	v293 = int32(1)
	v294 = v286 + v293
	v296 = v285 + v293
	v298 = v287 - v293
	if v298 != 0 {
		v285 = v296
		v286 = v294
		v287 = v298
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	m.G0 = v60 + int32(1056)
	if v907 < int32(0) {
		goto L213
	} else {
		goto L214
	}
L77:
	;
	if l1 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v328 = v60 + int32(32) + v317
	v329 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v328))) = uint8(v329)
	v332 = v328 + int32(1)
	if (v57^v332)&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	goto L80
L80:
	;
	v432 = v317
	v440 = v57
	goto L103
L81:
	;
	v407 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v407
	v414 = F_open(m, v60+int32(32), v407, v60+int32(16))
	mBase = m.M
	if v407 <= v414 {
		v907 = v414
		goto L76
	} else {
		goto L102
	}
L82:
	;
	goto L81
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v387))) = uint8(v386)
	if v386&int32(255) == int32(0) {
		goto L82
	} else {
		goto L98
	}
L84:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v385 = v57
	v386 = v338
	v387 = v332
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v57&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v342 = v57
	v344 = v332
	goto L90
L88:
	;
	v356 = v57
	v358 = v332
	goto L89
L89:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v363 = int32(-2139062144)
	if (int32(16843008)-v360|v360)&v363 != v363 {
		v385 = v356
		v386 = v360
		v387 = v358
		goto L83
	} else {
		goto L94
	}
L90:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	*(*uint8)(unsafe.Add(mBase, uint32(v344))) = uint8(v345)
	if v345 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L91:
	;
	v356 = v352
	v358 = v350
	goto L89
L92:
	;
	v349 = int32(1)
	v350 = v344 + v349
	v352 = v342 + v349
	if v352&int32(3) != 0 {
		v342 = v352
		v344 = v350
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v368 = v356
	v369 = v360
	v370 = v358
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370))) = v369
	v372 = int32(4)
	v373 = v370 + v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	v376 = v368 + v372
	v380 = int32(-2139062144)
	if (v374|(int32(16843008)-v374))&v380 == v380 {
		v368 = v376
		v369 = v374
		v370 = v373
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v385 = v376
	v386 = v374
	v387 = v373
	goto L83
L97:
	;
	goto L96
L98:
	;
	v394 = v385
	v396 = v387
	goto L99
L99:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v396)+1)) = uint8(v397)
	v399 = int32(1)
	if v397 != 0 {
		v394 = v394 + v399
		v396 = v396 + v399
		goto L99
	} else {
		goto L101
	}
L100:
	;
	goto L82
L101:
	;
	goto L100
L102:
	;
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v328))) = uint8(v417)
	goto L80
L103:
	;
	v467 = int32(47)
	v468 = F___strchrnul(m, v440, v467)
	mBase = m.M
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v470 == v467 {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	if l1 != 0 {
		goto L178
	} else {
		goto L179
	}
L105:
	;
	v480 = F_AllocateDir(m, v60+int32(32))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L10
	} else {
		goto L114
	}
L106:
	;
	if v474 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v474 = v468
	goto L109
L108:
	;
	v474 = int32(0)
	goto L109
L109:
	;
	goto L106
L110:
	;
	v477 = v474 - v440
	goto L105
L111:
	;
	goto L112
L112:
	;
	v476 = F_strlen(m, v440)
	mBase = m.M
	v477 = v476
	goto L105
L113:
	;
	if v488 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L114:
	;
	v485 = F_ReadDirExtended(m, v480, v60+int32(32), int32(15))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	if v485 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v488 = int32(1023) - v432
	v489 = v432 + (v60 + int32(32) | int32(1))
	v501 = v485
	goto L119
L117:
	;
	goto L118
L118:
	;
	F_FreeDir(m, v480)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L10
	} else {
		goto L143
	}
L119:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+19)))
	if v532 == int32(46) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L118
L121:
	;
	v595 = F_ReadDirExtended(m, v480, v60+int32(32), int32(15))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L10
	} else {
		goto L141
	}
L122:
	;
	v536 = v501 + int32(19)
	v537 = F_strlen(m, v536)
	mBase = m.M
	if v537 != v477 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v541 = v536
	v542 = v440
	v543 = v477
	goto L125
L124:
	;
	if v588 == int32(0) {
		goto L113
	} else {
		goto L140
	}
L125:
	;
	if v543 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v588 = int32(0)
	goto L124
L127:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541))))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	if v546 == v547 {
		v569 = v546
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	goto L126
L130:
	;
	v571 = int32(1)
	if v569 != 0 {
		v541 = v541 + v571
		v542 = v542 + v571
		v543 = v543 - v571
		goto L125
	} else {
		goto L139
	}
L131:
	;
	if base.Ui32((v546-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v557 = v546 | int32(32)
	goto L134
L133:
	;
	v557 = v546
	goto L134
L134:
	;
	if base.Ui32((v547-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v566 = v547 | int32(32)
	goto L137
L136:
	;
	v566 = v547
	goto L137
L137:
	;
	if v557 == v566 {
		v569 = v557
		goto L130
	} else {
		goto L138
	}
L138:
	;
	v588 = v557 - v566
	goto L124
L139:
	;
	goto L129
L140:
	;
	goto L121
L141:
	;
	if v595 != 0 {
		v501 = v595
		goto L119
	} else {
		goto L142
	}
L142:
	;
	goto L120
L143:
	;
	v907 = int32(-1)
	goto L76
L144:
	;
	F_FreeDir(m, v480)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L10
	} else {
		goto L176
	}
L145:
	;
	v753 = F_strlen(m, v749)
	mBase = m.M
	goto L144
L146:
	;
	v749 = v536
	goto L145
L147:
	;
	goto L148
L148:
	;
	v647 = v488 - int32(1)
	if (v489^v536)&int32(3) != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v746 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v743))) = uint8(v746)
	v749 = v742
	goto L145
L150:
	;
	v727 = v722
	v728 = v723
	v729 = v724
	goto L172
L151:
	;
	if v717 == int32(0) {
		v742 = v715
		v743 = v716
		goto L149
	} else {
		goto L171
	}
L152:
	;
	v715 = v536
	v716 = v489
	v717 = v647
	goto L151
L153:
	;
	goto L154
L154:
	;
	v651 = int32(0)
	if v536&int32(3) == v651 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if v684 == int32(0) {
		v742 = v681
		v743 = v682
		goto L149
	} else {
		goto L164
	}
L156:
	;
	v681 = v536
	v682 = v489
	v683 = v647
	v684 = base.B2i32(v647 != v651)
	goto L155
L157:
	;
	if v647 == int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v660 = v536
	v661 = v489
	v662 = v647
	goto L159
L159:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	*(*uint8)(unsafe.Add(mBase, uint32(v661))) = uint8(v664)
	if v664 == int32(0) {
		v722 = v660
		v723 = v661
		v724 = v662
		goto L150
	} else {
		goto L161
	}
L160:
	;
	v681 = v675
	v682 = v669
	v683 = v671
	v684 = v673
	goto L155
L161:
	;
	v668 = int32(1)
	v669 = v661 + v668
	v671 = v662 - v668
	v672 = int32(0)
	v673 = base.B2i32(v671 != v672)
	v675 = v660 + v668
	if v675&int32(3) == v672 {
		v681 = v675
		v682 = v669
		v683 = v671
		v684 = v673
		goto L155
	} else {
		goto L162
	}
L162:
	;
	if v671 != 0 {
		v660 = v675
		v661 = v669
		v662 = v671
		goto L159
	} else {
		goto L163
	}
L163:
	;
	goto L160
L164:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	if v687 == int32(0) {
		v715 = v681
		v716 = v682
		v717 = v683
		goto L151
	} else {
		goto L165
	}
L165:
	;
	if base.Ui32(v683) < base.Ui32(int32(4)) {
		v715 = v681
		v716 = v682
		v717 = v683
		goto L151
	} else {
		goto L166
	}
L166:
	;
	v693 = v681
	v694 = v682
	v695 = v683
	goto L167
L167:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	v701 = int32(-2139062144)
	if (int32(16843008)-v698|v698)&v701 != v701 {
		v722 = v693
		v723 = v694
		v724 = v695
		goto L150
	} else {
		goto L169
	}
L168:
	;
	v715 = v709
	v716 = v707
	v717 = v711
	goto L151
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v694))) = v698
	v706 = int32(4)
	v707 = v694 + v706
	v709 = v693 + v706
	v711 = v695 - v706
	if base.Ui32(int32(3)) < base.Ui32(v711) {
		v693 = v709
		v694 = v707
		v695 = v711
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v722 = v715
	v723 = v716
	v724 = v717
	goto L150
L172:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727))))
	*(*uint8)(unsafe.Add(mBase, uint32(v728))) = uint8(v731)
	if v731 == int32(0) {
		v742 = v727
		v743 = v728
		goto L149
	} else {
		goto L174
	}
L173:
	;
	v742 = v738
	v743 = v736
	goto L149
L174:
	;
	v735 = int32(1)
	v736 = v728 + v735
	v738 = v727 + v735
	v740 = v729 - v735
	if v740 != 0 {
		v727 = v738
		v728 = v736
		v729 = v740
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v759 = v60 + int32(32)
	v761 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v759+v432))) = uint8(v761)
	v763 = int32(1)
	v764 = v432 + v763
	v768 = F_strlen(m, v764+v759)
	mBase = m.M
	if v474 != 0 {
		v432 = v768 + v764
		v440 = v474 + v763
		goto L103
	} else {
		goto L177
	}
L177:
	;
	goto L104
L178:
	;
	v774 = v317 + v60 + int32(33)
	goto L184
L179:
	;
	goto L180
L180:
	;
	v890 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v890
	v895 = F_open(m, v60+int32(32), v890, v60)
	mBase = m.M
	v907 = v895
	goto L76
L181:
	;
	goto L180
L182:
	;
	v887 = F_strlen(m, v876)
	mBase = m.M
	goto L181
L184:
	;
	goto L185
L185:
	;
	v781 = int32(255)
	if (l1^v774)&int32(3) != 0 {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	v880 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v877))) = uint8(v880)
	goto L182
L187:
	;
	v861 = v856
	v862 = v857
	v863 = v858
	goto L209
L188:
	;
	if v851 == int32(0) {
		v876 = v849
		v877 = v850
		goto L186
	} else {
		goto L208
	}
L189:
	;
	v849 = v774
	v850 = l1
	v851 = v781
	goto L188
L190:
	;
	goto L191
L191:
	;
	if v774&int32(3) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if v818 == int32(0) {
		v876 = v815
		v877 = v816
		goto L186
	} else {
		goto L201
	}
L193:
	;
	v815 = v774
	v816 = l1
	v817 = v781
	v818 = int32(1)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v794 = v774
	v795 = l1
	v796 = v781
	goto L196
L196:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	*(*uint8)(unsafe.Add(mBase, uint32(v795))) = uint8(v798)
	if v798 == int32(0) {
		v856 = v794
		v857 = v795
		v858 = v796
		goto L187
	} else {
		goto L198
	}
L197:
	;
	v815 = v809
	v816 = v803
	v817 = v805
	v818 = v807
	goto L192
L198:
	;
	v802 = int32(1)
	v803 = v795 + v802
	v805 = v796 - v802
	v806 = int32(0)
	v807 = base.B2i32(v805 != v806)
	v809 = v794 + v802
	if v809&int32(3) == v806 {
		v815 = v809
		v816 = v803
		v817 = v805
		v818 = v807
		goto L192
	} else {
		goto L199
	}
L199:
	;
	if v805 != 0 {
		v794 = v809
		v795 = v803
		v796 = v805
		goto L196
	} else {
		goto L200
	}
L200:
	;
	goto L197
L201:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815))))
	if v821 == int32(0) {
		v849 = v815
		v850 = v816
		v851 = v817
		goto L188
	} else {
		goto L202
	}
L202:
	;
	if base.Ui32(v817) < base.Ui32(int32(4)) {
		v849 = v815
		v850 = v816
		v851 = v817
		goto L188
	} else {
		goto L203
	}
L203:
	;
	v827 = v815
	v828 = v816
	v829 = v817
	goto L204
L204:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v827)))
	v835 = int32(-2139062144)
	if (int32(16843008)-v832|v832)&v835 != v835 {
		v856 = v827
		v857 = v828
		v858 = v829
		goto L187
	} else {
		goto L206
	}
L205:
	;
	v849 = v843
	v850 = v841
	v851 = v845
	goto L188
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828))) = v832
	v840 = int32(4)
	v841 = v828 + v840
	v843 = v827 + v840
	v845 = v829 - v840
	if base.Ui32(int32(3)) < base.Ui32(v845) {
		v827 = v843
		v828 = v841
		v829 = v845
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v856 = v849
	v857 = v850
	v858 = v851
	goto L187
L209:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	*(*uint8)(unsafe.Add(mBase, uint32(v862))) = uint8(v865)
	if v865 == int32(0) {
		v876 = v861
		v877 = v862
		goto L186
	} else {
		goto L211
	}
L210:
	;
	v876 = v872
	v877 = v870
	goto L186
L211:
	;
	v869 = int32(1)
	v870 = v862 + v869
	v872 = v861 + v869
	v874 = v863 - v869
	if v874 != 0 {
		v861 = v872
		v862 = v870
		v863 = v874
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return int32(44)
L214:
	;
	goto L215
L215:
	;
	v947 = F_read(m, v907, v44, int32(54968))
	mBase = m.M
	if v947 <= int32(43) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	if v947 < int32(0) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L218
L218:
	;
	v959 = F_close(m, v907)
	mBase = m.M
	if int32(0) <= v959 {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v954 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v955 = v954
	goto L221
L220:
	;
	v955 = int32(28)
	goto L221
L221:
	;
	v956 = F_close(m, v907)
	mBase = m.M
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return v955
L222:
	;
	v963 = l2 + int32(22632)
	v965 = l2 + int32(22120)
	v967 = l2 + int32(18024)
	v969 = l2 + int32(24)
	v971 = l2 + int32(16024)
	v973 = v44 + int32(44)
	v1018 = int32(4)
	v1030 = v947
	goto L226
L223:
	;
	goto L224
L224:
	;
	v4039 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return v4039
L225:
	;
	F_emscripten_builtin_free(m, v44)
	mBase = m.M
	return v4020
L226:
	;
	v1053 = int32(28)
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(31)))))
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(30)))))
	v1056 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+28)))
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(29)))))
	v1062 = int32(8)
	v1068 = v1054 | (v1055|(v1056&int32(127)<<(uint(int32(16))%32)|v1061<<(uint(v1062)%32)))<<(uint(v1062)%32)
	if v1056 < int32(0) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	if v2348 < int32(3) {
		goto L409
	} else {
		goto L410
	}
L228:
	;
	v1073 = v1068 | int32(-2147483648)
	goto L230
L229:
	;
	v1073 = v1068
	goto L230
L230:
	;
	if base.Ui32(int32(49)) < base.Ui32(v1073) {
		v4020 = v1053
		goto L225
	} else {
		goto L231
	}
L231:
	;
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(39)))))
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(38)))))
	v1078 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+36)))
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(37)))))
	v1084 = int32(8)
	v1090 = v1076 | (v1077|(v1078&int32(127)<<(uint(int32(16))%32)|v1083<<(uint(v1084)%32)))<<(uint(v1084)%32)
	if v1078 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1095 = v1090 | int32(-2147483648)
	goto L234
L233:
	;
	v1095 = v1090
	goto L234
L234:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1095) {
		v4020 = v1053
		goto L225
	} else {
		goto L235
	}
L235:
	;
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(35)))))
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(34)))))
	v1100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+32)))
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(33)))))
	v1106 = int32(8)
	v1112 = v1098 | (v1099|(v1100&int32(127)<<(uint(int32(16))%32)|v1105<<(uint(v1106)%32)))<<(uint(v1106)%32)
	if v1100 < int32(0) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1117 = v1112 | int32(-2147483648)
	goto L238
L237:
	;
	v1117 = v1112
	goto L238
L238:
	;
	if base.Ui32(int32(1999)) < base.Ui32(v1117) {
		v4020 = v1053
		goto L225
	} else {
		goto L239
	}
L239:
	;
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(43)))))
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(42)))))
	v1122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+40)))
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(41)))))
	v1128 = int32(8)
	v1134 = v1120 | (v1121|(v1122&int32(127)<<(uint(int32(16))%32)|v1127<<(uint(v1128)%32)))<<(uint(v1128)%32)
	if v1122 < int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1139 = v1134 | int32(-2147483648)
	goto L242
L241:
	;
	v1139 = v1134
	goto L242
L242:
	;
	if base.Ui32(int32(49)) < base.Ui32(v1139) {
		v4020 = v1053
		goto L225
	} else {
		goto L243
	}
L243:
	;
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(21)))))
	v1143 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+20)))
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(22)))))
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(23)))))
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(27)))))
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(26)))))
	v1148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+24)))
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(25)))))
	v1154 = int32(8)
	v1160 = v1146 | (v1147|(v1148&int32(127)<<(uint(int32(16))%32)|v1153<<(uint(v1154)%32)))<<(uint(v1154)%32)
	if v1148 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1165 = v1160 | int32(-2147483648)
	goto L246
L245:
	;
	v1165 = v1160
	goto L246
L246:
	;
	if v1095 != v1165 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1168 = v1165
	goto L249
L248:
	;
	v1168 = int32(0)
	goto L249
L249:
	;
	if v1168 != 0 {
		v4020 = v1053
		goto L225
	} else {
		goto L250
	}
L250:
	;
	v1173 = int32(8)
	v1179 = (v1143&int32(127)<<(uint(int32(16))%32)|v1142<<(uint(v1173)%32)|v1144)<<(uint(v1173)%32) | v1145
	if v1143 < int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1184 = v1179 | int32(-2147483648)
	goto L253
L252:
	;
	v1184 = v1179
	goto L253
L253:
	;
	if v1095 != v1184 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1187 = v1184
	goto L256
L255:
	;
	v1187 = int32(0)
	goto L256
L256:
	;
	if v1187 != 0 {
		v4020 = v1053
		goto L225
	} else {
		goto L257
	}
L257:
	;
	v1190 = v1018 + int32(4)
	if v1030 < v1184+v1165+v1073*v1190+v1117+v1117*v1018+v1095*int32(6)+v1139+int32(44) {
		v4020 = v1053
		goto L225
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v1139
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1073
	v1207 = int32(0)
	if v1117 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1210 = v973
	v1211 = int32(0)
	v1213 = v1207
	goto L262
L260:
	;
	v1451 = v973
	v1452 = v1095
	v1454 = v1207
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1454
	if v1452 <= int32(0) {
		goto L293
	} else {
		goto L294
	}
L262:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210))))
	v1254 = v1252 & int32(127)
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+2)))
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+1)))
	v1257 = base.I32_extend8_s(v1252)
	if v1018 == int32(4) {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v1345 = int32(0)
	if v1345 < v1343 {
		goto L281
	} else {
		goto L282
	}
L264:
	;
	v1313 = v1213 + v971
	v1314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1313))) = uint8(v1314)
	if v1211 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L265:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+3)))
	v1263 = int32(8)
	v1269 = v1260 | (v1254<<(uint(int32(16))%32)|v1256<<(uint(v1263)%32)|v1255)<<(uint(v1263)%32)
	if v1257 < int32(0) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L267
L267:
	;
	v1276 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+7)))
	v1277 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+6)))
	v1278 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+4)))
	v1279 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+3)))
	v1280 = int64(8)
	v1284 = int64(16)
	v1297 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1210)+5)))
	v1304 = v1276 | (v1277|((v1278|(v1279<<(uint(v1280)%64)|(base.I64_extend_i32_u(v1255)|(base.I64_extend_i32_u(v1254)<<(uint(v1284)%64)|base.I64_extend_i32_u(v1256)<<(uint(v1280)%64)))<<(uint(v1284)%64)))<<(uint(v1284)%64)|v1297<<(uint(v1280)%64)))<<(uint(v1280)%64)
	if v1257 < int32(0) {
		goto L271
	} else {
		goto L272
	}
L268:
	;
	v1274 = v1269 | int32(-2147483648)
	goto L270
L269:
	;
	v1274 = v1269
	goto L270
L270:
	;
	v1312 = base.I64_extend_i32_s(v1274)
	goto L264
L271:
	;
	v1309 = v1304 | int64(-9223372036854775807-1)
	goto L273
L272:
	;
	v1309 = v1304
	goto L273
L273:
	;
	v1312 = v1309
	goto L264
L274:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v969+v1331<<(uint(int32(3))%32)))) = v1312
	v1338 = v1210 + v1018
	v1339 = int32(1)
	v1342 = v1213 + v1339
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1342 < v1343 {
		v1210 = v1338
		v1211 = v1331 + v1339
		v1213 = v1342
		goto L262
	} else {
		goto L280
	}
L275:
	;
	v1331 = int32(0)
	goto L274
L276:
	;
	goto L277
L277:
	;
	v1320 = v1211 - int32(1)
	v1324 = *(*int64)(unsafe.Add(mBase, uint32(v969+v1320<<(uint(int32(3))%32))))
	if v1324 < v1312 {
		v1331 = v1211
		goto L274
	} else {
		goto L278
	}
L278:
	;
	if v1312 < v1324 {
		v4020 = v1053
		goto L225
	} else {
		goto L279
	}
L279:
	;
	v1329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1313-int32(1)))) = uint8(v1329)
	v1331 = v1320
	goto L274
L280:
	;
	goto L263
L281:
	;
	v1349 = v1338
	v1350 = v1345
	v1352 = v1345
	v1353 = v1343
	goto L284
L282:
	;
	v1408 = v1338
	v1411 = v1345
	goto L283
L283:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1451 = v1408
	v1452 = v1450
	v1454 = v1411
	goto L261
L284:
	;
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349))))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1392 <= v1391 {
		v4020 = v1053
		goto L225
	} else {
		goto L286
	}
L285:
	;
	v1408 = v1404
	v1411 = v1401
	goto L283
L286:
	;
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350+v971))))
	if v1395 != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1352+v971))) = uint8(v1391)
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1401 = v1352 + int32(1)
	v1402 = v1398
	goto L289
L288:
	;
	v1401 = v1352
	v1402 = v1353
	goto L289
L289:
	;
	v1403 = int32(1)
	v1404 = v1349 + v1403
	v1406 = v1350 + v1403
	if v1406 < v1402 {
		v1349 = v1404
		v1350 = v1406
		v1352 = v1401
		v1353 = v1402
		goto L284
	} else {
		goto L290
	}
L290:
	;
	goto L285
L291:
	;
	v1830 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1797+v965))) = uint8(v1830)
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1834 <= v1830 {
		goto L317
	} else {
		goto L318
	}
L292:
	;
	v1622 = int32(3)
	v1623 = v1589 & v1622
	if base.Ui32(v1589-int32(1)) < base.Ui32(v1622) {
		goto L306
	} else {
		goto L307
	}
L293:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if int32(0) < v1496 {
		v1580 = v1451
		v1589 = v1496
		goto L292
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v1500 = v1451
	v1514 = v1207
	goto L297
L296:
	;
	v1788 = v1451
	v1797 = int32(0)
	goto L291
L297:
	;
	v1544 = v967 + v1514<<(uint(int32(4))%32)
	v1545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+3)))
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+2)))
	v1547 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1500))))
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+1)))
	v1553 = int32(8)
	v1559 = v1545 | (v1546|(v1547&int32(127)<<(uint(int32(16))%32)|v1552<<(uint(v1553)%32)))<<(uint(v1553)%32)
	if v1547 < int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1580 = v1575
	v1589 = v1570
	goto L292
L299:
	;
	v1564 = v1559 | int32(-2147483648)
	goto L301
L300:
	;
	v1564 = v1559
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544))) = v1564
	v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+4)))
	if base.Ui32(int32(1)) < base.Ui32(v1566) {
		v4020 = v1053
		goto L225
	} else {
		goto L302
	}
L302:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1544)+4)) = uint8(v1566)
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+5)))
	if v1570 <= v1571 {
		v4020 = v1053
		goto L225
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+8)) = v1571
	v1575 = v1500 + int32(6)
	v1577 = v1514 + int32(1)
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1577 < v1578 {
		v1500 = v1575
		v1514 = v1577
		goto L297
	} else {
		goto L304
	}
L304:
	;
	goto L298
L305:
	;
	if v1623 == int32(0) {
		v1788 = v1692
		v1797 = v1589
		goto L291
	} else {
		goto L312
	}
L306:
	;
	v1692 = v1580
	v1693 = int32(0)
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1632 = int32(0)
	v1634 = v1580
	v1635 = v1632
	v1640 = v1632
	goto L309
L309:
	;
	v1676 = v1635 + v965
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1676))) = uint8(v1677)
	v1679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1676)+1)) = uint8(v1679)
	v1681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1676)+2)) = uint8(v1681)
	v1683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1634)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1676)+3)) = uint8(v1683)
	v1685 = int32(4)
	v1686 = v1635 + v1685
	v1688 = v1634 + v1685
	v1690 = v1640 + v1685
	if v1690 != v1589&int32(-4) {
		v1634 = v1688
		v1635 = v1686
		v1640 = v1690
		goto L309
	} else {
		goto L311
	}
L310:
	;
	v1692 = v1688
	v1693 = v1686
	goto L305
L311:
	;
	goto L310
L312:
	;
	v1736 = v1692
	v1737 = v1693
	v1740 = int32(0)
	goto L313
L313:
	;
	v1779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1736))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1737+v965))) = uint8(v1779)
	v1781 = int32(1)
	v1784 = v1736 + v1781
	v1786 = v1740 + v1781
	if v1786 != v1623 {
		v1736 = v1784
		v1737 = v1737 + v1781
		v1740 = v1786
		goto L313
	} else {
		goto L315
	}
L314:
	;
	v1788 = v1784
	v1797 = v1589
	goto L291
L315:
	;
	goto L314
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1992
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v2031 {
		goto L339
	} else {
		goto L340
	}
L317:
	;
	v1988 = v1788
	v1992 = int32(0)
	goto L316
L318:
	;
	goto L319
L319:
	;
	v1839 = int32(0)
	v1841 = v1788
	v1845 = v1839
	v1847 = v1839
	v1882 = int64(0)
	goto L320
L320:
	;
	v1883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1841))))
	v1885 = v1883 & int32(127)
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+2)))
	v1887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+1)))
	v1888 = base.I32_extend8_s(v1883)
	if v1018 == int32(4) {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	v1988 = v1977
	v1992 = v1985
	goto L316
L322:
	;
	if v1943 < int64(0) {
		v4020 = v1053
		goto L225
	} else {
		goto L332
	}
L323:
	;
	v1891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+3)))
	v1894 = int32(8)
	v1900 = v1891 | (v1885<<(uint(int32(16))%32)|v1887<<(uint(v1894)%32)|v1886)<<(uint(v1894)%32)
	if v1888 < int32(0) {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	goto L325
L325:
	;
	v1907 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+7)))
	v1908 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+6)))
	v1909 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+4)))
	v1910 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+3)))
	v1911 = int64(8)
	v1915 = int64(16)
	v1928 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+5)))
	v1935 = v1907 | (v1908|((v1909|(v1910<<(uint(v1911)%64)|(base.I64_extend_i32_u(v1886)|(base.I64_extend_i32_u(v1885)<<(uint(v1915)%64)|base.I64_extend_i32_u(v1887)<<(uint(v1911)%64)))<<(uint(v1915)%64)))<<(uint(v1915)%64)|v1928<<(uint(v1911)%64)))<<(uint(v1911)%64)
	if v1888 < int32(0) {
		goto L329
	} else {
		goto L330
	}
L326:
	;
	v1905 = v1900 | int32(-2147483648)
	goto L328
L327:
	;
	v1905 = v1900
	goto L328
L328:
	;
	v1943 = base.I64_extend_i32_s(v1905)
	goto L322
L329:
	;
	v1940 = v1935 | int64(-9223372036854775807-1)
	goto L331
L330:
	;
	v1940 = v1935
	goto L331
L331:
	;
	v1943 = v1940
	goto L322
L332:
	;
	if v1943-v1882 < int64(2419199) {
		v4020 = v1053
		goto L225
	} else {
		goto L333
	}
L333:
	;
	v1949 = v1841 + v1018
	v1950 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1949))))
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1949)+1)))
	v1956 = int32(8)
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1949)+2)))
	v1963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1949)+3)))
	v1964 = (v1950&int32(127)<<(uint(int32(16))%32)|v1955<<(uint(v1956)%32)|v1959)<<(uint(v1956)%32) | v1963
	if v1950 < int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1969 = v1964 | int32(-2147483648)
	goto L336
L335:
	;
	v1969 = v1964
	goto L336
L336:
	;
	v1970 = int32(1)
	if base.B2i32(v1969 != v1847-v1970)&base.B2i32(v1969 != v1847+v1970) != 0 {
		v4020 = v1053
		goto L225
	} else {
		goto L337
	}
L337:
	;
	v1977 = v1841 + v1190
	v1980 = v963 + v1845<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1980)+8)) = base.I64_extend_i32_s(v1969)
	*(*int64)(unsafe.Add(mBase, uint32(v1980))) = v1943
	v1985 = v1845 + int32(1)
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1985 < v1986 {
		v1841 = v1977
		v1845 = v1985
		v1847 = v1969
		v1882 = v1943
		goto L320
	} else {
		goto L338
	}
L338:
	;
	goto L321
L339:
	;
	v2034 = v1988
	v2037 = v1830
	goto L342
L340:
	;
	v2153 = v1988
	goto L341
L341:
	;
	v2195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+4)))
	if v2195 != 0 {
		goto L358
	} else {
		goto L359
	}
L342:
	;
	v2076 = int32(0)
	if v1165 == v2076 {
		goto L345
	} else {
		goto L346
	}
L343:
	;
	v2094 = v2085
	v2095 = v2076
	goto L350
L344:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v967+v2037<<(uint(int32(4))%32))+12)) = uint8(v2086)
	v2092 = v2037 + int32(1)
	if v2092 != v2031 {
		v2034 = v2085
		v2037 = v2092
		goto L342
	} else {
		goto L349
	}
L345:
	;
	v2085 = v2034
	v2086 = int32(0)
	goto L344
L346:
	;
	goto L347
L347:
	;
	v2080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2034))))
	if base.Ui32(int32(1)) < base.Ui32(v2080) {
		v4020 = v1053
		goto L225
	} else {
		goto L348
	}
L348:
	;
	v2085 = v2034 + int32(1)
	v2086 = v2080
	goto L344
L349:
	;
	goto L343
L350:
	;
	if v1184 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L351:
	;
	v2153 = v2144
	goto L341
L352:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v967+v2095<<(uint(int32(4))%32))+13)) = uint8(v2145)
	v2151 = v2095 + int32(1)
	if v2151 != v2031 {
		v2094 = v2144
		v2095 = v2151
		goto L350
	} else {
		goto L357
	}
L353:
	;
	v2144 = v2094
	v2145 = int32(0)
	goto L352
L354:
	;
	goto L355
L355:
	;
	v2139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2094))))
	if base.Ui32(int32(1)) < base.Ui32(v2139) {
		v4020 = v1053
		goto L225
	} else {
		goto L356
	}
L356:
	;
	v2144 = v2094 + int32(1)
	v2145 = v2139
	goto L352
L357:
	;
	goto L351
L358:
	;
	v2197 = v44 - v2153 + v1030
	if v44 == v2153 {
		goto L362
	} else {
		goto L363
	}
L359:
	;
	v2348 = v1030
	goto L360
L360:
	;
	goto L227
L361:
	;
	if base.Ui32(v1018) < base.Ui32(int32(5)) {
		v1018 = v1018 << (uint(int32(1)) % 32)
		v1030 = v2197
		goto L226
	} else {
		goto L407
	}
L362:
	;
	goto L361
L363:
	;
	v2201 = v44 + v2197
	if base.Ui32(v2153-v2201) <= base.Ui32(int32(0)-v2197<<(uint(int32(1))%32)) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v2208 = F___memcpy(m, v44, v2153, v2197)
	mBase = m.M
	goto L361
L365:
	;
	goto L366
L366:
	;
	v2211 = (v44 ^ v2153) & int32(3)
	if base.Ui32(v44) < base.Ui32(v2153) {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	if v2313 == int32(0) {
		goto L362
	} else {
		goto L403
	}
L368:
	;
	if base.Ui32(v2291) <= base.Ui32(int32(3)) {
		v2312 = v2290
		v2313 = v2291
		v2314 = v2292
		goto L367
	} else {
		goto L399
	}
L369:
	;
	if v2211 != 0 {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	goto L371
L371:
	;
	if v2211 != 0 {
		v2273 = v2197
		goto L382
	} else {
		goto L383
	}
L372:
	;
	v2312 = v2153
	v2313 = v2197
	v2314 = v44
	goto L367
L373:
	;
	goto L374
L374:
	;
	if v44&int32(3) == int32(0) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v2290 = v2153
	v2291 = v2197
	v2292 = v44
	goto L368
L376:
	;
	goto L377
L377:
	;
	v2218 = v2153
	v2219 = v2197
	v2220 = v44
	goto L378
L378:
	;
	if v2219 == int32(0) {
		goto L362
	} else {
		goto L380
	}
L379:
	;
	v2290 = v2227
	v2291 = v2229
	v2292 = v2231
	goto L368
L380:
	;
	v2224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2218))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2220))) = uint8(v2224)
	v2226 = int32(1)
	v2227 = v2218 + v2226
	v2229 = v2219 - v2226
	v2231 = v2220 + v2226
	if v2231&int32(3) != 0 {
		v2218 = v2227
		v2219 = v2229
		v2220 = v2231
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	if v2273 == int32(0) {
		goto L362
	} else {
		goto L395
	}
L383:
	;
	if v2201&int32(3) != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v2238 = v2197
	goto L387
L385:
	;
	v2253 = v2197
	goto L386
L386:
	;
	if base.Ui32(v2253) <= base.Ui32(int32(3)) {
		v2273 = v2253
		goto L382
	} else {
		goto L391
	}
L387:
	;
	if v2238 == int32(0) {
		goto L362
	} else {
		goto L389
	}
L388:
	;
	v2253 = v2244
	goto L386
L389:
	;
	v2244 = v2238 - int32(1)
	v2245 = v44 + v2244
	v2247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2153+v2244))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2245))) = uint8(v2247)
	if v2245&int32(3) != 0 {
		v2238 = v2244
		goto L387
	} else {
		goto L390
	}
L390:
	;
	goto L388
L391:
	;
	v2260 = v2253
	goto L392
L392:
	;
	v2264 = v2260 - int32(4)
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2153+v2264)))
	*(*int32)(unsafe.Add(mBase, uint32(v44+v2264))) = v2267
	if base.Ui32(int32(3)) < base.Ui32(v2264) {
		v2260 = v2264
		goto L392
	} else {
		goto L394
	}
L393:
	;
	v2273 = v2264
	goto L382
L394:
	;
	goto L393
L395:
	;
	v2280 = v2273
	goto L396
L396:
	;
	v2284 = v2280 - int32(1)
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2153+v2284))))
	*(*uint8)(unsafe.Add(mBase, uint32(v44+v2284))) = uint8(v2287)
	if v2284 != 0 {
		v2280 = v2284
		goto L396
	} else {
		goto L398
	}
L397:
	;
	goto L362
L398:
	;
	goto L397
L399:
	;
	v2297 = v2290
	v2298 = v2291
	v2299 = v2292
	goto L400
L400:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2297)))
	*(*int32)(unsafe.Add(mBase, uint32(v2299))) = v2301
	v2303 = int32(4)
	v2304 = v2297 + v2303
	v2306 = v2299 + v2303
	v2308 = v2298 - v2303
	if base.Ui32(int32(3)) < base.Ui32(v2308) {
		v2297 = v2304
		v2298 = v2308
		v2299 = v2306
		goto L400
	} else {
		goto L402
	}
L401:
	;
	v2312 = v2304
	v2313 = v2308
	v2314 = v2306
	goto L367
L402:
	;
	goto L401
L403:
	;
	v2319 = v2312
	v2320 = v2313
	v2321 = v2314
	goto L404
L404:
	;
	v2323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2319))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2321))) = uint8(v2323)
	v2325 = int32(1)
	v2330 = v2320 - v2325
	if v2330 != 0 {
		v2319 = v2319 + v2325
		v2320 = v2330
		v2321 = v2321 + v2325
		goto L404
	} else {
		goto L406
	}
L405:
	;
	goto L362
L406:
	;
	goto L405
L407:
	;
	v2348 = v2197
	goto L360
L408:
	;
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v3427 < int32(2) {
		goto L501
	} else {
		goto L502
	}
L409:
	;
	v3382 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v3382 == int32(0) {
		v4020 = v1053
		goto L225
	} else {
		goto L500
	}
L410:
	;
	v2351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v2351 != int32(10) {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v2356 = v44 + v2348 - int32(1)
	v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2356))))
	if v2357 != int32(10) {
		goto L409
	} else {
		goto L412
	}
L412:
	;
	if int32(256) < v2031+int32(2) {
		v3391 = v2031
		goto L408
	} else {
		goto L413
	}
L413:
	;
	v2364 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2356))) = uint8(v2364)
	v2372 = F_tzparse(m, v44+int32(1), v44+int32(54968), v2364)
	mBase = m.M
	if v2372 == v2364 {
		goto L409
	} else {
		goto L414
	}
L414:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[1257])))
	if int32(0) < v2376 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v2387 = v2375
	v2388 = int32(0)
	v2392 = v2364
	goto L418
L416:
	;
	v2724 = v2375
	v2729 = v2364
	goto L417
L417:
	;
	if v2729 != v2376 {
		goto L409
	} else {
		goto L462
	}
L418:
	;
	v2428 = v44 + int32(73000) + v2388<<(uint(int32(4))%32)
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2428)))
	v2430 = v44 + int32(77088) + v2429
	v2431 = int32(0)
	if v2431 < v2387 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	v2724 = v2679
	v2729 = v2684
	goto L417
L420:
	;
	v2719 = v2388 + int32(1)
	if v2719 < v2376 {
		v2387 = v2679
		v2388 = v2719
		v2392 = v2684
		goto L418
	} else {
		goto L461
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2428))) = v2638
	v2679 = v2634
	v2684 = v2392 + int32(1)
	goto L420
L422:
	;
	v2435 = v2431
	goto L425
L423:
	;
	v2515 = v2431
	goto L424
L424:
	;
	v2550 = F_strlen(m, v2430)
	mBase = m.M
	v2551 = v2550 + v2515
	if int32(49) < v2551 {
		v2679 = v2387
		v2684 = v2392
		goto L420
	} else {
		goto L439
	}
L425:
	;
	v2477 = v2435 + v965
	v2480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2430))))
	v2481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2477))))
	if v2481 == int32(0) {
		v2500 = v2480
		v2501 = v2481
		goto L428
	} else {
		goto L429
	}
L426:
	;
	v2515 = v2387
	goto L424
L427:
	;
	if v2501-v2500 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L428:
	;
	goto L427
L429:
	;
	if v2480 != v2481 {
		v2500 = v2480
		v2501 = v2481
		goto L428
	} else {
		goto L430
	}
L430:
	;
	v2485 = v2477
	v2486 = v2430
	goto L431
L431:
	;
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2486)+1)))
	v2490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2485)+1)))
	if v2490 == int32(0) {
		v2500 = v2489
		v2501 = v2490
		goto L428
	} else {
		goto L433
	}
L432:
	;
	v2500 = v2489
	v2501 = v2490
	goto L428
L433:
	;
	v2493 = int32(1)
	if v2489 == v2490 {
		v2485 = v2485 + v2493
		v2486 = v2486 + v2493
		goto L431
	} else {
		goto L434
	}
L434:
	;
	goto L432
L435:
	;
	v2634 = v2387
	v2638 = v2435
	goto L421
L436:
	;
	goto L437
L437:
	;
	v2506 = v2435 + int32(1)
	if v2506 != v2387 {
		v2435 = v2506
		goto L425
	} else {
		goto L438
	}
L438:
	;
	goto L426
L439:
	;
	v2554 = v2515 + v965
	if (v2430^v2554)&int32(3) != 0 {
		goto L443
	} else {
		goto L444
	}
L440:
	;
	v2634 = v2551 + int32(1)
	v2638 = v2515
	goto L421
L441:
	;
	goto L440
L442:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2609))) = uint8(v2608)
	if v2608&int32(255) == int32(0) {
		goto L441
	} else {
		goto L457
	}
L443:
	;
	v2560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2430))))
	v2607 = v2430
	v2608 = v2560
	v2609 = v2554
	goto L442
L444:
	;
	goto L445
L445:
	;
	if v2430&int32(3) != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v2564 = v2430
	v2566 = v2554
	goto L449
L447:
	;
	v2578 = v2430
	v2580 = v2554
	goto L448
L448:
	;
	v2582 = *(*int32)(unsafe.Add(mBase, uint32(v2578)))
	v2585 = int32(-2139062144)
	if (int32(16843008)-v2582|v2582)&v2585 != v2585 {
		v2607 = v2578
		v2608 = v2582
		v2609 = v2580
		goto L442
	} else {
		goto L453
	}
L449:
	;
	v2567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2564))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2566))) = uint8(v2567)
	if v2567 == int32(0) {
		goto L441
	} else {
		goto L451
	}
L450:
	;
	v2578 = v2574
	v2580 = v2572
	goto L448
L451:
	;
	v2571 = int32(1)
	v2572 = v2566 + v2571
	v2574 = v2564 + v2571
	if v2574&int32(3) != 0 {
		v2564 = v2574
		v2566 = v2572
		goto L449
	} else {
		goto L452
	}
L452:
	;
	goto L450
L453:
	;
	v2590 = v2578
	v2591 = v2582
	v2592 = v2580
	goto L454
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2592))) = v2591
	v2594 = int32(4)
	v2595 = v2592 + v2594
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+4))
	v2598 = v2590 + v2594
	v2602 = int32(-2139062144)
	if (v2596|(int32(16843008)-v2596))&v2602 == v2602 {
		v2590 = v2598
		v2591 = v2596
		v2592 = v2595
		goto L454
	} else {
		goto L456
	}
L455:
	;
	v2607 = v2598
	v2608 = v2596
	v2609 = v2595
	goto L442
L456:
	;
	goto L455
L457:
	;
	v2616 = v2607
	v2618 = v2609
	goto L458
L458:
	;
	v2619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2616)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2618)+1)) = uint8(v2619)
	v2621 = int32(1)
	if v2619 != 0 {
		v2616 = v2616 + v2621
		v2618 = v2618 + v2621
		goto L458
	} else {
		goto L460
	}
L459:
	;
	goto L441
L460:
	;
	goto L459
L461:
	;
	goto L419
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v2724
	v2765 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2765 < int32(2) {
		v2826 = v2765
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[1258])))
	v2866 = int32(0)
	if v2826 == v2866 {
		v2993 = v2866
		goto L470
	} else {
		goto L471
	}
L464:
	;
	v2771 = v2765
	goto L465
L465:
	;
	v2811 = v2771 - int32(1)
	v2813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971+v2811))))
	v2817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2771+v971-int32(2)))))
	if v2813 != v2817 {
		v2826 = v2771
		goto L463
	} else {
		goto L467
	}
L466:
	;
	v2826 = int32(1)
	goto L463
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2811
	if base.Ui32(int32(2)) < base.Ui32(v2771) {
		v2771 = v2811
		goto L465
	} else {
		goto L468
	}
L468:
	;
	goto L466
L469:
	;
	if v2376 <= int32(0) {
		goto L409
	} else {
		goto L492
	}
L470:
	;
	if v2865 <= v2993 {
		goto L469
	} else {
		goto L482
	}
L471:
	;
	if v2865 <= int32(0) {
		v2993 = v2866
		goto L470
	} else {
		goto L472
	}
L472:
	;
	v2878 = *(*int64)(unsafe.Add(mBase, uint32(v2826<<(uint(int32(3))%32)+v969-int32(8))))
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2887 = v2866
	goto L473
L473:
	;
	v2925 = *(*int64)(unsafe.Add(mBase, uint32(v44+int32(54992)+v2887<<(uint(int32(3))%32))))
	v2926 = v2879
	goto L476
L474:
	;
	goto L469
L475:
	;
	if v2878 < v2925+v2980 {
		v2993 = v2887
		goto L470
	} else {
		goto L480
	}
L476:
	;
	v2970 = v2926 - int32(1)
	if v2970 < int32(0) {
		v2980 = int64(0)
		goto L475
	} else {
		goto L478
	}
L477:
	;
	v2978 = *(*int64)(unsafe.Add(mBase, uint32(v2975)+8))
	v2980 = v2978
	goto L475
L478:
	;
	v2975 = v963 + v2970<<(uint(int32(4))%32)
	v2976 = *(*int64)(unsafe.Add(mBase, uint32(v2975)))
	if v2925 < v2976 {
		v2926 = v2970
		goto L476
	} else {
		goto L479
	}
L479:
	;
	goto L477
L480:
	;
	v2984 = v2887 + int32(1)
	if v2984 != v2865 {
		v2887 = v2984
		goto L473
	} else {
		goto L481
	}
L481:
	;
	goto L474
L482:
	;
	v3036 = v2826
	v3040 = v2993
	goto L483
L483:
	;
	if int32(1999) < v3036 {
		goto L469
	} else {
		goto L485
	}
L484:
	;
	goto L469
L485:
	;
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v3078 = int32(3)
	v3084 = *(*int64)(unsafe.Add(mBase, uint32(v44+int32(54992)+v3040<<(uint(v3078)%32))))
	v3085 = v3077
	goto L487
L486:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v969+v3036<<(uint(v3078)%32)))) = v3084 + v3139
	v3142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+int32(70992)+v3040))))
	v3146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v3147 = v3145 + v3146
	*(*uint8)(unsafe.Add(mBase, uint32(v971+v3142))) = uint8(v3147)
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3150 = int32(1)
	v3151 = v3149 + v3150
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v3151
	v3154 = v3040 + v3150
	if v3154 != v2865 {
		v3036 = v3151
		v3040 = v3154
		goto L483
	} else {
		goto L491
	}
L487:
	;
	v3129 = v3085 - int32(1)
	if v3129 < int32(0) {
		v3139 = int64(0)
		goto L486
	} else {
		goto L489
	}
L488:
	;
	v3137 = *(*int64)(unsafe.Add(mBase, uint32(v3134)+8))
	v3139 = v3137
	goto L486
L489:
	;
	v3134 = v963 + v3129<<(uint(int32(4))%32)
	v3135 = *(*int64)(unsafe.Add(mBase, uint32(v3134)))
	if v3084 < v3135 {
		v3085 = v3129
		goto L487
	} else {
		goto L490
	}
L490:
	;
	goto L488
L491:
	;
	goto L484
L492:
	;
	v3200 = int32(1)
	v3203 = v44 + int32(72992)
	v3204 = int32(0)
	if v2376 != v3200 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	v3211 = v3204
	v3213 = int32(0)
	goto L496
L494:
	;
	v3283 = v3204
	goto L495
L495:
	;
	if v2376&v3200 == int32(0) {
		goto L409
	} else {
		goto L499
	}
L496:
	;
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v3253 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v3252 + v3253
	v3256 = int32(4)
	v3258 = v967 + v3252<<(uint(v3256)%32)
	v3261 = v3203 + v3211<<(uint(v3256)%32)
	v3262 = *(*int64)(unsafe.Add(mBase, uint32(v3261)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3258)+8)) = v3262
	v3264 = *(*int64)(unsafe.Add(mBase, uint32(v3261)))
	*(*int64)(unsafe.Add(mBase, uint32(v3258))) = v3264
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v3266 + v3253
	v3272 = v967 + v3266<<(uint(v3256)%32)
	v3273 = *(*int64)(unsafe.Add(mBase, uint32(v3261)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v3272)+8)) = v3273
	v3275 = *(*int64)(unsafe.Add(mBase, uint32(v3261)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v3272))) = v3275
	v3277 = int32(2)
	v3278 = v3211 + v3277
	v3280 = v3213 + v3277
	if v3280 != v2376&int32(2147483646) {
		v3211 = v3278
		v3213 = v3280
		goto L496
	} else {
		goto L498
	}
L497:
	;
	v3283 = v3278
	goto L495
L498:
	;
	goto L497
L499:
	;
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v3326 + int32(1)
	v3330 = int32(4)
	v3332 = v967 + v3326<<(uint(v3330)%32)
	v3335 = v3203 + v3283<<(uint(v3330)%32)
	v3336 = *(*int64)(unsafe.Add(mBase, uint32(v3335)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3332)+8)) = v3336
	v3338 = *(*int64)(unsafe.Add(mBase, uint32(v3335)))
	*(*int64)(unsafe.Add(mBase, uint32(v3332))) = v3338
	goto L409
L500:
	;
	v3391 = v3382
	goto L408
L501:
	;
	v3700 = int32(0)
	if v3427 <= v3700 {
		v3951 = v3700
		goto L544
	} else {
		goto L545
	}
L502:
	;
	v3430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+uint32(_consts[1259]))))
	v3433 = v967 + v3430<<(uint(int32(4))%32)
	v3435 = int32(1)
	goto L503
L503:
	;
	if v3391 <= v3430 {
		goto L506
	} else {
		goto L507
	}
L504:
	;
	v3544 = v3427 - int32(1)
	v3549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3544+v971))))
	v3552 = v967 + v3549<<(uint(int32(4))%32)
	v3554 = v3427 - int32(2)
	goto L524
L505:
	;
	goto L504
L506:
	;
	v3537 = v3435 + int32(1)
	if v3537 != v3427 {
		v3435 = v3537
		goto L503
	} else {
		goto L523
	}
L507:
	;
	v3479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3435+v971))))
	if v3391 <= v3479 {
		goto L506
	} else {
		goto L508
	}
L508:
	;
	v3483 = v967 + v3479<<(uint(int32(4))%32)
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v3483)))
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3433)))
	if v3484 != v3485 {
		goto L506
	} else {
		goto L509
	}
L509:
	;
	v3487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3483)+4)))
	v3488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3433)+4)))
	if v3487 != v3488 {
		goto L506
	} else {
		goto L510
	}
L510:
	;
	v3490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3483)+12)))
	v3491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3433)+12)))
	if v3490 != v3491 {
		goto L506
	} else {
		goto L511
	}
L511:
	;
	v3493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3483)+13)))
	v3494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3433)+13)))
	if v3493 != v3494 {
		goto L506
	} else {
		goto L512
	}
L512:
	;
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v3483)+8))
	v3497 = v965 + v3496
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v3433)+8))
	v3499 = v965 + v3498
	v3502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3499))))
	v3503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3497))))
	if v3503 == int32(0) {
		v3522 = v3502
		v3523 = v3503
		goto L514
	} else {
		goto L515
	}
L513:
	;
	if v3523-v3522 != 0 {
		goto L506
	} else {
		goto L521
	}
L514:
	;
	goto L513
L515:
	;
	if v3502 != v3503 {
		v3522 = v3502
		v3523 = v3503
		goto L514
	} else {
		goto L516
	}
L516:
	;
	v3507 = v3497
	v3508 = v3499
	goto L517
L517:
	;
	v3511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3508)+1)))
	v3512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3507)+1)))
	if v3512 == int32(0) {
		v3522 = v3511
		v3523 = v3512
		goto L514
	} else {
		goto L519
	}
L518:
	;
	v3522 = v3511
	v3523 = v3512
	goto L514
L519:
	;
	v3515 = int32(1)
	if v3511 == v3512 {
		v3507 = v3507 + v3515
		v3508 = v3508 + v3515
		goto L517
	} else {
		goto L520
	}
L520:
	;
	goto L518
L521:
	;
	v3528 = *(*int64)(unsafe.Add(mBase, uint32(v969+v3435<<(uint(int32(3))%32))))
	v3529 = *(*int64)(unsafe.Add(mBase, uint32(v969)))
	if v3528-v3529 != int64(12622780800) {
		goto L506
	} else {
		goto L522
	}
L522:
	;
	v3533 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v3533)
	goto L505
L523:
	;
	goto L505
L524:
	;
	if v3391 <= v3549 {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	goto L501
L526:
	;
	if int32(0) < v3554 {
		v3554 = v3554 - int32(1)
		goto L524
	} else {
		goto L543
	}
L527:
	;
	v3597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3554+v971))))
	if v3391 <= v3597 {
		goto L526
	} else {
		goto L528
	}
L528:
	;
	v3599 = *(*int32)(unsafe.Add(mBase, uint32(v3552)))
	v3602 = v967 + v3597<<(uint(int32(4))%32)
	v3603 = *(*int32)(unsafe.Add(mBase, uint32(v3602)))
	if v3599 != v3603 {
		goto L526
	} else {
		goto L529
	}
L529:
	;
	v3605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552)+4)))
	v3606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3602)+4)))
	if v3605 != v3606 {
		goto L526
	} else {
		goto L530
	}
L530:
	;
	v3608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552)+12)))
	v3609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3602)+12)))
	if v3608 != v3609 {
		goto L526
	} else {
		goto L531
	}
L531:
	;
	v3611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552)+13)))
	v3612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3602)+13)))
	if v3611 != v3612 {
		goto L526
	} else {
		goto L532
	}
L532:
	;
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(v3552)+8))
	v3615 = v965 + v3614
	v3616 = *(*int32)(unsafe.Add(mBase, uint32(v3602)+8))
	v3617 = v965 + v3616
	v3620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3617))))
	v3621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3615))))
	if v3621 == int32(0) {
		v3640 = v3620
		v3641 = v3621
		goto L534
	} else {
		goto L535
	}
L533:
	;
	if v3641-v3640 != 0 {
		goto L526
	} else {
		goto L541
	}
L534:
	;
	goto L533
L535:
	;
	if v3620 != v3621 {
		v3640 = v3620
		v3641 = v3621
		goto L534
	} else {
		goto L536
	}
L536:
	;
	v3625 = v3615
	v3626 = v3617
	goto L537
L537:
	;
	v3629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3626)+1)))
	v3630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3625)+1)))
	if v3630 == int32(0) {
		v3640 = v3629
		v3641 = v3630
		goto L534
	} else {
		goto L539
	}
L538:
	;
	v3640 = v3629
	v3641 = v3630
	goto L534
L539:
	;
	v3633 = int32(1)
	if v3629 == v3630 {
		v3625 = v3625 + v3633
		v3626 = v3626 + v3633
		goto L537
	} else {
		goto L540
	}
L540:
	;
	goto L538
L541:
	;
	v3643 = *(*int64)(unsafe.Add(mBase, uint32(v969+v3544<<(uint(int32(3))%32))))
	v3647 = *(*int64)(unsafe.Add(mBase, uint32(v969+v3554<<(uint(int32(3))%32))))
	if v3643-v3647 != int64(12622780800) {
		goto L526
	} else {
		goto L542
	}
L542:
	;
	v3651 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+17)) = uint8(v3651)
	goto L501
L543:
	;
	goto L525
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_consts[1260]))) = v3951
	v4020 = v3700
	goto L225
L545:
	;
	v3704 = v3700
	goto L547
L546:
	;
	v3951 = int32(0)
	goto L544
L547:
	;
	v3747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3704+v971))))
	if v3747 != 0 {
		goto L549
	} else {
		goto L550
	}
L548:
	;
	v3751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+uint32(_consts[1259]))))
	v3755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967+v3751<<(uint(int32(4))%32))+4)))
	if v3755 != int32(1) {
		goto L553
	} else {
		goto L554
	}
L549:
	;
	v3749 = v3704 + int32(1)
	if v3427 != v3749 {
		v3704 = v3749
		goto L547
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	goto L548
L552:
	;
	goto L546
L553:
	;
	v3850 = int32(1)
	if v3391 <= v3850 {
		goto L559
	} else {
		goto L560
	}
L554:
	;
	v3758 = v3751
	goto L555
L555:
	;
	if v3758 <= int32(0) {
		goto L553
	} else {
		goto L557
	}
L556:
	;
	v3951 = v3803
	goto L544
L557:
	;
	v3803 = v3758 - int32(1)
	v3807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967+v3803<<(uint(int32(4))%32))+4)))
	if v3807 != 0 {
		v3758 = v3803
		goto L555
	} else {
		goto L558
	}
L558:
	;
	goto L556
L559:
	;
	v3853 = v3850
	goto L561
L560:
	;
	v3853 = v3391
	goto L561
L561:
	;
	v3857 = int32(0)
	goto L562
L562:
	;
	v3902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+int32(18028)+v3857<<(uint(int32(4))%32)))))
	if v3902 != int32(1) {
		v3951 = v3857
		goto L544
	} else {
		goto L564
	}
L563:
	;
	goto L546
L564:
	;
	v3906 = v3857 + int32(1)
	if v3906 != v3853 {
		v3857 = v3906
		goto L562
	} else {
		goto L565
	}
L565:
	;
	goto L563
}
