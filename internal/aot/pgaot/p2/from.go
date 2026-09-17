package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginCopyFrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v439 int32
	_ = v439
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
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int64
	_ = v481
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v985 int32
	_ = v985
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1113 int64
	_ = v1113
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1276 int64
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	v5 = l4
	v19 = m.G0
	v21 = v19 - int32(272)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+264)) = v24
	v27 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+256)) = v27
	v30 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+240)) = v30
	v33 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = v33
	v36 = *(*int64)(unsafe.Add(mBase, _c_F_BeginCopyFrom[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+224)) = v36
	v39 = F_palloc0(m, int32(352))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(_a_F_BeginCopyFrom_0), int32(0), int32(_a_F_BeginCopyFrom_1), int32(_a_F_BeginCopyFrom_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+204)) = v49
	v52 = int32(_a_F_BeginCopyFrom_3)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5])) = v49
	v57 = v39 + int32(56)
	F_ProcessCopyOptions(m, l0, v57, int32(1), l7)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+62)))
	if v62 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+60)))
	if v67 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v69 = int32(_a_F_BeginCopyFrom_4)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v73 = F_CopyGetAttnums(m, v72, l1, l6)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v68 = int32(_a_F_BeginCopyFrom_5)
	goto L10
L9:
	;
	v68 = int32(_a_F_BeginCopyFrom_6)
	goto L10
L10:
	;
	v69 = v68
	goto L7
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v77 = base.I32_extend16_s(v76)
	v78 = F_palloc0(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v78
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+116)))
	if v81 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v842 = F_palloc0(m, v841)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L208
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L195
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L191
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L187
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L183
	}
L18:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v39)+140))
	if v194 != 0 {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	if v77 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	if v88 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L22:
	;
	base.MemoryFill(m, v78, int32(1), v77)
	goto L18
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v92 = F_CopyGetAttnums(m, v72, v91, v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v92 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v96 <= int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v101 = int32(0)
	goto L27
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v101<<(uint(int32(2))%32))))
	v124 = v122 - int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v127 = int32(0)
	if v126 == v127 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L18
L29:
	;
	if v165 == int32(0) {
		goto L17
	} else {
		goto L42
	}
L30:
	;
	v165 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v133 <= int32(0) {
		v159 = v127
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v165 = v159
	goto L29
L34:
	;
	v136 = int32(0)
	if v136 < v133 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v139 = v133
	goto L37
L36:
	;
	v139 = v136
	goto L37
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v142 = int32(0)
	goto L38
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v140+v142<<(uint(int32(2))%32))))
	v151 = base.B2i32(v150 == v122)
	if v150 == v122 {
		v159 = v151
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v159 = v151
	goto L33
L40:
	;
	v153 = v142 + int32(1)
	if v153 != v139 {
		v142 = v153
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v39)+120))
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v168+v124))) = uint8(v170)
	v173 = v101 + v170
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v173 < v174 {
		v101 = v173
		goto L27
	} else {
		goto L43
	}
L43:
	;
	goto L28
L44:
	;
	v213 = F_palloc0(m, v77)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L50
	}
L45:
	;
	v196 = F_palloc0(m, int32(12))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = int32(0)
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = int32(447)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v39)+220))
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+4)) = uint8(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v39)+140))
	if v204 != int32(1) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v39)+220))
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)) = uint8(v208)
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v213
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+128)))
	if v216 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+136)))
	if v329 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L52:
	;
	if v77 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v39)+124))
	if v223 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L55:
	;
	base.MemoryFill(m, v213, int32(1), v77)
	goto L51
L56:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v227 = F_CopyGetAttnums(m, v72, v226, v223)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v227 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v231 <= int32(0) {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v236 = int32(0)
	goto L60
L60:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253+v236<<(uint(int32(2))%32))))
	v259 = v257 - int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v262 = int32(0)
	if v261 == v262 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L51
L62:
	;
	if v300 == int32(0) {
		goto L16
	} else {
		goto L75
	}
