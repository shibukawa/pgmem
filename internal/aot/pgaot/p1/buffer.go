package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferLockAcquire(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v110 int64
	_ = v110
	var v126 int64
	_ = v126
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int64
	_ = v242
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int64
	_ = v272
	var v275 int64
	_ = v275
	var v281 int64
	_ = v281
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v331 int64
	_ = v331
	var v347 int64
	_ = v347
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v371 int64
	_ = v371
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v410 int64
	_ = v410
	var v412 int64
	_ = v412
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v471 int64
	_ = v471
	var v474 int64
	_ = v474
	var v481 int64
	_ = v481
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v628 int64
	_ = v628
	v3 = l2
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[0]))
	if v13 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v27 = int32(_a_F_BufferLockAcquire_0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[1])) = v29 + int32(1)
	v33 = int32(2)
	v46 = int32(0)
	goto L8
L2:
	;
	v17 = v13 << (uint(int32(4)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_BufferLockAcquire[2])))
	if v20 == l0 {
		v26 = v17 + int32(_a_F_BufferLockAcquire_1)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v24 = F_GetPrivateRefCountEntrySlow(m, l0, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	return
L7:
	;
	v26 = v24
	goto L1
L8:
	;
	v50 = int64(0)
	v53 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v50, v50)
	v57 = v53
	goto L10
L10:
	;
	switch v3 - v33 {
	case 0:
		goto L14
	case 1:
		goto L15
	default:
		goto L13
	}
L11:
	;
	if v86 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v89 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v57, v87)
	if v57 != v89 {
		v57 = v89
		goto L10
	} else {
		goto L25
	}
L13:
	;
	v84 = base.B2i32(v57&int64(9007199254740992) == int64(0))
	if v57&int64(9007199254740992) == int64(0) {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v77 = base.B2i32(v57&int64(13510798882111488) == int64(0))
	if v57&int64(13510798882111488) == int64(0) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v70 = base.B2i32(v57&int64(18014381329612800) == int64(0))
	if v57&int64(18014381329612800) == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = v57 | int64(9007199254740992)
	goto L18
L17:
	;
	v71 = v57
	goto L18
L18:
	;
	v86 = v70
	v87 = v71
	goto L12
L19:
	;
	v78 = v57 | int64(4503599627370496)
	goto L21
L20:
	;
	v78 = v57
	goto L21
L21:
	;
	v86 = v77
	v87 = v78
	goto L12
L22:
	;
	v85 = v57 - int64(-17179869184)
	goto L24
L23:
	;
	v85 = v57
	goto L24
L24:
	;
	v86 = v84
	v87 = v85
	goto L12
L25:
	;
	goto L11
L26:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[3]))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(v33)%32))+uint32(_c_F_BufferLockAcquire[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v598))) = v599
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v603 = v602
	v610 = v46
	goto L145
L27:
	;
	v93 = m.G0
	v95 = v93 - int32(32)
	m.G0 = v95
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	if v98 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v3
	if int32(0) < v46 {
		goto L138
	} else {
		goto L139
	}
L30:
	;
	v272 = int64(0)
	v275 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v272, v272)
	v281 = v275
	goto L70
L31:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L67
	}
L32:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+344)))
	if v99 != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L64
	}
L35:
	;
	v100 = int64(4194304)
	v102 = base.AtomicRmwOr64(m, l1, int32(24), v100)
	if v102&v100 != int64(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v110 = v102
	goto L39
L37:
	;
	goto L38
L38:
	;
	v209 = base.AtomicRmwOr64(m, l1, int32(24), int64(4294967296))
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+345)) = uint8(v3)
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+344)) = uint8(v213)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[6]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[7]))
	v222 = v217 + v219*int32(768)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v223 == int32(-1) {
		goto L61
	} else {
		goto L62
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = int32(_a_F_BufferLockAcquire_2)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+24)) = int32(_a_F_BufferLockAcquire_3)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(_a_F_BufferLockAcquire_4)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = int32(0)
	v126 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v126
	if v110&int64(4194304) != v126 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	goto L44
