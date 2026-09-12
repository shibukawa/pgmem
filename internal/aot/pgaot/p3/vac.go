package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_update_datfrozenxid(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
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
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
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
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v448 int32
	_ = v448
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
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
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int64
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var __phi569 int32
	_ = __phi569
	var v571 int32
	_ = v571
	var __phi571 int32
	_ = __phi571
	var v575 int32
	_ = v575
	var __phi575 int32
	_ = __phi575
	var v578 int32
	_ = v578
	var __phi578 int32
	_ = __phi578
	var v584 int64
	_ = v584
	var __phi584 int64
	_ = __phi584
	var v585 int64
	_ = v585
	var __phi585 int64
	_ = __phi585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int64
	_ = v602
	var v603 int64
	_ = v603
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int64
	_ = v617
	var v618 int64
	_ = v618
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v677 int64
	_ = v677
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int64
	_ = v697
	var v698 int64
	_ = v698
	var v704 int32
	_ = v704
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int64
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int64
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int64
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int64
	_ = v823
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v835 int64
	_ = v835
	var v836 int32
	_ = v836
	var v838 int64
	_ = v838
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int64
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int64
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1067 int64
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1106 int64
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1131 int64
	_ = v1131
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	v1 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(16908288)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = int64(0)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
	v38 = F_LockAcquire(m, v26, int32(7), v1, v1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	m.G0 = v26 + int32(16)
	v44 = F_GetOldestNonRemovableTransactionId(m, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v48 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = base.I32_wrap_i64(v48)
	v51 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v55 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v62)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	v57 = int32(0)
	v62 = F_systable_beginscan(m, v55, v57, v57, v57, v57, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v64 = F_systable_getnext(m, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v64 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v68 = v64
	v70 = v44
	v72 = v46
	goto L14
L12:
	;
	v144 = v44
	v146 = v46
	goto L13
L13:
	;
	v164 = v144
	v166 = v146
	v179 = int32(1)
	goto L7
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
	v87 = v85 + v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+136))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+140))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v90 == int32(114) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v144 = v136
	v146 = v137
	goto L13
L16:
	;
	v138 = F_systable_getnext(m, v62)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L43
	}
L17:
	;
	if v88 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v93 == int32(109) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v96 != int32(116) {
		v136 = v70
		v137 = v72
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v99 = int32(0)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v88))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == v99 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v125 = v70
	goto L23
L23:
	;
	if v89 == int32(0) {
		v136 = v125
		v137 = v72
		goto L16
	} else {
		goto L36
	}
L24:
	;
	if v111 != 0 {
		v164 = v70
		v166 = v72
		v179 = v99
		goto L7
	} else {
		goto L28
	}
L25:
	;
	v111 = base.B2i32(base.Ui32(v50) < base.Ui32(v88))
	goto L24
L26:
	;
	goto L27
L27:
	;
	v111 = int32(base.Ui32(v50-v88) >> (uint(int32(31)) % 32))
	goto L24
L28:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v70))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v123 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v123 = base.B2i32(base.Ui32(v88) < base.Ui32(v70))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v123 = int32(base.Ui32(v88-v70) >> (uint(int32(31)) % 32))
	goto L29
L33:
	;
	v124 = v88
	goto L35
L34:
	;
	v124 = v70
	goto L35
L35:
	;
	v125 = v124
	goto L23
L36:
	;
	goto L37
L37:
	;
	if int32(base.Ui32(v51-v89)>>(uint(int32(31))%32)) != 0 {
		v164 = v125
		v166 = v72
		v179 = int32(0)
		goto L7
	} else {
		goto L38
	}
L38:
	;
	goto L39
