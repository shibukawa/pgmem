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
	F_errmsg(m, int32(214411), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(490308), int32(107), int32(63384))
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
	F_errmsg(m, int32(214411), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(490308), int32(169), int32(63384))
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
	F_errmsg(m, int32(214411), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(490308), int32(237), int32(63384))
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
			v22 = int32(4470752)
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
						v45 = int32(4470752)
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
				v45 = int32(4470752)
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
func F_make_inner_pathkeys_for_merge(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	v4 = int32(0)
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v15 = v14
	goto L3
L2:
	;
	v15 = v4
	goto L3
L3:
	;
	if l1 == int32(0) {
		v233 = v4
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v233
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v233 = v4
	goto L4
L7:
	;
	goto L8
L8:
	;
	v26 = v4
	v27 = v15
	v28 = v4
	v30 = v4
	v31 = v4
	goto L10
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L42
	} else {
		goto L58
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v30<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+100))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L42
	} else {
		goto L55
	}
L12:
	;
	v44 = v40
	goto L15
L13:
	;
	goto L14
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+56))
	if v70 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v44
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
	if v55 != 0 {
		v44 = v55
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	goto L16
L18:
	;
	v74 = v70
	goto L21
L19:
	;
	goto L20
L20:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+120)))
	if v101 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v74
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+56))
	if v85 != 0 {
		v74 = v85
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	goto L22
L24:
	;
	v102 = int32(104)
	goto L26
L25:
	;
	v102 = int32(100)
	goto L26
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v38+v102)))
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L11
L28:
	;
	v107 = int32(100)
	goto L30
L29:
	;
	v107 = int32(104)
	goto L30
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v38+v107)))
	if v31 != v109 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v27 == int32(0) {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	v127 = v27
	v128 = v28
	v129 = v31
	goto L33
L33:
	;
	if v109 != v104 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v109 != v114 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v117 = v27 + int32(4)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v117) < base.Ui32(v119+v120<<(uint(int32(2))%32)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = v117
	goto L38
L37:
	;
	v125 = int32(0)
	goto L38
L38:
	;
	v127 = v125
	v128 = v113
	v129 = v109
	goto L33
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+16)))
	v134 = F_make_canonical_pathkey(m, l0, v104, v131, v132, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v138 = v128
	goto L41
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+40)))
	if v140 != 0 {
		v190 = v26
		goto L44
	} else {
		goto L45
	}
L42:
	;
	return int32(0)
L43:
	;
	v138 = v134
	goto L41
L44:
	;
	v199 = v30 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v200 <= v199 {
		v233 = v190
		goto L4
	} else {
		goto L54
	}
L45:
	;
	if v26 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v183 = F_lappend(m, v26, v138)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L42
	} else {
		goto L53
	}
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v143 <= int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v151 = int32(0)
	goto L49
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v146+v151<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v139 == v165 {
		v190 = v26
		goto L44
	} else {
		goto L51
	}
L50:
	;
	goto L46
L51:
	;
	v168 = v151 + int32(1)
	if v143 != v168 {
		v151 = v168
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v190 = v183
	goto L44
L54:
	;
	v26 = v190
	v27 = v127
	v28 = v128
	v30 = v199
	v31 = v129
	goto L10
L55:
	;
	F_errmsg_internal(m, int32(158495), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L42
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(485995), int32(1897), int32(394141))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L42
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errmsg_internal(m, int32(352345), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L42
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(485995), int32(1902), int32(394141))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L42
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
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
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
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
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
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
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v359 int64
	_ = v359
	var v362 int32
	_ = v362
	var v364 int64
	_ = v364
	var v366 int64
	_ = v366
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v380 int64
	_ = v380
	var v385 int32
	_ = v385
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int64
	_ = v395
	var v405 int64
	_ = v405
	var v420 float64
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
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
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int64
	_ = v1155
	var v1157 int64
	_ = v1157
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1355 int32
	_ = v1355
	var v1361 int32
	_ = v1361
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1658 int64
	_ = v1658
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int64
	_ = v1711
	var v1713 int64
	_ = v1713
	var v1732 int32
	_ = v1732
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1742 int64
	_ = v1742
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1772 int32
	_ = v1772
	var v1781 int32
	_ = v1781
	var v1792 int32
	_ = v1792
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1835 int32
	_ = v1835
	var v1842 int32
	_ = v1842
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if l1&int32(3) == int32(0) {
		v42 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v81 = v75 - int32(1636608432)
	if l1&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v75 = v67 - l1
	goto L1
L3:
	;
	v46 = v42
	goto L12
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v75 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v31 = l1
	goto L8
L8:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v67 = v35
	goto L2
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v61 = v46
	goto L15
L14:
	;
	goto L13
L15:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v67 = v61
	goto L2
L17:
	;
	goto L16
L18:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v342 = v340
	v351 = v341
	goto L58
L19:
	;
	v313 = int32(14)
	v315 = v309 ^ v310 - base.I32_rotl(v309, v313)
	v319 = v315 ^ v308 - base.I32_rotl(v315, int32(11))
	v323 = v319 ^ v309 - base.I32_rotl(v319, int32(25))
	v327 = v323 ^ v315 - base.I32_rotl(v323, int32(16))
	v331 = v327 ^ v319 - base.I32_rotl(v327, int32(4))
	v335 = v331 ^ v323 - base.I32_rotl(v331, v313)
	goto L18
L20:
	;
	switch v239 - int32(1) {
	case 0:
		v301 = v240
		v302 = v241
		v303 = v242
		goto L47
	case 1:
		v294 = v240
		v295 = v241
		v296 = v242
		goto L48
	case 2:
		v287 = v240
		v288 = v241
		v289 = v242
		goto L49
	case 3:
		v281 = v241
		v282 = v242
		goto L50
	case 4:
		v277 = v241
		v278 = v242
		goto L51
	case 5:
		v271 = v241
		v272 = v242
		goto L52
	case 6:
		v265 = v241
		v266 = v242
		goto L53
	case 7:
		v260 = v242
		goto L54
	case 8:
		v255 = v242
		goto L55
	case 9:
		v250 = v242
		goto L56
	case 10:
		goto L57
	default:
		v308 = v240
		v309 = v241
		v310 = v242
		goto L19
	}
L21:
	;
	v190 = l1
	v191 = v75
	v192 = v81
	v193 = v81
	v194 = v81
	goto L44
L22:
	;
	if base.Ui32(int32(11)) < base.Ui32(v75) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v75) < base.Ui32(int32(12)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v238 = l1
	v239 = v75
	v240 = v81
	v241 = v81
	v242 = v81
	goto L20
L26:
	;
	switch v137 - int32(1) {
	case 0:
		v187 = v138
		goto L33
	case 1:
		v182 = v138
		goto L34
	case 2:
		goto L35
	case 3:
		v175 = v139
		goto L36
	case 4:
		v172 = v139
		goto L37
	case 5:
		v167 = v139
		goto L38
	case 6:
		goto L39
	case 7:
		v158 = v140
		goto L40
	case 8:
		v153 = v140
		goto L41
	case 9:
		v148 = v140
		goto L42
	case 10:
		goto L43
	default:
		v308 = v138
		v309 = v139
		v310 = v140
		goto L19
	}
L27:
	;
	v136 = l1
	v137 = v75
	v138 = v81
	v139 = v81
	v140 = v81
	goto L26
L28:
	;
	goto L29
L29:
	;
	v88 = l1
	v89 = v75
	v90 = v81
	v91 = v81
	v92 = v81
	goto L30
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v95 = v94 + v91
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v99 = v98 + v92
	v101 = int32(4)
	v103 = v96 + v90 - v99 ^ base.I32_rotl(v99, v101)
	v107 = v95 - v103 ^ base.I32_rotl(v103, int32(6))
	v108 = v99 + v95
	v109 = v103 + v108
	v110 = v107 + v109
	v114 = v108 - v107 ^ base.I32_rotl(v107, int32(8))
	v118 = v109 - v114 ^ base.I32_rotl(v114, int32(16))
	v122 = v110 - v118 ^ base.I32_rotl(v118, int32(19))
	v123 = v114 + v110
	v124 = v118 + v123
	v125 = v122 + v124
	v129 = v123 - v122 ^ base.I32_rotl(v122, v101)
	v130 = int32(12)
	v131 = v88 + v130
	v133 = v89 - v130
	if base.Ui32(int32(11)) < base.Ui32(v133) {
		v88 = v131
		v89 = v133
		v90 = v124
		v91 = v125
		v92 = v129
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v136 = v131
	v137 = v133
	v138 = v124
	v139 = v125
	v140 = v129
	goto L26
L32:
	;
	goto L31
L33:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v308 = v187 + v188
	v309 = v139
	v310 = v140
	goto L19
L34:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v187 = v183<<(uint(int32(8))%32) + v182
	goto L33
L35:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+2)))
	v182 = v178<<(uint(int32(16))%32) + v138
	goto L34