L42:
	;
	goto L43
L43:
	;
	v169 = int32(_a_F_BufferLockAcquire_5)
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[8]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(8))+8))
	if v172 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	F_perform_spin_delay(m, v95+int32(8))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	v147 = int64(0)
	v150 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v147, v147)
	if v150&int64(4194304) != v147 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v189 = int64(4194304)
	v191 = base.AtomicRmwOr64(m, l1, int32(24), v189)
	if v191&v189 != int64(0) {
		v110 = v191
		goto L39
	} else {
		goto L59
	}
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[8])) = v187
	goto L49
L51:
	;
	if int32(999) < v170 {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v170 < int32(11) {
		goto L49
	} else {
		goto L58
	}
L54:
	;
	v177 = int32(900)
	if v177 <= v170 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v180 = v177
	goto L57
L56:
	;
	v180 = v170
	goto L57
L57:
	;
	v187 = v180 + int32(100)
	goto L50
L58:
	;
	v187 = v170 - int32(1)
	goto L50
L59:
	;
	goto L40
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v219
	v242 = base.AtomicRmwSub64(m, l1, int32(24), int64(4194304))
	m.G0 = v95 + int32(32)
	goto L30
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v222)+348)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v219
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+352)) = v223
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[6]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v232+v223*int32(768))+348)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v222)+348)) = int32(-1)
	goto L60
L64:
	;
	F_errmsg_internal(m, int32(_a_F_BufferLockAcquire_6), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_BufferLockAcquire_4), int32(_a_F_BufferLockAcquire_7), int32(_a_F_BufferLockAcquire_8))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_BufferLockAcquire_9), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_BufferLockAcquire_4), int32(_a_F_BufferLockAcquire_10), int32(_a_F_BufferLockAcquire_8))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	switch v3 - int32(2) {
	case 0:
		goto L74
	case 1:
		goto L75
	default:
		goto L73
	}
L71:
	;
	if v310 == int32(0) {
		goto L26
	} else {
		goto L86
	}
L72:
	;
	v313 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v281, v311)
	if v281 != v313 {
		v281 = v313
		goto L70
	} else {
		goto L85
	}
L73:
	;
	v308 = base.B2i32(v281&int64(9007199254740992) == int64(0))
	if v281&int64(9007199254740992) == int64(0) {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v301 = base.B2i32(v281&int64(13510798882111488) == int64(0))
	if v281&int64(13510798882111488) == int64(0) {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v294 = base.B2i32(v281&int64(18014381329612800) == int64(0))
	if v281&int64(18014381329612800) == int64(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v295 = v281 | int64(9007199254740992)
	goto L78
L77:
	;
	v295 = v281
	goto L78
L78:
	;
	v310 = v294
	v311 = v295
	goto L72
L79:
	;
	v302 = v281 | int64(4503599627370496)
	goto L81
L80:
	;
	v302 = v281
	goto L81
L81:
	;
	v310 = v301
	v311 = v302
	goto L72
L82:
	;
	v309 = v281 - int64(-17179869184)
	goto L84
L83:
	;
	v309 = v281
	goto L84
L84:
	;
	v310 = v308
	v311 = v309
	goto L72
L85:
	;
	goto L71
L86:
	;
	v317 = m.G0
	v319 = v317 - int32(32)
	m.G0 = v319
	v321 = int64(4194304)
	v323 = base.AtomicRmwOr64(m, l1, int32(24), v321)
	if v323&v321 != int64(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v331 = v323
	goto L90
L88:
	;
	goto L89
L89:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+344)))
	if v430 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+28)) = int32(_a_F_BufferLockAcquire_2)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+24)) = int32(_a_F_BufferLockAcquire_3)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+20)) = int32(_a_F_BufferLockAcquire_4)
	*(*int32)(unsafe.Add(mBase, uint32(v319)+16)) = int32(0)
	v347 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v319)+8)) = v347
	if v331&int64(4194304) != v347 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L89
L92:
	;
	goto L95
