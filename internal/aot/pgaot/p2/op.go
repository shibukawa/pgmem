package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSetOp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v70 int64
	_ = v70
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v305 int64
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v324 int64
	_ = v324
	var v331 int64
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int64
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v419 int64
	_ = v419
	var v420 int32
	_ = v420
	var v424 int64
	_ = v424
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int64
	_ = v431
	var v438 int64
	_ = v438
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int64
	_ = v466
	var v467 int64
	_ = v467
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v508 int32
	_ = v508
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSetOp[0]))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if int64(0) < v23 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v13 + int32(32)
	return v508
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v23 - int64(1)
	v508 = v15
	goto L6
L8:
	;
	goto L9
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v30 != 0 {
		v508 = int32(0)
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if v31 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v508 = int32(0)
	goto L6
L12:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v345 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	if v41 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v215 = l0 + int32(200)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	goto L77
L18:
	;
	F_ExecReScan(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v45 = m.T0[v44].(func(*base.Module, int32) int32)(m, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v165)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v170 = l0 + int32(200)
	v174 = int32(-1)
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	if v175 == int64(0) {
		v197 = v174
		goto L67
	} else {
		goto L68
	}
L23:
	;
	if v45 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	if v49&int32(2) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v55 = F_LookupTupleHashEntry(m, v37, v45, v13+int32(31), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	if v57 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v60 = v58 - v57
	goto L29
L28:
	;
	v60 = int32(0)
	goto L29
L29:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+31)))
	if v61 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v70
	goto L34
L31:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	v70 = v64 + int64(1)
	goto L30
L32:
	;
	goto L33
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = int64(0)
	v70 = int64(1)
	goto L30
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	F_MemoryContextReset(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	goto L52
L36:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	if v86 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_ExecReScan(m, v40)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v90 = m.T0[v89].(func(*base.Module, int32) int32)(m, v40)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	goto L35
L42:
	;
	if v90 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v94&int32(2) != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v97 = int32(0)
	v101 = F_LookupTupleHashEntry(m, v85, v90, v13+int32(31), v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	if v103 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v106 = v104 - v103
	goto L48
L47:
	;
	v106 = v97
	goto L48
L48:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+31)))
	if v107 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v106)))
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = v110 + int64(1)
	goto L34
L50:
	;
	goto L51
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = int64(1)
	goto L34
L52:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	if v129 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_ExecReScan(m, v38)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v133 = m.T0[v132].(func(*base.Module, int32) int32)(m, v38)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v133 == int32(0) {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)))
	if v137&int32(2) != 0 {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	v140 = int32(0)
	v142 = F_LookupTupleHashEntry(m, v128, v133, v140, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v142 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
	v146 = v144 - v145
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v146)+8)) = v147 + int64(1)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	F_MemoryContextReset(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	goto L52
L66:
	;
	goto L17
L67:
	;
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)) = uint8(v200)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v197
	goto L66
L68:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	v180 = int32(0)
	goto L69
L69:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178+v180*int32(12))+4))
	if v188 != int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v197 = v174
	goto L67
L71:
	;
	v197 = v180
	goto L67
L72:
	;
	goto L73
L73:
	;
	v192 = v180 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v192)) < base.Ui64(v175) {
		v180 = v192
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	m.T0[v339].(func(*base.Module, int32))(m, v216)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L124
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v331
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v336 = F_ExecStoreMinimalTuple(m, v334, v216, int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L123
	}
L77:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v227 != 0 {
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v331 = v324 - int64(1)
	goto L76
L79:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSetOp[0]))
	if v230 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+8)))
	v240 = v237
	goto L85
L83:
	;
	goto L82
L84:
	;
	if v272 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L85:
	;
	if v240&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v272 = v255
	goto L84
L87:
	;
	v272 = int32(0)
	goto L84
L88:
	;
	goto L89
L89:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v251 = v247 & (v248 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v251
	v255 = v246 + v248*int32(12)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v259 = v256 & (v257 ^ v251)
	if v259 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v262 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+8)) = uint8(v262)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v266 != int32(1) {
		v240 = base.B2i32(v259 == int32(0))
		goto L85
	} else {
		goto L93
	}
L93:
	;
	goto L86
L94:
	;
	v275 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v275)
	goto L11
L95:
	;
	goto L96
L96:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v228)+32))
	if v277 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v281 = v278 - v277
	goto L99
L98:
	;
	v281 = int32(0)
	goto L99
