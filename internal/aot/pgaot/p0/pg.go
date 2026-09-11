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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v14 = v10 + int32(168)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v10)+176)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v15
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+32)) = uint16(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v12 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v15
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+88)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+92)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+116)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = v59
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v61
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v63
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = v65
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v67
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+122)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+124)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+125)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+127)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+140)) = v75
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+126)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+136)) = v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+129)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v81
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+130)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+152)) = v83
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+131)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+156)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v91
	if l3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+172)) = l3
	} else {
		v94 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)) = uint8(v94)
	}
	if l4 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = l4
	} else {
		v97 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v97)
	}
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+33)) = uint8(v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v104 = F_heap_form_tuple(m, v101, v10+int32(48), v10)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		return
	} else {
		F_CatalogTupleInsert(m, l0, v104)
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return
		} else {
			F_pfree(m, v104)
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return
			} else {
				m.G0 = v10 + int32(192)
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[195]))
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v7 = *(*int32)(unsafe.Add(mBase, _consts[641]))
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	v55 = int32(1794320)
	v56 = int32(1793680)
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
	v48 = v23 + int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 != 0 {
		v23 = v48
		v24 = v49
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
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int64
	_ = v510
	var v512 int64
	_ = v512
	var v524 int32
	_ = v524
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v718 int32
	_ = v718
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int64
	_ = v787
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1177 int32
	_ = v1177
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1610 int32
	_ = v1610
	var v1611 int64
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1625 int32
	_ = v1625
	var v1626 int64
	_ = v1626
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int64
	_ = v1636
	var v1638 int64
	_ = v1638
	var v1640 int64
	_ = v1640
	var v1642 int64
	_ = v1642
	var v1644 int64
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1669 int64
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1683 int32
	_ = v1683
	var v1684 int64
	_ = v1684
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int64
	_ = v1694
	var v1696 int64
	_ = v1696
	var v1698 int64
	_ = v1698
	var v1700 int64
	_ = v1700
	var v1702 int64
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1790 int32
	_ = v1790
	*(*int32)(unsafe.Add(mBase, _consts[172])) = int32(9)
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
	v18 = int32(913)
	v20 = m.G0
	v22 = v20 - int32(144)
	m.G0 = v22
	switch int32(915) {
	case 0, 2:
		v32 = v18
		goto L4
	default:
		goto L5
	}
L3:
	;
	v60 = int32(-2)
	v62 = m.G0
	v64 = v62 - int32(144)
	m.G0 = v64
	switch int32(0) {
	case 0, 2:
		v74 = v60
		goto L17
	default:
		goto L18
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v32
	F_sigemptyset(m, v22+int32(8))
	mBase = m.M
	goto L7
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[402])) = v18
	v32 = int32(4729)
	goto L4
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+136)) = int32(268435456)
	v44 = v22 + int32(4)
	goto L11
L9:
	;
	m.G0 = v22 + int32(144)
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
	v55 = F___memcpy(m, int32(4519580), v44, int32(140))
	mBase = m.M
	goto L15
L14:
	;
	goto L15
L15:
	;
	goto L9
L16:
	;
	v102 = int32(915)
	v104 = m.G0
	v106 = v104 - int32(144)
	m.G0 = v106
	switch int32(917) {
	case 0, 2:
		v116 = v102
		goto L30
	default:
		goto L31
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v74
	F_sigemptyset(m, v64+int32(8))
	mBase = m.M
	goto L20
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[403])) = v60
	v74 = int32(4729)
	goto L17
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+136)) = int32(268435456)
	v86 = v64 + int32(4)
	goto L24
L22:
	;
	m.G0 = v64 + int32(144)
	goto L16
L24:
	;
	goto L25
L25:
	;
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v97 = F___memcpy(m, int32(4519720), v86, int32(140))
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L22
L29:
	;
	v144 = int32(-2)
	v146 = m.G0
	v148 = v146 - int32(144)
	m.G0 = v148
	switch int32(0) {
	case 0, 2:
		v158 = v144
		goto L43
	default:
		goto L44
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v116
	F_sigemptyset(m, v106+int32(8))
	mBase = m.M
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[85])) = v102
	v116 = int32(4729)
	goto L30
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+136)) = int32(268435456)
	v128 = v106 + int32(4)
	goto L37
L35:
	;
	m.G0 = v106 + int32(144)
	goto L29
L37:
	;
	goto L38
L38:
	;
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v139 = F___memcpy(m, int32(4521540), v128, int32(140))
	mBase = m.M
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L35
L42:
	;
	v186 = int32(-2)
	v188 = m.G0
	v190 = v188 - int32(144)
	m.G0 = v190
	switch int32(0) {
	case 0, 2:
		v200 = v186
		goto L56
	default:
		goto L57
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v158
	F_sigemptyset(m, v148+int32(8))
	mBase = m.M
	goto L46
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[442])) = v144
	v158 = int32(4729)
	goto L43
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+136)) = int32(268435456)
	v170 = v148 + int32(4)
	goto L50
L48:
	;
	m.G0 = v148 + int32(144)
	goto L42
L50:
	;
	goto L51
L51:
	;
	if v170 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v181 = F___memcpy(m, int32(4521400), v170, int32(140))
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L48
L55:
	;
	v228 = int32(916)
	v230 = m.G0
	v232 = v230 - int32(144)
	m.G0 = v232
	switch int32(918) {
	case 0, 2:
		v242 = v228
		goto L69
	default:
		goto L70
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v200
	F_sigemptyset(m, v190+int32(8))
	mBase = m.M
	goto L59
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[414])) = v186
	v200 = int32(4729)
	goto L56
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+136)) = int32(268435456)
	v212 = v190 + int32(4)
	goto L63
L61:
	;
	m.G0 = v190 + int32(144)
	goto L55
L63:
	;
	goto L64
L64:
	;
	if v212 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v223 = F___memcpy(m, int32(4521260), v212, int32(140))
	mBase = m.M
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L61
L68:
	;
	v270 = int32(944)
	v272 = m.G0
	v274 = v272 - int32(144)
	m.G0 = v274
	switch int32(946) {
	case 0, 2:
		v284 = v270
		goto L82
	default:
		goto L83
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v242
	F_sigemptyset(m, v232+int32(8))
	mBase = m.M
	goto L72
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[415])) = v228
	v242 = int32(4729)
	goto L69
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+136)) = int32(268435456)
	v254 = v232 + int32(4)
	goto L76
L74:
	;
	m.G0 = v232 + int32(144)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if v254 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v265 = F___memcpy(m, int32(4520840), v254, int32(140))
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L74
L81:
	;
	v312 = int32(0)
	v314 = m.G0
	v316 = v314 - int32(144)
	m.G0 = v316
	switch int32(2) {
	case 0, 2:
		v326 = v312
		goto L95
	default:
		goto L96
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+4)) = v284
	F_sigemptyset(m, v274+int32(8))
	mBase = m.M
	goto L85
L83:
	;
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v270
	v284 = int32(4729)
	goto L82
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+136)) = int32(268435456)
	v296 = v274 + int32(4)
	goto L89
L87:
	;
	m.G0 = v274 + int32(144)
	goto L81
L89:
	;
	goto L90
L90:
	;
	if v296 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v307 = F___memcpy(m, int32(4521120), v296, int32(140))
	mBase = m.M
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L87
L94:
	;
	F_sigprocmask(m, int32(4332504), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L107
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+4)) = v326
	F_sigemptyset(m, v316+int32(8))
	mBase = m.M
	goto L97
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[418])) = v312
	v326 = int32(4729)
	goto L95
L97:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+136)) = int32(268435457)
	v338 = v316 + int32(4)
	goto L102
L100:
	;
	m.G0 = v316 + int32(144)
	goto L94
L102:
	;
	goto L103
L103:
	;
	if v338 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v349 = F___memcpy(m, int32(4521820), v338, int32(140))
	mBase = m.M
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L100
L107:
	;
	F_on_shmem_exit(m, int32(945), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[443]))
	v364 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v364
	v368 = F_palloc(m, int32(2888))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _consts[444])) = v368
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v371
	v376 = F_binaryheap_allocate(m, int32(64), int32(946), v371)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = v376
	v383 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	v388 = F_AllocSetContextCreateInternal(m, v383, int32(202320), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[445])) = v388
	v391 = m.G0
	v393 = v391 - int32(16)
	m.G0 = v393
	v396 = *(*int32)(unsafe.Add(mBase, _consts[446]))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396))))
	if v397 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v483 = m.G0
	v485 = v483 - int32(3392)
	m.G0 = v485
	goto L142
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L139
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L136
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L131
	}
L116:
	;
	v413 = m.T0[v411].(func(*base.Module) int32)(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L123
	}
L117:
	;
	v411 = int32(947)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	if v403 != 0 {
		goto L115
	} else {
		goto L120
	}
L120:
	;
	v405 = int32(0)
	v407 = F_load_external_function(m, v396, int32(92998), v405, v405)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v407 == int32(0) {
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v411 = v407
	goto L116
L123:
	;
	*(*int32)(unsafe.Add(mBase, _consts[447])) = v413
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v413)+8))
	if v416 == int32(0) {
		goto L113
	} else {
		goto L124
	}
L124:
	;
	v421 = F_palloc0(m, int32(4))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, _consts[448])) = v421
	v425 = *(*int32)(unsafe.Add(mBase, _consts[447]))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v426 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	m.T0[v426].(func(*base.Module, int32))(m, v421)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_before_shmem_exit(m, int32(948), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	m.G0 = v393 + int32(16)
	goto L112
L131:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(98672), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errdetail(m, int32(533552), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(467367), int32(921), int32(16301))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = int32(92998)
	F_errmsg(m, int32(173716), v393)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(467367), int32(936), int32(16301))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
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
	F_errmsg(m, int32(299781), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(467367), int32(942), int32(16301))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
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
	v499 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = int32(0)
	goto L145
L143:
	;
	m.G0 = v485 + int32(3392)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L459
	}
L144:
	;
	goto L143
L145:
	;
	v503 = *(*int32)(unsafe.Add(mBase, _consts[449]))
	F_ProcessPgArchInterrupts(m)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[441]))
	if v507 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, uint32(v524)+4)) = int32(0)
	goto L153
L148:
	;
	v510 = F___time(m)
	mBase = m.M
	v512 = *(*int64)(unsafe.Add(mBase, _consts[450]))
	if v512 == int64(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, _consts[450])) = v510
	goto L147
L150:
	;
	goto L151
L151:
	;
	if base.Ui32(int32(59)) < base.Ui32(base.I32_wrap_i64(v510-v512)) {
		goto L144
	} else {
		goto L152
	}
L152:
	;
	goto L147
L153:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _consts[443]))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v539)+4)) = int32(0)
	v544 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	if v540 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	if v503 != 0 {
		goto L144
	} else {
		goto L456
	}
L155:
	;
	goto L154
L156:
	;
	v1400 = int32(0)
	v1403 = *(*int32)(unsafe.Add(mBase, _consts[441]))
	if v1403 != 0 {
		goto L155
	} else {
		goto L380
	}
L157:
	;
	v1307 = v1298 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1296)+4)) = v1307
	v1310 = v485 + int32(1296)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1296+v1307<<(uint(int32(2))%32))+8))
	if (v1314^v1310)&int32(3) != 0 {
		goto L362
	} else {
		goto L363
	}
L158:
	;
	v1272 = v1268
	v1273 = int32(0)
	goto L355
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L351
	}
L160:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	v687 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v686)+8)) = uint8(v687)
	*(*int32)(unsafe.Add(mBase, uint32(v686))) = int32(0)
	goto L195
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v544)+4)) = int32(0)
	v676 = v544
	goto L160
L162:
	;
	goto L163
L163:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v549 <= int32(0) {
		v676 = v544
		goto L160
	} else {
		goto L164
	}
L164:
	;
	v553 = v544
	v554 = v549
	goto L165
L165:
	;
	v564 = v554 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v553)+4)) = v564
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v553+v564<<(uint(int32(2))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+164)) = int32(20954)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+160)) = v569
	v579 = F_pg_snprintf(m, v485+int32(1344), int32(1024), int32(163820), v485+int32(160))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	v676 = v671
	goto L160
L167:
	;
	v587 = F___fstatat(m, int32(-100), v485+int32(1344), v485+int32(176), int32(0))
	mBase = m.M
	goto L168
L168:
	;
	if v587 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v591 = v485 + int32(1296)
	if (v569^v591)&int32(3) != 0 {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	goto L171
L171:
	;
	v667 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v667 != int32(44) {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v646))) = uint8(v645)
	if v645&int32(255) == int32(0) {
		goto L173
	} else {
		goto L189
	}
L175:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569))))
	v644 = v569
	v645 = v597
	v646 = v591
	goto L174
L176:
	;
	goto L177
L177:
	;
	if v569&int32(3) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v601 = v569
	v603 = v591
	goto L181
L179:
	;
	v615 = v569
	v617 = v591
	goto L180
L180:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v622 = int32(-2139062144)
	if (int32(16843008)-v619|v619)&v622 != v622 {
		v644 = v615
		v645 = v619
		v646 = v617
		goto L174
	} else {
		goto L185
	}
L181:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v604)
	if v604 == int32(0) {
		goto L173
	} else {
		goto L183
	}
L182:
	;
	v615 = v611
	v617 = v609
	goto L180
L183:
	;
	v608 = int32(1)
	v609 = v603 + v608
	v611 = v601 + v608
	if v611&int32(3) != 0 {
		v601 = v611
		v603 = v609
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v627 = v615
	v628 = v619
	v629 = v617
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v629))) = v628
	v631 = int32(4)
	v632 = v629 + v631
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v627)+4))
	v635 = v627 + v631
	v639 = int32(-2139062144)
	if (v633|(int32(16843008)-v633))&v639 == v639 {
		v627 = v635
		v628 = v633
		v629 = v632
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v644 = v635
	v645 = v633
	v646 = v632
	goto L174
L188:
	;
	goto L187
L189:
	;
	v653 = v644
	v655 = v646
	goto L190
L190:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v655)+1)) = uint8(v656)
	v658 = int32(1)
	if v656 != 0 {
		v653 = v653 + v658
		v655 = v655 + v658
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
	v671 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	if int32(0) < v672 {
		v553 = v671
		v554 = v672
		goto L165
	} else {
		goto L194
	}
L194:
	;
	goto L166
L195:
	;
	v696 = F_pg_snprintf(m, v485+int32(2368), int32(1024), int32(105496), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v700 = F_AllocateDir(m, v485+int32(2368))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v704 = F_ReadDir(m, v700, v485+int32(2368))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	if v704 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v707 = v704
	goto L202
L200:
	;
	goto L201
L201:
	;
	F_FreeDir(m, v700)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L344
	}
L202:
	;
	v718 = v707 + int32(19)
	if v718&int32(3) == int32(0) {
		v742 = v718
		goto L207
	} else {
		goto L208
	}
L203:
	;
	goto L201
L204:
	;
	v1216 = F_ReadDir(m, v700, v485+int32(2368))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L342
	}
L205:
	;
	if base.Ui32(v775-int32(47)) < base.Ui32(int32(-25)) {
		goto L204
	} else {
		goto L222
	}
L206:
	;
	v775 = v767 - v718
	goto L205
L207:
	;
	v746 = v742
	goto L216
L208:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	if v726 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v775 = int32(0)
	goto L205
L210:
	;
	goto L211
L211:
	;
	v731 = v718
	goto L212
L212:
	;
	v735 = v731 + int32(1)
	if v735&int32(3) == int32(0) {
		v742 = v735
		goto L207
	} else {
		goto L214
	}
L213:
	;
	v767 = v735
	goto L206
L214:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	if v740 != 0 {
		v731 = v735
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	v755 = int32(-2139062144)
	if (int32(16843008)-v752|v752)&v755 == v755 {
		v746 = v746 + int32(4)
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v761 = v746
	goto L219
L218:
	;
	goto L217
L219:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761))))
	if v765 != 0 {
		v761 = v761 + int32(1)
		goto L219
	} else {
		goto L221
	}
L220:
	;
	v767 = v761
	goto L206
L221:
	;
	goto L220
L222:
	;
	v780 = int32(294635)
	v784 = m.G0
	v786 = v784 - int32(32)
	v787 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v786)+24)) = v787
	*(*int64)(unsafe.Add(mBase, uint32(v786)+16)) = v787
	*(*int64)(unsafe.Add(mBase, uint32(v786)+8)) = v787
	*(*int64)(unsafe.Add(mBase, uint32(v786))) = v787
	v795 = int32(*(*uint8)(unsafe.Add(mBase, _consts[451])))
	if v795 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v865 = v775 - int32(6)
	if base.Ui32(v863) < base.Ui32(v865) {
		goto L204
	} else {
		goto L244
	}
L224:
	;
	v863 = int32(0)
	goto L223
L225:
	;
	goto L226
L226:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, _consts[452])))
	if v799 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v803 = v718
	goto L230
L228:
	;
	goto L229
L229:
	;
	v813 = v780
	v814 = v795
	goto L233
L230:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
	if v809 == v795 {
		v803 = v803 + int32(1)
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v863 = v803 - v718
	goto L223
L232:
	;
	goto L231
L233:
	;
	v821 = v786 + int32(base.Ui32(v814)>>(uint(int32(3))%32))&int32(28)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	v823 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = v822 | v823<<(uint(v814)%32)
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813)+1)))
	if v827 != 0 {
		v813 = v813 + v823
		v814 = v827
		goto L233
	} else {
		goto L235
	}
L234:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	if v830 == int32(0) {
		v855 = v718
		goto L236
	} else {
		goto L237
	}
L235:
	;
	goto L234
L236:
	;
	v863 = v855 - v718
	goto L223
L237:
	;
	v834 = v718
	v835 = v830
	goto L238
L238:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v786+int32(base.Ui32(v835)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v843)>>(uint(v835)%32))&int32(1) == int32(0) {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v855 = v851
	goto L236
L240:
	;
	v855 = v834
	goto L236
L241:
	;
	goto L242
L242:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+1)))
	v851 = v834 + int32(1)
	if v849 != 0 {
		v834 = v851
		v835 = v849
		goto L238
	} else {
		goto L243
	}
L243:
	;
	goto L239
L244:
	;
	v867 = v718 + v865
	v868 = int32(20954)
	v871 = int32(*(*uint8)(unsafe.Add(mBase, _consts[453])))
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	if v872 == int32(0) {
		v891 = v871
		v892 = v872
		goto L246
	} else {
		goto L247
	}
L245:
	;
	if v892-v891 != 0 {
		goto L204
	} else {
		goto L253
	}
L246:
	;
	goto L245
L247:
	;
	if v871 != v872 {
		v891 = v871
		v892 = v872
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v876 = v867
	v877 = v868
	goto L249
L249:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+1)))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+1)))
	if v881 == int32(0) {
		v891 = v880
		v892 = v881
		goto L246
	} else {
		goto L251
	}
L250:
	;
	v891 = v880
	v892 = v881
	goto L246
L251:
	;
	v884 = int32(1)
	if v880 == v881 {
		v876 = v876 + v884
		v877 = v877 + v884
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	if v865 != 0 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v901 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v485+int32(1344)+v865))) = uint8(v901)
	v904 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v904)))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)))
	if v906 <= int32(63) {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v896 = F__emscripten_memcpy_bulkmem(m, v485+int32(1344), v718, v865)
	mBase = m.M
	goto L257
L256:
	;
	goto L257
L257:
	;
	goto L254
L258:
	;
	v913 = v904 + v906*int32(41) + int32(264)
	v915 = v485 + int32(1344)
	if (v915^v913)&int32(3) != 0 {
		goto L264
	} else {
		goto L265
	}
L259:
	;
	goto L260
L260:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v905)+20))
	v1003 = v485 + int32(1344)
	v1004 = int32(0)
	v1005 = F_strlen(m, v1001)
	mBase = m.M
	if v1005 != int32(16) {
		v1018 = v1004
		goto L286
	} else {
		goto L287
	}
L261:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v904)))
	F_binaryheap_add_unordered(m, v990, v913)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L282
	}
L262:
	;
	goto L261
L263:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v970))) = uint8(v969)
	if v969&int32(255) == int32(0) {
		goto L262
	} else {
		goto L278
	}
L264:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915))))
	v968 = v915
	v969 = v921
	v970 = v913
	goto L263
L265:
	;
	goto L266
L266:
	;
	if v915&int32(3) != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v925 = v915
	v927 = v913
	goto L270
L268:
	;
	v939 = v915
	v941 = v913
	goto L269
L269:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v939)))
	v946 = int32(-2139062144)
	if (int32(16843008)-v943|v943)&v946 != v946 {
		v968 = v939
		v969 = v943
		v970 = v941
		goto L263
	} else {
		goto L274
	}
L270:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v925))))
	*(*uint8)(unsafe.Add(mBase, uint32(v927))) = uint8(v928)
	if v928 == int32(0) {
		goto L262
	} else {
		goto L272
	}
L271:
	;
	v939 = v935
	v941 = v933
	goto L269
L272:
	;
	v932 = int32(1)
	v933 = v927 + v932
	v935 = v925 + v932
	if v935&int32(3) != 0 {
		v925 = v935
		v927 = v933
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	v951 = v939
	v952 = v943
	v953 = v941
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v952
	v955 = int32(4)
	v956 = v953 + v955
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v951)+4))
	v959 = v951 + v955
	v963 = int32(-2139062144)
	if (v957|(int32(16843008)-v957))&v963 == v963 {
		v951 = v959
		v952 = v957
		v953 = v956
		goto L275
	} else {
		goto L277
	}
L276:
	;
	v968 = v959
	v969 = v957
	v970 = v956
	goto L263
L277:
	;
	goto L276
L278:
	;
	v977 = v968
	v979 = v970
	goto L279
L279:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v979)+1)) = uint8(v980)
	v982 = int32(1)
	if v980 != 0 {
		v977 = v977 + v982
		v979 = v979 + v982
		goto L279
	} else {
		goto L281
	}
L280:
	;
	goto L262
L281:
	;
	goto L280
L282:
	;
	v994 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v994)))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	if v996 != int32(64) {
		goto L204
	} else {
		goto L283
	}
L283:
	;
	F_binaryheap_build(m, v995)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	goto L204
L285:
	;
	if v1037 <= int32(0) {
		goto L204
	} else {
		goto L301
	}
L286:
	;
	v1019 = F_strlen(m, v1003)
	mBase = m.M
	if v1019 == int32(16) {
		goto L292
	} else {
		goto L293
	}
L287:
	;
	v1009 = F_strspn(m, v1001, int32(505149))
	mBase = m.M
	if v1009 != int32(8) {
		v1018 = v1004
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v1015 = F_strcmp(m, v1001+int32(8), int32(11754))
	mBase = m.M
	v1018 = base.B2i32(v1015 == int32(0))
	goto L286
L289:
	;
	v1036 = F_strcmp(m, v1001, v1003)
	mBase = m.M
	v1037 = v1036
	goto L285
L290:
	;
	if v1018 != 0 {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	v1029 = F_strcmp(m, v485+int32(1352), int32(11754))
	mBase = m.M
	if v1018 == base.B2i32(v1029 == int32(0)) {
		goto L289
	} else {
		goto L297
	}
L292:
	;
	v1023 = F_strspn(m, v1003, int32(505149))
	mBase = m.M
	if v1023 == int32(8) {
		goto L291
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	if v1018 != 0 {
		goto L290
	} else {
		goto L296
	}
L295:
	;
	goto L294
L296:
	;
	goto L289
L297:
	;
	goto L290
L298:
	;
	v1035 = int32(-1)
	goto L300
L299:
	;
	v1035 = int32(1)
	goto L300
L300:
	;
	v1037 = v1035
	goto L285
L301:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)))
	v1043 = F_binaryheap_remove_first(m, v1042)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1046 = v485 + int32(1344)
	if (v1046^v1043)&int32(3) != 0 {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1122)))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+4))
	if v1124 < v1125 {
		goto L325
	} else {
		goto L326
	}
