package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_CreateSharedMemoryAndSemaphores(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int64
	_ = v192
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int64
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v562 int32
	_ = v562
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v725 int32
	_ = v725
	var v728 int64
	_ = v728
	var v730 int64
	_ = v730
	var v732 int64
	_ = v732
	var v734 int64
	_ = v734
	var v736 int64
	_ = v736
	var v738 int32
	_ = v738
	var v745 int32
	_ = v745
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1095 int32
	_ = v1095
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v22 = F_CalculateShmemSize(m, v18+int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v22
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_0), v18)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v37 = m.G0
	v39 = v37 - int32(352)
	m.G0 = v39
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	v47 = F___fstatat(m, int32(-100), v42, v39+int32(192), int32(0))
	mBase = m.M
	goto L10
L7:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_1), int32(211), int32(_a_F_CreateSharedMemoryAndSemaphores_2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1])) = v738
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2])) = v738
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v738)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[3])) = v738 + v838
	goto L203
L10:
	;
	if int32(0) <= v47 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[5]))
	if base.B2i32(v53 == int32(1))&base.B2i32(v57 != int32(2)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L199
	}
L14:
	;
	if v57 == int32(2) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L195
	}
L17:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v39)+280))
	v197 = base.I32_wrap_i64(v192)
	goto L56
L18:
	;
	v66 = int32(1)
	if base.Ui32(v66) < base.Ui32(v53-v66) {
		v121 = int32(-1)
		v122 = v22
		v123 = int32(0)
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_3), int32(_a_F_CreateSharedMemoryAndSemaphores_4), int32(0), int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L55
	}
L21:
	;
	if v121 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[6]))
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v75 = v71 << (uint(int32(10)) % 32)
	goto L25
L24:
	;
	v75 = int32(_a_F_CreateSharedMemoryAndSemaphores_5)
	goto L25
L25:
	;
	if v75 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v86 = int32(_a_F_CreateSharedMemoryAndSemaphores_6) - base.I32_wrap_i64(base.I64_clz(base.I64_extend_i32_u(v75)-int64(1)))<<(uint(int32(26))%32)
	goto L28
L27:
	;
	v86 = int32(_a_F_CreateSharedMemoryAndSemaphores_6)
	goto L28
L28:
	;
	v87 = base.I32_rem_u_s(v22, v75)
	if v87 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v90 = v75 - v87
	goto L31
L30:
	;
	v90 = int32(0)
	goto L31
L31:
	;
	v91 = v90 + v22
	v92 = int32(-1)
	v93 = F_mmap(m, v91, v86, v92)
	mBase = m.M
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	if base.B2i32(v97 != int32(2))|base.B2i32(v93 != v92) != 0 {
		v121 = v93
		v122 = v91
		v123 = v95
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v103 = int32(-1)
	v106 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v106 == int32(0) {
		v121 = v103
		v122 = v91
		v123 = v95
		goto L21
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+176)) = v91
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_7), v39+int32(176))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(627), int32(_a_F_CreateSharedMemoryAndSemaphores_9))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v121 = v103
	v122 = v91
	v123 = v95
	goto L21
L37:
	;
	v129 = int32(_a_F_CreateSharedMemoryAndSemaphores_4)
	goto L39
L38:
	;
	v129 = int32(_a_F_CreateSharedMemoryAndSemaphores_10)
	goto L39
L39:
	;
	F_SetConfigOption(m, int32(_a_F_CreateSharedMemoryAndSemaphores_3), v129, int32(0), int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v121 != int32(-1) {
		v145 = v121
		v146 = v122
		v147 = v123
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v145 == int32(-1) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[4]))
	if v137 == int32(1) {
		v145 = v121
		v146 = v122
		v147 = v123
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v142 = F_mmap(m, v22, int32(33), int32(-1))
	mBase = m.M
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	v145 = v142
	v146 = v22
	v147 = v144
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = v147
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[8])) = v146
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[9])) = v145
	F_on_shmem_exit(m, int32(911), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_11), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v147 == int32(48) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v146
	F_errhint(m, int32(_a_F_CreateSharedMemoryAndSemaphores_12), v39+int32(16))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(663), int32(_a_F_CreateSharedMemoryAndSemaphores_9))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v188 = int32(40)
	v190 = v146
	goto L17
L55:
	;
	v188 = v22
	v190 = v22
	goto L17
