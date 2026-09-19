package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_advance_array_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v750 int32
	_ = v750
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1022 int32
	_ = v1022
	v8 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	if l6 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v27 + int32(16)
	return v1022
L2:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v59 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+26)) = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v36 = v34 - int32(1)
	if v36 <= l5 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v36*int32(48)))))
	if v42&int32(32) != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v36
	v46 = int32(0)
	v52 = F__bt_check_compare(m, l0, v31, l2, l3, l4, v46, v46, v27+int32(15), v27+int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v52 == int32(0) {
		v1022 = v8
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L2
L11:
	;
	if l6 != 0 {
		goto L240
	} else {
		goto L241
	}
L12:
	;
	v62 = int32(1)
	v851 = v62
	v858 = v62
	v860 = v8
	goto L11
L13:
	;
	goto L14
L14:
	;
	v65 = int32(1)
	v77 = int32(0)
	v78 = v8
	v80 = v65
	v86 = v8
	v87 = v65
	v88 = v8
	v89 = v8
	goto L16
L15:
	;
	v829 = int32(0)
	if base.B2i32(l6 == v829)|base.B2i32(v824 == v829) != 0 {
		v851 = v818
		v858 = v825
		v860 = v124
		goto L11
	} else {
		goto L238
	}
L16:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v94 = v91 + v77*int32(48)
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)))
	if v95 == int32(3) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v465&int32(1) == int32(0) {
		v818 = v466
		v824 = v470
		v825 = v471
		goto L15
	} else {
		goto L148
	}
L18:
	;
	if v77 < l5 {
		v465 = v78
		v466 = v80
		v470 = v86
		v471 = v87
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v98 = int32(0)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	if v99&int32(32) == v98 {
		v122 = v98
		v123 = v88
		v124 = v89
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v110 = int32(0)
	switch v31 + int32(1) {
	case 0:
		goto L24
	default:
		v122 = v110
		v123 = v88
		v124 = v89
		goto L18
	case 2:
		goto L25
	}
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v122 = v104 + v88<<(uint(int32(5))%32)
	v123 = v88 + int32(1)
	v124 = v89
	goto L18
L23:
	;
	v122 = v110
	v123 = v88
	v124 = int32(1)
	goto L18
L24:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)))
	if v116&int32(1) == int32(0) {
		v122 = v110
		v123 = v88
		v124 = v89
		goto L18
	} else {
		goto L27
	}
L25:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)))
	if v113&int32(2) != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v122 = v110
	v123 = v88
	v124 = v89
	goto L18
L27:
	;
	goto L23
L28:
	;
	v473 = v77 + int32(1)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v473 < v474 {
		v77 = v473
		v78 = v465
		v80 = v466
		v86 = v470
		v87 = v471
		v88 = v123
		v89 = v124
		goto L16
	} else {
		goto L147
	}
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v128 = v126 & int32(_a_F__bt_advance_array_keys_0)
	if v128 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if base.B2i32(l5 != v77)|v122 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if v131 <= l3 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+18)) = uint8(v133)
	goto L30
L33:
	;
	v465 = v360
	v466 = int32(0)
	v470 = int32(1)
	v471 = v362
	goto L28
L34:
	;
	v459 = int32(1)
	v465 = v360
	v466 = v459
	v470 = v459
	v471 = v362
	goto L28
L35:
	;
	v818 = int32(1)
	v824 = v86
	v825 = int32(0)
	goto L15
L36:
	;
	v465 = int32(1)
	v466 = v452
	v470 = v86
	v471 = v453
	goto L28
L37:
	;
	v139 = int32(0)
	v452 = v139
	v453 = v139
	goto L36
L38:
	;
	goto L39
L39:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)))
	if base.B2i32(v122|v128 == int32(0))|base.B2i32(v144 != int32(3)) != 0 {
		v465 = v78
		v466 = v80
		v470 = v86
		v471 = v87
		goto L28
	} else {
		goto L40
	}
L40:
	;
	if v78&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v179 | int32(_a_F__bt_advance_array_keys_1)
	v452 = v80
	v453 = v87
	goto L36
L42:
	;
	if v122 == int32(0) {
		v452 = v80
		v453 = v87
		goto L36
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v80 != 0 {
		goto L61
	} else {
		goto L62
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v152 != int32(-1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v31 != int32(-1) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+18)))
	if v168 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v160 = v152 - int32(1)
	goto L51
L50:
	;
	v160 = int32(0)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v160<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v166
	v452 = v80
	v453 = v87
	goto L36
L52:
	;
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v179 = v177 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v179
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+19)))
	if base.B2i32(base.B2i32(v177&int32(33554432) == v175) == base.B2i32(v31 == int32(-1)))|base.B2i32(v188 != int32(1)) == v175 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	if v169 == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v179 | int32(65)
	v452 = v80
	v453 = v87
	goto L36
L57:
	;
	goto L58
L58:
	;
	if v31 != int32(-1) {
		goto L41
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v179 | int32(_a_F__bt_advance_array_keys_2)
	v452 = v80
	v453 = v87
	goto L36
L60:
	;
	v263 = F_index_getattr_2(m, l2, v202, l4, v27+int32(15))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L82
	}
L61:
	;
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+4)))
	if v202 <= l3 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v205 = int32(0)
	if v122 == v205 {
		v465 = v205
		v466 = v80
		v470 = v86
		v471 = v87
		goto L28
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v208 != int32(-1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v211 = int32(1)
	if v31 != v211 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+18)))
	if v224 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v216 = v208 - v211
	goto L71
L70:
	;
	v216 = int32(0)
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218+v216<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v222
	v465 = v205
	v466 = v80
	v470 = v86
	v471 = v87
	goto L28
