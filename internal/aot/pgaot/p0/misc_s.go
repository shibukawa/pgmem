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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v201 int32
	_ = v201
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
	var v302 int32
	_ = v302
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
	var v488 int32
	_ = v488
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
	return v488
L7:
	;
	if v445 != 0 {
		v488 = v6
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
		v488 = v49
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v55 = v33 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v56 != v57 {
		v488 = v6
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v424 = v55
	v425 = v34 - int32(1)
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
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
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
	v488 = v49
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
	v488 = int32(-1)
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
		v488 = v123
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
	v302 = v6
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
	v201 = v6
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
	v216 = v201 + v213
	if v216 != v190 {
		v198 = v212
		v200 = v214
		v201 = v216
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
	v302 = v186
	goto L61
L105:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	if v302 == int32(0) {
		v488 = int32(1)
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_pfree(m, v302)
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
	if v302 == v341 {
		v488 = v341
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_pfree(m, v302)
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
	v488 = v6
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
	F_errmsg(m, int32(217136), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(497799), int32(107), int32(63932))
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
	F_errmsg(m, int32(217136), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(497799), int32(169), int32(63932))
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
	F_errmsg(m, int32(217136), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(497799), int32(237), int32(63932))
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
		v488 = v453
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
	v488 = v453
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
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	v3 = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	v11 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v11)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = l0 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v21 = l0 + int32(32)
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l5)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v11
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
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v14)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v13)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v11
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v10
	if v13 != 0 {
		v21 = v9
	} else {
		v21 = int32(0)
	}
	if v14 != 0 {
		v22 = v21
	} else {
		v22 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v24 != 0 {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v29 = v24 << (uint(int32(2)) % 32)
		if v29 != 0 {
			v30 = F__emscripten_memcpy_bulkmem(m, l1+int32(24), v27, v29)
			mBase = m.M
		} else {
		}
	} else {
	}
	if int32(0) < v22 {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v35 = int32(2)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v43 = v41 << (uint(v35) % 32)
		if v43 != 0 {
			v44 = F__emscripten_memcpy_bulkmem(m, l1+v34<<(uint(v35)%32)+int32(24), v40, v43)
			mBase = m.M
		} else {
		}
	} else {
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
					F_errmsg_internal(m, int32(46249), v8)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(492340), int32(96), int32(353948))
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
						F_sequence_close(m, v12, int32(3))
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
func F_ShmemInitHash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = int32(1)
	v18 = int32(32)
	v22 = int32(1073741823)
	if v22 <= l2 {
		v25 = v22
	} else {
		v25 = l2
	}
	v26 = int32(1)
	if base.Ui32(v25) <= base.Ui32(v26) {
		v39 = v16
	} else {
		v39 = int32(base.Ui32(int32(-1)<<(uint(v18-base.I32_clz(v25-v26))%32)^int32(-1))>>(uint(int32(8))%32)) + v26
	}
	v40 = int32(1)
	if base.Ui32(v39) <= base.Ui32(v40) {
		v47 = v16
	} else {
		v47 = v16 << (uint(v18-base.I32_clz(v39-v40)) % 32)
	}
	v49 = int32(256)
	for {
		if v49 < v47 {
			v49 = v49 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = int32(1105)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v49
	v68 = F_ShmemInitStruct(m, l0, v49<<(uint(int32(2))%32)+int32(432), v12+int32(15))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		return int32(0)
	} else {
		v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v68
		if v72 != 0 {
			v76 = l4 | int32(6660)
		} else {
			v76 = l4 | int32(2564)
		}
		v77 = F_hash_create(m, l0, l1, l3, v76)
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return int32(0)
		} else {
			m.G0 = v12 + int32(16)
			return v77
		}
	}
}
func F_SignalHandlerForShutdownRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[441])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
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
	var v19 int32
	_ = v19
	v5 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v7 = v5 + int32(20)
	*(*int32)(unsafe.Add(mBase, _consts[80])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	if v10 != 0 {
		v11 = int32(1)
		F_ModifyWaitEvent(m, v10, v11, v11, v7)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[80]))
			v17 = v16
			F_SetLatch(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v17 = v7
		F_SetLatch(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
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
	v13 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[122]))
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
					*(*int32)(unsafe.Add(mBase, _consts[122])) = v24
					*(*int32)(unsafe.Add(mBase, _consts[4])) = l0
					v43 = int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[122])) = v24 | int32(2)
					*(*int32)(unsafe.Add(mBase, _consts[4])) = l0
					v37 = int32(4513288)
					v39 = *(*int32)(unsafe.Add(mBase, _consts[159]))
					v41 = v39 + int32(1)
					*(*int32)(unsafe.Add(mBase, _consts[159])) = v41
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
							F_errmsg(m, int32(702005), v8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(492386), int32(45), int32(217701))
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0 & v7
	if base.Ui32((l0-int32(33))&v7) < base.Ui32(int32(94)) {
		v20 = int32(689521)
	} else {
		v20 = int32(29706)
	}
	v21 = F_pg_snprintf(m, int32(4413948), int32(5), v20, v5)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
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
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = F_scanRTEForColumn(m, l0, v14, v15, l3, l4, v6, v6)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L76
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L71
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L61
	}
L5:
	;
	return int32(0)
L6:
	;
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v24 = base.B2i32(v18 == int32(-6))
	if v18 == int32(-6) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v247 = v6
	goto L9
L9:
	;
	m.G0 = v12 - int32(-64)
	return v247
L10:
	;
	if v18 == int32(-6) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if int32(0) <= v18 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if v22 == int32(28) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v18 == int32(-6) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if int32(0) <= v18 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v22 == int32(43) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	if int32(0) < v18 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if int32(0) <= v18 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v22 == int32(18) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+44)) = l4
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	if v87 == int32(0) {
		v147 = l0
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v44 = v41 + v18<<(uint(int32(5))%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(32))))
	if v47 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v72 = base.I32_extend16_s(v18)
	v73 = F_SystemAttributeDefinition(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44-int32(28)))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(24))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(20))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(16))))
	v62 = F_makeVar(m, v47, v52, v55, v58, v61, l2)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+36)) = v66
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+40)) = uint16(v70)
	v82 = v62
	goto L22
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+68))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+76))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+96))
	v79 = F_makeVar(m, v75, v72, v76, v77, v78, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v82 = v79
	goto L22
L30:
	;
	if v86 <= int32(0) {
		v172 = v87
		goto L42
	} else {
		goto L43
	}
L31:
	;
	v91 = v87 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v87) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = int32(0)
	v102 = l0
	goto L35
L33:
	;
	v122 = l0
	goto L34
L34:
	;
	if v91 == int32(0) {
		v147 = v122
		goto L30
	} else {
		goto L38
	}
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v115 = v98 + int32(8)
	if v115 != v87&int32(-8) {
		v98 = v115
		v102 = v113
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v122 = v113
	goto L34
L37:
	;
	goto L36
L38:
	;
	v130 = int32(0)
	v134 = v122
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v140 = v130 + int32(1)
	if v140 != v91 {
		v130 = v140
		v134 = v138
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v147 = v138
	goto L30
L41:
	;
	goto L40
L42:
	;
	if v172 == int32(0) {
		v228 = l0
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	if v153 == int32(0) {
		v172 = v87
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v156 < v86 {
		v172 = v87
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158+v86<<(uint(int32(2))%32)-int32(4))))
	if v164 == int32(0) {
		v172 = v87
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v168 = F_bms_union(m, v167, v164)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = v168
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v172 = v171
	goto L42
L48:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+8)))
	F_markRTEForSelectPriv(m, v228, v237, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L60
	}
L49:
	;
	v177 = v172 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v172) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v183 = l0
	v188 = int32(0)
	goto L53
L51:
	;
	v203 = l0
	goto L52
L52:
	;
	if v177 == int32(0) {
		v228 = v203
		goto L48
	} else {
		goto L56
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v201 = v188 + int32(8)
	if v201 != v172&int32(-8) {
		v183 = v199
		v188 = v201
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v203 = v199
	goto L52
L55:
	;
	goto L54
L56:
	;
	v215 = v203
	v220 = int32(0)
	goto L57
L57:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v226 = v220 + int32(1)
	if v226 != v177 {
		v215 = v224
		v220 = v226
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v228 = v224
	goto L48
L59:
	;
	goto L58
L60:
	;
	v247 = v82
	goto L9
L61:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
	F_errmsg(m, int32(435157), v12)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(496161), int32(713), int32(274815))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l3
	F_errmsg(m, int32(269490), v10+int32(-48))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(496161), int32(726), int32(274815))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l3
	F_errmsg(m, int32(250622), v10+int32(-32))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	F_parser_errposition(m, l0, l4)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(496161), int32(737), int32(274815))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
	F_errmsg(m, int32(71586), v10+int32(-16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(496161), int32(751), int32(274815))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
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
	if base.Ui32(int32(8192)) <= base.Ui32(v4) {
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
			if int32(8192) <= v12 {
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
		if int32(8192) <= v12 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v520 int32
	_ = v520
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int64
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1182 int32
	_ = v1182
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1462 int32
	_ = v1462
	var v1470 int32
	_ = v1470
	var v1478 int32
	_ = v1478
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1496 int32
	_ = v1496
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1518 int32
	_ = v1518
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1542 int32
	_ = v1542
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1570 int32
	_ = v1570
	var v1574 int32
	_ = v1574
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1590 int32
	_ = v1590
	var v1594 int32
	_ = v1594
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1612 int64
	_ = v1612
	var v1614 int64
	_ = v1614
	var v1616 int64
	_ = v1616
	var v1618 int64
	_ = v1618
	var v1620 int64
	_ = v1620
	var v1622 int64
	_ = v1622
	var v1624 int64
	_ = v1624
	var v1626 int64
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1636 int32
	_ = v1636
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
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
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1728 int32
	_ = v1728
	var v1733 int32
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
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
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1831 int32
	_ = v1831
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1872 int32
	_ = v1872
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1929 int32
	_ = v1929
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2013 int32
	_ = v2013
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2029 int32
	_ = v2029
	var v2034 int32
	_ = v2034
	var v2039 int32
	_ = v2039
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2065 int32
	_ = v2065
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2083 int32
	_ = v2083
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(224)
	m.G0 = v17
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
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L26
	} else {
		goto L544
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L26
	} else {
		goto L528
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L26
	} else {
		goto L525
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L26
	} else {
		goto L509
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L26
	} else {
		goto L504
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L26
	} else {
		goto L499
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L26
	} else {
		goto L494
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1910 = m.ExcPending
	if v1910 != 0 {
		goto L26
	} else {
		goto L490
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L26
	} else {
		goto L487
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L26
	} else {
		goto L483
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L26
	} else {
		goto L479
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L26
	} else {
		goto L475
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L26
	} else {
		goto L471
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L26
	} else {
		goto L467
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L26
	} else {
		goto L462
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L26
	} else {
		goto L456
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L26
	} else {
		goto L451
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L26
	} else {
		goto L445
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L26
	} else {
		goto L440
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L26
	} else {
		goto L435
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L26
	} else {
		goto L430
	}
L22:
	;
	m.G0 = v17 + int32(224)
	return v1636
L23:
	;
	v24 = F_pstrdup(m, int32(757108))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v1636 = v7
	goto L22
L28:
	;
	v33 = F_strlen(m, l1)
	mBase = m.M
	if v33 != l2 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v35 {
	case 0:
		goto L34
	case 1:
		goto L33
	default:
		goto L32
	}
L30:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1602 != 0 {
		goto L425
	} else {
		goto L426
	}
L31:
	;
	v1580 = int32(0)
	v1581 = int32(2)
	if l5 == v1580 {
		v1590 = v1580
		v1594 = v1581
		goto L30
	} else {
		goto L423
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L26
	} else {
		goto L420
	}
L33:
	;
	v557 = F_pstrdup(m, l1)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L26
	} else {
		goto L147
	}
L34:
	;
	v36 = F_pstrdup(m, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v36
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v39)
	switch v39 - int32(110) {
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
	v230 = v95 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v230
	v232 = F_pstrdup(m, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L26
	} else {
		goto L90
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L26
	} else {
		goto L84
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L26
	} else {
		goto L66
	}
L39:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v96 == int32(44) {
		goto L36
	} else {
		goto L58
	}
L40:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v59 == int32(0) {
		goto L15
	} else {
		goto L47
	}
L41:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v51 == int32(1) {
		goto L17
	} else {
		goto L45
	}
L42:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v43 == int32(1) {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v46 != int32(44) {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v95 = v36 + int32(2)
	goto L39
L45:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v54 != int32(44) {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v95 = v36 + int32(2)
	goto L39
L47:
	;
	v65 = F_read_attr_value(m, v17+int32(192), int32(112))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v67 = int32(88549)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[354])))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v71 == int32(0) {
		v90 = v70
		v91 = v71
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v91-v90 != 0 {
		goto L38
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	if v70 != v71 {
		v90 = v70
		v91 = v71
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v75 = v65
	v76 = v67
	goto L53
L53:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v80 == int32(0) {
		v90 = v79
		v91 = v80
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v90 = v79
	v91 = v80
	goto L50
L55:
	;
	v83 = int32(1)
	if v79 == v80 {
		v75 = v75 + v83
		v76 = v76 + v83
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	v95 = v93
	goto L39
L58:
	;
	if v96 == int32(97) {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	v112 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95))))
	F_sanitize_char_2(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(4414000)
	F_errdetail(m, int32(641876), v17+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(497038), int32(1081), int32(403802))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L26
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L26
	} else {
		goto L67
	}
L67:
	;
	v136 = int32(0)
	goto L68
L68:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+v65))))
	if v150 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+uint32(_consts[355]))) = uint8(v191)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = int32(4414016)
	F_errmsg(m, int32(713913), v17+int32(80))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L26
	} else {
		goto L82
	}
L70:
	;
	goto L69
L71:
	;
	v186 = v136
	goto L70
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(94)) <= base.Ui32((v150-int32(33))&int32(255)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v162 = int32(63)
	goto L76
L75:
	;
	v162 = v150
	goto L76
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+uint32(_consts[355]))) = uint8(v162)
	v165 = v136 | int32(1)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v165))))
	if v167 == int32(0) {
		v186 = v165
		goto L70
	} else {
		goto L77
	}
L77:
	;
	if base.Ui32(int32(94)) <= base.Ui32((v167-int32(33))&int32(255)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v179 = int32(63)
	goto L80
L79:
	;
	v179 = v167
	goto L80
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+uint32(_consts[355]))) = uint8(v179)
	v181 = int32(30)
	v183 = v136 + int32(2)
	if v183 != v181 {
		v136 = v183
		goto L68
	} else {
		goto L81
	}
L81:
	;
	v186 = v181
	goto L70
L82:
	;
	F_errfinish(m, int32(497038), int32(1059), int32(403802))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L26
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L26
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L26
	} else {
		goto L86
	}
L86:
	;
	v216 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36))))
	F_sanitize_char_2(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L26
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(4414000)
	F_errdetail(m, int32(666444), v17)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L26
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(497038), int32(1066), int32(403802))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L26
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v232
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v235 == int32(109) {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	v241 = F_read_attr_value(m, v17+int32(192), int32(110))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L26
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v241
	v247 = F_read_attr_value(m, v17+int32(192), int32(114))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L26
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v247
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v251 = base.I32_extend8_s(v250)
	if int32(33) <= v251 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v256 = v251
	v260 = v247
	goto L97
L95:
	;
	v282 = v251
	goto L96
L96:
	;
	if v282 != 0 {
		goto L12
	} else {
		goto L102
	}
L97:
	;
	v269 = v256 & int32(255)
	if v269 == int32(44) {
		goto L12
	} else {
		goto L99
	}
L98:
	;
	v282 = v277
	goto L96
L99:
	;
	if v269 == int32(127) {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	v275 = v260 + int32(1)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v277 = base.I32_extend8_s(v276)
	if int32(32) < v277 {
		v256 = v277
		v260 = v275
		goto L97
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v295 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	goto L106
L104:
	;
	goto L105
L105:
	;
	v334 = int32(0)
	v338 = m.G0
	v340 = v338 - int32(16)
	m.G0 = v340
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v334
	v346 = F_open(m, int32(288357), v334, v340)
	mBase = m.M
	if v346 != int32(-1) {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v313 = F_read_any_attr(m, v17+int32(192), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L26
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v316 != 0 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	if v379 == int32(0) {
		goto L11
	} else {
		goto L123
	}
L111:
	;
	goto L115
L112:
	;
	v379 = v334
	goto L113
L113:
	;
	m.G0 = v340 + int32(16)
	goto L110
L114:
	;
	v374 = F_close(m, v346)
	mBase = m.M
	v379 = v372
	goto L113
L115:
	;
	v352 = v17 + int32(192)
	v353 = int32(18)
	goto L116
L116:
	;
	v358 = F_read(m, v346, v352, v353)
	mBase = m.M
	if v358 <= int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v372 = int32(1)
	goto L114
L118:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v362 == int32(27) {
		goto L116
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v367 = v353 - v358
	if v367 != 0 {
		v352 = v352 + v358
		v353 = v367
		goto L116
	} else {
		goto L122
	}
L121:
	;
	v372 = int32(0)
	goto L114
L122:
	;
	goto L117
L123:
	;
	v390 = base.I32_div_s(int32(20), int32(3))
	v392 = v390 << (uint(int32(2)) % 32)
	goto L124
L124:
	;
	v395 = F_palloc(m, v392+int32(1))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L26
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v395
	v399 = v17 + int32(192)
	v404 = v17 + int32(210)
	if base.Ui32(v399) < base.Ui32(v404) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	if v529 < int32(0) {
		goto L10
	} else {
		goto L144
	}
L127:
	;
	v520 = F___memset(m, v395, int32(0), v392)
	mBase = m.M
	v529 = int32(-1)
	goto L126
L128:
	;
	if v392 < v463-v395+int32(4) {
		goto L127
	} else {
		goto L140
	}
L129:
	;
	v408 = v399
	v409 = int32(0)
	v412 = v395
	v413 = int32(2)
	goto L132
L130:
	;
	v474 = v395
	goto L131
L131:
	;
	v529 = v474 - v395
	goto L126
L132:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	v419 = v415<<(uint(v413<<(uint(int32(3))%32))%32) | v409
	if int32(0) < v413 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v464 != int32(2) {
		goto L128
	} else {
		goto L139
	}
L134:
	;
	v462 = v419
	v463 = v412
	v464 = v413 - int32(1)
	goto L136
L135:
	;
	if v392 < v412-v395+int32(4) {
		goto L127
	} else {
		goto L137
	}
L136:
	;
	v466 = v408 + int32(1)
	if v466 != v404 {
		v408 = v466
		v409 = v462
		v412 = v463
		v413 = v464
		goto L132
	} else {
		goto L138
	}
L137:
	;
	v428 = int32(63)
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419&v428)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v412)+3)) = uint8(v432)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v419)>>(uint(int32(6))%32))&v428)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v412)+2)) = uint8(v440)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v419)>>(uint(int32(12))%32))&v428)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v412)+1)) = uint8(v448)
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v419)>>(uint(int32(18))%32))&v428)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v412))) = uint8(v456)
	v462 = int32(0)
	v463 = v412 + int32(4)
	v464 = int32(2)
	goto L136
L138:
	;
	goto L133
L139:
	;
	v474 = v463
	goto L131
L140:
	;
	v484 = int32(63)
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v462)>>(uint(int32(12))%32))&v484)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)) = uint8(v488)
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v462)>>(uint(int32(18))%32))&v484)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v496)
	if v464 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v462)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[356]))))
	v508 = v507
	goto L143
L142:
	;
	v508 = int32(61)
	goto L143
L143:
	;
	v509 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v463)+3)) = uint8(v509)
	*(*uint8)(unsafe.Add(mBase, uint32(v463)+2)) = uint8(v508)
	v529 = v463 + int32(4) - v395
	goto L126
L144:
	;
	v532 = int32(0)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*uint8)(unsafe.Add(mBase, uint32(v533+v529))) = uint8(v532)
	v537 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v538
	*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = base.I64_rotl(v537, int64(32))
	v548 = F_psprintf(m, int32(466248), v17+int32(32))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L26
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v548
	v551 = F_pstrdup(m, v548)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L26
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v551
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v1590 = int32(0)
	v1594 = v532
	goto L30
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v557
	v563 = F_read_attr_value(m, v17+int32(192), int32(99))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L26
	} else {
		goto L148
	}
L148:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v565 == int32(1) {
		goto L9
	} else {
		goto L149
	}
L149:
	;
	v568 = int32(114102)
	v571 = int32(*(*uint8)(unsafe.Add(mBase, _consts[357])))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	if v572 == int32(0) {
		v591 = v571
		v592 = v572
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v631 = F_read_attr_value(m, v17+int32(192), int32(114))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L26
	} else {
		goto L173
	}
L151:
	;
	if v592-v591 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L152:
	;
	goto L151
L153:
	;
	if v571 != v572 {
		v591 = v571
		v592 = v572
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v576 = v563
	v577 = v568
	goto L155
L155:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576)+1)))
	if v581 == int32(0) {
		v591 = v580
		v592 = v581
		goto L152
	} else {
		goto L157
	}
L156:
	;
	v591 = v580
	v592 = v581
	goto L152
L157:
	;
	v584 = int32(1)
	if v580 == v581 {
		v576 = v576 + v584
		v577 = v577 + v584
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v596 == int32(110) {
		goto L150
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v599 = int32(114413)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, _consts[358])))
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	if v603 == int32(0) {
		v622 = v602
		v623 = v603
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L161
L163:
	;
	if v623-v622 != 0 {
		goto L8
	} else {
		goto L171
	}
L164:
	;
	goto L163
L165:
	;
	if v602 != v603 {
		v622 = v602
		v623 = v603
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v607 = v563
	v608 = v599
	goto L167
L167:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608)+1)))
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+1)))
	if v612 == int32(0) {
		v622 = v611
		v623 = v612
		goto L164
	} else {
		goto L169
	}
L168:
	;
	v622 = v611
	v623 = v612
	goto L164
L169:
	;
	v615 = int32(1)
	if v611 == v612 {
		v607 = v607 + v615
		v608 = v608 + v615
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v625 != int32(121) {
		goto L8
	} else {
		goto L172
	}
L172:
	;
	goto L150
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v631
	goto L174
L174:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	v653 = F_read_any_attr(m, v17+int32(192), v17+int32(160))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L26
	} else {
		goto L176
	}
L175:
	;
	v658 = F_strlen(m, v653)
	mBase = m.M
	v662 = v658 * int32(3) >> (uint(int32(2)) % 32)
	goto L178
L176:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+160)))
	if v655 != int32(112) {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v663 = F_palloc(m, v662)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L26
	} else {
		goto L179
	}
L179:
	;
	v665 = F_strlen(m, v653)
	mBase = m.M
	v666 = int32(0)
	v673 = v653 + v665
	if base.Ui32(v653) < base.Ui32(v673) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v853 != v854 {
		goto L7
	} else {
		goto L224
	}
L181:
	;
	v840 = F___memset(m, v663, int32(0), v662)
	mBase = m.M
	v853 = int32(-1)
	goto L180
L182:
	;
	v675 = v653
	v679 = v666
	v680 = v663
	v681 = v666
	v683 = v666
	goto L185
L183:
	;
	v824 = v663
	goto L184
L184:
	;
	v853 = v824 - v663
	goto L180
L185:
	;
	v687 = v675 + int32(1)
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	if v688 != int32(61) {
		goto L193
	} else {
		goto L194
	}
L186:
	;
	if v812 != 0 {
		goto L181
	} else {
		goto L223
	}
L187:
	;
	if v810 != v673 {
		v675 = v810
		v679 = v812
		v680 = v813
		v681 = v814
		v683 = v816
		goto L185
	} else {
		goto L222
	}
L188:
	;
	if v662 < v680-v663+int32(1) {
		goto L181
	} else {
		goto L209
	}
L189:
	;
	v766 = int32(2)
	v767 = v687
	v770 = v683 << (uint(int32(6)) % 32)
	goto L188
L190:
	;
	v755 = v752 + v751<<(uint(int32(6))%32)
	v757 = v748 + int32(1)
	if v757 == int32(4) {
		v766 = v749
		v767 = v750
		v770 = v755
		goto L188
	} else {
		goto L208
	}
L191:
	;
	v748 = int32(3)
	v749 = int32(1)
	v750 = v675 + int32(2)
	v751 = v706
	v752 = v701
	goto L190
L192:
	;
	if base.Ui32(int32(125)) < base.Ui32((v725-int32(1))&int32(255)) {
		goto L181
	} else {
		goto L206
	}
L193:
	;
	v692 = v688 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v692) {
		v725 = v688
		v726 = v679
		v727 = v681
		v728 = v687
		v729 = v683
		goto L192
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v701 = int32(0)
	if v681 != 0 {
		v748 = v679
		v749 = v681
		v750 = v687
		v751 = v683
		v752 = v701
		goto L190
	} else {
		goto L198
	}
L196:
	;
	if int32(1)<<(uint(v692)%32)&int32(8388627) == int32(0) {
		v725 = v688
		v726 = v679
		v727 = v681
		v728 = v687
		v729 = v683
		goto L192
	} else {
		goto L197
	}
L197:
	;
	goto L181
L198:
	;
	switch v679 - int32(2) {
	case 0:
		goto L199
	case 1:
		goto L189
	default:
		goto L181
	}
L199:
	;
	if v687 == v673 {
		goto L181
	} else {
		goto L200
	}
L200:
	;
	v706 = v683 << (uint(int32(6)) % 32)
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687))))
	if v707 == int32(61) {
		goto L191
	} else {
		goto L201
	}
L201:
	;
	v711 = v707 - int32(9)
	if int32(1)<<(uint(v711)%32)&int32(8388627) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v719 = base.B2i32(base.Ui32(v711) <= base.Ui32(int32(23)))
	goto L204
L203:
	;
	v719 = int32(0)
	goto L204
L204:
	;
	if v719 != 0 {
		goto L181
	} else {
		goto L205
	}
L205:
	;
	v725 = v707
	v726 = int32(3)
	v727 = int32(1)
	v728 = v675 + int32(2)
	v729 = v706
	goto L192
L206:
	;
	v739 = int32(*(*int8)(unsafe.Add(mBase, uint32(v725)+uint32(_consts[359]))))
	if v739 < int32(0) {
		goto L181
	} else {
		goto L207
	}
L207:
	;
	v748 = v726
	v749 = v727
	v750 = v728
	v751 = v729
	v752 = v739
	goto L190
L208:
	;
	v810 = v750
	v812 = v757
	v813 = v680
	v814 = v749
	v816 = v755
	goto L187
L209:
	;
	v776 = int32(base.Ui32(v770) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v680))) = uint8(v776)
	v779 = v680 + int32(1)
	if base.Ui32(v766) < base.Ui32(int32(2)) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v783 = v766
	goto L212
L211:
	;
	v783 = int32(0)
	goto L212
L212:
	;
	if v783 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	if v662 < v779-v663+int32(1) {
		goto L181
	} else {
		goto L216
	}
L214:
	;
	v795 = v779
	goto L215
L215:
	;
	v796 = int32(0)
	if v766 == v796 {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	v791 = int32(base.Ui32(v770) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v680)+1)) = uint8(v791)
	v795 = v680 + int32(2)
	goto L215
L217:
	;
	v810 = v767
	v812 = int32(0)
	v813 = v808
	v814 = v766
	v816 = v796
	goto L187
L218:
	;
	if v662 < v795-v663+int32(1) {
		goto L181
	} else {
		goto L221
	}
L219:
	;
	if base.Ui32(int32(3)) <= base.Ui32(v766) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v808 = v795
	goto L217
L221:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v795))) = uint8(v770)
	v808 = v795 + int32(1)
	goto L217
L222:
	;
	goto L186
L223:
	;
	v824 = v813
	goto L184
L224:
	;
	v857 = l0 + int32(152)
	if v853 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	F_pfree(m, v663)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L26
	} else {
		goto L229
	}
L226:
	;
	v858 = F__emscripten_memcpy_bulkmem(m, v857, v663, v853)
	mBase = m.M
	v859 = v858
	goto L228
L227:
	;
	v859 = v857
	goto L228
L228:
	;
	goto L225
L229:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862))))
	if v863 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v865 = F_palloc(m, v648-v557)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L26
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v865
	v870 = v557 ^ int32(-1) + v648
	if v870 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v875 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v873+v870))) = uint8(v875)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v878 = F_strlen(m, v877)
	mBase = m.M
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v880 = F_strlen(m, v879)
	mBase = m.M
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v882 = F_strlen(m, v881)
	mBase = m.M
	if v882 != v880+v878 {
		goto L5
	} else {
		goto L236
	}
L233:
	;
	v871 = F__emscripten_memcpy_bulkmem(m, v865, l1, v870)
	mBase = m.M
	goto L235
L234:
	;
	goto L235
L235:
	;
	goto L232
L236:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v878) {
		goto L240
	} else {
		goto L241
	}
L237:
	;
	if v946 != 0 {
		goto L5
	} else {
		goto L255
	}
L238:
	;
	v946 = int32(0)
	goto L237
L239:
	;
	v920 = v915
	v921 = v916
	v922 = v917
	goto L249
L240:
	;
	if (v881|v877)&int32(3) != 0 {
		v915 = v881
		v916 = v877
		v917 = v878
		goto L239
	} else {
		goto L243
	}
L241:
	;
	v908 = v881
	v909 = v877
	v910 = v878
	goto L242
L242:
	;
	if v910 == int32(0) {
		goto L238
	} else {
		goto L248
	}
L243:
	;
	v892 = v881
	v893 = v877
	v894 = v878
	goto L244
L244:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v893)))
	if v897 != v898 {
		v915 = v892
		v916 = v893
		v917 = v894
		goto L239
	} else {
		goto L246
	}
L245:
	;
	v908 = v903
	v909 = v901
	v910 = v905
	goto L242
