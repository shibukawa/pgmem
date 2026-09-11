package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_add_page(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	F_tbm_mark_page_lossy(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 <= v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = v18 & v19
	v22 = v20
	v24 = v14
	v27 = int32(0)
	v28 = v17
	v29 = v15
	goto L6
L5:
	;
	v87 = base.I32_div_s(v84, int32(2))
	if v80 <= v87 {
		goto L3
	} else {
		goto L19
	}
L6:
	;
	v33 = v22
	v36 = v27
	v38 = v27
	goto L8
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v55
	v80 = v71
	v84 = v72
	goto L5
L8:
	;
	if v36&int32(1) != 0 {
		v80 = v24
		v84 = v29
		goto L5
	} else {
		goto L10
	}
L9:
	;
	if v57 == int32(0) {
		v80 = v24
		v84 = v29
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v46 = int32(1)
	v47 = v33 - v46
	v51 = base.B2i32(v45&(v47^v20) == int32(0))
	v52 = v51 | v38
	v55 = v45 & v47
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v57 = v33*int32(48) + v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
	if v58 != v46 {
		v33 = v55
		v36 = v51
		v38 = v52
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+5)))
	if v63 != 0 {
		v22 = v55
		v27 = v52
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v64&int32(255) == int32(0) {
		v22 = v55
		v27 = v52
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_tbm_mark_page_lossy(m, l0, v64)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = base.I32_div_s(v72, int32(2))
	if v74 < v71 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v55
	v24 = v71
	v27 = v52
	v28 = v76
	v29 = v72
	goto L6
L17:
	;
	goto L18
L18:
	;
	goto L7
L19:
	;
	v89 = int32(1073741823)
	if v89 <= v80 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v92 = v89
	goto L22
L21:
	;
	v92 = v80
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92 << (uint(int32(1)) % 32)
	goto L3
}
func F_tbm_intersect_page(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v504 int32
	_ = v504
	v3 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v504
L2:
	;
	v28 = int32(1)
	v32 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v267 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v39 = l0 + int32(8) + v32<<(uint(int32(2))%32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v504 = v253
	goto L1
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = v40
	v49 = v41 + v32<<(uint(int32(5))%32)
	v54 = v40
	v57 = int32(0)
	goto L10
L8:
	;
	v253 = v28
	goto L9
L9:
	;
	v263 = v32 + int32(1)
	if v263 != int32(8) {
		v28 = v253
		v32 = v263
		goto L5
	} else {
		goto L40
	}
L10:
	;
	if v48&int32(1) == int32(0) {
		v228 = v54
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v228
	v253 = base.B2i32(v228 == int32(0)) & v28
	goto L9
L12:
	;
	v235 = int32(1)
	if base.Ui32(v235) < base.Ui32(v48) {
		v48 = int32(base.Ui32(v48) >> (uint(v235) % 32))
		v49 = v49 + v235
		v54 = v228
		v57 = v57 + v235
		goto L10
	} else {
		goto L39
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v65 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v144 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v72 = v49 & int32(-256)
	v73 = int32(16)
	v77 = (v72 ^ int32(base.Ui32(v49)>>(uint(v73)%32))) * int32(-2048144789)
	v82 = (int32(base.Ui32(v77)>>(uint(int32(13))%32)) ^ v77) * int32(-1028477387)
	v86 = v70 & (int32(base.Ui32(v82)>>(uint(v73)%32)) ^ v82)
	v89 = v69 + v86*int32(48)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v90 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v95 = v89
	v98 = v86
	goto L17
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v72 != v108 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+5)))
	if v117 != int32(1) {
		goto L14
	} else {
		goto L23
	}
L19:
	;
	v112 = (v98 + int32(1)) & v70
	v115 = v69 + v112*int32(48)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
	if v116 != 0 {
		v95 = v115
		v98 = v112
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L14
L23:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(base.Ui32(v49)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v125)>>(uint(v49)%32))&int32(1) != 0 {
		v228 = v54
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L14
L25:
	;
	v228 = v54 & base.I32_rotl(int32(-2), v57)
	goto L12
L26:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v147 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v150 != v49 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v155 = int32(16)
	v159 = (int32(base.Ui32(v49)>>(uint(v155)%32)) ^ v49) * int32(-2048144789)
	v164 = (int32(base.Ui32(v159)>>(uint(int32(13))%32)) ^ v159) * int32(-1028477387)
	v168 = v154 & (int32(base.Ui32(v164)>>(uint(v155)%32)) ^ v164)
	v171 = v153 + v168*int32(48)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+4)))
	if v172 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L30:
	;
	v228 = v54
	goto L12
L31:
	;
	v177 = v171
	v180 = v168
	goto L32
L32:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v49 != v190 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+5)))
	if v199 != int32(1) {
		v228 = v54
		goto L12
	} else {
		goto L38
	}
