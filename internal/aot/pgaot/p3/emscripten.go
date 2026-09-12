package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__emscripten_memcpy_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	base.MemoryCopy(m, l0, l1, l2)
	return l0
}
func F__emscripten_memset_bulkmem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	base.MemoryFill(m, l0, l1, l2)
	return l0
}
func F__emscripten_stack_restore(m *base.Module, l0 int32) {
	m.G0 = l0
	return
}
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v9 int32
	_ = v9
	var v14 float64
	_ = v14
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v41 float64
	_ = v41
	var v48 float64
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v3 = float64(0)
	v9 = l0 << (uint(int32(3)) % 32)
	v14 = *(*float64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1328])))
	if base.F64_ne(v14, v3) != 0 {
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1329])))
		v18 = base.F64_max(l1, v17)
		v19 = base.F64_sub(v18, v17)
		if base.F64_lt(v19, float64(1.8446744073709552e+19))&base.F64_ge(v19, float64(0)) != 0 {
			v25 = base.I64_trunc_f64_u(v19)
			v27 = v25
		} else {
			v27 = int64(0)
		}
		if base.F64_lt(v14, float64(1.8446744073709552e+19))&base.F64_ge(v14, float64(0)) != 0 {
			v33 = base.I64_trunc_f64_u(v14)
			v35 = v33
		} else {
			v35 = int64(0)
		}
		v36 = base.I64_div_u_s(v27, v35)
		v41 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v36+int64(1)), v14), v17)
		*(*float64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1329]))) = v41
		v48 = base.F64_sub(v41, v18)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[1329]))) = int64(0)
		v48 = v3
	}
	v50 = m.Env.X_setitimer_js(m, l0, v48)
	mBase = m.M
	if l0 == int32(1) {
		v56 = int32(26)
	} else {
		v56 = int32(14)
	}
	if l0 == int32(2) {
		v59 = int32(27)
	} else {
		v59 = v56
	}
	v60 = F_raise(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		return
	} else {
		return
	}
}
func F_emscripten_builtin_calloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v3 = int32(0)
	if l0 == v3 {
		v22 = v3
	} else {
		v10 = base.I64_extend_i32_u(l0) * base.I64_extend_i32_u(l1)
		v11 = base.I32_wrap_i64(v10)
		if base.Ui32(l0|l1) < base.Ui32(int32(65536)) {
			v22 = v11
		} else {
			if base.I32_wrap_i64(int64(base.Ui64(v10)>>(uint(int64(32))%64))) != 0 {
				v19 = int32(-1)
			} else {
				v19 = v11
			}
			v22 = v19
		}
	}
	v23 = F_emscripten_builtin_malloc(m, v22)
	mBase = m.M
	if v23 == int32(0) {
	} else {
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23-int32(4)))))
		if v28&int32(3) == int32(0) {
		} else {
			v35 = F__emscripten_memset_bulkmem(m, v23, base.I32_extend8_s(int32(0)), v22)
			mBase = m.M
		}
	}
	return v23
}
func F_emscripten_builtin_realloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var __phi127 int32
	_ = __phi127
	var v129 int32
	_ = v129
	var __phi129 int32
	_ = __phi129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var __phi287 int32
	_ = __phi287
	var v289 int32
	_ = v289
	var __phi289 int32
	_ = __phi289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v369 int32
	_ = v369
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v481 int32
	_ = v481
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var __phi617 int32
	_ = __phi617
	var v621 int32
	_ = v621
	var __phi621 int32
	_ = __phi621
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var __phi774 int32
	_ = __phi774
	var v776 int32
	_ = v776
	var __phi776 int32
	_ = __phi776
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v807 int32
	_ = v807
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v855 int32
	_ = v855
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var __phi934 int32
	_ = __phi934
	var v936 int32
	_ = v936
	var __phi936 int32
	_ = __phi936
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v959 int32
	_ = v959
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v1016 int32
	_ = v1016
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1128 int32
	_ = v1128
	var v1170 int32
	_ = v1170
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = F_emscripten_builtin_malloc(m, l1)
	mBase = m.M
	return v16
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(-64)) <= base.Ui32(l1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(48)
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v26 = int32(11)
	if base.Ui32(l1) < base.Ui32(v26) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v1192 != 0 {
		goto L299
	} else {
		goto L300
	}
L8:
	;
	v32 = int32(16)
	goto L10
L9:
	;
	v32 = (l1 + v26) & int32(-8)
	goto L10
L10:
	;
	v34 = l0 - int32(8)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v37 = v35 & int32(-8)
	if v35&int32(3) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v1192 = v1170
	goto L7
L12:
	;
	if base.Ui32(v32) < base.Ui32(int32(256)) {
		v1170 = v3
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v55 = v37 + v34
	if base.Ui32(v32) <= base.Ui32(v37) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if base.Ui32(v32+int32(4)) <= base.Ui32(v37) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	if base.Ui32(v37-v32) <= base.Ui32(v49<<(uint(int32(1))%32)) {
		v1170 = v34
		goto L11
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v1192 = int32(0)
	goto L7
L19:
	;
	goto L18
L20:
	;
	v1170 = v34
	goto L11
L21:
	;
	v57 = v37 - v32
	if base.Ui32(v57) < base.Ui32(int32(16)) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[1352]))
	if v507 == v55 {
		goto L136
	} else {
		goto L137
	}
L24:
	;
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v60 | int32(2)
	v66 = v32 + v34
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v57 | int32(3)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v70 | v60
	v81 = v66 + v57
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v82&v60 != 0 {
		v199 = v66
		v200 = v57
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L20
L26:
	;
	goto L25
L27:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v208&int32(2) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L28:
	;
	if v82&int32(2) == int32(0) {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v90 = v89 + v57
	v91 = v66 - v89
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v91 != v93 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	if v109 == int32(0) {
		v199 = v91
		v200 = v90
		goto L27
	} else {
		goto L52
	}
L31:
	;
	v160 = int32(0)
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+12)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v98
	v199 = v91
	v200 = v90
	goto L27
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if base.Ui32(v89) <= base.Ui32(int32(255)) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v142 = int32(3)
	if v141&v142 != v142 {
		v199 = v91
		v200 = v90
		goto L27
	} else {
		goto L51
	}
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	if v95 != v98 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	if v91 != v95 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v100 = int32(4631368)
	v102 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v102 & base.I32_rotl(int32(-2), int32(base.Ui32(v89)>>(uint(int32(3))%32)))
	v199 = v91
	v200 = v90
	goto L27
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v111
	v160 = v95
	goto L30
L41:
	;
	goto L42
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v114 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v122 = v114
	v123 = v91 + int32(20)
	goto L45
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v117 == int32(0) {
		goto L31
	} else {
		goto L46
	}
L45:
	;
	__phi127 = v122
	__phi129 = v123
	v127 = __phi127
	v129 = __phi129
	goto L47
L46:
	;
	v122 = v117
	v123 = v91 + int32(16)
	goto L45
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+20))
	if v135 != 0 {
		__phi127 = v135
		__phi129 = v127 + int32(20)
		v127 = __phi127
		v129 = __phi129
		goto L47
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = int32(0)
	v160 = v127
	goto L30
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	if v138 != 0 {
		__phi127 = v138
		__phi129 = v127 + int32(16)
		v127 = __phi127
		v129 = __phi129
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v141 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v90 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v90
	goto L25
L52:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v171 = v169 << (uint(int32(2)) % 32)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v171)+uint32(_consts[1356])))
	if v174 == v91 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = v109
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v191 != 0 {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+uint32(_consts[1356]))) = v160
	if v160 != 0 {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	if v91 == v184 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v177 = int32(4631372)
	v179 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v179 & base.I32_rotl(int32(-2), v169)
	v199 = v91
	v200 = v90
	goto L27
L58:
	;
	if v160 == int32(0) {
		v199 = v91
		v200 = v90
		goto L27
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v160
	goto L58
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v160
	goto L58
L62:
	;
	goto L53
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+16)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v191)+24)) = v160
	goto L65
