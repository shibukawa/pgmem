package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v489 int32
	_ = v489
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
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return v489
L7:
	;
	if v445 != 0 {
		v489 = v6
		goto L6
	} else {
		goto L144
	}
L8:
	;
	v442 = l2
	v445 = v22
	v446 = l3
	goto L7
L9:
	;
	goto L10
L10:
	;
	if l3 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v442 = l2
	v445 = v22
	v446 = l3
	goto L7
L12:
	;
	goto L13
L13:
	;
	v31 = l0
	v32 = l1
	v33 = l2
	v34 = l3
	goto L14
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	switch v44 - int32(92) {
	case 0:
		goto L24
	case 1, 2:
		goto L22
	case 3:
		v424 = v33
		v425 = v34
		goto L16
	default:
		goto L25
	}
L15:
	;
	v442 = v431
	v445 = v427
	v446 = v429
	goto L7
L16:
	;
	v426 = int32(1)
	v427 = base.B2i32(v426 < v32)
	v429 = v425 - v426
	v431 = v424 + v426
	if v32 < int32(2) {
		v442 = v431
		v445 = v427
		v446 = v429
		goto L7
	} else {
		goto L142
	}
L17:
	;
	v417 = F_pg_strncoll(m, v186, v291, v31, v32, l4)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L137
	}
L18:
	;
	v412 = F_pg_strncoll(m, v33, v404-v33, v31, v32, l4)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L136
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L132
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L128
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L124
	}
L22:
	;
	if l4 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L23:
	;
	v61 = v31
	v62 = v32
	v63 = v33
	v64 = v34
	goto L31
L24:
	;
	if v34 <= int32(1) {
		goto L21
	} else {
		goto L28
	}
L25:
	;
	if v44 != int32(37) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v49 = int32(1)
	if base.Ui32(v34) <= base.Ui32(v49) {
		v489 = v49
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v54 != v55 {
		v489 = v6
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v57 = int32(1)
	v424 = v33 + v57
	v425 = v34 - v57
	goto L16
L30:
	;
	if v62 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v74 = int32(1)
	v75 = v64 - v74
	v77 = v63 + v74
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	switch v78 - int32(92) {
	case 0:
		goto L33
	case 1, 2:
		v98 = v78
		goto L30
	case 3:
		goto L35
	default:
		goto L36
	}
L32:
	;
	if v75 == int32(1) {
		goto L20
	} else {
		goto L42
	}
L33:
	;
	goto L32
L34:
	;
	if base.Ui32(int32(2)) < base.Ui32(v64) {
		v61 = v91
		v62 = v92
		v63 = v77
		v64 = v75
		goto L31
	} else {
		goto L41
	}
L35:
	;
	if v62 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v78 == int32(37) {
		v91 = v61
		v92 = v62
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v98 = v78
	goto L30
L38:
	;
	return int32(-1)
L39:
	;
	goto L40
L40:
	;
	v87 = int32(1)
	v91 = v61 + v87
	v92 = v62 - v87
	goto L34
L41:
	;
	v489 = v49
	goto L6
L42:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
	v98 = v97
	goto L30
L43:
	;
	return int32(-1)
L44:
	;
	goto L45
L45:
	;
	v105 = v61
	v106 = v62
	goto L46
L46:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v98&int32(255) != v118 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v489 = int32(-1)
	goto L6
L48:
	;
	v127 = int32(1)
	if v127 < v106 {
		v105 = v105 + v127
		v106 = v106 - v127
		goto L46
	} else {
		goto L56
	}
L49:
	;
	if l4 == int32(0) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v123 = F_SB_MatchText(m, v105, v106, v77, v75, l4)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v122 != 0 {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	if v123 != 0 {
		v489 = v123
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L48
L56:
	;
	goto L47
L57:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v44 == v348 {
		v424 = v33
		v425 = v34
		goto L16
	} else {
		goto L123
	}
L58:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v135 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	if v34 == int32(0) {
		v404 = v33
		goto L18
	} else {
		goto L60
	}
L60:
	;
	v142 = v34
	v144 = int32(0)
	v145 = v33
	goto L64
L61:
	;
	v307 = v32
	v313 = v31
	goto L105
L62:
	;
	v295 = v33
	v296 = v142
	v298 = v145 - v33
	v299 = v145
	v301 = v6
	goto L61
L63:
	;
	v185 = v183 - v33
	v186 = F_palloc(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L75
	}
L64:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	switch v152 - int32(92) {
	case 0:
		goto L67
	case 1, 2:
		v168 = v142
		v169 = v144
		v170 = v145
		goto L66
	case 3:
		goto L68
	default:
		goto L69
	}
L65:
	;
	v175 = int32(1)
	if v169&v175 == int32(0) {
		v404 = v172
		goto L18
	} else {
		goto L74
	}
L66:
	;
	v171 = int32(1)
	v172 = v170 + v171
	v174 = v168 - v171
	if v174 != 0 {
		v142 = v174
		v144 = v169
		v145 = v172
		goto L64
	} else {
		goto L73
	}
L67:
	;
	v162 = v142 - int32(1)
	if v162 == int32(0) {
		goto L19
	} else {
		goto L72
	}
L68:
	;
	if v144&int32(1) == int32(0) {
		goto L62
	} else {
		goto L71
	}
L69:
	;
	if v152 != int32(37) {
		v168 = v142
		v169 = v144
		v170 = v145
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v181 = v142
	v183 = v145
	v184 = v6
	goto L63
L72:
	;
	v165 = int32(1)
	v168 = v162
	v169 = v165
	v170 = v145 + v165
	goto L66
L73:
	;
	goto L65
L74:
	;
	v181 = int32(0)
	v183 = v172
	v184 = v175
	goto L63
L75:
	;
	if base.Ui32(v183) <= base.Ui32(v33) {
		v283 = v186
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v291 = v283 - v186
	if v184 != 0 {
		goto L17
	} else {
		goto L104
	}
L77:
	;
	v190 = v185 & int32(3)
	if v190 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v33-v183) {
		v283 = v223
		goto L76
	} else {
		goto L88
	}
L79:
	;
	v223 = v186
	v225 = v33
	goto L78
L80:
	;
	goto L81
L81:
	;
	v198 = v186
	v200 = v33
	v202 = v6
	goto L82
L82:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v206 != int32(92) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v223 = v212
	v225 = v214
	goto L78
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v206)
	v212 = v198 + int32(1)
	goto L86
L85:
	;
	v212 = v198
	goto L86
L86:
	;
	v213 = int32(1)
	v214 = v200 + v213
	v216 = v202 + v213
	if v216 != v190 {
		v198 = v212
		v200 = v214
		v202 = v216
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	v239 = v223
	v241 = v225
	goto L89
L89:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v247 != int32(92) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v283 = v274
	goto L76
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v247)
	v253 = v239 + int32(1)
	goto L93
L92:
	;
	v253 = v239
	goto L93
L93:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	if v254 != int32(92) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v254)
	v260 = v253 + int32(1)
	goto L96
L95:
	;
	v260 = v253
	goto L96
L96:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+2)))
	if v261 != int32(92) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v261)
	v267 = v260 + int32(1)
	goto L99
L98:
	;
	v267 = v260
	goto L99
L99:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+3)))
	if v268 != int32(92) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v267))) = uint8(v268)
	v274 = v267 + int32(1)
	goto L102
L101:
	;
	v274 = v267
	goto L102
L102:
	;
	v276 = v241 + int32(4)
	if v276 != v183 {
		v239 = v274
		v241 = v276
		goto L89
	} else {
		goto L103
	}
L103:
	;
	goto L90
L104:
	;
	v295 = v186
	v296 = v181
	v298 = v291
	v299 = v183
	v301 = v186
	goto L61
L105:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_SB_MatchText[0]))
	if v320 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v324 = F_pg_strncoll(m, v295, v298, v31, v313-v31, l4)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	if v307 != 0 {
		goto L118
	} else {
		goto L119
	}
L112:
	;
	if v324 != 0 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v326 = F_SB_MatchText(m, v313, v307, v299, v296, l4)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	if v326 != int32(1) {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	if v301 == int32(0) {
		v489 = int32(1)
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_pfree(m, v301)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	return int32(1)
L118:
	;
	v337 = int32(1)
	v307 = v307 - v337
	v313 = v313 + v337
	goto L105
L119:
	;
	v341 = int32(0)
	if v301 == v341 {
		v489 = v341
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_pfree(m, v301)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	return int32(0)
L123:
	;
	v489 = v6
	goto L6
L124:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_SB_MatchText_0), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_SB_MatchText_1), int32(107), int32(_a_F_SB_MatchText_2))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_SB_MatchText_0), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_SB_MatchText_1), int32(169), int32(_a_F_SB_MatchText_2))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
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
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_SB_MatchText_0), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_SB_MatchText_1), int32(237), int32(_a_F_SB_MatchText_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
	return base.B2i32(v412 == int32(0))
L137:
	;
	if v186 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_pfree(m, v186)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	return base.B2i32(v417 == int32(0))
L141:
	;
	goto L140
L142:
	;
	v434 = int32(1)
	if v434 < v425 {
		v31 = v31 + v434
		v32 = v32 - v434
		v33 = v431
		v34 = v429
		goto L14
	} else {
		goto L143
	}
L143:
	;
	goto L15
L144:
	;
	v453 = int32(1)
	if v446 <= int32(0) {
		v489 = v453
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v458 = v442
	v462 = v446
	goto L146
L146:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v469 != int32(37) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v489 = v453
	goto L6
L148:
	;
	return int32(-1)
L149:
	;
	goto L150
L150:
	;
	v474 = int32(1)
	if v474 < v462 {
		v458 = v458 + v474
		v462 = v462 - v474
		goto L146
	} else {
		goto L151
	}
L151:
	;
	goto L147
}
func F_ScanKeyEntryInitializeWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	v3 = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	v11 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v11)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = l0 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ScanKeyEntryInitializeWithInfo[0]))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l5)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v11
	return
}
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	v2 = l1
	v3 = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(4080218931200)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	F_fmgr_info(m, l3, l0+int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
func F_SerializeSnapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v14)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v11
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v10)
	if v10&int32(1) != 0 {
		v23 = v9
	} else {
		v23 = int32(0)
	}
	if v14 != 0 {
		v24 = v23
	} else {
		v24 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v26 == int32(0) {
	} else {
		v30 = v26 << (uint(int32(2)) % 32)
		if v30 == int32(0) {
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			base.MemoryCopy(m, l1+int32(24), v35, v30)
		}
	}
	if v24 <= int32(0) {
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v42 = v40 << (uint(int32(2)) % 32)
		if v42 == int32(0) {
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			base.MemoryCopy(m, l1+v45<<(uint(int32(2))%32)+int32(24), v51, v42)
		}
	}
	return
}
func F_SetMatViewPopulatedState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v17 = F_SearchSysCacheCopy(m, int32(57), v15, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
					F_errmsg_internal(m, int32(_a_F_SetMatViewPopulatedState_0), v8)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetMatViewPopulatedState_1), int32(96), int32(_a_F_SetMatViewPopulatedState_2))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
				*(*uint8)(unsafe.Add(mBase, uint32(v35+v36)+129)) = uint8(v2)
				F_CatalogTupleUpdate(m, v12, v17+int32(4), v17)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_pfree(m, v17)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_relation_close(m, v12, int32(3))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_CommandCounterIncrement(m)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_SignalHandlerForShutdownRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_SwitchToSharedLatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[0]))
	v7 = v5 + int32(20)
	*(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[1])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[2]))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(1)
	F_ModifyWaitEvent(m, v10, v11, v11, v7)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v17 = v7
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[1]))
	v17 = v16
	goto L3
L6:
	;
	return
L7:
	;
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[3]))
	if v28 == v24 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[4]))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v58 = F_pgmem_kill(m, v24, int32(23))
	mBase = m.M
	goto L7
L14:
	;
	m.G0 = v32 + int32(16)
	goto L6
L15:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+15)) = uint8(v38)
	goto L16
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[5]))
	v46 = F_write(m, v42, v32+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v46 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	goto L14
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[6]))
	if v50 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
func F_SwitchToUntrustedUser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)))) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_member_can_set_role(m, v18, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		if v19 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v22 = F_member_can_set_role(m, l0, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v22 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[1])) = v24
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[0])) = l0
					v43 = int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[1])) = v24 | int32(2)
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[0])) = l0
					v37 = int32(_a_F_SwitchToUntrustedUser_0)
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[2]))
					v41 = v39 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[2])) = v41
					v43 = v41
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v43
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v57 = F_GetUserNameFromId(m, v55, int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v60 = F_GetUserNameFromId(m, l0, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v57
							F_errmsg(m, int32(_a_F_SwitchToUntrustedUser_1), v8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_SwitchToUntrustedUser_2), int32(45), int32(_a_F_SwitchToUntrustedUser_3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
func F_sanitize_char_1(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	Fn13993(m, l0, int32(_a_F_sanitize_char_1_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_scalargesel(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = F_scalarineqsel_wrapper(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_scanNSItemForColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_scanRTEForColumn(m, l0, v15, v16, l3, l4, v6, v6)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			v24 = base.B2i32(v19 == int32(-6))
			v25 = int32(0)
			v26 = base.B2i32(v25 <= v19)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if base.B2i32(v24|v26 == v25)&base.B2i32(v30 == int32(28)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v274 = m.ExcPending
				if v274 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_scanNSItemForColumn_0))
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
						F_errmsg(m, int32(_a_F_scanNSItemForColumn_1), v13)
						mBase = m.M
						v281 = m.ExcPending
						if v281 != 0 {
							return int32(0)
						} else {
							F_parser_errposition(m, l0, l4)
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(713), int32(_a_F_scanNSItemForColumn_3))
								mBase = m.M
								v288 = m.ExcPending
								if v288 != 0 {
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
				if base.B2i32(v24|v26 == int32(0))&base.B2i32(v30 == int32(43)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v292 = m.ExcPending
					if v292 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(_a_F_scanNSItemForColumn_0))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l3
							F_errmsg(m, int32(_a_F_scanNSItemForColumn_4), v11+int32(-48))
							mBase = m.M
							v301 = m.ExcPending
							if v301 != 0 {
								return int32(0)
							} else {
								F_parser_errposition(m, l0, l4)
								mBase = m.M
								v303 = m.ExcPending
								if v303 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(726), int32(_a_F_scanNSItemForColumn_3))
									mBase = m.M
									v308 = m.ExcPending
									if v308 != 0 {
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
					if base.B2i32(v24|v26 == int32(0))&base.B2i32(v30 == int32(18)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v312 = m.ExcPending
						if v312 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_scanNSItemForColumn_0))
							mBase = m.M
							v315 = m.ExcPending
							if v315 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l3
								F_errmsg(m, int32(_a_F_scanNSItemForColumn_5), v11+int32(-32))
								mBase = m.M
								v321 = m.ExcPending
								if v321 != 0 {
									return int32(0)
								} else {
									F_parser_errposition(m, l0, l4)
									mBase = m.M
									v323 = m.ExcPending
									if v323 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(737), int32(_a_F_scanNSItemForColumn_3))
										mBase = m.M
										v328 = m.ExcPending
										if v328 != 0 {
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
						if int32(0) < v19 {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v51 = v48 + v19<<(uint(int32(5))%32)
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(32))))
							if v54 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v332 = m.ExcPending
								if v332 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50360452))
									mBase = m.M
									v335 = m.ExcPending
									if v335 != 0 {
										return int32(0)
									} else {
										v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v337
										*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l3
										F_errmsg(m, int32(_a_F_scanNSItemForColumn_6), v11+int32(-16))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(751), int32(_a_F_scanNSItemForColumn_3))
											mBase = m.M
											v349 = m.ExcPending
											if v349 != 0 {
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
								v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51-int32(28)))))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(24))))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(20))))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(16))))
								v69 = F_makeVar(m, v54, v59, v62, v65, v68, l2)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(8))))
									*(*int32)(unsafe.Add(mBase, uint32(v69)+36)) = v73
									v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51-int32(4)))))
									*(*uint16)(unsafe.Add(mBase, uint32(v69)+40)) = uint16(v77)
									v89 = v69
									*(*int32)(unsafe.Add(mBase, uint32(v89)+44)) = l4
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v92
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									if v95 == int32(0) {
										v156 = l0
									} else {
										v99 = v95 & int32(7)
										if base.Ui32(int32(8)) <= base.Ui32(v95) {
											v106 = int32(0)
											v108 = l0
											for {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
												v124 = v106 + int32(8)
												if v124 != v95&int32(-8) {
													v106 = v124
													v108 = v122
													continue
												} else {
													break
												}
												break
											}
											if v99 == int32(0) {
												v156 = v122
											} else {
												v131 = v122
												v140 = int32(0)
												v142 = v131
												for {
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													v151 = v140 + int32(1)
													if v151 != v99 {
														v140 = v151
														v142 = v149
														continue
													} else {
														break
													}
													break
												}
												v156 = v149
											}
										} else {
											v131 = l0
											v140 = int32(0)
											v142 = v131
											for {
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												v151 = v140 + int32(1)
												if v151 != v99 {
													v140 = v151
													v142 = v149
													continue
												} else {
													break
												}
												break
											}
											v156 = v149
										}
									}
									if v94 <= int32(0) {
										v185 = v95
										if v185 == int32(0) {
											v243 = l0
										} else {
											v189 = v185 & int32(7)
											if base.Ui32(int32(8)) <= base.Ui32(v185) {
												v195 = l0
												v198 = int32(0)
												for {
													v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
													v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
													v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
													v214 = v198 + int32(8)
													if v214 != v185&int32(-8) {
														v195 = v212
														v198 = v214
														continue
													} else {
														break
													}
													break
												}
												if v189 == int32(0) {
													v243 = v212
												} else {
													v218 = v212
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											} else {
												v218 = l0
												v229 = v218
												v232 = int32(0)
												for {
													v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
													v241 = v232 + int32(1)
													if v241 != v189 {
														v229 = v239
														v232 = v241
														continue
													} else {
														break
													}
													break
												}
												v243 = v239
											}
										}
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
										v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
										F_markRTEForSelectPriv(m, v243, v253, v254)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
											return int32(0)
										} else {
											v262 = v89
											m.G0 = v13 - int32(-64)
											return v262
										}
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
										if v165 == int32(0) {
											v185 = v95
											if v185 == int32(0) {
												v243 = l0
											} else {
												v189 = v185 & int32(7)
												if base.Ui32(int32(8)) <= base.Ui32(v185) {
													v195 = l0
													v198 = int32(0)
													for {
														v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
														v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
														v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
														v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
														v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
														v214 = v198 + int32(8)
														if v214 != v185&int32(-8) {
															v195 = v212
															v198 = v214
															continue
														} else {
															break
														}
														break
													}
													if v189 == int32(0) {
														v243 = v212
													} else {
														v218 = v212
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												} else {
													v218 = l0
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											}
											v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
											v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
											F_markRTEForSelectPriv(m, v243, v253, v254)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
												return int32(0)
											} else {
												v262 = v89
												m.G0 = v13 - int32(-64)
												return v262
											}
										} else {
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
											if v168 < v94 {
												v185 = v95
												if v185 == int32(0) {
													v243 = l0
												} else {
													v189 = v185 & int32(7)
													if base.Ui32(int32(8)) <= base.Ui32(v185) {
														v195 = l0
														v198 = int32(0)
														for {
															v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
															v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
															v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
															v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
															v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
															v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
															v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
															v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
															v214 = v198 + int32(8)
															if v214 != v185&int32(-8) {
																v195 = v212
																v198 = v214
																continue
															} else {
																break
															}
															break
														}
														if v189 == int32(0) {
															v243 = v212
														} else {
															v218 = v212
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													} else {
														v218 = l0
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												}
												v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
												v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
												F_markRTEForSelectPriv(m, v243, v253, v254)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return int32(0)
												} else {
													v262 = v89
													m.G0 = v13 - int32(-64)
													return v262
												}
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+v94<<(uint(int32(2))%32)-int32(4))))
												if v176 == int32(0) {
													v185 = v95
													if v185 == int32(0) {
														v243 = l0
													} else {
														v189 = v185 & int32(7)
														if base.Ui32(int32(8)) <= base.Ui32(v185) {
															v195 = l0
															v198 = int32(0)
															for {
																v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																v214 = v198 + int32(8)
																if v214 != v185&int32(-8) {
																	v195 = v212
																	v198 = v214
																	continue
																} else {
																	break
																}
																break
															}
															if v189 == int32(0) {
																v243 = v212
															} else {
																v218 = v212
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														} else {
															v218 = l0
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													}
													v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
													v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
													F_markRTEForSelectPriv(m, v243, v253, v254)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int32(0)
													} else {
														v262 = v89
														m.G0 = v13 - int32(-64)
														return v262
													}
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
													v180 = F_bms_union(m, v179, v176)
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v180
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
														v185 = v183
														if v185 == int32(0) {
															v243 = l0
														} else {
															v189 = v185 & int32(7)
															if base.Ui32(int32(8)) <= base.Ui32(v185) {
																v195 = l0
																v198 = int32(0)
																for {
																	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																	v214 = v198 + int32(8)
																	if v214 != v185&int32(-8) {
																		v195 = v212
																		v198 = v214
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v189 == int32(0) {
																	v243 = v212
																} else {
																	v218 = v212
																	v229 = v218
																	v232 = int32(0)
																	for {
																		v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																		v241 = v232 + int32(1)
																		if v241 != v189 {
																			v229 = v239
																			v232 = v241
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v243 = v239
																}
															} else {
																v218 = l0
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														}
														v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
														F_markRTEForSelectPriv(m, v243, v253, v254)
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															v262 = v89
															m.G0 = v13 - int32(-64)
															return v262
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v79 = base.I32_extend16_s(v19)
							v80 = F_SystemAttributeDefinition(m, v79)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+68))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+76))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+96))
								v86 = F_makeVar(m, v82, v79, v83, v84, v85, l2)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v89 = v86
									*(*int32)(unsafe.Add(mBase, uint32(v89)+44)) = l4
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v92
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									if v95 == int32(0) {
										v156 = l0
									} else {
										v99 = v95 & int32(7)
										if base.Ui32(int32(8)) <= base.Ui32(v95) {
											v106 = int32(0)
											v108 = l0
											for {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
												v124 = v106 + int32(8)
												if v124 != v95&int32(-8) {
													v106 = v124
													v108 = v122
													continue
												} else {
													break
												}
												break
											}
											if v99 == int32(0) {
												v156 = v122
											} else {
												v131 = v122
												v140 = int32(0)
												v142 = v131
												for {
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													v151 = v140 + int32(1)
													if v151 != v99 {
														v140 = v151
														v142 = v149
														continue
													} else {
														break
													}
													break
												}
												v156 = v149
											}
										} else {
											v131 = l0
											v140 = int32(0)
											v142 = v131
											for {
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												v151 = v140 + int32(1)
												if v151 != v99 {
													v140 = v151
													v142 = v149
													continue
												} else {
													break
												}
												break
											}
											v156 = v149
										}
									}
									if v94 <= int32(0) {
										v185 = v95
										if v185 == int32(0) {
											v243 = l0
										} else {
											v189 = v185 & int32(7)
											if base.Ui32(int32(8)) <= base.Ui32(v185) {
												v195 = l0
												v198 = int32(0)
												for {
													v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
													v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
													v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
													v214 = v198 + int32(8)
													if v214 != v185&int32(-8) {
														v195 = v212
														v198 = v214
														continue
													} else {
														break
													}
													break
												}
												if v189 == int32(0) {
													v243 = v212
												} else {
													v218 = v212
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											} else {
												v218 = l0
												v229 = v218
												v232 = int32(0)
												for {
													v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
													v241 = v232 + int32(1)
													if v241 != v189 {
														v229 = v239
														v232 = v241
														continue
													} else {
														break
													}
													break
												}
												v243 = v239
											}
										}
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
										v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
										F_markRTEForSelectPriv(m, v243, v253, v254)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
											return int32(0)
										} else {
											v262 = v89
											m.G0 = v13 - int32(-64)
											return v262
										}
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
										if v165 == int32(0) {
											v185 = v95
											if v185 == int32(0) {
												v243 = l0
											} else {
												v189 = v185 & int32(7)
												if base.Ui32(int32(8)) <= base.Ui32(v185) {
													v195 = l0
													v198 = int32(0)
													for {
														v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
														v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
														v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
														v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
														v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
														v214 = v198 + int32(8)
														if v214 != v185&int32(-8) {
															v195 = v212
															v198 = v214
															continue
														} else {
															break
														}
														break
													}
													if v189 == int32(0) {
														v243 = v212
													} else {
														v218 = v212
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												} else {
													v218 = l0
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											}
											v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
											v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
											F_markRTEForSelectPriv(m, v243, v253, v254)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
												return int32(0)
											} else {
												v262 = v89
												m.G0 = v13 - int32(-64)
												return v262
											}
										} else {
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
											if v168 < v94 {
												v185 = v95
												if v185 == int32(0) {
													v243 = l0
												} else {
													v189 = v185 & int32(7)
													if base.Ui32(int32(8)) <= base.Ui32(v185) {
														v195 = l0
														v198 = int32(0)
														for {
															v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
															v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
															v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
															v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
															v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
															v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
															v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
															v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
															v214 = v198 + int32(8)
															if v214 != v185&int32(-8) {
																v195 = v212
																v198 = v214
																continue
															} else {
																break
															}
															break
														}
														if v189 == int32(0) {
															v243 = v212
														} else {
															v218 = v212
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													} else {
														v218 = l0
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												}
												v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
												v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
												F_markRTEForSelectPriv(m, v243, v253, v254)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return int32(0)
												} else {
													v262 = v89
													m.G0 = v13 - int32(-64)
													return v262
												}
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+v94<<(uint(int32(2))%32)-int32(4))))
												if v176 == int32(0) {
													v185 = v95
													if v185 == int32(0) {
														v243 = l0
													} else {
														v189 = v185 & int32(7)
														if base.Ui32(int32(8)) <= base.Ui32(v185) {
															v195 = l0
															v198 = int32(0)
															for {
																v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																v214 = v198 + int32(8)
																if v214 != v185&int32(-8) {
																	v195 = v212
																	v198 = v214
																	continue
																} else {
																	break
																}
																break
															}
															if v189 == int32(0) {
																v243 = v212
															} else {
																v218 = v212
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														} else {
															v218 = l0
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													}
													v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
													v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
													F_markRTEForSelectPriv(m, v243, v253, v254)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int32(0)
													} else {
														v262 = v89
														m.G0 = v13 - int32(-64)
														return v262
													}
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
													v180 = F_bms_union(m, v179, v176)
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v180
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
														v185 = v183
														if v185 == int32(0) {
															v243 = l0
														} else {
															v189 = v185 & int32(7)
															if base.Ui32(int32(8)) <= base.Ui32(v185) {
																v195 = l0
																v198 = int32(0)
																for {
																	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																	v214 = v198 + int32(8)
																	if v214 != v185&int32(-8) {
																		v195 = v212
																		v198 = v214
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v189 == int32(0) {
																	v243 = v212
																} else {
																	v218 = v212
																	v229 = v218
																	v232 = int32(0)
																	for {
																		v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																		v241 = v232 + int32(1)
																		if v241 != v189 {
																			v229 = v239
																			v232 = v241
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v243 = v239
																}
															} else {
																v218 = l0
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														}
														v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
														F_markRTEForSelectPriv(m, v243, v253, v254)
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															v262 = v89
															m.G0 = v13 - int32(-64)
															return v262
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
		} else {
			v262 = v6
			m.G0 = v13 - int32(-64)
			return v262
		}
	}
}
func F_scanner_finish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if base.Ui32(int32(_a_F_scanner_finish_0)) <= base.Ui32(v4) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_pfree(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v11 = v10
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
			if int32(_a_F_scanner_finish_0) <= v12 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
				F_pfree(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v11 = v3
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
		if int32(_a_F_scanner_finish_0) <= v12 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
			F_pfree(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_scram_exchange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int64
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1174 int32
	_ = v1174
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1474 int32
	_ = v1474
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1584 int64
	_ = v1584
	var v1586 int64
	_ = v1586
	var v1588 int64
	_ = v1588
	var v1590 int64
	_ = v1590
	var v1592 int64
	_ = v1592
	var v1594 int64
	_ = v1594
	var v1596 int64
	_ = v1596
	var v1598 int64
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1609 int32
	_ = v1609
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1748 int32
	_ = v1748
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1882 int32
	_ = v1882
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1968 int32
	_ = v1968
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
	var v1994 int32
	_ = v1994
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2036 int32
	_ = v2036
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2055 int32
	_ = v2055
	var v2060 int32
	_ = v2060
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v7
	if l1 == v7 {
		goto L23
	} else {
		goto L24
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L26
	} else {
		goto L546
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L26
	} else {
		goto L530
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L26
	} else {
		goto L527
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L26
	} else {
		goto L511
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L26
	} else {
		goto L506
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L26
	} else {
		goto L501
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L26
	} else {
		goto L496
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L26
	} else {
		goto L492
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L26
	} else {
		goto L489
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L26
	} else {
		goto L485
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L26
	} else {
		goto L481
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L26
	} else {
		goto L477
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L26
	} else {
		goto L473
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L26
	} else {
		goto L469
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L26
	} else {
		goto L464
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L26
	} else {
		goto L458
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L26
	} else {
		goto L453
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L26
	} else {
		goto L447
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L26
	} else {
		goto L442
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L26
	} else {
		goto L437
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L26
	} else {
		goto L432
	}
L22:
	;
	m.G0 = v19 + int32(224)
	return v1609
L23:
	;
	v26 = F_pstrdup(m, int32(_a_F_scram_exchange_0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	return int32(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v26
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v1609 = v7
	goto L22
L28:
	;
	v35 = F_strlen(m, l1)
	mBase = m.M
	if v35 != l2 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v37 {
	case 0:
		goto L34
	case 1:
		goto L33
	default:
		goto L32
	}
L30:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1574 != 0 {
		goto L427
	} else {
		goto L428
	}
L31:
	;
	v1550 = int32(0)
	v1551 = int32(2)
	if l5 == v1550 {
		v1560 = v1550
		v1565 = v1551
		goto L30
	} else {
		goto L425
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L26
	} else {
		goto L422
	}
L33:
	;
	v544 = F_pstrdup(m, l1)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L26
	} else {
		goto L148
	}
L34:
	;
	v38 = F_pstrdup(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v38
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v41)
	switch v41 - int32(110) {
	case 0:
		goto L42
	default:
		goto L37
	case 2:
		goto L40
	case 11:
		goto L41
	}
L36:
	;
	v229 = v98 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v229
	v231 = F_pstrdup(m, v229)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L26
	} else {
		goto L89
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L26
	} else {
		goto L83
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L26
	} else {
		goto L65
	}
L39:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v99 == int32(44) {
		goto L36
	} else {
		goto L57
	}
L40:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v61 == int32(0) {
		goto L15
	} else {
		goto L47
	}
L41:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v53 == int32(1) {
		goto L17
	} else {
		goto L45
	}
L42:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v45 == int32(1) {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v48 != int32(44) {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v98 = v38 + int32(2)
	goto L39
L45:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v56 != int32(44) {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v98 = v38 + int32(2)
	goto L39
L47:
	;
	v67 = F_read_attr_value(m, v19+int32(192), int32(112))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v69 = int32(_a_F_scram_exchange_1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_exchange[0])))
	if base.B2i32(v72 == int32(0))|base.B2i32(v72 != v75) != 0 {
		v93 = v72
		v94 = v75
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v93-v94 != 0 {
		goto L38
	} else {
		goto L56
	}
L50:
	;
	goto L49
L51:
	;
	v78 = v67
	v79 = v69
	goto L52
L52:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v83
		v94 = v82
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v93 = v83
	v94 = v82
	goto L50
L54:
	;
	v86 = int32(1)
	if v83 == v82 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v98 = v96
	goto L39
L57:
	;
	if v99 == int32(97) {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	F_sanitize_char_2(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(_a_F_scram_exchange_3)
	F_errdetail(m, int32(_a_F_scram_exchange_4), v19+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L26
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1081), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	v139 = int32(0)
	goto L67
L67:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+v67))))
	if v155 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_scram_exchange[1]))) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(_a_F_scram_exchange_7)
	F_errmsg(m, int32(_a_F_scram_exchange_8), v19+int32(80))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L26
	} else {
		goto L81
	}
L69:
	;
	goto L68
L70:
	;
	v187 = v139
	goto L69
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(94)) <= base.Ui32((v155-int32(33))&int32(255)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v165 = int32(63)
	goto L75
L74:
	;
	v165 = v155
	goto L75
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_scram_exchange[1]))) = uint8(v165)
	v168 = v139 | int32(1)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v168))))
	if v170 == int32(0) {
		v187 = v168
		goto L69
	} else {
		goto L76
	}
L76:
	;
	if base.Ui32(int32(94)) <= base.Ui32((v170-int32(33))&int32(255)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v180 = int32(63)
	goto L79
L78:
	;
	v180 = v170
	goto L79
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_scram_exchange[1]))) = uint8(v180)
	v182 = int32(30)
	v184 = v139 + int32(2)
	if v184 != v182 {
		v139 = v184
		goto L67
	} else {
		goto L80
	}
