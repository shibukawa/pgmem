package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v7 != v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v4))) = float32(4)
	return v4
L2:
	;
	goto L3
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
	if base.Ui32(v16) < base.Ui32(v15) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v4))) = float32(3)
	return v4
L5:
	;
	goto L6
L6:
	;
	v21 = int32(4)
	v22 = v6 + v21
	v24 = v9 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)))
	if base.Ui32(v25) < base.Ui32(v26) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v25
	goto L9
L8:
	;
	v28 = v26
	goto L9
L9:
	;
	v29 = int32(0)
	v34 = int32(8)
	v35 = base.I32_div_s(v28, v34)
	if v34 <= v28 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v104 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	v104 = v101 + v97<<(uint(int32(3))%32)
	goto L10
L12:
	;
	v83 = v74
	goto L23
L13:
	;
	v41 = v29
	goto L16
L14:
	;
	v58 = v29
	goto L15
L15:
	;
	v65 = v28 - v35<<(uint(int32(3))%32)
	if v65 == int32(0) {
		v97 = v58
		v101 = v29
		goto L11
	} else {
		goto L22
	}
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v41))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v41))))
	if v47 != v49 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v58 = v35
	goto L15
L18:
	;
	v74 = int32(7)
	v75 = v41
	v77 = v47
	v78 = v49
	goto L12
L19:
	;
	goto L20
L20:
	;
	v53 = v41 + int32(1)
	if v53 != v35 {
		v41 = v53
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v58))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v58))))
	v74 = v65
	v75 = v58
	v77 = v71
	v78 = v69
	goto L12
L23:
	;
	if int32(base.Ui32(v77^v78)>>(uint(int32(8)-v83)%32)) != 0 {
		v83 = v83 - int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v97 = v75
	v101 = v83
	goto L11
L25:
	;
	goto L24
L26:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v4))) = float32(2)
	return v4
L27:
	;
	goto L28
L28:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v4))) = base.F32_div(float32(1), base.F32_convert_i32_u(v104))
	return v4
}
func F_inet_gist_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v334 int32
	_ = v334
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v413 int32
	_ = v413
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
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
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = v20 << (uint(int32(1)) % 32)
	v23 = F_palloc(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = F_palloc(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v31
	v36 = v19 + int32(4)
	v37 = int32(2)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+3)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = v20 - int32(1)
	if v42 < v37 {
		v209 = v39
		v210 = v40
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v36+v460<<(uint(int32(4))%32))))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+3)))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+2)))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+1)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v468 < int32(2) {
		v598 = v465
		v600 = v467
		v603 = v466
		goto L77
	} else {
		goto L78
	}
L5:
	;
	if v210 == int32(3) {
		goto L47
	} else {
		goto L48
	}
L6:
	;
	v46 = v38 + int32(4)
	v48 = v39
	v49 = v40
	v50 = v40
	v52 = v37
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v36+v52<<(uint(int32(4))%32))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+3)))
	if v48 < v70 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v156 == v154 {
		v209 = v153
		v210 = v156
		goto L5
	} else {
		goto L38
	}
L9:
	;
	v72 = v48
	goto L11
L10:
	;
	v72 = v70
	goto L11
L11:
	;
	if int32(0) < v72 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v76 = v67 + int32(4)
	v77 = int32(0)
	v82 = int32(8)
	v83 = base.I32_div_s(v72, v82)
	if v82 <= v72 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v153 = v72
	goto L14
L14:
	;
	if base.Ui32(v68) < base.Ui32(v50) {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v153 = v149 + v145<<(uint(int32(3))%32)
	goto L14
L16:
	;
	goto L15
L17:
	;
	v131 = v122
	goto L28
L18:
	;
	v89 = v77
	goto L21
L19:
	;
	v106 = v77
	goto L20
L20:
	;
	v113 = v72 - v83<<(uint(int32(3))%32)
	if v113 == int32(0) {
		v145 = v106
		v149 = v77
		goto L16
	} else {
		goto L27
	}
L21:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v89))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v89))))
	if v95 != v97 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v106 = v83
	goto L20
L23:
	;
	v122 = int32(7)
	v123 = v89
	v125 = v95
	v126 = v97
	goto L17
L24:
	;
	goto L25
L25:
	;
	v101 = v89 + int32(1)
	if v101 != v83 {
		v89 = v101
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v106))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v106))))
	v122 = v113
	v123 = v106
	v125 = v119
	v126 = v117
	goto L17
