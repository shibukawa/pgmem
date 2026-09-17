package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_step_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v571 int32
	_ = v571
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v732 int32
	_ = v732
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v885 int32
	_ = v885
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v946 int32
	_ = v946
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v982 int64
	_ = v982
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == v2 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L9
	} else {
		goto L290
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L9
	} else {
		goto L286
	}
L3:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_0), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L9
	} else {
		goto L284
	}
L4:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_1), int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L9
	} else {
		goto L282
	}
L5:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_2), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L9
	} else {
		goto L280
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	goto L69
L9:
	;
	return int32(0)
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	if base.Ui32(int32(_a_F_generate_series_step_numeric_3)) <= base.Ui32(v25) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
	if base.Ui32(int32(_a_F_generate_series_step_numeric_3)) <= base.Ui32(v46) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v25 == int32(_a_F_generate_series_step_numeric_3) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_4), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1731), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v69 == int32(3) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	if v46 == int32(_a_F_generate_series_step_numeric_3) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_7), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1742), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v73 = F_pg_detoast_datum(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	v144 = int32(1)
	v146 = int32(_a_F_generate_series_step_numeric_8)
	v148 = v2
	v149 = v2
	v150 = v2
	goto L30
L30:
	;
	v151 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L9
	} else {
		goto L56
	}
L31:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+4)))
	v76 = base.I32_extend16_s(v75)
	if base.Ui32(int32(_a_F_generate_series_step_numeric_3)) <= base.Ui32(v75) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v103 = base.B2i32(int32(0) <= v76)
	if int32(0) <= v76 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v76 == int32(-16384) {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_9), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1759), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v104 = int32(-8)
	goto L42
L41:
	;
	v104 = int32(-6)
	goto L42
L42:
	;
	v107 = int32(base.Ui32(int32(base.Ui32(v97)>>(uint(int32(2))%32))+v104) >> (uint(int32(1)) % 32))
	if int32(0) <= v76 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+6)))
	v118 = v108
	goto L45
L44:
	;
	v118 = v75<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v75&int32(63)
	goto L45
L45:
	;
	if v107 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v126 = v75 & int32(_a_F_generate_series_step_numeric_3)
	if v126 == int32(_a_F_generate_series_step_numeric_10) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v129 = v75 << (uint(int32(1)) % 32) & int32(_a_F_generate_series_step_numeric_11)
	goto L49
L48:
	;
	v129 = v126
	goto L49
L49:
	;
	v137 = base.B2i32(v76 < int32(0))
	if v76 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v138 = int32(base.Ui32(v75)>>(uint(int32(7))%32)) & int32(63)
	goto L52
L51:
	;
	v138 = v75 & int32(_a_F_generate_series_step_numeric_12)
	goto L52
L52:
	;
	if v76 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v141 = int32(6)
	goto L55
L54:
	;
	v141 = int32(8)
	goto L55
L55:
	;
	v144 = v107
	v146 = v73 + v141
	v148 = v118
	v149 = v129
	v150 = v138
	goto L30
L56:
	;
	v153 = int32(_a_F_generate_series_step_numeric_13)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_numeric[0]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v151)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_numeric[0])) = v156
	v159 = F_palloc(m, int32(72))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	base.MemoryFill(m, v159, int32(0), int32(72))
	F_set_var_from_num(m, v18, v159)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	F_set_var_from_num(m, v23, v159+int32(24))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v171 = v144 << (uint(int32(1)) % 32)
	v174 = F_palloc(m, v171+int32(2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v176 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v174))) = uint16(v176)
	v179 = v174 + int32(2)
	if v171 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	base.MemoryCopy(m, v179, v146, v171)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v159)+64))
	if v181 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+68)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v159)+64)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v159)+60)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v159)+56)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v159)+52)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v159)+48)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v151)+16)) = v159
	*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_numeric[0])) = v154
	goto L8
L67:
	;
	goto L66
L68:
	;
	v969 = F_make_result_opt_error(m, v206, int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L9
	} else {
		goto L278
	}
