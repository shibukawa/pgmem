package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
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
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v502 int32
	_ = v502
	v6 = int32(0)
	if l3 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = base.B2i32(int32(0) < l1)
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v16 != int32(37) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(1)
L4:
	;
	return int32(0)
L5:
	;
	if l1 <= int32(0) {
		v459 = l2
		v460 = l3
		v463 = v22
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v502
L7:
	;
	if v463 != 0 {
		v502 = v6
		goto L6
	} else {
		goto L148
	}
L8:
	;
	if l3 <= int32(0) {
		v459 = l2
		v460 = l3
		v463 = v22
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = l0
	v32 = l1
	v33 = l2
	v34 = l3
	goto L10
L10:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	switch v44 - int32(92) {
	case 0:
		goto L22
	case 1, 2:
		goto L19
	case 3:
		goto L20
	default:
		goto L23
	}
L11:
	;
	v459 = v452
	v460 = v450
	v463 = v448
	goto L7
L12:
	;
	v447 = int32(0)
	v448 = base.B2i32(v447 < v442)
	v449 = int32(1)
	v450 = v444 - v449
	v452 = v443 + v449
	if v442 <= v447 {
		v459 = v452
		v460 = v450
		v463 = v448
		goto L7
	} else {
		goto L146
	}
L13:
	;
	v438 = int32(1)
	v442 = v32 - v438
	v443 = v436
	v444 = v437
	v446 = v31 + v438
	goto L12
L14:
	;
	v429 = F_pg_strncoll(m, v194, v299, v31, v32, l4)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L141
	}
L15:
	;
	v424 = F_pg_strncoll(m, v33, v417-v33, v31, v32, l4)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L140
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L136
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L132
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L128
	}
L19:
	;
	if l4 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L20:
	;
	v138 = F_pg_mblen_with_len(m, v31, v32)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L59
	}
L21:
	;
	v63 = v31
	v64 = v32
	v65 = v33
	v66 = v34
	goto L31
L22:
	;
	if v34 <= int32(1) {
		goto L18
	} else {
		goto L26
	}
L23:
	;
	if v44 != int32(37) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v49 = int32(1)
	if base.Ui32(v34) <= base.Ui32(v49) {
		v502 = v49
		goto L6
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v55 = v33 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v56 == v57 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v436 = v55
	v437 = v34 - int32(1)
	goto L13
L28:
	;
	goto L29
L29:
	;
	return int32(0)
L30:
	;
	if v64 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	v76 = int32(1)
	v77 = v66 - v76
	v79 = v65 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	switch v80 - int32(92) {
	case 0:
		goto L33
	case 1, 2:
		v102 = v80
		goto L30
	case 3:
		goto L35
	default:
		goto L36
	}
L32:
	;
	if v77 == int32(1) {
		goto L17
	} else {
		goto L43
	}
L33:
	;
	goto L32
L34:
	;
	if base.Ui32(int32(2)) < base.Ui32(v66) {
		v63 = v93
		v64 = v94
		v65 = v79
		v66 = v77
		goto L31
	} else {
		goto L42
	}
L35:
	;
	if v64 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v80 == int32(37) {
		v93 = v63
		v94 = v64
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v102 = v80
	goto L30
L38:
	;
	return int32(-1)
L39:
	;
	goto L40
L40:
	;
	v89 = F_pg_mblen_with_len(m, v63, v64)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v93 = v63 + v89
	v94 = v64 - v89
	goto L34
L42:
	;
	v502 = int32(1)
	goto L6
L43:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
	v102 = v101
	goto L30
L44:
	;
	return int32(-1)
L45:
	;
	goto L46
L46:
	;
	v109 = v63
	v110 = v64
	goto L47
L47:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v102&int32(255) != v122 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	return int32(-1)
L49:
	;
	v130 = F_pg_mblen_with_len(m, v109, v110)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L57
	}
L50:
	;
	if l4 == int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v127 = F_MB_MatchText(m, v109, v110, v79, v77, l4)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v126 != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	if v127 != 0 {
		v502 = v127
		goto L6
	} else {
		goto L56
	}
L56:
	;
	goto L49
L57:
	;
	v133 = v110 - v130
	if int32(0) < v133 {
		v109 = v109 + v130
		v110 = v133
		goto L47
	} else {
		goto L58
	}
L58:
	;
	goto L48
L59:
	;
	v442 = v32 - v138
	v443 = v33
	v444 = v34
	v446 = v31 + v138
	goto L12
L60:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v44 == v358 {
		v436 = v33
		v437 = v34
		goto L13
	} else {
		goto L127
	}
L61:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v144 != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v34 == int32(0) {
		v417 = v33
		goto L15
	} else {
		goto L63
	}
L63:
	;
	v150 = v34
	v152 = v6
	v154 = v33
	goto L67
L64:
	;
	v315 = v32
	v320 = v31
	goto L108
L65:
	;
	v303 = v33
	v304 = v150
	v308 = v154
	v309 = v6
	v310 = v154 - v33
	goto L64
L66:
	;
	v193 = v191 - v33
	v194 = F_palloc(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L78
	}
L67:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	switch v160 - int32(92) {
	case 0:
		goto L70
	case 1, 2:
		v176 = v150
		v177 = v152
		v178 = v154
		goto L69
	case 3:
		goto L71
	default:
		goto L72
	}
L68:
	;
	v183 = int32(1)
	if v177&v183 == int32(0) {
		v417 = v180
		goto L15
	} else {
		goto L77
	}
L69:
	;
	v179 = int32(1)
	v180 = v178 + v179
	v182 = v176 - v179
	if v182 != 0 {
		v150 = v182
		v152 = v177
		v154 = v180
		goto L67
	} else {
		goto L76
	}
L70:
	;
	v170 = v150 - int32(1)
	if v170 == int32(0) {
		goto L16
	} else {
		goto L75
	}
L71:
	;
	if v152&int32(1) == int32(0) {
		goto L65
	} else {
		goto L74
	}
L72:
	;
	if v160 != int32(37) {
		v176 = v150
		v177 = v152
		v178 = v154
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v189 = v150
	v191 = v154
	v192 = v6
	goto L66
L75:
	;
	v173 = int32(1)
	v176 = v170
	v177 = v173
	v178 = v154 + v173
	goto L69
L76:
	;
	goto L68
L77:
	;
	v189 = int32(0)
	v191 = v180
	v192 = v183
	goto L66
L78:
	;
	if base.Ui32(v191) <= base.Ui32(v33) {
		v291 = v194
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v299 = v291 - v194
	if v192 != 0 {
		goto L14
	} else {
		goto L107
	}
L80:
	;
	v198 = v193 & int32(3)
	if v198 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v33-v191) {
		v291 = v231
		goto L79
	} else {
		goto L91
	}
L82:
	;
	v231 = v194
	v232 = v33
	goto L81
L83:
	;
	goto L84
L84:
	;
	v206 = v194
	v207 = v33
	v210 = v6
	goto L85
L85:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v214 != int32(92) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v231 = v220
	v232 = v222
	goto L81
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v214)
	v220 = v206 + int32(1)
	goto L89
L88:
	;
	v220 = v206
	goto L89
L89:
	;
	v221 = int32(1)
	v222 = v207 + v221
	v224 = v210 + v221
	if v224 != v198 {
		v206 = v220
		v207 = v222
		v210 = v224
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v247 = v231
	v248 = v232
	goto L92
L92:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v255 != int32(92) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v291 = v282
	goto L79
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v247))) = uint8(v255)
	v261 = v247 + int32(1)
	goto L96
L95:
	;
	v261 = v247
	goto L96
L96:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	if v262 != int32(92) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v262)
	v268 = v261 + int32(1)
	goto L99
L98:
	;
	v268 = v261
	goto L99
L99:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+2)))
	if v269 != int32(92) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v268))) = uint8(v269)
	v275 = v268 + int32(1)
	goto L102
L101:
	;
	v275 = v268
	goto L102
L102:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+3)))
	if v276 != int32(92) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v275))) = uint8(v276)
	v282 = v275 + int32(1)
	goto L105
L104:
	;
	v282 = v275
	goto L105
L105:
	;
	v284 = v248 + int32(4)
	if v284 != v191 {
		v247 = v282
		v248 = v284
		goto L92
	} else {
		goto L106
	}
L106:
	;
	goto L93
L107:
	;
	v303 = v194
	v304 = v189
	v308 = v191
	v309 = v194
	v310 = v299
	goto L64
L108:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v328 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v332 = F_pg_strncoll(m, v303, v310, v31, v320-v31, l4)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	if v315 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	if v332 != 0 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v334 = F_MB_MatchText(m, v320, v315, v308, v304, l4)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	if v334 != int32(1) {
		goto L114
	} else {
		goto L118
	}
L118:
	;
	if v309 == int32(0) {
		v502 = int32(1)
		goto L6
	} else {
		goto L119
	}
L119:
	;
	F_pfree(m, v309)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	return int32(1)
L121:
	;
	v347 = int32(0)
	if v309 == v347 {
		v502 = v347
		goto L6
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v354 = F_pg_mblen_with_len(m, v320, v315)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L126
	}
L124:
	;
	F_pfree(m, v309)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	return int32(0)
L126:
	;
	v315 = v315 - v354
	v320 = v354 + v320
	goto L108
L127:
	;
	return int32(0)
L128:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(217743), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(499445), int32(107), int32(64336))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(217743), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(499445), int32(169), int32(64336))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(217743), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(499445), int32(237), int32(64336))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	return base.B2i32(v424 == int32(0))
L141:
	;
	if v194 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_pfree(m, v194)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	return base.B2i32(v429 == int32(0))
L145:
	;
	goto L144
L146:
	;
	if int32(1) < v444 {
		v31 = v446
		v32 = v442
		v33 = v452
		v34 = v450
		goto L10
	} else {
		goto L147
	}
L147:
	;
	goto L11
L148:
	;
	v470 = int32(1)
	if v460 <= int32(0) {
		v502 = v470
		goto L6
	} else {
		goto L149
	}
L149:
	;
	v475 = v459
	v476 = v460
	goto L150
L150:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	if v486 != int32(37) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v502 = v470
	goto L6
L152:
	;
	return int32(-1)
L153:
	;
	goto L154
L154:
	;
	v491 = int32(1)
	if v491 < v476 {
		v475 = v475 + v491
		v476 = v476 - v491
		goto L150
	} else {
		goto L155
	}
L155:
	;
	goto L151
}
func F_MJFillInner(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 float64
	_ = v68
	var v74 int32
	_ = v74
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_MemoryContextReset(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v20
		if v11 != 0 {
			v22 = int32(4520560)
			v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) int32)(m, v11, v12, v9+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
				if v30 == int32(0) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v65 == int32(0) {
						v74 = v2
					} else {
						v68 = *(*float64)(unsafe.Add(mBase, uint32(v65)+248))
						*(*float64)(unsafe.Add(mBase, uint32(v65)+248)) = base.F64_add(v68, float64(1))
						v74 = v2
					}
					m.G0 = v9 + int32(16)
					return v74
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
					m.T0[v42].(func(*base.Module, int32))(m, v40)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = int32(4520560)
						v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v48
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
						v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v38+int32(4), v39, int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
							v60 = v58 & int32(65533)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v60)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)) = uint16(v63)
							v74 = v40
							m.G0 = v9 + int32(16)
							return v74
						}
					}
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
			m.T0[v42].(func(*base.Module, int32))(m, v40)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = int32(4520560)
				v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v48
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
				v54 = m.T0[v53].(func(*base.Module, int32, int32, int32) int32)(m, v38+int32(4), v39, int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
					v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
					v60 = v58 & int32(65533)
					*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v60)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
					*(*uint16)(unsafe.Add(mBase, uint32(v40)+6)) = uint16(v63)
					v74 = v40
					m.G0 = v9 + int32(16)
					return v74
				}
			}
		}
	}
}
func F_MakeSingleTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_MakeTupleTableSlot(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F___math_oflow(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v2 = float64(3.105036184601418e+231)
	if l0 != 0 {
		v4 = base.F64_neg(v2)
	} else {
		v4 = v2
	}
	v5 = F_fp_barrier_1(m, v4)
	return base.F64_mul(v2, v5)
}
func F___multi3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v31 int64
	_ = v31
	v10 = int64(32)
	v11 = int64(base.Ui64(l3) >> (uint(v10) % 64))
	v13 = int64(base.Ui64(l1) >> (uint(v10) % 64))
	v16 = int64(4294967295)
	v17 = l3 & v16
	v19 = l1 & v16
	v20 = v17 * v19
	v24 = int64(base.Ui64(v20)>>(uint(v10)%64)) + v17*v13
	v31 = v19*v11 + v24&v16
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l1*l4 + l2*l3 + v11*v13 + int64(base.Ui64(v24)>>(uint(v10)%64)) + int64(base.Ui64(v31)>>(uint(v10)%64))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v20&v16 | v31<<(uint(v10)%64)
	return
}
func F_makeAConst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v4 - int32(465) {
	case 0:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = F_palloc0(m, int32(20))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v16
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(1997159792712)
			v24 = v18
			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
			v27 = v24
			return v27
		}
	case 1:
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = F_palloc0(m, int32(20))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v7
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(2001454760008)
			v24 = v9
			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
			v27 = v24
			return v27
		}
	default:
		v27 = int32(0)
		return v27
	}
}
func F_makeFloat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(466)
		return v4
	}
}
func F_makeNullConst(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_get_typlenbyval(m, l0, v10+int32(14), v10+int32(13))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+14)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
		v23 = F_palloc0(m, int32(32))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(-1)
			*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)) = uint8(v21)
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v28)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(7)
			m.G0 = v10 + int32(16)
			return v23
		}
	}
}
func F_make_SAOP_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = F_get_array_type(m, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if v17 == int32(0) {
			v186 = int32(0)
			m.G0 = v15 + int32(32)
			return v186
		} else {
			if l6 != 0 {
				v25 = F_palloc0(m, int32(36))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(35)
					v30 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)) = uint8(v30)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l5
					v149 = v25
					v154 = F_palloc0(m, int32(36))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(20)
						v159 = F_get_opcode(m, l0)
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = l4
							v162 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v154)+20)) = uint8(v162)
							*(*int64)(unsafe.Add(mBase, uint32(v154)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v159
							*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v149
							*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v149
							v175 = F_list_make2_impl(m, v15+int32(8), v15+int32(4))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v175
								v186 = v154
								m.G0 = v15 + int32(32)
								return v186
							}
						}
					}
				}
			} else {
				if l5 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1)
					F_get_typlenbyvalalign(m, l2, v15+int32(30), v15+int32(29), v15+int32(28))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v51 = F_palloc(m, int32(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v54 = F_palloc(m, int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v116 = v51
								v117 = v54
								v119 = int32(-1)
								v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+30)))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)))
								v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+28)))
								v129 = F_construct_md_array(m, v116, v117, int32(1), v15+int32(24), v15+int32(20), l2, v126, v127, v128)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									v131 = int32(0)
									v133 = F_makeConst(m, v17, v119, l3, v119, v129, v131, v131)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v116)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v117)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												F_list_free(m, l5)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													v149 = v133
													v154 = F_palloc0(m, int32(36))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = l0
														*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(20)
														v159 = F_get_opcode(m, l0)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = l4
															v162 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v154)+20)) = uint8(v162)
															*(*int64)(unsafe.Add(mBase, uint32(v154)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v159
															*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v149
															*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v149
															v175 = F_list_make2_impl(m, v15+int32(8), v15+int32(4))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v175
																v186 = v154
																m.G0 = v15 + int32(32)
																return v186
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
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v56
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1)
					F_get_typlenbyvalalign(m, l2, v15+int32(30), v15+int32(29), v15+int32(28))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
						v71 = F_palloc(m, v68<<(uint(int32(2))%32))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
							v74 = F_palloc(m, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
								if v76 <= int32(0) {
									v116 = v71
									v117 = v74
								} else {
									v86 = int32(0)
									for {
										v93 = v86 << (uint(int32(2)) % 32)
										v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v95+v93)))
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v71+v93))) = v98
										v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)))
										*(*uint8)(unsafe.Add(mBase, uint32(v86+v74))) = uint8(v101)
										v104 = v86 + int32(1)
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
										if v104 < v105 {
											v86 = v104
											continue
										} else {
											break
										}
										break
									}
									v116 = v71
									v117 = v74
								}
								v119 = int32(-1)
								v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+30)))
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+29)))
								v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+28)))
								v129 = F_construct_md_array(m, v116, v117, int32(1), v15+int32(24), v15+int32(20), l2, v126, v127, v128)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									v131 = int32(0)
									v133 = F_makeConst(m, v17, v119, l3, v119, v129, v131, v131)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v116)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v117)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												F_list_free(m, l5)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													v149 = v133
													v154 = F_palloc0(m, int32(36))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = l0
														*(*int32)(unsafe.Add(mBase, uint32(v154))) = int32(20)
														v159 = F_get_opcode(m, l0)
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v154)+24)) = l4
															v162 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v154)+20)) = uint8(v162)
															*(*int64)(unsafe.Add(mBase, uint32(v154)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v159
															*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v149
															*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v149
															v175 = F_list_make2_impl(m, v15+int32(8), v15+int32(4))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = int32(-1)
																*(*int32)(unsafe.Add(mBase, uint32(v154)+28)) = v175
																v186 = v154
																m.G0 = v15 + int32(32)
																return v186
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
		}
	}
}
func F_make_ands_explicit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	if l0 == int32(0) {
		v6 = F_palloc0(m, int32(32))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(-1)
			v12 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+24)) = uint16(v12)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = int64(4294967297)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(4294967295)
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(68719476743)
			return v6
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v21 == int32(1) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			return v25
		} else {
			v28 = F_palloc0(m, int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v28))) = int64(21)
				return v28
			}
		}
	}
}
func F_make_plain_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v194 int64
	_ = v194
	v4 = l3
	v5 = l4
	v6 = l5
	v7 = l6
	v14 = F_palloc0(m, int32(168))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(318)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)) = uint8(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l7
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+9)) = uint8(v29)
	if l7 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = F_contain_leaked_vars(m, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v36 = v29
	goto L5
L5:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)) = uint8(v36)
	if l1 == v37 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v36 = v32 ^ int32(1)
	goto L5
