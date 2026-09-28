package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InsertPgClassTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(320)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+288)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+296)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+304)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+312)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = base.I64_extend_i32_u(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v12
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+32)) = uint16(v6)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = base.I64_extend_i32_u(v11 + int32(4))
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+68)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v36
	v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+72)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v38
	v40 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+76)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v40
	v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v42
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+84)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v44
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+88)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v46
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+92)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v48
	v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+96)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v50
	v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+100)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+128)) = v52
	v54 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+104)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+136)) = v54
	v56 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+108)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+144)) = v56
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+112)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v58
	v60 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+116)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v60
	v62 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+117)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = v62
	v64 = int64(*(*int8)(unsafe.Add(mBase, uint32(v11)+118)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = v64
	v66 = int64(*(*int8)(unsafe.Add(mBase, uint32(v11)+119)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v66
	v68 = int64(*(*int16)(unsafe.Add(mBase, uint32(v11)+120)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+192)) = v68
	v70 = int64(*(*int16)(unsafe.Add(mBase, uint32(v11)+122)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+200)) = v70
	v72 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+124)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+208)) = v72
	v74 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+125)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+216)) = v74
	v76 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+127)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+232)) = v76
	v78 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+128)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+240)) = v78
	v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+126)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+224)) = v80
	v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+129)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+248)) = v82
	v84 = int64(*(*int8)(unsafe.Add(mBase, uint32(v11)+130)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+256)) = v84
	v86 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+131)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+264)) = v86
	v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+132)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+272)) = v88
	v90 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+136)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+280)) = v90
	v92 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v11)+140)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+288)) = v92
	if l3 != v12 {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+296)) = l3
	} else {
		v97 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)) = uint8(v97)
	}
	if l4 != int64(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+304)) = l4
	} else {
		v102 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+32)) = uint8(v102)
	}
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+33)) = uint8(v104)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v109 = F_heap_form_tuple(m, v106, v9+int32(48), v9)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		return
	} else {
		F_CatalogTupleInsert(m, l0, v109)
		mBase = m.M
		v112 = m.ExcPending
		if v112 != 0 {
			return
		} else {
			F_pfree(m, v109)
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return
			} else {
				m.G0 = v9 + int32(320)
				return
			}
		}
	}
}
func F_MakePGDirectory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_MakePGDirectory[0]))
	v4 = F_mkdir(m, l0, v3)
	mBase = m.M
	return v4
}
func F_PGSemaphoreReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(0)
	v4 = m.Env.Pgmem_sem(m, int32(3), l0, v3)
	mBase = m.M
	if v4 < v3 {
		*(*int32)(unsafe.Add(mBase, _c_F_PGSemaphoreReset[0])) = int32(0) - v4
	} else {
	}
	return
}
func F_PG_char_to_encoding(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
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
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = int32(-1)
	if v2 == int32(0) {
		v89 = v14
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I64_extend_i32_s(v89)
L2:
	;
	m.G0 = v12 - int32(-64)
	goto L1
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
	if v17 == int32(0) {
		v89 = v14
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = F_strlen(m, v2)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v20) {
		v89 = v14
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = v2
	v24 = v17
	v25 = v12
	goto L6
L6:
	;
	v33 = F_isalnum(m, v24&int32(255))
	mBase = m.M
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v50)
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12))))
	v55 = int32(_a_F_PG_char_to_encoding_0)
	v56 = int32(_a_F_PG_char_to_encoding_1)
	goto L15
L8:
	;
	if base.Ui32((v24-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v46 = v25
	goto L10
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v47 != 0 {
		v23 = v23 + int32(1)
		v24 = v47
		v25 = v46
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v42 = v24 | int32(32)
	goto L13
L12:
	;
	v42 = v24
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v42)
	v46 = v25 + int32(1)
	goto L10
L14:
	;
	goto L7
L15:
	;
	v68 = v56 + (v55-v56)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v69))))
	v71 = v54 - v70
	if v71 != 0 {
		v74 = v71
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v89 = v14
	goto L2
L17:
	;
	v78 = base.B2i32(v74 < int32(0))
	if v74 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v72 = F_strcmp(m, v12, v69)
	mBase = m.M
	if v72 != 0 {
		v74 = v72
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v89 = v73
	goto L2
L20:
	;
	v79 = v68 - int32(8)
	goto L22
L21:
	;
	v79 = v55
	goto L22
L22:
	;
	if v74 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v56
	goto L25
L24:
	;
	v82 = v68 + int32(8)
	goto L25
L25:
	;
	if base.Ui32(v82) <= base.Ui32(v79) {
		v55 = v79
		v56 = v82
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
}
func F_PgArchiverMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v514 int64
	_ = v514
	var v516 int64
	_ = v516
	var v529 int32
	_ = v529
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
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
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int64
	_ = v730
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int64
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int64
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1087 int64
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1104 int32
	_ = v1104
	var v1108 int64
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1121 int32
	_ = v1121
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1228 int32
	_ = v1228
	var v1229 int64
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1535 int64
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int64
	_ = v1554
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1564 int64
	_ = v1564
	var v1566 int64
	_ = v1566
	var v1568 int64
	_ = v1568
	var v1570 int64
	_ = v1570
	var v1572 int64
	_ = v1572
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1601 int64
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1620 int64
	_ = v1620
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int64
	_ = v1630
	var v1632 int64
	_ = v1632
	var v1634 int64
	_ = v1634
	var v1636 int64
	_ = v1636
	var v1638 int64
	_ = v1638
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1733 int32
	_ = v1733
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v23 = int32(967)
	switch v23 {
	case 0, 2:
		goto L4
	default:
		goto L5
	}
L3:
	;
	v58 = int32(0)
	v60 = m.G0
	v62 = v60 - int32(32)
	m.G0 = v62
	switch v58 {
	case 0, 2:
		goto L14
	default:
		goto L15
	}
L4:
	;
	F_sigemptyset(m, v20+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(268435456)
	switch v23 {
	case 0:
		goto L9
	default:
		goto L7
	case 2:
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[0])) = int32(965)
	goto L4
L6:
	;
	goto L11
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(0)
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(-2)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v52 = F___sigaction(m, int32(1), v20+int32(12), int32(0))
	mBase = m.M
	m.G0 = v20 + int32(32)
	goto L3
L13:
	;
	v102 = m.G0
	v104 = v102 - int32(32)
	m.G0 = v104
	v107 = int32(969)
	switch v107 {
	case 0, 2:
		goto L24
	default:
		goto L25
	}
L14:
	;
	F_sigemptyset(m, v62+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = int32(268435456)
	switch v58 {
	case 0:
		goto L19
	default:
		goto L17
	case 2:
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[1])) = int32(-2)
	goto L14
L16:
	;
	goto L21
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = int32(0)
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = int32(-2)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v94 = F___sigaction(m, int32(2), v62+int32(12), int32(0))
	mBase = m.M
	m.G0 = v62 + int32(32)
	goto L13
L23:
	;
	v142 = int32(0)
	v144 = m.G0
	v146 = v144 - int32(32)
	m.G0 = v146
	switch v142 {
	case 0, 2:
		goto L34
	default:
		goto L35
	}
L24:
	;
	F_sigemptyset(m, v104+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(268435456)
	switch v107 {
	case 0:
		goto L29
	default:
		goto L27
	case 2:
		goto L28
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[2])) = int32(967)
	goto L24
L26:
	;
	goto L31
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L26
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(0)
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(-2)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v136 = F___sigaction(m, int32(15), v104+int32(12), int32(0))
	mBase = m.M
	m.G0 = v104 + int32(32)
	goto L23
L33:
	;
	v184 = int32(0)
	v186 = m.G0
	v188 = v186 - int32(32)
	m.G0 = v188
	switch v184 {
	case 0, 2:
		goto L44
	default:
		goto L45
	}
L34:
	;
	F_sigemptyset(m, v146+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v146)+24)) = int32(268435456)
	switch v142 {
	case 0:
		goto L39
	default:
		goto L37
	case 2:
		goto L38
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[3])) = int32(-2)
	goto L34
L36:
	;
	goto L41
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = int32(0)
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = int32(-2)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v178 = F___sigaction(m, int32(14), v146+int32(12), int32(0))
	mBase = m.M
	m.G0 = v146 + int32(32)
	goto L33
L43:
	;
	v228 = m.G0
	v230 = v228 - int32(32)
	m.G0 = v230
	v233 = int32(970)
	switch v233 {
	case 0, 2:
		goto L54
	default:
		goto L55
	}
L44:
	;
	F_sigemptyset(m, v188+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v188)+24)) = int32(268435456)
	switch v184 {
	case 0:
		goto L49
	default:
		goto L47
	case 2:
		goto L48
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[4])) = int32(-2)
	goto L44
L46:
	;
	goto L51
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = int32(0)
	goto L46
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = int32(-2)
	goto L46
L51:
	;
	goto L52
L52:
	;
	v220 = F___sigaction(m, int32(13), v188+int32(12), int32(0))
	mBase = m.M
	m.G0 = v188 + int32(32)
	goto L43
L53:
	;
	v270 = m.G0
	v272 = v270 - int32(32)
	m.G0 = v272
	v275 = int32(1007)
	switch v275 {
	case 0, 2:
		goto L64
	default:
		goto L65
	}
L54:
	;
	F_sigemptyset(m, v230+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = int32(268435456)
	switch v233 {
	case 0:
		goto L59
	default:
		goto L57
	case 2:
		goto L58
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[5])) = int32(968)
	goto L54
L56:
	;
	goto L61
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = int32(0)
	goto L56
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+12)) = int32(-2)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v262 = F___sigaction(m, int32(10), v230+int32(12), int32(0))
	mBase = m.M
	m.G0 = v230 + int32(32)
	goto L53
L63:
	;
	v312 = m.G0
	v314 = v312 - int32(32)
	m.G0 = v314
	v316 = int32(2)
	switch v316 {
	case 0, 2:
		goto L74
	default:
		goto L75
	}
L64:
	;
	F_sigemptyset(m, v272+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v272)+24)) = int32(268435456)
	switch v275 {
	case 0:
		goto L69
	default:
		goto L67
	case 2:
		goto L68
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[6])) = int32(1005)
	goto L64
L66:
	;
	goto L71
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v272)+12)) = int32(_a_F_PgArchiverMain_0)
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+12)) = int32(0)
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+12)) = int32(-2)
	goto L66
L71:
	;
	goto L72
L72:
	;
	v304 = F___sigaction(m, int32(12), v272+int32(12), int32(0))
	mBase = m.M
	m.G0 = v272 + int32(32)
	goto L63
L73:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_PgArchiverMain_1), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L83
	}
L74:
	;
	F_sigemptyset(m, v314+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = int32(268435456)
	switch v316 {
	case 0:
		goto L79
	default:
		goto L77
	case 2:
		goto L78
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[7])) = int32(0)
	goto L74
L76:
	;
	goto L80
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(_a_F_PgArchiverMain_0)
	v339 = int32(268435461)
	goto L76
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(0)
	v339 = int32(268435457)
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = int32(-2)
	v339 = int32(268435457)
	goto L76
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = v339
	goto L82
L82:
	;
	v346 = F___sigaction(m, int32(17), v314+int32(12), int32(0))
	mBase = m.M
	m.G0 = v314 + int32(32)
	goto L73
L83:
	;
	F_on_shmem_exit(m, int32(1006), int64(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[8]))
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v359))) = v361
	v365 = F_palloc(m, int32(2888))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10])) = v365
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v368
	v373 = F_binaryheap_allocate(m, int32(64), int32(1007), v368)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v376))) = v373
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[11]))
	v385 = F_AllocSetContextCreateInternal(m, v380, int32(_a_F_PgArchiverMain_2), int32(0), int32(_a_F_PgArchiverMain_3), int32(_a_F_PgArchiverMain_4))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[12])) = v385
	v388 = m.G0
	v390 = v388 - int32(16)
	m.G0 = v390
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[13]))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	if v394 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v480 = m.G0
	v482 = v480 - int32(3392)
	m.G0 = v482
	goto L118
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L115
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L112
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L107
	}
L92:
	;
	v410 = m.T0[v408].(func(*base.Module) int32)(m)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L99
	}
L93:
	;
	v408 = int32(1008)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[14]))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	if v400 != 0 {
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v402 = int32(0)
	v404 = F_load_external_function(m, v393, int32(_a_F_PgArchiverMain_5), v402, v402)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v404 == int32(0) {
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v408 = v404
	goto L92
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[15])) = v410
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	if v413 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L100:
	;
	v418 = F_palloc0(m, int32(4))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[16])) = v418
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[15]))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	if v423 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	m.T0[v423].(func(*base.Module, int32))(m, v418)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_before_shmem_exit(m, int32(1009), int64(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	m.G0 = v390 + int32(16)
	goto L88
L107:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(_a_F_PgArchiverMain_6), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v446 = F_errdetail(m, int32(_a_F_PgArchiverMain_7), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(924), int32(_a_F_PgArchiverMain_9))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = int32(_a_F_PgArchiverMain_5)
	F_errmsg(m, int32(_a_F_PgArchiverMain_10), v390)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(939), int32(_a_F_PgArchiverMain_9))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errmsg(m, int32(_a_F_PgArchiverMain_11), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(945), int32(_a_F_PgArchiverMain_9))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[17]))
	v500 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v499))) = v500
	v505 = base.AtomicRmwOr32(m, v500, int32(_a_F_PgArchiverMain_12), v500)
	goto L121
L119:
	;
	m.G0 = v482 + int32(3392)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L1
	} else {
		goto L413
	}
L120:
	;
	goto L119
L121:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[18]))
	F_ProcessPgArchInterrupts(m)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[19]))
	if v511 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v529)+4)) = int32(0)
	goto L129
L124:
	;
	v514 = F_time(m)
	mBase = m.M
	v516 = *(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[20]))
	if v516 == int64(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PgArchiverMain[20])) = v514
	goto L123
L126:
	;
	goto L127
L127:
	;
	if base.B2i32(v514 < v516)|base.B2i32(int64(59) < v514-v516) != 0 {
		goto L120
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[8]))
	v547 = base.AtomicRmwXchg32(m, v544, int32(4), int32(0))
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	if v547 == int32(1) {
		goto L137
	} else {
		goto L138
	}
L130:
	;
	if v507 != 0 {
		goto L120
	} else {
		goto L410
	}
L131:
	;
	goto L130
L132:
	;
	v1346 = int32(0)
	v1349 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[19]))
	if v1349 != 0 {
		goto L131
	} else {
		goto L333
	}
L133:
	;
	v1253 = v1244 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+4)) = v1253
	v1256 = v482 + int32(1296)
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1242+v1253<<(uint(int32(2))%32))+8))
	if (v1260^v1256)&int32(3) != 0 {
		goto L315
	} else {
		goto L316
	}
L134:
	;
	v1217 = int32(0)
	v1218 = v1214
	goto L308
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L1
	} else {
		goto L304
	}
L136:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v690 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v689)+8)) = uint8(v690)
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = int32(0)
	goto L171
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v549)+4)) = int32(0)
	v679 = v549
	goto L136
L138:
	;
	goto L139
L139:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	if v554 <= int32(0) {
		v679 = v549
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v557 = v554
	v558 = v549
	goto L141
L141:
	;
	v569 = v557 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v558)+4)) = v569
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v558+v569<<(uint(int32(2))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+164)) = int32(_a_F_PgArchiverMain_13)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+160)) = v574
	v579 = v482 + int32(1344)
	v584 = F_pg_snprintf(m, v579, int32(1024), int32(_a_F_PgArchiverMain_14), v482+int32(160))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v679 = v674
	goto L136
L143:
	;
	v590 = F___fstatat(m, int32(-100), v579, v482+int32(176), int32(0))
	mBase = m.M
	goto L144
L144:
	;
	if v590 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v594 = v482 + int32(1296)
	if (v574^v594)&int32(3) != 0 {
		goto L151
	} else {
		goto L152
	}
L146:
	;
	goto L147
L147:
	;
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[21]))
	if v670 != int32(44) {
		goto L135
	} else {
		goto L169
	}
L148:
	;
	goto L132
L149:
	;
	goto L148
L150:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v649))) = uint8(v648)
	if v648&int32(255) == int32(0) {
		goto L149
	} else {
		goto L165
	}
L151:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	v647 = v574
	v648 = v600
	v649 = v594
	goto L150
L152:
	;
	goto L153
L153:
	;
	if v574&int32(3) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v604 = v574
	v606 = v594
	goto L157
L155:
	;
	v618 = v574
	v620 = v594
	goto L156
L156:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	v625 = int32(-2139062144)
	if (int32(16843008)-v622|v622)&v625 != v625 {
		v647 = v618
		v648 = v622
		v649 = v620
		goto L150
	} else {
		goto L161
	}
L157:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604))))
	*(*uint8)(unsafe.Add(mBase, uint32(v606))) = uint8(v607)
	if v607 == int32(0) {
		goto L149
	} else {
		goto L159
	}
L158:
	;
	v618 = v614
	v620 = v612
	goto L156
L159:
	;
	v611 = int32(1)
	v612 = v606 + v611
	v614 = v604 + v611
	if v614&int32(3) != 0 {
		v604 = v614
		v606 = v612
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v630 = v618
	v631 = v622
	v632 = v620
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v632))) = v631
	v634 = int32(4)
	v635 = v632 + v634
	v637 = v630 + v634
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	v642 = int32(-2139062144)
	if (int32(16843008)-v639|v639)&v642 == v642 {
		v630 = v637
		v631 = v639
		v632 = v635
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v647 = v637
	v648 = v639
	v649 = v635
	goto L150
L164:
	;
	goto L163
L165:
	;
	v656 = v647
	v658 = v649
	goto L166
L166:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+1)) = uint8(v659)
	v661 = int32(1)
	if v659 != 0 {
		v656 = v656 + v661
		v658 = v658 + v661
		goto L166
	} else {
		goto L168
	}
L167:
	;
	goto L149
L168:
	;
	goto L167
L169:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	if int32(0) < v675 {
		v557 = v675
		v558 = v674
		goto L141
	} else {
		goto L170
	}
L170:
	;
	goto L142
L171:
	;
	v695 = v482 + int32(2368)
	v699 = F_pg_snprintf(m, v695, int32(1024), int32(_a_F_PgArchiverMain_15), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v701 = F_AllocateDir(m, v695)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v703 = F_ReadDir(m, v701, v695)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v703 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v706 = v703
	goto L178
L176:
	;
	goto L177
L177:
	;
	F_FreeDir(m, v701)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L297
	}
L178:
	;
	v717 = v706 + int32(19)
	v718 = F_strlen(m, v717)
	mBase = m.M
	if base.Ui32(v718-int32(47)) < base.Ui32(int32(-25)) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L177
L180:
	;
	v1162 = F_ReadDir(m, v701, v482+int32(2368))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L295
	}
L181:
	;
	v723 = int32(_a_F_PgArchiverMain_16)
	v727 = m.G0
	v729 = v727 - int32(32)
	v730 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v729)+24)) = v730
	*(*int64)(unsafe.Add(mBase, uint32(v729)+16)) = v730
	*(*int64)(unsafe.Add(mBase, uint32(v729)+8)) = v730
	*(*int64)(unsafe.Add(mBase, uint32(v729))) = v730
	v738 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PgArchiverMain[22])))
	if v738 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v808 = v718 - int32(6)
	if base.Ui32(v806) < base.Ui32(v808) {
		goto L180
	} else {
		goto L201
	}
L183:
	;
	v806 = int32(0)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PgArchiverMain[23])))
	if v742 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v746 = v717
	goto L189
L187:
	;
	goto L188
L188:
	;
	v756 = v723
	v757 = v738
	goto L192
L189:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746))))
	if v752 == v738 {
		v746 = v746 + int32(1)
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v806 = v746 - v717
	goto L182
L191:
	;
	goto L190
L192:
	;
	v764 = v729 + int32(base.Ui32(v757)>>(uint(int32(3))%32))&int32(28)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)))
	v766 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = v765 | v766<<(uint(v757)%32)
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756)+1)))
	if v770 != 0 {
		v756 = v756 + v766
		v757 = v770
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717))))
	if v773 == int32(0) {
		v796 = v717
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L193
L195:
	;
	v806 = v796 - v717
	goto L182
L196:
	;
	v777 = v717
	v778 = v773
	goto L197
L197:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v729+int32(base.Ui32(v778)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v786)>>(uint(v778)%32))&int32(1) == int32(0) {
		v796 = v777
		goto L195
	} else {
		goto L199
	}
L198:
	;
	v796 = v794
	goto L195
L199:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+1)))
	v794 = v777 + int32(1)
	if v792 != 0 {
		v777 = v794
		v778 = v792
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v810 = v808 + v717
	v811 = int32(_a_F_PgArchiverMain_13)
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810))))
	v817 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PgArchiverMain[24])))
	if base.B2i32(v814 == int32(0))|base.B2i32(v814 != v817) != 0 {
		v835 = v814
		v836 = v817
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v835-v836 != 0 {
		goto L180
	} else {
		goto L209
	}
L203:
	;
	goto L202
L204:
	;
	v820 = v810
	v821 = v811
	goto L205
L205:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821)+1)))
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+1)))
	if v825 == int32(0) {
		v835 = v825
		v836 = v824
		goto L203
	} else {
		goto L207
	}
L206:
	;
	v835 = v825
	v836 = v824
	goto L203
L207:
	;
	v828 = int32(1)
	if v825 == v824 {
		v820 = v820 + v828
		v821 = v821 + v828
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	if v808 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	base.MemoryCopy(m, v482+int32(1344), v717, v808)
	goto L212
L211:
	;
	goto L212
L212:
	;
	v842 = v482 + int32(1344)
	v844 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v842+v808))) = uint8(v844)
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	if v849 <= int32(63) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v856 = v847 + v849*int32(41) + int32(264)
	if (v842^v856)&int32(3) != 0 {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	goto L215
L215:
	;
	v943 = *(*int64)(unsafe.Add(mBase, uint32(v848)+24))
	v944 = int32(0)
	v946 = base.I32_wrap_i64(base.I64_extend_i32_u(v482 + int32(1344)))
	v947 = base.I32_wrap_i64(v943)
	v948 = F_strlen(m, v947)
	mBase = m.M
	if v948 != int32(16) {
		v961 = v944
		goto L241
	} else {
		goto L242
	}
L216:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	F_binaryheap_add_unordered(m, v931, base.I64_extend_i32_u(v856))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L237
	}
L217:
	;
	goto L216
L218:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v911))) = uint8(v910)
	if v910&int32(255) == int32(0) {
		goto L217
	} else {
		goto L233
	}
L219:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	v909 = v842
	v910 = v862
	v911 = v856
	goto L218
L220:
	;
	goto L221
L221:
	;
	if v842&int32(3) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v866 = v842
	v868 = v856
	goto L225
L223:
	;
	v880 = v842
	v882 = v856
	goto L224
L224:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	v887 = int32(-2139062144)
	if (int32(16843008)-v884|v884)&v887 != v887 {
		v909 = v880
		v910 = v884
		v911 = v882
		goto L218
	} else {
		goto L229
	}
L225:
	;
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	*(*uint8)(unsafe.Add(mBase, uint32(v868))) = uint8(v869)
	if v869 == int32(0) {
		goto L217
	} else {
		goto L227
	}
L226:
	;
	v880 = v876
	v882 = v874
	goto L224
L227:
	;
	v873 = int32(1)
	v874 = v868 + v873
	v876 = v866 + v873
	if v876&int32(3) != 0 {
		v866 = v876
		v868 = v874
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v892 = v880
	v893 = v884
	v894 = v882
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v894))) = v893
	v896 = int32(4)
	v897 = v894 + v896
	v899 = v892 + v896
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v892)+4))
	v904 = int32(-2139062144)
	if (int32(16843008)-v901|v901)&v904 == v904 {
		v892 = v899
		v893 = v901
		v894 = v897
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v909 = v899
	v910 = v901
	v911 = v897
	goto L218
L232:
	;
	goto L231
L233:
	;
	v918 = v909
	v920 = v911
	goto L234
L234:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v920)+1)) = uint8(v921)
	v923 = int32(1)
	if v921 != 0 {
		v918 = v918 + v923
		v920 = v920 + v923
		goto L234
	} else {
		goto L236
	}
L235:
	;
	goto L217
L236:
	;
	goto L235
L237:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if v938 != int32(64) {
		goto L180
	} else {
		goto L238
	}
L238:
	;
	F_binaryheap_build(m, v937)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	goto L180
L240:
	;
	if v980 <= int32(0) {
		goto L180
	} else {
		goto L256
	}
L241:
	;
	v962 = F_strlen(m, v946)
	mBase = m.M
	if v962 == int32(16) {
		goto L247
	} else {
		goto L248
	}
L242:
	;
	v952 = F_strspn(m, v947, int32(_a_F_PgArchiverMain_17))
	mBase = m.M
	if v952 != int32(8) {
		v961 = v944
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v958 = F_strcmp(m, v947+int32(8), int32(_a_F_PgArchiverMain_18))
	mBase = m.M
	v961 = base.B2i32(v958 == int32(0))
	goto L241
L244:
	;
	v979 = F_strcmp(m, v947, v946)
	mBase = m.M
	v980 = v979
	goto L240
L245:
	;
	if v961 != 0 {
		goto L253
	} else {
		goto L254
	}
L246:
	;
	v972 = F_strcmp(m, v946+int32(8), int32(_a_F_PgArchiverMain_18))
	mBase = m.M
	if v961 == base.B2i32(v972 == int32(0)) {
		goto L244
	} else {
		goto L252
	}
L247:
	;
	v966 = F_strspn(m, v946, int32(_a_F_PgArchiverMain_17))
	mBase = m.M
	if v966 == int32(8) {
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	if v961 != 0 {
		goto L245
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	goto L244
L252:
	;
	goto L245
L253:
	;
	v978 = int32(-1)
	goto L255
L254:
	;
	v978 = int32(1)
	goto L255
L255:
	;
	v980 = v978
	goto L240
L256:
	;
	v984 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)))
	v986 = F_binaryheap_remove_first(m, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v988 = base.I32_wrap_i64(v986)
	v990 = v482 + int32(1344)
	if (v990^v988)&int32(3) != 0 {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1069)))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	if v1071 < v1072 {
		goto L280
	} else {
		goto L281
	}
L259:
	;
	goto L258
L260:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1045))) = uint8(v1044)
	if v1044&int32(255) == int32(0) {
		goto L259
	} else {
		goto L275
	}
L261:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	v1043 = v990
	v1044 = v996
	v1045 = v988
	goto L260
L262:
	;
	goto L263
L263:
	;
	if v990&int32(3) != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1000 = v990
	v1002 = v988
	goto L267
L265:
	;
	v1014 = v990
	v1016 = v988
	goto L266
L266:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1021 = int32(-2139062144)
	if (int32(16843008)-v1018|v1018)&v1021 != v1021 {
		v1043 = v1014
		v1044 = v1018
		v1045 = v1016
		goto L260
	} else {
		goto L271
	}
L267:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1002))) = uint8(v1003)
	if v1003 == int32(0) {
		goto L259
	} else {
		goto L269
	}
L268:
	;
	v1014 = v1010
	v1016 = v1008
	goto L266
L269:
	;
	v1007 = int32(1)
	v1008 = v1002 + v1007
	v1010 = v1000 + v1007
	if v1010&int32(3) != 0 {
		v1000 = v1010
		v1002 = v1008
		goto L267
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	v1026 = v1014
	v1027 = v1018
	v1028 = v1016
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1027
	v1030 = int32(4)
	v1031 = v1028 + v1030
	v1033 = v1026 + v1030
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+4))
	v1038 = int32(-2139062144)
	if (int32(16843008)-v1035|v1035)&v1038 == v1038 {
		v1026 = v1033
		v1027 = v1035
		v1028 = v1031
		goto L272
	} else {
		goto L274
	}
L273:
	;
	v1043 = v1033
	v1044 = v1035
	v1045 = v1031
	goto L260
L274:
	;
	goto L273
L275:
	;
	v1052 = v1043
	v1054 = v1045
	goto L276
L276:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1054)+1)) = uint8(v1055)
	v1057 = int32(1)
	if v1055 != 0 {
		v1052 = v1052 + v1057
		v1054 = v1054 + v1057
		goto L276
	} else {
		goto L278
	}
