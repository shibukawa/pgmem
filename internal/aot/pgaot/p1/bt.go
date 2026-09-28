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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
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
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int64
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v417 int64
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v434 int64
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int64
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int64
	_ = v457
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int64
	_ = v571
	var v578 int32
	_ = v578
	var v579 int64
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int64
	_ = v582
	var v583 int32
	_ = v583
	var v585 int64
	_ = v585
	var v588 int32
	_ = v588
	var v589 int64
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int64
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v623 int64
	_ = v623
	var v624 int64
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int64
	_ = v633
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v651 int64
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v675 int32
	_ = v675
	var v676 int64
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int64
	_ = v679
	var v680 int32
	_ = v680
	var v682 int64
	_ = v682
	var v685 int32
	_ = v685
	var v686 int64
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int64
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int64
	_ = v722
	var v723 int64
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int64
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int64
	_ = v756
	var v758 int32
	_ = v758
	var v759 int64
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1056 int32
	_ = v1056
	v8 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(16)
	m.G0 = v29
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v33 = v31
	goto L3
L2:
	;
	v33 = int32(1)
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l6 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v29 + int32(16)
	return v1056
L5:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v63 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+30)) = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v38 - int32(1)
	if v40 <= l5 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v40*int32(56)))))
	if v46&int32(32) != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v40
	v50 = int32(0)
	v56 = F__bt_check_compare(m, l0, v33, l2, l3, l4, v50, v50, v29+int32(15), v29+int32(8))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v56 == int32(0) {
		v1056 = v8
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	if l6 != 0 {
		goto L243
	} else {
		goto L244
	}
L15:
	;
	v66 = int32(1)
	v883 = v66
	v890 = v66
	v892 = v8
	goto L14
L16:
	;
	goto L17
L17:
	;
	v69 = int32(1)
	v81 = int32(0)
	v82 = v8
	v84 = v69
	v87 = v8
	v91 = v69
	v92 = v8
	v93 = v8
	goto L19
L18:
	;
	v861 = int32(0)
	if base.B2i32(l6 == v861)|base.B2i32(v851 == v861) != 0 {
		v883 = v848
		v890 = v855
		v892 = v130
		goto L14
	} else {
		goto L241
	}
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v100 = v97 + v81*int32(56)
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
	if v101 == int32(3) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v477&int32(1) == int32(0) {
		v848 = v478
		v851 = v479
		v855 = v482
		goto L18
	} else {
		goto L151
	}
L21:
	;
	if v81 < l5 {
		v477 = v82
		v478 = v84
		v479 = v87
		v482 = v91
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v104 = int32(0)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v105&int32(32) == v104 {
		v128 = v104
		v129 = v92
		v130 = v93
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v116 = int32(0)
	switch v33 + int32(1) {
	case 0:
		goto L27
	default:
		v128 = v116
		v129 = v92
		v130 = v93
		goto L21
	case 2:
		goto L28
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v128 = v110 + v92<<(uint(int32(5))%32)
	v129 = v92 + int32(1)
	v130 = v93
	goto L21
L26:
	;
	v128 = v116
	v129 = v92
	v130 = int32(1)
	goto L21
L27:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	if v122&int32(1) == int32(0) {
		v128 = v116
		v129 = v92
		v130 = v93
		goto L21
	} else {
		goto L30
	}
L28:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	if v119&int32(2) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v128 = v116
	v129 = v92
	v130 = v93
	goto L21
L30:
	;
	goto L26
L31:
	;
	v486 = v81 + int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v486 < v487 {
		v81 = v486
		v82 = v477
		v84 = v478
		v87 = v479
		v91 = v482
		v92 = v129
		v93 = v130
		goto L19
	} else {
		goto L150
	}
L32:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v134 = v132 & int32(_a_F__bt_advance_array_keys_0)
	if v134 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if base.B2i32(l5 != v81)|v128 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+4)))
	if v137 <= l3 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)) = uint8(v139)
	goto L33
L36:
	;
	v477 = v369
	v478 = int32(0)
	v479 = int32(1)
	v482 = v371
	goto L31
L37:
	;
	v472 = int32(1)
	v477 = v369
	v478 = v472
	v479 = v472
	v482 = v371
	goto L31
L38:
	;
	v848 = int32(1)
	v851 = v87
	v855 = int32(0)
	goto L18
L39:
	;
	v477 = int32(1)
	v478 = v464
	v479 = v87
	v482 = v465
	goto L31
L40:
	;
	v145 = int32(0)
	v464 = v145
	v465 = v145
	goto L39
L41:
	;
	goto L42
L42:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+6)))
	if base.B2i32(v128|v134 == int32(0))|base.B2i32(v150 != int32(3)) != 0 {
		v477 = v82
		v478 = v84
		v479 = v87
		v482 = v91
		goto L31
	} else {
		goto L43
	}
L43:
	;
	if v82&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v186 | int32(_a_F__bt_advance_array_keys_1)
	v464 = v84
	v465 = v91
	goto L39