L7:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if l8 != 0 {
		goto L37
	} else {
		goto L38
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = int64(0)
	v129 = F_pull_varnos(m, l0, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L36
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v42 != int32(17) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v45 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 != int32(2) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = F_pull_varnos(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v53
	v56 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v57 == v56 {
		v65 = v56
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = F_pull_varnos(m, l0, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 < int32(2) {
		v65 = v56
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v65 = v64
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v70 = F_bms_union(m, v69, v66)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v73 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v76 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v79 = int32(0)
	if v73 == v79 {
		v120 = v79
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v120 != 0 {
		goto L7
	} else {
		goto L35
	}
L22:
	;
	goto L21
L23:
	;
	if v76 == int32(0) {
		v120 = v79
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v88 < v89 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = v88
	goto L27
L26:
	;
	v91 = v89
	goto L27
L27:
	;
	if v91 <= int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v94 = int32(1)
	goto L30
L29:
	;
	v94 = v91
	goto L30
L30:
	;
	v95 = int32(8)
	v100 = int32(0)
	goto L31
L31:
	;
	v107 = v100 << (uint(int32(2)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v76+v95+v107)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v73+v95))))
	v112 = v109 & v111
	v114 = base.B2i32(v112 != int32(0))
	if v112 != 0 {
		v120 = v114
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v120 = v114
	goto L22
L33:
	;
	v116 = v100 + int32(1)
	if v116 != v94 {
		v100 = v116
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+9)) = uint8(v124)
	goto L7
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v129
	goto L7
L37:
	;
	v135 = l8
	goto L39
L38:
	;
	v135 = v134
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v138 = F_bms_difference(m, v134, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v140 = int32(0)
	if v138 == v140 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v175
	F_bms_free(m, v138)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L54
	}
L42:
	;
	v175 = int32(0)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v147 = int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v148 <= v147 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v151 = v147
	goto L47
L46:
	;
	v151 = v148
	goto L47
L47:
	;
	v155 = int32(0)
	v157 = v140
	goto L48
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v138+int32(8)+v155<<(uint(int32(2))%32))))
	if v163 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v175 = v166
	goto L41
L50:
	;
	v166 = v157 + base.I32_popcnt(v163)
	goto L52
L51:
	;
	v166 = v157
	goto L52
L52:
	;
	v168 = v155 + int32(1)
	if v168 != v151 {
		v155 = v168
		v157 = v166
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v181 = v179 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v181
	v183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v183
	v185 = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v181
	v194 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v14)+104)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v14)+112)) = v194
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+120)) = uint8(v183)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+160)) = v194
	*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v14)+144)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = v185
	return v14
}
func F_manifest_process_file(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v303 int64
	_ = v303
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v324 int64
	_ = v324
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v349 int64
	_ = v349
	var v364 float64
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int64
	_ = v987
	var v989 int64
	_ = v989
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1137 int32
	_ = v1137
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1434 int64
	_ = v1434
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int64
	_ = v1487
	var v1489 int64
	_ = v1489
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1518 int64
	_ = v1518
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1548 int32
	_ = v1548
	var v1557 int32
	_ = v1557
	var v1568 int32
	_ = v1568
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1611 int32
	_ = v1611
	var v1618 int32
	_ = v1618
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v19 = F_strlen(m, l1)
	mBase = m.M
	v25 = v19 - int32(1636608432)
	if l1&int32(3) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v286 = v284
	v295 = v285
	goto L41
L2:
	;
	v257 = int32(14)
	v259 = v253 ^ v254 - base.I32_rotl(v253, v257)
	v263 = v259 ^ v252 - base.I32_rotl(v259, int32(11))
	v267 = v263 ^ v253 - base.I32_rotl(v263, int32(25))
	v271 = v267 ^ v259 - base.I32_rotl(v267, int32(16))
	v275 = v271 ^ v263 - base.I32_rotl(v271, int32(4))
	v279 = v275 ^ v267 - base.I32_rotl(v275, v257)
	goto L1
L3:
	;
	switch v183 - int32(1) {
	case 0:
		v245 = v184
		v246 = v185
		v247 = v186
		goto L30
	case 1:
		v238 = v184
		v239 = v185
		v240 = v186
		goto L31
	case 2:
		v231 = v184
		v232 = v185
		v233 = v186
		goto L32
	case 3:
		v225 = v185
		v226 = v186
		goto L33
	case 4:
		v221 = v185
		v222 = v186
		goto L34
	case 5:
		v215 = v185
		v216 = v186
		goto L35
	case 6:
		v209 = v185
		v210 = v186
		goto L36
	case 7:
		v204 = v186
		goto L37
	case 8:
		v199 = v186
		goto L38
	case 9:
		v194 = v186
		goto L39
	case 10:
		goto L40
	default:
		v252 = v184
		v253 = v185
		v254 = v186
		goto L2
	}
L4:
	;
	v134 = l1
	v135 = v19
	v136 = v25
	v137 = v25
	v138 = v25
	goto L27
L5:
	;
	if base.Ui32(int32(11)) < base.Ui32(v19) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(v19) < base.Ui32(int32(12)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v182 = l1
	v183 = v19
	v184 = v25
	v185 = v25
	v186 = v25
	goto L3
L9:
	;
	switch v81 - int32(1) {
	case 0:
		v131 = v82
		goto L16
	case 1:
		v126 = v82
		goto L17
	case 2:
		goto L18
	case 3:
		v119 = v83
		goto L19
	case 4:
		v116 = v83
		goto L20
	case 5:
		v111 = v83
		goto L21
	case 6:
		goto L22
	case 7:
		v102 = v84
		goto L23
	case 8:
		v97 = v84
		goto L24
	case 9:
		v92 = v84
		goto L25
	case 10:
		goto L26
	default:
		v252 = v82
		v253 = v83
		v254 = v84
		goto L2
	}
L10:
	;
	v80 = l1
	v81 = v19
	v82 = v25
	v83 = v25
	v84 = v25
	goto L9
L11:
	;
	goto L12
L12:
	;
	v32 = l1
	v33 = v19
	v34 = v25
	v35 = v25
	v36 = v25
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v39 = v38 + v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v43 = v42 + v36
	v45 = int32(4)
	v47 = v40 + v34 - v43 ^ base.I32_rotl(v43, v45)
	v51 = v39 - v47 ^ base.I32_rotl(v47, int32(6))
	v52 = v43 + v39
	v53 = v47 + v52
	v54 = v51 + v53
	v58 = v52 - v51 ^ base.I32_rotl(v51, int32(8))
	v62 = v53 - v58 ^ base.I32_rotl(v58, int32(16))
	v66 = v54 - v62 ^ base.I32_rotl(v62, int32(19))
	v67 = v58 + v54
	v68 = v62 + v67
	v69 = v66 + v68
	v73 = v67 - v66 ^ base.I32_rotl(v66, v45)
	v74 = int32(12)
	v75 = v32 + v74
	v77 = v33 - v74
	if base.Ui32(int32(11)) < base.Ui32(v77) {
		v32 = v75
		v33 = v77
		v34 = v68
		v35 = v69
		v36 = v73
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v80 = v75
	v81 = v77
	v82 = v68
	v83 = v69
	v84 = v73
	goto L9
L15:
	;
	goto L14
L16:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v252 = v131 + v132
	v253 = v83
	v254 = v84
	goto L2
L17:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	v131 = v127<<(uint(int32(8))%32) + v126
	goto L16
L18:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+2)))
	v126 = v122<<(uint(int32(16))%32) + v82
	goto L17
L19:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v252 = v120 + v82
	v253 = v119
	v254 = v84
	goto L2
L20:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
	v119 = v116 + v117
	goto L19
L21:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+5)))
	v116 = v112<<(uint(int32(8))%32) + v111
	goto L20
L22:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+6)))
	v111 = v107<<(uint(int32(16))%32) + v83
	goto L21
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v252 = v103 + v82
	v253 = v105 + v83
	v254 = v102
	goto L2
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
	v102 = v98<<(uint(int32(8))%32) + v97
	goto L23
L25:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
	v97 = v93<<(uint(int32(16))%32) + v92
	goto L24
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
	v92 = v88<<(uint(int32(24))%32) + v84
	goto L25
L27:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v141 = v140 + v137
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v145 = v144 + v138
	v147 = int32(4)
	v149 = v142 + v136 - v145 ^ base.I32_rotl(v145, v147)
	v153 = v141 - v149 ^ base.I32_rotl(v149, int32(6))
	v154 = v145 + v141
	v155 = v149 + v154
	v156 = v153 + v155
	v160 = v154 - v153 ^ base.I32_rotl(v153, int32(8))
	v164 = v155 - v160 ^ base.I32_rotl(v160, int32(16))
	v168 = v156 - v164 ^ base.I32_rotl(v164, int32(19))
	v169 = v160 + v156
	v170 = v164 + v169
	v171 = v168 + v170
	v175 = v169 - v168 ^ base.I32_rotl(v168, v147)
	v176 = int32(12)
	v177 = v134 + v176
	v179 = v135 - v176
	if base.Ui32(int32(11)) < base.Ui32(v179) {
		v134 = v177
		v135 = v179
		v136 = v170
		v137 = v171
		v138 = v175
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v182 = v177
	v183 = v179
	v184 = v170
	v185 = v171
	v186 = v175
	goto L3
L29:
	;
	goto L28
L30:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v252 = v245 + v248
	v253 = v246
	v254 = v247
	goto L2
L31:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	v245 = v241<<(uint(int32(8))%32) + v238
	v246 = v239
	v247 = v240
	goto L30
L32:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
	v238 = v234<<(uint(int32(16))%32) + v231
	v239 = v232
	v240 = v233
	goto L31
L33:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
	v231 = v227<<(uint(int32(24))%32) + v184
	v232 = v225
	v233 = v226
	goto L32
L34:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
	v225 = v221 + v223
	v226 = v222
	goto L33
L35:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
	v221 = v217<<(uint(int32(8))%32) + v215
	v222 = v216
	goto L34
L36:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
	v215 = v211<<(uint(int32(16))%32) + v209
	v216 = v210
	goto L35
L37:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
	v209 = v205<<(uint(int32(24))%32) + v185
	v210 = v204
	goto L36
L38:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
	v204 = v200<<(uint(int32(8))%32) + v199
	goto L37
L39:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
	v199 = v195<<(uint(int32(16))%32) + v194
	goto L38
L40:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+10)))
	v194 = v190<<(uint(int32(24))%32) + v186
	goto L39
L41:
	;
	if base.Ui32(v286) <= base.Ui32(v295) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v1618 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1618
	v286 = v1618
	v295 = v1611
	goto L41
L44:
	;
	return
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1568))) = int32(1)
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1580)+24))
	v1582 = F_MemoryContextStrdup(m, v1581, l1)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L63
	} else {
		goto L266
	}
L46:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1557 + int32(1)
	v1568 = v1548
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L63
	} else {
		goto L263
	}
L48:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if v303 == int64(4294967296) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v1077 = int32(0)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1080 = v1079 & (v279 ^ v271 - base.I32_rotl(v279, int32(24)))
	v1083 = v1078 + v1080<<(uint(int32(4))%32)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	if v1084 == v1077 {
		v1548 = v1083
		goto L46
	} else {
		goto L184
	}
L51:
	;
	v306 = int32(0)
	v308 = int64(2)
	v310 = v303 << (uint(int64(1)) % 64)
	if base.Ui64(v310) <= base.Ui64(v308) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L50
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L63
	} else {
		goto L181
	}
L54:
	;
	v313 = v308
	goto L56
L55:
	;
	v313 = v310
	goto L56
