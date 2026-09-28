package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	F_errfinish(m, int32(_a_F_PGLC_localeconv_8), int32(534), int32(_a_F_PGLC_localeconv_9))
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
	F_errfinish(m, int32(_a_F_PGLC_localeconv_8), int32(565), int32(_a_F_PGLC_localeconv_9))
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
func F_PMSignalShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_PMSignalShmemInit[0]))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PMSignalShmemInit[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+48)) = v5
	return
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
	F_errfinish(m, int32(_a_F_PersistHoldablePortal_2), int32(471), int32(_a_F_PersistHoldablePortal_3))
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
						F_errfinish(m, int32(_a_F_PrefetchBuffer_1), int32(798), int32(_a_F_PrefetchBuffer_2))
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
func F_PrepareRedoRemove(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	if base.Ui32(l0) <= base.Ui32(int32(2)) {
		F_PrepareRedoRemoveFull(m, base.I64_extend_i32_u(l0), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareRedoRemove[0]))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
		v15 = int64(base.Ui64(v13) >> (uint(int64(32)) % 64))
		if base.Ui32(base.I32_wrap_i64(v13)) < base.Ui32(l0) {
			v22 = (v15 - int64(1)) & int64(4294967295)
		} else {
			v22 = v15
		}
		F_PrepareRedoRemoveFull(m, base.I64_extend_i32_u(l0)|v22<<(uint(int64(32))%64), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ProcedureCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int64, l20 int64, l21 int64, l22 int32, l23 int64, l24 int32, l25 int64, l26 int32, l27 float32, l28 float32) {
	mBase := m.M
	_ = mBase
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v133 int32
	_ = v133
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
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
	var v169 int32
	_ = v169
	var v220 int32
	_ = v220
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v416 int32
	_ = v416
	var v419 int64
	_ = v419
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v462 int64
	_ = v462
	var v466 int64
	_ = v466
	var v468 int64
	_ = v468
	var v470 int64
	_ = v470
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v677 int32
	_ = v677
	var v678 int64
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int64
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v731 int32
	_ = v731
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int64
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v869 int32
	_ = v869
	var v887 int32
	_ = v887
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v968 int32
	_ = v968
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1040 int32
	_ = v1040
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1095 int32
	_ = v1095
	var v1126 int32
	_ = v1126
	var v1144 int32
	_ = v1144
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1234 int32
	_ = v1234
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int64
	_ = v1424
	var v1434 int32
	_ = v1434
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1558 int64
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int64
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1621 int32
	_ = v1621
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1656 int32
	_ = v1656
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1685 int32
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1711 int32
	_ = v1711
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1790 int32
	_ = v1790
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1871 int32
	_ = v1871
	var v1876 int32
	_ = v1876
	v44 = m.G0
	v46 = v44 - int32(640)
	m.G0 = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l18)+16))
	if base.Ui32(v48) < base.Ui32(int32(101)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v52 = base.B2i32(l19 == int64(0))
	if v52 == int32(0) {
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
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L20
	} else {
		goto L395
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L20
	} else {
		goto L392
	}
L5:
	;
	v55 = base.I32_wrap_i64(l19)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v56 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v66 = l18
	v67 = v48
	goto L7
L7:
	;
	v69 = base.B2i32(l20 == int64(0))
	if l20 == int64(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v59 <= int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v62 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v63 != int32(26) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v66 = v55
	v67 = v59
	goto L7
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L20
	} else {
		goto L389
	}
L13:
	;
	v84 = int32(0)
	goto L15
L14:
	;
	v71 = base.I32_wrap_i64(l20)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v72 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v86 = l18 + int32(24)
	v87 = F_check_valid_polymorphic_signature(m, l5, v86, v48)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	if v75 != v67 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v77 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	if v78 != int32(18) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v84 = v71 + int32(24)
	goto L15
L20:
	;
	return
L21:
	;
	if v87 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = F_check_valid_internal_signature(m, l5, v86, v48)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L20
	} else {
		goto L384
	}
L25:
	;
	if v91 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v96 = v66 + int32(24)
	v97 = int32(0)
	if v52|base.B2i32(v67 == v97) == v97 {
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
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L20
	} else {
		goto L379
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L20
	} else {
		goto L374
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L20
	} else {
		goto L369
	}
L31:
	;
	v133 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	if v84 == int32(0) {
		v405 = int64(0)
		goto L44
	} else {
		goto L45
	}
L34:
	;
	if v84 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v169 = v133 + int32(1)
	if v169 != v67 {
		v133 = v169
		goto L34
	} else {
		goto L43
	}
L37:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133+v84))))
	v151 = v149 - int32(105)
	if base.B2i32(v151 == int32(0))|base.B2i32(v151 == int32(13)) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v159 = v96 + v133<<(uint(int32(2))%32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v161 = F_check_valid_polymorphic_signature(m, v160, v86, v48)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v161 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v164 = F_check_valid_internal_signature(m, v163, v86, v48)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	if v164 != 0 {
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
	v406 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+630)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v46)+624)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v46)+616)) = v406
	*(*int64)(unsafe.Add(mBase, uint32(v46)+608)) = v406
	v416 = int32(0)
	base.MemoryFill(m, v46+int32(368), v416, int32(240))
	v419 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+358)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v46)+352)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v46)+344)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v46)+336)) = v419
	v428 = v46 + int32(272)
	v430 = F_strncpy(m, v428, l1, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+63)) = uint8(v416)
	goto L80
L45:
	;
	if v67 == int32(0) {
		v405 = int64(0)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v220 = int32(0)
	v254 = v220
	v255 = v220
	goto L47
L47:
	;
	v267 = v254 + v84
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	switch v268 - int32(98) {
	case 0, 7:
		goto L56
	default:
		goto L53
	case 13:
		goto L55
	case 18:
		v357 = v255
		goto L49
	case 20:
		goto L54
	}
L48:
	;
	v405 = base.I64_extend_i32_u(v357)
	goto L44
L49:
	;
	v359 = v254 + int32(1)
	if v359 != v67 {
		v254 = v359
		v255 = v357
		goto L47
	} else {
		goto L79
	}
L50:
	;
	v357 = int32(_a_F_ProcedureCreate_0)
	goto L49
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L20
	} else {
		goto L76
	}
L52:
	;
	v357 = int32(2283)
	goto L49
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L20
	} else {
		goto L73
	}
L54:
	;
	if v255 != 0 {
		goto L51
	} else {
		goto L65
	}
L55:
	;
	if base.B2i32(l12 != int32(112))|base.B2i32(v255 == int32(0)) != 0 {
		v357 = v255
		goto L49
	} else {
		goto L61
	}
L56:
	;
	v271 = int32(0)
	if v255 == v271 {
		v357 = v271
		goto L49
	} else {
		goto L57
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L20
	} else {
		goto L58
	}
L58:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_1), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(281), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	v293 = m.ExcPending
	if v293 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_1), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(285), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
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
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v96+v254<<(uint(int32(2))%32))))
	switch v306 - int32(2276) {
	case 0:
		v357 = v306
		goto L49
	case 1:
		goto L52
	default:
		goto L66
	}
L66:
	;
	if v306 == int32(_a_F_ProcedureCreate_4) {
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v311 = F_get_element_type(m, v306)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	if v311 != 0 {
		v357 = v311
		goto L49
	} else {
		goto L69
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_5), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(307), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
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
	v330 = int32(*(*int8)(unsafe.Add(mBase, uint32(v267))))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v330
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_6), v46+int32(16))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L20
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(312), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
	v350 = m.ExcPending
	if v350 != 0 {
		goto L20
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(292), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v46)+496)) = base.I64_extend_i32_u(v48)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+488)) = base.I64_extend_i32_s(l17)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+480)) = base.I64_extend_i32_s(l16)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+472)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+464)) = base.I64_extend_i32_u(l15)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+456)) = base.I64_extend_i32_u(l14)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+448)) = base.I64_extend_i32_u(l13)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+440)) = base.I64_extend_i32_s(l12)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+432)) = base.I64_extend_i32_u(l26)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+424)) = v405
	*(*int64)(unsafe.Add(mBase, uint32(v46)+416)) = base.I64_extend_i32_s(base.I32_reinterpret_f32(l28))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+408)) = base.I64_extend_i32_s(base.I32_reinterpret_f32(l27))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+400)) = base.I64_extend_i32_u(l7)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+392)) = base.I64_extend_i32_u(l6)
	v462 = base.I64_extend_i32_u(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+384)) = v462
	*(*int64)(unsafe.Add(mBase, uint32(v46)+376)) = base.I64_extend_i32_u(v428)
	if l22 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v466 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l22)+4)))
	v468 = v466
	goto L83
L82:
	;
	v468 = int64(0)
	goto L83
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+504)) = v468
	v470 = base.I64_extend_i32_u(l18)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+520)) = v470
	*(*int64)(unsafe.Add(mBase, uint32(v46)+512)) = base.I64_extend_i32_u(l5)
	if v52 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v69 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+528)) = l19
	goto L84
L86:
	;
	goto L87
L87:
	;
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+628)) = uint8(v477)
	goto L84
L88:
	;
	if l21 != int64(0) {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+536)) = l20
	goto L88
L90:
	;
	goto L91
L91:
	;
	v482 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+629)) = uint8(v482)
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
	*(*int64)(unsafe.Add(mBase, uint32(v46)+544)) = l21
	goto L92
L94:
	;
	goto L95
L95:
	;
	v487 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+630)) = uint8(v487)
	goto L92
L96:
	;
	if l23 != int64(0) {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v489 = F_nodeToString(m, l22)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L20
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v495 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+631)) = uint8(v495)
	goto L96
L100:
	;
	v491 = F_cstring_to_text(m, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L20
	} else {
		goto L101
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+552)) = base.I64_extend_i32_u(v491)
	goto L96
L102:
	;
	v502 = F_cstring_to_text(m, l9)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L20
	} else {
		goto L106
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+560)) = l23
	goto L102
L104:
	;
	goto L105
L105:
	;
	v500 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+632)) = uint8(v500)
	goto L102
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+568)) = base.I64_extend_i32_u(v502)
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
	v506 = F_cstring_to_text(m, l10)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L20
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+634)) = uint8(v510)
	goto L107
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+576)) = base.I64_extend_i32_u(v506)
	goto L107
L112:
	;
	if l25 != int64(0) {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v512 = F_nodeToString(m, l11)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L20
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v518 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+635)) = uint8(v518)
	goto L112
L116:
	;
	v514 = F_cstring_to_text(m, v512)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+584)) = base.I64_extend_i32_u(v514)
	goto L112
L118:
	;
	v527 = F_table_open(m, int32(1255), int32(3))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L20
	} else {
		goto L122
	}
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+592)) = l25
	goto L118
L120:
	;
	goto L121
L121:
	;
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+636)) = uint8(v523)
	goto L118
L122:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v527)+52))
	v532 = F_SearchSysCache3(m, int32(46), base.I64_extend_i32_u(l1), v470, v462)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L20
	} else {
		goto L130
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L20
	} else {
		goto L363
	}
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L20
	} else {
		goto L357
	}
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L20
	} else {
		goto L351
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L20
	} else {
		goto L344
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L20
	} else {
		goto L335
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L20
	} else {
		goto L331
	}
L129:
	;
	v1066 = F_new_object_addresses(m)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L20
	} else {
		goto L243
	}
L130:
	;
	if v532 != 0 {
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
	v998 = F_get_user_default_acl(m, int32(19), l6, l2)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L20
	} else {
		goto L236
	}
L134:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+22)))
	v539 = v537 + v538
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v541 = F_object_ownercheck(m, int32(1255), v540, l6)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L20
	} else {
		goto L135
	}
L135:
	;
	if v541 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	F_aclcheck_error(m, int32(2), int32(19), l1)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L20
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+96)))
	if v549 != l12&int32(255) {
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
	v556 = m.ExcPending
	if v556 != 0 {
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
	v559 = m.ExcPending
	if v559 != 0 {
		goto L20
	} else {
		goto L144
	}
L144:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_7), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L20
	} else {
		goto L145
	}
