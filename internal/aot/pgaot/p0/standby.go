package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_standby_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int64
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
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
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int64
	_ = v245
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int64
	_ = v604
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v645 int32
	_ = v645
	var v646 int64
	_ = v646
	var v648 int64
	_ = v648
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int64
	_ = v664
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int64
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int64
	_ = v707
	var v710 int64
	_ = v710
	var v714 int32
	_ = v714
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v21, v22, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = v20 & int32(240)
	switch v27 - int32(16) {
	case 0:
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L3
	case 16:
		goto L4
	default:
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L236
	}
L4:
	;
	m.G0 = v15 + int32(16)
	return
L5:
	;
	if v27 != 0 {
		goto L3
	} else {
		goto L235
	}
L6:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v33 = m.G0
	v35 = v33 - int32(176)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v37 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v35 + int32(176)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v732 = m.G0
	v734 = v732 - int32(16)
	m.G0 = v734
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v730)+8))
	if v736 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L8:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v266
	if base.Ui32(v266) < base.Ui32(int32(3)) {
		goto L77
	} else {
		goto L78
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if base.Ui32(v41) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_SnapBuildSerialize(m, v17, v30)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L76
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v78 == v40 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	if int32(0) <= v40-v41 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v49 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+68)) = uint32(v30)
	v53 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+64)) = uint32(v53)
	F_errmsg_internal(m, int32(32780), v35-int32(-64))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	F_SnapBuildWaitSnapshot(m, v32, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+52)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v60
	F_errdetail_internal(m, int32(45666), v35+int32(48))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(519946), int32(1277), int32(92891))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	goto L8
L23:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	if base.Ui64(v80) <= base.Ui64(v30) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+36)))
	if v118 != 0 {
		v123 = v37
		goto L34
	} else {
		goto L35
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v30 + int64(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v86 = v85
	goto L28
L27:
	;
	v86 = v40
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v88
	v96 = F_errstart(m, int32(15), v89)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v96 == int32(0) {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+84)) = uint32(v30)
	v102 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+80)) = uint32(v102)
	F_errmsg(m, int32(532965), v35+int32(80))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errdetail(m, int32(613449), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(519946), int32(1315), int32(92891))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L7
L34:
	;
	switch v123 + int32(1) {
	case 0:
		goto L41
	case 1:
		goto L40
	default:
		v214 = v123
		goto L39
	}
L35:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+37)))
	if v119 != 0 {
		v123 = v37
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v120 = F_SnapBuildRestore(m, v17, v30)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v120 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v123 = v122
	goto L34
L39:
	;
	if v214 != int32(1) {
		goto L8
	} else {
		goto L65
	}
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v165))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v164)) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v126 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v131
	v135 = F_errstart(m, int32(15), v126)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v135 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+132)) = uint32(v30)
	v139 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+128)) = uint32(v139)
	F_errmsg(m, int32(533014), v35+int32(128))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	F_SnapBuildWaitSnapshot(m, v32, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+116)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v35)+112)) = v146
	F_errdetail(m, int32(668600), v35+int32(112))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(519946), int32(1365), int32(92891))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	goto L8
L50:
	;
	if v177 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v177 = base.B2i32(base.Ui32(v164) <= base.Ui32(v165))
	goto L50
L52:
	;
	goto L53
L53:
	;
	v177 = base.B2i32(v164-v165 <= int32(0))
	goto L50
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v180
	v184 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v214 = v213
	goto L39
L57:
	;
	if v184 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+164)) = uint32(v30)
	v188 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+160)) = uint32(v188)
	F_errmsg(m, int32(532908), v35+int32(160))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	F_SnapBuildWaitSnapshot(m, v32, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+148)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v35)+144)) = v195
	F_errdetail(m, int32(668600), v35+int32(144))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(519946), int32(1389), int32(92891))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L8
L65:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v17)+60))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v218))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v217)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v230 == int32(0) {
		goto L8
	} else {
		goto L70
	}
L67:
	;
	v230 = base.B2i32(base.Ui32(v217) <= base.Ui32(v218))
	goto L66