L28:
	;
	if int32(base.Ui32(v125^v126)>>(uint(int32(8)-v131)%32)) != 0 {
		v131 = v131 - int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v145 = v123
	v149 = v131
	goto L16
L30:
	;
	goto L29
L31:
	;
	v154 = v50
	goto L33
L32:
	;
	v154 = v68
	goto L33
L33:
	;
	if base.Ui32(v49) < base.Ui32(v68) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v156 = v49
	goto L36
L35:
	;
	v156 = v68
	goto L36
L36:
	;
	v158 = v52 + int32(1)
	if v158 <= v42 {
		v48 = v153
		v49 = v156
		v50 = v154
		v52 = v158
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	v161 = int32(1)
	v163 = v161
	v165 = v161
	goto L39
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v36+v165<<(uint(int32(4))%32))))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v184 != v154 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L4
L41:
	;
	v204 = v163 + int32(1)
	v206 = v204 & int32(_a_F_inet_gist_picksplit_0)
	if v206 <= v42 {
		v163 = v204
		v165 = v206
		goto L39
	} else {
		goto L45
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v187 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v186 + v187
	*(*uint16)(unsafe.Add(mBase, uint32(v23+v186<<(uint(v187)%32)))) = uint16(v163)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v195 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v194 + v195
	*(*uint16)(unsafe.Add(mBase, uint32(v27+v194<<(uint(v195)%32)))) = uint16(v163)
	goto L41
L45:
	;
	goto L40
L46:
	;
	v354 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v354
	v358 = int32(1)
	if v358 < v42 {
		goto L67
	} else {
		goto L68
	}
L47:
	;
	v229 = int32(128)
	goto L49
L48:
	;
	v229 = int32(32)
	goto L49
L49:
	;
	if v229 <= v209 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v232 = v209
	goto L51
L51:
	;
	v248 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v248
	v253 = base.I32_div_s(v232, int32(8))
	if v42 <= v248 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if v232 < v229 {
		goto L4
	} else {
		goto L66
	}
L53:
	;
	goto L52
L54:
	;
	v334 = v232 + int32(1)
	if v334 != v229 {
		v232 = v334
		goto L51
	} else {
		goto L65
	}
L55:
	;
	v260 = int32(1)
	v262 = v260
	v265 = v260
	goto L56
L56:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v36+v265<<(uint(int32(4))%32))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282+v253)+4)))
	if int32(base.Ui32(int32(128))>>(uint(v232&int32(7))%32))&v284 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v310 <= int32(0) {
		goto L54
	} else {
		goto L63
	}
L58:
	;
	v306 = v262 + int32(1)
	v308 = v306 & int32(_a_F_inet_gist_picksplit_0)
	if base.Ui32(v308) <= base.Ui32(v42) {
		v262 = v306
		v265 = v308
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v289 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v288 + v289
	*(*uint16)(unsafe.Add(mBase, uint32(v23+v288<<(uint(v289)%32)))) = uint16(v262)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v297 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v296 + v297
	*(*uint16)(unsafe.Add(mBase, uint32(v27+v296<<(uint(v297)%32)))) = uint16(v262)
	goto L58
L62:
	;
	goto L57
L63:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if int32(0) < v313 {
		goto L53
	} else {
		goto L64
	}
L64:
	;
	goto L54
L65:
	;
	goto L46
L66:
	;
	goto L46
L67:
	;
	v362 = base.I32_div_s(v42, int32(2))
	v363 = v358
	goto L70
L68:
	;
	v393 = v358
	goto L69
L69:
	;
	if v42 < v393&int32(_a_F_inet_gist_picksplit_0) {
		goto L4
	} else {
		goto L73
	}
L70:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v381 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v380 + v381
	*(*uint16)(unsafe.Add(mBase, uint32(v23+v380<<(uint(v381)%32)))) = uint16(v363)
	v389 = v363 + v381
	if v389&int32(_a_F_inet_gist_picksplit_0) <= v362 {
		v363 = v389
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v393 = v389
	goto L69
L72:
	;
	goto L71
L73:
	;
	v413 = v393
	goto L74
L74:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v431 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v430 + v431
	*(*uint16)(unsafe.Add(mBase, uint32(v27+v430<<(uint(v431)%32)))) = uint16(v413)
	v439 = v413 + v431
	if base.Ui32(v439&int32(_a_F_inet_gist_picksplit_0)) <= base.Ui32(v42) {
		v413 = v439
		goto L74
	} else {
		goto L76
	}
L75:
	;
	goto L4
L76:
	;
	goto L75
L77:
	;
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v36+v615<<(uint(int32(4))%32))))
	v621 = F_palloc0(m, int32(20))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L114
	}
