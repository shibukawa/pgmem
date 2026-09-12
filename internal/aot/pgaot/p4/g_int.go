package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v566 int64
	_ = v566
	var v570 int32
	_ = v570
	var v571 int64
	_ = v571
	var v574 int64
	_ = v574
	var v575 int64
	_ = v575
	var v577 int32
	_ = v577
	var v579 int64
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v782 int64
	_ = v782
	var v788 int64
	_ = v788
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int64
	_ = v800
	var v806 int64
	_ = v806
	var v808 int32
	_ = v808
	var v812 int64
	_ = v812
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == v2 {
		v39 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v26 == int32(0) {
		v39 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(7) {
		v39 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v32 != int32(17) {
		v39 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v39 = v35 ^ int32(1)
	goto L2
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_fn_opclass_options(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v48 = int32(100)
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v50 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v48 = v47
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L10
	} else {
		goto L232
	}
L13:
	;
	m.G0 = v18 + int32(16)
	return v839
L14:
	;
	if v285 < int32(0) {
		v365 = v289
		goto L91
	} else {
		goto L92
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L10
	} else {
		goto L87
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L82
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L78
	}
L18:
	;
	v53 = F_pg_detoast_datum_copy(m, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v100 = F_pg_detoast_datum(m, v49)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L35
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v55 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v56 = F_array_contains_nulls(m, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v61 = F_ArrayGetNItems(m, v58, v53+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L27
	}
L25:
	;
	if v56 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v63 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)) = uint8(v63)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v65 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = v65
	goto L30
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v73 = (v66<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L30
L30:
	;
	F_isort(m, v73+v53, v61, v18+int32(15))
	mBase = m.M
	v78 = F__int_unique(m, v53)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v82 = v78 + int32(16)
	v83 = F_ArrayGetNItems(m, v80, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v86 = v48 << (uint(int32(1)) % 32)
	if v86 <= v83 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v89 = F_palloc(m, int32(16))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v78
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v94
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)))
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+14)) = uint8(v97)
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+12)) = uint16(v96)
	v839 = v89
	goto L13
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v103 = F_array_contains_nulls(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v107 = v100 + int32(16)
	v108 = F_ArrayGetNItems(m, v105, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	if v103 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v108 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v112 == v100 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v117 = F_ArrayGetNItems(m, v116, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L49
	}
L45:
	;
	v839 = v20
	goto L13
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v100)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	v839 = v20
	goto L13
L49:
	;
	v120 = v48 << (uint(int32(1)) % 32)
	if v117 < v120 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v839 = v20
	goto L13
L51:
	;
	goto L52
L52:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v122 == v100 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v124 = F_pg_detoast_datum_copy(m, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L56
	}
L54:
	;
	v126 = v100
	goto L55
L55:
	;
	v129 = F_resize_intArrayType(m, v126, v117<<(uint(int32(1))%32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	v126 = v124
	goto L55
L57:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v131 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v141 = (v134<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L60
L59:
	;
	v141 = v131
	goto L60
L60:
	;
	v143 = v117 - int32(1)
	v144 = v141 + v129
	if int32(2) <= v117 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v153 = v143
	v154 = v147
	v159 = v143
	goto L66
L62:
	;
	v147 = v117 - v48
	if int32(0) < v147 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v285 = v143
	v289 = v143
	goto L14
L65:
	;
	goto L64
L66:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v144+v153<<(uint(int32(2))%32))))
	v172 = v153
	v173 = v154
	v174 = v171
	goto L68
L67:
	;
	v285 = v212
	v289 = v210
	goto L14
L68:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v144-int32(4)+v172<<(uint(int32(2))%32))))
	if v190 != v174-int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v206 = v144 + v159<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v171
	v209 = int32(1)
	v210 = v159 - v209
	v212 = v200 - v209
	if v200 < int32(2) {
		v285 = v212
		v289 = v210
		goto L14
	} else {
		goto L76
	}
