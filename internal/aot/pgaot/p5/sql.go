package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_sql_stmt_retval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
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
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var __phi683 int32
	_ = __phi683
	var v687 int32
	_ = v687
	var __phi687 int32
	_ = __phi687
	var v688 int32
	_ = v688
	var __phi688 int32
	_ = __phi688
	var v689 int32
	_ = v689
	var __phi689 int32
	_ = __phi689
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v759 int32
	_ = v759
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v843 int32
	_ = v843
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v903 int32
	_ = v903
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	v6 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(160)
	m.G0 = v22
	*(*int32)(unsafe.Add(mBase, uint32(v22)+156)) = v6
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+155)) = uint8(v6)
	if l1 == int32(2278) {
		v903 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L82
	} else {
		goto L205
	}
L2:
	;
	m.G0 = v22 + int32(160)
	return v903
L3:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+155)))
	if v772 != int32(1) {
		v903 = v759
		goto L2
	} else {
		goto L183
	}
L5:
	;
	v674 = int32(1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+155)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	v678 = v664 + v674
	if v678 <= v363 {
		goto L167
	} else {
		goto L168
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L82
	} else {
		goto L162
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L82
	} else {
		goto L153
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L82
	} else {
		goto L147
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L82
	} else {
		goto L141
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L82
	} else {
		goto L135
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v32 <= int32(0) {
		v159 = v6
		v166 = v6
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v159 == int32(0) {
		goto L10
	} else {
		goto L57
	}
L13:
	;
	v35 = int32(0)
	if v35 < v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = v32
	goto L16
L15:
	;
	v38 = v35
	goto L16
L16:
	;
	v40 = v38 & int32(3)
	if int32(4) <= v32 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = int32(0)
	v54 = v6
	v57 = v6
	v64 = v6
	goto L20
L18:
	;
	v105 = v6
	v108 = v6
	v115 = v6
	goto L19
L19:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v125 = v105
	v127 = v6
	v128 = v108
	v135 = v115
	goto L48
L20:
	;
	v68 = v45 + v54<<(uint(int32(2))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+24)))
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v40 == int32(0) {
		v159 = v80
		v166 = v90
		goto L12
	} else {
		goto L47
	}
L22:
	;
	v74 = v72
	goto L24
L23:
	;
	v74 = v57
	goto L24
L24:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+24)))
	if v75 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = v71
	goto L27
L26:
	;
	v76 = v74
	goto L27
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)))
	if v77 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v78 = v70
	goto L30
L29:
	;
	v78 = v76
	goto L30
L30:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+24)))
	if v79 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v80 = v69
	goto L33
L32:
	;
	v80 = v78
	goto L33
L33:
	;
	if v73 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v87 = v68
	goto L36
L35:
	;
	v87 = v64
	goto L36
L36:
	;
	if v75 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v88 = v68 + int32(4)
	goto L39
L38:
	;
	v88 = v87
	goto L39
L39:
	;
	if v77 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v89 = v68 + int32(8)
	goto L42
L41:
	;
	v89 = v88
	goto L42
L42:
	;
	if v79 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v90 = v68 + int32(12)
	goto L45
L44:
	;
	v90 = v89
	goto L45
L45:
	;
	v91 = int32(4)
	v92 = v54 + v91
	v94 = v53 + v91
	if v94 != v38&int32(2147483644) {
		v53 = v94
		v54 = v92
		v57 = v80
		v64 = v90
		goto L20
	} else {
		goto L46
	}
L46:
	;
	goto L21
L47:
	;
	v105 = v92
	v108 = v80
	v115 = v90
	goto L19
L48:
	;
	v139 = v117 + v125<<(uint(int32(2))%32)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+24)))
	if v141 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v159 = v142
	v166 = v143
	goto L12
L50:
	;
	v142 = v140
	goto L52
L51:
	;
	v142 = v128
	goto L52
L52:
	;
	if v141 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v143 = v139
	goto L55
L54:
	;
	v143 = v135
	goto L55