L72:
	;
	v231 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v235 = v233 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v235
	v241 = int32(1)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+19)))
	if base.B2i32(base.B2i32(v233&int32(33554432) == v231) == base.B2i32(v31 == v241))|base.B2i32(v244 != v241) == v231 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	if v225 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_pfree(m, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v235 | int32(65)
	v465 = v205
	v466 = v80
	v470 = v86
	v471 = v87
	goto L28
L77:
	;
	goto L78
L78:
	;
	if v31 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v235 | int32(_a_F__bt_advance_array_keys_2)
	v465 = v205
	v466 = v80
	v470 = v86
	v471 = v87
	goto L28
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v235 | int32(_a_F__bt_advance_array_keys_1)
	v465 = v205
	v466 = v80
	v470 = v86
	v471 = v87
	goto L28
L82:
	;
	if v122 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v333 = int32(0)
	if base.B2i32(l6 == v333)|base.B2i32(v128 == v333) == v333 {
		goto L113
	} else {
		goto L114
	}
L84:
	;
	v329 = v324
	v331 = int32(0)
	goto L83
L85:
	;
	v266 = l6 & base.B2i32(l5 == v77)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v267 == int32(-1) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v285 = int32(1)
	v286 = v284 & v285
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)))
	if v287 == v285 {
		goto L94
	} else {
		goto L95
	}
L88:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)))
	F__bt_binsrch_skiparray_skey(m, v266, v31, v263, v270, v122, v94, v27+int32(8))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)))
	v282 = F__bt_binsrch_array_skey(m, v275+v77*int32(28), v266, v31, v263, v279, v122, v94, v27+int32(8))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L92
	}
L91:
	;
	v324 = v270
	goto L84
L92:
	;
	v329 = v279
	v331 = v282
	goto L83
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v320
	v324 = v287
	goto L84
L94:
	;
	if v286 != 0 {
		v320 = int32(0)
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	if v286 != 0 {
		goto L101
	} else {
		goto L102
	}
L97:
	;
	if v284&int32(33554432) != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v295 = int32(-1)
	goto L100
L99:
	;
	v295 = int32(1)
	goto L100
L100:
	;
	v320 = v295
	goto L93
L101:
	;
	if v284&int32(33554432) != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	v307 = F_FunctionCall2Coll(m, v301+v77*int32(28), v305, v263, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L8
	} else {
		goto L107
	}
L104:
	;
	v300 = int32(1)
	goto L106
L105:
	;
	v300 = int32(-1)
	goto L106
L106:
	;
	v320 = v300
	goto L93
L107:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+3)))
	if v309&int32(1) == int32(0) {
		v320 = v307
		goto L93
	} else {
		goto L108
	}
L108:
	;
	v315 = int32(0)
	if v307 < v315 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v319 = int32(1)
	goto L111
L110:
	;
	v319 = v315 - v307
	goto L111
L111:
	;
	v320 = v319
	goto L93
L112:
	;
	if v122 == int32(0) {
		v465 = v360
		v466 = v361
		v470 = v86
		v471 = v362
		goto L28
	} else {
		goto L117
	}
L113:
	;
	v342 = int32(0)
	v354 = base.B2i32(v332 == v342)
	v359 = int32(base.Ui32(v332) >> (uint(int32(31)) % 32))
	v360 = base.B2i32(v31 == int32(1))&base.B2i32(v342 < v332) | base.B2i32(v31 == int32(-1))&base.B2i32(v332 < v342)
	v361 = v354
	v362 = v354 & v87
	goto L112
L114:
	;
	goto L115
L115:
	;
	if v332 != 0 {
		goto L35
	} else {
		goto L116
	}
L116:
	;
	v356 = int32(0)
	v359 = v356
	v360 = v356
	v361 = int32(1)
	v362 = v87
	goto L112
L117:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v365 == int32(-1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+18)))
	if v361 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	if v438 == v331 {
		v465 = v360
		v466 = v361
		v470 = v86
		v471 = v362
		goto L28
	} else {
		goto L146
	}
L121:
	;
	if v368&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	if v329 != 0 {
		goto L134
	} else {
		goto L135
	}
L124:
	;
	v379 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v383 = v381 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v383
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+19)))
	if base.B2i32(v359 == base.B2i32(v381&int32(33554432) == v379))|base.B2i32(v390 != int32(1)) == v379 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	if v373 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	F_pfree(m, v373)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v383 | int32(65)
	goto L33
L129:
	;
	goto L130
L130:
	;
	if v359 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v383 | int32(_a_F__bt_advance_array_keys_2)
	goto L33
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v383 | int32(_a_F__bt_advance_array_keys_1)
	goto L33
L134:
	;
	if v368&int32(1) != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	goto L136
L136:
	;
	if v368&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v415&int32(-7864386) | int32(65)
	goto L34
L138:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	if v407 == int32(0) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	F_pfree(m, v407)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v429 & int32(-7864386)
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+18)))
	v434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v122)+16)))
	v435 = F_datumCopy(m, v263, v433, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L8
	} else {
		goto L145
	}
L142:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v94)+44))
	if v423 == int32(0) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	F_pfree(m, v423)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L8
	} else {
		goto L144
	}
L144:
	;
	goto L141
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v435
	goto L34
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v331
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441+v331<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+44)) = v445
	v465 = v360
	v466 = v361
	v470 = v86
	v471 = v362
	goto L28
L147:
	;
	goto L17
L148:
	;
	v480 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v484 = v482 - int32(1)
	if v480 <= v484 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v489 = base.B2i32(v31 == int32(1))
	v500 = v484
	v509 = v470
	goto L152
L150:
	;
	goto L151
L151:
	;
	F__bt_start_array_keys(m, l0, int32(0)-v31)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L8
	} else {
		goto L237
	}
L152:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v481)+8))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v481)+20))
	v518 = v515 + v500<<(uint(int32(5))%32)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v522 = v514 + v519*int32(48)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v523 == int32(-1) {
		goto L158
	} else {
		goto L159
	}
L153:
	;
	goto L151
L154:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	if v718 != int32(-1) {
		goto L220
	} else {
		goto L221
	}
L155:
	;
	v633 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v633)
	v635 = int32(1)
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v639 = int32(33554433)
	if v636&int32(_a_F__bt_advance_array_keys_2)|base.B2i32(v636&v639 == v639) != 0 {
		v717 = v635
		goto L154
	} else {
		goto L191
	}
L156:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	if v621 <= int32(0) {
		v717 = v509
		goto L154
	} else {
		goto L190
	}
L157:
	;
	if v531&int32(1) != 0 {
		goto L166
	} else {
		goto L167
	}
