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
	var v36 int32
	_ = v36
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v22 = F_AllocateDir(m, int32(165628))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_FreeDir(m, v22)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L53
	}
L2:
	;
	return int32(0)
L3:
	;
	v27 = F_ReadDir(m, v22, int32(165628))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v232 = v4
	goto L1
L6:
	;
	goto L7
L7:
	;
	v36 = v17 + int32(-32)
	v51 = v27
	v53 = v4
	goto L8
L8:
	;
	v64 = v51 + int32(19)
	v65 = int32(528047)
	v69 = m.G0
	v71 = v69 - int32(32)
	v72 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v71)+24)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v71)+16)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v71)+8)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v71))) = v72
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _consts[315])))
	if v80 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v232 = v219
	goto L1
L10:
	;
	v224 = F_ReadDir(m, v22, int32(165628))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L51
	}
L11:
	;
	if v148 != int32(40) {
		v219 = v53
		goto L10
	} else {
		goto L32
	}
L12:
	;
	v148 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _consts[316])))
	if v84 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v88 = v64
	goto L18
L16:
	;
	goto L17
L17:
	;
	v98 = v65
	v99 = v80
	goto L21
L18:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v94 == v80 {
		v88 = v88 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v148 = v88 - v64
	goto L11
L20:
	;
	goto L19
L21:
	;
	v106 = v71 + int32(base.Ui32(v99)>>(uint(int32(3))%32))&int32(28)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v107 | v108<<(uint(v99)%32)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v112 != 0 {
		v98 = v98 + v108
		v99 = v112
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v115 == int32(0) {
		v140 = v64
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v148 = v140 - v64
	goto L11
L25:
	;
	v119 = v64
	v120 = v115
	goto L26
L26:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v71+int32(base.Ui32(v120)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v128)>>(uint(v120)%32))&int32(1) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v140 = v136
	goto L24
L28:
	;
	v140 = v119
	goto L24
L29:
	;
	goto L30
L30:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	v136 = v119 + int32(1)
	if v134 != 0 {
		v119 = v136
		v120 = v134
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v152 = v51 + int32(59)
	v153 = int32(17768)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _consts[317])))
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v157 == int32(0) {
		v176 = v156
		v177 = v157
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v177-v176 != 0 {
		v219 = v53
		goto L10
	} else {
		goto L41
	}
L34:
	;
	goto L33
L35:
	;
	if v156 != v157 {
		v176 = v156
		v177 = v157
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v161 = v152
	v162 = v153
	goto L37
L37:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	if v166 == int32(0) {
		v176 = v165
		v177 = v166
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v176 = v165
	v177 = v166
	goto L34
L39:
	;
	v169 = int32(1)
	if v165 == v166 {
		v161 = v161 + v169
		v162 = v162 + v169
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v17 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v36 | int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v36 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v36 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v17 + int32(-32)
	v187 = F_sscanf(m, v64, int32(501230), v19)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v189 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+44)))
	v190 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+48)))
	v191 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+36)))
	v192 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+40)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if v194 != l0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v196 = l0
	goto L45
L44:
	;
	v196 = int32(0)
	goto L45
L45:
	;
	if v196 != 0 {
		v219 = v53
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v199 = v191<<(uint(int64(32))%64) | v192
	if base.Ui64(l2-int64(1)) < base.Ui64(v199) {
		v219 = v53
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v205 = v189<<(uint(int64(32))%64) | v190
	if base.B2i32(l1 != int64(0))&base.B2i32(base.Ui64(v205) <= base.Ui64(l1)) != 0 {
		v219 = v53
		goto L10
	} else {
		goto L48
	}
L48:
	;
	v209 = F_palloc(m, int32(24))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v209)+8)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v209))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v209)+16)) = v194
	v214 = F_lappend(m, v53, v209)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v219 = v214
	goto L10
L51:
	;
	if v224 != 0 {
		v51 = v224
		v53 = v219
		goto L8
	} else {
		goto L52
	}
L52:
	;
	goto L9
L53:
	;
	m.G0 = v19 - int32(-64)
	return v232
}
func F_WalReceiverMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int64
	_ = v325
	var v326 int64
	_ = v326
	var v334 int64
	_ = v334
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v463 int32
	_ = v463
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v547 int32
	_ = v547
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v577 int32
	_ = v577
	var v589 int32
	_ = v589
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v619 int32
	_ = v619
	var v631 int32
	_ = v631
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v673 int32
	_ = v673
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int64
	_ = v1004
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1048 int64
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1219 int32
	_ = v1219
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1255 int64
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1276 int64
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int64
	_ = v1294
	var v1295 int64
	_ = v1295
	var v1303 int64
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1312 int64
	_ = v1312
	var v1322 int64
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1331 int64
	_ = v1331
	var v1334 int64
	_ = v1334
	var v1340 int64
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int64
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1407 int64
	_ = v1407
	var v1408 int64
	_ = v1408
	var v1416 int64
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1425 int64
	_ = v1425
	var v1435 int64
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1444 int64
	_ = v1444
	var v1447 int64
	_ = v1447
	var v1453 int64
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1459 int64
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int64
	_ = v1503
	var v1504 int64
	_ = v1504
	var v1512 int64
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1518 int64
	_ = v1518
	var v1521 int64
	_ = v1521
	var v1530 int64
	_ = v1530
	var v1531 int64
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1551 int64
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int64
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1559 int64
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1571 int64
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1591 int64
	_ = v1591
	var v1593 int64
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1606 int64
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1631 int64
	_ = v1631
	var v1632 int64
	_ = v1632
	var v1636 int64
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1663 int32
	_ = v1663
	var v1664 int64
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1676 int64
	_ = v1676
	var v1677 int64
	_ = v1677
	var v1681 int64
	_ = v1681
	var v1733 int32
	_ = v1733
	var v1734 int64
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1756 int32
	_ = v1756
	var v1757 int64
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1778 int32
	_ = v1778
	var v1779 int64
	_ = v1779
	var v1784 int32
	_ = v1784
	var v1785 int64
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1801 int32
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1814 int64
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1836 int32
	_ = v1836
	var v1841 int32
	_ = v1841
	var v1843 int64
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1856 int64
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1860 int64
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1891 int64
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1920 int32
	_ = v1920
	var v1925 int32
	_ = v1925
	var v1928 int64
	_ = v1928
	var v1929 int64
	_ = v1929
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1949 int64
	_ = v1949
	var v1951 int64
	_ = v1951
	var v1952 int64
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1972 int32
	_ = v1972
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2012 int64
	_ = v2012
	var v2015 int64
	_ = v2015
	var v2021 int32
	_ = v2021
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2074 int64
	_ = v2074
	var v2076 int64
	_ = v2076
	var v2078 int64
	_ = v2078
	var v2080 int64
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2089 int64
	_ = v2089
	var v2090 int64
	_ = v2090
	var v2098 int64
	_ = v2098
	var v2100 int64
	_ = v2100
	var v2102 int64
	_ = v2102
	var v2104 int64
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2110 int64
	_ = v2110
	var v2118 int64
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2148 int32
	_ = v2148
	var v2153 int32
	_ = v2153
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2162 int64
	_ = v2162
	var v2163 int64
	_ = v2163
	var v2171 int64
	_ = v2171
	var v2173 int64
	_ = v2173
	var v2176 int64
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2267 int32
	_ = v2267
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2281 int32
	_ = v2281
	var v2286 int32
	_ = v2286
	var v2296 int32
	_ = v2296
	var v2303 int32
	_ = v2303
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2314 int64
	_ = v2314
	var v2317 int64
	_ = v2317
	var v2318 int64
	_ = v2318
	var v2319 int64
	_ = v2319
	var v2322 int64
	_ = v2322
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2336 int32
	_ = v2336
	var v2342 int32
	_ = v2342
	var v2346 int32
	_ = v2346
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2372 int32
	_ = v2372
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2407 int32
	_ = v2407
	var v2424 int32
	_ = v2424
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int64
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2464 int32
	_ = v2464
	var v2467 int64
	_ = v2467
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2486 int32
	_ = v2486
	var v2491 int32
	_ = v2491
	var v2495 int32
	_ = v2495
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2521 int32
	_ = v2521
	var v2525 int32
	_ = v2525
	var v2537 int32
	_ = v2537
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2547 int64
	_ = v2547
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2574 int32
	_ = v2574
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2587 int32
	_ = v2587
	var v2604 int32
	_ = v2604
	var v2607 int32
	_ = v2607
	var v2611 int32
	_ = v2611
	var v2620 int32
	_ = v2620
	var v2625 int32
	_ = v2625
	var v2629 int32
	_ = v2629
	var v2631 int32
	_ = v2631
	var v2639 int32
	_ = v2639
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2678 int32
	_ = v2678
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(1504)
	m.G0 = v18
	*(*int32)(unsafe.Add(mBase, _consts[264])) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+320)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v18)+316)) = v7
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[856]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+1456)) = int32(1)
	v35 = v30 + int32(1456)
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_s_lock(m, v35, int32(486724), int32(188), int32(274082))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	switch v41 {
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
	v69 = *(*int32)(unsafe.Add(mBase, _consts[702]))
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+1453)) = uint8(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v69
	v76 = v18 + int32(400)
	v78 = v30 + int32(104)
	goto L19
L8:
	;
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v53
	F_errstart_cold(m, int32(23), v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+1456)) = int32(0)
	F_ConditionVariableBroadcast(m, v30+int32(12))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(0)
	goto L9
L11:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	F_errmsg_internal(m, int32(345915), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(486724), int32(213), int32(274082))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
	v195 = v18 + int32(336)
	v197 = v30 + int32(1388)
	goto L51
L17:
	;
	v191 = F_strlen(m, v180)
	mBase = m.M
	goto L16
L19:
	;
	goto L20
L20:
	;
	v85 = int32(1023)
	if (v76^v78)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v184)
	goto L17
L22:
	;
	v165 = v160
	v166 = v161
	v167 = v162
	goto L44
L23:
	;
	if v155 == int32(0) {
		v180 = v153
		v181 = v154
		goto L21
	} else {
		goto L43
	}
L24:
	;
	v153 = v78
	v154 = v76
	v155 = v85
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v78&int32(3) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v122 == int32(0) {
		v180 = v119
		v181 = v120
		goto L21
	} else {
		goto L36
	}
L28:
	;
	v119 = v78
	v120 = v76
	v121 = v85
	v122 = int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v98 = v78
	v99 = v76
	v100 = v85
	goto L31
