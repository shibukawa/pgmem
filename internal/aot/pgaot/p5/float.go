package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FloatExceptionHandler(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_errcode(m, int32(16908418))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_FloatExceptionHandler_0), int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errdetail(m, int32(_a_F_FloatExceptionHandler_1), int32(0))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_FloatExceptionHandler_2), int32(3188), int32(_a_F_FloatExceptionHandler_3))
					v21 = m.ExcPending
					if v21 != 0 {
						return
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
func F_float_to_shortest_decimal_buf(m *base.Module, l0 float32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int64
	_ = v180
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int64
	_ = v317
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v322 int64
	_ = v322
	var v324 int32
	_ = v324
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v394 int64
	_ = v394
	var v396 int32
	_ = v396
	var v399 int64
	_ = v399
	var v401 int32
	_ = v401
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int64
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	v3 = int32(0)
	v20 = base.I32_reinterpret_f32(l0)
	v22 = v20 & int32(_a_F_float_to_shortest_decimal_buf_0)
	v25 = int32(255)
	v26 = int32(base.Ui32(v20)>>(uint(int32(23))%32)) & v25
	if v26|v22 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v1101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099+l1))) = uint8(v1101)
	return
L2:
	;
	v31 = base.B2i32(v26 != v25)
	goto L4
L3:
	;
	v31 = v3
	goto L4
L4:
	;
	if v31 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(int32(23)) < base.Ui32(v26-int32(127)) {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_float_to_shortest_decimal_buf[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v35)
	v38 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_float_to_shortest_decimal_buf[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v38)
	v1099 = int32(3)
	goto L1
L9:
	;
	goto L10
L10:
	;
	if v20 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v43)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v47 = l1 + int32(base.Ui32(v20)>>(uint(int32(31))%32))
	if v26 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = int64(8751735898823355977)
	if v20 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v55 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v55)
	if v20 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v54 = int32(9)
	goto L19
L18:
	;
	v54 = int32(8)
	goto L19
L19:
	;
	v1099 = v54
	goto L1
L20:
	;
	v61 = int32(2)
	goto L22
L21:
	;
	v61 = int32(1)
	goto L22
L22:
	;
	v1099 = v61
	goto L1
L23:
	;
	v709 = int32(0)
	if v20 < v709 {
		goto L91
	} else {
		goto L92
	}
L24:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_buf_1)) < base.Ui32(v650) {
		v692 = v650
		v697 = v655
		v708 = int32(8)
		goto L23
	} else {
		goto L82
	}
L25:
	;
	v78 = v22 << (uint(int32(2)) % 32)
	if v26 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v66 = int32(-1)
	v68 = int32(150) - v26
	if v22&(v66<<(uint(v68)%32)^v66) != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v650 = int32(base.Ui32(v22|int32(_a_F_float_to_shortest_decimal_buf_2)) >> (uint(v68) % 32))
	v655 = v3
	goto L24
L28:
	;
	v81 = v78 | int32(33554432)
	goto L30
L29:
	;
	v81 = v78
	goto L30
L30:
	;
	v84 = int32(2)
	v89 = v81 + (base.B2i32(v22 != int32(0)) | base.B2i32(base.Ui32(v26) < base.Ui32(v84)) ^ int32(-1))
	v91 = v81 | v84
	if v26 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v642 = v625 + v629
	v643 = v624 + v641
	if base.Ui32(v643) <= base.Ui32(int32(99999999)) {
		v650 = v643
		v655 = v642
		goto L24
	} else {
		goto L81
	}
L32:
	;
	v560 = int32(0)
	v561 = int32(10)
	v562 = base.I32_div_u_s(v546, v561)
	v564 = base.I32_div_u_s(v550, v561)
	if base.Ui32(v564) < base.Ui32(v562) {
		goto L75
	} else {
		goto L76
	}
L33:
	;
	v464 = int32(0)
	v465 = int32(10)
	v466 = base.I32_div_u_s(v450, v465)
	v468 = base.I32_div_u_s(v454, v465)
	if base.Ui32(v466) <= base.Ui32(v468) {
		goto L69
	} else {
		goto L70
	}
L34:
	;
	v95 = v26 - int32(152)
	goto L36
L35:
	;
	v95 = int32(-151)
	goto L36