L39:
	;
	if int32(base.Ui32(v89-v72)>>(uint(int32(31))%32)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = v89
	goto L42
L41:
	;
	v135 = v72
	goto L42
L42:
	;
	v136 = v125
	v137 = v135
	goto L16
L43:
	;
	if v138 != 0 {
		v68 = v138
		v70 = v136
		v72 = v137
		goto L14
	} else {
		goto L44
	}
L44:
	;
	goto L15
L45:
	;
	F_sequence_close(m, v55, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v179 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	m.G0 = v22 + int32(96)
	return
L48:
	;
	v187 = int32(0)
	v190 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	F_ScanKeyInit(m, v22+int32(32), int32(1), int32(3), int32(184), v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_systable_inplace_update_begin(m, v190, int32(2672), v22+int32(32), v22+int32(92), v22+int32(28))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	if v210 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_pfree(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L84
	}
L53:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_systable_inplace_update_finish(m, v278, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L83
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+88)) = v166
	v277 = v166
	goto L53
L55:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+22)))
	v213 = v211 + v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+84))
	if v214 == v164 {
		v244 = v164
		v245 = v187
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L80
	}
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v213)+88))
	if v166 != v246 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v164))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v214)) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+84)) = v164
	v244 = v164
	v245 = int32(1)
	goto L58
L61:
	;
	if v227 != 0 {
		goto L60
	} else {
		goto L65
	}
L62:
	;
	v227 = base.B2i32(base.Ui32(v214) < base.Ui32(v164))
	goto L61
L63:
	;
	goto L64
L64:
	;
	v227 = int32(base.Ui32(v214-v164) >> (uint(int32(31)) % 32))
	goto L61
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v213)+84))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v228))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v240 != 0 {
		goto L60
	} else {
		goto L70
	}
L67:
	;
	v240 = base.B2i32(base.Ui32(v50) < base.Ui32(v228))
	goto L66
L68:
	;
	goto L69
L69:
	;
	v240 = int32(base.Ui32(v50-v228) >> (uint(int32(31)) % 32))
	goto L66
L70:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v213)+84))
	v244 = v241
	v245 = v187
	goto L58
L71:
	;
	goto L74
L72:
	;
	v256 = v166
	goto L73
L73:
	;
	if v245 != 0 {
		v277 = v256
		goto L53
	} else {
		goto L78
	}
L74:
	;
	if int32(base.Ui32(v246-v166)>>(uint(int32(31))%32)) != 0 {
		goto L54
	} else {
		goto L75
	}
L75:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v213)+88))
	goto L76
L76:
	;
	if int32(base.Ui32(v51-v251)>>(uint(int32(31))%32)) != 0 {
		goto L54
	} else {
		goto L77
	}
L77:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v213)+88))
	v256 = v255
	goto L73
L78:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_systable_inplace_update_cancel(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v283 = v256
	v284 = int32(0)
	goto L52
L80:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v266
	F_errmsg_internal(m, int32(54332), v22)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(522856), int32(1776), int32(454292))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v283 = v277
	v284 = int32(1)
	goto L52
L84:
	;
	F_sequence_close(m, v190, int32(3))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v284 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v293 = int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v299 = F_LWLockAcquire(m, v295+int32(384), v293)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v346 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L101
	}
L89:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+36))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v302)+20))
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v302)+8))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302)+16))
	v308 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v308+int32(384))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if base.Ui32(v306) < base.Ui32(int32(3)) {
		v338 = v293
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v338 == int32(0) {
		goto L47
	} else {
		goto L100
	}
L92:
	;
	if v304 == int32(0) {
		v338 = v293
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v317 = base.I32_wrap_i64(v305)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v304))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v317)) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v329 != 0 {
		v338 = v293
		goto L91
	} else {
		goto L98
	}
L95:
	;
	v329 = base.B2i32(base.Ui32(v304) <= base.Ui32(v317))
	goto L94
L96:
	;
	goto L97
L97:
	;
	v329 = base.B2i32(int32(0) <= v317-v304)
	goto L94
L98:
	;
	v331 = int32(0)
	v334 = F_SearchSysCacheExists(m, int32(21), v303, v331, v331, v331)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v338 = v334 ^ int32(1)
	goto L91
L100:
	;
	goto L88
L101:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v353 = F_LWLockAcquire(m, v349+int32(5888), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v359 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)+188))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+12))
	m.T0[v504].(func(*base.Module, int32))(m, v363)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L152
	}
