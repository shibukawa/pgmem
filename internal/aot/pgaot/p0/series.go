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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
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
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v566 int32
	_ = v566
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v763 int32
	_ = v763
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v937 int32
	_ = v937
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v973 int64
	_ = v973
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
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
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == v2 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L9
	} else {
		goto L308
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L9
	} else {
		goto L304
	}
L3:
	;
	F_errmsg(m, int32(527458), int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L9
	} else {
		goto L302
	}
L4:
	;
	F_errmsg(m, int32(527508), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L9
	} else {
		goto L300
	}
L5:
	;
	F_errmsg(m, int32(527482), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L9
	} else {
		goto L298
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	goto L71
L9:
	;
	return int32(0)
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v24) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v45) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v24 == int32(49152) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(11323), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(499844), int32(1731), int32(490469))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
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
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v68 == int32(3) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	if v45 == int32(49152) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(11354), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(499844), int32(1742), int32(490469))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v72 = F_pg_detoast_datum(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L31
	}
L29:
	;
	v143 = int32(1737828)
	v144 = v2
	v146 = int32(1)
	v147 = v2
	v148 = v2
	goto L30
L30:
	;
	v149 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L56
	}
L31:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	v75 = base.I32_extend16_s(v74)
	if base.Ui32(int32(49152)) <= base.Ui32(v74) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v102 = base.B2i32(int32(0) <= v75)
	if int32(0) <= v75 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v75 == int32(-16384) {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(11294), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(499844), int32(1759), int32(490469))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
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
	v103 = int32(-8)
	goto L42
L41:
	;
	v103 = int32(-6)
	goto L42
L42:
	;
	v104 = int32(base.Ui32(v96)>>(uint(int32(2))%32)) + v103
	if int32(0) <= v75 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+6)))
	v115 = v105
	goto L45
L44:
	;
	v115 = v74<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v74&int32(63)
	goto L45
L45:
	;
	if base.Ui32(v104) < base.Ui32(int32(2)) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v118 = int32(1)
	v125 = v74 & int32(49152)
	if v125 == int32(32768) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v128 = v74 << (uint(v118) % 32) & int32(16384)
	goto L49
L48:
	;
	v128 = v125
	goto L49
L49:
	;
	v136 = base.B2i32(v75 < int32(0))
	if v75 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v137 = int32(base.Ui32(v74)>>(uint(int32(7))%32)) & int32(63)
	goto L52
L51:
	;
	v137 = v74 & int32(16383)
	goto L52
L52:
	;
	if v75 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v140 = int32(6)
	goto L55
L54:
	;
	v140 = int32(8)
	goto L55
L55:
	;
	v143 = v72 + v140
	v144 = v137
	v146 = int32(base.Ui32(v104) >> (uint(v118) % 32))
	v147 = v128
	v148 = v115
	goto L30
L56:
	;
	v151 = int32(4515248)
	v152 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v154
	v157 = F_palloc(m, int32(72))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v162 = F__emscripten_memset_bulkmem(m, v157, base.I32_extend8_s(int32(0)), int32(72))
	mBase = m.M
	goto L58
L58:
	;
	F_set_var_from_num(m, v17, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	F_set_var_from_num(m, v22, v162+int32(24))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v170 = v146 << (uint(int32(1)) % 32)
	v173 = F_palloc(m, v170+int32(2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v175 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v173))) = uint16(v175)
	v178 = v173 + int32(2)
	if v170 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v162)+64))
	if v181 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v179 = F__emscripten_memcpy_bulkmem(m, v178, v143, v170)
	mBase = m.M
	v180 = v179
	goto L65
L64:
	;
	v180 = v178
	goto L65
L65:
	;
	goto L62
L66:
	;
	F_pfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+68)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v162)+64)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v162)+60)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v162)+56)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v162)+52)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v162)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v149)+16)) = v162
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v152
	goto L8
L69:
	;
	goto L68
L70:
	;
	v960 = F_make_result_opt_error(m, v205, int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L9
	} else {
		goto L296
	}
L71:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+56))
	if v206 != int32(16384) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v206 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v205)+32))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v205)+24))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v582 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L75:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205)+32))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205)+24))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v211 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v210 == int32(0) {
		goto L70
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v210 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	if v209 == int32(16384) {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L70
L81:
	;
	if v218 == int32(0) {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v205)+28))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v205)+44))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v218 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L70