L93:
	;
	goto L94
L94:
	;
	v390 = int32(_a_F_BufferLockAcquire_5)
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[8]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v319+int32(8))+8))
	if v393 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	F_perform_spin_delay(m, v319+int32(8))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L97
	}
L96:
	;
	goto L94
L97:
	;
	v368 = int64(0)
	v371 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v368, v368)
	if v371&int64(4194304) != v368 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v410 = int64(4194304)
	v412 = base.AtomicRmwOr64(m, l1, int32(24), v410)
	if v412&v410 != int64(0) {
		v331 = v412
		goto L90
	} else {
		goto L110
	}
L100:
	;
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[8])) = v408
	goto L100
L102:
	;
	if int32(999) < v391 {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v391 < int32(11) {
		goto L100
	} else {
		goto L109
	}
L105:
	;
	v398 = int32(900)
	if v398 <= v391 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v401 = v398
	goto L108
L107:
	;
	v401 = v391
	goto L108
L108:
	;
	v408 = v401 + int32(100)
	goto L101
L109:
	;
	v408 = v391 - int32(1)
	goto L101
L110:
	;
	goto L91
L111:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[6]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[7]))
	v440 = v435 + v437*int32(768)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+348))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v440)+352))
	if v442 == int32(-1) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v468 != int32(-1) {
		goto L122
	} else {
		goto L123
	}
L114:
	;
	if v441 == int32(-1) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v441
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v440)+352))
	v451 = v446
	goto L114
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v435+v442*int32(768))+348)) = v441
	v451 = v442
	goto L114
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v440)+348)) = int64(0)
	goto L113
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v451
	goto L118
L120:
	;
	goto L121
L121:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[6]))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	*(*int32)(unsafe.Add(mBase, uint32(v457+v441*int32(768))+352)) = v451
	goto L118
L122:
	;
	v484 = base.AtomicRmwSub64(m, l1, int32(24), int64(4194304))
	if v430 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v471 = int64(0)
	v474 = base.AtomicRmwCmpxchg64(m, l1, int32(24), v471, v471)
	if v474&int64(4294967296) == v471 {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v481 = base.AtomicRmwAnd64(m, l1, int32(24), int64(-4294967297))
	goto L122
L125:
	;
	m.G0 = v319 + int32(32)
	goto L29
L126:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v489 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v488)+344)) = uint8(v489)
	goto L125
L127:
	;
	goto L128
L128:
	;
	v493 = base.AtomicRmwAnd64(m, l1, int32(24), int64(-8589934593))
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v497 = int32(0)
	v502 = v495
	goto L129
L129:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v502)+332))
	F_PGSemaphoreLock(m, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L131
	}
L130:
	;
	if v497 <= int32(0) {
		goto L125
	} else {
		goto L133
	}
L131:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+344)))
	if v515 != 0 {
		v497 = v497 + int32(1)
		v502 = v514
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v519 = v497
	goto L134
L134:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)+332))
	F_PGSemaphoreUnlock(m, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L6
	} else {
		goto L136
	}
L135:
	;
	goto L125
L136:
	;
	v534 = int32(1)
	if base.Ui32(v534) < base.Ui32(v519) {
		v519 = v519 - v534
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v573 = v46
	goto L141
L139:
	;
	goto L140
L140:
	;
	return
L141:
	;
	v578 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+332))
	F_PGSemaphoreUnlock(m, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L6
	} else {
		goto L143
	}
L142:
	;
	goto L140
L143:
	;
	v582 = int32(1)
	if base.Ui32(v582) < base.Ui32(v573) {
		v573 = v573 - v582
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v603)+332))
	F_PGSemaphoreLock(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L147
	}
L146:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = int32(0)
	v628 = base.AtomicRmwAnd64(m, l1, int32(24), int64(-8589934593))
	v46 = v610
	goto L8
