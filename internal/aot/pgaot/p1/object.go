package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_alloc_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
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
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v511 int32
	_ = v511
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = v21 + l1<<(uint(int32(5))%32)
	v26 = v24 + int32(224)
	v28 = F_LWLockAcquire(m, v26, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+244))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L168
	}
L4:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v833)))
	F_LWLockRelease(m, v849+v834<<(uint(int32(5))%32)+int32(224))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L167
	}
L5:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_c_F_alloc_object[0]))))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v717 = l0
	v718 = l1
	v719 = v32
	v727 = v19
	goto L7
L7:
	;
	v733 = int32(0)
	v736 = base.AtomicRmwOr32(m, v733, int32(_a_F_alloc_object_0), v733)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v717)+652))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v717)))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)+1468))
	if v737 != v739 {
		goto L141
	} else {
		goto L142
	}
L8:
	;
	v48 = l0 + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v41 = base.I32_div_u_s(int32(_a_F_alloc_object_1), v37)
	v46 = v41 - int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v45 = base.I32_div_u_s(int32(_a_F_alloc_object_2), v37)
	v46 = v45
	goto L8
L12:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v717 = l0
	v718 = l1
	v719 = v716
	v727 = v19
	goto L7
L13:
	;
	v56 = v49
	goto L16
L14:
	;
	goto L15
L15:
	;
	v530 = F_transfer_first_span(m, l0, v26, int32(2), int32(1))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L103
	}
L16:
	;
	v68 = int32(0)
	v71 = base.AtomicRmwOr32(m, v68, int32(_a_F_alloc_object_0), v68)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+1468))
	if v72 != v74 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v511 != 0 {
		goto L12
	} else {
		goto L102
	}
L18:
	;
	v76 = int32(0)
	v80 = F_LWLockAcquire(m, v73+int32(1476), v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v162 = int32(base.Ui32(v56) >> (uint(int32(27)) % 32))
	v165 = l0 + v162*int32(20)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	if v166 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+1468))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v83 != v84 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = v76
	goto L25
L23:
	;
	v140 = v82
	goto L24
L24:
	;
	F_LWLockRelease(m, v140+int32(1476))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L32
	}
L25:
	;
	v104 = v48 + v92*int32(20)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v105 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v83
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v140 = v123
	goto L24
L27:
	;
	v119 = v92 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v119) <= base.Ui32(v120) {
		v92 = v119
		goto L25
	} else {
		goto L31
	}
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+24)))
	if v108 != int32(1) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	F_dsm_detach(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = int64(0)
	goto L27
L31:
	;
	goto L26
L32:
	;
	goto L20
L33:
	;
	v170 = v166
	goto L35
L34:
	;
	v167 = F_get_segment_by_index(m, l0, v162)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v173 = v170 + v56&int32(134217727)
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+24)))
	v178 = base.I32_div_u_s((v46-v174)*int32(3), v46)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v179 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v170 = v169
	goto L35
L37:
	;
	v180 = int32(0)
	v183 = base.AtomicRmwOr32(m, v180, int32(_a_F_alloc_object_0), v180)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+1468))
	if v184 != v186 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v303 = int32(0)
	goto L39
L39:
	;
	if v178 <= int32(1) {
		goto L59
	} else {
		goto L60
	}
L40:
	;
	v188 = int32(0)
	v192 = F_LWLockAcquire(m, v185+int32(1476), v188)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v276 = int32(base.Ui32(v179) >> (uint(int32(27)) % 32))
	v279 = l0 + v276*int32(20)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	if v280 != 0 {
		goto L55
	} else {
		goto L56
	}
L43:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+1468))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v195 != v196 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v200 = v188
	goto L47
L45:
	;
	v252 = v194
	goto L46
L46:
	;
	F_LWLockRelease(m, v252+int32(1476))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	v216 = v48 + v200*int32(20)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	if v217 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v195
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v252 = v235
	goto L46
L49:
	;
	v231 = v200 + int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v231) <= base.Ui32(v232) {
		v200 = v231
		goto L47
	} else {
		goto L53
	}
L50:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+24)))
	if v220 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	F_dsm_detach(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v216))) = int64(0)
	goto L49
L53:
	;
	goto L48
L54:
	;
	goto L42
L55:
	;
	v284 = v280
	goto L57
L56:
	;
	v281 = F_get_segment_by_index(m, l0, v276)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v303 = v284 + v179&int32(134217727)
	goto L39
L58:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v284 = v283
	goto L57
L59:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v56 == v306 {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	goto L61
L61:
	;
	if v179 != 0 {
		v56 = v179
		goto L16
	} else {
		goto L101
	}
L62:
	;
	v364 = v24 + int32(240) + v178<<(uint(int32(2))%32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v365
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = int32(0)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v370 != 0 {
		goto L79
	} else {
		goto L80
	}
L63:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v356
	goto L62
L64:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v308
	if v303 == int32(0) {
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v315 = int32(0)
	v318 = base.AtomicRmwOr32(m, v315, int32(_a_F_alloc_object_0), v315)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+1468))
	if v319 != v321 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = int32(0)
	goto L63
L68:
	;
	v326 = F_LWLockAcquire(m, v320+int32(1476), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v338 = int32(base.Ui32(v314) >> (uint(int32(27)) % 32))
	v341 = l0 + v338*int32(20)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	if v342 != 0 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v330+int32(1476))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v346 = v342
	goto L76
L75:
	;
	v343 = F_get_segment_by_index(m, l0, v338)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v346+v314&int32(134217727))+8)) = v348
	if v303 == int32(0) {
		goto L62
	} else {
		goto L78
	}
L77:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	v346 = v345
	goto L76
L78:
	;
	goto L63
L79:
	;
	v371 = int32(0)
	v374 = base.AtomicRmwOr32(m, v371, int32(_a_F_alloc_object_0), v371)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+1468))
	if v375 != v377 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v173)+30)) = uint16(v178)
	goto L61
L82:
	;
	v379 = int32(0)
	v383 = F_LWLockAcquire(m, v376+int32(1476), v379)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v467 = int32(base.Ui32(v370) >> (uint(int32(27)) % 32))
	v470 = l0 + v467*int32(20)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	if v471 != 0 {
		goto L97
	} else {
		goto L98
	}
L85:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v385)+1468))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v386 != v387 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v391 = v379
	goto L89
L87:
	;
	v443 = v385
	goto L88
L88:
	;
	F_LWLockRelease(m, v443+int32(1476))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L96
	}
L89:
	;
	v407 = v48 + v391*int32(20)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+8))
	if v408 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v386
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v443 = v426
	goto L88
L91:
	;
	v422 = v391 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v422) <= base.Ui32(v423) {
		v391 = v422
		goto L89
	} else {
		goto L95
	}
L92:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+24)))
	if v411 != int32(1) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	F_dsm_detach(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = int64(0)
	goto L91
L95:
	;
	goto L90
L96:
	;
	goto L84
L97:
	;
	v475 = v471
	goto L99
L98:
	;
	v472 = F_get_segment_by_index(m, l0, v467)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475+v370&int32(134217727))+4)) = v56
	goto L81
L100:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)+12))
	v475 = v474
	goto L99
L101:
	;
	goto L17
L102:
	;
	goto L15
L103:
	;
	if v530 != 0 {
		goto L12
	} else {
		goto L104
	}
L104:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	if v532 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v537 = F_transfer_first_span(m, l0, v26, int32(0), int32(1))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if l1 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	if v537 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v574 = F_FreePageManagerGet(m, v571, v550, v19+int32(12))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L124
	}
L111:
	;
	v833 = l0
	v834 = l1
	v837 = int32(0)
	v843 = v19
	goto L4
L112:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v555 = F_LWLockAcquire(m, v551+int32(1476), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L118
	}
L113:
	;
	v549 = int32(0)
	v550 = int32(1)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v544 = F_alloc_object(m, l0, int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v544 == int32(0) {
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v549 = v544
	v550 = int32(16)
	goto L112
L118:
	;
	v557 = F_get_best_segment(m, l0, v550)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v557 != 0 {
		v570 = v557
		goto L110
	} else {
		goto L120
	}
L120:
	;
	v559 = F_make_new_segment(m, l0, v550)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	if v559 != 0 {
		v570 = v559
		goto L110
	} else {
		goto L122
	}
L122:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v561+int32(1476))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L111
L124:
	;
	if v574 == int32(0) {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v578+int32(1476))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v588 = base.I32_div_s(v570-v48, int32(20))
	v591 = v583<<(uint(int32(12))%32) | v588<<(uint(int32(27))%32)
	if l1 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v592 = v549
	goto L129
L128:
	;
	v592 = v591
	goto L129
L129:
	;
	F_init_span(m, l0, v592, v26, v591, v550, l1)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v595 = int32(0)
	if l1 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v604 = v595
	v613 = int32(0)
	goto L134
L132:
	;
	v657 = v595
	goto L133
L133:
	;
	v673 = v657
	v677 = v595
	goto L138
L134:
	;
	v618 = int32(2)
	v619 = v604 << (uint(v618) % 32)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v570)+16))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v619+(v620+v621<<(uint(v618)%32))))) = v592
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v570)+16))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v627+v628<<(uint(v618)%32)+v619)+4)) = v592
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v570)+16))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v634+v635<<(uint(v618)%32)+v619)+8)) = v592
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v570)+16))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v641+v642<<(uint(v618)%32)+v619)+12)) = v592
	v648 = int32(4)
	v649 = v604 + v648
	v651 = v613 + v648
	if v651 != v550&int32(16) {
		v604 = v649
		v613 = v651
		goto L134
	} else {
		goto L136
	}
