package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_internal_load_library(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int64
	_ = v275
	var v279 int64
	_ = v279
	var v283 int64
	_ = v283
	var v287 int64
	_ = v287
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int64
	_ = v749
	var v751 int64
	_ = v751
	var v753 int64
	_ = v753
	var v755 int64
	_ = v755
	var v757 int64
	_ = v757
	var v761 int64
	_ = v761
	var v763 int64
	_ = v763
	var v765 int64
	_ = v765
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	v13 = m.G0
	v15 = v13 - int32(208)
	m.G0 = v15
	v20 = int32(4443280)
	goto L7
L1:
	;
	F_emscripten_builtin_free(m, v157)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L163
	} else {
		goto L285
	}
L2:
	;
	v749 = *(*int64)(unsafe.Add(mBase, uint32(v585)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+104)) = v749
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v585)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v751
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v585)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v753
	v755 = *(*int64)(unsafe.Add(mBase, uint32(v585)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = v755
	v757 = *(*int64)(unsafe.Add(mBase, uint32(v585)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v757
	v761 = *(*int64)(unsafe.Add(mBase, uint32(v585)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v761
	v763 = *(*int64)(unsafe.Add(mBase, uint32(v585)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v763
	v765 = *(*int64)(unsafe.Add(mBase, uint32(v585)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = v765
	F_emscripten_builtin_free(m, v157)
	mBase = m.M
	v769 = m.G0
	v771 = v769 - int32(224)
	m.G0 = v771
	v776 = v15 + int32(48) | int32(4)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	if v777 != int32(1800) {
		goto L213
	} else {
		goto L214
	}
L3:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, _consts[925])))
	if v723 != 0 {
		goto L206
	} else {
		goto L207
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L163
	} else {
		goto L202
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L163
	} else {
		goto L198
	}
L6:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v676)+16))
	m.G0 = v15 + int32(208)
	return v686
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v62 = F___fstatat(m, int32(-100), l0, v15+int32(112), int32(0))
	mBase = m.M
	goto L21
L9:
	;
	v32 = v30 + int32(24)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	if v56-v55 != 0 {
		v20 = v30
		goto L7
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = l0
	v41 = v32
	goto L16
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v55 = v44
	v56 = v45
	goto L13
L18:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v676 = v30
	goto L6
L21:
	;
	if v62 == int32(-1) {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[926]))
	if v66 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	v71 = v66
	goto L26
L24:
	;
	goto L25
L25:
	;
	if l0&int32(3) == int32(0) {
		v121 = l0
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
	if v81 == v68 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v67 == v83 {
		v676 = v71
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v85 != 0 {
		v71 = v85
		goto L26
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L27
L33:
	;
	v157 = F_emscripten_builtin_malloc(m, v154+int32(25))
	mBase = m.M
	if v157 == int32(0) {
		goto L4
	} else {
		goto L50
	}
L34:
	;
	v154 = v146 - l0
	goto L33
L35:
	;
	v125 = v121
	goto L44
L36:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v105 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v154 = int32(0)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v110 = l0
	goto L40
L40:
	;
	v114 = v110 + int32(1)
	if v114&int32(3) == int32(0) {
		v121 = v114
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v146 = v114
	goto L34
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v119 != 0 {
		v110 = v114
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v134 = int32(-2139062144)
	if (int32(16843008)-v131|v131)&v134 == v134 {
		v125 = v125 + int32(4)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v140 = v125
	goto L47
L46:
	;
	goto L45
L47:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v144 != 0 {
		v140 = v140 + int32(1)
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v146 = v140
	goto L34
L49:
	;
	goto L48
L50:
	;
	if v157&int32(3) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v186 = v157 + int32(24)
	if (l0^v186)&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L52:
	;
	v165 = v157 + int32(24)
	if base.Ui32(v165) <= base.Ui32(v157) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157)+16)) = int64(0)
	goto L51
L55:
	;
	v171 = v157 + int32(4)
	if base.Ui32(v171) < base.Ui32(v165) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v173 = v165
	goto L58
L57:
	;
	v173 = v171
	goto L58
L58:
	;
	v180 = F__emscripten_memset_bulkmem(m, v157, base.I32_extend8_s(int32(0)), (v157^int32(-1)+v173)&int32(-4)+int32(4))
	mBase = m.M
	goto L59
L59:
	;
	goto L51
L60:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v261
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v15)+200))
	v264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = v263
	v267 = m.G0
	v269 = v267 - int32(16)
	m.G0 = v269
	if v186 == v264 {
		goto L83
	} else {
		goto L84
	}
L61:
	;
	goto L60
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v240)
	if v240&int32(255) == int32(0) {
		goto L61
	} else {
		goto L77
	}
L63:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v239 = l0
	v240 = v192
	v241 = v186
	goto L62
L64:
	;
	goto L65
L65:
	;
	if l0&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v196 = l0
	v198 = v186
	goto L69
L67:
	;
	v210 = l0
	v212 = v186
	goto L68
L68:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v217 = int32(-2139062144)
	if (int32(16843008)-v214|v214)&v217 != v217 {
		v239 = v210
		v240 = v214
		v241 = v212
		goto L62
	} else {
		goto L73
	}
L69:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v199)
	if v199 == int32(0) {
		goto L61
	} else {
		goto L71
	}
L70:
	;
	v210 = v206
	v212 = v204
	goto L68
L71:
	;
	v203 = int32(1)
	v204 = v198 + v203
	v206 = v196 + v203
	if v206&int32(3) != 0 {
		v196 = v206
		v198 = v204
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v222 = v210
	v223 = v214
	v224 = v212
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v223
	v226 = int32(4)
	v227 = v224 + v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v230 = v222 + v226
	v234 = int32(-2139062144)
	if (v228|(int32(16843008)-v228))&v234 == v234 {
		v222 = v230
		v223 = v228
		v224 = v227
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v239 = v230
	v240 = v228
	v241 = v227
	goto L62
L76:
	;
	goto L75
L77:
	;
	v248 = v239
	v250 = v241
	goto L78
L78:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)) = uint8(v251)
	v253 = int32(1)
	if v251 != 0 {
		v248 = v248 + v253
		v250 = v250 + v253
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L61
L80:
	;
	goto L79
L81:
	;
	m.G0 = v269 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = v570
	if v570 == int32(0) {
		goto L3
	} else {
		goto L165
	}
L82:
	;
	v560 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[925])) = uint8(v560)
	v570 = int32(0)
	goto L81
L83:
	;
	v275 = *(*int64)(unsafe.Add(mBase, _consts[927]))
	*(*int64)(unsafe.Add(mBase, _consts[928])) = v275
	v279 = *(*int64)(unsafe.Add(mBase, _consts[929]))
	*(*int64)(unsafe.Add(mBase, _consts[930])) = v279
	v283 = *(*int64)(unsafe.Add(mBase, _consts[931]))
	*(*int64)(unsafe.Add(mBase, _consts[932])) = v283
	v287 = *(*int64)(unsafe.Add(mBase, _consts[933]))
	*(*int64)(unsafe.Add(mBase, _consts[934])) = v287
	goto L82
L84:
	;
	goto L85
L85:
	;
	v292 = F_strlen(m, v186)
	mBase = m.M
	v299 = v292 + int32(1)
	goto L89
L86:
	;
	v397 = int32(0)
	v399 = *(*int32)(unsafe.Add(mBase, _consts[935]))
	if v397 < v399 {
		goto L122
	} else {
		goto L123
	}
L87:
	;
	if v311 != 0 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	v301 = int32(0)
	if v299 == v301 {
		v311 = v301
		goto L88
	} else {
		goto L91
	}
L90:
	;
	v311 = v306
	goto L88
L91:
	;
	v305 = v299 - int32(1)
	v306 = v186 + v305
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v307 != int32(47) {
		v299 = v305
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v314 = v311 + int32(1)
	goto L95
L94:
	;
	v314 = v186
	goto L95
L95:
	;
	v318 = F_strlen(m, v314)
	mBase = m.M
	v325 = v318 + int32(1)
	goto L98
L96:
	;
	if v337 != 0 {
		goto L102
	} else {
		goto L103
	}
L97:
	;
	goto L96
L98:
	;
	v327 = int32(0)
	if v325 == v327 {
		v337 = v327
		goto L97
	} else {
		goto L100
	}
L99:
	;
	v337 = v332
	goto L97
L100:
	;
	v331 = v325 - int32(1)
	v332 = v314 + v331
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v333 != int32(46) {
		v325 = v331
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v396 = v337 - v314
	goto L86
L103:
	;
	goto L104
L104:
	;
	if v314&int32(3) == int32(0) {
		v362 = v314
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v396 = v395
	goto L86
L106:
	;
	v395 = v387 - v314
	goto L105
L107:
	;
	v366 = v362
	goto L116
L108:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v346 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v395 = int32(0)
	goto L105
L110:
	;
	goto L111
L111:
	;
	v351 = v314
	goto L112
L112:
	;
	v355 = v351 + int32(1)
	if v355&int32(3) == int32(0) {
		v362 = v355
		goto L107
	} else {
		goto L114
	}
L113:
	;
	v387 = v355
	goto L106
L114:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
	if v360 != 0 {
		v351 = v355
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	v375 = int32(-2139062144)
	if (int32(16843008)-v372|v372)&v375 == v375 {
		v366 = v366 + int32(4)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v381 = v366
	goto L119
L118:
	;
	goto L117
L119:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	if v385 != 0 {
		v381 = v381 + int32(1)
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v387 = v381
	goto L106
L121:
	;
	goto L120
L122:
	;
	v406 = v397
	goto L125
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v186
	v542 = F_snprintf(m, int32(4542192), int32(512), int32(414230), v269)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L163
	} else {
		goto L164
	}
L125:
	;
	v415 = v406 * int32(12)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v415)+uint32(_consts[936])))
	if v418&int32(3) == int32(0) {
		v442 = v418
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L124
L127:
	;
	if v475 == v396 {
		goto L144
	} else {
		goto L145
	}
L128:
	;
	v475 = v467 - v418
	goto L127
L129:
	;
	v446 = v442
	goto L138
L130:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	if v426 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v475 = int32(0)
	goto L127
L132:
	;
	goto L133
L133:
	;
	v431 = v418
	goto L134
L134:
	;
	v435 = v431 + int32(1)
	if v435&int32(3) == int32(0) {
		v442 = v435
		goto L129
	} else {
		goto L136
	}
L135:
	;
	v467 = v435
	goto L128
L136:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v440 != 0 {
		v431 = v435
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v455 = int32(-2139062144)
	if (int32(16843008)-v452|v452)&v455 == v455 {
		v446 = v446 + int32(4)
		goto L138
	} else {
		goto L140
	}
L139:
	;
	v461 = v446
	goto L141
L140:
	;
	goto L139
L141:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v465 != 0 {
		v461 = v461 + int32(1)
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v467 = v461
	goto L128
L143:
	;
	goto L142
L144:
	;
	if v396 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L146
L146:
	;
	v524 = v406 + int32(1)
	if v524 != v399 {
		v406 = v524
		goto L125
	} else {
		goto L162
	}
L147:
	;
	if v520 == int32(0) {
		v570 = v415 + int32(2116448)
		goto L81
	} else {
		goto L161
	}
L148:
	;
	v520 = int32(0)
	goto L147
L149:
	;
	goto L150
L150:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	if v482 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v483 = v418
	v484 = v314
	v485 = v396
	v486 = v482
	goto L155
L152:
	;
	v508 = v314
	v512 = int32(0)
	goto L153
L153:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	v520 = v512 - v513
	goto L147
L154:
	;
	v508 = v503
	v512 = v505
	goto L153
L155:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	if v486 != v488 {
		v503 = v484
		v505 = v486
		goto L154
	} else {
		goto L157
	}
L156:
	;
	v503 = v497
	v505 = int32(0)
	goto L154
L157:
	;
	if v488 == int32(0) {
		v503 = v484
		v505 = v486
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v493 = v485 - int32(1)
	if v493 == int32(0) {
		v503 = v484
		v505 = v486
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v496 = int32(1)
	v497 = v484 + v496
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
	if v498 != 0 {
		v483 = v483 + v496
		v484 = v497
		v485 = v493
		v486 = v498
		goto L155
	} else {
		goto L160
	}
L160:
	;
	goto L156
L161:
	;
	goto L146
L162:
	;
	goto L126
L163:
	;
	return int32(0)
L164:
	;
	goto L82
L165:
	;
	v581 = F_pgmem_dlsym(m, v570, int32(471229))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L163
	} else {
		goto L166
	}
L166:
	;
	if v581 == int32(0) {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v585 = m.T0[v581].(func(*base.Module) int32)(m)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L163
	} else {
		goto L168
	}
L168:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v585)))
	if v587 != int32(64) {
		goto L2
	} else {
		goto L169
	}
L169:
	;
	v591 = v585 + int32(4)
	v592 = int32(1725416)
	v593 = int32(52)
	goto L173
L170:
	;
	if v655 != 0 {
		goto L2
	} else {
		goto L188
	}
L171:
	;
	v655 = int32(0)
	goto L170
L172:
	;
	v629 = v624
	v630 = v625
	v631 = v626
	goto L182
L173:
	;
	if (v591|v592)&int32(3) != 0 {
		v624 = v591
		v625 = v592
		v626 = v593
		goto L172
	} else {
		goto L176
	}
L175:
	;
	if v614 == int32(0) {
		goto L171
	} else {
		goto L181
	}
L176:
	;
	v601 = v591
	v602 = v592
	v603 = v593
	goto L177
L177:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	if v606 != v607 {
		v624 = v601
		v625 = v602
		v626 = v603
		goto L172
	} else {
		goto L179
	}
L178:
	;
	goto L175
L179:
	;
	v609 = int32(4)
	v610 = v602 + v609
	v612 = v601 + v609
	v614 = v603 - v609
	if base.Ui32(int32(3)) < base.Ui32(v614) {
		v601 = v612
		v602 = v610
		v603 = v614
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v624 = v612
	v625 = v610
	v626 = v614
	goto L172
L182:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629))))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	if v634 == v635 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v655 = v634 - v635
	goto L170
L184:
	;
	v637 = int32(1)
	v642 = v631 - v637
	if v642 != 0 {
		v629 = v629 + v637
		v630 = v630 + v637
		v631 = v642
		goto L182
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	goto L183
L187:
	;
	goto L171
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+20)) = v585
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v659 = F_pgmem_dlsym(m, v657, int32(95035))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L163
	} else {
		goto L189
	}
