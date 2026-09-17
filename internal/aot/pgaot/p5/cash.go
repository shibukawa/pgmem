package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_div_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_cash_div_cash_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_cash_div_cash_1), int32(725), int32(_a_F_cash_div_cash_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
		v30 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v26), base.F64_convert_i64_s(v4)))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_cash_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v5 < v3)
}
func F_cash_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v81 int32
	_ = v81
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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int64
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
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
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v388 int32
	_ = v388
	var v399 int64
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v438 int64
	_ = v438
	var v443 int64
	_ = v443
	var v446 int64
	_ = v446
	var v449 int64
	_ = v449
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v457 int64
	_ = v457
	var v464 int64
	_ = v464
	var v475 int64
	_ = v475
	var v476 int64
	_ = v476
	var v482 int64
	_ = v482
	var v485 int64
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int64
	_ = v574
	var v575 int64
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v588 int64
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v649 int64
	_ = v649
	var v650 int64
	_ = v650
	var v655 int32
	_ = v655
	var v658 int64
	_ = v658
	var v659 int64
	_ = v659
	var v664 int64
	_ = v664
	var v667 int64
	_ = v667
	var v670 int64
	_ = v670
	var v673 int64
	_ = v673
	var v674 int64
	_ = v674
	var v678 int64
	_ = v678
	var v685 int64
	_ = v685
	var v696 int64
	_ = v696
	var v697 int64
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v721 int64
	_ = v721
	var v739 int64
	_ = v739
	var v744 int32
	_ = v744
	var v767 int32
	_ = v767
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v792 int64
	_ = v792
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v989 int64
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int64
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	v17 = int64(0)
	v22 = m.G0
	v24 = v22 - int32(112)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_PGLC_localeconv(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+41)))
	v35 = int32(_a_F_cash_in_0)
	v36 = int32(46)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v38 == int32(0) {
		v48 = v35
		v49 = v36
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(int32(10)) < base.Ui32(v32) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v41 != 0 {
		v48 = v35
		v49 = v36
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v38 == int32(44) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = int32(_a_F_cash_in_1)
	goto L8
L7:
	;
	v46 = int32(_a_F_cash_in_0)
	goto L8
L8:
	;
	v48 = v46
	v49 = base.I32_extend8_s(v38)
	goto L3
L9:
	;
	v51 = int32(2)
	goto L11
L10:
	;
	v51 = v32
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v60 = v27
	goto L12
L12:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v81-int32(9)))&base.B2i32(v81 != int32(32)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v57 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v60 = v60 + int32(1)
	goto L12
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	v95 = v56
	goto L19
L18:
	;
	v95 = int32(_a_F_cash_in_2)
	goto L19
L19:
	;
	v96 = F_strlen(m, v95)
	mBase = m.M
	if v96 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v141 != 0 {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v141 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v102 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v103 = v60
	v104 = v95
	v105 = v96
	v106 = v102
	goto L28
L25:
	;
	v129 = v95
	v133 = int32(0)
	goto L26
L26:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v141 = v133 - v134
	goto L20
L27:
	;
	v129 = v124
	v133 = v126
	goto L26
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if base.B2i32(v106 != v108)|base.B2i32(v108 == int32(0)) != 0 {
		v124 = v104
		v126 = v106
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v124 = v118
	v126 = int32(0)
	goto L27
L30:
	;
	v114 = v105 - int32(1)
	if v114 == int32(0) {
		v124 = v104
		v126 = v106
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v117 = int32(1)
	v118 = v104 + v117
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v119 != 0 {
		v103 = v103 + v117
		v104 = v118
		v105 = v114
		v106 = v119
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v142 = int32(0)
	goto L35
L34:
	;
	v142 = v96
	goto L35
L35:
	;
	v144 = v60 + v142
	goto L36
L36:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v165-int32(9)))&base.B2i32(v165 != int32(32)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v55 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v144 = v144 + int32(1)
	goto L36
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	v178 = v54
	goto L43
L42:
	;
	v178 = int32(_a_F_cash_in_3)
	goto L43
L43:
	;
	v179 = int64(-1)
	if v53 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v59 != 0 {
		goto L79
	} else {
		goto L80
	}
L45:
	;
	v181 = v52
	goto L47
L46:
	;
	v181 = int32(_a_F_cash_in_4)
	goto L47
L47:
	;
	v182 = F_strlen(m, v181)
	mBase = m.M
	if v182 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v227 == int32(0) {
		v282 = v182
		v283 = v179
		goto L44
	} else {
		goto L61
	}
L49:
	;
	v227 = int32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v188 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v189 = v144
	v190 = v181
	v191 = v182
	v192 = v188
	goto L56
L53:
	;
	v215 = v181
	v219 = int32(0)
	goto L54
L54:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v227 = v219 - v220
	goto L48
L55:
	;
	v215 = v210
	v219 = v212
	goto L54
L56:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if base.B2i32(v192 != v194)|base.B2i32(v194 == int32(0)) != 0 {
		v210 = v190
		v212 = v192
		goto L55
	} else {
		goto L58
	}
L57:
	;
	v210 = v204
	v212 = int32(0)
	goto L55
L58:
	;
	v200 = v191 - int32(1)
	if v200 == int32(0) {
		v210 = v190
		v212 = v192
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v203 = int32(1)
	v204 = v190 + v203
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	if v205 != 0 {
		v189 = v189 + v203
		v190 = v204
		v191 = v200
		v192 = v205
		goto L56
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	if v165 == int32(40) {
		v282 = int32(1)
		v283 = v179
		goto L44
	} else {
		goto L62
	}
L62:
	;
	v233 = int32(0)
	v234 = F_strlen(m, v178)
	mBase = m.M
	if v234 == v233 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v279 != 0 {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v279 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v240 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v241 = v144
	v242 = v178
	v243 = v234
	v244 = v240
	goto L71
L68:
	;
	v267 = v178
	v271 = int32(0)
	goto L69
L69:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	v279 = v271 - v272
	goto L63
L70:
	;
	v267 = v262
	v271 = v264
	goto L69
L71:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if base.B2i32(v244 != v246)|base.B2i32(v246 == int32(0)) != 0 {
		v262 = v242
		v264 = v244
		goto L70
	} else {
		goto L73
	}
L72:
	;
	v262 = v256
	v264 = int32(0)
	goto L70
L73:
	;
	v252 = v243 - int32(1)
	if v252 == int32(0) {
		v262 = v242
		v264 = v244
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v255 = int32(1)
	v256 = v242 + v255
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	if v257 != 0 {
		v241 = v241 + v255
		v242 = v256
		v243 = v252
		v244 = v257
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v280 = v233
	goto L78
L77:
	;
	v280 = v234
	goto L78
L78:
	;
	v282 = v280
	v283 = int64(1)
	goto L44
L79:
	;
	v284 = v58
	goto L81
L80:
	;
	v284 = v48
	goto L81
L81:
	;
	v286 = v144 + v282
	goto L82
L82:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v307-int32(9)))&base.B2i32(v307 != int32(32)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v319 = int32(0)
	if v96 == v319 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v286 = v286 + int32(1)
	goto L82
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	if v364 != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v364 = int32(0)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v325 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v326 = v286
	v327 = v95
	v328 = v96
	v329 = v325
	goto L95
L92:
	;
	v352 = v95
	v356 = int32(0)
	goto L93
L93:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v364 = v356 - v357
	goto L87
L94:
	;
	v352 = v347
	v356 = v349
	goto L93
L95:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if base.B2i32(v329 != v331)|base.B2i32(v331 == int32(0)) != 0 {
		v347 = v327
		v349 = v329
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v347 = v341
	v349 = int32(0)
	goto L94
L97:
	;
	v337 = v328 - int32(1)
	if v337 == int32(0) {
		v347 = v327
		v349 = v329
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v340 = int32(1)
	v341 = v327 + v340
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+1)))
	if v342 != 0 {
		v326 = v326 + v340
		v327 = v341
		v328 = v337
		v329 = v342
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	v365 = v319
	goto L102
L101:
	;
	v365 = v96
	goto L102
L102:
	;
	v367 = v286 + v365
	goto L105
L103:
	;
	m.G0 = v24 + int32(112)
	return v996
L104:
	;
	if v628 < v399 {
		goto L155
	} else {
		goto L156
	}
L105:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if base.B2i32(base.Ui32(v388-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v388 == int32(32)) != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v401 = v367
	v403 = v388
	v404 = int32(0)
	v417 = v17
	v418 = v17
	goto L112
L107:
	;
	v367 = v367 + int32(1)
	goto L105
L108:
	;
	v399 = base.I64_extend8_s(base.I64_extend_i32_u(v51))
	if v388 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L106
L110:
	;
	goto L109
L111:
	;
	v611 = v367
	v627 = v17
	v628 = v17
	goto L104
L112:
	;
	v425 = v403 - int32(48)
	if v404&base.B2i32(v399 <= v418)|base.B2i32(base.Ui32(int32(9)) < base.Ui32(v425&int32(255))) == int32(0) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	if base.Ui32(int32(4)) < base.Ui32((v403-int32(53))&int32(255)) {
		v611 = v401
		v627 = v417
		v628 = v418
		goto L104
	} else {
		goto L148
	}
L114:
	;
	goto L113
L115:
	;
	v579 = v577 + int32(1)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	if v580 != 0 {
		v401 = v579
		v403 = v580
		v404 = v572
		v417 = v574
		v418 = v575
		goto L112
	} else {
		goto L147
	}
L116:
	;
	v434 = v24 + int32(96)
	v437 = int64(10)
	v438 = int64(0)
	v443 = int64(32)
	v446 = int64(base.Ui64(v417) >> (uint(v443) % 64))
	v449 = int64(4294967295)
	v452 = v417 & v449
	v453 = v437 * v452
	v457 = int64(base.Ui64(v453)>>(uint(v443)%64)) + v437*v446
	v464 = v452*v438 + v457&v449
	*(*int64)(unsafe.Add(mBase, uint32(v434)+8)) = v417*v438 + v417>>(uint(int64(63))%64)*v437 + v438*v446 + int64(base.Ui64(v457)>>(uint(v443)%64)) + int64(base.Ui64(v464)>>(uint(v443)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = v453&v449 | v464<<(uint(v443)%64)
	goto L119
L117:
	;
	goto L118
L118:
	;
	if (base.B2i32(v49 != base.I32_extend8_s(v403))|v404)&int32(1) == int32(0) {
		goto L130
	} else {
		goto L131
	}
L119:
	;
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v24)+104))
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v24)+96))
	if v475 == v476>>(uint(int64(63))%64) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v572 = v404
	v574 = v485
	v575 = v418 + base.I64_extend_i32_u(v404)&int64(1)
	v577 = v401
	goto L115
L121:
	;
	v482 = base.I64_extend_i32_u(v425) & int64(255)
	v485 = v476 - v482
	if base.B2i32(v482 != int64(0)) == base.B2i32(v485 < v476) {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v490 = int32(0)
	v491 = F_errsave_start(m, v26)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	if v491 == int32(0) {
		v996 = v490
		goto L103
	} else {
		goto L126
	}
L126:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = int32(_a_F_cash_in_5)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v27
	F_errmsg(m, int32(_a_F_cash_in_6), v24+int32(80))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errsave_finish(m, v26, int32(_a_F_cash_in_7), int32(293), int32(_a_F_cash_in_8))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v996 = v490
	goto L103
L130:
	;
	v572 = int32(1)
	v574 = v417
	v575 = v418
	v577 = v401
	goto L115
L131:
	;
	goto L132
L132:
	;
	v523 = F_strlen(m, v284)
	mBase = m.M
	if v523 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v568 != 0 {
		goto L114
	} else {
		goto L146
	}
L134:
	;
	v568 = int32(0)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v529 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v530 = v401
	v531 = v284
	v532 = v523
	v533 = v529
	goto L141
L138:
	;
	v556 = v284
	v560 = int32(0)
	goto L139
L139:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v568 = v560 - v561
	goto L133
L140:
	;
	v556 = v551
	v560 = v553
	goto L139
L141:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
	if base.B2i32(v533 != v535)|base.B2i32(v535 == int32(0)) != 0 {
		v551 = v531
		v553 = v533
		goto L140
	} else {
		goto L143
	}
L142:
	;
	v551 = v545
	v553 = int32(0)
	goto L140
L143:
	;
	v541 = v532 - int32(1)
	if v541 == int32(0) {
		v551 = v531
		v553 = v533
		goto L140
	} else {
		goto L144
	}
L144:
	;
	v544 = int32(1)
	v545 = v531 + v544
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530)+1)))
	if v546 != 0 {
		v530 = v530 + v544
		v531 = v545
		v532 = v541
		v533 = v546
		goto L141
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	v572 = v404
	v574 = v417
	v575 = v418
	v577 = v401 + v523 - int32(1)
	goto L115
L147:
	;
	v611 = v579
	v627 = v574
	v628 = v575
	goto L104
L148:
	;
	v588 = v417 - int64(1)
	if v588 < v417 {
		v611 = v401
		v627 = v588
		v628 = v418
		goto L104
	} else {
		goto L149
	}
L149:
	;
	v590 = int32(0)
	v591 = F_errsave_start(m, v26)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	if v591 == int32(0) {
		v996 = v590
		goto L103
	} else {
		goto L151
	}
L151:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = int32(_a_F_cash_in_5)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v27
	F_errmsg(m, int32(_a_F_cash_in_6), v24-int32(-64))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errsave_finish(m, v26, int32(_a_F_cash_in_7), int32(318), int32(_a_F_cash_in_8))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v996 = v590
	goto L103
L155:
	;
	v649 = v627
	v650 = v628
	goto L158
L156:
	;
	v739 = v627
	goto L157
L157:
	;
	v744 = v611
	goto L170
L158:
	;
	v655 = v24 + int32(48)
	v658 = int64(10)
	v659 = int64(0)
	v664 = int64(32)
	v667 = int64(base.Ui64(v649) >> (uint(v664) % 64))
	v670 = int64(4294967295)
	v673 = v649 & v670
	v674 = v658 * v673
	v678 = int64(base.Ui64(v674)>>(uint(v664)%64)) + v658*v667
	v685 = v673*v659 + v678&v670
	*(*int64)(unsafe.Add(mBase, uint32(v655)+8)) = v649*v659 + v649>>(uint(int64(63))%64)*v658 + v659*v667 + int64(base.Ui64(v678)>>(uint(v664)%64)) + int64(base.Ui64(v685)>>(uint(v664)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v655))) = v674&v670 | v685<<(uint(v664)%64)
	goto L160
L159:
	;
	v739 = v697
	goto L157
L160:
	;
	v696 = *(*int64)(unsafe.Add(mBase, uint32(v24)+56))
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v24)+48))
	if v696 != v697>>(uint(int64(63))%64) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v701 = int32(0)
	v702 = F_errsave_start(m, v26)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v721 = v650 + int64(1)
	if v721 != v399 {
		v649 = v697
		v650 = v721
		goto L158
	} else {
		goto L169
	}
