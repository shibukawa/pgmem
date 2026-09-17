package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__int_matchsel(m *base.Module, l0 int32) int32 {
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
	var v23 float64
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 float64
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v298 int32
	_ = v298
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
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
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v366 int32
	_ = v366
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 float64
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 float32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 float32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 float32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int64
	_ = v448
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 float32
	_ = v461
	var v462 float64
	_ = v462
	var v467 float64
	_ = v467
	var v468 int32
	_ = v468
	var v469 float64
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 float64
	_ = v476
	var v498 float64
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	v2 = int32(0)
	v17 = float32(0)
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	v23 = float64(0.005)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
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
	v502 = F_Float8GetDatum(m, v498)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L2
	} else {
		goto L107
	}
L2:
	;
	return int32(0)
L3:
	;
	if v33 == int32(0) {
		v498 = v23
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	if v39 != int32(1007) {
		v498 = v23
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(7) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v46 == int32(0) {
		v498 = v23
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+24)))
	if v52 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v49].(func(*base.Module, int32))(m, v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v498 = v23
	goto L1
L11:
	;
	v55 = float64(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v56 == int32(0) {
		v498 = v55
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0]))
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v59].(func(*base.Module, int32))(m, v56)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v498 = v55
	goto L1
L16:
	;
	if v62 != v391 {
		goto L80
	} else {
		goto L81
	}
L17:
	;
	v391 = v366
	goto L16
L18:
	;
	v162 = F_getExtensionOfObject(m, int32(1255), v64)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L38
	}
L19:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
	if v139 != int32(1) {
		v143 = v68
		v149 = int32(0)
		goto L18
	} else {
		goto L37
	}
L20:
	;
	v68 = v67
	goto L23
L21:
	;
	goto L22
L22:
	;
	v143 = int32(0)
	v149 = int32(1)
	goto L18
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v86 == v64 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	v88 = int32(_a_F__int_matchsel_0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__int_matchsel[1])))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if base.B2i32(v92 == int32(0))|base.B2i32(v92 != v95) != 0 {
		v113 = v92
		v114 = v95
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v118 != 0 {
		v68 = v118
		goto L23
	} else {
		goto L36
	}
L28:
	;
	if v113-v114 == int32(0) {
		goto L19
	} else {
		goto L35
	}
L29:
	;
	goto L28
L30:
	;
	v98 = v88
	v99 = v89
	goto L31
L31:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v103 == int32(0) {
		v113 = v103
		v114 = v102
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v113 = v103
	v114 = v102
	goto L29
L33:
	;
	v106 = int32(1)
	if v103 == v102 {
		v98 = v98 + v106
		v99 = v99 + v106
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L27
L36:
	;
	goto L24
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v366 = v142
	goto L17
L38:
	;
	if v162 == int32(0) {
		v366 = v2
		goto L17
	} else {
		goto L39
	}
L39:
	;
	v166 = m.G0
	v168 = v166 - int32(144)
	m.G0 = v168
	v172 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_ScanKeyInit(m, v168, int32(4), int32(3), int32(184), int32(3079))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_ScanKeyInit(m, v168+int32(48), int32(5), int32(3), int32(184), v162)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_ScanKeyInit(m, v168+int32(96), int32(6), int32(3), int32(65), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v199 = F_systable_beginscan(m, v172, int32(2674), int32(1), int32(0), int32(3), v168)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L45
	}
L44:
	;
	F_systable_endscan(m, v199)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L2
	} else {
		goto L68
	}
L45:
	;
	v201 = F_systable_getnext(m, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	if v201 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v205 = v201
	goto L50
L48:
	;
	goto L49
L49:
	;
	v298 = int32(0)
	goto L44
L50:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v205)+16))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+22)))
	v223 = v221 + v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v224 != int32(1247) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v275 = F_systable_getnext(m, v199)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L66
	}
L53:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+24)))
	if v227 != int32(101) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v232 = F_SearchSysCache1(m, int32(82), v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	if v232 == int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+22)))
	v240 = v236 + v237 + int32(4)
	v241 = int32(_a_F__int_matchsel_0)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__int_matchsel[1])))
	if base.B2i32(v244 == int32(0))|base.B2i32(v244 != v247) != 0 {
		v265 = v244
		v266 = v247
		goto L58
	} else {
		goto L59
	}
L57:
	;
	F_ReleaseCatCache(m, v232)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L64
	}
L58:
	;
	goto L57
L59:
	;
	v250 = v240
	v251 = v241
	goto L60
L60:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	if v255 == int32(0) {
		v265 = v255
		v266 = v254
		goto L58
	} else {
		goto L62
	}
L61:
	;
	v265 = v255
	v266 = v254
	goto L58