L147:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockAcquire[5]))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+344)))
	if v621 != 0 {
		v603 = v620
		v610 = v610 + int32(1)
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
}
func F_BufferLockProcessRelease(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v43 int64
	_ = v43
	var v67 int64
	_ = v67
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int64
	_ = v300
	var v308 int32
	_ = v308
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v322 int64
	_ = v322
	var v325 int64
	_ = v325
	var v346 int32
	_ = v346
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v363 int32
	_ = v363
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	if l2&int64(12884901888) != int64(4294967296) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(32)
	return
L2:
	;
	v30 = l2 & int64(18014381329612800)
	if base.B2i32(l1 != int32(2))&base.B2i32(v30 != int64(0)) != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = int64(4194304)
	v36 = base.AtomicRmwOr64(m, l0, int32(24), v34)
	if v36&v34 != int64(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v43 = v36
	goto L7
L5:
	;
	goto L6
L6:
	;
	v169 = int32(-1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v170 == v169 {
		v308 = v169
		v317 = int64(0)
		goto L29
	} else {
		goto L30
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(_a_F_BufferLockProcessRelease_0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(_a_F_BufferLockProcessRelease_1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(_a_F_BufferLockProcessRelease_2)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(0)
	v67 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v67
	if v43&int64(4194304) != v67 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	goto L12
L10:
	;
	goto L11
L11:
	;
	v124 = int32(_a_F_BufferLockProcessRelease_3)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[0]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(8))+8))
	if v127 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	F_perform_spin_delay(m, v21+int32(8))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	return
L15:
	;
	v95 = int64(0)
	v98 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v95, v95)
	if v98&int64(4194304) != v95 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v144 = int64(4194304)
	v146 = base.AtomicRmwOr64(m, l0, int32(24), v144)
	if v146&v144 != int64(0) {
		v43 = v146
		goto L7
	} else {
		goto L28
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[0])) = v142
	goto L18
L20:
	;
	if int32(999) < v125 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v125 < int32(11) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v132 = int32(900)
	if v132 <= v125 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v135 = v132
	goto L26
L25:
	;
	v135 = v125
	goto L26
L26:
	;
	v142 = v135 + int32(100)
	goto L19
L27:
	;
	v142 = v125 - int32(1)
	goto L19
L28:
	;
	goto L8
L29:
	;
	v319 = int64(0)
	v322 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v319, v319)
	v325 = v322
	goto L61
L30:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v186 = int32(-1)
	v190 = v169
	v191 = v176
	v192 = v170
	v193 = int32(1)
	v194 = base.B2i32(v30 == int64(0))
	v195 = v177
	v196 = v177 + v170*int32(768)
	v198 = int32(0)
	goto L31
L31:
	;
	v202 = v192 * int32(768)
	v203 = v195 + v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v196)+348))
	if v194&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	if v290 != 0 {
		goto L58
	} else {
		goto L59
	}
L33:
	;
	if v204 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L34:
	;
	v285 = v278
	v286 = v193
	v287 = v280
	v288 = v281
	v289 = v282
	v290 = v283
	goto L33
L35:
	;
	if v193&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+345)))
	if v207 != int32(3) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v278 = v186
	v280 = int32(0)
	v281 = v190
	v282 = v191
	v283 = v198
	goto L34
L38:
	;
	v285 = v270
	v286 = v277
	v287 = v194
	v288 = v273
	v289 = v274
	v290 = v275
	goto L33
L39:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v203)+348))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v203)+352))
	if v218 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+345)))
	if v213 != int32(2) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v270 = v186
	v273 = v190
	v274 = v191
	v275 = v198
	v277 = int32(0)
	goto L38
L42:
	;
	if v217 == int32(-1) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v217
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v203)+352))
	v227 = v222
	goto L42
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195+v218*int32(768))+348)) = v217
	v227 = v218
	goto L42
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v203)+348)) = int64(0)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v243 = v242 + v202
	if v186 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v227
	goto L46
L48:
	;
	goto L49
