package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InsertPgClassTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(192)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v11 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v12
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+32)) = uint16(v6)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+76)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+116)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+108)) = v58
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v60
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+116)) = v62
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+120)) = v64
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+122)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+124)) = v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+124)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+125)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = v70
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+127)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v74
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+126)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+136)) = v76
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+129)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v78
	v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+130)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+131)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = v88
	if l3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = l3
	} else {
		v91 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)) = uint8(v91)
	}
	if l4 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = l4
	} else {
		v94 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)) = uint8(v94)
	}
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+33)) = uint8(v96)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v101 = F_heap_form_tuple(m, v98, v9+int32(48), v9)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		return
	} else {
		F_CatalogTupleInsert(m, l0, v101)
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return
		} else {
			F_pfree(m, v101)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				m.G0 = v9 + int32(192)
				return
			}
		}
	}
}
func F_MakePGDirectory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_MakePGDirectory[0]))
	v4 = F_mkdir(m, l0, v3)
	mBase = m.M
	return v4
}
func F_PGProcShmemSize(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_PGProcShmemSize[0]))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_PGProcShmemSize[1]))
	v8 = F_add_size(m, int32(38), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_add_size(m, v4, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_mul_size(m, v12, int32(640))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_add_size(m, int32(0), v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v20 = F_mul_size(m, v12, int32(4))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_add_size(m, v17, v20)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v25 = F_mul_size(m, v12, int32(2))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								v27 = F_add_size(m, v22, v25)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return int32(0)
								} else {
									v30 = F_mul_size(m, v12, int32(1))
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return int32(0)
									} else {
										v32 = F_add_size(m, v27, v30)
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return int32(0)
										} else {
											return v32
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
func F_PG_char_to_encoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = int32(-1)
	if v2 == int32(0) {
		v89 = v14
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v89
L2:
	;
	m.G0 = v12 - int32(-64)
	goto L1
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
	if v17 == int32(0) {
		v89 = v14
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = F_strlen(m, v2)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v20) {
		v89 = v14
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = v2
	v24 = v17
	v25 = v12
	goto L6
L6:
	;
	v33 = F_isalnum(m, v24&int32(255))
	mBase = m.M
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v50)
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
	v55 = int32(_a_F_PG_char_to_encoding_0)
	v56 = int32(_a_F_PG_char_to_encoding_1)
	goto L15
L8:
	;
	if base.Ui32((v24-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v46 = v25
	goto L10
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v47 != 0 {
		v23 = v23 + int32(1)
		v24 = v47
		v25 = v46
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v42 = v24 | int32(32)
	goto L13
L12:
	;
	v42 = v24
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v42)
	v46 = v25 + int32(1)
	goto L10
L14:
	;
	goto L7
L15:
	;
	v68 = v56 + (v55-v56)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v69))))
	v71 = v54 - v70
	if v71 != 0 {
		v74 = v71
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v89 = v14
	goto L2
L17:
	;
	v78 = base.B2i32(v74 < int32(0))
	if v74 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v72 = F_strcmp(m, v12, v69)
	mBase = m.M
	if v72 != 0 {
		v74 = v72
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v89 = v73
	goto L2
L20:
	;
	v79 = v68 - int32(8)
	goto L22
L21:
	;
	v79 = v55
	goto L22
L22:
	;
	if v74 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v56
	goto L25
L24:
	;
	v82 = v68 + int32(8)
	goto L25
L25:
	;
	if base.Ui32(v82) <= base.Ui32(v79) {
		v55 = v79
		v56 = v82
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
}
func F_PgArchiverMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int64
	_ = v376
	var v378 int64
	_ = v378
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v542 int64
	_ = v542
	var v544 int64
	_ = v544
	var v556 int32
	_ = v556
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int64
	_ = v757
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1452 int32
	_ = v1452
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1554 int32
	_ = v1554
	var v1555 int64
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1569 int32
	_ = v1569
	var v1570 int64
	_ = v1570
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1580 int64
	_ = v1580
	var v1582 int64
	_ = v1582
	var v1584 int64
	_ = v1584
	var v1586 int64
	_ = v1586
	var v1588 int64
	_ = v1588
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1612 int32
	_ = v1612
	var v1613 int64
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1627 int32
	_ = v1627
	var v1628 int64
	_ = v1628
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int64
	_ = v1638
	var v1640 int64
	_ = v1640
	var v1642 int64
	_ = v1642
	var v1644 int64
	_ = v1644
	var v1646 int64
	_ = v1646
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1713 int32
	_ = v1713
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1739 int32
	_ = v1739
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[0])) = int32(9)
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = int32(914)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	switch int32(916) {
	case 0, 2:
		v32 = v18
		goto L4
	default:
		goto L5
	}
L3:
	;
	v64 = int32(-2)
	v66 = m.G0
	v68 = v66 - int32(32)
	m.G0 = v68
	switch int32(0) {
	case 0, 2:
		v78 = v64
		goto L17
	default:
		goto L18
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v32
	F_sigemptyset(m, v22+int32(16))
	mBase = m.M
	goto L7
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[1])) = v18
	v32 = int32(_a_F_PgArchiverMain_0)
	goto L4
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(268435456)
	v44 = v22 + int32(12)
	goto L11
L9:
	;
	m.G0 = v22 + int32(32)
	goto L3
L11:
	;
	goto L12
L12:
	;
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v50 = int32(20)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[2])) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[3])) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[4])) = v56
	goto L15
L14:
	;
	goto L15
L15:
	;
	goto L9
L16:
	;
	v110 = int32(916)
	v112 = m.G0
	v114 = v112 - int32(32)
	m.G0 = v114
	switch int32(918) {
	case 0, 2:
		v124 = v110
		goto L30
	default:
		goto L31
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = v78
	F_sigemptyset(m, v68+int32(16))
	mBase = m.M
	goto L20
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[5])) = v64
	v78 = int32(_a_F_PgArchiverMain_0)
	goto L17
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(268435456)
	v90 = v68 + int32(12)
	goto L24
L22:
	;
	m.G0 = v68 + int32(32)
	goto L16
L24:
	;
	goto L25
L25:
	;
	if v90 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v97 = int32(40)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[6])) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v90)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[7])) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v90)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[8])) = v102
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L22
L29:
	;
	v156 = int32(-2)
	v158 = m.G0
	v160 = v158 - int32(32)
	m.G0 = v160
	switch int32(0) {
	case 0, 2:
		v170 = v156
		goto L43
	default:
		goto L44
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v124
	F_sigemptyset(m, v114+int32(16))
	mBase = m.M
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[9])) = v110
	v124 = int32(_a_F_PgArchiverMain_0)
	goto L30
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(268435456)
	v136 = v114 + int32(12)
	goto L37
L35:
	;
	m.G0 = v114 + int32(32)
	goto L29
L37:
	;
	goto L38
L38:
	;
	if v136 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v143 = int32(300)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10])) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v136)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[11])) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[12])) = v148
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L35
L42:
	;
	v202 = int32(-2)
	v204 = m.G0
	v206 = v204 - int32(32)
	m.G0 = v206
	switch int32(0) {
	case 0, 2:
		v216 = v202
		goto L56
	default:
		goto L57
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v170
	F_sigemptyset(m, v160+int32(16))
	mBase = m.M
	goto L46
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[13])) = v156
	v170 = int32(_a_F_PgArchiverMain_0)
	goto L43
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = int32(268435456)
	v182 = v160 + int32(12)
	goto L50
L48:
	;
	m.G0 = v160 + int32(32)
	goto L42
L50:
	;
	goto L51
L51:
	;
	if v182 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v189 = int32(280)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[14])) = v190
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[15])) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[16])) = v194
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L48
L55:
	;
	v248 = int32(917)
	v250 = m.G0
	v252 = v250 - int32(32)
	m.G0 = v252
	switch int32(919) {
	case 0, 2:
		v262 = v248
		goto L69
	default:
		goto L70
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v216
	F_sigemptyset(m, v206+int32(16))
	mBase = m.M
	goto L59
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[17])) = v202
	v216 = int32(_a_F_PgArchiverMain_0)
	goto L56
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+24)) = int32(268435456)
	v228 = v206 + int32(12)
	goto L63
L61:
	;
	m.G0 = v206 + int32(32)
	goto L55
L63:
	;
	goto L64
L64:
	;
	if v228 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v235 = int32(260)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[18])) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v228)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[19])) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[20])) = v240
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L61
L68:
	;
	v294 = int32(945)
	v296 = m.G0
	v298 = v296 - int32(32)
	m.G0 = v298
	switch int32(947) {
	case 0, 2:
		v308 = v294
		goto L82
	default:
		goto L83
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+12)) = v262
	F_sigemptyset(m, v252+int32(16))
	mBase = m.M
	goto L72
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[21])) = v248
	v262 = int32(_a_F_PgArchiverMain_0)
	goto L69
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(268435456)
	v274 = v252 + int32(12)
	goto L76
L74:
	;
	m.G0 = v252 + int32(32)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if v274 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v281 = int32(200)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[22])) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v274)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[23])) = v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[24])) = v286
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L74
L81:
	;
	v340 = int32(0)
	v342 = m.G0
	v344 = v342 - int32(32)
	m.G0 = v344
	switch int32(2) {
	case 0, 2:
		v354 = v340
		goto L95
	default:
		goto L96
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298)+12)) = v308
	F_sigemptyset(m, v298+int32(16))
	mBase = m.M
	goto L85
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[25])) = v294
	v308 = int32(_a_F_PgArchiverMain_0)
	goto L82
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298)+24)) = int32(268435456)
	v320 = v298 + int32(12)
	goto L89
L87:
	;
	m.G0 = v298 + int32(32)
	goto L81
L89:
	;
	goto L90
L90:
	;
	if v320 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v327 = int32(240)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[26])) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[27])) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[28])) = v332
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L87
L94:
	;
	F_sigprocmask(m, int32(_a_F_PgArchiverMain_1), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L107
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+12)) = v354
	F_sigemptyset(m, v344+int32(16))
	mBase = m.M
	goto L97
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[29])) = v340
	v354 = int32(_a_F_PgArchiverMain_0)
	goto L95
L97:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+24)) = int32(268435457)
	v366 = v344 + int32(12)
	goto L102
L100:
	;
	m.G0 = v344 + int32(32)
	goto L94
L102:
	;
	goto L103
L103:
	;
	if v366 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v373 = int32(340)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[30])) = v374
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v366)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[31])) = v376
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v366)))
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[32])) = v378
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L100
L107:
	;
	F_on_shmem_exit(m, int32(946), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[33]))
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v396
	v400 = F_palloc(m, int32(2888))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35])) = v400
	v403 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v400)+4)) = v403
	v408 = F_binaryheap_allocate(m, int32(64), int32(947), v403)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v408
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[36]))
	v420 = F_AllocSetContextCreateInternal(m, v415, int32(_a_F_PgArchiverMain_2), int32(0), int32(_a_F_PgArchiverMain_3), int32(_a_F_PgArchiverMain_4))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[37])) = v420
	v423 = m.G0
	v425 = v423 - int32(16)
	m.G0 = v425
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[38]))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428))))
	if v429 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v515 = m.G0
	v517 = v515 - int32(3392)
	m.G0 = v517
	goto L142
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L139
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L136
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L131
	}
L116:
	;
	v445 = m.T0[v443].(func(*base.Module) int32)(m)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L123
	}
L117:
	;
	v443 = int32(948)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[39]))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	if v435 != 0 {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v437 = int32(0)
	v439 = F_load_external_function(m, v428, int32(_a_F_PgArchiverMain_5), v437, v437)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v439 == int32(0) {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v443 = v439
	goto L116
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[40])) = v445
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v445)+8))
	if v448 == int32(0) {
		goto L113
	} else {
		goto L124
	}
L124:
	;
	v453 = F_palloc0(m, int32(4))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[41])) = v453
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[40]))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	if v458 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	m.T0[v458].(func(*base.Module, int32))(m, v453)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_before_shmem_exit(m, int32(949), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	m.G0 = v425 + int32(16)
	goto L112
L131:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_PgArchiverMain_6), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errdetail(m, int32(_a_F_PgArchiverMain_7), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(921), int32(_a_F_PgArchiverMain_9))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = int32(_a_F_PgArchiverMain_5)
	F_errmsg(m, int32(_a_F_PgArchiverMain_10), v425)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(936), int32(_a_F_PgArchiverMain_9))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errmsg(m, int32(_a_F_PgArchiverMain_11), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(942), int32(_a_F_PgArchiverMain_9))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	v531 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = int32(0)
	goto L145
L143:
	;
	m.G0 = v517 + int32(3392)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L441
	}
L144:
	;
	goto L143
L145:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[43]))
	F_ProcessPgArchInterrupts(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[44]))
	if v539 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+4)) = int32(0)
	goto L153
L148:
	;
	v542 = F_time(m)
	mBase = m.M
	v544 = *(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[45]))
	if v544 == int64(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[45])) = v542
	goto L147
L150:
	;
	goto L151
L151:
	;
	if base.Ui32(int32(59)) < base.Ui32(base.I32_wrap_i64(v542-v544)) {
		goto L144
	} else {
		goto L152
	}
L152:
	;
	goto L147
L153:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[33]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+4)) = int32(0)
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	if v572 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	if v535 != 0 {
		goto L144
	} else {
		goto L438
	}
L155:
	;
	goto L154
L156:
	;
	v1366 = int32(0)
	v1369 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[44]))
	if v1369 != 0 {
		goto L155
	} else {
		goto L359
	}
L157:
	;
	v1273 = v1264 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+4)) = v1273
	v1276 = v517 + int32(1296)
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1262+v1273<<(uint(int32(2))%32))+8))
	if (v1280^v1276)&int32(3) != 0 {
		goto L341
	} else {
		goto L342
	}
L158:
	;
	v1237 = int32(0)
	v1238 = v1234
	goto L334
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L1
	} else {
		goto L330
	}
L160:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v717 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v716)+8)) = uint8(v717)
	*(*int32)(unsafe.Add(mBase, uint32(v716))) = int32(0)
	goto L195
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v576)+4)) = int32(0)
	v706 = v576
	goto L160
L162:
	;
	goto L163
L163:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	if v581 <= int32(0) {
		v706 = v576
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v584 = v581
	v585 = v576
	goto L165
L165:
	;
	v596 = v584 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v585)+4)) = v596
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v585+v596<<(uint(int32(2))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v517)+164)) = int32(_a_F_PgArchiverMain_12)
	*(*int32)(unsafe.Add(mBase, uint32(v517)+160)) = v601
	v606 = v517 + int32(1344)
	v611 = F_pg_snprintf(m, v606, int32(1024), int32(_a_F_PgArchiverMain_13), v517+int32(160))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	v706 = v701
	goto L160
L167:
	;
	v617 = F___fstatat(m, int32(-100), v606, v517+int32(176), int32(0))
	mBase = m.M
	goto L168
L168:
	;
	if v617 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v621 = v517 + int32(1296)
	if (v601^v621)&int32(3) != 0 {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	goto L171
L171:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[46]))
	if v697 != int32(44) {
		goto L159
	} else {
		goto L193
	}
L172:
	;
	goto L156
L173:
	;
	goto L172
L174:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v675)
	if v675&int32(255) == int32(0) {
		goto L173
	} else {
		goto L189
	}
L175:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	v674 = v601
	v675 = v627
	v676 = v621
	goto L174
L176:
	;
	goto L177
L177:
	;
	if v601&int32(3) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v631 = v601
	v633 = v621
	goto L181
L179:
	;
	v645 = v601
	v647 = v621
	goto L180
L180:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	v652 = int32(-2139062144)
	if (int32(16843008)-v649|v649)&v652 != v652 {
		v674 = v645
		v675 = v649
		v676 = v647
		goto L174
	} else {
		goto L185
	}
L181:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631))))
	*(*uint8)(unsafe.Add(mBase, uint32(v633))) = uint8(v634)
	if v634 == int32(0) {
		goto L173
	} else {
		goto L183
	}
L182:
	;
	v645 = v641
	v647 = v639
	goto L180
L183:
	;
	v638 = int32(1)
	v639 = v633 + v638
	v641 = v631 + v638
	if v641&int32(3) != 0 {
		v631 = v641
		v633 = v639
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v657 = v645
	v658 = v649
	v659 = v647
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659))) = v658
	v661 = int32(4)
	v662 = v659 + v661
	v664 = v657 + v661
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	v669 = int32(-2139062144)
	if (int32(16843008)-v666|v666)&v669 == v669 {
		v657 = v664
		v658 = v666
		v659 = v662
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v674 = v664
	v675 = v666
	v676 = v662
	goto L174
L188:
	;
	goto L187
L189:
	;
	v683 = v674
	v685 = v676
	goto L190
L190:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+1)) = uint8(v686)
	v688 = int32(1)
	if v686 != 0 {
		v683 = v683 + v688
		v685 = v685 + v688
		goto L190
	} else {
		goto L192
	}
L191:
	;
	goto L173
L192:
	;
	goto L191
L193:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	if int32(0) < v702 {
		v584 = v702
		v585 = v701
		goto L165
	} else {
		goto L194
	}
L194:
	;
	goto L166
L195:
	;
	v722 = v517 + int32(2368)
	v726 = F_pg_snprintf(m, v722, int32(1024), int32(_a_F_PgArchiverMain_14), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v728 = F_AllocateDir(m, v722)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v730 = F_ReadDir(m, v728, v722)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	if v730 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v733 = v730
	goto L202
L200:
	;
	goto L201
L201:
	;
	F_FreeDir(m, v728)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L323
	}
L202:
	;
	v744 = v733 + int32(19)
	v745 = F_strlen(m, v744)
	mBase = m.M
	if base.Ui32(v745-int32(47)) < base.Ui32(int32(-25)) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	goto L201
L204:
	;
	v1182 = F_ReadDir(m, v728, v517+int32(2368))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L321
	}
L205:
	;
	v750 = int32(_a_F_PgArchiverMain_15)
	v754 = m.G0
	v756 = v754 - int32(32)
	v757 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v756)+24)) = v757
	*(*int64)(unsafe.Add(mBase, uint32(v756)+16)) = v757
	*(*int64)(unsafe.Add(mBase, uint32(v756)+8)) = v757
	*(*int64)(unsafe.Add(mBase, uint32(v756))) = v757
	v765 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PgArchiverMain[47])))
	if v765 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v835 = v745 - int32(6)
	if base.Ui32(v833) < base.Ui32(v835) {
		goto L204
	} else {
		goto L225
	}
L207:
	;
	v833 = int32(0)
	goto L206
L208:
	;
	goto L209
L209:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PgArchiverMain[48])))
	if v769 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v773 = v744
	goto L213
L211:
	;
	goto L212
L212:
	;
	v783 = v750
	v784 = v765
	goto L216
L213:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
	if v779 == v765 {
		v773 = v773 + int32(1)
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v833 = v773 - v744
	goto L206
L215:
	;
	goto L214
L216:
	;
	v791 = v756 + int32(base.Ui32(v784)>>(uint(int32(3))%32))&int32(28)
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)))
	v793 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v791))) = v792 | v793<<(uint(v784)%32)
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+1)))
	if v797 != 0 {
		v783 = v783 + v793
		v784 = v797
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if v800 == int32(0) {
		v823 = v744
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L217
L219:
	;
	v833 = v823 - v744
	goto L206
L220:
	;
	v804 = v744
	v805 = v800
	goto L221
L221:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v756+int32(base.Ui32(v805)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v813)>>(uint(v805)%32))&int32(1) == int32(0) {
		v823 = v804
		goto L219
	} else {
		goto L223
	}
L222:
	;
	v823 = v821
	goto L219
L223:
	;
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804)+1)))
	v821 = v804 + int32(1)
	if v819 != 0 {
		v804 = v821
		v805 = v819
		goto L221
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	v837 = v835 + v744
	v838 = int32(_a_F_PgArchiverMain_12)
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	v844 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PgArchiverMain[49])))
	if base.B2i32(v841 == int32(0))|base.B2i32(v841 != v844) != 0 {
		v862 = v841
		v863 = v844
		goto L227
	} else {
		goto L228
	}
L226:
	;
	if v862-v863 != 0 {
		goto L204
	} else {
		goto L233
	}
L227:
	;
	goto L226
L228:
	;
	v847 = v837
	v848 = v838
	goto L229
L229:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848)+1)))
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847)+1)))
	if v852 == int32(0) {
		v862 = v852
		v863 = v851
		goto L227
	} else {
		goto L231
	}
L230:
	;
	v862 = v852
	v863 = v851
	goto L227
L231:
	;
	v855 = int32(1)
	if v852 == v851 {
		v847 = v847 + v855
		v848 = v848 + v855
		goto L229
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	if v835 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	base.MemoryCopy(m, v517+int32(1344), v744, v835)
	goto L236
L235:
	;
	goto L236
L236:
	;
	v869 = v517 + int32(1344)
	v871 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v869+v835))) = uint8(v871)
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v874)))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v876 <= int32(63) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v883 = v874 + v876*int32(41) + int32(264)
	if (v869^v883)&int32(3) != 0 {
		goto L243
	} else {
		goto L244
	}
L238:
	;
	goto L239
L239:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v875)+20))
	v971 = v517 + int32(1344)
	v972 = int32(0)
	v973 = F_strlen(m, v969)
	mBase = m.M
	if v973 != int32(16) {
		v986 = v972
		goto L265
	} else {
		goto L266
	}
L240:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v874)))
	F_binaryheap_add_unordered(m, v958, v883)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L261
	}
L241:
	;
	goto L240
L242:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v938))) = uint8(v937)
	if v937&int32(255) == int32(0) {
		goto L241
	} else {
		goto L257
	}
L243:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869))))
	v936 = v869
	v937 = v889
	v938 = v883
	goto L242
L244:
	;
	goto L245
L245:
	;
	if v869&int32(3) != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v893 = v869
	v895 = v883
	goto L249
L247:
	;
	v907 = v869
	v909 = v883
	goto L248
L248:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v907)))
	v914 = int32(-2139062144)
	if (int32(16843008)-v911|v911)&v914 != v914 {
		v936 = v907
		v937 = v911
		v938 = v909
		goto L242
	} else {
		goto L253
	}
L249:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	*(*uint8)(unsafe.Add(mBase, uint32(v895))) = uint8(v896)
	if v896 == int32(0) {
		goto L241
	} else {
		goto L251
	}
L250:
	;
	v907 = v903
	v909 = v901
	goto L248
L251:
	;
	v900 = int32(1)
	v901 = v895 + v900
	v903 = v893 + v900
	if v903&int32(3) != 0 {
		v893 = v903
		v895 = v901
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v919 = v907
	v920 = v911
	v921 = v909
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v921))) = v920
	v923 = int32(4)
	v924 = v921 + v923
	v926 = v919 + v923
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v919)+4))
	v931 = int32(-2139062144)
	if (int32(16843008)-v928|v928)&v931 == v931 {
		v919 = v926
		v920 = v928
		v921 = v924
		goto L254
	} else {
		goto L256
	}
L255:
	;
	v936 = v926
	v937 = v928
	v938 = v924
	goto L242
L256:
	;
	goto L255
L257:
	;
	v945 = v936
	v947 = v938
	goto L258
L258:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v947)+1)) = uint8(v948)
	v950 = int32(1)
	if v948 != 0 {
		v945 = v945 + v950
		v947 = v947 + v950
		goto L258
	} else {
		goto L260
	}
L259:
	;
	goto L241
L260:
	;
	goto L259
L261:
	;
	v962 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	if v964 != int32(64) {
		goto L204
	} else {
		goto L262
	}
L262:
	;
	F_binaryheap_build(m, v963)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	goto L204
L264:
	;
	if v1005 <= int32(0) {
		goto L204
	} else {
		goto L280
	}
L265:
	;
	v987 = F_strlen(m, v971)
	mBase = m.M
	if v987 == int32(16) {
		goto L271
	} else {
		goto L272
	}
L266:
	;
	v977 = F_strspn(m, v969, int32(_a_F_PgArchiverMain_16))
	mBase = m.M
	if v977 != int32(8) {
		v986 = v972
		goto L265
	} else {
		goto L267
	}
L267:
	;
	v983 = F_strcmp(m, v969+int32(8), int32(_a_F_PgArchiverMain_17))
	mBase = m.M
	v986 = base.B2i32(v983 == int32(0))
	goto L265
L268:
	;
	v1004 = F_strcmp(m, v969, v971)
	mBase = m.M
	v1005 = v1004
	goto L264
L269:
	;
	if v986 != 0 {
		goto L277
	} else {
		goto L278
	}
L270:
	;
	v997 = F_strcmp(m, v517+int32(1352), int32(_a_F_PgArchiverMain_17))
	mBase = m.M
	if v986 == base.B2i32(v997 == int32(0)) {
		goto L268
	} else {
		goto L276
	}
L271:
	;
	v991 = F_strspn(m, v971, int32(_a_F_PgArchiverMain_16))
	mBase = m.M
	if v991 == int32(8) {
		goto L270
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	if v986 != 0 {
		goto L269
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	goto L268
L276:
	;
	goto L269
L277:
	;
	v1003 = int32(-1)
	goto L279
L278:
	;
	v1003 = int32(1)
	goto L279
L279:
	;
	v1005 = v1003
	goto L264
L280:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1009)))
	v1011 = F_binaryheap_remove_first(m, v1010)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	if (v971^v1011)&int32(3) != 0 {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1089)))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+4))
	if v1090 < v1091 {
		goto L304
	} else {
		goto L305
	}
L283:
	;
	goto L282
L284:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1067))) = uint8(v1066)
	if v1066&int32(255) == int32(0) {
		goto L283
	} else {
		goto L299
	}
L285:
	;
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v1065 = v971
	v1066 = v1018
	v1067 = v1011
	goto L284
L286:
	;
	goto L287
L287:
	;
	if v971&int32(3) != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1022 = v971
	v1024 = v1011
	goto L291
L289:
	;
	v1036 = v971
	v1038 = v1011
	goto L290
L290:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1036)))
	v1043 = int32(-2139062144)
	if (int32(16843008)-v1040|v1040)&v1043 != v1043 {
		v1065 = v1036
		v1066 = v1040
		v1067 = v1038
		goto L284
	} else {
		goto L295
	}
L291:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1022))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1024))) = uint8(v1025)
	if v1025 == int32(0) {
		goto L283
	} else {
		goto L293
	}
L292:
	;
	v1036 = v1032
	v1038 = v1030
	goto L290
L293:
	;
	v1029 = int32(1)
	v1030 = v1024 + v1029
	v1032 = v1022 + v1029
	if v1032&int32(3) != 0 {
		v1022 = v1032
		v1024 = v1030
		goto L291
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	v1048 = v1036
	v1049 = v1040
	v1050 = v1038
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1050))) = v1049
	v1052 = int32(4)
	v1053 = v1050 + v1052
	v1055 = v1048 + v1052
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	v1060 = int32(-2139062144)
	if (int32(16843008)-v1057|v1057)&v1060 == v1060 {
		v1048 = v1055
		v1049 = v1057
		v1050 = v1053
		goto L296
	} else {
		goto L298
	}
L297:
	;
	v1065 = v1055
	v1066 = v1057
	v1067 = v1053
	goto L284