L63:
	;
	v300 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v268 <= int32(0) {
		v294 = v262
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v300 = v294
	goto L62
L67:
	;
	v271 = int32(0)
	if v271 < v268 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v274 = v268
	goto L70
L69:
	;
	v274 = v271
	goto L70
L70:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v277 = int32(0)
	goto L71
L71:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v275+v277<<(uint(int32(2))%32))))
	v286 = base.B2i32(v285 == v257)
	if v285 == v257 {
		v294 = v286
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v294 = v286
	goto L66
L73:
	;
	v288 = v277 + int32(1)
	if v288 != v274 {
		v277 = v288
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v39)+132))
	v305 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v303+v259))) = uint8(v305)
	v308 = v236 + v305
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v308 < v309 {
		v236 = v308
		goto L60
	} else {
		goto L76
	}
L76:
	;
	goto L61
L77:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v439 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L78:
	;
	v332 = F_palloc0(m, v77)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+168)) = v332
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v39)+160))
	v337 = F_CopyGetAttnums(m, v72, v335, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v337 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v341 <= int32(0) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v346 = int32(0)
	goto L83
L83:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v346<<(uint(int32(2))%32))))
	v369 = v367 - int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v372 = int32(0)
	if v371 == v372 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L77
L85:
	;
	if v410 == int32(0) {
		goto L15
	} else {
		goto L98
	}
L86:
	;
	v410 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v378 <= int32(0) {
		v404 = v372
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v410 = v404
	goto L85
L90:
	;
	v381 = int32(0)
	if v381 < v378 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v384 = v378
	goto L93
L92:
	;
	v384 = v381
	goto L93
L93:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
	v387 = int32(0)
	goto L94
L94:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v385+v387<<(uint(int32(2))%32))))
	v396 = base.B2i32(v395 == v367)
	if v395 == v367 {
		v404 = v396
		goto L89
	} else {
		goto L96
	}
L95:
	;
	v404 = v396
	goto L89
L96:
	;
	v398 = v387 + int32(1)
	if v398 != v384 {
		v387 = v398
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v39)+168))
	v415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v413+v369))) = uint8(v415)
	v418 = v346 + v415
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v418 < v419 {
		v346 = v418
		goto L83
	} else {
		goto L99
	}
L99:
	;
	goto L84
L100:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[6]))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	goto L103
L101:
	;
	v445 = v439
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v445
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	goto L107
L103:
	;
	v445 = v444
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+172)) = l2
	v474 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v474
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+48))
	v481 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+184)) = v481
	*(*int64)(unsafe.Add(mBase, uint32(v39)+192)) = v481
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+200)) = uint8(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+176)) = v480 + int32(4)
	v491 = F_palloc(m, int32(_a_F_BeginCopyFrom_7))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L115
	}
L105:
	;
	v459 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+24)) = uint8(v459)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	goto L112
L106:
	;
	v457 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+24)) = uint8(v457)
	goto L104
L107:
	;
	if v449 == v445 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v451 == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	goto L110
L110:
	;
	if v456 != 0 {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	goto L106
L112:
	;
	v465 = F_FindDefaultConversionProc(m, v461, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = v465
	if v465 == int32(0) {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	goto L104
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+328)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+324)) = v491
	v496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+336)) = uint8(v496)
	F_initStringInfo(m, v39+int32(264))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if l0 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+248)) = v502
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+252)) = v504
	goto L119
L118:
	;
	goto L119
L119:
	;
	v508 = F_palloc(m, v77*int32(28))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v511 = v76 << (uint(int32(16)) % 32) >> (uint(int32(14)) % 32)
	v512 = F_palloc(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v514 = F_palloc(m, v511)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v516 = F_palloc(m, v511)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v77 <= int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v824 = int32(0)
	v833 = v474
	goto L13
L125:
	;
	goto L126
L126:
	;
	v524 = int32(0)
	v529 = int32(1)
	v533 = v474
	goto L127
L127:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v547 = v72 + v541<<(uint(int32(4))%32) + v529*int32(100)
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+11)))
	if v548 != 0 {
		v696 = v524
		v699 = v533
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v824 = v696
	v833 = v699
	goto L13
L129:
	;
	if v529 != v77 {
		v524 = v696
		v529 = v529 + int32(1)
		v533 = v699
		goto L127
	} else {
		goto L182
	}
L130:
	;
	v550 = v547 - int32(80)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)+68))
	v553 = v529 - int32(1)
	v558 = v553 << (uint(int32(2)) % 32)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	m.T0[v561].(func(*base.Module, int32, int32, int32, int32))(m, v39, v551, v508+v553*int32(28), v512+v558)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v564 = v516 + v558
	v565 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v565
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	if v567 == v565 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v571 = int32(0)
	if v570 == v571 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	goto L134
