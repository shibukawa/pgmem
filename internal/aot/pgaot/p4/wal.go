package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetWalSummaries(m *base.Module, l0 int32, l1 int64, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v22 = F_AllocateDir(m, int32(_a_F_GetWalSummaries_0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = F_ReadDir(m, v22, int32(_a_F_GetWalSummaries_0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = v17 + int32(-32)
	v45 = v27
	v48 = v4
	goto L7
L5:
	;
	v228 = v4
	goto L6
L6:
	;
	F_FreeDir(m, v22)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L49
	}
L7:
	;
	v58 = v45 + int32(19)
	v59 = int32(_a_F_GetWalSummaries_1)
	v63 = m.G0
	v65 = v63 - int32(32)
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v66
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetWalSummaries[0])))
	if v74 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v228 = v213
	goto L6
L9:
	;
	v219 = F_ReadDir(m, v22, int32(_a_F_GetWalSummaries_0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L47
	}
L10:
	;
	if v142 != int32(40) {
		v213 = v48
		goto L9
	} else {
		goto L29
	}
L11:
	;
	v142 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetWalSummaries[1])))
	if v78 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v82 = v58
	goto L17
L15:
	;
	goto L16
L16:
	;
	v92 = v59
	v93 = v74
	goto L20
L17:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v88 == v74 {
		v82 = v82 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v142 = v82 - v58
	goto L10
L19:
	;
	goto L18
L20:
	;
	v100 = v65 + int32(base.Ui32(v93)>>(uint(int32(3))%32))&int32(28)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v101 | v102<<(uint(v93)%32)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v106 != 0 {
		v92 = v92 + v102
		v93 = v106
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v109 == int32(0) {
		v132 = v58
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v142 = v132 - v58
	goto L10
L24:
	;
	v113 = v58
	v114 = v109
	goto L25
L25:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v65+int32(base.Ui32(v114)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v122)>>(uint(v114)%32))&int32(1) == int32(0) {
		v132 = v113
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v132 = v130
	goto L23
L27:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	v130 = v113 + int32(1)
	if v128 != 0 {
		v113 = v130
		v114 = v128
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v146 = v45 + int32(59)
	v147 = int32(_a_F_GetWalSummaries_2)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetWalSummaries[2])))
	if base.B2i32(v150 == int32(0))|base.B2i32(v150 != v153) != 0 {
		v171 = v150
		v172 = v153
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v171-v172 != 0 {
		v213 = v48
		goto L9
	} else {
		goto L37
	}
L31:
	;
	goto L30
L32:
	;
	v156 = v146
	v157 = v147
	goto L33
L33:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+1)))
	if v161 == int32(0) {
		v171 = v161
		v172 = v160
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v171 = v161
	v172 = v160
	goto L31
L35:
	;
	v164 = int32(1)
	if v161 == v160 {
		v156 = v156 + v164
		v157 = v157 + v164
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v17 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v34 | int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v34 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v34 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v17 + int32(-32)
	v182 = F_sscanf(m, v58, int32(_a_F_GetWalSummaries_3), v19)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v184 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+44)))
	v185 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+48)))
	v186 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+36)))
	v187 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+40)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if v189 != l0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v191 = l0
	goto L41
L40:
	;
	v191 = int32(0)
	goto L41
L41:
	;
	if v191 != 0 {
		v213 = v48
		goto L9
	} else {
		goto L42
	}
L42:
	;
	v194 = v186<<(uint(int64(32))%64) | v187
	if base.Ui64(l2-int64(1)) < base.Ui64(v194) {
		v213 = v48
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v200 = v184<<(uint(int64(32))%64) | v185
	if base.B2i32(l1 != int64(0))&base.B2i32(base.Ui64(v200) <= base.Ui64(l1)) != 0 {
		v213 = v48
		goto L9
	} else {
		goto L44
	}
L44:
	;
	v204 = F_palloc(m, int32(24))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v204)+8)) = v200
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v204)+16)) = v189
	v209 = F_lappend(m, v48, v204)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v213 = v209
	goto L9
L47:
	;
	if v219 != 0 {
		v45 = v219
		v48 = v213
		goto L7
	} else {
		goto L48
	}
L48:
	;
	goto L8
L49:
	;
	m.G0 = v19 - int32(-64)
	return v228
}
func F_WalReceiverMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
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
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int64
	_ = v336
	var v337 int64
	_ = v337
	var v345 int64
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int64
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int64
	_ = v950
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v993 int64
	_ = v993
	var v999 int32
	_ = v999
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1204 int64
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1222 int64
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int64
	_ = v1240
	var v1241 int64
	_ = v1241
	var v1249 int64
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1258 int64
	_ = v1258
	var v1268 int64
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1277 int64
	_ = v1277
	var v1280 int64
	_ = v1280
	var v1286 int64
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int64
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int64
	_ = v1352
	var v1353 int64
	_ = v1353
	var v1361 int64
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1369 int32
	_ = v1369
	var v1370 int64
	_ = v1370
	var v1380 int64
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1389 int64
	_ = v1389
	var v1392 int64
	_ = v1392
	var v1398 int64
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int64
	_ = v1402
	var v1406 int32
	_ = v1406
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int64
	_ = v1447
	var v1448 int64
	_ = v1448
	var v1456 int64
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1462 int64
	_ = v1462
	var v1465 int64
	_ = v1465
	var v1474 int64
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1496 int32
	_ = v1496
	var v1497 int64
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int64
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int64
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1513 int64
	_ = v1513
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1534 int64
	_ = v1534
	var v1536 int64
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1549 int64
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1574 int64
	_ = v1574
	var v1575 int64
	_ = v1575
	var v1579 int64
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1606 int32
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1619 int64
	_ = v1619
	var v1620 int64
	_ = v1620
	var v1624 int64
	_ = v1624
	var v1674 int32
	_ = v1674
	var v1675 int64
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1686 int32
	_ = v1686
	var v1691 int64
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1711 int32
	_ = v1711
	var v1712 int64
	_ = v1712
	var v1716 int64
	_ = v1716
	var v1721 int32
	_ = v1721
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1745 int64
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1772 int64
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1784 int32
	_ = v1784
	var v1785 int64
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int64
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1798 int32
	_ = v1798
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1814 int32
	_ = v1814
	var v1816 int64
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1845 int32
	_ = v1845
	var v1850 int32
	_ = v1850
	var v1853 int64
	_ = v1853
	var v1854 int64
	_ = v1854
	var v1868 int32
	_ = v1868
	var v1870 int64
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1876 int64
	_ = v1876
	var v1878 int64
	_ = v1878
	var v1879 int64
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1900 int32
	_ = v1900
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1941 int64
	_ = v1941
	var v1944 int64
	_ = v1944
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v2004 int64
	_ = v2004
	var v2006 int64
	_ = v2006
	var v2008 int64
	_ = v2008
	var v2010 int64
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2019 int64
	_ = v2019
	var v2020 int64
	_ = v2020
	var v2028 int64
	_ = v2028
	var v2030 int64
	_ = v2030
	var v2032 int64
	_ = v2032
	var v2034 int64
	_ = v2034
	var v2040 int64
	_ = v2040
	var v2049 int64
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2079 int32
	_ = v2079
	var v2084 int32
	_ = v2084
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2093 int64
	_ = v2093
	var v2094 int64
	_ = v2094
	var v2102 int64
	_ = v2102
	var v2104 int64
	_ = v2104
	var v2107 int64
	_ = v2107
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2155 int32
	_ = v2155
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2199 int32
	_ = v2199
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2213 int32
	_ = v2213
	var v2218 int32
	_ = v2218
	var v2227 int32
	_ = v2227
	var v2236 int32
	_ = v2236
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2247 int64
	_ = v2247
	var v2250 int64
	_ = v2250
	var v2251 int64
	_ = v2251
	var v2252 int64
	_ = v2252
	var v2255 int64
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2321 int32
	_ = v2321
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2341 int32
	_ = v2341
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int64
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2408 int64
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2431 int64
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2457 int32
	_ = v2457
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2470 int32
	_ = v2470
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2504 int32
	_ = v2504
	var v2509 int32
	_ = v2509
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2523 int32
	_ = v2523
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2562 int32
	_ = v2562
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(1504)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[0])) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+320)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v19)+316)) = v7
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	v32 = int32(1456)
	v33 = v31 + v32
	v36 = base.AtomicRmwXchg32(m, v31, v32, int32(1))
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_s_lock(m, v33, int32(_a_F_WalReceiverMain_0), int32(188), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	switch v42 {
	case 0:
		goto L9
	case 1:
		goto L7
	default:
		goto L8
	case 5:
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[2]))
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1453)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v72
	v79 = v19 + int32(400)
	v81 = v31 + int32(104)
	goto L19
L8:
	;
	v55 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v33))), uint32(v55))
	F_errstart_cold(m, int32(23), v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L13
	}
L9:
	;
	v45 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v31)+1456)), uint32(v45))
	F_ConditionVariableBroadcast(m, v31+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(0)
	goto L9
L11:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_2), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(213), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v202 = v19 + int32(336)
	v204 = v31 + int32(1388)
	goto L50
L17:
	;
	v198 = F_strlen(m, v187)
	mBase = m.M
	goto L16
L19:
	;
	goto L20
L20:
	;
	v88 = int32(1023)
	if (v79^v81)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v191)
	goto L17
L22:
	;
	v172 = v167
	v173 = v168
	v174 = v169
	goto L43
L23:
	;
	if v162 == int32(0) {
		v187 = v160
		v188 = v161
		goto L21
	} else {
		goto L42
	}
L24:
	;
	v160 = v81
	v161 = v79
	v162 = v88
	goto L23
