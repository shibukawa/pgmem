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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	v23 = v14
	v26 = v17
	v27 = v15
	v28 = int32(0)
	goto L6
L5:
	;
	v77 = base.I32_div_s(v27, int32(2))
	if v23 <= v77 {
		goto L3
	} else {
		goto L18
	}
L6:
	;
	v33 = v22
	v35 = v28
	v39 = v28
	goto L8
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v55
	return
L8:
	;
	if v35&int32(1) != 0 {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+5)))
	if v61 != 0 {
		v22 = v55
		v28 = v52
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v46 = int32(1)
	v47 = v33 - v46
	v51 = base.B2i32(v45&(v47^v20) == int32(0))
	v52 = v51 | v39
	v55 = v45 & v47
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v57 = v33*int32(48) + v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
	if v58 != v46 {
		v33 = v55
		v35 = v51
		v39 = v52
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v62&int32(255) == int32(0) {
		v22 = v55
		v28 = v52
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_tbm_mark_page_lossy(m, l0, v62)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v72 = base.I32_div_s(v70, int32(2))
	if v72 < v69 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v55
	v23 = v69
	v26 = v74
	v27 = v70
	v28 = v52
	goto L6
L16:
	;
	goto L17
L17:
	;
	goto L7
L18:
	;
	v79 = int32(1073741823)
	if v79 <= v23 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v82 = v79
	goto L21
L20:
	;
	v82 = v23
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v82 << (uint(int32(1)) % 32)
	goto L3
}
func F_tbm_intersect_page(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v485 int32
	_ = v485
	v3 = int32(0)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v485
L2:
	;
	v29 = int32(1)
	v30 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v265 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v39 = l0 + int32(8) + v30<<(uint(int32(2))%32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v485 = v252
	goto L1
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = v40
	v50 = v41 + v30<<(uint(int32(5))%32)
	v52 = v40
	v55 = int32(0)
	goto L10
L8:
	;
	v252 = v29
	goto L9
L9:
	;
	v261 = v30 + int32(1)
	if v261 != int32(8) {
		v29 = v252
		v30 = v261
		goto L5
	} else {
		goto L40
	}
L10:
	;
	if v52&int32(1) == int32(0) {
		v223 = v49
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v223
	v252 = base.B2i32(v223 == int32(0)) & v29
	goto L9
L12:
	;
	v235 = int32(1)
	v240 = int32(base.Ui32(v52) >> (uint(v235) % 32))
	if v240 != 0 {
		v49 = v223
		v50 = v50 + v235
		v52 = v240
		v55 = v55 + v235
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
	v72 = v50 & int32(-256)
	v73 = int32(16)
	v77 = (v72 ^ int32(base.Ui32(v50)>>(uint(v73)%32))) * int32(-2048144789)
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
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(base.Ui32(v50)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v125)>>(uint(v50)%32))&int32(1) != 0 {
		v223 = v49
		goto L12
	} else {
		goto L24
	}
L24:
	;
	goto L14
L25:
	;
	v223 = v49 & base.I32_rotl(int32(-2), v55)
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
	if v150 != v50 {
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
	v159 = (int32(base.Ui32(v50)>>(uint(v155)%32)) ^ v50) * int32(-2048144789)
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
	v223 = v49
	goto L12
L31:
	;
	v177 = v171
	v180 = v168
	goto L32
L32:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v50 != v190 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+5)))
	if v199 != int32(1) {
		v223 = v49
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
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v348 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+20))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v272 = v264 & int32(-256)
	v273 = int32(16)
	v277 = (v272 ^ int32(base.Ui32(v264)>>(uint(v273)%32))) * int32(-2048144789)
	v282 = (int32(base.Ui32(v277)>>(uint(int32(13))%32)) ^ v277) * int32(-1028477387)
	v286 = v270 & (int32(base.Ui32(v282)>>(uint(v273)%32)) ^ v282)
	v289 = v269 + v286*int32(48)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+4)))
	if v290 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v295 = v289
	v298 = v286
	goto L44
L44:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v272 != v308 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+5)))
	if v317 != int32(1) {
		goto L41
	} else {
		goto L50
	}
L46:
	;
	v312 = (v298 + int32(1)) & v270
	v315 = v269 + v312*int32(48)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+4)))
	if v316 != 0 {
		v295 = v315
		v298 = v312
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
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v295+int32(base.Ui32(v264)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v325)>>(uint(v264)%32))&int32(1) == int32(0) {
		goto L41
	} else {
		goto L51
	}
L51:
	;
	v331 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v331)
	v485 = v3
	goto L1
L52:
	;
	v485 = int32(1)
	goto L1
L53:
	;
	goto L54
L54:
	;
	v352 = int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v353 == v352 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	v425 = v423 & v424
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v410)+12))
	v429 = v427 & v428
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v410)+16))
	v433 = v431 & v432
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v433
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v410)+20))
	v437 = v435 & v436
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v437
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v410)+24))
	v441 = v439 & v440
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v410)+28))
	v445 = v443 & v444
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v445
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v410)+32))
	v449 = v447 & v448
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v410)+36))
	v453 = v451 & v452
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v453
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v410)+40))
	v457 = v455 & v456
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v457
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v410)+44))
	v461 = v459 & v460
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v461
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+6)))
	v465 = v463 | v464
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v465)
	v485 = base.B2i32(v457|v461|v453|v449|v445|v441|v437|v433|v429|v425 == int32(0))
	goto L1
L56:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v356 != v264 {
		v485 = v352
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)+20))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	v363 = int32(16)
	v367 = (int32(base.Ui32(v264)>>(uint(v363)%32)) ^ v264) * int32(-2048144789)
	v372 = (int32(base.Ui32(v367)>>(uint(int32(13))%32)) ^ v367) * int32(-1028477387)
	v376 = v362 & (int32(base.Ui32(v372)>>(uint(v363)%32)) ^ v372)
	v379 = v361 + v376*int32(48)
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+4)))
	if v380 == int32(0) {
		v485 = v352
		goto L1
	} else {
		goto L60
	}
L59:
	;
	v410 = l1 + int32(40)
	goto L55
L60:
	;
	v385 = v379
	v388 = v376
	goto L61
L61:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v264 != v398 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+5)))
	if v407 != 0 {
		v485 = v352
		goto L1
	} else {
		goto L67
	}
L63:
	;
	v402 = (v388 + int32(1)) & v362
	v405 = v361 + v402*int32(48)
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+4)))
	if v406 != 0 {
		v385 = v405
		v388 = v402
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
	v485 = v352
	goto L1
L67:
	;
	v410 = v385
	goto L55
}