L49:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	*(*int32)(unsafe.Add(mBase, uint32(v233+v217*int32(768))+352)) = v227
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+348)) = int32(-1)
	v259 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+344)) = uint8(v259)
	v261 = int32(0)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v265 = int32(1)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+345)))
	switch v267 - v265 {
	case 0:
		v278 = v192
		v280 = v261
		v281 = v256
		v282 = v263
		v283 = v265
		goto L34
	case 1:
		v285 = v192
		v286 = v261
		v287 = v261
		v288 = v256
		v289 = v263
		v290 = v265
		goto L33
	case 2:
		v308 = v256
		v317 = int64(8589934592)
		goto L29
	default:
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+352)) = int32(-1)
	v256 = v192
	goto L50
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+352)) = v186
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v251+v186*int32(768))+348)) = v192
	v256 = v190
	goto L50
L54:
	;
	v270 = v192
	v273 = v256
	v274 = v263
	v275 = v265
	v277 = v193
	goto L38
L55:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v186 = v285
	v190 = v288
	v191 = v289
	v192 = v204
	v193 = v286
	v194 = v287
	v195 = v294
	v196 = v294 + v204*int32(768)
	v198 = v290
	goto L31
L56:
	;
	goto L57
L57:
	;
	goto L32
L58:
	;
	v300 = int64(8589934592)
	goto L60
L59:
	;
	v300 = int64(0)
	goto L60
L60:
	;
	v308 = v288
	v317 = v300
	goto L29
L61:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v346 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v308 == int32(-1) {
		goto L1
	} else {
		goto L67
	}
L63:
	;
	v349 = int64(-4299161601)
	goto L65
L64:
	;
	v349 = int64(-4194305)
	goto L65
L65:
	;
	v352 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v325, (v325&int64(-8594128897)|v317)&v349)
	if v325 != v352 {
		v325 = v352
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v363 = v308
	goto L68
L68:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v379 = v376 + v363*int32(768)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+348))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v379)+352))
	if v381 != int32(-1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v379)+348)) = int64(0)
	v410 = int32(0)
	v413 = base.AtomicRmwOr32(m, v410, int32(_a_F_BufferLockProcessRelease_4), v410)
	*(*uint8)(unsafe.Add(mBase, uint32(v379)+344)) = uint8(v410)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v379)+332))
	F_PGSemaphoreUnlock(m, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L14
	} else {
		goto L77
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376+v381*int32(768))+348)) = v380
	goto L72
L71:
	;
	goto L72
L72:
	;
	if v380 != int32(-1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_BufferLockProcessRelease[1]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	*(*int32)(unsafe.Add(mBase, uint32(v392+v380*int32(768))+352)) = v381
	*(*int64)(unsafe.Add(mBase, uint32(v379)+348)) = int64(0)
	v399 = int32(0)
	v402 = base.AtomicRmwOr32(m, v399, int32(_a_F_BufferLockProcessRelease_4), v399)
	*(*uint8)(unsafe.Add(mBase, uint32(v379)+344)) = uint8(v399)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v379)+332))
	F_PGSemaphoreUnlock(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L14
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L69
L76:
	;
	v363 = v380
	goto L68
L77:
	;
	goto L1
}
func F_PinBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v44 int64
	_ = v44
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v105 int64
	_ = v105
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = v9 + int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[0]))
	if v13 != int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v138
