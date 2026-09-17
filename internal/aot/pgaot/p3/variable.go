package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expandRecordVariable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
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
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v15 = l1
	v19 = int32(0)
	goto L4
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L23
	} else {
		goto L116
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L23
	} else {
		goto L113
	}
L3:
	;
	m.G0 = v12 + int32(160)
	return v462
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v25 = v24 + v19
	v26 = int32(0)
	if v25 <= v26 {
		v73 = l0
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v460 = F_get_expr_result_tupdesc(m, v456, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L23
	} else {
		goto L112
	}
L6:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+8)))
	if v87 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+v23<<(uint(int32(2))%32)-int32(4))))
	goto L6
L8:
	;
	v32 = v25 & int32(7)
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.Ui32(v25) < base.Ui32(int32(8)) {
		v73 = v47
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v47 = l0
	v50 = v25
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = l0
	v38 = v25
	v40 = v26
	goto L13
L13:
	;
	v41 = int32(1)
	v42 = v38 - v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v45 = v40 + v41
	if v45 != v32 {
		v35 = v43
		v38 = v42
		v40 = v45
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v47 = v43
	v50 = v42
	goto L9
L15:
	;
	goto L14
L16:
	;
	v55 = v47
	v58 = v50
	goto L17
L17:
	;
	v61 = int32(8)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v61 < v58 {
		v55 = v70
		v58 = v58 - v61
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v73 = v70
	goto L7
L19:
	;
	goto L18
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_expandRTE(m, v86, v90, v91, v92, v93, v91, v12+int32(32), v12+int32(156))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	if v170 != int32(2) {
		goto L43
	} else {
		goto L44
	}
L23:
	;
	return int32(0)
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if v104 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v107 = v105
	goto L27
L26:
	;
	v107 = int32(0)
	goto L27
L27:
	;
	v108 = F_CreateTemplateTupleDesc(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v114 = int32(0)
	v118 = int32(1)
	goto L29
L29:
	;
	v122 = int32(0)
	if v111 == v122 {
		v131 = v122
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v110 == int32(0) {
		v462 = v108
		goto L3
	} else {
		goto L34
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v125 <= v114 {
		v131 = v122
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v131 = v127 + v114<<(uint(int32(2))%32)
	goto L31
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if base.B2i32(v131 == int32(0))|base.B2i32(v136 <= v114) != 0 {
		v462 = v108
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v139 == int32(0) {
		v462 = v108
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v142 = base.I32_extend16_s(v118)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v139+v114<<(uint(int32(2))%32))))
	v149 = F_exprType(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v151 = F_exprTypmod(m, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	F_TupleDescInitEntry(m, v108, v142, v144, v149, v151, int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	v156 = F_exprCollation(m, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v108+v158<<(uint(int32(4))%32)+v142*int32(100))+16)) = v156
	goto L41
L41:
	;
	v166 = int32(1)
	v114 = v114 + v166
	v118 = v118 + v166
	goto L29
L42:
	;
	goto L5
L43:
	;
	switch v170 - int32(1) {
	case 0:
		goto L47
	default:
		v456 = v15
		goto L42
	case 5:
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+12))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v446+v87<<(uint(int32(2))%32)-int32(4))))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	if v453 == int32(6) {
		v15 = v452
		v19 = v25
		goto L4
	} else {
		goto L111
	}
L46:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+92)))
	if v303 != 0 {
		v456 = v15
		goto L42
	} else {
		goto L77
	}
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+76))
	if v176 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v214 == int32(0) {
		goto L2
	} else {
		goto L61
	}
L49:
	;
	goto L48
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v180 <= int32(0) {
		v214 = int32(0)
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v214 = int32(0)
	goto L49
L53:
	;
	v183 = int32(0)
	if v183 < v180 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v186 = v180
	goto L56
L55:
	;
	v186 = v183
	goto L56
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v191 = int32(0)
	goto L57
L57:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v187+v191<<(uint(int32(2))%32))))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199)+8)))
	if v200 == v87&int32(_a_F_expandRecordVariable_0) {
		v214 = v199
		goto L49
	} else {
		goto L59
	}
L58:
	;
	goto L52