L134:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+90)))
	if v610 != 0 {
		v696 = v524
		v699 = v533
		goto L129
	} else {
		goto L149
	}
L135:
	;
	if v609 != 0 {
		v696 = v524
		v699 = v533
		goto L129
	} else {
		goto L148
	}
L136:
	;
	v609 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v577 <= int32(0) {
		v603 = v571
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v609 = v603
	goto L135
L140:
	;
	v580 = int32(0)
	if v580 < v577 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v583 = v577
	goto L143
L142:
	;
	v583 = v580
	goto L143
L143:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v570)+12))
	v586 = int32(0)
	goto L144
L144:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v584+v586<<(uint(int32(2))%32))))
	v595 = base.B2i32(v594 == v529)
	if v594 == v529 {
		v603 = v595
		goto L139
	} else {
		goto L146
	}
L145:
	;
	v603 = v595
	goto L139
L146:
	;
	v597 = v586 + int32(1)
	if v597 != v583 {
		v586 = v597
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	goto L134
L149:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v612 = F_build_column_default(m, v611, v529)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	if v612 == int32(0) {
		v696 = v524
		v699 = v533
		goto L129
	} else {
		goto L151
	}
L151:
	;
	v616 = F_expression_planner(m, v612)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v619 = F_ExecInitExpr(m, v616, int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v619
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v623 = int32(0)
	if v622 == v623 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v661 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L155:
	;
	v661 = int32(0)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	if v629 <= int32(0) {
		v655 = v623
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v661 = v655
	goto L154
L159:
	;
	v632 = int32(0)
	if v632 < v629 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v635 = v629
	goto L162
L161:
	;
	v635 = v632
	goto L162
L162:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	v638 = int32(0)
	goto L163
L163:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v636+v638<<(uint(int32(2))%32))))
	v647 = base.B2i32(v646 == v529)
	if v646 == v529 {
		v655 = v647
		goto L158
	} else {
		goto L165
	}
L164:
	;
	v655 = v647
	goto L158
L165:
	;
	v649 = v638 + int32(1)
	if v649 != v635 {
		v638 = v649
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514+base.I32_extend16_s(v524)<<(uint(int32(2))%32)))) = v553
	v671 = v524 + int32(1)
	goto L169
L168:
	;
	v671 = v524
	goto L169
L169:
	;
	if v533&int32(1) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v696 = v671
	v699 = int32(1)
	goto L129
L171:
	;
	goto L172
L172:
	;
	v675 = int32(0)
	if v616 == v675 {
		v695 = v675
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v696 = v671
	v699 = v695
	goto L129
L174:
	;
	v681 = F_check_functions_in_node(m, v616, int32(862), int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	if v681 != 0 {
		v695 = int32(1)
		goto L173
	} else {
		goto L176
	}
L176:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	if v683 == int32(67) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v687 = int32(0)
	v689 = F_query_tree_walker_impl(m, v616, int32(863), v687, v687)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	goto L179
L179:
	;
	v693 = F_expression_tree_walker_impl(m, v616, int32(863), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L181
	}
L180:
	;
	v695 = v689
	goto L173
L181:
	;
	v695 = v693
	goto L173
L182:
	;
	goto L128
L183:
	;
	F_errcode(m, int32(_a_F_BeginCopyFrom_8))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = int32(_a_F_BeginCopyFrom_9)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v72 + v125<<(uint(int32(4))%32) + v124*int32(100) + int32(24)
	F_errmsg(m, int32(_a_F_BeginCopyFrom_10), v21+int32(112))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1612), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	F_errcode(m, int32(_a_F_BeginCopyFrom_8))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = int32(_a_F_BeginCopyFrom_13)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v72 + v260<<(uint(int32(4))%32) + v259*int32(100) + int32(24)
	F_errmsg(m, int32(_a_F_BeginCopyFrom_10), v21+int32(96))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1655), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(_a_F_BeginCopyFrom_8))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v72 + v370<<(uint(int32(4))%32) + v369*int32(100) + int32(24)
	F_errmsg_internal(m, int32(_a_F_BeginCopyFrom_14), v21+int32(80))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1679), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if base.Ui32(v793) <= base.Ui32(int32(41)) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v802 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[7]))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v802)+4))
	goto L201