L62:
	;
	v258 = int32(1)
	if v255 == v254 {
		v250 = v250 + v258
		v251 = v251 + v258
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	if v265-v266 == int32(0) {
		v298 = v231
		goto L44
	} else {
		goto L65
	}
L65:
	;
	goto L52
L66:
	;
	if v275 != 0 {
		v205 = v275
		goto L50
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	F_relation_close(m, v172, int32(1))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	m.G0 = v168 + int32(144)
	if v298 == int32(0) {
		v366 = v2
		goto L17
	} else {
		goto L70
	}
L70:
	;
	if v149 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0]))
	if v325 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v343 = v143
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+8)) = int32(_a_F__int_matchsel_0)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+4)) = v64
	v349 = F_GetSysCacheHashValue(m, int32(28), v162, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L79
	}
L74:
	;
	F_CacheRegisterSyscacheCallback(m, int32(28), int32(553), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[2]))
	v336 = F_MemoryContextAllocZero(m, v334, int32(24))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v338 = int32(_a_F__int_matchsel_1)
	v339 = *(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v336))) = v339
	*(*int32)(unsafe.Add(mBase, _c_F__int_matchsel[0])) = v336
	v343 = v336
	goto L73
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+20)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v343)+16)) = v349
	v353 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v343)+12)) = uint8(v353)
	v391 = v298
	goto L16
L80:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v393 == int32(0) {
		v498 = v23
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+20))
	v401 = F_pg_detoast_datum(m, v400)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L2
	} else {
		goto L85
	}
L83:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v396].(func(*base.Module, int32))(m, v393)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v498 = v23
	goto L1
L85:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v404 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v407 = float64(0)
	if v403 == int32(0) {
		v498 = v407
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v403 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v410].(func(*base.Module, int32))(m, v403)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v498 = v407
	goto L1
L91:
	;
	v467 = F_int_query_opr_selec(m, v401+v458<<(uint(int32(3))%32), v457, v459, v460, v461)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L2
	} else {
		goto L99
	}
L92:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+22)))
	v417 = *(*float32)(unsafe.Add(mBase, uint32(v414+v415)+8))
	v418 = int32(0)
	v422 = F_get_attstatsslot(m, v21, v403, int32(4), v418, int32(3))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v445
	v448 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v448
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v448
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v448
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v448
	v457 = v445
	v458 = v404
	v459 = v2
	v460 = v2
	v461 = v17
	v462 = float64(0)
	goto L91
L95:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v457 = v442
	v458 = v443
	v459 = v439
	v460 = v440
	v461 = v441
	v462 = base.F64_promote_f32(v417)
	goto L91