L70:
	;
	goto L69
L71:
	;
	v200 = v172
	v201 = v173
	v203 = v174
	goto L70
L72:
	;
	goto L73
L73:
	;
	v194 = int32(1)
	v195 = v173 - v194
	v197 = v172 - v194
	if v197 == int32(0) {
		v200 = v197
		v201 = v195
		v203 = v190
		goto L70
	} else {
		goto L74
	}
L74:
	;
	if v195 != 0 {
		v172 = v197
		v173 = v195
		v174 = v190
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v200 = v197
	v201 = v195
	v203 = v190
	goto L70
L76:
	;
	if int32(0) < v201 {
		v153 = v212
		v154 = v201
		v159 = v210
		goto L66
	} else {
		goto L77
	}
L77:
	;
	goto L67
L78:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(152677), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(493034), int32(180), int32(126998))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
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
	F_errcode(m, int32(261))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v245 = F_ArrayGetNItems(m, v244, v82)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v86 - int32(1)
	F_errmsg(m, int32(464269), v18)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(493034), int32(187), int32(126998))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(152677), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(493034), int32(202), int32(126998))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v375 = v365 + int32(1)
	if v375 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	if v285&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v317 = v289
	v318 = v285
	goto L95
L94:
	;
	v304 = v144 + v289<<(uint(int32(3))%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v144+v285<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v304)+4)) = v308
	v311 = int32(1)
	v317 = v289 - v311
	v318 = v285 - v311
	goto L95
L95:
	;
	if v285 == int32(0) {
		v365 = v317
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v321 = v318
	v327 = v317
	goto L97
L97:
	;
	v338 = v144 + v327<<(uint(int32(3))%32)
	v339 = int32(2)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v144+v321<<(uint(v339)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v338))) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v338)+4)) = v342
	v346 = v338 - int32(8)
	v348 = v321 - int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v144+v348<<(uint(v339)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v346)+4)) = v352
	v356 = v327 - v339
	if v348 != 0 {
		v321 = v321 - v339
		v327 = v356
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v365 = v356
	goto L91
L99:
	;
	goto L98
L100:
	;
	v530 = int32(1)
	v532 = v529 << (uint(v530) % 32)
	if v120 < v532 {
		goto L150
	} else {
		goto L151
	}
L101:
	;
	v529 = v117 - v375
	goto L100
L102:
	;
	goto L103
L103:
	;
	v379 = int32(3)
	v381 = v144 + v375<<(uint(v379)%32)
	v382 = v117 - v375
	v384 = v382 << (uint(v379) % 32)
	if v144 == v381 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v529 = v382
	goto L100
L105:
	;
	goto L104
L106:
	;
	v388 = v144 + v384
	if base.Ui32(v381-v388) <= base.Ui32(int32(0)-v384<<(uint(int32(1))%32)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v395 = F___memcpy(m, v144, v381, v384)
	mBase = m.M
	goto L104
L108:
	;
	goto L109
L109:
	;
	v398 = (v144 ^ v381) & int32(3)
	if base.Ui32(v144) < base.Ui32(v381) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	if v500 == int32(0) {
		goto L105
	} else {
		goto L146
	}
L111:
	;
	if base.Ui32(v478) <= base.Ui32(int32(3)) {
		v499 = v477
		v500 = v478
		v501 = v479
		goto L110
	} else {
		goto L142
	}
L112:
	;
	if v398 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if v398 != 0 {
		v460 = v384
		goto L125
	} else {
		goto L126
	}
L115:
	;
	v499 = v381
	v500 = v384
	v501 = v144
	goto L110
L116:
	;
	goto L117
L117:
	;
	if v144&int32(3) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v477 = v381
	v478 = v384
	v479 = v144
	goto L111
L119:
	;
	goto L120
L120:
	;
	v405 = v381
	v406 = v384
	v407 = v144
	goto L121
L121:
	;
	if v406 == int32(0) {
		goto L105
	} else {
		goto L123
	}
L122:
	;
	v477 = v414
	v478 = v416
	v479 = v418
	goto L111
L123:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	*(*uint8)(unsafe.Add(mBase, uint32(v407))) = uint8(v411)
	v413 = int32(1)
	v414 = v405 + v413
	v416 = v406 - v413
	v418 = v407 + v413
	if v418&int32(3) != 0 {
		v405 = v414
		v406 = v416
		v407 = v418
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v460 == int32(0) {
		goto L105
	} else {
		goto L138
	}
L126:
	;
	if v388&int32(3) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v425 = v384
	goto L130
L128:
	;
	v440 = v384
	goto L129
L129:
	;
	if base.Ui32(v440) <= base.Ui32(int32(3)) {
		v460 = v440
		goto L125
	} else {
		goto L134
	}
L130:
	;
	if v425 == int32(0) {
		goto L105
	} else {
		goto L132
	}
L131:
	;
	v440 = v431
	goto L129
L132:
	;
	v431 = v425 - int32(1)
	v432 = v144 + v431
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381+v431))))
	*(*uint8)(unsafe.Add(mBase, uint32(v432))) = uint8(v434)
	if v432&int32(3) != 0 {
		v425 = v431
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v447 = v440
	goto L135
L135:
	;
	v451 = v447 - int32(4)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v381+v451)))
	*(*int32)(unsafe.Add(mBase, uint32(v144+v451))) = v454
	if base.Ui32(int32(3)) < base.Ui32(v451) {
		v447 = v451
		goto L135
	} else {
		goto L137
	}
