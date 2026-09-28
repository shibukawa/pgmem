package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_logicalrep_read_prepare_common(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L54
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L51
	}
L4:
	;
	return
L5:
	;
	v14 = v11 & int32(255)
	if v14 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v17 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L48
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v17
	if v17 == int64(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v22 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v22
	if v22 == int64(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v27 = F_pq_getmsgint64(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v27
	v31 = F_pq_getmsgint(m, l0, int32(4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v31
	if v31 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v37 = l2 + int32(28)
	v38 = F_pq_getmsgstring(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L20
L17:
	;
	m.G0 = v9 - int32(-64)
	return
L18:
	;
	v156 = F_strlen(m, v145)
	mBase = m.M
	goto L17
L20:
	;
	goto L21
L21:
	;
	v46 = int32(199)
	if (v37^v38)&int32(3) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v146))) = uint8(v149)
	goto L18
L23:
	;
	v130 = v125
	v131 = v126
	v132 = v127
	goto L44
L24:
	;
	if v120 == int32(0) {
		v145 = v118
		v146 = v119
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v118 = v38
	v119 = v37
	v120 = v46
	goto L24
L26:
	;
	goto L27
L27:
	;
	v50 = int32(0)
	if base.B2i32(v38&int32(3) == v50)|int32(0) == v50 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v86 == int32(0) {
		v145 = v83
		v146 = v84
		goto L22
	} else {
		goto L37
	}
L29:
	;
	v62 = v38
	v63 = v37
	v64 = v46
	goto L32
L30:
	;
	goto L31
L31:
	;
	v83 = v38
	v84 = v37
	v85 = v46
	v86 = int32(1)
	goto L28
L32:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v66)
	if v66 == int32(0) {
		v125 = v62
		v126 = v63
		v127 = v64
		goto L23
	} else {
		goto L34
	}
L33:
	;
	v83 = v77
	v84 = v71
	v85 = v73
	v86 = v75
	goto L28
L34:
	;
	v70 = int32(1)
	v71 = v63 + v70
	v73 = v64 - v70
	v74 = int32(0)
	v75 = base.B2i32(v73 != v74)
	v77 = v62 + v70
	if v77&int32(3) == v74 {
		v83 = v77
		v84 = v71
		v85 = v73
		v86 = v75
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v73 != 0 {
		v62 = v77
		v63 = v71
		v64 = v73
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if base.B2i32(v89 == int32(0))|base.B2i32(base.Ui32(v85) < base.Ui32(int32(4))) != 0 {
		v118 = v83
		v119 = v84
		v120 = v85
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v96 = v83
	v97 = v84
	v98 = v85
	goto L39
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 != v104 {
		v125 = v96
		v126 = v97
		v127 = v98
		goto L23
	} else {
		goto L41
	}
L40:
	;
	v118 = v112
	v119 = v110
	v120 = v114
	goto L24
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v101
	v109 = int32(4)
	v110 = v97 + v109
	v112 = v96 + v109
	v114 = v98 - v109
	if base.Ui32(int32(3)) < base.Ui32(v114) {
		v96 = v112
		v97 = v110
		v98 = v114
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v125 = v118
	v126 = v119
	v127 = v120
	goto L23
L44:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v134)
	if v134 == int32(0) {
		v145 = v130
		v146 = v131
		goto L22
	} else {
		goto L46
	}
L45:
	;
	v145 = v141
	v146 = v139
	goto L22
L46:
	;
	v138 = int32(1)
	v139 = v131 + v138
	v141 = v130 + v138
	v143 = v132 - v138
	if v143 != 0 {
		v130 = v141
		v131 = v139
		v132 = v143
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v14
	F_errmsg_internal(m, int32(_a_F_logicalrep_read_prepare_common_0), v7+int32(-16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_logicalrep_read_prepare_common_1), int32(206), int32(_a_F_logicalrep_read_prepare_common_2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
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
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg_internal(m, int32(_a_F_logicalrep_read_prepare_common_3), v9)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_logicalrep_read_prepare_common_1), int32(211), int32(_a_F_logicalrep_read_prepare_common_2))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_logicalrep_read_prepare_common_4), v7+int32(-48))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_logicalrep_read_prepare_common_1), int32(214), int32(_a_F_logicalrep_read_prepare_common_2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	F_errmsg_internal(m, int32(_a_F_logicalrep_read_prepare_common_5), v7+int32(-32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_logicalrep_read_prepare_common_1), int32(218), int32(_a_F_logicalrep_read_prepare_common_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logicalrep_rel_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
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
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int64
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v298 int64
	_ = v298
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v316 int64
	_ = v316
	var v321 int32
	_ = v321
	var v322 int64
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
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
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int64
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int64
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v489 int64
	_ = v489
	var v493 int64
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v501 int64
	_ = v501
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l0
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[0]))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v55 = v19
	goto L3
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[1]))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v61 = F_hash_search(m, v55, v15+int32(76), int32(0), v15+int32(80))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L11
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[2]))
	v31 = F_AllocSetContextCreateInternal(m, v26, int32(_a_F_logicalrep_rel_open_0), int32(0), int32(_a_F_logicalrep_rel_open_1), int32(_a_F_logicalrep_rel_open_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v36 = v21
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = int64(309237645316)
	v46 = F_hash_create(m, int32(_a_F_logicalrep_rel_open_3), int64(128), v15+int32(80), int32(1064))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[1])) = v31
	v36 = v31
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[0])) = v46
	F_CacheRegisterRelcacheCallback(m, int32(1085))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[0]))
	v55 = v53
	goto L3