L85:
	;
	if int32(0) < v577 {
		goto L1
	} else {
		goto L185
	}
L86:
	;
	if v209 == int32(16384) {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v209 == int32(0) {
		goto L70
	} else {
		goto L137
	}
L89:
	;
	v231 = int32(0)
	if base.B2i32(v223 < v225)&base.B2i32(v231 < v211) == v231 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v577 = v402
	goto L85
L91:
	;
	v402 = v392
	goto L90
L92:
	;
	if v223 <= v262 {
		v297 = v223
		v299 = v231
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v262 = v225
	v266 = v231
	goto L92
L94:
	;
	goto L95
L95:
	;
	v243 = v225
	v247 = v231
	goto L96
L96:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226+v247<<(uint(int32(1))%32)))))
	if v253 != 0 {
		v392 = int32(1)
		goto L91
	} else {
		goto L98
	}
L97:
	;
	v262 = v257
	v266 = v255
	goto L92
L98:
	;
	v254 = int32(1)
	v255 = v247 + v254
	v257 = v243 - v254
	if v257 <= v223 {
		v262 = v257
		v266 = v255
		goto L92
	} else {
		goto L99
	}
L99:
	;
	if v255 < v211 {
		v243 = v257
		v247 = v255
		goto L96
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	if v262 != v297 {
		v338 = v266
		v339 = v299
		goto L109
	} else {
		goto L110
	}
L102:
	;
	if v210 <= int32(0) {
		v297 = v223
		v299 = v231
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v278 = v223
	v280 = v231
	goto L104
L104:
	;
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224+v280<<(uint(int32(1))%32)))))
	if v285 != 0 {
		v392 = int32(-1)
		goto L91
	} else {
		goto L106
	}
L105:
	;
	v297 = v289
	v299 = v287
	goto L101
L106:
	;
	v286 = int32(1)
	v287 = v280 + v286
	v289 = v278 - v286
	if v289 <= v262 {
		v297 = v289
		v299 = v287
		goto L101
	} else {
		goto L107
	}
L107:
	;
	if v287 < v210 {
		v278 = v289
		v280 = v287
		goto L104
	} else {
		goto L108
	}
L108:
	;
	goto L105
L109:
	;
	if v211 < v338 {
		goto L119
	} else {
		goto L120
	}
L110:
	;
	v308 = v266
	v309 = v299
	goto L111
L111:
	;
	if v211 <= v308 {
		v338 = v308
		v339 = v309
		goto L109
	} else {
		goto L113
	}
L112:
	;
	if base.I32_extend16_s(v324) < base.I32_extend16_s(v322) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	if v210 <= v309 {
		v338 = v308
		v339 = v309
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v313 = int32(1)
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226+v308<<(uint(v313)%32)))))
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309<<(uint(v313)%32)+v224))))
	if v322 == v324 {
		v308 = v308 + v313
		v309 = v309 + v313
		goto L111
	} else {
		goto L115
	}
L115:
	;
	goto L112
L116:
	;
	v331 = int32(1)
	goto L118
L117:
	;
	v331 = int32(-1)
	goto L118
L118:
	;
	v402 = v331
	goto L90
L119:
	;
	v342 = v338
	goto L121
L120:
	;
	v342 = v211
	goto L121
L121:
	;
	v349 = v338
	goto L122
L122:
	;
	if v342 == v349 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v392 = v375
	goto L91
L124:
	;
	if v210 < v339 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	goto L126
L126:
	;
	v375 = int32(1)
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226+v349<<(uint(v375)%32)))))
	if v381 == int32(0) {
		v349 = v349 + v375
		goto L122
	} else {
		goto L136
	}
L127:
	;
	v354 = v339
	goto L129
L128:
	;
	v354 = v210
	goto L129
L129:
	;
	v362 = v339
	goto L130
L130:
	;
	if v354 == v362 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v392 = int32(-1)
	goto L91
L132:
	;
	v402 = int32(0)
	goto L90
L133:
	;
	goto L134
L134:
	;
	v366 = int32(1)
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224+v362<<(uint(v366)%32)))))
	if v371 == int32(0) {
		v362 = v362 + v366
		goto L130
	} else {
		goto L135
	}
