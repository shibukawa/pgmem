package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PGLC_localeconv(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int64
	_ = v244
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v394 int64
	_ = v394
	var v406 int32
	_ = v406
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v580 int32
	_ = v580
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v688 int32
	_ = v688
	var v701 int32
	_ = v701
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v733 int32
	_ = v733
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int64
	_ = v956
	var v972 int32
	_ = v972
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1046 int32
	_ = v1046
	var v1058 int32
	_ = v1058
	var v1071 int32
	_ = v1071
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1167 int32
	_ = v1167
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1207 int32
	_ = v1207
	var v1218 int32
	_ = v1218
	var v1229 int32
	_ = v1229
	var v1240 int32
	_ = v1240
	var v1251 int32
	_ = v1251
	var v1257 int64
	_ = v1257
	var v1260 int64
	_ = v1260
	var v1263 int64
	_ = v1263
	var v1266 int64
	_ = v1266
	var v1269 int64
	_ = v1269
	var v1272 int64
	_ = v1272
	var v1275 int64
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1357 int32
	_ = v1357
	var v1382 int32
	_ = v1382
	var v1383 int64
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	v1 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(320)
	m.G0 = v27
	v31 = v27 + int32(184)
	v33 = v27 + int32(188)
	v35 = v27 + int32(192)
	v37 = v27 + int32(196)
	v39 = v27 + int32(204)
	v41 = v27 + int32(208)
	v43 = v1
	v44 = v1
	v45 = v1
	v46 = v1
	v47 = v1
	v48 = v1
	v49 = v1
	v50 = v1
	v51 = v1
	v52 = v1
	v53 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v53 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L3
L6:
	;
	v1382 = int32(m.ExcTag)
	v1383 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1382 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[0])) = v1102
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[1])) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	v1326 = v27 + int32(172)
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	F_emscripten_builtin_free(m, v1327)
	mBase = m.M
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+4))
	F_emscripten_builtin_free(m, v1329)
	mBase = m.M
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+8))
	F_emscripten_builtin_free(m, v1331)
	mBase = m.M
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+12))
	F_emscripten_builtin_free(m, v1333)
	mBase = m.M
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+16))
	F_emscripten_builtin_free(m, v1335)
	mBase = m.M
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+20))
	F_emscripten_builtin_free(m, v1337)
	mBase = m.M
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+24))
	F_emscripten_builtin_free(m, v1339)
	mBase = m.M
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+28))
	F_emscripten_builtin_free(m, v1341)
	mBase = m.M
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+32))
	F_emscripten_builtin_free(m, v1343)
	mBase = m.M
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1326)+36))
	F_emscripten_builtin_free(m, v1345)
	mBase = m.M
	goto L241
L8:
	;
	m.G0 = v27 + int32(320)
	return int32(_a_F_PGLC_localeconv_0)
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+220)) = int64(0)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[2])))
	if v71 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v1099 = v43
	v1100 = v44
	v1101 = v45
	v1102 = v46
	v1103 = v47
	v1104 = v48
	v1105 = v49
	v1106 = v50
	v1107 = v51
	v1108 = v52
	goto L11
L11:
	;
	if v1099 != 0 {
		goto L7
	} else {
		goto L224
	}
L12:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[3])))
	if v73 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	v83 = int32(_a_F_PGLC_localeconv_0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[4]))
	F_emscripten_builtin_free(m, v84)
	mBase = m.M
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[5]))
	F_emscripten_builtin_free(m, v86)
	mBase = m.M
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[6]))
	F_emscripten_builtin_free(m, v88)
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[7]))
	F_emscripten_builtin_free(m, v90)
	mBase = m.M
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[8]))
	F_emscripten_builtin_free(m, v92)
	mBase = m.M
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[9]))
	F_emscripten_builtin_free(m, v94)
	mBase = m.M
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[10]))
	F_emscripten_builtin_free(m, v96)
	mBase = m.M
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[11]))
	F_emscripten_builtin_free(m, v98)
	mBase = m.M
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[12]))
	F_emscripten_builtin_free(m, v100)
	mBase = m.M
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[13]))
	F_emscripten_builtin_free(m, v102)
	mBase = m.M
	goto L16
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[14]))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[16])) = int32(44)
	v125 = int32(0)
	v130 = m.G0
	v132 = v130 - int32(32)
	m.G0 = v132
	v137 = v125
	goto L21
L16:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[3])) = uint8(v105)
	goto L15
L17:
	;
	if v688 != 0 {
		goto L159
	} else {
		goto L160
	}
L18:
	;
	if v260 == int32(0) {
		v688 = int32(-1)
		goto L17
	} else {
		goto L46
	}
L19:
	;
	m.G0 = v132 + int32(32)
	goto L18
L20:
	;
	v260 = int32(0)
	goto L19
L21:
	;
	v142 = v137 << (uint(int32(2)) % 32)
	v148 = int32(1) << (uint(v137) % 32) & int32(2147483647)
	if v148|int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v167 = F___loc_is_allocated(m, v125)
	mBase = m.M
	if v167 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142+(v132+int32(8))))) = v159
	if v159 == int32(-1) {
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v125+v142)))
	v159 = v155
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v148 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v157 = v117
	goto L29
L28:
	;
	v157 = int32(_a_F_PGLC_localeconv_1)
	goto L29
L29:
	;
	v158 = F___get_locale(m, v137, v157)
	mBase = m.M
	v159 = v158
	goto L23
L30:
	;
	v164 = v137 + int32(1)
	if v164 != int32(6) {
		v137 = v164
		goto L21
	} else {
		goto L31
	}
L31:
	;
	goto L22
L32:
	;
	v170 = int32(_a_F_PGLC_localeconv_2)
	v172 = v132 + int32(8)
	v175 = F_memcmp(m, v172, v170, int32(24))
	mBase = m.M
	if v175 == int32(0) {
		v260 = v170
		goto L19
	} else {
		goto L35
	}
L33:
	;
	v239 = v125
	goto L34
L34:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v132)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+16)) = v244
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v132)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v239)+8)) = v246
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v132)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = v248
	v260 = v239
	goto L19
L35:
	;
	v178 = int32(_a_F_PGLC_localeconv_3)
	v181 = F_memcmp(m, v172, v178, int32(24))
	mBase = m.M
	if v181 == int32(0) {
		v260 = v178
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v184 = int32(0)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[17])))
	if v186 == v184 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v192 = v184
	goto L40
L38:
	;
	goto L39
L39:
	;
	v219 = int32(_a_F_PGLC_localeconv_4)
	v221 = v132 + int32(8)
	v224 = F_memcmp(m, v221, v219, int32(24))
	mBase = m.M
	if v224 == int32(0) {
		v260 = v219
		goto L19
	} else {
		goto L43
	}
L40:
	;
	v199 = F___get_locale(m, v192, int32(_a_F_PGLC_localeconv_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v192<<(uint(int32(2))%32))+uint32(_c_F_PGLC_localeconv[18]))) = v199
	v202 = v192 + int32(1)
	if v202 != int32(6) {
		v192 = v202
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[17])) = uint8(v206)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[19])) = v210
	goto L39
L42:
	;
	goto L41
L43:
	;
	v227 = int32(_a_F_PGLC_localeconv_5)
	v230 = F_memcmp(m, v221, v227, int32(24))
	mBase = m.M
	if v230 == int32(0) {
		v260 = v227
		goto L19
	} else {
		goto L44
	}
L44:
	;
	v234 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v234 == int32(0) {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	v239 = v234
	goto L34
L46:
	;
	v271 = int32(0)
	v276 = m.G0
	v278 = v276 - int32(32)
	m.G0 = v278
	v283 = v271
	goto L50
L47:
	;
	if v406 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L48:
	;
	m.G0 = v278 + int32(32)
	goto L47
L49:
	;
	v406 = int32(0)
	goto L48
L50:
	;
	v288 = v283 << (uint(int32(2)) % 32)
	v294 = int32(1) << (uint(v283) % 32) & int32(2147483647)
	if v294|int32(1) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v313 = F___loc_is_allocated(m, v271)
	mBase = m.M
	if v313 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288+(v278+int32(8))))) = v305
	if v305 == int32(-1) {
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v271+v288)))
	v305 = v301
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v294 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v303 = v119
	goto L58
L57:
	;
	v303 = int32(_a_F_PGLC_localeconv_1)
	goto L58
L58:
	;
	v304 = F___get_locale(m, v283, v303)
	mBase = m.M
	v305 = v304
	goto L52
L59:
	;
	v310 = v283 + int32(1)
	if v310 != int32(6) {
		v283 = v310
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	v316 = int32(_a_F_PGLC_localeconv_2)
	v318 = v278 + int32(8)
	v321 = F_memcmp(m, v318, v316, int32(24))
	mBase = m.M
	if v321 == int32(0) {
		v406 = v316
		goto L48
	} else {
		goto L64
	}
L62:
	;
	v385 = v271
	goto L63
L63:
	;
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v278)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v385)+16)) = v390
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v278)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v385)+8)) = v392
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v278)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v394
	v406 = v385
	goto L48
L64:
	;
	v324 = int32(_a_F_PGLC_localeconv_3)
	v327 = F_memcmp(m, v318, v324, int32(24))
	mBase = m.M
	if v327 == int32(0) {
		v406 = v324
		goto L48
	} else {
		goto L65
	}
L65:
	;
	v330 = int32(0)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[17])))
	if v332 == v330 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v338 = v330
	goto L69
L67:
	;
	goto L68
L68:
	;
	v365 = int32(_a_F_PGLC_localeconv_4)
	v367 = v278 + int32(8)
	v370 = F_memcmp(m, v367, v365, int32(24))
	mBase = m.M
	if v370 == int32(0) {
		v406 = v365
		goto L48
	} else {
		goto L72
	}
L69:
	;
	v345 = F___get_locale(m, v338, int32(_a_F_PGLC_localeconv_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v338<<(uint(int32(2))%32))+uint32(_c_F_PGLC_localeconv[18]))) = v345
	v348 = v338 + int32(1)
	if v348 != int32(6) {
		v338 = v348
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v352 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[17])) = uint8(v352)
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[18]))
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[19])) = v356
	goto L68
L71:
	;
	goto L70
L72:
	;
	v373 = int32(_a_F_PGLC_localeconv_5)
	v376 = F_memcmp(m, v367, v373, int32(24))
	mBase = m.M
	if v376 == int32(0) {
		v406 = v373
		goto L48
	} else {
		goto L73
	}
L73:
	;
	v380 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v380 == int32(0) {
		goto L49
	} else {
		goto L74
	}
L74:
	;
	v385 = v380
	goto L63
L75:
	;
	v416 = F___loc_is_allocated(m, v260)
	mBase = m.M
	if v416 != 0 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L77
L77:
	;
	v420 = v27 + int32(228)
	v421 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v420)+48)) = v421
	*(*int64)(unsafe.Add(mBase, uint32(v420)+40)) = v421
	*(*int64)(unsafe.Add(mBase, uint32(v420)+32)) = v421
	*(*int64)(unsafe.Add(mBase, uint32(v420)+24)) = v421
	*(*int64)(unsafe.Add(mBase, uint32(v420)+16)) = v421
	*(*int64)(unsafe.Add(mBase, uint32(v420)+8)) = v421
	*(*int64)(unsafe.Add(mBase, uint32(v420))) = v421
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[20]))
	if v260 != 0 {
		goto L83
	} else {
		goto L84
	}
L78:
	;
	v688 = int32(-1)
	goto L17
L79:
	;
	F_emscripten_builtin_free(m, v260)
	mBase = m.M
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	v450 = int32(0)
	goto L94
L83:
	;
	if v260 == int32(-1) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	if v438 == int32(_a_F_PGLC_localeconv_6) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v443 = int32(_a_F_PGLC_localeconv_6)
	goto L88
L87:
	;
	v443 = v260
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[20])) = v443
	goto L85
L89:
	;
	v448 = int32(-1)
	goto L91
L90:
	;
	v448 = v438
	goto L91
L91:
	;
	goto L82
L92:
	;
	if v448 != 0 {
		goto L142
	} else {
		goto L143
	}
L93:
	;
	v580 = int32(0)
	goto L135
L94:
	;
	if base.Ui32(int32(14)) < base.Ui32(v450-int32(3)) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v406 != 0 {
		goto L110
	} else {
		goto L111
	}
L96:
	;
	v508 = v450 + int32(1)
	if v508 != int32(18) {
		v450 = v508
		goto L94
	} else {
		goto L108
	}
L97:
	;
	v478 = v450 * int32(12)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+uint32(_c_F_PGLC_localeconv[21])))
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+uint32(_c_F_PGLC_localeconv[22]))))
	if v482 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+uint32(_c_F_PGLC_localeconv[23])))
	v488 = F_strlen(m, v485)
	mBase = m.M
	v490 = v488 + int32(1)
	v491 = F_emscripten_builtin_malloc(m, v490)
	mBase = m.M
	if v491 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+uint32(_c_F_PGLC_localeconv[23]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v420+v479))) = uint8(v502)
	goto L96
L101:
	;
	if v496 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v496 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v495 = F___memcpy(m, v491, v485, v490)
	mBase = m.M
	v496 = v495
	goto L101
L105:
	;
	goto L93
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420+v479))) = v496
	goto L96
L108:
	;
	goto L95
L109:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[23]))
	v528 = F_strlen(m, v525)
	mBase = m.M
	v530 = v528 + int32(1)
	v531 = F_emscripten_builtin_malloc(m, v530)
	mBase = m.M
	if v531 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L110:
	;
	if v406 == int32(-1) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	goto L116
L113:
	;
	v518 = int32(_a_F_PGLC_localeconv_6)
	goto L115
L114:
	;
	v518 = v406
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[20])) = v518
	goto L112
L116:
	;
	goto L118
L118:
	;
	goto L109
L119:
	;
	goto L93
L120:
	;
	if v536 == int32(0) {
		goto L119
	} else {
		goto L124
	}
L121:
	;
	v536 = int32(0)
	goto L120
L122:
	;
	goto L123
L123:
	;
	v535 = F___memcpy(m, v531, v525, v530)
	mBase = m.M
	v536 = v535
	goto L120
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v536
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[24]))
	v544 = F_strlen(m, v541)
	mBase = m.M
	v546 = v544 + int32(1)
	v547 = F_emscripten_builtin_malloc(m, v546)
	mBase = m.M
	if v547 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v552 == int32(0) {
		goto L119
	} else {
		goto L129
	}
L126:
	;
	v552 = int32(0)
	goto L125
L127:
	;
	goto L128
L128:
	;
	v551 = F___memcpy(m, v547, v541, v546)
	mBase = m.M
	v552 = v551
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = v552
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[25]))
	v560 = F_strlen(m, v557)
	mBase = m.M
	v562 = v560 + int32(1)
	v563 = F_emscripten_builtin_malloc(m, v562)
	mBase = m.M
	if v563 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v568 == int32(0) {
		goto L119
	} else {
		goto L134
	}
L131:
	;
	v568 = int32(0)
	goto L130
L132:
	;
	goto L133
L133:
	;
	v567 = F___memcpy(m, v563, v557, v562)
	mBase = m.M
	v568 = v567
	goto L130
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v420)+8)) = v568
	v646 = int32(0)
	goto L92
L135:
	;
	v604 = v580 * int32(12)
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+uint32(_c_F_PGLC_localeconv[22]))))
	if v605 == int32(1) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[16])) = int32(48)
	v646 = int32(-1)
	goto L92
L137:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v604)+uint32(_c_F_PGLC_localeconv[21])))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v420+v610)))
	F_emscripten_builtin_free(m, v612)
	mBase = m.M
	goto L139
L138:
	;
	goto L139
L139:
	;
	v615 = v580 + int32(1)
	if v615 != int32(18) {
		v580 = v615
		goto L135
	} else {
		goto L140
	}
L140:
	;
	goto L136
L141:
	;
	v660 = F___loc_is_allocated(m, v260)
	mBase = m.M
	if v660 != 0 {
		goto L152
	} else {
		goto L153
	}
L142:
	;
	if v448 == int32(-1) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	goto L148
L145:
	;
	v654 = int32(_a_F_PGLC_localeconv_6)
	goto L147
L146:
	;
	v654 = v448
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[20])) = v654
	goto L144
L148:
	;
	goto L150
L150:
	;
	goto L141
L151:
	;
	v662 = F___loc_is_allocated(m, v406)
	mBase = m.M
	if v662 != 0 {
		goto L156
	} else {
		goto L157
	}
L152:
	;
	F_emscripten_builtin_free(m, v260)
	mBase = m.M
	goto L154
L153:
	;
	goto L154
L154:
	;
	goto L151
L155:
	;
	v688 = v646
	goto L17
L156:
	;
	F_emscripten_builtin_free(m, v406)
	mBase = m.M
	goto L158
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L6
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v27)+228))
	v746 = F_strlen(m, v743)
	mBase = m.M
	v748 = v746 + int32(1)
	v749 = F_emscripten_builtin_malloc(m, v748)
	mBase = m.M
	if v749 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[14]))
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v712
	F_errmsg_internal(m, int32(_a_F_PGLC_localeconv_7), v27)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	F_errfinish(m, int32(_a_F_PGLC_localeconv_8), int32(560), int32(_a_F_PGLC_localeconv_9))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	goto L3
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+172)) = v754
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v47
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v27)+232))
	v768 = F_strlen(m, v765)
	mBase = m.M
	v770 = v768 + int32(1)
	v771 = F_emscripten_builtin_malloc(m, v770)
	mBase = m.M
	if v771 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L166:
	;
	v754 = int32(0)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v753 = F___memcpy(m, v749, v743, v748)
	mBase = m.M
	v754 = v753
	goto L165
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+176)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	v787 = v27 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v27)+236))
	v792 = F_strlen(m, v789)
	mBase = m.M
	v794 = v792 + int32(1)
	v795 = F_emscripten_builtin_malloc(m, v794)
	mBase = m.M
	if v795 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L170:
	;
	v776 = int32(0)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v775 = F___memcpy(m, v771, v765, v770)
	mBase = m.M
	v776 = v775
	goto L169
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+180)) = v800
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v27)+240))
	v814 = F_strlen(m, v811)
	mBase = m.M
	v816 = v814 + int32(1)
	v817 = F_emscripten_builtin_malloc(m, v816)
	mBase = m.M
	if v817 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v800 = int32(0)
	goto L173
L175:
	;
	goto L176
L176:
	;
	v799 = F___memcpy(m, v795, v789, v794)
	mBase = m.M
	v800 = v799
	goto L173
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+184)) = v822
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v27)+244))
	v836 = F_strlen(m, v833)
	mBase = m.M
	v838 = v836 + int32(1)
	v839 = F_emscripten_builtin_malloc(m, v838)
	mBase = m.M
	if v839 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v822 = int32(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v821 = F___memcpy(m, v817, v811, v816)
	mBase = m.M
	v822 = v821
	goto L177
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+188)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v27)+248))
	v858 = F_strlen(m, v855)
	mBase = m.M
	v860 = v858 + int32(1)
	v861 = F_emscripten_builtin_malloc(m, v860)
	mBase = m.M
	if v861 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L182:
	;
	v844 = int32(0)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v843 = F___memcpy(m, v839, v833, v838)
	mBase = m.M
	v844 = v843
	goto L181
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v27)+252))
	v880 = F_strlen(m, v877)
	mBase = m.M
	v882 = v880 + int32(1)
	v883 = F_emscripten_builtin_malloc(m, v882)
	mBase = m.M
	if v883 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v866 = int32(0)
	goto L185
L187:
	;
	goto L188
L188:
	;
	v865 = F___memcpy(m, v861, v855, v860)
	mBase = m.M
	v866 = v865
	goto L185
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+196)) = v888
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v27)+256))
	v902 = F_strlen(m, v899)
	mBase = m.M
	v904 = v902 + int32(1)
	v905 = F_emscripten_builtin_malloc(m, v904)
	mBase = m.M
	if v905 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L190:
	;
	v888 = int32(0)
	goto L189
L191:
	;
	goto L192
L192:
	;
	v887 = F___memcpy(m, v883, v877, v882)
	mBase = m.M
	v888 = v887
	goto L189
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+200)) = v910
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v27)+260))
	v924 = F_strlen(m, v921)
	mBase = m.M
	v926 = v924 + int32(1)
	v927 = F_emscripten_builtin_malloc(m, v926)
	mBase = m.M
	if v927 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L194:
	;
	v910 = int32(0)
	goto L193
L195:
	;
	goto L196