L78:
	;
	v472 = v464 + int32(4)
	v474 = v465
	v475 = v467
	v476 = v467
	v477 = int32(1)
	v479 = v466
	goto L79
L79:
	;
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v477<<(uint(int32(1))%32)))))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v36+v494<<(uint(int32(4))%32))))
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+2)))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+3)))
	if v474 < v503 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v588 == v590 {
		v598 = v586
		v600 = v590
		v603 = v587
		goto L77
	} else {
		goto L113
	}
L81:
	;
	v505 = v474
	goto L83
L82:
	;
	v505 = v503
	goto L83
L83:
	;
	if int32(0) < v505 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v509 = v498 + int32(4)
	v510 = int32(0)
	v515 = int32(8)
	v516 = base.I32_div_s(v505, v515)
	if v515 <= v505 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v586 = v505
	goto L86
L86:
	;
	if base.Ui32(v479) < base.Ui32(v499) {
		goto L103
	} else {
		goto L104
	}
L87:
	;
	v586 = v582 + v578<<(uint(int32(3))%32)
	goto L86
L88:
	;
	goto L87
L89:
	;
	v564 = v555
	goto L100
L90:
	;
	v522 = v510
	goto L93
L91:
	;
	v539 = v510
	goto L92
L92:
	;
	v546 = v505 - v516<<(uint(int32(3))%32)
	if v546 == int32(0) {
		v578 = v539
		v582 = v510
		goto L88
	} else {
		goto L99
	}
L93:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v522))))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v522))))
	if v528 != v530 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v539 = v516
	goto L92
L95:
	;
	v555 = int32(7)
	v556 = v522
	v558 = v528
	v559 = v530
	goto L89
L96:
	;
	goto L97
L97:
	;
	v534 = v522 + int32(1)
	if v534 != v516 {
		v522 = v534
		goto L93
	} else {
		goto L98
	}
L98:
	;
	goto L94
L99:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v539))))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v539))))
	v555 = v546
	v556 = v539
	v558 = v552
	v559 = v550
	goto L89
L100:
	;
	if int32(base.Ui32(v558^v559)>>(uint(int32(8)-v564)%32)) != 0 {
		v564 = v564 - int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v578 = v556
	v582 = v564
	goto L88
L102:
	;
	goto L101
L103:
	;
	v587 = v479
	goto L105
L104:
	;
	v587 = v499
	goto L105
L105:
	;
	if base.Ui32(v501) < base.Ui32(v475) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v588 = v475
	goto L108
L107:
	;
	v588 = v501
	goto L108
L108:
	;
	if base.Ui32(v476) < base.Ui32(v501) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v590 = v476
	goto L111
L110:
	;
	v590 = v501
	goto L111
L111:
	;
	v592 = v477 + int32(1)
	if v592 != v468 {
		v474 = v586
		v475 = v588
		v476 = v590
		v477 = v592
		v479 = v587
		goto L79
	} else {
		goto L112
	}
L112:
	;
	goto L80
L113:
	;
	v595 = int32(0)
	v598 = v595
	v600 = v595
	v603 = v595
	goto L77
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+3)) = uint8(v598)
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+2)) = uint8(v603)
	*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)) = uint8(v600)
	if v598 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v643 = base.I32_div_s(v598, int32(8))
	v646 = v598 - v643<<(uint(int32(3))%32)
	if v646 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v631 = base.I32_div_s(v598+int32(7), int32(8))
	if v631 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v634 = int32(4)
	base.MemoryCopy(m, v621+v634, v619+v634, v631)
	goto L115
L118:
	;
	v649 = v643 + v621 + int32(4)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	v653 = v650 & (int32(-256) >> (uint(v646) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v649))) = uint8(v653)
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)))
	v657 = v655
	goto L120
L119:
	;
	v657 = v600
	goto L120
L120:
	;
	if v657&int32(255) == int32(3) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v662 = int32(41)
	goto L123
L122:
	;
	v662 = int32(17)
	goto L123
L123:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v621))) = uint8(v662)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v621
	v665 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v36+v665<<(uint(int32(4))%32))))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+3)))
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+2)))
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669)+1)))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v673 < int32(2) {
		v803 = v670
		v805 = v672
		v808 = v671
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v820 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v36+v820<<(uint(int32(4))%32))))
	v826 = F_palloc0(m, int32(20))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L161
	}
L125:
	;
	v677 = v669 + int32(4)
	v679 = v670
	v680 = v672
	v681 = v672
	v682 = int32(1)
	v684 = v671
	goto L126
