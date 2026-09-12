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
				F_errmsg(m, int32(240257), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499393), int32(725), int32(324467))
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v83 int32
	_ = v83
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v170 int32
	_ = v170
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v391 int32
	_ = v391
	var v399 int64
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v436 int64
	_ = v436
	var v437 int64
	_ = v437
	var v442 int64
	_ = v442
	var v445 int64
	_ = v445
	var v448 int64
	_ = v448
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v456 int64
	_ = v456
	var v463 int64
	_ = v463
	var v474 int64
	_ = v474
	var v475 int64
	_ = v475
	var v481 int64
	_ = v481
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int64
	_ = v572
	var v573 int64
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v585 int64
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v624 int64
	_ = v624
	var v625 int64
	_ = v625
	var v646 int64
	_ = v646
	var v647 int64
	_ = v647
	var v652 int32
	_ = v652
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v661 int64
	_ = v661
	var v664 int64
	_ = v664
	var v667 int64
	_ = v667
	var v670 int64
	_ = v670
	var v671 int64
	_ = v671
	var v675 int64
	_ = v675
	var v682 int64
	_ = v682
	var v693 int64
	_ = v693
	var v694 int64
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v718 int64
	_ = v718
	var v736 int64
	_ = v736
	var v741 int32
	_ = v741
	var v764 int32
	_ = v764
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v790 int64
	_ = v790
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v985 int64
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int64
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
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
	v35 = int32(671939)
	v36 = int32(46)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v38 == int32(0) {
		v50 = v35
		v51 = v36
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
		v50 = v35
		v51 = v36
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v38&int32(255) == int32(44) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v48 = int32(671888)
	goto L8
L7:
	;
	v48 = int32(671939)
	goto L8
L8:
	;
	v50 = v48
	v51 = base.I32_extend8_s(v38)
	goto L3
L9:
	;
	v53 = int32(2)
	goto L11
L10:
	;
	v53 = v32
	goto L11
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v62 = v27
	goto L12
L12:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v83-int32(9)))&base.B2i32(v83 != int32(32)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v59&int32(255) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v62 = v62 + int32(1)
	goto L12
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	v99 = v58
	goto L19
L18:
	;
	v99 = int32(691930)
	goto L19
L19:
	;
	v100 = F_strlen(m, v99)
	mBase = m.M
	if v100 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v144 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v144 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v106 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v107 = v62
	v108 = v99
	v109 = v100
	v110 = v106
	goto L28
L25:
	;
	v132 = v99
	v136 = int32(0)
	goto L26
L26:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v144 = v136 - v137
	goto L20
L27:
	;
	v132 = v127
	v136 = v129
	goto L26
L28:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v110 != v112 {
		v127 = v108
		v129 = v110
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v127 = v121
	v129 = int32(0)
	goto L27
L30:
	;
	if v112 == int32(0) {
		v127 = v108
		v129 = v110
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v117 = v109 - int32(1)
	if v117 == int32(0) {
		v127 = v108
		v129 = v110
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v120 = int32(1)
	v121 = v108 + v120
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v122 != 0 {
		v107 = v107 + v120
		v108 = v121
		v109 = v117
		v110 = v122
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v145 = int32(0)
	goto L36
L35:
	;
	v145 = v100
	goto L36
L36:
	;
	if v57 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v148 = v56
	goto L39
L38:
	;
	v148 = int32(671941)
	goto L39
L39:
	;
	v149 = v62 + v145
	goto L40
L40:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v170-int32(9)))&base.B2i32(v170 != int32(32)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v182 = int64(-1)
	if v55&int32(255) != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v149 = v149 + int32(1)
	goto L40
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	if v61 != 0 {
		goto L82
	} else {
		goto L83
	}
L46:
	;
	v186 = v54
	goto L48
L47:
	;
	v186 = int32(671929)
	goto L48
L48:
	;
	v187 = F_strlen(m, v186)
	mBase = m.M
	if v187 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v231 == int32(0) {
		v286 = v187
		v287 = v182
		goto L45
	} else {
		goto L63
	}
L50:
	;
	v231 = int32(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v193 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v194 = v149
	v195 = v186
	v196 = v187
	v197 = v193
	goto L57
L54:
	;
	v219 = v186
	v223 = int32(0)
	goto L55
L55:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v231 = v223 - v224
	goto L49
L56:
	;
	v219 = v214
	v223 = v216
	goto L55
L57:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v197 != v199 {
		v214 = v195
		v216 = v197
		goto L56
	} else {
		goto L59
	}
L58:
	;
	v214 = v208
	v216 = int32(0)
	goto L56
L59:
	;
	if v199 == int32(0) {
		v214 = v195
		v216 = v197
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v204 = v196 - int32(1)
	if v204 == int32(0) {
		v214 = v195
		v216 = v197
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v207 = int32(1)
	v208 = v195 + v207
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v209 != 0 {
		v194 = v194 + v207
		v195 = v208
		v196 = v204
		v197 = v209
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	if v170 == int32(40) {
		v286 = int32(1)
		v287 = v182
		goto L45
	} else {
		goto L64
	}
L64:
	;
	v237 = int32(0)
	v238 = F_strlen(m, v148)
	mBase = m.M
	if v238 == v237 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v282 != 0 {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	v282 = int32(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v244 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v245 = v149
	v246 = v148
	v247 = v238
	v248 = v244
	goto L73
L70:
	;
	v270 = v148
	v274 = int32(0)
	goto L71
L71:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v282 = v274 - v275
	goto L65
L72:
	;
	v270 = v265
	v274 = v267
	goto L71
L73:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v248 != v250 {
		v265 = v246
		v267 = v248
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v265 = v259
	v267 = int32(0)
	goto L72
L75:
	;
	if v250 == int32(0) {
		v265 = v246
		v267 = v248
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v255 = v247 - int32(1)
	if v255 == int32(0) {
		v265 = v246
		v267 = v248
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v258 = int32(1)
	v259 = v246 + v258
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	if v260 != 0 {
		v245 = v245 + v258
		v246 = v259
		v247 = v255
		v248 = v260
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	v283 = v237
	goto L81
L80:
	;
	v283 = v238
	goto L81
L81:
	;
	v286 = v283
	v287 = int64(1)
	goto L45
L82:
	;
	v288 = v60
	goto L84
L83:
	;
	v288 = v50
	goto L84
L84:
	;
	v290 = v149 + v286
	goto L85
L85:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v311-int32(9)))&base.B2i32(v311 != int32(32)) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v323 = int32(0)
	if v100 == v323 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v290 = v290 + int32(1)
	goto L85
L88:
	;
	goto L89
L89:
	;
	goto L86
L90:
	;
	if v367 != 0 {
		goto L104
	} else {
		goto L105
	}
L91:
	;
	v367 = int32(0)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if v329 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v330 = v290
	v331 = v99
	v332 = v100
	v333 = v329
	goto L98
L95:
	;
	v355 = v99
	v359 = int32(0)
	goto L96
L96:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	v367 = v359 - v360
	goto L90
L97:
	;
	v355 = v350
	v359 = v352
	goto L96
L98:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v333 != v335 {
		v350 = v331
		v352 = v333
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v350 = v344
	v352 = int32(0)
	goto L97
L100:
	;
	if v335 == int32(0) {
		v350 = v331
		v352 = v333
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v340 = v332 - int32(1)
	if v340 == int32(0) {
		v350 = v331
		v352 = v333
		goto L97
	} else {
		goto L102
	}
L102:
	;
	v343 = int32(1)
	v344 = v331 + v343
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v345 != 0 {
		v330 = v330 + v343
		v331 = v344
		v332 = v340
		v333 = v345
		goto L98
	} else {
		goto L103
	}
L103:
	;
	goto L99
L104:
	;
	v368 = v323
	goto L106
L105:
	;
	v368 = v100
	goto L106
L106:
	;
	v370 = v290 + v368
	goto L109
L107:
	;
	m.G0 = v24 + int32(112)
	return v992
L108:
	;
	if v625 < v399 {
		goto L160
	} else {
		goto L161
	}
L109:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
	if base.Ui32(v391-int32(9)) < base.Ui32(int32(5)) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v403 = v370
	v404 = v391
	v406 = int32(0)
	v419 = v17
	v420 = v17
	goto L116
L111:
	;
	goto L110
L112:
	;
	v370 = v370 + int32(1)
	goto L109
L113:
	;
	if v391 == int32(32) {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v399 = base.I64_extend8_s(base.I64_extend_i32_u(v53))
	if v391 != 0 {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v608 = v370
	v624 = v17
	v625 = v17
	goto L108
L116:
	;
	v425 = v404 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v425&int32(255)) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	if base.Ui32(int32(4)) < base.Ui32((v404-int32(53))&int32(255)) {
		v608 = v403
		v624 = v419
		v625 = v420
		goto L108
	} else {
		goto L153
	}
L118:
	;
	goto L117
L119:
	;
	v576 = v569 + int32(1)
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	if v577 != 0 {
		v403 = v576
		v404 = v577
		v406 = v570
		v419 = v572
		v420 = v573
		goto L116
	} else {
		goto L152
	}
L120:
	;
	if (base.B2i32(v51 != base.I32_extend8_s(v404))|v406)&int32(1) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L121:
	;
	if v406&base.B2i32(v399 <= v420) != 0 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v433 = v24 + int32(96)
	v436 = int64(10)
	v437 = int64(0)
	v442 = int64(32)
	v445 = int64(base.Ui64(v419) >> (uint(v442) % 64))
	v448 = int64(4294967295)
	v451 = v419 & v448
	v452 = v436 * v451
	v456 = int64(base.Ui64(v452)>>(uint(v442)%64)) + v436*v445
	v463 = v451*v437 + v456&v448
	*(*int64)(unsafe.Add(mBase, uint32(v433)+8)) = v419*v437 + v419>>(uint(int64(63))%64)*v436 + v437*v445 + int64(base.Ui64(v456)>>(uint(v442)%64)) + int64(base.Ui64(v463)>>(uint(v442)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v433))) = v452&v448 | v463<<(uint(v442)%64)
	goto L123
L123:
	;
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v24)+104))
	v475 = *(*int64)(unsafe.Add(mBase, uint32(v24)+96))
	if v474 == v475>>(uint(int64(63))%64) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v569 = v403
	v570 = v406
	v572 = v484
	v573 = v420 + base.I64_extend_i32_u(v406)&int64(1)
	goto L119
L125:
	;
	v481 = base.I64_extend_i32_u(v425) & int64(255)
	v484 = v475 - v481
	if base.B2i32(v481 != int64(0)) == base.B2i32(v484 < v475) {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v488 = int32(0)
	v489 = F_errsave_start(m, v26)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	if v489 == int32(0) {
		v992 = v488
		goto L107
	} else {
		goto L130
	}
L130:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = int32(20987)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v27
	F_errmsg(m, int32(190444), v24+int32(80))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errsave_finish(m, v26, int32(499393), int32(293), int32(280388))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v992 = v488
	goto L107
L134:
	;
	v569 = v403
	v570 = int32(1)
	v572 = v419
	v573 = v420
	goto L119
L135:
	;
	goto L136
L136:
	;
	v521 = F_strlen(m, v288)
	mBase = m.M
	if v521 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v565 != 0 {
		goto L118
	} else {
		goto L151
	}
L138:
	;
	v565 = int32(0)
	goto L137
L139:
	;
	goto L140
L140:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v527 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v528 = v403
	v529 = v288
	v530 = v521
	v531 = v527
	goto L145
L142:
	;
	v553 = v288
	v557 = int32(0)
	goto L143
L143:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553))))
	v565 = v557 - v558
	goto L137
L144:
	;
	v553 = v548
	v557 = v550
	goto L143
L145:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	if v531 != v533 {
		v548 = v529
		v550 = v531
		goto L144
	} else {
		goto L147
	}
L146:
	;
	v548 = v542
	v550 = int32(0)
	goto L144
L147:
	;
	if v533 == int32(0) {
		v548 = v529
		v550 = v531
		goto L144
	} else {
		goto L148
	}
L148:
	;
	v538 = v530 - int32(1)
	if v538 == int32(0) {
		v548 = v529
		v550 = v531
		goto L144
	} else {
		goto L149
	}
L149:
	;
	v541 = int32(1)
	v542 = v529 + v541
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+1)))
	if v543 != 0 {
		v528 = v528 + v541
		v529 = v542
		v530 = v538
		v531 = v543
		goto L145
	} else {
		goto L150
	}
L150:
	;
	goto L146
L151:
	;
	v569 = v403 + v521 - int32(1)
	v570 = v406
	v572 = v419
	v573 = v420
	goto L119
L152:
	;
	v608 = v576
	v624 = v572
	v625 = v573
	goto L108
L153:
	;
	v585 = v419 - int64(1)
	if v585 < v419 {
		v608 = v403
		v624 = v585
		v625 = v420
		goto L108
	} else {
		goto L154
	}
L154:
	;
	v587 = int32(0)
	v588 = F_errsave_start(m, v26)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	if v588 == int32(0) {
		v992 = v587
		goto L107
	} else {
		goto L156
	}
L156:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = int32(20987)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v27
	F_errmsg(m, int32(190444), v24-int32(-64))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errsave_finish(m, v26, int32(499393), int32(318), int32(280388))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v992 = v587
	goto L107
L160:
	;
	v646 = v624
	v647 = v625
	goto L163
L161:
	;
	v736 = v624
	goto L162
L162:
	;
	v741 = v608
	goto L175
L163:
	;
	v652 = v24 + int32(48)
	v655 = int64(10)
	v656 = int64(0)
	v661 = int64(32)
	v664 = int64(base.Ui64(v646) >> (uint(v661) % 64))
	v667 = int64(4294967295)
	v670 = v646 & v667
	v671 = v655 * v670
	v675 = int64(base.Ui64(v671)>>(uint(v661)%64)) + v655*v664
	v682 = v670*v656 + v675&v667
	*(*int64)(unsafe.Add(mBase, uint32(v652)+8)) = v646*v656 + v646>>(uint(int64(63))%64)*v655 + v656*v664 + int64(base.Ui64(v675)>>(uint(v661)%64)) + int64(base.Ui64(v682)>>(uint(v661)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v652))) = v671&v667 | v682<<(uint(v661)%64)
	goto L165
L164:
	;
	v736 = v694
	goto L162
L165:
	;
	v693 = *(*int64)(unsafe.Add(mBase, uint32(v24)+56))
	v694 = *(*int64)(unsafe.Add(mBase, uint32(v24)+48))
	if v693 != v694>>(uint(int64(63))%64) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v698 = int32(0)
	v699 = F_errsave_start(m, v26)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v718 = v647 + int64(1)
	if v718 != v399 {
		v646 = v694
		v647 = v718
		goto L163
	} else {
		goto L174
	}
L169:
	;
	if v699 == int32(0) {
		v992 = v698
		goto L107
	} else {
		goto L170
	}
L170:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(20987)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27
	F_errmsg(m, int32(190444), v24)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errsave_finish(m, v26, int32(499393), int32(328), int32(280388))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v992 = v698
	goto L107
L174:
	;
	goto L164
L175:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741))))
	if base.Ui32((v764-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v741 = v741 + int32(1)
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v772 = v741
	v774 = v764
	v790 = v287
	goto L178
L177:
	;
	goto L176
L178:
	;
	switch v774 & int32(255) {
	case 0:
		goto L181
	default:
		goto L182
	case 9, 10, 11, 12, 13, 32, 41:
		v988 = int32(1)
		v989 = v790
		goto L180
	}
L180:
	;
	v990 = v988 + v772
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	v772 = v990
	v774 = v991
	v790 = v989
	goto L178
L181:
	;
	if int64(0) < v790 {
		goto L235
	} else {
		goto L236
	}
L182:
	;
	v795 = F_strlen(m, v186)
	mBase = m.M
	if v795 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v839 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L184:
	;
	v839 = int32(0)
	goto L183
L185:
	;
	goto L186
L186:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772))))
	if v801 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v802 = v772
	v803 = v186
	v804 = v795
	v805 = v801
	goto L191
