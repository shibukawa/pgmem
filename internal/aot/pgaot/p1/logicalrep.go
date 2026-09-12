package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_launcher_onexit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(0)
	return
}
func F_logicalrep_message_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	switch l0 - int32(65) {
	case 0:
		v38 = int32(515443)
		m.G0 = v6 + int32(16)
		return v38
	case 1:
		v38 = int32(528498)
		m.G0 = v6 + int32(16)
		return v38
	case 2:
		v38 = int32(519037)
		m.G0 = v6 + int32(16)
		return v38
	case 3:
		v38 = int32(536313)
		m.G0 = v6 + int32(16)
		return v38
	case 4:
		v38 = int32(524764)
		m.G0 = v6 + int32(16)
		return v38
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v30 = int32(4399152)
		v34 = F_pg_snprintf(m, v30, int32(20), int32(663034), v6)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = v30
			m.G0 = v6 + int32(16)
			return v38
		}
	case 8:
		v38 = int32(515676)
		m.G0 = v6 + int32(16)
		return v38
	case 10:
		v38 = int32(541829)
		m.G0 = v6 + int32(16)
		return v38
	case 12:
		v38 = int32(539657)
		m.G0 = v6 + int32(16)
		return v38
	case 14:
		v38 = int32(528463)
		m.G0 = v6 + int32(16)
		return v38
	case 15:
		v38 = int32(537650)
		m.G0 = v6 + int32(16)
		return v38
	case 17:
		v38 = int32(527687)
		m.G0 = v6 + int32(16)
		return v38
	case 18:
		v38 = int32(515747)
		m.G0 = v6 + int32(16)
		return v38
	case 19:
		v38 = int32(537001)
		m.G0 = v6 + int32(16)
		return v38
	case 20:
		v38 = int32(536979)
		m.G0 = v6 + int32(16)
		return v38
	case 24:
		v38 = int32(537802)
		m.G0 = v6 + int32(16)
		return v38
	case 33:
		v38 = int32(537629)
		m.G0 = v6 + int32(16)
		return v38
	case 34:
		v38 = int32(519030)
		m.G0 = v6 + int32(16)
		return v38
	case 47:
		v38 = int32(537643)
		m.G0 = v6 + int32(16)
		return v38
	case 49:
		v38 = int32(541845)
		m.G0 = v6 + int32(16)
		return v38
	}
}
func F_logicalrep_rel_close(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_sequence_close(m, v3, l1)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
		return
	}
}
func F_logicalrep_worker_find(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	if v8 <= v3 {
		v40 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v40
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v18 = v3
	goto L3
L3:
	;
	v23 = v12 + int32(16) + v18*int32(112)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)))
	if v24 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v40 = int32(0)
	goto L1
L5:
	;
	v35 = v18 + int32(1)
	if v35 != v8 {
		v18 = v35
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 == int32(3) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	if v30 != l0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	if v32 != l1 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v40 = v23
	goto L1
L10:
	;
	goto L4
}
func F_logicalrep_worker_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int64
	_ = v318
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v475 int32
	_ = v475
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int64
	_ = v532
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int64
	_ = v548
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v700 int64
	_ = v700
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v804 int64
	_ = v804
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(1568)
	m.G0 = v21
	v25 = F_errstart(m, int32(14), v8)
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
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = l3
	F_errmsg_internal(m, int32(685780), v21+int32(96))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[651]))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	F_errfinish(m, int32(493694), int32(338), int32(325027))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v21 + int32(1568)
	return v868
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v49 = F_LWLockAcquire(m, v45+int32(5504), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L154
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	v61 = v8
	v62 = v52
	v68 = v8
	goto L14
L13:
	;
	if l0 != int32(1) {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	if v62 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	v347 = v332
	v350 = v337
	goto L13
L16:
	;
	v227 = m.G0
	v228 = int32(16)
	v229 = v227 - v228
	m.G0 = v229
	F___gettimeofday(m, v229)
	mBase = m.M
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	v233 = int64(*(*int32)(unsafe.Add(mBase, uint32(v229)+8)))
	m.G0 = v229 + v228
	v241 = v233 + v232*int64(1000000) - int64(946684800000000)
	goto L42
L17:
	;
	v213 = v61
	v218 = int32(0)
	v220 = v68
	goto L16
L18:
	;
	goto L19
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v82 = int32(0)
	goto L20
L20:
	;
	v99 = v76 + int32(16) + v82*int32(112)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+16)))
	if v100 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v113 = v111 + int32(16)
	if v62 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v107 = v99
	v108 = v82
	goto L22