L36:
	;
	if int32(0) <= v95 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = int32(base.Ui32(v95*int32(_a_F_float_to_shortest_decimal_buf_3)) >> (uint(int32(18)) % 32))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v101<<(uint(int32(3))%32))+uint32(_c_F_float_to_shortest_decimal_buf[2])))
	v106 = v104 & int64(4294967295)
	v107 = base.I64_extend_i32_u(v89)
	v109 = int64(32)
	v111 = base.I32_wrap_i64(int64(base.Ui64(v106*v107) >> (uint(v109) % 64)))
	v113 = int64(base.Ui64(v104) >> (uint(v109) % 64))
	v114 = v107 * v113
	v116 = v111 + base.I32_wrap_i64(v114)
	v123 = v101 - v95
	v128 = v123 + int32(base.Ui32(v101*int32(_a_F_float_to_shortest_decimal_buf_4))>>(uint(int32(19))%32))
	v129 = int32(5) - v128
	v132 = v128 + int32(27)
	v134 = (base.B2i32(base.Ui32(v116) < base.Ui32(v111))+base.I32_wrap_i64(int64(base.Ui64(v114)>>(uint(v109)%64))))<<(uint(v129)%32) | int32(base.Ui32(v116)>>(uint(v132)%32))
	v135 = base.I64_extend_i32_u(v91)
	v139 = base.I32_wrap_i64(int64(base.Ui64(v106*v135) >> (uint(v109) % 64)))
	v140 = v135 * v113
	v142 = v139 + base.I32_wrap_i64(v140)
	v150 = (base.B2i32(base.Ui32(v142) < base.Ui32(v139))+base.I32_wrap_i64(int64(base.Ui64(v140)>>(uint(v109)%64))))<<(uint(v129)%32) | int32(base.Ui32(v142)>>(uint(v132)%32))
	v151 = base.I64_extend_i32_u(v81)
	v155 = base.I32_wrap_i64(int64(base.Ui64(v106*v151) >> (uint(v109) % 64)))
	v156 = v151 * v113
	v158 = v155 + base.I32_wrap_i64(v156)
	v166 = (base.B2i32(base.Ui32(v158) < base.Ui32(v155))+base.I32_wrap_i64(int64(base.Ui64(v156)>>(uint(v109)%64))))<<(uint(v129)%32) | int32(base.Ui32(v158)>>(uint(v132)%32))
	v167 = int32(0)
	if v101 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v310 = v95 * int32(-732923)
	v312 = int32(base.Ui32(v310) >> (uint(int32(20)) % 32))
	v313 = v95 + v312
	v317 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_float_to_shortest_decimal_buf_5)-v313<<(uint(int32(3))%32))))
	v319 = v317 & int64(4294967295)
	v320 = base.I64_extend_i32_u(v89)
	v322 = int64(32)
	v324 = base.I32_wrap_i64(int64(base.Ui64(v319*v320) >> (uint(v322) % 64)))
	v326 = int64(base.Ui64(v317) >> (uint(v322) % 64))
	v327 = v320 * v326
	v329 = v324 + base.I32_wrap_i64(v327)
	v340 = v312 - int32(base.Ui32(v313*int32(-1217359))>>(uint(int32(19))%32))
	v341 = int32(4) - v340
	v344 = v340 + int32(28)
	v346 = (base.B2i32(base.Ui32(v329) < base.Ui32(v324))+base.I32_wrap_i64(int64(base.Ui64(v327)>>(uint(v322)%64))))<<(uint(v341)%32) | int32(base.Ui32(v329)>>(uint(v344)%32))
	v347 = base.I64_extend_i32_u(v81)
	v351 = base.I32_wrap_i64(int64(base.Ui64(v319*v347) >> (uint(v322) % 64)))
	v352 = v347 * v326
	v354 = v351 + base.I32_wrap_i64(v352)
	v362 = (base.B2i32(base.Ui32(v354) < base.Ui32(v351))+base.I32_wrap_i64(int64(base.Ui64(v352)>>(uint(v322)%64))))<<(uint(v341)%32) | int32(base.Ui32(v354)>>(uint(v344)%32))
	v363 = base.I64_extend_i32_u(v91)
	v367 = base.I32_wrap_i64(int64(base.Ui64(v319*v363) >> (uint(v322) % 64)))
	v368 = v326 * v363
	v370 = v367 + base.I32_wrap_i64(v368)
	v378 = (base.B2i32(base.Ui32(v370) < base.Ui32(v367))+base.I32_wrap_i64(int64(base.Ui64(v368)>>(uint(v322)%64))))<<(uint(v341)%32) | int32(base.Ui32(v370)>>(uint(v344)%32))
	v380 = v378 - int32(1)
	if v312 != 0 {
		goto L61
	} else {
		goto L62
	}
L40:
	;
	v171 = int32(10)
	v172 = base.I32_div_u_s(v150-int32(1), v171)
	v174 = base.I32_div_u_s(v134, v171)
	if base.Ui32(v172) <= base.Ui32(v174) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v218 = v167
	goto L42
L42:
	;
	v224 = base.I32_rem_u_s(v81, int32(5))
	if v224 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v177 = v101 - int32(1)
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v177<<(uint(int32(3))%32))+uint32(_c_F_float_to_shortest_decimal_buf[2])))
	v184 = int64(32)
	v186 = base.I32_wrap_i64(int64(base.Ui64(v180&int64(4294967295)*v151) >> (uint(v184) % 64)))
	v189 = int64(base.Ui64(v180)>>(uint(v184)%64)) * v151
	v191 = v186 + base.I32_wrap_i64(v189)
	v202 = v123 + int32(base.Ui32(v177*int32(_a_F_float_to_shortest_decimal_buf_4))>>(uint(int32(19))%32))
	v210 = base.I32_rem_u_s((base.B2i32(base.Ui32(v191) < base.Ui32(v186))+base.I32_wrap_i64(int64(base.Ui64(v189)>>(uint(v184)%64))))<<(uint(int32(6)-v202)%32)|int32(base.Ui32(v191)>>(uint(v202+int32(26))%32)), int32(10))
	v211 = v210
	goto L45
L44:
	;
	v211 = v167
	goto L45
L45:
	;
	if base.Ui32(int32(33)) < base.Ui32(v95) {
		v543 = v166
		v546 = v150
		v548 = v101
		v549 = v211
		v550 = v134
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v218 = v211
	goto L42
L47:
	;
	v230 = v81
	v231 = v167
	goto L50
L48:
	;
	goto L49
L49:
	;
	v255 = int32(0)
	v257 = base.I32_rem_u_s(v91, int32(5))
	if v257 == v255 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v247 = v231 + int32(1)
	v248 = int32(5)
	v249 = base.I32_div_u_s(v230, v248)
	v251 = base.I32_rem_u_s(v249, v248)
	if v251 == int32(0) {
		v230 = v249
		v231 = v247
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if base.Ui32(v247) < base.Ui32(v101) {
		v543 = v166
		v546 = v150
		v548 = v101
		v549 = v218
		v550 = v134
		goto L32
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v447 = v166
	v450 = v150
	v452 = v101
	v453 = v218
	v454 = v134
	goto L33
L54:
	;
	v263 = v255
	v266 = v91
	goto L57
L55:
	;
	v290 = v255
	goto L56
L56:
	;
	v543 = v166
	v546 = v150 - base.B2i32(base.Ui32(v101) <= base.Ui32(v290))
	v548 = v101
	v549 = v218
	v550 = v134
	goto L32
L57:
	;
	v280 = v263 + int32(1)
	v281 = int32(5)
	v282 = base.I32_div_u_s(v266, v281)
	v284 = base.I32_rem_u_s(v282, v281)
	if v284 == int32(0) {
		v263 = v280
		v266 = v282
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v290 = v280
	goto L56
L59:
	;
	goto L58
L60:
	;
	v435 = int32(-1)
	if v81&(v435<<(uint(v312-int32(1))%32)^v435)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v310)) != 0 {
		v543 = v362
		v546 = v378
		v548 = v313
		v549 = v424
		v550 = v346
		goto L32
	} else {
		goto L68
	}
L61:
	;
	v381 = int32(10)
	v382 = base.I32_div_u_s(v380, v381)
	v384 = base.I32_div_u_s(v346, v381)
	if base.Ui32(v382) <= base.Ui32(v384) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v431 = v3
	goto L63
L63:
	;
	v447 = v362
	v450 = v380
	v452 = v313
	v453 = v431
	v454 = v346
	goto L33
L64:
	;
	v387 = int32(1) - v313
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v387<<(uint(int32(3))%32))+uint32(_c_F_float_to_shortest_decimal_buf[3])))
	v394 = int64(32)
	v396 = base.I32_wrap_i64(int64(base.Ui64(v390&int64(4294967295)*v347) >> (uint(v394) % 64)))
	v399 = int64(base.Ui64(v390)>>(uint(v394)%64)) * v347
	v401 = v396 + base.I32_wrap_i64(v399)
	v414 = v312 + (int32(base.Ui32(v387*int32(_a_F_float_to_shortest_decimal_buf_4))>>(uint(int32(19))%32)) ^ int32(-1))
	v422 = base.I32_rem_u_s((base.B2i32(base.Ui32(v401) < base.Ui32(v396))+base.I32_wrap_i64(int64(base.Ui64(v399)>>(uint(v394)%64))))<<(uint(int32(4)-v414)%32)|int32(base.Ui32(v401)>>(uint(v414+int32(28))%32)), int32(10))
	v424 = v422
	goto L66
