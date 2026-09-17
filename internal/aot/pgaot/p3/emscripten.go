package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__emscripten_stack_restore(m *base.Module, l0 int32) {
	m.G0 = l0
	return
}
func F__emscripten_timeout(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v8 int32
	_ = v8
	var v13 float64
	_ = v13
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v21 int64
	_ = v21
	var v26 float64
	_ = v26
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v3 = float64(0)
	v8 = l0 << (uint(int32(3)) % 32)
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F__emscripten_timeout[0])))
	if base.F64_ne(v13, v3) != 0 {
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F__emscripten_timeout[1])))
		v17 = base.F64_max(l1, v16)
		v21 = base.I64_div_u_s(base.I64_trunc_sat_f64_u(base.F64_sub(v17, v16)), base.I64_trunc_sat_f64_u(v13))
		v26 = base.F64_add(base.F64_mul(base.F64_convert_i64_u(v21+int64(1)), v13), v16)
		*(*float64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F__emscripten_timeout[1]))) = v26
		v33 = base.F64_sub(v26, v17)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8)+uint32(_c_F__emscripten_timeout[1]))) = int64(0)
		v33 = v3
	}
	v34 = m.Env.X_setitimer_js(m, l0, v33)
	mBase = m.M
	if l0 == int32(1) {
		v40 = int32(26)
	} else {
		v40 = int32(14)
	}
	if l0 == int32(2) {
		v43 = int32(27)
	} else {
		v43 = v40
	}
	v44 = F_raise(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		return
	} else {
		return
	}
}
func F_emscripten_builtin_calloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	v4 = int32(0)
	if l0 == v4 {
		v22 = v4
	} else {
		v10 = base.I64_extend_i32_u(l0) * base.I64_extend_i32_u(l1)
		v11 = base.I32_wrap_i64(v10)
		if base.Ui32(l0|l1) < base.Ui32(int32(_a_F_emscripten_builtin_calloc_0)) {
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
			v33 = int32(0)
			if v22 == v33 {
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v33)
				v40 = v23 + v22
				*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(1)))) = uint8(v33)
				if base.Ui32(v22) < base.Ui32(int32(3)) {
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v33)
					*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v33)
					*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(3)))) = uint8(v33)
					*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(2)))) = uint8(v33)
					if base.Ui32(v22) < base.Ui32(int32(7)) {
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)) = uint8(v33)
						*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(4)))) = uint8(v33)
						if base.Ui32(v22) < base.Ui32(int32(9)) {
						} else {
							v62 = int32(0)
							v65 = (v62 - v23) & int32(3)
							v66 = v23 + v65
							*(*int32)(unsafe.Add(mBase, uint32(v66))) = v62
							v74 = (v22 - v65) & int32(-4)
							v75 = v66 + v74
							*(*int32)(unsafe.Add(mBase, uint32(v75-int32(4)))) = v62
							if base.Ui32(v74) < base.Ui32(int32(9)) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v75-int32(8)))) = v62
								*(*int32)(unsafe.Add(mBase, uint32(v75-int32(12)))) = v62
								if base.Ui32(v74) < base.Ui32(int32(25)) {
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v66)+16)) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v75-int32(16)))) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v75-int32(20)))) = v62
									v101 = int32(24)
									*(*int32)(unsafe.Add(mBase, uint32(v75-v101))) = v62
									*(*int32)(unsafe.Add(mBase, uint32(v75-int32(28)))) = v62
									v110 = v66&int32(4) | v101
									v111 = v74 - v110
									if base.Ui32(v111) < base.Ui32(int32(32)) {
									} else {
										v116 = base.I64_extend_i32_u(v62) * int64(4294967297)
										v119 = v110 + v66
										v120 = v111
										for {
											*(*int64)(unsafe.Add(mBase, uint32(v119)+24)) = v116
											*(*int64)(unsafe.Add(mBase, uint32(v119)+16)) = v116
											*(*int64)(unsafe.Add(mBase, uint32(v119)+8)) = v116
											*(*int64)(unsafe.Add(mBase, uint32(v119))) = v116
											v128 = int32(32)
											v131 = v120 - v128
											if base.Ui32(int32(31)) < base.Ui32(v131) {
												v119 = v119 + v128
												v120 = v131
												continue
											} else {
												break
											}
											break
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
	var v126 int32
	_ = v126
	var __phi126 int32
	_ = __phi126
	var v128 int32
	_ = v128
	var __phi128 int32
	_ = __phi128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
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
	var v286 int32
	_ = v286
	var __phi286 int32
	_ = __phi286
	var v288 int32
	_ = v288
	var __phi288 int32
	_ = __phi288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
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
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v482 int32
	_ = v482
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var __phi616 int32
	_ = __phi616
	var v622 int32
	_ = v622
	var __phi622 int32
	_ = __phi622
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var __phi774 int32
	_ = __phi774
	var v776 int32
	_ = v776
	var __phi776 int32
	_ = __phi776
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
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
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var __phi934 int32
	_ = __phi934
	var v936 int32
	_ = v936
	var __phi936 int32
	_ = __phi936
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v961 int32
	_ = v961
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1017 int32
	_ = v1017
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1130 int32
	_ = v1130
	var v1170 int32
	_ = v1170
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1222 int32
	_ = v1222
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
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
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[0])) = int32(48)
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
	if v1194 != 0 {
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
	v1194 = v1170
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
	v55 = v34 + v37
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
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[1]))
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
	v1194 = int32(0)
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
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[2]))
	if v508 == v55 {
		goto L136
	} else {
		goto L137
	}