L135:
	;
	goto L131
L136:
	;
	goto L123
L137:
	;
	v405 = int32(0)
	if base.B2i32(v225 < v223)&base.B2i32(v405 < v210) == v405 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v577 = v576
	goto L85
L139:
	;
	v576 = v566
	goto L138
L140:
	;
	if v225 <= v436 {
		v471 = v225
		v473 = v405
		goto L149
	} else {
		goto L150
	}
L141:
	;
	v436 = v223
	v440 = v405
	goto L140
L142:
	;
	goto L143
L143:
	;
	v417 = v223
	v421 = v405
	goto L144
L144:
	;
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224+v421<<(uint(int32(1))%32)))))
	if v427 != 0 {
		v566 = int32(1)
		goto L139
	} else {
		goto L146
	}
L145:
	;
	v436 = v431
	v440 = v429
	goto L140
L146:
	;
	v428 = int32(1)
	v429 = v421 + v428
	v431 = v417 - v428
	if v431 <= v225 {
		v436 = v431
		v440 = v429
		goto L140
	} else {
		goto L147
	}
L147:
	;
	if v429 < v210 {
		v417 = v431
		v421 = v429
		goto L144
	} else {
		goto L148
	}
L148:
	;
	goto L145
L149:
	;
	if v436 != v471 {
		v512 = v440
		v513 = v473
		goto L157
	} else {
		goto L158
	}
L150:
	;
	if v211 <= int32(0) {
		v471 = v225
		v473 = v405
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v452 = v225
	v454 = v405
	goto L152
L152:
	;
	v459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226+v454<<(uint(int32(1))%32)))))
	if v459 != 0 {
		v566 = int32(-1)
		goto L139
	} else {
		goto L154
	}
L153:
	;
	v471 = v463
	v473 = v461
	goto L149
L154:
	;
	v460 = int32(1)
	v461 = v454 + v460
	v463 = v452 - v460
	if v463 <= v436 {
		v471 = v463
		v473 = v461
		goto L149
	} else {
		goto L155
	}
L155:
	;
	if v461 < v211 {
		v452 = v463
		v454 = v461
		goto L152
	} else {
		goto L156
	}
L156:
	;
	goto L153
L157:
	;
	if v210 < v512 {
		goto L167
	} else {
		goto L168
	}
L158:
	;
	v482 = v440
	v483 = v473
	goto L159
L159:
	;
	if v210 <= v482 {
		v512 = v482
		v513 = v483
		goto L157
	} else {
		goto L161
	}
L160:
	;
	if base.I32_extend16_s(v498) < base.I32_extend16_s(v496) {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	if v211 <= v483 {
		v512 = v482
		v513 = v483
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v487 = int32(1)
	v496 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224+v482<<(uint(v487)%32)))))
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v483<<(uint(v487)%32)+v226))))
	if v496 == v498 {
		v482 = v482 + v487
		v483 = v483 + v487
		goto L159
	} else {
		goto L163
	}
L163:
	;
	goto L160
L164:
	;
	v505 = int32(1)
	goto L166
L165:
	;
	v505 = int32(-1)
	goto L166
L166:
	;
	v576 = v505
	goto L138
L167:
	;
	v516 = v512
	goto L169
L168:
	;
	v516 = v210
	goto L169
L169:
	;
	v523 = v512
	goto L170
L170:
	;
	if v516 == v523 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v566 = v549
	goto L139
L172:
	;
	if v211 < v513 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	v549 = int32(1)
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224+v523<<(uint(v549)%32)))))
	if v555 == int32(0) {
		v523 = v523 + v549
		goto L170
	} else {
		goto L184
	}
L175:
	;
	v528 = v513
	goto L177
L176:
	;
	v528 = v211
	goto L177
L177:
	;
	v536 = v513
	goto L178
L178:
	;
	if v528 == v536 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v566 = int32(-1)
	goto L139
L180:
	;
	v576 = int32(0)
	goto L138
L181:
	;
	goto L182
L182:
	;
	v540 = int32(1)
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226+v536<<(uint(v540)%32)))))
	if v545 == int32(0) {
		v536 = v536 + v540
		goto L178
	} else {
		goto L183
	}