L45:
	;
	if v128 == int32(0) {
		v464 = v84
		v465 = v91
		goto L39
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v84 != 0 {
		goto L64
	} else {
		goto L65
	}
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v158 != int32(-1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v33 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+18)))
	if v174 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v166 = v158 - int32(1)
	goto L54
L53:
	;
	v166 = int32(0)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v168+v166<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = v172
	v464 = v84
	v465 = v91
	goto L39
L55:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = int64(0)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v186 = v184 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v186
	v190 = int32(0)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+19)))
	if base.B2i32(base.B2i32(v184&int32(33554432) == v190) == base.B2i32(v33 == int32(-1)))|base.B2i32(v195 != int32(1)) == v190 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	if v175 == int64(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, base.I32_wrap_i64(v175))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	goto L55
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v186 | int32(65)
	v464 = v84
	v465 = v91
	goto L39
L60:
	;
	goto L61
L61:
	;
	if v33 != int32(-1) {
		goto L44
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v186 | int32(_a_F__bt_advance_array_keys_2)
	v464 = v84
	v465 = v91
	goto L39
L63:
	;
	v271 = F_index_getattr_2(m, l2, v209, l4, v29+int32(15))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L85
	}
L64:
	;
	v209 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+4)))
	if v209 <= l3 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v212 = int32(0)
	if v128 == v212 {
		v477 = v212
		v478 = v84
		v479 = v87
		v482 = v91
		goto L31
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v215 != int32(-1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v218 = int32(1)
	if v33 != v218 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+18)))
	if v231 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v223 = v215 - v218
	goto L74
L73:
	;
	v223 = int32(0)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v225+v223<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = v229
	v477 = v212
	v478 = v84
	v479 = v87
	v482 = v91
	goto L31
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = int64(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v243 = v241 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v243
	v247 = int32(0)
	v249 = int32(1)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+19)))
	if base.B2i32(base.B2i32(v241&int32(33554432) == v247) == base.B2i32(v33 == v249))|base.B2i32(v252 != v249) == v247 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	if v232 == int64(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	F_pfree(m, base.I32_wrap_i64(v232))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v243 | int32(65)
	v477 = v212
	v478 = v84
	v479 = v87
	v482 = v91
	goto L31
L80:
	;
	goto L81
L81:
	;
	if v33 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v243 | int32(_a_F__bt_advance_array_keys_2)
	v477 = v212
	v478 = v84
	v479 = v87
	v482 = v91
	goto L31
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v243 | int32(_a_F__bt_advance_array_keys_1)
	v477 = v212
	v478 = v84
	v479 = v87
	v482 = v91
	goto L31
L85:
	;
	if v128 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v342 = int32(0)
	if base.B2i32(l6 == v342)|base.B2i32(v134 == v342) == v342 {
		goto L116
	} else {
		goto L117
	}
L87:
	;
	v338 = v333
	v340 = int32(0)
	goto L86
L88:
	;
	v274 = l6 & base.B2i32(l5 == v81)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v275 == int32(-1) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v293 = int32(1)
	v294 = v292 & v293
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	if v295 == v293 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	F__bt_binsrch_skiparray_skey(m, v274, v33, v271, v278, v128, v100, v29+int32(8))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L11
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	v290 = F__bt_binsrch_array_skey(m, v283+v81*int32(28), v274, v33, v271, v287, v128, v100, v29+int32(8))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L95
	}
L94:
	;
	v333 = v278
	goto L87
L95:
	;
	v338 = v287
	v340 = v290
	goto L86
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v329
	v333 = v295
	goto L87
L97:
	;
	if v294 != 0 {
		v329 = int32(0)
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	if v294 != 0 {
		goto L104
	} else {
		goto L105
	}
L100:
	;
	if v292&int32(33554432) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v303 = int32(-1)
	goto L103
L102:
	;
	v303 = int32(1)
	goto L103
L103:
	;
	v329 = v303
	goto L96
L104:
	;
	if v292&int32(33554432) != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	v315 = F_FunctionCall2Coll(m, v309+v81*int32(28), v313, v271, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L110
	}
L107:
	;
	v308 = int32(1)
	goto L109
L108:
	;
	v308 = int32(-1)
	goto L109
L109:
	;
	v329 = v308
	goto L96
L110:
	;
	v317 = base.I32_wrap_i64(v315)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
	if v318&int32(1) == int32(0) {
		v329 = v317
		goto L96
	} else {
		goto L111
	}
L111:
	;
	v324 = int32(0)
	if v317 < v324 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v328 = int32(1)
	goto L114
L113:
	;
	v328 = v324 - v317
	goto L114
L114:
	;
	v329 = v328
	goto L96
L115:
	;
	if v128 == int32(0) {
		v477 = v369
		v478 = v370
		v479 = v87
		v482 = v371
		goto L31
	} else {
		goto L120
	}
L116:
	;
	v351 = int32(0)
	v363 = base.B2i32(v341 == v351)
	v368 = int32(base.Ui32(v341) >> (uint(int32(31)) % 32))
	v369 = base.B2i32(v33 == int32(1))&base.B2i32(v351 < v341) | base.B2i32(v33 == int32(-1))&base.B2i32(v341 < v351)
	v370 = v363
	v371 = v363 & v91
	goto L115
L117:
	;
	goto L118
L118:
	;
	if v341 != 0 {
		goto L38
	} else {
		goto L119
	}
L119:
	;
	v365 = int32(0)
	v368 = v365
	v369 = v365
	v370 = int32(1)
	v371 = v91
	goto L115
L120:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v374 == int32(-1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+18)))
	if v370 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v450 == v340 {
		v477 = v369
		v478 = v370
		v479 = v87
		v482 = v371
		goto L31
	} else {
		goto L149
	}
