package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
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
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v322 int64
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v343 int64
	_ = v343
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int64
	_ = v398
	var v400 int64
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v422 int32
	_ = v422
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int64
	_ = v444
	var v446 int64
	_ = v446
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int64
	_ = v536
	var v538 int64
	_ = v538
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int64
	_ = v582
	var v584 int64
	_ = v584
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v606 int32
	_ = v606
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int64
	_ = v628
	var v630 int64
	_ = v630
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v652 int32
	_ = v652
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int64
	_ = v674
	var v676 int64
	_ = v676
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v698 int32
	_ = v698
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int64
	_ = v720
	var v722 int64
	_ = v722
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int64
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1092 int64
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1266 int32
	_ = v1266
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1302 int64
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1320 int64
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int64
	_ = v1338
	var v1339 int64
	_ = v1339
	var v1347 int64
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1356 int64
	_ = v1356
	var v1366 int64
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1375 int64
	_ = v1375
	var v1378 int64
	_ = v1378
	var v1384 int64
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1388 int64
	_ = v1388
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int64
	_ = v1450
	var v1451 int64
	_ = v1451
	var v1459 int64
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1467 int32
	_ = v1467
	var v1468 int64
	_ = v1468
	var v1478 int64
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1487 int64
	_ = v1487
	var v1490 int64
	_ = v1490
	var v1496 int64
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1500 int64
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1545 int64
	_ = v1545
	var v1546 int64
	_ = v1546
	var v1554 int64
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1560 int64
	_ = v1560
	var v1563 int64
	_ = v1563
	var v1572 int64
	_ = v1572
	var v1573 int64
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1594 int32
	_ = v1594
	var v1595 int64
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int64
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int64
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1611 int64
	_ = v1611
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1632 int64
	_ = v1632
	var v1634 int64
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1647 int64
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1672 int64
	_ = v1672
	var v1673 int64
	_ = v1673
	var v1677 int64
	_ = v1677
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1704 int32
	_ = v1704
	var v1705 int64
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1717 int64
	_ = v1717
	var v1718 int64
	_ = v1718
	var v1722 int64
	_ = v1722
	var v1772 int32
	_ = v1772
	var v1773 int64
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1784 int32
	_ = v1784
	var v1789 int64
	_ = v1789
	var v1793 int32
	_ = v1793
	var v1809 int32
	_ = v1809
	var v1810 int64
	_ = v1810
	var v1814 int64
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1841 int32
	_ = v1841
	var v1843 int64
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1863 int32
	_ = v1863
	var v1868 int32
	_ = v1868
	var v1870 int64
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1882 int32
	_ = v1882
	var v1883 int64
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int64
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1914 int64
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1930 int32
	_ = v1930
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1943 int32
	_ = v1943
	var v1948 int32
	_ = v1948
	var v1951 int64
	_ = v1951
	var v1952 int64
	_ = v1952
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1973 int64
	_ = v1973
	var v1975 int64
	_ = v1975
	var v1976 int64
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1997 int32
	_ = v1997
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int64
	_ = v2038
	var v2041 int64
	_ = v2041
	var v2047 int32
	_ = v2047
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2101 int64
	_ = v2101
	var v2103 int64
	_ = v2103
	var v2105 int64
	_ = v2105
	var v2107 int64
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2116 int64
	_ = v2116
	var v2117 int64
	_ = v2117
	var v2125 int64
	_ = v2125
	var v2127 int64
	_ = v2127
	var v2129 int64
	_ = v2129
	var v2131 int64
	_ = v2131
	var v2137 int64
	_ = v2137
	var v2146 int64
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2190 int64
	_ = v2190
	var v2191 int64
	_ = v2191
	var v2199 int64
	_ = v2199
	var v2201 int64
	_ = v2201
	var v2204 int64
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2222 int32
	_ = v2222
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2259 int32
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2280 int32
	_ = v2280
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2296 int32
	_ = v2296
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2310 int32
	_ = v2310
	var v2315 int32
	_ = v2315
	var v2325 int32
	_ = v2325
	var v2333 int32
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2344 int64
	_ = v2344
	var v2347 int64
	_ = v2347
	var v2348 int64
	_ = v2348
	var v2349 int64
	_ = v2349
	var v2352 int64
	_ = v2352
	var v2355 int32
	_ = v2355
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2366 int32
	_ = v2366
	var v2370 int32
	_ = v2370
	var v2374 int32
	_ = v2374
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2417 int32
	_ = v2417
	var v2421 int32
	_ = v2421
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2436 int32
	_ = v2436
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int64
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2494 int32
	_ = v2494
	var v2497 int64
	_ = v2497
	var v2500 int32
	_ = v2500
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2519 int32
	_ = v2519
	var v2520 int64
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2559 int32
	_ = v2559
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2584 int32
	_ = v2584
	var v2593 int32
	_ = v2593
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2612 int32
	_ = v2612
	var v2617 int32
	_ = v2617
	var v2620 int32
	_ = v2620
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2638 int32
	_ = v2638
	var v2642 int32
	_ = v2642
	var v2646 int32
	_ = v2646
	var v2651 int32
	_ = v2651
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
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1456)) = int32(1)
	v36 = v31 + int32(1456)
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_s_lock(m, v36, int32(_a_F_WalReceiverMain_0), int32(188), int32(_a_F_WalReceiverMain_1))
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
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[2]))
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1453)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v70
	v77 = v19 + int32(400)
	v79 = v31 + int32(104)
	goto L19
L8:
	;
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v54
	F_errstart_cold(m, int32(23), v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1456)) = int32(0)
	F_ConditionVariableBroadcast(m, v31+int32(12))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	v53 = m.ExcPending
	if v53 != 0 {
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
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(213), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
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
	v200 = v19 + int32(336)
	v202 = v31 + int32(1388)
	goto L50
L17:
	;
	v196 = F_strlen(m, v185)
	mBase = m.M
	goto L16
L19:
	;
	goto L20
L20:
	;
	v86 = int32(1023)
	if (v77^v79)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v189)
	goto L17
L22:
	;
	v170 = v165
	v171 = v166
	v172 = v167
	goto L43
L23:
	;
	if v160 == int32(0) {
		v185 = v158
		v186 = v159
		goto L21
	} else {
		goto L42
	}
L24:
	;
	v158 = v79
	v159 = v77
	v160 = v86
	goto L23
L25:
	;
	goto L26
L26:
	;
	v90 = int32(0)
	if base.B2i32(v79&int32(3) == v90)|int32(0) == v90 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v126 == int32(0) {
		v185 = v123
		v186 = v124
		goto L21
	} else {
		goto L36
	}
L28:
	;
	v102 = v79
	v103 = v77
	v104 = v86
	goto L31
L29:
	;
	goto L30
L30:
	;
	v123 = v79
	v124 = v77
	v125 = v86
	v126 = int32(1)
	goto L27
L31:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v106)
	if v106 == int32(0) {
		v165 = v102
		v166 = v103
		v167 = v104
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v123 = v117
	v124 = v111
	v125 = v113
	v126 = v115
	goto L27
L33:
	;
	v110 = int32(1)
	v111 = v103 + v110
	v113 = v104 - v110
	v114 = int32(0)
	v115 = base.B2i32(v113 != v114)
	v117 = v102 + v110
	if v117&int32(3) == v114 {
		v123 = v117
		v124 = v111
		v125 = v113
		v126 = v115
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v113 != 0 {
		v102 = v117
		v103 = v111
		v104 = v113
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if base.B2i32(v129 == int32(0))|base.B2i32(base.Ui32(v125) < base.Ui32(int32(4))) != 0 {
		v158 = v123
		v159 = v124
		v160 = v125
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v136 = v123
	v137 = v124
	v138 = v125
	goto L38
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v144 = int32(-2139062144)
	if (int32(16843008)-v141|v141)&v144 != v144 {
		v165 = v136
		v166 = v137
		v167 = v138
		goto L22
	} else {
		goto L40
	}
L39:
	;
	v158 = v152
	v159 = v150
	v160 = v154
	goto L23
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v141
	v149 = int32(4)
	v150 = v137 + v149
	v152 = v136 + v149
	v154 = v138 - v149
	if base.Ui32(int32(3)) < base.Ui32(v154) {
		v136 = v152
		v137 = v150
		v138 = v154
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v165 = v158
	v166 = v159
	v167 = v160
	goto L22
L43:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
	if v174 == int32(0) {
		v185 = v170
		v186 = v171
		goto L21
	} else {
		goto L45
	}
L44:
	;
	v185 = v181
	v186 = v179
	goto L21
L45:
	;
	v178 = int32(1)
	v179 = v171 + v178
	v181 = v170 + v178
	v183 = v172 - v178
	if v183 != 0 {
		v170 = v181
		v171 = v179
		v172 = v183
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v31)+32))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1452)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+332)) = v324
	v329 = m.G0
	v330 = int32(16)
	v331 = v329 - v330
	m.G0 = v331
	F_gettimeofday(m, v331)
	mBase = m.M
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v331)))
	v335 = int64(*(*int32)(unsafe.Add(mBase, uint32(v331)+8)))
	m.G0 = v331 + v330
	v343 = v335 + v334*int64(1000000) - int64(946684800000000)
	goto L78
