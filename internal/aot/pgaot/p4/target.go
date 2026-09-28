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
	var v62 int32
	_ = v62
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v268 int32
	_ = v268
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
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
		v322 = v2
		v331 = v22
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v333+int32(3584))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L68
	}
L12:
	;
	if v55 == v54 {
		v322 = v2
		v331 = v22
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v62 = v55
	v65 = v2
	v72 = v2
	goto L14
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62-int32(4))))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	if v77 == v79 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v322 = v303
	v331 = v310<<(uint(int32(4))%32) ^ v22
	goto L11
L16:
	;
	if v74 != v54 {
		v62 = v74
		v65 = v303
		v72 = v310
		goto L14
	} else {
		goto L67
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
		v303 = v65
		v310 = v72
		goto L16
	} else {
		goto L23
	}
L20:
	;
	if int32(1) < v83 {
		v303 = v65
		v310 = v72
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v86 == int32(0) {
		v303 = v65
		v310 = v72
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v90 = v62 - int32(8)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v91
	v303 = v90
	v310 = base.I32_wrap_i64(int64(base.Ui64(v91) >> (uint(int64(32)) % 64)))
	goto L16
L23:
	;
	if v96&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+108)))
	if v120&int32(8) != 0 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v103 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v106 = int32(3)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	if base.B2i32(base.Ui32(v105) < base.Ui32(v106))|base.B2i32(base.Ui32(v108) < base.Ui32(v106)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v105-v108 < int32(0) {
		goto L24
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(v108) <= base.Ui32(v105) {
		v303 = v65
		v310 = v72
		goto L16
	} else {
		goto L31
	}
L30:
	;
	v303 = v65
	v310 = v72
	goto L16
L31:
	;
	goto L24
L32:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v176+int32(3584))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L42
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+108)))
	if v125&int32(8) != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	if v128 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v132 = v77 + int32(32)
	if v128 == v132 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v124)+44))
	if base.B2i32(v134 == int32(0))|base.B2i32(v134 == v124+int32(40)) != 0 {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v142 = v128
	goto L38
L38:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	if v156 == v124 {
		v303 = v65
		v310 = v72
		goto L16
	} else {
		goto L40
	}
L39:
	;
	goto L32
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v158 != v132 {
		v142 = v158
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v186 = F_LWLockAcquire(m, v182+int32(3584), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	if v188&int32(8) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v285+int32(3584))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L65
	}
L45:
	;
	if v188&int32(1) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+108)))
	if v214&int32(8) != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v195 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v198 = int32(3)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	if base.B2i32(base.Ui32(v197) < base.Ui32(v198))|base.B2i32(base.Ui32(v200) < base.Ui32(v198)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v197-v200 < int32(0) {
		goto L46
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(v200) <= base.Ui32(v197) {
		goto L44
	} else {
		goto L53
	}
L52:
	;
	goto L44
L53:
	;
	goto L46
L54:
	;
	F_FlagRWConflict(m, v77, v213)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L64
	}
L55:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+108)))
	if v217&int32(8) != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	if v220 == int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v224 = v77 + int32(32)
	if v220 == v224 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v213)+44))
	if base.B2i32(v226 == int32(0))|base.B2i32(v226 == v213+int32(40)) != 0 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v235 = v220
	goto L60
L60:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v235)+20))
	if v248 == v213 {
		goto L44
	} else {
		goto L62
	}
L61:
	;
	goto L54
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v250 != v224 {
		v235 = v250
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L44
L65:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v295 = F_LWLockAcquire(m, v291+int32(3584), int32(1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v303 = v65
	v310 = v72
	goto L16
L67:
	;
	goto L15
L68:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v322 == int32(0) {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v342 = int32(1)
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v348 = F_LWLockAcquire(m, v344+int32(3840), v342)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[3]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+72))
	if v353 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v356&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v356 = int32(1)
	goto L75
L74:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+76)))
	v356 = v355
	goto L75
L75:
	;
	goto L72
L76:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	v364 = F_LWLockAcquire(m, v360+int32(72), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v367 = F_LWLockAcquire(m, v32, int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	v374 = F_LWLockAcquire(m, v370+int32(3584), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[4]))
	v379 = v18 + int32(8)
	v380 = int32(0)
	v382 = F_hash_search_with_hash_value(m, v377, v379, v331, v380, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v382 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v384)+4)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v322)+16))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v322)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v322)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = v392
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[4]))
	v398 = F_hash_search_with_hash_value(m, v395, v379, v331, int32(2), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	v415 = v342
	goto L85
L85:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v417+int32(3584))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L94
	}
L86:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v400 != v54 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v403 = v400
	goto L89
L88:
	;
	v403 = int32(0)
	goto L89
L89:
	;
	if v403 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[0]))
	v410 = F_hash_search_with_hash_value(m, v407, v40, v22, int32(2), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v415 = base.B2i32(v398 == int32(0))
	goto L85
L93:
	;
	goto L92
L94:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[3]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+72))
	if v427 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v430&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v430 = int32(1)
	goto L99
L98:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+76)))
	v430 = v429
	goto L99
L99:
	;
	goto L96
L100:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[2]))
	F_LWLockRelease(m, v434+int32(72))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[1]))
	F_LWLockRelease(m, v440+int32(3840))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	if v415 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTargetForConflictsIn[5]))
	v449 = F_hash_search_with_hash_value(m, v446, l0, v22, int32(2), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_DecrementParentLocks(m, l0)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
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