L135:
	;
	if v550&int32(1) == int32(0) {
		goto L12
	} else {
		goto L137
	}
L136:
	;
	goto L135
L137:
	;
	v657 = v649
	goto L133
L138:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v570)+16))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v689 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v687+v688<<(uint(v689)%32)+v673<<(uint(v689)%32)))) = v592
	v696 = int32(1)
	if v677 != 0 {
		v673 = v673 + v696
		v677 = v677 + v696
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L12
L140:
	;
	goto L139
L141:
	;
	v744 = F_LWLockAcquire(m, v738+int32(1476), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v756 = int32(base.Ui32(v719) >> (uint(int32(27)) % 32))
	v759 = v717 + v756*int32(20)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)+12))
	if v760 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	F_check_for_freed_segments_locked(m, v717)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v717)))
	F_LWLockRelease(m, v748+int32(1476))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v763 = F_get_segment_by_index(m, v717, v756)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	v766 = v760
	goto L149
L149:
	;
	v769 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718<<(uint(int32(1))%32))+uint32(_c_F_alloc_object[0]))))
	v770 = v766 + v719&int32(134217727)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v770)+12))
	v772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+26)))
	if v772 != int32(_a_F_alloc_object_3) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v759)+12))
	v766 = v765
	goto L149
L151:
	;
	v823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+24)))
	v825 = v823 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v770)+24)) = uint16(v825)
	if v825&int32(_a_F_alloc_object_3) != 0 {
		v833 = v717
		v834 = v718
		v837 = v820
		v843 = v727
		goto L4
	} else {
		goto L165
	}
L152:
	;
	v775 = int32(0)
	v778 = base.AtomicRmwOr32(m, v775, int32(_a_F_alloc_object_0), v775)
	v780 = v772*v769 + v771
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v717)+652))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v717)))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)+1468))
	if v781 != v783 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	v812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+22)))
	v814 = v812 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v770)+22)) = uint16(v814)
	v820 = v812*v769 + v771
	goto L151
L155:
	;
	v788 = F_LWLockAcquire(m, v782+int32(1476), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v800 = int32(base.Ui32(v780) >> (uint(int32(27)) % 32))
	v803 = v717 + v800*int32(20)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+12))
	if v804 != 0 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	F_check_for_freed_segments_locked(m, v717)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v717)))
	F_LWLockRelease(m, v792+int32(1476))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	goto L157
L161:
	;
	v808 = v804
	goto L163
L162:
	;
	v805 = F_get_segment_by_index(m, v717, v800)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L164
	}
L163:
	;
	v810 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v808+v780&int32(134217727)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v770)+26)) = uint16(v810)
	v820 = v780
	goto L151
L164:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v803)+12))
	v808 = v807
	goto L163
L165:
	;
	v831 = F_transfer_first_span(m, v717, v26, int32(1), int32(3))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v833 = v717
	v834 = v718
	v837 = v820
	v843 = v727
	goto L4
L167:
	;
	m.G0 = v843 + int32(16)
	return v837
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v550
	F_errmsg_internal(m, int32(_a_F_alloc_object_4), v19)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_alloc_object_5), int32(1769), int32(_a_F_alloc_object_6))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_free_object_addresses(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pfree(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 != 0 {
			F_pfree(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_get_object_address(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int64
	_ = v34
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int64
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int64
	_ = v702
	var v704 int64
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int64
	_ = v719
	var v726 int32
	_ = v726
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int64
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v839 int32
	_ = v839
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int64
	_ = v974
	var v977 int64
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int64
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int64
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int64
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1012 int64
	_ = v1012
	var v1014 int64
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1175 int64
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1303 int64
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int64
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1334 int64
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1393 int64
	_ = v1393
	var v1394 int64
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1423 int64
	_ = v1423
	var v1424 int64
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1486 int32
	_ = v1486
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1544 int32
	_ = v1544
	var v1555 int32
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1572 int64
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1668 int32
	_ = v1668
	var v1683 int32
	_ = v1683
	var v1696 int32
	_ = v1696
	var v1711 int32
	_ = v1711
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1745 int32
	_ = v1745
	var v1750 int32
	_ = v1750
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1783 int32
	_ = v1783
	var v1788 int32
	_ = v1788
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1830 int32
	_ = v1830
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1917 int32
	_ = v1917
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1971 int32
	_ = v1971
	var v1984 int32
	_ = v1984
	var v2001 int32
	_ = v2001
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2039 int32
	_ = v2039
	var v2051 int32
	_ = v2051
	var v2064 int32
	_ = v2064
	var v2081 int32
	_ = v2081
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2116 int64
	_ = v2116
	var v2118 int32
	_ = v2118
	v7 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(496)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v34 = *(*int64)(unsafe.Add(mBase, _c_F_get_object_address[0]))
	v44 = v7
	v50 = v7
	v52 = v7
	v53 = v7
	v55 = v34
	goto L2
L1:
	;
	m.G0 = v27 + int32(496)
	return
L2:
	;
	switch l1 {
	case 0, 9, 14, 15, 16, 17, 21, 27, 30, 34, 37, 39, 43:
		goto L41
	case 1, 19, 29, 35:
		goto L40
	case 2, 3:
		goto L35
	case 4, 6:
		goto L44
	case 5:
		goto L33
	case 7:
		goto L38
	case 8:
		goto L37
	case 10:
		goto L43
	case 11:
		goto L24
	case 12, 50:
		goto L22
	case 13:
		goto L20
	case 18, 20, 23, 38, 42, 52:
		goto L45
	case 22:
		goto L34
	case 24, 26:
		goto L36
	case 25:
		goto L39
	case 28, 36, 41, 45:
		goto L42
	case 31, 33:
		goto L25
	case 32:
		goto L26
	case 40:
		goto L23
	case 44:
		goto L32
	case 46:
		goto L28
	case 47:
		goto L30
	case 48:
		goto L31
	case 49:
		goto L29
	case 51:
		goto L27
	default:
		v1629 = v44
		goto L21
	}
L3:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L601
	}
L4:
	;
	if v1926 == int32(0) {
		goto L1
	} else {
		goto L528
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1911
	v1917 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1917
	v1926 = v1911
	v1927 = v1912
	v1929 = v1917
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1894
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1896
	v1926 = v1894
	v1927 = v1895
	v1929 = v1896
	goto L4
L7:
	;
	v1894 = v1882
	v1895 = int32(0)
	v1896 = v1884
	goto L6
L8:
	;
	v1875 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1875
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1871
	v1879 = int32(826)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1879
	v1926 = v1871
	v1927 = v1875
	v1929 = v1879
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+456)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+452)) = v1519
	*(*int32)(unsafe.Add(mBase, uint32(v27)+448)) = v1520
	F_errmsg(m, int32(_a_F_get_object_address_0), v27+int32(448))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L46
	} else {
		goto L526
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L46
	} else {
		goto L522
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+324)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(v27)+320)) = v1438
	F_errmsg(m, int32(_a_F_get_object_address_1), v27+int32(320))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L46
	} else {
		goto L520
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L46
	} else {
		goto L516
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L46
	} else {
		goto L512
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L46
	} else {
		goto L507
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L46
	} else {
		goto L503
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L46
	} else {
		goto L498
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L46
	} else {
		goto L494
	}
L18:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1926 = v1711
	v1927 = int32(0)
	v1929 = v1696
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1683
	v1696 = v1668
	goto L18
L20:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1646)))
	F_get_object_address_type(m, v27+int32(484), int32(12), v1647, l5)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L46
	} else {
		goto L492
	}
L21:
	;
	if v1629 != 0 {
		v1696 = v1629
		goto L18
	} else {
		goto L488
	}
L22:
	;
	F_get_object_address_type(m, l0, l1, l2, l5)
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L46
	} else {
		goto L487
	}
L23:
	;
	v1621 = int32(3381)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1621
	v1624 = F_get_statistics_object_oid(m, l2, l5)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L46
	} else {
		goto L486
	}
L24:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+4))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(3) <= v1514 {
		goto L452
	} else {
		goto L453
	}
L25:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1408)))
	v1410 = F_makeRangeVarFromNameList(m, v1409)
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		goto L46
	} else {
		goto L426
	}
L26:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+4))
	v1379 = int32(0)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1376)))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+4))
	v1383 = F_get_namespace_oid(m, v1382, l5)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L46
	} else {
		goto L418
	}
L27:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+4))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+4))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1239)))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1242)+4))
	v1244 = int32(_a_F_get_object_address_2)
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243))))
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_object_address[1])))
	if base.B2i32(v1247 == int32(0))|base.B2i32(v1247 != v1250) != 0 {
		v1268 = v1247
		v1269 = v1250
		goto L373
	} else {
		goto L374
	}
L28:
	;
	v1234 = int32(3602)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1234
	v1237 = F_get_ts_config_oid(m, l2, l5)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L46
	} else {
		goto L369
	}
L29:
	;
	v1229 = int32(3764)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1229
	v1232 = F_get_ts_template_oid(m, l2, l5)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L46
	} else {
		goto L368
	}
L30:
	;
	v1224 = int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1224
	v1227 = F_get_ts_dict_oid(m, l2, l5)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L46
	} else {
		goto L367
	}
L31:
	;
	v1219 = int32(3601)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1219
	v1222 = F_get_ts_parser_oid(m, l2, l5)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L46
	} else {
		goto L366
	}
L32:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+4))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+4))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1206)))
	v1211 = F_LookupTypeNameOid(m, v1210, l5)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L46
	} else {
		goto L363
	}
