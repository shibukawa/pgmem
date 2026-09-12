package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DeleteSharedSecurityLabel(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v7+int32(48), int32(2), int32(3), int32(184), l1)
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_table_open(m, int32(3592), int32(3))
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = F_systable_beginscan(m, v23, int32(3593), int32(1), int32(0), int32(2), v7)
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = F_systable_getnext(m, v29)
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = v31
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v29)
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	F_CatalogTupleDelete(m, v23, v33+int32(4))
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v41 = F_systable_getnext(m, v29)
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v41 != 0 {
		v33 = v41
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_sequence_close(m, v23, int32(3))
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	m.G0 = v7 + int32(96)
	return
}
func F_checkSharedDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(224)
	m.G0 = v18
	if l0 == int32(2613) {
		v33 = v5
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v222 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L2:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(int32(11999)) < base.Ui32(l1) {
		v33 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v33 = (base.B2i32(l0 != int32(2615)) | base.B2i32(l1 != int32(2200))) & base.B2i32(l0 != int32(1262))
	goto L3
L6:
	;
	v37 = F_palloc(m, int32(2560))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = l0
	F_errstart_cold(m, int32(21), v295)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L66
	}
L9:
	;
	return int32(0)
L10:
	;
	F_initStringInfo(m, v18+int32(100))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_initStringInfo(m, v18+int32(84))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v51 = F_table_open(m, int32(1214), int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_ScanKeyInit(m, v18+int32(128), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_ScanKeyInit(m, v18+int32(176), int32(6), int32(3), int32(184), l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v73 = F_systable_beginscan(m, v51, int32(1233), int32(1), int32(0), int32(2), v18+int32(128))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	F_systable_endscan(m, v73)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L49
	}
L17:
	;
	v75 = F_systable_getnext(m, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	if v75 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v220 = v5
	v222 = v5
	v225 = v37
	goto L16
L20:
	;
	goto L21
L21:
	;
	v81 = v75
	v85 = v5
	v87 = v5
	v90 = v37
	v91 = int32(128)
	goto L22
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+22)))
	v97 = v95 + v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v104 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v220 = v203
	v222 = v205
	v225 = v208
	goto L16
L24:
	;
	v213 = F_systable_getnext(m, v73)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L47
	}
L25:
	;
	v190 = F_palloc(m, int32(8))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L45
	}
L26:
	;
	v147 = int32(0)
	goto L39
L27:
	;
	if v91 <= v85 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	if v104 == v108 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v87 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v112 <= int32(0) {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v115 = int32(0)
	if v115 < v112 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v118 = v112
	goto L34
L33:
	;
	v118 = v115
	goto L34
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	goto L26
L35:
	;
	v124 = F_repalloc(m, v90, v91*int32(40))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	v128 = v90
	v129 = v91
	goto L37
L37:
	;
	v132 = v128 + v85*int32(20)
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int64)(unsafe.Add(mBase, uint32(v132))) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+8)) = v135
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v132)+12)) = uint8(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v141 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+16)) = base.B2i32(v139 != v141)
	v203 = v85 + int32(1)
	v205 = v87
	v208 = v128
	v209 = v129
	goto L24
L38:
	;
	v128 = v124
	v129 = v91 << (uint(int32(1)) % 32)
	goto L37
L39:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v119+v147<<(uint(int32(2))%32))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v104 != v165 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v170 + int32(1)
	v203 = v85
	v205 = v87
	v208 = v90
	v209 = v91
	goto L24