L36:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v308 = v176 + v138
	v309 = v175
	v310 = v140
	goto L19
L37:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+4)))
	v175 = v172 + v173
	goto L36
L38:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+5)))
	v172 = v168<<(uint(int32(8))%32) + v167
	goto L37
L39:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+6)))
	v167 = v163<<(uint(int32(16))%32) + v139
	goto L38
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v308 = v159 + v138
	v309 = v161 + v139
	v310 = v158
	goto L19
L41:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+8)))
	v158 = v154<<(uint(int32(8))%32) + v153
	goto L40
L42:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+9)))
	v153 = v149<<(uint(int32(16))%32) + v148
	goto L41
L43:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+10)))
	v148 = v144<<(uint(int32(24))%32) + v140
	goto L42
L44:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	v197 = v196 + v193
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	v201 = v200 + v194
	v203 = int32(4)
	v205 = v198 + v192 - v201 ^ base.I32_rotl(v201, v203)
	v209 = v197 - v205 ^ base.I32_rotl(v205, int32(6))
	v210 = v201 + v197
	v211 = v205 + v210
	v212 = v209 + v211
	v216 = v210 - v209 ^ base.I32_rotl(v209, int32(8))
	v220 = v211 - v216 ^ base.I32_rotl(v216, int32(16))
	v224 = v212 - v220 ^ base.I32_rotl(v220, int32(19))
	v225 = v216 + v212
	v226 = v220 + v225
	v227 = v224 + v226
	v231 = v225 - v224 ^ base.I32_rotl(v224, v203)
	v232 = int32(12)
	v233 = v190 + v232
	v235 = v191 - v232
	if base.Ui32(int32(11)) < base.Ui32(v235) {
		v190 = v233
		v191 = v235
		v192 = v226
		v193 = v227
		v194 = v231
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v238 = v233
	v239 = v235
	v240 = v226
	v241 = v227
	v242 = v231
	goto L20
L46:
	;
	goto L45
L47:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v308 = v301 + v304
	v309 = v302
	v310 = v303
	goto L19
L48:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	v301 = v297<<(uint(int32(8))%32) + v294
	v302 = v295
	v303 = v296
	goto L47
L49:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+2)))
	v294 = v290<<(uint(int32(16))%32) + v287
	v295 = v288
	v296 = v289
	goto L48
L50:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+3)))
	v287 = v283<<(uint(int32(24))%32) + v240
	v288 = v281
	v289 = v282
	goto L49
L51:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+4)))
	v281 = v277 + v279
	v282 = v278
	goto L50
L52:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+5)))
	v277 = v273<<(uint(int32(8))%32) + v271
	v278 = v272
	goto L51
L53:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+6)))
	v271 = v267<<(uint(int32(16))%32) + v265
	v272 = v266
	goto L52
L54:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+7)))
	v265 = v261<<(uint(int32(24))%32) + v241
	v266 = v260
	goto L53
L55:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+8)))
	v260 = v256<<(uint(int32(8))%32) + v255
	goto L54
L56:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+9)))
	v255 = v251<<(uint(int32(16))%32) + v250
	goto L55
L57:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+10)))
	v250 = v246<<(uint(int32(24))%32) + v242
	goto L56
L58:
	;
	if base.Ui32(v342) <= base.Ui32(v351) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v1842 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1842
	v342 = v1842
	v351 = v1835
	goto L58
L61:
	;
	return
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1792)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v1792))) = int32(1)
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+24))
	v1806 = F_MemoryContextStrdup(m, v1805, l1)
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L80
	} else {
		goto L334
	}
L63:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1781 + int32(1)
	v1792 = v1772
	goto L62
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L80
	} else {
		goto L331
	}
L65:
	;
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if v359 == int64(4294967296) {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v1245 = int32(0)
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1248 = v1247 & (v335 ^ v327 - base.I32_rotl(v335, int32(24)))
	v1251 = v1246 + v1248<<(uint(int32(4))%32)
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1251)))
	if v1252 == v1245 {
		v1772 = v1251
		goto L63
	} else {
		goto L235
	}