L55:
	;
	v144 = int32(1)
	v147 = v127 + v144
	if v147 != v40 {
		v125 = v125 + v144
		v127 = v147
		v128 = v142
		v135 = v143
		goto L48
	} else {
		goto L56
	}
L56:
	;
	goto L49
L57:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v170 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v187 = int32(0)
	if v185 == v187 {
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v159)+144))
	v185 = v173
	v186 = base.B2i32(v174 == int32(0))
	goto L58
L60:
	;
	goto L61
L61:
	;
	if base.Ui32(int32(3)) < base.Ui32(v170-int32(2)) {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v159)+96))
	if v182 == int32(0) {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v185 = v182
	v186 = int32(1)
	goto L58
L64:
	;
	v276 = F_get_typtype(m, l1)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	v275 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v197 <= int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v275 = int32(0)
	goto L64
L69:
	;
	goto L70
L70:
	;
	if v197 != int32(1) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v275 = v262
	goto L64
L72:
	;
	v203 = int32(0)
	if v203 < v197 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v244 = v187
	v245 = v187
	goto L74
L74:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250+v244<<(uint(int32(2))%32))))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+26)))
	v262 = v245 + (v255 ^ int32(1))
	goto L71
L75:
	;
	v206 = v197
	goto L77
L76:
	;
	v206 = v203
	goto L77
L77:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v212 = int32(0)
	v215 = v212
	v216 = v212
	v217 = v187
	goto L78
L78:
	;
	v222 = int32(2)
	v224 = v211 + v216<<(uint(v222)%32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+26)))
	v227 = int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+26)))
	v234 = v217 + (v226 ^ v227) + (v231 ^ v227)
	v236 = v216 + v222
	v238 = v215 + v222
	if v238 != v206&int32(2147483646) {
		v215 = v238
		v216 = v236
		v217 = v234
		goto L78
	} else {
		goto L80
	}
L79:
	;
	if v206&int32(1) == int32(0) {
		v262 = v234
		goto L71
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v244 = v236
	v245 = v234
	goto L74
L82:
	;
	return int32(0)
L83:
	;
	v283 = v276&int32(255) - int32(98)
	v290 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v283))|base.B2i32(int32(1)<<(uint(v283)%32)&int32(_a_F_check_sql_stmt_retval_0) == v290) == v290 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v275 != int32(1) {
		goto L9
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.B2i32(l1 != int32(2249))&base.B2i32(v276 != int32(99)) != 0 {
		goto L6
	} else {
		goto L100
	}
L87:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v304 = F_coerce_fn_result_column(m, v298, l1, int32(-1), v186, v22+int32(156), v22+int32(155))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L82
	} else {
		goto L88
	}
L88:
	;
	if v304 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v759 = int32(0)
	goto L4
L90:
	;
	goto L91
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L82
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L82
	} else {
		goto L93
	}
L93:
	;
	v314 = F_format_type_be(m, l1)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L82
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v314
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_1), v22+int32(112))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v323 = F_exprType(m, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L82
	} else {
		goto L96
	}
L96:
	;
	v325 = F_format_type_be(m, v323)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L82
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v325
	F_errdetail(m, int32(_a_F_check_sql_stmt_retval_2), v22+int32(96))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L82
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2279), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L82
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	if base.B2i32(l3 == int32(112))|base.B2i32(v275 != int32(1)) != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if l2 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v355 = F_coerce_fn_result_column(m, v349, l1, int32(-1), v186, v22+int32(156), v22+int32(155))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L82
	} else {
		goto L103
	}
L103:
	;
	if v355 == int32(0) {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v759 = int32(0)
	goto L4
L105:
	;
	v903 = int32(1)
	goto L2
L106:
	;
	goto L107
L107:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v185 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v664 = int32(0)
	goto L5
L109:
	;
	goto L110
L110:
	;
	v367 = int32(0)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v368 <= v367 {
		v664 = v367
		goto L5
	} else {
		goto L111
	}
L111:
	;
	v371 = int32(0)
	v373 = v371
	v381 = v371
	v382 = v367
	goto L112
L112:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v381<<(uint(int32(2))%32))))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+26)))
	if v397 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v664 = v496
	goto L5
