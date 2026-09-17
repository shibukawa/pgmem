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
		F_errmsg_internal(m, int32(_a_F_ExecBitmapAnd_0), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_ExecBitmapAnd_1), int32(44), int32(_a_F_ExecBitmapAnd_2))
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
		F_errmsg_internal(m, int32(_a_F_ExecBitmapIndexScan_0), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_ExecBitmapIndexScan_1), int32(40), int32(_a_F_ExecBitmapIndexScan_2))
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
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
	var v169 int32
	_ = v169
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
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
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
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
	var v301 int32
	_ = v301
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
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
	var v372 int32
	_ = v372
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
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 float64
	_ = v489
	var v490 float64
	_ = v490
	var v495 int32
	_ = v495
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v538 int32
	_ = v538
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v579 float64
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 float64
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 float64
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
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 float64
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 float64
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
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
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v831 float64
	_ = v831
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v853 float64
	_ = v853
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 float64
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v915 int32
	_ = v915
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
	m.G0 = v899 + int32(128)
	return v915
L2:
	;
	v894 = F_create_bitmap_and_path(m, l0, l1, v880)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L6
	} else {
		goto L176
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
	v878 = v22
	v880 = v4
	goto L2
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v899 = v22
	v915 = v40
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
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v45 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v53 = int32(0)
	v59 = v4
	v62 = v4
	goto L16
L14:
	;
	v538 = v4
	goto L15
L15:
	;
	if v538 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v62<<(uint(int32(2))%32))))
	v74 = F_palloc(m, int32(20))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v538 = v515
	goto L15
L18:
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
		goto L19
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v85 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v525 = v62 + int32(1)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v525 < v526 {
		v53 = v509
		v59 = v515
		v62 = v525
		goto L16
	} else {
		goto L91
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v88 = v86
	goto L23
L22:
	;
	v88 = int32(0)
	goto L23
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v89 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v92 = v90
	goto L26
L25:
	;
	v92 = int32(0)
	goto L26
L26:
	;
	if v88+v92 <= int32(100) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v85 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v495 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+16)) = uint8(v495)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43+v59<<(uint(int32(2))%32)))) = v74
	v509 = v53
	v515 = v59 + v495
	goto L20
L30:
	;
	if v219 == int32(0) {
		v346 = v215
		v349 = v218
		goto L49
	} else {
		goto L50
	}
L31:
	;
	v215 = v53
	v218 = int32(0)
	v219 = v89
	goto L30
L32:
	;
	goto L33
L33:
	;
	v99 = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v101 <= v99 {
		v215 = v53
		v218 = v99
		v219 = v89
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v108 = v53
	v111 = v99
	v113 = v99
	goto L35
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v113<<(uint(int32(2))%32))))
	v128 = int32(0)
	if v108 == v128 {
		v169 = v128
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v215 = v189
	v218 = v204
	v219 = v210
	goto L30
L37:
	;
	v204 = F_bms_add_member(m, v111, v190)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L47
	}
L38:
	;
	v183 = F_lappend(m, v108, v127)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L46
	}
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v131 <= int32(0) {
		v169 = v128
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v139 = v128
	goto L41
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153+v139<<(uint(int32(2))%32))))
	v158 = F_equal(m, v127, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	v169 = v161
	goto L38
L43:
	;
	if v158 != 0 {
		v189 = v108
		v190 = v139
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v161 = v139 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v161 < v162 {
		v139 = v161
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v189 = v183
	v190 = v169
	goto L37
L47:
	;
	v207 = v113 + int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v207 < v208 {
		v108 = v189
		v111 = v204
		v113 = v207
		goto L35
	} else {
		goto L48
	}
L48:
	;
	goto L36
L49:
	;
	v361 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+16)) = uint8(v361)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v349
	if v361 < v59 {
		goto L67
	} else {
		goto L68
	}
L50:
	;
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v233 <= v232 {
		v346 = v215
		v349 = v218
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v240 = v215
	v243 = v218
	v245 = v232
	goto L52
L52:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v245<<(uint(int32(2))%32))))
	v260 = int32(0)
	if v240 == v260 {
		v301 = v260
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v346 = v321
	v349 = v336
	goto L49
L54:
	;
	v336 = F_bms_add_member(m, v243, v322)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L6
	} else {
		goto L64
	}
L55:
	;
	v315 = F_lappend(m, v240, v259)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v263 <= int32(0) {
		v301 = v260
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v271 = v260
	goto L58
L58:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v271<<(uint(int32(2))%32))))
	v290 = F_equal(m, v259, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L60
	}
L59:
	;
	v301 = v293
	goto L55
