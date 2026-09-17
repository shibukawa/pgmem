package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtsvector_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
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
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
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
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	v2 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 == v2 {
		v35 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v35&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v22 == int32(0) {
		v35 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v25 != int32(7) {
		v35 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v28 != int32(17) {
		v35 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	v35 = v31 ^ int32(1)
	goto L2
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = F_get_fn_opclass_options(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v44 = int32(124)
	goto L9
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+14)))
	if v46 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v44 = v43
	goto L9
L12:
	;
	v528 = F_palloc(m, int32(16))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L10
	} else {
		goto L80
	}
L13:
	;
	v49 = F_pg_detoast_datum(m, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v458&int32(6) != int32(2) {
		goto L68
	} else {
		goto L69
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v55 = v51<<(uint(int32(2))%32) + int32(8)
	v56 = F_palloc(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v55 << (uint(int32(2)) % 32)
	v64 = v56 + int32(8)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v65 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = v49 + int32(8)
	v74 = v64
	v78 = v65
	v81 = v67
	goto L21
L19:
	;
	v227 = int32(0)
	goto L20
L20:
	;
	F_pg_qsort(m, v64, v227, int32(4), int32(1503))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L10
	} else {
		goto L35
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v88 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
	v90 = v88 & int32(2047)
	if v90 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v227 = v210
	goto L20
L23:
	;
	v93 = v67 + v65<<(uint(int32(2))%32) + int32(base.Ui32(v86)>>(uint(int32(12))%32))
	v94 = int32(-1)
	if v90 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v202 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v202
	v204 = int32(4)
	v209 = v78 - int32(1)
	if v209 != 0 {
		v74 = v74 + v204
		v78 = v209
		v81 = v81 + v204
		goto L21
	} else {
		goto L34
	}
L26:
	;
	v202 = v171 ^ int32(-1)
	goto L25
L27:
	;
	v102 = v93
	v104 = v94
	v113 = int32(0)
	goto L30
L28:
	;
	v144 = v93
	v146 = v94
	goto L29
L29:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v165 = *(*int32)(unsafe.Add(mBase, uint32((v159^int32(base.Ui32(v146)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_gtsvector_compress[0])))
	v171 = v165 ^ v146<<(uint(int32(8))%32)
	goto L26
L30:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v119 = int32(24)
	v122 = int32(2)
	v124 = *(*int32)(unsafe.Add(mBase, uint32((v118^int32(base.Ui32(v104)>>(uint(v119)%32)))<<(uint(v122)%32))+uint32(_c_F_gtsvector_compress[0])))
	v125 = int32(8)
	v127 = v124 ^ v104<<(uint(v125)%32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32((v117^int32(base.Ui32(v127)>>(uint(v119)%32)))<<(uint(v122)%32))+uint32(_c_F_gtsvector_compress[0])))
	v136 = v133 ^ v127<<(uint(v125)%32)
	v138 = v102 + v122
	v140 = v113 + v122
	if v140 != v88&int32(2046) {
		v102 = v138
		v104 = v136
		v113 = v140
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v88&int32(1) == int32(0) {
		v171 = v136
		goto L26
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v144 = v138
	v146 = v136
	goto L29
L34:
	;
	goto L22
L35:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if base.Ui32(v232) < base.Ui32(int32(2)) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if base.Ui32(v301) < base.Ui32(int32(2044)) {
		goto L47
	} else {
		goto L48
	}
L37:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v301 = v300
	v302 = v56
	goto L36
L38:
	;
	v236 = int32(1)
	v240 = v2
	goto L39
L39:
	;
	v251 = int32(2)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v64+v236<<(uint(v251)%32))))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v64+v240<<(uint(v251)%32))))
	if v254 == v258 {
		v268 = v240
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v273 = v268 + int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v273 == v274 {
		goto L37
	} else {
		goto L45
	}
L41:
	;
	v270 = v236 + int32(1)
	if v270 != v232 {
		v236 = v270
		v240 = v268
		goto L39
	} else {
		goto L44
	}
L42:
	;
	v261 = v240 + int32(1)
	if v236 == v261 {
		v268 = v236
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64+v261<<(uint(int32(2))%32)))) = v254
	v268 = v261
	goto L41
L44:
	;
	goto L40
L45:
	;
	v279 = v273<<(uint(int32(2))%32) + int32(8)
	v280 = F_repalloc(m, v56, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v283 = v279 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = v283
	v301 = v283
	v302 = v280
	goto L36
L47:
	;
	v516 = v302
	goto L12
L48:
	;
	goto L49
L49:
	;
	v319 = v44 + int32(8)
	v320 = F_palloc(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v322 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v319 << (uint(v322) % 32)
	v327 = int32(8)
	v328 = v320 + v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v335 = int32(base.Ui32(int32(base.Ui32(v329)>>(uint(v322)%32))-v327) >> (uint(v322) % 32))
	if base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v44))|v44&int32(3) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v335 == int32(0) {
		v516 = v320
		goto L12
	} else {
		goto L60
	}
L52:
	;
	if v44 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	v357 = v44
	goto L54
L54:
	;
	if v357 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L55:
	;
	v347 = v44 + v328
	v349 = v320 + int32(12)
	if base.Ui32(v349) < base.Ui32(v347) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v351 = v347
	goto L58
L57:
	;
	v351 = v349
	goto L58