L68:
	;
	v362 = int32(0)
	v364 = int64(2)
	v366 = v359 << (uint(int64(1)) % 64)
	if base.Ui64(v366) <= base.Ui64(v364) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L67
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L80
	} else {
		goto L232
	}
L71:
	;
	v369 = v364
	goto L73
L72:
	;
	v369 = v366
	goto L73
L73:
	;
	v370 = int64(1)
	if v369&(v369-v370) == int64(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v380 = v369
	goto L76
L75:
	;
	v380 = v370 << (uint(int64(64)-base.I64_clz(v369)) % 64)
	goto L76
L76:
	;
	if base.Ui64(v380<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v392 = F_MemoryContextAllocExtended(m, v387, base.I32_wrap_i64(v380)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L80
	} else {
		goto L229
	}
L80:
	;
	return
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v392
	v395 = int64(1)
	if v380&(v380-v395) == int64(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v405 = v380
	goto L84
L83:
	;
	v405 = v395 << (uint(int64(64)-base.I64_clz(v380)) % 64)
	goto L84
L84:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v405<<(uint(int64(4))%64)) {
		goto L70
	} else {
		goto L85
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = base.I32_wrap_i64(v405) - int32(1)
	v420 = base.F64_mul(base.F64_convert_i64_u(v405), float64(0.9))
	if base.F64_lt(v420, float64(4.294967296e+09))&base.F64_ge(v420, float64(0)) != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v405 == int64(4294967296) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v426 = base.I32_trunc_f64_u(v420)
	v428 = v426
	goto L86
L88:
	;
	goto L89
L89:
	;
	v428 = int32(0)
	goto L86
L90:
	;
	v429 = int32(-85899346)
	goto L92
L91:
	;
	v429 = v428
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v429
	if v386 != int64(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v436 = v362
	goto L97
L94:
	;
	goto L95
L95:
	;
	F_pfree(m, v385)
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L80
	} else {
		goto L228
	}
L96:
	;
	v790 = v786
	v794 = v362
	goto L159
L97:
	;
	v451 = v385 + v436<<(uint(int32(4))%32)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	if v452 != int32(1) {
		v786 = v436
		goto L96
	} else {
		goto L99
	}
L98:
	;
	v786 = int32(0)
	goto L96
L99:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	if v455&int32(3) == int32(0) {
		v479 = v455
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v518 = v512 - int32(1636608432)
	if v455&int32(3) != 0 {
		goto L121
	} else {
		goto L122
	}
L101:
	;
	v512 = v504 - v455
	goto L100
L102:
	;
	v483 = v479
	goto L111
L103:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	if v463 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v512 = int32(0)
	goto L100
L105:
	;
	goto L106
L106:
	;
	v468 = v455
	goto L107
L107:
	;
	v472 = v468 + int32(1)
	if v472&int32(3) == int32(0) {
		v479 = v472
		goto L102
	} else {
		goto L109
	}
L108:
	;
	v504 = v472
	goto L101
L109:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	if v477 != 0 {
		v468 = v472
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	v492 = int32(-2139062144)
	if (int32(16843008)-v489|v489)&v492 == v492 {
		v483 = v483 + int32(4)
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v498 = v483
	goto L114
L113:
	;
	goto L112
L114:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if v502 != 0 {
		v498 = v498 + int32(1)
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v504 = v498
	goto L101
L116:
	;
	goto L115
L117:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if (v772^v764-base.I32_rotl(v772, int32(24)))&v777 == v436 {
		v786 = v436
		goto L96
	} else {
		goto L157
	}
L118:
	;
	v750 = int32(14)
	v752 = v746 ^ v747 - base.I32_rotl(v746, v750)
	v756 = v752 ^ v745 - base.I32_rotl(v752, int32(11))
	v760 = v756 ^ v746 - base.I32_rotl(v756, int32(25))
	v764 = v760 ^ v752 - base.I32_rotl(v760, int32(16))
	v768 = v764 ^ v756 - base.I32_rotl(v764, int32(4))
	v772 = v768 ^ v760 - base.I32_rotl(v768, v750)
	goto L117
L119:
	;
	switch v676 - int32(1) {
	case 0:
		v738 = v677
		v739 = v678
		v740 = v679
		goto L146
	case 1:
		v731 = v677
		v732 = v678
		v733 = v679
		goto L147
	case 2:
		v724 = v677
		v725 = v678
		v726 = v679
		goto L148
	case 3:
		v718 = v678
		v719 = v679
		goto L149
	case 4:
		v714 = v678
		v715 = v679
		goto L150
	case 5:
		v708 = v678
		v709 = v679
		goto L151
	case 6:
		v702 = v678
		v703 = v679
		goto L152
	case 7:
		v697 = v679
		goto L153
	case 8:
		v692 = v679
		goto L154
	case 9:
		v687 = v679
		goto L155
	case 10:
		goto L156
	default:
		v745 = v677
		v746 = v678
		v747 = v679
		goto L118
	}
L120:
	;
	v627 = v455
	v628 = v512
	v629 = v518
	v630 = v518
	v631 = v518
	goto L143
L121:
	;
	if base.Ui32(int32(11)) < base.Ui32(v512) {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if base.Ui32(v512) < base.Ui32(int32(12)) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v675 = v455
	v676 = v512
	v677 = v518
	v678 = v518
	v679 = v518
	goto L119
L125:
	;
	switch v574 - int32(1) {
	case 0:
		v624 = v575
		goto L132
	case 1:
		v619 = v575
		goto L133
	case 2:
		goto L134
	case 3:
		v612 = v576
		goto L135
	case 4:
		v609 = v576
		goto L136
	case 5:
		v604 = v576
		goto L137
	case 6:
		goto L138
	case 7:
		v595 = v577
		goto L139
	case 8:
		v590 = v577
		goto L140
	case 9:
		v585 = v577
		goto L141
	case 10:
		goto L142
	default:
		v745 = v575
		v746 = v576
		v747 = v577
		goto L118
	}
L126:
	;
	v573 = v455
	v574 = v512
	v575 = v518
	v576 = v518
	v577 = v518
	goto L125
L127:
	;
	goto L128
L128:
	;
	v525 = v455
	v526 = v512
	v527 = v518
	v528 = v518
	v529 = v518
	goto L129
L129:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v532 = v531 + v528
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	v536 = v535 + v529
	v538 = int32(4)
	v540 = v533 + v527 - v536 ^ base.I32_rotl(v536, v538)
	v544 = v532 - v540 ^ base.I32_rotl(v540, int32(6))
	v545 = v536 + v532
	v546 = v540 + v545
	v547 = v544 + v546
	v551 = v545 - v544 ^ base.I32_rotl(v544, int32(8))
	v555 = v546 - v551 ^ base.I32_rotl(v551, int32(16))
	v559 = v547 - v555 ^ base.I32_rotl(v555, int32(19))
	v560 = v551 + v547
	v561 = v555 + v560
	v562 = v559 + v561
	v566 = v560 - v559 ^ base.I32_rotl(v559, v538)
	v567 = int32(12)
	v568 = v525 + v567
	v570 = v526 - v567
	if base.Ui32(int32(11)) < base.Ui32(v570) {
		v525 = v568
		v526 = v570
		v527 = v561
		v528 = v562
		v529 = v566
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v573 = v568
	v574 = v570
	v575 = v561
	v576 = v562
	v577 = v566
	goto L125
L131:
	;
	goto L130
L132:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
	v745 = v624 + v625
	v746 = v576
	v747 = v577
	goto L118
L133:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+1)))
	v624 = v620<<(uint(int32(8))%32) + v619
	goto L132
L134:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+2)))
	v619 = v615<<(uint(int32(16))%32) + v575
	goto L133
L135:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v745 = v613 + v575
	v746 = v612
	v747 = v577
	goto L118
L136:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+4)))
	v612 = v609 + v610
	goto L135
