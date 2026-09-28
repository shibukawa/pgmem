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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
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
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 float64
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 float64
	_ = v413
	var v414 int32
	_ = v414
	var v427 float64
	_ = v427
	var v429 int32
	_ = v429
	var v430 float64
	_ = v430
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v65 == int32(0) {
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
	v43 = int32(0)
	v45 = int32(_a_F_BuildIndex_1_1)
	v46 = base.AtomicRmwOr32(m, v43, v45, v43)
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(80))+232)) = int64(2)
	v54 = base.AtomicRmwOr32(m, v43, v45, v43)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v55 + v36
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[2])) = v61 - v36
	goto L4
L7:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+92)))
	if v446 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L8:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+56))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+56))
	v71 = F_plan_create_index_workers(m, v68, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
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
		goto L93
	}
L10:
	;
	if v71 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+180))
	if v75 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v90 <= int32(0) {
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[3]))
	v90 = v87
	goto L12
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+116))
	if v78 == int32(-1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[3]))
	if v78 < v82 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v84 = v78
	goto L18
L17:
	;
	v84 = v82
	goto L18
L18:
	;
	v90 = v84
	goto L12
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+121)))
	v96 = F_palloc0(m, int32(20))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+72)) = v101 + int32(1)
	goto L21
L21:
	;
	v108 = F_CreateParallelContext(m, int32(_a_F_BuildIndex_1_2), int32(_a_F_BuildIndex_1_3), v90)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v94 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v112 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v116 = int32(_a_F_BuildIndex_1_4)
	goto L25
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v119 = F_table_parallelscan_estimate(m, v118, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v114 = F_RegisterSnapshot(m, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v116 = v114
	goto L25
L28:
	;
	v121 = F_add_size(m, int32(160), v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108)+36))
	v128 = F_add_size(m, v123, (v121+int32(31))&int32(-32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+36)) = v128
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[5]))
	v134 = F_mul_size(m, v132, int32(1024))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v108)+36))
	v138 = int32(_a_F_BuildIndex_1_5)
	if base.Ui32(v138) < base.Ui32(v134) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v142 = v134 - v138
	goto L34
L33:
	;
	v142 = v134
	goto L34
L34:
	;
	if base.Ui32(int32(2147483647)) <= base.Ui32(v142) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v145 = int32(2147483647)
	goto L37
L36:
	;
	v145 = v142
	goto L37
L37:
	;
	v150 = F_add_size(m, v136, (v145+int32(31))&int32(-32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+36)) = v150
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
	v155 = F_add_size(m, v153, int32(2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+40)) = v155
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	if v159 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v108)+36))
	v161 = F_strlen(m, v159)
	mBase = m.M
	v166 = F_add_size(m, v160, v161&int32(-32)+int32(32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v178 = int32(1)
	goto L42
L42:
	;
	F_InitializeParallelDSM(m, v108)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+36)) = v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
	v171 = F_add_size(m, v169, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+40)) = v171
	v178 = v161 + int32(1)
	goto L42
L45:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v108)+44))
	if v181 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v184 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v199 = F_shm_toc_allocate(m, v198, v121)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L55
	}
L49:
	;
	F_UnregisterSnapshot(m, v116)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_DestroyParallelContext(m, v108)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+72)) = v194 - int32(1)
	goto L54
L54:
	;
	goto L9
L55:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+8)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v205
	v209 = v199 + int32(12)
	v210 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v209))), uint32(v210))
	*(*int64)(unsafe.Add(mBase, uint32(v209)+4)) = int64(-1)
	goto L56
L56:
	;
	v215 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v199)+24)), uint32(v215))
	*(*int64)(unsafe.Add(mBase, uint32(v199)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+28)) = v215
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v222, v199+int32(160), v116)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v228 = F_shm_toc_allocate(m, v227, v145)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_HnswInitLockTranche(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+132)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+112)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v199)+108)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v199)+88)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v199)+44)) = v232
	*(*int64)(unsafe.Add(mBase, uint32(v199)+48)) = int64(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v199)+40)), uint32(v232))
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	F_LWLockInitialize(m, v199+int32(56), v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	F_LWLockInitialize(m, v199+int32(72), v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	F_LWLockInitialize(m, v199+int32(92), v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	F_LWLockInitialize(m, v199+int32(116), v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	F_shm_toc_insert(m, v270, int64(-6917529027641081855), v199)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	F_shm_toc_insert(m, v274, int64(-6917529027641081854), v228)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	if v279 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v281 = F_shm_toc_allocate(m, v280, v178)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_LaunchParallelWorkers(m, v108)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L74
	}
L69:
	;
	if v178 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	base.MemoryCopy(m, v281, v284, v178)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	F_shm_toc_insert(m, v286, int64(-6917529027641081853), v281)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v108
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v294 + int32(1)
	if v294 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_WaitForParallelWorkersToFinish(m, v108)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v323 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L85
	}
L78:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v306 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_UnregisterSnapshot(m, v305)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	F_DestroyParallelContext(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v316)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+72)) = v317 - int32(1)
	goto L84
