package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BuildIndex_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
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
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
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
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
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
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 float64
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 float64
	_ = v415
	var v416 int32
	_ = v416
	var v429 float64
	_ = v429
	var v431 int32
	_ = v431
	var v432 float64
	_ = v432
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
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
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	F_InitBuildState_1(m, l3, l0, l1, l2, l4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[0]))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v57 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_1[1])))
	if v28&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = int32(_a_F_BuildIndex_1_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[2]))
	v36 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[2])) = v35 + v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v39 + v36
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(80))+232)) = int64(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v47 + v36
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[2])) = v53 - v36
	goto L4
L7:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+92)))
	if v448 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+56))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	v63 = F_plan_create_index_workers(m, v60, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v353 == int32(0) {
		goto L7
	} else {
		goto L91
	}
L10:
	;
	if v63 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+180))
	if v67 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v82 <= int32(0) {
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[3]))
	v82 = v79
	goto L12
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+108))
	if v70 == int32(-1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[3]))
	if v70 < v74 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v76 = v70
	goto L18
L17:
	;
	v76 = v74
	goto L18
L18:
	;
	v82 = v76
	goto L12
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+121)))
	v88 = F_palloc0(m, int32(20))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+72)) = v93 + int32(1)
	goto L21
L21:
	;
	v100 = F_CreateParallelContext(m, int32(_a_F_BuildIndex_1_1), int32(_a_F_BuildIndex_1_2), v82)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v86 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v104 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v108 = int32(_a_F_BuildIndex_1_3)
	goto L25
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v111 = F_table_parallelscan_estimate(m, v110, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v106 = F_RegisterSnapshot(m, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v108 = v106
	goto L25
L28:
	;
	v113 = F_add_size(m, int32(160), v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	v120 = F_add_size(m, v115, (v113+int32(31))&int32(-32))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+36)) = v120
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[5]))
	v126 = F_mul_size(m, v124, int32(1024))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	v130 = int32(_a_F_BuildIndex_1_4)
	if base.Ui32(v130) < base.Ui32(v126) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v134 = v126 - v130
	goto L34
L33:
	;
	v134 = v126
	goto L34
L34:
	;
	if base.Ui32(int32(2147483647)) <= base.Ui32(v134) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v137 = int32(2147483647)
	goto L37
L36:
	;
	v137 = v134
	goto L37
L37:
	;
	v142 = F_add_size(m, v128, (v137+int32(31))&int32(-32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+36)) = v142
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)+40))
	v147 = F_add_size(m, v145, int32(2))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+40)) = v147
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	if v151 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v100)+36))
	v153 = F_strlen(m, v151)
	mBase = m.M
	v158 = F_add_size(m, v152, v153&int32(-32)+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v170 = int32(1)
	goto L42
L42:
	;
	F_InitializeParallelDSM(m, v100)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+36)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v100)+40))
	v163 = F_add_size(m, v161, int32(1))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+40)) = v163
	v170 = v153 + int32(1)
	goto L42
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	if v173 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	switch v176 {
	case 0, 5:
		goto L50
	default:
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	v189 = F_shm_toc_allocate(m, v188, v113)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	F_DestroyParallelContext(m, v100)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	F_UnregisterSnapshot(m, v108)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+72)) = v184 - int32(1)
	goto L53
L53:
	;
	goto L9
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+8)) = uint8(v86)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v195
	v199 = v189 + int32(12)
	v200 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v199))), uint32(v200))
	*(*int64)(unsafe.Add(mBase, uint32(v199)+4)) = int64(-1)
	goto L55
L55:
	;
	v205 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v189)+24)), uint32(v205))
	*(*int64)(unsafe.Add(mBase, uint32(v189)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+28)) = v205
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v212, v189+int32(160), v108)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	v218 = F_shm_toc_allocate(m, v217, v137)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_HnswInitLockTranche(m)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+132)) = uint8(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+112)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v189)+108)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v189)+88)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v189)+44)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v189)+48)) = int64(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v189)+40)), uint32(v222))
	v237 = v189 + int32(56)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v237))) = uint16(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v237)+8)) = int64(-1)
	goto L59
L59:
	;
	v246 = v189 + int32(72)
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v246)+8)) = int64(-1)
	goto L60
L60:
	;
	v255 = v189 + int32(92)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v255))) = uint16(v257)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = int64(-1)
	goto L61
