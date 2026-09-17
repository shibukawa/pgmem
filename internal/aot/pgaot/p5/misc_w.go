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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v60 int32
	_ = v60
	m.G2 = int32(13128304)
	m.G1 = int32(_a_F___wasm_call_ctors_0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v16 = m.Wasi_snapshot_preview1.Environ_sizes_get(m, v10+int32(12), v10+int32(8))
	mBase = m.M
	if v16 != 0 {
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v23 = F_emscripten_builtin_malloc(m, v18<<(uint(int32(2))%32)+int32(4))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[0])) = v23
		if v23 == int32(0) {
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v28 = F_emscripten_builtin_malloc(m, v27)
			mBase = m.M
			if v28 != 0 {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[0]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = v35
				v37 = m.Wasi_snapshot_preview1.Environ_get(m, v30, v28)
				mBase = m.M
				if v37 == v35 {
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[0])) = int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[0])) = int32(0)
			}
		}
	}
	m.G0 = v10 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[1])) = int32(13128304)
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[2])) = int32(42)
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[3])) = int32(_a_F___wasm_call_ctors_1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[4]))
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[5])) = v60
	return
}
func F_websearch_to_tsquery_byid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14006(m, l0, int32(2), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v195 float64
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
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
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L155
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L150
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L146
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v28 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v33 == int32(0) {
		v136 = int32(0)
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L142
	}
L9:
	;
	if v31 == int32(701) {
		goto L28
	} else {
		goto L29
	}
L10:
	;
	v37 = v24 + int32(16)
	v38 = F_ArrayGetNItemsSafe(m, v28, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v46 = v37 + v41<<(uint(int32(3))%32)
	goto L14
L13:
	;
	v46 = int32(0)
	goto L14
L14:
	;
	if int32(8) <= v38 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v57 = v38
	v59 = v46
	goto L18
L16:
	;
	v83 = v38
	v85 = v46
	goto L17
L17:
	;
	if v83 <= int32(0) {
		v136 = v40
		goto L9
	} else {
		goto L22
	}
L18:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v73 != int32(255) {
		goto L3
	} else {
		goto L20
	}
L19:
	;
	v83 = v81
	v85 = v46 + int32(base.Ui32(v38-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L17
L20:
	;
	v81 = v57 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v57) {
		v57 = v81
		v59 = v59 + int32(1)
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v104 = v83
	v106 = int32(1)
	goto L23
L23:
	;
	if v106&v101 == int32(0) {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	v136 = v40
	goto L9
L25:
	;
	v123 = int32(1)
	if v123 < v104 {
		v104 = v104 - v123
		v106 = v106 << (uint(v123) % 32)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v543 != v24 {
		goto L138
	} else {
		goto L139
	}
L28:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v22)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v152 = F_ArrayGetNItemsSafe(m, v149, v24+int32(16))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	if v207 != 0 {
		goto L49
	} else {
		goto L50
	}
L31:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v148)&int64(9223372036854775807)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v530 = v152
	goto L27
L33:
	;
	goto L34
L34:
	;
	v159 = int32(0)
	if v152 <= v159 {
		v530 = v159
		goto L27
	} else {
		goto L35
	}
L35:
	;
	if v136 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v168 = v136
	goto L38
L37:
	;
	v168 = (v149<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L38
L38:
	;
	v171 = v152
	v174 = v159
	goto L39
L39:
	;
	v189 = base.I32_div_s(v171+v174, int32(2))
	v195 = *(*float64)(unsafe.Add(mBase, uint32(v24+v168+v189<<(uint(int32(3))%32))))
	v202 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v195)&int64(9223372036854775807))) | base.F64_lt(v148, v195)
	if v202 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v530 = v203
	goto L27
L41:
	;
	v203 = v174
	goto L43
L42:
	;
	v203 = v189 + int32(1)
	goto L43
L43:
	;
	if v202 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v204 = v189
	goto L46
L45:
	;
	v204 = v171
	goto L46
L46:
	;
	if v203 < v204 {
		v171 = v204
		v174 = v203
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+8)))
	v220 = base.I32_extend16_s(v219)
	if int32(0) < v220 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
	if v208 == v31 {
		v218 = v207
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v211 = F_lookup_type_cache(m, v31, int32(64))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+108))
	if v213 == int32(0) {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v211
	v218 = v211
	goto L48
L55:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+10)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v226 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+62)) = uint16(v226)
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+60)) = uint8(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v218 + int32(104)
	v239 = F_ArrayGetNItemsSafe(m, v225, v24+int32(16))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+11)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+10)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v321 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+62)) = uint16(v321)
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+60)) = uint8(v323)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v218 + int32(104)
	v334 = F_ArrayGetNItemsSafe(m, v320, v24+int32(16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L82
	}
L58:
	;
	if v239 <= int32(0) {
		v530 = v228
		goto L27
	} else {
		goto L59
	}
L59:
	;
	if v224 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v249 = v224
	goto L62
L61:
	;
	v249 = (v225<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L62
L62:
	;
	v251 = int32(1)
	v256 = v239
	v259 = v228
	goto L63
L63:
	;
	v272 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+68)) = uint8(v272)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v22
	v277 = base.I32_div_s(v256+v259, int32(2))
	v279 = v24 + v249 + v277*v219
	if v223&v251 == v272 {
		v285 = v279
		goto L66
	} else {
		goto L67
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L79
	}