L56:
	;
	v209 = int32(0)
	v210 = int32(_a_F_CreateSharedMemoryAndSemaphores_13)
	v216 = F___strchrnul(m, v210, int32(61))
	mBase = m.M
	if v210 == v216 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v258 != 0 {
		goto L74
	} else {
		goto L75
	}
L59:
	;
	v258 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v219 = v216 - v210
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+uint32(_c_F_CreateSharedMemoryAndSemaphores[10]))))
	if v221 != 0 {
		v252 = v209
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v258 = v252
	goto L58
L63:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[11]))
	if v223 == int32(0) {
		v252 = v209
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	if v226 == int32(0) {
		v252 = v209
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v230 = v223
	v231 = v226
	goto L66
L66:
	;
	v234 = F_strncmp(m, v210, v231, v219)
	mBase = m.M
	if v234 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v252 = v238 + int32(1)
	goto L62
L68:
	;
	goto L67
L69:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v238 = v237 + v219
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v239 == int32(61) {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v243 != 0 {
		v230 = v230 + int32(4)
		v231 = v243
		goto L66
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	v252 = v209
	goto L62
L74:
	;
	v259 = int32(0)
	v262 = F_strtox_2(m, v258, v259, v259, int64(4294967295))
	mBase = m.M
	goto L77
L75:
	;
	v264 = v209
	goto L76
L76:
	;
	v266 = F_pgmem_shmget(m, v197, v188, int32(1920))
	mBase = m.M
	if v266 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v264 = base.I32_wrap_i64(v262)
	goto L76
L78:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v39)+188))
	if v777 == int32(0) {
		v197 = v765
		goto L56
	} else {
		goto L189
	}
L79:
	;
	v765 = v197 + int32(1)
	goto L78
L80:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L186
	}
L81:
	;
	v707 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	v708 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+16)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(679834894)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v707
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v39)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v371)+32)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v371)+24)) = v713
	*(*int32)(unsafe.Add(mBase, uint32(v371)+12)) = int32(40)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+8)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(12)))) = v371
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[12])) = v197
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[13])) = v371
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[9]))
	if v725 == v708 {
		goto L183
	} else {
		goto L184
	}
L82:
	;
	v392 = int32(0)
	v393 = F_pgmem_shmget(m, v197, int32(40), v392)
	mBase = m.M
	if v393 < v392 {
		goto L115
	} else {
		goto L116
	}
L83:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	switch v270 - int32(2) {
	case 0, 18, 22:
		goto L82
	default:
		goto L88
	case 26:
		goto L89
	}
L84:
	;
	goto L85
L85:
	;
	F_on_shmem_exit(m, int32(912), v266)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L109
	}
L86:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(248), int32(_a_F_CreateSharedMemoryAndSemaphores_14))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L108
	}
L87:
	;
	F_errhint(m, v357, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L107
	}
L88:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L103
	}
L89:
	;
	v273 = int32(0)
	v275 = F_pgmem_shmget(m, v197, v273, int32(1920))
	mBase = m.M
	if v275 < v273 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7])) = int32(28)
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L100
	}
L91:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v279))|base.B2i32(int32(1)<<(uint(v279)%32)&int32(17825796) == int32(0)) != 0 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v289 = int32(0)
	v291 = F_pgmem_shmctl(m, v275, v289, v289)
	mBase = m.M
	if v289 <= v291 {
		goto L90
	} else {
		goto L95
	}
L94:
	;
	goto L82
L95:
	;
	v296 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v296 == int32(0) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v275
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_15), v39+int32(128))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(209), int32(_a_F_CreateSharedMemoryAndSemaphores_14))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L90
L100:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_16), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+116)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v197
	F_errdetail(m, int32(_a_F_CreateSharedMemoryAndSemaphores_17), v39+int32(112))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v357 = int32(_a_F_CreateSharedMemoryAndSemaphores_18)
	goto L87
L103:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_16), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = int32(1920)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v197
	F_errdetail(m, int32(_a_F_CreateSharedMemoryAndSemaphores_17), v39+int32(32))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	switch v270 - int32(48) {
	case 0:
		v357 = int32(_a_F_CreateSharedMemoryAndSemaphores_19)
		goto L87
	default:
		goto L86
	case 3:
		goto L106
	}
L106:
	;
	v357 = int32(_a_F_CreateSharedMemoryAndSemaphores_20)
	goto L87