L188:
	;
	v827 = v186
	v831 = int32(0)
	goto L189
L189:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827))))
	v839 = v831 - v832
	goto L183
L190:
	;
	v827 = v822
	v831 = v824
	goto L189
L191:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
	if v805 != v807 {
		v822 = v803
		v824 = v805
		goto L190
	} else {
		goto L193
	}
L192:
	;
	v822 = v816
	v824 = int32(0)
	goto L190
L193:
	;
	if v807 == int32(0) {
		v822 = v803
		v824 = v805
		goto L190
	} else {
		goto L194
	}
L194:
	;
	v812 = v804 - int32(1)
	if v812 == int32(0) {
		v822 = v803
		v824 = v805
		goto L190
	} else {
		goto L195
	}
L195:
	;
	v815 = int32(1)
	v816 = v803 + v815
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802)+1)))
	if v817 != 0 {
		v802 = v802 + v815
		v803 = v816
		v804 = v812
		v805 = v817
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	v988 = v795
	v989 = int64(-1)
	goto L180
L198:
	;
	goto L199
L199:
	;
	v843 = F_strlen(m, v148)
	mBase = m.M
	if v843 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if v887 == int32(0) {
		v988 = v843
		v989 = v790
		goto L180
	} else {
		goto L214
	}
L201:
	;
	v887 = int32(0)
	goto L200
L202:
	;
	goto L203
L203:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772))))
	if v849 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v850 = v772
	v851 = v148
	v852 = v843
	v853 = v849
	goto L208