L64:
	;
	goto L65
L65:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v194 == int32(0) {
		v199 = v91
		v200 = v90
		goto L27
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+20)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v194)+24)) = v160
	v199 = v91
	v200 = v90
	goto L27
L67:
	;
	if base.Ui32(v369) <= base.Ui32(int32(255)) {
		goto L114
	} else {
		goto L115
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v252 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v199+v252))) = v252
	if v199 != v236 {
		v369 = v252
		goto L67
	} else {
		goto L113
	}
L69:
	;
	if v269 == int32(0) {
		goto L68
	} else {
		goto L98
	}
L70:
	;
	v312 = int32(0)
	goto L69
L71:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[1352]))
	if v214 == v81 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v208 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v200 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v199+v200))) = v200
	v369 = v200
	goto L67
L74:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1352])) = v199
	v218 = int32(4631380)
	v220 = *(*int32)(unsafe.Add(mBase, _consts[1358]))
	v221 = v220 + v200
	*(*int32)(unsafe.Add(mBase, _consts[1358])) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v221 | int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v199 != v227 {
		goto L26
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v236 == v81 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v230
	*(*int32)(unsafe.Add(mBase, _consts[1353])) = v230
	goto L25
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1353])) = v199
	v240 = int32(4631376)
	v242 = *(*int32)(unsafe.Add(mBase, _consts[1355]))
	v243 = v242 + v200
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v243 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v199+v243))) = v243
	goto L25
