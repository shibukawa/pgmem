package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cdissect(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
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
	var v933 int32
	_ = v933
	var v942 int32
	_ = v942
	v5 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_cdissect[0]))
	if v15 != 0 {
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = m.T0[v22].(func(*base.Module) int32)(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(19)
L8:
	;
	goto L9
L9:
	;
	v27 = int32(15)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v28 - int32(40) {
	case 0:
		goto L14
	case 1, 3, 4, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		v942 = v27
		goto L10
	case 2:
		goto L15
	case 6:
		goto L17
	case 21:
		v907 = v5
		goto L11
	default:
		goto L18
	}
L10:
	;
	return v942
L11:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v915 <= int32(0) {
		goto L307
	} else {
		goto L308
	}
L12:
	;
	F_pfree(m, v416)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L4
	} else {
		goto L306
	}
L13:
	;
	v634 = (l3 - l2) >> (uint(int32(2)) % 32)
	v635 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v634) < base.Ui32(v635) {
		goto L230
	} else {
		goto L231
	}
L14:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v629 = F_cdissect(m, l0, v628, l2, l3)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L4
	} else {
		goto L229
	}
L15:
	;
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+1)))
	if v390&int32(2) == int32(0) {
		goto L13
	} else {
		goto L151
	}
L16:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v323 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v101 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v100<<(uint(v101)%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)+24))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v106&v101 != 0 {
		goto L37
	} else {
		goto L38
	}
L18:
	;
	if v28 == int32(124) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if v28 != int32(98) {
		v942 = v27
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v35 = int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = v36 + v37<<(uint(int32(3))%32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 == int32(-1) {
		v907 = v35
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v41 == v46 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v907 = base.B2i32(l2 != l3) | base.B2i32(base.I32_extend16_s(v44) < base.I32_extend16_s(v45))
	goto L11
L23:
	;
	goto L24
L24:
	;
	if l2 == l3 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v907 = base.B2i32(v45 != int32(0))
	goto L11
L26:
	;
	goto L27
L27:
	;
	v58 = (l3 - l2) >> (uint(int32(2)) % 32)
	v59 = v46 - v41
	v60 = base.I32_div_u_s(v58, v59)
	if v58-v60*v59 != 0 {
		v907 = v35
		goto L11
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v60) < base.Ui32(v45) {
		v907 = v35
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(v44 != int32(256))&base.B2i32(base.Ui32(v44) < base.Ui32(v60)) != 0 {
		v907 = v35
		goto L11
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v58) < base.Ui32(v59) {
		v907 = int32(0)
		goto L11
	} else {
		goto L31
	}
L31:
	;
	v71 = int32(2)
	v80 = l2
	v82 = v60
	goto L32
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+420))
	v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v64+v41<<(uint(v71)%32), v80, v59)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v907 = v94
	goto L11
L34:
	;
	v94 = base.B2i32(v91 != int32(0))
	if v91 != 0 {
		v907 = v94
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v97 = v82 - int32(1)
	if v97 != 0 {
		v80 = v80 + v59<<(uint(v71)%32)
		v82 = v97
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	if v104 != 0 {
		v136 = v104
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v104 != 0 {
		v244 = v104
		goto L85
	} else {
		goto L86
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v137 != 0 {
		v942 = v137
		goto L10
	} else {
		goto L47
	}
L41:
	;
	v109 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v116 = F_newdfa(m, l0, v99+int32(36), v112+int32(72), v109)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	if v116 == int32(0) {
		v136 = v109
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v120 == int32(98) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+60)) = v123
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+64)) = uint16(v125)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+66)) = uint16(v127)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v130<<(uint(int32(2))%32)))) = v116
	v136 = v116
	goto L40
L47:
	;
	v138 = F_getsubdfa(m, l0, v105)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v140 != 0 {
		v942 = v140
		goto L10
	} else {
		goto L49
	}