L48:
	;
	v319 = F_strlen(m, v308)
	mBase = m.M
	goto L47
L50:
	;
	goto L51
L51:
	;
	v209 = int32(63)
	if (v200^v202)&int32(3) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v312 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v312)
	goto L48
L53:
	;
	v293 = v288
	v294 = v289
	v295 = v290
	goto L74
L54:
	;
	if v283 == int32(0) {
		v308 = v281
		v309 = v282
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v281 = v202
	v282 = v200
	v283 = v209
	goto L54
L56:
	;
	goto L57
L57:
	;
	v213 = int32(0)
	if base.B2i32(v202&int32(3) == v213)|int32(0) == v213 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v249 == int32(0) {
		v308 = v246
		v309 = v247
		goto L52
	} else {
		goto L67
	}
L59:
	;
	v225 = v202
	v226 = v200
	v227 = v209
	goto L62
L60:
	;
	goto L61
L61:
	;
	v246 = v202
	v247 = v200
	v248 = v209
	v249 = int32(1)
	goto L58
L62:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v229)
	if v229 == int32(0) {
		v288 = v225
		v289 = v226
		v290 = v227
		goto L53
	} else {
		goto L64
	}
L63:
	;
	v246 = v240
	v247 = v234
	v248 = v236
	v249 = v238
	goto L58
L64:
	;
	v233 = int32(1)
	v234 = v226 + v233
	v236 = v227 - v233
	v237 = int32(0)
	v238 = base.B2i32(v236 != v237)
	v240 = v225 + v233
	if v240&int32(3) == v237 {
		v246 = v240
		v247 = v234
		v248 = v236
		v249 = v238
		goto L58
	} else {
		goto L65
	}
L65:
	;
	if v236 != 0 {
		v225 = v240
		v226 = v234
		v227 = v236
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if base.B2i32(v252 == int32(0))|base.B2i32(base.Ui32(v248) < base.Ui32(int32(4))) != 0 {
		v281 = v246
		v282 = v247
		v283 = v248
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v259 = v246
	v260 = v247
	v261 = v248
	goto L69
L69:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v267 = int32(-2139062144)
	if (int32(16843008)-v264|v264)&v267 != v267 {
		v288 = v259
		v289 = v260
		v290 = v261
		goto L53
	} else {
		goto L71
	}
L70:
	;
	v281 = v275
	v282 = v273
	v283 = v277
	goto L54
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v264
	v272 = int32(4)
	v273 = v260 + v272
	v275 = v259 + v272
	v277 = v261 - v272
	if base.Ui32(int32(3)) < base.Ui32(v277) {
		v259 = v275
		v260 = v273
		v261 = v277
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v288 = v281
	v289 = v282
	v290 = v283
	goto L53
L74:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	*(*uint8)(unsafe.Add(mBase, uint32(v294))) = uint8(v297)
	if v297 == int32(0) {
		v308 = v293
		v309 = v294
		goto L52
	} else {
		goto L76
	}
L75:
	;
	v308 = v304
	v309 = v302
	goto L52
L76:
	;
	v301 = int32(1)
	v302 = v294 + v301
	v304 = v293 + v301
	v306 = v295 - v301
	if v306 != 0 {
		v293 = v304
		v294 = v302
		v295 = v306
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+80)) = v343
	*(*int64)(unsafe.Add(mBase, uint32(v31)+96)) = v343
	*(*int64)(unsafe.Add(mBase, uint32(v31)+72)) = v343
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1456)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v348
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v353)+1464)) = int64(0)
	F_on_shmem_exit(m, int32(1029), v19+int32(332))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v362 = int32(914)
	v364 = m.G0
	v366 = v364 - int32(32)
	m.G0 = v366
	switch int32(916) {
	case 0, 2:
		v376 = v362
		goto L81
	default:
		goto L82
	}
L80:
	;
	v408 = int32(-2)
	v410 = m.G0
	v412 = v410 - int32(32)
	m.G0 = v412
	switch int32(0) {
	case 0, 2:
		v422 = v408
		goto L94
	default:
		goto L95
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v376
	F_sigemptyset(m, v366+int32(16))
	mBase = m.M
	goto L84
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[4])) = v362
	v376 = int32(_a_F_WalReceiverMain_3)
	goto L81
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = int32(268435456)
	v388 = v366 + int32(12)
	goto L88
L86:
	;
	m.G0 = v366 + int32(32)
	goto L80
L88:
	;
	goto L89
L89:
	;
	if v388 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v394 = int32(20)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[5])) = v396
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v388)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[6])) = v398
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v388)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[7])) = v400
	goto L92
L91:
	;
	goto L92
L92:
	;
	goto L86
L93:
	;
	v454 = int32(295)
	v456 = m.G0
	v458 = v456 - int32(32)
	m.G0 = v458
	switch int32(297) {
	case 0, 2:
		v468 = v454
		goto L107
	default:
		goto L108
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v412)+12)) = v422
	F_sigemptyset(m, v412+int32(16))
	mBase = m.M
	goto L97
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[8])) = v408
	v422 = int32(_a_F_WalReceiverMain_3)
	goto L94
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v412)+24)) = int32(268435456)
	v434 = v412 + int32(12)
	goto L101
L99:
	;
	m.G0 = v412 + int32(32)
	goto L93
L101:
	;
	goto L102
L102:
	;
	if v434 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v441 = int32(40)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v434)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[9])) = v442
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v434)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[10])) = v444
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v434)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[11])) = v446
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L99
L106:
	;
	v500 = int32(-2)
	v502 = m.G0
	v504 = v502 - int32(32)
	m.G0 = v504
	switch int32(0) {
	case 0, 2:
		v514 = v500
		goto L120
	default:
		goto L121
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458)+12)) = v468
	F_sigemptyset(m, v458+int32(16))
	mBase = m.M
	goto L110
L108:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12])) = v454
	v468 = int32(_a_F_WalReceiverMain_3)
	goto L107
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458)+24)) = int32(268435456)
	v480 = v458 + int32(12)
	goto L114
L112:
	;
	m.G0 = v458 + int32(32)
	goto L106
L114:
	;
	goto L115
L115:
	;
	if v480 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v487 = int32(300)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v480)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13])) = v488
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v480)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[14])) = v490
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v480)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[15])) = v492
	goto L118
L117:
	;
	goto L118
L118:
	;
	goto L112
L119:
	;
	v546 = int32(-2)
	v548 = m.G0
	v550 = v548 - int32(32)
	m.G0 = v550
	switch int32(0) {
	case 0, 2:
		v560 = v546
		goto L133
	default:
		goto L134
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504)+12)) = v514
	F_sigemptyset(m, v504+int32(16))
	mBase = m.M
	goto L123
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16])) = v500
	v514 = int32(_a_F_WalReceiverMain_3)
	goto L120
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504)+24)) = int32(268435456)
	v526 = v504 + int32(12)
	goto L127
L125:
	;
	m.G0 = v504 + int32(32)
	goto L119
L127:
	;
	goto L128
L128:
	;
	if v526 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v533 = int32(280)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v526)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[17])) = v534
	v536 = *(*int64)(unsafe.Add(mBase, uint32(v526)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[18])) = v536
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v526)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[19])) = v538
	goto L131
L130:
	;
	goto L131
L131:
	;
	goto L125
L132:
	;
	v592 = int32(917)
	v594 = m.G0
	v596 = v594 - int32(32)
	m.G0 = v596
	switch int32(919) {
	case 0, 2:
		v606 = v592
		goto L146
	default:
		goto L147
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+12)) = v560
	F_sigemptyset(m, v550+int32(16))
	mBase = m.M
	goto L136
L134:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[20])) = v546
	v560 = int32(_a_F_WalReceiverMain_3)
	goto L133
L136:
	;
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+24)) = int32(268435456)
	v572 = v550 + int32(12)
	goto L140
L138:
	;
	m.G0 = v550 + int32(32)
	goto L132
L140:
	;
	goto L141
L141:
	;
	if v572 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v579 = int32(260)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v572)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = v580
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v572)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[22])) = v582
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v572)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23])) = v584
	goto L144
L143:
	;
	goto L144
L144:
	;
	goto L138
L145:
	;
	v638 = int32(-2)
	v640 = m.G0
	v642 = v640 - int32(32)
	m.G0 = v642
	switch int32(0) {
	case 0, 2:
		v652 = v638
		goto L159
	default:
		goto L160
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v596)+12)) = v606
	F_sigemptyset(m, v596+int32(16))
	mBase = m.M
	goto L149
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])) = v592
	v606 = int32(_a_F_WalReceiverMain_3)
	goto L146
L149:
	;
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v596)+24)) = int32(268435456)
	v618 = v596 + int32(12)
	goto L153
L151:
	;
	m.G0 = v596 + int32(32)
	goto L145
L153:
	;
	goto L154
L154:
	;
	if v618 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v625 = int32(200)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v618)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[25])) = v626
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v618)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])) = v628
	v630 = *(*int64)(unsafe.Add(mBase, uint32(v618)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[27])) = v630
	goto L157