L196:
	;
	v909 = F___memcpy(m, v905, v899, v904)
	mBase = m.M
	v910 = v909
	goto L193
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+204)) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v39
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v27)+264))
	v946 = F_strlen(m, v943)
	mBase = m.M
	v948 = v946 + int32(1)
	v949 = F_emscripten_builtin_malloc(m, v948)
	mBase = m.M
	if v949 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L198:
	;
	v932 = int32(0)
	goto L197
L199:
	;
	goto L200
L200:
	;
	v931 = F___memcpy(m, v927, v921, v926)
	mBase = m.M
	v932 = v931
	goto L197
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+208)) = v954
	v956 = *(*int64)(unsafe.Add(mBase, uint32(v27)+268))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+212)) = v956
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v41
	v972 = int32(0)
	goto L205
L202:
	;
	v954 = int32(0)
	goto L201
L203:
	;
	goto L204
L204:
	;
	v953 = F___memcpy(m, v949, v943, v948)
	mBase = m.M
	v954 = v953
	goto L201
L205:
	;
	v995 = v972 * int32(12)
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v995)+uint32(_c_F_PGLC_localeconv[22]))))
	if v996 == int32(1) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v1009 = int32(0)
	if base.B2i32(v754 == v1009)|base.B2i32(v776 == v1009)|(base.B2i32(v800 == v1009)|base.B2i32(v822 == v1009))|(base.B2i32(v844 == v1009)|base.B2i32(v866 == v1009)|(base.B2i32(v888 == v1009)|base.B2i32(v910 == v1009))) != 0 {
		goto L212
	} else {
		goto L213
	}
L207:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v995)+uint32(_c_F_PGLC_localeconv[21])))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(228)+v1001)))
	F_emscripten_builtin_free(m, v1003)
	mBase = m.M
	goto L209
L208:
	;
	goto L209
L209:
	;
	v1006 = v972 + int32(1)
	if v1006 != int32(18) {
		v972 = v1006
		goto L205
	} else {
		goto L210
	}
L210:
	;
	goto L206
L211:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[0]))
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[1]))
	goto L220
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L6
	} else {
		goto L216
	}
L213:
	;
	if v932 == int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	if v954 != 0 {
		goto L211
	} else {
		goto L215
	}
L215:
	;
	goto L212
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	F_errcode(m, int32(_a_F_PGLC_localeconv_10))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L6
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	F_errmsg(m, int32(_a_F_PGLC_localeconv_11), int32(0))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L6
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v787
	F_errfinish(m, int32(_a_F_PGLC_localeconv_8), int32(591), int32(_a_F_PGLC_localeconv_9))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L6
	} else {
		goto L219
	}
L219:
	;
	goto L3
L220:
	;
	v1092 = v27 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1092))) = v27 + int32(12)
	goto L223
L221:
	;
	v1099 = int32(0)
	v1100 = v41
	v1101 = v1090
	v1102 = v1088
	v1103 = v787
	v1104 = v39
	v1105 = v37
	v1106 = v35
	v1107 = v33
	v1108 = v31
	goto L11
L223:
	;
	goto L221
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[1])) = v27 + int32(16)
	v1136 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[15]))
	v1138 = F_pg_get_encoding_from_locale(m, v1136, int32(1))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	v1149 = int32(0)
	if v1149 < v1138 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1152 = v1138
	goto L228
L227:
	;
	v1152 = v1149
	goto L228
L228:
	;
	F_db_encoding_convert(m, v1152, v27+int32(172))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_db_encoding_convert(m, v1152, v1103)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	v1178 = *(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[14]))
	v1180 = F_pg_get_encoding_from_locale(m, v1178, int32(1))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L6
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	v1191 = int32(0)
	if v1191 < v1180 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1194 = v1180
	goto L234
L233:
	;
	v1194 = v1191
	goto L234
L234:
	;
	F_db_encoding_convert(m, v1194, v1108)
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L6
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_db_encoding_convert(m, v1194, v1107)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L6
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_db_encoding_convert(m, v1194, v1106)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L6
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_db_encoding_convert(m, v1194, v1105)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L6
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_db_encoding_convert(m, v1194, v1104)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L6
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_db_encoding_convert(m, v1194, v1100)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[0])) = v1102
	*(*int32)(unsafe.Add(mBase, _c_F_PGLC_localeconv[1])) = v1101
	v1257 = *(*int64)(unsafe.Add(mBase, uint32(v27)+172))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[4])) = v1257
	v1260 = *(*int64)(unsafe.Add(mBase, uint32(v27)+180))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[6])) = v1260
	v1263 = *(*int64)(unsafe.Add(mBase, uint32(v27)+188))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[8])) = v1263
	v1266 = *(*int64)(unsafe.Add(mBase, uint32(v27)+196))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[10])) = v1266
	v1269 = *(*int64)(unsafe.Add(mBase, uint32(v27)+204))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[12])) = v1269
	v1272 = *(*int64)(unsafe.Add(mBase, uint32(v27)+212))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[26])) = v1272
	v1275 = *(*int64)(unsafe.Add(mBase, uint32(v27)+220))
	*(*int64)(unsafe.Add(mBase, _c_F_PGLC_localeconv[27])) = v1275
	v1278 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[3])) = uint8(v1278)
	*(*uint8)(unsafe.Add(mBase, _c_F_PGLC_localeconv[2])) = uint8(v1278)
	goto L8
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+288)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v27)+284)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(v27)+292)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v27)+296)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v27)+300)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v27)+304)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v27)+308)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v27)+316)) = v1103
	F_pg_re_throw(m)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L6
	} else {
		goto L242
	}
L242:
	;
	goto L5
L243:
	;
	v1387 = int32(v1383)
	m.G0 = v27
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1387)+4))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1390)))
	if v27+int32(12) == v1393 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	m.ExcPending = 1
	goto L252
L245:
	;
	if v1397 != 0 {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+4))
	v1397 = v1395
	goto L248
L247:
	;
	v1397 = int32(0)
	goto L248
L248:
	;
	goto L245
L249:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v27)+316))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v27)+312))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v27)+308))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v27)+304))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v27)+300))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v27)+296))
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v27)+292))
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v27)+288))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v27)+284))
	v43 = v1389
	v44 = v1404
	v45 = v1405
	v46 = v1406
	v47 = v1398
	v48 = v1403
	v49 = v1402
	v50 = v1401
	v51 = v1400
	v52 = v1399
	v53 = v1397
	goto L1
L250:
	;
	goto L251
L251:
	;
	F___wasm_longjmp(m, v1390, v1389)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	return int32(0)
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ParseCommitRecord(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int64
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	v4 = int32(0)
	base.MemoryFill(m, l2, v4, int32(288))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v10
	if v4 <= base.I32_extend8_s(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v15
	if v15&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v21
	v27 = l1 + int32(20)
	goto L5
L4:
	;
	v27 = l1 + int32(12)
	goto L5
L5:
	;
	if v15&int32(2) != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v32 = v27 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v30
	v38 = v32 + v30<<(uint(int32(2))%32)
	goto L8
L7:
	;
	v38 = v27
	goto L8
L8:
	;
	if v15&int32(4) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v44 = v38 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v42
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v51 = v44 + v47*int32(12)
	goto L11
L10:
	;
	v51 = v38
	goto L11
L11:
	;
	if v15&int32(256) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v57 = int32(4)
	v58 = v51 + v57
	*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v56
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v65 = v58 + v61<<(uint(v57)%32)
	goto L14
L13:
	;
	v65 = v51
	goto L14
L14:
	;
	if v15&int32(8) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v71 = int32(4)
	v72 = v65 + v71
	*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l2)+44)) = v70
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v79 = v72 + v75<<(uint(v71)%32)
	goto L17
L16:
	;
	v79 = v65
	goto L17
L17:
	;
	if v15&int32(16) == int32(0) {
		v220 = v15
		v221 = v79
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v220&int32(32) == int32(0) {
		goto L1
	} else {
		goto L52
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+52)) = v86
	v89 = v79 + int32(4)
	if v15&int32(128) == int32(0) {
		v220 = v15
		v221 = v89
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v95 = l2 + int32(56)
	goto L24
L21:
	;
	v215 = F_strlen(m, v89)
	mBase = m.M
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v220 = v219
	v221 = v215 + v89 + int32(1)
	goto L18
L22:
	;
	v212 = F_strlen(m, v201)
	mBase = m.M
	goto L21
L24:
	;
	goto L25
L25:
	;
	v102 = int32(199)
	if (v95^v89)&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v205)
	goto L22
L27:
	;
	v186 = v181
	v187 = v182
	v188 = v183
	goto L48
L28:
	;
	if v176 == int32(0) {
		v201 = v174
		v202 = v175
		goto L26
	} else {
		goto L47
	}
L29:
	;
	v174 = v89
	v175 = v95
	v176 = v102
	goto L28
L30:
	;
	goto L31
L31:
	;
	v106 = int32(0)
	if base.B2i32(v89&int32(3) == v106)|int32(0) == v106 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v142 == int32(0) {
		v201 = v139
		v202 = v140
		goto L26
	} else {
		goto L41
	}
L33:
	;
	v118 = v89
	v119 = v95
	v120 = v102
	goto L36
L34:
	;
	goto L35
L35:
	;
	v139 = v89
	v140 = v95
	v141 = v102
	v142 = int32(1)
	goto L32
L36:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v122)
	if v122 == int32(0) {
		v181 = v118
		v182 = v119
		v183 = v120
		goto L27
	} else {
		goto L38
	}
L37:
	;
	v139 = v133
	v140 = v127
	v141 = v129
	v142 = v131
	goto L32
L38:
	;
	v126 = int32(1)
	v127 = v119 + v126
	v129 = v120 - v126
	v130 = int32(0)
	v131 = base.B2i32(v129 != v130)
	v133 = v118 + v126
	if v133&int32(3) == v130 {
		v139 = v133
		v140 = v127
		v141 = v129
		v142 = v131
		goto L32
	} else {
		goto L39
	}
L39:
	;
	if v129 != 0 {
		v118 = v133
		v119 = v127
		v120 = v129
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if base.B2i32(v145 == int32(0))|base.B2i32(base.Ui32(v141) < base.Ui32(int32(4))) != 0 {
		v174 = v139
		v175 = v140
		v176 = v141
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v152 = v139
	v153 = v140
	v154 = v141
	goto L43
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v160 = int32(-2139062144)
	if (int32(16843008)-v157|v157)&v160 != v160 {
		v181 = v152
		v182 = v153
		v183 = v154
		goto L27
	} else {
		goto L45
	}
L44:
	;
	v174 = v168
	v175 = v166
	v176 = v170
	goto L28
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v157
	v165 = int32(4)
	v166 = v153 + v165
	v168 = v152 + v165
	v170 = v154 - v165
	if base.Ui32(int32(3)) < base.Ui32(v170) {
		v152 = v168
		v153 = v166
		v154 = v170
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v181 = v174
	v182 = v175
	v183 = v176
	goto L27
L48:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v190)
	if v190 == int32(0) {
		v201 = v186
		v202 = v187
		goto L26
	} else {
		goto L50
	}
L49:
	;
	v201 = v197
	v202 = v195
	goto L26
L50:
	;
	v194 = int32(1)
	v195 = v187 + v194
	v197 = v186 + v194
	v199 = v188 - v194
	if v199 != 0 {
		v186 = v197
		v187 = v195
		v188 = v199
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v221)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+280)) = v227
	*(*int64)(unsafe.Add(mBase, uint32(l2)+272)) = v226
	goto L1
}
func F_ParseTzFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v600 int32
	_ = v600
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int64
	_ = v915
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v955 int32
	_ = v955
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v1002 int32
	_ = v1002
	var v1024 int32
	_ = v1024
	var v1040 int32
	_ = v1040
	var v1051 int32
	_ = v1051
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(3328)
	m.G0 = v23
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v23 + int32(3328)
	return v1073
L2:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L14
	} else {
		goto L15
	}
L3:
	;
	v34 = v25
	v35 = l0
	goto L4
L4:
	;
	if base.Ui32((v34|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v59 = int32(-1)
	if l1 == int32(0) {
		v1073 = v59
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v56 != 0 {
		v34 = v56
		v35 = v35 + int32(1)
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v63
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+224)) = l0
	v71 = F_format_elog_string(m, int32(_a_F_ParseTzFile_0), v23+int32(224))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v71
	v1073 = v59
	goto L1
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v99
	goto L17
L15:
	;
	goto L16
L16:
	;
	v110 = v23 + int32(2288)
	F_get_share_path(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l0
	v105 = F_format_elog_string(m, int32(_a_F_ParseTzFile_1), v23)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v105
	v1073 = int32(-1)
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+212)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = v110
	v116 = v23 + int32(1264)
	v121 = F_pg_snprintf(m, v116, int32(1024), int32(_a_F_ParseTzFile_2), v23+int32(208))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v124 = F_AllocateFile(m, v116, int32(_a_F_ParseTzFile_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L23
	}
L21:
	;
	v1067 = F_FreeFile(m, v124)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L12
	} else {
		goto L284
	}
L22:
	;
	v200 = l4
	v204 = v6
	v213 = v6
	goto L42
L23:
	;
	if v124 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L27
L25:
	;
	goto L26
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v23 + int32(2288)
	v139 = v23 + int32(1264)
	v144 = F_pg_snprintf(m, v139, int32(1024), int32(_a_F_ParseTzFile_4), v23-int32(-64))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L29
	}
L27:
	;
	if int32(base.Ui32(v126)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1051 = l4
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	v146 = F_AllocateDir(m, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	if v146 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v151
	goto L34
L32:
	;
	goto L33
L33:
	;
	F_FreeDir(m, v146)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v139
	v159 = F_format_elog_string(m, int32(_a_F_ParseTzFile_5), v23+int32(32))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v159
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v163
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_ParseTzFile_6)
	v172 = F_format_elog_string(m, int32(_a_F_ParseTzFile_7), v23+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[3])) = v172
	v1073 = int32(-1)
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0])) = v134
	v180 = int32(-1)
	if base.B2i32(l1 == int32(0))&base.B2i32(v134 == int32(44)) != 0 {
		v1073 = v180
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v134
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = l0
	v193 = F_format_elog_string(m, int32(_a_F_ParseTzFile_8), v23+int32(48))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v193
	v1073 = v180
	goto L1
L42:
	;
	v219 = F_fgets(m, v23+int32(240), int32(1024), v124)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L48
	}
L43:
	;
	v1051 = v1024
	goto L21
L44:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L282
L45:
	;
	if v1002 < int32(0) {
		v1051 = v1002
		goto L21
	} else {
		goto L281
	}
L46:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v969 <= v200 {
		goto L274
	} else {
		goto L275
	}
L47:
	;
	v1051 = int32(-1)
	goto L21
L48:
	;
	if v219 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L52
L50:
	;
	goto L51
L51:
	;
	v243 = v204 + int32(1)
	v245 = v23 + int32(240)
	v246 = F_strlen(m, v245)
	mBase = m.M
	if v246 == int32(1023) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	if int32(base.Ui32(v223)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1051 = v200
		goto L21
	} else {
		goto L53
	}
L53:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v231
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = l0
	v239 = F_format_elog_string(m, int32(_a_F_ParseTzFile_8), v23+int32(80))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v239
	goto L47
L56:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v250
	goto L59
L57:
	;
	goto L58
L58:
	;
	v268 = v245
	goto L63
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = l0
	v259 = F_format_elog_string(m, int32(_a_F_ParseTzFile_9), v23+int32(96))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v259
	goto L47
L61:
	;
	v300 = v268
	v301 = int32(_a_F_ParseTzFile_10)
	v302 = int32(8)
	goto L72
L62:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L69
L63:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if base.Ui32(v282-int32(9)) < base.Ui32(int32(5)) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v282 != 0 {
		goto L61
	} else {
		goto L68
	}
L65:
	;
	goto L64
L66:
	;
	v268 = v268 + int32(1)
	goto L63
L67:
	;
	switch v282 - int32(32) {
	case 0:
		goto L66
	case 1, 2:
		goto L61
	case 3:
		goto L62
	default:
		goto L65
	}
L68:
	;
	goto L62
L69:
	;
	if int32(base.Ui32(v291)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1051 = v200
		goto L21
	} else {
		goto L70
	}
L70:
	;
	v204 = v243
	goto L42
L71:
	;
	if v347 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	if v302 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v347 = int32(0)
	goto L71
