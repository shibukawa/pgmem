package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_func_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	if v11 < v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = v8
	goto L6
L5:
	;
	v14 = v11
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v3
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 == int32(468) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_scanner_yyerror(m, int32(212159), l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v29 = v19 + int32(1)
	if v14 != v29 {
		v19 = v29
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_select_candidate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v592 int32
	_ = v592
	var v608 int32
	_ = v608
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var __phi635 int32
	_ = __phi635
	var v642 int32
	_ = v642
	var __phi642 int32
	_ = __phi642
	var v643 int32
	_ = v643
	var __phi643 int32
	_ = __phi643
	var v644 int32
	_ = v644
	var __phi644 int32
	_ = __phi644
	var v656 int32
	_ = v656
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v826 int32
	_ = v826
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v890 int32
	_ = v890
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v950 int32
	_ = v950
	var v974 int32
	_ = v974
	var v992 int32
	_ = v992
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(640)
	m.G0 = v19
	if l0 <= int32(100) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L13
	} else {
		goto L181
	}
L4:
	;
	v28 = v4
	v38 = v4
	goto L7
L5:
	;
	v76 = v4
	goto L6
L6:
	;
	if l2 != 0 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v42 = v28 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1+v42)))
	if v44 != int32(705) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v76 = v54
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v42))) = v55
	v61 = v28 + int32(1)
	if v61 != l0 {
		v28 = v61
		v38 = v54
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v47 = F_getBaseType(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v54 = v38 + int32(1)
	v55 = int32(705)
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	v54 = v38
	v55 = v47
	goto L9
L15:
	;
	goto L8
L16:
	;
	m.G0 = v19 + int32(640)
	return v992
L17:
	;
	v85 = l2
	v89 = l2
	v94 = v4
	v95 = v4
	v97 = v4
	goto L20
L18:
	;
	v226 = v4
	goto L19
L19:
	;
	v236 = int32(0)
	if v236 < l0 {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v99 = int32(0)
	if l0 <= v99 {
		v191 = v99
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = int32(0)
	if v214 == int32(1) {
		v992 = v211
		goto L16
	} else {
		goto L44
	}
L22:
	;
	if v94 < v191 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	v103 = v85 + int32(32)
	v104 = int32(0)
	if l0 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v111 = v104
	v112 = v99
	v115 = v104
	goto L27
L25:
	;
	v160 = v104
	v161 = v99
	goto L26
L26:
	;
	if l0&int32(1) == int32(0) {
		v191 = v161
		goto L22
	} else {
		goto L36
	}
L27:
	;
	v125 = v111 << (uint(int32(2)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v19+int32(240)))))
	if v129 != int32(705) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v160 = v153
	v161 = v151
	goto L26
L29:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v125+v103)))
	v136 = v112 + base.B2i32(v133 == v129)
	goto L31
L30:
	;
	v136 = v112
	goto L31
L31:
	;
	v140 = (v111 | int32(1)) << (uint(int32(2)) % 32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+(v19+int32(240)))))
	if v144 != int32(705) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v140+v103)))
	v151 = v136 + base.B2i32(v148 == v144)
	goto L34
L33:
	;
	v151 = v136
	goto L34
L34:
	;
	v152 = int32(2)
	v153 = v111 + v152
	v155 = v115 + v152
	if v155 != l0&int32(2147483646) {
		v111 = v153
		v112 = v151
		v115 = v155
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L28
L36:
	;
	v176 = v160 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+(v19+int32(240)))))
	if v180 == int32(705) {
		v191 = v161
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176+v103)))
	v191 = v161 + base.B2i32(v184 == v180)
	goto L22
L38:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v215 != 0 {
		v85 = v215
		v89 = v211
		v94 = v212
		v95 = v213
		v97 = v214
		goto L20
	} else {
		goto L43
	}
L39:
	;
	v211 = v85
	v212 = v191
	v213 = v85
	v214 = int32(1)
	goto L38
