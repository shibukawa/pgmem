package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v408 int64
	_ = v408
	var v412 int32
	_ = v412
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v419 int32
	_ = v419
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v478 int64
	_ = v478
	var v479 int64
	_ = v479
	var v482 int64
	_ = v482
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int64
	_ = v515
	var v521 int64
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int64
	_ = v532
	var v538 int64
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int64
	_ = v548
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int64
	_ = v562
	var v570 int64
	_ = v570
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 == v2 {
		v38 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v38&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v25 == int32(0) {
		v38 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v28 != int32(7) {
		v38 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v31 != int32(17) {
		v38 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+32)))
	v38 = v34 ^ int32(1)
	goto L2
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = F_get_fn_opclass_options(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v47 = int32(100)
	goto L9
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+18)))
	if v49 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	return int64(0)
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v47 = v46
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L10
	} else {
		goto L149
	}
L13:
	;
	m.G0 = v17 + int32(16)
	return base.I64_extend_i32_u(v600)
L14:
	;
	if v270 < int32(0) {
		v349 = v273
		goto L91
	} else {
		goto L92
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L87
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L82
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L78
	}
L18:
	;
	v52 = F_pg_detoast_datum_copy(m, v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v100 = F_pg_detoast_datum(m, v48)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L35
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v55 = F_array_contains_nulls(m, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v60 = F_ArrayGetNItemsSafe(m, v57, v52+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L27
	}
L25:
	;
	if v55 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)) = uint8(v62)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v64 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v72 = v64
	goto L30
L29:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v72 = (v65<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L30
L30:
	;
	F_isort(m, v72+v52, v60, v17+int32(15))
	mBase = m.M
	v77 = F__int_unique(m, v52)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v81 = v77 + int32(16)
	v82 = F_ArrayGetNItemsSafe(m, v79, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	v85 = v47 << (uint(int32(1)) % 32)
	if v85 <= v82 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v88 = F_palloc(m, int32(24))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = base.I64_extend_i32_u(v77)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = v94
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)))
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+18)) = uint8(v97)
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+16)) = uint16(v96)
	v600 = v88
	goto L13
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v103 = F_array_contains_nulls(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v107 = v100 + int32(16)
	v108 = F_ArrayGetNItemsSafe(m, v105, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	if v103 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v108 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v112 == v100 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v117 = F_ArrayGetNItemsSafe(m, v116, v107)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L49
	}
L45:
	;
	v600 = v19
	goto L13
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v100)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	v600 = v19
	goto L13
L49:
	;
	v120 = v47 << (uint(int32(1)) % 32)
	if v117 < v120 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v600 = v19
	goto L13
L51:
	;
	goto L52
L52:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v122 == v100 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v124 = F_pg_detoast_datum_copy(m, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L10
	} else {
		goto L56
	}
L54:
	;
	v126 = v100
	goto L55
L55:
	;
	v129 = F_resize_intArrayType(m, v126, v117<<(uint(int32(1))%32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	v126 = v124
	goto L55
L57:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v131 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v141 = (v134<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L60
L59:
	;
	v141 = v131
	goto L60
L60:
	;
	v143 = v117 - int32(1)
	v144 = v141 + v129
	if int32(2) <= v117 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v151 = v143
	v153 = v147
	v155 = v143
	goto L66
L62:
	;
	v147 = v117 - v47
	if int32(0) < v147 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v270 = v143
	v273 = v143
	goto L14
L65:
	;
	goto L64
L66:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144+v151<<(uint(int32(2))%32))))
	v169 = v151
	v170 = v168
	v171 = v153
	goto L68
L67:
	;
	v270 = v210
	v273 = v208
	goto L14
L68:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v144+v169<<(uint(int32(2))%32)-int32(4))))
	if v188 != v170-int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v204 = v144 + v155<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v168
	v207 = int32(1)
	v208 = v155 - v207
	v210 = v198 - v207
	if v198 < int32(2) {
		v270 = v210
		v273 = v208
		goto L14
	} else {
		goto L76
	}