L136:
	;
	v460 = v451
	goto L125
L137:
	;
	goto L136
L138:
	;
	v467 = v460
	goto L139
L139:
	;
	v471 = v467 - int32(1)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381+v471))))
	*(*uint8)(unsafe.Add(mBase, uint32(v144+v471))) = uint8(v474)
	if v471 != 0 {
		v467 = v471
		goto L139
	} else {
		goto L141
	}
L140:
	;
	goto L105
L141:
	;
	goto L140
L142:
	;
	v484 = v477
	v485 = v478
	v486 = v479
	goto L143
L143:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = v488
	v490 = int32(4)
	v491 = v484 + v490
	v493 = v486 + v490
	v495 = v485 - v490
	if base.Ui32(int32(3)) < base.Ui32(v495) {
		v484 = v491
		v485 = v495
		v486 = v493
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v499 = v491
	v500 = v495
	v501 = v493
	goto L110
L145:
	;
	goto L144
L146:
	;
	v506 = v499
	v507 = v500
	v508 = v501
	goto L147
L147:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	*(*uint8)(unsafe.Add(mBase, uint32(v508))) = uint8(v510)
	v512 = int32(1)
	v517 = v507 - v512
	if v517 != 0 {
		v506 = v506 + v512
		v507 = v517
		v508 = v508 + v512
		goto L147
	} else {
		goto L149
	}
L148:
	;
	goto L105
L149:
	;
	goto L148
L150:
	;
	v535 = v530
	v537 = v532
	goto L153
L151:
	;
	v760 = v532
	goto L152
L152:
	;
	if v760 <= int32(0) {
		goto L216
	} else {
		goto L217
	}
L153:
	;
	if int32(3) <= v537 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v760 = v755
	goto L152
L155:
	;
	v553 = int32(2)
	v554 = v535
	v566 = int64(9223372036854775807)
	goto L158
L156:
	;
	v584 = v535
	goto L157
L157:
	;
	v598 = int32(2)
	v600 = v144 + v584<<(uint(v598)%32)
	v601 = int32(4)
	v602 = v600 - v601
	v604 = v600 + v601
	v609 = (v537 + (v584 ^ int32(-1))) << (uint(v598) % 32)
	if v602 == v604 {
		goto L168
	} else {
		goto L169
	}
L158:
	;
	v570 = v144 + v553<<(uint(int32(2))%32)
	v571 = int64(*(*int32)(unsafe.Add(mBase, uint32(v570))))
	v574 = int64(*(*int32)(unsafe.Add(mBase, uint32(v570-int32(4)))))
	v575 = v571 - v574
	if v575 < v566 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v584 = v577
	goto L157
L160:
	;
	v577 = v553
	goto L162
L161:
	;
	v577 = v554
	goto L162
L162:
	;
	if v566 < v575 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v579 = v566
	goto L165
L164:
	;
	v579 = v575
	goto L165
L165:
	;
	v581 = v553 + int32(2)
	if v581 < v537 {
		v553 = v581
		v554 = v577
		v566 = v579
		goto L158
	} else {
		goto L166
	}
L166:
	;
	goto L159
L167:
	;
	v755 = v537 - int32(2)
	if v120 < v755 {
		v535 = v584
		v537 = v755
		goto L153
	} else {
		goto L213
	}
L168:
	;
	goto L167
L169:
	;
	v613 = v602 + v609
	if base.Ui32(v604-v613) <= base.Ui32(int32(0)-v609<<(uint(int32(1))%32)) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v620 = F___memcpy(m, v602, v604, v609)
	mBase = m.M
	goto L167
L171:
	;
	goto L172
L172:
	;
	v623 = (v602 ^ v604) & int32(3)
	if base.Ui32(v602) < base.Ui32(v604) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	if v725 == int32(0) {
		goto L168
	} else {
		goto L209
	}
L174:
	;
	if base.Ui32(v703) <= base.Ui32(int32(3)) {
		v724 = v702
		v725 = v703
		v726 = v704
		goto L173
	} else {
		goto L205
	}
L175:
	;
	if v623 != 0 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	if v623 != 0 {
		v685 = v609
		goto L188
	} else {
		goto L189
	}
L178:
	;
	v724 = v604
	v725 = v609
	v726 = v602
	goto L173
L179:
	;
	goto L180
L180:
	;
	if v602&int32(3) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v702 = v604
	v703 = v609
	v704 = v602
	goto L174
L182:
	;
	goto L183
L183:
	;
	v630 = v604
	v631 = v609
	v632 = v602
	goto L184
L184:
	;
	if v631 == int32(0) {
		goto L168
	} else {
		goto L186
	}
L185:
	;
	v702 = v639
	v703 = v641
	v704 = v643
	goto L174
L186:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	*(*uint8)(unsafe.Add(mBase, uint32(v632))) = uint8(v636)
	v638 = int32(1)
	v639 = v630 + v638
	v641 = v631 - v638
	v643 = v632 + v638
	if v643&int32(3) != 0 {
		v630 = v639
		v631 = v641
		v632 = v643
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	if v685 == int32(0) {
		goto L168
	} else {
		goto L201
	}
L189:
	;
	if v613&int32(3) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v650 = v609
	goto L193
L191:
	;
	v665 = v609
	goto L192
L192:
	;
	if base.Ui32(v665) <= base.Ui32(int32(3)) {
		v685 = v665
		goto L188
	} else {
		goto L197
	}
L193:
	;
	if v650 == int32(0) {
		goto L168
	} else {
		goto L195
	}
L194:
	;
	v665 = v656
	goto L192
L195:
	;
	v656 = v650 - int32(1)
	v657 = v602 + v656
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604+v656))))
	*(*uint8)(unsafe.Add(mBase, uint32(v657))) = uint8(v659)
	if v657&int32(3) != 0 {
		v650 = v656
		goto L193
	} else {
		goto L196
	}