L74:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v305 == v306 {
		v328 = v305
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	v330 = int32(1)
	if v328 != 0 {
		v300 = v300 + v330
		v301 = v301 + v330
		v302 = v302 - v330
		goto L72
	} else {
		goto L86
	}
L78:
	;
	if base.Ui32((v305-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v316 = v305 | int32(32)
	goto L81
L80:
	;
	v316 = v305
	goto L81
L81:
	;
	if base.Ui32((v306-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v325 = v306 | int32(32)
	goto L84
L83:
	;
	v325 = v306
	goto L84
L84:
	;
	if v316 == v325 {
		v328 = v316
		goto L77
	} else {
		goto L85
	}
L85:
	;
	v347 = v316 - v325
	goto L71
L86:
	;
	goto L76
L87:
	;
	v352 = F_pstrdup(m, v268+int32(8))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L12
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v405 = v268
	v406 = int32(_a_F_ParseTzFile_11)
	v407 = int32(9)
	goto L111
L90:
	;
	v397 = F_ParseTzFile(m, v382, l1+int32(1), l2, l3, v200)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L12
	} else {
		goto L108
	}
L91:
	;
	v355 = v23 + int32(3324)
	if v352 != 0 {
		v359 = v352
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v382 != 0 {
		goto L102
	} else {
		goto L103
	}
L93:
	;
	v361 = F_strspn(m, v359, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v362 = v361 + v359
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if v363 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	if v357 != 0 {
		v359 = v357
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v382 = int32(0)
	goto L92
L96:
	;
	v366 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v366
	v382 = v366
	goto L92
L97:
	;
	goto L98
L98:
	;
	v370 = F_strcspn(m, v362, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v371 = v370 + v362
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	if v372 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v371 + int32(1)
	v376 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v376)
	v382 = v362
	goto L92
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = int32(0)
	v382 = v362
	goto L92
L102:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v383 != 0 {
		goto L90
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v385
	goto L106
L105:
	;
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = l0
	v394 = F_format_elog_string(m, int32(_a_F_ParseTzFile_13), v23+int32(112))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v394
	goto L47
L108:
	;
	if int32(0) <= v397 {
		v1024 = v397
		goto L44
	} else {
		goto L109
	}
L109:
	;
	v1051 = v397
	goto L21
L110:
	;
	if v452 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L111:
	;
	if v407 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v452 = int32(0)
	goto L110
L113:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v410 == v411 {
		v433 = v410
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	v435 = int32(1)
	if v433 != 0 {
		v405 = v405 + v435
		v406 = v406 + v435
		v407 = v407 - v435
		goto L111
	} else {
		goto L125
	}
L117:
	;
	if base.Ui32((v410-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v421 = v410 | int32(32)
	goto L120
L119:
	;
	v421 = v410
	goto L120
L120:
	;
	if base.Ui32((v411-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v430 = v411 | int32(32)
	goto L123
L122:
	;
	v430 = v411
	goto L123
L123:
	;
	if v421 == v430 {
		v433 = v421
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v452 = v421 - v430
	goto L110
L125:
	;
	goto L115
L126:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	goto L129
L127:
	;
	goto L128
L128:
	;
	v462 = v23 + int32(3324)
	if v268 != 0 {
		v466 = v268
		goto L134
	} else {
		goto L135
	}
L129:
	;
	if int32(base.Ui32(v455)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1051 = v200
		goto L21
	} else {
		goto L130
	}
L130:
	;
	v204 = v243
	v213 = int32(1)
	goto L42
L131:
	;
	v695 = F_strlen(m, v493)
	mBase = m.M
	if base.Ui32(int32(11)) <= base.Ui32(v695) {
		goto L212
	} else {
		goto L213
	}
L132:
	;
	v679 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v679
	goto L208
L133:
	;
	if v489 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L134:
	;
	v468 = F_strspn(m, v466, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v469 = v468 + v466
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	if v470 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	if v464 != 0 {
		v466 = v464
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v489 = int32(0)
	goto L133
L137:
	;
	v473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v473
	v489 = v473
	goto L133
L138:
	;
	goto L139
L139:
	;
	v477 = F_strcspn(m, v469, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v478 = v477 + v469
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v479 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v478 + int32(1)
	v483 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v478))) = uint8(v483)
	v489 = v469
	goto L133
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = int32(0)
	v489 = v469
	goto L133
L143:
	;
	v673 = int32(_a_F_ParseTzFile_14)
	goto L132
L144:
	;
	goto L145
L145:
	;
	v493 = F_pstrdup(m, v489)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	v497 = v23 + int32(3324)
	goto L149
L147:
	;
	if v524 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L148:
	;
	v503 = F_strspn(m, v499, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v504 = v503 + v499
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	if v505 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	if v499 != 0 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v524 = int32(0)
	goto L147
L151:
	;
	v508 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v508
	v524 = v508
	goto L147
L152:
	;
	goto L153
L153:
	;
	v512 = F_strcspn(m, v504, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v513 = v512 + v504
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	if v514 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v513 + int32(1)
	v518 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v513))) = uint8(v518)
	v524 = v504
	goto L147
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = int32(0)
	v524 = v504
	goto L147
L157:
	;
	v673 = int32(_a_F_ParseTzFile_15)
	goto L132
L158:
	;
	goto L159
L159:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	if base.Ui32((v528-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666))))
	if v669 == int32(35) {
		v690 = v665
		v692 = v667
		v693 = v668
		goto L131
	} else {
		goto L207
	}
L161:
	;
	v635 = v23 + int32(3324)
	goto L198
L162:
	;
	v626 = F_pstrdup(m, v524)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L12
	} else {
		goto L195
	}
L163:
	;
	v541 = F_strtox_2(m, v524, v23+int32(3320), int32(10), int64(2147483648))
	mBase = m.M
	v542 = base.I32_wrap_i64(v541)
	goto L165
L164:
	;
	switch v528 - int32(43) {
	case 0, 2:
		goto L163
	default:
		goto L162
	}
L165:
	;
	v543 = int32(_a_F_ParseTzFile_16)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v23)+3320))
	if v524 == v544 {
		v673 = v543
		goto L132
	} else {
		goto L166
	}
L166:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	if v546 != 0 {
		v673 = v543
		goto L132
	} else {
		goto L167
	}
L167:
	;
	v547 = int32(0)
	v550 = v23 + int32(3324)
	goto L170
L168:
	;
	if v577 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L169:
	;
	v556 = F_strspn(m, v552, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v557 = v556 + v552
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v558 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	if v552 != 0 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v577 = int32(0)
	goto L168
L172:
	;
	v561 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v561
	v577 = v561
	goto L168
L173:
	;
	goto L174
L174:
	;
	v565 = F_strcspn(m, v557, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v566 = v565 + v557
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	if v567 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v566 + int32(1)
	v571 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v566))) = uint8(v571)
	v577 = v557
	goto L168
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = int32(0)
	v577 = v557
	goto L168
L178:
	;
	v690 = int32(0)
	v692 = v542
	v693 = v547
	goto L131
L179:
	;
	goto L180
L180:
	;
	v581 = int32(0)
	v585 = v577
	v586 = int32(_a_F_ParseTzFile_17)
	goto L182
L181:
	;
	if v623 != 0 {
		v665 = v581
		v666 = v577
		v667 = v542
		v668 = v547
		goto L160
	} else {
		goto L194
	}
L182:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	if v589 == v590 {
		v612 = v589
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v623 = int32(0)
	goto L181
L184:
	;
	v614 = int32(1)
	if v612 != 0 {
		v585 = v585 + v614
		v586 = v586 + v614
		goto L182
	} else {
		goto L193
	}
L185:
	;
	if base.Ui32((v589-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v600 = v589 | int32(32)
	goto L188
L187:
	;
	v600 = v589
	goto L188
L188:
	;
	if base.Ui32((v590-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v609 = v590 | int32(32)
	goto L191
L190:
	;
	v609 = v590
	goto L191
L191:
	;
	if v600 == v609 {
		v612 = v600
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v623 = v600 - v609
	goto L181
L193:
	;
	goto L183
L194:
	;
	v629 = v581
	v631 = v542
	v632 = int32(1)
	goto L161
L195:
	;
	v629 = v626
	v631 = int32(0)
	v632 = int32(0)
	goto L161
L196:
	;
	if v662 == int32(0) {
		v690 = v629
		v692 = v631
		v693 = v632
		goto L131
	} else {
		goto L206
	}
L197:
	;
	v641 = F_strspn(m, v637, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v642 = v641 + v637
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
	if v643 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	if v637 != 0 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v662 = int32(0)
	goto L196
L200:
	;
	v646 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = v646
	v662 = v646
	goto L196
L201:
	;
	goto L202
L202:
	;
	v650 = F_strcspn(m, v642, int32(_a_F_ParseTzFile_12))
	mBase = m.M
	v651 = v650 + v642
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	if v652 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = v651 + int32(1)
	v656 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v651))) = uint8(v656)
	v662 = v642
	goto L196
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = int32(0)
	v662 = v642
	goto L196
L206:
	;
	v665 = v629
	v666 = v662
	v667 = v631
	v668 = v632
	goto L160
L207:
	;
	v673 = int32(_a_F_ParseTzFile_18)
	goto L132
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = l0
	v687 = F_format_elog_string(m, v673, v23+int32(128))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L12
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v687
	goto L47
L210:
	;
	v798 = v772
	v799 = v775
	goto L233
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v790
	goto L47
L212:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v699
	goto L215
L213:
	;
	goto L214
L214:
	;
	if base.Ui32(int32(-100801)) <= base.Ui32(v692-int32(_a_F_ParseTzFile_19)) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v493
	v710 = F_format_elog_string(m, int32(_a_F_ParseTzFile_20), v23+int32(144))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L12
	} else {
		goto L216
	}
L216:
	;
	v790 = v710
	goto L211
L217:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	if v716 != 0 {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L219
L219:
	;
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v779
	goto L231
L220:
	;
	v723 = v493
	v724 = v716
	goto L223
L221:
	;
	goto L222
L222:
	;
	v772 = int32(0)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v775 = v200 - int32(1)
	if v772 <= v775 {
		goto L210
	} else {
		goto L230
	}
L223:
	;
	v737 = int32(255)
	v738 = v724 & v737
	if base.Ui32((v738-int32(65))&v737) < base.Ui32(int32(26)) {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L222
L225:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v723))) = uint8(v747)
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+1)))
	if v749 != 0 {
		v723 = v723 + int32(1)
		v724 = v749
		goto L223
	} else {
		goto L229
	}
L226:
	;
	v747 = v738 | int32(32)
	goto L228
L227:
	;
	v747 = v738
	goto L228
L228:
	;
	goto L225
L229:
	;
	goto L224
L230:
	;
	v955 = v772
	goto L46
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = v692
	v788 = F_format_elog_string(m, int32(_a_F_ParseTzFile_21), v23+int32(160))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L12
	} else {
		goto L232
	}
L232:
	;
	v790 = v788
	goto L211
L233:
	;
	v814 = (v798 + v799) >> (uint(int32(1)) % 32)
	v817 = v773 + v814*int32(24)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493))))
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818))))
	if base.B2i32(v821 == int32(0))|base.B2i32(v821 != v824) != 0 {
		v842 = v821
		v843 = v824
		goto L238
	} else {
		goto L239
	}
L234:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	if v856 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L235:
	;
	goto L234
L236:
	;
	if v853 <= v854 {
		v798 = v853
		v799 = v854
		goto L233
	} else {
		goto L248
	}
L237:
	;
	if v844 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L238:
	;
	v844 = v842 - v843
	goto L237
L239:
	;
	v827 = v493
	v828 = v818
	goto L240
L240:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828)+1)))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827)+1)))
	if v832 == int32(0) {
		v842 = v832
		v843 = v831
		goto L238
	} else {
		goto L242
	}
L241:
	;
	v842 = v832
	v843 = v831
	goto L238
L242:
	;
	v835 = int32(1)
	if v832 == v831 {
		v827 = v827 + v835
		v828 = v828 + v835
		goto L240
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	v853 = v798
	v854 = v814 - int32(1)
	goto L236
L245:
	;
	goto L246
L246:
	;
	if v844 == int32(0) {
		goto L235
	} else {
		goto L247
	}
L247:
	;
	v853 = v814 + int32(1)
	v854 = v799
	goto L236
L248:
	;
	v955 = v853
	goto L46
L249:
	;
	if v213 != 0 {
		goto L267
	} else {
		goto L268
	}
L250:
	;
	if v690 != 0 {
		v895 = v690
		goto L249
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	if v690 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	v859 = int32(0)
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v817)+8))
	if v860 != v692 {
		v895 = v859
		goto L249
	} else {
		goto L254
	}
L254:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+12)))
	if v862 != v693 {
		v895 = v859
		goto L249
	} else {
		goto L255
	}
L255:
	;
	v1002 = v200
	goto L45
L256:
	;
	v895 = int32(0)
	goto L249
L257:
	;
	goto L258
L258:
	;
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	if base.B2i32(v869 == int32(0))|base.B2i32(v869 != v872) != 0 {
		v890 = v869
		v891 = v872
		goto L260
	} else {
		goto L261
	}
L259:
	;
	if v890-v891 == int32(0) {
		v1002 = v200
		goto L45
	} else {
		goto L266
	}
L260:
	;
	goto L259
L261:
	;
	v875 = v856
	v876 = v690
	goto L262
L262:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876)+1)))
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875)+1)))
	if v880 == int32(0) {
		v890 = v880
		v891 = v879
		goto L260
	} else {
		goto L264
	}
L263:
	;
	v890 = v880
	v891 = v879
	goto L260
L264:
	;
	v883 = int32(1)
	if v880 == v879 {
		v875 = v875 + v883
		v876 = v876 + v883
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v895 = v690
	goto L249
L267:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v817)+12)) = uint8(v693)
	*(*int32)(unsafe.Add(mBase, uint32(v817)+8)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v817)+4)) = v895
	v1002 = v200
	goto L45
L268:
	;
	goto L269
L269:
	;
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v900
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v493
	v908 = F_format_elog_string(m, int32(_a_F_ParseTzFile_22), v23+int32(192))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L12
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[2])) = v908
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[1])) = v912
	goto L272
L272:
	;
	v915 = *(*int64)(unsafe.Add(mBase, uint32(v817)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v23)+176)) = base.I64_rotl(v915, int64(32))
	v925 = F_format_elog_string(m, int32(_a_F_ParseTzFile_23), v23+int32(176))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L12
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseTzFile[4])) = v925
	goto L47
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v969 << (uint(int32(1)) % 32)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v977 = F_repalloc(m, v974, v969*int32(48))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L12
	} else {
		goto L277
	}
L275:
	;
	v980 = v773
	goto L276
L276:
	;
	v981 = int32(24)
	v983 = v980 + v955*v981
	v986 = (v200 - v955) * v981
	if v986 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v977
	v980 = v977
	goto L276
L278:
	;
	base.MemoryCopy(m, v983+int32(24), v983, v986)
	goto L280
L279:
	;
	goto L280
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v983)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v983)+16)) = v243
	*(*uint8)(unsafe.Add(mBase, uint32(v983)+12)) = uint8(v693)
	*(*int32)(unsafe.Add(mBase, uint32(v983)+8)) = v692
	*(*int32)(unsafe.Add(mBase, uint32(v983)+4)) = v690
	*(*int32)(unsafe.Add(mBase, uint32(v983))) = v493
	v1002 = v200 + int32(1)
	goto L45
L281:
	;
	v1024 = v1002
	goto L44
L282:
	;
	if int32(base.Ui32(v1040)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v200 = v1024
		v204 = v243
		goto L42
	} else {
		goto L283
	}
L283:
	;
	goto L43
L284:
	;
	v1073 = v1051
	goto L1
}
func F_PersistHoldablePortal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int64
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v330 int32
	_ = v330
	var v343 int32
	_ = v343
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v389 int32
	_ = v389
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v444 int32
	_ = v444
	var v463 int32
	_ = v463
	var v464 int64
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(208)
	m.G0 = v21
	v25 = l0 + int32(108)
	v27 = l0 + int32(88)
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v37 = v2
	v38 = int32(-1)
	v39 = v2
	v40 = v2
	goto L1
L1:
	;
	goto L3
L2:
	;
	m.G0 = v21 + int32(208)
	return
L3:
	;
	if v38 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v463 = int32(m.ExcTag)
	v464 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v463 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v49 = int32(_a_F_PersistHoldablePortal_0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[0]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[0])) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v27
	v64 = F_CreateTupleDescCopy(m, v54)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v98 = v30
	v99 = v31
	v100 = v32
	v101 = v33
	v102 = v34
	v103 = v35
	v104 = v36
	v105 = v37
	v106 = v39
	v107 = v40
	goto L8
L8:
	;
	if v107 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v64
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[0])) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v27
	F_MarkPortalActive(m, l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[1]))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[2]))
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[3]))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[4]))
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[5]))
	goto L11
L11:
	;
	v92 = v21 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v21 + int32(12)
	goto L14
L12:
	;
	v98 = v48
	v99 = v25
	v100 = v50
	v101 = v84
	v102 = v82
	v103 = v86
	v104 = v88
	v105 = v90
	v106 = v27
	v107 = int32(0)
	goto L8
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[1])) = v102
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[2])) = v101
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[0])) = v100
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[4])) = v104
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[5])) = v105
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[3])) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_PopActiveSnapshot(m)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L5
	} else {
		goto L51
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[5])) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[2])) = v21 + int32(16)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v116 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[1])) = v102
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[2])) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L49
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[4])) = v116
	goto L21
L20:
	;
	goto L21
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[0])) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[3])) = v120
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_PushActiveSnapshot(m, v124)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v136&int32(2) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	v165 = F_CreateDestReceiver(m, int32(6))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_ExecutorRewind(m, v98)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v154 = v151 ^ int32(1)
	goto L23
L27:
	;
	v154 = int32(1)
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	v179 = int32(1)
	v180 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+36)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v165)+32)) = v180
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)) = uint8(v179)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v169
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_ExecutorRun(m, v98, v154, int64(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	m.T0[v200].(func(*base.Module, int32))(m, v199)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v212 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_ExecutorFinish(m, v98)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_ExecutorEnd(m, v98)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_FreeQueryDesc(m, v98)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[0])) = v250
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	if v252 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	goto L38
L36:
	;
	goto L37
L37:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_tuplestore_rescan(m, v287)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L5
	} else {
		goto L42
	}
L38:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	v285 = F_tuplestore_skiptuples(m, v273, int64(1000000), int32(1))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	if v285 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L15
L42:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v299&int32(2) == int32(0) {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	v304 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	v316 = F_tuplestore_skiptuples(m, v305, v304, int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if v316 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_errmsg_internal(m, int32(_a_F_PersistHoldablePortal_1), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_errfinish(m, int32(_a_F_PersistHoldablePortal_2), int32(470), int32(_a_F_PersistHoldablePortal_3))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[4])) = v104
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[5])) = v105
	*(*int32)(unsafe.Add(mBase, _c_F_PersistHoldablePortal[3])) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_pg_re_throw(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v21)+172)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v21)+180)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v21)+184)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v21)+188)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v21)+196)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21)+200)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v21)+204)) = v106
	F_MemoryContextDeleteChildren(m, v433)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	goto L4
L53:
	;
	v468 = int32(v464)
	m.G0 = v21
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	if v21+int32(12) == v474 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	m.ExcPending = 1
	goto L62
L55:
	;
	if v478 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	v478 = v476
	goto L58
L57:
	;
	v478 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v21)+204))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v21)+196))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v21)+192))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v21)+188))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v21)+184))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v21)+180))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v21)+172))
	v30 = v480
	v31 = v481
	v32 = v482
	v33 = v486
	v34 = v487
	v35 = v485
	v36 = v484
	v37 = v483
	v38 = v478
	v39 = v479
	v40 = v470
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v471, v470)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PrefetchBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+118)))
	if v13 == int32(116) {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
		if v16 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_PrefetchBuffer_0), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PrefetchBuffer_1), int32(662), int32(_a_F_PrefetchBuffer_2))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v19 == int32(0) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v23
				v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = v25
				v27 = F_smgropen(m, v10, v22)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v27
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
					if v31 != 0 {
						v39 = v31
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
						*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v33
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = v35
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+72))
						v39 = v37
					}
					*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v39 + int32(1)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v44 = v43
					v45 = m.G0
					v47 = v45 - int32(32)
					m.G0 = v47
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v55
					v60 = *(*int32)(unsafe.Add(mBase, _c_F_PrefetchBuffer[0]))
					if v60 != 0 {
						v65 = v60
						v68 = int32(0)
						v70 = F_hash_search(m, v65, v47+int32(12), v68, v68)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 != 0 {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
								m.G0 = v47 + int32(32)
								m.G0 = v10 + int32(32)
								return
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrefetchBuffer[1])))
								if v77&int32(1) != 0 {
									m.G0 = v47 + int32(32)
									m.G0 = v10 + int32(32)
									return
								} else {
									v81 = F_smgrprefetch(m, v44, l2, l3, int32(1))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										if v81 == int32(0) {
										} else {
											v85 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v85)
										}
										m.G0 = v47 + int32(32)
										m.G0 = v10 + int32(32)
										return
									}
								}
							}
						}
					} else {
						F_InitLocalBuffers(m)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_PrefetchBuffer[0]))
							v65 = v64
							v68 = int32(0)
							v70 = F_hash_search(m, v65, v47+int32(12), v68, v68)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								if v70 != 0 {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
									m.G0 = v47 + int32(32)
									m.G0 = v10 + int32(32)
									return
								} else {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrefetchBuffer[1])))
									if v77&int32(1) != 0 {
										m.G0 = v47 + int32(32)
										m.G0 = v10 + int32(32)
										return
									} else {
										v81 = F_smgrprefetch(m, v44, l2, l3, int32(1))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											if v81 == int32(0) {
											} else {
												v85 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v85)
											}
											m.G0 = v47 + int32(32)
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v44 = v19
				v45 = m.G0
				v47 = v45 - int32(32)
				m.G0 = v47
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v51
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v53
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v55
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_PrefetchBuffer[0]))
				if v60 != 0 {
					v65 = v60
					v68 = int32(0)
					v70 = F_hash_search(m, v65, v47+int32(12), v68, v68)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						if v70 != 0 {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
							m.G0 = v47 + int32(32)
							m.G0 = v10 + int32(32)
							return
						} else {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrefetchBuffer[1])))
							if v77&int32(1) != 0 {
								m.G0 = v47 + int32(32)
								m.G0 = v10 + int32(32)
								return
							} else {
								v81 = F_smgrprefetch(m, v44, l2, l3, int32(1))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									if v81 == int32(0) {
									} else {
										v85 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v85)
									}
									m.G0 = v47 + int32(32)
									m.G0 = v10 + int32(32)
									return
								}
							}
						}
					}
				} else {
					F_InitLocalBuffers(m)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_PrefetchBuffer[0]))
						v65 = v64
						v68 = int32(0)
						v70 = F_hash_search(m, v65, v47+int32(12), v68, v68)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 != 0 {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v72 ^ int32(-1)
								m.G0 = v47 + int32(32)
								m.G0 = v10 + int32(32)
								return
							} else {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrefetchBuffer[1])))
								if v77&int32(1) != 0 {
									m.G0 = v47 + int32(32)
									m.G0 = v10 + int32(32)
									return
								} else {
									v81 = F_smgrprefetch(m, v44, l2, l3, int32(1))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										if v81 == int32(0) {
										} else {
											v85 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v85)
										}
										m.G0 = v47 + int32(32)
										m.G0 = v10 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v90 != 0 {
			v116 = v90
			F_PrefetchSharedBuffer(m, l0, v116, l2, l3)
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return
			} else {
				m.G0 = v10 + int32(32)
				return
			}
		} else {
			v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v92
			v94 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v94
			v98 = F_smgropen(m, v10+int32(16), v91)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v98
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
				if v102 != 0 {
					v110 = v102
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v104
					v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v104))) = v106
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
					v110 = v108
				}
				*(*int32)(unsafe.Add(mBase, uint32(v98)+72)) = v110 + int32(1)
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v116 = v114
				F_PrefetchSharedBuffer(m, l0, v116, l2, l3)
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					m.G0 = v10 + int32(32)
					return
				}
			}
		}
	}
}
func F_PrepareRedoRemove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemove[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 <= v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L28
	}
L2:
	;
	m.G0 = v10 + int32(32)
	return