L40:
	;
	if v95 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v191 != v94 {
		v211 = v89
		v212 = v94
		v213 = v95
		v214 = v97
		goto L38
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v85
	v211 = v89
	v212 = v94
	v213 = v85
	v214 = v97 + int32(1)
	goto L38
L43:
	;
	goto L21
L44:
	;
	v226 = v211
	goto L19
L45:
	;
	v242 = v236
	goto L48
L46:
	;
	goto L47
L47:
	;
	v286 = int32(0)
	if v226 == v286 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v242<<(uint(int32(2))%32))))
	v264 = F_TypeCategory(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L13
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(128)+v242))) = uint8(v264)
	v268 = v242 + int32(1)
	if v268 != l0 {
		v242 = v268
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v407 = int32(0)
	if v76 == v407 {
		v992 = v407
		goto L16
	} else {
		goto L78
	}
L53:
	;
	v402 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v290 = int32(0)
	v296 = v290
	v300 = v226
	v302 = v290
	v303 = v286
	v305 = v226
	goto L56
L56:
	;
	v310 = int32(0)
	if base.B2i32(l0 <= v290) == v310 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = int32(0)
	if v382 == int32(1) {
		v992 = v385
		goto L16
	} else {
		goto L77
	}
L58:
	;
	v317 = v310
	v319 = int32(0)
	goto L61
L59:
	;
	v359 = v310
	goto L60
L60:
	;
	if v303 < v359 {
		goto L72
	} else {
		goto L73
	}
L61:
	;
	v333 = v319 << (uint(int32(2)) % 32)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333+(v19+int32(240)))))
	if v337 == int32(705) {
		v353 = v317
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v359 = v353
	goto L60
L63:
	;
	v356 = v319 + int32(1)
	if v356 != l0 {
		v317 = v353
		v319 = v356
		goto L61
	} else {
		goto L70
	}
L64:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v333+(v300+int32(32)))))
	if v337 != v341 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19+int32(128)+v319))))
	v347 = F_IsPreferredType(m, v346, v341)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L13
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v353 = v317 + int32(1)
	goto L63
L68:
	;
	if v347 == int32(0) {
		v353 = v317
		goto L63
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L62
L71:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	if v386 != 0 {
		v296 = v382
		v300 = v386
		v302 = v383
		v303 = v384
		v305 = v385
		goto L56
	} else {
		goto L76
	}
L72:
	;
	v382 = int32(1)
	v383 = v300
	v384 = v359
	v385 = v300
	goto L71
L73:
	;
	if v302 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	if v359 != v303 {
		v382 = v296
		v383 = v302
		v384 = v303
		v385 = v305
		goto L71
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v300
	v382 = v296 + int32(1)
	v383 = v300
	v384 = v303
	v385 = v305
	goto L71
L76:
	;
	goto L57
L77:
	;
	v402 = v385
	goto L52
L78:
	;
	if l0 <= int32(0) {
		v992 = v407
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v412 = int32(0)
	v416 = v412
	v418 = v412
	goto L81
L80:
	;
	v735 = int32(0)
	if l0 <= v76 {
		v992 = v735
		goto L16
	} else {
		goto L138
	}
L81:
	;
	v431 = v416 << (uint(int32(2)) % 32)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431+(v19+int32(240)))))
	if v435 != int32(705) {
		v608 = v418
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v623 = int32(0)
	v625 = v608 & int32(1)
	if v625&base.B2i32(v402 != v623) == v623 {
		goto L111
	} else {
		goto L112
	}
L83:
	;
	v621 = v416 + int32(1)
	if v621 != l0 {
		v416 = v621
		v418 = v608
		goto L81
	} else {
		goto L110
	}
L84:
	;
	v438 = int32(0)
	v441 = v19 + int32(16) + v416
	*(*uint8)(unsafe.Add(mBase, uint32(v441))) = uint8(v438)
	v446 = v19 + int32(128) + v416
	*(*uint8)(unsafe.Add(mBase, uint32(v446))) = uint8(v438)
	if v402 == v438 {
		v608 = int32(1)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v457 = v402
	v463 = v438
	v464 = v438
	v466 = v438
	goto L87
L86:
	;
	if v592&int32(255) == int32(83) {
		v608 = int32(1)
		goto L83
	} else {
		goto L109
	}
L87:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v431+v457)+32))
	F_get_type_category_preferred(m, v471, v19+int32(127), v19+int32(126))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L13
	} else {
		goto L89
	}
L88:
	;
	v578 = int32(1)
	if v571&v578 == int32(0) {
		v608 = v578
		goto L83
	} else {
		goto L108
	}
L89:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+127)))
	v480 = v463 & int32(255)
	if v480 != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v441))) = uint8(v575)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	if v577 != 0 {
		v457 = v577
		v463 = v568
		v464 = v575
		v466 = v571
		goto L87
	} else {
		goto L107
	}
L91:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	v562 = v544
	v568 = v463
	v571 = v553
	v575 = v464 | v557
	goto L90
L92:
	;
	v482 = v478 & int32(255)
	if v482 == v480 {
		v544 = v457
		v553 = v466
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v446))) = uint8(v478)
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	v562 = v457
	v568 = v478
	v571 = v466
	v575 = v540
	goto L90
L95:
	;
	if v482 != int32(83) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v489 = v457
	goto L99
L97:
	;
	v522 = v457
	v531 = v466
	goto L98
L98:
	;
	v535 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v446))) = uint8(v535)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	v562 = v522
	v568 = v535
	v571 = v531
	v575 = v538
	goto L90