L99:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+72))
	switch v283 {
	case 0:
		goto L106
	case 1:
		goto L105
	case 2:
		goto L104
	case 3:
		goto L103
	default:
		goto L102
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v324
	if v324 <= int64(0) {
		goto L77
	} else {
		goto L122
	}
L101:
	;
	v320 = int64(0)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
	if v320 < v321 {
		v331 = v320
		goto L76
	} else {
		goto L121
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L118
	}
L103:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
	if v301 <= v300 {
		goto L115
	} else {
		goto L116
	}
L104:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	if v292 <= int64(0) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
	if v288 < v289 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	if int64(0) < v284 {
		goto L101
	} else {
		goto L107
	}
L107:
	;
	v324 = int64(0)
	goto L100
L108:
	;
	v291 = v288
	goto L110
L109:
	;
	v291 = v289
	goto L110
L110:
	;
	v324 = v291
	goto L100
L111:
	;
	v324 = int64(0)
	goto L100
L112:
	;
	goto L113
L113:
	;
	v296 = int64(0)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v281)+8))
	if v297 != v296 {
		v324 = v296
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v331 = v296
	goto L76
L115:
	;
	v305 = v300 - v301
	goto L117
L116:
	;
	v305 = int64(0)
	goto L117
L117:
	;
	v324 = v305
	goto L100
L118:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v282)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v310
	F_errmsg_internal(m, int32(_a_F_ExecSetOp_0), v13)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_ExecSetOp_1), int32(158), int32(_a_F_ExecSetOp_2))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v324 = v320
	goto L100
L122:
	;
	goto L78
L123:
	;
	v508 = v336
	goto L6
L124:
	;
	goto L11
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	m.T0[v493].(func(*base.Module, int32))(m, v342)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L186
	}
L126:
	;
	v348 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v348)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+52))
	if v350 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	goto L145
L129:
	;
	F_ExecReScan(m, v344)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v354 = m.T0[v353].(func(*base.Module, int32) int32)(m, v344)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v354
	if v354 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v343)+52))
	if v364 != 0 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+4)))
	if v357&int32(2) == int32(0) {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v362)
	goto L11
L138:
	;
	goto L137
L139:
	;
	F_ExecReScan(m, v343)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v368 = m.T0[v367].(func(*base.Module, int32) int32)(m, v343)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v370 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v370)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v370)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v368
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v375 != 0 {
		goto L125
	} else {
		goto L144
	}
L144:
	;
	goto L128
L145:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v391 == int32(1) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L125
L147:
	;
	F_setop_load_group(m, l0+int32(128), v344, l0)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v396 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	if v396 == int64(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L149
L151:
	;
	v399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v399)
	goto L125
L152:
	;
	goto L153
L153:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v401 == int32(1) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	F_setop_load_group(m, l0+int32(152), v343, l0)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v406 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	if v406 != int64(0) {
		goto L162
	} else {
		goto L163
	}
L157:
	;
	goto L156
L158:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v479 != int32(1) {
		goto L145
	} else {
		goto L185
	}
L159:
	;
	v474 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v474)
	goto L158
L160:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+72))
	switch v430 {
	case 0:
		goto L174
	case 1:
		goto L169
	case 2:
		goto L173
	case 3:
		goto L172
	default:
		goto L171
	}
L161:
	;
	if v411 != 0 {
		goto L159
	} else {
		goto L167
	}
L162:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v411 = F_setop_compare_slots(m, v409, v410, l0)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v416 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v416)
	v419 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v427 = int64(0)
	v428 = v419
	goto L160
L165:
	;
	if int32(0) <= v411 {
		goto L161
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v420)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v420)
	v424 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v425 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v427 = v424
	v428 = v425
	goto L160
L168:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v467
	if v467 <= int64(0) {
		goto L158
	} else {
		goto L184
	}
L169:
	;
	if v428 < v427 {
		goto L181
	} else {
		goto L182
	}
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	goto L158
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L178
	}
L172:
	;
	if v428 < v427 {
		goto L170
	} else {
		goto L177
	}
L173:
	;
	v438 = int64(0)
	if base.B2i32(v427 != v438)|base.B2i32(v428 <= v438) != 0 {
		goto L170
	} else {
		goto L176
	}
L174:
	;
	v431 = int64(0)
	if base.B2i32(v428 <= v431)|base.B2i32(v427 <= v431) != 0 {
		goto L170
	} else {
		goto L175
	}
L175:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	v508 = v342
	goto L6
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	v508 = v342
	goto L6
L177:
	;
	v467 = v428 - v427
	goto L168
L178:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v429)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v451
	F_errmsg_internal(m, int32(_a_F_ExecSetOp_0), v13+int32(16))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ExecSetOp_1), int32(158), int32(_a_F_ExecSetOp_2))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	v466 = v428
	goto L183
L182:
	;
	v466 = v427
	goto L183
L183:
	;
	v467 = v466
	goto L168
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v467 - int64(1)
	v508 = v342
	goto L6
L185:
	;
	goto L146
