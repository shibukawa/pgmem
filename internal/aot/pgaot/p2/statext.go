package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_is_kind_built(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l1 - int32(100)
	v13 = v11 & int32(255)
	if base.B2i32(base.Ui32(int32(10)) <= base.Ui32(v13))|base.B2i32(int32(base.Ui32(int32(519))>>(uint(v13)%32))&int32(1) == v3) == v3 {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v11&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_statext_is_kind_built[0])))
		v31 = F_heap_attisnull(m, l0, v29, int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v31 ^ int32(1)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(_a_F_statext_is_kind_built_0), v8)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_statext_is_kind_built_1), int32(409), int32(_a_F_statext_is_kind_built_2))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_statext_mcv_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v774 int64
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v960 int32
	_ = v960
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	if l0 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L22
	} else {
		goto L231
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L22
	} else {
		goto L228
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L22
	} else {
		goto L225
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L22
	} else {
		goto L222
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L22
	} else {
		goto L219
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L22
	} else {
		goto L216
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v24 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v960 = v2
	goto L9
L9:
	;
	m.G0 = v22 + int32(128)
	return v960
L10:
	;
	v95 = F_palloc0(m, int32(48))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L22
	} else {
		goto L37
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	if base.Ui32(int32(17)) < base.Ui32(v47) {
		goto L10
	} else {
		goto L21
	}
L13:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v39 = int32(1)
	if v24&v39 != 0 {
		v47 = int32(base.Ui32(v24) >> (uint(v39) % 32))
		goto L12
	} else {
		goto L20
	}
L16:
	;
	v34 = int32(18)
	if v27 == v34 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v38 = v34
	goto L19
L18:
	;
	v38 = int32(2)
	goto L19
L19:
	;
	v47 = v38
	goto L12
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
	goto L12
L21:
	;
	goto L11
L22:
	;
	return int32(0)
L23:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v57 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v82
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_0), v22)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L22
	} else {
		goto L35
	}
L25:
	;
	v61 = int32(18)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v63 == v61 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v74 = int32(1)
	if v57&v74 != 0 {
		v82 = int32(base.Ui32(v57) >> (uint(v74) % 32))
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v66 = v61
	goto L30
L29:
	;
	v66 = int32(2)
	goto L30
L30:
	;
	if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v73 = int32(6)
	goto L33
L32:
	;
	v73 = v66
	goto L33
L33:
	;
	v82 = v73
	goto L24
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = int32(base.Ui32(v78) >> (uint(int32(2)) % 32))
	goto L24
L35:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1032), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L22
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v97 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v101 = v99 & v97
	if v101 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v102 = v97
	goto L40
L39:
	;
	v102 = int32(4)
	goto L40
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0+v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v104
	v107 = l0 + int32(1)
	if v101 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v110 = v107
	goto L43
L42:
	;
	v110 = l0 + int32(4)
	goto L43
L43:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v113
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v110)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v95)+12)) = uint16(v115)
	if v104 != int32(-509193790) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	if v111 != int32(1) {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	if v115 == int32(0) {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v115) {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	if v113 == int32(0) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	if base.Ui32(int32(_a_F_statext_mcv_deserialize_3)) <= base.Ui32(v113) {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v129 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v156 = v115 << (uint(int32(2)) % 32)
	v158 = v115 * int32(20)
	v167 = v156 + v158 + (v115*int32(3)+int32(16))*v113 + int32(18)
	if base.Ui32(v154) < base.Ui32(v167) {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v133 = int32(18)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v135 == v133 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v146 = int32(1)
	if v129&v146 != 0 {
		v154 = int32(base.Ui32(v129) >> (uint(v146) % 32))
		goto L50
	} else {
		goto L60
	}
L54:
	;
	v138 = v133
	goto L56
L55:
	;
	v138 = int32(2)
	goto L56
L56:
	;
	if base.Ui32((v135-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v145 = int32(6)
	goto L59
L58:
	;
	v145 = v138
	goto L59
L59:
	;
	v154 = v145
	goto L50
L60:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v154 = int32(base.Ui32(v150) >> (uint(int32(2)) % 32))
	goto L50
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L22
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v212 = v110 + int32(14)
	if v156 != 0 {
		goto L78
	} else {
		goto L79
	}
L64:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v173 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v198
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_4), v22+int32(48))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L22
	} else {
		goto L76
	}
L66:
	;
	v177 = int32(18)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v179 == v177 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v190 = int32(1)
	if v173&v190 != 0 {
		v198 = int32(base.Ui32(v173) >> (uint(v190) % 32))
		goto L65
	} else {
		goto L75
	}
L69:
	;
	v182 = v177
	goto L71
L70:
	;
	v182 = int32(2)
	goto L71
L71:
	;
	if base.Ui32((v179-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v189 = int32(6)
	goto L74
L73:
	;
	v189 = v182
	goto L74
L74:
	;
	v198 = v189
	goto L65
L75:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = int32(base.Ui32(v194) >> (uint(int32(2)) % 32))
	goto L65
L76:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1093), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L22
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	base.MemoryCopy(m, v95+int32(16), v212, v156)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v216 = v212 + v156
	v217 = F_palloc(m, v158)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L22
	} else {
		goto L81
	}
L81:
	;
	if v158 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	base.MemoryCopy(m, v217, v216, v158)
	goto L84
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(v115) < base.Ui32(int32(4)) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v334 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L86:
	;
	v287 = v266
	v288 = v267
	v293 = v2
	goto L94
L87:
	;
	v266 = int32(0)
	v267 = v167
	goto L86
L88:
	;
	goto L89
L89:
	;
	v227 = int32(0)
	v228 = v167
	v230 = v2
	goto L90
L90:
	;
	v247 = v217 + v227*int32(20)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+64))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v247)+44))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)+24))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v255 = v248 + (v249 + (v250 + (v251 + v228)))
	v256 = int32(4)
	v257 = v227 + v256
	v259 = v230 + v256
	if v259 != v115&int32(12) {
		v227 = v257
		v228 = v255
		v230 = v259
		goto L90
	} else {
		goto L92
	}