L205:
	;
	v875 = v148
	v879 = int32(0)
	goto L206
L206:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	v887 = v879 - v880
	goto L200
L207:
	;
	v875 = v870
	v879 = v872
	goto L206
L208:
	;
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	if v853 != v855 {
		v870 = v851
		v872 = v853
		goto L207
	} else {
		goto L210
	}
L209:
	;
	v870 = v864
	v872 = int32(0)
	goto L207
L210:
	;
	if v855 == int32(0) {
		v870 = v851
		v872 = v853
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v860 = v852 - int32(1)
	if v860 == int32(0) {
		v870 = v851
		v872 = v853
		goto L207
	} else {
		goto L212
	}
L212:
	;
	v863 = int32(1)
	v864 = v851 + v863
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+1)))
	if v865 != 0 {
		v850 = v850 + v863
		v851 = v864
		v852 = v860
		v853 = v865
		goto L208
	} else {
		goto L213
	}
L213:
	;
	goto L209
L214:
	;
	v890 = F_strlen(m, v99)
	mBase = m.M
	if v890 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	if v934 == int32(0) {
		v988 = v890
		v989 = v790
		goto L180
	} else {
		goto L229
	}
L216:
	;
	v934 = int32(0)
	goto L215
L217:
	;
	goto L218
L218:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772))))
	if v896 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v897 = v772
	v898 = v99
	v899 = v890
	v900 = v896
	goto L223
