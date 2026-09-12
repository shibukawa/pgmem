package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WalkInnerWith(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v4 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v7 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v12 = F_lcons(m, v10, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v54 = F_lcons(m, int32(0), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L15
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v24 = v4
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v24<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v33 = F_makeDependencyGraphWalker(m, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v36 = v24 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 < v37 {
		v24 = v36
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v49 = F_list_delete_first(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v49
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v96 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l2)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L27
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v66 = v4
	goto L19
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v66<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v75 = F_makeDependencyGraphWalker(m, v74, l2)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v77 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v80 = v78
	goto L24
L23:
	;
	v80 = int32(0)
	goto L24
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = F_lappend(m, v81, v73)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v82
	v86 = v66 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v86 < v87 {
		v66 = v86
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L20
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v99 = F_list_delete_first(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v99
	return
}
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
	*(*int32)(unsafe.Add(mBase, _consts[1])) = int32(4607936)
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
			v24 = F_parse_tsquery(m, v17, int32(1173), v6+int32(8), int32(2), int32(0))
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
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
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
	v753 = m.ExcPending
	if v753 != 0 {
		goto L3
	} else {
		goto L201
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L3
	} else {
		goto L197
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
	v703 = m.ExcPending
	if v703 != 0 {
		goto L3
	} else {
		goto L193
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
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v692 != v25 {
		goto L189
	} else {
		goto L190
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
	v680 = v157
	goto L28
L34:
	;
	goto L35
L35:
	;
	v164 = int32(0)
	if v157 <= v164 {
		v680 = v164
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
	v680 = v209
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
		v680 = v234
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
	v680 = v306
	goto L28
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v225
	F_errmsg_internal(m, int32(462348), v21+int32(16))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L3
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(311596), int32(70), int32(64767))
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
		v680 = v330
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
	v680 = v670
	goto L28
L90:
	;
	v382 = v367
	v383 = v365
	goto L93
L91:
	;
	v509 = v367
	goto L92
L92:
	;
	v526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+68)) = uint8(v526)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v23
	if v326&int32(1) != 0 {
		goto L135
	} else {
		goto L136
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
	v509 = v504
	goto L92
L95:
	;
	switch v356 {
	case 0:
		v504 = v491
		goto L129
	case 1:
		goto L131
	default:
		goto L130
	case 6:
		goto L132
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
	if v382&int32(3) == int32(0) {
		v453 = v382
		goto L114
	} else {
		goto L115
	}
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
		v491 = v382 + int32(base.Ui32(v401)>>(uint(v421)%32))
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
	v491 = v382 + v419
	goto L95
L111:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v491 = v382 + int32(base.Ui32(v426)>>(uint(int32(2))%32))
	goto L95
L112:
	;
	v491 = v486 + v382 + int32(1)
	goto L95
L113:
	;
	v486 = v478 - v382
	goto L112
L114:
	;
	v457 = v453
	goto L123
L115:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v437 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v486 = int32(0)
	goto L112
L117:
	;
	goto L118
L118:
	;
	v442 = v382
	goto L119
L119:
	;
	v446 = v442 + int32(1)
	if v446&int32(3) == int32(0) {
		v453 = v446
		goto L114
	} else {
		goto L121
	}
L120:
	;
	v478 = v446
	goto L113
L121:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	if v451 != 0 {
		v442 = v446
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v466 = int32(-2139062144)
	if (int32(16843008)-v463|v463)&v466 == v466 {
		v457 = v457 + int32(4)
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v472 = v457
	goto L126
L125:
	;
	goto L124
L126:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472))))
	if v476 != 0 {
		v472 = v472 + int32(1)
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v478 = v472
	goto L113
L128:
	;
	goto L127
L129:
	;
	v506 = v383 + int32(1)
	if v506 != v379 {
		v382 = v504
		v383 = v506
		goto L93
	} else {
		goto L133
	}
L130:
	;
	v504 = (v491 + int32(1)) & int32(-2)
	goto L129
L131:
	;
	v504 = (v491 + int32(7)) & int32(-8)
	goto L129
L132:
	;
	v504 = (v491 + int32(3)) & int32(-4)
	goto L129
L133:
	;
	goto L94
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v549
	v551 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+76)) = uint8(v551)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	v557 = m.T0[v556].(func(*base.Module, int32) int32)(m, v21+int32(44))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L3
	} else {
		goto L146
	}
L135:
	;
	switch v225 - int32(1) {
	case 0:
		goto L141
	case 1:
		goto L140
	default:
		goto L138
	case 3:
		goto L139
	}
L136:
	;
	goto L137
L137:
	;
	v549 = v509
	goto L134
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L142
	}
L139:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v549 = v533
	goto L134
L140:
	;
	v532 = int32(*(*int16)(unsafe.Add(mBase, uint32(v509))))
	v549 = v532
	goto L134
L141:
	;
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v509))))
	v549 = v531
	goto L134
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v226
	F_errmsg_internal(m, int32(462348), v21+int32(32))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L3
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(311596), int32(70), int32(64767))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L3
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	if v670 < v672 {
		v365 = v670
		v367 = v671
		v373 = v672
		goto L88
	} else {
		goto L188
	}