L96:
	;
	if v422 == int32(0) {
		v439 = v2
		v440 = v2
		v441 = v17
		v442 = v418
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v427 != v428+int32(3) {
		v439 = v2
		v440 = v2
		v441 = v17
		v442 = int32(0)
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v436 = *(*float32)(unsafe.Add(mBase, uint32(v432+v428<<(uint(int32(2))%32))))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v439 = v432
	v440 = v428
	v441 = v436
	v442 = v437
	goto L95
L99:
	;
	v469 = base.F64_mul(base.F64_sub(float64(1), v462), v467)
	F_free_attstatsslot(m, v21)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v472 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	m.T0[v473].(func(*base.Module, int32))(m, v472)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L2
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v476 = float64(0)
	if base.F64_lt(v469, v476) != 0 {
		v498 = v476
		goto L1
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	if base.F64_gt(v469, float64(1)) == int32(0) {
		v498 = v469
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v498 = float64(1)
	goto L1
L107:
	;
	m.G0 = v21 + int32(80)
	return v502
}
func F__int_overlap(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_copy(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 != 0 {
				v20 = F_array_contains_nulls(m, v12)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F__int_overlap_0), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F__int_overlap_1), int32(108), int32(_a_F__int_overlap_2))
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
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v22 != 0 {
							v23 = F_array_contains_nulls(m, v17)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								if v23 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67108994))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F__int_overlap_0), int32(0))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F__int_overlap_1), int32(109), int32(_a_F__int_overlap_2))
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
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
									v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									v27 = v12 + int32(16)
									v28 = F_ArrayGetNItemsSafe(m, v25, v27)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										if v28 == int32(0) {
											v83 = v2
											m.G0 = v9 + int32(16)
											return v83
										} else {
											v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
											v34 = v17 + int32(16)
											v35 = F_ArrayGetNItemsSafe(m, v32, v34)
											mBase = m.M
											v36 = m.ExcPending
											if v36 != 0 {
												return int32(0)
											} else {
												if v35 == int32(0) {
													v83 = v2
													m.G0 = v9 + int32(16)
													return v83
												} else {
													v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													v40 = F_ArrayGetNItemsSafe(m, v39, v27)
													mBase = m.M
													v41 = m.ExcPending
													if v41 != 0 {
														return int32(0)
													} else {
														v42 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v42)
														v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
														if v44 != 0 {
															v52 = v44
														} else {
															v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
															v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v52+v12, v40, v9+int32(15))
														mBase = m.M
														v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
														v58 = F_ArrayGetNItemsSafe(m, v57, v34)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return int32(0)
														} else {
															v60 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v60)
															v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
															if v62 != 0 {
																v70 = v62
															} else {
																v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
																v70 = (v63<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v70+v17, v58, v9+int32(14))
															mBase = m.M
															v75 = F_inner_int_overlap(m, v12, v17)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v12)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v17)
																	mBase = m.M
																	v80 = m.ExcPending
																	if v80 != 0 {
																		return int32(0)
																	} else {
																		v83 = v75
																		m.G0 = v9 + int32(16)
																		return v83
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
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v27 = v12 + int32(16)
							v28 = F_ArrayGetNItemsSafe(m, v25, v27)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if v28 == int32(0) {
									v83 = v2
									m.G0 = v9 + int32(16)
									return v83
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									v34 = v17 + int32(16)
									v35 = F_ArrayGetNItemsSafe(m, v32, v34)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										if v35 == int32(0) {
											v83 = v2
											m.G0 = v9 + int32(16)
											return v83
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											v40 = F_ArrayGetNItemsSafe(m, v39, v27)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												v42 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v42)
												v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												if v44 != 0 {
													v52 = v44
												} else {
													v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v52+v12, v40, v9+int32(15))
												mBase = m.M
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
												v58 = F_ArrayGetNItemsSafe(m, v57, v34)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return int32(0)
												} else {
													v60 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v60)
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
													if v62 != 0 {
														v70 = v62
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
														v70 = (v63<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v70+v17, v58, v9+int32(14))
													mBase = m.M
													v75 = F_inner_int_overlap(m, v12, v17)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v12)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v17)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int32(0)
															} else {
																v83 = v75
																m.G0 = v9 + int32(16)
																return v83
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
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				if v22 != 0 {
					v23 = F_array_contains_nulls(m, v17)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F__int_overlap_0), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F__int_overlap_1), int32(109), int32(_a_F__int_overlap_2))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
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
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v27 = v12 + int32(16)
							v28 = F_ArrayGetNItemsSafe(m, v25, v27)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if v28 == int32(0) {
									v83 = v2
									m.G0 = v9 + int32(16)
									return v83
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									v34 = v17 + int32(16)
									v35 = F_ArrayGetNItemsSafe(m, v32, v34)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										if v35 == int32(0) {
											v83 = v2
											m.G0 = v9 + int32(16)
											return v83
										} else {
											v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											v40 = F_ArrayGetNItemsSafe(m, v39, v27)
											mBase = m.M
											v41 = m.ExcPending
											if v41 != 0 {
												return int32(0)
											} else {
												v42 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v42)
												v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
												if v44 != 0 {
													v52 = v44
												} else {
													v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
													v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v52+v12, v40, v9+int32(15))
												mBase = m.M
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
												v58 = F_ArrayGetNItemsSafe(m, v57, v34)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return int32(0)
												} else {
													v60 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v60)
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
													if v62 != 0 {
														v70 = v62
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
														v70 = (v63<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v70+v17, v58, v9+int32(14))
													mBase = m.M
													v75 = F_inner_int_overlap(m, v12, v17)
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v12)
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v17)
															mBase = m.M
															v80 = m.ExcPending
															if v80 != 0 {
																return int32(0)
															} else {
																v83 = v75
																m.G0 = v9 + int32(16)
																return v83
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
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v27 = v12 + int32(16)
					v28 = F_ArrayGetNItemsSafe(m, v25, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						if v28 == int32(0) {
							v83 = v2
							m.G0 = v9 + int32(16)
							return v83
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v34 = v17 + int32(16)
							v35 = F_ArrayGetNItemsSafe(m, v32, v34)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									v83 = v2
									m.G0 = v9 + int32(16)
									return v83
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
									v40 = F_ArrayGetNItemsSafe(m, v39, v27)
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v42 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v42)
										v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										if v44 != 0 {
											v52 = v44
										} else {
											v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
											v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v52+v12, v40, v9+int32(15))
										mBase = m.M
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
										v58 = F_ArrayGetNItemsSafe(m, v57, v34)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v60)
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
											if v62 != 0 {
												v70 = v62
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
												v70 = (v63<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v70+v17, v58, v9+int32(14))
											mBase = m.M
											v75 = F_inner_int_overlap(m, v12, v17)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v12)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v17)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														v83 = v75
														m.G0 = v9 + int32(16)
														return v83
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