L80:
	;
	v187 = v182
	goto L69
L81:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1059), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L26
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L26
	} else {
		goto L85
	}
L85:
	;
	v215 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38))))
	F_sanitize_char_2(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L26
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_scram_exchange_3)
	F_errdetail(m, int32(_a_F_scram_exchange_9), v19)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1066), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L26
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v231
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v234 == int32(109) {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	v238 = v19 + int32(192)
	v240 = F_read_attr_value(m, v238, int32(110))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L26
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v240
	v244 = F_read_attr_value(m, v238, int32(114))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L26
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v244
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v248 = base.I32_extend8_s(v247)
	if int32(33) <= v248 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v253 = v248
	v258 = v244
	goto L96
L94:
	;
	v282 = v248
	goto L95
L95:
	;
	if v282 != 0 {
		goto L12
	} else {
		goto L100
	}
L96:
	;
	v268 = v253 & int32(255)
	if base.B2i32(v268 == int32(44))|base.B2i32(v268 == int32(127)) != 0 {
		goto L12
	} else {
		goto L98
	}
L97:
	;
	v282 = v277
	goto L95
L98:
	;
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v258)+1)))
	v277 = base.I32_extend8_s(v274)
	if int32(32) < v277 {
		v253 = v277
		v258 = v258 + int32(1)
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v297 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	goto L104
L102:
	;
	goto L103
L103:
	;
	v338 = v19 + int32(192)
	v340 = int32(0)
	v344 = m.G0
	v346 = v344 - int32(16)
	m.G0 = v346
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v340
	v352 = F_open(m, int32(_a_F_scram_exchange_10), v340, v346)
	mBase = m.M
	if v352 != int32(-1) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v317 = F_read_any_attr(m, v19+int32(192), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L26
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v320 != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	if v385 == int32(0) {
		goto L11
	} else {
		goto L121
	}
L109:
	;
	goto L113
L110:
	;
	v385 = v340
	goto L111
L111:
	;
	m.G0 = v346 + int32(16)
	goto L108
L112:
	;
	v380 = F_close(m, v352)
	mBase = m.M
	v385 = v378
	goto L111
L113:
	;
	v358 = v338
	v359 = int32(18)
	goto L114
L114:
	;
	v364 = F_read(m, v352, v358, v359)
	mBase = m.M
	if v364 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v378 = int32(1)
	goto L112
L116:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_scram_exchange[2]))
	if v368 == int32(27) {
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v373 = v359 - v364
	if v373 != 0 {
		v358 = v358 + v364
		v359 = v373
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v378 = int32(0)
	goto L112
L120:
	;
	goto L115
L121:
	;
	v396 = base.I32_div_s(int32(20), int32(3))
	v398 = v396 << (uint(int32(2)) % 32)
	goto L122
L122:
	;
	v401 = F_palloc(m, v398+int32(1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L26
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v401
	goto L127
L124:
	;
	if v516 < int32(0) {
		goto L10
	} else {
		goto L145
	}
L125:
	;
	if v398 != 0 {
		goto L142
	} else {
		goto L143
	}
L126:
	;
	if v398 < v458-v401+int32(4) {
		goto L125
	} else {
		goto L138
	}
L127:
	;
	v413 = v338
	v414 = int32(0)
	v417 = v401
	v418 = int32(2)
	goto L130
L129:
	;
	v516 = v458 - v401
	goto L124
L130:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v424 = v420<<(uint(v418<<(uint(int32(3))%32))%32) | v414
	if int32(0) < v418 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v459 != int32(2) {
		goto L126
	} else {
		goto L137
	}
L132:
	;
	v457 = v424
	v458 = v417
	v459 = v418 - int32(1)
	goto L134
L133:
	;
	if v398 < v417-v401+int32(4) {
		goto L125
	} else {
		goto L135
	}
L134:
	;
	v461 = v413 + int32(1)
	if base.Ui32(v461) < base.Ui32(v19+int32(210)) {
		v413 = v461
		v414 = v457
		v417 = v458
		v418 = v459
		goto L130
	} else {
		goto L136
	}
L135:
	;
	v433 = int32(63)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424&v433)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+3)) = uint8(v435)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v424)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417))) = uint8(v439)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v424)>>(uint(int32(6))%32))&v433)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+2)) = uint8(v445)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v424)>>(uint(int32(12))%32))&v433)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+1)) = uint8(v451)
	v457 = int32(0)
	v458 = v417 + int32(4)
	v459 = int32(2)
	goto L134
L136:
	;
	goto L131
L137:
	;
	goto L129
L138:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v479)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458)+1)) = uint8(v485)
	if v459 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	v495 = v494
	goto L141
L140:
	;
	v495 = int32(61)
	goto L141
L141:
	;
	v496 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v458)+3)) = uint8(v496)
	*(*uint8)(unsafe.Add(mBase, uint32(v458)+2)) = uint8(v495)
	v516 = v458 + int32(4) - v401
	goto L124
L142:
	;
	base.MemoryFill(m, v401, int32(0), v398)
	goto L144
L143:
	;
	goto L144
L144:
	;
	v516 = int32(-1)
	goto L124
L145:
	;
	v519 = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*uint8)(unsafe.Add(mBase, uint32(v520+v516))) = uint8(v519)
	v524 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v525
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = base.I64_rotl(v524, int64(32))
	v535 = F_psprintf(m, int32(_a_F_scram_exchange_11), v19+int32(32))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L26
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v535
	v538 = F_pstrdup(m, v535)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L26
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v538
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v1560 = int32(0)
	v1565 = v519
	goto L30
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v544
	v550 = F_read_attr_value(m, v19+int32(192), int32(99))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L26
	} else {
		goto L149
	}
L149:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v552 == int32(1) {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	v555 = int32(_a_F_scram_exchange_12)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_exchange[4])))
	if base.B2i32(v558 == int32(0))|base.B2i32(v558 != v561) != 0 {
		v579 = v558
		v580 = v561
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v620 = F_read_attr_value(m, v19+int32(192), int32(114))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L26
	} else {
		goto L172
	}
L152:
	;
	if v579-v580 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	goto L152
L154:
	;
	v564 = v550
	v565 = v555
	goto L155
L155:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+1)))
	if v569 == int32(0) {
		v579 = v569
		v580 = v568
		goto L153
	} else {
		goto L157
	}
L156:
	;
	v579 = v569
	v580 = v568
	goto L153
L157:
	;
	v572 = int32(1)
	if v569 == v568 {
		v564 = v564 + v572
		v565 = v565 + v572
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v584 == int32(110) {
		goto L151
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v587 = int32(_a_F_scram_exchange_13)
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_exchange[5])))
	if base.B2i32(v590 == int32(0))|base.B2i32(v590 != v593) != 0 {
		v611 = v590
		v612 = v593
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L161
L163:
	;
	if v611-v612 != 0 {
		goto L8
	} else {
		goto L170
	}
L164:
	;
	goto L163
L165:
	;
	v596 = v550
	v597 = v587
	goto L166
L166:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	if v601 == int32(0) {
		v611 = v601
		v612 = v600
		goto L164
	} else {
		goto L168
	}
L167:
	;
	v611 = v601
	v612 = v600
	goto L164
L168:
	;
	v604 = int32(1)
	if v601 == v600 {
		v596 = v596 + v604
		v597 = v597 + v604
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v614 != int32(121) {
		goto L8
	} else {
		goto L171
	}
L171:
	;
	goto L151
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v620
	goto L173
L173:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v644 = F_read_any_attr(m, v19+int32(192), v19+int32(160))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L26
	} else {
		goto L175
	}
L174:
	;
	v649 = F_strlen(m, v644)
	mBase = m.M
	v653 = v649 * int32(3) >> (uint(int32(2)) % 32)
	goto L177
L175:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+160)))
	if v646 != int32(112) {
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v654 = F_palloc(m, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L26
	} else {
		goto L178
	}
L178:
	;
	v656 = F_strlen(m, v644)
	mBase = m.M
	v657 = int32(0)
	if v657 < v656 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v842 != v843 {
		goto L7
	} else {
		goto L225
	}
L180:
	;
	if v653 != 0 {
		goto L222
	} else {
		goto L223
	}
L181:
	;
	v666 = v644 + v656
	v667 = v644
	v671 = v657
	v672 = v654
	v674 = v657
	v675 = v657
	goto L184
L182:
	;
	v813 = v654
	goto L183
L183:
	;
	v842 = v813 - v654
	goto L179
L184:
	;
	v679 = v667 + int32(1)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	if v680 != int32(61) {
		goto L192
	} else {
		goto L193
	}
L185:
	;
	if v801 != 0 {
		goto L180
	} else {
		goto L221
	}
L186:
	;
	if base.Ui32(v799) < base.Ui32(v666) {
		v667 = v799
		v671 = v801
		v672 = v802
		v674 = v804
		v675 = v805
		goto L184
	} else {
		goto L220
	}
L187:
	;
	if v653 < v672-v654+int32(1) {
		goto L180
	} else {
		goto L207
	}
L188:
	;
	v757 = v679
	v758 = int32(2)
	v761 = v675 << (uint(int32(6)) % 32)
	goto L187
L189:
	;
	v746 = v743 + v742<<(uint(int32(6))%32)
	v748 = v739 + int32(1)
	if v748 == int32(4) {
		v757 = v740
		v758 = v741
		v761 = v746
		goto L187
	} else {
		goto L206
	}
L190:
	;
	v739 = int32(3)
	v740 = v667 + int32(2)
	v741 = int32(1)
	v742 = v699
	v743 = v694
	goto L189
L191:
	;
	if base.Ui32(int32(125)) < base.Ui32((v718-int32(1))&int32(255)) {
		goto L180
	} else {
		goto L204
	}
L192:
	;
	v684 = v680 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v684))|base.B2i32(int32(1)<<(uint(v684)%32)&int32(_a_F_scram_exchange_14) == int32(0)) != 0 {
		v718 = v680
		v719 = v671
		v720 = v679
		v721 = v674
		v722 = v675
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v694 = int32(0)
	if v674 != 0 {
		v739 = v671
		v740 = v679
		v741 = v674
		v742 = v675
		v743 = v694
		goto L189
	} else {
		goto L196
	}
L195:
	;
	goto L180
L196:
	;
	switch v671 - int32(2) {
	case 0:
		goto L197
	case 1:
		goto L188
	default:
		goto L180
	}
L197:
	;
	if base.Ui32(v666) <= base.Ui32(v679) {
		goto L180
	} else {
		goto L198
	}
L198:
	;
	v699 = v675 << (uint(int32(6)) % 32)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	if v700 == int32(61) {
		goto L190
	} else {
		goto L199
	}
L199:
	;
	v704 = v700 - int32(9)
	if int32(1)<<(uint(v704)%32)&int32(_a_F_scram_exchange_14) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v712 = base.B2i32(base.Ui32(v704) <= base.Ui32(int32(23)))
	goto L202
L201:
	;
	v712 = int32(0)
	goto L202
L202:
	;
	if v712 != 0 {
		goto L180
	} else {
		goto L203
	}
L203:
	;
	v718 = v700
	v719 = int32(3)
	v720 = v667 + int32(2)
	v721 = int32(1)
	v722 = v699
	goto L191
L204:
	;
	v730 = int32(*(*int8)(unsafe.Add(mBase, uint32(v718)+uint32(_c_F_scram_exchange[6]))))
	if v730 < int32(0) {
		goto L180
	} else {
		goto L205
	}
L205:
	;
	v739 = v719
	v740 = v720
	v741 = v721
	v742 = v722
	v743 = v730
	goto L189
L206:
	;
	v799 = v740
	v801 = v748
	v802 = v672
	v804 = v741
	v805 = v746
	goto L186
L207:
	;
	v767 = int32(base.Ui32(v761) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v672))) = uint8(v767)
	v770 = v672 + int32(1)
	if base.Ui32(v758) < base.Ui32(int32(2)) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v774 = v758
	goto L210
L209:
	;
	v774 = int32(0)
	goto L210
L210:
	;
	if v774 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v653 < v770-v654+int32(1) {
		goto L180
	} else {
		goto L214
	}
L212:
	;
	v786 = v770
	goto L213
L213:
	;
	if v758 != 0 {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v782 = int32(base.Ui32(v761) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v672)+1)) = uint8(v782)
	v786 = v672 + int32(2)
	goto L213
L215:
	;
	v799 = v757
	v801 = int32(0)
	v802 = v796
	v804 = v797
	v805 = int32(0)
	goto L186
L216:
	;
	v796 = v786
	v797 = v758
	goto L215
L217:
	;
	goto L218
L218:
	;
	if v653 < v786-v654+int32(1) {
		goto L180
	} else {
		goto L219
	}
L219:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v786))) = uint8(v761)
	v796 = v786 + int32(1)
	v797 = int32(0)
	goto L215
L220:
	;
	goto L185
L221:
	;
	v813 = v802
	goto L183
L222:
	;
	base.MemoryFill(m, v654, int32(0), v653)
	goto L224
L223:
	;
	goto L224
L224:
	;
	v842 = int32(-1)
	goto L179
L225:
	;
	v846 = l0 + int32(152)
	if v842 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	base.MemoryCopy(m, v846, v654, v842)
	goto L228
L227:
	;
	goto L228
L228:
	;
	F_pfree(m, v654)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L26
	} else {
		goto L229
	}
L229:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850))))
	if v851 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v853 = F_palloc(m, v639-v544)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L26
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v853
	v858 = v544 ^ int32(-1) + v639
	if v858 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	base.MemoryCopy(m, v853, l1, v858)
	goto L234
L233:
	;
	goto L234
L234:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v862 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v860+v858))) = uint8(v862)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v865 = F_strlen(m, v864)
	mBase = m.M
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v867 = F_strlen(m, v866)
	mBase = m.M
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v869 = F_strlen(m, v868)
	mBase = m.M
	if v869 != v865+v867 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v865) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	if v933 != 0 {
		goto L5
	} else {
		goto L254
	}
L237:
	;
	v933 = int32(0)
	goto L236
L238:
	;
	v907 = v902
	v908 = v903
	v909 = v904
	goto L248
L239:
	;
	if (v868|v864)&int32(3) != 0 {
		v902 = v868
		v903 = v864
		v904 = v865
		goto L238
	} else {
		goto L242
	}
L240:
	;
	v895 = v868
	v896 = v864
	v897 = v865
	goto L241
L241:
	;
	if v897 == int32(0) {
		goto L237
	} else {
		goto L247
	}
L242:
	;
	v879 = v868
	v880 = v864
	v881 = v865
	goto L243
L243:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	if v884 != v885 {
		v902 = v879
		v903 = v880
		v904 = v881
		goto L238
	} else {
		goto L245
	}
L244:
	;
	v895 = v890
	v896 = v888
	v897 = v892
	goto L241
L245:
	;
	v887 = int32(4)
	v888 = v880 + v887
	v890 = v879 + v887
	v892 = v881 - v887
	if base.Ui32(int32(3)) < base.Ui32(v892) {
		v879 = v890
		v880 = v888
		v881 = v892
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v902 = v895
	v903 = v896
	v904 = v897
	goto L238
L248:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907))))
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v912 == v913 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v933 = v912 - v913
	goto L236
L250:
	;
	v915 = int32(1)
	v920 = v909 - v915
	if v920 != 0 {
		v907 = v907 + v915
		v908 = v908 + v915
		v909 = v920
		goto L248
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	goto L249
L253:
	;
	goto L237
L254:
	;
	v934 = v865 + v868
	if base.Ui32(int32(4)) <= base.Ui32(v867) {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	if v996 != 0 {
		goto L5
	} else {
		goto L273
	}
L256:
	;
	v996 = int32(0)
	goto L255
L257:
	;
	v970 = v965
	v971 = v966
	v972 = v967
	goto L267
L258:
	;
	if (v934|v866)&int32(3) != 0 {
		v965 = v934
		v966 = v866
		v967 = v867
		goto L257
	} else {
		goto L261
	}
L259:
	;
	v958 = v934
	v959 = v866
	v960 = v867
	goto L260
L260:
	;
	if v960 == int32(0) {
		goto L256
	} else {
		goto L266
	}
L261:
	;
	v942 = v934
	v943 = v866
	v944 = v867
	goto L262
L262:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
	if v947 != v948 {
		v965 = v942
		v966 = v943
		v967 = v944
		goto L257
	} else {
		goto L264
	}
L263:
	;
	v958 = v953
	v959 = v951
	v960 = v955
	goto L260
L264:
	;
	v950 = int32(4)
	v951 = v943 + v950
	v953 = v942 + v950
	v955 = v944 - v950
	if base.Ui32(int32(3)) < base.Ui32(v955) {
		v942 = v953
		v943 = v951
		v944 = v955
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v965 = v958
	v966 = v959
	v967 = v960
	goto L257
L267:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970))))
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	if v975 == v976 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v996 = v975 - v976
	goto L255
L269:
	;
	v978 = int32(1)
	v983 = v972 - v978
	if v983 != 0 {
		v970 = v970 + v978
		v971 = v971 + v978
		v972 = v983
		goto L267
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	goto L268
L272:
	;
	goto L256
L273:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v998 = F_pg_hmac_create(m, v997)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L26
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+156)) = int32(0)
	v1003 = l0 - int32(-64)
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1005 = F_pg_hmac_init(m, v998, v1003, v1004)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L26
	} else {
		goto L275
	}
L275:
	;
	if v1005 < int32(0) {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1010 = F_strlen(m, v1009)
	mBase = m.M
	if v998 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	if v1025 < int32(0) {
		goto L4
	} else {
		goto L284
	}
L278:
	;
	v1025 = int32(-1)
	goto L277
L279:
	;
	goto L280
L280:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1015 = F_pg_cryptohash_update(m, v1014, v1009, v1010)
	mBase = m.M
	if int32(0) <= v1015 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1025 = int32(0)
	goto L277
L282:
	;
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998)+8)) = int32(2)
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1022 = F_pg_cryptohash_error(m, v1021)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v998)+12)) = v1022
	v1025 = int32(-1)
	goto L277
L284:
	;
	if v998 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	if v1044 < int32(0) {
		goto L4
	} else {
		goto L292
	}
L286:
	;
	v1044 = int32(-1)
	goto L285
L287:
	;
	goto L288
L288:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1034 = F_pg_cryptohash_update(m, v1033, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1034 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1044 = int32(0)
	goto L285
L290:
	;
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998)+8)) = int32(2)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1041 = F_pg_cryptohash_error(m, v1040)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v998)+12)) = v1041
	v1044 = int32(-1)
	goto L285
L292:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v1048 = F_strlen(m, v1047)
	mBase = m.M
	if v998 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	if v1063 < int32(0) {
		goto L4
	} else {
		goto L300
	}
L294:
	;
	v1063 = int32(-1)
	goto L293
L295:
	;
	goto L296
L296:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1053 = F_pg_cryptohash_update(m, v1052, v1047, v1048)
	mBase = m.M
	if int32(0) <= v1053 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1063 = int32(0)
	goto L293
L298:
	;
	goto L299
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998)+8)) = int32(2)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1060 = F_pg_cryptohash_error(m, v1059)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v998)+12)) = v1060
	v1063 = int32(-1)
	goto L293
L300:
	;
	if v998 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	if v1082 < int32(0) {
		goto L4
	} else {
		goto L308
	}
L302:
	;
	v1082 = int32(-1)
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1072 = F_pg_cryptohash_update(m, v1071, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1072 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1082 = int32(0)
	goto L301
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998)+8)) = int32(2)
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1079 = F_pg_cryptohash_error(m, v1078)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v998)+12)) = v1079
	v1082 = int32(-1)
	goto L301
L308:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1086 = F_strlen(m, v1085)
	mBase = m.M
	if v998 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	if v1101 < int32(0) {
		goto L4
	} else {
		goto L316
	}
L310:
	;
	v1101 = int32(-1)
	goto L309
L311:
	;
	goto L312
L312:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1091 = F_pg_cryptohash_update(m, v1090, v1085, v1086)
	mBase = m.M
	if int32(0) <= v1091 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1101 = int32(0)
	goto L309
L314:
	;
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v998)+8)) = int32(2)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	v1098 = F_pg_cryptohash_error(m, v1097)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v998)+12)) = v1098
	v1101 = int32(-1)
	goto L309
L316:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1107 = F_pg_hmac_final(m, v998, v19+int32(192), v1106)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L26
	} else {
		goto L317
	}
L317:
	;
	if v1107 < int32(0) {
		goto L4
	} else {
		goto L318
	}
L318:
	;
	F_pg_hmac_free(m, v998)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L26
	} else {
		goto L319
	}
L319:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1113 <= int32(0) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1212 = v19 + int32(160)
	v1215 = F_scram_H(m, l0+int32(32), v1210, v1113, v1212, v19+int32(156))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L26
	} else {
		goto L329
	}
L321:
	;
	v1117 = l0 + int32(32)
	v1118 = int32(0)
	if v1113 != int32(1) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1133 = v1118
	v1136 = int32(0)
	goto L325
L323:
	;
	v1174 = v1118
	goto L324
L324:
	;
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(192)+v1174))))
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v1174))))
	v1190 = v1187 ^ v1189
	*(*uint8)(unsafe.Add(mBase, uint32(v1174+v1117))) = uint8(v1190)
	goto L320
L325:
	;
	v1144 = v19 + int32(192)
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144+v1133))))
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v1133))))
	v1149 = v1146 ^ v1148
	*(*uint8)(unsafe.Add(mBase, uint32(v1133+v1117))) = uint8(v1149)
	v1152 = v1133 | int32(1)
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152+v1144))))
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v1152))))
	v1158 = v1155 ^ v1157
	*(*uint8)(unsafe.Add(mBase, uint32(v1117+v1152))) = uint8(v1158)
	v1160 = int32(2)
	v1161 = v1133 + v1160
	v1163 = v1136 + v1160
	if v1163 != v1113&int32(2147483646) {
		v1133 = v1161
		v1136 = v1163
		goto L325
	} else {
		goto L327
	}
L326:
	;
	if v1113&int32(1) == int32(0) {
		goto L320
	} else {
		goto L328
	}
L327:
	;
	goto L326
L328:
	;
	v1174 = v1161
	goto L324
L329:
	;
	if v1215 < int32(0) {
		goto L3
	} else {
		goto L330
	}
L330:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v1219) {
		goto L334
	} else {
		goto L335
	}
L331:
	;
	if v1281 != 0 {
		goto L31
	} else {
		goto L349
	}
L332:
	;
	v1281 = int32(0)
	goto L331
L333:
	;
	v1255 = v1250
	v1256 = v1251
	v1257 = v1252
	goto L343
L334:
	;
	if (v1212|v1003)&int32(3) != 0 {
		v1250 = v1212
		v1251 = v1003
		v1252 = v1219
		goto L333
	} else {
		goto L337
	}
L335:
	;
	v1243 = v1212
	v1244 = v1003
	v1245 = v1219
	goto L336
L336:
	;
	if v1245 == int32(0) {
		goto L332
	} else {
		goto L342
	}
L337:
	;
	v1227 = v1212
	v1228 = v1003
	v1229 = v1219
	goto L338
L338:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1228)))
	if v1232 != v1233 {
		v1250 = v1227
		v1251 = v1228
		v1252 = v1229
		goto L333
	} else {
		goto L340
	}
L339:
	;
	v1243 = v1238
	v1244 = v1236
	v1245 = v1240
	goto L336
L340:
	;
	v1235 = int32(4)
	v1236 = v1228 + v1235
	v1238 = v1227 + v1235
	v1240 = v1229 - v1235
	if base.Ui32(int32(3)) < base.Ui32(v1240) {
		v1227 = v1238
		v1228 = v1236
		v1229 = v1240
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	v1250 = v1243
	v1251 = v1244
	v1252 = v1245
	goto L333
L343:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1255))))
	v1261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1256))))
	if v1260 == v1261 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1281 = v1260 - v1261
	goto L331
L345:
	;
	v1263 = int32(1)
	v1268 = v1257 - v1263
	if v1268 != 0 {
		v1255 = v1255 + v1263
		v1256 = v1256 + v1263
		v1257 = v1268
		goto L343
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	goto L344
L348:
	;
	goto L332
L349:
	;
	v1282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+192)))
	if v1282 != 0 {
		goto L31
	} else {
		goto L350
	}
L350:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1284 = F_pg_hmac_create(m, v1283)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L26
	} else {
		goto L351
	}
L351:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1289 = F_pg_hmac_init(m, v1284, l0+int32(96), v1288)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L26
	} else {
		goto L352
	}
L352:
	;
	if v1289 < int32(0) {
		goto L2
	} else {
		goto L353
	}
L353:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1294 = F_strlen(m, v1293)
	mBase = m.M
	if v1284 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	if v1309 < int32(0) {
		goto L2
	} else {
		goto L361
	}
L355:
	;
	v1309 = int32(-1)
	goto L354
L356:
	;
	goto L357
L357:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1299 = F_pg_cryptohash_update(m, v1298, v1293, v1294)
	mBase = m.M
	if int32(0) <= v1299 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1309 = int32(0)
	goto L354
L359:
	;
	goto L360
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+8)) = int32(2)
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1306 = F_pg_cryptohash_error(m, v1305)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+12)) = v1306
	v1309 = int32(-1)
	goto L354
L361:
	;
	if v1284 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	if v1328 < int32(0) {
		goto L2
	} else {
		goto L369
	}
L363:
	;
	v1328 = int32(-1)
	goto L362
L364:
	;
	goto L365
L365:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1318 = F_pg_cryptohash_update(m, v1317, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1318 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1328 = int32(0)
	goto L362
L367:
	;
	goto L368
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+8)) = int32(2)
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1325 = F_pg_cryptohash_error(m, v1324)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+12)) = v1325
	v1328 = int32(-1)
	goto L362
L369:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v1332 = F_strlen(m, v1331)
	mBase = m.M
	if v1284 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	if v1347 < int32(0) {
		goto L2
	} else {
		goto L377
	}
L371:
	;
	v1347 = int32(-1)
	goto L370
L372:
	;
	goto L373
L373:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1337 = F_pg_cryptohash_update(m, v1336, v1331, v1332)
	mBase = m.M
	if int32(0) <= v1337 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1347 = int32(0)
	goto L370
L375:
	;
	goto L376
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+8)) = int32(2)
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1344 = F_pg_cryptohash_error(m, v1343)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+12)) = v1344
	v1347 = int32(-1)
	goto L370
L377:
	;
	if v1284 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	if v1366 < int32(0) {
		goto L2
	} else {
		goto L385
	}
L379:
	;
	v1366 = int32(-1)
	goto L378
L380:
	;
	goto L381
L381:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1356 = F_pg_cryptohash_update(m, v1355, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1356 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1366 = int32(0)
	goto L378
L383:
	;
	goto L384
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+8)) = int32(2)
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1363 = F_pg_cryptohash_error(m, v1362)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+12)) = v1363
	v1366 = int32(-1)
	goto L378
L385:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1370 = F_strlen(m, v1369)
	mBase = m.M
	if v1284 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	if v1385 < int32(0) {
		goto L2
	} else {
		goto L393
	}
L387:
	;
	v1385 = int32(-1)
	goto L386
L388:
	;
	goto L389
L389:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1375 = F_pg_cryptohash_update(m, v1374, v1369, v1370)
	mBase = m.M
	if int32(0) <= v1375 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	v1385 = int32(0)
	goto L386
L391:
	;
	goto L392
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+8)) = int32(2)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1382 = F_pg_cryptohash_error(m, v1381)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+12)) = v1382
	v1385 = int32(-1)
	goto L386
L393:
	;
	v1389 = v19 + int32(192)
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1391 = F_pg_hmac_final(m, v1284, v1389, v1390)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L26
	} else {
		goto L394
	}
L394:
	;
	if v1391 < int32(0) {
		goto L2
	} else {
		goto L395
	}
L395:
	;
	F_pg_hmac_free(m, v1284)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L26
	} else {
		goto L396
	}
L396:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1399 = int32(2)
	v1402 = base.I32_div_s(v1398+v1399, int32(3))
	v1404 = v1402 << (uint(v1399) % 32)
	goto L397
L397:
	;
	v1407 = F_palloc(m, v1404+int32(1))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L26
	} else {
		goto L398
	}
L398:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) < v1409 {
		goto L402
	} else {
		goto L403
	}
L399:
	;
	if v1521 < int32(0) {
		goto L1
	} else {
		goto L420
	}
L400:
	;
	if v1404 != 0 {
		goto L417
	} else {
		goto L418
	}
L401:
	;
	if v1404 < v1463-v1407+int32(4) {
		goto L400
	} else {
		goto L413
	}
L402:
	;
	v1418 = v1389
	v1419 = int32(0)
	v1422 = v1407
	v1423 = int32(2)
	goto L405
L403:
	;
	v1474 = v1407
	goto L404
L404:
	;
	v1521 = v1474 - v1407
	goto L399
L405:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418))))
	v1429 = v1425<<(uint(v1423<<(uint(int32(3))%32))%32) | v1419
	if int32(0) < v1423 {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	if v1464 != int32(2) {
		goto L401
	} else {
		goto L412
	}
L407:
	;
	v1462 = v1429
	v1463 = v1422
	v1464 = v1423 - int32(1)
	goto L409
L408:
	;
	if v1404 < v1422-v1407+int32(4) {
		goto L400
	} else {
		goto L410
	}
L409:
	;
	v1466 = v1418 + int32(1)
	if base.Ui32(v1466) < base.Ui32(v1389+v1409) {
		v1418 = v1466
		v1419 = v1462
		v1422 = v1463
		v1423 = v1464
		goto L405
	} else {
		goto L411
	}
L410:
	;
	v1438 = int32(63)
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1429&v1438)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+3)) = uint8(v1440)
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1429)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422))) = uint8(v1444)
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1429)>>(uint(int32(6))%32))&v1438)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+2)) = uint8(v1450)
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1429)>>(uint(int32(12))%32))&v1438)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1422)+1)) = uint8(v1456)
	v1462 = int32(0)
	v1463 = v1422 + int32(4)
	v1464 = int32(2)
	goto L409
L411:
	;
	goto L406
L412:
	;
	v1474 = v1463
	goto L404
L413:
	;
	v1484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1462)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1463))) = uint8(v1484)
	v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1462)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+1)) = uint8(v1490)
	if v1464 == int32(0) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1462)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	v1500 = v1499
	goto L416
L415:
	;
	v1500 = int32(61)
	goto L416
L416:
	;
	v1501 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+3)) = uint8(v1501)
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+2)) = uint8(v1500)
	v1521 = v1463 + int32(4) - v1407
	goto L399
L417:
	;
	base.MemoryFill(m, v1407, int32(0), v1404)
	goto L419
L418:
	;
	goto L419
L419:
	;
	v1521 = int32(-1)
	goto L399
L420:
	;
	v1525 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1407+v1521))) = uint8(v1525)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v1407
	v1531 = F_psprintf(m, int32(_a_F_scram_exchange_16), v19+int32(144))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L26
	} else {
		goto L421
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1531
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
	v1560 = int32(1)
	v1565 = int32(1)
	goto L30
L422:
	;
	F_errmsg_internal(m, int32(_a_F_scram_exchange_17), int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L26
	} else {
		goto L423
	}
L423:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(457), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L26
	} else {
		goto L424
	}
L424:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L425:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1554 == int32(0) {
		v1560 = v1550
		v1565 = v1551
		goto L30
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1554
	v1560 = v1550
	v1565 = v1551
	goto L30
L427:
	;
	v1575 = F_strlen(m, v1574)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1575
	goto L429
L428:
	;
	goto L429
L429:
	;
	if v1560 == int32(0) {
		v1609 = v1565
		goto L22
	} else {
		goto L430
	}
L430:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1579 != int32(2) {
		v1609 = v1565
		goto L22
	} else {
		goto L431
	}
L431:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, _c_F_scram_exchange[7]))
	v1584 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+440)) = v1584
	v1586 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+432)) = v1586
	v1588 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+424)) = v1588
	v1590 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+416)) = v1590
	v1592 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+464)) = v1592
	v1594 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+456)) = v1594
	v1596 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+448)) = v1596
	v1598 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v1599 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1583)+480)) = uint8(v1599)
	*(*int64)(unsafe.Add(mBase, uint32(v1583)+472)) = v1598
	v1609 = v1565
	goto L22