L24:
	;
	v60 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v60 | int32(2)
	v66 = v34 + v32
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
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
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
	v161 = int32(0)
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
	v100 = int32(_a_F_emscripten_builtin_realloc_0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v102 & base.I32_rotl(int32(-2), int32(base.Ui32(v89)>>(uint(int32(3))%32)))
	v199 = v91
	v200 = v90
	goto L27
L40:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v111
	v161 = v95
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
	__phi126 = v123
	__phi128 = v122
	v126 = __phi126
	v128 = __phi128
	goto L47
L46:
	;
	v122 = v117
	v123 = v91 + int32(16)
	goto L45
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	if v135 != 0 {
		__phi126 = v128 + int32(20)
		__phi128 = v135
		v126 = __phi126
		v128 = __phi128
		goto L47
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = int32(0)
	v161 = v128
	goto L30
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v138 != 0 {
		__phi126 = v128 + int32(16)
		__phi128 = v138
		v126 = __phi126
		v128 = __phi128
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v141 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v90 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v90
	goto L25
L52:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v171 = v169 << (uint(int32(2)) % 32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+uint32(_c_F_emscripten_builtin_realloc[6])))
	if v172 == v91 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = v109
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v191 != 0 {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v161
	if v161 != 0 {
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
	v177 = int32(_a_F_emscripten_builtin_realloc_1)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v179 & base.I32_rotl(int32(-2), v169)
	v199 = v91
	v200 = v90
	goto L27
L58:
	;
	if v161 == int32(0) {
		v199 = v91
		v200 = v90
		goto L27
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v161
	goto L58
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v161
	goto L58
L62:
	;
	goto L53
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v191)+24)) = v161
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
	*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v194)+24)) = v161
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
	v313 = int32(0)
	goto L69
L71:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[2]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[2])) = v199
	v218 = int32(_a_F_emscripten_builtin_realloc_2)
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[8]))
	v221 = v220 + v200
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[8])) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v221 | int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
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
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
	if v236 == v81 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v230
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3])) = v230
	goto L25
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3])) = v199
	v240 = int32(_a_F_emscripten_builtin_realloc_3)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5]))
	v243 = v242 + v200
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v243
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
	v258 = int32(_a_F_emscripten_builtin_realloc_0)
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v260 & base.I32_rotl(int32(-2), int32(base.Ui32(v208)>>(uint(int32(3))%32)))
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
	v313 = v253
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
	__phi286 = v283
	__phi288 = v282
	v286 = __phi286
	v288 = __phi288
	goto L94
