package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyToBinaryOutFunc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_getTypeBinaryOutputInfo(m, l1, v6+int32(12), v6+int32(11))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		F_fmgr_info(m, v14, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_SplitToVariants(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v435 int32
	_ = v435
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v572 int32
	_ = v572
	var v583 int32
	_ = v583
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v636 int32
	_ = v636
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v798 int32
	_ = v798
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v963 int32
	_ = v963
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	v27 = m.G0
	v29 = v27 - int32(256)
	m.G0 = v29
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = v33
	v35 = l5
	goto L3
L2:
	;
	v34 = l1
	v35 = l6
	goto L3
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v40 = F_palloc(m, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v44 = F__emscripten_memset_bulkmem(m, v40, base.I32_extend8_s(int32(1)), l4)
	mBase = m.M
	goto L7
L7:
	;
	v46 = F_palloc(m, int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = int32(0)
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if l4 <= v35 {
		v911 = l5
		v913 = v46
		v921 = v29
		v923 = v44
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v50
	v54 = F_palloc(m, v50<<(uint(int32(2))%32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(16)
	v105 = F_palloc(m, int32(64))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L19
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v57
	if v57 <= int32(0) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v70 = int32(0)
	goto L15
L15:
	;
	v89 = v70 << (uint(int32(2)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89+v90)))
	v93 = F_pstrdup(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L9
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95+v89))) = v93
	v99 = v70 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v99 < v100 {
		v70 = v99
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v105
	goto L9
L20:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	*(*int32)(unsafe.Add(mBase, uint32(v963))) = v982 + int32(1)
	F_pfree(m, v973)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L4
	} else {
		goto L171
	}
L21:
	;
	v934 = F_pnstrdup(m, l3+v911, l4-v911)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L4
	} else {
		goto L165
	}
L22:
	;
	v137 = int32(1)
	v138 = l4 - v137
	v142 = v34
	v146 = l5
	v152 = v35
	goto L23
L23:
	;
	v167 = l3 + v146
	v169 = v142
	v179 = v152
	goto L25
L24:
	;
	v911 = v146
	v913 = v46
	v921 = v29
	v923 = v44
	goto L21
L25:
	;
	if v179 <= v146 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L24
L27:
	;
	goto L26
L28:
	;
	v693 = v169 + int32(4)
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v697 = v693 + v694<<(uint(int32(3))%32)
	if base.Ui32(v697) <= base.Ui32(v693) {
		v876 = int32(0)
		goto L120
	} else {
		goto L121
	}
L29:
	;
	if v169 == int32(0) {
		goto L27
	} else {
		goto L119
	}
L30:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v195 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v179 == v138 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v201 = int32(8)
	goto L34
L33:
	;
	v201 = int32(4)
	goto L34
L34:
	;
	if v179 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v203 = v201
	goto L37
L36:
	;
	v203 = int32(2)
	goto L37
L37:
	;
	v207 = l4 - v179
	v208 = l3 + v179
	v217 = v195
	goto L38
L38:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v169 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L29
L40:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+8)))
	if v381 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L41:
	;
	v362 = v248
	v364 = v266 - v208 + v268
	goto L40