L49:
	;
	v141 = int32(0)
	v143 = F_shortest(m, l0, v136, l2, l2, l3, v141, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v145 != 0 {
		v942 = v145
		goto L10
	} else {
		goto L51
	}
L51:
	;
	if v143 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	return int32(1)
L53:
	;
	goto L54
L54:
	;
	v157 = v143
	goto L55
L55:
	;
	v164 = F_longest(m, l0, v138, v157, l3, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L57
	}
L56:
	;
	v942 = int32(1)
	goto L10
L57:
	;
	if v164 == l3 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v167 = F_cdissect(m, l0, v99, l2, v157)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v206 != 0 {
		v942 = v206
		goto L10
	} else {
		goto L80
	}
L61:
	;
	if v167 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v171 = F_cdissect(m, l0, v105, v157, l3)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	v202 = v167
	goto L64
L64:
	;
	if v202 != int32(1) {
		v942 = v202
		goto L10
	} else {
		goto L79
	}
L65:
	;
	if v171 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v907 = int32(0)
	goto L11
L67:
	;
	goto L68
L68:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v177 <= int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v202 = v171
	goto L64
L70:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v193 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v180) <= base.Ui32(v177) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v183 = v177 << (uint(int32(3)) % 32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v186 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v188+v183)+4)) = v186
	goto L70
L73:
	;
	v195 = v193
	goto L76
L74:
	;
	goto L75
L75:
	;
	goto L69
L76:
	;
	F_zaptreesubs(m, l0, v195)
	mBase = m.M
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v195)+24))
	if v198 != 0 {
		v195 = v198
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	goto L77
L79:
	;
	goto L60
L80:
	;
	if l3 == v157 {
		v942 = int32(1)
		goto L10
	} else {
		goto L81
	}
L81:
	;
	v211 = int32(0)
	v213 = F_shortest(m, l0, v136, l2, v157+int32(4), l3, v211, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v215 != 0 {
		v942 = v215
		goto L10
	} else {
		goto L83
	}
L83:
	;
	if v213 != 0 {
		v157 = v213
		goto L55
	} else {
		goto L84
	}
L84:
	;
	goto L56
L85:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v245 != 0 {
		v942 = v245
		goto L10
	} else {
		goto L92
	}
L86:
	;
	v217 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = F_newdfa(m, l0, v99+int32(36), v220+int32(72), v217)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	if v224 == int32(0) {
		v244 = v217
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v228 == int32(98) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+60)) = v231
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+64)) = uint16(v233)
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+66)) = uint16(v235)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v237+v238<<(uint(int32(2))%32)))) = v224
	v244 = v224
	goto L85
L92:
	;
	v246 = F_getsubdfa(m, l0, v105)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v248 != 0 {
		v942 = v248
		goto L10
	} else {
		goto L94
	}
L94:
	;
	v250 = F_longest(m, l0, v244, l2, l3, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v252 != 0 {
		v942 = v252
		goto L10
	} else {
		goto L96
	}
L96:
	;
	if v250 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	return int32(1)
L98:
	;
	goto L99
L99:
	;
	v264 = v250
	goto L100
L100:
	;
	v271 = F_longest(m, l0, v246, v264, l3, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	v942 = int32(1)
	goto L10
L102:
	;
	if v271 == l3 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v274 = F_cdissect(m, l0, v99, l2, v264)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v313 != 0 {
		v942 = v313
		goto L10
	} else {
		goto L125
	}
L106:
	;
	if v274 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v278 = F_cdissect(m, l0, v105, v264, l3)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	v309 = v274
	goto L109
L109:
	;
	if v309 != int32(1) {
		v942 = v309
		goto L10
	} else {
		goto L124
	}
L110:
	;
	if v278 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v907 = int32(0)
	goto L11
L112:
	;
	goto L113
L113:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v284 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v309 = v278
	goto L109
L115:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	if v300 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v287) <= base.Ui32(v284) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v290 = v284 << (uint(int32(3)) % 32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v293 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v290+v291))) = v293
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v295+v290)+4)) = v293
	goto L115
L118:
	;
	v302 = v300
	goto L121
L119:
	;
	goto L120
L120:
	;
	goto L114
L121:
	;
	F_zaptreesubs(m, l0, v302)
	mBase = m.M
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v302)+24))
	if v305 != 0 {
		v302 = v305
		goto L121
	} else {
		goto L123
	}
L122:
	;
	goto L120
L123:
	;
	goto L122
L124:
	;
	goto L105
L125:
	;
	if l2 == v264 {
		v942 = int32(1)
		goto L10
	} else {
		goto L126
	}
