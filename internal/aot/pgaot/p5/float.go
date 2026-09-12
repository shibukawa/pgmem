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
			F_errmsg(m, int32(259975), int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errdetail(m, int32(639218), int32(0))
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(516971), int32(3188), int32(230777))
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
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int64
	_ = v186
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int64
	_ = v367
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v398 int64
	_ = v398
	var v402 int64
	_ = v402
	var v404 int32
	_ = v404
	var v407 int64
	_ = v407
	var v409 int32
	_ = v409
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v706 int32
	_ = v706
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
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
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int64
	_ = v938
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	v3 = int32(0)
	v20 = base.I32_reinterpret_f32(l0)
	v22 = v20 & int32(8388607)
	v25 = int32(255)
	v26 = int32(base.Ui32(v20)>>(uint(int32(23))%32)) & v25
	if v22|v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v1124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1122+l1))) = uint8(v1124)
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1314])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v35)
	v38 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1315])))
	*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v38)
	v1122 = int32(3)
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
	v1122 = v54
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
	v1122 = v61
	goto L1
L23:
	;
	v714 = v713 + v706
	v715 = int32(0)
	if v20 < v715 {
		goto L93
	} else {
		goto L94
	}
L24:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v654) {
		v696 = v654
		v706 = v664
		v713 = int32(8)
		goto L23
	} else {
		goto L84
	}
L25:
	;
	v79 = int32(2)
	v85 = v22 << (uint(v79) % 32)
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
	v654 = int32(base.Ui32(v22|int32(8388608)) >> (uint(v68) % 32))
	v664 = v3
	goto L24
L28:
	;
	v88 = v85 | int32(33554432)
	goto L30
L29:
	;
	v88 = v85
	goto L30
L30:
	;
	v89 = base.B2i32(v22 != int32(0)) | base.B2i32(base.Ui32(v26) < base.Ui32(v79)) ^ int32(-1) + v88
	v91 = v88 | int32(2)
	if v26 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v647 = v629 + v635
	v648 = v646 + v631
	if base.Ui32(v648) <= base.Ui32(int32(99999999)) {
		v654 = v648
		v664 = v647
		goto L24
	} else {
		goto L83
	}
L32:
	;
	v565 = int32(0)
	v566 = int32(10)
	v567 = base.I32_div_u_s(v555, v566)
	v569 = base.I32_div_u_s(v557, v566)
	if base.Ui32(v569) < base.Ui32(v567) {
		goto L77
	} else {
		goto L78
	}
L33:
	;
	v469 = int32(0)
	v470 = int32(10)
	v471 = base.I32_div_u_s(v459, v470)
	v473 = base.I32_div_u_s(v461, v470)
	if base.Ui32(v471) <= base.Ui32(v473) {
		goto L71
	} else {
		goto L72
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
	v101 = int32(base.Ui32(v95*int32(78913)) >> (uint(int32(18)) % 32))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v101<<(uint(int32(3))%32))+uint32(_consts[1316])))
	v108 = v106 & int64(4294967295)
	v109 = base.I64_extend_i32_u(v89)
	v111 = int64(32)
	v113 = base.I32_wrap_i64(int64(base.Ui64(v108*v109) >> (uint(v111) % 64)))
	v115 = int64(base.Ui64(v106) >> (uint(v111) % 64))
	v116 = v115 * v109
	v118 = v113 + base.I32_wrap_i64(v116)
	v125 = v101 - v95
	v130 = v125 + int32(base.Ui32(v101*int32(1217359))>>(uint(int32(19))%32))
	v131 = int32(5) - v130
	v134 = v130 + int32(27)
	v136 = (base.B2i32(base.Ui32(v118) < base.Ui32(v113))+base.I32_wrap_i64(int64(base.Ui64(v116)>>(uint(v111)%64))))<<(uint(v131)%32) | int32(base.Ui32(v118)>>(uint(v134)%32))
	v137 = base.I64_extend_i32_u(v91)
	v141 = base.I32_wrap_i64(int64(base.Ui64(v108*v137) >> (uint(v111) % 64)))
	v142 = v137 * v115
	v144 = v141 + base.I32_wrap_i64(v142)
	v152 = (base.B2i32(base.Ui32(v144) < base.Ui32(v141))+base.I32_wrap_i64(int64(base.Ui64(v142)>>(uint(v111)%64))))<<(uint(v131)%32) | int32(base.Ui32(v144)>>(uint(v134)%32))
	v153 = base.I64_extend_i32_u(v88)
	v157 = base.I32_wrap_i64(int64(base.Ui64(v108*v153) >> (uint(v111) % 64)))
	v158 = v153 * v115
	v160 = v157 + base.I32_wrap_i64(v158)
	v168 = (base.B2i32(base.Ui32(v160) < base.Ui32(v157))+base.I32_wrap_i64(int64(base.Ui64(v158)>>(uint(v111)%64))))<<(uint(v131)%32) | int32(base.Ui32(v160)>>(uint(v134)%32))
	v169 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v95) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v314 = v95 * int32(-732923)
	v316 = int32(base.Ui32(v314) >> (uint(int32(20)) % 32))
	v317 = v316 + v95
	v321 = *(*int64)(unsafe.Add(mBase, uint32(int32(1876512)-v317<<(uint(int32(3))%32))))
	v323 = v321 & int64(4294967295)
	v324 = base.I64_extend_i32_u(v89)
	v326 = int64(32)
	v328 = base.I32_wrap_i64(int64(base.Ui64(v323*v324) >> (uint(v326) % 64)))
	v330 = int64(base.Ui64(v321) >> (uint(v326) % 64))
	v331 = v330 * v324
	v333 = v328 + base.I32_wrap_i64(v331)
	v344 = v316 - int32(base.Ui32(v317*int32(-1217359))>>(uint(int32(19))%32))
	v345 = int32(4) - v344
	v348 = v344 + int32(28)
	v350 = (base.B2i32(base.Ui32(v333) < base.Ui32(v328))+base.I32_wrap_i64(int64(base.Ui64(v331)>>(uint(v326)%64))))<<(uint(v345)%32) | int32(base.Ui32(v333)>>(uint(v348)%32))
	v351 = base.I64_extend_i32_u(v88)
	v355 = base.I32_wrap_i64(int64(base.Ui64(v323*v351) >> (uint(v326) % 64)))
	v356 = v351 * v330
	v358 = v355 + base.I32_wrap_i64(v356)
	v366 = (base.B2i32(base.Ui32(v358) < base.Ui32(v355))+base.I32_wrap_i64(int64(base.Ui64(v356)>>(uint(v326)%64))))<<(uint(v345)%32) | int32(base.Ui32(v358)>>(uint(v348)%32))
	v367 = base.I64_extend_i32_u(v91)
	v371 = base.I32_wrap_i64(int64(base.Ui64(v323*v367) >> (uint(v326) % 64)))
	v372 = v330 * v367
	v374 = v371 + base.I32_wrap_i64(v372)
	v382 = (base.B2i32(base.Ui32(v374) < base.Ui32(v371))+base.I32_wrap_i64(int64(base.Ui64(v372)>>(uint(v326)%64))))<<(uint(v345)%32) | int32(base.Ui32(v374)>>(uint(v348)%32))
	v384 = v382 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v314) {
		goto L61
	} else {
		goto L62
	}