L33:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+4))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1158)))
	v1162 = F_LookupTypeNameOid(m, v1161, l5)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L46
	} else {
		goto L351
	}
L34:
	;
	v1099 = int32(2613)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1099
	v1102 = m.G0
	v1104 = v1102 - int32(16)
	m.G0 = v1104
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v1106 - int32(473) {
	case 0:
		goto L338
	case 1:
		goto L340
	default:
		goto L339
	}
L35:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v891)+4))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v892+v893<<(uint(int32(2))%32)-int32(4))))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+4))
	v904 = v900
	goto L277
L36:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v860)+4))
	v862 = F_get_index_am_oid(m, v861)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L46
	} else {
		goto L266
	}
L37:
	;
	v683 = int32(2607)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v683
	v686 = m.G0
	v688 = v686 - int32(16)
	m.G0 = v688
	F_DeconstructQualifiedName(m, l2, v688+int32(12), v688+int32(8))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L46
	} else {
		goto L236
	}
L38:
	;
	v678 = int32(3456)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v678
	v681 = F_get_collation_oid(m, l2, l5)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L46
	} else {
		goto L235
	}
L39:
	;
	v673 = int32(2617)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v673
	v676 = F_LookupOperWithArgs(m, l2, l5)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L46
	} else {
		goto L234
	}
L40:
	;
	v668 = int32(1255)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v668
	v671 = F_LookupFuncWithArgs(m, l1, l2, l5)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L46
	} else {
		goto L233
	}
L41:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	switch l1 {
	case 0:
		goto L196
	default:
		goto L197
	case 9:
		goto L209
	case 14:
		goto L201
	case 15:
		goto L208
	case 16:
		goto L203
	case 17:
		goto L202
	case 21:
		goto L204
	case 27:
		goto L200
	case 30:
		goto L199
	case 34:
		goto L206
	case 37:
		goto L205
	case 39:
		goto L198
	case 43:
		goto L207
	}
L42:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v321 <= int32(1) {
		goto L13
	} else {
		goto L118
	}
L43:
	;
	if l2 == int32(0) {
		goto L15
	} else {
		goto L104
	}
L44:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L93
	}
L45:
	;
	v60 = F_makeRangeVarFromNameList(m, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	return
L47:
	;
	v62 = F_relation_openrv_extended(m, v60, l4, l5)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	if v62 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	switch l1 - int32(18) {
	case 0:
		goto L54
	default:
		goto L53
	case 2:
		goto L59
	case 5:
		goto L55
	case 20:
		goto L58
	case 24:
		goto L57
	case 34:
		goto L56
	}
L50:
	;
	v234 = int32(0)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	v1911 = v234
	v1912 = v62
	goto L5
L52:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v62)+56))
	v234 = v233
	goto L51
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L46
	} else {
		goto L90
	}
L54:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+119)))
	if v194 == int32(102) {
		goto L52
	} else {
		goto L85
	}
L55:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+119)))
	if v169 == int32(109) {
		goto L52
	} else {
		goto L80
	}
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+119)))
	if v144 == int32(118) {
		goto L52
	} else {
		goto L75
	}
L57:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+119)))
	switch v119 - int32(112) {
	case 0, 2:
		goto L52
	default:
		goto L70
	}
L58:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+119)))
	if v94 == int32(83) {
		goto L52
	} else {
		goto L65
	}
L59:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	if v67|int32(32) == int32(105) {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L46
	} else {
		goto L61
	}
L61:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L46
	} else {
		goto L62
	}
L62:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v79 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_3), v27+int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L46
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1371), int32(_a_F_get_object_address_5))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L46
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L46
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L46
	} else {
		goto L67
	}
L67:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v104 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_6), v27+int32(48))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L46
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1378), int32(_a_F_get_object_address_5))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L46
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L46
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L46
	} else {
		goto L72
	}
L72:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v129 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_7), v27-int32(-64))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L46
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1386), int32(_a_F_get_object_address_5))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L46
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L46
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L46
	} else {
		goto L77
	}
L77:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v154 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_8), v27+int32(80))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L46
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1393), int32(_a_F_get_object_address_5))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L46
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L46
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L46
	} else {
		goto L82
	}
L82:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v179 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_9), v27+int32(96))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L46
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1400), int32(_a_F_get_object_address_5))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L46
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L46
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L46
	} else {
		goto L87
	}
L87:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v204 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_10), v27+int32(112))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L46
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1407), int32(_a_F_get_object_address_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L46
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_11), v27+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L46
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1410), int32(_a_F_get_object_address_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L46
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v239 <= int32(1) {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v242+v239<<(uint(int32(2))%32)-int32(4))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	v252 = F_list_copy_head(m, l2, v239-int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L46
	} else {
		goto L95
	}
L95:
	;
	v254 = F_makeRangeVarFromNameList(m, v252)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L46
	} else {
		goto L96
	}
L96:
	;
	v256 = F_relation_openrv(m, v254, l4)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L46
	} else {
		goto L97
	}
L97:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	v259 = F_get_attnum(m, v258, v249)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L46
	} else {
		goto L98
	}
L98:
	;
	if v259 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if l5 == int32(0) {
		goto L16
	} else {
		goto L102
	}
L100:
	;
	v270 = v258
	v271 = v256
	v272 = v259
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v272
	v1911 = v270
	v1912 = v271
	goto L5
L102:
	;
	F_relation_close(m, v256, l4)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L46
	} else {
		goto L103
	}
L103:
	;
	v267 = int32(0)
	v270 = v267
	v271 = v267
	v272 = v267
	goto L101
L104:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v276 <= int32(1) {
		goto L15
	} else {
		goto L105
	}
L105:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279+v276<<(uint(int32(2))%32)-int32(4))))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+4))
	v289 = F_list_copy_head(m, l2, v276-int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L46
	} else {
		goto L106
	}
L106:
	;
	v291 = F_makeRangeVarFromNameList(m, v289)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L46
	} else {
		goto L107
	}
L107:
	;
	v293 = F_relation_openrv(m, v291, l4)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L46
	} else {
		goto L108
	}
L108:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)+52))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v293)+56))
	v297 = F_get_attnum(m, v296, v286)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L46
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v318 = int32(2604)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v318
	v1926 = v313
	v1927 = v314
	v1929 = v318
	goto L4
L110:
	;
	if l5 == int32(0) {
		goto L14
	} else {
		goto L116
	}
L111:
	;
	if v297 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295)+24))
	if v301 == int32(0) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v304 = F_GetAttrDefaultOid(m, v296, v297)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L46
	} else {
		goto L114
	}
L114:
	;
	if v304 != 0 {
		v313 = v304
		v314 = v293
		goto L109
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	F_relation_close(m, v293, l4)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L46
	} else {
		goto L117
	}
L117:
	;
	v311 = int32(0)
	v313 = v311
	v314 = v311
	goto L109
L118:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v324+v321<<(uint(int32(2))%32)-int32(4))))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v334 = F_list_copy_head(m, l2, v321-int32(1))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L46
	} else {
		goto L119
	}
L119:
	;
	v336 = F_makeRangeVarFromNameList(m, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L46
	} else {
		goto L120
	}
L120:
	;
	v339 = F_table_openrv_extended(m, v336, int32(1), l5)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L46
	} else {
		goto L121
	}
L121:
	;
	if v339 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339)+56))
	v343 = v341
	goto L124
L123:
	;
	v343 = int32(0)
	goto L124
L124:
	;
	switch l1 - int32(28) {
	case 0:
		goto L130
	default:
		goto L129
	case 8:
		goto L133
	case 13:
		goto L131
	case 17:
		goto L132
	}
L125:
	;
	v1894 = v573
	v1895 = v574
	v1896 = v575
	goto L6
L126:
	;
	v571 = int32(0)
	v573 = v571
	v574 = v571
	v575 = v565
	goto L125
L127:
	;
	if v560 != 0 {
		v573 = v560
		v574 = v339
		v575 = v554
		goto L125
	} else {
		goto L194
	}
L128:
	;
	v512 = m.G0
	v514 = v512 - int32(16)
	m.G0 = v514
	v519 = F_SearchSysCache2(m, int32(60), base.I64_extend_i32_u(v343), base.I64_extend_i32_u(v331))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L46
	} else {
		goto L183
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L46
	} else {
		goto L179
	}
L130:
	;
	if v339 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L131:
	;
	if v339 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L132:
	;
	if v339 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	if v339 != 0 {
		goto L128
	} else {
		goto L134
	}
L134:
	;
	v565 = int32(2618)
	goto L126
L135:
	;
	v565 = int32(2620)
	goto L126
L136:
	;
	goto L137
L137:
	;
	v350 = int32(2620)
	v352 = m.G0
	v354 = v352 - int32(128)
	m.G0 = v354
	v358 = F_table_open(m, v350, int32(1))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L46
	} else {
		goto L138
	}
L138:
	;
	v361 = v354 + int32(16)
	F_ScanKeyInit(m, v361, int32(2), int32(3), int32(184), base.I64_extend_i32_u(v343))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L46
	} else {
		goto L139
	}
L139:
	;
	F_ScanKeyInit(m, v354+int32(72), int32(4), int32(3), int32(62), base.I64_extend_i32_u(v331))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L46
	} else {
		goto L140
	}
L140:
	;
	v380 = F_systable_beginscan(m, v358, int32(2701), int32(1), int32(0), int32(2), v361)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L46
	} else {
		goto L142
	}
L141:
	;
	F_systable_endscan(m, v380)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L46
	} else {
		goto L153
	}
L142:
	;
	v382 = F_systable_getnext(m, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L46
	} else {
		goto L143
	}
