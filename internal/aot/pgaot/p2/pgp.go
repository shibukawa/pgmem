package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgp_armor_headers(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_px_THROW_ERROR(m, v84)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L76
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L73
	}
L3:
	;
	v19 = int32(4470400)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	goto L30
L6:
	;
	return int32(0)
L7:
	;
	v25 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	v33 = F_get_call_result_type(m, l0, int32(0), v13+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v33 != int32(1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v38 = F_TupleDescGetAttInMetadata(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v38
	v42 = F_palloc(m, int32(12))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(1)
	v45 = v21 + v44
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v50 = v48 & v44
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v51 = v45
	goto L15
L14:
	;
	v51 = v21 + int32(4)
	goto L15
L15:
	;
	if v48 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v84 = F_pgp_extract_armor_headers(m, v51, v79, v42, v42+int32(4), v42+int32(8))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L27
	}
L17:
	;
	v54 = int32(4)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v56&int32(254) == int32(2) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v69 = int32(1)
	if v50 != 0 {
		v79 = int32(base.Ui32(v48)>>(uint(v69)%32)) - v69
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v65 = v54
	goto L22
L21:
	;
	v65 = base.B2i32(v56 == int32(18)) << (uint(v54) % 32)
	goto L22
L22:
	;
	if v56 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = v54
	goto L25
L24:
	;
	v68 = v65
	goto L25
L25:
	;
	v79 = v68
	goto L16
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	if v84 < int32(0) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v42
	goto L5
L29:
	;
	m.G0 = v13 + int32(16)
	return v261
L30:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	v102 = int64(*(*int32)(unsafe.Add(mBase, uint32(v101))))
	if base.Ui64(v102) <= base.Ui64(v100) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v114 = base.I32_wrap_i64(v100) << (uint(int32(2)) % 32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114+v115)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118+v114)))
	if v120&int32(3) == int32(0) {
		v144 = v120
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+20)) = int32(2)
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v109)
	v261 = int32(0)
	goto L29
L35:
	;
	v179 = F_pg_any_to_server(m, v120, v177, int32(6))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L52
	}
L36:
	;
	v177 = v169 - v120
	goto L35
L37:
	;
	v148 = v144
	goto L46
L38:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v128 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v177 = int32(0)
	goto L35
L40:
	;
	goto L41
L41:
	;
	v133 = v120
	goto L42