L69:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+16))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+56))
	if v207 != int32(_a_F_generate_series_step_numeric_11) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v207 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v206)+24))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v587 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L73:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206)+24))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v212 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	if v211 == int32(0) {
		goto L68
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	if v211 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v210 == int32(_a_F_generate_series_step_numeric_11) {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L68
L79:
	;
	if v219 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v206)+28))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v206)+44))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	if v219 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L68
L83:
	;
	if int32(0) < v582 {
		goto L1
	} else {
		goto L175
	}
L84:
	;
	if v210 == int32(_a_F_generate_series_step_numeric_11) {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v210 == int32(0) {
		goto L68
	} else {
		goto L131
	}
L87:
	;
	v232 = int32(0)
	if base.B2i32(v224 < v226)&base.B2i32(v232 < v212) == v232 {
		v263 = v226
		v267 = v232
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v582 = v405
	goto L83
L89:
	;
	v405 = v395
	goto L88
L90:
	;
	if base.B2i32(v211 <= int32(0))|base.B2i32(v224 <= v263) != 0 {
		v299 = v224
		v301 = v232
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v244 = v226
	v248 = v232
	goto L92
L92:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227+v248<<(uint(int32(1))%32)))))
	if v254 != 0 {
		v395 = int32(1)
		goto L89
	} else {
		goto L94
	}
L93:
	;
	v263 = v258
	v267 = v256
	goto L90
L94:
	;
	v255 = int32(1)
	v256 = v248 + v255
	v258 = v244 - v255
	if v258 <= v224 {
		v263 = v258
		v267 = v256
		goto L90
	} else {
		goto L95
	}
L95:
	;
	if v256 < v212 {
		v244 = v258
		v248 = v256
		goto L92
	} else {
		goto L96
	}
L96:
	;
	goto L93
L97:
	;
	if v263 != v299 {
		v341 = v267
		v342 = v301
		goto L104
	} else {
		goto L105
	}
L98:
	;
	v280 = v224
	v282 = v232
	goto L99
L99:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+v282<<(uint(int32(1))%32)))))
	if v287 != 0 {
		v395 = int32(-1)
		goto L89
	} else {
		goto L101
	}
L100:
	;
	v299 = v291
	v301 = v289
	goto L97
L101:
	;
	v288 = int32(1)
	v289 = v282 + v288
	v291 = v280 - v288
	if v291 <= v263 {
		v299 = v291
		v301 = v289
		goto L97
	} else {
		goto L102
	}
L102:
	;
	if v289 < v211 {
		v280 = v291
		v282 = v289
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	if v212 < v341 {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v310 = v267
	v311 = v301
	goto L106
L106:
	;
	if base.B2i32(v212 <= v310)|base.B2i32(v211 <= v311) != 0 {
		v341 = v310
		v342 = v311
		goto L104
	} else {
		goto L108
	}
L107:
	;
	if base.I32_extend16_s(v327) < base.I32_extend16_s(v325) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v316 = int32(1)
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227+v310<<(uint(v316)%32)))))
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311<<(uint(v316)%32)+v225))))
	if v325 == v327 {
		v310 = v310 + v316
		v311 = v311 + v316
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v334 = int32(1)
	goto L112
L111:
	;
	v334 = int32(-1)
	goto L112
L112:
	;
	v405 = v334
	goto L88
L113:
	;
	v345 = v341
	goto L115
L114:
	;
	v345 = v212
	goto L115
L115:
	;
	v352 = v341
	goto L116
L116:
	;
	if v345 == v352 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v395 = v378
	goto L89
L118:
	;
	if v211 < v342 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v378 = int32(1)
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227+v352<<(uint(v378)%32)))))
	if v384 == int32(0) {
		v352 = v352 + v378
		goto L116
	} else {
		goto L130
	}
L121:
	;
	v357 = v342
	goto L123
L122:
	;
	v357 = v211
	goto L123
L123:
	;
	v365 = v342
	goto L124
L124:
	;
	if v357 == v365 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v395 = int32(-1)
	goto L89
L126:
	;
	v405 = int32(0)
	goto L88
L127:
	;
	goto L128
L128:
	;
	v369 = int32(1)
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365<<(uint(v369)%32)+v225))))
	if v374 == int32(0) {
		v365 = v365 + v369
		goto L124
	} else {
		goto L129
	}
L129:
	;
	goto L125