L143:
	;
	if v382 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	if l5 != 0 {
		v409 = int32(0)
		goto L141
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v382)+16))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+22)))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405+v406)))
	v409 = v408
	goto L141
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L46
	} else {
		goto L148
	}
L148:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L46
	} else {
		goto L149
	}
L149:
	;
	v393 = F_get_rel_name(m, v343)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L46
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+4)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v331
	F_errmsg(m, int32(_a_F_get_object_address_12), v354)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L46
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_get_object_address_13), int32(1419), int32(_a_F_get_object_address_14))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L46
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
	F_relation_close(m, v358, int32(1))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L46
	} else {
		goto L154
	}
L154:
	;
	m.G0 = v354 + int32(128)
	v554 = v350
	v560 = v409
	goto L127
L155:
	;
	v565 = int32(2606)
	goto L126
L156:
	;
	goto L157
L157:
	;
	v422 = F_get_relation_constraint_oid(m, v343, v331, l5)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L46
	} else {
		goto L158
	}
L158:
	;
	v554 = int32(2606)
	v560 = v422
	goto L127
L159:
	;
	v565 = int32(3256)
	goto L126
L160:
	;
	goto L161
L161:
	;
	v427 = int32(3256)
	v429 = m.G0
	v431 = v429 - int32(128)
	m.G0 = v431
	v435 = F_table_open(m, v427, int32(1))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L46
	} else {
		goto L162
	}
L162:
	;
	v438 = v431 + int32(16)
	v439 = int32(3)
	F_ScanKeyInit(m, v438, v439, v439, int32(184), base.I64_extend_i32_u(v343))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L46
	} else {
		goto L163
	}
L163:
	;
	F_ScanKeyInit(m, v431+int32(72), int32(2), int32(3), int32(62), base.I64_extend_i32_u(v331))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L46
	} else {
		goto L164
	}
L164:
	;
	v457 = F_systable_beginscan(m, v435, int32(3258), int32(1), int32(0), int32(2), v438)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L46
	} else {
		goto L166
	}
L165:
	;
	F_systable_endscan(m, v457)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L46
	} else {
		goto L177
	}
L166:
	;
	v459 = F_systable_getnext(m, v457)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L46
	} else {
		goto L167
	}
L167:
	;
	if v459 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	if l5 != 0 {
		v486 = int32(0)
		goto L165
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v459)+16))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+22)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v482+v483)))
	v486 = v485
	goto L165
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L46
	} else {
		goto L172
	}
L172:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L46
	} else {
		goto L173
	}
L173:
	;
	v470 = F_get_rel_name(m, v343)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L46
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+4)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = v331
	F_errmsg(m, int32(_a_F_get_object_address_15), v431)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L46
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_get_object_address_16), int32(1246), int32(_a_F_get_object_address_17))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L46
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	F_relation_close(m, v435, int32(1))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L46
	} else {
		goto L178
	}
L178:
	;
	m.G0 = v431 + int32(128)
	v554 = v427
	v560 = v486
	goto L127
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_11), v27+int32(160))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L46
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1486), int32(_a_F_get_object_address_18))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L46
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	m.G0 = v514 + int32(16)
	v554 = int32(2618)
	v560 = v549
	goto L127
L183:
	;
	if v519 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if l5 != 0 {
		v549 = int32(0)
		goto L182
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v519)+16))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+22)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v542+v543)))
	F_ReleaseCatCache(m, v519)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L46
	} else {
		goto L193
	}
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L46
	} else {
		goto L188
	}
L188:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L46
	} else {
		goto L189
	}
L189:
	;
	v530 = F_get_rel_name(m, v343)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L46
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+4)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = v331
	F_errmsg(m, int32(_a_F_get_object_address_19), v514)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L46
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_get_object_address_20), int32(109), int32(_a_F_get_object_address_21))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L46
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	v549 = v545
	goto L182
L194:
	;
	F_relation_close(m, v339, int32(1))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L46
	} else {
		goto L195
	}
L195:
	;
	v565 = v554
	goto L126
L196:
	;
	v666 = F_get_am_type_oid(m, v581, int32(0), l5)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L46
	} else {
		goto L232
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L46
	} else {
		goto L229
	}
L198:
	;
	v647 = F_get_subscription_oid(m, v581, l5)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L46
	} else {
		goto L228
	}
L199:
	;
	v644 = F_get_publication_oid(m, v581, l5)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L46
	} else {
		goto L227
	}
L200:
	;
	v641 = F_ParameterAclLookup(m, v581, l5)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L46
	} else {
		goto L226
	}
L201:
	;
	v607 = m.G0
	v609 = v607 - int32(16)
	m.G0 = v609
	v613 = int64(0)
	v616 = F_GetSysCacheOid(m, int32(25), base.I64_extend_i32_u(v581), v613, v613, v613)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L46
	} else {
		goto L218
	}
L202:
	;
	v604 = F_get_foreign_server_oid(m, v581, l5)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L46
	} else {
		goto L217
	}
L203:
	;
	v601 = F_get_foreign_data_wrapper_oid(m, v581, l5)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L46
	} else {
		goto L216
	}
L204:
	;
	v598 = F_get_language_oid(m, v581, l5)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L46
	} else {
		goto L215
	}
L205:
	;
	v595 = F_get_namespace_oid(m, v581, l5)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L46
	} else {
		goto L214
	}
L206:
	;
	v592 = F_get_role_oid(m, v581, l5)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L46
	} else {
		goto L213
	}
L207:
	;
	v589 = F_get_tablespace_oid(m, v581, l5)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L46
	} else {
		goto L212
	}
L208:
	;
	v586 = F_get_extension_oid(m, v581, l5)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L46
	} else {
		goto L211
	}
L209:
	;
	v583 = F_get_database_oid(m, v581, l5)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L46
	} else {
		goto L210
	}
L210:
	;
	v1882 = v583
	v1884 = int32(1262)
	goto L7
L211:
	;
	v1882 = v586
	v1884 = int32(3079)
	goto L7
L212:
	;
	v1882 = v589
	v1884 = int32(1213)
	goto L7
L213:
	;
	v1882 = v592
	v1884 = int32(1260)
	goto L7
L214:
	;
	v1882 = v595
	v1884 = int32(2615)
	goto L7
L215:
	;
	v1882 = v598
	v1884 = int32(2612)
	goto L7
L216:
	;
	v1882 = v601
	v1884 = int32(2328)
	goto L7
L217:
	;
	v1882 = v604
	v1884 = int32(1417)
	goto L7
L218:
	;
	if v616|l5 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L46
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	m.G0 = v609 + int32(16)
	v1882 = v616
	v1884 = int32(3466)
	goto L7
L222:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L46
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v609))) = v581
	F_errmsg(m, int32(_a_F_get_object_address_22), v609)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L46
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_get_object_address_23), int32(590), int32(_a_F_get_object_address_24))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L46
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	v1882 = v641
	v1884 = int32(_a_F_get_object_address_25)
	goto L7
L227:
	;
	v1882 = v644
	v1884 = int32(_a_F_get_object_address_26)
	goto L7
L228:
	;
	v1882 = v647
	v1884 = int32(_a_F_get_object_address_27)
	goto L7
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_11), v27+int32(176))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L46
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1333), int32(_a_F_get_object_address_28))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L46
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	v1882 = v666
	v1884 = int32(2601)
	goto L7
L233:
	;
	v1668 = v668
	v1683 = v671
	goto L19
L234:
	;
	v1668 = v673
	v1683 = v676
	goto L19
L235:
	;
	v1668 = v678
	v1683 = v681
	goto L19
L236:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v688)+12))
	if v696 != 0 {
		goto L240
	} else {
		goto L241
	}
L237:
	;
	m.G0 = v688 + int32(16)
	v1668 = v683
	v1683 = v839
	goto L19
L238:
	;
	if v812|l5 != 0 {
		v839 = v812
		goto L237
	} else {
		goto L260
	}
L239:
	;
	v812 = int32(0)
	goto L238
L240:
	;
	v698 = F_LookupExplicitNamespace(m, v696, l5)
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L46
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L46
	} else {
		goto L249
	}
L243:
	;
	if v698 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v700 = int32(0)
	goto L246
L245:
	;
	v700 = l5
	goto L246
L246:
	;
	if v700 != 0 {
		goto L239
	} else {
		goto L247
	}
L247:
	;
	v702 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v688)+8)))
	v704 = int64(0)
	v706 = F_GetSysCacheOid(m, int32(18), v702, base.I64_extend_i32_u(v698), v704, v704)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L46
	} else {
		goto L248
	}
L248:
	;
	v812 = v706
	goto L238
L249:
	;
	v710 = int32(0)
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_address[2]))
	if v712 == v710 {
		v812 = v710
		goto L238
	} else {
		goto L250
	}
L250:
	;
	v715 = int32(0)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v716 <= v715 {
		goto L239
	} else {
		goto L251
	}
L251:
	;
	v719 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v688)+8)))
	v726 = v715
	goto L252
L252:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v712)+12))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v744+v726<<(uint(int32(2))%32))))
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_address[3]))
	if v748 != v750 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L239
L254:
	;
	v754 = int64(0)
	v756 = F_GetSysCacheOid(m, int32(18), v719, base.I64_extend_i32_u(v748), v754, v754)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L46
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v760 = v726 + int32(1)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v760 < v761 {
		v726 = v760
		goto L252
	} else {
		goto L259
	}
L257:
	;
	if v756 != 0 {
		v839 = v756
		goto L237
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	goto L253
L260:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L46
	} else {
		goto L261
	}
L261:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L46
	} else {
		goto L262
	}