L91:
	;
	if v115&int32(3) == int32(0) {
		v317 = v255
		goto L85
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v266 = v257
	v267 = v255
	goto L86
L94:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v217+v287*int32(20))+4))
	v309 = v308 + v288
	v310 = int32(1)
	v313 = v293 + v310
	if v313 != v115&int32(3) {
		v287 = v287 + v310
		v288 = v309
		v293 = v313
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v317 = v309
	goto L85
L96:
	;
	goto L95
L97:
	;
	if v359 != v317 {
		goto L108
	} else {
		goto L109
	}
L98:
	;
	v338 = int32(18)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v340 == v338 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v351 = int32(1)
	if v334&v351 != 0 {
		v359 = int32(base.Ui32(v334) >> (uint(v351) % 32))
		goto L97
	} else {
		goto L107
	}
L101:
	;
	v343 = v338
	goto L103
L102:
	;
	v343 = int32(2)
	goto L103
L103:
	;
	if base.Ui32((v340-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v350 = int32(6)
	goto L106
L105:
	;
	v350 = v343
	goto L106
L106:
	;
	v359 = v350
	goto L97
L107:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v359 = int32(base.Ui32(v355) >> (uint(int32(2)) % 32))
	goto L97
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L22
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v405 = F_palloc(m, v156)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L22
	} else {
		goto L125
	}
L111:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v365 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v390
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_4), v22+int32(80))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L22
	} else {
		goto L123
	}
L113:
	;
	v369 = int32(18)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v371 == v369 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v382 = int32(1)
	if v365&v382 != 0 {
		v390 = int32(base.Ui32(v365) >> (uint(v382) % 32))
		goto L112
	} else {
		goto L122
	}
L116:
	;
	v374 = v369
	goto L118
L117:
	;
	v374 = int32(2)
	goto L118
L118:
	;
	if base.Ui32((v371-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v381 = int32(6)
	goto L121
L120:
	;
	v381 = v374
	goto L121
L121:
	;
	v390 = v381
	goto L112
L122:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v390 = int32(base.Ui32(v386) >> (uint(int32(2)) % 32))
	goto L112
L123:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1125), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L22
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	v410 = int32(0)
	v411 = int32(0)
	goto L126
L126:
	;
	v427 = int32(2)
	v432 = v217 + v410*int32(20)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	v436 = F_palloc(m, v433<<(uint(v427)%32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L22
	} else {
		goto L128
	}
L127:
	;
	v444 = int32(7)
	v447 = (v156 + v444) & int32(120)
	v448 = v447 * v113
	v451 = int32(24)
	v452 = (v115 + v444) & v451
	v453 = v452 * v113
	v459 = (v113*v451 + int32(55)) & int32(_a_F_statext_mcv_deserialize_5)
	v467 = F_repalloc(m, v95, v448+(v453+v459)+(v440+v444)&int32(-8))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L22
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405+v410<<(uint(v427)%32)))) = v436
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	v440 = v439 + v411
	v442 = v410 + int32(1)
	if v442 != v115 {
		v410 = v442
		v411 = v440
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v469 = v467 + v459
	v470 = v469 + v448
	v474 = v158 + v216
	v475 = int32(0)
	v480 = v470 + v453
	goto L131
L131:
	;
	v494 = v217 + v475*int32(20)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+16)))
	if v495 == int32(1) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v729 = int32(1)
	if v113 <= v729 {
		goto L179
	} else {
		goto L180
	}