L93:
	;
	v282 = v277
	v283 = v81 + int32(16)
	goto L92
L94:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v288)+20))
	if v295 != 0 {
		__phi286 = v288 + int32(20)
		__phi288 = v295
		v286 = __phi286
		v288 = __phi288
		goto L94
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = int32(0)
	v313 = v288
	goto L69
L96:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	if v298 != 0 {
		__phi286 = v288 + int32(16)
		__phi288 = v298
		v286 = __phi286
		v288 = __phi288
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
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_emscripten_builtin_realloc[6])))
	if v324 == v81 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+24)) = v269
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v343 != 0 {
		goto L109
	} else {
		goto L110
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v313
	if v313 != 0 {
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
	v329 = int32(_a_F_emscripten_builtin_realloc_1)
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v331 & base.I32_rotl(int32(-2), v321)
	goto L68
L104:
	;
	if v313 == int32(0) {
		goto L68
	} else {
		goto L108
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+16)) = v313
	goto L104
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+20)) = v313
	goto L104
L108:
	;
	goto L99
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+16)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v343)+24)) = v313
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
	*(*int32)(unsafe.Add(mBase, uint32(v313)+20)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v346)+24)) = v313
	goto L68
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v252
	goto L25
L114:
	;
	v380 = v369 & int32(248)
	v382 = v380 + int32(_a_F_emscripten_builtin_realloc_4)
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v380)+uint32(_c_F_emscripten_builtin_realloc[9]))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v396)+12)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v396
	goto L25
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v388 | v384
	v396 = v382
	goto L117
L119:
	;
	goto L120
L120:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v380)+uint32(_c_F_emscripten_builtin_realloc[9])))
	v396 = v395
	goto L117
L121:
	;
	v407 = base.I32_clz(int32(base.Ui32(v369) >> (uint(int32(8)) % 32)))
	v410 = int32(1)
	v418 = int32(base.Ui32(v369)>>(uint(int32(38)-v407)%32))&v410 | v407<<(uint(v410)%32) ^ int32(62)
	goto L123
L122:
	;
	v418 = int32(31)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+28)) = v418
	*(*int64)(unsafe.Add(mBase, uint32(v199)+16)) = int64(0)
	v423 = v418 << (uint(int32(2)) % 32)
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	v429 = int32(1) << (uint(v418) % 32)
	if v427&v429 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v450)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v482)+12)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v450)+8)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v482
	goto L26
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v199
	goto L25
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v427 | v429
	*(*int32)(unsafe.Add(mBase, uint32(v423)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = v423 + int32(_a_F_emscripten_builtin_realloc_5)
	goto L125
L127:
	;
	goto L128
L128:
	;
	if v418 != int32(31) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v445 = int32(25) - int32(base.Ui32(v418)>>(uint(int32(1))%32))
	goto L131
L130:
	;
	v445 = int32(0)
	goto L131
L131:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v423)+uint32(_c_F_emscripten_builtin_realloc[6])))
	v450 = v447
	v451 = v369 << (uint(v445) % 32)
	goto L132
L132:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v457&int32(-8) == v369 {
		goto L124
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v467)+16)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v199)+24)) = v450
	goto L125
L134:
	;
	v467 = v450 + int32(base.Ui32(v451)>>(uint(int32(29))%32))&int32(4)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+16))
	if v468 != 0 {
		v450 = v468
		v451 = v451 << (uint(int32(1)) % 32)
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[8]))
	v512 = v511 + v37
	if base.Ui32(v512) <= base.Ui32(v32) {
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
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
	if v530 == v55 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v514 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v514 | int32(2)
	v520 = v34 + v32
	v521 = v512 - v32
	*(*int32)(unsafe.Add(mBase, uint32(v520)+4)) = v521 | v514
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[8])) = v521
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[2])) = v520
	goto L20
