package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EventCacheLookup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int64
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v300 int64
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int64
	_ = v328
	var v329 int32
	_ = v329
	var v333 int64
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0]))
	if v20 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L10
	} else {
		goto L151
	}
L2:
	;
	v555 = int32(0)
	v560 = F_hash_search(m, v547, v16+int32(24), v555, v555)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L10
	} else {
		goto L147
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[1]))
	v547 = v24
	goto L2
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2]))
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0])) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(34359738372)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v58
	v65 = F_hash_create(m, int32(_a_F_EventCacheLookup_0), int64(32), v16+int32(40), int32(1064))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L18
	}
L7:
	;
	F_MemoryContextReset(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[3]))
	if v33 != 0 {
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
	v38 = v33
	goto L14
L13:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v43 = F_AllocSetContextCreateInternal(m, v38, int32(_a_F_EventCacheLookup_1), int32(0), int32(_a_F_EventCacheLookup_2), int32(_a_F_EventCacheLookup_3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[3]))
	v38 = v37
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2])) = v43
	F_CacheRegisterSyscacheCallback(m, int32(26), int32(1790), int64(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L6
L18:
	;
	v69 = F_relation_open(m, int32(3466), int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v73 = F_index_open(m, int32(3467), int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v75 = int32(0)
	v78 = F_systable_beginscan_ordered(m, v69, v73, v75, v75, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v81 = F_systable_getnext_ordered(m, v78, int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v83 = v81
	goto L26
L24:
	;
	goto L25
L25:
	;
	F_systable_endscan_ordered(m, v78)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L10
	} else {
		goto L143
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
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
	v510 = F_systable_getnext_ordered(m, v78, int32(1))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L10
	} else {
		goto L141
	}
L29:
	;
	v102 = int32(0)
	v104 = v98 + int32(68)
	v105 = int32(_a_F_EventCacheLookup_4)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_EventCacheLookup[4])))
	if base.B2i32(v108 == v102)|base.B2i32(v108 != v111) != 0 {
		v129 = v108
		v130 = v111
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v253 = int32(_a_F_EventCacheLookup_5)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[5]))
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[5])) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v252
	v261 = F_palloc0(m, int32(12))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
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
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v98)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v263
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+140)))
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+4)) = uint8(v265)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268)+18)))
	if base.Ui32(v269&int32(2047)) <= base.Ui32(int32(6)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+35)))
	if v334 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L73:
	;
	v277 = F_getmissingattr(m, v267, int32(7), v16+int32(35))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L10
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+35)) = uint8(v279)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+20)))
	if v281&int32(1) == v279 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v333 = v277
	goto L72
L77:
	;
	v328 = F_nocachegetattr(m, v83, int32(7), v267)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L10
	} else {
		goto L95
	}
L78:
	;
	v286 = int32(*(*int16)(unsafe.Add(mBase, uint32(v267)+76)))
	if v286 < int32(0) {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+23)))
	if v320&int32(64) != 0 {
		goto L77
	} else {
		goto L94
	}
L81:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+22)))
	v291 = v268 + v289 + v286
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+80)))
	if v292 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v295 = int32(*(*int16)(unsafe.Add(mBase, uint32(v267)+78)))
	if base.I32_popcnt(v295) != int32(1) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v333 = base.I64_extend_i32_u(v291)
	goto L72
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L10
	} else {
		goto L91
	}
L86:
	;
	switch base.I32_ctz(v295) {
	case 0:
		goto L90
	case 1:
		goto L89
	case 2:
		goto L88
	case 3:
		goto L87
	default:
		goto L85
	}
L87:
	;
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v291)))
	v333 = v303
	goto L72
L88:
	;
	v302 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v291))))
	v333 = v302
	goto L72
L89:
	;
	v301 = int64(*(*int16)(unsafe.Add(mBase, uint32(v291))))
	v333 = v301
	goto L72
L90:
	;
	v300 = int64(*(*int8)(unsafe.Add(mBase, uint32(v291))))
	v333 = v300
	goto L72
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v295
	F_errmsg_internal(m, int32(_a_F_EventCacheLookup_10), v16+int32(16))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_EventCacheLookup_11), int32(123), int32(_a_F_EventCacheLookup_12))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v323 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+35)) = uint8(v323)
	v333 = int64(0)
	goto L72
L95:
	;
	v333 = v328
	goto L72
L96:
	;
	v337 = base.I32_wrap_i64(v333)
	v338 = F_pg_detoast_datum(m, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L10
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v477 = F_hash_search(m, v65, v16+int32(36), int32(1), v16+int32(34))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L10
	} else {
		goto L134
	}
L99:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v340 != int32(1) {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	if v343 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	if v344 != int32(25) {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v347 = int32(0)
	F_deconstruct_array_builtin(m, v338, int32(25), v16+int32(92), v347, v16+int32(88))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v356 = int32(0)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	if v356 < v357 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v360 = v347
	v364 = v356
	goto L107
L105:
	;
	v443 = v356
	goto L106
L106:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	F_pfree(m, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L10
	} else {
		goto L129
	}
L107:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v16)+92))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373+v360<<(uint(int32(3))%32))))
	v378 = F_text_to_cstring(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L10
	} else {
		goto L109
	}
L108:
	;
	v443 = v431
	goto L106
L109:
	;
	if v378 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v431 = F_bms_add_member(m, v364, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L126
	}
L111:
	;
	v430 = int32(0)
	goto L110
L112:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v386 == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v392 = int32(_a_F_EventCacheLookup_13)
	v393 = int32(_a_F_EventCacheLookup_14)
	goto L114
