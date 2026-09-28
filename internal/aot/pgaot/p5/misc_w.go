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
	m.G2 = int32(13210208)
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
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[1])) = int32(13210208)
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[2])) = int32(42)
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[3])) = int32(_a_F___wasm_call_ctors_1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[4]))
	*(*int32)(unsafe.Add(mBase, _c_F___wasm_call_ctors[5])) = v60
	return
}
func F_websearch_to_tsquery_byid(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14390(m, l0, int32(2), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
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
func F_width_bucket_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v206 int32
	_ = v206
	var v212 float64
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v304 int64
	_ = v304
	var v305 int64
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v506 int32
	_ = v506
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v511 int64
	_ = v511
	var v512 int64
	_ = v512
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int64
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	v21 = m.G0
	v23 = v21 - int32(112)
	m.G0 = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L3
	} else {
		goto L165
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L3
	} else {
		goto L161
	}
L3:
	;
	return int64(0)
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v31 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v36 == int32(0) {
		v149 = int32(0)
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L3
	} else {
		goto L157
	}
L8:
	;
	if v34 == int32(701) {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v40 = v27 + int32(16)
	v41 = F_ArrayGetNItemsSafe(m, v31, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v49 = v40 + v44<<(uint(int32(3))%32)
	goto L13
L12:
	;
	v49 = int32(0)
	goto L13
L13:
	;
	if int32(8) <= v41 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v60 = v41
	v61 = v49
	goto L17
L15:
	;
	v89 = v41
	v90 = v49
	goto L16
L16:
	;
	if v89 <= int32(0) {
		v149 = v43
		goto L8
	} else {
		goto L21
	}
L17:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v79 != int32(255) {
		goto L2
	} else {
		goto L19
	}
L18:
	;
	v89 = v87
	v90 = v49 + int32(base.Ui32(v41-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L16
L19:
	;
	v87 = v60 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v60) {
		v60 = v87
		v61 = v61 + int32(1)
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	v113 = v89
	v114 = int32(1)
	goto L22
L22:
	;
	if v114&v110 == int32(0) {
		goto L2
	} else {
		goto L24
	}
L23:
	;
	v149 = v43
	goto L8
L24:
	;
	v135 = int32(1)
	if v135 < v113 {
		v113 = v113 - v135
		v114 = v114 << (uint(v135) % 32)
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v607 != v27 {
		goto L153
	} else {
		goto L154
	}
L27:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v166 = F_ArrayGetNItemsSafe(m, v163, v27+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v224 != 0 {
		goto L48
	} else {
		goto L49
	}
L30:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v25&int64(9223372036854775807)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v591 = v166
	goto L26
L32:
	;
	goto L33
L33:
	;
	v172 = int32(0)
	if v166 <= v172 {
		v591 = v172
		goto L26
	} else {
		goto L34
	}
L34:
	;
	if v149 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v181 = v149
	goto L37
L36:
	;
	v181 = (v163<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L37
L37:
	;
	v185 = v166
	v188 = v172
	goto L38
L38:
	;
	v206 = base.I32_div_s(v185+v188, int32(2))
	v212 = *(*float64)(unsafe.Add(mBase, uint32(v27+v181+v206<<(uint(int32(3))%32))))
	v219 = base.F64_gt(v212, base.F64_reinterpret_i64(v25)) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v212)&int64(9223372036854775807)))
	if v219 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v591 = v220
	goto L26
L40:
	;
	v220 = v188
	goto L42
L41:
	;
	v220 = v206 + int32(1)
	goto L42
L42:
	;
	if v219 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v221 = v206
	goto L45
L44:
	;
	v221 = v185
	goto L45
L45:
	;
	if v220 < v221 {
		v185 = v221
		v188 = v220
		goto L38
	} else {
		goto L46
	}
L46:
	;
	goto L39
L47:
	;
	v237 = int32(*(*int16)(unsafe.Add(mBase, uint32(v235)+8)))
	if int32(0) < v237 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if v225 == v34 {
		v235 = v224
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v228 = F_lookup_type_cache(m, v34, int32(64))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v228)+108))
	if v230 == int32(0) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+16)) = v228
	v235 = v228
	goto L47
L54:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+10)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v243 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+74)) = uint16(v243)
	v245 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+72)) = uint8(v245)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v23)+60)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v235 + int32(104)
	v256 = F_ArrayGetNItemsSafe(m, v242, v27+int32(16))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+10)))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+11)))
	switch v342 - int32(99) {
	case 0:
		v364 = int32(1)
		goto L85
	case 1:
		goto L88
	default:
		goto L87
	case 6:
		goto L89
	case 16:
		goto L86
	}
