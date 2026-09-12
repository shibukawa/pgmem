package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbExtractScalar(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(1342177280)
	v12 = v10 & v11
	if v12 != v11 {
		if v10&int32(1073741824) != 0 {
			v19 = int32(16)
		} else {
			v19 = int32(17)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
		m.G0 = v8 + int32(32)
		return base.B2i32(v12 == int32(1342177280))
	} else {
		v21 = F_JsonbIteratorInit(m, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v21
			v31 = F_JsonbIteratorNext(m, v8+int32(28), v8+int32(8), int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v36 = F_JsonbIteratorNext(m, v8+int32(28), l1, int32(1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v43 = F_JsonbIteratorNext(m, v8+int32(28), v8+int32(8), int32(1))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v50 = F_JsonbIteratorNext(m, v8+int32(28), v8+int32(8), int32(1))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return base.B2i32(v12 == int32(1342177280))
						}
					}
				}
			}
		}
	}
}
func F_JsonbIteratorNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v511 int32
	_ = v511
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v603
L2:
	;
	v11 = v7
	goto L5
L3:
	;
	goto L4
L4:
	;
	v597 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v597
	v603 = v597
	goto L1
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	switch v14 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	case 4:
		goto L10
	default:
		goto L9
	}
L7:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v594 = F_iteratorFromContainer(m, v592, v593)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L18
	} else {
		goto L126
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L18
	} else {
		goto L123
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L18
	} else {
		goto L120
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(3)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v391)+20))
	v395 = v393 + v394
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391)+28))
	v401 = v392 + int32(4)
	v404 = v401 + v395<<(uint(int32(2))%32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	switch int32(base.Ui32(v405)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L92
	case 1:
		goto L91
	case 2:
		goto L89
	case 3:
		goto L90
	case 4:
		goto L93
	default:
		goto L88
	}
L11:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v244) <= base.Ui32(v243) {
		goto L56
	} else {
		goto L57
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(17)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v204)+20)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v205
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v217 = v205
	v219 = v214
	goto L50
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v38) <= base.Ui32(v37) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(16)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v21)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v24
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v24
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(1)
	return int32(4)
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	F_pfree(m, v11)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v56 = v50 + int32(4)
	v59 = v56 + v37<<(uint(int32(2))%32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch int32(base.Ui32(v60)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L22
	case 3:
		goto L23
	case 4:
		goto L26
	default:
		goto L21
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(5)
L20:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v178+v179<<(uint(int32(2))%32))))
	v185 = v183 & int32(268435455)
	if int32(0) <= v183 {
		goto L46
	} else {
		goto L47
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v125 = (v52 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v51 + v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v128 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v116)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L20
L23:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v51 + (v52+int32(3))&int32(-4)
	goto L20
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v51 + v52
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v71 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L20
L27:
	;
	v76 = v37
	v77 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v71 & int32(268435455)
	goto L20
L30:
	;
	v83 = v76 - int32(1)
	if int32(0) <= v83 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v71&int32(268435455) - v95
	goto L20
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v56+v83<<(uint(int32(2))%32))))
	v92 = v89&int32(268435455) + v77
	if int32(0) <= v89 {
		v76 = v83
		v77 = v92
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v95 = v77
	goto L34
L34:
	;
	goto L31
L35:
	;
	v95 = v92
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v166 + (v52 - v125)
	goto L20
L37:
	;
	v133 = v37
	v134 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	v166 = v128 & int32(268435455)
	goto L36
L40:
	;
	v140 = v133 - int32(1)
	if int32(0) <= v140 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v166 = v128&int32(268435455) - v152
	goto L36
L42:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v56+v140<<(uint(int32(2))%32))))
	v149 = v146&int32(268435455) + v134
	if int32(0) <= v146 {
		v133 = v140
		v134 = v149
		goto L40
	} else {
		goto L45
	}
L43:
	;
	v152 = v134
	goto L44
L44:
	;
	goto L41
L45:
	;
	v152 = v149
	goto L44
L46:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v177)+24))
	v190 = v188 + v185
	goto L48
L47:
	;
	v190 = v185
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+24)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = v193 + int32(1)
	v197 = int32(3)
	if l2 != 0 {
		v603 = v197
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v198 {
	case 0, 1, 2, 3, 32:
		v603 = v197
		goto L1
	default:
		goto L7
	}
L50:
	;
	v223 = v219 - int32(1)
	if int32(0) <= v223 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210)+28)) = v235
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+32)) = int32(3)
	return int32(6)