L156:
	;
	goto L157
L157:
	;
	goto L151
L158:
	;
	v684 = int32(0)
	v686 = m.G0
	v688 = v686 - int32(32)
	m.G0 = v688
	switch int32(2) {
	case 0, 2:
		v698 = v684
		goto L172
	default:
		goto L173
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642)+12)) = v652
	F_sigemptyset(m, v642+int32(16))
	mBase = m.M
	goto L162
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[28])) = v638
	v652 = int32(_a_F_WalReceiverMain_3)
	goto L159
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642)+24)) = int32(268435456)
	v664 = v642 + int32(12)
	goto L166
L164:
	;
	m.G0 = v642 + int32(32)
	goto L158
L166:
	;
	goto L167
L167:
	;
	if v664 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v671 = int32(240)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v664)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[29])) = v672
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v664)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[30])) = v674
	v676 = *(*int64)(unsafe.Add(mBase, uint32(v664)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[31])) = v676
	goto L170
L169:
	;
	goto L170
L170:
	;
	goto L164
L171:
	;
	F_load_file(m, int32(_a_F_WalReceiverMain_4), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L184
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v688)+12)) = v698
	F_sigemptyset(m, v688+int32(16))
	mBase = m.M
	goto L174
L173:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[32])) = v684
	v698 = int32(_a_F_WalReceiverMain_3)
	goto L172
L174:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v688)+24)) = int32(268435457)
	v710 = v688 + int32(12)
	goto L179
L177:
	;
	m.G0 = v688 + int32(32)
	goto L171
L179:
	;
	goto L180
L180:
	;
	if v710 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v717 = int32(340)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v710)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[33])) = v718
	v720 = *(*int64)(unsafe.Add(mBase, uint32(v710)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[34])) = v720
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v710)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[35])) = v722
	goto L183
L182:
	;
	goto L183
L183:
	;
	goto L177
L184:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	if v734 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_sigprocmask(m, int32(_a_F_WalReceiverMain_5), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L657
	}
L188:
	;
	v741 = int32(0)
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[37]))
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	if v746 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v747 = v744
	goto L191
L190:
	;
	v747 = int32(_a_F_WalReceiverMain_6)
	goto L191
L191:
	;
	v751 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v753 = m.T0[v752].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v77, int32(1), v741, v741, v747, v19+int32(324))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38])) = v753
	if v753 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v757)+8))
	v759 = m.T0[v758].(func(*base.Module, int32) int32)(m, v753)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L1
	} else {
		goto L653
	}
L196:
	;
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v768 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v768)+12))
	m.T0[v769].(func(*base.Module, int32, int32, int32))(m, v762, v19+int32(320), v19+int32(316))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(1)
	if v772 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	F_s_lock(m, v36, int32(_a_F_WalReceiverMain_0), int32(286), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	base.MemoryFill(m, v79, int32(0), int32(1024))
	if v759 != 0 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	goto L200
L202:
	;
	goto L208
L203:
	;
	goto L204
L204:
	;
	v903 = v31 + int32(1128)
	base.MemoryFill(m, v903, int32(0), int32(255))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v19)+320))
	if v907 != 0 {
		goto L236
	} else {
		goto L237
	}
L205:
	;
	goto L204
L206:
	;
	v899 = F_strlen(m, v888)
	mBase = m.M
	goto L205
L208:
	;
	goto L209
L209:
	;
	v789 = int32(1023)
	if (v79^v759)&int32(3) != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v892 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v889))) = uint8(v892)
	goto L206
L211:
	;
	v873 = v868
	v874 = v869
	v875 = v870
	goto L232
L212:
	;
	if v863 == int32(0) {
		v888 = v861
		v889 = v862
		goto L210
	} else {
		goto L231
	}
L213:
	;
	v861 = v759
	v862 = v79
	v863 = v789
	goto L212
L214:
	;
	goto L215
L215:
	;
	v793 = int32(0)
	if base.B2i32(v759&int32(3) == v793)|int32(0) == v793 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	if v829 == int32(0) {
		v888 = v826
		v889 = v827
		goto L210
	} else {
		goto L225
	}
L217:
	;
	v805 = v759
	v806 = v79
	v807 = v789
	goto L220
L218:
	;
	goto L219
L219:
	;
	v826 = v759
	v827 = v79
	v828 = v789
	v829 = int32(1)
	goto L216
L220:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	*(*uint8)(unsafe.Add(mBase, uint32(v806))) = uint8(v809)
	if v809 == int32(0) {
		v868 = v805
		v869 = v806
		v870 = v807
		goto L211
	} else {
		goto L222
	}
L221:
	;
	v826 = v820
	v827 = v814
	v828 = v816
	v829 = v818
	goto L216
L222:
	;
	v813 = int32(1)
	v814 = v806 + v813
	v816 = v807 - v813
	v817 = int32(0)
	v818 = base.B2i32(v816 != v817)
	v820 = v805 + v813
	if v820&int32(3) == v817 {
		v826 = v820
		v827 = v814
		v828 = v816
		v829 = v818
		goto L216
	} else {
		goto L223
	}
L223:
	;
	if v816 != 0 {
		v805 = v820
		v806 = v814
		v807 = v816
		goto L220
	} else {
		goto L224
	}
L224:
	;
	goto L221
L225:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	if base.B2i32(v832 == int32(0))|base.B2i32(base.Ui32(v828) < base.Ui32(int32(4))) != 0 {
		v861 = v826
		v862 = v827
		v863 = v828
		goto L212
	} else {
		goto L226
	}
L226:
	;
	v839 = v826
	v840 = v827
	v841 = v828
	goto L227
L227:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v839)))
	v847 = int32(-2139062144)
	if (int32(16843008)-v844|v844)&v847 != v847 {
		v868 = v839
		v869 = v840
		v870 = v841
		goto L211
	} else {
		goto L229
	}
L228:
	;
	v861 = v855
	v862 = v853
	v863 = v857
	goto L212
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v840))) = v844
	v852 = int32(4)
	v853 = v840 + v852
	v855 = v839 + v852
	v857 = v841 - v852
	if base.Ui32(int32(3)) < base.Ui32(v857) {
		v839 = v855
		v840 = v853
		v841 = v857
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v868 = v861
	v869 = v862
	v870 = v863
	goto L211
L232:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873))))
	*(*uint8)(unsafe.Add(mBase, uint32(v874))) = uint8(v877)
	if v877 == int32(0) {
		v888 = v873
		v889 = v874
		goto L210
	} else {
		goto L234
	}
L233:
	;
	v888 = v884
	v889 = v882
	goto L210
L234:
	;
	v881 = int32(1)
	v882 = v874 + v881
	v884 = v873 + v881
	v886 = v875 - v881
	if v886 != 0 {
		v873 = v884
		v874 = v882
		v875 = v886
		goto L232
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	goto L242
L237:
	;
	goto L238
L238:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1456)) = int32(0)
	v1030 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1453)) = uint8(v1030)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1384)) = v1027
	if v759 != 0 {
		goto L270
	} else {
		goto L271
	}
L239:
	;
	goto L238
L240:
	;
	v1024 = F_strlen(m, v1013)
	mBase = m.M
	goto L239
L242:
	;
	goto L243
L243:
	;
	v914 = int32(254)
	if (v903^v907)&int32(3) != 0 {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v1017 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1014))) = uint8(v1017)
	goto L240
L245:
	;
	v998 = v993
	v999 = v994
	v1000 = v995
	goto L266
L246:
	;
	if v988 == int32(0) {
		v1013 = v986
		v1014 = v987
		goto L244
	} else {
		goto L265
	}
L247:
	;
	v986 = v907
	v987 = v903
	v988 = v914
	goto L246
L248:
	;
	goto L249
L249:
	;
	v918 = int32(0)
	if base.B2i32(v907&int32(3) == v918)|int32(0) == v918 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	if v954 == int32(0) {
		v1013 = v951
		v1014 = v952
		goto L244
	} else {
		goto L259
	}
L251:
	;
	v930 = v907
	v931 = v903
	v932 = v914
	goto L254
L252:
	;
	goto L253
L253:
	;
	v951 = v907
	v952 = v903
	v953 = v914
	v954 = int32(1)
	goto L250
L254:
	;
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930))))
	*(*uint8)(unsafe.Add(mBase, uint32(v931))) = uint8(v934)
	if v934 == int32(0) {
		v993 = v930
		v994 = v931
		v995 = v932
		goto L245
	} else {
		goto L256
	}
L255:
	;
	v951 = v945
	v952 = v939
	v953 = v941
	v954 = v943
	goto L250
L256:
	;
	v938 = int32(1)
	v939 = v931 + v938
	v941 = v932 - v938
	v942 = int32(0)
	v943 = base.B2i32(v941 != v942)
	v945 = v930 + v938
	if v945&int32(3) == v942 {
		v951 = v945
		v952 = v939
		v953 = v941
		v954 = v943
		goto L250
	} else {
		goto L257
	}
L257:
	;
	if v941 != 0 {
		v930 = v945
		v931 = v939
		v932 = v941
		goto L254
	} else {
		goto L258
	}