L304:
	;
	goto L303
L305:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1101))) = uint8(v1100)
	if v1100&int32(255) == int32(0) {
		goto L304
	} else {
		goto L320
	}
L306:
	;
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046))))
	v1099 = v1046
	v1100 = v1052
	v1101 = v1043
	goto L305
L307:
	;
	goto L308
L308:
	;
	if v1046&int32(3) != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1056 = v1046
	v1058 = v1043
	goto L312
L310:
	;
	v1070 = v1046
	v1072 = v1043
	goto L311
L311:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	v1077 = int32(-2139062144)
	if (int32(16843008)-v1074|v1074)&v1077 != v1077 {
		v1099 = v1070
		v1100 = v1074
		v1101 = v1072
		goto L305
	} else {
		goto L316
	}
L312:
	;
	v1059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1058))) = uint8(v1059)
	if v1059 == int32(0) {
		goto L304
	} else {
		goto L314
	}
L313:
	;
	v1070 = v1066
	v1072 = v1064
	goto L311
L314:
	;
	v1063 = int32(1)
	v1064 = v1058 + v1063
	v1066 = v1056 + v1063
	if v1066&int32(3) != 0 {
		v1056 = v1066
		v1058 = v1064
		goto L312
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	v1082 = v1070
	v1083 = v1074
	v1084 = v1072
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1084))) = v1083
	v1086 = int32(4)
	v1087 = v1084 + v1086
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+4))
	v1090 = v1082 + v1086
	v1094 = int32(-2139062144)
	if (v1088|(int32(16843008)-v1088))&v1094 == v1094 {
		v1082 = v1090
		v1083 = v1088
		v1084 = v1087
		goto L317
	} else {
		goto L319
	}
L318:
	;
	v1099 = v1090
	v1100 = v1088
	v1101 = v1087
	goto L305
L319:
	;
	goto L318
L320:
	;
	v1108 = v1099
	v1110 = v1101
	goto L321
L321:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1110)+1)) = uint8(v1111)
	v1113 = int32(1)
	if v1111 != 0 {
		v1108 = v1108 + v1113
		v1110 = v1110 + v1113
		goto L321
	} else {
		goto L323
	}
L322:
	;
	goto L304
L323:
	;
	goto L322
L324:
	;
	goto L204
L325:
	;
	v1128 = v1123 + int32(20)
	v1129 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1128+v1124<<(uint(v1129)%32)))) = v1043
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	*(*int32)(unsafe.Add(mBase, uint32(v1123))) = v1133 + int32(1)
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1128+v1133<<(uint(v1129)%32))))
	if v1133 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L326:
	;
	goto L327
L327:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L339
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1128+v1177<<(uint(int32(2))%32)))) = v1140
	goto L324
L329:
	;
	v1177 = int32(0)
	goto L328
L330:
	;
	goto L331
L331:
	;
	v1145 = v1133
	goto L332
L332:
	;
	v1157 = int32(2)
	v1158 = base.I32_div_s(v1145-int32(1), v1157)
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1128+v1158<<(uint(v1157)%32))))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+16))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+12))
	v1165 = m.T0[v1164].(func(*base.Module, int32, int32, int32) int32)(m, v1140, v1162, v1163)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L334
	}
L333:
	;
	v1177 = v1158
	goto L328
L334:
	;
	if v1165 <= int32(0) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1177 = v1145
	goto L328
L336:
	;
	goto L337
L337:
	;
	v1169 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1128+v1145<<(uint(v1169)%32)))) = v1162
	if base.Ui32(v1169) < base.Ui32(v1145) {
		v1145 = v1158
		goto L332
	} else {
		goto L338
	}
L338:
	;
	goto L333
L339:
	;
	F_errmsg_internal(m, int32(109558), int32(0))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	F_errfinish(m, int32(465573), int32(161), int32(435706))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L342:
	;
	if v1216 != 0 {
		v707 = v1216
		goto L202
	} else {
		goto L343
	}
L343:
	;
	goto L203
L344:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1233)))
	if v1234 == int32(0) {
		goto L155
	} else {
		goto L345
	}
L345:
	;
	if int32(64) <= v1234 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+4)) = v1234
	v1268 = v1232
	goto L158
L347:
	;
	goto L348
L348:
	;
	F_binaryheap_build(m, v1233)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1243)))
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1244)))
	*(*int32)(unsafe.Add(mBase, uint32(v1243)+4)) = v1245
	if int32(0) < v1245 {
		v1268 = v1243
		goto L158
	} else {
		goto L350
	}
L350:
	;
	v1296 = v1243
	v1298 = v1245
	goto L157
L351:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+144)) = v485 + int32(1344)
	F_errmsg(m, int32(279590), v485+int32(144))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(467367), int32(682), int32(307020))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	v1283 = F_binaryheap_remove_first(m, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L357
	}
L356:
	;
	v1296 = v1286
	v1298 = v1293
	goto L157
L357:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, uint32(v1286+v1273<<(uint(int32(2))%32))+8)) = v1283
	v1292 = v1273 + int32(1)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	if v1292 < v1293 {
		v1272 = v1286
		v1273 = v1292
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	goto L156
L360:
	;
	goto L359
L361:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1369))) = uint8(v1368)
	if v1368&int32(255) == int32(0) {
		goto L360
	} else {
		goto L376
	}
L362:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314))))
	v1367 = v1314
	v1368 = v1320
	v1369 = v1310
	goto L361
L363:
	;
	goto L364
L364:
	;
	if v1314&int32(3) != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1324 = v1314
	v1326 = v1310
	goto L368
L366:
	;
	v1338 = v1314
	v1340 = v1310
	goto L367
L367:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1338)))
	v1345 = int32(-2139062144)
	if (int32(16843008)-v1342|v1342)&v1345 != v1345 {
		v1367 = v1338
		v1368 = v1342
		v1369 = v1340
		goto L361
	} else {
		goto L372
	}
L368:
	;
	v1327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1324))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1326))) = uint8(v1327)
	if v1327 == int32(0) {
		goto L360
	} else {
		goto L370
	}
L369:
	;
	v1338 = v1334
	v1340 = v1332
	goto L367
L370:
	;
	v1331 = int32(1)
	v1332 = v1326 + v1331
	v1334 = v1324 + v1331
	if v1334&int32(3) != 0 {
		v1324 = v1334
		v1326 = v1332
		goto L368
	} else {
		goto L371
	}
L371:
	;
	goto L369
L372:
	;
	v1350 = v1338
	v1351 = v1342
	v1352 = v1340
	goto L373
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1352))) = v1351
	v1354 = int32(4)
	v1355 = v1352 + v1354
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+4))
	v1358 = v1350 + v1354
	v1362 = int32(-2139062144)
	if (v1356|(int32(16843008)-v1356))&v1362 == v1362 {
		v1350 = v1358
		v1351 = v1356
		v1352 = v1355
		goto L373
	} else {
		goto L375
	}
L374:
	;
	v1367 = v1358
	v1368 = v1356
	v1369 = v1355
	goto L361
L375:
	;
	goto L374
L376:
	;
	v1376 = v1367
	v1378 = v1369
	goto L377
L377:
	;
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1378)+1)) = uint8(v1379)
	v1381 = int32(1)
	if v1379 != 0 {
		v1376 = v1376 + v1381
		v1378 = v1378 + v1381
		goto L377
	} else {
		goto L379
	}
L378:
	;
	goto L360
L379:
	;
	goto L378
L380:
	;
	v1406 = v1400
	v1407 = v1400
	goto L381
L381:
	;
	v1415 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L383
	}
L382:
	;
	goto L155
L383:
	;
	if v1415 == int32(0) {
		goto L155
	} else {
		goto L384
	}
L384:
	;
	F_ProcessPgArchInterrupts(m)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	v1422 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[190])) = v1422
	v1426 = *(*int32)(unsafe.Add(mBase, _consts[447]))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+4))
	if v1427 == v1422 {
		goto L388
	} else {
		goto L389
	}
L386:
	;
	F_pg_usleep(m, int32(1000000))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L454
	}
L387:
	;
	F_errfinish(m, int32(467367), v1739, int32(220650))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L453
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+112)) = v485 + int32(1296)
	v1469 = F_pg_snprintf(m, v485+int32(176), int32(1024), int32(165223), v485+int32(112))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L398
	}
L389:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v1432 = m.T0[v1427].(func(*base.Module, int32) int32)(m, v1431)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	if v1432 != 0 {
		goto L388
	} else {
		goto L391
	}
L391:
	;
	v1436 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	if v1436 == int32(0) {
		goto L155
	} else {
		goto L393
	}
L393:
	;
	F_errmsg(m, int32(423233), int32(0))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v1446 == int32(0) {
		v1739 = int32(431)
		goto L387
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+128)) = v1446
	F_errdetail_internal(m, int32(193943), v485+int32(128))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(467367), int32(431), int32(220650))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	goto L155
L398:
	;
	v1477 = F___fstatat(m, int32(-100), v485+int32(176), v485+int32(1200), int32(0))
	mBase = m.M
	goto L400
L399:
	;
	v1546 = F_pgarch_archiveXlog(m, v485+int32(1296))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L416
	}
L400:
	;
	if v1477 == int32(0) {
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v1481 != int32(44) {
		goto L399
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+100)) = int32(20954)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+96)) = v485 + int32(1296)
	v1495 = F_pg_snprintf(m, v485+int32(2368), int32(1024), int32(163820), v485+int32(96))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	v1499 = F_unlink(m, v485+int32(2368))
	mBase = m.M
	if v1499 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1504 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	v1522 = v1407 + int32(1)
	if v1522 < int32(3) {
		v1743 = v1406
		v1744 = v1522
		goto L386
	} else {
		goto L411
	}
L407:
	;
	if v1504 == int32(0) {
		goto L153
	} else {
		goto L408
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+64)) = v485 + int32(2368)
	F_errmsg(m, int32(665221), v485-int32(-64))
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(467367), int32(454), int32(220650))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	goto L153
L411:
	;
	v1527 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	if v1527 == int32(0) {
		goto L155
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+80)) = v485 + int32(2368)
	F_errmsg(m, int32(204527), v485+int32(80))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	F_errfinish(m, int32(467367), int32(464), int32(220650))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	goto L155
L416:
	;
	if v1546 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+36)) = int32(20954)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+32)) = v485 + int32(1296)
	v1559 = F_pg_snprintf(m, v485+int32(2368), int32(1024), int32(163820), v485+int32(32))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	v1662 = v485 + int32(1296)
	v1663 = int32(1)
	v1668 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	v1669 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v1670 = int32(4419940)
	v1672 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1672 + v1663
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1668)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1668)+40)) = v1676 + v1663
	goto L440
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+20)) = int32(350192)
	*(*int32)(unsafe.Add(mBase, uint32(v485)+16)) = v485 + int32(1296)
	v1572 = F_pg_snprintf(m, v485+int32(1344), int32(1024), int32(163820), v485+int32(16))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	v1578 = F_rename(m, v485+int32(2368), v485+int32(1344))
	mBase = m.M
	if int32(0) <= v1578 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1604 = v485 + int32(1296)
	v1610 = *(*int32)(unsafe.Add(mBase, _consts[454]))
	v1611 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v1612 = int32(4419940)
	v1614 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1615 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1614 + v1615
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+40)) = v1618 + v1615
	goto L431
L423:
	;
	v1583 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	if v1583 == int32(0) {
		goto L422
	} else {
		goto L425
	}
L425:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+4)) = v485 + int32(1344)
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v485 + int32(2368)
	F_errmsg(m, int32(279237), v485)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(467367), int32(837), int32(350250))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	goto L422
L429:
	;
	goto L153
L431:
	;
	goto L432
L432:
	;
	v1625 = v1610 + int32(48)
	v1626 = *(*int64)(unsafe.Add(mBase, uint32(v1625)))
	*(*int64)(unsafe.Add(mBase, uint32(v1625))) = v1626 + int64(1)
	goto L434
L434:
	;
	goto L435
L435:
	;
	v1633 = v1610 + int32(56)
	v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1633)+40)) = uint8(v1634)
	v1636 = *(*int64)(unsafe.Add(mBase, uint32(v1604)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+32)) = v1636
	v1638 = *(*int64)(unsafe.Add(mBase, uint32(v1604)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+24)) = v1638
	v1640 = *(*int64)(unsafe.Add(mBase, uint32(v1604)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+16)) = v1640
	v1642 = *(*int64)(unsafe.Add(mBase, uint32(v1604)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1633)+8)) = v1642
	v1644 = *(*int64)(unsafe.Add(mBase, uint32(v1604)))
	*(*int64)(unsafe.Add(mBase, uint32(v1633))) = v1644
	goto L437
L437:
	;
	goto L438
L438:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1610+int32(104)))) = v1611
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+40))
	v1652 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+40)) = v1651 + v1652
	v1655 = int32(4419940)
	v1657 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1657 - v1652
	goto L429
L439:
	;
	v1720 = v1406 + int32(1)
	if v1720 < int32(3) {
		v1743 = v1720
		v1744 = v1407
		goto L386
	} else {
		goto L449
	}
L440:
	;
	goto L442
L442:
	;
	v1683 = v1668 + int32(112)
	v1684 = *(*int64)(unsafe.Add(mBase, uint32(v1683)))
	*(*int64)(unsafe.Add(mBase, uint32(v1683))) = v1684 + int64(1)
	goto L443
L443:
	;
	goto L445
L445:
	;
	v1691 = v1668 + int32(120)
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1662)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1691)+40)) = uint8(v1692)
	v1694 = *(*int64)(unsafe.Add(mBase, uint32(v1662)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1691)+32)) = v1694
	v1696 = *(*int64)(unsafe.Add(mBase, uint32(v1662)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1691)+24)) = v1696
	v1698 = *(*int64)(unsafe.Add(mBase, uint32(v1662)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1691)+16)) = v1698
	v1700 = *(*int64)(unsafe.Add(mBase, uint32(v1662)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1691)+8)) = v1700
	v1702 = *(*int64)(unsafe.Add(mBase, uint32(v1662)))
	*(*int64)(unsafe.Add(mBase, uint32(v1691))) = v1702
	goto L446
L446:
	;
	goto L448
L448:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1668+int32(168)))) = v1669
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1668)+40))
	v1710 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1668)+40)) = v1709 + v1710
	v1713 = int32(4419940)
	v1715 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1715 - v1710
	goto L439
L449:
	;
	v1725 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	if v1725 == int32(0) {
		goto L155
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v485)+48)) = v485 + int32(1296)
	F_errmsg(m, int32(204614), v485+int32(48))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	v1739 = int32(500)
	goto L387
L453:
	;
	goto L155
L454:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, _consts[441]))
	if v1749 == int32(0) {
		v1406 = v1743
		v1407 = v1744
		goto L381
	} else {
		goto L455
	}
L455:
	;
	goto L382
L456:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v1768 = F_WaitLatch(m, v1764, int32(25), int32(60000), int32(83886080))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	if v1768&int32(16) == int32(0) {
		goto L142
	} else {
		goto L458
	}
L458:
	;
	goto L144
L459:
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	v7 = m.G0
	v9 = v7 - int32(2080)
	m.G0 = v9
	v11 = F_AllocateDir(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(2080)
	return
L2:
	;
	v20 = F_ReadDirExtended(m, v11, l0, int32(15))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L8
	}
L3:
	;
	return
L4:
	;
	if v11 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v16 == int32(44) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L2
L8:
	;
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = v20
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_FreeDir(m, v11)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L59
	}
L12:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+19)))
	if v28 != int32(46) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L11
L14:
	;
	v164 = F_ReadDirExtended(m, v11, l0, int32(15))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L57
	}
L15:
	;
	v41 = v23 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	v50 = F_pg_snprintf(m, v9+int32(32), int32(2048), int32(165202), v9+int32(16))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L20
	}
L16:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
	if v31 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
	if v34 != int32(46) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+21)))
	if v37 == int32(0) {
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
	F_errmsg(m, v150, v9)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L55
	}
L22:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L54
	}
L23:
	;
	v140 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L52
	}
L24:
	;
	v54 = int32(220933)
	goto L29
L25:
	;
	goto L26
L26:
	;
	v104 = F_get_dirent_type(m, v9+int32(32), v23, int32(0), int32(15))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L44
	}
L27:
	;
	if v91-v92 != 0 {
		goto L23
	} else {
		goto L41
	}
L29:
	;
	goto L30
L30:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v61 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v62 = v41
	v63 = v54
	v64 = int32(9)
	v65 = v61
	goto L35
L32:
	;
	v87 = v54
	v91 = int32(0)
	goto L33
L33:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	goto L27
L34:
	;
	v87 = v82
	v91 = v84
	goto L33
L35:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v65 != v67 {
		v82 = v63
		v84 = v65
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v82 = v76
	v84 = int32(0)
	goto L34
L37:
	;
	if v67 == int32(0) {
		v82 = v63
		v84 = v65
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v72 = v64 - int32(1)
	if v72 == int32(0) {
		v82 = v63
		v84 = v65
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v75 = int32(1)
	v76 = v63 + v75
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v77 != 0 {
		v62 = v62 + v75
		v63 = v76
		v64 = v72
		v65 = v77
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	goto L26
L42:
	;
	v127 = F_unlink(m, v9+int32(32))
	mBase = m.M
	if int32(0) <= v127 {
		goto L14
	} else {
		goto L49
	}
L43:
	;
	F_RemovePgTempFilesInDir(m, v9+int32(32), int32(0), int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L45
	}
L44:
	;
	switch v104 {
	case 0:
		goto L14
	default:
		goto L42
	case 3:
		goto L43
	}
L45:
	;
	v114 = F_rmdir(m, v9+int32(32))
	mBase = m.M
	if int32(0) <= v114 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	v119 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	if v119 == int32(0) {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	v146 = int32(278483)
	v147 = int32(3441)
	goto L22
L49:
	;
	v132 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	if v132 == int32(0) {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	v146 = int32(281042)
	v147 = int32(3449)
	goto L22
L52:
	;
	if v140 == int32(0) {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	v150 = int32(674150)
	v151 = int32(3455)
	goto L21
L54:
	;
	v150 = v146
	v151 = v147
	goto L21
L55:
	;
	F_errfinish(m, int32(469068), v151, int32(200893))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	goto L14
L57:
	;
	if v164 != 0 {
		v23 = v164
		goto L12
	} else {
		goto L58
	}
L58:
	;
	goto L13
L59:
	;
	goto L1
}
func F_pg_base64_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	return int64(base.Ui64(base.I64_extend_i32_u(l1)*int64(3)) >> (uint(int64(2)) % 64))
}
func F_pg_base64_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L26
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L26
	} else {
		goto L43
	}
L3:
	;
	v18 = l0
	v21 = l2
	v23 = v4
	v24 = v4
	v25 = v4
	goto L6
L4:
	;
	v161 = l2
	goto L5
L5:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_s(v161 - l2)
L6:
	;
	v30 = v18
	goto L9
L7:
	;
	if v157 != 0 {
		goto L1
	} else {
		goto L42
	}
L8:
	;
	goto L7
L9:
	;
	v41 = v30 + int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v44 = v42 - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v44) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v42 == int32(61) {
		goto L20
	} else {
		goto L21
	}
L11:
	;
	goto L10
L12:
	;
	if int32(1)<<(uint(v44)%32)&int32(8388627) == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if base.Ui32(v41) < base.Ui32(v16) {
		v30 = v41
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v153 = v21
	v157 = v25
	goto L8
L15:
	;
	if base.Ui32(v41) < base.Ui32(v16) {
		v18 = v41
		v21 = v146
		v23 = v148
		v24 = v149
		v25 = v150
		goto L6
	} else {
		goto L41
	}
L16:
	;
	v132 = int32(0)
	if base.B2i32(v128 == v132)&base.B2i32(base.Ui32(v129) < base.Ui32(int32(3))) == v132 {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	v123 = int32(base.Ui32(v118) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v123)
	v127 = v118
	v128 = v119
	v129 = v120
	v131 = v21 + int32(2)
	goto L16
L18:
	;
	v113 = int32(base.Ui32(v24) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v113)
	v118 = v24 << (uint(int32(6)) % 32)
	v119 = v56
	v120 = int32(2)
	goto L17
L19:
	;
	v95 = v92 + v24<<(uint(int32(6))%32)
	v97 = v25 + int32(1)
	if v97 != int32(4) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v56 = int32(0)
	if v23 != 0 {
		v92 = v56
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(int32(125)) < base.Ui32((v42-int32(1))&int32(255)) {
		goto L2
	} else {
		goto L31
	}
L23:
	;
	switch v25 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L18
	default:
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v146 = v21
	v148 = int32(1)
	v149 = v24 << (uint(int32(6)) % 32)
	v150 = int32(3)
	goto L15
L26:
	;
	return int64(0)
L27:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(391262), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(468817), int32(365), int32(388914))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[792]))))
	if v89 < int32(0) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v92 = v89
	goto L19
L33:
	;
	v146 = v21
	v148 = v23
	v149 = v95
	v150 = v97
	goto L15
L34:
	;
	goto L35
L35:
	;
	v101 = int32(base.Ui32(v95) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v101)
	v104 = base.B2i32(v23 == int32(0))
	if v23 == int32(0) {
		v118 = v95
		v119 = v104
		v120 = v23
		goto L17
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v23) {
		v118 = v95
		v119 = v104
		v120 = v23
		goto L17
	} else {
		goto L37
	}
L37:
	;
	v127 = v95
	v128 = int32(0)
	v129 = v23
	v131 = v21 + int32(1)
	goto L16
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v127)
	v143 = v131 + int32(1)
	goto L40
L39:
	;
	v143 = v131
	goto L40
L40:
	;
	v146 = v143
	v148 = v129
	v149 = v132
	v150 = int32(0)
	goto L15
L41:
	;
	v153 = v146
	v157 = v150
	goto L8
L42:
	;
	v161 = v153
	goto L5
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v183 = F_pg_mblen_range(m, v30, v16)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v183
	F_errmsg(m, int32(391203), v14)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(468817), int32(378), int32(388914))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L26
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(391126), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L26
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(598523), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(468817), int32(399), int32(388914))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L26
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_palloc(m, int32(36))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v24
	v31 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v31
	v37 = F_palloc(m, v31*int32(20))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v43 = F_palloc(m, v40*int32(56))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v49 = F_palloc(m, v46<<(uint(int32(2))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v57 = F_LWLockAcquire(m, v53+int32(512), int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v18 == int32(0) {
		v103 = v2
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v103 != 0 {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 <= int32(0) {
		v103 = v2
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[603]))
	v71 = int32(0)
	goto L10
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(36)+v71<<(uint(int32(2))%32))))
	v94 = v70 + v91*int32(640)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	if v95 == v18 {
		v103 = v94
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v103 = int32(0)
	goto L7
L12:
	;
	v98 = v71 + int32(1)
	if v98 != v63 {
		v71 = v98
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v123 = F_LWLockAcquire(m, v119+int32(23296), int32(1))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v395+int32(512))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L60
	}
L17:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v130 = F_LWLockAcquire(m, v126+int32(23424), int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v137 = F_LWLockAcquire(m, v133+int32(23552), int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v144 = F_LWLockAcquire(m, v140+int32(23680), int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v151 = F_LWLockAcquire(m, v147+int32(23808), int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v158 = F_LWLockAcquire(m, v154+int32(23936), int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v165 = F_LWLockAcquire(m, v161+int32(24064), int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v172 = F_LWLockAcquire(m, v168+int32(24192), int32(1))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v179 = F_LWLockAcquire(m, v175+int32(24320), int32(1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v186 = F_LWLockAcquire(m, v182+int32(24448), int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v193 = F_LWLockAcquire(m, v189+int32(24576), int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v200 = F_LWLockAcquire(m, v196+int32(24704), int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v207 = F_LWLockAcquire(m, v203+int32(24832), int32(1))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v214 = F_LWLockAcquire(m, v210+int32(24960), int32(1))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v221 = F_LWLockAcquire(m, v217+int32(25088), int32(1))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v228 = F_LWLockAcquire(m, v224+int32(25216), int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v103)+616))
	if v230 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v282+int32(25216))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L44
	}
L34:
	;
	F_GetSingleProcBlockerStatusData(m, v103, v20)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)+624))
	if v235 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	v239 = v230 + int32(620)
	if v235 == v239 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v241 = v235
	goto L40
L40:
	;
	F_GetSingleProcBlockerStatusData(m, v241-int32(628), v20)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L33
L42:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	if v262 != v239 {
		v241 = v262
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v288+int32(25088))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v294+int32(24960))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v300+int32(24832))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v306+int32(24704))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v312+int32(24576))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v318+int32(24448))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v324+int32(24320))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v330+int32(24192))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v336+int32(24064))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v342+int32(23936))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v348+int32(23808))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v354+int32(23680))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v360+int32(23552))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v366+int32(23424))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v372+int32(23296))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L16
L60:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v403 = F_palloc(m, v400<<(uint(int32(2))%32))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if int32(0) < v405 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v420 = v2
	v422 = v2
	goto L65