L189:
	;
	if v659 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	m.T0[v659].(func(*base.Module))(m)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L163
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v664 = *(*int32)(unsafe.Add(mBase, _consts[926]))
	if v664 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, _consts[937])) = v157
	v676 = v157
	goto L6
L195:
	;
	*(*int32)(unsafe.Add(mBase, _consts[926])) = v157
	goto L194
L196:
	;
	goto L197
L197:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _consts[937]))
	*(*int32)(unsafe.Add(mBase, uint32(v670))) = v157
	goto L194
L198:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L163
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(285733), v15)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L163
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(476042), int32(215), int32(16438))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L163
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L163
	} else {
		goto L203
	}
L203:
	;
	F_errmsg(m, int32(12915), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L163
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(476042), int32(234), int32(16438))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L163
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v725 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[925])) = uint8(v725)
	v729 = int32(4542192)
	goto L208
L207:
	;
	v729 = int32(0)
	goto L208
L208:
	;
	F_emscripten_builtin_free(m, v157)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L163
	} else {
		goto L209
	}
L209:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L163
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(196538), v15+int32(16))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L163
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(476042), int32(253), int32(16438))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L163
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	if int32(1000) <= v777 {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	goto L215
L215:
	;
	v835 = v776 + int32(20)
	v836 = int32(1725436)
	v839 = int32(*(*uint8)(unsafe.Add(mBase, _consts[938])))
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835))))
	if v840 == int32(0) {
		v859 = v839
		v860 = v840
		goto L227
	} else {
		goto L228
	}
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L163
	} else {
		goto L222
	}