L277:
	;
	goto L259
L278:
	;
	goto L277
L279:
	;
	goto L180
L280:
	;
	v1075 = v1070 + int32(24)
	v1076 = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v1075+v1071<<(uint(v1076)%32)))) = v986 & int64(4294967295)
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	*(*int32)(unsafe.Add(mBase, uint32(v1070))) = v1080 + int32(1)
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v1075+v1080<<(uint(v1076)%32))))
	if v1080 == int32(0) {
		v1121 = int32(0)
		goto L283
	} else {
		goto L284
	}
L281:
	;
	goto L282
L282:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L292
	}
L283:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1075+v1121<<(uint(int32(3))%32)))) = v1087
	goto L279
L284:
	;
	v1091 = v1080
	goto L285
L285:
	;
	v1104 = base.I32_div_s(v1091-int32(1), int32(2))
	v1108 = *(*int64)(unsafe.Add(mBase, uint32(v1075+v1104<<(uint(int32(3))%32))))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+16))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+12))
	v1111 = m.T0[v1110].(func(*base.Module, int64, int64, int32) int32)(m, v1087, v1108, v1109)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L287
	}
L286:
	;
	v1121 = v1104
	goto L283
L287:
	;
	if v1111 <= int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1121 = v1091
	goto L283
L289:
	;
	goto L290
L290:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1075+v1091<<(uint(int32(3))%32)))) = v1108
	if base.Ui32(int32(2)) < base.Ui32(v1091) {
		v1091 = v1104
		goto L285
	} else {
		goto L291
	}
L291:
	;
	goto L286
L292:
	;
	F_errmsg_internal(m, int32(_a_F_PgArchiverMain_19), int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_20), int32(159), int32(_a_F_PgArchiverMain_21))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	if v1162 != 0 {
		v706 = v1162
		goto L178
	} else {
		goto L296
	}
L296:
	;
	goto L179
L297:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1178)))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1179)))
	if v1180 == int32(0) {
		goto L131
	} else {
		goto L298
	}
L298:
	;
	if int32(64) <= v1180 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+4)) = v1180
	v1214 = v1178
	goto L134
L300:
	;
	goto L301
L301:
	;
	F_binaryheap_build(m, v1179)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1189)))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	*(*int32)(unsafe.Add(mBase, uint32(v1189)+4)) = v1191
	if int32(0) < v1191 {
		v1214 = v1189
		goto L134
	} else {
		goto L303
	}
L303:
	;
	v1242 = v1189
	v1244 = v1191
	goto L133
L304:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+144)) = v482 + int32(1344)
	F_errmsg(m, int32(_a_F_PgArchiverMain_22), v482+int32(144))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(685), int32(_a_F_PgArchiverMain_23))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1218)))
	v1229 = F_binaryheap_remove_first(m, v1228)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L310
	}
L309:
	;
	v1242 = v1232
	v1244 = v1239
	goto L133
L310:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[10]))
	*(*uint32)(unsafe.Add(mBase, uint32(v1232+v1217<<(uint(int32(2))%32))+8)) = uint32(v1229)
	v1238 = v1217 + int32(1)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+4))
	if v1238 < v1239 {
		v1217 = v1238
		v1218 = v1232
		goto L308
	} else {
		goto L311
	}
L311:
	;
	goto L309
L312:
	;
	goto L132
L313:
	;
	goto L312
L314:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1315))) = uint8(v1314)
	if v1314&int32(255) == int32(0) {
		goto L313
	} else {
		goto L329
	}
L315:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1260))))
	v1313 = v1260
	v1314 = v1266
	v1315 = v1256
	goto L314
L316:
	;
	goto L317
L317:
	;
	if v1260&int32(3) != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1270 = v1260
	v1272 = v1256
	goto L321
L319:
	;
	v1284 = v1260
	v1286 = v1256
	goto L320
L320:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1284)))
	v1291 = int32(-2139062144)
	if (int32(16843008)-v1288|v1288)&v1291 != v1291 {
		v1313 = v1284
		v1314 = v1288
		v1315 = v1286
		goto L314
	} else {
		goto L325
	}
L321:
	;
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1270))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1272))) = uint8(v1273)
	if v1273 == int32(0) {
		goto L313
	} else {
		goto L323
	}
L322:
	;
	v1284 = v1280
	v1286 = v1278
	goto L320
L323:
	;
	v1277 = int32(1)
	v1278 = v1272 + v1277
	v1280 = v1270 + v1277
	if v1280&int32(3) != 0 {
		v1270 = v1280
		v1272 = v1278
		goto L321
	} else {
		goto L324
	}
L324:
	;
	goto L322
L325:
	;
	v1296 = v1284
	v1297 = v1288
	v1298 = v1286
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1298))) = v1297
	v1300 = int32(4)
	v1301 = v1298 + v1300
	v1303 = v1296 + v1300
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+4))
	v1308 = int32(-2139062144)
	if (int32(16843008)-v1305|v1305)&v1308 == v1308 {
		v1296 = v1303
		v1297 = v1305
		v1298 = v1301
		goto L326
	} else {
		goto L328
	}
L327:
	;
	v1313 = v1303
	v1314 = v1305
	v1315 = v1301
	goto L314
L328:
	;
	goto L327
L329:
	;
	v1322 = v1313
	v1324 = v1315
	goto L330
L330:
	;
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1324)+1)) = uint8(v1325)
	v1327 = int32(1)
	if v1325 != 0 {
		v1322 = v1322 + v1327
		v1324 = v1324 + v1327
		goto L330
	} else {
		goto L332
	}
L331:
	;
	goto L313
L332:
	;
	goto L331
L333:
	;
	v1351 = v1346
	v1353 = v1346
	goto L334
L334:
	;
	v1361 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L1
	} else {
		goto L336
	}
L335:
	;
	goto L131
L336:
	;
	if v1361 == int32(0) {
		goto L131
	} else {
		goto L337
	}
L337:
	;
	F_ProcessPgArchInterrupts(m)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v1368 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[25])) = v1368
	v1372 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[15]))
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+4))
	if v1373 == v1368 {
		goto L343
	} else {
		goto L344
	}
L339:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[19]))
	if v1692 == int32(0) {
		v1351 = v1689
		v1353 = v1690
		goto L334
	} else {
		goto L409
	}
L340:
	;
	F_pg_usleep(m, int32(_a_F_PgArchiverMain_24))
	mBase = m.M
	v1689 = v1351 + int32(1)
	v1690 = v1353
	goto L339
L341:
	;
	F_pg_usleep(m, int32(_a_F_PgArchiverMain_24))
	mBase = m.M
	v1689 = v1351
	v1690 = v1353 + int32(1)
	goto L339
L342:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), v1676, int32(_a_F_PgArchiverMain_25))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L408
	}
L343:
	;
	v1407 = v482 + int32(1296)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+112)) = v1407
	v1410 = v482 + int32(176)
	v1415 = F_pg_snprintf(m, v1410, int32(1024), int32(_a_F_PgArchiverMain_26), v482+int32(112))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L353
	}
L344:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[16]))
	v1378 = m.T0[v1373].(func(*base.Module, int32) int32)(m, v1377)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	if v1378 != 0 {
		goto L343
	} else {
		goto L346
	}
L346:
	;
	v1382 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	if v1382 == int32(0) {
		goto L131
	} else {
		goto L348
	}
L348:
	;
	F_errmsg(m, int32(_a_F_PgArchiverMain_27), int32(0))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[25]))
	if v1392 == int32(0) {
		v1676 = int32(434)
		goto L342
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+128)) = v1392
	F_errdetail_internal(m, int32(_a_F_PgArchiverMain_28), v482+int32(128))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(434), int32(_a_F_PgArchiverMain_25))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	goto L131
L353:
	;
	v1421 = F___fstatat(m, int32(-100), v1410, v482+int32(1200), int32(0))
	mBase = m.M
	goto L355
L354:
	;
	v1481 = v482 + int32(1296)
	v1482 = F_pgarch_archiveXlog(m, v1481)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L1
	} else {
		goto L371
	}
L355:
	;
	if v1421 == int32(0) {
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[21]))
	if v1425 != int32(44) {
		goto L354
	} else {
		goto L357
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+100)) = int32(_a_F_PgArchiverMain_13)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+96)) = v1407
	v1432 = v482 + int32(2368)
	v1437 = F_pg_snprintf(m, v1432, int32(1024), int32(_a_F_PgArchiverMain_14), v482+int32(96))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	v1439 = F_unlink(m, v1432)
	mBase = m.M
	if v1439 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1444 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L1
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	if v1353 < int32(2) {
		goto L341
	} else {
		goto L366
	}
L362:
	;
	if v1444 == int32(0) {
		goto L129
	} else {
		goto L363
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+64)) = v1432
	F_errmsg(m, int32(_a_F_PgArchiverMain_29), v482-int32(-64))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(457), int32(_a_F_PgArchiverMain_25))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	goto L129
L366:
	;
	v1463 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L1
	} else {
		goto L367
	}
L367:
	;
	if v1463 == int32(0) {
		goto L131
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+80)) = v482 + int32(2368)
	F_errmsg(m, int32(_a_F_PgArchiverMain_30), v482+int32(80))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(467), int32(_a_F_PgArchiverMain_25))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	goto L131
L371:
	;
	if v1482 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+36)) = int32(_a_F_PgArchiverMain_13)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+32)) = v1481
	v1488 = v482 + int32(2368)
	v1493 = F_pg_snprintf(m, v1488, int32(1024), int32(_a_F_PgArchiverMain_14), v482+int32(32))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L1
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	v1594 = v482 + int32(1296)
	v1595 = int32(1)
	v1596 = int32(0)
	v1600 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[26]))
	v1601 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v1602 = int32(_a_F_PgArchiverMain_31)
	v1604 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27])) = v1604 + v1595
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+296))
	*(*int32)(unsafe.Add(mBase, uint32(v1600)+296)) = v1608 + v1595
	v1615 = base.AtomicRmwOr32(m, v1596, int32(_a_F_PgArchiverMain_32), v1596)
	goto L395
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+20)) = int32(_a_F_PgArchiverMain_33)
	*(*int32)(unsafe.Add(mBase, uint32(v482)+16)) = v1481
	v1499 = v482 + int32(1344)
	v1504 = F_pg_snprintf(m, v1499, int32(1024), int32(_a_F_PgArchiverMain_14), v482+int32(16))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L1
	} else {
		goto L376
	}
L376:
	;
	v1506 = F_rename(m, v1488, v1499)
	mBase = m.M
	if int32(0) <= v1506 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1528 = v482 + int32(1296)
	v1529 = int32(0)
	v1534 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[26]))
	v1535 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v1536 = int32(_a_F_PgArchiverMain_31)
	v1538 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27]))
	v1539 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27])) = v1538 + v1539
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+296))
	*(*int32)(unsafe.Add(mBase, uint32(v1534)+296)) = v1542 + v1539
	v1549 = base.AtomicRmwOr32(m, v1529, int32(_a_F_PgArchiverMain_32), v1529)
	goto L386
L378:
	;
	v1511 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L379
	}
L379:
	;
	if v1511 == int32(0) {
		goto L377
	} else {
		goto L380
	}
L380:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+4)) = v1499
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v1488
	F_errmsg(m, int32(_a_F_PgArchiverMain_34), v482)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_PgArchiverMain_8), int32(840), int32(_a_F_PgArchiverMain_35))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	goto L377
L384:
	;
	goto L129
L386:
	;
	goto L387
L387:
	;
	v1553 = v1534 + int32(304)
	v1554 = *(*int64)(unsafe.Add(mBase, uint32(v1553)))
	*(*int64)(unsafe.Add(mBase, uint32(v1553))) = v1554 + int64(1)
	goto L389
L389:
	;
	goto L390
L390:
	;
	v1561 = v1534 + int32(312)
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1528)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1561)+40)) = uint8(v1562)
	v1564 = *(*int64)(unsafe.Add(mBase, uint32(v1528)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+32)) = v1564
	v1566 = *(*int64)(unsafe.Add(mBase, uint32(v1528)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+24)) = v1566
	v1568 = *(*int64)(unsafe.Add(mBase, uint32(v1528)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+16)) = v1568
	v1570 = *(*int64)(unsafe.Add(mBase, uint32(v1528)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1561)+8)) = v1570
	v1572 = *(*int64)(unsafe.Add(mBase, uint32(v1528)))
	*(*int64)(unsafe.Add(mBase, uint32(v1561))) = v1572
	goto L392
L392:
	;
	goto L393
L393:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1534+int32(360)))) = v1535
	v1579 = int32(0)
	v1582 = base.AtomicRmwOr32(m, v1579, int32(_a_F_PgArchiverMain_32), v1579)
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+296))
	v1584 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1534)+296)) = v1583 + v1584
	v1587 = int32(_a_F_PgArchiverMain_31)
	v1589 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27])) = v1589 - v1584
	goto L384
L394:
	;
	if v1351 < int32(2) {
		goto L340
	} else {
		goto L404
	}
L395:
	;
	goto L397
L397:
	;
	v1619 = v1600 + int32(368)
	v1620 = *(*int64)(unsafe.Add(mBase, uint32(v1619)))
	*(*int64)(unsafe.Add(mBase, uint32(v1619))) = v1620 + int64(1)
	goto L398
L398:
	;
	goto L400
L400:
	;
	v1627 = v1600 + int32(376)
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1594)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1627)+40)) = uint8(v1628)
	v1630 = *(*int64)(unsafe.Add(mBase, uint32(v1594)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1627)+32)) = v1630
	v1632 = *(*int64)(unsafe.Add(mBase, uint32(v1594)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1627)+24)) = v1632
	v1634 = *(*int64)(unsafe.Add(mBase, uint32(v1594)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1627)+16)) = v1634
	v1636 = *(*int64)(unsafe.Add(mBase, uint32(v1594)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1627)+8)) = v1636
	v1638 = *(*int64)(unsafe.Add(mBase, uint32(v1594)))
	*(*int64)(unsafe.Add(mBase, uint32(v1627))) = v1638
	goto L401
L401:
	;
	goto L403
L403:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1600+int32(424)))) = v1601
	v1645 = int32(0)
	v1648 = base.AtomicRmwOr32(m, v1645, int32(_a_F_PgArchiverMain_32), v1645)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+296))
	v1650 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1600)+296)) = v1649 + v1650
	v1653 = int32(_a_F_PgArchiverMain_31)
	v1655 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[27])) = v1655 - v1650
	goto L394
L404:
	;
	v1663 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	if v1663 == int32(0) {
		goto L131
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v482)+48)) = v1594
	F_errmsg(m, int32(_a_F_PgArchiverMain_36), v482+int32(48))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	v1676 = int32(503)
	goto L342
L408:
	;
	goto L131
L409:
	;
	goto L335
L410:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchiverMain[17]))
	v1711 = F_WaitLatch(m, v1707, int32(25), int32(_a_F_PgArchiverMain_37), int32(83886080))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	if v1711&int32(16) == int32(0) {
		goto L118
	} else {
		goto L412
	}
L412:
	;
	goto L120
L413:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RemovePgTempFilesInDir(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	v7 = m.G0
	v9 = v7 - int32(2080)
	m.G0 = v9
	v11 = F_AllocateDir(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(2080)
	return
L2:
	;
	return
L3:
	;
	v13 = int32(0)
	if v11|base.B2i32(l1 == v13) == v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_RemovePgTempFilesInDir[0]))
	if v19 == int32(44) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v23 = F_ReadDirExtended(m, v11, l0, int32(15))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v23
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_FreeDir(m, v11)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L58
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+19)))
	if v31 != int32(46) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L11
L14:
	;
	v166 = F_ReadDirExtended(m, v11, l0, int32(15))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L56
	}
L15:
	;
	v44 = v26 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	v53 = F_pg_snprintf(m, v9+int32(32), int32(2048), int32(_a_F_RemovePgTempFilesInDir_0), v9+int32(16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L2
	} else {
		goto L20
	}
L16:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v34 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v37 != int32(46) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+21)))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	if l2 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(32)
	F_errmsg(m, v152, v9)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L54
	}
L22:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L53
	}
L23:
	;
	v142 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L51
	}
L24:
	;
	v57 = int32(_a_F_RemovePgTempFilesInDir_1)
	goto L29
L25:
	;
	goto L26
L26:
	;
	v108 = F_get_dirent_type(m, v9+int32(32), v26, int32(0), int32(15))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L43
	}
L27:
	;
	if v95-v96 != 0 {
		goto L23
	} else {
		goto L40
	}
L29:
	;
	goto L30
L30:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v64 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v65 = v44
	v66 = v57
	v67 = int32(9)
	v68 = v64
	goto L35
L32:
	;
	v91 = v57
	v95 = int32(0)
	goto L33
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	goto L27
L34:
	;
	v91 = v86
	v95 = v88
	goto L33
L35:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if base.B2i32(v68 != v70)|base.B2i32(v70 == int32(0)) != 0 {
		v86 = v66
		v88 = v68
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v86 = v80
	v88 = int32(0)
	goto L34
L37:
	;
	v76 = v67 - int32(1)
	if v76 == int32(0) {
		v86 = v66
		v88 = v68
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v79 = int32(1)
	v80 = v66 + v79
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
	if v81 != 0 {
		v65 = v65 + v79
		v66 = v80
		v67 = v76
		v68 = v81
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	goto L26
L41:
	;
	v129 = F_unlink(m, v9+int32(32))
	mBase = m.M
	if int32(0) <= v129 {
		goto L14
	} else {
		goto L48
	}
L42:
	;
	v111 = v9 + int32(32)
	F_RemovePgTempFilesInDir(m, v111, int32(0), int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	switch v108 {
	case 0:
		goto L14
	default:
		goto L41
	case 3:
		goto L42
	}
L44:
	;
	v116 = F_rmdir(m, v111)
	mBase = m.M
	if int32(0) <= v116 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v121 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	if v121 == int32(0) {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	v148 = int32(_a_F_RemovePgTempFilesInDir_2)
	v149 = int32(3426)
	goto L22
L48:
	;
	v134 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if v134 == int32(0) {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v148 = int32(_a_F_RemovePgTempFilesInDir_3)
	v149 = int32(3434)
	goto L22
L51:
	;
	if v142 == int32(0) {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v152 = int32(_a_F_RemovePgTempFilesInDir_4)
	v153 = int32(3440)
	goto L21
L53:
	;
	v152 = v148
	v153 = v149
	goto L21
L54:
	;
	F_errfinish(m, int32(_a_F_RemovePgTempFilesInDir_5), v153, int32(_a_F_RemovePgTempFilesInDir_6))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	goto L14
L56:
	;
	if v166 != 0 {
		v26 = v166
		goto L12
	} else {
		goto L57
	}
L57:
	;
	goto L13
L58:
	;
	goto L1
}
func F__PG_init_pg_trgm(m *base.Module) {
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	F_DefineCustomRealVariable(m, int32(_a_F__PG_init_pg_trgm_0), int32(_a_F__PG_init_pg_trgm_1), int32(_a_F__PG_init_pg_trgm_2), int32(_a_F__PG_init_pg_trgm_3), float64(0.30000001192092896), float64(0), float64(1), int32(6))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_DefineCustomRealVariable(m, int32(_a_F__PG_init_pg_trgm_4), int32(_a_F__PG_init_pg_trgm_5), int32(_a_F__PG_init_pg_trgm_2), int32(_a_F__PG_init_pg_trgm_6), float64(0.6000000238418579), float64(0), float64(1), int32(6))
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_DefineCustomRealVariable(m, int32(_a_F__PG_init_pg_trgm_7), int32(_a_F__PG_init_pg_trgm_8), int32(_a_F__PG_init_pg_trgm_2), int32(_a_F__PG_init_pg_trgm_9), float64(0.5), float64(0), float64(1), int32(6))
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_pg_trgm_10))
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F__PG_init_pgcrypto(m *base.Module) {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	F_DefineCustomEnumVariable(m, int32(_a_F__PG_init_pgcrypto_0), int32(_a_F__PG_init_pgcrypto_1), int32(_a_F__PG_init_pgcrypto_2), int32(_a_F__PG_init_pgcrypto_3), int32(0), int32(_a_F__PG_init_pgcrypto_4), int32(5))
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_pgcrypto_5))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pg_base64_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	return int64(base.Ui64(base.I64_extend_i32_u(l1)*int64(3)) >> (uint(int64(2)) % 64))
}
func F_pg_big5_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v4 {
		v7 = int32(1)
	} else {
		v7 = int32(2)
	}
	return v7
}
func F_pg_blocking_pids(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	v2 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_palloc(m, int32(36))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v22
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v29
	v34 = F_palloc_mul(m, int32(20), v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v39 = F_palloc_mul(m, int32(56), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v39
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v44 = F_palloc_mul(m, int32(4), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v44
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v52 = F_LWLockAcquire(m, v48+int32(512), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v54 = int32(0)
	if v16 == v54 {
		v92 = v54
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v92 != 0 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	goto L7
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[2]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 <= int32(0) {
		v92 = v54
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[3]))
	v72 = int32(0)
	goto L11
L11:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v62+int32(36)+v72<<(uint(int32(2))%32))))
	v83 = v70 + v80*int32(768)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v84 == v16 {
		v92 = v83
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v92 = int32(0)
	goto L8
L13:
	;
	v87 = v72 + int32(1)
	if v87 != v63 {
		v72 = v87
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v101 = F_LWLockAcquire(m, v97+int32(_a_F_pg_blocking_pids_0), int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v367+int32(512))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L61
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v108 = F_LWLockAcquire(m, v104+int32(_a_F_pg_blocking_pids_1), int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v115 = F_LWLockAcquire(m, v111+int32(_a_F_pg_blocking_pids_2), int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v122 = F_LWLockAcquire(m, v118+int32(_a_F_pg_blocking_pids_3), int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v129 = F_LWLockAcquire(m, v125+int32(_a_F_pg_blocking_pids_4), int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v136 = F_LWLockAcquire(m, v132+int32(_a_F_pg_blocking_pids_5), int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v143 = F_LWLockAcquire(m, v139+int32(_a_F_pg_blocking_pids_6), int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v150 = F_LWLockAcquire(m, v146+int32(_a_F_pg_blocking_pids_7), int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v157 = F_LWLockAcquire(m, v153+int32(_a_F_pg_blocking_pids_8), int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v164 = F_LWLockAcquire(m, v160+int32(_a_F_pg_blocking_pids_9), int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v171 = F_LWLockAcquire(m, v167+int32(_a_F_pg_blocking_pids_10), int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v178 = F_LWLockAcquire(m, v174+int32(_a_F_pg_blocking_pids_11), int32(1))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v185 = F_LWLockAcquire(m, v181+int32(_a_F_pg_blocking_pids_12), int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v192 = F_LWLockAcquire(m, v188+int32(_a_F_pg_blocking_pids_13), int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v199 = F_LWLockAcquire(m, v195+int32(_a_F_pg_blocking_pids_14), int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	v206 = F_LWLockAcquire(m, v202+int32(_a_F_pg_blocking_pids_15), int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v92)+364))
	if v208 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v256+int32(_a_F_pg_blocking_pids_15))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L45
	}
L35:
	;
	F_GetSingleProcBlockerStatusData(m, v92, v18)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208)+372))
	if v213 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L38:
	;
	goto L34
L39:
	;
	v217 = v208 + int32(368)
	if v213 == v217 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v224 = v213
	goto L41
L41:
	;
	F_GetSingleProcBlockerStatusData(m, v224-int32(376), v18)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L34
L43:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v238 != v217 {
		v224 = v238
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v262+int32(_a_F_pg_blocking_pids_14))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v268+int32(_a_F_pg_blocking_pids_13))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v274+int32(_a_F_pg_blocking_pids_12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v280+int32(_a_F_pg_blocking_pids_11))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v286+int32(_a_F_pg_blocking_pids_10))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v292+int32(_a_F_pg_blocking_pids_9))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v298+int32(_a_F_pg_blocking_pids_8))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v304+int32(_a_F_pg_blocking_pids_7))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v310+int32(_a_F_pg_blocking_pids_6))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v316+int32(_a_F_pg_blocking_pids_5))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v322+int32(_a_F_pg_blocking_pids_4))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v328+int32(_a_F_pg_blocking_pids_3))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v334+int32(_a_F_pg_blocking_pids_2))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v340+int32(_a_F_pg_blocking_pids_1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_pg_blocking_pids[1]))
	F_LWLockRelease(m, v346+int32(_a_F_pg_blocking_pids_0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L17
L61:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v375 = F_palloc(m, v372<<(uint(int32(3))%32))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v377 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v392 = v2
	v394 = v2
	goto L66
L64:
	;
	v669 = v2
	goto L65
L65:
	;
	v673 = F_construct_array_builtin(m, v375, v669, int32(23))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L114
	}
L66:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v399 = v396 + v394*int32(20)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v403 = v395 + v400*int32(56)
	v404 = int32(0)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	if v405 <= v404 {
		v502 = v404
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v669 = v650
	goto L65
L68:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v399)+12))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+15)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v519<<(uint(int32(2))%32))+uint32(_c_F_pg_blocking_pids[4])))
	goto L95
L69:
	;
	v409 = v405 & int32(3)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v411 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v405) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v418 = v404
	v421 = v411
	v427 = int32(0)
	goto L73
L71:
	;
	v461 = v404
	v464 = v411
	goto L72
L72:
	;
	v476 = v461
	v479 = v464
	v484 = v411
	goto L89
L73:
	;
	v433 = int32(56)
	v435 = v403 + v421*v433
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v435)+40))
	if v442 == v410 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v409 == int32(0) {
		v502 = v453
		goto L68
	} else {
		goto L88
	}
L75:
	;
	v444 = v435
	goto L77
L76:
	;
	v444 = v418
	goto L77
L77:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v435)+96))
	if v445 == v410 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v447 = v435 + v433
	goto L80
L79:
	;
	v447 = v444
	goto L80
L80:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v435)+152))
	if v448 == v410 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v450 = v435 + int32(112)
	goto L83
L82:
	;
	v450 = v447
	goto L83
L83:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v435)+208))
	if v451 == v410 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v453 = v435 + int32(168)
	goto L86
L85:
	;
	v453 = v450
	goto L86
L86:
	;
	v454 = int32(4)
	v455 = v421 + v454
	v457 = v427 + v454
	if v457 != v405&int32(2147483644) {
		v418 = v453
		v421 = v455
		v427 = v457
		goto L73
	} else {
		goto L87
	}
L87:
	;
	goto L74
L88:
	;
	v461 = v453
	v464 = v455
	goto L72
L89:
	;
	v493 = v403 + v479*int32(56)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	if v494 == v410 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v502 = v496
	goto L68
L91:
	;
	v496 = v493
	goto L93
L92:
	;
	v496 = v476
	goto L93
L93:
	;
	v497 = int32(1)
	v500 = v484 + v497
	if v500 != v409 {
		v476 = v496
		v479 = v479 + v497
		v484 = v500
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	if int32(0) < v523 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v526 = int32(2)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v502)+20))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v529+v530<<(uint(v526)%32))))
	v540 = v523
	v541 = int32(0)
	v548 = v392
	goto L99
L97:
	;
	v650 = v392
	goto L98
L98:
	;
	v654 = v394 + int32(1)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v654 < v655 {
		v392 = v650
		v394 = v654
		goto L66
	} else {
		goto L113
	}
L99:
	;
	v553 = v403 + v541*int32(56)
	if v553 == v502 {
		v624 = v540
		v632 = v548
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v650 = v632
	goto L98
L101:
	;
	v636 = v541 + int32(1)
	if v636 < v624 {
		v540 = v624
		v541 = v636
		v548 = v632
		goto L99
	} else {
		goto L112
	}
L102:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v553)+44))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v502)+44))
	if v555 == v556 {
		v624 = v540
		v632 = v548
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v553)+16))
	if v558&v534 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v375+v548<<(uint(int32(3))%32)))) = base.I64_extend_i32_s(v555)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	v624 = v619
	v632 = v548 + int32(1)
	goto L101
