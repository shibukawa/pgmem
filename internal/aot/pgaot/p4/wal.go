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
func F_WalRcvShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_WalRcvShmemRequest_0), int64(1480), int32(_a_F_WalRcvShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_WalReceiverMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int64
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1023 int64
	_ = v1023
	var v1024 int64
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int64
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int64
	_ = v1095
	var v1098 int64
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int64
	_ = v1118
	var v1119 int64
	_ = v1119
	var v1129 int64
	_ = v1129
	var v1133 int64
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int64
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int64
	_ = v1146
	var v1147 int64
	_ = v1147
	var v1151 int64
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int64
	_ = v1190
	var v1191 int64
	_ = v1191
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1406 int64
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1433 int32
	_ = v1433
	var v1438 int64
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int64
	_ = v1453
	var v1454 int64
	_ = v1454
	var v1462 int64
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1471 int64
	_ = v1471
	var v1481 int64
	_ = v1481
	var v1486 int32
	_ = v1486
	var v1490 int64
	_ = v1490
	var v1493 int64
	_ = v1493
	var v1499 int64
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int64
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1556 int32
	_ = v1556
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1567 int64
	_ = v1567
	var v1568 int64
	_ = v1568
	var v1576 int64
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1584 int32
	_ = v1584
	var v1585 int64
	_ = v1585
	var v1595 int64
	_ = v1595
	var v1600 int32
	_ = v1600
	var v1604 int64
	_ = v1604
	var v1607 int64
	_ = v1607
	var v1613 int64
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1617 int64
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int64
	_ = v1663
	var v1664 int64
	_ = v1664
	var v1672 int64
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1678 int64
	_ = v1678
	var v1681 int64
	_ = v1681
	var v1690 int64
	_ = v1690
	var v1691 int64
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1712 int32
	_ = v1712
	var v1713 int64
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int64
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int64
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1729 int64
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1751 int64
	_ = v1751
	var v1753 int64
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1766 int64
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1791 int64
	_ = v1791
	var v1792 int64
	_ = v1792
	var v1796 int64
	_ = v1796
	var v1801 int32
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1836 int64
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1866 int64
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1878 int64
	_ = v1878
	var v1879 int64
	_ = v1879
	var v1883 int64
	_ = v1883
	var v1933 int32
	_ = v1933
	var v1934 int64
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1945 int32
	_ = v1945
	var v1950 int64
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1970 int32
	_ = v1970
	var v1971 int64
	_ = v1971
	var v1975 int64
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1989 int64
	_ = v1989
	var v1992 int32
	_ = v1992
	var v2001 int32
	_ = v2001
	var v2002 int64
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int64
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2020 int32
	_ = v2020
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2034 int64
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2063 int32
	_ = v2063
	var v2068 int32
	_ = v2068
	var v2071 int64
	_ = v2071
	var v2072 int64
	_ = v2072
	var v2087 int32
	_ = v2087
	var v2089 int64
	_ = v2089
	var v2092 int64
	_ = v2092
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2100 int64
	_ = v2100
	var v2102 int64
	_ = v2102
	var v2103 int64
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2125 int32
	_ = v2125
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2167 int64
	_ = v2167
	var v2170 int64
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2233 int64
	_ = v2233
	var v2235 int64
	_ = v2235
	var v2237 int64
	_ = v2237
	var v2239 int64
	_ = v2239
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2248 int64
	_ = v2248
	var v2249 int64
	_ = v2249
	var v2257 int64
	_ = v2257
	var v2259 int64
	_ = v2259
	var v2261 int64
	_ = v2261
	var v2263 int64
	_ = v2263
	var v2269 int64
	_ = v2269
	var v2278 int64
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2312 int32
	_ = v2312
	var v2317 int32
	_ = v2317
	var v2322 int32
	_ = v2322
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2331 int64
	_ = v2331
	var v2332 int64
	_ = v2332
	var v2340 int64
	_ = v2340
	var v2342 int64
	_ = v2342
	var v2345 int64
	_ = v2345
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2356 int32
	_ = v2356
	var v2364 int32
	_ = v2364
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2418 int32
	_ = v2418
	var v2423 int32
	_ = v2423
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2434 int32
	_ = v2434
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2455 int32
	_ = v2455
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2494 int32
	_ = v2494
	var v2499 int32
	_ = v2499
	var v2510 int32
	_ = v2510
	var v2518 int32
	_ = v2518
	var v2522 int32
	_ = v2522
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2529 int64
	_ = v2529
	var v2532 int64
	_ = v2532
	var v2533 int64
	_ = v2533
	var v2534 int64
	_ = v2534
	var v2537 int64
	_ = v2537
	var v2540 int32
	_ = v2540
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2551 int32
	_ = v2551
	var v2555 int32
	_ = v2555
	var v2559 int32
	_ = v2559
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2574 int32
	_ = v2574
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2586 int32
	_ = v2586
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2595 int32
	_ = v2595
	var v2603 int32
	_ = v2603
	var v2607 int32
	_ = v2607
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2617 int32
	_ = v2617
	var v2623 int32
	_ = v2623
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2655 int32
	_ = v2655
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2665 int64
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2670 int32
	_ = v2670
	var v2674 int64
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2681 int64
	_ = v2681
	var v2684 int32
	_ = v2684
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2721 int32
	_ = v2721
	var v2726 int32
	_ = v2726
	var v2729 int32
	_ = v2729
	var v2733 int32
	_ = v2733
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2755 int32
	_ = v2755
	var v2760 int32
	_ = v2760
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1520)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+336)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v20)+332)) = v7
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[0]))
	v30 = int32(1456)
	v31 = v29 + v30
	v34 = base.AtomicRmwXchg32(m, v29, v30, int32(1))
	if v34 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_s_lock(m, v31, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	switch v38 {
	case 0:
		goto L9
	case 1:
		goto L7
	default:
		goto L8
	case 6:
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[1]))
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+1453)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v68
	v75 = v20 + int32(416)
	v77 = v29 + int32(104)
	goto L19
L8:
	;
	v51 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v31))), uint32(v51))
	F_errstart_cold(m, int32(24), v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L13
	}
L9:
	;
	v41 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+1456)), uint32(v41))
	F_ConditionVariableBroadcast(m, v29+int32(12))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(0)
	goto L9
L11:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(219), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
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
	v198 = v20 + int32(352)
	v200 = v29 + int32(1388)
	goto L50
L17:
	;
	v194 = F_strlen(m, v183)
	mBase = m.M
	goto L16
L19:
	;
	goto L20
L20:
	;
	v84 = int32(1023)
	if (v75^v77)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v187)
	goto L17
L22:
	;
	v168 = v163
	v169 = v164
	v170 = v165
	goto L43
L23:
	;
	if v158 == int32(0) {
		v183 = v156
		v184 = v157
		goto L21
	} else {
		goto L42
	}
L24:
	;
	v156 = v77
	v157 = v75
	v158 = v84
	goto L23
L25:
	;
	goto L26
L26:
	;
	v88 = int32(0)
	if base.B2i32(v77&int32(3) == v88)|int32(0) == v88 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v124 == int32(0) {
		v183 = v121
		v184 = v122
		goto L21
	} else {
		goto L36
	}
L28:
	;
	v100 = v77
	v101 = v75
	v102 = v84
	goto L31
L29:
	;
	goto L30
L30:
	;
	v121 = v77
	v122 = v75
	v123 = v84
	v124 = int32(1)
	goto L27
L31:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v104)
	if v104 == int32(0) {
		v163 = v100
		v164 = v101
		v165 = v102
		goto L22
	} else {
		goto L33
	}
L32:
	;
	v121 = v115
	v122 = v109
	v123 = v111
	v124 = v113
	goto L27
L33:
	;
	v108 = int32(1)
	v109 = v101 + v108
	v111 = v102 - v108
	v112 = int32(0)
	v113 = base.B2i32(v111 != v112)
	v115 = v100 + v108
	if v115&int32(3) == v112 {
		v121 = v115
		v122 = v109
		v123 = v111
		v124 = v113
		goto L27
	} else {
		goto L34
	}
L34:
	;
	if v111 != 0 {
		v100 = v115
		v101 = v109
		v102 = v111
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	if base.B2i32(v127 == int32(0))|base.B2i32(base.Ui32(v123) < base.Ui32(int32(4))) != 0 {
		v156 = v121
		v157 = v122
		v158 = v123
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v134 = v121
	v135 = v122
	v136 = v123
	goto L38
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v142 = int32(-2139062144)
	if (int32(16843008)-v139|v139)&v142 != v142 {
		v163 = v134
		v164 = v135
		v165 = v136
		goto L22
	} else {
		goto L40
	}
L39:
	;
	v156 = v150
	v157 = v148
	v158 = v152
	goto L23
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v139
	v147 = int32(4)
	v148 = v135 + v147
	v150 = v134 + v147
	v152 = v136 - v147
	if base.Ui32(int32(3)) < base.Ui32(v152) {
		v134 = v150
		v135 = v148
		v136 = v152
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v163 = v156
	v164 = v157
	v165 = v158
	goto L22
L43:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v172)
	if v172 == int32(0) {
		v183 = v168
		v184 = v169
		goto L21
	} else {
		goto L45
	}
L44:
	;
	v183 = v179
	v184 = v177
	goto L21
L45:
	;
	v176 = int32(1)
	v177 = v169 + v176
	v179 = v168 + v176
	v181 = v170 - v176
	if v181 != 0 {
		v168 = v179
		v169 = v177
		v170 = v181
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1452)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+348)) = v322
	v327 = m.G0
	v328 = int32(16)
	v329 = v327 - v328
	m.G0 = v329
	F_gettimeofday(m, v329)
	mBase = m.M
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v329)))
	v333 = int64(*(*int32)(unsafe.Add(mBase, uint32(v329)+8)))
	m.G0 = v329 + v328
	v341 = v333 + v332*int64(1000000) - int64(946684800000000)
	goto L78
L48:
	;
	v317 = F_strlen(m, v306)
	mBase = m.M
	goto L47
L50:
	;
	goto L51
L51:
	;
	v207 = int32(63)
	if (v198^v200)&int32(3) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v310)
	goto L48
L53:
	;
	v291 = v286
	v292 = v287
	v293 = v288
	goto L74
L54:
	;
	if v281 == int32(0) {
		v306 = v279
		v307 = v280
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v279 = v200
	v280 = v198
	v281 = v207
	goto L54
L56:
	;
	goto L57
L57:
	;
	v211 = int32(0)
	if base.B2i32(v200&int32(3) == v211)|int32(0) == v211 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v247 == int32(0) {
		v306 = v244
		v307 = v245
		goto L52
	} else {
		goto L67
	}
L59:
	;
	v223 = v200
	v224 = v198
	v225 = v207
	goto L62
L60:
	;
	goto L61
L61:
	;
	v244 = v200
	v245 = v198
	v246 = v207
	v247 = int32(1)
	goto L58
L62:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v227)
	if v227 == int32(0) {
		v286 = v223
		v287 = v224
		v288 = v225
		goto L53
	} else {
		goto L64
	}