L133:
	;
	v727 = v475 + int32(1)
	if v727 != v115 {
		v474 = v708
		v475 = v727
		v480 = v714
		goto L131
	} else {
		goto L178
	}
L134:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v498 <= int32(0) {
		v708 = v474
		v714 = v480
		goto L133
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if int32(0) < v561 {
		goto L152
	} else {
		goto L153
	}
L137:
	;
	v506 = v474
	v509 = int32(0)
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = int32(0)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v526 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v708 = v551
	v714 = v480
	goto L133
L140:
	;
	base.MemoryCopy(m, v22+int32(124), v506, v526)
	goto L142
L141:
	;
	goto L142
L142:
	;
	switch v526 - int32(1) {
	case 0:
		goto L144
	case 1:
		goto L147
	default:
		goto L145
	case 3:
		goto L146
	}
L143:
	;
	v551 = v506 + v526
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v405+v475<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v552+v509<<(uint(int32(2))%32)))) = v550
	v558 = v509 + int32(1)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v558 < v559 {
		v506 = v551
		v509 = v558
		goto L138
	} else {
		goto L151
	}
L144:
	;
	v549 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+124)))
	v550 = v549
	goto L143
L145:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L22
	} else {
		goto L148
	}
L146:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v550 = v533
	goto L143
L147:
	;
	v532 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+124)))
	v550 = v532
	goto L143
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v526
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_6), v22-int32(-64))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L22
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_7), int32(70), int32(_a_F_statext_mcv_deserialize_8))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L22
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	goto L139
L152:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v564 <= int32(0) {
		v708 = v474
		v714 = v480
		goto L133
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	switch v561 + int32(2) {
	case 0:
		goto L163
	case 1:
		goto L162
	default:
		v708 = v474
		v714 = v480
		goto L133
	}
L155:
	;
	v572 = v474
	v574 = int32(0)
	v575 = v561
	v578 = v480
	goto L156
L156:
	;
	if v575 != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v708 = v597
	v714 = v603
	goto L133
L158:
	;
	base.MemoryCopy(m, v578, v572, v575)
	goto L160
L159:
	;
	goto L160
L160:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v405+v475<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v574<<(uint(int32(2))%32)))) = v578
	v597 = v572 + v591
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v603 = v578 + (v598+int32(7))&int32(-8)
	v605 = v574 + int32(1)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v605 < v606 {
		v572 = v597
		v574 = v605
		v575 = v598
		v578 = v603
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L157
L162:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v655 <= int32(0) {
		v708 = v474
		v714 = v480
		goto L133
	} else {
		goto L171
	}
L163:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v610 <= int32(0) {
		v708 = v474
		v714 = v480
		goto L133
	} else {
		goto L164
	}
L164:
	;
	v618 = v474
	v621 = int32(0)
	v624 = v480
	goto L165
L165:
	;
	v637 = v618 + int32(4)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	if v638 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v708 = v645
	v714 = v650
	goto L133
L167:
	;
	base.MemoryCopy(m, v624, v637, v638)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v405+v475<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v640+v621<<(uint(int32(2))%32)))) = v624
	v645 = v638 + v637
	v650 = v624 + (v638+int32(7))&int32(-8)
	v652 = v621 + int32(1)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v652 < v653 {
		v618 = v645
		v621 = v652
		v624 = v650
		goto L165
	} else {
		goto L170
	}
L170:
	;
	goto L166
L171:
	;
	v663 = v474
	v666 = int32(0)
	v669 = v480
	goto L172
L172:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v663)))
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = v681<<(uint(int32(2))%32) + int32(16)
	v688 = v663 + int32(4)
	if v681 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v708 = v697
	v714 = v702
	goto L133
L174:
	;
	base.MemoryCopy(m, v669+int32(4), v688, v681)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v405+v475<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v692+v666<<(uint(int32(2))%32)))) = v669
	v697 = v688 + v681
	v702 = v669 + (v681+int32(11))&int32(-8)
	v704 = v666 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	if v704 < v705 {
		v663 = v697
		v666 = v704
		v669 = v702
		goto L172
	} else {
		goto L177
	}
L177:
	;
	goto L173
L178:
	;
	goto L132
L179:
	;
	v732 = v729
	goto L181