L57:
	;
	if v256 <= int32(0) {
		v591 = v245
		goto L26
	} else {
		goto L58
	}
L58:
	;
	if v241 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v266 = v241
	goto L61
L60:
	;
	v266 = (v242<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L61
L61:
	;
	v269 = int32(1)
	v275 = v256
	v278 = v245
	goto L62
L62:
	;
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+88)) = uint8(v294)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v25
	v299 = base.I32_div_s(v275+v278, int32(2))
	v301 = v27 + v266 + v299*v237
	if v240&v269 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v591 = v338
	goto L26
L64:
	;
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+104)) = uint8(v323)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v322
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v332 = m.T0[v331].(func(*base.Module, int32) int64)(m, v23+int32(56))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L3
	} else {
		goto L77
	}
L65:
	;
	if base.Ui32(v269) < base.Ui32(base.I32_popcnt(v237)) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v322 = base.I64_extend_i32_u(v301)
	goto L64
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L74
	}
L69:
	;
	switch base.I32_ctz(v237) {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	case 3:
		goto L70
	default:
		goto L68
	}
L70:
	;
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
	v322 = v305
	goto L64
L71:
	;
	v304 = int64(*(*int32)(unsafe.Add(mBase, uint32(v301))))
	v322 = v304
	goto L64
L72:
	;
	v303 = int64(*(*int16)(unsafe.Add(mBase, uint32(v301))))
	v322 = v303
	goto L64
L73:
	;
	v302 = int64(*(*int8)(unsafe.Add(mBase, uint32(v301))))
	v322 = v302
	goto L64
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v237
	F_errmsg_internal(m, int32(_a_F_width_bucket_array_0), v23+int32(16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_1), int32(123), int32(_a_F_width_bucket_array_2))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v337 = base.B2i32(v332&int64(2147483648) == int64(0))
	if v332&int64(2147483648) == int64(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v338 = v299 + int32(1)
	goto L80
L79:
	;
	v338 = v278
	goto L80
L80:
	;
	if v332&int64(2147483648) == int64(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v339 = v275
	goto L83
L82:
	;
	v339 = v299
	goto L83
L83:
	;
	if v338 < v339 {
		v275 = v339
		v278 = v338
		goto L62
	} else {
		goto L84
	}
L84:
	;
	goto L63
L85:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v367 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+74)) = uint16(v367)
	v369 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+72)) = uint8(v369)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v23)+60)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v235 + int32(104)
	v380 = F_ArrayGetNItemsSafe(m, v366, v27+int32(16))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L3
	} else {
		goto L93
	}
L86:
	;
	v364 = int32(2)
	goto L85
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L3
	} else {
		goto L90
	}
L88:
	;
	v364 = int32(8)
	goto L85
L89:
	;
	v364 = int32(4)
	goto L85
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = base.I32_extend8_s(v342)
	F_errmsg_internal(m, int32(_a_F_width_bucket_array_3), v23+int32(32))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_1), int32(322), int32(_a_F_width_bucket_array_4))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	if v380 <= int32(0) {
		v591 = v369
		goto L26
	} else {
		goto L94
	}
L94:
	;
	v385 = int32(0) - v364
	v387 = v364 - int32(1)
	if v365 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v394 = v365
	goto L97
L96:
	;
	v394 = (v366<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L97
L97:
	;
	v397 = int32(1)
	v406 = v369
	v408 = v380
	v410 = v27 + v394
	goto L98
L98:
	;
	v424 = base.I32_div_s(v406+v408, int32(2))
	if v406 < v424 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v591 = v583
	goto L26
L100:
	;
	v427 = v410
	v428 = v406
	goto L103
L101:
	;
	v487 = v410
	goto L102
L102:
	;
	v506 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+88)) = uint8(v506)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v25
	if v341&v397 != 0 {
		goto L121
	} else {
		goto L122
	}
L103:
	;
	if v237 == int32(-1) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v487 = v482
	goto L102
L105:
	;
	v482 = (v480 + v387) & v385
	v484 = v428 + int32(1)
	if v484 != v424 {
		v427 = v482
		v428 = v484
		goto L103
	} else {
		goto L119
	}
