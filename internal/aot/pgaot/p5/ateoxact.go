package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_GUC(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 float64
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 float64
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1093 int32
	_ = v1093
	var v1108 int32
	_ = v1108
	var v1123 int32
	_ = v1123
	var v1139 int32
	_ = v1139
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[0]))
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = l0
	v34 = l1
	v46 = v31
	v47 = int32(_a_F_AtEOXact_GUC_0)
	goto L4
L2:
	;
	v1139 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[1])) = v1139 - int32(1)
	return
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v64 = v46 - int32(16)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v65 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v1139 = v34
	goto L3
L6:
	;
	if v62 != 0 {
		v46 = v62
		v47 = v1123
		goto L4
	} else {
		goto L290
	}
L7:
	;
	v1123 = v46
	goto L6
L8:
	;
	goto L9
L9:
	;
	v69 = v46 - int32(72)
	v70 = int32(4)
	v71 = v46 + v70
	v73 = v46 - int32(44)
	v81 = v46 - v70
	v83 = v46 - int32(8)
	v85 = v46 - int32(40)
	v87 = v46 - int32(12)
	v89 = v46 - int32(48)
	v95 = v65
	v104 = v46
	goto L10
L10:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v119 < v34 {
		v1123 = v104
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v1123 = v1093
	goto L6
L12:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v33 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v1108 != 0 {
		v95 = v1108
		v104 = v1093
		goto L10
	} else {
		goto L289
	}
L14:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v95)+40))
	v858 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+40)) = v858
	if v857 == v858 {
		goto L225
	} else {
		goto L226
	}
L15:
	;
	F_pfree(m, v799)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L36
	} else {
		goto L224
	}
L16:
	;
	F_discard_stack_value(m, v69, v95+int32(32))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L36
	} else {
		goto L218
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v121
	F_pfree(m, v95)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L36
	} else {
		goto L217
	}
L18:
	;
	F_discard_stack_value(m, v69, v95+int32(32))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L36
	} else {
		goto L216
	}
L19:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v206+v95)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205+v95)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v207)))
	v217 = base.I32_wrap_i64(base.I64_reinterpret_f64(v215))
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v219 {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	case 4:
		goto L48
	default:
		v834 = v218
		goto L14
	}
L20:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v205 = int32(16)
	v206 = int32(24)
	v207 = v95 + int32(32)
	v208 = v200
	v209 = v95 + int32(40)
	goto L19
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	if v124 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if v119 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v205 = int32(20)
	v206 = int32(28)
	v207 = v95 + int32(48)
	v208 = int32(13)
	v209 = v95 + int32(56)
	goto L19
L24:
	;
	switch v124 - int32(1) {
	case 0:
		goto L16
	default:
		goto L20
	case 2:
		goto L23
	}
L25:
	;
	goto L26
L26:
	;
	if v121 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v119 - int32(1)
	v1093 = v104
	goto L13
L28:
	;
	goto L29
L29:
	;
	v137 = v119 - int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v138 < v137 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v137
	v1093 = v104
	goto L13
L31:
	;
	goto L32
L32:
	;
	switch v124 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	case 2:
		goto L33
	default:
		goto L17
	}
L33:
	;
	F_discard_stack_value(m, v69, v95+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L36
	} else {
		goto L43
	}
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v156 != int32(1) {
		goto L18
	} else {
		goto L42
	}
L35:
	;
	F_discard_stack_value(m, v69, v95+int32(32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	return
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v147 == int32(3) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_discard_stack_value(m, v69, v121+int32(48))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = int32(1)
	goto L17
L41:
	;
	goto L40
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+28)) = v161
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v95)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+48)) = v163
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v95)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+56)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = int32(3)
	goto L17
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+28)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v177 == int32(3) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_discard_stack_value(m, v69, v121+int32(48))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L36
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v95)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+56)) = v184
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+48)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = int32(3)
	goto L17
L47:
	;
	goto L46
L48:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	if v217 == v714 {
		goto L191
	} else {
		goto L192
	}
L49:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	if v418 != v217 {
		goto L129
	} else {
		goto L130
	}
L50:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v352)))
	if base.F64_eq(v215, v353) != 0 {
		goto L103
	} else {
		goto L104
	}