L79:
	;
	goto L80
L80:
	;
	v252 = v208&int32(-8) + v200
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if base.Ui32(v208) <= base.Ui32(int32(255)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v256 == v253 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
	if v253 != v81 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v258 = int32(4631368)
	v260 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v260 & base.I32_rotl(int32(-2), int32(base.Ui32(v208)>>(uint(int32(3))%32)))
	goto L68
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+12)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v256
	goto L68
L87:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v271)+12)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v271
	v312 = v253
	goto L69
L88:
	;
	goto L89
L89:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v274 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v282 = v274
	v283 = v81 + int32(20)
	goto L92
L91:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v277 == int32(0) {
		goto L70
	} else {
		goto L93
	}
L92:
	;
	__phi287 = v282
	__phi289 = v283
	v287 = __phi287
	v289 = __phi289
	goto L94
L93:
	;
	v282 = v277
	v283 = v81 + int32(16)
	goto L92
L94:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v287)+20))
	if v295 != 0 {
		__phi287 = v295
		__phi289 = v287 + int32(20)
		v287 = __phi287
		v289 = __phi289
		goto L94
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = int32(0)
	v312 = v287
	goto L69
L96:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	if v298 != 0 {
		__phi287 = v298
		__phi289 = v287 + int32(16)
		v287 = __phi287
		v289 = __phi289
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	v323 = v321 << (uint(int32(2)) % 32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[1356])))
	if v326 == v81 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+24)) = v269
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v343 != 0 {
		goto L109
	} else {
		goto L110
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_consts[1356]))) = v312
	if v312 != 0 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
	if v81 == v336 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v329 = int32(4631372)
	v331 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v331 & base.I32_rotl(int32(-2), v321)
	goto L68
L104:
	;
	if v312 == int32(0) {
		goto L68
	} else {
		goto L108
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+16)) = v312
	goto L104
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+20)) = v312
	goto L104
L108:
	;
	goto L99
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+16)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v343)+24)) = v312
	goto L111
L110:
	;
	goto L111
L111:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v346 == int32(0) {
		goto L68
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+20)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v346)+24)) = v312
	goto L68
L113:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v252
	goto L25
L114:
	;
	v380 = v369 & int32(-8)
	v382 = v380 + int32(4631408)
	v384 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	v388 = int32(1) << (uint(int32(base.Ui32(v369)>>(uint(int32(3))%32))) % 32)
	if v384&v388 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	goto L116
L116:
	;
	if base.Ui32(v369) <= base.Ui32(int32(16777215)) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380)+uint32(_consts[1359]))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v396)+12)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v396
	goto L25
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v388 | v384
	v396 = v382
	goto L117
L119:
	;
	goto L120
L120:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v380)+uint32(_consts[1359])))
	v396 = v395
	goto L117
L121:
	;
	v407 = base.I32_clz(int32(base.Ui32(v369) >> (uint(int32(8)) % 32)))
	v410 = int32(1)
	v417 = int32(base.Ui32(v369)>>(uint(int32(38)-v407)%32))&v410 - v407<<(uint(v410)%32) + int32(62)
	goto L123