L432:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L26
	} else {
		goto L433
	}
L433:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L26
	} else {
		goto L434
	}
L434:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_19), int32(0))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L26
	} else {
		goto L435
	}
L435:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(383), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L26
	} else {
		goto L436
	}
L436:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L437:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L26
	} else {
		goto L438
	}
L438:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L26
	} else {
		goto L439
	}
L439:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_20), int32(0))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L26
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(388), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L26
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L26
	} else {
		goto L443
	}
L443:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L26
	} else {
		goto L444
	}
L444:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_21), int32(0))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L26
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(996), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L26
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L26
	} else {
		goto L448
	}
L448:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L26
	} else {
		goto L449
	}
L449:
	;
	v1693 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+1)))
	F_sanitize_char_2(m, v1693)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L26
	} else {
		goto L450
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(_a_F_scram_exchange_3)
	F_errdetail(m, int32(_a_F_scram_exchange_22), v19+int32(48))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L26
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1004), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L26
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L26
	} else {
		goto L454
	}
L454:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L26
	} else {
		goto L455
	}
L455:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_21), int32(0))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L26
	} else {
		goto L456
	}
L456:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1018), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L26
	} else {
		goto L457
	}
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L26
	} else {
		goto L459
	}
L459:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L26
	} else {
		goto L460
	}
L460:
	;
	v1739 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+1)))
	F_sanitize_char_2(m, v1739)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L26
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(_a_F_scram_exchange_3)
	F_errdetail(m, int32(_a_F_scram_exchange_22), v19-int32(-64))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L26
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1034), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L26
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L26
	} else {
		goto L465
	}
L465:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L26
	} else {
		goto L466
	}
L466:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_23), int32(0))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L26
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1047), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L26
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L469:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L26
	} else {
		goto L470
	}
L470:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_24), int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L26
	} else {
		goto L471
	}
L471:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1075), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L26
	} else {
		goto L472
	}
L472:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L473:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L26
	} else {
		goto L474
	}
L474:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_25), int32(0))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L26
	} else {
		goto L475
	}
L475:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1096), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L26
	} else {
		goto L476
	}
L476:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L477:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L26
	} else {
		goto L478
	}
L478:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_26), int32(0))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L26
	} else {
		goto L479
	}
L479:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1110), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L26
	} else {
		goto L480
	}
L480:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L481:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L26
	} else {
		goto L482
	}
L482:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_27), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L26
	} else {
		goto L483
	}
L483:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1240), int32(_a_F_scram_exchange_28))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L26
	} else {
		goto L484
	}
L484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L485:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L26
	} else {
		goto L486
	}
L486:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_29), int32(0))
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L26
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1250), int32(_a_F_scram_exchange_28))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L26
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	F_errmsg_internal(m, int32(_a_F_scram_exchange_30), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L26
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1359), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L26
	} else {
		goto L491
	}
L491:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L492:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L26
	} else {
		goto L493
	}
L493:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_32), int32(0))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L26
	} else {
		goto L494
	}
L494:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1374), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L26
	} else {
		goto L495
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L26
	} else {
		goto L497
	}
L497:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L26
	} else {
		goto L498
	}
L498:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_33), int32(0))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L26
	} else {
		goto L499
	}
L499:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1393), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L26
	} else {
		goto L500
	}
L500:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L501:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L26
	} else {
		goto L502
	}
L502:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L26
	} else {
		goto L503
	}
L503:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_34), int32(0))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L26
	} else {
		goto L504
	}
L504:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1401), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L26
	} else {
		goto L505
	}
L505:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L506:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L26
	} else {
		goto L507
	}
L507:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_35), int32(0))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L26
	} else {
		goto L508
	}
L508:
	;
	F_errdetail(m, int32(_a_F_scram_exchange_36), int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L26
	} else {
		goto L509
	}
L509:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(421), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L26
	} else {
		goto L510
	}
L510:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L511:
	;
	if v998 == int32(0) {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v1983
	F_errmsg_internal(m, int32(_a_F_scram_exchange_37), v19+int32(96))
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L26
	} else {
		goto L525
	}
L513:
	;
	v1983 = int32(_a_F_scram_exchange_38)
	goto L512
L514:
	;
	goto L515
L515:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v998)+12))
	if v1968 != 0 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1980 = v1968
	goto L518
L517:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v998)+8))
	if v1972 == int32(2) {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	v1983 = v1980
	goto L512
L519:
	;
	v1975 = int32(_a_F_scram_exchange_39)
	goto L521
L520:
	;
	v1975 = int32(_a_F_scram_exchange_40)
	goto L521
L521:
	;
	if v1972 == int32(1) {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v1978 = int32(_a_F_scram_exchange_38)
	goto L524
L523:
	;
	v1978 = v1975
	goto L524
L524:
	;
	v1980 = v1978
	goto L518
L525:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1177), int32(_a_F_scram_exchange_41))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L26
	} else {
		goto L526
	}
L526:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L527:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v19)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v1999
	F_errmsg_internal(m, int32(_a_F_scram_exchange_42), v19+int32(112))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L26
	} else {
		goto L528
	}
L528:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1189), int32(_a_F_scram_exchange_41))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L26
	} else {
		goto L529
	}
L529:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L530:
	;
	if v1284 == int32(0) {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v2036
	F_errmsg_internal(m, int32(_a_F_scram_exchange_43), v19+int32(128))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L26
	} else {
		goto L544
	}
L532:
	;
	v2036 = int32(_a_F_scram_exchange_38)
	goto L531
L533:
	;
	goto L534
L534:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+12))
	if v2021 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v2033 = v2021
	goto L537
L536:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+8))
	if v2025 == int32(2) {
		goto L538
	} else {
		goto L539
	}
L537:
	;
	v2036 = v2033
	goto L531
L538:
	;
	v2028 = int32(_a_F_scram_exchange_39)
	goto L540
L539:
	;
	v2028 = int32(_a_F_scram_exchange_40)
	goto L540
L540:
	;
	if v2025 == int32(1) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2031 = int32(_a_F_scram_exchange_38)
	goto L543
L542:
	;
	v2031 = v2028
	goto L543
L543:
	;
	v2033 = v2031
	goto L537
L544:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1435), int32(_a_F_scram_exchange_44))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L26
	} else {
		goto L545
	}
L545:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L546:
	;
	F_errmsg_internal(m, int32(_a_F_scram_exchange_45), int32(0))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L26
	} else {
		goto L547
	}
L547:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1447), int32(_a_F_scram_exchange_44))
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L26
	} else {
		goto L548
	}
L548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_scram_get_mechanisms(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	F_appendStringInfoString(m, l1, int32(_a_F_scram_get_mechanisms_0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_appendStringInfoChar(m, l1, int32(0))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_script_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v16
	v18 = F_geterrposition(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v282 = int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v287 = F_strlen(m, v283)
	mBase = m.M
	v294 = v287 + v282
	goto L77
L2:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v275 = v270
	v281 = base.B2i32(int32(0) <= v270)
	goto L1
L3:
	;
	return
L4:
	;
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if (base.B2i32(v16 <= int32(0))|base.B2i32(v18 <= v14+v16))&base.B2i32(base.Ui32(v14) <= base.Ui32(v18)) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v174 = int32(0)
	if v14 < v174 {
		v275 = v14
		v281 = v174
		goto L1
	} else {
		goto L53
	}
L8:
	;
	v82 = v11 + int32(44)
	v84 = v11 + int32(40)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v89 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v29 = F_strlen(m, v13)
	mBase = m.M
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v30
	if v29 <= v30 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = int32(0)
	v42 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v13))))
	if v46 != int32(59) {
		v68 = v40
		v69 = v42
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v71 = v68 + int32(1)
	if v71 < v29 {
		v40 = v71
		v42 = v69
		goto L11
	} else {
		goto L23
	}
L14:
	;
	v50 = v40 + int32(1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v13))))
	if v52 == int32(13) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v55 = v50
	goto L17
L16:
	;
	v55 = v40
	goto L17
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v55)+1)))
	if v57 != int32(10) {
		v68 = v55
		v69 = v42
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v61 = v55 + int32(2)
	if v61 < v18 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v61
	v68 = v55
	v69 = v61
	goto L13
L20:
	;
	goto L21
L21:
	;
	if v61 <= v18 {
		v68 = v55
		v69 = v42
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v61 - v42
	goto L8
L23:
	;
	goto L12
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v158 = F_errposition(m, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L43
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v149
	goto L24
L26:
	;
	v107 = v103
	v110 = v104
	v111 = v105
	goto L33
L27:
	;
	v100 = F_strlen(m, v97)
	mBase = m.M
	if v100 <= int32(0) {
		v146 = v97
		v149 = v100
		v150 = v99
		goto L25
	} else {
		goto L32
	}
L28:
	;
	v97 = v13
	v99 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v93 = v13 + v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if int32(0) < v94 {
		v103 = v93
		v104 = v94
		v105 = v89
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v97 = v93
	v99 = v89
	goto L27
L32:
	;
	v103 = v97
	v104 = v100
	v105 = v99
	goto L26
L33:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107))))
	v115 = F_scanner_isspace(m, v114)
	mBase = m.M
	if v115 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v146 = v140
	v149 = int32(0)
	v150 = v104 + v105
	goto L25
L35:
	;
	v121 = v110
	goto L38
L36:
	;
	goto L37
L37:
	;
	v137 = int32(1)
	v140 = v107 + v137
	if v137 < v110 {
		v107 = v140
		v110 = v110 - v137
		v111 = v111 + v137
		goto L33
	} else {
		goto L42
	}
L38:
	;
	v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107+v121-int32(1)))))
	v129 = F_scanner_isspace(m, v128)
	mBase = m.M
	if v129 == int32(0) {
		v146 = v107
		v149 = v121
		v150 = v111
		goto L25
	} else {
		goto L40
	}
L39:
	;
	v146 = v107
	v149 = int32(0)
	v150 = v111
	goto L25
L40:
	;
	v132 = int32(1)
	if v132 < v121 {
		v121 = v121 - v132
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L34
L43:
	;
	v160 = v18 - v155
	if v160 < v156 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v162 = v160
	goto L46
L45:
	;
	v162 = v156
	goto L46
L46:
	;
	v163 = int32(0)
	if v163 <= v160 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = v162
	goto L49
L48:
	;
	v166 = v163
	goto L49
L49:
	;
	F_internalerrposition(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v170 = F_pnstrdup(m, v146, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v172 = F_internalerrquery(m, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L2
L53:
	;
	v178 = v11 + int32(44)
	v180 = v11 + int32(40)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v185 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L73
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v245
	goto L54
L56:
	;
	v203 = v199
	v206 = v200
	v207 = v201
	goto L63
L57:
	;
	v196 = F_strlen(m, v193)
	mBase = m.M
	if v196 <= int32(0) {
		v242 = v193
		v245 = v196
		v246 = v195
		goto L55
	} else {
		goto L62
	}
L58:
	;
	v193 = v13
	v195 = int32(0)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v189 = v13 + v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if int32(0) < v190 {
		v199 = v189
		v200 = v190
		v201 = v185
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v193 = v189
	v195 = v185
	goto L57
L62:
	;
	v199 = v193
	v200 = v196
	v201 = v195
	goto L56
L63:
	;
	v210 = int32(*(*int8)(unsafe.Add(mBase, uint32(v203))))
	v211 = F_scanner_isspace(m, v210)
	mBase = m.M
	if v211 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v242 = v236
	v245 = int32(0)
	v246 = v200 + v201
	goto L55
L65:
	;
	v217 = v206
	goto L68
L66:
	;
	goto L67
L67:
	;
	v233 = int32(1)
	v236 = v203 + v233
	if v233 < v206 {
		v203 = v236
		v206 = v206 - v233
		v207 = v207 + v233
		goto L63
	} else {
		goto L72
	}
L68:
	;
	v224 = int32(*(*int8)(unsafe.Add(mBase, uint32(v203+v217-int32(1)))))
	v225 = F_scanner_isspace(m, v224)
	mBase = m.M
	if v225 == int32(0) {
		v242 = v203
		v245 = v217
		v246 = v207
		goto L55
	} else {
		goto L70
	}
L69:
	;
	v242 = v203
	v245 = int32(0)
	v246 = v207
	goto L55
L70:
	;
	v228 = int32(1)
	if v228 < v217 {
		v217 = v217 - v228
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L64
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v242
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v255
	F_errcontext_msg(m, int32(_a_F_script_error_callback_0), v11+int32(32))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	goto L2
L75:
	;
	if v306 != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v296 = int32(0)
	if v294 == v296 {
		v306 = v296
		goto L76
	} else {
		goto L79
	}
L78:
	;
	v306 = v301
	goto L76
L79:
	;
	v300 = v294 - int32(1)
	v301 = v283 + v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v302 != int32(47) {
		v294 = v300
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v309 = v306 + int32(1)
	goto L83
L82:
	;
	v309 = v283
	goto L83
L83:
	;
	if v281 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	m.G0 = v11 + int32(48)
	return
L85:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v311 == int32(0) {
		v339 = v282
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L96
	}
L88:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L3
	} else {
		goto L94
	}
L89:
	;
	v316 = v275
	v317 = v310
	v319 = v282
	goto L90
L90:
	;
	v323 = v316 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v323
	if v323 < int32(0) {
		v339 = v319
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v339 = v330
	goto L88
L92:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	v330 = v319 + base.B2i32(v327 == int32(10))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	if v331 != 0 {
		v316 = v323
		v317 = v317 + int32(1)
		v319 = v330
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v309
	F_errcontext_msg(m, int32(_a_F_script_error_callback_1), v11)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	goto L84
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v309
	F_errcontext_msg(m, int32(_a_F_script_error_callback_2), v11+int32(16))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	goto L84
}
func F_secure_loaded_verify_locations(m *base.Module) int32 {
	return int32(0)
}
func F_secure_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_ProcessClientReadInterrupt(m, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if int32(0) < v15 {
		v82 = v15
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_ProcessClientReadInterrupt(m, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L32
	}
L4:
	;
	if base.Ui32(l2) < base.Ui32(v82) {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = F_pgmem_recv(m, v23, l1, l2)
	mBase = m.M
	if int32(0) <= v24 {
		v99 = v24
		goto L3
	} else {
		goto L8
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v27 != 0 {
		v99 = v24
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[0]))
	if v29 != int32(6) {
		v99 = v24
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[1]))
	v34 = int32(0)
	F_ModifyWaitEvent(m, v33, v34, int32(2), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[1]))
	v44 = F_WaitEventSetWait(m, v40, int32(-1), v8, int32(1), int32(100663296))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v46&int32(16) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v46&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	goto L7
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if v60 <= int32(0) {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	F_ProcessClientReadInterrupt(m, int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v82 = v60
	goto L4
L22:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_secure_read_0), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_secure_read_1), int32(241), int32(_a_F_secure_read_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v85 = l2
	goto L28
L27:
	;
	v85 = v82
	goto L28
L28:
	;
	if v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+512))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	base.MemoryCopy(m, l1, v86+v87, v85)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v90 + v85
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v93 - v85
	v99 = v85
	goto L3
L32:
	;
	m.G0 = v8 + int32(16)
	return v99
}
func F_serialize_deflist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	F_initStringInfo(m, v10+int32(16))
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
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v137 = F_cstring_to_text_with_len(m, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = int32(0)
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = v30 + v27<<(uint(int32(2))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = F_defGetString(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v38 = F_quote_identifier(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
	v42 = v10 + int32(16)
	F_appendStringInfo(m, v42, int32(_a_F_serialize_deflist_0), v10)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if base.Ui32(v47-int32(465)) < base.Ui32(int32(2)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v33+int32(4)) < base.Ui32(v113+v114<<(uint(int32(2))%32)) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	F_appendStringInfoString(m, v42, v35)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v54 = int32(92)
	v55 = F___strchrnul(m, v35, v54)
	mBase = m.M
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v57 == v54 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L11
L16:
	;
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v61 = v55
	goto L19
L18:
	;
	v61 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(69))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(39))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v75 = v35
	goto L25
L25:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(39))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	if base.B2i32(v79 != int32(92))&base.B2i32(v79 != int32(39)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	F_appendStringInfoChar(m, v10+int32(16), base.I32_extend8_s(v79))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_appendStringInfoChar(m, v10+int32(16), base.I32_extend8_s(v79))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v75 = v75 + int32(1)
	goto L25
L35:
	;
	goto L11
L36:
	;
	F_appendStringInfoString(m, v10+int32(16), int32(_a_F_serialize_deflist_1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v125 = v27 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v125 < v126 {
		v27 = v125
		goto L6
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L7
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	m.G0 = v10 + int32(32)
	return v137
}
func F_setRuleCheckAsUser_Query(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v12 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v40 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v3
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v21<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = l1
	v31 = v21 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v31 < v32 {
		v21 = v31
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v74 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v50 = int32(0)
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v50<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v58 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	F_setRuleCheckAsUser_Query(m, v61, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v65 = v50 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v65 < v66 {
		v50 = v65
		goto L10
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	goto L11
L18:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v105 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v77 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v78 <= v77 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v84 = v77
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	F_setRuleCheckAsUser_Query(m, v92, l1)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v96 = v84 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v96 < v97 {
		v84 = v96
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v112 = F_query_tree_walker_impl(m, l0, int32(1039), v9+int32(12), int32(3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	m.G0 = v9 + int32(16)
	return
L28:
	;
	goto L27
}
func F_setRuleCheckAsUser_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v3 == int32(67) {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_setRuleCheckAsUser_Query(m, l0, v6)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			v14 = F_expression_tree_walker_impl(m, l0, int32(1039), l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = v14
				return v17
			}
		}
	} else {
		v17 = int32(0)
		return v17
	}
}
func F_set_attnotnull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+118)))
	if v16 == int32(116) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L7
	} else {
		goto L50
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L46
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_CheckTableNotInUse(m, l1, int32(_a_F_set_attnotnull_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v32 = v25 + v26<<(uint(int32(4))%32) + l2*int32(100)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+11)))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v13 + int32(32)
	return
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32-int32(80))+86)))
	if v36 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v41 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L45
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v44 = F_SearchSysCacheCopyAttNum(m, v43, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v49 = int32(4)
	v52 = int32(118)
	*(*uint8)(unsafe.Add(mBase, uint32(v48+l2<<(uint(v49)%32))+15)) = uint8(v52)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
	v56 = v54 + v55
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56)+86)) = uint8(v57)
	F_CatalogTupleUpdate(m, v41, v44+v49, v44)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v63 = int32(0)
	if base.B2i32(l0 == v63)|base.B2i32(l3 == v63) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L42
	}
L19:
	;
	v69 = F_palloc0(m, int32(20))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(52)
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+74)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v56)+68))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
	v79 = F_makeVar(m, int32(1), v74, v75, v76, v77, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = int32(-1)
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+12)) = uint8(v83)
	v85 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v69
	v93 = F_list_make1_impl(m, v85, v13+int32(24))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v96 = F_ConstraintImpliedByRelConstraint(m, l1, v93, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	if v96 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v100 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v122 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	if v100 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v105 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v56 + v105
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v104 + v105
	F_errmsg_internal(m, int32(_a_F_set_attnotnull_1), v13+int32(16))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_set_attnotnull_2), int32(_a_F_set_attnotnull_3), int32(_a_F_set_attnotnull_4))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L18
L31:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+76)) = uint8(v192)
	goto L18
L32:
	;
	v160 = F_palloc0(m, int32(140))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L39
	}
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v125 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v132 = int32(0)
	goto L35
L35:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v128+v132<<(uint(int32(2))%32))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v144 == v121 {
		v185 = v143
		goto L31
	} else {
		goto L37
	}
L36:
	;
	goto L32
L37:
	;
	v147 = v132 + int32(1)
	if v125 != v147 {
		v132 = v147
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v121
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)) = uint8(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v169 = F_CreateTupleDescCopyConstr(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v160)+88)) = int64(0)
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+84)) = uint8(v174)
	v176 = int32(_a_F_set_attnotnull_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v160)+96)) = uint16(v176)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = F_lappend(m, v178, v160)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v179
	v185 = v160
	goto L31
L42:
	;
	F_relation_close(m, v41, int32(3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_pfree(m, v44)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L9
L45:
	;
	goto L9
L46:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_set_attnotnull_6), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_set_attnotnull_2), int32(_a_F_set_attnotnull_7), int32(_a_F_set_attnotnull_8))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg_internal(m, int32(_a_F_set_attnotnull_9), v13)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_set_attnotnull_2), int32(_a_F_set_attnotnull_10), int32(_a_F_set_attnotnull_11))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_cheapest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 float64
	_ = v166
	var v167 float64
	_ = v167
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 float64
	_ = v247
	var v248 float64
	_ = v248
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v356
	return
L2:
	;
	if v321 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 <= int32(0) {
		v352 = v2
		v356 = v2
		v357 = v2
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L13
	} else {
		goto L144
	}
L6:
	;
	v19 = v2
	v20 = v2
	v23 = v2
	v24 = v2
	v26 = v2
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v26<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L2
L9:
	;
	v330 = v26 + int32(1)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v330 < v331 {
		v19 = v321
		v20 = v322
		v23 = v325
		v24 = v326
		v26 = v330
		goto L7
	} else {
		goto L143
	}
L10:
	;
	v35 = F_lappend(m, v24, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if v19 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L13:
	;
	return
L14:
	;
	if v19 != 0 {
		v321 = v19
		v322 = v20
		v325 = v23
		v326 = v35
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v20 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v321 = int32(0)
	v322 = v33
	v325 = v23
	v326 = v35
	goto L9
L17:
	;
	goto L18
L18:
	;
	v40 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = v43
	goto L21
L20:
	;
	v44 = v40
	goto L21
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v45 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = v46
	goto L24
L23:
	;
	v47 = v40
	goto L24
L24:
	;
	v48 = int32(0)
	if v44 == v48 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v321 = v48
	v322 = v33
	v325 = v23
	v326 = v35
	goto L9
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v143 != v144 {
		goto L62
	} else {
		goto L63
	}
L27:
	;
	switch v142 {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		v321 = v48
		v322 = v20
		v325 = v23
		v326 = v35
		goto L9
	}
L28:
	;
	v142 = base.B2i32(v47 != int32(0))
	goto L27
L29:
	;
	goto L30
L30:
	;
	if v47 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v142 = int32(2)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v65 < v66 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v68 = v65
	goto L36
L35:
	;
	v68 = v66
	goto L36
L36:
	;
	if v68 <= int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v71 = int32(1)
	goto L39
L38:
	;
	v71 = v68
	goto L39
L39:
	;
	v72 = int32(8)
	v76 = int32(0)
	v78 = v76
	v79 = v76
	goto L42
L40:
	;
	v142 = int32(3)
	goto L27
L41:
	;
	v142 = v129
	goto L27
L42:
	;
	v89 = v79 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v44+v72+v89)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+(v47+v72))))
	if v91&(v93^int32(-1)) != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v66 < v65 {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v115 = v79 + int32(1)
	if v115 != v71 {
		v78 = v113
		v79 = v115
		goto L42
	} else {
		goto L51
	}
L45:
	;
	if base.B2i32(v78 == int32(1))|v93&(v91^int32(-1)) != 0 {
		v129 = int32(3)
		goto L41
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v93&(v91^int32(-1)) == int32(0) {
		v113 = v78
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v113 = int32(2)
	goto L44
L49:
	;
	if v78 == int32(2) {
		goto L40
	} else {
		goto L50
	}
L50:
	;
	v113 = int32(1)
	goto L44
L51:
	;
	goto L43
L52:
	;
	if v113 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v66 <= v65 {
		v129 = v113
		goto L41
	} else {
		goto L58
	}
L55:
	;
	v122 = int32(3)
	goto L57
L56:
	;
	v122 = int32(2)
	goto L57
L57:
	;
	v142 = v122
	goto L27
L58:
	;
	if v113 == int32(2) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v128 = int32(3)
	goto L61
L60:
	;
	v128 = int32(1)
	goto L61
L61:
	;
	v129 = v128
	goto L41
L62:
	;
	if v144 <= v143 {
		v321 = v48
		v322 = v20
		v325 = v23
		v326 = v35
		goto L9
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v147 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v20)+56))
	if base.F64_lt(v147, v148) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v321 = v48
	v322 = v33
	v325 = v23
	v326 = v35
	goto L9
L66:
	;
	v321 = v48
	v322 = v33
	v325 = v23
	v326 = v35
	goto L9
L67:
	;
	goto L68
L68:
	;
	if base.F64_gt(v147, v148) != 0 {
		v321 = v48
		v322 = v20
		v325 = v23
		v326 = v35
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v20)+48))
	if base.F64_lt(v151, v152) == int32(0) {
		v321 = v48
		v322 = v20
		v325 = v23
		v326 = v35
		goto L9
	} else {
		goto L70
	}
L70:
	;
	goto L25
L71:
	;
	v321 = v33
	v322 = v20
	v325 = v33
	v326 = v24
	goto L9
L72:
	;
	goto L73
L73:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	if v162 != v163 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	if v243 != v244 {
		goto L110
	} else {
		goto L111
	}
L75:
	;
	v240 = v33
	goto L74
L76:
	;
	if v163 <= v162 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v166 = *(*float64)(unsafe.Add(mBase, uint32(v23)+48))
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	if base.F64_lt(v166, v167) != 0 {
		v240 = v23
		goto L74
	} else {
		goto L80
	}
L79:
	;
	v240 = v23
	goto L74
L80:
	;
	if base.F64_gt(v166, v167) != 0 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v23)+56))
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	if base.F64_lt(v170, v171) != 0 {
		v240 = v23
		goto L74
	} else {
		goto L82
	}
L82:
	;
	if base.F64_gt(v170, v171) != 0 {
		goto L75
	} else {
		goto L83
	}
L83:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	if v174 == v175 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v235 != int32(2) {
		v240 = v23
		goto L74
	} else {
		goto L108
	}
L85:
	;
	v235 = int32(0)
	goto L84
L86:
	;
	goto L87
L87:
	;
	v184 = int32(0)
	goto L90
L88:
	;
	if v224 != 0 {
		goto L105
	} else {
		goto L106
	}
L89:
	;
	v219 = int32(0)
	if v206 != 0 {
		goto L102
	} else {
		goto L103
	}
L90:
	;
	v188 = int32(0)
	if v174 == v188 {
		v198 = v188
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v235 = int32(3)
	goto L84
L92:
	;
	if v175 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v192 <= v184 {
		v198 = int32(0)
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v198 = v194 + v184<<(uint(int32(2))%32)
	goto L92
L95:
	;
	v204 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	if base.B2i32(v198 == v204)|base.B2i32(v206 == v204) != 0 {
		goto L89
	} else {
		goto L100
	}
L96:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v184 < v199 {
		goto L95
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v201 = int32(0)
	v224 = base.B2i32(v198 == v201)
	v226 = v201
	goto L88
L99:
	;
	goto L98
L100:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v206+v184<<(uint(int32(2))%32))))
	if v214 == v216 {
		v184 = v184 + int32(1)
		goto L90
	} else {
		goto L101
	}
L101:
	;
	goto L91
L102:
	;
	v223 = int32(2)
	goto L104
L103:
	;
	v223 = v219
	goto L104
L104:
	;
	v224 = base.B2i32(v198 == v219)
	v226 = v223
	goto L88
L105:
	;
	v228 = v226
	goto L107
L106:
	;
	v228 = int32(1)
	goto L107
L107:
	;
	v235 = v228
	goto L84
L108:
	;
	goto L75
L109:
	;
	v321 = v33
	v322 = v20
	v325 = v240
	v326 = v24
	goto L9
L110:
	;
	if v244 <= v243 {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v247 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
	v248 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	if base.F64_lt(v247, v248) != 0 {
		v321 = v19
		v322 = v20
		v325 = v240
		v326 = v24
		goto L9
	} else {
		goto L114
	}
L113:
	;
	v321 = v19
	v322 = v20
	v325 = v240
	v326 = v24
	goto L9
L114:
	;
	if base.F64_gt(v247, v248) != 0 {
		goto L109
	} else {
		goto L115
	}
L115:
	;
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
	v252 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	if base.F64_lt(v251, v252) != 0 {
		v321 = v19
		v322 = v20
		v325 = v240
		v326 = v24
		goto L9
	} else {
		goto L116
	}
L116:
	;
	if base.F64_gt(v251, v252) != 0 {
		goto L109
	} else {
		goto L117
	}
L117:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	if v255 == v256 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v316 != int32(2) {
		v321 = v19
		v322 = v20
		v325 = v240
		v326 = v24
		goto L9
	} else {
		goto L142
	}
L119:
	;
	v316 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v265 = int32(0)
	goto L124
L122:
	;
	if v305 != 0 {
		goto L139
	} else {
		goto L140
	}
L123:
	;
	v300 = int32(0)
	if v287 != 0 {
		goto L136
	} else {
		goto L137
	}
L124:
	;
	v269 = int32(0)
	if v255 == v269 {
		v279 = v269
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v316 = int32(3)
	goto L118
L126:
	;
	if v256 != 0 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v273 <= v265 {
		v279 = int32(0)
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v279 = v275 + v265<<(uint(int32(2))%32)
	goto L126
L129:
	;
	v285 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	if base.B2i32(v279 == v285)|base.B2i32(v287 == v285) != 0 {
		goto L123
	} else {
		goto L134
	}
L130:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v265 < v280 {
		goto L129
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v282 = int32(0)
	v305 = base.B2i32(v279 == v282)
	v307 = v282
	goto L122
L133:
	;
	goto L132
L134:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v287+v265<<(uint(int32(2))%32))))
	if v295 == v297 {
		v265 = v265 + int32(1)
		goto L124
	} else {
		goto L135
	}
L135:
	;
	goto L125
L136:
	;
	v304 = int32(2)
	goto L138
L137:
	;
	v304 = v300
	goto L138
L138:
	;
	v305 = base.B2i32(v279 == v300)
	v307 = v304
	goto L122
L139:
	;
	v309 = v307
	goto L141
L140:
	;
	v309 = int32(1)
	goto L141
L141:
	;
	v316 = v309
	goto L118
L142:
	;
	goto L109
L143:
	;
	goto L8
L144:
	;
	F_errmsg_internal(m, int32(_a_F_set_cheapest_0), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_set_cheapest_1), int32(283), int32(_a_F_set_cheapest_2))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	v352 = v322
	v356 = v325
	v357 = v326
	goto L1
L148:
	;
	goto L149
L149:
	;
	v348 = F_lcons(m, v321, v326)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	v352 = v321
	v356 = v325
	v357 = v348
	goto L1
}
func F_set_errcontext_domain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_set_errcontext_domain[0]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_set_errcontext_domain[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_set_errcontext_domain_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_set_errcontext_domain_1), int32(1418), int32(_a_F_set_errcontext_domain_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if l0 != 0 {
			v26 = l0
		} else {
			v26 = int32(_a_F_set_errcontext_domain_3)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_c_F_set_errcontext_domain[1]))) = v26
		return
	}
}
func F_set_extra_field(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v5 == v9 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v11 {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	default:
		goto L4
	}
L4:
	;
	v25 = l0 + int32(56)
	goto L15
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v5 == v20 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v5 != v18 {
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v5 != v16 {
		goto L4
	} else {
		goto L12
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v5 != v14 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v5 != v12 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	goto L1
L12:
	;
	goto L1
L13:
	;
	goto L1
L14:
	;
	goto L4
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v28 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_pfree(m, v5)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	if v5 == v29 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	if v5 != v31 {
		v25 = v28
		goto L15
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	return
L23:
	;
	goto L1
}
func F_set_joinrel_size_estimates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v9 = F_calc_joinrel_size_estimate(m, l0, l1, l2, l3, v7, v8, l4, l5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v9
		return
	}
}
func F_set_limit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 float64
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_getTypeOutputInfo(m, int32(700), v6+int32(12), v6+int32(11))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v20 = F_OidOutputFunctionCall(m, v19, v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_SetConfigOption(m, int32(_a_F_set_limit_0), v20, int32(6), int32(13))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, _c_F_set_limit[0]))
				m.G0 = v6 + int32(16)
				return base.I32_reinterpret_f32(base.F32_demote_f64(v27))
			}
		}
	}
}
func F_set_spins_per_delay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_set_spins_per_delay[0])) = l0
	return
}
func F_setlocale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
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
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v401 int64
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v442 int64
	_ = v442
	var v445 int64
	_ = v445
	var v448 int64
	_ = v448
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v539 int64
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v797 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v797
L2:
	;
	if l0 == int32(6) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v590 = int32(0)
	v591 = int32(_a_F_setlocale_0)
	v596 = v3
	goto L197
