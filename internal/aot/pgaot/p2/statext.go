package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_is_kind_built(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = l1 - int32(100)
	v13 = v11 & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v13) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(485842), v8)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(491914), int32(409), int32(98411))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if int32(base.Ui32(int32(519))>>(uint(v13)%32))&int32(1) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_errmsg_internal(m, int32(485842), v8)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491914), int32(409), int32(98411))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11&int32(255)<<(uint(int32(2))%32))+uint32(_consts[724])))
			v30 = F_heap_attisnull(m, l0, v28, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v30 ^ int32(1)
			}
		}
	}
}
func F_statext_mcv_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int64
	_ = v824
	var v826 int64
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v999 int32
	_ = v999
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(128)
	m.G0 = v23
	if l0 == v2 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L23
	} else {
		goto L262
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L23
	} else {
		goto L259
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L23
	} else {
		goto L256
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L23
	} else {
		goto L253
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L23
	} else {
		goto L250
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L23
	} else {
		goto L247
	}
L7:
	;
	m.G0 = v23 + int32(128)
	return v999
L8:
	;
	v999 = v2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v27 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v100 = F_palloc0(m, int32(48))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L23
	} else {
		goto L41
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	if base.Ui32(int32(17)) < base.Ui32(v50) {
		goto L11
	} else {
		goto L22
	}
L14:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v42 = int32(1)
	if v27&v42 != 0 {
		v50 = int32(base.Ui32(v27) >> (uint(v42) % 32))
		goto L13
	} else {
		goto L21
	}
L17:
	;
	v37 = int32(18)
	if v30 == v37 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v41 = v37
	goto L20
L19:
	;
	v41 = int32(2)
	goto L20
L20:
	;
	v50 = v41
	goto L13
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = int32(base.Ui32(v46) >> (uint(int32(2)) % 32))
	goto L13
L22:
	;
	goto L12
L23:
	;
	return int32(0)
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v60 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v87
	F_errmsg_internal(m, int32(654017), v23)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L39
	}
L26:
	;
	v63 = int32(6)
	v65 = int32(18)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v67 == v65 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v79 = int32(1)
	if v60&v79 != 0 {
		v87 = int32(base.Ui32(v60) >> (uint(v79) % 32))
		goto L25
	} else {
		goto L38
	}
L29:
	;
	v70 = v65
	goto L31
L30:
	;
	v70 = int32(2)
	goto L31
L31:
	;
	if v67&int32(254) == int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v75 = v63
	goto L34
L33:
	;
	v75 = v70
	goto L34
L34:
	;
	if v67 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v78 = v63
	goto L37
L36:
	;
	v78 = v75
	goto L37
L37:
	;
	v87 = v78
	goto L25
L38:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = int32(base.Ui32(v83) >> (uint(int32(2)) % 32))
	goto L25
