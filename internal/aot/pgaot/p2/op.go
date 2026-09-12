package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSetOp(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
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
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v299 int64
	_ = v299
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v312 int64
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v331 int64
	_ = v331
	var v338 int64
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int64
	_ = v431
	var v433 int32
	_ = v433
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v440 int64
	_ = v440
	var v441 int64
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if int64(0) < v24 {
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
	m.G0 = v14 + int32(32)
	return v509
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v24 - int64(1)
	v509 = v16
	goto L6
L8:
	;
	goto L9
L9:
	;
	v30 = int32(0)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v31 != 0 {
		v509 = v30
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v32 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
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
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)))
	if v353 == int32(1) {
		goto L126
	} else {
		goto L127
	}
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v220 = l0 + int32(200)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	goto L76
L17:
	;
	F_ExecReScan(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v46 = m.T0[v45].(func(*base.Module, int32) int32)(m, v41)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v169)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v174 = l0 + int32(200)
	v178 = int32(-1)
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	if v179 == int64(0) {
		v201 = v178
		goto L66
	} else {
		goto L67
	}
L22:
	;
	if v46 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v50&int32(2) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v56 = F_LookupTupleHashEntry(m, v38, v46, v14+int32(31), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v58 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v61 = v59 - v58
	goto L28
L27:
	;
	v61 = int32(0)
	goto L28
L28:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+31)))
	if v62 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v71
	goto L33
L30:
	;
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
	v71 = v65 + int64(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = int64(0)
	v71 = int64(1)
	goto L29
L33:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	F_MemoryContextReset(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	goto L51
L35:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	if v88 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_ExecReScan(m, v41)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v92 = m.T0[v91].(func(*base.Module, int32) int32)(m, v41)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	goto L34
L41:
	;
	if v92 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)))
	if v96&int32(2) != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v99 = int32(0)
	v103 = F_LookupTupleHashEntry(m, v87, v92, v14+int32(31), v99)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	if v105 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v108 = v106 - v105
	goto L47
L46:
	;
	v108 = v99
	goto L47
L47:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+31)))
	if v109 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v112 + int64(1)
	goto L33
L49:
	;
	goto L50
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(1)
	goto L33
L51:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	if v132 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_ExecReScan(m, v39)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v136 = m.T0[v135].(func(*base.Module, int32) int32)(m, v39)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	if v136 == int32(0) {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+4)))
	if v140&int32(2) != 0 {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	v143 = int32(0)
	v145 = F_LookupTupleHashEntry(m, v131, v136, v143, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v145 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v131)+32))
	v149 = v147 - v148
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v149)+8)) = v150 + int64(1)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	F_MemoryContextReset(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L51
L65:
	;
	goto L16
L66:
	;
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)) = uint8(v204)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v201
	goto L65
L67:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v172)+20))
	v184 = int32(0)
	goto L68
L68:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v182+v184*int32(12))+4))
	if v192 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v201 = v178
	goto L66
L70:
	;
	v201 = v184
	goto L66
L71:
	;
	goto L72
L72:
	;
	v196 = v184 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v196)) < base.Ui64(v179) {
		v184 = v196
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v221)+8))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	m.T0[v346].(func(*base.Module, int32))(m, v221)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L123
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v343 = F_ExecStoreMinimalTuple(m, v341, v221, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L122
	}
L76:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v233 != 0 {
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v338 = v331 - int64(1)
	goto L75
L78:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v236 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v236 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+8)))
	v246 = v243
	goto L84
L82:
	;
	goto L81
L83:
	;
	if v278 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L84:
	;
	if v246&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v278 = v261
	goto L83
L86:
	;
	v278 = int32(0)
	goto L83
L87:
	;
	goto L88
L88:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v239)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v257 = v253 & (v254 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v257
	v261 = v252 + v254*int32(12)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v265 = v262 & (v263 ^ v257)
	if v265 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+8)) = uint8(v268)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v272 != int32(1) {
		v246 = base.B2i32(v265 == int32(0))
		goto L84
	} else {
		goto L92
	}
L92:
	;
	goto L85
L93:
	;
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v281)
	v509 = int32(0)
	goto L6
