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
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
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
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = l0
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0]))
	if v18 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L10
	} else {
		goto L142
	}
L2:
	;
	v530 = int32(0)
	v535 = F_hash_search(m, v524, v14+int32(24), v530, v530)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L10
	} else {
		goto L138
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[1]))
	v524 = v22
	goto L2
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2]))
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0])) = int32(1)
	v53 = int32(_a_F_EventCacheLookup_0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[3]))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[3])) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = int64(34359738372)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v57
	v67 = F_hash_create(m, int32(_a_F_EventCacheLookup_1), int32(32), v14+int32(40), int32(1064))
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
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[4]))
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
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_EventCacheLookup_2), int32(0), int32(_a_F_EventCacheLookup_3), int32(_a_F_EventCacheLookup_4))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[4]))
	v36 = v35
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2])) = v41
	F_CacheRegisterSyscacheCallback(m, int32(26), int32(1577), int32(0))
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
	v501 = m.ExcPending
	if v501 != 0 {
		goto L10
	} else {
		goto L134
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
	v487 = F_systable_getnext_ordered(m, v80, int32(1))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L10
	} else {
		goto L132
	}
L29:
	;
	v102 = int32(0)
	v104 = v98 + int32(68)
	v105 = int32(_a_F_EventCacheLookup_5)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventCacheLookup[5])))
	if base.B2i32(v108 == v102)|base.B2i32(v108 != v111) != 0 {
		v129 = v108
		v130 = v111
		goto L32
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v252
	v255 = F_palloc0(m, int32(12))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L71
	}
L31:
	;
	if v129-v130 == int32(0) {
		v252 = v102
		goto L30
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	v114 = v104
	v115 = v105
	goto L34
L34:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v119 == int32(0) {
		v129 = v119
		v130 = v118
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v129 = v119
	v130 = v118
	goto L32
L36:
	;
	v122 = int32(1)
	if v119 == v118 {
		v114 = v114 + v122
		v115 = v115 + v122
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v135 = int32(_a_F_EventCacheLookup_6)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventCacheLookup[6])))
	if base.B2i32(v138 == int32(0))|base.B2i32(v138 != v141) != 0 {
		v159 = v138
		v160 = v141
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v159-v160 == int32(0) {
		v252 = int32(1)
		goto L30
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	v144 = v104
	v145 = v135
	goto L42
L42:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	if v149 == int32(0) {
		v159 = v149
		v160 = v148
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v159 = v149
	v160 = v148
	goto L40
L44:
	;
	v152 = int32(1)
	if v149 == v148 {
		v144 = v144 + v152
		v145 = v145 + v152
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v165 = int32(_a_F_EventCacheLookup_7)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventCacheLookup[7])))
	if base.B2i32(v168 == int32(0))|base.B2i32(v168 != v171) != 0 {
		v189 = v168
		v190 = v171
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v189-v190 == int32(0) {
		v252 = int32(2)
		goto L30
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	v174 = v104
	v175 = v165
	goto L50
L50:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	if v179 == int32(0) {
		v189 = v179
		v190 = v178
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v189 = v179
	v190 = v178
	goto L48
L52:
	;
	v182 = int32(1)
	if v179 == v178 {
		v174 = v174 + v182
		v175 = v175 + v182
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v195 = int32(_a_F_EventCacheLookup_8)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventCacheLookup[8])))
	if base.B2i32(v198 == int32(0))|base.B2i32(v198 != v201) != 0 {
		v219 = v198
		v220 = v201
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v219-v220 == int32(0) {
		v252 = int32(3)
		goto L30
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v204 = v104
	v205 = v195
	goto L58
L58:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	if v209 == int32(0) {
		v219 = v209
		v220 = v208
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v219 = v209
	v220 = v208
	goto L56
L60:
	;
	v212 = int32(1)
	if v209 == v208 {
		v204 = v204 + v212
		v205 = v205 + v212
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v224 = int32(_a_F_EventCacheLookup_9)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventCacheLookup[9])))
	if base.B2i32(v227 == int32(0))|base.B2i32(v227 != v230) != 0 {
		v248 = v227
		v249 = v230
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v248-v249 != 0 {
		goto L28
	} else {
		goto L70
	}
L64:
	;
	goto L63
L65:
	;
	v233 = v104
	v234 = v224
	goto L66
L66:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if v238 == int32(0) {
		v248 = v238
		v249 = v237
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v248 = v238
	v249 = v237
	goto L64
L68:
	;
	v241 = int32(1)
	if v238 == v237 {
		v233 = v233 + v241
		v234 = v234 + v241
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v252 = int32(4)
	goto L30
L71:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v98)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v257
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+140)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+4)) = uint8(v259)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+18)))
	if base.Ui32(v263&int32(2047)) <= base.Ui32(int32(6)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+35)))
	if v325 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L73:
	;
	v271 = F_getmissingattr(m, v261, int32(7), v14+int32(35))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L10
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+35)) = uint8(v273)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+20)))
	if v275&int32(1) == v273 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v324 = v271
	goto L72
