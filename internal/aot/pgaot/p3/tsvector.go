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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var __phi70 int32
	_ = __phi70
	var v71 int32
	_ = v71
	var __phi71 int32
	_ = __phi71
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
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
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
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
	var v490 int32
	_ = v490
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
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
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 <= v2 {
		v306 = v2
		v310 = v22
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L8
	} else {
		goto L100
	}
L2:
	;
	v326 = v306 + v310<<(uint(int32(2))%32) + int32(8)
	v327 = F_palloc0(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L8
	} else {
		goto L71
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v22 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v263 = int32(0)
	v266 = v263
	v269 = v263
	goto L64
L5:
	;
	v28 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v28)
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
	v32 = F_palloc(m, int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_pg_qsort(m, v25, v22, int32(16), int32(1155))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v32
	v37 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v32))) = uint16(v37)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v41 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v41) <= base.Ui32(v30) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = v41
	goto L12
L11:
	;
	v44 = v30
	goto L12
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+2)) = uint16(v44)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v250 = v37
	goto L4
L13:
	;
	v52 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v52)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+8)))
	v56 = F_palloc(m, int32(4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v56
	v59 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v59)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v62 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v62) <= base.Ui32(v54) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = v62
	goto L17
L16:
	;
	v65 = v54
	goto L17
L17:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+2)) = uint16(v65)
	__phi70 = v25
	__phi71 = v25 + int32(16)
	__phi72 = v25
	v70 = __phi70
	v71 = __phi71
	v72 = __phi72
	goto L18
L18:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+18)))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)))
	if v86 == v87 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v240 = (v224 - v25 + int32(16)) >> (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v240
	v242 = int32(0)
	if v240 <= v242 {
		v306 = v242
		v310 = v240
		goto L2
	} else {
		goto L63
	}
L20:
	;
	v231 = v71 + int32(16)
	if (v231-v25)>>(uint(int32(4))%32) < v22 {
		__phi70 = v224
		__phi71 = v231
		__phi72 = v71
		v70 = __phi70
		v71 = __phi71
		v72 = __phi72
		goto L18
	} else {
		goto L62
	}
L21:
	;
	F_pfree(m, v89)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L43
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v86 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+18)) = uint16(v86)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+28)) = v140
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+24)))
	v143 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+22)) = uint16(v143)
	v146 = F_palloc(m, int32(4))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L39
	}
L25:
	;
	if v135 == int32(0) {
		goto L21
	} else {
		goto L38
	}
L26:
	;
	v135 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v96 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v89
	v98 = v90
	v99 = v86
	v100 = v96
	goto L33
L30:
	;
	v123 = v90
	v127 = int32(0)
	goto L31
L31:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	v135 = v127 - v128
	goto L25
L32:
	;
	v123 = v118
	v127 = v120
	goto L31
L33:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if base.B2i32(v100 != v102)|base.B2i32(v102 == int32(0)) != 0 {
		v118 = v98
		v120 = v100
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v118 = v112
	v120 = int32(0)
	goto L32
L35:
	;
	v108 = v99 - int32(1)
	if v108 == int32(0) {
		v118 = v98
		v120 = v100
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v111 = int32(1)
	v112 = v98 + v111
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v113 != 0 {
		v97 = v97 + v111
		v98 = v112
		v99 = v108
		v100 = v113
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L24
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = v146
	v149 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v146))) = uint16(v149)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v152 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v152) <= base.Ui32(v142) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v155 = v152
	goto L42
L41:
	;
	v155 = v142
	goto L42
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v151)+2)) = uint16(v155)
	v224 = v70 + int32(16)
	goto L20
L43:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
	if base.Ui32(int32(254)) < base.Ui32(v162) {
		v224 = v70
		goto L20
	} else {
		goto L44
	}
L44:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161+v162<<(uint(int32(1))%32)))))
	if v168 == int32(_a_F_make_tsvector_0) {
		v224 = v70
		goto L20
	} else {
		goto L45
	}
