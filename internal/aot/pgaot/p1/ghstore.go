package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
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
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v639 int32
	_ = v639
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v726 int32
	_ = v726
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v931 int32
	_ = v931
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v994 int32
	_ = v994
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1044 int64
	_ = v1044
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v27 == v2 {
		v44 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v44&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v31 == int32(0) {
		v44 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 != int32(7) {
		v44 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v37 != int32(17) {
		v44 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+32)))
	v44 = v40 ^ int32(1)
	goto L2
L7:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = F_get_fn_opclass_options(m, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v55 = int32(128)
	goto L9
L9:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v56)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)))
	if v59&int32(4) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int64(0)
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v55 = v52 << (uint(int32(3)) % 32)
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L10
	} else {
		goto L131
	}
L13:
	;
	v65 = v25 + int32(8)
	v68 = base.I32_wrap_i64(v23) & int32(_a_F_ghstore_consistent_0)
	switch v68 - int32(7) {
	case 0, 6:
		goto L20
	default:
		goto L12
	case 2:
		goto L19
	case 3:
		goto L17
	case 4:
		goto L18
	}
L14:
	;
	v1044 = int64(1)
	goto L15
L15:
	;
	m.G0 = v20 + int32(16)
	return v1044
L16:
	;
	v1044 = base.I64_extend_i32_u(v1014)
	goto L15
L17:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v809 = F_pg_detoast_datum(m, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L10
	} else {
		goto L110
	}
L18:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v603 = F_pg_detoast_datum(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L10
	} else {
		goto L89
	}
L19:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v431 = F_pg_detoast_datum_packed(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L65
	}
L20:
	;
	v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v72 = F_hstoreUpgrade(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v76 = v74 & int32(268435455)
	if v76 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v1014 = int32(1)
	goto L16
L23:
	;
	goto L24
L24:
	;
	v81 = v72 + int32(8)
	v84 = v81 + v76<<(uint(int32(3))%32)
	v90 = v2
	goto L25
L25:
	;
	v104 = v81 + v90<<(uint(int32(3))%32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v1014 = v411
	goto L16
L27:
	;
	if v120 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v120 = v105 & int32(1073741823)
	v121 = v84
	goto L27
L29:
	;
	goto L30
L30:
	;
	v110 = int32(1073741823)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104-int32(4))))
	v116 = v114 & v110
	v120 = v105&v110 - v116
	v121 = v116 + v84
	goto L27
L31:
	;
	v123 = int32(-1)
	if v120 != int32(1) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v245 = int32(0)
	goto L33
L33:
	;
	v246 = base.I32_rem_u_s(v245, v55)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(base.Ui32(v246)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v250)>>(uint(v246&int32(7))%32))&int32(1) == int32(0) {
		v1014 = int32(0)
		goto L16
	} else {
		goto L42
	}
L34:
	;
	v245 = v209 ^ int32(-1)
	goto L33
L35:
	;
	v131 = v121
	v132 = v123
	v133 = int32(0)
	goto L38
L36:
	;
	v179 = v121
	v180 = v123
	goto L37