L84:
	;
	goto L9
L85:
	;
	if v323 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v325
	F_errmsg(m, int32(_a_F_BuildIndex_1_6), v16)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+196)) = v96
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	F_HnswParallelScanAndInsert(m, v336, v337, v338, v339, int32(1))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_1_7), int32(1051), int32(_a_F_BuildIndex_1_8))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	F_WaitForParallelWorkersToAttach(m, v108)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L9
L93:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v356 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v427
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v430 = *(*float64)(unsafe.Add(mBase, uint32(v429)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v430
	goto L7
L95:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+8))
	v361 = v357 + int32(24)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	goto L98
L96:
	;
	goto L97
L97:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v404 = int32(1)
	v405 = int32(0)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v353)+188))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+140))
	v413 = m.T0[v412].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v353, v402, v403, v404, v405, v404, v405, int32(-1), int32(_a_F_BuildIndex_1_9), l3, v405)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L109
	}
L98:
	;
	v378 = base.AtomicRmwXchg32(m, v361, int32(0), int32(1))
	if v378 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+160)) = v357 + int32(40)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+204)) = v394
	v396 = *(*float64)(unsafe.Add(mBase, uint32(v357)+32))
	v397 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v357)+24)), uint32(v397))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L108
	}
L100:
	;
	F_s_lock(m, v361, int32(_a_F_BuildIndex_1_10))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v357)+28))
	if v362 != v382 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	v384 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v361))), uint32(v384))
	F_ConditionVariableSleep(m, v357+int32(12), int32(134217767))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	goto L99
L107:
	;
	goto L98
L108:
	;
	v427 = v396
	goto L94
L109:
	;
	v427 = v413
	goto L94
L110:
	;
	F_FlushPages(m, l3)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v451 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	F_WaitForParallelWorkersToFinish(m, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+118)))
	if v473 != int32(112) {
		goto L126
	} else {
		goto L127
	}
L117:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	if v456 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_UnregisterSnapshot(m, v455)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	F_DestroyParallelContext(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+72)) = v467 - int32(1)
	goto L123
L123:
	;
	goto L116
L124:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	F_MemoryContextDelete(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L135
	}
L125:
	;
	v488 = F_RelationGetNumberOfBlocksInFork(m, l1, l4)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L133
	}
L126:
	;
	if l4 != int32(3) {
		goto L124
	} else {
		goto L132
	}
L127:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[8]))
	if int32(0) < v477 {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v480 != 0 {
		goto L126
	} else {
		goto L129
	}
L129:
	;
	if l4 == int32(3) {
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v483 == int32(0) {
		goto L125
	} else {
		goto L131
	}
L131:
	;
	goto L124
L132:
	;
	goto L125
L133:
	;
	F_log_newpage_range(m, l1, l4, v488, int32(1))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	goto L124
L135:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l3)+184))
	F_MemoryContextDelete(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(2), base.I64_extend_i32_u(l0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			if l1 != 0 {
				v111 = int32(0)
				m.G0 = v8 - int32(-64)
				return v111
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutineByAmId_0), v8)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(82), int32(_a_F_GetIndexAmRoutineByAmId_2))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v34 = v32 + v33
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
			if v35 != int32(105) {
				if l1 != 0 {
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v111 = int32(0)
						m.G0 = v8 - int32(-64)
						return v111
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(_a_F_GetIndexAmRoutineByAmId_3)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v34 + int32(4)
							F_errmsg(m, int32(_a_F_GetIndexAmRoutineByAmId_4), v6+int32(-16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(97), int32(_a_F_GetIndexAmRoutineByAmId_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
				if v60 == int32(0) {
					if l1 != 0 {
						F_ReleaseCatCache(m, v12)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							v111 = int32(0)
							m.G0 = v8 - int32(-64)
							return v111
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v34 + int32(4)
								F_errmsg(m, int32(_a_F_GetIndexAmRoutineByAmId_5), v6+int32(-48))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(113), int32(_a_F_GetIndexAmRoutineByAmId_2))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
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
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v85 = F_OidFunctionCall0Coll(m, v60)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v87 = base.I32_wrap_i64(v85)
							if v87 != 0 {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
								if v88 == int32(444) {
									v111 = v87
									m.G0 = v8 - int32(-64)
									return v111
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v60
										F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutineByAmId_6), v6+int32(-32))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(43), int32(_a_F_GetIndexAmRoutineByAmId_7))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
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
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v60
									F_errmsg_internal(m, int32(_a_F_GetIndexAmRoutineByAmId_6), v6+int32(-32))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_GetIndexAmRoutineByAmId_1), int32(43), int32(_a_F_GetIndexAmRoutineByAmId_7))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
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
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+110)))
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
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+111)))
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
					F_errfinish(m, int32(_a_F_index_can_return_1), int32(815), int32(_a_F_index_can_return_2))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
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
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L57
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L53
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L49
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L46
	}
