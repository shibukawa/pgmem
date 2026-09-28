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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
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
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
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
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
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
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v713 int32
	_ = v713
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
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v804 int32
	_ = v804
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v832 int32
	_ = v832
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v944 int32
	_ = v944
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
		v944 = v27
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
	return v944
L11:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v916 <= int32(0) {
		goto L297
	} else {
		goto L298
	}
L12:
	;
	F_pfree(m, v414)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L4
	} else {
		goto L296
	}
L13:
	;
	v635 = (l3 - l2) >> (uint(int32(2)) % 32)
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v635) < base.Ui32(v636) {
		goto L225
	} else {
		goto L226
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
		goto L224
	}
L15:
	;
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+1)))
	if v389&int32(2) == int32(0) {
		goto L13
	} else {
		goto L150
	}
L16:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v322 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = int32(2)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+v99<<(uint(v100)%32))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)+24))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v105&v100 != 0 {
		goto L36
	} else {
		goto L37
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
		v944 = v27
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
	v907 = base.B2i32(l2 != l3) | base.B2i32(v44 < v45)
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
	v56 = (l3 - l2) >> (uint(int32(2)) % 32)
	v57 = v46 - v41
	v58 = base.I32_div_u_s(v56, v57)
	if v56-v58*v57|base.B2i32(base.Ui32(v58) < base.Ui32(v45)) != 0 {
		v907 = v35
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(v44 != int32(256))&base.B2i32(base.Ui32(v44) < base.Ui32(v58)) != 0 {
		v907 = v35
		goto L11
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v56) < base.Ui32(v57) {
		v907 = int32(0)
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v70 = int32(2)
	v80 = l2
	v81 = v58
	goto L31
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+420))
	v90 = m.T0[v89].(func(*base.Module, int32, int32, int32) int32)(m, v63+v41<<(uint(v70)%32), v80, v57)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	v907 = v93
	goto L11
L33:
	;
	v93 = base.B2i32(v90 != int32(0))
	if v90 != 0 {
		v907 = v93
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v96 = v81 - int32(1)
	if v96 != 0 {
		v80 = v80 + v57<<(uint(v70)%32)
		v81 = v96
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v103 != 0 {
		v135 = v103
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v103 != 0 {
		v243 = v103
		goto L84
	} else {
		goto L85
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v136 != 0 {
		v944 = v136
		goto L10
	} else {
		goto L46
	}
L40:
	;
	v108 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = F_newdfa(m, l0, v98+int32(36), v111+int32(72), v108)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v115 == int32(0) {
		v135 = v108
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v119 == int32(98) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+60)) = v122
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+64)) = uint16(v124)
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+66)) = uint16(v126)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v128+v129<<(uint(int32(2))%32)))) = v115
	v135 = v115
	goto L39
L46:
	;
	v137 = F_getsubdfa(m, l0, v104)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v139 != 0 {
		v944 = v139
		goto L10
	} else {
		goto L48
	}
L48:
	;
	v140 = int32(0)
	v142 = F_shortest(m, l0, v135, l2, l2, l3, v140, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v144 != 0 {
		v944 = v144
		goto L10
	} else {
		goto L50
	}
L50:
	;
	if v142 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return int32(1)
L52:
	;
	goto L53
L53:
	;
	v156 = v142
	goto L54
L54:
	;
	v163 = F_longest(m, l0, v137, v156, l3, int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	v944 = int32(1)
	goto L10
L56:
	;
	if v163 == l3 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v166 = F_cdissect(m, l0, v98, l2, v156)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v205 != 0 {
		v944 = v205
		goto L10
	} else {
		goto L79
	}
L60:
	;
	if v166 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v170 = F_cdissect(m, l0, v104, v156, l3)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	v201 = v166
	goto L63
L63:
	;
	if v201 != int32(1) {
		v944 = v201
		goto L10
	} else {
		goto L78
	}
L64:
	;
	if v170 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v907 = int32(0)
	goto L11
L66:
	;
	goto L67
L67:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v176 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v201 = v170
	goto L63
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v192 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v179) <= base.Ui32(v176) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v182 = v176 << (uint(int32(3)) % 32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v185 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v182+v183))) = v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v187+v182)+4)) = v185
	goto L69