L198:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v793<<(uint(int32(3))%32))+uint32(_c_F_BeginCopyFrom[8])))
	v800 = v798
	goto L200
L199:
	;
	v800 = int32(_a_F_BeginCopyFrom_15)
	goto L200
L200:
	;
	goto L197
L201:
	;
	if base.Ui32(v803) <= base.Ui32(int32(41)) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v810
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v800
	F_errmsg(m, int32(_a_F_BeginCopyFrom_16), v21-int32(-64))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L206
	}
L203:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v803<<(uint(int32(3))%32))+uint32(_c_F_BeginCopyFrom[8])))
	v810 = v808
	goto L205
L204:
	;
	v810 = int32(_a_F_BeginCopyFrom_15)
	goto L205
L205:
	;
	goto L202
L206:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1709), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+240)) = v842
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v846 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+56))
	v849 = v847
	goto L211
L210:
	;
	v849 = int32(0)
	goto L211
L211:
	;
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[9]))
	if v852 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v889 = v833 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+244)) = uint8(v889)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+236)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v39)+232)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v39)+216)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v39)+212)) = v508
	*(*int64)(unsafe.Add(mBase, uint32(v39)+344)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+208)) = uint16(v824)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+44)) = uint8(v5)
	if l5 != 0 {
		goto L219
	} else {
		goto L220
	}
L213:
	;
	goto L212
L214:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginCopyFrom[10])))
	if v856&int32(1) == int32(0) {
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v861 = int32(_a_F_BeginCopyFrom_17)
	v863 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	v864 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v863 + v864
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v852)))
	*(*int32)(unsafe.Add(mBase, uint32(v852))) = v867 + v864
	*(*int32)(unsafe.Add(mBase, uint32(v852)+220)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v852)+224)) = v849
	base.MemoryFill(m, v852+int32(232), int32(0), int32(160))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v852)))
	*(*int32)(unsafe.Add(mBase, uint32(v852))) = v878 + v864
	v884 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v884 - v864
	goto L213
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L298
	}
L217:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L1
	} else {
		goto L294
	}
L218:
	;
	goto L278
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(4)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(2)
	goto L218
L220:
	;
	goto L221
L221:
	;
	if l3 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(3)
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[12]))
	if v909 == int32(2) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v1023 = F_pstrdup(m, l3)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L244
	}
L225:
	;
	v912 = m.G0
	v914 = v912 - int32(16)
	m.G0 = v914
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v916 != 0 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	goto L227
L227:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v1021
	goto L218
L228:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+4))
	v919 = v917
	goto L230
L229:
	;
	v919 = int32(0)
	goto L230
L230:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+60)))
	F_pq_beginmessage(m, v914, int32(71))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_enlargeStringInfo(m, v914, int32(1))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	*(*uint8)(unsafe.Add(mBase, uint32(v927+v928))) = uint8(v920)
	*(*int32)(unsafe.Add(mBase, uint32(v914)+4)) = v927 + int32(1)
	F_enlargeStringInfo(m, v914, int32(2))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	v940 = int32(8)
	v946 = v919<<(uint(v940)%32) | int32(base.Ui32(v919&int32(_a_F_BeginCopyFrom_18))>>(uint(v940)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v937+v938))) = uint16(v946)
	*(*int32)(unsafe.Add(mBase, uint32(v914)+4)) = v937 + int32(2)
	if int32(0) < v919 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v954 = v920 << (uint(int32(8)) % 32)
	v957 = int32(0)
	goto L237
L235:
	;
	goto L236