L52:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v211+int32(4)+v223<<(uint(int32(2))%32))))
	v232 = v229&int32(268435455) + v217
	if int32(0) <= v229 {
		v217 = v232
		v219 = v223
		goto L50
	} else {
		goto L55
	}
L53:
	;
	v235 = v217
	goto L54
L54:
	;
	goto L51
L55:
	;
	v235 = v232
	goto L54
L56:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	F_pfree(m, v11)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L18
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v260 = v254 + int32(4)
	v263 = v260 + v243<<(uint(int32(2))%32)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	switch int32(base.Ui32(v264)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L65
	case 1:
		goto L64
	case 2:
		goto L62
	case 3:
		goto L63
	case 4:
		goto L66
	default:
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(7)
L60:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v381 != int32(1) {
		goto L8
	} else {
		goto L86
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v329 = (v256 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v255 + v329
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v332 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v320)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L60
L63:
	;
	v316 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v316)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v255 + (v256+int32(3))&int32(-4)
	goto L60
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v255 + v256
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v275 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L60
L67:
	;
	v280 = v243
	v281 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v275 & int32(268435455)
	goto L60
L70:
	;
	v287 = v280 - int32(1)
	if int32(0) <= v287 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v275&int32(268435455) - v299
	goto L60
L72:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v260+v287<<(uint(int32(2))%32))))
	v296 = v293&int32(268435455) + v281
	if int32(0) <= v293 {
		v280 = v287
		v281 = v296
		goto L70
	} else {
		goto L75
	}
L73:
	;
	v299 = v281
	goto L74
L74:
	;
	goto L71
L75:
	;
	v299 = v296
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v370 + (v256 - v329)
	goto L60
L77:
	;
	v337 = v243
	v338 = int32(0)
	goto L80
L78:
	;
	goto L79
L79:
	;
	v370 = v332 & int32(268435455)
	goto L76
L80:
	;
	v344 = v337 - int32(1)
	if int32(0) <= v344 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v370 = v332&int32(268435455) - v356
	goto L76
L82:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v260+v344<<(uint(int32(2))%32))))
	v353 = v350&int32(268435455) + v338
	if int32(0) <= v350 {
		v337 = v344
		v338 = v353
		goto L80
	} else {
		goto L85
	}
L83:
	;
	v356 = v338
	goto L84
L84:
	;
	goto L81
L85:
	;
	v356 = v353
	goto L84
L86:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v384)+32)) = int32(4)
	return int32(1)
L87:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+12))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v522)+20))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v523+v524<<(uint(int32(2))%32))))
	v530 = v528 & int32(268435455)
	if int32(0) <= v528 {
		goto L113
	} else {
		goto L114
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v470 = (v397 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v396 + v470
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	if v473 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L89:
	;
	v461 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v461)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L87
L90:
	;
	v457 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)) = uint8(v457)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	goto L87
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v396 + (v397+int32(3))&int32(-4)
	goto L87
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v396 + v397
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	if v416 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L87
L94:
	;
	v421 = v395
	v422 = int32(0)
	goto L97
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v416 & int32(268435455)
	goto L87
L97:
	;
	v428 = v421 - int32(1)
	if int32(0) <= v428 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v416&int32(268435455) - v440
	goto L87
L99:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v401+v428<<(uint(int32(2))%32))))
	v437 = v434&int32(268435455) + v422
	if int32(0) <= v434 {
		v421 = v428
		v422 = v437
		goto L97
	} else {
		goto L102
	}
L100:
	;
	v440 = v422
	goto L101
L101:
	;
	goto L98
L102:
	;
	v440 = v437
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v511 + (v397 - v470)
	goto L87
L104:
	;
	v478 = v395
	v479 = int32(0)
	goto L107
L105:
	;
	goto L106
L106:
	;
	v511 = v473 & int32(268435455)
	goto L103
L107:
	;
	v485 = v478 - int32(1)
	if int32(0) <= v485 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v511 = v473&int32(268435455) - v497
	goto L103
L109:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v401+v485<<(uint(int32(2))%32))))
	v494 = v491&int32(268435455) + v479
	if int32(0) <= v491 {
		v478 = v485
		v479 = v494
		goto L107
	} else {
		goto L112
	}