L105:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	v561 = int32(0)
	if base.B2i32(v560 == v561)|base.B2i32(int32(base.Ui32(v534)>>(uint(v560)%32))&int32(1) == v561) != 0 {
		v624 = v540
		v632 = v548
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v399)+16))
	if v569 <= int32(0) {
		v624 = v540
		v632 = v548
		goto L101
	} else {
		goto L107
	}
L107:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v553)+40))
	v577 = int32(0)
	goto L108
L108:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v518+v517<<(uint(v526)%32)+v577<<(uint(int32(2))%32))))
	if v592 == v572 {
		goto L104
	} else {
		goto L110
	}
L109:
	;
	v624 = v540
	v632 = v548
	goto L101
L110:
	;
	v595 = v577 + int32(1)
	if v569 != v595 {
		v577 = v595
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	goto L100
L113:
	;
	goto L67
L114:
	;
	return base.I64_extend_i32_u(v673)
}
func F_pg_char_and_wchar_strncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v8 = l0
	v9 = l1
	v10 = l2
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v13 != v14 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	return v13 - v14
L6:
	;
	goto L7
L7:
	;
	if v13 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v20 = int32(1)
	v25 = v10 - v20
	if v25 != 0 {
		v8 = v8 + v20
		v9 = v9 + int32(4)
		v10 = v25
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
}
func F_pg_char_to_encoding_private(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(-1)
	if l0 == int32(0) {
		v123 = v13
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v123
L2:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v16 == int32(0) {
		v123 = v13
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v19) {
		v123 = v13
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = l0
	v23 = v16
	v24 = v11
	goto L5
L5:
	;
	v31 = v23 & int32(255)
	goto L7
L6:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	v64 = int32(_a_F_pg_char_to_encoding_private_0)
	v65 = int32(_a_F_pg_char_to_encoding_private_1)
	goto L15
L7:
	;
	if base.B2i32(base.Ui32(v31-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v31|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v55 = v24
	goto L10
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v56 != 0 {
		v22 = v22 + int32(1)
		v23 = v56
		v24 = v55
		goto L5
	} else {
		goto L14
	}
L11:
	;
	v51 = v23 | int32(32)
	goto L13
L12:
	;
	v51 = v23
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v51)
	v55 = v24 + int32(1)
	goto L10
L14:
	;
	goto L6
L15:
	;
	v77 = v65 + (v64-v65)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78))))
	v80 = v63 - v79
	if v80 != 0 {
		v108 = v80
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v123 = v13
	goto L1
L17:
	;
	v112 = base.B2i32(v108 < int32(0))
	if v108 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L18:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v106 != 0 {
		v108 = v106
		goto L17
	} else {
		goto L26
	}
L20:
	;
	v106 = v104 - v105
	goto L19
L21:
	;
	v89 = v11
	v90 = v78
	goto L22
L22:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v104 = v94
	v105 = v93
	goto L20
L24:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v123 = v107
	goto L1
L27:
	;
	v113 = v77 - int32(8)
	goto L29
L28:
	;
	v113 = v64
	goto L29
L29:
	;
	if v108 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v116 = v65
	goto L32
L31:
	;
	v116 = v77 + int32(8)
	goto L32
L32:
	;
	if base.Ui32(v116) <= base.Ui32(v113) {
		v64 = v113
		v65 = v116
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
}
func F_pg_checksum_block_fallback(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	v8 = m.G0
	v9 = int32(128)
	v10 = v8 - v9
	base.MemoryCopy(m, v10, int32(_a_F_pg_checksum_block_fallback_0), v9)
	v17 = int32(0)
	for {
		v23 = l0 + v17<<(uint(int32(7))%32)
		v27 = int32(0)
		for {
			v32 = int32(2)
			v33 = v27 << (uint(v32) % 32)
			v34 = v10 + v33
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v23)))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			v38 = v36 ^ v37
			v39 = int32(16777619)
			v41 = int32(17)
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v38*v39 ^ int32(base.Ui32(v38)>>(uint(v41)%32))
			v46 = v33 | int32(4)
			v47 = v10 + v46
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v46+v23)))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			v51 = v49 ^ v50
			*(*int32)(unsafe.Add(mBase, uint32(v47))) = v51*v39 ^ int32(base.Ui32(v51)>>(uint(v41)%32))
			v59 = v27 + v32
			if v59 != int32(32) {
				v27 = v59
				continue
			} else {
				break
			}
			break
		}
		v63 = v17 + int32(1)
		if v63 != int32(64) {
			v17 = v63
			continue
		} else {
			break
		}
		break
	}
	v67 = int32(0)
	for {
		v76 = v10 + v67<<(uint(int32(2))%32)
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
		v78 = int32(16777619)
		v80 = int32(17)
		*(*int32)(unsafe.Add(mBase, uint32(v76))) = v77*v78 ^ int32(base.Ui32(v77)>>(uint(v80)%32))
		v84 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v84*v78 ^ int32(base.Ui32(v84)>>(uint(v80)%32))
		v91 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v91*v78 ^ int32(base.Ui32(v91)>>(uint(v80)%32))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v98*v78 ^ int32(base.Ui32(v98)>>(uint(v80)%32))
		v106 = v67 + int32(4)
		if v106 != int32(32) {
			v67 = v106
			continue
		} else {
			break
		}
		break
	}
	v110 = int32(0)
	for {
		v119 = v10 + v110<<(uint(int32(2))%32)
		v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
		v121 = int32(16777619)
		v123 = int32(17)
		*(*int32)(unsafe.Add(mBase, uint32(v119))) = v120*v121 ^ int32(base.Ui32(v120)>>(uint(v123)%32))
		v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v127*v121 ^ int32(base.Ui32(v127)>>(uint(v123)%32))
		v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v134*v121 ^ int32(base.Ui32(v134)>>(uint(v123)%32))
		v141 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = v141*v121 ^ int32(base.Ui32(v141)>>(uint(v123)%32))
		v149 = v110 + int32(4)
		if v149 != int32(32) {
			v110 = v149
			continue
		} else {
			break
		}
		break
	}
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v10)+116))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v10)+112))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+104))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v10)+72))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	return v152 ^ (v153 ^ (v154 ^ (v155 ^ (v156 ^ (v157 ^ (v158 ^ (v159 ^ (v160 ^ (v161 ^ (v162 ^ (v163 ^ (v164 ^ (v165 ^ (v166 ^ (v167 ^ (v168 ^ (v169 ^ (v170 ^ (v171 ^ (v172 ^ (v173 ^ (v174 ^ (v175 ^ (v176 ^ (v177 ^ (v178 ^ (v179 ^ (v180 ^ (v181 ^ (v182 ^ v183))))))))))))))))))))))))))))))
}
func F_pg_checksum_init(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	switch l1 - int32(1) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(-1)
		return int32(0)
	case 1:
		v12 = F_pg_cryptohash_create(m, int32(2))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
			v29 = v12
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	case 2:
		v18 = F_pg_cryptohash_create(m, int32(3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
			v29 = v18
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	case 3:
		v22 = F_pg_cryptohash_create(m, int32(4))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
			v29 = v22
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	case 4:
		v26 = F_pg_cryptohash_create(m, int32(5))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26
			v29 = v26
			if v29 == int32(0) {
				v42 = int32(-1)
				return v42
			} else {
				v33 = int32(0)
				v34 = F_pg_cryptohash_init(m, v29)
				mBase = m.M
				if v33 <= v34 {
					v42 = v33
					return v42
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_pg_cryptohash_free(m, v37)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v42 = int32(-1)
						return v42
					}
				}
			}
		}
	default:
		v42 = int32(0)
		return v42
	}
}
func F_pg_checksum_parse_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	v6 = l0
	v7 = int32(_a_F_pg_checksum_parse_type_0)
	goto L2
L1:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v10 == v11 {
		v33 = v10
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v44 = int32(0)
	goto L1
L4:
	;
	v35 = int32(1)
	if v33 != 0 {
		v6 = v6 + v35
		v7 = v7 + v35
		goto L2
	} else {
		goto L13
	}
L5:
	;
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = v10 | int32(32)
	goto L8
L7:
	;
	v21 = v10
	goto L8
L8:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = v11 | int32(32)
	goto L11
L10:
	;
	v30 = v11
	goto L11
L11:
	;
	if v21 == v30 {
		v33 = v21
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v44 = v21 - v30
	goto L1
L13:
	;
	goto L3
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return int32(1)
L15:
	;
	goto L16
L16:
	;
	v54 = l0
	v55 = int32(_a_F_pg_checksum_parse_type_1)
	goto L18
L17:
	;
	if v92 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v58 == v59 {
		v81 = v58
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v92 = int32(0)
	goto L17
L20:
	;
	v83 = int32(1)
	if v81 != 0 {
		v54 = v54 + v83
		v55 = v55 + v83
		goto L18
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = v58 | int32(32)
	goto L24
L23:
	;
	v69 = v58
	goto L24
L24:
	;
	if base.Ui32((v59-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = v59 | int32(32)
	goto L27
L26:
	;
	v78 = v59
	goto L27
L27:
	;
	if v69 == v78 {
		v81 = v69
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v92 = v69 - v78
	goto L17
L29:
	;
	goto L19
L30:
	;
	v95 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v95
	return v95
L31:
	;
	goto L32
L32:
	;
	v102 = l0
	v103 = int32(_a_F_pg_checksum_parse_type_2)
	goto L34
L33:
	;
	if v140 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v106 == v107 {
		v129 = v106
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v140 = int32(0)
	goto L33
L36:
	;
	v131 = int32(1)
	if v129 != 0 {
		v102 = v102 + v131
		v103 = v103 + v131
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v106-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = v106 | int32(32)
	goto L40
L39:
	;
	v117 = v106
	goto L40
L40:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v126 = v107 | int32(32)
	goto L43
L42:
	;
	v126 = v107
	goto L43
L43:
	;
	if v117 == v126 {
		v129 = v117
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v140 = v117 - v126
	goto L33
L45:
	;
	goto L35
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(2)
	return int32(1)
L47:
	;
	goto L48
L48:
	;
	v150 = l0
	v151 = int32(_a_F_pg_checksum_parse_type_3)
	goto L50
L49:
	;
	if v188 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v154 == v155 {
		v177 = v154
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v188 = int32(0)
	goto L49
L52:
	;
	v179 = int32(1)
	if v177 != 0 {
		v150 = v150 + v179
		v151 = v151 + v179
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if base.Ui32((v154-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v165 = v154 | int32(32)
	goto L56
L55:
	;
	v165 = v154
	goto L56
L56:
	;
	if base.Ui32((v155-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v174 = v155 | int32(32)
	goto L59
L58:
	;
	v174 = v155
	goto L59
L59:
	;
	if v165 == v174 {
		v177 = v165
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v188 = v165 - v174
	goto L49
L61:
	;
	goto L51
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(3)
	return int32(1)
L63:
	;
	goto L64
L64:
	;
	v198 = l0
	v199 = int32(_a_F_pg_checksum_parse_type_4)
	goto L66
L65:
	;
	if v236 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v202 == v203 {
		v225 = v202
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v236 = int32(0)
	goto L65
L68:
	;
	v227 = int32(1)
	if v225 != 0 {
		v198 = v198 + v227
		v199 = v199 + v227
		goto L66
	} else {
		goto L77
	}
L69:
	;
	if base.Ui32((v202-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v213 = v202 | int32(32)
	goto L72
L71:
	;
	v213 = v202
	goto L72
L72:
	;
	if base.Ui32((v203-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v222 = v203 | int32(32)
	goto L75
L74:
	;
	v222 = v203
	goto L75
L75:
	;
	if v213 == v222 {
		v225 = v213
		goto L68
	} else {
		goto L76
	}
L76:
	;
	v236 = v213 - v222
	goto L65
L77:
	;
	goto L67
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(4)
	return int32(1)
L79:
	;
	goto L80
L80:
	;
	v248 = l0
	v249 = int32(_a_F_pg_checksum_parse_type_5)
	goto L82
L81:
	;
	if v286 != 0 {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v252 == v253 {
		v275 = v252
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v286 = int32(0)
	goto L81
L84:
	;
	v277 = int32(1)
	if v275 != 0 {
		v248 = v248 + v277
		v249 = v249 + v277
		goto L82
	} else {
		goto L93
	}
L85:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v263 = v252 | int32(32)
	goto L88
L87:
	;
	v263 = v252
	goto L88
L88:
	;
	if base.Ui32((v253-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v272 = v253 | int32(32)
	goto L91
L90:
	;
	v272 = v253
	goto L91
L91:
	;
	if v263 == v272 {
		v275 = v263
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v286 = v263 - v272
	goto L81
L93:
	;
	goto L83
L94:
	;
	v287 = int32(0)
	goto L96
L95:
	;
	v287 = int32(5)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v287
	return base.B2i32(v286 == int32(0))
}
func F_pg_column_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L6
	} else {
		goto L33
	}
L2:
	;
	switch v35 + int32(2) {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		v100 = v35
		goto L11
	}
L3:
	;
	v16 = F_get_fn_expr_argtype(m, v11, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v35 = v34
	goto L2
L6:
	;
	return int64(0)
L7:
	;
	v20 = F_get_typlen(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v20 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v27 = F_MemoryContextAlloc(m, v25, int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v27
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v20
	v35 = v20
	goto L2
L11:
	;
	m.G0 = v8 + int32(16)
	return base.I64_extend_i32_s(v100)
L12:
	;
	v96 = F_strlen(m, base.I32_wrap_i64(v10))
	mBase = m.M
	v100 = v96 + int32(1)
	goto L11
L13:
	;
	v39 = base.I32_wrap_i64(v10)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v40 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v100 = v94
	goto L11
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v94 = int32(base.Ui32(v86) >> (uint(int32(2)) % 32))
	goto L14
L16:
	;
	v94 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
	goto L14
L17:
	;
	v43 = v39
	v47 = v10
	goto L20
L18:
	;
	v70 = v39
	v71 = v40
	goto L19
L19:
	;
	if v71&int32(1) == int32(0) {
		goto L15
	} else {
		goto L32
	}
L20:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v70 = v65
	v71 = v67
	goto L19
L22:
	;
	if v48 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v67 == int32(1) {
		v43 = v65
		v47 = base.I64_extend_i32_u(v65)
		goto L20
	} else {
		goto L31
	}
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)+6))
	v94 = v53 & int32(1073741823)
	goto L14
L26:
	;
	goto L27
L27:
	;
	if v48&int32(254) != int32(2) {
		v80 = int32(1)
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v47))+2))
	goto L29
L29:
	;
	v63 = F_EOH_get_flat_size(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v94 = v63
	goto L14
L31:
	;
	goto L21
L32:
	;
	v80 = v71
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
	F_errmsg_internal(m, int32(_a_F_pg_column_size_0), v8)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_pg_column_size_1), int32(_a_F_pg_column_size_2), int32(_a_F_pg_column_size_3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_config(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
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
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int64
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = m.G0
	v18 = v16 - int32(1024)
	m.G0 = v18
	v20 = int32(23)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v20
	v24 = F_palloc_mul(m, int32(8), v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_pstrdup(m, int32(_a_F_pg_config_0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27
	v30 = int32(_a_F_pg_config_1)
	goto L8
L5:
	;
	v153 = F_strlen(m, v18)
	mBase = m.M
	v160 = v153 + int32(1)
	goto L38
L6:
	;
	v147 = F_strlen(m, v136)
	mBase = m.M
	goto L5
L8:
	;
	goto L9
L9:
	;
	v37 = int32(1023)
	if (v18^v30)&int32(3) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
	goto L6
L11:
	;
	v121 = v116
	v122 = v117
	v123 = v118
	goto L32
L12:
	;
	if v111 == int32(0) {
		v136 = v109
		v137 = v110
		goto L10
	} else {
		goto L31
	}
L13:
	;
	v109 = v30
	v110 = v18
	v111 = v37
	goto L12
L14:
	;
	goto L15
L15:
	;
	goto L18
L16:
	;
	goto L25
L18:
	;
	goto L19
L19:
	;
	goto L16
L25:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_config[0])))
	if base.B2i32(v80 == int32(0))|int32(0) != 0 {
		v109 = v30
		v110 = v18
		v111 = v37
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v87 = v30
	v88 = v18
	v89 = v37
	goto L27
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v95 = int32(-2139062144)
	if (int32(16843008)-v92|v92)&v95 != v95 {
		v116 = v87
		v117 = v88
		v118 = v89
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v109 = v103
	v110 = v101
	v111 = v105
	goto L12
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v92
	v100 = int32(4)
	v101 = v88 + v100
	v103 = v87 + v100
	v105 = v89 - v100
	if base.Ui32(int32(3)) < base.Ui32(v105) {
		v87 = v103
		v88 = v101
		v89 = v105
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v116 = v109
	v117 = v110
	v118 = v111
	goto L11
L32:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v125)
	if v125 == int32(0) {
		v136 = v121
		v137 = v122
		goto L10
	} else {
		goto L34
	}
L33:
	;
	v136 = v132
	v137 = v130
	goto L10
L34:
	;
	v129 = int32(1)
	v130 = v122 + v129
	v132 = v121 + v129
	v134 = v123 - v129
	if v134 != 0 {
		v121 = v132
		v122 = v130
		v123 = v134
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	if v172 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	goto L36
L38:
	;
	v162 = int32(0)
	if v160 == v162 {
		v172 = v162
		goto L37
	} else {
		goto L40
	}
L39:
	;
	v172 = v167
	goto L37
L40:
	;
	v166 = v160 - int32(1)
	v167 = v18 + v166
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v168 != int32(47) {
		v160 = v166
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v173 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v173)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v175 = F_pstrdup(m, v18)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v175
	v179 = F_pstrdup(m, int32(_a_F_pg_config_2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v179
	F_get_doc_path(m, v18)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v184 = F_pstrdup(m, v18)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v184
	v188 = F_pstrdup(m, int32(_a_F_pg_config_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v188
	F_get_doc_path(m, v18)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v193 = F_pstrdup(m, v18)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v193
	v197 = F_pstrdup(m, int32(_a_F_pg_config_4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v197
	F_make_relative_path(m, v18, int32(_a_F_pg_config_5), int32(_a_F_pg_config_1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v204 = F_pstrdup(m, v18)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v204
	v208 = F_pstrdup(m, int32(_a_F_pg_config_6))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v208
	F_make_relative_path(m, v18, int32(_a_F_pg_config_7), int32(_a_F_pg_config_1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v215 = F_pstrdup(m, v18)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v215
	v219 = F_pstrdup(m, int32(_a_F_pg_config_8))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v219
	F_make_relative_path(m, v18, int32(_a_F_pg_config_9), int32(_a_F_pg_config_1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v226 = F_pstrdup(m, v18)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v226
	v230 = F_pstrdup(m, int32(_a_F_pg_config_10))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v230
	F_make_relative_path(m, v18, int32(_a_F_pg_config_11), int32(_a_F_pg_config_1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v237 = F_pstrdup(m, v18)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v237
	v241 = F_pstrdup(m, int32(_a_F_pg_config_12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v241
	F_get_pkglib_path(m, v18)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v246 = F_pstrdup(m, v18)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v246
	v250 = F_pstrdup(m, int32(_a_F_pg_config_13))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v250
	F_make_relative_path(m, v18, int32(_a_F_pg_config_14), int32(_a_F_pg_config_1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v257 = F_pstrdup(m, v18)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v257
	v261 = F_pstrdup(m, int32(_a_F_pg_config_15))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = v261
	F_make_relative_path(m, v18, int32(_a_F_pg_config_16), int32(_a_F_pg_config_1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v268 = F_pstrdup(m, v18)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v268
	v272 = F_pstrdup(m, int32(_a_F_pg_config_17))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v272
	F_get_share_path(m, v18)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v277 = F_pstrdup(m, v18)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v277
	v281 = F_pstrdup(m, int32(_a_F_pg_config_18))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+88)) = v281
	F_get_etc_path(m, int32(_a_F_pg_config_1), v18)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v287 = F_pstrdup(m, v18)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v287
	v291 = F_pstrdup(m, int32(_a_F_pg_config_19))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v291
	F_get_pkglib_path(m, v18)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v296 = int32(_a_F_pg_config_20)
	v297 = int32(1024)
	v299 = F_pg_ascii_verifystr(m, v18, v297)
	mBase = m.M
	if v299 == v297 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v306 = F_pstrdup(m, v18)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L86
	}
L82:
	;
	goto L81
L83:
	;
	v301 = F_strlen(m, v296)
	mBase = m.M
	goto L82
L84:
	;
	goto L85
L85:
	;
	v304 = F_strlcpy(m, v18+v299, v296, v297-v299)
	mBase = m.M
	goto L82
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v306
	v310 = F_pstrdup(m, int32(_a_F_pg_config_21))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+104)) = v310
	v314 = F_pstrdup(m, int32(_a_F_pg_config_22))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+108)) = v314
	v318 = F_pstrdup(m, int32(_a_F_pg_config_23))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v318
	v322 = F_pstrdup(m, int32(_a_F_pg_config_24))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+116)) = v322
	v326 = F_pstrdup(m, int32(_a_F_pg_config_25))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+120)) = v326
	v330 = F_pstrdup(m, int32(_a_F_pg_config_26))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+124)) = v330
	v334 = F_pstrdup(m, int32(_a_F_pg_config_27))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+128)) = v334
	v338 = F_pstrdup(m, int32(_a_F_pg_config_28))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+132)) = v338
	v342 = F_pstrdup(m, int32(_a_F_pg_config_29))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v342
	v346 = F_pstrdup(m, int32(_a_F_pg_config_30))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = v346
	v350 = F_pstrdup(m, int32(_a_F_pg_config_31))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+144)) = v350
	v354 = F_pstrdup(m, int32(_a_F_pg_config_32))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+148)) = v354
	v358 = F_pstrdup(m, int32(_a_F_pg_config_33))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = v358
	v362 = F_pstrdup(m, int32(_a_F_pg_config_34))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+156)) = v362
	v366 = F_pstrdup(m, int32(_a_F_pg_config_35))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+160)) = v366
	v370 = F_pstrdup(m, int32(_a_F_pg_config_36))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v370
	v374 = F_pstrdup(m, int32(_a_F_pg_config_37))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+168)) = v374
	v378 = F_pstrdup(m, int32(_a_F_pg_config_38))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+172)) = v378
	v382 = F_pstrdup(m, int32(_a_F_pg_config_39))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v382
	v386 = F_pstrdup(m, int32(_a_F_pg_config_40))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+180)) = v386
	m.G0 = v18 + int32(1024)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	if v392 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v394 = int32(0)
	goto L110
L108:
	;
	goto L109
L109:
	;
	m.G0 = v8 + int32(48)
	return int64(0)
L110:
	;
	v399 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v399
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v399
	v403 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v403)
	v407 = v24 + v394<<(uint(int32(3))%32)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v409 = F_cstring_to_text(m, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L109
L112:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v409)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	v414 = F_cstring_to_text(m, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v414)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v418, v419, v8+int32(16), v8+int32(14))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v427 = v394 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	if base.Ui32(v427) < base.Ui32(v428) {
		v394 = v427
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
}
func F_pg_control_checkpoint(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v81 int64
	_ = v81
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v7 = m.G0
	v9 = v7 - int32(304)
	m.G0 = v9
	v14 = F_get_call_result_type(m, l0, int32(0), v9+int32(108))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		if v14 == int32(1) {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[0]))
			v25 = F_LWLockAcquire(m, v21+int32(1152), int32(1))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[1]))
				v31 = F_get_controlfile(m, v28, v9+int32(31))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[0]))
					F_LWLockRelease(m, v34+int32(1152))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
						if v39 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v185 = m.ExcPending
							if v185 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_checkpoint_0), int32(0))
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_checkpoint_1), int32(90), int32(_a_F_pg_control_checkpoint_2))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v31)+40))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v43
							v46 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_control_checkpoint[2])))
							v47 = base.I64_div_u_s(v42, v46)
							v49 = base.I64_div_u_s(int64(4294967296), v46)
							v50 = base.I64_div_u_s(v47, v49)
							*(*uint32)(unsafe.Add(mBase, uint32(v9)+20)) = uint32(v50)
							v53 = v47 - v49*v50
							*(*uint32)(unsafe.Add(mBase, uint32(v9)+24)) = uint32(v53)
							v56 = v9 + int32(32)
							v61 = F_pg_snprintf(m, v56, int32(64), int32(_a_F_pg_control_checkpoint_3), v9+int32(16))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
								v64 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+112)) = uint8(v64)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+144)) = v63
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v31)+40))
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+113)) = uint8(v64)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v67
								v71 = F_cstring_to_text(m, v56)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int64(0)
								} else {
									v73 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+114)) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = base.I64_extend_i32_u(v71)
									v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(v31)+48)))
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+115)) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = v77
									v81 = int64(*(*int32)(unsafe.Add(mBase, uint32(v31)+52)))
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+116)) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+176)) = v81
									v85 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+56)))
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+117)) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v85
									v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+64)))
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+118)) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v9)+192)) = v89
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v31)+72))
									*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v93)
									v96 = int64(base.Ui64(v93) >> (uint(int64(32)) % 64))
									*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v96)
									v99 = F_psprintf(m, int32(_a_F_pg_control_checkpoint_4), v9)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int64(0)
									} else {
										v101 = F_cstring_to_text(m, v99)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int64(0)
										} else {
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+119)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+200)) = base.I64_extend_i32_u(v101)
											v107 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+80)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+120)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+208)) = v107
											v111 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+84)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+121)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+216)) = v111
											v115 = *(*int64)(unsafe.Add(mBase, uint32(v31)+88))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+122)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+224)) = v115
											v119 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+96)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+123)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+232)) = v119
											v123 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+100)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+124)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+240)) = v123
											v127 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+128)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+125)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+248)) = v127
											v131 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+104)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+126)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+256)) = v131
											v135 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+108)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+127)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+264)) = v135
											v139 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+120)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+128)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+272)) = v139
											v143 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+124)))
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+129)) = uint8(v103)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+280)) = v143
											v147 = *(*int64)(unsafe.Add(mBase, uint32(v31)+112))
											v152 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+130)) = uint8(v152)
											*(*int64)(unsafe.Add(mBase, uint32(v9)+288)) = v147*int64(1000000) - int64(946684800000000)
											v155 = *(*int32)(unsafe.Add(mBase, uint32(v9)+108))
											v160 = F_heap_form_tuple(m, v155, v9+int32(144), v9+int32(112))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int64(0)
											} else {
												v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
												v163 = F_HeapTupleHeaderGetDatum(m, v162)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int64(0)
												} else {
													m.G0 = v9 + int32(304)
													return v163
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v172 = m.ExcPending
			if v172 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_checkpoint_5), int32(0))
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_checkpoint_1), int32(82), int32(_a_F_pg_control_checkpoint_2))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_conversion_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_ConversionIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_crypt(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = F_text_to_cstring(m, v9)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v18 = F_text_to_cstring(m, v14)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v21 = F_palloc0(m, int32(128))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v24 = F_px_crypt(m, v16, v18, v21, int32(128))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v16)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int64(0)
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int64(0)
								} else {
									if v24 != 0 {
										v30 = F_cstring_to_text(m, v24)
										mBase = m.M
										v31 = m.ExcPending
										if v31 != 0 {
											return int64(0)
										} else {
											F_pfree(m, v21)
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return int64(0)
											} else {
												v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												if v34 != v9 {
													F_pfree(m, v9)
													mBase = m.M
													v37 = m.ExcPending
													if v37 != 0 {
														return int64(0)
													} else {
														v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														if v38 != v14 {
															F_pfree(m, v14)
															mBase = m.M
															v41 = m.ExcPending
															if v41 != 0 {
																return int64(0)
															} else {
																return base.I64_extend_i32_u(v30)
															}
														} else {
															return base.I64_extend_i32_u(v30)
														}
													}
												} else {
													v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													if v38 != v14 {
														F_pfree(m, v14)
														mBase = m.M
														v41 = m.ExcPending
														if v41 != 0 {
															return int64(0)
														} else {
															return base.I64_extend_i32_u(v30)
														}
													} else {
														return base.I64_extend_i32_u(v30)
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(579))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_pg_crypt_0), int32(0))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_crypt_1), int32(236), int32(_a_F_pg_crypt_2))
													mBase = m.M
													v59 = m.ExcPending
													if v59 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
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
func F_pg_cryptohash_error(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		return int32(_a_F_pg_cryptohash_error_0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v9 == int32(1) {
			v12 = int32(_a_F_pg_cryptohash_error_1)
		} else {
			v12 = int32(_a_F_pg_cryptohash_error_2)
		}
		if v9 == int32(2) {
			v15 = int32(_a_F_pg_cryptohash_error_0)
		} else {
			v15 = v12
		}
		return v15
	}
}
func F_pg_cryptohash_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v7 = m.Env.Pgmem_hash_reset(m, v6)
		mBase = m.M
		return int32(0)
	}
}
func F_pg_database_encoding_character_incrementer(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_database_encoding_character_incrementer[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 == int32(1) {
		v10 = int32(1853)
	} else {
		v10 = int32(1854)
	}
	if v7 == int32(6) {
		v13 = int32(1852)
	} else {
		v13 = v10
	}
	return v13
}
func F_pg_detoast_datum_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v4&int32(3) != 0 {
		v7 = F_detoast_attr(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v7
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(base.Ui32(v12) >> (uint(int32(2)) % 32))
		v15 = F_palloc(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				base.MemoryCopy(m, v15, l0, v14)
			} else {
			}
			return v15
		}
	}
}
func F_pg_euctw2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(142) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	return v87
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v71
	v76 = v18 + int32(1)
	v78 = v14 + int32(4)
	v79 = v15 + v72
	if int32(0) < v79 {
		v13 = v73
		v14 = v78
		v15 = v79
		v18 = v76
		goto L4
	} else {
		goto L18
	}
