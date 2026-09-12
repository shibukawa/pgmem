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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int64
	_ = v236
	var v239 int64
	_ = v239
	var v241 int32
	_ = v241
	var v244 int64
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v281 int64
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v330 int64
	_ = v330
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int64
	_ = v384
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(400)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = int32(4366320)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+280)) = v33
	*(*int64)(unsafe.Add(mBase, _consts[225])) = v33
	v38 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+272)) = v39
	*(*int64)(unsafe.Add(mBase, _consts[226])) = v39
	goto L3
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v48 = F_LWLockAcquire(m, v44+int32(4736), int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	if int32(0) < v51 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[665]))
	v68 = v51
	v70 = v59
	v74 = v2
	goto L8
L6:
	;
	goto L7
L7:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v523+int32(4736))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L127
	}
L8:
	;
	v86 = v70 + v74*int32(288)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
	if v87 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(1)
	if v90 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v493 = v68
	v495 = v70
	goto L12
L12:
	;
	v502 = v74 + int32(1)
	if v502 < v493 {
		v68 = v493
		v70 = v495
		v74 = v502
		goto L8
	} else {
		goto L126
	}
L13:
	;
	F_s_lock(m, v86, int32(487009), int32(268), int32(116630))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L18
L16:
	;
	goto L15
L17:
	;
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v103
	v110 = F__emscripten_memset_bulkmem(m, v21+int32(32), base.I32_extend8_s(v103), int32(80))
	mBase = m.M
	goto L21
L18:
	;
	v101 = F__emscripten_memcpy_bulkmem(m, v21+int32(112), v86, int32(288))
	mBase = m.M
	goto L20
L20:
	;
	goto L17
L21:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(16)))) = v111
	v113 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v113
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v21 + int32(136)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v118 == v111 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v127 = F_cstring_to_text(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v121)
	v126 = int32(309520)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(249)
	v126 = int32(309606)
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v127
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v130 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v21)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = base.B2i32(v136 == int32(2))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v21)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = base.B2i32(v140 != int32(0))
	if v140 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+3)) = uint8(v133)
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v130
	goto L27
L31:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v21)+208))
	if v147 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v140
	goto L31
L33:
	;
	goto L34
L34:
	;
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v145)
	goto L31
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v21)+212))
	if v151 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v147
	goto L35
L37:
	;
	goto L38
L38:
	;
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v149)
	goto L35
L39:
	;
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v21)+216))
	if v155 != int64(0) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v151
	goto L39
L41:
	;
	goto L42
L42:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v153)
	goto L39
L43:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v21)+232))
	if v163 != int64(0) {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v158 = F_Int64GetDatum(m, v155)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)) = uint8(v161)
	goto L43
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v158
	goto L43
L48:
	;
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v21)+216))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v21)+224))
	if v172 != 0 {
		goto L58
	} else {
		goto L59
	}
L49:
	;
	v166 = F_Int64GetDatum(m, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+10)) = uint8(v169)
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v166
	goto L48
L53:
	;
	v375 = int32(2)
	v377 = v21 + int32(32) + v365<<(uint(v375)%32)
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+248)))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v378
	v381 = v365 + v375
	if v378 != int32(1) {
		goto L102
	} else {
		goto L103
	}
L54:
	;
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21|v353))) = uint8(v362)
	v365 = v353
	goto L53
L55:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[668]))
	if v327 < int32(0) {
		v353 = v318
		goto L54
	} else {
		goto L96
	}
L56:
	;
	v318 = int32(12)
	v320 = v21 + int32(80)
	goto L55
L57:
	;
	v305 = F_cstring_to_text(m, int32(67303))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L95
	}
L58:
	;
	v281 = v171
	goto L60
L59:
	;
	v174 = m.G0
	v176 = v174 - int32(16)
	m.G0 = v176
	if v171 == int64(0) {
		v254 = int32(0)
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v281 == int64(0) {
		goto L57
	} else {
		goto L88
	}
L61:
	;
	m.G0 = v176 + int32(16)
	switch v254 {
	case 0:
		goto L84
	case 1:
		goto L83
	case 2:
		goto L82
	case 3:
		goto L81
	case 4:
		goto L80
	default:
		v318 = int32(11)
		v320 = v21 + int32(76)
		goto L55
	}
L62:
	;
	v181 = int32(4366320)
	v182 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v182)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+280)) = v183
	*(*int64)(unsafe.Add(mBase, _consts[225])) = v183
	v188 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v188)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v188)+272)) = v189
	*(*int64)(unsafe.Add(mBase, _consts[226])) = v189
	v194 = int64(*(*int32)(unsafe.Add(mBase, _consts[138])))
	v195 = base.I64_div_u_s(v189, v194)
	*(*int64)(unsafe.Add(mBase, uint32(v176)+8)) = v195
	F_KeepLogSeg(m, v189, v176+int32(8))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+440)) = int32(1)
	if v203 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	F_s_lock(m, v207+int32(440), int32(490437), int32(3760), int32(237463))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+440)) = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	v221 = base.I64_extend_i32_s(v220)
	v222 = base.I64_div_u_s(v171, v221)
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v176)+8))
	if base.Ui64(v223) <= base.Ui64(v222) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	v226 = int32(1)
	v228 = base.I64_div_u_s(v189, v221)
	v230 = *(*int32)(unsafe.Add(mBase, _consts[222]))
	v232 = base.I32_div_s(v220, int32(1048576))
	v233 = base.I32_div_s(v230, v232)
	v236 = base.I64_extend_i32_s(v233 + v226)
	if base.Ui64(v228) <= base.Ui64(v236) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v216)+232))
	if base.Ui64(v222) < base.Ui64(v244+int64(1)) {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	v239 = int64(1)
	goto L73
L72:
	;
	v239 = v228 - v236
	goto L73
L73:
	;
	if base.Ui64(v222) < base.Ui64(v239) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v241 = int32(2)
	goto L76
L75:
	;
	v241 = v226
	goto L76
L76:
	;
	v254 = v241
	goto L61
L77:
	;
	v248 = int32(4)
	goto L79
L78:
	;
	v248 = int32(3)
	goto L79
L79:
	;
	v254 = v248
	goto L61
L80:
	;
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v21)+216))
	v281 = v272
	goto L60
