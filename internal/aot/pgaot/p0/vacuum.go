package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vacuumLeafPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 float64
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v708 int32
	_ = v708
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int64
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(6192)
	m.G0 = v23
	if l2 < v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+12)))
	v49 = F__emscripten_memset_bulkmem(m, v23+int32(448), base.I32_extend8_s(int32(0)), int32(818))
	mBase = m.M
	goto L5
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(l2^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L1
L3:
	;
	goto L4
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v42 = v36 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v55 = F__emscripten_memset_bulkmem(m, v23+int32(32), base.I32_extend8_s(int32(0)), int32(409))
	mBase = m.M
	goto L6
L6:
	;
	if base.Ui32(v43) < base.Ui32(int32(25)) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L19
	} else {
		goto L129
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L19
	} else {
		goto L126
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L19
	} else {
		goto L119
	}
L10:
	;
	m.G0 = v23 + int32(6192)
	return
L11:
	;
	v61 = int32(base.Ui32(v43+int32(262120)) >> (uint(int32(2)) % 32))
	v63 = v61 & int32(65535)
	if v63 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v67 = l0 + int32(100)
	v70 = int32(1)
	v78 = v70
	v80 = v70
	v88 = v5
	goto L13
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32)+(v42+int32(24))-int32(4))))
	v100 = v42 + v97&int32(32767)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	switch v101 & int32(3) {
	case 0:
		goto L17
	case 1:
		goto L16
	default:
		v250 = v88
		goto L15
	}
L14:
	;
	if v250 == int32(0) {
		goto L10
	} else {
		goto L48
	}
L15:
	;
	v255 = v80 + int32(1)
	v257 = v255 & int32(65535)
	if base.Ui32(v257) <= base.Ui32(v63) {
		v78 = v257
		v80 = v255
		v88 = v250
		goto L13
	} else {
		goto L47
	}
L16:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v144))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v143)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v108 = m.T0[v107].(func(*base.Module, int32, int32) int32)(m, v100+int32(6), v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+4)))
	v132 = v130 & int32(16383)
	if v132 == int32(0) {
		v250 = v129
		goto L15
	} else {
		goto L25
	}
L19:
	;
	return
L20:
	;
	if v108 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v110)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v110)+16)) = base.F64_add(v111, float64(1))
	v118 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(32)+v78))) = uint8(v118)
	v129 = v88 + v118
	goto L18
L22:
	;
	goto L23
L23:
	;
	if l3 != 0 {
		v129 = v88
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v122)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v122)+8)) = base.F64_add(v123, float64(1))
	v129 = v88
	goto L18
L25:
	;
	if base.Ui32(v63) < base.Ui32(v132) {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v140 = v23 + int32(448) + v132<<(uint(int32(1))%32)
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140))))
	if v141 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v80)
	v250 = v129
	goto L15
L28:
	;
	if v156 == int32(0) {
		v250 = v88
		goto L15
	} else {
		goto L32
	}
L29:
	;
	v156 = base.B2i32(base.Ui32(v144) <= base.Ui32(v143))
	goto L28
L30:
	;
	goto L31
L31:
	;
	v156 = base.B2i32(int32(0) <= v143-v144)
	goto L28
L32:
	;
	v160 = v100 + int32(6)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	if v161 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v166 = v161
	goto L36
L34:
	;
	v208 = v67
	goto L35
L35:
	;
	v223 = F_palloc(m, int32(12))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L19
	} else {
		goto L46
	}
L36:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+2)))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	v184 = int32(16)
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+2)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166))))
	if v182|v183<<(uint(v184)%32) == v187|v188<<(uint(v184)%32) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v208 = v166 + int32(8)
	goto L35
L38:
	;
	if v198 != 0 {
		v250 = v88
		goto L15
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+4)))
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if v194 == v195 {
		v198 = int32(1)
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v198 = int32(0)
	goto L39
L43:
	;
	goto L42
L44:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v166)+8))
	if v199 != 0 {
		v166 = v199
		goto L36
	} else {
		goto L45
	}
L45:
	;
	goto L37
L46:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v223)+4)) = uint16(v225)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v227
	v229 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v229
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+6)) = uint8(v229)
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v223
	v250 = v88
	goto L15
