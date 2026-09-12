package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datum_to_jsonb_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v17 + int32(160)
	return
L4:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v976 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if l5 != 0 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	if base.Ui32(int32(4)) < base.Ui32(l3-int32(6)) {
		goto L4
	} else {
		goto L266
	}
L9:
	;
	v856 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L240
	}
L10:
	;
	v779 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L224
	}
L11:
	;
	v744 = F_pg_detoast_datum_packed(m, v740)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L220
	}
L12:
	;
	v738 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L219
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v673 = int32(0)
	v676 = F_JsonEncodeDateTime(m, v673, l0, int32(1184), v673)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L201
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v607 = int32(0)
	v610 = F_JsonEncodeDateTime(m, v607, l0, int32(1114), v607)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L183
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v541 = int32(0)
	v544 = F_JsonEncodeDateTime(m, v541, l0, int32(1082), v541)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L165
	}
L16:
	;
	v447 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L134
	}
L17:
	;
	v384 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L115
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+76)) = uint8(base.B2i32(l0 != int32(0)))
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	if l0 != 0 {
		goto L109
	} else {
		goto L110
	}
L20:
	;
	v113 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L44
	}
L21:
	;
	v47 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	switch l3 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L17
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		goto L9
	}
L23:
	;
	if base.Ui32(int32(4)) < base.Ui32(l3-int32(6)) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	switch l3 - int32(1) {
	case 0:
		goto L18
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		v740 = l0
		goto L11
	case 6:
		goto L10
	case 7:
		goto L21
	case 8:
		goto L20
	case 9:
		goto L12
	default:
		goto L9
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(233004), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(476938), int32(664), int32(297103))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v54 = v47 + int32(16)
	v55 = F_ArrayGetNItems(m, v52, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v55
	if v55 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v62 = F_pushJsonbValue(m, l2, int32(4), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_get_typlenbyvalalign(m, v49, v17+int32(150), v17+int32(149), v17+int32(148))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v62
	v67 = F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v67
	v945 = int32(1)
	goto L8
L38:
	;
	v79 = int32(1)
	F_json_categorize_type(m, v49, v79, v17+int32(144), v17+int32(140))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+150)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+149)))
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+148)))
	F_deconstruct_array(m, v47, v87, v88, v89, v17+int32(156), v17+int32(152), v17+int32(72))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+144))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)+140))
	F_array_dim_to_jsonb(m, l2, int32(0), v52, v54, v99, v100, v17+int32(32), v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	F_pfree(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v945 = v79
	goto L8
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v117 = F_lookup_rowtype_tupdesc(m, v115, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(base.Ui32(v119) >> (uint(int32(2)) % 32))
	v126 = F_pushJsonbValue(m, l2, int32(6), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if int32(0) < v129 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v133 = v117 + int32(20)
	v135 = v129
	v136 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	v360 = F_pushJsonbValue(m, l2, int32(7), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L104
	}
L50:
	;
	v154 = v133 + v135<<(uint(int32(4))%32) + v136*int32(100)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+91)))
	if v155 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v341 = v136 + int32(1)
	goto L54
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(1)
	v163 = v154 + int32(4)
	if v163&int32(3) == int32(0) {
		v187 = v163
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v341 < v342 {
		v135 = v342
		v136 = v341
		goto L50
	} else {
		goto L103
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v220
	v226 = F_pushJsonbValue(m, l2, int32(1), v17+int32(32))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L72
	}
L56:
	;
	v220 = v212 - v163
	goto L55
L57:
	;
	v191 = v187
	goto L66
L58:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v171 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v220 = int32(0)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v176 = v163
	goto L62
L62:
	;
	v180 = v176 + int32(1)
	if v180&int32(3) == int32(0) {
		v187 = v180
		goto L57
	} else {
		goto L64
	}
L63:
	;
	v212 = v180
	goto L56
L64:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v185 != 0 {
		v176 = v180
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v200 = int32(-2139062144)
	if (int32(16843008)-v197|v197)&v200 == v200 {
		v191 = v191 + int32(4)
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v206 = v191
	goto L69
L68:
	;
	goto L67
L69:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v210 != 0 {
		v206 = v206 + int32(1)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v212 = v206
	goto L56
L71:
	;
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v226
	v230 = v136 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+18)))
	if base.Ui32(v232&int32(2047)) <= base.Ui32(v136) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)))
	if v307 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L74:
	;
	v238 = F_getmissingattr(m, v117, v230, v17+int32(144))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)) = uint8(v240)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+20)))
	if v242&int32(1) == v240 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v306 = v238
	goto L73