L81:
	;
	v269 = F_cstring_to_text(m, int32(433652))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L87
	}
L82:
	;
	v265 = F_cstring_to_text(m, int32(454707))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L86
	}
L83:
	;
	v261 = F_cstring_to_text(m, int32(433722))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+11)) = uint8(v258)
	goto L56
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v261
	goto L56
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v265
	goto L56
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v269
	goto L56
L88:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(1)
	if v284 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_s_lock(m, v86, int32(487009), int32(364), int32(116630))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v86)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+216)) = v293
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v295
	if v292 == v295 {
		goto L57
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v300 = F_cstring_to_text(m, int32(433652))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v300
	goto L56
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v305
	v353 = int32(12)
	goto L54
L96:
	;
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v21)+216))
	v332 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	v333 = base.I64_extend_i32_s(v332)
	v334 = base.I64_div_u_s(v330, v333)
	v336 = base.I32_div_s(v332, int32(1048576))
	v337 = base.I32_div_s(v327, v336)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[669]))
	v340 = base.I32_div_s(v339, v336)
	if base.Ui32(v340) < base.Ui32(v337) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v342 = v337
	goto L99
L98:
	;
	v342 = v340
	goto L99
L99:
	;
	v349 = F_Int64GetDatum(m, (v334+base.I64_extend_i32_s(v342)+int64(1))*v333-v39)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v349
	v365 = v318
	goto L53
L101:
	;
	v401 = v365 + int32(3)
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v21)+384))
	if int64(0) < v402 {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v397 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v381))) = uint8(v397)
	goto L101
L103:
	;
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v21)+240))
	if v384 == int64(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v392 = F_Int64GetDatum(m, v384)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(32)+v381<<(uint(int32(2))%32)))) = v392
	goto L101
L106:
	;
	v417 = v365 + int32(4)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v21)+224))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	if v421 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L107:
	;
	v410 = F_Int64GetDatum(m, v402)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v401))) = uint8(v414)
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(32)+v401<<(uint(int32(2))%32)))) = v410
	goto L106
L111:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+314)))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+24)) = v476
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+313)))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+28)) = base.B2i32(v478 != int32(0))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	F_tuplestore_putvalues(m, v482, v483, v21+int32(32), v21)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L125
	}
L112:
	;
	v456 = int32(366995)
	if base.Ui32(int32(8)) < base.Ui32(v418) {
		v471 = v456
		goto L121
	} else {
		goto L122
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(32)+v417<<(uint(int32(2))%32)))) = int32(1)
	v451 = v365 + int32(5)
	goto L112
L114:
	;
	v437 = v365 + int32(5)
	if v418 != 0 {
		v451 = v437
		goto L112
	} else {
		goto L119
	}
L115:
	;
	v425 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v417))) = uint8(v425)
	goto L114
L116:
	;
	goto L117
L117:
	;
	switch v418 - int32(2) {
	case 0, 2:
		goto L113
	default:
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(32)+v417<<(uint(int32(2))%32)))) = int32(0)
	goto L114
L119:
	;
	v439 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21+v437))) = uint8(v439)
	goto L111
L120:
	;
	v472 = F_cstring_to_text(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L124
	}
L121:
	;
	goto L120
L122:
	;
	if int32(base.Ui32(int32(279))>>(uint(v418)%32))&int32(1) == int32(0) {
		v471 = v456
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v418<<(uint(int32(2))%32))+uint32(_consts[670])))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v471 = v470
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(32)+v451<<(uint(int32(2))%32)))) = v472
	goto L111
L125:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v491 = *(*int32)(unsafe.Add(mBase, _consts[665]))
	v493 = v489
	v495 = v491
	goto L12
L126:
	;
	goto L9
L127:
	;
	m.G0 = v21 + int32(400)
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[30]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v12)
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
				F_errmsg(m, int32(14313), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488926), int32(200), int32(157487))
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
		v34 = int32(*(*uint16)(unsafe.Add(mBase, _consts[169])))
		return base.B2i32(v34 != int32(0))
	}
}