L217:
	;
	v783 = base.I32_div_u_s(v777, int32(100))
	*(*int32)(unsafe.Add(mBase, uint32(v771)+160)) = v783
	v791 = F_pg_snprintf(m, v771+int32(192), int32(32), int32(469640), v771+int32(160))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L163
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v793 = int32(100)
	v794 = base.I32_div_s(v777, v793)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+176)) = v794
	*(*int32)(unsafe.Add(mBase, uint32(v771)+180)) = v777 - v794*v793
	v806 = F_pg_snprintf(m, v771+int32(192), int32(32), int32(448820), v771+int32(176))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L163
	} else {
		goto L221
	}
L220:
	;
	goto L216
L221:
	;
	goto L216
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+144)) = l0
	F_errmsg(m, int32(311303), v771+int32(144))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L163
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+128)) = int32(18)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+132)) = v771 + int32(192)
	F_errdetail(m, int32(568453), v771+int32(128))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L163
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(476042), int32(340), int32(203134))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L163
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	if v860-v859 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L227:
	;
	goto L226
L228:
	;
	if v839 != v840 {
		v859 = v839
		v860 = v840
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v844 = v835
	v845 = v836
	goto L230
L230:
	;
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845)+1)))
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844)+1)))
	if v849 == int32(0) {
		v859 = v848
		v860 = v849
		goto L227
	} else {
		goto L232
	}
