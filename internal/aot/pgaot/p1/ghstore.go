package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v44 int32
	_ = v44
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
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
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
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v795 int32
	_ = v795
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v839 int32
	_ = v839
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v24 == v2 {
		v41 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v41&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v28 == int32(0) {
		v41 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v31 != int32(7) {
		v41 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v34 != int32(17) {
		v41 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	v41 = v37 ^ int32(1)
	goto L2
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = F_get_fn_opclass_options(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v52 = int32(128)
	goto L9
L9:
	;
	v53 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v53)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
	if v56&int32(4) != 0 {
		v856 = v53
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v52 = v49 << (uint(int32(3)) % 32)
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L10
	} else {
		goto L132
	}
L13:
	;
	m.G0 = v17 + int32(16)
	return v856
L14:
	;
	v60 = v22 + int32(8)
	switch v20 - int32(7) {
	case 0, 6:
		goto L18
	default:
		goto L12
	case 2:
		goto L17
	case 3:
		goto L15
	case 4:
		goto L16
	}
L15:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v684 = F_pg_detoast_datum(m, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L10
	} else {
		goto L108
	}
L16:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v513 = F_pg_detoast_datum(m, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L10
	} else {
		goto L86
	}
L17:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v368 = F_pg_detoast_datum_packed(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L10
	} else {
		goto L63
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v64 = F_hstoreUpgrade(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v68 = v66 & int32(268435455)
	if v68 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v856 = v53
	goto L13
L21:
	;
	goto L22
L22:
	;
	v72 = v64 + int32(8)
	v75 = v72 + v68<<(uint(int32(3))%32)
	v82 = v2
	goto L23
L23:
	;
	v92 = v72 + v82<<(uint(int32(3))%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v856 = v351
	goto L13
L25:
	;
	v110 = int32(0)
	if v108 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v108 = v93 & int32(1073741823)
	v109 = v75
	goto L25
L27:
	;
	goto L28
L28:
	;
	v98 = int32(1073741823)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92-int32(4))))
	v104 = v102 & v98
	v108 = v93&v98 - v104
	v109 = v104 + v75
	goto L25
L29:
	;
	v112 = int32(1)
	v114 = int32(-1)
	if v108 != v112 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v205 = v110
	goto L31
L31:
	;
	v206 = base.I32_rem_u_s(v205, v52)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v206)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v210)>>(uint(v206&int32(7))%32))&int32(1) == int32(0) {
		v856 = v110
		goto L13
	} else {
		goto L41
	}
L32:
	;
	v120 = v109
	v121 = v114
	v122 = int32(0)
	goto L35
L33:
	;
	v162 = v109
	v163 = v114
	goto L34
L34:
	;
	if v108&v112 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v137 = int32(255)
	v139 = int32(2)
	v142 = *(*int32)(unsafe.Add(mBase, uint32((v121^v135)&v137<<(uint(v139)%32))+uint32(_consts[1191])))
	v143 = int32(8)
	v145 = v142 ^ int32(base.Ui32(v121)>>(uint(v143)%32))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32((v145^v146)&v137<<(uint(v139)%32))+uint32(_consts[1191])))
	v156 = v153 ^ int32(base.Ui32(v145)>>(uint(v143)%32))
	v158 = v120 + v139
	v160 = v122 + v139
	if v160 != v108&int32(-2) {
		v120 = v158
		v121 = v156
		v122 = v160
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v162 = v158
	v163 = v156
	goto L34
L37:
	;
	goto L36
L38:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32((v163^v176)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1191])))
	v188 = v184 ^ int32(base.Ui32(v163)>>(uint(int32(8))%32))
	goto L40
L39:
	;
	v188 = v163
	goto L40
L40:
	;
	v205 = v188 ^ int32(-1)
	goto L31
L41:
	;
	v223 = v72 + v82<<(uint(int32(3))%32) + int32(4)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v224&int32(1073741824) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v224 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v351 = int32(1)
	goto L44
L44:
	;
	if v351 == int32(0) {
		v856 = v351
		goto L13
	} else {
		goto L61
	}
L45:
	;
	if v241 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v241 = v224 & int32(1073741823)
	v242 = v75
	goto L45
L47:
	;
	goto L48
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v223-int32(4))))
	v237 = v235 & int32(1073741823)
	v241 = v224 - v237
	v242 = v237 + v75
	goto L45
L49:
	;
	v243 = int32(1)
	v245 = int32(-1)
	if v241 != v243 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v337 = int32(0)
	goto L51
L51:
	;
	v338 = base.I32_rem_u_s(v337, v52)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v338)>>(uint(int32(3))%32))))))
	v351 = int32(base.Ui32(v342)>>(uint(v338&int32(7))%32)) & int32(1)
	goto L44
L52:
	;
	v251 = v242
	v252 = v245
	v253 = int32(0)
	goto L55
