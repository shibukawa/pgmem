package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BIG5toCNS(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	if base.Ui32(l0) <= base.Ui32(int32(51519)) {
		goto L13
	} else {
		goto L14
	}
L1:
	;
	return v318 & int32(65535)
L2:
	;
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v313)
	v318 = int32(63)
	goto L1
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v307)
	v318 = v306 | int32(-32640)
	goto L1
L4:
	;
	v301 = int32(246)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v301)
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+2)))
	v318 = v303 | int32(-32640)
	goto L1
L5:
	;
	v300 = int32(2205736)
	goto L4
L6:
	;
	v300 = int32(2205732)
	goto L4
L7:
	;
	v300 = int32(2205728)
	goto L4
L8:
	;
	v300 = int32(2205724)
	goto L4
L9:
	;
	v300 = int32(2205720)
	goto L4
L10:
	;
	v300 = int32(2205716)
	goto L4
L11:
	;
	v170 = int32(46)
	v172 = int32(23)
	v173 = int32(0)
	goto L62
L12:
	;
	v306 = v157
	v307 = int32(149)
	goto L3
L13:
	;
	switch l0 - int32(51321) {
	case 0:
		v145 = int32(2205584)
		goto L16
	case 1, 3:
		goto L20
	case 2:
		goto L19
	case 4:
		goto L18
	default:
		goto L21
	}
L14:
	;
	goto L15
L15:
	;
	switch l0 - int32(63958) {
	case 0:
		v300 = int32(2205712)
		goto L4
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	case 5:
		goto L6
	case 6:
		goto L5
	default:
		goto L58
	}
L16:
	;
	v146 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v146)
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+2)))
	v318 = v148 | int32(-32640)
	goto L1
L17:
	;
	v145 = int32(2205596)
	goto L16
L18:
	;
	v145 = int32(2205592)
	goto L16
L19:
	;
	v145 = int32(2205588)
	goto L16
L20:
	;
	v21 = int32(23)
	v23 = int32(11)
	v24 = int32(0)
	goto L25
L21:
	;
	if l0 == int32(51362) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v141 != 0 {
		v157 = v141
		goto L12
	} else {
		goto L57
	}
L24:
	;
	v141 = v139 & int32(65535)
	goto L23
L25:
	;
	v29 = v23 << (uint(int32(2)) % 32)
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1478]))))
	v32 = base.B2i32(base.Ui32(l0) < base.Ui32(v31))
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v139 = int32(0)
	goto L24
L27:
	;
	goto L26
L28:
	;
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L50
	} else {
		goto L51
	}
L29:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1479]))))
	if base.Ui32(v33) <= base.Ui32(l0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[1480]))))
	if v35 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v40 = l0 - v31&int32(65280)
	if base.Ui32(int32(41280)) <= base.Ui32(l0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v43 = int32(255)
	v44 = l0 & v43
	v46 = v31 & v43
	v56 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v46))
	if base.Ui32(int32(160)) < base.Ui32(v46) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v83 = int32(255)
	v94 = v35 & v83
	if base.Ui32(int32(160)) < base.Ui32(v94) {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v57 = int32(0)
	goto L37
L36:
	;
	v57 = int32(-34)
	goto L37
L37:
	;
	if base.Ui32(int32(160)) < base.Ui32(v46) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v60 = int32(34)
	goto L40
L39:
	;
	v60 = int32(0)
	goto L40
L40:
	;
	if base.Ui32(int32(160)) < base.Ui32(v44) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v63 = v57
	goto L43
L42:
	;
	v63 = v60
	goto L43
L43:
	;
	v68 = int32(33)
	v69 = v44 - v46 + v40>>(uint(int32(8))%32)*int32(157) + v63 + v35&int32(255) - v68
	v70 = int32(94)
	v71 = base.I32_div_s(v69, v70)
	v139 = v69 - v71*v70 + v35&int32(65280) + v71<<(uint(int32(8))%32) + v68
	goto L24
L44:
	;
	v100 = int32(65438)
	goto L46
L45:
	;
	v100 = int32(65472)
	goto L46
L46:
	;
	v101 = l0&v83 - v31&v83 + int32(base.Ui32(v40)>>(uint(int32(8))%32))*int32(94) + v94 + v100
	v103 = int32(157)
	v104 = base.I32_div_s(base.I32_extend16_s(v101), v103)
	v107 = v101 - v104*v103
	if int32(62) < base.I32_extend16_s(v107) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v119 = int32(98)
	goto L49
L48:
	;
	v119 = int32(64)
	goto L49
L49:
	;
	v139 = v107 + v35&int32(65280) + v104<<(uint(int32(8))%32) + v119
	goto L24
L50:
	;
	v123 = v24
	goto L52
L51:
	;
	v123 = v23 + int32(1)
	goto L52
L52:
	;
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v126 = v23 - int32(1)
	goto L55
L54:
	;
	v126 = v21
	goto L55
L55:
	;
	if v123 <= v126 {
		v21 = v126
		v23 = (v123 + v126) >> (uint(int32(1)) % 32)
		v24 = v123
		goto L25
	} else {
		goto L56
	}
L56:
	;
	goto L27
L57:
	;
	goto L2
L58:
	;
	if l0 != int32(51530) {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v157 = int32(17474)
	goto L12
L60:
	;
	if v290 == int32(0) {
		goto L2
	} else {
		goto L94
	}
L61:
	;
	v290 = v288 & int32(65535)
	goto L60
L62:
	;
	v178 = v172 << (uint(int32(2)) % 32)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+uint32(_consts[1481]))))
	v181 = base.B2i32(base.Ui32(l0) < base.Ui32(v180))
	if base.Ui32(l0) < base.Ui32(v180) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v288 = int32(0)
	goto L61
L64:
	;
	goto L63
L65:
	;
	if base.Ui32(l0) < base.Ui32(v180) {
		goto L87
	} else {
		goto L88
	}
L66:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+uint32(_consts[1482]))))
	if base.Ui32(v182) <= base.Ui32(l0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+uint32(_consts[1483]))))
	if v184 == int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v189 = l0 - v180&int32(65280)
	if base.Ui32(int32(41280)) <= base.Ui32(l0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v192 = int32(255)
	v193 = l0 & v192
	v195 = v180 & v192
	v205 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v195))
	if base.Ui32(int32(160)) < base.Ui32(v195) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v232 = int32(255)
	v243 = v184 & v232
	if base.Ui32(int32(160)) < base.Ui32(v243) {
		goto L81
	} else {
		goto L82
	}
L72:
	;
	v206 = int32(0)
	goto L74
L73:
	;
	v206 = int32(-34)
	goto L74
L74:
	;
	if base.Ui32(int32(160)) < base.Ui32(v195) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v209 = int32(34)
	goto L77
L76:
	;
	v209 = int32(0)
	goto L77
L77:
	;
	if base.Ui32(int32(160)) < base.Ui32(v193) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v212 = v206
	goto L80
L79:
	;
	v212 = v209
	goto L80
L80:
	;
	v217 = int32(33)
	v218 = v193 - v195 + v189>>(uint(int32(8))%32)*int32(157) + v212 + v184&int32(255) - v217
	v219 = int32(94)
	v220 = base.I32_div_s(v218, v219)
	v288 = v218 - v220*v219 + v184&int32(65280) + v220<<(uint(int32(8))%32) + v217
	goto L61
L81:
	;
	v249 = int32(65438)
	goto L83
L82:
	;
	v249 = int32(65472)
	goto L83
L83:
	;
	v250 = l0&v232 - v180&v232 + int32(base.Ui32(v189)>>(uint(int32(8))%32))*int32(94) + v243 + v249
	v252 = int32(157)
	v253 = base.I32_div_s(base.I32_extend16_s(v250), v252)
	v256 = v250 - v253*v252
	if int32(62) < base.I32_extend16_s(v256) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v268 = int32(98)
	goto L86
L85:
	;
	v268 = int32(64)
	goto L86
L86:
	;
	v288 = v256 + v184&int32(65280) + v253<<(uint(int32(8))%32) + v268
	goto L61
L87:
	;
	v272 = v173
	goto L89
L88:
	;
	v272 = v172 + int32(1)
	goto L89
L89:
	;
	if base.Ui32(l0) < base.Ui32(v180) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v275 = v172 - int32(1)
	goto L92
L91:
	;
	v275 = v170
	goto L92
L92:
	;
	if v272 <= v275 {
		v170 = v275
		v172 = (v272 + v275) >> (uint(int32(1)) % 32)
		v173 = v272
		goto L62
	} else {
		goto L93
	}
L93:
	;
	goto L64
L94:
	;
	v306 = v290
	v307 = int32(150)
	goto L3
}
func F_BackgroundWriterMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int64
	_ = v438
	var v439 int64
	_ = v439
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v473 int32
	_ = v473
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
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int64
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v736 float32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v746 float32
	_ = v746
	var v751 float32
	_ = v751
	var v753 float32
	_ = v753
	var v754 float32
	_ = v754
	var v755 int32
	_ = v755
	var v757 float64
	_ = v757
	var v759 float32
	_ = v759
	var v765 float32
	_ = v765
	var v767 float64
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v776 float32
	_ = v776
	var v780 float32
	_ = v780
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v793 float32
	_ = v793
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int64
	_ = v864
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v901 int32
	_ = v901
	var v903 int64
	_ = v903
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v913 float32
	_ = v913
	var v924 int32
	_ = v924
	var v950 int32
	_ = v950
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int64
	_ = v992
	var v993 int32
	_ = v993
	var v994 int64
	_ = v994
	var v997 int64
	_ = v997
	var v998 int32
	_ = v998
	var v999 int64
	_ = v999
	var v1002 int64
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int64
	_ = v1004
	var v1012 int64
	_ = v1012
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int64
	_ = v1102
	var v1103 int64
	_ = v1103
	var v1111 int64
	_ = v1111
	var v1113 int64
	_ = v1113
	var v1121 int64
	_ = v1121
	var v1122 int64
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1128 int64
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1192 int32
	_ = v1192
	var v1200 int32
	_ = v1200
	var v1201 int64
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v28 = int32(-1)
	v30 = v3
	v31 = v3
	v32 = v3
	v34 = v3
	v42 = v25
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v28 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v1200 = int32(m.ExcTag)
	v1201 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1200 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L6:
	;
	v53 = v42 - int32(160)
	m.G0 = v53
	v56 = v53 - int32(5136)
	m.G0 = v56
	*(*int32)(unsafe.Add(mBase, _consts[264])) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		v1192 = v56
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v473 = v30
	v474 = v31
	v475 = v32
	v476 = v42
	v477 = v34
	goto L8
L8:
	;
	if v477 != 0 {
		goto L121
	} else {
		goto L122
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v70 = int32(913)
	v72 = m.G0
	v74 = v72 - int32(144)
	m.G0 = v74
	switch int32(915) {
	case 0, 2:
		v84 = v70
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v115 = int32(-2)
	v117 = m.G0
	v119 = v117 - int32(144)
	m.G0 = v119
	switch int32(0) {
	case 0, 2:
		v129 = v115
		goto L24
	default:
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v84
	F_sigemptyset(m, v74+int32(8))
	mBase = m.M
	goto L14
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[321])) = v70
	v84 = int32(4729)
	goto L11
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+136)) = int32(268435456)
	v96 = v74 + int32(4)
	goto L18
L16:
	;
	m.G0 = v74 + int32(144)
	goto L10
L18:
	;
	goto L19
L19:
	;
	if v96 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v107 = F___memcpy(m, int32(4635612), v96, int32(140))
	mBase = m.M
	goto L22
L21:
	;
	goto L22
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v160 = int32(915)
	v162 = m.G0
	v164 = v162 - int32(144)
	m.G0 = v164
	switch int32(917) {
	case 0, 2:
		v174 = v160
		goto L37
	default:
		goto L38
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v129
	F_sigemptyset(m, v119+int32(8))
	mBase = m.M
	goto L27
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v115
	v129 = int32(4729)
	goto L24
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+136)) = int32(268435456)
	v141 = v119 + int32(4)
	goto L31
L29:
	;
	m.G0 = v119 + int32(144)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if v141 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v152 = F___memcpy(m, int32(4635752), v141, int32(140))
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L29
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v205 = int32(-2)
	v207 = m.G0
	v209 = v207 - int32(144)
	m.G0 = v209
	switch int32(0) {
	case 0, 2:
		v219 = v205
		goto L50
	default:
		goto L51
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v174
	F_sigemptyset(m, v164+int32(8))
	mBase = m.M
	goto L40
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[323])) = v160
	v174 = int32(4729)
	goto L37
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+136)) = int32(268435456)
	v186 = v164 + int32(4)
	goto L44
L42:
	;
	m.G0 = v164 + int32(144)
	goto L36
L44:
	;
	goto L45
L45:
	;
	if v186 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v197 = F___memcpy(m, int32(4637572), v186, int32(140))
	mBase = m.M
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L42
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v250 = int32(-2)
	v252 = m.G0
	v254 = v252 - int32(144)
	m.G0 = v254
	switch int32(0) {
	case 0, 2:
		v264 = v250
		goto L63
	default:
		goto L64
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+4)) = v219
	F_sigemptyset(m, v209+int32(8))
	mBase = m.M
	goto L53
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[724])) = v205
	v219 = int32(4729)
	goto L50
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+136)) = int32(268435456)
	v231 = v209 + int32(4)
	goto L57
L55:
	;
	m.G0 = v209 + int32(144)
	goto L49
L57:
	;
	goto L58
L58:
	;
	if v231 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v242 = F___memcpy(m, int32(4637432), v231, int32(140))
	mBase = m.M
	goto L61
L60:
	;
	goto L61
L61:
	;
	goto L55
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v295 = int32(916)
	v297 = m.G0
	v299 = v297 - int32(144)
	m.G0 = v299
	switch int32(918) {
	case 0, 2:
		v309 = v295
		goto L76
	default:
		goto L77
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = v264
	F_sigemptyset(m, v254+int32(8))
	mBase = m.M
	goto L66
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v250
	v264 = int32(4729)
	goto L63
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+136)) = int32(268435456)
	v276 = v254 + int32(4)
	goto L70
L68:
	;
	m.G0 = v254 + int32(144)
	goto L62
L70:
	;
	goto L71
L71:
	;
	if v276 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v287 = F___memcpy(m, int32(4637292), v276, int32(140))
	mBase = m.M
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L68
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v340 = int32(-2)
	v342 = m.G0
	v344 = v342 - int32(144)
	m.G0 = v344
	switch int32(0) {
	case 0, 2:
		v354 = v340
		goto L89
	default:
		goto L90
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+4)) = v309
	F_sigemptyset(m, v299+int32(8))
	mBase = m.M
	goto L79
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[647])) = v295
	v309 = int32(4729)
	goto L76
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v299)+136)) = int32(268435456)
	v321 = v299 + int32(4)
	goto L83
L81:
	;
	m.G0 = v299 + int32(144)
	goto L75
L83:
	;
	goto L84
L84:
	;
	if v321 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v332 = F___memcpy(m, int32(4636872), v321, int32(140))
	mBase = m.M
	goto L87
L86:
	;
	goto L87
L87:
	;
	goto L81
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v385 = int32(0)
	v387 = m.G0
	v389 = v387 - int32(144)
	m.G0 = v389
	switch int32(2) {
	case 0, 2:
		v399 = v385
		goto L102
	default:
		goto L103
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v354
	F_sigemptyset(m, v344+int32(8))
	mBase = m.M
	goto L92
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v340
	v354 = int32(4729)
	goto L89
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344)+136)) = int32(268435456)
	v366 = v344 + int32(4)
	goto L96