L42:
	;
	v137 = v133 + int32(1)
	if v137&int32(3) == int32(0) {
		v144 = v137
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v169 = v137
	goto L36
L44:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v142 != 0 {
		v133 = v137
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v157 = int32(-2139062144)
	if (int32(16843008)-v154|v154)&v157 == v157 {
		v148 = v148 + int32(4)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v163 = v148
	goto L49
L48:
	;
	goto L47
L49:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v167 != 0 {
		v163 = v163 + int32(1)
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v169 = v163
	goto L36
L51:
	;
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v179
	if v117&int32(3) == int32(0) {
		v205 = v117
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v240 = F_pg_any_to_server(m, v117, v238, int32(6))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L70
	}
L54:
	;
	v238 = v230 - v117
	goto L53
L55:
	;
	v209 = v205
	goto L64
L56:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v189 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v238 = int32(0)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v194 = v117
	goto L60
L60:
	;
	v198 = v194 + int32(1)
	if v198&int32(3) == int32(0) {
		v205 = v198
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v230 = v198
	goto L54
L62:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if v203 != 0 {
		v194 = v198
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v218 = int32(-2139062144)
	if (int32(16843008)-v215|v215)&v218 == v218 {
		v209 = v209 + int32(4)
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v224 = v209
	goto L67
L66:
	;
	goto L65
L67:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v228 != 0 {
		v224 = v224 + int32(1)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v230 = v224
	goto L54
L69:
	;
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v240
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v246 = F_BuildTupleFromCStrings(m, v243, v13+int32(4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = v248 + int64(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v256 = F_HeapTupleHeaderGetDatum(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v261 = v256
	goto L29
L73:
	;
	F_errmsg_internal(m, int32(360976), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(488361), int32(937), int32(133155))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_expect_packet_end(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_pullf_read(m, l0, int32(32768), v5+int32(12))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(0) < v10 {
			F_px_debug(m, int32(495251), int32(0))
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v21 = int32(-100)
				m.G0 = v5 + int32(16)
				return v21
			}
		} else {
			v21 = v10
			m.G0 = v5 + int32(16)
			return v21
		}
	}
}
func F_pgp_extract_armor_headers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(-101)
	v17 = l0 + l1
	if base.Ui32(v17) <= base.Ui32(l0) {
		v126 = v16
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L117
	} else {
		goto L151
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L117
	} else {
		goto L148
	}
L3:
	;
	m.G0 = v14 + int32(16)
	return v486
L4:
	;
	if v126 <= int32(0) {
		v486 = v16
		goto L3
	} else {
		goto L43
	}
L5:
	;
	goto L4
L6:
	;
	v29 = int32(10)
	goto L8
L8:
	;
	goto L9
L9:
	;
	if v17-l0 < v29 {
		v126 = v16
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v35 = int32(520977)
	goto L12
L12:
	;
	goto L13
L13:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, _consts[1390])))
	v42 = l0
	goto L14
L14:
	;
	v49 = F_memchr(m, v42, v38, v17-v42)
	mBase = m.M
	if v49 == int32(0) {
		v126 = v16
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(12)))) = v49
	if base.Ui32(v17) <= base.Ui32(v52) {
		v89 = v52
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v52 = v29 + v49
	if base.Ui32(v17) < base.Ui32(v52) {
		v126 = v16
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v54 = F_memcmp(m, v49, v35, v29)
	mBase = m.M
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v56 = v49 + int32(1)
	if base.Ui32(v56) < base.Ui32(v17) {
		v42 = v56
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l0 == v49 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v126 = v16
	goto L5
L22:
	;
	goto L15
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49-int32(1)))))
	if v61 == int32(10) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v17) <= base.Ui32(v52) {
		v126 = v16
		goto L5
	} else {
		goto L25
	}
L25:
	;
	if v29 <= v17-v52 {
		v42 = v52
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v126 = v16
	goto L5
L27:
	;
	if v17-v89 < int32(5) {
		v126 = v16
		goto L5
	} else {
		goto L34
	}
L28:
	;
	v72 = v52
	goto L29
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v78 == int32(45) {
		v89 = v72
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v89 = v17
	goto L27
L31:
	;
	if base.Ui32(v78) < base.Ui32(int32(32)) {
		v126 = v16
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v84 = v72 + int32(1)
	if base.Ui32(v84) < base.Ui32(v17) {
		v72 = v84
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v99 = F_memcmp(m, v89, v35, int32(5))
	mBase = m.M
	if v99 != 0 {
		v126 = v16
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v101 = v89 + int32(5)
	if base.Ui32(v17) <= base.Ui32(v101) {
		v116 = v101
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v126 = v116 - v49
	goto L5
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	switch v103 - int32(10) {
	case 0, 3:
		goto L38
	default:
		v126 = v16
		goto L5
	}
L38:
	;
	if v103 == int32(13) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v110 = v89 + int32(6)
	goto L41
L40:
	;
	v110 = v101
	goto L41
L41:
	;
	if base.Ui32(v17) <= base.Ui32(v110) {
		v116 = v110
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	v116 = v110 + base.B2i32(v112 == int32(10))
	goto L36
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v131 = v130 + v126
	v140 = int32(-101)
	if base.Ui32(v17) <= base.Ui32(v131) {
		v240 = v140
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v240 <= int32(0) {
		v486 = v16
		goto L3
	} else {
		goto L83
	}
L45:
	;
	goto L44
L46:
	;
	v142 = int32(8)
	goto L47
L47:
	;
	goto L49
L49:
	;
	if v17-v131 < v142 {
		v240 = v140
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v147 = int32(533191)
	goto L51
L51:
	;
	goto L53
L53:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, _consts[1391])))
	v156 = v131
	goto L54
L54:
	;
	v163 = F_memchr(m, v156, v152, v17-v156)
	mBase = m.M
	if v163 == int32(0) {
		v240 = v140
		goto L45
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+int32(8)))) = v163
	if base.Ui32(v17) <= base.Ui32(v166) {
		v203 = v166
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v166 = v142 + v163
	if base.Ui32(v17) < base.Ui32(v166) {
		v240 = v140
		goto L45
	} else {
		goto L57
	}
L57:
	;
	v168 = F_memcmp(m, v163, v147, v142)
	mBase = m.M
	if v168 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v170 = v163 + int32(1)
	if base.Ui32(v170) < base.Ui32(v17) {
		v156 = v170
		goto L54
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v131 == v163 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v240 = v140
	goto L45
L62:
	;
	goto L55
L63:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163-int32(1)))))
	if v175 == int32(10) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if base.Ui32(v17) <= base.Ui32(v166) {
		v240 = v140
		goto L45
	} else {
		goto L65
	}
L65:
	;
	if v142 <= v17-v166 {
		v156 = v166
		goto L54
	} else {
		goto L66
	}
L66:
	;
	v240 = v140
	goto L45
L67:
	;
	if v17-v203 < int32(5) {
		v240 = v140
		goto L45
	} else {
		goto L74
	}
L68:
	;
	v186 = v166
	goto L69
L69:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v192 == int32(45) {
		v203 = v186
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v203 = v17
	goto L67
L71:
	;
	if base.Ui32(v192) < base.Ui32(int32(32)) {
		v240 = v140
		goto L45
	} else {
		goto L72
	}
L72:
	;
	v198 = v186 + int32(1)
	if base.Ui32(v198) < base.Ui32(v17) {
		v186 = v198
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v213 = F_memcmp(m, v203, v147, int32(5))
	mBase = m.M
	if v213 != 0 {
		v240 = v140
		goto L45
	} else {
		goto L75
	}
L75:
	;
	v215 = v203 + int32(5)
	if base.Ui32(v17) <= base.Ui32(v215) {
		v230 = v215
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v240 = v230 - v163
	goto L45
L77:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	switch v217 - int32(10) {
	case 0, 3:
		goto L78
	default:
		v240 = v140
		goto L45
	}
L78:
	;
	if v217 == int32(13) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v224 = v203 + int32(6)
	goto L81
L80:
	;
	v224 = v215
	goto L81
L81:
	;
	if base.Ui32(v17) <= base.Ui32(v224) {
		v230 = v224
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v230 = v224 + base.B2i32(v226 == int32(10))
	goto L76
L83:
	;
	v244 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if base.Ui32(v245) <= base.Ui32(v131) {
		v373 = v131
		v379 = v6
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v384 = v373 - v131
	v387 = F_palloc(m, v384+int32(1))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L117
	} else {
		goto L118
	}
L85:
	;
	v247 = v131
	v253 = v6
	goto L86
L86:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	switch v258 - int32(10) {
	case 0, 3:
		v373 = v247
		v379 = v253
		goto L84
	default:
		goto L88
	}
L87:
	;
	v373 = v371
	v379 = v369
	goto L84
L88:
	;
	v262 = v245 - v247
	v263 = int32(0)
	v266 = base.B2i32(v262 != v263)
	if v247&int32(3) == v263 {
		v292 = v247
		v294 = v262
		v295 = v266
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v365 == int32(0) {
		v486 = v16
		goto L3
	} else {
		goto L115
	}
L90:
	;
	v365 = int32(0)
	goto L89
L91:
	;
	v343 = v336
	v345 = v338
	goto L109
L92:
	;
	if v295 == int32(0) {
		goto L90
	} else {
		goto L100
	}
L93:
	;
	if v262 == int32(0) {
		v292 = v247
		v294 = v262
		v295 = v266
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v275 = v247
	v277 = v262
	goto L95
L95:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v280 == int32(10) {
		v336 = v275
		v338 = v277
		goto L91
	} else {
		goto L97
	}
L96:
	;
	v292 = v287
	v294 = v283
	v295 = v285
	goto L92
L97:
	;
	v282 = int32(1)
	v283 = v277 - v282
	v284 = int32(0)
	v285 = base.B2i32(v283 != v284)
	v287 = v275 + v282
	if v287&int32(3) == v284 {
		v292 = v287
		v294 = v283
		v295 = v285
		goto L92
	} else {
		goto L98
	}
L98:
	;
	if v283 != 0 {
		v275 = v287
		v277 = v283
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v299 == int32(10) {
		v329 = v292
		v331 = v294
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v331 == int32(0) {
		goto L90
	} else {
		goto L108
	}
L102:
	;
	if base.Ui32(v294) < base.Ui32(int32(4)) {
		v329 = v292
		v331 = v294
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v309 = v292
	v311 = v294
	goto L104
L104:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v316 = v315 ^ int32(168430090)
	v319 = int32(-2139062144)
	if (int32(16843008)-v316|v316)&v319 != v319 {
		v336 = v309
		v338 = v311
		goto L91
	} else {
		goto L106
	}
L105:
	;
	v329 = v324
	v331 = v326
	goto L101
L106:
	;
	v323 = int32(4)
	v324 = v309 + v323
	v326 = v311 - v323
	if base.Ui32(int32(3)) < base.Ui32(v326) {
		v309 = v324
		v311 = v326
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v336 = v329
	v338 = v331
	goto L91
L109:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
	if int32(10) == v348 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L90
L111:
	;
	v365 = v343
	goto L89
L112:
	;
	goto L113
L113:
	;
	v350 = int32(1)
	v353 = v345 - v350
	if v353 != 0 {
		v343 = v343 + v350
		v345 = v353
		goto L109
	} else {
		goto L114
	}
L114:
	;
	goto L110
L115:
	;
	v368 = int32(1)
	v369 = v253 + v368
	v371 = v365 + v368
	if base.Ui32(v371) < base.Ui32(v245) {
		v247 = v371
		v253 = v369
		goto L86
	} else {
		goto L116
	}
L116:
	;
	goto L87
L117:
	;
	return int32(0)
L118:
	;
	if v384 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v384))) = uint8(v394)
	v397 = v379 << (uint(int32(2)) % 32)
	v398 = F_palloc(m, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L117
	} else {
		goto L123
	}
L120:
	;
	v391 = F__emscripten_memcpy_bulkmem(m, v387, v131, v384)
	mBase = m.M
	v392 = v391
	goto L122
L121:
	;
	v392 = v387
	goto L122
L122:
	;
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v398
	v401 = F_palloc(m, v397)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L117
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v401
	v404 = int32(10)
	v405 = F___strchrnul(m, v392, v404)
	mBase = m.M
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	if v407 == v404 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v411 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v411 = v405
	goto L128
L127:
	;
	v411 = int32(0)
	goto L128
L128:
	;
	goto L125
L129:
	;
	v412 = v411
	v413 = v244
	v419 = v387
	goto L132
L130:
	;
	v464 = v244
	goto L131
L131:
	;
	if v464 != v379 {
		goto L1
	} else {
		goto L147
	}
L132:
	;
	if base.Ui32(v419) < base.Ui32(v412) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v464 = v452
	goto L131
L134:
	;
	v425 = v412 - int32(1)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425))))
	if v426 == int32(13) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v431 = v412
	goto L136
L136:
	;
	v432 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v431))) = uint8(v432)
	v435 = F_strstr(m, v419, int32(719688))
	mBase = m.M
	if v435 == v432 {
		v486 = v16
		goto L3
	} else {
		goto L140
	}
L137:
	;
	v429 = v425
	goto L139
L138:
	;
	v429 = v412
	goto L139
L139:
	;
	v431 = v429
	goto L136
L140:
	;
	v438 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v435))) = uint8(v438)
	if v413 == v379 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v441 = int32(2)
	v442 = v413 << (uint(v441) % 32)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v442+v443))) = v419
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v446+v442))) = v435 + v441
	v451 = int32(1)
	v452 = v413 + v451
	v454 = v412 + v451
	v455 = int32(10)
	v456 = F___strchrnul(m, v454, v455)
	mBase = m.M
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v458 == v455 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v462 != 0 {
		v412 = v462
		v413 = v452
		v419 = v454
		goto L132
	} else {
		goto L146
	}