L42:
	;
	if v235 == int32(0) {
		goto L27
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v235 == int32(0) {
		goto L28
	} else {
		goto L53
	}
L45:
	;
	v242 = v235
	v248 = v217
	goto L46
L46:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v266 < v207 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L29
L48:
	;
	v268 = F_strstr(m, v208, v242)
	mBase = m.M
	if v268 != 0 {
		goto L41
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v271 = v248 + int32(12)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 != 0 {
		v242 = v272
		v248 = v271
		goto L46
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	goto L47
L53:
	;
	v277 = v235
	v283 = v217
	goto L54
L54:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	if v301 < v207 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L29
L56:
	;
	if v301 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	goto L58
L58:
	;
	v350 = v283 + int32(12)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	if v351 != 0 {
		v277 = v351
		v283 = v350
		goto L54
	} else {
		goto L74
	}
L59:
	;
	if v346 == int32(0) {
		v362 = v283
		v364 = v301
		goto L40
	} else {
		goto L73
	}
L60:
	;
	v346 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v308 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v309 = v277
	v310 = v208
	v311 = v301
	v312 = v308
	goto L67
L64:
	;
	v334 = v208
	v338 = int32(0)
	goto L65
L65:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	v346 = v338 - v339
	goto L59
L66:
	;
	v334 = v329
	v338 = v331
	goto L65
L67:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v312 != v314 {
		v329 = v310
		v331 = v312
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v329 = v323
	v331 = int32(0)
	goto L66
L69:
	;
	if v314 == int32(0) {
		v329 = v310
		v331 = v312
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v319 = v311 - int32(1)
	if v319 == int32(0) {
		v329 = v310
		v331 = v312
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v322 = int32(1)
	v323 = v310 + v322
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v324 != 0 {
		v309 = v309 + v322
		v310 = v323
		v311 = v319
		v312 = v324
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	goto L58
L74:
	;
	goto L55
L75:
	;
	if v364 < int32(0) {
		goto L29
	} else {
		goto L78
	}
L76:
	;
	v386 = int32(0)
	goto L77
L77:
	;
	v387 = v386 + v179
	v388 = v44 - v137 + v387
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	if v389 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v386 = v364
	goto L77
L79:
	;
	v636 = v362 + int32(12)
	if v636 != 0 {
		v217 = v636
		goto L38
	} else {
		goto L118
	}
L80:
	;
	v392 = v386 + (v179 - v146)
	if v179-int32(1)+v392 <= l6 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	if int32(255) < v392 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	if int32(0) < v392 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v392 != 0 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	goto L85
L85:
	;
	v402 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v392+v29))) = uint8(v402)
	v404 = F_NormalizeSubWord(m, l0, v29, v203)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L90
	}
L86:
	;
	goto L85
L87:
	;
	v399 = F__emscripten_memcpy_bulkmem(m, v29, v167, v392)
	mBase = m.M
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L86
L90:
	;
	if v404 == int32(0) {
		goto L79
	} else {
		goto L91
	}
L91:
	;
	v410 = F_palloc(m, int32(16))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+12)) = int32(0)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v410)+4)) = v414
	v418 = F_palloc(m, v414<<(uint(int32(2))%32))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+8)) = v418
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v421
	if int32(0) < v421 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v435 = int32(0)
	goto L97
L95:
	;
	goto L96
L96:
	;
	v489 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v388))) = uint8(v489)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	if v491 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v452 = v435 << (uint(int32(2)) % 32)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v455+v452)))
	*(*int32)(unsafe.Add(mBase, uint32(v452+v453))) = v457
	v460 = v435 + int32(1)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v460 < v461 {
		v435 = v460
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	goto L98
L100:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v503 = v492
	v505 = v491
	v509 = v404
	goto L103
L101:
	;
	goto L102
L102:
	;
	F_pfree(m, v404)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L111
	}
L103:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	if v503 < v519 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533+v532<<(uint(int32(2))%32)))) = v505
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v540 = v538 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v410))) = v540
	v543 = v509 + int32(4)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	if v544 != 0 {
		v503 = v540
		v505 = v544
		v509 = v543
		goto L103
	} else {
		goto L110
	}
L106:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v532 = v503
	v533 = v521
	goto L105
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+4)) = v519 << (uint(int32(1)) % 32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v528 = F_repalloc(m, v525, v519<<(uint(int32(3))%32))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v410)+8)) = v528
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	v532 = v531
	v533 = v528
	goto L105
L110:
	;
	goto L104
L111:
	;
	v583 = v46
	goto L112
L112:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	if v599 != 0 {
		v583 = v599
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v601 = F_SplitToVariants(m, l0, int32(0), v410, l3, l4, v387, v387)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L4
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v583)+12)) = v601
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	F_pfree(m, v604)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_pfree(m, v410)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	goto L79
L118:
	;
	goto L39
L119:
	;
	goto L28
L120:
	;
	v878 = v179 + int32(1)
	if v878 < l4 {
		v169 = v876
		v179 = v878
		goto L25
	} else {
		goto L164
	}
L121:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v179))))
	v703 = v697
	v709 = v693
	goto L123
L122:
	;
	if v733&int32(256) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L123:
	;
	v732 = v709 + (v703-v709)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v735 = v733 & int32(255)
	if v735 == v700 {
		goto L122
	} else {
		goto L125
	}
L124:
	;
	v876 = int32(0)
	goto L120
L125:
	;
	v739 = base.B2i32(base.Ui32(v735) < base.Ui32(v700))
	if base.Ui32(v735) < base.Ui32(v700) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v740 = v732 + int32(8)
	goto L128
L127:
	;
	v740 = v709
	goto L128
L128:
	;
	if base.Ui32(v735) < base.Ui32(v700) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v741 = v703
	goto L131
L130:
	;
	v741 = v732
	goto L131
L131:
	;
	if base.Ui32(v740) < base.Ui32(v741) {
		v703 = v741
		v709 = v740
		goto L123
	} else {
		goto L132
	}
L132:
	;
	goto L124
L133:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	v876 = v849
	goto L120
L134:
	;
	if v179 == v138 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v753 = int32(8)
	goto L137
L136:
	;
	v753 = int32(4)
	goto L137
L137:
	;
	if v146 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v755 = v753
	goto L140
L139:
	;
	v755 = int32(2)
	goto L140