L246:
	;
	v900 = int32(4)
	v901 = v893 + v900
	v903 = v892 + v900
	v905 = v894 - v900
	if base.Ui32(int32(3)) < base.Ui32(v905) {
		v892 = v903
		v893 = v901
		v894 = v905
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v915 = v908
	v916 = v909
	v917 = v910
	goto L239
L249:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920))))
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921))))
	if v925 == v926 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v946 = v925 - v926
	goto L237
L251:
	;
	v928 = int32(1)
	v933 = v922 - v928
	if v933 != 0 {
		v920 = v920 + v928
		v921 = v921 + v928
		v922 = v933
		goto L249
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	goto L250
L254:
	;
	goto L238
L255:
	;
	v947 = v881 + v878
	if base.Ui32(int32(4)) <= base.Ui32(v880) {
		goto L259
	} else {
		goto L260
	}
L256:
	;
	if v1009 != 0 {
		goto L5
	} else {
		goto L274
	}
L257:
	;
	v1009 = int32(0)
	goto L256
L258:
	;
	v983 = v978
	v984 = v979
	v985 = v980
	goto L268
L259:
	;
	if (v947|v879)&int32(3) != 0 {
		v978 = v947
		v979 = v879
		v980 = v880
		goto L258
	} else {
		goto L262
	}
L260:
	;
	v971 = v947
	v972 = v879
	v973 = v880
	goto L261
L261:
	;
	if v973 == int32(0) {
		goto L257
	} else {
		goto L267
	}
L262:
	;
	v955 = v947
	v956 = v879
	v957 = v880
	goto L263
L263:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v955)))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v956)))
	if v960 != v961 {
		v978 = v955
		v979 = v956
		v980 = v957
		goto L258
	} else {
		goto L265
	}
L264:
	;
	v971 = v966
	v972 = v964
	v973 = v968
	goto L261
L265:
	;
	v963 = int32(4)
	v964 = v956 + v963
	v966 = v955 + v963
	v968 = v957 - v963
	if base.Ui32(int32(3)) < base.Ui32(v968) {
		v955 = v966
		v956 = v964
		v957 = v968
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	v978 = v971
	v979 = v972
	v980 = v973
	goto L258
L268:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983))))
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984))))
	if v988 == v989 {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1009 = v988 - v989
	goto L256
L270:
	;
	v991 = int32(1)
	v996 = v985 - v991
	if v996 != 0 {
		v983 = v983 + v991
		v984 = v984 + v991
		v985 = v996
		goto L268
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	goto L269
L273:
	;
	goto L257
L274:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1011 = F_pg_hmac_create(m, v1010)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L26
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = int32(0)
	v1016 = l0 - int32(-64)
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1018 = F_pg_hmac_init(m, v1011, v1016, v1017)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L26
	} else {
		goto L276
	}
L276:
	;
	if v1018 < int32(0) {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1023 = F_strlen(m, v1022)
	mBase = m.M
	if v1011 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	if v1038 < int32(0) {
		goto L4
	} else {
		goto L285
	}
L279:
	;
	v1038 = int32(-1)
	goto L278
L280:
	;
	goto L281
L281:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1028 = F_pg_cryptohash_update(m, v1027, v1022, v1023)
	mBase = m.M
	if int32(0) <= v1028 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1038 = int32(0)
	goto L278
L283:
	;
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = int32(2)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1035 = F_pg_cryptohash_error(m, v1034)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+12)) = v1035
	v1038 = int32(-1)
	goto L278
L285:
	;
	if v1011 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	if v1057 < int32(0) {
		goto L4
	} else {
		goto L293
	}
L287:
	;
	v1057 = int32(-1)
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1047 = F_pg_cryptohash_update(m, v1046, int32(669711), int32(1))
	mBase = m.M
	if int32(0) <= v1047 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1057 = int32(0)
	goto L286
L291:
	;
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = int32(2)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1054 = F_pg_cryptohash_error(m, v1053)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+12)) = v1054
	v1057 = int32(-1)
	goto L286
L293:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v1061 = F_strlen(m, v1060)
	mBase = m.M
	if v1011 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	if v1076 < int32(0) {
		goto L4
	} else {
		goto L301
	}
L295:
	;
	v1076 = int32(-1)
	goto L294
L296:
	;
	goto L297
L297:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1066 = F_pg_cryptohash_update(m, v1065, v1060, v1061)
	mBase = m.M
	if int32(0) <= v1066 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1076 = int32(0)
	goto L294
L299:
	;
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = int32(2)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1073 = F_pg_cryptohash_error(m, v1072)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+12)) = v1073
	v1076 = int32(-1)
	goto L294
L301:
	;
	if v1011 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	if v1095 < int32(0) {
		goto L4
	} else {
		goto L309
	}
L303:
	;
	v1095 = int32(-1)
	goto L302
L304:
	;
	goto L305
L305:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1085 = F_pg_cryptohash_update(m, v1084, int32(669711), int32(1))
	mBase = m.M
	if int32(0) <= v1085 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1095 = int32(0)
	goto L302
L307:
	;
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = int32(2)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1092 = F_pg_cryptohash_error(m, v1091)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+12)) = v1092
	v1095 = int32(-1)
	goto L302
L309:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1099 = F_strlen(m, v1098)
	mBase = m.M
	if v1011 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	if v1114 < int32(0) {
		goto L4
	} else {
		goto L317
	}
L311:
	;
	v1114 = int32(-1)
	goto L310
L312:
	;
	goto L313
L313:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1104 = F_pg_cryptohash_update(m, v1103, v1098, v1099)
	mBase = m.M
	if int32(0) <= v1104 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1114 = int32(0)
	goto L310
L315:
	;
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = int32(2)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1011)))
	v1111 = F_pg_cryptohash_error(m, v1110)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+12)) = v1111
	v1114 = int32(-1)
	goto L310
L317:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1120 = F_pg_hmac_final(m, v1011, v17+int32(192), v1119)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L26
	} else {
		goto L318
	}
L318:
	;
	if v1120 < int32(0) {
		goto L4
	} else {
		goto L319
	}
L319:
	;
	F_pg_hmac_free(m, v1011)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L26
	} else {
		goto L320
	}
L320:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1126 <= int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1224 = F_scram_H(m, l0+int32(32), v1219, v1126, v17+int32(160), v17+int32(156))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L26
	} else {
		goto L330
	}
L322:
	;
	v1130 = l0 + int32(32)
	v1131 = int32(0)
	if v1126 != int32(1) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1143 = v1131
	v1146 = int32(0)
	goto L326
L324:
	;
	v1182 = v1131
	goto L325
L325:
	;
	if v1126&int32(1) == int32(0) {
		goto L321
	} else {
		goto L329
	}
L326:
	;
	v1153 = v17 + int32(192)
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153+v1143))))
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143+v859))))
	v1158 = v1155 ^ v1157
	*(*uint8)(unsafe.Add(mBase, uint32(v1130+v1143))) = uint8(v1158)
	v1161 = v1143 | int32(1)
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1153+v1161))))
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1161+v859))))
	v1169 = v1166 ^ v1168
	*(*uint8)(unsafe.Add(mBase, uint32(v1130+v1161))) = uint8(v1169)
	v1171 = int32(2)
	v1172 = v1143 + v1171
	v1174 = v1146 + v1171
	if v1174 != v1126&int32(2147483646) {
		v1143 = v1172
		v1146 = v1174
		goto L326
	} else {
		goto L328
	}
L327:
	;
	v1182 = v1172
	goto L325
L328:
	;
	goto L327
L329:
	;
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(192)+v1182))))
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182+v859))))
	v1201 = v1198 ^ v1200
	*(*uint8)(unsafe.Add(mBase, uint32(v1130+v1182))) = uint8(v1201)
	goto L321
L330:
	;
	if v1224 < int32(0) {
		goto L3
	} else {
		goto L331
	}
L331:
	;
	v1229 = v17 + int32(160)
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(4)) <= base.Ui32(v1230) {
		goto L335
	} else {
		goto L336
	}
L332:
	;
	if v1292 != 0 {
		goto L31
	} else {
		goto L350
	}
L333:
	;
	v1292 = int32(0)
	goto L332
L334:
	;
	v1266 = v1261
	v1267 = v1262
	v1268 = v1263
	goto L344
L335:
	;
	if (v1229|v1016)&int32(3) != 0 {
		v1261 = v1229
		v1262 = v1016
		v1263 = v1230
		goto L334
	} else {
		goto L338
	}
L336:
	;
	v1254 = v1229
	v1255 = v1016
	v1256 = v1230
	goto L337
L337:
	;
	if v1256 == int32(0) {
		goto L333
	} else {
		goto L343
	}
L338:
	;
	v1238 = v1229
	v1239 = v1016
	v1240 = v1230
	goto L339
L339:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1238)))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1239)))
	if v1243 != v1244 {
		v1261 = v1238
		v1262 = v1239
		v1263 = v1240
		goto L334
	} else {
		goto L341
	}
L340:
	;
	v1254 = v1249
	v1255 = v1247
	v1256 = v1251
	goto L337
L341:
	;
	v1246 = int32(4)
	v1247 = v1239 + v1246
	v1249 = v1238 + v1246
	v1251 = v1240 - v1246
	if base.Ui32(int32(3)) < base.Ui32(v1251) {
		v1238 = v1249
		v1239 = v1247
		v1240 = v1251
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	v1261 = v1254
	v1262 = v1255
	v1263 = v1256
	goto L334
L344:
	;
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1266))))
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1267))))
	if v1271 == v1272 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	v1292 = v1271 - v1272
	goto L332
L346:
	;
	v1274 = int32(1)
	v1279 = v1268 - v1274
	if v1279 != 0 {
		v1266 = v1266 + v1274
		v1267 = v1267 + v1274
		v1268 = v1279
		goto L344
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	goto L345
L349:
	;
	goto L333
L350:
	;
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+192)))
	if v1293 != 0 {
		goto L31
	} else {
		goto L351
	}
L351:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1295 = F_pg_hmac_create(m, v1294)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L26
	} else {
		goto L352
	}
L352:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1300 = F_pg_hmac_init(m, v1295, l0+int32(96), v1299)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L26
	} else {
		goto L353
	}
L353:
	;
	if v1300 < int32(0) {
		goto L2
	} else {
		goto L354
	}
L354:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v1305 = F_strlen(m, v1304)
	mBase = m.M
	if v1295 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	if v1320 < int32(0) {
		goto L2
	} else {
		goto L362
	}
L356:
	;
	v1320 = int32(-1)
	goto L355
L357:
	;
	goto L358
L358:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1310 = F_pg_cryptohash_update(m, v1309, v1304, v1305)
	mBase = m.M
	if int32(0) <= v1310 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1320 = int32(0)
	goto L355
L360:
	;
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+8)) = int32(2)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1317 = F_pg_cryptohash_error(m, v1316)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+12)) = v1317
	v1320 = int32(-1)
	goto L355
L362:
	;
	if v1295 == int32(0) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	if v1339 < int32(0) {
		goto L2
	} else {
		goto L370
	}
L364:
	;
	v1339 = int32(-1)
	goto L363
L365:
	;
	goto L366
L366:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1329 = F_pg_cryptohash_update(m, v1328, int32(669711), int32(1))
	mBase = m.M
	if int32(0) <= v1329 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1339 = int32(0)
	goto L363
L368:
	;
	goto L369
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+8)) = int32(2)
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1336 = F_pg_cryptohash_error(m, v1335)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+12)) = v1336
	v1339 = int32(-1)
	goto L363
L370:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v1343 = F_strlen(m, v1342)
	mBase = m.M
	if v1295 == int32(0) {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	if v1358 < int32(0) {
		goto L2
	} else {
		goto L378
	}
L372:
	;
	v1358 = int32(-1)
	goto L371
L373:
	;
	goto L374
L374:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1348 = F_pg_cryptohash_update(m, v1347, v1342, v1343)
	mBase = m.M
	if int32(0) <= v1348 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1358 = int32(0)
	goto L371
L376:
	;
	goto L377
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+8)) = int32(2)
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1355 = F_pg_cryptohash_error(m, v1354)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+12)) = v1355
	v1358 = int32(-1)
	goto L371
L378:
	;
	if v1295 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	if v1377 < int32(0) {
		goto L2
	} else {
		goto L386
	}
L380:
	;
	v1377 = int32(-1)
	goto L379
L381:
	;
	goto L382
L382:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1367 = F_pg_cryptohash_update(m, v1366, int32(669711), int32(1))
	mBase = m.M
	if int32(0) <= v1367 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1377 = int32(0)
	goto L379
L384:
	;
	goto L385
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+8)) = int32(2)
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1374 = F_pg_cryptohash_error(m, v1373)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+12)) = v1374
	v1377 = int32(-1)
	goto L379
L386:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v1381 = F_strlen(m, v1380)
	mBase = m.M
	if v1295 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L387:
	;
	if v1396 < int32(0) {
		goto L2
	} else {
		goto L394
	}
L388:
	;
	v1396 = int32(-1)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1386 = F_pg_cryptohash_update(m, v1385, v1380, v1381)
	mBase = m.M
	if int32(0) <= v1386 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1396 = int32(0)
	goto L387
L392:
	;
	goto L393
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+8)) = int32(2)
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1295)))
	v1393 = F_pg_cryptohash_error(m, v1392)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1295)+12)) = v1393
	v1396 = int32(-1)
	goto L387
L394:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1402 = F_pg_hmac_final(m, v1295, v17+int32(192), v1401)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L26
	} else {
		goto L395
	}
L395:
	;
	if v1402 < int32(0) {
		goto L2
	} else {
		goto L396
	}
L396:
	;
	F_pg_hmac_free(m, v1295)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L26
	} else {
		goto L397
	}
L397:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1410 = int32(2)
	v1413 = base.I32_div_s(v1409+v1410, int32(3))
	v1415 = v1413 << (uint(v1410) % 32)
	goto L398
L398:
	;
	v1418 = F_palloc(m, v1415+int32(1))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L26
	} else {
		goto L399
	}
L399:
	;
	v1421 = v17 + int32(192)
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1426 = v1421 + v1422
	if base.Ui32(v1421) < base.Ui32(v1426) {
		goto L403
	} else {
		goto L404
	}
L400:
	;
	if v1551 < int32(0) {
		goto L1
	} else {
		goto L418
	}
L401:
	;
	v1542 = F___memset(m, v1418, int32(0), v1415)
	mBase = m.M
	v1551 = int32(-1)
	goto L400
L402:
	;
	if v1415 < v1485-v1418+int32(4) {
		goto L401
	} else {
		goto L414
	}
L403:
	;
	v1430 = v1421
	v1431 = int32(0)
	v1434 = v1418
	v1435 = int32(2)
	goto L406
L404:
	;
	v1496 = v1418
	goto L405
L405:
	;
	v1551 = v1496 - v1418
	goto L400
L406:
	;
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1430))))
	v1441 = v1437<<(uint(v1435<<(uint(int32(3))%32))%32) | v1431
	if int32(0) < v1435 {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	if v1486 != int32(2) {
		goto L402
	} else {
		goto L413
	}
L408:
	;
	v1484 = v1441
	v1485 = v1434
	v1486 = v1435 - int32(1)
	goto L410
L409:
	;
	if v1415 < v1434-v1418+int32(4) {
		goto L401
	} else {
		goto L411
	}
L410:
	;
	v1488 = v1430 + int32(1)
	if v1488 != v1426 {
		v1430 = v1488
		v1431 = v1484
		v1434 = v1485
		v1435 = v1486
		goto L406
	} else {
		goto L412
	}
L411:
	;
	v1450 = int32(63)
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441&v1450)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1434)+3)) = uint8(v1454)
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1441)>>(uint(int32(6))%32))&v1450)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1434)+2)) = uint8(v1462)
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1441)>>(uint(int32(12))%32))&v1450)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1434)+1)) = uint8(v1470)
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1441)>>(uint(int32(18))%32))&v1450)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1434))) = uint8(v1478)
	v1484 = int32(0)
	v1485 = v1434 + int32(4)
	v1486 = int32(2)
	goto L410
L412:
	;
	goto L407
L413:
	;
	v1496 = v1485
	goto L405
L414:
	;
	v1506 = int32(63)
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1484)>>(uint(int32(12))%32))&v1506)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1485)+1)) = uint8(v1510)
	v1518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1484)>>(uint(int32(18))%32))&v1506)+uint32(_consts[356]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1485))) = uint8(v1518)
	if v1486 == int32(0) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1484)>>(uint(int32(6))%32))&int32(63))+uint32(_consts[356]))))
	v1530 = v1529
	goto L417
L416:
	;
	v1530 = int32(61)
	goto L417
L417:
	;
	v1531 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v1485)+3)) = uint8(v1531)
	*(*uint8)(unsafe.Add(mBase, uint32(v1485)+2)) = uint8(v1530)
	v1551 = v1485 + int32(4) - v1418
	goto L400
L418:
	;
	v1555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1418+v1551))) = uint8(v1555)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v1418
	v1561 = F_psprintf(m, int32(175975), v17+int32(144))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L26
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1561
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
	v1590 = int32(1)
	v1594 = int32(1)
	goto L30
L420:
	;
	F_errmsg_internal(m, int32(352814), int32(0))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L26
	} else {
		goto L421
	}
L421:
	;
	F_errfinish(m, int32(497038), int32(457), int32(402644))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L26
	} else {
		goto L422
	}
L422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L423:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v1584 == int32(0) {
		v1590 = v1580
		v1594 = v1581
		goto L30
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1584
	v1590 = v1580
	v1594 = v1581
	goto L30
L425:
	;
	v1603 = F_strlen(m, v1602)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1603
	goto L427
L426:
	;
	goto L427
L427:
	;
	if v1590 == int32(0) {
		v1636 = v1594
		goto L22
	} else {
		goto L428
	}
L428:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1607 != int32(2) {
		v1636 = v1594
		goto L22
	} else {
		goto L429
	}
L429:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	v1612 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+416)) = v1612
	v1614 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+440)) = v1614
	v1616 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+432)) = v1616
	v1618 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+424)) = v1618
	v1620 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+464)) = v1620
	v1622 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+456)) = v1622
	v1624 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+448)) = v1624
	v1626 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v1627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1611)+480)) = uint8(v1627)
	*(*int64)(unsafe.Add(mBase, uint32(v1611)+472)) = v1626
	v1636 = v1594
	goto L22
L430:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L26
	} else {
		goto L431
	}
L431:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L26
	} else {
		goto L432
	}
L432:
	;
	F_errdetail(m, int32(570628), int32(0))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L26
	} else {
		goto L433
	}
L433:
	;
	F_errfinish(m, int32(497038), int32(383), int32(402644))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L26
	} else {
		goto L434
	}
L434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L435:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L26
	} else {
		goto L436
	}
L436:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L26
	} else {
		goto L437
	}
L437:
	;
	F_errdetail(m, int32(623974), int32(0))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L26
	} else {
		goto L438
	}
L438:
	;
	F_errfinish(m, int32(497038), int32(388), int32(402644))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L26
	} else {
		goto L439
	}
L439:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L440:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L26
	} else {
		goto L441
	}
L441:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L26
	} else {
		goto L442
	}
L442:
	;
	F_errdetail(m, int32(654800), int32(0))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L26
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(497038), int32(996), int32(403802))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L26
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L26
	} else {
		goto L446
	}
L446:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L26
	} else {
		goto L447
	}
L447:
	;
	v1719 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36)+1)))
	F_sanitize_char_2(m, v1719)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L26
	} else {
		goto L448
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(4414000)
	F_errdetail(m, int32(665436), v17+int32(48))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L26
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(497038), int32(1004), int32(403802))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L26
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L26
	} else {
		goto L452
	}
L452:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L26
	} else {
		goto L453
	}
L453:
	;
	F_errdetail(m, int32(654800), int32(0))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L26
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(497038), int32(1018), int32(403802))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L26
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L26
	} else {
		goto L457
	}
L457:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L26
	} else {
		goto L458
	}
L458:
	;
	v1765 = int32(*(*int8)(unsafe.Add(mBase, uint32(v36)+1)))
	F_sanitize_char_2(m, v1765)
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L26
	} else {
		goto L459
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = int32(4414000)
	F_errdetail(m, int32(665436), v17-int32(-64))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L26
	} else {
		goto L460
	}
L460:
	;
	F_errfinish(m, int32(497038), int32(1034), int32(403802))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L26
	} else {
		goto L461
	}
L461:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L462:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L26
	} else {
		goto L463
	}
L463:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L26
	} else {
		goto L464
	}
L464:
	;
	F_errdetail(m, int32(654688), int32(0))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L26
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(497038), int32(1047), int32(403802))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L26
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L26
	} else {
		goto L468
	}
L468:
	;
	F_errmsg(m, int32(441621), int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L26
	} else {
		goto L469
	}
L469:
	;
	F_errfinish(m, int32(497038), int32(1075), int32(403802))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L26
	} else {
		goto L470
	}
L470:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L471:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L26
	} else {
		goto L472
	}
L472:
	;
	F_errmsg(m, int32(271463), int32(0))
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L26
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(497038), int32(1096), int32(403802))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L26
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L26
	} else {
		goto L476
	}
L476:
	;
	F_errmsg(m, int32(414432), int32(0))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L26
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(497038), int32(1110), int32(403802))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L26
	} else {
		goto L478
	}
L478:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L26
	} else {
		goto L480
	}
L480:
	;
	F_errmsg(m, int32(414370), int32(0))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L26
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(497038), int32(1240), int32(403828))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L26
	} else {
		goto L482
	}
L482:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L483:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L26
	} else {
		goto L484
	}
L484:
	;
	F_errmsg(m, int32(414402), int32(0))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L26
	} else {
		goto L485
	}
L485:
	;
	F_errfinish(m, int32(497038), int32(1250), int32(403828))
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L26
	} else {
		goto L486
	}
L486:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L487:
	;
	F_errmsg_internal(m, int32(430854), int32(0))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L26
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(497038), int32(1359), int32(403951))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L26
	} else {
		goto L489
	}
L489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L490:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L26
	} else {
		goto L491
	}
L491:
	;
	F_errmsg(m, int32(404062), int32(0))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L26
	} else {
		goto L492
	}
L492:
	;
	F_errfinish(m, int32(497038), int32(1374), int32(403951))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L26
	} else {
		goto L493
	}
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L26
	} else {
		goto L495
	}
L495:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L26
	} else {
		goto L496
	}
L496:
	;
	F_errdetail(m, int32(641927), int32(0))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L26
	} else {
		goto L497
	}
L497:
	;
	F_errfinish(m, int32(497038), int32(1393), int32(403951))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L26
	} else {
		goto L498
	}
L498:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L499:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L26
	} else {
		goto L500
	}
L500:
	;
	F_errmsg(m, int32(404946), int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L26
	} else {
		goto L501
	}
L501:
	;
	F_errdetail(m, int32(641968), int32(0))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L26
	} else {
		goto L502
	}
L502:
	;
	F_errfinish(m, int32(497038), int32(1401), int32(403951))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L26
	} else {
		goto L503
	}
L503:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L504:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L26
	} else {
		goto L505
	}
L505:
	;
	F_errmsg(m, int32(361116), int32(0))
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L26
	} else {
		goto L506
	}
L506:
	;
	F_errdetail(m, int32(624747), int32(0))
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L26
	} else {
		goto L507
	}
L507:
	;
	F_errfinish(m, int32(497038), int32(421), int32(402644))
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L26
	} else {
		goto L508
	}
L508:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L509:
	;
	if v1011 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v2007
	F_errmsg_internal(m, int32(203043), v17+int32(96))
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L26
	} else {
		goto L523
	}
L511:
	;
	v2007 = int32(13904)
	goto L510
L512:
	;
	goto L513
L513:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+12))
	if v1992 != 0 {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2004 = v1992
	goto L516
L515:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+8))
	if v1996 == int32(2) {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v2007 = v2004
	goto L510
L517:
	;
	v1999 = int32(212308)
	goto L519
L518:
	;
	v1999 = int32(130145)
	goto L519
L519:
	;
	if v1996 == int32(1) {
		goto L520
	} else {
		goto L521
	}
L520:
	;
	v2002 = int32(13904)
	goto L522
L521:
	;
	v2002 = v1999
	goto L522
L522:
	;
	v2004 = v2002
	goto L516
L523:
	;
	F_errfinish(m, int32(497038), int32(1177), int32(338312))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L26
	} else {
		goto L524
	}
L524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L525:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v2023
	F_errmsg_internal(m, int32(199030), v17+int32(112))
	mBase = m.M
	v2029 = m.ExcPending
	if v2029 != 0 {
		goto L26
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(497038), int32(1189), int32(338312))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L26
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	if v1295 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v2059
	F_errmsg_internal(m, int32(203084), v17+int32(128))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L26
	} else {
		goto L542
	}
L530:
	;
	v2059 = int32(13904)
	goto L529
L531:
	;
	goto L532
L532:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+12))
	if v2044 != 0 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2056 = v2044
	goto L535
L534:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+8))
	if v2048 == int32(2) {
		goto L536
	} else {
		goto L537
	}
L535:
	;
	v2059 = v2056
	goto L529
L536:
	;
	v2051 = int32(212308)
	goto L538
L537:
	;
	v2051 = int32(130145)
	goto L538
L538:
	;
	if v2048 == int32(1) {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v2054 = int32(13904)
	goto L541
L540:
	;
	v2054 = v2051
	goto L541
L541:
	;
	v2056 = v2054
	goto L535
L542:
	;
	F_errfinish(m, int32(497038), int32(1435), int32(403977))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L26
	} else {
		goto L543
	}
L543:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L544:
	;
	F_errmsg_internal(m, int32(362757), int32(0))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L26
	} else {
		goto L545
	}
L545:
	;
	F_errfinish(m, int32(497038), int32(1447), int32(403977))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L26
	} else {
		goto L546
	}
L546:
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
	F_appendStringInfoString(m, l1, int32(556636))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v17
	v19 = F_geterrposition(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v287 = int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v292 = F_strlen(m, v288)
	mBase = m.M
	v299 = v292 + v287
	goto L79
L2:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v279 = v274
	v286 = base.B2i32(int32(0) <= v274)
	goto L1
L3:
	;
	return
L4:
	;
	if int32(0) < v19 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.Ui32(v15) <= base.Ui32(v19) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v177 = int32(0)
	if v15 < v177 {
		v279 = v15
		v286 = v177
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v85 = v12 + int32(44)
	v87 = v12 + int32(40)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v92 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	if v17 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v28 = F_strlen(m, v14)
	mBase = m.M
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v29
	if v28 <= v29 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	if v19 <= v15+v17 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v40 = int32(0)
	v45 = int32(0)
	goto L15
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v14))))
	if v48 != int32(59) {
		v69 = v40
		v71 = v45
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L8
L17:
	;
	v73 = v69 + int32(1)
	if v73 < v28 {
		v40 = v73
		v45 = v71
		goto L15
	} else {
		goto L27
	}
L18:
	;
	v52 = v40 + int32(1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v52))))
	if v54 == int32(13) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v57 = v52
	goto L21
L20:
	;
	v57 = v40
	goto L21
L21:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(1)+v57))))
	if v59 != int32(10) {
		v69 = v57
		v71 = v45
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v63 = v57 + int32(2)
	if v63 < v19 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v63
	v69 = v57
	v71 = v63
	goto L17
L24:
	;
	goto L25
L25:
	;
	if v63 <= v19 {
		v69 = v57
		v71 = v45
		goto L17
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v63 - v45
	goto L8
L27:
	;
	goto L16
L28:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v161 = F_errposition(m, int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L46
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v152
	goto L28
L30:
	;
	v110 = v106
	v113 = v107
	v114 = v108
	goto L38
L31:
	;
	v103 = F_strlen(m, v100)
	mBase = m.M
	if v103 <= int32(0) {
		v149 = v100
		v152 = v103
		v153 = v102
		goto L29
	} else {
		goto L36
	}
L32:
	;
	v100 = v14
	v102 = int32(0)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v96 = v14 + v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if int32(0) < v97 {
		v106 = v96
		v107 = v97
		v108 = v92
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v100 = v96
	v102 = v92
	goto L31
L36:
	;
	v106 = v100
	v107 = v103
	v108 = v102
	goto L30
L37:
	;
	v135 = v113
	goto L42
L38:
	;
	v117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v110))))
	v118 = F_scanner_isspace(m, v117)
	mBase = m.M
	if v118 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v149 = v124
	v152 = int32(0)
	v153 = v107 + v108
	goto L29
L40:
	;
	v121 = int32(1)
	v124 = v110 + v121
	if v121 < v113 {
		v110 = v124
		v113 = v113 - v121
		v114 = v114 + v121
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v140 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135+(v110-int32(1))))))
	v141 = F_scanner_isspace(m, v140)
	mBase = m.M
	if v141 == int32(0) {
		v149 = v110
		v152 = v135
		v153 = v114
		goto L29
	} else {
		goto L44
	}
L43:
	;
	v149 = v110
	v152 = int32(0)
	v153 = v114
	goto L29
