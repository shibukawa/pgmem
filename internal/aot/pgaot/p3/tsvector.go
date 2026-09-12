package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_close_tsvector_parser(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_pfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_init_tsvector_parser(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v6 = F_palloc(m, int32(28))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v15 = F_palloc(m, v10)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v15
			v19 = *(*int32)(unsafe.Add(mBase, _consts[451]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v20*int32(28))+uint32(_consts[970])))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = l2
			v29 = int32(1)
			v30 = int32(base.Ui32(l1)>>(uint(int32(2))%32)) & v29
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+22)) = uint8(v30)
			v35 = int32(base.Ui32(l1)>>(uint(v29)%32)) & v29
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+21)) = uint8(v35)
			v38 = l1 & v29
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)) = uint8(v38)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v25
			return v6
		}
	}
}
func F_make_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var __phi69 int32
	_ = __phi69
	var v71 int32
	_ = v71
	var __phi71 int32
	_ = __phi71
	var v75 int32
	_ = v75
	var __phi75 int32
	_ = __phi75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
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
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
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
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 <= v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L10
	} else {
		goto L107
	}
L2:
	;
	v329 = v310 + v315<<(uint(int32(2))%32) + int32(8)
	v330 = F_palloc0(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L77
	}
L3:
	;
	v310 = v2
	v315 = v21
	goto L2
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v21 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v268 = int32(0)
	v271 = v268
	v272 = v268
	goto L70
L7:
	;
	v27 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v27)
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)))
	v31 = F_palloc(m, int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_pg_qsort(m, v24, v21, int32(16), int32(1171))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v31
	v36 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v31))) = uint16(v36)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v40 = int32(16383)
	if base.Ui32(v40) <= base.Ui32(v29) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = v40
	goto L14
L13:
	;
	v43 = v29
	goto L14
L14:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)) = uint16(v43)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v257 = v36
	goto L6
L15:
	;
	v51 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v51)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)))
	v55 = F_palloc(m, int32(4))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v55
	v58 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v58)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v61 = int32(16383)
	if base.Ui32(v61) <= base.Ui32(v53) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v64 = v61
	goto L19
L18:
	;
	v64 = v53
	goto L19
L19:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v64)
	__phi69 = v24
	__phi71 = v24 + int32(16)
	__phi75 = v24
	v69 = __phi69
	v71 = __phi71
	v75 = __phi75
	goto L20
L20:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+18)))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+2)))
	if v84 == v85 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v246 = (v230 - v24 + int32(16)) >> (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v246
	v248 = int32(0)
	if v246 <= v248 {
		v310 = v248
		v315 = v246
		goto L2
	} else {
		goto L69
	}
L22:
	;
	v237 = v71 + int32(16)
	if (v237-v24)>>(uint(int32(4))%32) < v21 {
		__phi69 = v230
		__phi71 = v237
		__phi75 = v71
		v69 = __phi69
		v71 = __phi71
		v75 = __phi75
		goto L20
	} else {
		goto L68
	}
L23:
	;
	F_pfree(m, v87)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L10
	} else {
		goto L46
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v84 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+18)) = uint16(v84)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v137
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+24)))
	v140 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+22)) = uint16(v140)
	v143 = v69 + int32(24)
	v145 = F_palloc(m, int32(4))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L42
	}
L27:
	;
	if v132 == int32(0) {
		goto L23
	} else {
		goto L41
	}
L28:
	;
	v132 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v94 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v95 = v87
	v96 = v88
	v97 = v84
	v98 = v94
	goto L35
L32:
	;
	v120 = v88
	v124 = int32(0)
	goto L33
L33:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v132 = v124 - v125
	goto L27
L34:
	;
	v120 = v115
	v124 = v117
	goto L33