L124:
	;
	if v377&int32(1) != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	if v338 != 0 {
		goto L137
	} else {
		goto L138
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = int64(0)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v393 = v391 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v393
	v397 = int32(0)
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+19)))
	if base.B2i32(v368 == base.B2i32(v391&int32(33554432) == v397))|base.B2i32(v400 != int32(1)) == v397 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	if v382 == int64(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	F_pfree(m, base.I32_wrap_i64(v382))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v393 | int32(65)
	goto L36
L132:
	;
	goto L133
L133:
	;
	if v368 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v393 | int32(_a_F__bt_advance_array_keys_2)
	goto L36
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v393 | int32(_a_F__bt_advance_array_keys_1)
	goto L36
L137:
	;
	if v377&int32(1) != 0 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	if v377&int32(1) != 0 {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = int64(0)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v426&int32(-7864386) | int32(65)
	goto L37
L141:
	;
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	if v417 == int64(0) {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	F_pfree(m, base.I32_wrap_i64(v417))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L11
	} else {
		goto L143
	}
L143:
	;
	goto L140
L144:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v441 & int32(-7864386)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+18)))
	v446 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128)+16)))
	v447 = F_datumCopy(m, v271, v445, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L11
	} else {
		goto L148
	}
L145:
	;
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v100)+48))
	if v434 == int64(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	F_pfree(m, base.I32_wrap_i64(v434))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = v447
	goto L37
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v340
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v453+v340<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v100)+48)) = v457
	v477 = v369
	v478 = v370
	v479 = v87
	v482 = v371
	goto L31
L150:
	;
	goto L20
L151:
	;
	v493 = int32(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v497 = v495 - int32(1)
	if v493 <= v497 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v502 = base.B2i32(v33 == int32(1))
	v513 = v497
	v519 = v479
	goto L155
L153:
	;
	goto L154
L154:
	;
	F__bt_start_array_keys(m, l0, int32(0)-v33)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L11
	} else {
		goto L240
	}
L155:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	v533 = v530 + v513<<(uint(int32(5))%32)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v537 = v529 + v534*int32(56)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v538 == int32(-1) {
		goto L161
	} else {
		goto L162
	}
L156:
	;
	goto L154
L157:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v533)+4))
	if v744 != int32(-1) {
		goto L223
	} else {
		goto L224
	}
L158:
	;
	v653 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)) = uint8(v653)
	v655 = int32(1)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v659 = int32(33554433)
	if v656&int32(_a_F__bt_advance_array_keys_2)|base.B2i32(v656&v659 == v659) != 0 {
		v741 = v655
		goto L157
	} else {
		goto L194
	}
L159:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	if v641 <= int32(0) {
		v741 = v519
		goto L157
	} else {
		goto L193
	}
L160:
	;
	if v546&int32(1) != 0 {
		goto L169
	} else {
		goto L170
	}
L161:
	;
	if v502 == int32(0) {
		goto L158
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v33 != int32(1) {
		goto L159
	} else {
		goto L167
	}
L164:
	;
	v543 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)) = uint8(v543)
	v545 = int32(1)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v546&int32(_a_F__bt_advance_array_keys_1)|base.B2i32(v546&int32(33554433) == v545) != 0 {
		v741 = v545
		goto L157
	} else {
		goto L165
	}
L165:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	if v554 != 0 {
		goto L160
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v546 | int32(_a_F__bt_advance_array_keys_3)
	v848 = v478
	v851 = v545
	v855 = v482
	goto L18
L167:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v533)+12))
	if v538-int32(1) <= v560 {
		v741 = v519
		goto L157
	} else {
		goto L168
	}
L168:
	;
	v565 = v560 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v533)+12)) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	v571 = *(*int64)(unsafe.Add(mBase, uint32(v567+v565<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v571
	v848 = v478
	v851 = v519
	v855 = v482
	goto L18
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v546 & int32(-1048642)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	v579 = *(*int64)(unsafe.Add(mBase, uint32(v578)))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	v581 = int32(*(*int16)(unsafe.Add(mBase, uint32(v533)+16)))
	v582 = F_datumCopy(m, v579, v580, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L11
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v554)+20))
	v589 = m.T0[v588].(func(*base.Module, int32, int64, int32) int64)(m, v500, v585, v29+int32(8))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L11
	} else {
		goto L173
	}
