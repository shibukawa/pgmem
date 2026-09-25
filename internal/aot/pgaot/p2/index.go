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
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 float64
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 float64
	_ = v423
	var v424 int32
	_ = v424
	var v437 float64
	_ = v437
	var v439 int32
	_ = v439
	var v440 float64
	_ = v440
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
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
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+92)))
	if v456 == int32(0) {
		goto L108
	} else {
		goto L109
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
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v361 == int32(0) {
		goto L7
	} else {
		goto L91
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+108))
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
	switch v184 {
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
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v197 = F_shm_toc_allocate(m, v196, v121)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	F_DestroyParallelContext(m, v108)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	F_UnregisterSnapshot(m, v116)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+72)) = v192 - int32(1)
	goto L53
L53:
	;
	goto L9
L54:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v197))) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+8)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+4)) = v203
	v207 = v197 + int32(12)
	v208 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v207))), uint32(v208))
	*(*int64)(unsafe.Add(mBase, uint32(v207)+4)) = int64(-1)
	goto L55
L55:
	;
	v213 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v197)+24)), uint32(v213))
	*(*int64)(unsafe.Add(mBase, uint32(v197)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+28)) = v213
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v220, v197+int32(160), v116)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v226 = F_shm_toc_allocate(m, v225, v145)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_HnswInitLockTranche(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+132)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+112)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v197)+108)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v197)+88)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v197)+44)) = v230
	*(*int64)(unsafe.Add(mBase, uint32(v197)+48)) = int64(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v197)+40)), uint32(v230))
	v245 = v197 + int32(56)
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v247)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v245)+8)) = int64(-1)
	goto L59
L59:
	;
	v254 = v197 + int32(72)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v254))) = uint16(v256)
	*(*int32)(unsafe.Add(mBase, uint32(v254)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v254)+8)) = int64(-1)
	goto L60
L60:
	;
	v263 = v197 + int32(92)
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v263))) = uint16(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v263)+8)) = int64(-1)
	goto L61
L61:
	;
	v272 = v197 + int32(116)
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[7]))
	*(*uint16)(unsafe.Add(mBase, uint32(v272))) = uint16(v274)
	*(*int32)(unsafe.Add(mBase, uint32(v272)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = int64(-1)
	goto L62
L62:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	F_shm_toc_insert(m, v280, int64(-6917529027641081855), v197)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	F_shm_toc_insert(m, v284, int64(-6917529027641081854), v226)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	if v289 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v291 = F_shm_toc_allocate(m, v290, v178)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_LaunchParallelWorkers(m, v108)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L73
	}
L68:
	;
	if v178 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[6]))
	base.MemoryCopy(m, v291, v294, v178)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	F_shm_toc_insert(m, v296, int64(-6917529027641081853), v291)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v108
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v304 + int32(1)
	if v304 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_WaitForParallelWorkersToFinish(m, v108)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v331 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L83
	}
L77:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	switch v316 {
	case 0, 5:
		goto L79
	default:
		goto L78
	}
L78:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	F_DestroyParallelContext(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	F_UnregisterSnapshot(m, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+72)) = v325 - int32(1)
	goto L82
L82:
	;
	goto L9
L83:
	;
	if v331 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v333
	F_errmsg(m, int32(_a_F_BuildIndex_1_6), v16)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+196)) = v96
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	F_HnswParallelScanAndInsert(m, v344, v345, v346, v347, int32(1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_1_7), int32(1051), int32(_a_F_BuildIndex_1_8))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	F_WaitForParallelWorkersToAttach(m, v108)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L9
L91:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v364 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v437
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v440 = *(*float64)(unsafe.Add(mBase, uint32(v439)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v440
	goto L7
L93:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+8))
	v369 = v365 + int32(24)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	goto L96
L94:
	;
	goto L95
L95:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v414 = int32(1)
	v415 = int32(0)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v361)+188))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+140))
	v423 = m.T0[v422].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v361, v412, v413, v414, v415, v414, v415, int32(-1), int32(_a_F_BuildIndex_1_9), l3, v415)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L107
	}
L96:
	;
	v386 = base.AtomicRmwXchg32(m, v369, int32(0), int32(1))
	if v386 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+160)) = v365 + int32(40)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+204)) = v404
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v365)+32))
	v407 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v365)+24)), uint32(v407))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L106
	}
L98:
	;
	F_s_lock(m, v369, int32(_a_F_BuildIndex_1_7), int32(769), int32(_a_F_BuildIndex_1_10))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v365)+28))
	if v370 != v392 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	v394 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v369))), uint32(v394))
	F_ConditionVariableSleep(m, v365+int32(12), int32(134217767))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
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
	v437 = v406
	goto L92
L107:
	;
	v437 = v423
	goto L92
L108:
	;
	F_FlushPages(m, l3)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v461 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	F_WaitForParallelWorkersToFinish(m, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+118)))
	if v481 != int32(112) {
		goto L123
	} else {
		goto L124
	}
L115:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
	switch v466 {
	case 0, 5:
		goto L117
	default:
		goto L116
	}
L116:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	F_DestroyParallelContext(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	F_UnregisterSnapshot(m, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[4]))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v474)+72)) = v475 - int32(1)
	goto L120
L120:
	;
	goto L114
L121:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	F_MemoryContextDelete(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L132
	}
L122:
	;
	v496 = F_RelationGetNumberOfBlocksInFork(m, l1, l4)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
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
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_1[8]))
	if int32(0) < v485 {
		goto L122
	} else {
		goto L125
	}
L125:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v488 != 0 {
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
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v491 == int32(0) {
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
	F_log_newpage_range(m, l1, l4, v496, int32(1))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L121
L132:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l3)+184))
	F_MemoryContextDelete(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
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