L40:
	;
	v175 = int32(10)
	v176 = base.I32_div_u_s(v152-int32(1), v175)
	v178 = base.I32_div_u_s(v136, v175)
	if base.Ui32(v176) <= base.Ui32(v178) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v223 = v169
	goto L42
L42:
	;
	v228 = base.I32_rem_u_s(v88, int32(5))
	if v228 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v181 = v101 - int32(1)
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v181<<(uint(int32(3))%32))+uint32(_consts[1316])))
	v190 = int64(32)
	v192 = base.I32_wrap_i64(int64(base.Ui64(v186&int64(4294967295)*v153) >> (uint(v190) % 64)))
	v195 = int64(base.Ui64(v186)>>(uint(v190)%64)) * v153
	v197 = v192 + base.I32_wrap_i64(v195)
	v208 = v125 + int32(base.Ui32(v181*int32(1217359))>>(uint(int32(19))%32))
	v216 = base.I32_rem_u_s((base.B2i32(base.Ui32(v197) < base.Ui32(v192))+base.I32_wrap_i64(int64(base.Ui64(v195)>>(uint(v190)%64))))<<(uint(int32(6)-v208)%32)|int32(base.Ui32(v197)>>(uint(v208+int32(26))%32)), int32(10))
	v217 = v216
	goto L45
L44:
	;
	v217 = v169
	goto L45
L45:
	;
	if base.Ui32(int32(33)) < base.Ui32(v95) {
		v549 = v168
		v551 = v217
		v554 = v101
		v555 = v152
		v557 = v136
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v223 = v217
	goto L42
L47:
	;
	v233 = v88
	v235 = v169
	goto L50
L48:
	;
	goto L49
L49:
	;
	v259 = int32(0)
	v261 = base.I32_rem_u_s(v91, int32(5))
	if v261 == v259 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v251 = v235 + int32(1)
	v252 = int32(5)
	v253 = base.I32_div_u_s(v233, v252)
	v255 = base.I32_rem_u_s(v253, v252)
	if v255 == int32(0) {
		v233 = v253
		v235 = v251
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if base.Ui32(v251) < base.Ui32(v101) {
		v549 = v168
		v551 = v223
		v554 = v101
		v555 = v152
		v557 = v136
		goto L32
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v453 = v168
	v455 = v223
	v458 = v101
	v459 = v152
	v461 = v136
	goto L33
L54:
	;
	v266 = v259
	v270 = v91
	goto L57
L55:
	;
	v293 = v259
	goto L56
L56:
	;
	v549 = v168
	v551 = v223
	v554 = v101
	v555 = v152 - base.B2i32(base.Ui32(v101) <= base.Ui32(v293))
	v557 = v136
	goto L32
L57:
	;
	v284 = v266 + int32(1)
	v285 = int32(5)
	v286 = base.I32_div_u_s(v270, v285)
	v288 = base.I32_rem_u_s(v286, v285)
	if v288 == int32(0) {
		v266 = v284
		v270 = v286
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v293 = v284
	goto L56
L59:
	;
	goto L58
L60:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v314) {
		v549 = v366
		v551 = v431
		v554 = v317
		v555 = v382
		v557 = v350
		goto L32
	} else {
		goto L68
	}
L61:
	;
	v387 = int32(10)
	v388 = base.I32_div_u_s(v384, v387)
	v390 = base.I32_div_u_s(v350, v387)
	if base.Ui32(v388) <= base.Ui32(v390) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v437 = v3
	goto L63
L63:
	;
	v453 = v366
	v455 = v437
	v458 = v317
	v459 = v384
	v461 = v350
	goto L33
L64:
	;
	v393 = int32(1) - v317
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v393<<(uint(int32(3))%32))+uint32(_consts[1317])))
	v402 = int64(32)
	v404 = base.I32_wrap_i64(int64(base.Ui64(v398&int64(4294967295)*v351) >> (uint(v402) % 64)))
	v407 = int64(base.Ui64(v398)>>(uint(v402)%64)) * v351
	v409 = v404 + base.I32_wrap_i64(v407)
	v422 = v316 + (int32(base.Ui32(v393*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v430 = base.I32_rem_u_s((base.B2i32(base.Ui32(v409) < base.Ui32(v404))+base.I32_wrap_i64(int64(base.Ui64(v407)>>(uint(v402)%64))))<<(uint(int32(4)-v422)%32)|int32(base.Ui32(v409)>>(uint(v422+int32(28))%32)), int32(10))
	v431 = v430
	goto L66
L65:
	;
	v431 = v3
	goto L66
L66:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v314) {
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v437 = v431
	goto L63
L68:
	;
	v443 = int32(-1)
	if v88&(v443<<(uint(v316-int32(1))%32)^v443) != 0 {
		v549 = v366
		v551 = v431
		v554 = v317
		v555 = v382
		v557 = v350
		goto L32
	} else {
		goto L69
	}
L69:
	;
	v453 = v366
	v455 = v431
	v458 = v317
	v459 = v382
	v461 = v350
	goto L33
L70:
	;
	v536 = v521 & int32(255)
	v629 = v517
	v631 = v519
	v635 = v458
	v646 = (v534|base.B2i32(v536 != int32(5))|v519)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v536)) | base.B2i32(v519 == v526)
	goto L31
L71:
	;
	v517 = v469
	v519 = v453
	v521 = v455
	v526 = v461
	v534 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v479 = v469
	v480 = v453
	v482 = v455
	v484 = v471
	v486 = v473
	v487 = int32(1)
	goto L74
L74:
	;
	v497 = v479 + int32(1)
	v498 = int32(10)
	v499 = base.I32_div_u_s(v480, v498)
	v502 = v480 - v499*v498
	v507 = v487 & base.B2i32(v482&int32(255) == int32(0))
	v509 = base.I32_div_u_s(v484, v498)
	v511 = base.I32_div_u_s(v486, v498)
	if base.Ui32(v511) < base.Ui32(v509) {
		v479 = v497
		v480 = v499
		v482 = v502
		v484 = v509
		v486 = v511
		v487 = v507
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v517 = v497
	v519 = v499
	v521 = v502
	v526 = v486
	v534 = v507 ^ int32(1)
	goto L70
L76:
	;
	goto L75
L77:
	;
	v573 = v565
	v574 = v549
	v575 = v567
	v577 = v569
	goto L80
L78:
	;
	v604 = v565
	v605 = v549
	v607 = v551
	v613 = v557
	goto L79
L79:
	;
	v629 = v604
	v631 = v605
	v635 = v554
	v646 = base.B2i32(v605 == v613) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v607&int32(255)))
	goto L31