L114:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+155)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	v405 = v382
	v408 = v401
	v416 = v400
	goto L117
L115:
	;
	v487 = v373
	v496 = v382
	goto L116
L116:
	;
	v507 = v381 + int32(1)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v507 < v508 {
		v373 = v487
		v381 = v507
		v382 = v496
		goto L112
	} else {
		goto L134
	}
L117:
	;
	v422 = v405 + int32(1)
	if v363 < v422 {
		goto L8
	} else {
		goto L119
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+155)) = uint8(v472)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+156)) = v471
	v476 = v373 + int32(1)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v432)+76))
	v483 = F_coerce_fn_result_column(m, v396, v477, v478, v186, v22+int32(156), v22+int32(155))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L82
	} else {
		goto L132
	}
L119:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v430 = l2 + v424<<(uint(int32(4))%32) + v405*int32(100)
	v432 = v430 + int32(20)
	v433 = int32(0)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+111)))
	if base.B2i32(l4 == v433)|base.B2i32(v435&int32(1) == v433) == v433 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L118
L121:
	;
	v443 = int32(1)
	v446 = int32(0)
	v451 = F_makeConst(m, int32(23), int32(-1), v446, int32(4), v446, v443, v443)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L82
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if v435&int32(1) != 0 {
		v405 = v422
		goto L117
	} else {
		goto L131
	}
L124:
	;
	if v408 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v408)+4)))
	v457 = v453 + int32(1)
	goto L127
L126:
	;
	v457 = int32(1)
	goto L127
L127:
	;
	v459 = int32(0)
	v461 = F_makeTargetEntry(m, v451, base.I32_extend16_s(v457), v459, v459)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L82
	} else {
		goto L128
	}
L128:
	;
	v463 = F_lappend(m, v408, v461)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L82
	} else {
		goto L129
	}
L129:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+91)))
	if v465 != int32(1) {
		v471 = v463
		v472 = v443
		goto L120
	} else {
		goto L130
	}
L130:
	;
	v405 = v422
	v408 = v463
	v416 = v443
	goto L117
L131:
	;
	v471 = v408
	v472 = v416
	goto L120
L132:
	;
	if v483 == int32(0) {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v487 = v476
	v496 = v422
	goto L116
L134:
	;
	goto L113
L135:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L82
	} else {
		goto L136
	}
L136:
	;
	v536 = F_format_type_be(m, l1)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L82
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v536
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_1), v22)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L82
	} else {
		goto L138
	}
L138:
	;
	F_errdetail(m, int32(_a_F_check_sql_stmt_retval_5), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L82
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2226), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L82
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L82
	} else {
		goto L142
	}
L142:
	;
	v558 = F_format_type_be(m, l1)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L82
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v558
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_1), v22+int32(128))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L82
	} else {
		goto L144
	}
L144:
	;
	F_errdetail(m, int32(_a_F_check_sql_stmt_retval_6), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L82
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2264), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L82
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L82
	} else {
		goto L148
	}
L148:
	;
	v582 = F_format_type_be(m, l1)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L82
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v582
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_1), v22+int32(32))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L82
	} else {
		goto L150
	}
L150:
	;
	F_errdetail(m, int32(_a_F_check_sql_stmt_retval_7), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L82
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2368), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L82
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L82
	} else {
		goto L154
	}
L154:
	;
	v606 = F_format_type_be(m, l1)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L82
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v606
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_1), v22-int32(-64))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L82
	} else {
		goto L156
	}
L156:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	v615 = F_exprType(m, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L82
	} else {
		goto L157
	}
L157:
	;
	v617 = F_format_type_be(m, v615)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L82
	} else {
		goto L158
	}
L158:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	v620 = F_format_type_be(m, v619)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L82
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v617
	F_errdetail(m, int32(_a_F_check_sql_stmt_retval_8), v22+int32(48))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L82
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2404), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L82
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L82
	} else {
		goto L163
	}