L164:
	;
	if v702 == int32(0) {
		v996 = v701
		goto L103
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(_a_F_cash_in_5)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27
	F_errmsg(m, int32(_a_F_cash_in_6), v24)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errsave_finish(m, v26, int32(_a_F_cash_in_7), int32(328), int32(_a_F_cash_in_8))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v996 = v701
	goto L103
L169:
	;
	goto L159
L170:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if base.Ui32((v767-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v744 = v744 + int32(1)
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v776 = v744
	v777 = v767
	v792 = v283
	goto L173
L172:
	;
	goto L171
L173:
	;
	switch v777 {
	case 0:
		goto L176
	default:
		goto L177
	case 9, 10, 11, 12, 13, 32, 41:
		v992 = int32(1)
		v993 = v792
		goto L175
	}
L175:
	;
	v994 = v992 + v776
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994))))
	v776 = v994
	v777 = v995
	v792 = v993
	goto L173
L176:
	;
	if int64(0) < v792 {
		goto L227
	} else {
		goto L228
	}
L177:
	;
	v796 = F_strlen(m, v181)
	mBase = m.M
	if v796 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v841 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L179:
	;
	v841 = int32(0)
	goto L178
L180:
	;
	goto L181
L181:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
	if v802 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v803 = v776
	v804 = v181
	v805 = v796
	v806 = v802
	goto L186
