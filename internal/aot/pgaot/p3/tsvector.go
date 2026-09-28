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
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_init_tsvector_parser[0]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v20*int32(28))+uint32(_c_F_init_tsvector_parser[1])))
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v75 int32
	_ = v75
	var __phi75 int32
	_ = __phi75
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
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
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 <= v2 {
		v315 = v2
		v320 = v22
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L9
	} else {
		goto L106
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L9
	} else {
		goto L102
	}
L3:
	;
	v335 = v315 + v320<<(uint(int32(2))%32) + int32(8)
	v336 = F_palloc0(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L9
	} else {
		goto L73
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v266 = int32(0)
	v269 = v266
	v270 = v266
	goto L65
L6:
	;
	v28 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v28)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
	v33 = F_palloc_mul(m, v28, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_pg_qsort(m, v25, v22, int32(16), int32(1270))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L14
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v33
	v38 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33))) = uint16(v38)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v42 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v42) <= base.Ui32(v30) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v42
	goto L13
L12:
	;
	v45 = v30
	goto L13
L13:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)) = uint16(v45)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v254 = v38
	goto L5
L14:
	;
	v53 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v53)
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
	v58 = F_palloc_mul(m, v53, v53)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v58
	v61 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v58))) = uint16(v61)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v64 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v64) <= base.Ui32(v55) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = v64
	goto L18
L17:
	;
	v67 = v55
	goto L18
L18:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)) = uint16(v67)
	__phi72 = v25
	__phi74 = v25
	__phi75 = v25 + int32(16)
	v72 = __phi72
	v74 = __phi74
	v75 = __phi75
	goto L19
L19:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+18)))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+2)))
	if v88 == v89 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v243 = (v227 - v25 + int32(16)) >> (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v243
	v245 = int32(0)
	if v243 <= v245 {
		v315 = v245
		v320 = v243
		goto L3
	} else {
		goto L64
	}
L21:
	;
	v234 = v75 + int32(16)
	if (v234-v25)>>(uint(int32(4))%32) < v22 {
		__phi72 = v227
		__phi74 = v75
		__phi75 = v234
		v72 = __phi72
		v74 = __phi74
		v75 = __phi75
		goto L19
	} else {
		goto L63
	}
L22:
	;
	F_pfree(m, v91)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L44
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	if v88 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+18)) = uint16(v88)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+28)) = v142
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+24)))
	v145 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+22)) = uint16(v145)
	v149 = F_palloc_mul(m, v145, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L40
	}
L26:
	;
	if v137 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L27:
	;
	v137 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v98 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v99 = v91
	v100 = v92
	v101 = v88
	v102 = v98
	goto L34
L31:
	;
	v125 = v92
	v129 = int32(0)
	goto L32
L32:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v137 = v129 - v130
	goto L26
L33:
	;
	v125 = v120
	v129 = v122
	goto L32
L34:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if base.B2i32(v102 != v104)|base.B2i32(v104 == int32(0)) != 0 {
		v120 = v100
		v122 = v102
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v120 = v114
	v122 = int32(0)
	goto L33
L36:
	;
	v110 = v101 - int32(1)
	if v110 == int32(0) {
		v120 = v100
		v122 = v102
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v113 = int32(1)
	v114 = v100 + v113
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	if v115 != 0 {
		v99 = v99 + v113
		v100 = v114
		v101 = v110
		v102 = v115
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	goto L25
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v149
	v152 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v149))) = uint16(v152)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v155 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v155) <= base.Ui32(v144) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v158 = v155
	goto L43
L42:
	;
	v158 = v144
	goto L43
L43:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+2)) = uint16(v158)
	v227 = v72 + int32(16)
	goto L21
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	if base.Ui32(int32(254)) < base.Ui32(v165) {
		v227 = v72
		goto L21
	} else {
		goto L45
	}
