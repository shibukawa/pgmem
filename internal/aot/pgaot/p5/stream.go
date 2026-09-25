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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
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
	v77 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v77)
	return
L10:
	;
	v28 = l0 + int32(80)
	v30 = v24
	goto L11
L11:
	;
	v36 = v30 << (uint(int32(2)) % 32)
	v37 = v28 + v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v43 = v41 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v43)
	F_ReleaseBuffer(m, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(0)
	v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	if v30 < v49-int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v28+v53<<(uint(int32(2))%32)+v36))) = int32(0)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v61 = v30 + int32(1)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v63 != v61&int32(_a_F_read_stream_reset_0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = v61
	goto L20
L19:
	;
	v67 = int32(0)
	goto L20
L20:
	;
	v68 = base.I32_extend16_s(v67)
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v68 < v69 {
		v30 = v68
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int64
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v413 int32
	_ = v413
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v537 int32
	_ = v537
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int64
	_ = v554
	var v558 int64
	_ = v558
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int64
	_ = v577
	var v592 int32
	_ = v592
	var v596 int64
	_ = v596
	var v600 int64
	_ = v600
	var v602 int32
	_ = v602
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v770 int32
	_ = v770
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v968 int32
	_ = v968
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
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[0]))
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[1]))
	goto L16
L14:
	;
	goto L15
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[2]))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[3]))
	v78 = v73 - v75 - int32(8)
	if base.Ui32(v78) <= base.Ui32(v73) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	if base.Ui32(int32(_a_F_read_stream_start_pending_read_0)) < base.Ui32(v58-v60) {
		v96 = int32(_a_F_read_stream_start_pending_read_1)
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[0]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[1]))
	goto L18
L18:
	;
	v96 = v65 - v67
	goto L12
L19:
	;
	if base.Ui32(int32(_a_F_read_stream_start_pending_read_0)) < base.Ui32(v81) {
		v96 = int32(_a_F_read_stream_start_pending_read_1)
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
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[2]))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[3]))
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
	v99 = int32(_a_F_read_stream_start_pending_read_1)
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
	return v968
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
	v121 = v107
	goto L38
L38:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+70)))
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+76)))
	v124 = v121 + v123
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v125 < v124 {
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
		v968 = v116
		goto L35
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v106
	v121 = v106
	goto L38
L43:
	;
	v130 = v125
	goto L46
L44:
	;
	goto L45
L45:
	;
	v192 = v122 * int32(84)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v194 = v192 + v193
	v196 = v194 + int32(4)
	v198 = l0 + int32(80)
	v201 = v198 + v123<<(uint(int32(2))%32)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v203 = int32(0)
	v204 = m.G0
	v206 = v204 + int32(-64)
	m.G0 = v206
	v209 = v29 + int32(12)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v210 <= v203 {
		v744 = v210
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v156 = v130 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v156)
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(80)+v130<<(uint(int32(2))%32)))) = int32(0)
	v163 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v163 < v124 {
		v130 = v163
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
	m.G0 = v206 - int32(-64)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v830 = v828 + v829
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v830)
	if v803 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v744
	v770 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+34)) = uint16(v770)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+32)) = uint16(v744)
	*(*int32)(unsafe.Add(mBase, uint32(v196)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v196)+24)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v196)+20)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v194+int32(40)))) = int32(-1)
	goto L158
L51:
	;
	v214 = v210
	v217 = v203
	goto L52
L52:
	;
	v241 = v201 + v217<<(uint(int32(2))%32)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v242 != 0 {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	v744 = v738
	goto L50
L54:
	;
	if base.B2i32(v214 < int32(2))|v217 != 0 {
		v738 = v214
		goto L148
	} else {
		goto L149
	}
L55:
	;
	if v217 != 0 {
		goto L145
	} else {
		goto L146
	}
L56:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+24))
	v642 = int32(base.Ui32(v638)>>(uint(int32(24))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)) = uint8(v642)
	if v642 == int32(0) {
		goto L54
	} else {
		goto L144
	}
L57:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v637 = v631 + v242<<(uint(int32(6))%32) + int32(-64)
	goto L56
L58:
	;
	if int32(0) <= v242 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v252 = v217 + v202
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+8)))
	if v256 == int32(116) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[5]))
	v637 = v246 + (v242^int32(-1))<<(uint(int32(6))%32)
	goto L56
L62:
	;
	if v255 != 0 {
		goto L122
	} else {
		goto L123
	}