L158:
	;
	if v489 == int32(0) {
		goto L155
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	if v31 != int32(1) {
		goto L156
	} else {
		goto L164
	}
L161:
	;
	v528 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v528)
	v530 = int32(1)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v531&int32(_a_F__bt_advance_array_keys_1)|base.B2i32(v531&int32(33554433) == v530) != 0 {
		v717 = v530
		goto L154
	} else {
		goto L162
	}
L162:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	if v539 != 0 {
		goto L157
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v531 | int32(_a_F__bt_advance_array_keys_3)
	v818 = v466
	v824 = v530
	v825 = v471
	goto L15
L164:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	if v523-int32(1) <= v545 {
		v717 = v509
		goto L154
	} else {
		goto L165
	}
L165:
	;
	v550 = v545 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v518)+12)) = v550
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v552+v550<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v556
	v818 = v466
	v824 = v509
	v825 = v471
	goto L15
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v531 & int32(-1048642)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	v566 = int32(*(*int16)(unsafe.Add(mBase, uint32(v518)+16)))
	v567 = F_datumCopy(m, v564, v565, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L8
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v574 = m.T0[v573].(func(*base.Module, int32, int32, int32) int32)(m, v487, v570, v27+int32(8))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L8
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v567
	v818 = v466
	v824 = v530
	v825 = v471
	goto L15
L170:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	if v576 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+19)))
	if v579 != int32(1) {
		v717 = v530
		goto L154
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v518)+28))
	if v601 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L174:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v582&int32(33554432) != 0 {
		v717 = v530
		goto L154
	} else {
		goto L175
	}
L175:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v585 != 0 {
		v593 = v582
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v593&int32(-7864386) | int32(65)
	v818 = v466
	v824 = v530
	v825 = v471
	goto L15
L177:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	if v586 == int32(0) {
		v593 = v582
		goto L176
	} else {
		goto L178
	}
L178:
	;
	F_pfree(m, v586)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L8
	} else {
		goto L179
	}
L179:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v593 = v591
	goto L176
L180:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v613 != 0 {
		goto L186
	} else {
		goto L187
	}
L181:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v601)+44))
	v608 = F_FunctionCall2Coll(m, v601+int32(16), v606, v574, v607)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L8
	} else {
		goto L182
	}
L182:
	;
	if v608 != 0 {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v610 != 0 {
		v717 = v530
		goto L154
	} else {
		goto L184
	}
L184:
	;
	F_pfree(m, v574)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L8
	} else {
		goto L185
	}
L185:
	;
	v717 = v530
	goto L154
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v574
	v818 = v466
	v824 = v530
	v825 = v471
	goto L15
L187:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	if v614 == int32(0) {
		goto L186
	} else {
		goto L188
	}
L188:
	;
	F_pfree(m, v614)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	v625 = v621 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v518)+12)) = v625
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v627+v625<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v631
	v818 = v466
	v824 = v509
	v825 = v471
	goto L15
L191:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	if v644 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v636 | int32(_a_F__bt_advance_array_keys_4)
	v818 = v466
	v824 = v635
	v825 = v471
	goto L15
L193:
	;
	goto L194
L194:
	;
	if v636&int32(1) != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v636 & int32(-524354)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v518)+20))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	v658 = int32(*(*int16)(unsafe.Add(mBase, uint32(v518)+16)))
	v659 = F_datumCopy(m, v656, v657, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L8
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v644)+8))
	v666 = m.T0[v665].(func(*base.Module, int32, int32, int32) int32)(m, v487, v662, v27+int32(8))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L8
	} else {
		goto L199
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v659
	v818 = v466
	v824 = v635
	v825 = v471
	goto L15
L199:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	if v668 == int32(1) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+19)))
	if v671 != int32(1) {
		v717 = v635
		goto L154
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v518)+24))
	if v695 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L203:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if v674&int32(33554432) == int32(0) {
		v717 = v635
		goto L154
	} else {
		goto L204
	}
L204:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v679 != 0 {
		v687 = v674
		goto L205
	} else {
		goto L206
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v687&int32(-7864386) | int32(65)
	v818 = v466
	v824 = v635
	v825 = v471
	goto L15
L206:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	if v680 == int32(0) {
		v687 = v674
		goto L205
	} else {
		goto L207
	}
L207:
	;
	F_pfree(m, v680)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L8
	} else {
		goto L208
	}
L208:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v687 = v685
	goto L205
L209:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v707 != 0 {
		goto L215
	} else {
		goto L216
	}
L210:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v695)+12))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v695)+44))
	v702 = F_FunctionCall2Coll(m, v695+int32(16), v700, v666, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L8
	} else {
		goto L211
	}
L211:
	;
	if v702 != 0 {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v704 != 0 {
		v717 = v635
		goto L154
	} else {
		goto L213
	}
L213:
	;
	F_pfree(m, v666)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L8
	} else {
		goto L214
	}
L214:
	;
	v717 = v635
	goto L154
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v666
	v818 = v466
	v824 = v635
	v825 = v471
	goto L15
L216:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	if v708 == int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	F_pfree(m, v708)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L8
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	if int32(0) < v500 {
		v500 = v500 - int32(1)
		v509 = v717
		goto L152
	} else {
		goto L236
	}
L220:
	;
	if v31 == int32(1) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+18)))
	if v732 != 0 {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v724 = int32(0)
	goto L225
L224:
	;
	v724 = v718 - int32(1)
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+12)) = v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v518)+8))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v726+v724<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v730
	goto L219
L226:
	;
	v739 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v522)+44)) = v739
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v743 = v741 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v743
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+19)))
	if base.B2i32(v489 == base.B2i32(v741&int32(33554432) == v739))|base.B2i32(v750 != int32(1)) == v739 {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v522)+44))
	if v733 == int32(0) {
		goto L226
	} else {
		goto L228
	}
L228:
	;
	F_pfree(m, v733)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L8
	} else {
		goto L229
	}
L229:
	;
	goto L226
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v743 | int32(65)
	goto L219
L231:
	;
	goto L232
L232:
	;
	if v31 == int32(1) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v743 | int32(_a_F__bt_advance_array_keys_2)
	goto L219
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v743 | int32(_a_F__bt_advance_array_keys_1)
	goto L219
L236:
	;
	goto L153
L237:
	;
	v801 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v801)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+17)) = uint8(v801)
	v1022 = v480
	goto L1