L172:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v582
	v848 = v478
	v851 = v545
	v855 = v482
	goto L18
L173:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
	if v591 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+19)))
	if v594 != int32(1) {
		v741 = v545
		goto L157
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v533)+28))
	if v617 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L177:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v597&int32(33554432) != 0 {
		v741 = v545
		goto L157
	} else {
		goto L178
	}
L178:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v600 != 0 {
		v608 = v597
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v608&int32(-7864386) | int32(65)
	v848 = v478
	v851 = v545
	v855 = v482
	goto L18
L180:
	;
	v601 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	if v601 == int64(0) {
		v608 = v597
		goto L179
	} else {
		goto L181
	}
L181:
	;
	F_pfree(m, base.I32_wrap_i64(v601))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L11
	} else {
		goto L182
	}
L182:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v608 = v607
	goto L179
L183:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v632 != 0 {
		goto L189
	} else {
		goto L190
	}
L184:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v617)+12))
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v617)+48))
	v624 = F_FunctionCall2Coll(m, v617+int32(16), v622, v589, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	if v624 != int64(0) {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v628 != 0 {
		v741 = v545
		goto L157
	} else {
		goto L187
	}
L187:
	;
	F_pfree(m, base.I32_wrap_i64(v589))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L11
	} else {
		goto L188
	}
L188:
	;
	v741 = v545
	goto L157
L189:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v589
	v848 = v478
	v851 = v545
	v855 = v482
	goto L18
L190:
	;
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	if v633 == int64(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	F_pfree(m, base.I32_wrap_i64(v633))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L11
	} else {
		goto L192
	}
L192:
	;
	goto L189
L193:
	;
	v645 = v641 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v533)+12)) = v645
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	v651 = *(*int64)(unsafe.Add(mBase, uint32(v647+v645<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v651
	v848 = v478
	v851 = v519
	v855 = v482
	goto L18
L194:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	if v664 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v656 | int32(_a_F__bt_advance_array_keys_4)
	v848 = v478
	v851 = v655
	v855 = v482
	goto L18
L196:
	;
	goto L197
L197:
	;
	if v656&int32(1) != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v656 & int32(-524354)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v533)+20))
	v676 = *(*int64)(unsafe.Add(mBase, uint32(v675)+8))
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	v678 = int32(*(*int16)(unsafe.Add(mBase, uint32(v533)+16)))
	v679 = F_datumCopy(m, v676, v677, v678)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L11
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v682 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v664)+16))
	v686 = m.T0[v685].(func(*base.Module, int32, int64, int32) int64)(m, v500, v682, v29+int32(8))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L11
	} else {
		goto L202
	}
L201:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v679
	v848 = v478
	v851 = v655
	v855 = v482
	goto L18
L202:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
	if v688 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+19)))
	if v691 != int32(1) {
		v741 = v655
		goto L157
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v533)+24))
	if v716 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L206:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v694&int32(33554432) == int32(0) {
		v741 = v655
		goto L157
	} else {
		goto L207
	}
L207:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v699 != 0 {
		v707 = v694
		goto L208
	} else {
		goto L209
	}
L208:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v707&int32(-7864386) | int32(65)
	v848 = v478
	v851 = v655
	v855 = v482
	goto L18
L209:
	;
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	if v700 == int64(0) {
		v707 = v694
		goto L208
	} else {
		goto L210
	}
L210:
	;
	F_pfree(m, base.I32_wrap_i64(v700))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L11
	} else {
		goto L211
	}
L211:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v707 = v706
	goto L208
L212:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v731 != 0 {
		goto L218
	} else {
		goto L219
	}
L213:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v716)+12))
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v716)+48))
	v723 = F_FunctionCall2Coll(m, v716+int32(16), v721, v686, v722)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L11
	} else {
		goto L214
	}
L214:
	;
	if v723 != int64(0) {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v727 != 0 {
		v741 = v655
		goto L157
	} else {
		goto L216
	}
L216:
	;
	F_pfree(m, base.I32_wrap_i64(v686))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L11
	} else {
		goto L217
	}
L217:
	;
	v741 = v655
	goto L157
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v686
	v848 = v478
	v851 = v655
	v855 = v482
	goto L18
L219:
	;
	v732 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	if v732 == int64(0) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	F_pfree(m, base.I32_wrap_i64(v732))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L11
	} else {
		goto L221
	}
L221:
	;
	goto L218
L222:
	;
	if int32(0) < v513 {
		v513 = v513 - int32(1)
		v519 = v741
		goto L155
	} else {
		goto L239
	}