L60:
	;
	if v290 != 0 {
		v321 = v240
		v322 = v271
		goto L54
	} else {
		goto L61
	}
L61:
	;
	v293 = v271 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v293 < v294 {
		v271 = v293
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	v321 = v315
	v322 = v301
	goto L54
L64:
	;
	v339 = v245 + int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v339 < v340 {
		v240 = v321
		v243 = v336
		v245 = v339
		goto L52
	} else {
		goto L65
	}
L65:
	;
	goto L53
L66:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	F_cost_bitmap_tree_node(m, v474, v22+int32(48), v22+int32(32))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L6
	} else {
		goto L88
	}
L67:
	;
	v372 = v361
	goto L70
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v59<<(uint(int32(2))%32)))) = v74
	v509 = v346
	v515 = v59 + int32(1)
	goto L20
L70:
	;
	v388 = v43 + v372<<(uint(int32(2))%32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+16)))
	if v390 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	v395 = int32(0)
	if base.B2i32(v393 == v395)|base.B2i32(v394 == v395) != 0 {
		v441 = base.B2i32(v393|v394 == v395)
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	v447 = v372 + int32(1)
	if v447 != v59 {
		v372 = v447
		goto L70
	} else {
		goto L87
	}
L75:
	;
	if v441 != 0 {
		goto L66
	} else {
		goto L86
	}
L76:
	;
	goto L75
L77:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v409 != v410 {
		v441 = int32(0)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v412 = int32(1)
	if v409 <= v412 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v415 = v412
	goto L81
L80:
	;
	v415 = v409
	goto L81
L81:
	;
	v416 = int32(8)
	v421 = int32(0)
	goto L82
L82:
	;
	v429 = v421 << (uint(int32(2)) % 32)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v393+v416+v429)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v394+v416+v429)))
	v434 = base.B2i32(v431 == v433)
	if v431 != v433 {
		v441 = v434
		goto L76
	} else {
		goto L84
	}
L83:
	;
	v441 = v434
	goto L76
L84:
	;
	v437 = v421 + int32(1)
	if v437 != v415 {
		v421 = v437
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L74
L87:
	;
	goto L71
L88:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	F_cost_bitmap_tree_node(m, v482, v22+int32(40), v22+int32(24))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	v489 = *(*float64)(unsafe.Add(mBase, uint32(v22)+48))
	v490 = *(*float64)(unsafe.Add(mBase, uint32(v22)+40))
	if base.F64_lt(v489, v490) == int32(0) {
		v509 = v346
		v515 = v59
		goto L20
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v74
	v509 = v346
	v515 = v59
	goto L20
L91:
	;
	goto L17
L92:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v899 = v22
	v915 = v550
	goto L1
L93:
	;
	goto L94
L94:
	;
	F_pg_qsort(m, v43, v538, int32(4), int32(824))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	if v538 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v558 = F_create_bitmap_and_path(m, l0, l1, int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v560 = int32(0)
	v567 = v560
	v575 = v560
	v579 = float64(0)
	goto L100
L99:
	;
	v899 = v22
	v915 = v558
	goto L1
L100:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v43+v575<<(uint(int32(2))%32))))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v585
	v591 = F_list_make1_impl(m, int32(1), v22+int32(12))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L102
	}
L101:
	;
	if v862 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L102:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(1477468750106)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = l1
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v597
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v593)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v593
	v601 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v601
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v601
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v599 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v599)+4))
	v611 = v609
	goto L105
L104:
	;
	v611 = int32(0)
	goto L105
L105:
	;
	v612 = F_get_loop_count(m, l0, v608, v611)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	F_cost_bitmap_heap_scan(m, v22+int32(48), l0, l1, v599, v593, v612)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v616 = *(*float64)(unsafe.Add(mBase, uint32(v22)+104))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v584)+8))
	v619 = F_list_concat_copy(m, v617, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v584)+12))
	v622 = F_bms_copy(m, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v625 = v575 + int32(1)
	if v625 < v538 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v633 = v591
	v634 = v619
	v636 = v625
	v638 = v622
	v643 = v616
	goto L113
L111:
	;
	v843 = v591
	v844 = v619
	v853 = v616
	goto L112
L112:
	;
	F_list_free(m, v844)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L6
	} else {
		goto L163
	}
L113:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v43+v636<<(uint(int32(2))%32))))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)+12))
	v651 = int32(0)
	if base.B2i32(v650 == v651)|base.B2i32(v638 == v651) != 0 {
		v696 = v651
		goto L117
	} else {
		goto L118
	}