L31:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v102)
	if v102 == int32(0) {
		v160 = v98
		v161 = v99
		v162 = v100
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v119 = v113
	v120 = v107
	v121 = v109
	v122 = v111
	goto L27
L33:
	;
	v106 = int32(1)
	v107 = v99 + v106
	v109 = v100 - v106
	v110 = int32(0)
	v111 = base.B2i32(v109 != v110)
	v113 = v98 + v106
	if v113&int32(3) == v110 {
		v119 = v113
		v120 = v107
		v121 = v109
		v122 = v111
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v109 != 0 {
		v98 = v113
		v99 = v107
		v100 = v109
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v125 == int32(0) {
		v153 = v119
		v154 = v120
		v155 = v121
		goto L23
	} else {
		goto L37
	}
L37:
	;
	if base.Ui32(v121) < base.Ui32(int32(4)) {
		v153 = v119
		v154 = v120
		v155 = v121
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v131 = v119
	v132 = v120
	v133 = v121
	goto L39
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v139 = int32(-2139062144)
	if (int32(16843008)-v136|v136)&v139 != v139 {
		v160 = v131
		v161 = v132
		v162 = v133
		goto L22
	} else {
		goto L41
	}
L40:
	;
	v153 = v147
	v154 = v145
	v155 = v149
	goto L23
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v136
	v144 = int32(4)
	v145 = v132 + v144
	v147 = v131 + v144
	v149 = v133 - v144
	if base.Ui32(int32(3)) < base.Ui32(v149) {
		v131 = v147
		v132 = v145
		v133 = v149
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v160 = v153
	v161 = v154
	v162 = v155
	goto L22
L44:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v169)
	if v169 == int32(0) {
		v180 = v165
		v181 = v166
		goto L21
	} else {
		goto L46
	}
L45:
	;
	v180 = v176
	v181 = v174
	goto L21
L46:
	;
	v173 = int32(1)
	v174 = v166 + v173
	v176 = v165 + v173
	v178 = v167 - v173
	if v178 != 0 {
		v165 = v176
		v166 = v174
		v167 = v178
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1452)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+332)) = v315
	v320 = m.G0
	v321 = int32(16)
	v322 = v320 - v321
	m.G0 = v322
	F___gettimeofday(m, v322)
	mBase = m.M
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
	v326 = int64(*(*int32)(unsafe.Add(mBase, uint32(v322)+8)))
	m.G0 = v322 + v321
	v334 = v326 + v325*int64(1000000) - int64(946684800000000)
	goto L80
L49:
	;
	v310 = F_strlen(m, v299)
	mBase = m.M
	goto L48
L51:
	;
	goto L52
L52:
	;
	v204 = int32(63)
	if (v195^v197)&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v300))) = uint8(v303)
	goto L49
L54:
	;
	v284 = v279
	v285 = v280
	v286 = v281
	goto L76
L55:
	;
	if v274 == int32(0) {
		v299 = v272
		v300 = v273
		goto L53
	} else {
		goto L75
	}
L56:
	;
	v272 = v197
	v273 = v195
	v274 = v204
	goto L55
L57:
	;
	goto L58
L58:
	;
	if v197&int32(3) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v241 == int32(0) {
		v299 = v238
		v300 = v239
		goto L53
	} else {
		goto L68
	}
L60:
	;
	v238 = v197
	v239 = v195
	v240 = v204
	v241 = int32(1)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v217 = v197
	v218 = v195
	v219 = v204
	goto L63
L63:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v221)
	if v221 == int32(0) {
		v279 = v217
		v280 = v218
		v281 = v219
		goto L54
	} else {
		goto L65
	}
L64:
	;
	v238 = v232
	v239 = v226
	v240 = v228
	v241 = v230
	goto L59
L65:
	;
	v225 = int32(1)
	v226 = v218 + v225
	v228 = v219 - v225
	v229 = int32(0)
	v230 = base.B2i32(v228 != v229)
	v232 = v217 + v225
	if v232&int32(3) == v229 {
		v238 = v232
		v239 = v226
		v240 = v228
		v241 = v230
		goto L59
	} else {
		goto L66
	}
L66:
	;
	if v228 != 0 {
		v217 = v232
		v218 = v226
		v219 = v228
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v244 == int32(0) {
		v272 = v238
		v273 = v239
		v274 = v240
		goto L55
	} else {
		goto L69
	}
L69:
	;
	if base.Ui32(v240) < base.Ui32(int32(4)) {
		v272 = v238
		v273 = v239
		v274 = v240
		goto L55
	} else {
		goto L70
	}
L70:
	;
	v250 = v238
	v251 = v239
	v252 = v240
	goto L71
L71:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v258 = int32(-2139062144)
	if (int32(16843008)-v255|v255)&v258 != v258 {
		v279 = v250
		v280 = v251
		v281 = v252
		goto L54
	} else {
		goto L73
	}
L72:
	;
	v272 = v266
	v273 = v264
	v274 = v268
	goto L55
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v255
	v263 = int32(4)
	v264 = v251 + v263
	v266 = v250 + v263
	v268 = v252 - v263
	if base.Ui32(int32(3)) < base.Ui32(v268) {
		v250 = v266
		v251 = v264
		v252 = v268
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v279 = v272
	v280 = v273
	v281 = v274
	goto L54
L76:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v288)
	if v288 == int32(0) {
		v299 = v284
		v300 = v285
		goto L53
	} else {
		goto L78
	}
L77:
	;
	v299 = v295
	v300 = v293
	goto L53
L78:
	;
	v292 = int32(1)
	v293 = v285 + v292
	v295 = v284 + v292
	v297 = v286 - v292
	if v297 != 0 {
		v284 = v295
		v285 = v293
		v286 = v297
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+80)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v30)+96)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v30)+72)) = v334
	v339 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+1456)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v339
	v344 = *(*int32)(unsafe.Add(mBase, _consts[856]))
	*(*int64)(unsafe.Add(mBase, uint32(v344)+1464)) = int64(0)
	F_on_shmem_exit(m, int32(1028), v18+int32(332))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v353 = int32(913)
	v355 = m.G0
	v357 = v355 - int32(144)
	m.G0 = v357
	switch int32(915) {
	case 0, 2:
		v367 = v353
		goto L83
	default:
		goto L84
	}
L82:
	;
	v395 = int32(-2)
	v397 = m.G0
	v399 = v397 - int32(144)
	m.G0 = v399
	switch int32(0) {
	case 0, 2:
		v409 = v395
		goto L96
	default:
		goto L97
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+4)) = v367
	F_sigemptyset(m, v357+int32(8))
	mBase = m.M
	goto L86
L84:
	;
	*(*int32)(unsafe.Add(mBase, _consts[321])) = v353
	v367 = int32(4729)
	goto L83
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+136)) = int32(268435456)
	v379 = v357 + int32(4)
	goto L90
L88:
	;
	m.G0 = v357 + int32(144)
	goto L82
L90:
	;
	goto L91
L91:
	;
	if v379 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v390 = F___memcpy(m, int32(4635612), v379, int32(140))
	mBase = m.M
	goto L94
L93:
	;
	goto L94
L94:
	;
	goto L88
L95:
	;
	v437 = int32(295)
	v439 = m.G0
	v441 = v439 - int32(144)
	m.G0 = v441
	switch int32(297) {
	case 0, 2:
		v451 = v437
		goto L109
	default:
		goto L110
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v399)+4)) = v409
	F_sigemptyset(m, v399+int32(8))
	mBase = m.M
	goto L99
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v395
	v409 = int32(4729)
	goto L96
L99:
	;
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v399)+136)) = int32(268435456)
	v421 = v399 + int32(4)
	goto L103
L101:
	;
	m.G0 = v399 + int32(144)
	goto L95
L103:
	;
	goto L104
L104:
	;
	if v421 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v432 = F___memcpy(m, int32(4635752), v421, int32(140))
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L101
L108:
	;
	v479 = int32(-2)
	v481 = m.G0
	v483 = v481 - int32(144)
	m.G0 = v483
	switch int32(0) {
	case 0, 2:
		v493 = v479
		goto L122
	default:
		goto L123
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441)+4)) = v451
	F_sigemptyset(m, v441+int32(8))
	mBase = m.M
	goto L112
L110:
	;
	*(*int32)(unsafe.Add(mBase, _consts[323])) = v437
	v451 = int32(4729)
	goto L109
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v441)+136)) = int32(268435456)
	v463 = v441 + int32(4)
	goto L116
L114:
	;
	m.G0 = v441 + int32(144)
	goto L108
L116:
	;
	goto L117
L117:
	;
	if v463 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v474 = F___memcpy(m, int32(4637572), v463, int32(140))
	mBase = m.M
	goto L120
L119:
	;
	goto L120
L120:
	;
	goto L114
L121:
	;
	v521 = int32(-2)
	v523 = m.G0
	v525 = v523 - int32(144)
	m.G0 = v525
	switch int32(0) {
	case 0, 2:
		v535 = v521
		goto L135
	default:
		goto L136
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+4)) = v493
	F_sigemptyset(m, v483+int32(8))
	mBase = m.M
	goto L125
L123:
	;
	*(*int32)(unsafe.Add(mBase, _consts[724])) = v479
	v493 = int32(4729)
	goto L122
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+136)) = int32(268435456)
	v505 = v483 + int32(4)
	goto L129
L127:
	;
	m.G0 = v483 + int32(144)
	goto L121
L129:
	;
	goto L130
L130:
	;
	if v505 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v516 = F___memcpy(m, int32(4637432), v505, int32(140))
	mBase = m.M
	goto L133
L132:
	;
	goto L133
L133:
	;
	goto L127
L134:
	;
	v563 = int32(916)
	v565 = m.G0
	v567 = v565 - int32(144)
	m.G0 = v567
	switch int32(918) {
	case 0, 2:
		v577 = v563
		goto L148
	default:
		goto L149
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v525)+4)) = v535
	F_sigemptyset(m, v525+int32(8))
	mBase = m.M
	goto L138
L136:
	;
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v521
	v535 = int32(4729)
	goto L135
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v525)+136)) = int32(268435456)
	v547 = v525 + int32(4)
	goto L142
L140:
	;
	m.G0 = v525 + int32(144)
	goto L134
L142:
	;
	goto L143
L143:
	;
	if v547 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v558 = F___memcpy(m, int32(4637292), v547, int32(140))
	mBase = m.M
	goto L146
L145:
	;
	goto L146
L146:
	;
	goto L140
L147:
	;
	v605 = int32(-2)
	v607 = m.G0
	v609 = v607 - int32(144)
	m.G0 = v609
	switch int32(0) {
	case 0, 2:
		v619 = v605
		goto L161
	default:
		goto L162
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+4)) = v577
	F_sigemptyset(m, v567+int32(8))
	mBase = m.M
	goto L151
L149:
	;
	*(*int32)(unsafe.Add(mBase, _consts[647])) = v563
	v577 = int32(4729)
	goto L148
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+136)) = int32(268435456)
	v589 = v567 + int32(4)
	goto L155
