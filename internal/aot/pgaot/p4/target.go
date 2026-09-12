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
	var v64 int32
	_ = v64
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
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
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v267 int32
	_ = v267
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
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
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1109]))
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
	v25 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v32 = v25 + v22&int32(15)<<(uint(int32(7))%32) + int32(25344)
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
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1109]))
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
	v47 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v330+int32(3584))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L74
	}
L12:
	;
	v316 = v2
	v319 = v2
	goto L11
L13:
	;
	goto L14
L14:
	;
	if v55 == v54 {
		v316 = v2
		v319 = v2
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v62 = v55
	v64 = v2
	v72 = v2
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62-int32(4))))
	v79 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v77 == v79 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v316 = v309 << (uint(int32(4)) % 32)
	v319 = v301
	goto L11
L18:
	;
	if v74 != v54 {
		v62 = v74
		v64 = v301
		v72 = v309
		goto L16
	} else {
		goto L73
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	goto L22
L20:
	;
	goto L21
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	if v96&int32(8) != 0 {
		v301 = v64
		v309 = v72
		goto L18
	} else {
		goto L25
	}
L22:
	;
	if int32(1) < v83 {
		v301 = v64
		v309 = v72
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v86 == int32(0) {
		v301 = v64
		v309 = v72
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v90 = v62 - int32(8)
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v91
	v301 = v90
	v309 = base.I32_wrap_i64(int64(base.Ui64(v91) >> (uint(int64(32)) % 64)))
	goto L18
L25:
	;
	if v96&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v176+int32(3584))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L46
	}
L27:
	;
	v101 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v122 = v79
	v123 = v96
	goto L29
L29:
	;
	if v123&int32(8) != 0 {
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v104))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v103)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v116 == int32(0) {
		v301 = v64
		v309 = v72
		goto L18
	} else {
		goto L35
	}
L32:
	;
	v116 = base.B2i32(base.Ui32(v103) < base.Ui32(v104))
	goto L31
L33:
	;
	goto L34
L34:
	;
	v116 = int32(base.Ui32(v103-v104) >> (uint(int32(31)) % 32))
	goto L31
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	v122 = v120
	v123 = v121
	goto L29
L36:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+108)))
	if v126&int32(8) != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	if v129 == int32(0) {
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v133 = v77 + int32(32)
	if v129 == v133 {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+44))
	if v135 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	if v135 == v122+int32(40) {
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v142 = v129
	goto L42
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	if v156 == v122 {
		v301 = v64
		v309 = v72
		goto L18
	} else {
		goto L44
	}
L43:
	;
	goto L26
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v158 != v133 {
		v142 = v158
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v186 = F_LWLockAcquire(m, v182+int32(3584), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	if v188&int32(8) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v284+int32(3584))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L71
	}
L49:
	;
	if v188&int32(1) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v193 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v212 = v188
	goto L52
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v212&int32(8) != 0 {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v77)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v196))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v195)) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v208 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L55:
	;
	v208 = base.B2i32(base.Ui32(v195) < base.Ui32(v196))
	goto L54
L56:
	;
	goto L57
L57:
	;
	v208 = int32(base.Ui32(v195-v196) >> (uint(int32(31)) % 32))
	goto L54
L58:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v77)+108))
	v212 = v211
	goto L52
L59:
	;
	F_FlagRWConflict(m, v77, v214)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L70
	}
L60:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+108)))
	if v217&int32(8) != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
	if v220 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v224 = v77 + int32(32)
	if v220 == v224 {
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v214)+44))
	if v226 == int32(0) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	if v226 == v214+int32(40) {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v233 = v220
	goto L66
L66:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	if v247 == v214 {
		goto L48
	} else {
		goto L68
	}
L67:
	;
	goto L59
L68:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v249 != v224 {
		v233 = v249
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L48
L71:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v294 = F_LWLockAcquire(m, v290+int32(3584), int32(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v301 = v64
	v309 = v72
	goto L18
L73:
	;
	goto L17
L74:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v319 == int32(0) {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v339 = int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v345 = F_LWLockAcquire(m, v341+int32(3840), v339)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+72))
	if v352 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v354&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v354 = int32(1)
	goto L81
L80:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+76)))
	v354 = v353
	goto L81
L81:
	;
	goto L78
L82:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	v362 = F_LWLockAcquire(m, v358+int32(72), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v365 = F_LWLockAcquire(m, v32, int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v372 = F_LWLockAcquire(m, v368+int32(3584), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[1112]))
	v378 = v316 ^ v22
	v379 = int32(0)
	v381 = F_hash_search_with_hash_value(m, v375, v18+int32(8), v378, v379, v379)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v381 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+4)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+4)) = v389
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v391
	v394 = *(*int32)(unsafe.Add(mBase, _consts[1112]))
	v399 = F_hash_search_with_hash_value(m, v394, v18+int32(8), v378, int32(2), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	v416 = v339
	goto L91
L91:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v418+int32(3584))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L100
	}
L92:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v401 != v54 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v404 = v401
	goto L95
L94:
	;
	v404 = int32(0)
	goto L95
L95:
	;
	if v404 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[1109]))
	v411 = F_hash_search_with_hash_value(m, v408, v40, v22, int32(2), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v416 = base.B2i32(v399 == int32(0))
	goto L91
L99:
	;
	goto L98
L100:
	;
	F_LWLockRelease(m, v32)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+72))
	if v430 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v432&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v432 = int32(1)
	goto L105
L104:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+76)))
	v432 = v431
	goto L105
L105:
	;
	goto L102
L106:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	F_LWLockRelease(m, v436+int32(72))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v442+int32(3840))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	if v416 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _consts[1183]))
	v451 = F_hash_search_with_hash_value(m, v448, l0, v22, int32(2), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_DecrementParentLocks(m, l0)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v5 == v3 {
		v36 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v36
L2:
	;
	if l1 == int32(0) {
		v36 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v36 = v3
	goto L1
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v14<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v5 == v23 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v36 = int32(1)
	goto L1
L10:
	;
	goto L11
L11:
	;
	v27 = v14 + int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v27 < v28 {
		v14 = v27
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