L35:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v98 != v100 {
		v115 = v96
		v117 = v98
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v115 = v109
	v117 = int32(0)
	goto L34
L37:
	;
	if v100 == int32(0) {
		v115 = v96
		v117 = v98
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v105 = v97 - int32(1)
	if v105 == int32(0) {
		v115 = v96
		v117 = v98
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v108 = int32(1)
	v109 = v96 + v108
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v110 != 0 {
		v95 = v95 + v108
		v96 = v109
		v97 = v105
		v98 = v110
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	goto L26
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v145
	v148 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v145))) = uint16(v148)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v151 = int32(16383)
	if base.Ui32(v151) <= base.Ui32(v139) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v154 = v151
	goto L45
L44:
	;
	v154 = v139
	goto L45
L45:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v150)+2)) = uint16(v154)
	v230 = v69 + int32(16)
	goto L22
L46:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	if base.Ui32(int32(254)) < base.Ui32(v161) {
		v230 = v69
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160+v161<<(uint(int32(1))%32)))))
	if v167 == int32(16383) {
		v230 = v69
		goto L22
	} else {
		goto L48
	}
L48:
	;
	v170 = int32(16383)
	v172 = v75 + int32(24)
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172))))
	if base.Ui32(v170) <= base.Ui32(v173) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v176 = v170
	goto L51
L50:
	;
	v176 = v173
	goto L51
L51:
	;
	if v176 == v167 {
		v230 = v69
		goto L22
	} else {
		goto L52
	}
L52:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)))
	if base.Ui32(v178) <= base.Ui32(v161+int32(1)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = v178 << (uint(int32(1)) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+6)) = uint16(v183)
	v189 = F_repalloc(m, v160, v178<<(uint(int32(2))%32)&int32(131068))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L10
	} else {
		goto L56
	}
L54:
	;
	v194 = v173
	v195 = v160
	v196 = v161
	goto L55
L55:
	;
	v198 = v196 & int32(65535)
	if v198 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v189
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189))))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172))))
	v194 = v193
	v195 = v189
	v196 = v192
	goto L55
L57:
	;
	v221 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v195+v220<<(uint(v221)%32))+2)) = uint16(v219)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v225))))
	v228 = v226 + v221
	*(*uint16)(unsafe.Add(mBase, uint32(v225))) = uint16(v228)
	v230 = v69
	goto L22
L58:
	;
	v201 = int32(16383)
	v203 = v194 & int32(65535)
	if base.Ui32(v201) <= base.Ui32(v203) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v208 = int32(16383)
	v210 = v194 & int32(65535)
	if base.Ui32(v208) <= base.Ui32(v210) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v206 = v201
	goto L63
L62:
	;
	v206 = v203
	goto L63
L63:
	;
	v219 = v206
	v220 = int32(0)
	goto L57
L64:
	;
	v213 = v208
	goto L66
L65:
	;
	v213 = v210
	goto L66
L66:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195+v198<<(uint(int32(1))%32)))))
	if v213 == v217 {
		v230 = v69
		goto L22
	} else {
		goto L67
	}
L67:
	;
	v219 = v213
	v220 = v198
	goto L57
L68:
	;
	goto L21
L69:
	;
	v257 = v246
	goto L6
L70:
	;
	v288 = v267 + v272<<(uint(int32(4))%32)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288)+2)))
	v290 = v271 + v289
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288)+6)))
	if v291 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if int32(1048576) <= v303 {
		goto L1
	} else {
		goto L76
	}
L72:
	;
	v292 = int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)+8))
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296))))
	v303 = (v290+v292)&int32(-2) + v297<<(uint(v292)%32) + int32(2)
	goto L74
L73:
	;
	v303 = v290
	goto L74
L74:
	;
	v305 = v272 + int32(1)
	if v305 != v257 {
		v271 = v303
		v272 = v305
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	v310 = v303
	v315 = v257
	goto L2
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v329 << (uint(int32(2)) % 32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v330)+4)) = v335
	if int32(0) < v335 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v340 = v330 + int32(8)
	v343 = v340 + v335<<(uint(int32(2))%32)
	v346 = int32(0)
	v353 = v340
	v357 = v2
	goto L81
L79:
	;
	goto L80
L80:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v590 != 0 {
		goto L103
	} else {
		goto L104
	}
L81:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v362 = int32(1)
	v365 = v357 << (uint(int32(4)) % 32)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365+v366)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v361&v362 | v368<<(uint(v362)%32)&int32(4094) | v346<<(uint(int32(12))%32)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v380 = v379 + v365
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v380)+2)))
	if v382 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L80