L47:
	;
	goto L14
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[58]))) = int64(0)
	v264 = v42 + int32(4)
	v266 = v42 + int32(24)
	v267 = int32(1)
	v273 = v267
	v274 = v267
	v278 = int32(0)
	v283 = v5
	v284 = v5
	v287 = v5
	goto L49
L49:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v274<<(uint(int32(2))%32)+v266-int32(4))))
	v297 = v295 & int32(32767)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v297))))
	if v299&int32(3) != 0 {
		v480 = v278
		v485 = v283
		v486 = v284
		v489 = v287
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v499 = int32(65535)
	v500 = v486 & v499
	v502 = v480 & v499
	v504 = v489 & v499
	if v250 != v500+(v502+v504) {
		goto L7
	} else {
		goto L79
	}
L51:
	;
	v493 = v273 + int32(1)
	v494 = int32(65535)
	v495 = v493 & v494
	if base.Ui32(v495) <= base.Ui32(v61&v494) {
		v273 = v493
		v274 = v495
		v278 = v480
		v283 = v485
		v284 = v486
		v287 = v489
		goto L49
	} else {
		goto L78
	}
L52:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(448)+v274<<(uint(int32(1))%32)))))
	if v307 != 0 {
		v480 = v278
		v485 = v283
		v486 = v284
		v489 = v287
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v308 = int32(0)
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(32)+v274))))
	if v313 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v314 = v308
	goto L56
L55:
	;
	v314 = v273
	goto L56
L56:
	;
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297+v264))))
	v318 = v316 & int32(16383)
	if v318 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v465 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(5360)+v287&int32(65535)<<(uint(v465)%32)))) = uint16(v273)
	v470 = v287 + v465
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[58]))) = uint16(v470)
	v480 = v449
	v485 = v454
	v486 = v455
	v489 = v470
	goto L51
L58:
	;
	v323 = v318
	v325 = v308
	v327 = v278
	v331 = v314
	v332 = v283
	v333 = v284
	goto L61
L59:
	;
	goto L60
L60:
	;
	if v314&int32(65535) != 0 {
		v480 = v278
		v485 = v283
		v486 = v284
		v489 = v287
		goto L51
	} else {
		goto L77
	}
L61:
	;
	v340 = v323 & int32(65535)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340<<(uint(int32(2))%32)+v266-int32(4))))
	v348 = v346 & int32(32767)
	v349 = v42 + v348
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v350&int32(3) != 0 {
		goto L8
	} else {
		goto L63
	}
L62:
	;
	if v410&int32(65535) == int32(0) {
		v449 = v408
		v454 = v411
		v455 = v412
		goto L57
	} else {
		goto L75
	}
L63:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+int32(32)+v340))))
	if v356 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348+v264))))
	v416 = v414 & int32(16383)
	if v416 != 0 {
		v323 = v416
		v325 = v356
		v327 = v408
		v331 = v410
		v332 = v411
		v333 = v412
		goto L61
	} else {
		goto L74
	}
L65:
	;
	v363 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(4544)+v327&int32(65535)<<(uint(v363)%32)))) = uint16(v323)
	v368 = v327 + v363
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[59]))) = uint16(v368)
	v408 = v368
	v410 = v331
	v411 = v332
	v412 = v333
	goto L64
L66:
	;
	goto L67
L67:
	;
	if v331&int32(65535) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v374 = int32(1)
	v377 = v333 << (uint(v374) % 32) & int32(131070)
	*(*uint16)(unsafe.Add(mBase, uint32(v377+(v23+int32(2912))))) = uint16(v273)
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(3728)+v377))) = uint16(v323)
	v387 = v333 + v374
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[60]))) = uint16(v387)
	v408 = v327
	v410 = v273
	v411 = v332
	v412 = v387
	goto L64
L69:
	;
	goto L70
L70:
	;
	if v325&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v391 = int32(1)
	v394 = v332 << (uint(v391) % 32) & int32(131070)
	*(*uint16)(unsafe.Add(mBase, uint32(v394+(v23+int32(1280))))) = uint16(v323)
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(2096)+v394))) = uint16(v331)
	v404 = v332 + v391
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[61]))) = uint16(v404)
	v407 = v404
	goto L73
