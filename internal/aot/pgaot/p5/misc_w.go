package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WorkTableScanNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+108))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v7 = F_tuplestore_gettupleslot(m, v3, int32(1), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F___wasm_call_ctors(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v12 = m.Wasi_snapshot_preview1.Environ_sizes_get(m, v6+int32(12), v6+int32(8))
	mBase = m.M
	if v12 != 0 {
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v19 = F_emscripten_builtin_malloc(m, v14<<(uint(int32(2))%32)+int32(4))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v19
		if v19 == int32(0) {
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v24 = F_emscripten_builtin_malloc(m, v23)
			mBase = m.M
			if v24 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v31 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32)))) = v31
				v33 = m.Wasi_snapshot_preview1.Environ_get(m, v26, v24)
				mBase = m.M
				if v33 == v31 {
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = int32(0)
			}
		}
	}
	m.G0 = v6 + int32(16)
	*(*int32)(unsafe.Add(mBase, _consts[1])) = int32(4680248)
	*(*int32)(unsafe.Add(mBase, _consts[2])) = int32(42)
	return
}
func F_websearch_to_tsquery_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		v17 = F_text_to_cstring(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v24 = F_parse_tsquery(m, v17, int32(1174), v6+int32(8), int32(2), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func F_weight_checkdig(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = l0
	v10 = l1
	v11 = v6
	v12 = v3
	goto L3
L3:
	;
	v17 = (v11 - int32(48)) & int32(255)
	v21 = base.B2i32(base.Ui32(v17) < base.Ui32(int32(10)))
	if base.Ui32(v17) < base.Ui32(int32(10)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v35 = base.I32_rem_u_s(v23, int32(11))
	if v35 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L5:
	;
	goto L4
L6:
	;
	v22 = v10 * v17
	goto L8
L7:
	;
	v22 = int32(0)
	goto L8
L8:
	;
	v23 = v22 + v12
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v24 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v27 = int32(1)
	v29 = v10 - v21
	if base.Ui32(v27) < base.Ui32(v29) {
		v9 = v9 + v27
		v10 = v29
		v11 = v24
		v12 = v23
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	return int32(11) - v35
}
func F_width_bucket_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v153 float64
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v201 float64
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L3
	} else {
		goto L167
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L163
	}
L3:
	;
	return int32(0)
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v29 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v34 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L3
	} else {
		goto L159
	}
L8:
	;
	if v32 == int32(701) {
		goto L29
	} else {
		goto L30
	}
L9:
	;
	v141 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v38 = v25 + int32(16)
	v39 = F_ArrayGetNItems(m, v29, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v47 = v38 + v42<<(uint(int32(3))%32)
	goto L15
L14:
	;
	v47 = int32(0)
	goto L15
L15:
	;
	if int32(8) <= v39 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = v39
	v59 = v47
	goto L19
L17:
	;
	v85 = v39
	v86 = v47
	goto L18
L18:
	;
	if v85 <= int32(0) {
		v141 = v41
		goto L8
	} else {
		goto L23
	}
L19:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v75 != int32(255) {
		goto L2
	} else {
		goto L21
	}
L20:
	;
	v85 = v83
	v86 = v47 + int32(base.Ui32(v39-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L18
L21:
	;
	v83 = v58 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v58) {
		v58 = v83
		v59 = v59 + int32(1)
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v107 = v85
	v108 = int32(1)
	goto L24
L24:
	;
	if v108&v104 == int32(0) {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v141 = v41
	goto L8
L26:
	;
	v127 = int32(1)
	if v127 < v107 {
		v107 = v107 - v127
		v108 = v108 << (uint(v127) % 32)
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v580 != v25 {
		goto L155
	} else {
		goto L156
	}
L29:
	;
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v157 = F_ArrayGetNItems(m, v154, v25+int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	if v213 != 0 {
		goto L50
	} else {
		goto L51
	}
L32:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v153)&int64(9223372036854775807)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v568 = v157
	goto L28
L34:
	;
	goto L35
L35:
	;
	v164 = int32(0)
	if v157 <= v164 {
		v568 = v164
		goto L28
	} else {
		goto L36
	}
L36:
	;
	if v141 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v173 = v141
	goto L39
L38:
	;
	v173 = (v154<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L39
L39:
	;
	v176 = v157
	v181 = v164
	goto L40
L40:
	;
	v195 = base.I32_div_s(v176+v181, int32(2))
	v201 = *(*float64)(unsafe.Add(mBase, uint32(v25+v173+v195<<(uint(int32(3))%32))))
	v208 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v201)&int64(9223372036854775807))) | base.F64_lt(v153, v201)
	if v208 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v568 = v209
	goto L28
L42:
	;
	v209 = v181
	goto L44
L43:
	;
	v209 = v195 + int32(1)
	goto L44
L44:
	;
	if v208 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v210 = v195
	goto L47
L46:
	;
	v210 = v176
	goto L47
L47:
	;
	if v209 < v210 {
		v176 = v210
		v181 = v209
		goto L40
	} else {
		goto L48
	}
L48:
	;
	goto L41
L49:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+8)))
	v226 = base.I32_extend16_s(v225)
	if int32(0) < v226 {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v214 == v32 {
		v224 = v213
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v217 = F_lookup_type_cache(m, v32, int32(64))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+108))
	if v219 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v217
	v224 = v217
	goto L49
L56:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+10)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v232 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+62)) = uint16(v232)
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+60)) = uint8(v234)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v224 + int32(104)
	v245 = F_ArrayGetNItems(m, v231, v25+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+11)))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+10)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v328 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+62)) = uint16(v328)
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+60)) = uint8(v330)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v224 + int32(104)
	v341 = F_ArrayGetNItems(m, v327, v25+int32(16))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L3
	} else {
		goto L83
	}
