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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	v104 = v99 + v97<<(uint(int32(3))%32)
	goto L10
L12:
	;
	v85 = v76
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
		v99 = v29
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
	v75 = v41
	v76 = int32(7)
	v78 = v47
	v79 = v49
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
	v75 = v58
	v76 = v65
	v78 = v71
	v79 = v69
	goto L12
L23:
	;
	if int32(base.Ui32(v78^v79)>>(uint(int32(8)-v85)%32)) != 0 {
		v85 = v85 - int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v97 = v75
	v99 = v85
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v441 int32
	_ = v441
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v865 int32
	_ = v865
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v23 = v21 << (uint(int32(1)) % 32)
	v24 = F_palloc(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = F_palloc(m, v23)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v24
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v32
	v37 = v20 + int32(4)
	v38 = int32(2)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+3)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = v21 - int32(1)
	if v43 < v38 {
		v213 = v41
		v215 = v40
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v37+v490<<(uint(int32(4))%32))))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+3)))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+2)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v498 < int32(2) {
		v629 = v495
		v631 = v497
		v636 = v496
		goto L77
	} else {
		goto L78
	}
L5:
	;
	if v213 == int32(3) {
		goto L48
	} else {
		goto L49
	}
L6:
	;
	v47 = v39 + int32(4)
	v49 = v41
	v50 = v41
	v52 = v40
	v55 = v38
	goto L7
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v37+v55<<(uint(int32(4))%32))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+3)))
	if v52 < v73 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v157 == v158 {
		v213 = v158
		v215 = v156
		goto L5
	} else {
		goto L38
	}
L9:
	;
	v75 = v52
	goto L11
L10:
	;
	v75 = v73
	goto L11
L11:
	;
	if int32(0) < v75 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v79 = v69 + int32(4)
	v80 = int32(0)
	v85 = int32(8)
	v86 = base.I32_div_s(v75, v85)
	if v85 <= v75 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v156 = v75
	goto L14
L14:
	;
	if base.Ui32(v70) < base.Ui32(v49) {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v156 = v150 + v148<<(uint(int32(3))%32)
	goto L14
L16:
	;
	goto L15
L17:
	;
	v136 = v127
	goto L28
L18:
	;
	v92 = v80
	goto L21
L19:
	;
	v109 = v80
	goto L20
L20:
	;
	v116 = v75 - v86<<(uint(int32(3))%32)
	if v116 == int32(0) {
		v148 = v109
		v150 = v80
		goto L16
	} else {
		goto L27
	}
L21:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v92))))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v92))))
	if v98 != v100 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v109 = v86
	goto L20
L23:
	;
	v126 = v92
	v127 = int32(7)
	v129 = v98
	v130 = v100
	goto L17
L24:
	;
	goto L25
L25:
	;
	v104 = v92 + int32(1)
	if v104 != v86 {
		v92 = v104
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v109))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v109))))
	v126 = v109
	v127 = v116
	v129 = v122
	v130 = v120
	goto L17
L28:
	;
	if int32(base.Ui32(v129^v130)>>(uint(int32(8)-v136)%32)) != 0 {
		v136 = v136 - int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v148 = v126
	v150 = v136
	goto L16
L30:
	;
	goto L29
L31:
	;
	v157 = v49
	goto L33
L32:
	;
	v157 = v70
	goto L33
L33:
	;
	if base.Ui32(v50) < base.Ui32(v70) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v158 = v50
	goto L36
L35:
	;
	v158 = v70
	goto L36
L36:
	;
	v160 = v55 + int32(1)
	if v160 <= v43 {
		v49 = v157
		v50 = v158
		v52 = v156
		v55 = v160
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	v163 = int32(1)
	v165 = v163
	v167 = v163
	goto L39
L39:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v37+v167<<(uint(int32(4))%32))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	if v187 != v157 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L4
L41:
	;
	v207 = v165 + int32(1)
	v209 = v207 & int32(_a_F_inet_gist_picksplit_0)
	if v209 <= v43 {
		v165 = v207
		v167 = v209
		goto L39
	} else {
		goto L45
	}
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v190 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v189 + v190
	*(*uint16)(unsafe.Add(mBase, uint32(v24+v189<<(uint(v190)%32)))) = uint16(v165)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v198 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v197 + v198
	*(*uint16)(unsafe.Add(mBase, uint32(v28+v197<<(uint(v198)%32)))) = uint16(v165)
	goto L41
L45:
	;
	goto L40
L46:
	;
	v380 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v380
	v385 = base.I32_div_s(v43, int32(2))
	v386 = int32(1)
	if v386 < v43 {
		goto L67
	} else {
		goto L68
	}
L47:
	;
	if v347 < v233 {
		goto L4
	} else {
		goto L66
	}
L48:
	;
	v233 = int32(128)
	goto L50
L49:
	;
	v233 = int32(32)
	goto L50
L50:
	;
	if v233 <= v215 {
		v347 = v215
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v239 = v215
	goto L52
L52:
	;
	v253 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v253
	v258 = base.I32_div_s(v239, int32(8))
	if v43 <= v253 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L46
L54:
	;
	v341 = v239 + int32(1)
	if v341 != v233 {
		v239 = v341
		goto L52
	} else {
		goto L65
	}
L55:
	;
	v265 = int32(1)
	v267 = v265
	v268 = v265
	goto L56
L56:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v37+v268<<(uint(int32(4))%32))))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288+v258)+4)))
	if int32(base.Ui32(int32(128))>>(uint(v239&int32(7))%32))&v290 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v316 <= int32(0) {
		goto L54
	} else {
		goto L63
	}