L80:
	;
	v591 = v573 + int32(1)
	v592 = int32(10)
	v593 = base.I32_div_u_s(v574, v592)
	v595 = base.I32_div_u_s(v575, v592)
	v597 = base.I32_div_u_s(v577, v592)
	if base.Ui32(v597) < base.Ui32(v595) {
		v573 = v591
		v574 = v593
		v575 = v595
		v577 = v597
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v604 = v591
	v605 = v593
	v607 = v574 - v593*int32(10)
	v613 = v577
	goto L79
L82:
	;
	goto L81
L83:
	;
	v696 = v648
	v706 = v647
	v713 = int32(9)
	goto L23
L84:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v654) {
		v696 = v654
		v706 = v664
		v713 = int32(7)
		goto L23
	} else {
		goto L85
	}
L85:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v654) {
		v696 = v654
		v706 = v664
		v713 = int32(6)
		goto L23
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v654) {
		v696 = v654
		v706 = v664
		v713 = int32(5)
		goto L23
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(int32(999)) < base.Ui32(v654) {
		v696 = v654
		v706 = v664
		v713 = int32(4)
		goto L23
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(int32(99)) < base.Ui32(v654) {
		v696 = v654
		v706 = v664
		v713 = int32(3)
		goto L23
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(int32(9)) < base.Ui32(v654) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v693 = int32(2)
	goto L92
L91:
	;
	v693 = int32(1)
	goto L92
L92:
	;
	v696 = v654
	v706 = v664
	v713 = v693
	goto L23
L93:
	;
	v718 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v718)
	v721 = int32(1)
	goto L95
L94:
	;
	v721 = v715
	goto L95
L95:
	;
	if base.Ui32(v714+int32(3)) <= base.Ui32(int32(9)) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v942 = int32(0)
	if base.Ui32(v696) < base.Ui32(int32(10000)) {
		goto L138
	} else {
		goto L139
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v726))) = v938
	v940 = v937
	goto L96
L98:
	;
	v726 = l1 + v721
	v727 = int32(0)
	if v714 <= v727 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	if v706 != 0 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	v937 = int32(2) - v714
	v938 = int64(3472328296227679792)
	goto L97
L102:
	;
	goto L103
L103:
	;
	if int32(0) <= v706 {
		v937 = v727
		v938 = int64(3472328296227680304)
		goto L97
	} else {
		goto L104
	}
L104:
	;
	v940 = int32(1)
	goto L96
L105:
	;
	v786 = int32(0)
	if base.Ui32(v770) < base.Ui32(int32(10000)) {
		goto L114
	} else {
		goto L115
	}
L106:
	;
	v770 = v696
	v776 = v713
	goto L105
L107:
	;
	goto L108
L108:
	;
	v739 = v696
	v744 = v713
	goto L109
L109:
	;
	if v739&int32(1) != 0 {
		v770 = v739
		v776 = v744
		goto L105
	} else {
		goto L111
	}
L110:
	;
	v770 = v739
	v776 = v744
	goto L105