L163:
	;
	v642 = F_format_type_be(m, l1)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L82
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v642
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_9), v22+int32(80))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L82
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2444), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L82
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	__phi683 = v678
	__phi687 = v676
	__phi688 = v675
	__phi689 = v664
	v683 = __phi683
	v687 = __phi687
	v688 = __phi688
	v689 = __phi689
	goto L170
L168:
	;
	v739 = v676
	v740 = v675
	goto L169
L169:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+155)) = uint8(v740)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+156)) = v739
	v759 = v674
	goto L4
L170:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v689<<(uint(int32(4))%32))+29)))
	if v702 == int32(0) {
		goto L1
	} else {
		goto L172
	}
L171:
	;
	v739 = v727
	v740 = v728
	goto L169
L172:
	;
	if l4 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v705 = int32(1)
	v708 = int32(0)
	v713 = F_makeConst(m, int32(23), int32(-1), v708, int32(4), v708, v705, v705)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L82
	} else {
		goto L176
	}
L174:
	;
	v727 = v687
	v728 = v688
	goto L175
L175:
	;
	v730 = v683 + int32(1)
	if v730 <= v363 {
		__phi683 = v730
		__phi687 = v727
		__phi688 = v728
		__phi689 = v683
		v683 = __phi683
		v687 = __phi687
		v688 = __phi688
		v689 = __phi689
		goto L170
	} else {
		goto L182
	}
L176:
	;
	if v687 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v715 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v687)+4)))
	v719 = v715 + int32(1)
	goto L179
L178:
	;
	v719 = int32(1)
	goto L179
L179:
	;
	v721 = int32(0)
	v723 = F_makeTargetEntry(m, v713, base.I32_extend16_s(v719), v721, v721)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L82
	} else {
		goto L180
	}
L180:
	;
	v725 = F_lappend(m, v687, v723)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L82
	} else {
		goto L181
	}
L181:
	;
	v727 = v725
	v728 = v705
	goto L175
L182:
	;
	goto L171
L183:
	;
	v776 = F_palloc0(m, int32(168))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L82
	} else {
		goto L184
	}
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v776))) = int64(4294967363)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v781 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+24)) = uint8(v781)
	*(*int32)(unsafe.Add(mBase, uint32(v776)+8)) = v780
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v22)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v776)+76)) = v784
	v786 = int32(0)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	if v787 == v786 {
		v843 = v786
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v854 = F_palloc0(m, int32(136))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L82
	} else {
		goto L199
	}
L186:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	if v790 <= int32(0) {
		v843 = v786
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v797 = int32(0)
	v803 = v786
	goto L188
L188:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v787)+12))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v813+v797<<(uint(int32(2))%32))))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+26)))
	if v818 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v843 = v829
	goto L185
L190:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v817)+12))
	if v821 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v829 = v803
	goto L192
L192:
	;
	v831 = v797 + int32(1)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v787)+4))
	if v831 < v832 {
		v797 = v831
		v803 = v829
		goto L188
	} else {
		goto L198
	}
L193:
	;
	v823 = v821
	goto L195
L194:
	;
	v823 = int32(_a_F_check_sql_stmt_retval_10)
	goto L195
L195:
	;
	v824 = F_makeString(m, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L82
	} else {
		goto L196
	}
L196:
	;
	v826 = F_lappend(m, v803, v824)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L82
	} else {
		goto L197
	}
L197:
	;
	v829 = v826
	goto L192
L198:
	;
	goto L189
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854)+36)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v854)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = int32(101)
	v862 = F_makeAlias(m, int32(_a_F_check_sql_stmt_retval_11), v843)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L82
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854)+8)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(v854)+4)) = v862
	v866 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v854)+124)) = uint16(v866)
	v868 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v854)+20)) = uint8(v868)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v854
	v875 = F_list_make1_impl(m, int32(1), v22+int32(12))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L82
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v776)+52)) = v875
	v879 = F_palloc0(m, int32(8))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L82
	} else {
		goto L202
	}