L104:
	;
	v361 = int32(0)
	v363 = F_table_beginscan_catalog(m, v359, v361, v361)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v365 = F_heap_getnext(m, v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v365 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v487 = v244
	v489 = v283
	v491 = v1
	v493 = v1
	v495 = v356
	v496 = v356
	goto L103
L108:
	;
	goto L109
L109:
	;
	v369 = base.I32_wrap_i64(v346)
	v372 = v365
	v374 = v244
	v376 = v283
	v378 = v1
	v380 = v1
	v382 = v356
	v383 = v356
	goto L110
L110:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+22)))
	v391 = v389 + v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+84))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+88))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v391)+80))
	goto L113
L111:
	;
	v487 = v475
	v489 = v476
	v491 = v477
	v493 = v478
	v495 = v479
	v496 = v480
	goto L103
L112:
	;
	v481 = F_heap_getnext(m, v363)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L150
	}
L113:
	;
	if v394 == int32(-2) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v399 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v392))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	if v399 == int32(0) {
		v475 = v374
		v476 = v376
		v477 = v378
		v478 = v380
		v479 = v382
		v480 = v383
		goto L112
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v391 + int32(4)
	F_errmsg_internal(m, int32(454316), v22+int32(16))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(522856), int32(1905), int32(344684))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v475 = v374
	v476 = v376
	v477 = v378
	v478 = v380
	v479 = v382
	v480 = v383
	goto L112
L121:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v392))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v369)) == int32(0) {
		goto L133
	} else {
		goto L134
	}
L122:
	;
	if v427 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v427 = base.B2i32(base.Ui32(v50) < base.Ui32(v392))
	goto L122
L124:
	;
	goto L125
L125:
	;
	v427 = int32(base.Ui32(v50-v392) >> (uint(int32(31)) % 32))
	goto L122
L126:
	;
	goto L129
L127:
	;
	goto L128
L128:
	;
	v436 = int32(1)
	goto L121
L129:
	;
	if int32(base.Ui32(v51-v393)>>(uint(int32(31))%32)) == int32(0) {
		v436 = v378
		goto L121
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	goto L146
L132:
	;
	if v448 != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v448 = base.B2i32(base.Ui32(v369) < base.Ui32(v392))
	goto L132
L134:
	;
	goto L135
L135:
	;
	v448 = int32(base.Ui32(v369-v392) >> (uint(int32(31)) % 32))
	goto L132
L136:
	;
	v465 = v374
	v466 = int32(1)
	v467 = v383
	goto L131
L137:
	;
	goto L138
L138:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v374))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v392)) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v461 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v461 = base.B2i32(base.Ui32(v392) < base.Ui32(v374))
	goto L139
L141:
	;
	goto L142
L142:
	;
	v461 = int32(base.Ui32(v392-v374) >> (uint(int32(31)) % 32))
	goto L139
L143:
	;
	v465 = v374
	v466 = v380
	v467 = v383
	goto L131
L144:
	;
	goto L145
L145:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v465 = v392
	v466 = v380
	v467 = v464
	goto L131
L146:
	;
	if int32(base.Ui32(v393-v376)>>(uint(int32(31))%32)) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v472 = v393
	v473 = v471
	goto L149
L148:
	;
	v472 = v376
	v473 = v382
	goto L149
L149:
	;
	v475 = v465
	v476 = v472
	v477 = v436
	v478 = v466
	v479 = v473
	v480 = v467
	goto L112
L150:
	;
	if v481 != 0 {
		v372 = v481
		v374 = v475
		v376 = v476
		v378 = v477
		v380 = v478
		v382 = v479
		v383 = v480
		goto L110
	} else {
		goto L151
	}
L151:
	;
	goto L111
L152:
	;
	F_sequence_close(m, v359, int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v493 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v1229+int32(5888))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L321
	}
L155:
	;
	v512 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if v491 != 0 {
		goto L154
	} else {
		goto L163
	}