L140:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5]))
	v534 = v533 + v37
	if base.Ui32(v534) < base.Ui32(v32) {
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
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v575&int32(2) != 0 {
		v1170 = v3
		goto L11
	} else {
		goto L148
	}
L143:
	;
	v536 = v534 - v32
	if base.Ui32(int32(16)) <= base.Ui32(v536) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3])) = v570
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v568
	goto L20
L145:
	;
	v539 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v539 | int32(2)
	v545 = v34 + v32
	*(*int32)(unsafe.Add(mBase, uint32(v545)+4)) = v536 | v539
	v549 = v534 + v34
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = v536
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v549)+4)) = v551 & int32(-2)
	v568 = v536
	v570 = v545
	goto L144
L146:
	;
	goto L147
L147:
	;
	v555 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v35&v555 | v534 | int32(2)
	v561 = v534 + v34
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v561)+4)) = v562 | v555
	v566 = int32(0)
	v568 = v566
	v570 = v566
	goto L144
L148:
	;
	v580 = v575&int32(-8) + v37
	if base.Ui32(v580) < base.Ui32(v32) {
		v1170 = v3
		goto L11
	} else {
		goto L149
	}
L149:
	;
	v582 = v580 - v32
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if base.Ui32(v575) <= base.Ui32(int32(255)) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if base.Ui32(v582) <= base.Ui32(int32(15)) {
		goto L185
	} else {
		goto L186
	}
L151:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v586 == v583 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	if v583 != v55 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v588 = int32(_a_F_emscripten_builtin_realloc_0)
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v590 & base.I32_rotl(int32(-2), int32(base.Ui32(v575)>>(uint(int32(3))%32)))
	goto L150
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+12)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v583)+8)) = v586
	goto L150
L157:
	;
	if v599 == int32(0) {
		goto L150
	} else {
		goto L170
	}
L158:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+12)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v583)+8)) = v601
	v641 = v583
	goto L157
L159:
	;
	goto L160
L160:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v604 != 0 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v641 = int32(0)
	goto L157
L162:
	;
	v612 = v604
	v613 = v55 + int32(20)
	goto L164
L163:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v607 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L164:
	;
	__phi616 = v612
	__phi622 = v613
	v616 = __phi616
	v622 = __phi622
	goto L166
L165:
	;
	v612 = v607
	v613 = v55 + int32(16)
	goto L164
L166:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v616)+20))
	if v629 != 0 {
		__phi616 = v629
		__phi622 = v616 + int32(20)
		v616 = __phi616
		v622 = __phi622
		goto L166
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = int32(0)
	v641 = v616
	goto L157
L168:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v616)+16))
	if v632 != 0 {
		__phi616 = v632
		__phi622 = v616 + int32(16)
		v616 = __phi616
		v622 = __phi622
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v653 = v651 << (uint(int32(2)) % 32)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v653)+uint32(_c_F_emscripten_builtin_realloc[6])))
	if v654 == v55 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641)+24)) = v599
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v673 != 0 {
		goto L181
	} else {
		goto L182
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v641
	if v641 != 0 {
		goto L171
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v599)+16))
	if v55 == v666 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v659 = int32(_a_F_emscripten_builtin_realloc_1)
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v661 & base.I32_rotl(int32(-2), v651)
	goto L150
L176:
	;
	if v641 == int32(0) {
		goto L150
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v599)+16)) = v641
	goto L176
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v599)+20)) = v641
	goto L176
L180:
	;
	goto L171
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641)+16)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v673)+24)) = v641
	goto L183
L182:
	;
	goto L183
L183:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v676 == int32(0) {
		goto L150
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v641)+20)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v676)+24)) = v641
	goto L150
L185:
	;
	v696 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v35&v696 | v580 | int32(2)
	v702 = v34 + v580
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v702)+4)) = v703 | v696
	goto L20
L186:
	;
	goto L187