L11:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+80)))
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L7
	} else {
		goto L112
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L7
	} else {
		goto L108
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L7
	} else {
		goto L105
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	if v64 != 0 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L102
	}
L18:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
	if v65 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+56)))
	if v377 != int32(114) {
		goto L98
	} else {
		goto L99
	}
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	v69 = F_try_table_open(m, v68, l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v69
	if v69 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)) = uint8(v74)
	goto L20
L24:
	;
	goto L25
L25:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
	if v76 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	F_relation_close(m, v69, l1)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = int32(0)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
	if v81 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L20
L29:
	;
	F_free_attrmap(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v91 = F_makeRangeVar(m, v88, v89, int32(-1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = int32(0)
	goto L31
L33:
	;
	v94 = int32(0)
	v96 = F_RangeVarGetRelidExtended(m, v91, l1, int32(1), v94, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	if v96 == int32(0) {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	v101 = F_table_open(m, v96, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v101
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+48))
	v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+119)))
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61)+25)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	F_CheckSubscriptionRelkind(m, v106, v107, v108, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v112 = int32(_a_F_logicalrep_rel_open_4)
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[3]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+52))
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[3])) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v121 = F_make_attrmap(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v121
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[3])) = v113
	v126 = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v131 = F_bms_add_range(m, v126, v126, v128-int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if int32(0) < v133 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	F_bms_free(m, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L94
	}
L41:
	;
	v137 = v133
	v139 = v131
	v140 = v3
	v141 = v3
	goto L44
L42:
	;
	goto L43
L43:
	;
	if v131 != 0 {
		v445 = v131
		goto L12
	} else {
		goto L93
	}
L44:
	;
	v153 = v115 + v137<<(uint(int32(3))%32) + v140*int32(100)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+119)))
	if v154 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v254 != 0 {
		v445 = v254
		goto L12
	} else {
		goto L71
	}
L46:
	;
	v264 = v140 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v264 < v265 {
		v137 = v265
		v139 = v254
		v140 = v264
		v141 = v256
		goto L44
	} else {
		goto L70
	}
L47:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v162 = int32(_a_F_logicalrep_rel_open_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v158+v140<<(uint(int32(1))%32)))) = uint16(v162)
	v254 = v139
	v256 = v141
	goto L46
L48:
	;
	goto L49
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if int32(0) < v164 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	*(*uint16)(unsafe.Add(mBase, uint32(v240+v140<<(uint(int32(1))%32)))) = uint16(v174)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+int32(28))+90)))
	if v245 != 0 {
		goto L65
	} else {
		goto L66
	}