L130:
	;
	goto L117
L131:
	;
	v408 = int32(0)
	if base.B2i32(v226 < v224)&base.B2i32(v408 < v211) == v408 {
		v439 = v224
		v443 = v408
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v582 = v581
	goto L83
L133:
	;
	v581 = v571
	goto L132
L134:
	;
	if base.B2i32(v212 <= int32(0))|base.B2i32(v226 <= v439) != 0 {
		v475 = v226
		v477 = v408
		goto L141
	} else {
		goto L142
	}
L135:
	;
	v420 = v224
	v424 = v408
	goto L136
L136:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+v424<<(uint(int32(1))%32)))))
	if v430 != 0 {
		v571 = int32(1)
		goto L133
	} else {
		goto L138
	}
L137:
	;
	v439 = v434
	v443 = v432
	goto L134
L138:
	;
	v431 = int32(1)
	v432 = v424 + v431
	v434 = v420 - v431
	if v434 <= v226 {
		v439 = v434
		v443 = v432
		goto L134
	} else {
		goto L139
	}
L139:
	;
	if v432 < v211 {
		v420 = v434
		v424 = v432
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	if v439 != v475 {
		v517 = v443
		v518 = v477
		goto L148
	} else {
		goto L149
	}
L142:
	;
	v456 = v226
	v458 = v408
	goto L143
L143:
	;
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227+v458<<(uint(int32(1))%32)))))
	if v463 != 0 {
		v571 = int32(-1)
		goto L133
	} else {
		goto L145
	}
L144:
	;
	v475 = v467
	v477 = v465
	goto L141
L145:
	;
	v464 = int32(1)
	v465 = v458 + v464
	v467 = v456 - v464
	if v467 <= v439 {
		v475 = v467
		v477 = v465
		goto L141
	} else {
		goto L146
	}
L146:
	;
	if v465 < v212 {
		v456 = v467
		v458 = v465
		goto L143
	} else {
		goto L147
	}
L147:
	;
	goto L144
L148:
	;
	if v211 < v517 {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	v486 = v443
	v487 = v477
	goto L150
L150:
	;
	if base.B2i32(v211 <= v486)|base.B2i32(v212 <= v487) != 0 {
		v517 = v486
		v518 = v487
		goto L148
	} else {
		goto L152
	}
L151:
	;
	if base.I32_extend16_s(v503) < base.I32_extend16_s(v501) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v492 = int32(1)
	v501 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+v486<<(uint(v492)%32)))))
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v487<<(uint(v492)%32)+v227))))
	if v501 == v503 {
		v486 = v486 + v492
		v487 = v487 + v492
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v510 = int32(1)
	goto L156
L155:
	;
	v510 = int32(-1)
	goto L156
L156:
	;
	v581 = v510
	goto L132
L157:
	;
	v521 = v517
	goto L159
L158:
	;
	v521 = v211
	goto L159
L159:
	;
	v528 = v517
	goto L160
L160:
	;
	if v521 == v528 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v571 = v554
	goto L133
L162:
	;
	if v212 < v518 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	v554 = int32(1)
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225+v528<<(uint(v554)%32)))))
	if v560 == int32(0) {
		v528 = v528 + v554
		goto L160
	} else {
		goto L174
	}
L165:
	;
	v533 = v518
	goto L167
L166:
	;
	v533 = v212
	goto L167
L167:
	;
	v541 = v518
	goto L168
L168:
	;
	if v533 == v541 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v571 = int32(-1)
	goto L133
L170:
	;
	v581 = int32(0)
	goto L132
L171:
	;
	goto L172
L172:
	;
	v545 = int32(1)
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v541<<(uint(v545)%32)+v227))))
	if v550 == int32(0) {
		v541 = v541 + v545
		goto L168
	} else {
		goto L173
	}
L173:
	;
	goto L169
L174:
	;
	goto L161
L175:
	;
	goto L68
L176:
	;
	if v586 == int32(0) {
		goto L68
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	if v586 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	if v585 != int32(_a_F_generate_series_step_numeric_11) {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	goto L68
L181:
	;
	if v594 == int32(0) {
		goto L68
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v206)+28))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v206)+44))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	if v594 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	goto L1
L185:
	;
	if v957 < int32(0) {
		goto L1
	} else {
		goto L277
	}
