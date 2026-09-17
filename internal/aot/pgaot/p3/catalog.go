package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheCreateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v536 int32
	_ = v536
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v690 int32
	_ = v690
	var v718 int32
	_ = v718
	var v728 int32
	_ = v728
	var v755 int32
	_ = v755
	var v756 int64
	_ = v756
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v793 int32
	_ = v793
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(304)
	m.G0 = v29
	v32 = l0 + int32(48)
	v42 = int32(-1)
	v44 = v6
	v46 = v6
	v47 = v6
	v48 = v6
	v49 = v6
	v51 = v6
	v52 = v6
	v53 = v6
	v55 = v6
	goto L2
L1:
	;
	m.G0 = v29 + int32(304)
	return v793
L2:
	;
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v567
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v570
	v793 = v446
	goto L1
L4:
	;
	if v42 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v755 = int32(m.ExcTag)
	v756 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v755 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v446)+64)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v446))) = int32(1462113538)
	v468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v446)+60)) = v468
	*(*uint8)(unsafe.Add(mBase, uint32(v446)+36)) = uint8(v468)
	*(*int32)(unsafe.Add(mBase, uint32(v446)+32)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v446)+4)) = l3
	v476 = v456 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v446)+37)) = uint8(v476)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v479 = v478 + l4<<(uint(int32(3))%32)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v480 == v468 {
		goto L63
	} else {
		goto L64
	}
L8:
	;
	v297 = int32(_a_F_CatalogCacheCreateEntry_0)
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0]))
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0])) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	v306 = v296 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v306)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	v316 = F_palloc(m, v303+int32(76))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L6
	} else {
		goto L41
	}
L9:
	;
	v162 = int32(_a_F_CatalogCacheCreateEntry_0)
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0]))
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0])) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v67)
	v177 = F_palloc(m, int32(68))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v67 = base.B2i32(l1 == int32(0))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v99 = v44
	v100 = v46
	v101 = v47
	v102 = v48
	v103 = v49
	v104 = v55
	goto L12
L12:
	;
	if v99 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+20)))
	if v71&int32(4) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v290 = l1
	v292 = v46
	v293 = v47
	v294 = v48
	v295 = v49
	v296 = v67
	goto L8
L15:
	;
	goto L16
L16:
	;
	v76 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+200)) = uint16(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+196)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = l0
	v81 = int32(_a_F_CatalogCacheCreateEntry_1)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2])) = v29 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+204)) = v82
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3]))
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[4]))
	goto L17
L17:
	;
	v93 = v29 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v29 + int32(24)
	goto L20
L18:
	;
	v99 = v76
	v100 = v29 + int32(201)
	v101 = v82
	v102 = v89
	v103 = v91
	v104 = v67
	goto L12
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3])) = v29 + int32(32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v101
	v120 = v104 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v120)
	v122 = F_toast_flatten_tuple(m, l1, v111)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[4])) = v103
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3])) = v102
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2])) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v101
	v158 = v104 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v158)
	F_pg_re_throw(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3])) = v102
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2])) = v101
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[4])) = v103
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v130 == int32(0) {
		v290 = v122
		v292 = v100
		v293 = v101
		v294 = v102
		v295 = v103
		v296 = v104
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v101
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v120)
	F_pfree(m, v122)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v793 = int32(0)
	goto L1
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v179 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v190 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0])) = v163
	v446 = v177
	v447 = v46
	v448 = v47
	v449 = v48
	v450 = v49
	v452 = v51
	v453 = v52
	v454 = v53
	v456 = v67
	goto L7
L32:
	;
	v213 = v190 << (uint(int32(2)) % 32)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l2+v213)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v213+v32)))
	v226 = v184 + v216<<(uint(int32(4))%32) + v221*int32(100) - int32(80)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+68))
	if v227 != int32(19) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L31
L34:
	;
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v226)+72)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+82)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v67)
	v256 = F_datumCopy(m, v244, v246, v245)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L39
	}
L35:
	;
	v244 = v215
	goto L34
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v47
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v67)
	v239 = v29 + int32(208)
	v241 = F_strncpy(m, v239, v215, int32(64))
	mBase = m.M
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+63)) = uint8(v242)
	goto L38
L38:
	;
	v244 = v239
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213+(v177+int32(8))))) = v256
	v260 = v190 + int32(1)
	if v260 != v179 {
		v190 = v260
		goto L32
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+40)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+44)) = v320
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v290)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v316)+48)) = uint16(v322)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v328 = (v316 + int32(75)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v316)+56)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v316)+52)) = v324
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v331 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v290)+16))
	base.MemoryCopy(m, v328, v332, v331)
	goto L44
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0])) = v298
	if l1 != v290 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v293
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v306)
	F_pfree(m, v290)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v347 <= int32(0) {
		v446 = v316
		v447 = v292
		v448 = v293
		v449 = v294
		v450 = v295
		v452 = v51
		v453 = v52
		v454 = v53
		v456 = v296
		goto L7
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v351 = v316 + int32(40)
	v359 = int32(0)
	v368 = v51
	v369 = v52
	v370 = v53
	goto L50
L50:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v383 = v359 << (uint(int32(2)) % 32)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v32+v383)))
	if int32(0) < v385 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v446 = v316
	v447 = v292
	v448 = v293
	v449 = v294
	v450 = v295
	v452 = v429
	v453 = v430
	v454 = v431
	v456 = v296
	goto L7
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316+int32(8)+v383))) = v432
	v436 = v359 + int32(1)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v436 < v437 {
		v359 = v436
		v368 = v429
		v369 = v430
		v370 = v431
		goto L50
	} else {
		goto L62
	}