L238:
	;
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
	v836 = v834 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v836)
	v851 = v818
	v858 = v825
	v860 = v124
	goto L11
L239:
	;
	if v851 != 0 {
		goto L259
	} else {
		goto L260
	}
L240:
	;
	v862 = v851
	goto L242
L241:
	;
	v862 = v858
	goto L242
L242:
	;
	if v862&int32(1) != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v865 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = l5 + v865
	v875 = F__bt_check_compare(m, l0, v31, l2, l3, l4, int32(0), l6^v865, v27+int32(15), v27+int32(8))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L8
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	if l6 != 0 {
		goto L239
	} else {
		goto L256
	}
L246:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)))
	if v885 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	if v875 == int32(0) {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+18)))
	if v879 != 0 {
		goto L246
	} else {
		goto L249
	}
L249:
	;
	v880 = int32(1)
	if l1 == int32(0) {
		v1022 = v880
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v883 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v883)
	v1022 = v880
	goto L1
L251:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v890 = F__bt_advance_array_keys(m, l0, l1, l2, l3, l4, v888, int32(1))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L8
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	if l6 != 0 {
		goto L239
	} else {
		goto L255
	}
L254:
	;
	v1022 = int32(0)
	goto L1
L255:
	;
	v1022 = int32(0)
	goto L1
L256:
	;
	v1022 = int32(0)
	goto L1
L257:
	;
	v996 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v996)
	v998 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+17)) = uint8(v998)
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+18)))
	if v1001 != v996 {
		v1022 = v998
		goto L1
	} else {
		goto L285
	}
L258:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v973 == int32(1) {
		goto L279
	} else {
		goto L280
	}
L259:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+18)))
	if v860&(v921^int32(-1)) == int32(0) {
		goto L257
	} else {
		goto L269
	}
L260:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if l2 == v895 {
		goto L258
	} else {
		goto L261
	}
L261:
	;
	if v895 == int32(0) {
		goto L259
	} else {
		goto L262
	}
L262:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+7)))
	if v899&int32(32) == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v914 = int32(0)
	v918 = F__bt_tuple_before_array_skeys(m, l0, v31, v895, l4, v913, v914, v914, v30+int32(18))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L8
	} else {
		goto L267
	}
L264:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v29)+192))
	v911 = int32(*(*int16)(unsafe.Add(mBase, uint32(v910)+8)))
	v913 = v911
	goto L263
L265:
	;
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v895)+4)))
	if v904&int32(_a_F__bt_advance_array_keys_5) != 0 {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v913 = v904 & int32(4095)
	goto L263
L267:
	;
	if v918 != 0 {
		goto L258
	} else {
		goto L268
	}
L268:
	;
	goto L259
L269:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v927 == int32(0) {
		goto L257
	} else {
		goto L270
	}
L270:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v930)+52))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927)+7)))
	if v932&int32(32) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v948 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v948
	v958 = F__bt_check_compare(m, l0, v948-v31, v927, v946, v931, v948, v948, v27+int32(15), v27+int32(8))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L8
	} else {
		goto L275
	}
L272:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v930)+192))
	v944 = int32(*(*int16)(unsafe.Add(mBase, uint32(v943)+8)))
	v946 = v944
	goto L271
L273:
	;
	v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v927)+4)))
	if v937&int32(_a_F__bt_advance_array_keys_5) != 0 {
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v946 = v937 & int32(4095)
	goto L271
L275:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+15)))
	if v960 != 0 {
		goto L257
	} else {
		goto L276
	}
L276:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v947)+8))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961+v962*int32(48))+6)))
	if v966 == int32(3) {
		goto L257
	} else {
		goto L277
	}
L277:
	;
	goto L258
L278:
	;
	v981 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v981)
	v984 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+17)) = uint8(v984)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v986 == v981 {
		v1022 = v981
		goto L1
	} else {
		goto L283
	}
L279:
	;
	v976 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+30)))
	if v976 < int32(4) {
		goto L278
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v979 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+18)) = uint8(v979)
	goto L257
L282:
	;
	goto L281
L283:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	F__bt_parallel_primscan_schedule(m, l0, v989)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L8
	} else {
		goto L284
	}
L284:
	;
	v1022 = v981
	goto L1
L285:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)) = uint8(v860)
	if v31 != int32(1) {
		v1022 = v998
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v1007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v1009 = v1007 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v1009)
	v1022 = v998
	goto L1
}
func F__bt_buildadd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F__bt_buildadd[0]))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l3
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = int32(4)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v32 = v30 - v31
	if v32 <= v29 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v45 = (v39&int32(_a_F__bt_buildadd_0) + int32(7)) & int32(_a_F__bt_buildadd_1)
	if base.Ui32(int32(2705)) <= base.Ui32(v45) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v35 = v29
	goto L9
L8:
	;
	v35 = v32
	goto L9
L9:
	;
	v37 = v35 - int32(4)
	goto L6
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F__bt_check_third_page(m, v48, v49, base.B2i32(v38 == int32(0)), v28, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v38 != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L68
	}
L15:
	;
	v315 = F_PageAddItemExtended(m, v305, v302, v310, v301&int32(_a_F__bt_buildadd_3), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L64
	}
L16:
	;
	v301 = v298
	v302 = v299
	v305 = v28
	v306 = v27
	v310 = v300
	goto L15
L17:
	;
	v298 = v269
	v299 = l2
	v300 = v45
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L61
	}
L19:
	;
	if v26 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L20:
	;
	v56 = int32(0)
	goto L22
L21:
	;
	v56 = int32(8)
	goto L22
L22:
	;
	if base.Ui32(v56+v45) <= base.Ui32(v37) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.B2i32(base.Ui32(v26) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v61) <= base.Ui32(v37+v24)) != 0 {
		goto L19
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v67 = F_smgr_bulk_get_buf(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	F_PageInit(m, v67, int32(_a_F__bt_buildadd_2), int32(16))
	mBase = m.M
	goto L28
L28:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v73 = v67 + v72
	v74 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+14)) = uint16(v74)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+12)) = uint16(base.B2i32(v65 == v74))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = int64(0)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	v84 = v82 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)) = uint16(v84)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v86 + int32(1)
	v91 = v28 + int32(20)
	v94 = v91 + v26<<(uint(int32(2))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v98 = v28 + v95&int32(_a_F__bt_buildadd_7)
	if v38 == v74 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v114 = F_PageAddItemExtended(m, v67, v111, v110, int32(2), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L33
	}
L30:
	;
	v110 = int32(base.Ui32(v95) >> (uint(int32(17)) % 32))
	v111 = v98
	goto L29
L31:
	;
	goto L32
L32:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v98)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(537395200)
	v107 = int32(8)
	v110 = v107
	v111 = v18 + v107
	goto L29
