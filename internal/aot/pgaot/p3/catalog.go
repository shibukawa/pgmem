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
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
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
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v414 int64
	_ = v414
	var v415 int32
	_ = v415
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v438 int64
	_ = v438
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v443 int64
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v472 int64
	_ = v472
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v546 int32
	_ = v546
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v705 int32
	_ = v705
	var v735 int32
	_ = v735
	var v745 int32
	_ = v745
	var v774 int32
	_ = v774
	var v775 int64
	_ = v775
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int64
	_ = v795
	var v796 int64
	_ = v796
	var v797 int64
	_ = v797
	var v801 int32
	_ = v801
	var v812 int32
	_ = v812
	v6 = int32(0)
	v24 = int64(0)
	v29 = m.G0
	v31 = v29 - int32(320)
	m.G0 = v31
	v34 = l0 + int32(48)
	v44 = int32(-1)
	v46 = v6
	v48 = v6
	v49 = v6
	v50 = v6
	v51 = v6
	v52 = v6
	v63 = v24
	v64 = v24
	v65 = v24
	goto L2
L1:
	;
	m.G0 = v31 + int32(320)
	return v812
L2:
	;
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v580
	v812 = v456
	goto L1
L4:
	;
	if v44 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v774 = int32(m.ExcTag)
	v775 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v774 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456)+80)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v456)+8)) = int32(1462113538)
	v480 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v456)+76)) = v480
	*(*uint8)(unsafe.Add(mBase, uint32(v456)+52)) = uint8(v480)
	*(*int32)(unsafe.Add(mBase, uint32(v456)+48)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v456)+12)) = l3
	v488 = v458 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v456)+53)) = uint8(v488)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v491 = v490 + l4<<(uint(int32(3))%32)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	if v492 == v480 {
		goto L62
	} else {
		goto L63
	}
L8:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v308
	v320 = v307 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v320)
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0]))
	v326 = F_MemoryContextAlloc(m, v323, v311+int32(88))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L40
	}
L9:
	;
	v166 = int32(_a_F_CatalogCacheCreateEntry_0)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[1]))
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[1])) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v71)
	v181 = F_palloc(m, int32(88))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v71 = base.B2i32(l1 == int32(0))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v103 = v46
	v104 = v48
	v105 = v49
	v106 = v50
	v107 = v51
	v108 = v52
	goto L12
L12:
	;
	if v103 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+20)))
	if v75&int32(4) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v304 = l1
	v306 = v48
	v307 = v71
	v308 = v50
	v309 = v51
	v310 = v52
	goto L8
L15:
	;
	goto L16
L16:
	;
	v80 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+200)) = uint16(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+196)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = l0
	v85 = int32(_a_F_CatalogCacheCreateEntry_1)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2])) = v31 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+204)) = v86
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3]))
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[4]))
	goto L17
L17:
	;
	v97 = v31 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v31 + int32(24)
	goto L20
L18:
	;
	v103 = v80
	v104 = v31 + int32(201)
	v105 = v71
	v106 = v86
	v107 = v93
	v108 = v95
	goto L12
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3])) = v31 + int32(32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v106
	v124 = v105 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v124)
	v126 = F_toast_flatten_tuple(m, l1, v115)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[4])) = v108
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3])) = v107
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2])) = v106
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v106
	v162 = v105 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v162)
	F_pg_re_throw(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[3])) = v107
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[2])) = v106
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[4])) = v108
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v134 == int32(0) {
		v304 = v126
		v306 = v104
		v307 = v105
		v308 = v106
		v309 = v107
		v310 = v108
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v106
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v124)
	F_pfree(m, v126)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v812 = int32(0)
	goto L1
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v183 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v197 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[1])) = v167
	v456 = v181
	v457 = v48
	v458 = v71
	v459 = v50
	v460 = v51
	v461 = v52
	v472 = v63
	v473 = v64
	v474 = v65
	goto L7
L32:
	;
	v221 = int32(3)
	v222 = v197 << (uint(v221) % 32)
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l2+v222)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v34+v197<<(uint(int32(2))%32))))
	v237 = v188 + v225<<(uint(v221)%32) + v232*int32(100) - int32(72)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+68))
	if v238 == int32(19) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v71)
	v253 = F_strncpy(m, v31+int32(208), base.I32_wrap_i64(v224), int32(64))
	mBase = m.M
	v254 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+63)) = uint8(v254)
	goto L37
L35:
	;
	v256 = v224
	goto L36
L36:
	;
	v257 = int32(*(*int16)(unsafe.Add(mBase, uint32(v237)+72)))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+82)))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v71)
	v268 = F_datumCopy(m, v256, v258, v257)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L38
	}
L37:
	;
	v256 = base.I64_extend_i32_u(v31 + int32(208))
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v181+int32(16)+v222))) = v268
	v272 = v197 + int32(1)
	if v272 != v183 {
		v197 = v272
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L33
L40:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	*(*int32)(unsafe.Add(mBase, uint32(v326)+56)) = v328
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v326)+60)) = v330
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v326)+64)) = uint16(v332)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v336 = v326 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v326)+72)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v326)+68)) = v334
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	if v339 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v304)+16))
	base.MemoryCopy(m, v336, v340, v339)
	goto L43
L42:
	;
	goto L43