L298:
	;
	goto L297
L299:
	;
	v1074 = v1065
	v1076 = v1067
	goto L300
L300:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1074)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+1)) = uint8(v1077)
	v1079 = int32(1)
	if v1077 != 0 {
		v1074 = v1074 + v1079
		v1076 = v1076 + v1079
		goto L300
	} else {
		goto L302
	}
L301:
	;
	goto L283
L302:
	;
	goto L301
L303:
	;
	goto L204
L304:
	;
	v1094 = v1089 + int32(20)
	v1095 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1094+v1090<<(uint(v1095)%32)))) = v1011
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1089)))
	*(*int32)(unsafe.Add(mBase, uint32(v1089))) = v1099 + int32(1)
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1094+v1099<<(uint(v1095)%32))))
	if v1099 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L305:
	;
	goto L306
L306:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L318
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1094+v1141<<(uint(int32(2))%32)))) = v1106
	goto L303
L308:
	;
	v1141 = int32(0)
	goto L307
L309:
	;
	goto L310
L310:
	;
	v1111 = v1099
	goto L311
L311:
	;
	v1123 = int32(2)
	v1124 = base.I32_div_s(v1111-int32(1), v1123)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1094+v1124<<(uint(v1123)%32))))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+16))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1089)+12))
	v1131 = m.T0[v1130].(func(*base.Module, int32, int32, int32) int32)(m, v1106, v1128, v1129)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L313
	}
L312:
	;
	v1141 = v1124
	goto L307
L313:
	;
	if v1131 <= int32(0) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1141 = v1111
	goto L307
L315:
	;
	goto L316
L316:
	;
	v1135 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1094+v1111<<(uint(v1135)%32)))) = v1128
	if base.Ui32(v1135) < base.Ui32(v1111) {
		v1111 = v1124
		goto L311
	} else {
		goto L317
	}
L317:
	;
	goto L312
L318:
	;
	F_errmsg_internal(m, int32(_a_F_PgArchiverMain_18), int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_19), int32(161), int32(_a_F_PgArchiverMain_20))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	if v1182 != 0 {
		v733 = v1182
		goto L202
	} else {
		goto L322
	}
L322:
	;
	goto L203
L323:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	if v1200 == int32(0) {
		goto L155
	} else {
		goto L324
	}
L324:
	;
	if int32(64) <= v1200 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1198)+4)) = v1200
	v1234 = v1198
	goto L158
L326:
	;
	goto L327
L327:
	;
	F_binaryheap_build(m, v1199)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1210)))
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+4)) = v1211
	if int32(0) < v1211 {
		v1234 = v1209
		goto L158
	} else {
		goto L329
	}
L329:
	;
	v1262 = v1209
	v1264 = v1211
	goto L157
L330:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+144)) = v517 + int32(1344)
	F_errmsg(m, int32(_a_F_PgArchiverMain_21), v517+int32(144))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(682), int32(_a_F_PgArchiverMain_22))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1238)))
	v1249 = F_binaryheap_remove_first(m, v1248)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L1
	} else {
		goto L336
	}
L335:
	;
	v1262 = v1252
	v1264 = v1259
	goto L157
L336:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v1252+v1237<<(uint(int32(2))%32))+8)) = v1249
	v1258 = v1237 + int32(1)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+4))
	if v1258 < v1259 {
		v1237 = v1258
		v1238 = v1252
		goto L334
	} else {
		goto L337
	}
L337:
	;
	goto L335
L338:
	;
	goto L156
L339:
	;
	goto L338
L340:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1335))) = uint8(v1334)
	if v1334&int32(255) == int32(0) {
		goto L339
	} else {
		goto L355
	}
L341:
	;
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280))))
	v1333 = v1280
	v1334 = v1286
	v1335 = v1276
	goto L340
L342:
	;
	goto L343
L343:
	;
	if v1280&int32(3) != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1290 = v1280
	v1292 = v1276
	goto L347
L345:
	;
	v1304 = v1280
	v1306 = v1276
	goto L346
L346:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1304)))
	v1311 = int32(-2139062144)
	if (int32(16843008)-v1308|v1308)&v1311 != v1311 {
		v1333 = v1304
		v1334 = v1308
		v1335 = v1306
		goto L340
	} else {
		goto L351
	}
L347:
	;
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1292))) = uint8(v1293)
	if v1293 == int32(0) {
		goto L339
	} else {
		goto L349
	}
L348:
	;
	v1304 = v1300
	v1306 = v1298
	goto L346
L349:
	;
	v1297 = int32(1)
	v1298 = v1292 + v1297
	v1300 = v1290 + v1297
	if v1300&int32(3) != 0 {
		v1290 = v1300
		v1292 = v1298
		goto L347
	} else {
		goto L350
	}
L350:
	;
	goto L348
L351:
	;
	v1316 = v1304
	v1317 = v1308
	v1318 = v1306
	goto L352
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1318))) = v1317
	v1320 = int32(4)
	v1321 = v1318 + v1320
	v1323 = v1316 + v1320
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+4))
	v1328 = int32(-2139062144)
	if (int32(16843008)-v1325|v1325)&v1328 == v1328 {
		v1316 = v1323
		v1317 = v1325
		v1318 = v1321
		goto L352
	} else {
		goto L354
	}
L353:
	;
	v1333 = v1323
	v1334 = v1325
	v1335 = v1321
	goto L340
L354:
	;
	goto L353
L355:
	;
	v1342 = v1333
	v1344 = v1335
	goto L356
L356:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1344)+1)) = uint8(v1345)
	v1347 = int32(1)
	if v1345 != 0 {
		v1342 = v1342 + v1347
		v1344 = v1344 + v1347
		goto L356
	} else {
		goto L358
	}
L357:
	;
	goto L339
L358:
	;
	goto L357
L359:
	;
	v1371 = v1366
	v1373 = v1366
	goto L360
L360:
	;
	v1381 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L362
	}
L361:
	;
	goto L155
L362:
	;
	if v1381 == int32(0) {
		goto L155
	} else {
		goto L363
	}
L363:
	;
	F_ProcessPgArchInterrupts(m)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v1388 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[50])) = v1388
	v1392 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[40]))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+4))
	if v1393 == v1388 {
		goto L369
	} else {
		goto L370
	}
L365:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[44]))
	if v1698 == int32(0) {
		v1371 = v1695
		v1373 = v1696
		goto L360
	} else {
		goto L437
	}
L366:
	;
	F_pg_usleep(m, int32(_a_F_PgArchiverMain_23))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L1
	} else {
		goto L436
	}
L367:
	;
	F_pg_usleep(m, int32(_a_F_PgArchiverMain_23))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L435
	}
L368:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), v1680, int32(_a_F_PgArchiverMain_24))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L1
	} else {
		goto L434
	}
L369:
	;
	v1427 = v517 + int32(1296)
	*(*int32)(unsafe.Add(mBase, uint32(v517)+112)) = v1427
	v1430 = v517 + int32(176)
	v1435 = F_pg_snprintf(m, v1430, int32(1024), int32(_a_F_PgArchiverMain_25), v517+int32(112))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L1
	} else {
		goto L379
	}
L370:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[41]))
	v1398 = m.T0[v1393].(func(*base.Module, int32) int32)(m, v1397)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	if v1398 != 0 {
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v1402 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	if v1402 == int32(0) {
		goto L155
	} else {
		goto L374
	}
L374:
	;
	F_errmsg(m, int32(_a_F_PgArchiverMain_26), int32(0))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[50]))
	if v1412 == int32(0) {
		v1680 = int32(431)
		goto L368
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+128)) = v1412
	F_errdetail_internal(m, int32(_a_F_PgArchiverMain_27), v517+int32(128))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(431), int32(_a_F_PgArchiverMain_24))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	goto L155
L379:
	;
	v1441 = F___fstatat(m, int32(-100), v1430, v517+int32(1200), int32(0))
	mBase = m.M
	goto L381
L380:
	;
	v1501 = v517 + int32(1296)
	v1502 = F_pgarch_archiveXlog(m, v1501)
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L1
	} else {
		goto L397
	}
L381:
	;
	if v1441 == int32(0) {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[46]))
	if v1445 != int32(44) {
		goto L380
	} else {
		goto L383
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+100)) = int32(_a_F_PgArchiverMain_12)
	*(*int32)(unsafe.Add(mBase, uint32(v517)+96)) = v1427
	v1452 = v517 + int32(2368)
	v1457 = F_pg_snprintf(m, v1452, int32(1024), int32(_a_F_PgArchiverMain_13), v517+int32(96))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	v1459 = F_unlink(m, v1452)
	mBase = m.M
	if v1459 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1464 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L1
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	if v1373 < int32(2) {
		goto L367
	} else {
		goto L392
	}
L388:
	;
	if v1464 == int32(0) {
		goto L153
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+64)) = v1452
	F_errmsg(m, int32(_a_F_PgArchiverMain_28), v517-int32(-64))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(454), int32(_a_F_PgArchiverMain_24))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	goto L153
L392:
	;
	v1483 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	if v1483 == int32(0) {
		goto L155
	} else {
		goto L394
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+80)) = v517 + int32(2368)
	F_errmsg(m, int32(_a_F_PgArchiverMain_29), v517+int32(80))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(464), int32(_a_F_PgArchiverMain_24))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	goto L155
L397:
	;
	if v1502 != 0 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+36)) = int32(_a_F_PgArchiverMain_12)
	*(*int32)(unsafe.Add(mBase, uint32(v517)+32)) = v1501
	v1508 = v517 + int32(2368)
	v1513 = F_pg_snprintf(m, v1508, int32(1024), int32(_a_F_PgArchiverMain_13), v517+int32(32))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L1
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v1606 = v517 + int32(1296)
	v1607 = int32(1)
	v1612 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[51]))
	v1613 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v1614 = int32(_a_F_PgArchiverMain_30)
	v1616 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52]))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52])) = v1616 + v1607
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+40)) = v1620 + v1607
	goto L421
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+20)) = int32(_a_F_PgArchiverMain_31)
	*(*int32)(unsafe.Add(mBase, uint32(v517)+16)) = v1501
	v1519 = v517 + int32(1344)
	v1524 = F_pg_snprintf(m, v1519, int32(1024), int32(_a_F_PgArchiverMain_13), v517+int32(16))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v1526 = F_rename(m, v1508, v1519)
	mBase = m.M
	if int32(0) <= v1526 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1548 = v517 + int32(1296)
	v1554 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[51]))
	v1555 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v1556 = int32(_a_F_PgArchiverMain_30)
	v1558 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52]))
	v1559 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52])) = v1558 + v1559
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+40)) = v1562 + v1559
	goto L412
L404:
	;
	v1531 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	if v1531 == int32(0) {
		goto L403
	} else {
		goto L406
	}
L406:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+4)) = v1519
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = v1508
	F_errmsg(m, int32(_a_F_PgArchiverMain_32), v517)
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(837), int32(_a_F_PgArchiverMain_33))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	goto L403
L410:
	;
	goto L153
L412:
	;
	goto L413
L413:
	;
	v1569 = v1554 + int32(48)
	v1570 = *(*int64)(unsafe.Add(mBase, uint32(v1569)))
	*(*int64)(unsafe.Add(mBase, uint32(v1569))) = v1570 + int64(1)
	goto L415
L415:
	;
	goto L416
L416:
	;
	v1577 = v1554 + int32(56)
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1577)+40)) = uint8(v1578)
	v1580 = *(*int64)(unsafe.Add(mBase, uint32(v1548)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1577)+32)) = v1580
	v1582 = *(*int64)(unsafe.Add(mBase, uint32(v1548)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1577)+24)) = v1582
	v1584 = *(*int64)(unsafe.Add(mBase, uint32(v1548)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1577)+16)) = v1584
	v1586 = *(*int64)(unsafe.Add(mBase, uint32(v1548)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1577)+8)) = v1586
	v1588 = *(*int64)(unsafe.Add(mBase, uint32(v1548)))
	*(*int64)(unsafe.Add(mBase, uint32(v1577))) = v1588
	goto L418
L418:
	;
	goto L419
L419:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1554+int32(104)))) = v1555
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+40))
	v1596 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1554)+40)) = v1595 + v1596
	v1599 = int32(_a_F_PgArchiverMain_30)
	v1601 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52]))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52])) = v1601 - v1596
	goto L410
L420:
	;
	if v1371 < int32(2) {
		goto L366
	} else {
		goto L430
	}
L421:
	;
	goto L423
L423:
	;
	v1627 = v1612 + int32(112)
	v1628 = *(*int64)(unsafe.Add(mBase, uint32(v1627)))
	*(*int64)(unsafe.Add(mBase, uint32(v1627))) = v1628 + int64(1)
	goto L424
L424:
	;
	goto L426
L426:
	;
	v1635 = v1612 + int32(120)
	v1636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1606)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1635)+40)) = uint8(v1636)
	v1638 = *(*int64)(unsafe.Add(mBase, uint32(v1606)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1635)+32)) = v1638
	v1640 = *(*int64)(unsafe.Add(mBase, uint32(v1606)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1635)+24)) = v1640
	v1642 = *(*int64)(unsafe.Add(mBase, uint32(v1606)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1635)+16)) = v1642
	v1644 = *(*int64)(unsafe.Add(mBase, uint32(v1606)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1635)+8)) = v1644
	v1646 = *(*int64)(unsafe.Add(mBase, uint32(v1606)))
	*(*int64)(unsafe.Add(mBase, uint32(v1635))) = v1646
	goto L427
L427:
	;
	goto L429
L429:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1612+int32(168)))) = v1613
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1612)+40))
	v1654 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1612)+40)) = v1653 + v1654
	v1657 = int32(_a_F_PgArchiverMain_30)
	v1659 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52]))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[52])) = v1659 - v1654
	goto L420
L430:
	;
	v1667 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	if v1667 == int32(0) {
		goto L155
	} else {
		goto L432
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v517)+48)) = v1606
	F_errmsg(m, int32(_a_F_PgArchiverMain_34), v517+int32(48))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	v1680 = int32(500)
	goto L368
L434:
	;
	goto L155
L435:
	;
	v1695 = v1371
	v1696 = v1373 + int32(1)
	goto L365
L436:
	;
	v1695 = v1371 + int32(1)
	v1696 = v1373
	goto L365
L437:
	;
	goto L361
L438:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[42]))
	v1717 = F_WaitLatch(m, v1713, int32(25), int32(_a_F_PgArchiverMain_35), int32(83886080))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	if v1717&int32(16) == int32(0) {
		goto L142
	} else {
		goto L440
	}
L440:
	;
	goto L144
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RemovePgTempFilesInDir(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	v7 = m.G0
	v9 = v7 - int32(2080)
	m.G0 = v9
	v11 = F_AllocateDir(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(2080)
	return
L2:
	;
	return
L3:
	;
	v13 = int32(0)
	if v11|base.B2i32(l1 == v13) == v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_RemovePgTempFilesInDir[0]))
	if v19 == int32(44) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v23 = F_ReadDirExtended(m, v11, l0, int32(15))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v23
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_FreeDir(m, v11)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L58
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)))
	if v31 != int32(46) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L11
L14:
	;
	v166 = F_ReadDirExtended(m, v11, l0, int32(15))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L56
	}
L15:
	;
	v44 = v26 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	v53 = F_pg_snprintf(m, v9+int32(32), int32(2048), int32(_a_F_RemovePgTempFilesInDir_0), v9+int32(16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L20
	}
L16:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v34 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v37 != int32(46) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+21)))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	if l2 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(32)
	F_errmsg(m, v152, v9)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L54
	}
L22:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L53
	}
L23:
	;
	v142 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L51
	}
L24:
	;
	v57 = int32(_a_F_RemovePgTempFilesInDir_1)
	goto L29
L25:
	;
	goto L26
L26:
	;
	v108 = F_get_dirent_type(m, v9+int32(32), v26, int32(0), int32(15))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L43
	}
L27:
	;
	if v95-v96 != 0 {
		goto L23
	} else {
		goto L40
	}
L29:
	;
	goto L30
L30:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v64 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v65 = v44
	v66 = v57
	v67 = int32(9)
	v68 = v64
	goto L35
L32:
	;
	v91 = v57
	v95 = int32(0)
	goto L33
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	goto L27
L34:
	;
	v91 = v86
	v95 = v88
	goto L33
L35:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if base.B2i32(v68 != v70)|base.B2i32(v70 == int32(0)) != 0 {
		v86 = v66
		v88 = v68
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v86 = v80
	v88 = int32(0)
	goto L34
L37:
	;
	v76 = v67 - int32(1)
	if v76 == int32(0) {
		v86 = v66
		v88 = v68
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v79 = int32(1)
	v80 = v66 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v81 != 0 {
		v65 = v65 + v79
		v66 = v80
		v67 = v76
		v68 = v81
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	goto L26
L41:
	;
	v129 = F_unlink(m, v9+int32(32))
	mBase = m.M
	if int32(0) <= v129 {
		goto L14
	} else {
		goto L48
	}
L42:
	;
	v111 = v9 + int32(32)
	F_RemovePgTempFilesInDir(m, v111, int32(0), int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	switch v108 {
	case 0:
		goto L14
	default:
		goto L41
	case 3:
		goto L42
	}
L44:
	;
	v116 = F_rmdir(m, v111)
	mBase = m.M
	if int32(0) <= v116 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v121 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	if v121 == int32(0) {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	v148 = int32(_a_F_RemovePgTempFilesInDir_2)
	v149 = int32(3441)
	goto L22
L48:
	;
	v134 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if v134 == int32(0) {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v148 = int32(_a_F_RemovePgTempFilesInDir_3)
	v149 = int32(3449)
	goto L22
L51:
	;
	if v142 == int32(0) {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v152 = int32(_a_F_RemovePgTempFilesInDir_4)
	v153 = int32(3455)
	goto L21
L53:
	;
	v152 = v148
	v153 = v149
	goto L21
L54:
	;
	F_errfinish(m, int32(_a_F_RemovePgTempFilesInDir_5), v153, int32(_a_F_RemovePgTempFilesInDir_6))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	goto L14
L56:
	;
	if v166 != 0 {
		v26 = v166
		goto L12
	} else {
		goto L57
	}
L57:
	;
	goto L13
L58:
	;
	goto L1
}
func F__PG_init_pg_trgm(m *base.Module) {
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	F_DefineCustomRealVariable(m, int32(_a_F__PG_init_pg_trgm_0), int32(_a_F__PG_init_pg_trgm_1), int32(_a_F__PG_init_pg_trgm_2), int32(_a_F__PG_init_pg_trgm_3), float64(0.30000001192092896), float64(0), float64(1), int32(6))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_DefineCustomRealVariable(m, int32(_a_F__PG_init_pg_trgm_4), int32(_a_F__PG_init_pg_trgm_5), int32(_a_F__PG_init_pg_trgm_2), int32(_a_F__PG_init_pg_trgm_6), float64(0.6000000238418579), float64(0), float64(1), int32(6))
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_DefineCustomRealVariable(m, int32(_a_F__PG_init_pg_trgm_7), int32(_a_F__PG_init_pg_trgm_8), int32(_a_F__PG_init_pg_trgm_2), int32(_a_F__PG_init_pg_trgm_9), float64(0.5), float64(0), float64(1), int32(6))
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_pg_trgm_10))
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__PG_init_pgcrypto(m *base.Module) {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	F_DefineCustomEnumVariable(m, int32(_a_F__PG_init_pgcrypto_0), int32(_a_F__PG_init_pgcrypto_1), int32(_a_F__PG_init_pgcrypto_2), int32(_a_F__PG_init_pgcrypto_3), int32(0), int32(_a_F__PG_init_pgcrypto_4), int32(5))
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_pgcrypto_5))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pg_base64_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	return int64(base.Ui64(base.I64_extend_i32_u(l1)*int64(3)) >> (uint(int64(2)) % 64))
}
func F_pg_big5_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v4 {
		v7 = int32(1)
	} else {
		v7 = int32(2)
	}
	return v7
}
func F_pg_blocking_pids(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	v2 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_palloc(m, int32(36))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v22
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v29
	v35 = F_palloc(m, v29*int32(20))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v41 = F_palloc(m, v38*int32(56))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v47 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v47
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v55 = F_LWLockAcquire(m, v51+int32(512), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v16 == int32(0) {
		v104 = v2
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v104 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[2]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 <= int32(0) {
		v104 = v2
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[3]))
	v69 = int32(0)
	goto L10
L10:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v60+int32(36)+v69<<(uint(int32(2))%32))))
	v90 = v68 + v87*int32(640)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+44))
	if v91 == v16 {
		v104 = v90
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v104 = int32(0)
	goto L7
L12:
	;
	v94 = v69 + int32(1)
	if v94 != v61 {
		v69 = v94
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v117 = F_LWLockAcquire(m, v113+int32(_a_F_pg_blocking_pids_0), int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v383+int32(512))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L60
	}
L17:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v124 = F_LWLockAcquire(m, v120+int32(_a_F_pg_blocking_pids_1), int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v131 = F_LWLockAcquire(m, v127+int32(_a_F_pg_blocking_pids_2), int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v138 = F_LWLockAcquire(m, v134+int32(_a_F_pg_blocking_pids_3), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v145 = F_LWLockAcquire(m, v141+int32(_a_F_pg_blocking_pids_4), int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v152 = F_LWLockAcquire(m, v148+int32(_a_F_pg_blocking_pids_5), int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v159 = F_LWLockAcquire(m, v155+int32(_a_F_pg_blocking_pids_6), int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v166 = F_LWLockAcquire(m, v162+int32(_a_F_pg_blocking_pids_7), int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v173 = F_LWLockAcquire(m, v169+int32(_a_F_pg_blocking_pids_8), int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v180 = F_LWLockAcquire(m, v176+int32(_a_F_pg_blocking_pids_9), int32(1))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v187 = F_LWLockAcquire(m, v183+int32(_a_F_pg_blocking_pids_10), int32(1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v194 = F_LWLockAcquire(m, v190+int32(_a_F_pg_blocking_pids_11), int32(1))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v201 = F_LWLockAcquire(m, v197+int32(_a_F_pg_blocking_pids_12), int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v208 = F_LWLockAcquire(m, v204+int32(_a_F_pg_blocking_pids_13), int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v215 = F_LWLockAcquire(m, v211+int32(_a_F_pg_blocking_pids_14), int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v222 = F_LWLockAcquire(m, v218+int32(_a_F_pg_blocking_pids_15), int32(1))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v104)+616))
	if v224 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v272+int32(_a_F_pg_blocking_pids_15))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	F_GetSingleProcBlockerStatusData(m, v104, v18)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+624))
	if v229 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	v233 = v224 + int32(620)
	if v229 == v233 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v235 = v229
	goto L40
L40:
	;
	F_GetSingleProcBlockerStatusData(m, v235-int32(628), v18)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L33
L42:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v254 != v233 {
		v235 = v254
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v278+int32(_a_F_pg_blocking_pids_14))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v284+int32(_a_F_pg_blocking_pids_13))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v290+int32(_a_F_pg_blocking_pids_12))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v296+int32(_a_F_pg_blocking_pids_11))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v302+int32(_a_F_pg_blocking_pids_10))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v308+int32(_a_F_pg_blocking_pids_9))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v314+int32(_a_F_pg_blocking_pids_8))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v320+int32(_a_F_pg_blocking_pids_7))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v326+int32(_a_F_pg_blocking_pids_6))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v332+int32(_a_F_pg_blocking_pids_5))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v338+int32(_a_F_pg_blocking_pids_4))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v344+int32(_a_F_pg_blocking_pids_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v350+int32(_a_F_pg_blocking_pids_2))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v356+int32(_a_F_pg_blocking_pids_1))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v362+int32(_a_F_pg_blocking_pids_0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L16
L60:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v391 = F_palloc(m, v388<<(uint(int32(2))%32))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v393 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v408 = v2
	v410 = v2
	goto L65
L63:
	;
	v684 = v2
	goto L64
L64:
	;
	v688 = F_construct_array_builtin(m, v391, v684, int32(23))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L113
	}
L65:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v415 = v412 + v410*int32(20)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v419 = v411 + v416*int32(56)
	v420 = int32(0)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	if v421 <= v420 {
		v518 = v420
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v684 = v665
	goto L64
L67:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+15)))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v535<<(uint(int32(2))%32))+uint32(_c_F_pg_blocking_pids[4])))
	goto L94
L68:
	;
	v425 = v421 & int32(3)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v427 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v421) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v434 = v420
	v437 = v427
	v440 = int32(0)
	goto L72
L70:
	;
	v477 = v420
	v480 = v427
	goto L71
L71:
	;
	v492 = v477
	v495 = v480
	v497 = v427
	goto L88
L72:
	;
	v449 = int32(56)
	v451 = v419 + v437*v449
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v451)+40))
	if v458 == v426 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v425 == int32(0) {
		v518 = v469
		goto L67
	} else {
		goto L87
	}
L74:
	;
	v460 = v451
	goto L76
L75:
	;
	v460 = v434
	goto L76
L76:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v451)+96))
	if v461 == v426 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v463 = v451 + v449
	goto L79
L78:
	;
	v463 = v460
	goto L79
L79:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v451)+152))
	if v464 == v426 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v466 = v451 + int32(112)
	goto L82
L81:
	;
	v466 = v463
	goto L82
L82:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v451)+208))
	if v467 == v426 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v469 = v451 + int32(168)
	goto L85
L84:
	;
	v469 = v466
	goto L85
L85:
	;
	v470 = int32(4)
	v471 = v437 + v470
	v473 = v440 + v470
	if v473 != v421&int32(2147483644) {
		v434 = v469
		v437 = v471
		v440 = v473
		goto L72
	} else {
		goto L86
	}
L86:
	;
	goto L73
L87:
	;
	v477 = v469
	v480 = v471
	goto L71
L88:
	;
	v509 = v419 + v495*int32(56)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+40))
	if v510 == v426 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v518 = v512
	goto L67
L90:
	;
	v512 = v509
	goto L92
L91:
	;
	v512 = v492
	goto L92
L92:
	;
	v513 = int32(1)
	v516 = v497 + v513
	if v516 != v425 {
		v492 = v512
		v495 = v495 + v513
		v497 = v516
		goto L88
	} else {
		goto L93
	}
L93:
	;
	goto L89
L94:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	if int32(0) < v539 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v542 = int32(2)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v545+v546<<(uint(v542)%32))))
	v556 = v539
	v559 = int32(0)
	v564 = v408
	goto L98
L96:
	;
	v665 = v408
	goto L97
L97:
	;
	v669 = v410 + int32(1)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v669 < v670 {
		v408 = v665
		v410 = v669
		goto L65
	} else {
		goto L112
	}
L98:
	;
	v569 = v419 + v559*int32(56)
	if v569 == v518 {
		v639 = v556
		v647 = v564
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v665 = v647
	goto L97
L100:
	;
	v651 = v559 + int32(1)
	if v651 < v639 {
		v556 = v639
		v559 = v651
		v564 = v647
		goto L98
	} else {
		goto L111
	}
L101:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v569)+44))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v518)+44))
	if v571 == v572 {
		v639 = v556
		v647 = v564
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v569)+16))
	if v574&v550 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391+v564<<(uint(int32(2))%32)))) = v571
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	v639 = v634
	v647 = v564 + int32(1)
	goto L100