L33:
	;
	if v114 == int32(0) {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v118
	v120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v120
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	v124 = v122 - int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)) = uint16(v124)
	if v38 == v120 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v128 = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v91+(v26-v128)&int32(_a_F__bt_buildadd_3)<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v142 = F__bt_truncate(m, v129, v28+v137&int32(_a_F__bt_buildadd_7), v98, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v158 = v98
	goto L37
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v159 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
	v147 = F_PageIndexTupleOverwrite(m, v28, v128, v142, v144&int32(_a_F__bt_buildadd_0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v147 == int32(0) {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	F_pfree(m, v142)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v158 = v28 + v153&int32(_a_F__bt_buildadd_7)
	goto L37
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v164 = v162 + int32(1)
	v166 = F_palloc0(m, int32(32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = base.I32_rotr(v27, int32(16))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__bt_buildadd(m, l0, v228, v229, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L51
	}
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v169 = F_smgr_bulk_get_buf(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_PageInit(m, v169, int32(_a_F__bt_buildadd_2), int32(16))
	mBase = m.M
	goto L47
L47:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+16)))
	v175 = v169 + v174
	v176 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)) = uint16(v176)
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+12)) = uint16(base.B2i32(v164 == v176))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v164
	*(*int64)(unsafe.Add(mBase, uint32(v175))) = int64(0)
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+12)))
	v186 = v184 + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+12)) = uint16(v186)
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v169
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v190 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v189 + v190
	*(*int32)(unsafe.Add(mBase, uint32(v166)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v166)+16)) = v176
	*(*uint16)(unsafe.Add(mBase, uint32(v166)+12)) = uint16(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+8)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v166)+4)) = v189
	if v164 != 0 {
		v215 = int32(2457)
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+24)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v166
	goto L44
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+180))
	if v204 == int32(0) {
		v215 = int32(819)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v213 = base.I32_div_s(int32(_a_F__bt_buildadd_10)-v208<<(uint(int32(13))%32), int32(100))
	v215 = v213
	goto L48
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_pfree(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v236 = F_CopyIndexTuple(m, v158)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v236
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+v240)+4)) = v86
	v243 = v239 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v27
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_smgr_bulk_write(m, v247, v27, v28, int32(1))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v301 = int32(3)
	v302 = l2
	v305 = v67
	v306 = v86
	v310 = v45
	goto L15
L55:
	;
	v255 = F_palloc0(m, int32(8))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v269 = v26 + int32(1)
	if v38 == int32(0) {
		goto L17
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v255
	v258 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v255)+6)) = uint16(v258)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v261 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v260)+4)) = uint16(v261)
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+6)))
	v265 = v263 | int32(_a_F__bt_buildadd_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v260)+6)) = uint16(v265)
	goto L57
L59:
	;
	v272 = int32(2)
	if v269&int32(_a_F__bt_buildadd_3) != v272 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	v277 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(537395200)
	v281 = int32(8)
	v298 = v272
	v299 = v18 + v281
	v300 = v281
	goto L16
L61:
	;
	F_errmsg_internal(m, int32(_a_F__bt_buildadd_8), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F__bt_buildadd_5), int32(938), int32(_a_F__bt_buildadd_9))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
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
	if v315 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v301)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v305
	m.G0 = v18 + int32(16)
	return
L66:
	;
	goto L67
L67:
	;
	goto L14
L68:
	;
	F_errmsg_internal(m, int32(_a_F__bt_buildadd_4), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F__bt_buildadd_5), int32(735), int32(_a_F__bt_buildadd_6))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_checkpage(m *base.Module, l0 int32, l1 int32) {
	var v8 int32
	_ = v8
	Fn13931(m, l0, l1, int32(_a_F__bt_checkpage_0), int32(824), int32(_a_F__bt_checkpage_1), int32(813))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F__bt_end_vacuum_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	v9 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_vacuum_callback[0]))
	v13 = F_LWLockAcquire(m, v9+int32(2560), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_vacuum_callback[1]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_vacuum_callback[0]))
	F_LWLockRelease(m, v60+int32(2560))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v24 = int32(0)
	goto L5