L137:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+5)))
	v609 = v605<<(uint(int32(8))%32) + v604
	goto L136
L138:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+6)))
	v604 = v600<<(uint(int32(16))%32) + v576
	goto L137
L139:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	v745 = v596 + v575
	v746 = v598 + v576
	v747 = v595
	goto L118
L140:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+8)))
	v595 = v591<<(uint(int32(8))%32) + v590
	goto L139
L141:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+9)))
	v590 = v586<<(uint(int32(16))%32) + v585
	goto L140
L142:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+10)))
	v585 = v581<<(uint(int32(24))%32) + v577
	goto L141
L143:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v627)+4))
	v634 = v633 + v630
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v627)+8))
	v638 = v637 + v631
	v640 = int32(4)
	v642 = v635 + v629 - v638 ^ base.I32_rotl(v638, v640)
	v646 = v634 - v642 ^ base.I32_rotl(v642, int32(6))
	v647 = v638 + v634
	v648 = v642 + v647
	v649 = v646 + v648
	v653 = v647 - v646 ^ base.I32_rotl(v646, int32(8))
	v657 = v648 - v653 ^ base.I32_rotl(v653, int32(16))
	v661 = v649 - v657 ^ base.I32_rotl(v657, int32(19))
	v662 = v653 + v649
	v663 = v657 + v662
	v664 = v661 + v663
	v668 = v662 - v661 ^ base.I32_rotl(v661, v640)
	v669 = int32(12)
	v670 = v627 + v669
	v672 = v628 - v669
	if base.Ui32(int32(11)) < base.Ui32(v672) {
		v627 = v670
		v628 = v672
		v629 = v663
		v630 = v664
		v631 = v668
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v675 = v670
	v676 = v672
	v677 = v663
	v678 = v664
	v679 = v668
	goto L119
L145:
	;
	goto L144
L146:
	;
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	v745 = v738 + v741
	v746 = v739
	v747 = v740
	goto L118
L147:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+1)))
	v738 = v734<<(uint(int32(8))%32) + v731
	v739 = v732
	v740 = v733
	goto L146
L148:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+2)))
	v731 = v727<<(uint(int32(16))%32) + v724
	v732 = v725
	v733 = v726
	goto L147
L149:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+3)))
	v724 = v720<<(uint(int32(24))%32) + v677
	v725 = v718
	v726 = v719
	goto L148
L150:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+4)))
	v718 = v714 + v716
	v719 = v715
	goto L149
L151:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+5)))
	v714 = v710<<(uint(int32(8))%32) + v708
	v715 = v709
	goto L150
L152:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+6)))
	v708 = v704<<(uint(int32(16))%32) + v702
	v709 = v703
	goto L151
L153:
	;
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+7)))
	v702 = v698<<(uint(int32(24))%32) + v678
	v703 = v697
	goto L152
L154:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+8)))
	v697 = v693<<(uint(int32(8))%32) + v692
	goto L153
L155:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+9)))
	v692 = v688<<(uint(int32(16))%32) + v687
	goto L154
L156:
	;
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+10)))
	v687 = v683<<(uint(int32(24))%32) + v679
	goto L155
L157:
	;
	v781 = v436 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v781)) < base.Ui64(v386) {
		v436 = v781
		goto L97
	} else {
		goto L158
	}
L158:
	;
	goto L98
L159:
	;
	v805 = v385 + v790<<(uint(int32(4))%32)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	if v806 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L95
L161:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	if v809&int32(3) == int32(0) {
		v833 = v809
		goto L166
	} else {
		goto L167
	}
L162:
	;
	goto L163
L163:
	;
	v1176 = v790 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1176)) < base.Ui64(v386) {
		goto L224
	} else {
		goto L225
	}
L164:
	;
	v872 = v866 - int32(1636608432)
	if v809&int32(3) != 0 {
		goto L185
	} else {
		goto L186
	}
L165:
	;
	v866 = v858 - v809
	goto L164
L166:
	;
	v837 = v833
	goto L175
L167:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809))))
	if v817 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v866 = int32(0)
	goto L164
L169:
	;
	goto L170
L170:
	;
	v822 = v809
	goto L171
L171:
	;
	v826 = v822 + int32(1)
	if v826&int32(3) == int32(0) {
		v833 = v826
		goto L166
	} else {
		goto L173
	}
L172:
	;
	v858 = v826
	goto L165