L183:
	;
	v829 = v181
	v833 = int32(0)
	goto L184
L184:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	v841 = v833 - v834
	goto L178
L185:
	;
	v829 = v824
	v833 = v826
	goto L184
L186:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	if base.B2i32(v806 != v808)|base.B2i32(v808 == int32(0)) != 0 {
		v824 = v804
		v826 = v806
		goto L185
	} else {
		goto L188
	}
L187:
	;
	v824 = v818
	v826 = int32(0)
	goto L185
L188:
	;
	v814 = v805 - int32(1)
	if v814 == int32(0) {
		v824 = v804
		v826 = v806
		goto L185
	} else {
		goto L189
	}
L189:
	;
	v817 = int32(1)
	v818 = v804 + v817
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+1)))
	if v819 != 0 {
		v803 = v803 + v817
		v804 = v818
		v805 = v814
		v806 = v819
		goto L186
	} else {
		goto L190
	}
L190:
	;
	goto L187
L191:
	;
	v992 = v796
	v993 = int64(-1)
	goto L175
L192:
	;
	goto L193
L193:
	;
	v845 = F_strlen(m, v178)
	mBase = m.M
	if v845 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	if v890 == int32(0) {
		v992 = v845
		v993 = v792
		goto L175
	} else {
		goto L207
	}