L202:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v879))) = int64(4294967359)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v879
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v879
	v888 = F_list_make1_impl(m, int32(1), v22+int32(8))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L82
	} else {
		goto L203
	}
L203:
	;
	v891 = F_makeFromExpr(m, v888, int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L82
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v776)+60)) = v891
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v776)+44)) = uint8(v894)
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v776
	v903 = v759
	goto L2
L205:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L82
	} else {
		goto L206
	}
L206:
	;
	v927 = F_format_type_be(m, l1)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L82
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v927
	F_errmsg(m, int32(_a_F_check_sql_stmt_retval_1), v22+int32(16))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L82
	} else {
		goto L208
	}
L208:
	;
	F_errdetail(m, int32(_a_F_check_sql_stmt_retval_12), int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L82
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_check_sql_stmt_retval_3), int32(2415), int32(_a_F_check_sql_stmt_retval_4))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L82
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_map_sql_table_to_xmlschema(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	F_initStringInfo(m, v8+int32(96))
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
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L38
	}
L4:
	;
	v17 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v24 = v8 + int32(96)
	F_appendStringInfoString(m, v24, int32(_a_F_map_sql_table_to_xmlschema_0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	if v17 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v21 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l3
	F_appendStringInfo(m, v24, int32(_a_F_map_sql_table_to_xmlschema_1), v8-int32(-64))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v36 = v8 + int32(96)
	F_appendStringInfoString(m, v36, int32(_a_F_map_sql_table_to_xmlschema_2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+92)) = l0
	v45 = F_list_make1_impl(m, int32(1), v8+int32(60))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v47 = F_map_sql_typecoll_to_xmlschema_types(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoString(m, v36, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(_a_F_map_sql_table_to_xmlschema_3)
	F_appendStringInfo(m, v36, int32(_a_F_map_sql_table_to_xmlschema_4), v8+int32(48))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v58 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v85 = v8 + int32(96)
	F_appendStringInfoString(m, v85, int32(_a_F_map_sql_table_to_xmlschema_5))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L21:
	;
	v65 = int32(0)
	goto L22
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v58<<(uint(int32(4))%32)+v65*int32(100))+111)))
	if v73 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v77 = F_map_sql_identifier_to_xml_name(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L24:
	;
	v75 = v65 + int32(1)
	if v58 != v75 {
		v65 = v75
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L20
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	if l2 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_appendStringInfoString(m, v8+int32(96), int32(_a_F_map_sql_table_to_xmlschema_6))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_map_sql_table_to_xmlschema_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_map_sql_table_to_xmlschema_7)
	F_appendStringInfo(m, v85, int32(_a_F_map_sql_table_to_xmlschema_8), v8+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_map_sql_table_to_xmlschema_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_map_sql_table_to_xmlschema_9)
	F_appendStringInfo(m, v8+int32(96), int32(_a_F_map_sql_table_to_xmlschema_10), v8)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_map_sql_table_to_xmlschema_7)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_map_sql_table_to_xmlschema_11)
	F_appendStringInfo(m, v85, int32(_a_F_map_sql_table_to_xmlschema_10), v8+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	goto L30
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+96))
	m.G0 = v8 + int32(112)
	return v123
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l1
	F_errmsg_internal(m, int32(_a_F_map_sql_table_to_xmlschema_12), v8+int32(80))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_map_sql_table_to_xmlschema_13), int32(3533), int32(_a_F_map_sql_table_to_xmlschema_14))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_map_sql_typecoll_to_xmlschema_types(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	if l0 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v11)+140))
	m.G0 = v11 + int32(176)
	return v339
L2:
	;
	F_initStringInfo(m, v11+int32(140))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L35
	}