L39:
	;
	F_errfinish(m, int32(490781), int32(1032), int32(340820))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v102 = int32(1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v106 = v104 & v102
	if v106 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v107 = v102
	goto L44
L43:
	;
	v107 = int32(4)
	goto L44
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0+v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v109
	v112 = l0 + int32(1)
	if v106 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v115 = v112
	goto L47
L46:
	;
	v115 = l0 + int32(4)
	goto L47
L47:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v118
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v100)+12)) = uint16(v120)
	if v109 != int32(-509193790) {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	if v116 != int32(1) {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	if v120 == int32(0) {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v120) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	if v118 == int32(0) {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(int32(10001)) <= base.Ui32(v118) {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v134 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v163 = v120 << (uint(int32(2)) % 32)
	v165 = v120 * int32(20)
	v174 = v163 + v165 + (v120*int32(3)+int32(16))*v118 + int32(18)
	if base.Ui32(v161) < base.Ui32(v174) {
		goto L68
	} else {
		goto L69
	}
L55:
	;
	v137 = int32(6)
	v139 = int32(18)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v141 == v139 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v153 = int32(1)
	if v134&v153 != 0 {
		v161 = int32(base.Ui32(v134) >> (uint(v153) % 32))
		goto L54
	} else {
		goto L67
	}
L58:
	;
	v144 = v139
	goto L60
L59:
	;
	v144 = int32(2)
	goto L60
L60:
	;
	if v141&int32(254) == int32(2) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v149 = v137
	goto L63
L62:
	;
	v149 = v144
	goto L63
L63:
	;
	if v141 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v152 = v137
	goto L66
L65:
	;
	v152 = v149
	goto L66
L66:
	;
	v161 = v152
	goto L54
L67:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v161 = int32(base.Ui32(v157) >> (uint(int32(2)) % 32))
	goto L54
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L23
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v223 = v115 + int32(14)
	if v163 != 0 {
		goto L89
	} else {
		goto L90
	}
L71:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v180 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v207
	F_errmsg_internal(m, int32(654062), v23+int32(48))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L23
	} else {
		goto L86
	}
L73:
	;
	v183 = int32(6)
	v185 = int32(18)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v187 == v185 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v199 = int32(1)
	if v180&v199 != 0 {
		v207 = int32(base.Ui32(v180) >> (uint(v199) % 32))
		goto L72
	} else {
		goto L85
	}
L76:
	;
	v190 = v185
	goto L78
L77:
	;
	v190 = int32(2)
	goto L78
L78:
	;
	if v187&int32(254) == int32(2) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v195 = v183
	goto L81
L80:
	;
	v195 = v190
	goto L81
L81:
	;
	if v187 == int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v198 = v183
	goto L84
L83:
	;
	v198 = v195
	goto L84
L84:
	;
	v207 = v198
	goto L72
L85:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = int32(base.Ui32(v203) >> (uint(int32(2)) % 32))
	goto L72
L86:
	;
	F_errfinish(m, int32(490781), int32(1093), int32(340820))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L23
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v226 = F_palloc(m, v165)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L23
	} else {
		goto L92
	}
L89:
	;
	v224 = F__emscripten_memcpy_bulkmem(m, v100+int32(16), v223, v163)
	mBase = m.M
	goto L91
L90:
	;
	goto L91
L91:
	;
	goto L88
L92:
	;
	v228 = v223 + v163
	if v165 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v231 = int32(1)
	if v120 <= v231 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v229 = F__emscripten_memcpy_bulkmem(m, v226, v228, v165)
	mBase = m.M
	v230 = v229
	goto L96
L95:
	;
	v230 = v226
	goto L96
L96:
	;
	goto L93
L97:
	;
	v234 = v231
	goto L99
L98:
	;
	v234 = v120
	goto L99
L99:
	;
	v236 = v234 & int32(3)
	if v120 < int32(4) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v236 != 0 {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v296 = int32(0)
	v297 = v174
	goto L100
L102:
	;
	goto L103
L103:
	;
	v245 = int32(0)
	v246 = v174
	v250 = v2
	goto L104
L104:
	;
	v266 = int32(20)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v230+(v245|int32(3))*v266)+4))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v230+(v245|int32(2))*v266)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v230+(v245|int32(1))*v266)+4))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v230+v245*v266)+4))
	v289 = v269 + (v275 + (v281 + (v285 + v246)))
	v290 = int32(4)
	v291 = v245 + v290
	v293 = v250 + v290
	if v293 != v234&int32(12) {
		v245 = v291
		v246 = v289
		v250 = v293
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v296 = v291
	v297 = v289
	goto L100
L106:
	;
	goto L105
L107:
	;
	v316 = v296
	v317 = v297
	v318 = int32(0)
	goto L110
L108:
	;
	v347 = v297
	goto L109
L109:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v365 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v230+v316*int32(20))+4))
	v339 = v338 + v317
	v340 = int32(1)
	v343 = v318 + v340
	if v343 != v236 {
		v316 = v316 + v340
		v317 = v339
		v318 = v343
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v347 = v339
	goto L109
L112:
	;
	goto L111
L113:
	;
	if v392 != v347 {
		goto L127
	} else {
		goto L128
	}
L114:
	;
	v368 = int32(6)
	v370 = int32(18)
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v372 == v370 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	v384 = int32(1)
	if v365&v384 != 0 {
		v392 = int32(base.Ui32(v365) >> (uint(v384) % 32))
		goto L113
	} else {
		goto L126
	}
L117:
	;
	v375 = v370
	goto L119
L118:
	;
	v375 = int32(2)
	goto L119
L119:
	;
	if v372&int32(254) == int32(2) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v380 = v368
	goto L122
L121:
	;
	v380 = v375
	goto L122
L122:
	;
	if v372 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v383 = v368
	goto L125
L124:
	;
	v383 = v380
	goto L125
L125:
	;
	v392 = v383
	goto L113
L126:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v392 = int32(base.Ui32(v388) >> (uint(int32(2)) % 32))
	goto L113
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L23
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v439 = int32(1)
	if v120 <= v439 {
		goto L147
	} else {
		goto L148
	}
L130:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v398 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v425
	F_errmsg_internal(m, int32(654062), v23+int32(80))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L23
	} else {
		goto L145
	}