L58:
	;
	v312 = v267 + int32(1)
	v314 = v312 & int32(_a_F_inet_gist_picksplit_0)
	if base.Ui32(v314) <= base.Ui32(v43) {
		v267 = v312
		v268 = v314
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v295 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v294 + v295
	*(*uint16)(unsafe.Add(mBase, uint32(v24+v294<<(uint(v295)%32)))) = uint16(v267)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v303 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v302 + v303
	*(*uint16)(unsafe.Add(mBase, uint32(v28+v302<<(uint(v303)%32)))) = uint16(v267)
	goto L58
L62:
	;
	goto L57
L63:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if int32(0) < v319 {
		v347 = v239
		goto L47
	} else {
		goto L64
	}
L64:
	;
	goto L54
L65:
	;
	goto L53
L66:
	;
	goto L46
L67:
	;
	v389 = v386
	goto L70
L68:
	;
	v420 = v386
	goto L69
L69:
	;
	if v43 < v420&int32(_a_F_inet_gist_picksplit_0) {
		goto L4
	} else {
		goto L73
	}
L70:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v408 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v407 + v408
	*(*uint16)(unsafe.Add(mBase, uint32(v24+v407<<(uint(v408)%32)))) = uint16(v389)
	v416 = v389 + v408
	if v416&int32(_a_F_inet_gist_picksplit_0) <= v385 {
		v389 = v416
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v420 = v416
	goto L69
L72:
	;
	goto L71
L73:
	;
	v441 = v420
	goto L74
L74:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v460 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v459 + v460
	*(*uint16)(unsafe.Add(mBase, uint32(v28+v459<<(uint(v460)%32)))) = uint16(v441)
	v468 = v441 + v460
	if base.Ui32(v468&int32(_a_F_inet_gist_picksplit_0)) <= base.Ui32(v43) {
		v441 = v468
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
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v37+v647<<(uint(int32(4))%32))))
	v653 = F_palloc0(m, int32(20))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L114
	}
L78:
	;
	v502 = v494 + int32(4)
	v504 = v495
	v506 = v497
	v508 = v497
	v511 = v496
	v513 = int32(1)
	goto L79
L79:
	;
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+v513<<(uint(int32(1))%32)))))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v37+v525<<(uint(int32(4))%32))))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+2)))
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+1)))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+3)))
	if v504 < v535 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v621 == v620 {
		v629 = v618
		v631 = v621
		v636 = v619
		goto L77
	} else {
		goto L113
	}
L81:
	;
	v537 = v504
	goto L83
L82:
	;
	v537 = v535
	goto L83
L83:
	;
	if int32(0) < v537 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v541 = v529 + int32(4)
	v542 = int32(0)
	v547 = int32(8)
	v548 = base.I32_div_s(v537, v547)
	if v547 <= v537 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v618 = v537
	goto L86
L86:
	;
	if base.Ui32(v511) < base.Ui32(v530) {
		goto L103
	} else {
		goto L104
	}
L87:
	;
	v618 = v612 + v610<<(uint(int32(3))%32)
	goto L86
L88:
	;
	goto L87
L89:
	;
	v598 = v589
	goto L100