L45:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+v165<<(uint(int32(1))%32)))))
	if v171 == int32(_a_F_make_tsvector_0) {
		v227 = v72
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+24)))
	if base.B2i32(base.Ui32(v174) <= base.Ui32(int32(_a_F_make_tsvector_0)))&base.B2i32(v174 == v171) != 0 {
		v227 = v72
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)))
	if base.Ui32(v179) <= base.Ui32(v165+int32(1)) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v218 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v194+v216<<(uint(v218)%32))+2)) = uint16(v217)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	v225 = v223 + v218
	*(*uint16)(unsafe.Add(mBase, uint32(v222))) = uint16(v225)
	v227 = v72
	goto L21
L49:
	;
	v184 = v179 << (uint(int32(1)) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v184)
	v189 = F_repalloc_mul(m, v164, int32(2), v184&int32(_a_F_make_tsvector_1))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	v194 = v164
	v196 = v174
	v197 = v165
	goto L51
L51:
	;
	v199 = v197 & int32(_a_F_make_tsvector_2)
	if v199 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v189
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+24)))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189))))
	v194 = v189
	v196 = v192
	v197 = v193
	goto L51
L53:
	;
	v202 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v202) <= base.Ui32(v196) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v207 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v207) <= base.Ui32(v196) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v205 = v202
	goto L58
L57:
	;
	v205 = v196
	goto L58
L58:
	;
	v216 = int32(0)
	v217 = v205
	goto L48
L59:
	;
	v210 = v207
	goto L61
L60:
	;
	v210 = v196
	goto L61
L61:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194+v199<<(uint(int32(1))%32)))))
	if v210 == v214 {
		v227 = v72
		goto L21
	} else {
		goto L62
	}
L62:
	;
	v216 = v199
	v217 = v210
	goto L48
L63:
	;
	goto L20
L64:
	;
	v254 = v243
	goto L5
L65:
	;
	v287 = v265 + v270<<(uint(int32(4))%32)
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+2)))
	if base.Ui32((v288-int32(2048))&int32(_a_F_make_tsvector_2)) <= base.Ui32(int32(_a_F_make_tsvector_3)) {
		goto L2
	} else {
		goto L67
	}
L66:
	;
	if base.Ui32(int32(_a_F_make_tsvector_4)) <= base.Ui32(v308) {
		goto L1
	} else {
		goto L72
	}
L67:
	;
	v295 = v269 + v288
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+6)))
	if v296 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v297 = int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301))))
	v308 = (v295+v297)&int32(-2) + v302<<(uint(v297)%32) + int32(2)
	goto L70
L69:
	;
	v308 = v295
	goto L70
L70:
	;
	v310 = v270 + int32(1)
	if v310 != v254 {
		v269 = v308
		v270 = v310
		goto L65
	} else {
		goto L71
	}
L71:
	;
	goto L66
L72:
	;
	v315 = v308
	v320 = v254
	goto L3
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336))) = v335 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v336)+4)) = v341
	if int32(0) < v341 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v346 = v336 + int32(8)
	v349 = v346 + v341<<(uint(int32(2))%32)
	v352 = int32(0)
	v354 = v346
	v363 = v2
	goto L77
L75:
	;
	goto L76
L76:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v601 != 0 {
		goto L98
	} else {
		goto L99
	}
L77:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	v369 = int32(1)
	v372 = v363 << (uint(int32(4)) % 32)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372+v373)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v368&v369 | v375<<(uint(v369)%32)&int32(4094) | v352<<(uint(int32(12))%32)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v386 = v385 + v372
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386)+2)))
	if v387 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L76
L79:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	base.MemoryCopy(m, v352+v349, v389, v387)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v392 = v391 + v372
	v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+2)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	F_pfree(m, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	v397 = v352 + v393
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = v398 + v372
	v400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399)+6)))
	if v400 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v581 = v363 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v581 < v582 {
		v352 = v562
		v354 = v354 + int32(4)
		v363 = v581
		goto L77
	} else {
		goto L97
	}