L83:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v386 = v385 + v365
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386)+2)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	F_pfree(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L10
	} else {
		goto L87
	}
L84:
	;
	v383 = F__emscripten_memcpy_bulkmem(m, v346+v343, v381, v382)
	mBase = m.M
	goto L86
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	v391 = v346 + v387
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v393 = v392 + v365
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v393)+6)))
	if v394 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v571 = v357 + int32(1)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v571 < v572 {
		v346 = v553
		v353 = v353 + int32(4)
		v357 = v571
		goto L81
	} else {
		goto L102
	}
L89:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v393)+8))
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v395))))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v398 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v397 | v398
	v404 = (v391 + v398) & int32(-2)
	*(*uint16)(unsafe.Add(mBase, uint32(v343+v404))) = uint16(v396)
	if v396 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = v548 & int32(-2)
	v553 = v391
	goto L88
L92:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v538+v365)+8))
	F_pfree(m, v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L10
	} else {
		goto L101
	}
L93:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v410 = int32(2)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	v416 = int32(1)
	v427 = v340 + v409<<(uint(v410)%32) + (int32(base.Ui32(v413)>>(uint(int32(12))%32))+int32(base.Ui32(v413)>>(uint(v416)%32))&int32(2047)+v416)&int32(4194302) + v410
	v428 = int32(0)
	if v396 != v416 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v435 = v428
	v437 = int32(0)
	goto L97
L95:
	;
	v488 = v428
	goto L96
L96:
	;
	if v396&int32(1) == int32(0) {
		goto L92
	} else {
		goto L100
	}
L97:
	;
	v450 = int32(1)
	v451 = v435 << (uint(v450) % 32)
	v452 = v427 + v451
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	v454 = int32(16383)
	v455 = v453 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v452))) = uint16(v455)
	v457 = int32(2)
	v458 = v451 | v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459+v365)+8))
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v458+v461))))
	v465 = v463 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v452))) = uint16(v465)
	v467 = v458 + v427
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467))))
	v470 = v468 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v467))) = uint16(v470)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v472+v365)+8))
	v476 = v435 + v457
	v480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474+v476<<(uint(v450)%32)))))
	v482 = v480 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v467))) = uint16(v482)
	v485 = v437 + v457
	if v485 != v396&int32(65534) {
		v435 = v476
		v437 = v485
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v488 = v476
	goto L96
L99:
	;
	goto L98
L100:
	;
	v508 = v488 << (uint(int32(1)) % 32)
	v509 = v427 + v508
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v509))))
	v511 = int32(16383)
	v512 = v510 & v511
	*(*uint16)(unsafe.Add(mBase, uint32(v509))) = uint16(v512)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514+v365)+8))
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v516+v508)+2)))
	v520 = v518 & v511
	*(*uint16)(unsafe.Add(mBase, uint32(v509))) = uint16(v520)
	goto L92
L101:
	;
	v553 = v404 + v396<<(uint(int32(1))%32) + int32(2)
	goto L88
L102:
	;
	goto L82
L103:
	;
	F_pfree(m, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L10
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	m.G0 = v19 + int32(16)
	return v330
L106:
	;
	goto L105
L107:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(1048575)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v303
	F_errmsg(m, int32(656985), v19)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(490648), int32(194), int32(207790))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_filter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	F_deconstruct_array_builtin(m, v33, int32(18), v25+int32(12), v25+int32(8), v25+int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if int32(0) < v45 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L100
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v52 = int32(0)
	v55 = v2
	goto L9
L7:
	;
	v122 = v2
	goto L8
L8:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v138 = F_palloc0(m, int32(base.Ui32(v135)>>(uint(int32(2))%32)))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L22
	}
L9:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v50))))
	if v74 == int32(1) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v122 = v107 & int32(255)
	goto L8
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v49+v52<<(uint(int32(2))%32))))
	switch v81&int32(255) - int32(65) {
	case 0, 32:
		v106 = int32(8)
		goto L12
	case 1, 33:
		goto L13
	case 2, 34:
		goto L16
	case 3, 35:
		goto L15
	default:
		goto L14
	}