L186:
	;
	if v585 == int32(_a_F_generate_series_step_numeric_11) {
		goto L68
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if v585 == int32(0) {
		goto L1
	} else {
		goto L233
	}
L189:
	;
	v607 = int32(0)
	if base.B2i32(v599 < v601)&base.B2i32(v607 < v587) == v607 {
		v638 = v601
		v642 = v607
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v957 = v780
	goto L185
L191:
	;
	v780 = v770
	goto L190
L192:
	;
	if base.B2i32(v586 <= int32(0))|base.B2i32(v599 <= v638) != 0 {
		v674 = v599
		v676 = v607
		goto L199
	} else {
		goto L200
	}
L193:
	;
	v619 = v601
	v623 = v607
	goto L194
L194:
	;
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602+v623<<(uint(int32(1))%32)))))
	if v629 != 0 {
		v770 = int32(1)
		goto L191
	} else {
		goto L196
	}
L195:
	;
	v638 = v633
	v642 = v631
	goto L192
L196:
	;
	v630 = int32(1)
	v631 = v623 + v630
	v633 = v619 - v630
	if v633 <= v599 {
		v638 = v633
		v642 = v631
		goto L192
	} else {
		goto L197
	}
L197:
	;
	if v631 < v587 {
		v619 = v633
		v623 = v631
		goto L194
	} else {
		goto L198
	}
L198:
	;
	goto L195
L199:
	;
	if v638 != v674 {
		v716 = v642
		v717 = v676
		goto L206
	} else {
		goto L207
	}
L200:
	;
	v655 = v599
	v657 = v607
	goto L201
L201:
	;
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600+v657<<(uint(int32(1))%32)))))
	if v662 != 0 {
		v770 = int32(-1)
		goto L191
	} else {
		goto L203
	}
L202:
	;
	v674 = v666
	v676 = v664
	goto L199
L203:
	;
	v663 = int32(1)
	v664 = v657 + v663
	v666 = v655 - v663
	if v666 <= v638 {
		v674 = v666
		v676 = v664
		goto L199
	} else {
		goto L204
	}
L204:
	;
	if v664 < v586 {
		v655 = v666
		v657 = v664
		goto L201
	} else {
		goto L205
	}
L205:
	;
	goto L202
L206:
	;
	if v587 < v716 {
		goto L215
	} else {
		goto L216
	}
L207:
	;
	v685 = v642
	v686 = v676
	goto L208
L208:
	;
	if base.B2i32(v587 <= v685)|base.B2i32(v586 <= v686) != 0 {
		v716 = v685
		v717 = v686
		goto L206
	} else {
		goto L210
	}
L209:
	;
	if base.I32_extend16_s(v702) < base.I32_extend16_s(v700) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	v691 = int32(1)
	v700 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602+v685<<(uint(v691)%32)))))
	v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v686<<(uint(v691)%32)+v600))))
	if v700 == v702 {
		v685 = v685 + v691
		v686 = v686 + v691
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v709 = int32(1)
	goto L214
L213:
	;
	v709 = int32(-1)
	goto L214
L214:
	;
	v780 = v709
	goto L190
L215:
	;
	v720 = v716
	goto L217
L216:
	;
	v720 = v587
	goto L217
L217:
	;
	v727 = v716
	goto L218
L218:
	;
	if v720 == v727 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v770 = v753
	goto L191
L220:
	;
	if v586 < v717 {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	v753 = int32(1)
	v759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602+v727<<(uint(v753)%32)))))
	if v759 == int32(0) {
		v727 = v727 + v753
		goto L218
	} else {
		goto L232
	}
L223:
	;
	v732 = v717
	goto L225
L224:
	;
	v732 = v586
	goto L225
L225:
	;
	v740 = v717
	goto L226
L226:
	;
	if v732 == v740 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v770 = int32(-1)
	goto L191
L228:
	;
	v780 = int32(0)
	goto L190
L229:
	;
	goto L230
L230:
	;
	v744 = int32(1)
	v749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v740<<(uint(v744)%32)+v600))))
	if v749 == int32(0) {
		v740 = v740 + v744
		goto L226
	} else {
		goto L231
	}