L107:
	;
	goto L86
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v371 = F_pgmem_shmat(m, v266, v264)
	mBase = m.M
	if v371 == int32(-1) {
		goto L80
	} else {
		goto L110
	}
L110:
	;
	F_on_shmem_exit(m, int32(913), v371)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+160)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v39)+164)) = v266
	v380 = v39 + int32(288)
	v384 = F_pg_sprintf(m, v380, int32(_a_F_CreateSharedMemoryAndSemaphores_21), v39+int32(160))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_AddToDataDirLockFile(m, int32(7), v380)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v371 != 0 {
		goto L81
	} else {
		goto L114
	}
L114:
	;
	goto L82
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+188)) = int32(0)
	goto L79
L116:
	;
	goto L117
L117:
	;
	v398 = int32(0)
	v400 = v39 + int32(188)
	v403 = m.G0
	v405 = v403 - int32(192)
	m.G0 = v405
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v398
	v409 = int32(2)
	v413 = F_pgmem_shmctl(m, v393, v409, v405+int32(104))
	mBase = m.M
	if v413 < v398 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v39)+188))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	if v508 != 0 {
		goto L149
	} else {
		goto L150
	}
L119:
	;
	v491 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L145
	}
L120:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L140
	}
L121:
	;
	switch v456 - int32(2) {
	case 0:
		goto L119
	case 1:
		goto L79
	case 2:
		goto L118
	default:
		goto L120
	}
L122:
	;
	m.G0 = v405 + int32(192)
	goto L121
L123:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	switch v417 - int32(2) {
	case 0:
		goto L127
	default:
		goto L126
	case 22, 26:
		v456 = v409
		goto L122
	}
L124:
	;
	goto L125
L125:
	;
	v422 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	v427 = F_stat(m, v424, v405+int32(8))
	mBase = m.M
	if v427 < v422 {
		v456 = v422
		goto L122
	} else {
		goto L128
	}
L126:
	;
	v456 = int32(0)
	goto L122
L127:
	;
	v456 = int32(3)
	goto L122
L128:
	;
	v430 = F_pgmem_shmat(m, v393, v398)
	mBase = m.M
	if v430 == int32(-1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v433 = int32(2)
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[7]))
	switch v435 - v433 {
	case 0:
		goto L133
	default:
		goto L132
	case 22, 26:
		v456 = v433
		goto L122
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = v430
	v441 = int32(3)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if v442 != int32(679834894) {
		v456 = v441
		goto L122
	} else {
		goto L134
	}
L132:
	;
	v456 = int32(0)
	goto L122
L133:
	;
	v456 = int32(3)
	goto L122
L134:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v430)+24))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	if v445 != v446 {
		v456 = v441
		goto L122
	} else {
		goto L135
	}
L135:
	;
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v430)+32))
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v405)+96))
	if v448 != v449 {
		v456 = v441
		goto L122
	} else {
		goto L136
	}
L136:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v405)+176))
	if v453 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v454 = int32(1)
	goto L139
L138:
	;
	v454 = int32(4)
	goto L139
L139:
	;
	v456 = v454
	goto L122
L140:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+84)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v197
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_22), v39+int32(80))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v477
	F_errhint(m, int32(_a_F_CreateSharedMemoryAndSemaphores_23), v39-int32(-64))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(802), int32(_a_F_CreateSharedMemoryAndSemaphores_24))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	if v491 == int32(0) {
		v765 = v197
		goto L78
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v197
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_25), v39+int32(96))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(814), int32(_a_F_CreateSharedMemoryAndSemaphores_24))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v765 = v197
	goto L78
L149:
	;
	v509 = m.G0
	v511 = v509 - int32(48)
	m.G0 = v511
	v513 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v511)+44)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v511)+40)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v511)+36)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v511)+32)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v511)+28)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v511)+24)) = v513
	v534 = F_dsm_impl_op(m, int32(1), v508, v513, v511+int32(36), v511+int32(44), v511+int32(28), int32(14))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v701 = int32(0)
	v703 = F_pgmem_shmctl(m, v393, v701, v701)
	mBase = m.M
	v765 = int32(base.Ui32(v703)>>(uint(int32(31))%32)) + v197
	goto L78
L152:
	;
	if v534 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v536 = int32(2)
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v511)+28))
	if base.Ui32(v537) < base.Ui32(int32(12)) {
		v648 = v536
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	m.G0 = v511 + int32(48)
	goto L151
L156:
	;
	v666 = F_dsm_impl_op(m, v648, v508, int32(0), v511+int32(36), v511+int32(44), v511+int32(28), int32(15))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L181
	}