L183:
	;
	goto L179
L184:
	;
	goto L171
L185:
	;
	goto L70
L186:
	;
	if v581 == int32(0) {
		goto L70
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v581 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	if v580 != int32(16384) {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L70
L191:
	;
	if v589 == int32(0) {
		goto L70
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v205)+28))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v205)+44))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v589 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L1
L195:
	;
	if v948 < int32(0) {
		goto L1
	} else {
		goto L295
	}
L196:
	;
	if v580 == int32(16384) {
		goto L70
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	if v580 == int32(0) {
		goto L1
	} else {
		goto L247
	}
L199:
	;
	v602 = int32(0)
	if base.B2i32(v594 < v596)&base.B2i32(v602 < v582) == v602 {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	v948 = v773
	goto L195
L201:
	;
	v773 = v763
	goto L200
L202:
	;
	if v594 <= v633 {
		v668 = v594
		v670 = v602
		goto L211
	} else {
		goto L212
	}
L203:
	;
	v633 = v596
	v637 = v602
	goto L202
L204:
	;
	goto L205
L205:
	;
	v614 = v596
	v618 = v602
	goto L206
L206:
	;
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597+v618<<(uint(int32(1))%32)))))
	if v624 != 0 {
		v763 = int32(1)
		goto L201
	} else {
		goto L208
	}
L207:
	;
	v633 = v628
	v637 = v626
	goto L202
L208:
	;
	v625 = int32(1)
	v626 = v618 + v625
	v628 = v614 - v625
	if v628 <= v594 {
		v633 = v628
		v637 = v626
		goto L202
	} else {
		goto L209
	}
L209:
	;
	if v626 < v582 {
		v614 = v628
		v618 = v626
		goto L206
	} else {
		goto L210
	}
L210:
	;
	goto L207
L211:
	;
	if v633 != v668 {
		v709 = v637
		v710 = v670
		goto L219
	} else {
		goto L220
	}
L212:
	;
	if v581 <= int32(0) {
		v668 = v594
		v670 = v602
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v649 = v594
	v651 = v602
	goto L214
L214:
	;
	v656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595+v651<<(uint(int32(1))%32)))))
	if v656 != 0 {
		v763 = int32(-1)
		goto L201
	} else {
		goto L216
	}
L215:
	;
	v668 = v660
	v670 = v658
	goto L211
L216:
	;
	v657 = int32(1)
	v658 = v651 + v657
	v660 = v649 - v657
	if v660 <= v633 {
		v668 = v660
		v670 = v658
		goto L211
	} else {
		goto L217
	}
L217:
	;
	if v658 < v581 {
		v649 = v660
		v651 = v658
		goto L214
	} else {
		goto L218
	}
L218:
	;
	goto L215
L219:
	;
	if v582 < v709 {
		goto L229
	} else {
		goto L230
	}
L220:
	;
	v679 = v637
	v680 = v670
	goto L221
L221:
	;
	if v582 <= v679 {
		v709 = v679
		v710 = v680
		goto L219
	} else {
		goto L223
	}
L222:
	;
	if base.I32_extend16_s(v695) < base.I32_extend16_s(v693) {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	if v581 <= v680 {
		v709 = v679
		v710 = v680
		goto L219
	} else {
		goto L224
	}
L224:
	;
	v684 = int32(1)
	v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597+v679<<(uint(v684)%32)))))
	v695 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v680<<(uint(v684)%32)+v595))))
	if v693 == v695 {
		v679 = v679 + v684
		v680 = v680 + v684
		goto L221
	} else {
		goto L225
	}
L225:
	;
	goto L222
L226:
	;
	v702 = int32(1)
	goto L228
L227:
	;
	v702 = int32(-1)
	goto L228
L228:
	;
	v773 = v702
	goto L200
L229:
	;
	v713 = v709
	goto L231
L230:
	;
	v713 = v582
	goto L231
L231:
	;
	v720 = v709
	goto L232
L232:
	;
	if v713 == v720 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v763 = v746
	goto L201
L234:
	;
	if v581 < v710 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	v746 = int32(1)
	v752 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597+v720<<(uint(v746)%32)))))
	if v752 == int32(0) {
		v720 = v720 + v746
		goto L232
	} else {
		goto L246
	}
