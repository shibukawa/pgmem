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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
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
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
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
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L58
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L55
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L52
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
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L49
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
	v152 = F_strlen(m, v141)
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
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v145)
	goto L18
L23:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L45
L24:
	;
	if v116 == int32(0) {
		v141 = v114
		v142 = v115
		goto L22
	} else {
		goto L44
	}
L25:
	;
	v114 = v38
	v115 = v37
	v116 = v46
	goto L24
L26:
	;
	goto L27
L27:
	;
	if v38&int32(3) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v83 == int32(0) {
		v141 = v80
		v142 = v81
		goto L22
	} else {
		goto L37
	}
L29:
	;
	v80 = v38
	v81 = v37
	v82 = v46
	v83 = int32(1)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v59 = v38
	v60 = v37
	v61 = v46
	goto L32
L32:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v63)
	if v63 == int32(0) {
		v121 = v59
		v122 = v60
		v123 = v61
		goto L23
	} else {
		goto L34
	}
L33:
	;
	v80 = v74
	v81 = v68
	v82 = v70
	v83 = v72
	goto L28
L34:
	;
	v67 = int32(1)
	v68 = v60 + v67
	v70 = v61 - v67
	v71 = int32(0)
	v72 = base.B2i32(v70 != v71)
	v74 = v59 + v67
	if v74&int32(3) == v71 {
		v80 = v74
		v81 = v68
		v82 = v70
		v83 = v72
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v70 != 0 {
		v59 = v74
		v60 = v68
		v61 = v70
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v86 == int32(0) {
		v114 = v80
		v115 = v81
		v116 = v82
		goto L24
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v82) < base.Ui32(int32(4)) {
		v114 = v80
		v115 = v81
		v116 = v82
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v92 = v80
	v93 = v81
	v94 = v82
	goto L40
L40:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v100 = int32(-2139062144)
	if (int32(16843008)-v97|v97)&v100 != v100 {
		v121 = v92
		v122 = v93
		v123 = v94
		goto L23
	} else {
		goto L42
	}
L41:
	;
	v114 = v108
	v115 = v106
	v116 = v110
	goto L24
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v97
	v105 = int32(4)
	v106 = v93 + v105
	v108 = v92 + v105
	v110 = v94 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v92 = v108
		v93 = v106
		v94 = v110
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L23
L45:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v130)
	if v130 == int32(0) {
		v141 = v126
		v142 = v127
		goto L22
	} else {
		goto L47
	}
L46:
	;
	v141 = v137
	v142 = v135
	goto L22
L47:
	;
	v134 = int32(1)
	v135 = v127 + v134
	v137 = v126 + v134
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v137
		v127 = v135
		v128 = v139
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v14
	F_errmsg_internal(m, int32(400030), v7+int32(-16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(490650), int32(206), int32(243321))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
	F_errmsg_internal(m, int32(400066), v9)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(490650), int32(211), int32(243321))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	F_errmsg_internal(m, int32(400103), v7+int32(-48))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(490650), int32(214), int32(243321))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l1
	F_errmsg_internal(m, int32(400136), v7+int32(-32))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(490650), int32(218), int32(243321))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
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
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
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
	var v363 int32
	_ = v363
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int64
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l0
	v20 = *(*int32)(unsafe.Add(mBase, _consts[665]))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v56 = v20
	goto L3
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[666]))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v62 = F_hash_search(m, v56, v16+int32(76), int32(0), v16+int32(80))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L11
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v32 = F_AllocSetContextCreateInternal(m, v27, int32(61764), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v37 = v22
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = int64(309237645316)
	v47 = F_hash_create(m, int32(395055), int32(128), v16+int32(80), int32(1064))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[666])) = v32
	v37 = v32
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[665])) = v47
	F_CacheRegisterRelcacheCallback(m, int32(1015))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[665]))
	v56 = v54
	goto L3