L143:
	;
	v462 = v456
	goto L145
L144:
	;
	v462 = int32(0)
	goto L145
L145:
	;
	goto L142
L146:
	;
	goto L133
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v379
	v486 = int32(0)
	goto L3
L148:
	;
	F_errmsg_internal(m, int32(160180), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L117
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(486210), int32(473), int32(133120))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L117
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errmsg_internal(m, int32(160180), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L117
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(486210), int32(484), int32(133120))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L117
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgp_get_cipher_code(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	v9 = int32(168523)
	v10 = l0
	goto L4
L1:
	;
	return v438
L2:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v438 = v437
	goto L1
L3:
	;
	if v47 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v13 == v14 {
		v36 = v13
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v47 = int32(0)
	goto L3
L6:
	;
	v38 = int32(1)
	if v36 != 0 {
		v9 = v9 + v38
		v10 = v10 + v38
		goto L4
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v13-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = v13 | int32(32)
	goto L10
L9:
	;
	v24 = v13
	goto L10
L10:
	;
	if base.Ui32((v14-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v14 | int32(32)
	goto L13
L12:
	;
	v33 = v14
	goto L13
L13:
	;
	if v24 == v33 {
		v36 = v24
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v47 = v24 - v33
	goto L3
L15:
	;
	goto L5
L16:
	;
	v436 = int32(4361936)
	goto L2
L17:
	;
	goto L18
L18:
	;
	v57 = int32(540721)
	v58 = l0
	goto L20
L19:
	;
	if v95 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v61 == v62 {
		v84 = v61
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v95 = int32(0)
	goto L19
L22:
	;
	v86 = int32(1)
	if v84 != 0 {
		v57 = v57 + v86
		v58 = v58 + v86
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v61-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v72 = v61 | int32(32)
	goto L26
L25:
	;
	v72 = v61
	goto L26
L26:
	;
	if base.Ui32((v62-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = v62 | int32(32)
	goto L29
L28:
	;
	v81 = v62
	goto L29
L29:
	;
	if v72 == v81 {
		v84 = v72
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v95 = v72 - v81
	goto L19
L31:
	;
	goto L21
L32:
	;
	v436 = int32(4361956)
	goto L2
L33:
	;
	goto L34
L34:
	;
	v105 = int32(333887)
	v106 = l0
	goto L36
L35:
	;
	if v143 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == v110 {
		v132 = v109
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v143 = int32(0)
	goto L35
L38:
	;
	v134 = int32(1)
	if v132 != 0 {
		v105 = v105 + v134
		v106 = v106 + v134
		goto L36
	} else {
		goto L47
	}
L39:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = v109 | int32(32)
	goto L42
L41:
	;
	v120 = v109
	goto L42
L42:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v129 = v110 | int32(32)
	goto L45
L44:
	;
	v129 = v110
	goto L45
L45:
	;
	if v120 == v129 {
		v132 = v120
		goto L38
	} else {
		goto L46
	}
L46:
	;
	v143 = v120 - v129
	goto L35
L47:
	;
	goto L37
L48:
	;
	v436 = int32(4361976)
	goto L2
L49:
	;
	goto L50
L50:
	;
	v153 = int32(317106)
	v154 = l0
	goto L52
L51:
	;
	if v191 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == v158 {
		v180 = v157
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v191 = int32(0)
	goto L51
L54:
	;
	v182 = int32(1)
	if v180 != 0 {
		v153 = v153 + v182
		v154 = v154 + v182
		goto L52
	} else {
		goto L63
	}
L55:
	;
	if base.Ui32((v157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v168 = v157 | int32(32)
	goto L58
L57:
	;
	v168 = v157
	goto L58
L58:
	;
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v177 = v158 | int32(32)
	goto L61
L60:
	;
	v177 = v158
	goto L61
L61:
	;
	if v168 == v177 {
		v180 = v168
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v191 = v168 - v177
	goto L51
L63:
	;
	goto L53
L64:
	;
	v436 = int32(4361996)
	goto L2
L65:
	;
	goto L66
L66:
	;
	v201 = int32(169023)
	v202 = l0
	goto L68
L67:
	;
	if v239 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v205 == v206 {
		v228 = v205
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v239 = int32(0)
	goto L67
L70:
	;
	v230 = int32(1)
	if v228 != 0 {
		v201 = v201 + v230
		v202 = v202 + v230
		goto L68
	} else {
		goto L79
	}
L71:
	;
	if base.Ui32((v205-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v216 = v205 | int32(32)
	goto L74
L73:
	;
	v216 = v205
	goto L74
L74:
	;
	if base.Ui32((v206-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v225 = v206 | int32(32)
	goto L77
L76:
	;
	v225 = v206
	goto L77
L77:
	;
	if v216 == v225 {
		v228 = v216
		goto L70
	} else {
		goto L78
	}
L78:
	;
	v239 = v216 - v225
	goto L67
L79:
	;
	goto L69
L80:
	;
	v436 = int32(4362016)
	goto L2
L81:
	;
	goto L82
L82:
	;
	v249 = int32(540105)
	v250 = l0
	goto L84
L83:
	;
	if v287 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L84:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v253 == v254 {
		v276 = v253
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v287 = int32(0)
	goto L83
L86:
	;
	v278 = int32(1)
	if v276 != 0 {
		v249 = v249 + v278
		v250 = v250 + v278
		goto L84
	} else {
		goto L95
	}
L87:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v264 = v253 | int32(32)
	goto L90
L89:
	;
	v264 = v253
	goto L90
L90:
	;
	if base.Ui32((v254-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v273 = v254 | int32(32)
	goto L93
L92:
	;
	v273 = v254
	goto L93
L93:
	;
	if v264 == v273 {
		v276 = v264
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v287 = v264 - v273
	goto L83
L95:
	;
	goto L85
L96:
	;
	v436 = int32(4362036)
	goto L2
L97:
	;
	goto L98
L98:
	;
	v297 = int32(543663)
	v298 = l0
	goto L100
L99:
	;
	if v335 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L100:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if v301 == v302 {
		v324 = v301
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v335 = int32(0)
	goto L99
L102:
	;
	v326 = int32(1)
	if v324 != 0 {
		v297 = v297 + v326
		v298 = v298 + v326
		goto L100
	} else {
		goto L111
	}
L103:
	;
	if base.Ui32((v301-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v312 = v301 | int32(32)
	goto L106
L105:
	;
	v312 = v301
	goto L106
L106:
	;
	if base.Ui32((v302-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v321 = v302 | int32(32)
	goto L109
L108:
	;
	v321 = v302
	goto L109
L109:
	;
	if v312 == v321 {
		v324 = v312
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v335 = v312 - v321
	goto L99
L111:
	;
	goto L101
L112:
	;
	v436 = int32(4362056)
	goto L2
L113:
	;
	goto L114
L114:
	;
	v345 = int32(540485)
	v346 = l0
	goto L116
L115:
	;
	if v383 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L116:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if v349 == v350 {
		v372 = v349
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v383 = int32(0)
	goto L115
L118:
	;
	v374 = int32(1)
	if v372 != 0 {
		v345 = v345 + v374
		v346 = v346 + v374
		goto L116
	} else {
		goto L127
	}
L119:
	;
	if base.Ui32((v349-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v360 = v349 | int32(32)
	goto L122
L121:
	;
	v360 = v349
	goto L122
L122:
	;
	if base.Ui32((v350-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v369 = v350 | int32(32)
	goto L125
L124:
	;
	v369 = v350
	goto L125
L125:
	;
	if v360 == v369 {
		v372 = v360
		goto L118
	} else {
		goto L126
	}
L126:
	;
	v383 = v360 - v369
	goto L115
L127:
	;
	goto L117
L128:
	;
	v436 = int32(4362076)
	goto L2
L129:
	;
	goto L130
L130:
	;
	v394 = int32(317115)
	v395 = l0
	goto L132
L131:
	;
	if v432 != 0 {
		v438 = int32(-103)
		goto L1
	} else {
		goto L144
	}
L132:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v398 == v399 {
		v421 = v398
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v432 = int32(0)
	goto L131
L134:
	;
	v423 = int32(1)
	if v421 != 0 {
		v394 = v394 + v423
		v395 = v395 + v423
		goto L132
	} else {
		goto L143
	}
L135:
	;
	if base.Ui32((v398-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v409 = v398 | int32(32)
	goto L138
L137:
	;
	v409 = v398
	goto L138
L138:
	;
	if base.Ui32((v399-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v418 = v399 | int32(32)
	goto L141
L140:
	;
	v418 = v399
	goto L141
L141:
	;
	if v409 == v418 {
		v421 = v409
		goto L134
	} else {
		goto L142
	}
L142:
	;
	v432 = v409 - v418
	goto L131
L143:
	;
	goto L133
L144:
	;
	v436 = int32(4362096)
	goto L2
}
func F_pgp_get_cipher_key_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v2 = int32(0)
	v4 = l0 - int32(2)
	if base.Ui32(int32(8)) < base.Ui32(v4) {
		v19 = v2
	} else {
		if int32(base.Ui32(int32(487))>>(uint(v4)%32))&int32(1) == int32(0) {
			v19 = v2
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v4<<(uint(int32(2))%32))+uint32(_consts[1392])))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v19 = v18
		}
	}
	return v19
}
func F_pgp_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v20 int32
	_ = v20
	v4 = F_palloc0(m, int32(168))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v4)+60)) = int64(7)
		v10 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+72)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v4)+68)) = int32(6)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+52)) = int64(-4294967294)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+44)) = int64(-4294967293)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+80)) = v10
		v20 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+88)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v4
		return v20
	}
}
func F_pgp_rsa_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v12 = int32(-109)
	v13 = F_mpi_check(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v58 = v12
			return v58
		} else {
			v19 = F_mpi_check(m, v9)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v58 = v12
					return v58
				} else {
					v23 = F_mpi_check(m, v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						if v23 == int32(0) {
							v58 = v12
							return v58
						} else {
							v27 = int32(1)
							if v11 <= v27 {
								v30 = v27
							} else {
								v30 = v11
							}
							v31 = F_palloc(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v40 = m.Env.Pgmem_bn_op(m, int32(1), v34, v35, v36, v37, v38, v39, v31, v11)
								mBase = m.M
								if v40 < int32(0) {
									v51 = int32(-109)
									if v31 == int32(0) {
										v58 = v51
										return v58
									} else {
										v55 = F___memset(m, v31, int32(0), v30)
										mBase = m.M
										F_pfree(m, v31)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = v51
											return v58
										}
									}
								} else {
									v44 = F_bytes_to_mpi(m, v31, v40)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v44
										if v44 != 0 {
											v49 = int32(0)
										} else {
											v49 = int32(-109)
										}
										v51 = v49
										if v31 == int32(0) {
											v58 = v51
											return v58
										} else {
											v55 = F___memset(m, v31, int32(0), v30)
											mBase = m.M
											F_pfree(m, v31)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = v51
												return v58
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
func F_pgp_set_unicode_mode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = base.B2i32(l1 != v3)
	return v3
}
func F_pgp_skip_packet(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	goto L1
L1:
	;
	v14 = F_pullf_read(m, l0, int32(32768), v6+int32(12))
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v6 + int32(16)
	return v14
L3:
	;
	return int32(0)
L4:
	;
	if int32(0) < v14 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
}