L157:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v511)+44))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	if v541 != int32(-1706017486) {
		v648 = v536
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v540)+8))
	if base.Ui64(base.I64_extend_i32_u(v537)) < base.Ui64(base.I64_extend_i32_u(v545)*int64(24)+int64(12)) {
		v648 = v536
		goto L156
	} else {
		goto L159
	}
L159:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if base.Ui32(v545) < base.Ui32(v552) {
		v648 = v536
		goto L156
	} else {
		goto L160
	}
L160:
	;
	if v552 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v562 = int32(0)
	goto L164
L162:
	;
	goto L163
L163:
	;
	v627 = int32(3)
	v630 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L177
	}
L164:
	;
	v574 = v540 + int32(12) + v562*int32(24)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	if v575 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	goto L163
L166:
	;
	v610 = v562 + int32(1)
	if v610 != v552 {
		v562 = v610
		goto L164
	} else {
		goto L176
	}
L167:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	if v578&int32(1) != 0 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v583 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	if v583 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v511)+20)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(v511)+16)) = v578
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_26), v511+int32(16))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v606 = F_dsm_impl_op(m, int32(3), v578, int32(0), v511+int32(32), v511+int32(40), v511+int32(24), int32(15))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), int32(294), int32(_a_F_CreateSharedMemoryAndSemaphores_28))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	goto L166
L176:
	;
	goto L165
L177:
	;
	if v630 == int32(0) {
		v648 = v627
		goto L156
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v511))) = v508
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_29), v511)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), int32(304), int32(_a_F_CreateSharedMemoryAndSemaphores_28))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v648 = v627
	goto L156
L181:
	;
	goto L155
L182:
	;
	m.G0 = v39 + int32(352)
	goto L9
L183:
	;
	v738 = v371
	goto L182
L184:
	;
	goto L185
L185:
	;
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v371)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v725)+32)) = v728
	v730 = *(*int64)(unsafe.Add(mBase, uint32(v371)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v725)+24)) = v730
	v732 = *(*int64)(unsafe.Add(mBase, uint32(v371)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v725)+16)) = v732
	v734 = *(*int64)(unsafe.Add(mBase, uint32(v371)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v725)+8)) = v734
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
	*(*int64)(unsafe.Add(mBase, uint32(v725))) = v736
	v738 = v725
	goto L182
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+152)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+148)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v39)+144)) = v266
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_30), v39+int32(144))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(259), int32(_a_F_CreateSharedMemoryAndSemaphores_14))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v780 = F_pgmem_shmdt(m, v777)
	mBase = m.M
	if int32(0) <= v780 {
		v197 = v765
		goto L56
	} else {
		goto L190
	}
L190:
	;
	v785 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	if v785 == int32(0) {
		v197 = v765
		goto L56
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v777
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_31), v39+int32(48))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(839), int32(_a_F_CreateSharedMemoryAndSemaphores_24))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v197 = v765
	goto L56
L195:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_32), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(732), int32(_a_F_CreateSharedMemoryAndSemaphores_24))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v823
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_33), v39)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_8), int32(718), int32(_a_F_CreateSharedMemoryAndSemaphores_24))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v842 = m.G0
	v844 = v842 - int32(112)
	m.G0 = v844
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	v852 = F___fstatat(m, int32(-100), v847, v844+int32(16), int32(0))
	mBase = m.M
	goto L204
L204:
	;
	if v852 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v873 = F_mul_size(m, v841, int32(128))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L1
	} else {
		goto L212
	}
L208:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v862
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_33), v844)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_34), int32(210), int32(_a_F_CreateSharedMemoryAndSemaphores_35))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	v875 = m.G0
	v877 = v875 - int32(16)
	m.G0 = v877
	v880 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)+12))
	v885 = (v873 + int32(7)) & int32(-8)
	v886 = v881 + v885
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v880)+8))
	if base.Ui32(v887) < base.Ui32(v886) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_36), v1231)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L292
	}
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v880)+12)) = v886
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	m.G0 = v877 + int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[14])) = v899 + v881
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[15])) = v841
	v909 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[16])) = v909
	F_on_shmem_exit(m, int32(910), v909)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	F_errcode(m, int32(_a_F_CreateSharedMemoryAndSemaphores_37))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v885
	v1231 = v877
	goto L213
