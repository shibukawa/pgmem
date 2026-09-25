package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_get_replication_slots(m *base.Module, l0 int32) int32 {
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
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v115 int64
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int64
	_ = v201
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v235 int64
	_ = v235
	var v238 int32
	_ = v238
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int64
	_ = v283
	var v292 int64
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v341 int64
	_ = v341
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int64
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(400)
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
	return int32(0)
L2:
	;
	v28 = int64(0)
	v30 = int32(_a_F_pg_get_replication_slots_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v35 = base.AtomicRmwCmpxchg64(m, v31, int32(280), v28, v28)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v35
	v37 = int32(0)
	v40 = base.AtomicRmwOr32(m, v37, int32(_a_F_pg_get_replication_slots_1), v37)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v47 = base.AtomicRmwCmpxchg64(m, v43, int32(272), v28, v28)
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
	if int32(0) < v57 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v70 = v65
	v73 = v57
	v77 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	F_LWLockRelease(m, v529+int32(_a_F_pg_get_replication_slots_2))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L122
	}
L8:
	;
	v89 = v70 + v77*int32(288)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v90 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v95 = base.AtomicRmwXchg32(m, v89, int32(0), int32(1))
	if v95 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v499 = v70
	v501 = v73
	goto L12
L12:
	;
	v509 = v77 + int32(1)
	if v509 < v501 {
		v70 = v499
		v73 = v501
		v77 = v509
		goto L8
	} else {
		goto L121
	}
L13:
	;
	F_s_lock(m, v89, int32(_a_F_pg_get_replication_slots_3), int32(268), int32(_a_F_pg_get_replication_slots_4))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	base.MemoryCopy(m, v20+int32(112), v89, int32(288))
	v105 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89))), uint32(v105))
	base.MemoryFill(m, v20+int32(32), v105, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v105
	v115 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v115
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v20 + int32(136)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v120 == v105 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v129 = F_cstring_to_text(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v123 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v123)
	v128 = int32(_a_F_pg_get_replication_slots_5)
	goto L17
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v20 + int32(249)
	v128 = int32(_a_F_pg_get_replication_slots_6)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v132 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = base.B2i32(v138 == int32(2))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = base.B2i32(v142 != int32(0))
	if v142 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)) = uint8(v135)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v132
	goto L22
L26:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
	if v149 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v142
	goto L26
L28:
	;
	goto L29
L29:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)) = uint8(v147)
	goto L26
L30:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	if v153 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v149
	goto L30
L32:
	;
	goto L33
L33:
	;
	v151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)) = uint8(v151)
	goto L30
L34:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	if v157 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v153
	goto L34
L36:
	;
	goto L37
L37:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v155)
	goto L34
L38:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v20)+232))
	if v165 != int64(0) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v160 = F_Int64GetDatum(m, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)) = uint8(v163)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v160
	goto L38
L43:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	if v174 != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v168 = F_Int64GetDatum(m, v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)) = uint8(v171)
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v168
	goto L43
L48:
	;
	v385 = v20 + int32(32)
	v386 = int32(2)
	v388 = v385 + v377<<(uint(v386)%32)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+4)) = v389
	v392 = v377 + v386
	if v389 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L49:
	;
	v373 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20|v365))) = uint8(v373)
	v377 = v365
	goto L48
L50:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[6]))
	if v338 < int32(0) {
		v365 = v330
		goto L49
	} else {
		goto L91
	}
L51:
	;
	v329 = v20 + int32(80)
	v330 = int32(12)
	goto L50
L52:
	;
	v316 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_7))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L90
	}
L53:
	;
	v292 = v173
	goto L55
L54:
	;
	v176 = m.G0
	v178 = v176 - int32(16)
	m.G0 = v178
	if v173 == int64(0) {
		v265 = int32(0)
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v292 == int64(0) {
		goto L52
	} else {
		goto L83
	}
L56:
	;
	m.G0 = v178 + int32(16)
	switch v265 {
	case 0:
		goto L79
	case 1:
		goto L78
	case 2:
		goto L77
	case 3:
		goto L76
	case 4:
		goto L75
	default:
		v329 = v20 + int32(76)
		v330 = int32(11)
		goto L50
	}
L57:
	;
	v184 = int32(_a_F_pg_get_replication_slots_0)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v186 = int64(0)
	v189 = base.AtomicRmwCmpxchg64(m, v185, int32(280), v186, v186)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v189
	v191 = int32(0)
	v194 = base.AtomicRmwOr32(m, v191, int32(_a_F_pg_get_replication_slots_1), v191)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v201 = base.AtomicRmwCmpxchg64(m, v197, int32(272), v186, v186)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v201
	v204 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7])))
	v205 = base.I64_div_u_s(v201, v204)
	*(*int64)(unsafe.Add(mBase, uint32(v178)+8)) = v205
	F_KeepLogSeg(m, v201, v178+int32(8))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v215 = base.AtomicRmwXchg32(m, v212, int32(440), int32(1))
	if v215 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	F_s_lock(m, v217+int32(440), int32(_a_F_pg_get_replication_slots_8), int32(3760), int32(_a_F_pg_get_replication_slots_9))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v226)+232))
	v228 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v226)+440)), uint32(v228))
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v233 = base.I64_extend_i32_s(v232)
	v234 = base.I64_div_u_s(v173, v233)
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v178)+8))
	if base.Ui64(v235) <= base.Ui64(v234) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v238 = int32(1)
	v240 = base.I64_div_u_s(v201, v233)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[8]))
	v244 = base.I32_div_s(v232, int32(_a_F_pg_get_replication_slots_10))
	v245 = base.I32_div_s(v242, v244)
	v248 = base.I64_extend_i32_s(v245 + v238)
	if base.Ui64(v240) <= base.Ui64(v248) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if base.Ui64(v234) < base.Ui64(v227+int64(1)) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	v251 = int64(1)
	goto L68
