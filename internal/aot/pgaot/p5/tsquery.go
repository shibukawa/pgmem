package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_parse_tsquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v438 int32
	_ = v438
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v21 = l3 & int32(2)
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(1721)
	goto L3
L2:
	;
	v22 = int32(1722)
	goto L3
L3:
	;
	v24 = l3 & int32(1)
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(1720)
	goto L6
L5:
	;
	v25 = v22
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v25
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v30 = base.B2i32(v27 != int32(453))
	goto L9
L8:
	;
	v30 = int32(1)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(12884901888)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	v38 = int32(3)
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = int32(7)
	goto L12
L11:
	;
	v41 = v38
	goto L12
L12:
	;
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = v38
	goto L15
L14:
	;
	v42 = v41
	goto L15
L15:
	;
	v43 = F_init_tsvector_parser(m, l0, v42, l4)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v14)+60)) = int64(64)
	v51 = F_palloc(m, int32(64))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v51
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v55)
	F_makepol_1(m, v14+int32(28), l1, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	F_close_tsvector_parser(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if l4 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L16
	} else {
		goto L119
	}
L22:
	;
	m.G0 = v14 + int32(80)
	return v452
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v72 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v66 != int32(453) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v70 != 0 {
		v452 = int32(0)
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if v30 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if base.Ui32(v98) <= base.Ui32(int32(1073741815)) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	v94 = F_palloc(m, int32(8))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L16
	} else {
		goto L36
	}
L31:
	;
	v79 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	if v79 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v83
	F_errmsg(m, int32(_a_F_parse_tsquery_0), v14)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_1), int32(880), int32(_a_F_parse_tsquery_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = int64(32)
	v452 = v94
	goto L22
L37:
	;
	v129 = v98 + v101*int32(12) + int32(8)
	v130 = F_palloc0(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L47
	}
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v105 = base.I32_div_u_s(int32(1073741815)-v98, int32(12))
	if base.Ui32(v101) <= base.Ui32(v105) {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v108 = int32(0)
	v109 = F_errsave_start(m, l4)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L16
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v109 == int32(0) {
		v452 = v108
		goto L22
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_parse_tsquery_3), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, l4, int32(_a_F_parse_tsquery_1), int32(890), int32(_a_F_parse_tsquery_2))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v452 = v108
	goto L22
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v129 << (uint(int32(2)) % 32)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v135 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v138 = v136
	goto L50
L49:
	;
	v138 = int32(0)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v138
	v141 = v130 + int32(8)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v142 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if int32(0) < v143 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v220 = v138
	goto L53
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	if v227 != 0 {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	v151 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v220 = v215
	goto L53
L57:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158+v151<<(uint(int32(2))%32))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	switch v163 - int32(1) {
	case 0:
		goto L60
	case 1:
		goto L62
	case 2:
		goto L63
	default:
		goto L61
	}
L58:
	;
	goto L56
L59:
	;
	v201 = v151 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v201 < v202 {
		v151 = v201
		goto L57
	} else {
		goto L67
	}
L60:
	;
	v194 = v141 + v151*int32(12)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = v195
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v162)))
	*(*int64)(unsafe.Add(mBase, uint32(v194))) = v197
	goto L59
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L16
	} else {
		goto L64
	}
L62:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v162)))
	*(*int64)(unsafe.Add(mBase, uint32(v141+v151*int32(12)))) = v174
	goto L59
L63:
	;
	v169 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v141+v151*int32(12)))) = uint8(v169)
	goto L59
L64:
	;
	v180 = int32(*(*int8)(unsafe.Add(mBase, uint32(v162))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v180
	F_errmsg_internal(m, int32(_a_F_parse_tsquery_4), v14+int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_1), int32(917), int32(_a_F_parse_tsquery_2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	goto L58
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	base.MemoryCopy(m, v141+v220*int32(12), v231, v227)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	F_pfree(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)) = uint8(v237)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v237
	F_findoprnd_recurse(m, v141, v14+int32(76), v236, v14+int32(27))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v236 != v247 {
		goto L21
	} else {
		goto L73
	}
L73:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	if v249 == int32(0) {
		v452 = v130
		goto L22
	} else {
		goto L74
	}
L74:
	;
	v252 = m.G0
	v254 = v252 - int32(32)
	m.G0 = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v256 == int32(0) {
		v438 = v130
		goto L75
	} else {
		goto L76
	}
L75:
	;
	m.G0 = v254 + int32(32)
	v452 = v438
	goto L22
L76:
	;
	v260 = v130 + int32(8)
	v261 = F_maketree(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	v267 = F_clean_stopword_intree(m, v261, v254+int32(16), v254+int32(12))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	if v267 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	if v30 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v293 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v297 != int32(1) {
		goto L91
	} else {
		goto L92
	}
L82:
	;
	v289 = F_palloc(m, int32(8))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L16
	} else {
		goto L88
	}
L83:
	;
	v275 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L16
	} else {
		goto L84
	}
L84:
	;
	if v275 == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_parse_tsquery_5), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_6), int32(409), int32(_a_F_parse_tsquery_7))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L87
	}
