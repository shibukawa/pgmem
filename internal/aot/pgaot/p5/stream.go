package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v2)
	goto L1
L1:
	;
	v20 = F_read_stream_next_buffer(m, l0, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+76)))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v25 <= v24 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	return
L4:
	;
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ReleaseBuffer(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L2
L8:
	;
	goto L1
L9:
	;
	v78 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v78)
	return
L10:
	;
	v28 = l0 + int32(80)
	v30 = v24
	goto L11
L11:
	;
	v35 = base.I32_extend16_s(v30)
	v38 = v28 + v35<<(uint(int32(2))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v44 = v42 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v44)
	F_ReleaseBuffer(m, v39)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(0)
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	if v35 < v50-int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+(v54+v35)<<(uint(int32(2))%32)))) = int32(0)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v62 = v30 + int32(1)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v64 != v62&int32(65535) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v62
	goto L20
L19:
	;
	v68 = int32(0)
	goto L20
L20:
	;
	v69 = base.I32_extend16_s(v68)
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v69 < v70 {
		v30 = v69
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
}
func F_read_stream_start_pending_read(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v411 int32
	_ = v411
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v454 int64
	_ = v454
	var v456 int32
	_ = v456
	var v458 int64
	_ = v458
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v512 int64
	_ = v512
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int64
	_ = v553
	var v559 int64
	_ = v559
	var v567 int32
	_ = v567
	var v568 int64
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int64
	_ = v579
	var v597 int32
	_ = v597
	var v603 int64
	_ = v603
	var v609 int64
	_ = v609
	var v611 int32
	_ = v611
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v785 int32
	_ = v785
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v981 int32
	_ = v981
	v27 = m.G0
	v29 = v27 - int32(16)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v32 != int32(1) {
		v52 = v31
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v53 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v35 == v36 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v40 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v35
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v47 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v43 = v31
	goto L8
L7:
	;
	v43 = v31 | int32(2)
	goto L8
L8:
	;
	v52 = v43
	goto L1
L9:
	;
	v50 = v31 | int32(2)
	goto L11
L10:
	;
	v50 = v31
	goto L11
L11:
	;
	v52 = v50
	goto L1
L12:
	;
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+10)))
	v98 = v96 + v97
	if v98 != 0 {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[598]))
	v60 = *(*int32)(unsafe.Add(mBase, _consts[599]))
	goto L16
L14:
	;
	goto L15
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[600]))
	v75 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	v78 = v73 - v75 - int32(8)
	if base.Ui32(v78) <= base.Ui32(v73) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	if base.Ui32(int32(32766)) < base.Ui32(v58-v60) {
		v96 = int32(32767)
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[598]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[599]))
	goto L18
L18:
	;
	v96 = v65 - v67
	goto L12
L19:
	;
	if base.Ui32(int32(32766)) < base.Ui32(v81) {
		v96 = int32(32767)
		goto L12
	} else {
		goto L23
	}
L20:
	;
	v81 = v78
	goto L22
L21:
	;
	v81 = int32(0)
	goto L22
L22:
	;
	goto L19
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[600]))
	v89 = *(*int32)(unsafe.Add(mBase, _consts[601]))
	v92 = v87 - v89 - int32(8)
	if base.Ui32(v92) <= base.Ui32(v87) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v96 = v95
	goto L12
L25:
	;
	v95 = v92
	goto L27
L26:
	;
	v95 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v107
	if v106 < v107 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v99 = int32(32767)
	if v99 <= v98 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v106 = base.B2i32(v103 == int32(0))
	goto L28
L32:
	;
	v102 = v99
	goto L34
L33:
	;
	v102 = v98
	goto L34
L34:
	;
	v106 = v102
	goto L28
L35:
	;
	m.G0 = v29 + int32(16)
	return v981
L36:
	;
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v112 = base.I32_extend16_s(v110 + v106)
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v112 < v113 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v120 = v107
	goto L38