L126:
	;
	v319 = F_longest(m, l0, v244, l2, v264-int32(4), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v321 != 0 {
		v942 = v321
		goto L10
	} else {
		goto L128
	}
L128:
	;
	if v319 != 0 {
		v264 = v319
		goto L100
	} else {
		goto L129
	}
L129:
	;
	goto L101
L130:
	;
	return int32(1)
L131:
	;
	goto L132
L132:
	;
	v334 = v323
	goto L133
L133:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341+v342<<(uint(int32(2))%32))))
	if v346 != 0 {
		v374 = v346
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v942 = int32(1)
	goto L10
L135:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v375 != 0 {
		v942 = v375
		goto L10
	} else {
		goto L142
	}
L136:
	;
	v347 = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v354 = F_newdfa(m, l0, v334+int32(36), v350+int32(72), v347)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v354 == int32(0) {
		v374 = v347
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v358 == int32(98) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+60)) = v361
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v354)+64)) = uint16(v363)
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v354)+66)) = uint16(v365)
	goto L141
L140:
	;
	goto L141
L141:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v367+v368<<(uint(int32(2))%32)))) = v354
	v374 = v354
	goto L135
L142:
	;
	v377 = F_longest(m, l0, v374, l2, l3, int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	if v377 == l3 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v380 = F_cdissect(m, l0, v334, l2, l3)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v385 != 0 {
		v942 = v385
		goto L10
	} else {
		goto L149
	}
L147:
	;
	if v380 != int32(1) {
		v907 = v380
		goto L11
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v334)+24))
	if v387 != 0 {
		v334 = v387
		goto L133
	} else {
		goto L150
	}
L150:
	;
	goto L134
L151:
	;
	if v388 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if l2 == l3 {
		v907 = v5
		goto L11
	} else {
		goto L155
	}
L153:
	;
	v399 = v388
	goto L154
L154:
	;
	v402 = (l3 - l2) >> (uint(int32(2)) % 32)
	v403 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v402) < base.Ui32(v403) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v399 = int32(1)
	goto L154
L156:
	;
	v405 = v402
	goto L158
L157:
	;
	v405 = v403
	goto L158
L158:
	;
	if v403 == int32(256) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v408 = v402
	goto L161
L160:
	;
	v408 = v405
	goto L161
L161:
	;
	if base.Ui32(v399) < base.Ui32(v408) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v410 = v408
	goto L164
L163:
	;
	v410 = v399
	goto L164
L164:
	;
	v411 = int32(2)
	v416 = F_palloc_extended(m, v410<<(uint(v411)%32)+int32(4), v411)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	if v416 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	return int32(12)
L167:
	;
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = l2
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v424 = F_getsubdfa(m, l0, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v426 != 0 {
		goto L12
	} else {
		goto L170
	}
L170:
	;
	v432 = int32(1)
	v433 = l2
	v435 = v5
	goto L171
L171:
	;
	v442 = v432 - int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v416+v442<<(uint(int32(2))%32))))
	if l3 == v433 {
		v457 = v433
		goto L173
	} else {
		goto L174
	}
L172:
	;
	F_pfree(m, v416)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L228
	}
L173:
	;
	if base.Ui32(v432) < base.Ui32(v410) {
		goto L180
	} else {
		goto L181
	}
L174:
	;
	if v433 != v446 {
		v457 = v433
		goto L173
	} else {
		goto L175
	}
L175:
	;
	if base.Ui32(v432) < base.Ui32(v399) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	if (l3-v433)>>(uint(int32(2))%32) <= v399-v432 {
		v457 = v433
		goto L173
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v457 = v433 + int32(4)
	goto L173
L179:
	;
	goto L178
L180:
	;
	v462 = v457
	goto L182
L181:
	;
	v462 = l3
	goto L182
L182:
	;
	v463 = int32(0)
	v465 = F_shortest(m, l0, v424, v446, v462, l3, v463, v463)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416+v432<<(uint(int32(2))%32)))) = v465
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v468 != 0 {
		goto L12
	} else {
		goto L184
	}
L184:
	;
	if v465 == int32(0) {
		v542 = v435
		goto L189
	} else {
		goto L190
	}
L185:
	;
	goto L172
L186:
	;
	if int32(0) < v600 {
		v432 = v600
		v433 = v601
		v435 = v603
		goto L171
	} else {
		goto L227
	}
L187:
	;
	v576 = v563
	goto L221
L188:
	;
	if v556 <= int32(0) {
		goto L185
	} else {
		goto L220
	}
L189:
	;
	v552 = v542
	v556 = v442
	goto L188
L190:
	;
	if v435 < v432 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v472 = v435
	goto L193
L192:
	;
	v472 = v442
	goto L193
