package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DefineQueryRewrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v615 int64
	_ = v615
	var v622 int64
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	v16 = m.G0
	v18 = v16 - int32(272)
	m.G0 = v18
	v21 = F_table_open(m, l2, int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L10
	} else {
		goto L233
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L10
	} else {
		goto L229
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L10
	} else {
		goto L225
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L10
	} else {
		goto L221
	}
L5:
	;
	if l5 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L6:
	;
	v546 = int32(_a_F_DefineQueryRewrite_0)
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[0])))
	if base.B2i32(v549 == int32(0))|base.B2i32(v549 != v552) != 0 {
		v570 = v549
		v571 = v552
		goto L150
	} else {
		goto L151
	}
L7:
	;
	v493 = int32(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v494 <= v493 {
		goto L6
	} else {
		goto L138
	}
L8:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+119)))
	switch v238 - int32(109) {
	case 0, 9:
		goto L71
	default:
		goto L72
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L67
	}
L10:
	;
	return
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+119)))
	v26 = v24 - int32(109)
	v33 = int32(0)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v26))|base.B2i32(int32(1)<<(uint(v26)%32)&int32(553) == v33) == v33 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[1])))
	if v39 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L10
	} else {
		goto L62
	}
L15:
	;
	v43 = int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if base.Ui32(v44) < base.Ui32(int32(_a_F_DefineQueryRewrite_1)) {
		v53 = v43
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[2]))
	v57 = F_object_ownercheck(m, int32(1259), l2, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L23
	}
L18:
	;
	if v53 != 0 {
		goto L9
	} else {
		goto L22
	}
L19:
	;
	goto L18
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	if v48 == int32(99) {
		v53 = v43
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v51 = F_isTempToastNamespace(m, v48)
	mBase = m.M
	v53 = v51
	goto L19
L22:
	;
	goto L17
L23:
	;
	if v57 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62)+119)))
	switch v63 - int32(73) {
	case 0, 32:
		v73 = int32(20)
		goto L28
	default:
		goto L29
	case 10:
		goto L33
	case 29:
		goto L30
	case 36:
		goto L31
	case 45:
		goto L32
	}
L25:
	;
	goto L26
L26:
	;
	if l7 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	F_aclcheck_error(m, int32(2), v75, v76+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L34
	}
L28:
	;
	v75 = v73
	goto L27
L29:
	;
	v73 = int32(42)
	goto L28
L30:
	;
	v75 = int32(18)
	goto L27
L31:
	;
	v75 = int32(23)
	goto L27
L32:
	;
	v75 = int32(52)
	goto L27
L33:
	;
	v75 = int32(38)
	goto L27
L34:
	;
	goto L26
L35:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if int32(0) < v81 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if l4 == int32(1) {
		goto L8
	} else {
		goto L61
	}
L38:
	;
	v94 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	if l4 != int32(1) {
		goto L7
	} else {
		goto L60
	}
L41:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v94<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	if v105 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	v156 = v94 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v156 < v157 {
		v94 = v156
		goto L41
	} else {
		goto L59
	}
L44:
	;
	v109 = F_getInsertSelectQuery(m, v104, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	if v109 != v104 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	switch v112 - int32(1) {
	case 0:
		goto L48
	case 1:
		goto L47
	default:
		goto L43
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L54
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_2), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	F_errhint(m, int32(_a_F_DefineQueryRewrite_3), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(295), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_6), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errhint(m, int32(_a_F_DefineQueryRewrite_7), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(300), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	goto L42
L60:
	;
	goto L8
L61:
	;
	goto L6
L62:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v185 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_8), v18)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v192)+119)))
	F_errdetail_relkind_not_supported(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(265), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v208 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_9), v18+int32(96))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(271), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
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
	if l7 != 0 {
		goto L82
	} else {
		goto L83
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v248 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_10), v18+int32(16))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v258 = int32(*(*int8)(unsafe.Add(mBase, uint32(v257)+119)))
	F_errdetail_relkind_not_supported(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(316), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L10
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L10
	} else {
		goto L134
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L10
	} else {
		goto L130
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L126
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L10
	} else {
		goto L122
	}