L34:
	;
	v194 = (v180 + int32(1)) & v154
	v197 = v153 + v194*int32(48)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+4)))
	if v198 != 0 {
		v177 = v197
		v180 = v194
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L25
L38:
	;
	goto L25
L39:
	;
	goto L11
L40:
	;
	goto L6
L41:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v350 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	v274 = v266 & int32(-256)
	v275 = int32(16)
	v279 = (v274 ^ int32(base.Ui32(v266)>>(uint(v275)%32))) * int32(-2048144789)
	v284 = (int32(base.Ui32(v279)>>(uint(int32(13))%32)) ^ v279) * int32(-1028477387)
	v288 = v272 & (int32(base.Ui32(v284)>>(uint(v275)%32)) ^ v284)
	v291 = v271 + v288*int32(48)
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+4)))
	if v292 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v297 = v291
	v300 = v288
	goto L44
L44:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	if v274 != v310 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+5)))
	if v319 != int32(1) {
		goto L41
	} else {
		goto L50
	}
L46:
	;
	v314 = (v300 + int32(1)) & v272
	v317 = v271 + v314*int32(48)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+4)))
	if v318 != 0 {
		v297 = v317
		v300 = v314
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	goto L41
L50:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v297+int32(base.Ui32(v266)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v327)>>(uint(v266)%32))&int32(1) == int32(0) {
		goto L41
	} else {
		goto L51
	}
L51:
	;
	v333 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v333)
	v504 = v3
	goto L1
L52:
	;
	v504 = int32(1)
	goto L1
L53:
	;
	goto L54
L54:
	;
	v354 = int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v355 == v354 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v412)+8))
	v427 = v425 & v426
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v427
	v430 = l0 + int32(12)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v433 = v431 & v432
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v433
	v436 = l0 + int32(16)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v412)+16))
	v439 = v437 & v438
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = v439
	v442 = l0 + int32(20)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v412)+20))
	v445 = v443 & v444
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = v445
	v448 = l0 + int32(24)
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v412)+24))
	v451 = v449 & v450
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v451
	v454 = l0 + int32(28)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v412)+28))
	v457 = v455 & v456
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = v457
	v460 = l0 + int32(32)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v412)+32))
	v463 = v461 & v462
	*(*int32)(unsafe.Add(mBase, uint32(v460))) = v463
	v466 = l0 + int32(36)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v412)+36))
	v469 = v467 & v468
	*(*int32)(unsafe.Add(mBase, uint32(v466))) = v469
	v472 = l0 + int32(40)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v412)+40))
	v475 = v473 & v474
	*(*int32)(unsafe.Add(mBase, uint32(v472))) = v475
	v478 = l0 + int32(44)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v412)+44))
	v481 = v479 & v480
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = v481
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+6)))
	v485 = v483 | v484
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v485)
	v504 = base.B2i32(v475|v481|v469|v463|v457|v451|v445|v439|v433|v427 == int32(0))
	goto L1
L56:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v358 != v266 {
		v504 = v354
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+20))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v365 = int32(16)
	v369 = (int32(base.Ui32(v266)>>(uint(v365)%32)) ^ v266) * int32(-2048144789)
	v374 = (int32(base.Ui32(v369)>>(uint(int32(13))%32)) ^ v369) * int32(-1028477387)
	v378 = v364 & (int32(base.Ui32(v374)>>(uint(v365)%32)) ^ v374)
	v381 = v363 + v378*int32(48)
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+4)))
	if v382 == int32(0) {
		v504 = v354
		goto L1
	} else {
		goto L60
	}
L59:
	;
	v412 = l1 + int32(40)
	goto L55
L60:
	;
	v387 = v381
	v390 = v378
	goto L61
L61:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	if v266 != v400 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+5)))
	if v409 != 0 {
		v504 = v354
		goto L1
	} else {
		goto L67
	}
L63:
	;
	v404 = (v390 + int32(1)) & v364
	v407 = v363 + v404*int32(48)
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+4)))
	if v408 != 0 {
		v387 = v407
		v390 = v404
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	v504 = v354
	goto L1
L67:
	;
	v412 = v387
	goto L55
}