L77:
	;
	v319 = F_nocachegetattr(m, v85, int32(7), v261)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L10
	} else {
		goto L91
	}
L78:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v261)+116))
	if v280 < int32(0) {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+23)))
	if v311&int32(64) != 0 {
		goto L77
	} else {
		goto L90
	}
L81:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+22)))
	v285 = v262 + v283 + v280
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+122)))
	if v286 != int32(1) {
		v324 = v285
		goto L72
	} else {
		goto L82
	}
L82:
	;
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+120)))
	switch v289 - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L85
	default:
		goto L83
	case 3:
		goto L84
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L10
	} else {
		goto L87
	}
L84:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v324 = v294
	goto L72
L85:
	;
	v293 = int32(*(*int16)(unsafe.Add(mBase, uint32(v285))))
	v324 = v293
	goto L72
L86:
	;
	v292 = int32(*(*int8)(unsafe.Add(mBase, uint32(v285))))
	v324 = v292
	goto L72
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = base.I32_extend16_s(v289)
	F_errmsg_internal(m, int32(_a_F_EventCacheLookup_10), v14+int32(16))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_EventCacheLookup_11), int32(70), int32(_a_F_EventCacheLookup_12))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L10
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
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+35)) = uint8(v314)
	v324 = int32(0)
	goto L72
L91:
	;
	v324 = v319
	goto L72
L92:
	;
	v328 = F_pg_detoast_datum(m, v324)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L10
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v458 = F_hash_search(m, v67, v14+int32(36), int32(1), v14+int32(34))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L10
	} else {
		goto L126
	}
L95:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v330 != int32(1) {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	if v333 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	if v334 != int32(25) {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v337 = int32(0)
	F_deconstruct_array_builtin(m, v328, int32(25), v14+int32(92), v337, v14+int32(88))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	v346 = int32(0)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v346 < v347 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v350 = v346
	v352 = v337
	goto L103
L101:
	;
	v429 = v337
	goto L102
L102:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	F_pfree(m, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L10
	} else {
		goto L125
	}
L103:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361+v350<<(uint(int32(2))%32))))
	v366 = F_text_to_cstring(m, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L105
	}
L104:
	;
	v429 = v419
	goto L102
L105:
	;
	if v366 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v419 = F_bms_add_member(m, v352, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L10
	} else {
		goto L122
	}
L107:
	;
	v418 = int32(0)
	goto L106
L108:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v374 == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v380 = int32(_a_F_EventCacheLookup_13)
	v381 = int32(_a_F_EventCacheLookup_14)
	goto L110
L110:
	;
	v389 = v380 + (v381-v380)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v391 = F_pg_strcasecmp(m, v366, v390)
	mBase = m.M
	if v391 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L107
L112:
	;
	v418 = (v389 - int32(_a_F_EventCacheLookup_13)) >> (uint(int32(3)) % 32)
	goto L106
L113:
	;
	goto L114
L114:
	;
	v401 = base.B2i32(v391 < int32(0))
	if v391 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v402 = v389 - int32(8)
	goto L117
L116:
	;
	v402 = v381
	goto L117
L117:
	;
	if v391 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v405 = v380
	goto L120
L119:
	;
	v405 = v389 + int32(8)
	goto L120