L173:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	if v831 != 0 {
		v822 = v826
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v846 = int32(-2139062144)
	if (int32(16843008)-v843|v843)&v846 == v846 {
		v837 = v837 + int32(4)
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v852 = v837
	goto L178
L177:
	;
	goto L176
L178:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	if v856 != 0 {
		v852 = v852 + int32(1)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v858 = v852
	goto L165
L180:
	;
	goto L179
L181:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1132 = v1126 ^ v1118 - base.I32_rotl(v1126, int32(24))
	goto L221
L182:
	;
	v1104 = int32(14)
	v1106 = v1100 ^ v1101 - base.I32_rotl(v1100, v1104)
	v1110 = v1106 ^ v1099 - base.I32_rotl(v1106, int32(11))
	v1114 = v1110 ^ v1100 - base.I32_rotl(v1110, int32(25))
	v1118 = v1114 ^ v1106 - base.I32_rotl(v1114, int32(16))
	v1122 = v1118 ^ v1110 - base.I32_rotl(v1118, int32(4))
	v1126 = v1122 ^ v1114 - base.I32_rotl(v1122, v1104)
	goto L181
L183:
	;
	switch v1030 - int32(1) {
	case 0:
		v1092 = v1031
		v1093 = v1032
		v1094 = v1033
		goto L210
	case 1:
		v1085 = v1031
		v1086 = v1032
		v1087 = v1033
		goto L211
	case 2:
		v1078 = v1031
		v1079 = v1032
		v1080 = v1033
		goto L212
	case 3:
		v1072 = v1032
		v1073 = v1033
		goto L213
	case 4:
		v1068 = v1032
		v1069 = v1033
		goto L214
	case 5:
		v1062 = v1032
		v1063 = v1033
		goto L215
	case 6:
		v1056 = v1032
		v1057 = v1033
		goto L216
	case 7:
		v1051 = v1033
		goto L217
	case 8:
		v1046 = v1033
		goto L218
	case 9:
		v1041 = v1033
		goto L219
	case 10:
		goto L220
	default:
		v1099 = v1031
		v1100 = v1032
		v1101 = v1033
		goto L182
	}
L184:
	;
	v981 = v809
	v982 = v866
	v983 = v872
	v984 = v872
	v985 = v872
	goto L207
L185:
	;
	if base.Ui32(int32(11)) < base.Ui32(v866) {
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	if base.Ui32(v866) < base.Ui32(int32(12)) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	v1029 = v809
	v1030 = v866
	v1031 = v872
	v1032 = v872
	v1033 = v872
	goto L183
L189:
	;
	switch v928 - int32(1) {
	case 0:
		v978 = v929
		goto L196
	case 1:
		v973 = v929
		goto L197
	case 2:
		goto L198
	case 3:
		v966 = v930
		goto L199
	case 4:
		v963 = v930
		goto L200
	case 5:
		v958 = v930
		goto L201
	case 6:
		goto L202
	case 7:
		v949 = v931
		goto L203
	case 8:
		v944 = v931
		goto L204
	case 9:
		v939 = v931
		goto L205
	case 10:
		goto L206
	default:
		v1099 = v929
		v1100 = v930
		v1101 = v931
		goto L182
	}
L190:
	;
	v927 = v809
	v928 = v866
	v929 = v872
	v930 = v872
	v931 = v872
	goto L189
L191:
	;
	goto L192
L192:
	;
	v879 = v809
	v880 = v866
	v881 = v872
	v882 = v872
	v883 = v872
	goto L193
L193:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v879)+4))
	v886 = v885 + v882
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v879)+8))
	v890 = v889 + v883
	v892 = int32(4)
	v894 = v887 + v881 - v890 ^ base.I32_rotl(v890, v892)
	v898 = v886 - v894 ^ base.I32_rotl(v894, int32(6))
	v899 = v890 + v886
	v900 = v894 + v899
	v901 = v898 + v900
	v905 = v899 - v898 ^ base.I32_rotl(v898, int32(8))
	v909 = v900 - v905 ^ base.I32_rotl(v905, int32(16))
	v913 = v901 - v909 ^ base.I32_rotl(v909, int32(19))
	v914 = v905 + v901
	v915 = v909 + v914
	v916 = v913 + v915
	v920 = v914 - v913 ^ base.I32_rotl(v913, v892)
	v921 = int32(12)
	v922 = v879 + v921
	v924 = v880 - v921
	if base.Ui32(int32(11)) < base.Ui32(v924) {
		v879 = v922
		v880 = v924
		v881 = v915
		v882 = v916
		v883 = v920
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v927 = v922
	v928 = v924
	v929 = v915
	v930 = v916
	v931 = v920
	goto L189
L195:
	;
	goto L194
L196:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927))))
	v1099 = v978 + v979
	v1100 = v930
	v1101 = v931
	goto L182
L197:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+1)))
	v978 = v974<<(uint(int32(8))%32) + v973
	goto L196
L198:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+2)))
	v973 = v969<<(uint(int32(16))%32) + v929
	goto L197
L199:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v927)))
	v1099 = v967 + v929
	v1100 = v966
	v1101 = v931
	goto L182
L200:
	;
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+4)))
	v966 = v963 + v964
	goto L199
L201:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+5)))
	v963 = v959<<(uint(int32(8))%32) + v958
	goto L200
L202:
	;
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+6)))
	v958 = v954<<(uint(int32(16))%32) + v930
	goto L201
L203:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v927)))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v927)+4))
	v1099 = v950 + v929
	v1100 = v952 + v930
	v1101 = v949
	goto L182
L204:
	;
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+8)))
	v949 = v945<<(uint(int32(8))%32) + v944
	goto L203
L205:
	;
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+9)))
	v944 = v940<<(uint(int32(16))%32) + v939
	goto L204
L206:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+10)))
	v939 = v935<<(uint(int32(24))%32) + v931
	goto L205
L207:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	v988 = v987 + v984
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v981)))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v981)+8))
	v992 = v991 + v985
	v994 = int32(4)
	v996 = v989 + v983 - v992 ^ base.I32_rotl(v992, v994)
	v1000 = v988 - v996 ^ base.I32_rotl(v996, int32(6))
	v1001 = v992 + v988
	v1002 = v996 + v1001
	v1003 = v1000 + v1002
	v1007 = v1001 - v1000 ^ base.I32_rotl(v1000, int32(8))
	v1011 = v1002 - v1007 ^ base.I32_rotl(v1007, int32(16))
	v1015 = v1003 - v1011 ^ base.I32_rotl(v1011, int32(19))
	v1016 = v1007 + v1003
	v1017 = v1011 + v1016
	v1018 = v1015 + v1017
	v1022 = v1016 - v1015 ^ base.I32_rotl(v1015, v994)
	v1023 = int32(12)
	v1024 = v981 + v1023
	v1026 = v982 - v1023
	if base.Ui32(int32(11)) < base.Ui32(v1026) {
		v981 = v1024
		v982 = v1026
		v983 = v1017
		v984 = v1018
		v985 = v1022
		goto L207
	} else {
		goto L209
	}