L87:
	;
	goto L82
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = int64(32)
	v438 = v289
	goto L75
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v254)+24)) = int64(16)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v333 = int32(1)
	if base.Ui32((v332-v333)&int32(255)) <= base.Ui32(v333) {
		goto L98
	} else {
		goto L99
	}
L90:
	;
	goto L89
L91:
	;
	v300 = v267
	v301 = v296
	v303 = v293
	goto L94
L92:
	;
	v316 = v296
	v318 = v293
	goto L93
L93:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316)+8))
	v326 = v319&int32(4095) + int32(1)
	v327 = v318
	goto L90
L94:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v305 = F_calcstrlen(m, v304)
	mBase = m.M
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v306 == int32(1) {
		v326 = v305
		v327 = v303
		goto L90
	} else {
		goto L96
	}
L95:
	;
	v316 = v311
	v318 = v309
	goto L93
L96:
	;
	v309 = v305 + v303
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+8))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	if v312 != int32(1) {
		v300 = v310
		v301 = v311
		v303 = v309
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v340 = F_palloc(m, int32(192))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L16
	} else {
		goto L101
	}
L99:
	;
	v349 = v6
	v350 = v6
	goto L100
L100:
	;
	v352 = v350 * int32(12)
	v355 = v326 + v327 + v352 + int32(8)
	v356 = F_palloc(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L16
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+20)) = v340
	F_plainnode(m, v254+int32(20), v267)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v254)+28))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v254)+20))
	v349 = v348
	v350 = v347
	goto L100
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+4)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v355 << (uint(int32(2)) % 32)
	v363 = v356 + int32(8)
	if v352 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	base.MemoryCopy(m, v363, v349, v352)
	goto L106
L105:
	;
	goto L106
L106:
	;
	if int32(0) < v350 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v375 = v363 + v352
	v376 = int32(0)
	v377 = v350
	goto L110
L108:
	;
	goto L109
L109:
	;
	v438 = v356
	goto L75
L110:
	;
	v382 = v363 + v376*int32(12)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v383 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L109
L112:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v388 = v386 & int32(4095)
	if v388 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v419 = v375
	v420 = v377
	goto L114
L114:
	;
	v423 = v376 + int32(1)
	if v423 < v420 {
		v375 = v419
		v376 = v423
		v377 = v420
		goto L110
	} else {
		goto L118
	}
L115:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v390 = int32(12)
	base.MemoryCopy(m, v375, v260+v389*v390+int32(base.Ui32(v386)>>(uint(v390)%32)), v388)
	goto L117
L116:
	;
	goto L117
L117:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v398 = int32(4095)
	v401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v375+v397&v398))) = uint8(v401)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	v405 = v403 & v398
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v407 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v382)+8)) = v405 | (v375-(v363+v406*v407))<<(uint(v407)%32)
	v419 = v405 + v375 + int32(1)
	v420 = v406
	goto L114
L118:
	;
	goto L111
L119:
	;
	F_errmsg_internal(m, int32(_a_F_parse_tsquery_8), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L16
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_parse_tsquery_1), int32(793), int32(_a_F_parse_tsquery_9))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsquery_and(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14393(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_tsquery_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_copy(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_copy(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_CompareTSQ(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v13))
					}
				}
			}
		}
	}
}
func F_tsquery_rewrite_query(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
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
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int64
	_ = v93
	var v99 int32
	_ = v99
	var v110 int64
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_pg_detoast_datum_copy(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_pg_detoast_datum_packed(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v29 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L83
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L80
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L77
	}
L7:
	;
	m.G0 = v19 + int32(32)
	return base.I64_extend_i32_u(v270)
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v27 == v32 {
		v270 = v22
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0]))
	v39 = v22 + int32(8)
	v43 = F_QT2QTN(m, v39, v39+v29*int32(12))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_pfree(m, v27)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v270 = v22
	goto L7