L231:
	;
	goto L227
L232:
	;
	goto L219
L233:
	;
	v783 = int32(0)
	if base.B2i32(v601 < v599)&base.B2i32(v783 < v586) == v783 {
		v814 = v599
		v818 = v783
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v957 = v956
	goto L185
L235:
	;
	v956 = v946
	goto L234
L236:
	;
	if base.B2i32(v587 <= int32(0))|base.B2i32(v601 <= v814) != 0 {
		v850 = v601
		v852 = v783
		goto L243
	} else {
		goto L244
	}
L237:
	;
	v795 = v599
	v799 = v783
	goto L238
L238:
	;
	v805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600+v799<<(uint(int32(1))%32)))))
	if v805 != 0 {
		v946 = int32(1)
		goto L235
	} else {
		goto L240
	}
L239:
	;
	v814 = v809
	v818 = v807
	goto L236
L240:
	;
	v806 = int32(1)
	v807 = v799 + v806
	v809 = v795 - v806
	if v809 <= v601 {
		v814 = v809
		v818 = v807
		goto L236
	} else {
		goto L241
	}
L241:
	;
	if v807 < v586 {
		v795 = v809
		v799 = v807
		goto L238
	} else {
		goto L242
	}
L242:
	;
	goto L239
L243:
	;
	if v814 != v850 {
		v892 = v818
		v893 = v852
		goto L250
	} else {
		goto L251
	}
L244:
	;
	v831 = v601
	v833 = v783
	goto L245
L245:
	;
	v838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602+v833<<(uint(int32(1))%32)))))
	if v838 != 0 {
		v946 = int32(-1)
		goto L235
	} else {
		goto L247
	}
L246:
	;
	v850 = v842
	v852 = v840
	goto L243
L247:
	;
	v839 = int32(1)
	v840 = v833 + v839
	v842 = v831 - v839
	if v842 <= v814 {
		v850 = v842
		v852 = v840
		goto L243
	} else {
		goto L248
	}
L248:
	;
	if v840 < v587 {
		v831 = v842
		v833 = v840
		goto L245
	} else {
		goto L249
	}
L249:
	;
	goto L246
L250:
	;
	if v586 < v892 {
		goto L259
	} else {
		goto L260
	}
L251:
	;
	v861 = v818
	v862 = v852
	goto L252
L252:
	;
	if base.B2i32(v586 <= v861)|base.B2i32(v587 <= v862) != 0 {
		v892 = v861
		v893 = v862
		goto L250
	} else {
		goto L254
	}
L253:
	;
	if base.I32_extend16_s(v878) < base.I32_extend16_s(v876) {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v867 = int32(1)
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600+v861<<(uint(v867)%32)))))
	v878 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v862<<(uint(v867)%32)+v602))))
	if v876 == v878 {
		v861 = v861 + v867
		v862 = v862 + v867
		goto L252
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v885 = int32(1)
	goto L258
L257:
	;
	v885 = int32(-1)
	goto L258
L258:
	;
	v956 = v885
	goto L234
L259:
	;
	v896 = v892
	goto L261
L260:
	;
	v896 = v586
	goto L261
L261:
	;
	v903 = v892
	goto L262
L262:
	;
	if v896 == v903 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v946 = v929
	goto L235
L264:
	;
	if v587 < v893 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	v929 = int32(1)
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600+v903<<(uint(v929)%32)))))
	if v935 == int32(0) {
		v903 = v903 + v929
		goto L262
	} else {
		goto L276
	}
L267:
	;
	v908 = v893
	goto L269
L268:
	;
	v908 = v587
	goto L269
L269:
	;
	v916 = v893
	goto L270
L270:
	;
	if v908 == v916 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v946 = int32(-1)
	goto L235
L272:
	;
	v956 = int32(0)
	goto L234
L273:
	;
	goto L274
L274:
	;
	v920 = int32(1)
	v925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916<<(uint(v920)%32)+v602))))
	if v925 == int32(0) {
		v916 = v916 + v920
		goto L270
	} else {
		goto L275
	}
L275:
	;
	goto L271
L276:
	;
	goto L263
L277:
	;
	goto L68
