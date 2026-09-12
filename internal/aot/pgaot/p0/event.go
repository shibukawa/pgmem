package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EventCacheLookup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l0
	v18 = *(*int32)(unsafe.Add(mBase, _consts[868]))
	if v18 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L10
	} else {
		goto L147
	}
L2:
	;
	v528 = int32(0)
	v533 = F_hash_search(m, v522, v14+int32(24), v528, v528)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L10
	} else {
		goto L143
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[869]))
	v522 = v22
	goto L2
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[868])) = int32(1)
	v53 = int32(4442992)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = int64(34359738372)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v57
	v67 = F_hash_create(m, int32(308066), int32(32), v14+int32(40), int32(1064))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L18
	}
L7:
	;
	F_MemoryContextReset(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L6
L12:
	;
	v36 = v31
	goto L14
L13:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(380946), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v36 = v35
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[870])) = v41
	F_CacheRegisterSyscacheCallback(m, int32(26), int32(1592), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L6
L18:
	;
	v71 = F_relation_open(m, int32(3466), int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v75 = F_index_open(m, int32(3467), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v77 = int32(0)
	v80 = F_systable_beginscan_ordered(m, v71, v75, v77, v77, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v83 = F_systable_getnext_ordered(m, v80, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	if v83 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v85 = v83
	goto L26
L24:
	;
	goto L25
L25:
	;
	F_systable_endscan_ordered(m, v80)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L10
	} else {
		goto L139
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v98 = v96 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+140)))
	if v99 == int32(68) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v485 = F_systable_getnext_ordered(m, v80, int32(1))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L10
	} else {
		goto L137
	}
L29:
	;
	v102 = int32(0)
	v104 = v98 + int32(68)
	v105 = int32(78226)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[871])))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v109 == v102 {
		v128 = v108
		v129 = v109
		goto L32
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v247
	v250 = F_palloc0(m, int32(12))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L76
	}
L31:
	;
	if v129-v128 == int32(0) {
		v247 = v102
		goto L30
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	if v108 != v109 {
		v128 = v108
		v129 = v109
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v113 = v104
	v114 = v105
	goto L35
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v118 == int32(0) {
		v128 = v117
		v129 = v118
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v128 = v117
	v129 = v118
	goto L32
L37:
	;
	v121 = int32(1)
	if v117 == v118 {
		v113 = v113 + v121
		v114 = v114 + v121
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v134 = int32(406948)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, _consts[872])))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v138 == int32(0) {
		v157 = v137
		v158 = v138
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v158-v157 == int32(0) {
		v247 = int32(1)
		goto L30
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	if v137 != v138 {
		v157 = v137
		v158 = v138
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v142 = v104
	v143 = v134
	goto L44
L44:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	if v147 == int32(0) {
		v157 = v146
		v158 = v147
		goto L41
	} else {
		goto L46
	}
L45:
	;
	v157 = v146
	v158 = v147
	goto L41
L46:
	;
	v150 = int32(1)
	if v146 == v147 {
		v142 = v142 + v150
		v143 = v143 + v150
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v163 = int32(222815)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, _consts[873])))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v167 == int32(0) {
		v186 = v166
		v187 = v167
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v187-v186 == int32(0) {
		v247 = int32(2)
		goto L30
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	if v166 != v167 {
		v186 = v166
		v187 = v167
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v171 = v104
	v172 = v163
	goto L53
L53:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	if v176 == int32(0) {
		v186 = v175
		v187 = v176
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v186 = v175
	v187 = v176
	goto L50
L55:
	;
	v179 = int32(1)
	if v175 == v176 {
		v171 = v171 + v179
		v172 = v172 + v179
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v192 = int32(332495)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _consts[874])))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v196 == int32(0) {
		v215 = v195
		v216 = v196
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v216-v215 == int32(0) {
		v247 = int32(3)
		goto L30
	} else {
		goto L66
	}
L59:
	;
	goto L58
L60:
	;
	if v195 != v196 {
		v215 = v195
		v216 = v196
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v200 = v104
	v201 = v192
	goto L62
L62:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	if v205 == int32(0) {
		v215 = v204
		v216 = v205
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v215 = v204
	v216 = v205
	goto L59
L64:
	;
	v208 = int32(1)
	if v204 == v205 {
		v200 = v200 + v208
		v201 = v201 + v208
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v220 = int32(262987)
	v223 = int32(*(*uint8)(unsafe.Add(mBase, _consts[875])))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v224 == int32(0) {
		v243 = v223
		v244 = v224
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v244-v243 != 0 {
		goto L28
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	if v223 != v224 {
		v243 = v223
		v244 = v224
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v228 = v104
	v229 = v220
	goto L71
L71:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v233 == int32(0) {
		v243 = v232
		v244 = v233
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v243 = v232
	v244 = v233
	goto L68
L73:
	;
	v236 = int32(1)
	if v232 == v233 {
		v228 = v228 + v236
		v229 = v229 + v236
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v247 = int32(4)
	goto L30
L76:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v98)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+140)))
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+4)) = uint8(v254)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+18)))
	if base.Ui32(v258&int32(2047)) <= base.Ui32(int32(6)) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+35)))
	if v320 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L78:
	;
	v266 = F_getmissingattr(m, v256, int32(7), v14+int32(35))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L10
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+35)) = uint8(v268)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+20)))
	if v270&int32(1) == v268 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v319 = v266
	goto L77
