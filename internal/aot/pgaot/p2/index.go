package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int64
	_ = v223
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
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
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 float64
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 float64
	_ = v405
	var v406 int32
	_ = v406
	var v419 float64
	_ = v419
	var v421 int32
	_ = v421
	var v422 float64
	_ = v422
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v55 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v28 != int32(1) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v31 = int32(4548900)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v34 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v33 + v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v37 + v34
	*(*int64)(unsafe.Add(mBase, uint32(v24+int32(80))+232)) = int64(2)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v45 + v34
	v51 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v51 - v34
	goto L4
L7:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+92)))
	if v438 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+56))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+56))
	v61 = F_plan_create_index_workers(m, v58, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v345 == int32(0) {
		goto L7
	} else {
		goto L92
	}
L10:
	;
	if v61 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)+180))
	if v65 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v80 <= int32(0) {
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	v80 = v77
	goto L12
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+108))
	if v68 == int32(-1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[97]))
	if v68 < v72 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v74 = v68
	goto L18
L17:
	;
	v74 = v72
	goto L18
L18:
	;
	v80 = v74
	goto L12
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+121)))
	v86 = F_palloc0(m, int32(20))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+72)) = v91 + int32(1)
	goto L21
L21:
	;
	v98 = F_CreateParallelContext(m, int32(219134), int32(291597), v80)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v84 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v102 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v106 = int32(4212032)
	goto L25
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v109 = F_table_parallelscan_estimate(m, v108, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v104 = F_RegisterSnapshot(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v106 = v104
	goto L25
L28:
	;
	v111 = F_add_size(m, int32(160), v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v98)+36))
	v118 = F_add_size(m, v113, (v111+int32(31))&int32(-32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+36)) = v118
	v122 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v124 = F_mul_size(m, v122, int32(1024))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v98)+36))
	v128 = int32(3145728)
	if base.Ui32(v128) < base.Ui32(v124) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v132 = v124 - v128
	goto L34
L33:
	;
	v132 = v124
	goto L34
L34:
	;
	if base.Ui32(int32(2147483647)) <= base.Ui32(v132) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = int32(2147483647)
	goto L37
L36:
	;
	v135 = v132
	goto L37
L37:
	;
	v140 = F_add_size(m, v126, (v135+int32(31))&int32(-32))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+36)) = v140
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v98)+40))
	v145 = F_add_size(m, v143, int32(2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+40)) = v145
	v149 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	if v149 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v98)+36))
	v151 = F_strlen(m, v149)
	mBase = m.M
	v156 = F_add_size(m, v150, v151&int32(-32)+int32(32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v168 = int32(1)
	goto L42
L42:
	;
	F_InitializeParallelDSM(m, v98)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+36)) = v156
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v98)+40))
	v161 = F_add_size(m, v159, int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+40)) = v161
	v168 = v151 + int32(1)
	goto L42
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v98)+44))
	if v171 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	switch v174 {
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
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	v187 = F_shm_toc_allocate(m, v186, v111)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L54
	}
L49:
	;
	F_DestroyParallelContext(m, v98)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	F_UnregisterSnapshot(m, v106)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+72)) = v182 - int32(1)
	goto L53
L53:
	;
	goto L9
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+56))
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+4)) = v193
	v197 = v187 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(-4294967296)
	goto L55
L55:
	;
	v202 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+32)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v187)+24)) = v202
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v206, v187+int32(160), v106)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	v212 = F_shm_toc_allocate(m, v211, v135)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_HnswInitLockTranche(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+132)) = uint8(v216)
	*(*int32)(unsafe.Add(mBase, uint32(v187)+112)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v187)+108)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v187)+88)) = v216
	v223 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v187)+48)) = v223
	*(*int64)(unsafe.Add(mBase, uint32(v187)+40)) = v223
	v228 = v187 + int32(56)
	v230 = *(*int32)(unsafe.Add(mBase, _consts[1454]))
	*(*uint16)(unsafe.Add(mBase, uint32(v228))) = uint16(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v228)+8)) = int64(-1)
	goto L59
L59:
	;
	v237 = v187 + int32(72)
	v239 = *(*int32)(unsafe.Add(mBase, _consts[1454]))
	*(*uint16)(unsafe.Add(mBase, uint32(v237))) = uint16(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v237)+8)) = int64(-1)
	goto L60