L126:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v682<<(uint(int32(1))%32)))))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v36+v699<<(uint(int32(4))%32))))
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+2)))
	v706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+1)))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+3)))
	if v679 < v708 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v793 == v795 {
		v803 = v791
		v805 = v795
		v808 = v792
		goto L124
	} else {
		goto L160
	}
L128:
	;
	v710 = v679
	goto L130
L129:
	;
	v710 = v708
	goto L130
L130:
	;
	if int32(0) < v710 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v714 = v703 + int32(4)
	v715 = int32(0)
	v720 = int32(8)
	v721 = base.I32_div_s(v710, v720)
	if v720 <= v710 {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v791 = v710
	goto L133
L133:
	;
	if base.Ui32(v684) < base.Ui32(v704) {
		goto L150
	} else {
		goto L151
	}
L134:
	;
	v791 = v787 + v783<<(uint(int32(3))%32)
	goto L133
L135:
	;
	goto L134
L136:
	;
	v769 = v760
	goto L147
L137:
	;
	v727 = v715
	goto L140
L138:
	;
	v744 = v715
	goto L139
L139:
	;
	v751 = v710 - v721<<(uint(int32(3))%32)
	if v751 == int32(0) {
		v783 = v744
		v787 = v715
		goto L135
	} else {
		goto L146
	}
L140:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v727))))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+v727))))
	if v733 != v735 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v744 = v721
	goto L139
L142:
	;
	v760 = int32(7)
	v761 = v727
	v763 = v733
	v764 = v735
	goto L136
L143:
	;
	goto L144
L144:
	;
	v739 = v727 + int32(1)
	if v739 != v721 {
		v727 = v739
		goto L140
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+v744))))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v744))))
	v760 = v751
	v761 = v744
	v763 = v757
	v764 = v755
	goto L136
L147:
	;
	if int32(base.Ui32(v763^v764)>>(uint(int32(8)-v769)%32)) != 0 {
		v769 = v769 - int32(1)
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v783 = v761
	v787 = v769
	goto L135
L149:
	;
	goto L148
L150:
	;
	v792 = v684
	goto L152
L151:
	;
	v792 = v704
	goto L152
L152:
	;
	if base.Ui32(v706) < base.Ui32(v680) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v793 = v680
	goto L155
L154:
	;
	v793 = v706
	goto L155
L155:
	;
	if base.Ui32(v681) < base.Ui32(v706) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v795 = v681
	goto L158
L157:
	;
	v795 = v706
	goto L158
L158:
	;
	v797 = v682 + int32(1)
	if v797 != v673 {
		v679 = v791
		v680 = v793
		v681 = v795
		v682 = v797
		v684 = v792
		goto L126
	} else {
		goto L159
	}
L159:
	;
	goto L127
L160:
	;
	v800 = int32(0)
	v803 = v800
	v805 = v800
	v808 = v800
	goto L124
L161:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v826)+3)) = uint8(v803)
	*(*uint8)(unsafe.Add(mBase, uint32(v826)+2)) = uint8(v808)
	*(*uint8)(unsafe.Add(mBase, uint32(v826)+1)) = uint8(v805)
	if v803 <= int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v848 = base.I32_div_s(v803, int32(8))
	v851 = v803 - v848<<(uint(int32(3))%32)
	if v851 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v836 = base.I32_div_s(v803+int32(7), int32(8))
	if v836 == int32(0) {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v839 = int32(4)
	base.MemoryCopy(m, v826+v839, v824+v839, v836)
	goto L162
L165:
	;
	v854 = v848 + v826 + int32(4)
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854))))
	v858 = v855 & (int32(-256) >> (uint(v851) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v854))) = uint8(v858)
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826)+1)))
	v862 = v860
	goto L167
L166:
	;
	v862 = v805
	goto L167
L167:
	;
	if v862&int32(255) == int32(3) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v867 = int32(41)
	goto L170
L169:
	;
	v867 = int32(17)
	goto L170