L37:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v204 = *(*int32)(unsafe.Add(mBase, uint32((v180^v196)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v209 = v204 ^ int32(base.Ui32(v180)>>(uint(int32(8))%32))
	goto L34
L38:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	v150 = int32(255)
	v152 = int32(2)
	v156 = *(*int32)(unsafe.Add(mBase, uint32((v148^v132)&v150<<(uint(v152)%32))+uint32(_c_F_ghstore_consistent[0])))
	v157 = int32(8)
	v159 = v156 ^ int32(base.Ui32(v132)>>(uint(v157)%32))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32((v159^v160)&v150<<(uint(v152)%32))+uint32(_c_F_ghstore_consistent[0])))
	v171 = v168 ^ int32(base.Ui32(v159)>>(uint(v157)%32))
	v173 = v131 + v152
	v175 = v133 + v152
	if v175 != v120&int32(-2) {
		v131 = v173
		v132 = v171
		v133 = v175
		goto L38
	} else {
		goto L40
	}
L39:
	;
	if v120&int32(1) == int32(0) {
		v209 = v171
		goto L34
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v179 = v173
	v180 = v171
	goto L37
L42:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v259&int32(1073741824) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v264 = int32(1073741823)
	v265 = v105 & v264
	v270 = base.B2i32(int32(0) <= v259)
	if int32(0) <= v259 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v411 = int32(1)
	goto L45
L45:
	;
	if v411 == int32(0) {
		v1014 = v411
		goto L16
	} else {
		goto L63
	}
L46:
	;
	v271 = v259 - v265
	goto L48
L47:
	;
	v271 = v259 & v264
	goto L48
L48:
	;
	if v271 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v272 = int32(-1)
	if int32(0) <= v259 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v397 = int32(0)
	goto L51
L51:
	;
	v398 = base.I32_rem_u_s(v397, v55)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(base.Ui32(v398)>>(uint(int32(3))%32))))))
	v411 = int32(base.Ui32(v402)>>(uint(v398&int32(7))%32)) & int32(1)
	goto L45
L52:
	;
	v274 = v265
	goto L54
L53:
	;
	v274 = int32(0)
	goto L54
L54:
	;
	v275 = v84 + v274
	if v271 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v397 = v361 ^ int32(-1)
	goto L51
L56:
	;
	v283 = v275
	v284 = v272
	v285 = int32(0)
	goto L59
L57:
	;
	v331 = v275
	v332 = v272
	goto L58
L58:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v356 = *(*int32)(unsafe.Add(mBase, uint32((v332^v348)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v361 = v356 ^ int32(base.Ui32(v332)>>(uint(int32(8))%32))
	goto L55
L59:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	v302 = int32(255)
	v304 = int32(2)
	v308 = *(*int32)(unsafe.Add(mBase, uint32((v300^v284)&v302<<(uint(v304)%32))+uint32(_c_F_ghstore_consistent[0])))
	v309 = int32(8)
	v311 = v308 ^ int32(base.Ui32(v284)>>(uint(v309)%32))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32((v311^v312)&v302<<(uint(v304)%32))+uint32(_c_F_ghstore_consistent[0])))
	v323 = v320 ^ int32(base.Ui32(v311)>>(uint(v309)%32))
	v325 = v283 + v304
	v327 = v285 + v304
	if v327 != v271&int32(-2) {
		v283 = v325
		v284 = v323
		v285 = v327
		goto L59
	} else {
		goto L61
	}
L60:
	;
	if v271&int32(1) == int32(0) {
		v361 = v323
		goto L55
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v331 = v325
	v332 = v323
	goto L58
L63:
	;
	v428 = v90 + int32(1)
	if base.Ui32(v428) < base.Ui32(v76) {
		v90 = v428
		goto L25
	} else {
		goto L64
	}
L64:
	;
	goto L26
L65:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	v434 = int32(1)
	v435 = v433 & v434
	if v433 == v434 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v592 = base.I32_rem_u_s(v591, v55)
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(base.Ui32(v592)>>(uint(int32(3))%32))))))
	v1014 = int32(base.Ui32(v596)>>(uint(v592&int32(7))%32)) & int32(1)
	goto L16
L67:
	;
	if v435 != 0 {
		goto L78
	} else {
		goto L79
	}
L68:
	;
	if v462 != 0 {
		v465 = v462
		goto L67
	} else {
		goto L77
	}