L99:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	if v502 == int32(0) {
		v592 = v463
		goto L86
	} else {
		goto L101
	}
L100:
	;
	v522 = v502
	v531 = int32(1)
	goto L98
L101:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v431+v502)+32))
	F_get_type_category_preferred(m, v506, v19+int32(127), v19+int32(126))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+127)))
	if v480 == v513 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v544 = v502
	v553 = int32(1)
	goto L91
L104:
	;
	goto L105
L105:
	;
	if v513 != int32(83) {
		v489 = v502
		goto L99
	} else {
		goto L106
	}
L106:
	;
	goto L100
L107:
	;
	goto L88
L108:
	;
	v592 = v568
	goto L86
L109:
	;
	v727 = v402
	goto L80
L110:
	;
	goto L82
L111:
	;
	if v625 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	__phi635 = v402
	__phi642 = v402
	__phi643 = v623
	__phi644 = int32(0)
	v635 = __phi635
	v642 = __phi642
	v643 = __phi643
	v644 = __phi644
	goto L118
L114:
	;
	v632 = int32(0)
	goto L116
L115:
	;
	v632 = v402
	goto L116
L116:
	;
	v727 = v632
	goto L80
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v714))) = int32(0)
	if v713 == int32(1) {
		v992 = v642
		goto L16
	} else {
		goto L137
	}
L118:
	;
	v656 = int32(0)
	goto L121
L119:
	;
	if v643 == int32(1) {
		v992 = v402
		goto L16
	} else {
		goto L136
	}
L120:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	if v644 != 0 {
		goto L131
	} else {
		goto L132
	}
L121:
	;
	v670 = v656 << (uint(int32(2)) % 32)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670+(v19+int32(240)))))
	if v674 != int32(705) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v704 = v643 + int32(1)
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	if v705 != 0 {
		__phi635 = v705
		__phi643 = v704
		__phi644 = v635
		v635 = __phi635
		v643 = __phi643
		v644 = __phi644
		goto L118
	} else {
		goto L130
	}
L123:
	;
	v701 = v656 + int32(1)
	if v701 != l0 {
		v656 = v701
		goto L121
	} else {
		goto L129
	}
L124:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v670+(v635+int32(32)))))
	F_get_type_category_preferred(m, v678, v19+int32(127), v19+int32(126))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+127)))
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(128)+v656))))
	if v685 != v689 {
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(16)+v656))))
	if v694 != int32(1) {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+126)))
	if v697 != int32(1) {
		goto L120
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	goto L122
L130:
	;
	v713 = v704
	v714 = v635
	goto L117
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v706
	if v706 != 0 {
		__phi635 = v706
		v635 = __phi635
		goto L118
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	if v706 != 0 {
		__phi635 = v706
		__phi642 = v706
		__phi644 = int32(0)
		v635 = __phi635
		v642 = __phi642
		v644 = __phi644
		goto L118
	} else {
		goto L135
	}
L134:
	;
	v713 = v643
	v714 = v644
	goto L117
L135:
	;
	goto L119
L136:
	;
	v727 = v402
	goto L80
L137:
	;
	v727 = v642
	goto L80
L138:
	;
	if l0 <= int32(0) {
		v992 = v735
		goto L16
	} else {
		goto L139
	}
L139:
	;
	v743 = int32(705)
	v746 = v735
	goto L142
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = int32(0)
	v992 = v974
	goto L16
L141:
	;
	v992 = int32(0)
	goto L16
L142:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v746<<(uint(int32(2))%32))))
	if v761 == int32(705) {
		v767 = v743
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v771 = int32(0)
	if v767 == int32(705) {
		v992 = v771
		goto L16
	} else {
		goto L151
	}
L144:
	;
	v769 = v746 + int32(1)
	if v769 != l0 {
		v743 = v767
		v746 = v769
		goto L142
	} else {
		goto L150
	}
L145:
	;
	if v743 == int32(705) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v767 = v761
	goto L144
L147:
	;
	goto L148
L148:
	;
	if v743 != v761 {
		goto L141
	} else {
		goto L149
	}
L149:
	;
	v767 = v743
	goto L144
L150:
	;
	goto L143
L151:
	;
	if l0 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v727 == int32(0) {
		v992 = v771
		goto L16
	} else {
		goto L164
	}
L153:
	;
	v777 = l0 & int32(7)
	v778 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(l0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v792 = v778
	v795 = int32(0)
	goto L157
L155:
	;
	v826 = v778
	goto L156
L156:
	;
	if v777 == int32(0) {
		goto L152
	} else {
		goto L160
	}
L157:
	;
	v805 = v19 + int32(240) + v792<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v805))) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+4)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+8)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+12)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+16)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+20)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+24)) = v767
	*(*int32)(unsafe.Add(mBase, uint32(v805)+28)) = v767
	v814 = int32(8)
	v815 = v792 + v814
	v817 = v795 + v814
	if v817 != l0&int32(2147483640) {
		v792 = v815
		v795 = v817
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v826 = v815
	goto L156
L159:
	;
	goto L158
L160:
	;
	v838 = v778
	v844 = v826
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(240)+v844<<(uint(int32(2))%32)))) = v767
	v859 = int32(1)
	v862 = v838 + v859
	if v862 != v777 {
		v838 = v862
		v844 = v844 + v859
		goto L161
	} else {
		goto L163
	}