L13:
	;
	F_QTNTernary(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_QTNSort(m, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v49 = F_text_to_cstring(m, v27)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v54 = int32(0)
	v56 = F_SPI_prepare(m, v49, v54, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v56 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v60 = F_SPI_cursor_open(m, v56)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v60 == int32(0) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_SPI_cursor_fetch(m, v60, int32(100))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	if v68 == int32(0) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 != int32(2) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v76 = F_SPI_gettypeid(m, v71, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v76 != int32(3615) {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v84 = F_SPI_gettypeid(m, v82, int32(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v84 != int32(3615) {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v88 = int32(0)
	v93 = *(*int64)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[2]))
	if base.B2i32(v43 == v88)|base.B2i32(v93 == int64(0)) != 0 {
		v222 = base.B2i32(v43 != v88)
		v223 = v43
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	F_SPI_freetuptable(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L59
	}
L30:
	;
	v99 = v43
	v110 = int64(0)
	goto L31
L31:
	;
	v115 = base.I32_wrap_i64(v110) << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115+v118)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v124 = v19 + int32(30)
	v125 = F_SPI_getbinval(m, v120, v121, int32(1), v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v222 = v215
	v223 = v202
	goto L29
L33:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v127 != 0 {
		v192 = v99
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	F_SPI_freetuptable(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L55
	}
L35:
	;
	v197 = v110 + int64(1)
	v199 = *(*int64)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[2]))
	if base.Ui64(v197) < base.Ui64(v199) {
		v99 = v192
		v110 = v197
		goto L31
	} else {
		goto L54
	}
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[1]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v115+v130)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v135 = F_SPI_getbinval(m, v132, v133, int32(2), v124)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v137 != 0 {
		v192 = v99
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v138 = base.I32_wrap_i64(v125)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v139 == int32(0) {
		v192 = v99
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v143 = v138 + int32(8)
	v147 = F_QT2QTN(m, v143, v143+v139*int32(12))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_QTNTernary(m, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_QTNSort(m, v147)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v154 = base.I32_wrap_i64(v135)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	if v155 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v157 = v154 + int32(8)
	v161 = F_QT2QTN(m, v157, v157+v155*int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	v164 = int32(0)
	goto L45
L45:
	;
	v165 = int32(_a_F_tsquery_rewrite_query_0)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0])) = v37
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)) = uint8(v169)
	v173 = F_dofindsubquery(m, v99, v147, v164, v19+int32(31))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v164 = v161
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[0])) = v166
	F_QTNFree(m, v147)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_QTNFree(m, v164)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v181 = int32(0)
	if v173 == v181 {
		v202 = v181
		goto L34
	} else {
		goto L50
	}
L50:
	;
	F_QTNClearFlags(m, v173, int32(2))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_QTNTernary(m, v173)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_QTNSort(m, v173)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v192 = v173
	goto L35
L54:
	;
	v202 = v192
	goto L34
L55:
	;
	F_SPI_cursor_fetch(m, v60, int32(100))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v215 = base.B2i32(v202 != int32(0))
	v217 = *(*int64)(unsafe.Add(mBase, _c_F_tsquery_rewrite_query[2]))
	if v217 == int64(0) {
		v222 = v215
		v223 = v202
		goto L29
	} else {
		goto L57
	}
L57:
	;
	if v202 != 0 {
		v99 = v202
		v110 = int64(0)
		goto L31
	} else {
		goto L58
	}
L58:
	;
	goto L32
L59:
	;
	F_SPI_cursor_close(m, v60)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_SPI_freeplan(m, v56)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v245 = F_SPI_finish(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v222 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	F_pfree(m, v49)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L72
	}
L64:
	;
	F_QTNBinary(m, v223)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = int64(32)
	v259 = v22
	goto L63
L67:
	;
	v249 = F_QTN2QT(m, v223)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_QTNFree(m, v223)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v22 == v253 {
		v259 = v249
		goto L63
	} else {
		goto L70
	}
L70:
	;
	F_pfree(m, v22)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v259 = v249
	goto L63
L72:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v262 != v27 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_pfree(m, v27)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v270 = v259
	goto L7
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v49
	F_errmsg_internal(m, int32(_a_F_tsquery_rewrite_query_1), v19)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_tsquery_rewrite_query_2), int32(308), int32(_a_F_tsquery_rewrite_query_3))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v49
	F_errmsg_internal(m, int32(_a_F_tsquery_rewrite_query_4), v19+int32(16))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_tsquery_rewrite_query_2), int32(311), int32(_a_F_tsquery_rewrite_query_3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_tsquery_rewrite_query_5), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_tsquery_rewrite_query_2), int32(321), int32(_a_F_tsquery_rewrite_query_3))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