L25:
	;
	goto L26
L26:
	;
	v92 = int32(0)
	if base.B2i32(v81&int32(3) == v92)|int32(0) == v92 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v128 == int32(0) {
		v187 = v125
		v188 = v126
		goto L21
	} else {
		goto L36
	}
L28:
	;
	v104 = v81
	v105 = v79
	v106 = v88
	goto L31
L29:
	;
	goto L30
L30:
	;
	v125 = v81
	v126 = v79
	v127 = v88
	v128 = int32(1)
	goto L27
L31:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v108)
	if v108 == int32(0) {
		v167 = v104
		v168 = v105
		v169 = v106
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v125 = v119
	v126 = v113
	v127 = v115
	v128 = v117
	goto L27
L33:
	;
	v112 = int32(1)
	v113 = v105 + v112
	v115 = v106 - v112
	v116 = int32(0)
	v117 = base.B2i32(v115 != v116)
	v119 = v104 + v112
	if v119&int32(3) == v116 {
		v125 = v119
		v126 = v113
		v127 = v115
		v128 = v117
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v115 != 0 {
		v104 = v119
		v105 = v113
		v106 = v115
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if base.B2i32(v131 == int32(0))|base.B2i32(base.Ui32(v127) < base.Ui32(int32(4))) != 0 {
		v160 = v125
		v161 = v126
		v162 = v127
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v138 = v125
	v139 = v126
	v140 = v127
	goto L38
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v146 = int32(-2139062144)
	if (int32(16843008)-v143|v143)&v146 != v146 {
		v167 = v138
		v168 = v139
		v169 = v140
		goto L22
	} else {
		goto L40
	}
L39:
	;
	v160 = v154
	v161 = v152
	v162 = v156
	goto L23
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v143
	v151 = int32(4)
	v152 = v139 + v151
	v154 = v138 + v151
	v156 = v140 - v151
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v138 = v154
		v139 = v152
		v140 = v156
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v167 = v160
	v168 = v161
	v169 = v162
	goto L22
L43:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v176)
	if v176 == int32(0) {
		v187 = v172
		v188 = v173
		goto L21
	} else {
		goto L45
	}
L44:
	;
	v187 = v183
	v188 = v181
	goto L21
L45:
	;
	v180 = int32(1)
	v181 = v173 + v180
	v183 = v172 + v180
	v185 = v174 - v180
	if v185 != 0 {
		v172 = v183
		v173 = v181
		v174 = v185
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1452)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+332)) = v326
	v331 = m.G0
	v332 = int32(16)
	v333 = v331 - v332
	m.G0 = v333
	F_gettimeofday(m, v333)
	mBase = m.M
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v333)))
	v337 = int64(*(*int32)(unsafe.Add(mBase, uint32(v333)+8)))
	m.G0 = v333 + v332
	v345 = v337 + v336*int64(1000000) - int64(946684800000000)
	goto L78
L48:
	;
	v321 = F_strlen(m, v310)
	mBase = m.M
	goto L47
L50:
	;
	goto L51
L51:
	;
	v211 = int32(63)
	if (v202^v204)&int32(3) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v314 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v314)
	goto L48
L53:
	;
	v295 = v290
	v296 = v291
	v297 = v292
	goto L74
L54:
	;
	if v285 == int32(0) {
		v310 = v283
		v311 = v284
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v283 = v204
	v284 = v202
	v285 = v211
	goto L54
L56:
	;
	goto L57
L57:
	;
	v215 = int32(0)
	if base.B2i32(v204&int32(3) == v215)|int32(0) == v215 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v251 == int32(0) {
		v310 = v248
		v311 = v249
		goto L52
	} else {
		goto L67
	}
L59:
	;
	v227 = v204
	v228 = v202
	v229 = v211
	goto L62
L60:
	;
	goto L61
L61:
	;
	v248 = v204
	v249 = v202
	v250 = v211
	v251 = int32(1)
	goto L58
L62:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v231)
	if v231 == int32(0) {
		v290 = v227
		v291 = v228
		v292 = v229
		goto L53
	} else {
		goto L64
	}
L63:
	;
	v248 = v242
	v249 = v236
	v250 = v238
	v251 = v240
	goto L58
L64:
	;
	v235 = int32(1)
	v236 = v228 + v235
	v238 = v229 - v235
	v239 = int32(0)
	v240 = base.B2i32(v238 != v239)
	v242 = v227 + v235
	if v242&int32(3) == v239 {
		v248 = v242
		v249 = v236
		v250 = v238
		v251 = v240
		goto L58
	} else {
		goto L65
	}
L65:
	;
	if v238 != 0 {
		v227 = v242
		v228 = v236
		v229 = v238
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if base.B2i32(v254 == int32(0))|base.B2i32(base.Ui32(v250) < base.Ui32(int32(4))) != 0 {
		v283 = v248
		v284 = v249
		v285 = v250
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v261 = v248
	v262 = v249
	v263 = v250
	goto L69
L69:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v269 = int32(-2139062144)
	if (int32(16843008)-v266|v266)&v269 != v269 {
		v290 = v261
		v291 = v262
		v292 = v263
		goto L53
	} else {
		goto L71
	}
L70:
	;
	v283 = v277
	v284 = v275
	v285 = v279
	goto L54
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v266
	v274 = int32(4)
	v275 = v262 + v274
	v277 = v261 + v274
	v279 = v263 - v274
	if base.Ui32(int32(3)) < base.Ui32(v279) {
		v261 = v277
		v262 = v275
		v263 = v279
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v290 = v283
	v291 = v284
	v292 = v285
	goto L53
L74:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	*(*uint8)(unsafe.Add(mBase, uint32(v296))) = uint8(v299)
	if v299 == int32(0) {
		v310 = v295
		v311 = v296
		goto L52
	} else {
		goto L76
	}
L75:
	;
	v310 = v306
	v311 = v304
	goto L52
L76:
	;
	v303 = int32(1)
	v304 = v296 + v303
	v306 = v295 + v303
	v308 = v297 - v303
	if v308 != 0 {
		v295 = v306
		v296 = v304
		v297 = v308
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+80)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(v31)+96)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(v31)+72)) = v345
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v350
	v352 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v31)+1456)), uint32(v352))
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	v359 = base.AtomicRmwXchg64(m, v356, int32(1464), int64(0))
	F_on_shmem_exit(m, int32(1029), v19+int32(332))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v366 = int32(914)
	v368 = m.G0
	v370 = v368 - int32(32)
	m.G0 = v370
	switch int32(916) {
	case 0, 2:
		v380 = v366
		goto L81
	default:
		goto L82
	}
L80:
	;
	v399 = int32(-2)
	v401 = m.G0
	v403 = v401 - int32(32)
	m.G0 = v403
	switch int32(0) {
	case 0, 2:
		v413 = v399
		goto L87
	default:
		goto L88
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+12)) = v380
	F_sigemptyset(m, v370+int32(16))
	mBase = m.M
	goto L84
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[4])) = v366
	v380 = int32(_a_F_WalReceiverMain_3)
	goto L81
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+24)) = int32(268435456)
	v394 = F___sigaction(m, int32(1), v370+int32(12), int32(0))
	mBase = m.M
	m.G0 = v370 + int32(32)
	goto L80
L86:
	;
	v432 = int32(295)
	v434 = m.G0
	v436 = v434 - int32(32)
	m.G0 = v436
	switch int32(297) {
	case 0, 2:
		v446 = v432
		goto L93
	default:
		goto L94
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+12)) = v413
	F_sigemptyset(m, v403+int32(16))
	mBase = m.M
	goto L90
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[5])) = v399
	v413 = int32(_a_F_WalReceiverMain_3)
	goto L87
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+24)) = int32(268435456)
	v427 = F___sigaction(m, int32(2), v403+int32(12), int32(0))
	mBase = m.M
	m.G0 = v403 + int32(32)
	goto L86
L92:
	;
	v465 = int32(-2)
	v467 = m.G0
	v469 = v467 - int32(32)
	m.G0 = v469
	switch int32(0) {
	case 0, 2:
		v479 = v465
		goto L99
	default:
		goto L100
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+12)) = v446
	F_sigemptyset(m, v436+int32(16))
	mBase = m.M
	goto L96
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[6])) = v432
	v446 = int32(_a_F_WalReceiverMain_3)
	goto L93
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+24)) = int32(268435456)
	v460 = F___sigaction(m, int32(15), v436+int32(12), int32(0))
	mBase = m.M
	m.G0 = v436 + int32(32)
	goto L92
L98:
	;
	v498 = int32(-2)
	v500 = m.G0
	v502 = v500 - int32(32)
	m.G0 = v502
	switch int32(0) {
	case 0, 2:
		v512 = v498
		goto L105
	default:
		goto L106
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+12)) = v479
	F_sigemptyset(m, v469+int32(16))
	mBase = m.M
	goto L102
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[7])) = v465
	v479 = int32(_a_F_WalReceiverMain_3)
	goto L99
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+24)) = int32(268435456)
	v493 = F___sigaction(m, int32(14), v469+int32(12), int32(0))
	mBase = m.M
	m.G0 = v469 + int32(32)
	goto L98
L104:
	;
	v531 = int32(917)
	v533 = m.G0
	v535 = v533 - int32(32)
	m.G0 = v535
	switch int32(919) {
	case 0, 2:
		v545 = v531
		goto L111
	default:
		goto L112
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+12)) = v512
	F_sigemptyset(m, v502+int32(16))
	mBase = m.M
	goto L108
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[8])) = v498
	v512 = int32(_a_F_WalReceiverMain_3)
	goto L105
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+24)) = int32(268435456)
	v526 = F___sigaction(m, int32(13), v502+int32(12), int32(0))
	mBase = m.M
	m.G0 = v502 + int32(32)
	goto L104