L84:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401))))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	v404 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v403 | v404
	v410 = (v397 + v404) & int32(-2)
	*(*uint16)(unsafe.Add(mBase, uint32(v349+v410))) = uint16(v402)
	if v402 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v557 & int32(-2)
	v562 = v397
	goto L83
L87:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v547+v372)+8))
	F_pfree(m, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L9
	} else {
		goto L96
	}
L88:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	v416 = int32(2)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	v422 = int32(1)
	v433 = v346 + v415<<(uint(v416)%32) + (int32(base.Ui32(v419)>>(uint(int32(12))%32))+int32(base.Ui32(v419)>>(uint(v422)%32))&int32(2047)+v422)&int32(_a_F_make_tsvector_5) + v416
	v434 = int32(0)
	if v402 != v422 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v443 = v434
	v448 = int32(0)
	goto L92
L90:
	;
	v499 = v434
	goto L91
L91:
	;
	v516 = v499 << (uint(int32(1)) % 32)
	v517 = v433 + v516
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517))))
	v519 = int32(_a_F_make_tsvector_0)
	v520 = v518 & v519
	*(*uint16)(unsafe.Add(mBase, uint32(v517))) = uint16(v520)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v522+v372)+8))
	v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v524+v516)+2)))
	v528 = v526 & v519
	*(*uint16)(unsafe.Add(mBase, uint32(v517))) = uint16(v528)
	goto L87
L92:
	;
	v459 = int32(1)
	v460 = v443 << (uint(v459) % 32)
	v461 = v433 + v460
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v461))))
	v463 = int32(_a_F_make_tsvector_0)
	v464 = v462 & v463
	*(*uint16)(unsafe.Add(mBase, uint32(v461))) = uint16(v464)
	v466 = int32(2)
	v467 = v460 | v466
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v468+v372)+8))
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467+v470))))
	v474 = v472 & v463
	*(*uint16)(unsafe.Add(mBase, uint32(v461))) = uint16(v474)
	v476 = v433 + v467
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v476))))
	v479 = v477 & v463
	*(*uint16)(unsafe.Add(mBase, uint32(v476))) = uint16(v479)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v481+v372)+8))
	v485 = v443 + v466
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v483+v485<<(uint(v459)%32)))))
	v491 = v489 & v463
	*(*uint16)(unsafe.Add(mBase, uint32(v476))) = uint16(v491)
	v494 = v448 + v466
	if v494 != v402&int32(_a_F_make_tsvector_1) {
		v443 = v485
		v448 = v494
		goto L92
	} else {
		goto L94
	}
L93:
	;
	if v402&int32(1) == int32(0) {
		goto L87
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v499 = v485
	goto L91
L96:
	;
	v562 = v410 + v402<<(uint(int32(1))%32) + int32(2)
	goto L83
L97:
	;
	goto L78
L98:
	;
	F_pfree(m, v601)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L9
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	m.G0 = v20 + int32(32)
	return v336
L101:
	;
	goto L100
L102:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(2047)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v288
	F_errmsg(m, int32(_a_F_make_tsvector_6), v20)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_make_tsvector_7), int32(194), int32(_a_F_make_tsvector_8))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L9
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
	F_errcode(m, int32(261))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_make_tsvector_9)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v308
	F_errmsg(m, int32(_a_F_make_tsvector_10), v20+int32(16))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_make_tsvector_7), int32(207), int32(_a_F_make_tsvector_8))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_filter(m *base.Module, l0 int32) int64 {
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L47
	}
L6:
	;
	v50 = v2
	v54 = v2
	goto L9
L7:
	;
	v97 = v2
	goto L8