L53:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v316)+56))
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v388)+18)))
	if base.Ui32(v389&int32(2047)) < base.Ui32(v385) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v293
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v306)
	v427 = F_heap_getsysattr(m, v351, v385, v29+int32(31))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L61
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v293
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v306)
	v403 = F_getmissingattr(m, v381, v385, v29+int32(31))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v293
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v306)
	v415 = F_fastgetattr_4(m, v351, v385, v381, v29+int32(31))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	v429 = v368
	v430 = v369
	v431 = v403
	v432 = v403
	goto L52
L60:
	;
	v429 = v415
	v430 = v369
	v431 = v370
	v432 = v415
	goto L52
L61:
	;
	v429 = v368
	v430 = v427
	v431 = v370
	v432 = v427
	goto L52
L62:
	;
	goto L51
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v479
	v484 = v479
	goto L65
L64:
	;
	v484 = v480
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v446)+24)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v446)+28)) = v484
	v488 = v446 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v479)+4)) = v488
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v492 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v491 + v492
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[5]))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v496)+4)) = v497 + v492
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v501 <= v502<<(uint(v492)%32) {
		v793 = v446
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v448
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v476)
	v516 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	if v516 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v448
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v476)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v518
	F_errmsg_internal(m, int32(_a_F_CatalogCacheCreateEntry_2), v29)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v448
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v476)
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[1]))
	v567 = F_MemoryContextAllocZero(m, v564, v554<<(uint(int32(4))%32))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v448
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v476)
	F_errfinish(m, int32(_a_F_CatalogCacheCreateEntry_3), int32(992), int32(_a_F_CatalogCacheCreateEntry_4))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v570 = v554 << (uint(int32(1)) % 32)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v571 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v583 = v571
	v595 = int32(0)
	goto L77
L75:
	;
	goto L76
L76:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+276)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v29)+280)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v29)+284)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v29)+288)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v29)+292)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v29)+296)) = v448
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)) = uint8(v476)
	F_pfree(m, v718)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L6
	} else {
		goto L89
	}
L77:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v606 = v603 + v595<<(uint(int32(3))%32)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+4))
	v608 = int32(0)
	if base.B2i32(v607 == v608)|base.B2i32(v607 == v606) == v608 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L76
L79:
	;
	v618 = v607
	goto L82
L80:
	;
	v669 = v583
	goto L81
L81:
	;
	v690 = v595 + int32(1)
	if v690 < v669 {
		v583 = v669
		v595 = v690
		goto L77
	} else {
		goto L88
	}
L82:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v618-int32(20))))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v643)+4)) = v644
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v646
	v651 = v567 + v642&(v570-int32(1))<<(uint(int32(3))%32)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	if v652 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v669 = v662
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651))) = v651
	v656 = v651
	goto L86
L85:
	;
	v656 = v652
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v618))) = v651
	*(*int32)(unsafe.Add(mBase, uint32(v618)+4)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v656))) = v618
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v618
	if v606 != v644 {
		v618 = v644
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	goto L78
L89:
	;
	goto L5
L90:
	;
	v760 = int32(v756)
	m.G0 = v29
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v763)))
	if v29+int32(24) == v766 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	m.ExcPending = 1
	goto L99
L92:
	;
	if v770 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	v770 = v768
	goto L95
L94:
	;
	v770 = int32(0)
	goto L95
L95:
	;
	goto L92
L96:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+303)))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v29)+296))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v29)+292))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v29)+288))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v29)+284))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v29)+280))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v29)+276))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
	v42 = v770
	v44 = v762
	v46 = v773
	v47 = v772
	v48 = v774
	v49 = v775
	v51 = v777
	v52 = v778
	v53 = v776
	v55 = v771
	goto L2
L97:
	;
	goto L98
L98:
	;
	F___wasm_longjmp(m, v763, v762)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	return int32(0)
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	v4 = F_palloc0(m, int32(216))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+52)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(388)
		F_ExecOpenIndices(m, v4, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v4
		}
	}
}
func F_GetCatalogSnapshot(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetCatalogSnapshot[0]))
	if v4 != 0 {
		v9 = v4
		return v9
	} else {
		v5 = F_GetNonHistoricCatalogSnapshot(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = v5
			return v9
		}
	}
}