L258:
	;
	goto L255
L259:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951))))
	if base.B2i32(v957 == int32(0))|base.B2i32(base.Ui32(v953) < base.Ui32(int32(4))) != 0 {
		v986 = v951
		v987 = v952
		v988 = v953
		goto L246
	} else {
		goto L260
	}
L260:
	;
	v964 = v951
	v965 = v952
	v966 = v953
	goto L261
L261:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v964)))
	v972 = int32(-2139062144)
	if (int32(16843008)-v969|v969)&v972 != v972 {
		v993 = v964
		v994 = v965
		v995 = v966
		goto L245
	} else {
		goto L263
	}
L262:
	;
	v986 = v980
	v987 = v978
	v988 = v982
	goto L246
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v965))) = v969
	v977 = int32(4)
	v978 = v965 + v977
	v980 = v964 + v977
	v982 = v966 - v977
	if base.Ui32(int32(3)) < base.Ui32(v982) {
		v964 = v980
		v965 = v978
		v966 = v982
		goto L261
	} else {
		goto L264
	}
L264:
	;
	goto L262
L265:
	;
	v993 = v986
	v994 = v987
	v995 = v988
	goto L245
L266:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	*(*uint8)(unsafe.Add(mBase, uint32(v999))) = uint8(v1002)
	if v1002 == int32(0) {
		v1013 = v998
		v1014 = v999
		goto L244
	} else {
		goto L268
	}
L267:
	;
	v1013 = v1009
	v1014 = v1007
	goto L244
L268:
	;
	v1006 = int32(1)
	v1007 = v999 + v1006
	v1009 = v998 + v1006
	v1011 = v1000 - v1006
	if v1011 != 0 {
		v998 = v1009
		v999 = v1007
		v1000 = v1011
		goto L266
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	F_pfree(m, v759)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v19)+320))
	if v1035 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	goto L272
L274:
	;
	F_pfree(m, v1035)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v1043 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+16))
	v1045 = m.T0[v1044].(func(*base.Module, int32, int32) int32)(m, v1039, v19+int32(328))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L278
	}
L277:
	;
	goto L276
L278:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[39]))
	v1049 = *(*int64)(unsafe.Add(mBase, uint32(v1048)))
	goto L279
L279:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+208)) = v1049
	v1052 = v19 + int32(272)
	v1057 = F_pg_snprintf(m, v1052, int32(32), int32(_a_F_WalReceiverMain_7), v19+int32(208))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052))))
	if base.B2i32(v1061 == int32(0))|base.B2i32(v1061 != v1064) != 0 {
		v1082 = v1061
		v1083 = v1064
		goto L284
	} else {
		goto L285
	}
L281:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L652
	}
L282:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L1
	} else {
		goto L648
	}
L283:
	;
	if v1082-v1083 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L284:
	;
	goto L283
L285:
	;
	v1067 = v1045
	v1068 = v1052
	goto L286
L286:
	;
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068)+1)))
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+1)))
	if v1072 == int32(0) {
		v1082 = v1072
		v1083 = v1071
		goto L284
	} else {
		goto L288
	}
L287:
	;
	v1082 = v1072
	v1083 = v1071
	goto L284
L288:
	;
	v1075 = int32(1)
	if v1072 == v1071 {
		v1067 = v1067 + v1075
		v1068 = v1068 + v1075
		goto L286
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	v1087 = int32(1)
	v1092 = v322
	v1099 = v1087
	goto L293
L291:
	;
	v2559 = v1045
	goto L292
L292:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L1
	} else {
		goto L643
	}
L293:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v19)+328))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	if base.Ui32(v1107) <= base.Ui32(v1106) {
		goto L297
	} else {
		goto L298
	}
L294:
	;
	v2559 = v2516
	goto L292
L295:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	if v2333 < int32(0) {
		goto L584
	} else {
		goto L585
	}
L296:
	;
	if v1294 == int32(0) {
		v2325 = v1099
		goto L295
	} else {
		goto L581
	}
L297:
	;
	F_WalRcvFetchTimeLineHistoryFiles(m, v1107, v1106)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L1
	} else {
		goto L577
	}
L300:
	;
	if v323&v1087 != 0 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v1114 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+56))
	v1116 = m.T0[v1115].(func(*base.Module, int32) int32)(m, v1112)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+240)) = v1092
	v1273 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+232)) = uint8(v1273)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+248)) = v1275
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+336)))
	if v1280 != 0 {
		goto L342
	} else {
		goto L343
	}
L304:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+192)) = base.I64_extend_i32_s(v1116)
	v1121 = v19 + int32(336)
	v1126 = F_pg_snprintf(m, v1121, int32(64), int32(_a_F_WalReceiverMain_8), v19+int32(192))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v1131 = int32(0)
	v1136 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1136)+48))
	v1138 = m.T0[v1137].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1129, v1121, int32(1), v1131, v1131, v1131, v1131)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(1)
	if v1140 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	F_s_lock(m, v36, int32(_a_F_WalReceiverMain_0), int32(364), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v1149 = v19 + int32(336)
	goto L314
L310:
	;
	goto L309
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(0)
	goto L303
L312:
	;
	v1266 = F_strlen(m, v1255)
	mBase = m.M
	goto L311
L314:
	;
	goto L315
L315:
	;
	v1156 = int32(63)
	if (v202^v1149)&int32(3) != 0 {
		goto L319
	} else {
		goto L320
	}
L316:
	;
	v1259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1256))) = uint8(v1259)
	goto L312
L317:
	;
	v1240 = v1235
	v1241 = v1236
	v1242 = v1237
	goto L338
L318:
	;
	if v1230 == int32(0) {
		v1255 = v1228
		v1256 = v1229
		goto L316
	} else {
		goto L337
	}
L319:
	;
	v1228 = v1149
	v1229 = v202
	v1230 = v1156
	goto L318
L320:
	;
	goto L321
L321:
	;
	v1160 = int32(0)
	if base.B2i32(v1149&int32(3) == v1160)|int32(0) == v1160 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	if v1196 == int32(0) {
		v1255 = v1193
		v1256 = v1194
		goto L316
	} else {
		goto L331
	}
L323:
	;
	v1172 = v1149
	v1173 = v202
	v1174 = v1156
	goto L326
L324:
	;
	goto L325
L325:
	;
	v1193 = v1149
	v1194 = v202
	v1195 = v1156
	v1196 = int32(1)
	goto L322
L326:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1173))) = uint8(v1176)
	if v1176 == int32(0) {
		v1235 = v1172
		v1236 = v1173
		v1237 = v1174
		goto L317
	} else {
		goto L328
	}
L327:
	;
	v1193 = v1187
	v1194 = v1181
	v1195 = v1183
	v1196 = v1185
	goto L322
L328:
	;
	v1180 = int32(1)
	v1181 = v1173 + v1180
	v1183 = v1174 - v1180
	v1184 = int32(0)
	v1185 = base.B2i32(v1183 != v1184)
	v1187 = v1172 + v1180
	if v1187&int32(3) == v1184 {
		v1193 = v1187
		v1194 = v1181
		v1195 = v1183
		v1196 = v1185
		goto L322
	} else {
		goto L329
	}
L329:
	;
	if v1183 != 0 {
		v1172 = v1187
		v1173 = v1181
		v1174 = v1183
		goto L326
	} else {
		goto L330
	}
L330:
	;
	goto L327
L331:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193))))
	if base.B2i32(v1199 == int32(0))|base.B2i32(base.Ui32(v1195) < base.Ui32(int32(4))) != 0 {
		v1228 = v1193
		v1229 = v1194
		v1230 = v1195
		goto L318
	} else {
		goto L332
	}
L332:
	;
	v1206 = v1193
	v1207 = v1194
	v1208 = v1195
	goto L333
L333:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1206)))
	v1214 = int32(-2139062144)
	if (int32(16843008)-v1211|v1211)&v1214 != v1214 {
		v1235 = v1206
		v1236 = v1207
		v1237 = v1208
		goto L317
	} else {
		goto L335
	}
L334:
	;
	v1228 = v1222
	v1229 = v1220
	v1230 = v1224
	goto L318
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1207))) = v1211
	v1219 = int32(4)
	v1220 = v1207 + v1219
	v1222 = v1206 + v1219
	v1224 = v1208 - v1219
	if base.Ui32(int32(3)) < base.Ui32(v1224) {
		v1206 = v1222
		v1207 = v1220
		v1208 = v1224
		goto L333
	} else {
		goto L336
	}
L336:
	;
	goto L334
L337:
	;
	v1235 = v1228
	v1236 = v1229
	v1237 = v1230
	goto L317
L338:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1241))) = uint8(v1244)
	if v1244 == int32(0) {
		v1255 = v1240
		v1256 = v1241
		goto L316
	} else {
		goto L340
	}
L339:
	;
	v1255 = v1251
	v1256 = v1249
	goto L316
L340:
	;
	v1248 = int32(1)
	v1249 = v1241 + v1248
	v1251 = v1240 + v1248
	v1253 = v1242 - v1248
	if v1253 != 0 {
		v1240 = v1251
		v1241 = v1249
		v1242 = v1253
		goto L338
	} else {
		goto L341
	}