L63:
	;
	v244 = v238
	v245 = v232
	v246 = v234
	v247 = v236
	goto L58
L64:
	;
	v231 = int32(1)
	v232 = v224 + v231
	v234 = v225 - v231
	v235 = int32(0)
	v236 = base.B2i32(v234 != v235)
	v238 = v223 + v231
	if v238&int32(3) == v235 {
		v244 = v238
		v245 = v232
		v246 = v234
		v247 = v236
		goto L58
	} else {
		goto L65
	}
L65:
	;
	if v234 != 0 {
		v223 = v238
		v224 = v232
		v225 = v234
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if base.B2i32(v250 == int32(0))|base.B2i32(base.Ui32(v246) < base.Ui32(int32(4))) != 0 {
		v279 = v244
		v280 = v245
		v281 = v246
		goto L54
	} else {
		goto L68
	}
L68:
	;
	v257 = v244
	v258 = v245
	v259 = v246
	goto L69
L69:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v265 = int32(-2139062144)
	if (int32(16843008)-v262|v262)&v265 != v265 {
		v286 = v257
		v287 = v258
		v288 = v259
		goto L53
	} else {
		goto L71
	}
L70:
	;
	v279 = v273
	v280 = v271
	v281 = v275
	goto L54
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v262
	v270 = int32(4)
	v271 = v258 + v270
	v273 = v257 + v270
	v275 = v259 - v270
	if base.Ui32(int32(3)) < base.Ui32(v275) {
		v257 = v273
		v258 = v271
		v259 = v275
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v286 = v279
	v287 = v280
	v288 = v281
	goto L53
L74:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v295)
	if v295 == int32(0) {
		v306 = v291
		v307 = v292
		goto L52
	} else {
		goto L76
	}
L75:
	;
	v306 = v302
	v307 = v300
	goto L52
L76:
	;
	v299 = int32(1)
	v300 = v292 + v299
	v302 = v291 + v299
	v304 = v293 - v299
	if v304 != 0 {
		v291 = v302
		v292 = v300
		v293 = v304
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+80)) = v341
	*(*int64)(unsafe.Add(mBase, uint32(v29)+96)) = v341
	*(*int64)(unsafe.Add(mBase, uint32(v29)+72)) = v341
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v346
	v348 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+1456)), uint32(v348))
	F_on_shmem_exit(m, int32(1103), base.I64_extend_i32_u(v20+int32(348)))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v361 = m.G0
	v363 = v361 - int32(32)
	m.G0 = v363
	v366 = int32(967)
	switch v366 {
	case 0, 2:
		goto L81
	default:
		goto L82
	}
L80:
	;
	v401 = int32(0)
	v403 = m.G0
	v405 = v403 - int32(32)
	m.G0 = v405
	switch v401 {
	case 0, 2:
		goto L91
	default:
		goto L92
	}
L81:
	;
	F_sigemptyset(m, v363+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v363)+24)) = int32(268435456)
	switch v366 {
	case 0:
		goto L86
	default:
		goto L84
	case 2:
		goto L85
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[3])) = int32(965)
	goto L81
L83:
	;
	goto L88
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = int32(0)
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = int32(-2)
	goto L83
L88:
	;
	goto L89
L89:
	;
	v395 = F___sigaction(m, int32(1), v363+int32(12), int32(0))
	mBase = m.M
	m.G0 = v363 + int32(32)
	goto L80
L90:
	;
	v445 = m.G0
	v447 = v445 - int32(32)
	m.G0 = v447
	v450 = int32(974)
	switch v450 {
	case 0, 2:
		goto L101
	default:
		goto L102
	}
L91:
	;
	F_sigemptyset(m, v405+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v405)+24)) = int32(268435456)
	switch v401 {
	case 0:
		goto L96
	default:
		goto L94
	case 2:
		goto L95
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[4])) = int32(-2)
	goto L91
L93:
	;
	goto L98
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v405)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L93
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405)+12)) = int32(0)
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405)+12)) = int32(-2)
	goto L93
L98:
	;
	goto L99
L99:
	;
	v437 = F___sigaction(m, int32(2), v405+int32(12), int32(0))
	mBase = m.M
	m.G0 = v405 + int32(32)
	goto L90
L100:
	;
	v485 = int32(0)
	v487 = m.G0
	v489 = v487 - int32(32)
	m.G0 = v489
	switch v485 {
	case 0, 2:
		goto L111
	default:
		goto L112
	}
L101:
	;
	F_sigemptyset(m, v447+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v447)+24)) = int32(268435456)
	switch v450 {
	case 0:
		goto L106
	default:
		goto L104
	case 2:
		goto L105
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[5])) = int32(972)
	goto L101
L103:
	;
	goto L108
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v447)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v447)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L103
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v447)+12)) = int32(0)
	goto L103
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v447)+12)) = int32(-2)
	goto L103
L108:
	;
	goto L109
L109:
	;
	v479 = F___sigaction(m, int32(15), v447+int32(12), int32(0))
	mBase = m.M
	m.G0 = v447 + int32(32)
	goto L100
L110:
	;
	v527 = int32(0)
	v529 = m.G0
	v531 = v529 - int32(32)
	m.G0 = v531
	switch v527 {
	case 0, 2:
		goto L121
	default:
		goto L122
	}
L111:
	;
	F_sigemptyset(m, v489+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v489)+24)) = int32(268435456)
	switch v485 {
	case 0:
		goto L116
	default:
		goto L114
	case 2:
		goto L115
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[6])) = int32(-2)
	goto L111
L113:
	;
	goto L118
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v489)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489)+12)) = int32(0)
	goto L113
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v489)+12)) = int32(-2)
	goto L113
L118:
	;
	goto L119
L119:
	;
	v521 = F___sigaction(m, int32(14), v489+int32(12), int32(0))
	mBase = m.M
	m.G0 = v489 + int32(32)
	goto L110
L120:
	;
	v571 = m.G0
	v573 = v571 - int32(32)
	m.G0 = v573
	v576 = int32(970)
	switch v576 {
	case 0, 2:
		goto L131
	default:
		goto L132
	}
L121:
	;
	F_sigemptyset(m, v531+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v531)+24)) = int32(268435456)
	switch v527 {
	case 0:
		goto L126
	default:
		goto L124
	case 2:
		goto L125
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[7])) = int32(-2)
	goto L121
L123:
	;
	goto L128
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v531)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v531)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L123
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v531)+12)) = int32(0)
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v531)+12)) = int32(-2)
	goto L123
L128:
	;
	goto L129
L129:
	;
	v563 = F___sigaction(m, int32(13), v531+int32(12), int32(0))
	mBase = m.M
	m.G0 = v531 + int32(32)
	goto L120
L130:
	;
	v611 = int32(0)
	v613 = m.G0
	v615 = v613 - int32(32)
	m.G0 = v615
	switch v611 {
	case 0, 2:
		goto L141
	default:
		goto L142
	}
L131:
	;
	F_sigemptyset(m, v573+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v573)+24)) = int32(268435456)
	switch v576 {
	case 0:
		goto L136
	default:
		goto L134
	case 2:
		goto L135
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[8])) = int32(968)
	goto L131
L133:
	;
	goto L138
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v573)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v573)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v573)+12)) = int32(0)
	goto L133
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v573)+12)) = int32(-2)
	goto L133
L138:
	;
	goto L139
L139:
	;
	v605 = F___sigaction(m, int32(10), v573+int32(12), int32(0))
	mBase = m.M
	m.G0 = v573 + int32(32)
	goto L130
L140:
	;
	v655 = m.G0
	v657 = v655 - int32(32)
	m.G0 = v657
	v659 = int32(2)
	switch v659 {
	case 0, 2:
		goto L151
	default:
		goto L152
	}
L141:
	;
	F_sigemptyset(m, v615+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v615)+24)) = int32(268435456)
	switch v611 {
	case 0:
		goto L146
	default:
		goto L144
	case 2:
		goto L145
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[9])) = int32(-2)
	goto L141
L143:
	;
	goto L148
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v615)+12)) = int32(_a_F_WalReceiverMain_4)
	goto L143
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+12)) = int32(0)
	goto L143
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v615)+12)) = int32(-2)
	goto L143
L148:
	;
	goto L149
L149:
	;
	v647 = F___sigaction(m, int32(12), v615+int32(12), int32(0))
	mBase = m.M
	m.G0 = v615 + int32(32)
	goto L140
L150:
	;
	F_load_file(m, int32(_a_F_WalReceiverMain_5), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L160
	}
L151:
	;
	F_sigemptyset(m, v657+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v657)+24)) = int32(268435456)
	switch v659 {
	case 0:
		goto L156
	default:
		goto L154
	case 2:
		goto L155
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[10])) = int32(0)
	goto L151
L153:
	;
	goto L157
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v657)+12)) = int32(_a_F_WalReceiverMain_4)
	v682 = int32(268435461)
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+12)) = int32(0)
	v682 = int32(268435457)
	goto L153
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+12)) = int32(-2)
	v682 = int32(268435457)
	goto L153
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+24)) = v682
	goto L159
L159:
	;
	v689 = F___sigaction(m, int32(17), v657+int32(12), int32(0))
	mBase = m.M
	m.G0 = v657 + int32(32)
	goto L150
L160:
	;
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	if v698 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_WalReceiverMain_6), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L1
	} else {
		goto L667
	}
L164:
	;
	v705 = base.AtomicRmwXchg32(m, v31, int32(0), int32(1))
	if v705 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_s_lock(m, v31, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+1384)) = v709
	base.MemoryFill(m, v77, v709, int32(1279))
	v714 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+1453)) = uint8(v714)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+1456)), uint32(v709))
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[12]))
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	if v728 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	goto L167
L169:
	;
	v729 = v726
	goto L171
L170:
	;
	v729 = int32(_a_F_WalReceiverMain_7)
	goto L171
L171:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	v735 = m.T0[v734].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20+int32(416), v714, v709, v709, v729, v20+int32(340))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13])) = v735
	if v735 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v739 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	v741 = m.T0[v740].(func(*base.Module, int32) int32)(m, v735)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L1
	} else {
		goto L663
	}
L176:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)+12))
	m.T0[v751].(func(*base.Module, int32, int32, int32))(m, v744, v20+int32(336), v20+int32(332))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v756 = base.AtomicRmwXchg32(m, v31, int32(0), int32(1))
	if v756 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	F_s_lock(m, v31, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v741 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	goto L180
L182:
	;
	goto L188
L183:
	;
	goto L184
L184:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v20)+336))
	if v879 != 0 {
		goto L216
	} else {
		goto L217
	}
L185:
	;
	goto L184
L186:
	;
	v876 = F_strlen(m, v865)
	mBase = m.M
	goto L185
L188:
	;
	goto L189
L189:
	;
	v766 = int32(1023)
	if (v77^v741)&int32(3) != 0 {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v869 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v866))) = uint8(v869)
	goto L186