L8:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v116 = F_palloc0(m, int32(base.Ui32(v113)>>(uint(int32(2))%32)))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v54))))
	if v72 == int32(1) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v97 = v84 & int32(255)
	goto L8
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76+v54<<(uint(int32(3))%32)))))
	v81 = F_parse_weight(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v84 = v50 | int32(1)<<(uint(v81)%32)
	v86 = v54 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v86 < v87 {
		v50 = v84
		v54 = v86
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v118
	v121 = v116 + int32(8)
	v124 = v121 + v118<<(uint(int32(2))%32)
	if int32(0) < v118 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v128 = v28 + int32(8)
	v133 = v118
	v137 = v2
	v142 = v2
	v148 = v2
	goto L18
L16:
	;
	v332 = v2
	v337 = v2
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = v337
	if v118 != v337 {
		goto L37
	} else {
		goto L38
	}
L18:
	;
	v156 = v128 + v148<<(uint(int32(2))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157&int32(1) == int32(0) {
		v306 = v137
		v311 = v142
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v332 = v306
	v337 = v311
	goto L17
L20:
	;
	v324 = v148 + int32(1)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v324 < v325 {
		v133 = v325
		v137 = v306
		v142 = v311
		v148 = v324
		goto L18
	} else {
		goto L36
	}
L21:
	;
	v165 = int32(1)
	v170 = int32(base.Ui32(v157)>>(uint(v165)%32))&int32(2047) + v165
	v176 = v128 + v133<<(uint(int32(2))%32) + (v170+int32(base.Ui32(v157)>>(uint(int32(12))%32)))&int32(_a_F_tsvector_filter_0)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176))))
	if v177 == int32(0) {
		v306 = v137
		v311 = v142
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v183 = v124 + (v170+v137)&int32(-2)
	v184 = int32(2)
	v188 = int32(0)
	v191 = v188
	v192 = v188
	v202 = v177
	goto L23
L23:
	;
	v212 = int32(1)
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v184+v192<<(uint(v212)%32)))))
	if int32(base.Ui32(v97)>>(uint(int32(base.Ui32(v215)>>(uint(int32(14))%32)))%32))&v212 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v228 == int32(0) {
		v306 = v137
		v311 = v142
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v221 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v183+v184+v191<<(uint(v221)%32)))) = uint16(v215)
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176))))
	v228 = v191 + v221
	v229 = v225
	goto L27
L26:
	;
	v228 = v191
	v229 = v202
	goto L27
L27:
	;
	v231 = v192 + int32(1)
	if base.Ui32(v231) < base.Ui32(v229) {
		v191 = v228
		v192 = v231
		v202 = v229
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v237 = v121 + v142<<(uint(int32(2))%32)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v239 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v238 | v239
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v242&int32(4094) | v137<<(uint(int32(12))%32) | v239
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v255 = int32(base.Ui32(v251)>>(uint(v239)%32)) & int32(2047)
	if v255 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	base.MemoryCopy(m, v137+v124, v128+v35<<(uint(int32(2))%32)+int32(base.Ui32(v251)>>(uint(int32(12))%32)), v255)
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v183))) = uint16(v228)
	v262 = int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v264&v262 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v268 = int32(2)
	v271 = int32(1)
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121+v267<<(uint(v268)%32)+(int32(base.Ui32(v264)>>(uint(v271)%32))&int32(2047)+int32(base.Ui32(v264)>>(uint(int32(12))%32))+v271)&int32(_a_F_tsvector_filter_1)))))
	v289 = v283<<(uint(v271)%32) + v268
	goto L35
L34:
	;
	v289 = int32(2)
	goto L35
L35:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v291 = int32(1)
	v306 = v289 + ((int32(base.Ui32(v290)>>(uint(v291)%32))&int32(2047)+v291)&int32(4094) + v137)
	v311 = v142 + v262
	goto L20
L36:
	;
	goto L19