L132:
	;
	v401 = int32(6)
	v403 = int32(18)
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v405 == v403 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	v417 = int32(1)
	if v398&v417 != 0 {
		v425 = int32(base.Ui32(v398) >> (uint(v417) % 32))
		goto L131
	} else {
		goto L144
	}
L135:
	;
	v408 = v403
	goto L137
L136:
	;
	v408 = int32(2)
	goto L137
L137:
	;
	if v405&int32(254) == int32(2) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v413 = v401
	goto L140
L139:
	;
	v413 = v408
	goto L140
L140:
	;
	if v405 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v416 = v401
	goto L143
L142:
	;
	v416 = v413
	goto L143
L143:
	;
	v425 = v416
	goto L131
L144:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v425 = int32(base.Ui32(v421) >> (uint(int32(2)) % 32))
	goto L131
L145:
	;
	F_errfinish(m, int32(490781), int32(1125), int32(340820))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L23
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
	v442 = v439
	goto L149
L148:
	;
	v442 = v120
	goto L149
L149:
	;
	v444 = F_palloc(m, v163)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L23
	} else {
		goto L150
	}
L150:
	;
	v449 = int32(0)
	v450 = int32(0)
	goto L151
L151:
	;
	v467 = int32(2)
	v472 = v230 + v449*int32(20)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v476 = F_palloc(m, v473<<(uint(v467)%32))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L23
	} else {
		goto L153
	}
L152:
	;
	v484 = int32(7)
	v487 = (v163 + v484) & int32(120)
	v488 = v487 * v118
	v491 = int32(24)
	v492 = (v120 + v484) & v491
	v493 = v492 * v118
	v499 = (v118*v491 + int32(55)) & int32(1048568)
	v507 = F_repalloc(m, v100, v488+(v493+v499)+(v480+v484)&int32(-8))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L23
	} else {
		goto L155
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444+v449<<(uint(v467)%32)))) = v476
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v472)+8))
	v480 = v479 + v450
	v482 = v449 + int32(1)
	if v482 != v442 {
		v449 = v482
		v450 = v480
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v509 = v507 + v499
	v510 = v509 + v488
	v512 = int32(1)
	if v120 <= v512 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v515 = v512
	goto L158
L157:
	;
	v515 = v120
	goto L158
L158:
	;
	v518 = v165 + v228
	v522 = v510 + v493
	v526 = int32(0)
	goto L159
L159:
	;
	v539 = v230 + v526*int32(20)
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+16)))
	if v540 == int32(1) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v783 = int32(1)
	if v118 <= v783 {
		goto L211
	} else {
		goto L212
	}
L161:
	;
	v781 = v526 + int32(1)
	if v781 != v515 {
		v518 = v761
		v522 = v765
		v526 = v781
		goto L159
	} else {
		goto L210
	}
L162:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v543 <= int32(0) {
		v761 = v518
		v765 = v522
		goto L161
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	if int32(0) < v608 {
		goto L181
	} else {
		goto L182
	}
L165:
	;
	v551 = v518
	v553 = int32(0)
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+124)) = int32(0)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	if v574 != 0 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v761 = v598
	v765 = v522
	goto L161
L168:
	;
	switch v574 - int32(1) {
	case 0:
		goto L173
	case 1:
		goto L176
	default:
		goto L174
	case 3:
		goto L175
	}
L169:
	;
	v575 = F__emscripten_memcpy_bulkmem(m, v23+int32(124), v551, v574)
	mBase = m.M
	goto L171
L170:
	;
	goto L171
L171:
	;
	goto L168
L172:
	;
	v598 = v551 + v574
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v444+v526<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v599+v553<<(uint(int32(2))%32)))) = v597
	v605 = v553 + int32(1)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v605 < v606 {
		v551 = v598
		v553 = v605
		goto L166
	} else {
		goto L180
	}
L173:
	;
	v596 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+124)))
	v597 = v596
	goto L172
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L23
	} else {
		goto L177
	}
L175:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v23)+124))
	v597 = v580
	goto L172