L223:
	;
	if v33 == int32(1) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+18)))
	if v758 != 0 {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v750 = int32(0)
	goto L228
L227:
	;
	v750 = v744 - int32(1)
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533)+12)) = v750
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v533)+8))
	v756 = *(*int64)(unsafe.Add(mBase, uint32(v752+v750<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = v756
	goto L222
L229:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v537)+48)) = int64(0)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v770 = v768 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v770
	v774 = int32(0)
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+19)))
	if base.B2i32(v502 == base.B2i32(v768&int32(33554432) == v774))|base.B2i32(v777 != int32(1)) == v774 {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v759 = *(*int64)(unsafe.Add(mBase, uint32(v537)+48))
	if v759 == int64(0) {
		goto L229
	} else {
		goto L231
	}
L231:
	;
	F_pfree(m, base.I32_wrap_i64(v759))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L11
	} else {
		goto L232
	}
L232:
	;
	goto L229
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v770 | int32(65)
	goto L222
L234:
	;
	goto L235
L235:
	;
	if v33 == int32(1) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v770 | int32(_a_F__bt_advance_array_keys_2)
	goto L222
L237:
	;
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v770 | int32(_a_F__bt_advance_array_keys_1)
	goto L222
L239:
	;
	goto L156
L240:
	;
	v831 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)) = uint8(v831)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+17)) = uint8(v831)
	v1056 = v493
	goto L4
L241:
	;
	v866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+34)))
	v868 = v866 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+34)) = uint16(v868)
	v883 = v848
	v890 = v855
	v892 = v130
	goto L14
L242:
	;
	if v883 != 0 {
		goto L262
	} else {
		goto L263
	}
L243:
	;
	v896 = v883
	goto L245
L244:
	;
	v896 = v890
	goto L245
L245:
	;
	if v896&int32(1) != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v899 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = l5 + v899
	v909 = F__bt_check_compare(m, l0, v33, l2, l3, l4, int32(0), l6^v899, v29+int32(15), v29+int32(8))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L11
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	if l6 != 0 {
		goto L242
	} else {
		goto L259
	}
L249:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	if v919 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L250:
	;
	if v909 == int32(0) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)))
	if v913 != 0 {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	v914 = int32(1)
	if l1 == int32(0) {
		v1056 = v914
		goto L4
	} else {
		goto L253
	}
L253:
	;
	v917 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)) = uint8(v917)
	v1056 = v914
	goto L4
L254:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v924 = F__bt_advance_array_keys(m, l0, l1, l2, l3, l4, v922, int32(1))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L11
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	if l6 != 0 {
		goto L242
	} else {
		goto L258
	}
L257:
	;
	v1056 = int32(0)
	goto L4
L258:
	;
	v1056 = int32(0)
	goto L4
L259:
	;
	v1056 = int32(0)
	goto L4
L260:
	;
	v1030 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)) = uint8(v1030)
	v1032 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+17)) = uint8(v1032)
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)))
	if v1035 != v1030 {
		v1056 = v1032
		goto L4
	} else {
		goto L288
	}
L261:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v1007 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L262:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)))
	if v892&(v955^int32(-1)) == int32(0) {
		goto L260
	} else {
		goto L272
	}
L263:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if l2 == v929 {
		goto L261
	} else {
		goto L264
	}
L264:
	;
	if v929 == int32(0) {
		goto L262
	} else {
		goto L265
	}
L265:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929)+7)))
	if v933&int32(32) == int32(0) {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v948 = int32(0)
	v952 = F__bt_tuple_before_array_skeys(m, l0, v33, v929, l4, v947, v948, v948, v35+int32(18))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L11
	} else {
		goto L270
	}
L267:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v34)+192))
	v945 = int32(*(*int16)(unsafe.Add(mBase, uint32(v944)+8)))
	v947 = v945
	goto L266
L268:
	;
	v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v929)+4)))
	if v938&int32(_a_F__bt_advance_array_keys_5) != 0 {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v947 = v938 & int32(4095)
	goto L266
L270:
	;
	if v952 != 0 {
		goto L261
	} else {
		goto L271
	}
L271:
	;
	goto L262
L272:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v961 == int32(0) {
		goto L260
	} else {
		goto L273
	}
L273:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v964)+52))
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961)+7)))
	if v966&int32(32) == int32(0) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v982 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v982
	v992 = F__bt_check_compare(m, l0, v982-v33, v961, v980, v965, v982, v982, v29+int32(15), v29+int32(8))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L11
	} else {
		goto L278
	}
L275:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v964)+192))
	v978 = int32(*(*int16)(unsafe.Add(mBase, uint32(v977)+8)))
	v980 = v978
	goto L274
L276:
	;
	v971 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v961)+4)))
	if v971&int32(_a_F__bt_advance_array_keys_5) != 0 {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v980 = v971 & int32(4095)
	goto L274
L278:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+15)))
	if v994 != 0 {
		goto L260
	} else {
		goto L279
	}
L279:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v981)+8))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v1000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v995+v996*int32(56))+6)))
	if v1000 == int32(3) {
		goto L260
	} else {
		goto L280
	}
L280:
	;
	goto L261
L281:
	;
	v1015 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)) = uint8(v1015)
	v1018 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+17)) = uint8(v1018)
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v1020 == v1015 {
		v1056 = v1015
		goto L4
	} else {
		goto L286
	}