L6:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
	if v99 != 0 {
		goto L4
	} else {
		goto L31
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+131)))
	if v16 != int32(1) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v19 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L9
L11:
	;
	F_list_free(m, v19)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L12
	} else {
		goto L30
	}
L12:
	;
	return
L13:
	;
	if v19 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v24 <= v23 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v29 = v23
	goto L16
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v29<<(uint(int32(2))%32))))
	v42 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(v40))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	F_list_free(m, v19)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L25
	}
L18:
	;
	if v42 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47)+14)))
	F_ReleaseCatCache(m, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if v49 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v55 = v29 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v56 <= v55 {
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
	v29 = v55
	goto L16
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v67 + int32(4)
	F_errmsg(m, int32(_a_F_index_check_primary_key_0), v9+int32(-16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(222), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
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
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(0) < v100 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v108 = int32(0)
	v109 = v100
	goto L35
L33:
	;
	goto L34
L34:
	;
	m.G0 = v11 - int32(-64)
	return
L35:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(12)+v108<<(uint(int32(1))%32)))))
	if v117 == int32(0) {
		goto L3
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	v120 = base.I32_extend16_s(v117)
	if int32(0) <= v120 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v126 = F_SearchSysCache2(m, int32(7), v124, base.I64_extend_i32_u(v117))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	v139 = v109
	goto L40
L40:
	;
	v142 = v108 + int32(1)
	if v142 < v139 {
		v108 = v142
		v109 = v139
		goto L35
	} else {
		goto L45
	}
L41:
	;
	if v126 == int32(0) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
	v132 = v130 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+86)))
	if v133 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_ReleaseCatCache(m, v126)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v139 = v138
	goto L40