L145:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+96)))
	switch v565 - int32(97) {
	case 0:
		v571 = int32(_a_F_ProcedureCreate_8)
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
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(423), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L20
	} else {
		goto L152
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+160)) = l1
	v575 = F_errdetail(m, v571, v46+int32(160))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L20
	} else {
		goto L151
	}
L148:
	;
	v571 = int32(_a_F_ProcedureCreate_9)
	goto L147
L149:
	;
	v571 = int32(_a_F_ProcedureCreate_10)
	goto L147
L150:
	;
	v571 = int32(_a_F_ProcedureCreate_11)
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
	v588 = int32(_a_F_ProcedureCreate_12)
	goto L155
L154:
	;
	v588 = int32(_a_F_ProcedureCreate_13)
	goto L155
L155:
	;
	v590 = base.B2i32(l12 == int32(112))
	if l12 == int32(112) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v591 = int32(_a_F_ProcedureCreate_14)
	goto L158
L157:
	;
	v591 = v588
	goto L158
L158:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v539)+108))
	if l5 != v592 {
		goto L127
	} else {
		goto L159
	}
L159:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539)+100)))
	if v594 != l4 {
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
	v677 = v46 + int32(247)
	v678 = F_SysCacheGetAttr(m, int32(46), v532, int32(23), v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L20
	} else {
		goto L182
	}
L162:
	;
	v598 = F_build_function_result_tupdesc_t(m, v532)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L20
	} else {
		goto L163
	}
L163:
	;
	v600 = F_build_function_result_tupdesc_d(m, l12, l19, l20, l21)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	if v598|v600 == int32(0) {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v605 = int32(0)
	if base.B2i32(v598 == v605)|base.B2i32(v600 == v605) != 0 {
		goto L126
	} else {
		goto L166
	}
L166:
	;
	v610 = int32(0)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	if v614 != v615 {
		v666 = v610
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v666 == int32(0) {
		goto L126
	} else {
		goto L181
	}
L168:
	;
	goto L167
L169:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	if v617 != v618 {
		v666 = v610
		goto L168
	} else {
		goto L170
	}
L170:
	;
	if v614 <= int32(0) {
		v666 = int32(1)
		goto L168
	} else {
		goto L171
	}
L171:
	;
	v624 = v614 << (uint(int32(3)) % 32)
	v626 = int32(28)
	v632 = int32(0)
	goto L172
L172:
	;
	v639 = v632 * int32(100)
	v640 = v598 + v624 + v626 + v639
	v641 = int32(4)
	v643 = v639 + (v600 + v624 + v626)
	v646 = F_strcmp(m, v640+v641, v643+v641)
	mBase = m.M
	if v646 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v666 = int32(0)
	goto L168
L174:
	;
	goto L173
L175:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v640)+68))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v643)+68))
	if v647 != v648 {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v640)+76))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v643)+76))
	if v650 != v651 {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v640)+96))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v643)+96))
	if v653 != v654 {
		goto L174
	} else {
		goto L178
	}
L178:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+91)))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643)+91)))
	if v656 != v657 {
		goto L174
	} else {
		goto L179
	}
L179:
	;
	v659 = int32(1)
	v661 = v632 + v659
	if v614 != v661 {
		v632 = v661
		goto L172
	} else {
		goto L180
	}
L180:
	;
	v666 = v659
	goto L168
L181:
	;
	goto L161
L182:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+247)))
	if v680 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v826 = int32(*(*int16)(unsafe.Add(mBase, uint32(v539)+106)))
	if v826 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L184:
	;
	v684 = F_SysCacheGetAttr(m, int32(46), v532, int32(22), v677)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+247)))
	if v686 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v687 = int64(0)
	goto L188
L187:
	;
	v687 = v684
	goto L188
L188:
	;
	v690 = F_get_func_input_arg_names(m, v678, v687, v46+int32(260))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L20
	} else {
		goto L189
	}
L189:
	;
	v694 = F_get_func_input_arg_names(m, l21, l20, v46+int32(248))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L20
	} else {
		goto L190
	}
L190:
	;
	if v690 <= int32(0) {
		goto L183
	} else {
		goto L191
	}
L191:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v46)+248))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v46)+260))
	v731 = int32(0)
	goto L192
L192:
	;
	v745 = v731 << (uint(int32(2)) % 32)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v700+v745)))
	if v747 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L183
L194:
	;
	if v694 <= v731 {
		goto L125
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v781 = v731 + int32(1)
	if v781 != v690 {
		v731 = v781
		goto L192
	} else {
		goto L207
	}
L197:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v699+v745)))
	if v750 == int32(0) {
		goto L125
	} else {
		goto L198
	}
L198:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v747))))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	if base.B2i32(v755 == int32(0))|base.B2i32(v755 != v758) != 0 {
		v776 = v755
		v777 = v758
		goto L200
	} else {
		goto L201
	}
L199:
	;
	if v776-v777 != 0 {
		goto L125
	} else {
		goto L206
	}
L200:
	;
	goto L199
L201:
	;
	v761 = v747
	v762 = v750
	goto L202
L202:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762)+1)))
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761)+1)))
	if v766 == int32(0) {
		v776 = v766
		v777 = v765
		goto L200
	} else {
		goto L204
	}
L203:
	;
	v776 = v766
	v777 = v765
	goto L200
L204:
	;
	v769 = int32(1)
	if v766 == v765 {
		v761 = v761 + v769
		v762 = v762 + v769
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
	v968 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+365)) = uint8(v968)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+339)) = uint8(v968)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+336)) = uint8(v968)
	v981 = F_heap_modify_tuple(m, v532, v529, v46+int32(368), v46+int32(608), v46+int32(336))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
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
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	v831 = v829
	goto L212
L211:
	;
	v831 = int32(0)
	goto L212
L212:
	;
	if v831 < v826 {
		goto L124
	} else {
		goto L213
	}
L213:
	;
	v835 = F_SysCacheGetAttrNotNull(m, int32(46), v532, int32(24))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L20
	} else {
		goto L214
	}
L214:
	;
	v838 = F_text_to_cstring(m, base.I32_wrap_i64(v835))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L20
	} else {
		goto L215
	}
L215:
	;
	v840 = F_stringToNode(m, v838)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
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
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	v844 = v842
	goto L219
L218:
	;
	v844 = int32(0)
	goto L219
L219:
	;
	if v840 == int32(0) {
		goto L208
	} else {
		goto L220
	}
L220:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v847 <= int32(0) {
		goto L208
	} else {
		goto L221
	}
L221:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l22)+12))
	v851 = int32(*(*int16)(unsafe.Add(mBase, uint32(v539)+106)))
	v869 = int32(0)
	v887 = v850 + (v844-v851)<<(uint(int32(2))%32)
	goto L222
L222:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v887)))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v901+v869<<(uint(int32(2))%32))))
	v906 = F_exprType(m, v905)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L20
	} else {
		goto L224
	}
L223:
	;
	goto L208
L224:
	;
	v908 = F_exprType(m, v900)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L20
	} else {
		goto L225
	}
L225:
	;
	if v906 != v908 {
		goto L123
	} else {
		goto L226
	}
L226:
	;
	v912 = v887 + int32(4)
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l22)+12))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l22)+4))
	if base.Ui32(v912) < base.Ui32(v914+v915<<(uint(int32(2))%32)) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v920 = v912
	goto L229
L228:
	;
	v920 = int32(0)
	goto L229
L229:
	;
	v922 = v869 + int32(1)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v840)+4))
	if v922 < v923 {
		v869 = v922
		v887 = v920
		goto L222
	} else {
		goto L230
	}
L230:
	;
	goto L223
L231:
	;
	F_CatalogTupleUpdate(m, v527, v981+int32(4), v981)
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L20
	} else {
		goto L232
	}
L232:
	;
	F_ReleaseCatCache(m, v532)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L20
	} else {
		goto L233
	}
L233:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v981)+16))
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990)+22)))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v990+v991)))
	v995 = F_deleteDependencyRecordsFor(m, int32(1255), v993, int32(1))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L20
	} else {
		goto L234
	}
L234:
	;
	v1026 = v993
	v1032 = v981
	v1040 = v968
	goto L129
L235:
	;
	v1006 = F_GetNewOidWithIndex(m, v527, int32(2690), int32(1))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L20
	} else {
		goto L240
	}
L236:
	;
	if v998 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+600)) = base.I64_extend_i32_u(v998)
	goto L235
L238:
	;
	goto L239
L239:
	;
	v1002 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+637)) = uint8(v1002)
	goto L235
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46)+368)) = base.I64_extend_i32_u(v1006)
	v1014 = F_heap_form_tuple(m, v529, v46+int32(368), v46+int32(608))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	F_CatalogTupleInsert(m, v527, v1014)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+16))
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1018)+22)))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1018+v1019)))
	v1026 = v1021
	v1032 = v1014
	v1040 = v998
	goto L129
L243:
	;
	v1068 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1255)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v46)+260)) = int32(2615)
	v1079 = v46 + int32(260)
	F_add_exact_object_address(m, v1079, v1066)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v46)+260)) = int32(2612)
	F_add_exact_object_address(m, v1079, v1066)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L20
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v46)+260)) = int32(1247)
	F_add_exact_object_address(m, v1079, v1066)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L20
	} else {
		goto L246
	}
L246:
	;
	if v67 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1126 = int32(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v46)+260)) = int32(1247)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v96+v1126<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = v1144
	F_add_exact_object_address(m, v46+int32(260), v1066)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L20
	} else {
		goto L252
	}
L251:
	;
	goto L249
L252:
	;
	v1153 = v1126 + int32(1)
	if v1153 != v67 {
		v1126 = v1153
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
	v1200 = int32(0)
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l24)+4))
	if v1201 <= v1200 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1234 = v1200
	goto L257
L257:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l24)+12))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1247+v1234<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = v1251
	*(*int32)(unsafe.Add(mBase, uint32(v46)+260)) = int32(3576)
	F_add_exact_object_address(m, v46+int32(260), v1066)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L20
	} else {
		goto L259
	}
L258:
	;
	goto L254
L259:
	;
	v1262 = v1234 + int32(1)
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l24)+4))
	if v1262 < v1263 {
		v1234 = v1262
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+268)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+264)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v46)+260)) = int32(1255)
	F_add_exact_object_address(m, v46+int32(260), v1066)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L20
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1317 = int32(0)
	if base.B2i32(l11 == v1317)|base.B2i32(l7 != int32(14)) == v1317 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	goto L263
L265:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[0]))
	F_CheckUsageOnTypesInExpr(m, l11, int32(0), v1326)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L20
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	if l22 != 0 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	F_collectDependenciesOfExpr(m, v1066, l11)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L20
	} else {
		goto L269
	}
L269:
	;
	goto L267
L270:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[0]))
	F_CheckUsageOnTypesInExpr(m, l22, int32(0), v1333)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L20
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[1]))
	goto L275
L273:
	;
	F_collectDependenciesOfExpr(m, v1066, l22)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L20
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	v1346 = v46 + int32(248)
	v1347 = int32(0)
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+8))
	if v1349 <= v1347 {
		v1434 = v1347
		goto L276
	} else {
		goto L277
	}
L276:
	;
	if v1434 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L277:
	;
	v1354 = v1347
	goto L278
L278:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1066)))
	v1398 = v1395 + v1354*int32(12)
	v1399 = F_get_object_namespace(m, v1398)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L20
	} else {
		goto L282
	}
L279:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1346)+8)) = v1422
	v1424 = *(*int64)(unsafe.Add(mBase, uint32(v1398)))
	*(*int64)(unsafe.Add(mBase, uint32(v1346))) = v1424
	v1434 = int32(1)
	goto L276
L280:
	;
	goto L279
L281:
	;
	v1419 = v1354 + int32(1)
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+8))
	if v1419 < v1420 {
		v1354 = v1419
		goto L278
	} else {
		goto L289
	}
L282:
	;
	if v1399 == int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1403 = F_isAnyTempNamespace(m, v1399)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L20
	} else {
		goto L284
	}
L284:
	;
	if v1403 == int32(0) {
		goto L281
	} else {
		goto L285
	}