L208:
	;
	v1029 = v1024
	v1030 = v1026
	v1031 = v1017
	v1032 = v1018
	v1033 = v1022
	goto L183
L209:
	;
	goto L208
L210:
	;
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029))))
	v1099 = v1092 + v1095
	v1100 = v1093
	v1101 = v1094
	goto L182
L211:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+1)))
	v1092 = v1088<<(uint(int32(8))%32) + v1085
	v1093 = v1086
	v1094 = v1087
	goto L210
L212:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+2)))
	v1085 = v1081<<(uint(int32(16))%32) + v1078
	v1086 = v1079
	v1087 = v1080
	goto L211
L213:
	;
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+3)))
	v1078 = v1074<<(uint(int32(24))%32) + v1031
	v1079 = v1072
	v1080 = v1073
	goto L212
L214:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+4)))
	v1072 = v1068 + v1070
	v1073 = v1069
	goto L213
L215:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+5)))
	v1068 = v1064<<(uint(int32(8))%32) + v1062
	v1069 = v1063
	goto L214
L216:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+6)))
	v1062 = v1058<<(uint(int32(16))%32) + v1056
	v1063 = v1057
	goto L215
L217:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+7)))
	v1056 = v1052<<(uint(int32(24))%32) + v1032
	v1057 = v1051
	goto L216
L218:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+8)))
	v1051 = v1047<<(uint(int32(8))%32) + v1046
	goto L217
L219:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+9)))
	v1046 = v1042<<(uint(int32(16))%32) + v1041
	goto L218
L220:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+10)))
	v1041 = v1037<<(uint(int32(24))%32) + v1033
	goto L219
L221:
	;
	v1148 = v1132 & v1131
	v1153 = v392 + v1148<<(uint(int32(4))%32)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1153)))
	if v1154 != 0 {
		v1132 = v1148 + int32(1)
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v1155 = *(*int64)(unsafe.Add(mBase, uint32(v805)))
	*(*int64)(unsafe.Add(mBase, uint32(v1153))) = v1155
	v1157 = *(*int64)(unsafe.Add(mBase, uint32(v805)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1153)+8)) = v1157
	goto L163
L223:
	;
	goto L222
L224:
	;
	v1180 = v1176
	goto L226
L225:
	;
	v1180 = int32(0)
	goto L226
L226:
	;
	v1182 = v794 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1182)) < base.Ui64(v386) {
		v790 = v1180
		v794 = v1182
		goto L159
	} else {
		goto L227
	}
L227:
	;
	goto L160
L228:
	;
	goto L69
L229:
	;
	F_errmsg_internal(m, int32(394695), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L80
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(321768), int32(327), int32(335935))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L80
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errmsg_internal(m, int32(394695), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L80
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(321768), int32(327), int32(335935))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L80
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	v1258 = v1248
	v1259 = v1245
	v1262 = v1251
	goto L236
L236:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1262)+4))
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271))))
	if v1275 == int32(0) {
		v1294 = v1274
		v1295 = v1275
		goto L239
	} else {
		goto L240
	}
L237:
	;
	v1772 = v1750
	goto L63
L238:
	;
	if v1295-v1294 == int32(0) {
		goto L61
	} else {
		goto L246
	}
L239:
	;
	goto L238
L240:
	;
	if v1274 != v1275 {
		v1294 = v1274
		v1295 = v1275
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1279 = v1271
	v1280 = l1
	goto L242
L242:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280)+1)))
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279)+1)))
	if v1284 == int32(0) {
		v1294 = v1283
		v1295 = v1284
		goto L239
	} else {
		goto L244
	}
L243:
	;
	v1294 = v1283
	v1295 = v1284
	goto L239
L244:
	;
	v1287 = int32(1)
	if v1283 == v1284 {
		v1279 = v1279 + v1287
		v1280 = v1280 + v1287
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	if v1271&int32(3) == int32(0) {
		v1322 = v1271
		goto L249
	} else {
		goto L250
	}
L247:
	;
	v1361 = v1355 - int32(1636608432)
	if v1271&int32(3) != 0 {
		goto L268
	} else {
		goto L269
	}
L248:
	;
	v1355 = v1347 - v1271
	goto L247
L249:
	;
	v1326 = v1322
	goto L258
L250:
	;
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271))))
	if v1306 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1355 = int32(0)
	goto L247
L252:
	;
	goto L253
L253:
	;
	v1311 = v1271
	goto L254
L254:
	;
	v1315 = v1311 + int32(1)
	if v1315&int32(3) == int32(0) {
		v1322 = v1315
		goto L249
	} else {
		goto L256
	}
L255:
	;
	v1347 = v1315
	goto L248
L256:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315))))
	if v1320 != 0 {
		v1311 = v1315
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	v1335 = int32(-2139062144)
	if (int32(16843008)-v1332|v1332)&v1335 == v1335 {
		v1326 = v1326 + int32(4)
		goto L258
	} else {
		goto L260
	}
L259:
	;
	v1341 = v1326
	goto L261
L260:
	;
	goto L259
L261:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1341))))
	if v1345 != 0 {
		v1341 = v1341 + int32(1)
		goto L261
	} else {
		goto L263
	}
L262:
	;
	v1347 = v1341
	goto L248
L263:
	;
	goto L262
L264:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1621 = (v1615 ^ v1607 - base.I32_rotl(v1615, int32(24))) & v1620
	if base.Ui32(v1258) < base.Ui32(v1621) {
		goto L304
	} else {
		goto L305
	}
L265:
	;
	v1593 = int32(14)
	v1595 = v1589 ^ v1590 - base.I32_rotl(v1589, v1593)
	v1599 = v1595 ^ v1588 - base.I32_rotl(v1595, int32(11))
	v1603 = v1599 ^ v1589 - base.I32_rotl(v1599, int32(25))
	v1607 = v1603 ^ v1595 - base.I32_rotl(v1603, int32(16))
	v1611 = v1607 ^ v1599 - base.I32_rotl(v1607, int32(4))
	v1615 = v1611 ^ v1603 - base.I32_rotl(v1611, v1593)
	goto L264
