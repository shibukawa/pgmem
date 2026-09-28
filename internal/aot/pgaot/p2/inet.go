package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 float32
	_ = v110
	v4 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v9 != v12 {
		v110 = float32(4)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*float32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v4)))) = v110
	return v4
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
	if base.Ui32(v16) < base.Ui32(v15) {
		v110 = float32(3)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(4)
	v20 = v8 + v19
	v22 = v11 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+3)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)))
	if base.Ui32(v23) < base.Ui32(v24) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v23
	goto L6
L5:
	;
	v26 = v24
	goto L6
L6:
	;
	v27 = int32(0)
	v32 = int32(8)
	v33 = base.I32_div_s(v26, v32)
	if v32 <= v26 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	if v102 <= int32(0) {
		v110 = float32(2)
		goto L1
	} else {
		goto L23
	}
L8:
	;
	v102 = v99 + v95<<(uint(int32(3))%32)
	goto L7
L9:
	;
	v81 = v72
	goto L20
L10:
	;
	v39 = v27
	goto L13
L11:
	;
	v56 = v27
	goto L12
L12:
	;
	v63 = v26 - v33<<(uint(int32(3))%32)
	if v63 == int32(0) {
		v95 = v56
		v99 = v27
		goto L8
	} else {
		goto L19
	}
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v39))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v39))))
	if v45 != v47 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v56 = v33
	goto L12
L15:
	;
	v72 = int32(7)
	v73 = v39
	v75 = v45
	v76 = v47
	goto L9
L16:
	;
	goto L17
L17:
	;
	v51 = v39 + int32(1)
	if v51 != v33 {
		v39 = v51
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v56))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v56))))
	v72 = v63
	v73 = v56
	v75 = v69
	v76 = v67
	goto L9
L20:
	;
	if int32(base.Ui32(v75^v76)>>(uint(int32(8)-v81)%32)) != 0 {
		v81 = v81 - int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v95 = v73
	v99 = v81
	goto L8
L22:
	;
	goto L21
L23:
	;
	v110 = base.F32_div(float32(1), base.F32_convert_i32_u(v102))
	goto L1
}
func F_inet_gist_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v349 int32
	_ = v349
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v434 int32
	_ = v434
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v644 int32
	_ = v644
	var v648 int64
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
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
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v855 int32
	_ = v855
	var v859 int64
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = v22 << (uint(int32(1)) % 32)
	v25 = F_palloc(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v29 = base.I32_wrap_i64(v20)
	v30 = F_palloc(m, v24)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v25
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v34
	v39 = v21 + int32(8)
	v40 = int32(2)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+3)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = v22 - int32(1)
	if v45 < v40 {
		v216 = v42
		v217 = v43
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v39+v485*int32(24))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+3)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+2)))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v493 < int32(2) {
		v625 = v490
		v627 = v492
		v630 = v491
		goto L77
	} else {
		goto L78
	}
L5:
	;
	if v217 == int32(3) {
		goto L47
	} else {
		goto L48
	}
L6:
	;
	v49 = v41 + int32(4)
	v51 = v42
	v52 = v43
	v53 = v43
	v55 = v40
	goto L7
L7:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v39+v55*int32(24))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+3)))
	if v51 < v75 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v161 == v159 {
		v216 = v158
		v217 = v161
		goto L5
	} else {
		goto L38
	}
L9:
	;
	v77 = v51
	goto L11
L10:
	;
	v77 = v75
	goto L11
L11:
	;
	if int32(0) < v77 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v81 = v72 + int32(4)
	v82 = int32(0)
	v87 = int32(8)
	v88 = base.I32_div_s(v77, v87)
	if v87 <= v77 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v158 = v77
	goto L14
L14:
	;
	if base.Ui32(v73) < base.Ui32(v53) {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v158 = v154 + v150<<(uint(int32(3))%32)
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
	v94 = v82
	goto L21
L19:
	;
	v111 = v82
	goto L20
L20:
	;
	v118 = v77 - v88<<(uint(int32(3))%32)
	if v118 == int32(0) {
		v150 = v111
		v154 = v82
		goto L16
	} else {
		goto L27
	}
L21:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v94))))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v94))))
	if v100 != v102 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v111 = v88
	goto L20
L23:
	;
	v127 = int32(7)
	v128 = v94
	v130 = v100
	v131 = v102
	goto L17
L24:
	;
	goto L25