L51:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v217 == v288 {
		goto L78
	} else {
		goto L79
	}
L52:
	;
	v221 = v217 & int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v221 == v223 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v225 == v214 {
		v834 = v218
		goto L14
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	if v227 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	m.T0[v227].(func(*base.Module, int32, int32))(m, v221, v214)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L36
	} else {
		goto L60
	}
L58:
	;
	v231 = v222
	goto L59
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v221)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v214
	v235 = int32(1)
	if base.B2i32(v233 == int32(0))|base.B2i32(v233 == v214) != 0 {
		v834 = v235
		goto L14
	} else {
		goto L61
	}
L60:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v231 = v230
	goto L59
L61:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v240 {
	case 0:
		goto L67
	case 1:
		goto L66
	case 2:
		goto L65
	case 3:
		goto L64
	case 4:
		goto L63
	default:
		goto L62
	}
L62:
	;
	v254 = v64
	goto L73
L63:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v233 == v249 {
		v834 = v235
		goto L14
	} else {
		goto L72
	}
L64:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v233 != v247 {
		goto L62
	} else {
		goto L71
	}
L65:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v233 != v245 {
		goto L62
	} else {
		goto L70
	}
L66:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v233 != v243 {
		goto L62
	} else {
		goto L69
	}
L67:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v233 != v241 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v834 = v235
	goto L14
L69:
	;
	v834 = v235
	goto L14
L70:
	;
	v834 = v235
	goto L14
L71:
	;
	v834 = v235
	goto L14
L72:
	;
	goto L62
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v280 == int32(0) {
		v799 = v233
		v803 = v235
		goto L15
	} else {
		goto L75
	}
L74:
	;
	v834 = v235
	goto L14
L75:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v280)+40))
	if v233 == v283 {
		v834 = v235
		goto L14
	} else {
		goto L76
	}
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v280)+56))
	if v233 != v285 {
		v254 = v280
		goto L73
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v290 == v214 {
		v834 = v218
		goto L14
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v292 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	m.T0[v292].(func(*base.Module, int32, int32))(m, v217, v214)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L36
	} else {
		goto L85
	}
L83:
	;
	v296 = v287
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v217
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v214
	v300 = int32(1)
	if base.B2i32(v298 == int32(0))|base.B2i32(v298 == v214) != 0 {
		v834 = v300
		goto L14
	} else {
		goto L86
	}
L85:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v296 = v295
	goto L84
L86:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v305 {
	case 0:
		goto L92
	case 1:
		goto L91
	case 2:
		goto L90
	case 3:
		goto L89
	case 4:
		goto L88
	default:
		goto L87
	}
L87:
	;
	v319 = v64
	goto L98
L88:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v298 == v314 {
		v834 = v300
		goto L14
	} else {
		goto L97
	}
L89:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v298 != v312 {
		goto L87
	} else {
		goto L96
	}
L90:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v298 != v310 {
		goto L87
	} else {
		goto L95
	}
L91:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v298 != v308 {
		goto L87
	} else {
		goto L94
	}
L92:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v298 != v306 {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v834 = v300
	goto L14
L94:
	;
	v834 = v300
	goto L14
L95:
	;
	v834 = v300
	goto L14
L96:
	;
	v834 = v300
	goto L14
L97:
	;
	goto L87
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	if v345 == int32(0) {
		v799 = v298
		v803 = v300
		goto L15
	} else {
		goto L100
	}
L99:
	;
	v834 = v300
	goto L14
L100:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v345)+40))
	if v298 == v348 {
		v834 = v300
		goto L14
	} else {
		goto L101
	}
L101:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v345)+56))
	if v298 != v350 {
		v319 = v345
		goto L98
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v355 == v214 {
		v834 = v218
		goto L14
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v357 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	m.T0[v357].(func(*base.Module, float64, int32))(m, v215, v214)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L36
	} else {
		goto L110
	}
L108:
	;
	v361 = v352
	goto L109
L109:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v361))) = v215
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v214
	v365 = int32(1)
	if base.B2i32(v363 == int32(0))|base.B2i32(v363 == v214) != 0 {
		v834 = v365
		goto L14
	} else {
		goto L111
	}