L65:
	;
	v424 = v3
	goto L66
L66:
	;
	if v312 != int32(1) {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v431 = v424
	goto L63
L68:
	;
	v447 = v362
	v450 = v378
	v452 = v313
	v453 = v424
	v454 = v346
	goto L33
L69:
	;
	v512 = v447
	v513 = v464
	v518 = v453
	v519 = v454
	v529 = int32(0)
	goto L71
L70:
	;
	v474 = v447
	v475 = v464
	v476 = v466
	v477 = int32(1)
	v478 = v468
	v480 = v453
	goto L72
L71:
	;
	v531 = v518 & int32(255)
	v624 = v512
	v625 = v513
	v629 = v452
	v641 = (v529|base.B2i32(v531 != int32(5))|v512)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v531)) | base.B2i32(v512 == v519)
	goto L31
L72:
	;
	v492 = v475 + int32(1)
	v493 = int32(10)
	v494 = base.I32_div_u_s(v474, v493)
	v497 = v474 - v494*v493
	v502 = v477 & base.B2i32(v480&int32(255) == int32(0))
	v504 = base.I32_div_u_s(v476, v493)
	v506 = base.I32_div_u_s(v478, v493)
	if base.Ui32(v506) < base.Ui32(v504) {
		v474 = v494
		v475 = v492
		v476 = v504
		v477 = v502
		v478 = v506
		v480 = v497
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v512 = v494
	v513 = v492
	v518 = v497
	v519 = v478
	v529 = v502 ^ int32(1)
	goto L71
L74:
	;
	goto L73
L75:
	;
	v568 = v543
	v569 = v560
	v570 = v562
	v572 = v564
	goto L78
L76:
	;
	v599 = v543
	v600 = v560
	v605 = v549
	v606 = v550
	goto L77
L77:
	;
	v624 = v599
	v625 = v600
	v629 = v548
	v641 = base.B2i32(v599 == v606) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v605&int32(255)))
	goto L31
L78:
	;
	v586 = v569 + int32(1)
	v587 = int32(10)
	v588 = base.I32_div_u_s(v568, v587)
	v590 = base.I32_div_u_s(v570, v587)
	v592 = base.I32_div_u_s(v572, v587)
	if base.Ui32(v592) < base.Ui32(v590) {
		v568 = v588
		v569 = v586
		v570 = v590
		v572 = v592
		goto L78
	} else {
		goto L80
	}
L79:
	;
	v599 = v588
	v600 = v586
	v605 = v568 - v588*int32(10)
	v606 = v572
	goto L77
L80:
	;
	goto L79
L81:
	;
	v692 = v643
	v697 = v642
	v708 = int32(9)
	goto L23
L82:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_buf_6)) < base.Ui32(v650) {
		v692 = v650
		v697 = v655
		v708 = int32(7)
		goto L23
	} else {
		goto L83
	}
L83:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_buf_7)) < base.Ui32(v650) {
		v692 = v650
		v697 = v655
		v708 = int32(6)
		goto L23
	} else {
		goto L84
	}
L84:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_buf_8)) < base.Ui32(v650) {
		v692 = v650
		v697 = v655
		v708 = int32(5)
		goto L23
	} else {
		goto L85
	}
L85:
	;
	if base.Ui32(int32(999)) < base.Ui32(v650) {
		v692 = v650
		v697 = v655
		v708 = int32(4)
		goto L23
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(int32(99)) < base.Ui32(v650) {
		v692 = v650
		v697 = v655
		v708 = int32(3)
		goto L23
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(int32(9)) < base.Ui32(v650) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v688 = int32(2)
	goto L90
L89:
	;
	v688 = int32(1)
	goto L90
L90:
	;
	v692 = v650
	v697 = v655
	v708 = v688
	goto L23
L91:
	;
	v712 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v712)
	v715 = int32(1)
	goto L93
L92:
	;
	v715 = v709
	goto L93
L93:
	;
	v716 = v708 + v697
	if base.Ui32(v716+int32(3)) <= base.Ui32(int32(9)) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v927 = int32(0)
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_buf_9)) <= base.Ui32(v692) {
		goto L134
	} else {
		goto L135
	}
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v721))) = v923
	v925 = v922
	goto L94
L96:
	;
	v721 = l1 + v715
	v722 = int32(0)
	if v716 <= v722 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	if v697 != 0 {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v922 = int32(2) - v716
	v923 = int64(3472328296227679792)
	goto L95
L100:
	;
	goto L101
L101:
	;
	if int32(0) <= v697 {
		v922 = v722
		v923 = int64(3472328296227680304)
		goto L95
	} else {
		goto L102
	}
L102:
	;
	v925 = int32(1)
	goto L94
L103:
	;
	v781 = int32(0)
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_buf_9)) <= base.Ui32(v764) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	v764 = v692
	v769 = v708
	goto L103