L180:
	;
	v732 = v113
	goto L181
L181:
	;
	v735 = int32(1)
	if v115 <= v735 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v738 = v735
	goto L184
L183:
	;
	v738 = v115
	goto L184
L184:
	;
	v747 = v708
	v750 = int32(0)
	v755 = v470
	v758 = v469
	goto L185
L185:
	;
	v767 = v467 + int32(48) + v750*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v767)+16)) = v755
	*(*int32)(unsafe.Add(mBase, uint32(v767)+20)) = v758
	if v115 != 0 {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v911 = int32(1)
	if v115 <= v911 {
		goto L208
	} else {
		goto L209
	}
L187:
	;
	base.MemoryCopy(m, v755, v747, v115)
	goto L189
L188:
	;
	goto L189
L189:
	;
	v771 = v747 + v115
	v772 = *(*int64)(unsafe.Add(mBase, uint32(v771)))
	*(*int64)(unsafe.Add(mBase, uint32(v767))) = v772
	v774 = *(*int64)(unsafe.Add(mBase, uint32(v771)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v767)+8)) = v774
	v777 = v771 + int32(16)
	v778 = int32(0)
	if base.B2i32(v115 < int32(2)) == v778 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v909 = v750 + int32(1)
	if v909 != v732 {
		v747 = v888
		v750 = v909
		v755 = v755 + v452
		v758 = v758 + v447
		goto L185
	} else {
		goto L207
	}
L191:
	;
	v783 = v777
	v784 = v778
	v789 = v778
	goto L194
L192:
	;
	v849 = v777
	v850 = v778
	goto L193
L193:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v767)+16))
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867+v850))))
	if v869 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L194:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v767)+16))
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801+v784))))
	if v803 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v738&int32(1) == int32(0) {
		v888 = v840
		goto L190
	} else {
		goto L203
	}
L196:
	;
	v806 = int32(2)
	v807 = v784 << (uint(v806) % 32)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v767)+20))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v405+v807)))
	v812 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v783))))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v811+v812<<(uint(v806)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v807+v808))) = v816
	goto L198
L197:
	;
	goto L198
L198:
	;
	v820 = v784 | int32(1)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v767)+16))
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820+v821))))
	if v823 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v826 = int32(2)
	v827 = v820 << (uint(v826) % 32)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v767)+20))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v405+v827)))
	v832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v783)+2)))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v831+v832<<(uint(v826)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v827+v828))) = v836
	goto L201
L200:
	;
	goto L201
L201:
	;
	v840 = v783 + int32(4)
	v841 = int32(2)
	v842 = v784 + v841
	v844 = v789 + v841
	if v844 != v738&int32(14) {
		v783 = v840
		v784 = v842
		v789 = v844
		goto L194
	} else {
		goto L202
	}
L202:
	;
	goto L195
L203:
	;
	v849 = v840
	v850 = v842
	goto L193
L204:
	;
	v872 = int32(2)
	v873 = v850 << (uint(v872) % 32)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v767)+20))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v873+v405)))
	v878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v849))))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v877+v878<<(uint(v872)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v873+v874))) = v882
	goto L206
L205:
	;
	goto L206
L206:
	;
	v888 = v849 + int32(2)
	goto L190
L207:
	;
	goto L186
L208:
	;
	v914 = v911
	goto L210
L209:
	;
	v914 = v115
	goto L210
L210:
	;
	v917 = int32(0)
	goto L211
L211:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v405+v917<<(uint(int32(2))%32))))
	F_pfree(m, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L22
	} else {
		goto L213
	}
L212:
	;
	F_pfree(m, v405)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L22
	} else {
		goto L215
	}
L213:
	;
	v942 = v917 + int32(1)
	if v942 != v914 {
		v917 = v942
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v960 = v467
	goto L9
L216:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = int32(-509193790)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v973
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_9), v22+int32(112))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L22
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1057), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L22
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v991
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_10), v22+int32(96))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L22
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1061), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L22
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_11), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L22
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1064), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L22
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	v1022 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v1022
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_12), v22+int32(16))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L22
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1068), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L22
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_13), int32(0))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L22
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1071), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L22
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v1051
	F_errmsg_internal(m, int32(_a_F_statext_mcv_deserialize_14), v22+int32(32))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L22
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_statext_mcv_deserialize_1), int32(1074), int32(_a_F_statext_mcv_deserialize_2))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L22
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_statext_ndistinct_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L19
	} else {
		goto L85
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L82
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L19
	} else {
		goto L79
	}
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v236 = int32(0)
	goto L6
L6:
	;
	m.G0 = v11 - int32(-64)
	return v236