L82:
	;
	v314 = F_nocachegetattr(m, v85, int32(7), v256)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L96
	}
L83:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256)+116))
	if v275 < int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+23)))
	if v306&int32(64) != 0 {
		goto L82
	} else {
		goto L95
	}
L86:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
	v280 = v257 + v278 + v275
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+122)))
	if v281 != int32(1) {
		v319 = v280
		goto L77
	} else {
		goto L87
	}
L87:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+120)))
	switch v284 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	default:
		goto L88
	case 3:
		goto L89
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L10
	} else {
		goto L92
	}
L89:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v319 = v289
	goto L77
L90:
	;
	v288 = int32(*(*int16)(unsafe.Add(mBase, uint32(v280))))
	v319 = v288
	goto L77
L91:
	;
	v287 = int32(*(*int8)(unsafe.Add(mBase, uint32(v280))))
	v319 = v287
	goto L77
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = base.I32_extend16_s(v284)
	F_errmsg_internal(m, int32(460988), v14+int32(16))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(310395), int32(70), int32(64556))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	v309 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+35)) = uint8(v309)
	v319 = int32(0)
	goto L77
L96:
	;
	v319 = v314
	goto L77
L97:
	;
	v323 = F_pg_detoast_datum(m, v319)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L10
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v456 = F_hash_search(m, v67, v14+int32(36), int32(1), v14+int32(34))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L10
	} else {
		goto L131
	}
L100:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v325 != int32(1) {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	if v328 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	if v329 != int32(25) {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v332 = int32(0)
	F_deconstruct_array_builtin(m, v323, int32(25), v14+int32(92), v332, v14+int32(88))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v341 = int32(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v341 < v342 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v345 = v341
	v347 = v332
	goto L108
L106:
	;
	v427 = v332
	goto L107
L107:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	F_pfree(m, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L10
	} else {
		goto L130
	}
L108:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356+v345<<(uint(int32(2))%32))))
	v361 = F_text_to_cstring(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L110
	}
L109:
	;
	v427 = v417
	goto L107
L110:
	;
	v363 = int32(0)
	if v361 == v363 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v417 = F_bms_add_member(m, v347, v416)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L127
	}
L112:
	;
	v416 = v363
	goto L111
L113:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v370 == int32(0) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v376 = int32(1597776)
	v377 = int32(1599312)
	goto L115
L115:
	;
	v386 = v376 + (v377-v376)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v388 = F_pg_strcasecmp(m, v361, v387)
	mBase = m.M
	if v388 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L112
L117:
	;
	v416 = (v386 - int32(1597776)) >> (uint(int32(3)) % 32)
	goto L111
L118:
	;
	goto L119
L119:
	;
	v398 = base.B2i32(v388 < int32(0))
	if v388 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v399 = v386 - int32(8)
	goto L122
L121:
	;
	v399 = v377
	goto L122
L122:
	;
	if v388 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v402 = v376
	goto L125
L124:
	;
	v402 = v386 + int32(8)
	goto L125
L125:
	;
	if base.Ui32(v402) <= base.Ui32(v399) {
		v376 = v402
		v377 = v399
		goto L115
	} else {
		goto L126
	}
L126:
	;
	goto L116