L195:
	;
	v890 = int32(0)
	goto L194
L196:
	;
	goto L197
L197:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
	if v851 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v852 = v776
	v853 = v178
	v854 = v845
	v855 = v851
	goto L202
L199:
	;
	v878 = v178
	v882 = int32(0)
	goto L200
L200:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878))))
	v890 = v882 - v883
	goto L194
L201:
	;
	v878 = v873
	v882 = v875
	goto L200
L202:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853))))
	if base.B2i32(v855 != v857)|base.B2i32(v857 == int32(0)) != 0 {
		v873 = v853
		v875 = v855
		goto L201
	} else {
		goto L204
	}
L203:
	;
	v873 = v867
	v875 = int32(0)
	goto L201
L204:
	;
	v863 = v854 - int32(1)
	if v863 == int32(0) {
		v873 = v853
		v875 = v855
		goto L201
	} else {
		goto L205
	}
L205:
	;
	v866 = int32(1)
	v867 = v853 + v866
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+1)))
	if v868 != 0 {
		v852 = v852 + v866
		v853 = v867
		v854 = v863
		v855 = v868
		goto L202
	} else {
		goto L206
	}
L206:
	;
	goto L203
L207:
	;
	v893 = F_strlen(m, v95)
	mBase = m.M
	if v893 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v938 == int32(0) {
		v992 = v893
		v993 = v792
		goto L175
	} else {
		goto L221
	}