L4:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l1 != 0 {
		goto L140
	} else {
		goto L141
	}
L7:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v20
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v26
	v29 = int32(0)
	v30 = l1
	goto L9
L8:
	;
	v797 = int32(0)
	goto L1
L9:
	;
	goto L15
L10:
	;
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_setlocale[3])) = v442
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_setlocale[4])) = v445
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_setlocale[5])) = v448
	goto L3
L11:
	;
	v133 = v122 - v30
	if v133 <= int32(23) {
		goto L34
	} else {
		goto L35
	}
L12:
	;
	goto L11
L13:
	;
	v112 = v107
	goto L30
L14:
	;
	v107 = v99
	goto L13
L15:
	;
	if v30&int32(3) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v45 = v30
	goto L21
L19:
	;
	v59 = v30
	goto L20
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = int32(-2139062144)
	if (int32(16843008)-v65|v65)&v68 != v68 {
		v99 = v59
		goto L14
	} else {
		goto L25
	}
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if base.B2i32(v50 == int32(0))|base.B2i32(int32(59) == v50) != 0 {
		v122 = v45
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v59 = v56
	goto L20
L23:
	;
	v56 = v45 + int32(1)
	if v56&int32(3) != 0 {
		v45 = v56
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v74 = v59
	v76 = v65
	goto L26
L26:
	;
	v80 = v76 ^ int32(993737531)
	v83 = int32(-2139062144)
	if (int32(16843008)-v80|v80)&v83 != v83 {
		v99 = v74
		goto L14
	} else {
		goto L28
	}
L27:
	;
	v107 = v89
	goto L13
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v89 = v74 + int32(4)
	v93 = int32(-2139062144)
	if (v87|(int32(16843008)-v87))&v93 == v93 {
		v74 = v89
		v76 = v87
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v114 == int32(0) {
		v122 = v112
		goto L12
	} else {
		goto L32
	}
L31:
	;
	v122 = v112
	goto L12
L32:
	;
	if v114 != int32(59) {
		v112 = v112 + int32(1)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v133) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v312 = v30
	goto L36
L36:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v316 != 0 {
		v330 = v11
		goto L88
	} else {
		goto L89
	}
L37:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v133))) = uint8(v306)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v310 != 0 {
		goto L84
	} else {
		goto L85
	}
L38:
	;
	if v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v142 = v11 + v133
	if (v11^v30)&int32(3) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	base.MemoryCopy(m, v11, v30, v133)
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L37
L44:
	;
	if base.Ui32(v274) < base.Ui32(v142) {
		goto L78
	} else {
		goto L79
	}
L45:
	;
	if v11&int32(3) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v142) < base.Ui32(int32(4)) {
		goto L69
	} else {
		goto L70
	}
L48:
	;
	v178 = v142 & int32(-4)
	if base.Ui32(v142) < base.Ui32(int32(64)) {
		v228 = v172
		v229 = v173
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v172 = v30
	v173 = v11
	goto L48
L50:
	;
	goto L51
L51:
	;
	if v133 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v172 = v30
	v173 = v11
	goto L48
L53:
	;
	goto L54
L54:
	;
	v155 = v30
	v156 = v11
	goto L55
L55:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v160)
	v162 = int32(1)
	v163 = v155 + v162
	v165 = v156 + v162
	if v165&int32(3) == int32(0) {
		v172 = v163
		v173 = v165
		goto L48
	} else {
		goto L57
	}
L56:
	;
	v172 = v163
	v173 = v165
	goto L48
L57:
	;
	if base.Ui32(v165) < base.Ui32(v142) {
		v155 = v163
		v156 = v165
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if base.Ui32(v178) <= base.Ui32(v229) {
		v273 = v228
		v274 = v229
		goto L44
	} else {
		goto L65
	}
L60:
	;
	v182 = v178 + int32(-64)
	if base.Ui32(v182) < base.Ui32(v173) {
		v228 = v172
		v229 = v173
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v185 = v172
	v186 = v173
	goto L62
L62:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+8)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+16)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v185)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v185)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+24)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v185)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+28)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v185)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+32)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v185)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+36)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v185)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+40)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v185)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+44)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v185)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+48)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+52)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v185)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+56)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v185)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+60)) = v220
	v222 = int32(-64)
	v223 = v185 - v222
	v225 = v186 - v222
	if base.Ui32(v225) <= base.Ui32(v182) {
		v185 = v223
		v186 = v225
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v228 = v223
	v229 = v225
	goto L59
L64:
	;
	goto L63
L65:
	;
	v235 = v228
	v236 = v229
	goto L66
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v240
	v242 = int32(4)
	v243 = v235 + v242
	v245 = v236 + v242
	if base.Ui32(v245) < base.Ui32(v178) {
		v235 = v243
		v236 = v245
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v273 = v243
	v274 = v245
	goto L44
L68:
	;
	goto L67
L69:
	;
	v273 = v30
	v274 = v11
	goto L44
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(v133) < base.Ui32(int32(4)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v273 = v30
	v274 = v11
	goto L44
L73:
	;
	goto L74
L74:
	;
	v254 = v30
	v255 = v11
	goto L75
L75:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	*(*uint8)(unsafe.Add(mBase, uint32(v255))) = uint8(v259)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)) = uint8(v261)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+2)) = uint8(v263)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+3)) = uint8(v265)
	v267 = int32(4)
	v268 = v254 + v267
	v270 = v255 + v267
	if base.Ui32(v270) <= base.Ui32(v142-int32(4)) {
		v254 = v268
		v255 = v270
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v273 = v268
	v274 = v270
	goto L44
L77:
	;
	goto L76
L78:
	;
	v280 = v273
	v281 = v274
	goto L81
L79:
	;
	goto L80
L80:
	;
	goto L37
L81:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	*(*uint8)(unsafe.Add(mBase, uint32(v281))) = uint8(v285)
	v287 = int32(1)
	v290 = v281 + v287
	if v290 != v142 {
		v280 = v280 + v287
		v281 = v290
		goto L81
	} else {
		goto L83
	}
L82:
	;
	goto L80
L83:
	;
	goto L82
L84:
	;
	v311 = v122 + int32(1)
	goto L86
L85:
	;
	v311 = v30
	goto L86
L86:
	;
	v312 = v311
	goto L36
L87:
	;
	if v428 == int32(-1) {
		goto L8
	} else {
		goto L137
	}
L88:
	;
	v333 = int32(0)
	goto L103
L89:
	;
	v318 = F_getenv(m, int32(_a_F_setlocale_1))
	mBase = m.M
	if v318 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v319 != 0 {
		v330 = v318
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v324 = F_getenv(m, v29*int32(12)+int32(_a_F_setlocale_2))
	mBase = m.M
	if v324 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v325 != 0 {
		v330 = v324
		goto L88
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v327 = F_getenv(m, int32(_a_F_setlocale_3))
	mBase = m.M
	if v327 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v328 != 0 {
		v330 = v327
		goto L88
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v330 = int32(_a_F_setlocale_4)
	goto L88
L101:
	;
	goto L100
L102:
	;
	v352 = int32(_a_F_setlocale_4)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if v353 == int32(46) {
		v360 = v352
		goto L113
	} else {
		goto L114
	}
L103:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+v333))))
	v338 = int32(0)
	if base.B2i32(v337 == v338)|base.B2i32(v337 == int32(47)) == v338 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v351 = v333
	goto L102
L105:
	;
	v345 = int32(23)
	v347 = v333 + int32(1)
	if v347 != v345 {
		v333 = v347
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	v351 = v345
	goto L102
L109:
	;
	v428 = v420
	goto L87
L110:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	if v381 != 0 {
		goto L124
	} else {
		goto L125
	}
L111:
	;
	if v29 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L112:
	;
	v366 = F_strcmp(m, v364, int32(_a_F_setlocale_4))
	mBase = m.M
	if v366 == int32(0) {
		v371 = v364
		goto L111
	} else {
		goto L118
	}
L113:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	if v361 == int32(0) {
		v371 = v360
		goto L111
	} else {
		goto L117
	}
L114:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+v351))))
	if v357 != 0 {
		v360 = v352
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if v353 != int32(67) {
		v364 = v330
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v360 = v330
	goto L113
L117:
	;
	v364 = v360
	goto L112
L118:
	;
	v370 = F_strcmp(m, v364, int32(_a_F_setlocale_5))
	mBase = m.M
	if v370 != 0 {
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v371 = v364
	goto L111
L120:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)))
	if v375 == int32(46) {
		v420 = int32(_a_F_setlocale_6)
		goto L109
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v428 = int32(0)
	goto L87
L123:
	;
	goto L122
L124:
	;
	v384 = v381
	goto L127
L125:
	;
	goto L126
L126:
	;
	v399 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v399 != 0 {
		goto L131
	} else {
		goto L132
	}
L127:
	;
	v389 = F_strcmp(m, v364, v384+int32(8))
	mBase = m.M
	if v389 == int32(0) {
		v420 = v384
		goto L109
	} else {
		goto L129
	}
L128:
	;
	goto L126
L129:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v384)+32))
	if v392 != 0 {
		v384 = v392
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v401 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = v401
	v404 = v399 + int32(8)
	v405 = F___memcpy(m, v404, v364, v351)
	mBase = m.M
	v407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v404+v351))) = uint8(v407)
	v409 = int32(_a_F_setlocale_7)
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v399)+32)) = v410
	*(*int32)(unsafe.Add(mBase, _c_F_setlocale[6])) = v399
	goto L133
L132:
	;
	goto L133
L133:
	;
	if v29|v399 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v417 = v399
	goto L136
L135:
	;
	v417 = int32(_a_F_setlocale_6)
	goto L136
L136:
	;
	v420 = v417
	goto L109
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)+v29<<(uint(int32(2))%32)))) = v428
	v438 = v29 + int32(1)
	if v438 != int32(6) {
		v29 = v438
		v30 = v312
		goto L9
	} else {
		goto L138
	}
L138:
	;
	goto L10
L139:
	;
	if v575 != 0 {
		goto L194
	} else {
		goto L195
	}
L140:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v454 != 0 {
		v468 = l1
		goto L144
	} else {
		goto L145
	}
L141:
	;
	goto L142
L142:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_setlocale[5])))
	v575 = v574
	goto L139
L143:
	;
	if v566 == int32(-1) {
		v797 = v3
		goto L1
	} else {
		goto L193
	}
L144:
	;
	v471 = int32(0)
	goto L159
L145:
	;
	v456 = F_getenv(m, int32(_a_F_setlocale_1))
	mBase = m.M
	if v456 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v457 != 0 {
		v468 = v456
		goto L144
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v462 = F_getenv(m, l0*int32(12)+int32(_a_F_setlocale_2))
	mBase = m.M
	if v462 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v463 != 0 {
		v468 = v462
		goto L144
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v465 = F_getenv(m, int32(_a_F_setlocale_3))
	mBase = m.M
	if v465 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	if v466 != 0 {
		v468 = v465
		goto L144
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v468 = int32(_a_F_setlocale_4)
	goto L144
L157:
	;
	goto L156
L158:
	;
	v490 = int32(_a_F_setlocale_4)
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v491 == int32(46) {
		v498 = v490
		goto L169
	} else {
		goto L170
	}
L159:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468+v471))))
	v476 = int32(0)
	if base.B2i32(v475 == v476)|base.B2i32(v475 == int32(47)) == v476 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v489 = v471
	goto L158
L161:
	;
	v483 = int32(23)
	v485 = v471 + int32(1)
	if v485 != v483 {
		v471 = v485
		goto L159
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	goto L160
L164:
	;
	v489 = v483
	goto L158
L165:
	;
	v566 = v558
	goto L143
L166:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	if v519 != 0 {
		goto L180
	} else {
		goto L181
	}
L167:
	;
	if l0 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L168:
	;
	v504 = F_strcmp(m, v502, int32(_a_F_setlocale_4))
	mBase = m.M
	if v504 == int32(0) {
		v509 = v502
		goto L167
	} else {
		goto L174
	}
L169:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)))
	if v499 == int32(0) {
		v509 = v498
		goto L167
	} else {
		goto L173
	}
L170:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468+v489))))
	if v495 != 0 {
		v498 = v490
		goto L169
	} else {
		goto L171
	}
L171:
	;
	if v491 != int32(67) {
		v502 = v468
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v498 = v468
	goto L169
L173:
	;
	v502 = v498
	goto L168
L174:
	;
	v508 = F_strcmp(m, v502, int32(_a_F_setlocale_5))
	mBase = m.M
	if v508 != 0 {
		goto L166
	} else {
		goto L175
	}
L175:
	;
	v509 = v502
	goto L167
L176:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
	if v513 == int32(46) {
		v558 = int32(_a_F_setlocale_6)
		goto L165
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v566 = int32(0)
	goto L143
L179:
	;
	goto L178
L180:
	;
	v522 = v519
	goto L183
L181:
	;
	goto L182
L182:
	;
	v537 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v537 != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v527 = F_strcmp(m, v502, v522+int32(8))
	mBase = m.M
	if v527 == int32(0) {
		v558 = v522
		goto L165
	} else {
		goto L185
	}
L184:
	;
	goto L182
L185:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v522)+32))
	if v530 != 0 {
		v522 = v530
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v539 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v537))) = v539
	v542 = v537 + int32(8)
	v543 = F___memcpy(m, v542, v502, v489)
	mBase = m.M
	v545 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v542+v489))) = uint8(v545)
	v547 = int32(_a_F_setlocale_7)
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v537)+32)) = v548
	*(*int32)(unsafe.Add(mBase, _c_F_setlocale[6])) = v537
	goto L189
L188:
	;
	goto L189
L189:
	;
	if l0|v537 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v555 = v537
	goto L192
L191:
	;
	v555 = int32(_a_F_setlocale_6)
	goto L192
L192:
	;
	v558 = v555
	goto L165
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_setlocale[5]))) = v566
	v575 = v566
	goto L139
L194:
	;
	v579 = v575 + int32(8)
	goto L196
L195:
	;
	v579 = int32(_a_F_setlocale_8)
	goto L196
L196:
	;
	v797 = v579
	goto L1
L197:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[5]))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v590<<(uint(int32(2))%32))+uint32(_c_F_setlocale[5])))
	if v602 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v788 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v788)
	if v783 != int32(6) {
		goto L250
	} else {
		goto L251
	}
L199:
	;
	v606 = v602 + int32(8)
	goto L201
L200:
	;
	v606 = int32(_a_F_setlocale_8)
	goto L201
L201:
	;
	v607 = F_strlen(m, v606)
	mBase = m.M
	if base.Ui32(int32(512)) <= base.Ui32(v607) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v777 = v591 + v607
	v778 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v778)
	v780 = int32(1)
	v783 = v596 + base.B2i32(v602 == v599)
	v785 = v590 + v780
	if v785 != int32(6) {
		v590 = v785
		v591 = v777 + v780
		v596 = v783
		goto L197
	} else {
		goto L249
	}
L203:
	;
	if v607 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v614 = v591 + v607
	if (v591^v606)&int32(3) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	base.MemoryCopy(m, v591, v606, v607)
	goto L208
L207:
	;
	goto L208
L208:
	;
	goto L202
L209:
	;
	if base.Ui32(v746) < base.Ui32(v614) {
		goto L243
	} else {
		goto L244
	}
L210:
	;
	if v591&int32(3) == int32(0) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L212
L212:
	;
	if base.Ui32(v614) < base.Ui32(int32(4)) {
		goto L234
	} else {
		goto L235
	}
L213:
	;
	v650 = v614 & int32(-4)
	if base.Ui32(v614) < base.Ui32(int32(64)) {
		v700 = v644
		v701 = v645
		goto L224
	} else {
		goto L225
	}
L214:
	;
	v644 = v606
	v645 = v591
	goto L213
L215:
	;
	goto L216
L216:
	;
	if v607 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v644 = v606
	v645 = v591
	goto L213
L218:
	;
	goto L219
L219:
	;
	v627 = v606
	v628 = v591
	goto L220
L220:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	*(*uint8)(unsafe.Add(mBase, uint32(v628))) = uint8(v632)
	v634 = int32(1)
	v635 = v627 + v634
	v637 = v628 + v634
	if v637&int32(3) == int32(0) {
		v644 = v635
		v645 = v637
		goto L213
	} else {
		goto L222
	}
L221:
	;
	v644 = v635
	v645 = v637
	goto L213
L222:
	;
	if base.Ui32(v637) < base.Ui32(v614) {
		v627 = v635
		v628 = v637
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	if base.Ui32(v650) <= base.Ui32(v701) {
		v745 = v700
		v746 = v701
		goto L209
	} else {
		goto L230
	}
L225:
	;
	v654 = v650 + int32(-64)
	if base.Ui32(v654) < base.Ui32(v645) {
		v700 = v644
		v701 = v645
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v657 = v644
	v658 = v645
	goto L227
L227:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	*(*int32)(unsafe.Add(mBase, uint32(v658))) = v662
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v657)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+8)) = v666
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v657)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+12)) = v668
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v657)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+16)) = v670
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v657)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+20)) = v672
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v657)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+24)) = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v657)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+28)) = v676
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v657)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+32)) = v678
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v657)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+36)) = v680
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v657)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+40)) = v682
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v657)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+44)) = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v657)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+48)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v657)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+52)) = v688
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v657)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+56)) = v690
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v657)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+60)) = v692
	v694 = int32(-64)
	v695 = v657 - v694
	v697 = v658 - v694
	if base.Ui32(v697) <= base.Ui32(v654) {
		v657 = v695
		v658 = v697
		goto L227
	} else {
		goto L229
	}
L228:
	;
	v700 = v695
	v701 = v697
	goto L224
L229:
	;
	goto L228
L230:
	;
	v707 = v700
	v708 = v701
	goto L231
L231:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	*(*int32)(unsafe.Add(mBase, uint32(v708))) = v712
	v714 = int32(4)
	v715 = v707 + v714
	v717 = v708 + v714
	if base.Ui32(v717) < base.Ui32(v650) {
		v707 = v715
		v708 = v717
		goto L231
	} else {
		goto L233
	}
L232:
	;
	v745 = v715
	v746 = v717
	goto L209
L233:
	;
	goto L232
L234:
	;
	v745 = v606
	v746 = v591
	goto L209
L235:
	;
	goto L236
L236:
	;
	if base.Ui32(v607) < base.Ui32(int32(4)) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v745 = v606
	v746 = v591
	goto L209
L238:
	;
	goto L239
L239:
	;
	v726 = v606
	v727 = v591
	goto L240
L240:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	*(*uint8)(unsafe.Add(mBase, uint32(v727))) = uint8(v731)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+1)) = uint8(v733)
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+2)) = uint8(v735)
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+3)) = uint8(v737)
	v739 = int32(4)
	v740 = v726 + v739
	v742 = v727 + v739
	if base.Ui32(v742) <= base.Ui32(v614-int32(4)) {
		v726 = v740
		v727 = v742
		goto L240
	} else {
		goto L242
	}
L241:
	;
	v745 = v740
	v746 = v742
	goto L209
L242:
	;
	goto L241
L243:
	;
	v752 = v745
	v753 = v746
	goto L246
L244:
	;
	goto L245
L245:
	;
	goto L202
L246:
	;
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752))))
	*(*uint8)(unsafe.Add(mBase, uint32(v753))) = uint8(v757)
	v759 = int32(1)
	v762 = v753 + v759
	if v762 != v614 {
		v752 = v752 + v759
		v753 = v762
		goto L246
	} else {
		goto L248
	}
L247:
	;
	goto L245
L248:
	;
	goto L247
L249:
	;
	goto L198
L250:
	;
	v793 = int32(_a_F_setlocale_0)
	goto L252
L251:
	;
	v793 = v606
	goto L252
L252:
	;
	v797 = v793
	goto L1
}
func F_sha224_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(2), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_sha512_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(5), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_shell_archive_configured(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[0]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[2])) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_shell_archive_configured_0)
		v20 = F_format_elog_string(m, int32(_a_F_shell_archive_configured_1), v5)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[3])) = v20
			m.G0 = v5 + int32(16)
			return base.B2i32(v9 != int32(0))
		}
	} else {
		m.G0 = v5 + int32(16)
		return base.B2i32(v9 != int32(0))
	}
}
func F_show_instrumentation_count(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v18 int32
	_ = v18
	var v20 float64
	_ = v20
	var v25 int32
	_ = v25
	var v30 float64
	_ = v30
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)))
	if v7 != int32(1) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		if v10 == int32(0) {
			return
		} else {
			v13 = *(*float64)(unsafe.Add(mBase, uint32(v10)+232))
			if l1 == int32(2) {
				v18 = int32(248)
			} else {
				v18 = int32(240)
			}
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v10+v18)))
			if base.F64_gt(v20, float64(0)) == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
				if v25 == int32(0) {
					return
				} else {
					v30 = float64(0)
					if base.F64_gt(v13, v30) != 0 {
						v33 = base.F64_div(v20, v13)
					} else {
						v33 = v30
					}
					F_ExplainPropertyFloat(m, l0, int32(0), v33, int32(0), l3)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v30 = float64(0)
				if base.F64_gt(v13, v30) != 0 {
					v33 = base.F64_div(v20, v13)
				} else {
					v33 = v30
				}
				F_ExplainPropertyFloat(m, l0, int32(0), v33, int32(0), l3)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_show_trgm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v221 int32
	_ = v221
	var v240 int32
	_ = v240
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(1)
	v24 = v19 + v23
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v29 = v27 & v23
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = v24
	goto L5
L4:
	;
	v30 = v19 + int32(4)
	goto L5
L5:
	;
	if v27 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_generate_trgm_only(m, v14+int32(4), v30, v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v36 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v47 = int32(1)
	if v29 != 0 {
		v57 = int32(base.Ui32(v27)>>(uint(v47)%32)) - v47
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v39 = int32(16)
	goto L12
L11:
	;
	v39 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = int32(4)
	goto L15
L14:
	;
	v46 = v39
	goto L15
L15:
	;
	v57 = v46
	goto L6
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+4)) = uint8(v63)
	if int32(2) <= v61 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v62 + int32(5)
	F_pg_qsort(m, v68, v61, int32(3), int32(_a_F_show_trgm_0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v126 = v61
	goto L20
L20:
	;
	v130 = v126*int32(12) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v130
	v132 = int32(2)
	v137 = base.I32_div_u_s(int32(base.Ui32(v130)>>(uint(v132)%32))-int32(5), int32(3))
	v142 = F_palloc(m, v137<<(uint(v132)%32)+int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L29
	}
L21:
	;
	v76 = int32(0)
	v77 = int32(1)
	goto L22
L22:
	;
	v86 = int32(3)
	v88 = v68 + v77*v86
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[0]))
	v94 = m.T0[v93].(func(*base.Module, int32, int32) int32)(m, v88, v68+v76*v86)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v126 = v108 + int32(1)
	goto L20
L24:
	;
	v111 = v77 + int32(1)
	if v111 != v61 {
		v76 = v108
		v77 = v111
		goto L22
	} else {
		goto L28
	}
L25:
	;
	if v94 == int32(0) {
		v108 = v76
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v99 = v76 + int32(1)
	if v99 == v77 {
		v108 = v77
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v103 = v68 + v99*int32(3)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+2)) = uint8(v104)
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
	*(*uint16)(unsafe.Add(mBase, uint32(v103))) = uint16(v106)
	v108 = v99
	goto L24
L28:
	;
	goto L23
L29:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if base.Ui32(int32(3)) <= base.Ui32(int32(base.Ui32(v144)>>(uint(int32(2))%32))-int32(5)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v156 = v62 + int32(5)
	v159 = int32(0)
	goto L33
L31:
	;
	v317 = int32(0)
	goto L32
L32:
	;
	v319 = F_construct_array_builtin(m, v142, v317, int32(25))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L57
	}
L33:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[1]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v167*int32(28))+uint32(_c_F_show_trgm[2])))
	goto L35
L34:
	;
	v317 = v305
	goto L32
L35:
	;
	if int32(4) <= v172 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[1]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177*int32(28))+uint32(_c_F_show_trgm[2])))
	goto L39
L37:
	;
	v188 = int32(16)
	goto L38
L38:
	;
	v189 = F_palloc(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v188 = v182*int32(3) + int32(4)
	goto L38
L40:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[1]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193*int32(28))+uint32(_c_F_show_trgm[2])))
	goto L43
L41:
	;
	v291 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v142+v159<<(uint(v291)%32)))) = v189
	v295 = int32(3)
	v298 = v159 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v305 = base.I32_div_u_s(int32(base.Ui32(v299)>>(uint(v291)%32))-int32(5), v295)
	if base.Ui32(v298) < base.Ui32(v305) {
		v156 = v156 + v295
		v159 = v298
		goto L33
	} else {
		goto L56
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(28)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)) = uint8(v283)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+5)) = uint8(v285)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+6)) = uint8(v287)
	goto L41
L43:
	;
	if v198 < int32(2) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v202 = base.I32_extend8_s(v201)
	if v202 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+2)))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v258 | (v259<<(uint(int32(8))%32) | v201<<(uint(int32(16))%32))
	v268 = v189 + int32(4)
	v271 = F_pg_snprintf(m, v268, int32(12), int32(_a_F_show_trgm_1), v14)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L55
	}
L46:
	;
	goto L47
L47:
	;
	if base.B2i32(base.B2i32(base.Ui32(v201-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v201|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v202 != int32(32)) != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v221 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v221 < int32(0) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L50
L50:
	;
	if base.B2i32(base.B2i32(base.Ui32(v221-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v221|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v221 != int32(32)) != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v240 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156)+2)))
	if v240 < int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L53
L53:
	;
	if base.B2i32(base.Ui32(v240-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v240|int32(32)-int32(97)) < base.Ui32(int32(26)))|base.B2i32(v240 == int32(32)) != 0 {
		goto L42
	} else {
		goto L54
	}
L54:
	;
	goto L45
L55:
	;
	v273 = F_strlen(m, v268)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v273<<(uint(int32(2))%32) + int32(16)
	goto L41
L56:
	;
	goto L34
L57:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if base.Ui32(int32(3)) <= base.Ui32(int32(base.Ui32(v321)>>(uint(int32(2))%32))-int32(5)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v331 = int32(0)
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_pfree(m, v142)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L65
	}
L61:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v142+v331<<(uint(int32(2))%32))))
	F_pfree(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v347 = v331 + int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v354 = base.I32_div_u_s(int32(base.Ui32(v348)>>(uint(int32(2))%32))-int32(5), int32(3))
	if base.Ui32(v347) < base.Ui32(v354) {
		v331 = v347
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	F_pfree(m, v62)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v371 != v19 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_pfree(m, v19)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	m.G0 = v14 + int32(16)
	return v319
L70:
	;
	goto L69
}
func F_sigUsr1Handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_sigfillset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-15032385537)
	return
}
func F_significant_digits(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	v7 = l0
	goto L1
L1:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v14 = v12 - int32(43)
	v21 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v14))|base.B2i32(int32(1)<<(uint(v14)%32)&int32(37) == v21) == v21 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v28 = v7
	v29 = v12
	v31 = int32(1)
	goto L9
L3:
	;
	v7 = v7 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return v74
L7:
	;
	v74 = v31
	goto L6
L8:
	;
	v44 = v28
	v45 = v29
	v46 = int32(0)
	goto L15
L9:
	;
	switch v29 - int32(46) {
	case 0:
		v38 = v31
		goto L12
	case 1:
		goto L8
	case 2:
		goto L13
	default:
		goto L11
	}
L10:
	;
	if v29 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v28 = v28 + int32(1)
	v29 = v39
	v31 = v38
	goto L9
L13:
	;
	v38 = v31 + int32(1)
	goto L12
L14:
	;
	goto L8
L15:
	;
	v50 = base.B2i32(v45 != int32(46))
	if v50&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v45-int32(48))&int32(255))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v66 != 0 {
		v74 = v66
		goto L6
	} else {
		goto L21
	}
L17:
	;
	v60 = v46 + v50
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v61 != 0 {
		v44 = v44 + int32(1)
		v45 = v61
		v46 = v60
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v66 = v46
	goto L19
L19:
	;
	goto L16
L20:
	;
	v66 = v60
	goto L19
L21:
	;
	goto L7
}
func F_sjis_to_euc_jp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v16, v17, v18, int32(35), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v18 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v324)
	return v315 - v15
L4:
	;
	v313 = v14
	v315 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v14
	v30 = v15
	v31 = v18
	goto L7
L7:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v40 = base.I32_extend8_s(v39)
	if int32(0) <= v40 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v313 = v298
	v315 = v309
	goto L3
L9:
	;
	if int32(0) < v301 {
		v28 = v298
		v30 = v309
		v31 = v301
		goto L7
	} else {
		goto L70
	}
L10:
	;
	if v40 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v56 = F_pg_encoding_verifymbchar(m, int32(35), v30, v31)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v13 != 0 {
		v313 = v28
		v315 = v30
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v40)
	v49 = int32(1)
	v298 = v28 + v49
	v301 = v31 - v49
	v309 = v30 + v49
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(35), v30, v31)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v56 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v13 != 0 {
		v313 = v28
		v315 = v30
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32((v40+int32(95))&int32(255)) <= base.Ui32(int32(62)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	F_report_invalid_encoding(m, int32(35), v30, v31)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v298 = v284
	v301 = v31 - v56
	v309 = v30 + v56
	goto L9
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v40)
	v70 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v70)
	v284 = v28 + int32(2)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v78 = v75 | v39<<(uint(int32(8))%32)
	if base.Ui32(v78-int32(_a_F_sjis_to_euc_jp_0)) <= base.Ui32(int32(767)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v83 = v78
	v85 = v75
	v88 = v39
	v89 = int32(0)
	goto L31
L29:
	;
	v125 = v78
	v127 = v75
	v130 = v39
	goto L30
L30:
	;
	if v125 <= int32(_a_F_sjis_to_euc_jp_1) {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v96 = v89 << (uint(int32(3)) % 32)
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_sjis_to_euc_jp[0]))))
	if v99 == v83 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v125 = v118
	v127 = v119
	v130 = v120
	goto L30
L33:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_sjis_to_euc_jp[1]))))
	v106 = v101
	v107 = v101 & int32(255)
	v108 = int32(base.Ui32(v101) >> (uint(int32(8)) % 32))
	goto L35
L34:
	;
	v106 = v83
	v107 = v85
	v108 = v88
	goto L35
L35:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_sjis_to_euc_jp[2]))))
	if v111 == v106 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+uint32(_c_F_sjis_to_euc_jp[3]))))
	v118 = v113
	v119 = v113 & int32(255)
	v120 = int32(base.Ui32(v113) >> (uint(int32(8)) % 32))
	goto L38
L37:
	;
	v118 = v106
	v119 = v107
	v120 = v108
	goto L38
L38:
	;
	v122 = v89 + int32(2)
	if v122 != int32(388) {
		v83 = v118
		v85 = v119
		v88 = v120
		v89 = v122
		goto L31
	} else {
		goto L39
	}
L39:
	;
	goto L32
L40:
	;
	v142 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v127))
	if base.Ui32(int32(158)) < base.Ui32(v127) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v164 = int32(0)
	if base.B2i32(base.B2i32(v125 != int32(_a_F_sjis_to_euc_jp_2))&base.B2i32(base.Ui32(v125) < base.Ui32(int32(_a_F_sjis_to_euc_jp_3))) == v164)&base.B2i32(base.Ui32(int32(176)) < base.Ui32(v125-int32(_a_F_sjis_to_euc_jp_4))) == v164 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v143 = int32(2)
	goto L45
L44:
	;
	v143 = int32(96)
	goto L45
L45:
	;
	v147 = v143 + v127 + base.B2i32(base.Ui32(v127) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v147)
	v155 = v130<<(uint(int32(1))%32)&int32(126) | v142 + int32(159)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v155)
	v284 = v28 + int32(2)
	goto L24
L46:
	;
	v173 = int32(_a_F_sjis_to_euc_jp_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v28))) = uint16(v173)
	v284 = v28 + int32(2)
	goto L24
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(v125-int32(_a_F_sjis_to_euc_jp_3)) <= base.Ui32(int32(1279)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v184 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v127))
	if base.Ui32(int32(158)) < base.Ui32(v127) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(v125-int32(_a_F_sjis_to_euc_jp_6)) <= base.Ui32(int32(1279)) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v185 = int32(2)
	goto L54
