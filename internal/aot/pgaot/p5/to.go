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
	var v8 int32
	_ = v8
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v513 int32
	_ = v513
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
	var v542 int32
	_ = v542
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	v8 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(256)
	m.G0 = v28
	if l1 == v8 {
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
		v885 = l5
		v887 = v44
		v895 = v28
		v896 = v39
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
	v51 = F_palloc_mul(m, int32(4), v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v97 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v97
	v101 = F_palloc_mul(m, int32(4), v97)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v54
	if v54 <= int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v70 = v8
	goto L17
L17:
	;
	v84 = v70 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84+v85)))
	v88 = F_pstrdup(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L11
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v90+v84))) = v88
	v94 = v70 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v94 < v95 {
		v70 = v94
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
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v101
	goto L11
L22:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v916)))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	if v934 < v935 {
		goto L158
	} else {
		goto L159
	}
L23:
	;
	v907 = F_pnstrdup(m, l3+v885, l4-v885)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L156
	}
L24:
	;
	v133 = l4 - int32(1)
	v135 = v33
	v139 = l5
	v143 = v34
	goto L25
L25:
	;
	v159 = l3 + v139
	v161 = v135
	v169 = v143
	goto L27
L26:
	;
	v885 = v139
	v887 = v44
	v895 = v28
	v896 = v39
	goto L23
L27:
	;
	if v169 <= v139 {
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
	v686 = int32(0)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v687 == v686 {
		v851 = v686
		goto L118
	} else {
		goto L119
	}
L31:
	;
	if v161 == int32(0) {
		goto L29
	} else {
		goto L117
	}
L32:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v186 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v169 == v133 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v192 = int32(8)
	goto L36
L35:
	;
	v192 = int32(4)
	goto L36
L36:
	;
	if v169 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v194 = v192
	goto L39
L38:
	;
	v194 = int32(2)
	goto L39
L39:
	;
	v198 = l4 - v169
	v199 = l3 + v169
	v205 = v186
	goto L40
L40:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v161 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+8)))
	if v369 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L43:
	;
	v348 = v235
	v351 = v255 - v199 + v257
	goto L42
L44:
	;
	if v225 == int32(0) {
		goto L29
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v225 == int32(0) {
		goto L30
	} else {
		goto L55
	}
L47:
	;
	v232 = v225
	v235 = v205
	goto L48
L48:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v255 < v198 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L31
L50:
	;
	v257 = F_strstr(m, v199, v232)
	mBase = m.M
	if v257 != 0 {
		goto L43
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	if v259 != 0 {
		v232 = v259
		v235 = v235 + int32(12)
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
	v266 = v225
	v269 = v205
	goto L56
L56:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v289 < v198 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L31
L58:
	;
	if v289 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	if v338 != 0 {
		v266 = v338
		v269 = v269 + int32(12)
		goto L56
	} else {
		goto L75
	}
L61:
	;
	if v335 == int32(0) {
		v348 = v269
		v351 = v289
		goto L42
	} else {
		goto L74
	}
L62:
	;
	v335 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v296 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v297 = v266
	v298 = v199
	v299 = v289
	v300 = v296
	goto L69
L66:
	;
	v323 = v199
	v327 = int32(0)
	goto L67
L67:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v335 = v327 - v328
	goto L61
L68:
	;
	v323 = v318
	v327 = v320
	goto L67
L69:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	if base.B2i32(v300 != v302)|base.B2i32(v302 == int32(0)) != 0 {
		v318 = v298
		v320 = v300
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v318 = v312
	v320 = int32(0)
	goto L68
L71:
	;
	v308 = v299 - int32(1)
	if v308 == int32(0) {
		v318 = v298
		v320 = v300
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v311 = int32(1)
	v312 = v298 + v311
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v313 != 0 {
		v297 = v297 + v311
		v298 = v312
		v299 = v308
		v300 = v313
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
	if v351 < int32(0) {
		goto L31
	} else {
		goto L79
	}
L77:
	;
	v374 = int32(0)
	goto L78
L78:
	;
	v375 = v374 + v169
	v378 = v39 + v375 - int32(1)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v379 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v374 = v351
	goto L78
L80:
	;
	v205 = v348 + int32(12)
	goto L40
L81:
	;
	v382 = v374 + (v169 - v139)
	if base.B2i32(int32(255) < v382)|base.B2i32(v382+(v169-int32(1)) <= l6) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v388 = int32(0)
	if base.B2i32(v382 == v388)|base.B2i32(v382 <= v388) == v388 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	base.MemoryCopy(m, v28, v159, v382)
	goto L85
L84:
	;
	goto L85
L85:
	;
	v397 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v382+v28))) = uint8(v397)
	v399 = F_NormalizeSubWord(m, l0, v28, v194)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v399 == int32(0) {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v404 = F_palloc(m, int32(16))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+12)) = int32(0)
	if v44 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v490 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v378))) = uint8(v490)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	if v492 != 0 {
		goto L99
	} else {
		goto L100
	}