L63:
	;
	v726 = v2
	goto L64
L64:
	;
	v732 = F_construct_array_builtin(m, v403, v726, int32(23))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L114
	}
L65:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v429 = v426 + v422*int32(20)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	v433 = v425 + v430*int32(56)
	v434 = int32(0)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	if v435 <= v434 {
		v550 = v434
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v726 = v705
	goto L64
L67:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+15)))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v566<<(uint(int32(2))%32))+uint32(_consts[635])))
	goto L94
L68:
	;
	v439 = v435 & int32(3)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	v441 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v435) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v448 = v441
	v449 = int32(0)
	v451 = v434
	goto L72
L70:
	;
	v500 = v441
	v503 = v434
	goto L71
L71:
	;
	if v439 == int32(0) {
		v550 = v503
		goto L67
	} else {
		goto L87
	}
L72:
	;
	v467 = int32(56)
	v469 = v433 + (v448|int32(3))*v467
	v474 = v433 + (v448|int32(2))*v467
	v479 = v433 + (v448|int32(1))*v467
	v482 = v433 + v448*v467
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+40))
	if v483 == v440 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v500 = v496
	v503 = v494
	goto L71
L74:
	;
	v485 = v482
	goto L76
L75:
	;
	v485 = v451
	goto L76
L76:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v479)+40))
	if v486 == v440 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v488 = v479
	goto L79
L78:
	;
	v488 = v485
	goto L79
L79:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v474)+40))
	if v489 == v440 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v491 = v474
	goto L82
L81:
	;
	v491 = v488
	goto L82
L82:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v469)+40))
	if v492 == v440 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v494 = v469
	goto L85
L84:
	;
	v494 = v491
	goto L85
L85:
	;
	v495 = int32(4)
	v496 = v448 + v495
	v498 = v449 + v495
	if v498 != v435&int32(2147483644) {
		v448 = v496
		v449 = v498
		v451 = v494
		goto L72
	} else {
		goto L86
	}
L86:
	;
	goto L73
L87:
	;
	v519 = v500
	v522 = v503
	v525 = v441
	goto L88
L88:
	;
	v538 = v433 + v519*int32(56)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+40))
	if v539 == v440 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v550 = v541
	goto L67
L90:
	;
	v541 = v538
	goto L92
L91:
	;
	v541 = v522
	goto L92
L92:
	;
	v542 = int32(1)
	v545 = v525 + v542
	if v545 != v439 {
		v519 = v519 + v542
		v522 = v541
		v525 = v545
		goto L88
	} else {
		goto L93
	}
L93:
	;
	goto L89
L94:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	if int32(0) < v572 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v575 = int32(2)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v571)+4))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v550)+20))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v578+v579<<(uint(v575)%32))))
	v585 = int32(0)
	v591 = v572
	v597 = v420
	goto L98
L96:
	;
	v705 = v420
	goto L97
L97:
	;
	v711 = v422 + int32(1)
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v711 < v712 {
		v420 = v705
		v422 = v711
		goto L65
	} else {
		goto L113
	}
L98:
	;
	v604 = v433 + v585*int32(56)
	if v604 == v550 {
		v679 = v591
		v685 = v597
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v705 = v685
	goto L97
L100:
	;
	v691 = v585 + int32(1)
	if v691 < v679 {
		v585 = v691
		v591 = v679
		v597 = v685
		goto L98
	} else {
		goto L112
	}
L101:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v604)+44))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v550)+44))
	if v606 == v607 {
		v679 = v591
		v685 = v597
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v604)+16))
	if v609&v583 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403+v597<<(uint(int32(2))%32)))) = v606
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
	v679 = v672
	v685 = v597 + int32(1)
	goto L100
L104:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v604)+20))
	if v611 == int32(0) {
		v679 = v591
		v685 = v597
		goto L100
	} else {
		goto L105
	}
L105:
	;
	if int32(base.Ui32(v583)>>(uint(v611)%32))&int32(1) == int32(0) {
		v679 = v591
		v685 = v597
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v429)+16))
	if v619 <= int32(0) {
		v679 = v591
		v685 = v597
		goto L100
	} else {
		goto L107
	}
L107:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v604)+40))
	v626 = int32(0)
	goto L108
L108:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v565+v564<<(uint(v575)%32)+v626<<(uint(int32(2))%32))))
	if v644 == v622 {
		goto L103
	} else {
		goto L110
	}
L109:
	;
	v679 = v591
	v685 = v597
	goto L100
L110:
	;
	v647 = v626 + int32(1)
	if v619 != v647 {
		v626 = v647
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	goto L99
L113:
	;
	goto L66
L114:
	;
	return v732
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
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
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(-1)
	if l0 == int32(0) {
		v178 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v178
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		v178 = v13
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l0&int32(3) == int32(0) {
		v42 = l0
		goto L6
	} else {
		goto L7
	}
L4:
	;
	if base.Ui32(int32(63)) < base.Ui32(v75) {
		v178 = v13
		goto L1
	} else {
		goto L21
	}
L5:
	;
	v75 = v67 - l0
	goto L4
L6:
	;
	v46 = v42
	goto L15
L7:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v75 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v31 = l0
	goto L11
L11:
	;
	v35 = v31 + int32(1)
	if v35&int32(3) == int32(0) {
		v42 = v35
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v67 = v35
	goto L5
L13:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 != 0 {
		v31 = v35
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = int32(-2139062144)
	if (int32(16843008)-v52|v52)&v55 == v55 {
		v46 = v46 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v61 = v46
	goto L18
L17:
	;
	goto L16
L18:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v65 != 0 {
		v61 = v61 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v67 = v61
	goto L5
L20:
	;
	goto L19
L21:
	;
	v78 = l0
	v79 = v16
	v80 = v11
	goto L22
L22:
	;
	v87 = v79 & int32(255)
	goto L24
L23:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v115)
	v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	v120 = int32(1794320)
	v121 = int32(1793680)
	goto L32
L24:
	;
	if base.B2i32(base.Ui32(v87-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v87|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32((v79-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v111 = v80
	goto L27
L27:
	;
	v113 = v78 + int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v114 != 0 {
		v78 = v113
		v79 = v114
		v80 = v111
		goto L22
	} else {
		goto L31
	}
L28:
	;
	v107 = v79 | int32(32)
	goto L30
L29:
	;
	v107 = v79
	goto L30
L30:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v80))) = uint8(v107)
	v111 = v80 + int32(1)
	goto L27
L31:
	;
	goto L23
L32:
	;
	v133 = v121 + (v120-v121)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v135 = int32(*(*int8)(unsafe.Add(mBase, uint32(v134))))
	v136 = v119 - v135
	if v136 != 0 {
		v163 = v136
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v178 = v13
	goto L1
L34:
	;
	v167 = base.B2i32(v163 < int32(0))
	if v163 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v140 == int32(0) {
		v159 = v139
		v160 = v140
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v161 != 0 {
		v163 = v161
		goto L34
	} else {
		goto L44
	}
L37:
	;
	v161 = v160 - v159
	goto L36
L38:
	;
	if v139 != v140 {
		v159 = v139
		v160 = v140
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v144 = v11
	v145 = v134
	goto L40
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	if v149 == int32(0) {
		v159 = v148
		v160 = v149
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v159 = v148
	v160 = v149
	goto L37
L42:
	;
	v152 = int32(1)
	if v148 == v149 {
		v144 = v144 + v152
		v145 = v145 + v152
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v178 = v162
	goto L1
L45:
	;
	v168 = v133 - int32(8)
	goto L47
L46:
	;
	v168 = v120
	goto L47
L47:
	;
	if v163 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v171 = v121
	goto L50
L49:
	;
	v171 = v133 + int32(8)
	goto L50
L50:
	;
	if base.Ui32(v171) <= base.Ui32(v168) {
		v120 = v168
		v121 = v171
		goto L32
	} else {
		goto L51
	}
L51:
	;
	goto L33
}
func F_pg_checksum_init(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
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
			if v12 == int32(0) {
				v73 = int32(-1)
				return v73
			} else {
				v20 = int32(0)
				v21 = F_pg_cryptohash_init(m, v12)
				mBase = m.M
				if v20 <= v21 {
					v73 = v20
					return v73
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v73 = int32(-1)
						return v73
					}
				}
			}
		}
	case 2:
		v28 = F_pg_cryptohash_create(m, int32(3))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
			if v28 == int32(0) {
				v73 = int32(-1)
				return v73
			} else {
				v34 = int32(0)
				v35 = F_pg_cryptohash_init(m, v28)
				mBase = m.M
				if v34 <= v35 {
					v73 = v34
					return v73
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v73 = int32(-1)
						return v73
					}
				}
			}
		}
	case 3:
		v42 = F_pg_cryptohash_create(m, int32(4))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42
			if v42 == int32(0) {
				v73 = int32(-1)
				return v73
			} else {
				v48 = int32(0)
				v49 = F_pg_cryptohash_init(m, v42)
				mBase = m.M
				if v48 <= v49 {
					v73 = v48
					return v73
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v73 = int32(-1)
						return v73
					}
				}
			}
		}
	case 4:
		v56 = F_pg_cryptohash_create(m, int32(5))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56
			if v56 == int32(0) {
				v73 = int32(-1)
				return v73
			} else {
				v62 = int32(0)
				v63 = F_pg_cryptohash_init(m, v56)
				mBase = m.M
				if v62 <= v63 {
					v73 = v62
					return v73
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v66)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v73 = int32(-1)
						return v73
					}
				}
			}
		}
	default:
		v73 = int32(0)
		return v73
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
	v7 = int32(350163)
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
	v55 = int32(462127)
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
	v103 = int32(519141)
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
	v151 = int32(517582)
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
	v199 = int32(518917)
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
	v249 = int32(520466)
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
	var v36 int32
	_ = v36
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
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
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L49
	}
L2:
	;
	switch v36 + int32(2) {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		v154 = v36
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
	v36 = v34
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
	v36 = v20
	goto L2
L11:
	;
	m.G0 = v8 + int32(16)
	return v154
L12:
	;
	if v10&int32(3) == int32(0) {
		v115 = v10
		goto L34
	} else {
		goto L35
	}
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
	v154 = v91
	goto L11
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v91 = int32(base.Ui32(v83) >> (uint(int32(2)) % 32))
	goto L14
L16:
	;
	v91 = int32(base.Ui32(v77) >> (uint(int32(1)) % 32))
	goto L14
L17:
	;
	v44 = v10
	goto L20
L18:
	;
	v68 = v39
	v69 = v10
	goto L19
L19:
	;
	if v68&int32(1) == int32(0) {
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
	v68 = v64
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
		v77 = int32(1)
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
	v77 = v68
	goto L16
L32:
	;
	v154 = v148 + int32(1)
	goto L11
L33:
	;
	v148 = v140 - v10
	goto L32
L34:
	;
	v119 = v115
	goto L43
L35:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v99 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v148 = int32(0)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v104 = v10
	goto L39
L39:
	;
	v108 = v104 + int32(1)
	if v108&int32(3) == int32(0) {
		v115 = v108
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v140 = v108
	goto L33
L41:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v113 != 0 {
		v104 = v108
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v128 = int32(-2139062144)
	if (int32(16843008)-v125|v125)&v128 == v128 {
		v119 = v119 + int32(4)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v134 = v119
	goto L46
L45:
	;
	goto L44
L46:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v138 != 0 {
		v134 = v134 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v140 = v134
	goto L33
L48:
	;
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	F_errmsg_internal(m, int32(47385), v8)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(469433), int32(5289), int32(319553))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
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
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
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
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
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
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
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
	v26 = F_pstrdup(m, int32(493323))
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
	v29 = int32(4421136)
	goto L8
L5:
	;
	v148 = F_strlen(m, v18)
	mBase = m.M
	v155 = v148 + int32(1)
	goto L39
L6:
	;
	v142 = F_strlen(m, v131)
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
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v135)
	goto L6
L11:
	;
	v116 = v111
	v117 = v112
	v118 = v113
	goto L33
L12:
	;
	if v106 == int32(0) {
		v131 = v104
		v132 = v105
		goto L10
	} else {
		goto L32
	}
L13:
	;
	v104 = v29
	v105 = v18
	v106 = v36
	goto L12
L14:
	;
	goto L15
L15:
	;
	goto L17
L16:
	;
	goto L25
L17:
	;
	goto L16
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _consts[946])))
	if v76 == int32(0) {
		v104 = v29
		v105 = v18
		v106 = v36
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L27
L27:
	;
	v82 = v29
	v83 = v18
	v84 = v36
	goto L28
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 != v90 {
		v111 = v82
		v112 = v83
		v113 = v84
		goto L11
	} else {
		goto L30
	}
L29:
	;
	v104 = v98
	v105 = v96
	v106 = v100
	goto L12
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v87
	v95 = int32(4)
	v96 = v83 + v95
	v98 = v82 + v95
	v100 = v84 - v95
	if base.Ui32(int32(3)) < base.Ui32(v100) {
		v82 = v98
		v83 = v96
		v84 = v100
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v111 = v104
	v112 = v105
	v113 = v106
	goto L11
L33:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v120)
	if v120 == int32(0) {
		v131 = v116
		v132 = v117
		goto L10
	} else {
		goto L35
	}
L34:
	;
	v131 = v127
	v132 = v125
	goto L10
L35:
	;
	v124 = int32(1)
	v125 = v117 + v124
	v127 = v116 + v124
	v129 = v118 - v124
	if v129 != 0 {
		v116 = v127
		v117 = v125
		v118 = v129
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v167 != 0 {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v157 = int32(0)
	if v155 == v157 {
		v167 = v157
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v167 = v162
	goto L38
L41:
	;
	v161 = v155 - int32(1)
	v162 = v18 + v161
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	if v163 != int32(47) {
		v155 = v161
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v168)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v170 = F_pstrdup(m, v18)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v170
	v174 = F_pstrdup(m, int32(493391))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v174
	F_get_doc_path(m, v18)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v179 = F_pstrdup(m, v18)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v179
	v183 = F_pstrdup(m, int32(493337))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v183
	F_get_doc_path(m, v18)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v188 = F_pstrdup(m, v18)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v188
	v192 = F_pstrdup(m, int32(493380))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v192
	F_make_relative_path(m, v18, int32(386167), int32(4421136))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v199 = F_pstrdup(m, v18)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v199
	v203 = F_pstrdup(m, int32(493377))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v203
	F_make_relative_path(m, v18, int32(282210), int32(4421136))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v210 = F_pstrdup(m, v18)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v210
	v214 = F_pstrdup(m, int32(493408))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v214
	F_make_relative_path(m, v18, int32(201362), int32(4421136))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v221 = F_pstrdup(m, v18)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v221
	v225 = F_pstrdup(m, int32(493401))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v225
	F_make_relative_path(m, v18, int32(472128), int32(4421136))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v232 = F_pstrdup(m, v18)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v232
	v236 = F_pstrdup(m, int32(493398))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v236
	F_get_pkglib_path(m, v18)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v241 = F_pstrdup(m, v18)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+60)) = v241
	v247 = F_pstrdup(m, int32(493367))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23-int32(-64)))) = v247
	F_make_relative_path(m, v18, int32(373979), int32(4421136))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v254 = F_pstrdup(m, v18)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v254
	v258 = F_pstrdup(m, int32(493330))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v258
	F_make_relative_path(m, v18, int32(265527), int32(4421136))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v265 = F_pstrdup(m, v18)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v265
	v269 = F_pstrdup(m, int32(493358))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v269
	F_get_share_path(m, v18)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v274 = F_pstrdup(m, v18)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v274
	v278 = F_pstrdup(m, int32(493347))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v278
	F_get_etc_path(m, int32(4421136), v18)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v284 = F_pstrdup(m, v18)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+92)) = v284
	v288 = F_pstrdup(m, int32(490600))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v288
	F_get_pkglib_path(m, v18)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v293 = int32(296621)
	v294 = int32(1024)
	v296 = F_pg_ascii_verifystr(m, v18, v294)
	mBase = m.M
	if v296 == v294 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v303 = F_pstrdup(m, v18)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L87
	}
L83:
	;
	goto L82
L84:
	;
	v298 = F_strlen(m, v293)
	mBase = m.M
	goto L83
L85:
	;
	goto L86
L86:
	;
	v301 = F_strlcpy(m, v18+v296, v293, v294-v296)
	mBase = m.M
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v303
	v307 = F_pstrdup(m, int32(507416))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v307
	v311 = F_pstrdup(m, int32(636997))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = v311
	v315 = F_pstrdup(m, int32(512479))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v315
	v319 = F_pstrdup(m, int32(461925))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v319
	v323 = F_pstrdup(m, int32(491520))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+120)) = v323
	v327 = F_pstrdup(m, int32(510087))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+124)) = v327
	v331 = F_pstrdup(m, int32(491537))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v331
	v335 = F_pstrdup(m, int32(197936))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v335
	v339 = F_pstrdup(m, int32(500173))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = v339
	v343 = F_pstrdup(m, int32(512306))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+140)) = v343
	v347 = F_pstrdup(m, int32(491529))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v347
	v351 = F_pstrdup(m, int32(270256))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v351
	v355 = F_pstrdup(m, int32(478286))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v355
	v359 = F_pstrdup(m, int32(227657))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v359
	v363 = F_pstrdup(m, int32(500162))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v363
	v367 = F_pstrdup(m, int32(706478))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = v367
	v371 = F_pstrdup(m, int32(492434))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = v371
	v375 = F_pstrdup(m, int32(689047))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+172)) = v375
	v379 = F_pstrdup(m, int32(498113))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v379
	v383 = F_pstrdup(m, int32(519765))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v383
	m.G0 = v18 + int32(1024)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if v389 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v391 = int32(0)
	goto L111
L109:
	;
	goto L110
L110:
	;
	m.G0 = v8 + int32(32)
	return int32(0)
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(0)
	v398 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v398)
	v402 = v23 + v391<<(uint(int32(3))%32)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v404 = F_cstring_to_text(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v404
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v408 = F_cstring_to_text(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v408
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v411, v412, v8+int32(16), v8+int32(14))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v420 = v391 + int32(1)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	if base.Ui32(v420) < base.Ui32(v421) {
		v391 = v420
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L112
}
func F_pg_control_checkpoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v7 = m.G0
	v9 = v7 - int32(224)
	m.G0 = v9
	v14 = F_get_call_result_type(m, l0, int32(0), v9+int32(108))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(1) {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			v25 = F_LWLockAcquire(m, v21+int32(1152), int32(1))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[967]))
				v31 = F_get_controlfile(m, v28, v9+int32(31))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[29]))
					F_LWLockRelease(m, v34+int32(1152))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
						if v39 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v187 = m.ExcPending
							if v187 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(363935), int32(0))
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(469410), int32(90), int32(82687))
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v31)+40))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v43
							v46 = int64(*(*int32)(unsafe.Add(mBase, _consts[167])))
							v47 = base.I64_div_u_s(v42, v46)
							v49 = base.I64_div_u_s(int64(4294967296), v46)
							v50 = base.I64_div_u_s(v47, v49)
							*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v50)
							v53 = v47 - v49*v50
							*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)) = uint32(v53)
							v61 = F_pg_snprintf(m, v9+int32(32), int32(64), int32(478409), v9+int32(16))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
								v64 = F_Int64GetDatum(m, v63)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v66 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+112)) = uint8(v66)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v64
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v31)+40))
									v70 = F_Int64GetDatum(m, v69)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v72 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+113)) = uint8(v72)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v70
										v77 = F_cstring_to_text(m, v9+int32(32))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+114)) = uint8(v79)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v77
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+115)) = uint8(v79)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = v82
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+116)) = uint8(v79)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v86
											v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+56)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+117)) = uint8(v79)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v90
											v94 = *(*int64)(unsafe.Add(mBase, uint32(v31)+64))
											*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v94)
											v97 = int64(base.Ui64(v94) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v97)
											v100 = F_psprintf(m, int32(35756), v9)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												v102 = F_cstring_to_text(m, v100)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													v104 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+118)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = v102
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+119)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v107
													v111 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+120)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v111
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v31)+80))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+121)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v115
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v31)+84))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+122)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+184)) = v119
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v31)+88))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+123)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+188)) = v123
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v31)+120))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+124)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+192)) = v127
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v31)+92))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+125)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v131
													v135 = *(*int32)(unsafe.Add(mBase, uint32(v31)+96))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+126)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v135
													v139 = *(*int32)(unsafe.Add(mBase, uint32(v31)+112))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+127)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+204)) = v139
													v143 = *(*int32)(unsafe.Add(mBase, uint32(v31)+116))
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+128)) = uint8(v104)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+208)) = v143
													v147 = *(*int64)(unsafe.Add(mBase, uint32(v31)+104))
													v152 = F_Int64GetDatum(m, v147*int64(1000000)-int64(946684800000000))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														v154 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v9)+129)) = uint8(v154)
														*(*int32)(unsafe.Add(mBase, uint32(v9)+212)) = v152
														v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
														v162 = F_heap_form_tuple(m, v157, v9+int32(144), v9+int32(112))
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return int32(0)
														} else {
															v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+16))
															v165 = F_HeapTupleHeaderGetDatum(m, v164)
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(224)
																return v165
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
			v174 = m.ExcPending
			if v174 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(345117), int32(0))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(469410), int32(82), int32(82687))
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
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
		return int32(12790)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v9 == int32(1) {
			v12 = int32(286010)
		} else {
			v12 = int32(119679)
		}
		if v9 == int32(2) {
			v15 = int32(12790)
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
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v3&int32(3) != 0 {
		v6 = F_detoast_attr(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(base.Ui32(v11) >> (uint(int32(2)) % 32))
		v14 = F_palloc(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				v16 = F__emscripten_memcpy_bulkmem(m, v14, l0, v13)
				mBase = m.M
				v17 = v16
			} else {
				v17 = v14
			}
			return v17
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
	v46 = v42<<(uint(int32(8))%32) | int32(9371648)
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[266]))
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
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(633836)
					F_errmsg(m, int32(236931), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(465119), int32(1651), int32(230038))
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
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(633836)
				F_errmsg(m, int32(236931), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(465119), int32(1651), int32(230038))
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
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v4 < int32(0) {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if base.Ui32((v9-int32(48))&int32(255)) < base.Ui32(int32(10)) {
			v16 = int32(4)
		} else {
			v16 = int32(2)
		}
		v17 = v16
	} else {
		v17 = int32(1)
	}
	return v17
}
func F_pg_get_client_encoding(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[953]))
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
	var v23 int32
	_ = v23
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
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v17 = v3
	v18 = int32(-1)
	v19 = v3
	v20 = v3
	v23 = v3
	goto L4
