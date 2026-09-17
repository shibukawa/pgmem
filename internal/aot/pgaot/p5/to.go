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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v698 int32
	_ = v698
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v768 int32
	_ = v768
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	v26 = m.G0
	v28 = v26 - int32(256)
	m.G0 = v28
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = v32
	v34 = l5
	goto L3
L2:
	;
	v33 = l1
	v34 = l6
	goto L3
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v39 = F_palloc(m, l4)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	base.MemoryFill(m, v39, int32(1), l4)
	goto L9
L8:
	;
	goto L9
L9:
	;
	v44 = F_palloc(m, int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(0)
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if l4 <= v34 {
		v878 = l5
		v880 = v44
		v889 = v28
		v890 = v39
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
	v52 = F_palloc(m, v48<<(uint(int32(2))%32))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = int32(16)
	v102 = F_palloc(m, int32(64))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v55
	if v55 <= int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v68 = int32(0)
	goto L17
L17:
	;
	v86 = v68 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86+v87)))
	v90 = F_pstrdup(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L11
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v92+v86))) = v90
	v96 = v68 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v96 < v97 {
		v68 = v96
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v102
	goto L11
L22:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v909)))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v909)+4))
	if v927 < v928 {
		goto L155
	} else {
		goto L156
	}
L23:
	;
	v900 = F_pnstrdup(m, l3+v878, l4-v878)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L4
	} else {
		goto L153
	}
L24:
	;
	v134 = l4 - int32(1)
	v136 = v33
	v140 = l5
	v145 = v34
	goto L25
L25:
	;
	v160 = l3 + v140
	v162 = v136
	v171 = v145
	goto L27
L26:
	;
	v878 = v140
	v880 = v44
	v889 = v28
	v890 = v39
	goto L23
L27:
	;
	if v171 <= v140 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L26
L29:
	;
	goto L28
L30:
	;
	v679 = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v680 == v679 {
		v844 = v679
		goto L115
	} else {
		goto L116
	}
L31:
	;
	if v162 == int32(0) {
		goto L29
	} else {
		goto L114
	}
L32:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v187 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v171 == v134 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v193 = int32(8)
	goto L36
L35:
	;
	v193 = int32(4)
	goto L36
L36:
	;
	if v171 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v195 = v193
	goto L39
L38:
	;
	v195 = int32(2)
	goto L39
L39:
	;
	v199 = l4 - v171
	v200 = l3 + v171
	v209 = v187
	goto L40
L40:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v162 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+8)))
	if v370 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L43:
	;
	v349 = v256 - v200 + v258
	v352 = v239
	goto L42
L44:
	;
	if v226 == int32(0) {
		goto L29
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v226 == int32(0) {
		goto L30
	} else {
		goto L55
	}
L47:
	;
	v233 = v226
	v239 = v209
	goto L48
L48:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if v256 < v199 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L31
L50:
	;
	v258 = F_strstr(m, v200, v233)
	mBase = m.M
	if v258 != 0 {
		goto L43
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v239)+12))
	if v260 != 0 {
		v233 = v260
		v239 = v239 + int32(12)
		goto L48
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	goto L49
L55:
	;
	v267 = v226
	v273 = v209
	goto L56
L56:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v290 < v199 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L31
L58:
	;
	if v290 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	if v339 != 0 {
		v267 = v339
		v273 = v273 + int32(12)
		goto L56
	} else {
		goto L75
	}
L61:
	;
	if v336 == int32(0) {
		v349 = v290
		v352 = v273
		goto L42
	} else {
		goto L74
	}
L62:
	;
	v336 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v297 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v298 = v267
	v299 = v200
	v300 = v290
	v301 = v297
	goto L69
L66:
	;
	v324 = v200
	v328 = int32(0)
	goto L67
L67:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	v336 = v328 - v329
	goto L61
L68:
	;
	v324 = v319
	v328 = v321
	goto L67
L69:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if base.B2i32(v301 != v303)|base.B2i32(v303 == int32(0)) != 0 {
		v319 = v299
		v321 = v301
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v319 = v313
	v321 = int32(0)
	goto L68
L71:
	;
	v309 = v300 - int32(1)
	if v309 == int32(0) {
		v319 = v299
		v321 = v301
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v312 = int32(1)
	v313 = v299 + v312
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+1)))
	if v314 != 0 {
		v298 = v298 + v312
		v299 = v313
		v300 = v309
		v301 = v314
		goto L69
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	goto L60
L75:
	;
	goto L57
L76:
	;
	if v349 < int32(0) {
		goto L31
	} else {
		goto L79
	}
L77:
	;
	v375 = int32(0)
	goto L78
L78:
	;
	v376 = v375 + v171
	v379 = v39 + v376 - int32(1)
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v380 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v375 = v349
	goto L78
L80:
	;
	v209 = v352 + int32(12)
	goto L40
L81:
	;
	v383 = v375 + (v171 - v140)
	if base.B2i32(int32(255) < v383)|base.B2i32(v383+(v171-int32(1)) <= l6) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v389 = int32(0)
	if base.B2i32(v383 == v389)|base.B2i32(v383 <= v389) == v389 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	base.MemoryCopy(m, v28, v160, v383)
	goto L85
L84:
	;
	goto L85
L85:
	;
	v398 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v383+v28))) = uint8(v398)
	v400 = F_NormalizeSubWord(m, l0, v28, v195)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v400 == int32(0) {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v406 = F_palloc(m, int32(16))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+12)) = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v406)+4)) = v410
	v414 = F_palloc(m, v410<<(uint(int32(2))%32))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+8)) = v414
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v406))) = v417
	if int32(0) < v417 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v426 = int32(0)
	goto L93