L219:
	;
	m.G0 = v844 + int32(112)
	v918 = m.G0
	v920 = v918 - int32(16)
	m.G0 = v920
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[2]))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v923)+12))
	v926 = v924 + int32(8)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v923)+8))
	if base.Ui32(v927) < base.Ui32(v926) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923)+12)) = v926
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[1]))
	v942 = v941 + v924
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[17])) = v942
	v944 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v942))), uint32(v944))
	*(*int32)(unsafe.Add(mBase, uint32(v923)+20)) = v944
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v923)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v923)+12)) = (v923+v949+int32(127))&int32(-128) - v923
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[18])) = v944
	m.G0 = v920 + int32(16)
	F_CreateOrAttachShmemStructs(m)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	F_errcode(m, int32(_a_F_CreateSharedMemoryAndSemaphores_37))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920))) = int32(8)
	v1231 = v920
	goto L213
L225:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v966 = m.G0
	v968 = v966 - int32(1120)
	m.G0 = v968
	*(*int32)(unsafe.Add(mBase, uint32(v968)+76)) = int32(0)
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[19]))
	if v973 == int32(4) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[20]))
	if v1222 != 0 {
		goto L288
	} else {
		goto L289
	}
L227:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L1
	} else {
		goto L284
	}
L228:
	;
	v977 = F_AllocateDir(m, int32(_a_F_CreateSharedMemoryAndSemaphores_38))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[21]))
	v1116 = v1112*int32(5) - int32(-64)
	v1119 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L265
	}
L231:
	;
	v980 = F_ReadDir(m, v977, int32(_a_F_CreateSharedMemoryAndSemaphores_38))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	if v980 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v986 = v980
	goto L236
L234:
	;
	goto L235
L235:
	;
	F_FreeDir(m, v977)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L264
	}
L236:
	;
	v998 = v986 + int32(19)
	v999 = int32(_a_F_CreateSharedMemoryAndSemaphores_39)
	goto L240
L237:
	;
	goto L235
L238:
	;
	if v1037-v1038 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L240:
	;
	goto L241
L241:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	if v1006 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1007 = v998
	v1008 = v999
	v1009 = int32(5)
	v1010 = v1006
	goto L246
L243:
	;
	v1033 = v999
	v1037 = int32(0)
	goto L244
L244:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	goto L238
L245:
	;
	v1033 = v1028
	v1037 = v1030
	goto L244
L246:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	if base.B2i32(v1010 != v1012)|base.B2i32(v1012 == int32(0)) != 0 {
		v1028 = v1008
		v1030 = v1010
		goto L245
	} else {
		goto L248
	}
L247:
	;
	v1028 = v1022
	v1030 = int32(0)
	goto L245
L248:
	;
	v1018 = v1009 - int32(1)
	if v1018 == int32(0) {
		v1028 = v1008
		v1030 = v1010
		goto L245
	} else {
		goto L249
	}
L249:
	;
	v1021 = int32(1)
	v1022 = v1008 + v1021
	v1023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+1)))
	if v1023 != 0 {
		v1007 = v1007 + v1021
		v1008 = v1022
		v1009 = v1018
		v1010 = v1023
		goto L246
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+64)) = v998
	v1050 = v968 + int32(80)
	v1055 = F_pg_snprintf(m, v1050, int32(1036), int32(_a_F_CreateSharedMemoryAndSemaphores_40), v968-int32(-64))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1077 = F_ReadDir(m, v977, int32(_a_F_CreateSharedMemoryAndSemaphores_38))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L262
	}
L254:
	;
	v1059 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	if v1059 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+48)) = v1050
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_41), v968+int32(48))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1074 = F_unlink(m, v968+int32(80))
	mBase = m.M
	if v1074 != 0 {
		goto L227
	} else {
		goto L261
	}
L259:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), int32(337), int32(_a_F_CreateSharedMemoryAndSemaphores_42))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	goto L253
L262:
	;
	if v1077 != 0 {
		v986 = v1077
		goto L236
	} else {
		goto L263
	}
L263:
	;
	goto L237
L264:
	;
	goto L230
L265:
	;
	if v1119 != 0 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+16)) = v1116
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_43), v968+int32(16))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v1135 = v1116*int32(24) + int32(12)
	goto L271
L269:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), int32(198), int32(_a_F_CreateSharedMemoryAndSemaphores_44))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	v1153 = Fn13986(m, int64(32))
	mBase = m.M
	goto L273