L1:
	;
	m.G0 = v12 + int32(16)
	return v202
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v183
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188+v183))) = uint8(v190)
	v202 = v190
	goto L1
L3:
	;
	v202 = int32(1)
	goto L1
L4:
	;
	if v18 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if int32(base.Ui32(v143)>>(uint(int32(5))%32))&int32(1) != 0 {
		v182 = v35
		v183 = v36
		goto L2
	} else {
		goto L64
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
	if v17 == int32(0) {
		v35 = v19
		v36 = v20
		v37 = v23
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, 8)) = uint8(v33)
	v182 = v19
	v183 = v20
	goto L2
L11:
	;
	goto L14
L12:
	;
	goto L5
L13:
	;
	v150 = int32(m.ExcTag)
	v151 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v150 == int32(0) {
		goto L53
	} else {
		goto L54
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
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v138 < int32(0) {
		goto L50
	} else {
		goto L51
	}
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
	v68 = v66 + v67
	if v68&int32(3) == int32(0) {
		v92 = v68
		goto L28
	} else {
		goto L29
	}
L24:
	;
	goto L25
L25:
	;
	goto L15
L26:
	;
	v126 = v125 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v126
	if v36 < v126 {
		goto L43
	} else {
		goto L44
	}
L27:
	;
	v125 = v117 - v68
	goto L26
L28:
	;
	v96 = v92
	goto L37
L29:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v76 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v125 = int32(0)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v81 = v68
	goto L33
L33:
	;
	v85 = v81 + int32(1)
	if v85&int32(3) == int32(0) {
		v92 = v85
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v117 = v85
	goto L27
L35:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v90 != 0 {
		v81 = v85
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v105 = int32(-2139062144)
	if (int32(16843008)-v102|v102)&v105 == v105 {
		v96 = v96 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v111 = v96
	goto L40
L39:
	;
	goto L38
L40:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v115 != 0 {
		v111 = v111 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v117 = v111
	goto L27
L42:
	;
	goto L41
L43:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v66-int32(1)))))
	if v132 == int32(10) {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_enlargeStringInfo(m, l1, int32(128))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	goto L14
L48:
	;
	goto L12
L49:
	;
	goto L48
L50:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = v141
	goto L49
L51:
	;
	goto L52
L52:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v143 = v142
	goto L49
L53:
	;
	v155 = int32(v151)
	m.G0 = v12
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v12+int32(12) == v162 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	m.ExcPending = 1
	goto L62
L55:
	;
	if v165 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v165 = v164
	goto L58
L57:
	;
	v165 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	F___wasm_longjmp(m, v158, v157)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v17 = v157
	v18 = v165
	v19 = v35
	v20 = v36
	v23 = v37
	goto L4
L62:
	;
	return int32(0)
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v173 == v36 {
		v182 = v35
		v183 = v36
		goto L2
	} else {
		goto L65
	}
L65:
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
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
				v20 = int32(4425280)
				v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
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
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(345117), int32(0))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(463158), int32(3645), int32(125023))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
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
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v21
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									if v56 < v57 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v56<<(uint(int32(3))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v63
										v68 = F_psprintf(m, int32(56389), v9+int32(32))
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
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
													F_errmsg_internal(m, int32(440341), v9+int32(16))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(463158), int32(1834), int32(309308))
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(int32(2))%32))+uint32(_consts[75])))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v83
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
												v88 = F_BuildTupleFromCStrings(m, v85, v9+int32(40))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v90 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v90 + int32(1)
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
													F_pfree(m, v94)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v97 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
														*(*int64)(unsafe.Add(mBase, uint32(v54))) = v97 + int64(1)
														v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
														v105 = F_HeapTupleHeaderGetDatum(m, v104)
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v116 = v105
															m.G0 = v9 + int32(48)
															return v116
														}
													}
												}
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = int32(2)
											v112 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v112)
											v116 = int32(0)
											m.G0 = v9 + int32(48)
											return v116
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
				v68 = F_psprintf(m, int32(56389), v9+int32(32))
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
						v154 = m.ExcPending
						if v154 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
							F_errmsg_internal(m, int32(440341), v9+int32(16))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(463158), int32(1834), int32(309308))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(int32(2))%32))+uint32(_consts[75])))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v83
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
						v88 = F_BuildTupleFromCStrings(m, v85, v9+int32(40))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v90 + int32(1)
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
							F_pfree(m, v94)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								v97 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
								*(*int64)(unsafe.Add(mBase, uint32(v54))) = v97 + int64(1)
								v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
								v105 = F_HeapTupleHeaderGetDatum(m, v104)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v116 = v105
									m.G0 = v9 + int32(48)
									return v116
								}
							}
						}
					}
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = int32(2)
					v112 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v112)
					v116 = int32(0)
					m.G0 = v9 + int32(48)
					return v116
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v124 = m.ExcPending
		if v124 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
				F_errmsg(m, int32(55925), v9)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(463158), int32(3628), int32(125023))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[991])))
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(18280)+v13<<(uint(int32(4))%32))))
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
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[849]))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L122
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L119
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L116
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+52)) = int64(81604378650)
	v33 = F_SPI_prepare(m, int32(520516), int32(2), v10+int32(-12))
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
	v46 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), int32(496476))
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
	*(*int32)(unsafe.Add(mBase, _consts[849])) = v33
	goto L9
L13:
	;
	v48 = int32(8224)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+50)) = uint16(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v46
	v52 = *(*int32)(unsafe.Add(mBase, _consts[849]))
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
	v62 = *(*int64)(unsafe.Add(mBase, _consts[349]))
	if v62 != int64(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v386 = F_SPI_finish(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L111
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[350]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = int32(343967)
	v71 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v71 < v74 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v122 = F_SPI_getbinval(m, v69, v67, v119, v10+int32(-1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L33
	}
L19:
	;
	v119 = v81 + int32(1)
	goto L18
L20:
	;
	v81 = v71
	v82 = v74
	goto L23
L21:
	;
	goto L22
L22:
	;
	v107 = F_SystemAttributeByName(m, v70)
	mBase = m.M
	if v107 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v84 = int32(4)
	v89 = v67 + int32(20) + v82<<(uint(v84)%32) + v81*int32(100)
	v92 = F_namestrcmp(m, v89+v84, v70)
	mBase = m.M
	if v92 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+91)))
	if v95 != int32(1) {
		goto L19
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v99 = v81 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v99 < v100 {
		v81 = v99
		v82 = v100
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
	v119 = int32(-9)
	goto L18
L31:
	;
	goto L32
L32:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107)+74)))
	v119 = v111
	goto L18
L33:
	;
	v124 = int32(120066)
	v125 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v125 < v128 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v176 = F_SPI_getbinval(m, v69, v67, v173, v10+int32(-1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L49
	}
L35:
	;
	v173 = v135 + int32(1)
	goto L34
L36:
	;
	v135 = v125
	v136 = v128
	goto L39
L37:
	;
	goto L38
L38:
	;
	v161 = F_SystemAttributeByName(m, v124)
	mBase = m.M
	if v161 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v138 = int32(4)
	v143 = v67 + int32(20) + v136<<(uint(v138)%32) + v135*int32(100)
	v146 = F_namestrcmp(m, v143+v138, v124)
	mBase = m.M
	if v146 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+91)))
	if v149 != int32(1) {
		goto L35
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v153 = v135 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v153 < v154 {
		v135 = v153
		v136 = v154
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
	v173 = int32(-9)
	goto L34
L47:
	;
	goto L48
L48:
	;
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v161)+74)))
	v173 = v165
	goto L34
L49:
	;
	v178 = int32(436035)
	v179 = int32(0)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v179 < v182 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v230 = F_SPI_getbinval(m, v69, v67, v227, v10+int32(-1))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L65
	}
L51:
	;
	v227 = v189 + int32(1)
	goto L50
L52:
	;
	v189 = v179
	v190 = v182
	goto L55
L53:
	;
	goto L54
L54:
	;
	v215 = F_SystemAttributeByName(m, v178)
	mBase = m.M
	if v215 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v192 = int32(4)
	v197 = v67 + int32(20) + v190<<(uint(v192)%32) + v189*int32(100)
	v200 = F_namestrcmp(m, v197+v192, v178)
	mBase = m.M
	if v200 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+91)))
	if v203 != int32(1) {
		goto L51
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v207 = v189 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v207 < v208 {
		v189 = v207
		v190 = v208
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
	v227 = int32(-9)
	goto L50
L63:
	;
	goto L64
L64:
	;
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215)+74)))
	v227 = v219
	goto L50
L65:
	;
	v232 = int32(290644)
	v233 = int32(0)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v233 < v236 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v282 = F_SPI_getvalue(m, v69, v67, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L81
	}
L67:
	;
	v281 = v243 + int32(1)
	goto L66
L68:
	;
	v243 = v233
	v244 = v236
	goto L71
L69:
	;
	goto L70
L70:
	;
	v269 = F_SystemAttributeByName(m, v232)
	mBase = m.M
	if v269 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v246 = int32(4)
	v251 = v67 + int32(20) + v244<<(uint(v246)%32) + v243*int32(100)
	v254 = F_namestrcmp(m, v251+v246, v232)
	mBase = m.M
	if v254 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+91)))
	if v257 != int32(1) {
		goto L67
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v261 = v243 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v261 < v262 {
		v243 = v261
		v244 = v262
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
	v281 = int32(-9)
	goto L66
L79:
	;
	goto L80
L80:
	;
	v273 = int32(*(*int16)(unsafe.Add(mBase, uint32(v269)+74)))
	v281 = v273
	goto L66
L81:
	;
	v284 = int32(242732)
	v285 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v285 < v288 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v334 = F_SPI_getvalue(m, v69, v67, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v333 = v295 + int32(1)
	goto L82
L84:
	;
	v295 = v285
	v296 = v288
	goto L87
L85:
	;
	goto L86
L86:
	;
	v321 = F_SystemAttributeByName(m, v284)
	mBase = m.M
	if v321 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v298 = int32(4)
	v303 = v67 + int32(20) + v296<<(uint(v298)%32) + v295*int32(100)
	v306 = F_namestrcmp(m, v303+v298, v284)
	mBase = m.M
	if v306 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+91)))
	if v309 != int32(1) {
		goto L83
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v313 = v295 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v313 < v314 {
		v295 = v313
		v296 = v314
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
	v333 = int32(-9)
	goto L82
L95:
	;
	goto L96
L96:
	;
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321)+74)))
	v333 = v325
	goto L82
L97:
	;
	v336 = F_stringToNode(m, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v336 == int32(0) {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	if v340 != int32(1) {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	if v122&int32(255) != int32(49) {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	if v230 == int32(0) {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v349 != int32(60) {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v352 != int32(62) {
		goto L16
	} else {
		goto L104
	}
L104:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+2)))
	if v355 != 0 {
		goto L16
	} else {
		goto L105
	}
L105:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	if v358 != int32(1) {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	v365 = F_table_open(m, v176, int32(1))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v365)+52))
	F_get_query_def(m, v357, v10+int32(-32), int32(0), v367, int32(1), l1, l2, int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_appendStringInfoChar(m, v10+int32(-32), int32(59))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_sequence_close(m, v365, int32(1))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	goto L16
L111:
	;
	if v386 != int32(2) {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	m.G0 = v12 - int32(-64)
	if v390 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v396 = v391
	goto L115
L114:
	;
	v396 = int32(0)
	goto L115
L115:
	;
	return v396
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(520516)
	F_errmsg_internal(m, int32(649972), v12)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(463622), int32(822), int32(207308))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg_internal(m, int32(37448), v10+int32(-48))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(463622), int32(836), int32(207308))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errmsg_internal(m, int32(427137), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(463622), int32(858), int32(207308))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
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
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = v8 + int32(-16)
	v24 = int32(1603664)
	v26 = int32(9299)
	v29 = v2
	goto L3
L3:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(-38)))) = uint8(v31)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+24)) = uint16(v31)
	v39 = F_cstring_to_text(m, v26)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v68 = F_GetWaitEventCustomNames(m, int32(117440512), v8+int32(-4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v39
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v43 = F_cstring_to_text(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v47 = F_cstring_to_text(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v50, v51, v8+int32(-24), v8+int32(-40))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v59 = v29 + int32(1)
	v61 = v59 * int32(12)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_consts[767])))
	if v64 != 0 {
		v24 = v61 + int32(1603664)
		v26 = v64
		v29 = v59
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if int32(0) < v70 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v74 = v8 + int32(-32)
	v76 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v141 = F_GetWaitEventCustomNames(m, int32(184549376), v8+int32(-4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v83
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v83)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+20)) = uint16(v83)
	v92 = F_cstring_to_text(m, int32(256016))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v92
	v97 = v68 + v76<<(uint(int32(2))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v99 = F_cstring_to_text(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v99
	F_initStringInfo(m, v8+int32(-24))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
	F_appendStringInfo(m, v8+int32(-24), int32(359285), v8+int32(-48))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v116 = F_cstring_to_text(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v116
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v119, v120, v8+int32(-40), v8+int32(-44))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v128 = v76 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v128 < v129 {
		v76 = v128
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if int32(0) < v143 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v147 = v8 + int32(-32)
	v149 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	m.G0 = v10 - int32(-64)
	return int32(0)
L27:
	;
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)) = uint8(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+20)) = uint16(v156)
	v165 = F_cstring_to_text(m, int32(83645))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v165
	v170 = v141 + v149<<(uint(int32(2))%32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v172 = F_cstring_to_text(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v172
	F_initStringInfo(m, v8+int32(-24))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v179
	F_appendStringInfo(m, v8+int32(-24), int32(648160), v10)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v187 = F_cstring_to_text(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v187
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v190, v191, v8+int32(-40), v8+int32(-44))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v199 = v149 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	if v199 < v200 {
		v149 = v199
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v6 = F_palloc(m, int32(280))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v13 = F__emscripten_memset_bulkmem(m, v6, base.I32_extend8_s(int32(0)), int32(280))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l0
			if base.Ui32(l0) <= base.Ui32(int32(5)) {
				v20 = l0 << (uint(int32(2)) % 32)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[1030])))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v23
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[1031])))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v27
			} else {
			}
			v30 = F_pg_cryptohash_create(m, l0)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v30
				if v30 != 0 {
					v44 = v6
					return v44
				} else {
					v35 = F___memset(m, v13, int32(0), int32(280))
					mBase = m.M
					F_pfree(m, v13)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v44 = int32(0)
						return v44
					}
				}
			}
		} else {
			v44 = int32(0)
			return v44
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
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
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v20 = F__emscripten_memset_bulkmem(m, v7, base.I32_extend8_s(int32(0)), v18)
				mBase = m.M
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v23 = F_pg_cryptohash_final(m, v21, v20, v22)
				mBase = m.M
				if v23 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v28 == int32(0) {
						v43 = int32(12790)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
						if v35 == int32(1) {
							v38 = int32(286010)
						} else {
							v38 = int32(119679)
						}
						if v35 == int32(2) {
							v41 = int32(12790)
						} else {
							v41 = v38
						}
						v43 = v41
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v43
					F_pfree(m, v20)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						return int32(-1)
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v50 = F_pg_cryptohash_init(m, v49)
					mBase = m.M
					if v50 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v71 == int32(0) {
							v86 = int32(12790)
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
							if v78 == int32(1) {
								v81 = int32(286010)
							} else {
								v81 = int32(119679)
							}
							if v78 == int32(2) {
								v84 = int32(12790)
							} else {
								v84 = v81
							}
							v86 = v84
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v86
						F_pfree(m, v20)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							return int32(-1)
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v57 = F_pg_cryptohash_update(m, v53, l0+int32(152), v56)
						mBase = m.M
						if v57 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v71 == int32(0) {
								v86 = int32(12790)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
								if v78 == int32(1) {
									v81 = int32(286010)
								} else {
									v81 = int32(119679)
								}
								if v78 == int32(2) {
									v84 = int32(12790)
								} else {
									v84 = v81
								}
								v86 = v84
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v86
							F_pfree(m, v20)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								return int32(-1)
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v62 = F_pg_cryptohash_update(m, v60, v20, v61)
							mBase = m.M
							if v62 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v71 == int32(0) {
									v86 = int32(12790)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
									if v78 == int32(1) {
										v81 = int32(286010)
									} else {
										v81 = int32(119679)
									}
									if v78 == int32(2) {
										v84 = int32(12790)
									} else {
										v84 = v81
									}
									v86 = v84
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v86
								F_pfree(m, v20)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									return int32(-1)
								}
							} else {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v66 = F_pg_cryptohash_final(m, v65, l1, l2)
								mBase = m.M
								if int32(0) <= v66 {
									F_pfree(m, v20)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int32(0)
									} else {
										v96 = int32(0)
										return v96
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v71 == int32(0) {
										v86 = int32(12790)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
										if v78 == int32(1) {
											v81 = int32(286010)
										} else {
											v81 = int32(119679)
										}
										if v78 == int32(2) {
											v84 = int32(12790)
										} else {
											v84 = v81
										}
										v86 = v84
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v86
									F_pfree(m, v20)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
		v96 = int32(-1)
		return v96
	}
}
func F_pg_hmac_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
			v7 = F___memset(m, l0, int32(0), int32(280))
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v10
	v18 = F_get_call_result_type(m, l0, int32(0), v7+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(1) {
			v27 = F_getObjectTypeDescription(m, v7+int32(36), int32(1))
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
					v41 = F_getObjectIdentityParts(m, v7+int32(36), v7+int32(32), v7+int32(28), int32(1))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						if v41 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v45)
							v70 = v45
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v70)
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
							v77 = F_heap_form_tuple(m, v72, v7+int32(16), v7+int32(13))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
								v80 = F_HeapTupleHeaderGetDatum(m, v79)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(48)
									return v80
								}
							}
						} else {
							F_pfree(m, v41)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
								if v50 != 0 {
									v51 = F_strlist_to_textarray(m, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										v56 = v51
										v57 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v57)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v56
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										if v60 != 0 {
											v61 = F_strlist_to_textarray(m, v60)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												v66 = v61
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v66
												v70 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v70)
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v77 = F_heap_form_tuple(m, v72, v7+int32(16), v7+int32(13))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
													v80 = F_HeapTupleHeaderGetDatum(m, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v80
													}
												}
											}
										} else {
											v64 = F_construct_empty_array(m, int32(25))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v66 = v64
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v66
												v70 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v70)
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v77 = F_heap_form_tuple(m, v72, v7+int32(16), v7+int32(13))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
													v80 = F_HeapTupleHeaderGetDatum(m, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v80
													}
												}
											}
										}
									}
								} else {
									v54 = F_construct_empty_array(m, int32(25))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = v54
										v57 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v57)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v56
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
										if v60 != 0 {
											v61 = F_strlist_to_textarray(m, v60)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												v66 = v61
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v66
												v70 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v70)
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v77 = F_heap_form_tuple(m, v72, v7+int32(16), v7+int32(13))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
													v80 = F_HeapTupleHeaderGetDatum(m, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v80
													}
												}
											}
										} else {
											v64 = F_construct_empty_array(m, int32(25))
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v66 = v64
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v66
												v70 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v70)
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												v77 = F_heap_form_tuple(m, v72, v7+int32(16), v7+int32(13))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
													v80 = F_HeapTupleHeaderGetDatum(m, v79)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 + int32(48)
														return v80
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
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(345117), int32(0))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(463399), int32(4384), int32(117833))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
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
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v523 int32
	_ = v523
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
	return v523
L2:
	;
	v499 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L133
	}
L3:
	;
	v482 = F_ClosePipeStream(m, v36)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L6
	} else {
		goto L132
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L6
	} else {
		goto L128
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L6
	} else {
		goto L124
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
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L120
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
	v36 = F_OpenPipeStream(m, int32(475797), int32(216904))
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
	v53 = v2
	v55 = v32
	v56 = int32(100)
	v59 = v2
	v60 = v2
	goto L18
L18:
	;
	v62 = v16 + int32(224)
	if v62&int32(3) == int32(0) {
		v86 = v62
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v366 = F_ClosePipeStream(m, v36)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L6
	} else {
		goto L102
	}
L20:
	;
	v364 = F_fgets(m, v16+int32(224), int32(128), v36)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L100
	}
L21:
	;
	v146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v146)
	v151 = v16 + int32(224)
	goto L48
L22:
	;
	if v119 != 0 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v119 = v111 - v62
	goto L22
L24:
	;
	v90 = v86
	goto L33
L25:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v70 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v119 = int32(0)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v75 = v62
	goto L29
L29:
	;
	v79 = v75 + int32(1)
	if v79&int32(3) == int32(0) {
		v86 = v79
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v111 = v79
	goto L23
L31:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v84 != 0 {
		v75 = v79
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v99 = int32(-2139062144)
	if (int32(16843008)-v96|v96)&v99 == v99 {
		v90 = v90 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v105 = v90
	goto L36
L35:
	;
	goto L34
L36:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v109 != 0 {
		v105 = v105 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v111 = v105
	goto L23
L38:
	;
	goto L37
L39:
	;
	v122 = v119 + v16 + int32(223)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v123 == int32(10) {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v129 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	if v129 == int32(0) {
		v353 = v53
		v355 = v55
		v356 = v56
		v359 = v59
		v360 = v60
		goto L20
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(224)
	F_errmsg_internal(m, int32(676417), v16+int32(16))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(464016), int32(885), int32(132041))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v353 = v53
	v355 = v55
	v356 = v56
	v359 = v59
	v360 = v60
	goto L20
L47:
	;
	if base.B2i32(v153 == int32(0)) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v151))))
	if int32(0) < v153 {
		v151 = v151 + int32(1)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	goto L49
L51:
	;
	v164 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v184 = F_pg_get_encoding_from_locale(m, v16+int32(224), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L58
	}
L54:
	;
	if v164 == int32(0) {
		v353 = v53
		v355 = v55
		v356 = v56
		v359 = v59
		v360 = v60
		goto L20
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v16 + int32(224)
	F_errmsg_internal(m, int32(676593), v16-int32(-64))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(464016), int32(710), int32(373848))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v353 = v53
	v355 = v55
	v356 = v56
	v359 = v59
	v360 = v60
	goto L20
L58:
	;
	if v184 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v190 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(int32(35)) <= base.Ui32(v184) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	if v190 == int32(0) {
		v353 = v53
		v355 = v55
		v356 = v56
		v359 = v59
		v360 = v60
		goto L20
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v16 + int32(224)
	F_errmsg_internal(m, int32(675926), v16+int32(32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(464016), int32(717), int32(373848))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v353 = v53
	v355 = v55
	v356 = v56
	v359 = v59
	v360 = v60
	goto L20
L66:
	;
	v211 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if v184 == int32(0) {
		v353 = v53
		v355 = v55
		v356 = v56
		v359 = v59
		v360 = v60
		goto L20
	} else {
		goto L73
	}
L69:
	;
	if v211 == int32(0) {
		v353 = v53
		v355 = v55
		v356 = v56
		v359 = v59
		v360 = v60
		goto L20
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v16 + int32(224)
	F_errmsg_internal(m, int32(675878), v16+int32(48))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(464016), int32(722), int32(373848))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v353 = v53
	v355 = v55
	v356 = v56
	v359 = v59
	v360 = v60
	goto L20
L73:
	;
	v231 = v16 + int32(224)
	v233 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v234 = int32(99)
	v240 = int32(0)
	v245 = F_get_collation_actual_version(m, v234, v231)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	v247 = int32(1)
	v249 = F_CollationCreate(m, v231, v18, v233, v234, int32(1), v184, v231, v231, v240, v240, v245, v247, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	if v249 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	v255 = v59
	goto L78
L78:
	;
	v257 = v60 + int32(1)
	v263 = v16 + int32(224)
	v267 = int32(0)
	v269 = v16 + int32(96)
	goto L80
L79:
	;
	v255 = v59 + int32(1)
	goto L78
L80:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v276 != int32(46) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v317)
	if v267&int32(1) == v317 {
		v353 = v53
		v355 = v55
		v356 = v56
		v359 = v255
		v360 = v257
		goto L20
	} else {
		goto L93
	}
L82:
	;
	if v276 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v284 = v263
	goto L88
L84:
	;
	goto L81
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v276)
	v280 = int32(1)
	v263 = v263 + v280
	v269 = v269 + v280
	goto L80
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	v298 = v284 + int32(1)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if base.Ui32((v299&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		v284 = v298
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v263 = v298
	v267 = int32(1)
	goto L80
L90:
	;
	if v299 == int32(45) {
		v284 = v298
		goto L88
	} else {
		goto L91
	}
L91:
	;
	if base.Ui32((v299-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v284 = v298
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	if v56 <= v53 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v326 = F_repalloc(m, v55, v56*int32(24))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L97
	}
L95:
	;
	v330 = v55
	v331 = v56
	goto L96
L96:
	;
	v334 = v330 + v53*int32(12)
	v337 = F_pstrdup(m, v16+int32(224))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L98
	}
L97:
	;
	v330 = v326
	v331 = v56 << (uint(int32(1)) % 32)
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334))) = v337
	v342 = F_pstrdup(m, v16+int32(96))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+8)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v334)+4)) = v342
	v353 = v53 + int32(1)
	v355 = v330
	v356 = v331
	v359 = v255
	v360 = v257
	goto L20