L56:
	;
	v314 = int64(1)
	if v313&(v313-v314) == int64(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v324 = v313
	goto L59
L58:
	;
	v324 = v314 << (uint(int64(64)-base.I64_clz(v313)) % 64)
	goto L59
L59:
	;
	if base.Ui64(v324<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v336 = F_MemoryContextAllocExtended(m, v331, base.I32_wrap_i64(v324)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L63
	} else {
		goto L178
	}
L63:
	;
	return
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v336
	v339 = int64(1)
	if v324&(v324-v339) == int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v349 = v324
	goto L67
L66:
	;
	v349 = v339 << (uint(int64(64)-base.I64_clz(v324)) % 64)
	goto L67
L67:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v349<<(uint(int64(4))%64)) {
		goto L53
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = base.I32_wrap_i64(v349) - int32(1)
	v364 = base.F64_mul(base.F64_convert_i64_u(v349), float64(0.9))
	if base.F64_lt(v364, float64(4.294967296e+09))&base.F64_ge(v364, float64(0)) != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v349 == int64(4294967296) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v370 = base.I32_trunc_f64_u(v364)
	v372 = v370
	goto L69
L71:
	;
	goto L72
L72:
	;
	v372 = int32(0)
	goto L69
L73:
	;
	v373 = int32(-85899346)
	goto L75
L74:
	;
	v373 = v372
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v373
	if v330 != int64(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v380 = v306
	goto L80
L77:
	;
	goto L78
L78:
	;
	F_pfree(m, v329)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L63
	} else {
		goto L177
	}
L79:
	;
	v678 = v674
	v682 = v306
	goto L125
L80:
	;
	v395 = v329 + v380<<(uint(int32(4))%32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	if v396 != int32(1) {
		v674 = v380
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v674 = int32(0)
	goto L79
L82:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v400 = F_strlen(m, v399)
	mBase = m.M
	v406 = v400 - int32(1636608432)
	if v399&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if (v660^v652-base.I32_rotl(v660, int32(24)))&v665 == v380 {
		v674 = v380
		goto L79
	} else {
		goto L123
	}
L84:
	;
	v638 = int32(14)
	v640 = v634 ^ v635 - base.I32_rotl(v634, v638)
	v644 = v640 ^ v633 - base.I32_rotl(v640, int32(11))
	v648 = v644 ^ v634 - base.I32_rotl(v644, int32(25))
	v652 = v648 ^ v640 - base.I32_rotl(v648, int32(16))
	v656 = v652 ^ v644 - base.I32_rotl(v652, int32(4))
	v660 = v656 ^ v648 - base.I32_rotl(v656, v638)
	goto L83
L85:
	;
	switch v564 - int32(1) {
	case 0:
		v626 = v565
		v627 = v566
		v628 = v567
		goto L112
	case 1:
		v619 = v565
		v620 = v566
		v621 = v567
		goto L113
	case 2:
		v612 = v565
		v613 = v566
		v614 = v567
		goto L114
	case 3:
		v606 = v566
		v607 = v567
		goto L115
	case 4:
		v602 = v566
		v603 = v567
		goto L116
	case 5:
		v596 = v566
		v597 = v567
		goto L117
	case 6:
		v590 = v566
		v591 = v567
		goto L118
	case 7:
		v585 = v567
		goto L119
	case 8:
		v580 = v567
		goto L120
	case 9:
		v575 = v567
		goto L121
	case 10:
		goto L122
	default:
		v633 = v565
		v634 = v566
		v635 = v567
		goto L84
	}
L86:
	;
	v515 = v399
	v516 = v400
	v517 = v406
	v518 = v406
	v519 = v406
	goto L109
L87:
	;
	if base.Ui32(int32(11)) < base.Ui32(v400) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(v400) < base.Ui32(int32(12)) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v563 = v399
	v564 = v400
	v565 = v406
	v566 = v406
	v567 = v406
	goto L85
L91:
	;
	switch v462 - int32(1) {
	case 0:
		v512 = v463
		goto L98
	case 1:
		v507 = v463
		goto L99
	case 2:
		goto L100
	case 3:
		v500 = v464
		goto L101
	case 4:
		v497 = v464
		goto L102
	case 5:
		v492 = v464
		goto L103
	case 6:
		goto L104
	case 7:
		v483 = v465
		goto L105
	case 8:
		v478 = v465
		goto L106
	case 9:
		v473 = v465
		goto L107
	case 10:
		goto L108
	default:
		v633 = v463
		v634 = v464
		v635 = v465
		goto L84
	}
L92:
	;
	v461 = v399
	v462 = v400
	v463 = v406
	v464 = v406
	v465 = v406
	goto L91
L93:
	;
	goto L94
L94:
	;
	v413 = v399
	v414 = v400
	v415 = v406
	v416 = v406
	v417 = v406
	goto L95
L95:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	v420 = v419 + v416
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v413)+8))
	v424 = v423 + v417
	v426 = int32(4)
	v428 = v421 + v415 - v424 ^ base.I32_rotl(v424, v426)
	v432 = v420 - v428 ^ base.I32_rotl(v428, int32(6))
	v433 = v424 + v420
	v434 = v428 + v433
	v435 = v432 + v434
	v439 = v433 - v432 ^ base.I32_rotl(v432, int32(8))
	v443 = v434 - v439 ^ base.I32_rotl(v439, int32(16))
	v447 = v435 - v443 ^ base.I32_rotl(v443, int32(19))
	v448 = v439 + v435
	v449 = v443 + v448
	v450 = v447 + v449
	v454 = v448 - v447 ^ base.I32_rotl(v447, v426)
	v455 = int32(12)
	v456 = v413 + v455
	v458 = v414 - v455
	if base.Ui32(int32(11)) < base.Ui32(v458) {
		v413 = v456
		v414 = v458
		v415 = v449
		v416 = v450
		v417 = v454
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v461 = v456
	v462 = v458
	v463 = v449
	v464 = v450
	v465 = v454
	goto L91
L97:
	;
	goto L96
L98:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	v633 = v512 + v513
	v634 = v464
	v635 = v465
	goto L84
L99:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+1)))
	v512 = v508<<(uint(int32(8))%32) + v507
	goto L98
L100:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+2)))
	v507 = v503<<(uint(int32(16))%32) + v463
	goto L99
L101:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v633 = v501 + v463
	v634 = v500
	v635 = v465
	goto L84
L102:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+4)))
	v500 = v497 + v498
	goto L101
L103:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+5)))
	v497 = v493<<(uint(int32(8))%32) + v492
	goto L102
L104:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+6)))
	v492 = v488<<(uint(int32(16))%32) + v464
	goto L103
L105:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	v633 = v484 + v463
	v634 = v486 + v464
	v635 = v483
	goto L84
L106:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+8)))
	v483 = v479<<(uint(int32(8))%32) + v478
	goto L105
L107:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+9)))
	v478 = v474<<(uint(int32(16))%32) + v473
	goto L106
L108:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+10)))
	v473 = v469<<(uint(int32(24))%32) + v465
	goto L107
L109:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	v522 = v521 + v518
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	v526 = v525 + v519
	v528 = int32(4)
	v530 = v523 + v517 - v526 ^ base.I32_rotl(v526, v528)
	v534 = v522 - v530 ^ base.I32_rotl(v530, int32(6))
	v535 = v526 + v522
	v536 = v530 + v535
	v537 = v534 + v536
	v541 = v535 - v534 ^ base.I32_rotl(v534, int32(8))
	v545 = v536 - v541 ^ base.I32_rotl(v541, int32(16))
	v549 = v537 - v545 ^ base.I32_rotl(v545, int32(19))
	v550 = v541 + v537
	v551 = v545 + v550
	v552 = v549 + v551
	v556 = v550 - v549 ^ base.I32_rotl(v549, v528)
	v557 = int32(12)
	v558 = v515 + v557
	v560 = v516 - v557
	if base.Ui32(int32(11)) < base.Ui32(v560) {
		v515 = v558
		v516 = v560
		v517 = v551
		v518 = v552
		v519 = v556
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v563 = v558
	v564 = v560
	v565 = v551
	v566 = v552
	v567 = v556
	goto L85
L111:
	;
	goto L110
L112:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	v633 = v626 + v629
	v634 = v627
	v635 = v628
	goto L84
L113:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+1)))
	v626 = v622<<(uint(int32(8))%32) + v619
	v627 = v620
	v628 = v621
	goto L112
L114:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+2)))
	v619 = v615<<(uint(int32(16))%32) + v612
	v620 = v613
	v621 = v614
	goto L113
L115:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+3)))
	v612 = v608<<(uint(int32(24))%32) + v565
	v613 = v606
	v614 = v607
	goto L114
L116:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+4)))
	v606 = v602 + v604
	v607 = v603
	goto L115
L117:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+5)))
	v602 = v598<<(uint(int32(8))%32) + v596
	v603 = v597
	goto L116
L118:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+6)))
	v596 = v592<<(uint(int32(16))%32) + v590
	v597 = v591
	goto L117
L119:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+7)))
	v590 = v586<<(uint(int32(24))%32) + v566
	v591 = v585
	goto L118
L120:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+8)))
	v585 = v581<<(uint(int32(8))%32) + v580
	goto L119
L121:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+9)))
	v580 = v576<<(uint(int32(16))%32) + v575
	goto L120
L122:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+10)))
	v575 = v571<<(uint(int32(24))%32) + v567
	goto L121
L123:
	;
	v669 = v380 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v669)) < base.Ui64(v330) {
		v380 = v669
		goto L80
	} else {
		goto L124
	}
L124:
	;
	goto L81
L125:
	;
	v693 = v329 + v678<<(uint(int32(4))%32)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	if v694 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L78
L127:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	v698 = F_strlen(m, v697)
	mBase = m.M
	v704 = v698 - int32(1636608432)
	if v697&int32(3) != 0 {
		goto L134
	} else {
		goto L135
	}
L128:
	;
	goto L129
L129:
	;
	v1008 = v678 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1008)) < base.Ui64(v330) {
		goto L173
	} else {
		goto L174
	}
L130:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v964 = v958 ^ v950 - base.I32_rotl(v958, int32(24))
	goto L170
L131:
	;
	v936 = int32(14)
	v938 = v932 ^ v933 - base.I32_rotl(v932, v936)
	v942 = v938 ^ v931 - base.I32_rotl(v938, int32(11))
	v946 = v942 ^ v932 - base.I32_rotl(v942, int32(25))
	v950 = v946 ^ v938 - base.I32_rotl(v946, int32(16))
	v954 = v950 ^ v942 - base.I32_rotl(v950, int32(4))
	v958 = v954 ^ v946 - base.I32_rotl(v954, v936)
	goto L130
L132:
	;
	switch v862 - int32(1) {
	case 0:
		v924 = v863
		v925 = v864
		v926 = v865
		goto L159
	case 1:
		v917 = v863
		v918 = v864
		v919 = v865
		goto L160
	case 2:
		v910 = v863
		v911 = v864
		v912 = v865
		goto L161
	case 3:
		v904 = v864
		v905 = v865
		goto L162
	case 4:
		v900 = v864
		v901 = v865
		goto L163
	case 5:
		v894 = v864
		v895 = v865
		goto L164
	case 6:
		v888 = v864
		v889 = v865
		goto L165
	case 7:
		v883 = v865
		goto L166
	case 8:
		v878 = v865
		goto L167
	case 9:
		v873 = v865
		goto L168
	case 10:
		goto L169
	default:
		v931 = v863
		v932 = v864
		v933 = v865
		goto L131
	}
L133:
	;
	v813 = v697
	v814 = v698
	v815 = v704
	v816 = v704
	v817 = v704
	goto L156
L134:
	;
	if base.Ui32(int32(11)) < base.Ui32(v698) {
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if base.Ui32(v698) < base.Ui32(int32(12)) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v861 = v697
	v862 = v698
	v863 = v704
	v864 = v704
	v865 = v704
	goto L132
L138:
	;
	switch v760 - int32(1) {
	case 0:
		v810 = v761
		goto L145
	case 1:
		v805 = v761
		goto L146
	case 2:
		goto L147
	case 3:
		v798 = v762
		goto L148
	case 4:
		v795 = v762
		goto L149
	case 5:
		v790 = v762
		goto L150
	case 6:
		goto L151
	case 7:
		v781 = v763
		goto L152
	case 8:
		v776 = v763
		goto L153
	case 9:
		v771 = v763
		goto L154
	case 10:
		goto L155
	default:
		v931 = v761
		v932 = v762
		v933 = v763
		goto L131
	}
L139:
	;
	v759 = v697
	v760 = v698
	v761 = v704
	v762 = v704
	v763 = v704
	goto L138
L140:
	;
	goto L141
L141:
	;
	v711 = v697
	v712 = v698
	v713 = v704
	v714 = v704
	v715 = v704
	goto L142
L142:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v718 = v717 + v714
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v711)+8))
	v722 = v721 + v715
	v724 = int32(4)
	v726 = v719 + v713 - v722 ^ base.I32_rotl(v722, v724)
	v730 = v718 - v726 ^ base.I32_rotl(v726, int32(6))
	v731 = v722 + v718
	v732 = v726 + v731
	v733 = v730 + v732
	v737 = v731 - v730 ^ base.I32_rotl(v730, int32(8))
	v741 = v732 - v737 ^ base.I32_rotl(v737, int32(16))
	v745 = v733 - v741 ^ base.I32_rotl(v741, int32(19))
	v746 = v737 + v733
	v747 = v741 + v746
	v748 = v745 + v747
	v752 = v746 - v745 ^ base.I32_rotl(v745, v724)
	v753 = int32(12)
	v754 = v711 + v753
	v756 = v712 - v753
	if base.Ui32(int32(11)) < base.Ui32(v756) {
		v711 = v754
		v712 = v756
		v713 = v747
		v714 = v748
		v715 = v752
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v759 = v754
	v760 = v756
	v761 = v747
	v762 = v748
	v763 = v752
	goto L138
L144:
	;
	goto L143
L145:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759))))
	v931 = v810 + v811
	v932 = v762
	v933 = v763
	goto L131
L146:
	;
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+1)))
	v810 = v806<<(uint(int32(8))%32) + v805
	goto L145
L147:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+2)))
	v805 = v801<<(uint(int32(16))%32) + v761
	goto L146
L148:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	v931 = v799 + v761
	v932 = v798
	v933 = v763
	goto L131
L149:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+4)))
	v798 = v795 + v796
	goto L148
L150:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+5)))
	v795 = v791<<(uint(int32(8))%32) + v790
	goto L149
L151:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+6)))
	v790 = v786<<(uint(int32(16))%32) + v762
	goto L150
L152:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v759)))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v759)+4))
	v931 = v782 + v761
	v932 = v784 + v762
	v933 = v781
	goto L131
L153:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+8)))
	v781 = v777<<(uint(int32(8))%32) + v776
	goto L152
L154:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+9)))
	v776 = v772<<(uint(int32(16))%32) + v771
	goto L153
L155:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+10)))
	v771 = v767<<(uint(int32(24))%32) + v763
	goto L154