L44:
	;
	v144 = int32(1)
	if v144 < v135 {
		v135 = v135 - v144
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v163 = v19 - v158
	if v163 < v159 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v165 = v163
	goto L49
L48:
	;
	v165 = v159
	goto L49
L49:
	;
	v166 = int32(0)
	if v166 <= v163 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v169 = v165
	goto L52
L51:
	;
	v169 = v166
	goto L52
L52:
	;
	F_internalerrposition(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v173 = F_pnstrdup(m, v149, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v175 = F_internalerrquery(m, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	goto L2
L56:
	;
	v181 = v12 + int32(44)
	v183 = v12 + int32(40)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v188 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L75
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v248
	goto L57
L59:
	;
	v206 = v202
	v209 = v203
	v210 = v204
	goto L67
L60:
	;
	v199 = F_strlen(m, v196)
	mBase = m.M
	if v199 <= int32(0) {
		v245 = v196
		v248 = v199
		v249 = v198
		goto L58
	} else {
		goto L65
	}
L61:
	;
	v196 = v14
	v198 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v192 = v14 + v188
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if int32(0) < v193 {
		v202 = v192
		v203 = v193
		v204 = v188
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v196 = v192
	v198 = v188
	goto L60
L65:
	;
	v202 = v196
	v203 = v199
	v204 = v198
	goto L59
L66:
	;
	v231 = v209
	goto L71
L67:
	;
	v213 = int32(*(*int8)(unsafe.Add(mBase, uint32(v206))))
	v214 = F_scanner_isspace(m, v213)
	mBase = m.M
	if v214 == int32(0) {
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v245 = v220
	v248 = int32(0)
	v249 = v203 + v204
	goto L58
L69:
	;
	v217 = int32(1)
	v220 = v206 + v217
	if v217 < v209 {
		v206 = v220
		v209 = v209 - v217
		v210 = v210 + v217
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v236 = int32(*(*int8)(unsafe.Add(mBase, uint32(v231+(v206-int32(1))))))
	v237 = F_scanner_isspace(m, v236)
	mBase = m.M
	if v237 == int32(0) {
		v245 = v206
		v248 = v231
		v249 = v210
		goto L58
	} else {
		goto L73
	}
L72:
	;
	v245 = v206
	v248 = int32(0)
	v249 = v210
	goto L58
L73:
	;
	v240 = int32(1)
	if v240 < v231 {
		v231 = v231 - v240
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v245
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v258
	F_errcontext_msg(m, int32(689861), v12+int32(32))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	goto L2
L77:
	;
	if v311 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	goto L77
L79:
	;
	v301 = int32(0)
	if v299 == v301 {
		v311 = v301
		goto L78
	} else {
		goto L81
	}
L80:
	;
	v311 = v306
	goto L78
L81:
	;
	v305 = v299 - int32(1)
	v306 = v288 + v305
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v307 != int32(47) {
		v299 = v305
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v314 = v311 + int32(1)
	goto L85
L84:
	;
	v314 = v288
	goto L85
L85:
	;
	if v286 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	m.G0 = v12 + int32(48)
	return
L87:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v316 == int32(0) {
		v344 = v287
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L98
	}
L90:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L3
	} else {
		goto L96
	}
L91:
	;
	v321 = v279
	v322 = v315
	v323 = v287
	goto L92
L92:
	;
	v329 = v321 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v329
	if v329 < int32(0) {
		v344 = v323
		goto L90
	} else {
		goto L94
	}
L93:
	;
	v344 = v336
	goto L90
L94:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v336 = v323 + base.B2i32(v333 == int32(10))
	v338 = v322 + int32(1)
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v339 != 0 {
		v321 = v329
		v322 = v338
		v323 = v336
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v314
	F_errcontext_msg(m, int32(476726), v12)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	goto L86
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v314
	F_errcontext_msg(m, int32(715321), v12+int32(16))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	goto L86
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
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
	if v15 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	F_ProcessClientReadInterrupt(m, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	goto L8
L6:
	;
	v66 = v15
	goto L7
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+512))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	if base.Ui32(l2) < base.Ui32(v66) {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_pgl_recv(m, v23, l1, l2, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v66 = v60
	goto L7
L10:
	;
	if int32(0) <= v25 {
		v84 = v25
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v29 != 0 {
		v84 = v25
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v31 != int32(6) {
		v84 = v25
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v36 = int32(0)
	F_ModifyWaitEvent(m, v35, v36, int32(2), v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[366]))
	v46 = F_WaitEventSetWait(m, v42, int32(-1), v8, int32(1), int32(100663296))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v48&int32(16) != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	if v48&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if v60 <= int32(0) {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	F_ProcessClientReadInterrupt(m, int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L9
L23:
	;
	v72 = l2
	goto L25
L24:
	;
	v72 = v66
	goto L25
L25:
	;
	if v72 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v75 + v72
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v78 - v72
	v84 = v72
	goto L4
L27:
	;
	v73 = F__emscripten_memcpy_bulkmem(m, l1, v68+v69, v72)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	m.G0 = v8 + int32(16)
	return v84
L31:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(99431), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(498577), int32(241), int32(464040))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
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
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v143 = F_cstring_to_text_with_len(m, v141, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	F_appendStringInfo(m, v10+int32(16), int32(744924), v10)
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
	if base.Ui32(v47-int32(465)) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v116 = v33 + int32(4)
	if v116 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	F_appendStringInfoString(m, v10+int32(16), v35)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v56 = int32(92)
	v57 = F___strchrnul(m, v35, v56)
	mBase = m.M
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v59 == v56 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L11
L16:
	;
	if v63 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v63 = v57
	goto L19
L18:
	;
	v63 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(69))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
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
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v77 = v35
	goto L25
L25:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(39))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	v83 = v81 & int32(255)
	if base.B2i32(v83 != int32(92))&base.B2i32(v83 != int32(39)) == int32(0) {
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
	F_appendStringInfoChar(m, v10+int32(16), base.I32_extend8_s(v81))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_appendStringInfoChar(m, v10+int32(16), base.I32_extend8_s(v81))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v77 = v77 + int32(1)
	goto L25
L35:
	;
	goto L11
L36:
	;
	v131 = v27 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v131 < v132 {
		v27 = v131
		goto L6
	} else {
		goto L40
	}
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v119+v120<<(uint(int32(2))%32)) <= base.Ui32(v116) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_appendStringInfoString(m, v10+int32(16), int32(745866))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	goto L7
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	m.G0 = v10 + int32(32)
	return v143
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
	var v44 int32
	_ = v44
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
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v44 <= v43 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v50 = v43
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
	if v105 != 0 {
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
	v110 = F_query_tree_walker_impl(m, l0, int32(1039), v9+int32(12), int32(3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v4 == int32(67) {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_setRuleCheckAsUser_Query(m, l0, v7)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(1039), l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = v15
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
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
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L51
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L47
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
	F_CheckTableNotInUse(m, l1, int32(540573))
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
	v211 = m.ExcPending
	if v211 != 0 {
		goto L7
	} else {
		goto L46
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
	if l0 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L43
	}
L19:
	;
	if l3 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v68 = F_palloc0(m, int32(20))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(52)
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+74)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+68))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
	v78 = F_makeVar(m, int32(1), v73, v74, v75, v76, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = int32(-1)
	v82 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)) = uint8(v82)
	v84 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v68
	v92 = F_list_make1_impl(m, v84, v13+int32(24))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v95 = F_ConstraintImpliedByRelConstraint(m, l1, v92, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v99 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v121 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L28:
	;
	if v99 == int32(0) {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v104 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v56 + v104
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v103 + v104
	F_errmsg_internal(m, int32(152604), v13+int32(16))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(494319), int32(8113), int32(120239))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	goto L18
L32:
	;
	v191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v184)+76)) = uint8(v191)
	goto L18
L33:
	;
	v159 = F_palloc0(m, int32(140))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L40
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v124 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v131 = int32(0)
	goto L36
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v127+v131<<(uint(int32(2))%32))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v143 == v120 {
		v184 = v142
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	v146 = v131 + int32(1)
	if v124 != v146 {
		v131 = v146
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v120
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+4)) = uint8(v165)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v168 = F_CreateTupleDescCopyConstr(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+8)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v159)+88)) = int64(0)
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+84)) = uint8(v173)
	v175 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v159)+96)) = uint16(v175)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v178 = F_lappend(m, v177, v159)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v178
	v184 = v159
	goto L32
L43:
	;
	F_sequence_close(m, v41, int32(3))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	F_pfree(m, v44)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L9
L46:
	;
	goto L9
L47:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(143854), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(494319), int32(4460), int32(409582))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg_internal(m, int32(46437), v13)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(494319), int32(7868), int32(301833))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
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
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 float64
	_ = v245
	var v246 float64
	_ = v246
	var v249 float64
	_ = v249
	var v250 float64
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 <= int32(0) {
		v334 = v2
		v338 = v2
		v339 = v2
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L144
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v338
	return
L5:
	;
	v19 = v2
	v20 = v2
	v23 = v2
	v24 = v2
	v26 = v2
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v26<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v316 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L8:
	;
	v325 = v26 + int32(1)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v325 < v326 {
		v19 = v316
		v20 = v317
		v23 = v320
		v24 = v321
		v26 = v325
		goto L6
	} else {
		goto L139
	}
L9:
	;
	v35 = F_lappend(m, v24, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if v19 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L12:
	;
	return
L13:
	;
	if v19 != 0 {
		v316 = v19
		v317 = v20
		v320 = v23
		v321 = v35
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v20 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v316 = int32(0)
	v317 = v33
	v320 = v23
	v321 = v35
	goto L8
L16:
	;
	goto L17
L17:
	;
	v40 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v44 = v43
	goto L20
L19:
	;
	v44 = v40
	goto L20
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v45 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = v46
	goto L23
L22:
	;
	v47 = v40
	goto L23
L23:
	;
	v48 = int32(0)
	if v44 == v48 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v316 = v48
	v317 = v33
	v320 = v23
	v321 = v35
	goto L8
L25:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v144 != v145 {
		goto L62
	} else {
		goto L63
	}
L26:
	;
	switch v143 {
	case 0:
		goto L25
	case 1:
		goto L24
	default:
		v316 = v48
		v317 = v20
		v320 = v23
		v321 = v35
		goto L8
	}
L27:
	;
	v143 = base.B2i32(v47 != int32(0))
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v47 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v143 = int32(2)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v65 < v66 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v68 = v65
	goto L35
L34:
	;
	v68 = v66
	goto L35
L35:
	;
	if v68 <= int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v71 = int32(1)
	goto L38
L37:
	;
	v71 = v68
	goto L38
L38:
	;
	v72 = int32(8)
	v76 = int32(0)
	v78 = v76
	v79 = v76
	goto L41
L39:
	;
	v143 = int32(3)
	goto L26
L40:
	;
	v143 = v129
	goto L26
L41:
	;
	v89 = v79 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v44+v72+v89)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+(v47+v72))))
	if v91&(v93^int32(-1)) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	if v66 < v65 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v115 = v79 + int32(1)
	if v115 != v71 {
		v78 = v112
		v79 = v115
		goto L41
	} else {
		goto L51
	}
L44:
	;
	v99 = int32(3)
	if v78 == int32(1) {
		v129 = v99
		goto L40
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v93&(v91^int32(-1)) == int32(0) {
		v112 = v78
		goto L43
	} else {
		goto L49
	}
L47:
	;
	if v93&(v91^int32(-1)) != 0 {
		v129 = v99
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v112 = int32(2)
	goto L43
L49:
	;
	if v78 == int32(2) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v112 = int32(1)
	goto L43
L51:
	;
	goto L42
L52:
	;
	if v112 == int32(1) {
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
		v129 = v112
		goto L40
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
	v143 = v122
	goto L26
L58:
	;
	if v112 == int32(2) {
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
	goto L40
L62:
	;
	if v145 <= v144 {
		v316 = v48
		v317 = v20
		v320 = v23
		v321 = v35
		goto L8
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v20)+56))
	if base.F64_lt(v148, v149) != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v316 = v48
	v317 = v33
	v320 = v23
	v321 = v35
	goto L8
L66:
	;
	v316 = v48
	v317 = v33
	v320 = v23
	v321 = v35
	goto L8
L67:
	;
	goto L68
L68:
	;
	if base.F64_gt(v148, v149) != 0 {
		v316 = v48
		v317 = v20
		v320 = v23
		v321 = v35
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v20)+48))
	if base.F64_lt(v152, v153) == int32(0) {
		v316 = v48
		v317 = v20
		v320 = v23
		v321 = v35
		goto L8
	} else {
		goto L70
	}
L70:
	;
	goto L24
L71:
	;
	v316 = v33
	v317 = v20
	v320 = v33
	v321 = v24
	goto L8
L72:
	;
	goto L73
L73:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	if v163 != v164 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	if v241 != v242 {
		goto L108
	} else {
		goto L109
	}
L75:
	;
	v238 = v33
	goto L74
L76:
	;
	if v164 <= v163 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v23)+48))
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	if base.F64_lt(v167, v168) != 0 {
		v238 = v23
		goto L74
	} else {
		goto L80
	}
L79:
	;
	v238 = v23
	goto L74
L80:
	;
	if base.F64_gt(v167, v168) != 0 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v23)+56))
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	if base.F64_lt(v171, v172) != 0 {
		v238 = v23
		goto L74
	} else {
		goto L82
	}
L82:
	;
	if base.F64_gt(v171, v172) != 0 {
		goto L75
	} else {
		goto L83
	}
L83:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	if v175 == v176 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v233 != int32(2) {
		v238 = v23
		goto L74
	} else {
		goto L106
	}
L85:
	;
	v233 = int32(0)
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
	if v223 != 0 {
		goto L103
	} else {
		goto L104
	}
L89:
	;
	v217 = int32(0)
	v223 = base.B2i32(v197 == v217)
	v225 = base.B2i32(v206 != v217) << (uint(int32(1)) % 32)
	goto L88
L90:
	;
	v187 = int32(0)
	if v175 == v187 {
		v197 = v187
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v233 = int32(3)
	goto L84
L92:
	;
	if v176 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v191 <= v184 {
		v197 = int32(0)
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v197 = v193 + v184<<(uint(int32(2))%32)
	goto L92
L95:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v206 = v203 + v184<<(uint(int32(2))%32)
	if v197 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L96:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v184 < v198 {
		goto L95
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v200 = int32(0)
	v223 = base.B2i32(v197 == v200)
	v225 = v200
	goto L88
L99:
	;
	goto L98
L100:
	;
	if v206 == int32(0) {
		goto L89
	} else {
		goto L101
	}
L101:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v213 == v214 {
		v184 = v184 + int32(1)
		goto L90
	} else {
		goto L102
	}
L102:
	;
	goto L91
L103:
	;
	v227 = v225
	goto L105
L104:
	;
	v227 = int32(1)
	goto L105
L105:
	;
	v233 = v227
	goto L84
L106:
	;
	goto L75
L107:
	;
	v316 = v33
	v317 = v20
	v320 = v238
	v321 = v24
	goto L8
L108:
	;
	if v242 <= v241 {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v19)+56))
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	if base.F64_lt(v245, v246) != 0 {
		v316 = v19
		v317 = v20
		v320 = v238
		v321 = v24
		goto L8
	} else {
		goto L112
	}
L111:
	;
	v316 = v19
	v317 = v20
	v320 = v238
	v321 = v24
	goto L8
L112:
	;
	if base.F64_gt(v245, v246) != 0 {
		goto L107
	} else {
		goto L113
	}
L113:
	;
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v19)+48))
	v250 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	if base.F64_lt(v249, v250) != 0 {
		v316 = v19
		v317 = v20
		v320 = v238
		v321 = v24
		goto L8
	} else {
		goto L114
	}
L114:
	;
	if base.F64_gt(v249, v250) != 0 {
		goto L107
	} else {
		goto L115
	}
L115:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	if v253 == v254 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v311 != int32(2) {
		v316 = v19
		v317 = v20
		v320 = v238
		v321 = v24
		goto L8
	} else {
		goto L138
	}
L117:
	;
	v311 = int32(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v262 = int32(0)
	goto L122
L120:
	;
	if v301 != 0 {
		goto L135
	} else {
		goto L136
	}
L121:
	;
	v295 = int32(0)
	v301 = base.B2i32(v275 == v295)
	v303 = base.B2i32(v284 != v295) << (uint(int32(1)) % 32)
	goto L120
L122:
	;
	v265 = int32(0)
	if v253 == v265 {
		v275 = v265
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v311 = int32(3)
	goto L116
L124:
	;
	if v254 != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v269 <= v262 {
		v275 = int32(0)
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v275 = v271 + v262<<(uint(int32(2))%32)
	goto L124
L127:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	v284 = v281 + v262<<(uint(int32(2))%32)
	if v275 == int32(0) {
		goto L121
	} else {
		goto L132
	}
L128:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	if v262 < v276 {
		goto L127
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v278 = int32(0)
	v301 = base.B2i32(v275 == v278)
	v303 = v278
	goto L120
L131:
	;
	goto L130
L132:
	;
	if v284 == int32(0) {
		goto L121
	} else {
		goto L133
	}
L133:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if v291 == v292 {
		v262 = v262 + int32(1)
		goto L122
	} else {
		goto L134
	}
L134:
	;
	goto L123
L135:
	;
	v305 = v303
	goto L137
L136:
	;
	v305 = int32(1)
	goto L137
L137:
	;
	v311 = v305
	goto L116
L138:
	;
	goto L107
L139:
	;
	goto L7
L140:
	;
	v334 = v317
	v338 = v320
	v339 = v321
	goto L4
L141:
	;
	goto L142
L142:
	;
	v330 = F_lcons(m, v316, v321)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L12
	} else {
		goto L143
	}
L143:
	;
	v334 = v316
	v338 = v320
	v339 = v330
	goto L4
L144:
	;
	F_errmsg_internal(m, int32(16446), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(499267), int32(283), int32(77335))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, _consts[901]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[901])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(453809), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(497964), int32(1418), int32(277586))
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
			v28 = l0
		} else {
			v28 = int32(556035)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_consts[910]))) = v28
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_getTypeOutputInfo(m, int32(700), v7+int32(12), v7+int32(11))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v23 = F_OidOutputFunctionCall(m, v22, v9)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			F_SetConfigOption(m, int32(429399), v23, int32(6), int32(13))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v31 = *(*float64)(unsafe.Add(mBase, _consts[1099]))
				m.G0 = v7 + int32(16)
				return base.I32_reinterpret_f32(base.F32_demote_f64(v31))
			}
		}
	}
}
func F_set_spins_per_delay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[573])) = l0
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v273 int64
	_ = v273
	var v276 int64
	_ = v276
	var v279 int64
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v464 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v464
L2:
	;
	if l0 == int32(6) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v422 = int32(0)
	v423 = int32(4680576)
	v428 = v3
	goto L155
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
		goto L98
	} else {
		goto L99
	}
L7:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _consts[1140]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v20
	v23 = *(*int64)(unsafe.Add(mBase, _consts[1141]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	v26 = *(*int64)(unsafe.Add(mBase, _consts[1142]))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v26
	v29 = int32(0)
	v30 = l1
	goto L9
L8:
	;
	v464 = int32(0)
	goto L1
L9:
	;
	goto L15
L10:
	;
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int64)(unsafe.Add(mBase, _consts[1143])) = v273
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int64)(unsafe.Add(mBase, _consts[1144])) = v276
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*int64)(unsafe.Add(mBase, _consts[1145])) = v279
	goto L3
L11:
	;
	v134 = v123 - v30
	if v134 <= int32(23) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	goto L11
L13:
	;
	v113 = v108
	goto L31
L14:
	;
	v108 = v100
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
	v47 = v30
	goto L21
L19:
	;
	v60 = v30
	goto L20
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v69 = int32(-2139062144)
	if (int32(16843008)-v66|v66)&v69 != v69 {
		v100 = v60
		goto L14
	} else {
		goto L26
	}
L21:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v52 == int32(0) {
		v123 = v47
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v60 = v57
	goto L20
L23:
	;
	if int32(59) == v52 {
		v123 = v47
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v57 = v47 + int32(1)
	if v57&int32(3) != 0 {
		v47 = v57
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v75 = v60
	v77 = v66
	goto L27
L27:
	;
	v81 = v77 ^ int32(993737531)
	v84 = int32(-2139062144)
	if (int32(16843008)-v81|v81)&v84 != v84 {
		v100 = v75
		goto L14
	} else {
		goto L29
	}
L28:
	;
	v108 = v90
	goto L13
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v90 = v75 + int32(4)
	v94 = int32(-2139062144)
	if (v88|(int32(16843008)-v88))&v94 == v94 {
		v75 = v90
		v77 = v88
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v115 == int32(0) {
		v123 = v113
		goto L12
	} else {
		goto L33
	}
L32:
	;
	v123 = v113
	goto L12
L33:
	;
	if v115 != int32(59) {
		v113 = v113 + int32(1)
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v134 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v146 = v30
	goto L37
L37:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v150 != 0 {
		v164 = v11
		goto L46
	} else {
		goto L47
	}
L38:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v134))) = uint8(v140)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v144 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v137 = F__emscripten_memcpy_bulkmem(m, v11, v30, v134)
	mBase = m.M
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v145 = v123 + int32(1)
	goto L44
L43:
	;
	v145 = v30
	goto L44
L44:
	;
	v146 = v145
	goto L37
L45:
	;
	if v259 == int32(-1) {
		goto L8
	} else {
		goto L95
	}
L46:
	;
	v167 = int32(0)
	goto L61
L47:
	;
	v152 = F_getenv(m, int32(533324))
	mBase = m.M
	if v152 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v153 != 0 {
		v164 = v152
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v158 = F_getenv(m, v29*int32(12)+int32(4097008))
	mBase = m.M
	if v158 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 != 0 {
		v164 = v158
		goto L46
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v161 = F_getenv(m, int32(536505))
	mBase = m.M
	if v161 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v162 != 0 {
		v164 = v161
		goto L46
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v164 = int32(556070)
	goto L46
L59:
	;
	goto L58
L60:
	;
	v183 = int32(556070)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v184 == int32(46) {
		v191 = v183
		goto L71
	} else {
		goto L72
	}
L61:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v167))))
	if v171 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v182 = v167
	goto L60
L63:
	;
	goto L62
L64:
	;
	if v171 == int32(47) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v176 = int32(23)
	v178 = v167 + int32(1)
	if v178 != v176 {
		v167 = v178
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v182 = v176
	goto L60
L67:
	;
	v259 = v251
	goto L45
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	if v212 != 0 {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	if v29 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	v197 = F_strcmp(m, v195, int32(556070))
	mBase = m.M
	if v197 == int32(0) {
		v202 = v195
		goto L69
	} else {
		goto L76
	}
L71:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v192 == int32(0) {
		v202 = v191
		goto L69
	} else {
		goto L75
	}
L72:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v182))))
	if v188 != 0 {
		v191 = v183
		goto L71
	} else {
		goto L73
	}
L73:
	;
	if v184 != int32(67) {
		v195 = v164
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v191 = v164
	goto L71
L75:
	;
	v195 = v191
	goto L70
L76:
	;
	v201 = F_strcmp(m, v195, int32(509771))
	mBase = m.M
	if v201 != 0 {
		goto L68
	} else {
		goto L77
	}
L77:
	;
	v202 = v195
	goto L69
L78:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	if v206 == int32(46) {
		v251 = int32(4096916)
		goto L67
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v259 = int32(0)
	goto L45
L81:
	;
	goto L80
L82:
	;
	v215 = v212
	goto L85
L83:
	;
	goto L84
L84:
	;
	v230 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v230 != 0 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v220 = F_strcmp(m, v195, v215+int32(8))
	mBase = m.M
	if v220 == int32(0) {
		v251 = v215
		goto L67
	} else {
		goto L87
	}
L86:
	;
	goto L84
L87:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	if v223 != 0 {
		v215 = v223
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v232 = *(*int64)(unsafe.Add(mBase, _consts[1135]))
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v232
	v235 = v230 + int32(8)
	v236 = F___memcpy(m, v235, v195, v182)
	mBase = m.M
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v182))) = uint8(v238)
	v240 = int32(4680032)
	v241 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	*(*int32)(unsafe.Add(mBase, uint32(v230)+32)) = v241
	*(*int32)(unsafe.Add(mBase, _consts[1134])) = v230
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v29|v230 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v248 = v230
	goto L94
L93:
	;
	v248 = int32(4096916)
	goto L94
L94:
	;
	v251 = v248
	goto L67
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)+v29<<(uint(int32(2))%32)))) = v259
	v269 = v29 + int32(1)
	if v269 != int32(6) {
		v29 = v269
		v30 = v146
		goto L9
	} else {
		goto L96
	}
L96:
	;
	goto L10
L97:
	;
	if v407 != 0 {
		goto L152
	} else {
		goto L153
	}
L98:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v285 != 0 {
		v299 = l1
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1143])))
	v407 = v406
	goto L97
L101:
	;
	if v394 == int32(-1) {
		v464 = v3
		goto L1
	} else {
		goto L151
	}
L102:
	;
	v302 = int32(0)
	goto L117
L103:
	;
	v287 = F_getenv(m, int32(533324))
	mBase = m.M
	if v287 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v288 != 0 {
		v299 = v287
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v293 = F_getenv(m, l0*int32(12)+int32(4097008))
	mBase = m.M
	if v293 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L106
L108:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v294 != 0 {
		v299 = v293
		goto L102
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v296 = F_getenv(m, int32(536505))
	mBase = m.M
	if v296 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v297 != 0 {
		v299 = v296
		goto L102
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v299 = int32(556070)
	goto L102
L115:
	;
	goto L114
L116:
	;
	v318 = int32(556070)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v319 == int32(46) {
		v326 = v318
		goto L127
	} else {
		goto L128
	}
L117:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v302))))
	if v306 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v317 = v302
	goto L116
L119:
	;
	goto L118
L120:
	;
	if v306 == int32(47) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v311 = int32(23)
	v313 = v302 + int32(1)
	if v313 != v311 {
		v302 = v313
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v317 = v311
	goto L116
L123:
	;
	v394 = v386
	goto L101
L124:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	if v347 != 0 {
		goto L138
	} else {
		goto L139
	}
L125:
	;
	if l0 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L126:
	;
	v332 = F_strcmp(m, v330, int32(556070))
	mBase = m.M
	if v332 == int32(0) {
		v337 = v330
		goto L125
	} else {
		goto L132
	}
L127:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+1)))
	if v327 == int32(0) {
		v337 = v326
		goto L125
	} else {
		goto L131
	}
L128:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v317))))
	if v323 != 0 {
		v326 = v318
		goto L127
	} else {
		goto L129
	}
L129:
	;
	if v319 != int32(67) {
		v330 = v299
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v326 = v299
	goto L127
L131:
	;
	v330 = v326
	goto L126
L132:
	;
	v336 = F_strcmp(m, v330, int32(509771))
	mBase = m.M
	if v336 != 0 {
		goto L124
	} else {
		goto L133
	}
L133:
	;
	v337 = v330
	goto L125
L134:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	if v341 == int32(46) {
		v386 = int32(4096916)
		goto L123
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v394 = int32(0)
	goto L101
L137:
	;
	goto L136
L138:
	;
	v350 = v347
	goto L141
L139:
	;
	goto L140
L140:
	;
	v365 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v365 != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	v355 = F_strcmp(m, v330, v350+int32(8))
	mBase = m.M
	if v355 == int32(0) {
		v386 = v350
		goto L123
	} else {
		goto L143
	}
L142:
	;
	goto L140
L143:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v350)+32))
	if v358 != 0 {
		v350 = v358
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v367 = *(*int64)(unsafe.Add(mBase, _consts[1135]))
	*(*int64)(unsafe.Add(mBase, uint32(v365))) = v367
	v370 = v365 + int32(8)
	v371 = F___memcpy(m, v370, v330, v317)
	mBase = m.M
	v373 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v370+v317))) = uint8(v373)
	v375 = int32(4680032)
	v376 = *(*int32)(unsafe.Add(mBase, _consts[1134]))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+32)) = v376
	*(*int32)(unsafe.Add(mBase, _consts[1134])) = v365
	goto L147
L146:
	;
	goto L147
L147:
	;
	if l0|v365 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v383 = v365
	goto L150
L149:
	;
	v383 = int32(4096916)
	goto L150
L150:
	;
	v386 = v383
	goto L123
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1143]))) = v394
	v407 = v394
	goto L97
L152:
	;
	v411 = v407 + int32(8)
	goto L154
L153:
	;
	v411 = int32(544456)
	goto L154
L154:
	;
	v464 = v411
	goto L1
L155:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _consts[1143]))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v422<<(uint(int32(2))%32))+uint32(_consts[1143])))
	if v436 != 0 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v455 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v455)
	if v450 != int32(6) {
		goto L165
	} else {
		goto L166
	}
L157:
	;
	v440 = v436 + int32(8)
	goto L159
L158:
	;
	v440 = int32(544456)
	goto L159
L159:
	;
	v441 = F_strlen(m, v440)
	mBase = m.M
	if v441 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v444 = v423 + v441
	v445 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v445)
	v447 = int32(1)
	v450 = v428 + base.B2i32(v436 == v431)
	v452 = v422 + v447
	if v452 != int32(6) {
		v422 = v452
		v423 = v444 + v447
		v428 = v450
		goto L155
	} else {
		goto L164
	}
L161:
	;
	v442 = F__emscripten_memcpy_bulkmem(m, v423, v440, v441)
	mBase = m.M
	goto L163
L162:
	;
	goto L163
L163:
	;
	goto L160
L164:
	;
	goto L156
L165:
	;
	v460 = int32(4680576)
	goto L167
L166:
	;
	v460 = v440
	goto L167
L167:
	;
	v464 = v460
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[166]))
		*(*int32)(unsafe.Add(mBase, _consts[189])) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(427486)
		v20 = F_format_elog_string(m, int32(581643), v5)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[190])) = v20
			m.G0 = v5 + int32(16)
			return base.B2i32(v9 != int32(0))
		}
	} else {
		m.G0 = v5 + int32(16)
		return base.B2i32(v9 != int32(0))
	}
}
func F_shim_read(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	v3 = m.Env.Pgmem_recv(m, l0, l1)
	return v3
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v222 int32
	_ = v222
	var v241 int32
	_ = v241
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = int32(1)
	v22 = v17 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = v22
	goto L5
L4:
	;
	v28 = v17 + int32(4)
	goto L5
L5:
	;
	if v25 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_generate_trgm_only(m, v14+int32(4), v28, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v33 = int32(4)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v35&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v48 = int32(1)
	if v27 != 0 {
		v58 = int32(base.Ui32(v25)>>(uint(v48)%32)) - v48
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v44 = v33
	goto L12
L11:
	;
	v44 = base.B2i32(v35 == int32(18)) << (uint(v33) % 32)
	goto L12
L12:
	;
	if v35 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v33
	goto L15
L14:
	;
	v47 = v44
	goto L15
L15:
	;
	v58 = v47
	goto L6
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)) = uint8(v64)
	if int32(2) <= v62 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = v63 + int32(5)
	F_pg_qsort(m, v69, v62, int32(3), int32(6910))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v127 = v62
	goto L20
L20:
	;
	v131 = v127*int32(12) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v131
	v133 = int32(2)
	v138 = base.I32_div_u_s(int32(base.Ui32(v131)>>(uint(v133)%32))-int32(5), int32(3))
	v143 = F_palloc(m, v138<<(uint(v133)%32)+int32(4))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L29
	}