L237:
	;
	v725 = v710
	goto L239
L238:
	;
	v725 = v581
	goto L239
L239:
	;
	v733 = v710
	goto L240
L240:
	;
	if v725 == v733 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v763 = int32(-1)
	goto L201
L242:
	;
	v773 = int32(0)
	goto L200
L243:
	;
	goto L244
L244:
	;
	v737 = int32(1)
	v742 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595+v733<<(uint(v737)%32)))))
	if v742 == int32(0) {
		v733 = v733 + v737
		goto L240
	} else {
		goto L245
	}
L245:
	;
	goto L241
L246:
	;
	goto L233
L247:
	;
	v776 = int32(0)
	if base.B2i32(v596 < v594)&base.B2i32(v776 < v581) == v776 {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	v948 = v947
	goto L195
L249:
	;
	v947 = v937
	goto L248
L250:
	;
	if v596 <= v807 {
		v842 = v596
		v844 = v776
		goto L259
	} else {
		goto L260
	}
L251:
	;
	v807 = v594
	v811 = v776
	goto L250
L252:
	;
	goto L253
L253:
	;
	v788 = v594
	v792 = v776
	goto L254
L254:
	;
	v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595+v792<<(uint(int32(1))%32)))))
	if v798 != 0 {
		v937 = int32(1)
		goto L249
	} else {
		goto L256
	}
L255:
	;
	v807 = v802
	v811 = v800
	goto L250
L256:
	;
	v799 = int32(1)
	v800 = v792 + v799
	v802 = v788 - v799
	if v802 <= v596 {
		v807 = v802
		v811 = v800
		goto L250
	} else {
		goto L257
	}
L257:
	;
	if v800 < v581 {
		v788 = v802
		v792 = v800
		goto L254
	} else {
		goto L258
	}
L258:
	;
	goto L255
L259:
	;
	if v807 != v842 {
		v883 = v811
		v884 = v844
		goto L267
	} else {
		goto L268
	}
L260:
	;
	if v582 <= int32(0) {
		v842 = v596
		v844 = v776
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v823 = v596
	v825 = v776
	goto L262
L262:
	;
	v830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597+v825<<(uint(int32(1))%32)))))
	if v830 != 0 {
		v937 = int32(-1)
		goto L249
	} else {
		goto L264
	}
L263:
	;
	v842 = v834
	v844 = v832
	goto L259
L264:
	;
	v831 = int32(1)
	v832 = v825 + v831
	v834 = v823 - v831
	if v834 <= v807 {
		v842 = v834
		v844 = v832
		goto L259
	} else {
		goto L265
	}
L265:
	;
	if v832 < v582 {
		v823 = v834
		v825 = v832
		goto L262
	} else {
		goto L266
	}
L266:
	;
	goto L263
L267:
	;
	if v581 < v883 {
		goto L277
	} else {
		goto L278
	}
L268:
	;
	v853 = v811
	v854 = v844
	goto L269
L269:
	;
	if v581 <= v853 {
		v883 = v853
		v884 = v854
		goto L267
	} else {
		goto L271
	}
L270:
	;
	if base.I32_extend16_s(v869) < base.I32_extend16_s(v867) {
		goto L274
	} else {
		goto L275
	}
L271:
	;
	if v582 <= v854 {
		v883 = v853
		v884 = v854
		goto L267
	} else {
		goto L272
	}
L272:
	;
	v858 = int32(1)
	v867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595+v853<<(uint(v858)%32)))))
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v854<<(uint(v858)%32)+v597))))
	if v867 == v869 {
		v853 = v853 + v858
		v854 = v854 + v858
		goto L269
	} else {
		goto L273
	}
L273:
	;
	goto L270
L274:
	;
	v876 = int32(1)
	goto L276
L275:
	;
	v876 = int32(-1)
	goto L276
L276:
	;
	v947 = v876
	goto L248
L277:
	;
	v887 = v883
	goto L279
L278:
	;
	v887 = v581
	goto L279
L279:
	;
	v894 = v883
	goto L280
L280:
	;
	if v887 == v894 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v937 = v920
	goto L249
L282:
	;
	if v582 < v884 {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	goto L284
L284:
	;
	v920 = int32(1)
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595+v894<<(uint(v920)%32)))))
	if v926 == int32(0) {
		v894 = v894 + v920
		goto L280
	} else {
		goto L294
	}