L94:
	;
	goto L95
L95:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v234)+32))
	if v284 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v288 = v285 - v284
	goto L98
L97:
	;
	v288 = int32(0)
	goto L98
L98:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+72))
	switch v290 {
	case 0:
		goto L105
	case 1:
		goto L104
	case 2:
		goto L103
	case 3:
		goto L102
	default:
		goto L101
	}
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v331
	if v331 <= int64(0) {
		goto L76
	} else {
		goto L121
	}
L100:
	;
	v327 = int64(0)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	if v327 < v328 {
		v338 = v327
		goto L75
	} else {
		goto L120
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L117
	}
L102:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	if v308 <= v307 {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	if v299 <= int64(0) {
		goto L110
	} else {
		goto L111
	}
L104:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	if v295 < v296 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v291 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	if int64(0) < v291 {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v331 = int64(0)
	goto L99
L107:
	;
	v298 = v295
	goto L109
L108:
	;
	v298 = v296
	goto L109
L109:
	;
	v331 = v298
	goto L99
L110:
	;
	v331 = int64(0)
	goto L99
L111:
	;
	goto L112
L112:
	;
	v303 = int64(0)
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v288)+8))
	if v304 != v303 {
		v331 = v303
		goto L99
	} else {
		goto L113
	}
L113:
	;
	v338 = v303
	goto L75
L114:
	;
	v312 = v307 - v308
	goto L116
L115:
	;
	v312 = int64(0)
	goto L116
L116:
	;
	v331 = v312
	goto L99
L117:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v289)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v317
	F_errmsg_internal(m, int32(506617), v14)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(521418), int32(149), int32(93934))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v331 = v327
	goto L99
L121:
	;
	goto L77
L122:
	;
	v509 = v343
	goto L6
L123:
	;
	v509 = int32(0)
	goto L6
L124:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	m.T0[v505].(func(*base.Module, int32))(m, v350)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L190
	}
L125:
	;
	goto L146
L126:
	;
	v356 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v356)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v352)+52))
	if v358 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	if v31 != 0 {
		goto L124
	} else {
		goto L145
	}
L129:
	;
	F_ExecReScan(m, v352)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	v362 = m.T0[v361].(func(*base.Module, int32) int32)(m, v352)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v362
	if v362 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v351)+52))
	if v372 != 0 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+4)))
	if v365&int32(2) == int32(0) {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v370 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v370)
	v509 = v30
	goto L6
L138:
	;
	goto L137
L139:
	;
	F_ExecReScan(m, v351)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v376 = m.T0[v375].(func(*base.Module, int32) int32)(m, v351)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	v378 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v378)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v378)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v376
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v383&v378 == int32(0) {
		goto L125
	} else {
		goto L144
	}
L144:
	;
	goto L124
L145:
	;
	goto L125
L146:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v404 == int32(1) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L124
L148:
	;
	F_setop_load_group(m, l0+int32(128), v352, l0)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	if v409 == int64(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L150
L152:
	;
	v412 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v412)
	goto L124
L153:
	;
	goto L154
L154:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v414 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_setop_load_group(m, l0+int32(152), v351, l0)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v419 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	if v419 != int64(0) {
		goto L163
	} else {
		goto L164
	}
L158:
	;
	goto L157
L159:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v490 != int32(1) {
		goto L146
	} else {
		goto L189
	}
L160:
	;
	v485 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v485)
	goto L159
L161:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+72))
	switch v443 {
	case 0:
		goto L176
	case 1:
		goto L171
	case 2:
		goto L175
	case 3:
		goto L174
	default:
		goto L173
	}
L162:
	;
	if v424 != 0 {
		goto L160
	} else {
		goto L168
	}
L163:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v424 = F_setop_compare_slots(m, v422, v423, l0)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v429 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v429)
	v431 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v440 = v431
	v441 = int64(0)
	goto L161
L166:
	;
	if int32(0) <= v424 {
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v433 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v433)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v433)
	v437 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v438 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	v440 = v437
	v441 = v438
	goto L161
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v483
	v509 = v350
	goto L6
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v476
	if v476 <= int64(0) {
		goto L159
	} else {
		goto L188
	}