L266:
	;
	switch v1519 - int32(1) {
	case 0:
		v1581 = v1520
		v1582 = v1521
		v1583 = v1522
		goto L293
	case 1:
		v1574 = v1520
		v1575 = v1521
		v1576 = v1522
		goto L294
	case 2:
		v1567 = v1520
		v1568 = v1521
		v1569 = v1522
		goto L295
	case 3:
		v1561 = v1521
		v1562 = v1522
		goto L296
	case 4:
		v1557 = v1521
		v1558 = v1522
		goto L297
	case 5:
		v1551 = v1521
		v1552 = v1522
		goto L298
	case 6:
		v1545 = v1521
		v1546 = v1522
		goto L299
	case 7:
		v1540 = v1522
		goto L300
	case 8:
		v1535 = v1522
		goto L301
	case 9:
		v1530 = v1522
		goto L302
	case 10:
		goto L303
	default:
		v1588 = v1520
		v1589 = v1521
		v1590 = v1522
		goto L265
	}
L267:
	;
	v1470 = v1271
	v1471 = v1355
	v1472 = v1361
	v1473 = v1361
	v1474 = v1361
	goto L290
L268:
	;
	if base.Ui32(int32(11)) < base.Ui32(v1355) {
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	if base.Ui32(v1355) < base.Ui32(int32(12)) {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	v1518 = v1271
	v1519 = v1355
	v1520 = v1361
	v1521 = v1361
	v1522 = v1361
	goto L266
L272:
	;
	switch v1417 - int32(1) {
	case 0:
		v1467 = v1418
		goto L279
	case 1:
		v1462 = v1418
		goto L280
	case 2:
		goto L281
	case 3:
		v1455 = v1419
		goto L282
	case 4:
		v1452 = v1419
		goto L283
	case 5:
		v1447 = v1419
		goto L284
	case 6:
		goto L285
	case 7:
		v1438 = v1420
		goto L286
	case 8:
		v1433 = v1420
		goto L287
	case 9:
		v1428 = v1420
		goto L288
	case 10:
		goto L289
	default:
		v1588 = v1418
		v1589 = v1419
		v1590 = v1420
		goto L265
	}
L273:
	;
	v1416 = v1271
	v1417 = v1355
	v1418 = v1361
	v1419 = v1361
	v1420 = v1361
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1368 = v1271
	v1369 = v1355
	v1370 = v1361
	v1371 = v1361
	v1372 = v1361
	goto L276
L276:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+4))
	v1375 = v1374 + v1371
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1368)))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+8))
	v1379 = v1378 + v1372
	v1381 = int32(4)
	v1383 = v1376 + v1370 - v1379 ^ base.I32_rotl(v1379, v1381)
	v1387 = v1375 - v1383 ^ base.I32_rotl(v1383, int32(6))
	v1388 = v1379 + v1375
	v1389 = v1383 + v1388
	v1390 = v1387 + v1389
	v1394 = v1388 - v1387 ^ base.I32_rotl(v1387, int32(8))
	v1398 = v1389 - v1394 ^ base.I32_rotl(v1394, int32(16))
	v1402 = v1390 - v1398 ^ base.I32_rotl(v1398, int32(19))
	v1403 = v1394 + v1390
	v1404 = v1398 + v1403
	v1405 = v1402 + v1404
	v1409 = v1403 - v1402 ^ base.I32_rotl(v1402, v1381)
	v1410 = int32(12)
	v1411 = v1368 + v1410
	v1413 = v1369 - v1410
	if base.Ui32(int32(11)) < base.Ui32(v1413) {
		v1368 = v1411
		v1369 = v1413
		v1370 = v1404
		v1371 = v1405
		v1372 = v1409
		goto L276
	} else {
		goto L278
	}
L277:
	;
	v1416 = v1411
	v1417 = v1413
	v1418 = v1404
	v1419 = v1405
	v1420 = v1409
	goto L272
L278:
	;
	goto L277
L279:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416))))
	v1588 = v1467 + v1468
	v1589 = v1419
	v1590 = v1420
	goto L265
L280:
	;
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+1)))
	v1467 = v1463<<(uint(int32(8))%32) + v1462
	goto L279
L281:
	;
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+2)))
	v1462 = v1458<<(uint(int32(16))%32) + v1418
	goto L280
L282:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1416)))
	v1588 = v1456 + v1418
	v1589 = v1455
	v1590 = v1420
	goto L265
L283:
	;
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+4)))
	v1455 = v1452 + v1453
	goto L282
L284:
	;
	v1448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+5)))
	v1452 = v1448<<(uint(int32(8))%32) + v1447
	goto L283
L285:
	;
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+6)))
	v1447 = v1443<<(uint(int32(16))%32) + v1419
	goto L284
L286:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1416)))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+4))
	v1588 = v1439 + v1418
	v1589 = v1441 + v1419
	v1590 = v1438
	goto L265
L287:
	;
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+8)))
	v1438 = v1434<<(uint(int32(8))%32) + v1433
	goto L286
L288:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+9)))
	v1433 = v1429<<(uint(int32(16))%32) + v1428
	goto L287
L289:
	;
	v1424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+10)))
	v1428 = v1424<<(uint(int32(24))%32) + v1420
	goto L288
L290:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+4))
	v1477 = v1476 + v1473
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1470)))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+8))
	v1481 = v1480 + v1474
	v1483 = int32(4)
	v1485 = v1478 + v1472 - v1481 ^ base.I32_rotl(v1481, v1483)
	v1489 = v1477 - v1485 ^ base.I32_rotl(v1485, int32(6))
	v1490 = v1481 + v1477
	v1491 = v1485 + v1490
	v1492 = v1489 + v1491
	v1496 = v1490 - v1489 ^ base.I32_rotl(v1489, int32(8))
	v1500 = v1491 - v1496 ^ base.I32_rotl(v1496, int32(16))
	v1504 = v1492 - v1500 ^ base.I32_rotl(v1500, int32(19))
	v1505 = v1496 + v1492
	v1506 = v1500 + v1505
	v1507 = v1504 + v1506
	v1511 = v1505 - v1504 ^ base.I32_rotl(v1504, v1483)
	v1512 = int32(12)
	v1513 = v1470 + v1512
	v1515 = v1471 - v1512
	if base.Ui32(int32(11)) < base.Ui32(v1515) {
		v1470 = v1513
		v1471 = v1515
		v1472 = v1506
		v1473 = v1507
		v1474 = v1511
		goto L290
	} else {
		goto L292
	}
L291:
	;
	v1518 = v1513
	v1519 = v1515
	v1520 = v1506
	v1521 = v1507
	v1522 = v1511
	goto L266