L110:
	;
	v497 = v479
	goto L111
L111:
	;
	goto L108
L112:
	;
	v497 = v494
	goto L111
L113:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v522)+24))
	v535 = v533 + v530
	goto L115
L114:
	;
	v535 = v530
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+24)) = v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v537)+20))
	v540 = int32(2)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v538+v539<<(uint(v540)%32)+v543<<(uint(v540)%32))))
	v549 = v547 & int32(268435455)
	if int32(0) <= v547 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v537)+28))
	v554 = v552 + v549
	goto L118
L117:
	;
	v554 = v549
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537)+28)) = v554
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+20)) = v557 + int32(1)
	v561 = int32(2)
	if l2 != 0 {
		v603 = v561
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v562 {
	case 0, 1, 2, 3, 32:
		v603 = v561
		goto L1
	default:
		goto L7
	}
L120:
	;
	F_errmsg_internal(m, int32(371368), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L18
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(523722), int32(1002), int32(70043))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L18
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errmsg_internal(m, int32(21933), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L18
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(523722), int32(967), int32(70043))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v594
	v11 = v594
	goto L5
}
func F_JsonbToCString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_JsonbToCStringWorker(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_jsonb_array_element(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
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
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v9&int32(1073741824) == int32(0) {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v20 < int32(0) {
				v24 = v9 & int32(268435455)
				if base.Ui32(v24) < base.Ui32(int32(0)-v20) {
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
					return int32(0)
				} else {
					v34 = v20 + v24
					v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 != 0 {
							v37 = F_JsonbValueToJsonb(m, v35)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								return v37
							}
						} else {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
							return int32(0)
						}
					}
				}
			} else {
				v34 = v20
				v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 != 0 {
						v37 = F_JsonbValueToJsonb(m, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							return v37
						}
					} else {
						v40 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
						return int32(0)
					}
				}
			}
		}
	}
}
func F_jsonb_array_element_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v9&int32(1073741824) == int32(0) {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v20 < int32(0) {
				v24 = v9 & int32(268435455)
				if base.Ui32(v24) < base.Ui32(int32(0)-v20) {
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
					return int32(0)
				} else {
					v34 = v20 + v24
					v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
							if v39 == int32(0) {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								return int32(0)
							} else {
								v42 = F_JsonbValueAsText(m, v35)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									return v42
								}
							}
						}
					}
				}
			} else {
				v34 = v20
				v35 = F_getIthJsonbValueFromContainer(m, v5+int32(4), v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 == int32(0) {
						v45 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
						if v39 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							return int32(0)
						} else {
							v42 = F_JsonbValueAsText(m, v35)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								return v42
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_float4(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v41 = int32(0)
							m.G0 = v8 + int32(32)
							return v41
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v41 = int32(0)
						m.G0 = v8 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(330224))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v35 = F_DirectFunctionCall1Coll(m, int32(1338), int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v11 == v37 {
							v41 = v35
							m.G0 = v8 + int32(32)
							return v41
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v8 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(330224))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_jsonb_float8(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v41 = int32(0)
							m.G0 = v8 + int32(32)
							return v41
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v41 = int32(0)
						m.G0 = v8 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(285827))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v35 = F_DirectFunctionCall1Coll(m, int32(17), int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v11 == v37 {
							v41 = v35
							m.G0 = v8 + int32(32)
							return v41
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v8 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(285827))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_jsonb_in_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	switch l2 - int32(1) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(1)
		v13 = F_strlen(m, l1)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v13
		if base.Ui32(v13) < base.Ui32(int32(268435456)) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l1
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v83 == int32(0) {
				v86 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
				v90 = int32(4)
				v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
					v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
						v102 = int32(0)
						v105 = F_pushJsonbValue(m, l0, int32(5), v102)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
							v138 = v102
							m.G0 = v7 + int32(48)
							return v138
						}
					}
				}
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
				switch v108 - int32(16) {
				case 0:
					v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
						v138 = int32(0)
						m.G0 = v7 + int32(48)
						return v138
					}
				case 1:
					v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
						v138 = int32(0)
						m.G0 = v7 + int32(48)
						return v138
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(381985), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(527082), int32(454), int32(241887))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v17 = int32(23)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = F_errsave_start(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					v138 = v17
					m.G0 = v7 + int32(48)
					return v138
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(348477), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(268435455)
							F_errdetail(m, int32(627694), v7)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v18, int32(527082), int32(284), int32(296890))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v138 = v17
									m.G0 = v7 + int32(48)
									return v138
								}
							}
						}
					}
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(2)
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v50 = F_DirectInputFunctionCallSafe(m, int32(408), l1, int32(-1), v47, v7+int32(24))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			if v50 == int32(0) {
				v138 = int32(23)
				m.G0 = v7 + int32(48)
				return v138
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
				v55 = F_pg_detoast_datum(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v55
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v83 == int32(0) {
						v86 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
						v90 = int32(4)
						v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
							v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
								v102 = int32(0)
								v105 = F_pushJsonbValue(m, l0, int32(5), v102)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
									v138 = v102
									m.G0 = v7 + int32(48)
									return v138
								}
							}
						}
					} else {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						switch v108 - int32(16) {
						case 0:
							v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
								v138 = int32(0)
								m.G0 = v7 + int32(48)
								return v138
							}
						case 1:
							v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
								v138 = int32(0)
								m.G0 = v7 + int32(48)
								return v138
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(381985), int32(0))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(527082), int32(454), int32(241887))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(388456), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(527082), int32(424), int32(241887))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 8:
		v58 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v58)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(3)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v83 == int32(0) {
			v86 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
			v90 = int32(4)
			v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
				v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
					v102 = int32(0)
					v105 = F_pushJsonbValue(m, l0, int32(5), v102)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
						v138 = v102
						m.G0 = v7 + int32(48)
						return v138
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			switch v108 - int32(16) {
			case 0:
				v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
					v138 = int32(0)
					m.G0 = v7 + int32(48)
					return v138
				}
			case 1:
				v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
					v138 = int32(0)
					m.G0 = v7 + int32(48)
					return v138
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(381985), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(527082), int32(454), int32(241887))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	case 9:
		v62 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v62)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(3)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v83 == int32(0) {
			v86 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
			v90 = int32(4)
			v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
				v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
					v102 = int32(0)
					v105 = F_pushJsonbValue(m, l0, int32(5), v102)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
						v138 = v102
						m.G0 = v7 + int32(48)
						return v138
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			switch v108 - int32(16) {
			case 0:
				v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
					v138 = int32(0)
					m.G0 = v7 + int32(48)
					return v138
				}
			case 1:
				v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
					v138 = int32(0)
					m.G0 = v7 + int32(48)
					return v138
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(381985), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(527082), int32(454), int32(241887))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	case 10:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v83 == int32(0) {
			v86 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v86)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(4294967312)
			v90 = int32(4)
			v93 = F_pushJsonbValue(m, l0, v90, v7+v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
				v99 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v99
					v102 = int32(0)
					v105 = F_pushJsonbValue(m, l0, int32(5), v102)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
						v138 = v102
						m.G0 = v7 + int32(48)
						return v138
					}
				}
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
			switch v108 - int32(16) {
			case 0:
				v133 = F_pushJsonbValue(m, l0, int32(3), v7+int32(28))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v133
					v138 = int32(0)
					m.G0 = v7 + int32(48)
					return v138
				}
			case 1:
				v114 = F_pushJsonbValue(m, l0, int32(2), v7+int32(28))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
					v138 = int32(0)
					m.G0 = v7 + int32(48)
					return v138
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(381985), int32(0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(527082), int32(454), int32(241887))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v20) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v20) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v20) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v20) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_jsonb_object_agg_strict_transfn(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_agg_unique_strict_transfn(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_agg_unique_transfn(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_field_text(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
			if v19&int32(32) == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v86 = int32(0)
				m.G0 = v9 + int32(32)
				return v86
			} else {
				v26 = int32(4)
				v28 = int32(1)
				v29 = v17 + v28
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v34 = v32 & v28
				if v34 != 0 {
					v35 = v29
				} else {
					v35 = v17 + v26
				}
				if v32 == int32(1) {
					v38 = int32(4)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
					if v40&int32(254) == int32(2) {
						v49 = v38
					} else {
						v49 = base.B2i32(v40 == int32(18)) << (uint(v38) % 32)
					}
					if v40 == int32(1) {
						v52 = v38
					} else {
						v52 = v49
					}
					v63 = v52
				} else {
					v53 = int32(1)
					if v34 != 0 {
						v63 = int32(base.Ui32(v32)>>(uint(v53)%32)) - v53
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v66 = F_getKeyJsonValueFromContainer(m, v12+v26, v35, v63, v9+int32(12))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					if v66 == int32(0) {
						v75 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
						v86 = int32(0)
						m.G0 = v9 + int32(32)
						return v86
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						if v70 == int32(0) {
							v75 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v75)
							v86 = int32(0)
							m.G0 = v9 + int32(32)
							return v86
						} else {
							v73 = F_JsonbValueAsText(m, v66)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v86 = v73
								m.G0 = v9 + int32(32)
								return v86
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_path_query_array_internal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v29 = F_executeJsonPath(m, v17, v22, int32(1413), int32(1414), v12, base.B2i32(v26 == int32(0)), v9, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v31
	v37 = F_pushJsonbValue(m, v9+int32(12), int32(4), v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v40 != 0 {
		v55 = v39
		v56 = v40
		v57 = int32(0)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v58 = v55
	v59 = v56
	goto L15
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = int32(0)
	v55 = v39
	v56 = v44
	v57 = v44
	goto L7
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if int32(1) < v50 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = v46 + int32(4)
	goto L14
L13:
	;
	v53 = int32(0)
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = v53
	v56 = v54
	v57 = v41
	goto L7
L15:
	;
	if v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v91 = F_pushJsonbValue(m, v9+int32(12), int32(5), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L17:
	;
	if v59 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v66 = int32(0)
	v80 = v66
	v81 = v66
	goto L17
L19:
	;
	goto L20
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v70 = v58 + int32(4)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if base.Ui32(v70) < base.Ui32(v72+v73<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v78 = v70
	goto L23
L22:
	;
	v78 = int32(0)
	goto L23
L23:
	;
	v80 = v68
	v81 = v78
	goto L17
L24:
	;
	v85 = F_pushJsonbValue(m, v9+int32(12), int32(3), v59)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L16
L27:
	;
	v58 = v81
	v59 = v80
	goto L15
L28:
	;
	v93 = F_JsonbValueToJsonb(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v9 + int32(16)
	return v93
}
func F_jsonb_path_query_first_tz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_first_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_populate_record(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(442396), v3, int32(1), v3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_jsonb_populate_recordset(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	F_populate_recordset_worker(m, l0, int32(113271), int32(0), int32(1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_send(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_makeStringInfo(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v21 = F_JsonbToCStringWorker(m, v13, v9+int32(4), int32(base.Ui32(v17)>>(uint(int32(2))%32)), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_pq_begintypsend(m, v6)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_enlargeStringInfo(m, v6, int32(1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v31 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v28+v29))) = uint8(v31)
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v28 + v31
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						F_pq_sendtext(m, v6, v36, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_free_attrmap(m, v13)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44 << (uint(int32(2)) % 32)
								m.G0 = v6 + int32(16)
								return v43
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_subscript_fetch(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v14 = F_jsonb_get_element(m, v8, v10, v11, v12, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v14
			return
		}
	}
}
func F_jsonb_subscript_transform(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l1 == v6 {
		v232 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v232
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(-4294963494)
	m.G0 = v14 + int32(48)
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v232 = v6
	goto L1
L4:
	;
	goto L5
L5:
	;
	v24 = v14 + int32(32) | int32(4)
	v32 = v6
	v33 = v6
	goto L7
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L15
	} else {
		goto L64
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v33<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L15
	} else {
		goto L57
	}
L9:
	;
	if v41 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if v41 != 0 {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v45 = v44
	goto L14
L13:
	;
	v45 = v41
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(182652), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v57 = F_exprLocation(m, v45)
	mBase = m.M
	F_parser_errposition(m, l2, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(520660), int32(68), int32(302634))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	goto L8
L22:
	;
	v164 = int32(-1)
	v168 = F_coerce_type(m, l2, v68, v163, v162, v164, int32(0), int32(2), v164)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L15
	} else {
		goto L53
	}
L23:
	;
	if v158 == int32(705) {
		goto L21
	} else {
		goto L52
	}
L24:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v158 = v157
	goto L23
L25:
	;
	v152 = F_can_coerce_type(m, int32(1), v14+int32(44), v24, int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L15
	} else {
		goto L50
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L43
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v68 = F_transformExpr(m, l2, v41, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L15
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L38
	}
L30:
	;
	v70 = F_exprType(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v70
	v74 = int32(705)
	if v70 == v74 {
		v162 = int32(25)
		v163 = v74
		goto L22
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(107374182423)
	v85 = F_can_coerce_type(m, int32(1), v14+int32(44), v14+int32(32), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	if v85 == int32(0) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v94 = F_can_coerce_type(m, int32(1), v14+int32(44), v24, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v94 == int32(0) {
		v158 = v89
		goto L23
	} else {
		goto L36
	}
L36:
	;
	if v89 != int32(705) {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L24
L38:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(182652), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v112 = F_exprLocation(m, v111)
	mBase = m.M
	F_parser_errposition(m, l2, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(520660), int32(149), int32(302634))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v128 = F_format_type_be(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v128
	F_errmsg(m, int32(464807), v14+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	F_errhint(m, int32(607881), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v140 = F_exprLocation(m, v68)
	mBase = m.M
	F_parser_errposition(m, l2, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(520660), int32(102), int32(302634))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if v152 == int32(0) {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	goto L24
L52:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v162 = v158
	v163 = v161
	goto L22
L53:
	;
	if v168 == int32(0) {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v172 = F_lappend(m, v32, v168)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v175 = v33 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v176 <= v175 {
		v232 = v172
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v32 = v172
	v33 = v175
	goto L7
L57:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v187 = F_format_type_be(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v187
	F_errmsg(m, int32(464807), v14)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(607820), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	v197 = F_exprLocation(m, v68)
	mBase = m.M
	F_parser_errposition(m, l2, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(520660), int32(116), int32(302634))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L15
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(386923), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	v217 = F_exprLocation(m, int32(0))
	mBase = m.M
	F_parser_errposition(m, l2, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(520660), int32(137), int32(302634))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_jsonb_index_flags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = F_JsonbIteratorInit(m, l0+int32(4))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v10
	v20 = F_JsonbIteratorNext(m, v6+int32(28), v6+int32(8), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v20 == int32(4) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L136
	}
L7:
	;
	v33 = F_JsonbIteratorNext(m, v6+int32(28), v6+int32(8), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L133
	}
L9:
	;
	goto L8
L10:
	;
	v25 = v25 | int32(2)
	goto L7
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L128
	}
L12:
	;
	v381 = F_JsonbIteratorNext(m, v6+int32(28), v6+int32(8), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L126
	}
L13:
	;
	if v33 != int32(3) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v33 == int32(5) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if v52 != int32(1) {
		goto L11
	} else {
		goto L21
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(26363), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(520272), int32(5649), int32(166903))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v55 != int32(3) {
		v175 = v55
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v175 == int32(6) {
		goto L62
	} else {
		goto L63
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v64 = v59
	v65 = int32(320829)
	v66 = int32(3)
	goto L25
L24:
	;
	if v111 == int32(0) {
		v25 = int32(15)
		goto L7
	} else {
		goto L40
	}
L25:
	;
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v111 = int32(0)
	goto L24
L27:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		v92 = v69
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v94 = int32(1)
	if v92 != 0 {
		v64 = v64 + v94
		v65 = v65 + v94
		v66 = v66 - v94
		goto L25
	} else {
		goto L39
	}
L31:
	;
	if base.Ui32((v69-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v80 = v69 | int32(32)
	goto L34
L33:
	;
	v80 = v69
	goto L34
L34:
	;
	if base.Ui32((v70-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v89 = v70 | int32(32)
	goto L37
L36:
	;
	v89 = v70
	goto L37
L37:
	;
	if v80 == v89 {
		v92 = v80
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v111 = v80 - v89
	goto L24
L39:
	;
	goto L29
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v114 != int32(3) {
		v175 = v114
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v122 = v117
	v123 = int32(22949)
	v124 = int32(3)
	goto L43
L42:
	;
	if v169 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L43:
	;
	if v124 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v169 = int32(0)
	goto L42
L45:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v127 == v128 {
		v150 = v127
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v152 = int32(1)
	if v150 != 0 {
		v122 = v122 + v152
		v123 = v123 + v152
		v124 = v124 - v152
		goto L43
	} else {
		goto L57
	}
L49:
	;
	if base.Ui32((v127-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v138 = v127 | int32(32)
	goto L52
L51:
	;
	v138 = v127
	goto L52
L52:
	;
	if base.Ui32((v128-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v147 = v128 | int32(32)
	goto L55
L54:
	;
	v147 = v128
	goto L55
L55:
	;
	if v138 == v147 {
		v150 = v138
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v169 = v138 - v147
	goto L42
L57:
	;
	goto L47
L58:
	;
	v25 = v25 | int32(1)
	goto L7
L59:
	;
	goto L60
L60:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v175 = v174
	goto L22
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L120
	}
L62:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v183 = v178
	v184 = int32(348609)
	v185 = int32(6)
	goto L66
L63:
	;
	v234 = v175
	goto L64
L64:
	;
	if v234 != int32(7) {
		goto L61
	} else {
		goto L82
	}
L65:
	;
	if v230 == int32(0) {
		goto L10
	} else {
		goto L81
	}
L66:
	;
	if v185 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v230 = int32(0)
	goto L65
L68:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v188 == v189 {
		v211 = v188
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v213 = int32(1)
	if v211 != 0 {
		v183 = v183 + v213
		v184 = v184 + v213
		v185 = v185 - v213
		goto L66
	} else {
		goto L80
	}
L72:
	;
	if base.Ui32((v188-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v199 = v188 | int32(32)
	goto L75
L74:
	;
	v199 = v188
	goto L75
L75:
	;
	if base.Ui32((v189-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v208 = v189 | int32(32)
	goto L78
L77:
	;
	v208 = v189
	goto L78
L78:
	;
	if v199 == v208 {
		v211 = v199
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v230 = v199 - v208
	goto L65
L80:
	;
	goto L70
L81:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v234 = v233
	goto L64
L82:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v242 = v237
	v243 = int32(515272)
	v244 = int32(7)
	goto L84
L83:
	;
	if v289 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L84:
	;
	if v244 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v289 = int32(0)
	goto L83
L86:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v247 == v248 {
		v270 = v247
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	v272 = int32(1)
	if v270 != 0 {
		v242 = v242 + v272
		v243 = v243 + v272
		v244 = v244 - v272
		goto L84
	} else {
		goto L98
	}
L90:
	;
	if base.Ui32((v247-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v258 = v247 | int32(32)
	goto L93
L92:
	;
	v258 = v247
	goto L93
L93:
	;
	if base.Ui32((v248-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v267 = v248 | int32(32)
	goto L96
L95:
	;
	v267 = v248
	goto L96
L96:
	;
	if v258 == v267 {
		v270 = v258
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v289 = v258 - v267
	goto L83
L98:
	;
	goto L88
L99:
	;
	v25 = v25 | int32(4)
	goto L7
L100:
	;
	goto L101
L101:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v294 != int32(7) {
		goto L61
	} else {
		goto L102
	}
L102:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v302 = v297
	v303 = int32(298479)
	v304 = int32(7)
	goto L104
L103:
	;
	if v349 != 0 {
		goto L61
	} else {
		goto L119
	}
L104:
	;
	if v304 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v349 = int32(0)
	goto L103
L106:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v307 == v308 {
		v330 = v307
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	goto L105
L109:
	;
	v332 = int32(1)
	if v330 != 0 {
		v302 = v302 + v332
		v303 = v303 + v332
		v304 = v304 - v332
		goto L104
	} else {
		goto L118
	}
L110:
	;
	if base.Ui32((v307-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v318 = v307 | int32(32)
	goto L113
L112:
	;
	v318 = v307
	goto L113
L113:
	;
	if base.Ui32((v308-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v327 = v308 | int32(32)
	goto L116
L115:
	;
	v327 = v308
	goto L116
L116:
	;
	if v318 == v327 {
		v330 = v318
		goto L109
	} else {
		goto L117
	}
L117:
	;
	v349 = v318 - v327
	goto L103
L118:
	;
	goto L108
L119:
	;
	v25 = v25 | int32(8)
	goto L7
L120:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v361 = F_pnstrdup(m, v359, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v361
	F_errmsg(m, int32(761718), v6)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errhint(m, int32(701396), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(520272), int32(5644), int32(166903))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	if v381 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	m.G0 = v6 + int32(32)
	return v25
L128:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(348522), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errhint(m, int32(701396), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(520272), int32(5622), int32(166903))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(26363), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(520272), int32(5654), int32(166903))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(462510), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(520272), int32(5614), int32(166903))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
