package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBitmapAnd(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(245988), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(494612), int32(44), int32(425127))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_ExecBitmapHeapScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(700), int32(701))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ExecBitmapIndexScan(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(245865), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(491769), int32(40), int32(282283))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_choose_bitmap_and(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v361 int32
	_ = v361
	var v370 int32
	_ = v370
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v488 float64
	_ = v488
	var v489 float64
	_ = v489
	var v494 int32
	_ = v494
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 float64
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 float64
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 float64
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
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v642 float64
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
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
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 float64
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 float64
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v827 float64
	_ = v827
	var v831 int32
	_ = v831
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v849 float64
	_ = v849
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 float64
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v911 int32
	_ = v911
	v4 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	if l2 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v897 + int32(128)
	return v911
L2:
	;
	v890 = F_create_bitmap_and_path(m, l0, l1, v884)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L6
	} else {
		goto L179
	}
L3:
	;
	v27 = F_palloc(m, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v36 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return int32(0)
L7:
	;
	F_pg_qsort(m, v27, int32(0), int32(4), int32(824))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v876 = v22
	v884 = v4
	goto L2
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v897 = v22
	v911 = v40
	goto L1
L10:
	;
	goto L11
L11:
	;
	v43 = F_palloc(m, v36<<(uint(int32(2))%32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 <= v45 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v535 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L14:
	;
	v535 = v4
	goto L13
L15:
	;
	goto L16
L16:
	;
	v53 = v45
	v57 = v4
	v63 = v4
	goto L17
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v63<<(uint(int32(2))%32))))
	v74 = F_palloc(m, int32(20))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	v535 = v512
	goto L13
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v72
	v82 = v74 + int32(8)
	F_find_indexpath_quals(m, v72, v74+int32(4), v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v86 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v524 = v63 + int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v524 < v525 {
		v53 = v508
		v57 = v512
		v63 = v524
		goto L17
	} else {
		goto L93
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = v87
	goto L24
L23:
	;
	v88 = int32(0)
	goto L24
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v92 = v90
	goto L27
L26:
	;
	v92 = int32(0)
	goto L27
L27:
	;
	if v88+v92 <= int32(100) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v86 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v494 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+16)) = uint8(v494)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43+v57<<(uint(int32(2))%32)))) = v74
	v508 = v53
	v512 = v57 + v494
	goto L21
L31:
	;
	if v222 == int32(0) {
		v346 = v215
		v349 = v218
		goto L50
	} else {
		goto L51
	}
L32:
	;
	v215 = v53
	v218 = int32(0)
	v222 = v89
	goto L31
L33:
	;
	goto L34
L34:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v100 <= v99 {
		v215 = v53
		v218 = v99
		v222 = v89
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v108 = v53
	v111 = v99
	v116 = int32(0)
	goto L36
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v116<<(uint(int32(2))%32))))
	v128 = int32(0)
	if v108 == v128 {
		v167 = v128
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v215 = v189
	v218 = v204
	v222 = v210
	goto L31
L38:
	;
	v204 = F_bms_add_member(m, v111, v188)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L48
	}
L39:
	;
	v183 = F_lappend(m, v108, v127)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L47
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v131 <= int32(0) {
		v167 = v128
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v137 = v128
	goto L42
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v137<<(uint(int32(2))%32))))
	v158 = F_equal(m, v127, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L44
	}
L43:
	;
	v167 = v161
	goto L39
L44:
	;
	if v158 != 0 {
		v188 = v137
		v189 = v108
		goto L38
	} else {
		goto L45
	}
L45:
	;
	v161 = v137 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v161 < v162 {
		v137 = v161
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	v188 = v167
	v189 = v183
	goto L38
L48:
	;
	v207 = v116 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v207 < v208 {
		v108 = v189
		v111 = v204
		v116 = v207
		goto L36
	} else {
		goto L49
	}
L49:
	;
	goto L37
L50:
	;
	v361 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+16)) = uint8(v361)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v349
	if v361 < v57 {
		goto L68
	} else {
		goto L69
	}
L51:
	;
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v233 <= v232 {
		v346 = v215
		v349 = v218
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v240 = v215
	v243 = v218
	v248 = v232
	goto L53
L53:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v248<<(uint(int32(2))%32))))
	v260 = int32(0)
	if v240 == v260 {
		v299 = v260
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v346 = v321
	v349 = v336
	goto L50
L55:
	;
	v336 = F_bms_add_member(m, v243, v320)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L65
	}
L56:
	;
	v315 = F_lappend(m, v240, v259)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L64
	}
L57:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v263 <= int32(0) {
		v299 = v260
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v269 = v260
	goto L59
L59:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v269<<(uint(int32(2))%32))))
	v290 = F_equal(m, v259, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L61
	}
