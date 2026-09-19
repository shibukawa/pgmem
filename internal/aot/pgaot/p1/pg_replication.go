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
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v193 int64
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int64
	_ = v240
	var v243 int64
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int64
	_ = v275
	var v284 int64
	_ = v284
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int64
	_ = v333
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int64
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int64
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
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
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v43 = base.AtomicRmwCmpxchg64(m, v39, int32(272), v28, v28)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v43
	goto L3
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	v50 = F_LWLockAcquire(m, v46+int32(_a_F_pg_get_replication_slots_1), int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	if int32(0) < v53 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v66 = v61
	v69 = v53
	v73 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	F_LWLockRelease(m, v521+int32(_a_F_pg_get_replication_slots_1))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L122
	}
L8:
	;
	v85 = v66 + v73*int32(288)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v86 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v91 = base.AtomicRmwXchg32(m, v85, int32(0), int32(1))
	if v91 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v491 = v66
	v493 = v69
	goto L12
L12:
	;
	v501 = v73 + int32(1)
	if v501 < v493 {
		v66 = v491
		v69 = v493
		v73 = v501
		goto L8
	} else {
		goto L121
	}
L13:
	;
	F_s_lock(m, v85, int32(_a_F_pg_get_replication_slots_2), int32(268), int32(_a_F_pg_get_replication_slots_3))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	base.MemoryCopy(m, v20+int32(112), v85, int32(288))
	v101 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v85))), uint32(v101))
	base.MemoryFill(m, v20+int32(32), v101, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v101
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v20 + int32(136)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v116 == v101 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v125 = F_cstring_to_text(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v119 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v119)
	v124 = int32(_a_F_pg_get_replication_slots_4)
	goto L17
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v20 + int32(249)
	v124 = int32(_a_F_pg_get_replication_slots_5)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v128 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = base.B2i32(v134 == int32(2))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = base.B2i32(v138 != int32(0))
	if v138 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)) = uint8(v131)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v128
	goto L22
L26:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
	if v145 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v138
	goto L26
L28:
	;
	goto L29
L29:
	;
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)) = uint8(v143)
	goto L26
L30:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	if v149 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v145
	goto L30
L32:
	;
	goto L33
L33:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)) = uint8(v147)
	goto L30
L34:
	;
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	if v153 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v149
	goto L34
L36:
	;
	goto L37
L37:
	;
	v151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v151)
	goto L34
L38:
	;
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v20)+232))
	if v161 != int64(0) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v156 = F_Int64GetDatum(m, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)) = uint8(v159)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v156
	goto L38
L43:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	if v170 != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v164 = F_Int64GetDatum(m, v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)) = uint8(v167)
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v164
	goto L43
L48:
	;
	v377 = v20 + int32(32)
	v378 = int32(2)
	v380 = v377 + v369<<(uint(v378)%32)
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v381
	v384 = v369 + v378
	if v381 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L49:
	;
	v365 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20|v357))) = uint8(v365)
	v369 = v357
	goto L48
L50:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[6]))
	if v330 < int32(0) {
		v357 = v322
		goto L49
	} else {
		goto L91
	}
L51:
	;
	v321 = v20 + int32(80)
	v322 = int32(12)
	goto L50
L52:
	;
	v308 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_6))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L90
	}
L53:
	;
	v284 = v169
	goto L55
L54:
	;
	v172 = m.G0
	v174 = v172 - int32(16)
	m.G0 = v174
	if v169 == int64(0) {
		v257 = int32(0)
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v284 == int64(0) {
		goto L52
	} else {
		goto L83
	}
L56:
	;
	m.G0 = v174 + int32(16)
	switch v257 {
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
		v321 = v20 + int32(76)
		v322 = int32(11)
		goto L50
	}
L57:
	;
	v180 = int32(_a_F_pg_get_replication_slots_0)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v182 = int64(0)
	v185 = base.AtomicRmwCmpxchg64(m, v181, int32(280), v182, v182)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v185
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v193 = base.AtomicRmwCmpxchg64(m, v189, int32(272), v182, v182)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v193
	v196 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7])))
	v197 = base.I64_div_u_s(v193, v196)
	*(*int64)(unsafe.Add(mBase, uint32(v174)+8)) = v197
	F_KeepLogSeg(m, v193, v174+int32(8))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v207 = base.AtomicRmwXchg32(m, v204, int32(440), int32(1))
	if v207 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	F_s_lock(m, v209+int32(440), int32(_a_F_pg_get_replication_slots_7), int32(3760), int32(_a_F_pg_get_replication_slots_8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v218)+232))
	v220 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v218)+440)), uint32(v220))
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v225 = base.I64_extend_i32_s(v224)
	v226 = base.I64_div_u_s(v169, v225)
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
	if base.Ui64(v227) <= base.Ui64(v226) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v230 = int32(1)
	v232 = base.I64_div_u_s(v193, v225)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[8]))
	v236 = base.I32_div_s(v224, int32(_a_F_pg_get_replication_slots_9))
	v237 = base.I32_div_s(v234, v236)
	v240 = base.I64_extend_i32_s(v237 + v230)
	if base.Ui64(v232) <= base.Ui64(v240) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if base.Ui64(v226) < base.Ui64(v219+int64(1)) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	v243 = int64(1)
	goto L68