L5:
	;
	v33 = v16 + int32(12) + v24*int32(12)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 != v22 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v50 = v24 + int32(1)
	if v50 != v17 {
		v24 = v50
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v36 != v37 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v41 = v16 + v17*int32(12)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v42
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v17 - int32(1)
	goto L3
L10:
	;
	goto L6
L11:
	;
	return
}
func F__bt_find_extreme_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+208))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(2))%32)-int32(4))))
	v23 = F_get_opfamily_member(m, v21, l2, l2, base.I32_extend16_s(l3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L23
	}
L2:
	;
	return int32(0)
L3:
	;
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = F_get_opcode(m, v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L20
	}
L7:
	;
	if v27 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_fmgr_info(m, v27, v11+int32(20))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if int32(2) <= l5 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = v35
	v41 = int32(1)
	goto L13
L11:
	;
	v61 = v35
	goto L12
L12:
	;
	m.G0 = v11 + int32(48)
	return v61
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v52 = l4 + v41<<(uint(int32(2))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = F_FunctionCall2Coll(m, v11+int32(20), v49, v53, v39)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v61 = v57
	goto L12
L15:
	;
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v57 = v56
	goto L18
L17:
	;
	v57 = v39
	goto L18
L18:
	;
	v59 = v41 + int32(1)
	if v59 != l5 {
		v39 = v57
		v41 = v59
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+208))
	v78 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77+v78<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
	F_errmsg_internal(m, int32(_a_F__bt_find_extreme_element_0), v11)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F__bt_find_extreme_element_1), int32(2605), int32(_a_F__bt_find_extreme_element_2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v23
	F_errmsg_internal(m, int32(_a_F__bt_find_extreme_element_3), v11+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F__bt_find_extreme_element_1), int32(2608), int32(_a_F__bt_find_extreme_element_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_form_posting(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v8&int32(_a_F__bt_form_posting_0) == int32(0) {
		v25 = v8 & int32(_a_F__bt_form_posting_1)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		if v13&int32(32) == int32(0) {
			v25 = v8 & int32(_a_F__bt_form_posting_1)
		} else {
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
			v25 = v18 | v19<<(uint(int32(16))%32)
		}
	}
	v27 = l2 * int32(6)
	if int32(1) < l2 {
		v35 = (v25 + v27 + int32(7)) & int32(-8)
	} else {
		v35 = v25
	}
	v36 = F_palloc0(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		if v25 != 0 {
			base.MemoryCopy(m, v36, l0, v25)
		} else {
		}
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
		v44 = v41&int32(-8192) | v35
		if int32(2) <= l2 {
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)) = uint16(v25)
			v48 = int32(_a_F__bt_form_posting_0)
			v49 = l2 | v48
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v49)
			v52 = v44 | v48
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)) = uint16(v52)
			v55 = int32(base.Ui32(v25) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(v36))) = uint16(v55)
			if v27 == int32(0) {
				return v36
			} else {
				base.MemoryCopy(m, v36+v25&int32(-65536)+v25&int32(_a_F__bt_form_posting_2), l1, v27)
				return v36
			}
		} else {
			v68 = v44 & int32(_a_F__bt_form_posting_3)
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)) = uint16(v68)
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v36))) = v70
			v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v72)
			return v36
		}
	}
}
func F__bt_getmeta(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getmeta[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l1^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getmeta[1]))
		v26 = v20 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v26)+12)))
	if v29&int32(8) == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58 + int32(4)
				F_errmsg(m, int32(_a_F__bt_getmeta_0), v7)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F__bt_getmeta_1), int32(158), int32(_a_F__bt_getmeta_2))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
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
		v35 = v26 + int32(24)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
		if v36 != int32(_a_F__bt_getmeta_3) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58 + int32(4)
					F_errmsg(m, int32(_a_F__bt_getmeta_0), v7)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F__bt_getmeta_1), int32(158), int32(_a_F__bt_getmeta_2))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
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
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
			if base.Ui32(v39-int32(5)) <= base.Ui32(int32(-4)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(8589934596)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v77 + int32(4)
						F_errmsg(m, int32(_a_F__bt_getmeta_4), v7+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F__bt_getmeta_1), int32(167), int32(_a_F__bt_getmeta_2))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
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
				m.G0 = v7 + int32(32)
				return v35
			}
		}
	}
}
func F__bt_getstackbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v218 int32
	_ = v218
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v22 = v15
	v26 = v16
	goto L1
L1:
	;
	goto L3
L2:
	;
	return int32(0)
L3:
	;
	v46 = F__bt_getbuf(m, l0, v26, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v70&int32(20) != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v69 = v68 + v67
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+12)))
	if v70&int32(128) != 0 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	if v46 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getstackbuf[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(v46^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L5
L9:
	;
	goto L10
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getstackbuf[1]))
	v67 = v61 + v46<<(uint(int32(13))%32) + int32(-8192)
	goto L5
L11:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F__bt_finish_split(m, l0, l1, v46, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L4
L14:
	;
	goto L3
L15:
	;
	F__bt_relbuf(m, v46)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L43
	}
L16:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v79) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v191)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
	return v46
L18:
	;
	v87 = int32(base.Ui32(v79+int32(_a_F__bt_getstackbuf_0)) >> (uint(int32(2)) % 32))
	goto L20
L19:
	;
	v87 = int32(0)
	goto L20
L20:
	;
	v88 = int32(1)
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v92 = int32(2)
	goto L23
L22:
	;
	v92 = v88
	goto L23
L23:
	;
	if base.Ui32(v92) < base.Ui32(v22) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = v22
	goto L26
L25:
	;
	v94 = v92
	goto L26
L26:
	;
	v96 = v87 & int32(_a_F__bt_getstackbuf_1)
	if base.Ui32(v96) < base.Ui32(v94) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = v87 + v88
	goto L29
L28:
	;
	v98 = v94
	goto L29
L29:
	;
	if base.Ui32(v98&int32(_a_F__bt_getstackbuf_1)) <= base.Ui32(v96) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v108 = v98
	goto L33
L31:
	;
	goto L32
L32:
	;
	v159 = v98
	goto L39
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(20)+v108&int32(_a_F__bt_getstackbuf_1)<<(uint(int32(2))%32))))
	v126 = v67 + v123&int32(_a_F__bt_getstackbuf_2)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126))))
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+2)))
	if l3 == v127<<(uint(int32(16))%32)|v130 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v191 = v108
	goto L17
L36:
	;
	goto L37