L82:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if int32(2) <= v266 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L117
	}
L85:
	;
	if l5 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v273 != int32(1) {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+42)))
	if v276 == int32(1) {
		goto L79
	} else {
		goto L88
	}
L88:
	;
	if l3 != 0 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v272)+76))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	F_checkRuleResultList(m, v279, v280, int32(1), base.B2i32(v238 != int32(109)))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	if l6 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v356 = int32(_a_F_DefineQueryRewrite_0)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[0])))
	if base.B2i32(v359 == int32(0))|base.B2i32(v359 != v362) != 0 {
		v380 = v359
		v381 = v362
		goto L106
	} else {
		goto L107
	}
L92:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	if v286 == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	if v289 <= int32(0) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286)+4))
	v303 = int32(0)
	goto L95
L95:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v292+v303<<(uint(int32(2))%32))))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v313 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L10
	} else {
		goto L101
	}
L97:
	;
	v317 = v303 + int32(1)
	if v289 != v317 {
		v303 = v317
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L91
L101:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v326 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_11), v18-int32(-64))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(387), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	if v380-v381 == int32(0) {
		v591 = int32(105)
		goto L5
	} else {
		goto L112
	}
L106:
	;
	goto L105
L107:
	;
	v365 = l1
	v366 = v356
	goto L108
L108:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+1)))
	if v370 == int32(0) {
		v380 = v370
		v381 = v369
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v380 = v370
	v381 = v369
	goto L106
L110:
	;
	v373 = int32(1)
	if v370 == v369 {
		v365 = v365 + v373
		v366 = v366 + v373
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = int32(_a_F_DefineQueryRewrite_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v392 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_12), v18+int32(48))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(399), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_13), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L119
	}
L119:
	;
	F_errhint(m, int32(_a_F_DefineQueryRewrite_14), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(325), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_15), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(333), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L10
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_16), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(343), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L10
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L10
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_17), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(351), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L10
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L10
	} else {
		goto L135
	}
L135:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_18), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L10
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(359), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L10
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v506 = v493
	v509 = int32(0)
	goto L139
L139:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v512+v506<<(uint(int32(2))%32))))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+96))
	if v517 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L6
L141:
	;
	if v509 != 0 {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	v526 = v509
	goto L143
L143:
	;
	v528 = v506 + int32(1)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v528 < v529 {
		v506 = v528
		v509 = v526
		goto L139
	} else {
		goto L148
	}
L144:
	;
	if l3 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	if l5 == int32(0) {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v521 = int32(0)
	F_checkRuleResultList(m, v517, v520, v521, v521)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	v526 = int32(1)
	goto L143
L148:
	;
	goto L140
L149:
	;
	if v570-v571 == int32(0) {
		goto L1
	} else {
		goto L156
	}
L150:
	;
	goto L149
L151:
	;
	v555 = l1
	v556 = v546
	goto L152
L152:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+1)))
	if v560 == int32(0) {
		v570 = v560
		v571 = v559
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v570 = v560
	v571 = v559
	goto L150
L154:
	;
	v563 = int32(1)
	if v560 == v559 {
		v555 = v555 + v563
		v556 = v556 + v563
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v591 = int32(97)
	goto L5
L157:
	;
	v815 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2618)
	F_relation_close(m, v21, v815)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L10
	} else {
		goto L220
	}
L158:
	;
	v594 = int32(0)
	if l7 == v594 {
		v811 = v594
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v598 = F_nodeToString(m, l3)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L10
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	v600 = F_nodeToString(m, l7)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+200)) = int64(0)
	v605 = v18 + int32(136)
	v607 = F_strncpy(m, v605, l1, int32(64))
	mBase = m.M
	v608 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v607)+63)) = uint8(v608)
	goto L164