L21:
	;
	v77 = int32(1)
	v78 = int32(0)
	goto L22
L22:
	;
	v87 = int32(3)
	v89 = v69 + v77*v87
	v94 = *(*int32)(unsafe.Add(mBase, _consts[1098]))
	v95 = m.T0[v94].(func(*base.Module, int32, int32) int32)(m, v89, v69+v78*v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v127 = v109 + int32(1)
	goto L20
L24:
	;
	v112 = v77 + int32(1)
	if v112 != v62 {
		v77 = v112
		v78 = v109
		goto L22
	} else {
		goto L28
	}
L25:
	;
	if v95 == int32(0) {
		v109 = v78
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v100 = v78 + int32(1)
	if v77 == v100 {
		v109 = v77
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v104 = v69 + v100*int32(3)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
	*(*uint16)(unsafe.Add(mBase, uint32(v104))) = uint16(v105)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)) = uint8(v107)
	v109 = v100
	goto L24
L28:
	;
	goto L23
L29:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(int32(3)) <= base.Ui32(int32(base.Ui32(v145)>>(uint(int32(2))%32))-int32(5)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v156 = v63 + int32(5)
	v160 = int32(0)
	goto L33
L31:
	;
	v320 = int32(0)
	goto L32
L32:
	;
	v322 = F_construct_array_builtin(m, v143, v320, int32(25))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L58
	}
L33:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169*int32(28))+uint32(_consts[864])))
	goto L35
L34:
	;
	v320 = v308
	goto L32
L35:
	;
	if int32(4) <= v174 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179*int32(28))+uint32(_consts[864])))
	goto L39
L37:
	;
	v189 = int32(16)
	goto L38
L38:
	;
	v190 = F_palloc(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v189 = v184*int32(3) + int32(4)
	goto L38
L40:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194*int32(28))+uint32(_consts[864])))
	goto L43
L41:
	;
	v294 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v143+v160<<(uint(v294)%32)))) = v190
	v298 = int32(3)
	v301 = v160 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v308 = base.I32_div_u_s(int32(base.Ui32(v302)>>(uint(v294)%32))-int32(5), v298)
	if base.Ui32(v301) < base.Ui32(v308) {
		v156 = v156 + v298
		v160 = v301
		goto L33
	} else {
		goto L57
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = int32(28)
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)) = uint8(v285)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+5)) = uint8(v287)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+6)) = uint8(v289)
	goto L41
L43:
	;
	if v199 < int32(2) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v203 = base.I32_extend8_s(v202)
	if v203 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+2)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v259 | (v260<<(uint(int32(8))%32) | v202<<(uint(int32(16))%32))
	v269 = v190 + int32(4)
	v272 = F_pg_snprintf(m, v269, int32(12), int32(29614), v14)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L56
	}
L46:
	;
	goto L47
L47:
	;
	if base.B2i32(base.B2i32(base.Ui32(v202-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v202|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v203 != int32(32)) != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v222 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v222 < int32(0) {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L50
L50:
	;
	if base.B2i32(base.B2i32(base.Ui32(v222-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v222|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v222 != int32(32)) != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v241 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156)+2)))
	if v241 < int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L53
L53:
	;
	if v241 == int32(32) {
		goto L42
	} else {
		goto L54
	}
L54:
	;
	if base.B2i32(base.Ui32(v241-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v241|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L45
L56:
	;
	v274 = F_strlen(m, v269)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v274<<(uint(int32(2))%32) + int32(16)
	goto L41
L57:
	;
	goto L34
L58:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(int32(3)) <= base.Ui32(int32(base.Ui32(v324)>>(uint(int32(2))%32))-int32(5)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v333 = int32(0)
	goto L62
L60:
	;
	goto L61
L61:
	;
	F_pfree(m, v143)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L66
	}
L62:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v143+v333<<(uint(int32(2))%32))))
	F_pfree(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	v350 = v333 + int32(1)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v357 = base.I32_div_u_s(int32(base.Ui32(v351)>>(uint(int32(2))%32))-int32(5), int32(3))
	if base.Ui32(v350) < base.Ui32(v357) {
		v333 = v350
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	F_pfree(m, v63)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v374 != v17 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_pfree(m, v17)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	m.G0 = v14 + int32(16)
	return v322
L71:
	;
	goto L70
}
func F_sigUsr1Handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[477])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
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
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	v6 = l0
	goto L1
L1:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v12 = v10 - int32(43)
	if base.Ui32(int32(5)) < base.Ui32(v12) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v23 = v6
	v24 = v10
	v26 = int32(1)
	goto L9
L3:
	;
	goto L2
L4:
	;
	if int32(1)<<(uint(v12)%32)&int32(37) == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v6 = v6 + int32(1)
	goto L1
L6:
	;
	return v65
L7:
	;
	v65 = v26
	goto L6
L8:
	;
	v38 = v23
	v39 = v24
	v40 = int32(0)
	goto L15
L9:
	;
	switch v24 - int32(46) {
	case 0:
		v32 = v26
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
	if v24 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v23 = v23 + int32(1)
	v24 = v33
	v26 = v32
	goto L9
L13:
	;
	v32 = v26 + int32(1)
	goto L12
L14:
	;
	goto L8
L15:
	;
	if v39 != int32(46) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v58 != 0 {
		v65 = v58
		goto L6
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	if base.Ui32(int32(9)) < base.Ui32((v39-int32(48))&int32(255)) {
		v58 = v40
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v52 = v40 + base.B2i32(v39 != int32(46))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v53 != 0 {
		v38 = v38 + int32(1)
		v39 = v53
		v40 = v52
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v58 = v52
	goto L17
L23:
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
	var v29 int32
	_ = v29
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
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
	v322 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v322)
	return v312 - v15
L4:
	;
	v311 = v14
	v312 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v14
	v29 = v15
	v31 = v18
	goto L7
L7:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v40 = base.I32_extend8_s(v39)
	if int32(0) <= v40 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v311 = v296
	v312 = v307
	goto L3
L9:
	;
	if int32(0) < v299 {
		v28 = v296
		v29 = v307
		v31 = v299
		goto L7
	} else {
		goto L67
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
	v56 = F_pg_encoding_verifymbchar(m, int32(35), v29, v31)
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
		v311 = v28
		v312 = v29
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
	v296 = v28 + v49
	v299 = v31 - v49
	v307 = v29 + v49
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(35), v29, v31)
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
		v311 = v28
		v312 = v29
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
	F_report_invalid_encoding(m, int32(35), v29, v31)
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
	v296 = v282
	v299 = v31 - v56
	v307 = v29 + v56
	goto L9
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v40)
	v70 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v70)
	v282 = v28 + int32(2)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v78 = v75 | v39<<(uint(int32(8))%32)
	if base.Ui32(v78-int32(60736)) <= base.Ui32(int32(767)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v86 = v78
	v89 = v75
	v91 = int32(0)
	v92 = int32(61167)
	v93 = v39
	v94 = int32(2230784)
	goto L31
L29:
	;
	v118 = v78
	v121 = v75
	v125 = v39
	goto L30
L30:
	;
	if v118 <= int32(60222) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	if v92&int32(65535) == v86 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v118 = v106
	v121 = v107
	v125 = v108
	goto L30
L33:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+2)))
	v106 = v101
	v107 = v101 & int32(255)
	v108 = int32(base.Ui32(v101) >> (uint(int32(8)) % 32))
	goto L35
L34:
	;
	v106 = v86
	v107 = v89
	v108 = v93
	goto L35
L35:
	;
	v110 = v91 + int32(1)
	v112 = v110 << (uint(int32(3)) % 32)
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+uint32(_consts[1074]))))
	if v115 != int32(65535) {
		v86 = v106
		v89 = v107
		v91 = v110
		v92 = v115
		v93 = v108
		v94 = v112 + int32(2230784)
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v135 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v121))
	if base.Ui32(int32(158)) < base.Ui32(v121) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v157 = int32(0)
	if base.B2i32(base.B2i32(v118 != int32(60223))&base.B2i32(base.Ui32(v118) < base.Ui32(int32(61504))) == v157)&base.B2i32(base.Ui32(int32(176)) < base.Ui32(v118-int32(64588))) == v157 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v136 = int32(2)
	goto L42
L41:
	;
	v136 = int32(96)
	goto L42
L42:
	;
	v140 = v136 + v121 + base.B2i32(base.Ui32(v121) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v140)
	v148 = v125<<(uint(int32(1))%32)&int32(126) | v135 + int32(159)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v148)
	v282 = v28 + int32(2)
	goto L24
L43:
	;
	v166 = int32(44706)
	*(*uint16)(unsafe.Add(mBase, uint32(v28))) = uint16(v166)
	v282 = v28 + int32(2)
	goto L24
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(v118-int32(61504)) <= base.Ui32(int32(1279)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v177 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v121))
	if base.Ui32(int32(158)) < base.Ui32(v121) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(v118-int32(62784)) <= base.Ui32(int32(1279)) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v178 = int32(2)
	goto L51
L50:
	;
	v178 = int32(96)
	goto L51
L51:
	;
	v182 = v178 + v121 + base.B2i32(base.Ui32(v121) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v182)
	v192 = (v125<<(uint(int32(1))%32)+int32(34))&int32(126) | v177 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v192)
	v282 = v28 + int32(2)
	goto L24
L52:
	;
	v200 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v200)
	v205 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v121))
	if base.Ui32(int32(158)) < base.Ui32(v121) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v225 = int32(64064)
	if base.Ui32(v118) < base.Ui32(v225) {
		v282 = v28
		goto L24
	} else {
		goto L58
	}
L55:
	;
	v206 = int32(2)
	goto L57
L56:
	;
	v206 = int32(96)
	goto L57
L57:
	;
	v210 = v206 + v121 + base.B2i32(base.Ui32(v121) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)) = uint8(v210)
	v220 = (v125<<(uint(int32(1))%32)+int32(24))&int32(126) | v205 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v220)
	v282 = v28 + int32(3)
	goto L24
L58:
	;
	v230 = v118
	v231 = v28
	v235 = v225
	v236 = int32(0)
	v238 = int32(2230784)
	goto L59
L59:
	;
	if v230 != v235&int32(65535) {
		v270 = v230
		v271 = v231
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v282 = v271
	goto L24
L61:
	;
	v273 = v236 + int32(1)
	v275 = v273 << (uint(int32(3)) % 32)
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+uint32(_consts[1075]))))
	if v278 != int32(65535) {
		v230 = v270
		v231 = v271
		v235 = v278
		v236 = v273
		v238 = v275 + int32(2230784)
		goto L59
	} else {
		goto L66
	}
L62:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if int32(9371648) <= v245 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v248 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v248)
	v250 = int32(128)
	v251 = v245 | v250
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+2)) = uint8(v251)
	v256 = int32(base.Ui32(v245)>>(uint(int32(8))%32)) | v250
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)) = uint8(v256)
	v270 = v245
	v271 = v231 + int32(3)
	goto L61
L64:
	;
	goto L65
L65:
	;
	v260 = int32(128)
	v261 = v245 | v260
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)) = uint8(v261)
	v266 = int32(base.Ui32(v245)>>(uint(int32(8))%32)) | v260
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v266)
	v270 = v245
	v271 = v231 + int32(2)
	goto L61
L66:
	;
	goto L60
L67:
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v17) < base.Ui32(v18) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v117
L2:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(v18) < base.Ui32(v79) {
		goto L50
	} else {
		goto L51
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(v18) < base.Ui32(v34) {
		v117 = l1
		goto L1
	} else {
		goto L13
	}
L4:
	;
	if base.Ui32(v18) < base.Ui32(v17) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if base.Ui32(v15) < base.Ui32(v13) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.Ui32(v13) < base.Ui32(v15) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(v16) < base.Ui32(v14) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(v14) < base.Ui32(v16) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v25 < v26 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v26 < v25 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v30) <= base.Ui32(v29) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L3
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v34) < base.Ui32(v18) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.Ui32(v17) < base.Ui32(v34) {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if base.Ui32(v13) < base.Ui32(v36) {
		v117 = l1
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v36) < base.Ui32(v13) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(v14) < base.Ui32(v37) {
		v117 = l1
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v37) < base.Ui32(v14) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v43 < v44 {
		v117 = l1
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v44 < v43 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v47) < base.Ui32(v48) {
		v117 = l1
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	return l2
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v34) < base.Ui32(v17) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return l0
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(v15) < base.Ui32(v36) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return l2
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(v36) < base.Ui32(v15) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return l0
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(v16) < base.Ui32(v37) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return l2
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(v37) < base.Ui32(v16) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return l0
L39:
	;
	goto L40
L40:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v64 < v65 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return l2
L42:
	;
	goto L43
L43:
	;
	if v65 < v64 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return l0
L45:
	;
	goto L46
L46:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v70) < base.Ui32(v71) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v73 = l2
	goto L49
L48:
	;
	v73 = l0
	goto L49
L49:
	;
	return v73
L50:
	;
	if base.Ui32(v17) < base.Ui32(v79) {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	if base.Ui32(v79) < base.Ui32(v18) {
		v117 = l1
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(v13) < base.Ui32(v77) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	if base.Ui32(v77) < base.Ui32(v13) {
		v117 = l1
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32(v14) < base.Ui32(v78) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	if base.Ui32(v78) < base.Ui32(v14) {
		v117 = l1
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v86 < v87 {
		goto L50
	} else {
		goto L57
	}
L57:
	;
	if v87 < v86 {
		v117 = l1
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v91) < base.Ui32(v90) {
		v117 = l1
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L50
L60:
	;
	return l0
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v79) < base.Ui32(v17) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	return l2
L64:
	;
	goto L65
L65:
	;
	if base.Ui32(v15) < base.Ui32(v77) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	return l0
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(v77) < base.Ui32(v15) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	return l2
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(v16) < base.Ui32(v78) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	return l0
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(v78) < base.Ui32(v16) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	return l2
L76:
	;
	goto L77
L77:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v107 < v108 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	return l0
L79:
	;
	goto L80
L80:
	;
	if v108 < v107 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	return l2
L82:
	;
	goto L83
L83:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v113) < base.Ui32(v114) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v116 = l0
	goto L86
L85:
	;
	v116 = l2
	goto L86
L86:
	;
	v117 = v116
	goto L1
}
func F_spcache_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, _consts[221]))
	if v4 == int32(0) {
		v15 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v15)
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, _consts[221])) = v18
		*(*int32)(unsafe.Add(mBase, _consts[222])) = v18
		v27 = *(*int32)(unsafe.Add(mBase, _consts[212]))
		if v27 == v18 {
			v32 = *(*int32)(unsafe.Add(mBase, _consts[87]))
			v37 = F_AllocSetContextCreateInternal(m, v32, int32(399289), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[212])) = v37
				v44 = v37
				v46 = F_MemoryContextAllocZero(m, v44, int32(32))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
					v53 = F_MemoryContextAllocExtended(m, v44, int32(768), int32(5))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
						*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
						v61 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v61)
						*(*int32)(unsafe.Add(mBase, _consts[221])) = v46
						return
					}
				}
			}
		} else {
			F_MemoryContextReset(m, v27)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[212]))
				v44 = v43
				v46 = F_MemoryContextAllocZero(m, v44, int32(32))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
					v53 = F_MemoryContextAllocExtended(m, v44, int32(768), int32(5))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
						*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
						v61 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v61)
						*(*int32)(unsafe.Add(mBase, _consts[221])) = v46
						return
					}
				}
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[127])))
		if v8 == int32(0) {
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v15)
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v18)
			*(*int32)(unsafe.Add(mBase, _consts[221])) = v18
			*(*int32)(unsafe.Add(mBase, _consts[222])) = v18
			v27 = *(*int32)(unsafe.Add(mBase, _consts[212]))
			if v27 == v18 {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[87]))
				v37 = F_AllocSetContextCreateInternal(m, v32, int32(399289), int32(0), int32(8192), int32(8388608))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[212])) = v37
					v44 = v37
					v46 = F_MemoryContextAllocZero(m, v44, int32(32))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
						v53 = F_MemoryContextAllocExtended(m, v44, int32(768), int32(5))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
							*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
							v61 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v61)
							*(*int32)(unsafe.Add(mBase, _consts[221])) = v46
							return
						}
					}
				}
			} else {
				F_MemoryContextReset(m, v27)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, _consts[212]))
					v44 = v43
					v46 = F_MemoryContextAllocZero(m, v44, int32(32))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
						v53 = F_MemoryContextAllocExtended(m, v44, int32(768), int32(5))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
							*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
							v61 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v61)
							*(*int32)(unsafe.Add(mBase, _consts[221])) = v46
							return
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			if base.Ui32(v11) < base.Ui32(int32(256)) {
				return
			} else {
				v15 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _consts[126])) = uint8(v15)
				v18 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v18)
				*(*int32)(unsafe.Add(mBase, _consts[221])) = v18
				*(*int32)(unsafe.Add(mBase, _consts[222])) = v18
				v27 = *(*int32)(unsafe.Add(mBase, _consts[212]))
				if v27 == v18 {
					v32 = *(*int32)(unsafe.Add(mBase, _consts[87]))
					v37 = F_AllocSetContextCreateInternal(m, v32, int32(399289), int32(0), int32(8192), int32(8388608))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[212])) = v37
						v44 = v37
						v46 = F_MemoryContextAllocZero(m, v44, int32(32))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
							v53 = F_MemoryContextAllocExtended(m, v44, int32(768), int32(5))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
								*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
								v61 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v61)
								*(*int32)(unsafe.Add(mBase, _consts[221])) = v46
								return
							}
						}
					}
				} else {
					F_MemoryContextReset(m, v27)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _consts[212]))
						v44 = v43
						v46 = F_MemoryContextAllocZero(m, v44, int32(32))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
							v53 = F_MemoryContextAllocExtended(m, v44, int32(768), int32(5))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v46)+12)) = int64(120259084319)
								*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(32)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v53
								v61 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[127])) = uint8(v61)
								*(*int32)(unsafe.Add(mBase, _consts[221])) = v46
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
	var v25 int32
	_ = v25
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
		*(*int32)(unsafe.Add(mBase, uint32(v3)+108)) = int32(261)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+104)) = int32(262)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+100)) = int32(263)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+96)) = int32(264)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+92)) = int32(265)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+88)) = int32(266)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+84)) = int32(267)
		v25 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+80)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+76)) = int32(268)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = int32(269)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+68)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+64)) = int32(270)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+60)) = int32(271)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(272)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(273)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+44)) = int32(274)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+40)) = int32(275)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(276)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+32)) = v25
		*(*int32)(unsafe.Add(mBase, uint32(v3)+26)) = int32(50331649)
		*(*int64)(unsafe.Add(mBase, uint32(v3)+18)) = int64(16842753)
		*(*uint16)(unsafe.Add(mBase, uint32(v3)+16)) = uint16(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+136)) = v25
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
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
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int64
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int64
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l1 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if l3 != 0 {
		goto L10
	} else {
		goto L11
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = v14 * int32(48)
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v20 = F__emscripten_memcpy_bulkmem(m, v17, l1, v19)
	mBase = m.M
	goto L7
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+112)) = v139
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v149 {
		goto L36
	} else {
		goto L37
	}
L9:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+108)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+116)) = v135
	v139 = int32(0)
	v143 = v133
	goto L8
L10:
	;
	if v23 <= int32(0) {
		v132 = v23
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v63 = v23
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+108)) = v63
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+116)) = v73
	v75 = int32(0)
	if v63 <= v75 {
		v139 = v75
		v143 = v71
		goto L8
	} else {
		goto L24
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = v23 * int32(48)
	if v28 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v31 <= int32(0) {
		v132 = v31
		goto L9
	} else {
		goto L18
	}
L15:
	;
	v29 = F__emscripten_memcpy_bulkmem(m, v26, l3, v28)
	mBase = m.M
	goto L17
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	v38 = int32(0)
	goto L19
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v38*int32(48))+20))
	v50 = F_get_func_rettype(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v63 = v59
	goto L12
L21:
	;
	return
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v38<<(uint(int32(2))%32)))) = v50
	v58 = v38 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v58 < v59 {
		v38 = v58
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v78 <= int32(0) {
		v139 = v75
		v143 = v71
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v83 = v75
	v85 = int32(0)
	goto L26
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v71)+116))
	v96 = v93 + v85*int32(48)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v139 = v121
	v143 = v71
	goto L8
L28:
	;
	if v83 != v85 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v121 = v83
	v122 = int32(-1)
	goto L30
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v71)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v85<<(uint(int32(2))%32)))) = v122
	v129 = v85 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v129 < v130 {
		v83 = v121
		v85 = v129
		goto L26
	} else {
		goto L34
	}
L31:
	;
	v105 = v93 + v83*int32(48)
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v96)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+40)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v96)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+32)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v116
	goto L33
L32:
	;
	goto L33
L33:
	;
	v121 = v83 + int32(1)
	v122 = v83
	goto L30
L34:
	;
	goto L27
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+100)) = v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_MemoryContextReset(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L21
	} else {
		goto L52
	}
L36:
	;
	v152 = int32(0)
	v157 = v152
	v160 = v149
	v161 = v152
	v162 = v6
	v163 = v6
	goto L40
L37:
	;
	goto L38
L38:
	;
	v217 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v143)+96)) = uint16(v217)
	v227 = int32(0)
	goto L35
L39:
	;
	v214 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v143)+96)) = uint16(v214)
	v227 = v214
	goto L35
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v167 = v164 + v157*int32(48)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v168&int32(64) != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v200&v201 != 0 {
		goto L39
	} else {
		goto L51
	}
L42:
	;
	v203 = v157 + int32(1)
	if v203 < v198 {
		v157 = v203
		v160 = v198
		v161 = v199
		v162 = v200
		v163 = v201
		goto L40
	} else {
		goto L50
	}
L43:
	;
	v198 = v160
	v199 = v161
	v200 = v162
	v201 = int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	if v168&int32(128) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v198 = v160
	v199 = v161
	v200 = int32(1)
	v201 = v163
	goto L42
L47:
	;
	goto L48
L48:
	;
	if v168&int32(1) != 0 {
		goto L39
	} else {
		goto L49
	}
L49:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v143)+104))
	v180 = v177 + v161*int32(48)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v167)))
	*(*int64)(unsafe.Add(mBase, uint32(v180))) = v181
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v167)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+40)) = v183
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v167)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+32)) = v185
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v167)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+24)) = v187
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v167)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+16)) = v189
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v167)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v191
	v193 = int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v198 = v196
	v199 = v161 + v193
	v200 = v193
	v201 = v163
	goto L42
L50:
	;
	goto L41
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+97)) = uint8(v200)
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+96)) = uint8(v201)
	v227 = v199
	goto L35
L52:
	;
	v234 = int32(4515248)
	v235 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v237
	v240 = F_pairingheap_allocate(m, int32(256), v11)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v240
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+96)))
	if v243 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v247 = F_palloc(m, int32(40))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L21
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+97)))
	if v263 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v249 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v247)+32)) = uint16(v249)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+28)) = int32(131072)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+34)) = v249
	v255 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v247)+12)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v247)+20)) = v255
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	F_pairingheap_add(m, v259, v247)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	v272 = F_palloc(m, v267<<(uint(int32(3))%32)+int32(40))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L21
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v235
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	if v305 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L62:
	;
	v274 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v272)+34)) = uint8(v274)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	if v274 < v276 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v282 = v276 << (uint(int32(3)) % 32)
	if v282 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v272)+37)) = uint8(v285)
	*(*uint16)(unsafe.Add(mBase, uint32(v272)+35)) = uint16(v285)
	v289 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v272)+32)) = uint16(v289)
	*(*int32)(unsafe.Add(mBase, uint32(v272)+28)) = int32(65536)
	v293 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v272)+12)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v272)+20)) = v293
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	F_pairingheap_add(m, v297, v272)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L21
	} else {
		goto L70
	}
L66:
	;
	goto L65
L67:
	;
	v283 = F__emscripten_memcpy_bulkmem(m, v272+int32(40), v266, v282)
	mBase = m.M
	goto L69
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L61
L71:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+208)))
	if v345 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L72:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v308 <= int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v315 = v308
	v317 = int32(0)
	goto L74
L74:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(5120)+v317<<(uint(int32(2))%32))))
	if v327 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L71
L76:
	;
	F_pfree(m, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L21
	} else {
		goto L79
	}
L77:
	;
	v331 = v315
	goto L78
L78:
	;
	v333 = v317 + int32(1)
	if v333 < v331 {
		v315 = v331
		v317 = v333
		goto L74
	} else {
		goto L80
	}
L79:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	v331 = v330
	goto L78
L80:
	;
	goto L75
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+216)) = int64(0)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+272))
	if v387 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v348 <= int32(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v357 = int32(0)
	goto L84
L84:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(3488)+v357<<(uint(int32(2))%32))))
	F_pfree(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L21
	} else {
		goto L86
	}
L85:
	;
	goto L81
L86:
	;
	v371 = v357 + int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v371 < v372 {
		v357 = v371
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v403 != 0 {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+268)))
	if v390 != int32(1) {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	v397 = v387
	goto L91
L91:
	;
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v397)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v397)+16)) = v398 + int64(1)
	goto L88
L92:
	;
	F_pgstat_assoc_relation(m, v386)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L21
	} else {
		goto L93
	}
L93:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+272))
	v397 = v396
	goto L91
L94:
	;
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v403)))
	*(*int64)(unsafe.Add(mBase, uint32(v403))) = v404 + int64(1)
	goto L96
L95:
	;
	goto L96
L96:
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v78 = F_palloc(m, int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L29
	}