L122:
	;
	v417 = int32(31)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+28)) = v417
	*(*int64)(unsafe.Add(mBase, uint32(v199)+16)) = int64(0)
	v422 = v417 << (uint(int32(2)) % 32)
	v426 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	v428 = int32(1) << (uint(v417) % 32)
	if v426&v428 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v481)+12)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v452)+8)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v481
	goto L26
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v199
	goto L25
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v428 | v426
	*(*int32)(unsafe.Add(mBase, uint32(v422)+uint32(_consts[1356]))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = v422 + int32(4631672)
	goto L125
L127:
	;
	goto L128
L128:
	;
	if v417 != int32(31) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v444 = int32(25) - int32(base.Ui32(v417)>>(uint(int32(1))%32))
	goto L131
L130:
	;
	v444 = int32(0)
	goto L131
L131:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v422)+uint32(_consts[1356])))
	v449 = v369 << (uint(v444) % 32)
	v452 = v446
	goto L132
L132:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v456&int32(-8) == v369 {
		goto L124
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v466)+16)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = v452
	goto L125
L134:
	;
	v466 = v452 + int32(base.Ui32(v449)>>(uint(int32(29))%32))&int32(4)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+16))
	if v467 != 0 {
		v449 = v449 << (uint(int32(1)) % 32)
		v452 = v467
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _consts[1358]))
	v511 = v510 + v37
	if base.Ui32(v511) <= base.Ui32(v32) {
		v1170 = v3
		goto L11
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v529 == v55 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v513 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v513 | int32(2)
	v519 = v32 + v34
	v520 = v511 - v32
	*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v520 | v513
	*(*int32)(unsafe.Add(mBase, _consts[1358])) = v520
	*(*int32)(unsafe.Add(mBase, _consts[1352])) = v519
	goto L20
L140:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _consts[1355]))
	v533 = v532 + v37
	if base.Ui32(v533) < base.Ui32(v32) {
		v1170 = v3
		goto L11
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v574&int32(2) != 0 {
		v1170 = v3
		goto L11
	} else {
		goto L148
	}
L143:
	;
	v535 = v533 - v32
	if base.Ui32(int32(16)) <= base.Ui32(v535) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1353])) = v568
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v569
	goto L20
L145:
	;
	v538 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v538 | int32(2)
	v544 = v32 + v34
	*(*int32)(unsafe.Add(mBase, uint32(v544)+4)) = v535 | v538
	v548 = v533 + v34
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = v535
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v548)+4)) = v550 & int32(-2)
	v568 = v544
	v569 = v535
	goto L144
L146:
	;
	goto L147
L147:
	;
	v554 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v35&v554 | v533 | int32(2)
	v560 = v533 + v34
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v561 | v554
	v565 = int32(0)
	v568 = v565
	v569 = v565
	goto L144
L148:
	;
	v579 = v574&int32(-8) + v37
	if base.Ui32(v579) < base.Ui32(v32) {
		v1170 = v3
		goto L11
	} else {
		goto L149
	}
L149:
	;
	v581 = v579 - v32
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if base.Ui32(v574) <= base.Ui32(int32(255)) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if base.Ui32(v581) <= base.Ui32(int32(15)) {
		goto L185
	} else {
		goto L186
	}
L151:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v585 == v582 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	if v582 != v55 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v587 = int32(4631368)
	v589 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v589 & base.I32_rotl(int32(-2), int32(base.Ui32(v574)>>(uint(int32(3))%32)))
	goto L150
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585)+12)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v582)+8)) = v585
	goto L150
L157:
	;
	if v598 == int32(0) {
		goto L150
	} else {
		goto L170
	}
L158:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+12)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v582)+8)) = v600
	v637 = v582
	goto L157
L159:
	;
	goto L160
L160:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v603 != 0 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v637 = int32(0)
	goto L157
L162:
	;
	v611 = v603
	v612 = v55 + int32(20)
	goto L164
L163:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v606 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L164:
	;
	__phi617 = v611
	__phi621 = v612
	v617 = __phi617
	v621 = __phi621
	goto L166
L165:
	;
	v611 = v606
	v612 = v55 + int32(16)
	goto L164
L166:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v617)+20))
	if v628 != 0 {
		__phi617 = v628
		__phi621 = v617 + int32(20)
		v617 = __phi617
		v621 = __phi621
		goto L166
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v621))) = int32(0)
	v637 = v617
	goto L157