L285:
	;
	if base.B2i32(v1340 != int32(0))&base.B2i32(l2 == v1340) == int32(0) {
		goto L280
	} else {
		goto L286
	}
L286:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[1]))
	goto L287
L287:
	;
	if base.B2i32(v1411 != int32(0))&base.B2i32(v1399 == v1411) == int32(0) {
		goto L280
	} else {
		goto L288
	}
L288:
	;
	goto L281
L289:
	;
	v1434 = v1347
	goto L276
L290:
	;
	F_record_object_address_dependencies(m, l0, v1066, int32(110))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L20
	} else {
		goto L298
	}
L291:
	;
	v1474 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L20
	} else {
		goto L292
	}
L292:
	;
	if v1474 == int32(0) {
		goto L290
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+48)) = l1
	F_errmsg(m, int32(_a_F_ProcedureCreate_15), v46+int32(48))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L20
	} else {
		goto L294
	}
L294:
	;
	v1487 = F_getObjectDescription(m, v46+int32(248), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L20
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+32)) = v1487
	v1493 = F_errdetail(m, int32(_a_F_ProcedureCreate_16), v46+int32(32))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L20
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(690), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L20
	} else {
		goto L297
	}
L297:
	;
	goto L290
L298:
	;
	F_free_object_addresses(m, v1066)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L20
	} else {
		goto L299
	}
L299:
	;
	if v532 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	F_recordDependencyOnOwner(m, int32(1255), v1026, l6)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L20
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	F_recordDependencyOnCurrentExtension(m, l0, base.B2i32(v532 != int32(0)))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L20
	} else {
		goto L305
	}
L303:
	;
	F_recordDependencyOnNewAcl(m, int32(1255), v1026, l6, v1040)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L20
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	F_pfree(m, v1032)
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L20
	} else {
		goto L306
	}
L306:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[2]))
	if v1520 != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1522 = int32(0)
	F_RunObjectPostCreateHook(m, int32(1255), v1026, v1522, v1522)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L20
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	F_relation_close(m, v527, int32(3))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L20
	} else {
		goto L311
	}
L310:
	;
	goto L309
L311:
	;
	if l8 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	if v532 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L313:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L20
	} else {
		goto L314
	}
L314:
	;
	v1534 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcedureCreate[3])))
	if v1534 != int32(1) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1566 = F_OidFunctionCall1Coll(m, l8, int32(0), base.I64_extend_i32_u(v1026))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L20
	} else {
		goto L326
	}
L316:
	;
	v1537 = base.I32_wrap_i64(l25)
	if v1537 == int32(0) {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1541 = int32(_a_F_ProcedureCreate_17)
	v1543 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[4]))
	v1545 = v1543 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[4])) = v1545
	goto L318
L318:
	;
	v1549 = F_superuser(m)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L20
	} else {
		goto L319
	}
L319:
	;
	if v1549 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1551 = int32(5)
	goto L322
L321:
	;
	v1551 = int32(6)
	goto L322
L322:
	;
	F_ProcessGUCArray(m, v1537, v1551, int32(13), int32(2))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L20
	} else {
		goto L323
	}
L323:
	;
	v1558 = F_OidFunctionCall1Coll(m, l8, int32(0), base.I64_extend_i32_u(v1026))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L20
	} else {
		goto L324
	}
L324:
	;
	F_AtEOXact_GUC(m, int32(1), v1545)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L20
	} else {
		goto L325
	}
L325:
	;
	goto L312
L326:
	;
	goto L312
L327:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, _c_F_ProcedureCreate[5]))
	F_pgstat_create_transactional(m, int32(3), v1574, base.I64_extend_i32_u(v1026))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L20
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	m.G0 = v46 + int32(640)
	return
L330:
	;
	goto L329
L331:
	;
	F_errcode(m, int32(50884740))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L20
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+176)) = l1
	F_errmsg(m, int32(_a_F_ProcedureCreate_18), v46+int32(176))
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L20
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(405), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L20
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L20
	} else {
		goto L336
	}
L336:
	;
	if l12 == int32(112) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1608 = int32(_a_F_ProcedureCreate_19)
	goto L339
L338:
	;
	v1608 = int32(_a_F_ProcedureCreate_20)
	goto L339
L339:
	;
	F_errmsg(m, v1608, int32(0))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L20
	} else {
		goto L340
	}
L340:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v1613 = F_format_procedure(m, v1612)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L20
	} else {
		goto L341
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+148)) = v1613
	*(*int32)(unsafe.Add(mBase, uint32(v46)+144)) = v591
	F_errhint(m, int32(_a_F_ProcedureCreate_21), v46+int32(144))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L20
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(451), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
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
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L20
	} else {
		goto L345
	}
L345:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_20), int32(0))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L20
	} else {
		goto L346
	}
L346:
	;
	v1640 = F_errdetail(m, int32(_a_F_ProcedureCreate_22), int32(0))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L20
	} else {
		goto L347
	}
L347:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v1643 = F_format_procedure(m, v1642)
	mBase = m.M
	v1644 = m.ExcPending
	if v1644 != 0 {
		goto L20
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+132)) = v1643
	*(*int32)(unsafe.Add(mBase, uint32(v46)+128)) = v591
	F_errhint(m, int32(_a_F_ProcedureCreate_21), v46+int32(128))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L20
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(478), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L20
	} else {
		goto L350
	}
L350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L351:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L20
	} else {
		goto L352
	}
L352:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v46)+260))
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v731<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+112)) = v1669
	F_errmsg(m, int32(_a_F_ProcedureCreate_23), v46+int32(112))
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L20
	} else {
		goto L353
	}
L353:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v1677 = F_format_procedure(m, v1676)
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L20
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+100)) = v1677
	*(*int32)(unsafe.Add(mBase, uint32(v46)+96)) = v591
	F_errhint(m, int32(_a_F_ProcedureCreate_21), v46+int32(96))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L20
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(523), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L20
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L20
	} else {
		goto L358
	}
L358:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_24), int32(0))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L20
	} else {
		goto L359
	}
L359:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v1703 = F_format_procedure(m, v1702)
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L20
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+68)) = v1703
	*(*int32)(unsafe.Add(mBase, uint32(v46)+64)) = v591
	F_errhint(m, int32(_a_F_ProcedureCreate_21), v46-int32(-64))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L20
	} else {
		goto L361
	}
L361:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(549), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L20
	} else {
		goto L362
	}
L362:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L363:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L20
	} else {
		goto L364
	}
L364:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_25), int32(0))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L20
	} else {
		goto L365
	}
L365:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	v1729 = F_format_procedure(m, v1728)
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L20
	} else {
		goto L366
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+84)) = v1729
	*(*int32)(unsafe.Add(mBase, uint32(v46)+80)) = v591
	F_errhint(m, int32(_a_F_ProcedureCreate_21), v46+int32(80))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L20
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(573), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L20
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1749 = m.ExcPending
	if v1749 != 0 {
		goto L20
	} else {
		goto L370
	}
L370:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_26), int32(0))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L20
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+192)) = v164
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_27), v46+int32(192))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L20
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(262), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L20
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L20
	} else {
		goto L375
	}
L375:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_28), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L20
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+208)) = v161
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_27), v46+int32(208))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L20
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(254), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L20
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L379:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L20
	} else {
		goto L380
	}
L380:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_26), int32(0))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L20
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+224)) = v91
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_27), v46+int32(224))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L20
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(233), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L20
	} else {
		goto L383
	}
L383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L384:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L20
	} else {
		goto L385
	}
L385:
	;
	F_errmsg(m, int32(_a_F_ProcedureCreate_28), int32(0))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L20
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+240)) = v87
	F_errdetail_internal(m, int32(_a_F_ProcedureCreate_27), v46+int32(240))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L20
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(220), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L20
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_29), int32(0))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L20
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(205), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L20
	} else {
		goto L391
	}
L391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	F_errmsg_internal(m, int32(_a_F_ProcedureCreate_30), int32(0))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L20
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(181), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L20
	} else {
		goto L394
	}
L394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L395:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L20
	} else {
		goto L396
	}
L396:
	;
	v1865 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v1865
	F_errmsg_plural(m, int32(_a_F_ProcedureCreate_31), int32(_a_F_ProcedureCreate_32), v1865, v46)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L20
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(_a_F_ProcedureCreate_2), int32(163), int32(_a_F_ProcedureCreate_3))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L20
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessSyncingRelations(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int64
	_ = v201
	var v205 int32
	_ = v205
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
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
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int64
	_ = v850
	var v852 int64
	_ = v852
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int64
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int64
	_ = v946
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	switch v13 {
	case 0:
		goto L3
	case 1:
		goto L2
	case 2:
		goto L4
	case 3:
		goto L5
	default:
		goto L1
	}
L1:
	;
	return
L2:
	;
	v831 = m.G0
	v833 = v831 - int32(144)
	m.G0 = v833
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v839 = base.AtomicRmwXchg32(m, v836, int32(56), int32(1))
	if v839 != 0 {
		goto L217
	} else {
		goto L218
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L6
	} else {
		goto L214
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L6
	} else {
		goto L211
	}
L5:
	;
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = int32(0)
	F_FetchRelationStates(m, v18, v18, v16+int32(94))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[1]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[2]))
	v28 = int32(0)
	if v25|base.B2i32(v27 == v28) == v28 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[2]))
	if v53 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = int64(68719476740)
	v41 = F_hash_create(m, int32(_a_F_ProcessSyncingRelations_0), int64(256), v16+int32(16), int32(40))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v27|base.B2i32(v25 == int32(0)) != 0 {
		goto L8
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[1])) = v41
	goto L8
L13:
	;
	F_hash_destroy(m, v25)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[1])) = int32(0)
	goto L8
L15:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+94)))
	if v529 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if int32(0) < v56 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = v2
	v66 = v2
	goto L20
L18:
	;
	v508 = v2
	goto L19
L19:
	;
	if v508 == int32(0) {
		goto L15
	} else {
		goto L137
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v66<<(uint(int32(2))%32))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+94)))
	if v74 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v508 = v494
	goto L19
L22:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)))
	if v81 == int32(115) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+94)) = uint8(v79)
	goto L24
L26:
	;
	v501 = v66 + int32(1)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v501 < v502 {
		v63 = v494
		v66 = v501
		goto L20
	} else {
		goto L136
	}
L27:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v73)+8))
	if base.Ui64(l0) < base.Ui64(v84) {
		v494 = v63
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	v129 = F_LWLockAcquire(m, v125+int32(_a_F_ProcessSyncingRelations_1), int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L39
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = l0
	v87 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v87)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	F_LockSharedObject(m, int32(_a_F_ProcessSyncingRelations_2), v92, int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	if v63 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v100 = F_table_open(m, int32(_a_F_ProcessSyncingRelations_3), int32(3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	v102 = v63
	goto L34
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v108 = v16 + int32(16)
	F_ReplicationOriginNameForLogicalRep(m, v105, v106, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	v102 = v100
	goto L34
L36:
	;
	F_replorigin_drop_by_name(m, v108, int32(1), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+16)))
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v73)+8))
	F_UpdateSubscriptionRelState(m, v117, v118, v119, v120, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v494 = v102
	goto L26
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+32))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v136 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[4]))
	if v143 <= v136 {
		v186 = v136
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v186 != 0 {
		goto L53
	} else {
		goto L54
	}
L41:
	;
	goto L40
L42:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[5]))
	v155 = v136
	goto L43
L43:
	;
	v161 = v147 + int32(16) + v155<<(uint(int32(7))%32)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+16)))
	if v162 != int32(1) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v186 = int32(0)
	goto L41
L45:
	;
	v179 = v155 + int32(1)
	if v179 != v143 {
		v155 = v179
		goto L43
	} else {
		goto L52
	}
L46:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v165 == int32(4) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v161)+32))
	if v168 != v134 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161)+36))
	if base.B2i32(v170 != v135)|base.B2i32(int32(1) != v165) != 0 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v186 = v161
	goto L41
L52:
	;
	goto L44
L53:
	;
	v191 = int32(56)
	v192 = v186 + v191
	v195 = base.AtomicRmwXchg32(m, v186, v191, int32(1))
	if v195 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+32))
	v373 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[4]))
	if v381 <= v373 {
		v460 = v373
		goto L113
	} else {
		goto L114
	}