L8:
	;
	if v19 == int32(0) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L13
	}
L9:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v28 = v24<<(uint(int32(16))%32) | int32(-1912602624)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v33 = v30<<(uint(int32(8))%32) | v28
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v71 = v33 | v35
	v72 = int32(-4)
	v73 = v13 + int32(4)
	goto L7
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v46 = v42<<(uint(int32(8))%32) | int32(_a_F_pg_euctw2wchar_with_len_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v71 = v46 | v48
	v72 = int32(-3)
	v73 = v13 + int32(3)
	goto L7
L13:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v15 == int32(1) {
		v83 = v14
		v87 = v18
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v71 = v19
	v72 = int32(-1)
	v73 = v13 + int32(1)
	goto L7
L17:
	;
	v61 = v19 << (uint(int32(8)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v61
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v71 = v61 | v63
	v72 = int32(-2)
	v73 = v13 + int32(2)
	goto L7
L18:
	;
	v83 = v78
	v87 = v76
	goto L6
}
func F_pg_event_trigger_table_rewrite_reason(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_table_rewrite_reason[0]))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
		if v9 != 0 {
			m.G0 = v5 + int32(16)
			return base.I64_extend_i32_s(v9)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50463299))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_reason_0)
					F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_1), v5)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_2), int32(1661), int32(_a_F_pg_event_trigger_table_rewrite_reason_3))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50463299))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_event_trigger_table_rewrite_reason_0)
				F_errmsg(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_1), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_event_trigger_table_rewrite_reason_2), int32(1661), int32(_a_F_pg_event_trigger_table_rewrite_reason_3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_gb18030_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v2 < int32(0) {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if base.Ui32((v7-int32(48))&int32(255)) < base.Ui32(int32(10)) {
			v14 = int32(4)
		} else {
			v14 = int32(2)
		}
		v16 = v14
	} else {
		v16 = int32(1)
	}
	return v16
}
func F_pg_get_client_encoding(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_client_encoding[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return v3
}
func F_pg_get_expr_ext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		if v11 == int64(0) {
			v14 = int32(2)
		} else {
			v14 = int32(7)
		}
		v15 = F_pg_get_expr_worker(m, v4, v8, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			if v15 != 0 {
				return base.I64_extend_i32_u(v15)
			} else {
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
				return int64(0)
			}
		}
	}
}
func F_pg_get_line_append(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v17 = int32(-1)
	v18 = v3
	v19 = v3
	v20 = v3
	v22 = v3
	goto L4
L1:
	;
	m.G0 = v12 + int32(16)
	return v142
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v122
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v128+v122))) = uint8(v130)
	v142 = v130
	goto L1
L3:
	;
	v142 = int32(1)
	goto L1
L4:
	;
	if v17 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if int32(base.Ui32(v82)>>(uint(int32(5))%32))&int32(1) != 0 {
		v121 = v35
		v122 = v36
		goto L2
	} else {
		goto L43
	}
L6:
	;
	goto L11
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v35 = l1 + int32(4)
	v36 = v29
	v37 = int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v20 == int32(0) {
		v35 = v18
		v36 = v19
		v37 = v22
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, 8)) = uint8(v33)
	v121 = v18
	v122 = v19
	goto L2
L11:
	;
	goto L14
L12:
	;
	goto L5
L13:
	;
	v89 = int32(m.ExcTag)
	v90 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v89 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L14:
	;
	if v37 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L31
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, 4))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(1)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v58 = F_fgets(m, v53+v54, v56-v53, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if v37 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, 4))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(0)
	goto L22
L21:
	;
	goto L22
L22:
	;
	if v58 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v69 = F_strlen(m, v66+v67)
	mBase = m.M
	v70 = v69 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v70
	if v36 < v70 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	goto L15
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v66-int32(1)))))
	if v76 == int32(10) {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_enlargeStringInfo(m, l1, int32(128))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L14
L31:
	;
	goto L12
L32:
	;
	v94 = int32(v90)
	m.G0 = v12
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v12+int32(12) == v100 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	m.ExcPending = 1
	goto L41
L34:
	;
	if v104 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v104 = v102
	goto L37
L36:
	;
	v104 = int32(0)
	goto L37
L37:
	;
	goto L34
L38:
	;
	F___wasm_longjmp(m, v97, v96)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v17 = v104
	v18 = v35
	v19 = v36
	v20 = v96
	v22 = v37
	goto L4
L41:
	;
	return int32(0)
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v112 == v36 {
		v121 = v35
		v122 = v36
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L3
}
func F_pg_get_multixact_members(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int64
	_ = v136
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v12 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		if v14 == int32(0) {
			v17 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = int32(_a_F_pg_get_multixact_members_0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_members[0]))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_members[0])) = v24
				v27 = F_palloc(m, int32(12))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v30 = F_GetMultiXactIdMembers(m, v12, v27, int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v32 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v30
						v38 = F_get_call_result_type(m, l0, v32, v10+int32(24))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							if v38 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_pg_get_multixact_members_1), int32(0))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_get_multixact_members_2), int32(65), int32(_a_F_pg_get_multixact_members_3))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v42
								v44 = F_TupleDescGetAttInMetadata(m, v42)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v27
									*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v44
									*(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_members[0])) = v22
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
									if v57 < v58 {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v57<<(uint(int32(3))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v64
										v69 = F_psprintf(m, int32(_a_F_pg_get_multixact_members_4), v10+int32(16))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v69
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(3))%32))+4))
											v78 = m.G0
											v80 = v78 - int32(16)
											m.G0 = v80
											if base.Ui32(int32(6)) <= base.Ui32(v77) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v80))) = v77
													F_errmsg_internal(m, int32(_a_F_pg_get_multixact_members_5), v80)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_get_multixact_members_6), int32(1591), int32(_a_F_pg_get_multixact_members_7))
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(int32(2))%32))+uint32(_c_F_pg_get_multixact_members[1])))
												m.G0 = v80 + int32(16)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v99
												v104 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
												v107 = F_BuildTupleFromCStrings(m, v104, v10+int32(24))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int64(0)
												} else {
													v109 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v109 + int32(1)
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
													F_pfree(m, v113)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int64(0)
													} else {
														v116 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
														*(*int64)(unsafe.Add(mBase, uint32(v55))) = v116 + int64(1)
														v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = int32(1)
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
														v124 = F_HeapTupleHeaderGetDatum(m, v123)
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int64(0)
														} else {
															v136 = v124
															m.G0 = v10 + int32(32)
															return v136
														}
													}
												}
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int64(0)
										} else {
											v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = int32(2)
											v131 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v131)
											v136 = int64(0)
											m.G0 = v10 + int32(32)
											return v136
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
			if v57 < v58 {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v57<<(uint(int32(3))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v64
				v69 = F_psprintf(m, int32(_a_F_pg_get_multixact_members_4), v10+int32(16))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v69
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(3))%32))+4))
					v78 = m.G0
					v80 = v78 - int32(16)
					m.G0 = v80
					if base.Ui32(int32(6)) <= base.Ui32(v77) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v80))) = v77
							F_errmsg_internal(m, int32(_a_F_pg_get_multixact_members_5), v80)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_get_multixact_members_6), int32(1591), int32(_a_F_pg_get_multixact_members_7))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v77<<(uint(int32(2))%32))+uint32(_c_F_pg_get_multixact_members[1])))
						m.G0 = v80 + int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v99
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
						v107 = F_BuildTupleFromCStrings(m, v104, v10+int32(24))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int64(0)
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v109 + int32(1)
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							F_pfree(m, v113)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int64(0)
							} else {
								v116 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
								*(*int64)(unsafe.Add(mBase, uint32(v55))) = v116 + int64(1)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = int32(1)
								v123 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
								v124 = F_HeapTupleHeaderGetDatum(m, v123)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int64(0)
								} else {
									v136 = v124
									m.G0 = v10 + int32(32)
									return v136
								}
							}
						}
					}
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int64(0)
				} else {
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = int32(2)
					v131 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v131)
					v136 = int64(0)
					m.G0 = v10 + int32(32)
					return v136
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v144 = m.ExcPending
		if v144 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
				F_errmsg(m, int32(_a_F_pg_get_multixact_members_8), v10)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_get_multixact_members_2), int32(48), int32(_a_F_pg_get_multixact_members_3))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_get_partkeydef(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_get_partkeydef_worker(m, v3, int32(2), int32(0), int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int64(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v17)
				}
			}
		}
	}
}
func F_pg_get_publication_tables_b(m *base.Module, l0 int32) int64 {
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
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = int32(1)
		v10 = F_pg_get_publication_tables(m, l0, v3, v7, v8, v8)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return v10
		}
	}
}
func F_pg_get_ruledef(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_get_ruledef_worker(m, v3, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int64(0)
		} else {
			v15 = F_cstring_to_text(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v15)
				}
			}
		}
	}
}
func F_pg_get_timezone_offset(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_get_timezone_offset[0])))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v7 < int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
	return int32(1)
L2:
	;
	v13 = int32(1)
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_pg_get_timezone_offset_0)+v13<<(uint(int32(4))%32))))
	if v6 == v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v24 = v13 + int32(1)
	if v7 != v24 {
		v13 = v24
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L1
}
func F_pg_get_triggerdef(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_get_triggerdef_worker(m, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int64(0)
		} else {
			v15 = F_cstring_to_text(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v15)
				}
			}
		}
	}
}
func F_pg_get_triggerdef_ext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_pg_get_triggerdef_worker(m, v3, base.B2i32(v4 != int64(0)))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int64(0)
		} else {
			v17 = F_cstring_to_text(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v17)
				}
			}
		}
	}
}
func F_pg_get_viewdef_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
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
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	F_initStringInfo(m, v12+int32(28))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[0]))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L120
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L117
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L114
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = int64(81604378650)
	v33 = F_SPI_prepare(m, int32(_a_F_pg_get_viewdef_worker_0), int32(2), v12+int32(48))
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
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = base.I64_extend_i32_u(l0)
	v47 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), int64(572277))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_SPI_keepplan(m, v33)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[0])) = v33
	goto L9
L13:
	;
	v49 = int32(_a_F_pg_get_viewdef_worker_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+46)) = uint16(v49)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v47
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[0]))
	v58 = F_SPI_execute_plan(m, v53, v12+int32(48), v12+int32(46))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v58 != int32(5) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[1]))
	if v63 != int64(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v353 = F_SPI_finish(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L109
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_viewdef_worker[2]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = int32(_a_F_pg_get_viewdef_worker_2)
	v72 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v72 < v74 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v116 = v12 + int32(79)
	v117 = F_SPI_getbinval(m, v70, v68, v114, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L33
	}
L19:
	;
	v114 = v80 + int32(1)
	goto L18
L20:
	;
	v79 = v74
	v80 = v72
	goto L23
L21:
	;
	goto L22
L22:
	;
	v103 = F_SystemAttributeByName(m, v71)
	mBase = m.M
	if v103 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v86 = v68 + v79<<(uint(int32(3))%32) + v80*int32(100)
	v89 = F_namestrcmp(m, v86+int32(32), v71)
	mBase = m.M
	if v89 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+119)))
	if v92 != int32(1) {
		goto L19
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v96 = v80 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v96 < v97 {
		v79 = v97
		v80 = v96
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	goto L24
L30:
	;
	v114 = int32(-9)
	goto L18
L31:
	;
	goto L32
L32:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+74)))
	v114 = v107
	goto L18
L33:
	;
	v119 = int32(_a_F_pg_get_viewdef_worker_3)
	v120 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v120 < v122 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v163 = F_SPI_getbinval(m, v70, v68, v162, v116)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L49
	}
L35:
	;
	v162 = v128 + int32(1)
	goto L34
L36:
	;
	v127 = v122
	v128 = v120
	goto L39
L37:
	;
	goto L38
L38:
	;
	v151 = F_SystemAttributeByName(m, v119)
	mBase = m.M
	if v151 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v134 = v68 + v127<<(uint(int32(3))%32) + v128*int32(100)
	v137 = F_namestrcmp(m, v134+int32(32), v119)
	mBase = m.M
	if v137 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+119)))
	if v140 != int32(1) {
		goto L35
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v144 = v128 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v144 < v145 {
		v127 = v145
		v128 = v144
		goto L39
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	goto L40
L46:
	;
	v162 = int32(-9)
	goto L34
L47:
	;
	goto L48
L48:
	;
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v151)+74)))
	v162 = v155
	goto L34
L49:
	;
	v165 = int32(_a_F_pg_get_viewdef_worker_4)
	v166 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v166 < v168 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v209 = F_SPI_getbinval(m, v70, v68, v208, v116)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L65
	}
L51:
	;
	v208 = v174 + int32(1)
	goto L50
L52:
	;
	v173 = v168
	v174 = v166
	goto L55
L53:
	;
	goto L54
L54:
	;
	v197 = F_SystemAttributeByName(m, v165)
	mBase = m.M
	if v197 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	v180 = v68 + v173<<(uint(int32(3))%32) + v174*int32(100)
	v183 = F_namestrcmp(m, v180+int32(32), v165)
	mBase = m.M
	if v183 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+119)))
	if v186 != int32(1) {
		goto L51
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v190 = v174 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v190 < v191 {
		v173 = v191
		v174 = v190
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	goto L56
L62:
	;
	v208 = int32(-9)
	goto L50
L63:
	;
	goto L64
L64:
	;
	v201 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197)+74)))
	v208 = v201
	goto L50
L65:
	;
	v211 = int32(_a_F_pg_get_viewdef_worker_5)
	v212 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v212 < v214 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v255 = F_SPI_getvalue(m, v70, v68, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L81
	}
L67:
	;
	v254 = v220 + int32(1)
	goto L66
L68:
	;
	v219 = v214
	v220 = v212
	goto L71
L69:
	;
	goto L70
L70:
	;
	v243 = F_SystemAttributeByName(m, v211)
	mBase = m.M
	if v243 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v226 = v68 + v219<<(uint(int32(3))%32) + v220*int32(100)
	v229 = F_namestrcmp(m, v226+int32(32), v211)
	mBase = m.M
	if v229 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+119)))
	if v232 != int32(1) {
		goto L67
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v236 = v220 + int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v236 < v237 {
		v219 = v237
		v220 = v236
		goto L71
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	goto L72
L78:
	;
	v254 = int32(-9)
	goto L66
L79:
	;
	goto L80
L80:
	;
	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+74)))
	v254 = v247
	goto L66
L81:
	;
	v257 = int32(_a_F_pg_get_viewdef_worker_6)
	v258 = int32(0)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v258 < v260 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v301 = F_SPI_getvalue(m, v70, v68, v300)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v300 = v266 + int32(1)
	goto L82
L84:
	;
	v265 = v260
	v266 = v258
	goto L87
L85:
	;
	goto L86
L86:
	;
	v289 = F_SystemAttributeByName(m, v257)
	mBase = m.M
	if v289 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v272 = v68 + v265<<(uint(int32(3))%32) + v266*int32(100)
	v275 = F_namestrcmp(m, v272+int32(32), v257)
	mBase = m.M
	if v275 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+119)))
	if v278 != int32(1) {
		goto L83
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v282 = v266 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v282 < v283 {
		v265 = v283
		v266 = v282
		goto L87
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L88
L94:
	;
	v300 = int32(-9)
	goto L82
L95:
	;
	goto L96
L96:
	;
	v293 = int32(*(*int16)(unsafe.Add(mBase, uint32(v289)+74)))
	v300 = v293
	goto L82
L97:
	;
	v303 = F_stringToNode(m, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v303 == int32(0) {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if base.B2i32(v209 == int64(0))|(base.B2i32(v309 != int32(1))|base.B2i32(v117&int64(255) != int64(49))) != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v318 != int32(60) {
		goto L16
	} else {
		goto L101
	}
L101:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	if v321 != int32(62) {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+2)))
	if v324 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v327 != int32(1) {
		goto L16
	} else {
		goto L104
	}
L104:
	;
	v331 = v12 + int32(28)
	v335 = F_table_open(m, base.I32_wrap_i64(v163), int32(1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v335)+52))
	F_get_query_def(m, v326, v331, int32(0), v337, int32(1), l1, l2, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_appendStringInfoChar(m, v331, int32(59))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_relation_close(m, v335, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L16
L109:
	;
	if v353 != int32(2) {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	m.G0 = v12 + int32(80)
	if v357 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v363 = v358
	goto L113
L112:
	;
	v363 = int32(0)
	goto L113
L113:
	;
	return v363
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_get_viewdef_worker_0)
	F_errmsg_internal(m, int32(_a_F_pg_get_viewdef_worker_7), v12)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_pg_get_viewdef_worker_8), int32(825), int32(_a_F_pg_get_viewdef_worker_9))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_viewdef_worker_10), v12+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_pg_get_viewdef_worker_8), int32(839), int32(_a_F_pg_get_viewdef_worker_9))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_viewdef_worker_11), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_pg_get_viewdef_worker_8), int32(861), int32(_a_F_pg_get_viewdef_worker_9))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_viewdef_wrap(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_pg_get_viewdef_worker(m, v3, int32(7), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int64(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v16)
				}
			}
		}
	}
}
func F_pg_get_wait_events(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int64
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v24
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+34)) = uint8(v30)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+32)) = uint16(v30)
	v35 = v18 * int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_pg_get_wait_events[0])))
	v37 = F_cstring_to_text(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v66 = F_GetWaitEventCustomNames(m, int32(117440512), v9+int32(76))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = base.I64_extend_i32_u(v37)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_pg_get_wait_events[1])))
	v42 = F_cstring_to_text(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = base.I64_extend_i32_u(v42)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_pg_get_wait_events[2])))
	v47 = F_cstring_to_text(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = base.I64_extend_i32_u(v47)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v51, v52, v9+int32(48), v9+int32(32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v60 = v18 + int32(1)
	if v60 != int32(287) {
		v18 = v60
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L4
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	if int32(0) < v68 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v72 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v138 = F_GetWaitEventCustomNames(m, int32(184549376), v9+int32(76))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	v78 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v78
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)) = uint8(v84)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+28)) = uint16(v84)
	v89 = F_cstring_to_text(m, int32(_a_F_pg_get_wait_events_0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = base.I64_extend_i32_u(v89)
	v95 = v66 + v72<<(uint(int32(2))%32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = F_cstring_to_text(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = base.I64_extend_i32_u(v97)
	v102 = v9 + int32(32)
	F_initStringInfo(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v105
	F_appendStringInfo(m, v102, int32(_a_F_pg_get_wait_events_1), v9+int32(16))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v113 = F_cstring_to_text(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = base.I64_extend_i32_u(v113)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v117, v118, v9+int32(48), v9+int32(28))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v126 = v72 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	if v126 < v127 {
		v72 = v126
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	if int32(0) < v140 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v144 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	m.G0 = v9 + int32(80)
	return int64(0)
L27:
	;
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v150
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)) = uint8(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+28)) = uint16(v156)
	v161 = F_cstring_to_text(m, int32(_a_F_pg_get_wait_events_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = base.I64_extend_i32_u(v161)
	v167 = v138 + v144<<(uint(int32(2))%32)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v169 = F_cstring_to_text(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = base.I64_extend_i32_u(v169)
	v174 = v9 + int32(32)
	F_initStringInfo(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v177
	F_appendStringInfo(m, v174, int32(_a_F_pg_get_wait_events_3), v9)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v183 = F_cstring_to_text(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = base.I64_extend_i32_u(v183)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v187, v188, v9+int32(48), v9+int32(28))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v196 = v144 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	if v196 < v197 {
		v144 = v196
		goto L27
	} else {
		goto L35
	}
L35:
	;
	goto L28
}
func F_pg_hmac_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v5 = F_palloc(m, int32(280))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = int32(0)
			base.MemoryFill(m, v5, v9, int32(280))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v9
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
			if base.Ui32(l0) <= base.Ui32(int32(5)) {
				v18 = l0 << (uint(int32(2)) % 32)
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_hmac_create[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_hmac_create[1])))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v21
			} else {
			}
			v24 = F_pg_cryptohash_create(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v24
				if v24 != 0 {
					v36 = v5
					return v36
				} else {
					F___memset(m, v5, int32(0), int32(280))
					mBase = m.M
					F_pfree(m, v5)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v36 = int32(0)
						return v36
					}
				}
			}
		} else {
			v36 = int32(0)
			return v36
		}
	}
}
func F_pg_hmac_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	if l0 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v7 = F_palloc(m, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
				return int32(-1)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != 0 {
					base.MemoryFill(m, v7, int32(0), v17)
				} else {
				}
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v22 = F_pg_cryptohash_final(m, v20, v7, v21)
				mBase = m.M
				if v22 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v55 == int32(0) {
						v70 = int32(_a_F_pg_hmac_final_0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
						if v62 == int32(1) {
							v65 = int32(_a_F_pg_hmac_final_1)
						} else {
							v65 = int32(_a_F_pg_hmac_final_2)
						}
						if v62 == int32(2) {
							v68 = int32(_a_F_pg_hmac_final_0)
						} else {
							v68 = v65
						}
						v70 = v68
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
					F_pfree(m, v7)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						return int32(-1)
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = F_pg_cryptohash_init(m, v25)
					mBase = m.M
					if v26 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v55 == int32(0) {
							v70 = int32(_a_F_pg_hmac_final_0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
							if v62 == int32(1) {
								v65 = int32(_a_F_pg_hmac_final_1)
							} else {
								v65 = int32(_a_F_pg_hmac_final_2)
							}
							if v62 == int32(2) {
								v68 = int32(_a_F_pg_hmac_final_0)
							} else {
								v68 = v65
							}
							v70 = v68
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
						F_pfree(m, v7)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							return int32(-1)
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v33 = F_pg_cryptohash_update(m, v29, l0+int32(152), v32)
						mBase = m.M
						if v33 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if v55 == int32(0) {
								v70 = int32(_a_F_pg_hmac_final_0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
								if v62 == int32(1) {
									v65 = int32(_a_F_pg_hmac_final_1)
								} else {
									v65 = int32(_a_F_pg_hmac_final_2)
								}
								if v62 == int32(2) {
									v68 = int32(_a_F_pg_hmac_final_0)
								} else {
									v68 = v65
								}
								v70 = v68
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
							F_pfree(m, v7)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								return int32(-1)
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v38 = F_pg_cryptohash_update(m, v36, v7, v37)
							mBase = m.M
							if v38 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v55 == int32(0) {
									v70 = int32(_a_F_pg_hmac_final_0)
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									if v62 == int32(1) {
										v65 = int32(_a_F_pg_hmac_final_1)
									} else {
										v65 = int32(_a_F_pg_hmac_final_2)
									}
									if v62 == int32(2) {
										v68 = int32(_a_F_pg_hmac_final_0)
									} else {
										v68 = v65
									}
									v70 = v68
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
								F_pfree(m, v7)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									return int32(-1)
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v42 = F_pg_cryptohash_final(m, v41, l1, l2)
								mBase = m.M
								if int32(0) <= v42 {
									F_pfree(m, v7)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v51 = int32(0)
										return v51
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if v55 == int32(0) {
										v70 = int32(_a_F_pg_hmac_final_0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
										if v62 == int32(1) {
											v65 = int32(_a_F_pg_hmac_final_1)
										} else {
											v65 = int32(_a_F_pg_hmac_final_2)
										}
										if v62 == int32(2) {
											v68 = int32(_a_F_pg_hmac_final_0)
										} else {
											v68 = v65
										}
										v70 = v68
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v70
									F_pfree(m, v7)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										return int32(-1)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v51 = int32(-1)
		return v51
	}
}
func F_pg_hmac_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_pg_cryptohash_free(m, v2)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			F___memset(m, l0, int32(0), int32(280))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_pg_identify_object_as_address(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+60)) = uint32(v12)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+56)) = uint32(v11)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+52)) = uint32(v10)
	v19 = F_get_call_result_type(m, l0, int32(0), v6+int32(-56))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		if v19 == int32(1) {
			v26 = v6 + int32(-12)
			v28 = F_getObjectTypeDescription(m, v26, int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = F_cstring_to_text(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					v32 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v32)
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v30)
					v41 = F_getObjectIdentityParts(m, v26, v6+int32(-16), v6+int32(-20), int32(1))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						if v41 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v45)
							v72 = v45
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v72)
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
							v79 = F_heap_form_tuple(m, v74, v6+int32(-48), v6+int32(-51))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
								v82 = F_HeapTupleHeaderGetDatum(m, v81)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int64(0)
								} else {
									m.G0 = v8 - int32(-64)
									return v82
								}
							}
						} else {
							F_pfree(m, v41)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
								if v50 != 0 {
									v51 = F_strlist_to_textarray(m, v50)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int64(0)
									} else {
										v56 = v51
										v57 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v57)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v56)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
										if v61 != 0 {
											v62 = F_strlist_to_textarray(m, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int64(0)
											} else {
												v67 = v62
												*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_u(v67)
												v72 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v72)
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												v79 = F_heap_form_tuple(m, v74, v6+int32(-48), v6+int32(-51))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
													v82 = F_HeapTupleHeaderGetDatum(m, v81)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 - int32(-64)
														return v82
													}
												}
											}
										} else {
											v65 = F_construct_empty_array(m, int32(25))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												v67 = v65
												*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_u(v67)
												v72 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v72)
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												v79 = F_heap_form_tuple(m, v74, v6+int32(-48), v6+int32(-51))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
													v82 = F_HeapTupleHeaderGetDatum(m, v81)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 - int32(-64)
														return v82
													}
												}
											}
										}
									}
								} else {
									v54 = F_construct_empty_array(m, int32(25))
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										v56 = v54
										v57 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v57)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v56)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
										if v61 != 0 {
											v62 = F_strlist_to_textarray(m, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int64(0)
											} else {
												v67 = v62
												*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_u(v67)
												v72 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v72)
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												v79 = F_heap_form_tuple(m, v74, v6+int32(-48), v6+int32(-51))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
													v82 = F_HeapTupleHeaderGetDatum(m, v81)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 - int32(-64)
														return v82
													}
												}
											}
										} else {
											v65 = F_construct_empty_array(m, int32(25))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												v67 = v65
												*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_u(v67)
												v72 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v72)
												v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												v79 = F_heap_form_tuple(m, v74, v6+int32(-48), v6+int32(-51))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
													v82 = F_HeapTupleHeaderGetDatum(m, v81)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 - int32(-64)
														return v82
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_identify_object_as_address_0), int32(0))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_identify_object_as_address_1), int32(_a_F_pg_identify_object_as_address_2), int32(_a_F_pg_identify_object_as_address_3))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pg_import_system_collations(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(352)
	m.G0 = v17
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = F_superuser(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v17 + int32(352)
	return base.I64_extend_i32_s(v470)
L2:
	;
	v449 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L6
	} else {
		goto L114
	}
L3:
	;
	v431 = F_ClosePipeStream(m, v40)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L113
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L109
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L105
	}