L168:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
	if v631 != 0 {
		__phi617 = v631
		__phi621 = v617 + int32(16)
		v617 = __phi617
		v621 = __phi621
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v652 = v650 << (uint(int32(2)) % 32)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v652)+uint32(_consts[1356])))
	if v655 == v55 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637)+24)) = v598
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v672 != 0 {
		goto L181
	} else {
		goto L182
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652)+uint32(_consts[1356]))) = v637
	if v637 != 0 {
		goto L171
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v598)+16))
	if v55 == v665 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v658 = int32(4631372)
	v660 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v660 & base.I32_rotl(int32(-2), v650)
	goto L150
L176:
	;
	if v637 == int32(0) {
		goto L150
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+16)) = v637
	goto L176
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+20)) = v637
	goto L176
L180:
	;
	goto L171
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637)+16)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v672)+24)) = v637
	goto L183
L182:
	;
	goto L183
L183:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v675 == int32(0) {
		goto L150
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v637)+20)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v675)+24)) = v637
	goto L150
L185:
	;
	v695 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v35&v695 | v579 | int32(2)
	v701 = v34 + v579
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = v702 | v695
	goto L20
L186:
	;
	goto L187
L187:
	;
	v706 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v706 | int32(2)
	v712 = v32 + v34
	*(*int32)(unsafe.Add(mBase, uint32(v712)+4)) = v581 | int32(3)
	v716 = v34 + v579
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v716)+4)) = v717 | v706
	v728 = v712 + v581
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v729&v706 != 0 {
		v846 = v712
		v847 = v581
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L20
L189:
	;
	goto L188
L190:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v728)+4))
	if v855&int32(2) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L191:
	;
	if v729&int32(2) == int32(0) {
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	v737 = v736 + v581
	v738 = v712 - v736
	v740 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v738 != v740 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	if v756 == int32(0) {
		v846 = v738
		v847 = v737
		goto L190
	} else {
		goto L215
	}
L194:
	;
	v807 = int32(0)
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v745)+12)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v742)+8)) = v745
	v846 = v738
	v847 = v737
	goto L190
L196:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v738)+12))
	if base.Ui32(v736) <= base.Ui32(int32(255)) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v728)+4))
	v789 = int32(3)
	if v788&v789 != v789 {
		v846 = v738
		v847 = v737
		goto L190
	} else {
		goto L214
	}
L199:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	if v742 != v745 {
		goto L195
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v738)+24))
	if v738 != v742 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v747 = int32(4631368)
	v749 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v749 & base.I32_rotl(int32(-2), int32(base.Ui32(v736)>>(uint(int32(3))%32)))
	v846 = v738
	v847 = v737
	goto L190
L203:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v758)+12)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v742)+8)) = v758
	v807 = v742
	goto L193
L204:
	;
	goto L205
L205:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v738)+20))
	if v761 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v769 = v761
	v770 = v738 + int32(20)
	goto L208
L207:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v738)+16))
	if v764 == int32(0) {
		goto L194
	} else {
		goto L209
	}
L208:
	;
	__phi774 = v769
	__phi776 = v770
	v774 = __phi774
	v776 = __phi776
	goto L210
L209:
	;
	v769 = v764
	v770 = v738 + int32(16)
	goto L208
L210:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v774)+20))
	if v782 != 0 {
		__phi774 = v782
		__phi776 = v774 + int32(20)
		v774 = __phi774
		v776 = __phi776
		goto L210
	} else {
		goto L212
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v776))) = int32(0)
	v807 = v774
	goto L193
L212:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v774)+16))
	if v785 != 0 {
		__phi774 = v785
		__phi776 = v774 + int32(16)
		v774 = __phi774
		v776 = __phi776
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v728)+4)) = v788 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v738)+4)) = v737 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v728))) = v737
	goto L188
L215:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v738)+28))
	v818 = v816 << (uint(int32(2)) % 32)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v818)+uint32(_consts[1356])))
	if v821 == v738 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807)+24)) = v756
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v738)+16))
	if v838 != 0 {
		goto L226
	} else {
		goto L227
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v818)+uint32(_consts[1356]))) = v807
	if v807 != 0 {
		goto L216
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v756)+16))
	if v738 == v831 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v824 = int32(4631372)
	v826 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v826 & base.I32_rotl(int32(-2), v816)
	v846 = v738
	v847 = v737
	goto L190