L120:
	;
	if base.Ui32(v405) <= base.Ui32(v402) {
		v380 = v405
		v381 = v402
		goto L110
	} else {
		goto L121
	}
L121:
	;
	goto L111
L122:
	;
	F_pfree(m, v366)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	v424 = v350 + int32(1)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v424 < v425 {
		v350 = v424
		v352 = v419
		goto L103
	} else {
		goto L124
	}
L124:
	;
	goto L104
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+8)) = v429
	goto L94
L126:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+34)))
	if v460 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v458)+4))
	v464 = F_lappend(m, v463, v255)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L10
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v255
	v472 = F_list_make1_impl(m, int32(1), v14+int32(12))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L10
	} else {
		goto L131
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458)+4)) = v464
	goto L28
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458)+4)) = v472
	goto L28
L132:
	;
	if v487 != 0 {
		v85 = v487
		goto L26
	} else {
		goto L133
	}
L133:
	;
	goto L27
L134:
	;
	F_relation_close(m, v75, int32(1))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L10
	} else {
		goto L135
	}
L135:
	;
	F_relation_close(m, v71, int32(1))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L10
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[1])) = v67
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[3])) = v54
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0]))
	if v513 != int32(1) {
		v524 = v67
		goto L2
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0])) = int32(2)
	v524 = v67
	goto L2
L138:
	;
	if v535 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v538 = v537
	goto L141
L140:
	;
	v538 = v530
	goto L141
L141:
	;
	m.G0 = v14 + int32(96)
	return v538
L142:
	;
	F_errmsg_internal(m, int32(_a_F_EventCacheLookup_15), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L10
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_EventCacheLookup_16), int32(231), int32(_a_F_EventCacheLookup_17))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L10
	} else {
		goto L144
	}
L144:
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v107 int64
	_ = v107
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
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
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v209 int64
	_ = v209
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v239 int32
	_ = v239
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
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[0]))
	v22 = F_AllocSetContextCreateInternal(m, v17, int32(_a_F_EventTriggerInvoke_0), int32(0), int32(_a_F_EventTriggerInvoke_1), int32(_a_F_EventTriggerInvoke_2))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(_a_F_EventTriggerInvoke_3)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[0])) = v22
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[0])) = v25
	F_MemoryContextDelete(m, v22)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_EventTriggerInvoke_4), v12+int32(16))
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
	v51 = v12 - int32(-64)
	F_fmgr_info(m, v32, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	F_errfinish(m, int32(_a_F_EventTriggerInvoke_5), int32(1095), int32(_a_F_EventTriggerInvoke_6))
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v51
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)) = uint8(v54)
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = int64(0)
	v63 = v12 + int32(92)
	v65 = v12 + int32(32)
	F_pgstat_init_function_usage(m, v63, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, v63)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v80 = m.G0
	v82 = v80 - int32(16)
	m.G0 = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v84 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	F_MemoryContextReset(m, v22)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L22
	}
L16:
	;
	F___clock_gettime(m, int32(1), v82)
	mBase = m.M
	v87 = int32(_a_F_EventTriggerInvoke_7)
	v88 = *(*int64)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[1]))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
	v91 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
	v97 = v91 + v92*int64(1000000000) - v96
	*(*int64)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[1])) = v90 + v97
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	goto L19
L17:
	;
	goto L18
L18:
	;
	m.G0 = v82 + int32(16)
	goto L15
L19:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v102 + int64(1)
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v100 + v97
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v84)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v84)+16)) = v107 + (v97 - v88 + v90)
	goto L18
L22:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v121 < int32(2) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v127 = int32(1)
	goto L24
L24:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v127<<(uint(int32(2))%32))))
	v140 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L4
L26:
	;
	if v140 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v137
	F_errmsg_internal(m, int32(_a_F_EventTriggerInvoke_4), v12)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
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
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(_a_F_EventTriggerInvoke_5), int32(1095), int32(_a_F_EventTriggerInvoke_6))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v154 = v12 - int32(-64)
	F_fmgr_info(m, v137, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v157 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+110)) = uint16(v157)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = v154
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)) = uint8(v157)
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = int64(0)
	v166 = v12 + int32(92)
	v168 = v12 + int32(32)
	F_pgstat_init_function_usage(m, v166, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v173 = m.T0[v172].(func(*base.Module, int32) int32)(m, v166)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v182 = m.G0
	v184 = v182 - int32(16)
	m.G0 = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v186 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_MemoryContextReset(m, v22)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L43
	}