L341:
	;
	goto L339
L342:
	;
	v1281 = v19 + int32(336)
	goto L344
L343:
	;
	v1281 = v1273
	goto L344
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+236)) = v1281
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v1288 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1288)+32))
	v1290 = m.T0[v1289].(func(*base.Module, int32, int32) int32)(m, v1284, v19+int32(232))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1294 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	if v1290 == int32(0) {
		goto L296
	} else {
		goto L347
	}
L347:
	;
	if v1294 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v1298
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+164)) = uint32(v1092)
	v1302 = int64(base.Ui64(v1092) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+160)) = uint32(v1302)
	if v1099 != 0 {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	goto L350
L350:
	;
	v1320 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1321 = m.ExcPending
	if v1321 != 0 {
		goto L1
	} else {
		goto L359
	}
L351:
	;
	v1306 = int32(_a_F_WalReceiverMain_9)
	goto L353
L352:
	;
	v1306 = int32(_a_F_WalReceiverMain_10)
	goto L353
L353:
	;
	F_errmsg(m, v1306, v19+int32(160))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	if v1099 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1314 = int32(390)
	goto L357
L356:
	;
	v1314 = int32(394)
	goto L357
L357:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), v1314, int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	goto L350
L359:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41])) = v1320
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[42])) = v1320
	F_initStringInfo(m, int32(_a_F_WalReceiverMain_11))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
	;
	v1333 = m.G0
	v1334 = int32(16)
	v1335 = v1333 - v1334
	m.G0 = v1335
	F_gettimeofday(m, v1335)
	mBase = m.M
	v1338 = *(*int64)(unsafe.Add(mBase, uint32(v1335)))
	v1339 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1335)+8)))
	m.G0 = v1335 + v1334
	v1347 = v1339 + v1338*int64(1000000) - int64(946684800000000)
	goto L361
L361:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	v1355 = base.B2i32(v1349 <= int32(0))
	if v1349 <= int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1356 = int64(9223372036854775807)
	goto L364
L363:
	;
	v1356 = v1347 + base.I64_extend_i32_u(v1349)*int64(1000)
	goto L364
L364:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[44])) = v1356
	if v1349 <= int32(0) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1366 = int64(9223372036854775807)
	goto L367
L366:
	;
	v1366 = v1347 + base.I64_extend_i32_u(int32(base.Ui32(v1349)>>(uint(int32(1))%32)))*int64(1000)
	goto L367
L367:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[45])) = v1366
	v1371 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[46]))
	v1375 = v1347 + base.I64_extend_i32_u(v1371)*int64(1000000)
	if v1371 <= int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1378 = int64(9223372036854775807)
	goto L370
L369:
	;
	v1378 = v1375
	goto L370
L370:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[47])) = v1378
	if v1371 <= int32(0) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1384 = int64(9223372036854775807)
	goto L373
L372:
	;
	v1384 = v1375
	goto L373
L373:
	;
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[48])))
	if v1387 != 0 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1388 = v1384
	goto L376
L375:
	;
	v1388 = int64(9223372036854775807)
	goto L376
L376:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[49])) = v1388
	F_XLogWalRcvSendReply(m, int32(1), int32(0))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L1
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = int32(-1)
	v1401 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[50])))
	if v1401 == int32(1) {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L1
	} else {
		goto L573
	}
L380:
	;
	if v1411 != 0 {
		goto L384
	} else {
		goto L385
	}
L381:
	;
	v1406 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[51]))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+316))
	v1409 = base.B2i32(v1407 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[50])) = uint8(v1409)
	v1411 = v1409
	goto L383
L382:
	;
	v1411 = int32(0)
	goto L383
L383:
	;
	goto L380
L384:
	;
	goto L387
L385:
	;
	goto L386
L386:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L1
	} else {
		goto L569
	}
L387:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[52]))
	if v1429 != 0 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	goto L386
L389:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[53]))
	if v1433 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	goto L391
L393:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[53])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L1
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v1515 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+40))
	v1517 = m.T0[v1516].(func(*base.Module, int32, int32, int32) int32)(m, v1509, v19+int32(228), v19+int32(224))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L1
	} else {
		goto L414
	}
L396:
	;
	v1445 = m.G0
	v1446 = int32(16)
	v1447 = v1445 - v1446
	m.G0 = v1447
	F_gettimeofday(m, v1447)
	mBase = m.M
	v1450 = *(*int64)(unsafe.Add(mBase, uint32(v1447)))
	v1451 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1447)+8)))
	m.G0 = v1447 + v1446
	v1459 = v1451 + v1450*int64(1000000) - int64(946684800000000)
	goto L397
L397:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	v1467 = base.B2i32(v1461 <= int32(0))
	if v1461 <= int32(0) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1468 = int64(9223372036854775807)
	goto L400
L399:
	;
	v1468 = v1459 + base.I64_extend_i32_u(v1461)*int64(1000)
	goto L400
L400:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[44])) = v1468
	if v1461 <= int32(0) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1478 = int64(9223372036854775807)
	goto L403
L402:
	;
	v1478 = v1459 + base.I64_extend_i32_u(int32(base.Ui32(v1461)>>(uint(int32(1))%32)))*int64(1000)
	goto L403
L403:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[45])) = v1478
	v1483 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[46]))
	v1487 = v1459 + base.I64_extend_i32_u(v1483)*int64(1000000)
	if v1483 <= int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1490 = int64(9223372036854775807)
	goto L406
L405:
	;
	v1490 = v1487
	goto L406
L406:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[47])) = v1490
	if v1483 <= int32(0) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1496 = int64(9223372036854775807)
	goto L409
L408:
	;
	v1496 = v1487
	goto L409
L409:
	;
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[48])))
	if v1499 != 0 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1500 = v1496
	goto L412
L411:
	;
	v1500 = int64(9223372036854775807)
	goto L412
L412:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[49])) = v1500
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	goto L395
L414:
	;
	if v1517 != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	if int32(0) < v1517 {
		goto L419
	} else {
		goto L420
	}
L416:
	;
	goto L417
L417:
	;
	v2101 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[44]))
	v2103 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[45]))
	v2105 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[47]))
	v2107 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[49]))
	v2111 = m.G0
	v2112 = int32(16)
	v2113 = v2111 - v2112
	m.G0 = v2113
	F_gettimeofday(m, v2113)
	mBase = m.M
	v2116 = *(*int64)(unsafe.Add(mBase, uint32(v2113)))
	v2117 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2113)+8)))
	m.G0 = v2113 + v2112
	v2125 = v2117 + v2116*int64(1000000) - int64(946684800000000)
	goto L529
L418:
	;
	v2076 = int32(0)
	F_XLogWalRcvSendReply(m, v2076, v2076)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L527
	}
L419:
	;
	v1522 = v1517
	goto L422
L420:
	;
	goto L421
L421:
	;
	v2029 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L516
	}
L422:
	;
	v1540 = m.G0
	v1541 = int32(16)
	v1542 = v1540 - v1541
	m.G0 = v1542
	F_gettimeofday(m, v1542)
	mBase = m.M
	v1545 = *(*int64)(unsafe.Add(mBase, uint32(v1542)))
	v1546 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1542)+8)))
	m.G0 = v1542 + v1541
	v1554 = v1546 + v1545*int64(1000000) - int64(946684800000000)
	goto L424
L423:
	;
	if v2005 == int32(0) {
		goto L418
	} else {
		goto L515
	}
L424:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	if v1557 <= int32(0) {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[45])) = v1573
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[44])) = v1572
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v19)+228))
	v1579 = v1577 + int32(1)
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577))))
	v1582 = v1580 - int32(107)
	if v1582 != 0 {
		goto L435
	} else {
		goto L436
	}
L426:
	;
	v1560 = int64(9223372036854775807)
	v1572 = v1560
	v1573 = v1560
	goto L425
L427:
	;
	goto L428
L428:
	;
	v1563 = int64(1000)
	v1572 = base.I64_extend_i32_u(v1557)*v1563 + v1554
	v1573 = base.I64_extend_i32_u(int32(base.Ui32(v1557)>>(uint(int32(1))%32)))*v1563 + v1554
	goto L425
L429:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v2003 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v2003)+40))
	v2005 = m.T0[v2004].(func(*base.Module, int32, int32, int32) int32)(m, v1997, v19+int32(228), v19+int32(224))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L513
	}
L430:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v1966)+1464)) = v1951
	v1969 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	if v1969 < int32(0) {
		goto L429
	} else {
		goto L510
	}
L431:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L1
	} else {
		goto L506
	}
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L1
	} else {
		goto L502
	}
L433:
	;
	v1914 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41]))
	v1951 = v1914
	v1952 = v1595
	goto L430
L434:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L1
	} else {
		goto L498
	}