L278:
	;
	v971 = int32(_a_F_generate_series_step_numeric_13)
	v972 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_numeric[0]))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v205)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_numeric[0])) = v974
	F_add_var(m, v206, v206+int32(48), v206)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L9
	} else {
		goto L279
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_numeric[0])) = v972
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v205)))
	*(*int64)(unsafe.Add(mBase, uint32(v205))) = v982 + int64(1)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v986)+20)) = int32(1)
	return v969
L280:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1727), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L9
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1738), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L9
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1755), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L9
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L9
	} else {
		goto L287
	}
L287:
	;
	F_errmsg(m, int32(_a_F_generate_series_step_numeric_14), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L9
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_generate_series_step_numeric_5), int32(1767), int32(_a_F_generate_series_step_numeric_6))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L9
	} else {
		goto L289
	}
L289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L290:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+20)) = int32(2)
	v1046 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1046)
	return int32(0)
}
func F_generate_series_timestamp_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v94 int32
	_ = v94
	var v99 int64
	_ = v99
	var v101 float64
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int64
	_ = v113
	var v126 float64
	_ = v126
	var v133 int32
	_ = v133
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 != int32(460) {
		v133 = v2
		return v133
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		if v15 == int32(0) {
			v133 = v2
			return v133
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v18 != int32(15) {
				v133 = v2
				return v133
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v25 = F_estimate_expression_value(m, v21, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					v32 = F_estimate_expression_value(m, v29, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
						v37 = F_estimate_expression_value(m, v34, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if v39 != int32(7) {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
								if v46 != int32(7) {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
									if v53 != int32(7) {
										v133 = v2
										return v133
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
										if v57 != 0 {
											v126 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
											v133 = v11
											return v133
										} else {
											v58 = int32(7)
											if base.B2i32(v39 != v58)|base.B2i32(v46 != v58) != 0 {
												v133 = v2
												return v133
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
												v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
												if base.Ui64(v64-int64(9223372036854775807)) < base.Ui64(int64(2)) {
													v133 = v2
													return v133
												} else {
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
													v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
													if base.B2i32(base.Ui64(v70-int64(9223372036854775807)) < base.Ui64(int64(2)))|(base.B2i32(v70-v64 < v70)^base.B2i32(int64(0) < v64)) != 0 {
														v133 = v2
														return v133
													} else {
														v81 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
														v84 = F_Int64GetDatum(m, v70)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															v86 = F_Int64GetDatum(m, v64)
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return int32(0)
															} else {
																v88 = F_DirectFunctionCall2Coll(m, int32(1501), int32(0), v84, v86)
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int32(0)
																} else {
																	v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
																	v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
																	v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
																	v101 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v90), float64(30)), base.F64_convert_i32_s(v94)), float64(8.64e+10)), base.F64_convert_i64_s(v99))
																	if base.F64_eq(v101, float64(0)) != 0 {
																		v133 = v2
																	} else {
																		v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
																		v108 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
																		v113 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
																		v126 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v104), float64(30)), base.F64_convert_i32_s(v108)), float64(8.64e+10)), base.F64_convert_i64_s(v113)), v101), float64(1)))
																		*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
																		v133 = v11
																	}
																	return v133
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
									if v49 == int32(0) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
										if v53 != int32(7) {
											v133 = v2
											return v133
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
											if v57 != 0 {
												v126 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
												v133 = v11
												return v133
											} else {
												v58 = int32(7)
												if base.B2i32(v39 != v58)|base.B2i32(v46 != v58) != 0 {
													v133 = v2
													return v133
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
													if base.Ui64(v64-int64(9223372036854775807)) < base.Ui64(int64(2)) {
														v133 = v2
														return v133
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
														v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
														if base.B2i32(base.Ui64(v70-int64(9223372036854775807)) < base.Ui64(int64(2)))|(base.B2i32(v70-v64 < v70)^base.B2i32(int64(0) < v64)) != 0 {
															v133 = v2
															return v133
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
															v84 = F_Int64GetDatum(m, v70)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																v86 = F_Int64GetDatum(m, v64)
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return int32(0)
																} else {
																	v88 = F_DirectFunctionCall2Coll(m, int32(1501), int32(0), v84, v86)
																	mBase = m.M
																	v89 = m.ExcPending
																	if v89 != 0 {
																		return int32(0)
																	} else {
																		v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
																		v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
																		v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
																		v101 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v90), float64(30)), base.F64_convert_i32_s(v94)), float64(8.64e+10)), base.F64_convert_i64_s(v99))
																		if base.F64_eq(v101, float64(0)) != 0 {
																			v133 = v2
																		} else {
																			v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
																			v113 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
																			v126 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v104), float64(30)), base.F64_convert_i32_s(v108)), float64(8.64e+10)), base.F64_convert_i64_s(v113)), v101), float64(1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
																			v133 = v11
																		}
																		return v133
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v126 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
										v133 = v11
										return v133
									}
								}
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)))
								if v42 == int32(0) {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									if v46 != int32(7) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
										if v53 != int32(7) {
											v133 = v2
											return v133
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
											if v57 != 0 {
												v126 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
												v133 = v11
												return v133
											} else {
												v58 = int32(7)
												if base.B2i32(v39 != v58)|base.B2i32(v46 != v58) != 0 {
													v133 = v2
													return v133
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
													if base.Ui64(v64-int64(9223372036854775807)) < base.Ui64(int64(2)) {
														v133 = v2
														return v133
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
														v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
														if base.B2i32(base.Ui64(v70-int64(9223372036854775807)) < base.Ui64(int64(2)))|(base.B2i32(v70-v64 < v70)^base.B2i32(int64(0) < v64)) != 0 {
															v133 = v2
															return v133
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
															v84 = F_Int64GetDatum(m, v70)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																v86 = F_Int64GetDatum(m, v64)
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return int32(0)
																} else {
																	v88 = F_DirectFunctionCall2Coll(m, int32(1501), int32(0), v84, v86)
																	mBase = m.M
																	v89 = m.ExcPending
																	if v89 != 0 {
																		return int32(0)
																	} else {
																		v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
																		v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
																		v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
																		v101 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v90), float64(30)), base.F64_convert_i32_s(v94)), float64(8.64e+10)), base.F64_convert_i64_s(v99))
																		if base.F64_eq(v101, float64(0)) != 0 {
																			v133 = v2
																		} else {
																			v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
																			v113 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
																			v126 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v104), float64(30)), base.F64_convert_i32_s(v108)), float64(8.64e+10)), base.F64_convert_i64_s(v113)), v101), float64(1)))
																			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
																			v133 = v11
																		}
																		return v133
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
										if v49 == int32(0) {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
											if v53 != int32(7) {
												v133 = v2
												return v133
											} else {
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
												if v57 != 0 {
													v126 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
													v133 = v11
													return v133
												} else {
													v58 = int32(7)
													if base.B2i32(v39 != v58)|base.B2i32(v46 != v58) != 0 {
														v133 = v2
														return v133
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														v64 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
														if base.Ui64(v64-int64(9223372036854775807)) < base.Ui64(int64(2)) {
															v133 = v2
															return v133
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
															v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
															if base.B2i32(base.Ui64(v70-int64(9223372036854775807)) < base.Ui64(int64(2)))|(base.B2i32(v70-v64 < v70)^base.B2i32(int64(0) < v64)) != 0 {
																v133 = v2
																return v133
															} else {
																v81 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
																v84 = F_Int64GetDatum(m, v70)
																mBase = m.M
																v85 = m.ExcPending
																if v85 != 0 {
																	return int32(0)
																} else {
																	v86 = F_Int64GetDatum(m, v64)
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
																		return int32(0)
																	} else {
																		v88 = F_DirectFunctionCall2Coll(m, int32(1501), int32(0), v84, v86)
																		mBase = m.M
																		v89 = m.ExcPending
																		if v89 != 0 {
																			return int32(0)
																		} else {
																			v90 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
																			v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
																			v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
																			v101 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v90), float64(30)), base.F64_convert_i32_s(v94)), float64(8.64e+10)), base.F64_convert_i64_s(v99))
																			if base.F64_eq(v101, float64(0)) != 0 {
																				v133 = v2
																			} else {
																				v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
																				v108 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
																				v113 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
																				v126 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v104), float64(30)), base.F64_convert_i32_s(v108)), float64(8.64e+10)), base.F64_convert_i64_s(v113)), v101), float64(1)))
																				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
																				v133 = v11
																			}
																			return v133
																		}
																	}
																}
															}
														}
													}
												}
											}
										} else {
											v126 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
											v133 = v11
											return v133
										}
									}
								} else {
									v126 = float64(0)
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v126
									v133 = v11
									return v133
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_generate_series_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v97 int64
	_ = v97
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v134 int32
	_ = v134
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(272)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == v2 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return v206
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L48
	}