L236:
	;
	F_pq_endmessage(m, v914)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L241
	}
L237:
	;
	F_enlargeStringInfo(m, v914, int32(2))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L239
	}
L238:
	;
	goto L236
L239:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	*(*uint16)(unsafe.Add(mBase, uint32(v977+v978))) = uint16(v954)
	*(*int32)(unsafe.Add(mBase, uint32(v914)+4)) = v977 + int32(2)
	v985 = v957 + int32(1)
	if v985 != v919 {
		v957 = v985
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(1)
	v1009 = F_makeStringInfo(m)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v1009
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[14]))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+4))
	v1015 = m.T0[v1014].(func(*base.Module) int32)(m)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	m.G0 = v914 + int32(16)
	goto L218
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v1023
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+44)))
	if v1026 == int32(1) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(2)
	v1032 = F_OpenPipeStream(m, v1023, int32(_a_F_BeginCopyFrom_19))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+232)) = int64(1)
	v1054 = F_AllocateFile(m, v1023, int32(_a_F_BeginCopyFrom_19))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L254
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v1032
	if v1032 != 0 {
		goto L218
	} else {
		goto L249
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1041
	F_errmsg(m, int32(_a_F_BeginCopyFrom_20), v21)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1864), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v1054
	if v1054 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[15]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+60))
	if v1090 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L258:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1067
	F_errmsg(m, int32(_a_F_BeginCopyFrom_21), v21+int32(16))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	if base.B2i32(v1060 != int32(44))&base.B2i32(v1060 != int32(2)) == int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	F_errhint(m, int32(_a_F_BeginCopyFrom_22), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1883), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L265
	}
L264:
	;
	goto L263
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	if v1097 < int32(0) {
		goto L271
	} else {
		goto L272
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[15])) = int32(8)
	v1097 = int32(-1)
	goto L269
L268:
	;
	v1097 = v1090
	goto L269
L269:
	;
	goto L266
L270:
	;
	if v1107 != 0 {
		goto L217
	} else {
		goto L274
	}
L271:
	;
	v1103 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v1107 = v1103
	goto L270
L272:
	;
	goto L273
L273:
	;
	v1106 = F___fstatat(m, v1097, int32(_a_F_BeginCopyFrom_15), v21+int32(128), int32(_a_F_BeginCopyFrom_23))
	mBase = m.M
	v1107 = v1106
	goto L270
L274:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v21)+132))
	if v1108&int32(_a_F_BeginCopyFrom_24) == int32(_a_F_BeginCopyFrom_25) {
		goto L216
	} else {
		goto L275
	}
L275:
	;
	v1113 = *(*int64)(unsafe.Add(mBase, uint32(v21)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+240)) = v1113
	goto L218
L276:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	m.T0[v1312].(func(*base.Module, int32, int32))(m, v39, v72)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L293
	}
L277:
	;
	goto L276
L278:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[9]))
	if v1147 == int32(0) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginCopyFrom[10])))
	if v1151&int32(1) == int32(0) {
		goto L277
	} else {
		goto L280
	}
L280:
	;
	v1156 = int32(_a_F_BeginCopyFrom_17)
	v1158 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	v1159 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v1158 + v1159
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1147)))
	*(*int32)(unsafe.Add(mBase, uint32(v1147))) = v1162 + v1159
	goto L282
L281:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1147)))
	v1293 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1147))) = v1292 + v1293
	v1296 = int32(_a_F_BeginCopyFrom_17)
	v1298 = *(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[11])) = v1298 - v1293
	goto L277
L282:
	;
	goto L284
L284:
	;
	goto L285
L285:
	;
	v1257 = int32(0)
	v1260 = int32(0)
	goto L290
L290:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(256)+v1260<<(uint(int32(2))%32))))
	v1270 = int32(3)
	v1276 = *(*int64)(unsafe.Add(mBase, uint32(v21+int32(224)+v1260<<(uint(v1270)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1147+int32(232)+v1269<<(uint(v1270)%32)))) = v1276
	v1278 = int32(1)
	v1281 = v1257 + v1278
	if v1281 != int32(3) {
		v1257 = v1281
		v1260 = v1260 + v1278
		goto L290
	} else {
		goto L292
	}
L291:
	;
	goto L281
L292:
	;
	goto L291
L293:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BeginCopyFrom[5])) = v53
	m.G0 = v21 + int32(272)
	return v39