L94:
	;
	m.G0 = v344 + int32(144)
	goto L88
L96:
	;
	goto L97
L97:
	;
	if v366 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v377 = F___memcpy(m, int32(4637152), v366, int32(140))
	mBase = m.M
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L94
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v433 = m.G0
	v434 = int32(16)
	v435 = v433 - v434
	m.G0 = v435
	F___gettimeofday(m, v435)
	mBase = m.M
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v435)))
	v439 = int64(*(*int32)(unsafe.Add(mBase, uint32(v435)+8)))
	m.G0 = v435 + v434
	goto L114
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v399
	F_sigemptyset(m, v389+int32(8))
	mBase = m.M
	goto L104
L103:
	;
	*(*int32)(unsafe.Add(mBase, _consts[650])) = v385
	v399 = int32(4729)
	goto L102
L104:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389)+136)) = int32(268435457)
	v411 = v389 + int32(4)
	goto L109
L107:
	;
	m.G0 = v389 + int32(144)
	goto L101
L109:
	;
	goto L110
L110:
	;
	if v411 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v422 = F___memcpy(m, int32(4637852), v411, int32(140))
	mBase = m.M
	goto L113
L112:
	;
	goto L113
L113:
	;
	goto L107
L114:
	;
	*(*int64)(unsafe.Add(mBase, _consts[725])) = v439 + v438*int64(1000000) - int64(946684800000000)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	v454 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v459 = F_AllocSetContextCreateInternal(m, v454, int32(212839), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		v1192 = v56
		goto L5
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(4386348)
	goto L116
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v25
	goto L120
L118:
	;
	v473 = v56
	v474 = v53
	v475 = v459
	v476 = v56
	v477 = int32(0)
	goto L8
L120:
	;
	goto L118
L121:
	;
	v478 = int32(4465212)
	v480 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v480 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[385])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_EmitErrorReport(m)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_sigprocmask(m, int32(4377784), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L137
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_LWLockReleaseAll(m)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_UnlockBuffers(m)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_smgrdestroyall(m)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_FlushErrorState(m)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_MemoryContextReset(m, v475)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = int32(4386348)
	goto L135
L135:
	;
	v557 = int32(4465212)
	v559 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v559 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_pg_usleep(m, int32(1000000))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = int32(0)
	goto L123
L137:
	;
	v589 = int32(0)
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v609 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v609))) = int32(0)
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_ProcessMainLoopInterrupts(m)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v621 = m.G0
	v623 = v621 - int32(16)
	m.G0 = v623
	v626 = v623 + int32(12)
	v628 = v623 + int32(8)
	v630 = *(*int32)(unsafe.Add(mBase, _consts[726]))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	*(*int32)(unsafe.Add(mBase, uint32(v630))) = int32(1)
	if v631 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v635 = *(*int32)(unsafe.Add(mBase, _consts[726]))
	F_s_lock(m, v635, int32(484203), int32(399), int32(81888))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _consts[726]))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	v645 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	if v626 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v642)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v646
	v649 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v650 = base.I32_div_u_s(v643, v649)
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v646 + v650
	goto L148
L147:
	;
	goto L148
L148:
	;
	if v628 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v642)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v642)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v654
	v659 = *(*int32)(unsafe.Add(mBase, _consts[726]))
	v661 = v659
	goto L151
L150:
	;
	v661 = v642
	goto L151
L151:
	;
	v662 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v661))) = v662
	v664 = base.I32_rem_u_s(v643, v645)
	v665 = int32(4450176)
	v667 = *(*int64)(unsafe.Add(mBase, _consts[727]))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	*(*int64)(unsafe.Add(mBase, _consts[727])) = v667 + base.I64_extend_i32_u(v668)
	v673 = *(*int32)(unsafe.Add(mBase, _consts[728]))
	if v673 <= v662 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	m.G0 = v623 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v958 = *(*int32)(unsafe.Add(mBase, _consts[729]))
	v960 = *(*int32)(unsafe.Add(mBase, _consts[730]))
	v962 = *(*int32)(unsafe.Add(mBase, _consts[731]))
	v964 = *(*int32)(unsafe.Add(mBase, _consts[727]))
	v966 = *(*int32)(unsafe.Add(mBase, _consts[732]))
	v968 = *(*int32)(unsafe.Add(mBase, _consts[733]))
	v970 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	v972 = *(*int32)(unsafe.Add(mBase, _consts[735]))
	if v958|(v960|(v962|(v964|(v966|(v968|(v970|v972)))))) != 0 {
		goto L215
	} else {
		goto L216
	}
L153:
	;
	v677 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[736])) = uint8(v677)
	v950 = int32(1)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, _consts[736])))
	if v681 != 0 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _consts[737])) = v724
	*(*int32)(unsafe.Add(mBase, _consts[738])) = v664
	v733 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[736])) = uint8(v733)
	v736 = *(*float32)(unsafe.Add(mBase, _consts[739]))
	v737 = int32(0)
	if v668 != 0 {
		goto L168
	} else {
		goto L169
	}
L157:
	;
	v723 = v719
	v724 = v720
	v725 = v664
	v726 = v722
	v727 = v719
	goto L156
L158:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _consts[738]))
	v686 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v689 = *(*int32)(unsafe.Add(mBase, _consts[737]))
	v692 = v664 - v683 + v686*(v687-v689)
	v694 = *(*int32)(unsafe.Add(mBase, _consts[740]))
	if int32(0) < v694-v687 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, _consts[741])) = v664
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	*(*int32)(unsafe.Add(mBase, _consts[740])) = v715
	v718 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v719 = v718
	v720 = v715
	v722 = int32(0)
	goto L157
L161:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _consts[741]))
	v723 = v686
	v724 = v687
	v725 = v699
	v726 = v692
	v727 = v664 - v699
	goto L156
L162:
	;
	goto L163
L163:
	;
	if v687 != v694 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, _consts[740])) = v687
	*(*int32)(unsafe.Add(mBase, _consts[741])) = v664
	v719 = v686
	v720 = v687
	v722 = v692
	goto L157
L165:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _consts[741]))
	if v703 < v664 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v723 = v686
	v724 = v687
	v725 = v703
	v726 = v692
	v727 = v686 + v664 - v703
	goto L156
L167:
	;
	v755 = int32(0)
	v757 = *(*float64)(unsafe.Add(mBase, _consts[742]))
	v759 = *(*float32)(unsafe.Add(mBase, _consts[743]))
	if base.F32_ge(v753, v759) != 0 {
		goto L175
	} else {
		goto L176
	}
L168:
	;
	v740 = base.B2i32(v737 < v726)
	goto L170
L169:
	;
	v740 = v737
	goto L170
L170:
	;
	if v740 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v753 = base.F32_convert_i32_u(v668)
	v754 = v736
	goto L167
L172:
	;
	goto L173
L173:
	;
	v746 = base.F32_convert_i32_u(v668)
	v751 = base.F32_add(v736, base.F32_mul(base.F32_sub(base.F32_div(base.F32_convert_i32_u(v726), v746), v736), float32(0.0625)))
	*(*float32)(unsafe.Add(mBase, _consts[739])) = v751
	v753 = v746
	v754 = v751
	goto L167
L174:
	;
	if v773 != 0 {
		goto L181
	} else {
		goto L182
	}
L175:
	;
	v765 = v753
	goto L177
L176:
	;
	v765 = base.F32_add(v759, base.F32_mul(base.F32_sub(v753, v759), float32(0.0625)))
	goto L177
L177:
	;
	v767 = base.F64_mul(v757, base.F64_promote_f32(v765))
	if base.F64_lt(base.F64_abs(v767), float64(2.147483648e+09)) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v771 = base.I32_trunc_f64_s(v767)
	v773 = v771
	goto L174
L179:
	;
	goto L180
L180:
	;
	v773 = int32(-2147483648)
	goto L174
L181:
	;
	v776 = v765
	goto L183
L182:
	;
	v776 = float32(0)
	goto L183
L183:
	;
	*(*float32)(unsafe.Add(mBase, _consts[743])) = v776
	v780 = base.F32_div(base.F32_convert_i32_s(v723-v727), v754)
	if base.F32_lt(base.F32_abs(v780), float32(2.1474836e+09)) != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v790 = *(*int32)(unsafe.Add(mBase, _consts[744]))
	v793 = base.F32_div(base.F32_convert_i32_s(v723), base.F32_div(float32(120000), base.F32_convert_i32_s(v790)))
	if base.F32_lt(base.F32_abs(v793), float32(2.1474836e+09)) != 0 {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	v784 = base.I32_trunc_f32_s(v780)
	v786 = v784
	goto L184
L186:
	;
	goto L187
L187:
	;
	v786 = int32(-2147483648)
	goto L184
L188:
	;
	if v727 <= int32(0) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v797 = base.I32_trunc_f32_s(v793)
	v799 = v797
	goto L188
L190:
	;
	goto L191
L191:
	;
	v799 = int32(-2147483648)
	goto L188
L192:
	;
	v901 = int32(4450160)
	v903 = *(*int64)(unsafe.Add(mBase, _consts[735]))
	*(*int64)(unsafe.Add(mBase, _consts[735])) = v903 + base.I64_extend_i32_s(v890)
	v907 = v727 - v886
	if v907 <= int32(0) {
		goto L212
	} else {
		goto L213
	}
L193:
	;
	v885 = v786
	v886 = v727
	v890 = v755
	goto L192
L194:
	;
	v802 = v799 + v786
	if v773 < v802 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v804 = v802
	goto L197
L196:
	;
	v804 = v773
	goto L197
L197:
	;
	if v804 <= v786 {
		goto L193
	} else {
		goto L198
	}
L198:
	;
	v811 = v727
	v812 = v786
	v814 = v725
	v817 = v755
	goto L199
L199:
	;
	v829 = F_SyncOneBuffer(m, v814, int32(1), v473)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L201
	}
L200:
	;
	v885 = v873
	v886 = v852
	v890 = v872
	goto L192
L201:
	;
	v831 = int32(4386376)
	v833 = *(*int32)(unsafe.Add(mBase, _consts[741]))
	v835 = v833 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[741])) = v835
	v838 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	if v838 <= v835 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v840 = int32(4386380)
	v842 = *(*int32)(unsafe.Add(mBase, _consts[740]))
	*(*int32)(unsafe.Add(mBase, _consts[740])) = v842 + int32(1)
	v847 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[741])) = v847
	v850 = v847
	goto L204
L203:
	;
	v850 = v835
	goto L204
L204:
	;
	v851 = int32(1)
	v852 = v811 - v851
	if v829&v851 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	if base.Ui32(v811) < base.Ui32(int32(2)) {
		v885 = v873
		v886 = v852
		v890 = v872
		goto L192
	} else {
		goto L210
	}
L206:
	;
	v855 = int32(1)
	v856 = v812 + v855
	v858 = v817 + v855
	v860 = *(*int32)(unsafe.Add(mBase, _consts[728]))
	if v858 < v860 {
		v872 = v858
		v873 = v856
		goto L205
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v872 = v817
	v873 = int32(base.Ui32(v829)>>(uint(int32(1))%32)) + v812
	goto L205
L209:
	;
	v862 = int32(4450168)
	v864 = *(*int64)(unsafe.Add(mBase, _consts[733]))
	*(*int64)(unsafe.Add(mBase, _consts[733])) = v864 + int64(1)
	v885 = v856
	v886 = v852
	v890 = v858
	goto L192
L210:
	;
	if v873 < v804 {
		v811 = v852
		v812 = v873
		v814 = v850
		v817 = v872
		goto L199
	} else {
		goto L211
	}
L211:
	;
	goto L200
L212:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	v950 = base.B2i32(v727|v924 == int32(0))
	goto L152
L213:
	;
	if v885 == v786 {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v911 = int32(4089092)
	v913 = *(*float32)(unsafe.Add(mBase, _consts[739]))
	*(*float32)(unsafe.Add(mBase, _consts[739])) = base.F32_add(v913, base.F32_mul(base.F32_sub(base.F32_div(base.F32_convert_i32_u(v907), base.F32_convert_i32_u(v885-v786)), v913), float32(0.0625)))
	goto L212
L215:
	;
	v980 = int32(4465220)
	v982 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v983 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v982 + v983
	v987 = *(*int32)(unsafe.Add(mBase, _consts[745]))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v987)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v987)+336)) = v988 + v983
	v992 = *(*int64)(unsafe.Add(mBase, uint32(v987)+344))
	v993 = int32(4450160)
	v994 = *(*int64)(unsafe.Add(mBase, _consts[735]))
	*(*int64)(unsafe.Add(mBase, uint32(v987)+344)) = v992 + v994
	v997 = *(*int64)(unsafe.Add(mBase, uint32(v987)+352))
	v998 = int32(4450168)
	v999 = *(*int64)(unsafe.Add(mBase, _consts[733]))
	*(*int64)(unsafe.Add(mBase, uint32(v987)+352)) = v997 + v999
	v1002 = *(*int64)(unsafe.Add(mBase, uint32(v987)+360))
	v1003 = int32(4450176)
	v1004 = *(*int64)(unsafe.Add(mBase, _consts[727]))
	*(*int64)(unsafe.Add(mBase, uint32(v987)+360)) = v1002 + v1004
	*(*int32)(unsafe.Add(mBase, uint32(v987)+336)) = v988 + int32(2)
	v1012 = int64(0)
	*(*int64)(unsafe.Add(mBase, _consts[735])) = v1012
	*(*int64)(unsafe.Add(mBase, _consts[733])) = v1012
	*(*int64)(unsafe.Add(mBase, _consts[727])) = v1012
	*(*int64)(unsafe.Add(mBase, _consts[730])) = v1012
	v1025 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v1025 - v983
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_pgstat_report_wal(m, int32(1))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L219
	}
L218:
	;
	goto L217
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1043 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+4)) = int32(1)
	if v1044 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	F_s_lock(m, v1048+int32(4), int32(486761), int32(1404), int32(87239))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	*(*int32)(unsafe.Add(mBase, uint32(v1057)+4)) = int32(0)
	v1060 = int32(4378632)
	v1061 = *(*int32)(unsafe.Add(mBase, _consts[747]))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	*(*int32)(unsafe.Add(mBase, _consts[747])) = v1063
	if v1063 != v1061 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	goto L222
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_smgrdestroyall(m)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v1072 <= int32(0) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L226
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1140 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v1143 = *(*int32)(unsafe.Add(mBase, _consts[744]))
	v1145 = F_WaitLatch(m, v1140, int32(41), v1143, int32(83886083))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L240
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, _consts[189])))
	if v1080 == int32(1) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v1090 != 0 {
		goto L228
	} else {
		goto L234
	}
L231:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+316))
	v1088 = base.B2i32(v1086 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[189])) = uint8(v1088)
	v1090 = v1088
	goto L233
L232:
	;
	v1090 = int32(0)
	goto L233
L233:
	;
	goto L230
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1097 = m.G0
	v1098 = int32(16)
	v1099 = v1097 - v1098
	m.G0 = v1099
	F___gettimeofday(m, v1099)
	mBase = m.M
	v1102 = *(*int64)(unsafe.Add(mBase, uint32(v1099)))
	v1103 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1099)+8)))
	m.G0 = v1099 + v1098
	v1111 = v1103 + v1102*int64(1000000) - int64(946684800000000)
	goto L235
