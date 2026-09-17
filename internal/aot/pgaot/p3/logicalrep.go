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
	var v110 int32
	_ = v110
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
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
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
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
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int64
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = int64(309237645316)
	v46 = F_hash_create(m, int32(_a_F_logicalrep_rel_open_3), int32(128), v15+int32(80), int32(1064))
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
	F_CacheRegisterRelcacheCallback(m, int32(1015))
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
	v446 = m.ExcPending
	if v446 != 0 {
		goto L7
	} else {
		goto L110
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L7
	} else {
		goto L106
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L103
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
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L100
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
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+56)))
	if v366 != int32(114) {
		goto L96
	} else {
		goto L97
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
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	F_CheckSubscriptionRelkind(m, v106, v107, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v111 = int32(_a_F_logicalrep_rel_open_4)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[3]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+52))
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[3])) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v120 = F_make_attrmap(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[3])) = v112
	v125 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v130 = F_bms_add_range(m, v125, v125, v127-int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if int32(0) < v132 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	F_bms_free(m, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L92
	}
L41:
	;
	v136 = v132
	v138 = v130
	v139 = v3
	v140 = v3
	goto L44
L42:
	;
	goto L43
L43:
	;
	if v130 != 0 {
		v434 = v130
		goto L12
	} else {
		goto L91
	}
L44:
	;
	v152 = v114 + v136<<(uint(int32(4))%32) + v139*int32(100)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+111)))
	if v153 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v253 != 0 {
		v434 = v253
		goto L12
	} else {
		goto L71
	}
L46:
	;
	v263 = v139 + int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v263 < v264 {
		v136 = v264
		v138 = v253
		v139 = v263
		v140 = v255
		goto L44
	} else {
		goto L70
	}
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v161 = int32(_a_F_logicalrep_rel_open_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v157+v139<<(uint(int32(1))%32)))) = uint16(v161)
	v253 = v138
	v255 = v140
	goto L46
L48:
	;
	goto L49
L49:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if int32(0) < v163 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	*(*uint16)(unsafe.Add(mBase, uint32(v239+v139<<(uint(int32(1))%32)))) = uint16(v173)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+int32(20))+90)))
	if v244 != 0 {
		goto L65
	} else {
		goto L66
	}
L51:
	;
	v169 = v152 + int32(24)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v173 = int32(0)
	goto L54
L52:
	;
	goto L53
L53:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v236 = int32(_a_F_logicalrep_rel_open_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v232+v139<<(uint(int32(1))%32)))) = uint16(v236)
	v253 = v138
	v255 = v140
	goto L46
L54:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v170+v173<<(uint(int32(2))%32))))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if base.B2i32(v190 == int32(0))|base.B2i32(v190 != v193) != 0 {
		v211 = v190
		v212 = v193
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L53
L56:
	;
	if v211-v212 == int32(0) {
		goto L50
	} else {
		goto L63
	}
L57:
	;
	goto L56
L58:
	;
	v196 = v187
	v197 = v169
	goto L59
L59:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	if v201 == int32(0) {
		v211 = v201
		v212 = v200
		goto L57
	} else {
		goto L61
	}
L60:
	;
	v211 = v201
	v212 = v200
	goto L57
L61:
	;
	v204 = int32(1)
	if v201 == v200 {
		v196 = v196 + v204
		v197 = v197 + v204
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v217 = v173 + int32(1)
	if v217 != v163 {
		v173 = v217
		goto L54
	} else {
		goto L64
	}
L64:
	;
	goto L55
L65:
	;
	v245 = F_bms_add_member(m, v140, v173)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L68
	}
L66:
	;
	v247 = v140
	goto L67
L67:
	;
	v248 = F_bms_del_member(m, v138, v173)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L69
	}
L68:
	;
	v247 = v245
	goto L67
L69:
	;
	v253 = v248
	v255 = v247
	goto L46
L70:
	;
	goto L45