L69:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+1)))
	if base.Ui32((v439-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v465 = int32(4)
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v451 = int32(1)
	if v435 != 0 {
		v462 = int32(base.Ui32(v433)>>(uint(v451)%32)) - v451
		goto L68
	} else {
		goto L76
	}
L72:
	;
	if v439 == int32(18) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v450 = int32(16)
	goto L75
L74:
	;
	v450 = int32(0)
	goto L75
L75:
	;
	v462 = v450
	goto L68
L76:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	v462 = int32(base.Ui32(v455)>>(uint(int32(2))%32)) - int32(4)
	goto L68
L77:
	;
	v591 = int32(0)
	goto L66
L78:
	;
	v468 = int32(1)
	goto L80
L79:
	;
	v468 = int32(4)
	goto L80
L80:
	;
	v469 = v431 + v468
	v470 = int32(-1)
	if v465 != int32(1) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v591 = v556 ^ int32(-1)
	goto L66
L82:
	;
	v478 = v469
	v479 = v470
	v480 = int32(0)
	goto L85
L83:
	;
	v526 = v469
	v527 = v470
	goto L84
L84:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	v551 = *(*int32)(unsafe.Add(mBase, uint32((v527^v543)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v556 = v551 ^ int32(base.Ui32(v527)>>(uint(int32(8))%32))
	goto L81
L85:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	v497 = int32(255)
	v499 = int32(2)
	v503 = *(*int32)(unsafe.Add(mBase, uint32((v495^v479)&v497<<(uint(v499)%32))+uint32(_c_F_ghstore_consistent[0])))
	v504 = int32(8)
	v506 = v503 ^ int32(base.Ui32(v479)>>(uint(v504)%32))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+1)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32((v506^v507)&v497<<(uint(v499)%32))+uint32(_c_F_ghstore_consistent[0])))
	v518 = v515 ^ int32(base.Ui32(v506)>>(uint(v504)%32))
	v520 = v478 + v499
	v522 = v480 + v499
	if v522 != v465&int32(-2) {
		v478 = v520
		v479 = v518
		v480 = v522
		goto L85
	} else {
		goto L87
	}
L86:
	;
	if v465&int32(1) == int32(0) {
		v556 = v518
		goto L81
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v526 = v520
	v527 = v518
	goto L84
L89:
	;
	F_deconstruct_array_builtin(m, v603, int32(25), v20+int32(12), v20+int32(8), v20+int32(4))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v615 <= int32(0) {
		v1014 = int32(1)
		goto L16
	} else {
		goto L91
	}
L91:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v625 = v2
	goto L92
L92:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625+v619))))
	if v639 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v1014 = v789
	goto L16
L94:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v618+v625<<(uint(int32(3))%32))))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)))
	v648 = int32(base.Ui32(v646) >> (uint(int32(2)) % 32))
	v650 = v648 - int32(4)
	if v650 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v789 = int32(1)
	goto L96
L96:
	;
	if v789 == int32(0) {
		v1014 = v789
		goto L16
	} else {
		goto L108
	}
L97:
	;
	v652 = v645 + int32(4)
	v653 = int32(-1)
	if v648 != int32(5) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v775 = int32(0)
	goto L99
L99:
	;
	v776 = base.I32_rem_u_s(v775, v55)
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(base.Ui32(v776)>>(uint(int32(3))%32))))))
	v789 = int32(base.Ui32(v780)>>(uint(v776&int32(7))%32)) & int32(1)
	goto L96
L100:
	;
	v775 = v739 ^ int32(-1)
	goto L99
L101:
	;
	v661 = v652
	v662 = v653
	v663 = int32(0)
	goto L104
L102:
	;
	v709 = v652
	v710 = v653
	goto L103
L103:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v709))))
	v734 = *(*int32)(unsafe.Add(mBase, uint32((v710^v726)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v739 = v734 ^ int32(base.Ui32(v710)>>(uint(int32(8))%32))
	goto L100
L104:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661))))
	v680 = int32(255)
	v682 = int32(2)
	v686 = *(*int32)(unsafe.Add(mBase, uint32((v678^v662)&v680<<(uint(v682)%32))+uint32(_c_F_ghstore_consistent[0])))
	v687 = int32(8)
	v689 = v686 ^ int32(base.Ui32(v662)>>(uint(v687)%32))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+1)))
	v698 = *(*int32)(unsafe.Add(mBase, uint32((v689^v690)&v680<<(uint(v682)%32))+uint32(_c_F_ghstore_consistent[0])))
	v701 = v698 ^ int32(base.Ui32(v689)>>(uint(v687)%32))
	v703 = v661 + v682
	v705 = v663 + v682
	if v705 != v650&int32(-2) {
		v661 = v703
		v662 = v701
		v663 = v705
		goto L104
	} else {
		goto L106
	}