L235:
	;
	v1113 = *(*int64)(unsafe.Add(mBase, _consts[725]))
	if v1111 < v1113+int64(15000000) {
		goto L228
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1121 = *(*int64)(unsafe.Add(mBase, _consts[748]))
	v1122 = F_GetLastImportantRecPtr(m)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L237
	}
L237:
	;
	if base.Ui64(v1122) < base.Ui64(v1121) {
		goto L228
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1128 = F_LogStandbySnapshot(m)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L239
	}
L239:
	;
	*(*int64)(unsafe.Add(mBase, _consts[725])) = v1111
	*(*int64)(unsafe.Add(mBase, _consts[748])) = v1128
	goto L228
L240:
	;
	if v1145 != int32(8) {
		v589 = v950
		goto L138
	} else {
		goto L241
	}
L241:
	;
	if v589&v950 == int32(0) {
		v589 = v950
		goto L138
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1156 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	F_StrategyNotifyBgWriter(m, v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	v1163 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v1166 = *(*int32)(unsafe.Add(mBase, _consts[744]))
	v1170 = F_WaitLatch(m, v1163, int32(41), v1166*int32(50), int32(83886082))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v474
	F_StrategyNotifyBgWriter(m, int32(-1))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		v1192 = v476
		goto L5
	} else {
		goto L245
	}
L245:
	;
	v589 = v950
	goto L138
L246:
	;
	v1205 = int32(v1201)
	m.G0 = v1192
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1205)))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1208)))
	if v25 == v1210 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	m.ExcPending = 1
	goto L255
L248:
	;
	if v1213 != 0 {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	v1213 = v1212
	goto L251
L250:
	;
	v1213 = int32(0)
	goto L251
L251:
	;
	goto L248
L252:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = v1213
	v30 = v1215
	v31 = v1214
	v32 = v1216
	v34 = v1207
	v42 = v1192
	goto L1
L253:
	;
	goto L254
L254:
	;
	F___wasm_longjmp(m, v1208, v1207)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	return
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_BarrierArriveAndDetach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l0, int32(486974), int32(307), int32(296345))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v15 = v13 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v15 != v17 {
				v19 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
				return base.B2i32(v15 == v19)
			} else {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28 + int32(1)
				F_ConditionVariableBroadcast(m, l0+int32(24))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v15 == int32(0))
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = v13 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 != v17 {
			v19 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v19
			return base.B2i32(v15 == v19)
		} else {
			v24 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28 + int32(1)
			F_ConditionVariableBroadcast(m, l0+int32(24))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v15 == int32(0))
			}
		}
	}
}
func F_BarrierAttach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l0, int32(486974), int32(242), int32(320837))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15 + int32(1)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			return v19
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15 + int32(1)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v19
	}
}
func F_BasicOpenFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _consts[937]))
	v5 = F_BasicOpenFilePerm(m, l0, l1, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_BogusGetChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(644737), v5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(484174), int32(307), int32(61520))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_BogusRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		F_errmsg_internal(m, int32(644935), v6)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(484174), int32(299), int32(480590))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_BootStrapCommitTs(m *base.Module) {
	return
}
func F_BuildDescFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v5 = int32(0)
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = v13
	goto L3
L2:
	;
	v14 = v5
	goto L3
L3:
	;
	v19 = F_palloc(m, v14*int32(116)+int32(20))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(-4294965047)
	v35 = int32(0)
	v38 = v5
	goto L6
L6:
	;
	v43 = int32(0)
	if l0 == v43 {
		v54 = v43
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 == int32(0) {
		v63 = v43
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 <= v35 {
		v54 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = v50 + v35<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v64 = int32(0)
	if l2 == v64 {
		v75 = v64
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 <= v35 {
		v63 = v43
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v63 = v59 + v35<<(uint(int32(2))%32)
	goto L11
L14:
	;
	if l3 == int32(0) {
		v84 = v64
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v69 <= v35 {
		v75 = int32(0)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = v71 + v35<<(uint(int32(2))%32)
	goto L14
L17:
	;
	if v54 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v78 <= v35 {
		v84 = v64
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v84 = v80 + v35<<(uint(int32(2))%32)
	goto L17
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v95 = base.I32_extend16_s(v38 + int32(1))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	F_TupleDescInitEntry(m, v19, v95, v97, v98, v99, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L26
	}
L21:
	;
	return v19
L22:
	;
	if v63 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v75 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	if v84 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v19+int32(16)+v103<<(uint(int32(4))%32)+v95*int32(100)))) = v92
	v35 = v35 + int32(1)
	v38 = v95
	goto L6
}
func F_basque_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 < v12 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v512 < v12 {
		goto L149
	} else {
		goto L150
	}
L2:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v499)+8)) = v498
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v292 < v12 {
		goto L83
	} else {
		goto L84
	}
L4:
	;
	if v65 != 0 {
		goto L3
	} else {
		goto L18
	}
L5:
	;
	v24 = v12
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	goto L9
L8:
	;
	v65 = v61
	goto L4
L9:
	;
	if v12 == v24 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v61 = int32(0)
	goto L8
L11:
	;
	v65 = int32(-1)
	goto L4
L12:
	;
	goto L13
L13:
	;
	v36 = int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v12))))
	if int32(117) < v39 {
		v61 = v36
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v41 = v39 - int32(97)
	if v41 < int32(0) {
		v61 = v36
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v41)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v47)>>(uint(v41&int32(7))%32))&int32(1) == int32(0) {
		v61 = v36
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L17
L17:
	;
	goto L10
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v75 < v66 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v180 < v66 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	if v115 != 0 {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v77 = v66
	goto L23
L22:
	;
	v77 = v75
	goto L23
L23:
	;
	goto L25
L24:
	;
	v115 = v112
	goto L20
L25:
	;
	if v66 == v77 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v112 = int32(0)
	goto L24
L27:
	;
	v115 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v66))))
	if int32(117) < v90 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + int32(1)
	goto L34
L31:
	;
	v92 = v90 - int32(97)
	if v92 < int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v95 = int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v92)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v99)>>(uint(v92&int32(7))%32))&v95 != 0 {
		v112 = v95
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	goto L26
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v124 < v123 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v164 < int32(0) {
		goto L19
	} else {
		goto L51
	}
L37:
	;
	v126 = v123
	goto L39
L38:
	;
	v126 = v124
	goto L39
L39:
	;
	v133 = v123
	goto L41
L40:
	;
	v164 = v144
	goto L36
L41:
	;
	if v133 == v126 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v164 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v133))))
	if int32(117) < v139 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v156 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v156
	v133 = v156
	goto L41
L47:
	;
	v141 = v139 - int32(97)
	if v141 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v144 = int32(1)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v141)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v148)>>(uint(v141&int32(7))%32))&v144 != 0 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L46
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v167 + v164
	goto L2
L52:
	;
	if v223 != 0 {
		goto L3
	} else {
		goto L66
	}
L53:
	;
	v182 = v66
	goto L55
L54:
	;
	v182 = v180
	goto L55
L55:
	;
	goto L57
L56:
	;
	v223 = v219
	goto L52
L57:
	;
	if v66 == v182 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v219 = int32(0)
	goto L56
L59:
	;
	v223 = int32(-1)
	goto L52
L60:
	;
	goto L61
L61:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v66))))
	if int32(117) < v197 {
		v219 = v194
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v199 = v197 - int32(97)
	if v199 < int32(0) {
		v219 = v194
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v199)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v205)>>(uint(v199&int32(7))%32))&int32(1) == int32(0) {
		v219 = v194
		goto L56
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66 + int32(1)
	goto L65
L65:
	;
	goto L58
L66:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v233 < v232 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v276 < int32(0) {
		goto L3
	} else {
		goto L81
	}
L68:
	;
	v235 = v232
	goto L70
L69:
	;
	v235 = v233
	goto L70
L70:
	;
	v242 = v232
	goto L72
L71:
	;
	v276 = int32(1)
	goto L67
L72:
	;
	if v242 == v235 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v276 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248+v242))))
	if int32(117) < v250 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v252 = v250 - int32(97)
	if v252 < int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v252)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v258)>>(uint(v252&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v267 = v242 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v267
	v242 = v267
	goto L72
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v279 + v276
	goto L2
L82:
	;
	if v332 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v294 = v12
	goto L85
L84:
	;
	v294 = v292
	goto L85
L85:
	;
	goto L87
L86:
	;
	v332 = v329
	goto L82
L87:
	;
	if v12 == v294 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v329 = int32(0)
	goto L86
L89:
	;
	v332 = int32(-1)
	goto L82
L90:
	;
	goto L91
L91:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+v12))))
	if int32(117) < v307 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 + int32(1)
	goto L96
L93:
	;
	v309 = v307 - int32(97)
	if v309 < int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v312 = int32(1)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v309)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v316)>>(uint(v309&int32(7))%32))&v312 != 0 {
		v329 = v312
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L88
L97:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v342 < v333 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v447 < v333 {
		goto L132
	} else {
		goto L133
	}
L99:
	;
	if v382 != 0 {
		goto L98
	} else {
		goto L114
	}
L100:
	;
	v344 = v333
	goto L102
L101:
	;
	v344 = v342
	goto L102
L102:
	;
	goto L104
L103:
	;
	v382 = v379
	goto L99
L104:
	;
	if v333 == v344 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v379 = int32(0)
	goto L103
L106:
	;
	v382 = int32(-1)
	goto L99
L107:
	;
	goto L108
L108:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v333))))
	if int32(117) < v357 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333 + int32(1)
	goto L113
L110:
	;
	v359 = v357 - int32(97)
	if v359 < int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v362 = int32(1)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v359)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v366)>>(uint(v359&int32(7))%32))&v362 != 0 {
		v379 = v362
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	goto L105
L114:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v391 < v390 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v431 < int32(0) {
		goto L98
	} else {
		goto L130
	}
L116:
	;
	v393 = v390
	goto L118
L117:
	;
	v393 = v391
	goto L118
L118:
	;
	v400 = v390
	goto L120
L119:
	;
	v431 = v411
	goto L115
L120:
	;
	if v400 == v393 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v431 = int32(-1)
	goto L115
L123:
	;
	goto L124
L124:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v400))))
	if int32(117) < v406 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v423 = v400 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v423
	v400 = v423
	goto L120
L126:
	;
	v408 = v406 - int32(97)
	if v408 < int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v411 = int32(1)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v408)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v415)>>(uint(v408&int32(7))%32))&v411 != 0 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	goto L125
L130:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v498 = v434 + v431
	goto L2
L131:
	;
	if v490 != 0 {
		goto L1
	} else {
		goto L145
	}
L132:
	;
	v449 = v333
	goto L134
L133:
	;
	v449 = v447
	goto L134
L134:
	;
	goto L136
L135:
	;
	v490 = v486
	goto L131
L136:
	;
	if v333 == v449 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v486 = int32(0)
	goto L135
L138:
	;
	v490 = int32(-1)
	goto L131
L139:
	;
	goto L140
L140:
	;
	v461 = int32(1)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462+v333))))
	if int32(117) < v464 {
		v486 = v461
		goto L135
	} else {
		goto L141
	}
L141:
	;
	v466 = v464 - int32(97)
	if v466 < int32(0) {
		v486 = v461
		goto L135
	} else {
		goto L142
	}
L142:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v466)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v472)>>(uint(v466&int32(7))%32))&int32(1) == int32(0) {
		v486 = v461
		goto L135
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v333 + int32(1)
	goto L144
L144:
	;
	goto L137
L145:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v492 <= v491 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v498 = v491 + int32(1)
	goto L2
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v733
	v737 = v733 - int32(1)
	if v737 <= v12 {
		goto L212
	} else {
		goto L213
	}
L148:
	;
	if v552 < int32(0) {
		goto L147
	} else {
		goto L163
	}
L149:
	;
	v514 = v12
	goto L151
L150:
	;
	v514 = v512
	goto L151
L151:
	;
	v521 = v12
	goto L153
L152:
	;
	v552 = v532
	goto L148
L153:
	;
	if v521 == v514 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v552 = int32(-1)
	goto L148
L156:
	;
	goto L157
L157:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525+v521))))
	if int32(117) < v527 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v544 = v521 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v544
	v521 = v544
	goto L153
L159:
	;
	v529 = v527 - int32(97)
	if v529 < int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v532 = int32(1)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v529)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v536)>>(uint(v529&int32(7))%32))&v532 != 0 {
		goto L152
	} else {
		goto L161
	}
L161:
	;
	goto L158
L163:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v556 = v555 + v552
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v556
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v567 < v556 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v610 < int32(0) {
		goto L147
	} else {
		goto L178
	}
L165:
	;
	v569 = v556
	goto L167
L166:
	;
	v569 = v567
	goto L167
L167:
	;
	v576 = v556
	goto L169
L168:
	;
	v610 = int32(1)
	goto L164
L169:
	;
	if v576 == v569 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v610 = int32(-1)
	goto L164
L172:
	;
	goto L173
L173:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582+v576))))
	if int32(117) < v584 {
		goto L168
	} else {
		goto L174
	}
L174:
	;
	v586 = v584 - int32(97)
	if v586 < int32(0) {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v586)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v592)>>(uint(v586&int32(7))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v601 = v576 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v601
	v576 = v601
	goto L169
L178:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v614 = v613 + v610
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v614
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v616)+4)) = v614
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v626 < v625 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v666 < int32(0) {
		goto L147
	} else {
		goto L194
	}
L180:
	;
	v628 = v625
	goto L182
L181:
	;
	v628 = v626
	goto L182
L182:
	;
	v635 = v625
	goto L184
L183:
	;
	v666 = v646
	goto L179
L184:
	;
	if v635 == v628 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v666 = int32(-1)
	goto L179
L187:
	;
	goto L188
L188:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639+v635))))
	if int32(117) < v641 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v658 = v635 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v658
	v635 = v658
	goto L184
L190:
	;
	v643 = v641 - int32(97)
	if v643 < int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v646 = int32(1)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v643)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v650)>>(uint(v643&int32(7))%32))&v646 != 0 {
		goto L183
	} else {
		goto L192
	}
L192:
	;
	goto L189
L194:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v670 = v669 + v666
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v670
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v681 < v670 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v724 < int32(0) {
		goto L147
	} else {
		goto L209
	}
L196:
	;
	v683 = v670
	goto L198
L197:
	;
	v683 = v681
	goto L198
L198:
	;
	v690 = v670
	goto L200
L199:
	;
	v724 = int32(1)
	goto L195
L200:
	;
	if v690 == v683 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v724 = int32(-1)
	goto L195
L203:
	;
	goto L204
L204:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696+v690))))
	if int32(117) < v698 {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v700 = v698 - int32(97)
	if v700 < int32(0) {
		goto L199
	} else {
		goto L206
	}
L206:
	;
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v700)>>(uint(int32(3))%32)))+uint32(_consts[1467]))))
	if int32(base.Ui32(v706)>>(uint(v700&int32(7))%32))&int32(1) == int32(0) {
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v715 = v690 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v715
	v690 = v715
	goto L200
L209:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v727))) = v728 + v724
	goto L147
L210:
	;
	return v976
L211:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v816 = v810 - v812 + v815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v816
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v816 <= v819 {
		v921 = v816
		v923 = v815
		goto L241
	} else {
		goto L242
	}
L212:
	;
	v810 = v733
	v812 = v733
	goto L211
L213:
	;
	goto L214
L214:
	;
	v740 = v733
	v741 = v737
	v742 = v733
	goto L215
L215:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v741))))
	if v746&int32(224) != int32(96) {
		v810 = v740
		v812 = v742
		goto L211
	} else {
		goto L217
	}
L216:
	;
	v810 = v802
	v812 = v804
	goto L211