L25:
	;
	v106 = v94 + int32(1)
	if v106 != v88 {
		v94 = v106
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v111))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v111))))
	v127 = v118
	v128 = v111
	v130 = v124
	v131 = v122
	goto L17
L28:
	;
	if int32(base.Ui32(v130^v131)>>(uint(int32(8)-v136)%32)) != 0 {
		v136 = v136 - int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v150 = v128
	v154 = v136
	goto L16
L30:
	;
	goto L29
L31:
	;
	v159 = v53
	goto L33
L32:
	;
	v159 = v73
	goto L33
L33:
	;
	if base.Ui32(v52) < base.Ui32(v73) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v161 = v52
	goto L36
L35:
	;
	v161 = v73
	goto L36
L36:
	;
	v163 = v55 + int32(1)
	if v163 <= v45 {
		v51 = v158
		v52 = v161
		v53 = v159
		v55 = v163
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	v166 = int32(1)
	v168 = v166
	v170 = v166
	goto L39
L39:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v39+v170*int32(24))))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
	if v191 != v159 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L4
L41:
	;
	v211 = v168 + int32(1)
	v213 = v211 & int32(_a_F_inet_gist_picksplit_0)
	if v213 <= v45 {
		v168 = v211
		v170 = v213
		goto L39
	} else {
		goto L45
	}
L42:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v194 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v193 + v194
	*(*uint16)(unsafe.Add(mBase, uint32(v25+v193<<(uint(v194)%32)))) = uint16(v168)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v202 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v201 + v202
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v201<<(uint(v202)%32)))) = uint16(v168)
	goto L41
L45:
	;
	goto L40
L46:
	;
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v371
	v375 = int32(1)
	if v375 < v45 {
		goto L67
	} else {
		goto L68
	}
L47:
	;
	v238 = int32(128)
	goto L49
L48:
	;
	v238 = int32(32)
	goto L49
L49:
	;
	if v238 <= v216 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v241 = v216
	goto L51
L51:
	;
	v259 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v259
	v264 = base.I32_div_s(v241, int32(8))
	if v45 <= v259 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if v241 < v238 {
		goto L4
	} else {
		goto L66
	}
L53:
	;
	goto L52
L54:
	;
	v349 = v241 + int32(1)
	if v349 != v238 {
		v241 = v349
		goto L51
	} else {
		goto L65
	}
L55:
	;
	v271 = int32(1)
	v273 = v271
	v276 = v271
	goto L56
L56:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v39+v276*int32(24))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v264)+4)))
	if int32(base.Ui32(int32(128))>>(uint(v241&int32(7))%32))&v297 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v323 <= int32(0) {
		goto L54
	} else {
		goto L63
	}
L58:
	;
	v319 = v273 + int32(1)
	v321 = v319 & int32(_a_F_inet_gist_picksplit_0)
	if base.Ui32(v321) <= base.Ui32(v45) {
		v273 = v319
		v276 = v321
		goto L56
	} else {
		goto L62
	}
L59:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v301 + v302
	*(*uint16)(unsafe.Add(mBase, uint32(v25+v301<<(uint(v302)%32)))) = uint16(v273)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v310 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v309 + v310
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v309<<(uint(v310)%32)))) = uint16(v273)
	goto L58
L62:
	;
	goto L57
L63:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if int32(0) < v326 {
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
	v379 = base.I32_div_s(v45, int32(2))
	v380 = v375
	goto L70
L68:
	;
	v412 = v375
	goto L69
L69:
	;
	if v45 < v412&int32(_a_F_inet_gist_picksplit_0) {
		goto L4
	} else {
		goto L73
	}
L70:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v400 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v399 + v400
	*(*uint16)(unsafe.Add(mBase, uint32(v25+v399<<(uint(v400)%32)))) = uint16(v380)
	v408 = v380 + v400
	if v408&int32(_a_F_inet_gist_picksplit_0) <= v379 {
		v380 = v408
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v412 = v408
	goto L69
L72:
	;
	goto L71
L73:
	;
	v434 = v412
	goto L74
L74:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v454 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v453 + v454
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v453<<(uint(v454)%32)))) = uint16(v434)
	v462 = v434 + v454
	if base.Ui32(v462&int32(_a_F_inet_gist_picksplit_0)) <= base.Ui32(v45) {
		v434 = v462
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
	v644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v39+v644*int32(24))))
	v650 = F_palloc0(m, int32(20))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L114
	}
L78:
	;
	v497 = v489 + int32(4)
	v499 = v490
	v500 = v492
	v501 = v492
	v502 = int32(1)
	v504 = v491
	goto L79