L3:
	;
	F_initStringInfo(m, v11+int32(140))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L18
	} else {
		goto L34
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v15 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = v2
	v23 = v2
	goto L8
L6:
	;
	v77 = v2
	goto L7
L7:
	;
	if v77 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L8:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v26 < v32 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v77 = v65
	goto L7
L10:
	;
	v37 = v20
	v38 = v32
	v39 = v26
	goto L13
L11:
	;
	v65 = v20
	goto L12
L12:
	;
	v72 = v23 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v72 < v73 {
		v20 = v65
		v23 = v72
		goto L8
	} else {
		goto L21
	}
L13:
	;
	v48 = v31 + v38<<(uint(int32(4))%32) + v39*int32(100)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+111)))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v65 = v58
	goto L12
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+88))
	v53 = F_list_append_unique_oid(m, v37, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v58 = v37
	v59 = v38
	goto L17
L17:
	;
	v61 = v39 + int32(1)
	if v61 < v59 {
		v37 = v58
		v38 = v59
		v39 = v61
		goto L13
	} else {
		goto L20
	}
L18:
	;
	return int32(0)
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v58 = v53
	v59 = v57
	goto L17
L20:
	;
	goto L14
L21:
	;
	goto L9
L22:
	;
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v86 <= v85 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v124 = v77
	goto L2
L24:
	;
	goto L25
L25:
	;
	v89 = v77
	v93 = v85
	goto L26
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v93<<(uint(int32(2))%32))))
	v102 = F_getBaseType(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L28
	}
L27:
	;
	v124 = v107
	goto L2
L28:
	;
	if v101 != v102 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v105 = F_list_append_unique_oid(m, v89, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L18
	} else {
		goto L32
	}
L30:
	;
	v107 = v89
	goto L31
L31:
	;
	v109 = v93 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v109 < v110 {
		v89 = v107
		v93 = v109
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v107 = v105
	goto L31
L33:
	;
	goto L27
L34:
	;
	goto L1
L35:
	;
	if v124 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v138 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v139 <= v138 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v145 = v138
	goto L38
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150+v145<<(uint(int32(2))%32))))
	v156 = F_map_sql_type_to_xml_name(m, v154, int32(-1))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L18
	} else {
		goto L40
	}
L39:
	;
	goto L1
L40:
	;
	v159 = v11 + int32(160)
	F_initStringInfo(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	if v154 == int32(142) {
		v315 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_0)
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_appendStringInfoString(m, v11+int32(160), v315)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L18
	} else {
		goto L96
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v156
	F_appendStringInfo(m, v159, int32(_a_F_map_sql_typecoll_to_xmlschema_types_1), v11+int32(128))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v171 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_2)
	if v154 <= int32(1041) {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	F_appendStringInfoString(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_3))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L18
	} else {
		goto L95
	}
L46:
	;
	if v154 == int32(1266) {
		goto L91
	} else {
		goto L92
	}
L47:
	;
	F_appendStringInfoString(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_4))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L18
	} else {
		goto L90
	}
L48:
	;
	F_appendStringInfoString(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_5))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L18
	} else {
		goto L89
	}
L49:
	;
	F_appendStringInfoString(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_6))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L18
	} else {
		goto L88
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = int64(-9223372036854775807 - 1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = int64(9223372036854775807)
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_7), v11+int32(80))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L18
	} else {
		goto L87
	}
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = int64(-9223372034707292161)
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_8), v11-int32(-64))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L18
	} else {
		goto L86
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = int64(-140737488322561)
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_9), v11+int32(48))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L18
	} else {
		goto L85
	}
L53:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_typecoll_to_xmlschema_types[0]))
	if v242 != 0 {
		goto L81
	} else {
		goto L82
	}
L54:
	;
	v232 = v11 + int32(160)
	F_appendStringInfoString(m, v232, int32(_a_F_map_sql_typecoll_to_xmlschema_types_10))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L18
	} else {
		goto L79
	}
L55:
	;
	v210 = F_get_typtype(m, v154)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L74
	}