L78:
	;
	v249 = v133 + v136<<(uint(int32(4))%32)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if int32(0) <= v250 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231+int32(base.Ui32(v136)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v288)>>(uint(v136&int32(7))%32))&int32(1) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L81:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)))
	v255 = v231 + v253 + v250
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+6)))
	if v256 != int32(1) {
		v306 = v255
		goto L73
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v283 = F_nocachegetattr(m, v17+int32(72), v230, v117)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L92
	}
L84:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+4)))
	switch v259 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	default:
		goto L85
	case 3:
		goto L86
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L89
	}
L86:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v306 = v264
	goto L73
L87:
	;
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255))))
	v306 = v263
	goto L73
L88:
	;
	v262 = int32(*(*int8)(unsafe.Add(mBase, uint32(v255))))
	v306 = v262
	goto L73
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = base.I32_extend16_s(v259)
	F_errmsg_internal(m, int32(460988), v17+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(310395), int32(70), int32(64556))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v306 = v283
	goto L73
L93:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)) = uint8(v296)
	v306 = int32(0)
	goto L73
L94:
	;
	goto L95
L95:
	;
	v301 = F_nocachegetattr(m, v17+int32(72), v230, v117)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v306 = v301
	goto L73
L97:
	;
	F_datum_to_jsonb_internal(m, v306, v328&int32(1), l2, v329, v330, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L102
	}
L98:
	;
	v310 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v310
	v328 = int32(1)
	v329 = v310
	v330 = v310
	goto L97
L99:
	;
	goto L100
L100:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v154)+68))
	F_json_categorize_type(m, v317, int32(1), v17+int32(156), v17+int32(152))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v328 = v325
	v329 = v326
	v330 = v327
	goto L97
L102:
	;
	v341 = v230
	goto L54
L103:
	;
	goto L51
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v360
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	if int32(0) <= v363 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_DecrTupleDescRefCount(m, v117)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v945 = int32(1)
	goto L8
L108:
	;
	goto L107
L109:
	;
	v373 = int32(327532)
	goto L111
L110:
	;
	v373 = int32(344107)
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v373
	if l0 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v377 = int32(4)
	goto L114