L294:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1327
	F_errmsg(m, int32(_a_F_BeginCopyFrom_26), v21+int32(48))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1890), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v1346
	F_errmsg(m, int32(_a_F_BeginCopyFrom_27), v21+int32(32))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_BeginCopyFrom_11), int32(1895), int32(_a_F_BeginCopyFrom_12))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromBinaryOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
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
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+4)))
	v30 = v28
	goto L3
L2:
	;
	v30 = int32(0)
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v34 + int64(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v38-v39 <= int32(1) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L21
	} else {
		goto L132
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L21
	} else {
		goto L128
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L21
	} else {
		goto L124
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L21
	} else {
		goto L120
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L21
	} else {
		goto L116
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v108 + int32(1)
	goto L4
L10:
	;
	m.G0 = v24 + int32(16)
	return v507
L11:
	;
	v128 = int32(_a_F_CopyFromBinaryOneRow_0)
	if v111&v128 == v128 {
		goto L32
	} else {
		goto L33
	}
L12:
	;
	v46 = v39
	v49 = int32(0)
	v51 = v38
	v52 = v24 + int32(10)
	goto L16
L13:
	;
	goto L14
L14:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+v39))))
	v105 = v39 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v105
	v108 = v105
	v111 = v103
	v112 = v101
	v113 = v38
	goto L11
L15:
	;
	v507 = int32(0)
	goto L10
L16:
	;
	if v46 == v51 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v87 != int32(2) {
		goto L15
	} else {
		goto L31
	}
L18:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = v46
	v75 = v51
	goto L20
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v78 = int32(2) - v49
	v79 = v75 - v74
	if v78 < v79 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return int32(0)
L22:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v71 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v74 = v73
	v75 = v72
	goto L20
L24:
	;
	v81 = v78
	goto L26
L25:
	;
	v81 = v79
	goto L26
L26:
	;
	if v81 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v52, v74+v76, v81)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v84 = v74 + v81
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v84
	v87 = v49 + v81
	if v87 < int32(2) {
		v46 = v84
		v49 = v87
		v51 = v75
		v52 = v52 + v81
		goto L16
	} else {
		goto L30
	}
L30:
	;
	goto L17
L31:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
	v108 = v84
	v111 = v92
	v112 = v76
	v113 = v75
	goto L11
L32:
	;
	v132 = int32(0)
	if v132 < v113-v108 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v183 = int32(8)
	v190 = base.I32_extend16_s(v111<<(uint(v183)%32) | int32(base.Ui32(v111&int32(_a_F_CopyFromBinaryOneRow_1))>>(uint(v183)%32)))
	if v30 != v190 {
		goto L8
	} else {
		goto L51
	}
L35:
	;
	v139 = v108
	v142 = v132
	v143 = v112
	v144 = v113
	v145 = v24 + int32(9)
	goto L36
L36:
	;
	if v139 == v144 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v507 = int32(0)
	goto L10
L38:
	;
	goto L37
L39:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	v166 = v139
	v167 = v143
	v168 = v144
	goto L41
L41:
	;
	v170 = int32(1) - v142
	v171 = v168 - v166
	if v170 < v171 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v162 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v166 = v165
	v167 = v163
	v168 = v164
	goto L41
L44:
	;
	v173 = v170
	goto L46
L45:
	;
	v173 = v171
	goto L46
L46:
	;
	if v173 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	base.MemoryCopy(m, v145, v166+v167, v173)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v176 = v166 + v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v176
	v179 = v173 + v142
	if v179 <= int32(0) {
		v139 = v176
		v142 = v179
		v143 = v167
		v144 = v168
		v145 = v173 + v145
		goto L36
	} else {
		goto L50
	}
L50:
	;
	goto L4
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v192 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v507 = int32(1)
	goto L10
L53:
	;
	goto L54
L54:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v197 <= int32(0) {
		v507 = int32(1)
		goto L10
	} else {
		goto L55
	}