L285:
	;
	v899 = v884
	goto L287
L286:
	;
	v899 = v582
	goto L287
L287:
	;
	v907 = v884
	goto L288
L288:
	;
	if v899 == v907 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v937 = int32(-1)
	goto L249
L290:
	;
	v947 = int32(0)
	goto L248
L291:
	;
	goto L292
L292:
	;
	v911 = int32(1)
	v916 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597+v907<<(uint(v911)%32)))))
	if v916 == int32(0) {
		v907 = v907 + v911
		goto L288
	} else {
		goto L293
	}
L293:
	;
	goto L289
L294:
	;
	goto L281
L295:
	;
	goto L70
L296:
	;
	v962 = int32(4515248)
	v963 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v204)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v965
	F_add_var(m, v205, v205+int32(48), v205)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L9
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v963
	v973 = *(*int64)(unsafe.Add(mBase, uint32(v204)))
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v973 + int64(1)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v977)+20)) = int32(1)
	return v960
L298:
	;
	F_errfinish(m, int32(499844), int32(1727), int32(490469))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L9
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	F_errfinish(m, int32(499844), int32(1738), int32(490469))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L9
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L302:
	;
	F_errfinish(m, int32(499844), int32(1755), int32(490469))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L9
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L9
	} else {
		goto L305
	}