L37:
	;
	if v332 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v358 = v118
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v332<<(uint(int32(2))%32) + v358<<(uint(int32(4))%32) + int32(32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v365 != v28 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	base.MemoryCopy(m, v121+v337<<(uint(int32(2))%32), v124, v332)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v358 = v357
	goto L39
L43:
	;
	F_pfree(m, v28)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	m.G0 = v25 + int32(16)
	return base.I64_extend_i32_u(v116)
L46:
	;
	goto L45
L47:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_tsvector_filter_2), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_tsvector_filter_3), int32(858), int32(_a_F_tsvector_filter_4))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_strip(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(8)
		v19 = v14 + v18
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if int32(0) < v21 {
			v25 = v21 & int32(3)
			v26 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v21) {
				v32 = v26
				v35 = v2
				v40 = v2
				for {
					v45 = v19 + v32<<(uint(int32(2))%32)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
					v47 = int32(1)
					v49 = int32(2047)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					v69 = int32(base.Ui32(v46)>>(uint(v47)%32))&v49 + (int32(base.Ui32(v51)>>(uint(v47)%32))&v49 + v35 + int32(base.Ui32(v57)>>(uint(v47)%32))&v49 + int32(base.Ui32(v63)>>(uint(v47)%32))&v49)
					v70 = int32(4)
					v71 = v32 + v70
					v73 = v40 + v70
					if v73 != v21&int32(2147483644) {
						v32 = v71
						v35 = v69
						v40 = v73
						continue
					} else {
						break
					}
					break
				}
				if v25 == int32(0) {
					v119 = v69
				} else {
					v78 = v71
					v81 = v69
					v90 = v78
					v93 = v81
					v99 = v2
					for {
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v19+v90<<(uint(int32(2))%32))))
						v105 = int32(1)
						v109 = int32(base.Ui32(v104)>>(uint(v105)%32))&int32(2047) + v93
						v113 = v99 + v105
						if v113 != v25 {
							v90 = v90 + v105
							v93 = v109
							v99 = v113
							continue
						} else {
							break
						}
						break
					}
					v119 = v109
				}
			} else {
				v78 = v26
				v81 = v2
				v90 = v78
				v93 = v81
				v99 = v2
				for {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v19+v90<<(uint(int32(2))%32))))
					v105 = int32(1)
					v109 = int32(base.Ui32(v104)>>(uint(v105)%32))&int32(2047) + v93
					v113 = v99 + v105
					if v113 != v25 {
						v90 = v90 + v105
						v93 = v109
						v99 = v113
						continue
					} else {
						break
					}
					break
				}
				v119 = v109
			}
			v130 = v119 + int32(8)
		} else {
			v130 = v18
		}
		v143 = v130 + v21<<(uint(int32(2))%32)
		v144 = F_palloc0(m, v143)
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v144))) = v143 << (uint(int32(2)) % 32)
			v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v149
			if int32(0) < v149 {
				v154 = v144 + int32(8)
				v160 = v154 + v149<<(uint(int32(2))%32)
				v161 = v149
				v163 = int32(0)
				for {
					v172 = v163 << (uint(int32(2)) % 32)
					v173 = v19 + v172
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
					v178 = int32(base.Ui32(v174)>>(uint(int32(1))%32)) & int32(2047)
					if v178 != 0 {
						base.MemoryCopy(m, v160, v19+v161<<(uint(int32(2))%32)+int32(base.Ui32(v174)>>(uint(int32(12))%32)), v178)
					} else {
					}
					v186 = v154 + v172
					v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
					*(*int32)(unsafe.Add(mBase, uint32(v186))) = v187 & int32(-2)
					v191 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v186))) = v191&int32(4094) | (v160-(v154+v194<<(uint(int32(2))%32)))<<(uint(int32(12))%32)
					v203 = int32(1)
					v209 = v163 + v203
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					if v209 < v210 {
						v160 = v160 + int32(base.Ui32(v191)>>(uint(v203)%32))&int32(2047)
						v161 = v210
						v163 = v209
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v224 != v14 {
				F_pfree(m, v14)
				mBase = m.M
				v227 = m.ExcPending
				if v227 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v144)
				}
			} else {
				return base.I64_extend_i32_u(v144)
			}
		}
	}
}