L187:
	;
	v707 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32 | v35&v707 | int32(2)
	v713 = v34 + v32
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v582 | int32(3)
	v717 = v34 + v580
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v717)+4)) = v718 | v707
	v729 = v713 + v582
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v730&v707 != 0 {
		v847 = v713
		v848 = v582
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
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	if v856&int32(2) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L191:
	;
	if v730&int32(2) == int32(0) {
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	v738 = v737 + v582
	v739 = v713 - v737
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
	if v739 != v741 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	if v757 == int32(0) {
		v847 = v739
		v848 = v738
		goto L190
	} else {
		goto L215
	}
L194:
	;
	v809 = int32(0)
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v746)+12)) = v743
	*(*int32)(unsafe.Add(mBase, uint32(v743)+8)) = v746
	v847 = v739
	v848 = v738
	goto L190
L196:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v739)+12))
	if base.Ui32(v737) <= base.Ui32(int32(255)) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v729)+4))
	v790 = int32(3)
	if v789&v790 != v790 {
		v847 = v739
		v848 = v738
		goto L190
	} else {
		goto L214
	}
L199:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	if v743 != v746 {
		goto L195
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v739)+24))
	if v739 != v743 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v748 = int32(_a_F_emscripten_builtin_realloc_0)
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v750 & base.I32_rotl(int32(-2), int32(base.Ui32(v737)>>(uint(int32(3))%32)))
	v847 = v739
	v848 = v738
	goto L190
L203:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v739)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v759)+12)) = v743
	*(*int32)(unsafe.Add(mBase, uint32(v743)+8)) = v759
	v809 = v743
	goto L193
L204:
	;
	goto L205
L205:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v739)+20))
	if v762 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v770 = v762
	v771 = v739 + int32(20)
	goto L208
L207:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v739)+16))
	if v765 == int32(0) {
		goto L194
	} else {
		goto L209
	}
L208:
	;
	__phi774 = v771
	__phi776 = v770
	v774 = __phi774
	v776 = __phi776
	goto L210
L209:
	;
	v770 = v765
	v771 = v739 + int32(16)
	goto L208
L210:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v776)+20))
	if v783 != 0 {
		__phi774 = v776 + int32(20)
		__phi776 = v783
		v774 = __phi774
		v776 = __phi776
		goto L210
	} else {
		goto L212
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v774))) = int32(0)
	v809 = v776
	goto L193
L212:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v776)+16))
	if v786 != 0 {
		__phi774 = v776 + int32(16)
		__phi776 = v786
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
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v729)+4)) = v789 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v739)+4)) = v738 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v729))) = v738
	goto L188
L215:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v739)+28))
	v819 = v817 << (uint(int32(2)) % 32)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+uint32(_c_F_emscripten_builtin_realloc[6])))
	if v820 == v739 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v809)+24)) = v757
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v739)+16))
	if v839 != 0 {
		goto L226
	} else {
		goto L227
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v819)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v809
	if v809 != 0 {
		goto L216
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v757)+16))
	if v739 == v832 {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	v825 = int32(_a_F_emscripten_builtin_realloc_1)
	v827 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v827 & base.I32_rotl(int32(-2), v817)
	v847 = v739
	v848 = v738
	goto L190
L221:
	;
	if v809 == int32(0) {
		v847 = v739
		v848 = v738
		goto L190
	} else {
		goto L225
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v757)+16)) = v809
	goto L221
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v757)+20)) = v809
	goto L221
L225:
	;
	goto L216
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v809)+16)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v839)+24)) = v809
	goto L228
L227:
	;
	goto L228
L228:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v739)+20))
	if v842 == int32(0) {
		v847 = v739
		v848 = v738
		goto L190
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v809)+20)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v842)+24)) = v809
	v847 = v739
	v848 = v738
	goto L190
L230:
	;
	if base.Ui32(v1017) <= base.Ui32(int32(255)) {
		goto L277
	} else {
		goto L278
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+4)) = v900 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v847+v900))) = v900
	if v847 != v884 {
		v1017 = v900
		goto L230
	} else {
		goto L276
	}
L232:
	;
	if v917 == int32(0) {
		goto L231
	} else {
		goto L261
	}
L233:
	;
	v961 = int32(0)
	goto L232
L234:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[2]))
	if v862 == v729 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729)+4)) = v856 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v847)+4)) = v848 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v847+v848))) = v848
	v1017 = v848
	goto L230