L100:
	;
	if v364 != 0 {
		v53 = v353
		v55 = v355
		v56 = v356
		v59 = v359
		v60 = v360
		goto L18
	} else {
		goto L101
	}
L101:
	;
	goto L19
L102:
	;
	if int32(2) <= v353 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v360 != 0 {
		v523 = v430
		goto L1
	} else {
		goto L119
	}
L104:
	;
	v381 = int32(0)
	v390 = v359
	goto L110
L105:
	;
	F_pg_qsort(m, v355, v353, int32(12), int32(516))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v374 = int32(1)
	if v353 != v374 {
		v430 = v359
		goto L103
	} else {
		goto L109
	}
L108:
	;
	v377 = v353
	goto L104
L109:
	;
	v377 = v374
	goto L104
L110:
	;
	v394 = v355 + v381*int32(12)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v398 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v399 = int32(99)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v402 = int32(0)
	v405 = F_get_collation_actual_version(m, v399, v401)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L6
	} else {
		goto L112
	}
L111:
	;
	v430 = v415
	goto L103
L112:
	;
	v407 = int32(1)
	v409 = F_CollationCreate(m, v396, v18, v398, v399, int32(1), v395, v401, v401, v402, v402, v405, v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	if v409 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	v415 = v390
	goto L116
L116:
	;
	v417 = v381 + int32(1)
	if v417 != v377 {
		v381 = v417
		v390 = v415
		goto L110
	} else {
		goto L118
	}
L117:
	;
	v415 = v390 + int32(1)
	goto L116
L118:
	;
	goto L111
L119:
	;
	v495 = v430
	goto L2
L120:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(132100), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(464016), int32(844), int32(132041))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v18
	F_errmsg(m, int32(65489), v16+int32(80))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L6
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(464016), int32(849), int32(132041))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L6
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(475797)
	F_errmsg(m, int32(281660), v16)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(464016), int32(873), int32(132041))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v495 = v2
	goto L2
L133:
	;
	if v499 == int32(0) {
		v523 = v495
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(398303), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(464016), int32(964), int32(132041))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	v523 = v495
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v12)
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v15 = F_LWLockAcquire(m, v11+int32(4992), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[64]))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
		if v21 != 0 {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v26 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v26+int32(4992))
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
							F_errmsg_internal(m, int32(345117), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(463387), int32(434), int32(102866))
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
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
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
		v209 = int32(1)
	} else {
		v31 = int32(1233)
		v36 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v12))^int32(63))*v31+v31) >> (uint(int32(12)) % 32))
		v41 = *(*int64)(unsafe.Add(mBase, uint32(v36<<(uint(int32(3))%32))+uint32(_consts[659])))
		v43 = v36 + base.B2i32(base.Ui64(v41) <= base.Ui64(v12))
		if base.Ui64(v12) < base.Ui64(int64(100000000)) {
			v121 = v15
			v125 = v12
		} else {
			v47 = v12
			v51 = v15
			for {
				v56 = v14 + v43 - v51
				v57 = int32(8)
				v60 = base.I64_div_u_s(v47, int64(100000000))
				v64 = base.I32_wrap_i64(v60*int64(4194967296) + v47)
				v66 = base.I32_div_u_s(v64, int32(1000000))
				v67 = int32(1)
				v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66<<(uint(v67)%32))+uint32(_consts[660]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v56-v57))) = uint16(v71)
				v75 = int32(10000)
				v76 = base.I32_div_u_s(v64, v75)
				v77 = int32(100)
				v78 = base.I32_rem_u_s(v76, v77)
				v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78<<(uint(v67)%32))+uint32(_consts[660]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v56-int32(6)))) = uint16(v83)
				v89 = v64 - v76*v75
				v90 = int32(65535)
				v93 = base.I32_div_u_s(v89&v90, v77)
				v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93<<(uint(v67)%32))+uint32(_consts[660]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v56-int32(4)))) = uint16(v98)
				v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v89-v93*v77)&v90<<(uint(v67)%32))+uint32(_consts[660]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v56-int32(2)))) = uint16(v111)
				v114 = v51 + v57
				if base.Ui64(int64(9999999999999999)) < base.Ui64(v47) {
					v47 = v60
					v51 = v114
					continue
				} else {
					break
				}
				break
			}
			v121 = v114
			v125 = v60
		}
		v126 = base.I32_wrap_i64(v125)
		if base.Ui64(v125) < base.Ui64(int64(10000)) {
			v160 = v126
			v161 = v121
		} else {
			v130 = v14 + v43 - v121
			v131 = int32(4)
			v134 = base.I32_div_u_s(v126, int32(10000))
			v137 = v134*int32(-10000) + v126
			v138 = int32(100)
			v139 = base.I32_div_u_s(v137, v138)
			v140 = int32(1)
			v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139<<(uint(v140)%32))+uint32(_consts[660]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v130-v131))) = uint16(v144)
			v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v137-v139*v138)<<(uint(v140)%32))+uint32(_consts[660]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v130-int32(2)))) = uint16(v155)
			v160 = v134
			v161 = v121 | v131
		}
		if base.Ui32(v160) < base.Ui32(int32(100)) {
			v183 = v160
			v184 = v161
		} else {
			v168 = int32(2)
			v170 = int32(100)
			v171 = base.I32_div_u_s(v160, v170)
			v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v160-v171*v170)<<(uint(int32(1))%32))+uint32(_consts[660]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v43-v161-v168))) = uint16(v179)
			v183 = v171
			v184 = v161 + v168
		}
		if base.Ui32(int32(10)) <= base.Ui32(v183) {
			v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183<<(uint(int32(1))%32))+uint32(_consts[660]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v43-v184-int32(2)))) = uint16(v195)
			v209 = v43
		} else {
			v198 = v183 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v198)
			v209 = v43
		}
	}
	v210 = v209 + v13
	v212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v210))) = uint8(v212)
	return v210
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_consts[865])))
	if v14 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_report_invalid_encoding_db(m, v39, v56, v40)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L24
	}
L2:
	;
	return v72
L3:
	;
	v39 = l0
	v40 = l1
	v42 = v4
	v44 = v4
	goto L16
L4:
	;
	if int32(0) < l1 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l1 < l2 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v22 = l1
	goto L10
L9:
	;
	v22 = l2
	goto L10
L10:
	;
	if v22 <= int32(0) {
		v72 = v4
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v28 = v4
	goto L12
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v28))))
	if v32 == int32(0) {
		v72 = v28
		goto L2
	} else {
		goto L14
	}
L13:
	;
	return v22
L14:
	;
	v36 = v28 + int32(1)
	if v36 != v22 {
		v28 = v36
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == int32(0) {
		v72 = v42
		goto L2
	} else {
		goto L18
	}
L17:
	;
	v72 = v65
	goto L2
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50*int32(28))+uint32(_consts[955])))
	v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v39)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v40 < v56 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v62 = v44 + int32(1)
	if l2 < v62 {
		v72 = v42
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v65 = v42 + v56
	v66 = v40 - v56
	if int32(0) < v66 {
		v39 = v39 + v56
		v40 = v66
		v42 = v65
		v44 = v62
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(69827)
			F_errmsg(m, int32(180554), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(462360), int32(1511), int32(33434))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_md5_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	if l0&int32(3) == int32(0) {
		v32 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v66 = v65 + l2
	v69 = F_emscripten_builtin_malloc(m, v66+int32(1))
	mBase = m.M
	if v69 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v65 = v57 - l0
	goto L1
L3:
	;
	v36 = v32
	goto L12
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v65 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v21 = l0
	goto L8
L8:
	;
	v25 = v21 + int32(1)
	if v25&int32(3) == int32(0) {
		v32 = v25
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v57 = v25
	goto L2
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v30 != 0 {
		v21 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v45 = int32(-2139062144)
	if (int32(16843008)-v42|v42)&v45 == v45 {
		v36 = v36 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v51 = v36
	goto L15
L14:
	;
	goto L13
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 != 0 {
		v51 = v51 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v57 = v51
	goto L2
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(12790)
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if l2 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v76 = F__emscripten_memcpy_bulkmem(m, v69, l0, v65)
	mBase = m.M
	v77 = v76
	goto L24
L23:
	;
	v77 = v69
	goto L24
L24:
	;
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3499117)
	v85 = F_pg_md5_hash(m, v77, v66, l3+int32(3), l4)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v79 = F__emscripten_memcpy_bulkmem(m, v77+v65, l1, l2)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	return int32(0)
L30:
	;
	F_emscripten_builtin_free(m, v77)
	mBase = m.M
	return v85
}
func F_pg_mule_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	if l1 <= int32(0) {
		v67 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v67 - l0
L2:
	;
	v9 = l1
	v11 = l0
	goto L3
L3:
	;
	v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	if int32(0) <= v13 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v67 = v60
	goto L1
L5:
	;
	v60 = v59 + v11
	v61 = v9 - v59
	if int32(0) < v61 {
		v9 = v61
		v11 = v60
		goto L3
	} else {
		goto L24
	}
L6:
	;
	if v13 != 0 {
		v59 = int32(1)
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
		v42 = int32(2)
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v67 = v11
	goto L1
L10:
	;
	if base.Ui32(v9) < base.Ui32(v42) {
		v67 = v11
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v24 = int32(3)
	if v13&int32(-2) == int32(-102) {
		v42 = v24
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.Ui32((v13+int32(112))&int32(255)) < base.Ui32(int32(10)) {
		v42 = v24
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v13&int32(254) == int32(156) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = int32(4)
	goto L16
L15:
	;
	v41 = int32(1)
	goto L16
L16:
	;
	v42 = v41
	goto L10
L17:
	;
	if base.Ui32(v42) < base.Ui32(int32(2)) {
		v59 = v42
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+1)))
	if int32(0) <= v46 {
		v67 = v11
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v42 == int32(2) {
		v59 = v42
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+2)))
	if int32(0) <= v51 {
		v67 = v11
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v42) < base.Ui32(int32(4)) {
		v59 = v42
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+3)))
	if int32(0) <= v56 {
		v67 = v11
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v59 = v42
	goto L5
L24:
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
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int64
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l6)+260))
	if v23 == v8 {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[787])))
		v29 = l6 + v26<<(uint(int32(4))%32)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[991])))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[992]))))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
		v245 = v8
	} else {
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+272)))
		if v39 == int32(1) {
			v42 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
			if v38 < v42 {
				v66 = v42
				v68 = l6 + int32(280)
				if v38 < v66 {
					v78 = v66 - v38
				} else {
					v76 = *(*int64)(unsafe.Add(mBase, uint32(v23<<(uint(int32(3))%32)+v68-int32(8))))
					v78 = v38 - v76
				}
				v80 = v78 - int64(1)
				v81 = int64(12622780800)
				v82 = base.I64_rem_s(v80, v81)
				v83 = v80 - v82
				v85 = v83 + v81
				if v38 < v66 {
					v89 = v85
				} else {
					v89 = int64(-12622780800) - v83
				}
				v90 = v89 + v38
				*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v90
				v92 = int32(-1)
				if v90 < v66 {
					v245 = v92
				} else {
					v99 = *(*int64)(unsafe.Add(mBase, uint32(v23<<(uint(int32(3))%32)+v68-int32(8))))
					if v99 < v90 {
						v245 = v92
					} else {
						v103 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
						mBase = m.M
						v104 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
						v107 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						if v38 < v107 {
							v109 = int64(-12622780800) - v83
						} else {
							v109 = v85
						}
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v104 + v109
						v245 = v103
					}
				}
			} else {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+273)))
				if v45 == int32(0) {
					v49 = v23 - int32(1)
					v53 = *(*int64)(unsafe.Add(mBase, uint32(l6+v49<<(uint(int32(3))%32))+280))
					v112 = v49
					v114 = v53
					if v114 <= v38 {
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+l6)+uint32(_consts[993]))))
						v122 = l6 + v119<<(uint(int32(4))%32)
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[991])))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v125
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[992]))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v129
						v245 = v8
					} else {
						v132 = l6 + int32(280)
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
						if v133 <= v38 {
							v135 = int32(1)
							if v135 < v112 {
								v139 = v112
								v146 = v135
								for {
									v158 = int32(1)
									v159 = (v139 + v146) >> (uint(v158) % 32)
									v165 = *(*int64)(unsafe.Add(mBase, uint32(v132+v159<<(uint(int32(3))%32))))
									v166 = base.B2i32(v38 < v165)
									if v38 < v165 {
										v167 = v146
									} else {
										v167 = v159 + v158
									}
									if v38 < v165 {
										v168 = v159
									} else {
										v168 = v139
									}
									if v167 < v168 {
										v139 = v168
										v146 = v167
										continue
									} else {
										break
									}
									break
								}
								v177 = v167
							} else {
								v177 = v135
							}
							v189 = l6 + int32(18280)
							v190 = l6 + v177
							v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[994]))))
							v194 = int32(4)
							v196 = v189 + v193<<(uint(v194)%32)
							v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v197
							v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v199
							v204 = *(*int64)(unsafe.Add(mBase, uint32(v132+v177<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v204
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[993]))))
							v211 = v189 + v208<<(uint(v194)%32)
							v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v212
							v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v214
							v245 = v135
						} else {
							v217 = l6 + int32(18280)
							v218 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[787])))
							v219 = int32(4)
							v221 = v217 + v218<<(uint(v219)%32)
							v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v222
							v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v224
							v226 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v226
							v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[993]))))
							v231 = v217 + v228<<(uint(v219)%32)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v232
							v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v234
							v245 = int32(1)
						}
					}
				} else {
					v55 = l6 + int32(280)
					v57 = v23 - int32(1)
					v61 = *(*int64)(unsafe.Add(mBase, uint32(v55+v57<<(uint(int32(3))%32))))
					if v38 <= v61 {
						v112 = v57
						v114 = v61
						if v114 <= v38 {
							v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+l6)+uint32(_consts[993]))))
							v122 = l6 + v119<<(uint(int32(4))%32)
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[991])))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v125
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[992]))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v129
							v245 = v8
						} else {
							v132 = l6 + int32(280)
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
							if v133 <= v38 {
								v135 = int32(1)
								if v135 < v112 {
									v139 = v112
									v146 = v135
									for {
										v158 = int32(1)
										v159 = (v139 + v146) >> (uint(v158) % 32)
										v165 = *(*int64)(unsafe.Add(mBase, uint32(v132+v159<<(uint(int32(3))%32))))
										v166 = base.B2i32(v38 < v165)
										if v38 < v165 {
											v167 = v146
										} else {
											v167 = v159 + v158
										}
										if v38 < v165 {
											v168 = v159
										} else {
											v168 = v139
										}
										if v167 < v168 {
											v139 = v168
											v146 = v167
											continue
										} else {
											break
										}
										break
									}
									v177 = v167
								} else {
									v177 = v135
								}
								v189 = l6 + int32(18280)
								v190 = l6 + v177
								v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[994]))))
								v194 = int32(4)
								v196 = v189 + v193<<(uint(v194)%32)
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v197
								v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v199
								v204 = *(*int64)(unsafe.Add(mBase, uint32(v132+v177<<(uint(int32(3))%32))))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v204
								v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[993]))))
								v211 = v189 + v208<<(uint(v194)%32)
								v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v212
								v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v214
								v245 = v135
							} else {
								v217 = l6 + int32(18280)
								v218 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[787])))
								v219 = int32(4)
								v221 = v217 + v218<<(uint(v219)%32)
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v222
								v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v224
								v226 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v226
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[993]))))
								v231 = v217 + v228<<(uint(v219)%32)
								v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v232
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v234
								v245 = int32(1)
							}
						}
					} else {
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
						v66 = v63
						v68 = l6 + int32(280)
						if v38 < v66 {
							v78 = v66 - v38
						} else {
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v23<<(uint(int32(3))%32)+v68-int32(8))))
							v78 = v38 - v76
						}
						v80 = v78 - int64(1)
						v81 = int64(12622780800)
						v82 = base.I64_rem_s(v80, v81)
						v83 = v80 - v82
						v85 = v83 + v81
						if v38 < v66 {
							v89 = v85
						} else {
							v89 = int64(-12622780800) - v83
						}
						v90 = v89 + v38
						*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v90
						v92 = int32(-1)
						if v90 < v66 {
							v245 = v92
						} else {
							v99 = *(*int64)(unsafe.Add(mBase, uint32(v23<<(uint(int32(3))%32)+v68-int32(8))))
							if v99 < v90 {
								v245 = v92
							} else {
								v103 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
								mBase = m.M
								v104 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
								v107 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
								if v38 < v107 {
									v109 = int64(-12622780800) - v83
								} else {
									v109 = v85
								}
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v104 + v109
								v245 = v103
							}
						}
					}
				}
			}
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+273)))
			if v45 == int32(0) {
				v49 = v23 - int32(1)
				v53 = *(*int64)(unsafe.Add(mBase, uint32(l6+v49<<(uint(int32(3))%32))+280))
				v112 = v49
				v114 = v53
				if v114 <= v38 {
					v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+l6)+uint32(_consts[993]))))
					v122 = l6 + v119<<(uint(int32(4))%32)
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[991])))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v125
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[992]))))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v129
					v245 = v8
				} else {
					v132 = l6 + int32(280)
					v133 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
					if v133 <= v38 {
						v135 = int32(1)
						if v135 < v112 {
							v139 = v112
							v146 = v135
							for {
								v158 = int32(1)
								v159 = (v139 + v146) >> (uint(v158) % 32)
								v165 = *(*int64)(unsafe.Add(mBase, uint32(v132+v159<<(uint(int32(3))%32))))
								v166 = base.B2i32(v38 < v165)
								if v38 < v165 {
									v167 = v146
								} else {
									v167 = v159 + v158
								}
								if v38 < v165 {
									v168 = v159
								} else {
									v168 = v139
								}
								if v167 < v168 {
									v139 = v168
									v146 = v167
									continue
								} else {
									break
								}
								break
							}
							v177 = v167
						} else {
							v177 = v135
						}
						v189 = l6 + int32(18280)
						v190 = l6 + v177
						v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[994]))))
						v194 = int32(4)
						v196 = v189 + v193<<(uint(v194)%32)
						v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v197
						v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v199
						v204 = *(*int64)(unsafe.Add(mBase, uint32(v132+v177<<(uint(int32(3))%32))))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v204
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[993]))))
						v211 = v189 + v208<<(uint(v194)%32)
						v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v212
						v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v214
						v245 = v135
					} else {
						v217 = l6 + int32(18280)
						v218 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[787])))
						v219 = int32(4)
						v221 = v217 + v218<<(uint(v219)%32)
						v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v222
						v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v224
						v226 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v226
						v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[993]))))
						v231 = v217 + v228<<(uint(v219)%32)
						v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v232
						v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v234
						v245 = int32(1)
					}
				}
			} else {
				v55 = l6 + int32(280)
				v57 = v23 - int32(1)
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v55+v57<<(uint(int32(3))%32))))
				if v38 <= v61 {
					v112 = v57
					v114 = v61
					if v114 <= v38 {
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+l6)+uint32(_consts[993]))))
						v122 = l6 + v119<<(uint(int32(4))%32)
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[991])))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v125
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_consts[992]))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v129
						v245 = v8
					} else {
						v132 = l6 + int32(280)
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v132)))
						if v133 <= v38 {
							v135 = int32(1)
							if v135 < v112 {
								v139 = v112
								v146 = v135
								for {
									v158 = int32(1)
									v159 = (v139 + v146) >> (uint(v158) % 32)
									v165 = *(*int64)(unsafe.Add(mBase, uint32(v132+v159<<(uint(int32(3))%32))))
									v166 = base.B2i32(v38 < v165)
									if v38 < v165 {
										v167 = v146
									} else {
										v167 = v159 + v158
									}
									if v38 < v165 {
										v168 = v159
									} else {
										v168 = v139
									}
									if v167 < v168 {
										v139 = v168
										v146 = v167
										continue
									} else {
										break
									}
									break
								}
								v177 = v167
							} else {
								v177 = v135
							}
							v189 = l6 + int32(18280)
							v190 = l6 + v177
							v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[994]))))
							v194 = int32(4)
							v196 = v189 + v193<<(uint(v194)%32)
							v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v197
							v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v199
							v204 = *(*int64)(unsafe.Add(mBase, uint32(v132+v177<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v204
							v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[993]))))
							v211 = v189 + v208<<(uint(v194)%32)
							v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v212
							v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v214
							v245 = v135
						} else {
							v217 = l6 + int32(18280)
							v218 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[787])))
							v219 = int32(4)
							v221 = v217 + v218<<(uint(v219)%32)
							v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v222
							v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v224
							v226 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v226
							v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_consts[993]))))
							v231 = v217 + v228<<(uint(v219)%32)
							v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v232
							v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v234
							v245 = int32(1)
						}
					}
				} else {
					v63 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
					v66 = v63
					v68 = l6 + int32(280)
					if v38 < v66 {
						v78 = v66 - v38
					} else {
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v23<<(uint(int32(3))%32)+v68-int32(8))))
						v78 = v38 - v76
					}
					v80 = v78 - int64(1)
					v81 = int64(12622780800)
					v82 = base.I64_rem_s(v80, v81)
					v83 = v80 - v82
					v85 = v83 + v81
					if v38 < v66 {
						v89 = v85
					} else {
						v89 = int64(-12622780800) - v83
					}
					v90 = v89 + v38
					*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v90
					v92 = int32(-1)
					if v90 < v66 {
						v245 = v92
					} else {
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v23<<(uint(int32(3))%32)+v68-int32(8))))
						if v99 < v90 {
							v245 = v92
						} else {
							v103 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
							mBase = m.M
							v104 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							v107 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							if v38 < v107 {
								v109 = int64(-12622780800) - v83
							} else {
								v109 = v85
							}
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v104 + v109
							v245 = v103
						}
					}
				}
			}
		}
	}
	m.G0 = v21 + int32(16)
	return v245
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
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
							v59 = int64(0)
							F_pfree(m, v20)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v22)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v66 = v59
									m.G0 = v11 + int32(16)
									return base.B2i32(v66 == int64(0))
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
											if v46 == int32(0) {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v59 = v49
													F_pfree(m, v20)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v22)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v66 = v59
															m.G0 = v11 + int32(16)
															return base.B2i32(v66 == int64(0))
														}
													}
												}
											} else {
												if v46 == v45 {
													F_ReleaseCatCache(m, v24)
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return int32(0)
													} else {
														v59 = v49
														F_pfree(m, v20)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v22)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v66 = v59
																m.G0 = v11 + int32(16)
																return base.B2i32(v66 == int64(0))
															}
														}
													}
												} else {
													F_pfree(m, v46)
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v24)
														mBase = m.M
														v57 = m.ExcPending
														if v57 != 0 {
															return int32(0)
														} else {
															v59 = v49
															F_pfree(m, v20)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v22)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return int32(0)
																} else {
																	v66 = v59
																	m.G0 = v11 + int32(16)
																	return base.B2i32(v66 == int64(0))
																}
															}
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
											if v46 == int32(0) {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													v59 = v49
													F_pfree(m, v20)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v22)
														mBase = m.M
														v64 = m.ExcPending
														if v64 != 0 {
															return int32(0)
														} else {
															v66 = v59
															m.G0 = v11 + int32(16)
															return base.B2i32(v66 == int64(0))
														}
													}
												}
											} else {
												if v46 == v45 {
													F_ReleaseCatCache(m, v24)
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return int32(0)
													} else {
														v59 = v49
														F_pfree(m, v20)
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v22)
															mBase = m.M
															v64 = m.ExcPending
															if v64 != 0 {
																return int32(0)
															} else {
																v66 = v59
																m.G0 = v11 + int32(16)
																return base.B2i32(v66 == int64(0))
															}
														}
													}
												} else {
													F_pfree(m, v46)
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v24)
														mBase = m.M
														v57 = m.ExcPending
														if v57 != 0 {
															return int32(0)
														} else {
															v59 = v49
															F_pfree(m, v20)
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v22)
																mBase = m.M
																v64 = m.ExcPending
																if v64 != 0 {
																	return int32(0)
																} else {
																	v66 = v59
																	m.G0 = v11 + int32(16)
																	return base.B2i32(v66 == int64(0))
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
		} else {
			v66 = l2
			m.G0 = v11 + int32(16)
			return base.B2i32(v66 == int64(0))
		}
	}
}
func F_pg_popcount_masked_optimized(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v163 int64
	_ = v163
	v3 = int32(0)
	v8 = int64(0)
	v9 = int32(8168)
	if (l0+int32(3))&int32(-4) == l0 {
		v16 = l1 * int32(16843009)
		v17 = l0
		v20 = v3
		v21 = v9
		v24 = v8
		for {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v44 = base.I64_extend_i32_u(base.I32_popcnt(v25&v16)) + (base.I64_extend_i32_u(base.I32_popcnt(v29&v16)) + (base.I64_extend_i32_u(base.I32_popcnt(v33&v16)) + (v24 + base.I64_extend_i32_u(base.I32_popcnt(v37&v16)))))
			v45 = int32(16)
			v46 = v21 - v45
			v48 = v17 + v45
			v50 = v20 + int32(4)
			if v50 != int32(2040) {
				v17 = v48
				v20 = v50
				v21 = v46
				v24 = v44
				continue
			} else {
				break
			}
			break
		}
		v56 = v48
		v57 = v46
		v58 = v3
		v60 = v44
		for {
			v61 = int32(4)
			v62 = v57 - v61
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v67 = v60 + base.I64_extend_i32_u(base.I32_popcnt(v63&v16))
			v69 = v56 + v61
			v71 = v58 + int32(1)
			if v71 != int32(2) {
				v56 = v69
				v57 = v62
				v58 = v71
				v60 = v67
				continue
			} else {
				break
			}
			break
		}
		v74 = v69
		v78 = v62
		v81 = v67
	} else {
		v74 = l0
		v78 = v9
		v81 = v8
	}
	if v78 == int32(0) {
		v163 = v81
	} else {
		v85 = v78 & int32(3)
		if v85 == int32(0) {
			v110 = v74
			v112 = v78
			v117 = v81
		} else {
			v91 = v78
			v92 = v74
			v94 = int32(0)
			v96 = v81
			for {
				v97 = int32(1)
				v98 = v91 - v97
				v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
				v103 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v99&l1)+uint32(_consts[1032]))))
				v104 = v96 + v103
				v106 = v92 + v97
				v108 = v94 + v97
				if v108 != v85 {
					v91 = v98
					v92 = v106
					v94 = v108
					v96 = v104
					continue
				} else {
					break
				}
				break
			}
			v110 = v106
			v112 = v98
			v117 = v104
		}
		if base.Ui32(v78) < base.Ui32(int32(4)) {
			v163 = v117
		} else {
			v120 = v110
			v122 = v112
			v127 = v117
			for {
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+3)))
				v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v128&l1)+uint32(_consts[1032]))))
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+2)))
				v137 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v133&l1)+uint32(_consts[1032]))))
				v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
				v142 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v138&l1)+uint32(_consts[1032]))))
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
				v147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v143&l1)+uint32(_consts[1032]))))
				v151 = v132 + (v137 + (v142 + (v127 + v147)))
				v152 = int32(4)
				v155 = v122 - v152
				if v155 != 0 {
					v120 = v120 + v152
					v122 = v155
					v127 = v151
					continue
				} else {
					break
				}
				break
			}
			v163 = v151
		}
	}
	return v163
}
func F_pg_popcount_optimized(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v82 int32
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int64
	_ = v134
	var v135 int32
	_ = v135
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int64
	_ = v165
	v7 = int64(0)
	if l1 < int32(4) {
		v86 = l0
		v87 = l1
		v92 = v7
	} else {
		if l0 != (l0+int32(3))&int32(-4) {
			v86 = l0
			v87 = l1
			v92 = v7
		} else {
			v16 = l1 - int32(4)
			v20 = int32(base.Ui32(v16)>>(uint(int32(2))%32)) + int32(1)
			v22 = v20 & int32(3)
			if base.Ui32(v16) < base.Ui32(int32(12)) {
				v58 = l0
				v59 = l1
				v64 = v7
			} else {
				v28 = l0
				v29 = l1
				v30 = int32(0)
				v34 = v7
				for {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v50 = base.I64_extend_i32_u(base.I32_popcnt(v35)) + (base.I64_extend_i32_u(base.I32_popcnt(v38)) + (base.I64_extend_i32_u(base.I32_popcnt(v41)) + (v34 + base.I64_extend_i32_u(base.I32_popcnt(v44)))))
					v51 = int32(16)
					v52 = v29 - v51
					v54 = v28 + v51
					v56 = v30 + int32(4)
					if v56 != v20&int32(2147483644) {
						v28 = v54
						v29 = v52
						v30 = v56
						v34 = v50
						continue
					} else {
						break
					}
					break
				}
				v58 = v54
				v59 = v52
				v64 = v50
			}
			if v22 == int32(0) {
				v86 = v58
				v87 = v59
				v92 = v64
			} else {
				v69 = v59
				v70 = v58
				v71 = int32(0)
				v74 = v64
				for {
					v75 = int32(4)
					v76 = v69 - v75
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
					v80 = v74 + base.I64_extend_i32_u(base.I32_popcnt(v77))
					v82 = v70 + v75
					v84 = v71 + int32(1)
					if v84 != v22 {
						v69 = v76
						v70 = v82
						v71 = v84
						v74 = v80
						continue
					} else {
						break
					}
					break
				}
				v86 = v82
				v87 = v76
				v92 = v80
			}
		}
	}
	if v87 == int32(0) {
		v165 = v92
	} else {
		v96 = v87 & int32(3)
		if v96 == int32(0) {
			v119 = v86
			v121 = v87
			v125 = v92
		} else {
			v102 = v87
			v103 = v86
			v104 = int32(0)
			v106 = v92
			for {
				v107 = int32(1)
				v108 = v102 - v107
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
				v112 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v109)+uint32(_consts[1032]))))
				v113 = v106 + v112
				v115 = v103 + v107
				v117 = v104 + v107
				if v117 != v96 {
					v102 = v108
					v103 = v115
					v104 = v117
					v106 = v113
					continue
				} else {
					break
				}
				break
			}
			v119 = v115
			v121 = v108
			v125 = v113
		}
		if base.Ui32(v87) < base.Ui32(int32(4)) {
			v165 = v125
		} else {
			v128 = v119
			v130 = v121
			v134 = v125
			for {
				v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+3)))
				v138 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135)+uint32(_consts[1032]))))
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)))
				v142 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_consts[1032]))))
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
				v146 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v143)+uint32(_consts[1032]))))
				v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
				v150 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_consts[1032]))))
				v154 = v138 + (v142 + (v146 + (v134 + v150)))
				v155 = int32(4)
				v158 = v130 - v155
				if v158 != 0 {
					v128 = v128 + v155
					v130 = v158
					v134 = v154
					continue
				} else {
					break
				}
				break
			}
			v165 = v154
		}
	}
	return v165
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[623]))
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
					F_errmsg(m, int32(275588), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(464474), int32(293), int32(317543))
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[624])))
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
				F_errmsg(m, int32(321527), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(464474), int32(313), int32(362587))
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v2 = int32(16)
	v3 = l0 + v2
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - v2
	m.G0 = v10
	v13 = l0 + int32(32)
	v18 = F___memset(m, l0+int32(40), v4, int32(144))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(1)
	v31 = F___memcpy(m, v10, v13, int32(16))
	mBase = m.M
	v32 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v3))) = v32
	v38 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+8)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+28)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = v38
	v46 = F___syscall_ret(m, v4)
	mBase = m.M
	m.G0 = v10 + int32(16)
	F___gettimeofday(m, l0)
	mBase = m.M
	return
}
func F_pg_saslprep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v244 int32
	_ = v244
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v797 int32
	_ = v797
	var v806 int32
	_ = v806
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v18 = l0
	goto L6
