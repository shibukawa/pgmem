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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v406 int32
	_ = v406
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int64
	_ = v447
	var v449 int64
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v503 int64
	_ = v503
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v547 int64
	_ = v547
	var v554 int32
	_ = v554
	var v555 int64
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int64
	_ = v566
	var v581 int32
	_ = v581
	var v585 int64
	_ = v585
	var v589 int64
	_ = v589
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v756 int32
	_ = v756
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	if v31 != int32(1) {
		v51 = v30
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v52 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v34 == v35 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v39 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v34
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if int32(0) < v46 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v42 = v30
	goto L8
L7:
	;
	v42 = v30 | int32(2)
	goto L8
L8:
	;
	v51 = v42
	goto L1
L9:
	;
	v49 = v30 | int32(2)
	goto L11
L10:
	;
	v49 = v30
	goto L11
L11:
	;
	v51 = v49
	goto L1
L12:
	;
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+10)))
	v97 = v95 + v96
	if v97 != 0 {
		goto L29
	} else {
		goto L30
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[0]))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[1]))
	goto L16
L14:
	;
	goto L15
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[2]))
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[3]))
	v77 = v72 - v74 - int32(8)
	if base.Ui32(v77) <= base.Ui32(v72) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	if base.Ui32(int32(_a_F_read_stream_start_pending_read_0)) < base.Ui32(v57-v59) {
		v95 = int32(_a_F_read_stream_start_pending_read_1)
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[0]))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[1]))
	goto L18
L18:
	;
	v95 = v64 - v66
	goto L12
L19:
	;
	if base.Ui32(int32(_a_F_read_stream_start_pending_read_0)) < base.Ui32(v80) {
		v95 = int32(_a_F_read_stream_start_pending_read_1)
		goto L12
	} else {
		goto L23
	}
L20:
	;
	v80 = v77
	goto L22
L21:
	;
	v80 = int32(0)
	goto L22
L22:
	;
	goto L19
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[2]))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[3]))
	v91 = v86 - v88 - int32(8)
	if base.Ui32(v91) <= base.Ui32(v86) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v95 = v94
	goto L12
L25:
	;
	v94 = v91
	goto L27
L26:
	;
	v94 = int32(0)
	goto L27
L27:
	;
	goto L24
L28:
	;
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+52)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v106
	if v105 < v106 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v98 = int32(_a_F_read_stream_start_pending_read_1)
	if v98 <= v97 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v105 = base.B2i32(v102 == int32(0))
	goto L28
L32:
	;
	v101 = v98
	goto L34
L33:
	;
	v101 = v97
	goto L34
L34:
	;
	v105 = v101
	goto L28
L35:
	;
	m.G0 = v28 + int32(16)
	return v951
L36:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	v111 = base.I32_extend16_s(v109 + v105)
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v111 < v112 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v120 = v106
	goto L38
L38:
	;
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+70)))
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+76)))
	v123 = v120 + v122
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v124 < v123 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v111)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v115 = int32(0)
	if v115 < v109 {
		v951 = v115
		goto L35
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v105
	v120 = v105
	goto L38
L43:
	;
	v129 = v124
	goto L46
L44:
	;
	goto L45
L45:
	;
	v189 = v121 * int32(84)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v191 = v189 + v190
	v193 = v191 + int32(4)
	v195 = l0 + int32(80)
	v198 = v195 + v122<<(uint(int32(2))%32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v200 = int32(0)
	v201 = m.G0
	v203 = v201 + int32(-64)
	m.G0 = v203
	v206 = v28 + int32(12)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v207 <= v200 {
		v731 = v207
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v154 = v129 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v154)
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(80)+v129<<(uint(int32(2))%32)))) = int32(0)
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v161 < v123 {
		v129 = v161
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
	m.G0 = v203 - int32(-64)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v815 = v813 + v814
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v815)
	if v789 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v731
	v756 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v193)+34)) = uint16(v756)
	*(*uint16)(unsafe.Add(mBase, uint32(v193)+32)) = uint16(v731)
	*(*int32)(unsafe.Add(mBase, uint32(v193)+28)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v193)+24)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v193)+20)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v191+int32(40)))) = int32(-1)
	goto L158
L51:
	;
	v211 = v207
	v214 = v200
	goto L52
L52:
	;
	v237 = v198 + v214<<(uint(int32(2))%32)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v238 != 0 {
		goto L58
	} else {
		goto L59
	}