L262:
	;
	v821 = F_NameListToString(m, l2)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L46
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v688))) = v821
	F_errmsg(m, int32(_a_F_get_object_address_29), v688)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L46
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_get_object_address_30), int32(_a_F_get_object_address_31), int32(_a_F_get_object_address_32))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L46
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	v865 = F_list_copy_tail(m, l2, int32(1))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L46
	} else {
		goto L267
	}
L267:
	;
	switch l1 - int32(24) {
	case 0:
		goto L268
	default:
		goto L269
	case 2:
		goto L270
	}
L268:
	;
	v888 = F_get_opclass_oid(m, v862, v865, l5)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L46
	} else {
		goto L275
	}
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L46
	} else {
		goto L272
	}
L270:
	;
	v870 = F_get_opfamily_oid(m, v862, v865, l5)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L46
	} else {
		goto L271
	}
L271:
	;
	v1882 = v870
	v1884 = int32(2753)
	goto L7
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_11), v27+int32(192))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L46
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1678), int32(_a_F_get_object_address_33))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L46
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	v1882 = v888
	v1884 = int32(2616)
	goto L7
L276:
	;
	v949 = int32(0)
	v952 = F_list_copy_head(m, v891, v893-int32(1))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L46
	} else {
		goto L292
	}
L277:
	;
	v909 = v904 + int32(1)
	v910 = int32(*(*int8)(unsafe.Add(mBase, uint32(v904))))
	v911 = F___isspace(m, v910)
	mBase = m.M
	if v911 != 0 {
		v904 = v909
		goto L277
	} else {
		goto L279
	}
L278:
	;
	v912 = int32(1)
	switch v910&int32(255) - int32(43) {
	case 0:
		v918 = v912
		goto L281
	default:
		v920 = v910
		v921 = v904
		v922 = v912
		goto L280
	case 2:
		goto L282
	}
L279:
	;
	goto L278
L280:
	;
	v923 = int32(0)
	v925 = v920 - int32(48)
	if base.Ui32(v925) <= base.Ui32(int32(9)) {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	v919 = int32(*(*int8)(unsafe.Add(mBase, uint32(v909))))
	v920 = v919
	v921 = v909
	v922 = v918
	goto L280
L282:
	;
	v918 = int32(0)
	goto L281
L283:
	;
	v928 = v923
	v929 = v925
	v930 = v921
	goto L286
L284:
	;
	v942 = v923
	goto L285
L285:
	;
	if v922 != 0 {
		goto L289
	} else {
		goto L290
	}
L286:
	;
	v932 = int32(10)
	v934 = v928*v932 - v929
	v935 = int32(*(*int8)(unsafe.Add(mBase, uint32(v930)+1)))
	v939 = v935 - int32(48)
	if base.Ui32(v939) < base.Ui32(v932) {
		v928 = v934
		v929 = v939
		v930 = v930 + int32(1)
		goto L286
	} else {
		goto L288
	}
L287:
	;
	v942 = v934
	goto L285
L288:
	;
	goto L287
L289:
	;
	v948 = int32(0) - v942
	goto L291
L290:
	;
	v948 = v942
	goto L291
L291:
	;
	goto L276
L292:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v952)+12))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	v957 = F_get_index_am_oid(m, v956)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L46
	} else {
		goto L293
	}
L293:
	;
	v960 = F_list_copy_tail(m, v952, int32(1))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L46
	} else {
		goto L294
	}
L294:
	;
	v963 = F_get_opfamily_oid(m, v957, v960, int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L46
	} else {
		goto L295
	}
L295:
	;
	v965 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+492)) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v27)+488)) = v963
	*(*int32)(unsafe.Add(mBase, uint32(v27)+484)) = int32(2753)
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)+4))
	if v971 == v965 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v1012 = base.I64_extend_i32_u(v963)
	v1014 = base.I64_extend16_s(base.I64_extend_i32_u(v948))
	if l1 == int32(2) {
		goto L309
	} else {
		goto L310
	}
L297:
	;
	v974 = int64(0)
	v1006 = int32(0)
	v1008 = v949
	v1010 = v974
	v1011 = v974
	goto L296
L298:
	;
	goto L299
L299:
	;
	v977 = int64(0)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v971)+4))
	if v978 <= int32(0) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1006 = v1000
	v1008 = v1002
	v1010 = v1004
	v1011 = base.I64_extend_i32_u(v1001)
	goto L296
L301:
	;
	v981 = int32(0)
	v1000 = v981
	v1001 = v981
	v1002 = v949
	v1004 = v977
	goto L300
L302:
	;
	goto L303
L303:
	;
	v984 = v27 + int32(472)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v971)+12))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v986)))
	F_get_object_address_type(m, v984, int32(50), v987, l5)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L46
	} else {
		goto L304
	}
L304:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v27)+476))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v971)+4))
	if v991 < int32(2) {
		v1000 = v987
		v1001 = v990
		v1002 = v949
		v1004 = v977
		goto L300
	} else {
		goto L305
	}
L305:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v971)+12))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)+4))
	F_get_object_address_type(m, v984, int32(50), v996, l5)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L46
	} else {
		goto L306
	}
L306:
	;
	v999 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+476)))
	v1000 = v987
	v1001 = v990
	v1002 = v996
	v1004 = v999
	goto L300
L307:
	;
	if l5 != 0 {
		goto L327
	} else {
		goto L328
	}
L308:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+16))
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061)+22)))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1061+v1062)))
	F_ReleaseCatCache(m, v1059)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L46
	} else {
		goto L326
	}
L309:
	;
	v1019 = F_SearchSysCache4(m, int32(4), v1012, v1011, v1010, v1014)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L46
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1054 = F_SearchSysCache4(m, int32(5), v1012, v1011, v1010, v1014)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L46
	} else {
		goto L324
	}
L312:
	;
	if v1019 != 0 {
		v1059 = v1019
		v1060 = int32(2602)
		goto L308
	} else {
		goto L313
	}
L313:
	;
	if l5 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1882 = int32(0)
	v1884 = int32(2602)
	goto L7
L315:
	;
	goto L316
L316:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L46
	} else {
		goto L317
	}
L317:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L46
	} else {
		goto L318
	}
L318:
	;
	v1030 = F_TypeNameToString(m, v1006)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L46
	} else {
		goto L319
	}
L319:
	;
	v1032 = F_TypeNameToString(m, v1008)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L46
	} else {
		goto L320
	}
L320:
	;
	v1037 = F_getObjectDescription(m, v27+int32(484), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L46
	} else {
		goto L321
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+220)) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(v27)+216)) = v1032
	*(*int32)(unsafe.Add(mBase, uint32(v27)+212)) = v1030
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v948
	F_errmsg(m, int32(_a_F_get_object_address_34), v27+int32(208))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L46
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1755), int32(_a_F_get_object_address_35))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L46
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	if v1054 == int32(0) {
		goto L307
	} else {
		goto L325
	}
L325:
	;
	v1059 = v1054
	v1060 = int32(2603)
	goto L308
L326:
	;
	v1882 = v1064
	v1884 = v1060
	goto L7
L327:
	;
	v1882 = int32(0)
	v1884 = int32(2603)
	goto L7
L328:
	;
	goto L329
L329:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L46
	} else {
		goto L330
	}
L330:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L46
	} else {
		goto L331
	}
L331:
	;
	v1076 = F_TypeNameToString(m, v1006)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L46
	} else {
		goto L332
	}
L332:
	;
	v1078 = F_TypeNameToString(m, v1008)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L46
	} else {
		goto L333
	}
L333:
	;
	v1083 = F_getObjectDescription(m, v27+int32(484), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L46
	} else {
		goto L334
	}
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+236)) = v1083
	*(*int32)(unsafe.Add(mBase, uint32(v27)+232)) = v1078
	*(*int32)(unsafe.Add(mBase, uint32(v27)+228)) = v1076
	*(*int32)(unsafe.Add(mBase, uint32(v27)+224)) = v948
	F_errmsg(m, int32(_a_F_get_object_address_36), v27+int32(224))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L46
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1786), int32(_a_F_get_object_address_35))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L46
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	m.G0 = v1104 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1130
	v1137 = F_LargeObjectExists(m, v1130)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L46
	} else {
		goto L345
	}
L338:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1130 = v1129
	goto L337
L339:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L46
	} else {
		goto L342
	}
L340:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1110 = int32(0)
	v1113 = F_uint32in_subr(m, v1109, v1110, int32(_a_F_get_object_address_37), v1110)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L46
	} else {
		goto L341
	}
L341:
	;
	v1130 = v1113
	goto L337
L342:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1104))) = v1119
	F_errmsg_internal(m, int32(_a_F_get_object_address_38), v1104)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L46
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_get_object_address_39), int32(280), int32(_a_F_get_object_address_40))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L46
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	if v1137|l5 != 0 {
		v1696 = v1099
		goto L18
	} else {
		goto L346
	}
L346:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L46
	} else {
		goto L347
	}
L347:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L46
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+240)) = v1130
	F_errmsg(m, int32(_a_F_get_object_address_41), v27+int32(240))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L46
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1063), int32(_a_F_get_object_address_42))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L46
	} else {
		goto L350
	}
L350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L351:
	;
	v1164 = F_LookupTypeNameOid(m, v1159, l5)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L46
	} else {
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2605)
	v1168 = m.G0
	v1170 = v1168 - int32(16)
	m.G0 = v1170
	v1175 = int64(0)
	v1177 = F_GetSysCacheOid(m, int32(12), base.I64_extend_i32_u(v1162), base.I64_extend_i32_u(v1164), v1175, v1175)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L46
	} else {
		goto L353
	}