L105:
	;
	goto L106
L106:
	;
	v735 = v692
	v737 = v708
	goto L107
L107:
	;
	if v735&int32(1) != 0 {
		v764 = v735
		v769 = v737
		goto L103
	} else {
		goto L109
	}
L108:
	;
	v764 = v735
	v769 = v737
	goto L103
L109:
	;
	v758 = base.I32_div_u_s(v735, int32(10))
	if int32(0)-v735 == v758*int32(-10) {
		v735 = v758
		v737 = v737 - int32(1)
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v788 = v764
	v789 = v781
	goto L114
L112:
	;
	v834 = v764
	v835 = v781
	goto L113
L113:
	;
	if base.Ui32(v834) < base.Ui32(int32(100)) {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	v805 = l1 + v715 + v769 - v789
	v809 = base.I32_div_u_s(v788, int32(_a_F_float_to_shortest_decimal_buf_9))
	v812 = v788 + v809*int32(-10000)
	v813 = int32(100)
	v814 = base.I32_div_u_s(v812, v813)
	v815 = int32(1)
	v817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v814<<(uint(v815)%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v805-int32(3)))) = uint16(v817)
	v826 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v812-v814*v813)<<(uint(v815)%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v805-v815))) = uint16(v826)
	v829 = v789 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v788) {
		v788 = v809
		v789 = v829
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v834 = v809
	v835 = v829
	goto L113
L116:
	;
	goto L115
L117:
	;
	v876 = v716 - int32(1)
	v877 = l1 + v715
	if base.Ui32(int32(10)) <= base.Ui32(v874) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v873 = v835
	v874 = v834
	goto L117
L119:
	;
	goto L120
L120:
	;
	v858 = int32(_a_F_float_to_shortest_decimal_buf_10)
	v860 = int32(100)
	v861 = base.I32_div_u_s(v834&v858, v860)
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v834-v861*v860)&v858<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v715+v769+(v835^int32(-1))))) = uint16(v869)
	v873 = v835 | int32(2)
	v874 = v861
	goto L117
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v877))) = uint8(v891)
	if base.Ui32(int32(2)) <= base.Ui32(v769) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v884 = v874 << (uint(int32(1)) % 32)
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884)+uint32(_c_F_float_to_shortest_decimal_buf[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+(v715+v769-v873)))) = uint8(v885)
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884)+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	v891 = v887
	goto L121
L123:
	;
	goto L124
L124:
	;
	v891 = v874 | int32(48)
	goto L121
L125:
	;
	v895 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v877)+1)) = uint8(v895)
	v900 = v769 + int32(1)
	goto L127
L126:
	;
	v900 = int32(1)
	goto L127
L127:
	;
	v901 = v900 + v715
	v902 = l1 + v901
	v903 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v902))) = uint8(v903)
	v908 = base.B2i32(v876 < int32(0))
	if v876 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v909 = int32(45)
	goto L130
L129:
	;
	v909 = int32(43)
	goto L130
L130:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)) = uint8(v909)
	if v876 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v913 = int32(1) - v716
	goto L133
L132:
	;
	v913 = v876
	goto L133
L133:
	;
	v918 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v913<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v902)+2)) = uint16(v918)
	v1099 = v901 + int32(4)
	goto L1
L134:
	;
	v934 = v927
	v935 = v692
	goto L137
L135:
	;
	v980 = v927
	v981 = v692
	goto L136
L136:
	;
	if base.Ui32(v981) < base.Ui32(int32(100)) {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	v951 = v925 + v721 + v708 - v934
	v952 = int32(4)
	v955 = base.I32_div_u_s(v935, int32(_a_F_float_to_shortest_decimal_buf_9))
	v958 = v935 + v955*int32(-10000)
	v959 = int32(100)
	v960 = base.I32_div_u_s(v958, v959)
	v961 = int32(1)
	v963 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v960<<(uint(v961)%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v951-v952))) = uint16(v963)
	v972 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v958-v960*v959)<<(uint(v961)%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v951-int32(2)))) = uint16(v972)
	v975 = v934 + v952
	if base.Ui32(int32(99999999)) < base.Ui32(v935) {
		v934 = v975
		v935 = v955
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v980 = v975
	v981 = v955
	goto L136
L139:
	;
	goto L138
L140:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1020) {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	v1019 = v980
	v1020 = v981
	goto L140
L142:
	;
	goto L143
L143:
	;
	v1002 = int32(2)
	v1004 = int32(_a_F_float_to_shortest_decimal_buf_10)
	v1006 = int32(100)
	v1007 = base.I32_div_u_s(v981&v1004, v1006)
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v981-v1007*v1006)&v1004<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v925+v721+v708-v980-v1002))) = uint16(v1015)
	v1019 = v980 | v1002
	v1020 = v1007
	goto L140
L144:
	;
	v1036 = int32(1)
	if v925 == v1036 {
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v1030 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1020<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_buf[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v925+v721+v708-v1019-int32(2)))) = uint16(v1030)
	goto L144
L146:
	;
	goto L147
L147:
	;
	v1034 = v1020 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v925+v721))) = uint8(v1034)
	goto L144
L148:
	;
	v1099 = v1076 + int32(base.Ui32(v20)>>(uint(int32(31))%32))
	goto L1
L149:
	;
	if v716&int32(4) != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	if v697 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L152:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v721)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v721))) = v1041
	v1044 = int32(5)
	goto L154
L153:
	;
	v1044 = v1036
	goto L154
L154:
	;
	if v716&int32(2) != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1047 = v1044 + v721
	v1050 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1047))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1047-int32(1)))) = uint16(v1050)
	v1055 = v1044 | int32(2)
	goto L157
L156:
	;
	v1055 = v1044
	goto L157
L157:
	;
	if v716&int32(1) != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1058 = v1055 + v721
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1058))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1058-int32(1)))) = uint8(v1061)
	goto L160
L159:
	;
	goto L160
L160:
	;
	v1065 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v721+v716))) = uint8(v1065)
	v1076 = v708 + int32(1)
	goto L148