L68:
	;
	goto L69
L69:
	;
	v230 = base.B2i32(v217-v218 <= int32(0))
	goto L66
L70:
	;
	v233 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2)
	v239 = F_errstart(m, int32(15), v233)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v239 == int32(0) {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+100)) = uint32(v30)
	v245 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v35)+96)) = uint32(v245)
	F_errmsg(m, int32(532965), v35+int32(96))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errdetail(m, int32(656282), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(519946), int32(1412), int32(92891))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	goto L8
L76:
	;
	goto L8
L77:
	;
	v583 = int32(0)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)+16))
	if v585 == v583 {
		v595 = v583
		goto L167
	} else {
		goto L168
	}
L78:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v274 = F_MemoryContextAlloc(m, v270, v271<<(uint(int32(2))%32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v276 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	if v277 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v282 = v276
	v283 = int32(0)
	v285 = v277
	goto L83
L81:
	;
	v315 = v276
	goto L82
L82:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v326 = v315 << (uint(int32(2)) % 32)
	if v326 != 0 {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291+v283<<(uint(int32(2))%32))))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if int32(0) <= v295-v296 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v315 = v307
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274+v282<<(uint(int32(2))%32)))) = v295
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v307 = v282 + int32(1)
	v308 = v304
	goto L87
L86:
	;
	v307 = v282
	v308 = v285
	goto L87
L87:
	;
	v310 = v283 + int32(1)
	if base.Ui32(v310) < base.Ui32(v308) {
		v282 = v307
		v283 = v310
		v285 = v308
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	v331 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L93
	}
L90:
	;
	v327 = F__emscripten_memcpy_bulkmem(m, v324, v274, v326)
	mBase = m.M
	goto L92
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	if v331 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v333
	F_errmsg_internal(m, int32(61942), v35+int32(32))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v315
	F_pfree(m, v274)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	F_errfinish(m, int32(519946), int32(894), int32(253691))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v353 == int32(0) {
		goto L77
	} else {
		goto L100
	}
L100:
	;
	v361 = int32(0)
	goto L101
L101:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
	v370 = int32(2)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v369+v361<<(uint(v370)%32))))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if base.B2i32(base.Ui32(v370) < base.Ui32(v374))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v373)) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
	v393 = v387 - v391
	if int32(0) < v393 {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	goto L102
L104:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	if v386 != 0 {
		v391 = v361
		goto L103
	} else {
		goto L108
	}
L105:
	;
	v386 = base.B2i32(base.Ui32(v374) <= base.Ui32(v373))
	goto L104
L106:
	;
	goto L107
L107:
	;
	v386 = base.B2i32(int32(0) <= v373-v374)
	goto L104
L108:
	;
	v389 = v361 + int32(1)
	if base.Ui32(v389) < base.Ui32(v387) {
		v361 = v389
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v391 = v389
	goto L103
L110:
	;
	v551 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L161
	}
L111:
	;
	v396 = int32(2)
	v398 = v392 + v391<<(uint(v396)%32)
	v400 = v393 << (uint(v396) % 32)
	if v392 == v398 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	F_pfree(m, v392)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L160
	}
L114:
	;
	goto L110
L115:
	;
	goto L114
L116:
	;
	v404 = v392 + v400
	if base.Ui32(v398-v404) <= base.Ui32(int32(0)-v400<<(uint(int32(1))%32)) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v411 = F___memcpy(m, v392, v398, v400)
	mBase = m.M
	goto L114
L118:
	;
	goto L119
L119:
	;
	v414 = (v392 ^ v398) & int32(3)
	if base.Ui32(v392) < base.Ui32(v398) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v516 == int32(0) {
		goto L115
	} else {
		goto L156
	}
L121:
	;
	if base.Ui32(v494) <= base.Ui32(int32(3)) {
		v515 = v493
		v516 = v494
		v517 = v495
		goto L120
	} else {
		goto L152
	}
L122:
	;
	if v414 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v414 != 0 {
		v476 = v400
		goto L135
	} else {
		goto L136
	}