L67:
	;
	v243 = v232 - v240
	goto L68
L68:
	;
	if base.Ui64(v226) < base.Ui64(v243) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v245 = int32(2)
	goto L71
L70:
	;
	v245 = v230
	goto L71
L71:
	;
	v257 = v245
	goto L56
L72:
	;
	v251 = int32(4)
	goto L74
L73:
	;
	v251 = int32(3)
	goto L74
L74:
	;
	v257 = v251
	goto L56
L75:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v284 = v275
	goto L55
L76:
	;
	v272 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_10))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	v268 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_11))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v264 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_12))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+11)) = uint8(v261)
	goto L51
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v264
	goto L51
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v268
	goto L51
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v272
	goto L51
L83:
	;
	v289 = base.AtomicRmwXchg32(m, v85, int32(0), int32(1))
	if v289 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_s_lock(m, v85, int32(_a_F_pg_get_replication_slots_2), int32(364), int32(_a_F_pg_get_replication_slots_3))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = v296
	v298 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v85))), uint32(v298))
	if v295 == v298 {
		goto L52
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v304 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_10))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v304
	goto L51
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v308
	v357 = int32(12)
	goto L49
L91:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v336 = base.I64_extend_i32_s(v335)
	v337 = base.I64_div_u_s(v333, v336)
	v339 = base.I32_div_s(v335, int32(_a_F_pg_get_replication_slots_9))
	v340 = base.I32_div_s(v330, v339)
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[9]))
	v343 = base.I32_div_s(v342, v339)
	if base.Ui32(v343) < base.Ui32(v340) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v345 = v340
	goto L94
L93:
	;
	v345 = v343
	goto L94
L94:
	;
	v352 = F_Int64GetDatum(m, (v337+base.I64_extend_i32_s(v345)+int64(1))*v336-v43)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v352
	v369 = v322
	goto L48
L96:
	;
	v402 = v369 + int32(3)
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v20)+384))
	if int64(0) < v403 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v398 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v384+v20))) = uint8(v398)
	goto L96
L98:
	;
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v20)+240))
	if v387 == int64(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v393 = F_Int64GetDatum(m, v387)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384<<(uint(int32(2))%32)+v377))) = v393
	goto L96
L101:
	;
	v418 = v369 + int32(4)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v422 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L102:
	;
	v411 = F_Int64GetDatum(m, v403)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v402+v20))) = uint8(v415)
	goto L101
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)+v402<<(uint(int32(2))%32)))) = v411
	goto L101
L106:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+314)))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+24)) = v475
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+313)))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+28)) = base.B2i32(v477 != int32(0))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_tuplestore_putvalues(m, v481, v482, v20+int32(32), v20)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L120
	}
L107:
	;
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(v419)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v419)) != 0 {
		goto L116
	} else {
		goto L117
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432))) = int32(1)
	v449 = v369 + int32(5)
	goto L107
L109:
	;
	v439 = v369 + int32(5)
	if v419 != 0 {
		v449 = v439
		goto L107
	} else {
		goto L114
	}
L110:
	;
	v426 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v418))) = uint8(v426)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v430 = int32(2)
	v432 = v20 + int32(32) + v418<<(uint(v430)%32)
	switch v419 - v430 {
	case 0, 2:
		goto L108
	default:
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432))) = int32(0)
	goto L109
L114:
	;
	v441 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v439))) = uint8(v441)
	goto L106
L115:
	;
	v470 = F_cstring_to_text(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L119
	}
L116:
	;
	v469 = int32(_a_F_pg_get_replication_slots_13)
	goto L118
L117:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v419<<(uint(int32(2))%32))+uint32(_c_F_pg_get_replication_slots[10])))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v469 = v468
	goto L118
L118:
	;
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)+v449<<(uint(int32(2))%32)))) = v470
	goto L106
L120:
	;
	v488 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v491 = v490
	v493 = v488
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