L106:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v448 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v475 = F_strlen(m, v427)
	mBase = m.M
	v480 = v475 + v427 + int32(1)
	goto L105
L109:
	;
	v452 = int32(18)
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	if v454 == v452 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	v466 = int32(1)
	if v448&v466 != 0 {
		v480 = v427 + int32(base.Ui32(v448)>>(uint(v466)%32))
		goto L105
	} else {
		goto L118
	}
L112:
	;
	v457 = v452
	goto L114
L113:
	;
	v457 = int32(2)
	goto L114
L114:
	;
	if base.Ui32((v454-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v464 = int32(6)
	goto L117
L116:
	;
	v464 = v457
	goto L117
L117:
	;
	v480 = v427 + v464
	goto L105
L118:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v480 = v427 + int32(base.Ui32(v471)>>(uint(int32(2))%32))
	goto L105
L119:
	;
	goto L104
L120:
	;
	v530 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+104)) = uint8(v530)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v529
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	v537 = m.T0[v536].(func(*base.Module, int32) int64)(m, v23+int32(56))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L3
	} else {
		goto L134
	}
L121:
	;
	if base.I32_popcnt(v237) != v397 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v529 = base.I64_extend_i32_u(v487)
	goto L120
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L3
	} else {
		goto L130
	}
L125:
	;
	switch base.I32_ctz(v237) {
	case 0:
		goto L129
	case 1:
		goto L128
	case 2:
		goto L127
	case 3:
		goto L126
	default:
		goto L124
	}
L126:
	;
	v512 = *(*int64)(unsafe.Add(mBase, uint32(v487)))
	v529 = v512
	goto L120
L127:
	;
	v511 = int64(*(*int32)(unsafe.Add(mBase, uint32(v487))))
	v529 = v511
	goto L120
L128:
	;
	v510 = int64(*(*int16)(unsafe.Add(mBase, uint32(v487))))
	v529 = v510
	goto L120
L129:
	;
	v509 = int64(*(*int8)(unsafe.Add(mBase, uint32(v487))))
	v529 = v509
	goto L120
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v237
	F_errmsg_internal(m, int32(_a_F_width_bucket_array_0), v23+int32(48))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_1), int32(123), int32(_a_F_width_bucket_array_2))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	if v583 < v584 {
		v406 = v583
		v408 = v584
		v410 = v585
		goto L98
	} else {
		goto L152
	}
L134:
	;
	if v537&int64(2147483648) != int64(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v583 = v406
	v584 = v424
	v585 = v410
	goto L133
L136:
	;
	goto L137
L137:
	;
	if v237 == int32(-1) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v583 = v424 + int32(1)
	v584 = v408
	v585 = (v577 + v387) & v385
	goto L133
L139:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487))))
	if v545 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L141
L141:
	;
	v572 = F_strlen(m, v487)
	mBase = m.M
	v577 = v572 + v487 + int32(1)
	goto L138
L142:
	;
	v549 = int32(18)
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
	if v551 == v549 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v563 = int32(1)
	if v545&v563 != 0 {
		v577 = v487 + int32(base.Ui32(v545)>>(uint(v563)%32))
		goto L138
	} else {
		goto L151
	}
L145:
	;
	v554 = v549
	goto L147
L146:
	;
	v554 = int32(2)
	goto L147
L147:
	;
	if base.Ui32((v551-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v561 = int32(6)
	goto L150
L149:
	;
	v561 = v554
	goto L150
L150:
	;
	v577 = v487 + v561
	goto L138
L151:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	v577 = v487 + int32(base.Ui32(v568)>>(uint(int32(2))%32))
	goto L138
L152:
	;
	goto L99
L153:
	;
	F_pfree(m, v27)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L3
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	m.G0 = v23 + int32(112)
	return base.I64_extend_i32_s(v591)
L156:
	;
	goto L155
L157:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_width_bucket_array_5), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_6), int32(_a_F_width_bucket_array_7), int32(_a_F_width_bucket_array_8))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(_a_F_width_bucket_array_9), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_6), int32(_a_F_width_bucket_array_10), int32(_a_F_width_bucket_array_8))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L3
	} else {
		goto L166
	}
L166:
	;
	v675 = F_format_type_be(m, v34)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v675
	F_errmsg(m, int32(_a_F_width_bucket_array_11), v23)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_width_bucket_array_6), int32(_a_F_width_bucket_array_12), int32(_a_F_width_bucket_array_8))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
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