L209:
	;
	v938 = int32(0)
	goto L208
L210:
	;
	goto L211
L211:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v776))))
	if v899 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v900 = v776
	v901 = v95
	v902 = v893
	v903 = v899
	goto L216
L213:
	;
	v926 = v95
	v930 = int32(0)
	goto L214
L214:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
	v938 = v930 - v931
	goto L208
L215:
	;
	v926 = v921
	v930 = v923
	goto L214
L216:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901))))
	if base.B2i32(v903 != v905)|base.B2i32(v905 == int32(0)) != 0 {
		v921 = v901
		v923 = v903
		goto L215
	} else {
		goto L218
	}
L217:
	;
	v921 = v915
	v923 = int32(0)
	goto L215
L218:
	;
	v911 = v902 - int32(1)
	if v911 == int32(0) {
		v921 = v901
		v923 = v903
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v914 = int32(1)
	v915 = v901 + v914
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900)+1)))
	if v916 != 0 {
		v900 = v900 + v914
		v901 = v915
		v902 = v911
		v903 = v916
		goto L216
	} else {
		goto L220
	}
L220:
	;
	goto L217
L221:
	;
	v941 = int32(0)
	v942 = F_errsave_start(m, v26)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	if v942 == int32(0) {
		v996 = v941
		goto L103
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(_a_F_cash_in_5)
	F_errmsg(m, int32(_a_F_cash_in_9), v24+int32(16))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_errsave_finish(m, v26, int32(_a_F_cash_in_7), int32(355), int32(_a_F_cash_in_8))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v996 = v941
	goto L103
L227:
	;
	if v739 == int64(-9223372036854775807-1) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	v989 = v739
	goto L229
L229:
	;
	v990 = F_Int64GetDatum(m, v989)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L1
	} else {
		goto L238
	}