L59:
	;
	v203 = v191 + int32(1)
	if v203 != v186 {
		v191 = v203
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+26)))
	if v218 == int32(1) {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v222 != int32(6) {
		v456 = v221
		goto L42
	} else {
		goto L63
	}
L63:
	;
	v227 = int32(0)
	base.MemoryFill(m, v12+int32(36), v227, int32(120))
	if v25 == v227 {
		v286 = l0
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v286
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v297
	v301 = F_expandRecordVariable(m, v12+int32(32), v221)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L23
	} else {
		goto L76
	}
L65:
	;
	v232 = int32(7)
	v233 = v25 & v232
	if base.Ui32(v232) <= base.Ui32(v25-int32(1)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v241 = l0
	v246 = int32(0)
	goto L69
L67:
	;
	v263 = l0
	goto L68
L68:
	;
	v273 = v263
	v278 = int32(0)
	goto L73
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v259 = v246 + int32(8)
	if v259 != v25&int32(-8) {
		v241 = v257
		v246 = v259
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v233 == int32(0) {
		v286 = v257
		goto L64
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v263 = v257
	goto L68
L73:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v284 = v278 + int32(1)
	if v284 != v233 {
		v273 = v282
		v278 = v284
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v286 = v282
	goto L64
L75:
	;
	goto L74
L76:
	;
	v462 = v301
	goto L3
L77:
	;
	v304 = F_GetCTEForRTE(m, l0, v86, v25)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L23
	} else {
		goto L78
	}
L78:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304)+16))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v309 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v312 = int32(76)
	goto L81
L80:
	;
	v312 = int32(96)
	goto L81
L81:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v306+v312)))
	if v314 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if v352 == int32(0) {
		goto L1
	} else {
		goto L95
	}
L83:
	;
	goto L82
L84:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if v318 <= int32(0) {
		v352 = int32(0)
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v352 = int32(0)
	goto L83
L87:
	;
	v321 = int32(0)
	if v321 < v318 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v324 = v318
	goto L90
L89:
	;
	v324 = v321
	goto L90
L90:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	v329 = int32(0)
	goto L91
L91:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v325+v329<<(uint(int32(2))%32))))
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337)+8)))
	if v338 == v87&int32(_a_F_expandRecordVariable_0) {
		v352 = v337
		goto L83
	} else {
		goto L93
	}
L92:
	;
	goto L86
L93:
	;
	v341 = v329 + int32(1)
	if v341 != v324 {
		v329 = v341
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+26)))
	if v356 == int32(1) {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	if v360 != int32(6) {
		v456 = v359
		goto L42
	} else {
		goto L97
	}
L97:
	;
	v365 = int32(0)
	base.MemoryFill(m, v12+int32(36), v365, int32(120))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v86)+88))
	v369 = v368 + v25
	if v369 == v365 {
		v428 = l0
		goto L98
	} else {
		goto L99
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v428
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v304)+16))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v439
	v443 = F_expandRecordVariable(m, v12+int32(32), v359)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L23
	} else {
		goto L110
	}
L99:
	;
	v372 = int32(7)
	v373 = v369 & v372
	if base.Ui32(v372) <= base.Ui32(v19+v368+v24-int32(1)) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v383 = l0
	v385 = int32(0)
	goto L103
L101:
	;
	v405 = l0
	goto L102
L102:
	;
	v415 = v405
	v417 = int32(0)
	goto L107
L103:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v401 = v385 + int32(8)
	if v401 != v369&int32(-8) {
		v383 = v399
		v385 = v401
		goto L103
	} else {
		goto L105
	}