L63:
	;
	v259 = int32(1)
	v262 = F_LocalBufferAlloc(m, v254, v253, v252, v204+int32(-45))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
	v279 = F_IOContextForStrategy(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L66
	} else {
		goto L69
	}
L66:
	;
	return int32(0)
L67:
	;
	v266 = int32(3)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)))
	if v268 != int32(1) {
		v520 = v262
		v523 = int32(0)
		v528 = v259
		v537 = v266
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v271 = int32(_a_F_read_stream_start_pending_read_2)
	v273 = *(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[6])) = v273 + int64(1)
	v520 = v262
	v523 = int32(1)
	v528 = v259
	v537 = v266
	goto L62
L69:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[7]))
	F_ResourceOwnerEnlarge(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+24)) = v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v254)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+36)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v206)+32)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v206)+28)) = v291
	v296 = v204 + int32(-44)
	v297 = F_BufTableHashCode(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[8]))
	v307 = v300 + v297&int32(127)<<(uint(int32(7))%32) + int32(_a_F_read_stream_start_pending_read_3)
	v309 = F_LWLockAcquire(m, v307, int32(1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L66
	} else {
		goto L73
	}
L73:
	;
	v311 = F_BufTableLookup(m, v296, v297)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L66
	} else {
		goto L75
	}
L74:
	;
	v511 = int32(_a_F_read_stream_start_pending_read_4)
	v513 = *(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9])) = v513 + int64(1)
	v520 = v504
	v523 = int32(1)
	v528 = int32(0)
	v537 = v279
	goto L62
L75:
	;
	if int32(0) <= v311 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v320 = v317 + v311<<(uint(int32(6))%32)
	v321 = F_PinBuffer(m, v320, v278)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L66
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_LWLockRelease(m, v307)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L66
	} else {
		goto L82
	}
L79:
	;
	F_LWLockRelease(m, v307)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L66
	} else {
		goto L80
	}
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)) = uint8(v321)
	if v321 != 0 {
		v504 = v320
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v520 = v320
	v523 = int32(0)
	v528 = int32(0)
	v537 = v279
	goto L62
L82:
	;
	v329 = F_GetVictimBuffer(m, v278, v279)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L66
	} else {
		goto L83
	}
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v334 = F_LWLockAcquire(m, v307, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v338 = v332 + v329<<(uint(int32(6))%32)
	v340 = v338 + int32(-64)
	v344 = v338 - int32(44)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v346 = F_BufTableInsert(m, v204+int32(-44), v297, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L66
	} else {
		goto L85
	}
L85:
	;
	if v346 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+60)) = int32(_a_F_read_stream_start_pending_read_5)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+56)) = int32(_a_F_read_stream_start_pending_read_6)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+52)) = int32(_a_F_read_stream_start_pending_read_7)
	v356 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+48)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v206)+40)) = int64(0)
	v361 = v338 - int32(40)
	v362 = int32(_a_F_read_stream_start_pending_read_8)
	v364 = base.AtomicRmwOr32(m, v361, v356, v362)
	if v364&v362 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[7]))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	F_ResourceOwnerForget(m, v482, v483+int32(1), int32(_a_F_read_stream_start_pending_read_9))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L66
	} else {
		goto L114
	}
L89:
	;
	goto L92
L90:
	;
	v413 = v364
	goto L91
L91:
	;
	v431 = int32(_a_F_read_stream_start_pending_read_10)
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[10]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v204+int32(-24))+8))
	if v434 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	F_perform_spin_delay(m, v204+int32(-24))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L66
	} else {
		goto L94
	}
L93:
	;
	v413 = v399
	goto L91
L94:
	;
	v397 = int32(_a_F_read_stream_start_pending_read_8)
	v399 = base.AtomicRmwOr32(m, v361, int32(0), v397)
	if v399&v397 != 0 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+16)) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v206)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v340)+8)) = v453
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v206)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = v455
	v457 = int32(0)
	v461 = base.AtomicRmwOr32(m, v457, int32(_a_F_read_stream_start_pending_read_11), v457)
	v464 = int32(-2113667072)
	if v253 == int32(3) {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[10])) = v449
	goto L97