L91:
	;
	goto L92
L92:
	;
	v483 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v379))) = uint8(v483)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	if v485 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v447 = v426 << (uint(int32(2)) % 32)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v450+v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v447+v448))) = v452
	v455 = v426 + int32(1)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v455 < v456 {
		v426 = v455
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L92
L95:
	;
	goto L94
L96:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v492 = v486
	v498 = v400
	v500 = v485
	goto L99
L97:
	;
	goto L98
L98:
	;
	F_pfree(m, v400)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L107
	}
L99:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v492 < v512 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L98
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v526+v525<<(uint(int32(2))%32)))) = v500
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v533 = v531 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v406))) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v535 != 0 {
		v492 = v533
		v498 = v498 + int32(4)
		v500 = v535
		goto L99
	} else {
		goto L106
	}
L102:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v525 = v492
	v526 = v514
	goto L101
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+4)) = v512 << (uint(int32(1)) % 32)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v521 = F_repalloc(m, v518, v512<<(uint(int32(3))%32))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406)+8)) = v521
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v525 = v524
	v526 = v521
	goto L101
L106:
	;
	goto L100
L107:
	;
	v570 = v44
	goto L108
L108:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	if v590 != 0 {
		v570 = v590
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v592 = F_SplitToVariants(m, l0, int32(0), v406, l3, l4, v376, v376)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570)+12)) = v592
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	F_pfree(m, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_pfree(m, v406)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	goto L80
L114:
	;
	goto L30
L115:
	;
	v846 = v171 + int32(1)
	if v846 < l4 {
		v162 = v844
		v171 = v846
		goto L27
	} else {
		goto L152
	}
L116:
	;
	v684 = v162 + int32(4)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v171))))
	v692 = v684 + v680<<(uint(int32(3))%32)
	v698 = v684
	goto L118
L117:
	;
	if v171 == v134 {
		goto L129
	} else {
		goto L130
	}
L118:
	;
	v720 = v698 + (v692-v698)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	v723 = v721 & int32(255)
	if v723 == v689 {
		goto L117
	} else {
		goto L120
	}
L119:
	;
	v844 = int32(0)
	goto L115
L120:
	;
	v727 = base.B2i32(base.Ui32(v723) < base.Ui32(v689))
	if base.Ui32(v723) < base.Ui32(v689) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v728 = v720 + int32(8)
	goto L123
L122:
	;
	v728 = v698
	goto L123
L123:
	;
	if base.Ui32(v723) < base.Ui32(v689) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v729 = v692
	goto L126
L125:
	;
	v729 = v720
	goto L126
L126:
	;
	if base.Ui32(v728) < base.Ui32(v729) {
		v692 = v729
		v698 = v728
		goto L118
	} else {
		goto L127
	}
L127:
	;
	goto L119