L191:
	;
	v850 = v845
	v851 = v846
	v852 = v847
	goto L212
L192:
	;
	if v840 == int32(0) {
		v865 = v838
		v866 = v839
		goto L190
	} else {
		goto L211
	}
L193:
	;
	v838 = v741
	v839 = v77
	v840 = v766
	goto L192
L194:
	;
	goto L195
L195:
	;
	v770 = int32(0)
	if base.B2i32(v741&int32(3) == v770)|int32(0) == v770 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	if v806 == int32(0) {
		v865 = v803
		v866 = v804
		goto L190
	} else {
		goto L205
	}
L197:
	;
	v782 = v741
	v783 = v77
	v784 = v766
	goto L200
L198:
	;
	goto L199
L199:
	;
	v803 = v741
	v804 = v77
	v805 = v766
	v806 = int32(1)
	goto L196
L200:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782))))
	*(*uint8)(unsafe.Add(mBase, uint32(v783))) = uint8(v786)
	if v786 == int32(0) {
		v845 = v782
		v846 = v783
		v847 = v784
		goto L191
	} else {
		goto L202
	}
L201:
	;
	v803 = v797
	v804 = v791
	v805 = v793
	v806 = v795
	goto L196
L202:
	;
	v790 = int32(1)
	v791 = v783 + v790
	v793 = v784 - v790
	v794 = int32(0)
	v795 = base.B2i32(v793 != v794)
	v797 = v782 + v790
	if v797&int32(3) == v794 {
		v803 = v797
		v804 = v791
		v805 = v793
		v806 = v795
		goto L196
	} else {
		goto L203
	}
L203:
	;
	if v793 != 0 {
		v782 = v797
		v783 = v791
		v784 = v793
		goto L200
	} else {
		goto L204
	}
L204:
	;
	goto L201
L205:
	;
	v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803))))
	if base.B2i32(v809 == int32(0))|base.B2i32(base.Ui32(v805) < base.Ui32(int32(4))) != 0 {
		v838 = v803
		v839 = v804
		v840 = v805
		goto L192
	} else {
		goto L206
	}
L206:
	;
	v816 = v803
	v817 = v804
	v818 = v805
	goto L207
L207:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	v824 = int32(-2139062144)
	if (int32(16843008)-v821|v821)&v824 != v824 {
		v845 = v816
		v846 = v817
		v847 = v818
		goto L191
	} else {
		goto L209
	}
L208:
	;
	v838 = v832
	v839 = v830
	v840 = v834
	goto L192
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v817))) = v821
	v829 = int32(4)
	v830 = v817 + v829
	v832 = v816 + v829
	v834 = v818 - v829
	if base.Ui32(int32(3)) < base.Ui32(v834) {
		v816 = v832
		v817 = v830
		v818 = v834
		goto L207
	} else {
		goto L210
	}
L210:
	;
	goto L208
L211:
	;
	v845 = v838
	v846 = v839
	v847 = v840
	goto L191
L212:
	;
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850))))
	*(*uint8)(unsafe.Add(mBase, uint32(v851))) = uint8(v854)
	if v854 == int32(0) {
		v865 = v850
		v866 = v851
		goto L190
	} else {
		goto L214
	}
L213:
	;
	v865 = v861
	v866 = v859
	goto L190
L214:
	;
	v858 = int32(1)
	v859 = v851 + v858
	v861 = v850 + v858
	v863 = v852 - v858
	if v863 != 0 {
		v850 = v861
		v851 = v859
		v852 = v863
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v881 = v29 + int32(1128)
	goto L222
L217:
	;
	goto L218
L218:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v20)+332))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+1384)) = v1001
	v1003 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29)+1456)), uint32(v1003))
	if v741 != 0 {
		goto L250
	} else {
		goto L251
	}
L219:
	;
	goto L218
L220:
	;
	v998 = F_strlen(m, v987)
	mBase = m.M
	goto L219
L222:
	;
	goto L223
L223:
	;
	v888 = int32(254)
	if (v881^v879)&int32(3) != 0 {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	v991 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v988))) = uint8(v991)
	goto L220
L225:
	;
	v972 = v967
	v973 = v968
	v974 = v969
	goto L246
L226:
	;
	if v962 == int32(0) {
		v987 = v960
		v988 = v961
		goto L224
	} else {
		goto L245
	}
L227:
	;
	v960 = v879
	v961 = v881
	v962 = v888
	goto L226
L228:
	;
	goto L229
L229:
	;
	v892 = int32(0)
	if base.B2i32(v879&int32(3) == v892)|int32(0) == v892 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v928 == int32(0) {
		v987 = v925
		v988 = v926
		goto L224
	} else {
		goto L239
	}
L231:
	;
	v904 = v879
	v905 = v881
	v906 = v888
	goto L234
L232:
	;
	goto L233
L233:
	;
	v925 = v879
	v926 = v881
	v927 = v888
	v928 = int32(1)
	goto L230
L234:
	;
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904))))
	*(*uint8)(unsafe.Add(mBase, uint32(v905))) = uint8(v908)
	if v908 == int32(0) {
		v967 = v904
		v968 = v905
		v969 = v906
		goto L225
	} else {
		goto L236
	}
L235:
	;
	v925 = v919
	v926 = v913
	v927 = v915
	v928 = v917
	goto L230
L236:
	;
	v912 = int32(1)
	v913 = v905 + v912
	v915 = v906 - v912
	v916 = int32(0)
	v917 = base.B2i32(v915 != v916)
	v919 = v904 + v912
	if v919&int32(3) == v916 {
		v925 = v919
		v926 = v913
		v927 = v915
		v928 = v917
		goto L230
	} else {
		goto L237
	}
L237:
	;
	if v915 != 0 {
		v904 = v919
		v905 = v913
		v906 = v915
		goto L234
	} else {
		goto L238
	}
L238:
	;
	goto L235
L239:
	;
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v925))))
	if base.B2i32(v931 == int32(0))|base.B2i32(base.Ui32(v927) < base.Ui32(int32(4))) != 0 {
		v960 = v925
		v961 = v926
		v962 = v927
		goto L226
	} else {
		goto L240
	}
L240:
	;
	v938 = v925
	v939 = v926
	v940 = v927
	goto L241
L241:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v938)))
	v946 = int32(-2139062144)
	if (int32(16843008)-v943|v943)&v946 != v946 {
		v967 = v938
		v968 = v939
		v969 = v940
		goto L225
	} else {
		goto L243
	}
L242:
	;
	v960 = v954
	v961 = v952
	v962 = v956
	goto L226
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939))) = v943
	v951 = int32(4)
	v952 = v939 + v951
	v954 = v938 + v951
	v956 = v940 - v951
	if base.Ui32(int32(3)) < base.Ui32(v956) {
		v938 = v954
		v939 = v952
		v940 = v956
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v967 = v960
	v968 = v961
	v969 = v962
	goto L225
L246:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	*(*uint8)(unsafe.Add(mBase, uint32(v973))) = uint8(v976)
	if v976 == int32(0) {
		v987 = v972
		v988 = v973
		goto L224
	} else {
		goto L248
	}
L247:
	;
	v987 = v983
	v988 = v981
	goto L224
L248:
	;
	v980 = int32(1)
	v981 = v973 + v980
	v983 = v972 + v980
	v985 = v974 - v980
	if v985 != 0 {
		v972 = v983
		v973 = v981
		v974 = v985
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	F_pfree(m, v741)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v20)+336))
	if v1008 != 0 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L252
L254:
	;
	F_pfree(m, v1008)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	F_initStringInfo(m, int32(_a_F_WalReceiverMain_8))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	v1016 = int32(1)
	v1023 = v320
	v1024 = int64(0)
	v1027 = int32(0)
	v1031 = v1016
	v1034 = v7
	goto L260
L259:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L1
	} else {
		goto L662
	}
L260:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+16))
	v1047 = m.T0[v1046].(func(*base.Module, int32, int32, int32) int32)(m, v1039, v20+int32(344), v20+int32(280))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L262
	}
L261:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2711 = m.ExcPending
	if v2711 != 0 {
		goto L1
	} else {
		goto L658
	}
L262:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[14]))
	v1051 = *(*int64)(unsafe.Add(mBase, uint32(v1050)))
	goto L263
L263:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+224)) = v1051
	v1054 = v20 + int32(288)
	v1059 = F_pg_snprintf(m, v1054, int32(32), int32(_a_F_WalReceiverMain_9), v20+int32(224))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1054))))
	if base.B2i32(v1063 == int32(0))|base.B2i32(v1063 != v1066) != 0 {
		v1084 = v1063
		v1085 = v1066
		goto L268
	} else {
		goto L269
	}
L265:
	;
	v2518 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	if v2518 < int32(0) {
		goto L612
	} else {
		goto L613
	}
L266:
	;
	if v1398 == int32(0) {
		v2510 = v1031
		goto L265
	} else {
		goto L608
	}
L267:
	;
	if v1084-v1085 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L268:
	;
	goto L267
L269:
	;
	v1069 = v1047
	v1070 = v1054
	goto L270
L270:
	;
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070)+1)))
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1069)+1)))
	if v1074 == int32(0) {
		v1084 = v1074
		v1085 = v1073
		goto L268
	} else {
		goto L272
	}
L271:
	;
	v1084 = v1074
	v1085 = v1073
	goto L268
L272:
	;
	v1077 = int32(1)
	if v1074 == v1073 {
		v1069 = v1069 + v1077
		v1070 = v1070 + v1077
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	F_pfree(m, v1047)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L603
	}
L277:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v20)+344))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	if base.Ui32(v1092) <= base.Ui32(v1091) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	if v1091 != v1092 {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	goto L280
L280:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L599
	}
L281:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L1
	} else {
		goto L595
	}
L282:
	;
	F_WalRcvFetchTimeLineHistoryFiles(m, v1092, v1091)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L308
	}
L283:
	;
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
	if base.Ui64(v1023) <= base.Ui64(v1095) {
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v1098 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16])))
	if base.Ui64(v1098) < base.Ui64(v1023-v1095) {
		goto L282
	} else {
		goto L285
	}
L285:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[17]))
	v1103 = int32(0)
	if (v1027|base.B2i32(v1102 <= v1103))&int32(1) == v1103 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1113 = m.G0
	v1114 = int32(16)
	v1115 = v1113 - v1114
	m.G0 = v1115
	F_gettimeofday(m, v1115)
	mBase = m.M
	v1118 = *(*int64)(unsafe.Add(mBase, uint32(v1115)))
	v1119 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1115)+8)))
	m.G0 = v1115 + v1114
	goto L289
L287:
	;
	v1133 = v1024
	goto L288
L288:
	;
	if v1027&int32(1) != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1129 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[17])))
	v1133 = v1119 + v1118*int64(1000000) - int64(946684800000000) + v1129*int64(1000)
	goto L288
L290:
	;
	v1138 = int32(14)
	goto L292
L291:
	;
	v1138 = int32(15)
	goto L292