L292:
	;
	goto L291
L293:
	;
	v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518))))
	v1588 = v1581 + v1584
	v1589 = v1582
	v1590 = v1583
	goto L265
L294:
	;
	v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+1)))
	v1581 = v1577<<(uint(int32(8))%32) + v1574
	v1582 = v1575
	v1583 = v1576
	goto L293
L295:
	;
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+2)))
	v1574 = v1570<<(uint(int32(16))%32) + v1567
	v1575 = v1568
	v1576 = v1569
	goto L294
L296:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+3)))
	v1567 = v1563<<(uint(int32(24))%32) + v1520
	v1568 = v1561
	v1569 = v1562
	goto L295
L297:
	;
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+4)))
	v1561 = v1557 + v1559
	v1562 = v1558
	goto L296
L298:
	;
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+5)))
	v1557 = v1553<<(uint(int32(8))%32) + v1551
	v1558 = v1552
	goto L297
L299:
	;
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+6)))
	v1551 = v1547<<(uint(int32(16))%32) + v1545
	v1552 = v1546
	goto L298
L300:
	;
	v1541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+7)))
	v1545 = v1541<<(uint(int32(24))%32) + v1521
	v1546 = v1540
	goto L299
L301:
	;
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+8)))
	v1540 = v1536<<(uint(int32(8))%32) + v1535
	goto L300
L302:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+9)))
	v1535 = v1531<<(uint(int32(16))%32) + v1530
	goto L301
L303:
	;
	v1526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518)+10)))
	v1530 = v1526<<(uint(int32(24))%32) + v1522
	goto L302
L304:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1625 = v1258 + v1623
	goto L306
L305:
	;
	v1625 = v1258
	goto L306
L306:
	;
	v1628 = v1620 & (v1258 + int32(1))
	if base.Ui32(v1625-v1621) < base.Ui32(v1259) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1634 = v1246 + v1628<<(uint(int32(4))%32)
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1634)))
	if v1635 != 0 {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	goto L309
L309:
	;
	v1737 = v1259 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1737) {
		goto L326
	} else {
		goto L327
	}
L310:
	;
	v1636 = v1628
	v1641 = int32(0)
	goto L313
L311:
	;
	v1671 = v1628
	v1675 = v1634
	goto L312
L312:
	;
	if v1671 != v1258 {
		goto L320
	} else {
		goto L321
	}
L313:
	;
	v1653 = v1641 + int32(1)
	if int32(151) <= v1653 {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	v1671 = v1666
	v1675 = v1669
	goto L312
L315:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v1658 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1656), base.F64_convert_i64_u(v1658)), float64(0.1)) != 0 {
		v1835 = v1656
		goto L60
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1666 = (v1636 + int32(1)) & v1620
	v1669 = v1246 + v1666<<(uint(int32(4))%32)
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1669)))
	if v1670 != 0 {
		v1636 = v1666
		v1641 = v1653
		goto L313
	} else {
		goto L319
	}
L318:
	;
	goto L317
L319:
	;
	goto L314
L320:
	;
	v1688 = v1671
	v1692 = v1675
	goto L323
L321:
	;
	goto L322
L322:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1732 + int32(1)
	v1792 = v1262
	goto L62
L323:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1707 = v1704 & (v1688 - int32(1))
	v1710 = v1246 + v1707<<(uint(int32(4))%32)
	v1711 = *(*int64)(unsafe.Add(mBase, uint32(v1710)))
	*(*int64)(unsafe.Add(mBase, uint32(v1692))) = v1711
	v1713 = *(*int64)(unsafe.Add(mBase, uint32(v1710)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1692)+8)) = v1713
	if v1707 != v1258 {
		v1688 = v1707
		v1692 = v1710
		goto L323
	} else {
		goto L325
	}
L324:
	;
	goto L322
L325:
	;
	goto L324
L326:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v1742 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1740), base.F64_convert_i64_u(v1742)), float64(0.1)) != 0 {
		v1835 = v1740
		goto L60
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v1750 = v1246 + v1628<<(uint(int32(4))%32)
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1750)))
	if v1751 != 0 {
		v1258 = v1628
		v1259 = v1737
		v1262 = v1750
		goto L236
	} else {
		goto L330
	}
L329:
	;
	goto L328
L330:
	;
	goto L237
L331:
	;
	F_errmsg_internal(m, int32(455833), int32(0))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L80
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(321768), int32(630), int32(307124))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L80
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1792)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v1792)+4)) = v1806
	goto L61
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
	if v25 == int32(4088740) {
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
	if v41 != int32(4088740) {
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
	F_errmsg(m, int32(127189), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(487618), int32(4145), int32(218311))
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
	F_errmsg_internal(m, int32(675830), v14)
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
	F_errfinish(m, int32(487618), int32(4155), int32(218311))
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
	F_errmsg(m, int32(288569), int32(0))
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
	F_errfinish(m, int32(487618), int32(4163), int32(218311))
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
			v19 = *(*int32)(unsafe.Add(mBase, _consts[1106]))
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
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[1107])))
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
	F_errmsg_internal(m, int32(333218), v14)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(487199), int32(6242), int32(313322))
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
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(542237)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v70
						F_errmsg(m, int32(200112), v7)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(487252), int32(71), int32(499100))
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
	F_errmsg(m, int32(294086), v23)
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
	F_errhint(m, int32(620710), int32(0))
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
	F_errfinish(m, int32(491963), int32(1134), int32(35319))
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
	F_errmsg_internal(m, int32(18080), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(491963), int32(1092), int32(35319))
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
					F_errmsg_internal(m, int32(39263), v11)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(484970), int32(301), int32(238007))
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
					F_errmsg_internal(m, int32(246693), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
						F_errdetail_internal(m, int32(630314), v8)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490042), int32(2886), int32(238145))
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
						F_errmsg(m, int32(110830), int32(0))
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(489552), int32(521), int32(335613))
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
	F_errmsg_internal(m, int32(42992), v17+int32(16))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L2
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(487164), int32(690), int32(9754))
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
	F_errmsg_internal(m, int32(482896), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L2
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(487164), int32(318), int32(302702))
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
	F_errmsg_internal(m, int32(395832), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L2
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(487164), int32(509), int32(9754))
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
	F_errmsg_internal(m, int32(43098), v17)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L2
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(487164), int32(402), int32(302702))
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
	v4 = F_SlruSyncFileTag(m, int32(4365460), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