L3:
	;
	m.G0 = v13 + int32(272)
	goto L1
L4:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L12
	} else {
		goto L47
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L12
	} else {
		goto L43
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v24 == int32(4) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	goto L34
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v28 = F_pg_detoast_datum_packed(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v32 = v2
	goto L11
L11:
	;
	v33 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	v32 = v28
	goto L11
L14:
	;
	v35 = int32(_a_F_generate_series_timestamptz_0)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamptz[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamptz[0])) = v38
	v41 = F_palloc(m, int32(40))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v23
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+16)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = v47
	if v32 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v67 = base.I64_extend_i32_s(v61) + base.I64_extend_i32_s(v63)*int64(30)
	v76 = int64(32)
	v77 = int64(20)
	v79 = int64(base.Ui64(v67) >> (uint(v76) % 64))
	v82 = int64(4294967295)
	v83 = int64(500654080)
	v85 = v67 & v82
	v86 = v83 * v85
	v90 = int64(base.Ui64(v86)>>(uint(v76)%64)) + v83*v79
	v97 = v85*v77 + v90&v82
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v67*int64(0) + v67>>(uint(int64(63))%64)*int64(86400000000) + v77*v79 + int64(base.Ui64(v90)>>(uint(v76)%64)) + int64(base.Ui64(v97)>>(uint(v76)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v86&v82 | v97<<(uint(v76)%64)
	goto L22
L17:
	;
	v50 = v13 + int32(16)
	F_text_to_cstring_buffer(m, v32, v50, int32(256))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamptz[1]))
	v59 = v57
	goto L16