L217:
	;
	if int32(1)<<(uint(v746)%32)&int32(70566434) == int32(0) {
		v810 = v740
		v812 = v742
		goto L211
	} else {
		goto L218
	}
L218:
	;
	v759 = F_find_among_b(m, l0, int32(4144608), int32(109))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	return int32(0)
L220:
	;
	if v759 == int32(0) {
		v810 = v740
		v812 = v742
		goto L211
	} else {
		goto L221
	}
L221:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v765
	switch v759 - int32(1) {
	case 0:
		goto L227
	case 1:
		goto L226
	case 2:
		goto L225
	case 3:
		goto L224
	case 4:
		goto L223
	default:
		goto L222
	}
L222:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v802
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v806 = v802 - int32(1)
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v807 < v806 {
		v740 = v802
		v741 = v806
		v742 = v804
		goto L215
	} else {
		goto L240
	}
L223:
	;
	v797 = F_slice_from_s(m, l0, int32(6), int32(2144413))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L219
	} else {
		goto L238
	}
L224:
	;
	v791 = F_slice_from_s(m, l0, int32(7), int32(2144406))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L219
	} else {
		goto L236
	}
L225:
	;
	v785 = F_slice_from_s(m, l0, int32(7), int32(2144399))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L219
	} else {
		goto L234
	}
L226:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	if v765 < v777 {
		v810 = v740
		v812 = v742
		goto L211
	} else {
		goto L231
	}
L227:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)+8))
	if v765 < v770 {
		v810 = v740
		v812 = v742
		goto L211
	} else {
		goto L228
	}
L228:
	;
	v772 = F_slice_del(m, l0)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L219
	} else {
		goto L229
	}
L229:
	;
	if int32(0) <= v772 {
		goto L222
	} else {
		goto L230
	}
L230:
	;
	v976 = v772
	goto L210
L231:
	;
	v779 = F_slice_del(m, l0)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L219
	} else {
		goto L232
	}
L232:
	;
	if int32(0) <= v779 {
		goto L222
	} else {
		goto L233
	}
L233:
	;
	v976 = v779
	goto L210
L234:
	;
	if int32(0) <= v785 {
		goto L222
	} else {
		goto L235
	}
L235:
	;
	v976 = v785
	goto L210
L236:
	;
	if int32(0) <= v791 {
		goto L222
	} else {
		goto L237
	}
L237:
	;
	v976 = v791
	goto L210
L238:
	;
	if v797 < int32(0) {
		v976 = v797
		goto L210
	} else {
		goto L239
	}
L239:
	;
	goto L222
L240:
	;
	goto L216
L241:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v927 = v925 + (v921 - v923)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v927
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v927
	v931 = v927 - int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v931 <= v932 {
		goto L284
	} else {
		goto L285
	}
L242:
	;
	v822 = v816
	v824 = v815
	goto L243
L243:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826+v822-int32(1)))))
	if v830&int32(224) != int32(96) {
		v921 = v822
		v923 = v824
		goto L241
	} else {
		goto L245
	}
L244:
	;
	v921 = v915
	v923 = v917
	goto L241
L245:
	;
	if int32(1)<<(uint(v830)%32)&int32(71162402) == int32(0) {
		v921 = v822
		v923 = v824
		goto L241
	} else {
		goto L246
	}
L246:
	;
	v843 = F_find_among_b(m, l0, int32(4146800), int32(295))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L219
	} else {
		goto L247
	}
L247:
	;
	if v843 == int32(0) {
		v921 = v822
		v923 = v824
		goto L241
	} else {
		goto L248
	}
L248:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v847
	switch v843 - int32(1) {
	case 0:
		goto L259
	case 1:
		goto L258
	case 2:
		goto L257
	case 3:
		goto L256
	case 4:
		goto L255
	case 5:
		goto L254
	case 6:
		goto L253
	case 7:
		goto L252
	case 8:
		goto L251
	case 9:
		goto L250
	default:
		goto L249
	}
L249:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v915
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v918 < v915 {
		v822 = v915
		v824 = v917
		goto L243
	} else {
		goto L283
	}
L250:
	;
	v910 = F_slice_from_s(m, l0, int32(5), int32(2144913))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L219
	} else {
		goto L281
	}
L251:
	;
	v904 = F_slice_from_s(m, l0, int32(5), int32(2144908))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L219
	} else {
		goto L279
	}
L252:
	;
	v898 = F_slice_from_s(m, l0, int32(5), int32(2144903))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L219
	} else {
		goto L277
	}
L253:
	;
	v892 = F_slice_from_s(m, l0, int32(5), int32(2144898))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L219
	} else {
		goto L275
	}
L254:
	;
	v886 = F_slice_from_s(m, l0, int32(6), int32(2144892))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L219
	} else {
		goto L273
	}
L255:
	;
	v880 = F_slice_from_s(m, l0, int32(3), int32(2144889))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L219
	} else {
		goto L271
	}
L256:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	if v847 < v872 {
		v921 = v822
		v923 = v824
		goto L241
	} else {
		goto L268
	}
L257:
	;
	v867 = F_slice_from_s(m, l0, int32(3), int32(2144886))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L219
	} else {
		goto L266
	}
L258:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	if v847 < v859 {
		v921 = v822
		v923 = v824
		goto L241
	} else {
		goto L263
	}
L259:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v851)+8))
	if v847 < v852 {
		v921 = v822
		v923 = v824
		goto L241
	} else {
		goto L260
	}
L260:
	;
	v854 = F_slice_del(m, l0)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L219
	} else {
		goto L261
	}
L261:
	;
	if int32(0) <= v854 {
		goto L249
	} else {
		goto L262
	}
L262:
	;
	v976 = v854
	goto L210
L263:
	;
	v861 = F_slice_del(m, l0)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L219
	} else {
		goto L264
	}
L264:
	;
	if int32(0) <= v861 {
		goto L249
	} else {
		goto L265
	}
L265:
	;
	v976 = v861
	goto L210
L266:
	;
	if int32(0) <= v867 {
		goto L249
	} else {
		goto L267
	}
L267:
	;
	v976 = v867
	goto L210
L268:
	;
	v874 = F_slice_del(m, l0)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L219
	} else {
		goto L269
	}
L269:
	;
	if int32(0) <= v874 {
		goto L249
	} else {
		goto L270
	}
L270:
	;
	v976 = v874
	goto L210
L271:
	;
	if int32(0) <= v880 {
		goto L249
	} else {
		goto L272
	}
L272:
	;
	v976 = v880
	goto L210
L273:
	;
	if int32(0) <= v886 {
		goto L249
	} else {
		goto L274
	}
L274:
	;
	v976 = v886
	goto L210
L275:
	;
	if int32(0) <= v892 {
		goto L249
	} else {
		goto L276
	}
L276:
	;
	v976 = v892
	goto L210
L277:
	;
	if int32(0) <= v898 {
		goto L249
	} else {
		goto L278
	}
L278:
	;
	v976 = v898
	goto L210
L279:
	;
	if int32(0) <= v904 {
		goto L249
	} else {
		goto L280
	}
L280:
	;
	v976 = v904
	goto L210
L281:
	;
	if v910 < int32(0) {
		v976 = v910
		goto L210
	} else {
		goto L282
	}
L282:
	;
	goto L249
L283:
	;
	goto L244
L284:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v972
	v976 = int32(1)
	goto L210
L285:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934+v931))))
	if v936&int32(224) != int32(96) {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	if int32(1)<<(uint(v936)%32)&int32(35362) == int32(0) {
		goto L284
	} else {
		goto L287
	}
L287:
	;
	v949 = F_find_among_b(m, l0, int32(4152704), int32(19))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L219
	} else {
		goto L288
	}
L288:
	;
	if v949 == int32(0) {
		goto L284
	} else {
		goto L289
	}
L289:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v953
	switch v949 - int32(1) {
	case 0:
		goto L291
	case 1:
		goto L290
	default:
		goto L284
	}
L290:
	;
	v966 = F_slice_from_s(m, l0, int32(1), int32(2146105))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L219
	} else {
		goto L295
	}
L291:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v957)+8))
	if v953 < v958 {
		goto L284
	} else {
		goto L292
	}
L292:
	;
	v960 = F_slice_del(m, l0)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L219
	} else {
		goto L293
	}
L293:
	;
	if int32(0) <= v960 {
		goto L284
	} else {
		goto L294
	}
L294:
	;
	v976 = v960
	goto L210
L295:
	;
	if v966 < int32(0) {
		v976 = v966
		goto L210
	} else {
		goto L296
	}
L296:
	;
	goto L284
}
func F_binaryheap_add_unordered(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(116739), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(487529), int32(123), int32(443928))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, uint32(l0+v4<<(uint(int32(2))%32))+20)) = l1
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26 + int32(1)
		return
	}
}
func F_binaryheap_build(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v101 int32
	_ = v101
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = v10 - int32(2)
	if int32(-1) <= v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = l0 + int32(20)
	v18 = base.I32_div_s(v12, int32(2))
	v20 = v18
	goto L4
L2:
	;
	goto L3
L3:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v101)
	return
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16+v20<<(uint(int32(2))%32))))
	v36 = v20
	goto L6
L5:
	;
	goto L3
L6:
	;
	v41 = int32(1)
	v42 = v36 << (uint(v41) % 32)
	v44 = v42 | v41
	v46 = v42 + int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v46 < v47 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v36<<(uint(int32(2))%32)))) = v31
	if int32(0) < v20 {
		v20 = v20 - int32(1)
		goto L4
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	v49 = int32(2)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v16+v44<<(uint(v49)%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v16+v46<<(uint(v49)%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = m.T0[v58].(func(*base.Module, int32, int32, int32) int32)(m, v52, v56, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v65 = v44
	v66 = v47
	goto L11
L11:
	;
	if v66 <= v44 {
		goto L8
	} else {
		goto L17
	}
L12:
	;
	return
L13:
	;
	if v59 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = v46
	goto L16
L15:
	;
	v63 = v44
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = v63
	v66 = v64
	goto L11
L17:
	;
	v70 = v16 + v65<<(uint(int32(2))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v74 = m.T0[v73].(func(*base.Module, int32, int32, int32) int32)(m, v31, v71, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if int32(0) <= v74 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v36<<(uint(int32(2))%32)))) = v81
	v36 = v65
	goto L6
L20:
	;
	goto L5
}
func F_bitncommon(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v4 = int32(0)
	v9 = int32(8)
	v10 = base.I32_div_s(l2, v9)
	if v9 <= l2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v74 + v72<<(uint(int32(3))%32)
L2:
	;
	v60 = v51
	goto L13
L3:
	;
	v16 = v4
	goto L6
L4:
	;
	v33 = v4
	goto L5
L5:
	;
	v40 = l2 - v10<<(uint(int32(3))%32)
	if v40 == int32(0) {
		v72 = v33
		v74 = v4
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v16))))
	if v22 != v24 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v33 = v10
	goto L5
L8:
	;
	v50 = v16
	v51 = int32(7)
	v53 = v22
	v54 = v24
	goto L2
L9:
	;
	goto L10
L10:
	;
	v28 = v16 + int32(1)
	if v28 != v10 {
		v16 = v28
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v33))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v33))))
	v50 = v33
	v51 = v40
	v53 = v46
	v54 = v44
	goto L2