L156:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	v820 = v819 + v816
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v813)+8))
	v824 = v823 + v817
	v826 = int32(4)
	v828 = v821 + v815 - v824 ^ base.I32_rotl(v824, v826)
	v832 = v820 - v828 ^ base.I32_rotl(v828, int32(6))
	v833 = v824 + v820
	v834 = v828 + v833
	v835 = v832 + v834
	v839 = v833 - v832 ^ base.I32_rotl(v832, int32(8))
	v843 = v834 - v839 ^ base.I32_rotl(v839, int32(16))
	v847 = v835 - v843 ^ base.I32_rotl(v843, int32(19))
	v848 = v839 + v835
	v849 = v843 + v848
	v850 = v847 + v849
	v854 = v848 - v847 ^ base.I32_rotl(v847, v826)
	v855 = int32(12)
	v856 = v813 + v855
	v858 = v814 - v855
	if base.Ui32(int32(11)) < base.Ui32(v858) {
		v813 = v856
		v814 = v858
		v815 = v849
		v816 = v850
		v817 = v854
		goto L156
	} else {
		goto L158
	}
L157:
	;
	v861 = v856
	v862 = v858
	v863 = v849
	v864 = v850
	v865 = v854
	goto L132
L158:
	;
	goto L157
L159:
	;
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861))))
	v931 = v924 + v927
	v932 = v925
	v933 = v926
	goto L131
L160:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+1)))
	v924 = v920<<(uint(int32(8))%32) + v917
	v925 = v918
	v926 = v919
	goto L159
L161:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+2)))
	v917 = v913<<(uint(int32(16))%32) + v910
	v918 = v911
	v919 = v912
	goto L160
L162:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+3)))
	v910 = v906<<(uint(int32(24))%32) + v863
	v911 = v904
	v912 = v905
	goto L161
L163:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+4)))
	v904 = v900 + v902
	v905 = v901
	goto L162
L164:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+5)))
	v900 = v896<<(uint(int32(8))%32) + v894
	v901 = v895
	goto L163
L165:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+6)))
	v894 = v890<<(uint(int32(16))%32) + v888
	v895 = v889
	goto L164
L166:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+7)))
	v888 = v884<<(uint(int32(24))%32) + v864
	v889 = v883
	goto L165
L167:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+8)))
	v883 = v879<<(uint(int32(8))%32) + v878
	goto L166
L168:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+9)))
	v878 = v874<<(uint(int32(16))%32) + v873
	goto L167
L169:
	;
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+10)))
	v873 = v869<<(uint(int32(24))%32) + v865
	goto L168
L170:
	;
	v980 = v964 & v963
	v985 = v336 + v980<<(uint(int32(4))%32)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)))
	if v986 != 0 {
		v964 = v980 + int32(1)
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v693)))
	*(*int64)(unsafe.Add(mBase, uint32(v985))) = v987
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v693)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v985)+8)) = v989
	goto L129
L172:
	;
	goto L171
L173:
	;
	v1012 = v1008
	goto L175
L174:
	;
	v1012 = int32(0)
	goto L175
L175:
	;
	v1014 = v682 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1014)) < base.Ui64(v330) {
		v678 = v1012
		v682 = v1014
		goto L125
	} else {
		goto L176
	}
L176:
	;
	goto L126
L177:
	;
	goto L52
L178:
	;
	F_errmsg_internal(m, int32(401843), int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L63
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(327643), int32(327), int32(342217))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L63
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
	F_errmsg_internal(m, int32(401843), int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L63
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(327643), int32(327), int32(342217))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L63
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v1090 = v1080
	v1091 = v1077
	v1094 = v1083
	goto L185
L185:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+4))
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1103))))
	if v1107 == int32(0) {
		v1126 = v1106
		v1127 = v1107
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1548 = v1526
	goto L46
L187:
	;
	if v1127-v1126 == int32(0) {
		goto L44
	} else {
		goto L195
	}
L188:
	;
	goto L187
L189:
	;
	if v1106 != v1107 {
		v1126 = v1106
		v1127 = v1107
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1111 = v1103
	v1112 = l1
	goto L191
L191:
	;
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112)+1)))
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111)+1)))
	if v1116 == int32(0) {
		v1126 = v1115
		v1127 = v1116
		goto L188
	} else {
		goto L193
	}
L192:
	;
	v1126 = v1115
	v1127 = v1116
	goto L188
L193:
	;
	v1119 = int32(1)
	if v1115 == v1116 {
		v1111 = v1111 + v1119
		v1112 = v1112 + v1119
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v1131 = F_strlen(m, v1103)
	mBase = m.M
	v1137 = v1131 - int32(1636608432)
	if v1103&int32(3) != 0 {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1397 = (v1391 ^ v1383 - base.I32_rotl(v1391, int32(24))) & v1396
	if base.Ui32(v1090) < base.Ui32(v1397) {
		goto L236
	} else {
		goto L237
	}
L197:
	;
	v1369 = int32(14)
	v1371 = v1365 ^ v1366 - base.I32_rotl(v1365, v1369)
	v1375 = v1371 ^ v1364 - base.I32_rotl(v1371, int32(11))
	v1379 = v1375 ^ v1365 - base.I32_rotl(v1375, int32(25))
	v1383 = v1379 ^ v1371 - base.I32_rotl(v1379, int32(16))
	v1387 = v1383 ^ v1375 - base.I32_rotl(v1383, int32(4))
	v1391 = v1387 ^ v1379 - base.I32_rotl(v1387, v1369)
	goto L196
L198:
	;
	switch v1295 - int32(1) {
	case 0:
		v1357 = v1296
		v1358 = v1297
		v1359 = v1298
		goto L225
	case 1:
		v1350 = v1296
		v1351 = v1297
		v1352 = v1298
		goto L226
	case 2:
		v1343 = v1296
		v1344 = v1297
		v1345 = v1298
		goto L227
	case 3:
		v1337 = v1297
		v1338 = v1298
		goto L228
	case 4:
		v1333 = v1297
		v1334 = v1298
		goto L229
	case 5:
		v1327 = v1297
		v1328 = v1298
		goto L230
	case 6:
		v1321 = v1297
		v1322 = v1298
		goto L231
	case 7:
		v1316 = v1298
		goto L232
	case 8:
		v1311 = v1298
		goto L233
	case 9:
		v1306 = v1298
		goto L234
	case 10:
		goto L235
	default:
		v1364 = v1296
		v1365 = v1297
		v1366 = v1298
		goto L197
	}
L199:
	;
	v1246 = v1103
	v1247 = v1131
	v1248 = v1137
	v1249 = v1137
	v1250 = v1137
	goto L222
L200:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1131) {
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v1131) < base.Ui32(int32(12)) {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	v1294 = v1103
	v1295 = v1131
	v1296 = v1137
	v1297 = v1137
	v1298 = v1137
	goto L198
L204:
	;
	switch v1193 - int32(1) {
	case 0:
		v1243 = v1194
		goto L211
	case 1:
		v1238 = v1194
		goto L212
	case 2:
		goto L213
	case 3:
		v1231 = v1195
		goto L214
	case 4:
		v1228 = v1195
		goto L215
	case 5:
		v1223 = v1195
		goto L216
	case 6:
		goto L217
	case 7:
		v1214 = v1196
		goto L218
	case 8:
		v1209 = v1196
		goto L219
	case 9:
		v1204 = v1196
		goto L220
	case 10:
		goto L221
	default:
		v1364 = v1194
		v1365 = v1195
		v1366 = v1196
		goto L197
	}
L205:
	;
	v1192 = v1103
	v1193 = v1131
	v1194 = v1137
	v1195 = v1137
	v1196 = v1137
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1144 = v1103
	v1145 = v1131
	v1146 = v1137
	v1147 = v1137
	v1148 = v1137
	goto L208
L208:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+4))
	v1151 = v1150 + v1147
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1144)))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+8))
	v1155 = v1154 + v1148
	v1157 = int32(4)
	v1159 = v1152 + v1146 - v1155 ^ base.I32_rotl(v1155, v1157)
	v1163 = v1151 - v1159 ^ base.I32_rotl(v1159, int32(6))
	v1164 = v1155 + v1151
	v1165 = v1159 + v1164
	v1166 = v1163 + v1165
	v1170 = v1164 - v1163 ^ base.I32_rotl(v1163, int32(8))
	v1174 = v1165 - v1170 ^ base.I32_rotl(v1170, int32(16))
	v1178 = v1166 - v1174 ^ base.I32_rotl(v1174, int32(19))
	v1179 = v1170 + v1166
	v1180 = v1174 + v1179
	v1181 = v1178 + v1180
	v1185 = v1179 - v1178 ^ base.I32_rotl(v1178, v1157)
	v1186 = int32(12)
	v1187 = v1144 + v1186
	v1189 = v1145 - v1186
	if base.Ui32(int32(11)) < base.Ui32(v1189) {
		v1144 = v1187
		v1145 = v1189
		v1146 = v1180
		v1147 = v1181
		v1148 = v1185
		goto L208
	} else {
		goto L210
	}
L209:
	;
	v1192 = v1187
	v1193 = v1189
	v1194 = v1180
	v1195 = v1181
	v1196 = v1185
	goto L204
L210:
	;
	goto L209
L211:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192))))
	v1364 = v1243 + v1244
	v1365 = v1195
	v1366 = v1196
	goto L197
L212:
	;
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+1)))
	v1243 = v1239<<(uint(int32(8))%32) + v1238
	goto L211
L213:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+2)))
	v1238 = v1234<<(uint(int32(16))%32) + v1194
	goto L212
L214:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1192)))
	v1364 = v1232 + v1194
	v1365 = v1231
	v1366 = v1196
	goto L197
L215:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+4)))
	v1231 = v1228 + v1229
	goto L214
L216:
	;
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+5)))
	v1228 = v1224<<(uint(int32(8))%32) + v1223
	goto L215
L217:
	;
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+6)))
	v1223 = v1219<<(uint(int32(16))%32) + v1195
	goto L216
L218:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1192)))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1192)+4))
	v1364 = v1215 + v1194
	v1365 = v1217 + v1195
	v1366 = v1214
	goto L197
L219:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+8)))
	v1214 = v1210<<(uint(int32(8))%32) + v1209
	goto L218
L220:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+9)))
	v1209 = v1205<<(uint(int32(16))%32) + v1204
	goto L219
L221:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+10)))
	v1204 = v1200<<(uint(int32(24))%32) + v1196
	goto L220
L222:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+4))
	v1253 = v1252 + v1249
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1246)))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1246)+8))
	v1257 = v1256 + v1250
	v1259 = int32(4)
	v1261 = v1254 + v1248 - v1257 ^ base.I32_rotl(v1257, v1259)
	v1265 = v1253 - v1261 ^ base.I32_rotl(v1261, int32(6))
	v1266 = v1257 + v1253
	v1267 = v1261 + v1266
	v1268 = v1265 + v1267
	v1272 = v1266 - v1265 ^ base.I32_rotl(v1265, int32(8))
	v1276 = v1267 - v1272 ^ base.I32_rotl(v1272, int32(16))
	v1280 = v1268 - v1276 ^ base.I32_rotl(v1276, int32(19))
	v1281 = v1272 + v1268
	v1282 = v1276 + v1281
	v1283 = v1280 + v1282
	v1287 = v1281 - v1280 ^ base.I32_rotl(v1280, v1259)
	v1288 = int32(12)
	v1289 = v1246 + v1288
	v1291 = v1247 - v1288
	if base.Ui32(int32(11)) < base.Ui32(v1291) {
		v1246 = v1289
		v1247 = v1291
		v1248 = v1282
		v1249 = v1283
		v1250 = v1287
		goto L222
	} else {
		goto L224
	}
L223:
	;
	v1294 = v1289
	v1295 = v1291
	v1296 = v1282
	v1297 = v1283
	v1298 = v1287
	goto L198
L224:
	;
	goto L223
L225:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294))))
	v1364 = v1357 + v1360
	v1365 = v1358
	v1366 = v1359
	goto L197
L226:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+1)))
	v1357 = v1353<<(uint(int32(8))%32) + v1350
	v1358 = v1351
	v1359 = v1352
	goto L225
L227:
	;
	v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+2)))
	v1350 = v1346<<(uint(int32(16))%32) + v1343
	v1351 = v1344
	v1352 = v1345
	goto L226
L228:
	;
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+3)))
	v1343 = v1339<<(uint(int32(24))%32) + v1296
	v1344 = v1337
	v1345 = v1338
	goto L227
L229:
	;
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+4)))
	v1337 = v1333 + v1335
	v1338 = v1334
	goto L228
L230:
	;
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+5)))
	v1333 = v1329<<(uint(int32(8))%32) + v1327
	v1334 = v1328
	goto L229
L231:
	;
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+6)))
	v1327 = v1323<<(uint(int32(16))%32) + v1321
	v1328 = v1322
	goto L230
L232:
	;
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+7)))
	v1321 = v1317<<(uint(int32(24))%32) + v1297
	v1322 = v1316
	goto L231
L233:
	;
	v1312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+8)))
	v1316 = v1312<<(uint(int32(8))%32) + v1311
	goto L232
L234:
	;
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+9)))
	v1311 = v1307<<(uint(int32(16))%32) + v1306
	goto L233
L235:
	;
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294)+10)))
	v1306 = v1302<<(uint(int32(24))%32) + v1298
	goto L234
L236:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1401 = v1090 + v1399
	goto L238
L237:
	;
	v1401 = v1090
	goto L238
L238:
	;
	v1404 = v1396 & (v1090 + int32(1))
	if base.Ui32(v1401-v1397) < base.Ui32(v1091) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1410 = v1078 + v1404<<(uint(int32(4))%32)
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1410)))
	if v1411 != 0 {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	v1513 = v1091 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1513) {
		goto L258
	} else {
		goto L259
	}
L242:
	;
	v1412 = v1404
	v1417 = int32(0)
	goto L245
L243:
	;
	v1447 = v1404
	v1451 = v1410
	goto L244
L244:
	;
	if v1447 != v1090 {
		goto L252
	} else {
		goto L253
	}
L245:
	;
	v1429 = v1417 + int32(1)
	if int32(151) <= v1429 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1447 = v1442
	v1451 = v1445
	goto L244
L247:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v1434 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1432), base.F64_convert_i64_u(v1434)), float64(0.1)) != 0 {
		v1611 = v1432
		goto L43
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1442 = (v1412 + int32(1)) & v1396
	v1445 = v1078 + v1442<<(uint(int32(4))%32)
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	if v1446 != 0 {
		v1412 = v1442
		v1417 = v1429
		goto L245
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	goto L246
L252:
	;
	v1464 = v1447
	v1468 = v1451
	goto L255
L253:
	;
	goto L254
L254:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1508 + int32(1)
	v1568 = v1094
	goto L45
L255:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1483 = v1480 & (v1464 - int32(1))
	v1486 = v1078 + v1483<<(uint(int32(4))%32)
	v1487 = *(*int64)(unsafe.Add(mBase, uint32(v1486)))
	*(*int64)(unsafe.Add(mBase, uint32(v1468))) = v1487
	v1489 = *(*int64)(unsafe.Add(mBase, uint32(v1486)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1468)+8)) = v1489
	if v1483 != v1090 {
		v1464 = v1483
		v1468 = v1486
		goto L255
	} else {
		goto L257
	}
L256:
	;
	goto L254
L257:
	;
	goto L256
L258:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v1518 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1516), base.F64_convert_i64_u(v1518)), float64(0.1)) != 0 {
		v1611 = v1516
		goto L43
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1526 = v1078 + v1404<<(uint(int32(4))%32)
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1526)))
	if v1527 != 0 {
		v1090 = v1404
		v1091 = v1513
		v1094 = v1526
		goto L185
	} else {
		goto L262
	}
