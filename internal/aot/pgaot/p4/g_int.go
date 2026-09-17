package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_compress(m *base.Module, l0 int32) int32 {
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v407 int64
	_ = v407
	var v411 int32
	_ = v411
	var v412 int64
	_ = v412
	var v415 int64
	_ = v415
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v470 int32
	_ = v470
	var v477 int64
	_ = v477
	var v478 int64
	_ = v478
	var v481 int64
	_ = v481
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int64
	_ = v500
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int64
	_ = v514
	var v520 int64
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int64
	_ = v531
	var v537 int64
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int64
	_ = v561
	var v569 int64
	_ = v569
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+24)))
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
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+14)))
	if v49 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v47 = v46
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L10
	} else {
		goto L149
	}
L13:
	;
	m.G0 = v17 + int32(16)
	return v598
L14:
	;
	if v269 < int32(0) {
		v348 = v272
		goto L91
	} else {
		goto L92
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L87
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L82
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
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
	v99 = F_pg_detoast_datum(m, v48)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
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
	v88 = F_palloc(m, int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v77
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v93
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)))
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+14)) = uint8(v96)
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+12)) = uint16(v95)
	v598 = v88
	goto L13
L35:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v101 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = F_array_contains_nulls(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L10
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v106 = v99 + int32(16)
	v107 = F_ArrayGetNItemsSafe(m, v104, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	if v102 != 0 {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v107 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v111 == v99 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v116 = F_ArrayGetNItemsSafe(m, v115, v106)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L49
	}
L45:
	;
	v598 = v19
	goto L13
L46:
	;
	goto L47
L47:
	;
	F_pfree(m, v99)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	v598 = v19
	goto L13
L49:
	;
	v119 = v47 << (uint(int32(1)) % 32)
	if v116 < v119 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v598 = v19
	goto L13
L51:
	;
	goto L52
L52:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v121 == v99 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v123 = F_pg_detoast_datum_copy(m, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L56
	}
L54:
	;
	v125 = v99
	goto L55
L55:
	;
	v128 = F_resize_intArrayType(m, v125, v116<<(uint(int32(1))%32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	v125 = v123
	goto L55
L57:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v130 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v140 = (v133<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L60
L59:
	;
	v140 = v130
	goto L60
L60:
	;
	v142 = v116 - int32(1)
	v143 = v140 + v128
	if int32(2) <= v116 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v150 = v142
	v152 = v146
	v154 = v142
	goto L66
L62:
	;
	v146 = v116 - v47
	if int32(0) < v146 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v269 = v142
	v272 = v142
	goto L14
L65:
	;
	goto L64
L66:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v143+v150<<(uint(int32(2))%32))))
	v168 = v150
	v169 = v167
	v170 = v152
	goto L68
L67:
	;
	v269 = v209
	v272 = v207
	goto L14
L68:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v143+v168<<(uint(int32(2))%32)-int32(4))))
	if v187 != v169-int32(1) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v203 = v143 + v154<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v167
	v206 = int32(1)
	v207 = v154 - v206
	v209 = v197 - v206
	if v197 < int32(2) {
		v269 = v209
		v272 = v207
		goto L14
	} else {
		goto L76
	}
L70:
	;
	goto L69
L71:
	;
	v197 = v168
	v199 = v170
	v200 = v169
	goto L70
L72:
	;
	goto L73
L73:
	;
	v191 = int32(1)
	v192 = v170 - v191
	v194 = v168 - v191
	if v194 == int32(0) {
		v197 = v194
		v199 = v192
		v200 = v187
		goto L70
	} else {
		goto L74
	}
