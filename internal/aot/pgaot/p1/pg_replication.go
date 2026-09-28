package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_get_replication_slots(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v116 int64
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int64
	_ = v257
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int64
	_ = v295
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int64
	_ = v406
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int64
	_ = v489
	var v491 int64
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(496)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v28 = int64(0)
	v30 = int32(_a_F_pg_get_replication_slots_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v35 = base.AtomicRmwCmpxchg64(m, v31, int32(272), v28, v28)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v35
	v37 = int32(0)
	v40 = base.AtomicRmwOr32(m, v37, int32(_a_F_pg_get_replication_slots_1), v37)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v47 = base.AtomicRmwCmpxchg64(m, v43, int32(264), v28, v28)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v47
	goto L3
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	v54 = F_LWLockAcquire(m, v50+int32(_a_F_pg_get_replication_slots_2), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	if int32(0) < v57+v59 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[6]))
	v75 = v57
	v77 = v59
	v78 = v68
	v82 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	F_LWLockRelease(m, v556+int32(_a_F_pg_get_replication_slots_2))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L122
	}
L8:
	;
	v94 = v78 + v82*int32(296)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v95 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v100 = base.AtomicRmwXchg32(m, v94, int32(0), int32(1))
	if v100 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v525 = v75
	v526 = v77
	v527 = v78
	goto L12
L12:
	;
	v535 = v82 + int32(1)
	if v535 < v525+v526 {
		v75 = v525
		v77 = v526
		v78 = v527
		v82 = v535
		goto L8
	} else {
		goto L121
	}
L13:
	;
	F_s_lock(m, v94, int32(_a_F_pg_get_replication_slots_3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	base.MemoryCopy(m, v20+int32(200), v94, int32(296))
	v108 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94))), uint32(v108))
	base.MemoryFill(m, v20+int32(32), v108, int32(168))
	v116 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+13)) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = base.I64_extend_i32_u(v20 + int32(224))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v20)+288))
	if v123 == v108 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v132 = F_cstring_to_text(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v126)
	v131 = int32(_a_F_pg_get_replication_slots_4)
	goto L17
L19:
	;
	goto L20
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = base.I64_extend_i32_u(v20 + int32(337))
	v131 = int32(_a_F_pg_get_replication_slots_5)
	goto L17
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = base.I64_extend_i32_u(v132)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v20)+288))
	if v136 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v20)+292))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+64)) = base.I64_extend_i32_u(base.B2i32(v143 == int32(2)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
	v150 = base.B2i32(v148 != int32(-1))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+72)) = base.I64_extend_i32_u(v150)
	if v148 != int32(-1) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)) = uint8(v139)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = base.I64_extend_i32_u(v136)
	goto L22
L26:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v20)+296))
	if v165 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v161 = int64(*(*int32)(unsafe.Add(mBase, uint32(v157+v148*int32(768))+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = v161
	goto L26
L28:
	;
	goto L29
L29:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)) = uint8(v163)
	goto L26
L30:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v20)+300))
	if v170 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = base.I64_extend_i32_u(v165)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)) = uint8(v168)
	goto L30
L34:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v20)+304))
	if v175 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+96)) = base.I64_extend_i32_u(v170)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v173)
	goto L34
L38:
	;
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v20)+320))
	if v181 != int64(0) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+104)) = v175
	goto L38
L40:
	;
	goto L41
L41:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)) = uint8(v179)
	goto L38
L42:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+312))
	if v187 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v181
	goto L42
L44:
	;
	goto L45
L45:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)) = uint8(v185)
	goto L42
L46:
	;
	v395 = v20 + int32(32)
	v398 = v395 + v385<<(uint(int32(3))%32)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+336)))
	*(*int64)(unsafe.Add(mBase, uint32(v398)+8)) = base.I64_extend_i32_u(v399)
	v403 = v385 + int32(2)
	if v399 != int32(1) {
		goto L94
	} else {
		goto L95
	}
L47:
	;
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v373|v20))) = uint8(v383)
	v385 = v373
	goto L46
L48:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[8]))
	if v350 < int32(0) {
		v373 = v341
		goto L47
	} else {
		goto L89
	}
L49:
	;
	v341 = int32(12)
	v342 = v20 + int32(128)
	goto L48
L50:
	;
	v327 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_6))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L88
	}
L51:
	;
	v304 = v175
	goto L53
L52:
	;
	v189 = m.G0
	v191 = v189 - int32(16)
	m.G0 = v191
	if v175 == int64(0) {
		v274 = int32(0)
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v304 == int64(0) {
		goto L50
	} else {
		goto L81
	}
L54:
	;
	m.G0 = v191 + int32(16)
	switch v274 {
	case 0:
		goto L77
	case 1:
		goto L76
	case 2:
		goto L75
	case 3:
		goto L74
	case 4:
		goto L73
	default:
		v341 = int32(11)
		v342 = v20 + int32(120)
		goto L48
	}
L55:
	;
	v197 = int32(_a_F_pg_get_replication_slots_0)
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v199 = int64(0)
	v202 = base.AtomicRmwCmpxchg64(m, v198, int32(272), v199, v199)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v202
	v204 = int32(0)
	v207 = base.AtomicRmwOr32(m, v204, int32(_a_F_pg_get_replication_slots_1), v204)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v214 = base.AtomicRmwCmpxchg64(m, v210, int32(264), v199, v199)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v214
	v217 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[9])))
	v218 = base.I64_div_u_s(v214, v217)
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v218
	F_KeepLogSeg(m, v214, v191+int32(8))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v228 = base.AtomicRmwXchg32(m, v225, int32(440), int32(1))
	if v228 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_s_lock(m, v225+int32(440), int32(_a_F_pg_get_replication_slots_7))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v235)+224))
	v237 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v235)+440)), uint32(v237))
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[9]))
	v242 = base.I64_extend_i32_s(v241)
	v243 = base.I64_div_u_s(v175, v242)
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v191)+8))
	if base.Ui64(v244) <= base.Ui64(v243) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v247 = int32(1)
	v249 = base.I64_div_u_s(v214, v242)
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[10]))
	v253 = base.I32_div_s(v241, int32(_a_F_pg_get_replication_slots_8))
	v254 = base.I32_div_s(v251, v253)
	v257 = base.I64_extend_i32_s(v254 + v247)
	if base.Ui64(v249) <= base.Ui64(v257) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	if base.Ui64(v243) < base.Ui64(v236+int64(1)) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v260 = int64(1)
	goto L66