L221:
	;
	if v807 == int32(0) {
		v846 = v738
		v847 = v737
		goto L190
	} else {
		goto L225
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v756)+16)) = v807
	goto L221
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v756)+20)) = v807
	goto L221
L225:
	;
	goto L216
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807)+16)) = v838
	*(*int32)(unsafe.Add(mBase, uint32(v838)+24)) = v807
	goto L228
L227:
	;
	goto L228
L228:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v738)+20))
	if v841 == int32(0) {
		v846 = v738
		v847 = v737
		goto L190
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807)+20)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v841)+24)) = v807
	v846 = v738
	v847 = v737
	goto L190
L230:
	;
	if base.Ui32(v1016) <= base.Ui32(int32(255)) {
		goto L277
	} else {
		goto L278
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v846)+4)) = v899 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v846+v899))) = v899
	if v846 != v883 {
		v1016 = v899
		goto L230
	} else {
		goto L276
	}
L232:
	;
	if v916 == int32(0) {
		goto L231
	} else {
		goto L261
	}
L233:
	;
	v959 = int32(0)
	goto L232
L234:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _consts[1352]))
	if v861 == v728 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v728)+4)) = v855 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v846)+4)) = v847 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v846+v847))) = v847
	v1016 = v847
	goto L230
L237:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1352])) = v846
	v865 = int32(4631380)
	v867 = *(*int32)(unsafe.Add(mBase, _consts[1358]))
	v868 = v867 + v847
	*(*int32)(unsafe.Add(mBase, _consts[1358])) = v868
	*(*int32)(unsafe.Add(mBase, uint32(v846)+4)) = v868 | int32(1)
	v874 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v846 != v874 {
		goto L189
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v883 = *(*int32)(unsafe.Add(mBase, _consts[1353]))
	if v883 == v728 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v877 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v877
	*(*int32)(unsafe.Add(mBase, _consts[1353])) = v877
	goto L188
L241:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1353])) = v846
	v887 = int32(4631376)
	v889 = *(*int32)(unsafe.Add(mBase, _consts[1355]))
	v890 = v889 + v847
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v890
	*(*int32)(unsafe.Add(mBase, uint32(v846)+4)) = v890 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v846+v890))) = v890
	goto L188
L242:
	;
	goto L243
L243:
	;
	v899 = v855&int32(-8) + v847
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v728)+12))
	if base.Ui32(v855) <= base.Ui32(int32(255)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v728)+8))
	if v903 == v900 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v728)+24))
	if v900 != v728 {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v905 = int32(4631368)
	v907 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v907 & base.I32_rotl(int32(-2), int32(base.Ui32(v855)>>(uint(int32(3))%32)))
	goto L231
L248:
	;
	goto L249
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v903)+12)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v900)+8)) = v903
	goto L231
L250:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v728)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v918)+12)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v900)+8)) = v918
	v959 = v900
	goto L232
L251:
	;
	goto L252
L252:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v728)+20))
	if v921 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v929 = v921
	v930 = v728 + int32(20)
	goto L255
L254:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v728)+16))
	if v924 == int32(0) {
		goto L233
	} else {
		goto L256
	}
L255:
	;
	__phi934 = v929
	__phi936 = v930
	v934 = __phi934
	v936 = __phi936
	goto L257
L256:
	;
	v929 = v924
	v930 = v728 + int32(16)
	goto L255
L257:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v934)+20))
	if v942 != 0 {
		__phi934 = v942
		__phi936 = v934 + int32(20)
		v934 = __phi934
		v936 = __phi936
		goto L257
	} else {
		goto L259
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v936))) = int32(0)
	v959 = v934
	goto L232
L259:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v934)+16))
	if v945 != 0 {
		__phi934 = v945
		__phi936 = v934 + int32(16)
		v934 = __phi934
		v936 = __phi936
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v728)+28))
	v970 = v968 << (uint(int32(2)) % 32)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v970)+uint32(_consts[1356])))
	if v973 == v728 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v959)+24)) = v916
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v728)+16))
	if v990 != 0 {
		goto L272
	} else {
		goto L273
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v970)+uint32(_consts[1356]))) = v959
	if v959 != 0 {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v916)+16))
	if v728 == v983 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v976 = int32(4631372)
	v978 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v978 & base.I32_rotl(int32(-2), v968)
	goto L231
