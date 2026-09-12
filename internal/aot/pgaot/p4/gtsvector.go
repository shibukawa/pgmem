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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
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
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
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
	v518 = F_palloc(m, int32(16))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L10
	} else {
		goto L85
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
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v448&int32(6) != int32(2) {
		goto L73
	} else {
		goto L74
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
	v65 = v56 + int32(8)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v49 + int32(8)
	v76 = v68
	v83 = v65
	v85 = v66
	goto L21
L19:
	;
	v216 = int32(0)
	goto L20
L20:
	;
	F_pg_qsort(m, v65, v216, int32(4), int32(1519))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L36
	}
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v89 = int32(base.Ui32(v87) >> (uint(int32(1)) % 32))
	v91 = v89 & int32(2047)
	if v91 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v216 = v200
	goto L20
L23:
	;
	v92 = int32(1)
	v93 = v89 & v92
	v96 = v68 + v66<<(uint(int32(2))%32) + int32(base.Ui32(v87)>>(uint(int32(12))%32))
	v97 = int32(-1)
	if v91 != v92 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v192 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v192
	v194 = int32(4)
	v199 = v85 - int32(1)
	if v199 != 0 {
		v76 = v76 + v194
		v83 = v83 + v194
		v85 = v199
		goto L21
	} else {
		goto L35
	}
L26:
	;
	v102 = v96
	v103 = v97
	v104 = int32(0)
	goto L29
L27:
	;
	v146 = v96
	v147 = v97
	goto L28
L28:
	;
	if v93 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v119 = int32(24)
	v122 = int32(2)
	v126 = *(*int32)(unsafe.Add(mBase, uint32((v118^int32(base.Ui32(v103)>>(uint(v119)%32)))<<(uint(v122)%32))+uint32(_c_F_gtsvector_compress[0])))
	v127 = int32(8)
	v129 = v126 ^ v103<<(uint(v127)%32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32((v117^int32(base.Ui32(v129)>>(uint(v119)%32)))<<(uint(v122)%32))+uint32(_c_F_gtsvector_compress[0])))
	v140 = v137 ^ v129<<(uint(v127)%32)
	v142 = v102 + v122
	v144 = v104 + v122
	if v144 != v91-v93 {
		v102 = v142
		v103 = v140
		v104 = v144
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v146 = v142
	v147 = v140
	goto L28
L31:
	;
	goto L30
L32:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v169 = *(*int32)(unsafe.Add(mBase, uint32((v161^int32(base.Ui32(v147)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_gtsvector_compress[0])))
	v173 = v169 ^ v147<<(uint(int32(8))%32)
	goto L34
L33:
	;
	v173 = v147
	goto L34
L34:
	;
	v192 = v173 ^ int32(-1)
	goto L25
L35:
	;
	goto L22
L36:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if base.Ui32(v221) < base.Ui32(int32(2)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(v290) < base.Ui32(int32(2044)) {
		goto L48
	} else {
		goto L49
	}
L38:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v290 = v289
	v295 = v56
	goto L37
L39:
	;
	v225 = int32(1)
	v234 = v2
	goto L40
L40:
	;
	v240 = int32(2)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v65+v225<<(uint(v240)%32))))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v65+v234<<(uint(v240)%32))))
	if v243 == v247 {
		v257 = v234
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v262 = v257 + int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v262 == v263 {
		goto L38
	} else {
		goto L46
	}
L42:
	;
	v259 = v225 + int32(1)
	if v259 != v221 {
		v225 = v259
		v234 = v257
		goto L40
	} else {
		goto L45
	}
L43:
	;
	v250 = v234 + int32(1)
	if v225 == v250 {
		v257 = v225
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v250<<(uint(int32(2))%32)))) = v243
	v257 = v250
	goto L42
L45:
	;
	goto L41
L46:
	;
	v268 = v262<<(uint(int32(2))%32) + int32(8)
	v269 = F_repalloc(m, v56, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v272 = v268 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v272
	v290 = v272
	v295 = v269
	goto L37
L48:
	;
	v508 = v295
	goto L12
L49:
	;
	goto L50
L50:
	;
	v308 = v44 + int32(8)
	v309 = F_palloc(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v311 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v309)+4)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v308 << (uint(v311) % 32)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v319 = int32(8)
	v320 = int32(base.Ui32(v316)>>(uint(v311)%32)) - v319
	v322 = v309 + v319
	if v322&int32(3) != 0 {
		v342 = v44
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v320) < base.Ui32(int32(4)) {
		v508 = v309
		goto L12
	} else {
		goto L62
	}
L53:
	;
	v346 = F__emscripten_memset_bulkmem(m, v322, base.I32_extend8_s(int32(0)), v342)
	mBase = m.M
	goto L61
L54:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v44) {
		v342 = v44
		goto L53
	} else {
		goto L55
	}