L60:
	;
	v299 = v293
	goto L56
L61:
	;
	if v290 != 0 {
		v320 = v269
		v321 = v240
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v293 = v269 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v293 < v294 {
		v269 = v293
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v320 = v299
	v321 = v315
	goto L55
L65:
	;
	v339 = v248 + int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v339 < v340 {
		v240 = v321
		v243 = v336
		v248 = v339
		goto L53
	} else {
		goto L66
	}
L66:
	;
	goto L54
L67:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	F_cost_bitmap_tree_node(m, v473, v22+int32(48), v22+int32(32))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L6
	} else {
		goto L90
	}
L68:
	;
	v370 = v361
	goto L71
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v57<<(uint(int32(2))%32)))) = v74
	v508 = v346
	v512 = v57 + int32(1)
	goto L21
L71:
	;
	v388 = v43 + v370<<(uint(int32(2))%32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+16)))
	if v390 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v395 = int32(0)
	v402 = base.B2i32(v393|v394 == v395)
	if v393 == v395 {
		v441 = v402
		goto L77
	} else {
		goto L78
	}
L74:
	;
	goto L75
L75:
	;
	v446 = v370 + int32(1)
	if v446 != v57 {
		v370 = v446
		goto L71
	} else {
		goto L89
	}
L76:
	;
	if v441 != 0 {
		goto L67
	} else {
		goto L88
	}
L77:
	;
	goto L76
L78:
	;
	if v394 == int32(0) {
		v441 = v402
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v408 != v409 {
		v441 = int32(0)
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v411 = int32(1)
	if v408 <= v411 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v414 = v411
	goto L83
L82:
	;
	v414 = v408
	goto L83
L83:
	;
	v415 = int32(8)
	v420 = int32(0)
	goto L84
L84:
	;
	v428 = v420 << (uint(int32(2)) % 32)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v393+v415+v428)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428+(v394+v415))))
	v433 = base.B2i32(v430 == v432)
	if v432 != v430 {
		v441 = v433
		goto L77
	} else {
		goto L86
	}
L85:
	;
	v441 = v433
	goto L77
L86:
	;
	v436 = v420 + int32(1)
	if v436 != v414 {
		v420 = v436
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	goto L75
L89:
	;
	goto L72
L90:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	F_cost_bitmap_tree_node(m, v481, v22+int32(40), v22+int32(24))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	v488 = *(*float64)(unsafe.Add(mBase, uint32(v22)+48))
	v489 = *(*float64)(unsafe.Add(mBase, uint32(v22)+40))
	if base.F64_lt(v488, v489) == int32(0) {
		v508 = v346
		v512 = v57
		goto L21
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v74
	v508 = v346
	v512 = v57
	goto L21
L93:
	;
	goto L18
L94:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v897 = v22
	v911 = v549
	goto L1
L95:
	;
	goto L96
L96:
	;
	F_pg_qsort(m, v43, v535, int32(4), int32(824))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	if v535 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v557 = F_create_bitmap_and_path(m, l0, l1, int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L6
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v559 = int32(0)
	v574 = v559
	v575 = v559
	v578 = float64(0)
	goto L102
L101:
	;
	v897 = v22
	v911 = v557
	goto L1
L102:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v43+v575<<(uint(int32(2))%32))))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v584
	v590 = F_list_make1_impl(m, int32(1), v22+int32(12))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L104
	}
L103:
	;
	if v858 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L104:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(1477468750106)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = l1
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v596
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v592)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v592
	v600 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v600
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v598 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	v610 = v609
	goto L107
L106:
	;
	v610 = v600
	goto L107