L220:
	;
	v922 = v99
	v926 = int32(0)
	goto L221
L221:
	;
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922))))
	v934 = v926 - v927
	goto L215
L222:
	;
	v922 = v917
	v926 = v919
	goto L221
L223:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898))))
	if v900 != v902 {
		v917 = v898
		v919 = v900
		goto L222
	} else {
		goto L225
	}
L224:
	;
	v917 = v911
	v919 = int32(0)
	goto L222
L225:
	;
	if v902 == int32(0) {
		v917 = v898
		v919 = v900
		goto L222
	} else {
		goto L226
	}
L226:
	;
	v907 = v899 - int32(1)
	if v907 == int32(0) {
		v917 = v898
		v919 = v900
		goto L222
	} else {
		goto L227
	}
L227:
	;
	v910 = int32(1)
	v911 = v898 + v910
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897)+1)))
	if v912 != 0 {
		v897 = v897 + v910
		v898 = v911
		v899 = v907
		v900 = v912
		goto L223
	} else {
		goto L228
	}
L228:
	;
	goto L224
L229:
	;
	v937 = int32(0)
	v938 = F_errsave_start(m, v26)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	if v938 == int32(0) {
		v992 = v937
		goto L107
	} else {
		goto L231
	}
L231:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(20987)
	F_errmsg(m, int32(727181), v24+int32(16))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	F_errsave_finish(m, v26, int32(499393), int32(355), int32(280388))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v992 = v937
	goto L107