L261:
	;
	goto L260
L262:
	;
	goto L186
L263:
	;
	F_errmsg_internal(m, int32(463912), int32(0))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L63
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(327643), int32(630), int32(312820))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L63
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1568)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+4)) = v1582
	goto L44
}
func F_markNullableIfNeeded(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 == int32(0) {
		v54 = l0
	} else {
		v12 = v8 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(v8) {
			v18 = l0
			v20 = int32(0)
			for {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v33 = v20 + int32(8)
				if v33 != v8&int32(-8) {
					v18 = v31
					v20 = v33
					continue
				} else {
					break
				}
				break
			}
			v35 = v31
		} else {
			v35 = l0
		}
		if v12 == int32(0) {
			v54 = v35
		} else {
			v44 = v35
			v46 = int32(0)
			for {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				v52 = v46 + int32(1)
				if v52 != v12 {
					v44 = v50
					v46 = v52
					continue
				} else {
					break
				}
				break
			}
			v54 = v50
		}
	}
	if v7 <= int32(0) {
		return
	} else {
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
		if v62 == int32(0) {
			return
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
			if v65 < v7 {
				return
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+v7<<(uint(int32(2))%32)-int32(4))))
				if v73 == int32(0) {
					return
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v77 = F_bms_union(m, v76, v73)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v77
						return
					}
				}
			}
		}
	}
}
func F_markcanreach(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = m.T0[v8].(func(*base.Module) int32)(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(101)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v19 != l2 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v17 = v15
	goto L8
L7:
	;
	v17 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v17
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l3
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = v22
	goto L12
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	F_markcanreach(m, l0, v29, l2, l3)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v32 != 0 {
		v26 = v32
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_markst(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v4 = v2 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v4)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 != 0 {
		v7 = v6
		for {
			v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			v10 = v8 | int32(64)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v10)
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
			if v12 != 0 {
				v13 = v12
				for {
					F_markst(m, v13)
					mBase = m.M
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
					if v15 != 0 {
						v13 = v15
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
			if v17 != 0 {
				v7 = v17
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_mask_page_lsn_and_checksum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_match_kind(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(v3 == l1)
}
func F_materialize_finished_plan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v42 float64
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 float64
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v135 float64
	_ = v135
	var v138 int32
	_ = v138
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 float64
	_ = v158
	var v162 float64
	_ = v162
	var v170 float64
	_ = v170
	var v176 float64
	_ = v176
	var v182 float64
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 float64
	_ = v194
	var v195 float64
	_ = v195
	var v198 float64
	_ = v198
	var v201 float64
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	v7 = m.G0
	v9 = v7 - int32(96)
	m.G0 = v9
	v12 = F_palloc0(m, int32(72))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(360)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v19 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v18
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v19
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
		v42 = float64(0)
		if v29 == v19 {
			v128 = v19
			v135 = v42
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			if v45 <= int32(0) {
				v128 = v19
				v135 = v42
			} else {
				v48 = int32(0)
				if v48 < v45 {
					v51 = v45
				} else {
					v51 = v48
				}
				v52 = int32(1)
				if v45 == v52 {
					v100 = int32(0)
					v101 = v19
					v108 = v42
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
					v64 = int32(0)
					v65 = v19
					v68 = v19
					v72 = v42
					for {
						v73 = int32(2)
						v75 = v59 + v64<<(uint(v73)%32)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v76)+56))
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v76)+64))
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
						v82 = *(*float64)(unsafe.Add(mBase, uint32(v81)+56))
						v83 = *(*float64)(unsafe.Add(mBase, uint32(v81)+64))
						v85 = base.F64_add(base.F64_add(v72, base.F64_add(v77, v78)), base.F64_add(v82, v83))
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+38)))
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+38)))
						v91 = v86&v87 ^ int32(1) | v65
						v93 = v64 + v73
						v95 = v68 + v73
						if v95 != v51&int32(2147483646) {
							v64 = v93
							v65 = v91
							v68 = v95
							v72 = v85
							continue
						} else {
							break
						}
						break
					}
					v100 = v93
					v101 = v91
					v108 = v85
				}
				if v51&v52 == int32(0) {
					v128 = v101
					v135 = v108
				} else {
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v100<<(uint(int32(2))%32))))
					v116 = *(*float64)(unsafe.Add(mBase, uint32(v115)+56))
					v117 = *(*float64)(unsafe.Add(mBase, uint32(v115)+64))
					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+38)))
					v128 = v120 ^ int32(1) | v101
					v135 = base.F64_add(v108, base.F64_add(v116, v117))
				}
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9+int32(16)))) = v135
		v138 = v128 & int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(15)))) = uint8(v138)
		v140 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v141 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		v142 = base.F64_sub(v140, v141)
		*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v142
		v144 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
		v145 = base.F64_sub(v144, v141)
		*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v145
		v147 = int32(24)
		v148 = v9 + v147
		v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v150 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v155 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		*(*float64)(unsafe.Add(mBase, uint32(v148)+32)) = v150
		v158 = *(*float64)(unsafe.Add(mBase, _consts[384]))
		v162 = base.F64_add(base.F64_mul(base.F64_add(v158, v158), v150), base.F64_sub(v145, v142))
		v170 = base.F64_mul(v150, base.F64_convert_i32_u((v151+int32(7))&int32(-8)+v147))
		if base.F64_gt(v170, base.F64_convert_i32_u(v155<<(uint(int32(10))%32))) != 0 {
			v176 = *(*float64)(unsafe.Add(mBase, _consts[386]))
			v182 = base.F64_add(base.F64_mul(v176, base.F64_ceil(base.F64_mul(v170, float64(0.0001220703125)))), v162)
		} else {
			v182 = v162
		}
		v184 = int32(*(*uint8)(unsafe.Add(mBase, _consts[390])))
		*(*float64)(unsafe.Add(mBase, uint32(v148)+56)) = base.F64_add(v142, v182)
		*(*float64)(unsafe.Add(mBase, uint32(v148)+48)) = v142
		*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v149 + (v184 ^ int32(1))
		v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v192
		v194 = *(*float64)(unsafe.Add(mBase, uint32(v9)+72))
		v195 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = base.F64_add(v194, v195)
		v198 = *(*float64)(unsafe.Add(mBase, uint32(v9)+80))
		*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = base.F64_add(v195, v198)
		v201 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v201
		v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v204 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+36)) = uint8(v204)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v203
		v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)) = uint8(v207)
		m.G0 = v9 + int32(96)
		return v12
	}
}
func F_maybe_start_bgworkers(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v107 int32
	_ = v107
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v239 int32
	_ = v239
	v1 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[456])) = uint8(v17)
	*(*uint8)(unsafe.Add(mBase, _consts[457])) = uint8(v1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[458])))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[459]))
	if v25 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v25 == int32(4126372) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = v25
	v38 = v1
	v39 = int64(0)
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v43 = v33 - int32(20)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v44 != 0 {
		v230 = v38
		v231 = v39
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[456])) = uint8(v239)
	goto L1
L7:
	;
	goto L6
L8:
	;
	if v41 != int32(4126372) {
		v33 = v41
		v38 = v230
		v39 = v231
		goto L5
	} else {
		goto L68
	}
L9:
	;
	v46 = v33 - int32(1480)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33-int32(4)))))
	if v49 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ForgetBackgroundWorker(m, v46)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v55 = v33 - int32(16)
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int64(0) {
		v111 = v39
		goto L15
	} else {
		goto L16
	}
L13:
	;
	return
L14:
	;
	v230 = v38
	v231 = v39
	goto L8
L15:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v33-int32(1284))))
	v116 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	switch v116 {
	case 0, 1, 2:
		goto L30
	case 3:
		goto L31
	case 4:
		goto L32
	default:
		v230 = v38
		v231 = v111
		goto L8
	}
L16:
	;
	v60 = v33 - int32(1280)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v33-int32(24))))
	F_ForgetBackgroundWorker(m, v46)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v39 == int64(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	if v66 == int32(0) {
		v230 = v38
		v231 = v39
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v72 = F_kill(m, v66, int32(10))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v230 = v38
	v231 = v39
	goto L8
L23:
	;
	v79 = m.G0
	v80 = int32(16)
	v81 = v79 - v80
	m.G0 = v81
	F___gettimeofday(m, v81)
	mBase = m.M
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	v85 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+8)))
	m.G0 = v81 + v80
	goto L26
L24:
	;
	v96 = v61
	v97 = v39
	v98 = v56
	goto L25
L25:
	;
	goto L27
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	v96 = v94
	v97 = v85 + v84*int64(1000000) - int64(946684800000000)
	v98 = v95
	goto L25
L27:
	;
	if base.I64_extend_i32_s(v96*int32(1000))*int64(1000) <= v97-v98 {
		v111 = v97
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[457])) = uint8(v107)
	v230 = v38
	v231 = v97
	goto L8
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = int64(0)
	v124 = F_AssignPostmasterChildSlot(m, int32(5))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L13
	} else {
		goto L38
	}
L30:
	;
	if v114 != 0 {
		v230 = v38
		v231 = v111
		goto L8
	} else {
		goto L35
	}
L31:
	;
	if base.Ui32(v114) < base.Ui32(int32(2)) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	if base.Ui32(v114) < base.Ui32(int32(3)) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v230 = v38
	v231 = v111
	goto L8
L34:
	;
	v230 = v38
	v231 = v111
	goto L8
L35:
	;
	goto L29
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v168
	v210 = *(*int32)(unsafe.Add(mBase, _consts[82]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v46)+1472))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v46)+1460))
	*(*int32)(unsafe.Add(mBase, uint32(v210+v211*int32(1480))+20)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v46)+1456))
	if v217 != 0 {
		goto L63
	} else {
		goto L64
	}
L37:
	;
	v191 = m.G0
	v192 = int32(16)
	v193 = v191 - v192
	m.G0 = v193
	F___gettimeofday(m, v193)
	mBase = m.M
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	v197 = int64(*(*int32)(unsafe.Add(mBase, uint32(v193)+8)))
	m.G0 = v193 + v192
	goto L62
L38:
	;
	if v124 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v130 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L13
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+16)) = uint8(v146)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+8)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v46
	v153 = F_errstart(m, int32(14), v146)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L47
	}
L42:
	;
	if v130 == int32(0) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(129763), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(496666), int32(4145), int32(221747))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	goto L37
L47:
	;
	if v153 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v46
	F_errmsg_internal(m, int32(701637), v14)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v168 = F_postmaster_child_launch(m, int32(5), v165, v46, int32(1460), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	F_errfinish(m, int32(496666), int32(4155), int32(221747))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	if v168 != int32(-1) {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	v174 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	if v174 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_errmsg(m, int32(293842), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v185 = F_ReleasePostmasterChildSlot(m, v124)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L13
	} else {
		goto L61
	}
L59:
	;
	F_errfinish(m, int32(496666), int32(4163), int32(221747))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	goto L37
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v197 + v196*int64(1000000) - int64(946684800000000)
	goto L7
L63:
	;
	v219 = F_kill(m, v217, int32(10))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L13
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v222 = v38 + int32(1)
	if int32(99) < v222 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v230 = v222
	v231 = v111
	goto L8
L68:
	;
	goto L1
}
func F_mbtowc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v104 int32
	_ = v104
	if l1 == int32(0) {
		return int32(0)
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v11 = base.I32_extend8_s(v10)
		if int32(0) <= v11 {
			if l0 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v10
			} else {
			}
			return base.B2i32(v11 != int32(0))
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == int32(0) {
				if l0 == int32(0) {
					v104 = int32(1)
					return v104
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11 & int32(57343)
					return int32(1)
				}
			} else {
				v32 = v10 - int32(194)
				if base.Ui32(int32(50)) < base.Ui32(v32) {
					*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(25)
					v104 = int32(-1)
					return v104
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v37 = int32(base.Ui32(v35) >> (uint(int32(3)) % 32))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[1135])))
					if base.Ui32(int32(7)) < base.Ui32(v37-int32(16)|(v37+v44>>(uint(int32(26))%32))) {
						*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(25)
						v104 = int32(-1)
						return v104
					} else {
						v55 = v35 - int32(128) | v44<<(uint(int32(6))%32)
						if int32(0) <= v55 {
							if l0 == int32(0) {
								v104 = int32(2)
								return v104
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v55
								return int32(2)
							}
						} else {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
							v66 = v64 - int32(128)
							if base.Ui32(int32(63)) < base.Ui32(v66) {
								*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(25)
								v104 = int32(-1)
								return v104
							} else {
								v70 = v55 << (uint(int32(6)) % 32)
								v71 = v66 | v70
								if int32(0) <= v70 {
									if l0 == int32(0) {
										v104 = int32(3)
										return v104
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v71
										return int32(3)
									}
								} else {
									v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
									v82 = v80 - int32(128)
									if base.Ui32(int32(63)) < base.Ui32(v82) {
										*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(25)
										v104 = int32(-1)
										return v104
									} else {
										if l0 == int32(0) {
											v104 = int32(4)
											return v104
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0))) = v82 | v71<<(uint(int32(6))%32)
											return int32(4)
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
func F_mcv_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v105 int32
	_ = v105
	var v107 float32
	_ = v107
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v132 int32
	_ = v132
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	v9 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v142
	m.G0 = v14 + int32(80)
	return v141
L2:
	;
	v141 = v9
	v142 = v9
	goto L1
L3:
	;
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v19 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v75 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l3
	goto L5
L7:
	;
	v59 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L20
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v22 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v32 = v16
	goto L10
L10:
	;
	v38 = F_get_attstatsslot(m, v14+int32(44), v32, int32(1), int32(0), int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L17
	}
L11:
	;
	v141 = v9
	v142 = v9
	goto L1
L12:
	;
	goto L13
L13:
	;
	v25 = F_get_func_leakproof(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return float64(0)
L15:
	;
	if v25 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = v31
	goto L10
L17:
	;
	if v38 == int32(0) {
		v141 = v9
		v142 = v9
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)) = uint8(v42)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+32)) = uint8(v42)
	v46 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+26)) = uint16(v46)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v42)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
	if l4 == v42 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l3
	goto L5
L20:
	;
	if v59 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v141 = v9
	v142 = v9
	goto L1
L22:
	;
	goto L23
L23:
	;
	v63 = F_get_func_name(m, v22)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v63
	F_errmsg_internal(m, int32(339378), v14)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(496247), int32(6242), int32(319062))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v141 = v9
	v142 = v9
	goto L1
L27:
	;
	F_free_attstatsslot(m, v14+int32(44))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L14
	} else {
		goto L45
	}
