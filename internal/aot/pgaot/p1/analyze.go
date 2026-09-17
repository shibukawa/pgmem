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
	var v125 int32
	_ = v125
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
	*(*int64)(unsafe.Add(mBase, uint32(v93)+392)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v114 + int32(2)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_parse_analyze_fixedparams[4])) = v125 - v111
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
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
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v271 int64
	_ = v271
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int64
	_ = v286
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v311 int64
	_ = v311
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v318 int64
	_ = v318
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v342 int64
	_ = v342
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v349 int64
	_ = v349
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v360 int64
	_ = v360
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v374 int64
	_ = v374
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v388 int64
	_ = v388
	var v390 int64
	_ = v390
	var v391 int64
	_ = v391
	var v395 int64
	_ = v395
	var v397 int64
	_ = v397
	var v398 int64
	_ = v398
	var v402 int64
	_ = v402
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v409 int64
	_ = v409
	var v411 int64
	_ = v411
	var v412 int64
	_ = v412
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
	base.MemoryCopy(m, v14+int32(8), int32(_a_F_serializeAnalyzeReceive_0), int32(128))
	goto L6
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v16 == v42 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v140 < v139 {
		goto L33
	} else {
		goto L34
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v44 == v17 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	F_pfree(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v16
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v53
	if v17 <= v53 {
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
	v60 = F_palloc0(m, v17*int32(28))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v60
	v65 = v53
	goto L19
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v82 = v16 + v74<<(uint(int32(4))%32) + v65*int32(100) + int32(20)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	switch v84 {
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
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	F_fmgr_info(m, v118, v83+v65*int32(28))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L15
	} else {
		goto L31
	}
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v82)+68))
	v111 = v14 + int32(144)
	F_getTypeOutputInfo(m, v109, v111, v14+int32(143))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L15
	} else {
		goto L30
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+68))
	v87 = v14 + int32(144)
	F_getTypeBinaryOutputInfo(m, v85, v87, v14+int32(143))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v117 = v87
	goto L21
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+24)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v99
	F_errmsg(m, int32(_a_F_serializeAnalyzeReceive_1), v14)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_serializeAnalyzeReceive_2), int32(94), int32(_a_F_serializeAnalyzeReceive_3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	v117 = v111
	goto L21
L31:
	;
	v125 = v65 + int32(1)
	if v125 != v17 {
		v65 = v125
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L20
L33:
	;
	F_slot_getsomeattrs_int(m, l0, v139)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L15
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v144 = int32(_a_F_serializeAnalyzeReceive_4)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[0]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[0])) = v147
	v150 = l1 + int32(44)
	F_resetStringInfo(m, v150)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v150)+12)) = int32(68)
	goto L37
L36:
	;
	goto L35
L37:
	;
	F_enlargeStringInfo(m, v150, int32(2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v160 = int32(8)
	v166 = v17<<(uint(v160)%32) | int32(base.Ui32(v17&int32(_a_F_serializeAnalyzeReceive_5))>>(uint(v160)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v157+v158))) = uint16(v166)
	v169 = v157 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v169
	if int32(0) < v17 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v176 = int32(0)
	goto L42
L40:
	;
	v262 = v169
	goto L41
L41:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+64)) = v271 + base.I64_extend_i32_s(v262)
	*(*int32)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[0])) = v145
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	F_MemoryContextReset(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L15
	} else {
		goto L58
	}
L42:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v176))))
	if v187 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v262 = v259
	goto L41
L44:
	;
	v257 = v176 + int32(1)
	if v257 != v17 {
		v176 = v257
		goto L42
	} else {
		goto L57
	}
L45:
	;
	F_enlargeStringInfo(m, v150, int32(4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v204 = v201 + v176*int32(28)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v205+v176<<(uint(int32(2))%32))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v210 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v194))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v193 + int32(4)
	goto L44
L49:
	;
	v213 = F_OutputFunctionCall(m, v204, v209)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L15
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v218 = F_SendFunctionCall(m, v204, v209)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L15
	} else {
		goto L54
	}
L52:
	;
	v215 = F_strlen(m, v213)
	mBase = m.M
	F_pq_sendcountedtext(m, v150, v213, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	goto L44
L54:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	F_enlargeStringInfo(m, v150, int32(4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v227 = int32(2)
	v229 = int32(4)
	v230 = int32(base.Ui32(v220)>>(uint(v227)%32)) - v229
	v231 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v224+v225))) = base.I32_rotr(v230&v231, int32(8)) | base.I32_rotr(v230, int32(24))&v231
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v224 + v229
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	F_appendBinaryStringInfo(m, v150, v218+v229, int32(base.Ui32(v246)>>(uint(v227)%32))-v229)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
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
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+9)))
	if v281 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F___clock_gettime(m, int32(1), v14+int32(144))
	mBase = m.M
	v286 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	v287 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+152)))
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v14)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+72)) = v286 + (v287 + (v288*int64(1000000000) + v32))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v296 = v295
	goto L61
L60:
	;
	v296 = v280
	goto L61
L61:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+7)))
	if v297 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v301 = l1 + int32(80)
	v303 = v14 + int32(8)
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
	v306 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[1]))
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v303)))
	*(*int64)(unsafe.Add(mBase, uint32(v301))) = v304 + (v306 - v307)
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v301)+8))
	v313 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[2]))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v303)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v311 + (v313 - v314)
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v301)+16))
	v320 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[3]))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v303)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+16)) = v318 + (v320 - v321)
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v301)+24))
	v327 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[4]))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v303)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+24)) = v325 + (v327 - v328)
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v301)+32))
	v334 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[5]))
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v303)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+32)) = v332 + (v334 - v335)
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v301)+40))
	v341 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[6]))
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v303)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+40)) = v339 + (v341 - v342)
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v301)+48))
	v348 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[7]))
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v303)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+48)) = v346 + (v348 - v349)
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v301)+56))
	v355 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[8]))
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v303)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+56)) = v353 + (v355 - v356)
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v301)+64))
	v362 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[9]))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v303)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+64)) = v360 + (v362 - v363)
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v301)+72))
	v369 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[10]))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v303)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+72)) = v367 + (v369 - v370)
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v301)+80))
	v376 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[11]))
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v303)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+80)) = v374 + (v376 - v377)
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v301)+88))
	v383 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[12]))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v303)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+88)) = v381 + (v383 - v384)
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v301)+96))
	v390 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[13]))
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v303)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+96)) = v388 + (v390 - v391)
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v301)+104))
	v397 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[14]))
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v303)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+104)) = v395 + (v397 - v398)
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v301)+112))
	v404 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[15]))
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v303)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+112)) = v402 + (v404 - v405)
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v301)+120))
	v411 = *(*int64)(unsafe.Add(mBase, _c_F_serializeAnalyzeReceive[16]))
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v303)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+120)) = v409 + (v411 - v412)
	goto L65
L63:
	;
	goto L64
L64:
	;
	m.G0 = v14 + int32(160)
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