L125:
	;
	v515 = v398
	v516 = v400
	v517 = v392
	goto L120
L126:
	;
	goto L127
L127:
	;
	if v392&int32(3) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v493 = v398
	v494 = v400
	v495 = v392
	goto L121
L129:
	;
	goto L130
L130:
	;
	v421 = v398
	v422 = v400
	v423 = v392
	goto L131
L131:
	;
	if v422 == int32(0) {
		goto L115
	} else {
		goto L133
	}
L132:
	;
	v493 = v430
	v494 = v432
	v495 = v434
	goto L121
L133:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	*(*uint8)(unsafe.Add(mBase, uint32(v423))) = uint8(v427)
	v429 = int32(1)
	v430 = v421 + v429
	v432 = v422 - v429
	v434 = v423 + v429
	if v434&int32(3) != 0 {
		v421 = v430
		v422 = v432
		v423 = v434
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	if v476 == int32(0) {
		goto L115
	} else {
		goto L148
	}
L136:
	;
	if v404&int32(3) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v441 = v400
	goto L140
L138:
	;
	v456 = v400
	goto L139
L139:
	;
	if base.Ui32(v456) <= base.Ui32(int32(3)) {
		v476 = v456
		goto L135
	} else {
		goto L144
	}
L140:
	;
	if v441 == int32(0) {
		goto L115
	} else {
		goto L142
	}
L141:
	;
	v456 = v447
	goto L139
L142:
	;
	v447 = v441 - int32(1)
	v448 = v392 + v447
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v447))))
	*(*uint8)(unsafe.Add(mBase, uint32(v448))) = uint8(v450)
	if v448&int32(3) != 0 {
		v441 = v447
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v463 = v456
	goto L145
L145:
	;
	v467 = v463 - int32(4)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v398+v467)))
	*(*int32)(unsafe.Add(mBase, uint32(v392+v467))) = v470
	if base.Ui32(int32(3)) < base.Ui32(v467) {
		v463 = v467
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v476 = v467
	goto L135
L147:
	;
	goto L146
L148:
	;
	v483 = v476
	goto L149
L149:
	;
	v487 = v483 - int32(1)
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398+v487))))
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v487))) = uint8(v490)
	if v487 != 0 {
		v483 = v487
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L115
L151:
	;
	goto L150
L152:
	;
	v500 = v493
	v501 = v494
	v502 = v495
	goto L153
L153:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = v504
	v506 = int32(4)
	v507 = v500 + v506
	v509 = v502 + v506
	v511 = v501 - v506
	if base.Ui32(int32(3)) < base.Ui32(v511) {
		v500 = v507
		v501 = v511
		v502 = v509
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v515 = v507
	v516 = v511
	v517 = v509
	goto L120
L155:
	;
	goto L154
L156:
	;
	v522 = v515
	v523 = v516
	v524 = v517
	goto L157
L157:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v526)
	v528 = int32(1)
	v533 = v523 - v528
	if v533 != 0 {
		v522 = v522 + v528
		v523 = v533
		v524 = v524 + v528
		goto L157
	} else {
		goto L159
	}
L158:
	;
	goto L115
L159:
	;
	goto L158
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = int32(0)
	goto L110
L161:
	;
	if v551 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v17)+80))
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v393
	*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v554
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v553
	F_errmsg_internal(m, int32(61870), v35+int32(16))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v393
	goto L77
L165:
	;
	F_errfinish(m, int32(519946), int32(931), int32(253691))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	if v595 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	if v585 == v584+int32(12) {
		v595 = v583
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v585-int32(16))))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
	v595 = v594
	goto L167
L170:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v599 = v598
	goto L172
L171:
	;
	v599 = v595
	goto L172
L172:
	;
	v602 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	if v602 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v605
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v604
	F_errmsg_internal(m, int32(62766), v35)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v619 = m.G0
	v621 = v619 - int32(16)
	m.G0 = v621
	v624 = *(*int32)(unsafe.Add(mBase, _consts[510]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = int32(1)
	if v625 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	F_errfinish(m, int32(519946), int32(1185), int32(132740))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	F_s_lock(m, v624, int32(517668), int32(1688), int32(91771))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v624)+100))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v633))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v599)) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L182:
	;
	goto L181