L158:
	;
	if v512 == int32(0) {
		goto L154
	} else {
		goto L159
	}
L159:
	;
	F_errmsg(m, int32(151201), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errdetail(m, int32(616532), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(522856), int32(1951), int32(344684))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	goto L154
L163:
	;
	v529 = int32(0)
	v531 = m.G0
	v533 = v531 - int32(16)
	m.G0 = v533
	v536 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v540 = F_LWLockAcquire(m, v536+int32(6016), int32(1))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v547 = F_LWLockAcquire(m, v543+int32(3456), int32(1))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _consts[460]))
	v551 = *(*int64)(unsafe.Add(mBase, uint32(v550)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v551
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v550)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v553
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v550)+8))
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v550)))
	v558 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v558+int32(3456))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	if base.B2i32(v553 == v555)&base.B2i32(v551 == v556) != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v725+int32(6016))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L204
	}
L168:
	;
	__phi569 = v529
	__phi571 = v553
	__phi575 = int32(-1)
	__phi578 = v529
	__phi584 = int64(-1)
	__phi585 = v551
	v569 = __phi569
	v571 = __phi571
	v575 = __phi575
	v578 = __phi578
	v584 = __phi584
	v585 = __phi585
	goto L169
L169:
	;
	if v584 != v585 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v640 < int32(0) {
		goto L167
	} else {
		goto L199
	}
L171:
	;
	v589 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if int32(0) <= v575 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v639 = v569
	v640 = v575
	v641 = v578
	goto L173
L173:
	;
	v642 = v639 + v571
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+8))
	if base.Ui32(v643) < base.Ui32(int32(3)) {
		v667 = v641
		goto L183
	} else {
		goto L184
	}
L174:
	;
	if v578 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v613 = v578
	v614 = v589
	goto L176
L176:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+28))
	v617 = int64(*(*uint16)(unsafe.Add(mBase, _consts[527])))
	v618 = base.I64_rem_s(v585, v617)
	v624 = F_LWLockAcquire(m, v615+base.I32_wrap_i64(v618)<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L181
	}
L177:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v589)+12))
	v594 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v592+v575))) = uint8(v594)
	v597 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	v598 = v597
	goto L179
L178:
	;
	v598 = v589
	goto L179
L179:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v598)+28))
	v602 = int64(*(*uint16)(unsafe.Add(mBase, _consts[527])))
	v603 = base.I64_rem_s(v584, v602)
	F_LWLockRelease(m, v600+base.I32_wrap_i64(v603)<<(uint(int32(7))%32))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	v613 = int32(0)
	v614 = v611
	goto L176
L181:
	;
	v629 = F_SimpleLruReadPage(m, int32(4458936), v585, int32(1), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v632 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v633+v629<<(uint(int32(2))%32))))
	v639 = v637
	v640 = v629
	v641 = v613
	goto L173
L183:
	;
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v533)+8))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
	v671 = v669 + v670
	v675 = base.B2i32(base.Ui32(v671-int32(8173)) < base.Ui32(int32(-8193)))
	v677 = v668 + base.I64_extend_i32_u(v675)
	*(*int64)(unsafe.Add(mBase, uint32(v533)+8)) = v677
	if base.Ui32(v671-int32(8173)) < base.Ui32(int32(-8193)) {
		goto L194
	} else {
		goto L195
	}
L184:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v643)) == int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v657 == int32(0) {
		v667 = v641
		goto L183
	} else {
		goto L189
	}
L186:
	;
	v657 = base.B2i32(base.Ui32(v643) < base.Ui32(v487))
	goto L185
L187:
	;
	goto L188
L188:
	;
	v657 = int32(base.Ui32(v643-v487) >> (uint(int32(31)) % 32))
	goto L185
L189:
	;
	v662 = F_TransactionIdDidCommit(m, v643)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	if v662 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v664 = int32(2)
	goto L193
L192:
	;
	v664 = int32(0)
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v642)+8)) = v664
	v667 = int32(1)
	goto L183