L59:
	;
	if v245 <= int32(0) {
		v568 = v234
		goto L28
	} else {
		goto L60
	}
L60:
	;
	if v230 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v255 = v230
	goto L63
L62:
	;
	v255 = (v231<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L63
L63:
	;
	v257 = int32(1)
	v262 = v245
	v267 = v234
	goto L64
L64:
	;
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+68)) = uint8(v279)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v23
	v284 = base.I32_div_s(v262+v267, int32(2))
	v286 = v25 + v255 + v284*v225
	if v229&v257 == v279 {
		v292 = v286
		goto L67
	} else {
		goto L68
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L3
	} else {
		goto L80
	}
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v292
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+76)) = uint8(v294)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v302 = m.T0[v301].(func(*base.Module, int32) int32)(m, v21+int32(44))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L3
	} else {
		goto L72
	}
L68:
	;
	switch v225 - v257 {
	case 0:
		goto L71
	case 1:
		goto L70
	default:
		goto L66
	case 3:
		goto L69
	}
L69:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v292 = v291
	goto L67
L70:
	;
	v290 = int32(*(*int16)(unsafe.Add(mBase, uint32(v286))))
	v292 = v290
	goto L67
L71:
	;
	v289 = int32(*(*int8)(unsafe.Add(mBase, uint32(v286))))
	v292 = v289
	goto L67
L72:
	;
	v305 = base.B2i32(v302 < int32(0))
	if v302 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v306 = v267
	goto L75
L74:
	;
	v306 = v284 + int32(1)
	goto L75
L75:
	;
	if v302 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v307 = v284
	goto L78
L77:
	;
	v307 = v262
	goto L78
L78:
	;
	if v306 < v307 {
		v262 = v307
		v267 = v306
		goto L64
	} else {
		goto L79
	}
L79:
	;
	v568 = v306
	goto L28
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v225
	F_errmsg_internal(m, int32(483562), v21+int32(16))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(326784), int32(70), int32(67821))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L3
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
	if v341 <= int32(0) {
		v568 = v330
		goto L28
	} else {
		goto L84
	}
