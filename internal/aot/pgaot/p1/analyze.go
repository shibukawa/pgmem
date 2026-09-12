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
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = l4
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = int32(492)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v16
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
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = int32(41)
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[334]))
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
	v84 = *(*int32)(unsafe.Add(mBase, _consts[335]))
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
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[336])))
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
	v93 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v97 != int32(1) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v93)+392))
	if int32(1)&base.B2i32(v102 != int64(0)) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v106 = int32(4481700)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v108 + v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v112 + v109
	*(*int64)(unsafe.Add(mBase, uint32(v93)+392)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v112 + int32(2)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v123 - v109
	goto L33
}
func F_serializeAnalyzeReceive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
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
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
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
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int64
	_ = v281
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int64
	_ = v296
	var v297 int64
	_ = v297
	var v298 int64
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v342 int64
	_ = v342
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v398 int64
	_ = v398
	var v400 int64
	_ = v400
	var v401 int64
	_ = v401
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v408 int64
	_ = v408
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+9)))
	if v19 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F___clock_gettime(m, int32(1), v14+int32(8))
	mBase = m.M
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+16)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v32 = v26*int64(-1000000000) - v29
	v33 = v31
	goto L3
L2:
	;
	v32 = int64(0)
	v33 = v18
	goto L3
L3:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+7)))
	if v34 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L8
L5:
	;
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v16 == v43 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v41 = F__emscripten_memcpy_bulkmem(m, v14+int32(8), int32(4385304), int32(128))
	mBase = m.M
	goto L10
L10:
	;
	goto L7
L11:
	;
	v143 = l1 + int32(44)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v146 < v145 {
		goto L37
	} else {
		goto L38
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v45 == v17 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	F_pfree(m, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v16
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
	if v17 <= v54 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	return int32(0)
L20:
	;
	goto L18
L21:
	;
	v61 = F_palloc0(m, v17*int32(28))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v61
	v68 = v54
	goto L23
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v83 = v16 + int32(20) + v77<<(uint(int32(4))%32) + v68*int32(100)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	switch v85 {
	case 0:
		goto L26
	case 1:
		goto L28
	default:
		goto L27
	}
L24:
	;
	goto L11
L25:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	F_fmgr_info(m, v122, v84+v68*int32(28))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L19
	} else {
		goto L35
	}
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	F_getTypeOutputInfo(m, v112, v14+int32(144), v14+int32(143))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L19
	} else {
		goto L34
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L30
	}
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	F_getTypeBinaryOutputInfo(m, v86, v14+int32(144), v14+int32(143))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v121 = v14 + int32(144)
	goto L25
L30:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+24)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v102
	F_errmsg(m, int32(483699), v14)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(492449), int32(94), int32(240569))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v121 = v14 + int32(144)
	goto L25
L35:
	;
	v129 = v68 + int32(1)
	if v129 != v17 {
		v68 = v129
		goto L23
	} else {
		goto L36
	}
L36:
	;
	goto L24
L37:
	;
	F_slot_getsomeattrs_int(m, l0, v145)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L19
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v150 = int32(4487040)
	v151 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v153
	F_resetStringInfo(m, v143)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v143)+12)) = int32(68)
	goto L41
L40:
	;
	goto L39
L41:
	;
	F_enlargeStringInfo(m, v143, int32(2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v164 = int32(8)
	v170 = v17<<(uint(v164)%32) | int32(base.Ui32(v17&int32(65280))>>(uint(v164)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v161+v162))) = uint16(v170)
	v173 = v161 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v173
	if int32(0) < v17 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v180 = int32(0)
	goto L46
L44:
	;
	v272 = v173
	goto L45
L45:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+64)) = v281 + base.I64_extend_i32_s(v272)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v151
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_MemoryContextReset(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L19
	} else {
		goto L62
	}
L46:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v180))))
	if v191 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v272 = v269
	goto L45
L48:
	;
	v267 = v180 + int32(1)
	if v267 != v17 {
		v180 = v267
		goto L46
	} else {
		goto L61
	}
L49:
	;
	F_enlargeStringInfo(m, v143, int32(4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v208 = v205 + v180*int32(28)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209+v180<<(uint(int32(2))%32))))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v214 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v197+v198))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v197 + int32(4)
	goto L48
L53:
	;
	v217 = F_OutputFunctionCall(m, v208, v213)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L19
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v222 = F_SendFunctionCall(m, v208, v213)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L19
	} else {
		goto L58
	}