L11:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+80)))
	if v64 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L7
	} else {
		goto L113
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L109
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
	} else {
		goto L106
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
	if v65 != 0 {
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
	v398 = m.ExcPending
	if v398 != 0 {
		goto L7
	} else {
		goto L103
	}
L18:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+32)))
	if v66 != int32(1) {
		v87 = v66
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+56)))
	if v379 != int32(114) {
		goto L99
	} else {
		goto L100
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	if v91 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	if v87 != 0 {
		goto L19
	} else {
		goto L30
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v62)+36))
	v70 = F_try_table_open(m, v69, l1)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+40)) = v70
	if v70 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+32)) = uint8(v75)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+32)))
	if v77 != 0 {
		v87 = v77
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_sequence_close(m, v70, l1)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+40)) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+32)))
	if v82&int32(1) == v80 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	goto L20
L31:
	;
	F_free_attrmap(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v99 = F_makeRangeVar(m, v96, v97, int32(-1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = int32(0)
	goto L33
L35:
	;
	v102 = int32(0)
	v104 = F_RangeVarGetRelidExtended(m, v99, l1, int32(1), v102, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	if v104 == int32(0) {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v109 = F_table_open(m, v104, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+36)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v62)+40)) = v109
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+48))
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113)+119)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	F_CheckSubscriptionRelkind(m, v114, v115, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v119 = int32(4476144)
	v120 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+52))
	v125 = *(*int32)(unsafe.Add(mBase, _consts[666]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v128 = F_make_attrmap(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+44)) = v128
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v120
	v133 = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v138 = F_bms_add_range(m, v133, v133, v135-int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if int32(0) < v140 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	F_bms_free(m, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L95
	}
L43:
	;
	v145 = v140
	v149 = v138
	v150 = v3
	v151 = v3
	goto L46
L44:
	;
	goto L45
L45:
	;
	if v138 != 0 {
		v448 = v138
		goto L12
	} else {
		goto L94
	}
L46:
	;
	v163 = v122 + int32(20) + v145<<(uint(int32(4))%32) + v150*int32(100)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+91)))
	if v164 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if v264 != 0 {
		v448 = v264
		goto L12
	} else {
		goto L74
	}
L48:
	;
	v274 = v150 + int32(1)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v274 < v275 {
		v145 = v275
		v149 = v264
		v150 = v274
		v151 = v266
		goto L46
	} else {
		goto L73
	}
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v172 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v168+v150<<(uint(int32(1))%32)))) = uint16(v172)
	v264 = v149
	v266 = v151
	goto L48
L50:
	;
	goto L51
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if int32(0) < v174 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*uint16)(unsafe.Add(mBase, uint32(v249+v150<<(uint(int32(1))%32)))) = uint16(v181)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+90)))
	if v254 != 0 {
		goto L68
	} else {
		goto L69
	}
L53:
	;
	v178 = v163 + int32(4)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v181 = int32(0)
	goto L56
L54:
	;
	goto L55
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v246 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v242+v150<<(uint(int32(1))%32)))) = uint16(v246)
	v264 = v149
	v266 = v151
	goto L48
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v179+v181<<(uint(int32(2))%32))))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v201 == int32(0) {
		v220 = v200
		v221 = v201
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L55
L58:
	;
	if v221-v220 == int32(0) {
		goto L52
	} else {
		goto L66
	}
L59:
	;
	goto L58
L60:
	;
	if v200 != v201 {
		v220 = v200
		v221 = v201
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v205 = v197
	v206 = v178
	goto L62
L62:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v210 == int32(0) {
		v220 = v209
		v221 = v210
		goto L59
	} else {
		goto L64
	}
L63:
	;
	v220 = v209
	v221 = v210
	goto L59
L64:
	;
	v213 = int32(1)
	if v209 == v210 {
		v205 = v205 + v213
		v206 = v206 + v213
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v226 = v181 + int32(1)
	if v226 != v174 {
		v181 = v226
		goto L56
	} else {
		goto L67
	}
L67:
	;
	goto L57
L68:
	;
	v255 = F_bms_add_member(m, v151, v181)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L71
	}
L69:
	;
	v257 = v151
	goto L70
L70:
	;
	v258 = F_bms_del_member(m, v149, v181)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L7
	} else {
		goto L72
	}
L71:
	;
	v257 = v255
	goto L70
L72:
	;
	v264 = v258
	v266 = v257
	goto L48
L73:
	;
	goto L47
L74:
	;
	if v266 == int32(0) {
		goto L42
	} else {
		goto L75
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	v286 = int32(0)
	if v266 == v286 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v62)+4))
	v323 = F_logicalrep_get_attrs_str(m, v62, v266)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L7
	} else {
		goto L91
	}