L153:
	;
	m.G0 = v567 + int32(144)
	goto L147
L155:
	;
	goto L156
L156:
	;
	if v589 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v600 = F___memcpy(m, int32(4636872), v589, int32(140))
	mBase = m.M
	goto L159
L158:
	;
	goto L159
L159:
	;
	goto L153
L160:
	;
	v647 = int32(0)
	v649 = m.G0
	v651 = v649 - int32(144)
	m.G0 = v651
	switch int32(2) {
	case 0, 2:
		v661 = v647
		goto L174
	default:
		goto L175
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v609)+4)) = v619
	F_sigemptyset(m, v609+int32(8))
	mBase = m.M
	goto L164
L162:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v605
	v619 = int32(4729)
	goto L161
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v609)+136)) = int32(268435456)
	v631 = v609 + int32(4)
	goto L168
L166:
	;
	m.G0 = v609 + int32(144)
	goto L160
L168:
	;
	goto L169
L169:
	;
	if v631 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v642 = F___memcpy(m, int32(4637152), v631, int32(140))
	mBase = m.M
	goto L172
L171:
	;
	goto L172
L172:
	;
	goto L166
L173:
	;
	F_load_file(m, int32(211807), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L186
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651)+4)) = v661
	F_sigemptyset(m, v651+int32(8))
	mBase = m.M
	goto L176
L175:
	;
	*(*int32)(unsafe.Add(mBase, _consts[650])) = v647
	v661 = int32(4729)
	goto L174
L176:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v651)+136)) = int32(268435457)
	v673 = v651 + int32(4)
	goto L181
L179:
	;
	m.G0 = v651 + int32(144)
	goto L173
L181:
	;
	goto L182
L182:
	;
	if v673 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v684 = F___memcpy(m, int32(4637852), v673, int32(140))
	mBase = m.M
	goto L185
L184:
	;
	goto L185
L185:
	;
	goto L179
L186:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	if v693 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	F_sigprocmask(m, int32(4377784), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2669 = m.ExcPending
	if v2669 != 0 {
		goto L1
	} else {
		goto L682
	}
L190:
	;
	v702 = int32(0)
	v705 = *(*int32)(unsafe.Add(mBase, _consts[857]))
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	if v707 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v708 = v705
	goto L193
L192:
	;
	v708 = int32(211812)
	goto L193
L193:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	v714 = m.T0[v713].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v18+int32(400), int32(1), v702, v702, v708, v18+int32(324))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, _consts[858])) = v714
	if v714 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v718)+8))
	v720 = m.T0[v719].(func(*base.Module, int32) int32)(m, v714)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L1
	} else {
		goto L678
	}
L198:
	;
	v723 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v729 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	m.T0[v730].(func(*base.Module, int32, int32, int32))(m, v723, v18+int32(320), v18+int32(316))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(1)
	if v733 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	F_s_lock(m, v35, int32(486724), int32(286), int32(274082))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v744 = F__emscripten_memset_bulkmem(m, v78, base.I32_extend8_s(int32(0)), int32(1024))
	mBase = m.M
	goto L204
L203:
	;
	goto L202
L204:
	;
	if v720 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	goto L211
L206:
	;
	goto L207
L207:
	;
	v865 = F__emscripten_memset_bulkmem(m, v30+int32(1128), base.I32_extend8_s(int32(0)), int32(255))
	mBase = m.M
	goto L240
L208:
	;
	goto L207
L209:
	;
	v857 = F_strlen(m, v846)
	mBase = m.M
	goto L208
L211:
	;
	goto L212
L212:
	;
	v751 = int32(1023)
	if (v744^v720)&int32(3) != 0 {
		goto L216
	} else {
		goto L217
	}
L213:
	;
	v850 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v847))) = uint8(v850)
	goto L209
L214:
	;
	v831 = v826
	v832 = v827
	v833 = v828
	goto L236
L215:
	;
	if v821 == int32(0) {
		v846 = v819
		v847 = v820
		goto L213
	} else {
		goto L235
	}
L216:
	;
	v819 = v720
	v820 = v744
	v821 = v751
	goto L215
L217:
	;
	goto L218
L218:
	;
	if v720&int32(3) == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	if v788 == int32(0) {
		v846 = v785
		v847 = v786
		goto L213
	} else {
		goto L228
	}
L220:
	;
	v785 = v720
	v786 = v744
	v787 = v751
	v788 = int32(1)
	goto L219
L221:
	;
	goto L222
L222:
	;
	v764 = v720
	v765 = v744
	v766 = v751
	goto L223
L223:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764))))
	*(*uint8)(unsafe.Add(mBase, uint32(v765))) = uint8(v768)
	if v768 == int32(0) {
		v826 = v764
		v827 = v765
		v828 = v766
		goto L214
	} else {
		goto L225
	}
L224:
	;
	v785 = v779
	v786 = v773
	v787 = v775
	v788 = v777
	goto L219
L225:
	;
	v772 = int32(1)
	v773 = v765 + v772
	v775 = v766 - v772
	v776 = int32(0)
	v777 = base.B2i32(v775 != v776)
	v779 = v764 + v772
	if v779&int32(3) == v776 {
		v785 = v779
		v786 = v773
		v787 = v775
		v788 = v777
		goto L219
	} else {
		goto L226
	}
L226:
	;
	if v775 != 0 {
		v764 = v779
		v765 = v773
		v766 = v775
		goto L223
	} else {
		goto L227
	}
L227:
	;
	goto L224
L228:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v785))))
	if v791 == int32(0) {
		v819 = v785
		v820 = v786
		v821 = v787
		goto L215
	} else {
		goto L229
	}
L229:
	;
	if base.Ui32(v787) < base.Ui32(int32(4)) {
		v819 = v785
		v820 = v786
		v821 = v787
		goto L215
	} else {
		goto L230
	}
L230:
	;
	v797 = v785
	v798 = v786
	v799 = v787
	goto L231
L231:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v797)))
	v805 = int32(-2139062144)
	if (int32(16843008)-v802|v802)&v805 != v805 {
		v826 = v797
		v827 = v798
		v828 = v799
		goto L214
	} else {
		goto L233
	}
L232:
	;
	v819 = v813
	v820 = v811
	v821 = v815
	goto L215
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v798))) = v802
	v810 = int32(4)
	v811 = v798 + v810
	v813 = v797 + v810
	v815 = v799 - v810
	if base.Ui32(int32(3)) < base.Ui32(v815) {
		v797 = v813
		v798 = v811
		v799 = v815
		goto L231
	} else {
		goto L234
	}
L234:
	;
	goto L232
L235:
	;
	v826 = v819
	v827 = v820
	v828 = v821
	goto L214
L236:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831))))
	*(*uint8)(unsafe.Add(mBase, uint32(v832))) = uint8(v835)
	if v835 == int32(0) {
		v846 = v831
		v847 = v832
		goto L213
	} else {
		goto L238
	}
L237:
	;
	v846 = v842
	v847 = v840
	goto L213
L238:
	;
	v839 = int32(1)
	v840 = v832 + v839
	v842 = v831 + v839
	v844 = v833 - v839
	if v844 != 0 {
		v831 = v842
		v832 = v840
		v833 = v844
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v18)+320))
	if v866 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	goto L247
L242:
	;
	goto L243
L243:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+1456)) = int32(0)
	v985 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+1453)) = uint8(v985)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+1384)) = v982
	if v720 != 0 {
		goto L276
	} else {
		goto L277
	}
L244:
	;
	goto L243
L245:
	;
	v979 = F_strlen(m, v968)
	mBase = m.M
	goto L244
L247:
	;
	goto L248
L248:
	;
	v873 = int32(254)
	if (v865^v866)&int32(3) != 0 {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	v972 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v969))) = uint8(v972)
	goto L245
L250:
	;
	v953 = v948
	v954 = v949
	v955 = v950
	goto L272
L251:
	;
	if v943 == int32(0) {
		v968 = v941
		v969 = v942
		goto L249
	} else {
		goto L271
	}
L252:
	;
	v941 = v866
	v942 = v865
	v943 = v873
	goto L251
L253:
	;
	goto L254
L254:
	;
	if v866&int32(3) == int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	if v910 == int32(0) {
		v968 = v907
		v969 = v908
		goto L249
	} else {
		goto L264
	}
L256:
	;
	v907 = v866
	v908 = v865
	v909 = v873
	v910 = int32(1)
	goto L255
L257:
	;
	goto L258
L258:
	;
	v886 = v866
	v887 = v865
	v888 = v873
	goto L259
L259:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	*(*uint8)(unsafe.Add(mBase, uint32(v887))) = uint8(v890)
	if v890 == int32(0) {
		v948 = v886
		v949 = v887
		v950 = v888
		goto L250
	} else {
		goto L261
	}
L260:
	;
	v907 = v901
	v908 = v895
	v909 = v897
	v910 = v899
	goto L255
L261:
	;
	v894 = int32(1)
	v895 = v887 + v894
	v897 = v888 - v894
	v898 = int32(0)
	v899 = base.B2i32(v897 != v898)
	v901 = v886 + v894
	if v901&int32(3) == v898 {
		v907 = v901
		v908 = v895
		v909 = v897
		v910 = v899
		goto L255
	} else {
		goto L262
	}
L262:
	;
	if v897 != 0 {
		v886 = v901
		v887 = v895
		v888 = v897
		goto L259
	} else {
		goto L263
	}
L263:
	;
	goto L260
L264:
	;
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907))))
	if v913 == int32(0) {
		v941 = v907
		v942 = v908
		v943 = v909
		goto L251
	} else {
		goto L265
	}
L265:
	;
	if base.Ui32(v909) < base.Ui32(int32(4)) {
		v941 = v907
		v942 = v908
		v943 = v909
		goto L251
	} else {
		goto L266
	}
L266:
	;
	v919 = v907
	v920 = v908
	v921 = v909
	goto L267
L267:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v919)))
	v927 = int32(-2139062144)
	if (int32(16843008)-v924|v924)&v927 != v927 {
		v948 = v919
		v949 = v920
		v950 = v921
		goto L250
	} else {
		goto L269
	}
L268:
	;
	v941 = v935
	v942 = v933
	v943 = v937
	goto L251
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920))) = v924
	v932 = int32(4)
	v933 = v920 + v932
	v935 = v919 + v932
	v937 = v921 - v932
	if base.Ui32(int32(3)) < base.Ui32(v937) {
		v919 = v935
		v920 = v933
		v921 = v937
		goto L267
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	v948 = v941
	v949 = v942
	v950 = v943
	goto L250
L272:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v953))))
	*(*uint8)(unsafe.Add(mBase, uint32(v954))) = uint8(v957)
	if v957 == int32(0) {
		v968 = v953
		v969 = v954
		goto L249
	} else {
		goto L274
	}