L2:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v9 != int32(1) {
		v30 = l0
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v75 = int32(0)
	goto L4
L4:
	;
	return v75
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v32 = F_list_member(m, v31, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L12
	}
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+45)))
	if v14 != int32(1) {
		v30 = l0
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if v17 == int32(0) {
		v30 = l0
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+324))
	v21 = F_bms_make_singleton(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v26 = F_remove_nulling_relids(m, l0, v21, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v30 = v26
	goto L5
L12:
	;
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v35 = F_palloc(m, int32(8))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v46 - int32(6) {
	case 0, 3, 4, 5:
		goto L21
	case 1, 2, 6, 7, 8, 10:
		goto L18
	case 9:
		goto L19
	case 11:
		goto L20
	default:
		goto L22
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = l0
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v41 = F_lappend(m, v40, v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41
	return int32(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	v70 = F_expression_tree_walker_impl(m, l0, int32(895), l1)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L28
	}
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v66 != 0 {
		goto L1
	} else {
		goto L27
	}
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v63 != int32(1) {
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v52 = F_palloc(m, int32(8))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	if v46 != int32(319) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = l0
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v58 = F_lappend(m, v57, v52)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v58
	return int32(0)
L26:
	;
	goto L1
L27:
	;
	goto L18
L28:
	;
	v75 = v70
	goto L4
L29:
	;
	v81 = l1 + int32(32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = l0
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v86
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v88
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v88
	v93 = F_expression_tree_walker_impl(m, l0, int32(895), l1)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v97 = v95 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v98 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v101 = v99
	goto L33
L32:
	;
	v101 = int32(0)
	goto L33
L33:
	;
	if v101 <= v97 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v104 = F_lappend(m, v98, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L37
	}
L35:
	;
	v118 = v98
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v121 = v97 << (uint(int32(2)) % 32)
	v122 = v119 + v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = F_lappend(m, v123, v78)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v104
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v109 = F_lappend(m, v107, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v114 = F_lappend(m, v112, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v114
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = v117
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v129 = v128 + v121
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v132 = F_list_concat(m, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v137 = v136 + v121
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v140 = F_list_concat(m, v138, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v140
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v84
	v144 = F_lappend(m, v83, v78)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	if v97 < v82 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v147 = v82
	goto L46
L45:
	;
	v147 = v97
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v144
	return int32(0)
}
func F_sqrt_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int64
	_ = v22
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v256 int64
	_ = v256
	var v268 int64
	_ = v268
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v279 int32
	_ = v279
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v314 int64
	_ = v314
	var v328 int64
	_ = v328
	var v333 int64
	_ = v333
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v357 int64
	_ = v357
	var v369 float64
	_ = v369
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v405 int64
	_ = v405
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v465 int64
	_ = v465
	var v467 int64
	_ = v467
	var v468 int64
	_ = v468
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v504 int64
	_ = v504
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v510 int64
	_ = v510
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int64
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int64
	_ = v605
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v630 int64
	_ = v630
	var v632 int64
	_ = v632
	var v633 int64
	_ = v633
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v660 int32
	_ = v660
	var v666 int64
	_ = v666
	var v672 int32
	_ = v672
	var v691 int64
	_ = v691
	var v694 int64
	_ = v694
	var v695 int64
	_ = v695
	var v702 int64
	_ = v702
	var v707 int64
	_ = v707
	var v709 int64
	_ = v709
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v718 int64
	_ = v718
	var v720 int64
	_ = v720
	var v722 int64
	_ = v722
	var v724 int32
	_ = v724
	var v726 int64
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v742 int64
	_ = v742
	var v750 int64
	_ = v750
	var v757 int64
	_ = v757
	var v758 int64
	_ = v758
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v784 int64
	_ = v784
	var v788 int64
	_ = v788
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int64
	_ = v806
	var v807 int64
	_ = v807
	var v814 int32
	_ = v814
	var v815 int64
	_ = v815
	var v816 int64
	_ = v816
	var v817 int64
	_ = v817
	var v818 int64
	_ = v818
	var v823 int64
	_ = v823
	var v826 int64
	_ = v826
	var v829 int64
	_ = v829
	var v832 int64
	_ = v832
	var v833 int64
	_ = v833
	var v837 int64
	_ = v837
	var v844 int64
	_ = v844
	var v856 int32
	_ = v856
	var v857 int64
	_ = v857
	var v858 int64
	_ = v858
	var v862 int64
	_ = v862
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v919 int64
	_ = v919
	var v927 int64
	_ = v927
	var v934 int64
	_ = v934
	var v935 int64
	_ = v935
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v964 int64
	_ = v964
	var v969 int64
	_ = v969
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v983 int64
	_ = v983
	var v984 int64
	_ = v984
	var v991 int32
	_ = v991
	var v992 int64
	_ = v992
	var v993 int64
	_ = v993
	var v994 int64
	_ = v994
	var v995 int64
	_ = v995
	var v1000 int64
	_ = v1000
	var v1003 int64
	_ = v1003
	var v1006 int64
	_ = v1006
	var v1009 int64
	_ = v1009
	var v1010 int64
	_ = v1010
	var v1014 int64
	_ = v1014
	var v1021 int64
	_ = v1021
	var v1033 int32
	_ = v1033
	var v1034 int64
	_ = v1034
	var v1035 int64
	_ = v1035
	var v1039 int64
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1089 int32
	_ = v1089
	var v1108 int64
	_ = v1108
	var v1111 int64
	_ = v1111
	var v1112 int64
	_ = v1112
	var v1118 int64
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1125 int64
	_ = v1125
	var v1127 int64
	_ = v1127
	var v1128 int64
	_ = v1128
	var v1129 int64
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1134 int64
	_ = v1134
	var v1140 int32
	_ = v1140
	var v1158 int64
	_ = v1158
	var v1164 int64
	_ = v1164
	var v1170 int64
	_ = v1170
	var v1175 int64
	_ = v1175
	var v1177 int64
	_ = v1177
	var v1179 int64
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1186 int64
	_ = v1186
	var v1188 int64
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1192 int64
	_ = v1192
	var v1200 int32
	_ = v1200
	var v1218 int64
	_ = v1218
	var v1229 int64
	_ = v1229
	var v1236 int64
	_ = v1236
	var v1238 int64
	_ = v1238
	var v1250 int32
	_ = v1250
	var v1262 int64
	_ = v1262
	var v1263 int64
	_ = v1263
	var v1267 int64
	_ = v1267
	var v1274 int32
	_ = v1274
	var v1276 int64
	_ = v1276
	var v1281 int64
	_ = v1281
	var v1282 int64
	_ = v1282
	var v1284 int64
	_ = v1284
	var v1287 int64
	_ = v1287
	var v1288 int64
	_ = v1288
	var v1290 int64
	_ = v1290
	var v1291 int64
	_ = v1291
	var v1295 int64
	_ = v1295
	var v1302 int64
	_ = v1302
	var v1314 int32
	_ = v1314
	var v1315 int64
	_ = v1315
	var v1316 int64
	_ = v1316
	var v1319 int64
	_ = v1319
	var v1320 int64
	_ = v1320
	var v1323 int64
	_ = v1323
	var v1324 int64
	_ = v1324
	var v1325 int64
	_ = v1325
	var v1330 int64
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1339 int64
	_ = v1339
	var v1340 int64
	_ = v1340
	var v1343 int64
	_ = v1343
	var v1349 int64
	_ = v1349
	var v1350 int64
	_ = v1350
	var v1353 int64
	_ = v1353
	var v1359 int64
	_ = v1359
	var v1360 int64
	_ = v1360
	var v1361 int64
	_ = v1361
	var v1362 int64
	_ = v1362
	var v1375 int32
	_ = v1375
	var v1376 int64
	_ = v1376
	var v1377 int64
	_ = v1377
	var v1382 int64
	_ = v1382
	var v1383 int64
	_ = v1383
	var v1385 int64
	_ = v1385
	var v1388 int64
	_ = v1388
	var v1389 int64
	_ = v1389
	var v1391 int64
	_ = v1391
	var v1392 int64
	_ = v1392
	var v1396 int64
	_ = v1396
	var v1403 int64
	_ = v1403
	var v1415 int32
	_ = v1415
	var v1420 int64
	_ = v1420
	var v1421 int64
	_ = v1421
	var v1423 int64
	_ = v1423
	var v1426 int64
	_ = v1426
	var v1427 int64
	_ = v1427
	var v1429 int64
	_ = v1429
	var v1430 int64
	_ = v1430
	var v1434 int64
	_ = v1434
	var v1441 int64
	_ = v1441
	var v1453 int32
	_ = v1453
	var v1454 int64
	_ = v1454
	var v1455 int64
	_ = v1455
	var v1456 int64
	_ = v1456
	var v1465 int64
	_ = v1465
	var v1466 int64
	_ = v1466
	var v1468 int64
	_ = v1468
	var v1471 int64
	_ = v1471
	var v1472 int64
	_ = v1472
	var v1474 int64
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1479 int64
	_ = v1479
	var v1486 int64
	_ = v1486
	var v1498 int32
	_ = v1498
	var v1500 int64
	_ = v1500
	var v1503 int64
	_ = v1503
	var v1504 int64
	_ = v1504
	var v1509 int64
	_ = v1509
	var v1510 int64
	_ = v1510
	var v1513 int64
	_ = v1513
	var v1516 int64
	_ = v1516
	var v1517 int64
	_ = v1517
	var v1524 int64
	_ = v1524
	var v1535 int64
	_ = v1535
	var v1536 int64
	_ = v1536
	var v1538 int64
	_ = v1538
	var v1540 int64
	_ = v1540
	var v1545 int64
	_ = v1545
	var v1546 int64
	_ = v1546
	var v1547 int64
	_ = v1547
	var v1550 int64
	_ = v1550
	var v1552 int64
	_ = v1552
	var v1553 int64
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1557 int64
	_ = v1557
	var v1559 int64
	_ = v1559
	var v1560 int64
	_ = v1560
	var v1563 int64
	_ = v1563
	var v1566 int64
	_ = v1566
	var v1567 int64
	_ = v1567
	var v1574 int64
	_ = v1574
	var v1579 int32
	_ = v1579
	var v1580 int64
	_ = v1580
	var v1581 int64
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1600 int64
	_ = v1600
	var v1608 int64
	_ = v1608
	var v1615 int64
	_ = v1615
	var v1616 int64
	_ = v1616
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1642 int64
	_ = v1642
	var v1646 int64
	_ = v1646
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1664 int64
	_ = v1664
	var v1665 int64
	_ = v1665
	var v1671 int64
	_ = v1671
	var v1672 int64
	_ = v1672
	var v1673 int64
	_ = v1673
	var v1674 int64
	_ = v1674
	var v1679 int64
	_ = v1679
	var v1682 int64
	_ = v1682
	var v1685 int64
	_ = v1685
	var v1688 int64
	_ = v1688
	var v1689 int64
	_ = v1689
	var v1693 int64
	_ = v1693
	var v1700 int64
	_ = v1700
	var v1712 int32
	_ = v1712
	var v1713 int64
	_ = v1713
	var v1714 int64
	_ = v1714
	var v1718 int64
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1813 int32
	_ = v1813
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1849 int32
	_ = v1849
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1908 int32
	_ = v1908
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1945 int64
	_ = v1945
	var v1949 int64
	_ = v1949
	var v1955 int64
	_ = v1955
	var v1957 int64
	_ = v1957
	var v1958 int64
	_ = v1958
	var v1959 int64
	_ = v1959
	var v1960 int64
	_ = v1960
	var v1966 int64
	_ = v1966
	var v1968 int64
	_ = v1968
	var v1970 int64
	_ = v1970
	var v1978 int64
	_ = v1978
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2005 int64
	_ = v2005
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2033 int64
	_ = v2033
	var v2039 int64
	_ = v2039
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2063 int64
	_ = v2063
	var v2076 int32
	_ = v2076
	var v2078 int64
	_ = v2078
	var v2081 int64
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2156 int32
	_ = v2156
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2229 int32
	_ = v2229
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2284 int32
	_ = v2284
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2324 int32
	_ = v2324
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2392 int32
	_ = v2392
	var v2393 int64
	_ = v2393
	var v2412 int32
	_ = v2412
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2461 int32
	_ = v2461
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2516 int32
	_ = v2516
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2556 int32
	_ = v2556
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2624 int32
	_ = v2624
	var v2625 int64
	_ = v2625
	var v2646 int32
	_ = v2646
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2686 int64
	_ = v2686
	var v2688 int64
	_ = v2688
	var v2694 int32
	_ = v2694
	var v2698 int32
	_ = v2698
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2712 int32
	_ = v2712
	var v2714 int32
	_ = v2714
	var v2715 int64
	_ = v2715
	var v2720 int32
	_ = v2720
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2801 int32
	_ = v2801
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2829 int32
	_ = v2829
	var v2830 int32
	_ = v2830
	var v2833 int32
	_ = v2833
	var v2874 int32
	_ = v2874
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2919 int32
	_ = v2919
	var v2923 int32
	_ = v2923
	var v2929 int32
	_ = v2929
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2938 int32
	_ = v2938
	var v2942 int32
	_ = v2942
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2965 int32
	_ = v2965
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2984 int32
	_ = v2984
	var v2985 int32
	_ = v2985
	var v2989 int32
	_ = v2989
	var v2998 int32
	_ = v2998
	var v3000 int32
	_ = v3000
	var v3007 int32
	_ = v3007
	var v3014 int32
	_ = v3014
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3025 int32
	_ = v3025
	var v3030 int32
	_ = v3030
	var v3038 int32
	_ = v3038
	var v3042 int32
	_ = v3042
	var v3047 int32
	_ = v3047
	var v3051 int32
	_ = v3051
	var v3057 int32
	_ = v3057
	var v3068 int32
	_ = v3068
	var v3078 int32
	_ = v3078
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3120 int32
	_ = v3120
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3133 int32
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3140 int32
	_ = v3140
	var v3142 int32
	_ = v3142
	var v3148 int32
	_ = v3148
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3164 int32
	_ = v3164
	var v3168 int32
	_ = v3168
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3178 int32
	_ = v3178
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3199 int32
	_ = v3199
	var v3201 int32
	_ = v3201
	var v3206 int32
	_ = v3206
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3218 int32
	_ = v3218
	var v3220 int32
	_ = v3220
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3234 int32
	_ = v3234
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3252 int32
	_ = v3252
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3270 int32
	_ = v3270
	var v3275 int32
	_ = v3275
	var v3283 int32
	_ = v3283
	var v3287 int32
	_ = v3287
	var v3292 int32
	_ = v3292
	var v3296 int32
	_ = v3296
	var v3302 int32
	_ = v3302
	var v3313 int32
	_ = v3313
	var v3323 int32
	_ = v3323
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3359 int32
	_ = v3359
	var v3361 int32
	_ = v3361
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3375 int32
	_ = v3375
	var v3377 int32
	_ = v3377
	var v3378 int64
	_ = v3378
	var v3380 int64
	_ = v3380
	var v3383 int32
	_ = v3383
	var v3387 int32
	_ = v3387
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3398 int32
	_ = v3398
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3403 int64
	_ = v3403
	var v3405 int64
	_ = v3405
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3414 int32
	_ = v3414
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3421 int32
	_ = v3421
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3432 int32
	_ = v3432
	var v3438 int32
	_ = v3438
	var v3440 int32
	_ = v3440
	var v3447 int32
	_ = v3447
	var v3455 int32
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3460 int32
	_ = v3460
	var v3466 int32
	_ = v3466
	var v3468 int32
	_ = v3468
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3481 int32
	_ = v3481
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3496 int32
	_ = v3496
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3502 int32
	_ = v3502
	var v3507 int32
	_ = v3507
	var v3519 int32
	_ = v3519
	var v3523 int32
	_ = v3523
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3533 int32
	_ = v3533
	var v3538 int32
	_ = v3538
	var v3542 int32
	_ = v3542
	var v3554 int32
	_ = v3554
	var v3556 int32
	_ = v3556
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3565 int32
	_ = v3565
	var v3573 int32
	_ = v3573
	var v3575 int32
	_ = v3575
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3598 int32
	_ = v3598
	var v3600 int32
	_ = v3600
	var v3607 int32
	_ = v3607
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3618 int32
	_ = v3618
	var v3625 int32
	_ = v3625
	var v3630 int32
	_ = v3630
	var v3638 int32
	_ = v3638
	var v3642 int32
	_ = v3642
	var v3647 int32
	_ = v3647
	var v3651 int32
	_ = v3651
	var v3657 int32
	_ = v3657
	var v3668 int32
	_ = v3668
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3693 int32
	_ = v3693
	var v3697 int32
	_ = v3697
	var v3703 int32
	_ = v3703
	var v3704 int32
	_ = v3704
	var v3705 int32
	_ = v3705
	var v3707 int32
	_ = v3707
	var v3712 int32
	_ = v3712
	var v3716 int32
	_ = v3716
	var v3728 int32
	_ = v3728
	var v3730 int32
	_ = v3730
	var v3735 int32
	_ = v3735
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3739 int32
	_ = v3739
	var v3747 int32
	_ = v3747
	var v3749 int32
	_ = v3749
	var v3758 int32
	_ = v3758
	var v3759 int32
	_ = v3759
	var v3763 int32
	_ = v3763
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3781 int32
	_ = v3781
	var v3788 int32
	_ = v3788
	var v3789 int32
	_ = v3789
	var v3792 int32
	_ = v3792
	var v3799 int32
	_ = v3799
	var v3804 int32
	_ = v3804
	var v3812 int32
	_ = v3812
	var v3816 int32
	_ = v3816
	var v3821 int32
	_ = v3821
	var v3825 int32
	_ = v3825
	var v3831 int32
	_ = v3831
	var v3842 int32
	_ = v3842
	var v3852 int32
	_ = v3852
	var v3853 int32
	_ = v3853
	var v3862 int32
	_ = v3862
	var v3867 int32
	_ = v3867
	var v3877 int32
	_ = v3877
	var v3883 int32
	_ = v3883
	var v3890 int32
	_ = v3890
	var v3892 int32
	_ = v3892
	var v3912 int32
	_ = v3912
	var v3915 int32
	_ = v3915
	var v3916 int32
	_ = v3916
	var v3917 int32
	_ = v3917
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3926 int32
	_ = v3926
	var v3928 int32
	_ = v3928
	var v3929 int64
	_ = v3929
	var v3932 int32
	_ = v3932
	var v3933 int64
	_ = v3933
	var v3935 int32
	_ = v3935
	var v3940 int32
	_ = v3940
	var v3948 int32
	_ = v3948
	var v3951 int32
	_ = v3951
	var v3960 int32
	_ = v3960
	var v3962 int32
	_ = v3962
	var v3966 int32
	_ = v3966
	var v3967 int32
	_ = v3967
	var v3977 int32
	_ = v3977
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3986 int32
	_ = v3986
	var v3990 int32
	_ = v3990
	var v3991 int32
	_ = v3991
	var v3992 int32
	_ = v3992
	var v3995 int32
	_ = v3995
	var v3998 int32
	_ = v3998
	var v4003 int32
	_ = v4003
	var v4007 int32
	_ = v4007
	var v4013 int32
	_ = v4013
	var v4019 int32
	_ = v4019
	var v4020 int32
	_ = v4020
	var v4023 int32
	_ = v4023
	var v4026 int32
	_ = v4026
	var v4028 int32
	_ = v4028
	var v4029 int32
	_ = v4029
	var v4030 int32
	_ = v4030
	var v4033 int32
	_ = v4033
	var v4041 int32
	_ = v4041
	var v4045 int32
	_ = v4045
	var v4046 int32
	_ = v4046
	var v4049 int32
	_ = v4049
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4107 int32
	_ = v4107
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4125 int32
	_ = v4125
	var v4156 int32
	_ = v4156
	var v4157 int32
	_ = v4157
	var v4165 int32
	_ = v4165
	var v4201 int32
	_ = v4201
	var v4202 int32
	_ = v4202
	var v4232 int32
	_ = v4232
	var v4234 int32
	_ = v4234
	var v4235 int32
	_ = v4235
	var v4237 int32
	_ = v4237
	var v4239 int32
	_ = v4239
	var v4241 int32
	_ = v4241
	var v4242 int32
	_ = v4242
	var v4244 int32
	_ = v4244
	var v4245 int32
	_ = v4245
	var v4249 int32
	_ = v4249
	v4 = int32(0)
	v22 = int64(0)
	v34 = m.G0
	v36 = v34 - int32(512)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v38 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v36 + int32(512)
	return
L2:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v41 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v51 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_pfree(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	v47 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v47
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v47
	goto L1
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	v478 = v36 + int32(316)
	if v198 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L11:
	;
	v54 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+328)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+296)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+304)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+272)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+280)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+248)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+256)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+224)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+232)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+208)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+200)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+312)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+288)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+264)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+240)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+216)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v36)+192)) = v54
	v90 = int32(-1)
	v91 = int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v94 = v92 >> (uint(v91) % 32)
	if int32(0) <= l2+v91 {
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
	v431 = m.ExcPending
	if v431 != 0 {
		goto L8
	} else {
		goto L59
	}
L14:
	;
	v102 = int32(4)
	v105 = base.I32_div_s(l2+v102, v102)
	v110 = v105
	goto L16
L15:
	;
	v109 = base.I32_div_s(l2^int32(-1), int32(-4))
	v110 = v109
	goto L16
L16:
	;
	v112 = int32(1)
	v113 = v110 + v94 + v112
	if v113 <= v112 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v116 = v91
	goto L19
L18:
	;
	v116 = v113
	goto L19
L19:
	;
	v118 = int32(1)
	v122 = (v94^v90+v116)<<(uint(v118)%32) + v92 + v118
	if v122 <= v118 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v125 = v91
	goto L22
L21:
	;
	v125 = v122
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+336)) = v125
	if int32(5) <= v122 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v136 = int32(0)
	v140 = v125
	goto L26
L24:
	;
	v192 = v113
	v196 = v125
	v198 = v90
	goto L25
L25:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v220 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219))))
	if v196 < int32(2) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v163 = int32(2)
	v164 = int32(base.Ui32(v140) >> (uint(v163) % 32))
	v168 = v136 + int32(1)
	if v140&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v192 = v168
	v196 = v182
	v198 = v136
	goto L25
L28:
	;
	v179 = v164
	goto L30
L29:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v174))))
	v179 = v164 - base.B2i32(v175 < int32(2500))
	goto L30
L30:
	;
	v182 = v140 - v179<<(uint(int32(1))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v36+int32(336)+v168<<(uint(v163)%32)))) = v182
	if int32(4) < v182 {
		v136 = v168
		v140 = v182
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v369 = base.F64_sqrt(base.F64_convert_i64_s(v357))
	if base.F64_lt(base.F64_abs(v369), float64(9.223372036854776e+18)) != 0 {
		goto L51
	} else {
		goto L52
	}
L33:
	;
	v341 = v192
	v345 = int32(1)
	v346 = v4
	v357 = v220
	goto L32
L34:
	;
	goto L35
L35:
	;
	v224 = int32(1)
	v226 = v196 - v224
	if v196 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v232 = v226 & int32(-2)
	v238 = v224
	v240 = int32(0)
	v256 = v220
	goto L39
L37:
	;
	v296 = v224
	v298 = v226
	v303 = v4
	v314 = v220
	goto L38
L38:
	;
	if v226&v224 == int32(0) {
		v341 = v298
		v345 = v196
		v346 = v303
		v357 = v314
		goto L32
	} else {
		goto L48
	}
L39:
	;
	v268 = v256 * int64(10000)
	if v238 < v38 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v296 = v288
	v298 = v290
	v303 = v232
	v314 = v286
	goto L38
L41:
	;
	v273 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v238<<(uint(int32(1))%32)))))
	v275 = v268 + v273
	goto L43
L42:
	;
	v275 = v268
	goto L43
L43:
	;
	v277 = v275 * int64(10000)
	v279 = v238 + int32(1)
	if v279 < v38 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v284 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v279<<(uint(int32(1))%32)))))
	v286 = v277 + v284
	goto L46
L45:
	;
	v286 = v277
	goto L46
L46:
	;
	v287 = int32(2)
	v288 = v238 + v287
	v290 = v240 + v287
	if v290 != v232 {
		v238 = v288
		v240 = v290
		v256 = v286
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	v328 = v314 * int64(10000)
	if v38 <= v296 {
		v341 = v298
		v345 = v196
		v346 = v303
		v357 = v328
		goto L32
	} else {
		goto L49
	}
L49:
	;
	v333 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v296<<(uint(int32(1))%32)))))
	v341 = v298
	v345 = v196
	v346 = v303
	v357 = v328 + v333
	goto L32
L50:
	;
	v377 = v357 - v375*v375
	if base.B2i32(int64(0) <= v377)&base.B2i32(v377 <= v375<<(uint(int64(1))%64)) != 0 {
		v465 = v375
		v467 = v22
		v468 = v377
		goto L10
	} else {
		goto L54
	}
L51:
	;
	v373 = base.I64_trunc_f64_s(v369)
	v375 = v373
	goto L50
L52:
	;
	goto L53
L53:
	;
	v375 = int64(-9223372036854775807 - 1)
	goto L50
L54:
	;
	v405 = v375
	goto L55
L55:
	;
	v417 = base.I64_div_s(v357, v405)
	v420 = base.I64_div_s(v417+v405, int64(2))
	v422 = v420 << (uint(int64(1)) % 64)
	v424 = v357 - v420*v420
	if v424 < int64(0) {
		v405 = v420
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v465 = v420
	v467 = v422
	v468 = v424
	goto L10
L57:
	;
	if v422 < v424 {
		v405 = v420
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(228207), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(499844), int32(10455), int32(229254))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2139))) = v2129
	v2156 = int32(0)
	if v2156 <= v2135 {
		goto L234
	} else {
		goto L235
	}
L64:
	;
	v2018 = F_palloc(m, int32(12))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L8
	} else {
		goto L224
	}
L65:
	;
	v1994 = v345
	v1997 = v198
	v2005 = v465
	goto L64
L66:
	;
	goto L67
L67:
	;
	v489 = v341
	v493 = v345
	v494 = v346
	v495 = v198
	v504 = v465
	v506 = v467
	v507 = v468
	v510 = v22
	goto L68
L68:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(336)+v495<<(uint(int32(2))%32))))
	if v521 <= int32(8) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v1994 = v1932
	v1997 = int32(-1)
	v2005 = v1978
	goto L64
L70:
	;
	v1957 = v1955*v507 + v1945
	v1958 = int64(1)
	v1959 = v504 << (uint(v1958) % 64)
	v1960 = base.I64_div_s(v1957, v1959)
	v1966 = v1949 - v1960*v1960 + (v1957-v1960*v1959)*v1955
	v1968 = v1960 + v504*v1955
	v1970 = v1968 - v1958
	if v1966 < int64(0) {
		goto L220
	} else {
		goto L221
	}
L71:
	;
	if v533 != 0 {
		goto L196
	} else {
		goto L197
	}
L72:
	;
	v1768 = v601
	v1770 = v597
	v1772 = v597 * int32(10000)
	v1775 = v599
	v1797 = v556 * int32(-727379968)
	goto L71
L73:
	;
	v524 = v521 - v493
	v525 = int32(2)
	v526 = base.I32_div_s(v524, v525)
	if v524 < v525 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v605 = int64(63)
	v619 = v493
	v622 = v495
	v630 = v504
	v632 = v506
	v633 = v507
	v634 = v504 >> (uint(v605) % 64)
	v636 = v510
	v638 = v507 >> (uint(v605) % 64)
	goto L91
L76:
	;
	v529 = int64(0)
	v1928 = v489
	v1932 = v493
	v1933 = v494
	v1945 = v529
	v1949 = v529
	v1955 = int64(1)
	goto L70
L77:
	;
	goto L78
L78:
	;
	v532 = int32(1)
	v533 = v526 & v532
	v535 = v526 - v532
	if v535 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v1768 = v493
	v1770 = v489
	v1772 = int32(0)
	v1775 = v494
	v1797 = int32(10000)
	goto L71
L80:
	;
	goto L81
L81:
	;
	v543 = int32(0)
	v549 = v493
	v551 = v543
	v553 = v543
	v556 = int32(1)
	goto L82
L82:
	;
	v579 = v551 * int32(10000)
	if v549 < v38 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L72
L84:
	;
	v584 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219+v549<<(uint(int32(1))%32)))))
	v586 = v579 + v584
	goto L86
L85:
	;
	v586 = v579
	goto L86
L86:
	;
	v588 = v586 * int32(10000)
	v590 = v549 + int32(1)
	if v590 < v38 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v595 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219+v590<<(uint(int32(1))%32)))))
	v597 = v588 + v595
	goto L89
L88:
	;
	v597 = v588
	goto L89
L89:
	;
	v599 = v556 * int32(100000000)
	v600 = int32(2)
	v601 = v549 + v600
	v603 = v553 + v600
	if v526&int32(1073741822) != v603 {
		v549 = v601
		v551 = v597
		v553 = v603
		v556 = v599
		goto L82
	} else {
		goto L90
	}
L90:
	;
	goto L83
L91:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(336)+v622<<(uint(int32(2))%32))))
	if v647 <= int32(16) {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v1587 = F_palloc(m, int32(22))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L8
	} else {
		goto L181
	}
L93:
	;
	v1274 = v36 + int32(112)
	v1276 = v1267 >> (uint(int64(63)) % 64)
	v1281 = int64(32)
	v1282 = int64(base.Ui64(v1267) >> (uint(v1281) % 64))
	v1284 = int64(base.Ui64(v633) >> (uint(v1281) % 64))
	v1287 = int64(4294967295)
	v1288 = v1267 & v1287
	v1290 = v633 & v1287
	v1291 = v1288 * v1290
	v1295 = int64(base.Ui64(v1291)>>(uint(v1281)%64)) + v1288*v1284
	v1302 = v1290*v1282 + v1295&v1287
	*(*int64)(unsafe.Add(mBase, uint32(v1274)+8)) = v633*v1276 + v638*v1267 + v1282*v1284 + int64(base.Ui64(v1295)>>(uint(v1281)%64)) + int64(base.Ui64(v1302)>>(uint(v1281)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1274))) = v1291&v1287 | v1302<<(uint(v1281)%64)
	goto L168
L94:
	;
	v1120 = v652 & int32(1)
	if v1120 != 0 {
		goto L144
	} else {
		goto L145
	}
L95:
	;
	v1089 = v724
	v1108 = v720
	v1111 = v720 * int64(10000)
	v1112 = v722
	v1118 = v695 * int64(1000000000000)
	goto L94
L96:
	;
	v650 = v647 - v619
	v651 = int32(2)
	v652 = base.I32_div_s(v650, v651)
	if v650 < v651 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v729 = F_palloc(m, int32(22))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L8
	} else {
		goto L114
	}
L99:
	;
	v656 = int64(0)
	v1250 = v619
	v1262 = v656
	v1263 = v656
	v1267 = int64(1)
	goto L93
L100:
	;
	goto L101
L101:
	;
	v658 = base.I64_extend_i32_s(v652)
	v660 = base.B2i32(v658 == int64(1))
	if v658 == int64(1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v1089 = v619
	v1108 = v632
	v1111 = int64(0)
	v1112 = v636
	v1118 = int64(10000)
	goto L94
L103:
	;
	goto L104
L104:
	;
	v666 = int64(0)
	v672 = v619
	v691 = v666
	v694 = v666
	v695 = int64(1)
	goto L105
L105:
	;
	v702 = v691 * int64(10000)
	if v672 < v38 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L95
L107:
	;
	v707 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v672<<(uint(int32(1))%32)))))
	v709 = v702 + v707
	goto L109
L108:
	;
	v709 = v702
	goto L109
L109:
	;
	v711 = v709 * int64(10000)
	v713 = v672 + int32(1)
	if v713 < v38 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v718 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v713<<(uint(int32(1))%32)))))
	v720 = v711 + v718
	goto L112
L111:
	;
	v720 = v711
	goto L112
L112:
	;
	v722 = v695 * int64(100000000)
	v724 = v672 + int32(2)
	v726 = v694 + int64(2)
	if v658&int64(1073741822) != v726 {
		v672 = v724
		v691 = v720
		v694 = v726
		v695 = v722
		goto L105
	} else {
		goto L113
	}
L113:
	;
	goto L106
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+328)) = v729
	v732 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v729))) = uint16(v732)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+332)) = v729 + int32(2)
	if v634 < int64(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+316)) = v881
	*(*int32)(unsafe.Add(mBase, uint32(v36)+312)) = v875
	v906 = F_palloc(m, int32(22))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L8
	} else {
		goto L129
	}
L116:
	;
	v767 = v729 + int32(22)
	v768 = v732
	v784 = v757
	v788 = v758
	goto L121
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = int64(16384)
	v742 = int64(0)
	v757 = v742 - v630
	v758 = v742 - (v634 + base.I64_extend_i32_u(base.B2i32(v630 != v742)))
	goto L116
L118:
	;
	goto L119
L119:
	;
	v750 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = v750
	if v630|v634 == v750 {
		v875 = v732
		v881 = int32(0)
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v757 = v630
	v758 = v634
	goto L116
L121:
	;
	v797 = v36 + int32(176)
	v800 = m.G0
	v801 = int32(16)
	v802 = v800 - v801
	m.G0 = v802
	F___udivmodti4(m, v802, v784, v788, int64(10000), int64(0))
	mBase = m.M
	v806 = *(*int64)(unsafe.Add(mBase, uint32(v802)))
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v802)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v797)+8)) = v807
	*(*int64)(unsafe.Add(mBase, uint32(v797))) = v806
	m.G0 = v802 + v801
	goto L123
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+332)) = v856
	v875 = v867
	v881 = v768
	goto L115