L272:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v968)+76))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[22])) = v1170
	F_on_shmem_exit(m, int32(1097), v965)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L1
	} else {
		goto L277
	}
L273:
	;
	v1155 = v1153 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[23])) = v1155
	if v1155 == int32(0) {
		goto L271
	} else {
		goto L274
	}
L274:
	;
	v1165 = F_dsm_impl_op(m, int32(0), v1155, v1135, int32(_a_F_CreateSharedMemoryAndSemaphores_45), v968+int32(76), int32(_a_F_CreateSharedMemoryAndSemaphores_46), int32(21))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	if v1165 == int32(0) {
		goto L271
	} else {
		goto L276
	}
L276:
	;
	goto L272
L277:
	;
	v1177 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	if v1177 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+4)) = v1135
	v1181 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v968))) = v1181
	F_errmsg_internal(m, int32(_a_F_CreateSharedMemoryAndSemaphores_47), v968)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v965)+16)) = v1192
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_CreateSharedMemoryAndSemaphores[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+8)) = v1116
	*(*int64)(unsafe.Add(mBase, uint32(v1195))) = int64(2588949810)
	m.G0 = v968 + int32(1120)
	goto L226
L282:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), int32(223), int32(_a_F_CreateSharedMemoryAndSemaphores_44))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+32)) = v968 + int32(80)
	F_errmsg(m, int32(_a_F_CreateSharedMemoryAndSemaphores_48), v968+int32(32))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_27), int32(343), int32(_a_F_CreateSharedMemoryAndSemaphores_42))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	m.T0[v1222].(func(*base.Module))(m)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	m.G0 = v18 + int32(16)
	return
L291:
	;
	goto L290
L292:
	;
	F_errfinish(m, int32(_a_F_CreateSharedMemoryAndSemaphores_49), int32(258), int32(_a_F_CreateSharedMemoryAndSemaphores_50))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DeleteSharedComments(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13847(m, l0, l1, int32(2397), int32(2396))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_LockSharedObjectForSession(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(1262)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v2
	v19 = F_LockAcquire(m, v5, int32(8), int32(1), v2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_SharedFileSetAttach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = int32(44)
	v6 = l0 + v5
	v9 = base.AtomicRmwXchg32(m, l0, v5, int32(1))
	if v9 != 0 {
		F_s_lock(m, v6, int32(_a_F_SharedFileSetAttach_0), int32(60), int32(_a_F_SharedFileSetAttach_1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
				v19 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+44)), uint32(v19))
				F_on_dsm_detach(m, l1, int32(1096), l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				v25 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6))), uint32(v25))
				F_errstart_cold(m, int32(21), v25)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_SharedFileSetAttach_2), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SharedFileSetAttach_0), int32(73), int32(_a_F_SharedFileSetAttach_1))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v15 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
			v19 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+44)), uint32(v19))
			F_on_dsm_detach(m, l1, int32(1096), l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		} else {
			v25 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v6))), uint32(v25))
			F_errstart_cold(m, int32(21), v25)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_SharedFileSetAttach_2), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SharedFileSetAttach_0), int32(73), int32(_a_F_SharedFileSetAttach_1))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_UnlockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v9)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v4
	v18 = F_LockRelease(m, v7, l2, v4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_deleteSharedDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v7 = F_table_open(m, int32(1214), int32(3))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = int32(0)
		F_shdepDropDependency(m, v7, l0, l1, l2, base.B2i32(l2 == v9), v9, v9, v9)
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_relation_close(m, v7, int32(3))
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_shared_dependency_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v6) < base.Ui32(v7) {
		return int32(-1)
	} else {
		v11 = int32(1)
		if base.Ui32(v7) < base.Ui32(v6) {
			v34 = v11
			return v34
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui32(v13) < base.Ui32(v14) {
				return int32(-1)
			} else {
				if base.Ui32(v14) < base.Ui32(v13) {
					v34 = v11
					return v34
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v19) < base.Ui32(v20) {
						return int32(-1)
					} else {
						if base.Ui32(v20) < base.Ui32(v19) {
							v34 = v11
						} else {
							v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+12)))
							v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+12)))
							if v26 < v27 {
								v34 = int32(-1)
							} else {
								v34 = base.B2i32(v27 < v26)
							}
						}
						return v34
					}
				}
			}
		}
	}
}