L37:
	;
	v134 = v108 + int32(1)
	if base.Ui32(v134&int32(_a_F__bt_getstackbuf_1)) <= base.Ui32(v96) {
		v108 = v134
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v169 = v159 - int32(1)
	v171 = v169 & int32(_a_F__bt_getstackbuf_1)
	if base.Ui32(v171) < base.Ui32(v92) {
		goto L15
	} else {
		goto L41
	}
L40:
	;
	v191 = v169
	goto L17
L41:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v67+int32(20)+v171<<(uint(int32(2))%32))))
	v179 = v67 + v176&int32(_a_F__bt_getstackbuf_2)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179))))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179)+2)))
	if v180<<(uint(int32(16))%32)|v183 != l3 {
		v159 = v169
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if v76 != 0 {
		v22 = int32(0)
		v26 = v76
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L2
}
func F__bt_next(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
	if l1 == int32(1) {
		v10 = v6 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+100)) = v10
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
		if v10 <= v12 {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
			v35 = v5 + v32*int32(10)
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+108)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v36)
			v39 = v35 + int32(104)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v40
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			if v42 != 0 {
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42 + v43
			} else {
			}
			return int32(1)
		} else {
			v15 = F__bt_steppage(m, l0, int32(1))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v15 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
					v35 = v5 + v32*int32(10)
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+108)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v36)
					v39 = v35 + int32(104)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
					if v42 != 0 {
						v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42 + v43
					} else {
					}
					return int32(1)
				} else {
					return int32(0)
				}
			}
		}
	} else {
		v22 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+100)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
		if v24 <= v22 {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
			v35 = v5 + v32*int32(10)
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+108)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v36)
			v39 = v35 + int32(104)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v40
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
			if v42 != 0 {
				v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42 + v43
			} else {
			}
			return int32(1)
		} else {
			v26 = F__bt_steppage(m, l0, l1)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+100))
					v35 = v5 + v32*int32(10)
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+108)))
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v36)
					v39 = v35 + int32(104)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
					if v42 != 0 {
						v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v42 + v43
					} else {
					}
					return int32(1)
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F__bt_readfirstpage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+52)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+17)))
	if v10 == int32(1) {
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+17)) = uint8(v13)
		v15 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+88)) = uint16(v15)
	} else {
		if l2 == int32(1) {
			v19 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v5)+88)) = uint16(v19)
		} else {
			v21 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v5)+88)) = uint16(v21)
		}
	}
	v24 = F__bt_readpage(m, l0, l2, l1, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		if v24 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+40)))
			if v30 == int32(0) {
				F__bt_unlockbuf(m, v29)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					return int32(1)
				}
			} else {
				v37 = F_BufferGetLSNAtomic(m, v29)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v5)+72)) = v37
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
					F__bt_relbuf(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+56)) = int32(0)
						return int32(1)
					}
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v5)+56))
			F__bt_unlockbuf(m, v47)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v50 = F__bt_steppage(m, l0, l2)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					return v50
				}
			}
		}
	}
}
func F__bt_set_startikey(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = v25 + int32(20)
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v29 = int32(2)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(v29)%32))))
	v33 = int32(_a_F__bt_set_startikey_0)
	v35 = v25 + v32&v33
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27+v36<<(uint(v29)%32))))
	v43 = v25 + v40&v33
	v44 = int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
	if v46 <= int32(0) {
		v93 = v44
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v107 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v53 = v44
	goto L5
L5:
	;
	v69 = F_index_getattr_2(m, v43, v53, v24, v17+int32(15))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v93 = v46 + int32(1)
	goto L3
L7:
	;
	return
L8:
	;
	v73 = F_index_getattr_2(m, v35, v53, v24, v17+int32(14))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+14)))
	if v75 != v76 {
		v93 = v53
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v75 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v82 = v24 + int32(4) + v53<<(uint(int32(4))%32)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+6)))
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+4)))
	v85 = F_datum_image_eq(m, v69, v73, v83, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v53 != v46 {
		v53 = v53 + int32(1)
		goto L5
	} else {
		goto L16
	}
L14:
	;
	if v85 == int32(0) {
		v93 = v53
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L6
L17:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v364 != 0 {
		goto L1
	} else {
		goto L94
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v341
	v344 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v344)
	goto L17
L19:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v327
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v332)
	if v332 != int32(1) {
		goto L1
	} else {
		goto L93
	}
L20:
	;
	v327 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v111 = int32(0)
	v116 = v111
	v119 = v111
	v123 = v111
	goto L23
L23:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v131 = v128 + v123*int32(48)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if base.B2i32(v132&int32(_a_F__bt_set_startikey_1) == int32(0))|v132&int32(4) != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	if v307&int32(1) != 0 {
		v341 = v313
		goto L18
	} else {
		goto L92
	}
L25:
	;
	v313 = v123 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v313 < v314 {
		v116 = v307
		v119 = v309
		v123 = v313
		goto L23
	} else {
		goto L91
	}
L26:
	;
	v307 = v116
	v309 = v211
	goto L25
L27:
	;
	if v116&int32(1) == int32(0) {
		v327 = v123
		goto L19
	} else {
		goto L90
	}
L28:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+6)))
	if v140 != int32(3) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	if v93 < v143 {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v132&int32(32) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L32:
	;
	v147 = F_index_getattr_2(m, v43, v143, v24, v17+int32(13))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	v152 = F_index_getattr_2(m, v35, v149, v24, v17+int32(12))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v155&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v154&int32(1) != 0 {
		goto L27
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v154&int32(1) != 0 {
		goto L27
	} else {
		goto L40
	}
L38:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v160&int32(1) != 0 {
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v307 = v116
	v309 = v119
	goto L25
L40:
	;
	v166 = v131 + int32(16)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
	v169 = F_FunctionCall2Coll(m, v166, v167, v147, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	if v169 == int32(0) {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v173&int32(1) != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
	v178 = F_FunctionCall2Coll(m, v166, v176, v152, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	if v178 == int32(0) {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v307 = v116
	v309 = v119
	goto L25
L46:
	;
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	if v93 <= v186 {
		goto L27
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v211 = v119 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v215 = v212 + v119<<(uint(int32(5))%32)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v216 != int32(-1) {
		goto L58
	} else {
		goto L59
	}
L49:
	;
	v190 = F_index_getattr_2(m, v43, v186, v24, v17+int32(13))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v193&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v192&int32(1) == int32(0) {
		goto L27
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v192&int32(1) != 0 {
		goto L27
	} else {
		goto L55
	}
L54:
	;
	v307 = v116
	v309 = v119
	goto L25
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v131)+44))
	v206 = F_FunctionCall2Coll(m, v131+int32(16), v204, v190, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if v206 == int32(0) {
		goto L27
	} else {
		goto L57
	}
L57:
	;
	v307 = v116
	v309 = v119
	goto L25
L58:
	;
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	if v93 <= v219 {
		goto L27
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+19)))
	if v238 != 0 {
		goto L26
	} else {
		goto L65
	}
L61:
	;
	v223 = F_index_getattr_2(m, v43, v219, v24, v17+int32(13))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v229 = int32(0)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	v234 = F__bt_binsrch_array_skey(m, v225+v123*int32(28), v229, v229, v223, v231, v215, v131, v17+int32(8))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v237 != 0 {
		goto L27
	} else {
		goto L64
	}
L64:
	;
	v307 = int32(1)
	v309 = v211
	goto L25
L65:
	;
	v239 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	if v93 < v239 {
		goto L27
	} else {
		goto L66
	}
L66:
	;
	v243 = F_index_getattr_2(m, v43, v239, v24, v17+int32(13))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+4)))
	v248 = F_index_getattr_2(m, v35, v245, v24, v17+int32(12))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+19)))
	if v250 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+19)))
	if v273 != 0 {
		goto L26
	} else {
		goto L80
	}