L53:
	;
	v185 = int32(96)
	goto L54
L54:
	;
	v189 = v185 + v127 + base.B2i32(base.Ui32(v127) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v189)
	v199 = (v130<<(uint(int32(1))%32)+int32(34))&int32(126) | v184 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v199)
	v284 = v28 + int32(2)
	goto L24
L55:
	;
	v207 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v207)
	v212 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v127))
	if base.Ui32(int32(158)) < base.Ui32(v127) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v125) < base.Ui32(int32(_a_F_sjis_to_euc_jp_7)) {
		v284 = v28
		goto L24
	} else {
		goto L61
	}
L58:
	;
	v213 = int32(2)
	goto L60
L59:
	;
	v213 = int32(96)
	goto L60
L60:
	;
	v217 = v213 + v127 + base.B2i32(base.Ui32(v127) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)) = uint8(v217)
	v227 = (v130<<(uint(int32(1))%32)+int32(24))&int32(126) | v212 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v227)
	v284 = v28 + int32(3)
	goto L24
L61:
	;
	v234 = v125
	v235 = v28
	v240 = int32(0)
	goto L62
L62:
	;
	v247 = v240 << (uint(int32(3)) % 32)
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_sjis_to_euc_jp[1]))))
	if v234 != v250 {
		v277 = v234
		v278 = v235
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v284 = v278
	goto L24
L64:
	;
	v280 = v240 + int32(1)
	if v280 != int32(388) {
		v234 = v277
		v235 = v278
		v240 = v280
		goto L62
	} else {
		goto L69
	}
L65:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_sjis_to_euc_jp[4])))
	if int32(_a_F_sjis_to_euc_jp_8) <= v252 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v255 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v255)
	v257 = int32(128)
	v258 = v252 | v257
	*(*uint8)(unsafe.Add(mBase, uint32(v235)+2)) = uint8(v258)
	v263 = int32(base.Ui32(v252)>>(uint(int32(8))%32)) | v257
	*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)) = uint8(v263)
	v277 = v252
	v278 = v235 + int32(3)
	goto L64
L67:
	;
	goto L68
L68:
	;
	v267 = int32(128)
	v268 = v252 | v267
	*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)) = uint8(v268)
	v273 = int32(base.Ui32(v252)>>(uint(int32(8))%32)) | v267
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v273)
	v277 = v252
	v278 = v235 + int32(2)
	goto L64
L69:
	;
	goto L63
L70:
	;
	goto L8
}
func F_slice_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v7 = int32(-1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 < int32(0) {
		v29 = v7
		return v29
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v11 < v8 {
			v29 = v7
			return v29
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v13 < v11 {
				v29 = v7
				return v29
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v15 == int32(0) {
					v29 = v7
					return v29
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(4))))
					if v20 < v13 {
						v29 = v7
						return v29
					} else {
						v22 = int32(0)
						v25 = F_replace_s(m, l0, v8, v11, v22, v22, v22)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = v25
							return v29
						}
					}
				}
			}
		}
	}
}
func F_sort_pending_writebacks_med3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v17) < base.Ui32(v18) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return l0
L2:
	;
	return l2
L3:
	;
	return v101
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(v18) < base.Ui32(v71) {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(v18) < base.Ui32(v34) {
		v101 = l1
		goto L3
	} else {
		goto L15
	}
L6:
	;
	if base.Ui32(v18) < base.Ui32(v17) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(v15) < base.Ui32(v13) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(v13) < base.Ui32(v15) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if base.Ui32(v16) < base.Ui32(v14) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if base.Ui32(v14) < base.Ui32(v16) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v25 < v26 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v26 < v25 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v30) <= base.Ui32(v29) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v34) < base.Ui32(v18) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui32(v17) < base.Ui32(v34) {
		goto L2
	} else {
		goto L25
	}
L17:
	;
	if base.Ui32(v13) < base.Ui32(v36) {
		v101 = l1
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v36) < base.Ui32(v13) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v14) < base.Ui32(v37) {
		v101 = l1
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v37) < base.Ui32(v14) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v43 < v44 {
		v101 = l1
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v44 < v43 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v47) < base.Ui32(v48) {
		v101 = l1
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	if base.Ui32(v34) < base.Ui32(v17) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v15) < base.Ui32(v36) {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v36) < base.Ui32(v15) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v16) < base.Ui32(v37) {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v37) < base.Ui32(v16) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v58 < v59 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v59 < v58 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v62) < base.Ui32(v63) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v65 = l2
	goto L35
L34:
	;
	v65 = l0
	goto L35
L35:
	;
	return v65
L36:
	;
	if base.Ui32(v17) < base.Ui32(v71) {
		goto L1
	} else {
		goto L46
	}
L37:
	;
	if base.Ui32(v71) < base.Ui32(v18) {
		v101 = l1
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v13) < base.Ui32(v69) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(v69) < base.Ui32(v13) {
		v101 = l1
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v14) < base.Ui32(v70) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32(v70) < base.Ui32(v14) {
		v101 = l1
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v78 < v79 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v79 < v78 {
		v101 = l1
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v83) < base.Ui32(v82) {
		v101 = l1
		goto L3
	} else {
		goto L45
	}
L45:
	;
	goto L36
L46:
	;
	if base.Ui32(v71) < base.Ui32(v17) {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(v15) < base.Ui32(v69) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if base.Ui32(v69) < base.Ui32(v15) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(v16) < base.Ui32(v70) {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v70) < base.Ui32(v16) {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v93 < v94 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v94 < v93 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v97) < base.Ui32(v98) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v100 = l0
	goto L56
L55:
	;
	v100 = l2
	goto L56
L56:
	;
	v101 = v100
	goto L3
}
func F_spcache_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0]))
	if v4 == int32(0) {
		v17 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[1])) = uint8(v17)
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v20
		*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[3])) = v20
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
		if v29 == v20 {
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[5]))
			v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_spcache_init_0), int32(0), int32(_a_F_spcache_init_1), int32(_a_F_spcache_init_2))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4])) = v39
				v46 = v39
				v48 = F_MemoryContextAllocZero(m, v46, int32(32))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
					v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
						*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
						v63 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
						*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
						return
					}
				}
			}
		} else {
			F_MemoryContextReset(m, v29)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
				v46 = v45
				v48 = F_MemoryContextAllocZero(m, v46, int32(32))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
					v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
						*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
						v63 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
						*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
						return
					}
				}
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])))
		if v8&int32(1) == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[1])) = uint8(v17)
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v20
			*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[3])) = v20
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
			if v29 == v20 {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[5]))
				v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_spcache_init_0), int32(0), int32(_a_F_spcache_init_1), int32(_a_F_spcache_init_2))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4])) = v39
					v46 = v39
					v48 = F_MemoryContextAllocZero(m, v46, int32(32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
						v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
							*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
							return
						}
					}
				}
			} else {
				F_MemoryContextReset(m, v29)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
					v46 = v45
					v48 = F_MemoryContextAllocZero(m, v46, int32(32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
						v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
							*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
							return
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			if base.Ui32(v13) < base.Ui32(int32(256)) {
				return
			} else {
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[1])) = uint8(v17)
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v20
				*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[3])) = v20
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
				if v29 == v20 {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[5]))
					v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_spcache_init_0), int32(0), int32(_a_F_spcache_init_1), int32(_a_F_spcache_init_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4])) = v39
						v46 = v39
						v48 = F_MemoryContextAllocZero(m, v46, int32(32))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
							v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
								*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
								*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
								v63 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
								return
							}
						}
					}
				} else {
					F_MemoryContextReset(m, v29)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
						v46 = v45
						v48 = F_MemoryContextAllocZero(m, v46, int32(32))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
							v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
								*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
								*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
								v63 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_spghandler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v59 int64
	_ = v59
	v3 = F_palloc0(m, int32(140))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(16777223)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1970324836975030)
		v11 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+16)) = uint16(v11)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(261)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(262)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(263)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(264)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(265)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(266)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(267)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = int32(268)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(269)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(270)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = int32(271)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(272)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(273)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(274)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(275)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(276)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v3)+26)) = int32(50331649)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+18)) = int64(16842753)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = v11
		v59 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+128)) = v59
		*(*int64)(unsafe.Add(mBase, uint32(v3)+120)) = v59
		*(*int64)(unsafe.Add(mBase, uint32(v3)+112)) = v59
		return v3
	}
}
func F_spgrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int64
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int64
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v404 int32
	_ = v404
	var v405 int64
	_ = v405
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l1 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v14 * int32(48)
	if v18 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v21, l1, v18)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+112)) = v140
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v150 {
		goto L32
	} else {
		goto L33
	}
L6:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+108)) = v133
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v134)+116)) = v136
	v140 = int32(0)
	v144 = v134
	goto L5
L7:
	;
	if v24 <= int32(0) {
		v133 = v24
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v63 = v24
	goto L9
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+108)) = v63
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+116)) = v73
	v75 = int32(0)
	if v63 <= v75 {
		v140 = v75
		v144 = v71
		goto L5
	} else {
		goto L20
	}
L10:
	;
	v28 = v24 * int32(48)
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v29, l3, v28)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v31 <= int32(0) {
		v133 = v31
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(0)
	goto L15
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v38*int32(48))+20))
	v50 = F_get_func_rettype(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v63 = v59
	goto L9
L17:
	;
	return
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v38<<(uint(int32(2))%32)))) = v50
	v58 = v38 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v58 < v59 {
		v38 = v58
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v78 <= int32(0) {
		v140 = v75
		v144 = v71
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v83 = v75
	v85 = int32(0)
	goto L22
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v71)+116))
	v96 = v93 + v85*int32(48)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v140 = v121
	v144 = v71
	goto L5
L24:
	;
	if v83 != v85 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v121 = v83
	v122 = int32(-1)
	goto L26
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v71)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v85<<(uint(int32(2))%32)))) = v122
	v129 = v85 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v129 < v130 {
		v83 = v121
		v85 = v129
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v105 = v93 + v83*int32(48)
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v96)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+40)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v96)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+32)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v116
	goto L29
L28:
	;
	goto L29
L29:
	;
	v121 = v83 + int32(1)
	v122 = v83
	goto L26
L30:
	;
	goto L23
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+100)) = v220
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_MemoryContextReset(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L17
	} else {
		goto L48
	}
L32:
	;
	v153 = int32(0)
	v156 = v153
	v158 = v153
	v161 = v150
	v163 = v6
	v164 = v6
	goto L36
L33:
	;
	goto L34
L34:
	;
	v216 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+96)) = uint16(v216)
	v220 = int32(0)
	goto L31
L35:
	;
	v213 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+96)) = uint16(v213)
	v220 = v213
	goto L31
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v168 = v165 + v158*int32(48)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v169&int32(64) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v200&v201 != 0 {
		goto L35
	} else {
		goto L47
	}
L38:
	;
	v203 = v158 + int32(1)
	if v203 < v199 {
		v156 = v198
		v158 = v203
		v161 = v199
		v163 = v200
		v164 = v201
		goto L36
	} else {
		goto L46
	}
L39:
	;
	v198 = v156
	v199 = v161
	v200 = v163
	v201 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v169&int32(128) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v198 = v156
	v199 = v161
	v200 = int32(1)
	v201 = v164
	goto L38
L43:
	;
	goto L44
L44:
	;
	if v169&int32(1) != 0 {
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v144)+104))
	v181 = v178 + v156*int32(48)
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v168)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+40)) = v182
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v168)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+32)) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v168)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+24)) = v186
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v168)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+16)) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = v190
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	*(*int64)(unsafe.Add(mBase, uint32(v181))) = v192
	v194 = int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v198 = v156 + v194
	v199 = v197
	v200 = v194
	v201 = v164
	goto L38
L46:
	;
	goto L37
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+97)) = uint8(v200)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+96)) = uint8(v201)
	v220 = v198
	goto L31
L48:
	;
	v233 = int32(_a_F_spgrescan_0)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_spgrescan[0]))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_spgrescan[0])) = v236
	v239 = F_pairingheap_allocate(m, int32(256), v11)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v239
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+96)))
	if v242 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v246 = F_palloc(m, int32(40))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L17
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+97)))
	if v262 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v248 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v246)+32)) = uint16(v248)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = int32(_a_F_spgrescan_1)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+34)) = v248
	v254 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v246)+12)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v246)+20)) = v254
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	F_pairingheap_add(m, v258, v246)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	v271 = F_palloc(m, v266<<(uint(int32(3))%32)+int32(40))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L17
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spgrescan[0])) = v234
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	if v306 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v271)+34)) = uint8(v273)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	if v275 <= v273 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v271)+37)) = uint8(v286)
	*(*uint16)(unsafe.Add(mBase, uint32(v271)+35)) = uint16(v286)
	v290 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v271)+32)) = uint16(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+28)) = int32(_a_F_spgrescan_2)
	v294 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v271)+12)) = v294
	*(*int64)(unsafe.Add(mBase, uint32(v271)+20)) = v294
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	F_pairingheap_add(m, v298, v271)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L17
	} else {
		goto L62
	}
L60:
	;
	v279 = v275 << (uint(int32(3)) % 32)
	if v279 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	base.MemoryCopy(m, v271+int32(40), v265, v279)
	goto L59
L62:
	;
	goto L57
L63:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+208)))
	if v346 != int32(1) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v309 <= int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v316 = v309
	v318 = int32(0)
	goto L66
L66:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(_a_F_spgrescan_3)+v318<<(uint(int32(2))%32))))
	if v328 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L63
L68:
	;
	F_pfree(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L17
	} else {
		goto L71
	}
L69:
	;
	v332 = v316
	goto L70
L70:
	;
	v334 = v318 + int32(1)
	if v334 < v332 {
		v316 = v332
		v318 = v334
		goto L66
	} else {
		goto L72
	}
L71:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	v332 = v331
	goto L70
L72:
	;
	goto L67
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+216)) = int64(0)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+272))
	if v388 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v349 <= int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v358 = int32(0)
	goto L76
L76:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(3488)+v358<<(uint(int32(2))%32))))
	F_pfree(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L17
	} else {
		goto L78
	}
L77:
	;
	goto L73
L78:
	;
	v372 = v358 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v372 < v373 {
		v358 = v372
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v404 != 0 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+268)))
	if v391 != int32(1) {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	v398 = v388
	goto L83
L83:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v398)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v398)+16)) = v399 + int64(1)
	goto L80
L84:
	;
	F_pgstat_assoc_relation(m, v387)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+272))
	v398 = v397
	goto L83
L86:
	;
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v404)))
	*(*int64)(unsafe.Add(mBase, uint32(v404))) = v405 + int64(1)
	goto L88
L87:
	;
	goto L88
L88:
	;
	return
}
func F_split_pathtarget_at_srfs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v8 int32
	_ = v8
	F_split_pathtarget_at_srfs_extended(m, l0, l1, l2, l3, l4, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_split_pathtarget_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v128 = F_palloc(m, int32(8))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L43
	}
L2:
	;
	v54 = F_palloc(m, int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L25
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v9 != int32(1) {
		v30 = l0
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v51 = int32(0)
	goto L5
L5:
	;
	return v51
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v32 = F_list_member(m, v31, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L13
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+45)))
	if v14 != int32(1) {
		v30 = l0
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if v17 == int32(0) {
		v30 = l0
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+324))
	v21 = F_bms_make_singleton(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v26 = F_remove_nulling_relids(m, l0, v21, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = v26
	goto L6
L13:
	;
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	goto L1
L15:
	;
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v34 - int32(6) {
	case 0, 3, 4, 5:
		goto L1
	case 1, 2, 6, 7, 8, 10:
		goto L17
	case 9:
		goto L18
	case 11:
		goto L19
	default:
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	v46 = F_expression_tree_walker_impl(m, l0, int32(895), l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L24
	}
L18:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v42 != 0 {
		goto L2
	} else {
		goto L23
	}
L19:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v39 != int32(1) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	if v34 != int32(319) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	goto L2
L23:
	;
	goto L17
L24:
	;
	v51 = v46
	goto L5
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = l0
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v62
	v67 = F_expression_tree_walker_impl(m, l0, int32(895), l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v72 = v70
	goto L29
L28:
	;
	v72 = int32(0)
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v75 = v73 + int32(1)
	if v72 <= v75 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v78 = F_lappend(m, v69, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	v92 = v69
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v95 = v75 << (uint(int32(2)) % 32)
	v96 = v93 + v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = F_lappend(m, v97, v54)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L36
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v83 = F_lappend(m, v81, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v88 = F_lappend(m, v86, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v92 = v91
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v98
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v103 = v102 + v95
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v106 = F_list_concat(m, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v111 = v110 + v95
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v114 = F_list_concat(m, v112, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v114
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v58
	v118 = F_lappend(m, v57, v54)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	if v75 < v56 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v121 = v56
	goto L42
L41:
	;
	v121 = v75
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118
	return int32(0)
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = l0
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v134 = F_lappend(m, v133, v128)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v134
	return int32(0)
}
func F_sqrt_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v242 int64
	_ = v242
	var v252 int64
	_ = v252
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v298 int64
	_ = v298
	var v308 int64
	_ = v308
	var v313 int64
	_ = v313
	var v321 int32
	_ = v321
	var v335 int64
	_ = v335
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v373 int64
	_ = v373
	var v384 int64
	_ = v384
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v430 int64
	_ = v430
	var v434 int64
	_ = v434
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v461 int64
	_ = v461
	var v465 int64
	_ = v465
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int64
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int64
	_ = v556
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v578 int64
	_ = v578
	var v579 int64
	_ = v579
	var v582 int64
	_ = v582
	var v583 int64
	_ = v583
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int64
	_ = v603
	var v605 int64
	_ = v605
	var v607 int32
	_ = v607
	var v613 int64
	_ = v613
	var v619 int32
	_ = v619
	var v636 int64
	_ = v636
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v645 int64
	_ = v645
	var v650 int64
	_ = v650
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int32
	_ = v656
	var v661 int64
	_ = v661
	var v663 int64
	_ = v663
	var v665 int64
	_ = v665
	var v667 int32
	_ = v667
	var v669 int64
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v685 int64
	_ = v685
	var v693 int64
	_ = v693
	var v700 int64
	_ = v700
	var v701 int64
	_ = v701
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v722 int64
	_ = v722
	var v723 int64
	_ = v723
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int64
	_ = v743
	var v744 int64
	_ = v744
	var v751 int32
	_ = v751
	var v752 int64
	_ = v752
	var v753 int64
	_ = v753
	var v754 int64
	_ = v754
	var v755 int64
	_ = v755
	var v760 int64
	_ = v760
	var v763 int64
	_ = v763
	var v766 int64
	_ = v766
	var v769 int64
	_ = v769
	var v770 int64
	_ = v770
	var v774 int64
	_ = v774
	var v781 int64
	_ = v781
	var v793 int32
	_ = v793
	var v794 int64
	_ = v794
	var v795 int64
	_ = v795
	var v799 int64
	_ = v799
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v852 int64
	_ = v852
	var v860 int64
	_ = v860
	var v867 int64
	_ = v867
	var v868 int64
	_ = v868
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v893 int64
	_ = v893
	var v894 int64
	_ = v894
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int64
	_ = v910
	var v911 int64
	_ = v911
	var v918 int32
	_ = v918
	var v919 int64
	_ = v919
	var v920 int64
	_ = v920
	var v921 int64
	_ = v921
	var v922 int64
	_ = v922
	var v927 int64
	_ = v927
	var v930 int64
	_ = v930
	var v933 int64
	_ = v933
	var v936 int64
	_ = v936
	var v937 int64
	_ = v937
	var v941 int64
	_ = v941
	var v948 int64
	_ = v948
	var v960 int32
	_ = v960
	var v961 int64
	_ = v961
	var v962 int64
	_ = v962
	var v966 int64
	_ = v966
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v1005 int32
	_ = v1005
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1071 int32
	_ = v1071
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1158 int32
	_ = v1158
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1218 int32
	_ = v1218
	var v1219 int64
	_ = v1219
	var v1238 int32
	_ = v1238
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1282 int32
	_ = v1282
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1316 int32
	_ = v1316
	var v1323 int32
	_ = v1323
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1369 int32
	_ = v1369
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1429 int32
	_ = v1429
	var v1430 int64
	_ = v1430
	var v1448 int32
	_ = v1448
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int64
	_ = v1489
	var v1491 int64
	_ = v1491
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1512 int64
	_ = v1512
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1640 int32
	_ = v1640
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1753 int32
	_ = v1753
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1771 int32
	_ = v1771
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1789 int32
	_ = v1789
	var v1794 int32
	_ = v1794
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1832 int32
	_ = v1832
	var v1842 int32
	_ = v1842
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1989 int32
	_ = v1989
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2007 int32
	_ = v2007
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2057 int32
	_ = v2057
	var v2068 int32
	_ = v2068
	var v2078 int32
	_ = v2078
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2128 int32
	_ = v2128
	var v2130 int32
	_ = v2130
	var v2132 int32
	_ = v2132
	var v2133 int64
	_ = v2133
	var v2135 int64
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2159 int32
	_ = v2159
	var v2161 int32
	_ = v2161
	var v2162 int64
	_ = v2162
	var v2164 int64
	_ = v2164
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2225 int32
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2271 int32
	_ = v2271
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2320 int32
	_ = v2320
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2338 int32
	_ = v2338
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2356 int32
	_ = v2356
	var v2361 int32
	_ = v2361
	var v2369 int32
	_ = v2369
	var v2373 int32
	_ = v2373
	var v2378 int32
	_ = v2378
	var v2382 int32
	_ = v2382
	var v2388 int32
	_ = v2388
	var v2399 int32
	_ = v2399
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2447 int32
	_ = v2447
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2496 int32
	_ = v2496
	var v2505 int32
	_ = v2505
	var v2507 int32
	_ = v2507
	var v2514 int32
	_ = v2514
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2532 int32
	_ = v2532
	var v2537 int32
	_ = v2537
	var v2545 int32
	_ = v2545
	var v2549 int32
	_ = v2549
	var v2554 int32
	_ = v2554
	var v2558 int32
	_ = v2558
	var v2564 int32
	_ = v2564
	var v2575 int32
	_ = v2575
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2595 int32
	_ = v2595
	var v2598 int32
	_ = v2598
	var v2607 int32
	_ = v2607
	var v2616 int32
	_ = v2616
	var v2633 int64
	_ = v2633
	var v2638 int64
	_ = v2638
	var v2641 int64
	_ = v2641
	var v2642 int64
	_ = v2642
	var v2644 int64
	_ = v2644
	var v2649 int64
	_ = v2649
	var v2672 int64
	_ = v2672
	var v2677 int64
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2684 int64
	_ = v2684
	var v2690 int32
	_ = v2690
	var v2706 int64
	_ = v2706
	var v2713 int64
	_ = v2713
	var v2716 int64
	_ = v2716
	var v2721 int64
	_ = v2721
	var v2723 int64
	_ = v2723
	var v2725 int64
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2732 int64
	_ = v2732
	var v2734 int64
	_ = v2734
	var v2736 int32
	_ = v2736
	var v2738 int64
	_ = v2738
	var v2748 int32
	_ = v2748
	var v2764 int64
	_ = v2764
	var v2774 int64
	_ = v2774
	var v2779 int64
	_ = v2779
	var v2801 int64
	_ = v2801
	var v2817 int32
	_ = v2817
	var v2831 int64
	_ = v2831
	var v2832 int64
	_ = v2832
	var v2837 int64
	_ = v2837
	var v2841 int32
	_ = v2841
	var v2843 int64
	_ = v2843
	var v2848 int64
	_ = v2848
	var v2849 int64
	_ = v2849
	var v2851 int64
	_ = v2851
	var v2854 int64
	_ = v2854
	var v2855 int64
	_ = v2855
	var v2857 int64
	_ = v2857
	var v2858 int64
	_ = v2858
	var v2862 int64
	_ = v2862
	var v2869 int64
	_ = v2869
	var v2881 int32
	_ = v2881
	var v2882 int64
	_ = v2882
	var v2883 int64
	_ = v2883
	var v2886 int64
	_ = v2886
	var v2887 int64
	_ = v2887
	var v2890 int64
	_ = v2890
	var v2891 int64
	_ = v2891
	var v2892 int64
	_ = v2892
	var v2897 int64
	_ = v2897
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2906 int64
	_ = v2906
	var v2907 int64
	_ = v2907
	var v2910 int64
	_ = v2910
	var v2916 int64
	_ = v2916
	var v2917 int64
	_ = v2917
	var v2920 int64
	_ = v2920
	var v2926 int64
	_ = v2926
	var v2927 int64
	_ = v2927
	var v2928 int64
	_ = v2928
	var v2929 int64
	_ = v2929
	var v2942 int32
	_ = v2942
	var v2943 int64
	_ = v2943
	var v2944 int64
	_ = v2944
	var v2949 int64
	_ = v2949
	var v2950 int64
	_ = v2950
	var v2952 int64
	_ = v2952
	var v2955 int64
	_ = v2955
	var v2956 int64
	_ = v2956
	var v2958 int64
	_ = v2958
	var v2959 int64
	_ = v2959
	var v2963 int64
	_ = v2963
	var v2970 int64
	_ = v2970
	var v2982 int32
	_ = v2982
	var v2987 int64
	_ = v2987
	var v2988 int64
	_ = v2988
	var v2990 int64
	_ = v2990
	var v2993 int64
	_ = v2993
	var v2994 int64
	_ = v2994
	var v2996 int64
	_ = v2996
	var v2997 int64
	_ = v2997
	var v3001 int64
	_ = v3001
	var v3008 int64
	_ = v3008
	var v3020 int32
	_ = v3020
	var v3021 int64
	_ = v3021
	var v3022 int64
	_ = v3022
	var v3023 int64
	_ = v3023
	var v3032 int64
	_ = v3032
	var v3033 int64
	_ = v3033
	var v3035 int64
	_ = v3035
	var v3038 int64
	_ = v3038
	var v3039 int64
	_ = v3039
	var v3041 int64
	_ = v3041
	var v3042 int64
	_ = v3042
	var v3046 int64
	_ = v3046
	var v3053 int64
	_ = v3053
	var v3065 int32
	_ = v3065
	var v3067 int64
	_ = v3067
	var v3070 int64
	_ = v3070
	var v3071 int64
	_ = v3071
	var v3076 int64
	_ = v3076
	var v3077 int64
	_ = v3077
	var v3080 int64
	_ = v3080
	var v3083 int64
	_ = v3083
	var v3084 int64
	_ = v3084
	var v3091 int64
	_ = v3091
	var v3102 int64
	_ = v3102
	var v3103 int64
	_ = v3103
	var v3105 int64
	_ = v3105
	var v3107 int64
	_ = v3107
	var v3112 int64
	_ = v3112
	var v3113 int64
	_ = v3113
	var v3114 int64
	_ = v3114
	var v3117 int64
	_ = v3117
	var v3119 int64
	_ = v3119
	var v3120 int64
	_ = v3120
	var v3121 int64
	_ = v3121
	var v3124 int64
	_ = v3124
	var v3126 int64
	_ = v3126
	var v3127 int64
	_ = v3127
	var v3130 int64
	_ = v3130
	var v3133 int64
	_ = v3133
	var v3134 int64
	_ = v3134
	var v3141 int64
	_ = v3141
	var v3146 int32
	_ = v3146
	var v3147 int64
	_ = v3147
	var v3148 int64
	_ = v3148
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3157 int32
	_ = v3157
	var v3167 int64
	_ = v3167
	var v3175 int64
	_ = v3175
	var v3182 int64
	_ = v3182
	var v3183 int64
	_ = v3183
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3204 int64
	_ = v3204
	var v3205 int64
	_ = v3205
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3219 int32
	_ = v3219
	var v3221 int32
	_ = v3221
	var v3225 int64
	_ = v3225
	var v3226 int64
	_ = v3226
	var v3232 int64
	_ = v3232
	var v3233 int64
	_ = v3233
	var v3234 int64
	_ = v3234
	var v3235 int64
	_ = v3235
	var v3240 int64
	_ = v3240
	var v3243 int64
	_ = v3243
	var v3246 int64
	_ = v3246
	var v3249 int64
	_ = v3249
	var v3250 int64
	_ = v3250
	var v3254 int64
	_ = v3254
	var v3261 int64
	_ = v3261
	var v3273 int32
	_ = v3273
	var v3274 int64
	_ = v3274
	var v3275 int64
	_ = v3275
	var v3279 int64
	_ = v3279
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3292 int32
	_ = v3292
	var v3294 int32
	_ = v3294
	var v3323 int32
	_ = v3323
	var v3326 int32
	_ = v3326
	var v3328 int32
	_ = v3328
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3351 int32
	_ = v3351
	var v3356 int32
	_ = v3356
	var v3365 int32
	_ = v3365
	var v3367 int32
	_ = v3367
	var v3387 int32
	_ = v3387
	var v3395 int32
	_ = v3395
	var v3401 int32
	_ = v3401
	var v3405 int32
	_ = v3405
	var v3409 int32
	_ = v3409
	var v3427 int32
	_ = v3427
	var v3432 int32
	_ = v3432
	var v3434 int32
	_ = v3434
	var v3436 int32
	_ = v3436
	var v3438 int32
	_ = v3438
	var v3443 int32
	_ = v3443
	var v3445 int32
	_ = v3445
	var v3446 int32
	_ = v3446
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3483 int32
	_ = v3483
	var v3488 int32
	_ = v3488
	var v3498 int32
	_ = v3498
	var v3529 int32
	_ = v3529
	var v3544 int64
	_ = v3544
	var v3549 int64
	_ = v3549
	var v3552 int64
	_ = v3552
	var v3554 int64
	_ = v3554
	var v3555 int64
	_ = v3555
	var v3556 int64
	_ = v3556
	var v3557 int64
	_ = v3557
	var v3563 int64
	_ = v3563
	var v3565 int64
	_ = v3565
	var v3567 int64
	_ = v3567
	var v3575 int64
	_ = v3575
	var v3598 int64
	_ = v3598
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3613 int32
	_ = v3613
	var v3625 int64
	_ = v3625
	var v3631 int64
	_ = v3631
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3652 int64
	_ = v3652
	var v3664 int32
	_ = v3664
	var v3666 int64
	_ = v3666
	var v3669 int64
	_ = v3669
	var v3672 int32
	_ = v3672
	var v3681 int32
	_ = v3681
	var v3684 int32
	_ = v3684
	var v3711 int32
	_ = v3711
	var v3736 int32
	_ = v3736
	var v3743 int32
	_ = v3743
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3768 int32
	_ = v3768
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3773 int32
	_ = v3773
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3788 int32
	_ = v3788
	var v3789 int64
	_ = v3789
	var v3791 int64
	_ = v3791
	var v3794 int32
	_ = v3794
	var v3797 int32
	_ = v3797
	var v3809 int32
	_ = v3809
	var v3818 int32
	_ = v3818
	var v3820 int32
	_ = v3820
	var v3824 int32
	_ = v3824
	var v3825 int32
	_ = v3825
	var v3836 int32
	_ = v3836
	var v3839 int32
	_ = v3839
	var v3840 int32
	_ = v3840
	var v3843 int32
	_ = v3843
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3847 int32
	_ = v3847
	var v3848 int32
	_ = v3848
	var v3849 int32
	_ = v3849
	var v3852 int32
	_ = v3852
	var v3855 int32
	_ = v3855
	var v3860 int32
	_ = v3860
	var v3864 int32
	_ = v3864
	var v3870 int32
	_ = v3870
	var v3876 int32
	_ = v3876
	var v3877 int32
	_ = v3877
	var v3880 int32
	_ = v3880
	var v3883 int32
	_ = v3883
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3890 int32
	_ = v3890
	var v3898 int32
	_ = v3898
	var v3902 int32
	_ = v3902
	var v3903 int32
	_ = v3903
	var v3906 int32
	_ = v3906
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3935 int32
	_ = v3935
	var v3936 int32
	_ = v3936
	var v3960 int32
	_ = v3960
	var v3966 int32
	_ = v3966
	var v3995 int32
	_ = v3995
	var v3996 int32
	_ = v3996
	var v4000 int32
	_ = v4000
	var v4001 int32
	_ = v4001
	var v4014 int32
	_ = v4014
	var v4046 int32
	_ = v4046
	var v4047 int32
	_ = v4047
	var v4073 int32
	_ = v4073
	var v4075 int32
	_ = v4075
	var v4076 int32
	_ = v4076
	var v4078 int32
	_ = v4078
	var v4080 int32
	_ = v4080
	var v4082 int32
	_ = v4082
	var v4083 int32
	_ = v4083
	var v4085 int32
	_ = v4085
	var v4086 int32
	_ = v4086
	var v4090 int32
	_ = v4090
	v4 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(512)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v32 + int32(512)
	return
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v47 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_pfree(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v43 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v43
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v43
	goto L1
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	if int32(0) <= v187 {
		goto L60
	} else {
		goto L61
	}
L11:
	;
	v50 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+328)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+312)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+288)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+296)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+304)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+264)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+272)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+280)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+240)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+248)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+256)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+208)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v50
	v86 = int32(-1)
	v87 = int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = v88 >> (uint(v87) % 32)
	v94 = int32(4)
	v97 = base.I32_div_s(l2+v94, v94)
	v101 = base.I32_div_s(l2^v86, int32(-4))
	if int32(0) <= l2+v87 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L8
	} else {
		goto L54
	}