L104:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v569)+20))
	v577 = int32(0)
	if base.B2i32(v576 == v577)|base.B2i32(int32(base.Ui32(v550)>>(uint(v576)%32))&int32(1) == v577) != 0 {
		v639 = v556
		v647 = v564
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v415)+16))
	if v585 <= int32(0) {
		v639 = v556
		v647 = v564
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v569)+40))
	v593 = int32(0)
	goto L107
L107:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v534+v533<<(uint(v542)%32)+v593<<(uint(int32(2))%32))))
	if v608 == v588 {
		goto L103
	} else {
		goto L109
	}
L108:
	;
	v639 = v556
	v647 = v564
	goto L100
L109:
	;
	v611 = v593 + int32(1)
	if v585 != v611 {
		v593 = v611
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	goto L99
L112:
	;
	goto L66
L113:
	;
	return v688
}
func F_pg_char_and_wchar_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v8 = l0
	v9 = l1
	v10 = l2
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v13 != v14 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return v13 - v14
L6:
	;
	goto L7
L7:
	;
	if v13 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v20 = int32(1)
	v25 = v10 - v20
	if v25 != 0 {
		v8 = v8 + v20
		v9 = v9 + int32(4)
		v10 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
}
func F_pg_char_to_encoding_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(-1)
	if l0 == int32(0) {
		v123 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v123
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		v123 = v13
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v19) {
		v123 = v13
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = l0
	v23 = v16
	v24 = v11
	goto L5
L5:
	;
	v31 = v23 & int32(255)
	goto L7
L6:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	v64 = int32(_a_F_pg_char_to_encoding_private_0)
	v65 = int32(_a_F_pg_char_to_encoding_private_1)
	goto L15
L7:
	;
	if base.B2i32(base.Ui32(v31-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v31|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v55 = v24
	goto L10
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v56 != 0 {
		v22 = v22 + int32(1)
		v23 = v56
		v24 = v55
		goto L5
	} else {
		goto L14
	}
L11:
	;
	v51 = v23 | int32(32)
	goto L13
L12:
	;
	v51 = v23
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v51)
	v55 = v24 + int32(1)
	goto L10
L14:
	;
	goto L6
L15:
	;
	v77 = v65 + (v64-v65)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78))))
	v80 = v63 - v79
	if v80 != 0 {
		v108 = v80
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v123 = v13
	goto L1
L17:
	;
	v112 = base.B2i32(v108 < int32(0))
	if v108 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v106 != 0 {
		v108 = v106
		goto L17
	} else {
		goto L26
	}
L20:
	;
	v106 = v104 - v105
	goto L19
L21:
	;
	v89 = v11
	v90 = v78
	goto L22
L22:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v104 = v94
	v105 = v93
	goto L20
L24:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v123 = v107
	goto L1
L27:
	;
	v113 = v77 - int32(8)
	goto L29
L28:
	;
	v113 = v64
	goto L29
L29:
	;
	if v108 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v116 = v65
	goto L32
L31:
	;
	v116 = v77 + int32(8)
	goto L32
L32:
	;
	if base.Ui32(v116) <= base.Ui32(v113) {
		v64 = v113
		v65 = v116
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
}
func F_pg_checksum_init(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	switch l1 - int32(1) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
		return int32(0)
	case 1:
		v12 = F_pg_cryptohash_create(m, int32(2))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
			v29 = v12
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	case 2:
		v18 = F_pg_cryptohash_create(m, int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
			v29 = v18
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	case 3:
		v22 = F_pg_cryptohash_create(m, int32(4))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
			v29 = v22
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	case 4:
		v26 = F_pg_cryptohash_create(m, int32(5))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26
			v29 = v26
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	default:
		v42 = int32(0)
		return v42
	}
}
func F_pg_checksum_parse_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	v6 = l0
	v7 = int32(_a_F_pg_checksum_parse_type_0)
	goto L2
L1:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v10 == v11 {
		v33 = v10
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v44 = int32(0)
	goto L1
L4:
	;
	v35 = int32(1)
	if v33 != 0 {
		v6 = v6 + v35
		v7 = v7 + v35
		goto L2
	} else {
		goto L13
	}
L5:
	;
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = v10 | int32(32)
	goto L8
L7:
	;
	v21 = v10
	goto L8
L8:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = v11 | int32(32)
	goto L11
L10:
	;
	v30 = v11
	goto L11
L11:
	;
	if v21 == v30 {
		v33 = v21
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v44 = v21 - v30
	goto L1
L13:
	;
	goto L3
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(1)
L15:
	;
	goto L16
L16:
	;
	v54 = l0
	v55 = int32(_a_F_pg_checksum_parse_type_1)
	goto L18
L17:
	;
	if v92 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v58 == v59 {
		v81 = v58
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v92 = int32(0)
	goto L17
L20:
	;
	v83 = int32(1)
	if v81 != 0 {
		v54 = v54 + v83
		v55 = v55 + v83
		goto L18
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = v58 | int32(32)
	goto L24
L23:
	;
	v69 = v58
	goto L24
L24:
	;
	if base.Ui32((v59-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = v59 | int32(32)
	goto L27
L26:
	;
	v78 = v59
	goto L27
L27:
	;
	if v69 == v78 {
		v81 = v69
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v92 = v69 - v78
	goto L17
L29:
	;
	goto L19
L30:
	;
	v95 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v95
	return v95
L31:
	;
	goto L32
L32:
	;
	v102 = l0
	v103 = int32(_a_F_pg_checksum_parse_type_2)
	goto L34
L33:
	;
	if v140 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v106 == v107 {
		v129 = v106
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v140 = int32(0)
	goto L33
L36:
	;
	v131 = int32(1)
	if v129 != 0 {
		v102 = v102 + v131
		v103 = v103 + v131
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v106-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = v106 | int32(32)
	goto L40
L39:
	;
	v117 = v106
	goto L40
L40:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v126 = v107 | int32(32)
	goto L43
L42:
	;
	v126 = v107
	goto L43
L43:
	;
	if v117 == v126 {
		v129 = v117
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v140 = v117 - v126
	goto L33
L45:
	;
	goto L35
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	return int32(1)
L47:
	;
	goto L48
L48:
	;
	v150 = l0
	v151 = int32(_a_F_pg_checksum_parse_type_3)
	goto L50
L49:
	;
	if v188 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v154 == v155 {
		v177 = v154
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v188 = int32(0)
	goto L49
L52:
	;
	v179 = int32(1)
	if v177 != 0 {
		v150 = v150 + v179
		v151 = v151 + v179
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if base.Ui32((v154-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v165 = v154 | int32(32)
	goto L56
L55:
	;
	v165 = v154
	goto L56
L56:
	;
	if base.Ui32((v155-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v174 = v155 | int32(32)
	goto L59
L58:
	;
	v174 = v155
	goto L59
L59:
	;
	if v165 == v174 {
		v177 = v165
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v188 = v165 - v174
	goto L49
L61:
	;
	goto L51
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	return int32(1)
L63:
	;
	goto L64
L64:
	;
	v198 = l0
	v199 = int32(_a_F_pg_checksum_parse_type_4)
	goto L66
L65:
	;
	if v236 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v202 == v203 {
		v225 = v202
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v236 = int32(0)
	goto L65
L68:
	;
	v227 = int32(1)
	if v225 != 0 {
		v198 = v198 + v227
		v199 = v199 + v227
		goto L66
	} else {
		goto L77
	}
L69:
	;
	if base.Ui32((v202-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v213 = v202 | int32(32)
	goto L72
L71:
	;
	v213 = v202
	goto L72
L72:
	;
	if base.Ui32((v203-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v222 = v203 | int32(32)
	goto L75
L74:
	;
	v222 = v203
	goto L75
L75:
	;
	if v213 == v222 {
		v225 = v213
		goto L68
	} else {
		goto L76
	}
L76:
	;
	v236 = v213 - v222
	goto L65
L77:
	;
	goto L67
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	return int32(1)
L79:
	;
	goto L80
L80:
	;
	v248 = l0
	v249 = int32(_a_F_pg_checksum_parse_type_5)
	goto L82
L81:
	;
	if v286 != 0 {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v252 == v253 {
		v275 = v252
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v286 = int32(0)
	goto L81
L84:
	;
	v277 = int32(1)
	if v275 != 0 {
		v248 = v248 + v277
		v249 = v249 + v277
		goto L82
	} else {
		goto L93
	}
L85:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v263 = v252 | int32(32)
	goto L88
L87:
	;
	v263 = v252
	goto L88
L88:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v272 = v253 | int32(32)
	goto L91
L90:
	;
	v272 = v253
	goto L91
L91:
	;
	if v263 == v272 {
		v275 = v263
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v286 = v263 - v272
	goto L81
L93:
	;
	goto L83
L94:
	;
	v287 = int32(0)
	goto L96
L95:
	;
	v287 = int32(5)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v287
	return base.B2i32(v286 == int32(0))
}
func F_pg_column_size(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L32
	}
L2:
	;
	switch v35 + int32(2) {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		v96 = v35
		goto L11
	}
L3:
	;
	v16 = F_get_fn_expr_argtype(m, v11, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v35 = v34
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	v20 = F_get_typlen(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v20 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v27 = F_MemoryContextAlloc(m, v25, int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v27
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v20
	v35 = v20
	goto L2
L11:
	;
	m.G0 = v8 + int32(16)
	return v96
L12:
	;
	v92 = F_strlen(m, v10)
	mBase = m.M
	v96 = v92 + int32(1)
	goto L11
L13:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v39 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v96 = v91
	goto L11
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v91 = int32(base.Ui32(v83) >> (uint(int32(2)) % 32))
	goto L14
L16:
	;
	v91 = int32(base.Ui32(v76) >> (uint(int32(1)) % 32))
	goto L14
L17:
	;
	v44 = v10
	goto L20
L18:
	;
	v67 = v39
	v69 = v10
	goto L19
L19:
	;
	if v67&int32(1) == int32(0) {
		goto L15
	} else {
		goto L31
	}
L20:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v47 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v67 = v64
	v69 = v63
	goto L19
L22:
	;
	if v47 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 == int32(1) {
		v44 = v63
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+6))
	v91 = v52 & int32(1073741823)
	goto L14
L26:
	;
	goto L27
L27:
	;
	if v47&int32(254) != int32(2) {
		v76 = int32(1)
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2))
	v61 = F_EOH_get_flat_size(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v91 = v61
	goto L14
L30:
	;
	goto L21
L31:
	;
	v76 = v67
	goto L16
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	F_errmsg_internal(m, int32(_a_F_pg_column_size_0), v8)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_pg_column_size_1), int32(_a_F_pg_column_size_2), int32(_a_F_pg_column_size_3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
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
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
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
	v16 = m.G0
	v18 = v16 - int32(1024)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(23)
	v23 = F_palloc(m, int32(184))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = F_pstrdup(m, int32(_a_F_pg_config_0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
	v29 = int32(_a_F_pg_config_1)
	goto L8
L5:
	;
	v152 = F_strlen(m, v18)
	mBase = m.M
	v159 = v152 + int32(1)
	goto L38
L6:
	;
	v146 = F_strlen(m, v135)
	mBase = m.M
	goto L5
L8:
	;
	goto L9
L9:
	;
	v36 = int32(1023)
	if (v18^v29)&int32(3) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v139)
	goto L6
L11:
	;
	v120 = v115
	v121 = v116
	v122 = v117
	goto L32
L12:
	;
	if v110 == int32(0) {
		v135 = v108
		v136 = v109
		goto L10
	} else {
		goto L31
	}
L13:
	;
	v108 = v29
	v109 = v18
	v110 = v36
	goto L12
L14:
	;
	goto L15
L15:
	;
	goto L18
L16:
	;
	goto L25
L18:
	;
	goto L19
L19:
	;
	goto L16
L25:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_config[0])))
	if base.B2i32(v79 == int32(0))|int32(0) != 0 {
		v108 = v29
		v109 = v18
		v110 = v36
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v86 = v29
	v87 = v18
	v88 = v36
	goto L27
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v94 = int32(-2139062144)
	if (int32(16843008)-v91|v91)&v94 != v94 {
		v115 = v86
		v116 = v87
		v117 = v88
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v108 = v102
	v109 = v100
	v110 = v104
	goto L12
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v91
	v99 = int32(4)
	v100 = v87 + v99
	v102 = v86 + v99
	v104 = v88 - v99
	if base.Ui32(int32(3)) < base.Ui32(v104) {
		v86 = v102
		v87 = v100
		v88 = v104
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v115 = v108
	v116 = v109
	v117 = v110
	goto L11
L32:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
	if v124 == int32(0) {
		v135 = v120
		v136 = v121
		goto L10
	} else {
		goto L34
	}
L33:
	;
	v135 = v131
	v136 = v129
	goto L10
L34:
	;
	v128 = int32(1)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if v133 != 0 {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if v171 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v161 = int32(0)
	if v159 == v161 {
		v171 = v161
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v171 = v166
	goto L37
L40:
	;
	v165 = v159 - int32(1)
	v166 = v18 + v165
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v167 != int32(47) {
		v159 = v165
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v172)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v174 = F_pstrdup(m, v18)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v174
	v178 = F_pstrdup(m, int32(_a_F_pg_config_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v178
	F_get_doc_path(m, v18)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v183 = F_pstrdup(m, v18)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v183
	v187 = F_pstrdup(m, int32(_a_F_pg_config_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v187
	F_get_doc_path(m, v18)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v192 = F_pstrdup(m, v18)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v192
	v196 = F_pstrdup(m, int32(_a_F_pg_config_4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v196
	F_make_relative_path(m, v18, int32(_a_F_pg_config_5), int32(_a_F_pg_config_1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v203 = F_pstrdup(m, v18)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v203
	v207 = F_pstrdup(m, int32(_a_F_pg_config_6))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v207
	F_make_relative_path(m, v18, int32(_a_F_pg_config_7), int32(_a_F_pg_config_1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v214 = F_pstrdup(m, v18)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v214
	v218 = F_pstrdup(m, int32(_a_F_pg_config_8))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v218
	F_make_relative_path(m, v18, int32(_a_F_pg_config_9), int32(_a_F_pg_config_1))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v225 = F_pstrdup(m, v18)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v225
	v229 = F_pstrdup(m, int32(_a_F_pg_config_10))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v229
	F_make_relative_path(m, v18, int32(_a_F_pg_config_11), int32(_a_F_pg_config_1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v236 = F_pstrdup(m, v18)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v236
	v240 = F_pstrdup(m, int32(_a_F_pg_config_12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v240
	F_get_pkglib_path(m, v18)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v245 = F_pstrdup(m, v18)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v245
	v249 = F_pstrdup(m, int32(_a_F_pg_config_13))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v249
	F_make_relative_path(m, v18, int32(_a_F_pg_config_14), int32(_a_F_pg_config_1))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v256 = F_pstrdup(m, v18)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v256
	v260 = F_pstrdup(m, int32(_a_F_pg_config_15))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v260
	F_make_relative_path(m, v18, int32(_a_F_pg_config_16), int32(_a_F_pg_config_1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v267 = F_pstrdup(m, v18)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v267
	v271 = F_pstrdup(m, int32(_a_F_pg_config_17))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v271
	F_get_share_path(m, v18)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v276 = F_pstrdup(m, v18)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v276
	v280 = F_pstrdup(m, int32(_a_F_pg_config_18))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v280
	F_get_etc_path(m, int32(_a_F_pg_config_1), v18)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v286 = F_pstrdup(m, v18)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v286
	v290 = F_pstrdup(m, int32(_a_F_pg_config_19))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v290
	F_get_pkglib_path(m, v18)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v295 = int32(_a_F_pg_config_20)
	v296 = int32(1024)
	v298 = F_pg_ascii_verifystr(m, v18, v296)
	mBase = m.M
	if v298 == v296 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v305 = F_pstrdup(m, v18)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L86
	}
L82:
	;
	goto L81
L83:
	;
	v300 = F_strlen(m, v295)
	mBase = m.M
	goto L82
L84:
	;
	goto L85
L85:
	;
	v303 = F_strlcpy(m, v18+v298, v295, v296-v298)
	mBase = m.M
	goto L82
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v305
	v309 = F_pstrdup(m, int32(_a_F_pg_config_21))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v309
	v313 = F_pstrdup(m, int32(_a_F_pg_config_22))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = v313
	v317 = F_pstrdup(m, int32(_a_F_pg_config_23))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v317
	v321 = F_pstrdup(m, int32(_a_F_pg_config_24))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v321
	v325 = F_pstrdup(m, int32(_a_F_pg_config_25))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+120)) = v325
	v329 = F_pstrdup(m, int32(_a_F_pg_config_26))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+124)) = v329
	v333 = F_pstrdup(m, int32(_a_F_pg_config_27))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v333
	v337 = F_pstrdup(m, int32(_a_F_pg_config_28))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v337
	v341 = F_pstrdup(m, int32(_a_F_pg_config_29))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = v341
	v345 = F_pstrdup(m, int32(_a_F_pg_config_30))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+140)) = v345
	v349 = F_pstrdup(m, int32(_a_F_pg_config_31))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v349
	v353 = F_pstrdup(m, int32(_a_F_pg_config_32))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v353
	v357 = F_pstrdup(m, int32(_a_F_pg_config_33))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v357
	v361 = F_pstrdup(m, int32(_a_F_pg_config_34))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v361
	v365 = F_pstrdup(m, int32(_a_F_pg_config_35))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v365
	v369 = F_pstrdup(m, int32(_a_F_pg_config_36))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = v369
	v373 = F_pstrdup(m, int32(_a_F_pg_config_37))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = v373
	v377 = F_pstrdup(m, int32(_a_F_pg_config_38))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v377
	v381 = F_pstrdup(m, int32(_a_F_pg_config_39))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v381
	v385 = F_pstrdup(m, int32(_a_F_pg_config_40))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v385
	m.G0 = v18 + int32(1024)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v391 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v393 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	m.G0 = v8 + int32(32)
	return int32(0)
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(0)
	v400 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v400)
	v404 = v23 + v393<<(uint(int32(3))%32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v406 = F_cstring_to_text(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L109
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v406
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	v410 = F_cstring_to_text(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v410
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v413, v414, v8+int32(16), v8+int32(14))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v422 = v393 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if base.Ui32(v422) < base.Ui32(v423) {
		v393 = v422
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
}
func F_pg_control_checkpoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	v15 = F_get_call_result_type(m, l0, int32(0), v10+int32(108))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[0]))
			v26 = F_LWLockAcquire(m, v22+int32(1152), int32(1))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[1]))
				v32 = F_get_controlfile(m, v29, v10+int32(31))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[0]))
					F_LWLockRelease(m, v35+int32(1152))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
						if v40 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_checkpoint_0), int32(0))
								mBase = m.M
								v190 = m.ExcPending
								if v190 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_checkpoint_1), int32(90), int32(_a_F_pg_control_checkpoint_2))
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v43 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v44
							v47 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[2])))
							v48 = base.I64_div_u_s(v43, v47)
							v50 = base.I64_div_u_s(int64(4294967296), v47)
							v51 = base.I64_div_u_s(v48, v50)
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v51)
							v54 = v48 - v50*v51
							*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v54)
							v57 = v10 + int32(32)
							v62 = F_pg_snprintf(m, v57, int32(64), int32(_a_F_pg_control_checkpoint_3), v10+int32(16))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v64 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
								v65 = F_Int64GetDatum(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+112)) = uint8(v67)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = v65
									v70 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
									v71 = F_Int64GetDatum(m, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										v73 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+113)) = uint8(v73)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v71
										v76 = F_cstring_to_text(m, v57)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v78 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+114)) = uint8(v78)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+152)) = v76
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+115)) = uint8(v78)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v81
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+116)) = uint8(v78)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v85
											v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+56)))
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+117)) = uint8(v78)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v89
											v93 = *(*int64)(unsafe.Add(mBase, uint32(v32)+64))
											*(*uint32)(unsafe.Add(mBase, uint32(v10)+4)) = uint32(v93)
											v96 = int64(base.Ui64(v93) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v10))) = uint32(v96)
											v99 = F_psprintf(m, int32(_a_F_pg_control_checkpoint_4), v10)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												v101 = F_cstring_to_text(m, v99)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+118)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+168)) = v101
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = v106
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+120)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = v110
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+121)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v114
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v32)+84))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+122)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v118
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+123)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+188)) = v122
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v32)+120))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+124)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = v126
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v32)+92))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+125)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v130
													v134 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+126)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v134
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+127)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v138
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v32)+116))
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+128)) = uint8(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = v142
													v146 = *(*int64)(unsafe.Add(mBase, uint32(v32)+104))
													v151 = F_Int64GetDatum(m, v146*int64(1000000)-int64(946684800000000))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														v153 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+129)) = uint8(v153)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+212)) = v151
														v156 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
														v161 = F_heap_form_tuple(m, v156, v10+int32(144), v10+int32(112))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return int32(0)
														} else {
															v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
															v164 = F_HeapTupleHeaderGetDatum(m, v163)
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return int32(0)
															} else {
																m.G0 = v10 + int32(224)
																return v164
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v173 = m.ExcPending
			if v173 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_checkpoint_5), int32(0))
				mBase = m.M
				v177 = m.ExcPending
				if v177 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_checkpoint_1), int32(82), int32(_a_F_pg_control_checkpoint_2))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
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
func F_pg_conversion_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_ConversionIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_crypt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_text_to_cstring(m, v9)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_text_to_cstring(m, v14)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v21 = F_palloc0(m, int32(128))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = F_px_crypt(m, v16, v18, v21, int32(128))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v16)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									if v24 != 0 {
										v30 = F_cstring_to_text(m, v24)
										mBase = m.M
										v31 = m.ExcPending
										if v31 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v21)
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return int32(0)
											} else {
												v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v34 != v9 {
													F_pfree(m, v9)
													mBase = m.M
													v37 = m.ExcPending
													if v37 != 0 {
														return int32(0)
													} else {
														v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v38 != v14 {
															F_pfree(m, v14)
															mBase = m.M
															v41 = m.ExcPending
															if v41 != 0 {
																return int32(0)
															} else {
																return v30
															}
														} else {
															return v30
														}
													}
												} else {
													v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v38 != v14 {
														F_pfree(m, v14)
														mBase = m.M
														v41 = m.ExcPending
														if v41 != 0 {
															return int32(0)
														} else {
															return v30
														}
													} else {
														return v30
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(579))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_pg_crypt_0), int32(0))
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_crypt_1), int32(236), int32(_a_F_pg_crypt_2))
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
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
					}
				}
			}
		}
	}
}
func F_pg_cryptohash_error(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		return int32(_a_F_pg_cryptohash_error_0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v9 == int32(1) {
			v12 = int32(_a_F_pg_cryptohash_error_1)
		} else {
			v12 = int32(_a_F_pg_cryptohash_error_2)
		}
		if v9 == int32(2) {
			v15 = int32(_a_F_pg_cryptohash_error_0)
		} else {
			v15 = v12
		}
		return v15
	}
}
func F_pg_cryptohash_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = m.Env.Pgmem_hash_reset(m, v6)
		mBase = m.M
		return int32(0)
	}
}
func F_pg_database_encoding_character_incrementer(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_database_encoding_character_incrementer[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 == int32(1) {
		v10 = int32(1638)
	} else {
		v10 = int32(1639)
	}
	if v7 == int32(6) {
		v13 = int32(1637)
	} else {
		v13 = v10
	}
	return v13
}
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4&int32(3) != 0 {
		v7 = F_detoast_attr(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(base.Ui32(v12) >> (uint(int32(2)) % 32))
		v15 = F_palloc(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				base.MemoryCopy(m, v15, l0, v14)
			} else {
			}
			return v15
		}
	}
}
func F_pg_euctw2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(142) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	return v87
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v71
	v76 = v18 + int32(1)
	v78 = v14 + int32(4)
	v79 = v15 + v72
	if int32(0) < v79 {
		v13 = v73
		v14 = v78
		v15 = v79
		v18 = v76
		goto L4
	} else {
		goto L18
	}
L8:
	;
	if v19 == int32(0) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L13
	}
L9:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v28 = v24<<(uint(int32(16))%32) | int32(-1912602624)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v33 = v30<<(uint(int32(8))%32) | v28
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v71 = v33 | v35
	v72 = int32(-4)
	v73 = v13 + int32(4)
	goto L7
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v46 = v42<<(uint(int32(8))%32) | int32(_a_F_pg_euctw2wchar_with_len_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v71 = v46 | v48
	v72 = int32(-3)
	v73 = v13 + int32(3)
	goto L7
L13:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v15 == int32(1) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v71 = v19
	v72 = int32(-1)
	v73 = v13 + int32(1)
	goto L7
L17:
	;
	v61 = v19 << (uint(int32(8)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v61
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v71 = v61 | v63
	v72 = int32(-2)
	v73 = v13 + int32(2)
	goto L7
L18:
	;
	v83 = v78
	v87 = v76
	goto L6
}
func F_pg_event_trigger_table_rewrite_reason(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_table_rewrite_reason[0]))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		if v9 != 0 {
			m.G0 = v5 + int32(16)
			return v9
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50463299))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_reason_0)
					F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_1), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_2), int32(1651), int32(_a_F_pg_event_trigger_table_rewrite_reason_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50463299))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_reason_0)
				F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_1), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_2), int32(1651), int32(_a_F_pg_event_trigger_table_rewrite_reason_3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
func F_pg_gb18030_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v2 < int32(0) {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if base.Ui32((v7-int32(48))&int32(255)) < base.Ui32(int32(10)) {
			v14 = int32(4)
		} else {
			v14 = int32(2)
		}
		v16 = v14
	} else {
		v16 = int32(1)
	}
	return v16
}
func F_pg_get_client_encoding(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_client_encoding[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return v3
}
func F_pg_get_expr_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v11 != 0 {
			v12 = int32(7)
		} else {
			v12 = int32(2)
		}
		v13 = F_pg_get_expr_worker(m, v4, v8, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				v18 = v13
			} else {
				v15 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
				v18 = int32(0)
			}
			return v18
		}
	}
}
func F_pg_get_line_append(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v17 = int32(-1)
	v18 = v3
	v19 = v3
	v20 = v3
	v22 = v3
	goto L4
L1:
	;
	m.G0 = v12 + int32(16)
	return v142
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v122
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v128+v122))) = uint8(v130)
	v142 = v130
	goto L1
L3:
	;
	v142 = int32(1)
	goto L1
L4:
	;
	if v17 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if int32(base.Ui32(v82)>>(uint(int32(5))%32))&int32(1) != 0 {
		v121 = v35
		v122 = v36
		goto L2
	} else {
		goto L43
	}
L6:
	;
	goto L11
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v35 = l1 + int32(4)
	v36 = v29
	v37 = int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v20 == int32(0) {
		v35 = v18
		v36 = v19
		v37 = v22
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, 8)) = uint8(v33)
	v121 = v18
	v122 = v19
	goto L2
L11:
	;
	goto L14
L12:
	;
	goto L5