L194:
	;
	v680 = int32(0)
	goto L196
L195:
	;
	v680 = v671
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v680
	if v677 != v556 {
		__phi569 = v639
		__phi571 = v680
		__phi575 = v640
		__phi578 = v667
		__phi584 = v585
		__phi585 = v677
		v569 = __phi569
		v571 = __phi571
		v575 = __phi575
		v578 = __phi578
		v584 = __phi584
		v585 = __phi585
		goto L169
	} else {
		goto L197
	}
L197:
	;
	if v680 != v555 {
		__phi569 = v639
		__phi571 = v680
		__phi575 = v640
		__phi578 = v667
		__phi584 = v585
		__phi585 = v677
		v569 = __phi569
		v571 = __phi571
		v575 = __phi575
		v578 = __phi578
		v584 = __phi584
		v585 = __phi585
		goto L169
	} else {
		goto L198
	}
L198:
	;
	goto L170
L199:
	;
	v687 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	if v667 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)+12))
	v690 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v688+v640))) = uint8(v690)
	v693 = *(*int32)(unsafe.Add(mBase, _consts[526]))
	v694 = v693
	goto L202
L201:
	;
	v694 = v687
	goto L202
L202:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+28))
	v697 = int64(*(*uint16)(unsafe.Add(mBase, _consts[527])))
	v698 = base.I64_rem_s(v585, v697)
	F_LWLockRelease(m, v695+base.I32_wrap_i64(v698)<<(uint(int32(7))%32))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	goto L167
L204:
	;
	m.G0 = v533 + int32(16)
	v734 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v738 = F_LWLockAcquire(m, v734+int32(4992), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+40))
	if v742 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v763+int32(4992))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L213
	}
L207:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v742)) == int32(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	if v756 == int32(0) {
		goto L206
	} else {
		goto L212
	}
L209:
	;
	v756 = base.B2i32(base.Ui32(v742) < base.Ui32(v487))
	goto L208
L210:
	;
	goto L211
L211:
	;
	v756 = int32(base.Ui32(v742-v487) >> (uint(int32(31)) % 32))
	goto L208
L212:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _consts[90]))
	*(*int32)(unsafe.Add(mBase, uint32(v760)+40)) = v487
	goto L206
L213:
	;
	v768 = m.G0
	v770 = v768 - int32(32)
	m.G0 = v770
	*(*int64)(unsafe.Add(mBase, uint32(v770)+8)) = base.I64_extend_i32_u(int32(base.Ui32(v487) >> (uint(int32(15)) % 32)))
	v780 = F_SlruScanDirectory(m, int32(4456592), int32(288), v770+int32(8))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	if v780 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	F_AdvanceOldestClogXid(m, v487)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v806 = int32(32)
	m.G0 = v770 + v806
	v809 = m.G0
	v811 = v809 - v806
	m.G0 = v811
	v814 = base.I32_div_u_s(v487, int32(819))
	*(*int64)(unsafe.Add(mBase, uint32(v811)+8)) = base.I64_extend_i32_u(v814)
	v821 = F_SlruScanDirectory(m, int32(4456676), int32(288), v811+int32(8))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L224
	}
L218:
	;
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v770)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+28)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v770)+24)) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v770)+16)) = v784
	F_XLogBeginInsert(m)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v790 = int32(16)
	F_XLogRegisterData(m, v770+v790, v790)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v797 = F_XLogInsert(m, int32(3), int32(16))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_XLogFlush(m, v797)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v770)+8))
	F_SimpleLruTruncate(m, int32(4456592), v802)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	goto L217
L224:
	;
	if v821 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v811)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v811)+24)) = v487
	*(*int64)(unsafe.Add(mBase, uint32(v811)+16)) = v823
	F_XLogBeginInsert(m)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	m.G0 = v811 + int32(32)
	v845 = m.G0
	v847 = v845 - int32(144)
	m.G0 = v847
	v850 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v854 = F_LWLockAcquire(m, v850+int32(5248), int32(0))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L232
	}