L67:
	;
	v251 = v240 - v248
	goto L68
L68:
	;
	if base.Ui64(v234) < base.Ui64(v251) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v253 = int32(2)
	goto L71
L70:
	;
	v253 = v238
	goto L71
L71:
	;
	v265 = v253
	goto L56
L72:
	;
	v259 = int32(4)
	goto L74
L73:
	;
	v259 = int32(3)
	goto L74
L74:
	;
	v265 = v259
	goto L56
L75:
	;
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v292 = v283
	goto L55
L76:
	;
	v280 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_11))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	v276 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_12))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v272 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_13))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v269 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+11)) = uint8(v269)
	goto L51
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v272
	goto L51
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v276
	goto L51
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v280
	goto L51
L83:
	;
	v297 = base.AtomicRmwXchg32(m, v89, int32(0), int32(1))
	if v297 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_s_lock(m, v89, int32(_a_F_pg_get_replication_slots_3), int32(364), int32(_a_F_pg_get_replication_slots_4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v89)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = v304
	v306 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89))), uint32(v306))
	if v303 == v306 {
		goto L52
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v312 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_11))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v312
	goto L51
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v316
	v365 = int32(12)
	goto L49
L91:
	;
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v344 = base.I64_extend_i32_s(v343)
	v345 = base.I64_div_u_s(v341, v344)
	v347 = base.I32_div_s(v343, int32(_a_F_pg_get_replication_slots_10))
	v348 = base.I32_div_s(v338, v347)
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[9]))
	v351 = base.I32_div_s(v350, v347)
	if base.Ui32(v351) < base.Ui32(v348) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v353 = v348
	goto L94
L93:
	;
	v353 = v351
	goto L94
L94:
	;
	v360 = F_Int64GetDatum(m, (v345+base.I64_extend_i32_s(v353)+int64(1))*v344-v47)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v360
	v377 = v330
	goto L48
L96:
	;
	v410 = v377 + int32(3)
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v20)+384))
	if int64(0) < v411 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v406 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v20))) = uint8(v406)
	goto L96
L98:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v20)+240))
	if v395 == int64(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v401 = F_Int64GetDatum(m, v395)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392<<(uint(int32(2))%32)+v385))) = v401
	goto L96
L101:
	;
	v426 = v377 + int32(4)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v430 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L102:
	;
	v419 = F_Int64GetDatum(m, v411)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v423 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v410+v20))) = uint8(v423)
	goto L101
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)+v410<<(uint(int32(2))%32)))) = v419
	goto L101
L106:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+314)))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+24)) = v483
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+313)))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+28)) = base.B2i32(v485 != int32(0))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_tuplestore_putvalues(m, v489, v490, v20+int32(32), v20)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L120
	}
L107:
	;
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(v427)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v427)) != 0 {
		goto L116
	} else {
		goto L117
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(1)
	v457 = v377 + int32(5)
	goto L107
L109:
	;
	v447 = v377 + int32(5)
	if v427 != 0 {
		v457 = v447
		goto L107
	} else {
		goto L114
	}
L110:
	;
	v434 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v426))) = uint8(v434)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v438 = int32(2)
	v440 = v20 + int32(32) + v426<<(uint(v438)%32)
	switch v427 - v438 {
	case 0, 2:
		goto L108
	default:
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(0)
	goto L109
L114:
	;
	v449 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v447))) = uint8(v449)
	goto L106
L115:
	;
	v478 = F_cstring_to_text(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L119
	}
L116:
	;
	v477 = int32(_a_F_pg_get_replication_slots_14)
	goto L118
L117:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v427<<(uint(int32(2))%32))+uint32(_c_F_pg_get_replication_slots[10])))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	v477 = v476
	goto L118
L118:
	;
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)+v457<<(uint(int32(2))%32)))) = v478
	goto L106
L120:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v499 = v498
	v501 = v496
	goto L12
L121:
	;
	goto L9
L122:
	;
	m.G0 = v20 + int32(400)
	return int32(0)
}
func F_pg_replication_origin_session_is_setup(m *base.Module, l0 int32) int32 {
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
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
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
			return int32(0)
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_replication_origin_session_is_setup_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_replication_origin_session_is_setup_1), int32(200), int32(_a_F_pg_replication_origin_session_is_setup_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		v34 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_is_setup[2])))
		return base.B2i32(v34 != int32(0))
	}
}