L61:
	;
	v264 = v189 + int32(116)
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v264))) = uint16(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = int64(-1)
	goto L62
L62:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	F_shm_toc_insert(m, v272, int64(-6917529027641081855), v189)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	F_shm_toc_insert(m, v276, int64(-6917529027641081854), v218)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	if v281 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	v283 = F_shm_toc_allocate(m, v282, v170)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_LaunchParallelWorkers(m, v100)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L73
	}
L68:
	;
	if v170 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	base.MemoryCopy(m, v283, v286, v170)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	F_shm_toc_insert(m, v288, int64(-6917529027641081853), v283)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v100
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v296 + int32(1)
	if v296 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_WaitForParallelWorkersToFinish(m, v100)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v323 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L83
	}
L77:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	switch v308 {
	case 0, 5:
		goto L79
	default:
		goto L78
	}
L78:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	F_DestroyParallelContext(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	F_UnregisterSnapshot(m, v307)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+72)) = v317 - int32(1)
	goto L82
L82:
	;
	goto L9
L83:
	;
	if v323 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v325
	F_errmsg(m, int32(_a_F_BuildIndex_1_5), v16)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+196)) = v88
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	F_HnswParallelScanAndInsert(m, v336, v337, v338, v339, int32(1))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_1_6), int32(1051), int32(_a_F_BuildIndex_1_7))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	F_WaitForParallelWorkersToAttach(m, v100)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L9
L91:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v356 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v432 = *(*float64)(unsafe.Add(mBase, uint32(v431)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v432
	goto L7
L93:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v361 = v357 + int32(24)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	goto L96
L94:
	;
	goto L95
L95:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v406 = int32(1)
	v407 = int32(0)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v353)+188))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)+140))
	v415 = m.T0[v414].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v353, v404, v405, v406, v407, v406, v407, int32(-1), int32(_a_F_BuildIndex_1_8), l3, v407)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L107
	}
L96:
	;
	v378 = base.AtomicRmwXchg32(m, v361, int32(0), int32(1))
	if v378 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+160)) = v357 + int32(40)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+204)) = v396
	v398 = *(*float64)(unsafe.Add(mBase, uint32(v357)+32))
	v399 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v357)+24)), uint32(v399))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L106
	}
L98:
	;
	F_s_lock(m, v361, int32(_a_F_BuildIndex_1_6), int32(769), int32(_a_F_BuildIndex_1_9))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v357)+28))
	if v362 != v384 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	v386 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v361))), uint32(v386))
	F_ConditionVariableSleep(m, v357+int32(12), int32(134217767))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	goto L97
L105:
	;
	goto L96
L106:
	;
	v429 = v398
	goto L92
L107:
	;
	v429 = v415
	goto L92
L108:
	;
	F_FlushPages(m, l3)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v453 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	F_WaitForParallelWorkersToFinish(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+118)))
	if v473 != int32(112) {
		goto L123
	} else {
		goto L124
	}
L115:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v453)+12))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	switch v458 {
	case 0, 5:
		goto L117
	default:
		goto L116
	}
L116:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	F_DestroyParallelContext(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	F_UnregisterSnapshot(m, v457)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+72)) = v467 - int32(1)
	goto L120
L120:
	;
	goto L114
L121:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	F_MemoryContextDelete(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L132
	}
L122:
	;
	v488 = F_RelationGetNumberOfBlocksInFork(m, l1, l4)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L130
	}
L123:
	;
	if l4 != int32(3) {
		goto L121
	} else {
		goto L129
	}