L353:
	;
	if v1177|l5 == int32(0) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L46
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	m.G0 = v1170 + int32(16)
	v1668 = int32(2605)
	v1683 = v1177
	goto L19
L357:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L46
	} else {
		goto L358
	}
L358:
	;
	v1189 = F_format_type_be(m, v1162)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L46
	} else {
		goto L359
	}
L359:
	;
	v1191 = F_format_type_be(m, v1164)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L46
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1170)+4)) = v1191
	*(*int32)(unsafe.Add(mBase, uint32(v1170))) = v1189
	F_errmsg(m, int32(_a_F_get_object_address_43), v1170)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L46
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_get_object_address_44), int32(1243), int32(_a_F_get_object_address_45))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L46
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	v1213 = F_get_language_oid(m, v1208, l5)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L46
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3576)
	v1217 = F_get_transform_oid(m, v1211, v1213, l5)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L46
	} else {
		goto L365
	}
L365:
	;
	v1668 = int32(3576)
	v1683 = v1217
	goto L19
L366:
	;
	v1668 = v1219
	v1683 = v1222
	goto L19
L367:
	;
	v1668 = v1224
	v1683 = v1227
	goto L19
L368:
	;
	v1668 = v1229
	v1683 = v1232
	goto L19
L369:
	;
	v1668 = v1234
	v1683 = v1237
	goto L19
L370:
	;
	v1369 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1369
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1365
	v1373 = int32(1418)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1373
	v1926 = v1365
	v1927 = v1369
	v1929 = v1373
	goto L4
L371:
	;
	v1310 = F_GetForeignServerByName(m, v1241, int32(1))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L46
	} else {
		goto L394
	}
L372:
	;
	if v1268-v1269 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L373:
	;
	goto L372
L374:
	;
	v1253 = v1243
	v1254 = v1244
	goto L375
L375:
	;
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254)+1)))
	v1258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+1)))
	if v1258 == int32(0) {
		v1268 = v1258
		v1269 = v1257
		goto L373
	} else {
		goto L377
	}
L376:
	;
	v1268 = v1258
	v1269 = v1257
	goto L373
L377:
	;
	v1261 = int32(1)
	if v1258 == v1257 {
		v1253 = v1253 + v1261
		v1254 = v1254 + v1261
		goto L375
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	v1308 = int64(0)
	goto L371
L380:
	;
	goto L381
L381:
	;
	v1276 = F_SearchSysCache1(m, int32(10), base.I64_extend_i32_u(v1243))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L46
	} else {
		goto L382
	}
L382:
	;
	if v1276 == int32(0) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	if l5 != 0 {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	goto L385
L385:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+16))
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300)+22)))
	v1303 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1300+v1301))))
	F_ReleaseCatCache(m, v1276)
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L46
	} else {
		goto L393
	}
L386:
	;
	v1365 = int32(0)
	goto L370
L387:
	;
	goto L388
L388:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L46
	} else {
		goto L389
	}
L389:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L46
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1241
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1243
	F_errmsg(m, int32(_a_F_get_object_address_46), v27+int32(288))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L46
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1834), int32(_a_F_get_object_address_47))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L46
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	v1308 = v1303
	goto L371
L394:
	;
	if v1310 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	if l5 != 0 {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	goto L397
L397:
	;
	v1334 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1310))))
	v1335 = F_SearchSysCache2(m, int32(84), v1308, v1334)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L46
	} else {
		goto L405
	}
L398:
	;
	v1365 = int32(0)
	goto L370
L399:
	;
	goto L400
L400:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L46
	} else {
		goto L401
	}
L401:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L46
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+256)) = v1241
	F_errmsg(m, int32(_a_F_get_object_address_48), v27+int32(256))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L46
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1848), int32(_a_F_get_object_address_47))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L46
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	if v1335 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	if l5 != 0 {
		goto L409
	} else {
		goto L410
	}
L407:
	;
	goto L408
L408:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+16))
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+22)))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v1359+v1360)))
	F_ReleaseCatCache(m, v1335)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L46
	} else {
		goto L416
	}
L409:
	;
	v1365 = int32(0)
	goto L370
L410:
	;
	goto L411
L411:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L46
	} else {
		goto L412
	}
L412:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L46
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+276)) = v1241
	*(*int32)(unsafe.Add(mBase, uint32(v27)+272)) = v1243
	F_errmsg(m, int32(_a_F_get_object_address_46), v27+int32(272))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L46
	} else {
		goto L414
	}
L414:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1860), int32(_a_F_get_object_address_47))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L46
	} else {
		goto L415
	}
L415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L416:
	;
	v1365 = v1362
	goto L370
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1400
	v1405 = int32(_a_F_get_object_address_49)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1405
	v1926 = v1400
	v1927 = v1379
	v1929 = v1405
	goto L4
L418:
	;
	if v1383 == int32(0) {
		v1400 = v1379
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v1387 = F_GetPublicationByName(m, v1378, l5)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L46
	} else {
		goto L420
	}
L420:
	;
	if v1387 == int32(0) {
		v1400 = v1379
		goto L417
	} else {
		goto L421
	}
L421:
	;
	v1393 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1387))))
	v1394 = int64(0)
	v1396 = F_GetSysCacheOid(m, int32(50), base.I64_extend_i32_u(v1383), v1393, v1394, v1394)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L46
	} else {
		goto L422
	}
L422:
	;
	if l5 != 0 {
		v1400 = v1396
		goto L417
	} else {
		goto L423
	}
L423:
	;
	if v1396 == int32(0) {
		goto L12
	} else {
		goto L424
	}
L424:
	;
	v1400 = v1396
	goto L417
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1501
	v1508 = int32(_a_F_get_object_address_50)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1508
	v1926 = v1501
	v1927 = v1502
	v1929 = v1508
	goto L4
L426:
	;
	v1413 = F_relation_openrv_extended(m, v1410, int32(1), l5)
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L46
	} else {
		goto L427
	}
L427:
	;
	if v1413 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1415)+4))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+4))
	v1418 = F_GetPublicationByName(m, v1417, l5)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L46
	} else {
		goto L432
	}
L429:
	;
	goto L430
L430:
	;
	v1499 = int32(0)
	v1501 = v1499
	v1502 = v1499
	goto L425
L431:
	;
	F_relation_close(m, v1413, int32(1))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L46
	} else {
		goto L451
	}
L432:
	;
	if v1418 == int32(0) {
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1423 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1413)+56)))
	v1424 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1418))))
	v1425 = F_SearchSysCache2(m, int32(53), v1423, v1424)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L46
	} else {
		goto L434
	}
L434:
	;
	if v1425 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	if l5 != 0 {
		goto L431
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1425)+16))
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1453)+22)))
	v1455 = v1453 + v1454
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1455)+12)))
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1455)))
	F_ReleaseCatCache(m, v1425)
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L46
	} else {
		goto L444
	}
L438:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L46
	} else {
		goto L439
	}
L439:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L46
	} else {
		goto L440
	}
L440:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+48))
	v1438 = v1436 + int32(4)
	if l1 == int32(31) {
		goto L11
	} else {
		goto L441
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+340)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(v27)+336)) = v1438
	F_errmsg(m, int32(_a_F_get_object_address_51), v27+int32(336))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L46
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1932), int32(_a_F_get_object_address_52))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L46
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	v1463 = base.B2i32(v1456&int32(1) == int32(0))
	if v1463&base.B2i32(l1 == int32(31)) != 0 {
		goto L10
	} else {
		goto L445
	}
L445:
	;
	if base.B2i32(l1 != int32(33))|v1463 != 0 {
		v1501 = v1457
		v1502 = v1413
		goto L425
	} else {
		goto L446
	}
L446:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L46
	} else {
		goto L447
	}
L447:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L46
	} else {
		goto L448
	}
L448:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+356)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(v27)+352)) = v1477 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_53), v27+int32(352))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L46
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1957), int32(_a_F_get_object_address_52))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L46
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	goto L430
L452:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1511)+8))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+4))
	v1519 = v1518
	goto L454
L453:
	;
	v1519 = int32(0)
	goto L454
L454:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v1512)+4))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+4))
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1523))))
	switch v1524 - int32(76) {
	case 0:
		goto L458
	default:
		goto L457
	case 7:
		goto L456
	case 8:
		goto L460
	case 26:
		goto L461
	case 34:
		goto L459
	case 38:
		v1562 = int32(_a_F_get_object_address_54)
		goto L455
	}
L455:
	;
	v1565 = F_SearchSysCache1(m, int32(10), base.I64_extend_i32_u(v1520))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L46
	} else {
		goto L468
	}
L456:
	;
	v1562 = int32(_a_F_get_object_address_55)
	goto L455
L457:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L46
	} else {
		goto L462
	}
L458:
	;
	v1562 = int32(_a_F_get_object_address_56)
	goto L455
L459:
	;
	v1562 = int32(_a_F_get_object_address_57)
	goto L455
L460:
	;
	v1562 = int32(_a_F_get_object_address_58)
	goto L455
L461:
	;
	v1562 = int32(_a_F_get_object_address_59)
	goto L455
L462:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L46
	} else {
		goto L463
	}
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+416)) = base.I32_extend8_s(v1524)
	F_errmsg(m, int32(_a_F_get_object_address_60), v27+int32(416))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L46
	} else {
		goto L464
	}
L464:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+400)) = int64(326417514606)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+392)) = int64(360777252966)
	*(*int64)(unsafe.Add(mBase, uint32(v27)+384)) = int64(356482285682)
	F_errhint(m, int32(_a_F_get_object_address_61), v27+int32(384))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L46
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(2068), int32(_a_F_get_object_address_62))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L46
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	if l5 != 0 {
		v1871 = int32(0)
		goto L8
	} else {
		goto L480
	}
