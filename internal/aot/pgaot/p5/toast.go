package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_toast_table(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(304)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+112))
	if v18 != 0 {
		v433 = v8
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L20
	} else {
		goto L101
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L20
	} else {
		goto L98
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L20
	} else {
		goto L95
	}
L4:
	;
	m.G0 = v15 + int32(304)
	return v433
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_create_toast_table[0])))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l4 != int32(8) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v24 == int32(112) {
		v433 = v8
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[1]))
	if v44 == int32(0) {
		v433 = v8
		goto L4
	} else {
		goto L23
	}
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+117)))
	if v27 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[2]))
	if v31 != 0 {
		v433 = v8
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L15
L14:
	;
	goto L13
L15:
	;
	if base.Ui32(v32) < base.Ui32(int32(_a_F_create_toast_table_0)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[2]))
	if v36 != 0 {
		v433 = v8
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+152))
	v39 = m.T0[v38].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	if v39 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v433 = v8
	goto L4
L23:
	;
	goto L6
L24:
	;
	v50 = l5
	goto L26
L25:
	;
	v50 = int32(0)
	goto L26
L26:
	;
	if v50 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
	v58 = F_pg_snprintf(m, v15+int32(224), int32(64), int32(_a_F_create_toast_table_1), v15+int32(48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v19
	v67 = F_pg_snprintf(m, v15+int32(160), int32(64), int32(_a_F_create_toast_table_2), v15+int32(32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v71 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	F_TupleDescInitEntry(m, v71, int32(1), int32(_a_F_create_toast_table_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	F_TupleDescInitEntry(m, v71, int32(2), int32(_a_F_create_toast_table_4), int32(23), int32(-1), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	F_TupleDescInitEntry(m, v71, int32(3), int32(_a_F_create_toast_table_5), int32(17), int32(-1), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v95 = int32(3)
	v98 = int32(112)
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v94<<(uint(v95)%32))+112)) = uint8(v98)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v103 = v71 + v100<<(uint(v95)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+312)) = uint8(v98)
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+212)) = uint8(v98)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v108<<(uint(v95)%32))+113)) = uint8(v112)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v117 = v71 + v114<<(uint(v95)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+313)) = uint8(v112)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+213)) = uint8(v112)
	F_populate_compact_attribute(m, v71, v112)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	F_populate_compact_attribute(m, v71, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	F_populate_compact_attribute(m, v71, int32(2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v131 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v131 < v140 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+68))
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[3]))
	if v224 != 0 {
		goto L58
	} else {
		goto L59
	}
L38:
	;
	v144 = v71 + int32(28)
	v151 = v131
	v152 = v140
	v154 = v131
	goto L42
L39:
	;
	v208 = v131
	v215 = v140
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v71)+16)) = v208
	goto L37
L41:
	;
	v208 = v202
	v215 = v181
	goto L40
L42:
	;
	v160 = v144 + v140<<(uint(int32(3))%32) + v151*int32(100)
	v163 = v144 + v151<<(uint(int32(3))%32)
	if v140 != v152 {
		v181 = v152
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v202 = v140
	goto L41
L44:
	;
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v163)+2)))
	if v182 <= int32(0) {
		v202 = v151
		goto L41
	} else {
		goto L52
	}
L45:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+7)))
	if v165 != int32(118) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v181 = v151
	goto L44
L47:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	if v168 != int32(1) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+6)))
	if v171&int32(6) != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v163)+2)))
	if v174 <= int32(0) {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+90)))
	if v177 != int32(118) {
		v181 = v140
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L46
L52:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+90)))
	if v185 == int32(118) {
		v202 = v151
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+5)))
	v194 = (v154 + v188 - int32(1)) & (int32(0) - v188)
	if int32(_a_F_create_toast_table_6) < v194 {
		v202 = v151
		goto L41
	} else {
		goto L54
	}
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v163))) = uint16(v194)
	v200 = v151 + int32(1)
	if v200 != v140 {
		v151 = v200
		v152 = v181
		v154 = v194 + v182
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	if v232 != 0 {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	goto L56
L58:
	;
	v225 = int32(1)
	if v220 == v224 {
		v232 = v225
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v232 = int32(0)
	goto L57
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[4]))
	if v228 == v220 {
		v232 = v225
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[4]))
	v235 = v234
	goto L65
L64:
	;
	v235 = int32(99)
	goto L65
L65:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+117)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+119)))
	switch v238 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L67
	default:
		v244 = int32(0)
		goto L66
	}
L66:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v236)+92))
	v248 = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v236)+80))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+156))
	v253 = m.T0[v252].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L20
	} else {
		goto L68
	}
L67:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)+88))
	v244 = base.B2i32(v241 == int32(0))
	goto L66
L68:
	;
	v255 = int32(0)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v258 = int32(*(*int8)(unsafe.Add(mBase, uint32(v257)+118)))
	v259 = int32(1)
	v266 = F_heap_create_with_catalog(m, v15+int32(224), v235, v247, l1, v248, v248, v250, v253, v71, v255, int32(116), v258, v237&v259, v244, v255, l3, v255, v259, v259, l6, v255)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L20
	} else {
		goto L69
	}