L273:
	;
	v968 = v964
	v969 = v962
	goto L249
L274:
	;
	v961 = int32(1)
	v962 = v954 + v961
	v964 = v953 + v961
	v966 = v955 - v961
	if v966 != 0 {
		v953 = v964
		v954 = v962
		v955 = v966
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	F_pfree(m, v720)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v18)+320))
	if v990 != 0 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L278
L280:
	;
	F_pfree(m, v990)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v994 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v998 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v998)+16))
	v1000 = m.T0[v999].(func(*base.Module, int32, int32) int32)(m, v994, v18+int32(328))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L284
	}
L283:
	;
	goto L282
L284:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v1004 = *(*int64)(unsafe.Add(mBase, uint32(v1003)))
	goto L285
L285:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+208)) = v1004
	v1012 = F_pg_snprintf(m, v18+int32(272), int32(32), int32(37711), v18+int32(208))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v1015 = v18 + int32(272)
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015))))
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000))))
	if v1019 == int32(0) {
		v1038 = v1018
		v1039 = v1019
		goto L290
	} else {
		goto L291
	}
L287:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L1
	} else {
		goto L677
	}
L288:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L1
	} else {
		goto L673
	}
L289:
	;
	if v1039-v1038 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L290:
	;
	goto L289
L291:
	;
	if v1018 != v1019 {
		v1038 = v1018
		v1039 = v1019
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1023 = v1000
	v1024 = v1015
	goto L293
L293:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024)+1)))
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1023)+1)))
	if v1028 == int32(0) {
		v1038 = v1027
		v1039 = v1028
		goto L290
	} else {
		goto L295
	}
L294:
	;
	v1038 = v1027
	v1039 = v1028
	goto L290
L295:
	;
	v1031 = int32(1)
	if v1027 == v1028 {
		v1023 = v1023 + v1031
		v1024 = v1024 + v1031
		goto L293
	} else {
		goto L296
	}
L296:
	;
	goto L294
L297:
	;
	v1043 = int32(1)
	v1048 = v313
	v1055 = v1043
	goto L300
L298:
	;
	v2587 = v1000
	goto L299
L299:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L1
	} else {
		goto L668
	}
L300:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v18)+328))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	if base.Ui32(v1062) <= base.Ui32(v1061) {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	v2587 = v2543
	goto L299
L302:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	if v2303 < int32(0) {
		goto L591
	} else {
		goto L592
	}
L303:
	;
	if v1247 == int32(0) {
		v2296 = v1055
		goto L302
	} else {
		goto L588
	}
L304:
	;
	F_WalRcvFetchTimeLineHistoryFiles(m, v1062, v1061)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L1
	} else {
		goto L584
	}
L307:
	;
	if v314&v1043 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v1069 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+56))
	v1071 = m.T0[v1070].(func(*base.Module, int32) int32)(m, v1067)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+240)) = v1048
	v1226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+232)) = uint8(v1226)
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+248)) = v1228
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+336)))
	if v1233 != 0 {
		goto L350
	} else {
		goto L351
	}
L311:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+192)) = base.I64_extend_i32_s(v1071)
	v1081 = F_pg_snprintf(m, v18+int32(336), int32(64), int32(422768), v18+int32(192))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v1088 = int32(0)
	v1093 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+48))
	v1095 = m.T0[v1094].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1084, v18+int32(336), int32(1), v1088, v1088, v1088, v1088)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(1)
	if v1097 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	F_s_lock(m, v35, int32(486724), int32(364), int32(274082))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v1106 = v18 + int32(336)
	goto L321
L317:
	;
	goto L316
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(0)
	goto L310
L319:
	;
	v1219 = F_strlen(m, v1208)
	mBase = m.M
	goto L318
L321:
	;
	goto L322
L322:
	;
	v1113 = int32(63)
	if (v197^v1106)&int32(3) != 0 {
		goto L326
	} else {
		goto L327
	}
L323:
	;
	v1212 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1209))) = uint8(v1212)
	goto L319
L324:
	;
	v1193 = v1188
	v1194 = v1189
	v1195 = v1190
	goto L346
L325:
	;
	if v1183 == int32(0) {
		v1208 = v1181
		v1209 = v1182
		goto L323
	} else {
		goto L345
	}
L326:
	;
	v1181 = v1106
	v1182 = v197
	v1183 = v1113
	goto L325
L327:
	;
	goto L328
L328:
	;
	if v1106&int32(3) == int32(0) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	if v1150 == int32(0) {
		v1208 = v1147
		v1209 = v1148
		goto L323
	} else {
		goto L338
	}
L330:
	;
	v1147 = v1106
	v1148 = v197
	v1149 = v1113
	v1150 = int32(1)
	goto L329
L331:
	;
	goto L332
L332:
	;
	v1126 = v1106
	v1127 = v197
	v1128 = v1113
	goto L333
L333:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1127))) = uint8(v1130)
	if v1130 == int32(0) {
		v1188 = v1126
		v1189 = v1127
		v1190 = v1128
		goto L324
	} else {
		goto L335
	}
L334:
	;
	v1147 = v1141
	v1148 = v1135
	v1149 = v1137
	v1150 = v1139
	goto L329
L335:
	;
	v1134 = int32(1)
	v1135 = v1127 + v1134
	v1137 = v1128 - v1134
	v1138 = int32(0)
	v1139 = base.B2i32(v1137 != v1138)
	v1141 = v1126 + v1134
	if v1141&int32(3) == v1138 {
		v1147 = v1141
		v1148 = v1135
		v1149 = v1137
		v1150 = v1139
		goto L329
	} else {
		goto L336
	}
L336:
	;
	if v1137 != 0 {
		v1126 = v1141
		v1127 = v1135
		v1128 = v1137
		goto L333
	} else {
		goto L337
	}
L337:
	;
	goto L334
L338:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147))))
	if v1153 == int32(0) {
		v1181 = v1147
		v1182 = v1148
		v1183 = v1149
		goto L325
	} else {
		goto L339
	}
L339:
	;
	if base.Ui32(v1149) < base.Ui32(int32(4)) {
		v1181 = v1147
		v1182 = v1148
		v1183 = v1149
		goto L325
	} else {
		goto L340
	}
L340:
	;
	v1159 = v1147
	v1160 = v1148
	v1161 = v1149
	goto L341
L341:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1159)))
	v1167 = int32(-2139062144)
	if (int32(16843008)-v1164|v1164)&v1167 != v1167 {
		v1188 = v1159
		v1189 = v1160
		v1190 = v1161
		goto L324
	} else {
		goto L343
	}
L342:
	;
	v1181 = v1175
	v1182 = v1173
	v1183 = v1177
	goto L325
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1160))) = v1164
	v1172 = int32(4)
	v1173 = v1160 + v1172
	v1175 = v1159 + v1172
	v1177 = v1161 - v1172
	if base.Ui32(int32(3)) < base.Ui32(v1177) {
		v1159 = v1175
		v1160 = v1173
		v1161 = v1177
		goto L341
	} else {
		goto L344
	}
L344:
	;
	goto L342
L345:
	;
	v1188 = v1181
	v1189 = v1182
	v1190 = v1183
	goto L324
L346:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1193))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1194))) = uint8(v1197)
	if v1197 == int32(0) {
		v1208 = v1193
		v1209 = v1194
		goto L323
	} else {
		goto L348
	}
L347:
	;
	v1208 = v1204
	v1209 = v1202
	goto L323
L348:
	;
	v1201 = int32(1)
	v1202 = v1194 + v1201
	v1204 = v1193 + v1201
	v1206 = v1195 - v1201
	if v1206 != 0 {
		v1193 = v1204
		v1194 = v1202
		v1195 = v1206
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	v1234 = v18 + int32(336)
	goto L352
L351:
	;
	v1234 = v1226
	goto L352
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+236)) = v1234
	v1237 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v1241 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+32))
	v1243 = m.T0[v1242].(func(*base.Module, int32, int32) int32)(m, v1237, v18+int32(232))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v1247 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	if v1243 == int32(0) {
		goto L303
	} else {
		goto L355
	}
L355:
	;
	if v1247 != 0 {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+168)) = v1251
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+164)) = uint32(v1048)
	v1255 = int64(base.Ui64(v1048) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+160)) = uint32(v1255)
	v1260 = v1055 & int32(1)
	if v1260 != 0 {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	goto L358
L358:
	;
	v1276 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L367
	}
L359:
	;
	v1261 = int32(50383)
	goto L361
L360:
	;
	v1261 = int32(50442)
	goto L361
L361:
	;
	F_errmsg(m, v1261, v18+int32(160))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	if v1260 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1269 = int32(390)
	goto L365
L364:
	;
	v1269 = int32(394)
	goto L365
L365:
	;
	F_errfinish(m, int32(486724), v1269, int32(274082))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	goto L358
L367:
	;
	*(*int64)(unsafe.Add(mBase, _consts[860])) = v1276
	*(*int64)(unsafe.Add(mBase, _consts[861])) = v1276
	F_initStringInfo(m, int32(4380864))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	v1289 = m.G0
	v1290 = int32(16)
	v1291 = v1289 - v1290
	m.G0 = v1291
	F___gettimeofday(m, v1291)
	mBase = m.M
	v1294 = *(*int64)(unsafe.Add(mBase, uint32(v1291)))
	v1295 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1291)+8)))
	m.G0 = v1291 + v1290
	v1303 = v1295 + v1294*int64(1000000) - int64(946684800000000)
	goto L369
L369:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, _consts[862]))
	v1311 = base.B2i32(v1305 <= int32(0))
	if v1305 <= int32(0) {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1312 = int64(9223372036854775807)
	goto L372
L371:
	;
	v1312 = v1303 + base.I64_extend_i32_u(v1305)*int64(1000)
	goto L372
L372:
	;
	*(*int64)(unsafe.Add(mBase, _consts[863])) = v1312
	if v1305 <= int32(0) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1322 = int64(9223372036854775807)
	goto L375
L374:
	;
	v1322 = v1303 + base.I64_extend_i32_u(int32(base.Ui32(v1305)>>(uint(int32(1))%32)))*int64(1000)
	goto L375
L375:
	;
	*(*int64)(unsafe.Add(mBase, _consts[864])) = v1322
	v1327 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	v1331 = v1303 + base.I64_extend_i32_u(v1327)*int64(1000000)
	if v1327 <= int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1334 = int64(9223372036854775807)
	goto L378
L377:
	;
	v1334 = v1331
	goto L378
L378:
	;
	*(*int64)(unsafe.Add(mBase, _consts[865])) = v1334
	if v1327 <= int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1340 = int64(9223372036854775807)
	goto L381
L380:
	;
	v1340 = v1331
	goto L381
L381:
	;
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, _consts[866])))
	if v1343&int32(1) != 0 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1346 = v1340
	goto L384