L84:
	;
	if v324 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v351 = v324
	goto L87
L86:
	;
	v351 = (v327<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L87
L87:
	;
	v354 = base.B2i32(v226 != int32(-1))
	v356 = v325 - int32(99)
	v365 = v330
	v367 = v25 + v351
	v373 = v341
	goto L88
L88:
	;
	v379 = base.I32_div_s(v365+v373, int32(2))
	if v365 < v379 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v568 = v558
	goto L28
L90:
	;
	v382 = v367
	v383 = v365
	goto L93
L91:
	;
	v453 = v367
	goto L92
L92:
	;
	v470 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+68)) = uint8(v470)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v23
	if v326&int32(1) != 0 {
		goto L118
	} else {
		goto L119
	}
L93:
	;
	if v354 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v453 = v448
	goto L92
L95:
	;
	switch v356 {
	case 0:
		v448 = v435
		goto L112
	case 1:
		goto L114
	default:
		goto L113
	case 6:
		goto L115
	}
L96:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v401 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v430 = F_strlen(m, v382)
	mBase = m.M
	v435 = v430 + v382 + int32(1)
	goto L95
L99:
	;
	v404 = int32(6)
	v406 = int32(18)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+1)))
	if v408 == v406 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v421 = int32(1)
	if v401&v421 != 0 {
		v435 = v382 + int32(base.Ui32(v401)>>(uint(v421)%32))
		goto L95
	} else {
		goto L111
	}
L102:
	;
	v411 = v406
	goto L104
L103:
	;
	v411 = int32(2)
	goto L104
L104:
	;
	if v408&int32(254) == int32(2) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v416 = v404
	goto L107
L106:
	;
	v416 = v411
	goto L107
L107:
	;
	if v408 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v419 = v404
	goto L110
L109:
	;
	v419 = v416
	goto L110
L110:
	;
	v435 = v382 + v419
	goto L95
L111:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v435 = v382 + int32(base.Ui32(v426)>>(uint(int32(2))%32))
	goto L95
L112:
	;
	v450 = v383 + int32(1)
	if v450 != v379 {
		v382 = v448
		v383 = v450
		goto L93
	} else {
		goto L116
	}
L113:
	;
	v448 = (v435 + int32(1)) & int32(-2)
	goto L112
L114:
	;
	v448 = (v435 + int32(7)) & int32(-8)
	goto L112
L115:
	;
	v448 = (v435 + int32(3)) & int32(-4)
	goto L112
L116:
	;
	goto L94
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v493
	v495 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+76)) = uint8(v495)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v501 = m.T0[v500].(func(*base.Module, int32) int32)(m, v21+int32(44))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L3
	} else {
		goto L129
	}
L118:
	;
	switch v225 - int32(1) {
	case 0:
		goto L124
	case 1:
		goto L123
	default:
		goto L121
	case 3:
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v493 = v453
	goto L117
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L3
	} else {
		goto L125
	}
L122:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v493 = v477
	goto L117
L123:
	;
	v476 = int32(*(*int16)(unsafe.Add(mBase, uint32(v453))))
	v493 = v476
	goto L117
L124:
	;
	v475 = int32(*(*int8)(unsafe.Add(mBase, uint32(v453))))
	v493 = v475
	goto L117
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v226
	F_errmsg_internal(m, int32(483562), v21+int32(32))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(326784), int32(70), int32(67821))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	if v558 < v560 {
		v365 = v558
		v367 = v559
		v373 = v560
		goto L88
	} else {
		goto L154
	}
L129:
	;
	if v501 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v558 = v365
	v559 = v367
	v560 = v379
	goto L128
L131:
	;
	goto L132
L132:
	;
	if v354 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	switch v356 {
	case 0:
		v554 = v541
		goto L150
	case 1:
		goto L152
	default:
		goto L151
	case 6:
		goto L153
	}