L79:
	;
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v502<<(uint(int32(1))%32)))))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v39+v521*int32(24))))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+2)))
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+1)))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+3)))
	if v499 < v530 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v615 == v617 {
		v625 = v613
		v627 = v617
		v630 = v614
		goto L77
	} else {
		goto L113
	}
L81:
	;
	v532 = v499
	goto L83
L82:
	;
	v532 = v530
	goto L83
L83:
	;
	if int32(0) < v532 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v536 = v525 + int32(4)
	v537 = int32(0)
	v542 = int32(8)
	v543 = base.I32_div_s(v532, v542)
	if v542 <= v532 {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v613 = v532
	goto L86
L86:
	;
	if base.Ui32(v504) < base.Ui32(v526) {
		goto L103
	} else {
		goto L104
	}
L87:
	;
	v613 = v609 + v605<<(uint(int32(3))%32)
	goto L86
L88:
	;
	goto L87
L89:
	;
	v591 = v582
	goto L100
L90:
	;
	v549 = v537
	goto L93
L91:
	;
	v566 = v537
	goto L92
L92:
	;
	v573 = v532 - v543<<(uint(int32(3))%32)
	if v573 == int32(0) {
		v605 = v566
		v609 = v537
		goto L88
	} else {
		goto L99
	}
L93:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+v549))))
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536+v549))))
	if v555 != v557 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v566 = v543
	goto L92
L95:
	;
	v582 = int32(7)
	v583 = v549
	v585 = v555
	v586 = v557
	goto L89
L96:
	;
	goto L97
L97:
	;
	v561 = v549 + int32(1)
	if v561 != v543 {
		v549 = v561
		goto L93
	} else {
		goto L98
	}
L98:
	;
	goto L94
L99:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536+v566))))
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497+v566))))
	v582 = v573
	v583 = v566
	v585 = v579
	v586 = v577
	goto L89
L100:
	;
	if int32(base.Ui32(v585^v586)>>(uint(int32(8)-v591)%32)) != 0 {
		v591 = v591 - int32(1)
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v605 = v583
	v609 = v591
	goto L88
L102:
	;
	goto L101
L103:
	;
	v614 = v504
	goto L105
L104:
	;
	v614 = v526
	goto L105
L105:
	;
	if base.Ui32(v528) < base.Ui32(v500) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v615 = v500
	goto L108
L107:
	;
	v615 = v528
	goto L108
L108:
	;
	if base.Ui32(v501) < base.Ui32(v528) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v617 = v501
	goto L111
L110:
	;
	v617 = v528
	goto L111
L111:
	;
	v619 = v502 + int32(1)
	if v619 != v493 {
		v499 = v613
		v500 = v615
		v501 = v617
		v502 = v619
		v504 = v614
		goto L79
	} else {
		goto L112
	}
L112:
	;
	goto L80
L113:
	;
	v622 = int32(0)
	v625 = v622
	v627 = v622
	v630 = v622
	goto L77
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v650)+3)) = uint8(v625)
	*(*uint8)(unsafe.Add(mBase, uint32(v650)+2)) = uint8(v630)
	*(*uint8)(unsafe.Add(mBase, uint32(v650)+1)) = uint8(v627)
	if v625 <= int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v673 = base.I32_div_s(v625, int32(8))
	v676 = v625 - v673<<(uint(int32(3))%32)
	if v676 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v660 = base.I32_div_s(v625+int32(7), int32(8))
	if v660 == int32(0) {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v663 = int32(4)
	base.MemoryCopy(m, v650+v663, base.I32_wrap_i64(v648)+v663, v660)
	goto L115
L118:
	;
	v679 = v673 + v650 + int32(4)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	v683 = v680 & (int32(-256) >> (uint(v676) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v679))) = uint8(v683)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650)+1)))
	v687 = v685
	goto L120
L119:
	;
	v687 = v627
	goto L120
L120:
	;
	if v687&int32(255) == int32(3) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v692 = int32(41)
	goto L123
L122:
	;
	v692 = int32(17)
	goto L123
L123:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v650))) = uint8(v692)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = base.I64_extend_i32_u(v650)
	v696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30))))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v39+v696*int32(24))))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+3)))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+2)))
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+1)))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v704 < int32(2) {
		v836 = v701
		v838 = v703
		v841 = v702
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v855 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30))))
	v859 = *(*int64)(unsafe.Add(mBase, uint32(v39+v855*int32(24))))
	v861 = F_palloc0(m, int32(20))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L161
	}