L282:
	;
	v1010 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+34)))
	if v1010 < int32(4) {
		goto L281
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1013 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)) = uint8(v1013)
	goto L260
L285:
	;
	goto L284
L286:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	F__bt_parallel_primscan_schedule(m, l0, v1023)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L11
	} else {
		goto L287
	}
L287:
	;
	v1056 = v1015
	goto L4
L288:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+19)) = uint8(v892)
	if v33 != int32(1) {
		v1056 = v1032
		goto L4
	} else {
		goto L289
	}
L289:
	;
	v1041 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v1043 = v1041 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v1043)
	v1056 = v1032
	goto L4
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
	F_errfinish(m, int32(_a_F__bt_buildadd_5), int32(940), int32(_a_F__bt_buildadd_9))
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
	F_errfinish(m, int32(_a_F__bt_buildadd_5), int32(738), int32(_a_F__bt_buildadd_6))
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
	Fn14297(m, l0, l1, int32(_a_F__bt_checkpage_0), int32(829), int32(_a_F__bt_checkpage_1), int32(818))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F__bt_end_vacuum_callback(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	v10 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_vacuum_callback[0]))
	v14 = F_LWLockAcquire(m, v10+int32(2560), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_vacuum_callback[1]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v18 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_vacuum_callback[0]))
	F_LWLockRelease(m, v64+int32(2560))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L11
	}
L4:
	;
	v23 = base.I32_wrap_i64(l1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
	v26 = int32(0)
	goto L5
L5:
	;
	v36 = v17 + int32(12) + v26*int32(12)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != v24 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v53 = v26 + int32(1)
	if v53 != v18 {
		v26 = v53
		goto L5
	} else {
		goto L10
	}
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	if v39 != v40 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = v17 + v18*int32(12)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v18 - int32(1)
	goto L3
L10:
	;
	goto L6
L11:
	;
	return
}
func F__bt_find_extreme_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v73 int64
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15+v16<<(uint(int32(2))%32)-int32(4))))
	v24 = F_get_opfamily_member(m, v22, l2, l2, base.I32_extend16_s(l3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L23
	}
L2:
	;
	return int64(0)
L3:
	;
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = F_get_opcode(m, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
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
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L20
	}
L7:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_fmgr_info(m, v28, v12+int32(20))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
	if int32(2) <= l5 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = int32(1)
	v48 = v36
	goto L13
L11:
	;
	v73 = v36
	goto L12
L12:
	;
	m.G0 = v12 + int32(48)
	return v73
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v54 = l4 + v42<<(uint(int32(3))%32)
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	v56 = F_FunctionCall2Coll(m, v12+int32(20), v51, v55, v48)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v73 = v61
	goto L12
L15:
	;
	if v56 != int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	v61 = v60
	goto L18
L17:
	;
	v61 = v48
	goto L18
L18:
	;
	v63 = v42 + int32(1)
	if v63 != l5 {
		v42 = v63
		v48 = v61
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+208))
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
	F_errmsg_internal(m, int32(_a_F__bt_find_extreme_element_0), v12)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F__bt_find_extreme_element_1), int32(2607), int32(_a_F__bt_find_extreme_element_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v24
	F_errmsg_internal(m, int32(_a_F__bt_find_extreme_element_3), v12+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F__bt_find_extreme_element_1), int32(2610), int32(_a_F__bt_find_extreme_element_2))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
					F_errfinish(m, int32(_a_F__bt_getmeta_1), int32(159), int32(_a_F__bt_getmeta_2))
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
						F_errfinish(m, int32(_a_F__bt_getmeta_1), int32(159), int32(_a_F__bt_getmeta_2))
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
							F_errfinish(m, int32(_a_F__bt_getmeta_1), int32(168), int32(_a_F__bt_getmeta_2))
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
	v46 = F__bt_getbuf(m, l0, v26, int32(3))
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
	F_UnlockReleaseBuffer(m, v46)
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
				F_UnlockBuffer(m, v29)
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
					F_UnlockReleaseBuffer(m, v40)
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
			F_UnlockBuffer(m, v47)
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
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var __phi80 int32
	_ = __phi80
	var v82 int32
	_ = v82
	var __phi82 int32
	_ = __phi82
	var v86 int32
	_ = v86
	var __phi86 int32
	_ = __phi86
	var v96 int32
	_ = v96
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v262 int64
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int64
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int64
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int64
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v366 int32
	_ = v366
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v19 + int32(32)
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v29 = v27 + int32(20)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v31 = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(v31)%32))))
	v35 = int32(_a_F__bt_set_startikey_0)
	v37 = v27 + v34&v35
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29+v38<<(uint(v31)%32))))
	v45 = v27 + v42&v35
	v46 = F__bt_keep_natts_fast(m, v25, v37, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v48 <= int32(0) {
		v426 = v3
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v475 != 0 {
		goto L1
	} else {
		goto L135
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v447
	v453 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v453)
	goto L5
L7:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v426
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v431)
	if v431 != int32(1) {
		goto L1
	} else {
		goto L134
	}
L8:
	;
	v52 = int32(0)
	v57 = v3
	v63 = v3
	goto L9
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v71 = v68 + v63*int32(56)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72&int32(4) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	if v393&int32(1) != 0 {
		v447 = v410
		goto L6
	} else {
		goto L133
	}
L11:
	;
	v410 = v63 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v410 < v411 {
		v52 = v393
		v57 = v398
		v63 = v410
		goto L9
	} else {
		goto L132
	}
L12:
	;
	v393 = v52
	v398 = v283
	goto L11
L13:
	;
	if v52&int32(1) == int32(0) {
		v426 = v63
		goto L7
	} else {
		goto L131
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+4)))
	if v46 < v76 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+6)))
	if v212 != int32(3) {
		goto L70
	} else {
		goto L71
	}