L3:
	;
	v22 = v3
	goto L4
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(8)+v22<<(uint(int32(2))%32))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if l0 != v30 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v37 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v33 = v22 + int32(1)
	if v14 != v33 {
		v22 = v33
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L2
L10:
	;
	return
L11:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_PrepareRedoRemove_0), v10+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+45)))
	if v50 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_errfinish(m, int32(_a_F_PrepareRedoRemove_1), int32(2601), int32(_a_F_PrepareRedoRemove_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_RemoveTwoPhaseFile(m, l0, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemove[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 <= int32(0) {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v61 = v56 + int32(8)
	v66 = int32(0)
	goto L22
L22:
	;
	v72 = v61 + v66<<(uint(int32(2))%32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 != v29 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v79 = v57 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v79
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61+v79<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v29
	goto L2
L24:
	;
	v76 = v66 + int32(1)
	if v57 != v76 {
		v66 = v76
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v29
	F_errmsg_internal(m, int32(_a_F_PrepareRedoRemove_3), v10)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_PrepareRedoRemove_1), int32(650), int32(_a_F_PrepareRedoRemove_4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcedureCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 float32, l28 float32) {
	mBase := m.M
	_ = mBase
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v198 int32
	_ = v198
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v374 int32
	_ = v374
	var v383 int64
	_ = v383
	var v393 int32
	_ = v393
	var v396 int64
	_ = v396
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
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
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v673 int32
	_ = v673
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v804 int32
	_ = v804
	var v822 int32
	_ = v822
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v897 int32
	_ = v897
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v968 int32
	_ = v968
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1050 int32
	_ = v1050
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1152 int32
	_ = v1152
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	v41 = m.G0
	v43 = v41 - int32(496)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l18)+16))
	if base.Ui32(v45) < base.Ui32(int32(101)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l19 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L20
	} else {
		goto L370
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L20
	} else {
		goto L367
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l19)+4))
	if v48 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v58 = l18
	v59 = v45
	goto L7
L7:
	;
	if l20 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l19)+16))
	if v51 <= int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l19)+8))
	if v54 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l19)+12))
	if v55 != int32(26) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v58 = l19
	v59 = v51
	goto L7
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L20
	} else {
		goto L364
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l20)+4))
	if v60 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v72 = int32(0)
	goto L15
L15:
	;
	v74 = l18 + int32(24)
	v75 = F_check_valid_polymorphic_signature(m, l5, v74, v45)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l20)+16))
	if v63 != v59 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l20)+8))
	if v65 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l20)+12))
	if v66 != int32(18) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v72 = l20 + int32(24)
	goto L15
L20:
	;
	return
L21:
	;
	if v75 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = F_check_valid_internal_signature(m, l5, v74, v45)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L20
	} else {
		goto L359
	}
L25:
	;
	if v79 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = v58 + int32(24)
	v85 = int32(0)
	if base.B2i32(l19 == v85)|base.B2i32(v59 == v85) == v85 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	goto L28
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L20
	} else {
		goto L354
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1507 = m.ExcPending
	if v1507 != 0 {
		goto L20
	} else {
		goto L349
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L20
	} else {
		goto L344
	}
L31:
	;
	v123 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v198 = int32(0)
	if base.B2i32(v72 == v198)|base.B2i32(v59 == v198) == v198 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	if v72 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v156 = v123 + int32(1)
	if v156 != v59 {
		v123 = v156
		goto L34
	} else {
		goto L43
	}
L37:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123+v72))))
	v138 = v136 - int32(105)
	if base.B2i32(v138 == int32(0))|base.B2i32(v138 == int32(13)) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v146 = v84 + v123<<(uint(int32(2))%32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = F_check_valid_polymorphic_signature(m, v147, v74, v45)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v148 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v151 = F_check_valid_internal_signature(m, v150, v74, v45)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	if v151 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	goto L35
L44:
	;
	v239 = int32(0)
	v240 = v198
	goto L47
L45:
	;
	v374 = v198
	goto L46
L46:
	;
	v383 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+486)) = v383
	*(*int64)(unsafe.Add(mBase, uint32(v43)+480)) = v383
	*(*int64)(unsafe.Add(mBase, uint32(v43)+472)) = v383
	*(*int64)(unsafe.Add(mBase, uint32(v43)+464)) = v383
	v393 = int32(0)
	base.MemoryFill(m, v43+int32(336), v393, int32(120))
	v396 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+326)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v43)+320)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v43)+312)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v43)+304)) = v396
	v405 = v43 + int32(240)
	v407 = F_strncpy(m, v405, l1, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v407)+63)) = uint8(v393)
	goto L80
L47:
	;
	v249 = v239 + v72
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	switch v250 - int32(98) {
	case 0, 7:
		goto L56
	default:
		goto L53
	case 13:
		goto L55
	case 18:
		v339 = v240
		goto L49
	case 20:
		goto L54
	}
L48:
	;
	v374 = v339
	goto L46
L49:
	;
	v341 = v239 + int32(1)
	if v341 != v59 {
		v239 = v341
		v240 = v339
		goto L47
	} else {
		goto L79
	}
L50:
	;
	v339 = int32(_a_F_ProcedureCreate_0)
	goto L49
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L20
	} else {
		goto L76
	}
L52:
	;
	v339 = int32(2283)
	goto L49
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L20
	} else {
		goto L73
	}
L54:
	;
	if v240 != 0 {
		goto L51
	} else {
		goto L65
	}
L55:
	;
	if base.B2i32(l12 != int32(112))|base.B2i32(v240 == int32(0)) != 0 {
		v339 = v240
		goto L49
	} else {
		goto L61
	}
L56:
	;
	v253 = int32(0)
	if v240 == v253 {
		v339 = v253
		goto L49
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_1), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(279), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L20
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_1), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(283), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L20
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v84+v239<<(uint(int32(2))%32))))
	switch v288 - int32(2276) {
	case 0:
		v339 = v288
		goto L49
	case 1:
		goto L52
	default:
		goto L66
	}
L66:
	;
	if v288 == int32(_a_F_ProcedureCreate_4) {
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v293 = F_get_element_type(m, v288)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	if v293 != 0 {
		v339 = v293
		goto L49
	} else {
		goto L69
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_5), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(305), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v312 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v312
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_6), v43+int32(16))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L20
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(310), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_1), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L20
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(290), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	goto L48
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+400)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v43)+396)) = l17
	*(*int32)(unsafe.Add(mBase, uint32(v43)+392)) = l16
	*(*int32)(unsafe.Add(mBase, uint32(v43)+388)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v43)+384)) = l15
	*(*int32)(unsafe.Add(mBase, uint32(v43)+380)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v43)+376)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v43)+372)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v43)+368)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v43)+364)) = v374
	*(*float32)(unsafe.Add(mBase, uint32(v43)+360)) = l28
	*(*float32)(unsafe.Add(mBase, uint32(v43)+356)) = l27
	*(*int32)(unsafe.Add(mBase, uint32(v43)+352)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v43)+348)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v43)+344)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v43)+340)) = v405
	if l22 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l22)+4)))
	v427 = v426
	goto L83
L82:
	;
	v427 = int32(0)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+412)) = l18
	*(*int32)(unsafe.Add(mBase, uint32(v43)+408)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v43)+404)) = v427
	if l19 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if l20 != 0 {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+416)) = l19
	goto L84
L86:
	;
	goto L87
L87:
	;
	v432 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+484)) = uint8(v432)
	goto L84
L88:
	;
	if l21 != 0 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+420)) = l20
	goto L88
L90:
	;
	goto L91
L91:
	;
	v435 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+485)) = uint8(v435)
	goto L88
L92:
	;
	if l22 != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+424)) = l21
	goto L92
L94:
	;
	goto L95
L95:
	;
	v438 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+486)) = uint8(v438)
	goto L92
L96:
	;
	if l23 != 0 {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v440 = F_nodeToString(m, l22)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L20
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+487)) = uint8(v445)
	goto L96
L100:
	;
	v442 = F_cstring_to_text(m, v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L20
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+428)) = v442
	goto L96
L102:
	;
	v450 = F_cstring_to_text(m, l9)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L20
	} else {
		goto L106
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+432)) = l23
	goto L102
L104:
	;
	goto L105
L105:
	;
	v448 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+488)) = uint8(v448)
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+436)) = v450
	if l10 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if l11 != 0 {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v453 = F_cstring_to_text(m, l10)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L20
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+490)) = uint8(v456)
	goto L107
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+440)) = v453
	goto L107
L112:
	;
	if l25 != 0 {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v458 = F_nodeToString(m, l11)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L20
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v463 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+491)) = uint8(v463)
	goto L112
L116:
	;
	v460 = F_cstring_to_text(m, v458)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+444)) = v460
	goto L112
L118:
	;
	v470 = F_table_open(m, int32(1255), int32(3))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L20
	} else {
		goto L122
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+448)) = l25
	goto L118
L120:
	;
	goto L121
L121:
	;
	v466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+492)) = uint8(v466)
	goto L118
L122:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v470)+52))
	v474 = F_SearchSysCache3(m, int32(46), l1, l18, l2)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L20
	} else {
		goto L130
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L20
	} else {
		goto L338
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L20
	} else {
		goto L332
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L20
	} else {
		goto L326
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L20
	} else {
		goto L319
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L20
	} else {
		goto L310
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L20
	} else {
		goto L306
	}
L129:
	;
	v990 = F_new_object_addresses(m)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L20
	} else {
		goto L243
	}
L130:
	;
	if v474 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if l3 == int32(0) {
		goto L128
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v927 = F_get_user_default_acl(m, int32(19), l6, l2)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L20
	} else {
		goto L236
	}
L134:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v474)+16))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479)+22)))
	v481 = v479 + v480
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v483 = F_object_ownercheck(m, int32(1255), v482, l6)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L20
	} else {
		goto L135
	}
L135:
	;
	if v483 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	F_aclcheck_error(m, int32(2), int32(19), l1)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L20
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+96)))
	if v491 != l12&int32(255) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L138
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L20
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if l12 == int32(97) {
		goto L153
	} else {
		goto L154
	}
L143:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L20
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_7), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L20
	} else {
		goto L145
	}
L145:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+96)))
	switch v507 - int32(97) {
	case 0:
		v513 = int32(_a_F_ProcedureCreate_8)
		goto L147
	default:
		goto L146
	case 5:
		goto L150
	case 15:
		goto L149
	case 22:
		goto L148
	}
L146:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(421), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L20
	} else {
		goto L152
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+128)) = l1
	F_errdetail(m, v513, v43+int32(128))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L20
	} else {
		goto L151
	}
L148:
	;
	v513 = int32(_a_F_ProcedureCreate_9)
	goto L147
L149:
	;
	v513 = int32(_a_F_ProcedureCreate_10)
	goto L147
L150:
	;
	v513 = int32(_a_F_ProcedureCreate_11)
	goto L147
L151:
	;
	goto L146
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	v530 = int32(_a_F_ProcedureCreate_12)
	goto L155
L154:
	;
	v530 = int32(_a_F_ProcedureCreate_13)
	goto L155
L155:
	;
	v532 = base.B2i32(l12 == int32(112))
	if l12 == int32(112) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v533 = int32(_a_F_ProcedureCreate_14)
	goto L158
L157:
	;
	v533 = v530
	goto L158
L158:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v481)+108))
	if l5 != v534 {
		goto L127
	} else {
		goto L159
	}
L159:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+100)))
	if v536 != l4 {
		goto L127
	} else {
		goto L160
	}
L160:
	;
	if l5 != int32(2249) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v619 = v43 + int32(227)
	v620 = F_SysCacheGetAttr(m, int32(46), v474, int32(23), v619)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L20
	} else {
		goto L182
	}
L162:
	;
	v540 = F_build_function_result_tupdesc_t(m, v474)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L20
	} else {
		goto L163
	}
L163:
	;
	v542 = F_build_function_result_tupdesc_d(m, l12, l19, l20, l21)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	if v540|v542 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v547 = int32(0)
	if base.B2i32(v540 == v547)|base.B2i32(v542 == v547) != 0 {
		goto L126
	} else {
		goto L166
	}
L166:
	;
	v552 = int32(0)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	if v556 != v557 {
		v608 = v552
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v608 == int32(0) {
		goto L126
	} else {
		goto L181
	}
L168:
	;
	goto L167
L169:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	if v559 != v560 {
		v608 = v552
		goto L168
	} else {
		goto L170
	}
L170:
	;
	if v556 <= int32(0) {
		v608 = int32(1)
		goto L168
	} else {
		goto L171
	}
L171:
	;
	v566 = v556 << (uint(int32(4)) % 32)
	v568 = int32(20)
	v574 = int32(0)
	goto L172
L172:
	;
	v581 = v574 * int32(100)
	v582 = v540 + v566 + v568 + v581
	v583 = int32(4)
	v585 = v581 + (v542 + v566 + v568)
	v588 = F_strcmp(m, v582+v583, v585+v583)
	mBase = m.M
	if v588 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v608 = int32(0)
	goto L168
L174:
	;
	goto L173
L175:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v582)+68))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v585)+68))
	if v589 != v590 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v585)+76))
	if v592 != v593 {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v582)+96))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v585)+96))
	if v595 != v596 {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+91)))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+91)))
	if v598 != v599 {
		goto L174
	} else {
		goto L179
	}
L179:
	;
	v601 = int32(1)
	v603 = v574 + v601
	if v556 != v603 {
		v574 = v603
		goto L172
	} else {
		goto L180
	}
L180:
	;
	v608 = v601
	goto L168
L181:
	;
	goto L161
L182:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+227)))
	if v622 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v762 = int32(*(*int16)(unsafe.Add(mBase, uint32(v481)+106)))
	if v762 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L184:
	;
	v623 = int32(0)
	v627 = F_SysCacheGetAttr(m, int32(46), v474, int32(22), v619)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+227)))
	if v629 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v630 = v623
	goto L188
L187:
	;
	v630 = v627
	goto L188
L188:
	;
	v633 = F_get_func_input_arg_names(m, v620, v630, v43+int32(228))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L20
	} else {
		goto L189
	}
L189:
	;
	v637 = F_get_func_input_arg_names(m, l21, l20, v43+int32(220))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L20
	} else {
		goto L190
	}
L190:
	;
	if v633 <= int32(0) {
		goto L183
	} else {
		goto L191
	}
L191:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v43)+220))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v43)+228))
	v673 = v623
	goto L192
L192:
	;
	v684 = v673 << (uint(int32(2)) % 32)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v642+v684)))
	if v686 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L183
L194:
	;
	if v637 <= v673 {
		goto L125
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v720 = v673 + int32(1)
	if v720 != v633 {
		v673 = v720
		goto L192
	} else {
		goto L207
	}
L197:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v641+v684)))
	if v689 == int32(0) {
		goto L125
	} else {
		goto L198
	}
L198:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686))))
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689))))
	if base.B2i32(v694 == int32(0))|base.B2i32(v694 != v697) != 0 {
		v715 = v694
		v716 = v697
		goto L200
	} else {
		goto L201
	}
L199:
	;
	if v715-v716 != 0 {
		goto L125
	} else {
		goto L206
	}
L200:
	;
	goto L199
L201:
	;
	v700 = v686
	v701 = v689
	goto L202
L202:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+1)))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700)+1)))
	if v705 == int32(0) {
		v715 = v705
		v716 = v704
		goto L200
	} else {
		goto L204
	}
L203:
	;
	v715 = v705
	v716 = v704
	goto L200
L204:
	;
	v708 = int32(1)
	if v705 == v704 {
		v700 = v700 + v708
		v701 = v701 + v708
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	goto L196
L207:
	;
	goto L193
L208:
	;
	v897 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+333)) = uint8(v897)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+307)) = uint8(v897)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+304)) = uint8(v897)
	v910 = F_heap_modify_tuple(m, v474, v472, v43+int32(336), v43+int32(464), v43+int32(304))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L20
	} else {
		goto L231
	}
L209:
	;
	if l22 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	v767 = v765
	goto L212
L211:
	;
	v767 = int32(0)
	goto L212
L212:
	;
	if v767 < v762 {
		goto L124
	} else {
		goto L213
	}
L213:
	;
	v771 = F_SysCacheGetAttrNotNull(m, int32(46), v474, int32(24))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L20
	} else {
		goto L214
	}
L214:
	;
	v773 = F_text_to_cstring(m, v771)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L20
	} else {
		goto L215
	}
L215:
	;
	v775 = F_stringToNode(m, v773)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L20
	} else {
		goto L216
	}
L216:
	;
	if l22 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	v779 = v777
	goto L219
L218:
	;
	v779 = int32(0)
	goto L219
L219:
	;
	if v775 == int32(0) {
		goto L208
	} else {
		goto L220
	}
L220:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	if v782 <= int32(0) {
		goto L208
	} else {
		goto L221
	}
L221:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l22)+12))
	v786 = int32(*(*int16)(unsafe.Add(mBase, uint32(v481)+106)))
	v804 = int32(0)
	v822 = v785 + (v779-v786)<<(uint(int32(2))%32)
	goto L222
L222:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v822)))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v775)+12))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v833+v804<<(uint(int32(2))%32))))
	v838 = F_exprType(m, v837)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L20
	} else {
		goto L224
	}
L223:
	;
	goto L208
L224:
	;
	v840 = F_exprType(m, v832)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L20
	} else {
		goto L225
	}
L225:
	;
	if v838 != v840 {
		goto L123
	} else {
		goto L226
	}
L226:
	;
	v844 = v822 + int32(4)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l22)+12))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	if base.Ui32(v844) < base.Ui32(v846+v847<<(uint(int32(2))%32)) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v852 = v844
	goto L229
L228:
	;
	v852 = int32(0)
	goto L229
L229:
	;
	v854 = v804 + int32(1)
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	if v854 < v855 {
		v804 = v854
		v822 = v852
		goto L222
	} else {
		goto L230
	}
L230:
	;
	goto L223
L231:
	;
	F_CatalogTupleUpdate(m, v470, v910+int32(4), v910)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L20
	} else {
		goto L232
	}
L232:
	;
	F_ReleaseCatCache(m, v474)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L20
	} else {
		goto L233
	}
L233:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v910)+16))
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+22)))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v919+v920)))
	v924 = F_deleteDependencyRecordsFor(m, int32(1255), v922, int32(1))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L20
	} else {
		goto L234
	}
L234:
	;
	v950 = v922
	v953 = v910
	v968 = v897
	goto L129
L235:
	;
	v934 = F_GetNewOidWithIndex(m, v470, int32(2690), int32(1))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L20
	} else {
		goto L240
	}
L236:
	;
	if v927 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+452)) = v927
	goto L235
L238:
	;
	goto L239
L239:
	;
	v930 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+493)) = uint8(v930)
	goto L235
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+336)) = v934
	v941 = F_heap_form_tuple(m, v472, v43+int32(336), v43+int32(464))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	F_CatalogTupleInsert(m, v470, v941)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v941)+16))
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+22)))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v945+v946)))
	v950 = v948
	v953 = v941
	v968 = v927
	goto L129
L243:
	;
	v992 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1255)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = v992
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(2615)
	v1003 = v43 + int32(228)
	F_add_exact_object_address(m, v1003, v990)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(2612)
	F_add_exact_object_address(m, v1003, v990)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(1247)
	F_add_exact_object_address(m, v1003, v990)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	if v59 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1050 = int32(0)
	goto L250
L248:
	;
	goto L249
L249:
	;
	if l24 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(1247)
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v84+v1050<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = v1065
	F_add_exact_object_address(m, v43+int32(228), v990)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L20
	} else {
		goto L252
	}
L251:
	;
	goto L249
L252:
	;
	v1074 = v1050 + int32(1)
	if v1074 != v59 {
		v1050 = v1074
		goto L250
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	if l26 != 0 {
		goto L261
	} else {
		goto L262
	}
L255:
	;
	v1118 = int32(0)
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l24)+4))
	if v1119 <= v1118 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1152 = v1118
	goto L257
L257:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l24)+12))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1162+v1152<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = v1166
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(3576)
	F_add_exact_object_address(m, v43+int32(228), v990)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L20
	} else {
		goto L259
	}
L258:
	;
	goto L254
L259:
	;
	v1177 = v1152 + int32(1)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(l24)+4))
	if v1177 < v1178 {
		v1152 = v1177
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+232)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v43)+228)) = int32(1255)
	F_add_exact_object_address(m, v43+int32(228), v990)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L20
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	F_record_object_address_dependencies(m, l0, v990, int32(110))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L20
	} else {
		goto L265
	}
L264:
	;
	goto L263
L265:
	;
	F_free_object_addresses(m, v990)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L20
	} else {
		goto L266
	}