L60:
	;
	v246 = v187 + int32(92)
	v248 = *(*int32)(unsafe.Add(mBase, _consts[1454]))
	*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v246)+8)) = int64(-1)
	goto L61
L61:
	;
	v255 = v187 + int32(116)
	v257 = *(*int32)(unsafe.Add(mBase, _consts[1454]))
	*(*uint16)(unsafe.Add(mBase, uint32(v255))) = uint16(v257)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = int64(-1)
	goto L62
L62:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	F_shm_toc_insert(m, v263, int64(-6917529027641081855), v187)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	F_shm_toc_insert(m, v267, int64(-6917529027641081854), v212)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	if v272 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	v274 = F_shm_toc_allocate(m, v273, v168)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_LaunchParallelWorkers(m, v98)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L74
	}
L68:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	if v168 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v98)+52))
	F_shm_toc_insert(m, v280, int64(-6917529027641081853), v279)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	v278 = F__emscripten_memcpy_bulkmem(m, v274, v277, v168)
	mBase = m.M
	v279 = v278
	goto L72
L71:
	;
	v279 = v274
	goto L72
L72:
	;
	goto L69
L73:
	;
	goto L67
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v98
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v288 + int32(1)
	if v288 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_WaitForParallelWorkersToFinish(m, v98)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v315 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L84
	}
L78:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	switch v300 {
	case 0, 5:
		goto L80
	default:
		goto L79
	}
L79:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	F_DestroyParallelContext(m, v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	F_UnregisterSnapshot(m, v299)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+72)) = v309 - int32(1)
	goto L83
L83:
	;
	goto L9
L84:
	;
	if v315 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v317
	F_errmsg(m, int32(142415), v16)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+196)) = v86
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	F_HnswParallelScanAndInsert(m, v328, v329, v330, v331, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	F_errfinish(m, int32(523283), int32(1051), int32(322097))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	F_WaitForParallelWorkersToAttach(m, v98)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L9
L92:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v348 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l3)+160))
	v422 = *(*float64)(unsafe.Add(mBase, uint32(v421)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v422
	goto L7
L94:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+8))
	v353 = v349 + int32(24)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	goto L97
L95:
	;
	goto L96
L96:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v396 = int32(1)
	v397 = int32(0)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v345)+188))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+140))
	v405 = m.T0[v404].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v345, v394, v395, v396, v397, v396, v397, int32(-1), int32(7658), l3, v397)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L108
	}
L97:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = int32(1)
	if v368 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+160)) = v349 + int32(40)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+204)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v349)+24)) = int32(0)
	v391 = *(*float64)(unsafe.Add(mBase, uint32(v349)+32))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L107
	}
L99:
	;
	F_s_lock(m, v353, int32(523283), int32(769), int32(298284))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v349)+28))
	if v354 != v376 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L101
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = int32(0)
	F_ConditionVariableSleep(m, v349+int32(12), int32(134217767))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	goto L98
L106:
	;
	goto L97
L107:
	;
	v419 = v391
	goto L93
L108:
	;
	v419 = v405
	goto L93
L109:
	;
	F_FlushPages(m, l3)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l3)+196))
	if v443 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L111
L113:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	F_WaitForParallelWorkersToFinish(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+118)))
	if v463 != int32(112) {
		goto L124
	} else {
		goto L125
	}
L116:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v443)+12))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	switch v448 {
	case 0, 5:
		goto L118
	default:
		goto L117
	}
L117:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	F_DestroyParallelContext(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	F_UnregisterSnapshot(m, v447)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v456)+72)) = v457 - int32(1)
	goto L121
L121:
	;
	goto L115
L122:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	F_MemoryContextDelete(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L133
	}
L123:
	;
	v478 = F_RelationGetNumberOfBlocksInFork(m, l1, l4)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L131
	}
L124:
	;
	if l4 != int32(3) {
		goto L122
	} else {
		goto L130
	}