L228:
	;
	F_XLogRegisterData(m, v811+int32(16), int32(12))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v835 = F_XLogInsert(m, int32(18), int32(16))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v838 = *(*int64)(unsafe.Add(mBase, uint32(v811)+8))
	F_SimpleLruTruncate(m, int32(4456676), v838)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	goto L227
L232:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v861 = F_LWLockAcquire(m, v857+int32(1664), int32(1))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v864)))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v864)+12))
	v869 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v869+int32(1664))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	if v489-v867 <= int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	m.G0 = v847 + int32(144)
	F_SetTransactionIdLimit(m, v487, v496)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L1
	} else {
		goto L319
	}
L236:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v878+int32(5248))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v847)+104)) = int64(-1)
	v889 = F_SlruScanDirectory(m, int32(4456772), int32(294), v847+int32(104))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	goto L235
L240:
	;
	v891 = int32(1)
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v847)+104))
	v894 = v892 << (uint(int32(11)) % 32)
	if base.Ui32(v894) <= base.Ui32(v891) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v897 = v891
	goto L243
L242:
	;
	v897 = v894
	goto L243
L243:
	;
	if v867-v897 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v902+int32(5248))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	if v867 == v866 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	goto L235
L248:
	;
	if v489 != v866 {
		goto L261
	} else {
		goto L262
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+116)) = v865
	goto L248
L250:
	;
	goto L251
L251:
	;
	v911 = F_find_multixact_start(m, v867, v847+int32(116))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	if v911 != 0 {
		goto L248
	} else {
		goto L253
	}
L253:
	;
	v915 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	if v915 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+100)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v847)+96)) = v867
	F_errmsg(m, int32(278903), v847+int32(96))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v930+int32(5248))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	F_errfinish(m, int32(518132), int32(3259), int32(119512))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	goto L235
L261:
	;
	v938 = F_find_multixact_start(m, v489, v847+int32(120))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	v964 = v865
	goto L263
L263:
	;
	if v964 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L264:
	;
	if v938 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v944 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v847)+120))
	v964 = v963
	goto L263
L268:
	;
	if v944 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+80)) = v489
	F_errmsg(m, int32(279063), v847+int32(80))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v958 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v958+int32(5248))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L274
	}
L272:
	;
	F_errfinish(m, int32(518132), int32(3277), int32(119512))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	goto L235
L275:
	;
	v969 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v988 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L285
	}
L278:
	;
	if v969 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847))) = v489
	F_errmsg(m, int32(278977), v847)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v981+int32(5248))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L1
	} else {
		goto L284
	}
L282:
	;
	F_errfinish(m, int32(518132), int32(3294), int32(119512))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	goto L235
L285:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v847)+116))
	if v988 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1041 = int32(4556756)
	v1043 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	v1044 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v1043 + v1044
	v1048 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1048)+120)) = v1049 | v1044
	*(*int32)(unsafe.Add(mBase, uint32(v847)+140)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v847)+136)) = v990
	*(*int32)(unsafe.Add(mBase, uint32(v847)+132)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v847)+128)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v847)+124)) = v495
	F_XLogBeginInsert(m)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L292
	}
L287:
	;
	v993 = int32(1636)
	v994 = base.I32_div_u_s(v990, v993)
	v995 = int32(5)
	v998 = base.I32_div_u_s(v964, v993)
	v1000 = int32(base.Ui32(v998) >> (uint(v995) % 32))
	v1038 = v1000
	v1039 = int32(base.Ui32(v994) >> (uint(v995) % 32))
	v1040 = base.I64_extend_i32_u(v1000)
	goto L286