L74:
	;
	if v192 != 0 {
		v168 = v194
		v169 = v187
		v170 = v192
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v197 = v194
	v199 = v192
	v200 = v187
	goto L70
L76:
	;
	if int32(0) < v199 {
		v150 = v209
		v152 = v199
		v154 = v207
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_g_int_compress_0), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(180), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v238 = F_ArrayGetNItemsSafe(m, v237, v81)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v85 - int32(1)
	F_errmsg(m, int32(_a_F_g_int_compress_3), v17)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(187), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
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
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(_a_F_g_int_compress_0), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(202), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	v358 = int32(1)
	v360 = v348 + v358
	if v360 == int32(0) {
		v373 = v116
		goto L100
	} else {
		goto L101
	}
L92:
	;
	if v269&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v301 = v272
	v302 = v269
	goto L95
L94:
	;
	v288 = v143 + v272<<(uint(int32(3))%32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v143+v269<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v292
	v295 = int32(1)
	v301 = v272 - v295
	v302 = v269 - v295
	goto L95
L95:
	;
	if v269 == int32(0) {
		v348 = v301
		goto L91
	} else {
		goto L96
	}
L96:
	;
	v305 = v302
	v309 = v301
	goto L97
L97:
	;
	v321 = v143 + v309<<(uint(int32(3))%32)
	v322 = int32(2)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v143+v305<<(uint(v322)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v321)+4)) = v325
	v331 = v305 - int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v143+v331<<(uint(v322)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v321-int32(8)))) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v321-int32(4)))) = v335
	v341 = v309 - v322
	if v331 != 0 {
		v305 = v305 - v322
		v309 = v341
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v348 = v341
	goto L91
L99:
	;
	goto L98
L100:
	;
	v375 = v373 << (uint(int32(1)) % 32)
	if v119 < v375 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v363 = v116 - v360
	v365 = v363 << (uint(int32(3)) % 32)
	if v365 == int32(0) {
		v373 = v363
		goto L100
	} else {
		goto L102
	}
L102:
	;
	base.MemoryCopy(m, v143, v143+v360<<(uint(int32(3))%32), v365)
	v373 = v363
	goto L100
L103:
	;
	v379 = v358
	v380 = v375
	goto L106
L104:
	;
	v458 = v375
	goto L105
L105:
	;
	v470 = int32(0)
	if v458 <= v470 {
		v569 = int64(0)
		goto L125
	} else {
		goto L126
	}
L106:
	;
	if int32(3) <= v380 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v458 = v453
	goto L105
L108:
	;
	v395 = int32(2)
	v397 = v379
	v407 = int64(9223372036854775807)
	goto L111
L109:
	;
	v426 = v379
	goto L110
L110:
	;
	v442 = (v380 + (v426 ^ int32(-1))) << (uint(int32(2)) % 32)
	if v442 != 0 {
		goto L120
	} else {
		goto L121
	}
L111:
	;
	v411 = v143 + v395<<(uint(int32(2))%32)
	v412 = int64(*(*int32)(unsafe.Add(mBase, uint32(v411))))
	v415 = int64(*(*int32)(unsafe.Add(mBase, uint32(v411-int32(4)))))
	v416 = v412 - v415
	if v416 < v407 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v426 = v418
	goto L110
L113:
	;
	v418 = v395
	goto L115
L114:
	;
	v418 = v397
	goto L115
L115:
	;
	if v407 < v416 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v420 = v407
	goto L118
L117:
	;
	v420 = v416
	goto L118
L118:
	;
	v422 = v395 + int32(2)
	if v422 < v380 {
		v395 = v422
		v397 = v418
		v407 = v420
		goto L111
	} else {
		goto L119
	}
L119:
	;
	goto L112
L120:
	;
	v445 = v143 + v426<<(uint(int32(2))%32)
	v446 = int32(4)
	base.MemoryCopy(m, v445-v446, v445+v446, v442)
	goto L122
L121:
	;
	goto L122
L122:
	;
	v453 = v380 - int32(2)
	if v119 < v453 {
		v379 = v426
		v380 = v453
		goto L106
	} else {
		goto L123
	}
L123:
	;
	goto L107
L124:
	;
	if base.Ui32(int32(134217725)) <= base.Ui32(v581) {
		goto L12
	} else {
		goto L146
	}
L125:
	;
	if base.Ui64(v569-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L143
	} else {
		goto L144
	}
L126:
	;
	v477 = int64(*(*int32)(unsafe.Add(mBase, uint32(v143)+4)))
	v478 = int64(*(*int32)(unsafe.Add(mBase, uint32(v143))))
	v481 = v477 - v478 + int64(1)
	if base.Ui32(v458) < base.Ui32(int32(3)) {
		v569 = v481
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v487 = int32(base.Ui32(v458-int32(3)) >> (uint(int32(1)) % 32))
	if v487 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v555 = v143 + v546<<(uint(int32(2))%32)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555-int32(4))))
	if v556 == v559 {
		v569 = v547
		goto L125
	} else {
		goto L142
	}
L129:
	;
	v546 = int32(2)
	v547 = v481
	goto L128
L130:
	;
	goto L131
L131:
	;
	v491 = int32(1)
	v492 = v487 + v491
	v499 = int32(2)
	v500 = v481
	v505 = v470
	goto L132
L132:
	;
	v508 = v143 + v499<<(uint(int32(2))%32)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508-int32(4))))
	if v509 != v512 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	if v492&v491 == int32(0) {
		v569 = v537
		goto L125
	} else {
		goto L141
	}
L134:
	;
	v514 = int64(*(*int32)(unsafe.Add(mBase, uint32(v508)+4)))
	v520 = v500 + v514 - base.I64_extend_i32_s(v509) + int64(1)
	goto L136
L135:
	;
	v520 = v500
	goto L136
L136:
	;
	v521 = int32(2)
	v525 = v143 + (v499+v521)<<(uint(v521)%32)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v525-int32(4))))
	if v526 != v529 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v531 = int64(*(*int32)(unsafe.Add(mBase, uint32(v525)+4)))
	v537 = v520 + v531 - base.I64_extend_i32_s(v526) + int64(1)
	goto L139
L138:
	;
	v537 = v520
	goto L139
L139:
	;
	v539 = v499 + int32(4)
	v541 = v505 + int32(2)
	if v541 != v492&int32(-2) {
		v499 = v539
		v500 = v537
		v505 = v541
		goto L132
	} else {
		goto L140
	}
L140:
	;
	goto L133
L141:
	;
	v546 = v539
	v547 = v537
	goto L128
L142:
	;
	v561 = int64(*(*int32)(unsafe.Add(mBase, uint32(v555)+4)))
	v569 = v547 + v561 - base.I64_extend_i32_s(v556) + int64(1)
	goto L125
L143:
	;
	v581 = int32(-1)
	goto L145
L144:
	;
	v581 = base.I32_wrap_i64(v569)
	goto L145
L145:
	;
	goto L124
L146:
	;
	v584 = F_resize_intArrayType(m, v128, v458)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	v587 = F_palloc(m, int32(16))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v584
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v587)+4)) = v590
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v587)+8)) = v592
	v594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)))
	v595 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v587)+14)) = uint8(v595)
	*(*uint16)(unsafe.Add(mBase, uint32(v587)+12)) = uint16(v594)
	v598 = v587
	goto L13
L149:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_g_int_compress_4), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_g_int_compress_1), int32(276), int32(_a_F_g_int_compress_2))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
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
func F_g_int_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_g_int_options_0), int32(_a_F_g_int_options_1), int32(100), int32(1), int32(252))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