L24:
	;
	goto L25
L25:
	;
	v104 = v82 + int32(1)
	if v104 != v62 {
		v82 = v104
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v107 = v61
	v108 = v68
	goto L22
L27:
	;
	v123 = v109
	v132 = int32(0)
	v133 = v109
	goto L30
L28:
	;
	v174 = v109
	v184 = v109
	goto L29
L29:
	;
	if v62&int32(1) == int32(0) {
		v213 = v107
		v218 = v184
		v220 = v108
		goto L16
	} else {
		goto L39
	}
L30:
	;
	v140 = v113 + v123*int32(112)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+16)))
	if v141 != int32(1) {
		v150 = v133
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v174 = v167
	v184 = v165
	goto L29
L32:
	;
	v151 = int32(1)
	v155 = v113 + (v123|v151)*int32(112)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+16)))
	if v156 != v151 {
		v165 = v150
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v144 != int32(1) {
		v150 = v133
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+32))
	v150 = v133 + base.B2i32(v147 == l2)
	goto L32
L35:
	;
	v166 = int32(2)
	v167 = v123 + v166
	v169 = v132 + v166
	if v169 != v62&int32(2147483646) {
		v123 = v167
		v132 = v169
		v133 = v165
		goto L30
	} else {
		goto L38
	}
L36:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v159 != int32(1) {
		v165 = v150
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	v165 = v150 + base.B2i32(v162 == l2)
	goto L35
L38:
	;
	goto L31
L39:
	;
	v195 = v113 + v174*int32(112)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+16)))
	if v196 != int32(1) {
		v213 = v107
		v218 = v184
		v220 = v108
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	if v199 != int32(1) {
		v213 = v107
		v218 = v184
		v220 = v108
		goto L16
	} else {
		goto L41
	}
L41:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v195)+32))
	v213 = v107
	v218 = v184 + base.B2i32(v202 == l2)
	v220 = v108
	goto L16
L42:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	v244 = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	if base.B2i32(v213 == v244)|base.B2i32(v247 <= v218) == v244 {
		v347 = v243
		v350 = v247
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v252 = int32(0)
	if v243 <= v252 {
		v347 = v243
		v350 = v247
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v258 = v252
	v265 = int32(0)
	goto L45
L45:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v277 = v274 + v258*int32(112)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+32)))
	if v278 != int32(1) {
		v328 = v265
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v328&int32(1) != 0 {
		v61 = v213
		v62 = v332
		v68 = v220
		goto L14
	} else {
		goto L59
	}
L47:
	;
	v330 = v258 + int32(1)
	v332 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	if v330 < v332 {
		v258 = v330
		v265 = v328
		goto L45
	} else {
		goto L58
	}
L48:
	;
	v282 = v277 + int32(16)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
	if v283 != 0 {
		v328 = v265
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v282)+8))
	v286 = *(*int32)(unsafe.Add(mBase, _consts[653]))
	goto L50
L50:
	;
	if base.B2i32(base.I64_extend_i32_s(v286)*int64(1000) <= v241-v284) == int32(0) {
		v328 = v265
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v295 = v277 + int32(36)
	v298 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v298 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v282)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v300
	F_errmsg_internal(m, int32(453402), v21+int32(80))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v312 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+16)) = uint8(v312)
	*(*int32)(unsafe.Add(mBase, uint32(v282))) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v312
	v318 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v295)+8)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v318
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+68)) = uint8(v312)
	*(*int32)(unsafe.Add(mBase, uint32(v282)+64)) = int32(-1)
	v328 = int32(1)
	goto L47
L56:
	;
	F_errfinish(m, int32(493694), int32(393), int32(325027))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	goto L46
L59:
	;
	goto L15
L60:
	;
	v366 = int32(0)
	if v347 <= v366 {
		v475 = v366
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v218 < v350 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v361+int32(5504))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v868 = int32(0)
	goto L8
L64:
	;
	if l0 != int32(3) {
		goto L81
	} else {
		goto L82
	}