L292:
	;
	v1140 = F_errstart(m, v1138, int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	if v1140 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1142 = *(*int64)(unsafe.Add(mBase, uint32(v20)+280))
	*(*uint32)(unsafe.Add(mBase, uint32(v20+int32(192)))) = uint32(v1142)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v1144
	v1146 = int64(32)
	v1147 = int64(base.Ui64(v1142) >> (uint(v1146) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+188)) = uint32(v1147)
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+180)) = uint32(v1023)
	v1151 = int64(base.Ui64(v1023) >> (uint(v1146) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+176)) = uint32(v1151)
	F_errmsg(m, int32(_a_F_WalReceiverMain_10), v20+int32(176))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[18]))
	v1168 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[19]))
	v1170 = F_WaitLatch(m, v1165, int32(41), v1168, int32(134217786))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(398), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[18]))
	v1174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1173))) = v1174
	v1179 = base.AtomicRmwOr32(m, v1174, int32(_a_F_WalReceiverMain_11), v1174)
	goto L300
L300:
	;
	if int64(0) < v1133 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1185 = m.G0
	v1186 = int32(16)
	v1187 = v1185 - v1186
	m.G0 = v1187
	F_gettimeofday(m, v1187)
	mBase = m.M
	v1190 = *(*int64)(unsafe.Add(mBase, uint32(v1187)))
	v1191 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1187)+8)))
	m.G0 = v1187 + v1186
	goto L304
L302:
	;
	goto L303
L303:
	;
	v1201 = int32(1)
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[20]))
	if v1203 == int32(0) {
		v1024 = v1133
		v1027 = v1201
		goto L260
	} else {
		goto L306
	}
L304:
	;
	if v1133 <= v1191+v1190*int64(1000000)-int64(946684800000000) {
		goto L281
	} else {
		goto L305
	}
L305:
	;
	goto L303
L306:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1024 = v1133
	v1027 = v1201
	goto L260
L308:
	;
	if v321&v1016 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	if v1034 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1375 = v1034
	goto L311
L311:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+248)) = v1023
	v1377 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+240)) = uint8(v1377)
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v1379
	v1384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+352)))
	if v1384 != 0 {
		goto L353
	} else {
		goto L354
	}
L312:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v1216 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+56))
	v1218 = m.T0[v1217].(func(*base.Module, int32) int32)(m, v1214)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v1245 = base.AtomicRmwXchg32(m, v31, int32(0), int32(1))
	if v1245 != 0 {
		goto L318
	} else {
		goto L319
	}
L315:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+160)) = base.I64_extend_i32_s(v1218)
	v1223 = v20 + int32(352)
	v1228 = F_pg_snprintf(m, v1223, int32(64), int32(_a_F_WalReceiverMain_12), v20+int32(160))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v1233 = int32(0)
	v1238 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+48))
	v1240 = m.T0[v1239].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v1231, v1223, int32(1), v1233, v1233, v1233, v1233)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	goto L314
L318:
	;
	F_s_lock(m, v31, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1250 = v20 + int32(352)
	goto L325
L321:
	;
	goto L320
L322:
	;
	v1370 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v31))), uint32(v1370))
	v1375 = int32(1)
	goto L311
L323:
	;
	v1367 = F_strlen(m, v1356)
	mBase = m.M
	goto L322
L325:
	;
	goto L326
L326:
	;
	v1257 = int32(63)
	if (v200^v1250)&int32(3) != 0 {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	v1360 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1357))) = uint8(v1360)
	goto L323
L328:
	;
	v1341 = v1336
	v1342 = v1337
	v1343 = v1338
	goto L349
L329:
	;
	if v1331 == int32(0) {
		v1356 = v1329
		v1357 = v1330
		goto L327
	} else {
		goto L348
	}
L330:
	;
	v1329 = v1250
	v1330 = v200
	v1331 = v1257
	goto L329
L331:
	;
	goto L332
L332:
	;
	v1261 = int32(0)
	if base.B2i32(v1250&int32(3) == v1261)|int32(0) == v1261 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	if v1297 == int32(0) {
		v1356 = v1294
		v1357 = v1295
		goto L327
	} else {
		goto L342
	}
L334:
	;
	v1273 = v1250
	v1274 = v200
	v1275 = v1257
	goto L337
L335:
	;
	goto L336
L336:
	;
	v1294 = v1250
	v1295 = v200
	v1296 = v1257
	v1297 = int32(1)
	goto L333
L337:
	;
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1273))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1274))) = uint8(v1277)
	if v1277 == int32(0) {
		v1336 = v1273
		v1337 = v1274
		v1338 = v1275
		goto L328
	} else {
		goto L339
	}
L338:
	;
	v1294 = v1288
	v1295 = v1282
	v1296 = v1284
	v1297 = v1286
	goto L333
L339:
	;
	v1281 = int32(1)
	v1282 = v1274 + v1281
	v1284 = v1275 - v1281
	v1285 = int32(0)
	v1286 = base.B2i32(v1284 != v1285)
	v1288 = v1273 + v1281
	if v1288&int32(3) == v1285 {
		v1294 = v1288
		v1295 = v1282
		v1296 = v1284
		v1297 = v1286
		goto L333
	} else {
		goto L340
	}
L340:
	;
	if v1284 != 0 {
		v1273 = v1288
		v1274 = v1282
		v1275 = v1284
		goto L337
	} else {
		goto L341
	}
L341:
	;
	goto L338
L342:
	;
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1294))))
	if base.B2i32(v1300 == int32(0))|base.B2i32(base.Ui32(v1296) < base.Ui32(int32(4))) != 0 {
		v1329 = v1294
		v1330 = v1295
		v1331 = v1296
		goto L329
	} else {
		goto L343
	}
L343:
	;
	v1307 = v1294
	v1308 = v1295
	v1309 = v1296
	goto L344
L344:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1307)))
	v1315 = int32(-2139062144)
	if (int32(16843008)-v1312|v1312)&v1315 != v1315 {
		v1336 = v1307
		v1337 = v1308
		v1338 = v1309
		goto L328
	} else {
		goto L346
	}
L345:
	;
	v1329 = v1323
	v1330 = v1321
	v1331 = v1325
	goto L329
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1308))) = v1312
	v1320 = int32(4)
	v1321 = v1308 + v1320
	v1323 = v1307 + v1320
	v1325 = v1309 - v1320
	if base.Ui32(int32(3)) < base.Ui32(v1325) {
		v1307 = v1323
		v1308 = v1321
		v1309 = v1325
		goto L344
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	v1336 = v1329
	v1337 = v1330
	v1338 = v1331
	goto L328
L349:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1341))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1342))) = uint8(v1345)
	if v1345 == int32(0) {
		v1356 = v1341
		v1357 = v1342
		goto L327
	} else {
		goto L351
	}
L350:
	;
	v1356 = v1352
	v1357 = v1350
	goto L327
L351:
	;
	v1349 = int32(1)
	v1350 = v1342 + v1349
	v1352 = v1341 + v1349
	v1354 = v1343 - v1349
	if v1354 != 0 {
		v1341 = v1352
		v1342 = v1350
		v1343 = v1354
		goto L349
	} else {
		goto L352
	}
L352:
	;
	goto L350
L353:
	;
	v1385 = v20 + int32(352)
	goto L355
L354:
	;
	v1385 = v1377
	goto L355
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v1385
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v1392 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+32))
	v1394 = m.T0[v1393].(func(*base.Module, int32, int32) int32)(m, v1388, v20+int32(240))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	v1398 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	if v1394 == int32(0) {
		goto L266
	} else {
		goto L358
	}
L358:
	;
	if v1398 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v1402
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+132)) = uint32(v1023)
	v1406 = int64(base.Ui64(v1023) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+128)) = uint32(v1406)
	if v1031 != 0 {
		goto L362
	} else {
		goto L363
	}
L360:
	;
	goto L361
L361:
	;
	v1424 = base.AtomicRmwXchg32(m, v31, int32(0), int32(1))
	if v1424 != 0 {
		goto L370
	} else {
		goto L371
	}
L362:
	;
	v1410 = int32(_a_F_WalReceiverMain_13)
	goto L364
L363:
	;
	v1410 = int32(_a_F_WalReceiverMain_14)
	goto L364
L364:
	;
	F_errmsg(m, v1410, v20+int32(128))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	if v1031 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1418 = int32(481)
	goto L368
L367:
	;
	v1418 = int32(485)
	goto L368
L368:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), v1418, int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	goto L361
L370:
	;
	F_s_lock(m, v31, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L1
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v1428 == int32(2) {
		goto L374
	} else {
		goto L375
	}
L373:
	;
	goto L372
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(3)
	goto L376
L375:
	;
	goto L376
L376:
	;
	v1433 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v31))), uint32(v1433))
	v1438 = F_GetXLogReplayRecPtr(m, v1433)
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L1
	} else {
		goto L377
	}
L377:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = v1438
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[22])) = v1438
	v1448 = m.G0
	v1449 = int32(16)
	v1450 = v1448 - v1449
	m.G0 = v1450
	F_gettimeofday(m, v1450)
	mBase = m.M
	v1453 = *(*int64)(unsafe.Add(mBase, uint32(v1450)))
	v1454 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1450)+8)))
	m.G0 = v1450 + v1449
	v1462 = v1454 + v1453*int64(1000000) - int64(946684800000000)
	goto L378
L378:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[17]))
	v1470 = base.B2i32(v1464 <= int32(0))
	if v1464 <= int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1471 = int64(9223372036854775807)
	goto L381
L380:
	;
	v1471 = v1462 + base.I64_extend_i32_u(v1464)*int64(1000)
	goto L381
L381:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23])) = v1471
	if v1464 <= int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1481 = int64(9223372036854775807)
	goto L384
L383:
	;
	v1481 = v1462 + base.I64_extend_i32_u(int32(base.Ui32(v1464)>>(uint(int32(1))%32)))*int64(1000)
	goto L384
L384:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])) = v1481
	v1486 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[25]))
	v1490 = v1462 + base.I64_extend_i32_u(v1486)*int64(1000000)
	if v1486 <= int32(0) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1493 = int64(9223372036854775807)
	goto L387
L386:
	;
	v1493 = v1490
	goto L387
L387:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])) = v1493
	if v1486 <= int32(0) {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v1499 = int64(9223372036854775807)
	goto L390
L389:
	;
	v1499 = v1490
	goto L390
L390:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[27])))
	if v1502 != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v1503 = v1499
	goto L393
L392:
	;
	v1503 = int64(9223372036854775807)
	goto L393
L393:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[28])) = v1503
	v1506 = int32(0)
	F_XLogWalRcvSendReply(m, int32(1), v1506, v1506)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = int32(-1)
	v1517 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[29])))
	if v1517 == int32(1) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L1
	} else {
		goto L591
	}
L397:
	;
	if v1527 != 0 {
		goto L401
	} else {
		goto L402
	}