L72:
	;
	v194 = v192
	goto L75
L73:
	;
	goto L74
L74:
	;
	goto L68
L75:
	;
	F_zaptreesubs(m, l0, v194)
	mBase = m.M
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	if v197 != 0 {
		v194 = v197
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L74
L77:
	;
	goto L76
L78:
	;
	goto L59
L79:
	;
	if l3 == v156 {
		v944 = int32(1)
		goto L10
	} else {
		goto L80
	}
L80:
	;
	v210 = int32(0)
	v212 = F_shortest(m, l0, v135, l2, v156+int32(4), l3, v210, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v214 != 0 {
		v944 = v214
		goto L10
	} else {
		goto L82
	}
L82:
	;
	if v212 != 0 {
		v156 = v212
		goto L54
	} else {
		goto L83
	}
L83:
	;
	goto L55
L84:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v244 != 0 {
		v944 = v244
		goto L10
	} else {
		goto L91
	}
L85:
	;
	v216 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v223 = F_newdfa(m, l0, v98+int32(36), v219+int32(72), v216)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v223 == int32(0) {
		v243 = v216
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v227 == int32(98) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+60)) = v230
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v223)+64)) = uint16(v232)
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v223)+66)) = uint16(v234)
	goto L90
L89:
	;
	goto L90
L90:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v237<<(uint(int32(2))%32)))) = v223
	v243 = v223
	goto L84
L91:
	;
	v245 = F_getsubdfa(m, l0, v104)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v247 != 0 {
		v944 = v247
		goto L10
	} else {
		goto L93
	}
L93:
	;
	v249 = F_longest(m, l0, v243, l2, l3, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v251 != 0 {
		v944 = v251
		goto L10
	} else {
		goto L95
	}
L95:
	;
	if v249 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	return int32(1)
L97:
	;
	goto L98
L98:
	;
	v263 = v249
	goto L99
L99:
	;
	v270 = F_longest(m, l0, v245, v263, l3, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L101
	}
L100:
	;
	v944 = int32(1)
	goto L10
L101:
	;
	if v270 == l3 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v273 = F_cdissect(m, l0, v98, l2, v263)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v312 != 0 {
		v944 = v312
		goto L10
	} else {
		goto L124
	}
L105:
	;
	if v273 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v277 = F_cdissect(m, l0, v104, v263, l3)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L109
	}
L107:
	;
	v308 = v273
	goto L108
L108:
	;
	if v308 != int32(1) {
		v944 = v308
		goto L10
	} else {
		goto L123
	}
L109:
	;
	if v277 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v907 = int32(0)
	goto L11
L111:
	;
	goto L112
L112:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v283 <= int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v308 = v277
	goto L108
L114:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v299 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v286) <= base.Ui32(v283) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v289 = v283 << (uint(int32(3)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v292 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v289+v290))) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v294+v289)+4)) = v292
	goto L114
L117:
	;
	v301 = v299
	goto L120
L118:
	;
	goto L119
L119:
	;
	goto L113
L120:
	;
	F_zaptreesubs(m, l0, v301)
	mBase = m.M
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+24))
	if v304 != 0 {
		v301 = v304
		goto L120
	} else {
		goto L122
	}
L121:
	;
	goto L119
L122:
	;
	goto L121
L123:
	;
	goto L104
L124:
	;
	if l2 == v263 {
		v944 = int32(1)
		goto L10
	} else {
		goto L125
	}
L125:
	;
	v318 = F_longest(m, l0, v243, l2, v263-int32(4), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v320 != 0 {
		v944 = v320
		goto L10
	} else {
		goto L127
	}
L127:
	;
	if v318 != 0 {
		v263 = v318
		goto L99
	} else {
		goto L128
	}
L128:
	;
	goto L100
L129:
	;
	return int32(1)
L130:
	;
	goto L131
L131:
	;
	v333 = v322
	goto L132
L132:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v340+v341<<(uint(int32(2))%32))))
	if v345 != 0 {
		v372 = v345
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v944 = int32(1)
	goto L10
L134:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v374 != 0 {
		v944 = v374
		goto L10
	} else {
		goto L141
	}
L135:
	;
	v346 = int32(0)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v353 = F_newdfa(m, l0, v333+int32(36), v349+int32(72), v346)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	if v353 == int32(0) {
		v372 = v346
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v357 == int32(98) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+60)) = v360
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v353)+64)) = uint16(v362)
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v333)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v353)+66)) = uint16(v364)
	goto L140