L56:
	;
	switch v154 - int32(16) {
	case 0:
		goto L47
	case 1:
		goto L53
	case 2, 3, 6, 8:
		goto L55
	case 4:
		goto L50
	case 5:
		goto L52
	case 7:
		goto L51
	case 9:
		goto L54
	default:
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v154 <= int32(1113) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	switch v154 - int32(700) {
	case 0:
		goto L49
	case 1:
		goto L48
	default:
		goto L55
	}
L60:
	;
	if base.Ui32(v154-int32(1042)) < base.Ui32(int32(2)) {
		goto L54
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if v154 <= int32(1265) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	switch v154 - int32(1082) {
	case 0:
		goto L45
	case 1:
		goto L46
	default:
		goto L55
	}
L64:
	;
	if base.B2i32(v154 != int32(1114))&base.B2i32(v154 != int32(1184)) != 0 {
		goto L55
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if v154 == int32(1266) {
		goto L46
	} else {
		goto L72
	}
L67:
	;
	if v154 == int32(1184) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v197 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_11)
	goto L70
L69:
	;
	v197 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_12)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v197
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_13), v11+int32(112))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	v315 = v171
	goto L42
L72:
	;
	if v154 == int32(1700) {
		v315 = v171
		goto L42
	} else {
		goto L73
	}
L73:
	;
	goto L55
L74:
	;
	if v210 != int32(100) {
		v315 = v171
		goto L42
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+156)) = int32(-1)
	v218 = F_getBaseTypeAndTypmod(m, v154, v11+int32(156))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v11)+156))
	v221 = F_map_sql_type_to_xml_name(m, v218, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v221
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_14), v11+int32(16))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	v315 = v171
	goto L42
L79:
	;
	F_appendStringInfoString(m, v232, int32(_a_F_map_sql_typecoll_to_xmlschema_types_15))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	v315 = v171
	goto L42
L81:
	;
	v243 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_16)
	goto L83
L82:
	;
	v243 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_17)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v243
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_18), v11+int32(32))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	v315 = v171
	goto L42
L85:
	;
	v315 = v171
	goto L42
L86:
	;
	v315 = v171
	goto L42
L87:
	;
	v315 = v171
	goto L42
L88:
	;
	v315 = v171
	goto L42
L89:
	;
	v315 = v171
	goto L42
L90:
	;
	v315 = v171
	goto L42
L91:
	;
	v300 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_11)
	goto L93
L92:
	;
	v300 = int32(_a_F_map_sql_typecoll_to_xmlschema_types_12)
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v300
	F_appendStringInfo(m, v11+int32(160), int32(_a_F_map_sql_typecoll_to_xmlschema_types_19), v11+int32(96))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L18
	} else {
		goto L94
	}
L94:
	;
	v315 = v171
	goto L42
L95:
	;
	v315 = v171
	goto L42
L96:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v11)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v320
	F_appendStringInfo(m, v11+int32(140), int32(_a_F_map_sql_typecoll_to_xmlschema_types_20), v11)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L18
	} else {
		goto L97
	}
L97:
	;
	v328 = v145 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v328 < v329 {
		v145 = v328
		goto L38
	} else {
		goto L98
	}
L98:
	;
	goto L39
}
func F_read_sql_construct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
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
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	F_initStringInfo(m, v19+int32(-36))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = int32(_a_F_read_sql_construct_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[0])) = int32(2)
	v34 = int32(-1)
	v50 = v34
	v51 = int32(0)
	v52 = v34
	goto L6
L3:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(_a_F_read_sql_construct_1))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L59
	}
L4:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(_a_F_read_sql_construct_2))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L58
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l3
	F_errmsg(m, int32(_a_F_read_sql_construct_3), v21)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L55
	}
L6:
	;
	v54 = F_plpgsql_yylex(m, l9, l10, l11)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[0])) = v30
	if l7 != 0 {
		goto L36
	} else {
		goto L37
	}
L8:
	;
	if v50 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v59 = v58
	goto L11
L10:
	;
	v59 = v50
	goto L11
L11:
	;
	if base.B2i32(l0 != v54)&base.B2i32(l1 != v54)&base.B2i32(l2 != v54)|v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	switch v54 - int32(40) {
	case 0:
		goto L16
	case 1:
		goto L17
	default:
		goto L18
	}
L13:
	;
	goto L14
L14:
	;
	goto L7