L105:
	;
	if v648&int32(1) == int32(0) {
		v739 = v701
		goto L100
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	v709 = v703
	v710 = v701
	goto L103
L108:
	;
	v806 = v625 + int32(1)
	if v806 < v615 {
		v625 = v806
		goto L92
	} else {
		goto L109
	}
L109:
	;
	goto L93
L110:
	;
	F_deconstruct_array_builtin(m, v809, int32(25), v20+int32(12), v20+int32(8), v20+int32(4))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v820 <= int32(0) {
		v1014 = v2
		goto L16
	} else {
		goto L112
	}
L112:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v830 = v2
	goto L113
L113:
	;
	v842 = int32(0)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830+v824))))
	if v844 == v842 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v1014 = v994
	goto L16
L115:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v823+v830<<(uint(int32(3))%32))))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v850)))
	v853 = int32(base.Ui32(v851) >> (uint(int32(2)) % 32))
	v855 = v853 - int32(4)
	if v855 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v994 = v842
	goto L117
L117:
	;
	if v994 != 0 {
		v1014 = v994
		goto L16
	} else {
		goto L129
	}
L118:
	;
	v857 = v850 + int32(4)
	v858 = int32(-1)
	if v853 != int32(5) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v980 = int32(0)
	goto L120
L120:
	;
	v981 = base.I32_rem_u_s(v980, v55)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+int32(base.Ui32(v981)>>(uint(int32(3))%32))))))
	v994 = int32(base.Ui32(v985)>>(uint(v981&int32(7))%32)) & int32(1)
	goto L117
L121:
	;
	v980 = v944 ^ int32(-1)
	goto L120
L122:
	;
	v866 = v857
	v867 = v858
	v868 = int32(0)
	goto L125
L123:
	;
	v914 = v857
	v915 = v858
	goto L124
L124:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	v939 = *(*int32)(unsafe.Add(mBase, uint32((v915^v931)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_consistent[0])))
	v944 = v939 ^ int32(base.Ui32(v915)>>(uint(int32(8))%32))
	goto L121
L125:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	v885 = int32(255)
	v887 = int32(2)
	v891 = *(*int32)(unsafe.Add(mBase, uint32((v883^v867)&v885<<(uint(v887)%32))+uint32(_c_F_ghstore_consistent[0])))
	v892 = int32(8)
	v894 = v891 ^ int32(base.Ui32(v867)>>(uint(v892)%32))
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+1)))
	v903 = *(*int32)(unsafe.Add(mBase, uint32((v894^v895)&v885<<(uint(v887)%32))+uint32(_c_F_ghstore_consistent[0])))
	v906 = v903 ^ int32(base.Ui32(v894)>>(uint(v892)%32))
	v908 = v866 + v887
	v910 = v868 + v887
	if v910 != v855&int32(-2) {
		v866 = v908
		v867 = v906
		v868 = v910
		goto L125
	} else {
		goto L127
	}
L126:
	;
	if v853&int32(1) == int32(0) {
		v944 = v906
		goto L121
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v914 = v908
	v915 = v906
	goto L124
L129:
	;
	v1009 = v830 + int32(1)
	if v1009 < v820 {
		v830 = v1009
		goto L113
	} else {
		goto L130
	}
L130:
	;
	goto L114
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v68
	F_errmsg_internal(m, int32(_a_F_ghstore_consistent_1), v20)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ghstore_consistent_2), int32(609), int32(_a_F_ghstore_consistent_3))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