L164:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+248)) = base.I64_extend_i32_u(l5)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+240)) = int64(79)
	v615 = int64(56)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+232)) = (base.I64_extend_i32_u(l4)<<(uint(v615)%64) + int64(3458764513820540928)) >> (uint(v615) % 64)
	v622 = base.I64_extend_i32_u(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+224)) = v622
	*(*int64)(unsafe.Add(mBase, uint32(v18)+216)) = base.I64_extend_i32_u(v605)
	v626 = F_cstring_to_text(m, v598)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+256)) = base.I64_extend_i32_u(v626)
	v630 = F_cstring_to_text(m, v600)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+264)) = base.I64_extend_i32_u(v630)
	v636 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	v640 = F_SearchSysCache2(m, int32(60), v622, base.I64_extend_i32_u(l1))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L10
	} else {
		goto L170
	}
L168:
	;
	v710 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = int32(2618)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = int32(1259)
	v721 = v18 + int32(120)
	F_recordDependencyOn(m, v721, v18+int32(108), v591)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L10
	} else {
		goto L189
	}
L169:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v636)+52))
	v689 = F_heap_modify_tuple(m, v640, v682, v18+int32(208), v18+int32(200), v18+int32(120))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L10
	} else {
		goto L184
	}
L170:
	;
	if v640 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = int64(72340168543043584)
	if l6 != 0 {
		goto L169
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v667 = F_GetNewOidWithIndex(m, v636, int32(2692), int32(1))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L10
	} else {
		goto L180
	}
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L10
	} else {
		goto L175
	}
L175:
	;
	F_errcode(m, int32(_a_F_DefineQueryRewrite_19))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	v651 = F_get_rel_name(m, l2)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v651
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l1
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_20), v18+int32(32))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(105), int32(_a_F_DefineQueryRewrite_21))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L10
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
	*(*int64)(unsafe.Add(mBase, uint32(v18)+208)) = base.I64_extend_i32_u(v667)
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v636)+52))
	v676 = F_heap_form_tuple(m, v671, v18+int32(208), v18+int32(200))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L10
	} else {
		goto L181
	}
L181:
	;
	F_CatalogTupleInsert(m, v636, v676)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L10
	} else {
		goto L182
	}
L182:
	;
	F_pfree(m, v676)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	v709 = v667
	goto L168
L184:
	;
	F_CatalogTupleUpdate(m, v636, v689+int32(4), v689)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L10
	} else {
		goto L185
	}
L185:
	;
	F_ReleaseCatCache(m, v640)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v689)+16))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+22)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v697+v698)))
	F_pfree(m, v689)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	v705 = F_deleteDependencyRecordsFor(m, int32(2618), v700, int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	v709 = v700
	goto L168
L189:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[2]))
	F_CheckUsageOnTypesInExpr(m, l7, int32(0), v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_recordDependencyOnExpr(m, v721, l7, int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	if l3 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)))
	v737 = F_getInsertSelectQuery(m, v735, int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L10
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[3]))
	if v749 != 0 {
		goto L198
	} else {
		goto L199
	}
L195:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v737)+52))
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_DefineQueryRewrite[2]))
	F_CheckUsageOnTypesInExpr(m, l3, v739, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v737)+52))
	F_recordDependencyOnExpr(m, v721, l3, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	v751 = int32(0)
	F_RunObjectPostCreateHook(m, int32(2618), v709, v751, v751)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L10
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	F_relation_close(m, v636, int32(3))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L10
	} else {
		goto L202
	}
L201:
	;
	goto L200
L202:
	;
	v758 = m.G0
	v760 = v758 - int32(16)
	m.G0 = v760
	v764 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	v769 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(l2), int64(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L10
	} else {
		goto L205
	}