L28:
	;
	v126 = v9
	v127 = v9
	goto L27
L29:
	;
	goto L30
L30:
	;
	v79 = int32(0)
	v87 = v9
	v88 = v9
	goto L31
L31:
	;
	v91 = v79 << (uint(int32(2)) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91+v92)))
	if l4 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v126 = v112
	v127 = v113
	goto L27
L33:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)) = uint8(v97)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = m.T0[v102].(func(*base.Module, int32) int32)(m, v14+int32(8))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L14
	} else {
		goto L37
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v94
	goto L33
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v94
	goto L33
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v107 = *(*float32)(unsafe.Add(mBase, uint32(v105+v91)))
	v108 = base.F64_promote_f32(v107)
	if v103 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v110 = base.F64_add(v87, v108)
	goto L40
L39:
	;
	v110 = v87
	goto L40
L40:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
	if v111 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v112 = v87
	goto L43
L42:
	;
	v112 = v110
	goto L43
L43:
	;
	v113 = base.F64_add(v88, v108)
	v115 = v79 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v115 < v116 {
		v79 = v115
		v87 = v112
		v88 = v113
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	v141 = v126
	v142 = v127
	goto L1
}
func F_md5_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v16 == int32(1) {
			v19 = int32(4)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if v21&int32(254) == int32(2) {
				v30 = v19
			} else {
				v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
			}
			if v21 == int32(1) {
				v33 = v19
			} else {
				v33 = v30
			}
			v46 = v33
		} else {
			v34 = int32(1)
			if v16&v34 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = int32(1)
		if v16&v47 != 0 {
			v51 = v47
		} else {
			v51 = int32(4)
		}
		v57 = F_pg_md5_hash(m, v10+v51, v46, v5+int32(-48), v5+int32(-52))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			if v57 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(559109)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v70
						F_errmsg(m, int32(203280), v7)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496300), int32(71), int32(508499))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
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
				v82 = F_cstring_to_text(m, v5+int32(-48))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 - int32(-64)
					return v82
				}
			}
		}
	}
}
func F_mdwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int64
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	v7 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(1040)
	m.G0 = v23
	if l4 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L65
	}
L2:
	;
	m.G0 = v23 + int32(1040)
	return
L3:
	;
	v28 = F__mdfd_getseg(m, l0, l1, l2, l5, int32(9))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v33 = l2 & int32(131071)
	v34 = int32(131072) - v33
	if base.Ui32(l4) < base.Ui32(v34) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v36 = l4
	goto L8
L7:
	;
	v36 = v34
	goto L8
L8:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v36) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = int32(128)
	goto L11
L10:
	;
	v39 = v36
	goto L11
L11:
	;
	if v39 != l4 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = base.I64_extend_i32_u(v33 << (uint(int32(13)) % 32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v44
	v48 = int32(1)
	if base.Ui32(v36) < base.Ui32(int32(2)) {
		v174 = v48
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v190 = F_FileWriteV(m, v186, v23+int32(16), v174, v43, int32(167772184))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L35
	}
L14:
	;
	v51 = int32(1)
	v53 = l4 - v51
	if l4 == int32(2) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v53&v51 == int32(0) {
		v174 = v137
		goto L13
	} else {
		goto L30
	}
L16:
	;
	v135 = v23 + int32(16)
	v137 = v48
	v138 = v44
	v140 = v51
	goto L15
L17:
	;
	goto L18
L18:
	;
	v72 = v23 + int32(16)
	v74 = v48
	v75 = v44
	v77 = v51
	v84 = v7
	goto L19
L19:
	;
	v87 = v77 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3+v87)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v89 == v75+v90 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v135 = v121
	v137 = v122
	v138 = v123
	v140 = v125
	goto L15
L21:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l3+int32(4)+v87)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v107 != v105+v108 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v90 - int32(-8192)
	v103 = v72
	v104 = v74
	v105 = v75
	goto L21
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = int32(8192)
	v99 = v72 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v89
	v103 = v99
	v104 = v74 + int32(1)
	v105 = v89
	goto L21
L25:
	;
	v124 = int32(2)
	v125 = v77 + v124
	v127 = v84 + v124
	if v127 != v53&int32(-2) {
		v72 = v121
		v74 = v122
		v75 = v123
		v77 = v125
		v84 = v127
		goto L19
	} else {
		goto L29
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = int32(8192)
	v114 = v103 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v107
	v121 = v114
	v122 = v104 + int32(1)
	v123 = v107
	goto L25
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v108 - int32(-8192)
	v121 = v103
	v122 = v104
	v123 = v105
	goto L25
L29:
	;
	goto L20
L30:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l3+v140<<(uint(int32(2))%32))))
	if v138+v151 != v156 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = int32(8192)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v156
	v174 = v137 + int32(1)
	goto L13
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v151 - int32(-8192)
	v174 = v137
	goto L13
L34:
	;
	if l5 != 0 {
		goto L2
	} else {
		goto L62
	}
L35:
	;
	if int32(0) <= v190 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v203 = v190
	v205 = v174
	v206 = int32(0)
	v216 = v43
	goto L39
L37:
	;
	goto L38
L38:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L53
	}
L39:
	;
	v217 = v203 + v206
	if v217 == l4<<(uint(int32(13))%32) {
		goto L34
	} else {
		goto L41
	}
L40:
	;
	goto L38
L41:
	;
	v220 = v23 + int32(16)
	v225 = v220
	v226 = v205
	v227 = v203
	goto L44
L42:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v255 = v216 + base.I64_extend_i32_u(v203)
	v257 = F_FileWriteV(m, v251, v23+int32(16), v250, v255, int32(167772184))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L51
	}