L113:
	;
	v377 = int32(5)
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v377
	goto L4
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	if v384&int32(3) == int32(0) {
		v411 = v384
		goto L118
	} else {
		goto L119
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v444
	goto L4
L117:
	;
	v444 = v436 - v384
	goto L116
L118:
	;
	v415 = v411
	goto L127
L119:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v395 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v444 = int32(0)
	goto L116
L121:
	;
	goto L122
L122:
	;
	v400 = v384
	goto L123
L123:
	;
	v404 = v400 + int32(1)
	if v404&int32(3) == int32(0) {
		v411 = v404
		goto L118
	} else {
		goto L125
	}
L124:
	;
	v436 = v404
	goto L117
L125:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v409 != 0 {
		v400 = v404
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v424 = int32(-2139062144)
	if (int32(16843008)-v421|v421)&v424 == v424 {
		v415 = v415 + int32(4)
		goto L127
	} else {
		goto L129
	}
L128:
	;
	v430 = v415
	goto L130
L129:
	;
	goto L128
L130:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v434 != 0 {
		v430 = v430 + int32(1)
		goto L130
	} else {
		goto L132
	}
L131:
	;
	v436 = v430
	goto L117
L132:
	;
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	if v447&int32(3) == int32(0) {
		v503 = v447
		goto L150
	} else {
		goto L151
	}
L134:
	;
	v449 = int32(78)
	v450 = F___strchrnul(m, v447, v449)
	mBase = m.M
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	if v452 == v449 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v456 != 0 {
		goto L133
	} else {
		goto L139
	}
L136:
	;
	v456 = v450
	goto L138
L137:
	;
	v456 = int32(0)
	goto L138
L138:
	;
	goto L135
L139:
	;
	v457 = int32(110)
	v458 = F___strchrnul(m, v447, v457)
	mBase = m.M
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v460 == v457 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v464 != 0 {
		goto L133
	} else {
		goto L144
	}
L141:
	;
	v464 = v458
	goto L143
L142:
	;
	v464 = int32(0)
	goto L143
L143:
	;
	goto L140
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(2)
	v468 = int32(0)
	v471 = F_DirectFunctionCall3Coll(m, int32(408), v468, v447, v468, int32(-1))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v473 = F_pg_detoast_datum(m, v471)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v473
	F_pfree(m, v447)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L4
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v536
	goto L4
L149:
	;
	v536 = v528 - v447
	goto L148
L150:
	;
	v507 = v503
	goto L159
L151:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	if v487 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v536 = int32(0)
	goto L148
L153:
	;
	goto L154
L154:
	;
	v492 = v447
	goto L155
L155:
	;
	v496 = v492 + int32(1)
	if v496&int32(3) == int32(0) {
		v503 = v496
		goto L150
	} else {
		goto L157
	}
L156:
	;
	v528 = v496
	goto L149
L157:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	if v501 != 0 {
		v492 = v496
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	v516 = int32(-2139062144)
	if (int32(16843008)-v513|v513)&v516 == v516 {
		v507 = v507 + int32(4)
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v522 = v507
	goto L162
L161:
	;
	goto L160
L162:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v526 != 0 {
		v522 = v522 + int32(1)
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v528 = v522
	goto L149
L164:
	;
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v544
	if v544&int32(3) == int32(0) {
		v570 = v544
		goto L168
	} else {
		goto L169
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v603
	goto L4
L167:
	;
	v603 = v595 - v544
	goto L166
L168:
	;
	v574 = v570
	goto L177
L169:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	if v554 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v603 = int32(0)
	goto L166
L171:
	;
	goto L172
L172:
	;
	v559 = v544
	goto L173
L173:
	;
	v563 = v559 + int32(1)
	if v563&int32(3) == int32(0) {
		v570 = v563
		goto L168
	} else {
		goto L175
	}
L174:
	;
	v595 = v563
	goto L167
L175:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	if v568 != 0 {
		v559 = v563
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v583 = int32(-2139062144)
	if (int32(16843008)-v580|v580)&v583 == v583 {
		v574 = v574 + int32(4)
		goto L177
	} else {
		goto L179
	}
L178:
	;
	v589 = v574
	goto L180
L179:
	;
	goto L178
L180:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v593 != 0 {
		v589 = v589 + int32(1)
		goto L180
	} else {
		goto L182
	}
L181:
	;
	v595 = v589
	goto L167
L182:
	;
	goto L181
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v610
	if v610&int32(3) == int32(0) {
		v636 = v610
		goto L186
	} else {
		goto L187
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v669
	goto L4
L185:
	;
	v669 = v661 - v610
	goto L184
L186:
	;
	v640 = v636
	goto L195
L187:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	if v620 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v669 = int32(0)
	goto L184
L189:
	;
	goto L190
L190:
	;
	v625 = v610
	goto L191
L191:
	;
	v629 = v625 + int32(1)
	if v629&int32(3) == int32(0) {
		v636 = v629
		goto L186
	} else {
		goto L193
	}
L192:
	;
	v661 = v629
	goto L185
L193:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	if v634 != 0 {
		v625 = v629
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	v649 = int32(-2139062144)
	if (int32(16843008)-v646|v646)&v649 == v649 {
		v640 = v640 + int32(4)
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v655 = v640
	goto L198
L197:
	;
	goto L196
L198:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	if v659 != 0 {
		v655 = v655 + int32(1)
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v661 = v655
	goto L185
L200:
	;
	goto L199
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v676
	if v676&int32(3) == int32(0) {
		v702 = v676
		goto L204
	} else {
		goto L205
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v735
	goto L4
L203:
	;
	v735 = v727 - v676
	goto L202
L204:
	;
	v706 = v702
	goto L213
L205:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	if v686 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v735 = int32(0)
	goto L202
L207:
	;
	goto L208
L208:
	;
	v691 = v676
	goto L209
L209:
	;
	v695 = v691 + int32(1)
	if v695&int32(3) == int32(0) {
		v702 = v695
		goto L204
	} else {
		goto L211
	}
L210:
	;
	v727 = v695
	goto L203
L211:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
	if v700 != 0 {
		v691 = v695
		goto L209
	} else {
		goto L212
	}
L212:
	;
	goto L210
L213:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	v715 = int32(-2139062144)
	if (int32(16843008)-v712|v712)&v715 == v715 {
		v706 = v706 + int32(4)
		goto L213
	} else {
		goto L215
	}
L214:
	;
	v721 = v706
	goto L216
L215:
	;
	goto L214
L216:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721))))
	if v725 != 0 {
		v721 = v721 + int32(1)
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v727 = v721
	goto L203
L218:
	;
	goto L217
L219:
	;
	v740 = v738
	goto L11
L220:
	;
	F_makeJsonLexContext(m, v17+int32(72), v744, int32(1))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v751 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17-int32(-64)))) = v751
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = int32(1325)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(1326)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(1327)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(1328)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(1329)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(1330)
	v773 = F_pg_parse_json_or_errsave(m, v17+int32(72), v17+int32(32), v751)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_freeJsonLexContext(m, v17+int32(72))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v945 = int32(1)
	goto L8
L224:
	;
	v783 = F_JsonbIteratorInit(m, v779+int32(4))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v783
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v788 = v786 & int32(268435456)
	if v788 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v945 = base.B2i32(v788 == int32(0))
	goto L8
L227:
	;
	goto L230
L228:
	;
	goto L229
L229:
	;
	v830 = F_JsonbIteratorNext(m, v17+int32(32), v17+int32(72), int32(1))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L238
	}
L230:
	;
	v805 = int32(0)
	v811 = F_JsonbIteratorNext(m, v17+int32(32), v17+int32(72), v805)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v813 = int32(4)
	if base.Ui32(v813) <= base.Ui32(v811-v813) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v811 == int32(0) {
		goto L226
	} else {
		goto L236
	}
L234:
	;
	v821 = v805
	goto L235
L235:
	;
	v822 = F_pushJsonbValue(m, l2, v811, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L237
	}
L236:
	;
	v821 = v17 + int32(72)
	goto L235
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v822
	goto L230
L238:
	;
	v837 = F_JsonbIteratorNext(m, v17+int32(32), v17+int32(72), int32(1))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	goto L226
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	if v856&int32(3) == int32(0) {
		v883 = v856
		goto L243
	} else {
		goto L244
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v916
	if base.Ui32(v916) < base.Ui32(int32(268435456)) {
		goto L258
	} else {
		goto L259
	}
L242:
	;
	v916 = v908 - v856
	goto L241
L243:
	;
	v887 = v883
	goto L252
L244:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	if v867 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v916 = int32(0)
	goto L241
L246:
	;
	goto L247
L247:
	;
	v872 = v856
	goto L248
L248:
	;
	v876 = v872 + int32(1)
	if v876&int32(3) == int32(0) {
		v883 = v876
		goto L243
	} else {
		goto L250
	}
L249:
	;
	v908 = v876
	goto L242
L250:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876))))
	if v881 != 0 {
		v872 = v876
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	v896 = int32(-2139062144)
	if (int32(16843008)-v893|v893)&v896 == v896 {
		v887 = v887 + int32(4)
		goto L252
	} else {
		goto L254
	}
L253:
	;
	v902 = v887
	goto L255
L254:
	;
	goto L253
L255:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v906 != 0 {
		v902 = v902 + int32(1)
		goto L255
	} else {
		goto L257
	}
L256:
	;
	v908 = v902
	goto L242
L257:
	;
	goto L256
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v856
	v945 = int32(1)
	goto L8
L259:
	;
	v921 = F_errsave_start(m, int32(0))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	if v921 == int32(0) {
		goto L258
	} else {
		goto L261
	}
L261:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(314073), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(268435455)
	F_errdetail(m, int32(557190), v17)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	F_errsave_finish(m, int32(0), int32(476938), int32(284), int32(268578))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	goto L258
L266:
	;
	if v945 != 0 {
		goto L3
	} else {
		goto L267
	}
L267:
	;
	goto L4
L268:
	;
	v979 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+44)) = uint8(v979)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = int64(4294967312)
	v986 = F_pushJsonbValue(m, l2, int32(4), v17+int32(32))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	switch v1000 - int32(16) {
	case 0:
		goto L274
	case 1:
		goto L276
	default:
		goto L275
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v986
	v992 = F_pushJsonbValue(m, l2, int32(3), v17+int32(72))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v992
	v997 = F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v997
	goto L3
L274:
	;
	v1027 = F_pushJsonbValue(m, l2, int32(3), v17+int32(72))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L284
	}
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L1
	} else {
		goto L281
	}
L276:
	;
	if l5 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1005 = int32(1)
	goto L279
L278:
	;
	v1005 = int32(2)
	goto L279
L279:
	;
	v1008 = F_pushJsonbValue(m, l2, v1005, v17+int32(72))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1008
	goto L3
L281:
	;
	F_errmsg_internal(m, int32(345490), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(476938), int32(851), int32(297103))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1027
	goto L3
}