L468:
	;
	if v1565 == int32(0) {
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+16))
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569)+22)))
	v1572 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1569+v1570))))
	F_ReleaseCatCache(m, v1565)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L46
	} else {
		goto L470
	}
L470:
	;
	if v1519 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v1587 = F_SearchSysCache3(m, int32(22), v1572, base.I64_extend_i32_u(v1583), base.I64_extend_i32_u(v1524))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L46
	} else {
		goto L477
	}
L472:
	;
	v1583 = int32(0)
	goto L471
L473:
	;
	goto L474
L474:
	;
	v1579 = F_get_namespace_oid(m, v1519, int32(1))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L46
	} else {
		goto L475
	}
L475:
	;
	if v1579 == int32(0) {
		goto L467
	} else {
		goto L476
	}
L476:
	;
	v1583 = v1579
	goto L471
L477:
	;
	if v1587 == int32(0) {
		goto L467
	} else {
		goto L478
	}
L478:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1587)+16))
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1591)+22)))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1591+v1592)))
	F_ReleaseCatCache(m, v1587)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L46
	} else {
		goto L479
	}
L479:
	;
	v1871 = v1594
	goto L8
L480:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L46
	} else {
		goto L481
	}
L481:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L46
	} else {
		goto L482
	}
L482:
	;
	if v1519 != 0 {
		goto L9
	} else {
		goto L483
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+436)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(v27)+432)) = v1520
	F_errmsg(m, int32(_a_F_get_object_address_63), v27+int32(432))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L46
	} else {
		goto L484
	}
L484:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(2120), int32(_a_F_get_object_address_62))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L46
	} else {
		goto L485
	}
L485:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L486:
	;
	v1668 = v1621
	v1683 = v1624
	goto L19
L487:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1629 = v1628
	goto L21
L488:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L46
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l1
	F_errmsg_internal(m, int32(_a_F_get_object_address_11), v27)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L46
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1143), int32(_a_F_get_object_address_42))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L46
	} else {
		goto L491
	}
L491:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L492:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v27)+488))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1651)+4))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+4))
	v1654 = int32(2606)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1654
	v1657 = F_get_domain_constraint_oid(m, v1650, v1653, l5)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L46
	} else {
		goto L493
	}
L493:
	;
	v1668 = v1654
	v1683 = v1657
	goto L19
L494:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L46
	} else {
		goto L495
	}
L495:
	;
	F_errmsg(m, int32(_a_F_get_object_address_64), int32(0))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L46
	} else {
		goto L496
	}
L496:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1523), int32(_a_F_get_object_address_65))
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L46
	} else {
		goto L497
	}
L497:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L498:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L46
	} else {
		goto L499
	}
L499:
	;
	v1737 = F_NameListToString(m, v252)
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L46
	} else {
		goto L500
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+132)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v249
	F_errmsg(m, int32(_a_F_get_object_address_66), v27+int32(128))
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L46
	} else {
		goto L501
	}
L501:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1538), int32(_a_F_get_object_address_65))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L46
	} else {
		goto L502
	}
L502:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L503:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L46
	} else {
		goto L504
	}
L504:
	;
	F_errmsg(m, int32(_a_F_get_object_address_64), int32(0))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L46
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1576), int32(_a_F_get_object_address_67))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L46
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L46
	} else {
		goto L508
	}
L508:
	;
	v1775 = F_NameListToString(m, v289)
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L46
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+148)) = v1775
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v286
	F_errmsg(m, int32(_a_F_get_object_address_68), v27+int32(144))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L46
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1596), int32(_a_F_get_object_address_67))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L46
	} else {
		goto L511
	}
L511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L512:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L46
	} else {
		goto L513
	}
L513:
	;
	F_errmsg(m, int32(_a_F_get_object_address_69), int32(0))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L46
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1447), int32(_a_F_get_object_address_18))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L46
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L46
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1378
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1382
	F_errmsg(m, int32(_a_F_get_object_address_70), v27+int32(304))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L46
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(2001), int32(_a_F_get_object_address_71))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L46
	} else {
		goto L519
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1927), int32(_a_F_get_object_address_52))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L46
	} else {
		goto L521
	}
L521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L522:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L46
	} else {
		goto L523
	}
L523:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+372)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(v27)+368)) = v1843 + int32(4)
	F_errmsg(m, int32(_a_F_get_object_address_72), v27+int32(368))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L46
	} else {
		goto L524
	}
L524:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(1952), int32(_a_F_get_object_address_52))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L46
	} else {
		goto L525
	}
L525:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L526:
	;
	F_errfinish(m, int32(_a_F_get_object_address_4), int32(2115), int32(_a_F_get_object_address_62))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L46
	} else {
		goto L527
	}
L527:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L528:
	;
	if v50 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L529:
	;
	goto L3
L530:
	;
	if v1929 == int32(1259) {
		goto L567
	} else {
		goto L568
	}
L531:
	;
	if base.B2i32(v1929 != v50)|base.B2i32(v1926 != v52) == int32(0) {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v53 == v1953 {
		goto L529
	} else {
		goto L535
	}
L533:
	;
	goto L534
L534:
	;
	if v50 == int32(1259) {
		goto L530
	} else {
		goto L536
	}
L535:
	;
	goto L534
L536:
	;
	v1959 = int32(1)
	if v50 <= int32(3591) {
		goto L541
	} else {
		goto L542
	}
L537:
	;
	if v2030 != 0 {
		goto L562
	} else {
		goto L563
	}
L538:
	;
	goto L537
L539:
	;
	v2030 = int32(0)
	goto L538
L540:
	;
	if base.B2i32(base.Ui32(v50-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v50-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v2030 = v1959
		goto L538
	} else {
		goto L561
	}
L541:
	;
	if v50 <= int32(2670) {
		goto L544
	} else {
		goto L545
	}
L542:
	;
	goto L543
L543:
	;
	if v50 <= int32(_a_F_get_object_address_73) {
		goto L551
	} else {
		goto L552
	}
L544:
	;
	switch v50 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v2030 = v1959
		goto L538
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L539
	default:
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	v1971 = v50 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v1971))|base.B2i32(int32(1)<<(uint(v1971)%32)&int32(226492515) == int32(0)) != 0 {
		goto L540
	} else {
		goto L549
	}
L547:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v50-int32(2396)) {
		goto L539
	} else {
		goto L548
	}
L548:
	;
	v2030 = v1959
	goto L538
L549:
	;
	v2030 = v1959
	goto L538
L550:
	;
	if base.Ui32(v50-int32(3592)) < base.Ui32(int32(2)) {
		v2030 = v1959
		goto L538
	} else {
		goto L559
	}
L551:
	;
	v1984 = v50 - int32(_a_F_get_object_address_74)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v1984))|base.B2i32(int32(1)<<(uint(v1984)%32)&int32(963) == int32(0)) != 0 {
		goto L550
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	switch v50 - int32(_a_F_get_object_address_25) {
	case 0, 1, 2, 3, 4, 59, 60:
		v2030 = v1959
		goto L538
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L539
	default:
		goto L555
	}
L554:
	;
	v2030 = v1959
	goto L538
L555:
	;
	if base.Ui32(v50-int32(_a_F_get_object_address_75)) < base.Ui32(int32(3)) {
		v2030 = v1959
		goto L538
	} else {
		goto L556
	}
L556:
	;
	v2001 = v50 - int32(_a_F_get_object_address_27)
	if base.Ui32(int32(15)) < base.Ui32(v2001) {
		goto L539
	} else {
		goto L557
	}
L557:
	;
	if int32(1)<<(uint(v2001)%32)&int32(_a_F_get_object_address_76) != 0 {
		v2030 = v1959
		goto L538
	} else {
		goto L558
	}
L558:
	;
	goto L539
L559:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v50-int32(4060)) {
		goto L539
	} else {
		goto L560
	}
L560:
	;
	v2030 = v1959
	goto L538
L561:
	;
	goto L539
L562:
	;
	F_UnlockSharedObject(m, v50, v52, l4)
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L46
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	F_UnlockDatabaseObject(m, v50, v52, l4)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L46
	} else {
		goto L566
	}
L565:
	;
	goto L530
L566:
	;
	goto L530
L567:
	;
	if v1927 != 0 {
		goto L529
	} else {
		goto L599
	}
L568:
	;
	v2039 = int32(1)
	if v1929 <= int32(3591) {
		goto L573
	} else {
		goto L574
	}
L569:
	;
	if v2110 != 0 {
		goto L594
	} else {
		goto L595
	}
L570:
	;
	goto L569
L571:
	;
	v2110 = int32(0)
	goto L570
L572:
	;
	if base.B2i32(base.Ui32(v1929-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v1929-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v2110 = v2039
		goto L570
	} else {
		goto L593
	}
L573:
	;
	if v1929 <= int32(2670) {
		goto L576
	} else {
		goto L577
	}
L574:
	;
	goto L575
L575:
	;
	if v1929 <= int32(_a_F_get_object_address_73) {
		goto L583
	} else {
		goto L584
	}
L576:
	;
	switch v1929 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v2110 = v2039
		goto L570
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L571
	default:
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	v2051 = v1929 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v2051))|base.B2i32(int32(1)<<(uint(v2051)%32)&int32(226492515) == int32(0)) != 0 {
		goto L572
	} else {
		goto L581
	}
L579:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1929-int32(2396)) {
		goto L571
	} else {
		goto L580
	}
L580:
	;
	v2110 = v2039
	goto L570