L110:
	;
	v564 = int32(-2)
	v566 = m.G0
	v568 = v566 - int32(32)
	m.G0 = v568
	switch int32(0) {
	case 0, 2:
		v578 = v564
		goto L117
	default:
		goto L118
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+12)) = v545
	F_sigemptyset(m, v535+int32(16))
	mBase = m.M
	goto L114
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[9])) = v531
	v545 = int32(_a_F_WalReceiverMain_3)
	goto L111
L114:
	;
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v535)+24)) = int32(268435456)
	v559 = F___sigaction(m, int32(10), v535+int32(12), int32(0))
	mBase = m.M
	m.G0 = v535 + int32(32)
	goto L110
L116:
	;
	v597 = int32(0)
	v599 = m.G0
	v601 = v599 - int32(32)
	m.G0 = v601
	switch int32(2) {
	case 0, 2:
		v611 = v597
		goto L123
	default:
		goto L124
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568)+12)) = v578
	F_sigemptyset(m, v568+int32(16))
	mBase = m.M
	goto L120
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[10])) = v564
	v578 = int32(_a_F_WalReceiverMain_3)
	goto L117
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568)+24)) = int32(268435456)
	v592 = F___sigaction(m, int32(12), v568+int32(12), int32(0))
	mBase = m.M
	m.G0 = v568 + int32(32)
	goto L116
L122:
	;
	F_load_file(m, int32(_a_F_WalReceiverMain_4), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L128
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601)+12)) = v611
	F_sigemptyset(m, v601+int32(16))
	mBase = m.M
	goto L125
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11])) = v597
	v611 = int32(_a_F_WalReceiverMain_3)
	goto L123
L125:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v601)+24)) = int32(268435457)
	v625 = F___sigaction(m, int32(17), v601+int32(12), int32(0))
	mBase = m.M
	m.G0 = v601 + int32(32)
	goto L122
L128:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	if v634 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_WalReceiverMain_5), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L601
	}
L132:
	;
	v641 = int32(0)
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644))))
	if v646 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v647 = v644
	goto L135
L134:
	;
	v647 = int32(_a_F_WalReceiverMain_6)
	goto L135
L135:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	v653 = m.T0[v652].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v79, int32(1), v641, v641, v647, v19+int32(324))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14])) = v653
	if v653 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+8))
	v659 = m.T0[v658].(func(*base.Module, int32) int32)(m, v653)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L1
	} else {
		goto L597
	}
L140:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v668 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+12))
	m.T0[v669].(func(*base.Module, int32, int32, int32))(m, v662, v19+int32(320), v19+int32(316))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v674 = base.AtomicRmwXchg32(m, v33, int32(0), int32(1))
	if v674 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_s_lock(m, v33, int32(_a_F_WalReceiverMain_0), int32(286), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	base.MemoryFill(m, v81, int32(0), int32(1024))
	if v659 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L144
L146:
	;
	goto L152
L147:
	;
	goto L148
L148:
	;
	v803 = v31 + int32(1128)
	base.MemoryFill(m, v803, int32(0), int32(255))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v19)+320))
	if v807 != 0 {
		goto L180
	} else {
		goto L181
	}
L149:
	;
	goto L148
L150:
	;
	v799 = F_strlen(m, v788)
	mBase = m.M
	goto L149
L152:
	;
	goto L153
L153:
	;
	v689 = int32(1023)
	if (v81^v659)&int32(3) != 0 {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v792 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v789))) = uint8(v792)
	goto L150
L155:
	;
	v773 = v768
	v774 = v769
	v775 = v770
	goto L176
L156:
	;
	if v763 == int32(0) {
		v788 = v761
		v789 = v762
		goto L154
	} else {
		goto L175
	}
L157:
	;
	v761 = v659
	v762 = v81
	v763 = v689
	goto L156
L158:
	;
	goto L159
L159:
	;
	v693 = int32(0)
	if base.B2i32(v659&int32(3) == v693)|int32(0) == v693 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v729 == int32(0) {
		v788 = v726
		v789 = v727
		goto L154
	} else {
		goto L169
	}
L161:
	;
	v705 = v659
	v706 = v81
	v707 = v689
	goto L164
L162:
	;
	goto L163
L163:
	;
	v726 = v659
	v727 = v81
	v728 = v689
	v729 = int32(1)
	goto L160
L164:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	*(*uint8)(unsafe.Add(mBase, uint32(v706))) = uint8(v709)
	if v709 == int32(0) {
		v768 = v705
		v769 = v706
		v770 = v707
		goto L155
	} else {
		goto L166
	}
L165:
	;
	v726 = v720
	v727 = v714
	v728 = v716
	v729 = v718
	goto L160
L166:
	;
	v713 = int32(1)
	v714 = v706 + v713
	v716 = v707 - v713
	v717 = int32(0)
	v718 = base.B2i32(v716 != v717)
	v720 = v705 + v713
	if v720&int32(3) == v717 {
		v726 = v720
		v727 = v714
		v728 = v716
		v729 = v718
		goto L160
	} else {
		goto L167
	}
L167:
	;
	if v716 != 0 {
		v705 = v720
		v706 = v714
		v707 = v716
		goto L164
	} else {
		goto L168
	}
L168:
	;
	goto L165
L169:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	if base.B2i32(v732 == int32(0))|base.B2i32(base.Ui32(v728) < base.Ui32(int32(4))) != 0 {
		v761 = v726
		v762 = v727
		v763 = v728
		goto L156
	} else {
		goto L170
	}
L170:
	;
	v739 = v726
	v740 = v727
	v741 = v728
	goto L171
L171:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	v747 = int32(-2139062144)
	if (int32(16843008)-v744|v744)&v747 != v747 {
		v768 = v739
		v769 = v740
		v770 = v741
		goto L155
	} else {
		goto L173
	}
L172:
	;
	v761 = v755
	v762 = v753
	v763 = v757
	goto L156
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = v744
	v752 = int32(4)
	v753 = v740 + v752
	v755 = v739 + v752
	v757 = v741 - v752
	if base.Ui32(int32(3)) < base.Ui32(v757) {
		v739 = v755
		v740 = v753
		v741 = v757
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v768 = v761
	v769 = v762
	v770 = v763
	goto L155
L176:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
	*(*uint8)(unsafe.Add(mBase, uint32(v774))) = uint8(v777)
	if v777 == int32(0) {
		v788 = v773
		v789 = v774
		goto L154
	} else {
		goto L178
	}
L177:
	;
	v788 = v784
	v789 = v782
	goto L154
L178:
	;
	v781 = int32(1)
	v782 = v774 + v781
	v784 = v773 + v781
	v786 = v775 - v781
	if v786 != 0 {
		v773 = v784
		v774 = v782
		v775 = v786
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	goto L186
L181:
	;
	goto L182
L182:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
	v928 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1453)) = uint8(v928)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1384)) = v927
	v931 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v31)+1456)), uint32(v931))
	if v659 != 0 {
		goto L214
	} else {
		goto L215
	}
L183:
	;
	goto L182
L184:
	;
	v924 = F_strlen(m, v913)
	mBase = m.M
	goto L183
L186:
	;
	goto L187
L187:
	;
	v814 = int32(254)
	if (v803^v807)&int32(3) != 0 {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v917 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v914))) = uint8(v917)
	goto L184
L189:
	;
	v898 = v893
	v899 = v894
	v900 = v895
	goto L210
L190:
	;
	if v888 == int32(0) {
		v913 = v886
		v914 = v887
		goto L188
	} else {
		goto L209
	}
L191:
	;
	v886 = v807
	v887 = v803
	v888 = v814
	goto L190
L192:
	;
	goto L193
L193:
	;
	v818 = int32(0)
	if base.B2i32(v807&int32(3) == v818)|int32(0) == v818 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	if v854 == int32(0) {
		v913 = v851
		v914 = v852
		goto L188
	} else {
		goto L203
	}
L195:
	;
	v830 = v807
	v831 = v803
	v832 = v814
	goto L198
L196:
	;
	goto L197
L197:
	;
	v851 = v807
	v852 = v803
	v853 = v814
	v854 = int32(1)
	goto L194
L198:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830))))
	*(*uint8)(unsafe.Add(mBase, uint32(v831))) = uint8(v834)
	if v834 == int32(0) {
		v893 = v830
		v894 = v831
		v895 = v832
		goto L189
	} else {
		goto L200
	}
L199:
	;
	v851 = v845
	v852 = v839
	v853 = v841
	v854 = v843
	goto L194
L200:
	;
	v838 = int32(1)
	v839 = v831 + v838
	v841 = v832 - v838
	v842 = int32(0)
	v843 = base.B2i32(v841 != v842)
	v845 = v830 + v838
	if v845&int32(3) == v842 {
		v851 = v845
		v852 = v839
		v853 = v841
		v854 = v843
		goto L194
	} else {
		goto L201
	}
L201:
	;
	if v841 != 0 {
		v830 = v845
		v831 = v839
		v832 = v841
		goto L198
	} else {
		goto L202
	}
L202:
	;
	goto L199
L203:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	if base.B2i32(v857 == int32(0))|base.B2i32(base.Ui32(v853) < base.Ui32(int32(4))) != 0 {
		v886 = v851
		v887 = v852
		v888 = v853
		goto L190
	} else {
		goto L204
	}
L204:
	;
	v864 = v851
	v865 = v852
	v866 = v853
	goto L205