L1:
	;
	m.G0 = v13 + int32(16)
	return v806
L2:
	;
	F_pfree(m, v202)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L12
	} else {
		goto L227
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = int32(0)
	goto L2
L4:
	;
	v806 = int32(-1)
	goto L1
L5:
	;
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v20 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18))))
	if int32(0) < v20 {
		v18 = v18 + int32(1)
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
	v27 = F_pstrdup(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if l0&int32(3) == int32(0) {
		v57 = l0
		goto L17
	} else {
		goto L18
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v27
	if v27 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v806 = v3
	goto L1
L15:
	;
	if v90 != 0 {
		goto L32
	} else {
		goto L33
	}
L16:
	;
	v90 = v82 - l0
	goto L15
L17:
	;
	v61 = v57
	goto L26
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v90 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	v46 = l0
	goto L22
L22:
	;
	v50 = v46 + int32(1)
	if v50&int32(3) == int32(0) {
		v57 = v50
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v82 = v50
	goto L16
L24:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v55 != 0 {
		v46 = v50
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v70 = int32(-2139062144)
	if (int32(16843008)-v67|v67)&v70 == v70 {
		v61 = v61 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v76 = v61
	goto L29
L28:
	;
	goto L27
L29:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 != 0 {
		v76 = v76 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v82 = v76
	goto L16
L31:
	;
	goto L30
L32:
	;
	v93 = v3
	v94 = v90
	v96 = l0
	goto L35
L33:
	;
	v190 = v3
	goto L34
L34:
	;
	v202 = F_palloc(m, v190<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L76
	}
L35:
	;
	v101 = int32(-2)
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96))))
	if int32(0) <= v102 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	if v181 < int32(0) {
		v806 = v101
		goto L1
	} else {
		goto L74
	}
L37:
	;
	if base.Ui32(v94) < base.Ui32(v126) {
		v806 = v101
		goto L1
	} else {
		goto L50
	}
L38:
	;
	v126 = int32(1)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v107 = v102 & int32(255)
	if v107&int32(224) == int32(192) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v126 = int32(2)
	goto L37
L42:
	;
	goto L43
L43:
	;
	if v107&int32(240) == int32(224) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v126 = int32(3)
	goto L37
L45:
	;
	goto L46
L46:
	;
	if v107&int32(248) == int32(240) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v124 = int32(4)
	goto L49
L48:
	;
	v124 = int32(1)
	goto L49
L49:
	;
	v126 = v124
	goto L37
L50:
	;
	v128 = int32(0)
	switch v126 - int32(1) {
	case 0:
		goto L55
	case 1:
		goto L56
	case 2:
		goto L57
	case 3:
		goto L58
	default:
		v177 = v128
		goto L52
	}
L51:
	;
	if v177 == int32(0) {
		v806 = v101
		goto L1
	} else {
		goto L72
	}
L52:
	;
	goto L51
L53:
	;
	v177 = base.B2i32(base.Ui32(v169&int32(255)) < base.Ui32(int32(245)))
	goto L52
L54:
	;
	if base.I32_extend8_s(v164) < int32(-62) {
		v177 = v128
		goto L52
	} else {
		goto L71
	}
L55:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	v164 = v163
	goto L54
L56:
	;
	v137 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96)+1)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	switch v138 - int32(224) {
	case 0:
		goto L65
	default:
		goto L61
	case 13:
		goto L64
	case 16:
		goto L63
	case 20:
		goto L62
	}
L57:
	;
	v134 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96)+2)))
	if int32(-65) < v134 {
		v177 = v128
		goto L52
	} else {
		goto L60
	}
L58:
	;
	v131 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96)+3)))
	if int32(-65) < v131 {
		v177 = v128
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	goto L56
L61:
	;
	if v137 <= int32(-65) {
		v164 = v138
		goto L54
	} else {
		goto L70
	}
L62:
	;
	if int32(-113) < v137 {
		v177 = v128
		goto L52
	} else {
		goto L69
	}
L63:
	;
	if base.Ui32((v137-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v177 = v128
		goto L52
	} else {
		goto L68
	}
L64:
	;
	if int32(-97) < v137 {
		v177 = v128
		goto L52
	} else {
		goto L67
	}
L65:
	;
	v141 = int32(224)
	if base.Ui32(v141) <= base.Ui32((v137-int32(-64))&int32(255)) {
		v169 = v141
		goto L53
	} else {
		goto L66
	}
L66:
	;
	v177 = v128
	goto L52
L67:
	;
	v169 = int32(237)
	goto L53
L68:
	;
	v169 = int32(240)
	goto L53
L69:
	;
	v169 = int32(244)
	goto L53
L70:
	;
	v177 = v128
	goto L52
L71:
	;
	v169 = v164
	goto L53
L72:
	;
	v181 = v93 + int32(1)
	v183 = v94 - v126
	if v183 != 0 {
		v93 = v181
		v94 = v183
		v96 = v96 + v126
		goto L35
	} else {
		goto L73
	}
L73:
	;
	goto L36
L74:
	;
	if base.Ui32(int32(268435454)) < base.Ui32(v181) {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v190 = v181
	goto L34
L76:
	;
	if v202 == int32(0) {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	if v190 == int32(0) {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v209 = l0
	v215 = int32(0)
	goto L79
L79:
	;
	v219 = int32(*(*int8)(unsafe.Add(mBase, uint32(v209))))
	v221 = v219 & int32(255)
	if int32(0) <= v219 {
		v281 = v221
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v312 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v202+v190<<(uint(int32(2))%32)))) = v312
	v322 = v312
	v325 = v312
	goto L105
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202+v215<<(uint(int32(2))%32)))) = v281
	v283 = int32(*(*int8)(unsafe.Add(mBase, uint32(v209))))
	if int32(0) <= v283 {
		goto L92
	} else {
		goto L93
	}
L82:
	;
	if v221&int32(224) == int32(192) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v209))))
	v281 = v276&int32(63) | v273
	goto L81
L84:
	;
	v273 = v221 << (uint(int32(6)) % 32) & int32(1984)
	v274 = int32(1)
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v221&int32(240) == int32(224) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	v273 = v221<<(uint(int32(12))%32)&int32(61440) | v244&int32(63)<<(uint(int32(6))%32)
	v274 = int32(2)
	goto L83
L88:
	;
	goto L89
L89:
	;
	if v221&int32(248) != int32(240) {
		v281 = int32(-1)
		goto L81
	} else {
		goto L90
	}
L90:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	v261 = int32(63)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	v273 = v221<<(uint(int32(18))%32)&int32(1835008) | v260&v261<<(uint(int32(12))%32) | v266&v261<<(uint(int32(6))%32)
	v274 = int32(3)
	goto L83
L91:
	;
	v310 = v215 + int32(1)
	if v310 != v190 {
		v209 = v307 + v209
		v215 = v310
		goto L79
	} else {
		goto L104
	}
L92:
	;
	v307 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v288 = v283 & int32(255)
	if v288&int32(224) == int32(192) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v307 = int32(2)
	goto L91
L96:
	;
	goto L97
L97:
	;
	if v288&int32(240) == int32(224) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v307 = int32(3)
	goto L91
L99:
	;
	goto L100
L100:
	;
	if v288&int32(248) == int32(240) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v305 = int32(4)
	goto L103
L102:
	;
	v305 = int32(1)
	goto L103
L103:
	;
	v307 = v305
	goto L91
L104:
	;
	goto L80
L105:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v202+v325<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v332
	if base.Ui32(int32(-12130)) < base.Ui32(v332-int32(12289)) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v375 = v202 + v368<<(uint(int32(2))%32)
	v376 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = v376
	if v368 == v376 {
		goto L2
	} else {
		goto L120
	}
L107:
	;
	v371 = v325 + int32(1)
	if v371 != v190 {
		v322 = v368
		v325 = v371
		goto L105
	} else {
		goto L119
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202+v322<<(uint(int32(2))%32)))) = v361
	v368 = v322 + int32(1)
	goto L107
L109:
	;
	v345 = F_bsearch(m, v13+int32(12), int32(1814784), int32(6), int32(8), int32(4662))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L12
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v332
	if base.Ui32(v332-int32(65280)) <= base.Ui32(int32(-65108)) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v345 != 0 {
		v361 = int32(32)
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v361 = v332
	goto L108
L115:
	;
	goto L116
L116:
	;
	v356 = int32(8)
	v359 = F_bsearch(m, v13+int32(12), int32(1814832), v356, v356, int32(4662))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L117
	}
L117:
	;
	if v359 != 0 {
		v368 = v322
		goto L107
	} else {
		goto L118
	}
L118:
	;
	v361 = v332
	goto L108
L119:
	;
	goto L106
L120:
	;
	v381 = F_unicode_normalize(m, int32(2), v202)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
	;
	if v381 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_pfree(m, v202)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L12
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v368 <= int32(0) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L4
L126:
	;
	F_pfree(m, v202)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L12
	} else {
		goto L225
	}
L127:
	;
	F_pfree(m, v202)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L12
	} else {
		goto L223
	}
L128:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v516 != 0 {
		goto L161
	} else {
		goto L162
	}
L129:
	;
	v391 = int32(0)
	goto L130
L130:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v202+v391<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v404
	if base.Ui32(int32(1114112)) <= base.Ui32(v404) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v433 = int32(0)
	goto L140
L132:
	;
	v430 = v391 + int32(1)
	if v430 != v368 {
		v391 = v430
		goto L130
	} else {
		goto L139
	}
L133:
	;
	v414 = F_bsearch(m, v13+int32(12), int32(1814896), int32(36), int32(8), int32(4662))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L12
	} else {
		goto L134
	}
L134:
	;
	if v414 != 0 {
		goto L127
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v404
	if base.Ui32(v404-int32(983038)) <= base.Ui32(int32(-982494)) {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v427 = F_bsearch(m, v13+int32(12), int32(1815184), int32(396), int32(8), int32(4662))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L12
	} else {
		goto L137
	}
L137:
	;
	if v427 != 0 {
		goto L127
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	goto L131
L140:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v202+v433<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v446
	if base.Ui32(int32(-63808)) < base.Ui32(v446-int32(65277)) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v375-int32(4))))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v468 = int32(0)
	goto L149
L142:
	;
	goto L141
L143:
	;
	v458 = F_bsearch(m, v13+int32(12), int32(1818352), int32(34), int32(8), int32(4662))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L12
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v461 = v433 + int32(1)
	if v461 != v368 {
		v433 = v461
		goto L140
	} else {
		goto L148
	}
L146:
	;
	if v458 != 0 {
		goto L142
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	goto L128
L149:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v202+v468<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v481
	if base.Ui32(int32(-1114046)) < base.Ui32(v481-int32(1114110)) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v498 = F_is_code_in_table(m, v466)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L12
	} else {
		goto L157
	}
L151:
	;
	v493 = F_bsearch(m, v13+int32(12), int32(1818624), int32(360), int32(8), int32(4662))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L12
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v496 = v468 + int32(1)
	if v496 != v368 {
		v468 = v496
		goto L149
	} else {
		goto L156
	}
L154:
	;
	if v493 != 0 {
		goto L127
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	goto L150
L157:
	;
	if v498 == int32(0) {
		goto L127
	} else {
		goto L158
	}
L158:
	;
	v502 = F_is_code_in_table(m, v465)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L12
	} else {
		goto L159
	}
L159:
	;
	if v502 == int32(0) {
		goto L127
	} else {
		goto L160
	}
L160:
	;
	goto L128
L161:
	;
	v518 = v516
	v520 = int32(0)
	v524 = v381
	goto L164
L162:
	;
	v631 = int32(1)
	goto L163
L163:
	;
	v632 = F_palloc(m, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L12
	} else {
		goto L190
	}
L164:
	;
	if base.Ui32(v518) <= base.Ui32(int32(127)) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v631 = v615 + int32(1)
	goto L163
L166:
	;
	v590 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13+int32(12)))))
	if int32(0) <= v590 {
		goto L177
	} else {
		goto L178
	}
L167:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v518)
	goto L166
L168:
	;
	goto L169
L169:
	;
	if base.Ui32(v518) <= base.Ui32(int32(2047)) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v536 = v518&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v536)
	v541 = int32(base.Ui32(v518)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v541)
	goto L166
L171:
	;
	goto L172
L172:
	;
	if base.Ui32(v518) <= base.Ui32(int32(65535)) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v545 = int32(63)
	v547 = int32(128)
	v548 = v518&v545 | v547
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v548)
	v553 = int32(base.Ui32(v518)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v553)
	v560 = int32(base.Ui32(v518)>>(uint(int32(6))%32))&v545 | v547
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v560)
	goto L166
L174:
	;
	goto L175
L175:
	;
	v562 = int32(63)
	v564 = int32(128)
	v565 = v518&v562 | v564
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v565)
	v572 = int32(base.Ui32(v518)>>(uint(int32(6))%32))&v562 | v564
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v572)
	v579 = int32(base.Ui32(v518)>>(uint(int32(12))%32))&v562 | v564
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v579)
	v586 = int32(base.Ui32(v518)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v586)
	goto L166
L176:
	;
	v615 = v614 + v520
	v617 = v524 + int32(4)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	if v618 != 0 {
		v518 = v618
		v520 = v615
		v524 = v617
		goto L164
	} else {
		goto L189
	}
L177:
	;
	v614 = int32(1)
	goto L176
L178:
	;
	goto L179
L179:
	;
	v595 = v590 & int32(255)
	if v595&int32(224) == int32(192) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v614 = int32(2)
	goto L176
L181:
	;
	goto L182
L182:
	;
	if v595&int32(240) == int32(224) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v614 = int32(3)
	goto L176
L184:
	;
	goto L185
L185:
	;
	if v595&int32(248) == int32(240) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v612 = int32(4)
	goto L188
L187:
	;
	v612 = int32(1)
	goto L188
L188:
	;
	v614 = v612
	goto L176
L189:
	;
	goto L165
