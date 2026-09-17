package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v152 int32
	_ = v152
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
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v275 int64
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v323 int64
	_ = v323
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int64
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
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
	v30 = int32(_a_F_pg_get_replication_slots_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v31)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v32
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v32
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+272)) = v38
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v38
	goto L3
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	v47 = F_LWLockAcquire(m, v43+int32(_a_F_pg_get_replication_slots_1), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	if int32(0) < v50 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v63 = v58
	v66 = v50
	v71 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[3]))
	F_LWLockRelease(m, v511+int32(_a_F_pg_get_replication_slots_1))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L122
	}
L8:
	;
	v82 = v63 + v71*int32(288)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)))
	if v83 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(1)
	if v86 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v481 = v63
	v483 = v66
	goto L12
L12:
	;
	v491 = v71 + int32(1)
	if v491 < v483 {
		v63 = v481
		v66 = v483
		v71 = v491
		goto L8
	} else {
		goto L121
	}
L13:
	;
	F_s_lock(m, v82, int32(_a_F_pg_get_replication_slots_2), int32(268), int32(_a_F_pg_get_replication_slots_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	base.MemoryCopy(m, v20+int32(112), v82, int32(288))
	v98 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v98
	base.MemoryFill(m, v20+int32(32), v98, int32(80))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v98
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v20 + int32(136)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v112 == v98 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v121 = F_cstring_to_text(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v115)
	v120 = int32(_a_F_pg_get_replication_slots_4)
	goto L17
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v20 + int32(249)
	v120 = int32(_a_F_pg_get_replication_slots_5)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v124 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = base.B2i32(v130 == int32(2))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = base.B2i32(v134 != int32(0))
	if v134 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v127 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)) = uint8(v127)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v124
	goto L22
L26:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
	if v141 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v134
	goto L26
L28:
	;
	goto L29
L29:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)) = uint8(v139)
	goto L26
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	if v145 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v141
	goto L30
L32:
	;
	goto L33
L33:
	;
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+7)) = uint8(v143)
	goto L30
L34:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	if v149 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v145
	goto L34
L36:
	;
	goto L37
L37:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v147)
	goto L34
L38:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v20)+232))
	if v157 != int64(0) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v152 = F_Int64GetDatum(m, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v155 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+9)) = uint8(v155)
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v152
	goto L38
L43:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	if v166 != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v160 = F_Int64GetDatum(m, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+10)) = uint8(v163)
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v160
	goto L43
L48:
	;
	v367 = v20 + int32(32)
	v368 = int32(2)
	v370 = v367 + v359<<(uint(v368)%32)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v370)+4)) = v371
	v374 = v359 + v368
	if v371 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L49:
	;
	v355 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20|v347))) = uint8(v355)
	v359 = v347
	goto L48
L50:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[6]))
	if v320 < int32(0) {
		v347 = v312
		goto L49
	} else {
		goto L91
	}
L51:
	;
	v311 = v20 + int32(80)
	v312 = int32(12)
	goto L50
L52:
	;
	v298 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_6))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L90
	}
L53:
	;
	v275 = v165
	goto L55
L54:
	;
	v168 = m.G0
	v170 = v168 - int32(16)
	m.G0 = v170
	if v165 == int64(0) {
		v248 = int32(0)
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v275 == int64(0) {
		goto L52
	} else {
		goto L83
	}
L56:
	;
	m.G0 = v170 + int32(16)
	switch v248 {
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
		v311 = v20 + int32(76)
		v312 = int32(11)
		goto L50
	}
L57:
	;
	v175 = int32(_a_F_pg_get_replication_slots_0)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v176)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v176)+280)) = v177
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[1])) = v177
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+272)) = v183
	*(*int64)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[2])) = v183
	v188 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7])))
	v189 = base.I64_div_u_s(v183, v188)
	*(*int64)(unsafe.Add(mBase, uint32(v170)+8)) = v189
	F_KeepLogSeg(m, v183, v170+int32(8))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+440)) = int32(1)
	if v197 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	F_s_lock(m, v201+int32(440), int32(_a_F_pg_get_replication_slots_7), int32(3760), int32(_a_F_pg_get_replication_slots_8))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+440)) = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v215 = base.I64_extend_i32_s(v214)
	v216 = base.I64_div_u_s(v165, v215)
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
	if base.Ui64(v217) <= base.Ui64(v216) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v220 = int32(1)
	v222 = base.I64_div_u_s(v183, v215)
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[8]))
	v226 = base.I32_div_s(v214, int32(_a_F_pg_get_replication_slots_9))
	v227 = base.I32_div_s(v224, v226)
	v230 = base.I64_extend_i32_s(v227 + v220)
	if base.Ui64(v222) <= base.Ui64(v230) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v210)+232))
	if base.Ui64(v216) < base.Ui64(v238+int64(1)) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	v233 = int64(1)
	goto L68