L204:
	;
	v811 = v709
	goto L157
L205:
	;
	if v769 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v769)+16))
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+22)))
	v773 = v771 + v772
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773)+124)))
	if v774 != int32(1) {
		goto L210
	} else {
		goto L211
	}
L207:
	;
	goto L208
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L10
	} else {
		goto L217
	}
L209:
	;
	F_pfree(m, v769)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L10
	} else {
		goto L215
	}
L210:
	;
	v777 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v773)+124)) = uint8(v777)
	F_CatalogTupleUpdate(m, v764, v769+int32(4), v769)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L10
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	F_CacheInvalidateRelcacheByTuple(m, v769)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L10
	} else {
		goto L214
	}
L213:
	;
	goto L209
L214:
	;
	goto L209
L215:
	;
	F_relation_close(m, v764, int32(3))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L10
	} else {
		goto L216
	}
L216:
	;
	m.G0 = v760 + int32(16)
	goto L204
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v760))) = l2
	F_errmsg_internal(m, int32(_a_F_DefineQueryRewrite_22), v760)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_23), int32(65), int32(_a_F_DefineQueryRewrite_24))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L10
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	m.G0 = v18 + int32(272)
	return
L221:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L10
	} else {
		goto L222
	}
L222:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_25), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L10
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(422), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L10
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_26), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(427), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_27), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(431), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(_a_F_DefineQueryRewrite_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v881 + int32(4)
	F_errmsg(m, int32(_a_F_DefineQueryRewrite_28), v18+int32(80))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_DefineQueryRewrite_4), int32(447), int32(_a_F_DefineQueryRewrite_5))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L10
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_assign_query_collations(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v5 = F_query_tree_walker_impl(m, l1, int32(517), l0, int32(10))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_extract_query_dependencies_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	v3 = int32(0)
	if l0 == v3 {
		v157 = v3
		v158 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v157 == int32(0)) & v158
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v9 != int32(67) {
		v84 = l0
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+44)))
	if v103 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L4:
	;
	F_fix_expr_common(m, l1, v84)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L38
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 != int32(6) {
		v97 = l0
		v101 = v3
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v16 != int32(213) {
		v50 = v15
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v57 = base.B2i32(v16 == int32(213))
	v59 = v50
	goto L24
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v20 = F_extract_query_dependencies_walker(m, v19, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v24 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v157 = int32(1)
	v158 = v3
	goto L1
L12:
	;
	goto L13
L13:
	;
	v28 = v24
	goto L14
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v34 != int32(67) {
		v84 = v28
		goto L4
	} else {
		goto L16
	}
L15:
	;
	v157 = int32(1)
	v158 = v3
	goto L1
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v37 != int32(6) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v97 = v28
	v101 = int32(1)
	goto L3
L18:
	;
	goto L19
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 != int32(213) {
		v50 = v41
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = F_extract_query_dependencies_walker(m, v45, l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v49 != 0 {
		v28 = v49
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	if v83 != 0 {
		v97 = v83
		v101 = v57
		goto L3
	} else {
		goto L37
	}
L24:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch v61 - int32(241) {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		goto L30
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v59 = v81
	goto L24
L27:
	;
	v83 = v79
	goto L23
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v76 == int32(6) {
		v80 = v75
		goto L26
	} else {
		goto L36
	}
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 == int32(6) {
		v80 = v71
		goto L26
	} else {
		goto L35
	}
L30:
	;
	if v61 != int32(201) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v83 = int32(0)
	goto L23
L32:
	;
	goto L33
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 != int32(6) {
		v79 = v67
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v80 = v67
	goto L26
L35:
	;
	v79 = v71
	goto L27
L36:
	;
	v79 = v75
	goto L27
L37:
	;
	v157 = v57
	v158 = v3
	goto L1
L38:
	;
	v95 = F_expression_tree_walker_impl(m, v84, int32(886), l1)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v157 = base.B2i32(v9 == int32(67))
	v158 = v95
	goto L1
L40:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v106)+93)) = uint8(v107)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+52))
	if v109 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v151 = F_query_tree_walker_impl(m, v97, int32(886), l1, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L9
	} else {
		goto L55
	}
L44:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v120 = v3
	goto L46
L46:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v120<<(uint(int32(2))%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	switch v126 {
	case 0:
		goto L51
	case 1, 7:
		goto L50
	default:
		goto L48
	}
L47:
	;
	goto L43
L48:
	;
	v140 = v120 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v140 < v141 {
		v120 = v140
		goto L46
	} else {
		goto L54
	}
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+64))
	v134 = F_lappend_oid(m, v133, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L53
	}
L50:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	if v128 == int32(0) {
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v131 = v127
	goto L49
L52:
	;
	v131 = v128
	goto L49
L53:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+64)) = v134
	goto L48
L54:
	;
	goto L47
L55:
	;
	v157 = v101
	v158 = v151
	goto L1
}
func F_query_to_xml_and_xmlschema(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = F_text_to_cstring(m, v15)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = F_text_to_cstring(m, v23)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = int32(0)
						v32 = F_SPI_prepare(m, v19, v30, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							if v32 != 0 {
								v34 = F_SPI_cursor_open(m, v32)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int64(0)
								} else {
									if v34 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v19
											F_errmsg_internal(m, int32(_a_F_query_to_xml_and_xmlschema_0), v12+int32(16))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_query_to_xml_and_xmlschema_1), int32(3204), int32(_a_F_query_to_xml_and_xmlschema_2))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
										v41 = base.B2i32(v21 != int64(0))
										v42 = F_map_sql_table_to_xmlschema(m, v38, int32(0), v41, v25)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int64(0)
										} else {
											v44 = F_strlen(m, v42)
											mBase = m.M
											v46 = v44 + int32(1)
											v47 = F_SPI_palloc(m, v46)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return int64(0)
											} else {
												if v46 != 0 {
													base.MemoryCopy(m, v47, v42, v46)
												} else {
												}
												F_SPI_cursor_close(m, v34)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return int64(0)
												} else {
													v52 = F_SPI_finish(m)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return int64(0)
													} else {
														v55 = F_query_to_xml_internal(m, v19, int32(0), v47, v41, v25)
														mBase = m.M
														v56 = m.ExcPending
														if v56 != 0 {
															return int64(0)
														} else {
															v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
															v59 = F_cstring_to_text_with_len(m, v57, v58)
															mBase = m.M
															v60 = m.ExcPending
															if v60 != 0 {
																return int64(0)
															} else {
																m.G0 = v12 + int32(32)
																return base.I64_extend_i32_u(v59)
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v19
									F_errmsg_internal(m, int32(_a_F_query_to_xml_and_xmlschema_3), v12)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_query_to_xml_and_xmlschema_1), int32(3201), int32(_a_F_query_to_xml_and_xmlschema_2))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
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
func F_query_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
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
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v291 int32
	_ = v291
	v5 = int32(0)
	if l3&int32(64) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v24 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v23, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v22 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = F_palloc(m, int32(168))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	base.MemoryCopy(m, v16, l0, int32(168))
	v22 = v16
	goto L1
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+152))
	v28 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v27, l2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+152)) = v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v32 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v31, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v36 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v35, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v44 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v43, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	v48 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v47, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+144))
	v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+112))
	v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v55, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v22)+128))
	v60 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v59, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v22)+132))
	v64 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v63, l2)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v64
	if l3&int32(128) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if l3&int32(2) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v70 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v69, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)+116))
	if v85 == int32(0) {
		v145 = v5
		goto L26
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)+116))
	v74 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v73, l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v78 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v77, l2)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+124)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v22)+120))
	v82 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v81, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+120)) = v82
	goto L18
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v145
	goto L18
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v88 <= int32(0) {
		v145 = v5
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = v5
	v99 = v5
	goto L29
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v99<<(uint(int32(2))%32))))
	v109 = F_palloc(m, int32(56))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v145 = v133
	goto L26