L65:
	;
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v285
	v287 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+76)) = uint8(v287)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v295 = m.T0[v294].(func(*base.Module, int32) int32)(m, v20+int32(44))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L71
	}
L67:
	;
	switch v219 - v251 {
	case 0:
		goto L70
	case 1:
		goto L69
	default:
		goto L65
	case 3:
		goto L68
	}
L68:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v285 = v284
	goto L66
L69:
	;
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v279))))
	v285 = v283
	goto L66
L70:
	;
	v282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v279))))
	v285 = v282
	goto L66
L71:
	;
	v298 = base.B2i32(v295 < int32(0))
	if v295 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v299 = v259
	goto L74
L73:
	;
	v299 = v277 + int32(1)
	goto L74
L74:
	;
	if v295 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v300 = v277
	goto L77
L76:
	;
	v300 = v256
	goto L77
L77:
	;
	if v299 < v300 {
		v256 = v300
		v259 = v299
		goto L63
	} else {
		goto L78
	}
L78:
	;
	v530 = v299
	goto L27
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v219
	F_errmsg_internal(m, int32(_a_F_width_bucket_array_0), v20+int32(16))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_1), int32(70), int32(_a_F_width_bucket_array_2))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	if v334 <= int32(0) {
		v530 = v323
		goto L27
	} else {
		goto L83
	}
L83:
	;
	if v317 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v344 = v317
	goto L86
L85:
	;
	v344 = (v320<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L86
L86:
	;
	v347 = base.B2i32(v220 != int32(-1))
	v349 = v318 - int32(99)
	v356 = v323
	v358 = v334
	v359 = v24 + v344
	goto L87
L87:
	;
	v371 = base.I32_div_s(v356+v358, int32(2))
	if v356 < v371 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v530 = v522
	goto L27
L89:
	;
	v374 = v359
	v376 = v356
	goto L92
L90:
	;
	v442 = v359
	goto L91
L91:
	;
	v458 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+68)) = uint8(v458)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v22
	if v319&int32(1) != 0 {
		goto L1
	} else {
		goto L113
	}
L92:
	;
	if v347 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v442 = v437
	goto L91
L94:
	;
	switch v349 {
	case 0:
		v437 = v424
		goto L108
	case 1:
		goto L110
	default:
		goto L109
	case 6:
		goto L111
	}
L95:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v392 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v419 = F_strlen(m, v374)
	mBase = m.M
	v424 = v419 + v374 + int32(1)
	goto L94
L98:
	;
	v396 = int32(18)
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	if v398 == v396 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v410 = int32(1)
	if v392&v410 != 0 {
		v424 = v374 + int32(base.Ui32(v392)>>(uint(v410)%32))
		goto L94
	} else {
		goto L107
	}
L101:
	;
	v401 = v396
	goto L103
L102:
	;
	v401 = int32(2)
	goto L103
L103:
	;
	if base.Ui32((v398-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v408 = int32(6)
	goto L106
L105:
	;
	v408 = v401
	goto L106
L106:
	;
	v424 = v374 + v408
	goto L94
L107:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v424 = v374 + int32(base.Ui32(v415)>>(uint(int32(2))%32))
	goto L94
L108:
	;
	v439 = v376 + int32(1)
	if v439 != v371 {
		v374 = v437
		v376 = v439
		goto L92
	} else {
		goto L112
	}
L109:
	;
	v437 = (v424 + int32(1)) & int32(-2)
	goto L108
L110:
	;
	v437 = (v424 + int32(7)) & int32(-8)
	goto L108
L111:
	;
	v437 = (v424 + int32(3)) & int32(-4)
	goto L108
L112:
	;
	goto L93
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v442
	v462 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+76)) = uint8(v462)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v468 = m.T0[v467].(func(*base.Module, int32) int32)(m, v20+int32(44))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L115
	}