L6:
	;
	return int64(0)
L7:
	;
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = base.I32_wrap_i64(v19)
	v28 = int64(0)
	v31 = F_SearchSysCacheExists(m, int32(38), v19&int64(4294967295), v28, v28, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L101
	}
L11:
	;
	if v31 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v36 = F_palloc(m, int32(1200))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v40 = F_OpenPipeStream(m, int32(_a_F_pg_import_system_collations_0), int32(_a_F_pg_import_system_collations_1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v40 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v47 = F_fgets(m, v17+int32(224), int32(128), v40)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	if v47 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v55 = v2
	v57 = v36
	v58 = int32(100)
	v60 = v2
	v63 = v2
	goto L18
L18:
	;
	v68 = F_strlen(m, v17+int32(224))
	mBase = m.M
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v313 = F_ClosePipeStream(m, v40)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L83
	}
L20:
	;
	v311 = F_fgets(m, v17+int32(224), int32(128), v40)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L81
	}
L21:
	;
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v95)
	v98 = v17 + int32(224)
	v100 = v98
	goto L31
L22:
	;
	v71 = v68 + v17 + int32(223)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v72 == int32(10) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v78 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	if v78 == int32(0) {
		v297 = v55
		v299 = v57
		v300 = v58
		v302 = v60
		v305 = v63
		goto L20
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v17 + int32(224)
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_2), v17+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(890), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v297 = v55
	v299 = v57
	v300 = v58
	v302 = v60
	v305 = v63
	goto L20
L30:
	;
	if base.B2i32(v102 == int32(0)) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v100))))
	if int32(0) < v102 {
		v100 = v100 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	goto L32
L34:
	;
	v113 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v129 = v17 + int32(224)
	v131 = F_pg_get_encoding_from_locale(m, v129, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L41
	}
L37:
	;
	if v113 == int32(0) {
		v297 = v55
		v299 = v57
		v300 = v58
		v302 = v60
		v305 = v63
		goto L20
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v98
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_5), v17-int32(-64))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(715), int32(_a_F_pg_import_system_collations_6))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v297 = v55
	v299 = v57
	v300 = v58
	v302 = v60
	v305 = v63
	goto L20
L41:
	;
	if v131 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v137 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.B2i32(v131 != int32(7))&base.B2i32(base.Ui32(v131) <= base.Ui32(int32(34))) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	if v137 == int32(0) {
		v297 = v55
		v299 = v57
		v300 = v58
		v302 = v60
		v305 = v63
		goto L20
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v129
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_7), v17+int32(32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(722), int32(_a_F_pg_import_system_collations_6))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v297 = v55
	v299 = v57
	v300 = v58
	v302 = v60
	v305 = v63
	goto L20
L49:
	;
	v161 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v131 == int32(0) {
		v297 = v55
		v299 = v57
		v300 = v58
		v302 = v60
		v305 = v63
		goto L20
	} else {
		goto L56
	}
L52:
	;
	if v161 == int32(0) {
		v297 = v55
		v299 = v57
		v300 = v58
		v302 = v60
		v305 = v63
		goto L20
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v17 + int32(224)
	F_errmsg_internal(m, int32(_a_F_pg_import_system_collations_8), v17+int32(48))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(727), int32(_a_F_pg_import_system_collations_6))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v297 = v55
	v299 = v57
	v300 = v58
	v302 = v60
	v305 = v63
	goto L20
L56:
	;
	v181 = v17 + int32(224)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_pg_import_system_collations[0]))
	v184 = int32(99)
	v186 = int32(0)
	v189 = F_get_collation_actual_version(m, v184, v181)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v191 = int32(1)
	v193 = F_CollationCreate(m, v181, v24, v183, v184, int32(1), v131, v181, v181, v186, v186, v189, v191, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v193 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	v199 = v60
	goto L61
L61:
	;
	v201 = v63 + int32(1)
	v207 = v17 + int32(224)
	v209 = int32(0)
	v216 = v17 + int32(96)
	goto L63
L62:
	;
	v199 = v60 + int32(1)
	goto L61
L63:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v221 != int32(46) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v265 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v265)
	if v209 == v265 {
		v297 = v55
		v299 = v57
		v300 = v58
		v302 = v199
		v305 = v201
		goto L20
	} else {
		goto L74
	}
L65:
	;
	if v221 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v229 = v207
	goto L71
L67:
	;
	goto L64
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v221)
	v225 = int32(1)
	v207 = v207 + v225
	v216 = v216 + v225
	goto L63
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v244 = v229 + int32(1)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v250 = int32(255)
	if base.B2i32(base.Ui32((v245&int32(223)-int32(65))&v250) < base.Ui32(int32(26)))|(base.B2i32(v245 == int32(45))|base.B2i32(base.Ui32((v245-int32(48))&v250) < base.Ui32(int32(10)))) != 0 {
		v229 = v244
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v207 = v244
	v209 = int32(1)
	goto L63
L73:
	;
	goto L72
L74:
	;
	if v58 <= v55 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v272 = F_repalloc(m, v57, v58*int32(24))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L78
	}
L76:
	;
	v276 = v57
	v277 = v58
	goto L77
L77:
	;
	v280 = v276 + v55*int32(12)
	v283 = F_pstrdup(m, v17+int32(224))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L79
	}
L78:
	;
	v276 = v272
	v277 = v58 << (uint(int32(1)) % 32)
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v283
	v288 = F_pstrdup(m, v17+int32(96))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+8)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v280)+4)) = v288
	v297 = v55 + int32(1)
	v299 = v276
	v300 = v277
	v302 = v199
	v305 = v201
	goto L20
L81:
	;
	if v311 != 0 {
		v55 = v297
		v57 = v299
		v58 = v300
		v60 = v302
		v63 = v305
		goto L18
	} else {
		goto L82
	}
L82:
	;
	goto L19
L83:
	;
	if int32(2) <= v297 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v305 != 0 {
		v470 = v375
		goto L1
	} else {
		goto L100
	}
L85:
	;
	v329 = int32(0)
	v334 = v302
	goto L91
L86:
	;
	F_pg_qsort(m, v299, v297, int32(12), int32(555))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v321 = int32(1)
	if v297 != v321 {
		v375 = v302
		goto L84
	} else {
		goto L90
	}
L89:
	;
	v324 = v297
	goto L85
L90:
	;
	v324 = v321
	goto L85
L91:
	;
	v342 = v299 + v329*int32(12)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_pg_import_system_collations[0]))
	v347 = int32(99)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v350 = int32(0)
	v353 = F_get_collation_actual_version(m, v347, v349)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L93
	}
L92:
	;
	v375 = v363
	goto L84
L93:
	;
	v355 = int32(1)
	v357 = F_CollationCreate(m, v344, v24, v346, v347, int32(1), v343, v349, v349, v350, v350, v353, v355, v355)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	if v357 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L98
	}
L96:
	;
	v363 = v334
	goto L97
L97:
	;
	v365 = v329 + int32(1)
	if v365 != v324 {
		v329 = v365
		v334 = v363
		goto L91
	} else {
		goto L99
	}
L98:
	;
	v363 = v334 + int32(1)
	goto L97
L99:
	;
	goto L92
L100:
	;
	v441 = v375
	goto L2
L101:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_pg_import_system_collations_9), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(849), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v24
	F_errmsg(m, int32(_a_F_pg_import_system_collations_10), v17+int32(80))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(854), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_pg_import_system_collations_0)
	F_errmsg(m, int32(_a_F_pg_import_system_collations_11), v17)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(878), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v441 = v2
	goto L2
L114:
	;
	if v449 == int32(0) {
		v470 = v441
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_pg_import_system_collations_12), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_import_system_collations_3), int32(969), int32(_a_F_pg_import_system_collations_4))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	v470 = v441
	goto L1
}
func F_pg_index_column_has_property(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			if v5 <= int32(0) {
				v16 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
				return int64(0)
			} else {
				v22 = F_indexam_property(m, l0, v12, int32(0), base.I32_wrap_i64(v6), v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					return v22
				}
			}
		}
	}
}
func F_pg_is_in_recovery(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_is_in_recovery[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_is_in_recovery[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_is_in_recovery[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	return base.I64_extend_i32_u(v14)
}
func F_pg_largeobject_aclcheck_snapshot(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pg_largeobject_aclmask_snapshot(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 == int64(0))
	}
}
func F_pg_last_committed_xact(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pg_last_committed_xact[0]))
	v15 = F_LWLockAcquire(m, v11+int32(_a_F_pg_last_committed_xact_0), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_last_committed_xact[1]))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
		if v21 != 0 {
			v22 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_last_committed_xact[0]))
			F_LWLockRelease(m, v26+int32(_a_F_pg_last_committed_xact_0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v34 = F_get_call_result_type(m, l0, int32(0), v8+int32(8))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					if v34 != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_pg_last_committed_xact_1), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_last_committed_xact_2), int32(443), int32(_a_F_pg_last_committed_xact_3))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						if base.Ui32(v24) <= base.Ui32(int32(2)) {
							v40 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v40)
							v42 = int32(257)
							*(*uint16)(unsafe.Add(mBase, uint32(v8)+12)) = uint16(v42)
						} else {
							v44 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v44)
							*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v23
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v22
							*(*uint16)(unsafe.Add(mBase, uint32(v8)+13)) = uint16(v44)
							*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v24)
						}
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						v57 = F_heap_form_tuple(m, v52, v8+int32(16), v8+int32(12))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
							v60 = F_HeapTupleHeaderGetDatum(m, v59)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int64(0)
							} else {
								m.G0 = v8 + int32(48)
								return v60
							}
						}
					}
				}
			}
		} else {
			F_error_commit_ts_disabled(m)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_pg_lltoa(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	if int64(0) <= l0 {
		v12 = l0
		v13 = int32(0)
	} else {
		v7 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
		v12 = int64(0) - l0
		v13 = int32(1)
	}
	v14 = l1 + v13
	v15 = int32(0)
	if v12 == int64(0) {
		v24 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v24)
		v195 = int32(1)
	} else {
		v31 = int32(1233)
		v36 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(v12))^int32(63))*v31+v31) >> (uint(int32(12)) % 32))
		v39 = *(*int64)(unsafe.Add(mBase, uint32(v36<<(uint(int32(3))%32))+uint32(_c_F_pg_lltoa[0])))
		v41 = v36 + base.B2i32(base.Ui64(v39) <= base.Ui64(v12))
		if base.Ui64(int64(100000000)) <= base.Ui64(v12) {
			v45 = v12
			v48 = v15
			for {
				v54 = v14 + v41 - v48
				v55 = int32(8)
				v58 = base.I64_div_u_s(v45, int64(100000000))
				v62 = base.I32_wrap_i64(v45 + v58*int64(4194967296))
				v64 = base.I32_div_u_s(v62, int32(_a_F_pg_lltoa_0))
				v65 = int32(1)
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-v55))) = uint16(v67)
				v71 = int32(_a_F_pg_lltoa_1)
				v72 = base.I32_div_u_s(v62, v71)
				v73 = int32(100)
				v74 = base.I32_rem_u_s(v72, v73)
				v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(6)))) = uint16(v77)
				v83 = v62 - v72*v71
				v84 = int32(_a_F_pg_lltoa_2)
				v87 = base.I32_div_u_s(v83&v84, v73)
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(4)))) = uint16(v90)
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v83-v87*v73)&v84<<(uint(v65)%32))+uint32(_c_F_pg_lltoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v54-int32(2)))) = uint16(v101)
				v104 = v48 + v55
				if base.Ui64(int64(9999999999999999)) < base.Ui64(v45) {
					v45 = v58
					v48 = v104
					continue
				} else {
					break
				}
				break
			}
			v107 = v58
			v110 = v104
		} else {
			v107 = v12
			v110 = v15
		}
		v116 = base.I32_wrap_i64(v107)
		if base.Ui64(int64(10000)) <= base.Ui64(v107) {
			v120 = v14 + v41 - v110
			v121 = int32(4)
			v124 = base.I32_div_u_s(v116, int32(_a_F_pg_lltoa_1))
			v127 = v116 + v124*int32(-10000)
			v128 = int32(100)
			v129 = base.I32_div_u_s(v127, v128)
			v130 = int32(1)
			v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129<<(uint(v130)%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v120-v121))) = uint16(v132)
			v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v127-v129*v128)<<(uint(v130)%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v120-int32(2)))) = uint16(v141)
			v145 = v124
			v146 = v110 | v121
		} else {
			v145 = v116
			v146 = v110
		}
		if base.Ui32(int32(100)) <= base.Ui32(v145) {
			v154 = int32(2)
			v156 = int32(_a_F_pg_lltoa_2)
			v158 = int32(100)
			v159 = base.I32_div_u_s(v145&v156, v158)
			v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v145-v159*v158)&v156<<(uint(int32(1))%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v41-v146-v154))) = uint16(v167)
			v171 = v159
			v172 = v146 + v154
		} else {
			v171 = v145
			v172 = v146
		}
		if base.Ui32(int32(10)) <= base.Ui32(v171) {
			v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171<<(uint(int32(1))%32))+uint32(_c_F_pg_lltoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v41-v172-int32(2)))) = uint16(v181)
			v195 = v41
		} else {
			v184 = v171 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v184)
			v195 = v41
		}
	}
	v196 = v195 + v13
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v196))) = uint8(v198)
	return v196
}
func F_pg_mbcharcliplen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbcharcliplen[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_c_F_pg_mbcharcliplen[1])))
	if v14 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_report_invalid_encoding_db(m, v19, v36, v20)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L23
	}
L2:
	;
	return v69
L3:
	;
	if l1 <= int32(0) {
		v69 = v4
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if l1 < l2 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v19 = l0
	v20 = l1
	v22 = v4
	v24 = v4
	goto L7
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v25 == int32(0) {
		v69 = v22
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v69 = v45
	goto L2
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbcharcliplen[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30*int32(28))+uint32(_c_F_pg_mbcharcliplen[2])))
	v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, v19)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v20 < v36 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v42 = v24 + int32(1)
	if l2 < v42 {
		v69 = v22
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v45 = v22 + v36
	v46 = v20 - v36
	if int32(0) < v46 {
		v19 = v19 + v36
		v20 = v46
		v22 = v45
		v24 = v42
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v50 = l1
	goto L17
L16:
	;
	v50 = l2
	goto L17
L17:
	;
	if v50 <= int32(0) {
		v69 = v4
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v56 = v4
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v56))))
	if v60 == int32(0) {
		v69 = v56
		goto L2
	} else {
		goto L21
	}
L20:
	;
	v69 = v50
	goto L2
L21:
	;
	v64 = v56 + int32(1)
	if v64 != v50 {
		v56 = v64
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mcv_list_out(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_byteaout(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_pg_mcv_list_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_mcv_list_recv_0), int32(1509), int32(_a_F_pg_mcv_list_recv_1), int32(_a_F_pg_mcv_list_recv_2), int32(_a_F_pg_mcv_list_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_md5_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v9 = F_strlen(m, l0)
	mBase = m.M
	v10 = v9 + l2
	v13 = F_emscripten_builtin_malloc(m, v10+int32(1))
	mBase = m.M
	if v13 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(_a_F_pg_md5_encrypt_0)
		return int32(0)
	} else {
		if v9 != 0 {
			base.MemoryCopy(m, v13, l0, v9)
		} else {
		}
		if l2 != 0 {
			base.MemoryCopy(m, v13+v9, l1, l2)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_pg_md5_encrypt_1)
		v27 = F_pg_md5_hash(m, v13, v10, l3+int32(3), l4)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_emscripten_builtin_free(m, v13)
			mBase = m.M
			return v27
		}
	}
}
func F_pg_next_dst_boundary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	v8 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l6)+260))
	if v23 == v8 {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
		v29 = l6 + v26<<(uint(int32(4))%32)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pg_next_dst_boundary[1])))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
		v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pg_next_dst_boundary[2]))))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36
		v232 = v8
	} else {
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+272)))
		if v39 == int32(1) {
			v42 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
			if v38 < v42 {
				v56 = v42
				v58 = l6 + int32(280)
				if v38 < v56 {
					v68 = v56 - v38
				} else {
					v66 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
					v68 = v38 - v66
				}
				v70 = v68 - int64(1)
				v71 = int64(12622780800)
				v72 = base.I64_rem_s(v70, v71)
				v73 = v70 - v72
				v75 = v73 + v71
				v77 = int64(-12622780800) - v73
				if v38 < v56 {
					v79 = v75
				} else {
					v79 = v77
				}
				v80 = v79 + v38
				*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v80
				v82 = int32(-1)
				if v80 < v56 {
					v232 = v82
				} else {
					v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
					if v89 < v80 {
						v232 = v82
					} else {
						v93 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
						mBase = m.M
						v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
						v95 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						if v38 < v95 {
							v97 = v77
						} else {
							v97 = v75
						}
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v94 + v97
						v232 = v93
					}
				}
			} else {
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+273)))
				if v45 != int32(1) {
					v101 = l6 + int32(280)
					v103 = v23 - int32(1)
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
					if v107 <= v38 {
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v115 = l6 + v112<<(uint(int32(4))%32)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
						v232 = v8
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
						if v120 <= v38 {
							v122 = int32(1)
							if v122 < v103 {
								v126 = v122
								v133 = v103
								for {
									v145 = int32(1)
									v146 = (v126 + v133) >> (uint(v145) % 32)
									v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
									v153 = base.B2i32(v38 < v152)
									if v38 < v152 {
										v154 = v126
									} else {
										v154 = v146 + v145
									}
									if v38 < v152 {
										v155 = v146
									} else {
										v155 = v133
									}
									if v154 < v155 {
										v126 = v154
										v133 = v155
										continue
									} else {
										break
									}
									break
								}
								v157 = v154
							} else {
								v157 = v122
							}
							v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v177 = v157 + l6
							v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
							v181 = int32(4)
							v183 = v176 + v180<<(uint(v181)%32)
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
							v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
							v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v198 = v176 + v195<<(uint(v181)%32)
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
							v232 = v122
						} else {
							v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
							v206 = int32(4)
							v208 = v204 + v205<<(uint(v206)%32)
							v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
							v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
							v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v218 = v204 + v215<<(uint(v206)%32)
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
							v232 = int32(1)
						}
					}
				} else {
					v53 = *(*int64)(unsafe.Add(mBase, uint32(l6+int32(272)+v23<<(uint(int32(3))%32))))
					if v38 <= v53 {
						v101 = l6 + int32(280)
						v103 = v23 - int32(1)
						v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
						if v107 <= v38 {
							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v115 = l6 + v112<<(uint(int32(4))%32)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
							v232 = v8
						} else {
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
							if v120 <= v38 {
								v122 = int32(1)
								if v122 < v103 {
									v126 = v122
									v133 = v103
									for {
										v145 = int32(1)
										v146 = (v126 + v133) >> (uint(v145) % 32)
										v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
										v153 = base.B2i32(v38 < v152)
										if v38 < v152 {
											v154 = v126
										} else {
											v154 = v146 + v145
										}
										if v38 < v152 {
											v155 = v146
										} else {
											v155 = v133
										}
										if v154 < v155 {
											v126 = v154
											v133 = v155
											continue
										} else {
											break
										}
										break
									}
									v157 = v154
								} else {
									v157 = v122
								}
								v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
								v177 = v157 + l6
								v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
								v181 = int32(4)
								v183 = v176 + v180<<(uint(v181)%32)
								v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
								v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
								v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
								v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
								v198 = v176 + v195<<(uint(v181)%32)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
								v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
								v232 = v122
							} else {
								v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
								v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
								v206 = int32(4)
								v208 = v204 + v205<<(uint(v206)%32)
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
								v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
								v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
								v218 = v204 + v215<<(uint(v206)%32)
								v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
								*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
								v232 = int32(1)
							}
						}
					} else {
						v55 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						v56 = v55
						v58 = l6 + int32(280)
						if v38 < v56 {
							v68 = v56 - v38
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
							v68 = v38 - v66
						}
						v70 = v68 - int64(1)
						v71 = int64(12622780800)
						v72 = base.I64_rem_s(v70, v71)
						v73 = v70 - v72
						v75 = v73 + v71
						v77 = int64(-12622780800) - v73
						if v38 < v56 {
							v79 = v75
						} else {
							v79 = v77
						}
						v80 = v79 + v38
						*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v80
						v82 = int32(-1)
						if v80 < v56 {
							v232 = v82
						} else {
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
							if v89 < v80 {
								v232 = v82
							} else {
								v93 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
								mBase = m.M
								v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
								v95 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
								if v38 < v95 {
									v97 = v77
								} else {
									v97 = v75
								}
								*(*int64)(unsafe.Add(mBase, uint32(l3))) = v94 + v97
								v232 = v93
							}
						}
					}
				}
			}
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+273)))
			if v45 != int32(1) {
				v101 = l6 + int32(280)
				v103 = v23 - int32(1)
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
				if v107 <= v38 {
					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
					v115 = l6 + v112<<(uint(int32(4))%32)
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
					v232 = v8
				} else {
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
					if v120 <= v38 {
						v122 = int32(1)
						if v122 < v103 {
							v126 = v122
							v133 = v103
							for {
								v145 = int32(1)
								v146 = (v126 + v133) >> (uint(v145) % 32)
								v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
								v153 = base.B2i32(v38 < v152)
								if v38 < v152 {
									v154 = v126
								} else {
									v154 = v146 + v145
								}
								if v38 < v152 {
									v155 = v146
								} else {
									v155 = v133
								}
								if v154 < v155 {
									v126 = v154
									v133 = v155
									continue
								} else {
									break
								}
								break
							}
							v157 = v154
						} else {
							v157 = v122
						}
						v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
						v177 = v157 + l6
						v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
						v181 = int32(4)
						v183 = v176 + v180<<(uint(v181)%32)
						v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
						v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v198 = v176 + v195<<(uint(v181)%32)
						v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
						v232 = v122
					} else {
						v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
						v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
						v206 = int32(4)
						v208 = v204 + v205<<(uint(v206)%32)
						v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
						v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
						v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
						*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
						v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v218 = v204 + v215<<(uint(v206)%32)
						v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
						v232 = int32(1)
					}
				}
			} else {
				v53 = *(*int64)(unsafe.Add(mBase, uint32(l6+int32(272)+v23<<(uint(int32(3))%32))))
				if v38 <= v53 {
					v101 = l6 + int32(280)
					v103 = v23 - int32(1)
					v107 = *(*int64)(unsafe.Add(mBase, uint32(v101+v103<<(uint(int32(3))%32))))
					if v107 <= v38 {
						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6+v103)+uint32(_c_F_pg_next_dst_boundary[3]))))
						v115 = l6 + v112<<(uint(int32(4))%32)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[1])))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+uint32(_c_F_pg_next_dst_boundary[2]))))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v118
						v232 = v8
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
						if v120 <= v38 {
							v122 = int32(1)
							if v122 < v103 {
								v126 = v122
								v133 = v103
								for {
									v145 = int32(1)
									v146 = (v126 + v133) >> (uint(v145) % 32)
									v152 = *(*int64)(unsafe.Add(mBase, uint32(v101+v146<<(uint(int32(3))%32))))
									v153 = base.B2i32(v38 < v152)
									if v38 < v152 {
										v154 = v126
									} else {
										v154 = v146 + v145
									}
									if v38 < v152 {
										v155 = v146
									} else {
										v155 = v133
									}
									if v154 < v155 {
										v126 = v154
										v133 = v155
										continue
									} else {
										break
									}
									break
								}
								v157 = v154
							} else {
								v157 = v122
							}
							v176 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v177 = v157 + l6
							v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[4]))))
							v181 = int32(4)
							v183 = v176 + v180<<(uint(v181)%32)
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v184
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v186
							v191 = *(*int64)(unsafe.Add(mBase, uint32(v101+v157<<(uint(int32(3))%32))))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v191
							v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v198 = v176 + v195<<(uint(v181)%32)
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v199
							v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v201
							v232 = v122
						} else {
							v204 = l6 + int32(_a_F_pg_next_dst_boundary_0)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[0])))
							v206 = int32(4)
							v208 = v204 + v205<<(uint(v206)%32)
							v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v209
							v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v211
							v213 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v213
							v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+uint32(_c_F_pg_next_dst_boundary[3]))))
							v218 = v204 + v215<<(uint(v206)%32)
							v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v219
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(l5))) = v221
							v232 = int32(1)
						}
					}
				} else {
					v55 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
					v56 = v55
					v58 = l6 + int32(280)
					if v38 < v56 {
						v68 = v56 - v38
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
						v68 = v38 - v66
					}
					v70 = v68 - int64(1)
					v71 = int64(12622780800)
					v72 = base.I64_rem_s(v70, v71)
					v73 = v70 - v72
					v75 = v73 + v71
					v77 = int64(-12622780800) - v73
					if v38 < v56 {
						v79 = v75
					} else {
						v79 = v77
					}
					v80 = v79 + v38
					*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v80
					v82 = int32(-1)
					if v80 < v56 {
						v232 = v82
					} else {
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v58+v23<<(uint(int32(3))%32)-int32(8))))
						if v89 < v80 {
							v232 = v82
						} else {
							v93 = F_pg_next_dst_boundary(m, v21+int32(8), l1, l2, l3, l4, l5, l6)
							mBase = m.M
							v94 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							v95 = *(*int64)(unsafe.Add(mBase, uint32(l6)+280))
							if v38 < v95 {
								v97 = v77
							} else {
								v97 = v75
							}
							*(*int64)(unsafe.Add(mBase, uint32(l3))) = v94 + v97
							v232 = v93
						}
					}
				}
			}
		}
	}
	m.G0 = v21 + int32(16)
	return v232
}
func F_pg_node_tree_out(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_textout(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_pg_parameter_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
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
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = F_superuser_arg(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			v21 = F_convert_GUC_name_for_parameter_acl(m, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_cstring_to_text(m, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v26 = F_SearchSysCache1(m, int32(43), base.I64_extend_i32_u(v23))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 == int32(0) {
							v64 = int64(0)
							F_pfree(m, v21)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v23)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v72 = v64
									m.G0 = v12 + int32(16)
									return base.B2i32(v72 == int64(0))
								}
							}
						} else {
							v35 = F_SysCacheGetAttr(m, int32(43), v26, int32(3), v12+int32(15))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
								if v37 == int32(1) {
									v42 = F_acldefault(m, int32(27), int32(10))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v47 = int32(0)
										v48 = v42
										v51 = F_aclmask(m, v48, l1, int32(10), l2, int32(1))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v53 = int32(0)
											if base.B2i32(v48 == v53)|base.B2i32(v48 == v47) == v53 {
												F_pfree(m, v48)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v26)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														v64 = v51
														F_pfree(m, v21)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v72 = v64
																m.G0 = v12 + int32(16)
																return base.B2i32(v72 == int64(0))
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v26)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													v64 = v51
													F_pfree(m, v21)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v23)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v72 = v64
															m.G0 = v12 + int32(16)
															return base.B2i32(v72 == int64(0))
														}
													}
												}
											}
										}
									}
								} else {
									v44 = base.I32_wrap_i64(v35)
									v45 = F_pg_detoast_datum(m, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int32(0)
									} else {
										v47 = v44
										v48 = v45
										v51 = F_aclmask(m, v48, l1, int32(10), l2, int32(1))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v53 = int32(0)
											if base.B2i32(v48 == v53)|base.B2i32(v48 == v47) == v53 {
												F_pfree(m, v48)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v26)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														v64 = v51
														F_pfree(m, v21)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															F_pfree(m, v23)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																v72 = v64
																m.G0 = v12 + int32(16)
																return base.B2i32(v72 == int64(0))
															}
														}
													}
												}
											} else {
												F_ReleaseCatCache(m, v26)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													v64 = v51
													F_pfree(m, v21)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														F_pfree(m, v23)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															v72 = v64
															m.G0 = v12 + int32(16)
															return base.B2i32(v72 == int64(0))
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
		} else {
			v72 = l2
			m.G0 = v12 + int32(16)
			return base.B2i32(v72 == int64(0))
		}
	}
}
func F_pg_popcount_masked_optimized(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v5 = l0
	v7 = int64(0)
	v8 = int32(0)
	for {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
		v11 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
		v14 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v12&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
		v17 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v15&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		v20 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v18&l1)+uint32(_c_F_pg_popcount_masked_optimized[0]))))
		v24 = v11 + (v14 + (v17 + (v7 + v20)))
		v25 = int32(4)
		v28 = v8 + v25
		if v28 != int32(_a_F_pg_popcount_masked_optimized_0) {
			v5 = v5 + v25
			v7 = v24
			v8 = v28
			continue
		} else {
			break
		}
		break
	}
	return v24
}
func F_pg_popcount_optimized(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	v3 = int64(0)
	v4 = int32(0)
	if l1 == v4 {
		return int64(0)
	} else {
		v12 = l1 & int32(3)
		if base.Ui32(int32(4)) <= base.Ui32(l1) {
			v17 = l0
			v19 = v3
			v22 = v4
			for {
				v23 = int32(4)
				v24 = v17 + v23
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+3)))
				v26 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_pg_popcount_optimized[0]))))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
				v28 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_pg_popcount_optimized[0]))))
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
				v30 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pg_popcount_optimized[0]))))
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v32 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_pg_popcount_optimized[0]))))
				v36 = v26 + (v28 + (v30 + (v19 + v32)))
				v38 = v22 + v23
				if v38 != l1&int32(-4) {
					v17 = v24
					v19 = v36
					v22 = v38
					continue
				} else {
					break
				}
				break
			}
			if v12 == int32(0) {
				v65 = v36
			} else {
				v42 = v24
				v44 = v36
				v49 = v42
				v50 = int32(0)
				v51 = v44
				for {
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_pg_popcount_optimized[0]))))
					v57 = v51 + v56
					v58 = int32(1)
					v61 = v50 + v58
					if v61 != v12 {
						v49 = v49 + v58
						v50 = v61
						v51 = v57
						continue
					} else {
						break
					}
					break
				}
				v65 = v57
			}
		} else {
			v42 = l0
			v44 = v3
			v49 = v42
			v50 = int32(0)
			v51 = v44
			for {
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
				v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v55)+uint32(_c_F_pg_popcount_optimized[0]))))
				v57 = v51 + v56
				v58 = int32(1)
				v61 = v50 + v58
				if v61 != v12 {
					v49 = v49 + v58
					v50 = v61
					v51 = v57
					continue
				} else {
					break
				}
				break
			}
			v65 = v57
		}
		return v65
	}
}
func F_pg_prewarm(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v256 int64
	_ = v256
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v267 int64
	_ = v267
	var v272 int64
	_ = v272
	var v280 int64
	_ = v280
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v314 int64
	_ = v314
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int64
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int64
	_ = v391
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v419 int64
	_ = v419
	var v434 int32
	_ = v434
	var v436 int64
	_ = v436
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v20 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L13
	} else {
		goto L188
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L13
	} else {
		goto L184
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L13
	} else {
		goto L180
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L13
	} else {
		goto L176
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L13
	} else {
		goto L172
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L13
	} else {
		goto L167
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L13
	} else {
		goto L163
	}
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v23 == int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L13
	} else {
		goto L159
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = F_pg_detoast_datum_packed(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v121 == int32(1) {
		goto L5
	} else {
		goto L42
	}
L13:
	;
	return int64(0)
L14:
	;
	v32 = F_text_to_cstring(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v34 = int32(_a_F_pg_prewarm_0)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_prewarm[0])))
	if base.B2i32(v37 == int32(0))|base.B2i32(v37 != v40) != 0 {
		v58 = v37
		v59 = v40
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v60 == int32(0) {
		v120 = v6
		goto L12
	} else {
		goto L23
	}
L17:
	;
	v60 = v58 - v59
	goto L16
L18:
	;
	v43 = v32
	v44 = v34
	goto L19
L19:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v48
		v59 = v47
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v58 = v48
	v59 = v47
	goto L17
L21:
	;
	v51 = int32(1)
	if v48 == v47 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v63 = int32(_a_F_pg_prewarm_1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_prewarm[1])))
	if base.B2i32(v66 == int32(0))|base.B2i32(v66 != v69) != 0 {
		v87 = v66
		v88 = v69
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v87-v88 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	goto L24
L26:
	;
	v72 = v32
	v73 = v63
	goto L27
L27:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v77 == int32(0) {
		v87 = v77
		v88 = v76
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v87 = v77
	v88 = v76
	goto L25
L29:
	;
	v80 = int32(1)
	if v77 == v76 {
		v72 = v72 + v80
		v73 = v73 + v80
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v120 = int32(1)
	goto L12
L32:
	;
	goto L33
L33:
	;
	v93 = int32(_a_F_pg_prewarm_2)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_prewarm[2])))
	if base.B2i32(v96 == int32(0))|base.B2i32(v96 != v99) != 0 {
		v117 = v96
		v118 = v99
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v117-v118 != 0 {
		goto L6
	} else {
		goto L41
	}
L35:
	;
	goto L34
L36:
	;
	v102 = v32
	v103 = v93
	goto L37
L37:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v107
		v118 = v106
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v117 = v107
	v118 = v106
	goto L35
L39:
	;
	v110 = int32(1)
	if v107 == v106 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v120 = v6
	goto L12
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v125 = F_pg_detoast_datum_packed(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	v127 = F_text_to_cstring(m, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v129 = F_forkname_to_number(m, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v131 = F_get_rel_relkind(m, v26)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L47
	}
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[3]))
	v170 = F_pg_class_aclcheck(m, v165, v168, int64(2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L63
	}
L47:
	;
	if v131&int32(-33) == int32(73) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v137 = int32(1)
	v139 = F_IndexGetRelation(m, v26, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v158 = int32(1)
	v160 = F_relation_open(m, v26, v158)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L13
	} else {
		goto L61
	}
L51:
	;
	if v139 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v144 = F_relation_open(m, v26, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_LockRelationOid(m, v139, int32(1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L13
	} else {
		goto L56
	}
L55:
	;
	v588 = v144
	goto L1
L56:
	;
	v150 = F_relation_open(m, v26, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	if v26 == v139 {
		v164 = v150
		v165 = v139
		v166 = v137
		goto L46
	} else {
		goto L58
	}
L58:
	;
	v155 = F_IndexGetRelation(m, v26, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	if v155 == v139 {
		v164 = v150
		v165 = v139
		v166 = int32(0)
		goto L46
	} else {
		goto L60
	}
L60:
	;
	v588 = v150
	goto L1
L61:
	;
	if v26 == int32(0) {
		v588 = v160
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v164 = v160
	v165 = v26
	v166 = v158
	goto L46
L63:
	;
	if v170 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v173 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172)+119)))
	switch v173 - int32(73) {
	case 0, 32:
		v183 = int32(20)
		goto L68
	default:
		goto L69
	case 10:
		goto L73
	case 29:
		goto L70
	case 36:
		goto L71
	case 45:
		goto L72
	}
L65:
	;
	goto L66
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+119)))
	switch v191 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L76
	default:
		goto L77
	}