L31:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v107)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+48)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v107)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+40)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v107)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+32)) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v107)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+24)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+16)) = v119
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v107)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = v121
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v107)+24))
	v126 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v125, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v107)+28))
	v130 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v129, l2)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = v130
	v133 = F_lappend(m, v97, v109)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v136 = v99 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v136 < v137 {
		v97 = v133
		v99 = v136
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v173
	v175 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	if v178 == v175 {
		v291 = v175
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v169 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v164, l2)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v171 = F_copyObjectImpl(m, v164)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v173 = v169
	goto L36
L41:
	;
	v173 = v171
	goto L36
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v291
	return v22
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if int32(0) < v181 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v195 = v175
	v198 = v175
	goto L47
L45:
	;
	v272 = v175
	goto L46
L46:
	;
	v291 = v272
	goto L42
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v198<<(uint(int32(2))%32))))
	v208 = F_palloc(m, int32(136))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L5
	} else {
		goto L49
	}
L48:
	;
	v272 = v261
	goto L46
L49:
	;
	base.MemoryCopy(m, v208, v206, int32(136))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	switch v212 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	case 4:
		goto L53
	case 5:
		goto L52
	default:
		goto L50
	case 9:
		goto L51
	}
L50:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v206)+128))
	v258 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v257, l2)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L77
	}