L13:
	;
	v89 = int32(m.ExcTag)
	v90 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v89 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	if v37 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L31
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, 4))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(1)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v58 = F_fgets(m, v53+v54, v56-v53, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v37 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, 4))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(0)
	goto L22
L21:
	;
	goto L22
L22:
	;
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v69 = F_strlen(m, v66+v67)
	mBase = m.M
	v70 = v69 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v70
	if v36 < v70 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	goto L15
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v66-int32(1)))))
	if v76 == int32(10) {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_enlargeStringInfo(m, l1, int32(128))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L14
L31:
	;
	goto L12
L32:
	;
	v94 = int32(v90)
	m.G0 = v12
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v12+int32(12) == v100 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	m.ExcPending = 1
	goto L41
L34:
	;
	if v104 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v104 = v102
	goto L37
L36:
	;
	v104 = int32(0)
	goto L37
L37:
	;
	goto L34
L38:
	;
	F___wasm_longjmp(m, v97, v96)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v17 = v104
	v18 = v35
	v19 = v36
	v20 = v96
	v22 = v37
	goto L4
L41:
	;
	return int32(0)
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v112 == v36 {
		v121 = v35
		v122 = v36
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L3
}
func F_pg_get_multixact_members(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		if v13 == int32(0) {
			v16 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(_a_F_pg_get_multixact_members_0)
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_members[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_members[0])) = v23
				v26 = F_palloc(m, int32(12))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = F_GetMultiXactIdMembers(m, v11, v26, int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v29
						v37 = F_get_call_result_type(m, l0, v31, v9+int32(40))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_pg_get_multixact_members_1), int32(0))
									mBase = m.M
									v143 = m.ExcPending
									if v143 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_get_multixact_members_2), int32(3645), int32(_a_F_pg_get_multixact_members_3))
										mBase = m.M
										v148 = m.ExcPending
										if v148 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v41
								v43 = F_TupleDescGetAttInMetadata(m, v41)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v26
									*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v43
									*(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_members[0])) = v21
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									if v56 < v57 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v56<<(uint(int32(3))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v63
										v68 = F_psprintf(m, int32(_a_F_pg_get_multixact_members_4), v9+int32(32))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v68
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(3))%32))+4))
											if base.Ui32(int32(6)) <= base.Ui32(v76) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v152 = m.ExcPending
												if v152 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
													F_errmsg_internal(m, int32(_a_F_pg_get_multixact_members_5), v9+int32(16))
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_get_multixact_members_2), int32(1834), int32(_a_F_pg_get_multixact_members_6))
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(int32(2))%32))+uint32(_c_F_pg_get_multixact_members[1])))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v81
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
												v86 = F_BuildTupleFromCStrings(m, v83, v9+int32(40))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v88 + int32(1)
													v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
													F_pfree(m, v92)
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														v95 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
														*(*int64)(unsafe.Add(mBase, uint32(v54))) = v95 + int64(1)
														v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = int32(1)
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
														v103 = F_HeapTupleHeaderGetDatum(m, v102)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															v114 = v103
															m.G0 = v9 + int32(48)
															return v114
														}
													}
												}
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = int32(2)
											v110 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
											v114 = int32(0)
											m.G0 = v9 + int32(48)
											return v114
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
			if v56 < v57 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v56<<(uint(int32(3))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v63
				v68 = F_psprintf(m, int32(_a_F_pg_get_multixact_members_4), v9+int32(32))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v68
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(int32(3))%32))+4))
					if base.Ui32(int32(6)) <= base.Ui32(v76) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
							F_errmsg_internal(m, int32(_a_F_pg_get_multixact_members_5), v9+int32(16))
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_get_multixact_members_2), int32(1834), int32(_a_F_pg_get_multixact_members_6))
								mBase = m.M
								v163 = m.ExcPending
								if v163 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(int32(2))%32))+uint32(_c_F_pg_get_multixact_members[1])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v81
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
						v86 = F_BuildTupleFromCStrings(m, v83, v9+int32(40))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v88 + int32(1)
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
							F_pfree(m, v92)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								v95 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
								*(*int64)(unsafe.Add(mBase, uint32(v54))) = v95 + int64(1)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = int32(1)
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
								v103 = F_HeapTupleHeaderGetDatum(m, v102)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v114 = v103
									m.G0 = v9 + int32(48)
									return v114
								}
							}
						}
					}
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = int32(2)
					v110 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
					v114 = int32(0)
					m.G0 = v9 + int32(48)
					return v114
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v122 = m.ExcPending
		if v122 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v125 = m.ExcPending
			if v125 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
				F_errmsg(m, int32(_a_F_pg_get_multixact_members_7), v9)
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_get_multixact_members_2), int32(3628), int32(_a_F_pg_get_multixact_members_3))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
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
func F_pg_get_partkeydef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_get_partkeydef_worker(m, v3, int32(2), int32(0), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_get_ruledef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_get_ruledef_worker(m, v3, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = F_cstring_to_text(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v15
				}
			}
		}
	}
}
func F_pg_get_timezone_offset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_get_timezone_offset[0])))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v7 < int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	return int32(1)
L2:
	;
	v13 = int32(1)
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_pg_get_timezone_offset_0)+v13<<(uint(int32(4))%32))))
	if v6 == v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v24 = v13 + int32(1)
	if v7 != v24 {
		v13 = v24
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L1
}
func F_pg_get_triggerdef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_get_triggerdef_worker(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = F_cstring_to_text(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v15
				}
			}
		}
	}
}
func F_pg_get_triggerdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_get_triggerdef_worker(m, v3, base.B2i32(v4 != int32(0)))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_pg_get_viewdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
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
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	F_initStringInfo(m, v10+int32(-32))
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
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[0]))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L120
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L117
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L114
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+52)) = int64(81604378650)
	v33 = F_SPI_prepare(m, int32(_a_F_pg_get_viewdef_worker_0), int32(2), v10+int32(-12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l0
	v46 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), int32(_a_F_pg_get_viewdef_worker_1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_SPI_keepplan(m, v33)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[0])) = v33
	goto L9
L13:
	;
	v48 = int32(_a_F_pg_get_viewdef_worker_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+50)) = uint16(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v46
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[0]))
	v57 = F_SPI_execute_plan(m, v52, v10+int32(-12), v10+int32(-14))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v57 != int32(5) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[1]))
	if v62 != int64(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v351 = F_SPI_finish(m)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L109
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[2]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = int32(_a_F_pg_get_viewdef_worker_3)
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v71 < v73 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v115 = v10 + int32(-1)
	v116 = F_SPI_getbinval(m, v69, v67, v113, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L19:
	;
	v113 = v79 + int32(1)
	goto L18
L20:
	;
	v78 = v73
	v79 = v71
	goto L23
L21:
	;
	goto L22
L22:
	;
	v102 = F_SystemAttributeByName(m, v70)
	mBase = m.M
	if v102 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v85 = v67 + v78<<(uint(int32(4))%32) + v79*int32(100)
	v88 = F_namestrcmp(m, v85+int32(24), v70)
	mBase = m.M
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+111)))
	if v91 != int32(1) {
		goto L19
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v95 = v79 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v95 < v96 {
		v78 = v96
		v79 = v95
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	goto L24
L30:
	;
	v113 = int32(-9)
	goto L18
L31:
	;
	goto L32
L32:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v102)+74)))
	v113 = v106
	goto L18
L33:
	;
	v118 = int32(_a_F_pg_get_viewdef_worker_4)
	v119 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v119 < v121 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v162 = F_SPI_getbinval(m, v69, v67, v161, v115)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L49
	}
L35:
	;
	v161 = v127 + int32(1)
	goto L34
L36:
	;
	v126 = v121
	v127 = v119
	goto L39
L37:
	;
	goto L38
L38:
	;
	v150 = F_SystemAttributeByName(m, v118)
	mBase = m.M
	if v150 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v133 = v67 + v126<<(uint(int32(4))%32) + v127*int32(100)
	v136 = F_namestrcmp(m, v133+int32(24), v118)
	mBase = m.M
	if v136 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+111)))
	if v139 != int32(1) {
		goto L35
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v143 = v127 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v143 < v144 {
		v126 = v144
		v127 = v143
		goto L39
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	goto L40
L46:
	;
	v161 = int32(-9)
	goto L34
L47:
	;
	goto L48
L48:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+74)))
	v161 = v154
	goto L34
L49:
	;
	v164 = int32(_a_F_pg_get_viewdef_worker_5)
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v165 < v167 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v208 = F_SPI_getbinval(m, v69, v67, v207, v115)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L65
	}
L51:
	;
	v207 = v173 + int32(1)
	goto L50
L52:
	;
	v172 = v167
	v173 = v165
	goto L55
L53:
	;
	goto L54
L54:
	;
	v196 = F_SystemAttributeByName(m, v164)
	mBase = m.M
	if v196 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v179 = v67 + v172<<(uint(int32(4))%32) + v173*int32(100)
	v182 = F_namestrcmp(m, v179+int32(24), v164)
	mBase = m.M
	if v182 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+111)))
	if v185 != int32(1) {
		goto L51
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v189 = v173 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v189 < v190 {
		v172 = v190
		v173 = v189
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	goto L56
L62:
	;
	v207 = int32(-9)
	goto L50
L63:
	;
	goto L64
L64:
	;
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v196)+74)))
	v207 = v200
	goto L50
L65:
	;
	v210 = int32(_a_F_pg_get_viewdef_worker_6)
	v211 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v211 < v213 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v254 = F_SPI_getvalue(m, v69, v67, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L81
	}
L67:
	;
	v253 = v219 + int32(1)
	goto L66
L68:
	;
	v218 = v213
	v219 = v211
	goto L71
L69:
	;
	goto L70
L70:
	;
	v242 = F_SystemAttributeByName(m, v210)
	mBase = m.M
	if v242 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v225 = v67 + v218<<(uint(int32(4))%32) + v219*int32(100)
	v228 = F_namestrcmp(m, v225+int32(24), v210)
	mBase = m.M
	if v228 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+111)))
	if v231 != int32(1) {
		goto L67
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v235 = v219 + int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v235 < v236 {
		v218 = v236
		v219 = v235
		goto L71
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	goto L72
L78:
	;
	v253 = int32(-9)
	goto L66
L79:
	;
	goto L80
L80:
	;
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+74)))
	v253 = v246
	goto L66
L81:
	;
	v256 = int32(_a_F_pg_get_viewdef_worker_7)
	v257 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v257 < v259 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v300 = F_SPI_getvalue(m, v69, v67, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v299 = v265 + int32(1)
	goto L82
L84:
	;
	v264 = v259
	v265 = v257
	goto L87
L85:
	;
	goto L86
L86:
	;
	v288 = F_SystemAttributeByName(m, v256)
	mBase = m.M
	if v288 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v271 = v67 + v264<<(uint(int32(4))%32) + v265*int32(100)
	v274 = F_namestrcmp(m, v271+int32(24), v256)
	mBase = m.M
	if v274 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+111)))
	if v277 != int32(1) {
		goto L83
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v281 = v265 + int32(1)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v281 < v282 {
		v264 = v282
		v265 = v281
		goto L87
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L88
L94:
	;
	v299 = int32(-9)
	goto L82
L95:
	;
	goto L96
L96:
	;
	v292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v288)+74)))
	v299 = v292
	goto L82
L97:
	;
	v302 = F_stringToNode(m, v300)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v302 == int32(0) {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	if base.B2i32(v208 == int32(0))|(base.B2i32(v308 != int32(1))|base.B2i32(v116&int32(255) != int32(49))) != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v317 != int32(60) {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	if v320 != int32(62) {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+2)))
	if v323 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
	if v326 != int32(1) {
		goto L16
	} else {
		goto L104
	}
L104:
	;
	v330 = v10 + int32(-32)
	v333 = F_table_open(m, v162, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	F_get_query_def(m, v325, v330, int32(0), v335, int32(1), l1, l2, int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_appendStringInfoChar(m, v330, int32(59))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_relation_close(m, v333, int32(1))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L16
L109:
	;
	if v351 != int32(2) {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	m.G0 = v12 - int32(-64)
	if v355 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v361 = v356
	goto L113
L112:
	;
	v361 = int32(0)
	goto L113
L113:
	;
	return v361
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_get_viewdef_worker_0)
	F_errmsg_internal(m, int32(_a_F_pg_get_viewdef_worker_8), v12)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_pg_get_viewdef_worker_9), int32(822), int32(_a_F_pg_get_viewdef_worker_10))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_viewdef_worker_11), v10+int32(-48))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_pg_get_viewdef_worker_9), int32(836), int32(_a_F_pg_get_viewdef_worker_10))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_viewdef_worker_12), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_pg_get_viewdef_worker_9), int32(858), int32(_a_F_pg_get_viewdef_worker_10))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_viewdef_wrap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_get_viewdef_worker(m, v3, int32(7), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_get_wait_events(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+26)) = uint8(v24)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+24)) = uint16(v24)
	v33 = v18 * int32(12)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_pg_get_wait_events[0])))
	v35 = F_cstring_to_text(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v61 = F_GetWaitEventCustomNames(m, int32(117440512), v7+int32(-4))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_pg_get_wait_events[1])))
	v39 = F_cstring_to_text(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v39
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_pg_get_wait_events[2])))
	v43 = F_cstring_to_text(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v46, v47, v7+int32(-24), v7+int32(-40))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v55 = v18 + int32(1)
	if v55 != int32(273) {
		v18 = v55
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if int32(0) < v63 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v67 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v128 = F_GetWaitEventCustomNames(m, int32(184549376), v7+int32(-4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)) = uint16(v73)
	v82 = F_cstring_to_text(m, int32(_a_F_pg_get_wait_events_0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v82
	v87 = v61 + v67<<(uint(int32(2))%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v89 = F_cstring_to_text(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v89
	v93 = v7 + int32(-24)
	F_initStringInfo(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v96
	F_appendStringInfo(m, v93, int32(_a_F_pg_get_wait_events_1), v7+int32(-48))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	v104 = F_cstring_to_text(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v104
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v107, v108, v7+int32(-40), v7+int32(-44))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v116 = v67 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if v116 < v117 {
		v67 = v116
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if int32(0) < v130 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v134 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	m.G0 = v9 - int32(-64)
	return int32(0)
L27:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)) = uint8(v140)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)) = uint16(v140)
	v149 = F_cstring_to_text(m, int32(_a_F_pg_get_wait_events_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v149
	v154 = v128 + v134<<(uint(int32(2))%32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = F_cstring_to_text(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v156
	v160 = v7 + int32(-24)
	F_initStringInfo(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v163
	F_appendStringInfo(m, v160, int32(_a_F_pg_get_wait_events_3), v9)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	v169 = F_cstring_to_text(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v169
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v172, v173, v7+int32(-40), v7+int32(-44))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v181 = v134 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if v181 < v182 {
		v134 = v181
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L28
}
func F_pg_hmac_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v5 = F_palloc(m, int32(280))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = int32(0)
			base.MemoryFill(m, v5, v9, int32(280))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v9
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
			if base.Ui32(l0) <= base.Ui32(int32(5)) {
				v18 = l0 << (uint(int32(2)) % 32)
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_hmac_create[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_hmac_create[1])))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v21
			} else {
			}
			v24 = F_pg_cryptohash_create(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24
				if v24 != 0 {
					v36 = v5
					return v36
				} else {
					F___memset(m, v5, int32(0), int32(280))
					mBase = m.M
					F_pfree(m, v5)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v36 = int32(0)
						return v36
					}
				}
			}
		} else {
			v36 = int32(0)
			return v36
		}
	}
}
func F_pg_hmac_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	if l0 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v7 = F_palloc(m, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
				return int32(-1)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != 0 {
					base.MemoryFill(m, v7, int32(0), v17)
				} else {
				}
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v22 = F_pg_cryptohash_final(m, v20, v7, v21)
				mBase = m.M
				if v22 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v55 == int32(0) {
						v70 = int32(_a_F_pg_hmac_final_0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
						if v62 == int32(1) {
							v65 = int32(_a_F_pg_hmac_final_1)
						} else {
							v65 = int32(_a_F_pg_hmac_final_2)
						}
						if v62 == int32(2) {
							v68 = int32(_a_F_pg_hmac_final_0)
						} else {
							v68 = v65
						}
						v70 = v68
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
					F_pfree(m, v7)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						return int32(-1)
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = F_pg_cryptohash_init(m, v25)
					mBase = m.M
					if v26 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v55 == int32(0) {
							v70 = int32(_a_F_pg_hmac_final_0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
							if v62 == int32(1) {
								v65 = int32(_a_F_pg_hmac_final_1)
							} else {
								v65 = int32(_a_F_pg_hmac_final_2)
							}
							if v62 == int32(2) {
								v68 = int32(_a_F_pg_hmac_final_0)
							} else {
								v68 = v65
							}
							v70 = v68
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
						F_pfree(m, v7)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							return int32(-1)
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v33 = F_pg_cryptohash_update(m, v29, l0+int32(152), v32)
						mBase = m.M
						if v33 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v55 == int32(0) {
								v70 = int32(_a_F_pg_hmac_final_0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
								if v62 == int32(1) {
									v65 = int32(_a_F_pg_hmac_final_1)
								} else {
									v65 = int32(_a_F_pg_hmac_final_2)
								}
								if v62 == int32(2) {
									v68 = int32(_a_F_pg_hmac_final_0)
								} else {
									v68 = v65
								}
								v70 = v68
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
							F_pfree(m, v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								return int32(-1)
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v38 = F_pg_cryptohash_update(m, v36, v7, v37)
							mBase = m.M
							if v38 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v55 == int32(0) {
									v70 = int32(_a_F_pg_hmac_final_0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									if v62 == int32(1) {
										v65 = int32(_a_F_pg_hmac_final_1)
									} else {
										v65 = int32(_a_F_pg_hmac_final_2)
									}
									if v62 == int32(2) {
										v68 = int32(_a_F_pg_hmac_final_0)
									} else {
										v68 = v65
									}
									v70 = v68
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
								F_pfree(m, v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									return int32(-1)
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v42 = F_pg_cryptohash_final(m, v41, l1, l2)
								mBase = m.M
								if int32(0) <= v42 {
									F_pfree(m, v7)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v51 = int32(0)
										return v51
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v55 == int32(0) {
										v70 = int32(_a_F_pg_hmac_final_0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
										if v62 == int32(1) {
											v65 = int32(_a_F_pg_hmac_final_1)
										} else {
											v65 = int32(_a_F_pg_hmac_final_2)
										}
										if v62 == int32(2) {
											v68 = int32(_a_F_pg_hmac_final_0)
										} else {
											v68 = v65
										}
										v70 = v68
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
									F_pfree(m, v7)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										return int32(-1)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v51 = int32(-1)
		return v51
	}
}
func F_pg_hmac_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_pg_cryptohash_free(m, v2)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			F___memset(m, l0, int32(0), int32(280))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_pg_identify_object_as_address(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v9
	v18 = F_get_call_result_type(m, l0, int32(0), v7+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(1) {
			v25 = v7 + int32(36)
			v27 = F_getObjectTypeDescription(m, v25, int32(1))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = F_cstring_to_text(m, v27)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v31)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v29
					v39 = F_getObjectIdentityParts(m, v25, v7+int32(32), v7+int32(28), int32(1))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						if v39 == int32(0) {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v43)
							v68 = v43
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v68)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
							v75 = F_heap_form_tuple(m, v70, v7+int32(16), v7+int32(13))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
								v78 = F_HeapTupleHeaderGetDatum(m, v77)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v78
								}
							}
						} else {
							F_pfree(m, v39)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
								if v48 != 0 {
									v49 = F_strlist_to_textarray(m, v48)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										v54 = v49
										v55 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v55)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v54
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										if v58 != 0 {
											v59 = F_strlist_to_textarray(m, v58)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v64 = v59
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v64
												v68 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v68)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v75 = F_heap_form_tuple(m, v70, v7+int32(16), v7+int32(13))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
													v78 = F_HeapTupleHeaderGetDatum(m, v77)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v78
													}
												}
											}
										} else {
											v62 = F_construct_empty_array(m, int32(25))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v64 = v62
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v64
												v68 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v68)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v75 = F_heap_form_tuple(m, v70, v7+int32(16), v7+int32(13))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
													v78 = F_HeapTupleHeaderGetDatum(m, v77)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v78
													}
												}
											}
										}
									}
								} else {
									v52 = F_construct_empty_array(m, int32(25))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v54 = v52
										v55 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v55)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v54
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										if v58 != 0 {
											v59 = F_strlist_to_textarray(m, v58)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v64 = v59
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v64
												v68 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v68)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v75 = F_heap_form_tuple(m, v70, v7+int32(16), v7+int32(13))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
													v78 = F_HeapTupleHeaderGetDatum(m, v77)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v78
													}
												}
											}
										} else {
											v62 = F_construct_empty_array(m, int32(25))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v64 = v62
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v64
												v68 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v68)
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v75 = F_heap_form_tuple(m, v70, v7+int32(16), v7+int32(13))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
													v78 = F_HeapTupleHeaderGetDatum(m, v77)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v78
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_identify_object_as_address_0), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_identify_object_as_address_1), int32(_a_F_pg_identify_object_as_address_2), int32(_a_F_pg_identify_object_as_address_3))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
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
func F_pg_import_system_collations(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v455 int32
	_ = v455
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(352)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_superuser(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v16 + int32(352)
	return v455
L2:
	;
	v433 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L6
	} else {
		goto L114
	}
L3:
	;
	v416 = F_ClosePipeStream(m, v36)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L6
	} else {
		goto L113
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L109
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L105
	}
L6:
	;
	return int32(0)
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = int32(0)
	v27 = F_SearchSysCacheExists(m, int32(38), v18, v24, v24, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L6
	} else {
		goto L101
	}
L11:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v32 = F_palloc(m, int32(1200))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v36 = F_OpenPipeStream(m, int32(_a_F_pg_import_system_collations_0), int32(_a_F_pg_import_system_collations_1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v36 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v43 = F_fgets(m, v16+int32(224), int32(128), v36)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	if v43 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v51 = v2
	v53 = v32
	v54 = int32(100)
	v57 = v2
	v59 = v2
	goto L18
L18:
	;
	v63 = F_strlen(m, v16+int32(224))
	mBase = m.M
	if v63 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v300 = F_ClosePipeStream(m, v36)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L83
	}
L20:
	;
	v298 = F_fgets(m, v16+int32(224), int32(128), v36)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L81
	}
L21:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v90)
	v93 = v16 + int32(224)
	v95 = v93
	goto L31
L22:
	;
	v66 = v63 + v16 + int32(223)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 == int32(10) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v73 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	if v73 == int32(0) {
		v285 = v51
		v287 = v53
		v288 = v54
		v291 = v57
		v293 = v59
		goto L20
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(224)
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_2), v16+int32(16))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(885), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v285 = v51
	v287 = v53
	v288 = v54
	v291 = v57
	v293 = v59
	goto L20
L30:
	;
	if base.B2i32(v97 == int32(0)) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v95))))
	if int32(0) < v97 {
		v95 = v95 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	goto L32
L34:
	;
	v108 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v124 = v16 + int32(224)
	v126 = F_pg_get_encoding_from_locale(m, v124, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L41
	}
L37:
	;
	if v108 == int32(0) {
		v285 = v51
		v287 = v53
		v288 = v54
		v291 = v57
		v293 = v59
		goto L20
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v93
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_5), v16-int32(-64))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(710), int32(_a_F_pg_import_system_collations_6))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v285 = v51
	v287 = v53
	v288 = v54
	v291 = v57
	v293 = v59
	goto L20
L41:
	;
	if v126 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v132 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(35)) <= base.Ui32(v126) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if v132 == int32(0) {
		v285 = v51
		v287 = v53
		v288 = v54
		v291 = v57
		v293 = v59
		goto L20
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v124
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_7), v16+int32(32))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(717), int32(_a_F_pg_import_system_collations_6))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v285 = v51
	v287 = v53
	v288 = v54
	v291 = v57
	v293 = v59
	goto L20
L49:
	;
	v151 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v126 == int32(0) {
		v285 = v51
		v287 = v53
		v288 = v54
		v291 = v57
		v293 = v59
		goto L20
	} else {
		goto L56
	}
L52:
	;
	if v151 == int32(0) {
		v285 = v51
		v287 = v53
		v288 = v54
		v291 = v57
		v293 = v59
		goto L20
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v16 + int32(224)
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_8), v16+int32(48))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(722), int32(_a_F_pg_import_system_collations_6))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v285 = v51
	v287 = v53
	v288 = v54
	v291 = v57
	v293 = v59
	goto L20
L56:
	;
	v171 = v16 + int32(224)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_pg_import_system_collations[0]))
	v174 = int32(99)
	v176 = int32(0)
	v179 = F_get_collation_actual_version(m, v174, v171)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v181 = int32(1)
	v183 = F_CollationCreate(m, v171, v18, v173, v174, int32(1), v126, v171, v171, v176, v176, v179, v181, v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v183 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	v189 = v57
	goto L61
L61:
	;
	v191 = v59 + int32(1)
	v197 = v16 + int32(224)
	v199 = int32(0)
	v207 = v16 + int32(96)
	goto L63
L62:
	;
	v189 = v57 + int32(1)
	goto L61
L63:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v210 != int32(46) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v253)
	if v199 == v253 {
		v285 = v51
		v287 = v53
		v288 = v54
		v291 = v189
		v293 = v191
		goto L20
	} else {
		goto L74
	}
L65:
	;
	if v210 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v218 = v197
	goto L71
L67:
	;
	goto L64
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v210)
	v214 = int32(1)
	v197 = v197 + v214
	v207 = v207 + v214
	goto L63
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v232 = v218 + int32(1)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	v238 = int32(255)
	if base.B2i32(base.Ui32((v233&int32(223)-int32(65))&v238) < base.Ui32(int32(26)))|(base.B2i32(v233 == int32(45))|base.B2i32(base.Ui32((v233-int32(48))&v238) < base.Ui32(int32(10)))) != 0 {
		v218 = v232
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v197 = v232
	v199 = int32(1)
	goto L63
L73:
	;
	goto L72
L74:
	;
	if v54 <= v51 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v260 = F_repalloc(m, v53, v54*int32(24))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L78
	}
L76:
	;
	v264 = v53
	v265 = v54
	goto L77
L77:
	;
	v268 = v264 + v51*int32(12)
	v271 = F_pstrdup(m, v16+int32(224))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L79
	}
L78:
	;
	v264 = v260
	v265 = v54 << (uint(int32(1)) % 32)
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v271
	v276 = F_pstrdup(m, v16+int32(96))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+8)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v276
	v285 = v51 + int32(1)
	v287 = v264
	v288 = v265
	v291 = v189
	v293 = v191
	goto L20
L81:
	;
	if v298 != 0 {
		v51 = v285
		v53 = v287
		v54 = v288
		v57 = v291
		v59 = v293
		goto L18
	} else {
		goto L82
	}
L82:
	;
	goto L19
L83:
	;
	if int32(2) <= v285 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v293 != 0 {
		v455 = v362
		goto L1
	} else {
		goto L100
	}
L85:
	;
	v316 = int32(0)
	v322 = v291
	goto L91