L111:
	;
	v763 = base.I32_div_u_s(v739, int32(10))
	if int32(0)-v739 == v763*int32(-10) {
		v739 = v763
		v744 = v744 - int32(1)
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	if base.Ui32(v845) < base.Ui32(int32(100)) {
		goto L121
	} else {
		goto L122
	}
L114:
	;
	v843 = v786
	v845 = v770
	goto L113
L115:
	;
	goto L116
L116:
	;
	v793 = v786
	v794 = v770
	goto L117
L117:
	;
	v810 = l1 + v721 + v776 - v793
	v814 = base.I32_div_u_s(v794, int32(10000))
	v817 = v814*int32(-10000) + v794
	v818 = int32(100)
	v819 = base.I32_div_u_s(v817, v818)
	v820 = int32(1)
	v824 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v819<<(uint(v820)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v810-int32(3)))) = uint16(v824)
	v835 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v817-v819*v818)<<(uint(v820)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v810-v820))) = uint16(v835)
	v838 = v793 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v794) {
		v793 = v838
		v794 = v814
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v843 = v838
	v845 = v814
	goto L113
L119:
	;
	goto L118
L120:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v885) {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	v884 = v843
	v885 = v845
	goto L120
L122:
	;
	goto L123
L123:
	;
	v867 = int32(65535)
	v869 = int32(100)
	v870 = base.I32_div_u_s(v845&v867, v869)
	v880 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v845-v870*v869)&v867<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v721+v776+(v843^int32(-1))))) = uint16(v880)
	v884 = v843 | int32(2)
	v885 = v870
	goto L120
L124:
	;
	v904 = int32(1)
	v905 = v714 - v904
	v906 = l1 + v721
	*(*uint8)(unsafe.Add(mBase, uint32(v906))) = uint8(v903)
	if base.Ui32(int32(2)) <= base.Ui32(v776) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v892 = v885 << (uint(int32(1)) % 32)
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892)+uint32(_consts[1319]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+(v721+v776-v884)))) = uint8(v895)
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892)+uint32(_consts[1318]))))
	v903 = v899
	goto L124
L126:
	;
	goto L127
L127:
	;
	v903 = v885 | int32(48)
	goto L124
L128:
	;
	v911 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v906)+1)) = uint8(v911)
	v915 = v776 + int32(1)
	goto L130
L129:
	;
	v915 = v904
	goto L130
L130:
	;
	v916 = v915 + v721
	v917 = l1 + v916
	v918 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v917))) = uint8(v918)
	v923 = base.B2i32(v905 < int32(0))
	if v905 < int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v924 = int32(45)
	goto L133
L132:
	;
	v924 = int32(43)
	goto L133
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v917)+1)) = uint8(v924)
	if v905 < int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v928 = int32(1) - v714
	goto L136
L135:
	;
	v928 = v905
	goto L136
L136:
	;
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v928<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v917)+2)) = uint16(v933)
	v1122 = v916 + int32(4)
	goto L1
L137:
	;
	if base.Ui32(v1001) < base.Ui32(int32(100)) {
		goto L145
	} else {
		goto L146
	}
L138:
	;
	v1000 = v942
	v1001 = v696
	goto L137
L139:
	;
	goto L140
L140:
	;
	v949 = v696
	v950 = v942
	goto L141
L141:
	;
	v966 = v726 + v940 + v713 - v950
	v967 = int32(4)
	v970 = base.I32_div_u_s(v949, int32(10000))
	v973 = v970*int32(-10000) + v949
	v974 = int32(100)
	v975 = base.I32_div_u_s(v973, v974)
	v976 = int32(1)
	v980 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v975<<(uint(v976)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v966-v967))) = uint16(v980)
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v973-v975*v974)<<(uint(v976)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v966-int32(2)))) = uint16(v991)
	v994 = v950 + v967
	if base.Ui32(int32(99999999)) < base.Ui32(v949) {
		v949 = v970
		v950 = v994
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v1000 = v994
	v1001 = v970
	goto L137
L143:
	;
	goto L142
L144:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1040) {
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v1040 = v1001
	v1041 = v1000
	goto L144
L146:
	;
	goto L147
L147:
	;
	v1021 = int32(2)
	v1023 = int32(65535)
	v1025 = int32(100)
	v1026 = base.I32_div_u_s(v1001&v1023, v1025)
	v1036 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1001-v1026*v1025)&v1023<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v726+v940+v713-v1000-v1021))) = uint16(v1036)
	v1040 = v1026
	v1041 = v1000 | v1021
	goto L144
L148:
	;
	v1059 = int32(1)
	if v940 == v1059 {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v1053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1040<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v726+v940+v713-v1041-int32(2)))) = uint16(v1053)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v1057 = v1040 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v726+v940))) = uint8(v1057)
	goto L148
L152:
	;
	v1122 = v1099 + int32(base.Ui32(v20)>>(uint(int32(31))%32))
	goto L1
L153:
	;
	if v714&int32(4) != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	if v706 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L156:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v726)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v726))) = v1064
	v1067 = int32(5)
	goto L158
L157:
	;
	v1067 = v1059
	goto L158
L158:
	;
	if v714&int32(2) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1070 = v726 + v1067
	v1073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1070))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1070-int32(1)))) = uint16(v1073)
	v1078 = v1067 | int32(2)
	goto L161
L160:
	;
	v1078 = v1067
	goto L161
L161:
	;
	if v714&int32(1) != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v1081 = v726 + v1078
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1081))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1081-int32(1)))) = uint8(v1084)
	goto L164
L163:
	;
	goto L164
L164:
	;
	v1088 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v726+v714))) = uint8(v1088)
	v1099 = v713 + int32(1)
	goto L152
L165:
	;
	v1096 = int32(2) - v706
	goto L167
L166:
	;
	v1096 = v714
	goto L167