L124:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[8]))
	if int32(0) < v477 {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v480 != 0 {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	if l4 == int32(3) {
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v483 == int32(0) {
		goto L122
	} else {
		goto L128
	}
L128:
	;
	goto L121
L129:
	;
	goto L122
L130:
	;
	F_log_newpage_range(m, l1, l4, v488, int32(1))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L121
L132:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l3)+184))
	F_MemoryContextDelete(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	m.G0 = v16 + int32(16)
	return
}
func F_GetIndexAmRoutineByAmId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(2), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			if l1 != 0 {
				v109 = int32(0)
				m.G0 = v8 - int32(-64)
				return v109
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutineByAmId_0), v8)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(69), int32(_a_F_GetIndexAmRoutineByAmId_2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
			v33 = v31 + v32
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+72)))
			if v34 != int32(105) {
				if l1 != 0 {
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v109 = int32(0)
						m.G0 = v8 - int32(-64)
						return v109
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(_a_F_GetIndexAmRoutineByAmId_3)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v33 + int32(4)
							F_errmsg(m, int32(_a_F_GetIndexAmRoutineByAmId_4), v6+int32(-16))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(84), int32(_a_F_GetIndexAmRoutineByAmId_2))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
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
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
				if v59 == int32(0) {
					if l1 != 0 {
						F_ReleaseCatCache(m, v11)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v109 = int32(0)
							m.G0 = v8 - int32(-64)
							return v109
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v33 + int32(4)
								F_errmsg(m, int32(_a_F_GetIndexAmRoutineByAmId_5), v6+int32(-48))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(100), int32(_a_F_GetIndexAmRoutineByAmId_2))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
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
					F_ReleaseCatCache(m, v11)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v84 = F_OidFunctionCall0Coll(m, v59)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							if v84 != 0 {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
								if v86 == int32(438) {
									v109 = v84
									m.G0 = v8 - int32(-64)
									return v109
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
										F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutineByAmId_6), v6+int32(-32))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(43), int32(_a_F_GetIndexAmRoutineByAmId_7))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
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
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
									F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutineByAmId_6), v6+int32(-32))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(43), int32(_a_F_GetIndexAmRoutineByAmId_7))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int32(0)
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
		}
	}
}
func F_get_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v6)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+100)))
	v20 = F_build_index_paths(m, l0, l1, l2, l3, v16, int32(2), v12+int32(15))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	return
L3:
	;
	if v20 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = v6
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+109)))
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_add_path(m, l1, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+110)))
	if v44 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v58 = v35 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v58 < v59 {
		v35 = v58
		goto L6
	} else {
		goto L19
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+64))
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v40)+104))
	if base.F64_lt(v48, float64(1)) == int32(0) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v54 = F_lappend(m, v53, v40)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54
	goto L12
L19:
	;
	goto L7
L20:
	;
	v71 = int32(0)
	v74 = F_build_index_paths(m, l0, l1, l2, l3, v71, int32(1), v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	m.G0 = v12 + int32(16)
	return
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v77 = F_list_concat(m, v76, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v77
	goto L22
}
func F_index_can_return(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_index_can_return[0]))
	if v11 != v9 {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_index_can_return[1]))
		v15 = F_list_member_ptr(m, v14, v9)
		mBase = m.M
		v17 = v15
	} else {
		v17 = int32(1)
	}
	if v17 == int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
		if v21 != 0 {
			v22 = m.T0[v21].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = v22
				m.G0 = v7 + int32(16)
				return v27
			}
		} else {
			v27 = int32(0)
			m.G0 = v7 + int32(16)
			return v27
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v39 + int32(4)
				F_errmsg(m, int32(_a_F_index_can_return_0), v7)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_can_return_1), int32(837), int32(_a_F_index_can_return_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
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
func F_index_check_primary_key(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L12
	} else {
		goto L57
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L53
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L49
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L46
	}
L6:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
	if v94 != 0 {
		goto L4
	} else {
		goto L31
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+131)))
	if v15 != int32(1) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v18 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	F_list_free(m, v18)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L30
	}
L12:
	;
	return
L13:
	;
	if v18 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v22 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v23 <= v22 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v28 = v22
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v28<<(uint(int32(2))%32))))
	v39 = F_SearchSysCache1(m, int32(34), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	F_list_free(m, v18)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L25
	}
L18:
	;
	if v39 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v44)+14)))
	F_ReleaseCatCache(m, v39)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v46 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v52 = v28 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v53 <= v52 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L17
L24:
	;
	v28 = v52
	goto L16
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v64 + int32(4)
	F_errmsg(m, int32(_a_F_index_check_primary_key_0), v8+int32(-16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(221), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L12
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
	goto L6
L31:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(0) < v95 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v103 = int32(0)
	v106 = v95
	goto L35
L33:
	;
	goto L34
L34:
	;
	m.G0 = v10 - int32(-64)
	return
L35:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1+int32(12)+v103<<(uint(int32(1))%32)))))
	if v111 == int32(0) {
		goto L3
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	if int32(0) <= v111 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v118 = F_SearchSysCache2(m, int32(7), v117, v111)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	v132 = v106
	goto L40
L40:
	;
	v134 = v103 + int32(1)
	if v134 < v132 {
		v103 = v134
		v106 = v132
		goto L35
	} else {
		goto L45
	}
L41:
	;
	if v118 == int32(0) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+86)))
	if v125 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_ReleaseCatCache(m, v118)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v132 = v130
	goto L40