L43:
	;
	if v220 != v225 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if base.Ui32(v227) < base.Ui32(v229) {
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v250 = int32(0)
	goto L42
L46:
	;
	v235 = v226 - int32(1)
	if v235 != 0 {
		v225 = v225 + int32(8)
		v226 = v235
		v227 = v227 - v229
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v240 = F_memmove(m, v220, v225, v226<<(uint(int32(3))%32))
	mBase = m.M
	goto L50
L49:
	;
	goto L50
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v241 + v227
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v244 - v227
	v250 = v226
	goto L42
L51:
	;
	if int32(0) <= v257 {
		v203 = v257
		v205 = v250
		v206 = v217
		v216 = v255
		goto L39
	} else {
		goto L52
	}
L52:
	;
	goto L40
L53:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v291 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v289*int32(48))+32))
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l2 + l4 - int32(1)
	F_errmsg(m, int32(299359), v23)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if v282 == int32(51) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_errhint(m, int32(646394), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_errfinish(m, int32(501173), int32(1134), int32(36005))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v316 != int32(-1) {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	F_register_dirty_segment(m, l0, l1, v28)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L2
L65:
	;
	F_errmsg_internal(m, int32(18216), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(501173), int32(1092), int32(36005))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_memchr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v4 = int32(0)
	v7 = base.B2i32(l2 != v4)
	if l0&int32(3) == v4 {
		v33 = l0
		v35 = l2
		v36 = v7
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v84 = v77
	v86 = v79
	goto L20
L3:
	;
	if v36 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	if l2 == int32(0) {
		v33 = l0
		v35 = l2
		v36 = v7
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v16 = l0
	v18 = l2
	goto L6
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v21 == l1&int32(255) {
		v77 = v16
		v79 = v18
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v33 = v28
	v35 = v24
	v36 = v26
	goto L3
L8:
	;
	v23 = int32(1)
	v24 = v18 - v23
	v25 = int32(0)
	v26 = base.B2i32(v24 != v25)
	v28 = v16 + v23
	if v28&int32(3) == v25 {
		v33 = v28
		v35 = v24
		v36 = v26
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if v24 != 0 {
		v16 = v28
		v18 = v24
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v40 == l1&int32(255) {
		v70 = v33
		v72 = v35
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v72 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L13:
	;
	if base.Ui32(v35) < base.Ui32(int32(4)) {
		v70 = v33
		v72 = v35
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v50 = v33
	v52 = v35
	goto L15
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v57 = v56 ^ l1&int32(255)*int32(16843009)
	v60 = int32(-2139062144)
	if (int32(16843008)-v57|v57)&v60 != v60 {
		v77 = v50
		v79 = v52
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v70 = v65
	v72 = v67
	goto L12
L17:
	;
	v64 = int32(4)
	v65 = v50 + v64
	v67 = v52 - v64
	if base.Ui32(int32(3)) < base.Ui32(v67) {
		v50 = v65
		v52 = v67
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v77 = v70
	v79 = v72
	goto L2
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if l1&int32(255) == v89 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L1
L22:
	;
	return v84
L23:
	;
	goto L24
L24:
	;
	v92 = int32(1)
	v95 = v86 - v92
	if v95 != 0 {
		v84 = v84 + v92
		v86 = v95
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
}
func F_memmove(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v7 = l0 + l2
	if base.Ui32(l1-v7) <= base.Ui32(int32(0)-l2<<(uint(int32(1))%32)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v19 = (l0 ^ l1) & int32(3)
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	return v15
L7:
	;
	v14 = F__emscripten_memcpy_bulkmem(m, l0, l1, l2)
	mBase = m.M
	v15 = v14
	goto L9
L8:
	;
	v15 = l0
	goto L9
L9:
	;
	goto L6
L10:
	;
	if v121 == int32(0) {
		goto L1
	} else {
		goto L46
	}
L11:
	;
	if base.Ui32(v99) <= base.Ui32(int32(3)) {
		v120 = v98
		v121 = v99
		v122 = v100
		goto L10
	} else {
		goto L42
	}
L12:
	;
	if v19 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v19 != 0 {
		v81 = l2
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v120 = l1
	v121 = l2
	v122 = l0
	goto L10
L16:
	;
	goto L17
L17:
	;
	if l0&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v98 = l1
	v99 = l2
	v100 = l0
	goto L11
L19:
	;
	goto L20
L20:
	;
	v26 = l1
	v27 = l2
	v28 = l0
	goto L21
L21:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v98 = v35
	v99 = v37
	v100 = v39
	goto L11
L23:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v32)
	v34 = int32(1)
	v35 = v26 + v34
	v37 = v27 - v34
	v39 = v28 + v34
	if v39&int32(3) != 0 {
		v26 = v35
		v27 = v37
		v28 = v39
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v81 == int32(0) {
		goto L1
	} else {
		goto L38
	}
L26:
	;
	if v7&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v46 = l2
	goto L30
L28:
	;
	v61 = l2
	goto L29
L29:
	;
	if base.Ui32(v61) <= base.Ui32(int32(3)) {
		v81 = v61
		goto L25
	} else {
		goto L34
	}
L30:
	;
	if v46 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v61 = v52
	goto L29
L32:
	;
	v52 = v46 - int32(1)
	v53 = l0 + v52
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v52))))
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v55)
	if v53&int32(3) != 0 {
		v46 = v52
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v68 = v61
	goto L35
L35:
	;
	v72 = v68 - int32(4)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1+v72)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v72))) = v75
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v68 = v72
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v81 = v72
	goto L25
L37:
	;
	goto L36
L38:
	;
	v88 = v81
	goto L39
L39:
	;
	v92 = v88 - int32(1)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v92))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v92))) = uint8(v95)
	if v92 != 0 {
		v88 = v92
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L1
L41:
	;
	goto L40
L42:
	;
	v105 = v98
	v106 = v99
	v107 = v100
	goto L43
L43:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v109
	v111 = int32(4)
	v112 = v105 + v111
	v114 = v107 + v111
	v116 = v106 - v111
	if base.Ui32(int32(3)) < base.Ui32(v116) {
		v105 = v112
		v106 = v116
		v107 = v114
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v120 = v112
	v121 = v116
	v122 = v114
	goto L10
L45:
	;
	goto L44
L46:
	;
	v127 = v120
	v128 = v121
	v129 = v122
	goto L47
L47:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v131)
	v133 = int32(1)
	v138 = v128 - v133
	if v138 != 0 {
		v127 = v127 + v133
		v128 = v138
		v129 = v129 + v133
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L1
L49:
	;
	goto L48
}
func F_minmax_get_strategy_procinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = l1 - int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0+v14<<(uint(int32(2))%32))+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if l2 != v20 {
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = l2
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v23
		*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v23
	} else {
	}
	v39 = l3*int32(28) + v19 - int32(24)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v40 == int32(0) {
		v43 = int32(4)
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v14<<(uint(int32(2))%32))))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		v59 = v50 + v51<<(uint(v43)%32) + v14*int32(100) + int32(88)
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
		v62 = F_SearchSysCache4(m, v43, v49, v60, l2, base.I32_extend16_s(l3))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			if v62 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v91
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_errmsg_internal(m, int32(39996), v11)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493925), int32(301), int32(242165))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v70 = F_SysCacheGetAttrNotNull(m, int32(4), v62, int32(7))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v62)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = F_get_opcode(m, v70)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_fmgr_info_cxt(m, v74, v39, v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(16)
								return v39
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v11 + int32(16)
		return v39
	}
}
func F_minmax_multi_get_procinfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v15 == int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = base.I32_extend16_s(l1)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+216))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
		v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v21+v23*(v19-int32(1))<<(uint(int32(2))%32)+int32(44)-int32(4))))
		if v35 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(117833860))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(250910), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
						F_errdetail_internal(m, int32(655998), v8)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499163), int32(2886), int32(242303))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
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
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v40 = F_index_getprocinfo(m, v38, v19, int32(11))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v47 = v14 + int32(16)
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v47))) = v48
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v52
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(0)
				m.G0 = v8 + int32(16)
				return v14
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return v14
	}
}
func F_mkVoidAffix(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	v4 = int32(0)
	if l1 != 0 {
		v15 = l2
	} else {
		v15 = v4
	}
	if l1 != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v19 = v18
		v20 = int32(12)
	} else {
		v19 = l2
		v20 = int32(16)
	}
	v22 = F_palloc0(m, int32(16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(3)
		v26 = l0 + v20
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v22
		if v19 <= v15 {
			return
		} else {
			v31 = v19 - v15
			v32 = int32(3)
			v33 = v31 & v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v35 = int32(0)
			v38 = v19 + (v15 ^ int32(-1))
			if base.Ui32(v32) <= base.Ui32(v38) {
				v44 = v15
				v46 = v35
				v51 = v4
				for {
					v58 = v34 + v44*int32(24)
					v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+5)))
					v60 = int32(65532)
					v62 = int32(0)
					v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+29)))
					v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+53)))
					v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+77)))
					v82 = v46 + base.B2i32(v59&v60 == v62) + base.B2i32(v65&v60 == v62) + base.B2i32(v71&v60 == v62) + base.B2i32(v77&v60 == v62)
					v83 = int32(4)
					v84 = v44 + v83
					v86 = v51 + v83
					if v86 != v31&int32(-4) {
						v44 = v84
						v46 = v82
						v51 = v86
						continue
					} else {
						break
					}
					break
				}
				v89 = v84
				v91 = v82
			} else {
				v89 = v15
				v91 = v35
			}
			if v33 != 0 {
				v102 = v89
				v104 = v91
				v107 = v4
				for {
					v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v102*int32(24))+5)))
					v122 = v104 + base.B2i32(v117&int32(65532) == int32(0))
					v123 = int32(1)
					v126 = v107 + v123
					if v126 != v33 {
						v102 = v102 + v123
						v104 = v122
						v107 = v126
						continue
					} else {
						break
					}
					break
				}
				v131 = v122
			} else {
				v131 = v91
			}
			if v131 == int32(0) {
				return
			} else {
				v144 = v131 << (uint(int32(2)) % 32)
				if base.Ui32(int32(1025)) <= base.Ui32(v144) {
					v147 = F_palloc0(m, v144)
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return
					} else {
						v166 = v147
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v166
						v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v170 | v131<<(uint(int32(8))%32)
						if v38 == int32(0) {
							v237 = int32(0)
							v241 = v15
						} else {
							v182 = int32(0)
							v185 = v182
							v187 = v182
							v189 = v15
							for {
								v198 = v189 * int32(24)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v200 = v198 + v199
								v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+5)))
								if v201&int32(65532) == int32(0) {
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v206+v185<<(uint(int32(2))%32)))) = v200
									v213 = v185 + int32(1)
								} else {
									v213 = v185
								}
								v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v215 = v214 + v198
								v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+29)))
								if v216&int32(65532) == int32(0) {
									v221 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v221+v213<<(uint(int32(2))%32)))) = v215 + int32(24)
									v230 = v213 + int32(1)
								} else {
									v230 = v213
								}
								v231 = int32(2)
								v232 = v189 + v231
								v234 = v187 + v231
								if v234 != v31&int32(-2) {
									v185 = v230
									v187 = v234
									v189 = v232
									continue
								} else {
									break
								}
								break
							}
							v237 = v230
							v241 = v232
						}
						if v31&int32(1) == int32(0) {
						} else {
							v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v254 = v251 + v241*int32(24)
							v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+5)))
							if v255&int32(65532) != 0 {
							} else {
								v258 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v258+v237<<(uint(int32(2))%32)))) = v254
							}
						}
						return
					}
				} else {
					v152 = (v144 + int32(7)) & int32(4088)
					v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
					if base.Ui32(v152) <= base.Ui32(v153) {
						v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
						v160 = v153
						v161 = v155
						*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v160 - v152
						*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v161 + v152
						v166 = v161
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v166
						v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v170 | v131<<(uint(int32(8))%32)
						if v38 == int32(0) {
							v237 = int32(0)
							v241 = v15
						} else {
							v182 = int32(0)
							v185 = v182
							v187 = v182
							v189 = v15
							for {
								v198 = v189 * int32(24)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v200 = v198 + v199
								v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+5)))
								if v201&int32(65532) == int32(0) {
									v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v206+v185<<(uint(int32(2))%32)))) = v200
									v213 = v185 + int32(1)
								} else {
									v213 = v185
								}
								v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v215 = v214 + v198
								v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+29)))
								if v216&int32(65532) == int32(0) {
									v221 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v221+v213<<(uint(int32(2))%32)))) = v215 + int32(24)
									v230 = v213 + int32(1)
								} else {
									v230 = v213
								}
								v231 = int32(2)
								v232 = v189 + v231
								v234 = v187 + v231
								if v234 != v31&int32(-2) {
									v185 = v230
									v187 = v234
									v189 = v232
									continue
								} else {
									break
								}
								break
							}
							v237 = v230
							v241 = v232
						}
						if v31&int32(1) == int32(0) {
						} else {
							v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v254 = v251 + v241*int32(24)
							v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+5)))
							if v255&int32(65532) != 0 {
							} else {
								v258 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v258+v237<<(uint(int32(2))%32)))) = v254
							}
						}
						return
					} else {
						v156 = int32(8192)
						v158 = F_palloc0(m, v156)
						mBase = m.M
						v159 = m.ExcPending
						if v159 != 0 {
							return
						} else {
							v160 = v156
							v161 = v158
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v160 - v152
							*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v161 + v152
							v166 = v161
							*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v166
							v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v170 | v131<<(uint(int32(8))%32)
							if v38 == int32(0) {
								v237 = int32(0)
								v241 = v15
							} else {
								v182 = int32(0)
								v185 = v182
								v187 = v182
								v189 = v15
								for {
									v198 = v189 * int32(24)
									v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v200 = v198 + v199
									v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+5)))
									if v201&int32(65532) == int32(0) {
										v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v206+v185<<(uint(int32(2))%32)))) = v200
										v213 = v185 + int32(1)
									} else {
										v213 = v185
									}
									v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v215 = v214 + v198
									v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+29)))
									if v216&int32(65532) == int32(0) {
										v221 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v221+v213<<(uint(int32(2))%32)))) = v215 + int32(24)
										v230 = v213 + int32(1)
									} else {
										v230 = v213
									}
									v231 = int32(2)
									v232 = v189 + v231
									v234 = v187 + v231
									if v234 != v31&int32(-2) {
										v185 = v230
										v187 = v234
										v189 = v232
										continue
									} else {
										break
									}
									break
								}
								v237 = v230
								v241 = v232
							}
							if v31&int32(1) == int32(0) {
							} else {
								v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v254 = v251 + v241*int32(24)
								v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+5)))
								if v255&int32(65532) != 0 {
								} else {
									v258 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v258+v237<<(uint(int32(2))%32)))) = v254
								}
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_mode_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v2 = int32(0)
	v10 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v2
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v162
L2:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v157)
	v162 = int32(0)
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	if v20 == int64(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v25 = v23 + int32(76)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v30 = F_get_opcode(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v39 = v23
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+58)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	F_fmgr_info_cxt(m, v30, v25, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v39 = v38
	goto L7
L11:
	;
	v51 = int32(0)
	v55 = v51
	v56 = v51
	v58 = v2
	v60 = v2
	v62 = v10
	v63 = v10
	goto L17
L12:
	;
	F_tuplesort_performsort(m, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_tuplesort_rescan(m, v40)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v47)
	goto L11
L16:
	;
	goto L11
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v65 = int32(1)
	v73 = F_tuplesort_getdatum(m, v64, v65, v65, v14+int32(12), v14+int32(11), v14+int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L22
	}
L19:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v141 == int32(0) {
		v55 = v134
		v56 = v135
		v58 = v136
		v60 = v137
		v62 = v138
		v63 = v139
		goto L17
	} else {
		goto L53
	}
L20:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_pfree(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L52
	}
L21:
	;
	if v41&int32(1) != 0 {
		v134 = v55
		v135 = v56
		v136 = v58
		v137 = v60
		v138 = v123
		v139 = v124
		goto L19
	} else {
		goto L51
	}
L22:
	;
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v75 != 0 {
		goto L17
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if (v56|v41)&int32(1) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L26:
	;
	if v63 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = int64(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v134 = v82
	v135 = int32(1)
	v136 = v82
	v137 = v80
	v138 = v79
	v139 = v79
	goto L19
L28:
	;
	goto L29
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v83 != v60 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if (v56|v41)&int32(1) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v87 = F_FunctionCall2Coll(m, v25, v85, v86, v58)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	if v87 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	if v56&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v123 = v62 + int64(1)
	v124 = v63
	goto L21
L35:
	;
	goto L36
L36:
	;
	v96 = v63 + int64(1)
	if v96 <= v62 {
		v123 = v62
		v124 = v96
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v98 = int32(1)
	if v41&v98 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v134 = v58
	v135 = v98
	v136 = v58
	v137 = v60
	v138 = v96
	v139 = v96
	goto L19
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v55)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v127 = v58
	v128 = v98
	v129 = v96
	v130 = v96
	goto L20
L42:
	;
	F_pfree(m, v58)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v134 = v55
	v135 = int32(0)
	v136 = v113
	v137 = v112
	v138 = v62
	v139 = int64(1)
	goto L19
L45:
	;
	goto L44
L46:
	;
	F_pfree(m, v58)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v62 != int64(0) {
		v162 = v55
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L2
L51:
	;
	v127 = v55
	v128 = v56
	v129 = v123
	v130 = v124
	goto L20
L52:
	;
	v134 = v127
	v135 = v128
	v136 = v58
	v137 = v60
	v138 = v129
	v139 = v130
	goto L19
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v55 = v134
	v56 = v135
	v58 = v136
	v60 = v137
	v62 = v138
	v63 = v139
	goto L17
}
func F_mul_size(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = int32(0)
	if l0 == v3 {
		v16 = v3
		return v16
	} else {
		if l1 == int32(0) {
			v16 = v3
			return v16
		} else {
			v11 = base.I64_extend_i32_u(l1) * base.I64_extend_i32_u(l0)
			if base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(112965), int32(0))
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498656), int32(521), int32(341863))
							v36 = m.ExcPending
							if v36 != 0 {
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
				v16 = base.I32_wrap_i64(v11)
				return v16
			}
		}
	}
}
func F_multirangesel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 float64
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 float64
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v138 float64
	_ = v138
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 float32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 float32
	_ = v284
	var v288 int32
	_ = v288
	var v290 float64
	_ = v290
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int64
	_ = v412
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v454 float64
	_ = v454
	var v455 int32
	_ = v455
	var v459 float64
	_ = v459
	var v460 int32
	_ = v460
	var v465 float64
	_ = v465
	var v466 int32
	_ = v466
	var v472 float64
	_ = v472
	var v473 int32
	_ = v473
	var v478 float64
	_ = v478
	var v479 int32
	_ = v479
	var v484 float64
	_ = v484
	var v485 int32
	_ = v485
	var v490 float64
	_ = v490
	var v491 int32
	_ = v491
	var v496 float64
	_ = v496
	var v497 int32
	_ = v497
	var v502 float64
	_ = v502
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 float64
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 float64
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 float64
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 float64
	_ = v553
	var v554 int32
	_ = v554
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v603 float64
	_ = v603
	var v604 int32
	_ = v604
	var v607 float64
	_ = v607
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v637 float64
	_ = v637
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v677 float64
	_ = v677
	var v683 float64
	_ = v683
	var v693 float64
	_ = v693
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v740 float64
	_ = v740
	var v745 float64
	_ = v745
	var v748 float64
	_ = v748
	var v762 float64
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v802 float64
	_ = v802
	var v808 float64
	_ = v808
	var v823 float64
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 float64
	_ = v828
	var v846 float64
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	v10 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(176)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	v31 = F_get_restriction_variable(m, v22, v21, v20, v17+int32(44), v17+int32(40), v17+int32(39))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v850 = F_Float8GetDatum(m, v846)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L2
	} else {
		goto L251
	}
L2:
	;
	return int32(0)
L3:
	;
	if v31 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = float64(0.01)
	if v19 <= int32(4141) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	goto L6
L6:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 != int32(7) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v846 = v86
	goto L1
L8:
	;
	v86 = v80
	goto L7
L9:
	;
	v80 = float64(0.3333333333333333)
	goto L8
L10:
	;
	if v19 == int32(3585) {
		goto L9
	} else {
		goto L22
	}
L11:
	;
	v86 = float64(0.005)
	goto L7
L12:
	;
	v44 = v19 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v44) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if base.Ui32(v19-int32(4395)) < base.Ui32(int32(6)) {
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v48 = int32(1) << (uint(v44) % 32)
	if v48&int32(57359) != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v48&int32(6912) != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if int32(1)<<(uint(v44)%32)&int32(1152) == int32(0) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	if base.Ui32(v19-int32(4539)) < base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v19 != int32(4142) {
		v80 = v40
		goto L8
	} else {
		goto L21
	}
L21:
	;
	goto L9
L22:
	;
	if v19 != int32(4035) {
		v80 = v40
		goto L8
	} else {
		goto L23
	}
L23:
	;
	goto L9
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v91 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+24)))
	if v145 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	m.T0[v92].(func(*base.Module, int32))(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v98 = float64(0.01)
	if v19 <= int32(4141) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v846 = v144
	goto L1
L32:
	;
	v144 = v138
	goto L31
L33:
	;
	v138 = float64(0.3333333333333333)
	goto L32
L34:
	;
	if v19 == int32(3585) {
		goto L33
	} else {
		goto L46
	}
L35:
	;
	v144 = float64(0.005)
	goto L31
L36:
	;
	v102 = v19 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v102) {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(v19-int32(4395)) < base.Ui32(int32(6)) {
		goto L33
	} else {
		goto L43
	}
L39:
	;
	v106 = int32(1) << (uint(v102) % 32)
	if v106&int32(57359) != 0 {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	if v106&int32(6912) != 0 {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	if int32(1)<<(uint(v102)%32)&int32(1152) == int32(0) {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	if base.Ui32(v19-int32(4539)) < base.Ui32(int32(2)) {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	if v19 != int32(4142) {
		v138 = v98
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	if v19 != int32(4035) {
		v138 = v98
		goto L32
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v148 == int32(0) {
		v846 = v10
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+39)))
	if v154 != 0 {
		v164 = v19
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	m.T0[v151].(func(*base.Module, int32))(m, v148)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v846 = v10
	goto L1
L53:
	;
	if v164 <= int32(4034) {
		goto L65
	} else {
		goto L66
	}
L54:
	;
	v155 = F_get_commutator(m, v19)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	if v155 != 0 {
		v164 = v155
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v157 = float64(0.01)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v158 == int32(0) {
		v846 = v157
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	m.T0[v161].(func(*base.Module, int32))(m, v158)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	v846 = v157
	goto L1
L59:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v824 != 0 {
		goto L245
	} else {
		goto L246
	}
L60:
	;
	v762 = float64(0.01)
	if v164 <= int32(4141) {
		goto L233
	} else {
		goto L234
	}
L61:
	;
	if v262 == int32(0) {
		goto L60
	} else {
		goto L87
	}
L62:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	if v250 != v251 {
		goto L60
	} else {
		goto L84
	}
L63:
	;
	if v164 == int32(3585) {
		goto L60
	} else {
		goto L83
	}
L64:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v226 = F_multirange_get_typcache(m, l0, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L79
	}
L65:
	;
	v168 = v164 - int32(2866)
	if base.Ui32(int32(10)) < base.Ui32(v168) {
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	switch v164 - int32(4395) {
	case 0, 3:
		goto L60
	case 1, 4:
		goto L64
	case 2:
		goto L62
	default:
		goto L76
	}
L68:
	;
	v172 = int32(1) << (uint(v168) % 32)
	if v172&int32(705) != 0 {
		goto L60
	} else {
		goto L69
	}
L69:
	;
	if v172&int32(1042) != 0 {
		goto L64
	} else {
		goto L70
	}
L70:
	;
	if v168 != int32(3) {
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v180 = F_multirange_get_typcache(m, l0, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+296))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+200))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v183 != v186 {
		goto L60
	} else {
		goto L73
	}
L73:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+145)) = uint8(v188)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+146)) = uint8(v188)
	v193 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)) = uint8(v193)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+140)) = v190
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+110)) = uint8(v193)
	v198 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+108)) = uint16(v198)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v190
	v207 = F_range_serialize(m, v184, v17+int32(140), v17+int32(104), v193, v193)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v180)+296))
	v215 = F_make_multirange(m, v210, v211, int32(1), v17+int32(32))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v259 = v180
	v262 = v215
	goto L61