L123:
	;
	v814 = v36 + int32(160)
	v815 = *(*int64)(unsafe.Add(mBase, uint32(v36)+176))
	v816 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(184))))
	v817 = int64(55536)
	v818 = int64(0)
	v823 = int64(32)
	v826 = int64(base.Ui64(v815) >> (uint(v823) % 64))
	v829 = int64(4294967295)
	v832 = v815 & v829
	v833 = v817 * v832
	v837 = int64(base.Ui64(v833)>>(uint(v823)%64)) + v817*v826
	v844 = v832*v818 + v837&v829
	*(*int64)(unsafe.Add(mBase, uint32(v814)+8)) = v815*v818 + v816*v817 + v818*v826 + int64(base.Ui64(v837)>>(uint(v823)%64)) + int64(base.Ui64(v844)>>(uint(v823)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v814))) = v833&v829 | v844<<(uint(v823)%64)
	goto L124
L124:
	;
	v856 = v767 - int32(2)
	v857 = *(*int64)(unsafe.Add(mBase, uint32(v36)+160))
	v858 = v857 + v784
	*(*uint16)(unsafe.Add(mBase, uint32(v856))) = uint16(v858)
	v862 = int64(0)
	v867 = v768 + int32(1)
	if v788 == v862 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v868 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v784))
	goto L127
L126:
	;
	v868 = base.B2i32(v788 != v862)
	goto L127
L127:
	;
	if v868 != 0 {
		v767 = v856
		v768 = v867
		v784 = v815
		v788 = v816
		goto L121
	} else {
		goto L128
	}
L128:
	;
	goto L122
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+304)) = v906
	v909 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v906))) = uint16(v909)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+308)) = v906 + int32(2)
	if v638 < int64(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+288)) = v1051
	v2127 = v875
	v2129 = v1054
	v2132 = v619
	v2135 = v622
	v2139 = v36 + int32(292)
	goto L63
L131:
	;
	v944 = v909
	v946 = v906 + int32(22)
	v964 = v934
	v969 = v935
	goto L136
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+296)) = int64(16384)
	v919 = int64(0)
	v934 = v919 - v633
	v935 = v919 - (v638 + base.I64_extend_i32_u(base.B2i32(v633 != v919)))
	goto L131
L133:
	;
	goto L134
L134:
	;
	v927 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+296)) = v927
	if v633|v638 == v927 {
		v1051 = v909
		v1054 = int32(0)
		goto L130
	} else {
		goto L135
	}
L135:
	;
	v934 = v633
	v935 = v638
	goto L131
L136:
	;
	v974 = v36 + int32(144)
	v977 = m.G0
	v978 = int32(16)
	v979 = v977 - v978
	m.G0 = v979
	F___udivmodti4(m, v979, v964, v969, int64(10000), int64(0))
	mBase = m.M
	v983 = *(*int64)(unsafe.Add(mBase, uint32(v979)))
	v984 = *(*int64)(unsafe.Add(mBase, uint32(v979)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v974)+8)) = v984
	*(*int64)(unsafe.Add(mBase, uint32(v974))) = v983
	m.G0 = v979 + v978
	goto L138
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+308)) = v1033
	v1051 = v1044
	v1054 = v944
	goto L130
L138:
	;
	v991 = v36 + int32(128)
	v992 = *(*int64)(unsafe.Add(mBase, uint32(v36)+144))
	v993 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(152))))
	v994 = int64(55536)
	v995 = int64(0)
	v1000 = int64(32)
	v1003 = int64(base.Ui64(v992) >> (uint(v1000) % 64))
	v1006 = int64(4294967295)
	v1009 = v992 & v1006
	v1010 = v994 * v1009
	v1014 = int64(base.Ui64(v1010)>>(uint(v1000)%64)) + v994*v1003
	v1021 = v1009*v995 + v1014&v1006
	*(*int64)(unsafe.Add(mBase, uint32(v991)+8)) = v992*v995 + v993*v994 + v995*v1003 + int64(base.Ui64(v1014)>>(uint(v1000)%64)) + int64(base.Ui64(v1021)>>(uint(v1000)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v991))) = v1010&v1006 | v1021<<(uint(v1000)%64)
	goto L139
L139:
	;
	v1033 = v946 - int32(2)
	v1034 = *(*int64)(unsafe.Add(mBase, uint32(v36)+128))
	v1035 = v1034 + v964
	*(*uint16)(unsafe.Add(mBase, uint32(v1033))) = uint16(v1035)
	v1039 = int64(0)
	v1044 = v944 + int32(1)
	if v969 == v1039 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v1045 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v964))
	goto L142
L141:
	;
	v1045 = base.B2i32(v969 != v1039)
	goto L142
L142:
	;
	if v1045 != 0 {
		v944 = v1044
		v946 = v1033
		v964 = v992
		v969 = v993
		goto L136
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	if v38 <= v1089 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v1128 = v1108
	v1129 = v1112
	goto L146
L146:
	;
	v1130 = v652 + v619
	if v658 == int64(1) {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v1127 = v1111
	goto L149
L148:
	;
	v1125 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v1089<<(uint(int32(1))%32)))))
	v1127 = v1111 + v1125
	goto L149
L149:
	;
	v1128 = v1127
	v1129 = v1118
	goto L146
L150:
	;
	if v1120 == int32(0) {
		v1238 = v1218
		goto L163
	} else {
		goto L164
	}
L151:
	;
	v1200 = v1130
	v1218 = v1118
	v1229 = int64(0)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v1134 = int64(0)
	v1140 = v1130
	v1158 = v1134
	v1164 = v1134
	goto L154
L154:
	;
	v1170 = v1158 * int64(10000)
	if v1140 < v38 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v1200 = v1190
	v1218 = v1188
	v1229 = v1188 * int64(10000)
	goto L150
L156:
	;
	v1175 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v1140<<(uint(int32(1))%32)))))
	v1177 = v1170 + v1175
	goto L158
L157:
	;
	v1177 = v1170
	goto L158
L158:
	;
	v1179 = v1177 * int64(10000)
	v1181 = v1140 + int32(1)
	if v1181 < v38 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1186 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v1181<<(uint(int32(1))%32)))))
	v1188 = v1179 + v1186
	goto L161
L160:
	;
	v1188 = v1179
	goto L161
L161:
	;
	v1190 = v1140 + int32(2)
	v1192 = v1164 + int64(2)
	if v1192 != v658&int64(1073741822) {
		v1140 = v1190
		v1158 = v1188
		v1164 = v1192
		goto L154
	} else {
		goto L162
	}
L162:
	;
	goto L155
L163:
	;
	v1250 = v652 + v1130
	v1262 = v1238
	v1263 = v1128
	v1267 = v1129
	goto L93
L164:
	;
	if v38 <= v1200 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1238 = v1229
	goto L163
L166:
	;
	goto L167
L167:
	;
	v1236 = int64(*(*int16)(unsafe.Add(mBase, uint32(v219+v1200<<(uint(int32(1))%32)))))
	v1238 = v1229 + v1236
	goto L163
L168:
	;
	v1314 = v36 + int32(96)
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v36)+112))
	v1316 = v1315 + v1263
	v1319 = *(*int64)(unsafe.Add(mBase, uint32(v36)+120))
	v1320 = int64(63)
	v1323 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v1316) < base.Ui64(v1315))) + (v1319 + v1263>>(uint(v1320)%64))
	v1324 = int64(1)
	v1325 = v630 << (uint(v1324) % 64)
	v1330 = v634<<(uint(v1324)%64) | int64(base.Ui64(v630)>>(uint(v1320)%64))
	v1334 = m.G0
	v1335 = int32(16)
	v1336 = v1334 - v1335
	m.G0 = v1336
	v1339 = v1323 >> (uint(v1320) % 64)
	v1340 = v1339 ^ v1316
	v1343 = v1340 + int64(base.Ui64(v1323)>>(uint(v1320)%64))
	v1349 = v1330 >> (uint(v1320) % 64)
	v1350 = v1349 ^ v1325
	v1353 = v1350 + int64(base.Ui64(v1330)>>(uint(v1320)%64))
	F___udivmodti4(m, v1336, v1343, base.I64_extend_i32_u(base.B2i32(base.Ui64(v1343) < base.Ui64(v1340)))+(v1323^v1339), v1353, base.I64_extend_i32_u(base.B2i32(base.Ui64(v1353) < base.Ui64(v1350)))+(v1349^v1330))
	mBase = m.M
	v1359 = *(*int64)(unsafe.Add(mBase, uint32(v1336)+8))
	v1360 = v1349 ^ v1339
	v1361 = *(*int64)(unsafe.Add(mBase, uint32(v1336)))
	v1362 = v1360 ^ v1361
	*(*int64)(unsafe.Add(mBase, uint32(v1314))) = v1362 - v1360
	*(*int64)(unsafe.Add(mBase, uint32(v1314)+8)) = v1359 ^ v1360 - v1360 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1362) < base.Ui64(v1360)))
	m.G0 = v1336 + v1335
	goto L169
L169:
	;
	v1375 = v36 + int32(80)
	v1376 = *(*int64)(unsafe.Add(mBase, uint32(v36)+96))
	v1377 = *(*int64)(unsafe.Add(mBase, uint32(v36)+104))
	v1382 = int64(32)
	v1383 = int64(base.Ui64(v1325) >> (uint(v1382) % 64))
	v1385 = int64(base.Ui64(v1376) >> (uint(v1382) % 64))
	v1388 = int64(4294967295)
	v1389 = v1325 & v1388
	v1391 = v1376 & v1388
	v1392 = v1389 * v1391
	v1396 = int64(base.Ui64(v1392)>>(uint(v1382)%64)) + v1389*v1385
	v1403 = v1391*v1383 + v1396&v1388
	*(*int64)(unsafe.Add(mBase, uint32(v1375)+8)) = v1376*v1330 + v1377*v1325 + v1383*v1385 + int64(base.Ui64(v1396)>>(uint(v1382)%64)) + int64(base.Ui64(v1403)>>(uint(v1382)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1375))) = v1392&v1388 | v1403<<(uint(v1382)%64)
	goto L170
L170:
	;
	v1415 = v36 + int32(32)
	v1420 = int64(32)
	v1421 = int64(base.Ui64(v1267) >> (uint(v1420) % 64))
	v1423 = int64(base.Ui64(v630) >> (uint(v1420) % 64))
	v1426 = int64(4294967295)
	v1427 = v1267 & v1426
	v1429 = v630 & v1426
	v1430 = v1427 * v1429
	v1434 = int64(base.Ui64(v1430)>>(uint(v1420)%64)) + v1427*v1423
	v1441 = v1429*v1421 + v1434&v1426
	*(*int64)(unsafe.Add(mBase, uint32(v1415)+8)) = v630*v1276 + v634*v1267 + v1421*v1423 + int64(base.Ui64(v1434)>>(uint(v1420)%64)) + int64(base.Ui64(v1441)>>(uint(v1420)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1415))) = v1430&v1426 | v1441<<(uint(v1420)%64)
	goto L171
L171:
	;
	v1453 = v36 - int32(-64)
	v1454 = *(*int64)(unsafe.Add(mBase, uint32(v36)+80))
	v1455 = v1316 - v1454
	v1456 = *(*int64)(unsafe.Add(mBase, uint32(v36)+88))
	v1465 = int64(32)
	v1466 = int64(base.Ui64(v1267) >> (uint(v1465) % 64))
	v1468 = int64(base.Ui64(v1455) >> (uint(v1465) % 64))
	v1471 = int64(4294967295)
	v1472 = v1267 & v1471
	v1474 = v1455 & v1471
	v1475 = v1472 * v1474
	v1479 = int64(base.Ui64(v1475)>>(uint(v1465)%64)) + v1472*v1468
	v1486 = v1474*v1466 + v1479&v1471
	*(*int64)(unsafe.Add(mBase, uint32(v1453)+8)) = v1455*v1276 + (v1323-v1456-base.I64_extend_i32_u(base.B2i32(base.Ui64(v1316) < base.Ui64(v1454))))*v1267 + v1466*v1468 + int64(base.Ui64(v1479)>>(uint(v1465)%64)) + int64(base.Ui64(v1486)>>(uint(v1465)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1453))) = v1475&v1471 | v1486<<(uint(v1465)%64)
	goto L172
L172:
	;
	v1498 = v36 + int32(48)
	v1500 = v1376 * v1377
	v1503 = int64(32)
	v1504 = int64(base.Ui64(v1376) >> (uint(v1503) % 64))
	v1509 = int64(4294967295)
	v1510 = v1376 & v1509
	v1513 = v1510 * v1510
	v1516 = v1510 * v1504
	v1517 = int64(base.Ui64(v1513)>>(uint(v1503)%64)) + v1516
	v1524 = v1516 + v1517&v1509
	*(*int64)(unsafe.Add(mBase, uint32(v1498)+8)) = v1500 + v1500 + v1504*v1504 + int64(base.Ui64(v1517)>>(uint(v1503)%64)) + int64(base.Ui64(v1524)>>(uint(v1503)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v1498))) = v1513&v1509 | v1524<<(uint(v1503)%64)
	goto L173
L173:
	;
	v1535 = *(*int64)(unsafe.Add(mBase, uint32(v36)+72))
	v1536 = int64(63)
	v1538 = *(*int64)(unsafe.Add(mBase, uint32(v36)+56))
	v1540 = *(*int64)(unsafe.Add(mBase, uint32(v36)+48))
	v1545 = v1262 - v1540
	v1546 = *(*int64)(unsafe.Add(mBase, uint32(v36)+64))
	v1547 = v1545 + v1546
	v1550 = v1535 + (v1262>>(uint(v1536)%64) - v1538 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v1262) < base.Ui64(v1540)))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1547) < base.Ui64(v1545)))
	v1552 = v1550 >> (uint(v1536) % 64)
	v1553 = *(*int64)(unsafe.Add(mBase, uint32(v36)+32))
	v1554 = v1376 + v1553
	v1557 = *(*int64)(unsafe.Add(mBase, uint32(v36)+40))
	v1559 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v1554) < base.Ui64(v1376))) + (v1377 + v1557)
	v1560 = int64(0)
	v1563 = v1559 - base.I64_extend_i32_u(base.B2i32(v1554 == v1560))
	v1566 = v1554 - int64(1)
	v1567 = v1566 + v1554
	v1574 = v1547 + v1552&v1567
	v1579 = base.B2i32(v1550 < v1560)
	if v1550 < v1560 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1580 = v1563
	goto L176
L175:
	;
	v1580 = v1559
	goto L176
L176:
	;
	if v1550 < v1560 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v1581 = v1566
	goto L179
L178:
	;
	v1581 = v1554
	goto L179
L179:
	;
	v1585 = v622 - int32(1)
	if int32(0) < v622 {
		v619 = v1250
		v622 = v1585
		v630 = v1581
		v632 = v1559
		v633 = v1574
		v634 = v1580
		v636 = v1550
		v638 = v1550 + v1552&(v1563+v1559+base.I64_extend_i32_u(base.B2i32(base.Ui64(v1567) < base.Ui64(v1566)))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1574) < base.Ui64(v1547)))
		goto L91
	} else {
		goto L180
	}
L180:
	;
	goto L92
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+328)) = v1587
	v1590 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1587))) = uint16(v1590)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+332)) = v1587 + int32(2)
	if v1580 < int64(0) {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+312)) = v1731
	v2127 = v1731
	v2129 = v1733
	v2132 = v1250
	v2135 = v1585
	v2139 = v478
	goto L63
L183:
	;
	v1625 = v1587 + int32(22)
	v1626 = v1590
	v1642 = v1615
	v1646 = v1616
	goto L188
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = int64(16384)
	v1600 = int64(0)
	v1615 = v1600 - v1581
	v1616 = v1600 - (v1580 + base.I64_extend_i32_u(base.B2i32(v1581 != v1600)))
	goto L183
L185:
	;
	goto L186
L186:
	;
	v1608 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = v1608
	if v1581|v1580 == v1608 {
		v1731 = v1590
		v1733 = int32(0)
		goto L182
	} else {
		goto L187
	}
L187:
	;
	v1615 = v1581
	v1616 = v1580
	goto L183
L188:
	;
	v1654 = int32(16)
	v1655 = v36 + v1654
	v1658 = m.G0
	v1660 = v1658 - v1654
	m.G0 = v1660
	F___udivmodti4(m, v1660, v1642, v1646, int64(10000), int64(0))
	mBase = m.M
	v1664 = *(*int64)(unsafe.Add(mBase, uint32(v1660)))
	v1665 = *(*int64)(unsafe.Add(mBase, uint32(v1660)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1655)+8)) = v1665
	*(*int64)(unsafe.Add(mBase, uint32(v1655))) = v1664
	m.G0 = v1660 + v1654
	goto L190
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+332)) = v1712
	v1731 = v1723
	v1733 = v1626
	goto L182
L190:
	;
	v1671 = *(*int64)(unsafe.Add(mBase, uint32(v36)+16))
	v1672 = *(*int64)(unsafe.Add(mBase, uint32(v36+int32(24))))
	v1673 = int64(55536)
	v1674 = int64(0)
	v1679 = int64(32)
	v1682 = int64(base.Ui64(v1671) >> (uint(v1679) % 64))
	v1685 = int64(4294967295)
	v1688 = v1671 & v1685
	v1689 = v1673 * v1688
	v1693 = int64(base.Ui64(v1689)>>(uint(v1679)%64)) + v1673*v1682
	v1700 = v1688*v1674 + v1693&v1685
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v1671*v1674 + v1672*v1673 + v1674*v1682 + int64(base.Ui64(v1693)>>(uint(v1679)%64)) + int64(base.Ui64(v1700)>>(uint(v1679)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v1689&v1685 | v1700<<(uint(v1679)%64)
	goto L191
L191:
	;
	v1712 = v1625 - int32(2)
	v1713 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	v1714 = v1713 + v1642
	*(*uint16)(unsafe.Add(mBase, uint32(v1712))) = uint16(v1714)
	v1718 = int64(0)
	v1723 = v1626 + int32(1)
	if v1646 == v1718 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1724 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v1642))
	goto L194
L193:
	;
	v1724 = base.B2i32(v1646 != v1718)
	goto L194
L194:
	;
	if v1724 != 0 {
		v1625 = v1712
		v1626 = v1723
		v1642 = v1671
		v1646 = v1672
		goto L188
	} else {
		goto L195
	}
L195:
	;
	goto L189
L196:
	;
	if v38 <= v1768 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v1805 = v1770
	v1806 = v1775
	goto L198
L198:
	;
	v1807 = v493 + v526
	if v535 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	v1804 = v1772
	goto L201
L200:
	;
	v1802 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219+v1768<<(uint(int32(1))%32)))))
	v1804 = v1772 + v1802
	goto L201
L201:
	;
	v1805 = v1804
	v1806 = v1797
	goto L198
L202:
	;
	if v533 == int32(0) {
		v1917 = v1882
		goto L215
	} else {
		goto L216
	}
L203:
	;
	v1879 = v1807
	v1882 = v1797
	v1908 = int32(0)
	goto L202
L204:
	;
	goto L205
L205:
	;
	v1813 = int32(0)
	v1819 = v1807
	v1822 = v1813
	v1823 = v1813
	goto L206
L206:
	;
	v1849 = v1822 * int32(10000)
	if v1819 < v38 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v1879 = v1869
	v1882 = v1867
	v1908 = v1867 * int32(10000)
	goto L202
L208:
	;
	v1854 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219+v1819<<(uint(int32(1))%32)))))
	v1856 = v1849 + v1854
	goto L210
L209:
	;
	v1856 = v1849
	goto L210
L210:
	;
	v1858 = v1856 * int32(10000)
	v1860 = v1819 + int32(1)
	if v1860 < v38 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1865 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219+v1860<<(uint(int32(1))%32)))))
	v1867 = v1858 + v1865
	goto L213
L212:
	;
	v1867 = v1858
	goto L213
L213:
	;
	v1868 = int32(2)
	v1869 = v1819 + v1868
	v1871 = v1823 + v1868
	if v1871 != v526&int32(1073741822) {
		v1819 = v1869
		v1822 = v1867
		v1823 = v1871
		goto L206
	} else {
		goto L214
	}
L214:
	;
	goto L207
L215:
	;
	v1928 = v1805
	v1932 = v526 + v1807
	v1933 = v1806
	v1945 = base.I64_extend_i32_s(v1805)
	v1949 = base.I64_extend_i32_s(v1917)
	v1955 = base.I64_extend_i32_s(v1806)
	goto L70
L216:
	;
	if v38 <= v1879 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1917 = v1908
	goto L215
L218:
	;
	goto L219
L219:
	;
	v1915 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219+v1879<<(uint(int32(1))%32)))))
	v1917 = v1908 + v1915
	goto L215
L220:
	;
	v1978 = v1970
	goto L222
L221:
	;
	v1978 = v1968
	goto L222
L222:
	;
	if int32(0) < v495 {
		v489 = v1928
		v493 = v1932
		v494 = v1933
		v495 = v495 - int32(1)
		v504 = v1978
		v506 = v1960
		v507 = v1966 + (v1970+v1968)&(v1966>>(uint(int64(63))%64))
		v510 = v1966
		goto L68
	} else {
		goto L223
	}
L223:
	;
	goto L69
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+328)) = v2018
	v2021 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2018))) = uint16(v2021)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+332)) = v2018 + int32(2)
	if v2005 < int64(0) {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+312)) = v2093
	v2127 = v2093
	v2129 = v2095
	v2132 = v1994
	v2135 = v1997
	v2139 = v478
	goto L63
L226:
	;
	v2046 = v2018 + int32(12)
	v2047 = v2021
	v2063 = v2039
	goto L231
L227:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = int64(16384)
	v2039 = int64(0) - v2005
	goto L226
L228:
	;
	goto L229
L229:
	;
	v2033 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+320)) = v2033
	if v2005 == v2033 {
		v2093 = v2021
		v2095 = int32(0)
		goto L225
	} else {
		goto L230
	}
L230:
	;
	v2039 = v2005
	goto L226
L231:
	;
	v2076 = v2046 - int32(2)
	v2078 = base.I64_div_u_s(v2063, int64(10000))
	v2081 = v2078*int64(55536) + v2063
	*(*uint16)(unsafe.Add(mBase, uint32(v2076))) = uint16(v2081)
	v2084 = v2047 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v2063) {
		v2046 = v2076
		v2047 = v2084
		v2063 = v2078
		goto L231
	} else {
		goto L233
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+332)) = v2076
	v2093 = v2084
	v2095 = v2047
	goto L225
L233:
	;
	goto L232
L234:
	;
	v2170 = v2132
	v2172 = v2156
	v2173 = v2135
	v2174 = v2156
	goto L238
L235:
	;
	v3883 = v2127
	v3890 = v2156
	v3892 = v2156
	goto L236
L236:
	;
	v3912 = v3883 << (uint(int32(1)) % 32)
	v3915 = F_palloc(m, v3912+int32(2))
	mBase = m.M
	v3916 = m.ExcPending
	if v3916 != 0 {
		goto L8
	} else {
		goto L609
	}
L237:
	;
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v36)+312))
	v3883 = v3877
	v3890 = v2412
	v3892 = v2646
	goto L236
L238:
	;
	v2195 = int32(2)
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(336)+v2173<<(uint(v2195)%32))))
	v2201 = base.I32_div_s(v2198-v2170, v2195)
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v2170 < v2202 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(v36)+224))
	v3488 = *(*int32)(unsafe.Add(mBase, uint32(v36)+216))
	v3489 = *(*int32)(unsafe.Add(mBase, uint32(v36)+192))
	if v3489 == int32(0) {
		goto L498
	} else {
		goto L499
	}
L240:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2434 = v2201 + v2170
	if v2434 < v2433 {
		goto L276
	} else {
		goto L277
	}
L241:
	;
	v2204 = v2202 - v2170
	if v2201 < v2204 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	if v2172 != 0 {
		goto L271
	} else {
		goto L272
	}
L244:
	;
	v2206 = v2201
	goto L246
L245:
	;
	v2206 = v2204
	goto L246
L246:
	;
	if v2172 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	F_pfree(m, v2172)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L8
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v2210 = v2206 << (uint(int32(1)) % 32)
	v2213 = F_palloc(m, v2210+int32(2))
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L8
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+256)) = v2213
	v2216 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2213))) = uint16(v2216)
	v2219 = v2213 + int32(2)
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2210 != 0 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+248)) = int64(0)
	v2229 = v2201 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v2229
	if int32(0) < v2206 {
		goto L258
	} else {
		goto L259
	}
L253:
	;
	v2224 = F__emscripten_memcpy_bulkmem(m, v2219, v2220+v2170<<(uint(int32(1))%32), v2210)
	mBase = m.M
	v2225 = v2224
	goto L255
L254:
	;
	v2225 = v2219
	goto L255
L255:
	;
	goto L252
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+240)) = v2361
	*(*int32)(unsafe.Add(mBase, uint32(v36)+260)) = v2360
	v2412 = v2213
	goto L240
L257:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+244)) = int64(0)
	v2360 = v2324
	v2361 = int32(0)
	goto L256
L258:
	;
	v2238 = v2219
	v2239 = v2206
	v2240 = v2229
	goto L262
L259:
	;
	goto L260
L260:
	;
	if v2206 != 0 {
		v2360 = v2219
		v2361 = v2206
		goto L256
	} else {
		goto L270
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+244)) = v2240
	v2284 = v2239
	goto L266
L262:
	;
	v2267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2238))))
	if v2267 != 0 {
		goto L261
	} else {
		goto L264
	}
L263:
	;
	v2324 = v2210 + v2225
	goto L257
L264:
	;
	v2268 = int32(1)
	if v2268 < v2239 {
		v2238 = v2238 + int32(2)
		v2239 = v2239 - v2268
		v2240 = v2240 - v2268
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v2315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2238-int32(2)+v2284<<(uint(int32(1))%32)))))
	if v2315 != 0 {
		v2360 = v2238
		v2361 = v2284
		goto L256
	} else {
		goto L268
	}
L267:
	;
	v2324 = v2238
	goto L257
L268:
	;
	v2316 = int32(1)
	if v2316 < v2284 {
		v2284 = v2284 - v2316
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	v2324 = v2219
	goto L257
L271:
	;
	F_pfree(m, v2172)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L8
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v2393 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+256)) = v2393
	*(*int64)(unsafe.Add(mBase, uint32(v36)+248)) = v2393
	*(*int64)(unsafe.Add(mBase, uint32(v36)+240)) = v2393
	v2412 = int32(0)
	goto L240
L274:
	;
	goto L273
L275:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v36)+288))
	v2667 = v2665 << (uint(int32(1)) % 32)
	v2670 = F_palloc(m, v2667+int32(2))
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L8
	} else {
		goto L310
	}
L276:
	;
	v2436 = v2433 - v2434
	if v2201 < v2436 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	if v2174 != 0 {
		goto L306
	} else {
		goto L307
	}
L279:
	;
	v2438 = v2201
	goto L281
L280:
	;
	v2438 = v2436
	goto L281
L281:
	;
	if v2174 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	F_pfree(m, v2174)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L8
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v2442 = v2438 << (uint(int32(1)) % 32)
	v2445 = F_palloc(m, v2442+int32(2))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L8
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+280)) = v2445
	v2448 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2445))) = uint16(v2448)
	v2451 = v2445 + int32(2)
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2442 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+272)) = int64(0)
	v2461 = v2201 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+268)) = v2461
	if int32(0) < v2438 {
		goto L293
	} else {
		goto L294
	}
L288:
	;
	v2456 = F__emscripten_memcpy_bulkmem(m, v2451, v2452+v2434<<(uint(int32(1))%32), v2442)
	mBase = m.M
	v2457 = v2456
	goto L290
L289:
	;
	v2457 = v2451
	goto L290
L290:
	;
	goto L287
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+264)) = v2593
	*(*int32)(unsafe.Add(mBase, uint32(v36)+284)) = v2592
	v2646 = v2445
	goto L275
L292:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+268)) = int64(0)
	v2592 = v2556
	v2593 = int32(0)
	goto L291
L293:
	;
	v2470 = v2451
	v2471 = v2438
	v2472 = v2461
	goto L297
L294:
	;
	goto L295
L295:
	;
	if v2438 != 0 {
		v2592 = v2451
		v2593 = v2438
		goto L291
	} else {
		goto L305
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+268)) = v2472
	v2516 = v2471
	goto L301
L297:
	;
	v2499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2470))))
	if v2499 != 0 {
		goto L296
	} else {
		goto L299
	}
L298:
	;
	v2556 = v2442 + v2457
	goto L292
L299:
	;
	v2500 = int32(1)
	if v2500 < v2471 {
		v2470 = v2470 + int32(2)
		v2471 = v2471 - v2500
		v2472 = v2472 - v2500
		goto L297
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	v2547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2470-int32(2)+v2516<<(uint(int32(1))%32)))))
	if v2547 != 0 {
		v2592 = v2470
		v2593 = v2516
		goto L291
	} else {
		goto L303
	}
L302:
	;
	v2556 = v2470
	goto L292
L303:
	;
	v2548 = int32(1)
	if v2548 < v2516 {
		v2516 = v2516 - v2548
		goto L301
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v2556 = v2451
	goto L292
L306:
	;
	F_pfree(m, v2174)
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L8
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	v2625 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+280)) = v2625
	*(*int64)(unsafe.Add(mBase, uint32(v36)+272)) = v2625
	*(*int64)(unsafe.Add(mBase, uint32(v36)+264)) = v2625
	v2646 = int32(0)
	goto L275
L309:
	;
	goto L308
L310:
	;
	v2672 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2670))) = uint16(v2672)
	if v2672 < v2665 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v36)+308))
	if v2667 != 0 {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	goto L313
L313:
	;
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(v36)+232))
	if v2681 != 0 {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	goto L313
L315:
	;
	v2679 = F__emscripten_memcpy_bulkmem(m, v2670+int32(2), v2678, v2667)
	mBase = m.M
	goto L317
L316:
	;
	goto L317
L317:
	;
	goto L314
L318:
	;
	F_pfree(m, v2681)
	mBase = m.M
	v2683 = m.ExcPending
	if v2683 != 0 {
		goto L8
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v2685 = v36 + int32(224)
	v2686 = *(*int64)(unsafe.Add(mBase, uint32(v36)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v2685))) = v2686
	v2688 = *(*int64)(unsafe.Add(mBase, uint32(v36)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+216)) = v2688
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v2670
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v2670 + int32(2)
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v36)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+220)) = v2694 + v2201
	v2698 = v36 + int32(216)
	F_add_var(m, v2698, v36+int32(240), v2698)
	mBase = m.M
	v2704 = m.ExcPending
	if v2704 != 0 {
		goto L8
	} else {
		goto L322
	}