L266:
	;
	v1234 = int32(0)
	if base.B2i32(l11 == v1234)|base.B2i32(l7 != int32(14)) == v1234 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_recordDependencyOnExpr(m, l0, l11, int32(0))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L20
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if l22 != 0 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	goto L269
L271:
	;
	F_recordDependencyOnExpr(m, l0, l22, int32(0))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L20
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	if v474 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L274:
	;
	goto L273
L275:
	;
	F_recordDependencyOnOwner(m, int32(1255), v950, l6)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L20
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	F_recordDependencyOnCurrentExtension(m, l0, base.B2i32(v474 != int32(0)))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L20
	} else {
		goto L280
	}
L278:
	;
	F_recordDependencyOnNewAcl(m, int32(1255), v950, l6, v968)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L20
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	F_pfree(m, v953)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L20
	} else {
		goto L281
	}
L281:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[0]))
	if v1262 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1264 = int32(0)
	F_RunObjectPostCreateHook(m, int32(1255), v950, v1264, v1264)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L20
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_relation_close(m, v470, int32(3))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L20
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	if l8 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if v474 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L288:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L20
	} else {
		goto L289
	}
L289:
	;
	if l25 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1306 = F_OidFunctionCall1Coll(m, l8, int32(0), v950)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L20
	} else {
		goto L301
	}
L291:
	;
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcedureCreate[1])))
	if v1278&int32(1) == int32(0) {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v1284 = int32(_a_F_ProcedureCreate_15)
	v1286 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[2]))
	v1288 = v1286 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[2])) = v1288
	goto L293
L293:
	;
	v1292 = F_superuser(m)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L20
	} else {
		goto L294
	}
L294:
	;
	if v1292 != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1294 = int32(5)
	goto L297
L296:
	;
	v1294 = int32(6)
	goto L297
L297:
	;
	F_ProcessGUCArray(m, l25, v1294, int32(13), int32(2))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L20
	} else {
		goto L298
	}
L298:
	;
	v1300 = F_OidFunctionCall1Coll(m, l8, int32(0), v950)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L20
	} else {
		goto L299
	}
L299:
	;
	F_AtEOXact_GUC(m, int32(1), v1288)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L20
	} else {
		goto L300
	}
L300:
	;
	goto L287
L301:
	;
	goto L287
L302:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[3]))
	F_pgstat_create_transactional(m, int32(3), v1313, base.I64_extend_i32_u(v950))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L20
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	m.G0 = v43 + int32(496)
	return
L305:
	;
	goto L304
L306:
	;
	F_errcode(m, int32(50884740))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L20
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+144)) = l1
	F_errmsg(m, int32(_a_F_ProcedureCreate_16), v43+int32(144))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L20
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(403), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L20
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L20
	} else {
		goto L311
	}
L311:
	;
	if l12 == int32(112) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1347 = int32(_a_F_ProcedureCreate_17)
	goto L314
L313:
	;
	v1347 = int32(_a_F_ProcedureCreate_18)
	goto L314
L314:
	;
	F_errmsg(m, v1347, int32(0))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L20
	} else {
		goto L315
	}
L315:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v1352 = F_format_procedure(m, v1351)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L20
	} else {
		goto L316
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+116)) = v1352
	*(*int32)(unsafe.Add(mBase, uint32(v43)+112)) = v533
	F_errhint(m, int32(_a_F_ProcedureCreate_19), v43+int32(112))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L20
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(449), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L20
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L20
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_18), int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L20
	} else {
		goto L321
	}
L321:
	;
	F_errdetail(m, int32(_a_F_ProcedureCreate_20), int32(0))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L20
	} else {
		goto L322
	}
L322:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v1382 = F_format_procedure(m, v1381)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L20
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+100)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = v533
	F_errhint(m, int32(_a_F_ProcedureCreate_19), v43+int32(96))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L20
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(476), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L20
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L20
	} else {
		goto L327
	}
L327:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v43)+228))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1404+v673<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v1408
	F_errmsg(m, int32(_a_F_ProcedureCreate_21), v43+int32(80))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L20
	} else {
		goto L328
	}
L328:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v1416 = F_format_procedure(m, v1415)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L20
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v1416
	*(*int32)(unsafe.Add(mBase, uint32(v43)+64)) = v533
	F_errhint(m, int32(_a_F_ProcedureCreate_19), v43-int32(-64))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L20
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(521), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L20
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L20
	} else {
		goto L333
	}
L333:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_22), int32(0))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L20
	} else {
		goto L334
	}
L334:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v1442 = F_format_procedure(m, v1441)
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L20
	} else {
		goto L335
	}
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v1442
	*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v533
	F_errhint(m, int32(_a_F_ProcedureCreate_19), v43+int32(32))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L20
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(547), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L20
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L20
	} else {
		goto L339
	}
L339:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_23), int32(0))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L20
	} else {
		goto L340
	}
L340:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v1468 = F_format_procedure(m, v1467)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L20
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+52)) = v1468
	*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v533
	F_errhint(m, int32(_a_F_ProcedureCreate_19), v43+int32(48))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L20
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(571), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1481 = m.ExcPending
	if v1481 != 0 {
		goto L20
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L20
	} else {
		goto L345
	}
L345:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_24), int32(0))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L20
	} else {
		goto L346
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+160)) = v151
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_25), v43+int32(160))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L20
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(260), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L20
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L20
	} else {
		goto L350
	}
L350:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_26), int32(0))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L20
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+176)) = v148
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_25), v43+int32(176))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L20
	} else {
		goto L352
	}
L352:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(252), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L20
	} else {
		goto L353
	}
L353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L20
	} else {
		goto L355
	}
L355:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_24), int32(0))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L20
	} else {
		goto L356
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+192)) = v79
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_25), v43+int32(192))
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L20
	} else {
		goto L357
	}
L357:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(231), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L20
	} else {
		goto L358
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L20
	} else {
		goto L360
	}
L360:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_26), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L20
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+208)) = v75
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_25), v43+int32(208))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L20
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(218), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L20
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_27), int32(0))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L20
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(203), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L20
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_28), int32(0))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L20
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(179), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L20
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L20
	} else {
		goto L371
	}
L371:
	;
	v1605 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v1605
	F_errmsg_plural(m, int32(_a_F_ProcedureCreate_29), int32(_a_F_ProcedureCreate_30), v1605, v43)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L20
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(161), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L20
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessUtility(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessUtility[0]))
	if v11 != 0 {
		m.T0[v11].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4, l5, l6, l7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ProcessUtility(m, l0, l1, l2, l3, l4, l5, l6, l7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
func F_p_isEOF(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 != v5 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		v9 = v7
	} else {
		v9 = int32(0)
	}
	return base.B2i32(v9 == int32(0))
}
func F_p_isalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8<<(uint(int32(2))%32))))
			return base.B2i32(base.Ui32(int32(127)) < base.Ui32(v12)) | base.B2i32(base.Ui32(v12|int32(32)-int32(97)) < base.Ui32(int32(26)))
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+v25<<(uint(int32(2))%32))))
			if base.Ui32(v29) <= base.Ui32(int32(_a_F_p_isalpha_0)) {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v29)>>(uint(int32(8))%32)))+uint32(_c_F_p_isalpha[0]))))
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v29)>>(uint(int32(3))%32))&int32(31)|v38<<(uint(int32(5))%32))+uint32(_c_F_p_isalpha[0]))))
				v50 = int32(base.Ui32(v42)>>(uint(v29&int32(7))%32)) & int32(1)
			} else {
				v50 = base.B2i32(base.Ui32(v29) < base.Ui32(int32(_a_F_p_isalpha_1)))
			}
			return v50
		}
	} else {
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
		v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v54))))
		return base.B2i32(base.Ui32((v56|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)))
	}
}
func F_p_ishost(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	v2 = int32(0)
	v6 = F_palloc0(m, int32(48))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12 + v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v17 - v19
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v24 + v26<<(uint(int32(2))%32)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v31 + v33<<(uint(int32(2))%32)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v39 = F_palloc(m, int32(32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v41 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v39)+8)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(0)
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+29)) = uint8(v52)
	F_check_stack_depth(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v56 = F_TParserGet(m, v6)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	if v90 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v56 == int32(0) {
		v89 = v2
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	if v60 != int32(6) {
		v89 = v2
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 + v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v69 + v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v74 + v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v79 + v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v85
	v89 = int32(1)
	goto L11
L15:
	;
	v91 = v90
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_pfree(m, v6)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	F_pfree(m, v91)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v95
	if v95 != 0 {
		v91 = v95
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	return v89
}
func F_p_isspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v84
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+v70))))
	v84 = base.B2i32(v72 == int32(32)) | base.B2i32(base.Ui32((v72-int32(9))&int32(255)) < base.Ui32(int32(5)))
	goto L1
L5:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
	if base.Ui32(int32(127)) < base.Ui32(v13) {
		v84 = int32(0)
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+v26<<(uint(int32(2))%32))))
	if v30 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return base.B2i32(v13 == int32(32)) | base.B2i32(base.Ui32(v13-int32(9)) < base.Ui32(int32(5)))
L9:
	;
	return v66
L10:
	;
	v66 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v66 = base.B2i32(v59 != int32(0))
	goto L9
L14:
	;
	v39 = int32(_a_F_p_isspace_0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v49 = int32(_a_F_p_isspace_0)
	v50 = F_wcslen(m, v49)
	mBase = m.M
	v59 = v50<<(uint(int32(2))%32) + v49
	goto L13
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v42 != 0 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if v30 != v42 {
		v39 = v39 + int32(4)
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L21
L23:
	;
	v48 = v39
	goto L25
L24:
	;
	v48 = int32(0)
	goto L25
L25:
	;
	v59 = v48
	goto L13
}
func F_p_isstophost(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2 == int32(1) {
		v5 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)) = uint8(v5)
		v9 = int32(1)
	} else {
		v9 = int32(0)
	}
	return v9
}
func F_palloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_palloc[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)) = uint8(v2)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v10 = m.T0[v9].(func(*base.Module, int32, int32, int32) int32)(m, v4, l0, v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_palloc0(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_palloc0[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v2)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = m.T0[v10].(func(*base.Module, int32, int32, int32) int32)(m, v5, l0, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l0&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l0)) == int32(0) {
			if l0 == int32(0) {
				return v11
			} else {
				v26 = l0 + v11
				v28 = v11 + int32(4)
				if base.Ui32(v28) < base.Ui32(v26) {
					v30 = v26
				} else {
					v30 = v28
				}
				v35 = (v11^int32(-1)+v30)&int32(-4) + int32(4)
				if v35 == int32(0) {
					return v11
				} else {
					base.MemoryFill(m, v11, int32(0), v35)
					return v11
				}
			}
		} else {
			if l0 == int32(0) {
			} else {
				base.MemoryFill(m, v11, int32(0), l0)
			}
			return v11
		}
	}
}
func F_paramlist_param_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= v3 {
		v50 = v3
		m.G0 = v9 + int32(16)
		return v50
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		if v15 < v11 {
			v50 = v3
			m.G0 = v9 + int32(16)
			return v50
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != 0 {
				v21 = m.T0[v17].(func(*base.Module, int32, int32, int32, int32) int32)(m, v14, v11, int32(0), v9+int32(4))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v30 = v21
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
					if v31 == int32(0) {
						v50 = v3
						m.G0 = v9 + int32(16)
						return v50
					} else {
						v35 = F_palloc0(m, int32(28))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v11
							*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(8)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v40
							v44 = F_get_typcollation(m, v40)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v44
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v47
								v50 = v35
								m.G0 = v9 + int32(16)
								return v50
							}
						}
					}
				}
			} else {
				v30 = v14 + v11*int32(12) + int32(20)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
				if v31 == int32(0) {
					v50 = v3
					m.G0 = v9 + int32(16)
					return v50
				} else {
					v35 = F_palloc0(m, int32(28))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v11
						*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(8)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v40
						v44 = F_get_typcollation(m, v40)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v44
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v47
							v50 = v35
							m.G0 = v9 + int32(16)
							return v50
						}
					}
				}
			}
		}
	}
}
func F_parserOpenTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = v8 + int32(-20)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13
	v20 = int32(_a_F_parserOpenTable_0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_parserOpenTable[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v21
	*(*int32)(unsafe.Add(mBase, _c_F_parserOpenTable[0])) = v8 + int32(-12)
	goto L1
L1:
	;
	v28 = F_table_openrv_extended(m, l1, l2, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v168
	F_errdetail(m, int32(_a_F_parserOpenTable_1), v10)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L46
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L42
	}
L4:
	;
	return int32(0)
L5:
	;
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(-20))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_parserOpenTable[0])) = v143
	goto L41
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L36
	}
L11:
	;
	v36 = l0
	goto L14
L12:
	;
	goto L13
L13:
	;
	v120 = int32(0)
	goto L10
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	if v43 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v104 != 0 {
		v36 = v104
		goto L14
	} else {
		goto L35
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(0)
	if v49 < v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v52 = v46
	goto L21
L20:
	;
	v52 = v49
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v57 = int32(0)
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53+v57<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if base.B2i32(v69 == int32(0))|base.B2i32(v69 != v72) != 0 {
		v90 = v69
		v91 = v72
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v120 = int32(1)
	goto L10
L24:
	;
	if v90-v91 != 0 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	goto L24
L26:
	;
	v75 = v66
	v76 = v35
	goto L27
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v80 == int32(0) {
		v90 = v80
		v91 = v79
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v90 = v80
	v91 = v79
	goto L25
L29:
	;
	v83 = int32(1)
	if v80 == v79 {
		v75 = v75 + v83
		v76 = v76 + v83
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v94 = v57 + int32(1)
	if v52 != v94 {
		v57 = v94
		goto L22
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L23
L34:
	;
	goto L16
L35:
	;
	goto L15
L36:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v128
	F_errmsg(m, int32(_a_F_parserOpenTable_2), v8+int32(-48))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v120 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_parserOpenTable_3), int32(1469), int32(_a_F_parserOpenTable_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	m.G0 = v10 - int32(-64)
	return v28
L42:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v156 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v156
	F_errmsg(m, int32(_a_F_parserOpenTable_5), v8+int32(-32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_parserOpenTable_3), int32(1448), int32(_a_F_parserOpenTable_4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errhint(m, int32(_a_F_parserOpenTable_6), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_parserOpenTable_3), int32(1464), int32(_a_F_parserOpenTable_4))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_per_MultiFuncCall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	return v3
}
func F_pgl_longjmp(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_longjmp[0]))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F___wasm_longjmp(m, l0, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v6 = int32(_a_F_pgl_longjmp_0)
	v7 = int32(156)
	goto L6
L3:
	;
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v69 = int32(0)
	goto L3
L5:
	;
	v43 = v38
	v44 = v39
	v45 = v40
	goto L15
L6:
	;
	if (l0|v6)&int32(3) != 0 {
		v38 = l0
		v39 = v6
		v40 = v7
		goto L5
	} else {
		goto L9
	}
L8:
	;
	if v28 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v15 = l0
	v16 = v6
	v17 = v7
	goto L10
L10:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v20 != v21 {
		v38 = v15
		v39 = v16
		v40 = v17
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L8
L12:
	;
	v23 = int32(4)
	v24 = v16 + v23
	v26 = v15 + v23
	v28 = v17 - v23
	if base.Ui32(int32(3)) < base.Ui32(v28) {
		v15 = v26
		v16 = v24
		v17 = v28
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v38 = v26
	v39 = v24
	v40 = v28
	goto L5
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v48 == v49 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v69 = v48 - v49
	goto L3
L17:
	;
	v51 = int32(1)
	v56 = v45 - v51
	if v56 != 0 {
		v43 = v43 + v51
		v44 = v44 + v51
		v45 = v56
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgl_longjmp[1])))
	if v71 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgl_longjmp[2])) = uint8(v75)
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgl_longjmp[3])) = int32(100)
	m.Env.Emscripten_exit_with_live_runtime(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	return
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstatginindex_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstatginindex_internal(m, v2, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pgstatindexbyid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstatindexbyid_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstatindexbyid_1), int32(193), int32(_a_F_pgstatindexbyid_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v27 = F_relation_open(m, v3, int32(1))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = F_pgstatindex_impl(m, v27, l0)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					return v29
				}
			}
		}
	}
}
func F_pgstattuple(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_superuser(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v8 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pgstattuple_0), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pgstattuple_1), int32(178), int32(_a_F_pgstattuple_2))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v28 = F_textToQualifiedNameList(m, v4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = F_makeRangeVarFromNameList(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v33 = F_relation_openrv(m, v30, int32(1))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = F_pgstat_relation(m, v33, l0)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstattuplebyid_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pgstat_relation(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pkt_stream_flush(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v9 != 0 {
		v24 = int32(0)
		m.G0 = v7 + int32(16)
		return v24
	} else {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v10)
		v15 = F_pushf_write(m, l0, v7+int32(8), int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 < int32(0) {
				v24 = v15
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
				v24 = int32(0)
			}
			m.G0 = v7 + int32(16)
			return v24
		}
	}
}
func F_plain_crypt_verify(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
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
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
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
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v5
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v16 != int32(109) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(144)
	return v576
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v568
	v576 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v561 = F_psprintf(m, int32(_a_F_plain_crypt_verify_0), v10+int32(32))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L30
	} else {
		goto L147
	}
L4:
	;
	v512 = F_strlen(m, l0)
	mBase = m.M
	v517 = F_pg_md5_encrypt(m, l2, l0, v512, v10+int32(48), v10+int32(44))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L30
	} else {
		goto L134
	}
L5:
	;
	v128 = F_parse_scram_secret(m, l1, v10+int32(136), v10+int32(128), v10+int32(132), v10+int32(140), v10+int32(48), v10+int32(96))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L30
	} else {
		goto L31
	}
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v19 != int32(100) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v22 != int32(53) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v25 = F_strlen(m, l1)
	mBase = m.M
	if v25 != int32(35) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v29 = l1 + int32(3)
	v30 = int32(_a_F_plain_crypt_verify_1)
	v34 = m.G0
	v36 = v34 - int32(32)
	v37 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plain_crypt_verify[0])))
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v113 == int32(32) {
		goto L4
	} else {
		goto L29
	}
L11:
	;
	v113 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plain_crypt_verify[1])))
	if v49 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v29
	goto L17
L15:
	;
	goto L16
L16:
	;
	v63 = v30
	v64 = v45
	goto L20
L17:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v59 == v45 {
		v53 = v53 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v113 = v53 - v29
	goto L10
L19:
	;
	goto L18
L20:
	;
	v71 = v36 + int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 | v73<<(uint(v64)%32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v77 != 0 {
		v63 = v63 + v73
		v64 = v77
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v80 == int32(0) {
		v103 = v29
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v113 = v103 - v29
	goto L10
L24:
	;
	v84 = v29
	v85 = v80
	goto L25
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(base.Ui32(v85)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v93)>>(uint(v85)%32))&int32(1) == int32(0) {
		v103 = v84
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v103 = v101
	goto L23
L27:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v101 = v84 + int32(1)
	if v99 != 0 {
		v84 = v101
		v85 = v99
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	goto L5
L30:
	;
	return int32(0)
L31:
	;
	if v128 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v134 = int32(0)
	v135 = m.G0
	v137 = v135 - int32(192)
	m.G0 = v137
	*(*int32)(unsafe.Add(mBase, uint32(v137)+180)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v137)+40)) = v134
	v155 = F_parse_scram_secret(m, l1, v137+int32(184), v137+int32(176), v137+int32(180), v137+int32(188), v137+int32(112), v137+int32(80))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L30
	} else {
		goto L36
	}
L33:
	;
	if v485 != 0 {
		v576 = v134
		goto L1
	} else {
		goto L132
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L30
	} else {
		goto L129
	}
L35:
	;
	m.G0 = v137 + int32(192)
	goto L33
L36:
	;
	if v155 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v161 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L30
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v137)+188))
	v177 = F_strlen(m, v176)
	mBase = m.M
	v181 = v177 * int32(3) >> (uint(int32(2)) % 32)
	goto L44