L104:
	;
	if v373 == int32(0) {
		v428 = v399
		goto L98
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v405 = v399
	goto L102
L107:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v426 = v417 + int32(1)
	if v426 != v373 {
		v415 = v424
		v417 = v426
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v428 = v424
	goto L98
L109:
	;
	goto L108
L110:
	;
	v462 = v443
	goto L3
L111:
	;
	v456 = v452
	goto L42
L112:
	;
	v462 = v460
	goto L3
L113:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v480
	F_errmsg_internal(m, int32(_a_F_expandRecordVariable_1), v12+int32(16))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L23
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_expandRecordVariable_2), int32(1600), int32(_a_F_expandRecordVariable_3))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L23
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v498
	F_errmsg_internal(m, int32(_a_F_expandRecordVariable_4), v12)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L23
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_expandRecordVariable_2), int32(1660), int32(_a_F_expandRecordVariable_3))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L23
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_variable_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 float64
	_ = v188
	var v191 int32
	_ = v191
	var v192 float32
	_ = v192
	var v195 float32
	_ = v195
	var v198 float32
	_ = v198
	var v201 float32
	_ = v201
	var v203 float64
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v224 float64
	_ = v224
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 float64
	_ = v238
	var v242 float32
	_ = v242
	var v244 float64
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v263 float64
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 float32
	_ = v269
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	v6 = int32(0)
	v14 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v6
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)) = uint8(v6)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 == v6 {
		v348 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(96)
	return v348 & int32(1)
L2:
	;
	v28 = F_get_opcode(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_typlenbyval(m, v56, v17+int32(84), v17+int32(83))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L15
	}
L6:
	;
	if v28 == int32(0) {
		v348 = v6
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v35 = F_get_func_leakproof(m, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v35 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v39 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v39 == int32(0) {
		v348 = v6
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = F_get_func_name(m, v28)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v43
	F_errmsg_internal(m, int32(_a_F_get_variable_range_0), v17)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_get_variable_range_1), int32(_a_F_get_variable_range_2), int32(_a_F_get_variable_range_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v348 = v6
	goto L1
L15:
	;
	v64 = v17 + int32(16)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v68 = F_get_attstatsslot(m, v64, v65, int32(2), l1, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L22
	}
L16:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v340
	v348 = v330
	goto L1
L17:
	;
	F_free_attstatsslot(m, v17+int32(16))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L3
	} else {
		goto L59
	}
L18:
	;
	v294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	F_get_stats_slot_range(m, v17+int32(16), v28, v17+int32(52), l2, v294, v295, v17+int32(92), v17+int32(88), v17+int32(87))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L3
	} else {
		goto L58
	}
L19:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v164 <= int32(0) {
		v263 = v14
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v151 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v158 = F_get_attstatsslot(m, v17+int32(16), v154, int32(1), v151, int32(3))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L43
	}
L21:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = F_get_attstatsslot(m, v17+int32(16), v142, int32(1), int32(0), v139)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L40
	}
L22:
	;
	if v68 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v70 != l2 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v108 = v17 + int32(16)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = F_get_attstatsslot(m, v108, v109, int32(2), int32(0), int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L33
	}
L26:
	;
	F_free_attstatsslot(m, v17+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L32
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v73 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v80 = F_datumCopy(m, v77, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83+v84<<(uint(int32(2))%32)-int32(4))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v93 = F_datumCopy(m, v90, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)) = uint8(v95)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v93
	F_free_attstatsslot(m, v64)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v138 = int32(1)
	v139 = int32(1)
	goto L21
L32:
	;
	goto L25
L33:
	;
	if v113 == int32(0) {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	F_get_stats_slot_range(m, v108, v28, v17+int32(52), l2, v119, v120, v17+int32(92), v17+int32(88), v17+int32(87))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	F_free_attstatsslot(m, v108)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v131 = int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)))
	if v133&v131 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v136 = v131
	goto L39
L38:
	;
	v136 = int32(3)
	goto L39
L39:
	;
	v138 = v133
	v139 = v136
	goto L21
L40:
	;
	if v145 == int32(0) {
		v330 = v138
		goto L16
	} else {
		goto L41
	}
L41:
	;
	if v138&int32(1) != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	goto L19
L43:
	;
	if v158 == int32(0) {
		v330 = v151
		goto L16
	} else {
		goto L44
	}
L44:
	;
	goto L19
L45:
	;
	v264 = int32(0)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+22)))
	v269 = *(*float32)(unsafe.Add(mBase, uint32(v266+v267)+8))
	if base.F64_gt(base.F64_add(v263, base.F64_promote_f32(v269)), float64(0.99999)) == v264 {
		v319 = v264
		goto L17
	} else {
		goto L57
	}
L46:
	;
	v168 = v164 & int32(3)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v170 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v164) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v181 = v170
	v186 = v6
	v188 = v14
	goto L50
L48:
	;
	v217 = v170
	v224 = v14
	goto L49
L49:
	;
	v231 = v217
	v237 = v6
	v238 = v224
	goto L54
L50:
	;
	v191 = v169 + v181<<(uint(int32(2))%32)
	v192 = *(*float32)(unsafe.Add(mBase, uint32(v191)))
	v195 = *(*float32)(unsafe.Add(mBase, uint32(v191)+4))
	v198 = *(*float32)(unsafe.Add(mBase, uint32(v191)+8))
	v201 = *(*float32)(unsafe.Add(mBase, uint32(v191)+12))
	v203 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v188, base.F64_promote_f32(v192)), base.F64_promote_f32(v195)), base.F64_promote_f32(v198)), base.F64_promote_f32(v201))
	v204 = int32(4)
	v205 = v181 + v204
	v207 = v186 + v204
	if v207 != v164&int32(2147483644) {
		v181 = v205
		v186 = v207
		v188 = v203
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if v168 == int32(0) {
		v263 = v203
		goto L45
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v217 = v205
	v224 = v203
	goto L49
L54:
	;
	v242 = *(*float32)(unsafe.Add(mBase, uint32(v169+v231<<(uint(int32(2))%32))))
	v244 = base.F64_add(v238, base.F64_promote_f32(v242))
	v245 = int32(1)
	v248 = v237 + v245
	if v248 != v168 {
		v231 = v231 + v245
		v237 = v248
		v238 = v244
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v263 = v244
	goto L45
L56:
	;
	goto L55
L57:
	;
	goto L18
L58:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)))
	v319 = v304
	goto L17
L59:
	;
	v330 = v319
	goto L16
}
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v6)
	v22 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1055), v9+int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v9 + int32(32)
		return v22
	}
}