L70:
	;
	goto L69
L71:
	;
	v198 = v169
	v200 = v171
	v201 = v170
	goto L70
L72:
	;
	goto L73
L73:
	;
	v192 = int32(1)
	v193 = v171 - v192
	v195 = v169 - v192
	if v195 == int32(0) {
		v198 = v195
		v200 = v193
		v201 = v188
		goto L70
	} else {
		goto L74
	}
L74:
	;
	if v193 != 0 {
		v169 = v195
		v170 = v188
		v171 = v193
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v198 = v195
	v200 = v193
	v201 = v188
	goto L70
L76:
	;
	if int32(0) < v200 {
		v151 = v210
		v153 = v200
		v155 = v208
		goto L66
	} else {
		goto L77
	}
L77:
	;
	goto L67
L78:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_g_int_compress_0), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(181), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v239 = F_ArrayGetNItemsSafe(m, v238, v81)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v85 - int32(1)
	F_errmsg(m, int32(_a_F_g_int_compress_3), v17)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(188), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(_a_F_g_int_compress_0), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(203), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v359 = int32(1)
	v361 = v349 + v359
	if v361 == int32(0) {
		v374 = v117
		goto L100
	} else {
		goto L101
	}
L92:
	;
	if v270&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v302 = v273
	v303 = v270
	goto L95
L94:
	;
	v289 = v144 + v273<<(uint(int32(3))%32)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v144+v270<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v293
	v296 = int32(1)
	v302 = v273 - v296
	v303 = v270 - v296
	goto L95
L95:
	;
	if v270 == int32(0) {
		v349 = v302
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v306 = v303
	v310 = v302
	goto L97
L97:
	;
	v322 = v144 + v310<<(uint(int32(3))%32)
	v323 = int32(2)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v144+v306<<(uint(v323)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v326
	v332 = v306 - int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v144+v332<<(uint(v323)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v322-int32(8)))) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v322-int32(4)))) = v336
	v342 = v310 - v323
	if v332 != 0 {
		v306 = v306 - v323
		v310 = v342
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v349 = v342
	goto L91
L99:
	;
	goto L98
L100:
	;
	v376 = v374 << (uint(int32(1)) % 32)
	if v120 < v376 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v364 = v117 - v361
	v366 = v364 << (uint(int32(3)) % 32)
	if v366 == int32(0) {
		v374 = v364
		goto L100
	} else {
		goto L102
	}
L102:
	;
	base.MemoryCopy(m, v144, v144+v361<<(uint(int32(3))%32), v366)
	v374 = v364
	goto L100
L103:
	;
	v380 = v359
	v381 = v376
	goto L106
L104:
	;
	v459 = v376
	goto L105
L105:
	;
	v471 = int32(0)
	if v459 <= v471 {
		v570 = int64(0)
		goto L125
	} else {
		goto L126
	}
L106:
	;
	if int32(3) <= v381 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v459 = v454
	goto L105
L108:
	;
	v396 = int32(2)
	v398 = v380
	v408 = int64(9223372036854775807)
	goto L111
L109:
	;
	v427 = v380
	goto L110
L110:
	;
	v443 = (v381 + (v427 ^ int32(-1))) << (uint(int32(2)) % 32)
	if v443 != 0 {
		goto L120
	} else {
		goto L121
	}
L111:
	;
	v412 = v144 + v396<<(uint(int32(2))%32)
	v413 = int64(*(*int32)(unsafe.Add(mBase, uint32(v412))))
	v416 = int64(*(*int32)(unsafe.Add(mBase, uint32(v412-int32(4)))))
	v417 = v413 - v416
	if v417 < v408 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v427 = v419
	goto L110
L113:
	;
	v419 = v396
	goto L115
L114:
	;
	v419 = v398
	goto L115
L115:
	;
	if v408 < v417 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v421 = v408
	goto L118
L117:
	;
	v421 = v417
	goto L118
L118:
	;
	v423 = v396 + int32(2)
	if v423 < v381 {
		v396 = v423
		v398 = v419
		v408 = v421
		goto L111
	} else {
		goto L119
	}
L119:
	;
	goto L112
L120:
	;
	v446 = v144 + v427<<(uint(int32(2))%32)
	v447 = int32(4)
	base.MemoryCopy(m, v446-v447, v446+v447, v443)
	goto L122
L121:
	;
	goto L122
L122:
	;
	v454 = v381 - int32(2)
	if v120 < v454 {
		v380 = v427
		v381 = v454
		goto L106
	} else {
		goto L123
	}
L123:
	;
	goto L107
L124:
	;
	if base.Ui32(int32(67108864)) <= base.Ui32(v582) {
		goto L12
	} else {
		goto L146
	}
L125:
	;
	if base.Ui64(v570-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L143
	} else {
		goto L144
	}
L126:
	;
	v478 = int64(*(*int32)(unsafe.Add(mBase, uint32(v144)+4)))
	v479 = int64(*(*int32)(unsafe.Add(mBase, uint32(v144))))
	v482 = v478 - v479 + int64(1)
	if base.Ui32(v459) < base.Ui32(int32(3)) {
		v570 = v482
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v488 = int32(base.Ui32(v459-int32(3)) >> (uint(int32(1)) % 32))
	if v488 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v556 = v144 + v547<<(uint(int32(2))%32)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556-int32(4))))
	if v557 == v560 {
		v570 = v548
		goto L125
	} else {
		goto L142
	}