L288:
	;
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v847)+60)) = v964
	*(*int32)(unsafe.Add(mBase, uint32(v847)+56)) = v990
	v1004 = int32(1636)
	v1005 = base.I32_div_u_s(v964, v1004)
	v1006 = int32(5)
	v1007 = int32(base.Ui32(v1005) >> (uint(v1006) % 32))
	v1008 = base.I64_extend_i32_u(v1007)
	*(*int64)(unsafe.Add(mBase, uint32(v847)+72)) = v1008
	v1013 = base.I32_div_u_s(v990, v1004)
	v1015 = int32(base.Ui32(v1013) >> (uint(v1006) % 32))
	*(*int64)(unsafe.Add(mBase, uint32(v847-int32(-64)))) = base.I64_extend_i32_u(v1015)
	*(*int32)(unsafe.Add(mBase, uint32(v847)+36)) = v489
	v1019 = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v847)+48)) = base.I64_extend_i32_u(int32(base.Ui32(v489) >> (uint(v1019) % 32)))
	*(*int32)(unsafe.Add(mBase, uint32(v847)+32)) = v867
	*(*int64)(unsafe.Add(mBase, uint32(v847)+40)) = base.I64_extend_i32_u(int32(base.Ui32(v867) >> (uint(v1019) % 32)))
	F_errmsg_internal(m, int32(702773), v847+int32(32))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(518132), int32(3307), int32(119512))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	v1038 = v1007
	v1039 = v1015
	v1040 = v1008
	goto L286
L292:
	;
	F_XLogRegisterData(m, v847+int32(124), int32(20))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v1067 = F_XLogInsert(m, int32(6), int32(48))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	F_XLogFlush(m, v1067)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v1076 = F_LWLockAcquire(m, v1072+int32(1664), int32(0))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, _consts[103]))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+16)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+12)) = v489
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v1083+int32(1664))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	if v1038 != v1039 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1106 = base.I64_extend_i32_u(v1039)
	goto L301
L299:
	;
	goto L300
L300:
	;
	v1154 = int32(1)
	if v489 == v1154 {
		goto L314
	} else {
		goto L315
	}
L301:
	;
	v1111 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L303
	}
L302:
	;
	goto L300
L303:
	;
	if v1111 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v847)+16)) = v1106
	F_errmsg_internal(m, int32(27806), v847+int32(16))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	F_SlruDeleteSegment(m, v1106)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L309
	}
L307:
	;
	F_errfinish(m, int32(518132), int32(3132), int32(279173))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	if v1106 != int64(82040) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1131 = v1106 + int64(1)
	goto L312
L311:
	;
	v1131 = int64(0)
	goto L312
L312:
	;
	if v1131 != v1040 {
		v1106 = v1131
		goto L301
	} else {
		goto L313
	}
L313:
	;
	goto L302
L314:
	;
	v1160 = int32(2097151)
	goto L316
L315:
	;
	v1160 = int32(base.Ui32(v489-v1154) >> (uint(int32(11)) % 32))
	goto L316
L316:
	;
	F_SimpleLruTruncate(m, int32(4456772), base.I64_extend_i32_u(v1160))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1165)+120)) = v1166 & int32(-2)
	v1170 = int32(4556756)
	v1172 = *(*int32)(unsafe.Add(mBase, _consts[17]))
	*(*int32)(unsafe.Add(mBase, _consts[17])) = v1172 - int32(1)
	v1177 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v1177+int32(5248))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	goto L235
L319:
	;
	F_SetMultiXactIdLimit(m, v489, v495, int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	goto L154
L321:
	;
	goto L47
}
func F_vac_update_relstats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
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
	var v49 float32
	_ = v49
	var v50 float32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v19-int32(-64), int32(1), int32(3), int32(184), v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_systable_inplace_update_begin(m, v24, int32(2662), v19-int32(-64), v19+int32(60), v19+int32(56))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	if l1 != v46 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L85
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+96)) = l1
	goto L10
L9:
	;
	goto L10
L10:
	;
	v49 = base.F32_demote_f64(l2)
	v50 = *(*float32)(unsafe.Add(mBase, uint32(v45)+100))
	if base.F32_eq(v49, v50) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = base.B2i32(l1 != v46)
	goto L13
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v45)+100)) = v49
	v55 = int32(1)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	if l3 != v56 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+104)) = l3
	v60 = int32(1)
	goto L16