L20:
	;
	v54 = F_DecodeTimezoneNameToTz(m, v50)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v59 = v54
	goto L16
L22:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v41)+16))
	v110 = v108 + v109
	v111 = int64(0)
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v119 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v110) < base.Ui64(v108))) + (v115 + v109>>(uint(int64(63))%64))
	if v119 == v111 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v124 = base.B2i32(v110 != v111)
	goto L25
L24:
	;
	v124 = base.B2i32(v111 < v119)
	goto L25
L25:
	;
	v125 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v124 - base.B2i32(v119 < v125)
	if v119|v110 == v125 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v63 != int32(2147483647) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v41
	*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamptz[0])) = v36
	goto L8
L28:
	;
	v134 = int32(-2147483648)
	if base.B2i32(v63 != v134)|base.B2i32(v109 != int64(-9223372036854775807-1))|base.B2i32(v61 != v134) != 0 {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v109 != int64(9223372036854775807) {
		goto L27
	} else {
		goto L32
	}
L31:
	;
	goto L2
L32:
	;
	if v61 == int32(2147483647) {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	goto L27
L34:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v159)+8))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v159)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+32))
	if int32(0) < v162 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v159)+36))
	v170 = F_timestamptz_pl_interval_internal(m, v161, v159+int32(16), v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L41
	}
L36:
	;
	if v161 <= v160 {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v161 < v160 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	goto L4
L40:
	;
	goto L35
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = v170
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = v173 + int64(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+20)) = int32(1)
	v180 = F_Int64GetDatum(m, v161)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v206 = v180
	goto L3
L43:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_generate_series_timestamptz_1), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_generate_series_timestamptz_2), int32(_a_F_generate_series_timestamptz_3), int32(_a_F_generate_series_timestamptz_4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L12
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
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = int32(2)
	v203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v203)
	v206 = int32(0)
	goto L3
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_generate_series_timestamptz_5), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_generate_series_timestamptz_2), int32(_a_F_generate_series_timestamptz_6), int32(_a_F_generate_series_timestamptz_4))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