L45:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+24)))
	if base.B2i32(base.Ui32(v171) <= base.Ui32(int32(_a_F_make_tsvector_0)))&base.B2i32(v171 == v168) != 0 {
		v224 = v70
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)))
	if base.Ui32(v176) <= base.Ui32(v162+int32(1)) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v215 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v192+v213<<(uint(v215)%32))+2)) = uint16(v214)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	v222 = v220 + v215
	*(*uint16)(unsafe.Add(mBase, uint32(v219))) = uint16(v222)
	v224 = v70
	goto L20
L48:
	;
	v181 = v176 << (uint(int32(1)) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)) = uint16(v181)
	v187 = F_repalloc(m, v161, v176<<(uint(int32(2))%32)&int32(_a_F_make_tsvector_1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L51
	}
L49:
	;
	v192 = v161
	v193 = v171
	v194 = v162
	goto L50
L50:
	;
	v196 = v194 & int32(_a_F_make_tsvector_2)
	if v196 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v187
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+24)))
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187))))
	v192 = v187
	v193 = v190
	v194 = v191
	goto L50
L52:
	;
	v199 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v199) <= base.Ui32(v193) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v204 = int32(_a_F_make_tsvector_0)
	if base.Ui32(v204) <= base.Ui32(v193) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v202 = v199
	goto L57
L56:
	;
	v202 = v193
	goto L57
L57:
	;
	v213 = int32(0)
	v214 = v202
	goto L47
L58:
	;
	v207 = v204
	goto L60
L59:
	;
	v207 = v193
	goto L60
L60:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192+v196<<(uint(int32(1))%32)))))
	if v207 == v211 {
		v224 = v70
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v213 = v196
	v214 = v207
	goto L47
L62:
	;
	goto L19
L63:
	;
	v250 = v240
	goto L4
L64:
	;
	v284 = v262 + v269<<(uint(int32(4))%32)
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+2)))
	v286 = v266 + v285
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+6)))
	if v287 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if int32(_a_F_make_tsvector_3) <= v299 {
		goto L1
	} else {
		goto L70
	}
L66:
	;
	v288 = int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292))))
	v299 = (v286+v288)&int32(-2) + v293<<(uint(v288)%32) + int32(2)
	goto L68
L67:
	;
	v299 = v286
	goto L68
L68:
	;
	v301 = v269 + int32(1)
	if v301 != v250 {
		v266 = v299
		v269 = v301
		goto L64
	} else {
		goto L69
	}
L69:
	;
	goto L65
L70:
	;
	v306 = v299
	v310 = v250
	goto L2
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v326 << (uint(int32(2)) % 32)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = v332
	if int32(0) < v332 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v337 = v327 + int32(8)
	v340 = v337 + v332<<(uint(int32(2))%32)
	v343 = int32(0)
	v345 = v337
	v354 = v2
	goto L75
L73:
	;
	goto L74
L74:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v592 != 0 {
		goto L96
	} else {
		goto L97
	}
L75:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v360 = int32(1)
	v363 = v354 << (uint(int32(4)) % 32)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363+v364)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v345))) = v359&v360 | v366<<(uint(v360)%32)&int32(4094) | v343<<(uint(int32(12))%32)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v377 = v376 + v363
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+2)))
	if v378 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L74
L77:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v377)+12))
	base.MemoryCopy(m, v343+v340, v380, v378)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v383 = v382 + v363
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v383)+2)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	F_pfree(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v388 = v343 + v384
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v390 = v389 + v363
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+6)))
	if v391 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v572 = v354 + int32(1)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v572 < v573 {
		v343 = v553
		v345 = v345 + int32(4)
		v354 = v572
		goto L75
	} else {
		goto L95
	}
L82:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392))))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v395 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v345))) = v394 | v395
	v401 = (v388 + v395) & int32(-2)
	*(*uint16)(unsafe.Add(mBase, uint32(v340+v401))) = uint16(v393)
	if v393 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	*(*int32)(unsafe.Add(mBase, uint32(v345))) = v548 & int32(-2)
	v553 = v388
	goto L81
L85:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v538+v363)+8))
	F_pfree(m, v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L8
	} else {
		goto L94
	}