L176:
	;
	v579 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+124)))
	v597 = v579
	goto L172
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v574
	F_errmsg_internal(m, int32(481560), v23-int32(-64))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L23
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(325640), int32(70), int32(67612))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L23
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	goto L167
L181:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v611 <= int32(0) {
		v761 = v518
		v765 = v522
		goto L161
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	switch v608 + int32(2) {
	case 0:
		goto L193
	case 1:
		goto L192
	default:
		v761 = v518
		v765 = v522
		goto L161
	}
L184:
	;
	v619 = v518
	v621 = int32(0)
	v623 = v522
	v628 = v608
	goto L185
L185:
	;
	if v628 != 0 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v761 = v646
	v765 = v652
	goto L161
L187:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v444+v526<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v641+v621<<(uint(int32(2))%32)))) = v639
	v646 = v619 + v640
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v652 = v639 + (v647+int32(7))&int32(-8)
	v654 = v621 + int32(1)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v654 < v655 {
		v619 = v646
		v621 = v654
		v623 = v652
		v628 = v647
		goto L185
	} else {
		goto L191
	}
L188:
	;
	v638 = F__emscripten_memcpy_bulkmem(m, v623, v619, v628)
	mBase = m.M
	v639 = v638
	goto L190
L189:
	;
	v639 = v623
	goto L190
L190:
	;
	goto L187
L191:
	;
	goto L186
L192:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v706 <= int32(0) {
		v761 = v518
		v765 = v522
		goto L161
	} else {
		goto L202
	}
L193:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v659 <= int32(0) {
		v761 = v518
		v765 = v522
		goto L161
	} else {
		goto L194
	}
L194:
	;
	v667 = v518
	v669 = int32(0)
	v671 = v522
	goto L195
L195:
	;
	v687 = v667 + int32(4)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	if v688 != 0 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v761 = v696
	v765 = v701
	goto L161
L197:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v444+v526<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v691+v669<<(uint(int32(2))%32)))) = v690
	v696 = v688 + v687
	v701 = v690 + (v688+int32(7))&int32(-8)
	v703 = v669 + int32(1)
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v703 < v704 {
		v667 = v696
		v669 = v703
		v671 = v701
		goto L195
	} else {
		goto L201
	}
L198:
	;
	v689 = F__emscripten_memcpy_bulkmem(m, v671, v687, v688)
	mBase = m.M
	v690 = v689
	goto L200
L199:
	;
	v690 = v671
	goto L200
L200:
	;
	goto L197
L201:
	;
	goto L196
L202:
	;
	v714 = v518
	v718 = v522
	v719 = int32(0)
	goto L203
L203:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v733<<(uint(int32(2))%32) + int32(16)
	v739 = int32(4)
	v742 = v714 + v739
	if v733 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	v761 = v750
	v765 = v755
	goto L161
L205:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v444+v526<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v745+v719<<(uint(int32(2))%32)))) = v718
	v750 = v742 + v733
	v755 = v718 + (v733+int32(11))&int32(-8)
	v757 = v719 + int32(1)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	if v757 < v758 {
		v714 = v750
		v718 = v755
		v719 = v757
		goto L203
	} else {
		goto L209
	}
L206:
	;
	v743 = F__emscripten_memcpy_bulkmem(m, v718+v739, v742, v733)
	mBase = m.M
	goto L208
L207:
	;
	goto L208
L208:
	;
	goto L205
L209:
	;
	goto L204
L210:
	;
	goto L160
L211:
	;
	v786 = v783
	goto L213
L212:
	;
	v786 = v118
	goto L213
L213:
	;
	v797 = v761
	v803 = int32(0)
	v809 = v510
	v810 = v509
	goto L214
L214:
	;
	v818 = v507 + int32(48) + v803*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v818)+16)) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v818)+20)) = v810
	if v120 != 0 {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	v947 = int32(1)
	if v120 <= v947 {
		goto L239
	} else {
		goto L240
	}
L216:
	;
	v823 = v797 + v120
	v824 = *(*int64)(unsafe.Add(mBase, uint32(v823)))
	*(*int64)(unsafe.Add(mBase, uint32(v818))) = v824
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v823)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v818)+8)) = v826
	v829 = v823 + int32(16)
	v830 = int32(0)
	if base.B2i32(v120 < int32(2)) == v830 {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v821 = F__emscripten_memcpy_bulkmem(m, v809, v797, v120)
	mBase = m.M
	v822 = v821
	goto L219
L218:
	;
	v822 = v809
	goto L219
L219:
	;
	goto L216
L220:
	;
	v835 = v829
	v836 = v830
	v840 = v830
	goto L223
L221:
	;
	v900 = v829
	v901 = v830
	goto L222
L222:
	;
	if v234&int32(1) != 0 {
		goto L232
	} else {
		goto L233
	}
L223:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854+v836))))
	if v856 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v900 = v893
	v901 = v895
	goto L222
