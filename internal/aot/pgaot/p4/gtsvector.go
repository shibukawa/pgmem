package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtsvector_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v588 int64
	_ = v588
	v2 = int32(0)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = base.I32_wrap_i64(v17)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 == v2 {
		v37 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v37&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v24 == int32(0) {
		v37 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v27 != int32(7) {
		v37 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v30 != int32(17) {
		v37 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+32)))
	v37 = v33 ^ int32(1)
	goto L2
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = F_get_fn_opclass_options(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v46 = int32(124)
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+18)))
	if v48 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int64(0)
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v46 = v45
	goto L9
L12:
	;
	return v588
L13:
	;
	v588 = base.I64_extend_i32_u(v556)
	goto L12
L14:
	;
	v544 = F_palloc(m, int32(24))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L83
	}
L15:
	;
	v51 = F_pg_detoast_datum(m, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v471&int32(6) != int32(2) {
		goto L72
	} else {
		goto L73
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v57 = v53<<(uint(int32(2))%32) + int32(8)
	v58 = F_palloc(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v57 << (uint(int32(2)) % 32)
	v66 = v58 + int32(8)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v69 = v51 + int32(8)
	v76 = v66
	v80 = v67
	v82 = v69
	goto L23
L21:
	;
	v235 = int32(0)
	goto L22
L22:
	;
	F_pg_qsort(m, v66, v235, int32(4), int32(1716))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L37
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v91 = int32(base.Ui32(v89) >> (uint(int32(1)) % 32))
	v93 = v91 & int32(2047)
	if v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v235 = v217
	goto L22
L25:
	;
	v96 = v69 + v67<<(uint(int32(2))%32) + int32(base.Ui32(v89)>>(uint(int32(12))%32))
	v97 = int32(-1)
	if v93 != int32(1) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v209 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v209
	v211 = int32(4)
	v216 = v80 - int32(1)
	if v216 != 0 {
		v76 = v76 + v211
		v80 = v216
		v82 = v82 + v211
		goto L23
	} else {
		goto L36
	}
L28:
	;
	v209 = v176 ^ int32(-1)
	goto L27
L29:
	;
	v105 = v96
	v107 = v97
	v115 = int32(0)
	goto L32
L30:
	;
	v148 = v96
	v150 = v97
	goto L31
L31:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v170 = *(*int32)(unsafe.Add(mBase, uint32((v164^int32(base.Ui32(v150)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_gtsvector_compress[0])))
	v176 = v170 ^ v150<<(uint(int32(8))%32)
	goto L28
L32:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v123 = int32(24)
	v126 = int32(2)
	v128 = *(*int32)(unsafe.Add(mBase, uint32((v122^int32(base.Ui32(v107)>>(uint(v123)%32)))<<(uint(v126)%32))+uint32(_c_F_gtsvector_compress[0])))
	v129 = int32(8)
	v131 = v128 ^ v107<<(uint(v129)%32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32((v121^int32(base.Ui32(v131)>>(uint(v123)%32)))<<(uint(v126)%32))+uint32(_c_F_gtsvector_compress[0])))
	v140 = v137 ^ v131<<(uint(v129)%32)
	v142 = v105 + v126
	v144 = v115 + v126
	if v144 != v91&int32(2046) {
		v105 = v142
		v107 = v140
		v115 = v144
		goto L32
	} else {
		goto L34
	}
L33:
	;
	if v91&int32(1) == int32(0) {
		v176 = v140
		goto L28
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v148 = v142
	v150 = v140
	goto L31
L36:
	;
	goto L24
L37:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v240) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	if base.Ui32(v311) < base.Ui32(int32(2044)) {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v305 = v282<<(uint(int32(2))%32) + int32(8)
	v306 = F_repalloc(m, v58, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L10
	} else {
		goto L50
	}
L40:
	;
	v244 = int32(1)
	v248 = v2
	goto L43
L41:
	;
	goto L42
L42:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v311 = v301
	v312 = v58
	goto L38
L43:
	;
	v260 = int32(2)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v66+v244<<(uint(v260)%32))))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v66+v248<<(uint(v260)%32))))
	if v263 == v267 {
		v277 = v248
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v282 = v277 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v282 != v283 {
		goto L39
	} else {
		goto L49
	}
L45:
	;
	v279 = v244 + int32(1)
	if v279 != v240 {
		v244 = v279
		v248 = v277
		goto L43
	} else {
		goto L48
	}
L46:
	;
	v270 = v248 + int32(1)
	if v244 == v270 {
		v277 = v244
		goto L45
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66+v270<<(uint(int32(2))%32)))) = v263
	v277 = v270
	goto L45
L48:
	;
	goto L44
L49:
	;
	goto L42
L50:
	;
	v309 = v305 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v309
	v311 = v309
	v312 = v306
	goto L38
L51:
	;
	v531 = v312
	goto L14
L52:
	;
	goto L53
L53:
	;
	v330 = v46 + int32(8)
	v331 = F_palloc(m, v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	v333 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v331)+4)) = v333
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v330 << (uint(v333) % 32)
	v338 = int32(8)
	v339 = v331 + v338
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v346 = int32(base.Ui32(int32(base.Ui32(v340)>>(uint(v333)%32))-v338) >> (uint(v333) % 32))
	if base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v46))|v46&int32(3) == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v346 == int32(0) {
		v531 = v331
		goto L14
	} else {
		goto L64
	}
L56:
	;
	if v46 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v368 = v46
	goto L58
L58:
	;
	if v368 == int32(0) {
		goto L55
	} else {
		goto L63
	}
L59:
	;
	v358 = v46 + v339
	v360 = v331 + int32(12)
	if base.Ui32(v360) < base.Ui32(v358) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v362 = v358
	goto L62