L161:
	;
	v1073 = int32(2) - v697
	goto L163
L162:
	;
	v1073 = v716
	goto L163
L163:
	;
	v1076 = v1073
	goto L148
}
func F_float_to_shortest_decimal_bufn(m *base.Module, l0 float32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int64
	_ = v320
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v370 int32
	_ = v370
	var v371 int64
	_ = v371
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int64
	_ = v393
	var v397 int64
	_ = v397
	var v399 int32
	_ = v399
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v927 int64
	_ = v927
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	v3 = int32(0)
	v20 = base.I32_reinterpret_f32(l0)
	v22 = v20 & int32(_a_F_float_to_shortest_decimal_bufn_0)
	v25 = int32(255)
	v26 = int32(base.Ui32(v20)>>(uint(int32(23))%32)) & v25
	if v26|v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = base.B2i32(v26 != v25)
	goto L3
L2:
	;
	v31 = v3
	goto L3
L3:
	;
	if v31 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(23)) < base.Ui32(v26-int32(127)) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_float_to_shortest_decimal_bufn[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v35)
	v38 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_float_to_shortest_decimal_bufn[1])))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v38)
	return int32(3)
L8:
	;
	goto L9
L9:
	;
	if v20 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v44)
	goto L12
L11:
	;
	goto L12
L12:
	;
	v48 = l1 + int32(base.Ui32(v20)>>(uint(int32(31))%32))
	if v26 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(8751735898823355977)
	if v20 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v57 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v57)
	if v20 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v55 = int32(9)
	goto L18
L17:
	;
	v55 = int32(8)
	goto L18
L18:
	;
	return v55
L19:
	;
	v63 = int32(2)
	goto L21
L20:
	;
	v63 = int32(1)
	goto L21
L21:
	;
	return v63
L22:
	;
	v712 = int32(0)
	if v20 < v712 {
		goto L90
	} else {
		goto L91
	}
L23:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_bufn_1)) < base.Ui32(v653) {
		v695 = v653
		v700 = v658
		v711 = int32(8)
		goto L22
	} else {
		goto L81
	}
L24:
	;
	v81 = v22 << (uint(int32(2)) % 32)
	if v26 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v69 = int32(-1)
	v71 = int32(150) - v26
	if v22&(v69<<(uint(v71)%32)^v69) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v653 = int32(base.Ui32(v22|int32(_a_F_float_to_shortest_decimal_bufn_2)) >> (uint(v71) % 32))
	v658 = v3
	goto L23
L27:
	;
	v84 = v81 | int32(33554432)
	goto L29
L28:
	;
	v84 = v81
	goto L29
L29:
	;
	v87 = int32(2)
	v92 = v84 + (base.B2i32(v22 != int32(0)) | base.B2i32(base.Ui32(v26) < base.Ui32(v87)) ^ int32(-1))
	v94 = v84 | v87
	if v26 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v645 = v628 + v632
	v646 = v627 + v644
	if base.Ui32(v646) <= base.Ui32(int32(99999999)) {
		v653 = v646
		v658 = v645
		goto L23
	} else {
		goto L80
	}
L31:
	;
	v563 = int32(0)
	v564 = int32(10)
	v565 = base.I32_div_u_s(v549, v564)
	v567 = base.I32_div_u_s(v553, v564)
	if base.Ui32(v567) < base.Ui32(v565) {
		goto L74
	} else {
		goto L75
	}
L32:
	;
	v467 = int32(0)
	v468 = int32(10)
	v469 = base.I32_div_u_s(v453, v468)
	v471 = base.I32_div_u_s(v457, v468)
	if base.Ui32(v469) <= base.Ui32(v471) {
		goto L68
	} else {
		goto L69
	}
L33:
	;
	v98 = v26 - int32(152)
	goto L35
L34:
	;
	v98 = int32(-151)
	goto L35