L86:
	;
	F_pg_qsort(m, v287, v285, int32(12), int32(516))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v308 = int32(1)
	if v285 != v308 {
		v362 = v291
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v311 = v285
	goto L85
L90:
	;
	v311 = v308
	goto L85
L91:
	;
	v328 = v287 + v316*int32(12)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_pg_import_system_collations[0]))
	v333 = int32(99)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v336 = int32(0)
	v339 = F_get_collation_actual_version(m, v333, v335)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L93
	}
L92:
	;
	v362 = v349
	goto L84
L93:
	;
	v341 = int32(1)
	v343 = F_CollationCreate(m, v330, v18, v332, v333, int32(1), v329, v335, v335, v336, v336, v339, v341, v341)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	if v343 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L6
	} else {
		goto L98
	}
L96:
	;
	v349 = v322
	goto L97
L97:
	;
	v351 = v316 + int32(1)
	if v351 != v311 {
		v316 = v351
		v322 = v349
		goto L91
	} else {
		goto L99
	}
L98:
	;
	v349 = v322 + int32(1)
	goto L97
L99:
	;
	goto L92
L100:
	;
	v427 = v362
	goto L2
L101:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_pg_import_system_collations_9), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(844), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v18
	F_errmsg(m, int32(_a_F_pg_import_system_collations_10), v16+int32(80))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(849), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_pg_import_system_collations_0)
	F_errmsg(m, int32(_a_F_pg_import_system_collations_11), v16)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(873), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v427 = v2
	goto L2
L114:
	;
	if v433 == int32(0) {
		v455 = v427
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_pg_import_system_collations_12), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(964), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v455 = v427
	goto L1
}
func F_pg_index_column_has_property(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v5 <= int32(0) {
				v16 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
				return int32(0)
			} else {
				v21 = F_indexam_property(m, l0, v12, int32(0), v6, v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v21
				}
			}
		}
	}
}
func F_pg_is_in_recovery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_is_in_recovery[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_is_in_recovery[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_is_in_recovery[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	return v14
}
func F_pg_largeobject_aclcheck_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pg_largeobject_aclmask_snapshot(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 == int64(0))
	}
}
func F_pg_last_committed_xact(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pg_last_committed_xact[0]))
	v15 = F_LWLockAcquire(m, v11+int32(_a_F_pg_last_committed_xact_0), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_last_committed_xact[1]))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
		if v21 != 0 {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_last_committed_xact[0]))
			F_LWLockRelease(m, v26+int32(_a_F_pg_last_committed_xact_0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v34 = F_get_call_result_type(m, l0, int32(0), v8+int32(12))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					if v34 != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_pg_last_committed_xact_1), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_last_committed_xact_2), int32(434), int32(_a_F_pg_last_committed_xact_3))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						if base.Ui32(v24) <= base.Ui32(int32(2)) {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)) = uint8(v40)
							v42 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v8)+16)) = uint16(v42)
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							v58 = F_heap_form_tuple(m, v53, v8+int32(20), v8+int32(16))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
								v61 = F_HeapTupleHeaderGetDatum(m, v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return v61
								}
							}
						} else {
							v44 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v44)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v24
							v47 = F_Int64GetDatum(m, v23)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v47
								*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v22
								v51 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v8)+17)) = uint16(v51)
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
								v58 = F_heap_form_tuple(m, v53, v8+int32(20), v8+int32(16))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
									v61 = F_HeapTupleHeaderGetDatum(m, v60)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(32)
										return v61
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_error_commit_ts_disabled(m)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_pg_lltoa(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	if int64(0) <= l0 {
		v12 = l0
		v13 = int32(0)
	} else {
		v7 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
		v12 = int64(0) - l0
		v13 = int32(1)
	}
	v14 = l1 + v13
	v15 = int32(0)
	if v12 == int64(0) {
		v24 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v24)
		v195 = int32(1)
	} else {
		v31 = int32(1233)
		v36 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v12))^int32(63))*v31+v31) >> (uint(int32(12)) % 32))
		v39 = *(*int64)(unsafe.Add(mBase, uint32(v36<<(uint(int32(3))%32))+uint32(_c_F_pg_lltoa[0])))
		v41 = v36 + base.B2i32(base.Ui64(v39) <= base.Ui64(v12))
		if base.Ui64(int64(100000000)) <= base.Ui64(v12) {
			v45 = v12
			v48 = v15
			for {
				v54 = v14 + v41 - v48
				v55 = int32(8)
				v58 = base.I64_div_u_s(v45, int64(100000000))
				v62 = base.I32_wrap_i64(v45 + v58*int64(4194967296))
				v64 = base.I32_div_u_s(v62, int32(_a_F_pg_lltoa_0))
				v65 = int32(1)
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-v55))) = uint16(v67)
				v71 = int32(_a_F_pg_lltoa_1)
				v72 = base.I32_div_u_s(v62, v71)
				v73 = int32(100)
				v74 = base.I32_rem_u_s(v72, v73)
				v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(6)))) = uint16(v77)
				v83 = v62 - v72*v71
				v84 = int32(_a_F_pg_lltoa_2)
				v87 = base.I32_div_u_s(v83&v84, v73)
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(4)))) = uint16(v90)
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v83-v87*v73)&v84<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(2)))) = uint16(v101)
				v104 = v48 + v55
				if base.Ui64(int64(9999999999999999)) < base.Ui64(v45) {
					v45 = v58
					v48 = v104
					continue
				} else {
					break
				}
				break
			}
			v107 = v58
			v110 = v104
		} else {
			v107 = v12
			v110 = v15
		}
		v116 = base.I32_wrap_i64(v107)
		if base.Ui64(int64(10000)) <= base.Ui64(v107) {
			v120 = v14 + v41 - v110
			v121 = int32(4)
			v124 = base.I32_div_u_s(v116, int32(_a_F_pg_lltoa_1))
			v127 = v116 + v124*int32(-10000)
			v128 = int32(100)
			v129 = base.I32_div_u_s(v127, v128)
			v130 = int32(1)
			v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129<<(uint(v130)%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v120-v121))) = uint16(v132)
			v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v127-v129*v128)<<(uint(v130)%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v120-int32(2)))) = uint16(v141)
			v145 = v124
			v146 = v110 | v121
		} else {
			v145 = v116
			v146 = v110
		}
		if base.Ui32(int32(100)) <= base.Ui32(v145) {
			v154 = int32(2)
			v156 = int32(_a_F_pg_lltoa_2)
			v158 = int32(100)
			v159 = base.I32_div_u_s(v145&v156, v158)
			v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v145-v159*v158)&v156<<(uint(int32(1))%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v41-v146-v154))) = uint16(v167)
			v171 = v159
			v172 = v146 + v154
		} else {
			v171 = v145
			v172 = v146
		}
		if base.Ui32(int32(10)) <= base.Ui32(v171) {
			v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171<<(uint(int32(1))%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v41-v172-int32(2)))) = uint16(v181)
			v195 = v41
		} else {
			v184 = v171 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v184)
			v195 = v41
		}
	}
	v196 = v195 + v13
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v196))) = uint8(v198)
	return v196
}
func F_pg_mbcharcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbcharcliplen[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_c_F_pg_mbcharcliplen[1])))
	if v14 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_report_invalid_encoding_db(m, v19, v36, v20)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L23
	}
L2:
	;
	return v69
L3:
	;
	if l1 <= int32(0) {
		v69 = v4
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if l1 < l2 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v19 = l0
	v20 = l1
	v22 = v4
	v24 = v4
	goto L7
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v25 == int32(0) {
		v69 = v22
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v69 = v45
	goto L2
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbcharcliplen[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30*int32(28))+uint32(_c_F_pg_mbcharcliplen[2])))
	v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, v19)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v20 < v36 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = v24 + int32(1)
	if l2 < v42 {
		v69 = v22
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v45 = v22 + v36
	v46 = v20 - v36
	if int32(0) < v46 {
		v19 = v19 + v36
		v20 = v46
		v22 = v45
		v24 = v42
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v50 = l1
	goto L17
L16:
	;
	v50 = l2
	goto L17
L17:
	;
	if v50 <= int32(0) {
		v69 = v4
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v56 = v4
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v56))))
	if v60 == int32(0) {
		v69 = v56
		goto L2
	} else {
		goto L21
	}
L20:
	;
	v69 = v50
	goto L2
L21:
	;
	v64 = v56 + int32(1)
	if v64 != v50 {
		v56 = v64
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mcv_list_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_byteaout(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_mcv_list_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_pg_mcv_list_recv_0), int32(1511), int32(_a_F_pg_mcv_list_recv_1), int32(_a_F_pg_mcv_list_recv_2), int32(_a_F_pg_mcv_list_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_md5_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v9 = F_strlen(m, l0)
	mBase = m.M
	v10 = v9 + l2
	v13 = F_emscripten_builtin_malloc(m, v10+int32(1))
	mBase = m.M
	if v13 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(_a_F_pg_md5_encrypt_0)
		return int32(0)
	} else {
		if v9 != 0 {
			base.MemoryCopy(m, v13, l0, v9)
		} else {
		}
		if l2 != 0 {
			base.MemoryCopy(m, v13+v9, l1, l2)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_pg_md5_encrypt_1)
		v27 = F_pg_md5_hash(m, v13, v10, l3+int32(3), l4)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_emscripten_builtin_free(m, v13)
			mBase = m.M
			return v27
		}
	}
}
func F_pg_mule_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	if l1 <= int32(0) {
		v62 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v62 - l0
L2:
	;
	v9 = l1
	v10 = l0
	goto L3
L3:
	;
	v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	if int32(0) <= v13 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v62 = v56
	goto L1
L5:
	;
	v56 = v10 + v55
	v57 = v9 - v55
	if int32(0) < v57 {
		v9 = v57
		v10 = v56
		goto L3
	} else {
		goto L23
	}
L6:
	;
	if v13 != 0 {
		v55 = int32(1)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32((v13+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v38 = int32(2)
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = v10
	goto L1
L10:
	;
	if base.Ui32(v9) < base.Ui32(v38) {
		v62 = v10
		goto L1
	} else {
		goto L16
	}
L11:
	;
	if base.Ui32((v13+int32(112))&int32(255)) < base.Ui32(int32(12)) {
		v38 = int32(3)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if v13&int32(-2) == int32(-100) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = int32(4)
	goto L15
L14:
	;
	v37 = int32(1)
	goto L15
L15:
	;
	v38 = v37
	goto L10
L16:
	;
	if base.Ui32(v38) < base.Ui32(int32(2)) {
		v55 = v38
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v42 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+1)))
	if int32(0) <= v42 {
		v62 = v10
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v38 == int32(2) {
		v55 = v38
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+2)))
	if int32(0) <= v47 {
		v62 = v10
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v38) < base.Ui32(int32(4)) {
		v55 = v38
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+3)))
	if int32(0) <= v52 {
		v62 = v10
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v55 = v38
	goto L5
L23:
	;
	goto L4
}
func F_pg_next_dst_boundary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l6)+260))
	if v23 == v8 {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
		v29 = l6 + v26<<(uint(int32(4))%32)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pg_next_dst_boundary[1])))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pg_next_dst_boundary[2]))))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
		v232 = v8
	} else {
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+272)))
		if v39 == int32(1) {
			v42 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
			if v38 < v42 {
				v56 = v42
				v58 = l6 + int32(280)
				if v38 < v56 {
					v68 = v56 - v38
				} else {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
					v68 = v38 - v66
				}
				v70 = v68 - int64(1)
				v71 = int64(12622780800)
				v72 = base.I64_rem_s(v70, v71)
				v73 = v70 - v72
				v75 = v73 + v71
				v77 = int64(-12622780800) - v73
				if v38 < v56 {
					v79 = v75
				} else {
					v79 = v77
				}
				v80 = v79 + v38
				*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v80
				v82 = int32(-1)
				if v80 < v56 {
					v232 = v82
				} else {
					v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
					if v89 < v80 {
						v232 = v82
					} else {
						v93 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
						mBase = m.M
						v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
						v95 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						if v38 < v95 {
							v97 = v77
						} else {
							v97 = v75
						}
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v94 + v97
						v232 = v93
					}
				}
			} else {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+273)))
				if v45 != int32(1) {
					v101 = l6 + int32(280)
					v103 = v23 - int32(1)
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
					if v107 <= v38 {
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v115 = l6 + v112<<(uint(int32(4))%32)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
						v232 = v8
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
						if v120 <= v38 {
							v122 = int32(1)
							if v122 < v103 {
								v126 = v122
								v133 = v103
								for {
									v145 = int32(1)
									v146 = (v126 + v133) >> (uint(v145) % 32)
									v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
									v153 = base.B2i32(v38 < v152)
									if v38 < v152 {
										v154 = v126
									} else {
										v154 = v146 + v145
									}
									if v38 < v152 {
										v155 = v146
									} else {
										v155 = v133
									}
									if v154 < v155 {
										v126 = v154
										v133 = v155
										continue
									} else {
										break
									}
									break
								}
								v157 = v154
							} else {
								v157 = v122
							}
							v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v177 = v157 + l6
							v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
							v181 = int32(4)
							v183 = v176 + v180<<(uint(v181)%32)
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
							v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
							v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v198 = v176 + v195<<(uint(v181)%32)
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
							v232 = v122
						} else {
							v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
							v206 = int32(4)
							v208 = v204 + v205<<(uint(v206)%32)
							v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
							v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
							v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v218 = v204 + v215<<(uint(v206)%32)
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
							v232 = int32(1)
						}
					}
				} else {
					v53 = *(*int64)(unsafe.Add(mBase, uint32(l6+int32(272)+v23<<(uint(int32(3))%32))))
					if v38 <= v53 {
						v101 = l6 + int32(280)
						v103 = v23 - int32(1)
						v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
						if v107 <= v38 {
							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v115 = l6 + v112<<(uint(int32(4))%32)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
							v232 = v8
						} else {
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
							if v120 <= v38 {
								v122 = int32(1)
								if v122 < v103 {
									v126 = v122
									v133 = v103
									for {
										v145 = int32(1)
										v146 = (v126 + v133) >> (uint(v145) % 32)
										v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
										v153 = base.B2i32(v38 < v152)
										if v38 < v152 {
											v154 = v126
										} else {
											v154 = v146 + v145
										}
										if v38 < v152 {
											v155 = v146
										} else {
											v155 = v133
										}
										if v154 < v155 {
											v126 = v154
											v133 = v155
											continue
										} else {
											break
										}
										break
									}
									v157 = v154
								} else {
									v157 = v122
								}
								v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
								v177 = v157 + l6
								v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
								v181 = int32(4)
								v183 = v176 + v180<<(uint(v181)%32)
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
								v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
								v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
								v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
								v198 = v176 + v195<<(uint(v181)%32)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
								v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
								v232 = v122
							} else {
								v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
								v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
								v206 = int32(4)
								v208 = v204 + v205<<(uint(v206)%32)
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
								v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
								v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
								v218 = v204 + v215<<(uint(v206)%32)
								v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
								v232 = int32(1)
							}
						}
					} else {
						v55 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						v56 = v55
						v58 = l6 + int32(280)
						if v38 < v56 {
							v68 = v56 - v38
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
							v68 = v38 - v66
						}
						v70 = v68 - int64(1)
						v71 = int64(12622780800)
						v72 = base.I64_rem_s(v70, v71)
						v73 = v70 - v72
						v75 = v73 + v71
						v77 = int64(-12622780800) - v73
						if v38 < v56 {
							v79 = v75
						} else {
							v79 = v77
						}
						v80 = v79 + v38
						*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v80
						v82 = int32(-1)
						if v80 < v56 {
							v232 = v82
						} else {
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
							if v89 < v80 {
								v232 = v82
							} else {
								v93 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
								mBase = m.M
								v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
								v95 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
								if v38 < v95 {
									v97 = v77
								} else {
									v97 = v75
								}
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v94 + v97
								v232 = v93
							}
						}
					}
				}
			}
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+273)))
			if v45 != int32(1) {
				v101 = l6 + int32(280)
				v103 = v23 - int32(1)
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
				if v107 <= v38 {
					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
					v115 = l6 + v112<<(uint(int32(4))%32)
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
					v232 = v8
				} else {
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
					if v120 <= v38 {
						v122 = int32(1)
						if v122 < v103 {
							v126 = v122
							v133 = v103
							for {
								v145 = int32(1)
								v146 = (v126 + v133) >> (uint(v145) % 32)
								v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
								v153 = base.B2i32(v38 < v152)
								if v38 < v152 {
									v154 = v126
								} else {
									v154 = v146 + v145
								}
								if v38 < v152 {
									v155 = v146
								} else {
									v155 = v133
								}
								if v154 < v155 {
									v126 = v154
									v133 = v155
									continue
								} else {
									break
								}
								break
							}
							v157 = v154
						} else {
							v157 = v122
						}
						v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
						v177 = v157 + l6
						v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
						v181 = int32(4)
						v183 = v176 + v180<<(uint(v181)%32)
						v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
						v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v198 = v176 + v195<<(uint(v181)%32)
						v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
						v232 = v122
					} else {
						v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
						v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
						v206 = int32(4)
						v208 = v204 + v205<<(uint(v206)%32)
						v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
						v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
						v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
						v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v218 = v204 + v215<<(uint(v206)%32)
						v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
						v232 = int32(1)
					}
				}
			} else {
				v53 = *(*int64)(unsafe.Add(mBase, uint32(l6+int32(272)+v23<<(uint(int32(3))%32))))
				if v38 <= v53 {
					v101 = l6 + int32(280)
					v103 = v23 - int32(1)
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
					if v107 <= v38 {
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v115 = l6 + v112<<(uint(int32(4))%32)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
						v232 = v8
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
						if v120 <= v38 {
							v122 = int32(1)
							if v122 < v103 {
								v126 = v122
								v133 = v103
								for {
									v145 = int32(1)
									v146 = (v126 + v133) >> (uint(v145) % 32)
									v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
									v153 = base.B2i32(v38 < v152)
									if v38 < v152 {
										v154 = v126
									} else {
										v154 = v146 + v145
									}
									if v38 < v152 {
										v155 = v146
									} else {
										v155 = v133
									}
									if v154 < v155 {
										v126 = v154
										v133 = v155
										continue
									} else {
										break
									}
									break
								}
								v157 = v154
							} else {
								v157 = v122
							}
							v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v177 = v157 + l6
							v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
							v181 = int32(4)
							v183 = v176 + v180<<(uint(v181)%32)
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
							v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
							v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v198 = v176 + v195<<(uint(v181)%32)
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
							v232 = v122
						} else {
							v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
							v206 = int32(4)
							v208 = v204 + v205<<(uint(v206)%32)
							v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
							v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
							v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v218 = v204 + v215<<(uint(v206)%32)
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
							v232 = int32(1)
						}
					}
				} else {
					v55 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
					v56 = v55
					v58 = l6 + int32(280)
					if v38 < v56 {
						v68 = v56 - v38
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
						v68 = v38 - v66
					}
					v70 = v68 - int64(1)
					v71 = int64(12622780800)
					v72 = base.I64_rem_s(v70, v71)
					v73 = v70 - v72
					v75 = v73 + v71
					v77 = int64(-12622780800) - v73
					if v38 < v56 {
						v79 = v75
					} else {
						v79 = v77
					}
					v80 = v79 + v38
					*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v80
					v82 = int32(-1)
					if v80 < v56 {
						v232 = v82
					} else {
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
						if v89 < v80 {
							v232 = v82
						} else {
							v93 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
							mBase = m.M
							v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							v95 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							if v38 < v95 {
								v97 = v77
							} else {
								v97 = v75
							}
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v94 + v97
							v232 = v93
						}
					}
				}
			}
		}
	}
	m.G0 = v21 + int32(16)
	return v232
}
func F_pg_node_tree_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_textout(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_parameter_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = F_superuser_arg(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v20 = F_convert_GUC_name_for_parameter_acl(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_cstring_to_text(m, v20)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_SearchSysCache1(m, int32(43), v22)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 == int32(0) {
							v62 = int64(0)
							F_pfree(m, v20)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v22)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v69 = v62
									m.G0 = v11 + int32(16)
									return base.B2i32(v69 == int64(0))
								}
							}
						} else {
							v33 = F_SysCacheGetAttr(m, int32(43), v24, int32(3), v11+int32(15))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
								if v35 == int32(1) {
									v41 = F_acldefault(m, int32(27), int32(10))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v45 = int32(0)
										v46 = v41
										v49 = F_aclmask(m, v46, l1, int32(10), l2, int32(1))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											v51 = int32(0)
											if base.B2i32(v46 == v51)|base.B2i32(v45 == v46) == v51 {
												F_pfree(m, v46)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v24)
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														v62 = v49
														F_pfree(m, v20)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v22)
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v69 = v62
																m.G0 = v11 + int32(16)
																return base.B2i32(v69 == int64(0))
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v62 = v49
													F_pfree(m, v20)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v22)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v69 = v62
															m.G0 = v11 + int32(16)
															return base.B2i32(v69 == int64(0))
														}
													}
												}
											}
										}
									}
								} else {
									v43 = F_pg_detoast_datum(m, v33)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v45 = v33
										v46 = v43
										v49 = F_aclmask(m, v46, l1, int32(10), l2, int32(1))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											v51 = int32(0)
											if base.B2i32(v46 == v51)|base.B2i32(v45 == v46) == v51 {
												F_pfree(m, v46)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v24)
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														v62 = v49
														F_pfree(m, v20)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v22)
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v69 = v62
																m.G0 = v11 + int32(16)
																return base.B2i32(v69 == int64(0))
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v62 = v49
													F_pfree(m, v20)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v22)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v69 = v62
															m.G0 = v11 + int32(16)
															return base.B2i32(v69 == int64(0))
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
		} else {
			v69 = l2
			m.G0 = v11 + int32(16)
			return base.B2i32(v69 == int64(0))
		}
	}
}
func F_pg_popcount_masked_optimized(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v154 int64
	_ = v154
	v8 = int64(0)
	v9 = int32(_a_F_pg_popcount_masked_optimized_0)
	if (l0+int32(3))&int32(-4) == l0 {
		v16 = l1 * int32(16843009)
		v17 = l0
		v19 = v9
		v21 = int32(0)
		v24 = v8
		for {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v44 = base.I64_extend_i32_u(base.I32_popcnt(v25&v16)) + (base.I64_extend_i32_u(base.I32_popcnt(v29&v16)) + (base.I64_extend_i32_u(base.I32_popcnt(v33&v16)) + (v24 + base.I64_extend_i32_u(base.I32_popcnt(v37&v16)))))
			v45 = int32(16)
			v46 = v19 - v45
			v48 = v17 + v45
			v50 = v21 + int32(4)
			if v50 != int32(2040) {
				v17 = v48
				v19 = v46
				v21 = v50
				v24 = v44
				continue
			} else {
				break
			}
			break
		}
		v54 = v48
		v56 = v46
		v58 = int32(0)
		v61 = v44
		for {
			v62 = int32(4)
			v63 = v56 - v62
			v65 = v54 + v62
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
			v70 = v61 + base.I64_extend_i32_u(base.I32_popcnt(v66&v16))
			v72 = v58 + int32(1)
			if v72 != int32(2) {
				v54 = v65
				v56 = v63
				v58 = v72
				v61 = v70
				continue
			} else {
				break
			}
			break
		}
		v75 = v65
		v77 = v63
		v82 = v70
	} else {
		v75 = l0
		v77 = v9
		v82 = v8
	}
	if v77 == int32(0) {
		v154 = v82
	} else {
		v86 = v77 & int32(3)
		if v86 == int32(0) {
			v109 = v75
			v114 = v77
			v116 = v82
		} else {
			v90 = v75
			v94 = int32(0)
			v95 = v77
			v97 = v82
			for {
				v98 = int32(1)
				v99 = v90 + v98
				v101 = v95 - v98
				v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
				v104 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v102&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
				v105 = v97 + v104
				v107 = v94 + v98
				if v107 != v86 {
					v90 = v99
					v94 = v107
					v95 = v101
					v97 = v105
					continue
				} else {
					break
				}
				break
			}
			v109 = v99
			v114 = v101
			v116 = v105
		}
		if base.Ui32(v77) < base.Ui32(int32(4)) {
			v154 = v116
		} else {
			v119 = v109
			v124 = v114
			v126 = v116
			for {
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+3)))
				v129 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
				v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+2)))
				v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v130&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
				v135 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v133&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
				v138 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v136&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
				v142 = v129 + (v132 + (v135 + (v126 + v138)))
				v143 = int32(4)
				v146 = v124 - v143
				if v146 != 0 {
					v119 = v119 + v143
					v124 = v146
					v126 = v142
					continue
				} else {
					break
				}
				break
			}
			v154 = v142
		}
	}
	return v154
}
func F_pg_popcount_optimized(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int64
	_ = v155
	v7 = int64(0)
	if base.B2i32(l0 != (l0+int32(3))&int32(-4))|base.B2i32(l1 < int32(4)) != 0 {
		v86 = l0
		v87 = l1
		v92 = v7
	} else {
		v17 = l1 - int32(4)
		v21 = int32(base.Ui32(v17)>>(uint(int32(2))%32)) + int32(1)
		v23 = v21 & int32(3)
		if base.Ui32(int32(12)) <= base.Ui32(v17) {
			v28 = l0
			v29 = l1
			v32 = int32(0)
			v34 = v7
			for {
				v35 = int32(16)
				v36 = v29 - v35
				v38 = v28 + v35
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v54 = base.I64_extend_i32_u(base.I32_popcnt(v39)) + (base.I64_extend_i32_u(base.I32_popcnt(v42)) + (base.I64_extend_i32_u(base.I32_popcnt(v45)) + (v34 + base.I64_extend_i32_u(base.I32_popcnt(v48)))))
				v56 = v32 + int32(4)
				if v56 != v21&int32(2147483644) {
					v28 = v38
					v29 = v36
					v32 = v56
					v34 = v54
					continue
				} else {
					break
				}
				break
			}
			if v23 == int32(0) {
				v86 = v38
				v87 = v36
				v92 = v54
			} else {
				v60 = v38
				v61 = v36
				v66 = v54
				v68 = v60
				v69 = v61
				v70 = int32(0)
				v74 = v66
				for {
					v75 = int32(4)
					v76 = v69 - v75
					v78 = v68 + v75
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					v82 = v74 + base.I64_extend_i32_u(base.I32_popcnt(v79))
					v84 = v70 + int32(1)
					if v84 != v23 {
						v68 = v78
						v69 = v76
						v70 = v84
						v74 = v82
						continue
					} else {
						break
					}
					break
				}
				v86 = v78
				v87 = v76
				v92 = v82
			}
		} else {
			v60 = l0
			v61 = l1
			v66 = v7
			v68 = v60
			v69 = v61
			v70 = int32(0)
			v74 = v66
			for {
				v75 = int32(4)
				v76 = v69 - v75
				v78 = v68 + v75
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
				v82 = v74 + base.I64_extend_i32_u(base.I32_popcnt(v79))
				v84 = v70 + int32(1)
				if v84 != v23 {
					v68 = v78
					v69 = v76
					v70 = v84
					v74 = v82
					continue
				} else {
					break
				}
				break
			}
			v86 = v78
			v87 = v76
			v92 = v82
		}
	}
	if v87 == int32(0) {
		v155 = v92
	} else {
		v96 = v87 & int32(3)
		if v96 == int32(0) {
			v117 = v86
			v119 = v87
			v123 = v92
		} else {
			v100 = v86
			v102 = v87
			v104 = int32(0)
			v106 = v92
			for {
				v107 = int32(1)
				v108 = v100 + v107
				v110 = v102 - v107
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
				v112 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_pg_popcount_optimized[0]))))
				v113 = v106 + v112
				v115 = v104 + v107
				if v115 != v96 {
					v100 = v108
					v102 = v110
					v104 = v115
					v106 = v113
					continue
				} else {
					break
				}
				break
			}
			v117 = v108
			v119 = v110
			v123 = v113
		}
		if base.Ui32(v87) < base.Ui32(int32(4)) {
			v155 = v123
		} else {
			v126 = v117
			v128 = v119
			v132 = v123
			for {
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+3)))
				v134 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v133)+uint32(_c_F_pg_popcount_optimized[0]))))
				v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
				v136 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+uint32(_c_F_pg_popcount_optimized[0]))))
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
				v138 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v137)+uint32(_c_F_pg_popcount_optimized[0]))))
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
				v140 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_pg_popcount_optimized[0]))))
				v144 = v134 + (v136 + (v138 + (v132 + v140)))
				v145 = int32(4)
				v148 = v128 - v145
				if v148 != 0 {
					v126 = v126 + v145
					v128 = v148
					v132 = v144
					continue
				} else {
					break
				}
				break
			}
			v155 = v144
		}
	}
	return v155
}
func F_pg_prewarm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
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
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v274 int64
	_ = v274
	var v282 int64
	_ = v282
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v316 int64
	_ = v316
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int64
	_ = v393
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v421 int64
	_ = v421
	var v436 int32
	_ = v436
	var v438 int64
	_ = v438
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v20 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L13
	} else {
		goto L189
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L13
	} else {
		goto L185
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L13
	} else {
		goto L181
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L13
	} else {
		goto L177
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L13
	} else {
		goto L173
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L13
	} else {
		goto L168
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L13
	} else {
		goto L164
	}
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v23 == int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L13
	} else {
		goto L160
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v28 = F_pg_detoast_datum_packed(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v121 == int32(1) {
		goto L5
	} else {
		goto L42
	}
L13:
	;
	return int32(0)
L14:
	;
	v32 = F_text_to_cstring(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v34 = int32(_a_F_pg_prewarm_0)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_prewarm[0])))
	if base.B2i32(v37 == int32(0))|base.B2i32(v37 != v40) != 0 {
		v58 = v37
		v59 = v40
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v60 == int32(0) {
		v120 = v6
		goto L12
	} else {
		goto L23
	}
L17:
	;
	v60 = v58 - v59
	goto L16
L18:
	;
	v43 = v32
	v44 = v34
	goto L19
L19:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v48
		v59 = v47
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v58 = v48
	v59 = v47
	goto L17
L21:
	;
	v51 = int32(1)
	if v48 == v47 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v63 = int32(_a_F_pg_prewarm_1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_prewarm[1])))
	if base.B2i32(v66 == int32(0))|base.B2i32(v66 != v69) != 0 {
		v87 = v66
		v88 = v69
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v87-v88 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	goto L24
L26:
	;
	v72 = v32
	v73 = v63
	goto L27
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v77 == int32(0) {
		v87 = v77
		v88 = v76
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v87 = v77
	v88 = v76
	goto L25
L29:
	;
	v80 = int32(1)
	if v77 == v76 {
		v72 = v72 + v80
		v73 = v73 + v80
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v120 = int32(1)
	goto L12
L32:
	;
	goto L33
L33:
	;
	v93 = int32(_a_F_pg_prewarm_2)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_prewarm[2])))
	if base.B2i32(v96 == int32(0))|base.B2i32(v96 != v99) != 0 {
		v117 = v96
		v118 = v99
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v117-v118 != 0 {
		goto L6
	} else {
		goto L41
	}
L35:
	;
	goto L34
L36:
	;
	v102 = v32
	v103 = v93
	goto L37
L37:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v107
		v118 = v106
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v117 = v107
	v118 = v106
	goto L35
L39:
	;
	v110 = int32(1)
	if v107 == v106 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v120 = v6
	goto L12
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v125 = F_pg_detoast_datum_packed(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v127 = F_text_to_cstring(m, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v129 = F_forkname_to_number(m, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v131 = F_get_rel_relkind(m, v26)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[3]))
	v170 = F_pg_class_aclcheck(m, v165, v168, int64(2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L63
	}
L47:
	;
	if v131&int32(-33) == int32(73) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v137 = int32(1)
	v139 = F_IndexGetRelation(m, v26, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v158 = int32(1)
	v160 = F_relation_open(m, v26, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L13
	} else {
		goto L61
	}
L51:
	;
	if v139 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v144 = F_relation_open(m, v26, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_LockRelationOid(m, v139, int32(1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L13
	} else {
		goto L56
	}
L55:
	;
	v592 = v144
	goto L1
L56:
	;
	v150 = F_relation_open(m, v26, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	if v26 == v139 {
		v164 = v150
		v165 = v139
		v166 = v137
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v155 = F_IndexGetRelation(m, v26, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	if v155 == v139 {
		v164 = v150
		v165 = v139
		v166 = int32(0)
		goto L46
	} else {
		goto L60
	}
L60:
	;
	v592 = v150
	goto L1
L61:
	;
	if v26 == int32(0) {
		v592 = v160
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v164 = v160
	v165 = v26
	v166 = v158
	goto L46
L63:
	;
	if v170 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v173 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172)+119)))
	switch v173 - int32(73) {
	case 0, 32:
		v183 = int32(20)
		goto L68
	default:
		goto L69
	case 10:
		goto L73
	case 29:
		goto L70
	case 36:
		goto L71
	case 45:
		goto L72
	}
L65:
	;
	goto L66
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+119)))
	switch v191 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L76
	default:
		goto L77
	}
L67:
	;
	v186 = F_get_rel_name(m, v26)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L74
	}
L68:
	;
	v185 = v183
	goto L67
L69:
	;
	v183 = int32(41)
	goto L68
L70:
	;
	v185 = int32(18)
	goto L67
L71:
	;
	v185 = int32(23)
	goto L67
L72:
	;
	v185 = int32(51)
	goto L67
L73:
	;
	v185 = int32(37)
	goto L67
L74:
	;
	F_aclcheck_error(m, v170, v185, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	goto L66
L76:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	if v219 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v201 + int32(4)
	F_errmsg(m, int32(_a_F_pg_prewarm_3), v18+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v211 = int32(*(*int8)(unsafe.Add(mBase, uint32(v210)+119)))
	F_errdetail_relkind_not_supported(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(159), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L13
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
	v245 = v219
	goto L85
L84:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v223
	v227 = F_smgropen(m, v18+int32(88), v220)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L86
	}
L85:
	;
	v246 = F_smgrexists(m, v245, v129)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L91
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
	if v231 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v245 = v243
	goto L85
L88:
	;
	v239 = v231
	goto L90
L89:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
	v239 = v237
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+72)) = v239 + int32(1)
	goto L87
L91:
	;
	if v246 == int32(0) {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v250 = F_RelationGetNumberOfBlocksInFork(m, v164, v129)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	v252 = base.I64_extend_i32_u(v250)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v253 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
	if base.B2i32(v257 < int64(0))|base.B2i32(v252 <= v257) != 0 {
		goto L3
	} else {
		goto L97
	}
L95:
	;
	v262 = int64(0)
	goto L96
L96:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v263 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v262 = v257
	goto L96
L98:
	;
	if v60 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v274 = v252 - int64(1)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v268)))
	if base.B2i32(v269 < int64(0))|base.B2i32(v252 <= v269) != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v274 = v269
	goto L98