L435:
	;
	if v1582 != int32(12) {
		goto L431
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	if v1522 != int32(18) {
		goto L432
	} else {
		goto L491
	}
L438:
	;
	if base.Ui32(v1522) <= base.Ui32(int32(24)) {
		goto L434
	} else {
		goto L439
	}
L439:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1436)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+1428)) = int64(24)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1424)) = v1579
	v1594 = v19 + int32(1424)
	v1595 = F_pq_getmsgint64(m, v1594)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	v1597 = F_pq_getmsgint64(m, v1594)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v1599 = F_pq_getmsgint64(m, v1594)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	F_ProcessWalSndrMessage(m, v1597, v1599)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L1
	} else {
		goto L443
	}
L443:
	;
	v1604 = v1522 - int32(25)
	if v1604 == int32(0) {
		goto L433
	} else {
		goto L444
	}
L444:
	;
	v1611 = v1595
	v1617 = v1577 + int32(25)
	v1618 = v1604
	goto L445
L445:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[54]))
	v1628 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	if int32(0) <= v1628 {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	v1951 = v1870
	v1952 = v1870
	goto L430
L447:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[55])) = int32(0)
	v1663 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[56])))
	v1666 = m.G0
	v1668 = v1666 - int32(16)
	m.G0 = v1668
	if v1663 != 0 {
		goto L456
	} else {
		goto L457
	}
L448:
	;
	v1632 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[57]))
	v1634 = base.I64_div_u_s(v1611, base.I64_extend_i32_s(v1626))
	if v1632 == v1634 {
		v1657 = v1626
		goto L447
	} else {
		goto L451
	}
L449:
	;
	v1644 = v1626
	goto L450
L450:
	;
	v1647 = base.I64_div_u_s(v1611, base.I64_extend_i32_s(v1644))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[57])) = v1647
	v1649 = F_XLogFileInit(m, v1647, v1587)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L454
	}
L451:
	;
	F_XLogWalRcvClose(m, v1587)
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[54]))
	v1641 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	if int32(0) <= v1641 {
		v1657 = v1639
		goto L447
	} else {
		goto L453
	}
L453:
	;
	v1644 = v1639
	goto L450
L454:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[58])) = v1587
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40])) = v1649
	v1656 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[54]))
	v1657 = v1656
	goto L447
L455:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v1682))) = int32(167772240)
	v1686 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	v1690 = base.I32_wrap_i64(v1611) & (v1657 - int32(1))
	if base.Ui32(v1657) < base.Ui32(v1618+v1690) {
		goto L459
	} else {
		goto L460
	}
L456:
	;
	F___clock_gettime(m, int32(1), v1668)
	mBase = m.M
	v1672 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1668)+8)))
	v1673 = *(*int64)(unsafe.Add(mBase, uint32(v1668)))
	v1677 = v1672 + v1673*int64(1000000000)
	goto L458
L457:
	;
	v1677 = int64(0)
	goto L458
L458:
	;
	m.G0 = v1668 + int32(16)
	goto L455
L459:
	;
	v1694 = v1657 - v1690
	goto L461
L460:
	;
	v1694 = v1618
	goto L461
L461:
	;
	v1696 = F_pwrite(m, v1686, v1617, v1694, base.I64_extend_i32_s(v1690))
	mBase = m.M
	v1698 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[59]))
	*(*int32)(unsafe.Add(mBase, uint32(v1698))) = int32(0)
	v1704 = int32(1)
	v1705 = base.I64_extend_i32_s(v1696)
	v1709 = m.G0
	v1711 = v1709 - int32(16)
	m.G0 = v1711
	if v1677 != int64(0) {
		goto L463
	} else {
		goto L464
	}
L462:
	;
	if v1696 <= int32(0) {
		goto L479
	} else {
		goto L480
	}
L463:
	;
	F___clock_gettime(m, int32(1), v1711)
	mBase = m.M
	v1717 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1711)+8)))
	v1718 = *(*int64)(unsafe.Add(mBase, uint32(v1711)))
	v1722 = v1717 + (v1718*int64(1000000000) - v1677)
	goto L466
L464:
	;
	goto L465
L465:
	;
	v1809 = int32(888)
	v1810 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[60]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[60])) = v1810 + base.I64_extend_i32_u(v1704)
	v1814 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[61]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[61])) = v1814 + v1705
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v1704, v1705)
	mBase = m.M
	v1819 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[62])) = uint8(v1819)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[63])) = uint8(v1819)
	m.G0 = v1711 + int32(16)
	goto L462
L466:
	;
	v1772 = int32(888)
	v1773 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[64]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[64])) = v1773 + v1722
	v1777 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[0]))
	v1784 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1777))|base.B2i32(int32(1)<<(uint(v1777)%32)&int32(_a_F_WalReceiverMain_12) == v1784) == v1784 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v1789 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[65]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[65])) = v1789 + v1722
	v1793 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[62])) = uint8(v1793)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[66])) = uint8(v1793)
	goto L478
L477:
	;
	goto L478
L478:
	;
	goto L465
L479:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[55]))
	if v1830 == int32(0) {
		goto L482
	} else {
		goto L483
	}
L480:
	;
	goto L481
L481:
	;
	v1870 = v1611 + v1705
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41])) = v1870
	v1873 = v1618 - v1696
	if v1873 != 0 {
		v1611 = v1870
		v1617 = v1696 + v1617
		v1618 = v1873
		goto L445
	} else {
		goto L490
	}
L482:
	;
	v1834 = int32(51)
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[55])) = v1834
	v1837 = v1834
	goto L484
L483:
	;
	v1837 = v1830
	goto L484
L484:
	;
	v1839 = v19 + int32(1440)
	v1841 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[58]))
	v1843 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[57]))
	v1845 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[54]))
	F_XLogFileName(m, v1839, v1841, v1843, v1845)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[55])) = v1837
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v1694
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = v1690
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v1839
	F_errmsg(m, int32(_a_F_WalReceiverMain_13), v19+int32(144))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(953), int32(_a_F_WalReceiverMain_14))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L490:
	;
	goto L446
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1452)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+1444)) = int64(17)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+1440)) = v1579
	v1882 = v19 + int32(1440)
	v1883 = F_pq_getmsgint64(m, v1882)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	v1885 = F_pq_getmsgint64(m, v1882)
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	v1887 = F_pq_getmsgbyte(m, v1882)
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	F_ProcessWalSndrMessage(m, v1883, v1885)
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	if v1887 == int32(0) {
		goto L429
	} else {
		goto L496
	}
L496:
	;
	F_XLogWalRcvSendReply(m, int32(1), int32(0))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	goto L429
L498:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_15), int32(0))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(837), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L502:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_17), int32(0))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(861), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L506:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1580
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_18), v19+int32(32))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L1
	} else {
		goto L508
	}
L508:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(882), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L510:
	;
	v1973 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[57]))
	v1975 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[54])))
	v1976 = base.I64_div_u_s(v1952, v1975)
	if v1973 == v1976 {
		goto L429
	} else {
		goto L511
	}
L511:
	;
	F_XLogWalRcvClose(m, v1587)
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	goto L429
L513:
	;
	if int32(0) < v2005 {
		v1522 = v2005
		goto L422
	} else {
		goto L514
	}
L514:
	;
	goto L423
L515:
	;
	goto L421
L516:
	;
	if v2029 != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_19), int32(0))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L1
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	v2054 = int32(0)
	F_XLogWalRcvSendReply(m, v2054, v2054)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L1
	} else {
		goto L523
	}
L520:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v2035
	v2038 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41]))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+136)) = uint32(v2038)
	v2041 = int64(base.Ui64(v2038) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+132)) = uint32(v2041)
	F_errdetail(m, int32(_a_F_WalReceiverMain_20), v19+int32(128))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(475), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	goto L519
L523:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	F_XLogWalRcvFlush(m, int32(0), v2060)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v2068 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2068)+36))
	m.T0[v2069].(func(*base.Module, int32, int32))(m, v2064, v19+int32(328))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v19)+328))
	F_WalRcvFetchTimeLineHistoryFiles(m, v2072, v2073)
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	v2325 = v2054
	goto L295
L527:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	F_XLogWalRcvFlush(m, int32(0), v2081)
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	goto L417
L529:
	;
	if v2103 < v2101 {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v2127 = v2103
	goto L532
L531:
	;
	v2127 = v2101
	goto L532
L532:
	;
	if v2105 < v2127 {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v2129 = v2105
	goto L535
L534:
	;
	v2129 = v2127
	goto L535
L535:
	;
	if v2107 < v2129 {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	v2131 = v2107
	goto L538
L537:
	;
	v2131 = v2129
	goto L538
L538:
	;
	if v2131 <= v2125 {
		v2149 = int32(0)
		goto L540
	} else {
		goto L541
	}
L539:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[67]))
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v19)+224))
	v2154 = F_WaitLatchOrSocket(m, v2151, v2152, v2149, int32(83886094))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L544
	}
L540:
	;
	goto L539
L541:
	;
	v2137 = v2131 - v2125
	if base.B2i32(int64(0) < v2125)^base.B2i32(v2137 < v2131)|base.B2i32(int64(2147483646000) < v2137) != 0 {
		v2149 = int32(2147483647)
		goto L540
	} else {
		goto L542
	}
L542:
	;
	v2146 = base.I64_div_s(v2137+int64(999), int64(1000))
	v2149 = base.I32_wrap_i64(v2146)
	goto L540