L162:
	;
	goto L152
L163:
	;
	goto L162
L164:
	;
	v890 = v727
	goto L167
L165:
	;
	if v903 != 0 {
		v974 = v950
		goto L140
	} else {
		goto L180
	}
L166:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	if v907 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L167:
	;
	v903 = F_can_coerce_type(m, l0, v19+int32(240), v890+int32(32), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L13
	} else {
		goto L169
	}
L168:
	;
	v950 = int32(0)
	goto L165
L169:
	;
	if v903 != 0 {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	if v905 != 0 {
		v890 = v905
		goto L167
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	v974 = v890
	goto L140
L173:
	;
	goto L174
L174:
	;
	v914 = v907
	goto L175
L175:
	;
	v931 = F_can_coerce_type(m, l0, v19+int32(240), v914+int32(32), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L13
	} else {
		goto L177
	}
L176:
	;
	v950 = v890
	goto L165
L177:
	;
	if v931 != 0 {
		v992 = v771
		goto L16
	} else {
		goto L178
	}
L178:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	if v933 != 0 {
		v914 = v933
		goto L175
	} else {
		goto L179
	}
L179:
	;
	goto L176
L180:
	;
	goto L141
L181:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	v1013 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1013
	F_errmsg_plural(m, int32(254072), int32(254120), v1013, v19)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L13
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(500150), int32(1036), int32(356863))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L13
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_volatile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(44718), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499337), int32(1927), int32(385882))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28+v29)+101)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_get_func_variadictype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(44718), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499337), int32(1870), int32(366211))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28+v29)+88))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_makeFuncExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	v8 = F_palloc0(m, int32(36))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l3
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l4
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(15)
		return v8
	}
}
func F_parse_func_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	v7 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v7
	if l1 == v7 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L25
	} else {
		goto L50
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L25
	} else {
		goto L45
	}
L3:
	;
	F_errorConflictingDefElem(m, v42, l0)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
	} else {
		goto L44
	}
L4:
	;
	m.G0 = v13 + int32(32)
	return
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = v7
	goto L7
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v44 = int32(219899)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[479])))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 == int32(0) {
		v67 = v47
		v68 = v48
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L4
L9:
	;
	v138 = v36 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v138 < v139 {
		v36 = v138
		goto L7
	} else {
		goto L43
	}
L10:
	;
	if v68-v67 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L10
L12:
	;
	if v47 != v48 {
		v67 = v47
		v68 = v48
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = v43
	v53 = v44
	goto L14
L14:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v56
		v68 = v57
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v67 = v56
	v68 = v57
	goto L11
L16:
	;
	v60 = int32(1)
	if v56 == v57 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v72 == int32(1) {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v92 = int32(209678)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _consts[480])))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v96 == int32(0) {
		v115 = v95
		v116 = v96
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v75 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v77 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L9
L23:
	;
	goto L24
L24:
	;
	v82 = int32(0)
	v85 = F_LookupFuncName(m, v77, v82, v82, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v87 = F_get_func_rettype(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if v87 != int32(3115) {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v85
	goto L9
L29:
	;
	if v116-v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v95 != v96 {
		v115 = v95
		v116 = v96
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v100 = v43
	v101 = v92
	goto L33
L33:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v104
		v116 = v105
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v115 = v104
	v116 = v105
	goto L30
L35:
	;
	v108 = int32(1)
	if v104 == v105 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v118 == int32(1) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v121)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v123 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(111669150705)
	v130 = F_LookupFuncName(m, v123, int32(2), v13+int32(24), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L25
	} else {
		goto L42
	}
L40:
	;
	v133 = int32(0)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v133
	goto L9
L42:
	;
	v133 = v130
	goto L41
L43:
	;
	goto L8
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v164 = F_NameListToString(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(219322)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v164
	F_errmsg(m, int32(191407), v13)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(494533), int32(501), int32(490027))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v181
	F_errmsg_internal(m, int32(438096), v13+int32(16))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(494533), int32(561), int32(137552))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