L67:
	;
	v186 = F_get_rel_name(m, v26)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L74
	}
L68:
	;
	v185 = v183
	goto L67
L69:
	;
	v183 = int32(42)
	goto L68
L70:
	;
	v185 = int32(18)
	goto L67
L71:
	;
	v185 = int32(23)
	goto L67
L72:
	;
	v185 = int32(52)
	goto L67
L73:
	;
	v185 = int32(38)
	goto L67
L74:
	;
	F_aclcheck_error(m, v170, v185, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	goto L66
L76:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	if v219 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v201 + int32(4)
	F_errmsg(m, int32(_a_F_pg_prewarm_3), v18+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	v211 = int32(*(*int8)(unsafe.Add(mBase, uint32(v210)+119)))
	F_errdetail_relkind_not_supported(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(159), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v245 = v219
	goto L85
L84:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v223
	v227 = F_smgropen(m, v18+int32(88), v220)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L86
	}
L85:
	;
	v246 = F_smgrexists(m, v245, v129)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L91
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
	if v231 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v245 = v243
	goto L85
L88:
	;
	v239 = v231
	goto L90
L89:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
	v239 = v237
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+72)) = v239 + int32(1)
	goto L87
L91:
	;
	if v246 == int32(0) {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v250 = F_RelationGetNumberOfBlocksInFork(m, v164, v129)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	v252 = base.I64_extend_i32_u(v250)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v253 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	if base.B2i32(v256 < int64(0))|base.B2i32(v252 <= v256) != 0 {
		goto L3
	} else {
		goto L97
	}
L95:
	;
	v261 = int64(0)
	goto L96
L96:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v262 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v261 = v256
	goto L96
L98:
	;
	if v60 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L99:
	;
	v272 = v252 - int64(1)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v267 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if base.B2i32(v267 < int64(0))|base.B2i32(v252 <= v267) != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v272 = v267
	goto L98
L103:
	;
	F_relation_close(m, v164, int32(1))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L13
	} else {
		goto L154
	}
L104:
	;
	if v272 < v261 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	if v120 != 0 {
		goto L118
	} else {
		goto L119
	}
L107:
	;
	v436 = int64(0)
	goto L103
L108:
	;
	goto L109
L109:
	;
	v280 = v261
	goto L110
L110:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[4]))
	if v295 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v436 = v272 + int64(1) - v261
	goto L103
L112:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L13
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_PrefetchBuffer(m, v18+int32(104), v164, v129, base.I32_wrap_i64(v280))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L13
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	if base.B2i32(v280 == v272) == int32(0) {
		v280 = v280 + int64(1)
		goto L110
	} else {
		goto L117
	}
L117:
	;
	goto L111
L118:
	;
	if v272 < v261 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = base.I32_wrap_i64(v272) + int32(1)
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+104)) = uint32(v261)
	v380 = int32(0)
	v385 = F_read_stream_begin_relation(m, int32(13), v380, v164, v129, int32(3), v18+int32(104), v380)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L13
	} else {
		goto L140
	}
L121:
	;
	v436 = int64(0)
	goto L103
L122:
	;
	goto L123
L123:
	;
	v314 = v261
	goto L124
L124:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[4]))
	if v329 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v436 = v272 + int64(1) - v261
	goto L103
L126:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L13
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	if v332 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v164)+20))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v336
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v338
	v342 = F_smgropen(m, v18+int32(32), v335)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L133
	}
L131:
	;
	v359 = v332
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = int32(_a_F_pg_prewarm_6)
	F_smgrreadv(m, v359, v129, base.I32_wrap_i64(v314), v18+int32(104))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L13
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+12)) = v342
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)+72))
	if v346 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v359 = v358
	goto L132
L135:
	;
	v354 = v346
	goto L137
L136:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v342)+76))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v342)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v348
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v342)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v350
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v342)+72))
	v354 = v352
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342)+72)) = v354 + int32(1)
	goto L134
L138:
	;
	if base.B2i32(v314 == v272) == int32(0) {
		v314 = v314 + int64(1)
		goto L124
	} else {
		goto L139
	}
L139:
	;
	goto L125
L140:
	;
	if v261 <= v272 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v391 = v261
	goto L144
L142:
	;
	v419 = int64(0)
	goto L143
L143:
	;
	F_read_stream_end(m, v385)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L13
	} else {
		goto L153
	}
L144:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prewarm[4]))
	if v406 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v419 = v272 + int64(1) - v261
	goto L143
L146:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L13
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v410 = F_read_stream_next_buffer(m, v385, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	F_ReleaseBuffer(m, v410)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	if v391 != v272 {
		v391 = v391 + int64(1)
		goto L144
	} else {
		goto L152
	}
L152:
	;
	goto L145
L153:
	;
	v436 = v419
	goto L103
L154:
	;
	if v166 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_UnlockRelationOid(m, v165, int32(1))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L13
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	m.G0 = v18 + int32(112)
	return v436
L158:
	;
	goto L157
L159:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_7), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(83), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_8), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(88), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L167:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L13
	} else {
		goto L168
	}
L168:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_9), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L13
	} else {
		goto L169
	}
L169:
	;
	F_errhint(m, int32(_a_F_pg_prewarm_10), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L13
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(102), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L13
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L172:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L13
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(_a_F_pg_prewarm_11), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(108), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L13
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L13
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v127
	F_errmsg(m, int32(_a_F_pg_prewarm_12), v18+int32(80))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L13
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(166), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L13
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L13
	} else {
		goto L181
	}
L181:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v252 - int64(1)
	F_errmsg(m, int32(_a_F_pg_prewarm_13), v18-int32(-64))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L13
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(179), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L13
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L13
	} else {
		goto L185
	}
L185:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v252 - int64(1)
	F_errmsg(m, int32(_a_F_pg_prewarm_14), v18+int32(48))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L13
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(190), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L13
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L13
	} else {
		goto L189
	}
L189:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v588)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v598 + int32(4)
	F_errmsg(m, int32(_a_F_pg_prewarm_15), v18)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L13
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_pg_prewarm_4), int32(147), int32(_a_F_pg_prewarm_5))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L13
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_read_file_all_missing(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_convert_and_check_filename(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = int64(0)
			v16 = F_read_binary_file(m, v10, v12, int64(-1), base.B2i32(v9 != v12))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				if v16 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int64(0)
				} else {
					v24 = int32(4)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					F_pg_verifymbstr(m, v16+v24, int32(base.Ui32(v26)>>(uint(int32(2))%32))-v24)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v16)
					}
				}
			}
		}
	}
}
func F_pg_reload_conf(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reload_conf[0]))
	v5 = F_pgmem_kill(m, v3, int32(1))
	mBase = m.M
	if v5 == int32(0) {
		return int64(1)
	} else {
		v12 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			if v12 != 0 {
				F_errmsg(m, int32(_a_F_pg_reload_conf_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_reload_conf_1), int32(291), int32(_a_F_pg_reload_conf_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						return int64(0)
					}
				}
			} else {
				return int64(0)
			}
		}
	}
}
func F_pg_rotate_logfile(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rotate_logfile[0])))
	if v4 == int32(0) {
		v9 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			if v9 == int32(0) {
				v43 = int64(0)
				return v43
			} else {
				F_errmsg(m, int32(_a_F_pg_rotate_logfile_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_rotate_logfile_1), int32(311), int32(_a_F_pg_rotate_logfile_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return int64(0)
					}
				}
			}
		}
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_rotate_logfile[1])))
		if v28 == int32(1) {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_rotate_logfile[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v32+int32(12)))) = int32(1)
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_rotate_logfile[3]))
			v41 = F_pgmem_kill(m, v39, int32(10))
			mBase = m.M
		} else {
		}
		v43 = int64(1)
		return v43
	}
}
func F_pg_rusage_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	v3 = l0 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+24)) = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+16)) = int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(1)
	v13 = F___syscall_ret(m, int32(0))
	mBase = m.M
	F_gettimeofday(m, l0)
	mBase = m.M
	return
}
func F_pg_saslprep(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v400 int32
	_ = v400
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3
	v16 = F_strlen(m, l0)
	mBase = m.M
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v714
L2:
	;
	F_pfree(m, v126)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L49
	} else {
		goto L201
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = int32(0)
	goto L2
L4:
	;
	v714 = int32(-1)
	goto L1
L5:
	;
	v21 = v16
	v23 = v3
	v25 = l0
	goto L8
L6:
	;
	v119 = v3
	goto L7
L7:
	;
	v126 = F_palloc(m, v119<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L49
	} else {
		goto L50
	}
L8:
	;
	v26 = int32(-2)
	v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25))))
	if int32(0) <= v27 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v106 < int32(0) {
		v714 = v26
		goto L1
	} else {
		goto L47
	}
L10:
	;
	if base.Ui32(v21) < base.Ui32(v51) {
		v714 = v26
		goto L1
	} else {
		goto L23
	}
L11:
	;
	v51 = int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v32 = v27 & int32(255)
	if v32&int32(224) == int32(192) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = int32(2)
	goto L10
L15:
	;
	goto L16
L16:
	;
	if v32&int32(240) == int32(224) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v51 = int32(3)
	goto L10
L18:
	;
	goto L19
L19:
	;
	if v32&int32(248) == int32(240) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v49 = int32(4)
	goto L22
L21:
	;
	v49 = int32(1)
	goto L22
L22:
	;
	v51 = v49
	goto L10
L23:
	;
	v53 = int32(0)
	switch v51 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L29
	case 2:
		goto L30
	case 3:
		goto L31
	default:
		v102 = v53
		goto L25
	}
L24:
	;
	if v102 == int32(0) {
		v714 = v26
		goto L1
	} else {
		goto L45
	}
L25:
	;
	goto L24
L26:
	;
	v102 = base.B2i32(base.Ui32(v94&int32(255)) < base.Ui32(int32(245)))
	goto L25
L27:
	;
	if base.I32_extend8_s(v89) < int32(-62) {
		v102 = v53
		goto L25
	} else {
		goto L44
	}
L28:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v89 = v88
	goto L27
L29:
	;
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+1)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	switch v63 - int32(224) {
	case 0:
		goto L38
	default:
		goto L34
	case 13:
		goto L37
	case 16:
		goto L36
	case 20:
		goto L35
	}
L30:
	;
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+2)))
	if int32(-65) < v59 {
		v102 = v53
		goto L25
	} else {
		goto L33
	}
L31:
	;
	v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+3)))
	if int32(-65) < v56 {
		v102 = v53
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L29
L34:
	;
	if v62 <= int32(-65) {
		v89 = v63
		goto L27
	} else {
		goto L43
	}
L35:
	;
	if int32(-113) < v62 {
		v102 = v53
		goto L25
	} else {
		goto L42
	}
L36:
	;
	if base.Ui32((v62-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v102 = v53
		goto L25
	} else {
		goto L41
	}
L37:
	;
	if int32(-97) < v62 {
		v102 = v53
		goto L25
	} else {
		goto L40
	}
L38:
	;
	v66 = int32(224)
	if base.Ui32(v66) <= base.Ui32((v62-int32(-64))&int32(255)) {
		v94 = v66
		goto L26
	} else {
		goto L39
	}
L39:
	;
	v102 = v53
	goto L25
L40:
	;
	v94 = int32(237)
	goto L26
L41:
	;
	v94 = int32(240)
	goto L26
L42:
	;
	v94 = int32(244)
	goto L26
L43:
	;
	v102 = v53
	goto L25
L44:
	;
	v94 = v89
	goto L26
L45:
	;
	v106 = v23 + int32(1)
	v108 = v21 - v51
	if v108 != 0 {
		v21 = v108
		v23 = v106
		v25 = v51 + v25
		goto L8
	} else {
		goto L46
	}
L46:
	;
	goto L9
L47:
	;
	if base.Ui32(int32(268435454)) < base.Ui32(v106) {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v119 = v106
	goto L7
L49:
	;
	return int32(0)
L50:
	;
	if v126 == int32(0) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	if v119 == int32(0) {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v135 = l0
	v137 = int32(0)
	goto L53
L53:
	;
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135))))
	v149 = v147 & int32(255)
	if int32(0) <= v147 {
		v206 = v149
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v237 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v126+v119<<(uint(int32(2))%32)))) = v237
	v246 = v237
	v248 = v237
	goto L79
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126+v137<<(uint(int32(2))%32)))) = v206
	v208 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135))))
	if int32(0) <= v208 {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	if v149&int32(224) == int32(192) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v198))))
	v206 = v199 | v201&int32(63)
	goto L55
L58:
	;
	v198 = int32(1)
	v199 = v149 << (uint(int32(6)) % 32) & int32(1984)
	goto L57
L59:
	;
	goto L60
L60:
	;
	if v149&int32(240) == int32(224) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v198 = int32(2)
	v199 = v149<<(uint(int32(12))%32)&int32(_a_F_pg_saslprep_0) | v170&int32(63)<<(uint(int32(6))%32)
	goto L57
L62:
	;
	goto L63
L63:
	;
	if v149&int32(248) != int32(240) {
		v206 = int32(-1)
		goto L55
	} else {
		goto L64
	}
L64:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	v187 = int32(63)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+2)))
	v198 = int32(3)
	v199 = v149<<(uint(int32(18))%32)&int32(_a_F_pg_saslprep_1) | v186&v187<<(uint(int32(12))%32) | v192&v187<<(uint(int32(6))%32)
	goto L57
L65:
	;
	v235 = v137 + int32(1)
	if v235 != v119 {
		v135 = v232 + v135
		v137 = v235
		goto L53
	} else {
		goto L78
	}
L66:
	;
	v232 = int32(1)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v213 = v208 & int32(255)
	if v213&int32(224) == int32(192) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v232 = int32(2)
	goto L65
L70:
	;
	goto L71
L71:
	;
	if v213&int32(240) == int32(224) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v232 = int32(3)
	goto L65
L73:
	;
	goto L74
L74:
	;
	if v213&int32(248) == int32(240) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v230 = int32(4)
	goto L77
L76:
	;
	v230 = int32(1)
	goto L77
L77:
	;
	v232 = v230
	goto L65
L78:
	;
	goto L54
L79:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v126+v246<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v256
	if base.Ui32(int32(-12130)) < base.Ui32(v256-int32(_a_F_pg_saslprep_2)) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v299 = v126 + v293<<(uint(int32(2))%32)
	v300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = v300
	if v293 == v300 {
		goto L2
	} else {
		goto L94
	}
L81:
	;
	v295 = v246 + int32(1)
	if v295 != v119 {
		v246 = v295
		v248 = v293
		goto L79
	} else {
		goto L93
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126+v248<<(uint(int32(2))%32)))) = v285
	v293 = v248 + int32(1)
	goto L81
L83:
	;
	v269 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_3), int32(6), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L49
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v256
	if base.Ui32(v256-int32(_a_F_pg_saslprep_5)) <= base.Ui32(int32(-65108)) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if v269 != 0 {
		v285 = int32(32)
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v285 = v256
	goto L82
L89:
	;
	goto L90
L90:
	;
	v280 = int32(8)
	v283 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_6), v280, v280, int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L49
	} else {
		goto L91
	}
L91:
	;
	if v283 != 0 {
		v293 = v248
		goto L81
	} else {
		goto L92
	}
L92:
	;
	v285 = v256
	goto L82
L93:
	;
	goto L80
L94:
	;
	v305 = F_unicode_normalize(m, int32(2), v126)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L49
	} else {
		goto L95
	}
L95:
	;
	if v305 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_pfree(m, v126)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L49
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v293 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L4
L100:
	;
	F_pfree(m, v126)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L49
	} else {
		goto L199
	}
L101:
	;
	F_pfree(m, v126)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L49
	} else {
		goto L197
	}
L102:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v434 != 0 {
		goto L135
	} else {
		goto L136
	}
L103:
	;
	v314 = int32(0)
	goto L104
L104:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v126+v314<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v326
	if base.Ui32(int32(_a_F_pg_saslprep_7)) <= base.Ui32(v326) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v354 = int32(0)
	goto L114
L106:
	;
	v351 = v314 + int32(1)
	if v351 != v293 {
		v314 = v351
		goto L104
	} else {
		goto L113
	}
L107:
	;
	v331 = v12 + int32(12)
	v336 = F_bsearch(m, v331, int32(_a_F_pg_saslprep_8), int32(36), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L49
	} else {
		goto L108
	}
L108:
	;
	if v336 != 0 {
		goto L101
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v326
	if base.Ui32(v326-int32(_a_F_pg_saslprep_9)) <= base.Ui32(int32(-982494)) {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v347 = F_bsearch(m, v331, int32(_a_F_pg_saslprep_10), int32(396), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L49
	} else {
		goto L111
	}
L111:
	;
	if v347 != 0 {
		goto L101
	} else {
		goto L112
	}
L112:
	;
	goto L106
L113:
	;
	goto L105
L114:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v126+v354<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v366
	if base.Ui32(int32(-63808)) < base.Ui32(v366-int32(_a_F_pg_saslprep_11)) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v299-int32(4))))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v388 = int32(0)
	goto L123
L116:
	;
	goto L115
L117:
	;
	v378 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_12), int32(34), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L49
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v381 = v354 + int32(1)
	if v381 != v293 {
		v354 = v381
		goto L114
	} else {
		goto L122
	}
L120:
	;
	if v378 != 0 {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	goto L102
L123:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v126+v388<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v400
	if base.Ui32(int32(-1114046)) < base.Ui32(v400-int32(_a_F_pg_saslprep_13)) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v417 = F_is_code_in_table(m, v386)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L49
	} else {
		goto L131
	}
L125:
	;
	v412 = F_bsearch(m, v12+int32(12), int32(_a_F_pg_saslprep_14), int32(360), int32(8), int32(_a_F_pg_saslprep_4))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L49
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v415 = v388 + int32(1)
	if v415 != v293 {
		v388 = v415
		goto L123
	} else {
		goto L130
	}
L128:
	;
	if v412 != 0 {
		goto L101
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	goto L124
L131:
	;
	if v417 == int32(0) {
		goto L101
	} else {
		goto L132
	}
L132:
	;
	v421 = F_is_code_in_table(m, v385)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L49
	} else {
		goto L133
	}
L133:
	;
	if v421 == int32(0) {
		goto L101
	} else {
		goto L134
	}
L134:
	;
	goto L102
L135:
	;
	v436 = v434
	v440 = v305
	v442 = int32(0)
	goto L138
L136:
	;
	v548 = int32(1)
	goto L137
L137:
	;
	v549 = F_palloc(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L49
	} else {
		goto L164
	}
L138:
	;
	if base.Ui32(v436) <= base.Ui32(int32(127)) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v548 = v534 + int32(1)
	goto L137
L140:
	;
	v509 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(12)))))
	if int32(0) <= v509 {
		goto L151
	} else {
		goto L152
	}
L141:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v436)
	goto L140
L142:
	;
	goto L143
L143:
	;
	if base.Ui32(v436) <= base.Ui32(int32(2047)) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v453 = v436&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v453)
	v458 = int32(base.Ui32(v436)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v458)
	goto L140
L145:
	;
	goto L146
L146:
	;
	if base.Ui32(v436) <= base.Ui32(int32(_a_F_pg_saslprep_15)) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v462 = int32(63)
	v464 = int32(128)
	v465 = v436&v462 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v465)
	v470 = int32(base.Ui32(v436)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v470)
	v477 = int32(base.Ui32(v436)>>(uint(int32(6))%32))&v462 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v477)
	goto L140
L148:
	;
	goto L149
L149:
	;
	v479 = int32(63)
	v481 = int32(128)
	v482 = v436&v479 | v481
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v482)
	v489 = int32(base.Ui32(v436)>>(uint(int32(6))%32))&v479 | v481
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v489)
	v496 = int32(base.Ui32(v436)>>(uint(int32(12))%32))&v479 | v481
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v496)
	v503 = int32(base.Ui32(v436)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v503)
	goto L140
L150:
	;
	v534 = v533 + v442
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v535 != 0 {
		v436 = v535
		v440 = v440 + int32(4)
		v442 = v534
		goto L138
	} else {
		goto L163
	}
L151:
	;
	v533 = int32(1)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v514 = v509 & int32(255)
	if v514&int32(224) == int32(192) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v533 = int32(2)
	goto L150
L155:
	;
	goto L156
L156:
	;
	if v514&int32(240) == int32(224) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v533 = int32(3)
	goto L150
L158:
	;
	goto L159
L159:
	;
	if v514&int32(248) == int32(240) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v531 = int32(4)
	goto L162
L161:
	;
	v531 = int32(1)
	goto L162
L162:
	;
	v533 = v531
	goto L150
L163:
	;
	goto L139
L164:
	;
	if v549 == int32(0) {
		goto L100
	} else {
		goto L165
	}
L165:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v553 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v554 = v549
	v556 = v553
	v558 = v305
	goto L169
L167:
	;
	v652 = v549
	goto L168
L168:
	;
	v661 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v652))) = uint8(v661)
	F_pfree(m, v126)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L49
	} else {
		goto L195
	}
L169:
	;
	if base.Ui32(v556) <= base.Ui32(int32(127)) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v652 = v650
	goto L168
L171:
	;
	v625 = int32(*(*int8)(unsafe.Add(mBase, uint32(v554))))
	if int32(0) <= v625 {
		goto L182
	} else {
		goto L183
	}
L172:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v554))) = uint8(v556)
	goto L171
L173:
	;
	goto L174