L543:
	;
	if v2154&int32(8) != 0 {
		goto L553
	} else {
		goto L554
	}
L544:
	;
	if v2154&int32(1) == int32(0) {
		goto L543
	} else {
		goto L545
	}
L545:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[67]))
	*(*int32)(unsafe.Add(mBase, uint32(v2161))) = int32(0)
	goto L546
L546:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[52]))
	if v2165 != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v31)+1472))
	if v2168 == int32(0) {
		goto L543
	} else {
		goto L551
	}
L550:
	;
	goto L549
L551:
	;
	v2171 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+1472)) = v2171
	F_XLogWalRcvSendReply(m, int32(1), v2171)
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L1
	} else {
		goto L552
	}
L552:
	;
	goto L543
L553:
	;
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+224)) = int32(-1)
	v2222 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[50])))
	if v2222 == int32(1) {
		goto L565
	} else {
		goto L566
	}
L556:
	;
	v2185 = m.G0
	v2186 = int32(16)
	v2187 = v2185 - v2186
	m.G0 = v2187
	F_gettimeofday(m, v2187)
	mBase = m.M
	v2190 = *(*int64)(unsafe.Add(mBase, uint32(v2187)))
	v2191 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2187)+8)))
	m.G0 = v2187 + v2186
	v2199 = v2191 + v2190*int64(1000000) - int64(946684800000000)
	goto L557
L557:
	;
	v2201 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[44]))
	if v2201 <= v2199 {
		goto L379
	} else {
		goto L558
	}
L558:
	;
	v2204 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[45]))
	if v2204 <= v2199 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[45])) = int64(9223372036854775807)
	goto L561
L560:
	;
	goto L561
L561:
	;
	v2209 = base.B2i32(v2204 <= v2199)
	F_XLogWalRcvSendReply(m, v2209, v2209)
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L1
	} else {
		goto L562
	}
L562:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(0))
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	goto L555
L564:
	;
	if v2232 != 0 {
		goto L387
	} else {
		goto L568
	}
L565:
	;
	v2227 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[51]))
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+316))
	v2230 = base.B2i32(v2228 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[50])) = uint8(v2230)
	v2232 = v2230
	goto L567
L566:
	;
	v2232 = int32(0)
	goto L567
L567:
	;
	goto L564
L568:
	;
	goto L388
L569:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_21), int32(0))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(428), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L573:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L1
	} else {
		goto L574
	}
L574:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_22), int32(0))
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(573), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L577:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v19)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v2288
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v2290
	F_errmsg(m, int32(_a_F_WalReceiverMain_23), v19+int32(16))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(337), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L581:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v2304
	F_errmsg(m, int32(_a_F_WalReceiverMain_24), v19+int32(176))
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(606), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	v2325 = v1099
	goto L295
L584:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40])) = int32(-1)
	v2384 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L1
	} else {
		goto L594
	}
L585:
	;
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v19)+332))
	F_XLogWalRcvFlush(m, int32(0), v2337)
	mBase = m.M
	v2339 = m.ExcPending
	if v2339 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	v2341 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[58]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v2341
	v2344 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[57]))
	v2347 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[54])))
	v2348 = base.I64_div_u_s(int64(4294967296), v2347)
	v2349 = base.I64_div_u_s(v2344, v2348)
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+116)) = uint32(v2349)
	v2352 = v2344 - v2348*v2349
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+120)) = uint32(v2352)
	v2355 = v19 + int32(1440)
	v2360 = F_pg_snprintf(m, v2355, int32(64), int32(_a_F_WalReceiverMain_25), v19+int32(112))
	mBase = m.M
	v2361 = m.ExcPending
	if v2361 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	v2363 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[40]))
	v2364 = F_close(m, v2363)
	mBase = m.M
	if v2364 != 0 {
		goto L282
	} else {
		goto L588
	}
L588:
	;
	v2366 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[68]))
	if v2366 != int32(2) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	F_XLogArchiveForceDone(m, v2355)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L1
	} else {
		goto L592
	}
L590:
	;
	goto L591
L591:
	;
	F_XLogArchiveNotify(m, v19+int32(1440))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L1
	} else {
		goto L593
	}
L592:
	;
	goto L584
L593:
	;
	goto L584
L594:
	;
	if v2384 != 0 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_26), int32(0))
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L1
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+1456)) = int32(1)
	v2401 = v2396 + int32(1456)
	if v2397 != 0 {
		goto L600
	} else {
		goto L601
	}
L598:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(635), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	goto L597
L600:
	;
	F_s_lock(m, v2401, int32(_a_F_WalReceiverMain_0), int32(650), int32(_a_F_WalReceiverMain_27))
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L1
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+8))
	if v2407 != int32(2) {
		goto L604
	} else {
		goto L605
	}
L603:
	;
	goto L602
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2401))) = int32(0)
	if v2407 == int32(5) {
		goto L281
	} else {
		goto L607
	}
L605:
	;
	goto L606
L606:
	;
	v2427 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+1456)) = v2427
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+40)) = v2427
	*(*int64)(unsafe.Add(mBase, uint32(v2396)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+8)) = int32(3)
	F_WakeupRecovery(m)
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L1
	} else {
		goto L611
	}
L607:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_28), int32(0))
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(658), int32(_a_F_WalReceiverMain_27))
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L611:
	;
	goto L612
L612:
	;
	v2454 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[67]))
	*(*int32)(unsafe.Add(mBase, uint32(v2454))) = int32(0)
	goto L614
L613:
	;
	v2486 = *(*int64)(unsafe.Add(mBase, uint32(v2396)+32))
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+332)) = v2487
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+1456)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2396)+8)) = int32(2)
	v2494 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[69])))
	if v2494 != 0 {
		goto L628
	} else {
		goto L629
	}
L614:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[52]))
	if v2458 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2401)))
	*(*int32)(unsafe.Add(mBase, uint32(v2401))) = int32(1)
	if v2461 != 0 {
		goto L619
	} else {
		goto L620
	}
L618:
	;
	goto L617
L619:
	;
	F_s_lock(m, v2401, int32(_a_F_WalReceiverMain_0), int32(678), int32(_a_F_WalReceiverMain_27))
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L1
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2396)+8))
	switch v2469 - int32(4) {
	case 0:
		goto L623
	case 1:
		goto L625
	default:
		goto L624
	}
L622:
	;
	goto L621
L623:
	;
	goto L613
L624:
	;
	v2477 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2401))) = v2477
	v2480 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[67]))
	v2484 = F_WaitLatch(m, v2480, int32(33), v2477, int32(134217782))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L1
	} else {
		goto L627
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2401))) = int32(0)
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L627:
	;
	goto L612
L628:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+84)) = uint32(v2486)
	v2497 = int64(base.Ui64(v2486) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+80)) = uint32(v2497)
	v2500 = v19 + int32(1440)
	v2505 = F_pg_snprintf(m, v2500, int32(50), int32(_a_F_WalReceiverMain_29), v19+int32(80))
	mBase = m.M
	v2506 = m.ExcPending
	if v2506 != 0 {
		goto L1
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v2510 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	v2514 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2514)+16))
	v2516 = m.T0[v2515].(func(*base.Module, int32, int32) int32)(m, v2510, v19+int32(328))
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L1
	} else {
		goto L632
	}
L631:
	;
	v2507 = F_strlen(m, v2500)
	mBase = m.M
	goto L630
L632:
	;
	v2519 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[39]))
	v2520 = *(*int64)(unsafe.Add(mBase, uint32(v2519)))
	goto L633
L633:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v2520
	v2523 = v19 + int32(272)
	v2528 = F_pg_snprintf(m, v2523, int32(32), int32(_a_F_WalReceiverMain_7), v19-int32(-64))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	v2532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2516))))
	v2535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2523))))
	if base.B2i32(v2532 == int32(0))|base.B2i32(v2532 != v2535) != 0 {
		v2553 = v2532
		v2554 = v2535
		goto L636
	} else {
		goto L637
	}
L635:
	;
	if v2553-v2554 == int32(0) {
		v1092 = v2486
		v1099 = v2325
		goto L293
	} else {
		goto L642
	}
L636:
	;
	goto L635
L637:
	;
	v2538 = v2516
	v2539 = v2523
	goto L638
L638:
	;
	v2542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2539)+1)))
	v2543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2538)+1)))
	if v2543 == int32(0) {
		v2553 = v2543
		v2554 = v2542
		goto L636
	} else {
		goto L640
	}
L639:
	;
	v2553 = v2543
	v2554 = v2542
	goto L636
L640:
	;
	v2546 = int32(1)
	if v2543 == v2542 {
		v2538 = v2538 + v2546
		v2539 = v2539 + v2546
		goto L638
	} else {
		goto L641
	}
L641:
	;
	goto L639
L642:
	;
	goto L294