L61:
	;
	v362 = v360
	goto L62
L62:
	;
	v368 = (v339^int32(-1)+v362)&int32(-4) + int32(4)
	goto L58
L63:
	;
	base.MemoryFill(m, v339, int32(0), v368)
	goto L55
L64:
	;
	v379 = v312 + int32(8)
	v381 = v46 << (uint(int32(3)) % 32)
	v382 = int32(0)
	if v346 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v390 = v382
	v396 = int32(0)
	goto L68
L66:
	;
	v440 = v382
	goto L67
L67:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v379+v440<<(uint(int32(2))%32))))
	v460 = base.I32_rem_u_s(v459, v381)
	v463 = v339 + int32(base.Ui32(v460)>>(uint(int32(3))%32))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	v469 = v464 | int32(1)<<(uint(v460&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v469)
	v531 = v331
	goto L14
L68:
	;
	v406 = int32(2)
	v408 = v379 + v390<<(uint(v406)%32)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v410 = base.I32_rem_u_s(v409, v381)
	v411 = int32(3)
	v413 = v339 + int32(base.Ui32(v410)>>(uint(v411)%32))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v415 = int32(1)
	v416 = int32(7)
	v419 = v414 | v415<<(uint(v410&v416)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v413))) = uint8(v419)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v422 = base.I32_rem_u_s(v421, v381)
	v425 = v339 + int32(base.Ui32(v422)>>(uint(v411)%32))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425))))
	v431 = v426 | v415<<(uint(v422&v416)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v431)
	v434 = v390 + v406
	v436 = v396 + v406
	if v436 != v346&int32(1073741822) {
		v390 = v434
		v396 = v436
		goto L68
	} else {
		goto L70
	}
L69:
	;
	if v346&int32(1) == int32(0) {
		v531 = v331
		goto L14
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	v440 = v434
	goto L67
L72:
	;
	v556 = v18
	goto L13
L73:
	;
	goto L74
L74:
	;
	v476 = int32(0)
	if v476 < v46 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v483 = v476
	goto L78
L76:
	;
	goto L77
L77:
	;
	v523 = F_palloc(m, int32(8))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L82
	}
L78:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+(v47+int32(8))))))
	if v500 != int32(255) {
		v588 = v17 & int64(4294967295)
		goto L12
	} else {
		goto L80
	}
L79:
	;
	goto L77
L80:
	;
	v504 = v483 + int32(1)
	if v504 != v46 {
		v483 = v504
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v523))) = int64(25769803808)
	v531 = v523
	goto L14
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v544))) = base.I64_extend_i32_u(v531)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v544)+8)) = v548
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v544)+12)) = v550
	v552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	v553 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v544)+18)) = uint8(v553)
	*(*uint16)(unsafe.Add(mBase, uint32(v544)+16)) = uint16(v552)
	v556 = v544
	goto L13
}
func F_gtsvector_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v13)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v16 == v2 {
		v54 = v2
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v54)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v19&int32(2) != 0 {
			if v19&int32(4) != 0 {
				v54 = int32(1)
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_u(v54)
			} else {
				v29 = F_TS_execute(m, v11+int32(8), v10, int32(2), int32(1717))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v54 = v29
					m.G0 = v7 + int32(16)
					return base.I64_extend_i32_u(v54)
				}
			}
		} else {
			v33 = int32(8)
			v34 = v10 + v33
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v37 = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v34 + (int32(base.Ui32(v36)>>(uint(v37)%32))-v33)&int32(-4)
			v51 = F_TS_execute(m, v11+v33, v7+v33, v37, int32(1718))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				v54 = v51
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_u(v54)
			}
		}
	}
}
func F_gtsvector_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v137 int32
	_ = v137
	v2 = int32(0)
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v2 {
		v30 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v30&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v17 == int32(0) {
		v30 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v20 != int32(7) {
		v30 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v23 != int32(17) {
		v30 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	v30 = v26 ^ int32(1)
	goto L2
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_get_fn_opclass_options(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v39 = int32(124)
	goto L9
L9:
	;
	v42 = base.I32_wrap_i64(v9)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v43&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int64(0)
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v39 = v38
	goto L9
L12:
	;
	return v9 & int64(4294967295)
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v137)
	goto L12
L14:
	;
	v137 = int32(0)
	goto L13
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v47 = int32(4)
	v48 = v46 & v47
	if v43&v47 != 0 {
		v137 = int32(base.Ui32(v48) >> (uint(int32(2)) % 32))
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v79 = int32(2)
	v81 = int32(8)
	v84 = int32(base.Ui32(int32(base.Ui32(v78)>>(uint(v79)%32))-v81) >> (uint(v79) % 32))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v84 != int32(base.Ui32(int32(base.Ui32(v85)>>(uint(v79)%32))-v81)>>(uint(v79)%32)) {
		goto L14
	} else {
		goto L25
	}
L18:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v53 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v53)
	v55 = int32(0)
	if v39 <= v55 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v58 = int32(8)
	v62 = v55
	goto L21
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+(v11+v58)))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+(v10+v58)))))
	if v71 != v73 {
		goto L14
	} else {
		goto L23
	}
L22:
	;
	goto L12
L23:
	;
	v76 = v62 + int32(1)
	if v39 != v76 {
		v62 = v76
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v93)
	if v84 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v97 = int32(8)
	v102 = int32(0)
	goto L27
L27:
	;
	v111 = v102 << (uint(int32(2)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v11+v97+v111)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v10+v97+v111)))
	if v113 != v115 {
		goto L14
	} else {
		goto L29
	}
L28:
	;
	goto L12
L29:
	;
	v118 = v102 + int32(1)
	if v84 != v118 {
		v102 = v118
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