L225:
	;
	v859 = int32(2)
	v860 = v836 << (uint(v859) % 32)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v444+v860)))
	v865 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v835))))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v864+v865<<(uint(v859)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v860+v861))) = v869
	goto L227
L226:
	;
	goto L227
L227:
	;
	v873 = v836 | int32(1)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873+v874))))
	if v876 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v879 = int32(2)
	v880 = v873 << (uint(v879) % 32)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v444+v880)))
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v835)+2)))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v884+v885<<(uint(v879)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v880+v881))) = v889
	goto L230
L229:
	;
	goto L230
L230:
	;
	v893 = v835 + int32(4)
	v894 = int32(2)
	v895 = v836 + v894
	v897 = v840 + v894
	if v897 != v234&int32(14) {
		v835 = v893
		v836 = v895
		v840 = v897
		goto L223
	} else {
		goto L231
	}
L231:
	;
	goto L224
L232:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921+v901))))
	if v923 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	v939 = v900
	goto L234
L234:
	;
	v945 = v803 + int32(1)
	if v945 != v786 {
		v797 = v939
		v803 = v945
		v809 = v822 + v492
		v810 = v810 + v487
		goto L214
	} else {
		goto L238
	}
L235:
	;
	v926 = int32(2)
	v927 = v901 << (uint(v926) % 32)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v444+v927)))
	v932 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v900))))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v931+v932<<(uint(v926)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v927+v928))) = v936
	goto L237
L236:
	;
	goto L237
L237:
	;
	v939 = v900 + int32(2)
	goto L234
L238:
	;
	goto L215
L239:
	;
	v950 = v947
	goto L241
L240:
	;
	v950 = v120
	goto L241
L241:
	;
	v953 = int32(0)
	goto L242
L242:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v444+v953<<(uint(int32(2))%32))))
	F_pfree(m, v975)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L23
	} else {
		goto L244
	}
L243:
	;
	F_pfree(m, v444)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L23
	} else {
		goto L246
	}
L244:
	;
	v979 = v953 + int32(1)
	if v979 != v950 {
		v953 = v979
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v999 = v507
	goto L7
L247:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = int32(-509193790)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v1011
	F_errmsg_internal(m, int32(654866), v23+int32(112))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L23
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(490781), int32(1057), int32(340820))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L23
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v1029
	F_errmsg_internal(m, int32(654832), v23+int32(96))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L23
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(490781), int32(1061), int32(340820))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L23
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errmsg_internal(m, int32(76271), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L23
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(490781), int32(1064), int32(340820))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L23
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	v1060 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1060
	F_errmsg_internal(m, int32(76318), v23+int32(16))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L23
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(490781), int32(1068), int32(340820))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L23
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errmsg_internal(m, int32(76365), int32(0))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L23
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(490781), int32(1071), int32(340820))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L23
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v1089
	F_errmsg_internal(m, int32(76407), v23+int32(32))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L23
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(490781), int32(1074), int32(340820))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L23
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_statext_ndistinct_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 == v2 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L17
	} else {
		goto L84
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L81
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L17
	} else {
		goto L78
	}
L4:
	;
	m.G0 = v12 - int32(-64)
	return v246