L321:
	;
	goto L320
L322:
	;
	v2706 = v36 + int32(312)
	F_add_var(m, v2706, v2706, v36+int32(192))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L8
	} else {
		goto L323
	}
L323:
	;
	v2714 = v36 + int32(496)
	v2715 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2714))) = v2715
	*(*int64)(unsafe.Add(mBase, uint32(v36)+504)) = v2715
	v2720 = v36 + int32(472)
	*(*int64)(unsafe.Add(mBase, uint32(v2720))) = v2715
	*(*int64)(unsafe.Add(mBase, uint32(v36)+480)) = v2715
	*(*int64)(unsafe.Add(mBase, uint32(v36)+488)) = v2715
	*(*int64)(unsafe.Add(mBase, uint32(v36)+464)) = v2715
	v2735 = int32(0)
	F_div_var(m, v36+int32(216), v36+int32(192), v36+int32(488), v2735, v2735, v2735)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L8
	} else {
		goto L324
	}
L324:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v36)+204))
	F_mul_var(m, v36+int32(192), v36+int32(488), v36+int32(464), v2746)
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L8
	} else {
		goto L325
	}
L325:
	;
	v2752 = v36 + int32(464)
	F_sub_var(m, v36+int32(216), v2752, v2752)
	mBase = m.M
	v2756 = m.ExcPending
	if v2756 != 0 {
		goto L8
	} else {
		goto L326
	}
L326:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v36)+464))
	if v2757 == int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v36)+484))
	v2903 = *(*int32)(unsafe.Add(mBase, uint32(v36)+468))
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v36)+212))
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v36)+192))
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v36)+196))
	v2907 = int32(0)
	if base.B2i32(v2906 < v2903)&base.B2i32(v2907 < v2874) == v2907 {
		goto L346
	} else {
		goto L347
	}
L328:
	;
	v2874 = int32(0)
	goto L327
L329:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v36)+224))
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v36)+472))
	if v2760 == v2761 {
		v2874 = v2757
		goto L327
	} else {
		goto L330
	}
L330:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	goto L331
L331:
	;
	if base.B2i32(v2760 != v2763) == int32(0) {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	v2874 = v2830
	goto L327
L333:
	;
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(v36)+464))
	if v2830 == int32(0) {
		goto L328
	} else {
		goto L341
	}
L334:
	;
	v2801 = v36 + int32(488)
	F_sub_var(m, v2801, int32(1737832), v2801)
	mBase = m.M
	v2806 = m.ExcPending
	if v2806 != 0 {
		goto L8
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v2816 = v36 + int32(488)
	F_add_var(m, v2816, int32(1737832), v2816)
	mBase = m.M
	v2821 = m.ExcPending
	if v2821 != 0 {
		goto L8
	} else {
		goto L339
	}
L337:
	;
	v2808 = v36 + int32(464)
	F_add_var(m, v2808, v36+int32(192), v2808)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L8
	} else {
		goto L338
	}
L338:
	;
	goto L333
L339:
	;
	v2823 = v36 + int32(464)
	F_sub_var(m, v2823, v36+int32(192), v2823)
	mBase = m.M
	v2829 = m.ExcPending
	if v2829 != 0 {
		goto L8
	} else {
		goto L340
	}
L340:
	;
	goto L333
L341:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v36)+472))
	if v2760 != v2833 {
		goto L331
	} else {
		goto L342
	}
L342:
	;
	goto L332
L343:
	;
	if int32(0) <= v3078 {
		goto L390
	} else {
		goto L391
	}
L344:
	;
	v3078 = v3068
	goto L343
L345:
	;
	if v2906 <= v2938 {
		v2973 = v2906
		v2975 = v2907
		goto L354
	} else {
		goto L355
	}
L346:
	;
	v2938 = v2903
	v2942 = v2907
	goto L345
L347:
	;
	goto L348
L348:
	;
	v2919 = v2903
	v2923 = v2907
	goto L349
L349:
	;
	v2929 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2902+v2923<<(uint(int32(1))%32)))))
	if v2929 != 0 {
		v3068 = int32(1)
		goto L344
	} else {
		goto L351
	}
L350:
	;
	v2938 = v2933
	v2942 = v2931
	goto L345
L351:
	;
	v2930 = int32(1)
	v2931 = v2923 + v2930
	v2933 = v2919 - v2930
	if v2933 <= v2906 {
		v2938 = v2933
		v2942 = v2931
		goto L345
	} else {
		goto L352
	}
L352:
	;
	if v2931 < v2874 {
		v2919 = v2933
		v2923 = v2931
		goto L349
	} else {
		goto L353
	}
L353:
	;
	goto L350
L354:
	;
	if v2938 != v2973 {
		v3014 = v2942
		v3015 = v2975
		goto L362
	} else {
		goto L363
	}
L355:
	;
	if v2905 <= int32(0) {
		v2973 = v2906
		v2975 = v2907
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v2954 = v2906
	v2956 = v2907
	goto L357
L357:
	;
	v2961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2904+v2956<<(uint(int32(1))%32)))))
	if v2961 != 0 {
		v3068 = int32(-1)
		goto L344
	} else {
		goto L359
	}
L358:
	;
	v2973 = v2965
	v2975 = v2963
	goto L354
L359:
	;
	v2962 = int32(1)
	v2963 = v2956 + v2962
	v2965 = v2954 - v2962
	if v2965 <= v2938 {
		v2973 = v2965
		v2975 = v2963
		goto L354
	} else {
		goto L360
	}
L360:
	;
	if v2963 < v2905 {
		v2954 = v2965
		v2956 = v2963
		goto L357
	} else {
		goto L361
	}
L361:
	;
	goto L358
L362:
	;
	if v2874 < v3014 {
		goto L372
	} else {
		goto L373
	}
L363:
	;
	v2984 = v2942
	v2985 = v2975
	goto L364
L364:
	;
	if v2874 <= v2984 {
		v3014 = v2984
		v3015 = v2985
		goto L362
	} else {
		goto L366
	}
L365:
	;
	if base.I32_extend16_s(v3000) < base.I32_extend16_s(v2998) {
		goto L369
	} else {
		goto L370
	}
L366:
	;
	if v2905 <= v2985 {
		v3014 = v2984
		v3015 = v2985
		goto L362
	} else {
		goto L367
	}
L367:
	;
	v2989 = int32(1)
	v2998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2902+v2984<<(uint(v2989)%32)))))
	v3000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2985<<(uint(v2989)%32)+v2904))))
	if v2998 == v3000 {
		v2984 = v2984 + v2989
		v2985 = v2985 + v2989
		goto L364
	} else {
		goto L368
	}
L368:
	;
	goto L365
L369:
	;
	v3007 = int32(1)
	goto L371
L370:
	;
	v3007 = int32(-1)
	goto L371
L371:
	;
	v3078 = v3007
	goto L343
L372:
	;
	v3018 = v3014
	goto L374
L373:
	;
	v3018 = v2874
	goto L374
L374:
	;
	v3025 = v3014
	goto L375
L375:
	;
	if v3018 == v3025 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	v3068 = v3051
	goto L344
L377:
	;
	if v2905 < v3015 {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	goto L379
L379:
	;
	v3051 = int32(1)
	v3057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2902+v3025<<(uint(v3051)%32)))))
	if v3057 == int32(0) {
		v3025 = v3025 + v3051
		goto L375
	} else {
		goto L389
	}
L380:
	;
	v3030 = v3015
	goto L382
L381:
	;
	v3030 = v2905
	goto L382
L382:
	;
	v3038 = v3015
	goto L383
L383:
	;
	if v3030 == v3038 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v3068 = int32(-1)
	goto L344
L385:
	;
	v3078 = int32(0)
	goto L343
L386:
	;
	goto L387
L387:
	;
	v3042 = int32(1)
	v3047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2904+v3038<<(uint(v3042)%32)))))
	if v3047 == int32(0) {
		v3038 = v3038 + v3042
		goto L383
	} else {
		goto L388
	}
L388:
	;
	goto L384
L389:
	;
	goto L376
L390:
	;
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v36)+224))
	v3082 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	goto L393
L391:
	;
	v3330 = v2902
	v3331 = v2874
	goto L392
L392:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(v36)+488))
	v3361 = v3359 << (uint(int32(1)) % 32)
	v3364 = F_palloc(m, v3361+int32(2))
	mBase = m.M
	v3365 = m.ExcPending
	if v3365 != 0 {
		goto L8
	} else {
		goto L451
	}
L393:
	;
	if base.B2i32(v3081 != v3082) == int32(0) {
		goto L396
	} else {
		goto L397
	}
L394:
	;
	v3330 = v3149
	v3331 = v3150
	goto L392
L395:
	;
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v36)+484))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v36)+464))
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v36)+468))
	v3152 = int32(0)
	if base.B2i32(v2906 < v3151)&base.B2i32(v3152 < v3150) == v3152 {
		goto L406
	} else {
		goto L407
	}
L396:
	;
	v3120 = v36 + int32(488)
	F_add_var(m, v3120, int32(1737832), v3120)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L8
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v3135 = v36 + int32(488)
	F_sub_var(m, v3135, int32(1737832), v3135)
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L8
	} else {
		goto L401
	}
L399:
	;
	v3127 = v36 + int32(464)
	F_sub_var(m, v3127, v36+int32(192), v3127)
	mBase = m.M
	v3133 = m.ExcPending
	if v3133 != 0 {
		goto L8
	} else {
		goto L400
	}
L400:
	;
	goto L395
L401:
	;
	v3142 = v36 + int32(464)
	F_add_var(m, v3142, v36+int32(192), v3142)
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L8
	} else {
		goto L402
	}
L402:
	;
	goto L395
L403:
	;
	if int32(0) <= v3323 {
		goto L393
	} else {
		goto L450
	}
L404:
	;
	v3323 = v3313
	goto L403
L405:
	;
	if v2906 <= v3183 {
		v3218 = v2906
		v3220 = v3152
		goto L414
	} else {
		goto L415
	}
L406:
	;
	v3183 = v3151
	v3187 = v3152
	goto L405
L407:
	;
	goto L408
L408:
	;
	v3164 = v3151
	v3168 = v3152
	goto L409
L409:
	;
	v3174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3149+v3168<<(uint(int32(1))%32)))))
	if v3174 != 0 {
		v3313 = int32(1)
		goto L404
	} else {
		goto L411
	}
L410:
	;
	v3183 = v3178
	v3187 = v3176
	goto L405
L411:
	;
	v3175 = int32(1)
	v3176 = v3168 + v3175
	v3178 = v3164 - v3175
	if v3178 <= v2906 {
		v3183 = v3178
		v3187 = v3176
		goto L405
	} else {
		goto L412
	}
L412:
	;
	if v3176 < v3150 {
		v3164 = v3178
		v3168 = v3176
		goto L409
	} else {
		goto L413
	}
L413:
	;
	goto L410
L414:
	;
	if v3183 != v3218 {
		v3259 = v3187
		v3260 = v3220
		goto L422
	} else {
		goto L423
	}
L415:
	;
	if v2905 <= int32(0) {
		v3218 = v2906
		v3220 = v3152
		goto L414
	} else {
		goto L416
	}
L416:
	;
	v3199 = v2906
	v3201 = v3152
	goto L417
L417:
	;
	v3206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2904+v3201<<(uint(int32(1))%32)))))
	if v3206 != 0 {
		v3313 = int32(-1)
		goto L404
	} else {
		goto L419
	}
L418:
	;
	v3218 = v3210
	v3220 = v3208
	goto L414
L419:
	;
	v3207 = int32(1)
	v3208 = v3201 + v3207
	v3210 = v3199 - v3207
	if v3210 <= v3183 {
		v3218 = v3210
		v3220 = v3208
		goto L414
	} else {
		goto L420
	}
L420:
	;
	if v3208 < v2905 {
		v3199 = v3210
		v3201 = v3208
		goto L417
	} else {
		goto L421
	}
L421:
	;
	goto L418
L422:
	;
	if v3150 < v3259 {
		goto L432
	} else {
		goto L433
	}
L423:
	;
	v3229 = v3187
	v3230 = v3220
	goto L424
L424:
	;
	if v3150 <= v3229 {
		v3259 = v3229
		v3260 = v3230
		goto L422
	} else {
		goto L426
	}
L425:
	;
	if base.I32_extend16_s(v3245) < base.I32_extend16_s(v3243) {
		goto L429
	} else {
		goto L430
	}
L426:
	;
	if v2905 <= v3230 {
		v3259 = v3229
		v3260 = v3230
		goto L422
	} else {
		goto L427
	}
L427:
	;
	v3234 = int32(1)
	v3243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3149+v3229<<(uint(v3234)%32)))))
	v3245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3230<<(uint(v3234)%32)+v2904))))
	if v3243 == v3245 {
		v3229 = v3229 + v3234
		v3230 = v3230 + v3234
		goto L424
	} else {
		goto L428
	}
L428:
	;
	goto L425
L429:
	;
	v3252 = int32(1)
	goto L431
L430:
	;
	v3252 = int32(-1)
	goto L431
L431:
	;
	v3323 = v3252
	goto L403
L432:
	;
	v3263 = v3259
	goto L434
L433:
	;
	v3263 = v3150
	goto L434
L434:
	;
	v3270 = v3259
	goto L435
L435:
	;
	if v3263 == v3270 {
		goto L437
	} else {
		goto L438
	}
L436:
	;
	v3313 = v3296
	goto L404
L437:
	;
	if v2905 < v3260 {
		goto L440
	} else {
		goto L441
	}
L438:
	;
	goto L439
L439:
	;
	v3296 = int32(1)
	v3302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3149+v3270<<(uint(v3296)%32)))))
	if v3302 == int32(0) {
		v3270 = v3270 + v3296
		goto L435
	} else {
		goto L449
	}
L440:
	;
	v3275 = v3260
	goto L442
L441:
	;
	v3275 = v2905
	goto L442
L442:
	;
	v3283 = v3260
	goto L443
L443:
	;
	if v3275 == v3283 {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	v3313 = int32(-1)
	goto L404
L445:
	;
	v3323 = int32(0)
	goto L403
L446:
	;
	goto L447
L447:
	;
	v3287 = int32(1)
	v3292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2904+v3283<<(uint(v3287)%32)))))
	if v3292 == int32(0) {
		v3283 = v3283 + v3287
		goto L443
	} else {
		goto L448
	}
L448:
	;
	goto L444
L449:
	;
	goto L436
L450:
	;
	goto L394
L451:
	;
	v3366 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3364))) = uint16(v3366)
	if v3366 < v3359 {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v3372 = *(*int32)(unsafe.Add(mBase, uint32(v36)+508))
	if v3361 != 0 {
		goto L456
	} else {
		goto L457
	}
L453:
	;
	goto L454
L454:
	;
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(v36)+232))
	if v3375 != 0 {
		goto L459
	} else {
		goto L460
	}
L455:
	;
	goto L454
L456:
	;
	v3373 = F__emscripten_memcpy_bulkmem(m, v3364+int32(2), v3372, v3361)
	mBase = m.M
	goto L458
L457:
	;
	goto L458
L458:
	;
	goto L455
L459:
	;
	F_pfree(m, v3375)
	mBase = m.M
	v3377 = m.ExcPending
	if v3377 != 0 {
		goto L8
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	v3378 = *(*int64)(unsafe.Add(mBase, uint32(v2714)))
	*(*int64)(unsafe.Add(mBase, uint32(v2685))) = v3378
	v3380 = *(*int64)(unsafe.Add(mBase, uint32(v36)+488))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+216)) = v3380
	*(*int32)(unsafe.Add(mBase, uint32(v36)+232)) = v3364
	v3383 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+236)) = v3364 + v3383
	v3387 = v3331 << (uint(int32(1)) % 32)
	v3390 = F_palloc(m, v3387+v3383)
	mBase = m.M
	v3391 = m.ExcPending
	if v3391 != 0 {
		goto L8
	} else {
		goto L463
	}
L462:
	;
	goto L461
L463:
	;
	v3392 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3390))) = uint16(v3392)
	if v3392 < v3331 {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	if v3387 != 0 {
		goto L468
	} else {
		goto L469
	}
L465:
	;
	goto L466
L466:
	;
	v3400 = *(*int32)(unsafe.Add(mBase, uint32(v36)+208))
	if v3400 != 0 {
		goto L471
	} else {
		goto L472
	}
L467:
	;
	goto L466
L468:
	;
	v3398 = F__emscripten_memcpy_bulkmem(m, v3390+int32(2), v3330, v3387)
	mBase = m.M
	goto L470
L469:
	;
	goto L470
L470:
	;
	goto L467
L471:
	;
	F_pfree(m, v3400)
	mBase = m.M
	v3402 = m.ExcPending
	if v3402 != 0 {
		goto L8
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	v3403 = *(*int64)(unsafe.Add(mBase, uint32(v2720)))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+200)) = v3403
	v3405 = *(*int64)(unsafe.Add(mBase, uint32(v36)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+192)) = v3405
	*(*int32)(unsafe.Add(mBase, uint32(v36)+208)) = v3390
	*(*int32)(unsafe.Add(mBase, uint32(v36)+212)) = v3390 + int32(2)
	v3411 = *(*int32)(unsafe.Add(mBase, uint32(v36)+504))
	if v3411 != 0 {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	goto L473
L475:
	;
	F_pfree(m, v3411)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L8
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v3414 = *(*int32)(unsafe.Add(mBase, uint32(v36)+480))
	if v3414 != 0 {
		goto L479
	} else {
		goto L480
	}
L478:
	;
	goto L477
L479:
	;
	F_pfree(m, v3414)
	mBase = m.M
	v3416 = m.ExcPending
	if v3416 != 0 {
		goto L8
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(v36)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+316)) = v3417 + v2201
	v3421 = v36 + int32(312)
	F_add_var(m, v3421, v36+int32(216), v3421)
	mBase = m.M
	v3427 = m.ExcPending
	if v3427 != 0 {
		goto L8
	} else {
		goto L483
	}
L482:
	;
	goto L481
L483:
	;
	v3428 = *(*int32)(unsafe.Add(mBase, uint32(v36)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+196)) = v3428 + v2201
	v3432 = v36 + int32(192)
	F_add_var(m, v3432, v36+int32(264), v3432)
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L8
	} else {
		goto L484
	}
L484:
	;
	v3440 = v36 + int32(216)
	F_mul_var(m, v3440, v3440, v3440, int32(0))
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L8
	} else {
		goto L485
	}
L485:
	;
	if v2173 != 0 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	F_sub_var(m, v36+int32(192), v36+int32(216), v36+int32(288))
	mBase = m.M
	v3455 = m.ExcPending
	if v3455 != 0 {
		goto L8
	} else {
		goto L489
	}
L487:
	;
	goto L488
L488:
	;
	goto L239
L489:
	;
	v3456 = *(*int32)(unsafe.Add(mBase, uint32(v36)+296))
	if v3456 == int32(16384) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v3460 = v36 + int32(288)
	F_add_var(m, v3460, v36+int32(312), v3460)
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L8
	} else {
		goto L493
	}
L491:
	;
	goto L492
L492:
	;
	if int32(0) < v2173 {
		v2170 = v2201 + v2434
		v2172 = v2412
		v2173 = v2173 - int32(1)
		v2174 = v2646
		goto L238
	} else {
		goto L496
	}
L493:
	;
	v3468 = v36 + int32(312)
	F_sub_var(m, v3468, int32(1737832), v3468)
	mBase = m.M
	v3473 = m.ExcPending
	if v3473 != 0 {
		goto L8
	} else {
		goto L494
	}
L494:
	;
	v3475 = v36 + int32(288)
	F_add_var(m, v3475, v36+int32(312), v3475)
	mBase = m.M
	v3481 = m.ExcPending
	if v3481 != 0 {
		goto L8
	} else {
		goto L495
	}
L495:
	;
	goto L492
L496:
	;
	goto L237
L497:
	;
	v3862 = v36 + int32(312)
	F_sub_var(m, v3862, int32(1737832), v3862)
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L8
	} else {
		goto L608
	}
L498:
	;
	if v3488 == int32(0) {
		goto L237
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v36)+200))
	if v3488 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L501:
	;
	if v3487 != int32(16384) {
		goto L497
	} else {
		goto L502
	}
L502:
	;
	goto L237
L503:
	;
	if v3496 != 0 {
		goto L497
	} else {
		goto L506
	}
L504:
	;
	goto L505
L505:
	;
	v3499 = *(*int32)(unsafe.Add(mBase, uint32(v36)+220))
	v3500 = *(*int32)(unsafe.Add(mBase, uint32(v36)+236))
	v3501 = *(*int32)(unsafe.Add(mBase, uint32(v36)+196))
	v3502 = *(*int32)(unsafe.Add(mBase, uint32(v36)+212))
	if v3496 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L506:
	;
	goto L237
L507:
	;
	if int32(0) <= v3853 {
		goto L237
	} else {
		goto L607
	}
L508:
	;
	if v3487 == int32(16384) {
		goto L237
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	if v3487 == int32(0) {
		goto L497
	} else {
		goto L559
	}
L511:
	;
	v3507 = int32(0)
	if base.B2i32(v3499 < v3501)&base.B2i32(v3507 < v3489) == v3507 {
		goto L515
	} else {
		goto L516
	}
L512:
	;
	v3853 = v3678
	goto L507
L513:
	;
	v3678 = v3668
	goto L512
L514:
	;
	if v3499 <= v3538 {
		v3573 = v3499
		v3575 = v3507
		goto L523
	} else {
		goto L524
	}
L515:
	;
	v3538 = v3501
	v3542 = v3507
	goto L514
L516:
	;
	goto L517
L517:
	;
	v3519 = v3501
	v3523 = v3507
	goto L518
L518:
	;
	v3529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3502+v3523<<(uint(int32(1))%32)))))
	if v3529 != 0 {
		v3668 = int32(1)
		goto L513
	} else {
		goto L520
	}
L519:
	;
	v3538 = v3533
	v3542 = v3531
	goto L514
L520:
	;
	v3530 = int32(1)
	v3531 = v3523 + v3530
	v3533 = v3519 - v3530
	if v3533 <= v3499 {
		v3538 = v3533
		v3542 = v3531
		goto L514
	} else {
		goto L521
	}
L521:
	;
	if v3531 < v3489 {
		v3519 = v3533
		v3523 = v3531
		goto L518
	} else {
		goto L522
	}
L522:
	;
	goto L519
L523:
	;
	if v3538 != v3573 {
		v3614 = v3542
		v3615 = v3575
		goto L531
	} else {
		goto L532
	}
L524:
	;
	if v3488 <= int32(0) {
		v3573 = v3499
		v3575 = v3507
		goto L523
	} else {
		goto L525
	}
L525:
	;
	v3554 = v3499
	v3556 = v3507
	goto L526
L526:
	;
	v3561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3500+v3556<<(uint(int32(1))%32)))))
	if v3561 != 0 {
		v3668 = int32(-1)
		goto L513
	} else {
		goto L528
	}
L527:
	;
	v3573 = v3565
	v3575 = v3563
	goto L523
L528:
	;
	v3562 = int32(1)
	v3563 = v3556 + v3562
	v3565 = v3554 - v3562
	if v3565 <= v3538 {
		v3573 = v3565
		v3575 = v3563
		goto L523
	} else {
		goto L529
	}
L529:
	;
	if v3563 < v3488 {
		v3554 = v3565
		v3556 = v3563
		goto L526
	} else {
		goto L530
	}
L530:
	;
	goto L527
L531:
	;
	if v3489 < v3614 {
		goto L541
	} else {
		goto L542
	}
L532:
	;
	v3584 = v3542
	v3585 = v3575
	goto L533
L533:
	;
	if v3489 <= v3584 {
		v3614 = v3584
		v3615 = v3585
		goto L531
	} else {
		goto L535
	}
L534:
	;
	if base.I32_extend16_s(v3600) < base.I32_extend16_s(v3598) {
		goto L538
	} else {
		goto L539
	}
L535:
	;
	if v3488 <= v3585 {
		v3614 = v3584
		v3615 = v3585
		goto L531
	} else {
		goto L536
	}
L536:
	;
	v3589 = int32(1)
	v3598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3502+v3584<<(uint(v3589)%32)))))
	v3600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3585<<(uint(v3589)%32)+v3500))))
	if v3598 == v3600 {
		v3584 = v3584 + v3589
		v3585 = v3585 + v3589
		goto L533
	} else {
		goto L537
	}
L537:
	;
	goto L534
L538:
	;
	v3607 = int32(1)
	goto L540
L539:
	;
	v3607 = int32(-1)
	goto L540
L540:
	;
	v3678 = v3607
	goto L512
L541:
	;
	v3618 = v3614
	goto L543
L542:
	;
	v3618 = v3489
	goto L543
L543:
	;
	v3625 = v3614
	goto L544
L544:
	;
	if v3618 == v3625 {
		goto L546
	} else {
		goto L547
	}
L545:
	;
	v3668 = v3651
	goto L513
L546:
	;
	if v3488 < v3615 {
		goto L549
	} else {
		goto L550
	}
L547:
	;
	goto L548
L548:
	;
	v3651 = int32(1)
	v3657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3502+v3625<<(uint(v3651)%32)))))
	if v3657 == int32(0) {
		v3625 = v3625 + v3651
		goto L544
	} else {
		goto L558
	}
L549:
	;
	v3630 = v3615
	goto L551
L550:
	;
	v3630 = v3488
	goto L551
L551:
	;
	v3638 = v3615
	goto L552
L552:
	;
	if v3630 == v3638 {
		goto L554
	} else {
		goto L555
	}
L553:
	;
	v3668 = int32(-1)
	goto L513
L554:
	;
	v3678 = int32(0)
	goto L512
L555:
	;
	goto L556
L556:
	;
	v3642 = int32(1)
	v3647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3500+v3638<<(uint(v3642)%32)))))
	if v3647 == int32(0) {
		v3638 = v3638 + v3642
		goto L552
	} else {
		goto L557
	}
L557:
	;
	goto L553
L558:
	;
	goto L545
L559:
	;
	v3681 = int32(0)
	if base.B2i32(v3501 < v3499)&base.B2i32(v3681 < v3488) == v3681 {
		goto L563
	} else {
		goto L564
	}
L560:
	;
	v3853 = v3852
	goto L507
L561:
	;
	v3852 = v3842
	goto L560
L562:
	;
	if v3501 <= v3712 {
		v3747 = v3501
		v3749 = v3681
		goto L571
	} else {
		goto L572
	}
L563:
	;
	v3712 = v3499
	v3716 = v3681
	goto L562
L564:
	;
	goto L565
L565:
	;
	v3693 = v3499
	v3697 = v3681
	goto L566
L566:
	;
	v3703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3500+v3697<<(uint(int32(1))%32)))))
	if v3703 != 0 {
		v3842 = int32(1)
		goto L561
	} else {
		goto L568
	}
L567:
	;
	v3712 = v3707
	v3716 = v3705
	goto L562
L568:
	;
	v3704 = int32(1)
	v3705 = v3697 + v3704
	v3707 = v3693 - v3704
	if v3707 <= v3501 {
		v3712 = v3707
		v3716 = v3705
		goto L562
	} else {
		goto L569
	}
L569:
	;
	if v3705 < v3488 {
		v3693 = v3707
		v3697 = v3705
		goto L566
	} else {
		goto L570
	}
L570:
	;
	goto L567
L571:
	;
	if v3712 != v3747 {
		v3788 = v3716
		v3789 = v3749
		goto L579
	} else {
		goto L580
	}
L572:
	;
	if v3489 <= int32(0) {
		v3747 = v3501
		v3749 = v3681
		goto L571
	} else {
		goto L573
	}
L573:
	;
	v3728 = v3501
	v3730 = v3681
	goto L574
L574:
	;
	v3735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3502+v3730<<(uint(int32(1))%32)))))
	if v3735 != 0 {
		v3842 = int32(-1)
		goto L561
	} else {
		goto L576
	}
L575:
	;
	v3747 = v3739
	v3749 = v3737
	goto L571
L576:
	;
	v3736 = int32(1)
	v3737 = v3730 + v3736
	v3739 = v3728 - v3736
	if v3739 <= v3712 {
		v3747 = v3739
		v3749 = v3737
		goto L571
	} else {
		goto L577
	}
L577:
	;
	if v3737 < v3489 {
		v3728 = v3739
		v3730 = v3737
		goto L574
	} else {
		goto L578
	}
L578:
	;
	goto L575
L579:
	;
	if v3488 < v3788 {
		goto L589
	} else {
		goto L590
	}
L580:
	;
	v3758 = v3716
	v3759 = v3749
	goto L581
L581:
	;
	if v3488 <= v3758 {
		v3788 = v3758
		v3789 = v3759
		goto L579
	} else {
		goto L583
	}
L582:
	;
	if base.I32_extend16_s(v3774) < base.I32_extend16_s(v3772) {
		goto L586
	} else {
		goto L587
	}
L583:
	;
	if v3489 <= v3759 {
		v3788 = v3758
		v3789 = v3759
		goto L579
	} else {
		goto L584
	}
L584:
	;
	v3763 = int32(1)
	v3772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3500+v3758<<(uint(v3763)%32)))))
	v3774 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3759<<(uint(v3763)%32)+v3502))))
	if v3772 == v3774 {
		v3758 = v3758 + v3763
		v3759 = v3759 + v3763
		goto L581
	} else {
		goto L585
	}
L585:
	;
	goto L582
L586:
	;
	v3781 = int32(1)
	goto L588