L205:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v864)))
	v872 = int32(-2139062144)
	if (int32(16843008)-v869|v869)&v872 != v872 {
		v893 = v864
		v894 = v865
		v895 = v866
		goto L189
	} else {
		goto L207
	}
L206:
	;
	v886 = v880
	v887 = v878
	v888 = v882
	goto L190
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v865))) = v869
	v877 = int32(4)
	v878 = v865 + v877
	v880 = v864 + v877
	v882 = v866 - v877
	if base.Ui32(int32(3)) < base.Ui32(v882) {
		v864 = v880
		v865 = v878
		v866 = v882
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v893 = v886
	v894 = v887
	v895 = v888
	goto L189
L210:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898))))
	*(*uint8)(unsafe.Add(mBase, uint32(v899))) = uint8(v902)
	if v902 == int32(0) {
		v913 = v898
		v914 = v899
		goto L188
	} else {
		goto L212
	}
L211:
	;
	v913 = v909
	v914 = v907
	goto L188
L212:
	;
	v906 = int32(1)
	v907 = v899 + v906
	v909 = v898 + v906
	v911 = v900 - v906
	if v911 != 0 {
		v898 = v909
		v899 = v907
		v900 = v911
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	F_pfree(m, v659)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v19)+320))
	if v936 != 0 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L216
L218:
	;
	F_pfree(m, v936)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v944)+16))
	v946 = m.T0[v945].(func(*base.Module, int32, int32) int32)(m, v940, v19+int32(328))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L222
	}
L221:
	;
	goto L220
L222:
	;
	v949 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	v950 = *(*int64)(unsafe.Add(mBase, uint32(v949)))
	goto L223
L223:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+208)) = v950
	v953 = v19 + int32(272)
	v958 = F_pg_snprintf(m, v953, int32(32), int32(_a_F_WalReceiverMain_7), v19+int32(208))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946))))
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v953))))
	if base.B2i32(v962 == int32(0))|base.B2i32(v962 != v965) != 0 {
		v983 = v962
		v984 = v965
		goto L228
	} else {
		goto L229
	}
L225:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L596
	}
L226:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L1
	} else {
		goto L592
	}
L227:
	;
	if v983-v984 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L228:
	;
	goto L227
L229:
	;
	v968 = v946
	v969 = v953
	goto L230
L230:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969)+1)))
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v968)+1)))
	if v973 == int32(0) {
		v983 = v973
		v984 = v972
		goto L228
	} else {
		goto L232
	}
L231:
	;
	v983 = v973
	v984 = v972
	goto L228
L232:
	;
	v976 = int32(1)
	if v973 == v972 {
		v968 = v968 + v976
		v969 = v969 + v976
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v988 = int32(1)
	v993 = v324
	v999 = v988
	goto L237
L235:
	;
	v2470 = v946
	goto L236
L236:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L1
	} else {
		goto L587
	}
L237:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v19)+328))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	if base.Ui32(v1008) <= base.Ui32(v1007) {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	v2470 = v2427
	goto L236
L239:
	;
	v2236 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	if v2236 < int32(0) {
		goto L528
	} else {
		goto L529
	}
L240:
	;
	if v1196 == int32(0) {
		v2227 = v999
		goto L239
	} else {
		goto L525
	}
L241:
	;
	F_WalRcvFetchTimeLineHistoryFiles(m, v1008, v1007)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L1
	} else {
		goto L521
	}
L244:
	;
	if v325&v988 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1015 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+56))
	v1017 = m.T0[v1016].(func(*base.Module, int32) int32)(m, v1013)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+240)) = v993
	v1175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+232)) = uint8(v1175)
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+248)) = v1177
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+336)))
	if v1182 != 0 {
		goto L286
	} else {
		goto L287
	}
L248:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+192)) = base.I64_extend_i32_s(v1017)
	v1022 = v19 + int32(336)
	v1027 = F_pg_snprintf(m, v1022, int32(64), int32(_a_F_WalReceiverMain_8), v19+int32(192))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1032 = int32(0)
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+48))
	v1039 = m.T0[v1038].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1030, v1022, int32(1), v1032, v1032, v1032, v1032)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v1043 = base.AtomicRmwXchg32(m, v33, int32(0), int32(1))
	if v1043 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	F_s_lock(m, v33, int32(_a_F_WalReceiverMain_0), int32(364), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1050 = v19 + int32(336)
	goto L258
L254:
	;
	goto L253
L255:
	;
	v1170 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v33))), uint32(v1170))
	goto L247
L256:
	;
	v1167 = F_strlen(m, v1156)
	mBase = m.M
	goto L255
L258:
	;
	goto L259
L259:
	;
	v1057 = int32(63)
	if (v204^v1050)&int32(3) != 0 {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v1160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1157))) = uint8(v1160)
	goto L256
L261:
	;
	v1141 = v1136
	v1142 = v1137
	v1143 = v1138
	goto L282
L262:
	;
	if v1131 == int32(0) {
		v1156 = v1129
		v1157 = v1130
		goto L260
	} else {
		goto L281
	}
L263:
	;
	v1129 = v1050
	v1130 = v204
	v1131 = v1057
	goto L262
L264:
	;
	goto L265
L265:
	;
	v1061 = int32(0)
	if base.B2i32(v1050&int32(3) == v1061)|int32(0) == v1061 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	if v1097 == int32(0) {
		v1156 = v1094
		v1157 = v1095
		goto L260
	} else {
		goto L275
	}
L267:
	;
	v1073 = v1050
	v1074 = v204
	v1075 = v1057
	goto L270
L268:
	;
	goto L269
L269:
	;
	v1094 = v1050
	v1095 = v204
	v1096 = v1057
	v1097 = int32(1)
	goto L266
L270:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1074))) = uint8(v1077)
	if v1077 == int32(0) {
		v1136 = v1073
		v1137 = v1074
		v1138 = v1075
		goto L261
	} else {
		goto L272
	}
L271:
	;
	v1094 = v1088
	v1095 = v1082
	v1096 = v1084
	v1097 = v1086
	goto L266
L272:
	;
	v1081 = int32(1)
	v1082 = v1074 + v1081
	v1084 = v1075 - v1081
	v1085 = int32(0)
	v1086 = base.B2i32(v1084 != v1085)
	v1088 = v1073 + v1081
	if v1088&int32(3) == v1085 {
		v1094 = v1088
		v1095 = v1082
		v1096 = v1084
		v1097 = v1086
		goto L266
	} else {
		goto L273
	}
L273:
	;
	if v1084 != 0 {
		v1073 = v1088
		v1074 = v1082
		v1075 = v1084
		goto L270
	} else {
		goto L274
	}
L274:
	;
	goto L271
L275:
	;
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	if base.B2i32(v1100 == int32(0))|base.B2i32(base.Ui32(v1096) < base.Ui32(int32(4))) != 0 {
		v1129 = v1094
		v1130 = v1095
		v1131 = v1096
		goto L262
	} else {
		goto L276
	}
L276:
	;
	v1107 = v1094
	v1108 = v1095
	v1109 = v1096
	goto L277
L277:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1107)))
	v1115 = int32(-2139062144)
	if (int32(16843008)-v1112|v1112)&v1115 != v1115 {
		v1136 = v1107
		v1137 = v1108
		v1138 = v1109
		goto L261
	} else {
		goto L279
	}
L278:
	;
	v1129 = v1123
	v1130 = v1121
	v1131 = v1125
	goto L262
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1108))) = v1112
	v1120 = int32(4)
	v1121 = v1108 + v1120
	v1123 = v1107 + v1120
	v1125 = v1109 - v1120
	if base.Ui32(int32(3)) < base.Ui32(v1125) {
		v1107 = v1123
		v1108 = v1121
		v1109 = v1125
		goto L277
	} else {
		goto L280
	}
L280:
	;
	goto L278
L281:
	;
	v1136 = v1129
	v1137 = v1130
	v1138 = v1131
	goto L261
L282:
	;
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1141))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1142))) = uint8(v1145)
	if v1145 == int32(0) {
		v1156 = v1141
		v1157 = v1142
		goto L260
	} else {
		goto L284
	}
L283:
	;
	v1156 = v1152
	v1157 = v1150
	goto L260
L284:
	;
	v1149 = int32(1)
	v1150 = v1142 + v1149
	v1152 = v1141 + v1149
	v1154 = v1143 - v1149
	if v1154 != 0 {
		v1141 = v1152
		v1142 = v1150
		v1143 = v1154
		goto L282
	} else {
		goto L285
	}
L285:
	;
	goto L283
L286:
	;
	v1183 = v19 + int32(336)
	goto L288
L287:
	;
	v1183 = v1175
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+236)) = v1183
	v1186 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1190 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1190)+32))
	v1192 = m.T0[v1191].(func(*base.Module, int32, int32) int32)(m, v1186, v19+int32(232))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1196 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	if v1192 == int32(0) {
		goto L240
	} else {
		goto L291
	}
L291:
	;
	if v1196 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v1200
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+164)) = uint32(v993)
	v1204 = int64(base.Ui64(v993) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+160)) = uint32(v1204)
	if v999 != 0 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	goto L294
L294:
	;
	v1222 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L303
	}
L295:
	;
	v1208 = int32(_a_F_WalReceiverMain_9)
	goto L297
L296:
	;
	v1208 = int32(_a_F_WalReceiverMain_10)
	goto L297
L297:
	;
	F_errmsg(m, v1208, v19+int32(160))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	if v999 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1216 = int32(390)
	goto L301
L300:
	;
	v1216 = int32(394)
	goto L301