L13:
	;
	if int32(base.Ui32(v53^v54)>>(uint(int32(8)-v60)%32)) != 0 {
		v60 = v60 - int32(1)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v72 = v50
	v74 = v60
	goto L1
L15:
	;
	goto L14
}
func F_bitnot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v15 = F_palloc(m, int32(base.Ui32(v12)>>(uint(int32(2))%32)))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v19 = v17 & int32(-4)
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v21
			v23 = int32(8)
			v24 = v15 + v23
			v26 = v8 + v23
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			if base.Ui32(v26) < base.Ui32(v8+int32(base.Ui32(v27)>>(uint(int32(2))%32))) {
				v32 = v24
				v35 = v26
				for {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
					v40 = v38 ^ int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v40)
					v42 = int32(1)
					v43 = v32 + v42
					v45 = v35 + v42
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if base.Ui32(v45) < base.Ui32(v8+int32(base.Ui32(v46)>>(uint(int32(2))%32))) {
						v32 = v43
						v35 = v45
						continue
					} else {
						break
					}
					break
				}
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v53 = v43
				v57 = v51
				v59 = v52
			} else {
				v53 = v24
				v57 = v21
				v59 = v19
			}
			v66 = v59<<(uint(int32(1))%32)&int32(-8) - v57 + int32(-64)
			if int32(0) < v66 {
				v70 = v53 - int32(1)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
				v74 = v71 & (int32(255) << (uint(v66) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v74)
			} else {
			}
			return v15
		}
	}
}
func F_bitoctetlength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		return int32(base.Ui32(v7)>>(uint(int32(2))%32)) - int32(8)
	}
}
func F_bitsubstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v10 = F_bitsubstring(m, v3, v7, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_bitsubstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l3 != 0 {
		v30 = v12 + int32(1)
		v31 = int32(1)
		if l1 <= v31 {
			v34 = v31
		} else {
			v34 = l1
		}
		if base.B2i32(v34 <= v12)&base.B2i32(v34 < v30) == int32(0) {
			v41 = F_palloc(m, int32(8))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(32)
				return v41
			}
		} else {
			v48 = v30 - v34
			v50 = v48 + int32(7)
			v51 = int32(8)
			v52 = base.I32_div_s(v50, v51)
			v54 = v52 + v51
			v55 = F_palloc(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v48
				v59 = v54 << (uint(int32(2)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v55))) = v59
				v62 = v34 - int32(1)
				v64 = v62 & int32(7)
				if v64 == int32(0) {
					v67 = int32(8)
					if v52 != 0 {
						v74 = F__emscripten_memcpy_bulkmem(m, v55+v67, l0+int32(base.Ui32(v62)>>(uint(int32(3))%32))+v67, v52)
						mBase = m.M
					} else {
					}
					v175 = v48
					v177 = v59
				} else {
					if v50 < int32(8) {
						v175 = v48
						v177 = v59
					} else {
						v78 = int32(1)
						v80 = int32(8)
						v81 = v80 - v64
						v83 = v55 + v80
						v88 = l0 + int32(base.Ui32(v62)>>(uint(int32(3))%32)) + v80
						if base.Ui32(v80) <= base.Ui32(v48-v78) {
							v97 = v88
							v100 = v83
							v102 = int32(0)
							for {
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
								v108 = v107 << (uint(v64) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v108)
								v111 = v97 + int32(1)
								v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(v111) < base.Ui32(l0+int32(base.Ui32(v112)>>(uint(int32(2))%32))) {
									v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
									v119 = int32(base.Ui32(v117)>>(uint(v81)%32)) | v108
									*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v119)
								} else {
								}
								v122 = v100 + int32(1)
								v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
								v124 = v123 << (uint(v64) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v124)
								v126 = int32(2)
								v127 = v97 + v126
								v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(v127) < base.Ui32(l0+int32(base.Ui32(v128)>>(uint(v126)%32))) {
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
									v135 = int32(base.Ui32(v133)>>(uint(v81)%32)) | v124
									*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v135)
								} else {
								}
								v137 = int32(2)
								v138 = v100 + v137
								v140 = v102 + v137
								if v140 != v52&int32(268435454) {
									v97 = v127
									v100 = v138
									v102 = v140
									continue
								} else {
									break
								}
								break
							}
							v143 = v127
							v146 = v138
						} else {
							v143 = v88
							v146 = v83
						}
						if v52&v78 == int32(0) {
						} else {
							v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
							v156 = v155 << (uint(v64) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v156)
							v159 = v143 + int32(1)
							v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if base.Ui32(l0+int32(base.Ui32(v160)>>(uint(int32(2))%32))) <= base.Ui32(v159) {
							} else {
								v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
								v167 = int32(base.Ui32(v165)>>(uint(v81)%32)) | v156
								*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v167)
							}
						}
						v171 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						v175 = v171
						v177 = v172
					}
				}
				v185 = int32(base.Ui32(v177) >> (uint(int32(2)) % 32))
				v190 = v185<<(uint(int32(3))%32) - v175 + int32(-64)
				if int32(0) < v190 {
					v195 = v185 + v55 - int32(1)
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
					v199 = v196 & (int32(255) << (uint(v190) % 32))
					*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v199)
				} else {
				}
				return v55
			}
		}
	} else {
		if l2 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(17039490))
				mBase = m.M
				v209 = m.ExcPending
				if v209 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(432468), int32(0))
					mBase = m.M
					v213 = m.ExcPending
					if v213 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(484813), int32(1081), int32(324201))
						mBase = m.M
						v218 = m.ExcPending
						if v218 != 0 {
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
			v21 = l1 + l2
			if base.B2i32(l2 < int32(0)) != base.B2i32(v21 < l1) {
				v30 = v12 + int32(1)
			} else {
				v25 = v12 + int32(1)
				if v21 < v25 {
					v27 = v21
				} else {
					v27 = v25
				}
				v30 = v27
			}
			v31 = int32(1)
			if l1 <= v31 {
				v34 = v31
			} else {
				v34 = l1
			}
			if base.B2i32(v34 <= v12)&base.B2i32(v34 < v30) == int32(0) {
				v41 = F_palloc(m, int32(8))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v41))) = int64(32)
					return v41
				}
			} else {
				v48 = v30 - v34
				v50 = v48 + int32(7)
				v51 = int32(8)
				v52 = base.I32_div_s(v50, v51)
				v54 = v52 + v51
				v55 = F_palloc(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v48
					v59 = v54 << (uint(int32(2)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = v59
					v62 = v34 - int32(1)
					v64 = v62 & int32(7)
					if v64 == int32(0) {
						v67 = int32(8)
						if v52 != 0 {
							v74 = F__emscripten_memcpy_bulkmem(m, v55+v67, l0+int32(base.Ui32(v62)>>(uint(int32(3))%32))+v67, v52)
							mBase = m.M
						} else {
						}
						v175 = v48
						v177 = v59
					} else {
						if v50 < int32(8) {
							v175 = v48
							v177 = v59
						} else {
							v78 = int32(1)
							v80 = int32(8)
							v81 = v80 - v64
							v83 = v55 + v80
							v88 = l0 + int32(base.Ui32(v62)>>(uint(int32(3))%32)) + v80
							if base.Ui32(v80) <= base.Ui32(v48-v78) {
								v97 = v88
								v100 = v83
								v102 = int32(0)
								for {
									v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
									v108 = v107 << (uint(v64) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v108)
									v111 = v97 + int32(1)
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if base.Ui32(v111) < base.Ui32(l0+int32(base.Ui32(v112)>>(uint(int32(2))%32))) {
										v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
										v119 = int32(base.Ui32(v117)>>(uint(v81)%32)) | v108
										*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v119)
									} else {
									}
									v122 = v100 + int32(1)
									v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
									v124 = v123 << (uint(v64) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v124)
									v126 = int32(2)
									v127 = v97 + v126
									v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if base.Ui32(v127) < base.Ui32(l0+int32(base.Ui32(v128)>>(uint(v126)%32))) {
										v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
										v135 = int32(base.Ui32(v133)>>(uint(v81)%32)) | v124
										*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v135)
									} else {
									}
									v137 = int32(2)
									v138 = v100 + v137
									v140 = v102 + v137
									if v140 != v52&int32(268435454) {
										v97 = v127
										v100 = v138
										v102 = v140
										continue
									} else {
										break
									}
									break
								}
								v143 = v127
								v146 = v138
							} else {
								v143 = v88
								v146 = v83
							}
							if v52&v78 == int32(0) {
							} else {
								v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
								v156 = v155 << (uint(v64) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v156)
								v159 = v143 + int32(1)
								v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(l0+int32(base.Ui32(v160)>>(uint(int32(2))%32))) <= base.Ui32(v159) {
								} else {
									v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
									v167 = int32(base.Ui32(v165)>>(uint(v81)%32)) | v156
									*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v167)
								}
							}
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
							v172 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							v175 = v171
							v177 = v172
						}
					}
					v185 = int32(base.Ui32(v177) >> (uint(int32(2)) % 32))
					v190 = v185<<(uint(int32(3))%32) - v175 + int32(-64)
					if int32(0) < v190 {
						v195 = v185 + v55 - int32(1)
						v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
						v199 = v196 & (int32(255) << (uint(v190) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v199)
					} else {
					}
					return v55
				}
			}
		}
	}
}
func F_boolexpr_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3
	if v3 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
		return
	}
}
func F_boolsend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*uint8)(unsafe.Add(mBase, uint32(v16+v17))) = uint8(base.B2i32(v8 != int32(0)))
			v23 = v16 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v23
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = v23 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v26
		}
	}
}
func F_bottomup_sort_and_shrink_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v7 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	if v7 < v6 {
		return int32(-1)
	} else {
		if v6 < v7 {
			return int32(1)
		} else {
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			if v14 != v15 {
				v18 = int32(1)
				v20 = base.I32_extend16_s(v14)
				if v20&(v20-v18) != 0 {
					v27 = v18 << (uint(int32(32)-base.I32_clz(v20)) % 32)
				} else {
					v27 = v20
				}
				v28 = int32(1)
				v30 = base.I32_extend16_s(v15)
				if v30&(v30-v28) != 0 {
					v37 = v28 << (uint(int32(32)-base.I32_clz(v30)) % 32)
				} else {
					v37 = v30
				}
				if base.Ui32(v37) < base.Ui32(v27) {
					v51 = int32(-1)
				} else {
					if base.Ui32(v27) < base.Ui32(v37) {
						v51 = int32(1)
					} else {
						v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
						v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
						if v46 < v45 {
							v48 = int32(1)
						} else {
							v48 = int32(-1)
						}
						v51 = v48
					}
				}
			} else {
				v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
				if v46 < v45 {
					v48 = int32(1)
				} else {
					v48 = int32(-1)
				}
				v51 = v48
			}
			return v51
		}
	}
}
func F_boxes_bound_box(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v19 float64
	_ = v19
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v54 float64
	_ = v54
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_palloc(m, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
				v25 = v19
			} else {
				v25 = v13
			}
			if base.F64_lt(v13, v19) != 0 {
				v27 = v19
			} else {
				v27 = v25
			}
			v29 = v27
		} else {
			v29 = v13
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v29
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		if base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v31)&int64(9223372036854775807)) {
				v43 = v32
			} else {
				v43 = v31
			}
			if base.F64_gt(v31, v32) != 0 {
				v45 = v32
			} else {
				v45 = v43
			}
			v46 = v45
		} else {
			v46 = v31
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v46
		v48 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		if base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v54 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) {
				v60 = v54
			} else {
				v60 = v48
			}
			if base.F64_lt(v48, v54) != 0 {
				v62 = v54
			} else {
				v62 = v60
			}
			v64 = v62
		} else {
			v64 = v48
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v64
		v66 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
		v67 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
		if base.Ui64(base.I64_reinterpret_f64(v67)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) {
				v78 = v67
			} else {
				v78 = v66
			}
			if base.F64_gt(v66, v67) != 0 {
				v80 = v67
			} else {
				v80 = v78
			}
			v81 = v80
		} else {
			v81 = v66
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v81
		return v9
	}
}
func F_bpcharfastcmp_c(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	v11 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = F_pg_detoast_datum_packed(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(1)
	v20 = v15 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v21&v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v20
	goto L6
L5:
	;
	v24 = v15 + int32(4)
	goto L6
L6:
	;
	v25 = int32(1)
	v26 = v11 + v25
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v31 = v29 & v25
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = v26
	goto L9
L8:
	;
	v32 = v11 + int32(4)
	goto L9
L9:
	;
	if v29 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v63 = int32(-1)
	v65 = v60 - int32(1)
	if v63 <= v65 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v35 = int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v37&int32(254) == int32(2) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v50 = int32(1)
	if v31 != 0 {
		v60 = int32(base.Ui32(v29)>>(uint(v50)%32)) - v50
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v46 = v35
	goto L16
L15:
	;
	v46 = base.B2i32(v37 == int32(18)) << (uint(v35) % 32)
	goto L16
L16:
	;
	if v37 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v49 = v35
	goto L19
L18:
	;
	v49 = v46
	goto L19
L19:
	;
	v60 = v49
	goto L10
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v84 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v68 = v63
	goto L24
L23:
	;
	v68 = v65
	goto L24
L24:
	;
	v72 = v60
	goto L25
L25:
	;
	v76 = v72 - int32(1)
	if v76 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v83 = v72
	goto L21
L27:
	;
	v83 = v68 + int32(1)
	goto L21
L28:
	;
	goto L29
L29:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v76))))
	if v80 == int32(32) {
		v72 = v76
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	v117 = int32(-1)
	v119 = v114 - int32(1)
	if v117 <= v119 {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	v87 = int32(4)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v89&int32(254) == int32(2) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v102 = int32(1)
	if v84&v102 != 0 {
		v114 = int32(base.Ui32(v84)>>(uint(v102)%32)) - v102
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v98 = v87
	goto L37
L36:
	;
	v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
	goto L37
L37:
	;
	if v89 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v101 = v87
	goto L40
L39:
	;
	v101 = v98
	goto L40
L40:
	;
	v114 = v101
	goto L31
L41:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v114 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v138 = base.B2i32(v83 < v137)
	if v83 < v137 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v122 = v117
	goto L45
L44:
	;
	v122 = v119
	goto L45
L45:
	;
	v126 = v114
	goto L46
L46:
	;
	v130 = v126 - int32(1)
	if v130 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v137 = v126
	goto L42
L48:
	;
	v137 = v122 + int32(1)
	goto L42
L49:
	;
	goto L50
L50:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v130))))
	if v134 == int32(32) {
		v126 = v130
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v139 = v83
	goto L54
L53:
	;
	v139 = v137
	goto L54
L54:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v139) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	if l0 != v11 {
		goto L73
	} else {
		goto L74
	}
L56:
	;
	v201 = int32(0)
	goto L55
L57:
	;
	v175 = v170
	v176 = v171
	v177 = v172
	goto L67
L58:
	;
	if (v32|v24)&int32(3) != 0 {
		v170 = v32
		v171 = v24
		v172 = v139
		goto L57
	} else {
		goto L61
	}
L59:
	;
	v163 = v32
	v164 = v24
	v165 = v139
	goto L60
L60:
	;
	if v165 == int32(0) {
		goto L56
	} else {
		goto L66
	}
L61:
	;
	v147 = v32
	v148 = v24
	v149 = v139
	goto L62
L62:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v152 != v153 {
		v170 = v147
		v171 = v148
		v172 = v149
		goto L57
	} else {
		goto L64
	}
L63:
	;
	v163 = v158
	v164 = v156
	v165 = v160
	goto L60
L64:
	;
	v155 = int32(4)
	v156 = v148 + v155
	v158 = v147 + v155
	v160 = v149 - v155
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v147 = v158
		v148 = v156
		v149 = v160
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v170 = v163
	v171 = v164
	v172 = v165
	goto L57
L67:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	if v180 == v181 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v201 = v180 - v181
	goto L55
L69:
	;
	v183 = int32(1)
	v188 = v177 - v183
	if v188 != 0 {
		v175 = v175 + v183
		v176 = v176 + v183
		v177 = v188
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	goto L56
L73:
	;
	F_pfree(m, v11)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if l1 != v15 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	F_pfree(m, v15)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v201 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v210 = v201
	goto L83
L82:
	;
	v210 = base.B2i32(v137 < v83) - v138
	goto L83
L83:
	;
	return v210
}
func F_bracket(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v952 int64
	_ = v952
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1023 int32
	_ = v1023
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+6)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v14
	v18 = F_next(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v29 - int32(67) {
	case 0:
		goto L13
	default:
		goto L10
	case 2:
		goto L14
	case 6:
		goto L15
	case 15:
		goto L17
	case 26, 34:
		goto L9
	case 32:
		goto L11
	case 45:
		goto L16
	case 48:
		goto L12
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1269 != 0 {
		goto L417
	} else {
		goto L418
	}
L6:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1262 = F_range_(m, l0, v1253, v1254, v1259&int32(8))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L1
	} else {
		goto L414
	}
L7:
	;
	v1057 = F_next(m, l0)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L344
	}
L8:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1045 != int32(82) {
		v1253 = v1039
		v1254 = v1039
		goto L6
	} else {
		goto L343
	}
L9:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_okcolors(m, v756, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L255
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v752 != 0 {
		goto L252
	} else {
		goto L253
	}
L11:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v746 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12+v744))) = uint8(v746)
	v748 = F_next(m, l0)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L251
	}
L12:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v727)+8)) = v728 | int32(1024)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v735 = F_cclasscvec(m, l0, v726, v732&int32(8))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L245
	}
L13:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v551 = F_next(m, l0)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L186
	}