L128:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v720)+4))
	v844 = v818
	goto L115
L129:
	;
	v741 = int32(8)
	goto L131
L130:
	;
	v741 = int32(4)
	goto L131
L131:
	;
	if v140 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v743 = v741
	goto L134
L133:
	;
	v743 = int32(2)
	goto L134
L134:
	;
	if base.B2i32(v721&int32(256) == int32(0))|base.B2i32(int32(base.Ui32(v721)>>(uint(int32(9))%32))&v743 == int32(0))|base.B2i32(v171 <= l6) != 0 {
		goto L128
	} else {
		goto L135
	}
L135:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v39))))
	if v751 == int32(0) {
		goto L128
	} else {
		goto L136
	}
L136:
	;
	v755 = v171 + int32(1)
	if v755 == l4 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v758 = F_pnstrdup(m, v160, l4-v140)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v768 = v44
	goto L141
L140:
	;
	v904 = v758
	v909 = v44
	v918 = v28
	v919 = v39
	goto L22
L141:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v768)+12))
	if v785 != 0 {
		v768 = v785
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v786 = F_SplitToVariants(m, l0, v162, v44, l3, l4, v140, v171)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v768)+12)) = v786
	v790 = F_pnstrdup(m, v160, v755-v140)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v792 < v793 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806+v807<<(uint(int32(2))%32)))) = v790
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v812 + int32(1)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v755 < l4 {
		v136 = v816
		v140 = v755
		v145 = v755
		goto L25
	} else {
		goto L151
	}
L147:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v806 = v795
	v807 = v792
	goto L146
L148:
	;
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v793 << (uint(int32(1)) % 32)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v802 = F_repalloc(m, v799, v793<<(uint(int32(3))%32))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v802
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v806 = v802
	v807 = v805
	goto L146
L151:
	;
	v878 = v755
	v880 = v44
	v889 = v28
	v890 = v39
	goto L23
L152:
	;
	goto L29
L153:
	;
	v904 = v900
	v909 = v880
	v918 = v889
	v919 = v890
	goto L22
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v941+v942<<(uint(int32(2))%32)))) = v904
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v909)))
	*(*int32)(unsafe.Add(mBase, uint32(v909))) = v947 + int32(1)
	F_pfree(m, v919)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L4
	} else {
		goto L159
	}
L155:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v909)+8))
	v941 = v930
	v942 = v927
	goto L154
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v909)+4)) = v928 << (uint(int32(1)) % 32)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v909)+8))
	v937 = F_repalloc(m, v934, v928<<(uint(int32(3))%32))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v909)+8)) = v937
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v909)))
	v941 = v937
	v942 = v940
	goto L154
L159:
	;
	m.G0 = v918 + int32(256)
	return v909
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
				F_errmsg_internal(m, int32(_a_F_add_cast_to_0), v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_add_cast_to_1), int32(_a_F_add_cast_to_2), int32(_a_F_add_cast_to_3))
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
						F_appendStringInfo(m, l0, int32(_a_F_add_cast_to_4), v8+int32(16))
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = l1 - int32(8)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v8 < v5 {
		v12 = F_repalloc(m, v7, v5+int32(29))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 == int32(0) {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v5 + int32(20)
				v27 = v12 + int32(8)
				if v5 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					base.MemoryCopy(m, v27, v28, v5)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v27-int32(4)))) = v5
				return v27
			}
		}
	} else {
		v27 = l1
		if v5 != 0 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			base.MemoryCopy(m, v27, v28, v5)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v27-int32(4)))) = v5
		return v27
	}
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
	var v59 int32
	_ = v59
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
	v67 = int32(_a_F_to_ascii_encname_0)
	v68 = int32(_a_F_to_ascii_encname_1)
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
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v59 != 0 {
		v35 = v35 + int32(1)
		v36 = v59
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
	F_errmsg(m, int32(_a_F_to_ascii_encname_2), v7)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_to_ascii_encname_3), int32(128), int32(_a_F_to_ascii_encname_4))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14004(m, l0, int64(4), int64(16), int32(15))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_to_regcollation(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14005(m, l0, int32(1482))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_to_regnamespace(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14005(m, l0, int32(1484))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_to_regprocedure(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14005(m, l0, int32(1237))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