L171:
	;
	if v440 < v441 {
		goto L185
	} else {
		goto L186
	}
L172:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = int64(0)
	goto L159
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L182
	}
L174:
	;
	if v440 < v441 {
		goto L172
	} else {
		goto L181
	}
L175:
	;
	if v440 <= int64(0) {
		goto L172
	} else {
		goto L179
	}
L176:
	;
	if v440 <= int64(0) {
		goto L172
	} else {
		goto L177
	}
L177:
	;
	if v441 <= int64(0) {
		goto L172
	} else {
		goto L178
	}
L178:
	;
	v483 = int64(0)
	goto L169
L179:
	;
	if v441 != int64(0) {
		goto L172
	} else {
		goto L180
	}
L180:
	;
	v483 = int64(0)
	goto L169
L181:
	;
	v476 = v440 - v441
	goto L170
L182:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v442)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v460
	F_errmsg_internal(m, int32(506617), v14+int32(16))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(521418), int32(149), int32(93934))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	v475 = v440
	goto L187
L186:
	;
	v475 = v441
	goto L187
L187:
	;
	v476 = v475
	goto L170
L188:
	;
	v483 = v476 - int64(1)
	goto L169
L189:
	;
	goto L147
L190:
	;
	v509 = v30
	goto L6
}
func F_get_op_opfamily_strategy(m *base.Module, l0 int32, l1 int32) int32 {
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
	v5 = F_SearchSysCache3(m, int32(3), l0, int32(115), l1)
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
			v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13+v14)+16)))
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
func F_get_op_rettype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+88))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_op_error(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if l4 == int32(1) {
			F_errcode(m, int32(84439172))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = F_op_signature_string(m, l1, l2, l3)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v21
					F_errmsg(m, int32(214007), v10)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errhint(m, int32(613244), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_parser_errposition(m, l0, l5)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_errfinish(m, int32(520641), int32(633), int32(223273))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
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
		} else {
			F_errcode(m, int32(52461700))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = F_op_signature_string(m, l1, l2, l3)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v41
					F_errmsg(m, int32(210274), v10+int32(16))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if l3 != 0 {
							v51 = int32(612949)
						} else {
							v51 = int32(608084)
						}
						if l2 != 0 {
							v53 = v51
						} else {
							v53 = int32(608084)
						}
						F_errhint(m, v53, int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_parser_errposition(m, l0, l5)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(520641), int32(644), int32(223273))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
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
func F_op_hashjoinable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	if l0 != int32(2988) {
		if l0 != int32(1070) {
			v22 = F_SearchSysCache1(m, int32(40), l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v33 = int32(0)
					return v33 & int32(1)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v28)+78)))
					F_ReleaseCatCache(m, v22)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v30
						return v33 & int32(1)
					}
				}
			}
		} else {
			v8 = F_lookup_type_cache(m, l1, int32(16))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
				v33 = base.B2i32(v12 == int32(626))
				return v33 & int32(1)
			}
		}
	} else {
		v16 = F_lookup_type_cache(m, l1, int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
			v33 = base.B2i32(v18 == int32(6192))
			return v33 & int32(1)
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v36 int32
	_ = v36
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg_internal(m, int32(46906), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(525066), int32(1505), int32(172229))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+84))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_op_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	F_initStringInfo(m, v7+int32(32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l1 != 0 {
			v15 = F_format_type_be(m, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v15
				F_appendStringInfo(m, v7+int32(32), int32(774401), v7+int32(16))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v27 = F_NameListToString(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_appendStringInfoString(m, v7+int32(32), v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							v31 = F_format_type_be(m, l2)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v31
								F_appendStringInfo(m, v7+int32(32), int32(217217), v7)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
									m.G0 = v7 + int32(48)
									return v39
								}
							}
						}
					}
				}
			}
		} else {
			v27 = F_NameListToString(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_appendStringInfoString(m, v7+int32(32), v27)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = F_format_type_be(m, l2)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v31
						F_appendStringInfo(m, v7+int32(32), int32(217217), v7)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
							m.G0 = v7 + int32(48)
							return v39
						}
					}
				}
			}
		}
	}
}