L167:
	;
	v1099 = v1096
	goto L152
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
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v358 int32
	_ = v358
	var v359 int64
	_ = v359
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v370 int64
	_ = v370
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v401 int64
	_ = v401
	var v405 int64
	_ = v405
	var v407 int32
	_ = v407
	var v410 int64
	_ = v410
	var v412 int32
	_ = v412
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v667 int32
	_ = v667
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int64
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	v3 = int32(0)
	v20 = base.I32_reinterpret_f32(l0)
	v22 = v20 & int32(8388607)
	v25 = int32(255)
	v26 = int32(base.Ui32(v20)>>(uint(int32(23))%32)) & v25
	if v22|v26 != 0 {
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1314])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v35)
	v38 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1315])))
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
	v717 = v716 + v709
	v718 = int32(0)
	if v20 < v718 {
		goto L92
	} else {
		goto L93
	}
L23:
	;
	if base.Ui32(int32(9999999)) < base.Ui32(v657) {
		v699 = v657
		v709 = v667
		v716 = int32(8)
		goto L22
	} else {
		goto L83
	}
L24:
	;
	v82 = int32(2)
	v88 = v22 << (uint(v82) % 32)
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
	v657 = int32(base.Ui32(v22|int32(8388608)) >> (uint(v71) % 32))
	v667 = v3
	goto L23
L27:
	;
	v91 = v88 | int32(33554432)
	goto L29
L28:
	;
	v91 = v88
	goto L29
L29:
	;
	v92 = base.B2i32(v22 != int32(0)) | base.B2i32(base.Ui32(v26) < base.Ui32(v82)) ^ int32(-1) + v91
	v94 = v91 | int32(2)
	if v26 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v650 = v632 + v638
	v651 = v649 + v634
	if base.Ui32(v651) <= base.Ui32(int32(99999999)) {
		v657 = v651
		v667 = v650
		goto L23
	} else {
		goto L82
	}
L31:
	;
	v568 = int32(0)
	v569 = int32(10)
	v570 = base.I32_div_u_s(v558, v569)
	v572 = base.I32_div_u_s(v560, v569)
	if base.Ui32(v572) < base.Ui32(v570) {
		goto L76
	} else {
		goto L77
	}
L32:
	;
	v472 = int32(0)
	v473 = int32(10)
	v474 = base.I32_div_u_s(v462, v473)
	v476 = base.I32_div_u_s(v464, v473)
	if base.Ui32(v474) <= base.Ui32(v476) {
		goto L70
	} else {
		goto L71
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
	v104 = int32(base.Ui32(v98*int32(78913)) >> (uint(int32(18)) % 32))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v104<<(uint(int32(3))%32))+uint32(_consts[1316])))
	v111 = v109 & int64(4294967295)
	v112 = base.I64_extend_i32_u(v92)
	v114 = int64(32)
	v116 = base.I32_wrap_i64(int64(base.Ui64(v111*v112) >> (uint(v114) % 64)))
	v118 = int64(base.Ui64(v109) >> (uint(v114) % 64))
	v119 = v118 * v112
	v121 = v116 + base.I32_wrap_i64(v119)
	v128 = v104 - v98
	v133 = v128 + int32(base.Ui32(v104*int32(1217359))>>(uint(int32(19))%32))
	v134 = int32(5) - v133
	v137 = v133 + int32(27)
	v139 = (base.B2i32(base.Ui32(v121) < base.Ui32(v116))+base.I32_wrap_i64(int64(base.Ui64(v119)>>(uint(v114)%64))))<<(uint(v134)%32) | int32(base.Ui32(v121)>>(uint(v137)%32))
	v140 = base.I64_extend_i32_u(v94)
	v144 = base.I32_wrap_i64(int64(base.Ui64(v111*v140) >> (uint(v114) % 64)))
	v145 = v140 * v118
	v147 = v144 + base.I32_wrap_i64(v145)
	v155 = (base.B2i32(base.Ui32(v147) < base.Ui32(v144))+base.I32_wrap_i64(int64(base.Ui64(v145)>>(uint(v114)%64))))<<(uint(v134)%32) | int32(base.Ui32(v147)>>(uint(v137)%32))
	v156 = base.I64_extend_i32_u(v91)
	v160 = base.I32_wrap_i64(int64(base.Ui64(v111*v156) >> (uint(v114) % 64)))
	v161 = v156 * v118
	v163 = v160 + base.I32_wrap_i64(v161)
	v171 = (base.B2i32(base.Ui32(v163) < base.Ui32(v160))+base.I32_wrap_i64(int64(base.Ui64(v161)>>(uint(v114)%64))))<<(uint(v134)%32) | int32(base.Ui32(v163)>>(uint(v137)%32))
	v172 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v98) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v317 = v98 * int32(-732923)
	v319 = int32(base.Ui32(v317) >> (uint(int32(20)) % 32))
	v320 = v319 + v98
	v324 = *(*int64)(unsafe.Add(mBase, uint32(int32(1876512)-v320<<(uint(int32(3))%32))))
	v326 = v324 & int64(4294967295)
	v327 = base.I64_extend_i32_u(v92)
	v329 = int64(32)
	v331 = base.I32_wrap_i64(int64(base.Ui64(v326*v327) >> (uint(v329) % 64)))
	v333 = int64(base.Ui64(v324) >> (uint(v329) % 64))
	v334 = v333 * v327
	v336 = v331 + base.I32_wrap_i64(v334)
	v347 = v319 - int32(base.Ui32(v320*int32(-1217359))>>(uint(int32(19))%32))
	v348 = int32(4) - v347
	v351 = v347 + int32(28)
	v353 = (base.B2i32(base.Ui32(v336) < base.Ui32(v331))+base.I32_wrap_i64(int64(base.Ui64(v334)>>(uint(v329)%64))))<<(uint(v348)%32) | int32(base.Ui32(v336)>>(uint(v351)%32))
	v354 = base.I64_extend_i32_u(v91)
	v358 = base.I32_wrap_i64(int64(base.Ui64(v326*v354) >> (uint(v329) % 64)))
	v359 = v354 * v333
	v361 = v358 + base.I32_wrap_i64(v359)
	v369 = (base.B2i32(base.Ui32(v361) < base.Ui32(v358))+base.I32_wrap_i64(int64(base.Ui64(v359)>>(uint(v329)%64))))<<(uint(v348)%32) | int32(base.Ui32(v361)>>(uint(v351)%32))
	v370 = base.I64_extend_i32_u(v94)
	v374 = base.I32_wrap_i64(int64(base.Ui64(v326*v370) >> (uint(v329) % 64)))
	v375 = v333 * v370
	v377 = v374 + base.I32_wrap_i64(v375)
	v385 = (base.B2i32(base.Ui32(v377) < base.Ui32(v374))+base.I32_wrap_i64(int64(base.Ui64(v375)>>(uint(v329)%64))))<<(uint(v348)%32) | int32(base.Ui32(v377)>>(uint(v351)%32))
	v387 = v385 - int32(1)
	if base.Ui32(int32(1048576)) <= base.Ui32(v317) {
		goto L60
	} else {
		goto L61
	}