L587:
	;
	v3781 = int32(-1)
	goto L588
L588:
	;
	v3852 = v3781
	goto L560
L589:
	;
	v3792 = v3788
	goto L591
L590:
	;
	v3792 = v3488
	goto L591
L591:
	;
	v3799 = v3788
	goto L592
L592:
	;
	if v3792 == v3799 {
		goto L594
	} else {
		goto L595
	}
L593:
	;
	v3842 = v3825
	goto L561
L594:
	;
	if v3489 < v3789 {
		goto L597
	} else {
		goto L598
	}
L595:
	;
	goto L596
L596:
	;
	v3825 = int32(1)
	v3831 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3500+v3799<<(uint(v3825)%32)))))
	if v3831 == int32(0) {
		v3799 = v3799 + v3825
		goto L592
	} else {
		goto L606
	}
L597:
	;
	v3804 = v3789
	goto L599
L598:
	;
	v3804 = v3489
	goto L599
L599:
	;
	v3812 = v3789
	goto L600
L600:
	;
	if v3804 == v3812 {
		goto L602
	} else {
		goto L603
	}
L601:
	;
	v3842 = int32(-1)
	goto L561
L602:
	;
	v3852 = int32(0)
	goto L560
L603:
	;
	goto L604
L604:
	;
	v3816 = int32(1)
	v3821 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3502+v3812<<(uint(v3816)%32)))))
	if v3821 == int32(0) {
		v3812 = v3812 + v3816
		goto L600
	} else {
		goto L605
	}
L605:
	;
	goto L601
L606:
	;
	goto L593
L607:
	;
	goto L497
L608:
	;
	goto L237
L609:
	;
	v3917 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3915))) = uint16(v3917)
	if v3917 < v3883 {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v3923 = *(*int32)(unsafe.Add(mBase, uint32(v36)+332))
	if v3912 != 0 {
		goto L614
	} else {
		goto L615
	}
L611:
	;
	goto L612
L612:
	;
	v3926 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3926 != 0 {
		goto L617
	} else {
		goto L618
	}
L613:
	;
	goto L612
L614:
	;
	v3924 = F__emscripten_memcpy_bulkmem(m, v3915+int32(2), v3923, v3912)
	mBase = m.M
	goto L616
L615:
	;
	goto L616
L616:
	;
	goto L613
L617:
	;
	F_pfree(m, v3926)
	mBase = m.M
	v3928 = m.ExcPending
	if v3928 != 0 {
		goto L8
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	v3929 = *(*int64)(unsafe.Add(mBase, uint32(v36)+312))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v3929
	v3932 = l1 + int32(8)
	v3933 = *(*int64)(unsafe.Add(mBase, uint32(v36)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v3932))) = v3933
	v3935 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v3915 + v3935
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v3915
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v94
	v3940 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3932))) = v3940
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	v3948 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v3951 = l2 + v3948<<(uint(v3935)%32)
	if v3951+int32(4) < v3940 {
		goto L622
	} else {
		goto L623
	}
L620:
	;
	goto L619
L621:
	;
	v4067 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v4068 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v4068 {
		goto L650
	} else {
		goto L651
	}
L622:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L621
L623:
	;
	goto L624
L624:
	;
	v3960 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3962 = l2 & int32(3)
	v3966 = base.I32_div_s(v3951+int32(7), int32(4))
	v3967 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v3967 <= v3966 {
		goto L629
	} else {
		goto L630
	}
L625:
	;
	goto L621
L626:
	;
	if int32(0) <= v4033 {
		goto L625
	} else {
		goto L647
	}
L627:
	;
	v4013 = v4007
	goto L641
L628:
	;
	v3980 = int32(1)
	v3981 = v3966 - v3980
	v3984 = v3960 + v3981<<(uint(v3980)%32)
	v3985 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3984))))
	v3986 = int32(2)
	v3990 = *(*int32)(unsafe.Add(mBase, uint32(v3962<<(uint(v3986)%32))+uint32(_consts[802])))
	v3991 = base.I32_rem_s(v3985, v3990)
	v3992 = v3985 - v3991
	*(*uint16)(unsafe.Add(mBase, uint32(v3984))) = uint16(v3992)
	v3995 = base.I32_div_s(v3990, v3986)
	if v3991 < v3995 {
		v4033 = v3981
		goto L626
	} else {
		goto L636
	}
L629:
	;
	if v3962 == int32(0) {
		goto L625
	} else {
		goto L632
	}
L630:
	;
	goto L631
L631:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3966
	if v3962 != 0 {
		goto L628
	} else {
		goto L634
	}
L632:
	;
	if v3966 != v3967 {
		goto L625
	} else {
		goto L633
	}
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3966
	goto L628
L634:
	;
	v3977 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3960+v3966<<(uint(int32(1))%32)))))
	if v3977 <= int32(4999) {
		v4033 = v3966
		goto L626
	} else {
		goto L635
	}
L635:
	;
	v4007 = v3966
	goto L627
L636:
	;
	v3998 = v3990 + base.I32_extend16_s(v3992)
	if int32(9999) < v3998 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v4003 = v3998 + int32(55536)
	goto L639
L638:
	;
	v4003 = v3998
	goto L639
L639:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3984))) = uint16(v4003)
	if v3998 < int32(10000) {
		v4033 = v3981
		goto L626
	} else {
		goto L640
	}
L640:
	;
	v4007 = v3981
	goto L627
L641:
	;
	v4019 = int32(1)
	v4020 = v4013 - v4019
	v4023 = v3960 + v4020<<(uint(v4019)%32)
	v4026 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4023))))
	v4028 = base.B2i32(int32(9998) < v4026)
	if int32(9998) < v4026 {
		goto L643
	} else {
		goto L644
	}
L642:
	;
	v4033 = v4020
	goto L626
L643:
	;
	v4029 = int32(-9999)
	goto L645
L644:
	;
	v4029 = v4019
	goto L645
L645:
	;
	v4030 = v4029 + v4026
	*(*uint16)(unsafe.Add(mBase, uint32(v4023))) = uint16(v4030)
	if int32(9998) < v4026 {
		v4013 = v4020
		goto L641
	} else {
		goto L646
	}
L646:
	;
	goto L642
L647:
	;
	v4041 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4041 - int32(2)
	v4045 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4046 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4045 + v4046
	v4049 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v4049 + v4046
	goto L625
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4202
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4201
	v4232 = *(*int32)(unsafe.Add(mBase, uint32(v36)+328))
	if v4232 != 0 {
		goto L663
	} else {
		goto L664
	}
L649:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v4201 = v4165
	v4202 = int32(0)
	goto L648
L650:
	;
	v4078 = v4067
	v4079 = v4068
	goto L654
L651:
	;
	goto L652
L652:
	;
	if v4068 != 0 {
		v4201 = v4067
		v4202 = v4068
		goto L648
	} else {
		goto L662
	}
L653:
	;
	v4125 = v4079
	goto L658
L654:
	;
	v4107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4078))))
	if v4107 != 0 {
		goto L653
	} else {
		goto L656
	}
L655:
	;
	v4165 = v4067 + v4068<<(uint(int32(1))%32)
	goto L649
L656:
	;
	v4108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v4109 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v4108 - v4109
	if v4109 < v4079 {
		v4078 = v4078 + int32(2)
		v4079 = v4079 - v4109
		goto L654
	} else {
		goto L657
	}
L657:
	;
	goto L655
L658:
	;
	v4156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4078-int32(2)+v4125<<(uint(int32(1))%32)))))
	if v4156 != 0 {
		v4201 = v4078
		v4202 = v4125
		goto L648
	} else {
		goto L660
	}
L659:
	;
	v4165 = v4078
	goto L649
L660:
	;
	v4157 = int32(1)
	if v4157 < v4125 {
		v4125 = v4125 - v4157
		goto L658
	} else {
		goto L661
	}
L661:
	;
	goto L659
L662:
	;
	v4165 = v4067
	goto L649
L663:
	;
	F_pfree(m, v4232)
	mBase = m.M
	v4234 = m.ExcPending
	if v4234 != 0 {
		goto L8
	} else {
		goto L666
	}
L664:
	;
	goto L665
L665:
	;
	v4235 = *(*int32)(unsafe.Add(mBase, uint32(v36)+304))
	if v4235 != 0 {
		goto L667
	} else {
		goto L668
	}
L666:
	;
	goto L665
L667:
	;
	F_pfree(m, v4235)
	mBase = m.M
	v4237 = m.ExcPending
	if v4237 != 0 {
		goto L8
	} else {
		goto L670
	}
L668:
	;
	goto L669
L669:
	;
	if v3892 != 0 {
		goto L671
	} else {
		goto L672
	}
L670:
	;
	goto L669
L671:
	;
	F_pfree(m, v3892)
	mBase = m.M
	v4239 = m.ExcPending
	if v4239 != 0 {
		goto L8
	} else {
		goto L674
	}
L672:
	;
	goto L673
L673:
	;
	if v3890 != 0 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	goto L673
L675:
	;
	F_pfree(m, v3890)
	mBase = m.M
	v4241 = m.ExcPending
	if v4241 != 0 {
		goto L8
	} else {
		goto L678
	}
L676:
	;
	goto L677
L677:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, uint32(v36)+232))
	if v4242 != 0 {
		goto L679
	} else {
		goto L680
	}
L678:
	;
	goto L677
L679:
	;
	F_pfree(m, v4242)
	mBase = m.M
	v4244 = m.ExcPending
	if v4244 != 0 {
		goto L8
	} else {
		goto L682
	}
L680:
	;
	goto L681
L681:
	;
	v4245 = *(*int32)(unsafe.Add(mBase, uint32(v36)+208))
	if v4245 == int32(0) {
		goto L1
	} else {
		goto L683
	}
L682:
	;
	goto L681
L683:
	;
	F_pfree(m, v4245)
	mBase = m.M
	v4249 = m.ExcPending
	if v4249 != 0 {
		goto L8
	} else {
		goto L684
	}
L684:
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
	v3 = *(*int64)(unsafe.Add(mBase, _consts[95]))
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v97 = v4
		m.G0 = v10 + int32(16)
		return v97
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(683677)
					F_errmsg(m, int32(251825), v10)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(574636), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498003), int32(1783), int32(238792))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
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
			v16 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
				if v20 == int32(1) {
					v23 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						if v25 == int32(0) {
							v97 = v23
						} else {
							v29 = v25
							v30 = int32(0)
							v31 = v23
							for {
								if v30 != 0 {
									v36 = int32(255)
									v37 = v29 & v36
									if base.Ui32((v37-int32(65))&v36) < base.Ui32(int32(26)) {
										v46 = v37 | int32(32)
									} else {
										v46 = v37
									}
									v60 = v46
								} else {
									v47 = int32(255)
									v48 = v29 & v47
									if base.Ui32((v48-int32(97))&v47) < base.Ui32(int32(26)) {
										v57 = v48 - int32(32)
									} else {
										v57 = v48
									}
									v60 = v57 & int32(255)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v60)
								if base.Ui32((v60-int32(48))&int32(255)) < base.Ui32(int32(10)) {
									v77 = int32(1)
								} else {
									v77 = base.B2i32(base.Ui32((v60&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)))
								}
								v79 = v31 + int32(1)
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
								if v80 != 0 {
									v29 = v80
									v30 = v77
									v31 = v79
									continue
								} else {
									break
								}
								break
							}
							v97 = v23
						}
						m.G0 = v10 + int32(16)
						return v97
					}
				} else {
					v82 = l1 + int32(1)
					v83 = F_palloc(m, v82)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v85 = F_pg_strtitle(m, v83, v82, l0, l1, v16)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v88 = v85 + int32(1)
							if base.Ui32(v88) <= base.Ui32(v82) {
								v97 = v83
								m.G0 = v10 + int32(16)
								return v97
							} else {
								v90 = F_repalloc(m, v83, v88)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = F_pg_strtitle(m, v90, v88, l0, l1, v16)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										v97 = v90
										m.G0 = v10 + int32(16)
										return v97
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
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 float64
	_ = v101
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			v20 = v15 + v19
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v23 = v21 & v19
			if v21 == v19 {
				v26 = int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
				if v28&int32(254) == int32(2) {
					v37 = v26
				} else {
					v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
				}
				if v28 == int32(1) {
					v40 = v26
				} else {
					v40 = v37
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23 != 0 {
				v52 = v20
			} else {
				v52 = v15 + int32(4)
			}
			v53 = int32(1)
			v54 = v10 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v10 + int32(4)
			}
			if v57 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v88 = v77
			} else {
				v78 = int32(1)
				if v59 != 0 {
					v88 = int32(base.Ui32(v57)>>(uint(v78)%32)) - v78
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_calc_word_similarity(m, v52, v51, v60, v88, int32(3))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v15 {
							F_pfree(m, v15)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v101 = *(*float64)(unsafe.Add(mBase, _consts[1101]))
								return base.F64_le(v101, base.F64_promote_f32(v90))
							}
						} else {
							v101 = *(*float64)(unsafe.Add(mBase, _consts[1101]))
							return base.F64_le(v101, base.F64_promote_f32(v90))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v101 = *(*float64)(unsafe.Add(mBase, _consts[1101]))
							return base.F64_le(v101, base.F64_promote_f32(v90))
						}
					} else {
						v101 = *(*float64)(unsafe.Add(mBase, _consts[1101]))
						return base.F64_le(v101, base.F64_promote_f32(v90))
					}
				}
			}
		}
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v23 = F_AllocSetContextCreateInternal(m, v18, int32(25132), v2, int32(8192), int32(8388608))
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
	v27 = int32(4515248)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	v98 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v98
	v109 = F_construct_md_array(m, v91, v92, v98, v13+int32(12), v13+int32(8), int32(25), int32(-1), int32(0), int32(105))
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
	v91 = v34
	v92 = v37
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
		v91 = v42
		v92 = v45
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v52 = v2
	v53 = v2
	goto L12
L12:
	;
	v60 = v53 + v45
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v52<<(uint(int32(2))%32))))
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v91 = v42
	v92 = v45
	goto L3
L14:
	;
	v83 = v52 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v83 < v84 {
		v52 = v83
		v53 = v79
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
	v79 = v53
	goto L14
L18:
	;
	v71 = v53 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v42+v53<<(uint(int32(2))%32)))) = v68
	v79 = v71
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v30 = l0
	goto L8
L8:
	;
	if l3 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26+l1))) = uint8(v28)
	v30 = v26
	goto L8
L10:
	;
	v25 = F__emscripten_memcpy_bulkmem(m, v22, l0, l1)
	mBase = m.M
	v26 = v25
	goto L12
L11:
	;
	v26 = v22
	goto L12
L12:
	;
	goto L9
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v41 = l2
	goto L13
L15:
	;
	goto L16
L16:
	;
	v35 = l1 + v22 + int32(1)
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37+l3))) = uint8(v39)
	v41 = v35
	goto L13
L18:
	;
	v36 = F__emscripten_memcpy_bulkmem(m, v35, l2, l3)
	mBase = m.M
	v37 = v36
	goto L20
L19:
	;
	v37 = v35
	goto L20
L20:
	;
	goto L17
L21:
	;
	if v22 != v11 {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	goto L21
L23:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v50 = v30
	v51 = v41
	goto L25
L25:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v65 = v54
	v66 = v55
	goto L22
L27:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	F_pfree(m, v22)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	m.G0 = v11 + int32(1024)
	return v66 - v65
L32:
	;
	goto L31
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
	var v79 int32
	_ = v79
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
		v79 = l0
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return v79 - l0
L14:
	;
	v58 = l0
	v59 = v54
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v59)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v67)>>(uint(v59)%32))&int32(1) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v79 = v75
	goto L13
L17:
	;
	v79 = v58
	goto L13
L18:
	;
	goto L19
L19:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v75 = v58 + int32(1)
	if v73 != 0 {
		v58 = v75
		v59 = v73
		goto L15
	} else {
		goto L20
	}
L20:
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(1040)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v16 = F_pg_snprintf(m, v8+int32(16), int32(1024), int32(125562), v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[508]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
		v25 = F_BufFileOpenFileSet(m, v20, v8+int32(16), int32(0), int32(1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			if v25 != 0 {
				F_BufFileReadExact(m, v25, int32(4425436), int32(4))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[537]))
					v36 = int32(1073741823)
					if v36 <= v34 {
						v39 = v36
					} else {
						v39 = v34
					}
					if base.Ui32(int32(2)) <= base.Ui32(v39) {
						v47 = int32(32) - base.I32_clz(v39-int32(1))
					} else {
						v47 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, _consts[538])) = int32(1) << (uint(v47) % 32)
					v50 = int32(4515248)
					v51 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v54 = *(*int32)(unsafe.Add(mBase, _consts[539]))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v54
					v58 = F_palloc(m, int32(16)<<(uint(v47)%32))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
						*(*int32)(unsafe.Add(mBase, _consts[540])) = v58
						v65 = v34 << (uint(int32(4)) % 32)
						if v65 != 0 {
							F_BufFileReadExact(m, v25, v58, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_BufFileClose(m, v25)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v8 + int32(1040)
									return
								}
							}
						} else {
							F_BufFileClose(m, v25)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[967]))
	if v7 == int32(0) {
		if v5 == int32(10) {
			v15 = int32(1)
			v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
			if v17&v15 == int32(0) {
				v54 = v15
				return v54 & int32(1)
			} else {
				v25 = F_SearchSysCache1(m, int32(11), v5)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v25 != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+68)))
						F_ReleaseCatCache(m, v25)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = v32
							v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
							if v37 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v46 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
									*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
									v52 = v35 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
									v54 = v35
									return v54 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
								v52 = v35 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
								v54 = v35
								return v54 & int32(1)
							}
						}
					} else {
						v35 = int32(0)
						v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
						if v37 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v46 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
								*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
								v52 = v35 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
								v54 = v35
								return v54 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
							v52 = v35 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
							v54 = v35
							return v54 & int32(1)
						}
					}
				}
			}
		} else {
			v25 = F_SearchSysCache1(m, int32(11), v5)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+68)))
					F_ReleaseCatCache(m, v25)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v32
						v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
						if v37 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v46 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
								*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
								v52 = v35 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
								v54 = v35
								return v54 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
							v52 = v35 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
							v54 = v35
							return v54 & int32(1)
						}
					}
				} else {
					v35 = int32(0)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
					if v37 == int32(0) {
						F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v46 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
							*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
							v52 = v35 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
							v54 = v35
							return v54 & int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
						v52 = v35 & int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
						v54 = v35
						return v54 & int32(1)
					}
				}
			}
		}
	} else {
		if v7 != v5 {
			if v5 == int32(10) {
				v15 = int32(1)
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
				if v17&v15 == int32(0) {
					v54 = v15
					return v54 & int32(1)
				} else {
					v25 = F_SearchSysCache1(m, int32(11), v5)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v25 != 0 {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
							v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+68)))
							F_ReleaseCatCache(m, v25)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = v32
								v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
								if v37 == int32(0) {
									F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v46 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
										*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
										v52 = v35 & int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
										v54 = v35
										return v54 & int32(1)
									}
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
									v52 = v35 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
									v54 = v35
									return v54 & int32(1)
								}
							}
						} else {
							v35 = int32(0)
							v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
							if v37 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v46 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
									*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
									v52 = v35 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
									v54 = v35
									return v54 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
								v52 = v35 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
								v54 = v35
								return v54 & int32(1)
							}
						}
					}
				}
			} else {
				v25 = F_SearchSysCache1(m, int32(11), v5)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v25 != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+68)))
						F_ReleaseCatCache(m, v25)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = v32
							v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
							if v37 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v46 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
									*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
									v52 = v35 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
									v54 = v35
									return v54 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
								v52 = v35 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
								v54 = v35
								return v54 & int32(1)
							}
						}
					} else {
						v35 = int32(0)
						v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[968])))
						if v37 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1784), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v46 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[968])) = uint8(v46)
								*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
								v52 = v35 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
								v54 = v35
								return v54 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[967])) = v5
							v52 = v35 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[969])) = uint8(v52)
							v54 = v35
							return v54 & int32(1)
						}
					}
				}
			}
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[969])))
			v54 = v12
			return v54 & int32(1)
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
	var v330 int32
	_ = v330
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v716 int32
	_ = v716
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
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
	v155 = v140&int32(63) | (v98<<(uint(int32(18))%32)&int32(1835008) | v107<<(uint(int32(12))%32) | v123<<(uint(int32(6))%32))
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
	v155 = v98<<(uint(int32(12))%32)&int32(61440) | v107<<(uint(int32(6))%32) | v123
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
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v160)>>(uint(int32(3))%32)))+uint32(_consts[1070]))))
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
	v274 = v259&int32(63) | (v217<<(uint(int32(18))%32)&int32(1835008) | v226<<(uint(int32(12))%32) | v242<<(uint(int32(6))%32))
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
	v274 = v217<<(uint(int32(12))%32)&int32(61440) | v226<<(uint(int32(6))%32) | v242
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
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v279)>>(uint(int32(3))%32)))+uint32(_consts[1070]))))
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
	return v778
L78:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v494
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	if v497 <= v494 {
		goto L118
	} else {
		goto L119
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
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326+v318-int32(1)))))
	if v330&int32(224) != int32(96) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	if int32(1)<<(uint(v330)%32)&int32(1851442) == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v343 = F_find_among_b(m, l0, int32(4375840), int32(37))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	return int32(0)
L85:
	;
	if v343 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v5
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v350
	switch v343 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	default:
		goto L78
	}
L87:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L93
L88:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L84
	} else {
		goto L89
	}
L89:
	;
	if int32(0) <= v354 {
		goto L78
	} else {
		goto L90
	}
L90:
	;
	v778 = v354
	goto L77
L91:
	;
	if v486 != 0 {
		goto L78
	} else {
		goto L115
	}
L92:
	;
	v486 = v479
	goto L91
L93:
	;
	if v374 <= v375 {
		v479 = int32(-1)
		goto L92
	} else {
		goto L95
	}
L94:
	;
	v479 = int32(0)
	goto L92
L95:
	;
	v392 = int32(1)
	v393 = v374 - v392
	v395 = int32(*(*int8)(unsafe.Add(mBase, uint32(v371+v393))))
	v397 = v395 & int32(255)
	if v393 == v375 {
		v452 = v397
		v453 = v392
		goto L96
	} else {
		goto L97
	}
L96:
	;
	if int32(121) < v452 {
		goto L105
	} else {
		goto L106
	}
L97:
	;
	if int32(0) <= v395 {
		v452 = v397
		v453 = v392
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v403 = v397 & int32(63)
	v405 = v374 - int32(2)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v405))))
	v409 = v407 << (uint(int32(6)) % 32)
	if base.B2i32(v405 != v375)&base.B2i32(base.Ui32(v407) < base.Ui32(int32(192))) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v452 = v409&int32(1984) | v403
	v453 = int32(2)
	goto L96
L100:
	;
	goto L101
L101:
	;
	v422 = v409&int32(4032) | v403
	v424 = v374 - int32(3)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v424))))
	if base.B2i32(v424 != v375)&base.B2i32(base.Ui32(v426) < base.Ui32(int32(224))) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v452 = v426<<(uint(int32(12))%32)&int32(61440) | v422
	v453 = int32(3)
	goto L96
L103:
	;
	goto L104
L104:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374+(v371-int32(4))))))
	v452 = v426<<(uint(int32(12))%32)&int32(258048) | v444&int32(7)<<(uint(int32(18))%32) | v422
	v453 = int32(4)
	goto L96
L105:
	;
	v486 = v453
	goto L91
L106:
	;
	goto L107
L107:
	;
	v457 = v452 - int32(98)
	if v457 < int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v486 = v453
	goto L91
L109:
	;
	goto L110
L110:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(3))%32)))+uint32(_consts[1071]))))
	if int32(base.Ui32(v463)>>(uint(v457&int32(7))%32))&int32(1) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v486 = v453
	goto L91
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v374 - v453
	goto L114
L114:
	;
	goto L94
L115:
	;
	v487 = F_slice_del(m, l0)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L84
	} else {
		goto L116
	}
L116:
	;
	if int32(0) <= v487 {
		goto L78
	} else {
		goto L117
	}
L117:
	;
	v778 = v487
	goto L77
L118:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v497
	v502 = v494 - int32(1)
	if v502 <= v497 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v591 = v494
	v592 = v496
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v591
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v592)+4))
	if v591 < v594 {
		goto L150
	} else {
		goto L151
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v499
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v591 = v590
	v592 = v589
	goto L120
L122:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504+v502))))
	if v506&int32(224) != int32(96) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	if int32(1)<<(uint(v506)%32)&int32(1064976) == int32(0) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v519 = F_find_among_b(m, l0, int32(4376592), int32(7))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L84
	} else {
		goto L125
	}
L125:
	;
	if v519 == int32(0) {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v523
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L129
L127:
	;
	if v578 < int32(0) {
		goto L121
	} else {
		goto L147
	}
L129:
	;
	goto L130
L130:
	;
	goto L131
L131:
	;
	v534 = v523
	v536 = int32(1)
	goto L134
L133:
	;
	v578 = v560
	goto L127
L134:
	;
	if v534 <= v527 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L133
L136:
	;
	v578 = int32(-1)
	goto L127
L137:
	;
	goto L138
L138:
	;
	v541 = v534 - int32(1)
	v543 = int32(*(*int8)(unsafe.Add(mBase, uint32(v526+v541))))
	if int32(0) <= v543 {
		v560 = v541
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v564 = int32(1)
	if v564 < v536 {
		v534 = v560
		v536 = v536 - v564
		goto L134
	} else {
		goto L146
	}
L140:
	;
	if v541 <= v527 {
		v560 = v541
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v548 = v541
	goto L142
L142:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526+v548))))
	if base.Ui32(int32(191)) < base.Ui32(v553) {
		v560 = v548
		goto L139
	} else {
		goto L144
	}
L143:
	;
	v560 = v527
	goto L139
L144:
	;
	v557 = v548 - int32(1)
	if v527 < v557 {
		v548 = v557
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	goto L135
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v578
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v578
	v583 = F_slice_del(m, l0)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L84
	} else {
		goto L148
	}
L148:
	;
	if v583 < int32(0) {
		v778 = v583
		goto L77
	} else {
		goto L149
	}
L149:
	;
	goto L121
L150:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v775
	v778 = int32(1)
	goto L77
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v591
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v594
	v600 = v591 - int32(1)
	if v600 <= v594 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v597
	goto L150
L153:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602+v600))))
	if v604&int32(224) != int32(96) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	if int32(1)<<(uint(v604)%32)&int32(1572992) == int32(0) {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v617 = F_find_among_b(m, l0, int32(4376736), int32(5))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L84
	} else {
		goto L156
	}
L156:
	;
	if v617 == int32(0) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v597
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v622
	switch v617 - int32(1) {
	case 0:
		goto L160
	case 1:
		goto L159
	case 2:
		goto L158
	default:
		goto L150
	}
L158:
	;
	v767 = F_slice_from_s(m, l0, int32(4), int32(2224780))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L84
	} else {
		goto L190
	}
L159:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L165
L160:
	;
	v626 = F_slice_del(m, l0)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L84
	} else {
		goto L161
	}
L161:
	;
	if int32(0) <= v626 {
		goto L150
	} else {
		goto L162
	}
L162:
	;
	v778 = v626
	goto L77
L163:
	;
	if v758 != 0 {
		goto L150
	} else {
		goto L187
	}
L164:
	;
	v758 = v751
	goto L163
L165:
	;
	if v646 <= v647 {
		v751 = int32(-1)
		goto L164
	} else {
		goto L167
	}
L166:
	;
	v751 = int32(0)
	goto L164
L167:
	;
	v664 = int32(1)
	v665 = v646 - v664
	v667 = int32(*(*int8)(unsafe.Add(mBase, uint32(v643+v665))))
	v669 = v667 & int32(255)
	if v665 == v647 {
		v724 = v669
		v725 = v664
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if int32(118) < v724 {
		goto L177
	} else {
		goto L178
	}
L169:
	;
	if int32(0) <= v667 {
		v724 = v669
		v725 = v664
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v675 = v669 & int32(63)
	v677 = v646 - int32(2)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643+v677))))
	v681 = v679 << (uint(int32(6)) % 32)
	if base.B2i32(v677 != v647)&base.B2i32(base.Ui32(v679) < base.Ui32(int32(192))) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v724 = v681&int32(1984) | v675
	v725 = int32(2)
	goto L168
L172:
	;
	goto L173
L173:
	;
	v694 = v681&int32(4032) | v675
	v696 = v646 - int32(3)
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643+v696))))
	if base.B2i32(v696 != v647)&base.B2i32(base.Ui32(v698) < base.Ui32(int32(224))) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v724 = v698<<(uint(int32(12))%32)&int32(61440) | v694
	v725 = int32(3)
	goto L168
L175:
	;
	goto L176
L176:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646+(v643-int32(4))))))
	v724 = v698<<(uint(int32(12))%32)&int32(258048) | v716&int32(7)<<(uint(int32(18))%32) | v694
	v725 = int32(4)
	goto L168
L177:
	;
	v758 = v725
	goto L163
L178:
	;
	goto L179
L179:
	;
	v729 = v724 - int32(105)
	if v729 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v758 = v725
	goto L163
L181:
	;
	goto L182
L182:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v729)>>(uint(int32(3))%32)))+uint32(_consts[1072]))))
	if int32(base.Ui32(v735)>>(uint(v729&int32(7))%32))&int32(1) == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v758 = v725
	goto L163
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v646 - v725
	goto L186
L186:
	;
	goto L166
L187:
	;
	v761 = F_slice_from_s(m, l0, int32(3), int32(2224777))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L84
	} else {
		goto L188
	}
L188:
	;
	if int32(0) <= v761 {
		goto L150
	} else {
		goto L189
	}
L189:
	;
	v778 = v761
	goto L77
L190:
	;
	if int32(0) <= v767 {
		goto L150
	} else {
		goto L191
	}
L191:
	;
	v778 = v767
	goto L77
}