L51:
	;
	v170 = v153 + int32(32)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v174 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v237 = int32(_a_F_logicalrep_rel_open_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v233+v140<<(uint(int32(1))%32)))) = uint16(v237)
	v254 = v139
	v256 = v141
	goto L46
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v171+v174<<(uint(int32(2))%32))))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if base.B2i32(v191 == int32(0))|base.B2i32(v191 != v194) != 0 {
		v212 = v191
		v213 = v194
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L53
L56:
	;
	if v212-v213 == int32(0) {
		goto L50
	} else {
		goto L63
	}
L57:
	;
	goto L56
L58:
	;
	v197 = v188
	v198 = v170
	goto L59
L59:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v202 == int32(0) {
		v212 = v202
		v213 = v201
		goto L57
	} else {
		goto L61
	}
L60:
	;
	v212 = v202
	v213 = v201
	goto L57
L61:
	;
	v205 = int32(1)
	if v202 == v201 {
		v197 = v197 + v205
		v198 = v198 + v205
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v218 = v174 + int32(1)
	if v218 != v164 {
		v174 = v218
		goto L54
	} else {
		goto L64
	}
L64:
	;
	goto L55
L65:
	;
	v246 = F_bms_add_member(m, v141, v174)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L68
	}
L66:
	;
	v248 = v141
	goto L67
L67:
	;
	v249 = F_bms_del_member(m, v139, v174)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L69
	}
L68:
	;
	v248 = v246
	goto L67
L69:
	;
	v254 = v249
	v256 = v248
	goto L46
L70:
	;
	goto L45
L71:
	;
	if v256 == int32(0) {
		goto L40
	} else {
		goto L72
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v277 = int64(0)
	if v256 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v61)+4))
	v323 = F_logicalrep_get_attrs_str(m, v61, v256)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L7
	} else {
		goto L90
	}
L76:
	;
	v321 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v282 = v256 + int32(8)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v283 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v321 = base.I32_popcnt(v286)
	goto L75
L80:
	;
	goto L81
L81:
	;
	v289 = v283 << (uint(int32(2)) % 32)
	if v289 <= int32(7) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v321 = base.I32_wrap_i64(v316)
	goto L75
L83:
	;
	if v289 == int32(0) {
		v316 = v277
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v313 = F_pg_popcount_optimized(m, v282, v289)
	mBase = m.M
	v316 = v313
	goto L82
L86:
	;
	v294 = v289
	v295 = v282
	v296 = v277
	goto L87
L87:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+3)))
	v298 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v297)+uint32(_c_F_logicalrep_rel_open[4]))))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)))
	v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v299)+uint32(_c_F_logicalrep_rel_open[4]))))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	v302 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v301)+uint32(_c_F_logicalrep_rel_open[4]))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v304 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v303)+uint32(_c_F_logicalrep_rel_open[4]))))
	v308 = v298 + (v300 + (v302 + (v296 + v304)))
	v309 = int32(4)
	v312 = v294 - v309
	if v312 != 0 {
		v294 = v312
		v295 = v295 + v309
		v296 = v308
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v316 = v308
	goto L82
L89:
	;
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v322
	F_errmsg_plural(m, int32(_a_F_logicalrep_rel_open_6), int32(_a_F_logicalrep_rel_open_7), v321, v15+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(291), int32(_a_F_logicalrep_rel_open_9))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	goto L40
L94:
	;
	F_bms_free(m, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_logicalrep_rel_mark_updatable(m, v61)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v360 = F_FindLogicalRepLocalIndex(m, v358, v61, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	v362 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)) = uint8(v362)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+52)) = v360
	goto L19
L98:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[5]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	v386 = F_GetSubscriptionRelState(m, v382, v383, v61-int32(-64))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	m.G0 = v15 + int32(128)
	return v61
L101:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+56)) = uint8(v386)
	goto L100
L102:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v397
	F_errmsg_internal(m, int32(_a_F_logicalrep_rel_open_10), v15-int32(-64))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(376), int32(_a_F_logicalrep_rel_open_11))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v413
	F_errmsg_internal(m, int32(_a_F_logicalrep_rel_open_12), v15+int32(48))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(382), int32(_a_F_logicalrep_rel_open_11))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v432 = *(*int64)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v432
	F_errmsg(m, int32(_a_F_logicalrep_rel_open_13), v15)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(434), int32(_a_F_logicalrep_rel_open_11))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	v462 = int64(0)
	if v445 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v507 = *(*int64)(unsafe.Add(mBase, uint32(v61)+4))
	v508 = F_logicalrep_get_attrs_str(m, v61, v445)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L7
	} else {
		goto L129
	}
L115:
	;
	v506 = int32(0)
	goto L114
L116:
	;
	goto L117
L117:
	;
	v467 = v445 + int32(8)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	if v468 == int32(1) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v467)))
	v506 = base.I32_popcnt(v471)
	goto L114
L119:
	;
	goto L120
L120:
	;
	v474 = v468 << (uint(int32(2)) % 32)
	if v474 <= int32(7) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v506 = base.I32_wrap_i64(v501)
	goto L114