L186:
	;
	goto L11
}
func F_get_op_opfamily_strategy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v7 = F_SearchSysCache3(m, int32(3), base.I64_extend_i32_u(l0), int64(115), base.I64_extend_i32_u(l1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15+v16)+16)))
			F_ReleaseCatCache(m, v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	}
}
func F_get_op_rettype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v5 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)+88))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v16
			}
		}
	}
}
func F_op_error(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if l4 == int32(1) {
			F_errcode(m, int32(84439172))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = F_op_signature_string(m, l1, l2, l3)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v22
					F_errmsg(m, int32(_a_F_op_error_0), v11)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v30 = F_errdetail(m, int32(_a_F_op_error_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_op_error_2), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								F_parser_errposition(m, l0, l6)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_op_error_3), int32(646), int32(_a_F_op_error_4))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
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
		} else {
			F_errcode(m, int32(52461700))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				v46 = F_op_signature_string(m, l1, l2, l3)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v46
					F_errmsg(m, int32(_a_F_op_error_5), v11+int32(16))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v54 = int32(0)
						if l5&int32(8) == v54 {
							if l5&int32(1) != 0 {
								F_parser_errposition(m, l0, l6)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_op_error_3), int32(653), int32(_a_F_op_error_4))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							} else {
								if l5&int32(4) == int32(0) {
									v71 = F_errdetail(m, int32(_a_F_op_error_6), int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_op_error_3), int32(653), int32(_a_F_op_error_4))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v75 = F_errdetail(m, int32(_a_F_op_error_7), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_op_error_3), int32(653), int32(_a_F_op_error_4))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							if base.B2i32(l2 == v54)|base.B2i32(l3 == v54) != 0 {
								v79 = F_errdetail(m, int32(_a_F_op_error_8), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_op_error_9), int32(0))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_op_error_3), int32(653), int32(_a_F_op_error_4))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v87 = F_errdetail(m, int32(_a_F_op_error_10), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_op_error_2), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_op_error_3), int32(653), int32(_a_F_op_error_4))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
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
		}
	}
}
func F_op_hashjoinable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	if l0 <= int32(2987) {
		if l0 == int32(1070) {
			v28 = F_lookup_type_cache(m, l1, int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
				v52 = base.B2i32(v30 == int32(626))
				return v52 & int32(1)
			}
		} else {
			if l0 != int32(2860) {
				v41 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						v52 = int32(0)
						return v52 & int32(1)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47)+78)))
						F_ReleaseCatCache(m, v41)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = v49
							return v52 & int32(1)
						}
					}
				}
			} else {
				v10 = F_lookup_type_cache(m, l1, int32(16))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
					v52 = base.B2i32(v14 == int32(_a_F_op_hashjoinable_0))
					return v52 & int32(1)
				}
			}
		}
	} else {
		if l0 == int32(2988) {
			v34 = F_lookup_type_cache(m, l1, int32(16))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
				v52 = base.B2i32(v36 == int32(_a_F_op_hashjoinable_1))
				return v52 & int32(1)
			}
		} else {
			if l0 != int32(3882) {
				v41 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v41 == int32(0) {
						v52 = int32(0)
						return v52 & int32(1)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47)+78)))
						F_ReleaseCatCache(m, v41)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = v49
							return v52 & int32(1)
						}
					}
				}
			} else {
				v22 = F_lookup_type_cache(m, l1, int32(16))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
					v52 = base.B2i32(v24 == int32(3902))
					return v52 & int32(1)
				}
			}
		}
	}
}
func F_op_input_types(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(_a_F_op_input_types_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_op_input_types_1), int32(1663), int32(_a_F_op_input_types_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+80))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v34
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_op_is_safe_index_member(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v10 = int64(0)
	v12 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(l0), v10, v10)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L11
	}
L2:
	;
	return int32(0)
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if int32(0) < v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v51 = int32(0)
	goto L1
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)+v22<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)+24))
	switch v36 - int32(403) {
	case 0, 2:
		v51 = int32(1)
		goto L1
	default:
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v40 = v22 + int32(1)
	if v40 != v16 {
		v22 = v40
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	return v51
}
func F_op_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = v8 + int32(32)
	F_initStringInfo(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			v16 = F_format_type_be(m, l1)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16
				F_appendStringInfo(m, v11, int32(_a_F_op_signature_string_0), v8+int32(16))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v25 = v8 + int32(32)
					v26 = F_NameListToString(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_appendStringInfoString(m, v25, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = F_format_type_be(m, l2)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30
								F_appendStringInfo(m, v25, int32(_a_F_op_signature_string_1), v8)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
									m.G0 = v8 + int32(48)
									return v36
								}
							}
						}
					}
				}
			}
		} else {
			v25 = v8 + int32(32)
			v26 = F_NameListToString(m, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_appendStringInfoString(m, v25, v26)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = F_format_type_be(m, l2)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30
						F_appendStringInfo(m, v25, int32(_a_F_op_signature_string_1), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
							m.G0 = v8 + int32(48)
							return v36
						}
					}
				}
			}
		}
	}
}