L56:
	;
	F_s_lock(m, v192, int32(_a_F_ProcessSyncingRelations_4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v199)
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v186)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = v201
	if v199 == int32(119) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v205 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+40)) = uint8(v205)
	if base.Ui64(l0) < base.Ui64(v201) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v210 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v192))), uint32(v210))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)))
	if v213 == int32(119) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v208 = v201
	goto L65
L64:
	;
	v208 = l0
	goto L65
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v186)+48)) = v208
	goto L62
L66:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	if v216 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	F_LWLockRelease(m, v365+int32(_a_F_ProcessSyncingRelations_1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L6
	} else {
		goto L111
	}
L69:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
	F_SetLatch(m, v217+int32(316))
	mBase = m.M
	goto L72
L70:
	;
	goto L71
L71:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	F_LWLockRelease(m, v222+int32(_a_F_ProcessSyncingRelations_1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L6
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+94)))
	if v227 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	if v63 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v239 = v63
	goto L76
L76:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L83
	}
L77:
	;
	F_relation_close(m, v63, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v235 = int32(0)
	v237 = F_pgstat_report_stat(m, v235)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v239 = v235
	goto L76
L83:
	;
	v242 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+94)) = uint8(v242)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	goto L84
L84:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[6]))
	if v256 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+32))
	v266 = F_GetSubscriptionRelState(m, v263, v244, v16+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	if base.B2i32(v266 == int32(0))|base.B2i32(v266&int32(255) == int32(115)) != 0 {
		v494 = v239
		goto L26
	} else {
		goto L92
	}
L92:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	v280 = F_LWLockAcquire(m, v276+int32(_a_F_ProcessSyncingRelations_1), int32(1))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+32))
	v286 = int32(0)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[4]))
	if v293 <= v286 {
		v336 = v286
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	F_LWLockRelease(m, v342+int32(_a_F_ProcessSyncingRelations_1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L107
	}
L95:
	;
	goto L94
L96:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[5]))
	v305 = v286
	goto L97
L97:
	;
	v311 = v297 + int32(16) + v305<<(uint(int32(7))%32)
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+16)))
	if v312 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v336 = int32(0)
	goto L95
L99:
	;
	v329 = v305 + int32(1)
	if v329 != v293 {
		v305 = v329
		goto L97
	} else {
		goto L106
	}
L100:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v315 == int32(4) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v311)+32))
	if v318 != v285 {
		goto L99
	} else {
		goto L102
	}
L102:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v311)+36))
	if base.B2i32(v320 != v244)|base.B2i32(int32(1) != v315) != 0 {
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v336 = v311
	goto L95
L106:
	;
	goto L98
L107:
	;
	if v336 == int32(0) {
		v494 = v239
		goto L26
	} else {
		goto L108
	}
L108:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[7]))
	v354 = F_WaitLatch(m, v350, int32(41), int32(1000), int32(134217760))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[7]))
	v358 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v358
	v363 = base.AtomicRmwOr32(m, v358, int32(_a_F_ProcessSyncingRelations_5), v358)
	goto L110
L110:
	;
	goto L84
L111:
	;
	v494 = v63
	goto L26
L112:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	F_LWLockRelease(m, v467+int32(_a_F_ProcessSyncingRelations_1))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L6
	} else {
		goto L130
	}
L113:
	;
	goto L112
L114:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[5]))
	v387 = v385 + int32(16)
	if v381 != int32(1) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v396 = v373
	v397 = v373
	v401 = v373
	goto L118
L116:
	;
	v438 = v373
	v439 = v373
	goto L117
L117:
	;
	v446 = v387 + v439<<(uint(int32(7))%32)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+32))
	if v372 != v447 {
		v460 = v438
		goto L113
	} else {
		goto L128
	}
L118:
	;
	v404 = v387 + v397<<(uint(int32(7))%32)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+32))
	if v405 != v372 {
		v416 = v396
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v381&int32(1) == int32(0) {
		v460 = v428
		goto L113
	} else {
		goto L127
	}
L120:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v404)+160))
	if v417 != v372 {
		v428 = v416
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+16)))
	if v407 != int32(1) {
		v416 = v396
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	v416 = v396 + base.B2i32(base.Ui32(v410-int32(1)) < base.Ui32(int32(2)))
	goto L120
L123:
	;
	v429 = int32(2)
	v430 = v397 + v429
	v432 = v401 + v429
	if v432 != v381&int32(2147483646) {
		v396 = v428
		v397 = v430
		v401 = v432
		goto L118
	} else {
		goto L126
	}
L124:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+144)))
	if v419 != int32(1) {
		v428 = v416
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v404)+128))
	v428 = v416 + base.B2i32(base.Ui32(v422-int32(1)) < base.Ui32(int32(2)))
	goto L123
L126:
	;
	goto L119
L127:
	;
	v438 = v428
	v439 = v430
	goto L117
L128:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+16)))
	if v449 != int32(1) {
		v460 = v438
		goto L113
	} else {
		goto L129
	}
L129:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	v460 = v438 + base.B2i32(base.Ui32(v452-int32(1)) < base.Ui32(int32(2)))
	goto L113
L130:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[1]))
	v477 = F_hash_search(m, v473, v73, int32(1), v16+int32(16))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v479 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v477)+8)) = int64(0)
	goto L134
L133:
	;
	goto L134
L134:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	F_launch_sync_worker(m, int32(1), v460, v485, v477+int32(8))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	v494 = v63
	goto L26
L136:
	;
	goto L21
L137:
	;
	F_relation_close(m, v508, int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L15
L139:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[8]))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533)+36)))
	if v534 != int32(112) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L141
L141:
	;
	m.G0 = v16 + int32(96)
	v596 = m.G0
	v598 = v596 - int32(16)
	m.G0 = v598
	F_FetchRelationStates(m, int32(0), v598+int32(15), v598+int32(14))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L6
	} else {
		goto L165
	}
L142:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L6
	} else {
		goto L163
	}
L143:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L6
	} else {
		goto L144
	}
L144:
	;
	F_FetchRelationStates(m, v16+int32(95), int32(0), v16+int32(16))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v546 == int32(1) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L6
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+95)))
	if v554 != int32(1) {
		goto L142
	} else {
		goto L151
	}
L149:
	;
	v552 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[2]))
	if v558 != 0 {
		goto L142
	} else {
		goto L152
	}
L152:
	;
	v561 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	if v561 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[8]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v565
	F_errmsg(m, int32(_a_F_ProcessSyncingRelations_6), v16)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L6
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L6
	} else {
		goto L159
	}
L157:
	;
	F_errfinish(m, int32(_a_F_ProcessSyncingRelations_7), int32(601), int32(_a_F_ProcessSyncingRelations_8))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v578 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[8]))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	F_ApplyLauncherForgetWorkerStartTime(m, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
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
	v591 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L6
	} else {
		goto L164
	}
L164:
	;
	goto L141
L165:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+14)))
	if v607 == int32(1) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598)+15)))
	if v615 != int32(1) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v613 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L6
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	m.G0 = v598 + int32(16)
	return
L172:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	v623 = F_LWLockAcquire(m, v619+int32(_a_F_ProcessSyncingRelations_1), int32(1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L6
	} else {
		goto L173
	}
L173:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+32))
	v629 = int32(0)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[4]))
	if v637 <= v629 {
		v680 = v629
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v680 != 0 {
		goto L187
	} else {
		goto L188
	}
L175:
	;
	goto L174
L176:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[5]))
	v649 = v629
	goto L177
L177:
	;
	v655 = v641 + int32(16) + v649<<(uint(int32(7))%32)
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+16)))
	if v656 != int32(1) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v680 = int32(0)
	goto L175
L179:
	;
	v673 = v649 + int32(1)
	if v673 != v637 {
		v649 = v673
		goto L177
	} else {
		goto L186
	}
L180:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v659 == int32(4) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v655)+32))
	if v662 != v628 {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v655)+36))
	if base.B2i32(v664 != v629)|base.B2i32(int32(2) != v659) != 0 {
		goto L179
	} else {
		goto L183
	}
L183:
	;
	goto L184
L184:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v655)+20))
	if v670 != 0 {
		v680 = v655
		goto L175
	} else {
		goto L185
	}
L185:
	;
	goto L179
L186:
	;
	goto L178
L187:
	;
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	F_LWLockRelease(m, v686+int32(_a_F_ProcessSyncingRelations_1))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L6
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v692)+32))
	v694 = int32(0)
	v702 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[4]))
	if v702 <= v694 {
		v781 = v694
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L171
L191:
	;
	v788 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[3]))
	F_LWLockRelease(m, v788+int32(_a_F_ProcessSyncingRelations_1))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L209
	}
L192:
	;
	goto L191
L193:
	;
	v706 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[5]))
	v708 = v706 + int32(16)
	if v702 != int32(1) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v717 = v694
	v718 = v694
	v722 = v694
	goto L197
L195:
	;
	v759 = v694
	v760 = v694
	goto L196
L196:
	;
	v767 = v708 + v760<<(uint(int32(7))%32)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+32))
	if v693 != v768 {
		v781 = v759
		goto L192
	} else {
		goto L207
	}
L197:
	;
	v725 = v708 + v718<<(uint(int32(7))%32)
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)+32))
	if v726 != v693 {
		v737 = v717
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v702&int32(1) == int32(0) {
		v781 = v749
		goto L192
	} else {
		goto L206
	}
L199:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v725)+160))
	if v738 != v693 {
		v749 = v737
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725)+16)))
	if v728 != int32(1) {
		v737 = v717
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	v737 = v717 + base.B2i32(base.Ui32(v731-int32(1)) < base.Ui32(int32(2)))
	goto L199
L202:
	;
	v750 = int32(2)
	v751 = v718 + v750
	v753 = v722 + v750
	if v753 != v702&int32(2147483646) {
		v717 = v749
		v718 = v751
		v722 = v753
		goto L197
	} else {
		goto L205
	}
L203:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725)+144)))
	if v740 != int32(1) {
		v749 = v737
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v725)+128))
	v749 = v737 + base.B2i32(base.Ui32(v743-int32(1)) < base.Ui32(int32(2)))
	goto L202
L205:
	;
	goto L198
L206:
	;
	v759 = v749
	v760 = v751
	goto L196
L207:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767)+16)))
	if v770 != int32(1) {
		v781 = v759
		goto L192
	} else {
		goto L208
	}
L208:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v781 = v759 + base.B2i32(base.Ui32(v773-int32(1)) < base.Ui32(int32(2)))
	goto L192
L209:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	F_launch_sync_worker(m, int32(2), v781, int32(0), v796+int32(120))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	goto L171
L211:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessSyncingRelations_9), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L6
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_ProcessSyncingRelations_10), int32(180), int32(_a_F_ProcessSyncingRelations_11))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L6
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessSyncingRelations_12), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L6
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_ProcessSyncingRelations_10), int32(185), int32(_a_F_ProcessSyncingRelations_11))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L6
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_s_lock(m, v836+int32(56), int32(_a_F_ProcessSyncingRelations_4))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L6
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846)+40)))
	if v847 != int32(99) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	goto L219
L221:
	;
	v960 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v846)+56)), uint32(v960))
	m.G0 = v833 + int32(144)
	goto L1
L222:
	;
	v850 = *(*int64)(unsafe.Add(mBase, uint32(v846)+48))
	if base.Ui64(l0) < base.Ui64(v850) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v852 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v833)+120)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+112)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+104)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+96)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+88)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+80)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+72)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+64)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+56)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+48)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+40)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+32)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+24)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+16)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833)+8)) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v833))) = v852
	*(*int64)(unsafe.Add(mBase, uint32(v846)+48)) = l0
	v885 = int32(115)
	*(*uint8)(unsafe.Add(mBase, uint32(v846)+40)) = uint8(v885)
	v887 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v846)+56)), uint32(v887))
	v891 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[9]))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+20))
	goto L224