L125:
	;
	v708 = v700 + int32(4)
	v710 = v701
	v711 = v703
	v712 = v703
	v713 = int32(1)
	v715 = v702
	goto L126
L126:
	;
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v713<<(uint(int32(1))%32)))))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v39+v732*int32(24))))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+2)))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+1)))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+3)))
	if v710 < v741 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v826 == v828 {
		v836 = v824
		v838 = v828
		v841 = v825
		goto L124
	} else {
		goto L160
	}
L128:
	;
	v743 = v710
	goto L130
L129:
	;
	v743 = v741
	goto L130
L130:
	;
	if int32(0) < v743 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v747 = v736 + int32(4)
	v748 = int32(0)
	v753 = int32(8)
	v754 = base.I32_div_s(v743, v753)
	if v753 <= v743 {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v824 = v743
	goto L133
L133:
	;
	if base.Ui32(v715) < base.Ui32(v737) {
		goto L150
	} else {
		goto L151
	}
L134:
	;
	v824 = v820 + v816<<(uint(int32(3))%32)
	goto L133
L135:
	;
	goto L134
L136:
	;
	v802 = v793
	goto L147
L137:
	;
	v760 = v748
	goto L140
L138:
	;
	v777 = v748
	goto L139
L139:
	;
	v784 = v743 - v754<<(uint(int32(3))%32)
	if v784 == int32(0) {
		v816 = v777
		v820 = v748
		goto L135
	} else {
		goto L146
	}
L140:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708+v760))))
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747+v760))))
	if v766 != v768 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v777 = v754
	goto L139
L142:
	;
	v793 = int32(7)
	v794 = v760
	v796 = v766
	v797 = v768
	goto L136
L143:
	;
	goto L144
L144:
	;
	v772 = v760 + int32(1)
	if v772 != v754 {
		v760 = v772
		goto L140
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747+v777))))
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708+v777))))
	v793 = v784
	v794 = v777
	v796 = v790
	v797 = v788
	goto L136
L147:
	;
	if int32(base.Ui32(v796^v797)>>(uint(int32(8)-v802)%32)) != 0 {
		v802 = v802 - int32(1)
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v816 = v794
	v820 = v802
	goto L135
L149:
	;
	goto L148
L150:
	;
	v825 = v715
	goto L152
L151:
	;
	v825 = v737
	goto L152
L152:
	;
	if base.Ui32(v739) < base.Ui32(v711) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v826 = v711
	goto L155
L154:
	;
	v826 = v739
	goto L155
L155:
	;
	if base.Ui32(v712) < base.Ui32(v739) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v828 = v712
	goto L158
L157:
	;
	v828 = v739
	goto L158
L158:
	;
	v830 = v713 + int32(1)
	if v830 != v704 {
		v710 = v824
		v711 = v826
		v712 = v828
		v713 = v830
		v715 = v825
		goto L126
	} else {
		goto L159
	}
L159:
	;
	goto L127
L160:
	;
	v833 = int32(0)
	v836 = v833
	v838 = v833
	v841 = v833
	goto L124
L161:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v861)+3)) = uint8(v836)
	*(*uint8)(unsafe.Add(mBase, uint32(v861)+2)) = uint8(v841)
	*(*uint8)(unsafe.Add(mBase, uint32(v861)+1)) = uint8(v838)
	if v836 <= int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v886 = base.I32_div_s(v836, int32(8))
	v889 = v836 - v886<<(uint(int32(3))%32)
	if v889 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v871 = base.I32_div_s(v836+int32(7), int32(8))
	if v871 == int32(0) {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v874 = int32(4)
	base.MemoryCopy(m, v861+v874, base.I32_wrap_i64(v859)+v874, v871)
	goto L162
L165:
	;
	v892 = v886 + v861 + int32(4)
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892))))
	v896 = v893 & (int32(-256) >> (uint(v889) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v892))) = uint8(v896)
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v861)+1)))
	v900 = v898
	goto L167
L166:
	;
	v900 = v838
	goto L167
L167:
	;
	if v900&int32(255) == int32(3) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v905 = int32(41)
	goto L170
L169:
	;
	v905 = int32(17)
	goto L170