L39:
	;
	v178 = int32(10)
	v179 = base.I32_div_u_s(v155-int32(1), v178)
	v181 = base.I32_div_u_s(v139, v178)
	if base.Ui32(v179) <= base.Ui32(v181) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v226 = v172
	goto L41
L41:
	;
	v231 = base.I32_rem_u_s(v91, int32(5))
	if v231 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v184 = v104 - int32(1)
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v184<<(uint(int32(3))%32))+uint32(_consts[1316])))
	v193 = int64(32)
	v195 = base.I32_wrap_i64(int64(base.Ui64(v189&int64(4294967295)*v156) >> (uint(v193) % 64)))
	v198 = int64(base.Ui64(v189)>>(uint(v193)%64)) * v156
	v200 = v195 + base.I32_wrap_i64(v198)
	v211 = v128 + int32(base.Ui32(v184*int32(1217359))>>(uint(int32(19))%32))
	v219 = base.I32_rem_u_s((base.B2i32(base.Ui32(v200) < base.Ui32(v195))+base.I32_wrap_i64(int64(base.Ui64(v198)>>(uint(v193)%64))))<<(uint(int32(6)-v211)%32)|int32(base.Ui32(v200)>>(uint(v211+int32(26))%32)), int32(10))
	v220 = v219
	goto L44
L43:
	;
	v220 = v172
	goto L44