L41:
	;
	v168 = v147 + int32(1)
	if v118 != v168 {
		v147 = v168
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L25
L45:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v192
	v196 = F_lappend(m, v87, v190)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v203 = v85
	v205 = v196
	v208 = v90
	v209 = v91
	goto L24
L47:
	;
	if v213 != 0 {
		v81 = v213
		v85 = v203
		v87 = v205
		v90 = v208
		v91 = v209
		goto L22
	} else {
		goto L48
	}
L48:
	;
	goto L23
L49:
	;
	F_sequence_close(m, v51, int32(1))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	if int32(2) <= v220 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v247 = int32(0)
	v250 = v247
	v251 = v247
	v259 = v247
	goto L57
L52:
	;
	F_pg_qsort(m, v225, v220, int32(20), int32(475))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v241 = int32(0)
	if v220 != int32(1) {
		v320 = v241
		v329 = v241
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L51
L56:
	;
	goto L51
L57:
	;
	if v250 <= int32(99) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v320 = v280
	v329 = v282
	goto L1
L59:
	;
	v287 = v225 + v251*int32(20)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	v289 = int32(*(*int8)(unsafe.Add(mBase, uint32(v287)+12)))
	F_storeObjectDescription(m, v18+int32(84), v288, v287, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L9
	} else {
		goto L64
	}
L60:
	;
	v271 = v225 + v251*int32(20)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+16))
	v273 = int32(*(*int8)(unsafe.Add(mBase, uint32(v271)+12)))
	F_storeObjectDescription(m, v18+int32(100), v272, v271, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v280 = v250
	v282 = v259 + int32(1)
	goto L59
L63:
	;
	v280 = v250 + int32(1)
	v282 = v259
	goto L59
L64:
	;
	v293 = v251 + int32(1)
	if v293 != v220 {
		v250 = v280
		v251 = v293
		v259 = v282
		goto L57
	} else {
		goto L65
	}
L65:
	;
	goto L58
L66:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v309 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v309
	F_errmsg(m, int32(290259), v18)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(501120), int32(704), int32(169069))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
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
	F_pfree(m, v225)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L9
	} else {
		goto L107
	}
L72:
	;
	v444 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v338 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v339 <= v338 {
		v444 = v338
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v343 = v320
	v348 = v338
	v349 = int32(0)
	goto L76
L76:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358+v349<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(1262)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v365
	if v343 <= int32(99) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v444 = v405
	goto L71
L78:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v412 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L93
	}
L79:
	;
	v372 = v343 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v377 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L9
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v404 = v343
	v405 = v348 + int32(1)
	goto L78
L82:
	;
	if v377 == int32(0) {
		v404 = v372
		v405 = v348
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v381 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_appendStringInfoChar(m, v18+int32(100), int32(10))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L9
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v373
	if v373 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	v395 = int32(185642)
	goto L90
L89:
	;
	v395 = int32(185658)
	goto L90
L90:
	;
	F_appendStringInfo(m, v18+int32(100), v395, v18-int32(-64))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	F_pfree(m, v377)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	v404 = v372
	v405 = v348
	goto L78
L93:
	;
	if v412 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v414 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v436 = v349 + int32(1)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v436 < v437 {
		v343 = v404
		v348 = v405
		v349 = v436
		goto L76
	} else {
		goto L106
	}
L97:
	;
	F_appendStringInfoChar(m, v18+int32(84), int32(10))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L9
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v408
	if v408 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L99
L101:
	;
	v428 = int32(185642)
	goto L103
L102:
	;
	v428 = int32(185658)
	goto L103
L103:
	;
	F_appendStringInfo(m, v18+int32(84), v428, v18+int32(48))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	F_pfree(m, v412)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	goto L96
L106:
	;
	goto L77
L107:
	;
	F_list_free_deep(m, v222)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v458 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	m.G0 = v18 + int32(224)
	return base.B2i32(v458 != int32(0))
L110:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	F_pfree(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L9
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if int32(0) < v329 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	F_pfree(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	v467 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v467
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v467
	goto L109
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v329
	if v329 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	if int32(0) < v444 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v480 = int32(674013)
	goto L120
L119:
	;
	v480 = int32(674060)
	goto L120
L120:
	;
	F_appendStringInfo(m, v18+int32(100), v480, v18+int32(32))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L9
	} else {
		goto L121
	}
L121:
	;
	goto L117
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v444
	if v444 == int32(1) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v499
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v501
	goto L109
L125:
	;
	v494 = int32(674169)
	goto L127
L126:
	;
	v494 = int32(674108)
	goto L127
L127:
	;
	F_appendStringInfo(m, v18+int32(100), v494, v18+int32(16))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	goto L124
}
func F_shared_record_table_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v5 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v15 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_dsa_get_address(m, l3, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v13
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v14 = v9
	goto L1
L7:
	;
	v23 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v27 != v28 {
		v79 = v23
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_dsa_get_address(m, l3, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = v21
	goto L7
L11:
	;
	v22 = v19
	goto L7
L12:
	;
	return v79 ^ int32(1)
L13:
	;
	goto L12
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v30 != v31 {
		v79 = v23
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v27 <= int32(0) {
		v79 = int32(1)
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v37 = v27 << (uint(int32(4)) % 32)
	v39 = int32(20)
	v45 = int32(0)
	goto L17
L17:
	;
	v52 = v45 * int32(100)
	v53 = v14 + v37 + v39 + v52
	v54 = int32(4)
	v56 = v52 + (v22 + v37 + v39)
	v59 = F_strcmp(m, v53+v54, v56+v54)
	mBase = m.M
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v79 = int32(0)
	goto L13
L19:
	;
	goto L18
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+68))
	if v60 != v61 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	if v63 != v64 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
	if v66 != v67 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+91)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+91)))
	if v69 != v70 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v72 = int32(1)
	v74 = v45 + v72
	if v27 != v74 {
		v45 = v74
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v79 = v72
	goto L13
}
func F_shared_ts_free_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	v4 = int32(0)
	F_check_stack_depth(m)
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_dsa_get_address(m, v11, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v170, l1)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L53
	}