L103:
	;
	F_relation_close(m, v164, int32(1))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L154
	}
L104:
	;
	if v274 < v262 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if v120 != 0 {
		goto L118
	} else {
		goto L119
	}
L107:
	;
	v438 = int64(0)
	goto L103
L108:
	;
	goto L109
L109:
	;
	v282 = v262
	goto L110
L110:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[4]))
	if v297 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v438 = v274 + int64(1) - v262
	goto L103
L112:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L13
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_PrefetchBuffer(m, v18+int32(104), v164, v129, base.I32_wrap_i64(v282))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	if base.B2i32(v282 == v274) == int32(0) {
		v282 = v282 + int64(1)
		goto L110
	} else {
		goto L117
	}
L117:
	;
	goto L111
L118:
	;
	if v274 < v262 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = base.I32_wrap_i64(v274) + int32(1)
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+104)) = uint32(v262)
	v382 = int32(0)
	v387 = F_read_stream_begin_relation(m, int32(13), v382, v164, v129, int32(120), v18+int32(104), v382)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L13
	} else {
		goto L140
	}
L121:
	;
	v438 = int64(0)
	goto L103
L122:
	;
	goto L123
L123:
	;
	v316 = v262
	goto L124
L124:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[4]))
	if v331 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v438 = v274 + int64(1) - v262
	goto L103
L126:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L13
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	if v334 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v338
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v340
	v344 = F_smgropen(m, v18+int32(32), v337)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L13
	} else {
		goto L133
	}
L131:
	;
	v361 = v334
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = int32(_a_F_pg_prewarm_6)
	F_smgrreadv(m, v361, v129, base.I32_wrap_i64(v316), v18+int32(104))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L13
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v344
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+72))
	if v348 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v361 = v360
	goto L132
L135:
	;
	v356 = v348
	goto L137
L136:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v344)+76))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v350
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v344)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v352
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v344)+72))
	v356 = v354
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+72)) = v356 + int32(1)
	goto L134
L138:
	;
	if base.B2i32(v316 == v274) == int32(0) {
		v316 = v316 + int64(1)
		goto L124
	} else {
		goto L139
	}
L139:
	;
	goto L125
L140:
	;
	if v262 <= v274 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v393 = v262
	goto L144
L142:
	;
	v421 = int64(0)
	goto L143
L143:
	;
	F_read_stream_end(m, v387)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L13
	} else {
		goto L153
	}
L144:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[4]))
	if v408 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v421 = v274 + int64(1) - v262
	goto L143
L146:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L13
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v412 = F_read_stream_next_buffer(m, v387, int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L13
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	F_ReleaseBuffer(m, v412)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	if v393 != v274 {
		v393 = v393 + int64(1)
		goto L144
	} else {
		goto L152
	}
L152:
	;
	goto L145
L153:
	;
	v438 = v421
	goto L103
L154:
	;
	if v166 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_UnlockRelationOid(m, v165, int32(1))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L13
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v460 = F_Int64GetDatum(m, v438)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L13
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	m.G0 = v18 + int32(112)
	return v460
L160:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_7), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(83), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L13
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_8), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(88), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L13
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L13
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_9), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L13
	} else {
		goto L170
	}
L170:
	;
	F_errhint(m, int32(_a_F_pg_prewarm_10), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L13
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(102), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L13
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L174
	}
L174:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_11), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(108), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L13
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L13
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v127
	F_errmsg(m, int32(_a_F_pg_prewarm_12), v18+int32(80))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L13
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(166), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v252 - int64(1)
	F_errmsg(m, int32(_a_F_pg_prewarm_13), v18-int32(-64))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L13
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(179), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L13
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L13
	} else {
		goto L186
	}
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v252 - int64(1)
	F_errmsg(m, int32(_a_F_pg_prewarm_14), v18+int32(48))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L13
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(190), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L13
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L13
	} else {
		goto L190
	}
L190:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v592)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v602 + int32(4)
	F_errmsg(m, int32(_a_F_pg_prewarm_15), v18)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L13
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(147), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L13
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_read_file_all_missing(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_convert_and_check_filename(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v16 = F_read_binary_file(m, v10, int64(0), int64(-1), base.B2i32(v9 != int32(0)))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int32(0)
				} else {
					v24 = int32(4)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					F_pg_verifymbstr(m, v16+v24, int32(base.Ui32(v26)>>(uint(int32(2))%32))-v24)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			}
		}
	}
}
func F_pg_reload_conf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(1)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reload_conf[0]))
	v6 = F_kill(m, v4, v2)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v28 = v2
			return v28
		} else {
			v12 = int32(0)
			v15 = F_errstart(m, int32(19), v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 == int32(0) {
					v28 = v12
					return v28
				} else {
					F_errmsg(m, int32(_a_F_pg_reload_conf_0), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_reload_conf_1), int32(293), int32(_a_F_pg_reload_conf_2))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = v12
							return v28
						}
					}
				}
			}
		}
	}
}
func F_pg_rotate_logfile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v2 = int32(0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rotate_logfile[0])))
	if v4 == v2 {
		v9 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v30 = v2
				return v30
			} else {
				F_errmsg(m, int32(_a_F_pg_rotate_logfile_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_rotate_logfile_1), int32(313), int32(_a_F_pg_rotate_logfile_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		F_SendPostmasterSignal(m, int32(3))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v30 = int32(1)
			return v30
		}
	}
}
func F_pg_rusage_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	v3 = l0 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1)
	v13 = F___syscall_ret(m, int32(0))
	mBase = m.M
	F_gettimeofday(m, l0)
	mBase = m.M
	return
}
func F_pg_saslprep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v17 = l0
	goto L6
L1:
	;
	m.G0 = v12 + int32(16)
	return v728
L2:
	;
	F_pfree(m, v143)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L12
	} else {
		goto L210
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = int32(0)
	goto L2
L4:
	;
	v728 = int32(-1)
	goto L1
L5:
	;
	if v19 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17))))
	if int32(0) < v19 {
		v17 = v17 + int32(1)
		goto L6
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	goto L7
L9:
	;
	v26 = F_pstrdup(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v33 = F_strlen(m, l0)
	mBase = m.M
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
	if v26 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v728 = v3
	goto L1
L15:
	;
	v37 = v33
	v38 = v3
	v42 = l0
	goto L18
L16:
	;
	v134 = v3
	goto L17
L17:
	;
	v143 = F_palloc(m, v134<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L59
	}
L18:
	;
	v43 = int32(-2)
	v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42))))
	if int32(0) <= v44 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v123 < int32(0) {
		v728 = v43
		goto L1
	} else {
		goto L57
	}
L20:
	;
	if base.Ui32(v37) < base.Ui32(v68) {
		v728 = v43
		goto L1
	} else {
		goto L33
	}
L21:
	;
	v68 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v49 = v44 & int32(255)
	if v49&int32(224) == int32(192) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v68 = int32(2)
	goto L20
L25:
	;
	goto L26
L26:
	;
	if v49&int32(240) == int32(224) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v68 = int32(3)
	goto L20
L28:
	;
	goto L29
L29:
	;
	if v49&int32(248) == int32(240) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v66 = int32(4)
	goto L32
L31:
	;
	v66 = int32(1)
	goto L32
L32:
	;
	v68 = v66
	goto L20
L33:
	;
	v70 = int32(0)
	switch v68 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L39
	case 2:
		goto L40
	case 3:
		goto L41
	default:
		v119 = v70
		goto L35
	}
L34:
	;
	if v119 == int32(0) {
		v728 = v43
		goto L1
	} else {
		goto L55
	}
L35:
	;
	goto L34
L36:
	;
	v119 = base.B2i32(base.Ui32(v111&int32(255)) < base.Ui32(int32(245)))
	goto L35
L37:
	;
	if base.I32_extend8_s(v106) < int32(-62) {
		v119 = v70
		goto L35
	} else {
		goto L54
	}
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v106 = v105
	goto L37
L39:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	switch v80 - int32(224) {
	case 0:
		goto L48
	default:
		goto L44
	case 13:
		goto L47
	case 16:
		goto L46
	case 20:
		goto L45
	}
L40:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+2)))
	if int32(-65) < v76 {
		v119 = v70
		goto L35
	} else {
		goto L43
	}
L41:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+3)))
	if int32(-65) < v73 {
		v119 = v70
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L39
L44:
	;
	if v79 <= int32(-65) {
		v106 = v80
		goto L37
	} else {
		goto L53
	}
L45:
	;
	if int32(-113) < v79 {
		v119 = v70
		goto L35
	} else {
		goto L52
	}
L46:
	;
	if base.Ui32((v79-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v119 = v70
		goto L35
	} else {
		goto L51
	}
L47:
	;
	if int32(-97) < v79 {
		v119 = v70
		goto L35
	} else {
		goto L50
	}
L48:
	;
	v83 = int32(224)
	if base.Ui32(v83) <= base.Ui32((v79-int32(-64))&int32(255)) {
		v111 = v83
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v119 = v70
	goto L35
L50:
	;
	v111 = int32(237)
	goto L36
L51:
	;
	v111 = int32(240)
	goto L36
L52:
	;
	v111 = int32(244)
	goto L36
L53:
	;
	v119 = v70
	goto L35
L54:
	;
	v111 = v106
	goto L36
L55:
	;
	v123 = v38 + int32(1)
	v125 = v37 - v68
	if v125 != 0 {
		v37 = v125
		v38 = v123
		v42 = v68 + v42
		goto L18
	} else {
		goto L56
	}
L56:
	;
	goto L19
L57:
	;
	if base.Ui32(int32(268435454)) < base.Ui32(v123) {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v134 = v123
	goto L17
L59:
	;
	if v143 == int32(0) {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if v134 == int32(0) {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v150 = l0
	v156 = int32(0)
	goto L62
L62:
	;
	v162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150))))
	v164 = v162 & int32(255)
	if int32(0) <= v162 {
		v221 = v164
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v252 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v143+v134<<(uint(int32(2))%32)))) = v252
	v262 = v252
	v265 = v252
	goto L88
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143+v156<<(uint(int32(2))%32)))) = v221
	v223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150))))
	if int32(0) <= v223 {
		goto L75
	} else {
		goto L76
	}
L65:
	;
	if v164&int32(224) == int32(192) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v213))))
	v221 = v214 | v216&int32(63)
	goto L64
L67:
	;
	v213 = int32(1)
	v214 = v164 << (uint(int32(6)) % 32) & int32(1984)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if v164&int32(240) == int32(224) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	v213 = int32(2)
	v214 = v164<<(uint(int32(12))%32)&int32(_a_F_pg_saslprep_0) | v185&int32(63)<<(uint(int32(6))%32)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if v164&int32(248) != int32(240) {
		v221 = int32(-1)
		goto L64
	} else {
		goto L73
	}
L73:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	v202 = int32(63)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+2)))
	v213 = int32(3)
	v214 = v164<<(uint(int32(18))%32)&int32(_a_F_pg_saslprep_1) | v201&v202<<(uint(int32(12))%32) | v207&v202<<(uint(int32(6))%32)
	goto L66
L74:
	;
	v250 = v156 + int32(1)
	if v250 != v134 {
		v150 = v247 + v150
		v156 = v250
		goto L62
	} else {
		goto L87
	}
L75:
	;
	v247 = int32(1)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v228 = v223 & int32(255)
	if v228&int32(224) == int32(192) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v247 = int32(2)
	goto L74
L79:
	;
	goto L80
L80:
	;
	if v228&int32(240) == int32(224) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v247 = int32(3)
	goto L74
L82:
	;
	goto L83
L83:
	;
	if v228&int32(248) == int32(240) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v245 = int32(4)
	goto L86
L85:
	;
	v245 = int32(1)
	goto L86
L86:
	;
	v247 = v245
	goto L74
L87:
	;
	goto L63
L88:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v143+v265<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v271
	if base.Ui32(int32(-12130)) < base.Ui32(v271-int32(_a_F_pg_saslprep_2)) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v314 = v143 + v308<<(uint(int32(2))%32)
	v315 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v315
	if v308 == v315 {
		goto L2
	} else {
		goto L103
	}
L90:
	;
	v310 = v265 + int32(1)
	if v310 != v134 {
		v262 = v308
		v265 = v310
		goto L88
	} else {
		goto L102
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143+v262<<(uint(int32(2))%32)))) = v300
	v308 = v262 + int32(1)
	goto L90
L92:
	;
	v284 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_3), int32(6), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L12
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v271
	if base.Ui32(v271-int32(_a_F_pg_saslprep_5)) <= base.Ui32(int32(-65108)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	if v284 != 0 {
		v300 = int32(32)
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v300 = v271
	goto L91
L98:
	;
	goto L99
L99:
	;
	v295 = int32(8)
	v298 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_6), v295, v295, int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	if v298 != 0 {
		v308 = v262
		goto L90
	} else {
		goto L101
	}
L101:
	;
	v300 = v271
	goto L91
L102:
	;
	goto L89
L103:
	;
	v320 = F_unicode_normalize(m, int32(2), v143)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	if v320 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_pfree(m, v143)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v308 <= int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L4
L109:
	;
	F_pfree(m, v143)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L12
	} else {
		goto L208
	}
L110:
	;
	F_pfree(m, v143)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L12
	} else {
		goto L206
	}
L111:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v449 != 0 {
		goto L144
	} else {
		goto L145
	}
L112:
	;
	v329 = int32(0)
	goto L113
L113:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v143+v329<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v341
	if base.Ui32(int32(_a_F_pg_saslprep_7)) <= base.Ui32(v341) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v369 = int32(0)
	goto L123
L115:
	;
	v366 = v329 + int32(1)
	if v366 != v308 {
		v329 = v366
		goto L113
	} else {
		goto L122
	}
L116:
	;
	v346 = v12 + int32(12)
	v351 = F_bsearch(m, v346, int32(_a_F_pg_saslprep_8), int32(36), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L117
	}
L117:
	;
	if v351 != 0 {
		goto L110
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v341
	if base.Ui32(v341-int32(_a_F_pg_saslprep_9)) <= base.Ui32(int32(-982494)) {
		goto L115
	} else {
		goto L119
	}
L119:
	;
	v362 = F_bsearch(m, v346, int32(_a_F_pg_saslprep_10), int32(396), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	if v362 != 0 {
		goto L110
	} else {
		goto L121
	}
L121:
	;
	goto L115
L122:
	;
	goto L114
L123:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v143+v369<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v381
	if base.Ui32(int32(-63808)) < base.Ui32(v381-int32(_a_F_pg_saslprep_11)) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v314-int32(4))))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v403 = int32(0)
	goto L132
L125:
	;
	goto L124
L126:
	;
	v393 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_12), int32(34), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L12
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v396 = v369 + int32(1)
	if v396 != v308 {
		v369 = v396
		goto L123
	} else {
		goto L131
	}
L129:
	;
	if v393 != 0 {
		goto L125
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L111
L132:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v143+v403<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v415
	if base.Ui32(int32(-1114046)) < base.Ui32(v415-int32(_a_F_pg_saslprep_13)) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v432 = F_is_code_in_table(m, v401)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L12
	} else {
		goto L140
	}
L134:
	;
	v427 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_14), int32(360), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L12
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v430 = v403 + int32(1)
	if v430 != v308 {
		v403 = v430
		goto L132
	} else {
		goto L139
	}
L137:
	;
	if v427 != 0 {
		goto L110
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	goto L133
L140:
	;
	if v432 == int32(0) {
		goto L110
	} else {
		goto L141
	}
L141:
	;
	v436 = F_is_code_in_table(m, v400)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L12
	} else {
		goto L142
	}
L142:
	;
	if v436 == int32(0) {
		goto L110
	} else {
		goto L143
	}
L143:
	;
	goto L111
L144:
	;
	v451 = v449
	v454 = v320
	v457 = int32(0)
	goto L147
L145:
	;
	v563 = int32(1)
	goto L146
L146:
	;
	v564 = F_palloc(m, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L12
	} else {
		goto L173
	}
L147:
	;
	if base.Ui32(v451) <= base.Ui32(int32(127)) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v563 = v549 + int32(1)
	goto L146
L149:
	;
	v524 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
	if int32(0) <= v524 {
		goto L160
	} else {
		goto L161
	}
L150:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v451)
	goto L149
L151:
	;
	goto L152
L152:
	;
	if base.Ui32(v451) <= base.Ui32(int32(2047)) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v468 = v451&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v468)
	v473 = int32(base.Ui32(v451)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v473)
	goto L149
L154:
	;
	goto L155
L155:
	;
	if base.Ui32(v451) <= base.Ui32(int32(_a_F_pg_saslprep_15)) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v477 = int32(63)
	v479 = int32(128)
	v480 = v451&v477 | v479
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v480)
	v485 = int32(base.Ui32(v451)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v485)
	v492 = int32(base.Ui32(v451)>>(uint(int32(6))%32))&v477 | v479
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v492)
	goto L149
L157:
	;
	goto L158
L158:
	;
	v494 = int32(63)
	v496 = int32(128)
	v497 = v451&v494 | v496
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v497)
	v504 = int32(base.Ui32(v451)>>(uint(int32(6))%32))&v494 | v496
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v504)
	v511 = int32(base.Ui32(v451)>>(uint(int32(12))%32))&v494 | v496
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v511)
	v518 = int32(base.Ui32(v451)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v518)
	goto L149
L159:
	;
	v549 = v548 + v457
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	if v550 != 0 {
		v451 = v550
		v454 = v454 + int32(4)
		v457 = v549
		goto L147
	} else {
		goto L172
	}
L160:
	;
	v548 = int32(1)
	goto L159
L161:
	;
	goto L162
L162:
	;
	v529 = v524 & int32(255)
	if v529&int32(224) == int32(192) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v548 = int32(2)
	goto L159
L164:
	;
	goto L165
L165:
	;
	if v529&int32(240) == int32(224) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v548 = int32(3)
	goto L159
L167:
	;
	goto L168
L168:
	;
	if v529&int32(248) == int32(240) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v546 = int32(4)
	goto L171
L170:
	;
	v546 = int32(1)
	goto L171
L171:
	;
	v548 = v546
	goto L159
L172:
	;
	goto L148
L173:
	;
	if v564 == int32(0) {
		goto L109
	} else {
		goto L174
	}
L174:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v568 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v569 = v564
	v572 = v320
	v573 = v568
	goto L178
L176:
	;
	v667 = v564
	goto L177
L177:
	;
	v676 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v667))) = uint8(v676)
	F_pfree(m, v143)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L12
	} else {
		goto L204
	}
L178:
	;
	if base.Ui32(v573) <= base.Ui32(int32(127)) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v667 = v665
	goto L177
L180:
	;
	v640 = int32(*(*int8)(unsafe.Add(mBase, uint32(v569))))
	if int32(0) <= v640 {
		goto L191
	} else {
		goto L192
	}
L181:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v569))) = uint8(v573)
	goto L180
L182:
	;
	goto L183
L183:
	;
	if base.Ui32(v573) <= base.Ui32(int32(2047)) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v586 = v573&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)) = uint8(v586)
	v591 = int32(base.Ui32(v573)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v569))) = uint8(v591)
	goto L180
L185:
	;
	goto L186
L186:
	;
	if base.Ui32(v573) <= base.Ui32(int32(_a_F_pg_saslprep_15)) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v595 = int32(63)
	v597 = int32(128)
	v598 = v573&v595 | v597
	*(*uint8)(unsafe.Add(mBase, uint32(v569)+2)) = uint8(v598)
	v603 = int32(base.Ui32(v573)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v569))) = uint8(v603)
	v610 = int32(base.Ui32(v573)>>(uint(int32(6))%32))&v595 | v597
	*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)) = uint8(v610)
	goto L180