L44:
	;
	if base.Ui32(int32(33)) < base.Ui32(v98) {
		v552 = v171
		v554 = v220
		v557 = v104
		v558 = v155
		v560 = v139
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v226 = v220
	goto L41
L46:
	;
	v236 = v91
	v238 = v172
	goto L49
L47:
	;
	goto L48
L48:
	;
	v262 = int32(0)
	v264 = base.I32_rem_u_s(v94, int32(5))
	if v264 == v262 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v254 = v238 + int32(1)
	v255 = int32(5)
	v256 = base.I32_div_u_s(v236, v255)
	v258 = base.I32_rem_u_s(v256, v255)
	if v258 == int32(0) {
		v236 = v256
		v238 = v254
		goto L49
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(v254) < base.Ui32(v104) {
		v552 = v171
		v554 = v226
		v557 = v104
		v558 = v155
		v560 = v139
		goto L31
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v456 = v171
	v458 = v226
	v461 = v104
	v462 = v155
	v464 = v139
	goto L32
L53:
	;
	v269 = v262
	v273 = v94
	goto L56
L54:
	;
	v296 = v262
	goto L55
L55:
	;
	v552 = v171
	v554 = v226
	v557 = v104
	v558 = v155 - base.B2i32(base.Ui32(v104) <= base.Ui32(v296))
	v560 = v139
	goto L31
L56:
	;
	v287 = v269 + int32(1)
	v288 = int32(5)
	v289 = base.I32_div_u_s(v273, v288)
	v291 = base.I32_rem_u_s(v289, v288)
	if v291 == int32(0) {
		v269 = v287
		v273 = v289
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v296 = v287
	goto L55
L58:
	;
	goto L57
L59:
	;
	if base.Ui32(int32(32505855)) < base.Ui32(v317) {
		v552 = v369
		v554 = v434
		v557 = v320
		v558 = v385
		v560 = v353
		goto L31
	} else {
		goto L67
	}
L60:
	;
	v390 = int32(10)
	v391 = base.I32_div_u_s(v387, v390)
	v393 = base.I32_div_u_s(v353, v390)
	if base.Ui32(v391) <= base.Ui32(v393) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v440 = v3
	goto L62
L62:
	;
	v456 = v369
	v458 = v440
	v461 = v320
	v462 = v387
	v464 = v353
	goto L32
L63:
	;
	v396 = int32(1) - v320
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v396<<(uint(int32(3))%32))+uint32(_consts[1317])))
	v405 = int64(32)
	v407 = base.I32_wrap_i64(int64(base.Ui64(v401&int64(4294967295)*v354) >> (uint(v405) % 64)))
	v410 = int64(base.Ui64(v401)>>(uint(v405)%64)) * v354
	v412 = v407 + base.I32_wrap_i64(v410)
	v425 = v319 + (int32(base.Ui32(v396*int32(1217359))>>(uint(int32(19))%32)) ^ int32(-1))
	v433 = base.I32_rem_u_s((base.B2i32(base.Ui32(v412) < base.Ui32(v407))+base.I32_wrap_i64(int64(base.Ui64(v410)>>(uint(v405)%64))))<<(uint(int32(4)-v425)%32)|int32(base.Ui32(v412)>>(uint(v425+int32(28))%32)), int32(10))
	v434 = v433
	goto L65
L64:
	;
	v434 = v3
	goto L65
L65:
	;
	if base.Ui32(int32(2097151)) < base.Ui32(v317) {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v440 = v434
	goto L62
L67:
	;
	v446 = int32(-1)
	if v91&(v446<<(uint(v319-int32(1))%32)^v446) != 0 {
		v552 = v369
		v554 = v434
		v557 = v320
		v558 = v385
		v560 = v353
		goto L31
	} else {
		goto L68
	}
L68:
	;
	v456 = v369
	v458 = v434
	v461 = v320
	v462 = v385
	v464 = v353
	goto L32
L69:
	;
	v539 = v524 & int32(255)
	v632 = v520
	v634 = v522
	v638 = v461
	v649 = (v537|base.B2i32(v539 != int32(5))|v522)&base.B2i32(base.Ui32(int32(4)) < base.Ui32(v539)) | base.B2i32(v522 == v529)
	goto L30
L70:
	;
	v520 = v472
	v522 = v456
	v524 = v458
	v529 = v464
	v537 = int32(0)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v482 = v472
	v483 = v456
	v485 = v458
	v487 = v474
	v489 = v476
	v490 = int32(1)
	goto L73
L73:
	;
	v500 = v482 + int32(1)
	v501 = int32(10)
	v502 = base.I32_div_u_s(v483, v501)
	v505 = v483 - v502*v501
	v510 = v490 & base.B2i32(v485&int32(255) == int32(0))
	v512 = base.I32_div_u_s(v487, v501)
	v514 = base.I32_div_u_s(v489, v501)
	if base.Ui32(v514) < base.Ui32(v512) {
		v482 = v500
		v483 = v502
		v485 = v505
		v487 = v512
		v489 = v514
		v490 = v510
		goto L73
	} else {
		goto L75
	}
L74:
	;
	v520 = v500
	v522 = v502
	v524 = v505
	v529 = v489
	v537 = v510 ^ int32(1)
	goto L69
L75:
	;
	goto L74
L76:
	;
	v576 = v568
	v577 = v552
	v578 = v570
	v580 = v572
	goto L79
L77:
	;
	v607 = v568
	v608 = v552
	v610 = v554
	v616 = v560
	goto L78
L78:
	;
	v632 = v607
	v634 = v608
	v638 = v557
	v649 = base.B2i32(v608 == v616) | base.B2i32(base.Ui32(int32(4)) < base.Ui32(v610&int32(255)))
	goto L30
L79:
	;
	v594 = v576 + int32(1)
	v595 = int32(10)
	v596 = base.I32_div_u_s(v577, v595)
	v598 = base.I32_div_u_s(v578, v595)
	v600 = base.I32_div_u_s(v580, v595)
	if base.Ui32(v600) < base.Ui32(v598) {
		v576 = v594
		v577 = v596
		v578 = v598
		v580 = v600
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v607 = v594
	v608 = v596
	v610 = v577 - v596*int32(10)
	v616 = v580
	goto L78
L81:
	;
	goto L80
L82:
	;
	v699 = v651
	v709 = v650
	v716 = int32(9)
	goto L22
L83:
	;
	if base.Ui32(int32(999999)) < base.Ui32(v657) {
		v699 = v657
		v709 = v667
		v716 = int32(7)
		goto L22
	} else {
		goto L84
	}
L84:
	;
	if base.Ui32(int32(99999)) < base.Ui32(v657) {
		v699 = v657
		v709 = v667
		v716 = int32(6)
		goto L22
	} else {
		goto L85
	}
L85:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v657) {
		v699 = v657
		v709 = v667
		v716 = int32(5)
		goto L22
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(int32(999)) < base.Ui32(v657) {
		v699 = v657
		v709 = v667
		v716 = int32(4)
		goto L22
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(int32(99)) < base.Ui32(v657) {
		v699 = v657
		v709 = v667
		v716 = int32(3)
		goto L22
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(int32(9)) < base.Ui32(v657) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v696 = int32(2)
	goto L91
L90:
	;
	v696 = int32(1)
	goto L91
L91:
	;
	v699 = v657
	v709 = v667
	v716 = v696
	goto L22
L92:
	;
	v721 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v721)
	v724 = int32(1)
	goto L94
L93:
	;
	v724 = v718
	goto L94
L94:
	;
	if base.Ui32(v717+int32(3)) <= base.Ui32(int32(9)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v946 = int32(0)
	if base.Ui32(v699) < base.Ui32(int32(10000)) {
		goto L137
	} else {
		goto L138
	}
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v729))) = v942
	v944 = v941
	goto L95
L97:
	;
	v729 = l1 + v724
	v730 = int32(0)
	if v717 <= v730 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	if v709 != 0 {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	v941 = int32(2) - v717
	v942 = int64(3472328296227679792)
	goto L96
L101:
	;
	goto L102
L102:
	;
	if int32(0) <= v709 {
		v941 = v730
		v942 = int64(3472328296227680304)
		goto L96
	} else {
		goto L103
	}
L103:
	;
	v944 = int32(1)
	goto L95
L104:
	;
	v789 = int32(0)
	if base.Ui32(v773) < base.Ui32(int32(10000)) {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v773 = v699
	v779 = v716
	goto L104
L106:
	;
	goto L107
L107:
	;
	v742 = v699
	v747 = v716
	goto L108
L108:
	;
	if v742&int32(1) != 0 {
		v773 = v742
		v779 = v747
		goto L104
	} else {
		goto L110
	}
L109:
	;
	v773 = v742
	v779 = v747
	goto L104
L110:
	;
	v766 = base.I32_div_u_s(v742, int32(10))
	if int32(0)-v742 == v766*int32(-10) {
		v742 = v766
		v747 = v747 - int32(1)
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	if base.Ui32(v848) < base.Ui32(int32(100)) {
		goto L120
	} else {
		goto L121
	}
L113:
	;
	v846 = v789
	v848 = v773
	goto L112
L114:
	;
	goto L115
L115:
	;
	v796 = v789
	v797 = v773
	goto L116
L116:
	;
	v813 = l1 + v724 + v779 - v796
	v817 = base.I32_div_u_s(v797, int32(10000))
	v820 = v817*int32(-10000) + v797
	v821 = int32(100)
	v822 = base.I32_div_u_s(v820, v821)
	v823 = int32(1)
	v827 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v822<<(uint(v823)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v813-int32(3)))) = uint16(v827)
	v838 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v820-v822*v821)<<(uint(v823)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v813-v823))) = uint16(v838)
	v841 = v796 + int32(4)
	if base.Ui32(int32(99999999)) < base.Ui32(v797) {
		v796 = v841
		v797 = v817
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v846 = v841
	v848 = v817
	goto L112
L118:
	;
	goto L117
L119:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v888) {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	v887 = v846
	v888 = v848
	goto L119
L121:
	;
	goto L122
L122:
	;
	v870 = int32(65535)
	v872 = int32(100)
	v873 = base.I32_div_u_s(v848&v870, v872)
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v848-v873*v872)&v870<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(l1+v724+v779+(v846^int32(-1))))) = uint16(v883)
	v887 = v846 | int32(2)
	v888 = v873
	goto L119