L14:
	;
	v106 = v97
	goto L16
L15:
	;
	v106 = v101
	goto L16
L16:
	;
	v108 = int32(1)
	v109 = v106 + v90 + v108
	if v109 <= v108 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v112 = v87
	goto L19
L18:
	;
	v112 = v109
	goto L19
L19:
	;
	v114 = int32(1)
	v118 = (v90^v86+v112)<<(uint(v114)%32) + v88 + v114
	if v118 <= v114 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v121 = v87
	goto L22
L21:
	;
	v121 = v118
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+336)) = v121
	if int32(5) <= v118 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v131 = v121
	v134 = v4
	goto L26
L24:
	;
	v183 = v121
	v187 = v86
	goto L25
L25:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v207 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206))))
	if v183 < int32(2) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v154 = int32(2)
	v155 = int32(base.Ui32(v131) >> (uint(v154) % 32))
	v159 = v134 + int32(1)
	if v131&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v183 = v173
	v187 = v134
	goto L25
L28:
	;
	v170 = v155
	goto L30
L29:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165))))
	v170 = v155 - base.B2i32(v166 < int32(2500))
	goto L30
L30:
	;
	v173 = v131 - v170<<(uint(int32(1))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v159<<(uint(v154)%32)))) = v173
	if int32(4) < v173 {
		v131 = v173
		v134 = v159
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v346 = base.I64_trunc_sat_f64_s(base.F64_sqrt(base.F64_convert_i64_s(v335)))
	v348 = v335 - v346*v346
	if base.B2i32(int64(0) <= v348)&base.B2i32(v348 <= v346<<(uint(int64(1))%64)) != 0 {
		v430 = v346
		v434 = v348
		goto L10
	} else {
		goto L50
	}
L33:
	;
	v321 = int32(1)
	v335 = v207
	goto L32
L34:
	;
	goto L35
L35:
	;
	if v183 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v214 = int32(1)
	v215 = v183 - v214
	v226 = v214
	v231 = int32(0)
	v242 = v207
	goto L39
L37:
	;
	v282 = int32(1)
	v298 = v207
	goto L38
L38:
	;
	v308 = v298 * int64(10000)
	if v34 <= v282 {
		v321 = v183
		v335 = v308
		goto L32
	} else {
		goto L49
	}
L39:
	;
	v252 = v242 * int64(10000)
	if v226 < v34 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v215&v214 == int32(0) {
		v321 = v183
		v335 = v270
		goto L32
	} else {
		goto L48
	}
L41:
	;
	v257 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v226<<(uint(int32(1))%32)))))
	v259 = v252 + v257
	goto L43
L42:
	;
	v259 = v252
	goto L43
L43:
	;
	v261 = v259 * int64(10000)
	v263 = v226 + int32(1)
	if v263 < v34 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v268 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v263<<(uint(int32(1))%32)))))
	v270 = v261 + v268
	goto L46
L45:
	;
	v270 = v261
	goto L46
L46:
	;
	v271 = int32(2)
	v272 = v226 + v271
	v274 = v231 + v271
	if v274 != v215&int32(-2) {
		v226 = v272
		v231 = v274
		v242 = v270
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	v282 = v272
	v298 = v270
	goto L38
L49:
	;
	v313 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v282<<(uint(int32(1))%32)))))
	v321 = v183
	v335 = v308 + v313
	goto L32
L50:
	;
	v373 = v346
	goto L51
L51:
	;
	v384 = base.I64_div_s(v335, v373)
	v387 = base.I64_div_s(v384+v373, int64(2))
	v391 = v335 - v387*v387
	if base.B2i32(v391 < int64(0))|base.B2i32(v387<<(uint(int64(1))%64) < v391) != 0 {
		v373 = v387
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v430 = v387
	v434 = v391
	goto L10
L53:
	;
	goto L52
L54:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_sqrt_var_0), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_sqrt_var_1), int32(_a_F_sqrt_var_2), int32(_a_F_sqrt_var_3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L8
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
	v3768 = v3743 << (uint(int32(1)) % 32)
	v3771 = F_palloc(m, v3768+int32(2))
	mBase = m.M
	v3772 = m.ExcPending
	if v3772 != 0 {
		goto L8
	} else {
		goto L563
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+312)) = v3711
	v3736 = int32(0)
	v3743 = v3711
	v3749 = v3736
	v3750 = v3736
	goto L58
L60:
	;
	v449 = v321
	v453 = v187
	v461 = v430
	v465 = v434
	goto L63
L61:
	;
	v3598 = v430
	goto L62
L62:
	;
	v3610 = F_palloc(m, int32(12))
	mBase = m.M
	v3611 = m.ExcPending
	if v3611 != 0 {
		goto L8
	} else {
		goto L553
	}
L63:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v453<<(uint(int32(2))%32))))
	if v477 <= int32(8) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v3598 = v3575
	goto L62
L65:
	;
	v3554 = v465*v3552 + v3544
	v3555 = int64(1)
	v3556 = v461 << (uint(v3555) % 64)
	v3557 = base.I64_div_s(v3554, v3556)
	v3563 = v3549 - v3557*v3557 + (v3554-v3556*v3557)*v3552
	v3565 = v3557 + v461*v3552
	v3567 = v3565 - v3555
	if v3563 < int64(0) {
		goto L549
	} else {
		goto L550
	}
L66:
	;
	v3387 = v449 + v482
	if v490 == int32(0) {
		goto L535
	} else {
		goto L536
	}
L67:
	;
	v3348 = int32(_a_F_sqrt_var_4)
	v3349 = v3328 * v3348
	v3351 = v3326 * v3348
	if v34 <= v3323 {
		v3365 = v3351
		v3367 = v3349
		goto L66
	} else {
		goto L532
	}
L68:
	;
	if v482&int32(1) == int32(0) {
		v3365 = v550
		v3367 = v548
		goto L66
	} else {
		goto L531
	}
L69:
	;
	v480 = v477 - v449
	v481 = int32(2)
	v482 = base.I32_div_s(v480, v481)
	if v480 < v481 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v556 = int64(63)
	v566 = v449
	v570 = v453
	v578 = v461
	v579 = v461 >> (uint(v556) % 64)
	v582 = v465
	v583 = v465 >> (uint(v556) % 64)
	goto L87
L72:
	;
	v485 = int64(0)
	v3529 = v449
	v3544 = v485
	v3549 = v485
	v3552 = int64(1)
	goto L65
L73:
	;
	goto L74
L74:
	;
	v488 = int32(1)
	v490 = v482 - v488
	if v490 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v3323 = v449
	v3326 = v488
	v3328 = int32(0)
	goto L67
L76:
	;
	goto L77
L77:
	;
	v498 = int32(0)
	v504 = v449
	v507 = v488
	v508 = v498
	v509 = v498
	goto L78
L78:
	;
	v530 = v509 * int32(_a_F_sqrt_var_4)
	if v504 < v34 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L68
L80:
	;
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v504<<(uint(int32(1))%32)))))
	v537 = v530 + v535
	goto L82
L81:
	;
	v537 = v530
	goto L82
L82:
	;
	v539 = v537 * int32(_a_F_sqrt_var_4)
	v541 = v504 + int32(1)
	if v541 < v34 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v546 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v541<<(uint(int32(1))%32)))))
	v548 = v539 + v546
	goto L85
L84:
	;
	v548 = v539
	goto L85
L85:
	;
	v550 = v507 * int32(100000000)
	v551 = int32(2)
	v552 = v504 + v551
	v554 = v508 + v551
	if v482&int32(1073741822) != v554 {
		v504 = v552
		v507 = v550
		v508 = v554
		v509 = v548
		goto L78
	} else {
		goto L86
	}
L86:
	;
	goto L79
L87:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v570<<(uint(int32(2))%32))))
	if v594 <= int32(16) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v3154 = F_palloc(m, int32(22))
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L8
	} else {
		goto L516
	}
L89:
	;
	v2841 = v32 + int32(112)
	v2843 = v2837 >> (uint(int64(63)) % 64)
	v2848 = int64(32)
	v2849 = int64(base.Ui64(v2837) >> (uint(v2848) % 64))
	v2851 = int64(base.Ui64(v582) >> (uint(v2848) % 64))
	v2854 = int64(4294967295)
	v2855 = v2837 & v2854
	v2857 = v582 & v2854
	v2858 = v2855 * v2857
	v2862 = int64(base.Ui64(v2858)>>(uint(v2848)%64)) + v2855*v2851
	v2869 = v2857*v2849 + v2862&v2854
	*(*int64)(unsafe.Add(mBase, uint32(v2841)+8)) = v582*v2843 + v583*v2837 + v2849*v2851 + int64(base.Ui64(v2862)>>(uint(v2848)%64)) + int64(base.Ui64(v2869)>>(uint(v2848)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2841))) = v2858&v2854 | v2869<<(uint(v2848)%64)
	goto L503
L90:
	;
	v2680 = v566 + v599
	if v605 == int64(1) {
		goto L489
	} else {
		goto L490
	}
L91:
	;
	v2641 = int64(10000)
	v2642 = v2633 * v2641
	v2644 = v2638 * v2641
	if v34 <= v2616 {
		v2672 = v2642
		v2677 = v2644
		goto L90
	} else {
		goto L486
	}
L92:
	;
	if v599&int32(1) == int32(0) {
		v2672 = v663
		v2677 = v665
		goto L90
	} else {
		goto L485
	}
L93:
	;
	v597 = v594 - v566
	v598 = int32(2)
	v599 = base.I32_div_s(v597, v598)
	if v597 < v598 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v672 = F_palloc(m, int32(22))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L8
	} else {
		goto L111
	}
L96:
	;
	v603 = int64(0)
	v2817 = v566
	v2831 = v603
	v2832 = v603
	v2837 = int64(1)
	goto L89
L97:
	;
	goto L98
L98:
	;
	v605 = base.I64_extend_i32_s(v599)
	v607 = base.B2i32(v605 == int64(1))
	if v605 == int64(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v2616 = v566
	v2633 = int64(0)
	v2638 = int64(1)
	goto L91
L100:
	;
	goto L101
L101:
	;
	v613 = int64(0)
	v619 = v566
	v636 = v613
	v641 = int64(1)
	v642 = v613
	goto L102
L102:
	;
	v645 = v636 * int64(10000)
	if v619 < v34 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L92
L104:
	;
	v650 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v619<<(uint(int32(1))%32)))))
	v652 = v645 + v650
	goto L106
L105:
	;
	v652 = v645
	goto L106
L106:
	;
	v654 = v652 * int64(10000)
	v656 = v619 + int32(1)
	if v656 < v34 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v661 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v656<<(uint(int32(1))%32)))))
	v663 = v654 + v661
	goto L109
L108:
	;
	v663 = v654
	goto L109
L109:
	;
	v665 = v641 * int64(100000000)
	v667 = v619 + int32(2)
	v669 = v642 + int64(2)
	if v605&int64(1073741822) != v669 {
		v619 = v667
		v636 = v663
		v641 = v665
		v642 = v669
		goto L102
	} else {
		goto L110
	}
L110:
	;
	goto L103
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+328)) = v672
	v675 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v672))) = uint16(v675)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v672 + int32(2)
	if v579 < int64(0) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v32)+312)) = v812
	v839 = F_palloc(m, int32(22))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L8
	} else {
		goto L126
	}
L113:
	;
	v708 = v672 + int32(22)
	v709 = v675
	v722 = v700
	v723 = v701
	goto L118
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = int64(16384)
	v685 = int64(0)
	v700 = v685 - v578
	v701 = v685 - (v579 + base.I64_extend_i32_u(base.B2i32(v578 != v685)))
	goto L113
L115:
	;
	goto L116
L116:
	;
	v693 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = v693
	if v578|v579 == v693 {
		v812 = v675
		v814 = int32(0)
		goto L112
	} else {
		goto L117
	}
L117:
	;
	v700 = v578
	v701 = v579
	goto L113
L118:
	;
	v734 = v32 + int32(176)
	v737 = m.G0
	v738 = int32(16)
	v739 = v737 - v738
	m.G0 = v739
	F___udivmodti4(m, v739, v722, v723, int64(10000), int64(0))
	mBase = m.M
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v739)+8))
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v739)))
	*(*int64)(unsafe.Add(mBase, uint32(v734))) = v744
	*(*int64)(unsafe.Add(mBase, uint32(v734)+8)) = v743
	m.G0 = v739 + v738
	goto L120
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v793
	v812 = v804
	v814 = v709
	goto L112
L120:
	;
	v751 = v32 + int32(160)
	v752 = *(*int64)(unsafe.Add(mBase, uint32(v32)+176))
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v32)+184))
	v754 = int64(55536)
	v755 = int64(0)
	v760 = int64(32)
	v763 = int64(base.Ui64(v752) >> (uint(v760) % 64))
	v766 = int64(4294967295)
	v769 = v752 & v766
	v770 = v754 * v769
	v774 = int64(base.Ui64(v770)>>(uint(v760)%64)) + v754*v763
	v781 = v769*v755 + v774&v766
	*(*int64)(unsafe.Add(mBase, uint32(v751)+8)) = v752*v755 + v753*v754 + v755*v763 + int64(base.Ui64(v774)>>(uint(v760)%64)) + int64(base.Ui64(v781)>>(uint(v760)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v751))) = v770&v766 | v781<<(uint(v760)%64)
	goto L121
L121:
	;
	v793 = v708 - int32(2)
	v794 = *(*int64)(unsafe.Add(mBase, uint32(v32)+160))
	v795 = v794 + v722
	*(*uint16)(unsafe.Add(mBase, uint32(v793))) = uint16(v795)
	v799 = int64(0)
	v804 = v709 + int32(1)
	if v723 == v799 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v805 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v722))
	goto L124
L123:
	;
	v805 = base.B2i32(v723 != v799)
	goto L124
L124:
	;
	if v805 != 0 {
		v708 = v793
		v709 = v804
		v722 = v752
		v723 = v753
		goto L118
	} else {
		goto L125
	}
L125:
	;
	goto L119
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+304)) = v839
	v842 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v839))) = uint16(v842)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+308)) = v839 + int32(2)
	if v583 < int64(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+292)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+288)) = v979
	v1005 = int32(0)
	v1013 = v566
	v1017 = v570
	v1018 = v1005
	v1019 = v1005
	goto L142
L128:
	;
	v875 = v839 + int32(22)
	v876 = v842
	v893 = v867
	v894 = v868
	goto L133
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+296)) = int64(16384)
	v852 = int64(0)
	v867 = v852 - v582
	v868 = v852 - (v583 + base.I64_extend_i32_u(base.B2i32(v582 != v852)))
	goto L128
L130:
	;
	goto L131
L131:
	;
	v860 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+296)) = v860
	if v582|v583 == v860 {
		v979 = v842
		v981 = int32(0)
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v867 = v582
	v868 = v583
	goto L128
L133:
	;
	v901 = v32 + int32(144)
	v904 = m.G0
	v905 = int32(16)
	v906 = v904 - v905
	m.G0 = v906
	F___udivmodti4(m, v906, v893, v894, int64(10000), int64(0))
	mBase = m.M
	v910 = *(*int64)(unsafe.Add(mBase, uint32(v906)+8))
	v911 = *(*int64)(unsafe.Add(mBase, uint32(v906)))
	*(*int64)(unsafe.Add(mBase, uint32(v901))) = v911
	*(*int64)(unsafe.Add(mBase, uint32(v901)+8)) = v910
	m.G0 = v906 + v905
	goto L135
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+308)) = v960
	v979 = v971
	v981 = v876
	goto L127
L135:
	;
	v918 = v32 + int32(128)
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v32)+144))
	v920 = *(*int64)(unsafe.Add(mBase, uint32(v32)+152))
	v921 = int64(55536)
	v922 = int64(0)
	v927 = int64(32)
	v930 = int64(base.Ui64(v919) >> (uint(v927) % 64))
	v933 = int64(4294967295)
	v936 = v919 & v933
	v937 = v921 * v936
	v941 = int64(base.Ui64(v937)>>(uint(v927)%64)) + v921*v930
	v948 = v936*v922 + v941&v933
	*(*int64)(unsafe.Add(mBase, uint32(v918)+8)) = v919*v922 + v920*v921 + v922*v930 + int64(base.Ui64(v941)>>(uint(v927)%64)) + int64(base.Ui64(v948)>>(uint(v927)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v918))) = v937&v933 | v948<<(uint(v927)%64)
	goto L136
L136:
	;
	v960 = v875 - int32(2)
	v961 = *(*int64)(unsafe.Add(mBase, uint32(v32)+128))
	v962 = v961 + v893
	*(*uint16)(unsafe.Add(mBase, uint32(v960))) = uint16(v962)
	v966 = int64(0)
	v971 = v876 + int32(1)
	if v894 == v966 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v972 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v893))
	goto L139
L138:
	;
	v972 = base.B2i32(v894 != v966)
	goto L139
L139:
	;
	if v972 != 0 {
		v875 = v960
		v876 = v971
		v893 = v919
		v894 = v920
		goto L133
	} else {
		goto L140
	}
L140:
	;
	goto L134
L141:
	;
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v32)+312))
	v3743 = v2607
	v3749 = v1448
	v3750 = v1238
	goto L58
L142:
	;
	v1038 = int32(2)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v1017<<(uint(v1038)%32))))
	v1044 = base.I32_div_s(v1041-v1013, v1038)
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1013 < v1045 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v32)+224))
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	if v2218 == int32(0) {
		goto L382
	} else {
		goto L383
	}
L144:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1256 = v1013 + v1044
	if v1256 < v1255 {
		goto L180
	} else {
		goto L181
	}
L145:
	;
	v1047 = v1045 - v1013
	if v1044 < v1047 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	if v1019 != 0 {
		goto L175
	} else {
		goto L176
	}
L148:
	;
	v1049 = v1044
	goto L150
L149:
	;
	v1049 = v1047
	goto L150
L150:
	;
	if v1019 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	F_pfree(m, v1019)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L8
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v1053 = v1049 << (uint(int32(1)) % 32)
	v1056 = F_palloc(m, v1053+int32(2))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L8
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+256)) = v1056
	v1059 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1056))) = uint16(v1059)
	v1062 = v1056 + int32(2)
	if v1053 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v1062, v1063+v1013<<(uint(int32(1))%32), v1053)
	goto L158
L157:
	;
	goto L158
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+248)) = int64(0)
	v1071 = v1044 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+244)) = v1071
	if int32(0) < v1049 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+240)) = v1191
	*(*int32)(unsafe.Add(mBase, uint32(v32)+260)) = v1190
	v1238 = v1056
	goto L144
L160:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+244)) = int64(0)
	v1190 = v1158
	v1191 = int32(0)
	goto L159
L161:
	;
	v1080 = v1062
	v1081 = v1049
	v1085 = v1071
	goto L164
L162:
	;
	goto L163
L163:
	;
	if v1049 != 0 {
		v1190 = v1062
		v1191 = v1049
		goto L159
	} else {
		goto L174
	}
L164:
	;
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1080))))
	if v1105 != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v1158 = v1062 + v1053
	goto L160
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+244)) = v1085
	v1112 = v1081
	goto L169
L167:
	;
	goto L168
L168:
	;
	v1146 = int32(1)
	if v1146 < v1081 {
		v1080 = v1080 + int32(2)
		v1081 = v1081 - v1146
		v1085 = v1085 - v1146
		goto L164
	} else {
		goto L173
	}
L169:
	;
	v1141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1080+v1112<<(uint(int32(1))%32)-int32(2)))))
	if v1141 != 0 {
		v1190 = v1080
		v1191 = v1112
		goto L159
	} else {
		goto L171
	}
L170:
	;
	v1158 = v1080
	goto L160
L171:
	;
	v1142 = int32(1)
	if v1142 < v1112 {
		v1112 = v1112 - v1142
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	goto L165
L174:
	;
	v1158 = v1062
	goto L160
L175:
	;
	F_pfree(m, v1019)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L8
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v1219 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+256)) = v1219
	*(*int64)(unsafe.Add(mBase, uint32(v32)+248)) = v1219
	*(*int64)(unsafe.Add(mBase, uint32(v32)+240)) = v1219
	v1238 = int32(0)
	goto L144
L178:
	;
	goto L177
L179:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v32)+288))
	v1468 = v1466 << (uint(int32(1)) % 32)
	v1471 = F_palloc(m, v1468+int32(2))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		goto L8
	} else {
		goto L214
	}
L180:
	;
	v1258 = v1255 - v1256
	if v1044 < v1258 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	if v1018 != 0 {
		goto L210
	} else {
		goto L211
	}
L183:
	;
	v1260 = v1044
	goto L185
L184:
	;
	v1260 = v1258
	goto L185
L185:
	;
	if v1018 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	F_pfree(m, v1018)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L8
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1264 = v1260 << (uint(int32(1)) % 32)
	v1267 = F_palloc(m, v1264+int32(2))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L8
	} else {
		goto L190
	}
L189:
	;
	goto L188
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+280)) = v1267
	v1270 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1267))) = uint16(v1270)
	v1273 = v1267 + int32(2)
	if v1264 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v1273, v1274+v1256<<(uint(int32(1))%32), v1264)
	goto L193
L192:
	;
	goto L193
L193:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+272)) = int64(0)
	v1282 = v1044 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+268)) = v1282
	if int32(0) < v1260 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+264)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v32)+284)) = v1401
	v1448 = v1267
	goto L179
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+268)) = int64(0)
	v1401 = v1369
	v1402 = int32(0)
	goto L194
L196:
	;
	v1291 = v1273
	v1292 = v1260
	v1296 = v1282
	goto L199
L197:
	;
	goto L198
L198:
	;
	if v1260 != 0 {
		v1401 = v1273
		v1402 = v1260
		goto L194
	} else {
		goto L209
	}
L199:
	;
	v1316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1291))))
	if v1316 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v1369 = v1273 + v1264
	goto L195
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+268)) = v1296
	v1323 = v1292
	goto L204
L202:
	;
	goto L203
L203:
	;
	v1357 = int32(1)
	if v1357 < v1292 {
		v1291 = v1291 + int32(2)
		v1292 = v1292 - v1357
		v1296 = v1296 - v1357
		goto L199
	} else {
		goto L208
	}
L204:
	;
	v1352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1291+v1323<<(uint(int32(1))%32)-int32(2)))))
	if v1352 != 0 {
		v1401 = v1291
		v1402 = v1323
		goto L194
	} else {
		goto L206
	}
L205:
	;
	v1369 = v1291
	goto L195
L206:
	;
	v1353 = int32(1)
	if v1353 < v1323 {
		v1323 = v1323 - v1353
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	goto L200
L209:
	;
	v1369 = v1273
	goto L195
L210:
	;
	F_pfree(m, v1018)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L8
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v1430 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+280)) = v1430
	*(*int64)(unsafe.Add(mBase, uint32(v32)+272)) = v1430
	*(*int64)(unsafe.Add(mBase, uint32(v32)+264)) = v1430
	v1448 = int32(0)
	goto L179
L213:
	;
	goto L212
L214:
	;
	v1473 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1471))) = uint16(v1473)
	if base.B2i32(v1468 == v1473)|base.B2i32(v1466 <= v1473) == v1473 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v32)+308))
	base.MemoryCopy(m, v1471+int32(2), v1484, v1468)
	goto L217
L216:
	;
	goto L217
L217:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	if v1486 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	F_pfree(m, v1486)
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L8
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1489 = *(*int64)(unsafe.Add(mBase, uint32(v32)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v1489
	v1491 = *(*int64)(unsafe.Add(mBase, uint32(v32)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v1491
	*(*int32)(unsafe.Add(mBase, uint32(v32)+232)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v32)+236)) = v1471 + int32(2)
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+220)) = v1497 + v1044
	v1501 = v32 + int32(216)
	F_add_var(m, v1501, v32+int32(240), v1501)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L8
	} else {
		goto L222
	}
L221:
	;
	goto L220
L222:
	;
	v1507 = v32 + int32(312)
	v1509 = v32 + int32(192)
	F_add_var(m, v1507, v1507, v1509)
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L8
	} else {
		goto L223
	}
L223:
	;
	v1512 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+488)) = v1512
	*(*int64)(unsafe.Add(mBase, uint32(v32)+496)) = v1512
	*(*int64)(unsafe.Add(mBase, uint32(v32)+504)) = v1512
	*(*int64)(unsafe.Add(mBase, uint32(v32)+464)) = v1512
	*(*int64)(unsafe.Add(mBase, uint32(v32)+472)) = v1512
	*(*int64)(unsafe.Add(mBase, uint32(v32)+480)) = v1512
	v1525 = v32 + int32(488)
	v1526 = int32(0)
	F_div_var(m, v1501, v1509, v1525, v1526, v1526, v1526)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L8
	} else {
		goto L224
	}
L224:
	;
	v1532 = v32 + int32(464)
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v32)+204))
	F_mul_var(m, v1509, v1525, v1532, v1533)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L8
	} else {
		goto L225
	}
L225:
	;
	F_sub_var(m, v1501, v1532, v1532)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L8
	} else {
		goto L226
	}
L226:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v32)+464))
	if v1538 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v32)+484))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v32)+468))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v32)+212))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	v1669 = int32(0)
	if base.B2i32(v1668 < v1665)&base.B2i32(v1669 < v1640) == v1669 {
		v1700 = v1665
		v1704 = v1669
		goto L245
	} else {
		goto L246
	}
L228:
	;
	v1640 = int32(0)
	goto L227
L229:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v32)+224))
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v32)+472))
	if v1541 == v1542 {
		v1640 = v1538
		goto L227
	} else {
		goto L230
	}
L230:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	goto L231
L231:
	;
	if base.B2i32(v1541 != v1544) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	v1640 = v1600
	goto L227
L233:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v32)+464))
	if v1600 == int32(0) {
		goto L228
	} else {
		goto L241
	}
L234:
	;
	v1578 = v32 + int32(488)
	F_sub_var(m, v1578, int32(_a_F_sqrt_var_5), v1578)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L8
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1589 = v32 + int32(488)
	F_add_var(m, v1589, int32(_a_F_sqrt_var_5), v1589)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L8
	} else {
		goto L239
	}
L237:
	;
	v1583 = v32 + int32(464)
	F_add_var(m, v1583, v32+int32(192), v1583)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L8
	} else {
		goto L238
	}
L238:
	;
	goto L233
L239:
	;
	v1594 = v32 + int32(464)
	F_sub_var(m, v1594, v32+int32(192), v1594)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L8
	} else {
		goto L240
	}
L240:
	;
	goto L233
L241:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v32)+472))
	if v1541 != v1603 {
		goto L231
	} else {
		goto L242
	}
L242:
	;
	goto L232
L243:
	;
	if int32(0) <= v1842 {
		goto L286
	} else {
		goto L287
	}
L244:
	;
	v1842 = v1832
	goto L243
L245:
	;
	if base.B2i32(v1667 <= int32(0))|base.B2i32(v1668 <= v1700) != 0 {
		v1736 = v1668
		v1738 = v1669
		goto L252
	} else {
		goto L253
	}
L246:
	;
	v1681 = v1665
	v1685 = v1669
	goto L247
L247:
	;
	v1691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1664+v1685<<(uint(int32(1))%32)))))
	if v1691 != 0 {
		v1832 = int32(1)
		goto L244
	} else {
		goto L249
	}
L248:
	;
	v1700 = v1695
	v1704 = v1693
	goto L245
L249:
	;
	v1692 = int32(1)
	v1693 = v1685 + v1692
	v1695 = v1681 - v1692
	if v1695 <= v1668 {
		v1700 = v1695
		v1704 = v1693
		goto L245
	} else {
		goto L250
	}
L250:
	;
	if v1693 < v1640 {
		v1681 = v1695
		v1685 = v1693
		goto L247
	} else {
		goto L251
	}
L251:
	;
	goto L248
L252:
	;
	if v1700 != v1736 {
		v1778 = v1704
		v1779 = v1738
		goto L259
	} else {
		goto L260
	}
L253:
	;
	v1717 = v1668
	v1719 = v1669
	goto L254
L254:
	;
	v1724 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1666+v1719<<(uint(int32(1))%32)))))
	if v1724 != 0 {
		v1832 = int32(-1)
		goto L244
	} else {
		goto L256
	}
L255:
	;
	v1736 = v1728
	v1738 = v1726
	goto L252
L256:
	;
	v1725 = int32(1)
	v1726 = v1719 + v1725
	v1728 = v1717 - v1725
	if v1728 <= v1700 {
		v1736 = v1728
		v1738 = v1726
		goto L252
	} else {
		goto L257
	}
L257:
	;
	if v1726 < v1667 {
		v1717 = v1728
		v1719 = v1726
		goto L254
	} else {
		goto L258
	}
L258:
	;
	goto L255
L259:
	;
	if v1640 < v1778 {
		goto L268
	} else {
		goto L269
	}
L260:
	;
	v1747 = v1704
	v1748 = v1738
	goto L261
L261:
	;
	if base.B2i32(v1640 <= v1747)|base.B2i32(v1667 <= v1748) != 0 {
		v1778 = v1747
		v1779 = v1748
		goto L259
	} else {
		goto L263
	}
L262:
	;
	if base.I32_extend16_s(v1764) < base.I32_extend16_s(v1762) {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v1753 = int32(1)
	v1762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1664+v1747<<(uint(v1753)%32)))))
	v1764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1748<<(uint(v1753)%32)+v1666))))
	if v1762 == v1764 {
		v1747 = v1747 + v1753
		v1748 = v1748 + v1753
		goto L261
	} else {
		goto L264
	}
L264:
	;
	goto L262
L265:
	;
	v1771 = int32(1)
	goto L267
L266:
	;
	v1771 = int32(-1)
	goto L267
L267:
	;
	v1842 = v1771
	goto L243
L268:
	;
	v1782 = v1778
	goto L270
L269:
	;
	v1782 = v1640
	goto L270
L270:
	;
	v1789 = v1778
	goto L271
L271:
	;
	if v1782 == v1789 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1832 = v1815
	goto L244
L273:
	;
	if v1667 < v1779 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	v1815 = int32(1)
	v1821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1664+v1789<<(uint(v1815)%32)))))
	if v1821 == int32(0) {
		v1789 = v1789 + v1815
		goto L271
	} else {
		goto L285
	}
L276:
	;
	v1794 = v1779
	goto L278
L277:
	;
	v1794 = v1667
	goto L278
L278:
	;
	v1802 = v1779
	goto L279
L279:
	;
	if v1794 == v1802 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1832 = int32(-1)
	goto L244
L281:
	;
	v1842 = int32(0)
	goto L243
L282:
	;
	goto L283
L283:
	;
	v1806 = int32(1)
	v1811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1802<<(uint(v1806)%32)+v1666))))
	if v1811 == int32(0) {
		v1802 = v1802 + v1806
		goto L279
	} else {
		goto L284
	}
L284:
	;
	goto L280
L285:
	;
	goto L272
L286:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v32)+224))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	goto L289
L287:
	;
	v2085 = v1664
	v2086 = v1640
	goto L288
L288:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v32)+488))
	v2112 = v2110 << (uint(int32(1)) % 32)
	v2115 = F_palloc(m, v2112+int32(2))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L8
	} else {
		goto L343
	}
L289:
	;
	if base.B2i32(v1845 != v1846) == int32(0) {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v2085 = v1902
	v2086 = v1903
	goto L288
L291:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v32)+484))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v32)+464))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v32)+468))
	v1905 = int32(0)
	if base.B2i32(v1668 < v1904)&base.B2i32(v1905 < v1903) == v1905 {
		v1936 = v1904
		v1940 = v1905
		goto L301
	} else {
		goto L302
	}
L292:
	;
	v1880 = v32 + int32(488)
	F_add_var(m, v1880, int32(_a_F_sqrt_var_5), v1880)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L8
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v1891 = v32 + int32(488)
	F_sub_var(m, v1891, int32(_a_F_sqrt_var_5), v1891)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L8
	} else {
		goto L297
	}
