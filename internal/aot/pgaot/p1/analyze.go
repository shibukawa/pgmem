package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_parse_analyze_fixedparams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	v8 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	if int32(0) < l3 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = F_palloc(m, int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = l4
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 != int32(141) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+108)) = int32(527)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v16
	goto L5
L7:
	;
	v67 = F_transformStmt(m, v8, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L8:
	;
	v65 = v25
	goto L7
L9:
	;
	goto L10
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v25
	goto L14
L12:
	;
	v41 = v25
	goto L13
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v44 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	if v37 != 0 {
		v33 = v36
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v41 = v36
	goto L13
L16:
	;
	goto L15
L17:
	;
	v65 = v25
	goto L7
L18:
	;
	goto L19
L19:
	;
	v48 = F_palloc0(m, int32(20))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(242)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = int32(0)
	v65 = v48
	goto L7
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+156)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+160)) = v71
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[0]))
	switch v75 {
	case 0:
		v82 = v73
		goto L22
	case 1:
		goto L23
	default:
		goto L24
	}
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[1]))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v80 = F_JumbleQuery(m, v67)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[5])))
	if v77 != int32(1) {
		v82 = v73
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v82 = v80
	goto L22
L27:
	;
	m.T0[v84].(func(*base.Module, int32, int32, int32))(m, v8, v67, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_free_parsestate(m, v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v67)+16))
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[2]))
	if v93 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	return v67
L33:
	;
	goto L32
L34:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[3])))
	if v97&int32(1) == int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v93)+392))
	if int32(1)&base.B2i32(v104 != int64(0)) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v108 = int32(_a_F_parse_analyze_fixedparams_0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[4]))
	v111 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[4])) = v110 + v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v114 + v111
	v118 = int32(0)
	v120 = int32(_a_F_parse_analyze_fixedparams_1)
	v121 = base.AtomicRmwOr32(m, v118, v120, v118)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+392)) = v89
	v126 = base.AtomicRmwOr32(m, v118, v120, v118)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v127 + v111
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[4])) = v133 - v111
	goto L33
}
func F_serializeAnalyzeReceive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int64
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v279 int64
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int64
	_ = v294
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v333 int64
	_ = v333
	var v335 int64
	_ = v335
	var v336 int64
	_ = v336
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v361 int64
	_ = v361
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v382 int64
	_ = v382
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v389 int64
	_ = v389
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v396 int64
	_ = v396
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v403 int64
	_ = v403
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v410 int64
	_ = v410
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v420 int64
	_ = v420
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+9)))
	if v20 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F___clock_gettime(m, int32(1), v15+int32(8))
	mBase = m.M
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v30 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+16)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v33 = v27*int64(-1000000000) - v30
	v34 = v32
	goto L3
L2:
	;
	v33 = int64(0)
	v34 = v19
	goto L3
L3:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+7)))
	if v35 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	base.MemoryCopy(m, v15+int32(8), int32(_a_F_serializeAnalyzeReceive_0), int32(128))
	goto L6
L5:
	;
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v17 == v43 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v143 < v142 {
		goto L33
	} else {
		goto L34
	}
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v45 == v18 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	F_pfree(m, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v17
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
	if v18 <= v54 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L14
L17:
	;
	v61 = F_palloc0(m, v18*int32(28))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v61
	v66 = v54
	goto L19
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v84 = v17 + v76<<(uint(int32(3))%32) + v66*int32(100) + int32(28)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	switch v86 {
	case 0:
		goto L22
	case 1:
		goto L24
	default:
		goto L23
	}
L20:
	;
	goto L7
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	F_fmgr_info(m, v120, v85+v66*int32(28))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L15
	} else {
		goto L31
	}
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v84)+68))
	v113 = v15 + int32(144)
	F_getTypeOutputInfo(m, v111, v113, v15+int32(143))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L15
	} else {
		goto L30
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+68))
	v89 = v15 + int32(144)
	F_getTypeBinaryOutputInfo(m, v87, v89, v15+int32(143))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v119 = v89
	goto L21
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+24)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v101
	F_errmsg(m, int32(_a_F_serializeAnalyzeReceive_1), v15)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_serializeAnalyzeReceive_2), int32(95), int32(_a_F_serializeAnalyzeReceive_3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v119 = v113
	goto L21
L31:
	;
	v127 = v66 + int32(1)
	if v127 != v18 {
		v66 = v127
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	m.T0[v146].(func(*base.Module, int32, int32))(m, l0, v142)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v149 = int32(_a_F_serializeAnalyzeReceive_4)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[0]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[0])) = v152
	v155 = l1 + int32(44)
	F_resetStringInfo(m, v155)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v155)+12)) = int32(68)
	goto L37
L36:
	;
	goto L35