L53:
	;
	v731 = v725
	goto L50
L54:
	;
	if base.B2i32(v211 < int32(2))|v214 != 0 {
		v725 = v211
		goto L148
	} else {
		goto L149
	}
L55:
	;
	if v214 != 0 {
		goto L145
	} else {
		goto L146
	}
L56:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)+24))
	v631 = int32(base.Ui32(v627)>>(uint(int32(24))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)) = uint8(v631)
	if v631 == int32(0) {
		goto L54
	} else {
		goto L144
	}
L57:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v626 = v620 + v238<<(uint(int32(6))%32) + int32(-64)
	goto L56
L58:
	;
	if int32(0) <= v238 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v248 = v214 + v199
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+8)))
	if v252 == int32(116) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[5]))
	v626 = v242 + (v238^int32(-1))<<(uint(int32(6))%32)
	goto L56
L62:
	;
	if v251 != 0 {
		goto L122
	} else {
		goto L123
	}
L63:
	;
	v255 = int32(1)
	v258 = F_LocalBufferAlloc(m, v250, v249, v248, v201+int32(-45))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	v275 = F_IOContextForStrategy(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L66
	} else {
		goto L69
	}
L66:
	;
	return int32(0)
L67:
	;
	v262 = int32(3)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)))
	if v264 != int32(1) {
		v510 = v258
		v514 = int32(0)
		v519 = v255
		v527 = v262
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v267 = int32(_a_F_read_stream_start_pending_read_2)
	v269 = *(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[6])) = v269 + int64(1)
	v510 = v258
	v514 = int32(1)
	v519 = v255
	v527 = v262
	goto L62
L69:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[7]))
	F_ResourceOwnerEnlarge(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+20)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+24)) = v285
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+36)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v203)+32)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v203)+28)) = v287
	v292 = v201 + int32(-44)
	v293 = F_BufTableHashCode(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[8]))
	v303 = v296 + v293&int32(127)<<(uint(int32(7))%32) + int32(_a_F_read_stream_start_pending_read_3)
	v305 = F_LWLockAcquire(m, v303, int32(1))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L66
	} else {
		goto L73
	}
L73:
	;
	v307 = F_BufTableLookup(m, v292, v293)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L66
	} else {
		goto L75
	}
L74:
	;
	v501 = int32(_a_F_read_stream_start_pending_read_4)
	v503 = *(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[9])) = v503 + int64(1)
	v510 = v494
	v514 = int32(1)
	v519 = int32(0)
	v527 = v275
	goto L62
L75:
	;
	if int32(0) <= v307 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v316 = v313 + v307<<(uint(int32(6))%32)
	v317 = F_PinBuffer(m, v316, v274)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L66
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_LWLockRelease(m, v303)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L66
	} else {
		goto L82
	}
L79:
	;
	F_LWLockRelease(m, v303)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L66
	} else {
		goto L80
	}
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)) = uint8(v317)
	if v317 != 0 {
		v494 = v316
		goto L74
	} else {
		goto L81
	}
L81:
	;
	v510 = v316
	v514 = int32(0)
	v519 = int32(0)
	v527 = v275
	goto L62
L82:
	;
	v325 = F_GetVictimBuffer(m, v274, v275)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L66
	} else {
		goto L83
	}
L83:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v330 = F_LWLockAcquire(m, v303, int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v334 = v328 + v325<<(uint(int32(6))%32)
	v336 = v334 + int32(-64)
	v340 = v334 - int32(44)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v342 = F_BufTableInsert(m, v201+int32(-44), v293, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L66
	} else {
		goto L85
	}
L85:
	;
	if v342 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+60)) = int32(_a_F_read_stream_start_pending_read_5)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+56)) = int32(_a_F_read_stream_start_pending_read_6)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+52)) = int32(_a_F_read_stream_start_pending_read_7)
	v352 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+48)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v203)+40)) = int64(0)
	v357 = v334 - int32(40)
	v358 = int32(_a_F_read_stream_start_pending_read_8)
	v360 = base.AtomicRmwOr32(m, v357, v352, v358)
	if v360&v358 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[7]))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	F_ResourceOwnerForget(m, v472, v473+int32(1), int32(_a_F_read_stream_start_pending_read_9))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L66
	} else {
		goto L114
	}
L89:
	;
	goto L92
L90:
	;
	v406 = v360
	goto L91