L383:
	;
	v1346 = int64(9223372036854775807)
	goto L384
L384:
	;
	*(*int64)(unsafe.Add(mBase, _consts[867])) = v1346
	F_XLogWalRcvSendReply(m, int32(1), int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+224)) = int32(-1)
	v1359 = int32(*(*uint8)(unsafe.Add(mBase, _consts[189])))
	if v1359 == int32(1) {
		goto L389
	} else {
		goto L390
	}
L387:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L1
	} else {
		goto L580
	}
L388:
	;
	if v1369 != 0 {
		goto L392
	} else {
		goto L393
	}
L389:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+316))
	v1367 = base.B2i32(v1365 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[189])) = uint8(v1367)
	v1369 = v1367
	goto L391
L390:
	;
	v1369 = int32(0)
	goto L391
L391:
	;
	goto L388
L392:
	;
	goto L395
L393:
	;
	goto L394
L394:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L1
	} else {
		goto L576
	}
L395:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v1386 != 0 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	goto L394
L397:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L1
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, _consts[705]))
	if v1390 != 0 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	goto L399
L401:
	;
	*(*int32)(unsafe.Add(mBase, _consts[705])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L1
	} else {
		goto L404
	}
L402:
	;
	goto L403
L403:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v1474 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1474)+40))
	v1476 = m.T0[v1475].(func(*base.Module, int32, int32, int32) int32)(m, v1468, v18+int32(228), v18+int32(224))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L1
	} else {
		goto L422
	}
L404:
	;
	v1402 = m.G0
	v1403 = int32(16)
	v1404 = v1402 - v1403
	m.G0 = v1404
	F___gettimeofday(m, v1404)
	mBase = m.M
	v1407 = *(*int64)(unsafe.Add(mBase, uint32(v1404)))
	v1408 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1404)+8)))
	m.G0 = v1404 + v1403
	v1416 = v1408 + v1407*int64(1000000) - int64(946684800000000)
	goto L405
L405:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, _consts[862]))
	v1424 = base.B2i32(v1418 <= int32(0))
	if v1418 <= int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1425 = int64(9223372036854775807)
	goto L408
L407:
	;
	v1425 = v1416 + base.I64_extend_i32_u(v1418)*int64(1000)
	goto L408
L408:
	;
	*(*int64)(unsafe.Add(mBase, _consts[863])) = v1425
	if v1418 <= int32(0) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1435 = int64(9223372036854775807)
	goto L411
L410:
	;
	v1435 = v1416 + base.I64_extend_i32_u(int32(base.Ui32(v1418)>>(uint(int32(1))%32)))*int64(1000)
	goto L411
L411:
	;
	*(*int64)(unsafe.Add(mBase, _consts[864])) = v1435
	v1440 = *(*int32)(unsafe.Add(mBase, _consts[847]))
	v1444 = v1416 + base.I64_extend_i32_u(v1440)*int64(1000000)
	if v1440 <= int32(0) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1447 = int64(9223372036854775807)
	goto L414
L413:
	;
	v1447 = v1444
	goto L414
L414:
	;
	*(*int64)(unsafe.Add(mBase, _consts[865])) = v1447
	if v1440 <= int32(0) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1453 = int64(9223372036854775807)
	goto L417
L416:
	;
	v1453 = v1444
	goto L417
L417:
	;
	v1456 = int32(*(*uint8)(unsafe.Add(mBase, _consts[866])))
	if v1456&int32(1) != 0 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1459 = v1453
	goto L420
L419:
	;
	v1459 = int64(9223372036854775807)
	goto L420
L420:
	;
	*(*int64)(unsafe.Add(mBase, _consts[867])) = v1459
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1463 = m.ExcPending
	if v1463 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	goto L403
L422:
	;
	if v1476 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	if int32(0) < v1476 {
		goto L427
	} else {
		goto L428
	}
L424:
	;
	goto L425
L425:
	;
	v2074 = *(*int64)(unsafe.Add(mBase, _consts[863]))
	v2076 = *(*int64)(unsafe.Add(mBase, _consts[864]))
	v2078 = *(*int64)(unsafe.Add(mBase, _consts[865]))
	v2080 = *(*int64)(unsafe.Add(mBase, _consts[867]))
	v2084 = m.G0
	v2085 = int32(16)
	v2086 = v2084 - v2085
	m.G0 = v2086
	F___gettimeofday(m, v2086)
	mBase = m.M
	v2089 = *(*int64)(unsafe.Add(mBase, uint32(v2086)))
	v2090 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2086)+8)))
	m.G0 = v2086 + v2085
	v2098 = v2090 + v2089*int64(1000000) - int64(946684800000000)
	goto L535
L426:
	;
	v2050 = int32(0)
	F_XLogWalRcvSendReply(m, v2050, v2050)
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L533
	}
L427:
	;
	v1481 = v1476
	goto L430
L428:
	;
	goto L429
L429:
	;
	v2003 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L1
	} else {
		goto L522
	}
L430:
	;
	v1498 = m.G0
	v1499 = int32(16)
	v1500 = v1498 - v1499
	m.G0 = v1500
	F___gettimeofday(m, v1500)
	mBase = m.M
	v1503 = *(*int64)(unsafe.Add(mBase, uint32(v1500)))
	v1504 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1500)+8)))
	m.G0 = v1500 + v1499
	v1512 = v1504 + v1503*int64(1000000) - int64(946684800000000)
	goto L432
L431:
	;
	if v1980 == int32(0) {
		goto L426
	} else {
		goto L521
	}
L432:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, _consts[862]))
	if v1515 <= int32(0) {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	*(*int64)(unsafe.Add(mBase, _consts[864])) = v1531
	*(*int64)(unsafe.Add(mBase, _consts[863])) = v1530
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v18)+228))
	v1537 = v1535 + int32(1)
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535))))
	switch v1538 - int32(107) {
	case 0:
		goto L443
	default:
		goto L439
	case 12:
		goto L444
	}
L434:
	;
	v1518 = int64(9223372036854775807)
	v1530 = v1518
	v1531 = v1518
	goto L433
L435:
	;
	goto L436
L436:
	;
	v1521 = int64(1000)
	v1530 = base.I64_extend_i32_u(v1515)*v1521 + v1512
	v1531 = base.I64_extend_i32_u(int32(base.Ui32(v1515)>>(uint(int32(1))%32)))*v1521 + v1512
	goto L433
L437:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v1978 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1978)+40))
	v1980 = m.T0[v1979].(func(*base.Module, int32, int32, int32) int32)(m, v1972, v18+int32(228), v18+int32(224))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L1
	} else {
		goto L519
	}
L438:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, _consts[856]))
	*(*int64)(unsafe.Add(mBase, uint32(v1942)+1464)) = v1928
	v1945 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	if v1945 < int32(0) {
		goto L437
	} else {
		goto L516
	}
L439:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L1
	} else {
		goto L512
	}
L440:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L1
	} else {
		goto L508
	}
L441:
	;
	v1891 = *(*int64)(unsafe.Add(mBase, _consts[860]))
	v1928 = v1891
	v1929 = v1551
	goto L438
L442:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L1
	} else {
		goto L504
	}
L443:
	;
	if v1481 != int32(18) {
		goto L440
	} else {
		goto L497
	}
L444:
	;
	if base.Ui32(v1481) <= base.Ui32(int32(24)) {
		goto L442
	} else {
		goto L445
	}
L445:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1436)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1428)) = int64(24)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1424)) = v1537
	v1551 = F_pq_getmsgint64(m, v18+int32(1424))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L1
	} else {
		goto L446
	}
L446:
	;
	v1555 = F_pq_getmsgint64(m, v18+int32(1424))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	v1559 = F_pq_getmsgint64(m, v18+int32(1424))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	F_ProcessWalSndrMessage(m, v1555, v1559)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	v1564 = v1481 - int32(25)
	if v1564 == int32(0) {
		goto L441
	} else {
		goto L450
	}
L450:
	;
	v1571 = v1551
	v1575 = v1564
	v1577 = v1535 + int32(25)
	goto L451
L451:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v1587 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	if int32(0) <= v1587 {
		goto L454
	} else {
		goto L455
	}
L452:
	;
	v1928 = v1843
	v1929 = v1843
	goto L438
L453:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v1622 = int32(*(*uint8)(unsafe.Add(mBase, _consts[868])))
	v1625 = m.G0
	v1627 = v1625 - int32(16)
	m.G0 = v1627
	if v1622 != 0 {
		goto L462
	} else {
		goto L463
	}
L454:
	;
	v1591 = *(*int64)(unsafe.Add(mBase, _consts[869]))
	v1593 = base.I64_div_u_s(v1571, base.I64_extend_i32_s(v1585))
	if v1591 == v1593 {
		v1616 = v1585
		goto L453
	} else {
		goto L457
	}
L455:
	;
	v1603 = v1585
	goto L456
L456:
	;
	v1606 = base.I64_div_u_s(v1571, base.I64_extend_i32_s(v1603))
	*(*int64)(unsafe.Add(mBase, _consts[869])) = v1606
	v1608 = F_XLogFileInit(m, v1606, v1543)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L1
	} else {
		goto L460
	}
L457:
	;
	F_XLogWalRcvClose(m, v1543)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v1600 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	if int32(0) <= v1600 {
		v1616 = v1598
		goto L453
	} else {
		goto L459
	}
L459:
	;
	v1603 = v1598
	goto L456
L460:
	;
	*(*int32)(unsafe.Add(mBase, _consts[870])) = v1543
	*(*int32)(unsafe.Add(mBase, _consts[859])) = v1608
	v1615 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v1616 = v1615
	goto L453
L461:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v1641))) = int32(167772240)
	v1645 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	v1649 = base.I32_wrap_i64(v1571) & (v1616 - int32(1))
	if base.Ui32(v1616) < base.Ui32(v1575+v1649) {
		goto L465
	} else {
		goto L466
	}
L462:
	;
	F___clock_gettime(m, int32(1), v1627)
	mBase = m.M
	v1631 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1627)+8)))
	v1632 = *(*int64)(unsafe.Add(mBase, uint32(v1627)))
	v1636 = v1631 + v1632*int64(1000000000)
	goto L464
L463:
	;
	v1636 = int64(0)
	goto L464
L464:
	;
	m.G0 = v1627 + int32(16)
	goto L461
L465:
	;
	v1653 = v1616 - v1649
	goto L467
L466:
	;
	v1653 = v1575
	goto L467