L170:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v826))) = uint8(v867)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v826
	return v18
}
func F_inet_inclusion_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v8 = int32(1)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10&v8 != 0 {
		v13 = v8
	} else {
		v13 = int32(4)
	}
	v14 = l0 + v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v16 = int32(1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18&v16 != 0 {
		v21 = v16
	} else {
		v21 = int32(4)
	}
	v22 = l1 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v15 == v23 {
		v25 = int32(2)
		v26 = v14 + v25
		v28 = v22 + v25
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
		if base.Ui32(v29) < base.Ui32(v30) {
			v32 = v29
		} else {
			v32 = v30
		}
		v37 = base.I32_div_s(v32, int32(8))
		v38 = F_memcmp(m, v26, v28, v37)
		mBase = m.M
		if v38 != 0 {
			v124 = v38
			v134 = v124
		} else {
			v39 = int32(0)
			v42 = v32 - v37<<(uint(int32(3))%32)
			if v42 <= v39 {
				v124 = v39
				v134 = v124
			} else {
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v37))))
				v47 = int32(128)
				v48 = v46 & v47
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v37))))
				if v48 != v50&v47 {
					v125 = v48
					if v125 != 0 {
						v128 = int32(1)
					} else {
						v128 = int32(-1)
					}
					v134 = v128
				} else {
					if v42 == int32(1) {
						v124 = v39
						v134 = v124
					} else {
						v56 = int32(1)
						v58 = int32(128)
						v59 = v46 << (uint(v56) % 32) & v58
						if v59 != v50<<(uint(v56)%32)&v58 {
							v125 = v59
							if v125 != 0 {
								v128 = int32(1)
							} else {
								v128 = int32(-1)
							}
							v134 = v128
						} else {
							if v42 < int32(3) {
								v124 = v39
								v134 = v124
							} else {
								v67 = int32(2)
								v69 = int32(128)
								v70 = v46 << (uint(v67) % 32) & v69
								if v70 != v50<<(uint(v67)%32)&v69 {
									v125 = v70
									if v125 != 0 {
										v128 = int32(1)
									} else {
										v128 = int32(-1)
									}
									v134 = v128
								} else {
									if v42 == int32(3) {
										v124 = v39
										v134 = v124
									} else {
										v78 = int32(3)
										v80 = int32(128)
										v81 = v46 << (uint(v78) % 32) & v80
										if v81 != v50<<(uint(v78)%32)&v80 {
											v125 = v81
											if v125 != 0 {
												v128 = int32(1)
											} else {
												v128 = int32(-1)
											}
											v134 = v128
										} else {
											if v42 < int32(5) {
												v124 = v39
												v134 = v124
											} else {
												v89 = int32(4)
												v91 = int32(128)
												v92 = v46 << (uint(v89) % 32) & v91
												if v92 != v50<<(uint(v89)%32)&v91 {
													v125 = v92
													if v125 != 0 {
														v128 = int32(1)
													} else {
														v128 = int32(-1)
													}
													v134 = v128
												} else {
													if v42 == int32(5) {
														v124 = v39
														v134 = v124
													} else {
														v100 = int32(5)
														v102 = int32(128)
														v103 = v46 << (uint(v100) % 32) & v102
														if v103 != v50<<(uint(v100)%32)&v102 {
															v125 = v103
															if v125 != 0 {
																v128 = int32(1)
															} else {
																v128 = int32(-1)
															}
															v134 = v128
														} else {
															if v42 < int32(7) {
																v124 = v39
																v134 = v124
															} else {
																v111 = int32(6)
																v113 = int32(128)
																v114 = v46 << (uint(v111) % 32) & v113
																if v114 != v50<<(uint(v111)%32)&v113 {
																	v125 = v114
																	if v125 != 0 {
																		v128 = int32(1)
																	} else {
																		v128 = int32(-1)
																	}
																	v134 = v128
																} else {
																	v124 = v39
																	v134 = v124
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
						}
					}
				}
			}
		}
		if v134 != 0 {
			v176 = v134
			return v176
		} else {
			v136 = int32(1)
			v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v138&v136 != 0 {
				v141 = v136
			} else {
				v141 = int32(4)
			}
			v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v141)+1)))
			v144 = int32(1)
			v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v146&v144 != 0 {
				v149 = v144
			} else {
				v149 = int32(4)
			}
			v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v149)+1)))
			v152 = v143 - v151
			v153 = int32(0)
			if base.B2i32(v153 < v152)&base.B2i32(v153 <= l2)|base.B2i32(v143 == v151)&base.B2i32(base.Ui32(l2+int32(1)) <= base.Ui32(int32(2))) != 0 {
				v176 = int32(0)
				return v176
			} else {
				v165 = int32(0)
				if v165 <= v152 {
					v168 = l2
				} else {
					v168 = v165
				}
				if l2 <= int32(0) {
					v171 = v168
				} else {
					v171 = l2
				}
				return v171
			}
		}
	} else {
		v176 = v15 - v23
		return v176
	}
}
func F_inet_spg_config(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500938)
	return int32(0)
}