L110:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v361 = v360
	goto L109
L111:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v370 {
	case 0:
		goto L117
	case 1:
		goto L116
	case 2:
		goto L115
	case 3:
		goto L114
	case 4:
		goto L113
	default:
		goto L112
	}
L112:
	;
	v384 = v64
	goto L123
L113:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v363 == v379 {
		v834 = v365
		goto L14
	} else {
		goto L122
	}
L114:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v363 != v377 {
		goto L112
	} else {
		goto L121
	}
L115:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v363 != v375 {
		goto L112
	} else {
		goto L120
	}
L116:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v363 != v373 {
		goto L112
	} else {
		goto L119
	}
L117:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v363 != v371 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	v834 = v365
	goto L14
L119:
	;
	v834 = v365
	goto L14
L120:
	;
	v834 = v365
	goto L14
L121:
	;
	v834 = v365
	goto L14
L122:
	;
	goto L112
L123:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v410 == int32(0) {
		v799 = v363
		v803 = v365
		goto L15
	} else {
		goto L125
	}
L124:
	;
	v834 = v365
	goto L14
L125:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v410)+40))
	if v363 == v413 {
		v834 = v365
		goto L14
	} else {
		goto L126
	}
L126:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v410)+56))
	if v363 != v415 {
		v384 = v410
		goto L123
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
	v589 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+32)) = v589
	if v588 == v589 {
		goto L169
	} else {
		goto L170
	}
L129:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v46)+32))
	if v423 != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v420 != v214 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v565 = int32(0)
	goto L128
L132:
	;
	m.T0[v423].(func(*base.Module, int32, int32))(m, v217, v214)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L36
	} else {
		goto L135
	}
L133:
	;
	v428 = v418
	v429 = v417
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = v217
	if v428 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v428 = v427
	v429 = v426
	goto L134
L136:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v214
	v507 = int32(1)
	if base.B2i32(v505 == int32(0))|base.B2i32(v214 == v505) != 0 {
		v565 = v507
		goto L128
	} else {
		goto L149
	}
L137:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	if v428 == v434 {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v428 == v436 {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v428 == v438 {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v442 = v64
	goto L141
L141:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	if v469 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	F_pfree(m, v428)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L36
	} else {
		goto L148
	}
L143:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)+32))
	if v428 == v470 {
		goto L136
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	goto L142
L146:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v469)+48))
	if v428 != v472 {
		v442 = v469
		goto L141
	} else {
		goto L147
	}
L147:
	;
	goto L136
L148:
	;
	goto L136
L149:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v512 {
	case 0:
		goto L155
	case 1:
		goto L154
	case 2:
		goto L153
	case 3:
		goto L152
	case 4:
		goto L151
	default:
		goto L150
	}
L150:
	;
	v525 = v64
	goto L161
L151:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v505 == v521 {
		v565 = v507
		goto L128
	} else {
		goto L160
	}
L152:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v505 != v519 {
		goto L150
	} else {
		goto L159
	}
L153:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v505 != v517 {
		goto L150
	} else {
		goto L158
	}
L154:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v505 != v515 {
		goto L150
	} else {
		goto L157
	}
L155:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v505 != v513 {
		goto L150
	} else {
		goto L156
	}
L156:
	;
	v565 = v507
	goto L128
L157:
	;
	v565 = v507
	goto L128
L158:
	;
	v565 = v507
	goto L128
L159:
	;
	v565 = v507
	goto L128
L160:
	;
	goto L150
L161:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	if v552 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	F_pfree(m, v505)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L36
	} else {
		goto L168
	}
L163:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+40))
	if v505 == v553 {
		v565 = v507
		goto L128
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	goto L162
L166:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v552)+56))
	if v505 != v555 {
		v525 = v552
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v565 = v507
	goto L128
L168:
	;
	v565 = v507
	goto L128
L169:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	v666 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = v666
	if v665 == v666 {
		v834 = v565
		goto L14
	} else {
		goto L182
	}
L170:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
	if v588 == v594 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v588 == v596 {
		goto L169
	} else {
		goto L172
	}
L172:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v588 == v598 {
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v602 = v64
	goto L174
L174:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	if v629 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	F_pfree(m, v588)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L36
	} else {
		goto L181
	}
L176:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+32))
	if v588 == v630 {
		goto L169
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	goto L175
L179:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v629)+48))
	if v588 != v632 {
		v602 = v629
		goto L174
	} else {
		goto L180
	}