L224:
	;
	if base.B2i32(v892 == int32(2)) == int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L6
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+32))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v900)+36))
	v903 = int32(*(*int8)(unsafe.Add(mBase, uint32(v900)+40)))
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v900)+48))
	F_UpdateSubscriptionRelState(m, v901, v902, v903, v904, int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L6
	} else {
		goto L229
	}
L228:
	;
	goto L227
L229:
	;
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[10]))
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[11]))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+36))
	m.T0[v914].(func(*base.Module, int32, int32))(m, v909, v833+int32(140))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v918)+32))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v918)+36))
	v922 = v833 - int32(-64)
	F_ReplicationSlotNameForTablesync(m, v919, v920, v922)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L6
	} else {
		goto L231
	}
L231:
	;
	v926 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[10]))
	F_ReplicationSlotDropAtPubNode(m, v926, v922, int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L6
	} else {
		goto L232
	}
L232:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L6
	} else {
		goto L233
	}
L233:
	;
	v933 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L6
	} else {
		goto L234
	}
L234:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L6
	} else {
		goto L235
	}
L235:
	;
	v938 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[0]))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v938)+32))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v938)+36))
	F_ReplicationOriginNameForLogicalRep(m, v939, v940, v833)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L6
	} else {
		goto L236
	}
L236:
	;
	F_replorigin_session_reset(m)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L6
	} else {
		goto L237
	}
L237:
	;
	v946 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[12])) = v946
	*(*int64)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[13])) = v946
	v952 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_ProcessSyncingRelations[14])) = uint16(v952)
	goto L238
L238:
	;
	F_replorigin_drop_by_name(m, v833, int32(1), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L6
	} else {
		goto L239
	}
L239:
	;
	F_FinishSyncWorker(m)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = F_pg_database_locale(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+v10<<(uint(v11)%32))))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 < v11 {
			v25 = F_pg_database_locale(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_pg_iswalpha(m, v14, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v27
					return v29
				}
			}
		} else {
			v18 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			if v19 != v18 {
				v25 = F_pg_database_locale(m)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = F_pg_iswalpha(m, v14, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v27
						return v29
					}
				}
			} else {
				if base.Ui32(int32(127)) < base.Ui32(v14) {
					v29 = v18
					return v29
				} else {
					v25 = F_pg_database_locale(m)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = F_pg_iswalpha(m, v14, v25)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = v27
							return v29
						}
					}
				}
			}
		}
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	v2 = int32(0)
	v6 = F_palloc0(m, int32(40))
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12 + v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v17 - v19
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v22 + v24<<(uint(int32(2))%32)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v30 = F_palloc(m, int32(32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(0)
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)) = uint8(v43)
	F_check_stack_depth(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v47 = F_TParserGet(m, v6)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v81 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if v47 == int32(0) {
		v80 = v2
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	if v51 != int32(6) {
		v80 = v2
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55 + v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v60 + v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = v65 + v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v70 + v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v76
	v80 = int32(1)
	goto L8
L12:
	;
	v82 = v81
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_pfree(m, v6)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L19
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	F_pfree(m, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v86
	if v86 != 0 {
		v82 = v86
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	return v80
}
func F_p_isignore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	return v2
}
func F_p_isspace(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
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
	v4 = F_pg_database_locale(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+v10<<(uint(v11)%32))))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 < v11 {
			v24 = F_pg_database_locale(m)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
				if v26 == int32(0) {
					if base.Ui32(int32(127)) < base.Ui32(v14) {
						v38 = int32(0)
					} else {
						v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_p_isspace[0]))))
						v38 = base.B2i32(v32 < int32(0))
					}
					v41 = v38
					return v41
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
					v36 = m.T0[v35].(func(*base.Module, int32, int32) int32)(m, v14, v24)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = v36
						v41 = v38
						return v41
					}
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			if v18 != int32(1) {
				v24 = F_pg_database_locale(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					if v26 == int32(0) {
						if base.Ui32(int32(127)) < base.Ui32(v14) {
							v38 = int32(0)
						} else {
							v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_p_isspace[0]))))
							v38 = base.B2i32(v32 < int32(0))
						}
						v41 = v38
						return v41
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
						v36 = m.T0[v35].(func(*base.Module, int32, int32) int32)(m, v14, v24)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = v36
							v41 = v38
							return v41
						}
					}
				}
			} else {
				if base.Ui32(int32(127)) < base.Ui32(v14) {
					v41 = int32(0)
					return v41
				} else {
					v24 = F_pg_database_locale(m)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						if v26 == int32(0) {
							if base.Ui32(int32(127)) < base.Ui32(v14) {
								v38 = int32(0)
							} else {
								v32 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_p_isspace[0]))))
								v38 = base.B2i32(v32 < int32(0))
							}
							v41 = v38
							return v41
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
							v36 = m.T0[v35].(func(*base.Module, int32, int32) int32)(m, v14, v24)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = v36
								v41 = v38
								return v41
							}
						}
					}
				}
			}
		}
	}
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
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v2 == int32(1) {
		v5 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v5)
		v9 = int32(1)
	} else {
		v9 = int32(0)
	}
	return v9
}
func F_pairingheap_initialize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(444)
	return
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= v3 {
		v48 = v3
		m.G0 = v9 + int32(16)
		return v48
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		if v15 < v11 {
			v48 = v3
			m.G0 = v9 + int32(16)
			return v48
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			if v17 != 0 {
				v19 = m.T0[v17].(func(*base.Module, int32, int32, int32, int32) int32)(m, v14, v11, int32(0), v9)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v28 = v19
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
					if v29 == int32(0) {
						v48 = v3
						m.G0 = v9 + int32(16)
						return v48
					} else {
						v33 = F_palloc0(m, int32(28))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v11
							*(*int64)(unsafe.Add(mBase, uint32(v33))) = int64(8)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v38
							v42 = F_get_typcollation(m, v38)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v42
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v45
								v48 = v33
								m.G0 = v9 + int32(16)
								return v48
							}
						}
					}
				}
			} else {
				v28 = v14 + v11<<(uint(int32(4))%32) + int32(16)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
				if v29 == int32(0) {
					v48 = v3
					m.G0 = v9 + int32(16)
					return v48
				} else {
					v33 = F_palloc0(m, int32(28))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v11
						*(*int64)(unsafe.Add(mBase, uint32(v33))) = int64(8)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v38
						v42 = F_get_typcollation(m, v38)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v42
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v45
							v48 = v33
							m.G0 = v9 + int32(16)
							return v48
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
	var v171 int32
	_ = v171
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(524)
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
	v171 = F_errdetail(m, int32(_a_F_parserOpenTable_1), v10)
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
	F_errfinish(m, int32(_a_F_parserOpenTable_3), int32(1493), int32(_a_F_parserOpenTable_4))
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
	F_errfinish(m, int32(_a_F_parserOpenTable_3), int32(1472), int32(_a_F_parserOpenTable_4))
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
	F_errfinish(m, int32(_a_F_parserOpenTable_3), int32(1488), int32(_a_F_parserOpenTable_4))
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
func F_pgstatginindex_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstatginindex_internal(m, v2, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_pgstatindexbyid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstatindexbyid_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstatindexbyid_1), int32(194), int32(_a_F_pgstatindexbyid_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
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
			v28 = F_relation_open(m, base.I32_wrap_i64(v3), int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = F_pgstatindex_impl(m, v28, l0)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					return v30
				}
			}
		}
	}
}
func F_pgstattuple(m *base.Module, l0 int32) int64 {
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
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_superuser(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			if v8 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pgstattuple_0), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pgstattuple_1), int32(178), int32(_a_F_pgstattuple_2))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
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
				v28 = F_textToQualifiedNameList(m, v4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = F_makeRangeVarFromNameList(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v33 = F_relation_openrv(m, v30, int32(1))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = F_pgstat_relation(m, v33, l0)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
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
func F_pgstattuplebyid_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_pgstat_relation(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
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
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
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
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v667 int32
	_ = v667
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
	return v667
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v659
	v667 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
	v652 = F_psprintf(m, int32(_a_F_plain_crypt_verify_0), v10+int32(32))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L30
	} else {
		goto L156
	}
L4:
	;
	v539 = F_strlen(m, l0)
	mBase = m.M
	v544 = F_pg_md5_encrypt(m, l2, l0, v539, v10+int32(48), v10+int32(44))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L30
	} else {
		goto L132
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
	if v512 != 0 {
		v667 = v134
		goto L1
	} else {
		goto L130
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L30
	} else {
		goto L127
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
		v512 = v5
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
	F_errfinish(m, int32(_a_F_plain_crypt_verify_3), int32(545), int32(_a_F_plain_crypt_verify_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v512 = v5
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
		v512 = v373
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
	F_errfinish(m, int32(_a_F_plain_crypt_verify_3), int32(556), int32(_a_F_plain_crypt_verify_4))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L30
	} else {
		goto L98
	}
L98:
	;
	v512 = v373
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
	v419 = int32(0)
	if v396 == v419 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L109
L111:
	;
	v512 = base.B2i32(v507 == int32(0))
	goto L35
L112:
	;
	v507 = int32(0)
	goto L111
L113:
	;
	goto L114
L114:
	;
	v427 = v396 & int32(3)
	if base.Ui32(v396) < base.Ui32(int32(4)) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v507 = base.B2i32(v493 != int32(0))
	goto L111
L116:
	;
	v473 = v466
	v474 = v467
	v475 = v468
	v479 = v419
	goto L124
L117:
	;
	v466 = v416
	v467 = v418
	v468 = int32(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v434 = v416
	v435 = v418
	v436 = int32(0)
	v439 = v419
	goto L120
L120:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434))))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+1)))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+1)))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+2)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+2)))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+3)))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+3)))
	v456 = v436 | (v441 ^ v442) | (v445 ^ v446) | (v449 ^ v450) | (v453 ^ v454)
	v457 = int32(4)
	v458 = v435 + v457
	v460 = v434 + v457
	v462 = v439 + v457
	if v462 != v396&int32(-4) {
		v434 = v460
		v435 = v458
		v436 = v456
		v439 = v462
		goto L120
	} else {
		goto L122
	}
L121:
	;
	if v427 == int32(0) {
		v493 = v456
		goto L115
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v466 = v460
	v467 = v458
	v468 = v456
	goto L116
L124:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v483 = v475 | (v480 ^ v481)
	v484 = int32(1)
	v489 = v479 + v484
	if v489 != v427 {
		v473 = v473 + v484
		v474 = v474 + v484
		v475 = v483
		v479 = v489
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v493 = v483
	goto L115
L126:
	;
	goto L125
L127:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = v521
	F_errmsg_internal(m, int32(_a_F_plain_crypt_verify_6), v137+int32(16))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L30
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_plain_crypt_verify_3), int32(572), int32(_a_F_plain_crypt_verify_4))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L30
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
	v537 = F_psprintf(m, int32(_a_F_plain_crypt_verify_7), v10+int32(16))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L30
	} else {
		goto L131
	}
L131:
	;
	v659 = v537
	goto L2
L132:
	;
	if v544 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v659 = v548
	goto L2
L134:
	;
	goto L135
L135:
	;
	v550 = v10 + int32(48)
	v551 = F_strlen(m, v550)
	mBase = m.M
	v552 = F_strlen(m, l1)
	mBase = m.M
	if v551 != v552 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v646 = F_psprintf(m, int32(_a_F_plain_crypt_verify_7), v10)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L30
	} else {
		goto L155
	}
L137:
	;
	v554 = int32(0)
	if v551 == v554 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v642 != 0 {
		goto L136
	} else {
		goto L154
	}
L139:
	;
	v642 = int32(0)
	goto L138
L140:
	;
	goto L141
L141:
	;
	v562 = v551 & int32(3)
	if base.Ui32(v551) < base.Ui32(int32(4)) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v642 = base.B2i32(v628 != int32(0))
	goto L138
L143:
	;
	v608 = v601
	v609 = v602
	v610 = v603
	v614 = v554
	goto L151
L144:
	;
	v601 = v550
	v602 = l1
	v603 = int32(0)
	goto L143
L145:
	;
	goto L146
L146:
	;
	v569 = v550
	v570 = l1
	v571 = int32(0)
	v574 = v554
	goto L147