L301:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), v1216, int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	goto L294
L303:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[17])) = v1222
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[18])) = v1222
	F_initStringInfo(m, int32(_a_F_WalReceiverMain_11))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v1235 = m.G0
	v1236 = int32(16)
	v1237 = v1235 - v1236
	m.G0 = v1237
	F_gettimeofday(m, v1237)
	mBase = m.M
	v1240 = *(*int64)(unsafe.Add(mBase, uint32(v1237)))
	v1241 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1237)+8)))
	m.G0 = v1237 + v1236
	v1249 = v1241 + v1240*int64(1000000) - int64(946684800000000)
	goto L305
L305:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[19]))
	v1257 = base.B2i32(v1251 <= int32(0))
	if v1251 <= int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1258 = int64(9223372036854775807)
	goto L308
L307:
	;
	v1258 = v1249 + base.I64_extend_i32_u(v1251)*int64(1000)
	goto L308
L308:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[20])) = v1258
	if v1251 <= int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1268 = int64(9223372036854775807)
	goto L311
L310:
	;
	v1268 = v1249 + base.I64_extend_i32_u(int32(base.Ui32(v1251)>>(uint(int32(1))%32)))*int64(1000)
	goto L311
L311:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = v1268
	v1273 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[22]))
	v1277 = v1249 + base.I64_extend_i32_u(v1273)*int64(1000000)
	if v1273 <= int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1280 = int64(9223372036854775807)
	goto L314
L313:
	;
	v1280 = v1277
	goto L314
L314:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23])) = v1280
	if v1273 <= int32(0) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1286 = int64(9223372036854775807)
	goto L317
L316:
	;
	v1286 = v1277
	goto L317
L317:
	;
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])))
	if v1289 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1290 = v1286
	goto L320
L319:
	;
	v1290 = int64(9223372036854775807)
	goto L320
L320:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[25])) = v1290
	F_XLogWalRcvSendReply(m, int32(1), int32(0))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = int32(-1)
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])))
	if v1303 == int32(1) {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2171 = m.ExcPending
	if v2171 != 0 {
		goto L1
	} else {
		goto L517
	}
L324:
	;
	if v1313 != 0 {
		goto L328
	} else {
		goto L329
	}
L325:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[27]))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+316))
	v1311 = base.B2i32(v1309 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])) = uint8(v1311)
	v1313 = v1311
	goto L327
L326:
	;
	v1313 = int32(0)
	goto L327
L327:
	;
	goto L324
L328:
	;
	goto L331
L329:
	;
	goto L330
L330:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L513
	}
L331:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[28]))
	if v1331 != 0 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	goto L330
L333:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[29]))
	if v1335 != 0 {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	goto L335
L337:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[29])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1417 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+40))
	v1419 = m.T0[v1418].(func(*base.Module, int32, int32, int32) int32)(m, v1411, v19+int32(228), v19+int32(224))
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L358
	}
L340:
	;
	v1347 = m.G0
	v1348 = int32(16)
	v1349 = v1347 - v1348
	m.G0 = v1349
	F_gettimeofday(m, v1349)
	mBase = m.M
	v1352 = *(*int64)(unsafe.Add(mBase, uint32(v1349)))
	v1353 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1349)+8)))
	m.G0 = v1349 + v1348
	v1361 = v1353 + v1352*int64(1000000) - int64(946684800000000)
	goto L341
L341:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[19]))
	v1369 = base.B2i32(v1363 <= int32(0))
	if v1363 <= int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1370 = int64(9223372036854775807)
	goto L344
L343:
	;
	v1370 = v1361 + base.I64_extend_i32_u(v1363)*int64(1000)
	goto L344
L344:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[20])) = v1370
	if v1363 <= int32(0) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1380 = int64(9223372036854775807)
	goto L347
L346:
	;
	v1380 = v1361 + base.I64_extend_i32_u(int32(base.Ui32(v1363)>>(uint(int32(1))%32)))*int64(1000)
	goto L347
L347:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = v1380
	v1385 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[22]))
	v1389 = v1361 + base.I64_extend_i32_u(v1385)*int64(1000000)
	if v1385 <= int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1392 = int64(9223372036854775807)
	goto L350
L349:
	;
	v1392 = v1389
	goto L350
L350:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23])) = v1392
	if v1385 <= int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1398 = int64(9223372036854775807)
	goto L353
L352:
	;
	v1398 = v1389
	goto L353
L353:
	;
	v1401 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])))
	if v1401 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1402 = v1398
	goto L356
L355:
	;
	v1402 = int64(9223372036854775807)
	goto L356
L356:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[25])) = v1402
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	goto L339
L358:
	;
	if v1419 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	if int32(0) < v1419 {
		goto L363
	} else {
		goto L364
	}
L360:
	;
	goto L361
L361:
	;
	v2004 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[20]))
	v2006 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21]))
	v2008 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23]))
	v2010 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[25]))
	v2014 = m.G0
	v2015 = int32(16)
	v2016 = v2014 - v2015
	m.G0 = v2016
	F_gettimeofday(m, v2016)
	mBase = m.M
	v2019 = *(*int64)(unsafe.Add(mBase, uint32(v2016)))
	v2020 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2016)+8)))
	m.G0 = v2016 + v2015
	v2028 = v2020 + v2019*int64(1000000) - int64(946684800000000)
	goto L473
L362:
	;
	v1979 = int32(0)
	F_XLogWalRcvSendReply(m, v1979, v1979)
	mBase = m.M
	v1982 = m.ExcPending
	if v1982 != 0 {
		goto L1
	} else {
		goto L471
	}
L363:
	;
	v1424 = v1419
	goto L366
L364:
	;
	goto L365
L365:
	;
	v1932 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L460
	}
L366:
	;
	v1442 = m.G0
	v1443 = int32(16)
	v1444 = v1442 - v1443
	m.G0 = v1444
	F_gettimeofday(m, v1444)
	mBase = m.M
	v1447 = *(*int64)(unsafe.Add(mBase, uint32(v1444)))
	v1448 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1444)+8)))
	m.G0 = v1444 + v1443
	v1456 = v1448 + v1447*int64(1000000) - int64(946684800000000)
	goto L368
L367:
	;
	if v1908 == int32(0) {
		goto L362
	} else {
		goto L459
	}
L368:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[19]))
	if v1459 <= int32(0) {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = v1475
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[20])) = v1474
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v19)+228))
	v1481 = v1479 + int32(1)
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1479))))
	v1484 = v1482 - int32(107)
	if v1484 != 0 {
		goto L379
	} else {
		goto L380
	}
L370:
	;
	v1462 = int64(9223372036854775807)
	v1474 = v1462
	v1475 = v1462
	goto L369
L371:
	;
	goto L372
L372:
	;
	v1465 = int64(1000)
	v1474 = base.I64_extend_i32_u(v1459)*v1465 + v1456
	v1475 = base.I64_extend_i32_u(int32(base.Ui32(v1459)>>(uint(int32(1))%32)))*v1465 + v1456
	goto L369
L373:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1906 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+40))
	v1908 = m.T0[v1907].(func(*base.Module, int32, int32, int32) int32)(m, v1900, v19+int32(228), v19+int32(224))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L1
	} else {
		goto L457
	}
L374:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	v1870 = base.AtomicRmwXchg64(m, v1868, int32(1464), v1853)
	v1872 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	if v1872 < int32(0) {
		goto L373
	} else {
		goto L454
	}
L375:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L1
	} else {
		goto L450
	}
L376:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L1
	} else {
		goto L446
	}
L377:
	;
	v1816 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[17]))
	v1853 = v1816
	v1854 = v1497
	goto L374
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L442
	}
L379:
	;
	if v1484 != int32(12) {
		goto L375
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	if v1424 != int32(18) {
		goto L376
	} else {
		goto L435
	}
L382:
	;
	if base.Ui32(v1424) <= base.Ui32(int32(24)) {
		goto L378
	} else {
		goto L383
	}
L383:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1436)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+1428)) = int64(24)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1424)) = v1481
	v1496 = v19 + int32(1424)
	v1497 = F_pq_getmsgint64(m, v1496)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L384
	}
L384:
	;
	v1499 = F_pq_getmsgint64(m, v1496)
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	v1501 = F_pq_getmsgint64(m, v1496)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	F_ProcessWalSndrMessage(m, v1499, v1501)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	v1506 = v1424 - int32(25)
	if v1506 == int32(0) {
		goto L377
	} else {
		goto L388
	}
L388:
	;
	v1513 = v1497
	v1519 = v1506
	v1520 = v1479 + int32(25)
	goto L389
L389:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30]))
	v1530 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	if int32(0) <= v1530 {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	v1853 = v1772
	v1854 = v1772
	goto L374
L391:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[31])) = int32(0)
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[32])))
	v1568 = m.G0
	v1570 = v1568 - int32(16)
	m.G0 = v1570
	if v1565 != 0 {
		goto L400
	} else {
		goto L401
	}
L392:
	;
	v1534 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[33]))
	v1536 = base.I64_div_u_s(v1513, base.I64_extend_i32_s(v1528))
	if v1534 == v1536 {
		v1559 = v1528
		goto L391
	} else {
		goto L395
	}
L393:
	;
	v1546 = v1528
	goto L394
L394:
	;
	v1549 = base.I64_div_u_s(v1513, base.I64_extend_i32_s(v1546))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[33])) = v1549
	v1551 = F_XLogFileInit(m, v1549, v1489)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L398
	}
L395:
	;
	F_XLogWalRcvClose(m, v1489)
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30]))
	v1543 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	if int32(0) <= v1543 {
		v1559 = v1541
		goto L391
	} else {
		goto L397
	}
L397:
	;
	v1546 = v1541
	goto L394