L114:
	;
	v843 = v821
	v844 = v822
	v853 = v831
	goto L112
L115:
	;
	v835 = v636 + int32(1)
	if v835 != v538 {
		v633 = v821
		v634 = v822
		v636 = v835
		v638 = v826
		v643 = v831
		goto L113
	} else {
		goto L162
	}
L116:
	;
	if v696 != 0 {
		v821 = v633
		v822 = v634
		v826 = v638
		v831 = v643
		goto L115
	} else {
		goto L129
	}
L117:
	;
	goto L116
L118:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v638)+4))
	if v661 < v662 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v664 = v661
	goto L121
L120:
	;
	v664 = v662
	goto L121
L121:
	;
	if v664 <= int32(1) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v667 = int32(1)
	goto L124
L123:
	;
	v667 = v664
	goto L124
L124:
	;
	v668 = int32(8)
	v673 = int32(0)
	goto L125
L125:
	;
	v680 = v673 << (uint(int32(2)) % 32)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v638+v668+v680)))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v650+v668+v680)))
	v685 = v682 & v684
	v687 = base.B2i32(v685 != int32(0))
	if v685 != 0 {
		v696 = v687
		goto L117
	} else {
		goto L127
	}
L126:
	;
	v696 = v687
	goto L117
L127:
	;
	v689 = v673 + int32(1)
	if v689 != v667 {
		v673 = v689
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v649)+8))
	if v697 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v762 = F_lappend(m, v633, v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L139
	}
L131:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	if v700 <= int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v708 = int32(0)
	goto L133
L133:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v697)+12))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v723+v708<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v727
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v727
	v733 = F_list_make1_impl(m, int32(1), v22+int32(8))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L135
	}
L134:
	;
	goto L130
L135:
	;
	v736 = F_predicate_implied_by(m, v733, v634, int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	if v736 != 0 {
		v821 = v633
		v822 = v634
		v826 = v638
		v831 = v643
		goto L115
	} else {
		goto L137
	}
L137:
	;
	v739 = v708 + int32(1)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	if v739 < v740 {
		v708 = v739
		goto L133
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	v764 = F_create_bitmap_and_path(m, l0, l1, v762)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(1477468750106)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v769
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v764)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v764
	v773 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v773
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v773
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v771 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v771)+4))
	v783 = v781
	goto L143
L142:
	;
	v783 = int32(0)
	goto L143
L143:
	;
	v784 = F_get_loop_count(m, l0, v780, v783)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_cost_bitmap_heap_scan(m, v22+int32(48), l0, l1, v771, v764, v784)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v788 = *(*float64)(unsafe.Add(mBase, uint32(v22)+104))
	if base.F64_gt(v643, v788) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v649)+4))
	v791 = F_list_concat(m, v634, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	if v762 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v649)+8))
	v794 = F_list_concat(m, v791, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v649)+12))
	v797 = F_bms_add_members(m, v638, v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	v821 = v762
	v822 = v794
	v826 = v797
	v831 = v788
	goto L115
L152:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v762)+4))
	v803 = v799 - int32(1)
	goto L154
L153:
	;
	v803 = int32(-1)
	goto L154
L154:
	;
	v804 = int32(0)
	if base.B2i32(v762 == v804)|base.B2i32(v803 <= v804) != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v821 = v814
	v822 = v634
	v826 = v638
	v831 = v643
	goto L115
L156:
	;
	v814 = int32(0)
	goto L158
L157:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v762)+4))
	if v803 < v811 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L155
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v762)+4)) = v803
	goto L161
L160:
	;
	goto L161
L161:
	;
	v814 = v762
	goto L158
L162:
	;
	goto L114
L163:
	;
	v861 = base.B2i32(v575 == int32(0)) | base.F64_lt(v853, v579)
	if v861 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v862 = v843
	goto L166
L165:
	;
	v862 = v567
	goto L166
L166:
	;
	if v861 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v863 = v853
	goto L169
L168:
	;
	v863 = v579
	goto L169
L169:
	;
	if v625 != v538 {
		v567 = v862
		v575 = v625
		v579 = v863
		goto L100
	} else {
		goto L170
	}
L170:
	;
	goto L101
L171:
	;
	v868 = F_create_bitmap_and_path(m, l0, l1, int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L6
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	if v870 != int32(1) {
		v878 = v22
		v880 = v862
		goto L2
	} else {
		goto L175
	}
L174:
	;
	v899 = v22
	v915 = v868
	goto L1
L175:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v862)+12))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	v899 = v22
	v915 = v874
	goto L1
L176:
	;
	v899 = v878
	v915 = v894
	goto L1
}