L40:
	;
	if v161 == int32(0) {
		v485 = v5
		goto L35
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+32)) = l0
	F_errmsg(m, int32(_a_F_plain_crypt_verify_2), v137+int32(32))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L30
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_plain_crypt_verify_3), int32(547), int32(_a_F_plain_crypt_verify_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v485 = v5
	goto L35
L44:
	;
	v182 = F_palloc(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L30
	} else {
		goto L45
	}
L45:
	;
	v184 = F_strlen(m, v176)
	mBase = m.M
	v185 = int32(0)
	if v185 < v184 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v370 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L47:
	;
	if v181 != 0 {
		goto L89
	} else {
		goto L90
	}
L48:
	;
	v194 = v176 + v184
	v195 = v176
	v199 = v185
	v200 = v182
	v202 = v185
	v203 = v185
	goto L51
L49:
	;
	v341 = v182
	goto L50
L50:
	;
	v370 = v341 - v182
	goto L46
L51:
	;
	v207 = v195 + int32(1)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v208 != int32(61) {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	if v329 != 0 {
		goto L47
	} else {
		goto L88
	}
L53:
	;
	if base.Ui32(v327) < base.Ui32(v194) {
		v195 = v327
		v199 = v329
		v200 = v330
		v202 = v332
		v203 = v333
		goto L51
	} else {
		goto L87
	}
L54:
	;
	if v181 < v200-v182+int32(1) {
		goto L47
	} else {
		goto L74
	}
L55:
	;
	v285 = v207
	v286 = int32(2)
	v289 = v203 << (uint(int32(6)) % 32)
	goto L54
L56:
	;
	v274 = v271 + v270<<(uint(int32(6))%32)
	v276 = v267 + int32(1)
	if v276 == int32(4) {
		v285 = v268
		v286 = v269
		v289 = v274
		goto L54
	} else {
		goto L73
	}
L57:
	;
	v267 = int32(3)
	v268 = v195 + int32(2)
	v269 = int32(1)
	v270 = v227
	v271 = v222
	goto L56
L58:
	;
	if base.Ui32(int32(125)) < base.Ui32((v246-int32(1))&int32(255)) {
		goto L47
	} else {
		goto L71
	}
L59:
	;
	v212 = v208 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v212))|base.B2i32(int32(1)<<(uint(v212)%32)&int32(_a_F_plain_crypt_verify_5) == int32(0)) != 0 {
		v246 = v208
		v247 = v199
		v248 = v207
		v249 = v202
		v250 = v203
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v222 = int32(0)
	if v202 != 0 {
		v267 = v199
		v268 = v207
		v269 = v202
		v270 = v203
		v271 = v222
		goto L56
	} else {
		goto L63
	}
L62:
	;
	goto L47
L63:
	;
	switch v199 - int32(2) {
	case 0:
		goto L64
	case 1:
		goto L55
	default:
		goto L47
	}
L64:
	;
	if base.Ui32(v194) <= base.Ui32(v207) {
		goto L47
	} else {
		goto L65
	}
L65:
	;
	v227 = v203 << (uint(int32(6)) % 32)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v228 == int32(61) {
		goto L57
	} else {
		goto L66
	}
L66:
	;
	v232 = v228 - int32(9)
	if int32(1)<<(uint(v232)%32)&int32(_a_F_plain_crypt_verify_5) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v240 = base.B2i32(base.Ui32(v232) <= base.Ui32(int32(23)))
	goto L69
L68:
	;
	v240 = int32(0)
	goto L69
L69:
	;
	if v240 != 0 {
		goto L47
	} else {
		goto L70
	}
L70:
	;
	v246 = v228
	v247 = int32(3)
	v248 = v195 + int32(2)
	v249 = int32(1)
	v250 = v227
	goto L58
L71:
	;
	v258 = int32(*(*int8)(unsafe.Add(mBase, uint32(v246)+uint32(_c_F_plain_crypt_verify[2]))))
	if v258 < int32(0) {
		goto L47
	} else {
		goto L72
	}
L72:
	;
	v267 = v247
	v268 = v248
	v269 = v249
	v270 = v250
	v271 = v258
	goto L56
L73:
	;
	v327 = v268
	v329 = v276
	v330 = v200
	v332 = v269
	v333 = v274
	goto L53
L74:
	;
	v295 = int32(base.Ui32(v289) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v295)
	v298 = v200 + int32(1)
	if base.Ui32(v286) < base.Ui32(int32(2)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v302 = v286
	goto L77
L76:
	;
	v302 = int32(0)
	goto L77
L77:
	;
	if v302 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v181 < v298-v182+int32(1) {
		goto L47
	} else {
		goto L81
	}
L79:
	;
	v314 = v298
	goto L80
L80:
	;
	if v286 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v310 = int32(base.Ui32(v289) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)) = uint8(v310)
	v314 = v200 + int32(2)
	goto L80
L82:
	;
	v327 = v285
	v329 = int32(0)
	v330 = v324
	v332 = v325
	v333 = int32(0)
	goto L53
L83:
	;
	v324 = v314
	v325 = v286
	goto L82
L84:
	;
	goto L85
L85:
	;
	if v181 < v314-v182+int32(1) {
		goto L47
	} else {
		goto L86
	}
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v289)
	v324 = v314 + int32(1)
	v325 = int32(0)
	goto L82
L87:
	;
	goto L52
L88:
	;
	v341 = v330
	goto L50
L89:
	;
	base.MemoryFill(m, v182, int32(0), v181)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v370 = int32(-1)
	goto L46
L92:
	;
	v373 = int32(0)
	v376 = F_errstart(m, int32(15), v373)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L30
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v391 = F_pg_saslprep(m, l2, v137+int32(44))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L30
	} else {
		goto L99
	}
L95:
	;
	if v376 == int32(0) {
		v485 = v373
		goto L35
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = l0
	F_errmsg(m, int32(_a_F_plain_crypt_verify_2), v137)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L30
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_plain_crypt_verify_3), int32(558), int32(_a_F_plain_crypt_verify_4))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L30
	} else {
		goto L98
	}
L98:
	;
	v485 = v373
	goto L35
L99:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	if v391 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v394 = l2
	goto L102
L101:
	;
	v394 = v393
	goto L102
L102:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v137)+176))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v137)+180))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v137)+184))
	v399 = v137 + int32(144)
	v401 = v137 + int32(40)
	v402 = F_scram_SaltedPassword(m, v394, v395, v396, v182, v370, v397, v399, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L30
	} else {
		goto L103
	}
L103:
	;
	if v402 < int32(0) {
		goto L34
	} else {
		goto L104
	}
L104:
	;
	v408 = F_scram_ServerKey(m, v399, v395, v396, v137+int32(48), v401)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L30
	} else {
		goto L105
	}
L105:
	;
	if v408 < int32(0) {
		goto L34
	} else {
		goto L106
	}
L106:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	if v412 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_pfree(m, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L30
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v416 = v137 + int32(48)
	v418 = v137 + int32(80)
	if base.Ui32(int32(4)) <= base.Ui32(v396) {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	goto L109
L111:
	;
	v485 = base.B2i32(v480 == int32(0))
	goto L35
L112:
	;
	v480 = int32(0)
	goto L111
L113:
	;
	v454 = v449
	v455 = v450
	v456 = v451
	goto L123
L114:
	;
	if (v416|v418)&int32(3) != 0 {
		v449 = v416
		v450 = v418
		v451 = v396
		goto L113
	} else {
		goto L117
	}
L115:
	;
	v442 = v416
	v443 = v418
	v444 = v396
	goto L116
L116:
	;
	if v444 == int32(0) {
		goto L112
	} else {
		goto L122
	}
L117:
	;
	v426 = v416
	v427 = v418
	v428 = v396
	goto L118
L118:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	if v431 != v432 {
		v449 = v426
		v450 = v427
		v451 = v428
		goto L113
	} else {
		goto L120
	}
L119:
	;
	v442 = v437
	v443 = v435
	v444 = v439
	goto L116
L120:
	;
	v434 = int32(4)
	v435 = v427 + v434
	v437 = v426 + v434
	v439 = v428 - v434
	if base.Ui32(int32(3)) < base.Ui32(v439) {
		v426 = v437
		v427 = v435
		v428 = v439
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v449 = v442
	v450 = v443
	v451 = v444
	goto L113
L123:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	if v459 == v460 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v480 = v459 - v460
	goto L111
L125:
	;
	v462 = int32(1)
	v467 = v456 - v462
	if v467 != 0 {
		v454 = v454 + v462
		v455 = v455 + v462
		v456 = v467
		goto L123
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	goto L112
L129:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = v494
	F_errmsg_internal(m, int32(_a_F_plain_crypt_verify_6), v137+int32(16))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L30
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_plain_crypt_verify_3), int32(574), int32(_a_F_plain_crypt_verify_4))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L30
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	v510 = F_psprintf(m, int32(_a_F_plain_crypt_verify_7), v10+int32(16))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L30
	} else {
		goto L133
	}
L133:
	;
	v568 = v510
	goto L2
L134:
	;
	if v517 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v568 = v521
	goto L2
L136:
	;
	goto L137
L137:
	;
	v522 = int32(0)
	v524 = v10 + int32(48)
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v527 == v522)|base.B2i32(v527 != v530) != 0 {
		v548 = v527
		v549 = v530
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v548-v549 == int32(0) {
		v576 = v522
		goto L1
	} else {
		goto L145
	}
L139:
	;
	goto L138
L140:
	;
	v533 = v524
	v534 = l1
	goto L141
L141:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+1)))
	if v538 == int32(0) {
		v548 = v538
		v549 = v537
		goto L139
	} else {
		goto L143
	}
L142:
	;
	v548 = v538
	v549 = v537
	goto L139
L143:
	;
	v541 = int32(1)
	if v538 == v537 {
		v533 = v533 + v541
		v534 = v534 + v541
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v555 = F_psprintf(m, int32(_a_F_plain_crypt_verify_7), v10)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L30
	} else {
		goto L146
	}
L146:
	;
	v568 = v555
	goto L2
L147:
	;
	v568 = v561
	goto L2
}
func F_plainto_tsquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v9 = F_getTSCurrentConfig(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_DirectFunctionCall2Coll(m, int32(1160), int32(0), v9, v3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_plan_recursive_revoke(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v12 = l2 << (uint(int32(2)) % 32)
	v13 = l1 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if base.B2i32(v14 == int32(4))|l3&base.B2i32(v14 == int32(1)) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L24
	} else {
		goto L27
	}
L2:
	;
	return
L3:
	;
	v22 = l0 + int32(48)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12+v22)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v38 <= int32(0) {
		goto L2
	} else {
		goto L10
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+16)))
	if v32 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+16)))
	if v33 != int32(1) {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L2
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
	goto L4
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v45 = int32(0)
	goto L11
L11:
	;
	v54 = v45 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v22+v54)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+22)))
	v59 = v57 + v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v60 != v41 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v75 = int32(0)
	v76 = v38
	goto L18
L13:
	;
	v70 = v45 + int32(1)
	if v70 != v38 {
		v45 = v70
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	if v62 != int32(1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+v54)))
	if v66 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L12
L18:
	;
	v84 = v75 << (uint(int32(2)) % 32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v22+v84)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+56))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87+v88)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v90 != v91 {
		v103 = v76
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L2
L20:
	;
	v105 = v75 + int32(1)
	if v105 < v103 {
		v75 = v105
		v76 = v103
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1+v84)))
	if v94 == int32(4) {
		v103 = v76
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if l4 == int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_plan_recursive_revoke(m, l0, l1, v75, int32(0), l4)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v103 = v102
	goto L20
L26:
	;
	goto L19
L27:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_plan_recursive_revoke_0), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	F_errhint(m, int32(_a_F_plan_recursive_revoke_1), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_plan_recursive_revoke_2), int32(2494), int32(_a_F_plan_recursive_revoke_3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pop_arg_long_double(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v142 int64
	_ = v142
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = (v3 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7 + int32(16)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v25 = v12 & int64(281474976710655)
	v29 = int64(base.Ui64(v12)>>(uint(int64(48))%64)) & int64(32767)
	v30 = base.I32_wrap_i64(v29)
	if base.Ui32(v30-int32(_a_F_pop_arg_long_double_0)) <= base.Ui32(int32(2045)) {
		v39 = v25<<(uint(int64(4))%64) | int64(base.Ui64(v11)>>(uint(int64(60))%64))
		v44 = v11 & int64(1152921504606846975)
		if base.Ui64(int64(576460752303423489)) <= base.Ui64(v44) {
			v54 = v39 + int64(1)
		} else {
			if v44 != int64(576460752303423488) {
				v54 = v39
			} else {
				v54 = v39&int64(1) + v39
			}
		}
		v57 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v54))
		if base.Ui64(int64(4503599627370495)) < base.Ui64(v54) {
			v58 = int64(0)
		} else {
			v58 = v54
		}
		v135 = v58
		v142 = base.I64_extend_i32_u(v57) + base.I64_extend_i32_u(v30-int32(_a_F_pop_arg_long_double_1))
	} else {
		if base.B2i32(v11|v25 == int64(0))|base.B2i32(v29 != int64(32767)) == int32(0) {
			v135 = v25<<(uint(int64(4))%64) | int64(base.Ui64(v11)>>(uint(int64(60))%64)) | int64(2251799813685248)
			v142 = int64(2047)
		} else {
			if base.Ui32(int32(_a_F_pop_arg_long_double_2)) < base.Ui32(v30) {
				v135 = int64(0)
				v142 = int64(2047)
			} else {
				v84 = base.B2i32(v29 == int64(0))
				if v29 == int64(0) {
					v85 = int32(_a_F_pop_arg_long_double_1)
				} else {
					v85 = int32(_a_F_pop_arg_long_double_0)
				}
				v86 = v85 - v30
				if int32(112) < v86 {
					v89 = int64(0)
					v135 = v89
					v142 = v89
				} else {
					if v29 == int64(0) {
						v93 = v25
					} else {
						v93 = v25 | int64(281474976710656)
					}
					if v30 != v85 {
						F___ashlti3(m, v22+int32(16), v11, v93, int32(128)-v86)
						mBase = m.M
						v101 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v22)+24))
						v106 = base.B2i32(v101|v102 != int64(0))
					} else {
						v106 = int32(0)
					}
					F___lshrti3(m, v22, v11, v93, v86)
					mBase = m.M
					v108 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
					v111 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					v114 = v108<<(uint(int64(4))%64) | int64(base.Ui64(v111)>>(uint(int64(60))%64))
					v118 = base.I64_extend_i32_u(v106) | v111&int64(1152921504606846975)
					if base.Ui64(int64(576460752303423489)) <= base.Ui64(v118) {
						v128 = v114 + int64(1)
					} else {
						if v118 != int64(576460752303423488) {
							v128 = v114
						} else {
							v128 = v114&int64(1) + v114
						}
					}
					v132 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v128))
					if base.Ui64(int64(4503599627370495)) < base.Ui64(v128) {
						v133 = v128 ^ int64(4503599627370496)
					} else {
						v133 = v128
					}
					v135 = v133
					v142 = base.I64_extend_i32_u(v132)
				}
			}
		}
	}
	m.G0 = v22 + int32(32)
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = base.F64_reinterpret_i64(v12&int64(-9223372036854775807-1) | v142<<(uint(int64(52))%64) | v135)
	return
}
func F_positionsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_postgresql_fdw_validator(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
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
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
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
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v579 int32
	_ = v579
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_untransformRelOptions(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errfinish(m, int32(_a_F_postgresql_fdw_validator_0), int32(665), int32(_a_F_postgresql_fdw_validator_1))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L4
	} else {
		goto L163
	}
L2:
	;
	v579 = v8 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(v579)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v579)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v579)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v579))) = v41
	goto L158
L3:
	;
	m.G0 = v10 - int32(-64)
	return int32(1)
L4:
	;
	return int32(0)
L5:
	;
	if v13 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v19 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v19
	goto L10
L9:
	;
	v25 = v22
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = int32(0)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26+v31<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	switch v27 - int32(1417) {
	case 0:
		goto L16
	case 1:
		goto L15
	default:
		goto L2
	}
L12:
	;
	goto L3
L13:
	;
	v564 = v31 + int32(1)
	if v564 != v25 {
		v31 = v564
		goto L11
	} else {
		goto L157
	}
L14:
	;
	v535 = v8 + int32(-16)
	F_updateClosestMatch(m, v535, v532)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L150
	}
L15:
	;
	v463 = int32(_a_F_postgresql_fdw_validator_2)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[0])))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v466 == int32(0))|base.B2i32(v466 != v469) != 0 {
		v487 = v466
		v488 = v469
		goto L134
	} else {
		goto L135
	}
L16:
	;
	v42 = int32(_a_F_postgresql_fdw_validator_3)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[1])))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v45 == int32(0))|base.B2i32(v45 != v48) != 0 {
		v66 = v45
		v67 = v48
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v66-v67 == int32(0) {
		goto L13
	} else {
		goto L24
	}
L18:
	;
	goto L17
L19:
	;
	v51 = v42
	v52 = v41
	goto L20
L20:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v56
		v67 = v55
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v66 = v56
	v67 = v55
	goto L18
L22:
	;
	v59 = int32(1)
	if v56 == v55 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v71 = int32(_a_F_postgresql_fdw_validator_4)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[2])))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v74 == int32(0))|base.B2i32(v74 != v77) != 0 {
		v95 = v74
		v96 = v77
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v95-v96 == int32(0) {
		goto L13
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v80 = v71
	v81 = v41
	goto L28
L28:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v85
		v96 = v84
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v95 = v85
	v96 = v84
	goto L26
L30:
	;
	v88 = int32(1)
	if v85 == v84 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v100 = int32(_a_F_postgresql_fdw_validator_5)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[3])))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v103 == int32(0))|base.B2i32(v103 != v106) != 0 {
		v124 = v103
		v125 = v106
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v124-v125 == int32(0) {
		goto L13
	} else {
		goto L40
	}
L34:
	;
	goto L33
L35:
	;
	v109 = v100
	v110 = v41
	goto L36
L36:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v114 == int32(0) {
		v124 = v114
		v125 = v113
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v124 = v114
	v125 = v113
	goto L34
L38:
	;
	v117 = int32(1)
	if v114 == v113 {
		v109 = v109 + v117
		v110 = v110 + v117
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v129 = int32(_a_F_postgresql_fdw_validator_6)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[4])))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v132 == int32(0))|base.B2i32(v132 != v135) != 0 {
		v153 = v132
		v154 = v135
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v153-v154 == int32(0) {
		goto L13
	} else {
		goto L48
	}
L42:
	;
	goto L41
L43:
	;
	v138 = v129
	v139 = v41
	goto L44
L44:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v143 == int32(0) {
		v153 = v143
		v154 = v142
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v153 = v143
	v154 = v142
	goto L42
L46:
	;
	v146 = int32(1)
	if v143 == v142 {
		v138 = v138 + v146
		v139 = v139 + v146
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v158 = int32(_a_F_postgresql_fdw_validator_7)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[5])))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v161 == int32(0))|base.B2i32(v161 != v164) != 0 {
		v182 = v161
		v183 = v164
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v182-v183 == int32(0) {
		goto L13
	} else {
		goto L56
	}
L50:
	;
	goto L49
L51:
	;
	v167 = v158
	v168 = v41
	goto L52
L52:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
	if v172 == int32(0) {
		v182 = v172
		v183 = v171
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v182 = v172
	v183 = v171
	goto L50
L54:
	;
	v175 = int32(1)
	if v172 == v171 {
		v167 = v167 + v175
		v168 = v168 + v175
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v187 = int32(_a_F_postgresql_fdw_validator_8)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[6])))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v190 == int32(0))|base.B2i32(v190 != v193) != 0 {
		v211 = v190
		v212 = v193
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v211-v212 == int32(0) {
		goto L13
	} else {
		goto L64
	}
L58:
	;
	goto L57
L59:
	;
	v196 = v187
	v197 = v41
	goto L60
L60:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	if v201 == int32(0) {
		v211 = v201
		v212 = v200
		goto L58
	} else {
		goto L62
	}
L61:
	;
	v211 = v201
	v212 = v200
	goto L58
L62:
	;
	v204 = int32(1)
	if v201 == v200 {
		v196 = v196 + v204
		v197 = v197 + v204
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v216 = int32(_a_F_postgresql_fdw_validator_9)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[7])))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v219 == int32(0))|base.B2i32(v219 != v222) != 0 {
		v240 = v219
		v241 = v222
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v240-v241 == int32(0) {
		goto L13
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	v225 = v216
	v226 = v41
	goto L68
L68:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	if v230 == int32(0) {
		v240 = v230
		v241 = v229
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v240 = v230
	v241 = v229
	goto L66
L70:
	;
	v233 = int32(1)
	if v230 == v229 {
		v225 = v225 + v233
		v226 = v226 + v233
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v245 = int32(_a_F_postgresql_fdw_validator_10)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[8])))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v248 == int32(0))|base.B2i32(v248 != v251) != 0 {
		v269 = v248
		v270 = v251
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v269-v270 == int32(0) {
		goto L13
	} else {
		goto L80
	}
L74:
	;
	goto L73
L75:
	;
	v254 = v245
	v255 = v41
	goto L76
L76:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	if v259 == int32(0) {
		v269 = v259
		v270 = v258
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v269 = v259
	v270 = v258
	goto L74
L78:
	;
	v262 = int32(1)
	if v259 == v258 {
		v254 = v254 + v262
		v255 = v255 + v262
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v274 = int32(_a_F_postgresql_fdw_validator_11)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[9])))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v277 == int32(0))|base.B2i32(v277 != v280) != 0 {
		v298 = v277
		v299 = v280
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v298-v299 == int32(0) {
		goto L13
	} else {
		goto L88
	}
L82:
	;
	goto L81
L83:
	;
	v283 = v274
	v284 = v41
	goto L84
L84:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if v288 == int32(0) {
		v298 = v288
		v299 = v287
		goto L82
	} else {
		goto L86
	}
L85:
	;
	v298 = v288
	v299 = v287
	goto L82
L86:
	;
	v291 = int32(1)
	if v288 == v287 {
		v283 = v283 + v291
		v284 = v284 + v291
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v303 = int32(_a_F_postgresql_fdw_validator_12)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[10])))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v306 == int32(0))|base.B2i32(v306 != v309) != 0 {
		v327 = v306
		v328 = v309
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v327-v328 == int32(0) {
		goto L13
	} else {
		goto L96
	}
L90:
	;
	goto L89
L91:
	;
	v312 = v303
	v313 = v41
	goto L92
L92:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	if v317 == int32(0) {
		v327 = v317
		v328 = v316
		goto L90
	} else {
		goto L94
	}
L93:
	;
	v327 = v317
	v328 = v316
	goto L90
L94:
	;
	v320 = int32(1)
	if v317 == v316 {
		v312 = v312 + v320
		v313 = v313 + v320
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v332 = int32(_a_F_postgresql_fdw_validator_13)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[11])))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v335 == int32(0))|base.B2i32(v335 != v338) != 0 {
		v356 = v335
		v357 = v338
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v356-v357 == int32(0) {
		goto L13
	} else {
		goto L104
	}
L98:
	;
	goto L97
L99:
	;
	v341 = v332
	v342 = v41
	goto L100
L100:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+1)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	if v346 == int32(0) {
		v356 = v346
		v357 = v345
		goto L98
	} else {
		goto L102
	}