L53:
	;
	v293 = v242
	v294 = v245
	goto L54
L54:
	;
	if v241&v243 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v268 = int32(255)
	v270 = int32(2)
	v273 = *(*int32)(unsafe.Add(mBase, uint32((v252^v266)&v268<<(uint(v270)%32))+uint32(_consts[1191])))
	v274 = int32(8)
	v276 = v273 ^ int32(base.Ui32(v252)>>(uint(v274)%32))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32((v276^v277)&v268<<(uint(v270)%32))+uint32(_consts[1191])))
	v287 = v284 ^ int32(base.Ui32(v276)>>(uint(v274)%32))
	v289 = v251 + v270
	v291 = v253 + v270
	if v291 != v241&int32(-2) {
		v251 = v289
		v252 = v287
		v253 = v291
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v293 = v289
	v294 = v287
	goto L54
L57:
	;
	goto L56
L58:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v315 = *(*int32)(unsafe.Add(mBase, uint32((v294^v307)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1191])))
	v319 = v315 ^ int32(base.Ui32(v294)>>(uint(int32(8))%32))
	goto L60
L59:
	;
	v319 = v294
	goto L60
L60:
	;
	v337 = v319 ^ int32(-1)
	goto L51
L61:
	;
	v365 = v82 + int32(1)
	if base.Ui32(v365) < base.Ui32(v68) {
		v82 = v365
		goto L23
	} else {
		goto L62
	}
L62:
	;
	goto L24
L63:
	;
	v370 = int32(1)
	v371 = v368 + v370
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	v374 = v372 & v370
	if v372 == v370 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v502 = base.I32_rem_u_s(v501, v52)
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v502)>>(uint(int32(3))%32))))))
	v856 = int32(base.Ui32(v506)>>(uint(v502&int32(7))%32)) & int32(1)
	goto L13
L65:
	;
	if v374 != 0 {
		goto L73
	} else {
		goto L74
	}
L66:
	;
	if v400 != 0 {
		v402 = v400
		goto L65
	} else {
		goto L72
	}
L67:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	if base.Ui32((v378-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v402 = int32(4)
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v389 = int32(1)
	if v374 != 0 {
		v400 = int32(base.Ui32(v372)>>(uint(v389)%32)) - v389
		goto L66
	} else {
		goto L71
	}
L70:
	;
	v400 = base.B2i32(v378 == int32(18)) << (uint(int32(4)) % 32)
	goto L66
L71:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v400 = int32(base.Ui32(v393)>>(uint(int32(2))%32)) - int32(4)
	goto L66
L72:
	;
	v501 = int32(0)
	goto L64
L73:
	;
	v406 = v371
	goto L75
L74:
	;
	v406 = v368 + int32(4)
	goto L75
L75:
	;
	v407 = int32(1)
	if v402 == v407 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v402&v407 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v458 = v406
	v459 = int32(-1)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v416 = v406
	v417 = int32(-1)
	v418 = int32(0)
	goto L80
L80:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	v433 = int32(255)
	v435 = int32(2)
	v438 = *(*int32)(unsafe.Add(mBase, uint32((v417^v431)&v433<<(uint(v435)%32))+uint32(_consts[1191])))
	v439 = int32(8)
	v441 = v438 ^ int32(base.Ui32(v417)>>(uint(v439)%32))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+1)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32((v441^v442)&v433<<(uint(v435)%32))+uint32(_consts[1191])))
	v452 = v449 ^ int32(base.Ui32(v441)>>(uint(v439)%32))
	v454 = v416 + v435
	v456 = v418 + v435
	if v456 != v402&int32(-2) {
		v416 = v454
		v417 = v452
		v418 = v456
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v458 = v454
	v459 = v452
	goto L76
L82:
	;
	goto L81
L83:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	v480 = *(*int32)(unsafe.Add(mBase, uint32((v459^v472)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1191])))
	v484 = v480 ^ int32(base.Ui32(v459)>>(uint(int32(8))%32))
	goto L85
L84:
	;
	v484 = v459
	goto L85
L85:
	;
	v501 = v484 ^ int32(-1)
	goto L64
L86:
	;
	F_deconstruct_array_builtin(m, v513, int32(25), v17+int32(12), v17+int32(8), v17+int32(4))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v524 <= int32(0) {
		v856 = v53
		goto L13
	} else {
		goto L88
	}
L88:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v535 = v2
	goto L89
L89:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535+v528))))
	if v545 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v856 = v667
	goto L13
L91:
	;
	v548 = int32(2)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v527+v535<<(uint(v548)%32))))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	v554 = int32(base.Ui32(v552) >> (uint(v548) % 32))
	v556 = v554 - int32(4)
	if v556 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v667 = int32(1)
	goto L93
L93:
	;
	if v667 == int32(0) {
		v856 = v667
		goto L13
	} else {
		goto L106
	}