L2:
	;
	v116 = int64(0)
	v119 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v116, v116)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = v120 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[1]))
	F_ResourceOwnerRemember(m, v125, base.I64_extend_i32_s(v11), int32(_a_F_PinBuffer_0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L32
	}
L3:
	;
	v17 = v13 << (uint(int32(4)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_PinBuffer[2])))
	if v20 == v11 {
		v115 = v17 + int32(_a_F_PinBuffer_1)
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v24 = F_GetPrivateRefCountEntrySlow(m, v11, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return int32(0)
L8:
	;
	if v24 != 0 {
		v115 = v24
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v28 = int32(0)
	v30 = int64(0)
	v33 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v30, v30)
	if v33&int64(16777216) == v30 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = l2
	goto L12
L11:
	;
	v38 = v28
	goto L12
L12:
	;
	if v38 != 0 {
		v138 = v28
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v44 = v33
	goto L14
L14:
	;
	if v44&int64(4194304) != int64(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v138 = v28
	goto L1
L16:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(v105&int64(16777216) != int64(0)) != 0 {
		v44 = v105
		goto L14
	} else {
		goto L31
	}
L17:
	;
	v51 = F_WaitBufHdrUnlocked(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v56 = v44 + int64(1)
	v58 = v56 & int64(3932160)
	if l1 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v105 = v51
	goto L16
L21:
	;
	v63 = base.B2i32(v58 == int64(0))
	goto L23
L22:
	;
	v63 = base.B2i32(base.Ui64(v58) < base.Ui64(int64(1310720)))
	goto L23
L23:
	;
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v64 = v44 + int64(262145)
	goto L26
L25:
	;
	v64 = v56
	goto L26
L26:
	;
	v66 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v44, v64)
	if v44 != v66 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = v66
	goto L16
L28:
	;
	goto L29
L29:
	;
	v68 = int32(_a_F_PinBuffer_2)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v69<<(uint(int32(2))%32))+uint32(_c_F_PinBuffer[4]))) = v11
	v76 = v69 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_PinBuffer[5]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_PinBuffer[2]))) = v11
	*(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[3])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[0])) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v76)+uint32(_c_F_PinBuffer[6]))) = int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[1]))
	F_ResourceOwnerRemember(m, v94, base.I64_extend_i32_s(v11), int32(_a_F_PinBuffer_0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	return base.I32_wrap_i64(int64(base.Ui64(v64&int64(16777216)) >> (uint(int64(24)) % 64)))
L31:
	;
	goto L15
L32:
	;
	v138 = base.I32_wrap_i64(int64(base.Ui64(v119&int64(16777216)) >> (uint(int64(24)) % 64)))
	goto L1
}
func F_ReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v171 int64
	_ = v171
	var v184 int64
	_ = v184
	var v191 int64
	_ = v191
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int64
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v269 int64
	_ = v269
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	v29 = F_smgropen(m, v17+int32(16), v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v48 = v19
	goto L3
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+118)))
	if v50 == int32(116) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = v47
	goto L3
L7:
	;
	v43 = v35
	goto L9
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	v43 = v41
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v43 + int32(1)
	goto L6
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L94
	}
L11:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v53 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v57 = l3 - int32(1)
	if l2 == int32(-1) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	m.G0 = v17 + int32(112)
	return v363
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v65
	if base.Ui32(v57) < base.Ui32(int32(2)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(v57) <= base.Ui32(int32(1)) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v71 = int32(9)
	goto L21
L20:
	;
	v71 = int32(1)
	goto L21
L21:
	;
	v72 = F_ExtendBufferedRel(m, v17, l1, l4, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v363 = v72
	goto L15
L23:
	;
	if v50 != int32(116) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+40)) = uint8(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v48
	v350 = v17 + int32(32)
	if l3 == int32(3) {
		goto L86
	} else {
		goto L87
	}
L26:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v323 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L27:
	;
	v265 = v258*int32(320) + v255<<(uint(int32(6))%32)
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_ReadBufferExtended[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_ReadBufferExtended[4]))) = v269 + int64(1)
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_ReadBufferExtended[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v265)+uint32(_c_F_ReadBufferExtended[5]))) = v273
	v275 = int32(1)
	F_pgstat_count_backend_io_op(m, v258, v255, int32(2), v275, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[6])) = uint8(v275)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[7])) = uint8(v275)
	goto L70
L28:
	;
	v245 = int32(_a_F_ReadBufferExtended_4)
	v247 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[3])) = v247 + int64(1)
	v254 = v242
	v255 = v78
	v258 = int32(0)
	goto L27
L29:
	;
	F_UnpinBuffer(m, v138)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L66
	}