L17:
	;
	__phi80 = v75
	__phi82 = v76
	__phi86 = v75 + int32(4)
	v80 = __phi80
	v82 = __phi82
	v86 = __phi86
	goto L18
L18:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v96&int32(1) != 0 {
		goto L13
	} else {
		goto L20
	}
L19:
	;
	goto L13
L20:
	;
	v101 = F_index_getattr_2(m, v37, v82, v26, v19+int32(31))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86))))
	v106 = F_index_getattr_2(m, v45, v103, v26, v19+int32(30))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)))
	if v108 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v109&int32(1) != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v115 = v80 + int32(16)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v80)+48))
	v118 = F_FunctionCall2Coll(m, v115, v116, v101, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L26
	}
L25:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v80)+48))
	v165 = F_FunctionCall2Coll(m, v115, v163, v106, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L46
	}
L26:
	;
	v120 = base.I32_wrap_i64(v118)
	if v120 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v124 = int32(1)
	goto L29
L28:
	;
	v124 = int32(0) - v120
	goto L29
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v125&int32(16777216) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v128 = v124
	goto L32
L31:
	;
	v128 = v120
	goto L32
L32:
	;
	v131 = v128 | v125&int32(16)
	if v131 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+6)))
	switch v134 - int32(1) {
	case 0:
		goto L34
	case 1:
		goto L35
	default:
		goto L36
	case 3:
		goto L38
	case 4:
		goto L37
	}
L34:
	;
	if int32(0) <= v128 {
		goto L13
	} else {
		goto L45
	}
L35:
	;
	if int32(0) < v128 {
		goto L13
	} else {
		goto L44
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L41
	}
L37:
	;
	if int32(0) < v128 {
		goto L25
	} else {
		goto L40
	}
L38:
	;
	if int32(0) <= v128 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L13
L40:
	;
	goto L13
L41:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v145
	F_errmsg_internal(m, int32(_a_F__bt_set_startikey_1), v19+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F__bt_set_startikey_2), int32(1803), int32(_a_F__bt_set_startikey_3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	goto L25
L45:
	;
	goto L25
L46:
	;
	v167 = base.I32_wrap_i64(v165)
	if v167 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v171 = int32(1)
	goto L49
L48:
	;
	v171 = int32(0) - v167
	goto L49
L49:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v172&int32(16777216) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v175 = v171
	goto L52
L51:
	;
	v175 = v167
	goto L52
L52:
	;
	if v175|v172&int32(16) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v131 == int32(0) {
		goto L13
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+60)))
	if v208 <= v46 {
		__phi80 = v80 + int32(56)
		__phi82 = v208
		__phi86 = v80 + int32(60)
		v80 = __phi80
		v82 = __phi82
		v86 = __phi86
		goto L18
	} else {
		goto L69
	}
L56:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+6)))
	switch v181 - int32(1) {
	case 0:
		goto L61
	case 1:
		goto L60
	default:
		goto L57
	case 3:
		goto L59
	case 4:
		goto L58
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L3
	} else {
		goto L66
	}
L58:
	;
	if v175 <= int32(0) {
		goto L13
	} else {
		goto L65
	}
L59:
	;
	if v175 < int32(0) {
		goto L13
	} else {
		goto L64
	}
L60:
	;
	if int32(0) < v175 {
		goto L13
	} else {
		goto L63
	}
L61:
	;
	if int32(0) <= v175 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	v393 = v52
	v398 = v57
	goto L11
L63:
	;
	v393 = v52
	v398 = v57
	goto L11
L64:
	;
	v393 = v52
	v398 = v57
	goto L11
L65:
	;
	v393 = v52
	v398 = v57
	goto L11