L91:
	;
	v425 = int32(_a_F_read_stream_start_pending_read_10)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[10]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v201+int32(-24))+8))
	if v428 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L92:
	;
	F_perform_spin_delay(m, v201+int32(-24))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L66
	} else {
		goto L94
	}
L93:
	;
	v406 = v394
	goto L91
L94:
	;
	v392 = int32(_a_F_read_stream_start_pending_read_8)
	v394 = base.AtomicRmwOr32(m, v357, int32(0), v392)
	if v394&v392 != 0 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v203)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v336)+16)) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v203)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v336)+8)) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v203)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v336))) = v449
	v453 = int32(-2113667072)
	if v249 == int32(3) {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[10])) = v443
	goto L97
L99:
	;
	if int32(999) < v426 {
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v426 < int32(11) {
		goto L97
	} else {
		goto L106
	}
L102:
	;
	v433 = int32(900)
	if v433 <= v426 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v436 = v433
	goto L105
L104:
	;
	v436 = v426
	goto L105
L105:
	;
	v443 = v436 + int32(100)
	goto L98
L106:
	;
	v443 = v426 - int32(1)
	goto L98
L107:
	;
	v458 = v453
	goto L109
L108:
	;
	v458 = int32(33816576)
	goto L109
L109:
	;
	if v252 == int32(112) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v461 = v453
	goto L112
L111:
	;
	v461 = v458
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v406&int32(-38010881) | v461
	F_LWLockRelease(m, v303)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L66
	} else {
		goto L113
	}
L113:
	;
	v466 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)) = uint8(v466)
	v510 = v336
	v514 = v466
	v519 = v466
	v527 = v275
	goto L62
L114:
	;
	F_UnpinBufferNoOwner(m, v336)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L66
	} else {
		goto L115
	}
L115:
	;
	F_StrategyFreeBuffer(m, v336)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L66
	} else {
		goto L116
	}
L116:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[4]))
	v487 = v484 + v342<<(uint(int32(6))%32)
	v488 = F_PinBuffer(m, v487, v274)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L66
	} else {
		goto L117
	}
L117:
	;
	F_LWLockRelease(m, v303)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L66
	} else {
		goto L118
	}
L118:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)) = uint8(v488)
	if v488 != 0 {
		v494 = v487
		goto L74
	} else {
		goto L119
	}
L119:
	;
	v510 = v487
	v514 = int32(0)
	v519 = int32(0)
	v527 = v275
	goto L62
L120:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v510)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v614 + int32(1)
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)))
	if v618 != 0 {
		goto L55
	} else {
		goto L143
	}
L121:
	;
	v581 = v519*int32(320) + v527<<(uint(int32(6))%32)
	v585 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_c_F_read_stream_start_pending_read[11])))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_c_F_read_stream_start_pending_read[11]))) = v585 + int64(1)
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_c_F_read_stream_start_pending_read[12])))
	*(*int64)(unsafe.Add(mBase, uint32(v581)+uint32(_c_F_read_stream_start_pending_read[12]))) = v589
	v591 = int32(1)
	F_pgstat_count_backend_io_op(m, v519, v527, int32(2), v591, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[13])) = uint8(v591)
	*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[14])) = uint8(v591)
	goto L141
L122:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v251)+272))
	if v533 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	goto L124
L124:
	;
	if v514 == int32(0) {
		goto L120
	} else {
		goto L140
	}
L125:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v251)+272))
	if v554 != 0 {
		goto L134
	} else {
		goto L135
	}
L126:
	;
	if v514 == int32(0) {
		goto L120
	} else {
		goto L133
	}
L127:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+268)))
	if v536 != int32(1) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v547 = *(*int64)(unsafe.Add(mBase, uint32(v533)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v533)+112)) = v547 + int64(1)
	goto L126
L130:
	;
	F_pgstat_assoc_relation(m, v251)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L66
	} else {
		goto L131
	}
L131:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v251)+272))
	v543 = *(*int64)(unsafe.Add(mBase, uint32(v542)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v542)+112)) = v543 + int64(1)
	if v541 != 0 {
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
	v555 = *(*int64)(unsafe.Add(mBase, uint32(v554)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v554)+120)) = v555 + int64(1)
	goto L121
L135:
	;
	goto L136
L136:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+268)))
	if v559 != int32(1) {
		goto L121
	} else {
		goto L137
	}
L137:
	;
	F_pgstat_assoc_relation(m, v251)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L66
	} else {
		goto L138
	}