L90:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v404)+4)) = v408
	v411 = F_palloc_mul(m, int32(4), v408)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v456 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v404)+4)) = v456
	v460 = F_palloc_mul(m, int32(4), v456)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L98
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+8)) = v411
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = v414
	v416 = int32(0)
	if v414 <= v416 {
		goto L89
	} else {
		goto L94
	}
L94:
	;
	v427 = v416
	goto L95
L95:
	;
	v445 = v427 << (uint(int32(2)) % 32)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v448+v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v445+v446))) = v450
	v453 = v427 + int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v453 < v454 {
		v427 = v453
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L89
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v404)+8)) = v460
	goto L89
L99:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v496 = v399
	v502 = v493
	v513 = v492
	goto L102
L100:
	;
	goto L101
L101:
	;
	F_pfree(m, v399)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L4
	} else {
		goto L110
	}
L102:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	if v502 < v519 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533+v532<<(uint(int32(2))%32)))) = v513
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v540 = v538 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = v540
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	if v542 != 0 {
		v496 = v496 + int32(4)
		v502 = v540
		v513 = v542
		goto L102
	} else {
		goto L109
	}
L105:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v532 = v502
	v533 = v521
	goto L104
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+4)) = v519 << (uint(int32(1)) % 32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	v528 = F_repalloc(m, v525, v519<<(uint(int32(3))%32))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v404)+8)) = v528
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v532 = v531
	v533 = v528
	goto L104
L109:
	;
	goto L103
L110:
	;
	v580 = v44
	goto L111
L111:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v580)+12))
	if v597 != 0 {
		v580 = v597
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v599 = F_SplitToVariants(m, l0, int32(0), v404, l3, l4, v375, v375)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580)+12)) = v599
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v404)+8))
	F_pfree(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	F_pfree(m, v404)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	goto L80
L117:
	;
	goto L30
L118:
	;
	v853 = v169 + int32(1)
	if v853 < l4 {
		v161 = v851
		v169 = v853
		goto L27
	} else {
		goto L155
	}
L119:
	;
	v691 = v161 + int32(4)
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v169))))
	v699 = v691 + v687<<(uint(int32(3))%32)
	v702 = v691
	goto L121
L120:
	;
	if v169 == v133 {
		goto L132
	} else {
		goto L133
	}
L121:
	;
	v727 = v702 + (v699-v702)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v730 = v728 & int32(255)
	if v730 == v696 {
		goto L120
	} else {
		goto L123
	}
L122:
	;
	v851 = int32(0)
	goto L118
L123:
	;
	v734 = base.B2i32(base.Ui32(v730) < base.Ui32(v696))
	if base.Ui32(v730) < base.Ui32(v696) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v735 = v727 + int32(8)
	goto L126
L125:
	;
	v735 = v702
	goto L126
L126:
	;
	if base.Ui32(v730) < base.Ui32(v696) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v736 = v699
	goto L129
L128:
	;
	v736 = v727
	goto L129
L129:
	;
	if base.Ui32(v735) < base.Ui32(v736) {
		v699 = v736
		v702 = v735
		goto L121
	} else {
		goto L130
	}
L130:
	;
	goto L122
L131:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	v851 = v825
	goto L118