L69:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	v271 = F_table_open(m, v266, int32(5))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	v274 = F_palloc0(m, int32(144))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	v276 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v274)+76)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v274)+8)) = int64(562954248388610)
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = int64(8589934979)
	*(*int64)(unsafe.Add(mBase, uint32(v274)+84)) = v276
	*(*int64)(unsafe.Add(mBase, uint32(v274)+92)) = v276
	v286 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+100)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v274)+128)) = v286
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v274)+118)) = uint8(v290)
	*(*uint16)(unsafe.Add(mBase, uint32(v274)+116)) = uint16(v290)
	*(*int64)(unsafe.Add(mBase, uint32(v274)+132)) = int64(403)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+119)) = v286
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+140)) = v299
	*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(8495445313469)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+152)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v286
	v307 = int32(_a_F_create_toast_table_3)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+124)) = v307
	v309 = int32(_a_F_create_toast_table_4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v309
	v324 = F_list_make2_impl(m, v15+int32(28), v15+int32(24))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L20
	} else {
		goto L73
	}
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+92))
	v333 = int32(0)
	v340 = int32(1)
	v343 = F_index_create(m, v271, v15+int32(160), l2, v286, v286, v286, v274, v324, int32(403), v328, v15+int32(152), v15+int32(144), v333, v15+int32(140), v333, int64(0), int32(129), v333, v340, v340, v333)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L20
	} else {
		goto L74
	}
L74:
	;
	F_relation_close(m, v271, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	v350 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L20
	} else {
		goto L76
	}
L76:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[2]))
	if v353 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	F_pfree(m, v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L20
	} else {
		goto L88
	}
L78:
	;
	v357 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(v19), int64(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L20
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v371 = v15 - int32(-64)
	F_ScanKeyInit(m, v371, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v19))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L20
	} else {
		goto L84
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+300)) = v357
	if v357 == int32(0) {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v357)+16))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v362+v363)+112)) = v266
	F_CatalogTupleUpdate(m, v350, v357+int32(4), v357)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L20
	} else {
		goto L83
	}
L83:
	;
	v397 = v357
	goto L77
L84:
	;
	F_systable_inplace_update_begin(m, v350, int32(2662), v371, v15+int32(300), v15+int32(128))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
	if v385 == int32(0) {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v385)+16))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v388+v389)+112)) = v266
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
	F_systable_inplace_update_finish(m, v392, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
	v397 = v396
	goto L77
L88:
	;
	F_relation_close(m, v350, int32(3))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L20
	} else {
		goto L89
	}
L89:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_create_toast_table[2]))
	if v405 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v406 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v19
	v409 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = v409
	F_recordDependencyOn(m, v15+int32(128), v15-int32(-64), int32(105))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L20
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L20
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	v433 = int32(1)
	goto L4
L95:
	;
	F_errmsg_internal(m, int32(_a_F_create_toast_table_7), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L20
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_create_toast_table_8), int32(194), int32(_a_F_create_toast_table_9))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v19
	F_errmsg_internal(m, int32(_a_F_create_toast_table_10), v15+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L20
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_create_toast_table_8), int32(354), int32(_a_F_create_toast_table_9))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L20
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	F_errmsg_internal(m, int32(_a_F_create_toast_table_10), v15)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L20
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_create_toast_table_8), int32(374), int32(_a_F_create_toast_table_9))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_toast_fetch_datum_slice(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 != int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_toast_fetch_datum_slice_0), int32(0))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_toast_fetch_datum_slice_1), int32(405), int32(_a_F_toast_fetch_datum_slice_2))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v12 != int32(18) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_toast_fetch_datum_slice_0), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_toast_fetch_datum_slice_1), int32(405), int32(_a_F_toast_fetch_datum_slice_2))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+14))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+10))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+6))
			v19 = v17 & int32(1073741823)
			v21 = base.B2i32(base.Ui32(l1) < base.Ui32(v19))
			if base.Ui32(l1) < base.Ui32(v19) {
				v22 = l1
			} else {
				v22 = int32(0)
			}
			v23 = v19 - v22
			if base.Ui32(l1) < base.Ui32(v19) {
				v25 = l2
			} else {
				v25 = int32(0)
			}
			if int32(0) < v25 {
				v30 = v25 + int32(4)
			} else {
				v30 = v25
			}
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			v34 = base.B2i32(base.Ui32(v19) < base.Ui32(v31-int32(4)))
			if base.Ui32(v19) < base.Ui32(v31-int32(4)) {
				v35 = v30
			} else {
				v35 = v25
			}
			if v19 < v35+v22 {
				v38 = v23
			} else {
				v38 = v35
			}
			if v35 < int32(0) {
				v41 = v23
			} else {
				v41 = v38
			}
			v43 = v41 + int32(4)
			v44 = F_palloc(m, v43)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v48 = int32(2)
				v49 = v43 << (uint(v48) % 32)
				if base.Ui32(v19) < base.Ui32(v31-int32(4)) {
					v52 = v49 | v48
				} else {
					v52 = v49
				}
				*(*int32)(unsafe.Add(mBase, uint32(v44))) = v52
				if v41 != 0 {
					v55 = F_table_open(m, v15, int32(1))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+188))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+160))
						m.T0[v58].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v55, v16, v19, v22, v41, v44)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_relation_close(m, v55, int32(1))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								return v44
							}
						}
					}
				} else {
					return v44
				}
			}
		}
	}
}
func F_toast_tuple_try_compression(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v6 + l1<<(uint(int32(3))%32)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = v11 + l1*int32(12)
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+9)))
	v16 = F_toast_compress_datum(m, v10, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
		if base.I32_wrap_i64(v16) != 0 {
			if v18&int32(2) != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				F_pfree(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9))) = v16
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
					v27 = int32(2)
					v28 = v26 | v27
					*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v28)
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(base.Ui32(v31) >> (uint(v27) % 32))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
					v37 = v35 | int32(10)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v37)
					return
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v9))) = v16
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
				v27 = int32(2)
				v28 = v26 | v27
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(base.Ui32(v31) >> (uint(v27) % 32))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				v37 = v35 | int32(10)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v37)
				return
			}
		} else {
			v40 = v18 | int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v40)
			return
		}
	}
}