L90:
	;
	v554 = v542
	goto L93
L91:
	;
	v571 = v542
	goto L92
L92:
	;
	v578 = v537 - v548<<(uint(int32(3))%32)
	if v578 == int32(0) {
		v610 = v571
		v612 = v542
		goto L88
	} else {
		goto L99
	}
L93:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v554))))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+v554))))
	if v560 != v562 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v571 = v548
	goto L92
L95:
	;
	v588 = v554
	v589 = int32(7)
	v591 = v560
	v592 = v562
	goto L89
L96:
	;
	goto L97
L97:
	;
	v566 = v554 + int32(1)
	if v566 != v548 {
		v554 = v566
		goto L93
	} else {
		goto L98
	}
L98:
	;
	goto L94
L99:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+v571))))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v571))))
	v588 = v571
	v589 = v578
	v591 = v584
	v592 = v582
	goto L89
L100:
	;
	if int32(base.Ui32(v591^v592)>>(uint(int32(8)-v598)%32)) != 0 {
		v598 = v598 - int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v610 = v588
	v612 = v598
	goto L88
L102:
	;
	goto L101
L103:
	;
	v619 = v511
	goto L105
L104:
	;
	v619 = v530
	goto L105
L105:
	;
	if base.Ui32(v532) < base.Ui32(v508) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v620 = v508
	goto L108
L107:
	;
	v620 = v532
	goto L108
L108:
	;
	if base.Ui32(v506) < base.Ui32(v532) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v621 = v506
	goto L111
L110:
	;
	v621 = v532
	goto L111
L111:
	;
	v623 = v513 + int32(1)
	if v623 != v498 {
		v504 = v618
		v506 = v621
		v508 = v620
		v511 = v619
		v513 = v623
		goto L79
	} else {
		goto L112
	}
L112:
	;
	goto L80
L113:
	;
	v626 = int32(0)
	v629 = v626
	v631 = v626
	v636 = v626
	goto L77
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v653)+3)) = uint8(v629)
	*(*uint8)(unsafe.Add(mBase, uint32(v653)+2)) = uint8(v636)
	*(*uint8)(unsafe.Add(mBase, uint32(v653)+1)) = uint8(v631)
	if int32(0) < v629 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v660 = int32(4)
	v667 = base.I32_div_s(v629+int32(7), int32(8))
	if v667 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	goto L117
L117:
	;
	v673 = base.I32_div_s(v629, int32(8))
	v676 = v629 - v673<<(uint(int32(3))%32)
	if v676 != 0 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	goto L117
L119:
	;
	v668 = F__emscripten_memcpy_bulkmem(m, v653+v660, v651+v660, v667)
	mBase = m.M
	goto L121
L120:
	;
	goto L121
L121:
	;
	goto L118
L122:
	;
	v679 = v653 + v673 + int32(4)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	v683 = v680 & (int32(-256) >> (uint(v676) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v679))) = uint8(v683)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+1)))
	v687 = v685
	goto L124
L123:
	;
	v687 = v631
	goto L124
L124:
	;
	if v687&int32(255) == int32(3) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v692 = int32(41)
	goto L127
L126:
	;
	v692 = int32(17)
	goto L127
L127:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v653))) = uint8(v692)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v653
	v695 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v37+v695<<(uint(int32(4))%32))))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+3)))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+2)))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+1)))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v703 < int32(2) {
		v834 = v700
		v836 = v702
		v841 = v701
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v37+v852<<(uint(int32(4))%32))))
	v858 = F_palloc0(m, int32(20))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L165
	}
L129:
	;
	v707 = v699 + int32(4)
	v709 = v700
	v711 = v702
	v713 = v702
	v716 = v701
	v717 = int32(1)
	goto L130
L130:
	;
	v730 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+v717<<(uint(int32(1))%32)))))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v37+v730<<(uint(int32(4))%32))))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+2)))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+1)))
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+3)))
	if v709 < v740 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v826 == v825 {
		v834 = v823
		v836 = v826
		v841 = v824
		goto L128
	} else {
		goto L164
	}
L132:
	;
	v742 = v709
	goto L134
L133:
	;
	v742 = v740
	goto L134