L467:
	;
	v1655 = F_pwrite(m, v1645, v1577, v1653, base.I64_extend_i32_s(v1649))
	mBase = m.M
	v1657 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v1657))) = int32(0)
	v1663 = int32(1)
	v1664 = base.I64_extend_i32_s(v1655)
	v1668 = m.G0
	v1670 = v1668 - int32(16)
	m.G0 = v1670
	if v1636 != int64(0) {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	if v1655 <= int32(0) {
		goto L485
	} else {
		goto L486
	}
L469:
	;
	F___clock_gettime(m, int32(1), v1670)
	mBase = m.M
	v1676 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1670)+8)))
	v1677 = *(*int64)(unsafe.Add(mBase, uint32(v1670)))
	v1681 = v1676 + (v1677*int64(1000000000) - v1636)
	goto L472
L470:
	;
	goto L471
L471:
	;
	v1778 = int32(4452192)
	v1779 = *(*int64)(unsafe.Add(mBase, _consts[871]))
	*(*int64)(unsafe.Add(mBase, _consts[871])) = v1779 + base.I64_extend_i32_u(v1663)
	v1784 = int32(4451232)
	v1785 = *(*int64)(unsafe.Add(mBase, _consts[872]))
	*(*int64)(unsafe.Add(mBase, _consts[872])) = v1785 + v1664
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v1663, v1664)
	mBase = m.M
	v1790 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v1790)
	*(*uint8)(unsafe.Add(mBase, _consts[873])) = uint8(v1790)
	m.G0 = v1670 + int32(16)
	goto L468
L472:
	;
	v1733 = int32(4453152)
	v1734 = *(*int64)(unsafe.Add(mBase, _consts[874]))
	*(*int64)(unsafe.Add(mBase, _consts[874])) = v1734 + v1681
	v1738 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	if base.Ui32(int32(16)) < base.Ui32(v1738) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	goto L471
L483:
	;
	if int32(1)<<(uint(v1738)%32)&int32(115186) == int32(0) {
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v1756 = int32(4450048)
	v1757 = *(*int64)(unsafe.Add(mBase, _consts[875]))
	*(*int64)(unsafe.Add(mBase, _consts[875])) = v1757 + v1681
	v1761 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v1761)
	*(*uint8)(unsafe.Add(mBase, _consts[876])) = uint8(v1761)
	goto L482
L485:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v1801 == int32(0) {
		goto L488
	} else {
		goto L489
	}
L486:
	;
	goto L487
L487:
	;
	v1843 = v1571 + v1664
	*(*int64)(unsafe.Add(mBase, _consts[860])) = v1843
	v1846 = v1575 - v1655
	if v1846 != 0 {
		v1571 = v1843
		v1575 = v1846
		v1577 = v1655 + v1577
		goto L451
	} else {
		goto L496
	}
L488:
	;
	v1805 = int32(51)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v1805
	v1808 = v1805
	goto L490
L489:
	;
	v1808 = v1801
	goto L490
L490:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	v1814 = *(*int64)(unsafe.Add(mBase, _consts[869]))
	v1816 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	F_XLogFileName(m, v18+int32(1440), v1812, v1814, v1816)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v1808
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L1
	} else {
		goto L493
	}
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+152)) = v1653
	*(*int32)(unsafe.Add(mBase, uint32(v18)+148)) = v1649
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v18 + int32(1440)
	F_errmsg(m, int32(287431), v18+int32(144))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	F_errfinish(m, int32(486724), int32(953), int32(343957))
	mBase = m.M
	v1841 = m.ExcPending
	if v1841 != 0 {
		goto L1
	} else {
		goto L495
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	goto L452
L497:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1452)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+1444)) = int64(17)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+1440)) = v1537
	v1856 = F_pq_getmsgint64(m, v18+int32(1440))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v1860 = F_pq_getmsgint64(m, v18+int32(1440))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v1864 = F_pq_getmsgbyte(m, v18+int32(1440))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F_ProcessWalSndrMessage(m, v1856, v1860)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	if v1864 == int32(0) {
		goto L437
	} else {
		goto L502
	}
L502:
	;
	F_XLogWalRcvSendReply(m, int32(1), int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	goto L437
L504:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	F_errmsg_internal(m, int32(17856), int32(0))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	F_errfinish(m, int32(486724), int32(837), int32(321420))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L507
	}
L507:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L508:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1898 = m.ExcPending
	if v1898 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	F_errmsg_internal(m, int32(17808), int32(0))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(486724), int32(861), int32(321420))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L512:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L1
	} else {
		goto L513
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v1538
	F_errmsg_internal(m, int32(468449), v18+int32(32))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(486724), int32(882), int32(321420))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	v1949 = *(*int64)(unsafe.Add(mBase, _consts[869]))
	v1951 = int64(*(*int32)(unsafe.Add(mBase, _consts[262])))
	v1952 = base.I64_div_u_s(v1929, v1951)
	if v1949 == v1952 {
		goto L437
	} else {
		goto L517
	}
L517:
	;
	F_XLogWalRcvClose(m, v1543)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	goto L437
L519:
	;
	if int32(0) < v1980 {
		v1481 = v1980
		goto L430
	} else {
		goto L520
	}
L520:
	;
	goto L431
L521:
	;
	goto L429
L522:
	;
	if v2003 != 0 {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	F_errmsg(m, int32(210921), int32(0))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L1
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	v2028 = int32(0)
	F_XLogWalRcvSendReply(m, v2028, v2028)
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L529
	}
L526:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v2009
	v2012 = *(*int64)(unsafe.Add(mBase, _consts[860]))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+136)) = uint32(v2012)
	v2015 = int64(base.Ui64(v2012) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+132)) = uint32(v2015)
	F_errdetail(m, int32(631063), v18+int32(128))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L1
	} else {
		goto L527
	}
L527:
	;
	F_errfinish(m, int32(486724), int32(475), int32(274082))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	goto L525
L529:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	F_XLogWalRcvFlush(m, int32(0), v2034)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v2042 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2042)+36))
	m.T0[v2043].(func(*base.Module, int32, int32))(m, v2038, v18+int32(328))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v18)+328))
	F_WalRcvFetchTimeLineHistoryFiles(m, v2046, v2047)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v2296 = v2028
	goto L302
L533:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	F_XLogWalRcvFlush(m, int32(0), v2055)
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L534
	}
L534:
	;
	goto L425
L535:
	;
	if v2076 < v2074 {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	v2100 = v2076
	goto L538
L537:
	;
	v2100 = v2074
	goto L538
L538:
	;
	if v2078 < v2100 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v2102 = v2078
	goto L541
L540:
	;
	v2102 = v2100
	goto L541
L541:
	;
	if v2080 < v2102 {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v2104 = v2080
	goto L544
L543:
	;
	v2104 = v2102
	goto L544
L544:
	;
	if v2104 <= v2098 {
		v2121 = int32(0)
		goto L546
	} else {
		goto L547
	}
L545:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v18)+224))
	v2126 = F_WaitLatchOrSocket(m, v2123, v2124, v2121, int32(83886094))
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L1
	} else {
		goto L551
	}
L546:
	;
	goto L545
L547:
	;
	v2107 = int32(2147483647)
	v2110 = v2104 - v2098
	if base.B2i32(int64(0) < v2098)^base.B2i32(v2110 < v2104) != 0 {
		v2121 = v2107
		goto L546
	} else {
		goto L548
	}
L548:
	;
	if int64(2147483646000) < v2110 {
		v2121 = v2107
		goto L546
	} else {
		goto L549
	}
L549:
	;
	v2118 = base.I64_div_s(v2110+int64(999), int64(1000))
	v2121 = base.I32_wrap_i64(v2118)
	goto L546
L550:
	;
	if v2126&int32(8) != 0 {
		goto L560
	} else {
		goto L561
	}
L551:
	;
	if v2126&int32(1) == int32(0) {
		goto L550
	} else {
		goto L552
	}
L552:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v2133))) = int32(0)
	goto L553
L553:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v2137 != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L1
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v30)+1472))
	if v2140 == int32(0) {
		goto L550
	} else {
		goto L558
	}
L557:
	;
	goto L556
L558:
	;
	v2143 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+1472)) = v2143
	F_XLogWalRcvSendReply(m, int32(1), v2143)
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L1
	} else {
		goto L559
	}
L559:
	;
	goto L550
L560:
	;
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v2153 = m.ExcPending
	if v2153 != 0 {
		goto L1
	} else {
		goto L563
	}
L561:
	;
	goto L562
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+224)) = int32(-1)
	v2194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[189])))
	if v2194 == int32(1) {
		goto L572
	} else {
		goto L573
	}
L563:
	;
	v2157 = m.G0
	v2158 = int32(16)
	v2159 = v2157 - v2158
	m.G0 = v2159
	F___gettimeofday(m, v2159)
	mBase = m.M
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v2159)))
	v2163 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2159)+8)))
	m.G0 = v2159 + v2158
	v2171 = v2163 + v2162*int64(1000000) - int64(946684800000000)
	goto L564
L564:
	;
	v2173 = *(*int64)(unsafe.Add(mBase, _consts[863]))
	if v2173 <= v2171 {
		goto L387
	} else {
		goto L565
	}
L565:
	;
	v2176 = *(*int64)(unsafe.Add(mBase, _consts[864]))
	if v2176 <= v2171 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	*(*int64)(unsafe.Add(mBase, _consts[864])) = int64(9223372036854775807)
	goto L568
L567:
	;
	goto L568
L568:
	;
	v2181 = base.B2i32(v2176 <= v2171)
	F_XLogWalRcvSendReply(m, v2181, v2181)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(0))
	mBase = m.M
	v2186 = m.ExcPending
	if v2186 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	goto L562
L571:
	;
	if v2204 != 0 {
		goto L395
	} else {
		goto L575
	}
L572:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v2199)+316))
	v2202 = base.B2i32(v2200 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[189])) = uint8(v2202)
	v2204 = v2202
	goto L574
L573:
	;
	v2204 = int32(0)
	goto L574
L574:
	;
	goto L571
L575:
	;
	goto L396
L576:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	F_errmsg(m, int32(454211), int32(0))
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L1
	} else {
		goto L578
	}
L578:
	;
	F_errfinish(m, int32(486724), int32(428), int32(274082))
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L1
	} else {
		goto L579
	}
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	F_errmsg(m, int32(64941), int32(0))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L1
	} else {
		goto L582
	}
L582:
	;
	F_errfinish(m, int32(486724), int32(573), int32(274082))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L584:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L1
	} else {
		goto L585
	}
L585:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v18)+328))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v2259
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v2261
	F_errmsg(m, int32(50034), v18+int32(16))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	F_errfinish(m, int32(486724), int32(337), int32(274082))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L1
	} else {
		goto L587
	}