L581:
	;
	v2110 = v2039
	goto L570
L582:
	;
	if base.Ui32(v1929-int32(3592)) < base.Ui32(int32(2)) {
		v2110 = v2039
		goto L570
	} else {
		goto L591
	}
L583:
	;
	v2064 = v1929 - int32(_a_F_get_object_address_74)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v2064))|base.B2i32(int32(1)<<(uint(v2064)%32)&int32(963) == int32(0)) != 0 {
		goto L582
	} else {
		goto L586
	}
L584:
	;
	goto L585
L585:
	;
	switch v1929 - int32(_a_F_get_object_address_25) {
	case 0, 1, 2, 3, 4, 59, 60:
		v2110 = v2039
		goto L570
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L571
	default:
		goto L587
	}
L586:
	;
	v2110 = v2039
	goto L570
L587:
	;
	if base.Ui32(v1929-int32(_a_F_get_object_address_75)) < base.Ui32(int32(3)) {
		v2110 = v2039
		goto L570
	} else {
		goto L588
	}
L588:
	;
	v2081 = v1929 - int32(_a_F_get_object_address_27)
	if base.Ui32(int32(15)) < base.Ui32(v2081) {
		goto L571
	} else {
		goto L589
	}
L589:
	;
	if int32(1)<<(uint(v2081)%32)&int32(_a_F_get_object_address_76) != 0 {
		v2110 = v2039
		goto L570
	} else {
		goto L590
	}
L590:
	;
	goto L571
L591:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1929-int32(4060)) {
		goto L571
	} else {
		goto L592
	}
L592:
	;
	v2110 = v2039
	goto L570
L593:
	;
	goto L571
L594:
	;
	F_LockSharedObject(m, v1929, v1926, l4)
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L46
	} else {
		goto L597
	}
L595:
	;
	goto L596
L596:
	;
	F_LockDatabaseObject(m, v1929, v1926, l4)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L46
	} else {
		goto L598
	}
L597:
	;
	goto L567
L598:
	;
	goto L567
L599:
	;
	v2116 = *(*int64)(unsafe.Add(mBase, _c_F_get_object_address[0]))
	if v55 == v2116 {
		goto L529
	} else {
		goto L600
	}
L600:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = v1929
	v50 = v1929
	v52 = v1926
	v53 = v2118
	v55 = v2116
	goto L2
L601:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1927
	goto L1
}
func F_get_object_attnum_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_oid[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+20)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_oid[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_oid_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_oid[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_attnum_oid_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_attnum_oid_5), int32(2827), int32(_a_F_get_object_attnum_oid_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_catcache_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_catcache_name[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_catcache_name[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_catcache_name_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_catcache_name[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_catcache_name_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_catcache_name_5), int32(2827), int32(_a_F_get_object_catcache_name_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_field_end(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v12 < v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = int32(1)
	v16 = v11 - v15
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v16))))
	if v18 != v15 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v16<<(uint(int32(2))%32))))
	if v27 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(v32 == int32(0))|base.B2i32(v32 != v35) != 0 {
		v53 = v32
		v54 = v35
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v53-v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v38 = l1
	v39 = v27
	goto L9
L9:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v43
		v54 = v42
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v53 = v43
	v54 = v42
	goto L7
L11:
	;
	v46 = int32(1)
	if v43 == v42 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v11 < v12 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v14))) = uint8(v58)
	return v58
L15:
	;
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v62 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v73
	goto L1
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v66 != 0 {
		v73 = int32(0)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v69 = F_cstring_to_text_with_len(m, v62, v67-v62)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	return int32(0)
L24:
	;
	v73 = v69
	goto L18
}
func F_get_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 != 0 {
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v5 != 0 {
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v6
		}
	}
	return int32(0)
}
func F_get_object_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_type[0]))
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L27
	} else {
		goto L29
	}
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
	if v55 != int32(42) {
		v72 = v55
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v13 == l0 {
		v51 = v12
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v20 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_type[0])) = v47
	v51 = v47
	goto L2
L8:
	;
	v47 = v23 + int32(_a_F_get_object_type_0)
	goto L7
L9:
	;
	v47 = v23 + int32(_a_F_get_object_type_1)
	goto L7
L10:
	;
	v47 = v23 + int32(_a_F_get_object_type_2)
	goto L7
L11:
	;
	v23 = v20 * int32(40)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[1])))
	if l0 != v24 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v47 = v23 + int32(_a_F_get_object_type_3)
	goto L7
L13:
	;
	if v20 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[2])))
	if v30 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[3])))
	if v32 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_get_object_type[4])))
	if v34 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v20 = v20 + int32(4)
	goto L11
L20:
	;
	m.G0 = v9 + int32(16)
	return v72
L21:
	;
	v59 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v72 = int32(42)
	goto L20
L23:
	;
	v72 = int32(18)
	goto L20
L24:
	;
	v72 = int32(23)
	goto L20
L25:
	;
	v72 = int32(52)
	goto L20
L26:
	;
	v72 = int32(38)
	goto L20
L27:
	;
	return int32(0)
L28:
	;
	switch v59&int32(255) - int32(73) {
	case 0, 32:
		v72 = int32(20)
		goto L20
	default:
		goto L22
	case 10:
		goto L26
	case 29:
		goto L23
	case 36:
		goto L24
	case 45:
		goto L25
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_type_4), v9)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_get_object_type_5), int32(2827), int32(_a_F_get_object_type_6))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_object_ownercheck(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v14 = F_superuser_arg(m, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			if l0 == int32(2613) {
				v23 = int32(2995)
			} else {
				v23 = l0
			}
			v24 = F_get_object_catcache_oid(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 != int32(-1) {
					v29 = F_SearchSysCache1(m, v24, base.I64_extend_i32_u(l1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								v94 = F_get_object_class_descr(m, v23)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v94
									F_errmsg_internal(m, int32(_a_F_object_ownercheck_0), v11+int32(16))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_object_ownercheck_1), int32(_a_F_object_ownercheck_2), int32(_a_F_object_ownercheck_3))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
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
							v33 = F_get_object_attnum_owner(m, v23)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = F_SysCacheGetAttrNotNull(m, v24, v29, v33)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									F_ReleaseCatCache(m, v29)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v77 = v35
										v79 = F_has_privs_of_role(m, l2, base.I32_wrap_i64(v77))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v82 = v79
											m.G0 = v11 + int32(96)
											return v82
										}
									}
								}
							}
						}
					}
				} else {
					v40 = F_table_open(m, v23, int32(1))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v43 = v11 + int32(32)
						v44 = F_get_object_attnum_oid(m, v23)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_ScanKeyInit(m, v43, v44, int32(3), int32(184), base.I64_extend_i32_u(l1))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v51 = F_get_object_oid_index(m, v23)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									v53 = int32(1)
									v56 = F_systable_beginscan(m, v40, v51, v53, int32(0), v53, v43)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = F_systable_getnext(m, v56)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											if v58 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													v112 = F_get_object_class_descr(m, v23)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v112
														F_errmsg_internal(m, int32(_a_F_object_ownercheck_4), v11)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_object_ownercheck_1), int32(_a_F_object_ownercheck_5), int32(_a_F_object_ownercheck_3))
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
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
												v62 = F_get_object_attnum_owner(m, v23)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
													v67 = F_heap_getattr_2(m, v58, v62, v64, v11+int32(31))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														F_systable_endscan(m, v56)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v40, int32(1))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int32(0)
															} else {
																v77 = v67
																v79 = F_has_privs_of_role(m, l2, base.I32_wrap_i64(v77))
																mBase = m.M
																v80 = m.ExcPending
																if v80 != 0 {
																	return int32(0)
																} else {
																	v82 = v79
																	m.G0 = v11 + int32(96)
																	return v82
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
		} else {
			v82 = int32(1)
			m.G0 = v11 + int32(96)
			return v82
		}
	}
}
func F_parse_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v76
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = m.T0[v6].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v15 + int32(1)
	v19 = F_json_lex(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	if v12 != 0 {
		v76 = v12
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	if v19 != 0 {
		v76 = v19
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v21 - int32(1) {
	case 0:
		goto L13
	default:
		goto L12
	case 3:
		goto L11
	}
L11:
	;
	v63 = F_json_lex(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L38
	}
L12:
	;
	v50 = int32(11)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v53 != 0 {
		goto L32
	} else {
		goto L33
	}
L13:
	;
	v24 = F_parse_object_field(m, l0, l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v24 != 0 {
		v76 = v24
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L16
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v30 != int32(7) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v76 = v46
	goto L3
L18:
	;
	if v30 == int32(4) {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v44 = F_json_lex(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L28
	}
L21:
	;
	v35 = int32(11)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v39 = int32(13)
	goto L24
L23:
	;
	v39 = v35
	goto L24
L24:
	;
	if v30 == int32(12) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v42 = v35
	goto L27
L26:
	;
	v42 = v39
	goto L27
L27:
	;
	return v42
L28:
	;
	if v44 != 0 {
		v76 = v44
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v46 = F_parse_object_field(m, l0, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v46 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	v54 = int32(12)
	goto L34
L33:
	;
	v54 = v50
	goto L34
L34:
	;
	if v21 == int32(12) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v57 = v50
	goto L37
L36:
	;
	v57 = v54
	goto L37
L37:
	;
	return v57
L38:
	;
	if v63 != 0 {
		v76 = v63
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v65 - int32(1)
	if v5 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v70 = m.T0[v5].(func(*base.Module, int32) int32)(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v76 = int32(0)
	goto L3
L43:
	;
	if v70 != 0 {
		v76 = v70
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L42
}