L86:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v407 = int32(2)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v413 = int32(1)
	v424 = v337 + v406<<(uint(v407)%32) + (int32(base.Ui32(v410)>>(uint(int32(12))%32))+int32(base.Ui32(v410)>>(uint(v413)%32))&int32(2047)+v413)&int32(_a_F_make_tsvector_4) + v407
	v425 = int32(0)
	if v393 != v413 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v434 = v425
	v435 = int32(0)
	goto L90
L88:
	;
	v490 = v425
	goto L89
L89:
	;
	v507 = v490 << (uint(int32(1)) % 32)
	v508 = v424 + v507
	v509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508))))
	v510 = int32(_a_F_make_tsvector_0)
	v511 = v509 & v510
	*(*uint16)(unsafe.Add(mBase, uint32(v508))) = uint16(v511)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v513+v363)+8))
	v517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515+v507)+2)))
	v519 = v517 & v510
	*(*uint16)(unsafe.Add(mBase, uint32(v508))) = uint16(v519)
	goto L85
L90:
	;
	v450 = int32(1)
	v451 = v434 << (uint(v450) % 32)
	v452 = v424 + v451
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452))))
	v454 = int32(_a_F_make_tsvector_0)
	v455 = v453 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v452))) = uint16(v455)
	v457 = int32(2)
	v458 = v451 | v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v459+v363)+8))
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v458+v461))))
	v465 = v463 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v452))) = uint16(v465)
	v467 = v424 + v458
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467))))
	v470 = v468 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v467))) = uint16(v470)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v472+v363)+8))
	v476 = v434 + v457
	v480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v474+v476<<(uint(v450)%32)))))
	v482 = v480 & v454
	*(*uint16)(unsafe.Add(mBase, uint32(v467))) = uint16(v482)
	v485 = v435 + v457
	if v485 != v393&int32(_a_F_make_tsvector_5) {
		v434 = v476
		v435 = v485
		goto L90
	} else {
		goto L92
	}
L91:
	;
	if v393&int32(1) == int32(0) {
		goto L85
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v490 = v476
	goto L89
L94:
	;
	v553 = v401 + v393<<(uint(int32(1))%32) + int32(2)
	goto L81
L95:
	;
	goto L76
L96:
	;
	F_pfree(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L8
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	m.G0 = v20 + int32(16)
	return v327
L99:
	;
	goto L98
L100:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(_a_F_make_tsvector_6)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v299
	F_errmsg(m, int32(_a_F_make_tsvector_7), v20)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_make_tsvector_8), int32(194), int32(_a_F_make_tsvector_9))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
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
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L55
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v55 = v2
	v61 = v2
	goto L9
L7:
	;
	v129 = v2
	goto L8
L8:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v137 = F_palloc0(m, int32(base.Ui32(v134)>>(uint(int32(2))%32)))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L22
	}
L9:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v55))))
	if v73 == int32(1) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v129 = v106 & int32(255)
	goto L8