L66:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v196
	F_errmsg_internal(m, int32(_a_F__bt_set_startikey_1), v19)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F__bt_set_startikey_2), int32(1803), int32(_a_F__bt_set_startikey_3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	goto L19
L70:
	;
	v215 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v46 < v215 {
		goto L13
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v72&int32(32) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L73:
	;
	v219 = F_index_getattr_2(m, v37, v215, v26, v19+int32(31))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+4)))
	v224 = F_index_getattr_2(m, v45, v221, v26, v19+int32(30))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v227&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v226&int32(1) != 0 {
		goto L13
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v226&int32(1) != 0 {
		goto L13
	} else {
		goto L81
	}
L79:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v232&int32(1) != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v393 = v52
	v398 = v57
	goto L11
L81:
	;
	v238 = v71 + int32(16)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v71)+48))
	v241 = F_FunctionCall2Coll(m, v238, v239, v219, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	if v241 == int64(0) {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v245&int32(1) != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v71)+48))
	v250 = F_FunctionCall2Coll(m, v238, v248, v224, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	if v250 == int64(0) {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	v393 = v52
	v398 = v57
	goto L11
L87:
	;
	v258 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v46 <= v258 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v283 = v57 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v287 = v284 + v57<<(uint(int32(5))%32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v288 != int32(-1) {
		goto L99
	} else {
		goto L100
	}
L90:
	;
	v262 = F_index_getattr_2(m, v37, v258, v26, v19+int32(31))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v265&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v264&int32(1) == int32(0) {
		goto L13
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v264&int32(1) != 0 {
		goto L13
	} else {
		goto L96
	}
L95:
	;
	v393 = v52
	v398 = v57
	goto L11
L96:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v71)+48))
	v278 = F_FunctionCall2Coll(m, v71+int32(16), v276, v262, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	if v278 == int64(0) {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v393 = v52
	v398 = v57
	goto L11
L99:
	;
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v46 <= v291 {
		goto L13
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+19)))
	if v310 != 0 {
		goto L12
	} else {
		goto L106
	}
L102:
	;
	v295 = F_index_getattr_2(m, v37, v291, v26, v19+int32(31))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v301 = int32(0)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)))
	v306 = F__bt_binsrch_array_skey(m, v297+v63*int32(28), v301, v301, v295, v303, v287, v71, v19+int32(24))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v309 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	v393 = int32(1)
	v398 = v283
	goto L11
L106:
	;
	v311 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v46 < v311 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	v315 = F_index_getattr_2(m, v37, v311, v26, v19+int32(31))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v317 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+4)))
	v320 = F_index_getattr_2(m, v45, v317, v26, v19+int32(30))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L3
	} else {
		goto L109
	}
L109:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+19)))
	if v322 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+19)))
	if v345 != 0 {
		goto L12
	} else {
		goto L121
	}
L111:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)))
	if v323 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v287)+24))
	if v324 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v324)+12))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v324)+48))
	v329 = F_FunctionCall2Coll(m, v324+int32(16), v327, v315, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L3
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v287)+28))
	if v333 == int32(0) {
		goto L110
	} else {
		goto L118
	}
L116:
	;
	if v329 == int64(0) {
		goto L13
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v333)+48))
	v340 = F_FunctionCall2Coll(m, v333+int32(16), v338, v315, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L3
	} else {
		goto L119
	}
L119:
	;
	if v340 == int64(0) {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	goto L110
L121:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+30)))
	if v346 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = int32(0)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v287)+24))
	if v349 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v349)+48))
	v354 = F_FunctionCall2Coll(m, v349+int32(16), v352, v320, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v287)+28))
	if v358 == int32(0) {
		goto L12
	} else {
		goto L128
	}
L126:
	;
	if v354 == int64(0) {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v358)+48))
	v365 = F_FunctionCall2Coll(m, v358+int32(16), v363, v320, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	if v365 != int64(0) {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	goto L13
L131:
	;
	v447 = v63
	goto L6
L132:
	;
	goto L10
L133:
	;
	v426 = v410
	goto L7
L134:
	;
	goto L5
L135:
	;
	v476 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(17)))) = uint8(v476)
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)))) = v476
	goto L1
}
func F__bt_start_array_keys(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v3 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	v99 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+18)) = uint16(v99)
	return
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v25 = v22 + v18<<(uint(int32(5))%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v29 = v21 + v26*int32(56)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v30 != int32(-1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v88 = v18 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if v88 < v89 {
		v18 = v88
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v33 = int32(1)
	if l1 != v33 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v38 = v30 - v33
	goto L12
L11:
	;
	v38 = int32(0)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v44
	goto L6
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = int64(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v58 = v56 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v58
	v62 = int32(0)
	v64 = int32(1)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+19)))
	if base.B2i32(base.B2i32(v56&int32(33554432) == v62) == base.B2i32(l1 == v64))|base.B2i32(v67 != v64) == v62 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v47 == int64(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, base.I32_wrap_i64(v47))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v58 | int32(65)
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
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v58 | int32(_a_F__bt_start_array_keys_0)
	goto L6
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v58 | int32(_a_F__bt_start_array_keys_1)
	goto L6
L24:
	;
	goto L5
}
func F_bt_page_items(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_items_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
	var v115 int32
	_ = v115
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
											v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
												v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
												v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
													v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
												v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
													v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
													v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
														v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
											v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
												v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
												v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
													v115 = F_errdetail(m, int32(_a_F_bt_report_duplicate_4), v14)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_bt_report_duplicate_5), int32(908), int32(_a_F_bt_report_duplicate_6))
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