L125:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if int32(0) < v467 {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v470 != 0 {
		goto L124
	} else {
		goto L127
	}
L127:
	;
	if l4 == int32(3) {
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v473 == int32(0) {
		goto L123
	} else {
		goto L129
	}
L129:
	;
	goto L122
L130:
	;
	goto L123
L131:
	;
	F_log_newpage_range(m, l1, l4, v478, int32(1))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L122
L133:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l3)+184))
	F_MemoryContextDelete(m, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
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
					F_errmsg_internal(m, int32(58072), v8)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(521273), int32(69), int32(486342))
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
							*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(534698)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v33 + int32(4)
							F_errmsg(m, int32(202129), v6+int32(-16))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521273), int32(84), int32(486342))
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
								F_errmsg(m, int32(230728), v6+int32(-48))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(521273), int32(100), int32(486342))
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
										F_errmsg_internal(m, int32(116004), v6+int32(-32))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(521273), int32(43), int32(390779))
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
									F_errmsg_internal(m, int32(116004), v6+int32(-32))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(521273), int32(43), int32(390779))
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v72 != 0 {
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
	v32 = v6
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v32<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+109)))
	if v41 == int32(1) {
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
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+110)))
	if v46 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v60 = v32 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v60 < v61 {
		v32 = v60
		goto L6
	} else {
		goto L19
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+64))
	if v49 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v40)+104))
	if base.F64_lt(v50, float64(1)) == int32(0) {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v56 = F_lappend(m, v55, v40)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v56
	goto L12
L19:
	;
	goto L7
L20:
	;
	v73 = int32(0)
	v76 = F_build_index_paths(m, l0, l1, l2, l3, v73, int32(1), v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v79 = F_list_concat(m, v78, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v79
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v13 != v9 {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[122]))
		v17 = F_list_member_ptr(m, v16, v9)
		mBase = m.M
		v18 = v17
	} else {
		v18 = int32(1)
	}
	if v18 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
		if v22 != 0 {
			v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v28 = v23
				m.G0 = v7 + int32(16)
				return v28
			}
		} else {
			v28 = int32(0)
			m.G0 = v7 + int32(16)
			return v28
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v40 + int32(4)
				F_errmsg(m, int32(458845), v7)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(520736), int32(837), int32(256990))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
	var v104 int32
	_ = v104
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
	var v131 int32
	_ = v131
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
	F_errmsg(m, int32(459983), v8+int32(-16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(514633), int32(221), int32(21545))
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
	v104 = v95
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
	v131 = v104
	goto L40
L40:
	;
	v134 = v103 + int32(1)
	if v134 < v131 {
		v103 = v134
		v104 = v131
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
	v131 = v130
	goto L40
L45:
	;
	goto L36
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v38
	F_errmsg_internal(m, int32(42843), v8+int32(-32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(514633), int32(168), int32(22914))
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
	F_errmsg(m, int32(167095), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(514633), int32(234), int32(21545))
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
	F_errmsg(m, int32(154999), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(514633), int32(251), int32(21545))
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
	F_errmsg_internal(m, int32(50140), v10)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(514633), int32(262), int32(21545))
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
	F_errmsg(m, int32(557489), v8+int32(-48))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(514633), int32(269), int32(21545))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if l2 < int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l2^int32(-1))<<(uint(int32(2))%32))))
		v35 = v27
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		v35 = v29 + l2<<(uint(int32(13))%32) + int32(-8192)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l0
	if l2 < int32(0) {
		v40 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(l2^int32(-1))<<(uint(int32(6))%32))+16))
		v55 = v46
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+l2<<(uint(int32(6))%32)+int32(-64))+16))
		v55 = v54
	}
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(0)
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v55
	v64 = F_palloc(m, l4<<(uint(int32(3))%32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v64
		v71 = F_palloc(m, l4*int32(6))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v71
			if int32(0) < l4 {
				v79 = v58
				v81 = int32(0)
				for {
					v94 = v64 + v81<<(uint(int32(3))%32)
					v95 = int32(1)
					v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v81<<(uint(v95)%32)))))
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v98<<(uint(int32(2))%32)+(v35+int32(24))-int32(4))))
					v107 = v35 + v104&int32(32767)
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v108
					v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)) = uint16(v79)
					*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)) = uint16(v110)
					v115 = v71 + v81*int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v115)+2)) = v95
					*(*uint16)(unsafe.Add(mBase, uint32(v115))) = uint16(v98)
					v120 = v79 + v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v120
					v123 = v81 + v95
					if v123 != l4 {
						v79 = v120
						v81 = v123
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+188))
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+76))
			v142 = m.T0[v141].(func(*base.Module, int32, int32) int32)(m, l1, v16+int32(4))
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
				F_pfree(m, v144)
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
					F_pfree(m, v147)
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(32)
						return v142
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
	v13 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v13 != 0 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[125])))
		if v15&int32(1) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(352697), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(342435), int32(1218), int32(401413))
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	v14 = l1 - int32(1)
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v15 {
		v22 = l2 + v14<<(uint(int32(4))%32) + int32(20)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		if v23 < int32(0) {
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
			v28 = l0 + v23 + int32(8)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
			if v29 != int32(1) {
				v70 = v28
				m.G0 = v9 + int32(16)
				return v70
			} else {
				v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
				switch v32&int32(65535) - int32(1) {
				case 0:
					v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
					v70 = v37
					m.G0 = v9 + int32(16)
					return v70
				case 1:
					v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28))))
					v70 = v38
					m.G0 = v9 + int32(16)
					return v70
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v32
						F_errmsg_internal(m, int32(504959), v9)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(342394), int32(70), int32(73673))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v70 = v39
					m.G0 = v9 + int32(16)
					return v70
				}
			}
		}
	} else {
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if int32(base.Ui32(v55)>>(uint(v14)%32))&int32(1) != 0 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v19 != v15 {
		v22 = *(*int32)(unsafe.Add(mBase, _consts[122]))
		v23 = F_list_member_ptr(m, v22, v15)
		mBase = m.M
		v24 = v23
	} else {
		v24 = int32(1)
	}
	if v24 == int32(0) {
		v28 = F_EstimateSnapshotSpace(m, l2)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = F_add_size(m, int32(32), v28)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l7)+8)) = v32
				v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(l7))) = v34
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				*(*int64)(unsafe.Add(mBase, uint32(l7)+12)) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l7)+20)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(l7)+24)) = int64(0)
				v43 = l7 + int32(32)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v52 = *(*int64)(unsafe.Add(mBase, uint32(l2)+4))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+29)))
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
				*(*uint8)(unsafe.Add(mBase, uint32(v43)+16)) = uint8(v55)
				*(*uint8)(unsafe.Add(mBase, uint32(v43)+17)) = uint8(v54)
				*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v53
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v52
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v51
				if v54 != 0 {
					v62 = v50
				} else {
					v62 = int32(0)
				}
				if v55 != 0 {
					v63 = v62
				} else {
					v63 = v50
				}
				*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v63
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v65 != 0 {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v71 = F___memcpy(m, l7+int32(56), v68, v65<<(uint(int32(2))%32))
					mBase = m.M
				} else {
				}
				if int32(0) < v63 {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v75 = int32(2)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
					v84 = F___memcpy(m, v43+v74<<(uint(v75)%32)+int32(24), v80, v81<<(uint(v75)%32))
					mBase = m.M
				} else {
				}
				v88 = (v30 + int32(7)) & int32(-8)
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l7)+24)) = v88
					v93 = l5<<(uint(int32(3))%32) + int32(8)
					v94 = F_add_size(m, v88, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l7)+24))
						v97 = l7 + v96
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v97
						v101 = F__emscripten_memset_bulkmem(m, v97, base.I32_extend8_s(int32(0)), v93)
						mBase = m.M
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
						*(*int32)(unsafe.Add(mBase, uint32(v102))) = l5
						v109 = (v94 + int32(7)) & int32(-8)
						if l4 == int32(0) {
							m.G0 = v13 + int32(16)
							return
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+124))
							if v114 == int32(0) {
								m.G0 = v13 + int32(16)
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l7)+28)) = v109
								v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+124))
								m.T0[v120].(func(*base.Module, int32))(m, v109+l7)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return
								} else {
									m.G0 = v13 + int32(16)
									return
								}
							}
						}
					}
				} else {
					v109 = v88
					if l4 == int32(0) {
						m.G0 = v13 + int32(16)
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+124))
						if v114 == int32(0) {
							m.G0 = v13 + int32(16)
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l7)+28)) = v109
							v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+124))
							m.T0[v120].(func(*base.Module, int32))(m, v109+l7)
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
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
		v129 = m.ExcPending
		if v129 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v132 = m.ExcPending
			if v132 != 0 {
				return
			} else {
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v133 + int32(4)
				F_errmsg(m, int32(458845), v13)
				mBase = m.M
				v139 = m.ExcPending
				if v139 != 0 {
					return
				} else {
					F_errfinish(m, int32(520736), int32(520), int32(358296))
					mBase = m.M
					v144 = m.ExcPending
					if v144 != 0 {
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