L398:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30]))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+308))
	v1525 = base.B2i32(v1523 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[29])) = uint8(v1525)
	v1527 = v1525
	goto L400
L399:
	;
	v1527 = int32(0)
	goto L400
L400:
	;
	goto L397
L401:
	;
	goto L404
L402:
	;
	goto L403
L403:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L1
	} else {
		goto L587
	}
L404:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[20]))
	if v1546 != 0 {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	goto L403
L406:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L1
	} else {
		goto L409
	}
L407:
	;
	goto L408
L408:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[31]))
	if v1550 != 0 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L408
L410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[31])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v1632 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1632)+40))
	v1634 = m.T0[v1633].(func(*base.Module, int32, int32, int32) int32)(m, v1626, v20+int32(236), v20+int32(232))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L1
	} else {
		goto L431
	}
L413:
	;
	v1562 = m.G0
	v1563 = int32(16)
	v1564 = v1562 - v1563
	m.G0 = v1564
	F_gettimeofday(m, v1564)
	mBase = m.M
	v1567 = *(*int64)(unsafe.Add(mBase, uint32(v1564)))
	v1568 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1564)+8)))
	m.G0 = v1564 + v1563
	v1576 = v1568 + v1567*int64(1000000) - int64(946684800000000)
	goto L414
L414:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[17]))
	v1584 = base.B2i32(v1578 <= int32(0))
	if v1578 <= int32(0) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1585 = int64(9223372036854775807)
	goto L417
L416:
	;
	v1585 = v1576 + base.I64_extend_i32_u(v1578)*int64(1000)
	goto L417
L417:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23])) = v1585
	if v1578 <= int32(0) {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v1595 = int64(9223372036854775807)
	goto L420
L419:
	;
	v1595 = v1576 + base.I64_extend_i32_u(int32(base.Ui32(v1578)>>(uint(int32(1))%32)))*int64(1000)
	goto L420
L420:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])) = v1595
	v1600 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[25]))
	v1604 = v1576 + base.I64_extend_i32_u(v1600)*int64(1000000)
	if v1600 <= int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1607 = int64(9223372036854775807)
	goto L423
L422:
	;
	v1607 = v1604
	goto L423
L423:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[26])) = v1607
	if v1600 <= int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1613 = int64(9223372036854775807)
	goto L426
L425:
	;
	v1613 = v1604
	goto L426
L426:
	;
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[27])))
	if v1616 != 0 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1617 = v1613
	goto L429
L428:
	;
	v1617 = int64(9223372036854775807)
	goto L429
L429:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[28])) = v1617
	F_XLogWalRcvSendHSFeedback(m, int32(1))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L1
	} else {
		goto L430
	}
L430:
	;
	goto L412
L431:
	;
	if v1634 != 0 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	if int32(0) < v1634 {
		goto L436
	} else {
		goto L437
	}
L433:
	;
	goto L434
L434:
	;
	v2233 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23]))
	v2235 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[24]))
	v2237 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[26]))
	v2239 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[28]))
	v2243 = m.G0
	v2244 = int32(16)
	v2245 = v2243 - v2244
	m.G0 = v2245
	F_gettimeofday(m, v2245)
	mBase = m.M
	v2248 = *(*int64)(unsafe.Add(mBase, uint32(v2245)))
	v2249 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2245)+8)))
	m.G0 = v2245 + v2244
	v2257 = v2249 + v2248*int64(1000000) - int64(946684800000000)
	goto L547
L435:
	;
	v2206 = int32(0)
	F_XLogWalRcvSendReply(m, v2206, v2206, v2206)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L1
	} else {
		goto L545
	}
L436:
	;
	v1639 = v1634
	goto L439
L437:
	;
	goto L438
L438:
	;
	v2158 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L534
	}
L439:
	;
	v1658 = m.G0
	v1659 = int32(16)
	v1660 = v1658 - v1659
	m.G0 = v1660
	F_gettimeofday(m, v1660)
	mBase = m.M
	v1663 = *(*int64)(unsafe.Add(mBase, uint32(v1660)))
	v1664 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1660)+8)))
	m.G0 = v1660 + v1659
	v1672 = v1664 + v1663*int64(1000000) - int64(946684800000000)
	goto L441
L440:
	;
	if v2133 == int32(0) {
		goto L435
	} else {
		goto L533
	}
L441:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[17]))
	if v1675 <= int32(0) {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])) = v1691
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23])) = v1690
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v20)+236))
	v1697 = v1695 + int32(1)
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1695))))
	v1700 = v1698 - int32(107)
	if v1700 != 0 {
		goto L452
	} else {
		goto L453
	}
L443:
	;
	v1678 = int64(9223372036854775807)
	v1690 = v1678
	v1691 = v1678
	goto L442
L444:
	;
	goto L445
L445:
	;
	v1681 = int64(1000)
	v1690 = base.I64_extend_i32_u(v1675)*v1681 + v1672
	v1691 = base.I64_extend_i32_u(int32(base.Ui32(v1675)>>(uint(int32(1))%32)))*v1681 + v1672
	goto L442
L446:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v2131 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2131)+40))
	v2133 = m.T0[v2132].(func(*base.Module, int32, int32, int32) int32)(m, v2125, v20+int32(236), v20+int32(232))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L1
	} else {
		goto L531
	}
L447:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[0]))
	v2089 = base.AtomicRmwXchg64(m, v2087, int32(1464), v2071)
	v2092 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21]))
	F_WaitLSNWakeup(m, int32(1), v2092)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L1
	} else {
		goto L527
	}
L448:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L1
	} else {
		goto L523
	}
L449:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L1
	} else {
		goto L519
	}
L450:
	;
	v2034 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21]))
	v2071 = v2034
	v2072 = v1713
	goto L447
L451:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L1
	} else {
		goto L515
	}
L452:
	;
	if v1700 != int32(12) {
		goto L448
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	if v1639 != int32(18) {
		goto L449
	} else {
		goto L508
	}
L455:
	;
	if base.Ui32(v1639) <= base.Ui32(int32(24)) {
		goto L451
	} else {
		goto L456
	}
L456:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1452)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+1444)) = int64(24)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1440)) = v1697
	v1712 = v20 + int32(1440)
	v1713 = F_pq_getmsgint64(m, v1712)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	v1715 = F_pq_getmsgint64(m, v1712)
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v1717 = F_pq_getmsgint64(m, v1712)
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L1
	} else {
		goto L459
	}
L459:
	;
	F_ProcessWalSndrMessage(m, v1715, v1717)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	v1722 = v1639 - int32(25)
	if v1722 == int32(0) {
		goto L450
	} else {
		goto L461
	}
L461:
	;
	v1729 = v1713
	v1733 = v1695 + int32(25)
	v1735 = v1722
	goto L462
L462:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	v1747 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	if int32(0) <= v1747 {
		goto L465
	} else {
		goto L466
	}
L463:
	;
	v2071 = v1989
	v2072 = v1989
	goto L447
L464:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[32])) = int32(0)
	v1782 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[33])))
	v1785 = m.G0
	v1787 = v1785 - int32(16)
	m.G0 = v1787
	if v1782 != 0 {
		goto L473
	} else {
		goto L474
	}
L465:
	;
	v1751 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[34]))
	v1753 = base.I64_div_u_s(v1729, base.I64_extend_i32_s(v1745))
	if v1751 == v1753 {
		v1776 = v1745
		goto L464
	} else {
		goto L468
	}
L466:
	;
	v1763 = v1745
	goto L467
L467:
	;
	v1766 = base.I64_div_u_s(v1729, base.I64_extend_i32_s(v1763))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[34])) = v1766
	v1768 = F_XLogFileInit(m, v1766, v1705)
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L1
	} else {
		goto L471
	}
L468:
	;
	F_XLogWalRcvClose(m, v1705)
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	v1760 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	if int32(0) <= v1760 {
		v1776 = v1758
		goto L464
	} else {
		goto L470
	}
L470:
	;
	v1763 = v1758
	goto L467
L471:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[35])) = v1705
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15])) = v1768
	v1775 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	v1776 = v1775
	goto L464
L472:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v1801))) = int32(167772242)
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	v1809 = base.I32_wrap_i64(v1729) & (v1776 - int32(1))
	if base.Ui32(v1776) < base.Ui32(v1735+v1809) {
		goto L476
	} else {
		goto L477
	}
L473:
	;
	F___clock_gettime(m, int32(1), v1787)
	mBase = m.M
	v1791 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1787)+8)))
	v1792 = *(*int64)(unsafe.Add(mBase, uint32(v1787)))
	v1796 = v1791 + v1792*int64(1000000000)
	goto L475
L474:
	;
	v1796 = int64(0)
	goto L475
L475:
	;
	m.G0 = v1787 + int32(16)
	goto L472
L476:
	;
	v1813 = v1776 - v1809
	goto L478
L477:
	;
	v1813 = v1735
	goto L478
L478:
	;
	v1815 = F_pwrite(m, v1805, v1733, v1813, base.I64_extend_i32_s(v1809))
	mBase = m.M
	v1817 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[36]))
	v1818 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1817))) = v1818
	if v1815 <= v1818 {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[32]))
	if v1823 == int32(0) {
		goto L482
	} else {
		goto L483
	}
L480:
	;
	goto L481
L481:
	;
	v1865 = int32(1)
	v1866 = base.I64_extend_i32_u(v1815)
	v1870 = m.G0
	v1872 = v1870 - int32(16)
	m.G0 = v1872
	if v1796 != int64(0) {
		goto L491
	} else {
		goto L492
	}
L482:
	;
	v1827 = int32(51)
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[32])) = v1827
	v1830 = v1827
	goto L484
L483:
	;
	v1830 = v1823
	goto L484
L484:
	;
	v1832 = v20 + int32(1456)
	v1834 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[35]))
	v1836 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[34]))
	v1838 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16]))
	F_XLogFileName(m, v1832, v1834, v1836, v1838)
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L1
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[32])) = v1830
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+120)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v1809
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v1832
	F_errmsg(m, int32(_a_F_WalReceiverMain_15), v20+int32(112))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(1052), int32(_a_F_WalReceiverMain_16))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
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
	v1989 = v1729 + v1866
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21])) = v1989
	v1992 = v1735 - v1815
	if v1992 != 0 {
		v1729 = v1989
		v1733 = v1815 + v1733
		v1735 = v1992
		goto L462
	} else {
		goto L507
	}
L491:
	;
	F___clock_gettime(m, int32(1), v1872)
	mBase = m.M
	v1878 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1872)+8)))
	v1879 = *(*int64)(unsafe.Add(mBase, uint32(v1872)))
	v1883 = v1878 + (v1879*int64(1000000000) - v1796)
	goto L494
L492:
	;
	goto L493
L493:
	;
	v1970 = int32(888)
	v1971 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[37]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[37])) = v1971 + base.I64_extend_i32_u(v1865)
	v1975 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[38])) = v1975 + v1866
	F_pgstat_count_backend_io_op(m, int32(2), int32(3), int32(7), v1865, v1866)
	mBase = m.M
	v1980 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[39])) = uint8(v1980)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[40])) = uint8(v1980)
	m.G0 = v1872 + int32(16)
	goto L490