L56:
	;
	v219 = F_strlen(m, v217)
	mBase = m.M
	F_pq_sendcountedtext(m, v143, v217, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	goto L48
L58:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	F_enlargeStringInfo(m, v143, int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v231 = int32(2)
	v233 = int32(4)
	v234 = int32(base.Ui32(v224)>>(uint(v231)%32)) - v233
	v235 = int32(24)
	v237 = int32(65280)
	v239 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v228+v229))) = v234<<(uint(v235)%32) | v234&v237<<(uint(v239)%32) | (int32(base.Ui32(v234)>>(uint(v239)%32))&v237 | int32(base.Ui32(v234)>>(uint(v235)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v228 + v233
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	F_pq_sendbytes(m, v143, v222+v233, int32(base.Ui32(v256)>>(uint(v231)%32))-v233)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	goto L48
L61:
	;
	goto L47
L62:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+9)))
	if v291 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F___clock_gettime(m, int32(1), v14+int32(144))
	mBase = m.M
	v296 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	v297 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+152)))
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v14)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+72)) = v296 + (v297 + (v298*int64(1000000000) + v32))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v306 = v305
	goto L65
L64:
	;
	v306 = v290
	goto L65
L65:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+7)))
	if v307 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v311 = l1 + int32(80)
	v313 = v14 + int32(8)
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v311)))
	v316 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v313)))
	*(*int64)(unsafe.Add(mBase, uint32(v311))) = v314 + (v316 - v317)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v311)+8))
	v323 = *(*int64)(unsafe.Add(mBase, _consts[368]))
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v313)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+8)) = v321 + (v323 - v324)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v311)+16))
	v330 = *(*int64)(unsafe.Add(mBase, _consts[369]))
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v313)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+16)) = v328 + (v330 - v331)
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v311)+24))
	v337 = *(*int64)(unsafe.Add(mBase, _consts[370]))
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v313)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+24)) = v335 + (v337 - v338)
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v311)+32))
	v344 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v313)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+32)) = v342 + (v344 - v345)
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v311)+40))
	v351 = *(*int64)(unsafe.Add(mBase, _consts[371]))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v313)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+40)) = v349 + (v351 - v352)
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v311)+48))
	v358 = *(*int64)(unsafe.Add(mBase, _consts[372]))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v313)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+48)) = v356 + (v358 - v359)
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v311)+56))
	v365 = *(*int64)(unsafe.Add(mBase, _consts[373]))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v313)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+56)) = v363 + (v365 - v366)
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v311)+64))
	v372 = *(*int64)(unsafe.Add(mBase, _consts[374]))
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v313)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+64)) = v370 + (v372 - v373)
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v311)+72))
	v379 = *(*int64)(unsafe.Add(mBase, _consts[375]))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v313)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+72)) = v377 + (v379 - v380)
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v311)+80))
	v386 = *(*int64)(unsafe.Add(mBase, _consts[376]))
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v313)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+80)) = v384 + (v386 - v387)
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v311)+88))
	v393 = *(*int64)(unsafe.Add(mBase, _consts[377]))
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v313)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+88)) = v391 + (v393 - v394)
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v311)+96))
	v400 = *(*int64)(unsafe.Add(mBase, _consts[378]))
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v313)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+96)) = v398 + (v400 - v401)
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v311)+104))
	v407 = *(*int64)(unsafe.Add(mBase, _consts[379]))
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v313)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+104)) = v405 + (v407 - v408)
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v311)+112))
	v414 = *(*int64)(unsafe.Add(mBase, _consts[380]))
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v313)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+112)) = v412 + (v414 - v415)
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v311)+120))
	v421 = *(*int64)(unsafe.Add(mBase, _consts[381]))
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v313)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v311)+120)) = v419 + (v421 - v422)
	goto L69
L67:
	;
	goto L68
L68:
	;
	m.G0 = v14 + int32(160)
	return int32(1)
L69:
	;
	goto L68
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
	var v31 int32
	_ = v31
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v19 = F_AllocSetContextCreateInternal(m, v14, int32(342025), int32(0), int32(8192), int32(8388608))
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
			v31 = F__emscripten_memset_bulkmem(m, l0-int32(-64), base.I32_extend8_s(int32(0)), int32(144))
			mBase = m.M
			return
		}
	}
}