L295:
	;
	v1885 = v32 + int32(464)
	F_sub_var(m, v1885, v32+int32(192), v1885)
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L8
	} else {
		goto L296
	}
L296:
	;
	goto L291
L297:
	;
	v1896 = v32 + int32(464)
	F_add_var(m, v1896, v32+int32(192), v1896)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L8
	} else {
		goto L298
	}
L298:
	;
	goto L291
L299:
	;
	if int32(0) <= v2078 {
		goto L289
	} else {
		goto L342
	}
L300:
	;
	v2078 = v2068
	goto L299
L301:
	;
	if base.B2i32(v1667 <= int32(0))|base.B2i32(v1668 <= v1936) != 0 {
		v1972 = v1668
		v1974 = v1905
		goto L308
	} else {
		goto L309
	}
L302:
	;
	v1917 = v1904
	v1921 = v1905
	goto L303
L303:
	;
	v1927 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1902+v1921<<(uint(int32(1))%32)))))
	if v1927 != 0 {
		v2068 = int32(1)
		goto L300
	} else {
		goto L305
	}
L304:
	;
	v1936 = v1931
	v1940 = v1929
	goto L301
L305:
	;
	v1928 = int32(1)
	v1929 = v1921 + v1928
	v1931 = v1917 - v1928
	if v1931 <= v1668 {
		v1936 = v1931
		v1940 = v1929
		goto L301
	} else {
		goto L306
	}
L306:
	;
	if v1929 < v1903 {
		v1917 = v1931
		v1921 = v1929
		goto L303
	} else {
		goto L307
	}
L307:
	;
	goto L304
L308:
	;
	if v1936 != v1972 {
		v2014 = v1940
		v2015 = v1974
		goto L315
	} else {
		goto L316
	}
L309:
	;
	v1953 = v1668
	v1955 = v1905
	goto L310
L310:
	;
	v1960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1666+v1955<<(uint(int32(1))%32)))))
	if v1960 != 0 {
		v2068 = int32(-1)
		goto L300
	} else {
		goto L312
	}
L311:
	;
	v1972 = v1964
	v1974 = v1962
	goto L308
L312:
	;
	v1961 = int32(1)
	v1962 = v1955 + v1961
	v1964 = v1953 - v1961
	if v1964 <= v1936 {
		v1972 = v1964
		v1974 = v1962
		goto L308
	} else {
		goto L313
	}
L313:
	;
	if v1962 < v1667 {
		v1953 = v1964
		v1955 = v1962
		goto L310
	} else {
		goto L314
	}
L314:
	;
	goto L311
L315:
	;
	if v1903 < v2014 {
		goto L324
	} else {
		goto L325
	}
L316:
	;
	v1983 = v1940
	v1984 = v1974
	goto L317
L317:
	;
	if base.B2i32(v1903 <= v1983)|base.B2i32(v1667 <= v1984) != 0 {
		v2014 = v1983
		v2015 = v1984
		goto L315
	} else {
		goto L319
	}
L318:
	;
	if base.I32_extend16_s(v2000) < base.I32_extend16_s(v1998) {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	v1989 = int32(1)
	v1998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1902+v1983<<(uint(v1989)%32)))))
	v2000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1984<<(uint(v1989)%32)+v1666))))
	if v1998 == v2000 {
		v1983 = v1983 + v1989
		v1984 = v1984 + v1989
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	v2007 = int32(1)
	goto L323
L322:
	;
	v2007 = int32(-1)
	goto L323
L323:
	;
	v2078 = v2007
	goto L299
L324:
	;
	v2018 = v2014
	goto L326
L325:
	;
	v2018 = v1903
	goto L326
L326:
	;
	v2025 = v2014
	goto L327
L327:
	;
	if v2018 == v2025 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v2068 = v2051
	goto L300
L329:
	;
	if v1667 < v2015 {
		goto L332
	} else {
		goto L333
	}
L330:
	;
	goto L331
L331:
	;
	v2051 = int32(1)
	v2057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1902+v2025<<(uint(v2051)%32)))))
	if v2057 == int32(0) {
		v2025 = v2025 + v2051
		goto L327
	} else {
		goto L341
	}
L332:
	;
	v2030 = v2015
	goto L334
L333:
	;
	v2030 = v1667
	goto L334
L334:
	;
	v2038 = v2015
	goto L335
L335:
	;
	if v2030 == v2038 {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v2068 = int32(-1)
	goto L300
L337:
	;
	v2078 = int32(0)
	goto L299
L338:
	;
	goto L339
L339:
	;
	v2042 = int32(1)
	v2047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2038<<(uint(v2042)%32)+v1666))))
	if v2047 == int32(0) {
		v2038 = v2038 + v2042
		goto L335
	} else {
		goto L340
	}
L340:
	;
	goto L336
L341:
	;
	goto L328
L342:
	;
	goto L290
L343:
	;
	v2117 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2115))) = uint16(v2117)
	if base.B2i32(v2112 == v2117)|base.B2i32(v2110 <= v2117) == v2117 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v32)+508))
	base.MemoryCopy(m, v2115+int32(2), v2128, v2112)
	goto L346
L345:
	;
	goto L346
L346:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	if v2130 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	F_pfree(m, v2130)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L8
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v2133 = *(*int64)(unsafe.Add(mBase, uint32(v32)+496))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v2133
	v2135 = *(*int64)(unsafe.Add(mBase, uint32(v32)+488))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v2135
	*(*int32)(unsafe.Add(mBase, uint32(v32)+232)) = v2115
	v2138 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+236)) = v2115 + v2138
	v2142 = v2086 << (uint(int32(1)) % 32)
	v2145 = F_palloc(m, v2142+v2138)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L8
	} else {
		goto L351
	}
L350:
	;
	goto L349
L351:
	;
	v2147 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2145))) = uint16(v2147)
	if base.B2i32(v2142 == v2147)|base.B2i32(v2086 <= v2147) == v2147 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	base.MemoryCopy(m, v2145+int32(2), v2085, v2142)
	goto L354
L353:
	;
	goto L354
L354:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	if v2159 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	F_pfree(m, v2159)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L8
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v32)+472))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = v2162
	v2164 = *(*int64)(unsafe.Add(mBase, uint32(v32)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v2164
	*(*int32)(unsafe.Add(mBase, uint32(v32)+208)) = v2145
	*(*int32)(unsafe.Add(mBase, uint32(v32)+212)) = v2145 + int32(2)
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v32)+504))
	if v2170 != 0 {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	goto L357
L359:
	;
	F_pfree(m, v2170)
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L8
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v32)+480))
	if v2173 != 0 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	goto L361
L363:
	;
	F_pfree(m, v2173)
	mBase = m.M
	v2175 = m.ExcPending
	if v2175 != 0 {
		goto L8
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v32)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v2176 + v1044
	v2180 = v32 + int32(312)
	v2182 = v32 + int32(216)
	F_add_var(m, v2180, v2182, v2180)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L8
	} else {
		goto L367
	}
L366:
	;
	goto L365
L367:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+196)) = v2185 + v1044
	v2189 = v32 + int32(192)
	F_add_var(m, v2189, v32+int32(264), v2189)
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L8
	} else {
		goto L368
	}
L368:
	;
	F_mul_var(m, v2182, v2182, v2182, int32(0))
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L8
	} else {
		goto L369
	}
L369:
	;
	if v1017 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v2198 = v32 + int32(288)
	F_sub_var(m, v2189, v2182, v2198)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L8
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	goto L143
L373:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
	if v2201 == int32(_a_F_sqrt_var_6) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	F_add_var(m, v2198, v2180, v2198)
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L8
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	if int32(0) < v1017 {
		v1013 = v1044 + v1256
		v1017 = v1017 - int32(1)
		v1018 = v1448
		v1019 = v1238
		goto L142
	} else {
		goto L380
	}
L377:
	;
	F_sub_var(m, v2180, int32(_a_F_sqrt_var_5), v2180)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L8
	} else {
		goto L378
	}
L378:
	;
	F_add_var(m, v2198, v2180, v2198)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L8
	} else {
		goto L379
	}
L379:
	;
	goto L376
L380:
	;
	goto L141
L381:
	;
	v2595 = v32 + int32(312)
	F_sub_var(m, v2595, int32(_a_F_sqrt_var_5), v2595)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L8
	} else {
		goto L484
	}
L382:
	;
	if v2217 == int32(0) {
		goto L141
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	if v2217 == int32(0) {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	if v2216 != int32(_a_F_sqrt_var_6) {
		goto L381
	} else {
		goto L386
	}
L386:
	;
	goto L141
L387:
	;
	if v2225 != 0 {
		goto L381
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v32)+236))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v32)+212))
	if v2225 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	goto L141
L391:
	;
	if int32(0) <= v2586 {
		goto L141
	} else {
		goto L483
	}
L392:
	;
	if v2216 == int32(_a_F_sqrt_var_6) {
		goto L141
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	if v2216 == int32(0) {
		goto L381
	} else {
		goto L439
	}
L395:
	;
	v2236 = int32(0)
	if base.B2i32(v2228 < v2230)&base.B2i32(v2236 < v2218) == v2236 {
		v2267 = v2230
		v2271 = v2236
		goto L398
	} else {
		goto L399
	}
L396:
	;
	v2586 = v2409
	goto L391
L397:
	;
	v2409 = v2399
	goto L396
L398:
	;
	if base.B2i32(v2217 <= int32(0))|base.B2i32(v2228 <= v2267) != 0 {
		v2303 = v2228
		v2305 = v2236
		goto L405
	} else {
		goto L406
	}
L399:
	;
	v2248 = v2230
	v2252 = v2236
	goto L400
L400:
	;
	v2258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2231+v2252<<(uint(int32(1))%32)))))
	if v2258 != 0 {
		v2399 = int32(1)
		goto L397
	} else {
		goto L402
	}
L401:
	;
	v2267 = v2262
	v2271 = v2260
	goto L398
L402:
	;
	v2259 = int32(1)
	v2260 = v2252 + v2259
	v2262 = v2248 - v2259
	if v2262 <= v2228 {
		v2267 = v2262
		v2271 = v2260
		goto L398
	} else {
		goto L403
	}
L403:
	;
	if v2260 < v2218 {
		v2248 = v2262
		v2252 = v2260
		goto L400
	} else {
		goto L404
	}
L404:
	;
	goto L401
L405:
	;
	if v2267 != v2303 {
		v2345 = v2271
		v2346 = v2305
		goto L412
	} else {
		goto L413
	}
L406:
	;
	v2284 = v2228
	v2286 = v2236
	goto L407
L407:
	;
	v2291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2229+v2286<<(uint(int32(1))%32)))))
	if v2291 != 0 {
		v2399 = int32(-1)
		goto L397
	} else {
		goto L409
	}
L408:
	;
	v2303 = v2295
	v2305 = v2293
	goto L405
L409:
	;
	v2292 = int32(1)
	v2293 = v2286 + v2292
	v2295 = v2284 - v2292
	if v2295 <= v2267 {
		v2303 = v2295
		v2305 = v2293
		goto L405
	} else {
		goto L410
	}
L410:
	;
	if v2293 < v2217 {
		v2284 = v2295
		v2286 = v2293
		goto L407
	} else {
		goto L411
	}
L411:
	;
	goto L408
L412:
	;
	if v2218 < v2345 {
		goto L421
	} else {
		goto L422
	}
L413:
	;
	v2314 = v2271
	v2315 = v2305
	goto L414
L414:
	;
	if base.B2i32(v2218 <= v2314)|base.B2i32(v2217 <= v2315) != 0 {
		v2345 = v2314
		v2346 = v2315
		goto L412
	} else {
		goto L416
	}
L415:
	;
	if base.I32_extend16_s(v2331) < base.I32_extend16_s(v2329) {
		goto L418
	} else {
		goto L419
	}
L416:
	;
	v2320 = int32(1)
	v2329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2231+v2314<<(uint(v2320)%32)))))
	v2331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2315<<(uint(v2320)%32)+v2229))))
	if v2329 == v2331 {
		v2314 = v2314 + v2320
		v2315 = v2315 + v2320
		goto L414
	} else {
		goto L417
	}
L417:
	;
	goto L415
L418:
	;
	v2338 = int32(1)
	goto L420
L419:
	;
	v2338 = int32(-1)
	goto L420
L420:
	;
	v2409 = v2338
	goto L396
L421:
	;
	v2349 = v2345
	goto L423
L422:
	;
	v2349 = v2218
	goto L423
L423:
	;
	v2356 = v2345
	goto L424
L424:
	;
	if v2349 == v2356 {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	v2399 = v2382
	goto L397
L426:
	;
	if v2217 < v2346 {
		goto L429
	} else {
		goto L430
	}
L427:
	;
	goto L428
L428:
	;
	v2382 = int32(1)
	v2388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2231+v2356<<(uint(v2382)%32)))))
	if v2388 == int32(0) {
		v2356 = v2356 + v2382
		goto L424
	} else {
		goto L438
	}
L429:
	;
	v2361 = v2346
	goto L431
L430:
	;
	v2361 = v2217
	goto L431
L431:
	;
	v2369 = v2346
	goto L432
L432:
	;
	if v2361 == v2369 {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	v2399 = int32(-1)
	goto L397
L434:
	;
	v2409 = int32(0)
	goto L396
L435:
	;
	goto L436
L436:
	;
	v2373 = int32(1)
	v2378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2369<<(uint(v2373)%32)+v2229))))
	if v2378 == int32(0) {
		v2369 = v2369 + v2373
		goto L432
	} else {
		goto L437
	}
L437:
	;
	goto L433
L438:
	;
	goto L425
L439:
	;
	v2412 = int32(0)
	if base.B2i32(v2230 < v2228)&base.B2i32(v2412 < v2217) == v2412 {
		v2443 = v2228
		v2447 = v2412
		goto L442
	} else {
		goto L443
	}
L440:
	;
	v2586 = v2585
	goto L391
L441:
	;
	v2585 = v2575
	goto L440
L442:
	;
	if base.B2i32(v2218 <= int32(0))|base.B2i32(v2230 <= v2443) != 0 {
		v2479 = v2230
		v2481 = v2412
		goto L449
	} else {
		goto L450
	}
L443:
	;
	v2424 = v2228
	v2428 = v2412
	goto L444
L444:
	;
	v2434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2229+v2428<<(uint(int32(1))%32)))))
	if v2434 != 0 {
		v2575 = int32(1)
		goto L441
	} else {
		goto L446
	}
L445:
	;
	v2443 = v2438
	v2447 = v2436
	goto L442
L446:
	;
	v2435 = int32(1)
	v2436 = v2428 + v2435
	v2438 = v2424 - v2435
	if v2438 <= v2230 {
		v2443 = v2438
		v2447 = v2436
		goto L442
	} else {
		goto L447
	}
L447:
	;
	if v2436 < v2217 {
		v2424 = v2438
		v2428 = v2436
		goto L444
	} else {
		goto L448
	}
L448:
	;
	goto L445
L449:
	;
	if v2443 != v2479 {
		v2521 = v2447
		v2522 = v2481
		goto L456
	} else {
		goto L457
	}
L450:
	;
	v2460 = v2230
	v2462 = v2412
	goto L451
L451:
	;
	v2467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2231+v2462<<(uint(int32(1))%32)))))
	if v2467 != 0 {
		v2575 = int32(-1)
		goto L441
	} else {
		goto L453
	}
L452:
	;
	v2479 = v2471
	v2481 = v2469
	goto L449
L453:
	;
	v2468 = int32(1)
	v2469 = v2462 + v2468
	v2471 = v2460 - v2468
	if v2471 <= v2443 {
		v2479 = v2471
		v2481 = v2469
		goto L449
	} else {
		goto L454
	}
L454:
	;
	if v2469 < v2218 {
		v2460 = v2471
		v2462 = v2469
		goto L451
	} else {
		goto L455
	}
L455:
	;
	goto L452
L456:
	;
	if v2217 < v2521 {
		goto L465
	} else {
		goto L466
	}
L457:
	;
	v2490 = v2447
	v2491 = v2481
	goto L458
L458:
	;
	if base.B2i32(v2217 <= v2490)|base.B2i32(v2218 <= v2491) != 0 {
		v2521 = v2490
		v2522 = v2491
		goto L456
	} else {
		goto L460
	}
L459:
	;
	if base.I32_extend16_s(v2507) < base.I32_extend16_s(v2505) {
		goto L462
	} else {
		goto L463
	}
L460:
	;
	v2496 = int32(1)
	v2505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2229+v2490<<(uint(v2496)%32)))))
	v2507 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2491<<(uint(v2496)%32)+v2231))))
	if v2505 == v2507 {
		v2490 = v2490 + v2496
		v2491 = v2491 + v2496
		goto L458
	} else {
		goto L461
	}
L461:
	;
	goto L459
L462:
	;
	v2514 = int32(1)
	goto L464
L463:
	;
	v2514 = int32(-1)
	goto L464
L464:
	;
	v2585 = v2514
	goto L440
L465:
	;
	v2525 = v2521
	goto L467
L466:
	;
	v2525 = v2217
	goto L467
L467:
	;
	v2532 = v2521
	goto L468
L468:
	;
	if v2525 == v2532 {
		goto L470
	} else {
		goto L471
	}
L469:
	;
	v2575 = v2558
	goto L441
L470:
	;
	if v2218 < v2522 {
		goto L473
	} else {
		goto L474
	}
L471:
	;
	goto L472
L472:
	;
	v2558 = int32(1)
	v2564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2229+v2532<<(uint(v2558)%32)))))
	if v2564 == int32(0) {
		v2532 = v2532 + v2558
		goto L468
	} else {
		goto L482
	}
L473:
	;
	v2537 = v2522
	goto L475
L474:
	;
	v2537 = v2218
	goto L475
L475:
	;
	v2545 = v2522
	goto L476
L476:
	;
	if v2537 == v2545 {
		goto L478
	} else {
		goto L479
	}
L477:
	;
	v2575 = int32(-1)
	goto L441
L478:
	;
	v2585 = int32(0)
	goto L440
L479:
	;
	goto L480
L480:
	;
	v2549 = int32(1)
	v2554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2545<<(uint(v2549)%32)+v2231))))
	if v2554 == int32(0) {
		v2545 = v2545 + v2549
		goto L476
	} else {
		goto L481
	}
L481:
	;
	goto L477
L482:
	;
	goto L469
L483:
	;
	goto L381
L484:
	;
	goto L141
L485:
	;
	v2616 = v667
	v2633 = v663
	v2638 = v665
	goto L91
L486:
	;
	v2649 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2616<<(uint(int32(1))%32)))))
	v2672 = v2642 + v2649
	v2677 = v2644
	goto L90
L487:
	;
	v2817 = v2680 + v599
	v2831 = v2801
	v2832 = v2672
	v2837 = v2677
	goto L89
L488:
	;
	v2774 = v2764 * int64(10000)
	if v34 <= v2748 {
		v2801 = v2774
		goto L487
	} else {
		goto L502
	}
L489:
	;
	v2748 = v2680
	v2764 = int64(0)
	goto L488
L490:
	;
	goto L491
L491:
	;
	v2684 = int64(0)
	v2690 = v2680
	v2706 = v2684
	v2713 = v2684
	goto L492
L492:
	;
	v2716 = v2706 * int64(10000)
	if v2690 < v34 {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	if v599&int32(1) == int32(0) {
		v2801 = v2734
		goto L487
	} else {
		goto L501
	}
L494:
	;
	v2721 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2690<<(uint(int32(1))%32)))))
	v2723 = v2716 + v2721
	goto L496
L495:
	;
	v2723 = v2716
	goto L496
L496:
	;
	v2725 = v2723 * int64(10000)
	v2727 = v2690 + int32(1)
	if v2727 < v34 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v2732 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2727<<(uint(int32(1))%32)))))
	v2734 = v2725 + v2732
	goto L499
L498:
	;
	v2734 = v2725
	goto L499
L499:
	;
	v2736 = v2690 + int32(2)
	v2738 = v2713 + int64(2)
	if v2738 != v605&int64(1073741822) {
		v2690 = v2736
		v2706 = v2734
		v2713 = v2738
		goto L492
	} else {
		goto L500
	}
L500:
	;
	goto L493
L501:
	;
	v2748 = v2736
	v2764 = v2734
	goto L488
L502:
	;
	v2779 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2748<<(uint(int32(1))%32)))))
	v2801 = v2774 + v2779
	goto L487
L503:
	;
	v2881 = v32 + int32(96)
	v2882 = *(*int64)(unsafe.Add(mBase, uint32(v32)+112))
	v2883 = v2882 + v2832
	v2886 = *(*int64)(unsafe.Add(mBase, uint32(v32)+120))
	v2887 = int64(63)
	v2890 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v2883) < base.Ui64(v2882))) + (v2886 + v2832>>(uint(v2887)%64))
	v2891 = int64(1)
	v2892 = v578 << (uint(v2891) % 64)
	v2897 = v579<<(uint(v2891)%64) | int64(base.Ui64(v578)>>(uint(v2887)%64))
	v2901 = m.G0
	v2902 = int32(16)
	v2903 = v2901 - v2902
	m.G0 = v2903
	v2906 = v2890 >> (uint(v2887) % 64)
	v2907 = v2883 ^ v2906
	v2910 = v2907 + int64(base.Ui64(v2890)>>(uint(v2887)%64))
	v2916 = v2897 >> (uint(v2887) % 64)
	v2917 = v2916 ^ v2892
	v2920 = v2917 + int64(base.Ui64(v2897)>>(uint(v2887)%64))
	F___udivmodti4(m, v2903, v2910, base.I64_extend_i32_u(base.B2i32(base.Ui64(v2910) < base.Ui64(v2907)))+(v2906^v2890), v2920, base.I64_extend_i32_u(base.B2i32(base.Ui64(v2920) < base.Ui64(v2917)))+(v2916^v2897))
	mBase = m.M
	v2926 = *(*int64)(unsafe.Add(mBase, uint32(v2903)+8))
	v2927 = v2906 ^ v2916
	v2928 = *(*int64)(unsafe.Add(mBase, uint32(v2903)))
	v2929 = v2927 ^ v2928
	*(*int64)(unsafe.Add(mBase, uint32(v2881))) = v2929 - v2927
	*(*int64)(unsafe.Add(mBase, uint32(v2881)+8)) = v2927 ^ v2926 - v2927 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v2929) < base.Ui64(v2927)))
	m.G0 = v2903 + v2902
	goto L504
L504:
	;
	v2942 = v32 + int32(80)
	v2943 = *(*int64)(unsafe.Add(mBase, uint32(v32)+96))
	v2944 = *(*int64)(unsafe.Add(mBase, uint32(v32)+104))
	v2949 = int64(32)
	v2950 = int64(base.Ui64(v2892) >> (uint(v2949) % 64))
	v2952 = int64(base.Ui64(v2943) >> (uint(v2949) % 64))
	v2955 = int64(4294967295)
	v2956 = v2892 & v2955
	v2958 = v2943 & v2955
	v2959 = v2956 * v2958
	v2963 = int64(base.Ui64(v2959)>>(uint(v2949)%64)) + v2956*v2952
	v2970 = v2958*v2950 + v2963&v2955
	*(*int64)(unsafe.Add(mBase, uint32(v2942)+8)) = v2943*v2897 + v2944*v2892 + v2950*v2952 + int64(base.Ui64(v2963)>>(uint(v2949)%64)) + int64(base.Ui64(v2970)>>(uint(v2949)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2942))) = v2959&v2955 | v2970<<(uint(v2949)%64)
	goto L505
L505:
	;
	v2982 = v32 + int32(32)
	v2987 = int64(32)
	v2988 = int64(base.Ui64(v2837) >> (uint(v2987) % 64))
	v2990 = int64(base.Ui64(v578) >> (uint(v2987) % 64))
	v2993 = int64(4294967295)
	v2994 = v2837 & v2993
	v2996 = v578 & v2993
	v2997 = v2994 * v2996
	v3001 = int64(base.Ui64(v2997)>>(uint(v2987)%64)) + v2994*v2990
	v3008 = v2996*v2988 + v3001&v2993
	*(*int64)(unsafe.Add(mBase, uint32(v2982)+8)) = v578*v2843 + v579*v2837 + v2988*v2990 + int64(base.Ui64(v3001)>>(uint(v2987)%64)) + int64(base.Ui64(v3008)>>(uint(v2987)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2982))) = v2997&v2993 | v3008<<(uint(v2987)%64)
	goto L506
L506:
	;
	v3020 = v32 - int32(-64)
	v3021 = *(*int64)(unsafe.Add(mBase, uint32(v32)+80))
	v3022 = v2883 - v3021
	v3023 = *(*int64)(unsafe.Add(mBase, uint32(v32)+88))
	v3032 = int64(32)
	v3033 = int64(base.Ui64(v2837) >> (uint(v3032) % 64))
	v3035 = int64(base.Ui64(v3022) >> (uint(v3032) % 64))
	v3038 = int64(4294967295)
	v3039 = v2837 & v3038
	v3041 = v3022 & v3038
	v3042 = v3039 * v3041
	v3046 = int64(base.Ui64(v3042)>>(uint(v3032)%64)) + v3039*v3035
	v3053 = v3041*v3033 + v3046&v3038
	*(*int64)(unsafe.Add(mBase, uint32(v3020)+8)) = v3022*v2843 + (v2890-v3023-base.I64_extend_i32_u(base.B2i32(base.Ui64(v2883) < base.Ui64(v3021))))*v2837 + v3033*v3035 + int64(base.Ui64(v3046)>>(uint(v3032)%64)) + int64(base.Ui64(v3053)>>(uint(v3032)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3020))) = v3042&v3038 | v3053<<(uint(v3032)%64)
	goto L507
L507:
	;
	v3065 = v32 + int32(48)
	v3067 = v2943 * v2944
	v3070 = int64(32)
	v3071 = int64(base.Ui64(v2943) >> (uint(v3070) % 64))
	v3076 = int64(4294967295)
	v3077 = v2943 & v3076
	v3080 = v3077 * v3077
	v3083 = v3077 * v3071
	v3084 = int64(base.Ui64(v3080)>>(uint(v3070)%64)) + v3083
	v3091 = v3083 + v3084&v3076
	*(*int64)(unsafe.Add(mBase, uint32(v3065)+8)) = v3067 + v3067 + v3071*v3071 + int64(base.Ui64(v3084)>>(uint(v3070)%64)) + int64(base.Ui64(v3091)>>(uint(v3070)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3065))) = v3080&v3076 | v3091<<(uint(v3070)%64)
	goto L508
L508:
	;
	v3102 = *(*int64)(unsafe.Add(mBase, uint32(v32)+72))
	v3103 = int64(63)
	v3105 = *(*int64)(unsafe.Add(mBase, uint32(v32)+56))
	v3107 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
	v3112 = v2831 - v3107
	v3113 = *(*int64)(unsafe.Add(mBase, uint32(v32)+64))
	v3114 = v3112 + v3113
	v3117 = v3102 + (v2831>>(uint(v3103)%64) - v3105 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v2831) < base.Ui64(v3107)))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v3114) < base.Ui64(v3112)))
	v3119 = v3117 >> (uint(v3103) % 64)
	v3120 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
	v3121 = v2943 + v3120
	v3124 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
	v3126 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v3121) < base.Ui64(v2943))) + (v2944 + v3124)
	v3127 = int64(0)
	v3130 = v3126 - base.I64_extend_i32_u(base.B2i32(v3121 == v3127))
	v3133 = v3121 - int64(1)
	v3134 = v3133 + v3121
	v3141 = v3114 + v3134&v3119
	v3146 = base.B2i32(v3117 < v3127)
	if v3117 < v3127 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v3147 = v3130
	goto L511
L510:
	;
	v3147 = v3126
	goto L511
L511:
	;
	if v3117 < v3127 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v3148 = v3133
	goto L514
L513:
	;
	v3148 = v3121
	goto L514
L514:
	;
	if int32(0) < v570 {
		v566 = v2817
		v570 = v570 - int32(1)
		v578 = v3148
		v579 = v3147
		v582 = v3141
		v583 = v3117 + v3119&(v3130+v3126+base.I64_extend_i32_u(base.B2i32(base.Ui64(v3134) < base.Ui64(v3133)))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v3141) < base.Ui64(v3114)))
		goto L87
	} else {
		goto L515
	}
L515:
	;
	goto L88
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+328)) = v3154
	v3157 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3154))) = uint16(v3157)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3154 + int32(2)
	if v3147 < int64(0) {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v3294
	v3711 = v3292
	goto L59
L518:
	;
	v3190 = v3154 + int32(22)
	v3191 = v3157
	v3204 = v3182
	v3205 = v3183
	goto L523
L519:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = int64(16384)
	v3167 = int64(0)
	v3182 = v3167 - v3148
	v3183 = v3167 - (v3147 + base.I64_extend_i32_u(base.B2i32(v3148 != v3167)))
	goto L518
L520:
	;
	goto L521
L521:
	;
	v3175 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = v3175
	if v3148|v3147 == v3175 {
		v3292 = v3157
		v3294 = int32(0)
		goto L517
	} else {
		goto L522
	}
L522:
	;
	v3182 = v3148
	v3183 = v3147
	goto L518
L523:
	;
	v3215 = int32(16)
	v3216 = v32 + v3215
	v3219 = m.G0
	v3221 = v3219 - v3215
	m.G0 = v3221
	F___udivmodti4(m, v3221, v3204, v3205, int64(10000), int64(0))
	mBase = m.M
	v3225 = *(*int64)(unsafe.Add(mBase, uint32(v3221)+8))
	v3226 = *(*int64)(unsafe.Add(mBase, uint32(v3221)))
	*(*int64)(unsafe.Add(mBase, uint32(v3216))) = v3226
	*(*int64)(unsafe.Add(mBase, uint32(v3216)+8)) = v3225
	m.G0 = v3221 + v3215
	goto L525
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3273
	v3292 = v3284
	v3294 = v3191
	goto L517
L525:
	;
	v3232 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	v3233 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
	v3234 = int64(55536)
	v3235 = int64(0)
	v3240 = int64(32)
	v3243 = int64(base.Ui64(v3232) >> (uint(v3240) % 64))
	v3246 = int64(4294967295)
	v3249 = v3232 & v3246
	v3250 = v3234 * v3249
	v3254 = int64(base.Ui64(v3250)>>(uint(v3240)%64)) + v3234*v3243
	v3261 = v3249*v3235 + v3254&v3246
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v3232*v3235 + v3233*v3234 + v3235*v3243 + int64(base.Ui64(v3254)>>(uint(v3240)%64)) + int64(base.Ui64(v3261)>>(uint(v3240)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v3250&v3246 | v3261<<(uint(v3240)%64)
	goto L526
L526:
	;
	v3273 = v3190 - int32(2)
	v3274 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v3275 = v3274 + v3204
	*(*uint16)(unsafe.Add(mBase, uint32(v3273))) = uint16(v3275)
	v3279 = int64(0)
	v3284 = v3191 + int32(1)
	if v3205 == v3279 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v3285 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v3204))
	goto L529
L528:
	;
	v3285 = base.B2i32(v3205 != v3279)
	goto L529
L529:
	;
	if v3285 != 0 {
		v3190 = v3273
		v3191 = v3284
		v3204 = v3232
		v3205 = v3233
		goto L523
	} else {
		goto L530
	}
L530:
	;
	goto L524
L531:
	;
	v3323 = v552
	v3326 = v550
	v3328 = v548
	goto L67
L532:
	;
	v3356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3323<<(uint(int32(1))%32)))))
	v3365 = v3351
	v3367 = v3349 + v3356
	goto L66
L533:
	;
	v3529 = v3387 + v482
	v3544 = base.I64_extend_i32_s(v3367)
	v3549 = base.I64_extend_i32_s(v3498)
	v3552 = base.I64_extend_i32_s(v3365)
	goto L65
L534:
	;
	v3483 = v3461 * int32(_a_F_sqrt_var_4)
	if v34 <= v3457 {
		v3498 = v3483
		goto L533
	} else {
		goto L548
	}
L535:
	;
	v3457 = v3387
	v3461 = int32(0)
	goto L534
L536:
	;
	goto L537
L537:
	;
	v3395 = int32(0)
	v3401 = v3387
	v3405 = v3395
	v3409 = v3395
	goto L538
L538:
	;
	v3427 = v3405 * int32(_a_F_sqrt_var_4)
	if v3401 < v34 {
		goto L540
	} else {
		goto L541
	}
L539:
	;
	if v482&int32(1) == int32(0) {
		v3498 = v3445
		goto L533
	} else {
		goto L547
	}