L146:
	;
	if v557 < int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v670 = v365
	v671 = v367
	v672 = v379
	goto L145
L148:
	;
	goto L149
L149:
	;
	if v354 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	switch v356 {
	case 0:
		v666 = v653
		goto L184
	case 1:
		goto L186
	default:
		goto L185
	case 6:
		goto L187
	}
L151:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	if v563 == int32(1) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	if v509&int32(3) == int32(0) {
		v615 = v509
		goto L169
	} else {
		goto L170
	}
L154:
	;
	v566 = int32(6)
	v568 = int32(18)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
	if v570 == v568 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v583 = int32(1)
	if v563&v583 != 0 {
		v653 = v509 + int32(base.Ui32(v563)>>(uint(v583)%32))
		goto L150
	} else {
		goto L166
	}
L157:
	;
	v573 = v568
	goto L159
L158:
	;
	v573 = int32(2)
	goto L159
L159:
	;
	if v570&int32(254) == int32(2) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v578 = v566
	goto L162
L161:
	;
	v578 = v573
	goto L162
L162:
	;
	if v570 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v581 = v566
	goto L165
L164:
	;
	v581 = v578
	goto L165
L165:
	;
	v653 = v509 + v581
	goto L150
L166:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v653 = v509 + int32(base.Ui32(v588)>>(uint(int32(2))%32))
	goto L150
L167:
	;
	v653 = v648 + v509 + int32(1)
	goto L150
L168:
	;
	v648 = v640 - v509
	goto L167
L169:
	;
	v619 = v615
	goto L178
L170:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509))))
	if v599 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v648 = int32(0)
	goto L167
L172:
	;
	goto L173
L173:
	;
	v604 = v509
	goto L174
L174:
	;
	v608 = v604 + int32(1)
	if v608&int32(3) == int32(0) {
		v615 = v608
		goto L169
	} else {
		goto L176
	}
L175:
	;
	v640 = v608
	goto L168
L176:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	if v613 != 0 {
		v604 = v608
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v628 = int32(-2139062144)
	if (int32(16843008)-v625|v625)&v628 == v628 {
		v619 = v619 + int32(4)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v634 = v619
	goto L181
L180:
	;
	goto L179
L181:
	;
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	if v638 != 0 {
		v634 = v634 + int32(1)
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v640 = v634
	goto L168
L183:
	;
	goto L182
L184:
	;
	v670 = v379 + int32(1)
	v671 = v666
	v672 = v373
	goto L145
L185:
	;
	v666 = (v653 + int32(1)) & int32(-2)
	goto L184
L186:
	;
	v666 = (v653 + int32(7)) & int32(-8)
	goto L184
L187:
	;
	v666 = (v653 + int32(3)) & int32(-4)
	goto L184
L188:
	;
	goto L89
L189:
	;
	F_pfree(m, v25)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L3
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	m.G0 = v21 + int32(80)
	return v680
L192:
	;
	goto L191
L193:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L3
	} else {
		goto L194
	}
L194:
	;
	F_errmsg(m, int32(24009), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L3
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(473073), int32(6708), int32(22619))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L3
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L3
	} else {
		goto L198
	}
L198:
	;
	F_errmsg(m, int32(165385), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(473073), int32(6713), int32(22619))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L3
	} else {
		goto L202
	}
L202:
	;
	v757 = F_format_type_be(m, v32)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L3
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v757
	F_errmsg(m, int32(179969), v21)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L3
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(473073), int32(6733), int32(22619))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L3
	} else {
		goto L205
	}
L205:
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[467]))
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
	v33 = *(*int32)(unsafe.Add(mBase, _consts[451]))
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