L12:
	;
	v107 = v106 | v55
	v109 = v52 + int32(1)
	if v109 != v45 {
		v52 = v109
		v55 = v107
		goto L9
	} else {
		goto L21
	}
L13:
	;
	v106 = int32(4)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v106 = int32(1)
	goto L12
L16:
	;
	v106 = int32(2)
	goto L12
L17:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = base.I32_extend8_s(v81)
	F_errmsg(m, int32(712589), v25)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(494019), int32(869), int32(215288))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	goto L10
L22:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v140
	v143 = v138 + int32(8)
	v146 = v143 + v140<<(uint(int32(2))%32)
	if v140 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v359
	if v359 != v140 {
		goto L47
	} else {
		goto L48
	}
L24:
	;
	v357 = v2
	v359 = v2
	goto L23
L25:
	;
	goto L26
L26:
	;
	v150 = v28 + int32(8)
	v156 = v140
	v159 = v2
	v161 = v2
	v168 = v2
	goto L27
L27:
	;
	v178 = v150 + v168<<(uint(int32(2))%32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v179&int32(1) == int32(0) {
		v331 = v159
		v333 = v161
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v357 = v331
	v359 = v333
	goto L23
L29:
	;
	v349 = v168 + int32(1)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v349 < v350 {
		v156 = v350
		v159 = v331
		v161 = v333
		v168 = v349
		goto L27
	} else {
		goto L46
	}
L30:
	;
	v187 = int32(1)
	v192 = int32(base.Ui32(v179)>>(uint(v187)%32))&int32(2047) + v187
	v198 = v150 + v156<<(uint(int32(2))%32) + (v192+int32(base.Ui32(v179)>>(uint(int32(12))%32)))&int32(2097150)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	if v199 == int32(0) {
		v331 = v159
		v333 = v161
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v205 = v146 + (v192+v159)&int32(-2)
	v206 = int32(2)
	v210 = int32(0)
	v213 = v210
	v214 = v210
	v216 = v199
	goto L32
L32:
	;
	v234 = int32(1)
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198+v206+v213<<(uint(v234)%32)))))
	if int32(base.Ui32(v122)>>(uint(int32(base.Ui32(v237)>>(uint(int32(14))%32)))%32))&v234 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v250 == int32(0) {
		v331 = v159
		v333 = v161
		goto L29
	} else {
		goto L38
	}
L34:
	;
	v243 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v205+v206+v214<<(uint(v243)%32)))) = uint16(v237)
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	v250 = v214 + v243
	v251 = v247
	goto L36
L35:
	;
	v250 = v214
	v251 = v216
	goto L36
L36:
	;
	v253 = v213 + int32(1)
	if base.Ui32(v253) < base.Ui32(v251&int32(65535)) {
		v213 = v253
		v214 = v250
		v216 = v251
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v259 = int32(2)
	v262 = v143 + v161<<(uint(v259)%32)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v264 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v263 | v264
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v270 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v267&int32(4094) | v159<<(uint(v270)%32) | v264
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v284 = int32(base.Ui32(v277)>>(uint(v264)%32)) & int32(2047)
	if v284 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v205))) = uint16(v250)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v289 = int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v300&v289 != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v285 = F__emscripten_memcpy_bulkmem(m, v159+v146, v150+v35<<(uint(int32(2))%32)+int32(base.Ui32(v277)>>(uint(v270)%32)), v284)
	mBase = m.M
	goto L42
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v304 = int32(2)
	v307 = int32(1)
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+v303<<(uint(v304)%32)+(int32(base.Ui32(v300)>>(uint(v307)%32))&int32(2047)+int32(base.Ui32(v300)>>(uint(int32(12))%32))+v307)&int32(4194302)))))
	v324 = v319<<(uint(v307)%32) + v304
	goto L45
L44:
	;
	v324 = v259
	goto L45
L45:
	;
	v331 = v324 + ((int32(base.Ui32(v288)>>(uint(v289)%32))&int32(2047)+v289)&int32(4094) + v159)
	v333 = v161 + v289
	goto L29