L587:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L588:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v2275
	F_errmsg(m, int32(50976), v18+int32(176))
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(486724), int32(606), int32(274082))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	v2296 = v1055
	goto L302
L591:
	;
	*(*int32)(unsafe.Add(mBase, _consts[859])) = int32(-1)
	v2355 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L601
	}
L592:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v18)+332))
	F_XLogWalRcvFlush(m, int32(0), v2307)
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	v2311 = *(*int32)(unsafe.Add(mBase, _consts[870]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v2311
	v2314 = *(*int64)(unsafe.Add(mBase, _consts[869]))
	v2317 = int64(*(*int32)(unsafe.Add(mBase, _consts[262])))
	v2318 = base.I64_div_u_s(int64(4294967296), v2317)
	v2319 = base.I64_div_u_s(v2314, v2318)
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+116)) = uint32(v2319)
	v2322 = v2314 - v2318*v2319
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+120)) = uint32(v2322)
	v2330 = F_pg_snprintf(m, v18+int32(1440), int32(64), int32(501258), v18+int32(112))
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	v2333 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	v2334 = F_close(m, v2333)
	mBase = m.M
	if v2334 != 0 {
		goto L288
	} else {
		goto L595
	}
L595:
	;
	v2336 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	if v2336 != int32(2) {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	F_XLogArchiveForceDone(m, v18+int32(1440))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L1
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	F_XLogArchiveNotify(m, v18+int32(1440))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L600
	}
L599:
	;
	goto L591
L600:
	;
	goto L591
L601:
	;
	if v2355 != 0 {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	F_errmsg_internal(m, int32(138180), int32(0))
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L1
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, _consts[856]))
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2367)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+1456)) = int32(1)
	v2372 = v2367 + int32(1456)
	if v2368 != 0 {
		goto L607
	} else {
		goto L608
	}
L605:
	;
	F_errfinish(m, int32(486724), int32(635), int32(274082))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	goto L604
L607:
	;
	F_s_lock(m, v2372, int32(486724), int32(650), int32(245742))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L1
	} else {
		goto L610
	}
L608:
	;
	goto L609
L609:
	;
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v2367)+8))
	if v2378 != int32(2) {
		goto L611
	} else {
		goto L612
	}
L610:
	;
	goto L609
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2372))) = int32(0)
	if v2378 == int32(5) {
		goto L287
	} else {
		goto L614
	}
L612:
	;
	goto L613
L613:
	;
	v2398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+1456)) = v2398
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+40)) = v2398
	*(*int64)(unsafe.Add(mBase, uint32(v2367)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+8)) = int32(3)
	F_WakeupRecovery(m)
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L1
	} else {
		goto L618
	}
L614:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_errmsg_internal(m, int32(346540), int32(0))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	F_errfinish(m, int32(486724), int32(658), int32(245742))
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L618:
	;
	goto L619
L619:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v2424))) = int32(0)
	goto L621
L620:
	;
	v2456 = *(*int64)(unsafe.Add(mBase, uint32(v2367)+32))
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2367)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+332)) = v2457
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+1456)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+8)) = int32(2)
	v2464 = int32(*(*uint8)(unsafe.Add(mBase, _consts[877])))
	if v2464 != 0 {
		goto L635
	} else {
		goto L636
	}
L621:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v2428 != 0 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L1
	} else {
		goto L625
	}
L623:
	;
	goto L624
L624:
	;
	v2431 = *(*int32)(unsafe.Add(mBase, uint32(v2372)))
	*(*int32)(unsafe.Add(mBase, uint32(v2372))) = int32(1)
	if v2431 != 0 {
		goto L626
	} else {
		goto L627
	}
L625:
	;
	goto L624
L626:
	;
	F_s_lock(m, v2372, int32(486724), int32(678), int32(245742))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L1
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2367)+8))
	switch v2439 - int32(4) {
	case 0:
		goto L630
	case 1:
		goto L632
	default:
		goto L631
	}
L629:
	;
	goto L628
L630:
	;
	goto L620
L631:
	;
	v2447 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2372))) = v2447
	v2450 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v2454 = F_WaitLatch(m, v2450, int32(33), v2447, int32(134217782))
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L1
	} else {
		goto L634
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2372))) = int32(0)
	F_pgl_exit(m, int32(1))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L634:
	;
	goto L619
L635:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+84)) = uint32(v2456)
	v2467 = int64(base.Ui64(v2456) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+80)) = uint32(v2467)
	v2475 = F_pg_snprintf(m, v18+int32(1440), int32(50), int32(503922), v18+int32(80))
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	v2537 = *(*int32)(unsafe.Add(mBase, _consts[858]))
	v2541 = *(*int32)(unsafe.Add(mBase, _consts[448]))
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v2541)+16))
	v2543 = m.T0[v2542].(func(*base.Module, int32, int32) int32)(m, v2537, v18+int32(328))
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L1
	} else {
		goto L656
	}
L638:
	;
	v2478 = v18 + int32(1440)
	if v2478&int32(3) == int32(0) {
		v2502 = v2478
		goto L641
	} else {
		goto L642
	}
L639:
	;
	goto L637
L640:
	;
	goto L639
L641:
	;
	v2506 = v2502
	goto L650
L642:
	;
	v2486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478))))
	if v2486 == int32(0) {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	goto L639
L644:
	;
	goto L645
L645:
	;
	v2491 = v2478
	goto L646
L646:
	;
	v2495 = v2491 + int32(1)
	if v2495&int32(3) == int32(0) {
		v2502 = v2495
		goto L641
	} else {
		goto L648
	}
L647:
	;
	goto L640
L648:
	;
	v2500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2495))))
	if v2500 != 0 {
		v2491 = v2495
		goto L646
	} else {
		goto L649
	}
L649:
	;
	goto L647
L650:
	;
	v2512 = *(*int32)(unsafe.Add(mBase, uint32(v2506)))
	v2515 = int32(-2139062144)
	if (int32(16843008)-v2512|v2512)&v2515 == v2515 {
		v2506 = v2506 + int32(4)
		goto L650
	} else {
		goto L652
	}
L651:
	;
	v2521 = v2506
	goto L653
L652:
	;
	goto L651
L653:
	;
	v2525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2521))))
	if v2525 != 0 {
		v2521 = v2521 + int32(1)
		goto L653
	} else {
		goto L655
	}
L654:
	;
	goto L640
L655:
	;
	goto L654
L656:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v2547 = *(*int64)(unsafe.Add(mBase, uint32(v2546)))
	goto L657
L657:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v2547
	v2555 = F_pg_snprintf(m, v18+int32(272), int32(32), int32(37711), v18-int32(-64))
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	v2558 = v18 + int32(272)
	v2561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2558))))
	v2562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2543))))
	if v2562 == int32(0) {
		v2581 = v2561
		v2582 = v2562
		goto L660
	} else {
		goto L661
	}
L659:
	;
	if v2582-v2581 == int32(0) {
		v1048 = v2456
		v1055 = v2296
		goto L300
	} else {
		goto L667
	}
L660:
	;
	goto L659
L661:
	;
	if v2561 != v2562 {
		v2581 = v2561
		v2582 = v2562
		goto L660
	} else {
		goto L662
	}
L662:
	;
	v2566 = v2543
	v2567 = v2558
	goto L663
L663:
	;
	v2570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2567)+1)))
	v2571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2566)+1)))
	if v2571 == int32(0) {
		v2581 = v2570
		v2582 = v2571
		goto L660
	} else {
		goto L665
	}
L664:
	;
	v2581 = v2570
	v2582 = v2571
	goto L660
L665:
	;
	v2574 = int32(1)
	if v2570 == v2571 {
		v2566 = v2566 + v2574
		v2567 = v2567 + v2574
		goto L663
	} else {
		goto L666
	}
L666:
	;
	goto L664
L667:
	;
	goto L301
L668:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	F_errmsg(m, int32(23558), int32(0))
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v2587
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v18 + int32(272)
	F_errdetail(m, int32(579422), v18+int32(48))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	F_errfinish(m, int32(486724), int32(326), int32(274082))
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L673:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v18 + int32(1440)
	F_errmsg(m, int32(288755), v18+int32(96))
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L1
	} else {
		goto L675
	}
L675:
	;
	F_errfinish(m, int32(486724), int32(622), int32(274082))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L1
	} else {
		goto L676
	}
L676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L678:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L1
	} else {
		goto L679
	}
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v708
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v18)+324))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v2656
	F_errmsg(m, int32(197801), v18)
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	F_errfinish(m, int32(486724), int32(277), int32(274082))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L682:
	;
	F_errmsg_internal(m, int32(18950), int32(0))
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(486724), int32(265), int32(274082))
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L1
	} else {
		goto L684
	}