L37:
	;
	F_enlargeStringInfo(m, v155, int32(2))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v165 = int32(8)
	v171 = v18<<(uint(v165)%32) | int32(base.Ui32(v18&int32(_a_F_serializeAnalyzeReceive_5))>>(uint(v165)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v162+v163))) = uint16(v171)
	v174 = v162 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v174
	if int32(0) < v18 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v181 = int32(0)
	goto L42
L40:
	;
	v269 = v174
	goto L41
L41:
	;
	v279 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+64)) = v279 + base.I64_extend_i32_s(v269)
	*(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[0])) = v150
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_MemoryContextReset(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L15
	} else {
		goto L58
	}
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v181))))
	if v193 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v269 = v266
	goto L41
L44:
	;
	v264 = v181 + int32(1)
	if v264 != v18 {
		v181 = v264
		goto L42
	} else {
		goto L57
	}
L45:
	;
	F_enlargeStringInfo(m, v155, int32(4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L15
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v210 = v207 + v181*int32(28)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v211+v181<<(uint(int32(3))%32))))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v216 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v199+v200))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v199 + int32(4)
	goto L44
L49:
	;
	v219 = F_OutputFunctionCall(m, v210, v215)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L15
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v224 = F_SendFunctionCall(m, v210, v215)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L15
	} else {
		goto L54
	}
L52:
	;
	v221 = F_strlen(m, v219)
	mBase = m.M
	F_pq_sendcountedtext(m, v155, v219, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	F_enlargeStringInfo(m, v155, int32(4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v233 = int32(2)
	v235 = int32(4)
	v236 = int32(base.Ui32(v226)>>(uint(v233)%32)) - v235
	v237 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v230+v231))) = base.I32_rotr(v236&v237, int32(8)) | base.I32_rotr(v236, int32(24))&v237
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v230 + v235
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	F_appendBinaryStringInfo(m, v155, v224+v235, int32(base.Ui32(v252)>>(uint(v233)%32))-v235)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	goto L44
L57:
	;
	goto L43
L58:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+9)))
	if v289 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F___clock_gettime(m, int32(1), v15+int32(144))
	mBase = m.M
	v294 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	v295 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+152)))
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v15)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+72)) = v294 + (v295 + (v296*int64(1000000000) + v33))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v304 = v303
	goto L61
L60:
	;
	v304 = v288
	goto L61
L61:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+7)))
	if v305 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v309 = l1 + int32(80)
	v311 = v15 + int32(8)
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v309)))
	v314 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[1]))
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v311)))
	*(*int64)(unsafe.Add(mBase, uint32(v309))) = v312 + (v314 - v315)
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v309)+8))
	v321 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[2]))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v311)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+8)) = v319 + (v321 - v322)
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v309)+16))
	v328 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[3]))
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v311)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+16)) = v326 + (v328 - v329)
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v309)+24))
	v335 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[4]))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v311)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+24)) = v333 + (v335 - v336)
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v309)+32))
	v342 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[5]))
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v311)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+32)) = v340 + (v342 - v343)
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v309)+40))
	v349 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[6]))
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v311)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+40)) = v347 + (v349 - v350)
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v309)+48))
	v356 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[7]))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v311)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+48)) = v354 + (v356 - v357)
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v309)+56))
	v363 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[8]))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v311)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+56)) = v361 + (v363 - v364)
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v309)+64))
	v370 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[9]))
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v311)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+64)) = v368 + (v370 - v371)
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v309)+72))
	v377 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[10]))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v311)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+72)) = v375 + (v377 - v378)
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v309)+80))
	v384 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[11]))
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v311)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+80)) = v382 + (v384 - v385)
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v309)+88))
	v391 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[12]))
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v311)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+88)) = v389 + (v391 - v392)
	v396 = *(*int64)(unsafe.Add(mBase, uint32(v309)+96))
	v398 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[13]))
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v311)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+96)) = v396 + (v398 - v399)
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v309)+104))
	v405 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[14]))
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v311)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+104)) = v403 + (v405 - v406)
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v309)+112))
	v412 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[15]))
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v311)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+112)) = v410 + (v412 - v413)
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v309)+120))
	v419 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[16]))
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v311)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+120)) = v417 + (v419 - v420)
	goto L65
L63:
	;
	goto L64
L64:
	;
	m.G0 = v15 + int32(160)
	return int32(1)
L65:
	;
	goto L64
}
func F_serializeAnalyzeStartup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	switch v6 - int32(1) {
	case 0:
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v10)
	case 1:
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v10)
	default:
	}
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeStartup[0]))
	v19 = F_AllocSetContextCreateInternal(m, v14, int32(_a_F_serializeAnalyzeStartup_0), int32(0), int32(_a_F_serializeAnalyzeStartup_1), int32(_a_F_serializeAnalyzeStartup_2))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v19
		F_initStringInfo(m, l0+int32(44))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			base.MemoryFill(m, l0-int32(-64), int32(0), int32(144))
			return
		}
	}
}