L129:
	;
	v547 = int32(2)
	v548 = v482
	goto L128
L130:
	;
	goto L131
L131:
	;
	v492 = int32(1)
	v493 = v488 + v492
	v500 = int32(2)
	v501 = v482
	v506 = v471
	goto L132
L132:
	;
	v509 = v144 + v500<<(uint(int32(2))%32)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v509-int32(4))))
	if v510 != v513 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v493&v492 == int32(0) {
		v570 = v538
		goto L125
	} else {
		goto L141
	}
L134:
	;
	v515 = int64(*(*int32)(unsafe.Add(mBase, uint32(v509)+4)))
	v521 = v501 + v515 - base.I64_extend_i32_s(v510) + int64(1)
	goto L136
L135:
	;
	v521 = v501
	goto L136
L136:
	;
	v522 = int32(2)
	v526 = v144 + (v500+v522)<<(uint(v522)%32)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v526-int32(4))))
	if v527 != v530 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v532 = int64(*(*int32)(unsafe.Add(mBase, uint32(v526)+4)))
	v538 = v521 + v532 - base.I64_extend_i32_s(v527) + int64(1)
	goto L139
L138:
	;
	v538 = v521
	goto L139
L139:
	;
	v540 = v500 + int32(4)
	v542 = v506 + int32(2)
	if v542 != v493&int32(-2) {
		v500 = v540
		v501 = v538
		v506 = v542
		goto L132
	} else {
		goto L140
	}
L140:
	;
	goto L133
L141:
	;
	v547 = v540
	v548 = v538
	goto L128
L142:
	;
	v562 = int64(*(*int32)(unsafe.Add(mBase, uint32(v556)+4)))
	v570 = v548 + v562 - base.I64_extend_i32_s(v557) + int64(1)
	goto L125
L143:
	;
	v582 = int32(-1)
	goto L145
L144:
	;
	v582 = base.I32_wrap_i64(v570)
	goto L145
L145:
	;
	goto L124
L146:
	;
	v585 = F_resize_intArrayType(m, v129, v459)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	v588 = F_palloc(m, int32(24))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v588))) = base.I64_extend_i32_u(v585)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+8)) = v592
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v588)+12)) = v594
	v596 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)))
	v597 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v588)+18)) = uint8(v597)
	*(*uint16)(unsafe.Add(mBase, uint32(v588)+16)) = uint16(v596)
	v600 = v588
	goto L13
L149:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_g_int_compress_4), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(277), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_g_int_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_g_int_options_0), int32(_a_F_g_int_options_1), int32(100), int32(1), int32(252))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