L30:
	;
	v78 = F_IOContextForStrategy(m, l4)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v213 = F_LocalBufferAlloc(m, v48, l1, l2, v17+int32(28))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L64
	}
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[0]))
	F_ResourceOwnerEnlarge(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v90
	v95 = v17 + int32(32)
	v96 = F_BufTableHashCode(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[1]))
	v106 = v99 + v96&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ReadBufferExtended_3)
	v108 = F_LWLockAcquire(m, v106, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v110 = F_BufTableLookup(m, v95, v96)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if int32(0) <= v110 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[2]))
	v118 = v115 + v110*int32(56)
	v120 = F_PinBuffer(m, v118, l4, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_LWLockRelease(m, v106)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L45
	}
L42:
	;
	F_LWLockRelease(m, v106)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)) = uint8(v120)
	if v120 != 0 {
		v242 = v118
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v315 = v118
	goto L26
L45:
	;
	v127 = F_GetVictimBuffer(m, l4, v78)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[2]))
	v132 = F_LWLockAcquire(m, v106, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v134 = int32(56)
	v136 = v130 + v127*v134
	v138 = v136 - v134
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136-int32(36))))
	v144 = F_BufTableInsert(m, v17+int32(32), v96, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	if int32(0) <= v144 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	v148 = F_LockBufHdr(m, v138)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+16)) = v150
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v138)+8)) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v154
	v157 = v136 - int32(32)
	v158 = int64(2181300224)
	if v50 == int32(112) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v163 = v158
	goto L53
L52:
	;
	v163 = int64(33816576)
	goto L53
L53:
	;
	if l1 == int32(3) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v166 = v158
	goto L56
L55:
	;
	v166 = v163
	goto L56
L56:
	;
	v171 = base.AtomicRmwCmpxchg64(m, v157, int32(0), v148, v166|v148&int64(-38010881))
	if v148 != v171 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v184 = v171
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_LWLockRelease(m, v106)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L63
	}
L60:
	;
	v191 = base.AtomicRmwCmpxchg64(m, v157, int32(0), v184, v184&int64(-38010881)|v166)
	if v184 != v191 {
		v184 = v191
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	goto L61
L63:
	;
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)) = uint8(v209)
	v315 = v138
	goto L26
L64:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	if v215 == int32(0) {
		v315 = v213
		goto L26
	} else {
		goto L65
	}
L65:
	;
	v218 = int32(_a_F_ReadBufferExtended_6)
	v220 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[11])) = v220 + int64(1)
	v254 = v213
	v255 = int32(3)
	v258 = int32(1)
	goto L27
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[2]))
	v232 = v229 + v144*int32(56)
	v234 = F_PinBuffer(m, v232, l4, int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_LWLockRelease(m, v106)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)) = uint8(v234)
	if v234 == int32(0) {
		v315 = v232
		goto L26
	} else {
		goto L69
	}
L69:
	;
	v242 = v232
	goto L28
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[8])))
	if v285 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v288 = int32(_a_F_ReadBufferExtended_5)
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[9]))
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[9])) = v290 + v292
	goto L73
L72:
	;
	goto L73
L73:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v295 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v298 != int32(1) {
		v315 = v254
		goto L26
	} else {
		goto L77
	}
L75:
	;
	v304 = v295
	goto L76
L76:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v304)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v304)+120)) = v305 + int64(1)
	v315 = v254
	goto L26
L77:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v304 = v303
	goto L76
L79:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v315)+20))
	v340 = v338 + int32(1)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+28)))
	F_ZeroAndLockBuffer(m, v340, l3, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L85
	}
L80:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v326 != int32(1) {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	v332 = v323
	goto L82
L82:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v332)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v332)+112)) = v333 + int64(1)
	goto L79
L83:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v332 = v331
	goto L82
L85:
	;
	v363 = v340
	goto L15
L86:
	;
	v357 = int32(9)
	goto L88
L87:
	;
	v357 = int32(8)
	goto L88
L88:
	;
	v358 = F_StartReadBuffer(m, v350, v17+int32(28), l2, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v358 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v360 = F_WaitReadBuffers(m, v350)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v363 = v362
	goto L15
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_ReadBufferExtended_0), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ReadBufferExtended_1), int32(1296), int32(_a_F_ReadBufferExtended_2))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