L183:
	;
	m.G0 = v621 + int32(16)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v687 < int32(2) {
		goto L7
	} else {
		goto L198
	}
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v624)+240)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v624)+236)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = int32(0)
	F_LogicalConfirmReceivedLocation(m, v646)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L197
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = int32(0)
	goto L183
L186:
	;
	if v645 != 0 {
		goto L185
	} else {
		goto L190
	}
L187:
	;
	v645 = base.B2i32(base.Ui32(v599) <= base.Ui32(v633))
	goto L186
L188:
	;
	goto L189
L189:
	;
	v645 = base.B2i32(v599-v633 <= int32(0))
	goto L186
L190:
	;
	v646 = *(*int64)(unsafe.Add(mBase, uint32(v624)+120))
	if base.Ui64(v30) <= base.Ui64(v646) {
		goto L184
	} else {
		goto L191
	}
L191:
	;
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v624)+240))
	if v648 != int64(0) {
		goto L185
	} else {
		goto L192
	}
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v624)+240)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v624)+236)) = v599
	v653 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = v653
	v657 = F_errstart(m, int32(14), v653)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	if v657 == int32(0) {
		goto L183
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = v599
	*(*uint32)(unsafe.Add(mBase, uint32(v621)+8)) = uint32(v30)
	v664 = int64(base.Ui64(v30) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v621)+4)) = uint32(v664)
	F_errmsg_internal(m, int32(532487), v621)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(517668), int32(1731), int32(91771))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	goto L183
L197:
	;
	goto L183
L198:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+8))
	if v691 != v690+int32(4) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v698 = v691 - int32(188)
	goto L201
L200:
	;
	v698 = int32(0)
	goto L201
L201:
	;
	if v691 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v700 = v698
	goto L204
L203:
	;
	v700 = int32(0)
	goto L204
L204:
	;
	if v700 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v701 = *(*int64)(unsafe.Add(mBase, uint32(v700)+48))
	if v701 == int64(0) {
		goto L7
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v707 = *(*int64)(unsafe.Add(mBase, uint32(v706)+136))
	if v707 == int64(0) {
		goto L7
	} else {
		goto L210
	}
L208:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v30, v701)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	goto L7
L210:
	;
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	if v710 == int64(0) {
		goto L7
	} else {
		goto L211
	}
L211:
	;
	F_LogicalIncreaseRestartDecodingForSlot(m, v30, v710)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	goto L7
L213:
	;
	m.G0 = v734 + int32(16)
	goto L4
L214:
	;
	v740 = v730 + int32(4)
	if v736 == v740 {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v743 = v736
	goto L216
L216:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	v756 = v743 - int32(184)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v731))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v757)) == int32(0) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L213
L218:
	;
	if v769 == int32(0) {
		goto L213
	} else {
		goto L222
	}
L219:
	;
	v769 = base.B2i32(base.Ui32(v757) < base.Ui32(v731))
	goto L218
L220:
	;
	goto L221
L221:
	;
	v769 = int32(base.Ui32(v757-v731) >> (uint(int32(31)) % 32))
	goto L218
L222:
	;
	v774 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	if v774 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	*(*int32)(unsafe.Add(mBase, uint32(v734))) = v776
	F_errmsg_internal(m, int32(48891), v734)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v787 = v743 - int32(188)
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787))))
	if v788&int32(16) != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	F_errfinish(m, int32(515132), int32(3142), int32(448699))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v730)+84))
	m.T0[v792].(func(*base.Module, int32, int32, int64))(m, v730, v787, int64(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	F_ReorderBufferCleanupTXN(m, v730, v787)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	if v754 != v740 {
		v743 = v754
		goto L216
	} else {
		goto L234
	}
L234:
	;
	goto L217
L235:
	;
	goto L4
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v27
	F_errmsg_internal(m, int32(63439), v15)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(519616), int32(397), int32(429772))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