L67:
	;
	v233 = v222 - v230
	goto L68
L68:
	;
	if base.Ui64(v216) < base.Ui64(v233) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v235 = int32(2)
	goto L71
L70:
	;
	v235 = v220
	goto L71
L71:
	;
	v248 = v235
	goto L56
L72:
	;
	v242 = int32(4)
	goto L74
L73:
	;
	v242 = int32(3)
	goto L74
L74:
	;
	v248 = v242
	goto L56
L75:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v275 = v266
	goto L55
L76:
	;
	v263 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_10))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	v259 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_11))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v255 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_12))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v252 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+11)) = uint8(v252)
	goto L51
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v255
	goto L51
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v259
	goto L51
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v263
	goto L51
L83:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(1)
	if v278 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_s_lock(m, v82, int32(_a_F_pg_get_replication_slots_2), int32(364), int32(_a_F_pg_get_replication_slots_3))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v82)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+216)) = v287
	v289 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v289
	if v286 == v289 {
		goto L52
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v294 = F_cstring_to_text(m, int32(_a_F_pg_get_replication_slots_10))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v294
	goto L51
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v298
	v347 = int32(12)
	goto L49
L91:
	;
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v20)+216))
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[7]))
	v326 = base.I64_extend_i32_s(v325)
	v327 = base.I64_div_u_s(v323, v326)
	v329 = base.I32_div_s(v325, int32(_a_F_pg_get_replication_slots_9))
	v330 = base.I32_div_s(v320, v329)
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[9]))
	v333 = base.I32_div_s(v332, v329)
	if base.Ui32(v333) < base.Ui32(v330) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v335 = v330
	goto L94
L93:
	;
	v335 = v333
	goto L94
L94:
	;
	v342 = F_Int64GetDatum(m, (v327+base.I64_extend_i32_s(v335)+int64(1))*v326-v38)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v342
	v359 = v312
	goto L48
L96:
	;
	v392 = v359 + int32(3)
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v20)+384))
	if int64(0) < v393 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	v388 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v374+v20))) = uint8(v388)
	goto L96
L98:
	;
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v20)+240))
	if v377 == int64(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v383 = F_Int64GetDatum(m, v377)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374<<(uint(int32(2))%32)+v367))) = v383
	goto L96
L101:
	;
	v408 = v359 + int32(4)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v20)+224))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v20)+200))
	if v412 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L102:
	;
	v401 = F_Int64GetDatum(m, v393)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v20))) = uint8(v405)
	goto L101
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)+v392<<(uint(int32(2))%32)))) = v401
	goto L101
L106:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+314)))
	*(*int32)(unsafe.Add(mBase, uint32(v370)+24)) = v465
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+313)))
	*(*int32)(unsafe.Add(mBase, uint32(v370)+28)) = base.B2i32(v467 != int32(0))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_tuplestore_putvalues(m, v471, v472, v20+int32(32), v20)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L120
	}
L107:
	;
	if base.B2i32(int32(base.Ui32(int32(279))>>(uint(v409)%32))&int32(1) == int32(0))|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v409)) != 0 {
		goto L116
	} else {
		goto L117
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = int32(1)
	v439 = v359 + int32(5)
	goto L107
L109:
	;
	v429 = v359 + int32(5)
	if v409 != 0 {
		v439 = v429
		goto L107
	} else {
		goto L114
	}
L110:
	;
	v416 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v408))) = uint8(v416)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v420 = int32(2)
	v422 = v20 + int32(32) + v408<<(uint(v420)%32)
	switch v409 - v420 {
	case 0, 2:
		goto L108
	default:
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422))) = int32(0)
	goto L109
L114:
	;
	v431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v429))) = uint8(v431)
	goto L106
L115:
	;
	v460 = F_cstring_to_text(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L119
	}
L116:
	;
	v459 = int32(_a_F_pg_get_replication_slots_13)
	goto L118
L117:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v409<<(uint(int32(2))%32))+uint32(_c_F_pg_get_replication_slots[10])))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	v459 = v458
	goto L118
L118:
	;
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(32)+v439<<(uint(int32(2))%32)))) = v460
	goto L106
L120:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[4]))
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_replication_slots[5]))
	v481 = v480
	v483 = v478
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