L65:
	;
	v260 = v249 - v257
	goto L66
L66:
	;
	if base.Ui64(v243) < base.Ui64(v260) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v262 = int32(2)
	goto L69
L68:
	;
	v262 = v247
	goto L69
L69:
	;
	v274 = v262
	goto L54
L70:
	;
	v268 = int32(4)
	goto L72
L71:
	;
	v268 = int32(3)
	goto L72
L72:
	;
	v274 = v268
	goto L54
L73:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v20)+304))
	v304 = v295
	goto L53
L74:
	;
	v291 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_9))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L80
	}
L75:
	;
	v286 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_10))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L79
	}
L76:
	;
	v281 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_11))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v278 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+11)) = uint8(v278)
	goto L49
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = base.I64_extend_i32_u(v281)
	goto L49
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = base.I64_extend_i32_u(v286)
	goto L49
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = base.I64_extend_i32_u(v291)
	goto L49
L81:
	;
	v309 = base.AtomicRmwXchg32(m, v94, int32(0), int32(1))
	if v309 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_s_lock(m, v94, int32(_a_F_pg_get_replication_slots_3))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v94)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+304)) = v314
	v316 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v94))), uint32(v316))
	if v313 == int32(-1) {
		goto L50
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v322 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_9))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = base.I64_extend_i32_u(v322)
	goto L49
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = base.I64_extend_i32_u(v327)
	v373 = int32(12)
	goto L47
L89:
	;
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v20)+304))
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[9]))
	v356 = base.I64_extend_i32_s(v355)
	v357 = base.I64_div_u_s(v353, v356)
	v359 = base.I32_div_s(v355, int32(_a_F_pg_get_replication_slots_8))
	v360 = base.I32_div_s(v350, v359)
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[11]))
	v363 = base.I32_div_s(v362, v359)
	if base.Ui32(v363) < base.Ui32(v360) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v365 = v360
	goto L92
L91:
	;
	v365 = v363
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = (v357+base.I64_extend_i32_s(v365)+int64(1))*v356 - v47
	v385 = v341
	goto L46
L93:
	;
	v419 = v385 + int32(3)
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v20)+472))
	if int64(0) < v420 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v403))) = uint8(v415)
	goto L93
L95:
	;
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v20)+328))
	if v406 == int64(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v403<<(uint(int32(3))%32)+v395))) = v406
	goto L93
L97:
	;
	v433 = v385 + int32(4)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v20)+312))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v20)+288))
	if v437 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(32)+v419<<(uint(int32(3))%32)))) = v420
	goto L97
L99:
	;
	goto L100
L100:
	;
	v430 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v419))) = uint8(v430)
	goto L97
L101:
	;
	v489 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v20)+402)))
	*(*int64)(unsafe.Add(mBase, uint32(v398)+48)) = v489
	v491 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v20)+401)))
	*(*int64)(unsafe.Add(mBase, uint32(v398)+56)) = v491
	v494 = v385 + int32(8)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v20)+488))
	if v495 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L102:
	;
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(v434)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v434)) != 0 {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v447))) = int64(1)
	v463 = v385 + int32(5)
	goto L102
L104:
	;
	v454 = v385 + int32(5)
	if v434 != 0 {
		v463 = v454
		goto L102
	} else {
		goto L109
	}
L105:
	;
	v441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v433))) = uint8(v441)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v447 = v20 + int32(32) + v433<<(uint(int32(3))%32)
	switch v434 - int32(2) {
	case 0, 2:
		goto L103
	default:
		goto L108
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v447))) = int64(0)
	goto L104
L109:
	;
	v456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v454))) = uint8(v456)
	goto L101
L110:
	;
	v484 = F_cstring_to_text(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L114
	}
L111:
	;
	v483 = int32(_a_F_pg_get_replication_slots_12)
	goto L113
L112:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v434<<(uint(int32(2))%32))+uint32(_c_F_pg_get_replication_slots[12])))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	v483 = v482
	goto L113
L113:
	;
	goto L110
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(32)+v463<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v484)
	goto L101
L115:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_tuplestore_putvalues(m, v513, v514, v20+int32(32), v20)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L120
	}
L116:
	;
	v499 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v494+v20))) = uint8(v499)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v495<<(uint(int32(2))%32))+uint32(_c_F_pg_get_replication_slots[13])))
	v509 = F_cstring_to_text(m, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(32)+v494<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v509)
	goto L115
L120:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[6]))
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	v525 = v524
	v526 = v520
	v527 = v522
	goto L12
L121:
	;
	goto L9
L122:
	;
	m.G0 = v20 + int32(496)
	return int64(0)
}
func F_pg_replication_origin_session_is_setup(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_is_setup[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_is_setup[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_is_setup[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_replication_origin_session_is_setup_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_replication_origin_session_is_setup_1), int32(217), int32(_a_F_pg_replication_origin_session_is_setup_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		v34 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_is_setup[2])))
		return base.I64_extend_i32_u(base.B2i32(v34 != int32(0)))
	}
}