L193:
	;
	if l3 != v465 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	if base.Ui32(v410) <= base.Ui32(v432) {
		v542 = v472
		goto L189
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	if base.Ui32(v432) < base.Ui32(v399) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v600 = v432 + int32(1)
	v601 = v465
	v603 = v472
	goto L186
L198:
	;
	v563 = v432
	v568 = v472
	goto L187
L199:
	;
	goto L200
L200:
	;
	v485 = v472
	goto L202
L201:
	;
	F_pfree(m, v416)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L219
	}
L202:
	;
	v492 = v485 + int32(1)
	if v432 < v492 {
		goto L201
	} else {
		goto L204
	}
L203:
	;
	if v530 == int32(1) {
		v552 = v485
		v556 = v492
		goto L188
	} else {
		goto L217
	}
L204:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if v496 <= int32(0) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v522 = int32(2)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v416+v485<<(uint(v522)%32))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v416+v492<<(uint(v522)%32))))
	v530 = F_cdissect(m, l0, v521, v525, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L215
	}
L206:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	if v512 != 0 {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v499) <= base.Ui32(v496) {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v502 = v496 << (uint(int32(3)) % 32)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v505 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v502+v503))) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v507+v502)+4)) = v505
	goto L206
L209:
	;
	v514 = v512
	goto L212
L210:
	;
	goto L211
L211:
	;
	goto L205
L212:
	;
	F_zaptreesubs(m, l0, v514)
	mBase = m.M
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	if v517 != 0 {
		v514 = v517
		goto L212
	} else {
		goto L214
	}
L213:
	;
	goto L211
L214:
	;
	goto L213
L215:
	;
	if v530 == int32(0) {
		v485 = v492
		goto L202
	} else {
		goto L216
	}
L216:
	;
	goto L203
L217:
	;
	F_pfree(m, v416)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	return v530
L219:
	;
	v907 = int32(0)
	goto L11
L220:
	;
	v563 = v556
	v568 = v552
	goto L187
L221:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v416+v576<<(uint(int32(2))%32))))
	if base.Ui32(v588) < base.Ui32(l3) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	goto L185
L223:
	;
	v600 = v576
	v601 = v588 + int32(4)
	v603 = v568
	goto L186
L224:
	;
	goto L225
L225:
	;
	v592 = int32(1)
	if v592 < v576 {
		v576 = v576 - v592
		goto L221
	} else {
		goto L226
	}
L226:
	;
	goto L222
L227:
	;
	goto L185
L228:
	;
	return int32(1)
L229:
	;
	v907 = v629
	goto L11
L230:
	;
	v637 = v634
	goto L232
L231:
	;
	v637 = v635
	goto L232
L232:
	;
	if v635 == int32(256) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v640 = v634
	goto L235
L234:
	;
	v640 = v637
	goto L235
L235:
	;
	v641 = int32(1)
	if v388 <= v641 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v644 = v641
	goto L238
L237:
	;
	v644 = v388
	goto L238
L238:
	;
	if base.Ui32(v644) < base.Ui32(v640) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v646 = v640
	goto L241
L240:
	;
	v646 = v644
	goto L241
L241:
	;
	v647 = int32(2)
	v652 = F_palloc_extended(m, v646<<(uint(v647)%32)+int32(4), v647)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	if v652 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	return int32(12)
L244:
	;
	goto L245
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652))) = l2
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v660 = F_getsubdfa(m, l0, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v662 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	F_pfree(m, v652)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L4
	} else {
		goto L305
	}
L248:
	;
	v667 = int32(1)
	v671 = l3
	v672 = v5
	goto L249
L249:
	;
	v676 = int32(2)
	v680 = v667 - int32(1)
	v683 = v652 + v680<<(uint(v676)%32)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	v686 = F_longest(m, l0, v660, v684, v671, int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L4
	} else {
		goto L251
	}
L250:
	;
	F_pfree(m, v652)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L4
	} else {
		goto L304
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652+v667<<(uint(v676)%32)))) = v686
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v689 != 0 {
		goto L247
	} else {
		goto L252
	}
L252:
	;
	if v686 == int32(0) {
		v771 = v672
		goto L257
	} else {
		goto L258
	}
L253:
	;
	goto L250
L254:
	;
	if int32(0) < v839 {
		v667 = v839
		v671 = v843
		v672 = v844
		goto L249
	} else {
		goto L303
	}
L255:
	;
	v805 = v792
	goto L293