L5:
	;
	v246 = v2
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v94 = int32(1)
	v97 = v16 & v94
	if v97 != 0 {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	if base.Ui32(int32(11)) < base.Ui32(v42) {
		goto L8
	} else {
		goto L16
	}
L11:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v30 = int32(1)
	if v16&v30 != 0 {
		v42 = int32(base.Ui32(v16)>>(uint(v30)%32)) - v30
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v42 = base.B2i32(v19 == int32(18)) << (uint(int32(4)) % 32)
	goto L10
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L16:
	;
	goto L9
L17:
	;
	return int32(0)
L18:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v52 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v82
	F_errmsg_internal(m, int32(653854), v12)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L17
	} else {
		goto L30
	}
L20:
	;
	v55 = int32(4)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v57&int32(254) == int32(2) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v70 = int32(1)
	if v52&v70 != 0 {
		v82 = int32(base.Ui32(v52)>>(uint(v70)%32)) - v70
		goto L19
	} else {
		goto L29
	}
L23:
	;
	v66 = v55
	goto L25
L24:
	;
	v66 = base.B2i32(v57 == int32(18)) << (uint(v55) % 32)
	goto L25
L25:
	;
	if v57 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v69 = v55
	goto L28
L27:
	;
	v69 = v66
	goto L28
L28:
	;
	v82 = v69
	goto L19
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v82 = int32(base.Ui32(v76)>>(uint(int32(2))%32)) - int32(4)
	goto L19
L30:
	;
	F_errfinish(m, int32(491633), int32(264), int32(340844))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v98 = v94
	goto L34
L33:
	;
	v98 = int32(4)
	goto L34
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0+v98)))
	if v100 != int32(-1554858076) {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v104 = l0 + int32(1)
	if v97 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v107 = v104
	goto L38
L37:
	;
	v107 = l0 + int32(4)
	goto L38
L38:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v108 != int32(1) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v111 == int32(0) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v115 = v111 << (uint(int32(4)) % 32)
	v117 = v115 | int32(12)
	if v16 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(v145) < base.Ui32(v117) {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v120 = int32(4)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v122&int32(254) == int32(2) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v135 = int32(1)
	if v97 != 0 {
		v145 = int32(base.Ui32(v16)>>(uint(v135)%32)) - v135
		goto L41
	} else {
		goto L51
	}
L45:
	;
	v131 = v120
	goto L47
L46:
	;
	v131 = base.B2i32(v122 == int32(18)) << (uint(v120) % 32)
	goto L47
L47:
	;
	if v122 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v134 = v120
	goto L50
L49:
	;
	v134 = v131
	goto L50
L50:
	;
	v145 = v134
	goto L41
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v145 = int32(base.Ui32(v139)>>(uint(int32(2))%32)) - int32(4)
	goto L41
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v198 = F_palloc0(m, v115+int32(16))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L17
	} else {
		goto L69
	}
L55:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v151 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v181
	F_errmsg_internal(m, int32(653854), v10+int32(-48))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L17
	} else {
		goto L67
	}
L57:
	;
	v154 = int32(4)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v156&int32(254) == int32(2) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v169 = int32(1)
	if v151&v169 != 0 {
		v181 = int32(base.Ui32(v151)>>(uint(v169)%32)) - v169
		goto L56
	} else {
		goto L66
	}
L60:
	;
	v165 = v154
	goto L62
L61:
	;
	v165 = base.B2i32(v156 == int32(18)) << (uint(v154) % 32)
	goto L62
L62:
	;
	if v156 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v168 = v154
	goto L65
L64:
	;
	v168 = v165
	goto L65
L65:
	;
	v181 = v168
	goto L56
L66:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v181 = int32(base.Ui32(v175)>>(uint(int32(2))%32)) - int32(4)
	goto L56
L67:
	;
	F_errfinish(m, int32(491633), int32(290), int32(340844))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L17
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
	*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v198))) = int64(7035076516)
	v207 = v107 + int32(12)
	v209 = int32(0)
	goto L70
L70:
	;
	v217 = v198 + int32(16) + v209<<(uint(int32(4))%32)
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v207)))
	*(*int64)(unsafe.Add(mBase, uint32(v217))) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = v220
	v224 = F_palloc(m, v220<<(uint(int32(1))%32))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L17
	} else {
		goto L72
	}
L71:
	;
	v246 = v198
	goto L4
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = v224
	v228 = v207 + int32(12)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v231 = v229 << (uint(int32(1)) % 32)
	if v231 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v235 = int32(1)
	v239 = v209 + v235
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	if base.Ui32(v239) < base.Ui32(v240) {
		v207 = v228 + v234<<(uint(v235)%32)
		v209 = v239
		goto L70
	} else {
		goto L77
	}
L74:
	;
	v232 = F__emscripten_memcpy_bulkmem(m, v224, v228, v231)
	mBase = m.M
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	goto L71
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(-1554858076)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v100
	F_errmsg_internal(m, int32(653632), v10+int32(-16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(491633), int32(279), int32(340844))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L17
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v108
	F_errmsg_internal(m, int32(664302), v10+int32(-32))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L17
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(491633), int32(282), int32(340844))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errmsg_internal(m, int32(108897), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(491633), int32(284), int32(340844))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L17
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