L127:
	;
	F_pfree(m, v361)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	v422 = v345 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v422 < v423 {
		v345 = v422
		v347 = v417
		goto L108
	} else {
		goto L129
	}
L129:
	;
	goto L109
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+8)) = v427
	goto L99
L131:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+34)))
	if v458 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	v462 = F_lappend(m, v461, v250)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L10
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v250
	v470 = F_list_make1_impl(m, int32(1), v14+int32(12))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L10
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = v462
	goto L28
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = v470
	goto L28
L137:
	;
	if v485 != 0 {
		v85 = v485
		goto L26
	} else {
		goto L138
	}
L138:
	;
	goto L27
L139:
	;
	F_relation_close(m, v75, int32(1))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	F_relation_close(m, v71, int32(1))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[869])) = v67
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v54
	v511 = *(*int32)(unsafe.Add(mBase, _consts[868]))
	if v511 != int32(1) {
		v522 = v67
		goto L2
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, _consts[868])) = int32(2)
	v522 = v67
	goto L2
L143:
	;
	if v533 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	v536 = v535
	goto L146
L145:
	;
	v536 = v528
	goto L146
L146:
	;
	m.G0 = v14 + int32(96)
	return v536
L147:
	;
	F_errmsg_internal(m, int32(22964), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(476024), int32(231), int32(98153))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_EventTriggerInvoke(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v180 int32
	_ = v180
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
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v223 int64
	_ = v223
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v253 int32
	_ = v253
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v22 = F_AllocSetContextCreateInternal(m, v17, int32(58083), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(4442992)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	F_MemoryContextDelete(m, v22)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L45
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v35 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v32
	F_errmsg_internal(m, int32(49462), v12+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v49 = v12 + int32(100)
	F_fmgr_info(m, v32, v12-int32(-64))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	F_errfinish(m, int32(472622), int32(1095), int32(379488))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v54 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+110)) = uint16(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v12 - int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)) = uint8(v54)
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = int64(0)
	F_pgstat_init_function_usage(m, v12+int32(92), v12+int32(32))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = m.T0[v73].(func(*base.Module, int32) int32)(m, v12+int32(92))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v78 = v12 + int32(32)
	v86 = m.G0
	v88 = v86 - int32(16)
	m.G0 = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v90 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	F_MemoryContextReset(m, v22)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L22
	}
L16:
	;
	F___clock_gettime(m, int32(1), v88)
	mBase = m.M
	v93 = int32(4422768)
	v94 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v78)+16))
	v97 = int64(*(*int32)(unsafe.Add(mBase, uint32(v88)+8)))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v78)+24))
	v103 = v97 + v98*int64(1000000000) - v102
	*(*int64)(unsafe.Add(mBase, _consts[265])) = v96 + v103
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	goto L19
L17:
	;
	goto L18
L18:
	;
	m.G0 = v88 + int32(16)
	goto L15
L19:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = v108 + int64(1)
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v90)+8)) = v106 + v103
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v90)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v90)+16)) = v113 + (v103 - v94 + v96)
	goto L18
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v127 <= int32(1) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v135 = int32(1)
	goto L24
L24:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v135<<(uint(int32(2))%32))))
	v148 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L4
L26:
	;
	if v148 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v145
	F_errmsg_internal(m, int32(49462), v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(472622), int32(1095), int32(379488))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	F_fmgr_info(m, v145, v12-int32(-64))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v165 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+110)) = uint16(v165)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v12 - int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(108)))) = uint8(v165)
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = int64(0)
	F_pgstat_init_function_usage(m, v12+int32(92), v12+int32(32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v185 = m.T0[v184].(func(*base.Module, int32) int32)(m, v12+int32(92))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v188 = v12 + int32(32)
	v196 = m.G0
	v198 = v196 - int32(16)
	m.G0 = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v200 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_MemoryContextReset(m, v22)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L43
	}
L37:
	;
	F___clock_gettime(m, int32(1), v198)
	mBase = m.M
	v203 = int32(4422768)
	v204 = *(*int64)(unsafe.Add(mBase, _consts[265]))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v188)+16))
	v207 = int64(*(*int32)(unsafe.Add(mBase, uint32(v198)+8)))
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v188)+24))
	v213 = v207 + v208*int64(1000000000) - v212
	*(*int64)(unsafe.Add(mBase, _consts[265])) = v206 + v213
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v188)+8))
	goto L40
L38:
	;
	goto L39