L114:
	;
	if v522 < v523 {
		v356 = v522
		v358 = v523
		v359 = v524
		goto L87
	} else {
		goto L137
	}
L115:
	;
	if v468 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v522 = v356
	v523 = v371
	v524 = v359
	goto L114
L117:
	;
	goto L118
L118:
	;
	if v347 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	switch v349 {
	case 0:
		v519 = v506
		goto L133
	case 1:
		goto L135
	default:
		goto L134
	case 6:
		goto L136
	}
L120:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442))))
	if v474 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v501 = F_strlen(m, v442)
	mBase = m.M
	v506 = v501 + v442 + int32(1)
	goto L119
L123:
	;
	v478 = int32(18)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v442)+1)))
	if v480 == v478 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v492 = int32(1)
	if v474&v492 != 0 {
		v506 = v442 + int32(base.Ui32(v474)>>(uint(v492)%32))
		goto L119
	} else {
		goto L132
	}
L126:
	;
	v483 = v478
	goto L128
L127:
	;
	v483 = int32(2)
	goto L128
L128:
	;
	if base.Ui32((v480-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v490 = int32(6)
	goto L131
L130:
	;
	v490 = v483
	goto L131
L131:
	;
	v506 = v442 + v490
	goto L119
L132:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v506 = v442 + int32(base.Ui32(v497)>>(uint(int32(2))%32))
	goto L119
L133:
	;
	v522 = v371 + int32(1)
	v523 = v358
	v524 = v519
	goto L114
L134:
	;
	v519 = (v506 + int32(1)) & int32(-2)
	goto L133
L135:
	;
	v519 = (v506 + int32(7)) & int32(-8)
	goto L133
L136:
	;
	v519 = (v506 + int32(3)) & int32(-4)
	goto L133
L137:
	;
	goto L88
L138:
	;
	F_pfree(m, v24)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	m.G0 = v20 + int32(80)
	return v530
L141:
	;
	goto L140
L142:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_width_bucket_array_3), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_4), int32(_a_F_width_bucket_array_5), int32(_a_F_width_bucket_array_6))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_width_bucket_array_7), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_4), int32(_a_F_width_bucket_array_8), int32(_a_F_width_bucket_array_6))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v607 = F_format_type_be(m, v31)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v607
	F_errmsg(m, int32(_a_F_width_bucket_array_9), v20)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_4), int32(_a_F_width_bucket_array_10), int32(_a_F_width_bucket_array_6))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v220
	F_errmsg_internal(m, int32(_a_F_width_bucket_array_0), v20+int32(32))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_1), int32(70), int32(_a_F_width_bucket_array_2))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	v9 = m.G0
	v11 = v9 - int32(_a_F_write_pipe_chunks_0)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_write_pipe_chunks[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if v15 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_write_pipe_chunks[1])) = int32(8)
		v22 = int32(-1)
	} else {
		v22 = v15
	}
	v23 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11))) = uint16(v23)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v23)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_write_pipe_chunks[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v28
	v32 = int32(1)
	switch l2 - v32 {
	case 0:
		v39 = int32(17)
		v40 = int32(16)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v40)
		v44 = v39
	default:
		v44 = v32
	case 7:
		v39 = int32(33)
		v40 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v40)
		v44 = v39
	case 15:
		v39 = int32(65)
		v40 = int32(64)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v40)
		v44 = v39
	}
	if int32(4088) <= l1 {
		v49 = l0
		v50 = l1
		for {
			v57 = int32(4087)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)) = uint16(v57)
			base.MemoryCopy(m, v11+int32(9), v49, v57)
			v62 = F_write(m, v22, v11, int32(_a_F_write_pipe_chunks_0))
			mBase = m.M
			v64 = v50 - v57
			v66 = v49 + v57
			if base.Ui32(int32(_a_F_write_pipe_chunks_1)) < base.Ui32(v50) {
				v49 = v66
				v50 = v64
				continue
			} else {
				break
			}
			break
		}
		v69 = v66
		v70 = v64
	} else {
		v69 = l0
		v70 = l1
	}
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)) = uint16(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)) = uint8(v44)
	if v70 != 0 {
		base.MemoryCopy(m, v11+int32(9), v69, v70)
	} else {
	}
	v84 = F_write(m, v22, v11, v70+int32(9))
	mBase = m.M
	m.G0 = v11 + int32(_a_F_write_pipe_chunks_0)
	return
}
func F_writetup_heap_1(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v12 - int32(6)
	v17 = v9 + int32(12)
	F_LogicalTapeWrite(m, l1, v17, int32(4))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(10)
		F_LogicalTapeWrite(m, l1, v11+v21, v12-v21)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v27&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v17, int32(4))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