L196:
	;
	goto L194
L197:
	;
	v672 = v665
	goto L198
L198:
	;
	v676 = v672 - int32(4)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v604+v676)))
	*(*int32)(unsafe.Add(mBase, uint32(v602+v676))) = v679
	if base.Ui32(int32(3)) < base.Ui32(v676) {
		v672 = v676
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v685 = v676
	goto L188
L200:
	;
	goto L199
L201:
	;
	v692 = v685
	goto L202
L202:
	;
	v696 = v692 - int32(1)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604+v696))))
	*(*uint8)(unsafe.Add(mBase, uint32(v602+v696))) = uint8(v699)
	if v696 != 0 {
		v692 = v696
		goto L202
	} else {
		goto L204
	}
L203:
	;
	goto L168
L204:
	;
	goto L203
L205:
	;
	v709 = v702
	v710 = v703
	v711 = v704
	goto L206
L206:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	*(*int32)(unsafe.Add(mBase, uint32(v711))) = v713
	v715 = int32(4)
	v716 = v709 + v715
	v718 = v711 + v715
	v720 = v710 - v715
	if base.Ui32(int32(3)) < base.Ui32(v720) {
		v709 = v716
		v710 = v720
		v711 = v718
		goto L206
	} else {
		goto L208
	}
L207:
	;
	v724 = v716
	v725 = v720
	v726 = v718
	goto L173