L99:
	;
	if int32(999) < v432 {
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v432 < int32(11) {
		goto L97
	} else {
		goto L106
	}
L102:
	;
	v439 = int32(900)
	if v439 <= v432 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v442 = v439
	goto L105
L104:
	;
	v442 = v432
	goto L105
L105:
	;
	v449 = v442 + int32(100)
	goto L98
L106:
	;
	v449 = v432 - int32(1)
	goto L98
L107:
	;
	v469 = v464
	goto L109
L108:
	;
	v469 = int32(33816576)
	goto L109
L109:
	;
	if v256 == int32(112) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v472 = v464
	goto L112
L111:
	;
	v472 = v469
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = v413&int32(-38010881) | v472
	F_LWLockRelease(m, v307)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L66
	} else {
		goto L113
	}
L113:
	;
	v477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)) = uint8(v477)
	v520 = v340
	v523 = v457
	v528 = v477
	v537 = v279
	goto L62
L114:
	;
	F_UnpinBufferNoOwner(m, v340)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L66
	} else {
		goto L115
	}
L115:
	;
	F_StrategyFreeBuffer(m, v340)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L66
	} else {
		goto L116
	}
L116:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v497 = v494 + v346<<(uint(int32(6))%32)
	v498 = F_PinBuffer(m, v497, v278)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L66
	} else {
		goto L117
	}
L117:
	;
	F_LWLockRelease(m, v307)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L66
	} else {
		goto L118
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)) = uint8(v498)
	if v498 != 0 {
		v504 = v497
		goto L74
	} else {
		goto L119
	}
L119:
	;
	v520 = v497
	v523 = int32(0)
	v528 = int32(0)
	v537 = v279
	goto L62
L120:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v520)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = v625 + int32(1)
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)))
	if v629 != 0 {
		goto L55
	} else {
		goto L143
	}
L121:
	;
	v592 = v528*int32(320) + v537<<(uint(int32(6))%32)
	v596 = *(*int64)(unsafe.Add(mBase, uint32(v592)+uint32(_c_F_read_stream_start_pending_read[11])))
	*(*int64)(unsafe.Add(mBase, uint32(v592)+uint32(_c_F_read_stream_start_pending_read[11]))) = v596 + int64(1)
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v592)+uint32(_c_F_read_stream_start_pending_read[12])))
	*(*int64)(unsafe.Add(mBase, uint32(v592)+uint32(_c_F_read_stream_start_pending_read[12]))) = v600
	v602 = int32(1)
	F_pgstat_count_backend_io_op(m, v528, v537, int32(2), v602, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[13])) = uint8(v602)
	*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[14])) = uint8(v602)
	goto L141
L122:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v255)+272))
	if v544 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	goto L124
L124:
	;
	if v523 == int32(0) {
		goto L120
	} else {
		goto L140
	}
L125:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v255)+272))
	if v565 != 0 {
		goto L134
	} else {
		goto L135
	}
L126:
	;
	if v523 == int32(0) {
		goto L120
	} else {
		goto L133
	}
L127:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+268)))
	if v547 != int32(1) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v558 = *(*int64)(unsafe.Add(mBase, uint32(v544)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+112)) = v558 + int64(1)
	goto L126
L130:
	;
	F_pgstat_assoc_relation(m, v255)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L66
	} else {
		goto L131
	}
L131:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v255)+272))
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v553)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v553)+112)) = v554 + int64(1)
	if v552 != 0 {
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
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v565)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+120)) = v566 + int64(1)
	goto L121
L135:
	;
	goto L136
L136:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+268)))
	if v570 != int32(1) {
		goto L121
	} else {
		goto L137
	}
L137:
	;
	F_pgstat_assoc_relation(m, v255)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L66
	} else {
		goto L138
	}
L138:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+19)))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v255)+272))
	v577 = *(*int64)(unsafe.Add(mBase, uint32(v576)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v576)+120)) = v577 + int64(1)
	if v575 != 0 {
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
	v612 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[15])))
	if v612 != int32(1) {
		goto L120
	} else {
		goto L142
	}
L142:
	;
	v615 = int32(_a_F_read_stream_start_pending_read_12)
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[16]))
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[16])) = v617 + v619
	goto L120
L143:
	;
	goto L54
L144:
	;
	goto L55
L145:
	;
	v744 = v217
	goto L50
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = int32(1)
	v803 = int32(0)
	goto L49
L148:
	;
	v741 = v217 + int32(1)
	if v741 < v738 {
		v214 = v738
		v217 = v741
		goto L52
	} else {
		goto L157
	}
