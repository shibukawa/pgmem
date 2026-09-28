package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__int_matchsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 float32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v74 int32
	_ = v74
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
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
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v306 int32
	_ = v306
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int64
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 float32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 float32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 float32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int64
	_ = v457
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 float32
	_ = v470
	var v471 float64
	_ = v471
	var v476 float64
	_ = v476
	var v477 int32
	_ = v477
	var v478 float64
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 float64
	_ = v485
	var v493 float64
	_ = v493
	var v509 int64
	_ = v509
	v2 = int32(0)
	v17 = float32(0)
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	v23 = int64(4572414629676717179)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v33 = F_get_restriction_variable(m, v24, v25, v26, v21+int32(48), v21+int32(44), v21+int32(43))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(80)
	return v509
L2:
	;
	return int64(0)
L3:
	;
	if v33 == int32(0) {
		v509 = v23
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	if v39 != int32(1007) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v42 == int32(0) {
		v509 = v23
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 != int32(7) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v45].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v509 = v23
	goto L1
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v52 == int32(0) {
		v509 = v23
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+32)))
	if v58 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v55].(func(*base.Module, int32))(m, v52)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v509 = v23
	goto L1
L15:
	;
	v61 = int64(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v62 == int32(0) {
		v509 = v61
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v73 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0]))
	if v73 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v65].(func(*base.Module, int32))(m, v62)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v509 = v61
	goto L1
L20:
	;
	if v68 != v400 {
		goto L84
	} else {
		goto L85
	}
L21:
	;
	v400 = v375
	goto L20
L22:
	;
	v168 = F_getExtensionOfObject(m, int32(1255), v70)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L42
	}
L23:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+12)))
	if v145 != int32(1) {
		v149 = v74
		v155 = int32(0)
		goto L22
	} else {
		goto L41
	}
L24:
	;
	v74 = v73
	goto L27
L25:
	;
	goto L26
L26:
	;
	v149 = int32(0)
	v155 = int32(1)
	goto L22
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v92 == v70 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v94 = int32(_a_F__int_matchsel_0)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__int_matchsel[1])))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if base.B2i32(v98 == int32(0))|base.B2i32(v98 != v101) != 0 {
		v119 = v98
		v120 = v101
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v124 != 0 {
		v74 = v124
		goto L27
	} else {
		goto L40
	}
L32:
	;
	if v119-v120 == int32(0) {
		goto L23
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v104 = v94
	v105 = v95
	goto L35
L35:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v109
		v120 = v108
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v119 = v109
	v120 = v108
	goto L33
L37:
	;
	v112 = int32(1)
	if v109 == v108 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L31
L40:
	;
	goto L28
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	v375 = v148
	goto L21
L42:
	;
	if v168 == int32(0) {
		v375 = v2
		goto L21
	} else {
		goto L43
	}
L43:
	;
	v172 = m.G0
	v174 = v172 - int32(176)
	m.G0 = v174
	v178 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	F_ScanKeyInit(m, v174, int32(4), int32(3), int32(184), int64(3079))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_ScanKeyInit(m, v174+int32(56), int32(5), int32(3), int32(184), base.I64_extend_i32_u(v168))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	F_ScanKeyInit(m, v174+int32(112), int32(6), int32(3), int32(65), int64(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v206 = F_systable_beginscan(m, v178, int32(2674), int32(1), int32(0), int32(3), v174)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L49
	}
L48:
	;
	F_systable_endscan(m, v206)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L2
	} else {
		goto L72
	}
L49:
	;
	v208 = F_systable_getnext(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	if v208 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v212 = v208
	goto L54
L52:
	;
	goto L53
L53:
	;
	v306 = int32(0)
	goto L48
L54:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+22)))
	v230 = v228 + v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v231 != int32(1247) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v283 = F_systable_getnext(m, v206)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L2
	} else {
		goto L70
	}
L57:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+24)))
	if v234 != int32(101) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v240 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v238))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	if v240 == int32(0) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+16))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+22)))
	v248 = v244 + v245 + int32(4)
	v249 = int32(_a_F__int_matchsel_0)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__int_matchsel[1])))
	if base.B2i32(v252 == int32(0))|base.B2i32(v252 != v255) != 0 {
		v273 = v252
		v274 = v255
		goto L62
	} else {
		goto L63
	}
L61:
	;
	F_ReleaseCatCache(m, v240)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L68
	}
L62:
	;
	goto L61
L63:
	;
	v258 = v248
	v259 = v249
	goto L64