L235:
	;
	if v736 == int64(-9223372036854775807-1) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	v985 = v736
	goto L237
L237:
	;
	v986 = F_Int64GetDatum(m, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L246
	}
L238:
	;
	v962 = int32(0)
	v963 = F_errsave_start(m, v26)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v985 = int64(0) - v736
	goto L237
L241:
	;
	if v963 == int32(0) {
		v992 = v962
		goto L107
	} else {
		goto L242
	}
L242:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = int32(20987)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v27
	F_errmsg(m, int32(190444), v24+int32(32))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_errsave_finish(m, v26, int32(499393), int32(368), int32(280388))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v992 = v962
	goto L107
L246:
	;
	v992 = v986
	goto L107
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
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v39 int32
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
	F_appendStringInfoString(m, v9, int32(735596))
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
	v30 = base.I64_div_u_s(v24, int64(100000000000000))
	v33 = base.I32_rem_u_s(base.I32_wrap_i64(v30), int32(1000))
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
	v37 = base.I64_div_u_s(v24, int64(100000000000000000))
	F_append_num_word(m, v9, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_appendStringInfoString(m, v9, int32(741649))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_append_num_word(m, v9, base.I64_extend_i32_u(v33))
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
	F_appendStringInfoString(m, v9, int32(741638))
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
	F_appendStringInfoString(m, v9, int32(741673))
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
	F_appendStringInfoString(m, v9, int32(741663))
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
	F_appendStringInfoString(m, v9, int32(745423))
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
	F_appendStringInfoString(m, v9, int32(240991))
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
	v97 = int32(745514)
	goto L42
L41:
	;
	v97 = int32(745500)
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
	v109 = int32(97253)
	goto L47
L46:
	;
	v109 = int32(123797)
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