L79:
	;
	v321 = int32(0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v293 = int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v294 <= v293 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v297 = v293
	goto L84
L83:
	;
	v297 = v294
	goto L84
L84:
	;
	v301 = int32(0)
	v303 = v286
	goto L85
L85:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v266+int32(8)+v301<<(uint(int32(2))%32))))
	if v309 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v321 = v312
	goto L78
L87:
	;
	v312 = v303 + base.I32_popcnt(v309)
	goto L89
L88:
	;
	v312 = v303
	goto L89
L89:
	;
	v314 = v301 + int32(1)
	if v314 != v297 {
		v301 = v314
		v303 = v312
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v322
	F_errmsg_plural(m, int32(200419), int32(198433), v321, v16+int32(16))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(490900), int32(279), int32(129659))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	goto L42
L95:
	;
	F_bms_free(m, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	F_logicalrep_rel_mark_updatable(m, v62)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v62)+40))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v361 = F_FindLogicalRepLocalIndex(m, v359, v62, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+32)) = uint8(v363)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+52)) = v361
	goto L19
L99:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _consts[667]))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v62)+36))
	v388 = F_GetSubscriptionRelState(m, v384, v385, v62-int32(-64))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L7
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	m.G0 = v16 + int32(128)
	return v62
L102:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+56)) = uint8(v388)
	goto L101
L103:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v399
	F_errmsg_internal(m, int32(56480), v16-int32(-64))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(490900), int32(364), int32(278626))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
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
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v415
	F_errmsg_internal(m, int32(278691), v16+int32(48))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(490900), int32(370), int32(278626))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v62)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v434
	F_errmsg(m, int32(69815), v16)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(490900), int32(422), int32(278626))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L7
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
	F_errcode(m, int32(325))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	v464 = int32(0)
	if v448 == v464 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v62)+4))
	v501 = F_logicalrep_get_attrs_str(m, v62, v448)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L7
	} else {
		goto L128
	}
L116:
	;
	v499 = int32(0)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v471 = int32(1)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v472 <= v471 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v475 = v471
	goto L121
L120:
	;
	v475 = v472
	goto L121
L121:
	;
	v479 = int32(0)
	v481 = v464
	goto L122
L122:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v448+int32(8)+v479<<(uint(int32(2))%32))))
	if v487 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v499 = v490
	goto L115
L124:
	;
	v490 = v481 + base.I32_popcnt(v487)
	goto L126
L125:
	;
	v490 = v481
	goto L126
L126:
	;
	v492 = v479 + int32(1)
	if v492 != v475 {
		v479 = v492
		v481 = v490
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v501
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v500
	F_errmsg_plural(m, int32(200501), int32(198516), v499, v16+int32(32))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(490900), int32(268), int32(129659))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v14 = F_LWLockAcquire(m, v10+int32(5504), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[661]))
		v21 = v18 + l0*int32(112)
		v23 = v21 + int32(16)
		*(*int32)(unsafe.Add(mBase, _consts[580])) = v23
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)))
		if v25 != 0 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
			if v26 != 0 {
				v68 = *(*int32)(unsafe.Add(mBase, _consts[44]))
				F_LWLockRelease(m, v68+int32(5504))
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
							F_errmsg(m, int32(322744), v7)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(490041), int32(740), int32(322638))
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
				v28 = *(*int32)(unsafe.Add(mBase, _consts[128]))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v28
				F_before_shmem_exit(m, int32(988), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					F_LWLockRelease(m, v35+int32(5504))
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
			v44 = *(*int32)(unsafe.Add(mBase, _consts[44]))
			F_LWLockRelease(m, v44+int32(5504))
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
						F_errmsg(m, int32(322685), v7+int32(16))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							F_errfinish(m, int32(490041), int32(731), int32(322638))
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