L494:
	;
	v1933 = int32(888)
	v1934 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[41])) = v1934 + v1883
	v1938 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[42]))
	v1945 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1938))|base.B2i32(int32(1)<<(uint(v1938)%32)&int32(_a_F_WalReceiverMain_17) == v1945) == v1945 {
		goto L504
	} else {
		goto L505
	}
L504:
	;
	v1950 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[43]))
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[43])) = v1950 + v1883
	v1954 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[39])) = uint8(v1954)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[44])) = uint8(v1954)
	goto L506
L505:
	;
	goto L506
L506:
	;
	goto L493
L507:
	;
	goto L463
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1468)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+1460)) = int64(17)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+1456)) = v1697
	v2001 = v20 + int32(1456)
	v2002 = F_pq_getmsgint64(m, v2001)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	v2004 = F_pq_getmsgint64(m, v2001)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	v2006 = F_pq_getmsgbyte(m, v2001)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	F_ProcessWalSndrMessage(m, v2002, v2004)
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L1
	} else {
		goto L512
	}
L512:
	;
	if v2006 == int32(0) {
		goto L446
	} else {
		goto L513
	}
L513:
	;
	v2013 = int32(0)
	F_XLogWalRcvSendReply(m, int32(1), v2013, v2013)
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L1
	} else {
		goto L514
	}
L514:
	;
	goto L446
L515:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L1
	} else {
		goto L516
	}
L516:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_18), int32(0))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(939), int32(_a_F_WalReceiverMain_19))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L518
	}
L518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L519:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_20), int32(0))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(963), int32(_a_F_WalReceiverMain_19))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L1
	} else {
		goto L522
	}
L522:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L523:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L1
	} else {
		goto L524
	}
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v1698
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_21), v20+int32(32))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(984), int32(_a_F_WalReceiverMain_19))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L527:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	if v2096 < int32(0) {
		goto L446
	} else {
		goto L528
	}
L528:
	;
	v2100 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[34]))
	v2102 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16])))
	v2103 = base.I64_div_u_s(v2072, v2102)
	if v2100 == v2103 {
		goto L446
	} else {
		goto L529
	}
L529:
	;
	F_XLogWalRcvClose(m, v1705)
	mBase = m.M
	v2106 = m.ExcPending
	if v2106 != 0 {
		goto L1
	} else {
		goto L530
	}
L530:
	;
	goto L446
L531:
	;
	if int32(0) < v2133 {
		v1639 = v2133
		goto L439
	} else {
		goto L532
	}
L532:
	;
	goto L440
L533:
	;
	goto L438
L534:
	;
	if v2158 != 0 {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_22), int32(0))
	mBase = m.M
	v2163 = m.ExcPending
	if v2163 != 0 {
		goto L1
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	v2183 = int32(0)
	F_XLogWalRcvSendReply(m, v2183, v2183, v2183)
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L1
	} else {
		goto L541
	}
L538:
	;
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v2164
	v2167 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[21]))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+104)) = uint32(v2167)
	v2170 = int64(base.Ui64(v2167) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+100)) = uint32(v2170)
	v2175 = F_errdetail(m, int32(_a_F_WalReceiverMain_23), v20+int32(96))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(576), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	goto L537
L541:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	F_XLogWalRcvFlush(m, int32(0), v2190)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[13]))
	v2198 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[11]))
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v2198)+36))
	m.T0[v2199].(func(*base.Module, int32, int32))(m, v2194, v20+int32(344))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L1
	} else {
		goto L543
	}
L543:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v20)+344))
	F_WalRcvFetchTimeLineHistoryFiles(m, v2202, v2203)
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	v2510 = v2183
	goto L265
L545:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	F_XLogWalRcvFlush(m, int32(0), v2212)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L1
	} else {
		goto L546
	}
L546:
	;
	goto L434
L547:
	;
	if v2235 < v2233 {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v2259 = v2235
	goto L550
L549:
	;
	v2259 = v2233
	goto L550
L550:
	;
	if v2237 < v2259 {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v2261 = v2237
	goto L553
L552:
	;
	v2261 = v2259
	goto L553
L553:
	;
	if v2239 < v2261 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v2263 = v2239
	goto L556
L555:
	;
	v2263 = v2261
	goto L556
L556:
	;
	if v2263 <= v2257 {
		v2281 = int32(0)
		goto L558
	} else {
		goto L559
	}
L557:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[18]))
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v20)+232))
	v2286 = F_WaitLatchOrSocket(m, v2283, v2284, v2281, int32(83886094))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L1
	} else {
		goto L562
	}
L558:
	;
	goto L557
L559:
	;
	v2269 = v2263 - v2257
	if base.B2i32(int64(0) < v2257)^base.B2i32(v2269 < v2263)|base.B2i32(int64(2147483646000) < v2269) != 0 {
		v2281 = int32(2147483647)
		goto L558
	} else {
		goto L560
	}
L560:
	;
	v2278 = base.I64_div_s(v2269+int64(999), int64(1000))
	v2281 = base.I32_wrap_i64(v2278)
	goto L558
L561:
	;
	if v2286&int32(8) != 0 {
		goto L571
	} else {
		goto L572
	}
L562:
	;
	if v2286&int32(1) == int32(0) {
		goto L561
	} else {
		goto L563
	}
L563:
	;
	v2293 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[18]))
	v2294 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2293))) = v2294
	v2299 = base.AtomicRmwOr32(m, v2294, int32(_a_F_WalReceiverMain_11), v2294)
	goto L564
L564:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[20]))
	if v2301 != 0 {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L1
	} else {
		goto L568
	}
L566:
	;
	goto L567
L567:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v29)+1472))
	if v2304 == int32(0) {
		goto L561
	} else {
		goto L569
	}
L568:
	;
	goto L567
L569:
	;
	v2307 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+1472)) = v2307
	v2312 = base.AtomicRmwOr32(m, v2307, int32(_a_F_WalReceiverMain_24), v2307)
	F_XLogWalRcvSendReply(m, v2307, v2307, int32(1))
	mBase = m.M
	v2317 = m.ExcPending
	if v2317 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	goto L561
L571:
	;
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L1
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+232)) = int32(-1)
	v2364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[29])))
	if v2364 == int32(1) {
		goto L583
	} else {
		goto L584
	}
L574:
	;
	v2326 = m.G0
	v2327 = int32(16)
	v2328 = v2326 - v2327
	m.G0 = v2328
	F_gettimeofday(m, v2328)
	mBase = m.M
	v2331 = *(*int64)(unsafe.Add(mBase, uint32(v2328)))
	v2332 = int64(*(*int32)(unsafe.Add(mBase, uint32(v2328)+8)))
	m.G0 = v2328 + v2327
	v2340 = v2332 + v2331*int64(1000000) - int64(946684800000000)
	goto L575
L575:
	;
	v2342 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[23]))
	if v2342 <= v2340 {
		goto L396
	} else {
		goto L576
	}
L576:
	;
	v2345 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[24]))
	if v2345 <= v2340 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[24])) = int64(9223372036854775807)
	goto L579
L578:
	;
	goto L579
L579:
	;
	v2350 = base.B2i32(v2345 <= v2340)
	F_XLogWalRcvSendReply(m, v2350, v2350, int32(0))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L1
	} else {
		goto L580
	}
L580:
	;
	F_XLogWalRcvSendHSFeedback(m, int32(0))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	goto L573
L582:
	;
	if v2374 != 0 {
		goto L404
	} else {
		goto L586
	}
L583:
	;
	v2369 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[30]))
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v2369)+308))
	v2372 = base.B2i32(v2370 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[29])) = uint8(v2372)
	v2374 = v2372
	goto L585
L584:
	;
	v2374 = int32(0)
	goto L585
L585:
	;
	goto L582
L586:
	;
	goto L405
L587:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_25), int32(0))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(529), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L1
	} else {
		goto L590
	}
L590:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L591:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_26), int32(0))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(674), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L595:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_27), int32(0))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(411), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L1
	} else {
		goto L598
	}
L598:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L599:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v20)+344))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v2447
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v2449
	F_errmsg(m, int32(_a_F_WalReceiverMain_28), v20+int32(16))
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(355), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L603:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L1
	} else {
		goto L604
	}
L604:
	;
	F_errmsg(m, int32(_a_F_WalReceiverMain_29), int32(0))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L605
	}
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v1047
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v20 + int32(288)
	v2479 = F_errdetail(m, int32(_a_F_WalReceiverMain_30), v20+int32(208))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L1
	} else {
		goto L606
	}
L606:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(343), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L608:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v2488
	F_errmsg(m, int32(_a_F_WalReceiverMain_31), v20+int32(144))
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L1
	} else {
		goto L609
	}
L609:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(707), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	v2510 = v1031
	goto L265
L611:
	;
	goto L261
L612:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15])) = int32(-1)
	v2569 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2570 = m.ExcPending
	if v2570 != 0 {
		goto L1
	} else {
		goto L622
	}
L613:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v20)+348))
	F_XLogWalRcvFlush(m, int32(0), v2522)
	mBase = m.M
	v2524 = m.ExcPending
	if v2524 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v2526
	v2529 = *(*int64)(unsafe.Add(mBase, _c_F_WalReceiverMain[34]))
	v2532 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[16])))
	v2533 = base.I64_div_u_s(int64(4294967296), v2532)
	v2534 = base.I64_div_u_s(v2529, v2533)
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+84)) = uint32(v2534)
	v2537 = v2529 - v2533*v2534
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+88)) = uint32(v2537)
	v2540 = v20 + int32(1456)
	v2545 = F_pg_snprintf(m, v2540, int32(64), int32(_a_F_WalReceiverMain_32), v20+int32(80))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	v2548 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[15]))
	v2549 = F_close(m, v2548)
	mBase = m.M
	if v2549 != 0 {
		goto L611
	} else {
		goto L616
	}
L616:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[45]))
	if v2551 != int32(2) {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	F_XLogArchiveForceDone(m, v2540)
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L1
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	F_XLogArchiveNotify(m, v20+int32(1456))
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L1
	} else {
		goto L621
	}
L620:
	;
	goto L612
L621:
	;
	goto L612
L622:
	;
	if v2569 != 0 {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_33), int32(0))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L1
	} else {
		goto L626
	}
L624:
	;
	goto L625
L625:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[0]))
	v2582 = int32(1456)
	v2583 = v2581 + v2582
	v2586 = base.AtomicRmwXchg32(m, v2581, v2582, int32(1))
	if v2586 != 0 {
		goto L628
	} else {
		goto L629
	}
L626:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(736), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	goto L625
L628:
	;
	F_s_lock(m, v2583, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L1
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2581)+8))
	if base.Ui32(v2590-int32(4)) <= base.Ui32(int32(-3)) {
		goto L632
	} else {
		goto L633
	}
L631:
	;
	goto L630