L122:
	;
	if v474 == int32(0) {
		v501 = v462
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v498 = F_pg_popcount_optimized(m, v467, v474)
	mBase = m.M
	v501 = v498
	goto L121
L125:
	;
	v479 = v474
	v480 = v467
	v481 = v462
	goto L126
L126:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+3)))
	v483 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482)+uint32(_c_F_logicalrep_rel_open[4]))))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+2)))
	v485 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v484)+uint32(_c_F_logicalrep_rel_open[4]))))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+1)))
	v487 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v486)+uint32(_c_F_logicalrep_rel_open[4]))))
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v489 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_logicalrep_rel_open[4]))))
	v493 = v483 + (v485 + (v487 + (v481 + v489)))
	v494 = int32(4)
	v497 = v479 - v494
	if v497 != 0 {
		v479 = v497
		v480 = v480 + v494
		v481 = v493
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v501 = v493
	goto L121
L128:
	;
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v507
	F_errmsg_plural(m, int32(_a_F_logicalrep_rel_open_14), int32(_a_F_logicalrep_rel_open_15), v506, v15+int32(32))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(280), int32(_a_F_logicalrep_rel_open_9))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_logicalrep_sync_worker_count(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_sync_worker_count[0]))
	if v10 <= v2 {
		v89 = v2
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_sync_worker_count[1]))
		v16 = v14 + int32(16)
		if v10 != int32(1) {
			v25 = v2
			v26 = v2
			v30 = v2
			for {
				v33 = v16 + v26<<(uint(int32(7))%32)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
				if v34 != l0 {
					v45 = v25
				} else {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+16)))
					if v36 != int32(1) {
						v45 = v25
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v45 = v25 + base.B2i32(base.Ui32(v39-int32(1)) < base.Ui32(int32(2)))
					}
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+160))
				if v46 != l0 {
					v57 = v45
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+144)))
					if v48 != int32(1) {
						v57 = v45
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)+128))
						v57 = v45 + base.B2i32(base.Ui32(v51-int32(1)) < base.Ui32(int32(2)))
					}
				}
				v58 = int32(2)
				v59 = v26 + v58
				v61 = v30 + v58
				if v61 != v10&int32(2147483646) {
					v25 = v57
					v26 = v59
					v30 = v61
					continue
				} else {
					break
				}
				break
			}
			if v10&int32(1) == int32(0) {
				v89 = v57
			} else {
				v67 = v57
				v68 = v59
				v75 = v16 + v68<<(uint(int32(7))%32)
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
				if l0 != v76 {
					v89 = v67
				} else {
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+16)))
					if v78 != int32(1) {
						v89 = v67
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
						v89 = v67 + base.B2i32(base.Ui32(v81-int32(1)) < base.Ui32(int32(2)))
					}
				}
			}
		} else {
			v67 = v2
			v68 = v2
			v75 = v16 + v68<<(uint(int32(7))%32)
			v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+32))
			if l0 != v76 {
				v89 = v67
			} else {
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+16)))
				if v78 != int32(1) {
					v89 = v67
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
					v89 = v67 + base.B2i32(base.Ui32(v81-int32(1)) < base.Ui32(int32(2)))
				}
			}
		}
	}
	return v89
}
func F_logicalrep_worker_attach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[0]))
	v14 = F_LWLockAcquire(m, v10+int32(_a_F_logicalrep_worker_attach_0), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[1]))
		v21 = v18 + l0<<(uint(int32(7))%32)
		v23 = v21 + int32(16)
		*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[2])) = v23
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)))
		if v25 != 0 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
			if v26 != 0 {
				v68 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[0]))
				F_LWLockRelease(m, v68+int32(_a_F_logicalrep_worker_attach_0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(_a_F_logicalrep_worker_attach_1), v7)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_logicalrep_worker_attach_2), int32(792), int32(_a_F_logicalrep_worker_attach_3))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[3]))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v28
				F_before_shmem_exit(m, int32(1053), int64(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[0]))
					F_LWLockRelease(m, v35+int32(_a_F_logicalrep_worker_attach_0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_worker_attach[0]))
			F_LWLockRelease(m, v44+int32(_a_F_logicalrep_worker_attach_0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
						F_errmsg(m, int32(_a_F_logicalrep_worker_attach_4), v7+int32(16))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_logicalrep_worker_attach_2), int32(783), int32(_a_F_logicalrep_worker_attach_3))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
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
}