L64:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)))
	if v263 == int32(0) {
		v273 = v263
		v274 = v262
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v273 = v263
	v274 = v262
	goto L62
L66:
	;
	v266 = int32(1)
	if v263 == v262 {
		v258 = v258 + v266
		v259 = v259 + v266
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v273-v274 == int32(0) {
		v306 = v238
		goto L48
	} else {
		goto L69
	}
L69:
	;
	goto L56
L70:
	;
	if v283 != 0 {
		v212 = v283
		goto L54
	} else {
		goto L71
	}
L71:
	;
	goto L55
L72:
	;
	F_relation_close(m, v178, int32(1))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	m.G0 = v174 + int32(176)
	if v306 == int32(0) {
		v375 = v2
		goto L21
	} else {
		goto L74
	}
L74:
	;
	if v155 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0]))
	if v333 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v351 = v149
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351)+8)) = int32(_a_F__int_matchsel_0)
	*(*int32)(unsafe.Add(mBase, uint32(v351)+4)) = v70
	v358 = F_GetSysCacheHashValue(m, int32(28), base.I64_extend_i32_u(v168), int64(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L83
	}
L78:
	;
	F_CacheRegisterSyscacheCallback(m, int32(28), int32(594), int64(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[2]))
	v344 = F_MemoryContextAllocZero(m, v342, int32(24))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v346 = int32(_a_F__int_matchsel_1)
	v347 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v347
	*(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0])) = v344
	v351 = v344
	goto L77
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351)+20)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v351)+16)) = v358
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v351)+12)) = uint8(v362)
	v400 = v306
	goto L20
L84:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v402 == int32(0) {
		v509 = v23
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+24))
	v410 = F_pg_detoast_datum(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L89
	}
L87:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v405].(func(*base.Module, int32))(m, v402)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	v509 = v23
	goto L1
L89:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v413 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v416 = int64(0)
	if v412 == int32(0) {
		v509 = v416
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v412 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v419].(func(*base.Module, int32))(m, v412)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v509 = v416
	goto L1
L95:
	;
	v476 = F_int_query_opr_selec(m, v410+v467<<(uint(int32(3))%32), v466, v468, v469, v470)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L2
	} else {
		goto L103
	}
L96:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v412)+16))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+22)))
	v426 = *(*float32)(unsafe.Add(mBase, uint32(v423+v424)+8))
	v427 = int32(0)
	v431 = F_get_attstatsslot(m, v21, v412, int32(4), v427, int32(3))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v454 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v454
	v457 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v457
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v457
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v457
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v457
	v466 = v454
	v467 = v413
	v468 = v2
	v469 = v2
	v470 = v17
	v471 = float64(0)
	goto L95
L99:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v466 = v451
	v467 = v452
	v468 = v448
	v469 = v449
	v470 = v450
	v471 = base.F64_promote_f32(v426)
	goto L95