L190:
	;
	if v632 == int32(0) {
		goto L126
	} else {
		goto L191
	}
L191:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v636 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v637 = v632
	v639 = v636
	v643 = v381
	goto L195
L193:
	;
	v736 = v632
	goto L194
L194:
	;
	v746 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v736))) = uint8(v746)
	F_pfree(m, v202)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L12
	} else {
		goto L221
	}
L195:
	;
	if base.Ui32(v639) <= base.Ui32(int32(127)) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v736 = v732
	goto L194
L197:
	;
	v707 = int32(*(*int8)(unsafe.Add(mBase, uint32(v637))))
	if int32(0) <= v707 {
		goto L208
	} else {
		goto L209
	}
L198:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v637))) = uint8(v639)
	goto L197
L199:
	;
	goto L200
L200:
	;
	if base.Ui32(v639) <= base.Ui32(int32(2047)) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v655 = v639&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+1)) = uint8(v655)
	v660 = int32(base.Ui32(v639)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v637))) = uint8(v660)
	goto L197
L202:
	;
	goto L203
L203:
	;
	if base.Ui32(v639) <= base.Ui32(int32(65535)) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v664 = int32(63)
	v666 = int32(128)
	v667 = v639&v664 | v666
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+2)) = uint8(v667)
	v672 = int32(base.Ui32(v639)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v637))) = uint8(v672)
	v679 = int32(base.Ui32(v639)>>(uint(int32(6))%32))&v664 | v666
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+1)) = uint8(v679)
	goto L197
L205:
	;
	goto L206
L206:
	;
	v681 = int32(63)
	v683 = int32(128)
	v684 = v639&v681 | v683
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+3)) = uint8(v684)
	v691 = int32(base.Ui32(v639)>>(uint(int32(6))%32))&v681 | v683
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+2)) = uint8(v691)
	v698 = int32(base.Ui32(v639)>>(uint(int32(12))%32))&v681 | v683
	*(*uint8)(unsafe.Add(mBase, uint32(v637)+1)) = uint8(v698)
	v705 = int32(base.Ui32(v639)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v637))) = uint8(v705)
	goto L197
L207:
	;
	v732 = v731 + v637
	v734 = v643 + int32(4)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	if v735 != 0 {
		v637 = v732
		v639 = v735
		v643 = v734
		goto L195
	} else {
		goto L220
	}
L208:
	;
	v731 = int32(1)
	goto L207
L209:
	;
	goto L210
L210:
	;
	v712 = v707 & int32(255)
	if v712&int32(224) == int32(192) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v731 = int32(2)
	goto L207
L212:
	;
	goto L213
L213:
	;
	if v712&int32(240) == int32(224) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v731 = int32(3)
	goto L207
L215:
	;
	goto L216
L216:
	;
	if v712&int32(248) == int32(240) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v729 = int32(4)
	goto L219
L218:
	;
	v729 = int32(1)
	goto L219
L219:
	;
	v731 = v729
	goto L207
L220:
	;
	goto L196
L221:
	;
	F_pfree(m, v381)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L12
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v632
	v806 = v746
	goto L1
L223:
	;
	F_pfree(m, v381)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L12
	} else {
		goto L224
	}
L224:
	;
	v806 = int32(-3)
	goto L1
L225:
	;
	F_pfree(m, v381)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L12
	} else {
		goto L226
	}
L226:
	;
	goto L4
L227:
	;
	v806 = int32(-3)
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
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
		v20 = *(*int32)(unsafe.Add(mBase, _consts[4]))
		v22 = F_pg_class_aclcheck(m, v10, v20, int64(258))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			if v22 != 0 {
				F_sequence_close(m, v24, int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v68 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
					v73 = int32(0)
					m.G0 = v8 + int32(32)
					return v73
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+118)))
				switch v26 - int32(112) {
				case 0:
					v52 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+16)))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
						F_UnlockReleaseBuffer(m, v56)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_sequence_close(m, v24, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								if v55 != int32(1) {
									v68 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
									v73 = int32(0)
									m.G0 = v8 + int32(32)
									return v73
								} else {
									v64 = F_Int64GetDatum(m, v54)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v73 = v64
										m.G0 = v8 + int32(32)
										return v73
									}
								}
							}
						}
					}
				default:
					v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
					if v34 == int32(1) {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+316))
						v42 = base.B2i32(v40 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v42)
						v44 = v42
					} else {
						v44 = int32(0)
					}
					if v44 == int32(0) {
						v52 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+16)))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
							F_UnlockReleaseBuffer(m, v56)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_sequence_close(m, v24, int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									if v55 != int32(1) {
										v68 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
										v73 = int32(0)
										m.G0 = v8 + int32(32)
										return v73
									} else {
										v64 = F_Int64GetDatum(m, v54)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v73 = v64
											m.G0 = v8 + int32(32)
											return v73
										}
									}
								}
							}
						}
					} else {
						F_sequence_close(m, v24, int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v68 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
							v73 = int32(0)
							m.G0 = v8 + int32(32)
							return v73
						}
					}
				case 4:
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
					if v29 != int32(1) {
						F_sequence_close(m, v24, int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v68 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
							v73 = int32(0)
							m.G0 = v8 + int32(32)
							return v73
						}
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
						if v34 == int32(1) {
							v39 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+316))
							v42 = base.B2i32(v40 != int32(2))
							*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v42)
							v44 = v42
						} else {
							v44 = int32(0)
						}
						if v44 == int32(0) {
							v52 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+16)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
								F_UnlockReleaseBuffer(m, v56)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_sequence_close(m, v24, int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v55 != int32(1) {
											v68 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
											v73 = int32(0)
											m.G0 = v8 + int32(32)
											return v73
										} else {
											v64 = F_Int64GetDatum(m, v54)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v73 = v64
												m.G0 = v8 + int32(32)
												return v73
											}
										}
									}
								}
							}
						} else {
							F_sequence_close(m, v24, int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v68 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
								v73 = int32(0)
								m.G0 = v8 + int32(32)
								return v73
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v13 = v11 >> (uint(int64(63)) % 64)
	if base.Ui64(v11^v13-v13) <= base.Ui64(int64(10239)) {
		v64 = int32(148397)
		v65 = v11
		v68 = int32(1611376)
	} else {
		v21 = base.I64_div_s(v11, int64(512))
		v23 = v21 >> (uint(int64(63)) % 64)
		if base.Ui64(v21^v23-v23) < base.Ui64(int64(20479)) {
			v64 = int32(512689)
			v65 = v21
			v68 = int32(1611388)
		} else {
			v31 = base.I64_div_s(v11, int64(524288))
			v33 = v31 >> (uint(int64(63)) % 64)
			if base.Ui64(v31^v33-v33) < base.Ui64(int64(20479)) {
				v64 = int32(512717)
				v65 = v31
				v68 = int32(1611400)
			} else {
				v41 = base.I64_div_s(v11, int64(536870912))
				v43 = v41 >> (uint(int64(63)) % 64)
				if base.Ui64(v41^v43-v43) < base.Ui64(int64(20479)) {
					v64 = int32(512720)
					v65 = v41
					v68 = int32(1611412)
				} else {
					v51 = base.I64_div_s(v11, int64(549755813888))
					v53 = v51 >> (uint(int64(63)) % 64)
					if base.Ui64(v51^v53-v53) < base.Ui64(int64(20479)) {
						v64 = int32(512697)
						v65 = v51
						v68 = int32(1611424)
					} else {
						v61 = base.I64_div_s(v11, int64(562949953421312))
						v64 = int32(512700)
						v65 = v61
						v68 = int32(1611436)
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
	v86 = F_pg_snprintf(m, v8+int32(16), int32(64), int32(184767), v8)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		return int32(0)
	} else {
		v92 = F_cstring_to_text(m, v8+int32(16))
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(80)
			return v92
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
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_strtox_2(m, v16, v13+int32(20), int32(10), int64(-1))
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
	return v151
L3:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v15)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L31
	}
L4:
	;
	v32 = F_strtox_2(m, v22+int32(1), v13+int32(20), int32(10), int64(-1))
	mBase = m.M
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != int32(58) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if base.I32_wrap_i64(v21) == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v32&int64(4294967295) == int64(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if base.Ui64(v32) < base.Ui64(v21) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v49 = F_makeStringInfo(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v53 = int32(24)
	F_appendBinaryStringInfo(m, v49, v13+v53, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v59 = v33 + int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v114 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(0)
	F_pfree(m, v49)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L30
	}
L14:
	;
	v63 = v59
	v72 = int64(0)
	goto L15
L15:
	;
	v77 = F_strtox_2(m, v63, v13+int32(20), int32(10), int64(-1))
	mBase = m.M
	goto L17
L16:
	;
	goto L13
L17:
	;
	if base.Ui64(v77) < base.Ui64(v72) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if base.Ui64(v77) < base.Ui64(v21) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if base.Ui64(v32) <= base.Ui64(v77) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v77 != v72 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v77
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v85 + int32(1)
	F_appendBinaryStringInfo(m, v49, v13+int32(24), int32(8))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v95 != int32(44) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	if v95 == int32(0) {
		goto L13
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v101 = v81 + int32(1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v102 != 0 {
		v63 = v101
		v72 = v77
		goto L15
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	goto L16
L30:
	;
	v151 = v113
	goto L2
L31:
	;
	if v133 == int32(0) {
		v151 = v132
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(81202)
	F_errmsg(m, int32(674710), v13)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errsave_finish(m, v15, int32(464611), int32(324), int32(81227))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	v151 = v132
	goto L2
}
func F_pg_snapshot_xip(m *base.Module, l0 int32) int32 {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v21 = F_MemoryContextAlloc(m, v17, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v25 = int32(base.Ui32(v23) >> (uint(int32(2)) % 32))
					if v25 != 0 {
						v26 = F__emscripten_memcpy_bulkmem(m, v21, v11, v25)
						mBase = m.M
						v27 = v26
					} else {
						v27 = v21
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v27
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
					v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+4)))
					if base.Ui64(v33) < base.Ui64(v35) {
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v34+base.I32_wrap_i64(v33)<<(uint(int32(3))%32))+24))
						*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33 + int64(1)
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(1)
						v48 = F_Int64GetDatum(m, v41)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							return v48
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(2)
							v56 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
		v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+4)))
		if base.Ui64(v33) < base.Ui64(v35) {
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v34+base.I32_wrap_i64(v33)<<(uint(int32(3))%32))+24))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33 + int64(1)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(1)
			v48 = F_Int64GetDatum(m, v41)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				return v48
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(2)
				v56 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
				return int32(0)
			}
		}
	}
}
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
	v13 = v9 + int32(28)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+20)) = int64(0)
	if l1 != 0 {
		v20 = l0
	} else {
		v20 = v9 + int32(7)
	}
	v21 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v21) {
		v24 = v21
	} else {
		v24 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v20 + v24 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v20
	F_dopr(m, v9+int32(8), l2, l3)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return int32(0)
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v38 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v38)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		m.G0 = v9 + int32(32)
		if v42 != 0 {
			v49 = int32(-1)
		} else {
			v49 = v41 + (v37 - v40)
		}
		return v49
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
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(115739)
					F_errmsg(m, int32(633504), v6)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(467508), int32(730), int32(115739))
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
			v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[151])))
			if v36 != int32(1) {
				v40 = *(*int32)(unsafe.Add(mBase, _consts[29]))
				v44 = F_LWLockAcquire(m, v40+int32(256), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _consts[67]))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					if base.Ui32(int32(12001)) <= base.Ui32(v48) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _consts[67]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(12000)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v85
							F_errmsg_internal(m, int32(37396), v33)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(465288), int32(634), int32(437135))
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
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(12000)
						v54 = *(*int32)(unsafe.Add(mBase, _consts[67]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(0)
						v58 = *(*int32)(unsafe.Add(mBase, _consts[29]))
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
					F_errmsg_internal(m, int32(342342), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(465288), int32(627), int32(437135))
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(l0) <= base.Ui32(int32(153)) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v214
L2:
	;
	v214 = l1
	goto L1
L3:
	;
	switch l0 - int32(1) {
	case 0:
		v214 = int32(504710)
		goto L1
	case 1:
		goto L104
	case 2:
		goto L103
	case 3:
		goto L102
	case 4:
		goto L101
	case 5:
		goto L100
	case 6:
		goto L99
	case 7:
		goto L98
	case 8:
		goto L97
	case 9:
		goto L96
	default:
		goto L44
	case 11:
		goto L95
	case 12:
		goto L94
	case 13:
		goto L93
	case 14:
		goto L92
	case 15:
		goto L91
	case 17:
		goto L90
	case 19:
		goto L89
	case 20:
		goto L88
	case 21:
		goto L87
	case 22:
		goto L85
	case 23:
		goto L84
	case 25:
		goto L83
	case 26:
		goto L82
	case 27:
		goto L81
	case 28:
		goto L80
	case 29:
		goto L79
	case 30:
		goto L78
	case 31:
		goto L77
	case 32:
		goto L76
	case 33:
		goto L75
	case 34:
		goto L74
	case 36:
		goto L73
	case 37:
		goto L72
	case 38:
		goto L71
	case 39:
		goto L70
	case 40:
		goto L69
	case 41:
		goto L68
	case 42:
		goto L67
	case 43:
		goto L66
	case 44:
		goto L65
	case 47:
		goto L64
	case 50:
		goto L63
	case 51:
		goto L62
	case 52:
		goto L61
	case 53:
		goto L60
	case 54:
		goto L59
	case 56:
		goto L58
	case 58:
		goto L56
	case 59:
		goto L55
	case 60:
		goto L54
	case 62:
		goto L53
	case 63:
		goto L52
	case 65:
		goto L51
	case 67:
		goto L50
	case 68:
		goto L49
	case 70:
		goto L48
	case 72:
		goto L47
	case 73:
		goto L46
	case 74:
		goto L45
	case 137:
		goto L57
	case 141:
		goto L86
	}
L4:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L5:
	;
	v12 = l0
	goto L7
L6:
	;
	v12 = int32(0)
	goto L7
L7:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12<<(uint(int32(1))%32))+uint32(_consts[1034]))))
	v19 = v17 + int32(4026836)
	goto L4
L8:
	;
	goto L12
L9:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L10:
	;
	v134 = F_strlen(m, v123)
	mBase = m.M
	goto L9
L12:
	;
	goto L13
L13:
	;
	v28 = int32(255)
	if (l1^v19)&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v127)
	goto L10
L15:
	;
	v108 = v103
	v109 = v104
	v110 = v105
	goto L37
L16:
	;
	if v98 == int32(0) {
		v123 = v96
		v124 = v97
		goto L14
	} else {
		goto L36
	}
L17:
	;
	v96 = v19
	v97 = l1
	v98 = v28
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v19&int32(3) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v65 == int32(0) {
		v123 = v62
		v124 = v63
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v62 = v19
	v63 = l1
	v64 = v28
	v65 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v41 = v19
	v42 = l1
	v43 = v28
	goto L24
L24:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v45)
	if v45 == int32(0) {
		v103 = v41
		v104 = v42
		v105 = v43
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v62 = v56
	v63 = v50
	v64 = v52
	v65 = v54
	goto L20
L26:
	;
	v49 = int32(1)
	v50 = v42 + v49
	v52 = v43 - v49
	v53 = int32(0)
	v54 = base.B2i32(v52 != v53)
	v56 = v41 + v49
	if v56&int32(3) == v53 {
		v62 = v56
		v63 = v50
		v64 = v52
		v65 = v54
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v52 != 0 {
		v41 = v56
		v42 = v50
		v43 = v52
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v68 == int32(0) {
		v96 = v62
		v97 = v63
		v98 = v64
		goto L16
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v64) < base.Ui32(int32(4)) {
		v96 = v62
		v97 = v63
		v98 = v64
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v74 = v62
	v75 = v63
	v76 = v64
	goto L32
L32:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v82 = int32(-2139062144)
	if (int32(16843008)-v79|v79)&v82 != v82 {
		v103 = v74
		v104 = v75
		v105 = v76
		goto L15
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	v97 = v88
	v98 = v92
	goto L16
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v79
	v87 = int32(4)
	v88 = v75 + v87
	v90 = v74 + v87
	v92 = v76 - v87
	if base.Ui32(int32(3)) < base.Ui32(v92) {
		v74 = v90
		v75 = v88
		v76 = v92
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v103 = v96
	v104 = v97
	v105 = v98
	goto L15
L37:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v112)
	if v112 == int32(0) {
		v123 = v108
		v124 = v109
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v123 = v119
	v124 = v117
	goto L14
L39:
	;
	v116 = int32(1)
	v117 = v109 + v116
	v119 = v108 + v116
	v121 = v110 - v116
	if v121 != 0 {
		v108 = v119
		v109 = v117
		v110 = v121
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v139 == int32(63) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	if v139 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L3
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v209 = F_pg_snprintf(m, l1, int32(256), int32(442142), v7)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L105
	} else {
		goto L106
	}
L45:
	;
	v214 = int32(484793)
	goto L1
L46:
	;
	v214 = int32(476899)
	goto L1
L47:
	;
	v214 = int32(485302)
	goto L1
L48:
	;
	v214 = int32(504032)
	goto L1
L49:
	;
	v214 = int32(491552)
	goto L1
L50:
	;
	v214 = int32(509346)
	goto L1
L51:
	;
	v214 = int32(485572)
	goto L1
L52:
	;
	v214 = int32(507695)
	goto L1
L53:
	;
	v214 = int32(499419)
	goto L1
L54:
	;
	v214 = int32(484469)
	goto L1
L55:
	;
	v214 = int32(495291)
	goto L1
L56:
	;
	v214 = int32(476649)
	goto L1
L57:
	;
	v214 = int32(494322)
	goto L1
L58:
	;
	v214 = int32(502138)
	goto L1
L59:
	;
	v214 = int32(476656)
	goto L1
L60:
	;
	v214 = int32(493308)
	goto L1
L61:
	;
	v214 = int32(498373)
	goto L1
L62:
	;
	v214 = int32(490560)
	goto L1
L63:
	;
	v214 = int32(512152)
	goto L1
L64:
	;
	v214 = int32(499876)
	goto L1
L65:
	;
	v214 = int32(512383)
	goto L1
L66:
	;
	v214 = int32(486549)
	goto L1
L67:
	;
	v214 = int32(484799)
	goto L1
L68:
	;
	v214 = int32(491544)
	goto L1
L69:
	;
	v214 = int32(508329)
	goto L1
L70:
	;
	v214 = int32(504051)
	goto L1
L71:
	;
	v214 = int32(489336)
	goto L1
L72:
	;
	v214 = int32(496119)
	goto L1
L73:
	;
	v214 = int32(504183)
	goto L1
L74:
	;
	v214 = int32(505348)
	goto L1
L75:
	;
	v214 = int32(502111)
	goto L1
L76:
	;
	v214 = int32(508336)
	goto L1
L77:
	;
	v214 = int32(494944)
	goto L1
L78:
	;
	v214 = int32(493316)
	goto L1
L79:
	;
	v214 = int32(498382)
	goto L1
L80:
	;
	v214 = int32(495297)
	goto L1
L81:
	;
	v214 = int32(501954)
	goto L1
L82:
	;
	v214 = int32(492550)
	goto L1
L83:
	;
	v214 = int32(490945)
	goto L1
L84:
	;
	v214 = int32(499425)
	goto L1
L85:
	;
	v214 = int32(504038)
	goto L1
L86:
	;
	v214 = int32(496109)
	goto L1
L87:
	;
	v214 = int32(504704)
	goto L1
L88:
	;
	v214 = int32(488861)
	goto L1
L89:
	;
	v214 = int32(485376)
	goto L1
L90:
	;
	v214 = int32(499772)
	goto L1
L91:
	;
	v214 = int32(502126)
	goto L1
L92:
	;
	v214 = int32(489346)
	goto L1
L93:
	;
	v214 = int32(511610)
	goto L1
L94:
	;
	v214 = int32(511533)
	goto L1
L95:
	;
	v214 = int32(510876)
	goto L1
L96:
	;
	v214 = int32(476893)
	goto L1
L97:
	;
	v214 = int32(504069)
	goto L1
L98:
	;
	v214 = int32(505166)
	goto L1
L99:
	;
	v214 = int32(477925)
	goto L1
L100:
	;
	v214 = int32(498773)
	goto L1
L101:
	;
	v214 = int32(485588)
	goto L1
L102:
	;
	v214 = int32(501723)
	goto L1
L103:
	;
	v214 = int32(506949)
	goto L1
L104:
	;
	v214 = int32(491863)
	goto L1
L105:
	;
	return int32(0)
L106:
	;
	goto L2
}
func F_pg_strlower(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v11 - int32(98) {
	case 0:
		v34 = int32(0)
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
		v38 = F_convert_case(m, l0, l1, l2, l3, v34, v35, v34, v34)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = v38
			m.G0 = v9 + int32(16)
			return v40
		}
	case 1:
		v14 = F_strlower_libc(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v40 = v14
			m.G0 = v9 + int32(16)
			return v40
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(201260)
			F_errmsg_internal(m, int32(471342), v9)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(468462), int32(1284), int32(201260))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
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
func F_pg_sync_replication_slots(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_CheckSlotPermissions(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
		if v13 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
			v21 = base.B2i32(v19 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v21)
			v23 = v21
		} else {
			v23 = int32(0)
		}
		if v23 != 0 {
			v25 = F_ValidateSlotSyncParams(m, int32(21))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				F_load_file(m, int32(202329), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = m.G0
					v33 = v31 - int32(16)
					m.G0 = v33
					v36 = *(*int32)(unsafe.Add(mBase, _consts[548]))
					v38 = *(*int32)(unsafe.Add(mBase, _consts[293]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
					v40 = m.T0[v39].(func(*base.Module, int32) int32)(m, v36)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v40 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = int32(226208)
									*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(354749)
									F_errmsg(m, int32(661777), v33)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(469255), int32(1052), int32(226225))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
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
							m.G0 = v33 + int32(16)
							F_initStringInfo(m, v5+int32(28))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, _consts[549]))
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
								if v72 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v71
									F_appendStringInfo(m, v5+int32(28), int32(459549), v5+int32(16))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, _consts[548]))
										v88 = int32(0)
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
										v95 = *(*int32)(unsafe.Add(mBase, _consts[293]))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
										v97 = m.T0[v96].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v87, v88, v88, v88, v91, v5+int32(44))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											if v97 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return int32(0)
													} else {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
														*(*int32)(unsafe.Add(mBase, uint32(v5))) = v139
														v141 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v141
														F_errmsg(m, int32(188766), v5)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(464260), int32(929), int32(109261))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
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
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
												F_pfree(m, v101)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													F_SyncReplicationSlots(m, v97)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, _consts[293]))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+64))
														m.T0[v108].(func(*base.Module, int32))(m, v97)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return int32(0)
														} else {
															m.G0 = v5 + int32(48)
															return int32(0)
														}
													}
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, v5+int32(28), int32(459552))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, _consts[548]))
										v88 = int32(0)
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
										v95 = *(*int32)(unsafe.Add(mBase, _consts[293]))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
										v97 = m.T0[v96].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v87, v88, v88, v88, v91, v5+int32(44))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											if v97 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return int32(0)
													} else {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
														*(*int32)(unsafe.Add(mBase, uint32(v5))) = v139
														v141 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v141
														F_errmsg(m, int32(188766), v5)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(464260), int32(929), int32(109261))
															mBase = m.M
															v150 = m.ExcPending
															if v150 != 0 {
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
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
												F_pfree(m, v101)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													F_SyncReplicationSlots(m, v97)
													mBase = m.M
													v105 = m.ExcPending
													if v105 != 0 {
														return int32(0)
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, _consts[293]))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+64))
														m.T0[v108].(func(*base.Module, int32))(m, v97)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return int32(0)
														} else {
															m.G0 = v5 + int32(48)
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
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(201624), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(464260), int32(906), int32(109261))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v3 = m.G0
	v5 = v3 - int32(2208)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if v7 != 0 {
		v10 = v7
	} else {
		v10 = v9
	}
	if base.Ui32(v10-int32(1663)) <= base.Ui32(int32(1)) {
		v16 = F_cstring_to_text(m, int32(706478))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v69 = v16
			m.G0 = v5 + int32(2208)
			return v69
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+52)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = int32(459271)
		v29 = F_pg_snprintf(m, v5+int32(1184), int32(1024), int32(36192), v5+int32(48))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v37 = F___fstatat(m, int32(-100), v5+int32(1184), v5-int32(-64), int32(256))
			mBase = m.M
			if v37 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(1184)
						F_errmsg(m, int32(279590), v5)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(469103), int32(341), int32(249512))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
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
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)+68))
				if v40&int32(61440) != int32(40960) {
					v47 = F_cstring_to_text(m, v5+int32(1184))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v69 = v47
						m.G0 = v5 + int32(2208)
						return v69
					}
				} else {
					v54 = F_readlink(m, v5+int32(1184), v5+int32(160), int32(1024))
					mBase = m.M
					if v54 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v5 + int32(1184)
								F_errmsg(m, int32(279356), v5+int32(16))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(469103), int32(355), int32(249512))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
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
						if base.Ui32(int32(1024)) <= base.Ui32(v54) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v5 + int32(1184)
									F_errmsg(m, int32(307273), v5+int32(32))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(469103), int32(360), int32(249512))
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
						} else {
							v60 = v5 + int32(160)
							v62 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v60+v54))) = uint8(v62)
							v66 = F_cstring_to_text(m, v60)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v69 = v66
								m.G0 = v5 + int32(2208)
								return v69
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
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v314 int64
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int64
	_ = v394
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
	var v415 int64
	_ = v415
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
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
	v24 = F_palloc0(m, int32(23784))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[788])))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_get_share_path(m, int32(4425696))
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
	v215 = F_pstrdup(m, int32(4425696))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L57
	}