L305:
	;
	F_errmsg(m, int32(240048), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L9
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(499844), int32(1767), int32(490469))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L9
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+20)) = int32(2)
	v1037 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1037)
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
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v92 int32
	_ = v92
	var v97 int64
	_ = v97
	var v99 float64
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int64
	_ = v111
	var v124 float64
	_ = v124
	var v128 int32
	_ = v128
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 != int32(460) {
		v128 = v2
		return v128
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		if v15 == int32(0) {
			v128 = v2
			return v128
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			if v18 != int32(15) {
				v128 = v2
				return v128
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
										v128 = v2
										return v128
									} else {
										v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
										if v57 != 0 {
											v124 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
											v128 = v11
											return v128
										} else {
											if v39 != int32(7) {
												v128 = v2
												return v128
											} else {
												if v46 != int32(7) {
													v128 = v2
													return v128
												} else {
													v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
													v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
													if base.Ui64(v63-int64(9223372036854775807)) < base.Ui64(int64(2)) {
														v128 = v2
														return v128
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
														v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
														if base.Ui64(v69-int64(9223372036854775807)) < base.Ui64(int64(2)) {
															v128 = v2
															return v128
														} else {
															if base.B2i32(v69-v63 < v69)^base.B2i32(int64(0) < v63) != 0 {
																v128 = v2
																return v128
															} else {
																v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
																v82 = F_Int64GetDatum(m, v69)
																mBase = m.M
																v83 = m.ExcPending
																if v83 != 0 {
																	return int32(0)
																} else {
																	v84 = F_Int64GetDatum(m, v63)
																	mBase = m.M
																	v85 = m.ExcPending
																	if v85 != 0 {
																		return int32(0)
																	} else {
																		v86 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v82, v84)
																		mBase = m.M
																		v87 = m.ExcPending
																		if v87 != 0 {
																			return int32(0)
																		} else {
																			v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																			v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																			v97 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
																			v99 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v88), float64(30)), base.F64_convert_i32_s(v92)), float64(8.64e+10)), base.F64_convert_i64_s(v97))
																			if base.F64_eq(v99, float64(0)) != 0 {
																				v128 = v2
																			} else {
																				v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
																				v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
																				v111 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																				v124 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v102), float64(30)), base.F64_convert_i32_s(v106)), float64(8.64e+10)), base.F64_convert_i64_s(v111)), v99), float64(1)))
																				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
																				v128 = v11
																			}
																			return v128
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
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
									if v49 == int32(0) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
										if v53 != int32(7) {
											v128 = v2
											return v128
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
											if v57 != 0 {
												v124 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
												v128 = v11
												return v128
											} else {
												if v39 != int32(7) {
													v128 = v2
													return v128
												} else {
													if v46 != int32(7) {
														v128 = v2
														return v128
													} else {
														v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
														if base.Ui64(v63-int64(9223372036854775807)) < base.Ui64(int64(2)) {
															v128 = v2
															return v128
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
															v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
															if base.Ui64(v69-int64(9223372036854775807)) < base.Ui64(int64(2)) {
																v128 = v2
																return v128
															} else {
																if base.B2i32(v69-v63 < v69)^base.B2i32(int64(0) < v63) != 0 {
																	v128 = v2
																	return v128
																} else {
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
																	v82 = F_Int64GetDatum(m, v69)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int32(0)
																	} else {
																		v84 = F_Int64GetDatum(m, v63)
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
																			return int32(0)
																		} else {
																			v86 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v82, v84)
																			mBase = m.M
																			v87 = m.ExcPending
																			if v87 != 0 {
																				return int32(0)
																			} else {
																				v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																				v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																				v97 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
																				v99 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v88), float64(30)), base.F64_convert_i32_s(v92)), float64(8.64e+10)), base.F64_convert_i64_s(v97))
																				if base.F64_eq(v99, float64(0)) != 0 {
																					v128 = v2
																				} else {
																					v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
																					v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
																					v111 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																					v124 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v102), float64(30)), base.F64_convert_i32_s(v106)), float64(8.64e+10)), base.F64_convert_i64_s(v111)), v99), float64(1)))
																					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
																					v128 = v11
																				}
																				return v128
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
									} else {
										v124 = float64(0)
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
										v128 = v11
										return v128
									}
								}
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)))
								if v42 == int32(0) {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									if v46 != int32(7) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
										if v53 != int32(7) {
											v128 = v2
											return v128
										} else {
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
											if v57 != 0 {
												v124 = float64(0)
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
												v128 = v11
												return v128
											} else {
												if v39 != int32(7) {
													v128 = v2
													return v128
												} else {
													if v46 != int32(7) {
														v128 = v2
														return v128
													} else {
														v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
														v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
														if base.Ui64(v63-int64(9223372036854775807)) < base.Ui64(int64(2)) {
															v128 = v2
															return v128
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
															v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
															if base.Ui64(v69-int64(9223372036854775807)) < base.Ui64(int64(2)) {
																v128 = v2
																return v128
															} else {
																if base.B2i32(v69-v63 < v69)^base.B2i32(int64(0) < v63) != 0 {
																	v128 = v2
																	return v128
																} else {
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
																	v82 = F_Int64GetDatum(m, v69)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int32(0)
																	} else {
																		v84 = F_Int64GetDatum(m, v63)
																		mBase = m.M
																		v85 = m.ExcPending
																		if v85 != 0 {
																			return int32(0)
																		} else {
																			v86 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v82, v84)
																			mBase = m.M
																			v87 = m.ExcPending
																			if v87 != 0 {
																				return int32(0)
																			} else {
																				v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																				v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																				v97 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
																				v99 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v88), float64(30)), base.F64_convert_i32_s(v92)), float64(8.64e+10)), base.F64_convert_i64_s(v97))
																				if base.F64_eq(v99, float64(0)) != 0 {
																					v128 = v2
																				} else {
																					v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
																					v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
																					v111 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																					v124 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v102), float64(30)), base.F64_convert_i32_s(v106)), float64(8.64e+10)), base.F64_convert_i64_s(v111)), v99), float64(1)))
																					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
																					v128 = v11
																				}
																				return v128
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
									} else {
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+24)))
										if v49 == int32(0) {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
											if v53 != int32(7) {
												v128 = v2
												return v128
											} else {
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+24)))
												if v57 != 0 {
													v124 = float64(0)
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
													v128 = v11
													return v128
												} else {
													if v39 != int32(7) {
														v128 = v2
														return v128
													} else {
														if v46 != int32(7) {
															v128 = v2
															return v128
														} else {
															v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
															v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
															if base.Ui64(v63-int64(9223372036854775807)) < base.Ui64(int64(2)) {
																v128 = v2
																return v128
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
																v69 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
																if base.Ui64(v69-int64(9223372036854775807)) < base.Ui64(int64(2)) {
																	v128 = v2
																	return v128
																} else {
																	if base.B2i32(v69-v63 < v69)^base.B2i32(int64(0) < v63) != 0 {
																		v128 = v2
																		return v128
																	} else {
																		v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
																		v82 = F_Int64GetDatum(m, v69)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return int32(0)
																		} else {
																			v84 = F_Int64GetDatum(m, v63)
																			mBase = m.M
																			v85 = m.ExcPending
																			if v85 != 0 {
																				return int32(0)
																			} else {
																				v86 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v82, v84)
																				mBase = m.M
																				v87 = m.ExcPending
																				if v87 != 0 {
																					return int32(0)
																				} else {
																					v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
																					v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
																					v97 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
																					v99 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v88), float64(30)), base.F64_convert_i32_s(v92)), float64(8.64e+10)), base.F64_convert_i64_s(v97))
																					if base.F64_eq(v99, float64(0)) != 0 {
																						v128 = v2
																					} else {
																						v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
																						v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
																						v111 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
																						v124 = base.F64_floor(base.F64_add(base.F64_div(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v102), float64(30)), base.F64_convert_i32_s(v106)), float64(8.64e+10)), base.F64_convert_i64_s(v111)), v99), float64(1)))
																						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
																						v128 = v11
																					}
																					return v128
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
										} else {
											v124 = float64(0)
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
											v128 = v11
											return v128
										}
									}
								} else {
									v124 = float64(0)
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v124
									v128 = v11
									return v128
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int64
	_ = v68
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v98 int64
	_ = v98
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
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
	return v205
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L50
	}