L101:
	;
	v356 = v346
	v357 = v345
	goto L98
L102:
	;
	v349 = int32(1)
	if v346 == v345 {
		v341 = v341 + v349
		v342 = v342 + v349
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v361 = int32(_a_F_postgresql_fdw_validator_14)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[12])))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v364 == int32(0))|base.B2i32(v364 != v367) != 0 {
		v385 = v364
		v386 = v367
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v385-v386 == int32(0) {
		goto L13
	} else {
		goto L112
	}
L106:
	;
	goto L105
L107:
	;
	v370 = v361
	v371 = v41
	goto L108
L108:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
	if v375 == int32(0) {
		v385 = v375
		v386 = v374
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v385 = v375
	v386 = v374
	goto L106
L110:
	;
	v378 = int32(1)
	if v375 == v374 {
		v370 = v370 + v378
		v371 = v371 + v378
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v390 = int32(_a_F_postgresql_fdw_validator_15)
	v393 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[13])))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v393 == int32(0))|base.B2i32(v393 != v396) != 0 {
		v414 = v393
		v415 = v396
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if v414-v415 == int32(0) {
		goto L13
	} else {
		goto L120
	}
L114:
	;
	goto L113
L115:
	;
	v399 = v390
	v400 = v41
	goto L116
L116:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+1)))
	if v404 == int32(0) {
		v414 = v404
		v415 = v403
		goto L114
	} else {
		goto L118
	}
L117:
	;
	v414 = v404
	v415 = v403
	goto L114
L118:
	;
	v407 = int32(1)
	if v404 == v403 {
		v399 = v399 + v407
		v400 = v400 + v407
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v420 = v8 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(v420)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v420)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v420)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v420))) = v41
	goto L121
L121:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_3))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_4))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_5))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_6))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_7))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_8))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_9))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_10))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_11))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_12))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	F_updateClosestMatch(m, v420, int32(_a_F_postgresql_fdw_validator_13))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v532 = int32(_a_F_postgresql_fdw_validator_14)
	v533 = int32(_a_F_postgresql_fdw_validator_15)
	goto L14
L133:
	;
	if v487-v488 == int32(0) {
		goto L13
	} else {
		goto L140
	}
L134:
	;
	goto L133
L135:
	;
	v472 = v463
	v473 = v41
	goto L136
L136:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+1)))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+1)))
	if v477 == int32(0) {
		v487 = v477
		v488 = v476
		goto L134
	} else {
		goto L138
	}
L137:
	;
	v487 = v477
	v488 = v476
	goto L134
L138:
	;
	v480 = int32(1)
	if v477 == v476 {
		v472 = v472 + v480
		v473 = v473 + v480
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v492 = int32(_a_F_postgresql_fdw_validator_16)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postgresql_fdw_validator[14])))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if base.B2i32(v495 == int32(0))|base.B2i32(v495 != v498) != 0 {
		v516 = v495
		v517 = v498
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v516-v517 == int32(0) {
		goto L13
	} else {
		goto L148
	}
L142:
	;
	goto L141
L143:
	;
	v501 = v492
	v502 = v41
	goto L144
L144:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+1)))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+1)))
	if v506 == int32(0) {
		v516 = v506
		v517 = v505
		goto L142
	} else {
		goto L146
	}
L145:
	;
	v516 = v506
	v517 = v505
	goto L142
L146:
	;
	v509 = int32(1)
	if v506 == v505 {
		v501 = v501 + v509
		v502 = v502 + v509
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v522 = v8 + int32(-16)
	*(*int32)(unsafe.Add(mBase, uint32(v522)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v522)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v522)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v522))) = v41
	goto L149
L149:
	;
	v532 = int32(_a_F_postgresql_fdw_validator_2)
	v533 = int32(_a_F_postgresql_fdw_validator_16)
	goto L14
L150:
	;
	F_updateClosestMatch(m, v535, v533)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v548
	F_errmsg(m, int32(_a_F_postgresql_fdw_validator_17), v8+int32(-32))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	if v540 == int32(0) {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v540
	F_errhint(m, int32(_a_F_postgresql_fdw_validator_18), v8+int32(-48))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	goto L1
L157:
	;
	goto L12
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v595
	F_errmsg(m, int32(_a_F_postgresql_fdw_validator_17), v10)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errhint(m, int32(_a_F_postgresql_fdw_validator_19), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	goto L1
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pre_sync_fname(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l1 != 0 {
		m.G0 = v7 + int32(48)
		return
	} else {
		v16 = m.G0
		v18 = v16 - int32(16)
		m.G0 = v18
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[0]))
		if v21 != 0 {
			v22 = F_GetCurrentTimestamp(m)
			mBase = m.M
			v24 = *(*int64)(unsafe.Add(mBase, _c_F_pre_sync_fname[1]))
			F_TimestampDifference(m, v24, v22, v18+int32(12), v18+int32(8))
			mBase = m.M
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(44)))) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7+int32(40)))) = v32
			*(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[0])) = int32(0)
		} else {
		}
		m.G0 = v18 + int32(16)
		if base.B2i32(v21 != int32(0)) == int32(0) {
			v70 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[2]))
			v71 = F_OpenTransientFilePerm(m, l0, int32(0), v70)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				if v71 < int32(0) {
					v76 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[3]))
					if v76 == int32(2) {
						m.G0 = v7 + int32(48)
						return
					} else {
						v80 = F_errstart(m, l2, int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							if v80 == int32(0) {
								m.G0 = v7 + int32(48)
								return
							} else {
								v98 = int32(_a_F_pre_sync_fname_0)
								v99 = int32(3810)
								F_errcode_for_file_access(m)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
									F_errmsg(m, v98, v7)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_pre_sync_fname_1), v99, int32(_a_F_pre_sync_fname_2))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											m.G0 = v7 + int32(48)
											return
										}
									}
								}
							}
						}
					}
				} else {
					v86 = F_fsync(m, v71)
					mBase = m.M
					v87 = F_CloseTransientFile(m, v71)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						if v87 == int32(0) {
							m.G0 = v7 + int32(48)
							return
						} else {
							v92 = F_errstart(m, l2, int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								if v92 == int32(0) {
									m.G0 = v7 + int32(48)
									return
								} else {
									v98 = int32(_a_F_pre_sync_fname_3)
									v99 = int32(3823)
									F_errcode_for_file_access(m)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
										F_errmsg(m, v98, v7)
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_pre_sync_fname_1), v99, int32(_a_F_pre_sync_fname_2))
											mBase = m.M
											v108 = m.ExcPending
											if v108 != 0 {
												return
											} else {
												m.G0 = v7 + int32(48)
												return
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
			v47 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				if v47 == int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[2]))
					v71 = F_OpenTransientFilePerm(m, l0, int32(0), v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						if v71 < int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[3]))
							if v76 == int32(2) {
								m.G0 = v7 + int32(48)
								return
							} else {
								v80 = F_errstart(m, l2, int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									if v80 == int32(0) {
										m.G0 = v7 + int32(48)
										return
									} else {
										v98 = int32(_a_F_pre_sync_fname_0)
										v99 = int32(3810)
										F_errcode_for_file_access(m)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
											F_errmsg(m, v98, v7)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_pre_sync_fname_1), v99, int32(_a_F_pre_sync_fname_2))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													m.G0 = v7 + int32(48)
													return
												}
											}
										}
									}
								}
							}
						} else {
							v86 = F_fsync(m, v71)
							mBase = m.M
							v87 = F_CloseTransientFile(m, v71)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								if v87 == int32(0) {
									m.G0 = v7 + int32(48)
									return
								} else {
									v92 = F_errstart(m, l2, int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										if v92 == int32(0) {
											m.G0 = v7 + int32(48)
											return
										} else {
											v98 = int32(_a_F_pre_sync_fname_3)
											v99 = int32(3823)
											F_errcode_for_file_access(m)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
												F_errmsg(m, v98, v7)
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_pre_sync_fname_1), v99, int32(_a_F_pre_sync_fname_2))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														m.G0 = v7 + int32(48)
														return
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
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = l0
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
					v56 = base.I32_div_s(v54, int32(_a_F_pre_sync_fname_4))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v56
					F_errmsg(m, int32(_a_F_pre_sync_fname_5), v7+int32(16))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pre_sync_fname_1), int32(3800), int32(_a_F_pre_sync_fname_2))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[2]))
							v71 = F_OpenTransientFilePerm(m, l0, int32(0), v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								if v71 < int32(0) {
									v76 = *(*int32)(unsafe.Add(mBase, _c_F_pre_sync_fname[3]))
									if v76 == int32(2) {
										m.G0 = v7 + int32(48)
										return
									} else {
										v80 = F_errstart(m, l2, int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											if v80 == int32(0) {
												m.G0 = v7 + int32(48)
												return
											} else {
												v98 = int32(_a_F_pre_sync_fname_0)
												v99 = int32(3810)
												F_errcode_for_file_access(m)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
													F_errmsg(m, v98, v7)
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_pre_sync_fname_1), v99, int32(_a_F_pre_sync_fname_2))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															m.G0 = v7 + int32(48)
															return
														}
													}
												}
											}
										}
									}
								} else {
									v86 = F_fsync(m, v71)
									mBase = m.M
									v87 = F_CloseTransientFile(m, v71)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										if v87 == int32(0) {
											m.G0 = v7 + int32(48)
											return
										} else {
											v92 = F_errstart(m, l2, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												if v92 == int32(0) {
													m.G0 = v7 + int32(48)
													return
												} else {
													v98 = int32(_a_F_pre_sync_fname_3)
													v99 = int32(3823)
													F_errcode_for_file_access(m)
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
														F_errmsg(m, v98, v7)
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_pre_sync_fname_1), v99, int32(_a_F_pre_sync_fname_2))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																m.G0 = v7 + int32(48)
																return
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
}
func F_preprocess_aggrefs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
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
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v609 int32
	_ = v609
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v670 int32
	_ = v670
	v3 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(432)
	m.G0 = v25
	if l0 == v3 {
		v670 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v25 + int32(432)
	return v670
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v30 == int32(9) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v628
	v670 = int32(0)
	goto L1
L4:
	;
	v303 = F_palloc0(m, int32(20))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L9
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L9
	} else {
		goto L72
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = F_SearchSysCache1(m, int32(0), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v264 = F_expression_tree_walker_impl(m, l0, int32(850), l1)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L71
	}
L9:
	;
	return int32(0)
L10:
	;
	if v35 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v52 = v25 + int32(16)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v88 = F_resolve_aggregate_transtype(m, v87, v50, v52)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L22
	}
L13:
	;
	goto L12
L14:
	;
	goto L15
L15:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v59 < v60 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = v59
	goto L19
L17:
	;
	goto L18
L18:
	;
	goto L12
L19:
	;
	v68 = v63 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v68))) = v72
	v75 = v63 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v75 < v76 {
		v63 = v75
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v88
	v91 = int32(-1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v92 == int32(0) {
		v105 = v91
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+42)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_get_typlenbyval(m, v107, v25+int32(422), v25+int32(421))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L28
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v98 = F_exprType(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	if v98 != v88 {
		v105 = v91
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v102 = F_exprTypmod(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v105 = v102
	goto L23
L28:
	;
	v118 = F_SysCacheGetAttr(m, int32(0), v35, int32(21), v25+int32(420))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+420)))
	if v120 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_getTypeInputInfo(m, v88, v25+int32(428), v25+int32(424))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	v139 = v3
	goto L32
L32:
	;
	F_ReleaseCatCache(m, v35)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L37
	}
L33:
	;
	v129 = F_text_to_cstring(m, v118)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v25)+428))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v25)+424))
	v134 = F_OidInputFunctionCall(m, v131, v129, v132, int32(-1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v129)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v139 = v134
	goto L32
L37:
	;
	v142 = F_contain_volatile_functions(m, l0)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	if v142 != 0 {
		v288 = v3
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	if v144 == int32(0) {
		v288 = v3
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v147 <= int32(0) {
		v288 = v3
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v155 = int32(0)
	v157 = int32(-1)
	v160 = v3
	goto L42
L42:
	;
	v175 = v157 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v155<<(uint(int32(2))%32))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	if v176 != v185 {
		v242 = v160
		goto L45
	} else {
		goto L46
	}
L43:
	;
	F_list_free(m, v160)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L68
	}
L44:
	;
	goto L43
L45:
	;
	v244 = v155 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v244 < v245 {
		v155 = v244
		v157 = v175
		v160 = v242
		goto L42
	} else {
		goto L67
	}
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+20))
	if v187 != v188 {
		v242 = v160
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+48)))
	if v190 != v191 {
		v242 = v160
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+49)))
	if v193 != v194 {
		v242 = v160
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+50)))
	if v196 != v197 {
		v242 = v160
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v184)+32))
	v201 = F_equal(m, v199, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	if v201 == int32(0) {
		v242 = v160
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v184)+36))
	v207 = F_equal(m, v205, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	if v207 == int32(0) {
		v242 = v160
		goto L45
	} else {
		goto L54
	}
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v184)+40))
	v213 = F_equal(m, v211, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	if v213 == int32(0) {
		v242 = v160
		goto L45
	} else {
		goto L56
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v184)+44))
	v219 = F_equal(m, v217, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	if v219 == int32(0) {
		v242 = v160
		goto L45
	} else {
		goto L58
	}
L58:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v223 != v224 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+12)))
	if v236 != int32(1) {
		v242 = v160
		goto L45
	} else {
		goto L65
	}
L60:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	if v226 != v227 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	if v229 != v230 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v184)+28))
	v234 = F_equal(m, v232, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	if v234 != 0 {
		goto L44
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	v240 = F_lappend_int(m, v160, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v242 = v240
	goto L45
L67:
	;
	v288 = v242
	goto L4
L68:
	;
	if v175 == int32(-1) {
		v288 = int32(0)
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253+v175<<(uint(int32(2))%32))))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v259 = F_lappend(m, v258, l0)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+4)) = v259
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v628 = v175
	v632 = v262
	goto L3
L71:
	;
	v670 = v264
	goto L1
L72:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v270
	F_errmsg_internal(m, int32(_a_F_preprocess_aggrefs_walker_0), v25)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_preprocess_aggrefs_walker_1), int32(153), int32(_a_F_preprocess_aggrefs_walker_2))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+16)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v303))) = int32(327)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v25)+428)) = l0
	v313 = F_list_make1_impl(m, int32(1), v25+int32(12))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v303)+12)) = uint8(base.B2i32(v106 != int32(119)))
	*(*int32)(unsafe.Add(mBase, uint32(v303)+4)) = v313
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+328))
	if v319 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v322 = v320
	goto L79
L78:
	;
	v322 = int32(0)
	goto L79
L79:
	;
	v323 = F_lappend(m, v319, v303)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+328)) = v323
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v326 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_get_typlenbyval(m, v88, v25+int32(424), v25+int32(419))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L9
	} else {
		goto L86
	}
L82:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v329 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v332 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+340)) = uint8(v332)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+336))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+336)) = v334 + v332
	goto L81
L85:
	;
	goto L84
L86:
	;
	if base.B2i32(v288 == int32(0))|base.B2i32(v106 == int32(119)) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+8)) = v609
	v628 = v322
	v632 = v609
	goto L3
L88:
	;
	v444 = F_palloc0(m, int32(56))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L109
	}
L89:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v349 <= int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+424)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+420)))
	v355 = int32(1)
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+419)))
	v363 = int32(0)
	v372 = v349
	goto L91
L91:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+332))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v385 = int32(2)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v363<<(uint(v385)%32))))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v383+v388<<(uint(v385)%32))))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	if v49 != v393 {
		v414 = v372
		goto L94
	} else {
		goto L95
	}
L92:
	;
	if v388 != int32(-1) {
		v609 = v388
		goto L87
	} else {
		goto L108
	}
L93:
	;
	goto L92
L94:
	;
	v417 = v363 + int32(1)
	if v417 < v414 {
		v363 = v417
		v372 = v414
		goto L91
	} else {
		goto L107
	}
L95:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)+28))
	if v88 != v395 {
		v414 = v372
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v392)+16))
	if v46 != v397 {
		v414 = v372
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v392)+20))
	if v45 != v399 {
		v414 = v372
		goto L94
	} else {
		goto L98
	}
L98:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v392)+24))
	if v47 != v401 {
		v414 = v372
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+52)))
	if v354&v355 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if v403&int32(1) == int32(0) {
		v414 = v372
		goto L94
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if v403&int32(1) != 0 {
		v414 = v372
		goto L94
	} else {
		goto L104
	}
L103:
	;
	goto L93
L104:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v392)+48))
	v411 = F_datumIsEqual(m, v139, v410, v357&v355, v352)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	if v411 != 0 {
		goto L93
	} else {
		goto L106
	}
L106:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v414 = v413
	goto L94
L107:
	;
	goto L88
L108:
	;
	goto L88
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = int32(328)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+4)) = v448
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v444)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v444)+8)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v444)+32)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v444)+28)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v444)+20)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v444)+16)) = v46
	v458 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+424)))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+36)) = v458
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+419)))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+48)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v444)+44)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+40)) = uint8(v460)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+420)))
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+52)) = uint8(v464)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l1)+332))
	if v466 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+4))
	v469 = v467
	goto L112
L111:
	;
	v469 = int32(0)
	goto L112
L112:
	;
	v470 = F_lappend(m, v466, v444)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+332)) = v470
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+340)))
	if v473 != 0 {
		v609 = v469
		goto L87
	} else {
		goto L114
	}
L114:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v444)+24))
	if v474 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+340)) = uint8(v477)
	v609 = v469
	goto L87
L116:
	;
	goto L117
L117:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v444)+28))
	if v479 != int32(2281) {
		v609 = v469
		goto L87
	} else {
		goto L118
	}
L118:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v444)+16))
	if v482 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v487 != int32(_a_F_preprocess_aggrefs_walker_3) {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v444)+20))
	if v483 != 0 {
		v487 = v482
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v484 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+341)) = uint8(v484)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v444)+16))
	v487 = v486
	goto L119