L231:
	;
	v859 = v848
	v860 = v849
	goto L227
L232:
	;
	v852 = int32(1)
	if v848 == v849 {
		v844 = v844 + v852
		v845 = v845 + v852
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	F_initStringInfo(m, v771+int32(192))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L163
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L163
	} else {
		goto L281
	}
L237:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	if v868 != int32(100) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v771)+196))
	if v871 != 0 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	goto L240
L240:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v776)+8))
	if v891 != int32(32) {
		goto L246
	} else {
		goto L247
	}
L241:
	;
	F_appendStringInfoChar(m, v771+int32(192), int32(10))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L163
	} else {
		goto L244
	}
L242:
	;
	v878 = v868
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+88)) = v878
	*(*int32)(unsafe.Add(mBase, uint32(v771)+84)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+80)) = int32(503457)
	F_appendStringInfo(m, v771+int32(192), int32(616477), v771+int32(80))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L163
	} else {
		goto L245
	}
L244:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v878 = v877
	goto L243
L245:
	;
	goto L240
L246:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v771)+196))
	if v894 != 0 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L248
L248:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v776)+12))
	if v914 != int32(64) {
		goto L254
	} else {
		goto L255
	}
L249:
	;
	F_appendStringInfoChar(m, v771+int32(192), int32(10))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L163
	} else {
		goto L252
	}