L7:
	;
	v33 = int32(4425696)
	goto L10
L8:
	;
	v92 = v85 + int32(4425696)
	v93 = int32(349812)
	v95 = int32(1024) - v85
	if v95 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	v85 = v76 - v33
	goto L8
L10:
	;
	v61 = v33
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v70 = int32(-2139062144)
	if (int32(16843008)-v67|v67)&v70 == v70 {
		v61 = v61 + int32(4)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v76 = v61
	goto L22
L21:
	;
	goto L20
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 != 0 {
		v76 = v76 + int32(1)
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L9
L24:
	;
	goto L23
L25:
	;
	v211 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[788])) = uint8(v211)
	goto L6
L26:
	;
	v207 = F_strlen(m, v203)
	mBase = m.M
	goto L25
L27:
	;
	v203 = v93
	goto L26
L28:
	;
	goto L29
L29:
	;
	v101 = v95 - int32(1)
	if (v92^v93)&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v200)
	v203 = v196
	goto L26
L31:
	;
	v181 = v176
	v182 = v177
	v183 = v178
	goto L53
L32:
	;
	if v171 == int32(0) {
		v196 = v169
		v197 = v170
		goto L30
	} else {
		goto L52
	}
L33:
	;
	v169 = v93
	v170 = v92
	v171 = v101
	goto L32
L34:
	;
	goto L35
L35:
	;
	goto L37
L36:
	;
	if base.B2i32(v101 != int32(0)) == int32(0) {
		v196 = v93
		v197 = v92
		goto L30
	} else {
		goto L45
	}
L37:
	;
	goto L36
L45:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _consts[789])))
	if v141 == int32(0) {
		v169 = v93
		v170 = v92
		v171 = v101
		goto L32
	} else {
		goto L46
	}
L46:
	;
	if base.Ui32(v101) < base.Ui32(int32(4)) {
		v169 = v93
		v170 = v92
		v171 = v101
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v147 = v93
	v148 = v92
	v149 = v101
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v155 = int32(-2139062144)
	if (int32(16843008)-v152|v152)&v155 != v155 {
		v176 = v147
		v177 = v148
		v178 = v149
		goto L31
	} else {
		goto L50
	}
L49:
	;
	v169 = v163
	v170 = v161
	v171 = v165
	goto L32
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v152
	v160 = int32(4)
	v161 = v148 + v160
	v163 = v147 + v160
	v165 = v149 - v160
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v147 = v163
		v148 = v161
		v149 = v165
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L31
L53:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v185)
	if v185 == int32(0) {
		v196 = v181
		v197 = v182
		goto L30
	} else {
		goto L55
	}
L54:
	;
	v196 = v192
	v197 = v190
	goto L30
L55:
	;
	v189 = int32(1)
	v190 = v182 + v189
	v192 = v181 + v189
	v194 = v183 - v189
	if v194 != 0 {
		v181 = v192
		v182 = v190
		v183 = v194
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	if v215&int32(3) == int32(0) {
		v240 = v215
		goto L60
	} else {
		goto L61
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v273 + int32(1)
	v280 = F_AllocateDir(m, v215)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L75
	}
L59:
	;
	v273 = v265 - v215
	goto L58
L60:
	;
	v244 = v240
	goto L69
L61:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v224 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v273 = int32(0)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v229 = v215
	goto L65
L65:
	;
	v233 = v229 + int32(1)
	if v233&int32(3) == int32(0) {
		v240 = v233
		goto L60
	} else {
		goto L67
	}
L66:
	;
	v265 = v233
	goto L59
L67:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v238 != 0 {
		v229 = v233
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v253 = int32(-2139062144)
	if (int32(16843008)-v250|v250)&v253 == v253 {
		v244 = v244 + int32(4)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v259 = v244
	goto L72
L71:
	;
	goto L70
L72:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v263 != 0 {
		v259 = v259 + int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v265 = v259
	goto L59
L74:
	;
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v280
	if v280 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	m.G0 = v21 + int32(16)
	v303 = F_pg_tzenumerate_next(m, v24)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L83
	}
L79:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v215
	F_errmsg(m, int32(278356), v21)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(462134), int32(409), int32(77432))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
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
	if v303 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v306 = v9 + int32(16)
	v307 = v303
	goto L87
L85:
	;
	goto L86
L86:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) <= v451 {
		goto L126
	} else {
		goto L127
	}
L87:
	;
	v314 = *(*int64)(unsafe.Add(mBase, _consts[96]))
	v323 = F_timestamp2tm(m, v314, v9+int32(88), v9+int32(44), v9+int32(40), v9+int32(36), v307)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v443 = F_pg_tzenumerate_next(m, v24)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L124
	}
L90:
	;
	if v323 != 0 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v325 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v325&int32(3) == int32(0) {
		v349 = v325
		goto L97
	} else {
		goto L98
	}
L93:
	;
	goto L94
L94:
	;
	v385 = F_cstring_to_text(m, v307)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L113
	}
L95:
	;
	if base.Ui32(int32(31)) < base.Ui32(v382) {
		goto L89
	} else {
		goto L112
	}
L96:
	;
	v382 = v374 - v325
	goto L95
L97:
	;
	v353 = v349
	goto L106
L98:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	if v333 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v382 = int32(0)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v338 = v325
	goto L102
L102:
	;
	v342 = v338 + int32(1)
	if v342&int32(3) == int32(0) {
		v349 = v342
		goto L97
	} else {
		goto L104
	}
L103:
	;
	v374 = v342
	goto L96
L104:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v347 != 0 {
		v338 = v342
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v362 = int32(-2139062144)
	if (int32(16843008)-v359|v359)&v362 == v362 {
		v353 = v353 + int32(4)
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v368 = v353
	goto L109
L108:
	;
	goto L107
L109:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v372 != 0 {
		v368 = v368 + int32(1)
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v374 = v368
	goto L96
L111:
	;
	goto L110
L112:
	;
	goto L94
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v385
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v388 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v390 = v388
	goto L116
L115:
	;
	v390 = int32(706478)
	goto L116
L116:
	;
	v391 = F_cstring_to_text(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+100)) = v391
	v394 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v306)+8)) = v394
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v394
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = base.I64_extend_i32_s(int32(0)-v399) * int64(1000000)
	v406 = v9 + int32(8)
	v408 = F_palloc(m, int32(16))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v411 = int64(*(*int32)(unsafe.Add(mBase, uint32(v406)+12)))
	v412 = int64(*(*int32)(unsafe.Add(mBase, uint32(v406)+16)))
	v415 = v411 + v412*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v415-int64(2147483648)) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = v408
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+108)) = base.B2i32(int32(0) < v429)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v433, v434, v9+int32(96), v9+int32(92))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L123
	}
L120:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v408)+12)) = uint32(v415)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v408)+8)) = v421
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v406)))
	*(*int64)(unsafe.Add(mBase, uint32(v408))) = v423
	goto L122
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	goto L89
L124:
	;
	if v443 != 0 {
		v307 = v443
		goto L87
	} else {
		goto L125
	}
L125:
	;
	goto L88
L126:
	;
	v460 = v451
	goto L129
L127:
	;
	goto L128
L128:
	;
	F_pfree(m, v24)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L134
	}
L129:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v460<<(uint(int32(2))%32))))
	F_FreeDir(m, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(48)+v470<<(uint(int32(2))%32))))
	F_pfree(m, v474)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v479 = v477 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v479
	if int32(0) <= v479 {
		v460 = v479
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	v6 = m.G0
	v8 = v6 - int32(256)
	m.G0 = v8
	v11 = l0 >> (uint(int32(31)) % 32)
	v13 = l0 ^ v11 - v11
	v15 = base.I32_div_s(v13, int32(3600))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v15
	v23 = F_pg_snprintf(m, v8+int32(192), int32(64), int32(405873), v8+int32(48))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = v13 - v15*int32(3600)
	if v29 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v188 = v8 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v188
	if int32(0) < l0 {
		goto L42
	} else {
		goto L43
	}
L4:
	;
	v33 = v8 + int32(192)
	if v33&int32(3) == int32(0) {
		v57 = v33
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v93 = base.I32_div_s(base.I32_extend16_s(v29), int32(60))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = base.I32_extend16_s(v93)
	v104 = F_pg_snprintf(m, v90+(v8+int32(192)), int32(64)-v90, int32(405872), v8+int32(32))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v90 = v82 - v33
	goto L5
L7:
	;
	v61 = v57
	goto L16
L8:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v90 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v46 = v33
	goto L12
L12:
	;
	v50 = v46 + int32(1)
	if v50&int32(3) == int32(0) {
		v57 = v50
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v82 = v50
	goto L6
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v55 != 0 {
		v46 = v50
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v70 = int32(-2139062144)
	if (int32(16843008)-v67|v67)&v70 == v70 {
		v61 = v61 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v76 = v61
	goto L19
L18:
	;
	goto L17
L19:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 != 0 {
		v76 = v76 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v82 = v76
	goto L6
L21:
	;
	goto L20
L22:
	;
	v108 = v29 - v93*int32(60)
	if v108&int32(65535) == int32(0) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v114 = v8 + int32(192)
	if v114&int32(3) == int32(0) {
		v138 = v114
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = base.I32_extend16_s(v108)
	v182 = F_pg_snprintf(m, v171+(v8+int32(192)), int32(64)-v171, int32(405872), v8+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L41
	}
L25:
	;
	v171 = v163 - v114
	goto L24
L26:
	;
	v142 = v138
	goto L35
L27:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v122 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v171 = int32(0)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v127 = v114
	goto L31
L31:
	;
	v131 = v127 + int32(1)
	if v131&int32(3) == int32(0) {
		v138 = v131
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v163 = v131
	goto L25
L33:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v136 != 0 {
		v127 = v131
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v151 = int32(-2139062144)
	if (int32(16843008)-v148|v148)&v151 == v151 {
		v142 = v142 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v157 = v142
	goto L38
L37:
	;
	goto L36
L38:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v161 != 0 {
		v157 = v157 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v163 = v157
	goto L25
L40:
	;
	goto L39
L41:
	;
	goto L3
L42:
	;
	v200 = int32(165425)
	goto L44
L43:
	;
	v200 = int32(165410)
	goto L44
L44:
	;
	v201 = F_pg_snprintf(m, v8-int32(-64), int32(128), v200, v8)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v205 = F_pg_tzset(m, v8-int32(-64))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v8 + int32(256)
	return v205
}
func F_pg_ultoa_n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	if l0 == int32(0) {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(1)
	} else {
		v19 = int32(1233)
		v24 = int32(base.Ui32((base.I32_clz(l0)^int32(31))*v19+v19) >> (uint(int32(12)) % 32))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_consts[808])))
		v31 = v24 + base.B2i32(base.Ui32(v29) <= base.Ui32(l0))
		v32 = int32(0)
		if base.Ui32(l0) < base.Ui32(int32(10000)) {
			v78 = v32
			v79 = l0
		} else {
			v36 = l0
			v38 = v32
			for {
				v45 = l1 + v31 - v38
				v46 = int32(4)
				v49 = base.I32_div_u_s(v36, int32(10000))
				v52 = v49*int32(-10000) + v36
				v53 = int32(100)
				v54 = base.I32_div_u_s(v52, v53)
				v55 = int32(1)
				v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54<<(uint(v55)%32))+uint32(_consts[660]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v45-v46))) = uint16(v59)
				v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v52-v54*v53)<<(uint(v55)%32))+uint32(_consts[660]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v45-int32(2)))) = uint16(v70)
				v73 = v38 + v46
				if base.Ui32(int32(99999999)) < base.Ui32(v36) {
					v36 = v49
					v38 = v73
					continue
				} else {
					break
				}
				break
			}
			v78 = v73
			v79 = v49
		}
		if base.Ui32(v79) < base.Ui32(int32(100)) {
			v108 = v79
			v109 = v78
		} else {
			v89 = int32(2)
			v91 = int32(65535)
			v93 = int32(100)
			v94 = base.I32_div_u_s(v79&v91, v93)
			v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v79-v94*v93)&v91<<(uint(int32(1))%32))+uint32(_consts[660]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v31-v78-v89))) = uint16(v104)
			v108 = v94
			v109 = v78 | v89
		}
		if base.Ui32(int32(10)) <= base.Ui32(v108) {
			v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108<<(uint(int32(1))%32))+uint32(_consts[660]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v31-v109-int32(2)))) = uint16(v120)
			return v31
		} else {
			v124 = v108 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v124)
			return v31
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
	var v70 float64
	_ = v70
	var v75 float64
	_ = v75
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
	var v96 int32
	_ = v96
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
	v16 = int32(1000000)
	v17 = base.I32_div_u_s(l0, v16)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = base.I64_extend_i32_u(v17)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = (l0 - v17*v16) * int32(1000)
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = int32(28)
	if v12 == int32(0) {
		v96 = v30
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
	v107 = int32(0) - v96
	if base.Ui32(int32(-4095)) <= base.Ui32(v107) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if base.Ui32(int32(999999999)) < base.Ui32(v33) {
		v96 = v30
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if v36 < int64(0) {
		v96 = v30
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
	v96 = int32(0)
	goto L4
L10:
	;
	v67 = v57 << (uint(int32(3)) % 32)
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1033])))
	if base.F64_eq(v70, float64(0)) != 0 {
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
	v75 = m.Env.Emscripten_get_now(m)
	mBase = m.M
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1033])))
	v77 = v75
	v78 = v76
	goto L16
L15:
	;
	v77 = v61
	v78 = v70
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
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0) - v107
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v98 int32
	_ = v98
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v254 int32
	_ = v254
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	if base.Ui32(l1) < base.Ui32(int32(16)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v352 <= int32(0) {
		v487 = v354
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v352 = l1
	v354 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v31 = l0
	v33 = int32(11)
	v34 = l1
	goto L5
L5:
	;
	if v33 != int32(11) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	switch v254 {
	case 0:
		v352 = l1
		v354 = l0
		goto L1
	default:
		goto L21
	case 11:
		v328 = v275
		v331 = v277
		goto L20
	}
L7:
	;
	v274 = int32(16)
	v275 = v31 + v274
	v277 = v34 - v274
	if base.Ui32(int32(15)) < base.Ui32(v277) {
		v31 = v275
		v33 = v254
		v34 = v277
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+15)))
	v133 = int32(2)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(v133)%32))+uint32(_consts[1029])))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+14)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138<<(uint(v133)%32))+uint32(_consts[1029])))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+13)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144<<(uint(v133)%32))+uint32(_consts[1029])))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+12)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150<<(uint(v133)%32))+uint32(_consts[1029])))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+11)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156<<(uint(v133)%32))+uint32(_consts[1029])))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+10)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162<<(uint(v133)%32))+uint32(_consts[1029])))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+9)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168<<(uint(v133)%32))+uint32(_consts[1029])))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+8)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174<<(uint(v133)%32))+uint32(_consts[1029])))
	v180 = int32(255)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v113&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v114&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v115&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v116&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v117&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v118&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v119&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v120&v180<<(uint(v133)%32))+uint32(_consts[1029])))
	v254 = int32(base.Ui32(v137)>>(uint(int32(base.Ui32(v143)>>(uint(int32(base.Ui32(v149)>>(uint(int32(base.Ui32(v155)>>(uint(int32(base.Ui32(v161)>>(uint(int32(base.Ui32(v167)>>(uint(int32(base.Ui32(v173)>>(uint(int32(base.Ui32(v179)>>(uint(int32(base.Ui32(v186)>>(uint(int32(base.Ui32(v193)>>(uint(int32(base.Ui32(v200)>>(uint(int32(base.Ui32(v207)>>(uint(int32(base.Ui32(v214)>>(uint(int32(base.Ui32(v221)>>(uint(int32(base.Ui32(v228)>>(uint(int32(base.Ui32(v235)>>(uint(v33)%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)) & int32(31)
	goto L7
L9:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+7)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v113 = v56
	v114 = v57
	v115 = v58
	v116 = v59
	v117 = v60
	v118 = v61
	v119 = v62
	v120 = v63
	goto L8
L10:
	;
	goto L11
L11:
	;
	v64 = int32(11)
	v66 = v31 + int32(16)
	if base.Ui32(v66) <= base.Ui32(v31) {
		v254 = v64
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	v72 = (v68 + int64(9187201950435737471)) & int64(-9187201950435737472)
	v88 = v31 + int32(8)
	if base.Ui32(v66) <= base.Ui32(v88) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v98 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(56)) % 64)))
	v99 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(48)) % 64)))
	v100 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(40)) % 64)))
	v101 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(32)) % 64)))
	v102 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(24)) % 64)))
	v103 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(16)) % 64)))
	v104 = base.I32_wrap_i64(int64(base.Ui64(v68) >> (uint(int64(8)) % 64)))
	v105 = base.I32_wrap_i64(v68)
	if v96&int64(-9187201950435737472) != int64(0) {
		v113 = v98
		v114 = v99
		v115 = v100
		v116 = v101
		v117 = v102
		v118 = v103
		v119 = v104
		v120 = v105
		goto L8
	} else {
		goto L17
	}
L14:
	;
	v95 = v72
	v96 = v68
	goto L13
L15:
	;
	goto L16
L16:
	;
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	v95 = (v90 + int64(9187201950435737471)) & v72
	v96 = v90 | v68
	goto L13
L17:
	;
	if v95 == int64(-9187201950435737472) {
		v254 = v64
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v113 = v98
	v114 = v99
	v115 = v100
	v116 = v101
	v117 = v102
	v118 = v103
	v119 = v104
	v120 = v105
	goto L8
L19:
	;
	goto L6
L20:
	;
	v352 = v331
	v354 = v328
	goto L1
L21:
	;
	v282 = v275
	v285 = v277
	goto L22
L22:
	;
	v305 = int32(1)
	v306 = v285 + v305
	v308 = v282 - v305
	v309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v308))))
	v311 = v309 & int32(255)
	if int32(0) <= v309 {
		v282 = v308
		v285 = v306
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v352 = v306
	v354 = v308
	goto L1
L24:
	;
	if v311&int32(248) == int32(240) {
		v328 = v308
		v331 = v306
		goto L20
	} else {
		goto L25
	}
L25:
	;
	if v311&int32(224) == int32(192) {
		v328 = v308
		v331 = v306
		goto L20
	} else {
		goto L26
	}
L26:
	;
	if v311&int32(240) != int32(224) {
		v282 = v308
		v285 = v306
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	return v487 - l0
L29:
	;
	v379 = v352
	v381 = v354
	goto L30
L30:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v404 = base.I32_extend8_s(v403)
	if int32(0) <= v404 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v487 = v480
	goto L28
L32:
	;
	v480 = v479 + v381
	v481 = v379 - v479
	if int32(0) < v481 {
		v379 = v481
		v381 = v480
		goto L30
	} else {
		goto L66
	}
L33:
	;
	if v404 != 0 {
		v479 = int32(1)
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v403&int32(224) == int32(192) {
		v425 = int32(2)
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v487 = v381
	goto L28
L37:
	;
	if base.Ui32(v379) < base.Ui32(v425) {
		v487 = v381
		goto L28
	} else {
		goto L43
	}
L38:
	;
	if v403&int32(240) == int32(224) {
		v425 = int32(3)
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if v403&int32(248) == int32(240) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v424 = int32(4)
	goto L42
L41:
	;
	v424 = int32(1)
	goto L42
L42:
	;
	v425 = v424
	goto L37
L43:
	;
	v427 = int32(0)
	switch v425 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L49
	case 2:
		goto L50
	case 3:
		goto L51
	default:
		v476 = v427
		goto L45
	}
L44:
	;
	if v476 == int32(0) {
		v487 = v381
		goto L28
	} else {
		goto L65
	}
L45:
	;
	goto L44
L46:
	;
	v476 = base.B2i32(base.Ui32(v468&int32(255)) < base.Ui32(int32(245)))
	goto L45
L47:
	;
	if base.I32_extend8_s(v463) < int32(-62) {
		v476 = v427
		goto L45
	} else {
		goto L64
	}
L48:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v463 = v462
	goto L47
L49:
	;
	v436 = int32(*(*int8)(unsafe.Add(mBase, uint32(v381)+1)))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	switch v437 - int32(224) {
	case 0:
		goto L58
	default:
		goto L54
	case 13:
		goto L57
	case 16:
		goto L56
	case 20:
		goto L55
	}
L50:
	;
	v433 = int32(*(*int8)(unsafe.Add(mBase, uint32(v381)+2)))
	if int32(-65) < v433 {
		v476 = v427
		goto L45
	} else {
		goto L53
	}
L51:
	;
	v430 = int32(*(*int8)(unsafe.Add(mBase, uint32(v381)+3)))
	if int32(-65) < v430 {
		v476 = v427
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L49
L54:
	;
	if v436 <= int32(-65) {
		v463 = v437
		goto L47
	} else {
		goto L63
	}
L55:
	;
	if int32(-113) < v436 {
		v476 = v427
		goto L45
	} else {
		goto L62
	}
L56:
	;
	if base.Ui32((v436-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v476 = v427
		goto L45
	} else {
		goto L61
	}
L57:
	;
	if int32(-97) < v436 {
		v476 = v427
		goto L45
	} else {
		goto L60
	}
L58:
	;
	v440 = int32(224)
	if base.Ui32(v440) <= base.Ui32((v436-int32(-64))&int32(255)) {
		v468 = v440
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v476 = v427
	goto L45
L60:
	;
	v468 = int32(237)
	goto L46
L61:
	;
	v468 = int32(240)
	goto L46
L62:
	;
	v468 = int32(244)
	goto L46
L63:
	;
	v476 = v427
	goto L45
L64:
	;
	v468 = v463
	goto L46
L65:
	;
	v479 = v425
	goto L32
L66:
	;
	goto L31
}