L38:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+70)))
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+76)))
	v125 = v120 + v124
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v126 < v125 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v112)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v116 = int32(0)
	if v116 < v110 {
		v981 = v116
		goto L35
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v106
	v120 = v106
	goto L38
L43:
	;
	v131 = v126
	goto L46
L44:
	;
	goto L45
L45:
	;
	v193 = v123 * int32(84)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v195 = v193 + v194
	v197 = v195 + int32(4)
	v199 = l0 + int32(80)
	v202 = v199 + v124<<(uint(int32(2))%32)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v204 = int32(0)
	v205 = m.G0
	v207 = v205 + int32(-64)
	m.G0 = v207
	v210 = v29 + int32(12)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v211 <= v204 {
		v759 = v211
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v157 = v131 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v157)
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(80)+v131<<(uint(int32(2))%32)))) = int32(0)
	v164 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v164 < v125 {
		v131 = v164
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	goto L47
L49:
	;
	m.G0 = v207 - int32(-64)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v845 = v843 + v844
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v845)
	if v822 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v759
	v785 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v197)+34)) = uint16(v785)
	*(*uint16)(unsafe.Add(mBase, uint32(v197)+32)) = uint16(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v197)+24)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v197)+20)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v195+int32(40)))) = int32(-1)
	goto L160
L51:
	;
	v215 = v211
	v222 = v204
	goto L52
L52:
	;
	v242 = v202 + v222<<(uint(int32(2))%32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v243 != 0 {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	v759 = v752
	goto L50
L54:
	;
	if v215 < int32(2) {
		v752 = v215
		goto L148
	} else {
		goto L149
	}
L55:
	;
	if v222 != 0 {
		goto L145
	} else {
		goto L146
	}
L56:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+24))
	v653 = int32(base.Ui32(v649)>>(uint(int32(24))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)) = uint8(v653)
	if v653 == int32(0) {
		goto L54
	} else {
		goto L144
	}
L57:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v648 = v642 + v243<<(uint(int32(6))%32) + int32(-64)
	goto L56
L58:
	;
	if int32(0) <= v243 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v253 = v222 + v203
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+8)))
	if v257 == int32(116) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v648 = v247 + (v243^int32(-1))<<(uint(int32(6))%32)
	goto L56
L62:
	;
	if v256 != 0 {
		goto L122
	} else {
		goto L123
	}
L63:
	;
	v260 = int32(1)
	v263 = F_LocalBufferAlloc(m, v255, v254, v253, v205+int32(-45))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	v278 = F_IOContextForStrategy(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L66
	} else {
		goto L69
	}
L66:
	;
	return int32(0)
L67:
	;
	v267 = int32(3)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)))
	if v268 != int32(1) {
		v519 = v263
		v521 = v268
		v528 = v260
		v537 = v267
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v271 = int32(4447064)
	v273 = *(*int64)(unsafe.Add(mBase, _consts[345]))
	*(*int64)(unsafe.Add(mBase, _consts[345])) = v273 + int64(1)
	v519 = v263
	v521 = v268
	v528 = v260
	v537 = v267
	goto L62
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	F_ResourceOwnerEnlarge(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+24)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+36)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v207)+32)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v207)+28)) = v290
	v296 = F_BufTableHashCode(m, v205+int32(-44))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v306 = v299 + v296&int32(127)<<(uint(int32(7))%32) + int32(6912)
	v308 = F_LWLockAcquire(m, v306, int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L66
	} else {
		goto L73
	}
L73:
	;
	v312 = F_BufTableLookup(m, v205+int32(-44), v296)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L66
	} else {
		goto L75
	}
L74:
	;
	v510 = int32(4447032)
	v512 = *(*int64)(unsafe.Add(mBase, _consts[341]))
	*(*int64)(unsafe.Add(mBase, _consts[341])) = v512 + int64(1)
	v519 = v503
	v521 = int32(1)
	v528 = int32(0)
	v537 = v278
	goto L62
L75:
	;
	if int32(0) <= v312 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v321 = v318 + v312<<(uint(int32(6))%32)
	v322 = F_PinBuffer(m, v321, v277)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L66
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_LWLockRelease(m, v306)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L66
	} else {
		goto L82
	}