L123:
	;
	goto L122
L124:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v444)+20))
	if v490 != int32(_a_F_preprocess_aggrefs_walker_4) {
		v609 = v469
		goto L87
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v493 = int32(0)
	v494 = m.G0
	v496 = v494 - int32(16)
	m.G0 = v496
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v498 == v493 {
		goto L131
	} else {
		goto L132
	}
L127:
	;
	goto L126
L128:
	;
	if v563 != 0 {
		v609 = v469
		goto L87
	} else {
		goto L153
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L9
	} else {
		goto L150
	}
L130:
	;
	m.G0 = v496 + int32(16)
	goto L128
L131:
	;
	v563 = int32(1)
	goto L130
L132:
	;
	goto L133
L133:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v503 <= int32(0) {
		v563 = int32(1)
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v508 = v493
	goto L135
L135:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528+v508<<(uint(int32(2))%32))))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	v534 = F_exprType(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L9
	} else {
		goto L137
	}
L136:
	;
	v563 = v537
	goto L130
L137:
	;
	v536 = int32(2249)
	v537 = base.B2i32(v534 != v536)
	if v534 == v536 {
		v563 = v537
		goto L130
	} else {
		goto L138
	}
L138:
	;
	v541 = F_SearchSysCache1(m, int32(82), v534)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L9
	} else {
		goto L139
	}
L139:
	;
	if v541 == int32(0) {
		goto L129
	} else {
		goto L140
	}
L140:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541)+16))
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+22)))
	v547 = v545 + v546
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+78)))
	if v548 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_ReleaseCatCache(m, v541)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L148
	}
L142:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v547)+112))
	if v549 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v547)+108))
	if v550 != 0 {
		goto L141
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	F_ReleaseCatCache(m, v541)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L9
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	v563 = int32(0)
	goto L130
L148:
	;
	v557 = v508 + int32(1)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v557 < v558 {
		v508 = v557
		goto L135
	} else {
		goto L149
	}
L149:
	;
	goto L136
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = v534
	F_errmsg_internal(m, int32(_a_F_preprocess_aggrefs_walker_5), v496)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L9
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_preprocess_aggrefs_walker_6), int32(2134), int32(_a_F_preprocess_aggrefs_walker_7))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L9
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	v598 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+341)) = uint8(v598)
	v609 = v469
	goto L87
}
func F_printsimple_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_pq_beginmessage(m, v9, int32(84))
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
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	F_enlargeStringInfo(m, v9, int32(2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v21 = int32(8)
	v25 = v14<<(uint(v21)%32) | int32(base.Ui32(v14)>>(uint(v21)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v18+v19))) = uint16(v25)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v18 + int32(2)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v30 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = v30
	v38 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pq_endmessage(m, v9)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v45 = l2 + v35<<(uint(int32(4))%32) + v38*int32(100)
	F_pq_sendstring(m, v9, v45+int32(24))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	F_enlargeStringInfo(m, v9, int32(4))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v53+v54))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v53 + int32(4)
	F_enlargeStringInfo(m, v9, int32(2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v67 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v64+v65))) = uint16(v67)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v64 + int32(2)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)+88))
	F_enlargeStringInfo(m, v9, int32(4))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v81 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v76+v77))) = base.I32_rotr(v72, int32(24))&v81 | base.I32_rotr(v72&v81, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v76 + int32(4)
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+92)))
	F_enlargeStringInfo(m, v9, int32(2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v99 = int32(8)
	v103 = v92<<(uint(v99)%32) | int32(base.Ui32(v92)>>(uint(v99)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v96+v97))) = uint16(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v96 + int32(2)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	F_enlargeStringInfo(m, v9, int32(4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v117 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v112+v113))) = base.I32_rotr(v108, int32(24))&v117 | base.I32_rotr(v108&v117, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v112 + int32(4)
	F_enlargeStringInfo(m, v9, int32(2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v134 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v131+v132))) = uint16(v134)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v131 + int32(2)
	v140 = v38 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v140 < v141 {
		v35 = v141
		v38 = v140
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
L17:
	;
	m.G0 = v9 + int32(16)
	return
}
func F_processIndirection(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L17
	} else {
		goto L30
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v76
L3:
	;
	v76 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = l0
	v19 = v3
	goto L8
L6:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v70 == v73 {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	if v67 == int32(0) {
		v76 = v65
		goto L2
	} else {
		goto L26
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v21 - int32(14) {
	case 0:
		goto L12
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		v65 = v15
		v67 = v19
		goto L7
	case 12:
		goto L13
	default:
		goto L14
	}
L9:
	;
	v65 = int32(0)
	v67 = v62
	goto L7
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v63 != 0 {
		v15 = v63
		v19 = v62
		goto L8
	} else {
		goto L25
	}
L11:
	;
	v61 = v15 + int32(4)
	v62 = v15
	goto L10
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v52 == int32(0) {
		v65 = v15
		v67 = v19
		goto L7
	} else {
		goto L23
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v30 = F_get_typ_typrelid(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v21 != int32(55) {
		v65 = v15
		v67 = v19
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	if v26 == int32(2) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v70 = v15
	v72 = v15
	goto L6
L17:
	;
	return int32(0)
L18:
	;
	if v30 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37))))
	v40 = F_get_attname(m, v30, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v42 = F_quote_identifier(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v42
	F_appendStringInfo(m, v14, int32(_a_F_processIndirection_0), v9+int32(16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v61 = v51
	v62 = v19
	goto L10
L23:
	;
	F_printSubscripts(m, v15, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v61 = v15 + int32(36)
	v62 = v19
	goto L10
L25:
	;
	goto L9
L26:
	;
	v70 = v65
	v72 = v67
	goto L6
L27:
	;
	v75 = v72
	goto L29
L28:
	;
	v75 = v70
	goto L29
L29:
	;
	v76 = v75
	goto L2
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v91 = F_format_type_be(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v91
	F_errmsg_internal(m, int32(_a_F_processIndirection_1), v9)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_processIndirection_2), int32(_a_F_processIndirection_3), int32(_a_F_processIndirection_4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_processTypesSpec(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_typenameTypeId(m, int32(0), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if int32(2) <= v10 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v16 = F_typenameTypeId(m, int32(0), v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = v16
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if int32(3) <= v20 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_processTypesSpec_0), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_processTypesSpec_1), int32(1128), int32(_a_F_processTypesSpec_2))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
					return
				}
			}
		} else {
			v18 = v7
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if int32(3) <= v20 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_processTypesSpec_0), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_processTypesSpec_1), int32(1128), int32(_a_F_processTypesSpec_2))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
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
				return
			}
		}
	}
}
func F_process_implied_equality(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	v10 = F_copyObjectImpl(m, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = F_copyObjectImpl(m, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = F_make_opclause(m, l1, v10, v14, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l7 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return v249
L6:
	;
	v33 = F_pull_varnos(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v32 = v16
	goto L6
L8:
	;
	goto L9
L9:
	;
	v20 = F_eval_const_expressions(m, l0, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v20 == int32(0) {
		v32 = int32(0)
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 != int32(7) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = v20
	goto L6
L13:
	;
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	if v27 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v32 = v20
	goto L6
L16:
	;
	goto L17
L17:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v29 != 0 {
		v249 = int32(0)
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v32 = v20
	goto L6
L19:
	;
	if v33 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v37 = F_bms_copy(m, l5)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v143 = v33
	goto L22
L22:
	;
	v152 = int32(0)
	v158 = F_make_restrictinfo(m, l0, v32, int32(1), v152, v152, base.B2i32(v33 == v152), l6, v143, v152, v152)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L48
	}
L23:
	;
	v140 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+319)) = uint8(v140)
	v143 = v132
	goto L22
L24:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v40 = int32(0)
	if base.B2i32(v37 == v40)|base.B2i32(v39 == v40) != 0 {
		v86 = base.B2i32(v37|v39 == v40)
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v86 != 0 {
		v132 = v37
		goto L23
	} else {
		goto L36
	}
L26:
	;
	goto L25
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v54 != v55 {
		v86 = int32(0)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v57 = int32(1)
	if v54 <= v57 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v60 = v57
	goto L31
L30:
	;
	v60 = v54
	goto L31
L31:
	;
	v61 = int32(8)
	v66 = int32(0)
	goto L32
L32:
	;
	v74 = v66 << (uint(int32(2)) % 32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v37+v61+v74)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v39+v61+v74)))
	v79 = base.B2i32(v76 == v78)
	if v76 != v78 {
		v86 = v79
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v86 = v79
	goto L26
L34:
	;
	v82 = v66 + int32(1)
	if v82 != v60 {
		v66 = v82
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v91 == int32(0) {
		v132 = v37
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v94 <= int32(0) {
		v132 = v37
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v99 = v37
	v100 = int32(0)
	goto L39
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v100<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	if v112 != int32(1) {
		v126 = v99
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v132 = v126
	goto L23
L41:
	;
	v128 = v100 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v128 < v129 {
		v99 = v126
		v100 = v128
		goto L39
	} else {
		goto L47
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	v116 = F_bms_is_member(m, v115, v99)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v116 == int32(0) {
		v126 = v99
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	v121 = F_bms_del_member(m, v99, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v124 = F_bms_del_members(m, v121, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v126 = v124
	goto L41
L47:
	;
	goto L40
L48:
	;
	v160 = int32(0)
	if v143 == v160 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v205 == int32(2) {
		goto L65
	} else {
		goto L66
	}
L50:
	;
	v205 = int32(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v168 = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v169 <= v168 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v172 = v168
	goto L55
L54:
	;
	v172 = v169
	goto L55
L55:
	;
	v176 = int32(0)
	v178 = v160
	goto L56
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v143+int32(8)+v176<<(uint(int32(2))%32))))
	if v185 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v205 = v197
	goto L49
L58:
	;
	goto L57
L59:
	;
	v186 = int32(2)
	if v178 != 0 {
		v197 = v186
		goto L58
	} else {
		goto L62
	}
L60:
	;
	v192 = v178
	goto L61
L61:
	;
	v194 = v176 + int32(1)
	if v194 != v172 {
		v176 = v194
		v178 = v192
		goto L56
	} else {
		goto L64
	}
L62:
	;
	v187 = int32(1)
	if base.Ui32(v187) < base.Ui32(base.I32_popcnt(v185)) {
		v197 = v186
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v192 = v187
	goto L61
L64:
	;
	v197 = v192
	goto L58
L65:
	;
	v209 = F_pull_var_clause(m, v32, int32(26))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+10)))
	if v216 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	F_add_vars_to_targetlist(m, l0, v209, v143)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_list_free(m, v209)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v158)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L83
	}
L72:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v217 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v220 != int32(17) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	if v223 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v226 != int32(2) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v232 = F_exprType(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v234 = F_op_mergejoinable(m, v229, v232)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v234 == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v238 = F_contain_volatile_functions(m, v158)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v238 != 0 {
		goto L71
	} else {
		goto L81
	}
L81:
	;
	v240 = F_get_mergejoin_opfamilies(m, v229)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+96)) = v240
	goto L71
L83:
	;
	v249 = v158
	goto L5
}
func F_prsd_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	F_pfree(m, v6)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = v9
	if v9 != 0 {
		v6 = v9
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	F_pfree(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v21 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	F_pfree(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_pfree(m, v4)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	return int32(0)
}
func F_pub_collist_validate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	if l1 == v3 {
		v65 = v3
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L31
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L27
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L23
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L19
	}
L5:
	;
	m.G0 = v11 - int32(-64)
	return v65
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		v65 = v3
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = v3
	v26 = v3
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = F_get_attnum(m, v27, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v65 = v54
	goto L5
L10:
	;
	return int32(0)
L11:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v34 <= int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v42<<(uint(int32(4))%32)+v34*int32(100))+10)))
	if v49 == int32(118) {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v52 = F_bms_is_member(m, v34, v24)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v54 = F_bms_add_member(m, v24, v34)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v57 = v26 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 < v58 {
		v24 = v54
		v26 = v57
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v79 + int32(4)
	F_errmsg(m, int32(_a_F_pub_collist_validate_0), v11)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(571), int32(_a_F_pub_collist_validate_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_errcode(m, int32(_a_F_pub_collist_validate_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v33
	F_errmsg(m, int32(_a_F_pub_collist_validate_4), v9+int32(-16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(577), int32(_a_F_pub_collist_validate_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errcode(m, int32(_a_F_pub_collist_validate_3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v33
	F_errmsg(m, int32(_a_F_pub_collist_validate_5), v9+int32(-48))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(583), int32(_a_F_pub_collist_validate_2))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(_a_F_pub_collist_validate_6))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v33
	F_errmsg(m, int32(_a_F_pub_collist_validate_7), v9+int32(-32))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(589), int32(_a_F_pub_collist_validate_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pull_up_sublinks_qual_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l1 == v7 {
		v384 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v384
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v24 - int32(21) {
	case 0:
		goto L3
	case 1:
		goto L4
	default:
		v384 = l1
		goto L1
	}
L3:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v269 {
	case 0:
		goto L68
	default:
		v384 = l1
		goto L1
	case 2:
		goto L69
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v27 {
	case 0:
		goto L5
	default:
		v384 = l1
		goto L1
	case 2:
		goto L6
	}
L5:
	;
	v221 = int32(0)
	v223 = F_convert_EXISTS_sublink_to_join(m, l0, l1, v221, l3)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L23
	} else {
		goto L55
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v34 != int32(17) {
		v162 = v7
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v32 + int32(16)
	if v162 != 0 {
		v384 = v162
		goto L1
	} else {
		goto L41
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v37 == int32(0) {
		v162 = v7
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 != int32(2) {
		v162 = v7
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if int32(1) < v44 {
		v162 = v7
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+132))
	if v47 != 0 {
		v162 = v7
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+128))
	if v48 != 0 {
		v162 = v7
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	if v49 != 0 {
		v162 = v7
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v50 == int32(0) {
		v162 = v7
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 != int32(1) {
		v162 = v7
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v58 != int32(5) {
		v162 = v7
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v61 == int32(0) {
		v162 = v7
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 < int32(2) {
		v162 = v7
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v72 = F_contain_volatile_functions(m, v61)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	if v72 != 0 {
		v162 = v7
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v76 == int32(0) {
		v137 = v7
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v147 = F_exprType(m, v70)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L23
	} else {
		goto L39
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		v137 = v7
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v90 = v7
	v95 = int32(0)
	goto L29
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v95<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v106
	v110 = F_list_make1_impl(m, int32(1), v32)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v137 = v124
	goto L26
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l0
	v116 = F_convert_testexpr_mutator(m, v70, v32+int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v118 = F_eval_const_expressions(m, l0, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v120 != int32(7) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v162 = int32(0)
	goto L7
L35:
	;
	goto L36
L36:
	;
	v124 = F_lappend(m, v90, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v127 = v95 + int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v127 < v128 {
		v90 = v124
		v95 = v127
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v153 = F_make_SAOP_expr(m, v68, v71, v147, v151, v67, v137, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v162 = v153
	goto L7
L41:
	;
	v175 = F_convert_ANY_sublink_to_join(m, l0, l1, l3)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	if v175 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v175
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v183 = F_pull_up_sublinks_jointree_recurse(m, l0, v180, v20+int32(12))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L23
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if l5 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v192 = F_pull_up_sublinks_qual_recurse(m, l0, v186, v175+int32(12), l3, v175+int32(16), v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+28)) = v192
	v384 = int32(0)
	goto L1
L48:
	;
	v384 = l1
	goto L1
L49:
	;
	goto L50
L50:
	;
	v198 = F_convert_ANY_sublink_to_join(m, l0, l1, l5)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	if v198 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v198
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	v208 = F_pull_up_sublinks_jointree_recurse(m, l0, v205, v20+int32(12))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+16)) = v208
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v198)+28))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v217 = F_pull_up_sublinks_qual_recurse(m, l0, v211, v198+int32(12), l5, v198+int32(16), v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+28)) = v217
	v384 = int32(0)
	goto L1
L55:
	;
	if v223 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v223
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v231 = F_pull_up_sublinks_jointree_recurse(m, l0, v228, v20+int32(12))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L23
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if l5 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v231
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)+28))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v240 = F_pull_up_sublinks_qual_recurse(m, l0, v234, v223+int32(12), l3, v223+int32(16), v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+28)) = v240
	v384 = v221
	goto L1
L61:
	;
	v384 = l1
	goto L1
L62:
	;
	goto L63
L63:
	;
	v246 = F_convert_EXISTS_sublink_to_join(m, l0, l1, int32(0), l5)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	if v246 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v246
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v256 = F_pull_up_sublinks_jointree_recurse(m, l0, v253, v20+int32(12))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+16)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v246)+28))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v265 = F_pull_up_sublinks_qual_recurse(m, l0, v259, v246+int32(12), l5, v246+int32(16), v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+28)) = v265
	v384 = int32(0)
	goto L1
L68:
	;
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v313 == v312 {
		v384 = v312
		goto L1
	} else {
		goto L83
	}
L69:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v275 != int32(22) {
		v384 = l1
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	if v278 != 0 {
		v384 = l1
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v280 = F_convert_EXISTS_sublink_to_join(m, l0, v272, int32(1), l3)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L23
	} else {
		goto L74
	}
L73:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+16))
	v299 = F_pull_up_sublinks_jointree_recurse(m, l0, v296, v20+int32(8))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L23
	} else {
		goto L81
	}
L74:
	;
	if v280 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v280)+12)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v280
	v295 = v280
	goto L73
L76:
	;
	goto L77
L77:
	;
	if l5 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v288 = F_convert_EXISTS_sublink_to_join(m, l0, v272, int32(1), l5)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L23
	} else {
		goto L79
	}
L79:
	;
	if v288 == int32(0) {
		v384 = l1
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+12)) = v292
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v288
	v295 = v288
	goto L73
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v299
	v302 = int32(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v295)+28))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v309 = F_pull_up_sublinks_qual_recurse(m, l0, v303, v295+int32(16), v306, v302, v302)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+28)) = v309
	v384 = v302
	goto L1
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	if int32(0) < v316 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v321 = int32(0)
	v327 = v7
	goto L87
L85:
	;
	v358 = v7
	goto L86
L86:
	;
	if v358 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337+v321<<(uint(int32(2))%32))))
	v342 = F_pull_up_sublinks_qual_recurse(m, l0, v341, l2, l3, l4, l5)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L23
	} else {
		goto L89
	}
L88:
	;
	v358 = v346
	goto L86
L89:
	;
	if v342 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v344 = F_lappend(m, v327, v342)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L23
	} else {
		goto L93
	}
L91:
	;
	v346 = v327
	goto L92
L92:
	;
	v348 = v321 + int32(1)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	if v348 < v349 {
		v321 = v348
		v327 = v346
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v346 = v344
	goto L92
L94:
	;
	goto L88
L95:
	;
	v384 = int32(0)
	goto L1
L96:
	;
	goto L97
L97:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v371 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v384 = v375
	goto L1
L99:
	;
	goto L100
L100:
	;
	v376 = F_make_andclause(m, v358)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L23
	} else {
		goto L101
	}
L101:
	;
	v384 = v376
	goto L1
}
func F_pushf_free_all(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	m.T0[v9].(func(*base.Module, int32))(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v13 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	goto L20
L14:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	F_pfree(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L18
	}
L15:
	;
	base.MemoryFill(m, v13, int32(0), v15)
	goto L17
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L13
L19:
	;
	F_pfree(m, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L23
	}
L20:
	;
	base.MemoryFill(m, v4, int32(0), int32(24))
	goto L22
L22:
	;
	goto L19
L23:
	;
	if v7 != 0 {
		v4 = v7
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L5
}
func F_pwrite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v17 = m.Wasi_snapshot_preview1.Fd_pwrite(m, l0, v8+int32(8), int32(1), l3, v8+int32(4))
	mBase = m.M
	if v17 == int32(0) {
		v24 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pwrite[0])) = v17
		v24 = int32(-1)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	m.G0 = v8 + int32(16)
	if v24 != 0 {
		v30 = int32(-1)
	} else {
		v30 = v25
	}
	return v30
}
func F_pwritev(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = m.Wasi_snapshot_preview1.Fd_pwrite(m, l0, l1, l2, l3, v8+int32(12))
	mBase = m.M
	if v12 == int32(0) {
		v19 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pwritev[0])) = v12
		v19 = int32(-1)
	}
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	m.G0 = v8 + int32(16)
	if v19 != 0 {
		v25 = int32(-1)
	} else {
		v25 = v20
	}
	return v25
}