L55:
	;
	if v44&int32(3) != 0 {
		v342 = v44
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v329 = v44 + v322
	if base.Ui32(v329) <= base.Ui32(v322) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v334 = v309 + int32(12)
	if base.Ui32(v334) < base.Ui32(v329) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v336 = v329
	goto L60
L59:
	;
	v336 = v334
	goto L60
L60:
	;
	v342 = (v322^int32(-1)+v336)&int32(-4) + int32(4)
	goto L53
L61:
	;
	goto L52
L62:
	;
	v352 = v295 + int32(8)
	v354 = v44 << (uint(int32(3)) % 32)
	v355 = int32(1)
	v357 = int32(base.Ui32(v320) >> (uint(int32(2)) % 32))
	if base.Ui32(v357) <= base.Ui32(v355) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v360 = v355
	goto L65
L64:
	;
	v360 = v357
	goto L65
L65:
	;
	v363 = int32(0)
	if base.Ui32(int32(8)) <= base.Ui32(v320) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v369 = v363
	v372 = int32(0)
	goto L69
L67:
	;
	v416 = v363
	goto L68
L68:
	;
	if v360&int32(1) == int32(0) {
		v508 = v309
		goto L12
	} else {
		goto L72
	}
L69:
	;
	v384 = int32(2)
	v386 = v352 + v369<<(uint(v384)%32)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v388 = base.I32_rem_u_s(v387, v354)
	v389 = int32(3)
	v391 = v322 + int32(base.Ui32(v388)>>(uint(v389)%32))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	v393 = int32(1)
	v394 = int32(7)
	v397 = v392 | v393<<(uint(v388&v394)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v397)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	v400 = base.I32_rem_u_s(v399, v354)
	v403 = v322 + int32(base.Ui32(v400)>>(uint(v389)%32))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	v409 = v404 | v393<<(uint(v400&v394)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v403))) = uint8(v409)
	v412 = v369 + v384
	v414 = v372 + v384
	if v414 != v360&int32(1073741822) {
		v369 = v412
		v372 = v414
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v416 = v412
	goto L68
L71:
	;
	goto L70
L72:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v352+v416<<(uint(int32(2))%32))))
	v437 = base.I32_rem_u_s(v436, v354)
	v440 = v322 + int32(base.Ui32(v437)>>(uint(int32(3))%32))
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	v446 = v441 | int32(1)<<(uint(v437&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v446)
	v508 = v309
	goto L12
L73:
	;
	return v16
L74:
	;
	goto L75
L75:
	;
	v454 = int32(0)
	if v44 <= v454 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v498 = F_palloc(m, int32(8))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L10
	} else {
		goto L84
	}
L77:
	;
	v459 = v454
	goto L78
L78:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+(v45+int32(8))))))
	if v475 == int32(255) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	return v16
L80:
	;
	v479 = v459 + int32(1)
	if v44 != v479 {
		v459 = v479
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	goto L76
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v498))) = int64(25769803808)
	v508 = v498
	goto L12
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = v508
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+4)) = v521
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v518)+8)) = v523
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)))
	v526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v518)+14)) = uint8(v526)
	*(*uint16)(unsafe.Add(mBase, uint32(v518)+12)) = uint16(v525)
	return v518
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v13)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v16 == v2 {
		v54 = v2
		m.G0 = v7 + int32(16)
		return v54
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v19&int32(2) != 0 {
			if v19&int32(4) != 0 {
				v54 = int32(1)
				m.G0 = v7 + int32(16)
				return v54
			} else {
				v29 = F_TS_execute(m, v9+int32(8), v11, int32(2), int32(1520))
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
			v34 = v11 + v33
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v37 = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v34 + (int32(base.Ui32(v36)>>(uint(v37)%32))-v33)&int32(-4)
			v51 = F_TS_execute(m, v9+v33, v7+v33, v37, int32(1521))
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == v2 {
		v29 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v29&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v16 == int32(0) {
		v29 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != int32(7) {
		v29 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != int32(17) {
		v29 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)))
	v29 = v25 ^ int32(1)
	goto L2
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = F_get_fn_opclass_options(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v38 = int32(124)
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v39&int32(2) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v38 = v37
	goto L9
L12:
	;
	return v8
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v122)
	goto L12
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v43 = int32(4)
	v44 = v42 & v43
	if v39&v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v74 = int32(2)
	v76 = int32(8)
	v77 = int32(base.Ui32(v73)>>(uint(v74)%32)) - v76
	v79 = int32(base.Ui32(v77) >> (uint(v74) % 32))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v79 != int32(base.Ui32(int32(base.Ui32(v80)>>(uint(v74)%32))-v76)>>(uint(v74)%32)) {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	v122 = int32(base.Ui32(v44) >> (uint(int32(2)) % 32))
	goto L13
L18:
	;
	goto L19
L19:
	;
	if v44 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v122 = v2
	goto L13
L21:
	;
	goto L22
L22:
	;
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v49)
	v51 = int32(0)
	if v38 <= v51 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v54 = int32(8)
	v58 = v51
	goto L24
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+(v10+v54)))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+(v9+v54)))))
	if v66 == v68 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v122 = v2
	goto L13
L26:
	;
	v71 = v58 + int32(1)
	if v38 != v71 {
		v58 = v71
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L12
L30:
	;
	v122 = v2
	goto L13
L31:
	;
	goto L32
L32:
	;
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v88)
	if base.Ui32(v77) < base.Ui32(int32(4)) {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	v92 = int32(8)
	v96 = int32(1)
	if base.Ui32(v79) <= base.Ui32(v96) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = v96
	goto L36
L35:
	;
	v99 = v79
	goto L36
L36:
	;
	v101 = int32(0)
	goto L37
L37:
	;
	v109 = v101 << (uint(int32(2)) % 32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v10+v92+v109)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+(v9+v92))))
	if v111 != v113 {
		v122 = v2
		goto L13
	} else {
		goto L39
	}
L38:
	;
	goto L12
L39:
	;
	v116 = v101 + int32(1)
	if v99 != v116 {
		v101 = v116
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
}