L147:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+1)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+1)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+2)))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+2)))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+3)))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569)+3)))
	v591 = v571 | (v576 ^ v577) | (v580 ^ v581) | (v584 ^ v585) | (v588 ^ v589)
	v592 = int32(4)
	v593 = v570 + v592
	v595 = v569 + v592
	v597 = v574 + v592
	if v597 != v551&int32(-4) {
		v569 = v595
		v570 = v593
		v571 = v591
		v574 = v597
		goto L147
	} else {
		goto L149
	}
L148:
	;
	if v562 == int32(0) {
		v628 = v591
		goto L142
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	v601 = v595
	v602 = v593
	v603 = v591
	goto L143
L151:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v608))))
	v618 = v610 | (v615 ^ v616)
	v619 = int32(1)
	v624 = v614 + v619
	if v624 != v562 {
		v608 = v608 + v619
		v609 = v609 + v619
		v610 = v618
		v614 = v624
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v628 = v618
	goto L142
L153:
	;
	goto L152
L154:
	;
	v667 = int32(0)
	goto L1
L155:
	;
	v659 = v646
	goto L2
L156:
	;
	v659 = v652
	goto L2
}
func F_plainto_tsquery(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14391(m, l0, int32(1275))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
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
	v22 = l0 - int32(-64)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12+v22)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	if l3 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+72))
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+72))
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	F_errfinish(m, int32(_a_F_plan_recursive_revoke_2), int32(2507), int32(_a_F_plan_recursive_revoke_3))
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
func F_positionsel(m *base.Module, l0 int32) int64 {
	return int64(4591870180066957722)
}
func F_postgresql_fdw_validator(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
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
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	F_errfinish(m, int32(_a_F_postgresql_fdw_validator_0), int32(708), int32(_a_F_postgresql_fdw_validator_1))
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
	return int64(1)
L4:
	;
	return int64(0)
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
								v99 = int32(3795)
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
									v99 = int32(3808)
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
										v99 = int32(3795)
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
											v99 = int32(3808)
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
						F_errfinish(m, int32(_a_F_pre_sync_fname_1), int32(3785), int32(_a_F_pre_sync_fname_2))
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
												v99 = int32(3795)
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
													v99 = int32(3808)
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
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
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
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
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v383 int64
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
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
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v586 int32
	_ = v586
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v650 int32
	_ = v650
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(432)
	m.G0 = v26
	if l0 == v3 {
		v650 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v26 + int32(432)
	return v650
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v31 == int32(9) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v606
	v650 = int32(0)
	goto L1
L4:
	;
	v275 = F_palloc0(m, int32(20))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L66
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L63
	}
L6:
	;
	v35 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v36 = F_SearchSysCache1(m, int32(0), v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v235 = F_expression_tree_walker_impl(m, l0, int32(898), l1)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L62
	}
L9:
	;
	return int32(0)
L10:
	;
	if v36 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v53 = v26 + int32(16)
	v54 = F_get_aggregate_argtypes(m, l0, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = F_resolve_aggregate_transtype(m, v56, v51, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v57
	v60 = int32(-1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v61 == int32(0) {
		v74 = v60
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+42)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_get_typlenbyval(m, v76, v26+int32(422), v26+int32(421))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v67 = F_exprType(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v67 != v57 {
		v74 = v60
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v71 = F_exprTypmod(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v74 = v71
	goto L14
L19:
	;
	v87 = F_SysCacheGetAttr(m, int32(0), v36, int32(21), v26+int32(420))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+420)))
	if v89 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_getTypeInputInfo(m, v57, v26+int32(428), v26+int32(424))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	v109 = int64(0)
	goto L23
L23:
	;
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L28
	}
L24:
	;
	v99 = F_text_to_cstring(m, base.I32_wrap_i64(v87))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v26)+428))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v26)+424))
	v104 = F_OidInputFunctionCall(m, v101, v99, v102, int32(-1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v99)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v109 = v104
	goto L23
L28:
	;
	v112 = F_contain_volatile_functions(m, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v112 != 0 {
		v259 = v3
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
	if v114 == int32(0) {
		v259 = v3
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		v259 = v3
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v125 = int32(0)
	v127 = int32(-1)
	v130 = v3
	goto L33
L33:
	;
	v146 = v127 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v125<<(uint(int32(2))%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	if v147 != v156 {
		v213 = v130
		goto L36
	} else {
		goto L37
	}
L34:
	;
	F_list_free(m, v130)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L59
	}
L35:
	;
	goto L34
L36:
	;
	v215 = v125 + int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v215 < v216 {
		v125 = v215
		v127 = v146
		v130 = v213
		goto L33
	} else {
		goto L58
	}
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	if v158 != v159 {
		v213 = v130
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+48)))
	if v161 != v162 {
		v213 = v130
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+49)))
	if v164 != v165 {
		v213 = v130
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+50)))
	if v167 != v168 {
		v213 = v130
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	v172 = F_equal(m, v170, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	if v172 == int32(0) {
		v213 = v130
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v155)+36))
	v178 = F_equal(m, v176, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	if v178 == int32(0) {
		v213 = v130
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v155)+40))
	v184 = F_equal(m, v182, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	if v184 == int32(0) {
		v213 = v130
		goto L36
	} else {
		goto L47
	}
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v155)+44))
	v190 = F_equal(m, v188, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if v190 == int32(0) {
		v213 = v130
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v194 != v195 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+12)))
	if v207 != int32(1) {
		v213 = v130
		goto L36
	} else {
		goto L56
	}
L51:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v155)+8))
	if v197 != v198 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	if v200 != v201 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	v205 = F_equal(m, v203, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	if v205 != 0 {
		goto L35
	} else {
		goto L55
	}
L55:
	;
	goto L50
L56:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v211 = F_lappend_int(m, v130, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v213 = v211
	goto L36
L58:
	;
	v259 = v213
	goto L4
L59:
	;
	if v146 == int32(-1) {
		v259 = int32(0)
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224+v146<<(uint(int32(2))%32))))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v230 = F_lappend(m, v229, l0)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v230
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	v606 = v146
	v610 = v233
	goto L3
L62:
	;
	v650 = v235
	goto L1
L63:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v241
	F_errmsg_internal(m, int32(_a_F_preprocess_aggrefs_walker_0), v26)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_preprocess_aggrefs_walker_1), int32(153), int32(_a_F_preprocess_aggrefs_walker_2))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = int32(331)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+428)) = l0
	v285 = F_list_make1_impl(m, int32(1), v26+int32(12))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v275)+12)) = uint8(base.B2i32(v75 != int32(119)))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+4)) = v285
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+344))
	if v291 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v294 = v292
	goto L70
L69:
	;
	v294 = int32(0)
	goto L70
L70:
	;
	v295 = F_lappend(m, v291, v275)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+344)) = v295
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v298 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	F_get_typlenbyval(m, v57, v26+int32(424), v26+int32(419))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L9
	} else {
		goto L77
	}
L73:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v301 == int32(0) {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v304 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+356)) = uint8(v304)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+352))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+352)) = v306 + v304
	goto L72
L76:
	;
	goto L75
L77:
	;
	if base.B2i32(v259 == int32(0))|base.B2i32(v75 == int32(119)) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+8)) = v586
	v606 = v294
	v610 = v586
	goto L3
L79:
	;
	v418 = F_palloc0(m, int32(64))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L9
	} else {
		goto L100
	}
L80:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v321 <= int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v324 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+424)))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+420)))
	v327 = int32(1)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+419)))
	v335 = int32(0)
	v344 = v321
	goto L82
L82:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+348))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v358 = int32(2)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v335<<(uint(v358)%32))))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v356+v361<<(uint(v358)%32))))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	if v50 != v366 {
		v387 = v344
		goto L85
	} else {
		goto L86
	}
L83:
	;
	if v361 != int32(-1) {
		v586 = v361
		goto L78
	} else {
		goto L99
	}
L84:
	;
	goto L83
L85:
	;
	v390 = v335 + int32(1)
	if v390 < v387 {
		v335 = v390
		v344 = v387
		goto L82
	} else {
		goto L98
	}
L86:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365)+28))
	if v57 != v368 {
		v387 = v344
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v365)+16))
	if v47 != v370 {
		v387 = v344
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v365)+20))
	if v46 != v372 {
		v387 = v344
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
	if v48 != v374 {
		v387 = v344
		goto L85
	} else {
		goto L90
	}
L90:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+56)))
	if v326&v327 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v376&int32(1) == int32(0) {
		v387 = v344
		goto L85
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v376&int32(1) != 0 {
		v387 = v344
		goto L85
	} else {
		goto L95
	}
L94:
	;
	goto L84
L95:
	;
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v365)+48))
	v384 = F_datumIsEqual(m, v109, v383, v329&v327, v324)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	if v384 != 0 {
		goto L84
	} else {
		goto L97
	}
L97:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	v387 = v386
	goto L85
L98:
	;
	goto L79
L99:
	;
	goto L79
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = int32(332)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+24)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v418)+8)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v418)+32)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v418)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v418)+20)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v418)+16)) = v47
	v432 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+424)))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+36)) = v432
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+419)))
	*(*int64)(unsafe.Add(mBase, uint32(v418)+48)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v418)+44)) = v45
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+40)) = uint8(v434)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+420)))
	*(*uint8)(unsafe.Add(mBase, uint32(v418)+56)) = uint8(v438)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+348))
	if v440 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v443 = v441
	goto L103
L102:
	;
	v443 = int32(0)
	goto L103
L103:
	;
	v444 = F_lappend(m, v440, v418)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+348)) = v444
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+356)))
	if v447 != 0 {
		v586 = v443
		goto L78
	} else {
		goto L105
	}
L105:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v418)+24))
	if v448 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v451 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+356)) = uint8(v451)
	v586 = v443
	goto L78
L107:
	;
	goto L108
L108:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v418)+28))
	if v453 != int32(2281) {
		v586 = v443
		goto L78
	} else {
		goto L109
	}
L109:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v418)+16))
	if v456 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v461 != int32(_a_F_preprocess_aggrefs_walker_3) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v418)+20))
	if v457 != 0 {
		v461 = v456
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v458 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+357)) = uint8(v458)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v418)+16))
	v461 = v460
	goto L110
L114:
	;
	goto L113
L115:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v418)+20))
	if v464 != int32(_a_F_preprocess_aggrefs_walker_4) {
		v586 = v443
		goto L78
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v467 = int32(0)
	v468 = m.G0
	v470 = v468 - int32(16)
	m.G0 = v470
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v472 == v467 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	goto L117
L119:
	;
	if v539 != 0 {
		v586 = v443
		goto L78
	} else {
		goto L144
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L9
	} else {
		goto L141
	}
L121:
	;
	m.G0 = v470 + int32(16)
	goto L119
L122:
	;
	v539 = int32(1)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	if v477 <= int32(0) {
		v539 = int32(1)
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v482 = v467
	goto L126
L126:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v472)+12))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v503+v482<<(uint(int32(2))%32))))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v509 = F_exprType(m, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L9
	} else {
		goto L128
	}
L127:
	;
	v539 = v512
	goto L121
L128:
	;
	v511 = int32(2249)
	v512 = base.B2i32(v509 != v511)
	if v509 == v511 {
		v539 = v512
		goto L121
	} else {
		goto L129
	}
L129:
	;
	v517 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v509))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	if v517 == int32(0) {
		goto L120
	} else {
		goto L131
	}
L131:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517)+16))
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+22)))
	v523 = v521 + v522
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+78)))
	if v524 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_ReleaseCatCache(m, v517)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L9
	} else {
		goto L139
	}