L237:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[2])) = v847
	v866 = int32(_a_F_emscripten_builtin_realloc_2)
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[8]))
	v869 = v868 + v848
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[8])) = v869
	*(*int32)(unsafe.Add(mBase, uint32(v847)+4)) = v869 | int32(1)
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
	if v847 != v875 {
		goto L189
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v884 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3]))
	if v884 == v729 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v878 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v878
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3])) = v878
	goto L188
L241:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[3])) = v847
	v888 = int32(_a_F_emscripten_builtin_realloc_3)
	v890 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5]))
	v891 = v890 + v848
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v847)+4)) = v891 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v847+v891))) = v891
	goto L188
L242:
	;
	goto L243
L243:
	;
	v900 = v856&int32(-8) + v848
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v729)+12))
	if base.Ui32(v856) <= base.Ui32(int32(255)) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	if v904 == v901 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v729)+24))
	if v901 != v729 {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v906 = int32(_a_F_emscripten_builtin_realloc_0)
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v908 & base.I32_rotl(int32(-2), int32(base.Ui32(v856)>>(uint(int32(3))%32)))
	goto L231
L248:
	;
	goto L249
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v904)+12)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v901)+8)) = v904
	goto L231
L250:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v919)+12)) = v901
	*(*int32)(unsafe.Add(mBase, uint32(v901)+8)) = v919
	v961 = v901
	goto L232
L251:
	;
	goto L252
L252:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v729)+20))
	if v922 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v930 = v922
	v931 = v729 + int32(20)
	goto L255
L254:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	if v925 == int32(0) {
		goto L233
	} else {
		goto L256
	}
L255:
	;
	__phi934 = v931
	__phi936 = v930
	v934 = __phi934
	v936 = __phi936
	goto L257
L256:
	;
	v930 = v925
	v931 = v729 + int32(16)
	goto L255
L257:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v936)+20))
	if v943 != 0 {
		__phi934 = v936 + int32(20)
		__phi936 = v943
		v934 = __phi934
		v936 = __phi936
		goto L257
	} else {
		goto L259
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v934))) = int32(0)
	v961 = v936
	goto L232
L259:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v936)+16))
	if v946 != 0 {
		__phi934 = v936 + int32(16)
		__phi936 = v946
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
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v729)+28))
	v971 = v969 << (uint(int32(2)) % 32)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v971)+uint32(_c_F_emscripten_builtin_realloc[6])))
	if v972 == v729 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v961)+24)) = v917
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v729)+16))
	if v991 != 0 {
		goto L272
	} else {
		goto L273
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v971)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v961
	if v961 != 0 {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v917)+16))
	if v729 == v984 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v977 = int32(_a_F_emscripten_builtin_realloc_1)
	v979 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v979 & base.I32_rotl(int32(-2), v969)
	goto L231
L267:
	;
	if v961 == int32(0) {
		goto L231
	} else {
		goto L271
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v917)+16)) = v961
	goto L267
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v917)+20)) = v961
	goto L267
L271:
	;
	goto L262
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v961)+16)) = v991
	*(*int32)(unsafe.Add(mBase, uint32(v991)+24)) = v961
	goto L274
L273:
	;
	goto L274
L274:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v729)+20))
	if v994 == int32(0) {
		goto L231
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v961)+20)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v994)+24)) = v961
	goto L231
L276:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[5])) = v900
	goto L188
L277:
	;
	v1028 = v1017 & int32(248)
	v1030 = v1028 + int32(_a_F_emscripten_builtin_realloc_4)
	v1032 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4]))
	v1036 = int32(1) << (uint(int32(base.Ui32(v1017)>>(uint(int32(3))%32))) % 32)
	if v1032&v1036 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L278:
	;
	goto L279
L279:
	;
	if base.Ui32(v1017) <= base.Ui32(int32(16777215)) {
		goto L284
	} else {
		goto L285
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+uint32(_c_F_emscripten_builtin_realloc[9]))) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+12)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v847)+12)) = v1030
	*(*int32)(unsafe.Add(mBase, uint32(v847)+8)) = v1044
	goto L188