L14:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v228 = F_next(m, l0)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L91
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = F_next(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L32
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = F_next(m, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v36 = v34
	goto L20
L19:
	;
	v36 = int32(11)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
	goto L3
L21:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v41 != int32(82) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v44&int32(8) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v59 != 0 {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	v49 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v49)
	F_subcoloronechr(m, l0, v38, l1, l2, v12+int32(14))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v55 = F_allcases(m, l0, v38)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	F_subcolorcvec(m, l0, v55, l1, l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L3
L31:
	;
	v1051 = v38
	goto L7
L32:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v63 != int32(112) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v91 = F_next(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L41
	}
L34:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v85 = v66
	goto L33
L35:
	;
	goto L36
L36:
	;
	goto L37
L37:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = F_next(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v85 = v76
	goto L33
L39:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v79 == int32(112) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if base.Ui32(v85) <= base.Ui32(v60) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v96 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v100 != 0 {
		goto L3
	} else {
		goto L48
	}
L45:
	;
	v98 = v96
	goto L47
L46:
	;
	v98 = int32(3)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v98
	goto L3
L48:
	;
	v101 = v85 - v60
	if v101 == int32(4) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v1039 = v104
	goto L8
L50:
	;
	goto L51
L51:
	;
	v106 = v101 >> (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v108 | int32(1024)
	v117 = int32(523033)
	v118 = int32(1593776)
	goto L53
L52:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v225 != 0 {
		goto L3
	} else {
		goto L90
	}
L53:
	;
	if v117&int32(3) == int32(0) {
		v146 = v117
		goto L57
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v221 != 0 {
		goto L87
	} else {
		goto L88
	}
L55:
	;
	if v179 == v106 {
		goto L72
	} else {
		goto L73
	}
L56:
	;
	v179 = v171 - v117
	goto L55
L57:
	;
	v150 = v146
	goto L66
L58:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v130 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v179 = int32(0)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v135 = v117
	goto L62
L62:
	;
	v139 = v135 + int32(1)
	if v139&int32(3) == int32(0) {
		v146 = v139
		goto L57
	} else {
		goto L64
	}
L63:
	;
	v171 = v139
	goto L56
L64:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v144 != 0 {
		v135 = v139
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v159 = int32(-2139062144)
	if (int32(16843008)-v156|v156)&v159 == v159 {
		v150 = v150 + int32(4)
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v165 = v150
	goto L69
L68:
	;
	goto L67
L69:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v169 != 0 {
		v165 = v165 + int32(1)
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v171 = v165
	goto L56
L71:
	;
	goto L70
L72:
	;
	if v106 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	v217 = v118 + int32(8)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v218 != 0 {
		v117 = v218
		v118 = v217
		goto L53
	} else {
		goto L86
	}
L75:
	;
	if v213 == int32(0) {
		goto L52
	} else {
		goto L85
	}
L76:
	;
	v213 = int32(0)
	goto L75
L77:
	;
	v185 = v117
	v186 = v60
	v187 = v106
	goto L78
L78:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	if v190 != v191 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L76
L80:
	;
	v213 = v190 - v191
	goto L75
L81:
	;
	goto L82
L82:
	;
	if v190 == int32(0) {
		goto L76
	} else {
		goto L83
	}
L83:
	;
	v196 = int32(1)
	v201 = v187 - v196
	if v201 != 0 {
		v185 = v185 + v196
		v186 = v186 + int32(4)
		v187 = v201
		goto L78
	} else {
		goto L84
	}
L84:
	;
	goto L79
L85:
	;
	goto L74
L86:
	;
	goto L54
L87:
	;
	v223 = v221
	goto L89
L88:
	;
	v223 = int32(3)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v223
	goto L3
L90:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	v1039 = v226
	goto L8
L91:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v230 != int32(112) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v258 = F_next(m, l0)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L100
	}
L93:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v252 = v233
	goto L92
L94:
	;
	goto L95
L95:
	;
	goto L96
L96:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v244 = F_next(m, l0)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L98
	}
L97:
	;
	v252 = v243
	goto L92
L98:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v246 == int32(112) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	if base.Ui32(v252) <= base.Ui32(v227) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v263 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v267 != 0 {
		goto L3
	} else {
		goto L107
	}
L104:
	;
	v265 = v263
	goto L106
L105:
	;
	v265 = int32(3)
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v265
	goto L3
L107:
	;
	v268 = v252 - v227
	if v268 == int32(4) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v406 = v404 & int32(8)
	if v403 != int32(120) {
		goto L152
	} else {
		goto L153
	}
L109:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	v403 = v271
	goto L108
L110:
	;
	goto L111
L111:
	;
	v273 = v268 >> (uint(int32(2)) % 32)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+8)) = v275 | int32(1024)
	v284 = int32(523033)
	v285 = int32(1593776)
	goto L113
L112:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v392 != 0 {
		goto L3
	} else {
		goto L150
	}
L113:
	;
	if v284&int32(3) == int32(0) {
		v313 = v284
		goto L117
	} else {
		goto L118
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v388 != 0 {
		goto L147
	} else {
		goto L148
	}
L115:
	;
	if v346 == v273 {
		goto L132
	} else {
		goto L133
	}
L116:
	;
	v346 = v338 - v284
	goto L115
L117:
	;
	v317 = v313
	goto L126
L118:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	if v297 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v346 = int32(0)
	goto L115
L120:
	;
	goto L121
L121:
	;
	v302 = v284
	goto L122
L122:
	;
	v306 = v302 + int32(1)
	if v306&int32(3) == int32(0) {
		v313 = v306
		goto L117
	} else {
		goto L124
	}
L123:
	;
	v338 = v306
	goto L116
L124:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306))))
	if v311 != 0 {
		v302 = v306
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v326 = int32(-2139062144)
	if (int32(16843008)-v323|v323)&v326 == v326 {
		v317 = v317 + int32(4)
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v332 = v317
	goto L129
L128:
	;
	goto L127
L129:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v336 != 0 {
		v332 = v332 + int32(1)
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v338 = v332
	goto L116
L131:
	;
	goto L130
L132:
	;
	if v273 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	goto L134
L134:
	;
	v384 = v285 + int32(8)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v385 != 0 {
		v284 = v385
		v285 = v384
		goto L113
	} else {
		goto L146
	}
L135:
	;
	if v380 == int32(0) {
		goto L112
	} else {
		goto L145
	}
L136:
	;
	v380 = int32(0)
	goto L135
L137:
	;
	v352 = v284
	v353 = v227
	v354 = v273
	goto L138
L138:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	if v357 != v358 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L136
L140:
	;
	v380 = v357 - v358
	goto L135
L141:
	;
	goto L142
L142:
	;
	if v357 == int32(0) {
		goto L136
	} else {
		goto L143
	}
L143:
	;
	v363 = int32(1)
	v368 = v354 - v363
	if v368 != 0 {
		v352 = v352 + v363
		v353 = v353 + int32(4)
		v354 = v368
		goto L138
	} else {
		goto L144
	}
L144:
	;
	goto L139
L145:
	;
	goto L134
L146:
	;
	goto L114
L147:
	;
	v390 = v388
	goto L149
L148:
	;
	v390 = int32(3)
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v390
	goto L3
L150:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+4)))
	v403 = v393
	goto L108
L151:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v547 != 0 {
		goto L3
	} else {
		goto L184
	}
L152:
	;
	if v406 != 0 {
		goto L171
	} else {
		goto L172
	}
L153:
	;
	if v404&int32(4096) == int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v413 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v457 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v456 + v457
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v454)+8))
	v461 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v460+v456<<(uint(v461)%32)))) = int32(120)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v466 + v457
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v454)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v470+v466<<(uint(v461)%32)))) = int32(121)
	if v406 == int32(0) {
		v544 = v454
		goto L151
	} else {
		goto L170
	}
L156:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+4))
	if v414 < int32(4) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	v430 = F_palloc_extended(m, int32(44), int32(2))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L163
	}
L159:
	;
	F_pfree(m, v413)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v413)+16))
	if v417 < int32(0) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+24)) = int32(-1)
	v422 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v413)+12)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = v422
	v454 = v413
	goto L155
L162:
	;
	goto L158
L163:
	;
	if v430 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v430))) = int64(17179869184)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v430)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+20)) = v430 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+8)) = v430 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v430
	v454 = v430
	goto L155
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v447 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v447
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v450 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v452 = v450
	goto L169
L168:
	;
	v452 = int32(12)
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v452
	v454 = v447
	goto L155
L170:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v479 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v478 + v479
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v454)+8))
	v483 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v482+v478<<(uint(v483)%32)))) = int32(88)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v488 + v479
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v454)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v492+v488<<(uint(v483)%32)))) = int32(89)
	v544 = v454
	goto L151
L171:
	;
	v498 = F_allcases(m, l0, v403)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v500 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v544 = v498
	goto L151
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = v535 + int32(1)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v539+v535<<(uint(int32(2))%32)))) = v403
	v544 = v534
	goto L151
L176:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	if v501 <= int32(0) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	v518 = F_palloc_extended(m, int32(32), int32(2))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L183
	}
L179:
	;
	F_pfree(m, v500)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500)+16))
	if v504 < int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v500)+24)) = int32(-1)
	v509 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v500)+12)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v500))) = v509
	v534 = v500
	v535 = v509
	goto L175
L182:
	;
	goto L178
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v518)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v518))) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v518)+20)) = v518 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v518)+8)) = v518 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v518
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v534 = v518
	v535 = v533
	goto L175
L184:
	;
	F_subcolorcvec(m, l0, v544, l1, l2)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L3
L186:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v553 != int32(112) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v581 = F_next(m, l0)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L195
	}
L188:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v575 = v556
	goto L187
L189:
	;
	goto L190
L190:
	;
	goto L191
L191:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v567 = F_next(m, l0)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L193
	}
L192:
	;
	v575 = v566
	goto L187
L193:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v569 == int32(112) {
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	if base.Ui32(v575) <= base.Ui32(v550) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v586 != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v590 != 0 {
		goto L3
	} else {
		goto L202
	}
L199:
	;
	v588 = v586
	goto L201
L200:
	;
	v588 = int32(4)
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v588
	goto L3
L202:
	;
	v593 = (v575 - v550) >> (uint(int32(2)) % 32)
	v600 = int32(282460)
	v601 = int32(0)
	v602 = int32(1594544)
	goto L204
L203:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v712 != 0 {
		goto L3
	} else {
		goto L241
	}
L204:
	;
	if v600&int32(3) == int32(0) {
		v629 = v600
		goto L208
	} else {
		goto L209
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v708 != 0 {
		goto L238
	} else {
		goto L239
	}
L206:
	;
	if v662 == v593 {
		goto L223
	} else {
		goto L224
	}
L207:
	;
	v662 = v654 - v600
	goto L206
L208:
	;
	v633 = v629
	goto L217
L209:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v613 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v662 = int32(0)
	goto L206
L211:
	;
	goto L212
L212:
	;
	v618 = v600
	goto L213
L213:
	;
	v622 = v618 + int32(1)
	if v622&int32(3) == int32(0) {
		v629 = v622
		goto L208
	} else {
		goto L215
	}
L214:
	;
	v654 = v622
	goto L207
L215:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622))))
	if v627 != 0 {
		v618 = v622
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v642 = int32(-2139062144)
	if (int32(16843008)-v639|v639)&v642 == v642 {
		v633 = v633 + int32(4)
		goto L217
	} else {
		goto L219
	}
L218:
	;
	v648 = v633
	goto L220
L219:
	;
	goto L218
L220:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648))))
	if v652 != 0 {
		v648 = v648 + int32(1)
		goto L220
	} else {
		goto L222
	}
L221:
	;
	v654 = v648
	goto L207
L222:
	;
	goto L221
L223:
	;
	if v593 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	goto L225
L225:
	;
	v700 = v602 + int32(4)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	v703 = v601 + int32(1)
	if v703 != int32(14) {
		v600 = v701
		v601 = v703
		v602 = v700
		goto L204
	} else {
		goto L237
	}
L226:
	;
	if v696 == int32(0) {
		goto L203
	} else {
		goto L236
	}
L227:
	;
	v696 = int32(0)
	goto L226
L228:
	;
	v668 = v600
	v669 = v550
	v670 = v593
	goto L229
L229:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668))))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	if v673 != v674 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L227
L231:
	;
	v696 = v673 - v674
	goto L226
L232:
	;
	goto L233
L233:
	;
	if v673 == int32(0) {
		goto L227
	} else {
		goto L234
	}
L234:
	;
	v679 = int32(1)
	v684 = v670 - v679
	if v684 != 0 {
		v668 = v668 + v679
		v669 = v669 + int32(4)
		v670 = v684
		goto L229
	} else {
		goto L235
	}
L235:
	;
	goto L230
L236:
	;
	goto L225
L237:
	;
	goto L205
L238:
	;
	v710 = v708
	goto L240
L239:
	;
	v710 = int32(4)
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v710
	goto L3
L241:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+8)) = v714 | int32(1024)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v721 = F_cclasscvec(m, l0, v601, v718&int32(8))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v723 != 0 {
		goto L3
	} else {
		goto L243
	}
L243:
	;
	F_subcolorcvec(m, l0, v721, l1, l2)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	goto L3
L245:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v737 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	F_subcolorcvec(m, l0, v735, l1, l2)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	v742 = F_next(m, l0)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L250
	}
L249:
	;
	goto L248
L250:
	;
	goto L3
L251:
	;
	goto L3
L252:
	;
	v754 = v752
	goto L254
L253:
	;
	v754 = int32(15)
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v754
	goto L3
L255:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v760 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	m.G0 = v12 + int32(16)
	return
L257:
	;
	v761 = int32(0)
	v766 = v761
	v767 = v761
	goto L259
L258:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)+20))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v794)+12))
	v797 = int32(24)
	v799 = v795 + v796*v797
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v802 != 0 {
		goto L269
	} else {
		goto L270
	}
L259:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766+v12))))
	if v773 == int32(1) {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	if v767&int32(1) == int32(0) {
		goto L256
	} else {
		goto L268
	}
L261:
	;
	F_charclasscomplement(m, l0, v766, l1, l2)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v785 = v766 + int32(1)
	if v785 != int32(14) {
		v766 = v785
		goto L259
	} else {
		goto L267
	}
L264:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v778 != 0 {
		goto L256
	} else {
		goto L265
	}
L265:
	;
	v779 = int32(1)
	v781 = v766 + v779
	if v781 != int32(14) {
		v766 = v781
		v767 = v779
		goto L259
	} else {
		goto L266
	}
L266:
	;
	goto L258
L267:
	;
	goto L260
L268:
	;
	goto L258
L269:
	;
	v807 = v802
	goto L272
L270:
	;
	v830 = v795
	goto L271
L271:
	;
	if base.Ui32(v830) < base.Ui32(v799+v797) {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+20))
	v814 = int32(*(*int16)(unsafe.Add(mBase, uint32(v807)+4)))
	v819 = v813 + v814*int32(24) + int32(20)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)))
	*(*int32)(unsafe.Add(mBase, uint32(v819))) = v820 | int32(4)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v807)+16))
	if v824 != 0 {
		v807 = v824
		goto L272
	} else {
		goto L274
	}
L273:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v825)+20))
	v830 = v826
	goto L271
L274:
	;
	goto L273
L275:
	;
	v841 = v830
	v844 = int32(1)
	goto L278
L276:
	;
	goto L277
L277:
	;
	goto L286
L278:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v841)+20))
	if v847&int32(4) != 0 {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	if v858&int32(1) == int32(0) {
		goto L256
	} else {
		goto L285
	}
L280:
	;
	if base.Ui32(v841) < base.Ui32(v799) {
		v841 = v841 + int32(24)
		v844 = v858
		goto L278
	} else {
		goto L284
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v841)+20)) = v847 & int32(-5)
	v858 = v844
	goto L280
L282:
	;
	goto L283
L283:
	;
	v858 = base.B2i32(v847&int32(3) != int32(0)) & v844
	goto L280
L284:
	;
	goto L279
L285:
	;
	goto L277
L286:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v884 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v963 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v963 != 0 {
		goto L318
	} else {
		goto L319
	}
L288:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v884)+8))
	v893 = int32(*(*int16)(unsafe.Add(mBase, uint32(v884)+4)))
	if v893 < int32(0) {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	goto L290
L290:
	;
	goto L287
L291:
	;
	goto L286
L292:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v884)+16))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v884)+20))
	if v927 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L293:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v884)))
	v898 = v896 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v898) {
		goto L292
	} else {
		goto L294
	}
L294:
	;
	if int32(1)<<(uint(v898)%32)&int32(163841) == int32(0) {
		goto L292
	} else {
		goto L295
	}
L295:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v885)+80))
	if v907 != 0 {
		goto L292
	} else {
		goto L296
	}
L296:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v884)+36))
	if v908 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	if v920 != 0 {
		goto L301
	} else {
		goto L302
	}
L298:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v885)+52))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+20))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v884)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v912+v893*int32(24))+12)) = v916
	v920 = v916
	goto L297
L299:
	;
	goto L300
L300:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v884)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v908)+32)) = v918
	v920 = v918
	goto L297
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920)+36)) = v908
	goto L303
L302:
	;
	goto L303
L303:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v884)+32)) = int64(0)
	goto L292
L304:
	;
	if v926 != 0 {
		goto L308
	} else {
		goto L309
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v892)+20)) = v926
	goto L304
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v927)+16)) = v926
	goto L304
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v926)+20)) = v927
	goto L310
L309:
	;
	goto L310
L310:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v892)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v892)+12)) = v933 - int32(1)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v884)+24))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v884)+28))
	if v938 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	v944 = v884 + int32(8)
	if v937 != 0 {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v891)+16)) = v937
	goto L311