L188:
	;
	goto L189
L189:
	;
	v612 = int32(63)
	v614 = int32(128)
	v615 = v573&v612 | v614
	*(*uint8)(unsafe.Add(mBase, uint32(v569)+3)) = uint8(v615)
	v622 = int32(base.Ui32(v573)>>(uint(int32(6))%32))&v612 | v614
	*(*uint8)(unsafe.Add(mBase, uint32(v569)+2)) = uint8(v622)
	v629 = int32(base.Ui32(v573)>>(uint(int32(12))%32))&v612 | v614
	*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)) = uint8(v629)
	v636 = int32(base.Ui32(v573)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v569))) = uint8(v636)
	goto L180
L190:
	;
	v665 = v664 + v569
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v666 != 0 {
		v569 = v665
		v572 = v572 + int32(4)
		v573 = v666
		goto L178
	} else {
		goto L203
	}
L191:
	;
	v664 = int32(1)
	goto L190
L192:
	;
	goto L193
L193:
	;
	v645 = v640 & int32(255)
	if v645&int32(224) == int32(192) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v664 = int32(2)
	goto L190
L195:
	;
	goto L196
L196:
	;
	if v645&int32(240) == int32(224) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v664 = int32(3)
	goto L190
L198:
	;
	goto L199
L199:
	;
	if v645&int32(248) == int32(240) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v662 = int32(4)
	goto L202
L201:
	;
	v662 = int32(1)
	goto L202
L202:
	;
	v664 = v662
	goto L190
L203:
	;
	goto L179
L204:
	;
	F_pfree(m, v320)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L12
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v564
	v728 = v676
	goto L1
L206:
	;
	F_pfree(m, v320)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L12
	} else {
		goto L207
	}
L207:
	;
	v728 = int32(-3)
	goto L1
L208:
	;
	F_pfree(m, v320)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L12
	} else {
		goto L209
	}
L209:
	;
	goto L4
L210:
	;
	v728 = int32(-3)
	goto L1
}
func F_pg_sequence_last_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_init_sequence(m, v10, v8+int32(28), v8+int32(24))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[0]))
		v22 = F_pg_class_aclcheck(m, v10, v20, int64(258))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			if v22 != 0 {
				F_relation_close(m, v24, int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v65 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
					v69 = int32(0)
					m.G0 = v8 + int32(32)
					return v69
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+118)))
				switch v26 - int32(112) {
				case 0:
					v47 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
						F_UnlockReleaseBuffer(m, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_relation_close(m, v24, int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v50 != int32(1) {
									v65 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
									v69 = int32(0)
									m.G0 = v8 + int32(32)
									return v69
								} else {
									v59 = F_Int64GetDatum(m, v49)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v69 = v59
										m.G0 = v8 + int32(32)
										return v69
									}
								}
							}
						}
					}
				default:
					v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])))
					if v34 == int32(1) {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[2]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+316))
						v42 = base.B2i32(v40 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])) = uint8(v42)
						v44 = v42
					} else {
						v44 = int32(0)
					}
					if v44 != 0 {
						F_relation_close(m, v24, int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v65 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
							v69 = int32(0)
							m.G0 = v8 + int32(32)
							return v69
						}
					} else {
						v47 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
							F_UnlockReleaseBuffer(m, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_relation_close(m, v24, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									if v50 != int32(1) {
										v65 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
										v69 = int32(0)
										m.G0 = v8 + int32(32)
										return v69
									} else {
										v59 = F_Int64GetDatum(m, v49)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v69 = v59
											m.G0 = v8 + int32(32)
											return v69
										}
									}
								}
							}
						}
					}
				case 4:
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
					if v29 != int32(1) {
						F_relation_close(m, v24, int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v65 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
							v69 = int32(0)
							m.G0 = v8 + int32(32)
							return v69
						}
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])))
						if v34 == int32(1) {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[2]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+316))
							v42 = base.B2i32(v40 != int32(2))
							*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])) = uint8(v42)
							v44 = v42
						} else {
							v44 = int32(0)
						}
						if v44 != 0 {
							F_relation_close(m, v24, int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v65 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
								v69 = int32(0)
								m.G0 = v8 + int32(32)
								return v69
							}
						} else {
							v47 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
								F_UnlockReleaseBuffer(m, v51)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v24, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										if v50 != int32(1) {
											v65 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
											v69 = int32(0)
											m.G0 = v8 + int32(32)
											return v69
										} else {
											v59 = F_Int64GetDatum(m, v49)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v69 = v59
												m.G0 = v8 + int32(32)
												return v69
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
func F_pg_size_pretty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v13 = v11 >> (uint(int64(63)) % 64)
	if base.Ui64(v11^v13-v13) <= base.Ui64(int64(10239)) {
		v64 = int32(_a_F_pg_size_pretty_0)
		v65 = v11
		v68 = int32(_a_F_pg_size_pretty_1)
	} else {
		v21 = base.I64_div_s(v11, int64(512))
		v23 = v21 >> (uint(int64(63)) % 64)
		if base.Ui64(v21^v23-v23) < base.Ui64(int64(20479)) {
			v64 = int32(_a_F_pg_size_pretty_2)
			v65 = v21
			v68 = int32(_a_F_pg_size_pretty_3)
		} else {
			v31 = base.I64_div_s(v11, int64(524288))
			v33 = v31 >> (uint(int64(63)) % 64)
			if base.Ui64(v31^v33-v33) < base.Ui64(int64(20479)) {
				v64 = int32(_a_F_pg_size_pretty_4)
				v65 = v31
				v68 = int32(_a_F_pg_size_pretty_5)
			} else {
				v41 = base.I64_div_s(v11, int64(536870912))
				v43 = v41 >> (uint(int64(63)) % 64)
				if base.Ui64(v41^v43-v43) < base.Ui64(int64(20479)) {
					v64 = int32(_a_F_pg_size_pretty_6)
					v65 = v41
					v68 = int32(_a_F_pg_size_pretty_7)
				} else {
					v51 = base.I64_div_s(v11, int64(549755813888))
					v53 = v51 >> (uint(int64(63)) % 64)
					if base.Ui64(v51^v53-v53) < base.Ui64(int64(20479)) {
						v64 = int32(_a_F_pg_size_pretty_8)
						v65 = v51
						v68 = int32(_a_F_pg_size_pretty_9)
					} else {
						v61 = base.I64_div_s(v11, int64(562949953421312))
						v64 = int32(_a_F_pg_size_pretty_10)
						v65 = v61
						v68 = int32(_a_F_pg_size_pretty_11)
					}
				}
			}
		}
	}
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)))
	if v69 == int32(1) {
		v78 = base.I64_div_s(v65>>(uint(int64(63))%64)|int64(1)+v65, int64(2))
		v79 = v78
	} else {
		v79 = v65
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v79
	v83 = v8 + int32(16)
	v86 = F_pg_snprintf(m, v83, int32(64), int32(_a_F_pg_size_pretty_12), v8)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		return int32(0)
	} else {
		v90 = F_cstring_to_text(m, v83)
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(80)
			return v90
		}
	}
}
func F_pg_sjis_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = int32(1)
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v5+int32(95))&int32(255)) < base.Ui32(int32(63)) {
		v12 = v2
	} else {
		v12 = int32(2)
	}
	if int32(0) <= v5 {
		v15 = v2
	} else {
		v15 = v12
	}
	return v15
}
func F_pg_snapshot_in(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = v13 + int32(20)
	v21 = F_strtox_2(m, v16, v18, int32(10), int64(-1))
	mBase = m.M
	goto L1
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v23 != int32(58) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v13 + int32(48)
	return v154
L3:
	;
	v135 = int32(0)
	v136 = F_errsave_start(m, v15)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L26
	}
L4:
	;
	v30 = F_strtox_2(m, v22+int32(1), v18, int32(10), int64(-1))
	mBase = m.M
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if base.B2i32(base.I32_wrap_i64(v21) == int32(0))|base.B2i32(v35 != int32(58))|(base.B2i32(v30&int64(4294967295) == int64(0))|base.B2i32(base.Ui64(v30) < base.Ui64(v21))) != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v50 = F_makeStringInfo(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v54 = int32(24)
	F_appendBinaryStringInfo(m, v50, v13+v54, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
	F_pfree(m, v50)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L25
	}
L11:
	;
	v67 = v34 + int32(1)
	v73 = int64(0)
	goto L12
L12:
	;
	v78 = F_strtox_2(m, v67, v13+int32(20), int32(10), int64(-1))
	mBase = m.M
	goto L14
L13:
	;
	goto L10
L14:
	;
	if base.B2i32(base.Ui64(v78) < base.Ui64(v73))|base.B2i32(base.Ui64(v78) < base.Ui64(v21))|base.B2i32(base.Ui64(v30) <= base.Ui64(v78)) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v78 != v73 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v78
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v88 + int32(1)
	F_appendBinaryStringInfo(m, v50, v13+int32(24), int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v98 != int32(44) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	if v98 == int32(0) {
		goto L10
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v105 != 0 {
		v67 = v84 + int32(1)
		v73 = v78
		goto L12
	} else {
		goto L24
	}
L23:
	;
	goto L3
L24:
	;
	goto L13
L25:
	;
	v154 = v116
	goto L2
L26:
	;
	if v136 == int32(0) {
		v154 = v135
		goto L2
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_pg_snapshot_in_0)
	F_errmsg(m, int32(_a_F_pg_snapshot_in_1), v13)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errsave_finish(m, v15, int32(_a_F_pg_snapshot_in_2), int32(324), int32(_a_F_pg_snapshot_in_3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v154 = v135
	goto L2
}
func F_pg_snapshot_xip(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v23 = F_MemoryContextAlloc(m, v19, int32(base.Ui32(v20)>>(uint(int32(2))%32)))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v27 = int32(base.Ui32(v25) >> (uint(int32(2)) % 32))
					if v27 != 0 {
						base.MemoryCopy(m, v23, v13, v27)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v23
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
					v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+4)))
					if base.Ui64(v36) < base.Ui64(v38) {
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v37+base.I32_wrap_i64(v36)<<(uint(int32(3))%32))+24))
						*(*int64)(unsafe.Add(mBase, uint32(v35))) = v36 + int64(1)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(1)
						v51 = F_Int64GetDatum(m, v44)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							return v51
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
							v59 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
		v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+4)))
		if base.Ui64(v36) < base.Ui64(v38) {
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v37+base.I32_wrap_i64(v36)<<(uint(int32(3))%32))+24))
			*(*int64)(unsafe.Add(mBase, uint32(v35))) = v36 + int64(1)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(1)
			v51 = F_Int64GetDatum(m, v44)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				return v51
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
				v59 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
				return int32(0)
			}
		}
	}
}
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = int64(0)
	if l1 != 0 {
		v15 = l0
	} else {
		v15 = v8 + int32(7)
	}
	v16 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v16) {
		v19 = v16
	} else {
		v19 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15 + v19 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v15
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v26)
	F_dopr(m, v8+int32(8), l2, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v35 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v35)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
		m.G0 = v8 + int32(32)
		if v39 != 0 {
			v46 = int32(-1)
		} else {
			v46 = v38 + (v34 - v37)
		}
		return v46
	}
}
func F_pg_stop_making_pinned_objects(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_superuser(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_pg_stop_making_pinned_objects_0)
					F_errmsg(m, int32(_a_F_pg_stop_making_pinned_objects_1), v6)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stop_making_pinned_objects_2), int32(730), int32(_a_F_pg_stop_making_pinned_objects_0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v31 = m.G0
			v33 = v31 - int32(16)
			m.G0 = v33
			v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[0])))
			if v36 != int32(1) {
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[1]))
				v44 = F_LWLockAcquire(m, v40+int32(256), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[2]))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					if base.Ui32(int32(_a_F_pg_stop_making_pinned_objects_3)) <= base.Ui32(v48) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[2]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_pg_stop_making_pinned_objects_4)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v85
							F_errmsg_internal(m, int32(_a_F_pg_stop_making_pinned_objects_5), v33)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_stop_making_pinned_objects_6), int32(634), int32(_a_F_pg_stop_making_pinned_objects_7))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(_a_F_pg_stop_making_pinned_objects_4)
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(0)
						v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[1]))
						F_LWLockRelease(m, v58+int32(256))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							m.G0 = v33 + int32(16)
							m.G0 = v6 + int32(16)
							return int32(0)
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pg_stop_making_pinned_objects_8), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stop_making_pinned_objects_6), int32(627), int32(_a_F_pg_stop_making_pinned_objects_7))
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
			}
		}
	}
}
func F_pg_strerror_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_pg_strerror_r_0)
	if base.Ui32(int32(153)) < base.Ui32(l0) {
		v24 = v10
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v453
L2:
	;
	v453 = l1
	goto L1
L3:
	;
	switch l0 - int32(1) {
	case 0:
		v453 = int32(_a_F_pg_strerror_r_1)
		goto L1
	case 1:
		goto L172
	case 2:
		goto L171
	case 3:
		goto L170
	case 4:
		goto L169
	case 5:
		goto L168
	case 6:
		goto L167
	case 7:
		goto L166
	case 8:
		goto L165
	case 9:
		goto L164
	default:
		goto L112
	case 11:
		goto L163
	case 12:
		goto L162
	case 13:
		goto L161
	case 14:
		goto L160
	case 15:
		goto L159
	case 17:
		goto L158
	case 19:
		goto L157
	case 20:
		goto L156
	case 21:
		goto L155
	case 22:
		goto L153
	case 23:
		goto L152
	case 25:
		goto L151
	case 26:
		goto L150
	case 27:
		goto L149
	case 28:
		goto L148
	case 29:
		goto L147
	case 30:
		goto L146
	case 31:
		goto L145
	case 32:
		goto L144
	case 33:
		goto L143
	case 34:
		goto L142
	case 36:
		goto L141
	case 37:
		goto L140
	case 38:
		goto L139
	case 39:
		goto L138
	case 40:
		goto L137
	case 41:
		goto L136
	case 42:
		goto L135
	case 43:
		goto L134
	case 44:
		goto L133
	case 47:
		goto L132
	case 50:
		goto L131
	case 51:
		goto L130
	case 52:
		goto L129
	case 53:
		goto L128
	case 54:
		goto L127
	case 56:
		goto L126
	case 58:
		goto L124
	case 59:
		goto L123
	case 60:
		goto L122
	case 62:
		goto L121
	case 63:
		goto L120
	case 65:
		goto L119
	case 67:
		goto L118
	case 68:
		goto L117
	case 70:
		goto L116
	case 72:
		goto L115
	case 73:
		goto L114
	case 74:
		goto L113
	case 137:
		goto L125
	case 141:
		goto L154
	}
L4:
	;
	if v373|base.B2i32(l1 == int32(0)) != 0 {
		goto L3
	} else {
		goto L109
	}
L5:
	;
	v25 = F_strlen(m, v24)
	mBase = m.M
	if base.Ui32(int32(256)) <= base.Ui32(v25) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if l0 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v24 = v20 + int32(_a_F_pg_strerror_r_2)
	goto L5
L8:
	;
	v20 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_strerror_r[0]))))
	if v17 == int32(0) {
		v24 = v10
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v20 = v17
	goto L7
L12:
	;
	goto L17
L13:
	;
	goto L14
L14:
	;
	v202 = v25 + int32(1)
	if base.Ui32(int32(512)) <= base.Ui32(v202) {
		goto L63
	} else {
		goto L64
	}
L15:
	;
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+255)) = uint8(v198)
	v373 = int32(68)
	goto L4
L17:
	;
	goto L18
L18:
	;
	v35 = l1 + int32(255)
	if (l1^v24)&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if base.Ui32(v167) < base.Ui32(v35) {
		goto L56
	} else {
		goto L57
	}
L23:
	;
	if l1&int32(3) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v35) < base.Ui32(int32(4)) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v71 = v35 & int32(-4)
	if base.Ui32(v35) < base.Ui32(int32(64)) {
		v121 = v65
		v122 = v66
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v65 = v24
	v66 = l1
	goto L26
L28:
	;
	goto L29
L29:
	;
	goto L31
L31:
	;
	goto L32
L32:
	;
	v48 = v24
	v49 = l1
	goto L33
L33:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v53)
	v55 = int32(1)
	v56 = v48 + v55
	v58 = v49 + v55
	if v58&int32(3) == int32(0) {
		v65 = v56
		v66 = v58
		goto L26
	} else {
		goto L35
	}
L34:
	;
	v65 = v56
	v66 = v58
	goto L26
L35:
	;
	if base.Ui32(v58) < base.Ui32(v35) {
		v48 = v56
		v49 = v58
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if base.Ui32(v71) <= base.Ui32(v122) {
		v166 = v121
		v167 = v122
		goto L22
	} else {
		goto L43
	}
L38:
	;
	v75 = v71 + int32(-64)
	if base.Ui32(v75) < base.Ui32(v66) {
		v121 = v65
		v122 = v66
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v78 = v65
	v79 = v66
	goto L40
L40:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+36)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+40)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+48)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+52)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v78)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+56)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+60)) = v113
	v115 = int32(-64)
	v116 = v78 - v115
	v118 = v79 - v115
	if base.Ui32(v118) <= base.Ui32(v75) {
		v78 = v116
		v79 = v118
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v121 = v116
	v122 = v118
	goto L37
L42:
	;
	goto L41
L43:
	;
	v128 = v121
	v129 = v122
	goto L44
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v133
	v135 = int32(4)
	v136 = v128 + v135
	v138 = v129 + v135
	if base.Ui32(v138) < base.Ui32(v71) {
		v128 = v136
		v129 = v138
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v166 = v136
	v167 = v138
	goto L22
L46:
	;
	goto L45
L47:
	;
	v166 = v24
	v167 = l1
	goto L22
L48:
	;
	goto L49
L49:
	;
	goto L51
L51:
	;
	goto L52
L52:
	;
	v147 = v24
	v148 = l1
	goto L53
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v152)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)) = uint8(v154)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)) = uint8(v156)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)) = uint8(v158)
	v160 = int32(4)
	v161 = v147 + v160
	v163 = v148 + v160
	if base.Ui32(v163) <= base.Ui32(v35-int32(4)) {
		v147 = v161
		v148 = v163
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v166 = v161
	v167 = v163
	goto L22
L55:
	;
	goto L54
L56:
	;
	v173 = v166
	v174 = v167
	goto L59
L57:
	;
	goto L58
L58:
	;
	goto L15
L59:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v178)
	v180 = int32(1)
	v183 = v174 + v180
	if v183 != v35 {
		v173 = v173 + v180
		v174 = v183
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	goto L60
L62:
	;
	v373 = int32(0)
	goto L4
L63:
	;
	if v202 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v209 = l1 + v202
	if (l1^v24)&int32(3) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	base.MemoryCopy(m, l1, v24, v202)
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L62
L69:
	;
	if base.Ui32(v341) < base.Ui32(v209) {
		goto L103
	} else {
		goto L104
	}
L70:
	;
	if l1&int32(3) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(v209) < base.Ui32(int32(4)) {
		goto L94
	} else {
		goto L95
	}
L73:
	;
	v245 = v209 & int32(-4)
	if base.Ui32(v209) < base.Ui32(int32(64)) {
		v295 = v239
		v296 = v240
		goto L84
	} else {
		goto L85
	}
L74:
	;
	v239 = v24
	v240 = l1
	goto L73
L75:
	;
	goto L76
L76:
	;
	if v202 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v239 = v24
	v240 = l1
	goto L73
L78:
	;
	goto L79
L79:
	;
	v222 = v24
	v223 = l1
	goto L80
L80:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v227)
	v229 = int32(1)
	v230 = v222 + v229
	v232 = v223 + v229
	if v232&int32(3) == int32(0) {
		v239 = v230
		v240 = v232
		goto L73
	} else {
		goto L82
	}
L81:
	;
	v239 = v230
	v240 = v232
	goto L73
L82:
	;
	if base.Ui32(v232) < base.Ui32(v209) {
		v222 = v230
		v223 = v232
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if base.Ui32(v245) <= base.Ui32(v296) {
		v340 = v295
		v341 = v296
		goto L69
	} else {
		goto L90
	}
L85:
	;
	v249 = v245 + int32(-64)
	if base.Ui32(v249) < base.Ui32(v240) {
		v295 = v239
		v296 = v240
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v252 = v239
	v253 = v240
	goto L87
L87:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+12)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+16)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+20)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v252)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+28)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+32)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+36)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v252)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+40)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+44)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v252)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+48)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v252)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+52)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v252)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+56)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v252)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+60)) = v287
	v289 = int32(-64)
	v290 = v252 - v289
	v292 = v253 - v289
	if base.Ui32(v292) <= base.Ui32(v249) {
		v252 = v290
		v253 = v292
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v295 = v290
	v296 = v292
	goto L84
L89:
	;
	goto L88
L90:
	;
	v302 = v295
	v303 = v296
	goto L91
L91:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = v307
	v309 = int32(4)
	v310 = v302 + v309
	v312 = v303 + v309
	if base.Ui32(v312) < base.Ui32(v245) {
		v302 = v310
		v303 = v312
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v340 = v310
	v341 = v312
	goto L69
L93:
	;
	goto L92
L94:
	;
	v340 = v24
	v341 = l1
	goto L69
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(v202) < base.Ui32(int32(4)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v340 = v24
	v341 = l1
	goto L69
L98:
	;
	goto L99
L99:
	;
	v321 = v24
	v322 = l1
	goto L100
L100:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	*(*uint8)(unsafe.Add(mBase, uint32(v322))) = uint8(v326)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+1)) = uint8(v328)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+2)) = uint8(v330)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+3)) = uint8(v332)
	v334 = int32(4)
	v335 = v321 + v334
	v337 = v322 + v334
	if base.Ui32(v337) <= base.Ui32(v209-int32(4)) {
		v321 = v335
		v322 = v337
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v340 = v335
	v341 = v337
	goto L69
L102:
	;
	goto L101
L103:
	;
	v347 = v340
	v348 = v341
	goto L106
L104:
	;
	goto L105
L105:
	;
	goto L62
L106:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	*(*uint8)(unsafe.Add(mBase, uint32(v348))) = uint8(v352)
	v354 = int32(1)
	v357 = v348 + v354
	if v357 != v209 {
		v347 = v347 + v354
		v348 = v357
		goto L106
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	goto L107
L109:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v377 == int32(63) {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	if v377 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	goto L3
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	v447 = F_pg_snprintf(m, l1, int32(256), int32(_a_F_pg_strerror_r_3), v8)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L173
	} else {
		goto L174
	}
L113:
	;
	v453 = int32(_a_F_pg_strerror_r_4)
	goto L1
L114:
	;
	v453 = int32(_a_F_pg_strerror_r_5)
	goto L1
L115:
	;
	v453 = int32(_a_F_pg_strerror_r_6)
	goto L1
L116:
	;
	v453 = int32(_a_F_pg_strerror_r_7)
	goto L1
L117:
	;
	v453 = int32(_a_F_pg_strerror_r_8)
	goto L1
L118:
	;
	v453 = int32(_a_F_pg_strerror_r_9)
	goto L1
L119:
	;
	v453 = int32(_a_F_pg_strerror_r_10)
	goto L1
L120:
	;
	v453 = int32(_a_F_pg_strerror_r_11)
	goto L1
L121:
	;
	v453 = int32(_a_F_pg_strerror_r_12)
	goto L1
L122:
	;
	v453 = int32(_a_F_pg_strerror_r_13)
	goto L1
L123:
	;
	v453 = int32(_a_F_pg_strerror_r_14)
	goto L1
L124:
	;
	v453 = int32(_a_F_pg_strerror_r_15)
	goto L1
L125:
	;
	v453 = int32(_a_F_pg_strerror_r_16)
	goto L1
L126:
	;
	v453 = int32(_a_F_pg_strerror_r_17)
	goto L1
L127:
	;
	v453 = int32(_a_F_pg_strerror_r_18)
	goto L1
L128:
	;
	v453 = int32(_a_F_pg_strerror_r_19)
	goto L1
L129:
	;
	v453 = int32(_a_F_pg_strerror_r_20)
	goto L1
L130:
	;
	v453 = int32(_a_F_pg_strerror_r_21)
	goto L1
L131:
	;
	v453 = int32(_a_F_pg_strerror_r_22)
	goto L1
L132:
	;
	v453 = int32(_a_F_pg_strerror_r_23)
	goto L1
L133:
	;
	v453 = int32(_a_F_pg_strerror_r_24)
	goto L1
L134:
	;
	v453 = int32(_a_F_pg_strerror_r_25)
	goto L1
L135:
	;
	v453 = int32(_a_F_pg_strerror_r_26)
	goto L1
L136:
	;
	v453 = int32(_a_F_pg_strerror_r_27)
	goto L1
L137:
	;
	v453 = int32(_a_F_pg_strerror_r_28)
	goto L1
L138:
	;
	v453 = int32(_a_F_pg_strerror_r_29)
	goto L1
L139:
	;
	v453 = int32(_a_F_pg_strerror_r_30)
	goto L1
L140:
	;
	v453 = int32(_a_F_pg_strerror_r_31)
	goto L1
L141:
	;
	v453 = int32(_a_F_pg_strerror_r_32)
	goto L1
L142:
	;
	v453 = int32(_a_F_pg_strerror_r_33)
	goto L1
L143:
	;
	v453 = int32(_a_F_pg_strerror_r_34)
	goto L1
L144:
	;
	v453 = int32(_a_F_pg_strerror_r_35)
	goto L1
L145:
	;
	v453 = int32(_a_F_pg_strerror_r_36)
	goto L1
L146:
	;
	v453 = int32(_a_F_pg_strerror_r_37)
	goto L1
L147:
	;
	v453 = int32(_a_F_pg_strerror_r_38)
	goto L1
L148:
	;
	v453 = int32(_a_F_pg_strerror_r_39)
	goto L1
L149:
	;
	v453 = int32(_a_F_pg_strerror_r_40)
	goto L1
L150:
	;
	v453 = int32(_a_F_pg_strerror_r_41)
	goto L1
L151:
	;
	v453 = int32(_a_F_pg_strerror_r_42)
	goto L1
L152:
	;
	v453 = int32(_a_F_pg_strerror_r_43)
	goto L1
L153:
	;
	v453 = int32(_a_F_pg_strerror_r_44)
	goto L1
L154:
	;
	v453 = int32(_a_F_pg_strerror_r_45)
	goto L1
L155:
	;
	v453 = int32(_a_F_pg_strerror_r_46)
	goto L1
L156:
	;
	v453 = int32(_a_F_pg_strerror_r_47)
	goto L1
L157:
	;
	v453 = int32(_a_F_pg_strerror_r_48)
	goto L1
L158:
	;
	v453 = int32(_a_F_pg_strerror_r_49)
	goto L1
L159:
	;
	v453 = int32(_a_F_pg_strerror_r_50)
	goto L1
L160:
	;
	v453 = int32(_a_F_pg_strerror_r_51)
	goto L1
L161:
	;
	v453 = int32(_a_F_pg_strerror_r_52)
	goto L1
L162:
	;
	v453 = int32(_a_F_pg_strerror_r_53)
	goto L1
L163:
	;
	v453 = int32(_a_F_pg_strerror_r_54)
	goto L1
L164:
	;
	v453 = int32(_a_F_pg_strerror_r_55)
	goto L1
L165:
	;
	v453 = int32(_a_F_pg_strerror_r_56)
	goto L1
L166:
	;
	v453 = int32(_a_F_pg_strerror_r_57)
	goto L1
L167:
	;
	v453 = int32(_a_F_pg_strerror_r_58)
	goto L1
L168:
	;
	v453 = int32(_a_F_pg_strerror_r_59)
	goto L1
L169:
	;
	v453 = int32(_a_F_pg_strerror_r_60)
	goto L1
L170:
	;
	v453 = int32(_a_F_pg_strerror_r_61)
	goto L1
L171:
	;
	v453 = int32(_a_F_pg_strerror_r_62)
	goto L1
L172:
	;
	v453 = int32(_a_F_pg_strerror_r_63)
	goto L1
L173:
	;
	return int32(0)
L174:
	;
	goto L2
}
func F_pg_strlower(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn13965(m, l0, l1, l2, l3, l4, int32(0), int32(_a_F_pg_strlower_0), int32(1284))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_pg_sync_replication_slots(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	F_CheckSlotPermissions(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[0])))
		if v14 == int32(1) {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[1]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
			v22 = base.B2i32(v20 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[0])) = uint8(v22)
			v24 = v22
		} else {
			v24 = int32(0)
		}
		if v24 != 0 {
			v26 = F_ValidateSlotSyncParams(m, int32(21))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_load_file(m, int32(_a_F_pg_sync_replication_slots_0), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = m.G0
					v34 = v32 - int32(16)
					m.G0 = v34
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[2]))
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
					v41 = m.T0[v40].(func(*base.Module, int32) int32)(m, v37)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						if v41 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_pg_sync_replication_slots_1)
									*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(_a_F_pg_sync_replication_slots_2)
									F_errmsg(m, int32(_a_F_pg_sync_replication_slots_3), v34)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_sync_replication_slots_4), int32(1052), int32(_a_F_pg_sync_replication_slots_5))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							m.G0 = v34 + int32(16)
							v68 = v6 + int32(28)
							F_initStringInfo(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[4]))
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
								if v73 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v72
									F_appendStringInfo(m, v68, int32(_a_F_pg_sync_replication_slots_6), v6+int32(16))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[2]))
										v87 = int32(0)
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v86, v87, v87, v87, v90, v6+int32(44))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											if v96 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
														*(*int32)(unsafe.Add(mBase, uint32(v6))) = v138
														v140 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v140
														F_errmsg(m, int32(_a_F_pg_sync_replication_slots_7), v6)
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_sync_replication_slots_8), int32(929), int32(_a_F_pg_sync_replication_slots_9))
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
												F_pfree(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													F_SyncReplicationSlots(m, v96)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+64))
														m.T0[v107].(func(*base.Module, int32))(m, v96)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int32(0)
														} else {
															m.G0 = v6 + int32(48)
															return int32(0)
														}
													}
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, v6+int32(28), int32(_a_F_pg_sync_replication_slots_10))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[2]))
										v87 = int32(0)
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v86, v87, v87, v87, v90, v6+int32(44))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											if v96 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
														*(*int32)(unsafe.Add(mBase, uint32(v6))) = v138
														v140 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v140
														F_errmsg(m, int32(_a_F_pg_sync_replication_slots_7), v6)
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_sync_replication_slots_8), int32(929), int32(_a_F_pg_sync_replication_slots_9))
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
												F_pfree(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													F_SyncReplicationSlots(m, v96)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int32(0)
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+64))
														m.T0[v107].(func(*base.Module, int32))(m, v96)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int32(0)
														} else {
															m.G0 = v6 + int32(48)
															return int32(0)
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_sync_replication_slots_11), int32(0))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_sync_replication_slots_8), int32(906), int32(_a_F_pg_sync_replication_slots_9))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
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
func F_pg_table_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_RelationIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_tablespace_location(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v4 = m.G0
	v6 = v4 - int32(2208)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tablespace_location[0]))
	if v8 != 0 {
		v11 = v8
	} else {
		v11 = v10
	}
	if base.Ui32(v11-int32(1663)) <= base.Ui32(int32(1)) {
		v17 = F_cstring_to_text(m, int32(_a_F_pg_tablespace_location_0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v63 = v17
			m.G0 = v6 + int32(2208)
			return v63
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(_a_F_pg_tablespace_location_1)
		v25 = v6 + int32(1184)
		v30 = F_pg_snprintf(m, v25, int32(1024), int32(_a_F_pg_tablespace_location_2), v6+int32(48))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v36 = F___fstatat(m, int32(-100), v25, v6-int32(-64), int32(256))
			mBase = m.M
			if v36 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v6 + int32(1184)
						F_errmsg(m, int32(_a_F_pg_tablespace_location_3), v6)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_tablespace_location_4), int32(341), int32(_a_F_pg_tablespace_location_5))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
				if v39&int32(_a_F_pg_tablespace_location_6) != int32(_a_F_pg_tablespace_location_7) {
					v44 = F_cstring_to_text(m, v25)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v63 = v44
						m.G0 = v6 + int32(2208)
						return v63
					}
				} else {
					v49 = v6 + int32(160)
					v51 = F_readlink(m, v6+int32(1184), v49, int32(1024))
					mBase = m.M
					if v51 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v6 + int32(1184)
								F_errmsg(m, int32(_a_F_pg_tablespace_location_8), v6+int32(16))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_tablespace_location_4), int32(355), int32(_a_F_pg_tablespace_location_5))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(1024)) <= base.Ui32(v51) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v6 + int32(1184)
									F_errmsg(m, int32(_a_F_pg_tablespace_location_9), v6+int32(32))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_tablespace_location_4), int32(360), int32(_a_F_pg_tablespace_location_5))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v57 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v51+v49))) = uint8(v57)
							v59 = F_cstring_to_text(m, v49)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v63 = v59
								m.G0 = v6 + int32(2208)
								return v63
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_timezone_names(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int64
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int64
	_ = v230
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v248 int64
	_ = v248
	var v251 int64
	_ = v251
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v2
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = F_palloc0(m, int32(_a_F_pg_timezone_names_0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_names[0])))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_get_share_path(m, int32(_a_F_pg_timezone_names_1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v163 = F_pstrdup(m, int32(_a_F_pg_timezone_names_1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L39
	}
L7:
	;
	v33 = int32(_a_F_pg_timezone_names_1)
	v34 = F_strlen(m, v33)
	mBase = m.M
	v36 = v34 + v33
	v37 = int32(_a_F_pg_timezone_names_2)
	v39 = int32(1024) - v34
	if v39 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_names[0])) = uint8(v159)
	goto L6
L9:
	;
	v155 = F_strlen(m, v151)
	mBase = m.M
	goto L8
L10:
	;
	v151 = v37
	goto L9
L11:
	;
	goto L12
L12:
	;
	v45 = v39 - int32(1)
	if (v36^v37)&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v148)
	v151 = v144
	goto L9