L281:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[4])) = v1036 | v1032
	v1044 = v1030
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+uint32(_c_F_emscripten_builtin_realloc[9])))
	v1044 = v1043
	goto L280
L284:
	;
	v1055 = base.I32_clz(int32(base.Ui32(v1017) >> (uint(int32(8)) % 32)))
	v1058 = int32(1)
	v1066 = int32(base.Ui32(v1017)>>(uint(int32(38)-v1055)%32))&v1058 | v1055<<(uint(v1058)%32) ^ int32(62)
	goto L286
L285:
	;
	v1066 = int32(31)
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+28)) = v1066
	*(*int64)(unsafe.Add(mBase, uint32(v847)+16)) = int64(0)
	v1071 = v1066 << (uint(int32(2)) % 32)
	v1075 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7]))
	v1077 = int32(1) << (uint(v1066) % 32)
	if v1075&v1077 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+12)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+8)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v847)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v847)+12)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v847)+8)) = v1130
	goto L189
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+12)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v847)+8)) = v847
	goto L188
L289:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_realloc[7])) = v1075 | v1077
	*(*int32)(unsafe.Add(mBase, uint32(v1071)+uint32(_c_F_emscripten_builtin_realloc[6]))) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v847)+24)) = v1071 + int32(_a_F_emscripten_builtin_realloc_5)
	goto L288
L290:
	;
	goto L291
L291:
	;
	if v1066 != int32(31) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1093 = int32(25) - int32(base.Ui32(v1066)>>(uint(int32(1))%32))
	goto L294
L293:
	;
	v1093 = int32(0)
	goto L294
L294:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1071)+uint32(_c_F_emscripten_builtin_realloc[6])))
	v1098 = v1095
	v1099 = v1017 << (uint(v1093) % 32)
	goto L295
L295:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	if v1105&int32(-8) == v1017 {
		goto L287
	} else {
		goto L297
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+16)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v847)+24)) = v1098
	goto L288
L297:
	;
	v1115 = v1098 + int32(base.Ui32(v1099)>>(uint(int32(29))%32))&int32(4)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+16))
	if v1116 != 0 {
		v1098 = v1116
		v1099 = v1099 << (uint(int32(1)) % 32)
		goto L295
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	return v1194 + int32(8)
L300:
	;
	goto L301
L301:
	;
	v1198 = F_emscripten_builtin_malloc(m, l1)
	mBase = m.M
	if v1198 == int32(0) {
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
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(4))))
	if v1207&int32(3) != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1210 = int32(-4)
	goto L307
L306:
	;
	v1210 = int32(-8)
	goto L307
L307:
	;
	v1213 = v1210 + v1207&int32(-8)
	if base.Ui32(v1213) < base.Ui32(l1) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1215 = v1213
	goto L310
L309:
	;
	v1215 = l1
	goto L310
L310:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v1215) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return v1198
L312:
	;
	if v1215 != 0 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	goto L314
L314:
	;
	v1222 = v1198 + v1215
	if (v1198^l0)&int32(3) == int32(0) {
		goto L319
	} else {
		goto L320
	}
L315:
	;
	base.MemoryCopy(m, v1198, l0, v1215)
	goto L317
L316:
	;
	goto L317
L317:
	;
	goto L311
L318:
	;
	if base.Ui32(v1354) < base.Ui32(v1222) {
		goto L352
	} else {
		goto L353
	}
L319:
	;
	if v1198&int32(3) == int32(0) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	goto L321
L321:
	;
	if base.Ui32(v1222) < base.Ui32(int32(4)) {
		goto L343
	} else {
		goto L344
	}
L322:
	;
	v1258 = v1222 & int32(-4)
	if base.Ui32(v1222) < base.Ui32(int32(64)) {
		v1308 = v1252
		v1309 = v1253
		goto L333
	} else {
		goto L334
	}
L323:
	;
	v1252 = l0
	v1253 = v1198
	goto L322
L324:
	;
	goto L325