L138:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+19)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v251)+272))
	v566 = *(*int64)(unsafe.Add(mBase, uint32(v565)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v565)+120)) = v566 + int64(1)
	if v564 != 0 {
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
	v601 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[15])))
	if v601 != int32(1) {
		goto L120
	} else {
		goto L142
	}
L142:
	;
	v604 = int32(_a_F_read_stream_start_pending_read_11)
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[16]))
	v608 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[17]))
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[16])) = v606 + v608
	goto L120
L143:
	;
	goto L54
L144:
	;
	goto L55
L145:
	;
	v731 = v214
	goto L50
L146:
	;
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = int32(1)
	v789 = int32(0)
	goto L49
L148:
	;
	v728 = v214 + int32(1)
	if v728 < v725 {
		v211 = v725
		v214 = v728
		goto L52
	} else {
		goto L157
	}
L149:
	;
	v693 = int32(_a_F_read_stream_start_pending_read_12)
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18]))
	v696 = int32(1)
	v697 = v695 + v696
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18])) = v697
	*(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[18])) = v697 - v696
	v708 = int32(_a_F_read_stream_start_pending_read_13) - v199&int32(_a_F_read_stream_start_pending_read_14)
	if v211 <= v708 {
		v725 = v211
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v712 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L66
	} else {
		goto L151
	}
L151:
	;
	if v712 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v199
	F_errmsg_internal(m, int32(_a_F_read_stream_start_pending_read_15), v203)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L66
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v725 = v708
	goto L148
L155:
	;
	F_errfinish(m, int32(_a_F_read_stream_start_pending_read_7), int32(1382), int32(_a_F_read_stream_start_pending_read_16))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
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
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_start_pending_read[19]))
	if v767 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v768 = F_AsyncReadBuffers(m, v193, v206)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L66
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+28)) = v772 | int32(8)
	v776 = int32(1)
	if v51&int32(2) == int32(0) {
		v789 = v776
		goto L49
	} else {
		goto L163
	}
L162:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	*(*uint16)(unsafe.Add(mBase, uint32(v193)+32)) = uint16(v770)
	v789 = v768
	goto L49
L163:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v783 = F_smgrprefetch(m, v781, v782, v199, v731)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L66
	} else {
		goto L164
	}
L164:
	;
	v789 = v776
	goto L49
L165:
	;
	v848 = int32(0)
	if v120 <= v847 {
		v892 = v848
		goto L173
	} else {
		goto L174
	}
L166:
	;
	v819 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v819 < int32(2) {
		v847 = v813
		goto L165
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*uint16)(unsafe.Add(mBase, uint32(v825+v189))) = uint16(v122)
	v828 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v829 = int32(1)
	v830 = v828 + v829
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v830)
	v832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)))
	v834 = v832 + v829
	v836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v836 != v834&int32(_a_F_read_stream_start_pending_read_17) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v823 = v819 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v823)
	v847 = v813
	goto L165
L170:
	;
	v840 = v834
	goto L172
L171:
	;
	v840 = int32(0)
	goto L172
L172:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v840)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v842 + v843
	v847 = v842
	goto L165
L173:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v892)
	v920 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v922 = base.I32_extend16_s(v122 + (v892 + v847) - v920)
	if v922 <= int32(0) {
		goto L179
	} else {
		goto L180
	}
L174:
	;
	v850 = int32(2)
	v856 = v120 - v847
	v858 = v848
	goto L175
L175:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v195+v847<<(uint(v850)%32)+v122<<(uint(v850)%32)+v858<<(uint(int32(2))%32))))
	if v885 == int32(0) {
		v892 = v858
		goto L173
	} else {
		goto L177
	}
L176:
	;
	v892 = v856
	goto L173
L177:
	;
	v889 = v858 + int32(1)
	if v889 != v856 {
		v858 = v889
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v934 + v847
	v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v938 = v937 - v847
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v938)
	v940 = v847 + v122
	if v920 <= base.I32_extend16_s(v940) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v926 = v922 << (uint(int32(2)) % 32)
	if v926 == int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	base.MemoryCopy(m, v195, v195+v920<<(uint(int32(2))%32), v926)
	goto L179
L182:
	;
	v944 = v920
	goto L184
L183:
	;
	v944 = int32(0)
	goto L184
L184:
	;
	v945 = v940 - v944
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v945)
	v951 = int32(1)
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