L256:
	;
	if v777 <= int32(0) {
		goto L253
	} else {
		goto L292
	}
L257:
	;
	v777 = v680
	v782 = v771
	goto L256
L258:
	;
	if v672 < v667 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v693 = v672
	goto L261
L260:
	;
	v693 = v680
	goto L261
L261:
	;
	if l3 != v686 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v792 = v667
	v797 = v693
	goto L255
L263:
	;
	if base.Ui32(v646) <= base.Ui32(v667) {
		v771 = v693
		goto L257
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	if base.Ui32(v667) < base.Ui32(v644) {
		goto L262
	} else {
		goto L272
	}
L266:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	if v696 == v686 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	if base.Ui32(v644) <= base.Ui32(v667) {
		goto L262
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v839 = v667 + int32(1)
	v843 = l3
	v844 = v693
	goto L254
L270:
	;
	if v644-v667 < (l3-v686)>>(uint(int32(2))%32) {
		goto L262
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	v716 = v693
	goto L274
L273:
	;
	F_pfree(m, v652)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L291
	}
L274:
	;
	v721 = v716 + int32(1)
	if v667 < v721 {
		goto L273
	} else {
		goto L276
	}
L275:
	;
	if v759 == int32(1) {
		v777 = v721
		v782 = v716
		goto L256
	} else {
		goto L289
	}
L276:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	if v725 <= int32(0) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v751 = int32(2)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v652+v716<<(uint(v751)%32))))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v652+v721<<(uint(v751)%32))))
	v759 = F_cdissect(m, l0, v750, v754, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L4
	} else {
		goto L287
	}
L278:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	if v741 != 0 {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v728) <= base.Ui32(v725) {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v731 = v725 << (uint(int32(3)) % 32)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v734 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v731+v732))) = v734
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v736+v731)+4)) = v734
	goto L278
L281:
	;
	v743 = v741
	goto L284
L282:
	;
	goto L283
L283:
	;
	goto L277
L284:
	;
	F_zaptreesubs(m, l0, v743)
	mBase = m.M
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v743)+24))
	if v746 != 0 {
		v743 = v746
		goto L284
	} else {
		goto L286
	}
L285:
	;
	goto L283
L286:
	;
	goto L285
L287:
	;
	if v759 == int32(0) {
		v716 = v721
		goto L274
	} else {
		goto L288
	}
L288:
	;
	goto L275
L289:
	;
	F_pfree(m, v652)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L290
	}
L290:
	;
	return v759
L291:
	;
	v907 = int32(0)
	goto L11
L292:
	;
	v792 = v777
	v797 = v782
	goto L255
L293:
	;
	v815 = v652 + v805<<(uint(int32(2))%32)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815-int32(4))))
	if base.Ui32(v816) <= base.Ui32(v819) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	goto L253
L295:
	;
	v831 = int32(1)
	if v831 < v805 {
		v805 = v805 - v831
		goto L293
	} else {
		goto L302
	}
L296:
	;
	v822 = v816 - int32(4)
	if base.Ui32(v819) < base.Ui32(v822) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v839 = v805
	v843 = v822
	v844 = v797
	goto L254
L298:
	;
	goto L299
L299:
	;
	if base.Ui32(v644) <= base.Ui32(v805) {
		goto L295
	} else {
		goto L300
	}
L300:
	;
	if v644-v805 < (l3-v819)>>(uint(int32(2))%32) {
		goto L295
	} else {
		goto L301
	}
L301:
	;
	v839 = v805
	v843 = v822
	v844 = v797
	goto L254
L302:
	;
	goto L294
L303:
	;
	goto L253
L304:
	;
	v866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v907 = base.B2i32(l2 != l3) | base.B2i32(v866 != int32(0))
	goto L11
L305:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v907 = v885
	goto L11
L306:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v907 = v901
	goto L11
L307:
	;
	return v907
L308:
	;
	goto L309
L309:
	;
	if v907 != 0 {
		v942 = v907
		goto L10
	} else {
		goto L310
	}
L310:
	;
	v919 = int32(0)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v920) <= base.Ui32(v915) {
		v942 = v919
		goto L10
	} else {
		goto L311
	}
L311:
	;
	v923 = v915 << (uint(int32(3)) % 32)
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v928 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v923+v924))) = (l2 - v926) >> (uint(v928) % 32)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v931+v923)+4)) = (l3 - v933) >> (uint(v928) % 32)
	v942 = v919
	goto L10
}