L94:
	;
	v558 = v551 + int32(4)
	v559 = int32(-1)
	if v554 != int32(5) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v653 = int32(0)
	goto L96
L96:
	;
	v654 = base.I32_rem_u_s(v653, v52)
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v654)>>(uint(int32(3))%32))))))
	v667 = int32(base.Ui32(v658)>>(uint(v654&int32(7))%32)) & int32(1)
	goto L93
L97:
	;
	v565 = v558
	v566 = v559
	v567 = int32(0)
	goto L100
L98:
	;
	v607 = v558
	v608 = v559
	goto L99
L99:
	;
	if v552&int32(4) != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	v582 = int32(255)
	v584 = int32(2)
	v587 = *(*int32)(unsafe.Add(mBase, uint32((v566^v580)&v582<<(uint(v584)%32))+uint32(_consts[1191])))
	v588 = int32(8)
	v590 = v587 ^ int32(base.Ui32(v566)>>(uint(v588)%32))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32((v590^v591)&v582<<(uint(v584)%32))+uint32(_consts[1191])))
	v601 = v598 ^ int32(base.Ui32(v590)>>(uint(v588)%32))
	v603 = v565 + v584
	v605 = v567 + v584
	if v605 != v556&int32(-2) {
		v565 = v603
		v566 = v601
		v567 = v605
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v607 = v603
	v608 = v601
	goto L99
L102:
	;
	goto L101
L103:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v631 = *(*int32)(unsafe.Add(mBase, uint32((v608^v623)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1191])))
	v635 = v631 ^ int32(base.Ui32(v608)>>(uint(int32(8))%32))
	goto L105
L104:
	;
	v635 = v608
	goto L105
L105:
	;
	v653 = v635 ^ int32(-1)
	goto L96
L106:
	;
	v681 = v535 + int32(1)
	if v681 < v524 {
		v535 = v681
		goto L89
	} else {
		goto L107
	}
L107:
	;
	goto L90
L108:
	;
	F_deconstruct_array_builtin(m, v684, int32(25), v17+int32(12), v17+int32(8), v17+int32(4))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v695 <= int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v856 = int32(0)
	goto L13
L111:
	;
	goto L112
L112:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v707 = v2
	goto L113
L113:
	;
	v715 = int32(0)
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707+v700))))
	if v717 == v715 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v856 = v839
	goto L13
L115:
	;
	v720 = int32(2)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v699+v707<<(uint(v720)%32))))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)))
	v726 = int32(base.Ui32(v724) >> (uint(v720) % 32))
	v728 = v726 - int32(4)
	if v728 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v839 = v715
	goto L117
L117:
	;
	if v839 != 0 {
		v856 = v839
		goto L13
	} else {
		goto L130
	}
L118:
	;
	v730 = v723 + int32(4)
	v731 = int32(-1)
	if v726 != int32(5) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v825 = int32(0)
	goto L120
L120:
	;
	v826 = base.I32_rem_u_s(v825, v52)
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v826)>>(uint(int32(3))%32))))))
	v839 = int32(base.Ui32(v830)>>(uint(v826&int32(7))%32)) & int32(1)
	goto L117
L121:
	;
	v737 = v730
	v738 = v731
	v739 = int32(0)
	goto L124
L122:
	;
	v779 = v730
	v780 = v731
	goto L123
L123:
	;
	if v724&int32(4) != 0 {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737))))
	v754 = int32(255)
	v756 = int32(2)
	v759 = *(*int32)(unsafe.Add(mBase, uint32((v738^v752)&v754<<(uint(v756)%32))+uint32(_consts[1191])))
	v760 = int32(8)
	v762 = v759 ^ int32(base.Ui32(v738)>>(uint(v760)%32))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737)+1)))
	v770 = *(*int32)(unsafe.Add(mBase, uint32((v762^v763)&v754<<(uint(v756)%32))+uint32(_consts[1191])))
	v773 = v770 ^ int32(base.Ui32(v762)>>(uint(v760)%32))
	v775 = v737 + v756
	v777 = v739 + v756
	if v777 != v728&int32(-2) {
		v737 = v775
		v738 = v773
		v739 = v777
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v779 = v775
	v780 = v773
	goto L123
L126:
	;
	goto L125
L127:
	;
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	v803 = *(*int32)(unsafe.Add(mBase, uint32((v780^v795)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1191])))
	v807 = v803 ^ int32(base.Ui32(v780)>>(uint(int32(8))%32))
	goto L129
L128:
	;
	v807 = v780
	goto L129
L129:
	;
	v825 = v807 ^ int32(-1)
	goto L120
L130:
	;
	v851 = v707 + int32(1)
	if v851 < v695 {
		v707 = v851
		goto L113
	} else {
		goto L131
	}
L131:
	;
	goto L114
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v20
	F_errmsg_internal(m, int32(482362), v17)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(492660), int32(608), int32(92756))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L10
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