L65:
	;
	v369 = int32(1)
	v371 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v375 = v373 + int32(16)
	if v347 != v369 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v384 = v371
	v390 = int32(0)
	v392 = v366
	goto L69
L67:
	;
	v435 = v371
	v443 = v366
	goto L68
L68:
	;
	if v347&v369 == int32(0) {
		v475 = v443
		goto L64
	} else {
		goto L78
	}
L69:
	;
	v401 = v375 + v384*int32(112)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+16)))
	if v402 != int32(1) {
		v411 = v392
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v435 = v428
	v443 = v426
	goto L68
L71:
	;
	v412 = int32(1)
	v416 = v375 + (v384|v412)*int32(112)
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)))
	if v417 != v412 {
		v426 = v411
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	if v405 != int32(3) {
		v411 = v392
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401)+32))
	v411 = v392 + base.B2i32(v408 == l2)
	goto L71
L74:
	;
	v427 = int32(2)
	v428 = v384 + v427
	v430 = v390 + v427
	if v430 != v347&int32(2147483646) {
		v384 = v428
		v390 = v430
		v392 = v426
		goto L69
	} else {
		goto L77
	}
L75:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	if v420 != int32(3) {
		v426 = v411
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)+32))
	v426 = v411 + base.B2i32(v423 == l2)
	goto L74
L77:
	;
	goto L70
L78:
	;
	v454 = v375 + v435*int32(112)
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454)+16)))
	if v455 != int32(1) {
		v475 = v443
		goto L64
	} else {
		goto L79
	}
L79:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v458 != int32(3) {
		v475 = v443
		goto L64
	} else {
		goto L80
	}
L80:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v454)+32))
	v475 = v443 + base.B2i32(v461 == l2)
	goto L64
L81:
	;
	if v213 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	if v475 < v485 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v489+int32(5504))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v868 = int32(0)
	goto L8
L85:
	;
	v496 = int32(0)
	v498 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v498+int32(5504))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v526 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)) = uint8(v526)
	*(*int64)(unsafe.Add(mBase, uint32(v213)+8)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = l0
	v530 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+60)) = v530
	v532 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v213)+48)) = v532
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+40)) = uint8(v530)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v213)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v213)+28)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v213)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v213)+20)) = v530
	v542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+18)))
	v544 = v542 + v526
	*(*uint16)(unsafe.Add(mBase, uint32(v213)+18)) = uint16(v544)
	v547 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v548 = int64(-9223372036854775807 - 1)
	*(*int64)(unsafe.Add(mBase, uint32(v213)+104)) = v548
	*(*int64)(unsafe.Add(mBase, uint32(v213)+96)) = v532
	*(*int64)(unsafe.Add(mBase, uint32(v213)+88)) = v548
	*(*int64)(unsafe.Add(mBase, uint32(v213)+80)) = v548
	*(*int64)(unsafe.Add(mBase, uint32(v213)+72)) = v532
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+68)) = uint8(base.B2i32(l0 == int32(3)))
	if l0 != int32(3) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v505 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v505 == int32(0) {
		v868 = v496
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(118444), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(133470)
	F_errhint(m, int32(649736), v21)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(493694), int32(438), int32(325027))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v868 = v496
	goto L8
L95:
	;
	v562 = int32(-1)
	goto L97
L96:
	;
	v562 = v547
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+64)) = v562
	v565 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v565+int32(5504))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v575 = F__emscripten_memset_bulkmem(m, v21+int32(104), base.I32_extend8_s(int32(0)), int32(1460))
	mBase = m.M
	goto L99
L99:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+296)) = int64(8589934595)
	v583 = F_pg_snprintf(m, v21+int32(308), int32(1024), int32(161350), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	switch v585 {
	case 0:
		goto L103
	case 1:
		goto L104
	case 2:
		goto L102
	case 3:
		goto L105
	default:
		goto L101
	}
L101:
	;
	v671 = v213 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+304)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1428)) = v220
	v676 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1560)) = v676
	v682 = F_RegisterDynamicBackgroundWorker(m, v21+int32(104), v21+int32(100))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L118
	}
L102:
	;
	v652 = F_pg_snprintf(m, v21+int32(1332), int32(96), int32(277884), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L115
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L112
	}
L104:
	;
	v615 = F_pg_snprintf(m, v21+int32(1332), int32(96), int32(277953), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L109
	}