L71:
	;
	if v255 == int32(0) {
		goto L40
	} else {
		goto L72
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v275 = int32(0)
	if v255 == v275 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v61)+4))
	v312 = F_logicalrep_get_attrs_str(m, v61, v255)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L88
	}
L76:
	;
	v310 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v282 = int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v283 <= v282 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v286 = v282
	goto L81
L80:
	;
	v286 = v283
	goto L81
L81:
	;
	v290 = int32(0)
	v292 = v275
	goto L82
L82:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v255+int32(8)+v290<<(uint(int32(2))%32))))
	if v298 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v310 = v301
	goto L75
L84:
	;
	v301 = v292 + base.I32_popcnt(v298)
	goto L86
L85:
	;
	v301 = v292
	goto L86
L86:
	;
	v303 = v290 + int32(1)
	if v303 != v286 {
		v290 = v303
		v292 = v301
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v312
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v311
	F_errmsg_plural(m, int32(_a_F_logicalrep_rel_open_6), int32(_a_F_logicalrep_rel_open_7), v310, v15+int32(16))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(279), int32(_a_F_logicalrep_rel_open_9))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
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
	goto L40
L92:
	;
	F_bms_free(m, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	F_logicalrep_rel_mark_updatable(m, v61)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
	v349 = F_FindLogicalRepLocalIndex(m, v347, v61, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	v351 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)) = uint8(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+52)) = v349
	goto L19
L96:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_logicalrep_rel_open[4]))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
	v375 = F_GetSubscriptionRelState(m, v371, v372, v61-int32(-64))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L7
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	m.G0 = v15 + int32(128)
	return v61
L99:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+56)) = uint8(v375)
	goto L98
L100:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v386
	F_errmsg_internal(m, int32(_a_F_logicalrep_rel_open_10), v15-int32(-64))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(364), int32(_a_F_logicalrep_rel_open_11))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v402
	F_errmsg_internal(m, int32(_a_F_logicalrep_rel_open_12), v15+int32(48))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(370), int32(_a_F_logicalrep_rel_open_11))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v421
	F_errmsg(m, int32(_a_F_logicalrep_rel_open_13), v15)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(422), int32(_a_F_logicalrep_rel_open_11))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	v450 = int32(0)
	if v434 == v450 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v61)+4))
	v487 = F_logicalrep_get_attrs_str(m, v61, v434)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L7
	} else {
		goto L125
	}
L113:
	;
	v485 = int32(0)
	goto L112
L114:
	;
	goto L115
L115:
	;
	v457 = int32(1)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	if v458 <= v457 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v461 = v457
	goto L118
L117:
	;
	v461 = v458
	goto L118
L118:
	;
	v465 = int32(0)
	v467 = v450
	goto L119
L119:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v434+int32(8)+v465<<(uint(int32(2))%32))))
	if v473 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v485 = v476
	goto L112
L121:
	;
	v476 = v467 + base.I32_popcnt(v473)
	goto L123
L122:
	;
	v476 = v467
	goto L123
L123:
	;
	v478 = v465 + int32(1)
	if v478 != v461 {
		v465 = v478
		v467 = v476
		goto L119
	} else {
		goto L124
	}
L124:
	;
	goto L120
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v486
	F_errmsg_plural(m, int32(_a_F_logicalrep_rel_open_14), int32(_a_F_logicalrep_rel_open_15), v485, v15+int32(32))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_logicalrep_rel_open_8), int32(268), int32(_a_F_logicalrep_rel_open_9))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
		v21 = v18 + l0*int32(112)
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
								F_errfinish(m, int32(_a_F_logicalrep_worker_attach_2), int32(740), int32(_a_F_logicalrep_worker_attach_3))
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
				F_before_shmem_exit(m, int32(988), int32(0))
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
							F_errfinish(m, int32(_a_F_logicalrep_worker_attach_2), int32(731), int32(_a_F_logicalrep_worker_attach_3))
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