L43:
	;
	if l1 != v304 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v308
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v320)
	F_pfree(m, v304)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v353 <= int32(0) {
		v456 = v326
		v457 = v306
		v458 = v307
		v459 = v308
		v460 = v309
		v461 = v310
		v472 = v63
		v473 = v64
		v474 = v65
		goto L7
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v357 = v326 + int32(56)
	v365 = int32(0)
	v384 = v63
	v385 = v64
	v386 = v65
	goto L49
L49:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v34+v365<<(uint(int32(2))%32))))
	if int32(0) < v396 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v456 = v326
	v457 = v306
	v458 = v307
	v459 = v308
	v460 = v309
	v461 = v310
	v472 = v440
	v473 = v441
	v474 = v442
	goto L7
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v326+int32(16)+v365<<(uint(int32(3))%32)))) = v443
	v446 = v365 + int32(1)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v446 < v447 {
		v365 = v446
		v384 = v440
		v385 = v441
		v386 = v442
		goto L49
	} else {
		goto L61
	}
L52:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v326)+72))
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399)+18)))
	if base.Ui32(v400&int32(2047)) < base.Ui32(v396) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v308
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v320)
	v438 = F_heap_getsysattr(m, v357, v396, v31+int32(31))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L6
	} else {
		goto L60
	}
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v308
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v320)
	v414 = F_getmissingattr(m, v389, v396, v31+int32(31))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v385
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v308
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v320)
	v426 = F_fastgetattr_4(m, v357, v396, v389, v31+int32(31))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L59
	}
L58:
	;
	v440 = v384
	v441 = v385
	v442 = v414
	v443 = v414
	goto L51
L59:
	;
	v440 = v426
	v441 = v385
	v442 = v386
	v443 = v426
	goto L51
L60:
	;
	v440 = v384
	v441 = v438
	v442 = v386
	v443 = v438
	goto L51
L61:
	;
	goto L50
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v491
	v496 = v491
	goto L64
L63:
	;
	v496 = v492
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v456))) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v456)+4)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v456
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v502 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v501 + v502
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[5]))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+4)) = v507 + v502
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v511 <= v512<<(uint(v502)%32) {
		v812 = v456
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v472
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v459
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v488)
	v526 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v526 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v472
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v459
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v488)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v528
	F_errmsg_internal(m, int32(_a_F_CatalogCacheCreateEntry_2), v31)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L6
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v472
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v459
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v488)
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_CatalogCacheCreateEntry[0]))
	v577 = F_MemoryContextAllocZero(m, v574, v564<<(uint(int32(4))%32))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v472
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v459
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v488)
	F_errfinish(m, int32(_a_F_CatalogCacheCreateEntry_3), int32(1010), int32(_a_F_CatalogCacheCreateEntry_4))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v580 = v564 << (uint(int32(1)) % 32)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v581 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v593 = v581
	v601 = int32(0)
	goto L76
L74:
	;
	goto L75
L75:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+280)) = v472
	*(*int64)(unsafe.Add(mBase, uint32(v31)+272)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v31)+288)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v31)+300)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v31)+304)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v31)+308)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v31)+312)) = v459
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)) = uint8(v488)
	F_pfree(m, v735)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L88
	}
L76:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v618 = v615 + v601<<(uint(int32(3))%32)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)+4))
	v620 = int32(0)
	if base.B2i32(v619 == v620)|base.B2i32(v619 == v618) == v620 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v630 = v619
	goto L81
L79:
	;
	v682 = v593
	goto L80
L80:
	;
	v705 = v601 + int32(1)
	if v705 < v682 {
		v593 = v682
		v601 = v705
		goto L76
	} else {
		goto L87
	}
L81:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+4)) = v656
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	*(*int32)(unsafe.Add(mBase, uint32(v656))) = v658
	v663 = v577 + v654&(v580-int32(1))<<(uint(int32(3))%32)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+4))
	if v664 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v682 = v675
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v663)+4)) = v663
	*(*int32)(unsafe.Add(mBase, uint32(v663))) = v663
	goto L85
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+4)) = v663
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	*(*int32)(unsafe.Add(mBase, uint32(v630))) = v670
	*(*int32)(unsafe.Add(mBase, uint32(v670)+4)) = v630
	*(*int32)(unsafe.Add(mBase, uint32(v663))) = v630
	if v618 != v656 {
		v630 = v656
		goto L81
	} else {
		goto L86
	}
L86:
	;
	goto L82
L87:
	;
	goto L77
L88:
	;
	goto L5
L89:
	;
	v779 = int32(v775)
	m.G0 = v31
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	if v31+int32(24) == v785 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	m.ExcPending = 1
	goto L98
L91:
	;
	if v789 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	v789 = v787
	goto L94
L93:
	;
	v789 = int32(0)
	goto L94
L94:
	;
	goto L91
L95:
	;
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+319)))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v31)+312))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v31)+308))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v31)+304))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v31)+300))
	v795 = *(*int64)(unsafe.Add(mBase, uint32(v31)+288))
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v31)+280))
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v31)+272))
	v44 = v789
	v46 = v781
	v48 = v792
	v49 = v790
	v50 = v791
	v51 = v793
	v52 = v794
	v63 = v796
	v64 = v797
	v65 = v795
	goto L2
L96:
	;
	goto L97
L97:
	;
	F___wasm_longjmp(m, v782, v781)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	return int32(0)
L99:
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
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(394)
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