L398:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[34])) = v1489
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16])) = v1551
	v1558 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30]))
	v1559 = v1558
	goto L391
L399:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v1584))) = int32(167772240)
	v1588 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	v1592 = base.I32_wrap_i64(v1513) & (v1559 - int32(1))
	if base.Ui32(v1559) < base.Ui32(v1519+v1592) {
		goto L403
	} else {
		goto L404
	}
L400:
	;
	F___clock_gettime(m, int32(1), v1570)
	mBase = m.M
	v1574 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1570)+8)))
	v1575 = *(*int64)(unsafe.Add(mBase, uint32(v1570)))
	v1579 = v1574 + v1575*int64(1000000000)
	goto L402
L401:
	;
	v1579 = int64(0)
	goto L402
L402:
	;
	m.G0 = v1570 + int32(16)
	goto L399
L403:
	;
	v1596 = v1559 - v1592
	goto L405
L404:
	;
	v1596 = v1519
	goto L405
L405:
	;
	v1598 = F_pwrite(m, v1588, v1520, v1596, base.I64_extend_i32_s(v1592))
	mBase = m.M
	v1600 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v1600))) = int32(0)
	v1606 = int32(1)
	v1607 = base.I64_extend_i32_s(v1598)
	v1611 = m.G0
	v1613 = v1611 - int32(16)
	m.G0 = v1613
	if v1579 != int64(0) {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	if v1598 <= int32(0) {
		goto L423
	} else {
		goto L424
	}
L407:
	;
	F___clock_gettime(m, int32(1), v1613)
	mBase = m.M
	v1619 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1613)+8)))
	v1620 = *(*int64)(unsafe.Add(mBase, uint32(v1613)))
	v1624 = v1619 + (v1620*int64(1000000000) - v1579)
	goto L410
L408:
	;
	goto L409
L409:
	;
	v1711 = int32(888)
	v1712 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[36])) = v1712 + base.I64_extend_i32_u(v1606)
	v1716 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[37]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[37])) = v1716 + v1607
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v1606, v1607)
	mBase = m.M
	v1721 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[38])) = uint8(v1721)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[39])) = uint8(v1721)
	m.G0 = v1613 + int32(16)
	goto L406
L410:
	;
	v1674 = int32(888)
	v1675 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[40])) = v1675 + v1624
	v1679 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[0]))
	v1686 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1679))|base.B2i32(int32(1)<<(uint(v1679)%32)&int32(_a_F_WalReceiverMain_12) == v1686) == v1686 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v1691 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41])) = v1691 + v1624
	v1695 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[38])) = uint8(v1695)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[42])) = uint8(v1695)
	goto L422
L421:
	;
	goto L422
L422:
	;
	goto L409
L423:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[31]))
	if v1732 == int32(0) {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	goto L425
L425:
	;
	v1772 = v1513 + v1607
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[17])) = v1772
	v1775 = v1519 - v1598
	if v1775 != 0 {
		v1513 = v1772
		v1519 = v1775
		v1520 = v1598 + v1520
		goto L389
	} else {
		goto L434
	}
L426:
	;
	v1736 = int32(51)
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[31])) = v1736
	v1739 = v1736
	goto L428
L427:
	;
	v1739 = v1732
	goto L428
L428:
	;
	v1741 = v19 + int32(1440)
	v1743 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[34]))
	v1745 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[33]))
	v1747 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30]))
	F_XLogFileName(m, v1741, v1743, v1745, v1747)
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L1
	} else {
		goto L429
	}
L429:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[31])) = v1739
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L1
	} else {
		goto L431
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v1596
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = v1592
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v1741
	F_errmsg(m, int32(_a_F_WalReceiverMain_13), v19+int32(144))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(953), int32(_a_F_WalReceiverMain_14))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L434:
	;
	goto L390
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1452)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+1444)) = int64(17)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1440)) = v1481
	v1784 = v19 + int32(1440)
	v1785 = F_pq_getmsgint64(m, v1784)
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	v1787 = F_pq_getmsgint64(m, v1784)
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v1789 = F_pq_getmsgbyte(m, v1784)
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	F_ProcessWalSndrMessage(m, v1785, v1787)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		goto L1
	} else {
		goto L439
	}
L439:
	;
	if v1789 == int32(0) {
		goto L373
	} else {
		goto L440
	}
L440:
	;
	F_XLogWalRcvSendReply(m, int32(1), int32(0))
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	goto L373
L442:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_15), int32(0))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(837), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L446:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_17), int32(0))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(861), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L450:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1482
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_18), v19+int32(32))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(882), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L1
	} else {
		goto L453
	}
L453:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L454:
	;
	v1876 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[33]))
	v1878 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30])))
	v1879 = base.I64_div_u_s(v1854, v1878)
	if v1876 == v1879 {
		goto L373
	} else {
		goto L455
	}
L455:
	;
	F_XLogWalRcvClose(m, v1489)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	goto L373
L457:
	;
	if int32(0) < v1908 {
		v1424 = v1908
		goto L366
	} else {
		goto L458
	}
L458:
	;
	goto L367
L459:
	;
	goto L365
L460:
	;
	if v1932 != 0 {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_19), int32(0))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L1
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	v1957 = int32(0)
	F_XLogWalRcvSendReply(m, v1957, v1957)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L1
	} else {
		goto L467
	}
L464:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v1938
	v1941 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[17]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+136)) = uint32(v1941)
	v1944 = int64(base.Ui64(v1941) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+132)) = uint32(v1944)
	F_errdetail(m, int32(_a_F_WalReceiverMain_20), v19+int32(128))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(475), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L466
	}
L466:
	;
	goto L463
L467:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	F_XLogWalRcvFlush(m, int32(0), v1963)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1971 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+36))
	m.T0[v1972].(func(*base.Module, int32, int32))(m, v1967, v19+int32(328))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v19)+328))
	F_WalRcvFetchTimeLineHistoryFiles(m, v1975, v1976)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	v2227 = v1957
	goto L239
L471:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	F_XLogWalRcvFlush(m, int32(0), v1984)
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L472
	}
L472:
	;
	goto L361
L473:
	;
	if v2006 < v2004 {
		goto L474
	} else {
		goto L475
	}
L474:
	;
	v2030 = v2006
	goto L476
L475:
	;
	v2030 = v2004
	goto L476
L476:
	;
	if v2008 < v2030 {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v2032 = v2008
	goto L479
L478:
	;
	v2032 = v2030
	goto L479
L479:
	;
	if v2010 < v2032 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2034 = v2010
	goto L482
L481:
	;
	v2034 = v2032
	goto L482
L482:
	;
	if v2034 <= v2028 {
		v2052 = int32(0)
		goto L484
	} else {
		goto L485
	}
L483:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v19)+224))
	v2057 = F_WaitLatchOrSocket(m, v2054, v2055, v2052, int32(83886094))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L488
	}
L484:
	;
	goto L483
L485:
	;
	v2040 = v2034 - v2028
	if base.B2i32(int64(0) < v2028)^base.B2i32(v2040 < v2034)|base.B2i32(int64(2147483646000) < v2040) != 0 {
		v2052 = int32(2147483647)
		goto L484
	} else {
		goto L486
	}
L486:
	;
	v2049 = base.I64_div_s(v2040+int64(999), int64(1000))
	v2052 = base.I32_wrap_i64(v2049)
	goto L484
L487:
	;
	if v2057&int32(8) != 0 {
		goto L497
	} else {
		goto L498
	}
L488:
	;
	if v2057&int32(1) == int32(0) {
		goto L487
	} else {
		goto L489
	}
L489:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v2064))) = int32(0)
	goto L490
L490:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[28]))
	if v2068 != 0 {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L1
	} else {
		goto L494
	}
L492:
	;
	goto L493
L493:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v31)+1472))
	if v2071 == int32(0) {
		goto L487
	} else {
		goto L495
	}
L494:
	;
	goto L493
L495:
	;
	v2074 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1472)) = v2074
	F_XLogWalRcvSendReply(m, int32(1), v2074)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	goto L487
L497:
	;
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L1
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = int32(-1)
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])))
	if v2125 == int32(1) {
		goto L509
	} else {
		goto L510
	}
L500:
	;
	v2088 = m.G0
	v2089 = int32(16)
	v2090 = v2088 - v2089
	m.G0 = v2090
	F_gettimeofday(m, v2090)
	mBase = m.M
	v2093 = *(*int64)(unsafe.Add(mBase, uint32(v2090)))
	v2094 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2090)+8)))
	m.G0 = v2090 + v2089
	v2102 = v2094 + v2093*int64(1000000) - int64(946684800000000)
	goto L501
L501:
	;
	v2104 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[20]))
	if v2104 <= v2102 {
		goto L323
	} else {
		goto L502
	}
L502:
	;
	v2107 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21]))
	if v2107 <= v2102 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = int64(9223372036854775807)
	goto L505
L504:
	;
	goto L505
L505:
	;
	v2112 = base.B2i32(v2107 <= v2102)
	F_XLogWalRcvSendReply(m, v2112, v2112)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(0))
	mBase = m.M
	v2117 = m.ExcPending
	if v2117 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	goto L499
L508:
	;
	if v2135 != 0 {
		goto L331
	} else {
		goto L512
	}
L509:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[27]))
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2130)+316))
	v2133 = base.B2i32(v2131 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])) = uint8(v2133)
	v2135 = v2133
	goto L511
L510:
	;
	v2135 = int32(0)
	goto L511
L511:
	;
	goto L508
L512:
	;
	goto L332
L513:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_21), int32(0))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(428), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L517:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_22), int32(0))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(573), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L521:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v19)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v2191
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v2193
	F_errmsg(m, int32(_a_F_WalReceiverMain_23), v19+int32(16))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L1
	} else {
		goto L523
	}