L15:
	;
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v76 = v51 + int32(1)
	goto L15
L17:
	;
	if int32(0) < v51 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	switch v54 - int32(91) {
	case 0:
		goto L16
	default:
		v76 = v51
		goto L15
	case 2:
		goto L17
	}
L19:
	;
	v76 = v51 - int32(1)
	goto L15
L20:
	;
	goto L21
L21:
	;
	goto L3
L22:
	;
	v80 = base.B2i32(v54 != int32(59))
	goto L24
L23:
	;
	v80 = int32(0)
	goto L24
L24:
	;
	if v80 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v76 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+64))
	goto L35
L28:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_read_sql_construct_4))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if l5 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l3
	F_errmsg(m, int32(_a_F_read_sql_construct_5), v19+int32(-48))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v97 = F_plpgsql_scanner_errposition(m, v96, l11)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_read_sql_construct_6), int32(2822), int32(_a_F_read_sql_construct_7))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v50 = v59
	v51 = v76
	v52 = v104 + v106
	goto L6
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v59
	goto L38
L37:
	;
	goto L38
L38:
	;
	if l8 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v54
	goto L41
L40:
	;
	goto L41
L41:
	;
	if v52 <= v59 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if l5 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_plpgsql_append_source_text(m, v19+int32(-36), v59, v52, l11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	F_plpgsql_yyerror(m, l10, int32(0), l11, int32(_a_F_read_sql_construct_8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v123 = F_palloc0(m, int32(80))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v125 = F_pstrdup(m, v121)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v125
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v130
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[2]))
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+20)) = uint8(v134)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v133
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_pfree(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if l6 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	m.G0 = v21 - int32(-64)
	return v123
L52:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_sql_construct[3])))
	if v145 != int32(1) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v150 = int32(_a_F_read_sql_construct_9)
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[4]))
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[4])) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(_a_F_read_sql_construct_10)
	v160 = int32(_a_F_read_sql_construct_11)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[6])) = v19 + int32(-20)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v19 + int32(-8)
	v170 = F_raw_parser(m, v149, v148)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[4])) = v151
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_read_sql_construct[6])) = v175
	goto L51
L55:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v189 = F_plpgsql_scanner_errposition(m, v188, l11)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_read_sql_construct_6), int32(2816), int32(_a_F_read_sql_construct_7))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sql_postrewrite_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v3 = int32(0)
	if l0 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L23
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L19
	}
L3:
	;
	if l1 != 0 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v3
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12+v16<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L3
L8:
	;
	v33 = v16 + int32(1)
	if v9 != v33 {
		v16 = v33
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 != int32(213) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v30 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L7
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+58)))
	v45 = F_check_sql_stmt_retval(m, l0, v41, v42, v43, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	return
L16:
	;
	return
L17:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+56)))
	if v45 != v47 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_sql_postrewrite_callback_0), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_sql_postrewrite_callback_1), int32(2075), int32(_a_F_sql_postrewrite_callback_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_sql_postrewrite_callback_3), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_sql_postrewrite_callback_1), int32(1270), int32(_a_F_sql_postrewrite_callback_4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_unpack_sql_state(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v2 = int32(_a_F_unpack_sql_state_0)
	v3 = int32(63)
	v5 = int32(48)
	v6 = l0&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _c_F_unpack_sql_state[0])) = uint8(v6)
	v14 = int32(base.Ui32(l0)>>(uint(int32(24))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _c_F_unpack_sql_state[1])) = uint8(v14)
	v22 = int32(base.Ui32(l0)>>(uint(int32(18))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _c_F_unpack_sql_state[2])) = uint8(v22)
	v30 = int32(base.Ui32(l0)>>(uint(int32(12))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _c_F_unpack_sql_state[3])) = uint8(v30)
	v38 = int32(base.Ui32(l0)>>(uint(int32(6))%32))&v3 + v5
	*(*uint8)(unsafe.Add(mBase, _c_F_unpack_sql_state[4])) = uint8(v38)
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_unpack_sql_state[5])) = uint8(v41)
	return v2
}