L139:
	;
	goto L140
L140:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v333)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v366+v367<<(uint(int32(2))%32)))) = v353
	v372 = v353
	goto L134
L141:
	;
	v376 = F_longest(m, l0, v372, l2, l3, int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	if v376 == l3 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v379 = F_cdissect(m, l0, v333, l2, l3)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v384 != 0 {
		v944 = v384
		goto L10
	} else {
		goto L148
	}
L146:
	;
	if v379 != int32(1) {
		v907 = v379
		goto L11
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v333)+24))
	if v386 != 0 {
		v333 = v386
		goto L132
	} else {
		goto L149
	}
L149:
	;
	goto L133
L150:
	;
	if v387 <= int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if l2 == l3 {
		v907 = v5
		goto L11
	} else {
		goto L154
	}
L152:
	;
	v398 = v387
	goto L153
L153:
	;
	v403 = (l3 - l2) >> (uint(int32(2)) % 32)
	v404 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+18)))
	if base.Ui32(v403) < base.Ui32(v404) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v398 = int32(1)
	goto L153
L155:
	;
	v406 = v403
	goto L157
L156:
	;
	v406 = v404
	goto L157
L157:
	;
	if v404 == int32(256) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v409 = v403
	goto L160
L159:
	;
	v409 = v406
	goto L160
L160:
	;
	if base.Ui32(v398) < base.Ui32(v409) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v411 = v409
	goto L163
L162:
	;
	v411 = v398
	goto L163
L163:
	;
	v414 = F_palloc_mul_extended(m, int32(4), v411+int32(1))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	if v414 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	return int32(12)
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = l2
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v422 = F_getsubdfa(m, l0, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v424 != 0 {
		goto L12
	} else {
		goto L169
	}
L169:
	;
	v429 = l2
	v430 = int32(1)
	v432 = v5
	goto L170
L170:
	;
	v443 = int32(2)
	v449 = v430 - int32(1)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v414+v449<<(uint(v443)%32))))
	if base.B2i32(base.Ui32(v430) < base.Ui32(v398))&base.B2i32((l3-v429)>>(uint(v443)%32) <= v398-v430)|(base.B2i32(l3 == v429)|base.B2i32(v453 != v429)) != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	F_pfree(m, v414)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L223
	}
L172:
	;
	v457 = v429
	goto L174
L173:
	;
	v457 = v429 + int32(4)
	goto L174
L174:
	;
	if base.Ui32(v430) < base.Ui32(v411) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v462 = v457
	goto L177
L176:
	;
	v462 = l3
	goto L177
L177:
	;
	v463 = int32(0)
	v465 = F_shortest(m, l0, v422, v453, v462, l3, v463, v463)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414+v430<<(uint(int32(2))%32)))) = v465
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v468 != 0 {
		goto L12
	} else {
		goto L179
	}
L179:
	;
	if v465 == int32(0) {
		v542 = v432
		goto L184
	} else {
		goto L185
	}
L180:
	;
	goto L171
L181:
	;
	if int32(0) < v601 {
		v429 = v600
		v430 = v601
		v432 = v603
		goto L170
	} else {
		goto L222
	}
L182:
	;
	v577 = v564
	goto L216
L183:
	;
	if v556 <= int32(0) {
		goto L180
	} else {
		goto L215
	}
L184:
	;
	v549 = v542
	v556 = v449
	goto L183
L185:
	;
	if v432 < v449 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v472 = v432
	goto L188
L187:
	;
	v472 = v449
	goto L188
L188:
	;
	if l3 != v465 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	if base.Ui32(v411) <= base.Ui32(v430) {
		v542 = v472
		goto L184
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	if base.Ui32(v430) < base.Ui32(v398) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v600 = v465
	v601 = v430 + int32(1)
	v603 = v472
	goto L181
L193:
	;
	v564 = v430
	v565 = v472
	goto L182
L194:
	;
	goto L195
L195:
	;
	v485 = v472
	goto L197
L196:
	;
	F_pfree(m, v414)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L214
	}