L37:
	;
	F___clock_gettime(m, int32(1), v184)
	mBase = m.M
	v189 = int32(_a_F_EventTriggerInvoke_7)
	v190 = *(*int64)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[1]))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v168)+16))
	v193 = int64(*(*int32)(unsafe.Add(mBase, uint32(v184)+8)))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v184)))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v168)+24))
	v199 = v193 + v194*int64(1000000000) - v198
	*(*int64)(unsafe.Add(mBase, _c_F_EventTriggerInvoke[1])) = v192 + v199
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
	goto L40
L38:
	;
	goto L39
L39:
	;
	m.G0 = v184 + int32(16)
	goto L36
L40:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
	*(*int64)(unsafe.Add(mBase, uint32(v186))) = v204 + int64(1)
	goto L42
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v186)+8)) = v202 + v199
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v186)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v186)+16)) = v209 + (v199 - v190 + v192)
	goto L39
L43:
	;
	v224 = v127 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v224 < v225 {
		v127 = v224
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[0]))
	v16 = F_LWLockAcquire(m, v12+int32(_a_F_GetWaitEventCustomNames_0), int32(1))
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
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[1]))
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
	v91 = F_palloc(m, v88<<(uint(int32(2))%32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v88 = v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + v24))))))))))))))))))))))))))))))
	goto L6
L5:
	;
	v88 = v24
	goto L6
L6:
	;
	goto L3
L7:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[1]))
	F_hash_seq_init(m, v9+int32(12), v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v102 = int32(0)
	goto L9
L9:
	;
	v107 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[0]))
	F_LWLockRelease(m, v122+int32(_a_F_GetWaitEventCustomNames_0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	if v107 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+67)))
	if v109<<(uint(int32(24))%32) != l0 {
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
	v116 = F_pstrdup(m, v107)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v102<<(uint(int32(2))%32)))) = v116
	v102 = v102 + int32(1)
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v102
	m.G0 = v9 + int32(32)
	return v91
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v2 = m.G0
	v4 = v2 + int32(-64)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[0])))
	if v7 != int32(1) {
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[1]))
		if v11 == int32(0) {
		} else {
			v14 = int32(_a_F_InitializeWaitEventSupport_0)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[2]))
			v16 = F_close(m, v15)
			mBase = m.M
			v17 = int32(_a_F_InitializeWaitEventSupport_1)
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[3]))
			v19 = F_close(m, v18)
			mBase = m.M
			v21 = int32(-1)
			*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[2])) = v21
			*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[3])) = v21
			*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[1])) = int32(0)
			v29 = int32(_a_F_InitializeWaitEventSupport_2)
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[4]))
			*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[4])) = v31 - int32(1)
			v35 = int32(_a_F_InitializeWaitEventSupport_2)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[4]))
			*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[4])) = v37 - int32(1)
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
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[2])) = v48
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v4)+60))
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[3])) = v62
		v66 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[5]))
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[1])) = v66
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
				v73 = int32(1110)
				v75 = m.G0
				v77 = v75 - int32(32)
				m.G0 = v77
				switch int32(1112) {
				case 0, 2:
					v87 = v73
				default:
					*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[6])) = v73
					v87 = int32(_a_F_InitializeWaitEventSupport_3)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v87
				F_sigemptyset(m, v77+int32(16))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = int32(268435456)
				v99 = v77 + int32(12)
				if v99 != 0 {
					v106 = int32(460)
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
					*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[7])) = v107
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v99)+8))
					*(*int64)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[8])) = v109
					v111 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
					*(*int64)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[9])) = v111
				} else {
				}
				m.G0 = v77 + int32(32)
				m.G0 = v4 - int32(-64)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v124 = m.ExcPending
		if v124 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_InitializeWaitEventSupport_4), int32(0))
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_InitializeWaitEventSupport_5), int32(296), int32(_a_F_InitializeWaitEventSupport_6))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
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