L39:
	;
	m.G0 = v198 + int32(16)
	goto L36
L40:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v200)))
	*(*int64)(unsafe.Add(mBase, uint32(v200))) = v218 + int64(1)
	goto L42
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v200)+8)) = v216 + v213
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v200)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v200)+16)) = v223 + (v213 - v204 + v206)
	goto L39
L43:
	;
	v238 = v135 + int32(1)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v238 < v239 {
		v135 = v238
		goto L24
	} else {
		goto L44
	}
L44:
	;
	goto L25
L45:
	;
	m.G0 = v12 + int32(112)
	return
}
func F_GetWaitEventCustomNames(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v16 = F_LWLockAcquire(m, v12+int32(6144), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+412))
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v93 = F_palloc(m, v90<<(uint(int32(2))%32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+376))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+364))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+352))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+340))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+328))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+316))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+304))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+292))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+280))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+268))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+256))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)+244))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+232))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+220))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+208))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+196))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)+184))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+172))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+160))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+148))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v23)+136))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v23)+124))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+112))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(-64))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v90 = v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v54 + (v55 + (v56 + (v57 + (v58 + v24))))))))))))))))))))))))))))))
	goto L6
L5:
	;
	v90 = v24
	goto L6
L6:
	;
	goto L3
L7:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	F_hash_seq_init(m, v9+int32(12), v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v103 = int32(0)
	goto L9
L9:
	;
	v109 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v124+int32(6144))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v109 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+67)))
	if v111<<(uint(int32(24))%32) != l0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L10
L15:
	;
	v118 = F_pstrdup(m, v109)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93+v103<<(uint(int32(2))%32)))) = v118
	v103 = v103 + int32(1)
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v103
	m.G0 = v9 + int32(32)
	return v93
}
func F_InitializeWaitEventSupport(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v2 = m.G0
	v4 = v2 + int32(-64)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v7 != int32(1) {
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[625]))
		if v11 == int32(0) {
		} else {
			v14 = int32(4062880)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[626]))
			v16 = F_close(m, v15)
			mBase = m.M
			v17 = int32(4062884)
			v18 = *(*int32)(unsafe.Add(mBase, _consts[627]))
			v19 = F_close(m, v18)
			mBase = m.M
			v21 = int32(-1)
			*(*int32)(unsafe.Add(mBase, _consts[626])) = v21
			*(*int32)(unsafe.Add(mBase, _consts[627])) = v21
			*(*int32)(unsafe.Add(mBase, _consts[625])) = int32(0)
			v29 = int32(4358988)
			v31 = *(*int32)(unsafe.Add(mBase, _consts[592]))
			*(*int32)(unsafe.Add(mBase, _consts[592])) = v31 - int32(1)
			v35 = int32(4358988)
			v37 = *(*int32)(unsafe.Add(mBase, _consts[592]))
			*(*int32)(unsafe.Add(mBase, _consts[592])) = v37 - int32(1)
		}
	}
	v43 = F_pipe(m, v2+int32(-8))
	mBase = m.M
	if int32(0) <= v43 {
		v46 = int32(2048)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v46
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v4)+56))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = v46
		v52 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v52
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = v52
		*(*int32)(unsafe.Add(mBase, _consts[626])) = v48
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v4)+60))
		*(*int32)(unsafe.Add(mBase, _consts[627])) = v62
		v66 = *(*int32)(unsafe.Add(mBase, _consts[522]))
		*(*int32)(unsafe.Add(mBase, _consts[625])) = v66
		F_ReserveExternalFD(m)
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			F_ReserveExternalFD(m)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				v73 = int32(1106)
				v75 = m.G0
				v77 = v75 - int32(144)
				m.G0 = v77
				switch int32(1108) {
				case 0, 2:
					v87 = v73
				default:
					*(*int32)(unsafe.Add(mBase, _consts[628])) = v73
					v87 = int32(4729)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v87
				F_sigemptyset(m, v77+int32(8))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v77)+136)) = int32(268435456)
				v99 = v77 + int32(4)
				if v99 != 0 {
					v110 = F___memcpy(m, int32(4611124), v99, int32(140))
					mBase = m.M
				} else {
				}
				m.G0 = v77 + int32(144)
				m.G0 = v4 - int32(-64)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(281090), int32(0))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return
			} else {
				F_errfinish(m, int32(470450), int32(296), int32(76502))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
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