L55:
	;
	v201 = l0 + int32(264)
	v214 = int32(0)
	goto L56
L56:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v225 = int32(4)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	v229 = int32(2)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228+v214<<(uint(v229)%32))))
	v237 = v33 + v224<<(uint(v225)%32) + v232*int32(100) - int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v237 + v225
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237)+76))
	v243 = v232 - int32(1)
	v245 = v243 << (uint(v229) % 32)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v31+v245)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v250-v251 < v225 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v507 = v501
	goto L10
L58:
	;
	v351 = l3 + v243
	v354 = v32 + v243*int32(28)
	if v334 == int32(-1) {
		goto L83
	} else {
		goto L84
	}
L59:
	;
	v257 = v251
	v260 = v24 + int32(12)
	v262 = int32(0)
	v263 = v250
	goto L63
L60:
	;
	goto L61
L61:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v324+v251)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v251 + int32(4)
	v334 = v326
	goto L58
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L21
	} else {
		goto L78
	}
L63:
	;
	if v257 == v263 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v296 != int32(4) {
		goto L62
	} else {
		goto L77
	}
L65:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L21
	} else {
		goto L68
	}
L66:
	;
	v283 = v257
	v284 = v263
	goto L67
L67:
	;
	v286 = int32(4) - v262
	v287 = v284 - v283
	if v286 < v287 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v280 != 0 {
		goto L62
	} else {
		goto L69
	}
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v283 = v282
	v284 = v281
	goto L67
L70:
	;
	v289 = v286
	goto L72
L71:
	;
	v289 = v287
	goto L72
L72:
	;
	if v289 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	base.MemoryCopy(m, v260, v290+v283, v289)
	goto L75
L74:
	;
	goto L75
L75:
	;
	v293 = v283 + v289
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v293
	v296 = v289 + v262
	if v296 < int32(4) {
		v257 = v293
		v260 = v260 + v289
		v262 = v296
		v263 = v284
		goto L63
	} else {
		goto L76
	}
L76:
	;
	goto L64
L77:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v334 = v301
	goto L58
L78:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_2), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2023), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L21
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2+v245))) = v477
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = int32(0)
	v501 = int32(1)
	v503 = v214 + v501
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v503 < v504 {
		v214 = v503
		goto L56
	} else {
		goto L115
	}
L83:
	;
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v357)
	v360 = F_ReceiveFunctionCall(m, v354, int32(0), v247, v241)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L21
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v364 = int32(16711935)
	v370 = base.I32_rotr(v334, int32(24))&v364 | base.I32_rotr(v334&v364, int32(8))
	if v370 < int32(0) {
		goto L7
	} else {
		goto L87
	}
L86:
	;
	v477 = v360
	goto L82
L87:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v374 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v373))) = uint8(v374)
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v374
	goto L88
L88:
	;
	F_enlargeStringInfo(m, v201, v370)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v370 <= v384-v385 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v370
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v467 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465+v370))) = uint8(v467)
	v469 = F_ReceiveFunctionCall(m, v354, v201, v247, v241)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L21
	} else {
		goto L113
	}
L91:
	;
	if v370 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v398 = v385
	v400 = int32(0)
	v401 = v382
	goto L97
L94:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	base.MemoryCopy(m, v382, v388+v385, v370)
	goto L96
L95:
	;
	goto L96
L96:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v391 + v370
	goto L90
L97:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v398 == v415 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	if v370 != v439 {
		goto L6
	} else {
		goto L112
	}
L99:
	;
	goto L98
L100:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L21
	} else {
		goto L103
	}
L101:
	;
	v422 = v415
	v423 = v398
	goto L102
L102:
	;
	v424 = v370 - v400
	v425 = v422 - v423
	if v424 < v425 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v419 != 0 {
		v439 = v400
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v422 = v421
	v423 = v420
	goto L102
L105:
	;
	v427 = v424
	goto L107
L106:
	;
	v427 = v425
	goto L107
L107:
	;
	if v427 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	base.MemoryCopy(m, v401, v428+v423, v427)
	goto L110
L109:
	;
	goto L110
L110:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v432 = v431 + v427
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v432
	v435 = v427 + v400
	if v435 < v370 {
		v398 = v432
		v400 = v435
		v401 = v427 + v401
		goto L97
	} else {
		goto L111
	}
L111:
	;
	v439 = v435
	goto L99
L112:
	;
	goto L90
L113:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+268))
	if v471 != v472 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	v474 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v474)
	v477 = v469
	goto L82