L76:
	;
	switch v164 - int32(4539) {
	case 0:
		goto L60
	case 1:
		goto L64
	default:
		goto L77
	}
L77:
	;
	if v164 != int32(4035) {
		goto L62
	} else {
		goto L78
	}
L78:
	;
	goto L64
L79:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+296))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v229 != v231 {
		goto L60
	} else {
		goto L80
	}
L80:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)+20))
	v234 = F_pg_detoast_datum(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v234
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)+296))
	v242 = F_make_multirange(m, v237, v238, int32(1), v17+int32(32))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v259 = v226
	v262 = v242
	goto L61
L83:
	;
	goto L62
L84:
	;
	v253 = F_multirange_get_typcache(m, l0, v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v257 = F_pg_detoast_datum(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	v259 = v253
	v262 = v257
	goto L61
L87:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v265 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L88:
	;
	v745 = float64(0)
	v748 = base.F64_mul(base.F64_sub(float64(1), v297), v740)
	if base.F64_lt(v748, v745) != 0 {
		v823 = v745
		goto L59
	} else {
		goto L226
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L2
	} else {
		goto L223
	}
L90:
	;
	if v164 == int32(4142) {
		v740 = v10
		goto L88
	} else {
		goto L221
	}
L91:
	;
	if base.B2i32(v164 != int32(4540))&base.B2i32(v164 != int32(2874)) == int32(0) {
		goto L218
	} else {
		goto L219
	}
L92:
	;
	v637 = float64(0.01)
	if v164 <= int32(4141) {
		goto L206
	} else {
		goto L207
	}
L93:
	;
	F_free_attstatsslot(m, v17+int32(104))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L2
	} else {
		goto L198
	}
L94:
	;
	v603 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v349, v343, int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L2
	} else {
		goto L197
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L2
	} else {
		goto L194
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L2
	} else {
		goto L191
	}
L97:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	if v298 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L98:
	;
	v294 = v10
	v297 = float64(0)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+22)))
	v272 = *(*float32)(unsafe.Add(mBase, uint32(v269+v270)+8))
	v278 = F_get_attstatsslot(m, v17+int32(140), v265, int32(6), int32(0), int32(2))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	if v278 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v17)+164))
	if v280 != int32(1) {
		goto L96
	} else {
		goto L105
	}
L103:
	;
	v290 = v10
	goto L104
L104:
	;
	v294 = v290
	v297 = base.F64_promote_f32(v272)
	goto L97
L105:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v17)+160))
	v284 = *(*float32)(unsafe.Add(mBase, uint32(v283)))
	F_free_attstatsslot(m, v17+int32(140))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v290 = base.F64_promote_f32(v284)
	goto L104
L107:
	;
	if v164 <= int32(4141) {
		goto L114
	} else {
		goto L115
	}
L108:
	;
	goto L109
L109:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v259)+296))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+216))
	v320 = F_statistic_proc_security_check(m, v17+int32(44), v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L120
	}
L110:
	;
	if v164 == int32(4035) {
		v740 = v10
		goto L88
	} else {
		goto L119
	}
L111:
	;
	v740 = base.F64_sub(float64(1), v294)
	goto L88
L112:
	;
	v740 = float64(1)
	goto L88
L113:
	;
	v740 = v294
	goto L88
L114:
	;
	switch v164 - int32(2862) {
	case 0, 5, 6, 14, 15:
		v740 = v10
		goto L88
	case 1, 12:
		goto L113
	case 2, 8, 9:
		goto L112
	case 3:
		goto L111
	case 4, 7, 10, 11, 13:
		goto L89
	default:
		goto L110
	}
L115:
	;
	goto L116
L116:
	;
	v306 = v164 - int32(4396)
	if base.Ui32(int32(4)) < base.Ui32(v306) {
		goto L90
	} else {
		goto L117
	}
L117:
	;
	if v306 == int32(2) {
		goto L90
	} else {
		goto L118
	}
L118:
	;
	v740 = v10
	goto L88
L119:
	;
	goto L89
L120:
	;
	if v320 == int32(0) {
		goto L92
	} else {
		goto L121
	}
L121:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)+272))
	if v324 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v327 = F_statistic_proc_security_check(m, v17+int32(44), v324)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v331 == int32(0) {
		goto L92
	} else {
		goto L127
	}
L125:
	;
	if v327 == int32(0) {
		goto L92
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v339 = F_get_attstatsslot(m, v17+int32(140), v331, int32(7), int32(0), int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	if v339 == int32(0) {
		goto L92
	} else {
		goto L129
	}
L129:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	if v343 < int32(2) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_free_attstatsslot(m, v17+int32(140))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L2
	} else {
		goto L190
	}
L131:
	;
	v348 = v343 << (uint(int32(3)) % 32)
	v349 = F_palloc(m, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L132
	}
L132:
	;
	v351 = F_palloc(m, v348)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L133
	}
L133:
	;
	v353 = int32(0)
	goto L134
L134:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367+v353<<(uint(int32(2))%32))))
	v372 = F_pg_detoast_datum(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L136
	}
L135:
	;
	v389 = v164 - int32(2870)
	if base.Ui32(int32(4)) < base.Ui32(v389) {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v375 = v353 << (uint(int32(3)) % 32)
	F_range_deserialize(m, v318, v372, v349+v375, v375+v351, v17+int32(79))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+79)))
	if v382 == int32(1) {
		goto L95
	} else {
		goto L138
	}
L138:
	;
	v386 = v353 + int32(1)
	if v386 != v343 {
		v353 = v386
		goto L134
	} else {
		goto L139
	}
L139:
	;
	goto L135
L140:
	;
	F_multirange_get_bounds(m, v318, v262, int32(0), v17+int32(96), v17+int32(80))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L148
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+136)) = int32(0)
	v412 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = v412
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v412
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v412
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v412
	goto L140
L142:
	;
	if v389 == int32(2) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v394 == int32(0) {
		goto L130
	} else {
		goto L144
	}
L144:
	;
	v402 = F_get_attstatsslot(m, v17+int32(104), v394, int32(6), int32(0), int32(1))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	if v402 == int32(0) {
		goto L130
	} else {
		goto L146
	}
L146:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v17)+120))
	if int32(2) <= v407 {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v607 = float64(-1)
	goto L93
L148:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	F_multirange_get_bounds(m, v318, v262, v429-int32(1), v17+int32(80), v17+int32(88))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	if v164 <= int32(4141) {
		goto L163
	} else {
		goto L164
	}
L150:
	;
	v553 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v349, v343, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L2
	} else {
		goto L189
	}
L151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L2
	} else {
		goto L186
	}
L152:
	;
	if v164 == int32(4035) {
		goto L94
	} else {
		goto L185
	}
L153:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+92)))
	if v515 == int32(1) {
		goto L180
	} else {
		goto L181
	}
L154:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v17)+116))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v17)+120))
	v513 = F_calc_hist_selectivity_contains(m, v318, v17+int32(96), v17+int32(88), v349, v343, v511, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L2
	} else {
		goto L179
	}
L155:
	;
	v496 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v351, v343, int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L2
	} else {
		goto L177
	}
L156:
	;
	v490 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(88), v351, v343, int32(1))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L176
	}
L157:
	;
	v484 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(88), v349, v343, int32(1))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L2
	} else {
		goto L175
	}
L158:
	;
	v478 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v351, v343, int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L174
	}
L159:
	;
	v472 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v349, v343, int32(1))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L2
	} else {
		goto L173
	}
L160:
	;
	v465 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v349, v343, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L2
	} else {
		goto L172
	}
L161:
	;
	v459 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v349, v343, int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L2
	} else {
		goto L171
	}
L162:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+100)))
	if v448 != int32(1) {
		goto L153
	} else {
		goto L169
	}
L163:
	;
	switch v164 - int32(2862) {
	case 0:
		goto L150
	case 1:
		goto L161
	case 2:
		goto L159
	case 3:
		goto L160
	case 4, 10, 11, 13:
		goto L151
	case 5, 6, 7:
		goto L155
	case 8, 9:
		goto L154
	case 12:
		goto L162
	case 14, 15:
		goto L156
	default:
		goto L152
	}
L164:
	;
	goto L165
L165:
	;
	switch v164 - int32(4396) {
	case 0, 1:
		goto L158
	case 2:
		goto L151
	case 3, 4:
		goto L157
	default:
		goto L166
	}
L166:
	;
	if v164 == int32(4142) {
		goto L94
	} else {
		goto L167
	}
L167:
	;
	if v164 != int32(4540) {
		goto L151
	} else {
		goto L168
	}
L168:
	;
	goto L162
L169:
	;
	v454 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(88), v351, v343, int32(1))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L2
	} else {
		goto L170
	}
L170:
	;
	v607 = v454
	goto L93
L171:
	;
	v607 = v459
	goto L93
L172:
	;
	v607 = base.F64_sub(float64(1), v465)
	goto L93
L173:
	;
	v607 = base.F64_sub(float64(1), v472)
	goto L93
L174:
	;
	v607 = v478
	goto L93
L175:
	;
	v607 = base.F64_sub(float64(1), v484)
	goto L93
L176:
	;
	v607 = v490
	goto L93
L177:
	;
	v502 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(88), v349, v343, int32(1))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L2
	} else {
		goto L178
	}
L178:
	;
	v607 = base.F64_sub(float64(1), base.F64_add(v496, base.F64_sub(float64(1), v502)))
	goto L93
L179:
	;
	v607 = v513
	goto L93
L180:
	;
	v522 = F_calc_hist_selectivity_scalar(m, v318, v17+int32(96), v349, v343, int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L2
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v17)+116))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v17)+120))
	v531 = F_calc_hist_selectivity_contained(m, v318, v17+int32(96), v17+int32(88), v349, v343, v529, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L2
	} else {
		goto L184
	}
L183:
	;
	v607 = base.F64_sub(float64(1), v522)
	goto L93
L184:
	;
	v607 = v531
	goto L93
L185:
	;
	goto L151
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v164
	F_errmsg_internal(m, int32(43725), v17+int32(16))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(496212), int32(690), int32(9862))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v607 = v553
	goto L93
L190:
	;
	goto L92
L191:
	;
	F_errmsg_internal(m, int32(491810), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L2
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(496212), int32(318), int32(308164))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L2
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errmsg_internal(m, int32(402980), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(496212), int32(509), int32(9862))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L2
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	v607 = base.F64_sub(float64(1), v603)
	goto L93
L198:
	;
	F_free_attstatsslot(m, v17+int32(140))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L199
	}
L199:
	;
	if base.F64_lt(v607, float64(0)) == int32(0) {
		v693 = v607
		goto L91
	} else {
		goto L200
	}
L200:
	;
	goto L92
L201:
	;
	v693 = v683
	goto L91
L202:
	;
	v683 = v677
	goto L201
L203:
	;
	v677 = float64(0.3333333333333333)
	goto L202
L204:
	;
	if v164 == int32(3585) {
		goto L203
	} else {
		goto L216
	}
L205:
	;
	v683 = float64(0.005)
	goto L201
L206:
	;
	v641 = v164 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v641) {
		goto L204
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	if base.Ui32(v164-int32(4395)) < base.Ui32(int32(6)) {
		goto L203
	} else {
		goto L213
	}
L209:
	;
	v645 = int32(1) << (uint(v641) % 32)
	if v645&int32(57359) != 0 {
		goto L203
	} else {
		goto L210
	}
L210:
	;
	if v645&int32(6912) != 0 {
		goto L205
	} else {
		goto L211
	}
L211:
	;
	if int32(1)<<(uint(v641)%32)&int32(1152) == int32(0) {
		goto L204
	} else {
		goto L212
	}
L212:
	;
	goto L205
L213:
	;
	if base.Ui32(v164-int32(4539)) < base.Ui32(int32(2)) {
		goto L205
	} else {
		goto L214
	}
L214:
	;
	if v164 != int32(4142) {
		v677 = v637
		goto L202
	} else {
		goto L215
	}
L215:
	;
	goto L203
L216:
	;
	if v164 != int32(4035) {
		v677 = v637
		goto L202
	} else {
		goto L217
	}
L217:
	;
	goto L203
L218:
	;
	v740 = base.F64_add(base.F64_mul(base.F64_sub(float64(1), v294), v693), v294)
	goto L88
L219:
	;
	goto L220
L220:
	;
	v740 = base.F64_mul(base.F64_sub(float64(1), v294), v693)
	goto L88
L221:
	;
	if v164 == int32(4540) {
		v740 = v294
		goto L88
	} else {
		goto L222
	}
L222:
	;
	goto L89
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v164
	F_errmsg_internal(m, int32(43831), v17)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L2
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(496212), int32(402), int32(308164))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L2
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	if base.F64_gt(v748, float64(1)) == int32(0) {
		v823 = v748
		goto L59
	} else {
		goto L227
	}
L227:
	;
	v823 = float64(1)
	goto L59
L228:
	;
	v823 = v808
	goto L59
L229:
	;
	v808 = v802
	goto L228
L230:
	;
	v802 = float64(0.3333333333333333)
	goto L229
L231:
	;
	if v164 == int32(3585) {
		goto L230
	} else {
		goto L243
	}
L232:
	;
	v808 = float64(0.005)
	goto L228
L233:
	;
	v766 = v164 - int32(2862)
	if base.Ui32(int32(15)) < base.Ui32(v766) {
		goto L231
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	if base.Ui32(v164-int32(4395)) < base.Ui32(int32(6)) {
		goto L230
	} else {
		goto L240
	}
L236:
	;
	v770 = int32(1) << (uint(v766) % 32)
	if v770&int32(57359) != 0 {
		goto L230
	} else {
		goto L237
	}
L237:
	;
	if v770&int32(6912) != 0 {
		goto L232
	} else {
		goto L238
	}
L238:
	;
	if int32(1)<<(uint(v766)%32)&int32(1152) == int32(0) {
		goto L231
	} else {
		goto L239
	}
L239:
	;
	goto L232
L240:
	;
	if base.Ui32(v164-int32(4539)) < base.Ui32(int32(2)) {
		goto L232
	} else {
		goto L241
	}
L241:
	;
	if v164 != int32(4142) {
		v802 = v762
		goto L229
	} else {
		goto L242
	}
L242:
	;
	goto L230
L243:
	;
	if v164 != int32(4035) {
		v802 = v762
		goto L229
	} else {
		goto L244
	}
L244:
	;
	goto L230
L245:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	m.T0[v825].(func(*base.Module, int32))(m, v824)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L2
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v828 = float64(0)
	if base.F64_lt(v823, v828) != 0 {
		v846 = v828
		goto L1
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	if base.F64_gt(v823, float64(1)) == int32(0) {
		v846 = v823
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v846 = float64(1)
	goto L1
L251:
	;
	m.G0 = v17 + int32(176)
	return v850
}
func F_multixactoffsetssyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(4415268), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