L170:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v861))) = uint8(v905)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = base.I64_extend_i32_u(v861)
	return v20 & int64(4294967295)
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
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
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
			v32 = v14
		} else {
			v32 = v22
		}
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
		v38 = base.I32_div_s(v33, int32(8))
		v39 = F_memcmp(m, v26, v28, v38)
		mBase = m.M
		if v39 != 0 {
			v125 = v39
			v135 = v125
		} else {
			v40 = int32(0)
			v43 = v33 - v38<<(uint(int32(3))%32)
			if v43 <= v40 {
				v125 = v40
				v135 = v125
			} else {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v38))))
				v48 = int32(128)
				v49 = v47 & v48
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v38))))
				if v49 != v51&v48 {
					v126 = v49
					if v126 != 0 {
						v129 = int32(1)
					} else {
						v129 = int32(-1)
					}
					v135 = v129
				} else {
					if v43 == int32(1) {
						v125 = v40
						v135 = v125
					} else {
						v57 = int32(1)
						v59 = int32(128)
						v60 = v47 << (uint(v57) % 32) & v59
						if v60 != v51<<(uint(v57)%32)&v59 {
							v126 = v60
							if v126 != 0 {
								v129 = int32(1)
							} else {
								v129 = int32(-1)
							}
							v135 = v129
						} else {
							if v43 < int32(3) {
								v125 = v40
								v135 = v125
							} else {
								v68 = int32(2)
								v70 = int32(128)
								v71 = v47 << (uint(v68) % 32) & v70
								if v71 != v51<<(uint(v68)%32)&v70 {
									v126 = v71
									if v126 != 0 {
										v129 = int32(1)
									} else {
										v129 = int32(-1)
									}
									v135 = v129
								} else {
									if v43 == int32(3) {
										v125 = v40
										v135 = v125
									} else {
										v79 = int32(3)
										v81 = int32(128)
										v82 = v47 << (uint(v79) % 32) & v81
										if v82 != v51<<(uint(v79)%32)&v81 {
											v126 = v82
											if v126 != 0 {
												v129 = int32(1)
											} else {
												v129 = int32(-1)
											}
											v135 = v129
										} else {
											if v43 < int32(5) {
												v125 = v40
												v135 = v125
											} else {
												v90 = int32(4)
												v92 = int32(128)
												v93 = v47 << (uint(v90) % 32) & v92
												if v93 != v51<<(uint(v90)%32)&v92 {
													v126 = v93
													if v126 != 0 {
														v129 = int32(1)
													} else {
														v129 = int32(-1)
													}
													v135 = v129
												} else {
													if v43 == int32(5) {
														v125 = v40
														v135 = v125
													} else {
														v101 = int32(5)
														v103 = int32(128)
														v104 = v47 << (uint(v101) % 32) & v103
														if v104 != v51<<(uint(v101)%32)&v103 {
															v126 = v104
															if v126 != 0 {
																v129 = int32(1)
															} else {
																v129 = int32(-1)
															}
															v135 = v129
														} else {
															if v43 < int32(7) {
																v125 = v40
																v135 = v125
															} else {
																v112 = int32(6)
																v114 = int32(128)
																v115 = v47 << (uint(v112) % 32) & v114
																if v115 != v51<<(uint(v112)%32)&v114 {
																	v126 = v115
																	if v126 != 0 {
																		v129 = int32(1)
																	} else {
																		v129 = int32(-1)
																	}
																	v135 = v129
																} else {
																	v125 = v40
																	v135 = v125
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
		if v135 != 0 {
			v177 = v135
			return v177
		} else {
			v137 = int32(1)
			v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v139&v137 != 0 {
				v142 = v137
			} else {
				v142 = int32(4)
			}
			v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v142)+1)))
			v145 = int32(1)
			v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v147&v145 != 0 {
				v150 = v145
			} else {
				v150 = int32(4)
			}
			v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v150)+1)))
			v153 = v144 - v152
			v154 = int32(0)
			if base.B2i32(v154 < v153)&base.B2i32(v154 <= l2)|base.B2i32(v144 == v152)&base.B2i32(base.Ui32(l2+int32(1)) <= base.Ui32(int32(2))) != 0 {
				v177 = int32(0)
				return v177
			} else {
				v166 = int32(0)
				if v166 <= v153 {
					v169 = l2
				} else {
					v169 = v166
				}
				if l2 <= int32(0) {
					v172 = v169
				} else {
					v172 = l2
				}
				return v172
			}
		}
	} else {
		v177 = v15 - v23
		return v177
	}
}
func F_inet_spg_config(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v2)+12)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9783935500938)
	return int64(0)
}