L134:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	if v507 == int32(1) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	v536 = F_strlen(m, v453)
	mBase = m.M
	v541 = v536 + v453 + int32(1)
	goto L133
L137:
	;
	v510 = int32(6)
	v512 = int32(18)
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	if v514 == v512 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v527 = int32(1)
	if v507&v527 != 0 {
		v541 = v453 + int32(base.Ui32(v507)>>(uint(v527)%32))
		goto L133
	} else {
		goto L149
	}
L140:
	;
	v517 = v512
	goto L142
L141:
	;
	v517 = int32(2)
	goto L142
L142:
	;
	if v514&int32(254) == int32(2) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v522 = v510
	goto L145
L144:
	;
	v522 = v517
	goto L145
L145:
	;
	if v514 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v525 = v510
	goto L148
L147:
	;
	v525 = v522
	goto L148
L148:
	;
	v541 = v453 + v525
	goto L133
L149:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v541 = v453 + int32(base.Ui32(v532)>>(uint(int32(2))%32))
	goto L133
L150:
	;
	v558 = v379 + int32(1)
	v559 = v554
	v560 = v373
	goto L128
L151:
	;
	v554 = (v541 + int32(1)) & int32(-2)
	goto L150
L152:
	;
	v554 = (v541 + int32(7)) & int32(-8)
	goto L150
L153:
	;
	v554 = (v541 + int32(3)) & int32(-4)
	goto L150
L154:
	;
	goto L89
L155:
	;
	F_pfree(m, v25)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L3
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	m.G0 = v21 + int32(80)
	return v568
L158:
	;
	goto L157
L159:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(25719), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(494818), int32(6708), int32(24188))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	F_errmsg(m, int32(174946), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(494818), int32(6713), int32(24188))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L3
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	v645 = F_format_type_be(m, v32)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v645
	F_errmsg(m, int32(189645), v21)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L3
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(494818), int32(6733), int32(24188))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_write_pipe_chunks(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v9 = m.G0
	v11 = v9 - int32(4096)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[466]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	if v15 < int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
		v20 = v18
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
		v20 = v19
	}
	if v20 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(8)
		v27 = int32(-1)
	} else {
		v27 = v20
	}
	v28 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11))) = uint16(v28)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v28)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v33
	v37 = int32(1)
	switch l2 - v37 {
	case 0:
		v44 = int32(16)
		v45 = int32(17)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v44)
		v49 = v45
	default:
		v49 = v37
	case 7:
		v44 = int32(32)
		v45 = int32(33)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v44)
		v49 = v45
	case 15:
		v44 = int32(64)
		v45 = int32(65)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v44)
		v49 = v45
	}
	if l1 < int32(4088) {
		v75 = l0
		v79 = l1
	} else {
		v54 = l0
		v55 = l1
		for {
			v62 = int32(4087)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)) = uint16(v62)
			v65 = F__emscripten_memcpy_bulkmem(m, v11+int32(9), v54, v62)
			mBase = m.M
			v68 = F_write(m, v27, v11, int32(4096))
			mBase = m.M
			v69 = int32(4087)
			v70 = v54 + v69
			v74 = v55 - v69
			if base.Ui32(int32(8174)) < base.Ui32(v55) {
				v54 = v70
				v55 = v74
				continue
			} else {
				break
			}
			break
		}
		v75 = v70
		v79 = v74
	}
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)) = uint16(v79)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v49)
	if v79 != 0 {
		v87 = F__emscripten_memcpy_bulkmem(m, v11+int32(9), v75, v79)
		mBase = m.M
	} else {
	}
	v91 = F_write(m, v27, v11, v79+int32(9))
	mBase = m.M
	m.G0 = v11 + int32(4096)
	return
}
func F_writetup_heap_1(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11 - int32(6)
	F_LogicalTapeWrite(m, l1, v8+int32(12), int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = int32(10)
		F_LogicalTapeWrite(m, l1, v10+v20, v11-v20)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v26&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v8+int32(12), int32(4))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