L149:
	;
	v706 = int32(_a_F_read_stream_start_pending_read_13)
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18]))
	v709 = int32(1)
	v710 = v708 + v709
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18])) = v710
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18])) = v710 - v709
	v721 = int32(_a_F_read_stream_start_pending_read_14) - v202&int32(_a_F_read_stream_start_pending_read_15)
	if v214 <= v721 {
		v738 = v214
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v725 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L66
	} else {
		goto L151
	}
L151:
	;
	if v725 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+8)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v202
	F_errmsg_internal(m, int32(_a_F_read_stream_start_pending_read_16), v206)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L66
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v738 = v721
	goto L148
L155:
	;
	F_errfinish(m, int32(_a_F_read_stream_start_pending_read_7), int32(1382), int32(_a_F_read_stream_start_pending_read_17))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L66
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	goto L53
L158:
	;
	v781 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[19]))
	if v781 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v782 = F_AsyncReadBuffers(m, v196, v209)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L66
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v196)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+28)) = v786 | int32(8)
	v790 = int32(1)
	if v52&int32(2) == int32(0) {
		v803 = v790
		goto L49
	} else {
		goto L163
	}
L162:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+32)) = uint16(v784)
	v803 = v782
	goto L49
L163:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v797 = F_smgrprefetch(m, v795, v796, v202, v744)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L66
	} else {
		goto L164
	}
L164:
	;
	v803 = v790
	goto L49
L165:
	;
	v863 = int32(0)
	if v121 <= v862 {
		v908 = v863
		goto L173
	} else {
		goto L174
	}
L166:
	;
	v834 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v834 < int32(2) {
		v862 = v828
		goto L165
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v840+v192))) = uint16(v123)
	v843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v844 = int32(1)
	v845 = v843 + v844
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v845)
	v847 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)))
	v849 = v847 + v844
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v851 != v849&int32(_a_F_read_stream_start_pending_read_18) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v838 = v834 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v838)
	v862 = v828
	goto L165
L170:
	;
	v855 = v849
	goto L172
L171:
	;
	v855 = int32(0)
	goto L172
L172:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v855)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v857 + v858
	v862 = v857
	goto L165
L173:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v908)
	v937 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v939 = base.I32_extend16_s(v123 + (v908 + v862) - v937)
	if v939 <= int32(0) {
		goto L179
	} else {
		goto L180
	}
L174:
	;
	v865 = int32(2)
	v871 = v121 - v862
	v873 = v863
	goto L175
L175:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v198+v862<<(uint(v865)%32)+v123<<(uint(v865)%32)+v873<<(uint(int32(2))%32))))
	if v901 == int32(0) {
		v908 = v873
		goto L173
	} else {
		goto L177
	}
L176:
	;
	v908 = v871
	goto L173
L177:
	;
	v905 = v873 + int32(1)
	if v905 != v871 {
		v873 = v905
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v951 + v862
	v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v955 = v954 - v862
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v955)
	v957 = v862 + v123
	if v937 <= base.I32_extend16_s(v957) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v943 = v939 << (uint(int32(2)) % 32)
	if v943 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	base.MemoryCopy(m, v198, v198+v937<<(uint(int32(2))%32), v943)
	goto L179
L182:
	;
	v961 = v937
	goto L184
L183:
	;
	v961 = int32(0)
	goto L184