L72:
	;
	v407 = v332
	goto L73
L73:
	;
	v408 = v327
	v410 = v323
	v411 = v407
	v412 = v333
	goto L64
L74:
	;
	goto L62
L75:
	;
	if v356 == int32(0) {
		v480 = v408
		v485 = v411
		v486 = v412
		v489 = v287
		goto L51
	} else {
		goto L76
	}
L76:
	;
	v423 = int32(1)
	v426 = v411 << (uint(v423) % 32) & int32(131070)
	v430 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v426+(v23+int32(1280))))) = uint16(v430)
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(2096)+v426))) = uint16(v410)
	v437 = v411 + v423
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[61]))) = uint16(v437)
	v480 = v408
	v485 = v437
	v486 = v412
	v489 = v287
	goto L51
L77:
	;
	v449 = v278
	v454 = v283
	v455 = v284
	goto L57
L78:
	;
	goto L50
L79:
	;
	v508 = int32(4509780)
	v510 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v510 + int32(1)
	v515 = l0 + int32(16)
	v518 = int32(2)
	F_spgPageIndexMultiDelete(m, v515, v42, v23+int32(5360), v504, v518, v518, int32(-1), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	v526 = int32(3)
	F_spgPageIndexMultiDelete(m, v515, v42, v23+int32(4544), v502, v526, v526, int32(-1), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L19
	} else {
		goto L81
	}
L81:
	;
	if v500 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v535 = int32(1)
	v538 = v42 + int32(24)
	v539 = int32(0)
	if v486&int32(65535) != v535 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v690 = int32(0)
	goto L84
L84:
	;
	v691 = int32(3)
	F_spgPageIndexMultiDelete(m, v515, v42, v23+int32(3728), v690, v691, v691, int32(-1), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L19
	} else {
		goto L94
	}
L85:
	;
	v553 = v539
	v555 = int32(0)
	goto L88
L86:
	;
	v626 = v539
	goto L87
L87:
	;
	if v500&v535 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v568 = v553 << (uint(int32(1)) % 32)
	v570 = v23 + int32(3728)
	v572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v568+v570))))
	v573 = int32(2)
	v576 = int32(4)
	v577 = v572<<(uint(v573)%32) + v538 - v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	v580 = v23 + int32(2912)
	v582 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580+v568))))
	v587 = v582<<(uint(v573)%32) + v538 - v576
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	*(*int32)(unsafe.Add(mBase, uint32(v577))) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v578
	v592 = v568 | v573
	v596 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v592+v570))))
	v601 = v596<<(uint(v573)%32) + v538 - v576
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v606 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v580+v592))))
	v611 = v606<<(uint(v573)%32) + v538 - v576
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v612
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v602
	v616 = v553 + v573
	v618 = v555 + v573
	if v618 != v500&int32(65534) {
		v553 = v616
		v555 = v618
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v626 = v616
	goto L87
L90:
	;
	goto L89
L91:
	;
	v641 = v626 << (uint(int32(1)) % 32)
	v645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v641+(v23+int32(3728))))))
	v646 = int32(2)
	v649 = int32(4)
	v650 = v645<<(uint(v646)%32) + v538 - v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v650)))
	v655 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(2912)+v641))))
	v660 = v655<<(uint(v646)%32) + v538 - v649
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	*(*int32)(unsafe.Add(mBase, uint32(v650))) = v661
	*(*int32)(unsafe.Add(mBase, uint32(v660))) = v651
	goto L93
L92:
	;
	goto L93
L93:
	;
	v690 = v486 & int32(65535)
	goto L84
L94:
	;
	v698 = v485 & int32(65535)
	if v698 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v708 = int32(0)
	goto L98
L96:
	;
	goto L97
L97:
	;
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L19
	} else {
		goto L101
	}
L98:
	;
	v724 = int32(1)
	v725 = v708 << (uint(v724) % 32)
	v729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v725+(v23+int32(2096))))))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v729<<(uint(int32(2))%32)+(v42+int32(24))-int32(4))))
	v738 = v42 + int32(4) + v735&int32(32767)
	v742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(1280)+v725))))
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v738))))
	v748 = v742&int32(16383) | v745&int32(49152)
	*(*uint16)(unsafe.Add(mBase, uint32(v738))) = uint16(v748)
	v751 = v708 + v724
	if base.Ui32(v751) < base.Ui32(v698) {
		v708 = v751
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	goto L99
L101:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+118)))
	if v776 != int32(112) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v849 = int32(4509780)
	v851 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v851 - int32(1)
	goto L10