L11:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v48+v55<<(uint(int32(2))%32))))
	switch v80&int32(255) - int32(65) {
	case 0, 32:
		v105 = int32(8)
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
	v106 = v105 | v61
	v108 = v55 + int32(1)
	if v108 != v45 {
		v55 = v108
		v61 = v106
		goto L9
	} else {
		goto L21
	}
L13:
	;
	v105 = int32(4)
	goto L12
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v105 = int32(1)
	goto L12
L16:
	;
	v105 = int32(2)
	goto L12
L17:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = base.I32_extend8_s(v80)
	F_errmsg(m, int32(_a_F_tsvector_filter_0), v25)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_tsvector_filter_1), int32(869), int32(_a_F_tsvector_filter_2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v139
	v142 = v137 + int32(8)
	v145 = v142 + v139<<(uint(int32(2))%32)
	if int32(0) < v139 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v149 = v28 + int32(8)
	v154 = v139
	v161 = v2
	v165 = v2
	v172 = v2
	goto L26
L24:
	;
	v358 = v2
	v362 = v2
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v362
	if v139 != v362 {
		goto L45
	} else {
		goto L46
	}
L26:
	;
	v177 = v149 + v172<<(uint(int32(2))%32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v178&int32(1) == int32(0) {
		v332 = v161
		v336 = v165
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v358 = v332
	v362 = v336
	goto L25
L28:
	;
	v347 = v172 + int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v347 < v348 {
		v154 = v348
		v161 = v332
		v165 = v336
		v172 = v347
		goto L26
	} else {
		goto L44
	}
L29:
	;
	v186 = int32(1)
	v191 = int32(base.Ui32(v178)>>(uint(v186)%32))&int32(2047) + v186
	v197 = v149 + v154<<(uint(int32(2))%32) + (v191+int32(base.Ui32(v178)>>(uint(int32(12))%32)))&int32(_a_F_tsvector_filter_3)
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197))))
	if v198 == int32(0) {
		v332 = v161
		v336 = v165
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v204 = v145 + (v191+v161)&int32(-2)
	v205 = int32(2)
	v209 = int32(0)
	v212 = v209
	v216 = v209
	v222 = v198
	goto L31
L31:
	;
	v233 = int32(1)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197+v205+v216<<(uint(v233)%32)))))
	if int32(base.Ui32(v129)>>(uint(int32(base.Ui32(v236)>>(uint(int32(14))%32)))%32))&v233 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v249 == int32(0) {
		v332 = v161
		v336 = v165
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v242 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v204+v205+v212<<(uint(v242)%32)))) = uint16(v236)
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197))))
	v249 = v212 + v242
	v250 = v246
	goto L35
L34:
	;
	v249 = v212
	v250 = v222
	goto L35
L35:
	;
	v252 = v216 + int32(1)
	if base.Ui32(v252) < base.Ui32(v250&int32(_a_F_tsvector_filter_4)) {
		v212 = v249
		v216 = v252
		v222 = v250
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v260 = v142 + v165<<(uint(int32(2))%32)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v262 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v261 | v262
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v265&int32(4094) | v161<<(uint(int32(12))%32) | v262
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v278 = int32(base.Ui32(v274)>>(uint(v262)%32)) & int32(2047)
	if v278 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	base.MemoryCopy(m, v161+v145, v149+v35<<(uint(int32(2))%32)+int32(base.Ui32(v274)>>(uint(int32(12))%32)), v278)
	goto L40
L39:
	;
	goto L40
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v204))) = uint16(v249)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v286 = int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v297&v286 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v301 = int32(2)
	v304 = int32(1)
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142+v300<<(uint(v301)%32)+(int32(base.Ui32(v297)>>(uint(v304)%32))&int32(2047)+int32(base.Ui32(v297)>>(uint(int32(12))%32))+v304)&int32(_a_F_tsvector_filter_5)))))
	v322 = v316<<(uint(v304)%32) + v301
	goto L43
L42:
	;
	v322 = int32(2)
	goto L43
L43:
	;
	v332 = v322 + ((int32(base.Ui32(v285)>>(uint(v286)%32))&int32(2047)+v286)&int32(4094) + v161)
	v336 = v165 + v286
	goto L28
L44:
	;
	goto L27
L45:
	;
	if v358 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v381 = v139
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v358<<(uint(int32(2))%32) + v381<<(uint(int32(4))%32) + int32(32)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v388 != v28 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	base.MemoryCopy(m, v142+v362<<(uint(int32(2))%32), v145, v358)
	goto L50
L49:
	;
	goto L50
L50:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v381 = v380
	goto L47
L51:
	;
	F_pfree(m, v28)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	m.G0 = v25 + int32(16)
	return v137
L54:
	;
	goto L53
L55:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_tsvector_filter_6), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_tsvector_filter_1), int32(845), int32(_a_F_tsvector_filter_2))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
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
			return int32(0)
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
			v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v224 != v14 {
				F_pfree(m, v14)
				mBase = m.M
				v227 = m.ExcPending
				if v227 != 0 {
					return int32(0)
				} else {
					return v144
				}
			} else {
				return v144
			}
		}
	}
}