L79:
	;
	F_LWLockRelease(m, v306)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L66
	} else {
		goto L80
	}
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)) = uint8(v322)
	if v322 != 0 {
		v503 = v321
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v519 = v321
	v521 = int32(0)
	v528 = int32(0)
	v537 = v278
	goto L62
L82:
	;
	v330 = F_GetVictimBuffer(m, v277, v278)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L66
	} else {
		goto L83
	}
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v335 = F_LWLockAcquire(m, v306, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v339 = v333 + v330<<(uint(int32(6))%32)
	v341 = v339 + int32(-64)
	v345 = v339 - int32(44)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v347 = F_BufTableInsert(m, v205+int32(-44), v296, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L66
	} else {
		goto L85
	}
L85:
	;
	if v347 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+60)) = int32(239002)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+56)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+52)) = int32(514763)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v207)+40)) = int64(0)
	v362 = v339 - int32(40)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v364 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v363 | v364
	if v363&v364 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	F_ResourceOwnerForget(m, v481, v482+int32(1), int32(1656256))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L66
	} else {
		goto L114
	}
L89:
	;
	goto L92
L90:
	;
	v411 = v363
	goto L91
L91:
	;
	v434 = int32(4155052)
	v435 = *(*int32)(unsafe.Add(mBase, _consts[602]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v205+int32(-24))+8))
	if v437 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	F_perform_spin_delay(m, v205+int32(-24))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L66
	} else {
		goto L94
	}
L93:
	;
	v411 = v399
	goto L91
L94:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	v400 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v399 | v400
	if v399&v400 != 0 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v207)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v454
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v207)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v341)+16)) = v456
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v207)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v341)+8)) = v458
	v462 = int32(-2113667072)
	if v254 == int32(3) {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, _consts[602])) = v452
	goto L97
L99:
	;
	if int32(999) < v435 {
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v435 < int32(11) {
		goto L97
	} else {
		goto L106
	}
L102:
	;
	v442 = int32(900)
	if v442 <= v435 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v445 = v442
	goto L105
L104:
	;
	v445 = v435
	goto L105
L105:
	;
	v452 = v445 + int32(100)
	goto L98
L106:
	;
	v452 = v435 - int32(1)
	goto L98
L107:
	;
	v467 = v462
	goto L109
L108:
	;
	v467 = int32(33816576)
	goto L109
L109:
	;
	if v257 == int32(112) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v470 = v462
	goto L112
L111:
	;
	v470 = v467
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v411&int32(-38010881) | v470
	F_LWLockRelease(m, v306)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L66
	} else {
		goto L113
	}
L113:
	;
	v475 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)) = uint8(v475)
	v519 = v341
	v521 = v475
	v528 = v475
	v537 = v278
	goto L62
L114:
	;
	F_UnpinBufferNoOwner(m, v341)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L66
	} else {
		goto L115
	}
L115:
	;
	F_StrategyFreeBuffer(m, v341)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L66
	} else {
		goto L116
	}
L116:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v496 = v493 + v347<<(uint(int32(6))%32)
	v497 = F_PinBuffer(m, v496, v277)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L66
	} else {
		goto L117
	}
L117:
	;
	F_LWLockRelease(m, v306)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L66
	} else {
		goto L118
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)) = uint8(v497)
	if v497 != 0 {
		v503 = v496
		goto L74
	} else {
		goto L119
	}
L119:
	;
	v519 = v496
	v521 = int32(0)
	v528 = int32(0)
	v537 = v278
	goto L62
L120:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v519)+20))
	v635 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v634 + v635
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)))
	if v638&v635 != 0 {
		goto L55
	} else {
		goto L143
	}
L121:
	;
	v597 = v528*int32(320) + v537<<(uint(int32(6))%32)
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v597)+uint32(_consts[603])))
	*(*int64)(unsafe.Add(mBase, uint32(v597)+uint32(_consts[603]))) = v603 + int64(1)
	v609 = *(*int64)(unsafe.Add(mBase, uint32(v597)+uint32(_consts[604])))
	*(*int64)(unsafe.Add(mBase, uint32(v597)+uint32(_consts[604]))) = v609
	v611 = int32(1)
	F_pgstat_count_backend_io_op(m, v528, v537, int32(2), v611, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _consts[90])) = uint8(v611)
	*(*uint8)(unsafe.Add(mBase, _consts[605])) = uint8(v611)
	goto L141
L122:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v256)+272))
	if v543 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	goto L124
L124:
	;
	if v521 == int32(0) {
		goto L120
	} else {
		goto L140
	}