L174:
	;
	if base.Ui32(v556) <= base.Ui32(int32(2047)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v571 = v556&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+1)) = uint8(v571)
	v576 = int32(base.Ui32(v556)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v554))) = uint8(v576)
	goto L171
L176:
	;
	goto L177
L177:
	;
	if base.Ui32(v556) <= base.Ui32(int32(_a_F_pg_saslprep_15)) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v580 = int32(63)
	v582 = int32(128)
	v583 = v556&v580 | v582
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+2)) = uint8(v583)
	v588 = int32(base.Ui32(v556)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v554))) = uint8(v588)
	v595 = int32(base.Ui32(v556)>>(uint(int32(6))%32))&v580 | v582
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+1)) = uint8(v595)
	goto L171
L179:
	;
	goto L180
L180:
	;
	v597 = int32(63)
	v599 = int32(128)
	v600 = v556&v597 | v599
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+3)) = uint8(v600)
	v607 = int32(base.Ui32(v556)>>(uint(int32(6))%32))&v597 | v599
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+2)) = uint8(v607)
	v614 = int32(base.Ui32(v556)>>(uint(int32(12))%32))&v597 | v599
	*(*uint8)(unsafe.Add(mBase, uint32(v554)+1)) = uint8(v614)
	v621 = int32(base.Ui32(v556)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v554))) = uint8(v621)
	goto L171
L181:
	;
	v650 = v649 + v554
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v558)+4))
	if v651 != 0 {
		v554 = v650
		v556 = v651
		v558 = v558 + int32(4)
		goto L169
	} else {
		goto L194
	}
L182:
	;
	v649 = int32(1)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v630 = v625 & int32(255)
	if v630&int32(224) == int32(192) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v649 = int32(2)
	goto L181
L186:
	;
	goto L187
L187:
	;
	if v630&int32(240) == int32(224) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v649 = int32(3)
	goto L181
L189:
	;
	goto L190
L190:
	;
	if v630&int32(248) == int32(240) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v647 = int32(4)
	goto L193
L192:
	;
	v647 = int32(1)
	goto L193
L193:
	;
	v649 = v647
	goto L181
L194:
	;
	goto L170
L195:
	;
	F_pfree(m, v305)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L49
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v549
	v714 = v661
	goto L1
L197:
	;
	F_pfree(m, v305)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L49
	} else {
		goto L198
	}
L198:
	;
	v714 = int32(-3)
	goto L1
L199:
	;
	F_pfree(m, v305)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L49
	} else {
		goto L200
	}
L200:
	;
	goto L4
L201:
	;
	v714 = int32(-3)
	goto L1
}
func F_pg_sequence_last_value(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int64
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_init_sequence(m, v10, v8+int32(28), v8+int32(24))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[0]))
		v22 = F_pg_class_aclcheck(m, v10, v20, int64(258))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			if v22 != 0 {
				F_relation_close(m, v24, int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					v64 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
					v68 = int64(0)
					m.G0 = v8 + int32(32)
					return v68
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+118)))
				switch v26 - int32(112) {
				case 0:
					v47 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
						F_UnlockReleaseBuffer(m, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							F_relation_close(m, v24, int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								if v50 == int32(0) {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
									v68 = int64(0)
								} else {
									v68 = v49
								}
								m.G0 = v8 + int32(32)
								return v68
							}
						}
					}
				default:
					v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])))
					if v34 == int32(1) {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[2]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+308))
						v42 = base.B2i32(v40 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])) = uint8(v42)
						v44 = v42
					} else {
						v44 = int32(0)
					}
					if v44 != 0 {
						F_relation_close(m, v24, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v68 = int64(0)
							m.G0 = v8 + int32(32)
							return v68
						}
					} else {
						v47 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int64(0)
						} else {
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
							F_UnlockReleaseBuffer(m, v51)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								F_relation_close(m, v24, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int64(0)
								} else {
									if v50 == int32(0) {
										v64 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
										v68 = int64(0)
									} else {
										v68 = v49
									}
									m.G0 = v8 + int32(32)
									return v68
								}
							}
						}
					}
				case 4:
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+24)))
					if v29 != int32(1) {
						F_relation_close(m, v24, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v68 = int64(0)
							m.G0 = v8 + int32(32)
							return v68
						}
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])))
						if v34 == int32(1) {
							v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[2]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+308))
							v42 = base.B2i32(v40 != int32(2))
							*(*uint8)(unsafe.Add(mBase, _c_F_pg_sequence_last_value[1])) = uint8(v42)
							v44 = v42
						} else {
							v44 = int32(0)
						}
						if v44 != 0 {
							F_relation_close(m, v24, int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int64(0)
							} else {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
								v68 = int64(0)
								m.G0 = v8 + int32(32)
								return v68
							}
						} else {
							v47 = F_read_seq_tuple(m, v24, v8+int32(20), v8)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
								F_UnlockReleaseBuffer(m, v51)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									F_relation_close(m, v24, int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int64(0)
									} else {
										if v50 == int32(0) {
											v64 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
											v68 = int64(0)
										} else {
											v68 = v49
										}
										m.G0 = v8 + int32(32)
										return v68
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
func F_pg_size_pretty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = v10 >> (uint(int64(63)) % 64)
	if base.Ui64(v10^v12-v12) <= base.Ui64(int64(10239)) {
		v63 = int32(_a_F_pg_size_pretty_0)
		v64 = v10
		v67 = int32(_a_F_pg_size_pretty_1)
	} else {
		v20 = base.I64_div_s(v10, int64(512))
		v22 = v20 >> (uint(int64(63)) % 64)
		if base.Ui64(v20^v22-v22) < base.Ui64(int64(20479)) {
			v63 = int32(_a_F_pg_size_pretty_2)
			v64 = v20
			v67 = int32(_a_F_pg_size_pretty_3)
		} else {
			v30 = base.I64_div_s(v10, int64(524288))
			v32 = v30 >> (uint(int64(63)) % 64)
			if base.Ui64(v30^v32-v32) < base.Ui64(int64(20479)) {
				v63 = int32(_a_F_pg_size_pretty_4)
				v64 = v30
				v67 = int32(_a_F_pg_size_pretty_5)
			} else {
				v40 = base.I64_div_s(v10, int64(536870912))
				v42 = v40 >> (uint(int64(63)) % 64)
				if base.Ui64(v40^v42-v42) < base.Ui64(int64(20479)) {
					v63 = int32(_a_F_pg_size_pretty_6)
					v64 = v40
					v67 = int32(_a_F_pg_size_pretty_7)
				} else {
					v50 = base.I64_div_s(v10, int64(549755813888))
					v52 = v50 >> (uint(int64(63)) % 64)
					if base.Ui64(v50^v52-v52) < base.Ui64(int64(20479)) {
						v63 = int32(_a_F_pg_size_pretty_8)
						v64 = v50
						v67 = int32(_a_F_pg_size_pretty_9)
					} else {
						v60 = base.I64_div_s(v10, int64(562949953421312))
						v63 = int32(_a_F_pg_size_pretty_10)
						v64 = v60
						v67 = int32(_a_F_pg_size_pretty_11)
					}
				}
			}
		}
	}
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+8)))
	if v68 == int32(1) {
		v77 = base.I64_div_s(v64>>(uint(int64(63))%64)|int64(1)+v64, int64(2))
		v78 = v77
	} else {
		v78 = v64
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v78
	v82 = v8 + int32(16)
	v85 = F_pg_snprintf(m, v82, int32(64), int32(_a_F_pg_size_pretty_12), v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		return int64(0)
	} else {
		v89 = F_cstring_to_text(m, v82)
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int64(0)
		} else {
			m.G0 = v8 + int32(80)
			return base.I64_extend_i32_u(v89)
		}
	}
}
func F_pg_sjis_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = int32(1)
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v5+int32(95))&int32(255)) < base.Ui32(int32(63)) {
		v12 = v2
	} else {
		v12 = int32(2)
	}
	if int32(0) <= v5 {
		v15 = v2
	} else {
		v15 = v12
	}
	return v15
}
func F_pg_snapshot_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int64
	_ = v161
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = v13 + int32(20)
	v21 = F_strtox_2(m, v16, v18, int32(10), int64(-1))
	mBase = m.M
	goto L1
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v23 != int32(58) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	m.G0 = v13 + int32(48)
	return v161
L3:
	;
	v136 = int64(0)
	v137 = F_errsave_start(m, v15)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L26
	}
L4:
	;
	v30 = F_strtox_2(m, v22+int32(1), v18, int32(10), int64(-1))
	mBase = m.M
	goto L5
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if base.B2i32(base.I32_wrap_i64(v21) == int32(0))|base.B2i32(v35 != int32(58))|(base.B2i32(v30&int64(4294967295) == int64(0))|base.B2i32(base.Ui64(v30) < base.Ui64(v21))) != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v50 = F_makeStringInfo(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int64(0)
L8:
	;
	v54 = int32(24)
	F_appendBinaryStringInfo(m, v50, v13+v54, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v59 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
	F_pfree(m, v50)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L25
	}
L11:
	;
	v67 = v34 + int32(1)
	v73 = int64(0)
	goto L12
L12:
	;
	v78 = F_strtox_2(m, v67, v13+int32(20), int32(10), int64(-1))
	mBase = m.M
	goto L14
L13:
	;
	goto L10
L14:
	;
	if base.B2i32(base.Ui64(v78) < base.Ui64(v73))|base.B2i32(base.Ui64(v78) < base.Ui64(v21))|base.B2i32(base.Ui64(v30) <= base.Ui64(v78)) != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v78 != v73 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v78
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v88 + int32(1)
	F_appendBinaryStringInfo(m, v50, v13+int32(24), int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v98 != int32(44) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	if v98 == int32(0) {
		goto L10
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v105 != 0 {
		v67 = v84 + int32(1)
		v73 = v78
		goto L12
	} else {
		goto L24
	}
L23:
	;
	goto L3
L24:
	;
	goto L13
L25:
	;
	v161 = base.I64_extend_i32_u(v116)
	goto L2
L26:
	;
	if v137 == int32(0) {
		v161 = v136
		goto L2
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_pg_snapshot_in_0)
	F_errmsg(m, int32(_a_F_pg_snapshot_in_1), v13)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errsave_finish(m, v15, int32(_a_F_pg_snapshot_in_2), int32(325), int32(_a_F_pg_snapshot_in_3))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v161 = v136
	goto L2
}
func F_pg_snapshot_xip(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v23 = F_MemoryContextAlloc(m, v19, int32(base.Ui32(v20)>>(uint(int32(2))%32)))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v27 = int32(base.Ui32(v25) >> (uint(int32(2)) % 32))
					if v27 != 0 {
						base.MemoryCopy(m, v23, v13, v27)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v23
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
					v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+4)))
					if base.Ui64(v36) < base.Ui64(v38) {
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v37+base.I32_wrap_i64(v36)<<(uint(int32(3))%32))+24))
						*(*int64)(unsafe.Add(mBase, uint32(v35))) = v36 + int64(1)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(1)
						return v44
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int64(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(2)
							v57 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
		v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
		v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v37)+4)))
		if base.Ui64(v36) < base.Ui64(v38) {
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v37+base.I32_wrap_i64(v36)<<(uint(int32(3))%32))+24))
			*(*int64)(unsafe.Add(mBase, uint32(v35))) = v36 + int64(1)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(1)
			return v44
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int64(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(2)
				v57 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v57)
				return int64(0)
			}
		}
	}
}
func F_pg_snprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = int64(0)
	if l1 != 0 {
		v15 = l0
	} else {
		v15 = v8 + int32(7)
	}
	v16 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v16) {
		v19 = v16
	} else {
		v19 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15 + v19 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v15
	v26 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)) = uint8(v26)
	F_dopr(m, v8+int32(8), l2, l3)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		v35 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v35)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
		m.G0 = v8 + int32(32)
		if v39 != 0 {
			v46 = int32(-1)
		} else {
			v46 = v38 + (v34 - v37)
		}
		return v46
	}
}
func F_pg_stop_making_pinned_objects(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_superuser(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_pg_stop_making_pinned_objects_0)
					F_errmsg(m, int32(_a_F_pg_stop_making_pinned_objects_1), v6)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stop_making_pinned_objects_2), int32(730), int32(_a_F_pg_stop_making_pinned_objects_0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v31 = m.G0
			v33 = v31 - int32(16)
			m.G0 = v33
			v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[0])))
			if v36 != int32(1) {
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[1]))
				v44 = F_LWLockAcquire(m, v40+int32(256), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[2]))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					if base.Ui32(int32(_a_F_pg_stop_making_pinned_objects_3)) <= base.Ui32(v48) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int64(0)
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[2]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(_a_F_pg_stop_making_pinned_objects_4)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v85
							F_errmsg_internal(m, int32(_a_F_pg_stop_making_pinned_objects_5), v33)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_stop_making_pinned_objects_6), int32(633), int32(_a_F_pg_stop_making_pinned_objects_7))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(_a_F_pg_stop_making_pinned_objects_4)
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(0)
						v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stop_making_pinned_objects[1]))
						F_LWLockRelease(m, v58+int32(256))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							m.G0 = v33 + int32(16)
							m.G0 = v6 + int32(16)
							return int64(0)
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_pg_stop_making_pinned_objects_8), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stop_making_pinned_objects_6), int32(626), int32(_a_F_pg_stop_making_pinned_objects_7))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_strerror_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(_a_F_pg_strerror_r_0)
	if base.Ui32(int32(153)) < base.Ui32(l0) {
		v24 = v10
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v453
L2:
	;
	v453 = l1
	goto L1
L3:
	;
	switch l0 - int32(1) {
	case 0:
		v453 = int32(_a_F_pg_strerror_r_1)
		goto L1
	case 1:
		goto L172
	case 2:
		goto L171
	case 3:
		goto L170
	case 4:
		goto L169
	case 5:
		goto L168
	case 6:
		goto L167
	case 7:
		goto L166
	case 8:
		goto L165
	case 9:
		goto L164
	default:
		goto L112
	case 11:
		goto L163
	case 12:
		goto L162
	case 13:
		goto L161
	case 14:
		goto L160
	case 15:
		goto L159
	case 17:
		goto L158
	case 19:
		goto L157
	case 20:
		goto L156
	case 21:
		goto L155
	case 22:
		goto L153
	case 23:
		goto L152
	case 25:
		goto L151
	case 26:
		goto L150
	case 27:
		goto L149
	case 28:
		goto L148
	case 29:
		goto L147
	case 30:
		goto L146
	case 31:
		goto L145
	case 32:
		goto L144
	case 33:
		goto L143
	case 34:
		goto L142
	case 36:
		goto L141
	case 37:
		goto L140
	case 38:
		goto L139
	case 39:
		goto L138
	case 40:
		goto L137
	case 41:
		goto L136
	case 42:
		goto L135
	case 43:
		goto L134
	case 44:
		goto L133
	case 47:
		goto L132
	case 50:
		goto L131
	case 51:
		goto L130
	case 52:
		goto L129
	case 53:
		goto L128
	case 54:
		goto L127
	case 56:
		goto L126
	case 58:
		goto L124
	case 59:
		goto L123
	case 60:
		goto L122
	case 62:
		goto L121
	case 63:
		goto L120
	case 65:
		goto L119
	case 67:
		goto L118
	case 68:
		goto L117
	case 70:
		goto L116
	case 72:
		goto L115
	case 73:
		goto L114
	case 74:
		goto L113
	case 137:
		goto L125
	case 141:
		goto L154
	}
L4:
	;
	if v373|base.B2i32(l1 == int32(0)) != 0 {
		goto L3
	} else {
		goto L109
	}
L5:
	;
	v25 = F_strlen(m, v24)
	mBase = m.M
	if base.Ui32(int32(256)) <= base.Ui32(v25) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	if l0 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v24 = v20 + int32(_a_F_pg_strerror_r_2)
	goto L5
L8:
	;
	v20 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_strerror_r[0]))))
	if v17 == int32(0) {
		v24 = v10
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v20 = v17
	goto L7
L12:
	;
	goto L17
L13:
	;
	goto L14
L14:
	;
	v202 = v25 + int32(1)
	if base.Ui32(int32(512)) <= base.Ui32(v202) {
		goto L63
	} else {
		goto L64
	}
L15:
	;
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+255)) = uint8(v198)
	v373 = int32(68)
	goto L4
L17:
	;
	goto L18
L18:
	;
	v35 = l1 + int32(255)
	if (l1^v24)&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if base.Ui32(v167) < base.Ui32(v35) {
		goto L56
	} else {
		goto L57
	}
L23:
	;
	if l1&int32(3) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v35) < base.Ui32(int32(4)) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v71 = v35 & int32(-4)
	if base.Ui32(v35) < base.Ui32(int32(64)) {
		v121 = v65
		v122 = v66
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v65 = v24
	v66 = l1
	goto L26
L28:
	;
	goto L29
L29:
	;
	goto L31
L31:
	;
	goto L32
L32:
	;
	v48 = v24
	v49 = l1
	goto L33
L33:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v53)
	v55 = int32(1)
	v56 = v48 + v55
	v58 = v49 + v55
	if v58&int32(3) == int32(0) {
		v65 = v56
		v66 = v58
		goto L26
	} else {
		goto L35
	}
L34:
	;
	v65 = v56
	v66 = v58
	goto L26
L35:
	;
	if base.Ui32(v58) < base.Ui32(v35) {
		v48 = v56
		v49 = v58
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if base.Ui32(v71) <= base.Ui32(v122) {
		v166 = v121
		v167 = v122
		goto L22
	} else {
		goto L43
	}
L38:
	;
	v75 = v71 + int32(-64)
	if base.Ui32(v75) < base.Ui32(v66) {
		v121 = v65
		v122 = v66
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v78 = v65
	v79 = v66
	goto L40
L40:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v78)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v78)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+36)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v78)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+40)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+44)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v78)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+48)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+52)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v78)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+56)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v78)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+60)) = v113
	v115 = int32(-64)
	v116 = v78 - v115
	v118 = v79 - v115
	if base.Ui32(v118) <= base.Ui32(v75) {
		v78 = v116
		v79 = v118
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v121 = v116
	v122 = v118
	goto L37
L42:
	;
	goto L41
L43:
	;
	v128 = v121
	v129 = v122
	goto L44
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v133
	v135 = int32(4)
	v136 = v128 + v135
	v138 = v129 + v135
	if base.Ui32(v138) < base.Ui32(v71) {
		v128 = v136
		v129 = v138
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v166 = v136
	v167 = v138
	goto L22
L46:
	;
	goto L45
L47:
	;
	v166 = v24
	v167 = l1
	goto L22
L48:
	;
	goto L49
L49:
	;
	goto L51
L51:
	;
	goto L52
L52:
	;
	v147 = v24
	v148 = l1
	goto L53
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v152)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)) = uint8(v154)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+2)) = uint8(v156)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+3)) = uint8(v158)
	v160 = int32(4)
	v161 = v147 + v160
	v163 = v148 + v160
	if base.Ui32(v163) <= base.Ui32(v35-int32(4)) {
		v147 = v161
		v148 = v163
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v166 = v161
	v167 = v163
	goto L22
L55:
	;
	goto L54
L56:
	;
	v173 = v166
	v174 = v167
	goto L59
L57:
	;
	goto L58
L58:
	;
	goto L15
L59:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v178)
	v180 = int32(1)
	v183 = v174 + v180
	if v183 != v35 {
		v173 = v173 + v180
		v174 = v183
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	goto L60
L62:
	;
	v373 = int32(0)
	goto L4
L63:
	;
	if v202 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v209 = l1 + v202
	if (l1^v24)&int32(3) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	base.MemoryCopy(m, l1, v24, v202)
	goto L68
L67:
	;
	goto L68
L68:
	;
	goto L62
L69:
	;
	if base.Ui32(v341) < base.Ui32(v209) {
		goto L103
	} else {
		goto L104
	}
L70:
	;
	if l1&int32(3) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(v209) < base.Ui32(int32(4)) {
		goto L94
	} else {
		goto L95
	}
L73:
	;
	v245 = v209 & int32(-4)
	if base.Ui32(v209) < base.Ui32(int32(64)) {
		v295 = v239
		v296 = v240
		goto L84
	} else {
		goto L85
	}
L74:
	;
	v239 = v24
	v240 = l1
	goto L73
L75:
	;
	goto L76
L76:
	;
	if v202 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v239 = v24
	v240 = l1
	goto L73
L78:
	;
	goto L79
L79:
	;
	v222 = v24
	v223 = l1
	goto L80
L80:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v227)
	v229 = int32(1)
	v230 = v222 + v229
	v232 = v223 + v229
	if v232&int32(3) == int32(0) {
		v239 = v230
		v240 = v232
		goto L73
	} else {
		goto L82
	}
L81:
	;
	v239 = v230
	v240 = v232
	goto L73
L82:
	;
	if base.Ui32(v232) < base.Ui32(v209) {
		v222 = v230
		v223 = v232
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if base.Ui32(v245) <= base.Ui32(v296) {
		v340 = v295
		v341 = v296
		goto L69
	} else {
		goto L90
	}
L85:
	;
	v249 = v245 + int32(-64)
	if base.Ui32(v249) < base.Ui32(v240) {
		v295 = v239
		v296 = v240
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v252 = v239
	v253 = v240
	goto L87
L87:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+12)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v252)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+16)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+20)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v252)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = v269
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v252)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+28)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v252)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+32)) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v252)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+36)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v252)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+40)) = v277
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+44)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v252)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+48)) = v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v252)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+52)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v252)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+56)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v252)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+60)) = v287
	v289 = int32(-64)
	v290 = v252 - v289
	v292 = v253 - v289
	if base.Ui32(v292) <= base.Ui32(v249) {
		v252 = v290
		v253 = v292
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v295 = v290
	v296 = v292
	goto L84
L89:
	;
	goto L88
L90:
	;
	v302 = v295
	v303 = v296
	goto L91
L91:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = v307
	v309 = int32(4)
	v310 = v302 + v309
	v312 = v303 + v309
	if base.Ui32(v312) < base.Ui32(v245) {
		v302 = v310
		v303 = v312
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v340 = v310
	v341 = v312
	goto L69
L93:
	;
	goto L92
L94:
	;
	v340 = v24
	v341 = l1
	goto L69
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(v202) < base.Ui32(int32(4)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v340 = v24
	v341 = l1
	goto L69
L98:
	;
	goto L99
L99:
	;
	v321 = v24
	v322 = l1
	goto L100
L100:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	*(*uint8)(unsafe.Add(mBase, uint32(v322))) = uint8(v326)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+1)) = uint8(v328)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+2)) = uint8(v330)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v322)+3)) = uint8(v332)
	v334 = int32(4)
	v335 = v321 + v334
	v337 = v322 + v334
	if base.Ui32(v337) <= base.Ui32(v209-int32(4)) {
		v321 = v335
		v322 = v337
		goto L100
	} else {
		goto L102
	}
L101:
	;
	v340 = v335
	v341 = v337
	goto L69
L102:
	;
	goto L101
L103:
	;
	v347 = v340
	v348 = v341
	goto L106
L104:
	;
	goto L105
L105:
	;
	goto L62
L106:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	*(*uint8)(unsafe.Add(mBase, uint32(v348))) = uint8(v352)
	v354 = int32(1)
	v357 = v348 + v354
	if v357 != v209 {
		v347 = v347 + v354
		v348 = v357
		goto L106
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	goto L107
L109:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v377 == int32(63) {
		goto L3
	} else {
		goto L110
	}
L110:
	;
	if v377 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	goto L3
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	v447 = F_pg_snprintf(m, l1, int32(256), int32(_a_F_pg_strerror_r_3), v8)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L173
	} else {
		goto L174
	}
L113:
	;
	v453 = int32(_a_F_pg_strerror_r_4)
	goto L1
L114:
	;
	v453 = int32(_a_F_pg_strerror_r_5)
	goto L1
L115:
	;
	v453 = int32(_a_F_pg_strerror_r_6)
	goto L1
L116:
	;
	v453 = int32(_a_F_pg_strerror_r_7)
	goto L1
L117:
	;
	v453 = int32(_a_F_pg_strerror_r_8)
	goto L1
L118:
	;
	v453 = int32(_a_F_pg_strerror_r_9)
	goto L1
L119:
	;
	v453 = int32(_a_F_pg_strerror_r_10)
	goto L1
L120:
	;
	v453 = int32(_a_F_pg_strerror_r_11)
	goto L1
L121:
	;
	v453 = int32(_a_F_pg_strerror_r_12)
	goto L1
L122:
	;
	v453 = int32(_a_F_pg_strerror_r_13)
	goto L1
L123:
	;
	v453 = int32(_a_F_pg_strerror_r_14)
	goto L1
L124:
	;
	v453 = int32(_a_F_pg_strerror_r_15)
	goto L1
L125:
	;
	v453 = int32(_a_F_pg_strerror_r_16)
	goto L1
L126:
	;
	v453 = int32(_a_F_pg_strerror_r_17)
	goto L1
L127:
	;
	v453 = int32(_a_F_pg_strerror_r_18)
	goto L1
L128:
	;
	v453 = int32(_a_F_pg_strerror_r_19)
	goto L1
L129:
	;
	v453 = int32(_a_F_pg_strerror_r_20)
	goto L1
L130:
	;
	v453 = int32(_a_F_pg_strerror_r_21)
	goto L1
L131:
	;
	v453 = int32(_a_F_pg_strerror_r_22)
	goto L1
L132:
	;
	v453 = int32(_a_F_pg_strerror_r_23)
	goto L1
L133:
	;
	v453 = int32(_a_F_pg_strerror_r_24)
	goto L1
L134:
	;
	v453 = int32(_a_F_pg_strerror_r_25)
	goto L1
L135:
	;
	v453 = int32(_a_F_pg_strerror_r_26)
	goto L1
L136:
	;
	v453 = int32(_a_F_pg_strerror_r_27)
	goto L1
L137:
	;
	v453 = int32(_a_F_pg_strerror_r_28)
	goto L1
L138:
	;
	v453 = int32(_a_F_pg_strerror_r_29)
	goto L1
L139:
	;
	v453 = int32(_a_F_pg_strerror_r_30)
	goto L1
L140:
	;
	v453 = int32(_a_F_pg_strerror_r_31)
	goto L1
L141:
	;
	v453 = int32(_a_F_pg_strerror_r_32)
	goto L1
L142:
	;
	v453 = int32(_a_F_pg_strerror_r_33)
	goto L1
L143:
	;
	v453 = int32(_a_F_pg_strerror_r_34)
	goto L1
L144:
	;
	v453 = int32(_a_F_pg_strerror_r_35)
	goto L1
L145:
	;
	v453 = int32(_a_F_pg_strerror_r_36)
	goto L1
L146:
	;
	v453 = int32(_a_F_pg_strerror_r_37)
	goto L1
L147:
	;
	v453 = int32(_a_F_pg_strerror_r_38)
	goto L1
L148:
	;
	v453 = int32(_a_F_pg_strerror_r_39)
	goto L1
L149:
	;
	v453 = int32(_a_F_pg_strerror_r_40)
	goto L1
L150:
	;
	v453 = int32(_a_F_pg_strerror_r_41)
	goto L1
L151:
	;
	v453 = int32(_a_F_pg_strerror_r_42)
	goto L1
L152:
	;
	v453 = int32(_a_F_pg_strerror_r_43)
	goto L1
L153:
	;
	v453 = int32(_a_F_pg_strerror_r_44)
	goto L1
L154:
	;
	v453 = int32(_a_F_pg_strerror_r_45)
	goto L1
L155:
	;
	v453 = int32(_a_F_pg_strerror_r_46)
	goto L1
L156:
	;
	v453 = int32(_a_F_pg_strerror_r_47)
	goto L1
L157:
	;
	v453 = int32(_a_F_pg_strerror_r_48)
	goto L1
L158:
	;
	v453 = int32(_a_F_pg_strerror_r_49)
	goto L1
L159:
	;
	v453 = int32(_a_F_pg_strerror_r_50)
	goto L1
L160:
	;
	v453 = int32(_a_F_pg_strerror_r_51)
	goto L1
L161:
	;
	v453 = int32(_a_F_pg_strerror_r_52)
	goto L1
L162:
	;
	v453 = int32(_a_F_pg_strerror_r_53)
	goto L1