L184:
	;
	v962 = v957 - v961
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v962)
	v968 = int32(1)
	goto L35
}
func F_stream_cleanup_files(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(1056)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
	v12 = v7 + int32(32)
	v17 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_stream_cleanup_files_0), v7+int32(16))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_stream_cleanup_files[0]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
		F_BufFileDeleteFileSet(m, v21, v12, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			v29 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_stream_cleanup_files_1), v7)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_stream_cleanup_files[0]))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
				F_BufFileDeleteFileSet(m, v33, v12, int32(1))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					m.G0 = v7 + int32(1056)
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
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v6 = m.G0
	v8 = v6 - int32(1056)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[0]))
	if v11 < int32(0) {
		v15 = F_GetCurrentTimestamp(m)
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, _c_F_stream_start_internal[1])) = v15
	} else {
	}
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[2]))
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
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
						if v38 != 0 {
							v56 = v37
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
							v61 = v8 + int32(32)
							v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
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
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
										F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v81 = int32(_a_F_stream_start_internal_5)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
												if l1 != 0 {
													v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
														v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
															v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
															F_subxact_info_read(m, v115, l0)
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
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
										v81 = int32(_a_F_stream_start_internal_5)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
										if l1 != 0 {
											v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
												v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
													F_subxact_info_read(m, v115, l0)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
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
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[8]))
							*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v41
							v44 = F_palloc(m, int32(44))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v47)+60)) = v44
								F_FileSetInit(m, v44)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
									v54 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
									v56 = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
									v61 = v8 + int32(32)
									v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
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
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
												F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														v81 = int32(_a_F_stream_start_internal_5)
														v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
														v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
														v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
														v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
														if l1 != 0 {
															v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
																*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(1056)
																		return
																	}
																}
															}
														} else {
															v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
																v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
																	v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
																	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
																	F_subxact_info_read(m, v115, l0)
																	mBase = m.M
																	v117 = m.ExcPending
																	if v117 != 0 {
																		return
																	} else {
																		F_PopActiveSnapshot(m)
																		mBase = m.M
																		v121 = m.ExcPending
																		if v121 != 0 {
																			return
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v123 = m.ExcPending
																			if v123 != 0 {
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
												v81 = int32(_a_F_stream_start_internal_5)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
												if l1 != 0 {
													v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
														v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
															v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
															F_subxact_info_read(m, v115, l0)
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
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
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[3]))
				*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
				if v38 != 0 {
					v56 = v37
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
					v61 = v8 + int32(32)
					v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
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
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
								F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										v81 = int32(_a_F_stream_start_internal_5)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
										if l1 != 0 {
											v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
												v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
													F_subxact_info_read(m, v115, l0)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
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
								v81 = int32(_a_F_stream_start_internal_5)
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
								v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
								*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
								v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
								if l1 != 0 {
									v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
										F_PopActiveSnapshot(m)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												m.G0 = v8 + int32(1056)
												return
											}
										}
									}
								} else {
									v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
										v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
											v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
											F_subxact_info_read(m, v115, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
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
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[8]))
					*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v41
					v44 = F_palloc(m, int32(44))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+60)) = v44
						F_FileSetInit(m, v44)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v34
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
							v56 = v54
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l0
							v61 = v8 + int32(32)
							v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_stream_start_internal_0), v8+int32(16))
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
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
										F_errmsg_internal(m, int32(_a_F_stream_start_internal_1), v8)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_stream_start_internal_2), int32(_a_F_stream_start_internal_3), int32(_a_F_stream_start_internal_4))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v81 = int32(_a_F_stream_start_internal_5)
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
												v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
												if l1 != 0 {
													v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return
															} else {
																m.G0 = v8 + int32(1056)
																return
															}
														}
													}
												} else {
													v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
													mBase = m.M
													v104 = m.ExcPending
													if v104 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
														v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
															v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
															F_subxact_info_read(m, v115, l0)
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return
															} else {
																F_PopActiveSnapshot(m)
																mBase = m.M
																v121 = m.ExcPending
																if v121 != 0 {
																	return
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v123 = m.ExcPending
																	if v123 != 0 {
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
										v81 = int32(_a_F_stream_start_internal_5)
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4]))
										v85 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[6]))
										*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
										if l1 != 0 {
											v92 = F_BufFileCreateFileSet(m, v89, v8+int32(32))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v92
												F_PopActiveSnapshot(m)
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													F_CommandCounterIncrement(m)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														m.G0 = v8 + int32(1056)
														return
													}
												}
											}
										} else {
											v103 = F_BufFileOpenFileSet(m, v89, v8+int32(32), int32(2), int32(0))
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[7])) = v103
												v109 = F_BufFileSeek(m, v103, int32(0), int64(0), int32(2))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[4])) = v82
													v114 = *(*int32)(unsafe.Add(mBase, _c_F_stream_start_internal[5]))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+32))
													F_subxact_info_read(m, v115, l0)
													mBase = m.M
													v117 = m.ExcPending
													if v117 != 0 {
														return
													} else {
														F_PopActiveSnapshot(m)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
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
	v10 = int32(_a_F_stream_stop_cb_wrapper_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_stream_stop_cb_wrapper[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, _c_F_stream_stop_cb_wrapper[0])) = v8 + int32(4)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_stream_stop_cb_wrapper_1)
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
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_stream_stop_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_stream_stop_cb_wrapper_3), v8)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_stream_stop_cb_wrapper_4), int32(1358), int32(_a_F_stream_stop_cb_wrapper_5))
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
			*(*int32)(unsafe.Add(mBase, _c_F_stream_stop_cb_wrapper[0])) = v57
			m.G0 = v8 + int32(32)
			return
		}
	}
}