L180:
	;
	goto L169
L181:
	;
	goto L169
L182:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	if v665 == v671 {
		v834 = v565
		goto L14
	} else {
		goto L183
	}
L183:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v665 == v673 {
		v834 = v565
		goto L14
	} else {
		goto L184
	}
L184:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v665 == v675 {
		v834 = v565
		goto L14
	} else {
		goto L185
	}
L185:
	;
	v680 = v64
	goto L186
L186:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	if v706 == int32(0) {
		v799 = v665
		v803 = v565
		goto L15
	} else {
		goto L188
	}
L187:
	;
	v834 = v565
	goto L14
L188:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v706)+32))
	if v665 == v709 {
		v834 = v565
		goto L14
	} else {
		goto L189
	}
L189:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v706)+48))
	if v665 != v711 {
		v680 = v706
		goto L186
	} else {
		goto L190
	}
L190:
	;
	goto L187
L191:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v716 == v214 {
		v834 = v218
		goto L14
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v46)+36))
	if v718 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L193
L195:
	;
	m.T0[v718].(func(*base.Module, int32, int32))(m, v217, v214)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L36
	} else {
		goto L198
	}
L196:
	;
	v722 = v713
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722))) = v217
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v214
	v726 = int32(1)
	if base.B2i32(v724 == int32(0))|base.B2i32(v724 == v214) != 0 {
		v834 = v726
		goto L14
	} else {
		goto L199
	}
L198:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v722 = v721
	goto L197
L199:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v731 {
	case 0:
		goto L205
	case 1:
		goto L204
	case 2:
		goto L203
	case 3:
		goto L202
	case 4:
		goto L201
	default:
		goto L200
	}
L200:
	;
	v745 = v64
	goto L211
L201:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v724 == v740 {
		v834 = v726
		goto L14
	} else {
		goto L210
	}
L202:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v724 != v738 {
		goto L200
	} else {
		goto L209
	}
L203:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v724 != v736 {
		goto L200
	} else {
		goto L208
	}
L204:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v724 != v734 {
		goto L200
	} else {
		goto L207
	}
L205:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v724 != v732 {
		goto L200
	} else {
		goto L206
	}
L206:
	;
	v834 = v726
	goto L14
L207:
	;
	v834 = v726
	goto L14
L208:
	;
	v834 = v726
	goto L14
L209:
	;
	v834 = v726
	goto L14
L210:
	;
	goto L200
L211:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	if v771 == int32(0) {
		v799 = v724
		v803 = v726
		goto L15
	} else {
		goto L213
	}
L212:
	;
	v834 = v726
	goto L14
L213:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v771)+40))
	if v724 == v774 {
		v834 = v726
		goto L14
	} else {
		goto L214
	}
L214:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v771)+56))
	if v724 != v776 {
		v745 = v771
		goto L211
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	goto L17
L217:
	;
	v1093 = v104
	goto L13
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v121
	if v121 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v62
	F_pfree(m, v95)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L36
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	F_pfree(m, v95)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L36
	} else {
		goto L223
	}
L222:
	;
	v1093 = v47
	goto L13
L223:
	;
	v1093 = v104
	goto L13
L224:
	;
	v834 = v803
	goto L14
L225:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v95)+56))
	v941 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+56)) = v941
	if v940 == v941 {
		goto L247
	} else {
		goto L248
	}
L226:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v857 == v862 {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v864 {
	case 0:
		goto L233
	case 1:
		goto L232
	case 2:
		goto L231
	case 3:
		goto L230
	case 4:
		goto L229
	default:
		goto L228
	}
L228:
	;
	v877 = v64
	goto L239
L229:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v857 == v873 {
		goto L225
	} else {
		goto L238
	}
L230:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v857 != v871 {
		goto L228
	} else {
		goto L237
	}
L231:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v857 != v869 {
		goto L228
	} else {
		goto L236
	}
L232:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v857 != v867 {
		goto L228
	} else {
		goto L235
	}
L233:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v857 != v865 {
		goto L228
	} else {
		goto L234
	}
L234:
	;
	goto L225
L235:
	;
	goto L225
L236:
	;
	goto L225
L237:
	;
	goto L225
L238:
	;
	goto L228
L239:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	if v904 != 0 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	F_pfree(m, v857)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L36
	} else {
		goto L246
	}