L58:
	;
	v357 = (v328^int32(-1)+v351)&int32(-4) + int32(4)
	goto L54
L59:
	;
	base.MemoryFill(m, v328, int32(0), v357)
	goto L51
L60:
	;
	v368 = v302 + int32(8)
	v370 = v44 << (uint(int32(3)) % 32)
	v371 = int32(0)
	if v335 != int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v379 = v371
	v385 = int32(0)
	goto L64
L62:
	;
	v428 = v371
	goto L63
L63:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v368+v428<<(uint(int32(2))%32))))
	v447 = base.I32_rem_u_s(v446, v370)
	v450 = v328 + int32(base.Ui32(v447)>>(uint(int32(3))%32))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	v456 = v451 | int32(1)<<(uint(v447&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v450))) = uint8(v456)
	v516 = v320
	goto L12
L64:
	;
	v394 = int32(2)
	v396 = v368 + v379<<(uint(v394)%32)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v398 = base.I32_rem_u_s(v397, v370)
	v399 = int32(3)
	v401 = v328 + int32(base.Ui32(v398)>>(uint(v399)%32))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	v403 = int32(1)
	v404 = int32(7)
	v407 = v402 | v403<<(uint(v398&v404)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v407)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	v410 = base.I32_rem_u_s(v409, v370)
	v413 = v328 + int32(base.Ui32(v410)>>(uint(v399)%32))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v419 = v414 | v403<<(uint(v410&v404)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v413))) = uint8(v419)
	v422 = v379 + v394
	v424 = v385 + v394
	if v424 != v335&int32(1073741822) {
		v379 = v422
		v385 = v424
		goto L64
	} else {
		goto L66
	}
L65:
	;
	if v335&int32(1) == int32(0) {
		v516 = v320
		goto L12
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v428 = v422
	goto L63
L68:
	;
	return v16
L69:
	;
	goto L70
L70:
	;
	v464 = int32(0)
	if v44 <= v464 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v508 = F_palloc(m, int32(8))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L10
	} else {
		goto L79
	}
L72:
	;
	v469 = v464
	goto L73
L73:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469+(v45+int32(8))))))
	if v485 == int32(255) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	return v16
L75:
	;
	v489 = v469 + int32(1)
	if v44 != v489 {
		v469 = v489
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L71
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v508))) = int64(25769803808)
	v516 = v508
	goto L12
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v516
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+4)) = v531
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+8)) = v533
	v535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)))
	v536 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v528)+14)) = uint8(v536)
	*(*uint16)(unsafe.Add(mBase, uint32(v528)+12)) = uint16(v535)
	return v528
}
func F_gtsvector_consistent(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v13)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v16 == v2 {
		v54 = v2
		m.G0 = v7 + int32(16)
		return v54
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v19&int32(2) != 0 {
			if v19&int32(4) != 0 {
				v54 = int32(1)
				m.G0 = v7 + int32(16)
				return v54
			} else {
				v29 = F_TS_execute(m, v11+int32(8), v10, int32(2), int32(1504))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v54 = v29
					m.G0 = v7 + int32(16)
					return v54
				}
			}
		} else {
			v33 = int32(8)
			v34 = v10 + v33
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v37 = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v34 + (int32(base.Ui32(v36)>>(uint(v37)%32))-v33)&int32(-4)
			v51 = F_TS_execute(m, v11+v33, v7+v33, v37, int32(1505))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v54 = v51
				m.G0 = v7 + int32(16)
				return v54
			}
		}
	}
}
func F_gtsvector_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 == v2 {
		v28 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v28&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v15 == int32(0) {
		v28 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v18 != int32(7) {
		v28 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v21 != int32(17) {
		v28 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
	v28 = v24 ^ int32(1)
	goto L2
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = F_get_fn_opclass_options(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v37 = int32(124)
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v38&int32(2) != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(0)
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v37 = v36
	goto L9
L12:
	;
	return v7
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v124)
	goto L12
L14:
	;
	v124 = int32(0)
	goto L13
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v42 = int32(4)
	v43 = v41 & v42
	if v38&v42 != 0 {
		v124 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v72 = int32(2)
	v74 = int32(8)
	v77 = int32(base.Ui32(int32(base.Ui32(v71)>>(uint(v72)%32))-v74) >> (uint(v72) % 32))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v77 != int32(base.Ui32(int32(base.Ui32(v78)>>(uint(v72)%32))-v74)>>(uint(v72)%32)) {
		goto L14
	} else {
		goto L25
	}
L18:
	;
	if v43 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v48)
	v50 = int32(0)
	if v37 <= v50 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v53 = int32(8)
	v57 = v50
	goto L21
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+(v9+v53)))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57+(v8+v53)))))
	if v64 != v66 {
		goto L14
	} else {
		goto L23
	}
L22:
	;
	goto L12
L23:
	;
	v69 = v57 + int32(1)
	if v37 != v69 {
		v57 = v69
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v86)
	if v77 == int32(0) {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v90 = int32(8)
	v95 = int32(0)
	goto L27
L27:
	;
	v102 = v95 << (uint(int32(2)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v9+v90+v102)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v8+v90+v102)))
	if v104 != v106 {
		goto L14
	} else {
		goto L29
	}
L28:
	;
	goto L12
L29:
	;
	v109 = v95 + int32(1)
	if v77 != v109 {
		v95 = v109
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