L313:
	;
	goto L314
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+24)) = v937
	goto L311
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v937)+28)) = v938
	goto L317
L316:
	;
	goto L317
L317:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v891)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v891)+8)) = v946 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = int32(0)
	v952 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v944)+16)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v944)+8)) = v952
	*(*int64)(unsafe.Add(mBase, uint32(v944))) = v952
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v885)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+16)) = v958
	*(*int32)(unsafe.Add(mBase, uint32(v885)+32)) = v884
	goto L291
L318:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v966 <= v967 {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	goto L320
L322:
	;
	F_createarc(m, v961, int32(112), int32(-2), l1, l2)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L342
	}
L323:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v969 == int32(0) {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v990 == int32(0) {
		goto L322
	} else {
		goto L334
	}
L326:
	;
	v972 = v969
	goto L327
L327:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v972)+12))
	if v981 != l2 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	goto L322
L329:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v972)+16))
	if v989 != 0 {
		v972 = v989
		goto L327
	} else {
		goto L333
	}
L330:
	;
	v983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v972)+4)))
	if v983 != int32(65534) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v972)))
	if v986 == int32(112) {
		goto L256
	} else {
		goto L332
	}
L332:
	;
	goto L329
L333:
	;
	goto L328
L334:
	;
	v993 = v990
	goto L335
L335:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v993)+8))
	if v1002 != l1 {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	goto L322
L337:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v993)+24))
	if v1010 != 0 {
		v993 = v1010
		goto L335
	} else {
		goto L341
	}
L338:
	;
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993)+4)))
	if v1004 != int32(65534) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1007 == int32(112) {
		goto L256
	} else {
		goto L340
	}
L340:
	;
	goto L337
L341:
	;
	goto L336
L342:
	;
	goto L256
L343:
	;
	v1051 = v1039
	goto L7
L344:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v1059 - int32(73) {
	case 0:
		goto L346
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L5
	case 9:
		goto L347
	default:
		goto L348
	}
L345:
	;
	if v1051 == v1239 {
		goto L411
	} else {
		goto L412
	}
L346:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1069 = F_next(m, l0)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L352
	}
L347:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1065 = F_next(m, l0)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	if v1059 != int32(112) {
		goto L5
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1067 != 0 {
		goto L3
	} else {
		goto L351
	}
L351:
	;
	v1239 = v1064
	goto L345
L352:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1071 != int32(112) {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v1099 = F_next(m, l0)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L361
	}
L354:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1094 = v1074
	goto L353
L355:
	;
	goto L356
L356:
	;
	goto L357
L357:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1085 = F_next(m, l0)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L359
	}
L358:
	;
	v1094 = v1084
	goto L353
L359:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1087 == int32(112) {
		goto L357
	} else {
		goto L360
	}
L360:
	;
	goto L358
L361:
	;
	if base.Ui32(v1094) <= base.Ui32(v1068) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1104 != 0 {
		goto L365
	} else {
		goto L366
	}
L363:
	;
	goto L364
L364:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1108 != 0 {
		goto L3
	} else {
		goto L368
	}
L365:
	;
	v1106 = v1104
	goto L367
L366:
	;
	v1106 = int32(3)
	goto L367
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1106
	goto L3
L368:
	;
	v1109 = v1094 - v1068
	if v1109 == int32(4) {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1068)))
	v1239 = v1112
	goto L345
L370:
	;
	goto L371
L371:
	;
	v1114 = v1109 >> (uint(int32(2)) % 32)
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+8)) = v1116 | int32(1024)
	v1126 = int32(523033)
	v1127 = int32(1593776)
	goto L373
L372:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1233 != 0 {
		goto L3
	} else {
		goto L410
	}
L373:
	;
	if v1126&int32(3) == int32(0) {
		v1154 = v1126
		goto L377
	} else {
		goto L378
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1229 != 0 {
		goto L407
	} else {
		goto L408
	}
L375:
	;
	if v1187 == v1114 {
		goto L392
	} else {
		goto L393
	}
L376:
	;
	v1187 = v1179 - v1126
	goto L375
L377:
	;
	v1158 = v1154
	goto L386
L378:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126))))
	if v1138 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1187 = int32(0)
	goto L375
L380:
	;
	goto L381
L381:
	;
	v1143 = v1126
	goto L382
L382:
	;
	v1147 = v1143 + int32(1)
	if v1147&int32(3) == int32(0) {
		v1154 = v1147
		goto L377
	} else {
		goto L384
	}
L383:
	;
	v1179 = v1147
	goto L376
L384:
	;
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147))))
	if v1152 != 0 {
		v1143 = v1147
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1158)))
	v1167 = int32(-2139062144)
	if (int32(16843008)-v1164|v1164)&v1167 == v1167 {
		v1158 = v1158 + int32(4)
		goto L386
	} else {
		goto L388
	}
L387:
	;
	v1173 = v1158
	goto L389
L388:
	;
	goto L387
L389:
	;
	v1177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173))))
	if v1177 != 0 {
		v1173 = v1173 + int32(1)
		goto L389
	} else {
		goto L391
	}
L390:
	;
	v1179 = v1173
	goto L376
L391:
	;
	goto L390
L392:
	;
	if v1114 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L393:
	;
	goto L394
L394:
	;
	v1225 = v1127 + int32(8)
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)))
	if v1226 != 0 {
		v1126 = v1226
		v1127 = v1225
		goto L373
	} else {
		goto L406
	}
L395:
	;
	if v1221 == int32(0) {
		goto L372
	} else {
		goto L405
	}
L396:
	;
	v1221 = int32(0)
	goto L395
L397:
	;
	v1193 = v1126
	v1194 = v1068
	v1195 = v1114
	goto L398
L398:
	;
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193))))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1194)))
	if v1198 != v1199 {
		goto L400
	} else {
		goto L401
	}
L399:
	;
	goto L396
L400:
	;
	v1221 = v1198 - v1199
	goto L395
L401:
	;
	goto L402
L402:
	;
	if v1198 == int32(0) {
		goto L396
	} else {
		goto L403
	}
L403:
	;
	v1204 = int32(1)
	v1209 = v1195 - v1204
	if v1209 != 0 {
		v1193 = v1193 + v1204
		v1194 = v1194 + int32(4)
		v1195 = v1209
		goto L398
	} else {
		goto L404
	}
L404:
	;
	goto L399
L405:
	;
	goto L394
L406:
	;
	goto L374
L407:
	;
	v1231 = v1229
	goto L409
L408:
	;
	v1231 = int32(3)
	goto L409
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1231
	goto L3
L410:
	;
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127)+4)))
	v1239 = v1234
	goto L345
L411:
	;
	v1253 = v1051
	v1254 = v1051
	goto L6
L412:
	;
	goto L413
L413:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1245)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1245)+8)) = v1246 | int32(512)
	v1253 = v1051
	v1254 = v1239
	goto L6
L414:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1264 != 0 {
		goto L3
	} else {
		goto L415
	}
L415:
	;
	F_subcolorcvec(m, l0, v1262, l1, l2)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	goto L3
L417:
	;
	v1271 = v1269
	goto L419
L418:
	;
	v1271 = int32(11)
	goto L419
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1271
	goto L3
}
func F_brinbuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v17 = v13 | v14<<(uint(v10)%32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	if base.Ui32(v18+v19-int32(1)) < base.Ui32(v17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = v18
	goto L4
L2:
	;
	goto L3
L3:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v77 = F_add_values_to_range(m, l0, v75, v76, l2, l3)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L12
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v38 = F_brin_form_tuple(m, v34, v27, v35, v11+int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v45 = F_brin_doinsert(m, v40, v41, v42, l5+int32(24), v43, v38, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v47, float64(1))
	F_pfree(m, v38)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v53 + v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	F_brin_memtuple_initialize(m, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	if base.Ui32(v61+v62-int32(1)) < base.Ui32(v17) {
		v27 = v61
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_brinbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	if l1 != 0 {
		v10 = l1
		return v10
	} else {
		v6 = F_palloc0(m, int32(40))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = v6
			return v10
		}
	}
}
func F_brinvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v3 = int32(0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v13 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v211 = l1
	goto L3
L3:
	;
	return v211
L4:
	;
	v19 = F_palloc0(m, int32(40))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v23 = l1
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_RelationGetNumberOfBlocksInFork(m, v24, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v23 = v19
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v32 = F_IndexGetRelation(m, v30, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v35 = F_table_open(m, v32, int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = F_RelationGetNumberOfBlocksInFork(m, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v3
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_FreeSpaceMapVacuum(m, v38)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L62
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v58 = int32(0)
	v60 = F_ReadBufferExtended(m, v38, v58, v47, v58, v37)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	F_ReleaseBuffer(m, v60)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L60
	}
L23:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+14)))
	if v80 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v60 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65+(v60^int32(-1))<<(uint(int32(2))%32))))
	v79 = v71
	goto L23
L26:
	;
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v79 = v73 + v60<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L28:
	;
	F_LockRelationForExtension(m, v38, int32(5))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v60 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	F_UnlockRelationForExtension(m, v38, int32(5))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_LockBuffer(m, v60, int32(2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+14)))
	if v92 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_brin_initialize_empty_new_buffer(m, v38, v60)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_LockBuffer(m, v60, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	F_LockBuffer(m, v60, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	goto L22
L39:
	;
	goto L30
L40:
	;
	goto L22
L41:
	;
	if v132 == int32(61586) {
		goto L40
	} else {
		goto L47
	}
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106+(v60^int32(-1))<<(uint(int32(2))%32))))
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112+v113)+6)))
	if v115 != int32(61585) {
		v132 = v115
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v122 = v119 + v60<<(uint(int32(13))%32)
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122-int32(8176)))))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+v125-int32(8186)))))
	if v129 == int32(61585) {
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L40
L46:
	;
	v132 = v129
	goto L41
L47:
	;
	if v60 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v154 = int32(0)
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+16)))
	v156 = v79 + v155
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+6)))
	if v157 != int32(61587) {
		v172 = v154
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138+(v60^int32(-1))<<(uint(int32(6))%32))+16))
	v153 = v144
	goto L48
L50:
	;
	goto L51
L51:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146+v60<<(uint(int32(6))%32)+int32(-64))+16))
	v153 = v152
	goto L48
L52:
	;
	F_RecordPageWithFreeSpace(m, v38, v153, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L59
	}
L53:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)))
	if v160&int32(1) != 0 {
		v172 = v154
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v163 = int32(4)
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+14)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+12)))
	v166 = v164 - v165
	if v166 <= v163 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v172 = v169 - int32(4)
	goto L52
L56:
	;
	v169 = v163
	goto L58
L57:
	;
	v169 = v166
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L40
L60:
	;
	v184 = v47 + int32(1)
	if v184 != v40 {
		v47 = v184
		goto L16
	} else {
		goto L61
	}
L61:
	;
	goto L17
L62:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = v23 + int32(8)
	F_brinsummarize(m, v200, v35, int32(-1), int32(0), v204, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_sequence_close(m, v35, int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v211 = v23
	goto L3
}
func F_btadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v13 = int32(0)
	goto L3
L3:
	;
	v14 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v10 = F_get_opclass_input_type(m, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v13 = v10
	goto L3
L7:
	;
	return
L8:
	;
	if v14 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v19 <= v18 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v23 = l1
	v24 = v18
	v26 = v13
	goto L11
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v24<<(uint(int32(2))%32))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L7
L13:
	;
	v64 = v24 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v64 < v65 {
		v23 = v60
		v24 = v64
		v26 = v61
		goto L11
	} else {
		goto L27
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = l0
	v58 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v58)
	v60 = v23
	v61 = v26
	goto L13
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v37 != int32(1) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v40 != v41 {
		goto L14
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	if v26 != v40 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v45 = F_opclass_for_family_datatype(m, int32(403), l0, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v47 = v23
	v48 = v26
	goto L22
L22:
	;
	if v47 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v47 = v45
	v48 = v40
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = v47
	v50 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v50)
	v60 = v47
	v61 = v48
	goto L13
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+28)) = l0
	v53 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+24)) = uint16(v53)
	v60 = int32(0)
	v61 = v48
	goto L13
L27:
	;
	goto L12
}
func F_btbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	v5 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_palloc(m, int32(27320))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int64(-4294967296)
			*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[133]))) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v12
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 <= int32(0) {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
				v26 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v26
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v10
				*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v34
				return v5
			} else {
				v22 = F_palloc(m, v16*int32(48))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = v22
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
					v26 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = v26
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v26
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v26
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v26
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v10
					*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v34
					return v5
				}
			}
		}
	}
}
func F_btboolskipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(193)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(4294967296)
	return int32(0)
}
func F_btendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
	if v4 == int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[133])))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+36))
	if int32(0) < v7 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F__bt_killitems(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+56))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	goto L5
L8:
	;
	F_ReleaseBuffer(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(0)
	goto L1
L10:
	;
	F_ReleaseBuffer(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_consts[133]))) = int32(0)
	goto L12
L14:
	;
	F_pfree(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	if v30 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_MemoryContextDelete(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v33 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_pfree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v3)+44))
	if v36 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	F_pfree(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_pfree(m, v3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	return
}
func F_btfloat48cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v7 float64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = base.F64_promote_f32(v6)
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v15 = base.I64_reinterpret_f64(v12) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v24 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v24
	} else {
		v20 = int32(1)
		if base.F64_gt(v7, v12) != 0 {
			v32 = v20
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v32 = v20
			} else {
				v24 = v20
				v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v24
			}
		}
	}
	return v32
}
func F_btgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v3)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L1
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	if v14 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v57 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v17 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v49 = F__bt_first(m, l0, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L17
	}
L7:
	;
	v43 = F__bt_next(m, l0, l1)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L15
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v24 = F_palloc(m, int32(5432))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v29 = v20
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if int32(1357) < v30 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v24
	v29 = v24
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v30 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(2))%32)))) = v39
	goto L7
L15:
	;
	if v43 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	return int32(1)
L17:
	;
	if v49 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	return int32(1)
L19:
	;
	v58 = F__bt_start_prim_scan(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L2
L22:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L21
}
func F_btnametextcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6&int32(3) == int32(0) {
		v35 = v6
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v69 = int32(1)
	v70 = v8 + v69
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v75 = v73 & v69
	if v75 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v68 = v60 - v6
	goto L3
L5:
	;
	v39 = v35
	goto L14
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v68 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v24 = v6
	goto L10
L10:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v60 = v28
	goto L4
L12:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v54 = v39
	goto L17
L16:
	;
	goto L15
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v60 = v54
	goto L4
L19:
	;
	goto L18
L20:
	;
	v76 = v70
	goto L22
L21:
	;
	v76 = v8 + int32(4)
	goto L22
L22:
	;
	if v73 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v106 = F_varstr_cmp(m, v6, v68, v76, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L34
	}
L24:
	;
	v79 = int32(4)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v81&int32(254) == int32(2) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v94 = int32(1)
	if v75 != 0 {
		v104 = int32(base.Ui32(v73)>>(uint(v94)%32)) - v94
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v90 = v79
	goto L29
L28:
	;
	v90 = base.B2i32(v81 == int32(18)) << (uint(v79) % 32)
	goto L29
L29:
	;
	if v81 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v79
	goto L32
L31:
	;
	v93 = v90
	goto L32
L32:
	;
	v104 = v93
	goto L23
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v104 = int32(base.Ui32(v98)>>(uint(int32(2))%32)) - int32(4)
	goto L23
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v108 != v8 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_pfree(m, v8)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return v106
L38:
	;
	goto L37
}
func F_btoidcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v4) < base.Ui32(v3)) - base.B2i32(base.Ui32(v3) < base.Ui32(v4))
}
func F_btoidskipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(204)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(205)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(-4294967296)
	return int32(0)
}
func F_btoptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(4), int32(24), int32(732784), int32(3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_btrecordimagecmp(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_bttext_pattern_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v3 = int32(4470560)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v7
	F_varstr_sortsupport(m, v6, int32(25), int32(950))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v4
		return int32(0)
	}
}
func F_bttextnamecmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v113 int32
	_ = v113
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(1)
	v14 = v9 + v13
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v19 = v17 & v13
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v14
	goto L5