L45:
	;
	goto L36
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v40
	F_errmsg_internal(m, int32(_a_F_index_check_primary_key_3), v9+int32(-32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(169), int32(_a_F_index_check_primary_key_4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
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
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_index_check_primary_key_5), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(235), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_index_check_primary_key_6), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(252), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
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
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v120
	F_errmsg_internal(m, int32(_a_F_index_check_primary_key_7), v11)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(263), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
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
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v132 + int32(4)
	F_errmsg(m, int32(_a_F_index_check_primary_key_8), v9+int32(-48))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_index_check_primary_key_1), int32(270), int32(_a_F_index_check_primary_key_2))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+(l2^int32(-1))*int32(56))+16))
		v56 = v47
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, _c_F_index_compute_xid_horizon_for_tuples[3]))
		v50 = int32(56)
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+l2*v50-v50)+16))
		v56 = v55
	}
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = int64(0)
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v56
	v64 = F_palloc_mul(m, int32(8), l4)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v64
		v70 = F_palloc_mul(m, int32(6), l4)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v70
			if int32(0) < l4 {
				v78 = v59
				v80 = int32(0)
				for {
					v92 = int32(1)
					v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v80<<(uint(v92)%32)))))
					v99 = *(*int32)(unsafe.Add(mBase, uint32(v36+int32(20)+v95<<(uint(int32(2))%32))))
					v102 = v36 + v99&int32(_a_F_index_compute_xid_horizon_for_tuples_0)
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
					v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+4)))
					v107 = v64 + v80<<(uint(int32(3))%32)
					*(*uint16)(unsafe.Add(mBase, uint32(v107)+6)) = uint16(v78)
					*(*uint16)(unsafe.Add(mBase, uint32(v107)+4)) = uint16(v104)
					*(*int32)(unsafe.Add(mBase, uint32(v107))) = v103
					v113 = v70 + v80*int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v113)+2)) = v92
					*(*uint16)(unsafe.Add(mBase, uint32(v113))) = uint16(v95)
					v118 = v78 + v92
					*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v118
					v121 = v80 + v92
					if v121 != l4 {
						v78 = v118
						v80 = v121
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
			v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+76))
			v141 = m.T0[v140].(func(*base.Module, int32, int32) int32)(m, l1, v17+int32(4))
			mBase = m.M
			v142 = m.ExcPending
			if v142 != 0 {
				return int32(0)
			} else {
				v143 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
				F_pfree(m, v143)
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return int32(0)
				} else {
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
					F_pfree(m, v146)
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int32(0)
					} else {
						m.G0 = v17 + int32(32)
						return v141
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v3)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+188))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v23 = m.T0[v22].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v12, l0+int32(60), v15, l1, l0+int32(66), v8+int32(15))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		if v23 == int32(0) {
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v47 == int32(0) {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v50)
			} else {
			}
			m.G0 = v8 + int32(16)
			return v23
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+272))
			if v30 == int32(0) {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+268)))
				if v33 != int32(1) {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v47 == int32(0) {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v50)
					} else {
					}
					m.G0 = v8 + int32(16)
					return v23
				} else {
					F_pgstat_assoc_relation(m, v29)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+272))
						v40 = v39
						v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v40)+32)) = v41 + int64(1)
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v47 == int32(0) {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v50)
						} else {
						}
						m.G0 = v8 + int32(16)
						return v23
					}
				}
			} else {
				v40 = v30
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v40)+32)) = v41 + int64(1)
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v47 == int32(0) {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v50)
				} else {
				}
				m.G0 = v8 + int32(16)
				return v23
			}
		}
	}
}
func F_index_getattr_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v73 int64
	_ = v73
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v14 {
		v21 = l2 + l1<<(uint(int32(3))%32) + int32(20)
		v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21))))
		if v22 < int32(0) {
			v67 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				v73 = v67
				m.G0 = v10 + int32(16)
				return v73
			}
		} else {
			v27 = l0 + v22 + int32(8)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
			if v28 == int32(1) {
				v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+2)))
				if base.I32_popcnt(v31) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
						F_errmsg_internal(m, int32(_a_F_index_getattr_1_0), v10)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_index_getattr_1_1), int32(123), int32(_a_F_index_getattr_1_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					switch base.I32_ctz(v31) {
					case 0:
						v36 = int64(*(*int8)(unsafe.Add(mBase, uint32(v27))))
						v73 = v36
						m.G0 = v10 + int32(16)
						return v73
					case 1:
						v37 = int64(*(*int16)(unsafe.Add(mBase, uint32(v27))))
						v73 = v37
						m.G0 = v10 + int32(16)
						return v73
					case 2:
						v38 = int64(*(*int32)(unsafe.Add(mBase, uint32(v27))))
						v73 = v38
						m.G0 = v10 + int32(16)
						return v73
					case 3:
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
						v73 = v39
						m.G0 = v10 + int32(16)
						return v73
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
							F_errmsg_internal(m, int32(_a_F_index_getattr_1_0), v10)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_index_getattr_1_1), int32(123), int32(_a_F_index_getattr_1_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
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
				v73 = base.I64_extend_i32_u(v27)
				m.G0 = v10 + int32(16)
				return v73
			}
		}
	} else {
		v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		v57 = int32(1)
		if int32(base.Ui32(v56)>>(uint(l1-v57)%32))&v57 != 0 {
			v67 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				v73 = v67
				m.G0 = v10 + int32(16)
				return v73
			}
		} else {
			v62 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v62)
			v73 = int64(0)
			m.G0 = v10 + int32(16)
			return v73
		}
	}
}
func F_index_parallelscan_initialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_initialize[0]))
	if v13 != v11 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_index_parallelscan_initialize[1]))
		v17 = F_list_member_ptr(m, v16, v11)
		mBase = m.M
		v19 = v17
	} else {
		v19 = int32(1)
	}
	if v19 == int32(0) {
		v23 = F_EstimateSnapshotSpace(m, l2)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = F_add_size(m, int32(28), v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v27
				v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v29
				v31 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(l3)+12)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v33
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v35
				v38 = l3 + int32(28)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v48 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v38)+16)) = uint8(v50)
				*(*uint16)(unsafe.Add(mBase, uint32(v38)+18)) = uint16(v35)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v49
				*(*int64)(unsafe.Add(mBase, uint32(v38))) = v48
				*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v47
				*(*uint8)(unsafe.Add(mBase, uint32(v38)+17)) = uint8(v46)
				if v46&int32(1) != 0 {
					v61 = v45
				} else {
					v61 = v35
				}
				if v50 != 0 {
					v62 = v61
				} else {
					v62 = v45
				}
				*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v62
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v64 == int32(0) {
				} else {
					v68 = v64 << (uint(int32(2)) % 32)
					if v68 == int32(0) {
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						base.MemoryCopy(m, l3+int32(52), v73, v68)
					}
				}
				if v62 <= int32(0) {
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
					v80 = v78 << (uint(int32(2)) % 32)
					if v80 == int32(0) {
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						base.MemoryCopy(m, v38+v83<<(uint(int32(2))%32)+int32(24), v89, v80)
					}
				}
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+124))
				if v93 != 0 {
					v97 = (v25 + int32(7)) & int32(-8)
					*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v97
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+124))
					m.T0[v101].(func(*base.Module, int32))(m, v97+l3)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v115 + int32(4)
				F_errmsg(m, int32(_a_F_index_parallelscan_initialize_0), v9)
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_index_parallelscan_initialize_1), int32(511), int32(_a_F_index_parallelscan_initialize_2))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
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