L123:
	;
	v907 = int32(1)
	v908 = v717 - v907
	v909 = l1 + v724
	*(*uint8)(unsafe.Add(mBase, uint32(v909))) = uint8(v906)
	if base.Ui32(int32(2)) <= base.Ui32(v779) {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v895 = v888 << (uint(int32(1)) % 32)
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+uint32(_consts[1319]))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1+(v724+v779-v887)))) = uint8(v898)
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v895)+uint32(_consts[1318]))))
	v906 = v902
	goto L123
L125:
	;
	goto L126
L126:
	;
	v906 = v888 | int32(48)
	goto L123
L127:
	;
	v914 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v909)+1)) = uint8(v914)
	v918 = v779 + int32(1)
	goto L129
L128:
	;
	v918 = v907
	goto L129
L129:
	;
	v919 = v918 + v724
	v920 = l1 + v919
	v921 = int32(101)
	*(*uint8)(unsafe.Add(mBase, uint32(v920))) = uint8(v921)
	v926 = base.B2i32(v908 < int32(0))
	if v908 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v927 = int32(45)
	goto L132
L131:
	;
	v927 = int32(43)
	goto L132
L132:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v920)+1)) = uint8(v927)
	if v908 < int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v931 = int32(1) - v717
	goto L135
L134:
	;
	v931 = v908
	goto L135
L135:
	;
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v931<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v920)+2)) = uint16(v936)
	return v919 + int32(4)
L136:
	;
	if base.Ui32(v1005) < base.Ui32(int32(100)) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	v1004 = v946
	v1005 = v699
	goto L136
L138:
	;
	goto L139
L139:
	;
	v953 = v699
	v954 = v946
	goto L140
L140:
	;
	v970 = v729 + v944 + v716 - v954
	v971 = int32(4)
	v974 = base.I32_div_u_s(v953, int32(10000))
	v977 = v974*int32(-10000) + v953
	v978 = int32(100)
	v979 = base.I32_div_u_s(v977, v978)
	v980 = int32(1)
	v984 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v979<<(uint(v980)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v970-v971))) = uint16(v984)
	v995 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v977-v979*v978)<<(uint(v980)%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v970-int32(2)))) = uint16(v995)
	v998 = v954 + v971
	if base.Ui32(int32(99999999)) < base.Ui32(v953) {
		v953 = v974
		v954 = v998
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v1004 = v998
	v1005 = v974
	goto L136
L142:
	;
	goto L141
L143:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1044) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	v1044 = v1005
	v1045 = v1004
	goto L143
L145:
	;
	goto L146
L146:
	;
	v1025 = int32(2)
	v1027 = int32(65535)
	v1029 = int32(100)
	v1030 = base.I32_div_u_s(v1005&v1027, v1029)
	v1040 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v1005-v1030*v1029)&v1027<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v729+v944+v716-v1004-v1025))) = uint16(v1040)
	v1044 = v1030
	v1045 = v1004 | v1025
	goto L143
L147:
	;
	v1063 = int32(1)
	if v944 == v1063 {
		goto L152
	} else {
		goto L153
	}
L148:
	;
	v1057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1044<<(uint(int32(1))%32))+uint32(_consts[1318]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v729+v944+v716-v1045-int32(2)))) = uint16(v1057)
	goto L147
L149:
	;
	goto L150
L150:
	;
	v1061 = v1044 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v729+v944))) = uint8(v1061)
	goto L147
L151:
	;
	return v1103 + int32(base.Ui32(v20)>>(uint(int32(31))%32))
L152:
	;
	if v717&int32(4) != 0 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	if v709 < int32(0) {
		goto L164
	} else {
		goto L165
	}
L155:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v729)+1))
	*(*int32)(unsafe.Add(mBase, uint32(v729))) = v1068
	v1071 = int32(5)
	goto L157
L156:
	;
	v1071 = v1063
	goto L157
L157:
	;
	if v717&int32(2) != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1074 = v729 + v1071
	v1077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1074))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1074-int32(1)))) = uint16(v1077)
	v1082 = v1071 | int32(2)
	goto L160
L159:
	;
	v1082 = v1071
	goto L160
L160:
	;
	if v717&int32(1) != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1085 = v729 + v1082
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1085-int32(1)))) = uint8(v1088)
	goto L163
L162:
	;
	goto L163
L163:
	;
	v1092 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v729+v717))) = uint8(v1092)
	v1103 = v716 + int32(1)
	goto L151
L164:
	;
	v1100 = int32(2) - v709
	goto L166
L165:
	;
	v1100 = v717
	goto L166
L166:
	;
	v1103 = v1100
	goto L151
}