L632:
	;
	v2595 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2583))), uint32(v2595))
	if v2590 == int32(6) {
		goto L259
	} else {
		goto L635
	}
L633:
	;
	goto L634
L634:
	;
	v2613 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+40)) = v2613
	*(*int64)(unsafe.Add(mBase, uint32(v2581)+32)) = int64(0)
	v2617 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+8)) = v2617
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2581)+1456)), uint32(v2613))
	v2623 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[46]))
	F_SetLatch(m, v2623+v2617)
	mBase = m.M
	goto L639
L635:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v2603 = m.ExcPending
	if v2603 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_34), int32(0))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(759), int32(_a_F_WalReceiverMain_35))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L639:
	;
	goto L640
L640:
	;
	v2645 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[18]))
	v2646 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2645))) = v2646
	v2651 = base.AtomicRmwOr32(m, v2646, int32(_a_F_WalReceiverMain_11), v2646)
	goto L642
L642:
	;
	v2653 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[20]))
	if v2653 != 0 {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2655 = m.ExcPending
	if v2655 != 0 {
		goto L1
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	v2658 = base.AtomicRmwXchg32(m, v2583, int32(0), int32(1))
	if v2658 != 0 {
		goto L647
	} else {
		goto L648
	}
L646:
	;
	goto L645
L647:
	;
	F_s_lock(m, v2583, int32(_a_F_WalReceiverMain_0))
	mBase = m.M
	v2661 = m.ExcPending
	if v2661 != 0 {
		goto L1
	} else {
		goto L650
	}
L648:
	;
	goto L649
L649:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2581)+8))
	switch v2662 - int32(5) {
	case 0:
		goto L653
	case 1:
		goto L652
	default:
		goto L651
	}
L650:
	;
	goto L649
L651:
	;
	v2698 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2583))), uint32(v2698))
	v2702 = *(*int32)(unsafe.Add(mBase, _c_F_WalReceiverMain[18]))
	v2706 = F_WaitLatch(m, v2702, int32(33), v2698, int32(134217787))
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L1
	} else {
		goto L657
	}
L652:
	;
	v2692 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2583))), uint32(v2692))
	F_proc_exit(m, int32(1))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L1
	} else {
		goto L656
	}
L653:
	;
	v2665 = *(*int64)(unsafe.Add(mBase, uint32(v2581)+32))
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v2581)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+348)) = v2666
	*(*int32)(unsafe.Add(mBase, uint32(v2581)+8)) = int32(2)
	v2670 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v2581)+1456)), uint32(v2670))
	v2674 = int64(0)
	v2676 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalReceiverMain[47])))
	if v2676 == v2670 {
		v1023 = v2665
		v1024 = v2674
		v1027 = v2670
		v1031 = v2510
		v1034 = v1375
		goto L260
	} else {
		goto L654
	}
L654:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+52)) = uint32(v2665)
	v2681 = int64(base.Ui64(v2665) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+48)) = uint32(v2681)
	v2684 = v20 + int32(1456)
	v2689 = F_pg_snprintf(m, v2684, int32(50), int32(_a_F_WalReceiverMain_36), v20+int32(48))
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	v2691 = F_strlen(m, v2684)
	mBase = m.M
	v1023 = v2665
	v1024 = v2674
	v1027 = v2670
	v1031 = v2510
	v1034 = v1375
	goto L260
L656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L657:
	;
	goto L640
L658:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v20 + int32(1456)
	F_errmsg(m, int32(_a_F_WalReceiverMain_37), v20-int32(-64))
	mBase = m.M
	v2721 = m.ExcPending
	if v2721 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(723), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L1
	} else {
		goto L661
	}
L661:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L663:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v729
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v20)+340))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v2738
	F_errmsg(m, int32(_a_F_WalReceiverMain_38), v20)
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L665
	}
L665:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(295), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L667:
	;
	F_errmsg_internal(m, int32(_a_F_WalReceiverMain_39), int32(0))
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	F_errfinish(m, int32(_a_F_WalReceiverMain_2), int32(269), int32(_a_F_WalReceiverMain_3))
	mBase = m.M
	v2760 = m.ExcPending
	if v2760 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
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
									F_errfinish(m, int32(_a_F_WalSndSegmentOpen_2), int32(3349), int32(_a_F_WalSndSegmentOpen_3))
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
								F_errfinish(m, int32(_a_F_WalSndSegmentOpen_2), int32(3355), int32(_a_F_WalSndSegmentOpen_3))
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
func F_WalSummarizerShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_WalSummarizerShmemRequest_0), int64(48), int32(_a_F_WalSummarizerShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_WalWriterMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int64
	_ = v546
	var v548 int64
	_ = v548
	var v550 int32
	_ = v550
	var v553 int64
	_ = v553
	var v557 int32
	_ = v557
	var v558 int64
	_ = v558
	var v561 int64
	_ = v561
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int64
	_ = v573
	var v575 int64
	_ = v575
	var v577 int64
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int64
	_ = v592
	var v598 int64
	_ = v598
	var v601 int32
	_ = v601
	var v605 int64
	_ = v605
	var v607 int64
	_ = v607
	var v611 int64
	_ = v611
	var v612 int64
	_ = v612
	var v615 int32
	_ = v615
	var v618 int64
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int64
	_ = v627
	var v628 int64
	_ = v628
	var v636 int64
	_ = v636
	var v638 int32
	_ = v638
	var v640 int64
	_ = v640
	var v648 int64
	_ = v648
	var v650 int32
	_ = v650
	var v660 int32
	_ = v660
	var v661 int64
	_ = v661
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int64
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int64
	_ = v693
	var v696 int64
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int64
	_ = v708
	var v710 int64
	_ = v710
	var v712 int64
	_ = v712
	var v714 int64
	_ = v714
	var v716 int64
	_ = v716
	var v718 int64
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int64
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v823 int32
	_ = v823
	var v824 int64
	_ = v824
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(176)
	m.G0 = v16
	v19 = int32(-1)
	v22 = v3
	v24 = v3
	goto L1
L1:
	;
	goto L3
L3:
	;
	if v19 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v823 = int32(m.ExcTag)
	v824 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v823 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v400 = v22
	v401 = v24
	goto L8
L8:
	;
	if v401 != 0 {
		goto L95
	} else {
		goto L96
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v42 = m.G0
	v44 = v42 - int32(32)
	m.G0 = v44
	v47 = int32(967)
	switch v47 {
	case 0, 2:
		goto L11
	default:
		goto L12
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v83 = int32(0)
	v85 = m.G0
	v87 = v85 - int32(32)
	m.G0 = v87
	switch v83 {
	case 0, 2:
		goto L21
	default:
		goto L22
	}
L11:
	;
	F_sigemptyset(m, v44+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = int32(268435456)
	switch v47 {
	case 0:
		goto L16
	default:
		goto L14
	case 2:
		goto L15
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[0])) = int32(965)
	goto L11
L13:
	;
	goto L18
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(_a_F_WalWriterMain_0)
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(0)
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = int32(-2)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v76 = F___sigaction(m, int32(1), v44+int32(12), int32(0))
	mBase = m.M
	m.G0 = v44 + int32(32)
	goto L10
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v128 = m.G0
	v130 = v128 - int32(32)
	m.G0 = v130
	v133 = int32(969)
	switch v133 {
	case 0, 2:
		goto L31
	default:
		goto L32
	}
L21:
	;
	F_sigemptyset(m, v87+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = int32(268435456)
	switch v83 {
	case 0:
		goto L26
	default:
		goto L24
	case 2:
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[1])) = int32(-2)
	goto L21
L23:
	;
	goto L28
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(_a_F_WalWriterMain_0)
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(0)
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(-2)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v119 = F___sigaction(m, int32(2), v87+int32(12), int32(0))
	mBase = m.M
	m.G0 = v87 + int32(32)
	goto L20
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v169 = int32(0)
	v171 = m.G0
	v173 = v171 - int32(32)
	m.G0 = v173
	switch v169 {
	case 0, 2:
		goto L41
	default:
		goto L42
	}
L31:
	;
	F_sigemptyset(m, v130+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v130)+24)) = int32(268435456)
	switch v133 {
	case 0:
		goto L36
	default:
		goto L34
	case 2:
		goto L35
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[2])) = int32(967)
	goto L31
L33:
	;
	goto L38
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(_a_F_WalWriterMain_0)
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(0)
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(-2)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v162 = F___sigaction(m, int32(15), v130+int32(12), int32(0))
	mBase = m.M
	m.G0 = v130 + int32(32)
	goto L30
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v212 = int32(0)
	v214 = m.G0
	v216 = v214 - int32(32)
	m.G0 = v216
	switch v212 {
	case 0, 2:
		goto L51
	default:
		goto L52
	}
L41:
	;
	F_sigemptyset(m, v173+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = int32(268435456)
	switch v169 {
	case 0:
		goto L46
	default:
		goto L44
	case 2:
		goto L45
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[3])) = int32(-2)
	goto L41
L43:
	;
	goto L48
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = int32(_a_F_WalWriterMain_0)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = int32(0)
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = int32(-2)
	goto L43
L48:
	;
	goto L49
L49:
	;
	v205 = F___sigaction(m, int32(14), v173+int32(12), int32(0))
	mBase = m.M
	m.G0 = v173 + int32(32)
	goto L40
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v257 = m.G0
	v259 = v257 - int32(32)
	m.G0 = v259
	v262 = int32(970)
	switch v262 {
	case 0, 2:
		goto L61
	default:
		goto L62
	}
L51:
	;
	F_sigemptyset(m, v216+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(268435456)
	switch v212 {
	case 0:
		goto L56
	default:
		goto L54
	case 2:
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[4])) = int32(-2)
	goto L51
L53:
	;
	goto L58
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = int32(_a_F_WalWriterMain_0)
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = int32(0)
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = int32(-2)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v248 = F___sigaction(m, int32(13), v216+int32(12), int32(0))
	mBase = m.M
	m.G0 = v216 + int32(32)
	goto L50
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v298 = int32(0)
	v300 = m.G0
	v302 = v300 - int32(32)
	m.G0 = v302
	switch v298 {
	case 0, 2:
		goto L71
	default:
		goto L72
	}
L61:
	;
	F_sigemptyset(m, v259+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v259)+24)) = int32(268435456)
	switch v262 {
	case 0:
		goto L66
	default:
		goto L64
	case 2:
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[5])) = int32(968)
	goto L61
L63:
	;
	goto L68
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+12)) = int32(_a_F_WalWriterMain_0)
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+12)) = int32(0)
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+12)) = int32(-2)
	goto L63
L68:
	;
	goto L69
L69:
	;
	v291 = F___sigaction(m, int32(10), v259+int32(12), int32(0))
	mBase = m.M
	m.G0 = v259 + int32(32)
	goto L60
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v343 = m.G0
	v345 = v343 - int32(32)
	m.G0 = v345
	v347 = int32(2)
	switch v347 {
	case 0, 2:
		goto L81
	default:
		goto L82
	}