L103:
	;
	v780 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v780 <= int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v783 != 0 {
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L19
	} else {
		goto L109
	}
L107:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v784 != 0 {
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[62]))) = v787
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[63]))) = uint8(v789)
	F_XLogRegisterData(m, v23+int32(6176), int32(16))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[58]))))
	F_XLogRegisterData(m, v23+int32(5360), v798<<(uint(int32(1))%32))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	v805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[59]))))
	F_XLogRegisterData(m, v23+int32(4544), v805<<(uint(int32(1))%32))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	v812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[60]))))
	F_XLogRegisterData(m, v23+int32(3728), v812<<(uint(int32(1))%32))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L19
	} else {
		goto L113
	}
L113:
	;
	v819 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[60]))))
	F_XLogRegisterData(m, v23+int32(2912), v819<<(uint(int32(1))%32))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	v826 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[61]))))
	F_XLogRegisterData(m, v23+int32(2096), v826<<(uint(int32(1))%32))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	v833 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[61]))))
	F_XLogRegisterData(m, v23+int32(1280), v833<<(uint(int32(1))%32))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	v844 = F_XLogInsert(m, int32(16), int32(96))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L19
	} else {
		goto L118
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = base.I64_rotr(v844, int64(32))
	goto L102
L119:
	;
	if l2 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v902 + int32(4)
	F_errmsg_internal(m, int32(695719), v23+int32(16))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L19
	} else {
		goto L124
	}
L121:
	;
	v886 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v886+(l2^int32(-1))<<(uint(int32(6))%32))+16))
	v901 = v892
	goto L120
L122:
	;
	goto L123
L123:
	;
	v894 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v894+l2<<(uint(int32(6))%32)+int32(-64))+16))
	v901 = v900
	goto L120
L124:
	;
	F_errfinish(m, int32(496764), int32(179), int32(408773))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L19
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
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v921 & int32(3)
	F_errmsg_internal(m, int32(483626), v23)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L19
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(496764), int32(266), int32(408773))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L19
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errmsg_internal(m, int32(164567), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L19
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(496764), int32(326), int32(408773))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vacuum_delay_point(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v82 float64
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 float64
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 float64
	_ = v106
	var v115 float64
	_ = v115
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v145 float64
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int64
	_ = v180
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
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
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v354 int32
	_ = v354
	v7 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v21 != 0 {
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
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L83
	}
L7:
	;
	m.G0 = v14 + int32(16)
	return
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[304])))
	if v23 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	if v27 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v33 = *(*int32)(unsafe.Add(mBase, _consts[305]))
	if v33 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	if v53 != 0 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	if v23 == int32(0) {
		goto L7
	} else {
		goto L20
	}
L15:
	;
	if v31 != int32(4) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[305])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[304])))
	if v47&int32(1) != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	goto L13
L21:
	;
	if base.F64_gt(v115, float64(0)) == int32(0) {
		goto L7
	} else {
		goto L29
	}
L22:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[307]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v58 = int32(4509840)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	v60 = v57 + v59
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v60
	v62 = int32(4412300)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[309]))
	v66 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	v67 = v64 + v66
	*(*int32)(unsafe.Add(mBase, _consts[309])) = v67
	v70 = *(*int32)(unsafe.Add(mBase, _consts[310]))
	if base.Ui32(v60) < base.Ui32(v70) {
		v94 = v7
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	v103 = *(*int32)(unsafe.Add(mBase, _consts[310]))
	if v101 < v103 {
		goto L7
	} else {
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[308])) = int32(0)
	v115 = v94
	goto L21