L4:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if v128 == int32(0) {
		goto L3
	} else {
		goto L42
	}
L5:
	;
	v95 = v4
	goto L31
L6:
	;
	v59 = v4
	goto L20
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	switch v14 {
	case 0:
		goto L4
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	default:
		goto L3
	}
L9:
	;
	v28 = v4
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v28<<(uint(int32(2))%32))))
	if base.B2i32(l2 <= int32(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L3
L12:
	;
	v46 = v28 + int32(1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if base.Ui32(v46) < base.Ui32(v47) {
		v28 = v46
		goto L10
	} else {
		goto L19
	}
L13:
	;
	F_shared_ts_free_recurse(m, l0, v35, l2-int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v35&int32(1) != 0 {
		goto L12
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v42, v35)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	goto L11
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+(v12+int32(12))))))
	if v64 == int32(255) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L3
L22:
	;
	v82 = v59 + int32(1)
	if v82 != int32(256) {
		v59 = v82
		goto L20
	} else {
		goto L30
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(268)+v64<<(uint(int32(2))%32))))
	if int32(0) < l2 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_shared_ts_free_recurse(m, l0, v70, l2-int32(8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v70&int32(1) != 0 {
		goto L22
	} else {
		goto L28
	}
L27:
	;
	goto L22
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v77, v70)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L22
L30:
	;
	goto L21
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)+int32(base.Ui32(v95)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v104)>>(uint(v95)%32))&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L3
L33:
	;
	v125 = v95 + int32(1)
	if v125 != int32(256) {
		v95 = v125
		goto L31
	} else {
		goto L41
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v95<<(uint(int32(2))%32))))
	if int32(0) < l2 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_shared_ts_free_recurse(m, l0, v113, l2-int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v113&int32(1) != 0 {
		goto L33
	} else {
		goto L39
	}
L38:
	;
	goto L33
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v120, v113)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	goto L32
L42:
	;
	v131 = int32(8)
	v141 = v4
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v12+v131+v141<<(uint(int32(2))%32))))
	if base.B2i32(l2 <= int32(0)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L3
L45:
	;
	v159 = v141 + int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if base.Ui32(v159) < base.Ui32(v160) {
		v141 = v159
		goto L43
	} else {
		goto L52
	}
L46:
	;
	F_shared_ts_free_recurse(m, l0, v148, l2-v131)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v148&int32(1) != 0 {
		goto L45
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v155, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	goto L44
L53:
	;
	return
}