L208:
	;
	goto L207
L209:
	;
	v731 = v724
	v732 = v725
	v733 = v726
	goto L210
L210:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731))))
	*(*uint8)(unsafe.Add(mBase, uint32(v733))) = uint8(v735)
	v737 = int32(1)
	v742 = v732 - v737
	if v742 != 0 {
		v731 = v731 + v737
		v732 = v742
		v733 = v733 + v737
		goto L210
	} else {
		goto L212
	}
L211:
	;
	goto L168
L212:
	;
	goto L211
L213:
	;
	goto L154
L214:
	;
	if base.Ui32(int32(134217725)) <= base.Ui32(v822) {
		goto L12
	} else {
		goto L229
	}
L215:
	;
	if base.Ui64(v812-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L226
	} else {
		goto L227
	}
L216:
	;
	v812 = int64(0)
	goto L215
L217:
	;
	goto L218
L218:
	;
	v778 = int64(*(*int32)(unsafe.Add(mBase, uint32(v144)+4)))
	v779 = int64(*(*int32)(unsafe.Add(mBase, uint32(v144))))
	v782 = v778 - v779 + int64(1)
	if base.Ui32(v760) < base.Ui32(int32(3)) {
		v812 = v782
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v788 = v782
	v789 = int32(2)
	goto L220
L220:
	;
	v794 = v144 + v789<<(uint(int32(2))%32)
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v794-int32(4))))
	if v795 != v798 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v812 = v806
	goto L215
L222:
	;
	v800 = int64(*(*int32)(unsafe.Add(mBase, uint32(v794)+4)))
	v806 = v788 + v800 - base.I64_extend_i32_s(v795) + int64(1)
	goto L224
L223:
	;
	v806 = v788
	goto L224
L224:
	;
	v808 = v789 + int32(2)
	if v808 < v760 {
		v788 = v806
		v789 = v808
		goto L220
	} else {
		goto L225
	}
L225:
	;
	goto L221
L226:
	;
	v822 = int32(-1)
	goto L228
L227:
	;
	v822 = base.I32_wrap_i64(v812)
	goto L228
L228:
	;
	goto L214
L229:
	;
	v825 = F_resize_intArrayType(m, v129, v760)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	v828 = F_palloc(m, int32(16))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828))) = v825
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+4)) = v831
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+8)) = v833
	v835 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+12)))
	v836 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v828)+14)) = uint8(v836)
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+12)) = uint16(v835)
	v839 = v828
	goto L13
L232:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L10
	} else {
		goto L233
	}
L233:
	;
	F_errmsg(m, int32(464195), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(493034), int32(276), int32(126998))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_g_int_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(169296), int32(270864), int32(100), int32(1), int32(252))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