L241:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v904)+40))
	if v857 == v905 {
		goto L225
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	goto L240
L244:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v904)+56))
	if v857 != v907 {
		v877 = v904
		goto L239
	} else {
		goto L245
	}
L245:
	;
	goto L225
L246:
	;
	goto L225
L247:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v1023 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L248:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v940 == v945 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v947 {
	case 0:
		goto L255
	case 1:
		goto L254
	case 2:
		goto L253
	case 3:
		goto L252
	case 4:
		goto L251
	default:
		goto L250
	}
L250:
	;
	v960 = v64
	goto L261
L251:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	if v940 == v956 {
		goto L247
	} else {
		goto L260
	}
L252:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v940 != v954 {
		goto L250
	} else {
		goto L259
	}
L253:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
	if v940 != v952 {
		goto L250
	} else {
		goto L258
	}
L254:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	if v940 != v950 {
		goto L250
	} else {
		goto L257
	}
L255:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v940 != v948 {
		goto L250
	} else {
		goto L256
	}
L256:
	;
	goto L247
L257:
	;
	goto L247
L258:
	;
	goto L247
L259:
	;
	goto L247
L260:
	;
	goto L250
L261:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	if v987 != 0 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	F_pfree(m, v940)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L36
	} else {
		goto L268
	}
L263:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v987)+40))
	if v940 == v988 {
		goto L247
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	goto L262
L266:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v987)+56))
	if v940 != v990 {
		v960 = v987
		goto L261
	} else {
		goto L267
	}
L267:
	;
	goto L247
L268:
	;
	goto L247
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v46-int32(32)))) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v46-int32(24)))) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v121
	if v121 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L270:
	;
	if v208 == int32(0) {
		goto L269
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	if v208 != 0 {
		goto L269
	} else {
		goto L278
	}
L273:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[2]))
	if v1029 != 0 {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	v1037 = int32(_a_F_AtEOXact_GUC_1)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v1036)+4)) = v83
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[3])) = v83
	goto L269
L275:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[3]))
	v1036 = v1031
	goto L274
L276:
	;
	goto L277
L277:
	;
	v1033 = int32(_a_F_AtEOXact_GUC_1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[2])) = v1033
	v1036 = v1033
	goto L274
L278:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+4)) = v1044
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v1044))) = v1046
	goto L269
L279:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46-int32(52)))))
	if v1063&int32(64) == int32(0) {
		v1093 = v1062
		goto L13
	} else {
		goto L287
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v62
	F_pfree(m, v95)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L36
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	F_pfree(m, v95)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L36
	} else {
		goto L285
	}
L283:
	;
	if v834 != 0 {
		v1062 = v47
		goto L279
	} else {
		goto L284
	}
L284:
	;
	v1093 = v47
	goto L13
L285:
	;
	if v834 == int32(0) {
		v1093 = v104
		goto L13
	} else {
		goto L286
	}
L286:
	;
	v1062 = v104
	goto L279
L287:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v1068&int32(4) != 0 {
		v1093 = v1062
		goto L13
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v1068 | int32(4)
	v1074 = int32(_a_F_AtEOXact_GUC_2)
	v1075 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v1075
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_GUC[4])) = v71
	v1093 = v1062
	goto L13
L289:
	;
	goto L11
L290:
	;
	goto L5
}
func F_AtEOXact_MultiXact(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[0]))
	v3 = int32(_a_F_AtEOXact_MultiXact_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[1]))
	v5 = int32(2)
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2+v4<<(uint(v5)%32)))) = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[2]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11+v13<<(uint(v5)%32)))) = v8
	v20 = int32(_a_F_AtEOXact_MultiXact_1)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[3])) = v20
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[4])) = v20
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[5])) = v8
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_MultiXact[6])) = v8
	return
}