L26:
	;
	v72 = base.F64_convert_i32_s(v67)
	v73 = base.F64_convert_i32_s(v70)
	if base.F64_gt(v72, base.F64_mul(base.F64_div(v73, base.F64_convert_i32_s(v56)), float64(0.5))) == int32(0) {
		v94 = v7
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v82 = *(*float64)(unsafe.Add(mBase, _consts[311]))
	v84 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v85 - v67
	*(*int32)(unsafe.Add(mBase, _consts[309])) = int32(0)
	v94 = base.F64_div(base.F64_mul(v82, v72), v73)
	goto L25
L28:
	;
	v106 = *(*float64)(unsafe.Add(mBase, _consts[311]))
	v115 = base.F64_div(base.F64_mul(v106, base.F64_convert_i32_s(v101)), base.F64_convert_i32_s(v103))
	goto L21
L29:
	;
	v123 = *(*float64)(unsafe.Add(mBase, _consts[311]))
	v125 = base.F64_mul(v123, float64(4))
	if base.F64_gt(v115, v125) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v127 = v125
	goto L32
L31:
	;
	v127 = v115
	goto L32
L32:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[312])))
	if v129 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v134 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v139 = v134 + v135*int64(1000000000)
	goto L35
L34:
	;
	v139 = int64(0)
	goto L35
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = int32(150994951)
	v145 = base.F64_mul(v127, float64(1000))
	if base.F64_lt(base.F64_abs(v145), float64(2.147483648e+09)) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_pg_usleep(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L40
	}
L37:
	;
	v149 = base.I32_trunc_f64_s(v145)
	v151 = v149
	goto L36
L38:
	;
	goto L39
L39:
	;
	v151 = int32(-2147483648)
	goto L36
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(0)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _consts[312])))
	if v159 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v273 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L42:
	;
	F___clock_gettime(m, int32(1), v14)
	mBase = m.M
	v164 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+8)))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v168 = v164 + v165*int64(1000000000)
	v169 = v168 - v139
	v171 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if int32(0) <= v171 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v174 = int32(4412280)
	v176 = *(*int64)(unsafe.Add(mBase, _consts[319]))
	v177 = v176 + v169
	*(*int64)(unsafe.Add(mBase, _consts[319])) = v177
	v180 = *(*int64)(unsafe.Add(mBase, _consts[320]))
	if v168-v180 < int64(1000000000) {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if l0 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	F_pgstat_progress_parallel_incr_param(m, int32(10), v177)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, _consts[319])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[320])) = v168
	goto L41
L48:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v195 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	if v234 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	goto L41
L52:
	;
	goto L51
L53:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v199 != int32(1) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v202 = int32(4509780)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v205 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v204 + v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v208 + v205
	v216 = v195 + int32(296)
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v216)))
	*(*int64)(unsafe.Add(mBase, uint32(v216))) = v217 + v169
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v220 + v205
	v226 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v226 - v205
	goto L52
L55:
	;
	goto L41
L56:
	;
	goto L55
L57:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
	if v238 != int32(1) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v241 = int32(4509780)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v244 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v243 + v244
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v247 + v244
	v255 = v234 + int32(312)
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
	*(*int64)(unsafe.Add(mBase, uint32(v255))) = v256 + v169
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v259 + v244
	v265 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v265 - v244
	goto L56
L59:
	;
	v276 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v281 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[308])) = v281
	v284 = *(*int32)(unsafe.Add(mBase, _consts[313]))
	if v284 == v281 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v276 == int32(0) {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v334 == int32(0) {
		goto L7
	} else {
		goto L81
	}
L65:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _consts[314]))
	if v289 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L78
	}
L67:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _consts[315]))
	v296 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if int32(0) < v294 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v316 = v289
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, _consts[310])) = v316
	goto L64
L70:
	;
	v299 = v294
	goto L72
L71:
	;
	v299 = v296
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, _consts[310])) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v284)+32))
	if v301 == int32(0) {
		goto L64
	} else {
		goto L73
	}
L73:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+uint32(_consts[318])))
	if v306 <= int32(0) {
		goto L66
	} else {
		goto L74
	}
L74:
	;
	v309 = int32(1)
	v310 = base.I32_div_s(v299, v306)
	if v310 <= v309 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v313 = v309
	goto L77
L76:
	;
	v313 = v310
	goto L77
L77:
	;
	v316 = v313
	goto L69
L78:
	;
	F_errmsg_internal(m, int32(570146), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(496739), int32(1754), int32(101080))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	goto L7
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