L114:
	;
	v401 = v392 + (v393-v392)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v403 = F_pg_strcasecmp(m, v378, v402)
	mBase = m.M
	if v403 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	goto L111
L116:
	;
	v430 = (v401 - int32(_a_F_EventCacheLookup_13)) >> (uint(int32(3)) % 32)
	goto L110
L117:
	;
	goto L118
L118:
	;
	v413 = base.B2i32(v403 < int32(0))
	if v403 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v414 = v401 - int32(8)
	goto L121
L120:
	;
	v414 = v393
	goto L121
L121:
	;
	if v403 < int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v417 = v392
	goto L124
L123:
	;
	v417 = v401 + int32(8)
	goto L124
L124:
	;
	if base.Ui32(v417) <= base.Ui32(v414) {
		v392 = v417
		v393 = v414
		goto L114
	} else {
		goto L125
	}
L125:
	;
	goto L115
L126:
	;
	F_pfree(m, v378)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	v436 = v360 + int32(1)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	if v436 < v437 {
		v360 = v436
		v364 = v431
		goto L107
	} else {
		goto L128
	}
L128:
	;
	goto L108
L129:
	;
	if v338 != v337 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_pfree(m, v338)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L10
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v443
	goto L98
L133:
	;
	goto L132
L134:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)))
	if v479 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[5])) = v254
	goto L28
L136:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	v483 = F_lappend(m, v482, v261)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L10
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v261
	v491 = F_list_make1_impl(m, int32(1), v16+int32(12))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L10
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v477)+4)) = v483
	goto L135
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v477)+4)) = v491
	goto L135
L141:
	;
	if v510 != 0 {
		v83 = v510
		goto L26
	} else {
		goto L142
	}
L142:
	;
	goto L27
L143:
	;
	F_relation_close(m, v73, int32(1))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L10
	} else {
		goto L144
	}
L144:
	;
	F_relation_close(m, v69, int32(1))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[1])) = v65
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0]))
	if v536 != int32(1) {
		v547 = v65
		goto L2
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_EventCacheLookup[0])) = int32(2)
	v547 = v65
	goto L2
L147:
	;
	if v560 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v563 = v562
	goto L150
L149:
	;
	v563 = v555
	goto L150
L150:
	;
	m.G0 = v16 + int32(96)
	return v563
L151:
	;
	F_errmsg_internal(m, int32(_a_F_EventCacheLookup_15), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_EventCacheLookup_16), int32(232), int32(_a_F_EventCacheLookup_17))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
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
	var v70 int64
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
	var v173 int64
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
	v49 = v12 + int32(96)
	v51 = v12 + int32(60)
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
	F_errfinish(m, int32(_a_F_EventTriggerInvoke_5), int32(1105), int32(_a_F_EventTriggerInvoke_6))
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
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+106)) = uint16(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v51
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)) = uint8(v54)
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = int64(0)
	v63 = v12 + int32(88)
	v65 = v12 + int32(24)
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = m.T0[v69].(func(*base.Module, int32) int64)(m, v63)
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
	F_errfinish(m, int32(_a_F_EventTriggerInvoke_5), int32(1105), int32(_a_F_EventTriggerInvoke_6))
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
	v154 = v12 + int32(60)
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
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+106)) = uint16(v157)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+92)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v154
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+8)) = uint8(v157)
	*(*int64)(unsafe.Add(mBase, uint32(v49))) = int64(0)
	v166 = v12 + int32(88)
	v168 = v12 + int32(24)
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
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v173 = m.T0[v172].(func(*base.Module, int32) int64)(m, v166)
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
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
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[1]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)+8))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v24)+808))
	if v26 != int64(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v93 = F_palloc_mul(m, int32(4), base.I32_wrap_i64(v91))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v24)+752))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v24)+728))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v24)+704))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v24)+680))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v24)+656))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v24)+632))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v24)+608))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v24)+584))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v24)+560))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v24)+536))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)+512))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v24)+488))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v24)+464))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v24)+440))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v24)+416))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v24)+392))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v24)+368))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v24)+344))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v24)+320))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v24)+296))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v24)+272))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v24)+248))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v24)+224))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v24)+200))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v24)+176))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v24)+152))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v24)+128))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v24)+104))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v24)+80))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v24)+56))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v24)+32))
	v91 = v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + v25))))))))))))))))))))))))))))))
	goto L6
L5:
	;
	v91 = v25
	goto L6
L6:
	;
	goto L3
L7:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[1]))
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
	v104 = int32(0)
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
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_GetWaitEventCustomNames[0]))
	F_LWLockRelease(m, v124+int32(_a_F_GetWaitEventCustomNames_0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v93+v104<<(uint(int32(2))%32)))) = v118
	v104 = v104 + int32(1)
	goto L9
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v104
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
				v76 = m.G0
				v78 = v76 - int32(32)
				m.G0 = v78
				v81 = int32(1214)
				switch v81 {
				case 0, 2:
				default:
					*(*int32)(unsafe.Add(mBase, _c_F_InitializeWaitEventSupport[6])) = int32(1212)
				}
				F_sigemptyset(m, v78+int32(16))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = int32(268435456)
				switch v81 {
				case 0:
					*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(-2)
				default:
					*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = int32(268435460)
					*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(_a_F_InitializeWaitEventSupport_3)
				case 2:
					*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(0)
				}
				v110 = F___sigaction(m, int32(23), v78+int32(12), int32(0))
				mBase = m.M
				m.G0 = v78 + int32(32)
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
			F_errmsg_internal(m, int32(_a_F_InitializeWaitEventSupport_4), int32(0))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_InitializeWaitEventSupport_5), int32(297), int32(_a_F_InitializeWaitEventSupport_6))
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
