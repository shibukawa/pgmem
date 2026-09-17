package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckTargetForConflictsIn(m *base.Module, l0 int32) {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v269 int32
	_ = v269
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[0]))
	v22 = F_get_hash_value(m, v21, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v32 = v25 + v22&int32(15)<<(uint(int32(7))%32) + int32(_a_F_CheckTargetForConflictsIn_0)
	v34 = F_LWLockAcquire(m, v32, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[0]))
	v38 = int32(0)
	v40 = F_hash_search_with_hash_value(m, v37, l0, v22, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v18 + int32(16)
	return
L5:
	;
	if v40 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v51 = F_LWLockAcquire(m, v47+int32(3584), int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L4
L10:
	;
	v54 = v40 + int32(16)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v55 == int32(0) {
		v323 = v2
		v332 = v22
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v334+int32(3584))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L70
	}
L12:
	;
	if v55 == v54 {
		v323 = v2
		v332 = v22
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v61 = v55
	v65 = v2
	v72 = v2
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v61-int32(4))))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	if v77 == v79 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v323 = v304
	v332 = v311<<(uint(int32(4))%32) ^ v22
	goto L11
L16:
	;
	if v74 != v54 {
		v61 = v74
		v65 = v304
		v72 = v311
		goto L14
	} else {
		goto L69
	}
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[3]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	goto L20
L18:
	;
	goto L19
L19:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	if v96&int32(8) != 0 {
		v304 = v65
		v311 = v72
		goto L16
	} else {
		goto L23
	}
L20:
	;
	if int32(1) < v83 {
		v304 = v65
		v311 = v72
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v86 == int32(0) {
		v304 = v65
		v311 = v72
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v90 = v61 - int32(8)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v91
	v304 = v90
	v311 = base.I32_wrap_i64(int64(base.Ui64(v91) >> (uint(int64(32)) % 64)))
	goto L16
L23:
	;
	if v96&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v177+int32(3584))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L43
	}
L25:
	;
	v101 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v122 = v79
	v123 = v96
	goto L27
L27:
	;
	if v123&int32(8) != 0 {
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v104))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v103)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v116 == int32(0) {
		v304 = v65
		v311 = v72
		goto L16
	} else {
		goto L33
	}
L30:
	;
	v116 = base.B2i32(base.Ui32(v103) < base.Ui32(v104))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v116 = int32(base.Ui32(v103-v104) >> (uint(int32(31)) % 32))
	goto L29
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	v122 = v120
	v123 = v121
	goto L27
L34:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+108)))
	if v126&int32(8) != 0 {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	if v129 == int32(0) {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v133 = v77 + int32(32)
	if v129 == v133 {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+44))
	if base.B2i32(v135 == int32(0))|base.B2i32(v135 == v122+int32(40)) != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v143 = v129
	goto L39
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	if v157 == v122 {
		v304 = v65
		v311 = v72
		goto L16
	} else {
		goto L41
	}
L40:
	;
	goto L24
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v159 != v133 {
		v143 = v159
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v187 = F_LWLockAcquire(m, v183+int32(3584), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	if v189&int32(8) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v286+int32(3584))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L67
	}
L46:
	;
	if v189&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v194 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v213 = v189
	goto L49
L49:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	if v213&int32(8) != 0 {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v197))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v196)) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v209 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L52:
	;
	v209 = base.B2i32(base.Ui32(v196) < base.Ui32(v197))
	goto L51
L53:
	;
	goto L54
L54:
	;
	v209 = int32(base.Ui32(v196-v197) >> (uint(int32(31)) % 32))
	goto L51
L55:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	v213 = v212
	goto L49
L56:
	;
	F_FlagRWConflict(m, v77, v215)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L66
	}
L57:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+108)))
	if v218&int32(8) != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	if v221 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v225 = v77 + int32(32)
	if v221 == v225 {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v215)+44))
	if base.B2i32(v227 == int32(0))|base.B2i32(v227 == v215+int32(40)) != 0 {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v235 = v221
	goto L62
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	if v249 == v215 {
		goto L45
	} else {
		goto L64
	}
L63:
	;
	goto L56
L64:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v251 != v225 {
		v235 = v251
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	goto L45
L67:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v296 = F_LWLockAcquire(m, v292+int32(3584), int32(1))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v304 = v65
	v311 = v72
	goto L16
L69:
	;
	goto L15
L70:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v323 == int32(0) {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v343 = int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v349 = F_LWLockAcquire(m, v345+int32(3840), v343)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[3]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+72))
	if v354 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v357&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v357 = int32(1)
	goto L77
L76:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+76)))
	v357 = v356
	goto L77
L77:
	;
	goto L74
L78:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	v365 = F_LWLockAcquire(m, v361+int32(72), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v368 = F_LWLockAcquire(m, v32, int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v375 = F_LWLockAcquire(m, v371+int32(3584), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[4]))
	v380 = v18 + int32(8)
	v381 = int32(0)
	v383 = F_hash_search_with_hash_value(m, v378, v380, v332, v381, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v383 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v385)+4)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v386))) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v323)+16))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v323)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+4)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v323)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = v393
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[4]))
	v399 = F_hash_search_with_hash_value(m, v396, v380, v332, int32(2), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	v416 = v343
	goto L87
L87:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v418+int32(3584))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L96
	}
L88:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v401 != v54 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v404 = v401
	goto L91
L90:
	;
	v404 = int32(0)
	goto L91
L91:
	;
	if v404 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[0]))
	v411 = F_hash_search_with_hash_value(m, v408, v40, v22, int32(2), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v416 = base.B2i32(v399 == int32(0))
	goto L87
L95:
	;
	goto L94
L96:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[3]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+72))
	if v428 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v431&int32(1) != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v431 = int32(1)
	goto L101
L100:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+76)))
	v431 = v430
	goto L101
L101:
	;
	goto L98
L102:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	F_LWLockRelease(m, v435+int32(72))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v441+int32(3840))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	if v416 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[5]))
	v450 = F_hash_search_with_hash_value(m, v447, l0, v22, int32(2), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_DecrementParentLocks(m, l0)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L4
}
func F_ExecTargetListLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	if l0 == int32(0) {
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v6
	}
}
func F_targetIsInSortList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.B2i32(v5 == v3)|base.B2i32(l1 == v3) != 0 {
		v38 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v38
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v38 = v3
	goto L1
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v15<<(uint(int32(2))%32))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v5 == v24 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v38 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v28 = v15 + int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 < v29 {
		v15 = v28
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
}