L125:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v256)+272))
	if v567 != 0 {
		goto L134
	} else {
		goto L135
	}
L126:
	;
	if v521 == int32(0) {
		goto L120
	} else {
		goto L133
	}
L127:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+268)))
	if v546 != int32(1) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v543)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v543)+112)) = v559 + int64(1)
	goto L126
L130:
	;
	F_pgstat_assoc_relation(m, v256)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L66
	} else {
		goto L131
	}
L131:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v256)+272))
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v552)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v552)+112)) = v553 + int64(1)
	if v551&int32(1) != 0 {
		goto L125
	} else {
		goto L132
	}
L132:
	;
	goto L120
L133:
	;
	goto L125
L134:
	;
	v568 = *(*int64)(unsafe.Add(mBase, uint32(v567)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v567)+120)) = v568 + int64(1)
	goto L121
L135:
	;
	goto L136
L136:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+268)))
	if v572 != int32(1) {
		goto L121
	} else {
		goto L137
	}
L137:
	;
	F_pgstat_assoc_relation(m, v256)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L66
	} else {
		goto L138
	}
L138:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+19)))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v256)+272))
	v579 = *(*int64)(unsafe.Add(mBase, uint32(v578)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v578)+120)) = v579 + int64(1)
	if v577&int32(1) != 0 {
		goto L121
	} else {
		goto L139
	}
L139:
	;
	goto L120
L140:
	;
	goto L121
L141:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, _consts[52])))
	if v621 != int32(1) {
		goto L120
	} else {
		goto L142
	}
L142:
	;
	v624 = int32(4543488)
	v626 = *(*int32)(unsafe.Add(mBase, _consts[51]))
	v628 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	*(*int32)(unsafe.Add(mBase, _consts[51])) = v626 + v628
	goto L120
L143:
	;
	goto L54
L144:
	;
	goto L55
L145:
	;
	v759 = v222
	goto L50
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(1)
	v822 = int32(0)
	goto L49
L148:
	;
	v756 = v222 + int32(1)
	if v756 < v752 {
		v215 = v752
		v222 = v756
		goto L52
	} else {
		goto L159
	}
L149:
	;
	if v222 != 0 {
		v752 = v215
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v716 = int32(4543420)
	v718 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v718 + int32(1)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v714)+36))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v722*int32(80))+uint32(_consts[607])))
	v728 = m.T0[v727].(func(*base.Module, int32, int32, int32) int32)(m, v714, v715, v203)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L66
	} else {
		goto L151
	}
L151:
	;
	v730 = int32(4543420)
	v732 = *(*int32)(unsafe.Add(mBase, _consts[412]))
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v732 - int32(1)
	if v215 <= v728 {
		v752 = v215
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v739 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L66
	} else {
		goto L153
	}
L153:
	;
	if v739 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v203
	F_errmsg_internal(m, int32(47382), v207)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L66
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v752 = v728
	goto L148
L157:
	;
	F_errfinish(m, int32(514763), int32(1382), int32(312994))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L66
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	goto L53
L160:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _consts[608]))
	if v796 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v797 = F_AsyncReadBuffers(m, v197, v210)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L66
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v197)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+28)) = v801 | int32(8)
	v805 = int32(1)
	if v52&int32(2) == int32(0) {
		v822 = v805
		goto L49
	} else {
		goto L165
	}
L164:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*uint16)(unsafe.Add(mBase, uint32(v197)+32)) = uint16(v799)
	v822 = v797
	goto L49
L165:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v812 = F_smgrprefetch(m, v810, v811, v203, v759)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L66
	} else {
		goto L166
	}
L166:
	;
	v822 = v805
	goto L49
L167:
	;
	v878 = v877 + v124
	v879 = int32(0)
	if v120 <= v877 {
		v919 = v879
		goto L175
	} else {
		goto L176
	}
L168:
	;
	v849 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v849 < int32(2) {
		v877 = v843
		goto L167
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v855+v193))) = uint16(v124)
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v859 = int32(1)
	v860 = v858 + v859
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v860)
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)))
	v864 = v862 + v859
	v866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v866 != v864&int32(65535) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v853 = v849 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v853)
	v877 = v843
	goto L167
L172:
	;
	v870 = v864
	goto L174