L4:
	;
	v20 = v9 + int32(4)
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v17 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v21&int32(3) == int32(0) {
		v73 = v21
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v24 = int32(4)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v26&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v39 = int32(1)
	if v19 != 0 {
		v49 = int32(base.Ui32(v17)>>(uint(v39)%32)) - v39
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v35 = v24
	goto L12
L11:
	;
	v35 = base.B2i32(v26 == int32(18)) << (uint(v24) % 32)
	goto L12
L12:
	;
	if v26 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = v24
	goto L15
L14:
	;
	v38 = v35
	goto L15
L15:
	;
	v49 = v38
	goto L6
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = F_varstr_cmp(m, v20, v49, v21, v106, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L18:
	;
	v106 = v98 - v21
	goto L17
L19:
	;
	v77 = v73
	goto L28
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v57 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v106 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v62 = v21
	goto L24
L24:
	;
	v66 = v62 + int32(1)
	if v66&int32(3) == int32(0) {
		v73 = v66
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v98 = v66
	goto L18
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != 0 {
		v62 = v66
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v86 = int32(-2139062144)
	if (int32(16843008)-v83|v83)&v86 == v86 {
		v77 = v77 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v92 = v77
	goto L31
L30:
	;
	goto L29
L31:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v96 != 0 {
		v92 = v92 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v98 = v92
	goto L18
L33:
	;
	goto L32
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v110 != v9 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_pfree(m, v9)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return v108
L38:
	;
	goto L37
}
func F_build_coercion_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l2 != 0 {
		v17 = F_SearchSysCache1(m, int32(47), l2)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l2
					F_errmsg_internal(m, int32(44089), v14+int32(16))
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490929), int32(853), int32(264978))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23+v24)+104)))
				F_ReleaseCatCache(m, v17)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v30 = v26
					switch l1 - int32(1) {
					case 0:
						*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l0
						v38 = F_list_make1_impl(m, int32(1), v14+int32(12))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v30 < int32(2) {
								v67 = v38
								v69 = F_makeFuncExpr(m, l2, l3, v67, int32(0), l6)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v69)+32)) = l7
									v138 = v69
									m.G0 = v14 + int32(32)
									return v138
								}
							} else {
								v44 = int32(0)
								v48 = F_makeConst(m, int32(23), int32(-1), v44, int32(4), l4, v44, int32(1))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v50 = F_lappend(m, v38, v48)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										if v30 != int32(3) {
											v67 = v50
											v69 = F_makeFuncExpr(m, l2, l3, v67, int32(0), l6)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v69)+32)) = l7
												v138 = v69
												m.G0 = v14 + int32(32)
												return v138
											}
										} else {
											v56 = int32(0)
											v57 = int32(1)
											v62 = F_makeConst(m, int32(16), int32(-1), v56, v57, base.B2i32(l5 == int32(3)), v56, v57)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v64 = F_lappend(m, v50, v62)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													v67 = v64
													v69 = F_makeFuncExpr(m, l2, l3, v67, int32(0), l6)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v69)+32)) = l7
														v138 = v69
														m.G0 = v14 + int32(32)
														return v138
													}
												}
											}
										}
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
							F_errmsg_internal(m, int32(264951), v14)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490929), int32(996), int32(264978))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 2:
						v73 = F_palloc0(m, int32(32))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(29)
							v78 = F_palloc0(m, int32(16))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(34)
								v82 = F_exprTypmod(m, l0)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v82
									v85 = F_exprType(m, l0)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v89 = F_getBaseTypeAndTypmod(m, v85, v14+int32(24))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											v91 = F_get_element_type(m, v89)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v91
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v94
												v98 = F_get_element_type(m, l3)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
													v102 = F_coerce_to_target_type(m, int32(0), v78, v101, v98, l4, l5, l6, l7)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														if v102 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(440092), int32(0))
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(490929), int32(961), int32(264978))
																	mBase = m.M
																	v171 = m.ExcPending
																	if v171 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = l3
															*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v102
															*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = l0
															v109 = F_exprTypmod(m, v102)
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v73)+28)) = l7
																*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
																*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v109
																v138 = v73
																m.G0 = v14 + int32(32)
																return v138
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
					case 3:
						v128 = F_palloc0(m, int32(24))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = l6
							*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(28)
							v138 = v128
							m.G0 = v14 + int32(32)
							return v138
						}
					}
				}
			}
		}
	} else {
		v30 = int32(0)
		switch l1 - int32(1) {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l0
			v38 = F_list_make1_impl(m, int32(1), v14+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				if v30 < int32(2) {
					v67 = v38
					v69 = F_makeFuncExpr(m, l2, l3, v67, int32(0), l6)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v69)+32)) = l7
						v138 = v69
						m.G0 = v14 + int32(32)
						return v138
					}
				} else {
					v44 = int32(0)
					v48 = F_makeConst(m, int32(23), int32(-1), v44, int32(4), l4, v44, int32(1))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = F_lappend(m, v38, v48)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							if v30 != int32(3) {
								v67 = v50
								v69 = F_makeFuncExpr(m, l2, l3, v67, int32(0), l6)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v69)+32)) = l7
									v138 = v69
									m.G0 = v14 + int32(32)
									return v138
								}
							} else {
								v56 = int32(0)
								v57 = int32(1)
								v62 = F_makeConst(m, int32(16), int32(-1), v56, v57, base.B2i32(l5 == int32(3)), v56, v57)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = F_lappend(m, v50, v62)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v67 = v64
										v69 = F_makeFuncExpr(m, l2, l3, v67, int32(0), l6)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v69)+32)) = l7
											v138 = v69
											m.G0 = v14 + int32(32)
											return v138
										}
									}
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v117 = m.ExcPending
			if v117 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
				F_errmsg_internal(m, int32(264951), v14)
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490929), int32(996), int32(264978))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 2:
			v73 = F_palloc0(m, int32(32))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(29)
				v78 = F_palloc0(m, int32(16))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(34)
					v82 = F_exprTypmod(m, l0)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v82
						v85 = F_exprType(m, l0)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v89 = F_getBaseTypeAndTypmod(m, v85, v14+int32(24))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								v91 = F_get_element_type(m, v89)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v91
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v94
									v98 = F_get_element_type(m, l3)
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return int32(0)
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
										v102 = F_coerce_to_target_type(m, int32(0), v78, v101, v98, l4, l5, l6, l7)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
										} else {
											if v102 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(440092), int32(0))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(490929), int32(961), int32(264978))
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = l3
												*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v102
												*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = l0
												v109 = F_exprTypmod(m, v102)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v73)+28)) = l7
													*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l6
													*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v109
													v138 = v73
													m.G0 = v14 + int32(32)
													return v138
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
		case 3:
			v128 = F_palloc0(m, int32(24))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(28)
				v138 = v128
				m.G0 = v14 + int32(32)
				return v138
			}
		}
	}
}
func F_build_tlist_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	if l0 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = int32(12)
		v15 = v10*v11 + v11
	} else {
		v15 = int32(12)
	}
	v16 = F_palloc(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)) = uint16(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
		v24 = v16 + int32(12)
		if l0 == v20 {
			v73 = v24
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v27 <= int32(0) {
				v73 = v24
			} else {
				v31 = v24
				v34 = int32(0)
				for {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v34<<(uint(int32(2))%32))))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					if v43 == int32(0) {
						v64 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)) = uint8(v64)
						v66 = v31
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v46 != int32(319) {
							if v46 != int32(6) {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)) = uint8(v64)
								v66 = v31
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v31))) = v51
								v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+8)))
								*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v53)
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
								*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)) = uint16(v55)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v57
								v66 = v31 + int32(12)
							}
						} else {
							v61 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v61)
							v66 = v31
						}
					}
					v69 = v34 + int32(1)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v69 < v70 {
						v31 = v66
						v34 = v69
						continue
					} else {
						break
					}
					break
				}
				v73 = v66
			}
		}
		v82 = base.I32_div_s(v73-v24, int32(12))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v82
		return v16
	}
}
func F_byte_increment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v3 != int32(255) {
		v7 = v3 + int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v7)
	} else {
	}
	return base.B2i32(v3 != int32(255))
}
func F_byteagt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v21&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = v19
	goto L10
L9:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L10
L10:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v19
	goto L13
L12:
	;
	v33 = v30
	goto L13
L13:
	;
	v46 = v33
	goto L4
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v78 = int32(1)
	if v16&v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v50 = int32(4)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v65 = int32(1)
	if v47&v65 != 0 {
		v77 = int32(base.Ui32(v47)>>(uint(v65)%32)) - v65
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v61 = v50
	goto L21
L20:
	;
	v61 = base.B2i32(v52 == int32(18)) << (uint(v50) % 32)
	goto L21
L21:
	;
	if v52 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = v50
	goto L24
L23:
	;
	v64 = v61
	goto L24
L24:
	;
	v77 = v64
	goto L15
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v82 = v78
	goto L28
L27:
	;
	v82 = int32(4)
	goto L28
L28:
	;
	v83 = v9 + v82
	v84 = int32(1)
	if v47&v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v84
	goto L31
L30:
	;
	v88 = int32(4)
	goto L31
L31:
	;
	v89 = v14 + v88
	if v46 < v77 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v91 = v46
	goto L34
L33:
	;
	v91 = v77
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v153 = int32(0)
	goto L35
L37:
	;
	v127 = v122
	v128 = v123
	v129 = v124
	goto L47
L38:
	;
	if (v83|v89)&int32(3) != 0 {
		v122 = v83
		v123 = v89
		v124 = v91
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v115 = v83
	v116 = v89
	v117 = v91
	goto L40
L40:
	;
	if v117 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v99 = v83
	v100 = v89
	v101 = v91
	goto L42
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != v105 {
		v122 = v99
		v123 = v100
		v124 = v101
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v115 = v110
	v116 = v108
	v117 = v112
	goto L40
L44:
	;
	v107 = int32(4)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L37
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 == v133 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v153 = v132 - v133
	goto L35
L49:
	;
	v135 = int32(1)
	v140 = v129 - v135
	if v140 != 0 {
		v127 = v127 + v135
		v128 = v128 + v135
		v129 = v140
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v158 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v162 = int32(0)
	return base.B2i32(v153 == v162)&base.B2i32(v77 < v46) | base.B2i32(v162 < v153)
L60:
	;
	goto L59
}
func F_byteain(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 != int32(92) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v244
L2:
	;
	v100 = v14
	v102 = v15
	v106 = v2
	goto L25
L3:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v18 != int32(120) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v14&int32(3) == int32(0) {
		v44 = v14
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v79 = v77 - int32(2)
	v84 = F_palloc(m, int32(base.Ui32(v79)>>(uint(int32(1))%32))+int32(4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	v77 = v69 - v14
	goto L5
L7:
	;
	v48 = v44
	goto L16
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v77 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v33 = v14
	goto L12
L12:
	;
	v37 = v33 + int32(1)
	if v37&int32(3) == int32(0) {
		v44 = v37
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v69 = v37
	goto L6
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v42 != 0 {
		v33 = v37
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v57 = int32(-2139062144)
	if (int32(16843008)-v54|v54)&v57 == v57 {
		v48 = v48 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v63 = v48
	goto L19
L18:
	;
	goto L17
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 != 0 {
		v63 = v63 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v69 = v63
	goto L6
L21:
	;
	goto L20
L22:
	;
	return int32(0)
L23:
	;
	v92 = F_hex_decode_safe(m, v14+int32(2), v79, v84+int32(4), v13)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = base.I32_wrap_i64(v92)<<(uint(int32(2))%32) + int32(16)
	v244 = v84
	goto L1
L25:
	;
	if v102 != int32(92) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v239 = v100 + v236
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	v100 = v239
	v102 = v240
	v106 = v106 + int32(1)
	goto L25
L28:
	;
	v152 = v106 + int32(4)
	v153 = F_palloc(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L22
	} else {
		goto L45
	}
L29:
	;
	if v102 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v113&int32(252) == int32(48) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v236 = int32(1)
	goto L27
L33:
	;
	v132 = F_errsave_start(m, v13)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L22
	} else {
		goto L40
	}
L34:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	if v118&int32(248) != int32(48) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v113 != int32(92) {
		goto L33
	} else {
		goto L39
	}
L37:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+3)))
	if v123&int32(248) != int32(48) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v236 = int32(4)
	goto L27
L39:
	;
	v236 = int32(2)
	goto L27
L40:
	;
	if v132 == int32(0) {
		v244 = v2
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(498387)
	F_errmsg(m, int32(186323), v11+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v13, int32(491527), int32(342), int32(273970))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v244 = v2
	goto L1
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v152 << (uint(int32(2)) % 32)
	v160 = v153 + int32(4)
	v161 = v14
	goto L46
L46:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v168 != int32(92) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v218 = int32(0)
	v219 = F_errsave_start(m, v13)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L22
	} else {
		goto L59
	}
L48:
	;
	if v168 == int32(0) {
		v244 = v153
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v178&int32(252) == int32(48) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v168)
	v174 = int32(1)
	v160 = v160 + v174
	v161 = v161 + v174
	goto L46
L52:
	;
	goto L47
L53:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
	if v183&int32(248) != int32(48) {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v178 != int32(92) {
		goto L52
	} else {
		goto L58
	}
L56:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
	if v188&int32(248) != int32(48) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v202 = v188 + (v183<<(uint(int32(3))%32)&int32(56) | v178<<(uint(int32(6))%32)) - int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v202)
	v160 = v160 + int32(1)
	v161 = v161 + int32(4)
	goto L46
L58:
	;
	v210 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v210)
	v160 = v160 + int32(1)
	v161 = v161 + int32(2)
	goto L46
L59:
	;
	if v219 == int32(0) {
		v244 = v218
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L22
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(498387)
	F_errmsg(m, int32(186323), v11)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	F_errsave_finish(m, v13, int32(491527), int32(383), int32(273970))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	v244 = v218
	goto L1
}
func F_bytealike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = v7 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v20 = v18 & int32(1)
			if v20 != 0 {
				v21 = v12
			} else {
				v21 = v7 + int32(4)
			}
			if v18 == int32(1) {
				v24 = int32(4)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v26&int32(254) == int32(2) {
					v35 = v24
				} else {
					v35 = base.B2i32(v26 == int32(18)) << (uint(v24) % 32)
				}
				if v26 == int32(1) {
					v38 = v24
				} else {
					v38 = v35
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v20 != 0 {
					v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v14 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v14 + int32(4)
			}
			if v54 == int32(1) {
				v60 = int32(4)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v62&int32(254) == int32(2) {
					v71 = v60
				} else {
					v71 = base.B2i32(v62 == int32(18)) << (uint(v60) % 32)
				}
				if v62 == int32(1) {
					v74 = v60
				} else {
					v74 = v71
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v56 != 0 {
					v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_SB_MatchText(m, v21, v49, v57, v85, int32(0))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v87 == int32(1))
			}
		}
	}
}