L107:
	;
	v611 = F_get_loop_count(m, l0, v608, v610)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	F_cost_bitmap_heap_scan(m, v22+int32(48), l0, l1, v598, v592, v611)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v615 = *(*float64)(unsafe.Add(mBase, uint32(v22)+104))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v583)+8))
	v618 = F_list_concat_copy(m, v616, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	v621 = F_bms_copy(m, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	v624 = v575 + int32(1)
	if v624 < v535 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v632 = v621
	v633 = v618
	v636 = v590
	v638 = v624
	v642 = v615
	goto L115
L113:
	;
	v840 = v618
	v843 = v590
	v849 = v615
	goto L114
L114:
	;
	F_list_free(m, v840)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L6
	} else {
		goto L166
	}
L115:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v43+v638<<(uint(int32(2))%32))))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+12))
	v650 = int32(0)
	if v649 == v650 {
		v691 = v650
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v840 = v818
	v843 = v821
	v849 = v827
	goto L114
L117:
	;
	v831 = v638 + int32(1)
	if v831 != v535 {
		v632 = v817
		v633 = v818
		v636 = v821
		v638 = v831
		v642 = v827
		goto L115
	} else {
		goto L165
	}
L118:
	;
	if v691 != 0 {
		v817 = v632
		v818 = v633
		v821 = v636
		v827 = v642
		goto L117
	} else {
		goto L132
	}
L119:
	;
	goto L118
L120:
	;
	if v632 == int32(0) {
		v691 = v650
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v659 < v660 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v662 = v659
	goto L124
L123:
	;
	v662 = v660
	goto L124
L124:
	;
	if v662 <= int32(1) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v665 = int32(1)
	goto L127
L126:
	;
	v665 = v662
	goto L127
L127:
	;
	v666 = int32(8)
	v671 = int32(0)
	goto L128
L128:
	;
	v678 = v671 << (uint(int32(2)) % 32)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v632+v666+v678)))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v678+(v649+v666))))
	v683 = v680 & v682
	v685 = base.B2i32(v683 != int32(0))
	if v683 != 0 {
		v691 = v685
		goto L119
	} else {
		goto L130
	}
L129:
	;
	v691 = v685
	goto L119
L130:
	;
	v687 = v671 + int32(1)
	if v687 != v665 {
		v671 = v687
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v648)+8))
	if v695 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	v760 = F_lappend(m, v636, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L6
	} else {
		goto L142
	}
L134:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v698 <= int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v706 = int32(0)
	goto L136
L136:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v695)+12))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v721+v706<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v725
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v725
	v731 = F_list_make1_impl(m, int32(1), v22+int32(8))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L6
	} else {
		goto L138
	}
L137:
	;
	goto L133
L138:
	;
	v734 = F_predicate_implied_by(m, v731, v633, int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	if v734 != 0 {
		v817 = v632
		v818 = v633
		v821 = v636
		v827 = v642
		goto L117
	} else {
		goto L140
	}
L140:
	;
	v737 = v706 + int32(1)
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v737 < v738 {
		v706 = v737
		goto L136
	} else {
		goto L141
	}
L141:
	;
	goto L137
L142:
	;
	v762 = F_create_bitmap_and_path(m, l0, l1, v760)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(1477468750106)
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v767
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v762)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v762
	v771 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v771
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v769 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	v781 = v780
	goto L146
L145:
	;
	v781 = v771
	goto L146
L146:
	;
	v782 = F_get_loop_count(m, l0, v779, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	F_cost_bitmap_heap_scan(m, v22+int32(48), l0, l1, v769, v762, v782)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v786 = *(*float64)(unsafe.Add(mBase, uint32(v22)+104))
	if base.F64_gt(v642, v786) != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	v789 = F_list_concat(m, v633, v788)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L6
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	if v760 != 0 {
		goto L155
	} else {
		goto L156
	}
L152:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v648)+8))
	v792 = F_list_concat(m, v789, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v648)+12))
	v795 = F_bms_add_members(m, v632, v794)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	v817 = v795
	v818 = v792
	v821 = v760
	v827 = v786
	goto L117
L155:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	v801 = v797 - int32(1)
	goto L157
L156:
	;
	v801 = int32(-1)
	goto L157
L157:
	;
	v802 = int32(0)
	if v760 == v802 {
		v810 = v802
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v817 = v632
	v818 = v633
	v821 = v810
	v827 = v642
	goto L117
L159:
	;
	goto L158
L160:
	;
	if v801 <= int32(0) {
		v810 = v802
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v801 < v807 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760)+4)) = v801
	goto L164
L163:
	;
	goto L164
L164:
	;
	v810 = v760
	goto L159
L165:
	;
	goto L116
L166:
	;
	v857 = base.B2i32(v575 == int32(0)) | base.F64_lt(v849, v578)
	if v857 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v858 = v843
	goto L169
L168:
	;
	v858 = v574
	goto L169
L169:
	;
	if v857 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v859 = v849
	goto L172
L171:
	;
	v859 = v578
	goto L172
L172:
	;
	if v624 != v535 {
		v574 = v858
		v575 = v624
		v578 = v859
		goto L102
	} else {
		goto L173
	}
L173:
	;
	goto L103
L174:
	;
	v864 = F_create_bitmap_and_path(m, l0, l1, int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L6
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v866 != int32(1) {
		v876 = v22
		v884 = v858
		goto L2
	} else {
		goto L178
	}
L177:
	;
	v897 = v22
	v911 = v864
	goto L1
L178:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v869)))
	v897 = v22
	v911 = v870
	goto L1
L179:
	;
	v897 = v876
	v911 = v890
	goto L1
}