L173:
	;
	v870 = int32(0)
	goto L174
L174:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v870)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v872 + v873
	v877 = v872
	goto L167
L175:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v919)
	v948 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v950 = base.I32_extend16_s(v124 + (v919 + v877) - v948)
	if int32(0) < v950 {
		goto L181
	} else {
		goto L182
	}
L176:
	;
	v881 = v120 - v877
	v883 = v879
	goto L177
L177:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v199+(v883+v878)<<(uint(int32(2))%32))))
	if v912 == int32(0) {
		v919 = v883
		goto L175
	} else {
		goto L179
	}
L178:
	;
	v919 = v881
	goto L175
L179:
	;
	v916 = v883 + int32(1)
	if v916 != v881 {
		v883 = v916
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v953 = int32(2)
	v957 = v950 << (uint(v953) % 32)
	if v957 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	goto L183
L183:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v960 + v877
	v963 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v964 = v963 - v877
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v964)
	if v948 <= base.I32_extend16_s(v878) {
		goto L188
	} else {
		goto L189
	}
L184:
	;
	goto L183
L185:
	;
	v958 = F__emscripten_memcpy_bulkmem(m, v199, v199+v948<<(uint(v953)%32), v957)
	mBase = m.M
	goto L187
L186:
	;
	goto L187
L187:
	;
	goto L184
L188:
	;
	v969 = v948
	goto L190
L189:
	;
	v969 = int32(0)
	goto L190