L540:
	;
	v3432 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3401<<(uint(int32(1))%32)))))
	v3434 = v3427 + v3432
	goto L542
L541:
	;
	v3434 = v3427
	goto L542
L542:
	;
	v3436 = v3434 * int32(_a_F_sqrt_var_4)
	v3438 = v3401 + int32(1)
	if v3438 < v34 {
		goto L543
	} else {
		goto L544
	}
L543:
	;
	v3443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3438<<(uint(int32(1))%32)))))
	v3445 = v3436 + v3443
	goto L545
L544:
	;
	v3445 = v3436
	goto L545
L545:
	;
	v3446 = int32(2)
	v3447 = v3401 + v3446
	v3449 = v3409 + v3446
	if v3449 != v482&int32(1073741822) {
		v3401 = v3447
		v3405 = v3445
		v3409 = v3449
		goto L538
	} else {
		goto L546
	}
L546:
	;
	goto L539
L547:
	;
	v3457 = v3447
	v3461 = v3445
	goto L534
L548:
	;
	v3488 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3457<<(uint(int32(1))%32)))))
	v3498 = v3483 + v3488
	goto L533
L549:
	;
	v3575 = v3567
	goto L551
L550:
	;
	v3575 = v3565
	goto L551
L551:
	;
	if int32(0) < v453 {
		v449 = v3529
		v453 = v453 - int32(1)
		v461 = v3575
		v465 = v3563 + (v3567+v3565)&(v3563>>(uint(int64(63))%64))
		goto L63
	} else {
		goto L552
	}
L552:
	;
	goto L64
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+328)) = v3610
	v3613 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3610))) = uint16(v3613)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3610 + int32(2)
	if v3598 < int64(0) {
		goto L556
	} else {
		goto L557
	}
L554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v3684
	v3711 = v3681
	goto L59
L555:
	;
	v3638 = v3610 + int32(12)
	v3639 = v3613
	v3652 = v3631
	goto L560
L556:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = int64(16384)
	v3631 = int64(0) - v3598
	goto L555
L557:
	;
	goto L558
L558:
	;
	v3625 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = v3625
	if v3598 == v3625 {
		v3681 = v3613
		v3684 = int32(0)
		goto L554
	} else {
		goto L559
	}
L559:
	;
	v3631 = v3598
	goto L555
L560:
	;
	v3664 = v3638 - int32(2)
	v3666 = base.I64_div_u_s(v3652, int64(10000))
	v3669 = v3666*int64(55536) + v3652
	*(*uint16)(unsafe.Add(mBase, uint32(v3664))) = uint16(v3669)
	v3672 = v3639 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v3652) {
		v3638 = v3664
		v3639 = v3672
		v3652 = v3666
		goto L560
	} else {
		goto L562
	}
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3664
	v3681 = v3672
	v3684 = v3639
	goto L554
L562:
	;
	goto L561
L563:
	;
	v3773 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3771))) = uint16(v3773)
	if base.B2i32(v3768 == v3773)|base.B2i32(v3743 <= v3773) == v3773 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v3784 = *(*int32)(unsafe.Add(mBase, uint32(v32)+332))
	base.MemoryCopy(m, v3771+int32(2), v3784, v3768)
	goto L566
L565:
	;
	goto L566
L566:
	;
	v3786 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3786 != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	F_pfree(m, v3786)
	mBase = m.M
	v3788 = m.ExcPending
	if v3788 != 0 {
		goto L8
	} else {
		goto L570
	}
L568:
	;
	goto L569
L569:
	;
	v3789 = *(*int64)(unsafe.Add(mBase, uint32(v32)+320))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v3789
	v3791 = *(*int64)(unsafe.Add(mBase, uint32(v32)+312))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v3791
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v3771
	v3794 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v3771 + v3794
	v3797 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3797
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	v3809 = l2 + v90<<(uint(v3794)%32)
	if v3809+int32(4) < v3797 {
		goto L572
	} else {
		goto L573
	}
L570:
	;
	goto L569
L571:
	;
	v3924 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3925 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v3925 {
		goto L599
	} else {
		goto L600
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L571
L573:
	;
	goto L574
L574:
	;
	v3818 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3820 = l2 & int32(3)
	v3824 = base.I32_div_s(v3809+int32(7), int32(4))
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v3825 <= v3824 {
		goto L579
	} else {
		goto L580
	}
L575:
	;
	goto L571
L576:
	;
	if int32(0) <= v3890 {
		goto L575
	} else {
		goto L596
	}
L577:
	;
	v3870 = v3864
	goto L590
L578:
	;
	v3839 = int32(1)
	v3840 = v3824 - v3839
	v3843 = v3818 + v3840<<(uint(v3839)%32)
	v3844 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3843))))
	v3845 = int32(2)
	v3847 = *(*int32)(unsafe.Add(mBase, uint32(v3820<<(uint(v3845)%32))+uint32(_c_F_sqrt_var[0])))
	v3848 = base.I32_rem_s(v3844, v3847)
	v3849 = v3844 - v3848
	*(*uint16)(unsafe.Add(mBase, uint32(v3843))) = uint16(v3849)
	v3852 = base.I32_div_s(v3847, v3845)
	if v3848 < v3852 {
		v3890 = v3840
		goto L576
	} else {
		goto L585
	}
L579:
	;
	if base.B2i32(v3820 == int32(0))|base.B2i32(v3824 != v3825) != 0 {
		goto L575
	} else {
		goto L582
	}
L580:
	;
	goto L581
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3824
	if v3820 != 0 {
		goto L578
	} else {
		goto L583
	}
L582:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3824
	goto L578
L583:
	;
	v3836 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3818+v3824<<(uint(int32(1))%32)))))
	if v3836 <= int32(_a_F_sqrt_var_7) {
		v3890 = v3824
		goto L576
	} else {
		goto L584
	}
L584:
	;
	v3864 = v3824
	goto L577
L585:
	;
	v3855 = v3847 + base.I32_extend16_s(v3849)
	if int32(_a_F_sqrt_var_8) < v3855 {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v3860 = v3855 + int32(_a_F_sqrt_var_9)
	goto L588
L587:
	;
	v3860 = v3855
	goto L588
L588:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3843))) = uint16(v3860)
	if v3855 < int32(_a_F_sqrt_var_4) {
		v3890 = v3840
		goto L576
	} else {
		goto L589
	}
L589:
	;
	v3864 = v3840
	goto L577
L590:
	;
	v3876 = int32(1)
	v3877 = v3870 - v3876
	v3880 = v3818 + v3877<<(uint(v3876)%32)
	v3883 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3880))))
	v3885 = base.B2i32(int32(_a_F_sqrt_var_10) < v3883)
	if int32(_a_F_sqrt_var_10) < v3883 {
		goto L592
	} else {
		goto L593
	}
L591:
	;
	v3890 = v3877
	goto L576
L592:
	;
	v3886 = int32(-9999)
	goto L594
L593:
	;
	v3886 = v3876
	goto L594
L594:
	;
	v3887 = v3886 + v3883
	*(*uint16)(unsafe.Add(mBase, uint32(v3880))) = uint16(v3887)
	if int32(_a_F_sqrt_var_10) < v3883 {
		v3870 = v3877
		goto L590
	} else {
		goto L595
	}
L595:
	;
	goto L591
L596:
	;
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v3898 - int32(2)
	v3902 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3903 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3902 + v3903
	v3906 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3906 + v3903
	goto L575
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4047
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4046
	v4073 = *(*int32)(unsafe.Add(mBase, uint32(v32)+328))
	if v4073 != 0 {
		goto L613
	} else {
		goto L614
	}
L598:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v4046 = v4014
	v4047 = int32(0)
	goto L597
L599:
	;
	v3935 = v3924
	v3936 = v3925
	goto L602
L600:
	;
	goto L601
L601:
	;
	if v3925 != 0 {
		v4046 = v3924
		v4047 = v3925
		goto L597
	} else {
		goto L612
	}
L602:
	;
	v3960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3935))))
	if v3960 != 0 {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	v4014 = v3924 + v3925<<(uint(int32(1))%32)
	goto L598
L604:
	;
	v3966 = v3936
	goto L607
L605:
	;
	goto L606
L606:
	;
	v4000 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v4001 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v4000 - v4001
	if v4001 < v3936 {
		v3935 = v3935 + int32(2)
		v3936 = v3936 - v4001
		goto L602
	} else {
		goto L611
	}
L607:
	;
	v3995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3935+v3966<<(uint(int32(1))%32)-int32(2)))))
	if v3995 != 0 {
		v4046 = v3935
		v4047 = v3966
		goto L597
	} else {
		goto L609
	}
L609:
	;
	v3996 = int32(1)
	if v3996 < v3966 {
		v3966 = v3966 - v3996
		goto L607
	} else {
		goto L610
	}
L610:
	;
	v4014 = v3935
	goto L598
L611:
	;
	goto L603
L612:
	;
	v4014 = v3924
	goto L598
L613:
	;
	F_pfree(m, v4073)
	mBase = m.M
	v4075 = m.ExcPending
	if v4075 != 0 {
		goto L8
	} else {
		goto L616
	}
L614:
	;
	goto L615
L615:
	;
	v4076 = *(*int32)(unsafe.Add(mBase, uint32(v32)+304))
	if v4076 != 0 {
		goto L617
	} else {
		goto L618
	}
L616:
	;
	goto L615
L617:
	;
	F_pfree(m, v4076)
	mBase = m.M
	v4078 = m.ExcPending
	if v4078 != 0 {
		goto L8
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	if v3749 != 0 {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	goto L619
L621:
	;
	F_pfree(m, v3749)
	mBase = m.M
	v4080 = m.ExcPending
	if v4080 != 0 {
		goto L8
	} else {
		goto L624
	}
L622:
	;
	goto L623
L623:
	;
	if v3750 != 0 {
		goto L625
	} else {
		goto L626
	}
L624:
	;
	goto L623
L625:
	;
	F_pfree(m, v3750)
	mBase = m.M
	v4082 = m.ExcPending
	if v4082 != 0 {
		goto L8
	} else {
		goto L628
	}
L626:
	;
	goto L627
L627:
	;
	v4083 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	if v4083 != 0 {
		goto L629
	} else {
		goto L630
	}
L628:
	;
	goto L627
L629:
	;
	F_pfree(m, v4083)
	mBase = m.M
	v4085 = m.ExcPending
	if v4085 != 0 {
		goto L8
	} else {
		goto L632
	}
L630:
	;
	goto L631
L631:
	;
	v4086 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	if v4086 == int32(0) {
		goto L1
	} else {
		goto L633
	}
L632:
	;
	goto L631
L633:
	;
	F_pfree(m, v4086)
	mBase = m.M
	v4090 = m.ExcPending
	if v4090 != 0 {
		goto L8
	} else {
		goto L634
	}
L634:
	;
	goto L1
}
func F_stat(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	v5 = F___fstatat(m, int32(-100), l0, l1, int32(0))
	return v5
}
func F_statement_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_statement_timestamp[0]))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_str_initcap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v100 = v4
		m.G0 = v10 + int32(16)
		return v100
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_str_initcap_0)
					F_errmsg(m, int32(_a_F_str_initcap_1), v10)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_str_initcap_2), int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_str_initcap_3), int32(1783), int32(_a_F_str_initcap_4))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
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
			v17 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
				if v21 == int32(1) {
					v24 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
						if v26 == int32(0) {
							v100 = v24
						} else {
							v29 = v24
							v31 = v26
							v32 = int32(1)
							for {
								if v32&int32(1) == int32(0) {
									v40 = int32(255)
									v41 = v31 & v40
									if base.Ui32((v41-int32(65))&v40) < base.Ui32(int32(26)) {
										v50 = v41 | int32(32)
									} else {
										v50 = v41
									}
									v64 = v50
								} else {
									v51 = int32(255)
									v52 = v31 & v51
									if base.Ui32((v52-int32(97))&v51) < base.Ui32(int32(26)) {
										v61 = v52 - int32(32)
									} else {
										v61 = v52
									}
									v64 = v61 & int32(255)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v64)
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
								if v80 != 0 {
									v29 = v29 + int32(1)
									v31 = v80
									v32 = base.B2i32(base.Ui32((v64&int32(223)-int32(91))&int32(255)) < base.Ui32(int32(230))) & base.B2i32(base.Ui32(base.I32_extend8_s(v64)-int32(58)) < base.Ui32(int32(-10)))
									continue
								} else {
									break
								}
								break
							}
							v100 = v24
						}
						m.G0 = v10 + int32(16)
						return v100
					}
				} else {
					v84 = l1 + int32(1)
					v85 = F_palloc(m, v84)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v87 = F_pg_strtitle(m, v85, v84, l0, l1, v17)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							v90 = v87 + int32(1)
							if base.Ui32(v90) <= base.Ui32(v84) {
								v100 = v85
								m.G0 = v10 + int32(16)
								return v100
							} else {
								v92 = F_repalloc(m, v85, v90)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									v94 = F_pg_strtitle(m, v92, v90, l0, l1, v17)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										v100 = v92
										m.G0 = v10 + int32(16)
										return v100
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
func F_strict_word_similarity_commutator_op(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14022(m, l0, int32(_a_F_strict_word_similarity_commutator_op_0), int32(3))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_strip_phvs_in_index_operand_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v3 = int32(0)
	if l0 == v3 {
		v21 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v21
L2:
	;
	v6 = l0
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v9 != int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v15 = F_expression_tree_mutator_impl(m, v6, int32(826), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L4
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v12 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v13 != 0 {
		v6 = v13
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v21 = v3
	goto L1
L9:
	;
	return int32(0)
L10:
	;
	v21 = v15
	goto L1
}
func F_strlist_to_textarray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v2
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0]))
	v23 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_strlist_to_textarray_0), v2, int32(_a_F_strlist_to_textarray_1), int32(_a_F_strlist_to_textarray_2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = int32(_a_F_strlist_to_textarray_3)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0])) = v23
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0])) = v28
	v98 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v98
	v109 = F_construct_md_array(m, v89, v90, v98, v13+int32(12), v13+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v34 = F_palloc(m, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = F_palloc(m, v39<<(uint(int32(2))%32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v37 = F_palloc(m, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v89 = v34
	v90 = v37
	goto L3
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v45 = F_palloc(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47 <= int32(0) {
		v89 = v42
		v90 = v45
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v55 = v2
	v58 = v2
	goto L12
L12:
	;
	v60 = v45 + v55
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v58<<(uint(int32(2))%32))))
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v89 = v42
	v90 = v45
	goto L3
L14:
	;
	v83 = v58 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v83 < v84 {
		v55 = v80
		v58 = v83
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v66)
	v68 = F_cstring_to_text(m, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v77)
	v80 = v55
	goto L14
L18:
	;
	v71 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v42+v55<<(uint(int32(2))%32)))) = v68
	v80 = v71
	goto L14
L19:
	;
	goto L13
L20:
	;
	F_MemoryContextDelete(m, v23)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	m.G0 = v13 + int32(16)
	return v109
}
func F_strncoll_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(1024)
	m.G0 = v11
	v15 = l1 + l3 + int32(2)
	if base.Ui32(int32(1025)) <= base.Ui32(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = F_palloc(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v22 = v11
	goto L3
L3:
	;
	if l1 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v22 = v18
	goto L3
L6:
	;
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v29 = l0
	goto L8
L8:
	;
	if l3 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	base.MemoryCopy(m, v22, l0, l1)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v22))) = uint8(v27)
	v29 = v22
	goto L8
L12:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if base.B2i32(v43 == int32(0))|base.B2i32(v43 != v46) != 0 {
		v64 = v43
		v65 = v46
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v39 = l2
	goto L12
L14:
	;
	goto L15
L15:
	;
	v34 = l1 + v22 + int32(1)
	if l3 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, v34, l2, l3)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34+l3))) = uint8(v37)
	v39 = v34
	goto L12
L19:
	;
	if v22 != v11 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v49 = v29
	v50 = v39
	goto L22
L22:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v64 = v54
		v65 = v53
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v64 = v54
	v65 = v53
	goto L20
L24:
	;
	v57 = int32(1)
	if v54 == v53 {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_pfree(m, v22)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	m.G0 = v11 + int32(1024)
	return v64 - v65
L29:
	;
	goto L28
}
func F_strrchr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = F_strlen(m, l0)
	mBase = m.M
	v12 = v5 + int32(1)
	goto L2
L1:
	;
	return v24
L2:
	;
	v14 = int32(0)
	if v12 == v14 {
		v24 = v14
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v24 = v19
	goto L1
L4:
	;
	v18 = v12 - int32(1)
	v19 = l0 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v20 != l1&int32(255) {
		v12 = v18
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
}
func F_strspn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v6 = m.G0
	v8 = v6 - int32(32)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v9
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = l0
	goto L7
L5:
	;
	goto L6
L6:
	;
	v37 = l1
	v38 = v17
	goto L10
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v32 == v17 {
		v26 = v26 + int32(1)
		goto L7
	} else {
		goto L9
	}
L8:
	;
	return v26 - l0
L9:
	;
	goto L8
L10:
	;
	v45 = v8 + int32(base.Ui32(v38)>>(uint(int32(3))%32))&int32(28)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46 | v47<<(uint(v38)%32)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v51 != 0 {
		v37 = v37 + v47
		v38 = v51
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v54 == int32(0) {
		v77 = l0
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return v77 - l0
L14:
	;
	v58 = l0
	v59 = v54
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v59)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v67)>>(uint(v59)%32))&int32(1) == int32(0) {
		v77 = v58
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v77 = v75
	goto L13
L17:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v75 = v58 + int32(1)
	if v73 != 0 {
		v58 = v75
		v59 = v73
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_subxact_info_read(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(1040)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v13 = v8 + int32(16)
	v16 = F_pg_snprintf(m, v13, int32(1024), int32(_a_F_subxact_info_read_0), v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[0]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
		v23 = F_BufFileOpenFileSet(m, v20, v13, int32(0), int32(1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 != 0 {
				F_BufFileReadExact(m, v23, int32(_a_F_subxact_info_read_1), int32(4))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[1]))
					v34 = int32(1073741823)
					if v34 <= v32 {
						v37 = v34
					} else {
						v37 = v32
					}
					if base.Ui32(int32(2)) <= base.Ui32(v37) {
						v45 = int32(32) - base.I32_clz(v37-int32(1))
					} else {
						v45 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[2])) = int32(1) << (uint(v45) % 32)
					v48 = int32(_a_F_subxact_info_read_2)
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[3]))
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[4]))
					*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[3])) = v52
					v56 = F_palloc(m, int32(16)<<(uint(v45)%32))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[3])) = v49
						*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[5])) = v56
						v63 = v32 << (uint(int32(4)) % 32)
						if v63 != 0 {
							F_BufFileReadExact(m, v23, v56, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_BufFileClose(m, v23)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									m.G0 = v8 + int32(1040)
									return
								}
							}
						} else {
							F_BufFileClose(m, v23)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								m.G0 = v8 + int32(1040)
								return
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(1040)
				return
			}
		}
	}
}
func F_superuser(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_superuser[0]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_superuser[1]))
	if base.B2i32(v5 == v1)|base.B2i32(v5 != v9) == v1 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])))
		v57 = v15
		return v57 & int32(1)
	} else {
		if v9 == int32(10) {
			v18 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[3])))
			if v20&v18 == int32(0) {
				v57 = v18
				return v57 & int32(1)
			} else {
				v28 = F_SearchSysCache1(m, int32(11), v9)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if v28 != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v33)+68)))
						F_ReleaseCatCache(m, v28)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = v35
							v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
							if v40 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v49 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v49)
									*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
									v55 = v38 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
									v57 = v38
									return v57 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
								v55 = v38 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
								v57 = v38
								return v57 & int32(1)
							}
						}
					} else {
						v38 = int32(0)
						v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
						if v40 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v49)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
								v55 = v38 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
								v57 = v38
								return v57 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
							v55 = v38 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
							v57 = v38
							return v57 & int32(1)
						}
					}
				}
			}
		} else {
			v28 = F_SearchSysCache1(m, int32(11), v9)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v28 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v33)+68)))
					F_ReleaseCatCache(m, v28)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = v35
						v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
						if v40 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v49)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
								v55 = v38 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
								v57 = v38
								return v57 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
							v55 = v38 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
							v57 = v38
							return v57 & int32(1)
						}
					}
				} else {
					v38 = int32(0)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
					if v40 == int32(0) {
						F_CacheRegisterSyscacheCallback(m, int32(11), int32(1768), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v49)
							*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
							v55 = v38 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
							v57 = v38
							return v57 & int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
						v55 = v38 & int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v55)
						v57 = v38
						return v57 & int32(1)
					}
				}
			}
		}
	}
}
func F_swedish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v318 < v321 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	if v62 < int32(0) {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v17 = v10
	v19 = int32(3)
	goto L9
L8:
	;
	v62 = v47
	goto L2
L9:
	;
	if v7 <= v17 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v62 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v24 = v17 + int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v17))))
	if base.Ui32(v26) < base.Ui32(int32(192)) {
		v47 = v24
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(1)
	if v48 < v19 {
		v17 = v47
		v19 = v19 - v48
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v7 <= v24 {
		v47 = v24
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v33 = v24
	goto L17
L17:
	;
	v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9+v33))))
	if int32(-65) < v36 {
		v47 = v33
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v47 = v7
	goto L14
L19:
	;
	v40 = v33 + int32(1)
	if v40 != v7 {
		v33 = v40
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L10
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = v10
	goto L25
L23:
	;
	if v184 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v184 = v156
	goto L23
L25:
	;
	if v80 <= v89 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v184 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v96 = int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v81))))
	if base.Ui32(v98) < base.Ui32(int32(192)) {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(246) < v155 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v102 = v89 + int32(1)
	if v102 == v80 {
		v155 = v98
		v156 = v96
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v81))))
	v107 = v105 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v98) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111+v81))))
	v123 = v121 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v98) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v111 = v89 + int32(2)
	if v111 != v80 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v155 = v98<<(uint(int32(6))%32)&int32(1984) | v107
	v156 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v127))))
	v155 = v140&int32(63) | (v98<<(uint(int32(18))%32)&int32(_a_F_swedish_UTF_8_stem_0) | v107<<(uint(int32(12))%32) | v123<<(uint(int32(6))%32))
	v156 = int32(4)
	goto L30
L39:
	;
	v127 = v89 + int32(3)
	if v127 != v80 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v155 = v98<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v107<<(uint(int32(6))%32) | v123
	v156 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v173 = v156 + v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v173
	v89 = v173
	goto L25
L44:
	;
	v160 = v155 - int32(97)
	if v160 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v160)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[0]))))
	if int32(base.Ui32(v166)>>(uint(v160&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = v198
	goto L51
L49:
	;
	if v304 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v304 = v275
	goto L49
L51:
	;
	if v199 <= v208 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v304 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v215 = int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208+v200))))
	if base.Ui32(v217) < base.Ui32(int32(192)) {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(246) < v274 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v221 = v208 + int32(1)
	if v221 == v199 {
		v274 = v217
		v275 = v215
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v200))))
	v226 = v224 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v217) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v200))))
	v242 = v240 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v217) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v230 = v208 + int32(2)
	if v230 != v199 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v274 = v217<<(uint(int32(6))%32)&int32(1984) | v226
	v275 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v246))))
	v274 = v259&int32(63) | (v217<<(uint(int32(18))%32)&int32(_a_F_swedish_UTF_8_stem_0) | v226<<(uint(int32(12))%32) | v242<<(uint(int32(6))%32))
	v275 = int32(4)
	goto L56
L65:
	;
	v246 = v208 + int32(3)
	if v246 != v199 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v274 = v217<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v226<<(uint(int32(6))%32) | v242
	v275 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v279 = v274 - int32(97)
	if v279 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v279)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[0]))))
	if int32(base.Ui32(v285)>>(uint(v279&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v293 = v275 + v208
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	v208 = v293
	goto L51
L73:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v308 = v307 + v304
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	if v311 < v308 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v313 = v308
	goto L76
L75:
	;
	v313 = v311
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+4)) = v313
	goto L1
L77:
	;
	return v782
L78:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v495
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v498 <= v495 {
		goto L116
	} else {
		goto L117
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v318
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v321
	if v318 <= v321 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	goto L78
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v328 = int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v318-v328))))
	if base.B2i32(v330&int32(224) != int32(96))|base.B2i32(v328<<(uint(v330)%32)&int32(_a_F_swedish_UTF_8_stem_2) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v344 = F_find_among_b(m, l0, int32(_a_F_swedish_UTF_8_stem_3), int32(37))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v344 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v351
	switch v344 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	default:
		goto L78
	}
L86:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L92
L87:
	;
	v355 = F_slice_del(m, l0)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	if int32(0) <= v355 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v782 = v355
	goto L77
L90:
	;
	if v487 != 0 {
		goto L78
	} else {
		goto L113
	}
L91:
	;
	v487 = v480
	goto L90
L92:
	;
	if v371 <= v372 {
		v480 = int32(-1)
		goto L91
	} else {
		goto L94
	}
L93:
	;
	v480 = int32(0)
	goto L91
L94:
	;
	v389 = int32(1)
	v390 = v371 - v389
	v392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v373+v390))))
	v394 = v392 & int32(255)
	if base.B2i32(v390 == v372)|base.B2i32(int32(0) <= v392) != 0 {
		v452 = v394
		v456 = v389
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if int32(121) < v452 {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v401 = v394 & int32(63)
	v403 = v371 - int32(2)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v403))))
	v407 = v405 << (uint(int32(6)) % 32)
	if base.B2i32(v403 != v372)&base.B2i32(base.Ui32(v405) < base.Ui32(int32(192))) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v452 = v407&int32(1984) | v401
	v456 = int32(2)
	goto L95
L98:
	;
	goto L99
L99:
	;
	v420 = v407&int32(4032) | v401
	v422 = v371 - int32(3)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v422))))
	if base.B2i32(v422 != v372)&base.B2i32(base.Ui32(v424) < base.Ui32(int32(224))) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v452 = v424<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v420
	v456 = int32(3)
	goto L95
L101:
	;
	goto L102
L102:
	;
	v442 = int32(4)
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v373-v442))))
	v452 = v424<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_4) | v444&int32(7)<<(uint(int32(18))%32) | v420
	v456 = v442
	goto L95
L103:
	;
	v487 = v456
	goto L90
L104:
	;
	goto L105
L105:
	;
	v458 = v452 - int32(98)
	if v458 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v487 = v456
	goto L90
L107:
	;
	goto L108
L108:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v458)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[1]))))
	if int32(base.Ui32(v464)>>(uint(v458&int32(7))%32))&int32(1) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v487 = v456
	goto L90
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v371 - v456
	goto L112
L112:
	;
	goto L93
L113:
	;
	v488 = F_slice_del(m, l0)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L83
	} else {
		goto L114
	}
L114:
	;
	if int32(0) <= v488 {
		goto L78
	} else {
		goto L115
	}
L115:
	;
	v782 = v488
	goto L77
L116:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v498
	v503 = v495 - int32(1)
	if v503 <= v498 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v594 = v495
	v595 = v497
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v594
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v594 < v597 {
		goto L146
	} else {
		goto L147
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v500
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v594 = v593
	v595 = v592
	goto L118
L120:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v503))))
	if base.B2i32(v507&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v507)%32)&int32(_a_F_swedish_UTF_8_stem_5) == int32(0)) != 0 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v521 = F_find_among_b(m, l0, int32(_a_F_swedish_UTF_8_stem_6), int32(7))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L83
	} else {
		goto L122
	}
L122:
	;
	if v521 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v525
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L126
L124:
	;
	if v581 < int32(0) {
		goto L119
	} else {
		goto L143
	}
L126:
	;
	goto L127
L127:
	;
	goto L128
L128:
	;
	v536 = v525
	v538 = int32(1)
	goto L131
L130:
	;
	v581 = v563
	goto L124
L131:
	;
	if v536 <= v529 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L130
L133:
	;
	v581 = int32(-1)
	goto L124
L134:
	;
	goto L135
L135:
	;
	v543 = v536 - int32(1)
	v545 = int32(*(*int8)(unsafe.Add(mBase, uint32(v528+v543))))
	if base.B2i32(int32(0) <= v545)|base.B2i32(v543 <= v529) != 0 {
		v563 = v543
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v567 = int32(1)
	if v567 < v538 {
		v536 = v563
		v538 = v538 - v567
		goto L131
	} else {
		goto L142
	}
L137:
	;
	v551 = v543
	goto L138
L138:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528+v551))))
	if base.Ui32(int32(191)) < base.Ui32(v556) {
		v563 = v551
		goto L136
	} else {
		goto L140
	}
L139:
	;
	v563 = v529
	goto L136
L140:
	;
	v560 = v551 - int32(1)
	if v529 < v560 {
		v551 = v560
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	goto L132
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v581
	v586 = F_slice_del(m, l0)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L83
	} else {
		goto L144
	}
L144:
	;
	if v586 < int32(0) {
		v782 = v586
		goto L77
	} else {
		goto L145
	}
L145:
	;
	goto L119
L146:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v779
	v782 = int32(1)
	goto L77
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v594
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v597
	v603 = v594 - int32(1)
	if v603 <= v597 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v600
	goto L146
L149:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605+v603))))
	if base.B2i32(v607&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v607)%32)&int32(_a_F_swedish_UTF_8_stem_7) == int32(0)) != 0 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v621 = F_find_among_b(m, l0, int32(_a_F_swedish_UTF_8_stem_8), int32(5))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L83
	} else {
		goto L151
	}
L151:
	;
	if v621 == int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v600
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v626
	switch v621 - int32(1) {
	case 0:
		goto L155
	case 1:
		goto L154
	case 2:
		goto L153
	default:
		goto L146
	}
L153:
	;
	v771 = F_slice_from_s(m, l0, int32(4), int32(_a_F_swedish_UTF_8_stem_9))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L83
	} else {
		goto L184
	}
L154:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L160
L155:
	;
	v630 = F_slice_del(m, l0)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L83
	} else {
		goto L156
	}
L156:
	;
	if int32(0) <= v630 {
		goto L146
	} else {
		goto L157
	}
L157:
	;
	v782 = v630
	goto L77
L158:
	;
	if v762 != 0 {
		goto L146
	} else {
		goto L181
	}
L159:
	;
	v762 = v755
	goto L158
L160:
	;
	if v646 <= v647 {
		v755 = int32(-1)
		goto L159
	} else {
		goto L162
	}
L161:
	;
	v755 = int32(0)
	goto L159
L162:
	;
	v664 = int32(1)
	v665 = v646 - v664
	v667 = int32(*(*int8)(unsafe.Add(mBase, uint32(v648+v665))))
	v669 = v667 & int32(255)
	if base.B2i32(v665 == v647)|base.B2i32(int32(0) <= v667) != 0 {
		v727 = v669
		v731 = v664
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if int32(118) < v727 {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	v676 = v669 & int32(63)
	v678 = v646 - int32(2)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648+v678))))
	v682 = v680 << (uint(int32(6)) % 32)
	if base.B2i32(v678 != v647)&base.B2i32(base.Ui32(v680) < base.Ui32(int32(192))) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v727 = v682&int32(1984) | v676
	v731 = int32(2)
	goto L163
L166:
	;
	goto L167
L167:
	;
	v695 = v682&int32(4032) | v676
	v697 = v646 - int32(3)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648+v697))))
	if base.B2i32(v697 != v647)&base.B2i32(base.Ui32(v699) < base.Ui32(int32(224))) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v727 = v699<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v695
	v731 = int32(3)
	goto L163
L169:
	;
	goto L170
L170:
	;
	v717 = int32(4)
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+v648-v717))))
	v727 = v699<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_4) | v719&int32(7)<<(uint(int32(18))%32) | v695
	v731 = v717
	goto L163
L171:
	;
	v762 = v731
	goto L158
L172:
	;
	goto L173
L173:
	;
	v733 = v727 - int32(105)
	if v733 < int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v762 = v731
	goto L158
L175:
	;
	goto L176
L176:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v733)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[2]))))
	if int32(base.Ui32(v739)>>(uint(v733&int32(7))%32))&int32(1) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v762 = v731
	goto L158
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v646 - v731
	goto L180
L180:
	;
	goto L161
L181:
	;
	v765 = F_slice_from_s(m, l0, int32(3), int32(_a_F_swedish_UTF_8_stem_10))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L83
	} else {
		goto L182
	}
L182:
	;
	if int32(0) <= v765 {
		goto L146
	} else {
		goto L183
	}
L183:
	;
	v782 = v765
	goto L77
L184:
	;
	if int32(0) <= v771 {
		goto L146
	} else {
		goto L185
	}
L185:
	;
	v782 = v771
	goto L77
}