L133:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523)+112))
	if v525 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v523)+108))
	if v526 != 0 {
		goto L132
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	F_ReleaseCatCache(m, v517)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L9
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v539 = int32(0)
	goto L121
L139:
	;
	v533 = v482 + int32(1)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	if v533 < v534 {
		v482 = v533
		goto L126
	} else {
		goto L140
	}
L140:
	;
	goto L127
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = v509
	F_errmsg_internal(m, int32(_a_F_preprocess_aggrefs_walker_5), v470)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_preprocess_aggrefs_walker_6), int32(2201), int32(_a_F_preprocess_aggrefs_walker_7))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	v575 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+357)) = uint8(v575)
	v586 = v443
	goto L78
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
	v45 = l2 + v35<<(uint(int32(3))%32) + v38*int32(100)
	F_pq_sendstring(m, v9, v45+int32(32))
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
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
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+100)))
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
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
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
								F_errfinish(m, int32(_a_F_processTypesSpec_1), int32(1148), int32(_a_F_processTypesSpec_2))
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
							F_errfinish(m, int32(_a_F_processTypesSpec_1), int32(1148), int32(_a_F_processTypesSpec_2))
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
func F_process_concurrent_changes(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
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
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
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
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
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
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int64
	_ = v456
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
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
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int64
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	v3 = l2
	v18 = m.G0
	v20 = v18 - int32(1120)
	m.G0 = v20
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[0]))
	if v26 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[1]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	v71 = int32(76)
	v72 = v70 + v71
	v75 = base.AtomicRmwXchg32(m, v70, v71, int32(1))
	if v75 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_concurrent_changes[2])))
	if v30&int32(1) == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v35 = int32(_a_F_process_concurrent_changes_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v37 + v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v41 + v38
	v45 = int32(0)
	v47 = int32(_a_F_process_concurrent_changes_1)
	v48 = base.AtomicRmwOr32(m, v45, v47, v45)
	*(*int64)(unsafe.Add(mBase, uint32(v26+int32(8))+232)) = int64(5)
	v56 = base.AtomicRmwOr32(m, v45, v47, v45)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v57 + v38
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v63 - v38
	goto L2
L5:
	;
	F_s_lock(m, v72, int32(_a_F_process_concurrent_changes_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+16)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v70)+8)) = l0
	v81 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v70)+76)), uint32(v81))
	v85 = v70 + int32(100)
	F_ConditionVariablePrepareToSleep(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	goto L11
L11:
	;
	v107 = base.AtomicRmwXchg32(m, v72, int32(0), int32(1))
	if v107 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L21
	}
L13:
	;
	F_s_lock(m, v72, int32(_a_F_process_concurrent_changes_2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v112 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v70)+76)), uint32(v112))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v115 != v111 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	F_ConditionVariableSleep(m, v85, int32(134217776))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L12
L20:
	;
	goto L11
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v70)+96))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v122
	v127 = v20 - int32(-64)
	v132 = F_pg_snprintf(m, v127, int32(1024), int32(_a_F_process_concurrent_changes_3), v20+int32(48))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v136 = int32(0)
	v138 = F_BufFileOpenFileSet(m, v70+int32(20), v127, v136, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1091)) = uint8(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+52))
	v145 = F_MakeSingleTupleTableSlot(m, v143, int32(_a_F_process_concurrent_changes_4))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+52))
	v148 = F_table_slot_callbacks(m, v142)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v150 = F_MakeSingleTupleTableSlot(m, v147, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v142)+52))
	v154 = F_MakeSingleTupleTableSlot(m, v152, int32(_a_F_process_concurrent_changes_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+152))
	if v157 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v160 = F_MakePerTupleExprContext(m, v156)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	v162 = v157
	goto L30
L30:
	;
	v163 = int32(_a_F_process_concurrent_changes_5)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[4]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[4])) = v166
	v169 = v150 + int32(32)
	v170 = int32(0)
	v174 = v170
	v176 = v170
	goto L36
L31:
	;
	v162 = v160
	goto L30
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L8
	} else {
		goto L138
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L8
	} else {
		goto L135
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L8
	} else {
		goto L131
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L8
	} else {
		goto L128
	}
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[5]))
	if v190 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	F_ExecDropSingleTupleTableSlot(m, v145)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L8
	} else {
		goto L124
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v195 = int32(1)
	v197 = F_BufFileReadCommon(m, v138, v20+int32(1091), v195, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v197 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1091)))
	switch v199 - int32(85) {
	case 0, 15:
		goto L47
	default:
		goto L46
	case 20:
		goto L48
	case 32:
		goto L49
	}
L44:
	;
	goto L45
L45:
	;
	goto L37
L46:
	;
	F_restore_tuple(m, v138, v142, v145)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L54
	}
L47:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L52
	}
L48:
	;
	v209 = v176&int32(255) - int32(85)
	if base.B2i32(v209 == int32(0))|base.B2i32(v209 == int32(15)) != 0 {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	F_restore_tuple(m, v138, v142, v154)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1091)))
	v174 = int32(1)
	v176 = v205
	goto L36
L51:
	;
	goto L46
L52:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	goto L46
L54:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1091)))
	switch v223 - int32(85) {
	case 0:
		goto L58
	default:
		goto L57
	case 15:
		goto L59
	case 20:
		goto L56
	}
L55:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)+152))
	if v577 != 0 {
		goto L120
	} else {
		goto L121
	}
L56:
	;
	v494 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L8
	} else {
		goto L113
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L8
	} else {
		goto L110
	}
L58:
	;
	if v174&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L59:
	;
	v226 = F_find_target_tuple(m, v142, l1, v145, v150)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	if v226 == int32(0) {
		goto L35
	} else {
		goto L61
	}
L61:
	;
	v231 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v234 = int32(0)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v142)+188))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+96))
	v241 = m.T0[v240].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v142, v169, v231, int32(2), v234, v234, v234, v20+int32(1096))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	if v241 != 0 {
		goto L34
	} else {
		goto L64
	}
L64:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[0]))
	if v247 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v561 = v174
	goto L55
L66:
	;
	goto L65
L67:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_concurrent_changes[2])))
	if v251&int32(1) == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v256 = int32(_a_F_process_concurrent_changes_0)
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	v259 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v258 + v259
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v262 + v259
	v266 = int32(0)
	v268 = int32(_a_F_process_concurrent_changes_1)
	v269 = base.AtomicRmwOr32(m, v266, v268, v266)
	v274 = v247 + int32(280)
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v274)))
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = v275 + int64(1)
	v281 = base.AtomicRmwOr32(m, v266, v268, v266)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v282 + v259
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v288 - v259
	goto L66
L69:
	;
	v295 = v154
	goto L71
L70:
	;
	v295 = v145
	goto L71
L71:
	;
	v296 = F_find_target_tuple(m, v142, l1, v295, v150)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	if v296 == int32(0) {
		goto L33
	} else {
		goto L73
	}
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v145)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	if int32(0) < v301 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v311 = int32(0)
	goto L77
L75:
	;
	goto L76
L76:
	;
	v397 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L8
	} else {
		goto L98
	}
L77:
	;
	v325 = v311 << (uint(int32(3)) % 32)
	v326 = v300 + int32(28) + v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+6)))
	if v327&int32(4) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L76
L79:
	;
	v376 = v311 + int32(1)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	if v376 < v377 {
		v311 = v376
		goto L77
	} else {
		goto L97
	}
L80:
	;
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v326)+2)))
	if v330 != int32(_a_F_process_concurrent_changes_6) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v334 = v311 + int32(1)
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145)+6)))
	if v335 <= v311 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	m.T0[v338].(func(*base.Module, int32, int32))(m, v145, v334)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v145)+20))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341+v311))))
	if v343 != 0 {
		goto L79
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v344 = int32(*(*int16)(unsafe.Add(mBase, uint32(v145)+6)))
	if v344 <= v311 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+16))
	m.T0[v347].(func(*base.Module, int32, int32))(m, v145, v334)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L8
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v350+v325)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v353 != int32(1) {
		goto L79
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v356 != int32(18) {
		goto L79
	} else {
		goto L92
	}
L92:
	;
	v359 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+6)))
	if v359 <= v311 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	m.T0[v362].(func(*base.Module, int32, int32))(m, v150, v334)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L8
	} else {
		goto L96
	}
L94:
	;
	v366 = v350
	goto L95
L95:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v368+v325)))
	*(*int64)(unsafe.Add(mBase, uint32(v366+v325))) = v370
	goto L79
L96:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	v366 = v365
	goto L95
L97:
	;
	goto L78
L98:
	;
	v400 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v142)+188))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)+100))
	v411 = m.T0[v410].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v142, v169, v145, v397, int32(1), v400, v400, v400, v20+int32(1096), v20+int32(1116), v20+int32(1092))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	if v411 != 0 {
		goto L32
	} else {
		goto L100
	}
L100:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v20)+1092))
	switch v414 {
	case 0:
		goto L101
	default:
		v416 = int32(1)
		goto L102
	case 2:
		goto L103
	}
L101:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[0]))
	if v428 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v419 = int32(0)
	v421 = F_ExecInsertIndexTuples(m, v417, v418, v416, v145, v419, v419)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L104
	}
L103:
	;
	v416 = int32(5)
	goto L102
L104:
	;
	goto L101
L105:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+12))
	m.T0[v475].(func(*base.Module, int32))(m, v154)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L109
	}
L106:
	;
	goto L105
L107:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_concurrent_changes[2])))
	if v432&int32(1) == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v437 = int32(_a_F_process_concurrent_changes_0)
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	v440 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v439 + v440
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v443 + v440
	v447 = int32(0)
	v449 = int32(_a_F_process_concurrent_changes_1)
	v450 = base.AtomicRmwOr32(m, v447, v449, v447)
	v455 = v428 + int32(272)
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v455)))
	*(*int64)(unsafe.Add(mBase, uint32(v455))) = v456 + int64(1)
	v462 = base.AtomicRmwOr32(m, v447, v449, v447)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	*(*int32)(unsafe.Add(mBase, uint32(v428))) = v463 + v440
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v469 - v440
	goto L106
L109:
	;
	v561 = int32(0)
	goto L55
L110:
	;
	v483 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+1091)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v483
	F_errmsg_internal(m, int32(_a_F_process_concurrent_changes_7), v20)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_process_concurrent_changes_8), int32(2789), int32(_a_F_process_concurrent_changes_9))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L8
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
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v142)+188))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+80))
	m.T0[v499].(func(*base.Module, int32, int32, int32, int32, int32))(m, v142, v145, v494, int32(8), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v504 = int32(0)
	v507 = F_ExecInsertIndexTuples(m, v502, v503, v504, v145, v504, v504)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	v513 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[0]))
	if v513 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v561 = v174
	goto L55
L117:
	;
	goto L116
L118:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_concurrent_changes[2])))
	if v517&int32(1) == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v522 = int32(_a_F_process_concurrent_changes_0)
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	v525 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v524 + v525
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = v528 + v525
	v532 = int32(0)
	v534 = int32(_a_F_process_concurrent_changes_1)
	v535 = base.AtomicRmwOr32(m, v532, v534, v532)
	v540 = v513 + int32(264)
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v540)))
	*(*int64)(unsafe.Add(mBase, uint32(v540))) = v541 + int64(1)
	v547 = base.AtomicRmwOr32(m, v532, v534, v532)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = v548 + v525
	v554 = *(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[3])) = v554 - v525
	goto L117
L120:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+20))
	F_MemoryContextReset(m, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L8
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1091)))
	v174 = v561
	v176 = v581
	goto L36
L123:
	;
	goto L122
L124:
	;
	F_ExecDropSingleTupleTableSlot(m, v150)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	F_ExecDropSingleTupleTableSlot(m, v154)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_concurrent_changes[4])) = v164
	F_BufFileClose(m, v138)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L8
	} else {
		goto L127
	}
L127:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v592 + int32(1)
	m.G0 = v20 + int32(1120)
	return
L128:
	;
	F_errmsg_internal(m, int32(_a_F_process_concurrent_changes_10), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_process_concurrent_changes_8), int32(2756), int32(_a_F_process_concurrent_changes_9))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L8
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L8
	} else {
		goto L132
	}
L132:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v142)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(_a_F_process_concurrent_changes_11)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v619 + int32(4)
	F_errmsg(m, int32(_a_F_process_concurrent_changes_12), v20+int32(16))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_process_concurrent_changes_8), int32(2889), int32(_a_F_process_concurrent_changes_13))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errmsg_internal(m, int32(_a_F_process_concurrent_changes_10), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L8
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_process_concurrent_changes_8), int32(2772), int32(_a_F_process_concurrent_changes_9))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L8
	} else {
		goto L139
	}