L7:
	;
	v91 = int32(1)
	v94 = v13 & v91
	if v94 != 0 {
		goto L34
	} else {
		goto L35
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L9:
	;
	if base.Ui32(int32(11)) < base.Ui32(v40) {
		goto L7
	} else {
		goto L18
	}
L10:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v16-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v28 = int32(1)
	if v13&v28 != 0 {
		v40 = int32(base.Ui32(v13)>>(uint(v28)%32)) - v28
		goto L9
	} else {
		goto L17
	}
L13:
	;
	if v16 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v27 = int32(16)
	goto L16
L15:
	;
	v27 = int32(0)
	goto L16
L16:
	;
	v40 = v27
	goto L9
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L18:
	;
	goto L8
L19:
	;
	return int32(0)
L20:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v50 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_0), v11)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L19
	} else {
		goto L32
	}
L22:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v56 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v67 = int32(1)
	if v50&v67 != 0 {
		v79 = int32(base.Ui32(v50)>>(uint(v67)%32)) - v67
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v59 = int32(16)
	goto L27
L26:
	;
	v59 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v66 = int32(4)
	goto L30
L29:
	;
	v66 = v59
	goto L30
L30:
	;
	v79 = v66
	goto L21
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(264), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v95 = v91
	goto L36
L35:
	;
	v95 = int32(4)
	goto L36
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0+v95)))
	if v97 != int32(-1554858076) {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v101 = l0 + int32(1)
	if v94 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v104 = v101
	goto L40
L39:
	;
	v104 = l0 + int32(4)
	goto L40
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v105 != int32(1) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v108 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v13 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v139 = v108 << (uint(int32(4)) % 32)
	v141 = v139 | int32(12)
	if base.Ui32(v137) < base.Ui32(v141) {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v116 == int32(18) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v127 = int32(1)
	if v94 != 0 {
		v137 = int32(base.Ui32(v13)>>(uint(v127)%32)) - v127
		goto L43
	} else {
		goto L53
	}
L47:
	;
	v119 = int32(16)
	goto L49
L48:
	;
	v119 = int32(0)
	goto L49
L49:
	;
	if base.Ui32((v116-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v126 = int32(4)
	goto L52
L51:
	;
	v126 = v119
	goto L52
L52:
	;
	v137 = v126
	goto L43
L53:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v137 = int32(base.Ui32(v131)>>(uint(int32(2))%32)) - int32(4)
	goto L43
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L19
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v193 = F_palloc0(m, v139+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L19
	} else {
		goto L71
	}
L57:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v147 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v176
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_0), v9+int32(-48))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L19
	} else {
		goto L69
	}
L59:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v153 == int32(18) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v164 = int32(1)
	if v147&v164 != 0 {
		v176 = int32(base.Ui32(v147)>>(uint(v164)%32)) - v164
		goto L58
	} else {
		goto L68
	}
L62:
	;
	v156 = int32(16)
	goto L64
L63:
	;
	v156 = int32(0)
	goto L64
L64:
	;
	if base.Ui32((v153-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v163 = int32(4)
	goto L67
L66:
	;
	v163 = v156
	goto L67
L67:
	;
	v176 = v163
	goto L58
L68:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v176 = int32(base.Ui32(v170)>>(uint(int32(2))%32)) - int32(4)
	goto L58
L69:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(290), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = int64(7035076516)
	v201 = v104 + int32(12)
	v205 = int32(0)
	goto L72
L72:
	;
	v211 = v193 + int32(16) + v205<<(uint(int32(4))%32)
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	*(*int64)(unsafe.Add(mBase, uint32(v211))) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+8)) = v214
	v218 = F_palloc(m, v214<<(uint(int32(1))%32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L19
	} else {
		goto L74
	}
L73:
	;
	v236 = v193
	goto L6
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+12)) = v218
	v222 = v201 + int32(12)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v225 = v223 << (uint(int32(1)) % 32)
	if v225 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	base.MemoryCopy(m, v218, v222, v225)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v228 = int32(1)
	v232 = v205 + v228
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	if base.Ui32(v232) < base.Ui32(v233) {
		v201 = v222 + v227<<(uint(v228)%32)
		v205 = v232
		goto L72
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(-1554858076)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v97
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_3), v9+int32(-16))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(279), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L19
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v105
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_4), v9+int32(-32))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L19
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(282), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L19
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
	F_errmsg_internal(m, int32(_a_F_statext_ndistinct_deserialize_5), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L19
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_statext_ndistinct_deserialize_1), int32(284), int32(_a_F_statext_ndistinct_deserialize_2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L19
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