L684:
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v8 = m.G0
	v10 = v8 - int32(1136)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[879]))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[880])))
	if v16 != int32(1) {
		v27 = v13
	} else {
		v20 = *(*int64)(unsafe.Add(mBase, _consts[881]))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
		v22 = base.I64_div_u_s(v20, v21)
		if l1 != v22 {
			v27 = v13
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[882]))
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
	v43 = F_pg_snprintf(m, v10+int32(112), int32(1024), int32(501251), v10+int32(32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return
	} else {
		v48 = F_BasicOpenFile(m, v10+int32(112), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+1168)) = v48
			if v48 < int32(0) {
				v54 = *(*int32)(unsafe.Add(mBase, _consts[140]))
				if v54 == int32(44) {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v83 = *(*int32)(unsafe.Add(mBase, _consts[262]))
					F_XLogFileName(m, v10+int32(48), v81, l1, v83)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(44)
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(48)
								F_errmsg(m, int32(433170), v10)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_errfinish(m, int32(487056), int32(3089), int32(277167))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
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
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(112)
							F_errmsg(m, int32(293799), v10+int32(16))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								F_errfinish(m, int32(487056), int32(3095), int32(277167))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
	v5 = int32(4381040)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[883]))
	v7 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, _consts[884])) = v7
	*(*int32)(unsafe.Add(mBase, _consts[885])) = v7
	v16 = m.G0
	v17 = int32(16)
	v18 = v16 - v17
	m.G0 = v18
	F___gettimeofday(m, v18)
	mBase = m.M
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	m.G0 = v18 + v17
	v30 = v22 + v21*int64(1000000) - int64(946684800000000)
	F_enlargeStringInfo(m, int32(4381040), int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		v35 = int32(4381044)
		v36 = *(*int32)(unsafe.Add(mBase, _consts[885]))
		v37 = int32(4381040)
		v38 = *(*int32)(unsafe.Add(mBase, _consts[883]))
		v40 = int64(56)
		v42 = int64(65280)
		v44 = int64(40)
		v47 = int64(16711680)
		v49 = int64(24)
		v51 = int64(4278190080)
		v53 = int64(8)
		*(*int64)(unsafe.Add(mBase, uint32(v36+v38))) = v30<<(uint(v40)%64) | v30&v42<<(uint(v44)%64) | (v30&v47<<(uint(v49)%64) | v30&v51<<(uint(v53)%64)) | (int64(base.Ui64(v30)>>(uint(v53)%64))&v51 | int64(base.Ui64(v30)>>(uint(v49)%64))&v47 | (int64(base.Ui64(v30)>>(uint(v44)%64))&v42 | int64(base.Ui64(v30)>>(uint(v40)%64))))
		*(*int32)(unsafe.Add(mBase, _consts[885])) = v36 + int32(8)
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
		v83 = *(*int32)(unsafe.Add(mBase, _consts[883]))
		v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
		*(*int64)(unsafe.Add(mBase, uint32(v81)+17)) = v84
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
		v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
		v91 = *(*int32)(unsafe.Add(mBase, _consts[314]))
		v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
		m.T0[v92].(func(*base.Module, int32, int32, int32))(m, int32(100), v88, v89)
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			v96 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			if v96 != 0 {
				F_ProcessInterrupts(m)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, _consts[314]))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
					v102 = m.T0[v101].(func(*base.Module) int32)(m)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						if v102 == int32(0) {
							v107 = *(*int64)(unsafe.Add(mBase, _consts[886]))
							v109 = *(*int32)(unsafe.Add(mBase, _consts[887]))
							v111 = base.I32_div_s(v109, int32(2))
							if v30 < v107+base.I64_extend_i32_s(v111)*int64(1000) {
								v118 = *(*int32)(unsafe.Add(mBase, _consts[314]))
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
				v100 = *(*int32)(unsafe.Add(mBase, _consts[314]))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
				v102 = m.T0[v101].(func(*base.Module) int32)(m)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					if v102 == int32(0) {
						v107 = *(*int64)(unsafe.Add(mBase, _consts[886]))
						v109 = *(*int32)(unsafe.Add(mBase, _consts[887]))
						v111 = base.I32_div_s(v109, int32(2))
						if v30 < v107+base.I64_extend_i32_s(v111)*int64(1000) {
							v118 = *(*int32)(unsafe.Add(mBase, _consts[314]))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
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
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int64
	_ = v571
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = int32(-1)
	v17 = v3
	v18 = v3
	v19 = v3
	v20 = v11
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v14 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v570 = int32(m.ExcTag)
	v571 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v570 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L6:
	;
	v25 = v20 - int32(160)
	m.G0 = v25
	*(*int32)(unsafe.Add(mBase, _consts[264])) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		v568 = v25
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v404 = v17
	v405 = v18
	v406 = v19
	v407 = v20
	goto L8
L8:
	;
	if v406 != 0 {
		goto L119
	} else {
		goto L120
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v37 = int32(913)
	v39 = m.G0
	v41 = v39 - int32(144)
	m.G0 = v41
	switch int32(915) {
	case 0, 2:
		v51 = v37
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v81 = int32(915)
	v83 = m.G0
	v85 = v83 - int32(144)
	m.G0 = v85
	switch int32(917) {
	case 0, 2:
		v95 = v81
		goto L24
	default:
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v51
	F_sigemptyset(m, v41+int32(8))
	mBase = m.M
	goto L14
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[321])) = v37
	v51 = int32(4729)
	goto L11
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+136)) = int32(268435456)
	v63 = v41 + int32(4)
	goto L18
L16:
	;
	m.G0 = v41 + int32(144)
	goto L10
L18:
	;
	goto L19
L19:
	;
	if v63 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = F___memcpy(m, int32(4635612), v63, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v125 = int32(915)
	v127 = m.G0
	v129 = v127 - int32(144)
	m.G0 = v129
	switch int32(917) {
	case 0, 2:
		v139 = v125
		goto L37
	default:
		goto L38
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v95
	F_sigemptyset(m, v85+int32(8))
	mBase = m.M
	goto L27
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v81
	v95 = int32(4729)
	goto L24
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+136)) = int32(268435456)
	v107 = v85 + int32(4)
	goto L31
L29:
	;
	m.G0 = v85 + int32(144)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if v107 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v118 = F___memcpy(m, int32(4635752), v107, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v169 = int32(-2)
	v171 = m.G0
	v173 = v171 - int32(144)
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
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v139
	F_sigemptyset(m, v129+int32(8))
	mBase = m.M
	goto L40
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[323])) = v125
	v139 = int32(4729)
	goto L37
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+136)) = int32(268435456)
	v151 = v129 + int32(4)
	goto L44
L42:
	;
	m.G0 = v129 + int32(144)
	goto L36
L44:
	;
	goto L45
L45:
	;
	if v151 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v162 = F___memcpy(m, int32(4637572), v151, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v213 = int32(-2)
	v215 = m.G0
	v217 = v215 - int32(144)
	m.G0 = v217
	switch int32(0) {
	case 0, 2:
		v227 = v213
		goto L63
	default:
		goto L64
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v183
	F_sigemptyset(m, v173+int32(8))
	mBase = m.M
	goto L53
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[724])) = v169
	v183 = int32(4729)
	goto L50
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+136)) = int32(268435456)
	v195 = v173 + int32(4)
	goto L57
L55:
	;
	m.G0 = v173 + int32(144)
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
	v206 = F___memcpy(m, int32(4637432), v195, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v257 = int32(916)
	v259 = m.G0
	v261 = v259 - int32(144)
	m.G0 = v261
	switch int32(918) {
	case 0, 2:
		v271 = v257
		goto L76
	default:
		goto L77
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v227
	F_sigemptyset(m, v217+int32(8))
	mBase = m.M
	goto L66
L64:
	;
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v213
	v227 = int32(4729)
	goto L63
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+136)) = int32(268435456)
	v239 = v217 + int32(4)
	goto L70
L68:
	;
	m.G0 = v217 + int32(144)
	goto L62
L70:
	;
	goto L71
L71:
	;
	if v239 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v250 = F___memcpy(m, int32(4637292), v239, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v301 = int32(-2)
	v303 = m.G0
	v305 = v303 - int32(144)
	m.G0 = v305
	switch int32(0) {
	case 0, 2:
		v315 = v301
		goto L89
	default:
		goto L90
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+4)) = v271
	F_sigemptyset(m, v261+int32(8))
	mBase = m.M
	goto L79
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[647])) = v257
	v271 = int32(4729)
	goto L76
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+136)) = int32(268435456)
	v283 = v261 + int32(4)
	goto L83
L81:
	;
	m.G0 = v261 + int32(144)
	goto L75
L83:
	;
	goto L84
L84:
	;
	if v283 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v294 = F___memcpy(m, int32(4636872), v283, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	v345 = int32(0)
	v347 = m.G0
	v349 = v347 - int32(144)
	m.G0 = v349
	switch int32(2) {
	case 0, 2:
		v359 = v345
		goto L102
	default:
		goto L103
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+4)) = v315
	F_sigemptyset(m, v305+int32(8))
	mBase = m.M
	goto L92
L90:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v301
	v315 = int32(4729)
	goto L89
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+136)) = int32(268435456)
	v327 = v305 + int32(4)
	goto L96
L94:
	;
	m.G0 = v305 + int32(144)
	goto L88
L96:
	;
	goto L97
L97:
	;
	if v327 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v338 = F___memcpy(m, int32(4637152), v327, int32(140))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v25
	v389 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v394 = F_AllocSetContextCreateInternal(m, v389, int32(212811), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		v568 = v25
		goto L5
	} else {
		goto L114
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+4)) = v359
	F_sigemptyset(m, v349+int32(8))
	mBase = m.M
	goto L104
L103:
	;
	*(*int32)(unsafe.Add(mBase, _consts[650])) = v345
	v359 = int32(4729)
	goto L102
L104:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v349)+136)) = int32(268435457)
	v371 = v349 + int32(4)
	goto L109
L107:
	;
	m.G0 = v349 + int32(144)
	goto L101
L109:
	;
	goto L110
L110:
	;
	if v371 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v382 = F___memcpy(m, int32(4637852), v371, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v394
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v11 + int32(4)
	goto L118
L116:
	;
	v404 = v25
	v405 = v394
	v406 = int32(0)
	v407 = v25
	goto L8
L118:
	;
	goto L116
L119:
	;
	v408 = int32(4465212)
	v410 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v410 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[385])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	F_EmitErrorReport(m)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, _consts[416])) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	F_sigprocmask(m, int32(4377784), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L134
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_LWLockReleaseAll(m)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L124
	}
L124:
	;
	v430 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_UnlockBuffers(m)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_smgrdestroyall(m)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	F_FlushErrorState(m)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_MemoryContextReset(m, v405)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L132
	}
L132:
	;
	v472 = int32(4465212)
	v474 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v474 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	F_pg_usleep(m, int32(1000000))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L133
	}
L133:
	;
	goto L121
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_SetWalWriterSleeping(m, int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L135
	}
L135:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v500 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+60)) = v500
	v504 = int32(0)
	v508 = int32(50)
	goto L136
L136:
	;
	v512 = base.B2i32(v508 < int32(2))
	if v512 != v504&int32(1) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_SetWalWriterSleeping(m, base.B2i32(v508 < int32(2)))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L141
	}
L139:
	;
	v522 = v504
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	v526 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v526))) = int32(0)
	goto L142
L141:
	;
	v522 = v512
	goto L140
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_ProcessMainLoopInterrupts(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	v535 = F_XLogBackgroundFlush(m)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v404
	v545 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v548 = *(*int32)(unsafe.Add(mBase, _consts[271]))
	if v535 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v555 = int32(50)
	goto L148
L147:
	;
	v555 = v508 - base.B2i32(int32(0) < v508)
	goto L148
L148:
	;
	if int32(0) < v555 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v558 = v548
	goto L151
L150:
	;
	v558 = v548 * int32(25)
	goto L151
L151:
	;
	v560 = F_WaitLatch(m, v545, int32(41), v558, int32(83886097))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		v568 = v407
		goto L5
	} else {
		goto L152
	}
L152:
	;
	v504 = v522
	v508 = v555
	goto L136
L153:
	;
	v575 = int32(v571)
	m.G0 = v568
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v578)))
	if v11+int32(4) == v582 {
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
	if v585 != 0 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v585 = v584
	goto L158
L157:
	;
	v585 = int32(0)
	goto L158
L158:
	;
	goto L155
L159:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v14 = v585
	v17 = v586
	v18 = v587
	v19 = v577
	v20 = v568
	goto L1
L160:
	;
	goto L161
L161:
	;
	F___wasm_longjmp(m, v578, v577)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
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