L134:
	;
	if int32(0) < v742 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v746 = v734 + int32(4)
	v747 = int32(0)
	v752 = int32(8)
	v753 = base.I32_div_s(v742, v752)
	if v752 <= v742 {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v823 = v742
	goto L137
L137:
	;
	if base.Ui32(v716) < base.Ui32(v735) {
		goto L154
	} else {
		goto L155
	}
L138:
	;
	v823 = v817 + v815<<(uint(int32(3))%32)
	goto L137
L139:
	;
	goto L138
L140:
	;
	v803 = v794
	goto L151
L141:
	;
	v759 = v747
	goto L144
L142:
	;
	v776 = v747
	goto L143
L143:
	;
	v783 = v742 - v753<<(uint(int32(3))%32)
	if v783 == int32(0) {
		v815 = v776
		v817 = v747
		goto L139
	} else {
		goto L150
	}
L144:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707+v759))))
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v759))))
	if v765 != v767 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v776 = v753
	goto L143
L146:
	;
	v793 = v759
	v794 = int32(7)
	v796 = v765
	v797 = v767
	goto L140
L147:
	;
	goto L148
L148:
	;
	v771 = v759 + int32(1)
	if v771 != v753 {
		v759 = v771
		goto L144
	} else {
		goto L149
	}
L149:
	;
	goto L145
L150:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v776))))
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707+v776))))
	v793 = v776
	v794 = v783
	v796 = v789
	v797 = v787
	goto L140
L151:
	;
	if int32(base.Ui32(v796^v797)>>(uint(int32(8)-v803)%32)) != 0 {
		v803 = v803 - int32(1)
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v815 = v793
	v817 = v803
	goto L139
L153:
	;
	goto L152
L154:
	;
	v824 = v716
	goto L156
L155:
	;
	v824 = v735
	goto L156
L156:
	;
	if base.Ui32(v737) < base.Ui32(v713) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v825 = v713
	goto L159
L158:
	;
	v825 = v737
	goto L159
L159:
	;
	if base.Ui32(v711) < base.Ui32(v737) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v826 = v711
	goto L162
L161:
	;
	v826 = v737
	goto L162
L162:
	;
	v828 = v717 + int32(1)
	if v828 != v703 {
		v709 = v823
		v711 = v826
		v713 = v825
		v716 = v824
		v717 = v828
		goto L130
	} else {
		goto L163
	}
L163:
	;
	goto L131
L164:
	;
	v831 = int32(0)
	v834 = v831
	v836 = v831
	v841 = v831
	goto L128
L165:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+3)) = uint8(v834)
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+2)) = uint8(v841)
	*(*uint8)(unsafe.Add(mBase, uint32(v858)+1)) = uint8(v836)
	if int32(0) < v834 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v865 = int32(4)
	v872 = base.I32_div_s(v834+int32(7), int32(8))
	if v872 != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	v878 = base.I32_div_s(v834, int32(8))
	v881 = v834 - v878<<(uint(int32(3))%32)
	if v881 != 0 {
		goto L173
	} else {
		goto L174
	}
L169:
	;
	goto L168
L170:
	;
	v873 = F__emscripten_memcpy_bulkmem(m, v858+v865, v856+v865, v872)
	mBase = m.M
	goto L172
L171:
	;
	goto L172
L172:
	;
	goto L169
L173:
	;
	v884 = v858 + v878 + int32(4)
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884))))
	v888 = v885 & (int32(-256) >> (uint(v881) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v884))) = uint8(v888)
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v858)+1)))
	v892 = v890
	goto L175
L174:
	;
	v892 = v836
	goto L175
L175:
	;
	if v892&int32(255) == int32(3) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v897 = int32(41)
	goto L178
L177:
	;
	v897 = int32(17)
	goto L178
L178:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v858))) = uint8(v897)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v858
	return v19
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
	var v135 int32
	_ = v135
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
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
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
			v175 = v134
			return v175
		} else {
			v135 = int32(0)
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
			if base.B2i32(v153 < v152)&base.B2i32(v153 <= l2) != 0 {
				v175 = v135
				return v175
			} else {
				if base.B2i32(v143 == v151)&base.B2i32(base.Ui32(l2+int32(1)) <= base.Ui32(int32(2))) != 0 {
					v175 = v135
					return v175
				} else {
					v164 = int32(0)
					if v164 <= v152 {
						v167 = l2
					} else {
						v167 = v164
					}
					if l2 <= int32(0) {
						v170 = v167
					} else {
						v170 = l2
					}
					return v170
				}
			}
		}
	} else {
		v175 = v15 - v23
		return v175
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