L250:
	;
	v901 = v891
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+72)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v771)+68)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+64)) = int32(502533)
	F_appendStringInfo(m, v771+int32(192), int32(616477), v771-int32(-64))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L163
	} else {
		goto L253
	}
L252:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v776)+8))
	v901 = v900
	goto L251
L253:
	;
	goto L248
L254:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v771)+196))
	if v917 != 0 {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	goto L256
L256:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v776)+16))
	if v937 != 0 {
		goto L262
	} else {
		goto L263
	}
L257:
	;
	F_appendStringInfoChar(m, v771+int32(192), int32(10))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L163
	} else {
		goto L260
	}
L258:
	;
	v924 = v914
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+56)) = v924
	*(*int32)(unsafe.Add(mBase, uint32(v771)+52)) = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+48)) = int32(510858)
	F_appendStringInfo(m, v771+int32(192), int32(616477), v771+int32(48))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L163
	} else {
		goto L261
	}
L260:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v776)+12))
	v924 = v923
	goto L259
L261:
	;
	goto L256
L262:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v771)+196))
	if v938 != 0 {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	goto L264
L264:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v771)+196))
	if v962 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L265:
	;
	F_appendStringInfoChar(m, v771+int32(192), int32(10))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L163
	} else {
		goto L268
	}
L266:
	;
	v949 = int32(330316)
	goto L267
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+40)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v771)+36)) = int32(347024)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+32)) = int32(513953)
	F_appendStringInfo(m, v771+int32(192), int32(568071), v771+int32(32))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L163
	} else {
		goto L272
	}
L268:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v776)+16))
	if v946 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v947 = int32(330316)
	goto L271
L270:
	;
	v947 = int32(347024)
	goto L271
L271:
	;
	v949 = v947
	goto L267
L272:
	;
	goto L264
L273:
	;
	F_appendStringInfoString(m, v771+int32(192), int32(607152))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L163
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L163
	} else {
		goto L277
	}
L276:
	;
	goto L275
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+16)) = l0
	F_errmsg(m, int32(311374), v771+int32(16))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L163
	} else {
		goto L278
	}
L278:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v771)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = v980
	F_errdetail_internal(m, int32(197405), v771)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L163
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(476042), int32(414), int32(203134))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L163
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+112)) = l0
	F_errmsg(m, int32(311500), v771+int32(112))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L163
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+100)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v771)+96)) = int32(1725436)
	F_errdetail(m, int32(627911), v771+int32(96))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L163
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(476042), int32(355), int32(203134))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L163
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	F_errmsg(m, int32(303950), v15+int32(32))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L163
	} else {
		goto L286
	}
L286:
	;
	F_errhint(m, int32(575809), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L163
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(476042), int32(291), int32(16438))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L163
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