L35:
	;
	if int32(0) <= v98 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v104 = int32(base.Ui32(v98*int32(_a_F_float_to_shortest_decimal_bufn_3)) >> (uint(int32(18)) % 32))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[2])))
	v109 = v107 & int64(4294967295)
	v110 = base.I64_extend_i32_u(v92)
	v112 = int64(32)
	v114 = base.I32_wrap_i64(int64(base.Ui64(v109*v110) >> (uint(v112) % 64)))
	v116 = int64(base.Ui64(v107) >> (uint(v112) % 64))
	v117 = v110 * v116
	v119 = v114 + base.I32_wrap_i64(v117)
	v126 = v104 - v98
	v131 = v126 + int32(base.Ui32(v104*int32(_a_F_float_to_shortest_decimal_bufn_4))>>(uint(int32(19))%32))
	v132 = int32(5) - v131
	v135 = v131 + int32(27)
	v137 = (base.B2i32(base.Ui32(v119) < base.Ui32(v114))+base.I32_wrap_i64(int64(base.Ui64(v117)>>(uint(v112)%64))))<<(uint(v132)%32) | int32(base.Ui32(v119)>>(uint(v135)%32))
	v138 = base.I64_extend_i32_u(v94)
	v142 = base.I32_wrap_i64(int64(base.Ui64(v109*v138) >> (uint(v112) % 64)))
	v143 = v138 * v116
	v145 = v142 + base.I32_wrap_i64(v143)
	v153 = (base.B2i32(base.Ui32(v145) < base.Ui32(v142))+base.I32_wrap_i64(int64(base.Ui64(v143)>>(uint(v112)%64))))<<(uint(v132)%32) | int32(base.Ui32(v145)>>(uint(v135)%32))
	v154 = base.I64_extend_i32_u(v84)
	v158 = base.I32_wrap_i64(int64(base.Ui64(v109*v154) >> (uint(v112) % 64)))
	v159 = v154 * v116
	v161 = v158 + base.I32_wrap_i64(v159)
	v169 = (base.B2i32(base.Ui32(v161) < base.Ui32(v158))+base.I32_wrap_i64(int64(base.Ui64(v159)>>(uint(v112)%64))))<<(uint(v132)%32) | int32(base.Ui32(v161)>>(uint(v135)%32))
	v170 = int32(0)
	if v104 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v313 = v98 * int32(-732923)
	v315 = int32(base.Ui32(v313) >> (uint(int32(20)) % 32))
	v316 = v98 + v315
	v320 = *(*int64)(unsafe.Add(mBase, uint32(int32(_a_F_float_to_shortest_decimal_bufn_5)-v316<<(uint(int32(3))%32))))
	v322 = v320 & int64(4294967295)
	v323 = base.I64_extend_i32_u(v92)
	v325 = int64(32)
	v327 = base.I32_wrap_i64(int64(base.Ui64(v322*v323) >> (uint(v325) % 64)))
	v329 = int64(base.Ui64(v320) >> (uint(v325) % 64))
	v330 = v323 * v329
	v332 = v327 + base.I32_wrap_i64(v330)
	v343 = v315 - int32(base.Ui32(v316*int32(-1217359))>>(uint(int32(19))%32))
	v344 = int32(4) - v343
	v347 = v343 + int32(28)
	v349 = (base.B2i32(base.Ui32(v332) < base.Ui32(v327))+base.I32_wrap_i64(int64(base.Ui64(v330)>>(uint(v325)%64))))<<(uint(v344)%32) | int32(base.Ui32(v332)>>(uint(v347)%32))
	v350 = base.I64_extend_i32_u(v84)
	v354 = base.I32_wrap_i64(int64(base.Ui64(v322*v350) >> (uint(v325) % 64)))
	v355 = v350 * v329
	v357 = v354 + base.I32_wrap_i64(v355)
	v365 = (base.B2i32(base.Ui32(v357) < base.Ui32(v354))+base.I32_wrap_i64(int64(base.Ui64(v355)>>(uint(v325)%64))))<<(uint(v344)%32) | int32(base.Ui32(v357)>>(uint(v347)%32))
	v366 = base.I64_extend_i32_u(v94)
	v370 = base.I32_wrap_i64(int64(base.Ui64(v322*v366) >> (uint(v325) % 64)))
	v371 = v329 * v366
	v373 = v370 + base.I32_wrap_i64(v371)
	v381 = (base.B2i32(base.Ui32(v373) < base.Ui32(v370))+base.I32_wrap_i64(int64(base.Ui64(v371)>>(uint(v325)%64))))<<(uint(v344)%32) | int32(base.Ui32(v373)>>(uint(v347)%32))
	v383 = v381 - int32(1)
	if v315 != 0 {
		goto L60
	} else {
		goto L61
	}
L39:
	;
	v174 = int32(10)
	v175 = base.I32_div_u_s(v153-int32(1), v174)
	v177 = base.I32_div_u_s(v137, v174)
	if base.Ui32(v175) <= base.Ui32(v177) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v221 = v170
	goto L41
L41:
	;
	v227 = base.I32_rem_u_s(v84, int32(5))
	if v227 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v180 = v104 - int32(1)
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v180<<(uint(int32(3))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[2])))
	v187 = int64(32)
	v189 = base.I32_wrap_i64(int64(base.Ui64(v183&int64(4294967295)*v154) >> (uint(v187) % 64)))
	v192 = int64(base.Ui64(v183)>>(uint(v187)%64)) * v154
	v194 = v189 + base.I32_wrap_i64(v192)
	v205 = v126 + int32(base.Ui32(v180*int32(_a_F_float_to_shortest_decimal_bufn_4))>>(uint(int32(19))%32))
	v213 = base.I32_rem_u_s((base.B2i32(base.Ui32(v194) < base.Ui32(v189))+base.I32_wrap_i64(int64(base.Ui64(v192)>>(uint(v187)%64))))<<(uint(int32(6)-v205)%32)|int32(base.Ui32(v194)>>(uint(v205+int32(26))%32)), int32(10))
	v214 = v213
	goto L44
L43:
	;
	v214 = v170
	goto L44