L70:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+13)))
	if v251 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	if v252 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	v257 = F_FunctionCall2Coll(m, v252+int32(16), v255, v243, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	if v261 == int32(0) {
		goto L69
	} else {
		goto L77
	}
L75:
	;
	if v257 == int32(0) {
		goto L27
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v261)+44))
	v268 = F_FunctionCall2Coll(m, v261+int32(16), v266, v243, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	if v268 == int32(0) {
		goto L27
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v274 != 0 {
		goto L27
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	if v277 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)+44))
	v282 = F_FunctionCall2Coll(m, v277+int32(16), v280, v248, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	if v286 == int32(0) {
		goto L26
	} else {
		goto L87
	}
L85:
	;
	if v282 == int32(0) {
		goto L27
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286)+44))
	v293 = F_FunctionCall2Coll(m, v286+int32(16), v291, v248, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	if v293 != 0 {
		goto L26
	} else {
		goto L89
	}
L89:
	;
	goto L27
L90:
	;
	v341 = v123
	goto L18
L91:
	;
	goto L24
L92:
	;
	v327 = v313
	goto L19
L93:
	;
	goto L17
L94:
	;
	v365 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(13)))) = uint8(v365)
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)))) = v365
	goto L1
}
func F__bt_start_array_keys(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v3 < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	v94 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+18)) = uint16(v94)
	return
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v23 = v20 + v17<<(uint(int32(5))%32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v27 = v19 + v24*int32(48)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v28 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v84 = v17 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v84 < v85 {
		v17 = v84
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v31 = int32(1)
	if l1 != v31 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+18)))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v36 = v28 - v31
	goto L12
L11:
	;
	v36 = int32(0)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v42
	goto L6
L13:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v55 = v53 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55
	v61 = int32(1)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+19)))
	if base.B2i32(base.B2i32(v53&int32(33554432) == v51) == base.B2i32(l1 == v61))|base.B2i32(v64 != v61) == v51 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v45 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55 | int32(65)
	goto L6
L19:
	;
	goto L20
L20:
	;
	if l1 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55 | int32(_a_F__bt_start_array_keys_0)
	goto L6
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v55 | int32(_a_F__bt_start_array_keys_1)
	goto L6
L24:
	;
	goto L5
}
func F_bt_page_items(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_items_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_bt_report_duplicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v17 | v18<<(uint(int32(16))%32)
	v28 = F_psprintf(m, int32(_a_F_bt_report_duplicate_0), v14+int32(128))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
		v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v30 | v31<<(uint(int32(16))%32)
		v41 = F_psprintf(m, int32(_a_F_bt_report_duplicate_0), v14+int32(112))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v43
			v50 = F_psprintf(m, int32(_a_F_bt_report_duplicate_0), v14+int32(96))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if l3 != v52 {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l3
					v62 = F_psprintf(m, int32(_a_F_bt_report_duplicate_1), v14+int32(80))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						v64 = v62
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v65 < int32(0) {
							v75 = int32(_a_F_bt_report_duplicate_2)
							if l5 < int32(0) {
								v85 = int32(_a_F_bt_report_duplicate_2)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errcode(m, int32(33557032))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
										F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
											v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
											*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
											*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
											*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
											F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
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
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l5
								v83 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14+int32(48))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									v85 = v83
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
											F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
												v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
												*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
												*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
												F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
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
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v65
							v73 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14-int32(-64))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								v75 = v73
								if l5 < int32(0) {
									v85 = int32(_a_F_bt_report_duplicate_2)
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
											F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
												v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
												*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
												*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
												F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
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
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l5
									v83 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14+int32(48))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v85 = v83
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
												F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
													v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
													*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
													*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
													F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
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
				} else {
					v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					if l4 != v54 {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l3
						v62 = F_psprintf(m, int32(_a_F_bt_report_duplicate_1), v14+int32(80))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							v64 = v62
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v65 < int32(0) {
								v75 = int32(_a_F_bt_report_duplicate_2)
								if l5 < int32(0) {
									v85 = int32(_a_F_bt_report_duplicate_2)
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
											F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
												v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
												*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
												*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
												F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
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
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l5
									v83 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14+int32(48))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v85 = v83
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
												F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
													v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
													*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
													*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
													F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
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
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v65
								v73 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14-int32(-64))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v75 = v73
									if l5 < int32(0) {
										v85 = int32(_a_F_bt_report_duplicate_2)
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
												F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
													v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
													*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
													*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
													F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
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
										*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l5
										v83 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14+int32(48))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											v85 = v83
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return
											} else {
												F_errcode(m, int32(33557032))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
													F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return
													} else {
														v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
														*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
														v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
														*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
														*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
														*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
														*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
														F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return
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
					} else {
						v64 = int32(_a_F_bt_report_duplicate_2)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						if v65 < int32(0) {
							v75 = int32(_a_F_bt_report_duplicate_2)
							if l5 < int32(0) {
								v85 = int32(_a_F_bt_report_duplicate_2)
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errcode(m, int32(33557032))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
										F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
											v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
											*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
											*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
											*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
											*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
											F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
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
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l5
								v83 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14+int32(48))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									v85 = v83
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
											F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
												v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
												*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
												*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
												F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
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
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v65
							v73 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14-int32(-64))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								v75 = v73
								if l5 < int32(0) {
									v85 = int32(_a_F_bt_report_duplicate_2)
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errcode(m, int32(33557032))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
											F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
												v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
												*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
												*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
												*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
												F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
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
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l5
									v83 = F_psprintf(m, int32(_a_F_bt_report_duplicate_7), v14+int32(48))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v85 = v83
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_errcode(m, int32(33557032))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94 + int32(4)
												F_errmsg(m, int32(_a_F_bt_report_duplicate_3), v14+int32(32))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v103 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+28)) = uint32(v103)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v28
													v108 = int64(base.Ui64(v103) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v14)+24)) = uint32(v108)
													*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v85
													*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v64
													*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v75
													*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
													F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(906), int32(_a_F_bt_report_duplicate_6))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
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
}