L523:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(337), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L525:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v2207
	F_errmsg(m, int32(_a_F_WalReceiverMain_24), v19+int32(176))
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(606), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	v2227 = v999
	goto L239
L528:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16])) = int32(-1)
	v2287 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L1
	} else {
		goto L538
	}
L529:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	F_XLogWalRcvFlush(m, int32(0), v2240)
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v2244
	v2247 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[33]))
	v2250 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30])))
	v2251 = base.I64_div_u_s(int64(4294967296), v2250)
	v2252 = base.I64_div_u_s(v2247, v2251)
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+116)) = uint32(v2252)
	v2255 = v2247 - v2251*v2252
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+120)) = uint32(v2255)
	v2258 = v19 + int32(1440)
	v2263 = F_pg_snprintf(m, v2258, int32(64), int32(_a_F_WalReceiverMain_25), v19+int32(112))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	v2267 = F_close(m, v2266)
	mBase = m.M
	if v2267 != 0 {
		goto L226
	} else {
		goto L532
	}
L532:
	;
	v2269 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[44]))
	if v2269 != int32(2) {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	F_XLogArchiveForceDone(m, v2258)
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L1
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	F_XLogArchiveNotify(m, v19+int32(1440))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L1
	} else {
		goto L537
	}
L536:
	;
	goto L528
L537:
	;
	goto L528
L538:
	;
	if v2287 != 0 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_26), int32(0))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L1
	} else {
		goto L542
	}
L540:
	;
	goto L541
L541:
	;
	v2299 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	v2300 = int32(1456)
	v2301 = v2299 + v2300
	v2304 = base.AtomicRmwXchg32(m, v2299, v2300, int32(1))
	if v2304 != 0 {
		goto L544
	} else {
		goto L545
	}
L542:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(635), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	goto L541
L544:
	;
	F_s_lock(m, v2301, int32(_a_F_WalReceiverMain_0), int32(650), int32(_a_F_WalReceiverMain_27))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	v2310 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+8))
	if v2310 != int32(2) {
		goto L548
	} else {
		goto L549
	}
L547:
	;
	goto L546
L548:
	;
	v2313 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2301))), uint32(v2313))
	if v2310 == int32(5) {
		goto L225
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v2331 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2299)+40)) = v2331
	*(*int64)(unsafe.Add(mBase, uint32(v2299)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2299)+8)) = int32(3)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2299)+1456)), uint32(v2331))
	v2341 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[45]))
	F_SetLatch(m, v2341+int32(4))
	mBase = m.M
	goto L555
L551:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2321 = m.ExcPending
	if v2321 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_28), int32(0))
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L1
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(658), int32(_a_F_WalReceiverMain_27))
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	goto L556
L556:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	*(*int32)(unsafe.Add(mBase, uint32(v2362))) = int32(0)
	goto L558
L557:
	;
	v2396 = *(*int64)(unsafe.Add(mBase, uint32(v2299)+32))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+332)) = v2397
	*(*int32)(unsafe.Add(mBase, uint32(v2299)+8)) = int32(2)
	v2401 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2299)+1456)), uint32(v2401))
	v2405 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[46])))
	if v2405 != 0 {
		goto L572
	} else {
		goto L573
	}
L558:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[28]))
	if v2366 != 0 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L1
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	v2371 = base.AtomicRmwXchg32(m, v2301, int32(0), int32(1))
	if v2371 != 0 {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	goto L561
L563:
	;
	F_s_lock(m, v2301, int32(_a_F_WalReceiverMain_0), int32(678), int32(_a_F_WalReceiverMain_27))
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L1
	} else {
		goto L566
	}
L564:
	;
	goto L565
L565:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+8))
	switch v2377 - int32(4) {
	case 0:
		goto L567
	case 1:
		goto L569
	default:
		goto L568
	}
L566:
	;
	goto L565
L567:
	;
	goto L557
L568:
	;
	v2386 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2301))), uint32(v2386))
	v2390 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	v2394 = F_WaitLatch(m, v2390, int32(33), v2386, int32(134217782))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	v2380 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2301))), uint32(v2380))
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L571:
	;
	goto L556
L572:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+84)) = uint32(v2396)
	v2408 = int64(base.Ui64(v2396) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+80)) = uint32(v2408)
	v2411 = v19 + int32(1440)
	v2416 = F_pg_snprintf(m, v2411, int32(50), int32(_a_F_WalReceiverMain_29), v19+int32(80))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v2425 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v2425)+16))
	v2427 = m.T0[v2426].(func(*base.Module, int32, int32) int32)(m, v2421, v19+int32(328))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L576
	}
L575:
	;
	v2418 = F_strlen(m, v2411)
	mBase = m.M
	goto L574
L576:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	v2431 = *(*int64)(unsafe.Add(mBase, uint32(v2430)))
	goto L577
L577:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v2431
	v2434 = v19 + int32(272)
	v2439 = F_pg_snprintf(m, v2434, int32(32), int32(_a_F_WalReceiverMain_7), v19-int32(-64))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v2443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2427))))
	v2446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2434))))
	if base.B2i32(v2443 == int32(0))|base.B2i32(v2443 != v2446) != 0 {
		v2464 = v2443
		v2465 = v2446
		goto L580
	} else {
		goto L581
	}
L579:
	;
	if v2464-v2465 == int32(0) {
		v993 = v2396
		v999 = v2227
		goto L237
	} else {
		goto L586
	}
L580:
	;
	goto L579
L581:
	;
	v2449 = v2427
	v2450 = v2434
	goto L582
L582:
	;
	v2453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2450)+1)))
	v2454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2449)+1)))
	if v2454 == int32(0) {
		v2464 = v2454
		v2465 = v2453
		goto L580
	} else {
		goto L584
	}
L583:
	;
	v2464 = v2454
	v2465 = v2453
	goto L580
L584:
	;
	v2457 = int32(1)
	if v2454 == v2453 {
		v2449 = v2449 + v2457
		v2450 = v2450 + v2457
		goto L582
	} else {
		goto L585
	}
L585:
	;
	goto L583
L586:
	;
	goto L238
L587:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_30), int32(0))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v2470
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v19 + int32(272)
	F_errdetail(m, int32(_a_F_WalReceiverMain_31), v19+int32(48))
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(326), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2509 = m.ExcPending
	if v2509 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L592:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2515 = m.ExcPending
	if v2515 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v19 + int32(1440)
	F_errmsg(m, int32(_a_F_WalReceiverMain_32), v19+int32(96))
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(622), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L597:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v647
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v2540
	F_errmsg(m, int32(_a_F_WalReceiverMain_33), v19)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(277), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L601:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_34), int32(0))
	mBase = m.M
	v2557 = m.ExcPending
	if v2557 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(265), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WalSndSegmentOpen(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(1136)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[1])))
	if v16 != int32(1) {
		v27 = v13
	} else {
		v20 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[2]))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
		v22 = base.I64_div_u_s(v20, v21)
		if l1 != v22 {
			v27 = v13
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[3]))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v25
			v27 = v25
		}
	}
	v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v27
	v31 = base.I64_div_u_s(int64(4294967296), v28)
	v32 = base.I64_div_u_s(l1, v31)
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+36)) = uint32(v32)
	v35 = l1 - v31*v32
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+40)) = uint32(v35)
	v38 = v10 + int32(112)
	v43 = F_pg_snprintf(m, v38, int32(1024), int32(_a_F_WalSndSegmentOpen_0), v10+int32(32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return
	} else {
		v46 = F_BasicOpenFile(m, v38, int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+1168)) = v46
			if v46 < int32(0) {
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[4]))
				if v52 == int32(44) {
					v76 = v10 + int32(48)
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v79 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[5]))
					F_XLogFileName(m, v76, v77, l1, v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_WalSndSegmentOpen[4])) = int32(44)
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
								F_errmsg(m, int32(_a_F_WalSndSegmentOpen_1), v10)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_WalSndSegmentOpen_2), int32(3089), int32(_a_F_WalSndSegmentOpen_3))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
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
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v38
							F_errmsg(m, int32(_a_F_WalSndSegmentOpen_4), v10+int32(16))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_WalSndSegmentOpen_2), int32(3095), int32(_a_F_WalSndSegmentOpen_3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
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
				m.G0 = v10 + int32(1136)
				return
			}
		}
	}
}
func F_WalSndWriteData(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	v5 = int32(_a_F_WalSndWriteData_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[0]))
	v7 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[1])) = v7
	*(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[2])) = v7
	v16 = m.G0
	v17 = int32(16)
	v18 = v16 - v17
	m.G0 = v18
	F_gettimeofday(m, v18)
	mBase = m.M
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	m.G0 = v18 + v17
	v30 = v22 + v21*int64(1000000) - int64(946684800000000)
	F_enlargeStringInfo(m, int32(_a_F_WalSndWriteData_0), int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		v35 = int32(_a_F_WalSndWriteData_1)
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[2]))
		v37 = int32(_a_F_WalSndWriteData_0)
		v38 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[0]))
		v40 = int64(56)
		v42 = int64(65280)
		v44 = int64(40)
		v47 = int64(16711680)
		v49 = int64(24)
		v51 = int64(4278190080)
		v53 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v36+v38))) = v30<<(uint(v40)%64) | v30&v42<<(uint(v44)%64) | (v30&v47<<(uint(v49)%64) | v30&v51<<(uint(v53)%64)) | (int64(base.Ui64(v30)>>(uint(v53)%64))&v51 | int64(base.Ui64(v30)>>(uint(v49)%64))&v47 | (int64(base.Ui64(v30)>>(uint(v44)%64))&v42 | int64(base.Ui64(v30)>>(uint(v40)%64))))
		*(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[2])) = v36 + int32(8)
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
		v83 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[0]))
		v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
		*(*int64)(unsafe.Add(mBase, uint32(v81)+17)) = v84
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
		v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
		v91 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[3]))
		v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
		m.T0[v92].(func(*base.Module, int32, int32, int32))(m, int32(100), v88, v89)
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			v96 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[4]))
			if v96 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[3]))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
					v102 = m.T0[v101].(func(*base.Module) int32)(m)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						if v102 == int32(0) {
							v107 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndWriteData[5]))
							v109 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[6]))
							v111 = base.I32_div_s(v109, int32(2))
							if v30 < v107+base.I64_extend_i32_s(v111)*int64(1000) {
								v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[3]))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
								v120 = m.T0[v119].(func(*base.Module) int32)(m)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									if v120 == int32(0) {
										return
									} else {
										F_ProcessPendingWrites(m)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								F_ProcessPendingWrites(m)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_WalSndShutdown(m)
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v100 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[3]))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
				v102 = m.T0[v101].(func(*base.Module) int32)(m)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					if v102 == int32(0) {
						v107 = *(*int64)(unsafe.Add(mBase, _c_F_WalSndWriteData[5]))
						v109 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[6]))
						v111 = base.I32_div_s(v109, int32(2))
						if v30 < v107+base.I64_extend_i32_s(v111)*int64(1000) {
							v118 = *(*int32)(unsafe.Add(mBase, _c_F_WalSndWriteData[3]))
							v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
							v120 = m.T0[v119].(func(*base.Module) int32)(m)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								if v120 == int32(0) {
									return
								} else {
									F_ProcessPendingWrites(m)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							F_ProcessPendingWrites(m)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_WalSndShutdown(m)
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
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
func F_WalWriterMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v456 int64
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v12 = int32(-1)
	v15 = v3
	v16 = v3
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v12 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v455 = int32(m.ExcTag)
	v456 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v455 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[0])) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v317 = v15
	v318 = v16
	goto L8