L163:
	;
	v453 = int32(_a_F_pg_strerror_r_54)
	goto L1
L164:
	;
	v453 = int32(_a_F_pg_strerror_r_55)
	goto L1
L165:
	;
	v453 = int32(_a_F_pg_strerror_r_56)
	goto L1
L166:
	;
	v453 = int32(_a_F_pg_strerror_r_57)
	goto L1
L167:
	;
	v453 = int32(_a_F_pg_strerror_r_58)
	goto L1
L168:
	;
	v453 = int32(_a_F_pg_strerror_r_59)
	goto L1
L169:
	;
	v453 = int32(_a_F_pg_strerror_r_60)
	goto L1
L170:
	;
	v453 = int32(_a_F_pg_strerror_r_61)
	goto L1
L171:
	;
	v453 = int32(_a_F_pg_strerror_r_62)
	goto L1
L172:
	;
	v453 = int32(_a_F_pg_strerror_r_63)
	goto L1
L173:
	;
	return int32(0)
L174:
	;
	goto L2
}
func F_pg_strlower(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v11 == v6 {
		v14 = int32(0)
		if base.B2i32(l1 == v14)|base.B2i32(l3 == v14) == v14 {
			v21 = int32(1)
			v22 = l3 - v21
			v24 = l1 - v21
			if base.Ui32(v22) < base.Ui32(v24) {
				v26 = v22
			} else {
				v26 = v24
			}
			if v26 == int32(0) {
				v86 = int32(0)
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v86))))
				if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
					v103 = v94 | int32(32)
				} else {
					v103 = v94
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))) = uint8(v103)
				v111 = v86 + int32(1)
			} else {
				v30 = int32(1)
				v31 = v26 + v30
				v41 = int32(0)
				v46 = v6
				for {
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v41))))
					if base.Ui32((v49-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v58 = v49 | int32(32)
					} else {
						v58 = v49
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v41))) = uint8(v58)
					v61 = v41 | int32(1)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v61))))
					if base.Ui32((v64-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v73 = v64 | int32(32)
					} else {
						v73 = v64
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v61))) = uint8(v73)
					v75 = int32(2)
					v76 = v41 + v75
					v78 = v46 + v75
					if v78 != v31&int32(-2) {
						v41 = v76
						v46 = v78
						continue
					} else {
						break
					}
					break
				}
				if v31&v30 == int32(0) {
					v111 = v76
				} else {
					v86 = v76
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v86))))
					if base.Ui32((v94-int32(65))&int32(255)) < base.Ui32(int32(26)) {
						v103 = v94 | int32(32)
					} else {
						v103 = v94
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))) = uint8(v103)
					v111 = v86 + int32(1)
				}
			}
			if base.Ui32(l1) <= base.Ui32(v111) {
				v145 = l3
				return v145
			} else {
				v127 = v26 + int32(1)
				v134 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v127))) = uint8(v134)
				return l3
			}
		} else {
			v120 = int32(0)
			if l1 == v120 {
				v145 = l3
				return v145
			} else {
				v127 = v120
				v134 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v127))) = uint8(v134)
				return l3
			}
		}
	} else {
		v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v138 = m.T0[v137].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return int32(0)
		} else {
			v145 = v138
			return v145
		}
	}
}
func F_pg_sync_replication_slots(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	F_CheckSlotPermissions(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[0])))
		if v14 == int32(1) {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[1]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+308))
			v22 = base.B2i32(v20 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[0])) = uint8(v22)
			v24 = v22
		} else {
			v24 = int32(0)
		}
		if v24 != 0 {
			v26 = F_ValidateSlotSyncParams(m, int32(21))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				F_load_file(m, int32(_a_F_pg_sync_replication_slots_0), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					v32 = m.G0
					v34 = v32 - int32(16)
					m.G0 = v34
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[2]))
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
					v41 = m.T0[v40].(func(*base.Module, int32) int32)(m, v37)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						if v41 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(_a_F_pg_sync_replication_slots_1)
									*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(_a_F_pg_sync_replication_slots_2)
									F_errmsg(m, int32(_a_F_pg_sync_replication_slots_3), v34)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_sync_replication_slots_4), int32(1215), int32(_a_F_pg_sync_replication_slots_5))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							m.G0 = v34 + int32(16)
							v68 = v6 + int32(28)
							F_initStringInfo(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[4]))
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
								if v73 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v72
									F_appendStringInfo(m, v68, int32(_a_F_pg_sync_replication_slots_6), v6+int32(16))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int64(0)
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[2]))
										v87 = int32(0)
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v86, v87, v87, v87, v90, v6+int32(44))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											if v96 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int64(0)
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
														*(*int32)(unsafe.Add(mBase, uint32(v6))) = v138
														v140 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v140
														F_errmsg(m, int32(_a_F_pg_sync_replication_slots_7), v6)
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_sync_replication_slots_8), int32(963), int32(_a_F_pg_sync_replication_slots_9))
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
												F_pfree(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int64(0)
												} else {
													F_SyncReplicationSlots(m, v96)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int64(0)
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+64))
														m.T0[v107].(func(*base.Module, int32))(m, v96)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int64(0)
														} else {
															m.G0 = v6 + int32(48)
															return int64(0)
														}
													}
												}
											}
										}
									}
								} else {
									F_appendStringInfoString(m, v6+int32(28), int32(_a_F_pg_sync_replication_slots_10))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int64(0)
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[2]))
										v87 = int32(0)
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
										v96 = m.T0[v95].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v86, v87, v87, v87, v90, v6+int32(44))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											if v96 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(100663808))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int64(0)
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
														*(*int32)(unsafe.Add(mBase, uint32(v6))) = v138
														v140 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
														*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v140
														F_errmsg(m, int32(_a_F_pg_sync_replication_slots_7), v6)
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_sync_replication_slots_8), int32(963), int32(_a_F_pg_sync_replication_slots_9))
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
												F_pfree(m, v100)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int64(0)
												} else {
													F_SyncReplicationSlots(m, v96)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return int64(0)
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sync_replication_slots[3]))
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+64))
														m.T0[v107].(func(*base.Module, int32))(m, v96)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int64(0)
														} else {
															m.G0 = v6 + int32(48)
															return int64(0)
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_sync_replication_slots_11), int32(0))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_sync_replication_slots_8), int32(940), int32(_a_F_pg_sync_replication_slots_9))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_table_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_RelationIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_tablespace_location(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = m.G0
	v7 = v5 - int32(2208)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tablespace_location[0]))
	if v4 != 0 {
		v11 = v4
	} else {
		v11 = v10
	}
	if base.Ui32(v11-int32(1663)) <= base.Ui32(int32(1)) {
		v17 = F_pstrdup(m, int32(_a_F_pg_tablespace_location_0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v63 = v17
			m.G0 = v7 + int32(2208)
			v123 = F_cstring_to_text(m, v63)
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v123)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_pg_tablespace_location_1)
		v25 = v7 + int32(1184)
		v30 = F_pg_snprintf(m, v25, int32(1024), int32(_a_F_pg_tablespace_location_2), v7+int32(48))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int64(0)
		} else {
			v36 = F___fstatat(m, int32(-100), v25, v7-int32(-64), int32(256))
			mBase = m.M
			if v36 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(1184)
						F_errmsg(m, int32(_a_F_pg_tablespace_location_3), v7)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_tablespace_location_4), int32(68), int32(_a_F_pg_tablespace_location_5))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
				if v39&int32(_a_F_pg_tablespace_location_6) != int32(_a_F_pg_tablespace_location_7) {
					v44 = F_pstrdup(m, v25)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						v63 = v44
						m.G0 = v7 + int32(2208)
						v123 = F_cstring_to_text(m, v63)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v123)
						}
					}
				} else {
					v49 = v7 + int32(160)
					v51 = F_readlink(m, v7+int32(1184), v49, int32(1024))
					mBase = m.M
					if v51 < int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int64(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1184)
								F_errmsg(m, int32(_a_F_pg_tablespace_location_8), v7+int32(16))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_tablespace_location_4), int32(81), int32(_a_F_pg_tablespace_location_5))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(1024)) <= base.Ui32(v51) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v7 + int32(1184)
									F_errmsg(m, int32(_a_F_pg_tablespace_location_9), v7+int32(32))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_tablespace_location_4), int32(86), int32(_a_F_pg_tablespace_location_5))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v57 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v51+v49))) = uint8(v57)
							v59 = F_pstrdup(m, v49)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								v63 = v59
								m.G0 = v7 + int32(2208)
								v123 = F_cstring_to_text(m, v63)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v123)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_timezone_names(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int64
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int64
	_ = v232
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v253 int64
	_ = v253
	var v259 int32
	_ = v259
	var v261 int64
	_ = v261
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v330 int32
	_ = v330
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+92)) = v2
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = F_palloc0(m, int32(_a_F_pg_timezone_names_0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_names[0])))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_get_share_path(m, int32(_a_F_pg_timezone_names_1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v163 = F_pstrdup(m, int32(_a_F_pg_timezone_names_1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L39
	}
L7:
	;
	v33 = int32(_a_F_pg_timezone_names_1)
	v34 = F_strlen(m, v33)
	mBase = m.M
	v36 = v34 + v33
	v37 = int32(_a_F_pg_timezone_names_2)
	v39 = int32(1024) - v34
	if v39 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_names[0])) = uint8(v159)
	goto L6
L9:
	;
	v155 = F_strlen(m, v151)
	mBase = m.M
	goto L8
L10:
	;
	v151 = v37
	goto L9
L11:
	;
	goto L12
L12:
	;
	v45 = v39 - int32(1)
	if (v36^v37)&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v148)
	v151 = v144
	goto L9
L14:
	;
	v129 = v124
	v130 = v125
	v131 = v126
	goto L35
L15:
	;
	if v119 == int32(0) {
		v144 = v117
		v145 = v118
		goto L13
	} else {
		goto L34
	}
L16:
	;
	v117 = v37
	v118 = v36
	v119 = v45
	goto L15
L17:
	;
	goto L18
L18:
	;
	v49 = int32(0)
	if int32(0)|base.B2i32(v45 == v49) == v49 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v85 == int32(0) {
		v144 = v82
		v145 = v83
		goto L13
	} else {
		goto L28
	}
L20:
	;
	v61 = v37
	v62 = v36
	v63 = v45
	goto L23
L21:
	;
	goto L22
L22:
	;
	v82 = v37
	v83 = v36
	v84 = v45
	v85 = base.B2i32(v45 != v49)
	goto L19
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v65)
	if v65 == int32(0) {
		v124 = v61
		v125 = v62
		v126 = v63
		goto L14
	} else {
		goto L25
	}
L24:
	;
	v82 = v76
	v83 = v70
	v84 = v72
	v85 = v74
	goto L19
L25:
	;
	v69 = int32(1)
	v70 = v62 + v69
	v72 = v63 - v69
	v73 = int32(0)
	v74 = base.B2i32(v72 != v73)
	v76 = v61 + v69
	if v76&int32(3) == v73 {
		v82 = v76
		v83 = v70
		v84 = v72
		v85 = v74
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if v72 != 0 {
		v61 = v76
		v62 = v70
		v63 = v72
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if base.B2i32(v88 == int32(0))|base.B2i32(base.Ui32(v84) < base.Ui32(int32(4))) != 0 {
		v117 = v82
		v118 = v83
		v119 = v84
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v95 = v82
	v96 = v83
	v97 = v84
	goto L30
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v103 = int32(-2139062144)
	if (int32(16843008)-v100|v100)&v103 != v103 {
		v124 = v95
		v125 = v96
		v126 = v97
		goto L14
	} else {
		goto L32
	}
L31:
	;
	v117 = v111
	v118 = v109
	v119 = v113
	goto L15
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v100
	v108 = int32(4)
	v109 = v96 + v108
	v111 = v95 + v108
	v113 = v97 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v95 = v111
		v96 = v109
		v97 = v113
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v124 = v117
	v125 = v118
	v126 = v119
	goto L14
L35:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v133)
	if v133 == int32(0) {
		v144 = v129
		v145 = v130
		goto L13
	} else {
		goto L37
	}
L36:
	;
	v144 = v140
	v145 = v138
	goto L13
L37:
	;
	v137 = int32(1)
	v138 = v130 + v137
	v140 = v129 + v137
	v142 = v131 - v137
	if v142 != 0 {
		v129 = v140
		v130 = v138
		v131 = v142
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v165 = F_strlen(m, v163)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v165 + int32(1)
	v172 = F_AllocateDir(m, v163)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	if v172 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	m.G0 = v21 + int32(16)
	v195 = F_pg_tzenumerate_next(m, v24)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L48
	}
L44:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v163
	F_errmsg(m, int32(_a_F_pg_timezone_names_3), v21)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_names_4), int32(409), int32(_a_F_pg_timezone_names_5))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	if v195 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v198 = v9 + int32(16)
	v199 = v195
	goto L52
L50:
	;
	goto L51
L51:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) <= v291 {
		goto L74
	} else {
		goto L75
	}
L52:
	;
	v206 = *(*int64)(unsafe.Add(mBase, _c_F_pg_timezone_names[1]))
	v215 = F_timestamp2tm(m, v206, v9+int32(88), v9+int32(44), v9+int32(40), v9+int32(36), v199)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	v283 = F_pg_tzenumerate_next(m, v24)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L72
	}
L55:
	;
	if v215 != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v217 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v218 = F_strlen(m, v217)
	mBase = m.M
	if base.Ui32(int32(31)) < base.Ui32(v218) {
		goto L54
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v221 = F_cstring_to_text(m, v199)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = base.I64_extend_i32_u(v221)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v225 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v227 = v225
	goto L64
L63:
	;
	v227 = int32(_a_F_pg_timezone_names_6)
	goto L64
L64:
	;
	v228 = F_cstring_to_text(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = base.I64_extend_i32_u(v228)
	v232 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v198)+8)) = v232
	*(*int64)(unsafe.Add(mBase, uint32(v198))) = v232
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v9)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = base.I64_extend_i32_s(int32(0)-v237) * int64(1000000)
	v244 = v9 + int32(8)
	v246 = F_palloc(m, int32(16))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v249 = int64(*(*int32)(unsafe.Add(mBase, uint32(v244)+12)))
	v250 = int64(*(*int32)(unsafe.Add(mBase, uint32(v244)+16)))
	v253 = v249 + v250*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v253-int64(2147483648)) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = base.I64_extend_i32_u(v246)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = base.I64_extend_i32_u(base.B2i32(int32(0) < v268))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v273, v274, v9+int32(96), v9+int32(92))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L71
	}
L68:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v246)+12)) = uint32(v253)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	*(*int64)(unsafe.Add(mBase, uint32(v246))) = v261
	goto L70
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L54
L72:
	;
	if v283 != 0 {
		v199 = v283
		goto L52
	} else {
		goto L73
	}
L73:
	;
	goto L53
L74:
	;
	v298 = v291
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_pfree(m, v24)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(8)+v298<<(uint(int32(2))%32))))
	F_FreeDir(m, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L76
L79:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v24+int32(48)+v310<<(uint(int32(2))%32))))
	F_pfree(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v319 = v317 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v319
	if int32(0) <= v319 {
		v298 = v319
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	m.G0 = v9 + int32(128)
	return int64(0)
}
func F_pg_ts_config_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSConfigIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_ts_template_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSTemplateIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_tzset_offset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	v12 = l0 >> (uint(int32(31)) % 32)
	v14 = l0 ^ v12 - v12
	v16 = base.I32_div_s(v14, int32(3600))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v16
	v19 = v9 + int32(192)
	v24 = F_pg_snprintf(m, v19, int32(64), int32(_a_F_pg_tzset_offset_0), v9+int32(48))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v30 = v14 - v16*int32(3600)
		if v30 == int32(0) {
			v69 = v9 + int32(192)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
			v73 = v9 - int32(-64)
			if int32(0) < l0 {
				v79 = int32(_a_F_pg_tzset_offset_1)
			} else {
				v79 = int32(_a_F_pg_tzset_offset_2)
			}
			v80 = F_pg_snprintf(m, v73, int32(128), v79, v9)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				v82 = F_pg_tzset(m, v73)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(256)
					return v82
				}
			}
		} else {
			v33 = F_strlen(m, v19)
			mBase = m.M
			v36 = base.I32_div_s(base.I32_extend16_s(v30), int32(60))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = base.I32_extend16_s(v36)
			v45 = F_pg_snprintf(m, v33+v19, int32(64)-v33, int32(_a_F_pg_tzset_offset_3), v9+int32(32))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v49 = v30 - v36*int32(60)
				if v49&int32(_a_F_pg_tzset_offset_4) == int32(0) {
					v69 = v9 + int32(192)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
					v73 = v9 - int32(-64)
					if int32(0) < l0 {
						v79 = int32(_a_F_pg_tzset_offset_1)
					} else {
						v79 = int32(_a_F_pg_tzset_offset_2)
					}
					v80 = F_pg_snprintf(m, v73, int32(128), v79, v9)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = F_pg_tzset(m, v73)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(256)
							return v82
						}
					}
				} else {
					v54 = F_strlen(m, v19)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = base.I32_extend16_s(v49)
					v63 = F_pg_snprintf(m, v54+v19, int32(64)-v54, int32(_a_F_pg_tzset_offset_3), v9+int32(16))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v69 = v9 + int32(192)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
						v73 = v9 - int32(-64)
						if int32(0) < l0 {
							v79 = int32(_a_F_pg_tzset_offset_1)
						} else {
							v79 = int32(_a_F_pg_tzset_offset_2)
						}
						v80 = F_pg_snprintf(m, v73, int32(128), v79, v9)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = F_pg_tzset(m, v73)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(256)
								return v82
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_ultoa_n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v3 = int32(0)
	if l0 == v3 {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(1)
	} else {
		v19 = int32(1233)
		v24 = int32(base.Ui32((base.I32_clz(l0)^int32(31))*v19+v19) >> (uint(int32(12)) % 32))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v24<<(uint(int32(2))%32))+uint32(_c_F_pg_ultoa_n[0])))
		v29 = v24 + base.B2i32(base.Ui32(v27) <= base.Ui32(l0))
		if base.Ui32(int32(_a_F_pg_ultoa_n_0)) <= base.Ui32(l0) {
			v33 = l0
			v35 = v3
			for {
				v42 = l1 + v29 - v35
				v43 = int32(4)
				v46 = base.I32_div_u_s(v33, int32(_a_F_pg_ultoa_n_0))
				v49 = v33 + v46*int32(-10000)
				v50 = int32(100)
				v51 = base.I32_div_u_s(v49, v50)
				v52 = int32(1)
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51<<(uint(v52)%32))+uint32(_c_F_pg_ultoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v42-v43))) = uint16(v54)
				v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v49-v51*v50)<<(uint(v52)%32))+uint32(_c_F_pg_ultoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v42-int32(2)))) = uint16(v63)
				v66 = v35 + v43
				if base.Ui32(int32(99999999)) < base.Ui32(v33) {
					v33 = v46
					v35 = v66
					continue
				} else {
					break
				}
				break
			}
			v69 = v46
			v71 = v66
		} else {
			v69 = l0
			v71 = v3
		}
		if base.Ui32(int32(100)) <= base.Ui32(v69) {
			v82 = int32(2)
			v84 = int32(_a_F_pg_ultoa_n_1)
			v86 = int32(100)
			v87 = base.I32_div_u_s(v69&v84, v86)
			v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v69-v87*v86)&v84<<(uint(int32(1))%32))+uint32(_c_F_pg_ultoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v29-v71-v82))) = uint16(v95)
			v99 = v87
			v100 = v71 | v82
		} else {
			v99 = v69
			v100 = v71
		}
		if base.Ui32(int32(10)) <= base.Ui32(v99) {
			v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99<<(uint(int32(1))%32))+uint32(_c_F_pg_ultoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v29-v100-int32(2)))) = uint16(v109)
			return v29
		} else {
			v113 = v99 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v113)
			return v29
		}
	}
}
func F_pg_usleep(m *base.Module, l0 int32) {
	if int32(0) < l0 {
		m.Env.Pgmem_usleep(m, l0)
	} else {
	}
	return
}
func F_pg_utf8_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	if base.Ui32(l1) < base.Ui32(int32(16)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v326 <= int32(0) {
		v455 = v328
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v326 = l1
	v328 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = l0
	v29 = int32(11)
	v31 = l1
	goto L5
L5:
	;
	if v29 != int32(11) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	switch v252 {
	case 0:
		v326 = l1
		v328 = l0
		goto L1
	default:
		goto L16
	case 11:
		v305 = v254
		v308 = v256
		goto L15
	}
L7:
	;
	v253 = int32(16)
	v254 = v28 + v253
	v256 = v31 - v253
	if base.Ui32(int32(15)) < base.Ui32(v256) {
		v28 = v254
		v29 = v252
		v31 = v256
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+15)))
	v141 = int32(2)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v144 = int32(255)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v123&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v124&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v125&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v126&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v127&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v128&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v129&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v130<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v131&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v132&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v133&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v134&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v135&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v136&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v137&v144<<(uint(v141)%32))+uint32(_c_F_pg_utf8_verifystr[0])))
	v252 = int32(base.Ui32(v143)>>(uint(int32(base.Ui32(v148)>>(uint(int32(base.Ui32(v153)>>(uint(int32(base.Ui32(v158)>>(uint(int32(base.Ui32(v163)>>(uint(int32(base.Ui32(v168)>>(uint(int32(base.Ui32(v173)>>(uint(int32(base.Ui32(v178)>>(uint(int32(base.Ui32(v181)>>(uint(int32(base.Ui32(v186)>>(uint(int32(base.Ui32(v191)>>(uint(int32(base.Ui32(v196)>>(uint(int32(base.Ui32(v201)>>(uint(int32(base.Ui32(v206)>>(uint(int32(base.Ui32(v211)>>(uint(int32(base.Ui32(v216)>>(uint(v29)%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)))%32)) & int32(31)
	goto L7
L9:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+14)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+13)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+12)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+11)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+7)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+6)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v123 = v50
	v124 = v51
	v125 = v52
	v126 = v53
	v127 = v54
	v128 = v55
	v129 = v56
	v130 = v57
	v131 = v58
	v132 = v59
	v133 = v60
	v134 = v61
	v135 = v62
	v136 = v63
	v137 = v64
	goto L8
L10:
	;
	goto L11
L11:
	;
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v66 = int64(48)
	v68 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v66) % 64)))
	v69 = int64(40)
	v71 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v69) % 64)))
	v72 = int64(32)
	v74 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v72) % 64)))
	v75 = int64(24)
	v77 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v75) % 64)))
	v78 = int64(16)
	v80 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v78) % 64)))
	v81 = int64(8)
	v83 = base.I32_wrap_i64(int64(base.Ui64(v65) >> (uint(v81) % 64)))
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v87 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(int64(56)) % 64)))
	v90 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v66) % 64)))
	v93 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v69) % 64)))
	v96 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v72) % 64)))
	v99 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v75) % 64)))
	v102 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v78) % 64)))
	v105 = base.I32_wrap_i64(int64(base.Ui64(v84) >> (uint(v81) % 64)))
	v106 = base.I32_wrap_i64(v65)
	v107 = base.I32_wrap_i64(v84)
	if (v84|v65)&int64(-9187201950435737472) != int64(0) {
		v123 = v68
		v124 = v71
		v125 = v74
		v126 = v77
		v127 = v80
		v128 = v83
		v129 = v106
		v130 = v87
		v131 = v90
		v132 = v93
		v133 = v96
		v134 = v99
		v135 = v102
		v136 = v105
		v137 = v107
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v114 = int64(9187201950435737471)
	v119 = int64(-9187201950435737472)
	if (v84+v114)&(v65+v114)&v119 == v119 {
		v252 = int32(11)
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v123 = v68
	v124 = v71
	v125 = v74
	v126 = v77
	v127 = v80
	v128 = v83
	v129 = v106
	v130 = v87
	v131 = v90
	v132 = v93
	v133 = v96
	v134 = v99
	v135 = v102
	v136 = v105
	v137 = v107
	goto L8
L14:
	;
	goto L6
L15:
	;
	v326 = v308
	v328 = v305
	goto L1
L16:
	;
	v261 = v254
	v264 = v256
	goto L17
L17:
	;
	v281 = int32(1)
	v282 = v264 + v281
	v284 = v261 - v281
	v285 = int32(*(*int8)(unsafe.Add(mBase, uint32(v284))))
	v287 = v285 & int32(255)
	if int32(0) <= v285 {
		v261 = v284
		v264 = v282
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v326 = v282
	v328 = v284
	goto L1
L19:
	;
	if base.B2i32(v287&int32(248) == int32(240))|base.B2i32(v287&int32(224) == int32(192)) != 0 {
		v305 = v284
		v308 = v282
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v287&int32(240) != int32(224) {
		v261 = v284
		v264 = v282
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	return v455 - l0
L23:
	;
	v350 = v326
	v352 = v328
	goto L24
L24:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v372 = base.I32_extend8_s(v371)
	if int32(0) <= v372 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v455 = v448
	goto L22
L26:
	;
	v448 = v447 + v352
	v449 = v350 - v447
	if int32(0) < v449 {
		v350 = v449
		v352 = v448
		goto L24
	} else {
		goto L60
	}
L27:
	;
	if v372 != 0 {
		v447 = int32(1)
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v371&int32(224) == int32(192) {
		v393 = int32(2)
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v455 = v352
	goto L22
L31:
	;
	if base.Ui32(v350) < base.Ui32(v393) {
		v455 = v352
		goto L22
	} else {
		goto L37
	}
L32:
	;
	if v371&int32(240) == int32(224) {
		v393 = int32(3)
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v371&int32(248) == int32(240) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v392 = int32(4)
	goto L36
L35:
	;
	v392 = int32(1)
	goto L36
L36:
	;
	v393 = v392
	goto L31
L37:
	;
	v395 = int32(0)
	switch v393 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L43
	case 2:
		goto L44
	case 3:
		goto L45
	default:
		v444 = v395
		goto L39
	}
L38:
	;
	if v444 == int32(0) {
		v455 = v352
		goto L22
	} else {
		goto L59
	}
L39:
	;
	goto L38
L40:
	;
	v444 = base.B2i32(base.Ui32(v436&int32(255)) < base.Ui32(int32(245)))
	goto L39
L41:
	;
	if base.I32_extend8_s(v431) < int32(-62) {
		v444 = v395
		goto L39
	} else {
		goto L58
	}
L42:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	v431 = v430
	goto L41
L43:
	;
	v404 = int32(*(*int8)(unsafe.Add(mBase, uint32(v352)+1)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	switch v405 - int32(224) {
	case 0:
		goto L52
	default:
		goto L48
	case 13:
		goto L51
	case 16:
		goto L50
	case 20:
		goto L49
	}
L44:
	;
	v401 = int32(*(*int8)(unsafe.Add(mBase, uint32(v352)+2)))
	if int32(-65) < v401 {
		v444 = v395
		goto L39
	} else {
		goto L47
	}
L45:
	;
	v398 = int32(*(*int8)(unsafe.Add(mBase, uint32(v352)+3)))
	if int32(-65) < v398 {
		v444 = v395
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L43
L48:
	;
	if v404 <= int32(-65) {
		v431 = v405
		goto L41
	} else {
		goto L57
	}
L49:
	;
	if int32(-113) < v404 {
		v444 = v395
		goto L39
	} else {
		goto L56
	}
L50:
	;
	if base.Ui32((v404-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v444 = v395
		goto L39
	} else {
		goto L55
	}
L51:
	;
	if int32(-97) < v404 {
		v444 = v395
		goto L39
	} else {
		goto L54
	}
L52:
	;
	v408 = int32(224)
	if base.Ui32(v408) <= base.Ui32((v404-int32(-64))&int32(255)) {
		v436 = v408
		goto L40
	} else {
		goto L53
	}
L53:
	;
	v444 = v395
	goto L39
L54:
	;
	v436 = int32(237)
	goto L40
L55:
	;
	v436 = int32(240)
	goto L40
L56:
	;
	v436 = int32(244)
	goto L40
L57:
	;
	v444 = v395
	goto L39
L58:
	;
	v436 = v431
	goto L40
L59:
	;
	v447 = v393
	goto L26
L60:
	;
	goto L25
}