L643:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_30), int32(0))
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v2559
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v19 + int32(272)
	F_errdetail(m, int32(_a_F_WalReceiverMain_31), v19+int32(48))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(326), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L648:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v19 + int32(1440)
	F_errmsg(m, int32(_a_F_WalReceiverMain_32), v19+int32(96))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(622), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L653:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v747
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v19)+324))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v2629
	F_errmsg(m, int32(_a_F_WalReceiverMain_33), v19)
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(277), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2638 = m.ExcPending
	if v2638 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L657:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_34), int32(0))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_0), int32(265), int32(_a_F_WalReceiverMain_1))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int64
	_ = v393
	var v395 int64
	_ = v395
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
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
	v560 = int32(m.ExcTag)
	v561 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v560 == int32(0) {
		goto L153
	} else {
		goto L154
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
	v421 = v15
	v422 = v16
	goto L8
L8:
	;
	if v422 != 0 {
		goto L119
	} else {
		goto L120
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
	v75 = int32(916)
	v77 = m.G0
	v79 = v77 - int32(32)
	m.G0 = v79
	switch int32(918) {
	case 0, 2:
		v89 = v75
		goto L24
	default:
		goto L25
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
	v54 = v32 + int32(12)
	goto L18
L16:
	;
	m.G0 = v32 + int32(32)
	goto L10
L18:
	;
	goto L19
L19:
	;
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v60 = int32(20)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[2])) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[3])) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[4])) = v66
	goto L22
L21:
	;
	goto L22
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v122 = int32(916)
	v124 = m.G0
	v126 = v124 - int32(32)
	m.G0 = v126
	switch int32(918) {
	case 0, 2:
		v136 = v122
		goto L37
	default:
		goto L38
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v89
	F_sigemptyset(m, v79+int32(16))
	mBase = m.M
	goto L27
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[5])) = v75
	v89 = int32(_a_F_WalWriterMain_0)
	goto L24
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = int32(268435456)
	v101 = v79 + int32(12)
	goto L31
L29:
	;
	m.G0 = v79 + int32(32)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if v101 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v108 = int32(40)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[6])) = v109
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v101)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[7])) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[8])) = v113
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L29
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v169 = int32(-2)
	v171 = m.G0
	v173 = v171 - int32(32)
	m.G0 = v173
	switch int32(0) {
	case 0, 2:
		v183 = v169
		goto L50
	default:
		goto L51
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = v136
	F_sigemptyset(m, v126+int32(16))
	mBase = m.M
	goto L40
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[9])) = v122
	v136 = int32(_a_F_WalWriterMain_0)
	goto L37
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+24)) = int32(268435456)
	v148 = v126 + int32(12)
	goto L44
L42:
	;
	m.G0 = v126 + int32(32)
	goto L36
L44:
	;
	goto L45
L45:
	;
	if v148 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v155 = int32(300)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10])) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[11])) = v158
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[12])) = v160
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L42
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v216 = int32(-2)
	v218 = m.G0
	v220 = v218 - int32(32)
	m.G0 = v220
	switch int32(0) {
	case 0, 2:
		v230 = v216
		goto L63
	default:
		goto L64
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v183
	F_sigemptyset(m, v173+int32(16))
	mBase = m.M
	goto L53
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[13])) = v169
	v183 = int32(_a_F_WalWriterMain_0)
	goto L50
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = int32(268435456)
	v195 = v173 + int32(12)
	goto L57
L55:
	;
	m.G0 = v173 + int32(32)
	goto L49
L57:
	;
	goto L58
L58:
	;
	if v195 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v202 = int32(280)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v195)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[14])) = v203
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v195)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[15])) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v195)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[16])) = v207
	goto L61
L60:
	;
	goto L61
L61:
	;
	goto L55
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v263 = int32(917)
	v265 = m.G0
	v267 = v265 - int32(32)
	m.G0 = v267
	switch int32(919) {
	case 0, 2:
		v277 = v263
		goto L76
	default:
		goto L77
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = v230
	F_sigemptyset(m, v220+int32(16))
	mBase = m.M
	goto L66
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17])) = v216
	v230 = int32(_a_F_WalWriterMain_0)
	goto L63
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(268435456)
	v242 = v220 + int32(12)
	goto L70
L68:
	;
	m.G0 = v220 + int32(32)
	goto L62
L70:
	;
	goto L71
L71:
	;
	if v242 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v249 = int32(260)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[18])) = v250
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v242)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19])) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v242)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[20])) = v254
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L68
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v310 = int32(-2)
	v312 = m.G0
	v314 = v312 - int32(32)
	m.G0 = v314
	switch int32(0) {
	case 0, 2:
		v324 = v310
		goto L89
	default:
		goto L90
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+12)) = v277
	F_sigemptyset(m, v267+int32(16))
	mBase = m.M
	goto L79
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[21])) = v263
	v277 = int32(_a_F_WalWriterMain_0)
	goto L76
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+24)) = int32(268435456)
	v289 = v267 + int32(12)
	goto L83
L81:
	;
	m.G0 = v267 + int32(32)
	goto L75
L83:
	;
	goto L84
L84:
	;
	if v289 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v296 = int32(200)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[22])) = v297
	v299 = *(*int64)(unsafe.Add(mBase, uint32(v289)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[23])) = v299
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v289)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[24])) = v301
	goto L87
L86:
	;
	goto L87
L87:
	;
	goto L81
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v357 = int32(0)
	v359 = m.G0
	v361 = v359 - int32(32)
	m.G0 = v361
	switch int32(2) {
	case 0, 2:
		v371 = v357
		goto L102
	default:
		goto L103
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+12)) = v324
	F_sigemptyset(m, v314+int32(16))
	mBase = m.M
	goto L92
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[25])) = v310
	v324 = int32(_a_F_WalWriterMain_0)
	goto L89
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = int32(268435456)
	v336 = v314 + int32(12)
	goto L96
L94:
	;
	m.G0 = v314 + int32(32)
	goto L88
L96:
	;
	goto L97
L97:
	;
	if v336 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v343 = int32(240)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v336)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[26])) = v344
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v336)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[27])) = v346
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v336)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[28])) = v348
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L94
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v15
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[29]))
	v409 = F_AllocSetContextCreateInternal(m, v404, int32(_a_F_WalWriterMain_1), int32(0), int32(_a_F_WalWriterMain_2), int32(_a_F_WalWriterMain_3))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L5
	} else {
		goto L114
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+12)) = v371
	F_sigemptyset(m, v361+int32(16))
	mBase = m.M
	goto L104
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[30])) = v357
	v371 = int32(_a_F_WalWriterMain_0)
	goto L102
L104:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+24)) = int32(268435457)
	v383 = v361 + int32(12)
	goto L109
L107:
	;
	m.G0 = v361 + int32(32)
	goto L101
L109:
	;
	goto L110
L110:
	;
	if v383 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v390 = int32(340)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[31])) = v391
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v383)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[32])) = v393
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v383)))
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[33])) = v395
	goto L113
L112:
	;
	goto L113
L113:
	;
	goto L107
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[34])) = v409
	goto L115
L115:
	;
	v415 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v415)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v415))) = v9 + int32(12)
	goto L118
L116:
	;
	v421 = v409
	v422 = int32(0)
	goto L8
L118:
	;
	goto L116
L119:
	;
	v423 = int32(_a_F_WalWriterMain_4)
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[35]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[35])) = v425 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[36])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_EmitErrorReport(m)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L5
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[37])) = v9 + int32(16)
	F_sigprocmask(m, int32(_a_F_WalWriterMain_5), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L134
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_LWLockReleaseAll(m)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[38]))
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_UnlockBuffers(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_smgrdestroyall(m)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[34])) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_FlushErrorState(m)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_MemoryContextReset(m, v421)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	v475 = int32(_a_F_WalWriterMain_4)
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[35]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[35])) = v477 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_pg_usleep(m, int32(_a_F_WalWriterMain_6))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	goto L121
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_SetWalWriterSleeping(m, int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[39]))
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v500)+60)) = v502
	v506 = int32(0)
	v509 = int32(50)
	goto L136
L136:
	;
	v512 = base.B2i32(v509 < int32(2))
	if v512 != v506&int32(1) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_SetWalWriterSleeping(m, v512)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	v519 = v506
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[41]))
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = int32(0)
	goto L142
L141:
	;
	v519 = v512
	goto L140
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_ProcessMainLoopInterrupts(m)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	v529 = F_XLogBackgroundFlush(m)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v421
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[41]))
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[42]))
	if v529 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v547 = int32(50)
	goto L148
L147:
	;
	v547 = v509 - base.B2i32(int32(0) < v509)
	goto L148
L148:
	;
	if int32(0) < v547 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v550 = v540
	goto L151
L150:
	;
	v550 = v540 * int32(25)
	goto L151
L151:
	;
	v552 = F_WaitLatch(m, v537, int32(41), v550, int32(83886097))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v506 = v519
	v509 = v547
	goto L136
L153:
	;
	v565 = int32(v561)
	m.G0 = v9
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v568)))
	if v9+int32(12) == v571 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	m.ExcPending = 1
	goto L162
L155:
	;
	if v575 != 0 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	v575 = v573
	goto L158
L157:
	;
	v575 = int32(0)
	goto L158
L158:
	;
	goto L155
L159:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v12 = v575
	v15 = v576
	v16 = v567
	goto L1
L160:
	;
	goto L161
L161:
	;
	F___wasm_longjmp(m, v568, v567)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	return
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