L105:
	;
	v591 = F_pg_snprintf(m, v21+int32(1332), int32(96), int32(277876), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = l2
	v600 = F_pg_snprintf(m, v21+int32(104), int32(96), int32(44472), v21+int32(48))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v607 = F_pg_snprintf(m, v21+int32(200), int32(96), int32(220311), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+1432)) = l6
	goto L101
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l2
	v625 = F_pg_snprintf(m, v21+int32(104), int32(96), int32(55216), v21-int32(-64))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v632 = F_pg_snprintf(m, v21+int32(200), int32(96), int32(220555), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L101
L112:
	;
	F_errmsg_internal(m, int32(367160), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(493694), int32(506), int32(325027))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l2
	v661 = F_pg_snprintf(m, v21+int32(104), int32(96), int32(44419), v21+int32(32))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v668 = F_pg_snprintf(m, v21+int32(200), int32(96), int32(219959), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	goto L101
L118:
	;
	if v682 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v686 = int32(0)
	v688 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v692 = F_LWLockAcquire(m, v688+int32(5504), v686)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	v753 = int32(0)
	goto L130
L122:
	;
	v694 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)) = uint8(v694)
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v671)+16)) = v694
	v700 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v671)+8)) = v700
	*(*int64)(unsafe.Add(mBase, uint32(v671))) = v700
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+68)) = uint8(v694)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+64)) = int32(-1)
	v709 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v709+int32(5504))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v716 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	if v716 == int32(0) {
		v868 = v686
		goto L8
	} else {
		goto L125
	}
L125:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(118484), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(160764)
	F_errhint(m, int32(649736), v21+int32(16))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(493694), int32(524), int32(325027))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v868 = v686
	goto L8
L130:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v760 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v768 = F_LWLockAcquire(m, v764+int32(5504), int32(1))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)))
	if v770 != int32(1) {
		v812 = v770
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v832 = F_WaitLatch(m, v828, int32(41), int32(10), int32(134217734))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L149
	}
L138:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v814+int32(5504))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L146
	}
L139:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	if v773 != 0 {
		v812 = v770
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v775+int32(5504))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v782 = F_GetBackgroundWorkerPid(m, v739, v21+int32(1564))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v782 != int32(2) {
		goto L137
	} else {
		goto L143
	}
L143:
	;
	v786 = int32(0)
	v788 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v792 = F_LWLockAcquire(m, v788+int32(5504), v786)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+18)))
	if v794 != v544&int32(65535) {
		v812 = v786
		goto L138
	} else {
		goto L145
	}
L145:
	;
	v798 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+16)) = uint8(v798)
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v798
	*(*int32)(unsafe.Add(mBase, uint32(v671)+16)) = v798
	v804 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v671)+8)) = v804
	*(*int64)(unsafe.Add(mBase, uint32(v671))) = v804
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+68)) = uint8(v798)
	*(*int32)(unsafe.Add(mBase, uint32(v213)+64)) = int32(-1)
	v812 = v786
	goto L138
L146:
	;
	if v753&int32(1) == int32(0) {
		v868 = v812
		goto L8
	} else {
		goto L147
	}
L147:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	F_SetLatch(m, v824)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v868 = v812
	goto L8
L149:
	;
	if v832&int32(1) == int32(0) {
		goto L130
	} else {
		goto L150
	}
L150:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	*(*int32)(unsafe.Add(mBase, uint32(v839))) = int32(0)
	goto L151
L151:
	;
	v842 = int32(1)
	v844 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v844 == int32(0) {
		v753 = v842
		goto L130
	} else {
		goto L152
	}
L152:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v753 = v842
	goto L130
L154:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(553204), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(493694), int32(344), int32(325027))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logicalrep_worker_wakeup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v7 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v11 = F_LWLockAcquire(m, v7+int32(5504), int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[535]))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v53+int32(5504))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[650]))
	v23 = int32(0)
	goto L5
L5:
	;
	v28 = v18 + int32(16) + v23*int32(112)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+16)))
	if v29 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_SetLatch(m, v38+int32(20))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L7:
	;
	goto L6
L8:
	;
	v41 = v23 + int32(1)
	if v41 != v14 {
		v23 = v41
		goto L5
	} else {
		goto L14
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v32 == int32(3) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v35 != l0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v37 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v38 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	goto L3
L15:
	;
	goto L3
L16:
	;
	return
}