L139:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v142)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(_a_F_process_concurrent_changes_14)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v655 + int32(4)
	F_errmsg(m, int32(_a_F_process_concurrent_changes_12), v20+int32(32))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_process_concurrent_changes_8), int32(2851), int32(_a_F_process_concurrent_changes_15))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L8
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
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
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
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
	return v251
L6:
	;
	v35 = F_pull_varnos(m, l0, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v34 = v16
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
		v34 = int32(0)
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
	v34 = v20
	goto L6
L13:
	;
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+32)))
	if v27 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v34 = v20
	goto L6
L16:
	;
	goto L17
L17:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
	if v29 != int64(0) {
		v251 = int32(0)
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v34 = v20
	goto L6
L19:
	;
	if v35 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v39 = F_bms_copy(m, l5)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v145 = v35
	goto L22
L22:
	;
	v154 = int32(0)
	v160 = F_make_restrictinfo(m, l0, v34, int32(1), v154, v154, base.B2i32(v35 == v154), l6, v145, v154, v154)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L48
	}
L23:
	;
	v142 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+335)) = uint8(v142)
	v145 = v134
	goto L22
L24:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v42 = int32(0)
	if base.B2i32(v39 == v42)|base.B2i32(v41 == v42) != 0 {
		v88 = base.B2i32(v39|v41 == v42)
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v88 != 0 {
		v134 = v39
		goto L23
	} else {
		goto L36
	}
L26:
	;
	goto L25
L27:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v56 != v57 {
		v88 = int32(0)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v59 = int32(1)
	if v56 <= v59 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v62 = v59
	goto L31
L30:
	;
	v62 = v56
	goto L31
L31:
	;
	v63 = int32(8)
	v68 = int32(0)
	goto L32
L32:
	;
	v76 = v68 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v39+v63+v76)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41+v63+v76)))
	v81 = base.B2i32(v78 == v80)
	if v78 != v80 {
		v88 = v81
		goto L26
	} else {
		goto L34
	}
L33:
	;
	v88 = v81
	goto L26
L34:
	;
	v84 = v68 + int32(1)
	if v84 != v62 {
		v68 = v84
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v93 == int32(0) {
		v134 = v39
		goto L23
	} else {
		goto L37
	}
L37:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96 <= int32(0) {
		v134 = v39
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v101 = v39
	v102 = int32(0)
	goto L39
L39:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v102<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
	if v114 != int32(1) {
		v128 = v101
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v134 = v128
	goto L23
L41:
	;
	v130 = v102 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v130 < v131 {
		v101 = v128
		v102 = v130
		goto L39
	} else {
		goto L47
	}
L42:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	v118 = F_bms_is_member(m, v117, v101)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v118 == int32(0) {
		v128 = v101
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113)+24))
	v123 = F_bms_del_member(m, v101, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v126 = F_bms_del_members(m, v123, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v128 = v126
	goto L41
L47:
	;
	goto L40
L48:
	;
	v162 = int32(0)
	if v145 == v162 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v207 == int32(2) {
		goto L65
	} else {
		goto L66
	}
L50:
	;
	v207 = int32(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v170 = int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	if v171 <= v170 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v174 = v170
	goto L55
L54:
	;
	v174 = v171
	goto L55
L55:
	;
	v178 = int32(0)
	v180 = v162
	goto L56
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v145+int32(8)+v178<<(uint(int32(2))%32))))
	if v187 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v207 = v199
	goto L49
L58:
	;
	goto L57
L59:
	;
	v188 = int32(2)
	if v180 != 0 {
		v199 = v188
		goto L58
	} else {
		goto L62
	}
L60:
	;
	v194 = v180
	goto L61
L61:
	;
	v196 = v178 + int32(1)
	if v196 != v174 {
		v178 = v196
		v180 = v194
		goto L56
	} else {
		goto L64
	}
L62:
	;
	v189 = int32(1)
	if base.Ui32(v189) < base.Ui32(base.I32_popcnt(v187)) {
		v199 = v188
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v194 = v189
	goto L61
L64:
	;
	v199 = v194
	goto L58
L65:
	;
	v211 = F_pull_var_clause(m, v34, int32(26))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+10)))
	if v218 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	F_add_vars_to_targetlist(m, l0, v211, v145)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_list_free(m, v211)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	F_distribute_restrictinfo_to_rels(m, l0, v160)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L83
	}
L72:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v219 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v222 != int32(17) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v219)+28))
	if v225 == int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v228 != int32(2) {
		goto L71
	} else {
		goto L76
	}
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v234 = F_exprType(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v236 = F_op_mergejoinable(m, v231, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v236 == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v240 = F_contain_volatile_functions(m, v160)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v240 != 0 {
		goto L71
	} else {
		goto L81
	}
L81:
	;
	v242 = F_get_mergejoin_opfamilies(m, v231)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+96)) = v242
	goto L71
L83:
	;
	v251 = v160
	goto L5
}
func F_prsd_end(m *base.Module, l0 int32) int64 {
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
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
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
	return int64(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v9
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
	F_pfree(m, v4)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	return int64(0)
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
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v42<<(uint(int32(3))%32)+v34*int32(100))+18)))
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
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(709), int32(_a_F_pub_collist_validate_2))
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
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(715), int32(_a_F_pub_collist_validate_2))
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
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(721), int32(_a_F_pub_collist_validate_2))
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
	F_errfinish(m, int32(_a_F_pub_collist_validate_1), int32(727), int32(_a_F_pub_collist_validate_2))
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
	var v89 int32
	_ = v89
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
	var v136 int32
	_ = v136
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
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
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
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l1 == v7 {
		v401 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v401
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v24 - int32(21) {
	case 0:
		goto L3
	case 1:
		goto L4
	default:
		v401 = l1
		goto L1
	}
L3:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v271 {
	case 0:
		goto L68
	default:
		v401 = l1
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
		v401 = l1
		goto L1
	case 2:
		goto L6
	}
L5:
	;
	v223 = int32(0)
	v225 = F_convert_EXISTS_sublink_to_join(m, l0, l1, v223, l3)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
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
		v161 = v7
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v32 + int32(16)
	if v161 != 0 {
		v401 = v161
		goto L1
	} else {
		goto L41
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v37 == int32(0) {
		v161 = v7
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 != int32(2) {
		v161 = v7
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
		v161 = v7
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
		v161 = v7
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
		v161 = v7
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	if v49 != 0 {
		v161 = v7
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v50 == int32(0) {
		v161 = v7
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 != int32(1) {
		v161 = v7
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
		v161 = v7
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v61 == int32(0) {
		v161 = v7
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 < int32(2) {
		v161 = v7
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
		v161 = v7
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	if v76 == int32(0) {
		v136 = v7
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
		v136 = v7
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v89 = v7
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
	v136 = v124
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
	v161 = int32(0)
	goto L7
L35:
	;
	goto L36
L36:
	;
	v124 = F_lappend(m, v89, v118)
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
		v89 = v124
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
	v153 = F_make_SAOP_expr(m, v68, v71, v147, v151, v67, v136, int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v161 = v153
	goto L7
L41:
	;
	v175 = int32(0)
	v177 = F_convert_ANY_sublink_to_join(m, l0, l1, v175, l3)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	if v177 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+12)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v177
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v185 = F_pull_up_sublinks_jointree_recurse(m, l0, v182, v20+int32(12))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v177)+16)) = v185
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v177)+28))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v194 = F_pull_up_sublinks_qual_recurse(m, l0, v188, v177+int32(12), l3, v177+int32(16), v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+28)) = v194
	v401 = v175
	goto L1
L48:
	;
	v401 = l1
	goto L1
L49:
	;
	goto L50
L50:
	;
	v200 = F_convert_ANY_sublink_to_join(m, l0, l1, int32(0), l5)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	if v200 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v200
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	v210 = F_pull_up_sublinks_jointree_recurse(m, l0, v207, v20+int32(12))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v210
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v200)+28))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v219 = F_pull_up_sublinks_qual_recurse(m, l0, v213, v200+int32(12), l5, v200+int32(16), v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+28)) = v219
	v401 = int32(0)
	goto L1
L55:
	;
	if v225 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v225
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v233 = F_pull_up_sublinks_jointree_recurse(m, l0, v230, v20+int32(12))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v233
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v225)+28))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v242 = F_pull_up_sublinks_qual_recurse(m, l0, v236, v225+int32(12), l3, v225+int32(16), v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+28)) = v242
	v401 = v223
	goto L1
L61:
	;
	v401 = l1
	goto L1
L62:
	;
	goto L63
L63:
	;
	v248 = F_convert_EXISTS_sublink_to_join(m, l0, l1, int32(0), l5)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	if v248 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v248)+12)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v248
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	v258 = F_pull_up_sublinks_jointree_recurse(m, l0, v255, v20+int32(12))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+16)) = v258
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v248)+28))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v267 = F_pull_up_sublinks_qual_recurse(m, l0, v261, v248+int32(12), l5, v248+int32(16), v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+28)) = v267
	v401 = int32(0)
	goto L1
L68:
	;
	v328 = int32(0)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v329 == v328 {
		v401 = v328
		goto L1
	} else {
		goto L92
	}
L69:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v274 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	if v277 != int32(22) {
		v401 = l1
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	switch v280 {
	case 0:
		goto L74
	default:
		v401 = l1
		goto L1
	case 2:
		goto L75
	}
L72:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+16))
	v315 = F_pull_up_sublinks_jointree_recurse(m, l0, v312, v20+int32(8))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L23
	} else {
		goto L90
	}
L73:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+12)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v307
	v311 = v307
	goto L72
L74:
	;
	v295 = F_convert_EXISTS_sublink_to_join(m, l0, v274, int32(1), l3)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L23
	} else {
		goto L83
	}
L75:
	;
	v282 = F_convert_ANY_sublink_to_join(m, l0, v274, int32(1), l3)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L23
	} else {
		goto L76
	}
L76:
	;
	if v282 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+12)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v282
	v311 = v282
	goto L72
L78:
	;
	goto L79
L79:
	;
	if l5 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v290 = F_convert_ANY_sublink_to_join(m, l0, v274, int32(1), l5)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L23
	} else {
		goto L81
	}
L81:
	;
	if v290 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v307 = v290
	goto L73
L83:
	;
	if v295 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v295
	v311 = v295
	goto L72
L85:
	;
	goto L86
L86:
	;
	if l5 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v303 = F_convert_EXISTS_sublink_to_join(m, l0, v274, int32(1), l5)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L23
	} else {
		goto L88
	}
L88:
	;
	if v303 == int32(0) {
		v401 = l1
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v307 = v303
	goto L73
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+16)) = v315
	v318 = int32(0)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v311)+28))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v325 = F_pull_up_sublinks_qual_recurse(m, l0, v319, v311+int32(16), v322, v318, v318)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L23
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311)+28)) = v325
	v401 = v318
	goto L1
L92:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if int32(0) < v332 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v337 = int32(0)
	v342 = v7
	goto L96
L94:
	;
	v373 = v7
	goto L95
L95:
	;
	if v373 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L96:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v353+v337<<(uint(int32(2))%32))))
	v358 = F_pull_up_sublinks_qual_recurse(m, l0, v357, l2, l3, l4, l5)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L23
	} else {
		goto L98
	}
L97:
	;
	v373 = v362
	goto L95
L98:
	;
	if v358 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v360 = F_lappend(m, v342, v358)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L23
	} else {
		goto L102
	}
L100:
	;
	v362 = v342
	goto L101
L101:
	;
	v364 = v337 + int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v364 < v365 {
		v337 = v364
		v342 = v362
		goto L96
	} else {
		goto L103
	}
L102:
	;
	v362 = v360
	goto L101
L103:
	;
	goto L97
L104:
	;
	v401 = int32(0)
	goto L1
L105:
	;
	goto L106
L106:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v387 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v401 = v391
	goto L1
L108:
	;
	goto L109
L109:
	;
	v392 = F_make_andclause(m, v373)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L23
	} else {
		goto L110
	}
L110:
	;
	v401 = v392
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