L44:
	;
	if base.Ui32(int32(33)) < base.Ui32(v98) {
		v546 = v169
		v549 = v153
		v551 = v104
		v552 = v214
		v553 = v137
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v221 = v214
	goto L41
L46:
	;
	v233 = v84
	v234 = v170
	goto L49
L47:
	;
	goto L48
L48:
	;
	v258 = int32(0)
	v260 = base.I32_rem_u_s(v94, int32(5))
	if v260 == v258 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v250 = v234 + int32(1)
	v251 = int32(5)
	v252 = base.I32_div_u_s(v233, v251)
	v254 = base.I32_rem_u_s(v252, v251)
	if v254 == int32(0) {
		v233 = v252
		v234 = v250
		goto L49
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(v250) < base.Ui32(v104) {
		v546 = v169
		v549 = v153
		v551 = v104
		v552 = v221
		v553 = v137
		goto L31
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v450 = v169
	v453 = v153
	v455 = v104
	v456 = v221
	v457 = v137
	goto L32
L53:
	;
	v266 = v258
	v269 = v94
	goto L56
L54:
	;
	v293 = v258
	goto L55
L55:
	;
	v546 = v169
	v549 = v153 - base.B2i32(base.Ui32(v104) <= base.Ui32(v293))
	v551 = v104
	v552 = v221
	v553 = v137
	goto L31
L56:
	;
	v283 = v266 + int32(1)
	v284 = int32(5)
	v285 = base.I32_div_u_s(v269, v284)
	v287 = base.I32_rem_u_s(v285, v284)
	if v287 == int32(0) {
		v266 = v283
		v269 = v285
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v293 = v283
	goto L55
L58:
	;
	goto L57
L59:
	;
	v438 = int32(-1)
	if v84&(v438<<(uint(v315-int32(1))%32)^v438)|base.B2i32(base.Ui32(int32(32505855)) < base.Ui32(v313)) != 0 {
		v546 = v365
		v549 = v381
		v551 = v316
		v552 = v427
		v553 = v349
		goto L31
	} else {
		goto L67
	}
L60:
	;
	v384 = int32(10)
	v385 = base.I32_div_u_s(v383, v384)
	v387 = base.I32_div_u_s(v349, v384)
	if base.Ui32(v385) <= base.Ui32(v387) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v434 = v3
	goto L62
L62:
	;
	v450 = v365
	v453 = v383
	v455 = v316
	v456 = v434
	v457 = v349
	goto L32
L63:
	;
	v390 = int32(1) - v316
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v390<<(uint(int32(3))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[3])))
	v397 = int64(32)
	v399 = base.I32_wrap_i64(int64(base.Ui64(v393&int64(4294967295)*v350) >> (uint(v397) % 64)))
	v402 = int64(base.Ui64(v393)>>(uint(v397)%64)) * v350
	v404 = v399 + base.I32_wrap_i64(v402)
	v417 = v315 + (int32(base.Ui32(v390*int32(_a_F_float_to_shortest_decimal_bufn_4))>>(uint(int32(19))%32)) ^ int32(-1))
	v425 = base.I32_rem_u_s((base.B2i32(base.Ui32(v404) < base.Ui32(v399))+base.I32_wrap_i64(int64(base.Ui64(v402)>>(uint(v397)%64))))<<(uint(int32(4)-v417)%32)|int32(base.Ui32(v404)>>(uint(v417+int32(28))%32)), int32(10))
	v427 = v425
	goto L65
L64:
	;
	v427 = v3
	goto L65
L65:
	;
	if v315 != int32(1) {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v434 = v427
	goto L62
L67:
	;
	v450 = v365
	v453 = v381
	v455 = v316
	v456 = v427
	v457 = v349
	goto L32
L68:
	;
	v515 = v450
	v516 = v467
	v521 = v456
	v522 = v457
	v532 = int32(0)
	goto L70
L69:
	;
	v477 = v450
	v478 = v467
	v479 = v469
	v480 = int32(1)
	v481 = v471
	v483 = v456
	goto L71
L70:
	;
	v534 = v521 & int32(255)
	v627 = v515
	v628 = v516
	v632 = v455
	v644 = (v532|base.B2i32(v534 != int32(5))|v515)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v534)) | base.B2i32(v515 == v522)
	goto L30
L71:
	;
	v495 = v478 + int32(1)
	v496 = int32(10)
	v497 = base.I32_div_u_s(v477, v496)
	v500 = v477 - v497*v496
	v505 = v480 & base.B2i32(v483&int32(255) == int32(0))
	v507 = base.I32_div_u_s(v479, v496)
	v509 = base.I32_div_u_s(v481, v496)
	if base.Ui32(v509) < base.Ui32(v507) {
		v477 = v497
		v478 = v495
		v479 = v507
		v480 = v505
		v481 = v509
		v483 = v500
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v515 = v497
	v516 = v495
	v521 = v500
	v522 = v481
	v532 = v505 ^ int32(1)
	goto L70
L73:
	;
	goto L72
L74:
	;
	v571 = v546
	v572 = v563
	v573 = v565
	v575 = v567
	goto L77
L75:
	;
	v602 = v546
	v603 = v563
	v608 = v552
	v609 = v553
	goto L76
L76:
	;
	v627 = v602
	v628 = v603
	v632 = v551
	v644 = base.B2i32(v602 == v609) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v608&int32(255)))
	goto L30
L77:
	;
	v589 = v572 + int32(1)
	v590 = int32(10)
	v591 = base.I32_div_u_s(v571, v590)
	v593 = base.I32_div_u_s(v573, v590)
	v595 = base.I32_div_u_s(v575, v590)
	if base.Ui32(v595) < base.Ui32(v593) {
		v571 = v591
		v572 = v589
		v573 = v593
		v575 = v595
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v602 = v591
	v603 = v589
	v608 = v571 - v591*int32(10)
	v609 = v575
	goto L76
L79:
	;
	goto L78
L80:
	;
	v695 = v646
	v700 = v645
	v711 = int32(9)
	goto L22
L81:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_bufn_6)) < base.Ui32(v653) {
		v695 = v653
		v700 = v658
		v711 = int32(7)
		goto L22
	} else {
		goto L82
	}
L82:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_bufn_7)) < base.Ui32(v653) {
		v695 = v653
		v700 = v658
		v711 = int32(6)
		goto L22
	} else {
		goto L83
	}
L83:
	;
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_bufn_8)) < base.Ui32(v653) {
		v695 = v653
		v700 = v658
		v711 = int32(5)
		goto L22
	} else {
		goto L84
	}
L84:
	;
	if base.Ui32(int32(999)) < base.Ui32(v653) {
		v695 = v653
		v700 = v658
		v711 = int32(4)
		goto L22
	} else {
		goto L85
	}
L85:
	;
	if base.Ui32(int32(99)) < base.Ui32(v653) {
		v695 = v653
		v700 = v658
		v711 = int32(3)
		goto L22
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(int32(9)) < base.Ui32(v653) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v691 = int32(2)
	goto L89
L88:
	;
	v691 = int32(1)
	goto L89
L89:
	;
	v695 = v653
	v700 = v658
	v711 = v691
	goto L22
L90:
	;
	v715 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v715)
	v718 = int32(1)
	goto L92
L91:
	;
	v718 = v712
	goto L92