L325:
	;
	if v1215 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1252 = l0
	v1253 = v1198
	goto L322
L327:
	;
	goto L328
L328:
	;
	v1235 = l0
	v1236 = v1198
	goto L329
L329:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1235))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1236))) = uint8(v1240)
	v1242 = int32(1)
	v1243 = v1235 + v1242
	v1245 = v1236 + v1242
	if v1245&int32(3) == int32(0) {
		v1252 = v1243
		v1253 = v1245
		goto L322
	} else {
		goto L331
	}
L330:
	;
	v1252 = v1243
	v1253 = v1245
	goto L322
L331:
	;
	if base.Ui32(v1245) < base.Ui32(v1222) {
		v1235 = v1243
		v1236 = v1245
		goto L329
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	if base.Ui32(v1258) <= base.Ui32(v1309) {
		v1353 = v1308
		v1354 = v1309
		goto L318
	} else {
		goto L339
	}
L334:
	;
	v1262 = v1258 + int32(-64)
	if base.Ui32(v1262) < base.Ui32(v1253) {
		v1308 = v1252
		v1309 = v1253
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1265 = v1252
	v1266 = v1253
	goto L336
L336:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1265)))
	*(*int32)(unsafe.Add(mBase, uint32(v1266))) = v1270
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+4)) = v1272
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+8)) = v1274
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+12)) = v1276
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+16)) = v1278
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+20)) = v1280
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+24)) = v1282
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+28)) = v1284
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+32)) = v1286
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+36)) = v1288
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+40)) = v1290
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+44)) = v1292
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+48)) = v1294
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+52)) = v1296
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+56)) = v1298
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1265)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1266)+60)) = v1300
	v1302 = int32(-64)
	v1303 = v1265 - v1302
	v1305 = v1266 - v1302
	if base.Ui32(v1305) <= base.Ui32(v1262) {
		v1265 = v1303
		v1266 = v1305
		goto L336
	} else {
		goto L338
	}
L337:
	;
	v1308 = v1303
	v1309 = v1305
	goto L333
L338:
	;
	goto L337
L339:
	;
	v1315 = v1308
	v1316 = v1309
	goto L340
L340:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1315)))
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1320
	v1322 = int32(4)
	v1323 = v1315 + v1322
	v1325 = v1316 + v1322
	if base.Ui32(v1325) < base.Ui32(v1258) {
		v1315 = v1323
		v1316 = v1325
		goto L340
	} else {
		goto L342
	}
L341:
	;
	v1353 = v1323
	v1354 = v1325
	goto L318
L342:
	;
	goto L341
L343:
	;
	v1353 = l0
	v1354 = v1198
	goto L318
L344:
	;
	goto L345
L345:
	;
	if base.Ui32(v1215) < base.Ui32(int32(4)) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1353 = l0
	v1354 = v1198
	goto L318
L347:
	;
	goto L348
L348:
	;
	v1334 = l0
	v1335 = v1198
	goto L349
L349:
	;
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1335))) = uint8(v1339)
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1335)+1)) = uint8(v1341)
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1335)+2)) = uint8(v1343)
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1335)+3)) = uint8(v1345)
	v1347 = int32(4)
	v1348 = v1334 + v1347
	v1350 = v1335 + v1347
	if base.Ui32(v1350) <= base.Ui32(v1222-int32(4)) {
		v1334 = v1348
		v1335 = v1350
		goto L349
	} else {
		goto L351
	}
L350:
	;
	v1353 = v1348
	v1354 = v1350
	goto L318
L351:
	;
	goto L350
L352:
	;
	v1360 = v1353
	v1361 = v1354
	goto L355
L353:
	;
	goto L354
L354:
	;
	goto L311
L355:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1361))) = uint8(v1365)
	v1367 = int32(1)
	v1370 = v1361 + v1367
	if v1370 != v1222 {
		v1360 = v1360 + v1367
		v1361 = v1370
		goto L355
	} else {
		goto L357
	}
L356:
	;
	goto L354
L357:
	;
	goto L356
}