L46:
	;
	goto L28
L47:
	;
	v380 = v143 + v359<<(uint(int32(2))%32)
	if v380 == v146 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v526 = v140
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v357<<(uint(int32(2))%32) + v526<<(uint(int32(4))%32) + int32(32)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v533 != v28 {
		goto L96
	} else {
		goto L97
	}
L50:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v526 = v525
	goto L49
L51:
	;
	goto L50
L52:
	;
	v384 = v380 + v357
	if base.Ui32(v146-v384) <= base.Ui32(int32(0)-v357<<(uint(int32(1))%32)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v391 = F___memcpy(m, v380, v146, v357)
	mBase = m.M
	goto L50
L54:
	;
	goto L55
L55:
	;
	v394 = (v380 ^ v146) & int32(3)
	if base.Ui32(v380) < base.Ui32(v146) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v496 == int32(0) {
		goto L51
	} else {
		goto L92
	}
L57:
	;
	if base.Ui32(v474) <= base.Ui32(int32(3)) {
		v495 = v473
		v496 = v474
		v497 = v475
		goto L56
	} else {
		goto L88
	}
L58:
	;
	if v394 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	if v394 != 0 {
		v456 = v357
		goto L71
	} else {
		goto L72
	}
L61:
	;
	v495 = v146
	v496 = v357
	v497 = v380
	goto L56
L62:
	;
	goto L63
L63:
	;
	if v380&int32(3) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v473 = v146
	v474 = v357
	v475 = v380
	goto L57
L65:
	;
	goto L66
L66:
	;
	v401 = v146
	v402 = v357
	v403 = v380
	goto L67
L67:
	;
	if v402 == int32(0) {
		goto L51
	} else {
		goto L69
	}
L68:
	;
	v473 = v410
	v474 = v412
	v475 = v414
	goto L57
L69:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v407)
	v409 = int32(1)
	v410 = v401 + v409
	v412 = v402 - v409
	v414 = v403 + v409
	if v414&int32(3) != 0 {
		v401 = v410
		v402 = v412
		v403 = v414
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	if v456 == int32(0) {
		goto L51
	} else {
		goto L84
	}
L72:
	;
	if v384&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v421 = v357
	goto L76
L74:
	;
	v436 = v357
	goto L75
L75:
	;
	if base.Ui32(v436) <= base.Ui32(int32(3)) {
		v456 = v436
		goto L71
	} else {
		goto L80
	}
L76:
	;
	if v421 == int32(0) {
		goto L51
	} else {
		goto L78
	}
L77:
	;
	v436 = v427
	goto L75
L78:
	;
	v427 = v421 - int32(1)
	v428 = v380 + v427
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v427))))
	*(*uint8)(unsafe.Add(mBase, uint32(v428))) = uint8(v430)
	if v428&int32(3) != 0 {
		v421 = v427
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v443 = v436
	goto L81
L81:
	;
	v447 = v443 - int32(4)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v146+v447)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v447))) = v450
	if base.Ui32(int32(3)) < base.Ui32(v447) {
		v443 = v447
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v456 = v447
	goto L71
L83:
	;
	goto L82
L84:
	;
	v463 = v456
	goto L85
L85:
	;
	v467 = v463 - int32(1)
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v467))))
	*(*uint8)(unsafe.Add(mBase, uint32(v380+v467))) = uint8(v470)
	if v467 != 0 {
		v463 = v467
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L51
L87:
	;
	goto L86
L88:
	;
	v480 = v473
	v481 = v474
	v482 = v475
	goto L89
L89:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v484
	v486 = int32(4)
	v487 = v480 + v486
	v489 = v482 + v486
	v491 = v481 - v486
	if base.Ui32(int32(3)) < base.Ui32(v491) {
		v480 = v487
		v481 = v491
		v482 = v489
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v495 = v487
	v496 = v491
	v497 = v489
	goto L56
L91:
	;
	goto L90
L92:
	;
	v502 = v495
	v503 = v496
	v504 = v497
	goto L93
L93:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	*(*uint8)(unsafe.Add(mBase, uint32(v504))) = uint8(v506)
	v508 = int32(1)
	v513 = v503 - v508
	if v513 != 0 {
		v502 = v502 + v508
		v503 = v513
		v504 = v504 + v508
		goto L93
	} else {
		goto L95
	}
L94:
	;
	goto L51
L95:
	;
	goto L94
L96:
	;
	F_pfree(m, v28)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	m.G0 = v25 + int32(16)
	return v138
L99:
	;
	goto L98
L100:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(152127), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(494019), int32(845), int32(215288))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_strip(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(8)
		v18 = v13 + v17
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		if int32(0) < v20 {
			v24 = v20 & int32(3)
			v25 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v20) {
				v31 = v25
				v32 = v2
				v36 = v2
				for {
					v43 = v18 + v31<<(uint(int32(2))%32)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
					v45 = int32(1)
					v47 = int32(2047)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
					v67 = int32(base.Ui32(v44)>>(uint(v45)%32))&v47 + (int32(base.Ui32(v49)>>(uint(v45)%32))&v47 + v32 + int32(base.Ui32(v55)>>(uint(v45)%32))&v47 + int32(base.Ui32(v61)>>(uint(v45)%32))&v47)
					v68 = int32(4)
					v69 = v31 + v68
					v71 = v36 + v68
					if v71 != v20&int32(2147483644) {
						v31 = v69
						v32 = v67
						v36 = v71
						continue
					} else {
						break
					}
					break
				}
				v74 = v69
				v75 = v67
			} else {
				v74 = v25
				v75 = v2
			}
			if v24 != 0 {
				v85 = v74
				v86 = v75
				v92 = v2
				for {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v18+v85<<(uint(int32(2))%32))))
					v99 = int32(1)
					v103 = int32(base.Ui32(v98)>>(uint(v99)%32))&int32(2047) + v86
					v107 = v92 + v99
					if v107 != v24 {
						v85 = v85 + v99
						v86 = v103
						v92 = v107
						continue
					} else {
						break
					}
					break
				}
				v111 = v103
			} else {
				v111 = v75
			}
			v123 = v111 + int32(8)
		} else {
			v123 = v17
		}
		v135 = v123 + v20<<(uint(int32(2))%32)
		v136 = F_palloc0(m, v135)
		mBase = m.M
		v137 = m.ExcPending
		if v137 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v136))) = v135 << (uint(int32(2)) % 32)
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v141
			if int32(0) < v141 {
				v146 = v136 + int32(8)
				v152 = v141
				v153 = int32(0)
				v155 = v146 + v141<<(uint(int32(2))%32)
				for {
					v162 = int32(2)
					v166 = v153 << (uint(v162) % 32)
					v167 = v18 + v166
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
					v175 = int32(base.Ui32(v168)>>(uint(int32(1))%32)) & int32(2047)
					if v175 != 0 {
						v176 = F__emscripten_memcpy_bulkmem(m, v155, v18+v152<<(uint(v162)%32)+int32(base.Ui32(v168)>>(uint(int32(12))%32)), v175)
						mBase = m.M
						v177 = v176
					} else {
						v177 = v155
					}
					v178 = v166 + v146
					v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
					*(*int32)(unsafe.Add(mBase, uint32(v178))) = v179 & int32(-2)
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
					v185 = v183 & int32(4094)
					*(*int32)(unsafe.Add(mBase, uint32(v178))) = v185 | v179&int32(-4096)
					v190 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v178))) = (v177-(v146+v190<<(uint(int32(2))%32)))<<(uint(int32(12))%32) | v185
					v199 = int32(1)
					v205 = v153 + v199
					v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					if v205 < v206 {
						v152 = v206
						v153 = v205
						v155 = v177 + int32(base.Ui32(v183)>>(uint(v199)%32))&int32(2047)
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v219 != v13 {
				F_pfree(m, v13)
				mBase = m.M
				v222 = m.ExcPending
				if v222 != 0 {
					return int32(0)
				} else {
					return v136
				}
			} else {
				return v136
			}
		}
	}
}