L190:
	;
	v970 = v878 - v969
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v970)
	v981 = int32(1)
	goto L35
}
func F_stream_cleanup_files(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(1056)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l1
	v16 = F_pg_snprintf(m, v6+int32(32), int32(1024), int32(178142), v6+int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[507]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
		F_BufFileDeleteFileSet(m, v20, v6+int32(32), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			v32 = F_pg_snprintf(m, v6+int32(32), int32(1024), int32(132702), v6)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _consts[507]))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
				F_BufFileDeleteFileSet(m, v36, v6+int32(32), int32(1))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					m.G0 = v6 + int32(1056)
					return
				}
			}
		}
	}
}
func F_stream_start_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	if v11 < int32(0) {
		v15 = F_GetCurrentTimestamp(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _consts[289])) = v15
	} else {
	}
	v18 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if base.B2i32(v19 == int32(2)) == int32(0) {
		F_StartTransactionCommand(m)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_maybe_reread_subscription(m)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = F_GetTransactionSnapshot(m)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_PushActiveSnapshot(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, _consts[528]))
						*(*int32)(unsafe.Add(mBase, _consts[10])) = v34
						v37 = *(*int32)(unsafe.Add(mBase, _consts[507]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
						if v38 != 0 {
							v56 = v37
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
							v66 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(178142), v8+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								v70 = F_errstart(m, int32(14), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									if v70 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
										F_errmsg_internal(m, int32(178362), v8)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errfinish(m, int32(515057), int32(4339), int32(402103))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v83 = int32(4548768)
												v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
												v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
												if l1 != 0 {
													v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
														*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
														F_PopActiveSnapshot(m)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
														v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
															v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
															F_subxact_info_read(m, v117, l0)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v83 = int32(4548768)
										v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
										*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
										v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
										if l1 != 0 {
											v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
												*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
												F_PopActiveSnapshot(m)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
												v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
													v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
													F_subxact_info_read(m, v117, l0)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _consts[531]))
							*(*int32)(unsafe.Add(mBase, _consts[10])) = v41
							v44 = F_palloc(m, int32(44))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _consts[507]))
								*(*int32)(unsafe.Add(mBase, uint32(v47)+60)) = v44
								F_FileSetInit(m, v44)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[10])) = v34
									v54 = *(*int32)(unsafe.Add(mBase, _consts[507]))
									v56 = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
									v66 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(178142), v8+int32(16))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										v70 = F_errstart(m, int32(14), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											if v70 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
												F_errmsg_internal(m, int32(178362), v8)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errfinish(m, int32(515057), int32(4339), int32(402103))
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return
													} else {
														v83 = int32(4548768)
														v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
														v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
														v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
														if l1 != 0 {
															v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
																*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
																F_PopActiveSnapshot(m)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														} else {
															v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
																v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
																	v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
																	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
																	F_subxact_info_read(m, v117, l0)
																	mBase = m.M
																	v119 = m.ExcPending
																	if v119 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v123 = m.ExcPending
																		if v123 != 0 {
																			return
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v125 = m.ExcPending
																			if v125 != 0 {
																				return
																			} else {
																				m.G0 = v8 + int32(1056)
																				return
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v83 = int32(4548768)
												v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
												v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
												if l1 != 0 {
													v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
														*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
														F_PopActiveSnapshot(m)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
														v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
															v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
															F_subxact_info_read(m, v117, l0)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
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
					}
				}
			}
		}
	} else {
		v28 = F_GetTransactionSnapshot(m)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			F_PushActiveSnapshot(m, v28)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[528]))
				*(*int32)(unsafe.Add(mBase, _consts[10])) = v34
				v37 = *(*int32)(unsafe.Add(mBase, _consts[507]))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
				if v38 != 0 {
					v56 = v37
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
					v66 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(178142), v8+int32(16))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						v70 = F_errstart(m, int32(14), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							if v70 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
								F_errmsg_internal(m, int32(178362), v8)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									F_errfinish(m, int32(515057), int32(4339), int32(402103))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v83 = int32(4548768)
										v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
										*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
										v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
										if l1 != 0 {
											v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
												*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
												F_PopActiveSnapshot(m)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
												v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
													v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
													F_subxact_info_read(m, v117, l0)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v83 = int32(4548768)
								v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
								*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
								v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
								if l1 != 0 {
									v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
										*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
										F_PopActiveSnapshot(m)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												m.G0 = v8 + int32(1056)
												return
											}
										}
									}
								} else {
									v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
										v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
											v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
											v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
											F_subxact_info_read(m, v117, l0)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return
											} else {
												F_PopActiveSnapshot(m)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _consts[531]))
					*(*int32)(unsafe.Add(mBase, _consts[10])) = v41
					v44 = F_palloc(m, int32(44))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _consts[507]))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+60)) = v44
						F_FileSetInit(m, v44)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[10])) = v34
							v54 = *(*int32)(unsafe.Add(mBase, _consts[507]))
							v56 = v54
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
							v66 = F_pg_snprintf(m, v8+int32(32), int32(1024), int32(178142), v8+int32(16))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								v70 = F_errstart(m, int32(14), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									if v70 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(32)
										F_errmsg_internal(m, int32(178362), v8)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errfinish(m, int32(515057), int32(4339), int32(402103))
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												v83 = int32(4548768)
												v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
												v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
												v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
												if l1 != 0 {
													v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
														*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
														F_PopActiveSnapshot(m)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
														v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
															v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
															v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
															F_subxact_info_read(m, v117, l0)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v123 = m.ExcPending
																if v123 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v125 = m.ExcPending
																	if v125 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v83 = int32(4548768)
										v84 = *(*int32)(unsafe.Add(mBase, _consts[10]))
										v87 = *(*int32)(unsafe.Add(mBase, _consts[529]))
										*(*int32)(unsafe.Add(mBase, _consts[10])) = v87
										v90 = *(*int32)(unsafe.Add(mBase, _consts[507]))
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
										if l1 != 0 {
											v94 = F_BufFileCreateFileSet(m, v91, v8+int32(32))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
												*(*int32)(unsafe.Add(mBase, _consts[530])) = v94
												F_PopActiveSnapshot(m)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v105 = F_BufFileOpenFileSet(m, v91, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[530])) = v105
												v111 = F_BufFileSeek(m, v105, int32(0), int64(0), int32(2))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[10])) = v84
													v116 = *(*int32)(unsafe.Add(mBase, _consts[507]))
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
													F_subxact_info_read(m, v117, l0)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
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
			}
		}
	}
}
func F_stream_stop_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(4541672)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v8 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(244057)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(993)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v8 + int32(16)
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+147)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v34 == v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(524306)
				F_errmsg(m, int32(331621), v8)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_errfinish(m, int32(517668), int32(1358), int32(227271))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v34].(func(*base.Module, int32, int32))(m, v12, l1)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			*(*int32)(unsafe.Add(mBase, _consts[142])) = v57
			m.G0 = v8 + int32(32)
			return
		}
	}
}