L45:
	;
	goto L36
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v38
	F_errmsg_internal(m, int32(_a_F_index_check_primary_key_3), v8+int32(-32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(168), int32(_a_F_index_check_primary_key_4))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_index_check_primary_key_5), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(234), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_index_check_primary_key_6), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(251), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
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
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v111
	F_errmsg_internal(m, int32(_a_F_index_check_primary_key_7), v10)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(262), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v124 + int32(4)
	F_errmsg(m, int32(_a_F_index_check_primary_key_8), v8+int32(-48))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(269), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_compute_xid_horizon_for_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l2 < int32(0) {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_index_compute_xid_horizon_for_tuples[0]))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+(l2^int32(-1))<<(uint(int32(2))%32))))
		v36 = v28
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_index_compute_xid_horizon_for_tuples[1]))
		v36 = v30 + l2<<(uint(int32(13))%32) + int32(-8192)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l0
	if l2 < int32(0) {
		v41 = *(*int32)(unsafe.Add(mBase, _c_F_index_compute_xid_horizon_for_tuples[2]))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+(l2^int32(-1))<<(uint(int32(6))%32))+16))
		v56 = v47
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, _c_F_index_compute_xid_horizon_for_tuples[3]))
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+l2<<(uint(int32(6))%32)+int32(-64))+16))
		v56 = v55
	}
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = int64(0)
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v56
	v65 = F_palloc(m, l4<<(uint(int32(3))%32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v65
		v72 = F_palloc(m, l4*int32(6))
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v72
			if int32(0) < l4 {
				v80 = v59
				v82 = int32(0)
				for {
					v94 = int32(1)
					v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v82<<(uint(v94)%32)))))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(20)+v97<<(uint(int32(2))%32))))
					v104 = v36 + v101&int32(_a_F_index_compute_xid_horizon_for_tuples_0)
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
					v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+4)))
					v109 = v65 + v82<<(uint(int32(3))%32)
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+6)) = uint16(v80)
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+4)) = uint16(v106)
					*(*int32)(unsafe.Add(mBase, uint32(v109))) = v105
					v115 = v72 + v82*int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v115)+2)) = v94
					*(*uint16)(unsafe.Add(mBase, uint32(v115))) = uint16(v97)
					v120 = v80 + v94
					*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v120
					v123 = v82 + v94
					if v123 != l4 {
						v80 = v120
						v82 = v123
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
			v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
			v143 = m.T0[v142].(func(*base.Module, int32, int32) int32)(m, l1, v17+int32(4))
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return int32(0)
			} else {
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				F_pfree(m, v145)
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return int32(0)
				} else {
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
					F_pfree(m, v148)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int32(0)
					} else {
						m.G0 = v17 + int32(32)
						return v143
					}
				}
			}
		}
	}
}
func F_index_fetch_heap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v3)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_index_fetch_heap[0]))
	if v13 != 0 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_fetch_heap[1])))
		if v15&int32(1) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_index_fetch_heap_0), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_fetch_heap_1), int32(1218), int32(_a_F_index_fetch_heap_2))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+188))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
			v31 = m.T0[v30].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20, l0+int32(60), v23, l1, l0+int32(66), v8+int32(15))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				if v31 == int32(0) {
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v55 == int32(0) {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
					} else {
					}
					m.G0 = v8 + int32(16)
					return v31
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
					if v38 == int32(0) {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
						if v41 != int32(1) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v55 == int32(0) {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
							} else {
							}
							m.G0 = v8 + int32(16)
							return v31
						} else {
							F_pgstat_assoc_relation(m, v37)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
								v48 = v47
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
								if v55 == int32(0) {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
								} else {
								}
								m.G0 = v8 + int32(16)
								return v31
							}
						}
					} else {
						v48 = v38
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v55 == int32(0) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
						} else {
						}
						m.G0 = v8 + int32(16)
						return v31
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+188))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
		v31 = m.T0[v30].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v20, l0+int32(60), v23, l1, l0+int32(66), v8+int32(15))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			if v31 == int32(0) {
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v55 == int32(0) {
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
				} else {
				}
				m.G0 = v8 + int32(16)
				return v31
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+272))
				if v38 == int32(0) {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+268)))
					if v41 != int32(1) {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v55 == int32(0) {
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
						} else {
						}
						m.G0 = v8 + int32(16)
						return v31
					} else {
						F_pgstat_assoc_relation(m, v37)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
							v48 = v47
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v55 == int32(0) {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
							} else {
							}
							m.G0 = v8 + int32(16)
							return v31
						}
					}
				} else {
					v48 = v38
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v48)+32)) = v49 + int64(1)
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v55 == int32(0) {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v58)
					} else {
					}
					m.G0 = v8 + int32(16)
					return v31
				}
			}
		}
	}
}
func F_index_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v13 {
		v16 = int32(4)
		v20 = l2 + l1<<(uint(v16)%32) + v16
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		if v21 < int32(0) {
			v64 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = v64
				m.G0 = v9 + int32(16)
				return v70
			}
		} else {
			v26 = l0 + v21 + int32(8)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
			if v27 != int32(1) {
				v70 = v26
				m.G0 = v9 + int32(16)
				return v70
			} else {
				v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
				switch v30&int32(_a_F_index_getattr_1_0) - int32(1) {
				case 0:
					v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26))))
					v70 = v35
					m.G0 = v9 + int32(16)
					return v70
				case 1:
					v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26))))
					v70 = v36
					m.G0 = v9 + int32(16)
					return v70
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v30
						F_errmsg_internal(m, int32(_a_F_index_getattr_1_1), v9)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_index_getattr_1_2), int32(70), int32(_a_F_index_getattr_1_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v70 = v37
					m.G0 = v9 + int32(16)
					return v70
				}
			}
		}
	} else {
		v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		v54 = int32(1)
		if int32(base.Ui32(v53)>>(uint(l1-v54)%32))&v54 != 0 {
			v64 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = v64
				m.G0 = v9 + int32(16)
				return v70
			}
		} else {
			v59 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v59)
			v70 = int32(0)
			m.G0 = v9 + int32(16)
			return v70
		}
	}
}
func F_index_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
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
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_initialize[0]))
	if v17 != v15 {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_initialize[1]))
		v21 = F_list_member_ptr(m, v20, v15)
		mBase = m.M
		v23 = v21
	} else {
		v23 = int32(1)
	}
	if v23 == int32(0) {
		v27 = F_EstimateSnapshotSpace(m, l2)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = F_add_size(m, int32(32), v27)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l7)+8)) = v31
				v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(l7))) = v33
				v35 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(l7)+12)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l7)+20)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(l7)+24)) = int64(0)
				v42 = l7 + int32(32)
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v52 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v42)+16)) = uint8(v54)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v52
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v51
				*(*uint8)(unsafe.Add(mBase, uint32(v42)+17)) = uint8(v50)
				if v50&int32(1) != 0 {
					v63 = v49
				} else {
					v63 = int32(0)
				}
				if v54 != 0 {
					v64 = v63
				} else {
					v64 = v49
				}
				*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v66 == int32(0) {
				} else {
					v70 = v66 << (uint(int32(2)) % 32)
					if v70 == int32(0) {
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						base.MemoryCopy(m, l7+int32(56), v75, v70)
					}
				}
				if v64 <= int32(0) {
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
					v82 = v80 << (uint(int32(2)) % 32)
					if v82 == int32(0) {
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						base.MemoryCopy(m, v42+v85<<(uint(int32(2))%32)+int32(24), v91, v82)
					}
				}
				v97 = (v29 + int32(7)) & int32(-8)
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l7)+24)) = v97
					v102 = l5<<(uint(int32(3))%32) + int32(8)
					v103 = F_add_size(m, v97, v102)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l7)+24))
						v106 = l7 + v105
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v106
						if v102 != 0 {
							base.MemoryFill(m, v106, int32(0), v102)
						} else {
						}
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
						*(*int32)(unsafe.Add(mBase, uint32(v110))) = l5
						v117 = (v103 + int32(7)) & int32(-8)
						if l4 == int32(0) {
							m.G0 = v13 + int32(16)
							return
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+124))
							if v122 == int32(0) {
								m.G0 = v13 + int32(16)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l7)+28)) = v117
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+124))
								m.T0[v128].(func(*base.Module, int32))(m, v117+l7)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return
								} else {
									m.G0 = v13 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v117 = v97
					if l4 == int32(0) {
						m.G0 = v13 + int32(16)
						return
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+124))
						if v122 == int32(0) {
							m.G0 = v13 + int32(16)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l7)+28)) = v117
							v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+124))
							m.T0[v128].(func(*base.Module, int32))(m, v117+l7)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								m.G0 = v13 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v137 = m.ExcPending
		if v137 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return
			} else {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v141 + int32(4)
				F_errmsg(m, int32(_a_F_index_parallelscan_initialize_0), v13)
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_index_parallelscan_initialize_1), int32(520), int32(_a_F_index_parallelscan_initialize_2))
					mBase = m.M
					v152 = m.ExcPending
					if v152 != 0 {
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