L197:
	;
	v492 = v485 + int32(1)
	if v430 < v492 {
		goto L196
	} else {
		goto L199
	}
L198:
	;
	if v530 == int32(1) {
		v549 = v485
		v556 = v492
		goto L183
	} else {
		goto L212
	}
L199:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	if v496 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v522 = int32(2)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v414+v485<<(uint(v522)%32))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v414+v492<<(uint(v522)%32))))
	v530 = F_cdissect(m, l0, v521, v525, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L4
	} else {
		goto L210
	}
L201:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	if v512 != 0 {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v499) <= base.Ui32(v496) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v502 = v496 << (uint(int32(3)) % 32)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v505 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v502+v503))) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v507+v502)+4)) = v505
	goto L201
L204:
	;
	v514 = v512
	goto L207
L205:
	;
	goto L206
L206:
	;
	goto L200
L207:
	;
	F_zaptreesubs(m, l0, v514)
	mBase = m.M
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	if v517 != 0 {
		v514 = v517
		goto L207
	} else {
		goto L209
	}
L208:
	;
	goto L206
L209:
	;
	goto L208
L210:
	;
	if v530 == int32(0) {
		v485 = v492
		goto L197
	} else {
		goto L211
	}
L211:
	;
	goto L198
L212:
	;
	F_pfree(m, v414)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	return v530
L214:
	;
	v907 = int32(0)
	goto L11
L215:
	;
	v564 = v556
	v565 = v549
	goto L182
L216:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v414+v577<<(uint(int32(2))%32))))
	if base.Ui32(v588) < base.Ui32(l3) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L180
L218:
	;
	v600 = v588 + int32(4)
	v601 = v577
	v603 = v565
	goto L181
L219:
	;
	goto L220
L220:
	;
	v592 = int32(1)
	if v592 < v577 {
		v577 = v577 - v592
		goto L216
	} else {
		goto L221
	}
L221:
	;
	goto L217
L222:
	;
	goto L180
L223:
	;
	return int32(1)
L224:
	;
	v907 = v629
	goto L11
L225:
	;
	v638 = v635
	goto L227
L226:
	;
	v638 = v636
	goto L227
L227:
	;
	if v636 == int32(256) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v641 = v635
	goto L230
L229:
	;
	v641 = v638
	goto L230
L230:
	;
	v642 = int32(1)
	if v387 <= v642 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v645 = v642
	goto L233
L232:
	;
	v645 = v387
	goto L233
L233:
	;
	if base.Ui32(v645) < base.Ui32(v641) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v647 = v641
	goto L236
L235:
	;
	v647 = v645
	goto L236
L236:
	;
	v650 = F_palloc_mul_extended(m, int32(4), v647+int32(1))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	if v650 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	return int32(12)
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v650))) = l2
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v658 = F_getsubdfa(m, l0, v657)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v660 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	F_pfree(m, v650)
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L4
	} else {
		goto L295
	}
L243:
	;
	v666 = int32(1)
	v667 = v5
	v669 = l3
	goto L244
L244:
	;
	v674 = int32(2)
	v678 = v666 - int32(1)
	v681 = v650 + v678<<(uint(v674)%32)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v684 = F_longest(m, l0, v658, v682, v669, int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L246
	}
L245:
	;
	F_pfree(m, v650)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L4
	} else {
		goto L294
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v650+v666<<(uint(v674)%32)))) = v684
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v687 != 0 {
		goto L242
	} else {
		goto L247
	}
L247:
	;
	if v684 == int32(0) {
		v771 = v667
		goto L252
	} else {
		goto L253
	}
L248:
	;
	goto L245
L249:
	;
	if int32(0) < v841 {
		v666 = v841
		v667 = v842
		v669 = v844
		goto L244
	} else {
		goto L293
	}
L250:
	;
	v804 = v791
	goto L284
L251:
	;
	if v776 <= int32(0) {
		goto L248
	} else {
		goto L283
	}
L252:
	;
	v776 = v678
	v779 = v771
	goto L251
L253:
	;
	if v667 < v678 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v691 = v667
	goto L256
L255:
	;
	v691 = v678
	goto L256