L267:
	;
	if v959 == int32(0) {
		goto L231
	} else {
		goto L271
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v916)+16)) = v959
	goto L267
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v916)+20)) = v959
	goto L267
L271:
	;
	goto L262
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v959)+16)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v990)+24)) = v959
	goto L274
L273:
	;
	goto L274
L274:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v728)+20))
	if v993 == int32(0) {
		goto L231
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v959)+20)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v993)+24)) = v959
	goto L231
L276:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1355])) = v899
	goto L188
L277:
	;
	v1027 = v1016 & int32(-8)
	v1029 = v1027 + int32(4631408)
	v1031 = *(*int32)(unsafe.Add(mBase, _consts[1354]))
	v1035 = int32(1) << (uint(int32(base.Ui32(v1016)>>(uint(int32(3))%32))) % 32)
	if v1031&v1035 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L278:
	;
	goto L279
L279:
	;
	if base.Ui32(v1016) <= base.Ui32(int32(16777215)) {
		goto L284
	} else {
		goto L285
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1027)+uint32(_consts[1359]))) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+12)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v846)+12)) = v1029
	*(*int32)(unsafe.Add(mBase, uint32(v846)+8)) = v1043
	goto L188
L281:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1354])) = v1035 | v1031
	v1043 = v1029
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+uint32(_consts[1359])))
	v1043 = v1042
	goto L280
L284:
	;
	v1054 = base.I32_clz(int32(base.Ui32(v1016) >> (uint(int32(8)) % 32)))
	v1057 = int32(1)
	v1064 = int32(base.Ui32(v1016)>>(uint(int32(38)-v1054)%32))&v1057 - v1054<<(uint(v1057)%32) + int32(62)
	goto L286
L285:
	;
	v1064 = int32(31)
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v846)+28)) = v1064
	*(*int64)(unsafe.Add(mBase, uint32(v846)+16)) = int64(0)
	v1069 = v1064 << (uint(int32(2)) % 32)
	v1073 = *(*int32)(unsafe.Add(mBase, _consts[1357]))
	v1075 = int32(1) << (uint(v1064) % 32)
	if v1073&v1075 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1128)+12)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+8)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v846)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v846)+12)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(v846)+8)) = v1128
	goto L189
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v846)+12)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v846)+8)) = v846
	goto L188
L289:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1357])) = v1075 | v1073
	*(*int32)(unsafe.Add(mBase, uint32(v1069)+uint32(_consts[1356]))) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v846)+24)) = v1069 + int32(4631672)
	goto L288
L290:
	;
	goto L291
L291:
	;
	if v1064 != int32(31) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1091 = int32(25) - int32(base.Ui32(v1064)>>(uint(int32(1))%32))
	goto L294
L293:
	;
	v1091 = int32(0)
	goto L294
L294:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+uint32(_consts[1356])))
	v1096 = v1016 << (uint(v1091) % 32)
	v1099 = v1093
	goto L295
L295:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+4))
	if v1103&int32(-8) == v1016 {
		goto L287
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1113)+16)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v846)+24)) = v1099
	goto L288
L297:
	;
	v1113 = v1099 + int32(base.Ui32(v1096)>>(uint(int32(29))%32))&int32(4)
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+16))
	if v1114 != 0 {
		v1096 = v1096 << (uint(int32(1)) % 32)
		v1099 = v1114
		goto L295
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	return v1192 + int32(8)
L300:
	;
	goto L301
L301:
	;
	v1196 = F_emscripten_builtin_malloc(m, l1)
	mBase = m.M
	if v1196 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	return int32(0)
L303:
	;
	goto L304
L304:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(4))))
	if v1205&int32(3) != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1208 = int32(-4)
	goto L307
L306:
	;
	v1208 = int32(-8)
	goto L307
L307:
	;
	v1211 = v1208 + v1205&int32(-8)
	if base.Ui32(v1211) < base.Ui32(l1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1213 = v1211
	goto L310
L309:
	;
	v1213 = l1
	goto L310
L310:
	;
	if v1213 != 0 {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return v1196
L312:
	;
	v1214 = F__emscripten_memcpy_bulkmem(m, v1196, l0, v1213)
	mBase = m.M
	goto L314
L313:
	;
	goto L314
L314:
	;
	goto L311
}