L3:
	;
	m.G0 = v13 + int32(272)
	goto L1
L4:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L49
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L12
	} else {
		goto L45
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
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	goto L36
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
	v35 = int32(4515248)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
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
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v68 = base.I64_extend_i32_s(v62) + base.I64_extend_i32_s(v64)*int64(30)
	v77 = int64(32)
	v78 = int64(20)
	v80 = int64(base.Ui64(v68) >> (uint(v77) % 64))
	v83 = int64(4294967295)
	v84 = int64(500654080)
	v86 = v68 & v83
	v87 = v84 * v86
	v91 = int64(base.Ui64(v87)>>(uint(v77)%64)) + v84*v80
	v98 = v86*v78 + v91&v83
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v68*int64(0) + v68>>(uint(int64(63))%64)*int64(86400000000) + v78*v80 + int64(base.Ui64(v91)>>(uint(v77)%64)) + int64(base.Ui64(v98)>>(uint(v77)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v87&v83 | v98<<(uint(v77)%64)
	goto L22
L17:
	;
	F_text_to_cstring_buffer(m, v32, v13+int32(16), int32(256))
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
	v59 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	v60 = v59
	goto L16
L20:
	;
	v56 = F_DecodeTimezoneNameToTz(m, v13+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v60 = v56
	goto L16
L22:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v41)+16))
	v111 = v109 + v110
	v112 = int64(0)
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v120 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v111) < base.Ui64(v109))) + (v116 + v110>>(uint(int64(63))%64))
	if v120 == v112 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v125 = base.B2i32(v111 != v112)
	goto L25
L24:
	;
	v125 = base.B2i32(v112 < v120)
	goto L25
L25:
	;
	v126 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = v125 - base.B2i32(v120 < v126)
	if v120|v111 == v126 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	if v64 != int32(2147483647) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v41
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
	goto L8
L28:
	;
	if v64 != int32(-2147483648) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v110 != int64(9223372036854775807) {
		goto L27
	} else {
		goto L34
	}
L31:
	;
	if v110 != int64(-9223372036854775807-1) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if v62 != int32(-2147483648) {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	if v62 == int32(2147483647) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v158)+8))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v158)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+32))
	if int32(0) < v161 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v158)+36))
	v169 = F_timestamptz_pl_interval_internal(m, v160, v158+int32(16), v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L43
	}
L38:
	;
	if v160 <= v159 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v160 < v159 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L4
L42:
	;
	goto L37
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v158))) = v169
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	*(*int64)(unsafe.Add(mBase, uint32(v157))) = v172 + int64(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+20)) = int32(1)
	v179 = F_Int64GetDatum(m, v160)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	v205 = v179
	goto L3
L45:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(240048), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(495703), int32(6794), int32(310543))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+20)) = int32(2)
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v202)
	v205 = int32(0)
	goto L3
L50:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(350177), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(495703), int32(6799), int32(310543))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