L14:
	;
	v129 = v124
	v130 = v125
	v131 = v126
	goto L35
L15:
	;
	if v119 == int32(0) {
		v144 = v117
		v145 = v118
		goto L13
	} else {
		goto L34
	}
L16:
	;
	v117 = v37
	v118 = v36
	v119 = v45
	goto L15
L17:
	;
	goto L18
L18:
	;
	v49 = int32(0)
	if int32(1)|base.B2i32(v45 == v49) == v49 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v85 == int32(0) {
		v144 = v82
		v145 = v83
		goto L13
	} else {
		goto L28
	}
L20:
	;
	v61 = v37
	v62 = v36
	v63 = v45
	goto L23
L21:
	;
	goto L22
L22:
	;
	v82 = v37
	v83 = v36
	v84 = v45
	v85 = base.B2i32(v45 != v49)
	goto L19
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v65)
	if v65 == int32(0) {
		v124 = v61
		v125 = v62
		v126 = v63
		goto L14
	} else {
		goto L25
	}
L24:
	;
	v82 = v76
	v83 = v70
	v84 = v72
	v85 = v74
	goto L19
L25:
	;
	v69 = int32(1)
	v70 = v62 + v69
	v72 = v63 - v69
	v73 = int32(0)
	v74 = base.B2i32(v72 != v73)
	v76 = v61 + v69
	if v76&int32(3) == v73 {
		v82 = v76
		v83 = v70
		v84 = v72
		v85 = v74
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if v72 != 0 {
		v61 = v76
		v62 = v70
		v63 = v72
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.B2i32(v88 == int32(0))|base.B2i32(base.Ui32(v84) < base.Ui32(int32(4))) != 0 {
		v117 = v82
		v118 = v83
		v119 = v84
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v95 = v82
	v96 = v83
	v97 = v84
	goto L30
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v103 = int32(-2139062144)
	if (int32(16843008)-v100|v100)&v103 != v103 {
		v124 = v95
		v125 = v96
		v126 = v97
		goto L14
	} else {
		goto L32
	}
L31:
	;
	v117 = v111
	v118 = v109
	v119 = v113
	goto L15
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v100
	v108 = int32(4)
	v109 = v96 + v108
	v111 = v95 + v108
	v113 = v97 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v95 = v111
		v96 = v109
		v97 = v113
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v124 = v117
	v125 = v118
	v126 = v119
	goto L14
L35:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v133)
	if v133 == int32(0) {
		v144 = v129
		v145 = v130
		goto L13
	} else {
		goto L37
	}
L36:
	;
	v144 = v140
	v145 = v138
	goto L13
L37:
	;
	v137 = int32(1)
	v138 = v130 + v137
	v140 = v129 + v137
	v142 = v131 - v137
	if v142 != 0 {
		v129 = v140
		v130 = v138
		v131 = v142
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v165 = F_strlen(m, v163)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v165 + int32(1)
	v172 = F_AllocateDir(m, v163)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	if v172 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	m.G0 = v21 + int32(16)
	v195 = F_pg_tzenumerate_next(m, v24)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v163
	F_errmsg(m, int32(_a_F_pg_timezone_names_3), v21)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_names_4), int32(409), int32(_a_F_pg_timezone_names_5))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	if v195 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v198 = v9 + int32(16)
	v199 = v195
	goto L52
L50:
	;
	goto L51
L51:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) <= v287 {
		goto L74
	} else {
		goto L75
	}
L52:
	;
	v206 = *(*int64)(unsafe.Add(mBase, _c_F_pg_timezone_names[1]))
	v215 = F_timestamp2tm(m, v206, v9+int32(88), v9+int32(44), v9+int32(40), v9+int32(36), v199)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	v279 = F_pg_tzenumerate_next(m, v24)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L72
	}
L55:
	;
	if v215 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v217 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v218 = F_strlen(m, v217)
	mBase = m.M
	if base.Ui32(int32(31)) < base.Ui32(v218) {
		goto L54
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v221 = F_cstring_to_text(m, v199)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v221
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v224 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v226 = v224
	goto L64
L63:
	;
	v226 = int32(_a_F_pg_timezone_names_6)
	goto L64
L64:
	;
	v227 = F_cstring_to_text(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v227
	v230 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v198)+8)) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v198))) = v230
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = base.I64_extend_i32_s(int32(0)-v235) * int64(1000000)
	v242 = v9 + int32(8)
	v244 = F_palloc(m, int32(16))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v247 = int64(*(*int32)(unsafe.Add(mBase, uint32(v242)+12)))
	v248 = int64(*(*int32)(unsafe.Add(mBase, uint32(v242)+16)))
	v251 = v247 + v248*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v251-int64(2147483648)) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v244
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+108)) = base.B2i32(int32(0) < v265)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v269, v270, v9+int32(96), v9+int32(92))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L71
	}
L68:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v244)+12)) = uint32(v251)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+8)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = v259
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L54
L72:
	;
	if v279 != 0 {
		v199 = v279
		goto L52
	} else {
		goto L73
	}
L73:
	;
	goto L53
L74:
	;
	v294 = v287
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_pfree(m, v24)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v294<<(uint(int32(2))%32))))
	F_FreeDir(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(48)+v306<<(uint(int32(2))%32))))
	F_pfree(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v315 = v313 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v315
	if int32(0) <= v315 {
		v294 = v315
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	m.G0 = v9 + int32(112)
	return int32(0)
}
func F_pg_ts_config_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSConfigIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_ts_template_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSTemplateIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_tzset_offset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	v12 = l0 >> (uint(int32(31)) % 32)
	v14 = l0 ^ v12 - v12
	v16 = base.I32_div_s(v14, int32(3600))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v16
	v19 = v9 + int32(192)
	v24 = F_pg_snprintf(m, v19, int32(64), int32(_a_F_pg_tzset_offset_0), v9+int32(48))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v30 = v14 - v16*int32(3600)
		if v30 == int32(0) {
			v69 = v9 + int32(192)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
			v73 = v9 - int32(-64)
			if int32(0) < l0 {
				v79 = int32(_a_F_pg_tzset_offset_1)
			} else {
				v79 = int32(_a_F_pg_tzset_offset_2)
			}
			v80 = F_pg_snprintf(m, v73, int32(128), v79, v9)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = F_pg_tzset(m, v73)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(256)
					return v82
				}
			}
		} else {
			v33 = F_strlen(m, v19)
			mBase = m.M
			v36 = base.I32_div_s(base.I32_extend16_s(v30), int32(60))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = base.I32_extend16_s(v36)
			v45 = F_pg_snprintf(m, v33+v19, int32(64)-v33, int32(_a_F_pg_tzset_offset_3), v9+int32(32))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v49 = v30 - v36*int32(60)
				if v49&int32(_a_F_pg_tzset_offset_4) == int32(0) {
					v69 = v9 + int32(192)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
					v73 = v9 - int32(-64)
					if int32(0) < l0 {
						v79 = int32(_a_F_pg_tzset_offset_1)
					} else {
						v79 = int32(_a_F_pg_tzset_offset_2)
					}
					v80 = F_pg_snprintf(m, v73, int32(128), v79, v9)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = F_pg_tzset(m, v73)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(256)
							return v82
						}
					}
				} else {
					v54 = F_strlen(m, v19)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = base.I32_extend16_s(v49)
					v63 = F_pg_snprintf(m, v54+v19, int32(64)-v54, int32(_a_F_pg_tzset_offset_3), v9+int32(16))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v69 = v9 + int32(192)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
						v73 = v9 - int32(-64)
						if int32(0) < l0 {
							v79 = int32(_a_F_pg_tzset_offset_1)
						} else {
							v79 = int32(_a_F_pg_tzset_offset_2)
						}
						v80 = F_pg_snprintf(m, v73, int32(128), v79, v9)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = F_pg_tzset(m, v73)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(256)
								return v82
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_ultoa_n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v3 = int32(0)
	if l0 == v3 {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(1)
	} else {
		v19 = int32(1233)
		v24 = int32(base.Ui32((base.I32_clz(l0)^int32(31))*v19+v19) >> (uint(int32(12)) % 32))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_pg_ultoa_n[0])))
		v29 = v24 + base.B2i32(base.Ui32(v27) <= base.Ui32(l0))
		if base.Ui32(int32(_a_F_pg_ultoa_n_0)) <= base.Ui32(l0) {
			v33 = l0
			v35 = v3
			for {
				v42 = l1 + v29 - v35
				v43 = int32(4)
				v46 = base.I32_div_u_s(v33, int32(_a_F_pg_ultoa_n_0))
				v49 = v33 + v46*int32(-10000)
				v50 = int32(100)
				v51 = base.I32_div_u_s(v49, v50)
				v52 = int32(1)
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51<<(uint(v52)%32))+uint32(_c_F_pg_ultoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v42-v43))) = uint16(v54)
				v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v49-v51*v50)<<(uint(v52)%32))+uint32(_c_F_pg_ultoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v42-int32(2)))) = uint16(v63)
				v66 = v35 + v43
				if base.Ui32(int32(99999999)) < base.Ui32(v33) {
					v33 = v46
					v35 = v66
					continue
				} else {
					break
				}
				break
			}
			v69 = v46
			v71 = v66
		} else {
			v69 = l0
			v71 = v3
		}
		if base.Ui32(int32(100)) <= base.Ui32(v69) {
			v82 = int32(2)
			v84 = int32(_a_F_pg_ultoa_n_1)
			v86 = int32(100)
			v87 = base.I32_div_u_s(v69&v84, v86)
			v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v69-v87*v86)&v84<<(uint(int32(1))%32))+uint32(_c_F_pg_ultoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v29-v71-v82))) = uint16(v95)
			v99 = v87
			v100 = v71 | v82
		} else {
			v99 = v69
			v100 = v71
		}
		if base.Ui32(int32(10)) <= base.Ui32(v99) {
			v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99<<(uint(int32(1))%32))+uint32(_c_F_pg_ultoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v29-v100-int32(2)))) = uint16(v109)
			return v29
		} else {
			v113 = v99 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v113)
			return v29
		}
	}
}
func F_pg_usleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v46 float64
	_ = v46
	var v51 float64
	_ = v51
	var v57 int32
	_ = v57
	var v61 float64
	_ = v61
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v73 float64
	_ = v73
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v87 int32
	_ = v87
	var v90 float64
	_ = v90
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if int32(0) < l0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(_a_F_pg_usleep_0)
	v17 = base.I32_div_u_s(l0, v16)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = base.I64_extend_i32_u(v17)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = (l0 - v17*v16) * int32(1000)
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = int32(28)
	if v12 == int32(0) {
		v94 = v30
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v12 + int32(16)
	return
L4:
	;
	m.G0 = v28 + int32(16)
	v107 = int32(0) - v94
	if base.Ui32(int32(-4095)) <= base.Ui32(v107) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(int32(999999999)) < base.Ui32(v33) {
		v94 = v30
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if v36 < int64(0) {
		v94 = v30
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v46 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	v51 = v46
	goto L8
L8:
	;
	v57 = int32(0)
	v61 = v51
	goto L10
L9:
	;
	v94 = int32(0)
	goto L4
L10:
	;
	v67 = v57 << (uint(int32(3)) % 32)
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_pg_usleep[0])))
	if base.F64_eq(v68, float64(0)) != 0 {
		v84 = v61
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v90 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	if base.F64_lt(base.F64_sub(v90, v46), base.F64_add(base.F64_mul(base.F64_convert_i64_s(v36), float64(1000)), base.F64_div(base.F64_convert_i32_s(v33), float64(1e+06)))) != 0 {
		v51 = v90
		goto L8
	} else {
		goto L21
	}
L12:
	;
	v87 = v57 + int32(1)
	if v87 != int32(3) {
		v57 = v87
		v61 = v84
		goto L10
	} else {
		goto L20
	}
L13:
	;
	if base.F64_eq(v61, float64(0)) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v73 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_pg_usleep[0])))
	v77 = v73
	v78 = v76
	goto L16
L15:
	;
	v77 = v61
	v78 = v68
	goto L16
L16:
	;
	if base.F64_ge(v77, v78) == int32(0) {
		v84 = v77
		goto L12
	} else {
		goto L17
	}
L17:
	;
	F__emscripten_timeout(m, v57, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v84 = v77
	goto L12
L20:
	;
	goto L11
L21:
	;
	goto L9
L22:
	;
	goto L3
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_usleep[1])) = int32(0) - v107
	goto L25
L24:
	;
	goto L25
L25:
	;
	goto L22
}
func F_pg_utf8_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	if base.Ui32(l1) < base.Ui32(int32(16)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v326 <= int32(0) {
		v455 = v328
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v326 = l1
	v328 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = l0
	v29 = int32(11)
	v31 = l1
	goto L5
L5:
	;
	if v29 != int32(11) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	switch v252 {
	case 0:
		v326 = l1
		v328 = l0
		goto L1
	default:
		goto L16
	case 11:
		v305 = v254
		v308 = v256
		goto L15
	}
L7:
	;
	v253 = int32(16)
	v254 = v28 + v253
	v256 = v31 - v253
	if base.Ui32(int32(15)) < base.Ui32(v256) {
		v28 = v254
		v29 = v252
		v31 = v256
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	v141 = int32(2)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v144 = int32(255)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v123&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v124&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v125&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v126&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v127&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v128&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v129&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v130<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v131&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v132&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v133&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v134&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v135&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v136&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v137&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v252 = int32(base.Ui32(v143)>>(uint(int32(base.Ui32(v148)>>(uint(int32(base.Ui32(v153)>>(uint(int32(base.Ui32(v158)>>(uint(int32(base.Ui32(v163)>>(uint(int32(base.Ui32(v168)>>(uint(int32(base.Ui32(v173)>>(uint(int32(base.Ui32(v178)>>(uint(int32(base.Ui32(v181)>>(uint(int32(base.Ui32(v186)>>(uint(int32(base.Ui32(v191)>>(uint(int32(base.Ui32(v196)>>(uint(int32(base.Ui32(v201)>>(uint(int32(base.Ui32(v206)>>(uint(int32(base.Ui32(v211)>>(uint(int32(base.Ui32(v216)>>(uint(v29)%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)) & int32(31)
	goto L7
L9:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+14)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+13)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+12)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+11)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v123 = v50
	v124 = v51
	v125 = v52
	v126 = v53
	v127 = v54
	v128 = v55
	v129 = v56
	v130 = v57
	v131 = v58
	v132 = v59
	v133 = v60
	v134 = v61
	v135 = v62
	v136 = v63
	v137 = v64
	goto L8
L10:
	;
	goto L11
L11:
	;
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v66 = int64(48)
	v68 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v66) % 64)))
	v69 = int64(40)
	v71 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v69) % 64)))
	v72 = int64(32)
	v74 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v72) % 64)))
	v75 = int64(24)
	v77 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v75) % 64)))
	v78 = int64(16)
	v80 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v78) % 64)))
	v81 = int64(8)
	v83 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v81) % 64)))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v87 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(int64(56)) % 64)))
	v90 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v66) % 64)))
	v93 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v69) % 64)))
	v96 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v72) % 64)))
	v99 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v75) % 64)))
	v102 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v78) % 64)))
	v105 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v81) % 64)))
	v106 = base.I32_wrap_i64(v65)
	v107 = base.I32_wrap_i64(v84)
	if (v84|v65)&int64(-9187201950435737472) != int64(0) {
		v123 = v68
		v124 = v71
		v125 = v74
		v126 = v77
		v127 = v80
		v128 = v83
		v129 = v106
		v130 = v87
		v131 = v90
		v132 = v93
		v133 = v96
		v134 = v99
		v135 = v102
		v136 = v105
		v137 = v107
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v114 = int64(9187201950435737471)
	v119 = int64(-9187201950435737472)
	if (v84+v114)&(v65+v114)&v119 == v119 {
		v252 = int32(11)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v123 = v68
	v124 = v71
	v125 = v74
	v126 = v77
	v127 = v80
	v128 = v83
	v129 = v106
	v130 = v87
	v131 = v90
	v132 = v93
	v133 = v96
	v134 = v99
	v135 = v102
	v136 = v105
	v137 = v107
	goto L8
L14:
	;
	goto L6
L15:
	;
	v326 = v308
	v328 = v305
	goto L1
L16:
	;
	v261 = v254
	v264 = v256
	goto L17
L17:
	;
	v281 = int32(1)
	v282 = v264 + v281
	v284 = v261 - v281
	v285 = int32(*(*int8)(unsafe.Add(mBase, uint32(v284))))
	v287 = v285 & int32(255)
	if int32(0) <= v285 {
		v261 = v284
		v264 = v282
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v326 = v282
	v328 = v284
	goto L1
L19:
	;
	if base.B2i32(v287&int32(248) == int32(240))|base.B2i32(v287&int32(224) == int32(192)) != 0 {
		v305 = v284
		v308 = v282
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v287&int32(240) != int32(224) {
		v261 = v284
		v264 = v282
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	return v455 - l0
L23:
	;
	v350 = v326
	v352 = v328
	goto L24
L24:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v372 = base.I32_extend8_s(v371)
	if int32(0) <= v372 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v455 = v448
	goto L22
L26:
	;
	v448 = v447 + v352
	v449 = v350 - v447
	if int32(0) < v449 {
		v350 = v449
		v352 = v448
		goto L24
	} else {
		goto L60
	}
L27:
	;
	if v372 != 0 {
		v447 = int32(1)
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v371&int32(224) == int32(192) {
		v393 = int32(2)
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v455 = v352
	goto L22
L31:
	;
	if base.Ui32(v350) < base.Ui32(v393) {
		v455 = v352
		goto L22
	} else {
		goto L37
	}
L32:
	;
	if v371&int32(240) == int32(224) {
		v393 = int32(3)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v371&int32(248) == int32(240) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v392 = int32(4)
	goto L36
L35:
	;
	v392 = int32(1)
	goto L36
L36:
	;
	v393 = v392
	goto L31
L37:
	;
	v395 = int32(0)
	switch v393 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L43
	case 2:
		goto L44
	case 3:
		goto L45
	default:
		v444 = v395
		goto L39
	}
L38:
	;
	if v444 == int32(0) {
		v455 = v352
		goto L22
	} else {
		goto L59
	}
L39:
	;
	goto L38
L40:
	;
	v444 = base.B2i32(base.Ui32(v436&int32(255)) < base.Ui32(int32(245)))
	goto L39
L41:
	;
	if base.I32_extend8_s(v431) < int32(-62) {
		v444 = v395
		goto L39
	} else {
		goto L58
	}
L42:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v431 = v430
	goto L41
L43:
	;
	v404 = int32(*(*int8)(unsafe.Add(mBase, uint32(v352)+1)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	switch v405 - int32(224) {
	case 0:
		goto L52
	default:
		goto L48
	case 13:
		goto L51
	case 16:
		goto L50
	case 20:
		goto L49
	}
L44:
	;
	v401 = int32(*(*int8)(unsafe.Add(mBase, uint32(v352)+2)))
	if int32(-65) < v401 {
		v444 = v395
		goto L39
	} else {
		goto L47
	}
L45:
	;
	v398 = int32(*(*int8)(unsafe.Add(mBase, uint32(v352)+3)))
	if int32(-65) < v398 {
		v444 = v395
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L43
L48:
	;
	if v404 <= int32(-65) {
		v431 = v405
		goto L41
	} else {
		goto L57
	}
L49:
	;
	if int32(-113) < v404 {
		v444 = v395
		goto L39
	} else {
		goto L56
	}
L50:
	;
	if base.Ui32((v404-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v444 = v395
		goto L39
	} else {
		goto L55
	}
L51:
	;
	if int32(-97) < v404 {
		v444 = v395
		goto L39
	} else {
		goto L54
	}
L52:
	;
	v408 = int32(224)
	if base.Ui32(v408) <= base.Ui32((v404-int32(-64))&int32(255)) {
		v436 = v408
		goto L40
	} else {
		goto L53
	}
L53:
	;
	v444 = v395
	goto L39
L54:
	;
	v436 = int32(237)
	goto L40
L55:
	;
	v436 = int32(240)
	goto L40
L56:
	;
	v436 = int32(244)
	goto L40
L57:
	;
	v444 = v395
	goto L39
L58:
	;
	v436 = v431
	goto L40
L59:
	;
	v447 = v393
	goto L26
L60:
	;
	goto L25
}