L140:
	;
	if int32(base.Ui32(v733)>>(uint(int32(9))%32))&v755 == int32(0) {
		goto L133
	} else {
		goto L141
	}
L141:
	;
	if v179 <= l6 {
		goto L133
	} else {
		goto L142
	}
L142:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v44))))
	if v761 == int32(0) {
		goto L133
	} else {
		goto L143
	}
L143:
	;
	v765 = v179 + int32(1)
	if v765 == l4 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v768 = F_pnstrdup(m, v167, l4-v146)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v798 = v46
	goto L153
L147:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v770 < v771 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v785+v784<<(uint(int32(2))%32)))) = v768
	v963 = v46
	v971 = v29
	v973 = v44
	goto L20
L149:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v784 = v770
	v785 = v773
	goto L148
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v771 << (uint(int32(1)) % 32)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v780 = F_repalloc(m, v777, v771<<(uint(int32(3))%32))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v780
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v784 = v783
	v785 = v780
	goto L148
L153:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v798)+12))
	if v816 != 0 {
		v798 = v816
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v817 = F_SplitToVariants(m, l0, v169, v46, l3, l4, v146, v179)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L156
	}
L155:
	;
	goto L154
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v798)+12)) = v817
	v821 = F_pnstrdup(m, v167, v765-v146)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v823 < v824 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v838+v837<<(uint(int32(2))%32)))) = v821
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v843 + int32(1)
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v765 < l4 {
		v142 = v847
		v146 = v765
		v152 = v765
		goto L23
	} else {
		goto L163
	}
L159:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v837 = v823
	v838 = v826
	goto L158
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v824 << (uint(int32(1)) % 32)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v833 = F_repalloc(m, v830, v824<<(uint(int32(3))%32))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v833
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v837 = v836
	v838 = v833
	goto L158
L163:
	;
	v911 = v765
	v913 = v46
	v921 = v29
	v923 = v44
	goto L21
L164:
	;
	goto L27
L165:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v913)))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v913)+4))
	if v936 < v937 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v951+v950<<(uint(int32(2))%32)))) = v934
	v963 = v913
	v971 = v921
	v973 = v923
	goto L20
L167:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v913)+8))
	v950 = v936
	v951 = v939
	goto L166
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913)+4)) = v937 << (uint(int32(1)) % 32)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v913)+8))
	v946 = F_repalloc(m, v943, v937<<(uint(int32(3))%32))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913)+8)) = v946
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v913)))
	v950 = v949
	v951 = v946
	goto L166
L171:
	;
	m.G0 = v971 + int32(256)
	return v963
}
func F_add_cast_to(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(82), l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_errmsg_internal(m, int32(49363), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(482839), int32(13488), int32(233999))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+68))
			v32 = F_get_namespace_name_or_temp(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v34 = F_quote_identifier(m, v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v38 = F_quote_identifier(m, v30+int32(4))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v38
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v34
						F_appendStringInfo(m, l0, int32(173450), v8+int32(16))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v11)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_assign_to(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = l1 - int32(8)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v8 < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = F_repalloc(m, v7, v5+int32(29))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = l1
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v28 == v29 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	return int32(0)
L5:
	;
	if v12 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_pfree(m, v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v5 + int32(20)
	v28 = v12 + int32(8)
	goto L3
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173-int32(4)))) = v5
	return v173
L11:
	;
	v173 = v28
	goto L10
L12:
	;
	v33 = v28 + v5
	if base.Ui32(v29-v33) <= base.Ui32(int32(0)-v5<<(uint(int32(1))%32)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = F___memcpy(m, v28, v29, v5)
	mBase = m.M
	v173 = v40
	goto L10
L14:
	;
	goto L15
L15:
	;
	v43 = (v28 ^ v29) & int32(3)
	if base.Ui32(v28) < base.Ui32(v29) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v145 == int32(0) {
		goto L11
	} else {
		goto L52
	}
L17:
	;
	if base.Ui32(v123) <= base.Ui32(int32(3)) {
		v144 = v122
		v145 = v123
		v146 = v124
		goto L16
	} else {
		goto L48
	}
L18:
	;
	if v43 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v43 != 0 {
		v105 = v5
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v144 = v29
	v145 = v5
	v146 = v28
	goto L16
L22:
	;
	goto L23
L23:
	;
	if v28&int32(3) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v122 = v29
	v123 = v5
	v124 = v28
	goto L17
L25:
	;
	goto L26
L26:
	;
	v50 = v29
	v51 = v5
	v52 = v28
	goto L27
L27:
	;
	if v51 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v122 = v59
	v123 = v61
	v124 = v63
	goto L17
L29:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v56)
	v58 = int32(1)
	v59 = v50 + v58
	v61 = v51 - v58
	v63 = v52 + v58
	if v63&int32(3) != 0 {
		v50 = v59
		v51 = v61
		v52 = v63
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v105 == int32(0) {
		goto L11
	} else {
		goto L44
	}
L32:
	;
	if v33&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v70 = v5
	goto L36
L34:
	;
	v85 = v5
	goto L35
L35:
	;
	if base.Ui32(v85) <= base.Ui32(int32(3)) {
		v105 = v85
		goto L31
	} else {
		goto L40
	}
L36:
	;
	if v70 == int32(0) {
		goto L11
	} else {
		goto L38
	}
L37:
	;
	v85 = v76
	goto L35
L38:
	;
	v76 = v70 - int32(1)
	v77 = v28 + v76
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v76))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v79)
	if v77&int32(3) != 0 {
		v70 = v76
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v92 = v85
	goto L41
L41:
	;
	v96 = v92 - int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v29+v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+v96))) = v99
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v92 = v96
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v105 = v96
	goto L31