L15:
	;
	v60 = v55
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+108))
	if l4 != v61 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+108)) = l4
	v65 = int32(1)
	goto L19
L18:
	;
	v65 = v60
	goto L19
L19:
	;
	if l10 != 0 {
		v90 = v65
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v45)+136))
	if l8 != 0 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	if l5 != 0 {
		v74 = v65
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+124)))
	if v75 != int32(1) {
		v82 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+116)))
	if v66&int32(1) == int32(0) {
		v74 = v65
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+116)) = uint8(v71)
	v74 = int32(1)
	goto L22
L25:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+125)))
	if v83 != int32(1) {
		v90 = v82
		goto L20
	} else {
		goto L28
	}
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v78 != 0 {
		v82 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+124)) = uint8(v79)
	v82 = int32(1)
	goto L25
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v86 != 0 {
		v90 = v82
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+125)) = uint8(v87)
	v90 = int32(1)
	goto L20
L30:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v92)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v94 = int32(0)
	if base.Ui32(l6) < base.Ui32(int32(3)) {
		v137 = v90
		v138 = v94
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v45)+140))
	if l9 != 0 {
		goto L50
	} else {
		goto L51
	}
L34:
	;
	if v91 == l6 {
		v137 = v90
		v138 = v94
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l6))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v91)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v109 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v109 = base.B2i32(base.Ui32(v91) < base.Ui32(l6))
	goto L36
L38:
	;
	goto L39
L39:
	;
	v109 = int32(base.Ui32(v91-l6) >> (uint(int32(31)) % 32))
	goto L36
L40:
	;
	v112 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+136)) = l6
	v130 = int32(1)
	v132 = v109 ^ v130
	if l8 == int32(0) {
		v137 = v130
		v138 = v132
		goto L33
	} else {
		goto L49
	}
L43:
	;
	v114 = base.I32_wrap_i64(v112)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v91))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v114)) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v126 == int32(0) {
		v137 = v90
		v138 = v94
		goto L33
	} else {
		goto L48
	}
L45:
	;
	v126 = base.B2i32(base.Ui32(v114) < base.Ui32(v91))
	goto L44
L46:
	;
	goto L47
L47:
	;
	v126 = int32(base.Ui32(v114-v91) >> (uint(int32(31)) % 32))
	goto L44
L48:
	;
	goto L42
L49:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v135)
	v137 = v130
	v138 = v132
	goto L33
L50:
	;
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v141)
	goto L52
L51:
	;
	goto L52
L52:
	;
	if l7 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	F_sequence_close(m, v24, int32(3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L70
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	F_systable_inplace_update_cancel(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L69
	}
L55:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_systable_inplace_update_finish(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L68
	}
L56:
	;
	v166 = int32(0)
	if v137 == v166 {
		goto L54
	} else {
		goto L67
	}
L57:
	;
	if l7 == v140 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v148 = int32(base.Ui32(v140-l7) >> (uint(int32(31)) % 32))
	goto L59
L59:
	;
	if v148 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v151 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+140)) = l7
	v160 = v148 ^ int32(1)
	if l9 == int32(0) {
		v170 = v160
		goto L55
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if int32(base.Ui32(v151-v140)>>(uint(int32(31))%32)) == int32(0) {
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v163)
	v170 = v160
	goto L55
L67:
	;
	v170 = v166
	goto L55
L68:
	;
	v179 = v170
	goto L53
L69:
	;
	v179 = v166
	goto L53
L70:
	;
	if v138 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v179 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	v187 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v187 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v194 + int32(4)
	F_errmsg_internal(m, int32(753654), v19+int32(32))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(522856), int32(1595), int32(134102))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L71
L78:
	;
	m.G0 = v19 + int32(112)
	return
L79:
	;
	v215 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v215 == int32(0) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v222 + int32(4)
	F_errmsg_internal(m, int32(753727), v19+int32(16))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(522856), int32(1601), int32(134102))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L78
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	F_errmsg_internal(m, int32(353030), v19)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(522856), int32(1474), int32(134102))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