L115:
	;
	goto L57
L116:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v190
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_5), v24)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(1130), int32(_a_F_CopyFromBinaryOneRow_6))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_7), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2032), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_2), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L21
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2042), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L21
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L21
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_8), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L21
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(2055), int32(_a_F_CopyFromBinaryOneRow_4))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L21
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_CopyFromBinaryOneRow_9), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_CopyFromBinaryOneRow_3), int32(1122), int32(_a_F_CopyFromBinaryOneRow_6))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromErrorCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v145 int32
	_ = v145
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v11 == int32(1) {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
			F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_0), v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				m.G0 = v9 + int32(176)
				return
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
		if v23 == int32(1) {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
				if v22 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v31
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v30
					F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_1), v9+int32(32))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(176)
						return
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v30
					F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_2), v9+int32(16))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v9 + int32(176)
						return
					}
				}
			}
		} else {
			if v22 != 0 {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
				if v47 != 0 {
					v48 = F_strlen(m, v47)
					mBase = m.M
					if v48 <= int32(100) {
						v51 = F_pstrdup(m, v47)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v64 = v51
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
								v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v71
								*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v70
								*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v69
								F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_3), v9+int32(144))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									F_pfree(m, v64)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										m.G0 = v9 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v54 = F_pg_mbcliplen(m, v47, v48, int32(100))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v58 = F_palloc(m, v54+int32(4))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								if v54 != 0 {
									base.MemoryCopy(m, v58, v47, v54)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v58+v54))) = int32(_a_F_CopyFromErrorCallback_4)
								v64 = v58
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
									v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v64
									*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v71
									*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v70
									*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v69
									F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_3), v9+int32(144))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_pfree(m, v64)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											m.G0 = v9 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
						v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v88
						*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v87
						*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v86
						F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_5), v9+int32(112))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							m.G0 = v9 + int32(176)
							return
						}
					}
				}
			} else {
				v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)))
				if v97 == int32(1) {
					v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
					v101 = F_strlen(m, v100)
					mBase = m.M
					if v101 <= int32(100) {
						v104 = F_pstrdup(m, v100)
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							v117 = v104
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
								v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v117
								*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v123
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v122
								F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_6), v9-int32(-64))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return
								} else {
									F_pfree(m, v117)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										m.G0 = v9 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v107 = F_pg_mbcliplen(m, v100, v101, int32(100))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							v111 = F_palloc(m, v107+int32(4))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								if v107 != 0 {
									base.MemoryCopy(m, v111, v100, v107)
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v111+v107))) = int32(_a_F_CopyFromErrorCallback_4)
								v117 = v111
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
									v123 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v117
									*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v123
									*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v122
									F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_6), v9-int32(-64))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return
									} else {
										F_pfree(m, v117)
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return
										} else {
											m.G0 = v9 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
						v138 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v138
						*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v137
						F_errcontext_msg(m, int32(_a_F_CopyFromErrorCallback_2), v9+int32(96))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return
						} else {
							m.G0 = v9 + int32(176)
							return
						}
					}
				}
			}
		}
	}
}
func F_CopyFromTextLikeInFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_getTypeInputInfo(m, l1, v7+int32(12), l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		F_fmgr_info(m, v13, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_RemoveFromWaitQueue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v20 - v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v24 - v21
	v30 = v10 + v12<<(uint(int32(2))%32)
	v32 = v30 + int32(44)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = v33 - v21
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	if v35 == v37 {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39 & base.I32_rotl(int32(-2), v12)
	} else {
	}
	v44 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(l0)+92)) = int64(0)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(v44)%32))+uint32(_c_F_RemoveFromWaitQueue[0])))
	F_CleanUpLock(m, v10, v9, v50, l1, int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		return
	} else {
		return
	}
}