L71:
	;
	F_sigemptyset(m, v302+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v302)+24)) = int32(268435456)
	switch v298 {
	case 0:
		goto L76
	default:
		goto L74
	case 2:
		goto L75
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[6])) = int32(-2)
	goto L71
L73:
	;
	goto L78
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v302)+12)) = int32(_a_F_WalWriterMain_0)
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+12)) = int32(0)
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v302)+12)) = int32(-2)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v334 = F___sigaction(m, int32(12), v302+int32(12), int32(0))
	mBase = m.M
	m.G0 = v302 + int32(32)
	goto L70
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v22
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[7]))
	v388 = F_AllocSetContextCreateInternal(m, v383, int32(_a_F_WalWriterMain_1), int32(0), int32(_a_F_WalWriterMain_2), int32(_a_F_WalWriterMain_3))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L90
	}
L81:
	;
	F_sigemptyset(m, v345+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v345)+24)) = int32(268435456)
	switch v347 {
	case 0:
		goto L86
	default:
		goto L84
	case 2:
		goto L85
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[8])) = int32(0)
	goto L81
L83:
	;
	goto L87
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = int32(_a_F_WalWriterMain_0)
	v370 = int32(268435461)
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = int32(0)
	v370 = int32(268435457)
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+12)) = int32(-2)
	v370 = int32(268435457)
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v345)+24)) = v370
	goto L89
L89:
	;
	v377 = F___sigaction(m, int32(17), v345+int32(12), int32(0))
	mBase = m.M
	m.G0 = v345 + int32(32)
	goto L80
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[9])) = v388
	goto L91
L91:
	;
	v394 = v16 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v16 + int32(12)
	goto L94
L92:
	;
	v400 = v388
	v401 = int32(0)
	goto L8
L94:
	;
	goto L92
L95:
	;
	v402 = int32(_a_F_WalWriterMain_4)
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10])) = v404 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[11])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_EmitErrorReport(m)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[12])) = v16 + int32(16)
	F_pgmem_sigprocmask(m, int32(_a_F_WalWriterMain_5), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L5
	} else {
		goto L109
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_LWLockReleaseAll(m)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_pgaio_error_cleanup(m)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_UnlockBuffers(m)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_ReleaseAuxProcessResources(m, int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_smgrdestroyall(m)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_AtEOXact_Files(m, int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_AtEOXact_HashTables(m, int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[9])) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_FlushErrorState(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_MemoryContextReset(m, v400)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v454 = int32(_a_F_WalWriterMain_4)
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[10])) = v456 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_pg_usleep(m, int32(_a_F_WalWriterMain_6))
	mBase = m.M
	goto L97
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_SetWalWriterSleeping(m, int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[14]))
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v478)+64)) = v480
	v484 = int32(0)
	v488 = int32(50)
	goto L111
L111:
	;
	v497 = base.B2i32(v488 < int32(2))
	if v497 != v484&int32(1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_SetWalWriterSleeping(m, v497)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L116
	}
L114:
	;
	v504 = v484
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[16]))
	v508 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = v508
	v513 = base.AtomicRmwOr32(m, v508, int32(_a_F_WalWriterMain_7), v508)
	goto L117
L116:
	;
	v504 = v497
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_ProcessMainLoopInterrupts(m)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	v519 = m.G0
	v521 = v519 - int32(32)
	m.G0 = v521
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v526 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[18])))
	if v526 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	m.G0 = v521 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	F_pgstat_report_wal(m, int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L5
	} else {
		goto L170
	}
L120:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v524)+308))
	v532 = base.B2i32(v530 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[18])) = uint8(v532)
	if v530 != int32(2) {
		v777 = int32(0)
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v524)+300))
	v538 = base.AtomicRmwXchg32(m, v524, int32(440), int32(1))
	if v538 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L122
L124:
	;
	F_s_lock(m, v524+int32(440), int32(_a_F_WalWriterMain_8))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L5
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v544 = int32(_a_F_WalWriterMain_9)
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v545)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+16)) = v546
	v548 = *(*int64)(unsafe.Add(mBase, uint32(v545)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+24)) = v548
	v550 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v545)+440)), uint32(v550))
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v521)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+16)) = v553 & int64(-8192)
	v557 = int32(_a_F_WalWriterMain_10)
	v558 = int64(0)
	v561 = base.AtomicRmwCmpxchg64(m, v545, int32(272), v558, v558)
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19])) = v561
	v566 = base.AtomicRmwOr32(m, v550, int32(_a_F_WalWriterMain_11), v550)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v573 = base.AtomicRmwCmpxchg64(m, v569, int32(264), v558, v558)
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[20])) = v573
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v521)+16))
	v577 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19]))
	v578 = base.B2i32(base.Ui64(v577) < base.Ui64(v575))
	if base.Ui64(v577) < base.Ui64(v575) {
		v618 = v575
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L126
L128:
	;
	v622 = m.G0
	v623 = int32(16)
	v624 = v622 - v623
	m.G0 = v624
	F_gettimeofday(m, v624)
	mBase = m.M
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v624)))
	v628 = int64(*(*int32)(unsafe.Add(mBase, uint32(v624)+8)))
	m.G0 = v624 + v623
	v636 = v628 + v627*int64(1000000) - int64(946684800000000)
	goto L138
L129:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v583 = base.AtomicRmwXchg32(m, v580, int32(440), int32(1))
	if v583 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_s_lock(m, v580+int32(440), int32(_a_F_WalWriterMain_8))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L5
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v589 = int32(0)
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v591)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+16)) = v592
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v591)+440)), uint32(v589))
	v598 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19]))
	if base.Ui64(v598) < base.Ui64(v592) {
		v618 = v592
		goto L128
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[21]))
	if v601 < int32(0) {
		v777 = v589
		goto L119
	} else {
		goto L135
	}
L135:
	;
	v605 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[22]))
	v607 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[20]))
	v611 = int64(*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[23])))
	v612 = base.I64_div_u_s(v607-int64(1), v611)
	if v605 == v612 {
		v777 = v589
		goto L119
	} else {
		goto L136
	}
L136:
	;
	F_XLogFileClose(m)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L5
	} else {
		goto L137
	}
L137:
	;
	v777 = v589
	goto L119
L138:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[24]))
	if v638 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v675 = int32(_a_F_WalWriterMain_12)
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[25])) = v677 + int32(1)
	v681 = F_WaitXLogInsertionsToFinish(m, v618)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L152
	}
L140:
	;
	v648 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19]))
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[26]))
	goto L145
L141:
	;
	v640 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[27]))
	if v640 != int64(0) {
		goto L140
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[27])) = v636
	*(*int64)(unsafe.Add(mBase, uint32(v521)+24)) = v618
	goto L139
L144:
	;
	goto L143
L145:
	;
	if base.I64_extend_i32_s(v650)*int64(1000) <= v636-v640 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[27])) = v636
	*(*int64)(unsafe.Add(mBase, uint32(v521)+24)) = v618
	goto L139
L147:
	;
	goto L148
L148:
	;
	v660 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[24]))
	v661 = int64(13)
	if v660 <= base.I32_wrap_i64(int64(base.Ui64(v618)>>(uint(v661)%64))-int64(base.Ui64(v648)>>(uint(v661)%64))) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[27])) = v636
	*(*int64)(unsafe.Add(mBase, uint32(v521)+24)) = v618
	goto L139
L150:
	;
	goto L151
L151:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v521)+24)) = int64(0)
	goto L139
L152:
	;
	v684 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[28]))
	v688 = F_LWLockAcquire(m, v684+int32(1024), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L5
	} else {
		goto L153
	}
L153:
	;
	v691 = int32(_a_F_WalWriterMain_9)
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v693 = int64(0)
	v696 = base.AtomicRmwCmpxchg64(m, v692, int32(272), v693, v693)
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19])) = v696
	v698 = int32(0)
	v701 = base.AtomicRmwOr32(m, v698, int32(_a_F_WalWriterMain_11), v698)
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v708 = base.AtomicRmwCmpxchg64(m, v704, int32(264), v693, v693)
	*(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[20])) = v708
	v710 = *(*int64)(unsafe.Add(mBase, uint32(v521)+16))
	if base.Ui64(v710) <= base.Ui64(v708) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[28]))
	F_LWLockRelease(m, v723+int32(1024))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L5
	} else {
		goto L160
	}
L155:
	;
	v712 = *(*int64)(unsafe.Add(mBase, uint32(v521)+24))
	v714 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19]))
	if base.Ui64(v712) <= base.Ui64(v714) {
		goto L154
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v716 = *(*int64)(unsafe.Add(mBase, uint32(v521)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v521)+8)) = v716
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v521)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v521))) = v718
	F_XLogWrite(m, v521, v535, v578)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L5
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	goto L154
L160:
	;
	v728 = int32(_a_F_WalWriterMain_12)
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[25]))
	v731 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[25])) = v730 - v731
	v736 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[18])))
	if v736 == v731 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[17]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+308))
	v743 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[18])) = uint8(base.B2i32(v742 != v743))
	v748 = base.B2i32(v742 == v743)
	goto L163
L162:
	;
	v748 = v731
	goto L163
L163:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[29])))
	if v750 != int32(1) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v764 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[30])) = uint8(v764)
	v768 = *(*int64)(unsafe.Add(mBase, _c_F_WalWriterMain[19]))
	F_WaitLSNWakeup(m, int32(3), v768)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L5
	} else {
		goto L168
	}
L165:
	;
	v754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_WalWriterMain[29])) = uint8(v754)
	v757 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[31]))
	if v757 <= v754 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	F_WalSndWakeup(m, int32(1), v748)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L5
	} else {
		goto L167
	}
L167:
	;
	goto L164
L168:
	;
	v771 = int32(1)
	F_AdvanceXLInsertBuffer(m, int64(0), v535, v771)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L5
	} else {
		goto L169
	}
L169:
	;
	v777 = v771
	goto L119
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+172)) = v400
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[16]))
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_WalWriterMain[26]))
	if v777 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v803 = int32(50)
	goto L173
L172:
	;
	v803 = v488 - base.B2i32(int32(0) < v488)
	goto L173
L173:
	;
	if int32(0) < v803 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v806 = v796
	goto L176
L175:
	;
	v806 = v796 * int32(25)
	goto L176
L176:
	;
	v808 = F_WaitLatch(m, v793, int32(41), v806, int32(83886097))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	v484 = v504
	v488 = v803
	goto L111
L178:
	;
	v828 = int32(v824)
	m.G0 = v16
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	if v16+int32(12) == v834 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	m.ExcPending = 1
	goto L187
L180:
	;
	if v838 != 0 {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	v838 = v836
	goto L183
L182:
	;
	v838 = int32(0)
	goto L183
L183:
	;
	goto L180
L184:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v16)+172))
	v19 = v838
	v22 = v839
	v24 = v830
	goto L1
L185:
	;
	goto L186
L186:
	;
	F___wasm_longjmp(m, v831, v830)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	return
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