L100:
	;
	if v431 == int32(0) {
		v448 = v2
		v449 = v2
		v450 = v17
		v451 = v427
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v436 != v437+int32(3) {
		v448 = v2
		v449 = v2
		v450 = v17
		v451 = int32(0)
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v445 = *(*float32)(unsafe.Add(mBase, uint32(v441+v437<<(uint(int32(2))%32))))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v448 = v441
	v449 = v437
	v450 = v445
	v451 = v446
	goto L99
L103:
	;
	v478 = base.F64_mul(base.F64_sub(float64(1), v471), v476)
	F_free_attstatsslot(m, v21)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v481 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v482].(func(*base.Module, int32))(m, v481)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v485 = float64(0)
	if base.F64_lt(v478, v485) != 0 {
		v493 = v485
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L107
L109:
	;
	v509 = base.I64_reinterpret_f64(v493)
	goto L1
L110:
	;
	if base.F64_gt(v478, float64(1)) == int32(0) {
		v493 = v478
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v493 = float64(1)
	goto L109
}
func F__int_overlap(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_copy(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v20 != 0 {
				v21 = F_array_contains_nulls(m, v13)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					if v21 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F__int_overlap_0), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F__int_overlap_1), int32(109), int32(_a_F__int_overlap_2))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
						if v23 != 0 {
							v24 = F_array_contains_nulls(m, v18)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int64(0)
							} else {
								if v24 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F__int_overlap_0), int32(0))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F__int_overlap_1), int32(110), int32(_a_F__int_overlap_2))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v28 = v13 + int32(16)
									v29 = F_ArrayGetNItemsSafe(m, v26, v28)
									mBase = m.M
									v30 = m.ExcPending
									if v30 != 0 {
										return int64(0)
									} else {
										if v29 == int32(0) {
											v86 = v7
											m.G0 = v10 + int32(16)
											return v86
										} else {
											v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
											v35 = v18 + int32(16)
											v36 = F_ArrayGetNItemsSafe(m, v33, v35)
											mBase = m.M
											v37 = m.ExcPending
											if v37 != 0 {
												return int64(0)
											} else {
												if v36 == int32(0) {
													v86 = v7
													m.G0 = v10 + int32(16)
													return v86
												} else {
													v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v41 = F_ArrayGetNItemsSafe(m, v40, v28)
													mBase = m.M
													v42 = m.ExcPending
													if v42 != 0 {
														return int64(0)
													} else {
														v43 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v43)
														v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v45 != 0 {
															v53 = v45
														} else {
															v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v53 = (v46<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v53+v13, v41, v10+int32(15))
														mBase = m.M
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
														v59 = F_ArrayGetNItemsSafe(m, v58, v35)
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return int64(0)
														} else {
															v61 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v61)
															v63 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
															if v63 != 0 {
																v71 = v63
															} else {
																v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
																v71 = (v64<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v71+v18, v59, v10+int32(14))
															mBase = m.M
															v76 = F_inner_int_overlap(m, v13, v18)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return int64(0)
															} else {
																F_pfree(m, v13)
																mBase = m.M
																v79 = m.ExcPending
																if v79 != 0 {
																	return int64(0)
																} else {
																	F_pfree(m, v18)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int64(0)
																	} else {
																		v86 = base.I64_extend_i32_u(v76)
																		m.G0 = v10 + int32(16)
																		return v86
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
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v28 = v13 + int32(16)
							v29 = F_ArrayGetNItemsSafe(m, v26, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int64(0)
							} else {
								if v29 == int32(0) {
									v86 = v7
									m.G0 = v10 + int32(16)
									return v86
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
									v35 = v18 + int32(16)
									v36 = F_ArrayGetNItemsSafe(m, v33, v35)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										if v36 == int32(0) {
											v86 = v7
											m.G0 = v10 + int32(16)
											return v86
										} else {
											v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v41 = F_ArrayGetNItemsSafe(m, v40, v28)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return int64(0)
											} else {
												v43 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v43)
												v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												if v45 != 0 {
													v53 = v45
												} else {
													v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v53 = (v46<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v53+v13, v41, v10+int32(15))
												mBase = m.M
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
												v59 = F_ArrayGetNItemsSafe(m, v58, v35)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int64(0)
												} else {
													v61 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v61)
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
													if v63 != 0 {
														v71 = v63
													} else {
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
														v71 = (v64<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v71+v18, v59, v10+int32(14))
													mBase = m.M
													v76 = F_inner_int_overlap(m, v13, v18)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int64(0)
													} else {
														F_pfree(m, v13)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int64(0)
														} else {
															F_pfree(m, v18)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int64(0)
															} else {
																v86 = base.I64_extend_i32_u(v76)
																m.G0 = v10 + int32(16)
																return v86
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
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				if v23 != 0 {
					v24 = F_array_contains_nulls(m, v18)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						if v24 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F__int_overlap_0), int32(0))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F__int_overlap_1), int32(110), int32(_a_F__int_overlap_2))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v28 = v13 + int32(16)
							v29 = F_ArrayGetNItemsSafe(m, v26, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int64(0)
							} else {
								if v29 == int32(0) {
									v86 = v7
									m.G0 = v10 + int32(16)
									return v86
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
									v35 = v18 + int32(16)
									v36 = F_ArrayGetNItemsSafe(m, v33, v35)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										if v36 == int32(0) {
											v86 = v7
											m.G0 = v10 + int32(16)
											return v86
										} else {
											v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v41 = F_ArrayGetNItemsSafe(m, v40, v28)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return int64(0)
											} else {
												v43 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v43)
												v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												if v45 != 0 {
													v53 = v45
												} else {
													v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v53 = (v46<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v53+v13, v41, v10+int32(15))
												mBase = m.M
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
												v59 = F_ArrayGetNItemsSafe(m, v58, v35)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int64(0)
												} else {
													v61 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v61)
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
													if v63 != 0 {
														v71 = v63
													} else {
														v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
														v71 = (v64<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v71+v18, v59, v10+int32(14))
													mBase = m.M
													v76 = F_inner_int_overlap(m, v13, v18)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return int64(0)
													} else {
														F_pfree(m, v13)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int64(0)
														} else {
															F_pfree(m, v18)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return int64(0)
															} else {
																v86 = base.I64_extend_i32_u(v76)
																m.G0 = v10 + int32(16)
																return v86
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
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v28 = v13 + int32(16)
					v29 = F_ArrayGetNItemsSafe(m, v26, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						if v29 == int32(0) {
							v86 = v7
							m.G0 = v10 + int32(16)
							return v86
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
							v35 = v18 + int32(16)
							v36 = F_ArrayGetNItemsSafe(m, v33, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								if v36 == int32(0) {
									v86 = v7
									m.G0 = v10 + int32(16)
									return v86
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v41 = F_ArrayGetNItemsSafe(m, v40, v28)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int64(0)
									} else {
										v43 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v43)
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										if v45 != 0 {
											v53 = v45
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v53 = (v46<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v53+v13, v41, v10+int32(15))
										mBase = m.M
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
										v59 = F_ArrayGetNItemsSafe(m, v58, v35)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int64(0)
										} else {
											v61 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v61)
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
											if v63 != 0 {
												v71 = v63
											} else {
												v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
												v71 = (v64<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v71+v18, v59, v10+int32(14))
											mBase = m.M
											v76 = F_inner_int_overlap(m, v13, v18)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int64(0)
											} else {
												F_pfree(m, v13)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int64(0)
												} else {
													F_pfree(m, v18)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int64(0)
													} else {
														v86 = base.I64_extend_i32_u(v76)
														m.G0 = v10 + int32(16)
														return v86
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
func F__int_unique(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = l0 + int32(16)
	v12 = F_ArrayGetNItemsSafe(m, v9, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v16 == int32(0) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v26 = (v19<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		} else {
			v26 = v16
		}
		if base.Ui32(int32(2)) <= base.Ui32(v12) {
			v29 = l0 + v26
			v32 = int32(0)
			v33 = int32(1)
			for {
				v39 = int32(2)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v29+v33<<(uint(v39)%32))))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v29+v32<<(uint(v39)%32))))
				if v42 == v46 {
					v55 = v32
				} else {
					v49 = v32 + int32(1)
					if v49 == v33 {
						v55 = v33
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v29+v49<<(uint(int32(2))%32)))) = v42
						v55 = v49
					}
				}
				v58 = v33 + int32(1)
				if v58 != v12 {
					v32 = v55
					v33 = v58
					continue
				} else {
					break
				}
				break
			}
			v65 = v55 + int32(1)
		} else {
			v65 = v12
		}
		if v65 <= int32(0) {
			v73 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				return v73
			}
		} else {
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v77 = F_ArrayGetNItemsSafe(m, v76, v11)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				if v77 == v65 {
					v170 = l0
					return v170
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v80 != 0 {
						v88 = v80
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v88 = (v81<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					v91 = v88 + v65<<(uint(int32(2))%32)
					v92 = F_repalloc(m, l0, v91)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v92))) = v91 << (uint(int32(2)) % 32)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
						if v97 <= int32(0) {
							v170 = v92
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v92)+16)) = v65
							if v97 == int32(1) {
								v170 = v92
							} else {
								v104 = v92 + int32(16)
								v105 = int32(1)
								v106 = v97 - v105
								v107 = int32(7)
								v108 = v106 & v107
								if base.Ui32(v107) <= base.Ui32(v97-int32(2)) {
									v119 = v105
									v121 = int32(0)
									for {
										v127 = v104 + v119<<(uint(int32(2))%32)
										v128 = int64(4294967297)
										*(*int64)(unsafe.Add(mBase, uint32(v127)+24)) = v128
										*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = v128
										*(*int64)(unsafe.Add(mBase, uint32(v127)+8)) = v128
										*(*int64)(unsafe.Add(mBase, uint32(v127))) = v128
										v136 = int32(8)
										v137 = v119 + v136
										v139 = v121 + v136
										if v139 != v106&int32(-8) {
											v119 = v137
											v121 = v139
											continue
										} else {
											break
										}
										break
									}
									if v108 == int32(0) {
										v170 = v92
									} else {
										v145 = v137
										v153 = int32(0)
										v154 = v145
										for {
											v163 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v104+v154<<(uint(int32(2))%32)))) = v163
											v168 = v153 + v163
											if v168 != v108 {
												v153 = v168
												v154 = v154 + v163
												continue
											} else {
												break
											}
											break
										}
										v170 = v92
									}
								} else {
									v145 = v105
									v153 = int32(0)
									v154 = v145
									for {
										v163 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v104+v154<<(uint(int32(2))%32)))) = v163
										v168 = v153 + v163
										if v168 != v108 {
											v153 = v168
											v154 = v154 + v163
											continue
										} else {
											break
										}
										break
									}
									v170 = v92
								}
							}
						}
						return v170
					}
				}
			}
		}
	}
}