L92:
	;
	v719 = v711 + v700
	if base.Ui32(v719+int32(3)) <= base.Ui32(int32(9)) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v931 = int32(0)
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_bufn_9)) <= base.Ui32(v695) {
		goto L133
	} else {
		goto L134
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v927
	v929 = v926
	goto L93
L95:
	;
	v724 = l1 + v718
	v725 = int32(0)
	if v719 <= v725 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	if v700 != 0 {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	v926 = int32(2) - v719
	v927 = int64(3472328296227679792)
	goto L94
L99:
	;
	goto L100
L100:
	;
	if int32(0) <= v700 {
		v926 = v725
		v927 = int64(3472328296227680304)
		goto L94
	} else {
		goto L101
	}
L101:
	;
	v929 = int32(1)
	goto L93
L102:
	;
	v784 = int32(0)
	if base.Ui32(int32(_a_F_float_to_shortest_decimal_bufn_9)) <= base.Ui32(v767) {
		goto L110
	} else {
		goto L111
	}
L103:
	;
	v767 = v695
	v772 = v711
	goto L102
L104:
	;
	goto L105
L105:
	;
	v738 = v695
	v740 = v711
	goto L106
L106:
	;
	if v738&int32(1) != 0 {
		v767 = v738
		v772 = v740
		goto L102
	} else {
		goto L108
	}
L107:
	;
	v767 = v738
	v772 = v740
	goto L102
L108:
	;
	v761 = base.I32_div_u_s(v738, int32(10))
	if int32(0)-v738 == v761*int32(-10) {
		v738 = v761
		v740 = v740 - int32(1)
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v791 = v767
	v792 = v784
	goto L113
L111:
	;
	v837 = v767
	v838 = v784
	goto L112
L112:
	;
	if base.Ui32(v837) < base.Ui32(int32(100)) {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v808 = l1 + v718 + v772 - v792
	v812 = base.I32_div_u_s(v791, int32(_a_F_float_to_shortest_decimal_bufn_9))
	v815 = v791 + v812*int32(-10000)
	v816 = int32(100)
	v817 = base.I32_div_u_s(v815, v816)
	v818 = int32(1)
	v820 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v817<<(uint(v818)%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v808-int32(3)))) = uint16(v820)
	v829 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v815-v817*v816)<<(uint(v818)%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v808-v818))) = uint16(v829)
	v832 = v792 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v791) {
		v791 = v812
		v792 = v832
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v837 = v812
	v838 = v832
	goto L112
L115:
	;
	goto L114
L116:
	;
	v879 = v719 - int32(1)
	v880 = l1 + v718
	if base.Ui32(int32(10)) <= base.Ui32(v877) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v876 = v838
	v877 = v837
	goto L116
L118:
	;
	goto L119
L119:
	;
	v861 = int32(_a_F_float_to_shortest_decimal_bufn_10)
	v863 = int32(100)
	v864 = base.I32_div_u_s(v837&v861, v863)
	v872 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v837-v864*v863)&v861<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v718+v772+(v838^int32(-1))))) = uint16(v872)
	v876 = v838 | int32(2)
	v877 = v864
	goto L116
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v880))) = uint8(v894)
	if base.Ui32(int32(2)) <= base.Ui32(v772) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v887 = v877 << (uint(int32(1)) % 32)
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+uint32(_c_F_float_to_shortest_decimal_bufn[5]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+(v718+v772-v876)))) = uint8(v888)
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	v894 = v890
	goto L120
L122:
	;
	goto L123
L123:
	;
	v894 = v877 | int32(48)
	goto L120
L124:
	;
	v898 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v880)+1)) = uint8(v898)
	v903 = v772 + int32(1)
	goto L126
L125:
	;
	v903 = int32(1)
	goto L126
L126:
	;
	v904 = v903 + v718
	v905 = l1 + v904
	v906 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v905))) = uint8(v906)
	v911 = base.B2i32(v879 < int32(0))
	if v879 < int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v912 = int32(45)
	goto L129
L128:
	;
	v912 = int32(43)
	goto L129
L129:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v905)+1)) = uint8(v912)
	if v879 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v916 = int32(1) - v719
	goto L132
L131:
	;
	v916 = v879
	goto L132
L132:
	;
	v921 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v916<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v905)+2)) = uint16(v921)
	return v904 + int32(4)
L133:
	;
	v938 = v931
	v939 = v695
	goto L136
L134:
	;
	v984 = v931
	v985 = v695
	goto L135
L135:
	;
	if base.Ui32(v985) < base.Ui32(int32(100)) {
		goto L140
	} else {
		goto L141
	}
L136:
	;
	v955 = v929 + v724 + v711 - v938
	v956 = int32(4)
	v959 = base.I32_div_u_s(v939, int32(_a_F_float_to_shortest_decimal_bufn_9))
	v962 = v939 + v959*int32(-10000)
	v963 = int32(100)
	v964 = base.I32_div_u_s(v962, v963)
	v965 = int32(1)
	v967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v964<<(uint(v965)%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v955-v956))) = uint16(v967)
	v976 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v962-v964*v963)<<(uint(v965)%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v955-int32(2)))) = uint16(v976)
	v979 = v938 + v956
	if base.Ui32(int32(99999999)) < base.Ui32(v939) {
		v938 = v979
		v939 = v959
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v984 = v979
	v985 = v959
	goto L135
L138:
	;
	goto L137
L139:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1024) {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	v1023 = v984
	v1024 = v985
	goto L139
L141:
	;
	goto L142
L142:
	;
	v1006 = int32(2)
	v1008 = int32(_a_F_float_to_shortest_decimal_bufn_10)
	v1010 = int32(100)
	v1011 = base.I32_div_u_s(v985&v1008, v1010)
	v1019 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v985-v1011*v1010)&v1008<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v929+v724+v711-v984-v1006))) = uint16(v1019)
	v1023 = v984 | v1006
	v1024 = v1011
	goto L139
L143:
	;
	v1040 = int32(1)
	if v929 == v1040 {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v1034 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1024<<(uint(int32(1))%32))+uint32(_c_F_float_to_shortest_decimal_bufn[4]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v929+v724+v711-v1023-int32(2)))) = uint16(v1034)
	goto L143
L145:
	;
	goto L146
L146:
	;
	v1038 = v1024 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v929+v724))) = uint8(v1038)
	goto L143
L147:
	;
	return v1080 + int32(base.Ui32(v20)>>(uint(int32(31))%32))
L148:
	;
	if v719&int32(4) != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	if v700 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L151:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v724)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v1045
	v1048 = int32(5)
	goto L153
L152:
	;
	v1048 = v1040
	goto L153
L153:
	;
	if v719&int32(2) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1051 = v1048 + v724
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1051))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1051-int32(1)))) = uint16(v1054)
	v1059 = v1048 | int32(2)
	goto L156
L155:
	;
	v1059 = v1048
	goto L156
L156:
	;
	if v719&int32(1) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v1062 = v1059 + v724
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1062-int32(1)))) = uint8(v1065)
	goto L159
L158:
	;
	goto L159
L159:
	;
	v1069 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v724+v719))) = uint8(v1069)
	v1080 = v711 + int32(1)
	goto L147
L160:
	;
	v1077 = int32(2) - v700
	goto L162
L161:
	;
	v1077 = v719
	goto L162
L162:
	;
	v1080 = v1077
	goto L147
}