L256:
	;
	if l3 != v684 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v791 = v666
	v794 = v691
	goto L250
L258:
	;
	if base.Ui32(v647) <= base.Ui32(v666) {
		v771 = v691
		goto L252
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	if base.Ui32(v666) < base.Ui32(v645) {
		goto L257
	} else {
		goto L263
	}
L261:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	if (base.B2i32(v645-v666 < (l3-v684)>>(uint(int32(2))%32))|base.B2i32(base.Ui32(v645) <= base.Ui32(v666)))&base.B2i32(v701 == v684) != 0 {
		goto L257
	} else {
		goto L262
	}
L262:
	;
	v841 = v666 + int32(1)
	v842 = v691
	v844 = l3
	goto L249
L263:
	;
	v713 = v691
	goto L265
L264:
	;
	F_pfree(m, v650)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L282
	}
L265:
	;
	v721 = v713 + int32(1)
	if v666 < v721 {
		goto L264
	} else {
		goto L267
	}
L266:
	;
	if v759 == int32(1) {
		v776 = v721
		v779 = v713
		goto L251
	} else {
		goto L280
	}
L267:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v723)+8))
	if v725 <= int32(0) {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v751 = int32(2)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v650+v713<<(uint(v751)%32))))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v650+v721<<(uint(v751)%32))))
	v759 = F_cdissect(m, l0, v750, v754, v758)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L4
	} else {
		goto L278
	}
L269:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v723)+20))
	if v741 != 0 {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v728) <= base.Ui32(v725) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v731 = v725 << (uint(int32(3)) % 32)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v734 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v731+v732))) = v734
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v736+v731)+4)) = v734
	goto L269
L272:
	;
	v743 = v741
	goto L275
L273:
	;
	goto L274
L274:
	;
	goto L268
L275:
	;
	F_zaptreesubs(m, l0, v743)
	mBase = m.M
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v743)+24))
	if v746 != 0 {
		v743 = v746
		goto L275
	} else {
		goto L277
	}
L276:
	;
	goto L274
L277:
	;
	goto L276
L278:
	;
	if v759 == int32(0) {
		v713 = v721
		goto L265
	} else {
		goto L279
	}
L279:
	;
	goto L266
L280:
	;
	F_pfree(m, v650)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	return v759
L282:
	;
	v907 = int32(0)
	goto L11
L283:
	;
	v791 = v776
	v794 = v779
	goto L250
L284:
	;
	v815 = v650 + v804<<(uint(int32(2))%32)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v815)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v815-int32(4))))
	if base.Ui32(v816) <= base.Ui32(v819) {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	goto L248
L286:
	;
	v832 = int32(1)
	if v832 < v804 {
		v804 = v804 - v832
		goto L284
	} else {
		goto L292
	}
L287:
	;
	v822 = v816 - int32(4)
	if base.Ui32(v819) < base.Ui32(v822) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v841 = v804
	v842 = v794
	v844 = v822
	goto L249
L289:
	;
	goto L290
L290:
	;
	if base.B2i32(v645-v804 < (l3-v819)>>(uint(int32(2))%32))|base.B2i32(base.Ui32(v645) <= base.Ui32(v804)) != 0 {
		goto L286
	} else {
		goto L291
	}
L291:
	;
	v841 = v804
	v842 = v794
	v844 = v822
	goto L249
L292:
	;
	goto L285
L293:
	;
	goto L248
L294:
	;
	v867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v907 = base.B2i32(l2 != l3) | base.B2i32(v867 != int32(0))
	goto L11
L295:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v907 = v886
	goto L11
L296:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v907 = v902
	goto L11
L297:
	;
	return v907
L298:
	;
	goto L299
L299:
	;
	if v907 != 0 {
		v944 = v907
		goto L10
	} else {
		goto L300
	}
L300:
	;
	v920 = int32(0)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if base.Ui32(v921) <= base.Ui32(v916) {
		v944 = v920
		goto L10
	} else {
		goto L301
	}
L301:
	;
	v924 = v916 << (uint(int32(3)) % 32)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v929 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v924+v925))) = (l2 - v927) >> (uint(v929) % 32)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v932+v924)+4)) = (l3 - v934) >> (uint(v929) % 32)
	v944 = v920
	goto L10
}