L43:
	;
	goto L42
L44:
	;
	v112 = v105
	goto L45
L45:
	;
	v116 = v112 - int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28+v116))) = uint8(v119)
	if v116 != 0 {
		v112 = v116
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L11
L47:
	;
	goto L46
L48:
	;
	v129 = v122
	v130 = v123
	v131 = v124
	goto L49
L49:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133
	v135 = int32(4)
	v136 = v129 + v135
	v138 = v131 + v135
	v140 = v130 - v135
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		v129 = v136
		v130 = v140
		v131 = v138
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v144 = v136
	v145 = v140
	v146 = v138
	goto L16
L51:
	;
	goto L50
L52:
	;
	v151 = v144
	v152 = v145
	v153 = v146
	goto L53
L53:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v155)
	v157 = int32(1)
	v162 = v152 - v157
	if v162 != 0 {
		v151 = v151 + v157
		v152 = v162
		v153 = v153 + v157
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L11
L55:
	;
	goto L54
}
func F_coerce_to_specific_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_coerce_to_specific_type_typmod(m, l0, l1, l2, int32(-1), l3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_to_ascii_encname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_copy(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = m.G0
	v24 = v22 + int32(-64)
	m.G0 = v24
	v26 = int32(-1)
	if v14 == int32(0) {
		v101 = v26
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v101 < int32(0) {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	m.G0 = v24 - int32(-64)
	goto L3
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v29 == int32(0) {
		v101 = v26
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v32 = F_strlen(m, v14)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v32) {
		v101 = v26
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v35 = v14
	v36 = v29
	v37 = v24
	goto L8
L8:
	;
	v45 = F_isalnum(m, v36&int32(255))
	mBase = m.M
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v62)
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v24))))
	v67 = int32(1814800)
	v68 = int32(1814160)
	goto L17
L10:
	;
	if base.Ui32((v36-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v58 = v37
	goto L12
L12:
	;
	v60 = v35 + int32(1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v61 != 0 {
		v35 = v60
		v36 = v61
		v37 = v58
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v54 = v36 | int32(32)
	goto L15
L14:
	;
	v54 = v36
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v54)
	v58 = v37 + int32(1)
	goto L12
L16:
	;
	goto L9
L17:
	;
	v80 = v68 + (v67-v68)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81))))
	v83 = v66 - v82
	if v83 != 0 {
		v86 = v83
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v101 = v26
	goto L4
L19:
	;
	v90 = base.B2i32(v86 < int32(0))
	if v86 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v84 = F_strcmp(m, v24, v81)
	mBase = m.M
	if v84 != 0 {
		v86 = v84
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v101 = v85
	goto L4
L22:
	;
	v91 = v80 - int32(8)
	goto L24
L23:
	;
	v91 = v67
	goto L24
L24:
	;
	if v86 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = v68
	goto L27
L26:
	;
	v94 = v80 + int32(8)
	goto L27
L27:
	;
	if base.Ui32(v94) <= base.Ui32(v91) {
		v67 = v91
		v68 = v94
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L18
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v125 = F_encode_to_ascii(m, v10, v101)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L36
	}
L32:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
	F_errmsg(m, int32(372940), v7)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(486591), int32(128), int32(369583))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	m.G0 = v7 + int32(16)
	return v125
}
func F_to_hex64(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v14 = v9 - v8
	v15 = v14
	v16 = v12
	goto L1
L1:
	;
	v22 = v15 - int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)&int32(15))+uint32(_consts[371]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v28)
	if base.Ui64(v16) < base.Ui64(int64(16)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v14 - v22
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v22) {
		v15 = v22
		v16 = int64(base.Ui64(v16) >> (uint(int64(4)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v39
L9:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v39+int32(4), v22, v36)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_regcollation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1112]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1497), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regnamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1112]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1499), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regprocedure(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1112]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1252), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