L230:
	;
	v966 = int32(0)
	v967 = F_errsave_start(m, v26)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v989 = int64(0) - v739
	goto L229
L233:
	;
	if v967 == int32(0) {
		v996 = v966
		goto L103
	} else {
		goto L234
	}
L234:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = int32(_a_F_cash_in_5)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v27
	F_errmsg(m, int32(_a_F_cash_in_6), v24+int32(32))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_errsave_finish(m, v26, int32(_a_F_cash_in_7), int32(368), int32(_a_F_cash_in_8))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v996 = v966
	goto L103
L238:
	;
	v996 = v990
	goto L103
}
func F_cash_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v3 <= v5)
}
func F_cash_mul_flt8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_cash_mul_float8(m, v3, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_cash_words(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	F_initStringInfo(m, v9)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v12 < int64(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v24 = v12
	goto L5
L5:
	;
	v26 = base.I64_div_u_s(v24, int64(100))
	v28 = base.I64_rem_u_s(v26, int64(1000))
	if base.Ui64(int64(100000000000000000)) <= base.Ui64(v24) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v24 = int64(0) - v12
	goto L5
L7:
	;
	v32 = base.I64_div_u_s(v24, int64(100000000000000000))
	F_append_num_word(m, v9, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v39 = base.I64_div_u_s(v24, int64(100000000000000))
	v42 = base.I32_rem_u_s(base.I32_wrap_i64(v39), int32(1000))
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_append_num_word(m, v9, base.I64_extend_i32_u(v42))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = base.I64_div_u_s(v24, int64(100000000000))
	v53 = base.I32_rem_u_s(base.I32_wrap_i64(v50), int32(1000))
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_append_num_word(m, v9, base.I64_extend_i32_u(v53))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = base.I64_div_u_s(v24, int64(100000000))
	v63 = base.I64_rem_u_s(v61, int64(1000))
	if v63 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_append_num_word(m, v9, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = base.I64_div_u_s(v24, int64(100000))
	v74 = base.I64_rem_u_s(v72, int64(1000))
	if v74 != int64(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	F_append_num_word(m, v9, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v28 != int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_5))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	F_append_num_word(m, v9, v28)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui64(v24) <= base.Ui64(int64(99)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_cash_words_6))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v93 = int64(100)
	if base.Ui64(v24-v93) < base.Ui64(v93) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v97 = int32(_a_F_cash_words_7)
	goto L42
L41:
	;
	v97 = int32(_a_F_cash_words_8)
	goto L42
L42:
	;
	F_appendStringInfoString(m, v9, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v102 = v24 - v26*int64(100)
	F_append_num_word(m, v9, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v102 == int64(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v109 = int32(_a_F_cash_words_9)
	goto L47
L46:
	;
	v109 = int32(_a_F_cash_words_10)
	goto L47
L47:
	;
	F_appendStringInfoString(m, v9, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if base.Ui32((v113-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v124)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v129 = F_cstring_to_text_with_len(m, v127, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v122 = v113 - int32(32)
	goto L52
L51:
	;
	v122 = v113
	goto L52
L52:
	;
	v124 = v122 & int32(255)
	goto L49
L53:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	F_pfree(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v9 + int32(16)
	return v129
}