L8:
	;
	if v318 != 0 {
		goto L63
	} else {
		goto L64
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v28 = int32(914)
	v30 = m.G0
	v32 = v30 - int32(32)
	m.G0 = v32
	switch int32(916) {
	case 0, 2:
		v42 = v28
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v62 = int32(916)
	v64 = m.G0
	v66 = v64 - int32(32)
	m.G0 = v66
	switch int32(918) {
	case 0, 2:
		v76 = v62
		goto L17
	default:
		goto L18
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v42
	F_sigemptyset(m, v32+int32(16))
	mBase = m.M
	goto L14
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[1])) = v28
	v42 = int32(_a_F_WalWriterMain_0)
	goto L11
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = int32(268435456)
	v56 = F___sigaction(m, int32(1), v32+int32(12), int32(0))
	mBase = m.M
	m.G0 = v32 + int32(32)
	goto L10
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v96 = int32(916)
	v98 = m.G0
	v100 = v98 - int32(32)
	m.G0 = v100
	switch int32(918) {
	case 0, 2:
		v110 = v96
		goto L23
	default:
		goto L24
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v76
	F_sigemptyset(m, v66+int32(16))
	mBase = m.M
	goto L20
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[2])) = v62
	v76 = int32(_a_F_WalWriterMain_0)
	goto L17
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(268435456)
	v90 = F___sigaction(m, int32(2), v66+int32(12), int32(0))
	mBase = m.M
	m.G0 = v66 + int32(32)
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v130 = int32(-2)
	v132 = m.G0
	v134 = v132 - int32(32)
	m.G0 = v134
	switch int32(0) {
	case 0, 2:
		v144 = v130
		goto L29
	default:
		goto L30
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v110
	F_sigemptyset(m, v100+int32(16))
	mBase = m.M
	goto L26
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[3])) = v96
	v110 = int32(_a_F_WalWriterMain_0)
	goto L23
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+24)) = int32(268435456)
	v124 = F___sigaction(m, int32(15), v100+int32(12), int32(0))
	mBase = m.M
	m.G0 = v100 + int32(32)
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v164 = int32(-2)
	v166 = m.G0
	v168 = v166 - int32(32)
	m.G0 = v168
	switch int32(0) {
	case 0, 2:
		v178 = v164
		goto L35
	default:
		goto L36
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v144
	F_sigemptyset(m, v134+int32(16))
	mBase = m.M
	goto L32
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[4])) = v130
	v144 = int32(_a_F_WalWriterMain_0)
	goto L29
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = int32(268435456)
	v158 = F___sigaction(m, int32(14), v134+int32(12), int32(0))
	mBase = m.M
	m.G0 = v134 + int32(32)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v198 = int32(917)
	v200 = m.G0
	v202 = v200 - int32(32)
	m.G0 = v202
	switch int32(919) {
	case 0, 2:
		v212 = v198
		goto L41
	default:
		goto L42
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v178
	F_sigemptyset(m, v168+int32(16))
	mBase = m.M
	goto L38
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[5])) = v164
	v178 = int32(_a_F_WalWriterMain_0)
	goto L35
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = int32(268435456)
	v192 = F___sigaction(m, int32(13), v168+int32(12), int32(0))
	mBase = m.M
	m.G0 = v168 + int32(32)
	goto L34
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v232 = int32(-2)
	v234 = m.G0
	v236 = v234 - int32(32)
	m.G0 = v236
	switch int32(0) {
	case 0, 2:
		v246 = v232
		goto L47
	default:
		goto L48
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v212
	F_sigemptyset(m, v202+int32(16))
	mBase = m.M
	goto L44
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[6])) = v198
	v212 = int32(_a_F_WalWriterMain_0)
	goto L41
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+24)) = int32(268435456)
	v226 = F___sigaction(m, int32(10), v202+int32(12), int32(0))
	mBase = m.M
	m.G0 = v202 + int32(32)
	goto L40
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v266 = int32(0)
	v268 = m.G0
	v270 = v268 - int32(32)
	m.G0 = v270
	switch int32(2) {
	case 0, 2:
		v280 = v266
		goto L53
	default:
		goto L54
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+12)) = v246
	F_sigemptyset(m, v236+int32(16))
	mBase = m.M
	goto L50
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[7])) = v232
	v246 = int32(_a_F_WalWriterMain_0)
	goto L47
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236)+24)) = int32(268435456)
	v260 = F___sigaction(m, int32(12), v236+int32(12), int32(0))
	mBase = m.M
	m.G0 = v236 + int32(32)
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[8]))
	v305 = F_AllocSetContextCreateInternal(m, v300, int32(_a_F_WalWriterMain_1), int32(0), int32(_a_F_WalWriterMain_2), int32(_a_F_WalWriterMain_3))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L5
	} else {
		goto L58
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v280
	F_sigemptyset(m, v270+int32(16))
	mBase = m.M
	goto L55
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[9])) = v266
	v280 = int32(_a_F_WalWriterMain_0)
	goto L53
L55:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270)+24)) = int32(268435457)
	v294 = F___sigaction(m, int32(17), v270+int32(12), int32(0))
	mBase = m.M
	m.G0 = v270 + int32(32)
	goto L52
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10])) = v305
	goto L59
L59:
	;
	v311 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v311)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v9 + int32(12)
	goto L62
L60:
	;
	v317 = v305
	v318 = int32(0)
	goto L8
L62:
	;
	goto L60
L63:
	;
	v319 = int32(_a_F_WalWriterMain_4)
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[11])) = v321 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[12])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_EmitErrorReport(m)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[13])) = v9 + int32(16)
	F_pgmem_sigprocmask(m, int32(_a_F_WalWriterMain_5), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L77
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_LWLockReleaseAll(m)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L5
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v338))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_UnlockBuffers(m)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_smgrdestroyall(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10])) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_FlushErrorState(m)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_MemoryContextReset(m, v317)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v371 = int32(_a_F_WalWriterMain_4)
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[11])) = v373 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_pg_usleep(m, int32(_a_F_WalWriterMain_6))
	mBase = m.M
	goto L65
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_SetWalWriterSleeping(m, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[15]))
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+60)) = v397
	v401 = int32(0)
	v404 = int32(50)
	goto L79
L79:
	;
	v407 = base.B2i32(v404 < int32(2))
	if v407 != v401&int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_SetWalWriterSleeping(m, v407)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L84
	}
L82:
	;
	v414 = v401
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v417))) = int32(0)
	goto L85
L84:
	;
	v414 = v407
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_ProcessMainLoopInterrupts(m)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	v424 = F_XLogBackgroundFlush(m)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v317
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[18]))
	if v424 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v442 = int32(50)
	goto L91
L90:
	;
	v442 = v404 - base.B2i32(int32(0) < v404)
	goto L91
L91:
	;
	if int32(0) < v442 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v445 = v435
	goto L94
L93:
	;
	v445 = v435 * int32(25)
	goto L94
L94:
	;
	v447 = F_WaitLatch(m, v432, int32(41), v445, int32(83886097))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v401 = v414
	v404 = v442
	goto L79
L96:
	;
	v460 = int32(v456)
	m.G0 = v9
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	if v9+int32(12) == v466 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	m.ExcPending = 1
	goto L105
L98:
	;
	if v470 != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	v470 = v468
	goto L101
L100:
	;
	v470 = int32(0)
	goto L101
L101:
	;
	goto L98
L102:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v12 = v470
	v15 = v471
	v16 = v462
	goto L1
L103:
	;
	goto L104
L104:
	;
	F___wasm_longjmp(m, v463, v462)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	return
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