L132:
	;
	v748 = int32(8)
	goto L134
L133:
	;
	v748 = int32(4)
	goto L134
L134:
	;
	if v139 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v750 = v748
	goto L137
L136:
	;
	v750 = int32(2)
	goto L137
L137:
	;
	if base.B2i32(v728&int32(256) == int32(0))|base.B2i32(int32(base.Ui32(v728)>>(uint(int32(9))%32))&v750 == int32(0))|base.B2i32(v169 <= l6) != 0 {
		goto L131
	} else {
		goto L138
	}
L138:
	;
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v39))))
	if v758 == int32(0) {
		goto L131
	} else {
		goto L139
	}
L139:
	;
	v762 = v169 + int32(1)
	if v762 == l4 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v765 = F_pnstrdup(m, v159, l4-v139)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L4
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v769 = v44
	goto L144
L143:
	;
	v909 = v765
	v916 = v44
	v924 = v28
	v925 = v39
	goto L22
L144:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	if v792 != 0 {
		v769 = v792
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v793 = F_SplitToVariants(m, l0, v161, v44, l3, l4, v139, v169)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L4
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v769)+12)) = v793
	v797 = F_pnstrdup(m, v159, v762-v139)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v799 < v800 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v814+v813<<(uint(int32(2))%32)))) = v797
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v819 + int32(1)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v762 < l4 {
		v135 = v823
		v139 = v762
		v143 = v762
		goto L25
	} else {
		goto L154
	}
L150:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v813 = v799
	v814 = v802
	goto L149
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v800 << (uint(int32(1)) % 32)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v809 = F_repalloc(m, v806, v800<<(uint(int32(3))%32))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v809
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v813 = v812
	v814 = v809
	goto L149
L154:
	;
	v885 = v762
	v887 = v44
	v895 = v28
	v896 = v39
	goto L23
L155:
	;
	goto L29
L156:
	;
	v909 = v907
	v916 = v887
	v924 = v895
	v925 = v896
	goto L22
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v948+v949<<(uint(int32(2))%32)))) = v909
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v916)))
	*(*int32)(unsafe.Add(mBase, uint32(v916))) = v954 + int32(1)
	F_pfree(m, v925)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L162
	}
L158:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v916)+8))
	v948 = v937
	v949 = v934
	goto L157
L159:
	;
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v916)+4)) = v935 << (uint(int32(1)) % 32)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v916)+8))
	v944 = F_repalloc(m, v941, v935<<(uint(int32(3))%32))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v916)+8)) = v944
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v916)))
	v948 = v944
	v949 = v947
	goto L157
L162:
	;
	m.G0 = v924 + int32(256)
	return v916
}
func F_add_cast_to(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_errmsg_internal(m, int32(_a_F_add_cast_to_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_add_cast_to_1), int32(_a_F_add_cast_to_2), int32(_a_F_add_cast_to_3))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
			v33 = F_get_namespace_name_or_temp(m, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = F_quote_identifier(m, v33)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v39 = F_quote_identifier(m, v31+int32(4))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v35
						F_appendStringInfo(m, l0, int32(_a_F_add_cast_to_4), v8+int32(16))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = v7 - int32(8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v10 < v6 {
		v14 = F_repalloc(m, v9, v6+int32(29))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				return int32(-1)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v6 + int32(20)
				v26 = v14 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
				v28 = v26
				if v6 != 0 {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					base.MemoryCopy(m, v28, v29, v6)
				} else {
				}
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v31-int32(4)))) = v6
				return int32(0)
			}
		}
	} else {
		v28 = v7
		if v6 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			base.MemoryCopy(m, v28, v29, v6)
		} else {
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v31-int32(4)))) = v6
		return int32(0)
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
func F_to_ascii_encname(m *base.Module, l0 int32) int64 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(v125)
}
func F_to_hex64(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14388(m, l0, int64(4), int64(16), int32(15))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_to_regcollation(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14389(m, l0, int32(1694))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_to_regnamespace(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14389(m, l0, int32(1696))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_to_regprocedure(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14389(m, l0, int32(1365))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