L51:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)+120))
	if l3&int32(256) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L52:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v206)+80))
	v244 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v243, l2)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L71
	}
L53:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v206)+76))
	v240 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v239, l2)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L70
	}
L54:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v206)+68))
	v236 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v235, l2)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L69
	}
L55:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v206)+52))
	if l3&int32(4) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	if l3&int32(1) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	v214 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v213, l2)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+32)) = v214
	goto L50
L59:
	;
	v220 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v217, l2)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v223 = F_copyObjectImpl(m, v217)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+36)) = v220
	goto L50
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+36)) = v223
	goto L50
L64:
	;
	v229 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v226, l2)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v232 = F_copyObjectImpl(m, v226)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+52)) = v229
	goto L50
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+52)) = v232
	goto L50
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+68)) = v236
	goto L50
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+76)) = v240
	goto L50
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+80)) = v244
	goto L50
L72:
	;
	v250 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v247, l2)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v253 = F_copyObjectImpl(m, v247)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+120)) = v250
	goto L50
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+120)) = v253
	goto L50
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+128)) = v258
	v261 = F_lappend(m, v195, v208)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v264 = v198 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	if v264 < v265 {
		v195 = v261
		v198 = v264
		goto L47
	} else {
		goto L79
	}
L79:
	;
	goto L48
}
func F_query_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v9 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v8, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v9 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v14 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v13, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v17 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v16, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v17 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v19, l2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v23 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v22, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	if v23 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v26 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v25, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v26 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v29 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v28, l2)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v29 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v32 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v31, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v35 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v34, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v35 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v38 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v37, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	if v38 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v41 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v40, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v41 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if l3&int32(128) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if l3&int32(2) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v46 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v45, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v57 == int32(0) {
		goto L25
	} else {
		goto L37
	}
L29:
	;
	if v46 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v48, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v49 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v52 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v55 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v54, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v55 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v68 = int32(0)
	goto L39
L39:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v68<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v76 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v75, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L1
L41:
	;
	goto L40
L42:
	;
	if v76 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v79 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v78, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v79 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v82 = v68 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v82 < v83 {
		v68 = v82
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L25
L47:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v97 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v96, l2)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if l3&int32(8) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v97 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v104 = F_range_table_walker_impl(m, v103, l1, l2, l3)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	return int32(0)
L55:
	;
	if v104 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
}
